// searchFunc.go

/*
*	©2019 H.F.M. MIT license
 */

package main

import (
	"fmt"
	"strings"

	"github.com/gotk3/gotk3/gtk"
	g "github.com/hfmrow/searchEngine/genLib"
	gi "github.com/hfmrow/searchEngine/gtk3Import"
)

/* Scan files and display it */
func fillListstore() (err error) {
	find := g.SearchNew()

	statusbar.CleanAll()
	statusbar.Set("Ready to search.", 0)

	if computeEntry(find) {
		/* Running search job */
		find.SearchCompile()
		if g.ScanFile(mainOptions.LastDirectory).IsExists && find.Ready {

			timer.Lapse("Searching")
			mainOptions.foundFilesList, err = find.FindDepth(mainOptions.LastDirectory, mainOptions.Depth)
			timer.Lapse("Display")
			if err != nil {
				return err
			}
			/* Show results */
			doDisplay()
			timer.Stop()

			statusbar.Set(fmt.Sprintf("%d files scanned", find.BrowsedFiles), 0)
			statusbar.Set(fmt.Sprintf("%d", len(mainOptions.foundFilesList)), 1)
			statusbar.Set(g.TruncatePath(mainOptions.LastDirectory, 3), 2)
			statusbar.Set(timer.Results[0], 3)
			statusbar.Set(timer.Results[1], 4)
		}
		mainOptions.Write()
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
func makeWords(find *g.Search, line string, wWord, op string) {
	words := strings.Split(line, " ")
	for _, word := range words {
		find.SearchAdd(word, wWord, op)
	}
}

/* search job */
func doDisplay() {
	var dispTime string
	mainObjects.fileListstore.Clear()
	/* Store and display found files */
	if len(mainOptions.foundFilesList) != 0 {
		for _, file := range mainOptions.foundFilesList {
			scanned := g.ScanFile(file)
			if scanned.IsExists {
				switch mainObjects.SearchComboboxTextDateType.GetActive() {
				case 0:
					dispTime = scanned.AtimeYMDhms
				case 1:
					dispTime = scanned.MtimeYMDhms
				case 2:
					dispTime = scanned.AtimeFriendlyHR
				case 3:
					dispTime = scanned.MtimeFriendlyHR
				}
				gi.ListStoreAddRow(mainObjects.fileListstore,
					[]string{scanned.Base,
						scanned.SizeHR,
						scanned.Ext,
						dispTime,
						scanned.Path})
			}
		}
	}
}

/*	Get entry from controls */
func computeEntry(find *g.Search) bool {

	var wWordAnd, wWordOr, wWordNot, entryAnd, entryOr, entryNot string

	entryAnd = getEntryText(mainObjects.SearchComboboxTextAnd)
	entryOr = getEntryText(mainObjects.SearchComboboxTextOr)
	entryNot = getEntryText(mainObjects.SearchComboboxTextNot)

	/* Check if there is some entry */
	if len(entryAnd+entryOr+entryNot) == 0 {
		return false
	}

	/* Record entry */
	if (len(entryAnd) != 0) && (g.GetStrIndex(mainOptions.SearchList.And, entryAnd) == -1) {
		mainOptions.SearchList.And = append(mainOptions.SearchList.And, entryAnd)
		fillComboboxText(mainObjects.SearchComboboxTextAnd, entryAnd)
	}

	if (len(entryOr) != 0) && (g.GetStrIndex(mainOptions.SearchList.Or, entryOr) == -1) {
		mainOptions.SearchList.Or = append(mainOptions.SearchList.Or, entryOr)
		fillComboboxText(mainObjects.SearchComboboxTextOr, entryOr)
	}

	if (len(entryNot) != 0) && (g.GetStrIndex(mainOptions.SearchList.Not, entryNot) == -1) {
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
func computeOptions(find *g.Search) {

	mainOptions.LastDirectory = mainObjects.SearchFilechooserbutton.GetFilename()
	/* prevent the use of zero as depth */
	if mainObjects.SearchSpinbuttonDepth.GetValue() == 0 {
		mainObjects.SearchSpinbuttonDepth.SetValue(-1)
	}
	mainOptions.Depth = int(mainObjects.SearchSpinbuttonDepth.GetValue())

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
