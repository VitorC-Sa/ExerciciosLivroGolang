package Section_1_7

import (
	"encoding/json"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

//Modifique o servidor/função Lissajous para ler valores de parâmetros do URL. (localhost:8080/?cycles=20...)
//Curl example: curl localhost:8080 -G -d "size=500" -d "cycles=100" -d "delay=16"  --output tmp.gif & code tmp.gif
func Ex01() {
	log.Println("Starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var settings = NewLissajousObj()

		parseInt := func(s string) *int {
			if s != "" {
				v, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				return &v
			}
			return nil
		}
		parseFloat := func(s string) *float64 {
			if s != "" {
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					panic(err)
				}
				return &v
			}
			return nil
		}

		queryValues := r.URL.Query()
		var m = map[string]interface{}{
			"cycles":  parseInt(queryValues.Get("cycles")),
			"res":     parseFloat(queryValues.Get("res")),
			"size":    parseInt(queryValues.Get("size")),
			"nframes": parseInt(queryValues.Get("nframes")),
			"delay":   parseInt(queryValues.Get("delay")),
		}

		jsonbody, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
			return
		}

		if err := json.Unmarshal(jsonbody, &settings); err != nil {
			log.Fatal(err)
			return
		}

		log.Println(settings)
		lissajous(w, settings)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	log.Println("Ending server...")
}

type lissajousObj struct {
	Cycles  int
	Res     float64
	Size    int
	Nframes int
	Delay   int
}

func NewLissajousObj() lissajousObj {
	return lissajousObj{
		Cycles:  DEFAULT_CYCLES,
		Res:     DEFAULT_RES,
		Size:    DEFAULT_SIZE,
		Nframes: DEFAULT_NFRAMES,
		Delay:   DEFAULT_DELAY,
	}
}

const (
	DEFAULT_CYCLES  = 5     //número de revoluções completas do oscilador x
	DEFAULT_RES     = 0.001 //resolução angular
	DEFAULT_SIZE    = 200   //canvas da imagem cobre de [-size..+size]
	DEFAULT_NFRAMES = 64    //número de quadros da animação
	DEFAULT_DELAY   = 8     //tempo entre quadros em unidades de 10ms
)

var palette = []color.Color{color.RGBA{6, 255, 0, 1}, color.RGBA{0, 0, 0, 1}}

const (
	blackIndex = 0 //primeira cor da paleta
	greenIndex = 1 //próxima cor da paleta
)

func lissajous(out io.Writer, settings lissajousObj) {
	freq := rand.Float64() * 3.0 //frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: settings.Nframes}
	phase := 0.0 //diferença de fase
	for i := 0; i < settings.Nframes; i++ {
		rect := image.Rect(0, 0, 2*settings.Size+1, 2*settings.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(settings.Cycles)*2*math.Pi; t += settings.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(settings.Size+int(x*float64(settings.Size)+0.5), settings.Size+int(y*float64(settings.Size)+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, settings.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
