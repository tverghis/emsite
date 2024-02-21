#!/bin/bash

if [ -f "~/images" ]; then
    echo "Removing existing images archive"
    rm ~/images.tar
fi

if [ -d "~/images" ]; then
    echo "Removing existing images directory"
    rm -rf ~/images
fi

echo "Downloading new image archive"
curl -o ~/images.tar http://em.tverghis.space/download

if [ $? -ne 0 ]; then
    echo "Download failed!"
    exit 1
fi

echo "Extracting archive"
mkdir ~/images
tar -xvf ~/images.tar -C ~/images

if [ $? -ne 0 ]; then
    echo "Extraction failed!"
    exit 1
fi

echo "Update completed successfully"
