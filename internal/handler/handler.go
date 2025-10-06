package handler

import (
	"archive/zip"
	"fmt"
	"io"
	utils "kuroneko/internal/ui/utils"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
)

func GenerateCbzFiles(folderPath string, w fyne.Window, progressCallback func(float64)) {
	allowedFileExtensions := []string{".jpeg", ".png", ".jpg"}

	if folderPath == "" {
		utils.ThrowError(w, "Invalid folder path!")
		return
	}

	fmt.Println("Running function on folder:", folderPath)

	mangaImagesByDir := make(map[string][]string)

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
						mangaImagesByDir[dirPath] = append(mangaImagesByDir[dirPath], subEntryPath)
					}
				}
			}
		}
	}

	if len(mangaImagesByDir) == 0 {
		utils.ThrowError(w, "Not found manga images!")
		return
	}

	fmt.Println("Manga images found:", mangaImagesByDir)

	successCount := 0
	errorCount := 0
	totalDirs := len(mangaImagesByDir)
	processedDirs := 0

	for dir, images := range mangaImagesByDir {
		if len(images) > 0 {
			cbzPath := filepath.Join(folderPath, filepath.Base(dir)+".cbz")
			err := createZip(cbzPath, images)
			if err != nil {
				utils.ThrowError(w, fmt.Sprintf("Error creating CBZ for %s: %v", dir, err))
				errorCount++
			} else {
				fmt.Printf("Created CBZ: %s\n", cbzPath)
				successCount++
			}
		}
		processedDirs++
		if progressCallback != nil {
			progressCallback(float64(processedDirs) / float64(totalDirs))
		}
	}

	if successCount > 0 {
		utils.ShowSuccess(w, fmt.Sprintf("Successfully created %d CBZ file(s)!", successCount))
	}
	if errorCount > 0 {
		utils.ThrowError(w, fmt.Sprintf("Failed to create %d CBZ file(s). Check console for details.", errorCount))
	}
}

func createZip(zipPath string, files []string) error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		err := addFileToZip(zipWriter, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(filePath)
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}
