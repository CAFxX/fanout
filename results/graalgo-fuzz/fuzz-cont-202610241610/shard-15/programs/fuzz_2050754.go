package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint64; F1 bool; F2 int8 }
type GenS3 struct { F0 int16; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 string; F2 int }
func genFn0(p0 int8, p1 int8, p2 int32) uint32 {
    var v4 uint16 = 32767
    v5 := GenS1{F0: 0, F1: 172}
    var v6 uint64 = 562
    p0 = (p0 >> ((2) & 63))
    if true {
        var v7 float64 = 1e10
        var v8 float32 = -2.25
        _ = v7
        _ = v8
    } else {
        if true {
            if true {
                v9 := []int{4294967295, -1}
                p0 = (p1 % (((^(p1))) | 1))
                _ = v9
            }
            fmt.Println(p2, v6, p0)
            var v10 int64 = -1
            _ = v10
        }
        if ((!(((float64(-2.25) * 3.14159) * (3.14159 * 1e10)) != 0.0)) && true) {
            v4 -= v4
        } else {
            var v11 uint32 = 4294967295
            v12 := GenS1{F0: (v6 - v6), F1: 164}
            _ = v11
            _ = v12
        }
        v13 := []int{}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint = 2
    var v2 uint8 = 128
    v3 := make(map[int]int)
    var v14 float64 = 1e10
    if true {
        v14 = ((float64((0.0 - -2.25)) / (v14 * ((0.0 + ((v14 / 0.1) / 0.1)) - v14))) - 0.1)
        v15 := GenS2{F0: 0, F1: false, F2: -1}
        _ = v15
    } else {
        v16 := GenS2{F0: 17, F1: true, F2: 100}
        fmt.Println(v2)
        _ = v16
    }
    switch 128 {
    case 715:
        if false {
            v14 = 0.0
        } else {
            v2 = 128
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 uint32 = 65535
                var v19 int16 = -32768
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 int64 = 1
            _ = v20
        }
    default:
        var v21 uint8 = 172
        _ = v21
    }
    v22 := []int{9223372036854775807, -445, 1, -2147483648}
    v14 += (1.5 - float64(-2.25))
    v1 -= 729
    if (!true) {
        switch 32768 {
        case 1:
            if true {
                if true {
                    for fi23 := 0; fi23 < 15; fi23++ {
                        if true {
                            for fi24 := 0; fi24 < 11; fi24++ {
                                if ((((("\n\t" + ("a" + "hello")) + "foo bar") != "a") || false) || (4294967295 <= 4294967295)) {
                                    v2 += (v2 ^ (^(v2)))
                                    v25 := GenS3{F0: 42, F1: false, F2: v2}
                                    _ = v25
                                } else {
                                    switch fi24 {
                                    case -686:
                                        v1 -= v1
                                    case -9223372036854775808:
                                        var v26 uint = v1
                                        _ = v26
                                    default:
                                        var v27 uint = (v1 / ((v1) | 1))
                                        _ = v27
                                    }
                                    var v28 bool = (42 > -128)
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                    _ = v28
                                }
                                _ = len(v22)
                                var v29 float64 = v14
                                _ = fi24
                                _ = v29
                            }
                            var v30 int8 = 100
                            _ = v30
                        }
                        _ = fi23
                    }
                    fmt.Println(v1, v1, v1)
                    var v31 uint8 = (v2 - 255)
                    _ = v31
                }
            } else {
                v2 -= (^(v2))
            }
        case 128:
            var v32 uint16 = 1
            _ = v32
        }
    } else {
        var v33 int32 = int32(v1)
        v1 = 65535
        var v34 uint = 42
        _ = v33
        _ = v34
    }
    delete(v3, -9223372036854775808)
    var v35 float32 = float32(0.1)
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    var v36 string = "foo bar"
    var v37 uint8 = v2
    var v38 uint32 = 4294967295
    if true {
        v39 := GenS4{F0: 556}
        if (true || false) {
            v40 := v3[9223372036854775807]
            fmt.Println(v36, v36)
            v37 += 218
            _ = v40
        }
        _ = v39
    }
    var v41 int = 267
    v41 = (v41 * 4294967295)
    v42 := GenS5{F0: ((v14 / (v14 / v14)) != v14), F1: ("a" + (("a" + "hello") + v36)), F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v41
    _ = v42
}
