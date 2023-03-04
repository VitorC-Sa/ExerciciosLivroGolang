package Section_1_4

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

/*
	Precisa gerar uma seed na main "rand.Seed(time.Now().UTC().UnixNano())"
	e chamar a função  "lissajous" passando o "os.Stdout"
*/

var palette = []color.Color{color.RGBA{6, 255, 0, 1}, color.RGBA{0, 0, 0, 1}}

const (
	blackIndex = 0 //primeira cor da paleta
	greenIndex = 1 //próxima cor da paleta
)

func Lissajous(out io.Writer) {
	const (
		cycles  = 5     //número de revoluções completas do oscilador x
		res     = 0.001 //resolução angular
		size    = 100   //canvas da imagem cobre de [-size..+size]
		nframes = 64    //número de quadros da animação
		delay   = 8     //tempo entre quadros em unidades de 10ms
	)
	freq := rand.Float64() * 3.0 //frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //diferença de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
