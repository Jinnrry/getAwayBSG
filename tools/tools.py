import os
from datetime import datetime, timedelta, timezone


def writeLog(type, info, filename='log.txt'):
    utc_dt = datetime.utcnow().replace(tzinfo=timezone.utc)
    cn_dt = utc_dt.astimezone(timezone(timedelta(hours=8)))
    dir = os.path.dirname(__file__)
    path = dir + "/" + filename
    with open(path, "a") as log:
        log.write(cn_dt.strftime("%Y-%m-%d %H:%M:%S") + "\n")
        log.write(type + ":" + info + "\n")
