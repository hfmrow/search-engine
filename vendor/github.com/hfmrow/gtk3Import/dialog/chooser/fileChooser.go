// fileChooser.go

package gtk3Import

import (
	"fmt"
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
	"save":          gtk.FILE_CHOOSER_ACTION_SAVE,
}

// FileChooser: Display a file chooser dialog.
// dlgType: "open", "save", "create-folder", "select-folder" as dlgType.
// title = "": auto choice based on dialog type.
// options: 1-keepAbove, 2-enablePreviewImages, 3-setModal, 4-askOverwrite
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
		case "open":
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
								fmt.Printf("Image '%s' cannot be loaded, got error: %s", fc.GetFilename(), err.Error())
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

	switch int(fileChooser.Run()) {
	case -3:
		result = true
		outFilename = fileChooser.GetFilename()
	}

	fileChooser.Destroy()
	return
}
