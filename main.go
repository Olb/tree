package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	if printFiles {
		return printTree(out, path, printFiles, "")
	}
	return printTreeWithoutFiles(out, path, printFiles, "")
}

func printTree(out io.Writer, path string, printFiles bool, prefix string) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return err
	}
	for idx, file := range files {
		if !file.IsDir() {
			var size string
			if file.Size() == 0 {
				size = " (empty)"
			} else {
				size = fmt.Sprintf(" (%db)", file.Size())
			}
			toPrint := prefix + "├───" + file.Name() + size
			if idx == len(files)-1 {
				toPrint = prefix + "└───" + file.Name() + size
			}
			fmt.Fprint(out, toPrint+"\n")
		} else {
			toPrint := prefix + "├───" + file.Name()
			if idx == len(files)-1 {
				toPrint = prefix + "└───" + file.Name()
			}
			fmt.Fprint(out, toPrint+"\n")
		}

		if file.IsDir() {
			nextPref := "\t"
			if idx != len(files)-1 {
				nextPref = "│\t"
			}
			printTree(out, path+"/"+file.Name(), printFiles, prefix+nextPref)
		}
	}
	return nil
}

func printTreeWithoutFiles(out io.Writer, path string, printFiles bool, prefix string) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return err
	}
	var dirs []os.FileInfo
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f)
		}
	}
	for idx, file := range dirs {

		if file.IsDir() {
			toPrint := prefix + "├───" + file.Name()
			if idx == len(dirs)-1 {
				toPrint = prefix + "└───" + file.Name()
			}
			fmt.Fprint(out, toPrint+"\n")
			nextPref := "\t"
			if idx != len(dirs)-1 {
				nextPref = "│\t"
			}
			printTreeWithoutFiles(out, path+"/"+file.Name(), printFiles, prefix+nextPref)
		}
	}
	return nil
}
