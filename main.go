// main.go

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
	"fmt"
	"log"

	gidgcr "github.com/hfmrow/gtk3Import/dialog/chooser"
	gimc "github.com/hfmrow/gtk3Import/misc"
)

func main() {

	/* Build options */
	// devMode: is used in some functions to control the behavior of the program
	// When software is ready to be published, this flag must be set at "false"
	// that means:
	// - options file will be stored in $HOME/.config/[Creat]/[softwareName],
	// - translate function if used, will no more auto-update "sts" map sentences,
	// - all built-in assets will be used instead of the files themselves.
	//   Be aware to update assets via "Goh" and translations with "Got" before all.
	devMode = true
	absoluteRealPath, optFilename = getAbsRealPath()

	// Initialization of assets according to the chosen mode (devMode).
	// you can set this flag to your liking without reference to devMode.
	assetsDeclarationsUseEmbedded(!devMode)

	// Create temp directory .. or not
	doTempDir = false

	/* Init & read options file */
	mainOptions = new(MainOpt) // Assignate options' structure.
	mainOptions.Read()         // Read values from options' file if exists.

	maintitle = fmt.Sprintf("%s %s  %s %s %s",
		Name,
		Vers,
		YearCreat,
		Creat,
		LicenseAbrv)

	/* Init gtk display */
	mainStartGtk(maintitle,
		mainOptions.MainWinWidth,
		mainOptions.MainWinHeight, true)
}

/*******************************************\
/* Executed before signals initialisation. */
/******************************************/
func mainApplication() {
	/* Init AboutBox */
	mainOptions.About.InitFillInfos(
		mainObjects.MainWindow,
		"About "+Name,
		Name,
		Vers,
		Creat,
		YearCreat,
		LicenseAbrv,
		LicenseShort,
		Repository,
		Descr,
		"",
		tickIcon48)

	/* Translate init */
	translate = MainTranslateNew(absoluteRealPath+mainOptions.LanguageFilename, devMode)

	/* Get cmdline  if exist */
	cmdLineParser()

	/* Init Clipboard */
	Check(clipboard.Init())

	/* Init Statusbar */
	statusbar.Init(mainObjects.Statusbar, []string{"Status:", "Found:", "Selected: ", "Base directory:", "Search time:", "Display time:"})

	/* Init titlebar */
	titlebar = gimc.TitleBarStructureNew(mainObjects.MainWindow)

	/* Init Comboboxes */
	fillAllComboboxText()

	/* Init SpinButtons */
	initSpinButtons()

	/* Init Popup menu */
	initPopup()

	/* Init Treeview	*/
	initTreeview()

	/* Storing originals labels for newer than and older than buttons */
	origLabelNT, _ = mainObjects.SearchButtonNewerThan.GetLabel()
	origLabelOT, _ = mainObjects.SearchButtonOlderThan.GetLabel()

	/* Init calendar window */
	mainOptions.calendar = gidgcr.CalendarNew(mainObjects.MainWindow, nil, "title", calendarPers48, []string{"Reset", "Ok"}, reset48, tickIcon48)
	mainOptions.calendar.DisplayTime = true
	mainOptions.calendar.ButtonsSize = 18
	mainOptions.calendar.ButtonOkPosition = 1
	initCalData()
}

/******************************************\
/* Executed after signals initialisation. */
/*****************************************/
func afterSignals() {
	/* Set options */
	mainOptions.UpdateObjects()
	mainObjects.SearchCheckbuttonCharClassesStrictMode.SetVisible(false)
	mainOptions.UpdateOnChanges = false
	SearchCheckbuttonRegexToggled()
	SearchCheckbuttonCharClassesClicked()
	if mainObjects.SearchCheckbuttonRegex.GetActive() {
		mainObjects.SearchCheckbuttonWildCard.SetSensitive(false)
	}
}

/*************************************\
/* Executed just before closing all. */
/************************************/
func onShutdown() bool {
	var err error
	// Update mainOptions with GtkObjects and save it
	if err = mainOptions.Write(); err == nil {
		// What you want to execute before closing the app.
		// Return:
		// true for exit applicaton
		// false does not exit application
	}
	if err != nil {
		log.Fatalf("Unexpected error on exit: %s", err.Error())
	}
	return true
}
