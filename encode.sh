#!/bin/bash

ffmpeg -i $1 -vf scale=1920x1080 -b:v 1800k \
  -minrate 900k -maxrate 2610k -tile-columns 2 -g 240 -threads 4 \
  -quality good -crf 31 -c:v libvpx-vp9 -an \
  -pass 1 -speed 4 $2.webm && \
ffmpeg -i $1 -vf scale=1920x1080 -b:v 1800k \
  -minrate 900k -maxrate 2610k -tile-columns 3 -g 240 -threads 4 \
  -quality good -crf 31 -c:v libvpx-vp9 -an \
  -pass 2 -speed 4 -y $2.webm
