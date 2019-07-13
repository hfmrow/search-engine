// gohImages.go

// Source file auto-generated on Wed, 10 Jul 2019 21:07:55 using Gotk3ObjHandler v1.3 ©2019 H.F.M

/*
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

/************************************************************/
/* Images declarations, used to initialize objects with it */
/* The functions: setImage, setWinIcon and setButtonImage */
/* accept both kind of variables: filename or []byte     */
/* content in case of using embedded binary data. The   */
/* variables names are the same. You can use function  */
/* "func assetsDeclarationsUseEmbedded(bool)"         */
/* to toggle between filenames and embedded binary.  */
/****************************************************/
func assignImages() {
	setImage(mainObjects.ImageTop, searchEngine700x48)
	setWinIcon(mainObjects.MainWindow, find48)
	setButtonImage(mainObjects.SearchButton, find18)
	setButtonImage(mainObjects.SearchButtonExit, cancel18x18)
	setButtonImage(mainObjects.SearchButtonExport, saveAll18)
	setButtonImage(mainObjects.SearchButtonNewerThan, calendar18)
	setButtonImage(mainObjects.SearchButtonOlderThan, calendar18)
	setButtonImage(mainObjects.SearchButtonResetComboEntry, reset18x18)
	setButtonImage(mainObjects.TimeButtonOkNewer, checked18x18)
	setButtonImage(mainObjects.TimeButtonOkOlder, checked18x18)
	setButtonImage(mainObjects.TimeButtonResetNewer, reset18x18)
	setButtonImage(mainObjects.TimeButtonResetOlder, reset18x18)
	setImage(mainObjects.TimeImageTopNewer, searchEngine400x27)
	setImage(mainObjects.TimeImageTopOlder, searchEngine400x27)
	setWinIcon(mainObjects.TimeWindowNewer, "")
	setWinIcon(mainObjects.TimeWindowOlder, "")
}

// Assets var declarations, this step permit to make a "bridge" between the differents
// types used: (string or []byte) and to simply switch from one to another.
var calendar18 interface{}         // assets/images/calendar-18.png
var calendar48 interface{}         // assets/images/calendar-48.png
var cancel18x18 interface{}        // assets/images/cancel-18x18.png
var checked18x18 interface{}       // assets/images/checked-18x18.png
var copyDocument20 interface{}     // assets/images/Copy-document-20.png
var find18 interface{}             // assets/images/find-18.png
var find48 interface{}             // assets/images/find-48.png
var folderOpen20 interface{}       // assets/images/folder-open-20.png
var globalNetwork20 interface{}    // assets/images/Global-Network-20.png
var mainGlade interface{}          // assets/glade/main.glade
var open18 interface{}             // assets/images/open-18.png
var play20 interface{}             // assets/images/Play-20.png
var reset18x18 interface{}         // assets/images/reset-18x18.png
var saveAll18 interface{}          // assets/images/save-all-18.png
var searchEngine400x27 interface{} // assets/images/Search-Engine-400x27.png
var searchEngine700x48 interface{} // assets/images/Search-Engine-700x48.png
var signError20 interface{}        // assets/images/Sign-Error-20.png
