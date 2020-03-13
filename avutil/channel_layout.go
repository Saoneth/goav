package avutil

//#include <libavutil/channel_layout.h>
//#include <stdlib.h>
import "C"

const (
	AV_CH_FRONT_LEFT            = int64(C.AV_CH_FRONT_LEFT)
	AV_CH_FRONT_RIGHT           = int64(C.AV_CH_FRONT_RIGHT)
	AV_CH_FRONT_CENTER          = int64(C.AV_CH_FRONT_CENTER)
	AV_CH_BACK_LEFT             = int64(C.AV_CH_BACK_LEFT)
	AV_CH_BACK_RIGHT            = int64(C.AV_CH_BACK_RIGHT)
	AV_CH_FRONT_LEFT_OF_CENTER  = int64(C.AV_CH_FRONT_LEFT_OF_CENTER)
	AV_CH_FRONT_RIGHT_OF_CENTER = int64(C.AV_CH_FRONT_RIGHT_OF_CENTER)

	AV_CH_LAYOUT_MONO   = int64(C.AV_CH_LAYOUT_MONO)
	AV_CH_LAYOUT_STEREO = int64(C.AV_CH_LAYOUT_STEREO)
)

func AvGetDefaultChannelLayout(numChannels int) int64 {
	num := C.int(numChannels)
	ret := C.av_get_default_channel_layout(num)
	return int64(ret)
}
