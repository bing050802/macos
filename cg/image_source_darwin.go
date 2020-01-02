// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cg

import "github.com/richardwilkes/macos/cf"

// #import <ImageIO/ImageIO.h>
import "C"

type ImageSource = C.CGImageSourceRef

func ImageSourceCreateWithData(data cf.Data, options cf.Dictionary) ImageSource {
	return C.CGImageSourceCreateWithData(C.CFDataRef(data), C.CFDictionaryRef(options))
}

func (is ImageSource) CreateImageAtIndex(index int, options cf.Dictionary) Image {
	return C.CGImageSourceCreateImageAtIndex(is, C.size_t(index), C.CFDictionaryRef(options))
}

func (is ImageSource) Release() {
	C.CFRelease(C.CFTypeRef(is))
}
