// translate.go

// File generated on Thu, 03 Oct 2019 01:29:58 using Gotk3ObjectsTranslate v1.3 2019 H.F.M

/*
* 	This program comes with absolutely no warranty.
*	See the The MIT License (MIT) for details:
*	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

// initGtkObjectsText: read translations from structure and set them to objects.
func (trans *MainTranslate) initGtkObjectsText() {
	trans.setTextToGtkObjects(&mainObjects.ImageTop.Widget, "ImageTop")
	trans.setTextToGtkObjects(&mainObjects.SearchButton.Widget, "SearchButton")
	trans.setTextToGtkObjects(&mainObjects.SearchButtonExit.Widget, "SearchButtonExit")
	trans.setTextToGtkObjects(&mainObjects.SearchButtonExport.Widget, "SearchButtonExport")
	trans.setTextToGtkObjects(&mainObjects.SearchButtonNewerThan.Widget, "SearchButtonNewerThan")
	trans.setTextToGtkObjects(&mainObjects.SearchButtonOlderThan.Widget, "SearchButtonOlderThan")
	trans.setTextToGtkObjects(&mainObjects.SearchButtonResetComboEntry.Widget, "SearchButtonResetComboEntry")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonCaseSensitive.Widget, "SearchCheckbuttonCaseSensitive")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonCharClasses.Widget, "SearchCheckbuttonCharClasses")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonCharClassesStrictMode.Widget, "SearchCheckbuttonCharClassesStrictMode")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonFollowSL.Widget, "SearchCheckbuttonFollowSL")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonRegex.Widget, "SearchCheckbuttonRegex")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonSplitedAnd.Widget, "SearchCheckbuttonSplitedAnd")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonSplitedNot.Widget, "SearchCheckbuttonSplitedNot")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonSplitedOr.Widget, "SearchCheckbuttonSplitedOr")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonWildCard.Widget, "SearchCheckbuttonWildCard")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonWordAnd.Widget, "SearchCheckbuttonWordAnd")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonWordNot.Widget, "SearchCheckbuttonWordNot")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonWordOr.Widget, "SearchCheckbuttonWordOr")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextAnd.Widget, "SearchComboboxTextAnd")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextDateType.Widget, "SearchComboboxTextDateType")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextEntryAnd.Widget, "SearchComboboxTextEntryAnd")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextEntryNot.Widget, "SearchComboboxTextEntryNot")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextEntryOr.Widget, "SearchComboboxTextEntryOr")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextNot.Widget, "SearchComboboxTextNot")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextOr.Widget, "SearchComboboxTextOr")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextType.Widget, "SearchComboboxTextType")
	trans.setTextToGtkObjects(&mainObjects.SearchFilechooserbutton.Widget, "SearchFilechooserbutton")
	trans.setTextToGtkObjects(&mainObjects.SearchSpinbuttonDepth.Widget, "SearchSpinbuttonDepth")
	trans.setTextToGtkObjects(&mainObjects.SearchTreeview.Widget, "SearchTreeview")
	trans.setTextToGtkObjects(&mainObjects.TimeButtonOkNewer.Widget, "TimeButtonOkNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeButtonOkOlder.Widget, "TimeButtonOkOlder")
	trans.setTextToGtkObjects(&mainObjects.TimeButtonResetNewer.Widget, "TimeButtonResetNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeButtonResetOlder.Widget, "TimeButtonResetOlder")
	trans.setTextToGtkObjects(&mainObjects.TimeImageTopNewer.Widget, "TimeImageTopNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeImageTopOlder.Widget, "TimeImageTopOlder")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonHourNewer.Widget, "TimeSpinbuttonHourNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonHourOlder.Widget, "TimeSpinbuttonHourOlder")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonMinutsNewer.Widget, "TimeSpinbuttonMinutsNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonMinutsOlder.Widget, "TimeSpinbuttonMinutsOlder")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonSecondsNewer.Widget, "TimeSpinbuttonSecondsNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonSecondsOlder.Widget, "TimeSpinbuttonSecondsOlder")
}
// Translations structure declaration. To be used in main application.
var translate = new(MainTranslate)

// sts: some sentences/words used in the application. Mostly used in Development mode.
// You must add there all sentences used in your application. Or not ...
// They'll be added to language file each time application started
// when "devMode" is set at true.
var sts = map[string]string{
	`no`: `No`,
	`yes`: `Yes`,
}


// Translations structure with methods
type MainTranslate struct {
	// Public
	ProgInfos    progInfo
	Language     language
	Options      parsingFlags
	ObjectsCount int
	Objects      []object
	Sentences    map[string]string
	// Private
	objectsLoaded bool
}

// MainTranslateNew: Initialise new translation structure and assign language file content to GtkObjects.
// devModeActive, indicate that the new sentences must be added to previous language file.
func MainTranslateNew(filename string, devModeActive ...bool) (mt *MainTranslate) {
	mt = new(MainTranslate)
	if _, err := os.Stat(filename); err == nil {
		mt.read(filename)
		mt.initGtkObjectsText()
		if len(devModeActive) != 0 {
			if devModeActive[0] {
				mt.Sentences = sts
				err := mt.write(filename)
				if err != nil {
					fmt.Printf("%s\n%s\n", "Cannot write actual sentences to language file.", err.Error())
				}
			}
		}
	} else {
		fmt.Printf("%s\n%s\n", "Error loading language file !\nNot an error when is just creating from glade Xml or GOH project file.", err.Error())
	}
	return mt
}

// readFile: language file.
func (trans *MainTranslate) read(filename string) (err error) {
	var textFileBytes []byte
	if textFileBytes, err = ioutil.ReadFile(filename); err == nil {
		if err = json.Unmarshal(textFileBytes, &trans); err == nil {
			trans.objectsLoaded = true
		}
	}
	return err
}

// Write json datas to file
func (trans *MainTranslate) write(filename string) (err error) {
	var out bytes.Buffer
	var jsonData []byte
	if jsonData, err = json.Marshal(&trans); err == nil && trans.objectsLoaded {
		if err = json.Indent(&out, jsonData, "", "\t"); err == nil {
			err = ioutil.WriteFile(filename, out.Bytes(), 0644)
		}
	}
	return err
}

type parsingFlags struct {
	SkipLowerCase  bool
	SkipEmptyLabel bool
	SkipEmptyName  bool
	DoBackup       bool
}

type progInfo struct {
	Name              string
	Version           string
	Creat             string
	MainObjStructName string
	GladeXmlFilename  string
	TranslateFilename string
}

type language struct {
	LangNameLong string
	LangNameShrt string
	Author       string
	Date         string
	Updated      string
	Contributors []string
}

type object struct {
	Class         string
	Id            string
	Label         string
	LabelMarkup   bool
	LabelWrap     bool
	Tooltip       string
	TooltipMarkup bool
	Text          string
	Uri           string
	Comment       string
	Idx           int
}

// Define available property within objects
type propObject struct {
	Class         string
	Label         bool
	LabelMarkup   bool
	LabelWrap     bool
	Tooltip       bool
	TooltipMarkup bool
	Text          bool
	Uri           bool
}

// Property that exists for Gtk3 Object ...	(Used for Class capability)
var propPerObjects = []propObject{
	{Class: "GtkButton", Label: true, Tooltip: true, TooltipMarkup: true},
	{Class: "GtkToolButton", Label: true, Tooltip: true, TooltipMarkup: true},
	{Class: "GtkToggleButton", Label: true, Tooltip: true, TooltipMarkup: true},
	{Class: "GtkLabel", Label: true, LabelMarkup: true, Tooltip: true, TooltipMarkup: true, LabelWrap: true},
	{Class: "GtkSpinButton", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkEntry", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkCheckButton", Label: true, Tooltip: true, TooltipMarkup: true},
	{Class: "GtkProgressBar", Tooltip: true, TooltipMarkup: true, Text: true},
	{Class: "GtkSearchBar", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkImage", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkRadioButton", Label: true, LabelMarkup: false, Tooltip: true, TooltipMarkup: true},
	{Class: "GtkComboBoxText", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkComboBox", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkLinkButton", Label: true, Tooltip: true, TooltipMarkup: true, Uri: true},
	{Class: "GtkSwitch", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkTreeView", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkFileChooserButton", Tooltip: true, TooltipMarkup: true},
	{Class: "GtkTextView", Tooltip: true, TooltipMarkup: true},
}

// setTextToGtkObjects: read translations from structure and set them to object.
// like this: setTextToGtkObjects(&mainObjects.TransLabelHint.Widget, "TransLabelHint")
func (trans *MainTranslate) setTextToGtkObjects(obj *gtk.Widget, objectId string) {
	for _, currObject := range trans.Objects {
		if currObject.Id == objectId {
			for _, props := range propPerObjects {
				if currObject.Class == props.Class {
					if props.Label {
						obj.SetProperty("label", currObject.Label)
						if props.LabelMarkup {
							obj.SetProperty("use-markup", currObject.LabelMarkup)
							obj.SetProperty("label", strings.ReplaceAll(currObject.Label, "&", "&amp;"))
						}
					}
					if props.LabelWrap {
						obj.SetProperty("wrap", currObject.LabelWrap)
					}
					if props.Tooltip && !currObject.TooltipMarkup {
						obj.SetProperty("tooltip_text", currObject.Tooltip)
					}
					if props.Tooltip && currObject.TooltipMarkup {
						obj.SetProperty("tooltip_markup", strings.ReplaceAll(currObject.Tooltip, "&", "&amp;"))
					}
					if props.Text {
						obj.SetProperty("text", currObject.Text)
					}
					if props.Uri {
						obj.SetProperty("uri", currObject.Uri)
					}
				}
			}
		}
	}
}
