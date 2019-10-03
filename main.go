// main.go

// Source file auto-generated on Wed, 02 Oct 2019 17:33:10 using Gotk3ObjHandler v1.3.8 ©2018-19 H.F.M

/*
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"fmt"

	gimc "github.com/hfmrow/gtk3Import/misc"
)

func main() {
	var err error

	/* Be or not to be ... in dev mode ... */
	devMode = true

	/* Build directory for tempDir */
	doTempDir = false

	/* Naming widgets as Gtk objects names to use in css.
	   Set to false if they already named in Glade.*/
	namingWidget = true

	/* Set to true when you choose using embedded assets functionality */
	assetsDeclarationsUseEmbedded(!devMode)

	/* Init Options */
	mainOptions = new(MainOpt)
	mainOptions.Init()

	/* Read Options */
	err = mainOptions.Read()
	if err != nil {
		fmt.Printf("%s\n%v\n", "Options file not found or error on parsing.", err)
	}

	/* Init AboutBox */
	mainOptions.AboutOptions.InitFillInfos(
		"About "+Name,
		Name,
		Vers,
		Creat,
		YearCreat,
		LicenseAbrv,
		LicenseShort,
		Repository,
		Descr,
		searchEngine400x27,
		checked18x18)

	maintitle = fmt.Sprintf("%s %s  %s %s %s",
		mainOptions.AboutOptions.AppName,
		mainOptions.AboutOptions.AppVers,
		mainOptions.AboutOptions.YearCreat,
		mainOptions.AboutOptions.AppCreats,
		mainOptions.AboutOptions.LicenseAbrv)

	/* Init gtk display */
	mainStartGtk(maintitle,
		mainOptions.MainWinWidth,
		mainOptions.MainWinHeight, true)

}

func mainApplication() {
	/* Translate init */
	translate = MainTranslateNew(absoluteRealPath+mainOptions.LanguageFilename, devMode)

	/* Get cmdline  if exist */
	cmdLineParser()

	/* Set options */
	mainObjects.SearchCheckbuttonCharClassesStrictMode.SetVisible(false)
	mainOptions.UpdateOnChanges = false

	SearchCheckbuttonRegexToggled()

	/* Update gtk controls with stored values into mainOptions */
	mainOptions.UpdateObjects()

	/* Init Clipboard */
	Check(clipboard.Init())

	/* Init Statusbar */
	statusbar.StructureSetup(mainObjects.Statusbar, []string{"Status:", "Found:", "Selected: ", "Base directory:", "Search time:", "Display time:"})

	/* Init titlebar */
	titlebar = gimc.TitleBarNew(mainObjects.MainWindow, maintitle)

	/* Init Objects */
	fillAllComboboxText()
	initSpinButtons()
	// initComboboxes() // TODO Remove this function if all ok

	/* Storing originals labels for newer than and older than buttons */
	origLabelNT, _ = mainObjects.SearchButtonNewerThan.GetLabel()
	origLabelOT, _ = mainObjects.SearchButtonOlderThan.GetLabel()

	/* Init some  disp */
	initTreeview()                                                               /* Init Treeview	*/
	mainObjects.popUpMenu = initPopup(mainObjects.SearchTreeview, tvs.ListStore) /* Init Popup menu */

	/* Init display */
	SearchCheckbuttonCharClassesClicked()
}
