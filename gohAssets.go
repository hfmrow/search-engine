// gohAssets.go

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
	"embed"
	"log"
)

//go:embed assets/glade
//go:embed assets/images
var embeddedFiles embed.FS

// This functionality does not require explicit encoding of the files, at each
// compilation, the files are inserted into the resulting binary. Thus, updating
// assets is only required when new files are added to be embedded in order to
// create and declare the variables to which the files are linked.
// assetsDeclarationsUseEmbedded: Use native Go 'embed' package to include files
// content at runtime.
func assetsDeclarationsUseEmbedded(embedded ...bool) {
	mainGlade = readEmbedFile("assets/glade/main.glade")
	calendarPers48 = readEmbedFile("assets/images/calendar-pers-48.png")
	clearHist48 = readEmbedFile("assets/images/clear-hist-48.png")
	copyDocument20 = readEmbedFile("assets/images/Copy-document-20.png")
	crossIcon48 = readEmbedFile("assets/images/Cross-icon-48.png")
	find48 = readEmbedFile("assets/images/find-48.png")
	floppySave48 = readEmbedFile("assets/images/floppy-save-48.png")
	folder48 = readEmbedFile("assets/images/folder-48.png")
	globalNetwork20 = readEmbedFile("assets/images/Global-Network-20.png")
	logout48 = readEmbedFile("assets/images/logout-48.png")
	play20 = readEmbedFile("assets/images/Play-20.png")
	reset48 = readEmbedFile("assets/images/reset-48.png")
	searchEngineTop370x32 = readEmbedFile("assets/images/search-engine-top-370x32.png")
	searchEngineTop550x48 = readEmbedFile("assets/images/search-engine-top-550x48.png")
	searchFolder48 = readEmbedFile("assets/images/search-folder-48.png")
	stop48 = readEmbedFile("assets/images/Stop-48.png")
	tickIcon48 = readEmbedFile("assets/images/Tick-icon-48.png")
}

// readEmbedFile: read 'embed' file system and return []byte data.
func readEmbedFile(filename string) (out []byte) {
	var err error
	out, err = embeddedFiles.ReadFile(filename)
	if err != nil {
		log.Printf("Unable to read embedded file: %s, %v\n", filename, err)
	}
	return
}
