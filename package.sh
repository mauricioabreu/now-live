#!/bin/bash

docker run -v "$(pwd)/media/:/media" --net=host --rm google/shaka-packager packager \
  'in=udp://127.0.0.1:40001?reuse=1,stream=video,segment_template=/media/live_288p/$Number$.ts,playlist_name=/media/live_288p/main.m3u8,iframe_playlist_name=/media/live_288p/iframe.m3u8' \
  'in=udp://127.0.0.1:40002?reuse=1,stream=video,segment_template=/media/live_432p/$Number$.ts,playlist_name=/media/live_432p/main.m3u8,iframe_playlist_name=/media/live_432p/iframe.m3u8' \
  --hls_master_playlist_output /media/live.m3u8 \
  --hls_playlist_type LIVE