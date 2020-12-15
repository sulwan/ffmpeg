// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

// AvpictureLayout copies pixel data from an Picture into a buffer.
func (p *Picture) AvpictureLayout(pf PixelFormat, w, h int, d *string, ds int) int {
	return int(C.avpicture_layout((*C.struct_AVPicture)(p), (C.enum_AVPixelFormat)(pf), C.int(w), C.int(h), (*C.uchar)(unsafe.Pointer(d)), C.int(ds)))
}

// AvPictureCopy -Copy image src to dst.
func (p *Picture) AvPictureCopy(d *Picture, pf PixelFormat, w, h int) {
	C.av_picture_copy((*C.struct_AVPicture)(d), (*C.struct_AVPicture)(p), (C.enum_AVPixelFormat)(pf), C.int(w), C.int(h))
}

// AvPicturePad - Pad image.
func (p *Picture) AvPicturePad(d *Picture, h, w int, pf PixelFormat, t, b, l, r int, c *int) int {
	return int(C.av_picture_pad((*C.struct_AVPicture)(d), (*C.struct_AVPicture)(p), (C.int)(h), (C.int)(w), (C.enum_AVPixelFormat)(pf), (C.int)(t), (C.int)(b), (C.int)(l), (C.int)(r), (*C.int)(unsafe.Pointer(c))))
}

// AvpictureGetSize - Calculate the size in bytes that a picture of the given width and height would occupy if stored in the given picture format.
func AvpictureGetSize(pf PixelFormat, w, h int) int {
	return int(C.avpicture_get_size((C.enum_AVPixelFormat)(pf), C.int(w), C.int(h)))
}
