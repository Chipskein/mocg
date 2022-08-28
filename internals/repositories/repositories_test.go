package repositories

import (
	"testing"
)

type TestDirectories struct {
	repository LocalRepository
	shouldFail bool
}

func TestReadDirectory(t *testing.T) {
	var tests = []TestDirectories{
		{LocalRepository{"/home/chipskein/Music/", "/home/chipskein/sources/mocg/audios/", nil}, false},
		{LocalRepository{"/home/chipskein/Music/", "/home/chipskein/fasojfklasjfklasjhfjkahsf", nil}, false},
		{LocalRepository{"/home/chipskein/Music/", "/", nil}, false},
		{LocalRepository{"/home/chipskein/Music/", "/home/chipskein/", nil}, false},
		{LocalRepository{"/home/chipskein/Music/", "../../audios", nil}, false},
		{LocalRepository{"/home/chipskein/Music/", "", nil}, false},
		{LocalRepository{"/wrong/default/path", "fhjasjfhajkshfjkashfkjahfjkasf", nil}, true},
		{LocalRepository{"/another/wrong/default/path", "fjkasjfkasjfkjas", nil}, true},
		{LocalRepository{"../../audios/", "/invalid/directory/convert", nil}, false},
	}

	for _, test := range tests {
		files, err := test.repository.ReadDirectoryOrDefault()
		if err != nil {
			t.Errorf("\n[ERROR]CURRENT_DIRECTORY %s \nShould Fail:%t ", test.repository.CURRENT_DIRECTORY, test.shouldFail)
			t.Errorf("\n[ERROR]DEFAULT_DIRECTORY %s \nShould Fail:%t ", test.repository.DEFAULT_DIRECTORY, test.shouldFail)
			continue
		}
		if files != nil {
			t.Logf("\nFound %d DIRECTORY %s\nShould Fail:%t", len(files), test.repository.CURRENT_DIRECTORY, test.shouldFail)
		}
	}
}
func TestListFiles(t *testing.T) {
	var tests = []TestDirectories{
		{LocalRepository{"/home/chipskein/Music", "/home/chipskein/sources/mocg/audios/", nil}, false},
		{LocalRepository{"/home/chipskein/Music", "../", nil}, false},
		{LocalRepository{"/home/chipskein/Music", "/wrong/path", nil}, false},
	}
	for _, test := range tests {
		err := test.repository.ListFiles()
		if err != nil {
			t.Errorf("\n[ERRROR]Could not read Directory %s\n", test.repository.CURRENT_DIRECTORY)
			continue
		}
		for _, file := range test.repository.Files {
			t.Logf("\nFilename:%s\nExtension:%s\nIsADirectory:%t\nFullPath:%s\nSize:%d\n", file.Name, file.Extension, file.IsADirectory, file.FullPath, file.Size)
		}
	}
}
