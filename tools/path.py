import os
from pathlib import Path
dir = os.path.dirname(__file__)
rootPath = os.path.dirname(dir)

rootPath=Path(rootPath).as_posix()
tempPath = os.path.join(rootPath + "/tmp/")
