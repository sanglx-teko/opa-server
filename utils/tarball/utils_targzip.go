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

func Compress_tarball(destinationfile string, sourcedir string) (err error) {
	if destinationfile == "" {
		err = fmt.Errorf("Usage : gotar destinationfile.tar.gz source")
		return
	}
	if sourcedir == "" {
		err = fmt.Errorf("Usage : gotar destinationfile.tar.gz source-directory")
		return
	}

	dir, err := os.Open(sourcedir)
	defer dir.Close()

	if err != nil {
		return
	}

	files, err := dir.Readdir(0) // grab the files list

	if err != nil {
		return
	}

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

	for _, fileInfo := range files {

		if fileInfo.IsDir() {
			continue
		}
		file, err := os.Open(dir.Name() + string(filepath.Separator) + fileInfo.Name())
		defer file.Close()

		if err != nil {
			break
		}

		header := new(tar.Header)
		header.Name = file.Name()
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
