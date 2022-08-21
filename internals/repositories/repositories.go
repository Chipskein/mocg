package repositories

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type File struct {
	Name      string
	Size      int64
	FullPath  string
	Extension string
}

const DEFAULT_DIRECTORY = "/home/chipskein/Music"

func GetAllFilesFromLocalDirectory(DIRECTORY string) map[string]File {

	fmt.Println(fmt.Sprintf("WORKING WITH DIRECTORY => %s", DIRECTORY))
	files, err := ioutil.ReadDir(DIRECTORY)
	if err != nil {
		fmt.Println(fmt.Sprintf("[ERROR] %s Directory not found", DIRECTORY))
		DIRECTORY = DEFAULT_DIRECTORY
		fmt.Println(fmt.Sprintf("[WARNING] USING default directory insted %s ", DIRECTORY))

	}
	files, err = ioutil.ReadDir(DIRECTORY)
	if err != nil {
		log.Panic(fmt.Sprintf("Could not Read DEFAULT DIRECTORY %s", DIRECTORY))
	}

	Files := make(map[string]File)
	for _, file := range files {
		if !file.IsDir() {
			var filename = file.Name()
			var filesize = file.Size()
			var fullpath = fmt.Sprintf("%s/%s", DIRECTORY, filename)
			var extension = filepath.Ext(fullpath)
			var MusicFile = File{filename, filesize, fullpath, extension}
			Files[filename] = MusicFile
		}
	}
	fmt.Println(fmt.Sprintf("Found %d files in directory %s", len(Files), DIRECTORY))
	return Files
}
