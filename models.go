package main

import (
	"github.com/Mimoja/MFT-Common"
	"path/filepath"
	"time"
)

type Page struct {
	Title  string
	Error  string
	IsOkay bool
}

type MainPage struct {
	Page
	Imports int
	FlashImages int
}

type ReportOverviewPage struct {
	Page
	LastReports []ImportRef
}

type ImportRef struct {
	ImportTime string
	ID         string
	Name       string
}

type ReportPage struct {
	Page
	Config      *MFTCommon.AppRunConfiguration
	Data struct {
		Import      MFTCommon.ImportEntry
		FlashImages []FlashDocument
	}
}

type FlashDocument struct {
	FlashImage MFTCommon.FlashImage
	Certificates []CertificateDocument
	EFI map[string]interface{}
}

type CertificateDocument struct {
	Valid      bool
	Subject    []string
	Issuer     []string
	Serial     string
	Raw        map[string]interface{}
	Algorithm  string
	SelfSigned bool
}

/*
type ImportEntry struct {
	ImportDataDefinition string         `json:",omitempty"`
	MetaData             DownloadEntry  `json:",omitempty"`
	Contents             []StorageEntry `json:",omitempty"`
	ImportTime           string         `json:",omitempty"`
	Success              bool           `json:",omitempty"`
}
*/

func NewImportRef(importEntry MFTCommon.ImportEntry) ImportRef {

	name := filepath.Base(importEntry.MetaData.DownloadPath)

	if name == "." {
		name = ""
	}

	importTime, _ := time.Parse(time.RFC3339, importEntry.ImportTime)
	return ImportRef{
		ImportTime: importTime.Format("02.01.2006 15:04:05"),
		ID:         importEntry.MetaData.PackageID.GetID(),
		Name:       name,
	}
}

