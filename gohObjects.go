// gohObjects.go

/*
	Source file auto-generated on Fri, 02 Apr 2021 13:10:55 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2018-21 H.F.M - Search Engine v1.9 github.com/hfmrow/search-engine
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
	SearchComboboxTextDateZone             *gtk.ComboBoxText
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
	mainObjects.SearchComboboxTextDateZone = loadObject("SearchComboboxTextDateZone").(*gtk.ComboBoxText)
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
	mainObjects.TopImageEventbox = loadObject("TopImageEventbox").(*gtk.EventBox)
}
