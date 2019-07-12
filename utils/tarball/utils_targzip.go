package tarball

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CopyFile ...
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// CompressTarball ...
func CompressTarball(destinationfile string, sourcedir string) (err error) {
	if destinationfile == "" {
		err = fmt.Errorf("Usage : gotar destinationfile.tar.gz source")
		return
	}
	if sourcedir == "" {
		err = fmt.Errorf("Usage : gotar destinationfile.tar.gz source-directory")
		return
	}

	dir, err := os.Open(sourcedir)
	type tempFileInfoStruct struct {
		Path     string
		FileInfo os.FileInfo
	}
	files := []*tempFileInfoStruct{}
	defer dir.Close()
	if err = filepath.Walk(sourcedir, func(path string, info os.FileInfo, err error) error {
		files = append(files, &tempFileInfoStruct{
			Path:     path,
			FileInfo: info,
		})
		return nil
	}); err != nil {
		fmt.Println("could not walk into directory")
	}

	// if err != nil {
	// 	return
	// }

	// files, err := dir.Readdir(0) // grab the files list
	// if err != nil {
	// 	return
	// }

	tarfile, err := os.Create(destinationfile)
	defer tarfile.Close()

	if err != nil {
		return
	}

	var fileWriter io.WriteCloser = tarfile

	if strings.HasSuffix(destinationfile, ".gz") {
		fileWriter = gzip.NewWriter(tarfile) // add a gzip filter
		defer fileWriter.Close()             // if user add .gz in the destination filename
	}

	tarfileWriter := tar.NewWriter(fileWriter)
	defer tarfileWriter.Close()

	for _, f := range files {
		// fmt.Println("file info:", fileInfo)
		fileInfo := f.FileInfo
		if fileInfo.IsDir() {
			continue
		}
		// fmt.Println("file info name:", fileInfo)
		// fmt.Println("file name:", fileI + string(filepath.Separator) + fileInfo.Name())
		file, err := os.Open(f.Path)
		defer file.Close()

		if err != nil {
			break
		}
		// fmt.Println("file Name:", file.Name())

		header := new(tar.Header)
		name := file.Name()
		fileName := strings.Join(strings.Split(name, "/")[2:], "/")
		header.Name = fileName
		header.Size = fileInfo.Size()
		header.Mode = int64(fileInfo.Mode())
		header.ModTime = fileInfo.ModTime()

		err = tarfileWriter.WriteHeader(header)
		if err != nil {
			break
		}

		_, err = io.Copy(tarfileWriter, file)
		if err != nil {
			break
		}
	}

	return

}
