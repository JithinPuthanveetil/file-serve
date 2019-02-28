package fileserv

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/file-serve/model"
)

// DownloadFile enables download
func (r *ShowFileRequest) DownloadFile() (*model.Server, error) {
	file, err := os.Open(r.Root)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if !fileInfo.IsDir() {
		buf := make([]byte, fileInfo.Size())
		_, err = file.Read(buf)
		if err != nil {
			return nil, err
		}
		root := filepath.Base(fileInfo.Name())
		return &model.Server{
			Root:       strings.Replace(root, " ", "_", strings.Count(root, " ")),
			Data:       buf,
			IsDownload: true,
		}, nil
	}

	return downloadDirectory(r.Root)
}

func downloadDirectory(root string) (*model.Server, error) {
	zipRoot := strings.Replace(filepath.Base(root), " ", "_", strings.Count(filepath.Base(root), " ")) + ".zip"

	file, err := os.Create(zipRoot)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)

	err = addFiles(zipWriter, root, "")
	if err != nil {
		return nil, err
	}

	err = zipWriter.Close()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(zipRoot)
	if err != nil {
		return nil, err
	}
	err = os.Remove(zipRoot)
	if err != nil {
		return nil, err
	}

	return &model.Server{
		Root:       zipRoot,
		Data:       data,
		IsDownload: true,
	}, nil
}

func addFiles(zipWriter *zip.Writer, root, path string) error {
	fileInfo, err := os.Stat(root)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(root)
		if err != nil {
			return err
		}

		for _, f := range files {
			err := addFiles(zipWriter, root+"/"+f.Name(), path+"/"+f.Name())
			if err != nil {
				return err
			}
		}
	} else {
		data, err := ioutil.ReadFile(root)
		if err != nil {
			return err
		}
		f, err := zipWriter.Create(path)
		if err != nil {
			return err
		}

		_, err = f.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}
