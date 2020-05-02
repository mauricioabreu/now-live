# Now Live - an open streaming platform

Have some fun streaming videos :-)

## What is this?

I created this project to experiment some video tools. Could this be an open streaming platform in the future? Hell yeah!

## How to use?

`make package` starts a packager container ready to receive a UDP streaming.

`make ingest` produce and ingest some video so the packager can produce HLS playlists.

After using it, you can run `make clean` to remove all the media files.