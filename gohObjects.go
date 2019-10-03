// gohObjects.go

// Source file auto-generated on Wed, 02 Oct 2019 23:28:15 using Gotk3ObjHandler v1.3.8 ©2018-19 H.F.M

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
	SearchCheckbuttonFollowSL              *gtk.CheckButton
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
	SearchTreeviewSelection                *gtk.TreeSelection
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
	mainObjects.SearchCheckbuttonFollowSL = loadObject("SearchCheckbuttonFollowSL").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonRegex = loadObject("SearchCheckbuttonRegex").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonSplitedAnd = loadObject("SearchCheckbuttonSplitedAnd").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonSplitedNot = loadObject("SearchCheckbuttonSplitedNot").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonSplitedOr = loadObject("SearchCheckbuttonSplitedOr").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonWildCard = loadObject("SearchCheckbuttonWildCard").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonWordAnd = loadObject("SearchCheckbuttonWordAnd").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonWordNot = loadObject("SearchCheckbuttonWordNot").(*gtk.CheckButton)
	mainObjects.SearchCheckbuttonWordOr = loadObject("SearchCheckbuttonWordOr").(*gtk.CheckButton)
	mainObjects.SearchComboboxTextAnd = loadObject("SearchComboboxTextAnd").(*gtk.ComboBoxText)
	mainObjects.SearchComboboxTextDateType = loadObject("SearchComboboxTextDateType").(*gtk.ComboBoxText)
	mainObjects.SearchComboboxTextEntryAnd = loadObject("SearchComboboxTextEntryAnd").(*gtk.Entry)
	mainObjects.SearchComboboxTextEntryNot = loadObject("SearchComboboxTextEntryNot").(*gtk.Entry)
	mainObjects.SearchComboboxTextEntryOr = loadObject("SearchComboboxTextEntryOr").(*gtk.Entry)
	mainObjects.SearchComboboxTextNot = loadObject("SearchComboboxTextNot").(*gtk.ComboBoxText)
	mainObjects.SearchComboboxTextOr = loadObject("SearchComboboxTextOr").(*gtk.ComboBoxText)
	mainObjects.SearchComboboxTextType = loadObject("SearchComboboxTextType").(*gtk.ComboBoxText)
	mainObjects.SearchFilechooserbutton = loadObject("SearchFilechooserbutton").(*gtk.FileChooserButton)
	mainObjects.SearchSpinbuttonDepth = loadObject("SearchSpinbuttonDepth").(*gtk.SpinButton)
	mainObjects.SearchTreeview = loadObject("SearchTreeview").(*gtk.TreeView)
	mainObjects.SearchTreeviewSelection = loadObject("SearchTreeviewSelection").(*gtk.TreeSelection)
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
	mainObjects.TimeSpinbuttonHourOlder = loadObject("TimeSpinbuttonHourOlder").(*gtk.SpinButton)
	mainObjects.TimeSpinbuttonMinutsNewer = loadObject("TimeSpinbuttonMinutsNewer").(*gtk.SpinButton)
	mainObjects.TimeSpinbuttonMinutsOlder = loadObject("TimeSpinbuttonMinutsOlder").(*gtk.SpinButton)
	mainObjects.TimeSpinbuttonSecondsNewer = loadObject("TimeSpinbuttonSecondsNewer").(*gtk.SpinButton)
	mainObjects.TimeSpinbuttonSecondsOlder = loadObject("TimeSpinbuttonSecondsOlder").(*gtk.SpinButton)
	mainObjects.TimeWindowNewer = loadObject("TimeWindowNewer").(*gtk.Window)
	mainObjects.TimeWindowOlder = loadObject("TimeWindowOlder").(*gtk.Window)
	mainObjects.TopImageEventbox = loadObject("TopImageEventbox").(*gtk.EventBox)

}



/*************************************/
/* GtkObjects Widget naming. Usualy */
/* used in css to identify objects */
/**********************************/
func widgetNaming() {
	mainObjects.ImageTop.SetName("ImageTop")
	mainObjects.MainWindow.SetName("MainWindow")
	mainObjects.SearchButton.SetName("SearchButton")
	mainObjects.SearchButtonExit.SetName("SearchButtonExit")
	mainObjects.SearchButtonExport.SetName("SearchButtonExport")
	mainObjects.SearchButtonNewerThan.SetName("SearchButtonNewerThan")
	mainObjects.SearchButtonOlderThan.SetName("SearchButtonOlderThan")
	mainObjects.SearchButtonResetComboEntry.SetName("SearchButtonResetComboEntry")
	mainObjects.SearchCheckbuttonCaseSensitive.SetName("SearchCheckbuttonCaseSensitive")
	mainObjects.SearchCheckbuttonCharClasses.SetName("SearchCheckbuttonCharClasses")
	mainObjects.SearchCheckbuttonCharClassesStrictMode.SetName("SearchCheckbuttonCharClassesStrictMode")
	mainObjects.SearchCheckbuttonFollowSL.SetName("SearchCheckbuttonFollowSL")
	mainObjects.SearchCheckbuttonRegex.SetName("SearchCheckbuttonRegex")
	mainObjects.SearchCheckbuttonSplitedAnd.SetName("SearchCheckbuttonSplitedAnd")
	mainObjects.SearchCheckbuttonSplitedNot.SetName("SearchCheckbuttonSplitedNot")
	mainObjects.SearchCheckbuttonSplitedOr.SetName("SearchCheckbuttonSplitedOr")
	mainObjects.SearchCheckbuttonWildCard.SetName("SearchCheckbuttonWildCard")
	mainObjects.SearchCheckbuttonWordAnd.SetName("SearchCheckbuttonWordAnd")
	mainObjects.SearchCheckbuttonWordNot.SetName("SearchCheckbuttonWordNot")
	mainObjects.SearchCheckbuttonWordOr.SetName("SearchCheckbuttonWordOr")
	mainObjects.SearchComboboxTextAnd.SetName("SearchComboboxTextAnd")
	mainObjects.SearchComboboxTextDateType.SetName("SearchComboboxTextDateType")
	mainObjects.SearchComboboxTextEntryAnd.SetName("SearchComboboxTextEntryAnd")
	mainObjects.SearchComboboxTextEntryNot.SetName("SearchComboboxTextEntryNot")
	mainObjects.SearchComboboxTextEntryOr.SetName("SearchComboboxTextEntryOr")
	mainObjects.SearchComboboxTextNot.SetName("SearchComboboxTextNot")
	mainObjects.SearchComboboxTextOr.SetName("SearchComboboxTextOr")
	mainObjects.SearchComboboxTextType.SetName("SearchComboboxTextType")
	mainObjects.SearchFilechooserbutton.SetName("SearchFilechooserbutton")
	mainObjects.SearchSpinbuttonDepth.SetName("SearchSpinbuttonDepth")
	mainObjects.SearchTreeview.SetName("SearchTreeview")
	mainObjects.Statusbar.SetName("Statusbar")
	mainObjects.TimeButtonOkNewer.SetName("TimeButtonOkNewer")
	mainObjects.TimeButtonOkOlder.SetName("TimeButtonOkOlder")
	mainObjects.TimeButtonResetNewer.SetName("TimeButtonResetNewer")
	mainObjects.TimeButtonResetOlder.SetName("TimeButtonResetOlder")
	mainObjects.TimeCalendarNewer.SetName("TimeCalendarNewer")
	mainObjects.TimeCalendarOlder.SetName("TimeCalendarOlder")
	mainObjects.TimeImageTopNewer.SetName("TimeImageTopNewer")
	mainObjects.TimeImageTopOlder.SetName("TimeImageTopOlder")
	mainObjects.TimeSpinbuttonHourNewer.SetName("TimeSpinbuttonHourNewer")
	mainObjects.TimeSpinbuttonHourOlder.SetName("TimeSpinbuttonHourOlder")
	mainObjects.TimeSpinbuttonMinutsNewer.SetName("TimeSpinbuttonMinutsNewer")
	mainObjects.TimeSpinbuttonMinutsOlder.SetName("TimeSpinbuttonMinutsOlder")
	mainObjects.TimeSpinbuttonSecondsNewer.SetName("TimeSpinbuttonSecondsNewer")
	mainObjects.TimeSpinbuttonSecondsOlder.SetName("TimeSpinbuttonSecondsOlder")
	mainObjects.TimeWindowNewer.SetName("TimeWindowNewer")
	mainObjects.TimeWindowOlder.SetName("TimeWindowOlder")
	mainObjects.TopImageEventbox.SetName("TopImageEventbox")

}
