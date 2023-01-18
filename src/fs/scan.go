package fs

import (
	"os"

	"github.com/schollz/progressbar/v3"
)

var PROGRESS_BAR *progressbar.ProgressBar

func Scan(origin string) (int64, error) {
	entries, err := os.ReadDir(origin)
	if err != nil || len(entries) == 0 {
		return 0, err
	}

	var count int64 = 0
	for _, entry := range entries {
		if entry.IsDir() {
			c, err := Scan(origin + "/" + entry.Name())
			if err != nil {
				return 0, err
			}
			count += c
		} else {
			count++
		}
	}

	return count, nil
}

/*

func folder(origin string, target string) error {
	entries, err := os.ReadDir(origin)
	if err != nil {
		return err
	}

	if DEBUG {
		fmt.Println("COPY FOLDER:", origin, "->", target)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			err := os.Mkdir(target+"/"+entry.Name(), PERMS_DIRECTORY)
			if err != nil {
				return err
			}

			folder(origin+"/"+entry.Name(), target+"/"+entry.Name())
		} else {
			err := file(origin+"/"+entry.Name(), target+"/"+entry.Name())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
*/
