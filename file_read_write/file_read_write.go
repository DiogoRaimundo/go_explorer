/*
	File permissions are flags that are handled by a user on the system. This is represented by the permission indicator.

	On Linux, we can use "ls -alh" to check files' permission indicator:
		drwxr-xr-x   3 Uday.Hiwarale  admin  96B Mar 10 18:55    .
		drwxr-xr-x  55 Uday.Hiwarale  admin  1.7K Mar 10 18:55   ..
		-rwxr-xr--   1 Uday.Hiwarale  admin  0B Mar 10 18:55     go-fs.go

	"drwxr-xr-x" is in fact a bit plus 3 groups of 3bits (collectively called the file mode): d rwx r-x r-x
	The first bit indicates either the location on the file system is a directory (d) or a file (-).
	Each group of 3 bits represents "read-write-execute" permissions.
	The first group represents the file permission for the file owner which is indicated by the username in the third column of the result (which is "Uday.Hiwarale").
	The second group represents the file permission for the user group that owns the file.
	The third group represents the file permission for the rest of the system users.

	Typically, a file mode value is represented as an octal number:
		Octal Number	Permission					rwx Flags	Binary
		7				read, write and execute		rwx			111
		6				read and write				rw-			110
		5				read and execute			r-x			101
		4				read only					r--			100
		3				write and execute			-wx			011
		2				write only					-w-			010
		1				execute only				--x			001
		0				none						---			000
*/

package fileReadWrite

import (
	"fmt"
	"os"
	"time"
)

func Run() {
	printCurrentDirectory()

	createOrRewriteFile()

	printFileContent()

	appendTimeToFile()

	printFileContent()
}

const fileName string = "localfile.txt"

func printCurrentDirectory() {
	path, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	fmt.Println("Current directory:", path)
	fmt.Println()
}

func createOrRewriteFile() {
	currentDate := time.Now().UnixNano()

	dataToWrite := []byte(fmt.Sprintf("This file was created/rewritten (%d)", currentDate))
	err := os.WriteFile(fileName, dataToWrite, 0777)

	if err != nil {
		panic(err)
	}
}

func printFileContent() {
	data, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println("File content:", string(data))
	fmt.Println()
}

func appendTimeToFile() {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	currentDate := time.Now().UnixNano()

	_, err = f.WriteString(fmt.Sprintf("\nThis line wasn't there originally (%d)", currentDate))
	if err != nil {
		panic(err)
	}
}
