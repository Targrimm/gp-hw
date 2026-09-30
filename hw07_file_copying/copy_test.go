package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	maxInt64 := int64(^uint(0) >> 1)

	tests := []struct {
		name   string
		input  string
		file   string
		offset int64
		limit  int64
		err    error
	}{
		{
			name:   "Success Offset 0 Limit 0",
			input:  "testdata/input.txt",
			file:   "testdata/out_offset0_limit0.txt",
			offset: 0,
			limit:  0,
		},
		{
			name:   "Success Limit over original file",
			input:  "testdata/input.txt",
			file:   "testdata/out_offset0_limit0.txt",
			offset: 0,
			limit:  maxInt64,
		},
		{
			name:   "Success Offset 0 Limit 10",
			input:  "testdata/input.txt",
			file:   "testdata/out_offset0_limit10.txt",
			offset: 0,
			limit:  10,
		},
		{
			name:   "Success Offset 0 Limit 1000",
			input:  "testdata/input.txt",
			file:   "testdata/out_offset0_limit1000.txt",
			offset: 0,
			limit:  1000,
		},
		{
			name:   "Success Offset 0 Limit 10000",
			input:  "testdata/input.txt",
			file:   "testdata/out_offset0_limit10000.txt",
			offset: 0,
			limit:  10000,
		},
		{
			name:   "Success Offset 100 Limit 1000",
			input:  "testdata/input.txt",
			file:   "testdata/out_offset100_limit1000.txt",
			offset: 100,
			limit:  1000,
		},
		{
			name:   "Success Offset 6000 Limit 0",
			input:  "testdata/input.txt",
			file:   "testdata/out_offset6000_limit1000.txt",
			offset: 6000,
			limit:  1000,
		},
		{
			name:   "Fail Offset more than original file size",
			input:  "testdata/input.txt",
			err:    ErrOffsetExceedsFileSize,
			offset: maxInt64,
			limit:  0,
		},
		{
			name:  "Fail Offset more than original file size",
			input: "testdata",
			err:   ErrUnsupportedFile,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outPath := filepath.Join(t.TempDir(), "test_out.txt")
			err := Copy(test.input, outPath, test.offset, test.limit)
			if test.err != nil {
				require.ErrorIs(t, err, test.err)
				return
			}
			require.NoError(t, err)
			result, err := os.ReadFile(outPath)
			require.NoError(t, err)
			testResult, err := os.ReadFile(test.file)
			require.NoError(t, err)
			require.Equal(t, string(result), string(testResult))
			require.Equal(t, result, testResult)
		})
	}
}
