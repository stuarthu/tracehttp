package proxy

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

type Server struct{}

func (*Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, e := httputil.DumpRequest(r, true)
	log.Println("\ndump request====\n", string(b), e, "\ndump request finished====\n")
	t := &http.Transport{}
	r.URL.Scheme = "https"
	resp, e := t.RoundTrip(r)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, e = httputil.DumpResponse(resp, true)
	log.Println("\ndump response====\n", string(b), e, "\ndump response finished====\n")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
