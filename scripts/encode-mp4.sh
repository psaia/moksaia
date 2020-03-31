#!/bin/sh
ffmpeg -i $1 -preset slow -tune film -y -filter:v "scale=-1:800" -vcodec h264 -crf 29 -an $2
