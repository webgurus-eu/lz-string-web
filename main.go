package main

import (
	"syscall/js"

	lzstring "github.com/webgurus-eu/lz-string"
)

func main() {
	c := make(chan struct{}, 0)

	println("Go WebAssembly Initialized")

	js.Global().Set("encode", js.FuncOf(encode))
	js.Global().Set("decode", js.FuncOf(decode))

	<-c
}

func encode(this js.Value, inputs []js.Value) interface{} {
	firstArg := inputs[0].String()
	return js.ValueOf(lzstring.CompressToBase64(firstArg))
}

func decode(this js.Value, inputs []js.Value) interface{} {
	firstArg := inputs[0].String()
	res, err := lzstring.DecompressFromBase64(firstArg)
	if err != nil {
		return js.ValueOf(err)
	}
	return js.ValueOf(res)
}
