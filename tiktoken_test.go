package tiktoken

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	ass := assert.New(t)
	enc, err := GetWhisperEncoding("multilingual", 100, "./")
	ass.Nil(err, "Encoding  init should not be nil")

	tokens := enc.Encode("<|startofprev|> Nvidia<|startoftranscript|><|en|><|transcribe|>", []string{"all"}, nil)
	sourceTokens := []int{50362, 46284, 50258, 50259, 50360}
	ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

	tokens = enc.Encode("<|startofprev|>สวัสดีค่ะ<|startoftranscript|><|th|><|transcribe|>", []string{"all"}, nil)
	sourceTokens = []int{50362, 13715, 7643, 5981, 13715, 9345, 6033, 8163, 4294, 8055, 50258, 50289, 50360}
	ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

	tokens = enc.Encode("<|startoftranscript|><|th|><|transcribe|><|notimestamps|>", []string{"all"}, nil)
	sourceTokens = []int{50258, 50289, 50360, 50364}
	ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

	tokens = enc.Encode("<|startoftranscript|><|en|><|transcribe|><|notimestamps|>", []string{"all"}, nil)
	sourceTokens = []int{50258, 50259, 50360, 50364}
	ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")
}

func TestDecoding(t *testing.T) {
	ass := assert.New(t)
	enc, err := GetWhisperEncoding("multilingual", 100, "./")
	ass.Nil(err, "Encoding  init should not be nil")

	sourceTokens := []int{50258, 50259, 50360, 50364}
	txt := enc.Decode(sourceTokens)
	ass.Equal("<|startoftranscript|><|en|><|transcribe|><|notimestamps|>", txt, "Decoding should be equal")

	sourceTokens = []int{50258, 50289, 50360, 50364}
	txt = enc.Decode(sourceTokens)
	ass.Equal("<|startoftranscript|><|th|><|transcribe|><|notimestamps|>", txt, "Decoding should be equal")

	sourceTokens = []int{13715, 7643, 5981, 13715, 9345, 6033, 8163, 4294, 8055}
	txt = enc.Decode(sourceTokens)
	ass.Equal("สวัสดีค่ะ", txt, "Decoding should be equal")

}
