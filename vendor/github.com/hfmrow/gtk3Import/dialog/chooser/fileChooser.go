// fileChooser.go

package gtk3Import

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

/****************************
* FileChooser implementation.
 ****************************/
var FileChooserAction = map[string]gtk.FileChooserAction{
	"select-folder": gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER,
	"create-folder": gtk.FILE_CHOOSER_ACTION_CREATE_FOLDER,
	"open":          gtk.FILE_CHOOSER_ACTION_OPEN,
	"open-entry":    gtk.FILE_CHOOSER_ACTION_SAVE, // Open with 'entry' like in 'save' dialog
	"save":          gtk.FILE_CHOOSER_ACTION_SAVE,
}

var UserAbortError = errors.New("Action aborted by user !")

// FileChooser: Display a file chooser dialog.
// dlgType: "open", "save", "create-folder", "select-folder" as dlgType.
// title = "": auto choice based on dialog type.
// options: 1-keepAbove, 2-enablePreviewImages, 3-setModal, 4-askOverwrite
// Default:	1- true, 2- false, 3- true, 4- true
func FileChooser(window *gtk.Window, dlgType, title, filename string, options ...bool) (outFilename string, result bool, err error) {
	var preview, folder bool
	var fileChooser *gtk.FileChooserDialog
	kpAbove, modal, overwrt := true, true, true
	firstBtn, scndBtn := "Cancel", "Ok"

	switch len(options) {
	case 1:
		kpAbove = options[0]
	case 2:
		kpAbove = options[0]
		preview = options[1]
	case 3:
		kpAbove = options[0]
		preview = options[1]
		modal = options[2]
	case 4:
		kpAbove = options[0]
		preview = options[1]
		modal = options[2]
		overwrt = options[3]
	}

	if len(title) == 0 {
		switch dlgType {

		case "create-folder":
			title = "Create folder"
			folder = true

		case "select-folder":
			title = "Select directory"
			folder = true

		case "open", "open-entry":
			title = "Select file to open"

		case "save":
			title = "Select file to save"
		}
	}

	if fileChooser, err = gtk.FileChooserDialogNewWith2Buttons(title, window, FileChooserAction[dlgType],
		firstBtn, gtk.RESPONSE_CANCEL, scndBtn, gtk.RESPONSE_ACCEPT); err != nil {
		return
	}

	if preview {
		if previewImage, err := gtk.ImageNew(); err == nil {
			previewImage.Show()
			var pixbuf *gdk.Pixbuf
			fileChooser.SetPreviewWidget(previewImage)
			fileChooser.Connect("update-preview", func(fc *gtk.FileChooserDialog) {
				if _, err = os.Stat(fc.GetFilename()); !os.IsNotExist(err) {
					if pixbuf, err = gdk.PixbufNewFromFile(fc.GetFilename()); err == nil {
						fileChooser.SetPreviewWidgetActive(true)
						if pixbuf.GetWidth() > 640 || pixbuf.GetHeight() > 480 {
							if pixbuf, err = gdk.PixbufNewFromFileAtScale(fc.GetFilename(), 200, 200, true); err != nil {
								log.Fatalf("Image '%s' cannot be loaded, got error: %s", fc.GetFilename(), err.Error())
							}
						}
						previewImage.SetFromPixbuf(pixbuf)
					} else {
						fileChooser.SetPreviewWidgetActive(false)
					}
				}
			})
		}
	}

	fileChooser.SetFilename(filename)

	if dlgType == "save" {
		fileChooser.SetCurrentName(filepath.Base(filename))
	}

	if folder {
		fileChooser.SetCurrentFolder(filename)
	} else {
		fileChooser.SetCurrentFolder(filepath.Dir(filename))
	}
	fileChooser.SetDoOverwriteConfirmation(overwrt)
	fileChooser.SetModal(modal)
	fileChooser.SetSkipPagerHint(true)
	fileChooser.SetSkipTaskbarHint(true)
	fileChooser.SetKeepAbove(kpAbove)

	if dlgType == "open-entry" {
		fileChooser.SetDoOverwriteConfirmation(false)
		fileChooser.SetCurrentName(filepath.Base(filename))
	}

	// fmt.Println(gtk.RESPONSE_ACCEPT)       //-3
	// fmt.Println(gtk.RESPONSE_DELETE_EVENT) //-4
	// fmt.Println(gtk.RESPONSE_OK)           //-5
	// fmt.Println(gtk.RESPONSE_CANCEL)       //-6
	// fmt.Println(gtk.RESPONSE_CLOSE)        //-7
	// fmt.Println(gtk.RESPONSE_YES)          //-8

	resp := fileChooser.Run()

	switch resp {
	case gtk.RESPONSE_CANCEL, gtk.RESPONSE_DELETE_EVENT:

		result = false
		err = UserAbortError
	case gtk.RESPONSE_ACCEPT:

		result = true
		outFilename = fileChooser.GetFilename()

		if dlgType == "open-entry" {
			if _, err = os.Stat(outFilename); err != nil {
				result = false
			}
		}
	}

	fileChooser.Destroy()
	return
}
