// gohObjects.go

// Source file auto-generated on Wed, 10 Jul 2019 21:07:55 using Gotk3ObjHandler v1.3 ©2019 H.F.M

/*
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"github.com/gotk3/gotk3/gtk"
)

// Control over all used objects from glade.
var mainObjects *MainControlsObj

/******************************/
/* Main structure Declaration */
/******************************/
type MainControlsObj struct {
	fileListstore                          *gtk.ListStore /*MANUAL*/
	ImageTop                               *gtk.Image
	mainUiBuilder                          *gtk.Builder
	MainWindow                             *gtk.Window
	popUpMenu                              *gtk.Menu /*MANUAL*/
	SearchButton                           *gtk.Button
	SearchButtonExit                       *gtk.Button
	SearchButtonExport                     *gtk.Button
	SearchButtonNewerThan                  *gtk.Button
	SearchButtonOlderThan                  *gtk.Button
	SearchButtonResetComboEntry            *gtk.Button
	SearchCheckbuttonCaseSensitive         *gtk.CheckButton
	SearchCheckbuttonCharClasses           *gtk.CheckButton
	SearchCheckbuttonCharClassesStrictMode *gtk.CheckButton
	SearchCheckbuttonRegex                 *gtk.CheckButton
	SearchCheckbuttonSplitedAnd            *gtk.CheckButton
	SearchCheckbuttonSplitedNot            *gtk.CheckButton
	SearchCheckbuttonSplitedOr             *gtk.CheckButton
	SearchCheckbuttonWildCard              *gtk.CheckButton
	SearchCheckbuttonWordAnd               *gtk.CheckButton
	SearchCheckbuttonWordNot               *gtk.CheckButton
	SearchCheckbuttonWordOr                *gtk.CheckButton
	SearchComboboxTextAnd                  *gtk.ComboBoxText
	SearchComboboxTextDateType             *gtk.ComboBoxText
	SearchComboboxTextEntryAnd             *gtk.Entry
	SearchComboboxTextEntryNot             *gtk.Entry
	SearchComboboxTextEntryOr              *gtk.Entry
	SearchComboboxTextNot                  *gtk.ComboBoxText
	SearchComboboxTextOr                   *gtk.ComboBoxText
	SearchComboboxTextType                 *gtk.ComboBoxText
	SearchFilechooserbutton                *gtk.FileChooserButton
	SearchSpinbuttonDepth                  *gtk.SpinButton
	SearchTreeview                         *gtk.TreeView
	Statusbar                              *gtk.Statusbar
	TimeButtonOkNewer                      *gtk.Button
	TimeButtonOkOlder                      *gtk.Button
	TimeButtonResetNewer                   *gtk.Button
	TimeButtonResetOlder                   *gtk.Button
	TimeCalendarNewer                      *gtk.Calendar
	TimeCalendarOlder                      *gtk.Calendar
	TimeImageTopNewer                      *gtk.Image
	TimeImageTopOlder                      *gtk.Image
	TimeSpinbuttonHourNewer                *gtk.SpinButton
	TimeSpinbuttonHourOlder                *gtk.SpinButton
	TimeSpinbuttonMinutsNewer              *gtk.SpinButton
	TimeSpinbuttonMinutsOlder              *gtk.SpinButton
	TimeSpinbuttonSecondsNewer             *gtk.SpinButton
	TimeSpinbuttonSecondsOlder             *gtk.SpinButton
	TimeWindowNewer                        *gtk.Window
	TimeWindowOlder                        *gtk.Window
	TopImageEventbox                       *gtk.EventBox
}

/******************************/
/* GtkObjects  Initialisation */
/******************************/
// gladeObjParser: Initialise Gtk3 Objects into mainObjects structure.
func gladeObjParser() {
	mainObjects.ImageTop = loadObject("ImageTop").(*gtk.Image)
	mainObjects.MainWindow = loadObject("MainWindow").(*gtk.Window)
	mainObjects.SearchButton = loadObject("SearchButton").(*gtk.Button)
	mainObjects.SearchButtonExit = loadObject("SearchButtonExit").(*gtk.Button)
	mainObjects.SearchButtonExport = loadObject("SearchButtonExport").(*gtk.Button)
	mainObjects.SearchButtonNewerThan = loadObject("SearchButtonNewerThan").(*gtk.Button)
	mainObjects.SearchButtonOlderThan = loadObject("SearchButtonOlderThan").(*gtk.Button)
	mainObjects.SearchButtonResetComboEntry = loadObject("SearchButtonResetComboEntry").(*gtk.Button)
	mainObjects.SearchCheckbuttonCaseSensitive = loadObject("SearchCheckbuttonCaseSensitive").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonCharClasses = loadObject("SearchCheckbuttonCharClasses").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonCharClassesStrictMode = loadObject("SearchCheckbuttonCharClassesStrictMode").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonRegex = loadObject("SearchCheckbuttonRegex").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonSplitedAnd = loadObject("SearchCheckbuttonSplitedAnd").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonSplitedNot = loadObject("SearchCheckbuttonSplitedNot").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonSplitedOr = loadObject("SearchCheckbuttonSplitedOr").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonWildCard = loadObject("SearchCheckbuttonWildCard").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonWordAnd = loadObject("SearchCheckbuttonWordAnd").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonWordNot = loadObject("SearchCheckbuttonWordNot").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonWordOr = loadObject("SearchCheckbuttonWordOr").(*gtk.CheckButton)
	mainObjects.SearchComboboxTextAnd = loadObject("SearchComboboxTextAnd").(*gtk.ComboBoxText)
	mainObjects.SearchComboboxTextAnd.SetActive(0) /*MANUAL*/
	mainObjects.SearchComboboxTextDateType = loadObject("SearchComboboxTextDateType").(*gtk.ComboBoxText)
	mainObjects.SearchComboboxTextEntryAnd = loadObject("SearchComboboxTextEntryAnd").(*gtk.Entry)
	mainObjects.SearchComboboxTextEntryNot = loadObject("SearchComboboxTextEntryNot").(*gtk.Entry)
	mainObjects.SearchComboboxTextEntryOr = loadObject("SearchComboboxTextEntryOr").(*gtk.Entry)
	mainObjects.SearchComboboxTextNot = loadObject("SearchComboboxTextNot").(*gtk.ComboBoxText)
	mainObjects.SearchComboboxTextNot.SetActive(0) /*MANUAL*/
	mainObjects.SearchComboboxTextOr = loadObject("SearchComboboxTextOr").(*gtk.ComboBoxText)
	mainObjects.SearchComboboxTextOr.SetActive(0) /*MANUAL*/
	mainObjects.SearchComboboxTextType = loadObject("SearchComboboxTextType").(*gtk.ComboBoxText)
	mainObjects.SearchComboboxTextType.SetActive(1) /*MANUAL*/
	mainObjects.SearchFilechooserbutton = loadObject("SearchFilechooserbutton").(*gtk.FileChooserButton)
	mainObjects.SearchSpinbuttonDepth = loadObject("SearchSpinbuttonDepth").(*gtk.SpinButton)
	ad, err := gtk.AdjustmentNew(-1, -1, 100, 1, 0, 0)                 /*MANUAL*/
	Check(err, "Error on:", "SearchSpinbuttonDepth", "Initialisation") /*MANUAL*/
	mainObjects.SearchSpinbuttonDepth.Configure(ad, 1, 0)              /*MANUAL*/
	mainObjects.SearchTreeview = loadObject("SearchTreeview").(*gtk.TreeView)
	mainObjects.Statusbar = loadObject("Statusbar").(*gtk.Statusbar)
	mainObjects.TimeButtonOkNewer = loadObject("TimeButtonOkNewer").(*gtk.Button)
	mainObjects.TimeButtonOkOlder = loadObject("TimeButtonOkOlder").(*gtk.Button)
	mainObjects.TimeButtonResetNewer = loadObject("TimeButtonResetNewer").(*gtk.Button)
	mainObjects.TimeButtonResetOlder = loadObject("TimeButtonResetOlder").(*gtk.Button)
	mainObjects.TimeCalendarNewer = loadObject("TimeCalendarNewer").(*gtk.Calendar)
	mainObjects.TimeCalendarOlder = loadObject("TimeCalendarOlder").(*gtk.Calendar)
	mainObjects.TimeImageTopNewer = loadObject("TimeImageTopNewer").(*gtk.Image)
	mainObjects.TimeImageTopOlder = loadObject("TimeImageTopOlder").(*gtk.Image)
	mainObjects.TimeSpinbuttonHourNewer = loadObject("TimeSpinbuttonHourNewer").(*gtk.SpinButton)
	ad, err = gtk.AdjustmentNew(0, 0, 23, 1, 0, 0)                       /*MANUAL*/
	Check(err, "Error on:", "TimeSpinbuttonHourNewer", "Initialisation") /*MANUAL*/
	mainObjects.TimeSpinbuttonHourNewer.Configure(ad, 1, 0)              /*MANUAL*/
	mainObjects.TimeSpinbuttonHourOlder = loadObject("TimeSpinbuttonHourOlder").(*gtk.SpinButton)
	ad, err = gtk.AdjustmentNew(0, 0, 59, 1, 0, 0)                       /*MANUAL*/
	Check(err, "Error on:", "TimeSpinbuttonHourOlder", "Initialisation") /*MANUAL*/
	mainObjects.TimeSpinbuttonHourOlder.Configure(ad, 1, 0)              /*MANUAL*/
	mainObjects.TimeSpinbuttonMinutsNewer = loadObject("TimeSpinbuttonMinutsNewer").(*gtk.SpinButton)
	ad, err = gtk.AdjustmentNew(0, 0, 59, 1, 0, 0)                         /*MANUAL*/
	Check(err, "Error on:", "TimeSpinbuttonMinutsNewer", "Initialisation") /*MANUAL*/
	mainObjects.TimeSpinbuttonMinutsNewer.Configure(ad, 1, 0)              /*MANUAL*/
	mainObjects.TimeSpinbuttonMinutsOlder = loadObject("TimeSpinbuttonMinutsOlder").(*gtk.SpinButton)
	ad, err = gtk.AdjustmentNew(0, 0, 23, 1, 0, 0)                         /*MANUAL*/
	Check(err, "Error on:", "TimeSpinbuttonMinutsOlder", "Initialisation") /*MANUAL*/
	mainObjects.TimeSpinbuttonMinutsOlder.Configure(ad, 1, 0)              /*MANUAL*/
	mainObjects.TimeSpinbuttonSecondsNewer = loadObject("TimeSpinbuttonSecondsNewer").(*gtk.SpinButton)
	ad, err = gtk.AdjustmentNew(0, 0, 59, 1, 0, 0)                          /*MANUAL*/
	Check(err, "Error on:", "TimeSpinbuttonSecondsNewer", "Initialisation") /*MANUAL*/
	mainObjects.TimeSpinbuttonSecondsNewer.Configure(ad, 1, 0)              /*MANUAL*/
	mainObjects.TimeSpinbuttonSecondsOlder = loadObject("TimeSpinbuttonSecondsOlder").(*gtk.SpinButton)
	ad, err = gtk.AdjustmentNew(0, 0, 59, 1, 0, 0)                          /*MANUAL*/
	Check(err, "Error on:", "TimeSpinbuttonSecondsOlder", "Initialisation") /*MANUAL*/
	mainObjects.TimeSpinbuttonSecondsOlder.Configure(ad, 1, 0)              /*MANUAL*/
	mainObjects.TimeWindowNewer = loadObject("TimeWindowNewer").(*gtk.Window)
	mainObjects.TimeWindowOlder = loadObject("TimeWindowOlder").(*gtk.Window)
	mainObjects.TopImageEventbox = loadObject("TopImageEventbox").(*gtk.EventBox)
}
