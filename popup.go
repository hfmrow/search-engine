// popup.go

/*
*	©2019 H.F.M. MIT license
 */

package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	glfs "github.com/hfmrow/genLib/files"
	glts "github.com/hfmrow/genLib/tools"
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

	menuItem.Connect("activate", func() { launchFile() })
	menuItem1.Connect("activate", func() { go openInBrowser() })
	menuItem2.Connect("activate", func() { openDir() })
	menuItem3.Connect("activate", func() { toClipboard() })
	menuItem4.Connect("activate", func() { deleteEntry() })
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
func toClipboard() {
	var err error
	var selected [][]string
	var tmpFileNames string
	if selected, err = getSelectedAsString(); err == nil {
		for _, row := range selected {
			tmpFileNames += filepath.Join(row[4], row[0]) + "\n"
		}
	} else if err != errNoSelection { // don't warn on "unselected row" case.
		DialogMessage(mainObjects.MainWindow, "info", "Information !",
			"You got a problem with some selected row(s)\n\n"+err.Error()+"\n",
			"", "Ok")
	}

	clipboard.SetText(tmpFileNames)
	clipboard.Store()
}

// getSelectedAsString: Retrieve values of ll selected rows
func getSelectedAsString() (outSlice [][]string, err error) {
	var iters []*gtk.TreeIter
	if iters, err = tvs.GetSelectedIters(); err == nil {
		outSlice, err = tvs.GetSelectedRows(iters...)
	}
	if len(outSlice) == 0 {
		return outSlice, errNoSelection
	}
	return
}

// Open in browser
func openInBrowser() {
	var err error
	var result []byte
	var selected [][]string
	if selected, err = getSelectedAsString(); err == nil {
		_, _ = glib.IdleAdd(func() { // idle run to permit gtk3 working right during goroutine
			result, err = glts.ExecCommand(mainOptions.AppLauncher,
				mainOptions.WebSearchEngine+glfs.BaseNoExt(filepath.Base(selected[0][0])))
			if err != nil && err != errNoSelection { // don't warn on "unselected row" case.
				DialogMessage(mainObjects.MainWindow, "info", "Information !",
					"You got a problem with:\n"+filepath.Join(selected[0][4], selected[0][0])+
						"\n\n"+err.Error()+"\n\nTerminal output:\n"+string(result), "", "Ok")
			}
		})
	}
}

// Run filename
func launchFile() {
	var selected [][]string
	var err error
	var filename string
	if selected, err = getSelectedAsString(); err == nil {
		glib.IdleAdd(func() { // idle run to permit gtk3 working right during goroutine
			filename = filepath.Join(selected[0][4], selected[0][0])
			go glts.ExecCommand(mainOptions.AppLauncher, filename)
		})
	}
}

// Open directory filename
func openDir() {
	var selected [][]string
	var err error
	var filename string
	if selected, err = getSelectedAsString(); err == nil {
		glib.IdleAdd(func() { // idle run to permit gtk3 working right during goroutine
			filename = filepath.Join(selected[0][4], selected[0][0])
			go glts.ExecCommand(mainOptions.AppLauncher, filepath.Dir(filename))
		})
	}
}

// deleteEntry
func deleteEntry() {
	var err error
	var selected [][]string
	var tmpFileNames string
	var iters []*gtk.TreeIter
	if iters, err = tvs.GetSelectedIters(); err == nil {
		if selected, err = tvs.GetSelectedRows(iters...); err == nil {
			for _, row := range selected {
				tmpFileNames += filepath.Join(row[4], row[0]) + "\n"
			}
			if DialogMessage(mainObjects.MainWindow, "info", "Information !",
				"\nAre you sure, you want to delete file(s):\n\n"+strings.Trim(tmpFileNames, "\n"),
				"", sts["yes"], sts["no"]) == 0 {

				for idx, row := range selected {
					if err = os.Remove(filepath.Join(row[4], row[0])); err == nil {
						err = tvs.RemoveSelectedRows(iters[idx])
					}
					if err != nil {
						break
					}
				}
			}
		}
	}
	if err != nil && err != errNoSelection { // don't warn on "unselected row" case.
		DialogMessage(mainObjects.MainWindow, "info", "Information !",
			"You got a problem with some selected row(s)\n\n"+err.Error()+"\n",
			"", "Ok")
	}
}
