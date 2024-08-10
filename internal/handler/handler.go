package handler

import (
	"fmt"
	utils "kuroneko/internal/ui/utils"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
)

func GenerateCbzFiles(folderPath string, w fyne.Window) {
	allowedFileExtensions := []string{".jpeg", ".png", ".jpg"}

	if folderPath == "" {
		utils.ThrowError(w, "Invalid folder path!")
		return
	}

	fmt.Println("Running function on folder:", folderPath)

	var mangaImagesFilePaths []string

	entries, err := os.ReadDir(folderPath)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dirPath := filepath.Join(folderPath, entry.Name())
			subEntries, err := os.ReadDir(dirPath)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			for _, subEntry := range subEntries {
				if !subEntry.IsDir() {
					subEntryPath := filepath.Join(dirPath, subEntry.Name())
					if utils.Includes(allowedFileExtensions, filepath.Ext(subEntry.Name())) {
						mangaImagesFilePaths = append(mangaImagesFilePaths, subEntryPath)
					}
				}
			}
		}
	}

	if len(mangaImagesFilePaths) == 0 {
		utils.ThrowError(w, "Not found manga images!")
		return
	}

	fmt.Println("Manga images found:", mangaImagesFilePaths)
}
