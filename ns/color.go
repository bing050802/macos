// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ns

/*
#import <Cocoa/Cocoa.h>

typedef void *NSColorPtr;
typedef void *NSColorSpacePtr;

CFArrayRef nsColorAlternatingContentBackgroundColors() {
	return (CFArrayRef)[NSColor alternatingContentBackgroundColors];
}

NSColorPtr nsColorUsingColorSpace(NSColorPtr color, NSColorSpacePtr space) {
	return [(NSColor *)color colorUsingColorSpace:(NSColorSpace *)space];
}

void nsColorGetRedGreenBlueAlpha(NSColorPtr color, CGFloat *r, CGFloat *g, CGFloat *b, CGFloat *a) {
	[(NSColor *)color getRed:r green:g blue:b alpha:a];
}

void nsColorDispose(NSColorPtr color) {
	[((NSColor *)color) release];
}
*/
import "C"
import "github.com/richardwilkes/macos/cf"

type Color struct {
	native C.NSColorPtr
}

func AlternatingContentBackgroundColors() []*Color {
	data := cf.Array(C.nsColorAlternatingContentBackgroundColors())
	colors := make([]*Color, data.GetCount())
	for i := range colors {
		colors[i] = &Color{native: C.NSColorPtr(data.GetValueAtIndex(i))}
	}
	return colors
}

func (c *Color) ColorUsingColorSpace(space *ColorSpace) *Color {
	return &Color{native: C.nsColorUsingColorSpace(c.native, space.native)}
}

func (c *Color) GetRedGreenBlueAlpha() (r, g, b, a float64) {
	var rr, gg, bb, aa C.CGFloat
	C.nsColorGetRedGreenBlueAlpha(c.native, &rr, &gg, &bb, &aa)
	return float64(rr), float64(gg), float64(bb), float64(aa)
}

func (c *Color) Release() {
	if c.native != nil {
		C.nsColorDispose(c.native)
		c.native = nil
	}
}
