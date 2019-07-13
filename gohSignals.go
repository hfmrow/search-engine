// gohSignals.go

// Source file auto-generated on Wed, 10 Jul 2019 21:07:55 using Gotk3ObjHandler v1.3 ©2019 H.F.M

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
	mainObjects.SearchButton.Connect("clicked", SearchButtonClicked)                                 /*MANUAL*/
	mainObjects.SearchButtonExit.Connect("clicked", windowDestroy)                                   /*MANUAL*/
	mainObjects.SearchButtonExport.Connect("clicked", SearchButtonExportClicked)                     /*MANUAL*/
	mainObjects.SearchButtonNewerThan.Connect("clicked", SearchButtonNewerThanClicked)               /*MANUAL*/
	mainObjects.SearchButtonOlderThan.Connect("clicked", SearchButtonOlderThanClicked)               /*MANUAL*/
	mainObjects.SearchButtonResetComboEntry.Connect("clicked", SearchButtonResetComboEntryClicked)   /*MANUAL*/
	mainObjects.SearchCheckbuttonCharClasses.Connect("clicked", SearchCheckbuttonCharClassesClicked) /*MANUAL*/
	mainObjects.SearchCheckbuttonRegex.Connect("toggled", SearchCheckbuttonRegexToggled)             /*MANUAL*/
	mainObjects.SearchComboboxTextDateType.Connect("changed", SearchComboboxTextDateTypeChanged)     /*MANUAL*/
	mainObjects.SearchComboboxTextEntryAnd.Connect("activate", SearchButtonClicked)                  /*MANUAL*/
	mainObjects.SearchComboboxTextEntryNot.Connect("activate", SearchButtonClicked)                  /*MANUAL*/
	mainObjects.SearchComboboxTextEntryOr.Connect("activate", SearchButtonClicked)                   /*MANUAL*/
	mainObjects.SearchComboboxTextType.Connect("changed", SearchComboboxTextTypeChanged)             /*MANUAL*/
	mainObjects.SearchTreeview.Connect("button-press-event", SearchTreeviewButtonPressEvent)         /*MANUAL*/
	mainObjects.TimeButtonOkNewer.Connect("clicked", TimeButtonOkNewerClicked)                       /*MANUAL*/
	mainObjects.TimeButtonOkOlder.Connect("clicked", TimeButtonOkOlderClicked)                       /*MANUAL*/
	mainObjects.TimeButtonResetNewer.Connect("clicked", TimeButtonResetNewerClicked)                 /*MANUAL*/
	mainObjects.TimeButtonResetOlder.Connect("clicked", TimeButtonResetOlderClicked)                 /*MANUAL*/
	mainObjects.TimeCalendarNewer.Connect("day-selected-double-click", TimeButtonOkNewerClicked)     /*MANUAL*/
	mainObjects.TimeCalendarOlder.Connect("day-selected-double-click", TimeButtonOkOlderClicked)     /*MANUAL*/
	mainObjects.TopImageEventbox.Connect("button-release-event", TopImageEventboxClicked)            /*MANUAL*/
}
