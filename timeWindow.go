// timeWindow.go

/*
	©2019 H.F.M. MIT license
*/

package main

import (
	"strings"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

/************************/
/* Handling TIME window */
/************************/
// TimeButtonOkNewerClicked
func TimeButtonOkNewerClicked() {
	getCal(mainObjects.TimeCalendarNewer, &mainOptions.searchNewerThan,
		mainObjects.TimeSpinbuttonHourNewer,
		mainObjects.TimeSpinbuttonMinutsNewer,
		mainObjects.TimeSpinbuttonSecondsNewer)
	mainOptions.searchNewerThan.Ready = true
	genericHideWindow(mainObjects.TimeWindowNewer)
	/* Display dateTime */
	dateTime := strings.Split(time.Date(int(mainOptions.searchNewerThan.y),
		time.Month(mainOptions.searchNewerThan.m),
		int(mainOptions.searchNewerThan.d),
		int(mainOptions.searchNewerThan.H),
		int(mainOptions.searchNewerThan.M),
		int(mainOptions.searchNewerThan.S),
		0, time.Local).String(), " ")[:2]
	//	markupLabel(mainObjects.SearchLabelNewerThan, " "+strings.Join(dateTime, " ")+" ")
	//	origLabelNT, _ = mainObjects.SearchButtonNewerThan.GetLabel()
	mainObjects.SearchButtonNewerThan.SetLabel(" " + strings.Join(dateTime, " ") + " ")
}

// TimeButtonOkOlderClicked
func TimeButtonOkOlderClicked() {
	getCal(mainObjects.TimeCalendarOlder, &mainOptions.searchOlderThan,
		mainObjects.TimeSpinbuttonHourOlder,
		mainObjects.TimeSpinbuttonMinutsOlder,
		mainObjects.TimeSpinbuttonSecondsOlder)
	mainOptions.searchOlderThan.Ready = true
	genericHideWindow(mainObjects.TimeWindowOlder)
	/* Display dateTime */
	dateTime := strings.Split(time.Date(int(mainOptions.searchOlderThan.y),
		time.Month(mainOptions.searchOlderThan.m),
		int(mainOptions.searchOlderThan.d),
		int(mainOptions.searchOlderThan.H),
		int(mainOptions.searchOlderThan.M),
		int(mainOptions.searchOlderThan.S),
		0, time.Local).String(), " ")[:2]
	//		markupLabel(mainObjects.SearchLabelOlderThan, " "+strings.Join(dateTime, " ")+" ")
	//	origLabelOT, _ = mainObjects.SearchButtonOlderThan.GetLabel()
	mainObjects.SearchButtonOlderThan.SetLabel(" " + strings.Join(dateTime, " ") + " ")

}

// TimeButtonResetNewerClicked
func TimeButtonResetNewerClicked() {
	mainOptions.searchNewerThan.Ready = false
	mainObjects.SearchButtonNewerThan.SetLabel(origLabelNT)
	genericHideWindow(mainObjects.TimeWindowNewer)
}

// TimeButtonResetOlderClicked
func TimeButtonResetOlderClicked() {
	mainOptions.searchOlderThan.Ready = false
	mainObjects.SearchButtonOlderThan.SetLabel(origLabelOT)
	genericHideWindow(mainObjects.TimeWindowOlder)
}

// Get datetime values
func getCal(cal *gtk.Calendar, sdt *searchTimeCal, spinH, spinM, spinS *gtk.SpinButton) {
	sdt.y, sdt.m, sdt.d = cal.GetDate()
	sdt.H = spinH.GetValue()
	sdt.M = spinM.GetValue()
	sdt.S = spinS.GetValue()
	sdt.m++
}

// Set datetime values
func setCal(cal *gtk.Calendar, sdt *searchTimeCal, spinH, spinM, spinS *gtk.SpinButton) {
	var err error

	/* Set datetime default values */
	Now := time.Now()
	y := uint(Now.Year())
	m := uint(Now.Month())
	d := uint(Now.Day())
	H := float64(Now.Hour())
	M := float64(Now.Minute())
	S := float64(Now.Second())
	/* Set values from last selection if exist */
	if sdt.Ready {
		y = sdt.y
		m = sdt.m
		d = sdt.d
		H = sdt.H
		M = sdt.M
		S = sdt.S
	}
	/* Set control values */
	if err = cal.SetProperty("day", d); err == nil {
		if err = cal.SetProperty("month", m-1); err == nil {
			if err = cal.SetProperty("year", y); err == nil {
				spinH.SetValue(H)
				spinM.SetValue(M)
				spinS.SetValue(S)
			}
		}
	}
	if err != nil {
		DialogMessage(mainObjects.MainWindow, "error", "Error occured during calendar setup", "\n\n"+err.Error(), "", "Ok")
	}
}

// Display Calendar window
func displayTimeWin(window *gtk.Window, title string) {
	window.SetTitle(title)
	window.SetSkipTaskbarHint(true)
	window.SetKeepAbove(true)
	window.SetSizeRequest(400, 10)
	window.SetResizable(false)
	window.SetModal(true)
	window.Connect("delete_event", genericHideWindow)
	window.ShowAll()
}

// Signal handler delete_event (hidding window)
func genericHideWindow(w *gtk.Window) bool {
	if w.GetVisible() {
		w.Hide()
	}
	return true
}

func TopImageEventboxClicked() {
	mainOptions.AboutOptions.Show() /*	Init Aboutbox	*/
}
