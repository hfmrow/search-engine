// gohStartGtk.go

// Source file auto-generated on Fri, 06 Sep 2019 04:25:58 using Gotk3ObjHandler v1.3.6 ©2019 H.F.M

/*
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

/******************************/
/* Gtk3 Window Initialisation */
/******************************/
func mainStartGtk(winTitle string, width, height int, center bool) {
	mainObjects = new(MainControlsObj)
	gtk.Init(nil)
	if newBuilder(mainGlade) == nil {
		// Init tempDir and Remove it on quit if requested.
		if doTempDir {
			tempDir = tempMake(Name)
			defer os.RemoveAll(tempDir)
		}
		// Parse Gtk objects
		gladeObjParser()
		// Naming widgets as Gtk objects
		if namingWidget {
			widgetNaming()
		}
		// Objects Signals initialisations
		signalsPropHandler()
		/* Fill control with images */
		assignImages()
		// Set Window Properties
		if center {
			mainObjects.MainWindow.SetPosition(gtk.WIN_POS_CENTER)
		}
		mainObjects.MainWindow.SetTitle(winTitle)
		mainObjects.MainWindow.SetDefaultSize(width, height)
		mainObjects.MainWindow.Connect("delete-event", windowDestroy)
		mainObjects.MainWindow.ShowAll()
		// Start main application ...
		mainApplication()
		/* Update gtk conctrols with stored values into mainOptions */
		mainOptions.UpdateObjects()
		//	Start Gui loop
		gtk.Main()
	} else {
		log.Fatal("Builder initialisation error.")
	}
}
