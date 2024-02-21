#!/bin/bash

if [ -f "/home/pi/images.tar" ]; then
    echo "Removing existing images archive"
    rm /home/pi/images.tar
fi

echo "Downloading new image archive"
curl -o /home/pi/images.tar http://em.tverghis.space/download

if [ $? -ne 0 ]; then
    echo "Download failed!"
    exit 1
fi

if [ -d "/home/pi/images" ]; then
    echo "Removing existing images directory"
    rm -rf /home/pi/images
else
    echo "No pre-exiting images directory"
fi

mkdir /home/pi/images

echo "Extracting archive"
tar -xvf /home/pi/images.tar -C /home/pi/images

if [ $? -ne 0 ]; then
    echo "Extraction failed!"
    exit 1
fi

echo "Update completed successfully"
