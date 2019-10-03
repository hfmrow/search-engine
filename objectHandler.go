// objectHandler.go

/*
	©2018 H.F.M. MIT license
*/

package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	glfs "github.com/hfmrow/genLib/files"
	glsg "github.com/hfmrow/genLib/strings"
	gidgcr "github.com/hfmrow/gtk3Import/dialog/chooser"
)

func SearchFilechooserbuttonFileSet() {
	mainOptions.LastDirectory = mainObjects.SearchFilechooserbutton.GetFilename()
	updateStatusbar()
}

// Let's go search
func SearchButtonClicked() {
	if err := fillListstore(); err != nil {
		DialogMessage(mainObjects.MainWindow, "error", "Error occured during search", "\n\n"+err.Error(), "", "Ok")
	}
}

// SearchTreeviewButtonPressEvent
func SearchTreeviewButtonPressEvent(tw *gtk.TreeView, event *gdk.Event) bool {
	eventButton := gdk.EventButtonNewFromEvent(event)
	selection, err := tw.GetSelection()
	Check(err) // Is there anything out there ?
	if selection.CountSelectedRows() > 0 && err == nil {
		if eventButton.Button() == 3 {
			mainObjects.popUpMenu.PopupAtPointer(event)

			// return true to stop propagate event, so nothing get selected on RMB click
			return true
		}
	}
	// return false to propagate event
	return false
}

// Save results as text file
func SearchButtonExportClicked() {
	var err error
	var ok bool
	var filename string

	if len(mainOptions.foundFilesList) != 0 {
		filename, ok, err = gidgcr.FileChooser(mainObjects.MainWindow, "save", "", "")
		if ok {
			err = glfs.WriteFile(filename,
				[]byte(strings.Join(mainOptions.foundFilesList, glsg.GetOsLineEnd())))
		} else {
			if len(filename) != 0 && ok {
				err = errors.New("Error writing file ...")
			}
		}
	} else {
		err = errors.New("Nothing to save ...")
	}
	if err != nil {
		DialogMessage(mainObjects.MainWindow, "info", "Attention !!!", "\nYou got a problem:\n\n"+err.Error(), "", "Ok")
	}
}

// Handle SearchButtonNewerThanClicked
func SearchButtonNewerThanClicked() {
	displayTimeWin(mainObjects.TimeWindowNewer, "Choose date time for newer than files")
	setCal(mainObjects.TimeCalendarNewer, &mainOptions.searchNewerThan,
		mainObjects.TimeSpinbuttonHourNewer,
		mainObjects.TimeSpinbuttonMinutsNewer,
		mainObjects.TimeSpinbuttonSecondsNewer)
}

// Handle SearchButtonOlderThanClicked
func SearchButtonOlderThanClicked() {
	displayTimeWin(mainObjects.TimeWindowOlder, "Choose date time for older than files")
	setCal(mainObjects.TimeCalendarOlder, &mainOptions.searchOlderThan,
		mainObjects.TimeSpinbuttonHourOlder,
		mainObjects.TimeSpinbuttonMinutsOlder,
		mainObjects.TimeSpinbuttonSecondsOlder)
}

//Switch visibility when Regex behind used
func SearchCheckbuttonRegexToggled() {
	mainObjects.SearchCheckbuttonWordAnd.SetSensitive(!mainObjects.SearchCheckbuttonRegex.GetActive())
	mainObjects.SearchCheckbuttonWordOr.SetSensitive(!mainObjects.SearchCheckbuttonRegex.GetActive())
	mainObjects.SearchCheckbuttonWordNot.SetSensitive(!mainObjects.SearchCheckbuttonRegex.GetActive())

	mainObjects.SearchCheckbuttonCaseSensitive.SetSensitive(!mainObjects.SearchCheckbuttonRegex.GetActive())
	mainObjects.SearchCheckbuttonWildCard.SetSensitive(!mainObjects.SearchCheckbuttonRegex.GetActive())
	mainObjects.SearchComboboxTextOr.SetSensitive(!mainObjects.SearchCheckbuttonRegex.GetActive())
	mainObjects.SearchComboboxTextNot.SetSensitive(!mainObjects.SearchCheckbuttonRegex.GetActive())
	mainObjects.SearchCheckbuttonCharClasses.SetSensitive(!mainObjects.SearchCheckbuttonRegex.GetActive())
}

// Switch visibility when Char classes behind used
func SearchCheckbuttonCharClassesClicked() {
	mainObjects.SearchCheckbuttonCharClassesStrictMode.SetVisible(mainObjects.SearchCheckbuttonCharClasses.GetActive())
	mainObjects.SearchCheckbuttonWildCard.SetSensitive(!mainObjects.SearchCheckbuttonCharClasses.GetActive())
	mainObjects.SearchCheckbuttonRegex.SetSensitive(!mainObjects.SearchCheckbuttonCharClasses.GetActive())
}

// Reset comboBoxeS
func SearchButtonResetComboEntryClicked() {
	/* remove all entries from gtk object */
	fillAllComboboxText(true)
	/* remove all entries from internal variables */
	mainOptions.SearchList.And = mainOptions.SearchList.And[:0]
	mainOptions.SearchList.Or = mainOptions.SearchList.Or[:0]
	mainOptions.SearchList.Not = mainOptions.SearchList.Not[:0]
	/* updating option file with cleaned entries */
	mainOptions.Write()
	fmt.Println("Combobox entries reseted ...")
}

func SearchComboboxTextTypeChanged() {
	if mainOptions.UpdateOnChanges {
		if err := fillListstore(); err != nil {
			DialogMessage(mainObjects.MainWindow, "error", "Error occured during search", "\n\n"+err.Error(), "", "Ok")
		}
	}
}

func SearchComboboxTextDateTypeChanged() {
	if mainOptions.UpdateOnChanges {
		doDisplay()
	}
}
