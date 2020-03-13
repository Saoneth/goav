package main

import (
	"log"

	"github.com/saoneth/goav/avcodec"
	"github.com/saoneth/goav/avdevice"
	"github.com/saoneth/goav/avfilter"
	"github.com/saoneth/goav/avformat"
	"github.com/saoneth/goav/avutil"
	"github.com/saoneth/goav/swresample"
	"github.com/saoneth/goav/swscale"
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
