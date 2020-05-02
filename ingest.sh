#!/bin/bash

docker run --net="host" --rm -v $(pwd):/files jrottenberg/ffmpeg:4.1 -hide_banner -re -f lavfi -i 'testsrc2=size=1280x720:rate=60,format=yuv420p' \
    -f lavfi -i 'sine=frequency=440:sample_rate=48000:beep_factor=4' \
    -c:a libfdk_aac -b:a 128x \
    -c:v libx264 -x264opts keyint=30:min-keyint=30:scenecut=-1 -tune zerolatency \
    -b:v 626k -g 30 -r 30 -s 512x288 -preset superfast -profile:v high -level 4.1 \
    -c:a aac -b:a 32k -f mpegts udp://127.0.0.1:40001 \
    -c:a libfdk_aac -b:a 128x \
    -c:v libx264 -x264opts keyint=30:min-keyint=30:scenecut=-1 -tune zerolatency \
    -b:v 1485k -g 30 -r 30 -s 768x432 -preset superfast -profile:v high -level 4.1 \
    -c:a aac -b:a 32k -f mpegts udp://127.0.0.1:40002 \
