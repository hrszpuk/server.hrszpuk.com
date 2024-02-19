package main

import (
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ValidateDirectory(dirname string) error {
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		log.Printf("'%s' does not exist! Generating...\n", dirname)
		err := os.Mkdir(dirname, os.ModeDir)
		return err
	}
	log.Printf("found '%s'.", dirname)
	return nil
}

func CountFiles(dirname string) (int, error) {
	dir, err := os.Open(dirname)
	defer dir.Close()

	if err != nil {
		return 0, nil
	}

	names, err := dir.Readdirnames(-1)
	return len(names), err
}

func DirectoryStatusCheck(dirname string) {
	ValidateDirectory(dirname)
	count, err := CountFiles(dirname)
	log.Printf("Found '%d' files in '%s' directory.\n", count, dirname)
	check(err)
}
