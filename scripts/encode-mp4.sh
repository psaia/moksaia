#!/bin/sh
# ./scripts/encode-mp4.sh ./docs/raw/video.mp4 ./docs/mp4/video.mp4
ffmpeg -i $1 -preset slow -tune film -y -filter:v "scale=-1:800" -vcodec h264 -crf 29 -an $2
