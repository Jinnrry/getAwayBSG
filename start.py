import sys

from tools import path

sys.path.append(path.rootPath)

from spider import ziroom

ziroom.start()