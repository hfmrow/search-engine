// colorChooser.go

package gtk3Import

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

/************************************
*  ColorChooserDialog implementation.
 ************************************/

// Convert gdk.RGBA to Hex string value: "#EF2929".
func ColGdkRGBA2Hex(value *gdk.RGBA) string {
	// Convert int string value to Hex with 2 digits.
	var build2DigitsHex = func(intValueStr string) string {
		xi, _ := strconv.Atoi(intValueStr)
		xs := fmt.Sprintf("%X", xi)
		if len(xs) == 1 {
			xs = "0" + xs
		}
		return xs
	}

	regOne := regexp.MustCompile(`[()]`)
	regTwo := regexp.MustCompile(`[,]`)
	tmpStrSl := regOne.Split(value.String(), -1)
	tmpStrSl = regTwo.Split(tmpStrSl[1], -1)
	rr := build2DigitsHex(tmpStrSl[0])
	gg := build2DigitsHex(tmpStrSl[1])
	bb := build2DigitsHex(tmpStrSl[2])
	aa := "FF"
	if len(tmpStrSl) > 3 {
		aa = build2DigitsHex(tmpStrSl[3])
	}
	return fmt.Sprintf("#%s%s%s%s", rr, gg, bb, aa)
}

// ColorChooserDialogAndGetHexValue: Open color chooser dialog and retrieve value. String return is
// converted like this: "rgb(239,41,41)" to "#EF2929".
func ColorChooserDialogAndGetFloat64Val(parentWindow *gtk.Window, title string, orgCol []float64) (outFloat64 []float64, outStr string) {
	// Build Color chooser dialog en give it some basic parameters.
	if ColorChooserDialog, err := gtk.ColorChooserDialogNew(title, parentWindow); err == nil {
		if len(orgCol) != 0 {
			RGBA := gdk.NewRGBA(orgCol...)
			fmt.Println(RGBA.String())
			ColorChooserDialog.ColorChooser.SetRGBA(RGBA)
		}
		ColorChooserDialog.SetSkipTaskbarHint(true)
		ColorChooserDialog.SetKeepAbove(true)
		ColorChooserDialog.SetSizeRequest(10, 10)
		ColorChooserDialog.SetResizable(false)
		ColorChooserDialog.SetModal(true)
		ColorChooserDialog.ColorChooser.SetUseAlpha(true)
		switch ColorChooserDialog.Run() {
		case gtk.RESPONSE_OK:
			ColorChooserDialog.Close()
			return ColorChooserDialog.ColorChooser.GetRGBA().Floats(), ColGdkRGBA2Hex(ColorChooserDialog.ColorChooser.GetRGBA())
		case gtk.RESPONSE_CANCEL:
			ColorChooserDialog.Close()
		}
	}
	return outFloat64, outStr
}
