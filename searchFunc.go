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

	glfsff "github.com/hfmrow/genLib/files/findFiles"
	glss "github.com/hfmrow/genLib/slices"
	gltsbh "github.com/hfmrow/genLib/tools/bench"

	"gopkg.in/djherbis/times.v1"
)

/* Scan files and display it */
func fillListstore() (err error) {
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
					if mainOptions.foundFilesList, err = find.FindDepth(mainOptions.LastDirectory,
						mainOptions.Depth, false, mainObjects.SearchCheckbuttonFollowSL.GetActive()); err != nil {
						return err
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
	return err
}

// Get text from gtk.comboBoxTextEntry object
func getEntryText(cbxEntry *gtk.ComboBoxText) string {
	return cbxEntry.GetActiveText()
}

/* Fill / Clean comboBoxText */
func fillComboboxText(cbxt *gtk.ComboBoxText, text string, removAll ...bool) {
	if len(removAll) == 0 {
		cbxt.PrependText(text)
	} else if removAll[0] {
		cbxt.RemoveAll()
	}
}

/* Fill All comboBoxText */
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

/* Build words to fit search structure */
func makeWords(find *glfsff.Search, line string, wWord, op string) {
	words := strings.Split(line, " ")
	for _, word := range words {
		find.SearchAdd(word, wWord, op)
	}
}

/* search job */
func doDisplay() {
	var err error
	var dispTime, extDir string
	var stats os.FileInfo
	var mTime, aTime time.Time
	// var isText bool

	tvs.StoreDetach()
	tvs.ListStore.Clear()

	/* Store and display found files */
	if len(mainOptions.foundFilesList) != 0 {
		for _, file := range mainOptions.foundFilesList {
			extDir = ""
			if stats, err = os.Lstat(file); err == nil && !os.IsNotExist(err) {
				switch {
				case stats.IsDir():
					extDir = "Dir"
				case stats.Mode()&os.ModeSymlink != 0:
					extDir = "Link"
				default:
					// if isText, _, err = g.IsTextFileSimple(file, -1, 0.6, 97); err == nil && isText {
					// 	extDir = filepath.Ext(file) + "/txt"
					// } else {
					if ext := filepath.Ext(file); len(ext) > 1 {
						extDir = ext[1:]
					}
					// }
				}

				mTime = times.Get(stats).ModTime()
				aTime = times.Get(stats).AccessTime()
				switch mainObjects.SearchComboboxTextDateType.GetActive() {
				case 0:
					dispTime = aTime.String()[:16]
				case 1:
					dispTime = mTime.String()[:16]
				case 2:
					dispTime = humanize.Time(aTime)
				case 3:
					dispTime = humanize.Time(mTime)
				}
				tvs.AddRow(nil, tvs.ColValuesStringSliceToIfaceSlice(
					filepath.Base(file),
					humanize.Bytes(uint64(stats.Size())),
					// fmt.Sprintf("%d", stats.Size()),
					extDir,
					dispTime,
					filepath.Dir(file)))
			}
		}
	}
	tvs.StoreAttach()

}

/*	Get entry from controls */
func computeEntry(find *glfsff.Search) bool {

	var wWordAnd, wWordOr, wWordNot, entryAnd, entryOr, entryNot string

	entryAnd = getEntryText(mainObjects.SearchComboboxTextAnd)
	entryOr = getEntryText(mainObjects.SearchComboboxTextOr)
	entryNot = getEntryText(mainObjects.SearchComboboxTextNot)

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

/*	Get datas from controls */
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

	if mainOptions.searchNewerThan.Ready {
		find.SearchTime.SetNewerThan(mainOptions.searchNewerThan.Ready,
			int(mainOptions.searchNewerThan.d),
			int(mainOptions.searchNewerThan.m),
			int(mainOptions.searchNewerThan.y),
			int(mainOptions.searchNewerThan.H),
			int(mainOptions.searchNewerThan.M),
			int(mainOptions.searchNewerThan.S))
	}

	if mainOptions.searchOlderThan.Ready {
		find.SearchTime.SetOlderThan(mainOptions.searchOlderThan.Ready,
			int(mainOptions.searchOlderThan.d),
			int(mainOptions.searchOlderThan.m),
			int(mainOptions.searchOlderThan.y),
			int(mainOptions.searchOlderThan.H),
			int(mainOptions.searchOlderThan.M),
			int(mainOptions.searchOlderThan.S))
	}

}
