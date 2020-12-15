// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

//BitStreamFilter *av_bitstream_filter_next (const BitStreamFilter *f)
func (f *BitStreamFilter) AvBitstreamFilterNext() *BitStreamFilter {
	return (*BitStreamFilter)(C.av_bitstream_filter_next((*C.struct_AVBitStreamFilter)(f)))
}

//Filter bitstream.
func (bfx *BitStreamFilterContext) AvBitstreamFilterFilter(ctxt *Context, a string, p **uint8, ps *int, b *uint8, bs, k int) int {
	return int(C.av_bitstream_filter_filter((*C.struct_AVBitStreamFilterContext)(bfx), (*C.struct_AVCodecContext)(ctxt), C.CString(a), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

//Create and initialize a bitstream filter context given a bitstream filter name.
func AvBitstreamFilterInit(n string) *BitStreamFilterContext {
	return (*BitStreamFilterContext)(C.av_bitstream_filter_init(C.CString(n)))
}
