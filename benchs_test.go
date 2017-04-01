package passgo

import "testing"

func benchmarkPassGen(flags []string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		PassGen(flags)
	}
}

func BenchmarkPassGenD10(b *testing.B)       { benchmarkPassGen([]string{"10", "-d"}, b) }
func BenchmarkPassGenD100(b *testing.B)      { benchmarkPassGen([]string{"100", "-d"}, b) }
func BenchmarkPassGenDL100(b *testing.B)     { benchmarkPassGen([]string{"100", "-dl"}, b) }
func BenchmarkPassGenDU100(b *testing.B)     { benchmarkPassGen([]string{"100", "-du"}, b) }
func BenchmarkPassGenDS100(b *testing.B)     { benchmarkPassGen([]string{"100", "-ds"}, b) }
func BenchmarkPassGenDLU100(b *testing.B)    { benchmarkPassGen([]string{"100", "-dlu"}, b) }
func BenchmarkPassGenDLS100(b *testing.B)    { benchmarkPassGen([]string{"100", "-dls"}, b) }
func BenchmarkPassGenDSU100(b *testing.B)    { benchmarkPassGen([]string{"100", "-dsu"}, b) }
func BenchmarkPassGenDLSU100(b *testing.B)   { benchmarkPassGen([]string{"100", "-dlsu"}, b) }
func BenchmarkPassGenS100(b *testing.B)      { benchmarkPassGen([]string{"100", "-s"}, b) }
func BenchmarkPassGenSL100(b *testing.B)     { benchmarkPassGen([]string{"100", "-sl"}, b) }
func BenchmarkPassGenSU100(b *testing.B)     { benchmarkPassGen([]string{"100", "-su"}, b) }
func BenchmarkPassGenSLU100(b *testing.B)    { benchmarkPassGen([]string{"100", "-slu"}, b) }
func BenchmarkPassGenL100(b *testing.B)      { benchmarkPassGen([]string{"100", "-l"}, b) }
func BenchmarkPassGenLU100(b *testing.B)     { benchmarkPassGen([]string{"100", "-lu"}, b) }
func BenchmarkPassGenU100(b *testing.B)      { benchmarkPassGen([]string{"100", "-u"}, b) }
func BenchmarkPassGenD1000(b *testing.B)     { benchmarkPassGen([]string{"1000", "-d"}, b) }
func BenchmarkPassGenDL1000(b *testing.B)    { benchmarkPassGen([]string{"1000", "-dl"}, b) }
func BenchmarkPassGenDU1000(b *testing.B)    { benchmarkPassGen([]string{"1000", "-du"}, b) }
func BenchmarkPassGenDS1000(b *testing.B)    { benchmarkPassGen([]string{"1000", "-ds"}, b) }
func BenchmarkPassGenDLU1000(b *testing.B)   { benchmarkPassGen([]string{"1000", "-dlu"}, b) }
func BenchmarkPassGenDLS1000(b *testing.B)   { benchmarkPassGen([]string{"1000", "-dls"}, b) }
func BenchmarkPassGenDSU1000(b *testing.B)   { benchmarkPassGen([]string{"1000", "-dsu"}, b) }
func BenchmarkPassGenDLSU1000(b *testing.B)  { benchmarkPassGen([]string{"1000", "-dlsu"}, b) }
func BenchmarkPassGenS1000(b *testing.B)     { benchmarkPassGen([]string{"1000", "-s"}, b) }
func BenchmarkPassGenSL1000(b *testing.B)    { benchmarkPassGen([]string{"1000", "-sl"}, b) }
func BenchmarkPassGenSU1000(b *testing.B)    { benchmarkPassGen([]string{"1000", "-su"}, b) }
func BenchmarkPassGenSLU1000(b *testing.B)   { benchmarkPassGen([]string{"1000", "-slu"}, b) }
func BenchmarkPassGenL1000(b *testing.B)     { benchmarkPassGen([]string{"1000", "-l"}, b) }
func BenchmarkPassGenLU1000(b *testing.B)    { benchmarkPassGen([]string{"1000", "-lu"}, b) }
func BenchmarkPassGenU1000(b *testing.B)     { benchmarkPassGen([]string{"1000", "-u"}, b) }
func BenchmarkPassGenD10000(b *testing.B)    { benchmarkPassGen([]string{"10000", "-d"}, b) }
func BenchmarkPassGenDL10000(b *testing.B)   { benchmarkPassGen([]string{"10000", "-dl"}, b) }
func BenchmarkPassGenDU10000(b *testing.B)   { benchmarkPassGen([]string{"10000", "-du"}, b) }
func BenchmarkPassGenDS10000(b *testing.B)   { benchmarkPassGen([]string{"10000", "-ds"}, b) }
func BenchmarkPassGenDLU10000(b *testing.B)  { benchmarkPassGen([]string{"10000", "-dlu"}, b) }
func BenchmarkPassGenDLS10000(b *testing.B)  { benchmarkPassGen([]string{"10000", "-dls"}, b) }
func BenchmarkPassGenDSU10000(b *testing.B)  { benchmarkPassGen([]string{"10000", "-dsu"}, b) }
func BenchmarkPassGenDLSU10000(b *testing.B) { benchmarkPassGen([]string{"10000", "-dlsu"}, b) }
func BenchmarkPassGenS10000(b *testing.B)    { benchmarkPassGen([]string{"10000", "-s"}, b) }
func BenchmarkPassGenSL10000(b *testing.B)   { benchmarkPassGen([]string{"10000", "-sl"}, b) }
func BenchmarkPassGenSU10000(b *testing.B)   { benchmarkPassGen([]string{"10000", "-su"}, b) }
func BenchmarkPassGenSLU10000(b *testing.B)  { benchmarkPassGen([]string{"10000", "-slu"}, b) }
func BenchmarkPassGenL10000(b *testing.B)    { benchmarkPassGen([]string{"10000", "-l"}, b) }
func BenchmarkPassGenLU10000(b *testing.B)   { benchmarkPassGen([]string{"10000", "-lu"}, b) }
func BenchmarkPassGenU10000(b *testing.B)    { benchmarkPassGen([]string{"10000", "-u"}, b) }
