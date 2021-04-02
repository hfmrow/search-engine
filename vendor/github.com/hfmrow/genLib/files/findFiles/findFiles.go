// findFiles.go

// Modestly similary to the linux find function

package findFiles

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	glcc "github.com/hfmrow/genLib/strings/cClass"
	times "gopkg.in/djherbis/times.v1"
)

type Search struct {
	BrowsedFiles    int
	Ready           bool
	SearchTime      searchTime
	ShowDir         bool
	CaseSensitive   bool
	POSIXcharClass  bool
	POSIXstrictMode bool
	Regex           bool
	UseEscapeChar   bool
	Wildcard        bool
	WholeWord       bool
	Type            searchType

	searchInto string
	searchFor  []word

	readyToCompile bool

	searchForRegexpAnd []*regexp.Regexp
	searchForRegexpOr  []*regexp.Regexp
	searchForRegexpNot []*regexp.Regexp
}

type StoreFiles struct {
	FilePath string
	FileInfo os.FileInfo
}

type searchTime struct {
	newerThan time.Time
	olderThan time.Time
	access    bool
	modif     bool
	ntReady   bool
	otReady   bool
}

func (st *searchTime) AccessTime() {
	st.access = true
	st.modif = false
}

func (st *searchTime) ModifTime() {
	st.access = false
	st.modif = true
}

func (st *searchTime) SetNewerThan(inTime time.Time) {
	var blankTime time.Time
	if inTime != blankTime {
		st.ntReady = true
		st.newerThan = inTime
	} else {
		st.ntReady = false
	}
}

func (st *searchTime) SetOlderThan(inTime time.Time) {
	var blankTime time.Time
	if inTime != blankTime {
		st.otReady = true
		st.olderThan = inTime
	} else {
		st.otReady = false
		st.olderThan = time.Now()
	}
}

type searchType struct {
	typeAll  bool
	typeFile bool
	typeDir  bool
	typeLink bool
}

func (t *searchType) All() {
	t.typeReset()
	t.typeAll = true
}
func (t *searchType) File() {
	t.typeReset()
	t.typeFile = true
}
func (t *searchType) Dir() {
	t.typeReset()
	t.typeDir = true
}
func (t *searchType) Link() {
	t.typeReset()
	t.typeLink = true
}
func (t *searchType) typeReset() {
	t.typeAll = false
	t.typeFile = false
	t.typeDir = false
	t.typeLink = false
}

type word struct {
	Word      string
	WholeWord string
	Logical   string /* "&",(And) "|",(Or) "!",(Not) */
}

// SearchCompile: create regex string and compile it..
func (s *Search) SearchCompile() (err error) {
	var regexStr string
	var tmpRegexp *regexp.Regexp

	if s.readyToCompile {
		s.Ready = true
		if s.Regex {
			if tmpRegexp, err = regexp.Compile(s.searchFor[0].Word); err != nil {
				return
			}
			s.searchForRegexpAnd = append(s.searchForRegexpAnd, tmpRegexp)
		} else {
			for _, values := range s.searchFor {
				regexStr = values.Word

				switch {
				case s.POSIXcharClass:
					regexStr = glcc.StringToCharacterClasses(regexStr, s.CaseSensitive, s.POSIXstrictMode)
				case s.Wildcard:
					if !s.UseEscapeChar {
						regexStr = strings.Replace(regexStr, "?", "¤¤¤¤¤¤", -1)
						regexStr = strings.Replace(regexStr, "*", "¤¤¤¤¤", -1)
						regexStr = regexp.QuoteMeta(regexStr)
						regexStr = strings.Replace(regexStr, "¤¤¤¤¤¤", "?", -1)
						regexStr = strings.Replace(regexStr, "¤¤¤¤¤", "*", -1)
					}
					regexStr = strings.Replace(regexStr, "*", `.+`, -1)
					regexStr = strings.Replace(regexStr, "?", `.`, -1)
				case !s.UseEscapeChar:
					regexStr = regexp.QuoteMeta(regexStr)
				}

				if values.WholeWord == "w" {
					regexStr = `\b` + regexStr + `\b`
				}
				regexStr = "(" + regexStr + ")"
				if !s.CaseSensitive {
					regexStr = "(?i)" + regexStr
				}

				regexStr += ".*"

				tmpRegexp = regexp.MustCompile(regexStr)

				switch values.Logical {
				case "&":
					s.searchForRegexpAnd = append(s.searchForRegexpAnd, tmpRegexp)
				case "|":
					s.searchForRegexpOr = append(s.searchForRegexpOr, tmpRegexp)
				case "!":
					s.searchForRegexpNot = append(s.searchForRegexpNot, tmpRegexp)
				}
				regexStr = ""
			}
			if len(regexStr) != 0 {
				s.searchForRegexpAnd = append(s.searchForRegexpAnd, tmpRegexp)
			}
		}
	}
	return
}

func SearchNew() *Search {
	s := new(Search)
	s.Type.All()
	s.CaseSensitive = true
	s.SearchTime.ModifTime()
	return s
}

// SearchAdd: Adding entry to be searched. Format: "wordToFind", "w", "&"
// "w", "", means WholeWord or not. "&", "|", "!", means and, or, not
func (s *Search) SearchAdd(searchFor, wWord, logicalOp string) {
	if len(searchFor) != 0 {
		s.readyToCompile = true
		s.searchFor = append(s.searchFor, word{searchFor, wWord, logicalOp})
	}
}

func (s *Search) searchInAdd(searchIn string) {
	s.searchInto = searchIn
}

// SIMLINK version ***********
// FindDepth: Search for file in dir and subdir depending on depth argument. depth = -1 means infinite.
// This function get parameter from a Search structure which contain all search options.
// Note: time comparisons are calculated only with UTC format by golang ...
func (toFind *Search) FindDepth(root string, depth int, showDir, followSymlink bool) (files []StoreFiles, err error) {
	if files, err = toFind.passFindDepth(root, depth, showDir, followSymlink); err == nil {
		var ok, dir, link bool
		if !toFind.Type.typeAll {
			for idx := len(files) - 1; idx >= 0; idx-- {
				dir = files[idx].FileInfo.Mode()&os.ModeDir != 0
				link = files[idx].FileInfo.Mode()&os.ModeSymlink != 0
				ok = false
				switch {
				case toFind.Type.typeDir && dir:
					ok = showDir
				case toFind.Type.typeLink && link:
					ok = true
				case toFind.Type.typeFile && (!dir && !link):
					ok = true
				}
				if !ok { // Remove entry that does not match
					files = append(files[:idx], files[idx+1:]...)
				}
			}
		}
		// Note, time comparisons are calculated only with UTC format by golang...
		if toFind.SearchTime.ntReady || toFind.SearchTime.otReady {
			var infos times.Timespec
			var timeCheck time.Time
			if toFind.SearchTime.olderThan == timeCheck {
				toFind.SearchTime.olderThan = time.Now().UTC()
			}
			for idx := len(files) - 1; idx >= 0; idx-- {
				infos = times.Get(files[idx].FileInfo)
				// Select modification or access
				if toFind.SearchTime.access {
					timeCheck = infos.AccessTime()
					// fmt.Printf("A-%s: %v < %v < %v\n", files[idx].FileInfo.Name(), toFind.SearchTime.newerThan, timeCheck, toFind.SearchTime.olderThan)
				} else {
					timeCheck = infos.ModTime()
					// fmt.Printf("M-%s: %v < %v < %v\n", files[idx].FileInfo.Name(), toFind.SearchTime.newerThan, timeCheck, toFind.SearchTime.olderThan)
				} // Compare and  Remove entry that does not match
				if !(timeCheck.After(toFind.SearchTime.newerThan) && timeCheck.Before(toFind.SearchTime.olderThan)) {
					files = append(files[:idx], files[idx+1:]...)
				}
			}
			return
		}
	}
	return
}

// passFindDepth: this is the first pass for ScanDir ethos
func (toFind *Search) passFindDepth(root string, depth int, showDir, followSymlink bool) (files []StoreFiles, err error) {
	var targetSL string
	var depthRecurse int
	var osFile *os.File
	var osFileInfos []os.FileInfo
	var statSL os.FileInfo
	var tmpFilesRecurse []StoreFiles

	if toFind.Ready {
		// root = strings.TrimSuffix(root, string(filepath.Separator))
		if osFile, err = os.Open(root); err == nil {
			defer osFile.Close()
			if osFileInfos, err = osFile.Readdir(-1); err != nil {
				return files, err
			}
			toFind.BrowsedFiles += len(osFileInfos)
			for _, file := range osFileInfos {

				depthRecurse = depth
				if file.IsDir() {
					if depth != 0 {
						depthRecurse--
						tmpFilesRecurse, err = toFind.passFindDepth(filepath.Join(root, file.Name()), depthRecurse, showDir, followSymlink)
						if err != nil && !os.IsPermission(err) {
							return files, err
						}
					}
					files = append(files, tmpFilesRecurse...)
				} else if followSymlink && file.Mode()&os.ModeSymlink != 0 {
					if targetSL, err = os.Readlink(filepath.Join(root, file.Name())); err == nil {
						if statSL, err = os.Lstat(targetSL); err == nil {
							if statSL.IsDir() {
								if depth != 0 {
									depthRecurse--
									tmpFilesRecurse, err = toFind.passFindDepth(filepath.Join(root, file.Name()), depthRecurse, showDir, followSymlink)
									if err != nil && !os.IsPermission(err) {
										return files, err
									}
								}
								files = append(files, tmpFilesRecurse...)
							}
						}
					} else if err != nil && !os.IsPermission(err) {
						return files, err
					}
				}
				toFind.searchInto = file.Name()
				if toFind.search() {
					files = append(files, StoreFiles{
						FilePath: filepath.Join(root, file.Name()),
						FileInfo: file})
				}
			}
			if len(files) == 0 {
				return files, err
			}
		} else {
			err = errors.New(fmt.Sprintf("Nothing to search in ! : %s\n", err.Error()))
		}
	}
	return files, err
}

// search: filter input to keep only that match
func (toFind *Search) search() (resp bool) {
	// And
	for _, elem := range toFind.searchForRegexpAnd {
		if elem.MatchString(toFind.searchInto) {
			resp = true
		} else {
			resp = false
			break
		}
	} // Or
	for _, elem := range toFind.searchForRegexpOr {
		if elem.MatchString(toFind.searchInto) {
			resp = true
			break
		}
	} // Not
	for _, elem := range toFind.searchForRegexpNot {
		if elem.MatchString(toFind.searchInto) {
			resp = false
			break
		}
	}
	return
}

// // Find: Search for file in dir all subdir. Equal performances between this one and the (Depth) version
// // This function get parameter from a Search structure which contain all search options.
// func (toFind *Search) Find(root string) (files []string, err error) {
// 	var tmpFiles []string
// 	var ok, dir, link bool

// 	if toFind.Ready {
// 		err = filepath.Walk(root,
// 			func(path string, info os.FileInfo, err error) error {
// 				toFind.SearchInAdd(info.Name())
// 				if toFind.Type.typeAll {
// 					ok = search(toFind)
// 				} else {
// 					dir = (info.Mode()&os.ModeDir != 0)
// 					link = (info.Mode()&os.ModeSymlink != 0)
// 					switch {
// 					case toFind.Type.typeDir && dir:
// 						ok = search(toFind)
// 					case toFind.Type.typeLink && link:
// 						ok = search(toFind)
// 					case toFind.Type.typeFile && (!dir && !link):
// 						ok = search(toFind)
// 					}
// 				}
// 				toFind.BrowsedFiles++
// 				if ok {
// 					tmpFiles = append(tmpFiles, path)
// 					ok = false
// 				}
// 				return nil
// 			})
// 		if err == nil {

// 			if len(tmpFiles) == 0 {
// 				return files, err
// 			}
// 			files = toFind.searchFiltering(tmpFiles)
// 		}
// 	} else {
// 		err = errors.New("Error: Nothing to search in ...")
// 	}
// 	return files, err
// }

/*************************
*	EXAMPLE how to use.  *
*************************/

/*
package main

import (
	"fmt"

	"github.com/hfmrow/genLib"
)

func main() {
	count := 1
	var files []string
	var err error
	find := genLib.SearchNew()
	genLib.Use(count)
	root := "/home/syndicate/Videos/"
	// toFind2 := "Une"
	//	toFind1 := "extraordinaires"
	toFind := `Blade Runner`

	find.SearchAdd(toFind, "&")
	//	find.SearchAdd(toFind1, "!")
	//	find.SearchAdd(toFind2, "|")

	//	find.POSIXcharClass = true
	//	find.POSIXstrictMode = true
	//	find.Regex = true
	find.Wildcard = true
	//	find.WholeWord = false
	find.CaseSensitive = false
	find.Type.All()
	find.SearchTime.AccessTime()
	find.SearchTime.SetNewerThan(1, 8, 2017, 8, 1, 46)
	find.SearchTime.SetOlderThan(1, 4, 2018, 8, 0, 33)
	find.SearchCompile()

	for idx := 0; idx < count; idx++ {

		//	files = find.Find( find)
		files, err = find.Find(root)
	}
	genLib.Check(err)
	outPutFiles := genLib.ScanFiles(files)
	for _, file := range outPutFiles {
		fmt.Printf("FileType: %s	Name: %v\nFullname: %v\nSize: %v	Modif: %v	Access: %v\n",
			file.Type, file.Base, file.PathBase, file.SizeHR, file.MtimeYMDhmsShort, file.AtimeYMDhmsShort)
	}
	fmt.Printf("\nTotal: %v\n", len(files))
}
*/
