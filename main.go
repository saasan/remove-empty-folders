package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/saasan/go-s2dir"
	"github.com/saasan/go-term"
)

func removeEmptyDir(dirname string) error {
	dirs, _, err := s2dir.Read(dirname)
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		path := filepath.Join(dirname, dir.Name())
		if err := removeEmptyDir(path); err != nil {
			return err
		}
	}

	isEmpty, err := s2dir.IsEmpty(dirname)
	if err != nil {
		return err
	}

	if !isEmpty {
		return nil
	}

	fmt.Printf("空フォルダを削除: %s\n", dirname)
	if err := os.Remove(dirname); err != nil {
		return fmt.Errorf("%s を削除できません。: %w", dirname, err)
	}

	return nil
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("処理対象: %s\n", arg)
		if err := removeEmptyDir(arg); err != nil {
			fmt.Println(err)
		}
	}

	if runtime.GOOS == "windows" {
		fmt.Println("Press any key to continue...")
		term.WaitAnyKey()
	}
}
