package rename

import (
	// "fmt"
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
)

func rename(path string, file fs.DirEntry, source regexp.Regexp, dest string, run bool) error {

	if file.IsDir() {
		return nil
	}

	filename := file.Name()

	if !source.MatchString(filename) {
		return nil
	}

	newName := source.ReplaceAllString(filename, dest)

	fmt.Printf("%s => %s\n", filename, newName)
	if run {
		return os.Rename(fmt.Sprintf("%s/%s", path, filename), fmt.Sprintf("%s/%s", path, string(newName)))
	}
	return nil
}

func BatchRename(path string, reg string, outReg string, run bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}
	inReg := regexp.MustCompile(reg)
	log.Println("start")
	for _, file := range files {
		err := rename(path, file, *inReg, outReg, run)
		if err != nil {
			log.Printf("Rename %s fail. %s", file.Name(), err.Error())
		}
	}
	log.Println("Finish")
}
