// gohStartGtk.go

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
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

/*******************************/
/* Gtk3 Window Initialisation */
/*****************************/
func mainStartGtk(winTitle string, width, height int, center bool) {
	mainObjects = new(MainControlsObj)
	gtk.Init(nil)
	if err := newBuilder(mainGlade); err == nil {

		/* Init tempDir and plan to delete it when leaving. */
		if doTempDir {
			tempDir = tempMake(Name)
			defer os.RemoveAll(tempDir)
		}

		/* Parse Gtk objects */
		gladeObjParser()

		/* Fill control with images */
		assignImages()

		/* Set Window Properties */
		if center {
			mainObjects.MainWindow.SetPosition(gtk.WIN_POS_CENTER)
		}
		mainObjects.MainWindow.SetTitle(winTitle)
		mainObjects.MainWindow.SetDefaultSize(width, height)
		mainObjects.MainWindow.Connect("delete-event", windowDestroy)
		mainObjects.MainWindow.ShowAll()

		/* Start main application before signals init... */
		mainApplication()

		/* Objects Signals initialisations */
		signalsPropHandler()

		/* Execute after signals initialisation */
		afterSignals()

		/* Start Gui loop */
		gtk.Main()
	} else {
		log.Fatal("Builder initialisation error.", err.Error())
	}
}

// windowDestroy: on closing/destroying the gui window.
func windowDestroy() {
	if onShutdown() {
		gtk.MainQuit()
	}
}
