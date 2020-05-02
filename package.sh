#!/bin/bash

docker run -v "$(pwd)/media/:/media" --net=host --rm google/shaka-packager packager \
  'in=udp://127.0.0.1:40001?reuse=1,stream=video,segment_template=/media/live_360p/$Number$.ts,playlist_name=/media/live_360p/main.m3u8,iframe_playlist_name=/media/live_360p/iframe.m3u8' \
  'in=udp://127.0.0.1:40002?reuse=1,stream=video,segment_template=/media/live_480p_$Number$.ts,playlist_name=/media/live_480p/main.m3u8,iframe_playlist_name=/media/live_480p/iframe.m3u8' \
  --hls_master_playlist_output /media/live.m3u8 \
  --hls_playlist_type LIVE