package avutil

/*
	#cgo pkg-config: libswresample
	#include <libavutil/channel_layout.h>
*/
import "C"

const (
	AV_CH_LAYOUT_STEREO = C.AV_CH_LAYOUT_STEREO
)

func AvGetDefaultChannelLayout(numChannels int) int64 {
	num := C.int(numChannels)
	ret := C.av_get_default_channel_layout(num)
	return int64(ret)
}
