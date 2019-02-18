package model

// Server struct
type Server struct {
	Root       string
	FolderName []string
	FileName   []string
	Data       []byte
	IsDownload bool
}
