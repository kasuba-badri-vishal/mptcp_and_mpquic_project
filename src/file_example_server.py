import socket                   # Import socket module

port = 6000                # Reserve a port for your service.
s = socket.socket()             # Create a socket object
host = "192.168.0.107"    # Get local machine name
s.bind((host, port))            # Bind to the port
s.listen(5)                     # Now wait for client connection.

print 'Server listening....'

while True:
    conn, addr = s.accept()     # Establish connection with client.
    print 'Got connection from', addr
    data = conn.recv(1024)
    print('Server received', repr(data))

    filename='Anaconda3-2020.02-Linux-x86_64.sh'
    f = open(filename,'rb')
    l = f.read(1024)
    while (l):
       conn.send(l)
    #    print('Sent ',repr(l))
       l = f.read(1024)
    f.close()

    print('Done sending')
    conn.send('Thank you for connecting')
    conn.close()