import socket
from progmp import ProgMP

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(("127.0.0.1", 1061))

schedulerName = "python_api_example"
with open("api_example.progmp", "r") as src:
    schedProgStr = src.read()
try:
    ProgMp.loadScheduler(schedProgStr)
except:
    print "Scheduler loading error."

try:
    ProgMp.setScheduler(s, schedulerName)
except:
    print "Scheduler not found, maybe no MPTCP?"

ProgMp.setUser(s, 2)
ProgMp.setRegister(s, Progmp.R1(), 5)
s.send("Multipath is awesome!")
