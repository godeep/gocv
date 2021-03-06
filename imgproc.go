package gocv

/*
#include <stdlib.h>
#include "imgproc.h"
*/
import "C"
import (
	"image"
	"image/color"
	"unsafe"
)

// CvtColor converts an image from one color space to another.
// It converts the src Mat image to the dst Mat using the
// code param containing the desired ColorConversionCode color space.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d7/d1b/group__imgproc__misc.html#ga4e0972be5de079fed4e3a10e24ef5ef0
//
func CvtColor(src Mat, dst Mat, code ColorConversionCode) {
	C.CvtColor(src.p, dst.p, C.int(code))
}

// Blur blurs an image Mat using a box filter.
// The function convolves the src Mat image into the dst Mat using
// the specified Gaussian kernel params.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d4/d86/group__imgproc__filter.html#gaabe8c836e97159a9193fb0b11ac52cf1
//
func Blur(src Mat, dst Mat, ksize image.Point) {
	pSize := C.struct_Size{
		height: C.int(ksize.X),
		width:  C.int(ksize.Y),
	}

	C.Blur(src.p, dst.p, pSize)
}

// GaussianBlur blurs an image Mat using a Gaussian filter.
// The function convolves the src Mat image into the dst Mat using
// the specified Gaussian kernel params.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d4/d86/group__imgproc__filter.html#gaabe8c836e97159a9193fb0b11ac52cf1
//
func GaussianBlur(src Mat, dst Mat, ksize image.Point, sigmaX float64,
	sigmaY float64, borderType int) {
	pSize := C.struct_Size{
		height: C.int(ksize.X),
		width:  C.int(ksize.Y),
	}

	C.GaussianBlur(src.p, dst.p, pSize, C.double(sigmaX), C.double(sigmaY), C.int(borderType))
}

// Canny finds edges in an image using the Canny algorithm.
// The function finds edges in the input image image and marks
// them in the output map edges using the Canny algorithm.
// The smallest value between threshold1 and threshold2 is used
// for edge linking. The largest value is used to
// find initial segments of strong edges.
// See http://en.wikipedia.org/wiki/Canny_edge_detector
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/dd/d1a/group__imgproc__feature.html#ga04723e007ed888ddf11d9ba04e2232de
//
func Canny(src Mat, edges Mat, t1 float32, t2 float32) {
	C.Canny(src.p, edges.p, C.double(t1), C.double(t2))
}

// HoughLines implements the standard or standard multi-scale Hough transform
// algorithm for line detection. For a good explanation of Hough transform, see:
// http://homepages.inf.ed.ac.uk/rbf/HIPR2/hough.htm
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/dd/d1a/group__imgproc__feature.html#ga46b4e588934f6c8dfd509cc6e0e4545a
//
func HoughLines(src Mat, lines Mat, rho float32, theta float32, threshold int) {
	C.HoughLines(src.p, lines.p, C.double(rho), C.double(theta), C.int(threshold))
}

// HoughLinesP implements the probabilistic Hough transform
// algorithm for line detection. For a good explanation of Hough transform, see:
// http://homepages.inf.ed.ac.uk/rbf/HIPR2/hough.htm
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/dd/d1a/group__imgproc__feature.html#ga8618180a5948286384e3b7ca02f6feeb
//
func HoughLinesP(src Mat, lines Mat, rho float32, theta float32, threshold int) {
	C.HoughLinesP(src.p, lines.p, C.double(rho), C.double(theta), C.int(threshold))
}

// Rectangle draws a simple, thick, or filled up-right rectangle.
// It renders a rectangle with the desired characteristics to the target Mat image.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d6/d6e/group__imgproc__draw.html#ga346ac30b5c74e9b5137576c9ee9e0e8c
//
func Rectangle(img Mat, r image.Rectangle, c color.RGBA, thickness int) {
	cRect := C.struct_Rect{
		x:      C.int(r.Min.X),
		y:      C.int(r.Min.Y),
		width:  C.int(r.Size().X),
		height: C.int(r.Size().Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Rectangle(img.p, cRect, sColor, C.int(thickness))
}

// HersheyFont are the font libraries included in OpenCV.
// Only a subset of the available Hershey fonts are supported by OpenCV.
//
// For more information, see:
// http://sources.isc.org/utils/misc/hershey-font.txt
//
type HersheyFont int

const (
	// FontHersheySimplex is normal size sans-serif font.
	FontHersheySimplex HersheyFont = 0
	// FontHersheyPlain issmall size sans-serif font.
	FontHersheyPlain = 1
	// FontHersheyDuplex normal size sans-serif font
	// (more complex than FontHersheySIMPLEX).
	FontHersheyDuplex = 2
	// FontHersheyComplex i a normal size serif font.
	FontHersheyComplex = 3
	// FontHersheyTriplex is a normal size serif font
	// (more complex than FontHersheyCOMPLEX).
	FontHersheyTriplex = 4
	// FontHersheyComplexSmall is a smaller version of FontHersheyCOMPLEX.
	FontHersheyComplexSmall = 5
	// FontHersheyScriptSimplex is a hand-writing style font.
	FontHersheyScriptSimplex = 6
	// FontHersheyScriptComplex is a more complex variant of FontHersheyScriptSimplex.
	FontHersheyScriptComplex = 7
	// FontItalic is the flag for italic font.
	FontItalic = 16
)

// GetTextSize calculates the width and height of a text string.
// It returns an image.Point with the size required to draw text using
// a specific font face, scale, and thickness.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d6/d6e/group__imgproc__draw.html#ga3d2abfcb995fd2db908c8288199dba82
//
func GetTextSize(text string, fontFace HersheyFont, fontScale float64, thickness int) image.Point {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	sz := C.GetTextSize(cText, C.int(fontFace), C.double(fontScale), C.int(thickness))
	return image.Pt(int(sz.width), int(sz.height))
}

// PutText draws a text string.
// It renders the specified text string into the img Mat at the location
// passed in the "org" param, using the desired font face, font scale,
// color, and line thinkness.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d6/d6e/group__imgproc__draw.html#ga5126f47f883d730f633d74f07456c576
//
func PutText(img Mat, text string, org image.Point, fontFace HersheyFont, fontScale float64, c color.RGBA, thickness int) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	pOrg := C.struct_Point{
		x: C.int(org.X),
		y: C.int(org.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.PutText(img.p, cText, pOrg, C.int(fontFace), C.double(fontScale), sColor, C.int(thickness))
	return
}
