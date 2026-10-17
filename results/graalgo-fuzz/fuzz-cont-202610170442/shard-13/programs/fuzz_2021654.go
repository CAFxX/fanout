package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 string; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int16; F2 int16 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint }
func genFn0(p0 bool) bool {
    p0 = (150 >= 42)
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := []int{42, -872, (^(fi5))}
        _ = fi5
        _ = v6
    }
    var v7 int = 42
    _ = p0
    _ = v7
    return (2 >= 42)
}

func main() {
    var v1 string = ""
    v2 := []uint32{2147483648, 0, 2, 592}
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    if len(v2) > 0 {
        v2[0] = 495
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v1 = v1
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: (fi9 - -632), F1: fi9}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        if (128 == 0) {
            if (-906 > 1) {
                for fi12 := 0; fi12 < 3; fi12++ {
                    v4[(fi12 * (-(fi11)))] = ((true || true) || (4294967295 < 18446744073709551615))
                    v13 := GenS2{F0: ("foo bar" + v1), F1: 0, F2: 543}
                    _ = fi12
                    _ = v13
                }
                if (!true) {
                    if true {
                        v1 = ((v1 + (v1 + v1)) + (v1 + "foo bar"))
                        for fi14 := 0; fi14 < 16; fi14++ {
                            var v15 uint = 2
                            var v16 uint32 = 2147483648
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    } else {
                        var v17 int16 = -32768
                        v17 = v17
                        if false {
                            v17 = (-(v17))
                        }
                        _ = v17
                    }
                    if true {
                        switch fi11 {
                        case -618:
                            if (((!true) && (!(-1 > 0))) && true) {
                                if (true && true) {
                                    v18 := GenS3{F0: 1, F1: 1000, F2: 1000}
                                    v1 = "a"
                                    _ = v18
                                }
                                fmt.Println(v1, fi11)
                            } else {
                                var v19 uint64 = 504
                                switch (("" + v1) + (("a" + v1) + "\n\t")) {
                                case "\n\t":
                                    v1 = (v1 + v1)
                                }
                                _ = v19
                            }
                        case 42:
                            _ = len(v4)
                        }
                        var v20 uint = 312
                        _ = v20
                    } else {
                        var v21 int8 = 1
                        var v22 bool = true
                        _ = v21
                        _ = v22
                    }
                }
            } else {
                var v23 uint16 = 2
                switch 32767 {
                case 127:
                    v1 = v1
                default:
                    v23 = 42
                }
                v24 := v3[v1]
                _ = v23
                _ = v24
            }
        }
        _ = fi11
    }
    var v25 int = 116
    var v26 int32 = 1
    v25 = v25
    var v27 uint = 65535
    v1 = (((v1 + (v1 + (v1 + v1))) + (v1 + v1)) + v1)
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 int32 = v26
        v30 := GenS4{F0: 370, F1: (v27 & v27)}
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := GenS5{F0: 81, F1: 0, F2: (v27 / (((v27 * 255)) | 1))}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
}
