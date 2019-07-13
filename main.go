// main.go

// Source file auto-generated on Wed, 10 Jul 2019 21:07:55 using Gotk3ObjHandler v1.3 ©2019 H.F.M

/*
	SearchEngine v1.8 ©2018 H.F.M

	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:

	Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
	associated documentation files (the "Software"), to dealin the Software without restriction,
	including without limitation the rights to use, copy, modify, merge, publish, distribute,
	sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all copies or
	substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
	NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
	NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
	DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT
	OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gotk3/gotk3/gtk"
	gi "github.com/hfmrow/searchEngine/gtk3Import"
)

func main() {
	devMode = true

	/* Set to true when you choose using embedded assets functionality  */
	assetsDeclarationsUseEmbedded(!devMode)

	/* Init Options */
	mainOptions = new(MainOpt)
	mainOptions.Init()

	/* Read Options */
	err = mainOptions.Read()
	if err != nil {
		fmt.Printf("%s\n%v\n", "Reading options error.", err)
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

	/* Init gtk display */
	mainStartGtk(fmt.Sprintf("%s %s  %s %s %s",
		mainOptions.AboutOptions.AppName,
		mainOptions.AboutOptions.AppVers,
		mainOptions.AboutOptions.YearCreat,
		mainOptions.AboutOptions.AppCreats,
		mainOptions.AboutOptions.LicenseAbrv),
		mainOptions.MainWinWidth,
		mainOptions.MainWinHeight, true)

}

func mainApplication() {
	/* Translate init. */
	translate = MainTranslateNew(absoluteRealPath+mainOptions.LanguageFilename, devMode)
	sts = translate.Sentences

	/* Get cmdline  if exist */
	cmdLineParser(os.Args)

	/* Set options */
	mainObjects.SearchCheckbuttonCharClassesStrictMode.SetVisible(false)
	mainOptions.UpdateOnChanges = false

	/* Update gtk controls with stored values into mainOptions */
	mainOptions.UpdateObjects()

	/*	Init Clipboard	*/
	Check(clipboard.Init())

	/*	Init Statusbar	*/
	statusbar.Init(mainObjects.Statusbar, []string{"Status:", "Found:", "Base directory:", "Search time:", "Display time:"})

	/*	Init Comboboxes	*/
	fillAllComboboxText()

	/* Storing originals labels for newer than and older than buttons */
	origLabelNT, _ = mainObjects.SearchButtonNewerThan.GetLabel()
	origLabelOT, _ = mainObjects.SearchButtonOlderThan.GetLabel()

	/* Init some  disp */
	initTreeview()                                                                           /* Init Treeview	*/
	mainObjects.popUpMenu = initPopup(mainObjects.SearchTreeview, mainObjects.fileListstore) /* Init Popup menu */

	/* Init display */
	SearchCheckbuttonCharClassesClicked()
}

func initTreeview() {

	/* Initialiste liststore Columns */
	mainObjects.fileListstore = gi.TreeViewListStoreSetup(mainObjects.SearchTreeview, true, columnsNames, false)

	// -----------************

	// getIterValue: rebuild full filepath from treeview entry (iter).
	var getIterValue = func(iter *gtk.TreeIter) (outStr string) {
		path, err := mainObjects.fileListstore.GetValue(iter, 4)
		if err == nil {
			name, err := mainObjects.fileListstore.GetValue(iter, 0)
			if err == nil {
				strPath, err := path.GetString()
				if err == nil {
					strName, err := name.GetString()
					if err == nil {
						return filepath.Join(strPath, strName)
					}
				}
			}
		}
		return ""
	}

	// changeOrder: called when columns are sorted.
	var changeOrder = func() {
		mainOptions.foundFilesList = mainOptions.foundFilesList[:0]
		iter, ok := mainObjects.fileListstore.GetIterFirst()
		for ok {
			mainOptions.foundFilesList = append(mainOptions.foundFilesList, getIterValue(iter))
			//			fmt.Println(getIterValue(iter))
			ok = mainObjects.fileListstore.IterNext(iter)
		}
		//		fmt.Println("reordered !!!")
	}

	// Refresh slice containing filepath to reflect treeview order
	model, _ := mainObjects.SearchTreeview.GetModel()
	model.Connect("rows-reordered", changeOrder)

	// -----------************

	// sel, _ := mainObjects.SearchTreeview.GetSelection()
	// sel.SetMode(gtk.SELECTION_MULTIPLE)
	mainObjects.SearchTreeview.SetProperty("activate-on-single-click", false)

	for idx, _ := range columnsNames {
		column := mainObjects.SearchTreeview.GetColumn(idx)
		column.SetSizing(gtk.TREE_VIEW_COLUMN_AUTOSIZE)
		column = mainObjects.SearchTreeview.GetColumn(idx)
		column.SetSizing(gtk.TREE_VIEW_COLUMN_AUTOSIZE)
		column = mainObjects.SearchTreeview.GetColumn(idx)
		column.SetSizing(gtk.TREE_VIEW_COLUMN_AUTOSIZE)
		column = mainObjects.SearchTreeview.GetColumn(idx)
		column.SetSizing(gtk.TREE_VIEW_COLUMN_AUTOSIZE)
		column = mainObjects.SearchTreeview.GetColumn(idx)
		column.SetSizing(gtk.TREE_VIEW_COLUMN_AUTOSIZE)
		// column := mainObjects.SearchTreeview.GetColumn(0)
		// column.SetSortColumnID(-1)
	}
}
