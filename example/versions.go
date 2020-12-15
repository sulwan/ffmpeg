package main

import (
	"log"

	"github.com/sulwan/ffmpeg/avcodec"
	"github.com/sulwan/ffmpeg/avdevice"
	"github.com/sulwan/ffmpeg/avfilter"
	"github.com/sulwan/ffmpeg/avformat"
	"github.com/sulwan/ffmpeg/avutil"
	"github.com/sulwan/ffmpeg/swresample"
	"github.com/sulwan/ffmpeg/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
