// gohImages.go

/*
	Source file auto-generated on Fri, 02 Apr 2021 13:10:55 using Gotk3 Objects Handler v1.7.5 ©2018-21 hfmrow
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2018-21 H.F.M - Search Engine v1.9 github.com/hfmrow/search-engine
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

/**********************************************************/
/* This section preserve user modifications on update.   */
/* Images declarations, used to initialize objects with */
/* The SetPict() func, accept both kind of variables:  */
/* filename or []byte content in case of using        */
/* embedded binary data. The variables names are the */
/* same. "assetsDeclarationsUseEmbedded(bool)" func */
/* could be used to toggle between filenames and   */
/* embedded binary type. See SetPict()            */
/* declaration to learn more on how to use it.   */
/************************************************/
func assignImages() {
	size := 18
	SetPict(mainObjects.ImageTop, searchEngineTop550x48)
	SetPict(mainObjects.MainWindow, find48)
	SetPict(mainObjects.SearchButton, searchFolder48, size)
	SetPict(mainObjects.SearchButtonExit, logout48, size-2)
	SetPict(mainObjects.SearchButtonExport, floppySave48, size-4)
	SetPict(mainObjects.SearchButtonNewerThan, calendarPers48, size)
	SetPict(mainObjects.SearchButtonOlderThan, calendarPers48, size)
	SetPict(mainObjects.SearchButtonResetComboEntry, clearHist48, size-2)
}

/**********************************************************/
/* This section is rewritten on assets update.           */
/* Assets var declarations, this step permit to make a  */
/* bridge between the differents types used, string or */
/* []byte, and to simply switch from one to another.  */
/*****************************************************/
var mainGlade interface{}             // assets/glade/main.glade
var calendarPers48 interface{}        // assets/images/calendar-pers-48.png
var clearHist48 interface{}           // assets/images/clear-hist-48.png
var copyDocument20 interface{}        // assets/images/Copy-document-20.png
var crossIcon48 interface{}           // assets/images/Cross-icon-48.png
var find48 interface{}                // assets/images/find-48.png
var floppySave48 interface{}          // assets/images/floppy-save-48.png
var folder48 interface{}              // assets/images/folder-48.png
var globalNetwork20 interface{}       // assets/images/Global-Network-20.png
var logout48 interface{}              // assets/images/logout-48.png
var play20 interface{}                // assets/images/Play-20.png
var reset48 interface{}               // assets/images/reset-48.png
var searchEngineTop370x32 interface{} // assets/images/search-engine-top-370x32.png
var searchEngineTop550x48 interface{} // assets/images/search-engine-top-550x48.png
var searchFolder48 interface{}        // assets/images/search-folder-48.png
var stop48 interface{}                // assets/images/Stop-48.png
var tickIcon48 interface{}            // assets/images/Tick-icon-48.png
