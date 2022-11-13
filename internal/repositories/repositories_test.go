package repositories

import (
	"testing"
)

type TestDirectories struct {
	repository LocalRepository
	shouldFail bool
}
type TestFiles struct {
	fullpath string
}

func TestReadDirectory(t *testing.T) {
	var tests = []TestDirectories{
		{LocalRepository{"/home/chipskein/Music/", "/home/chipskein/sources/mocg/audios/", nil, true}, false},
		{LocalRepository{"/home/chipskein/Music/", "/home/chipskein/fasojfklasjfklasjhfjkahsf", nil, true}, false},
		{LocalRepository{"/home/chipskein/Music/", "/", nil, true}, false},
		{LocalRepository{"/home/chipskein/Music/", "/home/chipskein/", nil, true}, false},
		{LocalRepository{"/home/chipskein/Music/", "../../audios", nil, true}, false},
		{LocalRepository{"/home/chipskein/Music/", "", nil, true}, false},
		{LocalRepository{"/wrong/default/path", "fhjasjfhajkshfjkashfkjahfjkasf", nil, true}, true},
		{LocalRepository{"/another/wrong/default/path", "fjkasjfkasjfkjas", nil, true}, true},
		{LocalRepository{"../../audios/", "/invalid/directory/convert", nil, true}, false},
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
func TestMapFiles(t *testing.T) {
	var tests = []TestDirectories{
		{LocalRepository{"/home/chipskein/Music", "/home/chipskein/sources/mocg/audios/", nil, true}, false},
		{LocalRepository{"/home/chipskein/Music", "../", nil, true}, false},
		{LocalRepository{"/home/chipskein/Music", "/wrong/path", nil, true}, false},
	}
	for _, test := range tests {
		err := test.repository.MapFiles()
		if err != nil {
			t.Errorf("\n[ERRROR]Could not read Directory %s\n", test.repository.CURRENT_DIRECTORY)
			continue
		}
		for _, file := range test.repository.Files {
			t.Logf("\nFilename:%s\nExtension:%s\nIsADirectory:%t\nFullPath:%s\nSize:%d\n", file.Name, file.Extension, file.IsADirectory, file.FullPath, file.Size)
		}
	}
}
func TestReadFIle(t *testing.T) {

	f := ReadFile("../testAudios/ogg/music1.ogg")
	defer f.Close()
	b1 := make([]byte, 50)
	n1, _ := f.Read(b1)
	t.Logf("%d bytes: %s\n", n1, string(b1[:n1]))
	o2, _ := f.Seek(51, 0)
	b2 := make([]byte, 2)
	n2, _ := f.Read(b2)
	t.Logf("%d bytes @ %d: ", n2, o2)
	t.Logf("%v\n", string(b2[:n2]))

	f.Close()
}
