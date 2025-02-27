package test_common

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/grafana/agent/converter/diag"
	"github.com/grafana/agent/pkg/cluster"
	"github.com/grafana/agent/pkg/flow"
	"github.com/grafana/agent/pkg/flow/logging"
	"github.com/stretchr/testify/require"
)

const (
	flowSuffix  = ".river"
	diagsSuffix = ".diags"
)

// TestDirectory will execute tests for converting from a source configuration
// file to a flow configuration file for all files in a provided folder path.
//
// For each file in the folderPath which ends with the sourceSuffix:
//
//  1. Execute the convert func on the content of each file.
//  2. Remove an Info diags from the results of calling convert in step 1.
//  3. If the current filename.sourceSuffix has a matching filename.diags, read
//     the contents of filename.diags and validate that they match in order
//     with the diags from step 2.
//  4. If the current filename.sourceSuffix has a matching filename.river, read
//     the contents of filename.river and validate that they match the river
//     configuration generated by calling convert in step 1.
func TestDirectory(t *testing.T, folderPath string, sourceSuffix string, convert func(in []byte) ([]byte, diag.Diagnostics)) {
	require.NoError(t, filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, _ error) error {
		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, sourceSuffix) {
			tc := getTestCaseName(path, sourceSuffix)
			t.Run(tc, func(t *testing.T) {
				riverFile := strings.TrimSuffix(path, sourceSuffix) + flowSuffix
				diagsFile := strings.TrimSuffix(path, sourceSuffix) + diagsSuffix
				if !fileExists(riverFile) && !fileExists(diagsFile) {
					t.Fatalf("no expected diags or river for %s - missing test expectations?", path)
				}

				actualRiver, actualDiags := convert(getSourceContents(t, path))

				// Skip Info level diags for this testing. These would create
				// a lot of unnecessary noise.
				actualDiags.RemoveDiagsBySeverity(diag.SeverityLevelInfo)

				expectedDiags := getExpectedDiags(t, diagsFile)
				validateDiags(t, expectedDiags, actualDiags)

				expectedRiver := getExpectedRiver(t, riverFile)
				validateRiver(t, expectedRiver, actualRiver)
			})
		}

		return nil
	}))
}

// getSourceContents reads the source file and retrieve its contents.
func getSourceContents(t *testing.T, path string) []byte {
	sourceBytes, err := os.ReadFile(path)
	require.NoError(t, err)
	return sourceBytes
}

// getTestCaseName gets the test case name based on the path and source suffix.
func getTestCaseName(path string, sourceSuffix string) string {
	caseName := filepath.Base(path)
	return strings.TrimSuffix(caseName, sourceSuffix)
}

// getExpectedDiags will retrieve any expected diags for the test.
func getExpectedDiags(t *testing.T, diagsFile string) []string {
	expectedDiags := []string{}
	if _, err := os.Stat(diagsFile); err == nil {
		errorBytes, err := os.ReadFile(diagsFile)
		require.NoError(t, err)
		errorsString := string(normalizeLineEndings(errorBytes))
		expectedDiags = strings.Split(errorsString, "\n")

		// Some error messages have \n in them and need this
		for ix := range expectedDiags {
			expectedDiags[ix] = strings.ReplaceAll(expectedDiags[ix], "\\n", "\n")
		}
	}

	return expectedDiags
}

// validateDiags makes sure the expected diags and actual diags are a match
func validateDiags(t *testing.T, expectedDiags []string, actualDiags diag.Diagnostics) {
	for ix, diag := range actualDiags {
		if len(expectedDiags) > ix {
			if expectedDiags[ix] != diag.String() {
				printActualDiags(actualDiags)
			}
			require.Equal(t, expectedDiags[ix], diag.String())
		} else {
			printActualDiags(actualDiags)
			require.Fail(t, "unexpected diag count reach for diag: "+diag.String())
		}
	}

	// If we expect more diags than we got
	if len(expectedDiags) > len(actualDiags) {
		printActualDiags(actualDiags)
		require.Fail(t, "missing expected diag: "+expectedDiags[len(actualDiags)])
	}
}

func printActualDiags(actualDiags diag.Diagnostics) {
	fmt.Println("============== ACTUAL =============")
	fmt.Println(string(normalizeLineEndings([]byte(actualDiags.Error()))))
	fmt.Println("===================================")
}

// normalizeLineEndings will replace '\r\n' with '\n'.
func normalizeLineEndings(data []byte) []byte {
	normalized := bytes.ReplaceAll(data, []byte{'\r', '\n'}, []byte{'\n'})
	return normalized
}

// getExpectedRiver reads the expected river output file and retrieve its contents.
func getExpectedRiver(t *testing.T, filePath string) []byte {
	if _, err := os.Stat(filePath); err == nil {
		outputBytes, err := os.ReadFile(filePath)
		require.NoError(t, err)
		return normalizeLineEndings(outputBytes)
	}

	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// validateRiver makes sure the expected river and actual river are a match
func validateRiver(t *testing.T, expectedRiver []byte, actualRiver []byte) {
	if len(expectedRiver) > 0 {
		if !reflect.DeepEqual(expectedRiver, actualRiver) {
			fmt.Println("============== ACTUAL =============")
			fmt.Println(string(normalizeLineEndings(actualRiver)))
			fmt.Println("===================================")
		}

		require.Equal(t, string(expectedRiver), string(normalizeLineEndings(actualRiver)))

		attemptLoadingFlowConfig(t, actualRiver)
	}
}

// attemptLoadingFlowConfig will attempt to load the Flow config and report any errors.
func attemptLoadingFlowConfig(t *testing.T, river []byte) {
	cfg, err := flow.ReadFile(t.Name(), river)
	require.NoError(t, err, "the output River config failed to parse: %s", string(normalizeLineEndings(river)))

	// The below check suffers from test race conditions on Windows. Our goal here is to verify config conversions,
	// which is platform independent, so we can skip this check on Windows as a workaround.
	if runtime.GOOS == "windows" {
		return
	}

	logger, err := logging.New(os.Stderr, logging.DefaultOptions)
	require.NoError(t, err)
	f := flow.New(flow.Options{
		Logger:         logger,
		Clusterer:      &cluster.Clusterer{Node: cluster.NewLocalNode("")},
		DataPath:       t.TempDir(),
		HTTPListenAddr: ":0",
	})
	err = f.LoadFile(cfg, nil)

	// Many components will fail to build as e.g. the cert files are missing, so we ignore these errors.
	// This is not ideal, but we still validate for other potential issues.
	if err != nil && strings.Contains(err.Error(), "Failed to build component") {
		t.Log("ignoring error: " + err.Error())
		return
	}
	require.NoError(t, err, "failed to load the River config: %s", string(normalizeLineEndings(river)))
}
