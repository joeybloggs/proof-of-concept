package allocations

import "testing"

// func BenchmarkReflectString(b *testing.B) {

// 	s := "test"

// 	for n := 0; n < b.N; n++ {
// 		testReflectAllocations(s)
// 	}
// }

// func BenchmarkReflectFloat64(b *testing.B) {

// 	f := 1.123

// 	for n := 0; n < b.N; n++ {
// 		testReflectAllocations(f)
// 	}
// }

// func BenchmarkReflectStringPtr(b *testing.B) {

// 	var s *string
// 	tmp := "test"
// 	s = &tmp

// 	for n := 0; n < b.N; n++ {
// 		testReflectAllocations(s)
// 	}
// }

// func BenchmarkReflectFloat64Ptr(b *testing.B) {

// 	var f *float64
// 	tmp := 1.123
// 	f = &tmp

// 	for n := 0; n < b.N; n++ {
// 		testReflectAllocations(f)
// 	}
// }

// func BenchmarkBasicString(b *testing.B) {

// 	s := "test"

// 	for n := 0; n < b.N; n++ {
// 		testBasicAllocations(s)
// 	}
// }

// func BenchmarkBasicFloat64(b *testing.B) {

// 	f := 1.123

// 	for n := 0; n < b.N; n++ {
// 		testBasicAllocations(f)
// 	}
// }

func BenchmarkHybridString(b *testing.B) {

	s := "test"

	for n := 0; n < b.N; n++ {
		testHybridAllocations(s)
	}
}

func BenchmarkHybridFloat64(b *testing.B) {

	f := 1.123

	for n := 0; n < b.N; n++ {
		testHybridAllocations(f)
	}
}

func BenchmarkHybridStringPtr(b *testing.B) {

	var s *string
	tmp := "test"
	s = &tmp

	for n := 0; n < b.N; n++ {
		testHybridAllocations(s)
	}
}

func BenchmarkHybridFloat64Ptr(b *testing.B) {

	var f *float64
	tmp := 1.123
	f = &tmp

	for n := 0; n < b.N; n++ {
		testHybridAllocations(f)
	}
}
