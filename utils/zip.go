package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

func Compress(src string, dest *os.File) error {

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	gw := gzip.NewWriter(dest)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	return compress(srcFile, "", tw)
}

// compress 打包成并返回 writer 流
func compress(file *os.File, prefix string, tw *tar.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.ReadDir(-1)
		if err != nil {
			return err
		}
		for _, fileInfo := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fileInfo.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = prefix + "/" + info.Name()
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
