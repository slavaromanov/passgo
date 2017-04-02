package passgo

import "testing"

func benchmarkNewPass(flags []string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewPass(flags)
	}
}

func BenchmarkNewPassD10(b *testing.B)       { benchmarkNewPass([]string{"10", "-d"}, b) }
func BenchmarkNewPassD100(b *testing.B)      { benchmarkNewPass([]string{"100", "-d"}, b) }
func BenchmarkNewPassDL100(b *testing.B)     { benchmarkNewPass([]string{"100", "-dl"}, b) }
func BenchmarkNewPassDU100(b *testing.B)     { benchmarkNewPass([]string{"100", "-du"}, b) }
func BenchmarkNewPassDS100(b *testing.B)     { benchmarkNewPass([]string{"100", "-ds"}, b) }
func BenchmarkNewPassDLU100(b *testing.B)    { benchmarkNewPass([]string{"100", "-dlu"}, b) }
func BenchmarkNewPassDLS100(b *testing.B)    { benchmarkNewPass([]string{"100", "-dls"}, b) }
func BenchmarkNewPassDSU100(b *testing.B)    { benchmarkNewPass([]string{"100", "-dsu"}, b) }
func BenchmarkNewPassDLSU100(b *testing.B)   { benchmarkNewPass([]string{"100", "-dlsu"}, b) }
func BenchmarkNewPassS100(b *testing.B)      { benchmarkNewPass([]string{"100", "-s"}, b) }
func BenchmarkNewPassSL100(b *testing.B)     { benchmarkNewPass([]string{"100", "-sl"}, b) }
func BenchmarkNewPassSU100(b *testing.B)     { benchmarkNewPass([]string{"100", "-su"}, b) }
func BenchmarkNewPassSLU100(b *testing.B)    { benchmarkNewPass([]string{"100", "-slu"}, b) }
func BenchmarkNewPassL100(b *testing.B)      { benchmarkNewPass([]string{"100", "-l"}, b) }
func BenchmarkNewPassLU100(b *testing.B)     { benchmarkNewPass([]string{"100", "-lu"}, b) }
func BenchmarkNewPassU100(b *testing.B)      { benchmarkNewPass([]string{"100", "-u"}, b) }
func BenchmarkNewPassD1000(b *testing.B)     { benchmarkNewPass([]string{"1000", "-d"}, b) }
func BenchmarkNewPassDL1000(b *testing.B)    { benchmarkNewPass([]string{"1000", "-dl"}, b) }
func BenchmarkNewPassDU1000(b *testing.B)    { benchmarkNewPass([]string{"1000", "-du"}, b) }
func BenchmarkNewPassDS1000(b *testing.B)    { benchmarkNewPass([]string{"1000", "-ds"}, b) }
func BenchmarkNewPassDLU1000(b *testing.B)   { benchmarkNewPass([]string{"1000", "-dlu"}, b) }
func BenchmarkNewPassDLS1000(b *testing.B)   { benchmarkNewPass([]string{"1000", "-dls"}, b) }
func BenchmarkNewPassDSU1000(b *testing.B)   { benchmarkNewPass([]string{"1000", "-dsu"}, b) }
func BenchmarkNewPassDLSU1000(b *testing.B)  { benchmarkNewPass([]string{"1000", "-dlsu"}, b) }
func BenchmarkNewPassS1000(b *testing.B)     { benchmarkNewPass([]string{"1000", "-s"}, b) }
func BenchmarkNewPassSL1000(b *testing.B)    { benchmarkNewPass([]string{"1000", "-sl"}, b) }
func BenchmarkNewPassSU1000(b *testing.B)    { benchmarkNewPass([]string{"1000", "-su"}, b) }
func BenchmarkNewPassSLU1000(b *testing.B)   { benchmarkNewPass([]string{"1000", "-slu"}, b) }
func BenchmarkNewPassL1000(b *testing.B)     { benchmarkNewPass([]string{"1000", "-l"}, b) }
func BenchmarkNewPassLU1000(b *testing.B)    { benchmarkNewPass([]string{"1000", "-lu"}, b) }
func BenchmarkNewPassU1000(b *testing.B)     { benchmarkNewPass([]string{"1000", "-u"}, b) }
func BenchmarkNewPassD10000(b *testing.B)    { benchmarkNewPass([]string{"10000", "-d"}, b) }
func BenchmarkNewPassDL10000(b *testing.B)   { benchmarkNewPass([]string{"10000", "-dl"}, b) }
func BenchmarkNewPassDU10000(b *testing.B)   { benchmarkNewPass([]string{"10000", "-du"}, b) }
func BenchmarkNewPassDS10000(b *testing.B)   { benchmarkNewPass([]string{"10000", "-ds"}, b) }
func BenchmarkNewPassDLU10000(b *testing.B)  { benchmarkNewPass([]string{"10000", "-dlu"}, b) }
func BenchmarkNewPassDLS10000(b *testing.B)  { benchmarkNewPass([]string{"10000", "-dls"}, b) }
func BenchmarkNewPassDSU10000(b *testing.B)  { benchmarkNewPass([]string{"10000", "-dsu"}, b) }
func BenchmarkNewPassDLSU10000(b *testing.B) { benchmarkNewPass([]string{"10000", "-dlsu"}, b) }
func BenchmarkNewPassS10000(b *testing.B)    { benchmarkNewPass([]string{"10000", "-s"}, b) }
func BenchmarkNewPassSL10000(b *testing.B)   { benchmarkNewPass([]string{"10000", "-sl"}, b) }
func BenchmarkNewPassSU10000(b *testing.B)   { benchmarkNewPass([]string{"10000", "-su"}, b) }
func BenchmarkNewPassSLU10000(b *testing.B)  { benchmarkNewPass([]string{"10000", "-slu"}, b) }
func BenchmarkNewPassL10000(b *testing.B)    { benchmarkNewPass([]string{"10000", "-l"}, b) }
func BenchmarkNewPassLU10000(b *testing.B)   { benchmarkNewPass([]string{"10000", "-lu"}, b) }
func BenchmarkNewPassU10000(b *testing.B)    { benchmarkNewPass([]string{"10000", "-u"}, b) }
