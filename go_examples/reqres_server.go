package main

import (
	//"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	quic "github.com/lucas-clemente/quic-go"
	// "github.com/lucas-clemente/quic-go/internal/utils"
	"io"
	"math/big"
)

var addr = "0.0.0.0:4242"

const (
	MsgLen    = 750
	MinFields = 5
)

// We start a server echoing data on the first stream the client opens,
// then connect with a client, send the message, and wait for its receipt.
func main() {
	addrF := flag.String("addr", "0.0.0.0:4242", "Address to bind")
	flag.Parse()
	addr = *addrF
	err := echoServer()
	if err != nil {
		panic(err)
	}
}

// Start a server that performs similar traffic to Siri servers
func echoServer() error {
	cfgServer := &quic.Config{}
	tlsConfig := generateTLSConfig()
	fmt.Println("Address", addr)
	listener, err := quic.ListenAddr(addr, tlsConfig, cfgServer)
	if err != nil {
		return err
	}
	sess, err := listener.Accept()
	if err != nil {
		return err
	}
	stream, err := sess.AcceptStream()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, MsgLen)
	_, err = io.ReadAtLeast(stream, buf, 10)
	if err != nil {
		stream.Close()
		stream.Close()
		fmt.Println("Error", err)
		return err
	}
	msg := string(buf)
	fmt.Println("MESSAGE:", msg)
	//splitMsg := strings.Split(msg, "&")
	//res := msgID + "&" + strings.Repeat("0", resSize-len(msgID)-2) + "\n"
	res := "Hello, this is your server"
	_, err = stream.Write([]byte(res))
	if err != nil {
		stream.Close()
		stream.Close()
		return err
	}
	return err
}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{Certificates: []tls.Certificate{tlsCert}}
}
