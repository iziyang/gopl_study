package server3

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)
//
//var mu sync.Mutex
//var count int

var palett = []color.Color{color.White, color.Black}

const (
	whiteInde = 0
	blackInde = 1
)

func main() {
	//http.HandleFunc("/", handler3)
	//http.HandleFunc("/count", counter)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

// 处理程序回显请求的 URL 的路径部分
func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q = %q\n]", k, v)
	}

}

//counter 回显目前为止调用的次数
//func counter(w http.ResponseWriter, r *http.Request) {
//	mu.Lock()
//	fmt.Fprintf(w, "Count %d\n", count)
//	mu.Unlock()
//
//}

func Lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palett)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackInde)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
