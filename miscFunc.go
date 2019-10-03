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
	/* Initialiste liststore Columns */
	if tvs, err = gitw.TreeViewStructureNew(mainObjects.SearchTreeview, true, false); err == nil {
		for _, col := range columnsNames {
			tvs.AddColumn(col[0], col[1], true, true, true, true, true, true)
		}
		err = tvs.StoreSetup(new(gtk.ListStore))
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
	// statusbar.Set(glsg.TruncatePath(mainOptions.LastDirectory, 3), 3)
	if len(timer.Results) > 1 {
		statusbar.Set(fmt.Sprintf("%dm %ds %dms", timer.NumTime[0].Min, timer.NumTime[0].Sec, timer.NumTime[0].Ms), 4)
		statusbar.Set(fmt.Sprintf("%dm %ds %dms", timer.NumTime[1].Min, timer.NumTime[1].Sec, timer.NumTime[1].Ms), 5)
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
		mainOptions.AboutOptions.AppName,
		mainOptions.AboutOptions.AppVers,
		mainOptions.AboutOptions.YearCreat,
		mainOptions.AboutOptions.AppCreats,
		mainOptions.AboutOptions.LicenseShort,
		mainOptions.AboutOptions.Description,
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
		if ad, err = gtk.AdjustmentNew(0, 0, 23, 1, 0, 0); err == nil {
			mainObjects.TimeSpinbuttonHourNewer.Configure(ad, 1, 0)
			if ad, err = gtk.AdjustmentNew(0, 0, 59, 1, 0, 0); err == nil {
				mainObjects.TimeSpinbuttonHourOlder.Configure(ad, 1, 0)
				if ad, err = gtk.AdjustmentNew(0, 0, 59, 1, 0, 0); err == nil {
					mainObjects.TimeSpinbuttonMinutsNewer.Configure(ad, 1, 0)
					if ad, err = gtk.AdjustmentNew(0, 0, 23, 1, 0, 0); err == nil {
						mainObjects.TimeSpinbuttonMinutsOlder.Configure(ad, 1, 0)
						if ad, err = gtk.AdjustmentNew(0, 0, 59, 1, 0, 0); err == nil {
							mainObjects.TimeSpinbuttonSecondsNewer.Configure(ad, 1, 0)
							if ad, err = gtk.AdjustmentNew(0, 0, 59, 1, 0, 0); err == nil {
								mainObjects.TimeSpinbuttonSecondsOlder.Configure(ad, 1, 0)
							}
						}
					}
				}
			}
		}
	}
	if err != nil {
		Check(err)
	}
}
