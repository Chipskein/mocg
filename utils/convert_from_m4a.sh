#!/bin/sh
cd ../internal/testAudios/m4a
for file in * ;do
    filename=$(basename -- "$file")
    filename="${filename%.*}"
    echo "ffmpeg -i $file -ar 48000 $filename.ogg -ar 48000 $filename.wav $filename.mp3 $filename.flac" ;
    ffmpeg -i "$file" -ar 48000 "$filename.ogg" -ar 48000 "$filename.wav" -ar 48000 "$filename.mp3" -ar 48000 "$filename.flac"; 
done;