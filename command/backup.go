package command

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/nurali-techie/picopy/cli"
)

type backupCmd struct {
}

func NewBackupCommand() cli.Command {
	return new(backupCmd)
}

func (c *backupCmd) Execute(ctx context.Context, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing argument, check help")
	}
	fromDir := strings.TrimSuffix(args[0], "/")
	toDir := strings.TrimSuffix(args[1], "/")
	return backupFiles(fromDir, toDir)
}

func backupFiles(fromDir string, toDir string) error {
	fromFiles, err := mapFiles(fromDir)
	if err != nil {
		return err
	}
	fmt.Println("from.len=", len(fromFiles))

	toFiles, err := mapFiles(toDir)
	if err != nil {
		return err
	}
	fmt.Println("to.len=", len(toFiles))

	for fromFileKey, fromFilePath := range fromFiles {
		if _, ok := toFiles[fromFileKey]; !ok {
			fmt.Println("copy=", fromFilePath)
		}
	}

	return nil
}

func mapFiles(dir string) (map[string]string, error) {
	_, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fileMap := make(map[string]string)
	fileMap = visitDir(dir, fileMap)

	// fmt.Println("map.len=", len(fileMap))
	// for fileName, filePath := range fileMap {
	// 	fmt.Printf("%s = %s\n", fileName, filePath)
	// }

	return fileMap, nil
}

func visitDir(dir string, fileMap map[string]string) map[string]string {
	c, err := os.ReadDir(dir)
	if err != nil {
		return fileMap
	}

	for _, entry := range c {
		if entry.IsDir() {
			visitDir(fmt.Sprintf("%s/%s", dir, entry.Name()), fileMap)
		} else {
			fileInfo, err := entry.Info()
			// TODO handle err != nil
			if err == nil {
				// TODO revisit key strategy for uniqueness
				fileKey := fmt.Sprintf("%s_%d", fileInfo.Name(), fileInfo.Size())
				filePath := fmt.Sprintf("%s/%s", dir, entry.Name())
				fileMap[fileKey] = filePath
			}
		}
		// fmt.Printf("name=%s, dir=%t\n", entry.Name(), entry.IsDir())
	}

	return fileMap
}
