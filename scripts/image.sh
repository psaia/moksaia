#!/bin/sh
# ./image.sh ./docs/raw/video.mp4 ./docs/images/image.png
ffmpeg -i $1 -ss 00:00:00 -vframes 1 $2
