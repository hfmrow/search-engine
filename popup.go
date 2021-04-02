// popup.go

/*
*	©2019 H.F.M. MIT license
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"gopkg.in/djherbis/times.v1"

	glfs "github.com/hfmrow/genLib/files"
	glsg "github.com/hfmrow/genLib/strings"
	glts "github.com/hfmrow/genLib/tools"
)

// initPopup: The popup for the TreeView.
func initPopup() *gtk.Menu {
	popupMenu = PopupMenuIconStructNew()

	popupMenu.AddItem("_Launch/Open", launchFile, popupMenu.OPT_ICON|popupMenu.OPT_NORMAL, play20)
	popupMenu.AddItem("_Web browser", openInBrowser, popupMenu.OPT_ICON|popupMenu.OPT_NORMAL, globalNetwork20)
	popupMenu.AddItem("Open _directory", openDir, popupMenu.OPT_ICON|popupMenu.OPT_NORMAL, folder48)
	popupMenu.AddItem("_Copy to clipboard", toClipboard, popupMenu.OPT_ICON|popupMenu.OPT_NORMAL, copyDocument20)
	popupMenu.AddItem("", nil, popupMenu.OPT_SEPARATOR)
	popupMenu.AddItem("Time to _Newer than", func() { toCalendar(false) }, popupMenu.OPT_ICON|popupMenu.OPT_NORMAL, calendarPers48)
	popupMenu.AddItem("Time to _Older than", func() { toCalendar(true) }, popupMenu.OPT_ICON|popupMenu.OPT_NORMAL, calendarPers48)
	popupMenu.AddItem("", nil, popupMenu.OPT_SEPARATOR)
	popupMenu.AddItem("_Delete", deleteEntry, popupMenu.OPT_ICON|popupMenu.OPT_NORMAL, stop48)
	return popupMenu.MenuBuild()
}

// toCalendar: default, set to newer than otherwise, set to older than
func toCalendar(toOlder bool) {
	var err error
	var selected [][]string
	if selected, err = tvs.GetSelectedRows(); err == nil {
		for _, entry := range storeFoundFiles {
			if entry.FilePath == filepath.Join(selected[0][4], selected[0][0]) {
				var timeInfo time.Time
				infos := times.Get(entry.FileInfo)
				switch mainObjects.SearchComboboxTextDateType.GetActive() {
				case 0, 2:
					timeInfo = infos.AccessTime()
				case 1, 3:
					timeInfo = infos.ModTime()
				}
				if mainObjects.SearchComboboxTextDateZone.GetActive() == 0 {
					timeInfo = timeInfo.UTC()
				}
				if !toOlder {
					mainOptions.calDataNewerThan.FromTime(timeInfo)
					setCalendarbuttonLabel(mainObjects.SearchButtonNewerThan, mainOptions.calDataNewerThan)
				} else {
					mainOptions.calDataOlderThan.FromTime(timeInfo)
					setCalendarbuttonLabel(mainObjects.SearchButtonOlderThan, mainOptions.calDataOlderThan)
				}
			}
		}
	}
}

// toClipboard: Copy path(s) to clipboard
func toClipboard() {
	var err error
	var selected [][]string
	var tmpFileNames string
	if selected, err = tvs.GetSelectedRows(); err == nil {
		for _, row := range selected {
			tmpFileNames += filepath.Join(row[4], row[0]) + "\n"
		}
	} else if err != errNoSelection { // don't warn on "unselected row" case.
		DialogMessage(mainObjects.MainWindow, "info", "Information !",
			"You got a problem with some selected row(s)\n\n"+err.Error()+"\n",
			"", "Ok")
	}
	clipboard.SetText(strings.Trim(tmpFileNames, glsg.GetTextEOL([]byte(tmpFileNames))))
	clipboard.Store()
}

// openInBrowser:
func openInBrowser() {

	var (
		selected [][]string
		err      error
		filename string
	)

	if selected, err = tvs.GetSelectedRows(); err == nil {
		filename = mainOptions.WebSearchEngine + glfs.BaseNoExt(filepath.Base(selected[0][0]))
		open(filename)
	} /* else {

	}*/

	// var err error
	// var result []byte
	// var selected [][]string
	// if selected, err = tvs.GetSelectedRows(); err == nil {
	// 	// Using Goroutine to freeing current thread and make it independent.
	// 	go glib.IdleAdd(func() { // IdleAdd run to permit gtk3 working right during goroutine
	// 		openresult, err = glts.ExecCommand(mainOptions.AppLauncher,
	// 			mainOptions.WebSearchEngine+glfs.BaseNoExt(filepath.Base(selected[0][0])))
	// 	})
	// }
}

// launchFile:
func launchFile() {

	var (
		selected [][]string
		err      error
		filename string
	)

	if selected, err = tvs.GetSelectedRows(); err == nil {
		filename = filepath.Join(selected[0][4], selected[0][0])
		open(filename)
	} /* else {

	}*/
}

// openDir:
func openDir() {

	var (
		selected [][]string
		err      error
		filename string
	)

	if selected, err = tvs.GetSelectedRows(); err == nil {
		filename = filepath.Join(selected[0][4], selected[0][0])
		open(filepath.Dir(filename))
	} /* else {

	}*/
}

// open: show file or dir depending on "path".
func open(path string) {

	var goFunc = func() {
		if outTerm, err := glts.ExecCommand([]string{mainOptions.AppLauncher, path}); err != nil {

			// Error is handled by "xdg-open" command

			fmt.Println(err, outTerm)
		}
	}
	// IdleAdd to permit gtk3 working right during goroutine
	glib.IdleAdd(func() {
		// Using goroutine to permit the usage of another
		// thread and freeing the current one.
		go goFunc()
	})
}

// deleteEntry:
func deleteEntry() {
	var err error
	var tmpFileNames string
	var iters []*gtk.TreeIter
	// Iters needed to be able to remove selected rows
	if iters = tvs.GetSelectedIters(); len(iters) > 0 {
		for _, iter := range iters {
			tmpFileNames += tvs.GetColValue(iter, 0).(string) + "\n" // Get filename
		}
		if DialogMessage(mainObjects.MainWindow, "info", "Information !",
			fmt.Sprintf("\nAre you sure, want you to delete %d file(s):\n\n", len(iters))+strings.Trim(tmpFileNames, "\n"),
			"", sts["yes"], sts["no"]) == 0 {

			for _, iter := range iters {
				if err = os.RemoveAll(
					// Get Full path
					filepath.Join(tvs.GetColValue(iter, 4).(string),
						tvs.GetColValue(iter, 0).(string))); err == nil || os.IsNotExist(err) {
					// do it one by one in case of error, process will be stopped.
					tvs.RemoveSelectedRows(iter)
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
