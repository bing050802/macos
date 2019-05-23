package macos

// #import <CoreFoundation/CoreFoundation.h>
import "C"
import "unsafe"

type CFString = C.CFStringRef

func CFStringCreateWithString(str string) CFString {
	return CFStringCreateWithBytes([]byte(str))
}

func CFStringCreateWithBytes(bytes []byte) CFString {
	return C.CFStringCreateWithBytes(0, (*C.UInt8)(unsafe.Pointer(&bytes[0])), C.CFIndex(len(bytes)), C.kCFStringEncodingUTF8, 0)
}

func (s CFString) GetLength() int {
	return int(C.CFStringGetLength(s))
}

func (s CFString) String() string {
	var allocedPtr *C.char
	strPtr := C.CFStringGetCStringPtr(s, C.kCFStringEncodingUTF8)
	if strPtr == nil {
		stringLength := s.GetLength()
		maxBytes := 4*stringLength + 1
		allocedPtr = (*C.char)(C.malloc(C.size_t(maxBytes)))
		C.CFStringGetCString(s, allocedPtr, C.CFIndex(maxBytes), C.kCFStringEncodingUTF8)
		strPtr = allocedPtr
	}
	str := C.GoString(strPtr)
	if allocedPtr != nil {
		C.free(unsafe.Pointer(allocedPtr))
	}
	return str
}

func (s CFString) Release() {
	C.CFRelease(CFType(s))
}