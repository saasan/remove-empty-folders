package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/saasan/go-s2file"
)

func TestRemoveEmptyDir(t *testing.T) {
	// 一時ディレクトリを作成
	tmpDir, err := os.MkdirTemp("", "TestRemoveEmptyDir")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	var (
		dir1_1  = filepath.Join(tmpDir, "dir1_1")
		dir1_3  = filepath.Join(tmpDir, "dir1_1", "dir1_2", "dir1_3")
		dir2_1  = filepath.Join(tmpDir, "dir2_1")
		dir2_2  = filepath.Join(tmpDir, "dir2_1", "dir2_2")
		dir2_3  = filepath.Join(tmpDir, "dir2_1", "dir2_2", "dir2_3")
		dir3_1  = filepath.Join(tmpDir, "dir3_1")
		dir3_3  = filepath.Join(tmpDir, "dir3_1", "dir3_2", "dir3_3")
		file2_2 = filepath.Join(tmpDir, "dir2_1", "file2_2.txt")
		file3_3 = filepath.Join(tmpDir, "dir3_1", "dir3_2", "file3_3.txt")
	)

	// サブディレクトリを作成
	os.MkdirAll(dir1_3, os.ModePerm)
	os.MkdirAll(dir2_3, os.ModePerm)
	os.MkdirAll(dir3_3, os.ModePerm)
	// ファイルを作成
	os.WriteFile(file2_2, []byte("file2"), os.ModePerm)
	os.WriteFile(file3_3, []byte("file3"), os.ModePerm)

	removeEmptyDir(dir1_1)
	removeEmptyDir(dir2_1)
	removeEmptyDir(dir3_1)

	if s2file.Exists(dir1_1) {
		t.Error("dir1_1失敗")
	}

	if !s2file.Exists(file2_2) {
		t.Error("file2_2失敗")
	}
	if s2file.Exists(dir2_2) {
		t.Error("dir2_2失敗")
	}

	if !s2file.Exists(file3_3) {
		t.Error("file3_3失敗")
	}
	if s2file.Exists(dir3_3) {
		t.Error("dir3_3失敗")
	}
}
