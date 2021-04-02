// gohSignals.go

/*
	Source file auto-generated on Fri, 02 Apr 2021 13:10:55 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2018-21 H.F.M - Search Engine v1.9 github.com/hfmrow/search-engine
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

/***************************/
/* Signals Implementations */
/***************************/
// signalsPropHandler: initialise signals used by gtk objects ...
func signalsPropHandler() {
	mainObjects.SearchButton.Connect("clicked", SearchButtonClicked)
	mainObjects.SearchButtonExit.Connect("clicked", windowDestroy)
	mainObjects.SearchButtonExport.Connect("clicked", SearchButtonExportClicked)
	mainObjects.SearchButtonNewerThan.Connect("clicked", SearchButtonNewerThanClicked)
	mainObjects.SearchButtonOlderThan.Connect("clicked", SearchButtonOlderThanClicked)
	mainObjects.SearchButtonResetComboEntry.Connect("clicked", SearchButtonResetComboEntryClicked)
	mainObjects.SearchCheckbuttonCharClasses.Connect("clicked", SearchCheckbuttonCharClassesClicked)
	mainObjects.SearchCheckbuttonRegex.Connect("toggled", SearchCheckbuttonRegexToggled)
	mainObjects.SearchComboboxTextDateType.Connect("changed", ComboboxTextGenericChanged)
	mainObjects.SearchComboboxTextDateZone.Connect("changed", ComboboxTextGenericChanged)
	mainObjects.SearchComboboxTextEntryAnd.Connect("activate", SearchButtonClicked)
	mainObjects.SearchComboboxTextEntryNot.Connect("activate", SearchButtonClicked)
	mainObjects.SearchComboboxTextEntryOr.Connect("activate", SearchButtonClicked)
	mainObjects.SearchComboboxTextType.Connect("changed", SearchComboboxTextTypeChanged)
	mainObjects.SearchFilechooserbutton.Connect("file-set", SearchFilechooserbuttonFileSet)
	mainObjects.SearchTreeview.Connect("button-press-event", popupMenu.CheckRMBFromTreeView)
	mainObjects.TopImageEventbox.Connect("button-release-event", TopImageEventboxClicked)
}
