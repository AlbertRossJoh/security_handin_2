package shared

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	NodesFilename  = "nodes"
	LeaderFilename = "leader"
)

func WriteToSharedFile(dockerId string, filePath string) {

	f, err := os.OpenFile(filePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)

	if err != nil {
		log.Println("[WriteToSharedFile] Couldn't open the shared file! On server ", dockerId)
	}
	defer f.Close()
	if _, err := f.WriteString(dockerId + "\n"); err != nil {
		log.Println("Couldn't write to shared file! On server ", dockerId)
	}
}

func ReplaceSharedFileContents(fileContents []string, filename string) {
	sharedFilePath := GetPath(filename)

	f, err := os.OpenFile(sharedFilePath,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644)

	if err != nil {
		log.Println("[ReplaceSharedFileContents] Couldn't open the shared file!")
	}
	defer f.Close()
	if err := f.Truncate(0); err != nil {
		log.Println("[ReplaceSharedFileContents] Couldn't write to shared file!")
	}

	f.Seek(0, 0)

	for i := 0; i < len(fileContents); i++ {
		WriteToSharedFile(fileContents[i], filename)
	}
}

func GetFileContents(dockerId string, filePath string) []string {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println("File does not exist: " + filePath)
	}

	f, err := os.OpenFile(filePath,
		os.O_RDONLY,
		0644)

	if err != nil {
		log.Println("[GetFileContents] Couldn't open the shared file! On server ", dockerId)
	}
	defer f.Close()
	bytes := make([]byte, 1_000_000)

	f.Read(bytes)
	contents := string(bytes)
	strs := strings.Split(contents, "\n")

	if len(strs) > 0 {
		// Removes whitespace from end of file
		strs = strs[:len(strs)-1]
	}
	return strs
}

func RemoveDockerIdFromFile(dockerId string, filename string) error {
	fileContents := GetFileContents("Unknown DockerId", filename)

	removeIndex, err := indexOf(fileContents, dockerId)

	if err != nil {
		return errors.New("DockerId not found in file")
	}

	newFileContents, _ := removeAtIndex(fileContents, removeIndex)

	ReplaceSharedFileContents(newFileContents, filename)

	return nil
}

// Finds the first instance of `str` in `arr`.
func indexOf(arr []string, str string) (int, error) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return i, nil
		}
	}

	return -1, errors.New("string not found in array")
}

func removeAtIndex(arr []string, index int) ([]string, error) {

	// add error checks...
	newArr := make([]string, len(arr)-1)

	for i, j := 0, 0; i < len(newArr); {
		if i == index {
			j++
		}

		newArr[i] = arr[j]

		i++
		j++
	}

	return newArr, nil
}

func GetPath(file string) string {
	return filepath.FromSlash("/var/hosts/" + file)
}
