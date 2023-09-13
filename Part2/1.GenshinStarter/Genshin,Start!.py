import os
import random, time
import win32com.client 
from ctypes import *
import win32api,win32con

screenX = win32api.GetSystemMetrics(win32con.SM_CXSCREEN)//1.1
screenY = win32api.GetSystemMetrics(win32con.SM_CYSCREEN)//1.1

started = False
 
def get_color(x, y):
    gdi32 = windll.gdi32
    user32 = windll.user32
    hdc = user32.GetDC(None)  # 获取颜色值
    pixel = gdi32.GetPixel(hdc, x, y)  # 提取RGB值
    r = pixel & 0x0000ff
    g = (pixel & 0x00ff00) >> 8
    b = pixel >> 16
    return [r, g, b]

def voter(count=100):
    for i in range(count):
        color = get_color(random.randint(0,screenX), random.randint(0,screenY))
        if color[0] + color[1] + color[2] < 240*3:
            # print(color)
            return False
    return True

if __name__ == "__main__":

    try:
        assert os.name == "nt"
    except:
        print("Playing Genshin on *nix is cool, but this script only embrace Windows ;)")
        exit(0)

    原神路径 = ""

    try:
        StartMenuPath = r'C:\ProgramData\Microsoft\Windows\Start Menu\Programs'
        assert os.path.isdir(StartMenuPath + "\原神")
        LinkPath = StartMenuPath + "\\原神\\原神.lnk"
        shell = win32com.client.Dispatch("WScript.Shell")
        shortcut = shell.CreateShortCut(LinkPath)
        原神路径 = "\\".join([*shortcut.Targetpath.split("\\")[:-1], *["Genshin Impact Game", "YuanShen.exe"]])

        print(f"Current Start Path-> {原神路径}")
    except Exception as e:
        print("原神？启动失败！" + e)


    # playsound("shed_a_light.wav", False)

    print("Wait for it...")
    while(True):
        if voter() and not started:
            srarted = True
            print("原神？启动！")
            os.system(os.path.abspath(os.path.dirname(__file__))+"\shed_a_light.wav")
            os.system(f"\"{原神路径}\" -popupwindow -screen-fullscreen 0") #跳脸输出
            exit(0)
        time.sleep(2)
