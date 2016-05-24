package walker

import (
	"io/ioutil"
	"os"
)

type WalkFn func(path string, file os.FileInfo) error

func Walk(path string, walkfn WalkFn) error {
	errc := make(chan error)
	go func(errc chan error) {
		file, err := os.Lstat(path)
		if file.IsDir() {
			dirList, err := ioutil.ReadDir(path)
			if err != nil {
				errc <- err
			}
			for _, content := range dirList {
				Walk(path+"/"+content.Name(), walkfn)
			}
		} else {
			err = walkfn(path, file)
			if err != nil {
				errc <- err
			}
		}
		defer close(errc)
	}(errc)

	for err := range errc {
		if err != nil {
			return err
		}
	}

	return nil
}
