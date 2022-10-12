package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Facade — структурный паттерн проектирования, который предоставляет простой интерфейс к набору взаимосвязанных
	классов или объектов некоторой подсистемы, что облегчает ее использование.

	Use-cases:
		- Когда нужно представить простой или урезанный интерфейс к сложной подсистеме
		- Когда вы хотите разложить подсистему на отдельные слои

	Props:
		-  Изолирует клиентов от компонентов сложной подсистемы

	Cons:
		- Фасад рискует стать божественным объектом, привязанным ко всем классам программы

	Examples:
		- Использование комплексной сторонней библеотеки

*/

// sets of types of some complex subsystem

type AudioType int

const (
	MP3 AudioType = iota
	WAV
)

type AudioFile struct {
	path      string
	audioType AudioType
}

func (f AudioFile) extractExtension(path string) AudioType {
	return MP3
}

type CompressionCodec interface {
}

type SBCCompressionCodec struct {
}

type AACCompressionCodec struct {
}

type BitrateReader struct {
}

func (r BitrateReader) read(file AudioFile, codec CompressionCodec) []byte {
	fmt.Println("Reading file")
	return []byte{}
}

type BitrateWriter struct {
}

func (w BitrateWriter) write(data []byte, codec CompressionCodec) AudioFile {
	fmt.Println("Writing file")
	return AudioFile{audioType: WAV}
}

// AudioConverterFacade is Facade pattern for 3rd-party audio lib
type AudioConverterFacade struct {
}

func (f AudioConverterFacade) convert(path string, audioType AudioType) AudioFile {
	file := AudioFile{path: path}
	file.audioType = file.extractExtension(path)

	var sourceCodec CompressionCodec
	var destCodec CompressionCodec

	if file.audioType == MP3 {
		sourceCodec = SBCCompressionCodec{}
	} else if file.audioType == WAV {
		sourceCodec = AACCompressionCodec{}
	}

	if audioType == MP3 {
		destCodec = SBCCompressionCodec{}
	} else if audioType == WAV {
		destCodec = AACCompressionCodec{}
	}

	buf := BitrateReader{}.read(file, sourceCodec)
	return BitrateWriter{}.write(buf, destCodec)
}

func main() {
	convertedFile := AudioConverterFacade{}.convert("./example.mp3", WAV)
	fmt.Println(convertedFile.audioType)
}
