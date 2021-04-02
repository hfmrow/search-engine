// objectHandler.go

/*
	©2018 H.F.M. MIT license
*/

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gotk3/gotk3/gtk"

	glsg "github.com/hfmrow/genLib/strings"
	gidgcr "github.com/hfmrow/gtk3Import/dialog/chooser"
)

func SearchFilechooserbuttonFileSet() {
	mainOptions.LastDirectory = mainObjects.SearchFilechooserbutton.GetFilename()
	updateStatusbar()
}

// Let's go search
func SearchButtonClicked() {
	fillListstore()
}

// Save results as text file
func SearchButtonExportClicked() {
	var err error
	var ok bool
	var filename string
	var tmpFilesList string

	if len(storeFoundFiles) != 0 {
		filename, ok, err = gidgcr.FileChooser(mainObjects.MainWindow, "save", "", "")
		if ok {
			for _, file := range storeFoundFiles {
				tmpFilesList += file.FilePath + glsg.GetOsLineEnd()
			}
			err = ioutil.WriteFile(filename, []byte(tmpFilesList), os.ModePerm)
		} else {
			if len(filename) != 0 && ok {
				err = errors.New("Error writing file ...")
			}
		}
	} else {
		err = errors.New("There is nothing to export ...")
	}
	if err != nil {
		DialogMessage(mainObjects.MainWindow, "info", "Attention !!!", "\nYou got a problem:\n\n"+err.Error(), "", "Ok")
	}
}

// TODO find a way to make it working ...
func timeControl(older, newer time.Time) (out bool) {
	// var mess = DialogMessage(mainObjects.MainWindow, "error", "Error date selection", "\n\nDate time selection will never retrieve any file.", "", "Ok")
	// var blankTime time.Time
	// if newer != blankTime || older != blankTime {
	// 	switch {
	// 	case newer.After(older):
	// 		mess()
	// 		return false
	// 	case time.Now().After(older):
	// 		return true
	// 	case newer.After(time.Now()):
	// 		mess()
	// 		return false
	// 	}
	// } else {
	// 	return true
	// }
	// mess()
	return true
}

// initCalData: Init calendar
func initCalData() {
	mainOptions.calDataNewerThan = gidgcr.CalendarDataNew()
	mainOptions.calDataNewerThan.Init() // Set values to nul. is needed for files search

	mainOptions.calDataOlderThan = gidgcr.CalendarDataNew()
	mainOptions.calDataOlderThan.Init() // Set values to nul. is needed for files search
}

// Handle SearchButtonNewerThanClicked
func SearchButtonNewerThanClicked() {

	if mainOptions.calDataNewerThan.ToTime() == mainOptions.calDataNewerThan.BlankTime {
		mainOptions.calDataNewerThan.Init(time.Now())
	}
	mainOptions.calendar.TitleWindow = "Choose newer than date time"
	mainOptions.calendar.Result = mainOptions.calDataNewerThan
	ok, err := mainOptions.calendar.Run()
	if err != nil {
		log.Fatalf("SearchButtonNewerThanClicked: %s\n", err.Error())
	}
	if ok && timeControl(mainOptions.calDataOlderThan.ToTime(), mainOptions.calDataNewerThan.ToTime()) {
		mainOptions.calDataNewerThan = mainOptions.calendar.Result
		setCalendarbuttonLabel(mainObjects.SearchButtonNewerThan, mainOptions.calDataNewerThan)
	} else {
		mainOptions.calDataNewerThan.Init()
		mainObjects.SearchButtonNewerThan.SetLabel(origLabelNT)
	}
	fmt.Println(mainOptions.calDataNewerThan.ToLayout())
}

// setCalendarbuttonLabel: fill label of desined button with content of CalendarData struct
func setCalendarbuttonLabel(button *gtk.Button, calData *gidgcr.CalendarData) {
	button.SetLabel(fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d",
		calData.Year,
		calData.Month,
		calData.Day,
		calData.Hour,
		calData.Min,
		calData.Sec))
}

// Handle SearchButtonOlderThanClicked
func SearchButtonOlderThanClicked() {

	if mainOptions.calDataOlderThan.ToTime() == mainOptions.calDataOlderThan.BlankTime {
		mainOptions.calDataOlderThan.Init(time.Now())
	}
	mainOptions.calendar.TitleWindow = "Choose older than date time"
	mainOptions.calendar.Result = mainOptions.calDataOlderThan
	ok, err := mainOptions.calendar.Run()
	if err != nil {
		log.Fatalf("SearchButtonOlderThanClicked: %s\n", err.Error())
	}
	if ok && timeControl(mainOptions.calDataOlderThan.ToTime(), mainOptions.calDataNewerThan.ToTime()) {
		setCalendarbuttonLabel(mainObjects.SearchButtonOlderThan, mainOptions.calDataOlderThan)
	} else {
		mainOptions.calDataOlderThan.Init()
		mainObjects.SearchButtonOlderThan.SetLabel(origLabelOT)
	}
	fmt.Println(mainOptions.calDataOlderThan.ToLayout())
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
		fillListstore()
	}
}

func ComboboxTextGenericChanged() {
	doDisplay()
}

// genericHideWindow: Signal handler delete_event (hidding window)
func genericHideWindow(w *gtk.Window) bool {
	if w.GetVisible() {
		w.Hide()
	}
	return true
}

// TopImageEventboxClicked: display Aboutbox
func TopImageEventboxClicked() {
	mainOptions.About.Show()
}
