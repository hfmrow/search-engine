// popup.go

/*
*	©2019 H.F.M. MIT license
 */

package main

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gtk"
	g "github.com/hfmrow/searchEngine/genLib"
	gi "github.com/hfmrow/searchEngine/gtk3Import"
)

// initPopup: The popup for TreeView control ...
func initPopup(tw *gtk.TreeView, ls *gtk.ListStore) *gtk.Menu {

	var MenuItemNewWithImage = func(label string, icon interface{}) (menuItem *gtk.MenuItem, err error) {
		box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
		if err == nil {
			image, err := gtk.ImageNew()
			if err == nil {
				setImage(image, icon, 14)
				label, err := gtk.LabelNewWithMnemonic(label)
				if err == nil {
					menuItem, err = gtk.MenuItemNew()
					if err == nil {
						label.SetHAlign(gtk.ALIGN_START)
						box.Add(image)
						box.PackEnd(label, true, true, 8)
						menuItem.Container.Add(box)
						menuItem.ShowAll()
					}
				}
			}
		}
		return menuItem, err
	}

	var getRowNb = func() (rowNb []int) {
		selected, err := tw.GetSelection()
		Check(err)
		rows := selected.GetSelectedRows(ls)
		for l := rows; l != nil; l = l.Next() {
			path := l.Data().(*gtk.TreePath)
			rowNb = append(rowNb, path.GetIndices()[0])
		}
		return rowNb
	}
	/*	Make menu	*/
	menu, err := gtk.MenuNew()
	Check(err)
	var menuItem, menuItem1, menuItem2, menuItem3, menuItem4 *gtk.MenuItem
	var separatorMenuItem *gtk.SeparatorMenuItem
	if menuItem, err = MenuItemNewWithImage("_Launch/Open", play20); err == nil {
		if menuItem1, err = MenuItemNewWithImage("_Web browser", globalNetwork20); err == nil {
			if menuItem2, err = MenuItemNewWithImage("_Open directory", folderOpen20); err == nil {
				if menuItem3, err = MenuItemNewWithImage("_Copy path", copyDocument20); err == nil {
					if separatorMenuItem, err = gtk.SeparatorMenuItemNew(); err == nil {
						menuItem4, err = MenuItemNewWithImage("_Delete", signError20)
					}
				}
			}
		}
	}
	Check(err)

	menuItem.Connect("activate", func() { go launchFile(getRowNb()) })
	menuItem1.Connect("activate", func() { go openInBrowser(getRowNb()) })
	menuItem2.Connect("activate", func() { go openDir(getRowNb()) })
	menuItem3.Connect("activate", func() { toClipboard(getRowNb()) })
	menuItem4.Connect("activate", func() { deleteEntry(getRowNb()) })
	/*	Add options items to menu	*/
	menu.Append(menuItem)
	menu.Append(menuItem1)
	menu.Append(menuItem2)
	menu.Append(menuItem3)
	menu.Append(separatorMenuItem)
	menu.Append(menuItem4)
	/*	Show menu	*/
	menu.ShowAll()
	return menu
}

// Copy path to clipboard
func toClipboard(rows []int) {
	var tmpFileNames string
	for _, row := range rows {
		tmpFileNames += mainOptions.foundFilesList[row] + "\n"
	}
	clipboard.SetText(tmpFileNames)
	clipboard.Store()
}

// Open in browser
func openInBrowser(rows []int) {
	result, err := g.ExecCommand(mainOptions.AppLauncher,
		mainOptions.SearchEngine+g.SplitFilepath(mainOptions.foundFilesList[rows[0]]).BaseNoExt)
	if err != nil {
		gi.DlgMessage(mainObjects.MainWindow, "info", "Information !",
			"You got a problem with:\n"+mainOptions.foundFilesList[rows[0]]+"\n\n"+err.Error()+"\n\nTerminal output:\n"+string(result),
			"", "Ok")
	}
}

// Run filename
func launchFile(rows []int) {
	for _, row := range rows {
		command := mainOptions.foundFilesList[row]
		result, err := g.ExecCommand(mainOptions.AppLauncher, command)
		if err != nil {
			gi.DlgMessage(mainObjects.MainWindow, "info", "Information !",
				"You got a problem with:\n"+mainOptions.foundFilesList[row]+"\nCommand: "+command+"\n"+err.Error()+"\n\nTerminal output:\n"+string(result),
				"", "Ok")
		}
	}
}

// Open directory filename
func openDir(rows []int) {
	result, err := g.ExecCommand(mainOptions.AppLauncher, g.SplitFilepath(mainOptions.foundFilesList[rows[0]]).Path)
	if err != nil {
		gi.DlgMessage(mainObjects.MainWindow, "info", "Information !",
			"You got a problem with:\n"+mainOptions.foundFilesList[rows[0]]+"\n\n"+err.Error()+"\n\nTerminal output:\n"+string(result),
			"", "Ok")
	}
}

// deleteEntry
func deleteEntry(rows []int) {
	if gi.DlgMessage(mainObjects.MainWindow, "info", "Information", fmt.Sprintf("\n\nRemove current selected files ? %d", len(rows)), "", "Cancel", "Ok") != 0 {
		for idx := len(rows) - 1; idx >= 0; idx-- {
			fileName := mainOptions.foundFilesList[rows[idx]]
			fmt.Printf("Removing '%s'\n", fileName)
			if err := os.Remove(fileName); err == nil {
				mainOptions.foundFilesList = append(mainOptions.foundFilesList[:rows[idx]], mainOptions.foundFilesList[rows[idx]+1:]...)
			} else {
				gi.DlgMessage(mainObjects.MainWindow, "error", "Error", fmt.Sprintf("\n\nUnable to remove:\n%s\n\n%s", fileName, err.Error()), "", "Ok")
			}
		}
		doDisplay()
	}
}
