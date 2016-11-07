/*
 * Author: Trey Watford
 * Email: treyjustinwatford@gmail.com
 * Description: This is a tool designed to find all .git/config files on a
 * file system and replace given text with a new string.  This tool is meant to be
 * used from the command line and will work on both Windows and Unix based systems.
 * Feel free to use this tool as needed and please let me know of any bugs so that
 * they can be fixed.
 *
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//define flags with default values
var root = flag.String("root", getRoot(), "starting directory")
var replace = flag.String("replace", "", "text to be replaced")
var newText = flag.String("newText", "", "text to replace with")
var force = flag.Bool("force", false, "true forces file modification without prompt")

//main driver function
func main() {
	text := ""   //dummy string to hold output open
	flag.Parse() //parse flags
	if flag.NFlag() <= 3 {
		err := filepath.Walk(*root, visit)
		if err != nil {
			fmt.Printf("Error, returned %v\n", err)
		} else {
			fmt.Print("\nAll remotes updated successfully.\n")
			fmt.Print("Press any key then [ENTER] to quit.\n")
			fmt.Scanf("%s", &text) //used to hold output screen open
		}
	}
}

//function to return the root directory of a file system depending on OS
func getRoot() string {
	if runtime.GOOS == "windows" {
		return "C:\\"
	}
	return "/"
}

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	desiredFile := filepath.FromSlash(".git/config") //returns os-specific path
	if strings.Contains(path, desiredFile) {
		err := modifyRemotes(path, *replace, *newText)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

//modify remote repo changes occurrences of one string to another in a given file
func modifyRemotes(path string, match string, newStr string) error {
	//create input file
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	//alert user of file read
	fmt.Printf("\n\nread file %s.", path)
	lines := strings.Split(string(input), "\n") //split file into lines
	indices := []int{}                          //index slice for matched lines

	//count lines containing replace and copy indices into index slice
	for i, line := range lines {
		if strings.Contains(line, match) {
			indices = append(indices, i)
		}
	}
	if len(indices) > 0 {
		//alert user that a number of lines in file will be modified
		fmt.Printf("\n%v lines will be modified.", len(indices))
		if *force != true {
			if getUserInput() {
				for v := range indices {
					fmt.Printf("modifying line %s --->", lines[indices[v]])
					lines[indices[v]] = strings.Replace(lines[indices[v]], match, newStr, 1)
					fmt.Printf("%s\n", lines[indices[v]])
				}
			} else {
				fmt.Printf("\nNothing to modify, closing file...")
			}
		} else {
			for v := range indices {
				fmt.Printf("modifying line %s --->", lines[indices[v]])
				lines[indices[v]] = strings.Replace(lines[indices[v]], match, newStr, 1)
				fmt.Printf("%s\n", lines[indices[v]])
			}
		}
	}

	//overwrite file
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("\nFile saved --success!\n\n")
	return nil
}

//function returns true if user inputs y and false for n
func getUserInput() bool {
	input := ""
	loop := true
	for loop == true {
		fmt.Print("Continue [y/n]? : ")
		fmt.Scanf("%s", &input)
		input = strings.ToLower(input)
		if strings.Compare(input, "y") == 0 {
			loop = false
		} else if strings.Compare(input, "n") == 0 {
			fmt.Printf("\nNo changes will be made.")
			return false
		} else {
			fmt.Printf("\nPlease enter y or n\n")
		}
	}
	return true
}
