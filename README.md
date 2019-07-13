# Search Engine
*This program is designed to search files over directory, subdirectory, and retrieving information based on date/time, type, patterns contained in name. i'm using it as a replacement for native search function.*

- If you just want to use it, simply download the compiled version under the "release" tab. You can see at [the Right way to install H.F.M's softwares](https://github.com/hfmrow/instHFMsofts) page to integrate this software in your debian environment.
	
- If you want to play inside code, see below "How to compile" section.

## How it's made
- Programed with go language: [golang](https://golang.org/doc/) 
- GUI provided by [Gotk3 (gtk3 v3.22)](https://github.com/gotk3/gotk3), GUI library for Go (minimum required v3.16).
- I use homemade software: "Gotk3ObjHandler" to embed images/icons, UI-information and manage/generate gtk3 objects code from [glade ui designer](https://glade.gnome.org/). and "Gotk3ObjTranslate" to generate the language files and the assignment of a tooltip on the gtk3 objects (both are not published at the moment, in fact, they need documentations and, for the moment, I have not had the time to do them).

## Fonctionnalities
- Search files based on patterns (contained in filename).
- Search files based on file type (all, files, dir, link).
- Search files based on date time (newer than, older than).
- Wildcard, regex, case sensitive, character classes can be used.
- Date/time display can be in human readable or classic format.
- Modification or last acces date time display selection.
- And, or, not as logical operators may be used.
- Whole word functionality.
- Patterns can be splitted (at space char) to add multiples searches for each logical operator.
- Directory depth can be defined.
- History of previous pettern are recorded and can be erased as you whish
- Export function is available to save the results list in txt format.
- Each function have his tooltip for explanations.

## Some pictures and explanations  

*This is the main screen.*  
![Main](readME-Pic/mainScr.png  "Main")  

*Search result.*  
![Result](readME-Pic/searchAction.png  "Result")  

*Search option. Date/Time Older than*  
![Search Date/Time](readME-Pic/olderThan.png  "Search Date/Time")  

*Date/Time display.*  
![Date/Time display](readME-Pic/modAccess.png  "Date/Time display")  

*File type selection.*  
![File type selection](readME-Pic/fileType.png  "File type selection")  

*History display*  
![History display](readME-Pic/histSearch.png  "History display")  

*Tooltip display*  
![Tooltip display](readME-Pic/tooltipDisp.png  "Tooltip display")  

## How to compile
- Be sure you have golang installed in right way. [Go installation](https://golang.org/doc/install).
- Open terminal window and at command prompt, type: `go get github.com/hfmrow/searchengine`
- See [Gotk3 Installation instructions](https://github.com/gotk3/gotk3/wiki#installation) for gui installation instruction.
- To change gtk3 interface you need to use a home made software, (not published actually). So don't change gtk3 interface (glade file) ...
- To change language file you need to use another home made software, (not published actually). So don't change language file ...
- To Produce a standalone executable, you must change inside "main.go" file:

		func main() {
			devMode = true
		...
into

		func main() {
			devMode = false
		...

This operation indicate that externals datas (Image/Icons) must be embedded into the executable file.

### Misc/Os informations
- I'm working on:

| DISTRIB  | LinuxMint |
| -------- | --------- |
| VERSION  | 19.1  |
| CODENAME  | tessa  |
| RELEASE  | #21~18.04.1-Ubuntu SMP Thu Jun 27 04:04:37 UTC 2019  |
| UBUNTU_CODENAME  | bionic  |
| KERNEL  | 5.0.0-20-generic  |
| HDWPLATFORM  | x86_64  |
| GTK  | libgtk-3-0:amd64 3.22.30-1ubuntu3  |
| GLIB  | Ubuntu GLIBC 2.27-3ubuntu1  |

- The compilation have not been tested under Windows or Mac OS, but all file access functions, line-end manipulations or charset implementation are made with OS portability in mind.  

## You got an issue ?
- Give informations (as above), about used plateform and OS version.
- Provide a method to reproduce the problem.
