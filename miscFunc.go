// miscFunc.go

/*
*	©2019 H.F.M. MIT license
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gotk3/gotk3/gtk"

	glsg "github.com/hfmrow/genLib/strings"

	gitw "github.com/hfmrow/gtk3Import/treeview"
)

// initTreeview
func initTreeview() {
	var err error
	/* Init liststore Columns */
	if tvs, err = gitw.TreeViewStructureNew(mainObjects.SearchTreeview, true, false); err == nil {
		tvs.AddColumns(columnsNames, true, true, true, true, true, true)

		tvs.SelectionChangedFunc = updateStatusbar

		if err = tvs.StoreSetup(new(gtk.ListStore)); err == nil {

			// Assign sorted column (col 1 will be sorted using values of col 5)
			tvs.Columns[columnsMap["size"]].Column.SetSortColumnID(columnsMap["sizeSort"])
			tvs.Columns[columnsMap["time"]].Column.SetSortColumnID(columnsMap["dateSort"])

			// Callback function for selection "changed" signal.
			tvs.SelectionChangedFunc = updateStatusbar
		}
	}
	if err != nil {
		Check(err)
	}
}

// updateStatusbar
func updateStatusbar() {
	statusbar.Set(fmt.Sprintf("%d files scanned", filesScanned), 0)
	statusbar.Set(fmt.Sprintf("%d", tvs.CountRows()), 1)
	statusbar.Set(fmt.Sprintf("%d", tvs.Selection.CountSelectedRows()), 2)

	if len(timer.Lapses) > 1 {
		statusbar.Set(fmt.Sprintf("%dm %ds %dms", timer.Lapses[0].Min, timer.Lapses[0].Sec, timer.Lapses[0].Ms), 4)
		statusbar.Set(fmt.Sprintf("%dm %ds %dms", timer.Lapses[1].Min, timer.Lapses[1].Sec, timer.Lapses[1].Ms), 5)
	}
	titlebar.Update([]string{glsg.TruncatePath(mainOptions.LastDirectory, 3)})
}

// Handling arguments from command line.
func cmdLineParser() {
	if len(os.Args) > 1 {
		if fi, err := os.Stat(os.Args[1]); !os.IsNotExist(err) {
			if fi.IsDir() {
				mainOptions.LastDirectory = os.Args[1]
				return
			}
			mainOptions.LastDirectory = filepath.Dir(os.Args[1])
			return
		}
	}
	name := filepath.Base(os.Args[0])
	fmt.Printf("%s %s %s %s\n%s\n%s\n\nUsage: %s \"PathToScan\"\n",
		mainOptions.About.AppName,
		mainOptions.About.AppVers,
		mainOptions.About.YearCreat,
		mainOptions.About.AppCreats,
		mainOptions.About.LicenseShort,
		mainOptions.About.Description,
		name)
}

// initComboboxes
func initComboboxes() {
	mainObjects.SearchComboboxTextAnd.SetActive(0)
	mainObjects.SearchComboboxTextNot.SetActive(0)
	mainObjects.SearchComboboxTextOr.SetActive(0)
	mainObjects.SearchComboboxTextType.SetActive(1)
}

// initSpinButtons
func initSpinButtons() {
	var err error
	var ad *gtk.Adjustment
	if ad, err = gtk.AdjustmentNew(-1, -1, 100, 1, 0, 0); err == nil {
		mainObjects.SearchSpinbuttonDepth.Configure(ad, 1, 0)
	}
	if err != nil {
		Check(err)
	}
}
