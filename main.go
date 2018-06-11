package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	address    = flag.String("a", ":18080", `bind address.`)
	privateKey *rsa.PrivateKey
)

const (
	cert = `-----BEGIN RSA PRIVATE KEY-----
MIIBOwIBAAJBALecq3BwAI4YJZwhJ+snnDFj3lF3DMqNPorV6y5ZKXCiCMqj8OeO
mxk4YZW9aaV9ckl/zlAOI0mpB3pDT+Xlj2sCAwEAAQJAW6/aVD05qbsZHMvZuS2A
a5FpNNj0BDlf38hOtkhDzz/hkYb+EBYLLvldhgsD0OvRNy8yhz7EjaUqLCB0juIN
4QIhAMsJQ3xiJemnJ2pD65iRNCC/Kr7jtxbbBwa6ZFLjp12pAiEA54JCn41fF8GZ
90b9L5dtFQB2/yIcGX4Xo7bCvl8DaPMCIBgOZ+2T33QYtwXTOFXiVm/O1qy5ZFcT
6ng0m3BqwsjJAiEAqna/l7wAyP1E4U7kHqbhKxWsiTAUgLDXtzRbMNHFMQECIQCA
xlpXEPqnC3P8if0G9xHomqJ531rOJuzB8fNtRFmxnA==
-----END RSA PRIVATE KEY-----`
)

func init() {
	flag.Parse()
	block, _ := pem.Decode([]byte(cert))

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Panicln(err)
	}
	privateKey = key
}

func main() {
	http.HandleFunc("/rpc/releaseTicket.action", releaseTicketHandler)
	http.HandleFunc("/rpc/obtainTicket.action", obtainTicketHandler)
	http.HandleFunc("/rpc/ping.action", pingTicketHandler)
	http.HandleFunc("/rpc/prolongTicket.action", prolongTicketHandler)

	http.Handle("/yaaw/", http.StripPrefix("/yaaw", http.FileServer(yaaw)))
	http.Handle("/aria/", http.StripPrefix("/aria", http.FileServer(aria)))

	http.ListenAndServe(*address, nil)
}

func releaseTicketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Release Ticket request")
	salt := r.URL.Query().Get("salt")
	str := fmt.Sprintf("<ReleaseTicketResponse><message></message>"+
		"<responseCode>OK</responseCode><salt>%s</salt></ReleaseTicketResponse>", salt)

	writeAnswer(w, str)
}

func obtainTicketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Obtain Ticket request")
	salt := r.URL.Query().Get("salt")
	userName := r.URL.Query().Get("userName")

	str := fmt.Sprintf("<ObtainTicketResponse>"+
		"<message></message>"+
		"<prolongationPeriod>607875500</prolongationPeriod>"+
		"<responseCode>OK</responseCode>"+
		"<salt>%s</salt>"+
		"<ticketId>1</ticketId>"+
		"<ticketProperties>licensee=%s\tlicenseType=0\t</ticketProperties>"+
		"</ObtainTicketResponse>", salt, userName)

	writeAnswer(w, str)
}

func prolongTicketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Prolong Ticket request")
	salt := r.URL.Query().Get("salt")

	str := fmt.Sprintf("<ProlongTicketResponse>"+
		"<message></message>"+
		"<responseCode>OK</responseCode>"+
		"<salt>%s</salt>"+
		"<ticketId>1</ticketId>"+
		"</ProlongTicketResponse>", salt)

	writeAnswer(w, str)
}

func pingTicketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping Ticket request")
	salt := r.URL.Query().Get("salt")

	str := fmt.Sprintf("<PingResponse>"+
		"<message></message>"+
		"<responseCode>OK</responseCode>"+
		"<salt>%s</salt>"+
		"</PingResponse>", salt)

	writeAnswer(w, str)
}

func writeAnswer(w http.ResponseWriter, str string) {
	h := crypto.MD5.New()
	h.Write([]byte(str))
	hashed := h.Sum(nil)

	bs, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.MD5, hashed)

	str = fmt.Sprintf("<!-- %x -->\n%s", string(bs), str)
	w.Write([]byte(str))
}

type handle struct {
	reverseProxy string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse(this.reverseProxy)
	if err != nil {
		log.Fatalln(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	r.Host = remote.Host
	proxy.ServeHTTP(w, r)
	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
}

func main2() {
	bind := flag.String("l", "0.0.0.0:8888", "listen on ip:port")
	remote := flag.String("r", "http://idea.lanyus.com:80", "reverse proxy addr")
	flag.Parse()
	log.Printf("Listening on %s, forwarding to %s", *bind, *remote)
	h := &handle{reverseProxy: *remote}
	err := http.ListenAndServe(*bind, h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
