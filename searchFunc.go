// searchFunc.go

/*
*	©2019 H.F.M. MIT license
 */

package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gotk3/gotk3/gtk"
	"gopkg.in/djherbis/times.v1"

	glfsff "github.com/hfmrow/genLib/files/findFiles"
	glss "github.com/hfmrow/genLib/slices"
	gltsbh "github.com/hfmrow/genLib/tools/bench"
)

// fillListstore: Scan files and display it
func fillListstore() {
	var err error
	var fileInfo os.FileInfo
	find := glfsff.SearchNew()
	timer = gltsbh.BenchNew(false)
	statusbar.CleanAll()
	statusbar.Set("Ready to search.", 0)

	if computeEntry(find) {
		/* Running search job */
		if err = find.SearchCompile(); err == nil {
			if fileInfo, err = os.Stat(mainOptions.LastDirectory); err == nil {

				if fileInfo.IsDir() && find.Ready {
					timer.Lapse("Searching")
					if storeFoundFiles, err = find.FindDepth(mainOptions.LastDirectory,
						mainOptions.Depth, true, mainObjects.SearchCheckbuttonFollowSL.GetActive()); err != nil {
						DialogMessage(mainObjects.MainWindow, "error", "Error occured during search", "\n\n"+err.Error(), "", sts["ok"])
						return
					}
					timer.Lapse("Display")
					/* Show results */
					doDisplay()

					timer.Stop()
					filesScanned = find.BrowsedFiles
					updateStatusbar()
				}
				mainOptions.Write()
			}
		}
	}
	if err != nil {
		DialogMessage(mainObjects.MainWindow, "error", "Error occured during search", "\n\n"+err.Error(), "", "Ok")
	}
}

// getEntryText:
func getEntryText(cbxEntry *gtk.ComboBoxText) string {
	return cbxEntry.GetActiveText()
}

// fillComboboxText: Fill / Clean comboBoxText
func fillComboboxText(cbxt *gtk.ComboBoxText, text string, removAll ...bool) {
	if len(removAll) == 0 {
		cbxt.PrependText(text)
	} else if removAll[0] {
		cbxt.RemoveAll()
	}
}

// fillAllComboboxText:
func fillAllComboboxText(removAll ...bool) {
	if len(removAll) != 0 {
		fillComboboxText(mainObjects.SearchComboboxTextAnd, "", removAll[0])
		fillComboboxText(mainObjects.SearchComboboxTextOr, "", removAll[0])
		fillComboboxText(mainObjects.SearchComboboxTextNot, "", removAll[0])
		return
	}

	var recurseFill = func(cbxTxt *gtk.ComboBoxText, wordsList []string) {
		for _, word := range wordsList {
			fillComboboxText(cbxTxt, word)
		}
	}

	if len(mainOptions.SearchList.And) != 0 {
		recurseFill(mainObjects.SearchComboboxTextAnd, mainOptions.SearchList.And)
	}
	if len(mainOptions.SearchList.Or) != 0 {
		recurseFill(mainObjects.SearchComboboxTextOr, mainOptions.SearchList.Or)
	}
	if len(mainOptions.SearchList.Not) != 0 {
		recurseFill(mainObjects.SearchComboboxTextNot, mainOptions.SearchList.Not)
	}
}

// makeWords: Build words to fit search structure
func makeWords(find *glfsff.Search, line string, wWord, op string) {
	words := strings.Split(line, " ")
	for _, word := range words {
		find.SearchAdd(word, wWord, op)
	}
}

// doDisplay: display found files
func doDisplay() {

	var (
		dispTime,
		extDir string
		timeUnix int64
		mTime,
		aTime time.Time
	)

	tvs.StoreDetach()
	defer tvs.StoreAttach()
	tvs.ListStore.Clear()

	/* Store and display found files */
	if len(storeFoundFiles) != 0 {
		for _, file := range storeFoundFiles {
			extDir = ""
			switch {
			case file.FileInfo.IsDir():
				extDir = "Dir"
			case file.FileInfo.Mode()&os.ModeSymlink != 0:
				extDir = "Link"
			default:
				if ext := filepath.Ext(file.FileInfo.Name()); len(ext) > 1 {
					extDir = ext[1:]
				}
			}
			if mainObjects.SearchComboboxTextDateZone.GetActive() == 0 {
				mTime = times.Get(file.FileInfo).ModTime().UTC()
				aTime = times.Get(file.FileInfo).AccessTime().UTC()
			} else {
				mTime = times.Get(file.FileInfo).ModTime().Local()
				aTime = times.Get(file.FileInfo).AccessTime().Local()
			}
			switch mainObjects.SearchComboboxTextDateType.GetActive() {
			case 0:
				dispTime = aTime.String()[:16]
				timeUnix = aTime.Unix()
			case 1:
				dispTime = mTime.String()[:16]
				timeUnix = mTime.Unix()
			case 2:
				dispTime = humanize.Time(aTime)
				timeUnix = aTime.Unix()
			case 3:
				dispTime = humanize.Time(mTime)
				timeUnix = mTime.Unix()
			}
			tvs.AddRow(nil,
				file.FileInfo.Name(),
				humanize.Bytes(uint64(file.FileInfo.Size())),
				extDir,
				dispTime,
				filepath.Dir(file.FilePath),
				file.FileInfo.Size(), // invisible
				timeUnix)             // invisible
		}
	}
}

// computeEntry: Get entry from controls
func computeEntry(find *glfsff.Search) bool {

	var (
		wWordAnd,
		wWordOr,
		wWordNot string

		entryAnd = getEntryText(mainObjects.SearchComboboxTextAnd)
		entryOr  = getEntryText(mainObjects.SearchComboboxTextOr)
		entryNot = getEntryText(mainObjects.SearchComboboxTextNot)
	)

	/* Check if there is some entry */
	if len(entryAnd+entryOr+entryNot) == 0 {
		return false
	}

	/* Record entry */
	if (len(entryAnd) != 0) && (glss.GetStrIndex(mainOptions.SearchList.And, entryAnd) == -1) {
		mainOptions.SearchList.And = append(mainOptions.SearchList.And, entryAnd)
		fillComboboxText(mainObjects.SearchComboboxTextAnd, entryAnd)
	}

	if (len(entryOr) != 0) && (glss.GetStrIndex(mainOptions.SearchList.Or, entryOr) == -1) {
		mainOptions.SearchList.Or = append(mainOptions.SearchList.Or, entryOr)
		fillComboboxText(mainObjects.SearchComboboxTextOr, entryOr)
	}

	if (len(entryNot) != 0) && (glss.GetStrIndex(mainOptions.SearchList.Not, entryNot) == -1) {
		mainOptions.SearchList.Not = append(mainOptions.SearchList.Not, entryNot)
		fillComboboxText(mainObjects.SearchComboboxTextNot, entryNot)
	}

	/* Compute entry */
	if mainObjects.SearchCheckbuttonWordAnd.GetActive() {
		wWordAnd = "w"
	}
	if mainObjects.SearchCheckbuttonWordOr.GetActive() {
		wWordOr = "w"
	}
	if mainObjects.SearchCheckbuttonWordNot.GetActive() {
		wWordNot = "w"
	}

	if mainObjects.SearchCheckbuttonSplitedAnd.GetActive() {
		makeWords(find, entryAnd, wWordAnd, "&")
	} else {
		find.SearchAdd(entryAnd, wWordAnd, "&")
	}
	if mainObjects.SearchCheckbuttonSplitedOr.GetActive() {
		makeWords(find, entryOr, wWordOr, "|")
	} else {
		find.SearchAdd(entryOr, wWordOr, "|")
	}
	if mainObjects.SearchCheckbuttonSplitedNot.GetActive() {
		makeWords(find, entryNot, wWordNot, "!")
	} else {
		find.SearchAdd(entryNot, wWordNot, "!")
	}

	computeOptions(find)
	return true
}

// computeOptions: Get options from controls
func computeOptions(find *glfsff.Search) {

	mainOptions.LastDirectory = mainObjects.SearchFilechooserbutton.GetFilename()

	mainOptions.Depth = mainObjects.SearchSpinbuttonDepth.GetValueAsInt()

	find.CaseSensitive = mainObjects.SearchCheckbuttonCaseSensitive.GetActive()
	find.Regex = mainObjects.SearchCheckbuttonRegex.GetActive()
	find.Wildcard = mainObjects.SearchCheckbuttonWildCard.GetActive()
	find.WholeWord = mainObjects.SearchCheckbuttonWordAnd.GetActive()

	switch mainObjects.SearchComboboxTextType.GetActive() {
	case 0:
		find.Type.All()
	case 1:
		find.Type.File()
	case 2:
		find.Type.Dir()
	case 3:
		find.Type.Link()
	default:
	}

	find.POSIXcharClass = mainObjects.SearchCheckbuttonCharClasses.GetActive()
	find.POSIXstrictMode = mainObjects.SearchCheckbuttonCharClassesStrictMode.GetActive()

	// send date/time to search options corresponding to current choosen zone.
	if mainObjects.SearchComboboxTextDateZone.GetActive() == 1 {
		find.SearchTime.SetNewerThan(mainOptions.calDataNewerThan.ToTimeAsUTC())
		find.SearchTime.SetOlderThan(mainOptions.calDataOlderThan.ToTimeAsUTC())
	} else {
		find.SearchTime.SetNewerThan(mainOptions.calDataNewerThan.ToTime())
		find.SearchTime.SetOlderThan(mainOptions.calDataOlderThan.ToTime())
	}
}
