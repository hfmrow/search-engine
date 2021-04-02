// gohOptions.go

/*
	Source file auto-generated on Fri, 02 Apr 2021 13:10:55 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2018-21 H.F.M - Search Engine v1.9 github.com/hfmrow/search-engine
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	glfsff "github.com/hfmrow/genLib/files/findFiles"
	gltsbh "github.com/hfmrow/genLib/tools/bench"
	gltses "github.com/hfmrow/genLib/tools/errors"

	gidg "github.com/hfmrow/gtk3Import/dialog"
	gidgcr "github.com/hfmrow/gtk3Import/dialog/chooser"
	gimc "github.com/hfmrow/gtk3Import/misc"
	gitw "github.com/hfmrow/gtk3Import/treeview"
)

// App infos
var Name = "Search Engine"
var Vers = "v1.9"
var Descr = "This program is designed to search files over directory,\nsubdirectory, and retrieving information based on\ndate, type, patterns contained in name."
var Creat = "H.F.M"
var YearCreat = "2018-21"
var LicenseShort = "This program comes with absolutely no warranty.\nSee the The MIT License (MIT) for details:\nhttps://opensource.org/licenses/mit-license.php"
var LicenseAbrv = "License (MIT)"
var Repository = "github.com/hfmrow/search-engine"

// Vars declarations
var (
	maintitle,
	tempDir,
	absoluteRealPath,
	optFilename string

	mainOptions *MainOpt
	devMode     bool

	doTempDir      bool
	errNoSelection = errors.New("There is no selection currently.")

	columnsNames = [][]string{
		{"Name", "text"},
		{"Size", "text"},
		{"Type", "text"},
		{"Time", "text"},
		{"Path", "text"},
		{"sizeSort", "int64"}, // This one will be invisible (int64)
		{"dateSort", "int64"}} // This one will be invisible (int64)

	columnsMap = map[string]int{
		`name`:     0,
		`size`:     1,
		`type`:     2,
		`time`:     3,
		`path`:     4,
		`sizeSort`: 5,
		`dateSort`: 6,
	}

	filesScanned int

	// Functions mapping
	timer           = gltsbh.BenchNew(false)
	statusbar       = gimc.StatusBar{}
	clipboard       = gimc.Clipboard{}
	titlebar        *gimc.TitleBar
	tvs             *gitw.TreeViewStructure
	Check           = gltses.Check
	DialogMessage   = gidg.DialogMessage
	storeFoundFiles []glfsff.StoreFiles

	// Popup
	popupMenu              *gimc.PopupMenuIconStruct
	PopupMenuIconStructNew = gimc.PopupMenuIconStructNew

	// To store original label content for newer than and older than buttons.
	origLabelNT,
	origLabelOT string
)

type searchList struct {
	And []string
	Or  []string
	Not []string
}

type MainOpt struct {
	/* Public, will be saved and restored */
	About                       *gidg.AboutInfos
	MainWinWidth, MainWinHeight int
	LanguageFilename            string
	LastDirectory               string
	CaseSensitive               bool
	CharClass                   bool
	CharClasStrict              bool
	WildCard                    bool
	Regex                       bool
	FollowSymlinks              bool
	FileType                    int
	DateType                    int
	DateZone                    int
	Depth                       int
	WordAnd                     bool
	WordOr                      bool
	WordNot                     bool
	SplitAnd                    bool
	SplitOr                     bool
	SplitNot                    bool
	UpdateOnChanges             bool
	SearchList                  searchList
	AppLauncher                 string
	WebSearchEngine             string
	FileExplorer                string
	OptionsPath                 string

	/* Private, will NOT be saved */
	calendar         *gidgcr.Calendar
	calDataNewerThan *gidgcr.CalendarData
	calDataOlderThan *gidgcr.CalendarData

	displayFilesList [][]string
}

// Main options initialisation
func (opt *MainOpt) Init() {
	opt.About = new(gidg.AboutInfos)

	opt.LanguageFilename = "assets/lang/eng.lang"

	opt.WebSearchEngine = `https://www.google.com/search?q=`
	opt.FileExplorer = "thunar"
	opt.AppLauncher = "xdg-open"
	opt.Depth = -1

	opt.MainWinWidth = 800
	opt.MainWinHeight = 600
}

// Variables -> Objects.
func (opt *MainOpt) UpdateObjects() {
	mainObjects.MainWindow.Resize(opt.MainWinWidth, opt.MainWinHeight)

	mainObjects.SearchFilechooserbutton.SetFilename(opt.LastDirectory)
	mainObjects.SearchCheckbuttonCaseSensitive.SetActive(opt.CaseSensitive)
	mainObjects.SearchCheckbuttonWildCard.SetActive(opt.WildCard)
	mainObjects.SearchCheckbuttonRegex.SetActive(opt.Regex)
	mainObjects.SearchCheckbuttonCharClasses.SetActive(opt.CharClass)
	mainObjects.SearchCheckbuttonCharClassesStrictMode.SetActive(opt.CharClasStrict)

	mainObjects.SearchSpinbuttonDepth.SetValue(float64(opt.Depth))
	mainObjects.SearchCheckbuttonWordAnd.SetActive(opt.WordAnd)
	mainObjects.SearchCheckbuttonWordOr.SetActive(opt.WordOr)
	mainObjects.SearchCheckbuttonWordNot.SetActive(opt.WordNot)
	mainObjects.SearchCheckbuttonSplitedAnd.SetActive(opt.SplitAnd)
	mainObjects.SearchCheckbuttonSplitedOr.SetActive(opt.SplitOr)
	mainObjects.SearchCheckbuttonSplitedNot.SetActive(opt.SplitNot)

	mainObjects.SearchComboboxTextType.SetActive(opt.FileType)
	mainObjects.SearchComboboxTextDateType.SetActive(opt.DateType)
	mainObjects.SearchComboboxTextDateZone.SetActive(opt.DateZone)

	mainObjects.SearchCheckbuttonFollowSL.SetActive(opt.FollowSymlinks)
}

// Objects -> Variables.
func (opt *MainOpt) UpdateOptions() {

	opt.MainWinWidth, opt.MainWinHeight = mainObjects.MainWindow.GetSize()

	opt.LastDirectory = mainObjects.SearchFilechooserbutton.GetFilename()
	opt.CaseSensitive = mainObjects.SearchCheckbuttonCaseSensitive.GetActive()
	opt.WildCard = mainObjects.SearchCheckbuttonWildCard.GetActive()
	opt.Regex = mainObjects.SearchCheckbuttonRegex.GetActive()
	opt.CharClass = mainObjects.SearchCheckbuttonCharClasses.GetActive()
	opt.CharClasStrict = mainObjects.SearchCheckbuttonCharClassesStrictMode.GetActive()

	opt.Depth = int(mainObjects.SearchSpinbuttonDepth.GetValue())
	opt.WordAnd = mainObjects.SearchCheckbuttonWordAnd.GetActive()
	opt.WordOr = mainObjects.SearchCheckbuttonWordOr.GetActive()
	opt.WordNot = mainObjects.SearchCheckbuttonWordNot.GetActive()
	opt.SplitAnd = mainObjects.SearchCheckbuttonSplitedAnd.GetActive()
	opt.SplitOr = mainObjects.SearchCheckbuttonSplitedOr.GetActive()
	opt.SplitNot = mainObjects.SearchCheckbuttonSplitedNot.GetActive()

	opt.FileType = mainObjects.SearchComboboxTextType.GetActive()
	opt.DateType = mainObjects.SearchComboboxTextDateType.GetActive()
	opt.DateZone = mainObjects.SearchComboboxTextDateZone.GetActive()

	opt.FollowSymlinks = mainObjects.SearchCheckbuttonFollowSL.GetActive()
}

// Read Options from file
func (opt *MainOpt) Read() (err error) {
	var textFileBytes []byte
	opt.Init()
	if textFileBytes, err = ioutil.ReadFile(optFilename); err == nil {
		err = json.Unmarshal(textFileBytes, &opt)
	}
	if err != nil {
		fmt.Printf("Error while reading options file: %s\n", err.Error())
	}
	return
}

// Write Options to file
func (opt *MainOpt) Write() (err error) {
	var jsonData []byte
	var out bytes.Buffer
	opt.UpdateOptions()
	opt.About.DlgBoxStruct = nil // remove dialog object before saving
	if jsonData, err = json.Marshal(&opt); err == nil {
		if err = json.Indent(&out, jsonData, "", "\t"); err == nil {
			err = ioutil.WriteFile(optFilename, out.Bytes(), os.ModePerm)
		}
	}
	return err
}
