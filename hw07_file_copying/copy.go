package main

import (
	"errors"
	"io"
	"os"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	chunkSize                int64 = 1024
	ErrUnsupportedFile             = errors.New("unsupported file")
	ErrOffsetExceedsFileSize       = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	stat, err := os.Stat(fromPath)

	if err != nil {
		return err
	}
	if !stat.Mode().IsRegular() || stat.Size() <= 0 {
		return ErrUnsupportedFile
	}

	if offset > stat.Size() {
		return ErrOffsetExceedsFileSize
	}

	var realLimit int64
	switch {
	case limit <= 0:
		realLimit = stat.Size()
	case limit > stat.Size()-offset:
		realLimit = stat.Size() - offset
	default:
		realLimit = limit
	}

	originalFile, err := os.OpenFile(fromPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer originalFile.Close()

	_, err = originalFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	destinationFile, err := os.OpenFile(toPath, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	progressBar := pb.StartNew(int(realLimit))
	defer progressBar.Finish()

	buffer := make([]byte, chunkSize)
	var total int64

	for total < realLimit {
		realChunkSize := chunkSize
		if (total + chunkSize) > realLimit {
			realChunkSize = realLimit - total
		}

		n, e := originalFile.Read(buffer[:realChunkSize])
		if e != nil && errors.Is(e, io.EOF) && errors.Is(e, io.ErrUnexpectedEOF) {
			return e
		}

		if n == 0 {
			break
		}

		_, err = destinationFile.Write(buffer[:n])
		if e != nil {
			return e
		}

		total += int64(n)

		progressBar.Add(n)
	}

	// Place your code here.
	return nil
}
