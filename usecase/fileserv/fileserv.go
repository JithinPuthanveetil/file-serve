package fileserv

import (
	"os"
	"strings"

	"github.com/file-serve/model"
)

// ShowFileRequest request
type ShowFileRequest struct {
	Root string
}

// ShowFile serve files
func (r *ShowFileRequest) ShowFile() (*model.Server, error) {
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

		return &model.Server{
			Root: fileInfo.Name(),
			Data: buf,
		}, nil
	}

	return showDirectory(file)
}

func showDirectory(file *os.File) (*model.Server, error) {
	var a model.Server

	f, err := file.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, v := range f {
		if !strings.HasPrefix(v.Name(), ".") {
			if v.IsDir() {
				a.FolderName = append(a.FolderName, v.Name())
			} else {
				a.FileName = append(a.FileName, v.Name())
			}
			a.Root = file.Name()
		}
	}

	return &a, nil
}
