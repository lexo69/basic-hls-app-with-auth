## Basic HLS 'streaming' app with authorization written in Go as a prof of concept.

```console
curl -X POST http://127.0.0.1:9000/playlist -d '{"login":"test","password":"12345"}' -H "Content-Type: application/json"
```
Will generate and return a playlist for a user which then can be played.
You can convert a video to HLS format with
```console
ffmpeg -i input_file.mp4 -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls output_file.m3u8 
```
