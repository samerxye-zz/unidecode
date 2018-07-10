package unidecode

import (
	"testing"
)

const asciiControlCharacters = "\x00\x01\x02\x03\x04\x05\x06\x07" +
	"\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17" +
	"\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f"

const gsmCharacterSet = "\u0040\u00A3\u0024\u00A5\u00E8\u00E9\u00F9\u00EC" +
	"\u00F2\u00E7\u000A\u00D8\u00F8\u000D\u00C5\u00E5" +
	"\u0394\u005F\u03A6\u0393\u039B\u03A9\u03A0\u03A8" +
	"\u03A3\u0398\u039E\u0020\u000C\u005E\u007B\u007D" +
	"\u005C\u005B\u007E\u005D\u007C\u20AC\u00C6\u00E6" +
	"\u00DF\u00C9\u0020\u0021\u0022\u0023\u00A4\u0025" +
	"\u0026\u0027\u0028\u0029\u002A\u002B\u002C\u002D" +
	"\u002E\u002F\u0030\u0031\u0032\u0033\u0034\u0035" +
	"\u0036\u0037\u0038\u0039\u003A\u003B\u003C\u003D" +
	"\u003E\u003F\u00A1\u0041\u0042\u0043\u0044\u0045" +
	"\u0046\u0047\u0048\u0049\u004A\u004B\u004C\u004D" +
	"\u004E\u004F\u0050\u0051\u0052\u0053\u0054\u0055" +
	"\u0056\u0057\u0058\u0059\u005A\u00C4\u00D6\u00D1" +
	"\u00DC\u00A7\u00BF\u0061\u0062\u0063\u0064\u0065" +
	"\u0066\u0067\u0068\u0069\u006A\u006B\u006C\u006D" +
	"\u006E\u006F\u0070\u0071\u0072\u0073\u0074\u0075" +
	"\u0076\u0077\u0078\u0079\u007A\u00E4\u00F6\u00F1" +
	"\u00FC\u00E0"

func testTransliteration(original string, decoded string, t *testing.T) {
	if r := Unidecode(original); r != decoded {
		t.Errorf("Expected '%s', got '%s'\n", decoded, r)
	}
}

func TestASCII(t *testing.T) {
	s := "ABCDEF"
	testTransliteration(s, s, t)
}

func TestASCIIControlCharacters(t *testing.T) {
	o := asciiControlCharacters
	d := "\u0000 \n\u000c\r"
	testTransliteration(o, d, t)
}

func TestKnosos(t *testing.T) {
	o := "Κνωσός"
	d := "Knosos"
	testTransliteration(o, d, t)
}

func TestBeiJing(t *testing.T) {
	o := "\u5317\u4EB0"
	d := "Bei Jing "
	testTransliteration(o, d, t)
}

func TestEmoji(t *testing.T) {
	o := "Hey Luna t belle 😵😂"
	d := "Hey Luna t belle "
	testTransliteration(o, d, t)
}

func TestGSM(t *testing.T) {
	s := gsmCharacterSet
	testTransliteration(s, s, t)
}

func BenchmarkUnidecode(b *testing.B) {
	cases := []string{
		"ABCDEF",
		"Κνωσός",
		"\u5317\u4EB0",
	}
	for ii := 0; ii < b.N; ii++ {
		for _, v := range cases {
			_ = Unidecode(v)
		}
	}
}

func BenchmarkDecodeTable(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		decodeTransliterations()
	}
}

func init() {
	decodeTransliterations()
}
