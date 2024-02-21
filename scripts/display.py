import sys
import os
import time
import random

from PIL import Image
from itertools import cycle

from inky.auto import auto

saturation = 0.5
sleep_sec = 30 * 60 # 30 mins
img_path = "/home/pi/images"

cur_img = None

inky = auto(ask_user=True, verbose=True)

def select_image():
    if not os.path.exists(img_path):
        return None

    images = os.listdir(img_path)
    img = None

    while img == cur_img:
        img = random.choice(images)

    return img


def show_image(f):
    cur_img = f

    img = Image.open(os.path.join(img_path, f))
    resized = img.resize(inky.resolution)

    inky.set_image(resized, saturation=saturation)
    inky.show()

if __name__ == "__main__":
    while True:
        if not os.path.exists(img_path):
            print("Image directory does not exist, doing nothing")
            sys.exit(0)

        img = select_image()
        print(f"Displaying image `{img}`")
        show_image(img)

        print(f"Sleeping {sleep_sec} seconds")
        time.sleep(sleep_sec)
