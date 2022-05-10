package archive

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type ArchiveInstaller struct {
	BearerToken string
}

func Installer(bearerToken string) *ArchiveInstaller {
	return &ArchiveInstaller{
		BearerToken: bearerToken,
	}
}

// Install is downloading artifact archive and unzip it into destination
// folder
func (i *ArchiveInstaller) Install(uri url.URL, installationDir string) error {

	uri.Scheme = "https"

	// create request
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return err
	}

	if i.BearerToken != "" {
		req.Header.Add("PRIVATE-TOKEN", i.BearerToken)
		req.Header.Add("Bearer", i.BearerToken)
	}

	// Get the data
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// write body to 'archive.zip'
	archivePath := filepath.Join(installationDir, "archive.zip")
	out, err := os.OpenFile(archivePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// unzip the archive
	err = extractZip(archivePath, installationDir)
	if err != nil {
		return err
	}

	return nil
}

func extractZip(src string, dest string) error {
	archive, err := zip.OpenReader(src)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	for _, f := range archive.File {
		filePath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	return nil
}
