package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS5 struct { F0 string; F1 uint32 }
func genFn0() int16 {
    var v5 int8 = -1
    var v6 string = ("foo bar" + "hello")
    _ = v5
    _ = v6
    return -32768
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 617
    var v3 uint16 = 0
    var v4 uint16 = (v3 * v3)
    var v7 uint8 = 60
    var v8 uint32 = 1
    v2 = -32768
    var v9 int64 = 42
    if (1e10 < -2.25) {
        _ = len(v1)
        v2 = (v2 | (v2 / (((^(v2))) | 1)))
    } else {
        var v10 uint32 = v8
        if true {
            switch -252 {
            case 4294967295:
                var v11 uint8 = v7
                _ = v11
            case -679:
                for fi12 := 0; fi12 < 10; fi12++ {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        if true {
                            if (!(-68 != -6)) {
                                v14 := GenS1{F0: false}
                                _ = v14
                            } else {
                                var v15 uint64 = 4294967295
                                v16 := GenS1{F0: false}
                                _ = v15
                                _ = v16
                            }
                            v17 := GenS1{F0: ((fi13 != (fi12 / ((-774) | 1))) || true)}
                            var v18 string = ("\n\t" + "")
                            _ = v17
                            _ = v18
                        }
                        _ = fi13
                    }
                    if false {
                        for fi19 := 0; fi19 < 7; fi19++ {
                            fmt.Println(v9)
                            _ = fi19
                        }
                        for fi20 := 0; fi20 < 18; fi20++ {
                            var v21 uint64 = 93
                            _ = fi20
                            _ = v21
                        }
                    }
                    var v22 int8 = 100
                    _ = fi12
                    _ = v22
                }
            case 127:
                for fi23 := 0; fi23 < 9; fi23++ {
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v7 += 236
                        fmt.Println(fi24, v3, v3)
                        var v25 int32 = -1
                        switch "a" {
                        case "hello":
                            fmt.Println(fi24)
                        default:
                            var v26 uint = 255
                            _ = v26
                        }
                        _ = fi24
                        _ = v25
                    }
                    if true {
                        var v27 uint32 = v8
                        switch (^(fi23)) {
                        case 723:
                            v28 := GenS1{F0: false}
                            _ = v28
                        case 9223372036854775807:
                            var v29 int64 = 42
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 15; fi30++ {
                            var v31 int64 = (v9 / ((0) | 1))
                            var v32 float64 = ((float64(0.1) / float64(1e10)) * float64(1.5))
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        _ = v27
                    } else {
                        fmt.Println(v10)
                    }
                    _ = fi23
                }
            default:
                v2 -= genFn0()
            }
        }
        _ = v10
    }
    v3 = (v4 & (v3 - v3))
    var v33 bool = (false && ((v7 >> ((18446744073709551615) & 63)) < 42))
    v34 := []int{-668, 128}
    var v35 uint = 0
    v36 := GenS2{F0: -32, F1: -32768, F2: 42}
    v37 := GenS3{F0: 2, F1: v8}
    v38 := GenS4{F0: (v7 * (v7 * uint8(v35))), F1: (v7 ^ (v7 ^ 0)), F2: (v8 & (v8 | (v8 * v8)))}
    v39 := GenS5{F0: "\n\t", F1: (v8 - (^(v8)))}
    _ = len(v1)
    if (128 != (v7 >> ((1) & 63))) {
        var v40 int8 = 100
        _ = v40
    } else {
        for fi41 := 0; fi41 < 9; fi41++ {
            v42 := GenS2{F0: -1, F1: (v2 << (((v35 / ((255) | 1))) & 63)), F2: 4294967295}
            var v43 int16 = 32767
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    var v44 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v44
    _ = v7
    _ = v8
    _ = v9
}
