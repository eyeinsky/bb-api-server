package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os/exec"
	"sync/atomic"
)

type Radio struct {
	name string
	args []string
}

var p1 = Radio{"p1", []string{"http://http-live.sr.se/p1-mp3-192"}}
var kuku = Radio{"kuku", []string{"http://striiming.trio.ee:8008/kuku.mp3"}}
var r2 = Radio{"r2", []string{"-playlist", "http://r2.err.ee/gfx/raadio2128.m3u"}}
var viker = Radio{"viker", []string{"http://icecast.err.ee:80/vikerraadio.mp3"}}
var brock = Radio{"brock", []string{"http://fm02-icecast.mtg-r.net/fm02_aac"}}

var radios = map[string]Radio{
	"p1":    p1,
	"kuku":  kuku,
	"r2":    r2,
	"viker": viker,
	"brock": brock,
}

func main() {
	var current atomic.Value
	current.Store("off")

	kill := func() {
		exec.Command("killall", "mplayer").Run()
		current.Store("off")
	}

	mux := http.NewServeMux()

	makeHandler := func(name *string, radio *Radio) func(w http.ResponseWriter, req *http.Request) {
		n := *name
		r := *radio
		return func(w http.ResponseWriter, req *http.Request) {
			kill()
			exec.Command("mplayer", r.args...).Start()
			fmt.Fprintf(w, html.EscapeString(n))
			current.Store(r.name)
		}
	}

	for name, radio := range radios {
		mux.HandleFunc("/"+name, makeHandler(&name, &radio))
	}

	mux.HandleFunc("/playing", func(w http.ResponseWriter, r *http.Request) {
		station := current.Load().(string)
		fmt.Println("Playing " + station)
		fmt.Fprintf(w, html.EscapeString(station))
	})

	mux.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request) {
		kill()
		fmt.Fprintf(w, "off")
	})

	fmt.Println("Serving")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
