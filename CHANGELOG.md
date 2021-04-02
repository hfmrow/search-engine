## Informations

At the bottom you can find a compiled standalone ".deb" version with its checksum. The ".tar.gz" & ".zip" sources contain a "vendor" directory ensuring you can always compile it even if the official libraries have been changed.

## Changelog

All notable changes to this project will be documented in this file.

### [1.9] 2021-04-02

#### Added

- Add the time zone selection (UTC, Local), which is used in the file view and to search by time function.

- Context menu: adding option to set date/time at newer and older than from selected file.

- Context menu delete file(s): confirmation window display the names and their count.

#### Fixed

- Search choosing "Directory" as type of file, will work correctly now instead of before where directory name was displayed only in the "All" mode.

- Sort by size, works properly and reflects a natural order.

- Sort by date, works correctly even if you have enabled Human readable mode.

- Some strange behavior with context menu (handling  has been rewritten).

#### Changed

- Artwork changed, image / icons.

- Calendar section rewritten.

- Repository name was changed to [https://github.com/hfmrow/search-engine](https://github.com/hfmrow/search-engine) instead of `https://github.com/hfmrow/searchEngine`

---

### [1.8.5] 2019-10-03

#### Added

- Option to allow the analysis of symlinked directories.

#### Fixed

- And operand: now correctly manages the functionality of the splitted words.

- Solving issues on errors occuring sometimes on opening text file with "GDK_FATAL_ERROR" or while opening directory that getting permission error with a"GVFS-WARNING"

#### Changed

- some parts of the search function have been rewritten to avoid some minor issues.

- Display current path in the title bar rather than the tatusbar.

- Lot of parts of the code have been rewritten for more stability and a really faster processing. And, resulting a lower weight executable file.
