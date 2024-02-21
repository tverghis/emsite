import sys
import os
import random

from PIL import Image

from inky.auto import auto

saturation = 0.5
img_path = "/home/pi/images"

inky = auto(ask_user=True, verbose=True)

def select_image():
    images = os.listdir(img_path)
    return random.choice(images)

def show_image(f):
    img = Image.open(os.path.join(img_path, f))
    resized = img.resize(inky.resolution)

    inky.set_image(resized, saturation=saturation)
    inky.show()

if __name__ == "__main__":
    if not os.path.exists(img_path):
        print("Image directory does not exist, doing nothing")
        sys.exit(0)

    img = select_image()
    print(f"Displaying image `{img}`")
    show_image(img)
