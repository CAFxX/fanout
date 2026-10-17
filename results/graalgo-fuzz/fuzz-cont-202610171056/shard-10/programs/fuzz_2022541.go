package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 bool }
func genFn0(p0 int8) int64 {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := []int{}
        var v6 int64 = 158
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint64 = 4294967295
    v7 += (v7 / ((946) | 1))
    _ = p0
    _ = v7
    return 0
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{(v1 << ((1) & 63))}
    var v3 int64 = 1
    v8 := GenS1{F0: v1, F1: "\n\t", F2: 0}
    v3 += (v3 / ((42) | 1))
    var v9 int32 = 0
    for fi10 := 0; fi10 < 6; fi10++ {
        v2 = append(v2, 2147483647)
        var v11 uint8 = 255
        _ = fi10
        _ = v11
    }
    v3 = v3
    v9 -= 257
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 string = ""
            switch -1 {
            case 9223372036854775807:
                if ((v9 & 0) <= (v9 ^ v9)) {
                    v9 = (v9 % ((int32(fi12)) | 1))
                } else {
                    v15 := GenS2{F0: 883, F1: v1, F2: (v3 + (v3 | v3))}
                    v1 += v1
                    var v16 float64 = (-2.25 - 3.14159)
                    v16 = 0.0
                    v17 := GenS1{F0: (^(v1)), F1: "", F2: 306}
                    v18 := GenS2{F0: 65535, F1: v1, F2: v3}
                    v14 = (((v14 + ("" + "hello")) + "foo bar") + (v14 + ("" + "foo bar")))
                    v3 = (^(v3))
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            case -429:
                v19 := GenS1{F0: 1, F1: "\n\t", F2: 18446744073709551615}
                _ = v19
            case 128:
                if ((!(1 != v9)) && (584 < 571)) {
                    _ = len(v2)
                } else {
                    var v20 int32 = (v9 * 2147483647)
                    _ = v20
                }
            default:
                var v21 uint8 = 127
                _ = v21
            }
            _ = fi13
            _ = v14
        }
        v1 = 42
        switch (("foo bar" + (("" + "\n\t") + "hello")) + (("" + "a") + ("hello" + "hello"))) {
        case "hello":
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(v1)
                _ = fi22
            }
        case "a":
            if (1e10 > 1e10) {
                switch "a" {
                case "":
                    var v23 int = 32768
                    _ = v23
                case "\n\t":
                    for fi24 := 0; fi24 < 11; fi24++ {
                        switch (fi12 / ((32767) | 1)) {
                        case 32767:
                            v25 := GenS1{F0: 2, F1: "a", F2: 0}
                            _ = v25
                        case 32768:
                            var v26 int32 = (v9 + (-(v9)))
                            _ = v26
                        case -1:
                            var v27 string = (("a" + ("\n\t" + "")) + ("\n\t" + "hello"))
                            _ = v27
                        default:
                            switch fi12 {
                            case 128:
                                if true {
                                    fmt.Println(v9, v1)
                                    _ = len(v2)
                                    v3 = (v3 & (^(v3)))
                                }
                            case 127:
                                var v28 int16 = 0
                                _ = v28
                            default:
                                v29 := GenS2{F0: 65535, F1: v1, F2: int64(fi12)}
                                _ = v29
                            }
                        }
                        var v30 float64 = (0.0 / 1.0)
                        _ = fi24
                        _ = v30
                    }
                default:
                    var v31 uint64 = 4294967295
                    _ = v31
                }
            }
        default:
            fmt.Println(v3, v9, v1)
        }
        _ = fi12
    }
    if false {
        v32 := GenS3{F0: (v1 & v1), F1: -160, F2: 565}
        if false {
            fmt.Println(v1, v9)
            v33 := GenS4{F0: 2, F1: v1}
            _ = v33
        }
        _ = v32
    }
    var v34 uint8 = 128
    v34 += v34
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += v1
    v34 -= (v34 * 0)
    v1 -= 1
    v35 := GenS5{F0: false}
    if (!false) {
        var v36 int16 = 42
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v8
    _ = v9
}
