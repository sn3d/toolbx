package testutil

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateTestData(src string) (string, error) {

	// create temporary testutil dir
	dest, err := os.MkdirTemp("", "toolbx-")
	if err != nil {
		return "", err
	}

	// copy content
	err = copyDir(src, dest)
	return dest, err
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
