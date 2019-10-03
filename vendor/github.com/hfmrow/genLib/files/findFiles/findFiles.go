// findFiles.go

// Modestly similary to the linux find function

package findFiles

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	glcc "github.com/hfmrow/genLib/strings/cClass"
	glte "github.com/hfmrow/genLib/tools/errors"
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

// Date format is set to: Day, Month, Year -- Hour, Minute, Second.
// Local time is used.
func (st *searchTime) SetNewerThan(ready bool, dateTime ...int) {
	st.newerThan = time.Time{}
	var H, M, S, day, year int
	var Month time.Month
	if ready {
		if len(dateTime) != 0 {
			for idx, value := range dateTime {
				switch idx {
				case 0:
					day = value
				case 1:
					Month = time.Month(value)
				case 2:
					year = value
				case 3:
					H = value
				case 4:
					M = value
				case 5:
					S = value
				}
			}
			st.newerThan = time.Date(year, Month, day, H, M, S, 0, time.Local)
			st.ntReady = ready
		}
	}
	//	fmt.Printf("Newer than: %v\n", st.newerThan)	/*	Control	*/
}

// Date format is set to: Day, Month, Year -- Hour, Minute, Second.
// Local time is used.
func (st *searchTime) SetOlderThan(ready bool, dateTime ...int) {
	st.olderThan = time.Time{}
	var H, M, S, day, year int
	var Month time.Month
	if ready {
		if len(dateTime) != 0 {
			for idx, value := range dateTime {
				switch idx {
				case 0:
					day = value
				case 1:
					Month = time.Month(value)
				case 2:
					year = value
				case 3:
					H = value
				case 4:
					M = value
				case 5:
					S = value
				}
			}
			st.olderThan = time.Date(year, Month, day, H, M, S, 0, time.Local)
			st.otReady = ready
		}
		//	fmt.Printf("Older than: %v\n", st.olderThan)	/*	Control	*/
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

	if !s.SearchTime.otReady {
		s.SearchTime.olderThan = time.Now()
	}

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
				// if tmpRegexp, err = regexp.Compile(regexStr); err != nil {
				// 	return
				// }
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
	// fmt.Printf("And Or: %v\n", strings.Join(s.searchForRegStr, ""))	/*	Control	*/
	// fmt.Printf("Not: %v\n", strings.Join(s.searchForRegStrNot, "|"))	/*	Control	*/
	return
}
func SearchNew() *Search { // TODO comms to remove if there no issue on search function
	s := new(Search)
	s.Type.All()
	s.CaseSensitive = true
	s.SearchTime.ModifTime()
	return s
}

// SearchAdd: Adding entry to be searched. Format: "wordToFind", "w", "&"
// "w", "", mean WholeWord or not. "&", "|", "!", mean and, or, not
func (s *Search) SearchAdd(searchFor, wWord, logicalOp string) {
	if len(searchFor) != 0 {
		s.readyToCompile = true
		s.searchFor = append(s.searchFor, word{searchFor, wWord, logicalOp})
	}
}
func (s *Search) SearchInAdd(searchIn string) {
	s.searchInto = searchIn
}

// FindDepth: Search for file in dir and subdir depending on depth argument. depth = -1 mean infinite.
// This function get parameter from a Search structure which contain all search options.
func (toFind *Search) FindDepthTest(root string, depth int) (files []string, err error) {
	var tmpFilesRecurse, tmpFiles []string
	var depthRecurse int

	if toFind.Ready {
		filesInfos, err := ioutil.ReadDir(root)
		if err != nil {
			return files, err
		}
		for _, file := range filesInfos {
			depthRecurse = depth
			if file.IsDir() {
				if depth != 0 {
					depthRecurse--
					tmpFilesRecurse, err = toFind.FindDepthTest(filepath.Join(root, file.Name()), depthRecurse)
					if err != nil {
						return files, err
					}
				}
				tmpFiles = append(tmpFiles, tmpFilesRecurse...)
			}
			toFind.SearchInAdd(file.Name())
			if search(toFind) {
				tmpFiles = append(tmpFiles, filepath.Join(root, file.Name()))
			}
		}
		if len(tmpFiles) == 0 {
			return files, err
		}
		files = toFind.searchFilteringTest(tmpFiles)

	} else {
		err = errors.New("Error: Nothing to search in ...")
	}
	return files, err
}

// Filtering results with provided options ...
func (toFind *Search) searchFilteringTest(tmpFiles []string) (files []string) {
	var skip, dir, link bool
	for _, file := range tmpFiles {
		fileInfos, err := os.Stat(file)
		glte.Check(err)
		if !toFind.Type.typeAll {
			dir = (fileInfos.Mode()&os.ModeDir != 0)
			link = (fileInfos.Mode()&os.ModeSymlink != 0)
			switch {
			case toFind.Type.typeDir && !dir:
				skip = true
			case toFind.Type.typeLink && !link:
				skip = true
			case toFind.Type.typeFile && !dir && !link:
				skip = true
			}
		}
		if !skip {
			skip = false
			if toFind.SearchTime.ntReady || toFind.SearchTime.otReady {
				for _, file := range tmpFiles {
					fileInfos, err := os.Stat(file)
					glte.Check(err)
					infos := times.Get(fileInfos)
					if toFind.SearchTime.access {
						if infos.AccessTime().After(toFind.SearchTime.newerThan) && infos.AccessTime().Before(toFind.SearchTime.olderThan) {
							files = append(files, file)
						}
					} else {
						if infos.ModTime().After(toFind.SearchTime.newerThan) && !infos.ModTime().After(toFind.SearchTime.olderThan) {
							files = append(files, file)
						}
					}
				}
			} else { /* No Time control. */
				return tmpFiles
			}
			files = append(files, file)
		} else { /* 	Not desired Type. */
			return tmpFiles
		}
	}
	return files
}

/*
// CLASSIC version ***********
// FindDepth: Search for file in dir and subdir depending on depth argument. depth = -1 mean infinite.
// This function get parameter from a Search structure which contain all search options.
func (toFind *Search) FindDepthN(root string, depth int) (files []string, err error) {
	var tmpFilesRecurse, tmpFiles []string
	var ok, dir, link bool
	var depthRecurse int

	if toFind.Ready {
		filesInfos, err := ioutil.ReadDir(root)
		if err != nil && !os.IsPermission(err) {
			return files, err
		}
		toFind.BrowsedFiles += len(filesInfos)
		for _, file := range filesInfos {
			depthRecurse = depth
			if file.IsDir() {
				if depth != 0 {
					depthRecurse--
					tmpFilesRecurse, err = toFind.FindDepth(filepath.Join(root, file.Name()), depthRecurse)
					if err != nil && !os.IsPermission(err) {
						return files, err
					}
				}
				tmpFiles = append(tmpFiles, tmpFilesRecurse...)
			}
			toFind.SearchInAdd(file.Name())
			if toFind.Type.typeAll {
				ok = search(toFind)
			} else {
				dir = (file.Mode()&os.ModeDir != 0)
				link = (file.Mode()&os.ModeSymlink != 0)
				switch {
				case toFind.Type.typeDir && dir:
					ok = search(toFind)
				case toFind.Type.typeLink && link:
					ok = search(toFind)
				case toFind.Type.typeFile && (!dir && !link):
					ok = search(toFind)
				}
			}
			if ok {
				ok = false
				tmpFiles = append(tmpFiles, filepath.Join(root, file.Name()))
			}
		}

		if len(tmpFiles) == 0 {
			return files, err
		}
		files = toFind.searchFiltering(tmpFiles)

	} else {
		err = errors.New("Error: Nothing to search in ...")
	}
	return files, err
}
*/

// SIMLINK version ***********
// FindDepth: Search for file in dir and subdir depending on depth argument. depth = -1 mean infinite.
// This function get parameter from a Search structure which contain all search options.
func (toFind *Search) FindDepth(root string, depth int, showDir, followSymlink bool) (files []string, err error) {
	var tmpFilesRecurse, tmpFiles []string
	var targetSL string
	var ok, dir, link bool
	var depthRecurse int
	var osFile *os.File
	var osFileInfos []os.FileInfo
	var statSL os.FileInfo

	if toFind.Ready {
		root = strings.TrimSuffix(root, string(filepath.Separator))
		if osFile, err = os.Open(root); err == nil {
			defer osFile.Close()
			if osFileInfos, err = osFile.Readdir(-1); err != nil {
				return files, err
			}
			toFind.BrowsedFiles += len(osFileInfos)
			for _, file := range osFileInfos {

				depthRecurse = depth
				if file.IsDir() {
					// if showDir {
					// 	tmpFiles = append(tmpFiles, filepath.Join(root, file.Name()))
					// }
					if depth != 0 {
						depthRecurse--
						tmpFilesRecurse, err = toFind.FindDepth(filepath.Join(root, file.Name()), depthRecurse, showDir, followSymlink)
						if err != nil && !os.IsPermission(err) {
							return files, err
						}
					}
					tmpFiles = append(tmpFiles, tmpFilesRecurse...)
				} else if followSymlink && file.Mode()&os.ModeSymlink != 0 {
					if targetSL, err = os.Readlink(filepath.Join(root, file.Name())); err == nil {
						if statSL, err = os.Lstat(targetSL); err == nil {
							if statSL.IsDir() {
								// if showDir {
								// 	tmpFiles = append(tmpFiles, filepath.Join(root, file.Name()))
								// }
								if depth != 0 {
									depthRecurse--
									tmpFilesRecurse, err = toFind.FindDepth(filepath.Join(root, file.Name()), depthRecurse, showDir, followSymlink)
									if err != nil && !os.IsPermission(err) {
										return files, err
									}
								}
								tmpFiles = append(tmpFiles, tmpFilesRecurse...)
							}
						}
					} else if err != nil && !os.IsPermission(err) {
						return files, err
					}
				}
				toFind.SearchInAdd(file.Name())
				if toFind.Type.typeAll {
					ok = search(toFind)
				} else {
					dir = file.Mode()&os.ModeDir != 0
					link = file.Mode()&os.ModeSymlink != 0
					switch {
					case toFind.Type.typeDir && dir:
						if showDir {
							ok = search(toFind)
						}
					case toFind.Type.typeLink && link:
						ok = search(toFind)
					case toFind.Type.typeFile && (!dir && !link):
						ok = search(toFind)
					}
				}
				if ok {
					ok = false
					tmpFiles = append(tmpFiles, filepath.Join(root, file.Name()))
				}
			}

			if len(tmpFiles) == 0 {
				return files, err
			}
			files = toFind.searchFiltering(tmpFiles)

		} else {
			err = errors.New(fmt.Sprintf("Nothing to search in ! : %s\n", err.Error()))
		}
	}
	return files, err
}

// FindDepth: Search for file in dir and subdir depending on depth argument. depth = -1 mean infinite.
// This function get parameter from a Search structure which contain all search options.
func (toFind *Search) FindDepthN(root string, depth int, showDir, followSymlink bool) (files []string, err error) {
	var tmpFilesRecurse, tmpFiles []string
	var targetSL string
	var ok, dir, link bool
	var depthRecurse int
	var osFile *os.File
	var osFileInfos []os.FileInfo
	var statSL os.FileInfo

	if toFind.Ready {

		root = strings.TrimSuffix(root, string(filepath.Separator))

		if osFile, err = os.Open(root); err == nil {
			defer osFile.Close()
			if osFileInfos, err = osFile.Readdir(-1); err != nil {
				return files, err
			}
			toFind.BrowsedFiles += len(osFileInfos)
			for _, file := range osFileInfos {

				if followSymlink && file.Mode()&os.ModeSymlink != 0 {
					if targetSL, err = os.Readlink(filepath.Join(root, file.Name())); err == nil {
						if statSL, err = os.Lstat(targetSL); err == nil {
							if statSL.IsDir() {
								if showDir {
									tmpFiles = append(tmpFiles, filepath.Join(root, file.Name()))
								}
								if depth != 0 {
									depthRecurse--
									tmpFilesRecurse, err = toFind.FindDepthN(filepath.Join(root, file.Name()), depthRecurse, showDir, followSymlink)
									if err != nil && !os.IsPermission(err) {
										return files, err
									}
								}
								tmpFiles = append(tmpFiles, tmpFilesRecurse...)
							}
						}
					} else if err != nil && !os.IsPermission(err) {
						return files, err
					}
				}
				// fmt.Printf("%d - %s\n", idx, file.Name())
				depthRecurse = depth
				if file.IsDir() {
					if showDir {
						tmpFiles = append(tmpFiles, filepath.Join(root, file.Name()))
					}
					if depth != 0 {
						depthRecurse--
						tmpFilesRecurse, err = toFind.FindDepthN(filepath.Join(root, file.Name()), depthRecurse, showDir, followSymlink)
						if err != nil && !os.IsPermission(err) {
							return files, err
						}
					}
					tmpFiles = append(tmpFiles, tmpFilesRecurse...)
				}
				toFind.SearchInAdd(file.Name())
				if toFind.Type.typeAll {
					ok = search(toFind)
				} else {
					dir = (file.Mode()&os.ModeDir != 0)
					link = (file.Mode()&os.ModeSymlink != 0)
					switch {
					case toFind.Type.typeDir && dir:
						ok = search(toFind)
					case toFind.Type.typeLink && link:
						ok = search(toFind)
					case toFind.Type.typeFile && (!dir && !link):
						ok = search(toFind)
					}
				}
				if ok {
					ok = false
					tmpFiles = append(tmpFiles, filepath.Join(root, file.Name()))
				}
			}

			if len(tmpFiles) == 0 {
				return files, err
			}
			files = toFind.searchFiltering(tmpFiles)

		} else {
			err = errors.New(fmt.Sprintf("Nothing to search in ! : %s\n", err.Error()))
		}
	}
	return files, err
}

// Filtering results with provided options ...
func (toFind *Search) searchFiltering(tmpFiles []string) (files []string) {
	if toFind.SearchTime.ntReady || toFind.SearchTime.otReady {
		for _, file := range tmpFiles {
			fileInfos, err := os.Stat(file)
			if err != nil && !os.IsPermission(err) {
				fmt.Printf("%v\n", err)
				return files
			}
			infos := times.Get(fileInfos)
			if toFind.SearchTime.access {
				if infos.AccessTime().After(toFind.SearchTime.newerThan) && infos.AccessTime().Before(toFind.SearchTime.olderThan) {
					files = append(files, file)
				}
			} else {
				if infos.ModTime().After(toFind.SearchTime.newerThan) && !infos.ModTime().After(toFind.SearchTime.olderThan) {
					files = append(files, file)
				}
			}
		}
		//		fmt.Printf("Time control.	Newer than: %v	Older than: %v\n", toFind.SearchTime.newerThan, toFind.SearchTime.olderThan)	/*	Control	*/
	} else {
		//		fmt.Println("No time control.")	/*	Control	*/
		return tmpFiles
	}
	return files
}

// Find: Search for file in dir all subdir. Equal performances between this one and the (Depth) version
// This function get parameter from a Search structure which contain all search options.
func (toFind *Search) Find(root string) (files []string, err error) {
	var tmpFiles []string
	var ok, dir, link bool

	if toFind.Ready {
		err = filepath.Walk(root,
			func(path string, info os.FileInfo, err error) error {
				toFind.SearchInAdd(info.Name())
				if toFind.Type.typeAll {
					ok = search(toFind)
				} else {
					dir = (info.Mode()&os.ModeDir != 0)
					link = (info.Mode()&os.ModeSymlink != 0)
					switch {
					case toFind.Type.typeDir && dir:
						ok = search(toFind)
					case toFind.Type.typeLink && link:
						ok = search(toFind)
					case toFind.Type.typeFile && (!dir && !link):
						ok = search(toFind)
					}
				}
				toFind.BrowsedFiles++
				if ok {
					tmpFiles = append(tmpFiles, path)
					ok = false
				}
				return nil
			})
		if err == nil {

			if len(tmpFiles) == 0 {
				return files, err
			}
			files = toFind.searchFiltering(tmpFiles)
		}
	} else {
		err = errors.New("Error: Nothing to search in ...")
	}
	return files, err
}

func search(toFind *Search) (resp bool) {
	// And
	for _, elem := range toFind.searchForRegexpAnd {
		if elem.MatchString(toFind.searchInto) {
			resp = true
		} else {
			resp = false
			break
		}
	}
	// Or
	for _, elem := range toFind.searchForRegexpOr {
		if elem.MatchString(toFind.searchInto) {
			resp = true
		}
	}
	// Not
	for _, elem := range toFind.searchForRegexpNot {
		if elem.MatchString(toFind.searchInto) {
			resp = false
		}
	}
	return resp
}

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
