import socket
import os
import time
import argparse

def client(host,port):
    start_time = time.time()
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.connect((host, port))
    print('Client has been assigned socket name', sock.getsockname())
    sock.send(b"Hello server!")
    with open('received_file', 'wb') as f:
        print('file opened')
        while True:
            data = sock.recv(1024)
            if not data:
                break
            f.write(data)
    f.close()
    print('Successfully get the file')
    sock.close()
    print('connection closed')
    end_time = time.time()
    total_time = end_time-start_time
    print("Total Time taken is ",total_time)

def server(interface,port):
    
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    # sock.setsockopt(socket.IPPROTO_TCP,socket.TCP_NODELAY,1)
    print(sock)
    sock.bind((interface,port))
    sock.listen(5)
    print('Listening at', sock.getsockname())
    while True:
        client_sock,address = sock.accept()
        print('We have accepted a connection from', address)
        print('Socket name:', client_sock.getsockname())
        print('Socket peer:', client_sock.getpeername())
        data = client_sock.recv(1024)
        filename='bigfile.sh'
        f = open(filename,'rb')
        l = f.read(1024)
        while (l):
            client_sock.send(l)
            l = f.read(1024)
        f.close()
        print('Done sending')
        client_sock.send(b'Thank you for connecting')
        client_sock.close()


if __name__ == "__main__":
    choices = {'client':client,'server':server}
    parser = argparse.ArgumentParser(description="Client-Server File Transfer Using MPTCP")
    parser.add_argument('role',choices=choices,help='Which role, Client or Server?')
    parser.add_argument('host',help='Interface at which Server listens at')
    parser.add_argument('-p',metavar='PORT',type=int,default=1060,help='TCP Port (default 1060)')
    args = parser.parse_args()
    function = choices[args.role]
    function(args.host,args.p)
