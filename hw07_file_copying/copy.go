package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	srcSize := fromInfo.Size()
	if srcSize == 0 {
		return ErrUnsupportedFile
	}
	if offset > srcSize {
		return ErrOffsetExceedsFileSize
	}
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	if offset != 0 {
		_, err = fromFile.Seek(offset, io.SeekStart)
		if err != nil {
			return err
		}
	}
	dstFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	bar := initProgressBar(srcSize, limit, offset)
	barReader := bar.NewProxyReader(fromFile)
	defer bar.Finish()
	if limit != 0 {
		barReader = bar.NewProxyReader(io.LimitReader(fromFile, limit))
	}
	_, err = io.Copy(dstFile, barReader)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}

func defineBarSize(fileSize int64, limit int64, offset int64) int64 {
	if offset+limit > fileSize {
		return fileSize - offset
	}
	if limit != 0 {
		return limit
	}
	return fileSize - offset
}

func initProgressBar(srcSize int64, limit int64, offset int64) *pb.ProgressBar {
	barSize := defineBarSize(srcSize, limit, offset)

	return pb.Full.Start64(barSize)
}
