package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint; F1 bool; F2 uint }
type GenS3 struct { F0 bool; F1 int16; F2 string }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int32) int {
    v7 := GenS1{F0: 1, F1: false, F2: 255}
    fmt.Println(p0)
    var v8 int = 127
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: 0, F1: false, F2: 516}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 953
}
func genFn1() int32 {
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int16 = 32767
        v13 := GenS2{F0: 42, F1: (false && (true && true)), F2: 2}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float64 = 1.5
    _ = v14
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -54
    v3 := []string{("foo bar" + ((("foo bar" + "foo bar") + "a") + "\n\t"))}
    v4 := make(map[uint32]int)
    var v5 float32 = 1.5
    var v6 uint8 = 128
    switch "" {
    case "hello":
        for fi15 := 0; fi15 < 7; fi15++ {
            if false {
                for fi16 := 0; fi16 < 16; fi16++ {
                    if (42 != -791) {
                        var v17 int16 = -995
                        v5 = ((3.14159 / float32(0.1)) / 3.14159)
                        v5 -= (((0.0 * 1.5) + (v5 * v5)) / float32(3.14159))
                        _ = v17
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            var v18 float64 = (0.0 / 1e10)
                            _ = v18
                        case "hello":
                            fmt.Println(fi15, v5)
                        default:
                            v19 := GenS2{F0: 912, F1: (!(4294967295 >= 18446744073709551615)), F2: 0}
                            _ = v19
                        }
                        v6 = v6
                    }
                    v20 := GenS3{F0: true, F1: -15, F2: "a"}
                    if (true || (!false)) {
                        v2 = 113
                        v2 += v2
                        if true {
                            fmt.Println(v6, fi15)
                        }
                    } else {
                        if (42 > 102) {
                            switch fi16 {
                            case -9223372036854775808:
                                v6 = v6
                            case 32767:
                                v21 := GenS4{F0: (v6 | v6)}
                                _ = v21
                            case 4294967295:
                                v6 = 78
                            }
                            if true {
                                var v22 int32 = -1
                                var v23 bool = ((fi15 / (((fi15 - fi16)) | 1)) > fi16)
                                _ = v22
                                _ = v23
                            }
                        }
                        var v24 float64 = 0.1
                        v25 := GenS4{F0: (v6 % ((2) | 1))}
                        _ = v24
                        _ = v25
                    }
                    v26 := v1[fi15]
                    _ = fi16
                    _ = v20
                    _ = v26
                }
                var v27 float32 = float32(0.0)
                v27 = 1e10
                _ = v27
            }
            v5 -= v5
            _ = fi15
        }
    case "a":
        if true {
            _ = len(v3)
        }
    }
    v6 = (^(v6))
    fmt.Println(v5, v6, v2)
    v28 := GenS1{F0: 65535, F1: ((0.1 / float64(v5)) != ((-2.25 * 1e10) + ((-2.25 / 1.0) * 1e10))), F2: 0}
    for fi29 := 0; fi29 < 16; fi29++ {
        if true {
            switch genFn0(-291) {
            case 594:
                v5 = 0.0
            }
            v6 += 128
            v2 = v2
        } else {
            delete(v1, fi29)
            v6 = v6
            var v30 string = ""
            _ = v30
        }
        _ = fi29
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (1.5 == -2.25) {
        v6 = v6
    }
    var v31 uint8 = (v6 + v6)
    fmt.Println(v6, v31, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
