// translate.go

// File generated on Mon, 11 Mar 2019 10:27:21 using Gotk3ObjectsTranslate v1.0 2019 H.F.M

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
	"path/filepath"

	"github.com/gotk3/gotk3/gtk"
)

// initGtkObjectsText: read translations from structure and set them to objects.
func (trans *MainTranslate) initGtkObjectsText() {
	trans.setTextToGtkObjects(&mainObjects.SearchButtonResetComboEntry.Widget, "SearchButtonResetComboEntry")
	trans.setTextToGtkObjects(&mainObjects.SearchButtonNewerThan.Widget, "SearchButtonNewerThan")
	trans.setTextToGtkObjects(&mainObjects.SearchButtonOlderThan.Widget, "SearchButtonOlderThan")
	trans.setTextToGtkObjects(&mainObjects.SearchButtonExport.Widget, "SearchButtonExport")
	trans.setTextToGtkObjects(&mainObjects.SearchButtonExit.Widget, "SearchButtonExit")
	trans.setTextToGtkObjects(&mainObjects.SearchButton.Widget, "SearchButton")
	trans.setTextToGtkObjects(&mainObjects.TimeButtonResetNewer.Widget, "TimeButtonResetNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeButtonOkNewer.Widget, "TimeButtonOkNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeButtonResetOlder.Widget, "TimeButtonResetOlder")
	trans.setTextToGtkObjects(&mainObjects.TimeButtonOkOlder.Widget, "TimeButtonOkOlder")
	trans.setTextToGtkObjects(&mainObjects.SearchSpinbuttonDepth.Widget, "SearchSpinbuttonDepth")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonHourNewer.Widget, "TimeSpinbuttonHourNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonMinutsNewer.Widget, "TimeSpinbuttonMinutsNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonSecondsNewer.Widget, "TimeSpinbuttonSecondsNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonHourOlder.Widget, "TimeSpinbuttonHourOlder")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonMinutsOlder.Widget, "TimeSpinbuttonMinutsOlder")
	trans.setTextToGtkObjects(&mainObjects.TimeSpinbuttonSecondsOlder.Widget, "TimeSpinbuttonSecondsOlder")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextEntryAnd.Widget, "SearchComboboxTextEntryAnd")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextEntryOr.Widget, "SearchComboboxTextEntryOr")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextEntryNot.Widget, "SearchComboboxTextEntryNot")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonSplitedAnd.Widget, "SearchCheckbuttonSplitedAnd")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonWordAnd.Widget, "SearchCheckbuttonWordAnd")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonSplitedOr.Widget, "SearchCheckbuttonSplitedOr")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonWordOr.Widget, "SearchCheckbuttonWordOr")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonSplitedNot.Widget, "SearchCheckbuttonSplitedNot")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonWordNot.Widget, "SearchCheckbuttonWordNot")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonCaseSensitive.Widget, "SearchCheckbuttonCaseSensitive")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonWildCard.Widget, "SearchCheckbuttonWildCard")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonRegex.Widget, "SearchCheckbuttonRegex")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonCharClasses.Widget, "SearchCheckbuttonCharClasses")
	trans.setTextToGtkObjects(&mainObjects.SearchCheckbuttonCharClassesStrictMode.Widget, "SearchCheckbuttonCharClassesStrictMode")
	trans.setTextToGtkObjects(&mainObjects.ImageTop.Widget, "ImageTop")
	trans.setTextToGtkObjects(&mainObjects.TimeImageTopNewer.Widget, "TimeImageTopNewer")
	trans.setTextToGtkObjects(&mainObjects.TimeImageTopOlder.Widget, "TimeImageTopOlder")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextAnd.Widget, "SearchComboboxTextAnd")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextOr.Widget, "SearchComboboxTextOr")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextNot.Widget, "SearchComboboxTextNot")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextType.Widget, "SearchComboboxTextType")
	trans.setTextToGtkObjects(&mainObjects.SearchComboboxTextDateType.Widget, "SearchComboboxTextDateType")
	trans.setTextToGtkObjects(&mainObjects.SearchTreeview.Widget, "SearchTreeview")
	trans.setTextToGtkObjects(&mainObjects.SearchFilechooserbutton.Widget, "SearchFilechooserbutton")
}

// sentences: some sentences/words used in the application.
var sentences = map[string]string{}

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
// devModeActive, indicate that the new sentences must be added to original language file
func MainTranslateNew(filename string, devModeActive ...bool) (mt *MainTranslate) {
	mt = new(MainTranslate)
	if _, err := os.Stat(filename); err == nil {
		mt.read(filename)
		mt.initGtkObjectsText()
		if len(devModeActive) != 0 {
			if devModeActive[0] {
				mt.Sentences = sentences
				err := mt.write(filename)
				if err != nil {
					fmt.Printf("%s\n%s\n", "Cannot write actual sentences to language file.", err.Error())
				}
			}
		}
	} else {
		fmt.Printf("%s\n%s\n", "Error loading language file !", err.Error())
	}
	return mt
}

// readFile: language file.
func (trans *MainTranslate) read(filename string) (err error) {
	var textFileBytes []byte
	if textFileBytes, err = ioutil.ReadFile(filename); err == nil {
		if err = json.Unmarshal(textFileBytes, &trans); err == nil {
			trans.ProgInfos.GladeXmlFilenameRel, _ = filepath.Rel(filepath.Dir(filename), trans.ProgInfos.GladeXmlFilename)
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
	DoBackup       bool
}

type progInfo struct {
	Name                string
	Version             string
	Creat               string
	MainObjStructName   string
	GladeXmlFilename    string
	GladeXmlFilenameRel string
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
	Class   string
	Id      string
	Label   string
	Tooltip string
	Text    string
	Uri     string
	Markup  bool
	Comment string
}

// Define available property within objects
type propObject struct {
	Class   string
	Label   bool
	Tooltip bool
	Markup  bool
	Text    bool
	Uri     bool
}

// Property that exists for GObject ...	(Used for Class)
var propPerObjects = []propObject{
	{Class: "GtkButton", Label: true, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkToggleButton", Label: true, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkLabel", Label: true, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkSpinButton", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkEntry", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkCheckButton", Label: true, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkProgressBar", Label: false, Tooltip: true, Markup: true, Text: true, Uri: false},
	{Class: "GtkSearchBar", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkImage", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkRadioButton", Label: true, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkComboBoxText", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkComboBox", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkLinkButton", Label: true, Tooltip: true, Markup: true, Text: false, Uri: true},
	{Class: "GtkSwitch", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkTreeView", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
	{Class: "GtkFileChooserButton", Label: false, Tooltip: true, Markup: true, Text: false, Uri: false},
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
					}
					if props.Tooltip && !props.Markup {
						obj.SetProperty("tooltip_text", currObject.Tooltip)
					}
					if props.Tooltip && props.Markup {
						obj.SetProperty("tooltip_markup", currObject.Tooltip)
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
