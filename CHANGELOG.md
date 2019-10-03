# Changelog

## SearchEngine ©2018-19 H.F.M

All notable changes to this project will be documented in this file.

## [1.8.5] 2019-10-03

### Added

- Option to allow the analysis of symlinked directories.

### Fixed

- And operand: now correctly manages the functionality of the splitted words.
- Solving issues on errors occuring sometimes on opening text file with "GDK_FATAL_ERROR" or while opening directory that getting permission error with a"GVFS-WARNING"

### Changed

- some parts of the search function have been rewritten to avoid some minor issues.
- Display current path in the title bar rather than the tatusbar.
- Lot of parts of the code have been rewritten for more stability and a really faster processing. And, resulting a lower weight executable file.