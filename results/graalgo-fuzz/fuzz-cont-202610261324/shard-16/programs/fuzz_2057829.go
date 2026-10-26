package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint }
func genFn0() uint64 {
    var v5 int32 = 0
    v6 := GenS1{F0: (v5 * v5), F1: 1}
    v5 = 1
    _ = v5
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 int64, p1 int16) string {
    v7 := []int{9223372036854775807, 128, 32768, int(p1)}
    p0 -= p0
    _ = len(v7)
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    return ("foo bar" + ("hello" + "foo bar"))
}

func main() {
    v1 := []uint32{65535, 4294967295, 0, 246}
    var v2 int8 = 100
    v3 := []uint32{}
    v4 := make(map[int]uint32)
    fmt.Println(v2)
    for fi8 := 0; fi8 < 12; fi8++ {
        v2 -= (-(v2))
        _ = fi8
    }
    if (1e10 == (1e10 + 1.5)) {
        v2 = (v2 + v2)
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println(v2)
    if (float32((float64(-2.25) / (-2.25 / 1.5))) >= -2.25) {
        fmt.Println(v2)
    }
    fmt.Println(v2)
    _ = len(v1)
    for fi9 := 0; fi9 < 7; fi9++ {
        if false {
            if (127 != (v2 * (v2 % ((int8(fi9)) | 1)))) {
                for fi10 := 0; fi10 < 13; fi10++ {
                    switch "a" {
                    case "foo bar":
                        if true {
                            if true {
                                v11 := GenS2{F0: true, F1: 65535, F2: false}
                                var v12 uint16 = 615
                                var v13 int = fi10
                                switch "" {
                                case "a":
                                    fmt.Println(v2, fi9, v13)
                                case "":
                                    var v14 uint16 = (v12 | (v12 % (((^(v12))) | 1)))
                                    _ = v14
                                case "hello":
                                    var v15 float32 = (-2.25 + -2.25)
                                    _ = v15
                                }
                                v16 := GenS3{F0: -9223372036854775808}
                                _ = v11
                                _ = v12
                                _ = v13
                                _ = v16
                            } else {
                                v2 = (v2 ^ v2)
                            }
                            v3 = append(v3, 42)
                        }
                    case "\n\t":
                        if ((4294967295 < 65535) || (2 < 2)) {
                            fmt.Println(fi10, v2)
                            v17 := v4[fi10]
                            fmt.Println(v17, v2, v2)
                            _ = v17
                        }
                    default:
                        switch 0 {
                        case 1:
                            var v18 int64 = 0
                            _ = v18
                        case 32768:
                            if (128 <= 1) {
                                if true {
                                    switch ("\n\t" + "hello") {
                                    case "a":
                                        v2 += (v2 / ((v2) | 1))
                                    case "\n\t":
                                        v2 -= (v2 + v2)
                                    case "foo bar":
                                        switch 42 {
                                        case 42:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case -1:
                                            if false {
                                                v19 := GenS4{F0: (!(false && false))}
                                                _ = v19
                                            }
                                        case 32767:
                                            switch "a" {
                                            case "hello":
                                                var v20 int32 = -2147483648
                                                _ = v20
                                            }
                                        default:
                                            v2 -= v2
                                        }
                                    }
                                    var v21 int32 = -2147483648
                                    switch (fi9 << ((2) & 63)) {
                                    case 127:
                                        var v22 uint8 = 255
                                        var v23 int64 = 9223372036854775807
                                        _ = v22
                                        _ = v23
                                    case 42:
                                        _ = len(v4)
                                    }
                                    _ = v21
                                } else {
                                    v24 := GenS5{F0: -286, F1: 42}
                                    _ = v24
                                }
                            } else {
                                if ((3.14159 - float64(0.0)) < -2.25) {
                                    if (2 <= 177) {
                                        switch 24 {
                                        case 822:
                                            v2 = v2
                                        }
                                        var v25 int64 = 42
                                        switch (fi9 << ((1) & 63)) {
                                        case 127:
                                            v2 = (^(v2))
                                        case -2147483648:
                                            var v26 float32 = 1e10
                                            _ = v26
                                        default:
                                            v27 := v4[(fi9 / (((fi10 ^ fi9)) | 1))]
                                            _ = v27
                                        }
                                        _ = v25
                                    }
                                    fmt.Println(v2)
                                }
                            }
                        }
                    }
                    _ = fi10
                }
            } else {
                var v28 float32 = (1e10 - float32(3.14159))
                _ = v28
            }
            switch fi9 {
            case 425:
                fmt.Println(fi9, v2)
            default:
                for fi29 := 0; fi29 < 20; fi29++ {
                    fmt.Println(v2, fi29)
                    v4[-544] = 2147483648
                    fmt.Println(v2, fi29)
                    _ = fi29
                }
            }
        } else {
            var v30 uint64 = 272
            fmt.Println(fi9)
            var v31 float32 = (float32(0.1) - (float32((1e10 + -2.25)) - (0.1 / 1.5)))
            _ = v30
            _ = v31
        }
        v32 := GenS6{F0: 307}
        var v33 int64 = 0
        _ = fi9
        _ = v32
        _ = v33
    }
    var v34 uint16 = 42
    fmt.Println(v34, v2)
    var v35 uint64 = 18446744073709551615
    var v36 bool = ((1 >= -9223372036854775808) || (false && true))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
