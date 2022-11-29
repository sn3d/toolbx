package tempfs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type TempFs struct {
	rootDir string
}

func New(src string) (*TempFs, error) {
	// create temporary root dir
	rootDir, err := os.MkdirTemp("", "tempfs-")
	if err != nil {
		return nil, err
	}

	// copy content
	err = copyDir(src, rootDir)
	if err != nil {
		return nil, err
	}

	return &TempFs{rootDir}, nil
}

func (tfs *TempFs) GetRoot() string {
	return tfs.rootDir
}

// Get returns you absolute path for given path
func (tfs *TempFs) Get(path string) string {
	return filepath.Join(tfs.rootDir, path)
}

func copyDir(src string, dest string) error {

	if dest[:len(src)] == src {
		return fmt.Errorf("cannot copy a folder into the folder itself")
	}

	f, err := os.Open(src)
	if err != nil {
		return err
	}

	file, err := f.Stat()
	if err != nil {
		return err
	}
	if !file.IsDir() {
		return fmt.Errorf("Source " + file.Name() + " is not a directory!")
	}

	os.Mkdir(dest, 0755)

	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			err = copyDir(src+"/"+f.Name(), dest+"/"+f.Name())
			if err != nil {
				return err
			}
		}

		if !f.IsDir() {
			content, err := ioutil.ReadFile(src + "/" + f.Name())
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(dest+"/"+f.Name(), content, 0755)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
