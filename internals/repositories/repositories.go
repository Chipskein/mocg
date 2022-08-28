package repositories

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type File struct {
	Name      string
	Size      int64
	FullPath  string
	Extension string
}

var DEFAULT_DIRECTORY = "/home/chipskein/Music"
var CURRENT_DIRECTORY string

func GetAllFilesFromLocalDirectory(DIRECTORY string) (map[string]File, error) {

	fmt.Println(fmt.Sprintf("WORKING WITH DIRECTORY => %s", DIRECTORY))
	files, err := ioutil.ReadDir(DIRECTORY)
	if err != nil {
		fmt.Println(fmt.Sprintf("[ERROR] %s Directory not found", DIRECTORY))
		DIRECTORY = DEFAULT_DIRECTORY
		fmt.Println(fmt.Sprintf("[WARNING] USING default directory insted %s ", DIRECTORY))
	}
	files, err = ioutil.ReadDir(DIRECTORY)
	if err != nil {
		return nil, err
	}

	Files := make(map[string]File)
	for _, file := range files {
		if !file.IsDir() {
			var filename = file.Name()

			var fullpath = fmt.Sprintf("%s/%s", DIRECTORY, filename)
			var extension = filepath.Ext(fullpath)
			if !HasSupport(extension) {
				//fmt.Printf("%s: Extension not supported\n", filename)
				continue
			}
			var filesize = file.Size()
			var MusicFile = File{filename, filesize, fullpath, extension}
			Files[filename] = MusicFile
		}
	}
	fmt.Println(fmt.Sprintf("Found %d files in directory %s", len(Files), DIRECTORY))
	return Files, nil
}
func HasSupport(extension string) bool {
	switch extension {
	case ".ogg":
		return true
	default:
		return false
	}
}
