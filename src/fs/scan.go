package fs

import (
	"os"

	"github.com/schollz/progressbar/v3"
)

var PROGRESS_BAR *progressbar.ProgressBar

func Scan(origin string) (int64, error) {
	entries, err := os.ReadDir(origin)
	if err != nil {
		return 0, err
	}

	var COUNTER int64 = 1

	if len(entries) == 0 {
		return COUNTER, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			c, _ := Scan(origin + "/" + entry.Name())
			COUNTER += c
		} else {
			COUNTER++
		}
	}

	return COUNTER, nil
}

// folder 1111
// file 3298
// total 4409
