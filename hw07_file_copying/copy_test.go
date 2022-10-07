package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("offset exceeds file size", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/test1.txt", 6618, 0)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
	})
	t.Run("unsupported file", func(t *testing.T) {
		err := Copy("/dev/urandom", "/tmp/test2.txt", 0, 0)
		require.ErrorIs(t, err, ErrUnsupportedFile)
	})
	t.Run("file not found", func(t *testing.T) {
		err := Copy("testdata/never_existed_file.txt", "/tmp/test3.txt", 0, 0)
		require.ErrorIs(t, err, os.ErrNotExist)
	})
	t.Run("copy full file", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/test4.txt", 0, 0)
		require.NoError(t, err, os.ErrNotExist)
	})
	t.Run("copy file with offset > 0", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/test5.txt", 1000, 0)
		require.NoError(t, err, os.ErrNotExist)
	})
	t.Run("copy file with offset > 0 && limit > 0", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/test6.txt", 1000, 1000)
		require.NoError(t, err, os.ErrNotExist)
	})
	t.Run("copy file with limit > 0", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/test7.txt", 0, 1000)
		require.NoError(t, err, os.ErrNotExist)
	})
	t.Run("offset + limit > fileSize case", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/test8.txt", 4000, 3000)
		require.NoError(t, err, os.ErrNotExist)
	})
}

func Test_defineBarSize(t *testing.T) {
	tests := []struct {
		name     string
		fileSize int64
		limit    int64
		offset   int64
		want     int64
	}{
		{
			name:     "copy full file",
			fileSize: 10000,
			limit:    0,
			offset:   0,
			want:     10000,
		},
		{
			name:     "offset > 0, limit = 0, expected to exhaust limit",
			fileSize: 10000,
			limit:    0,
			offset:   1000,
			want:     9000,
		},
		{
			name:     "offset + limit < fileSize",
			fileSize: 10000,
			limit:    9000,
			offset:   0,
			want:     9000,
		},
		{
			name:     "offset + limit > fileSize",
			fileSize: 10000,
			limit:    8000,
			offset:   7000,
			want:     3000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defineBarSize(tt.fileSize, tt.limit, tt.offset); got != tt.want {
				t.Errorf("defineBarSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
