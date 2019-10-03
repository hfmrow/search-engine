// gohSignals.go

// Source file auto-generated on Wed, 02 Oct 2019 23:28:15 using Gotk3ObjHandler v1.3.8 ©2018-19 H.F.M

/*
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
	mainObjects.SearchCheckbuttonFollowSL.Connect("notify", blankNotify)
	mainObjects.SearchCheckbuttonRegex.Connect("toggled", SearchCheckbuttonRegexToggled)
	mainObjects.SearchComboboxTextDateType.Connect("changed", SearchComboboxTextDateTypeChanged)
	mainObjects.SearchComboboxTextEntryAnd.Connect("activate", SearchButtonClicked)
	mainObjects.SearchComboboxTextEntryNot.Connect("activate", SearchButtonClicked)
	mainObjects.SearchComboboxTextEntryOr.Connect("activate", SearchButtonClicked)
	mainObjects.SearchComboboxTextType.Connect("changed", SearchComboboxTextTypeChanged)
	mainObjects.SearchFilechooserbutton.Connect("file-set", SearchFilechooserbuttonFileSet)
	mainObjects.SearchTreeview.Connect("button-press-event", SearchTreeviewButtonPressEvent)
	mainObjects.SearchTreeviewSelection.Connect("changed", updateStatusbar)
	mainObjects.TimeButtonOkNewer.Connect("clicked", TimeButtonOkNewerClicked)
	mainObjects.TimeButtonOkOlder.Connect("clicked", TimeButtonOkOlderClicked)
	mainObjects.TimeButtonResetNewer.Connect("clicked", TimeButtonResetNewerClicked)
	mainObjects.TimeButtonResetOlder.Connect("clicked", TimeButtonResetOlderClicked)
	mainObjects.TimeCalendarNewer.Connect("day-selected-double-click", TimeButtonOkNewerClicked)
	mainObjects.TimeCalendarOlder.Connect("day-selected-double-click", TimeButtonOkOlderClicked)
	mainObjects.TopImageEventbox.Connect("button-release-event", TopImageEventboxClicked)

}
