package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 bool; F1 uint16 }

func main() {
    v1 := []int{504, -895, 295, -1}
    v2 := []uint32{}
    v3 := []uint32{125, 0}
    v4 := make(map[int]bool)
    var v5 uint = 42
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(v5)
        v5 = 439
        _ = fi6
    }
    v5 = 255
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v5)
    case "foo bar":
        _ = len(v4)
    case "a":
        v5 = v5
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        if (1 < v5) {
            v5 += (v5 + 1)
            v5 -= (v5 - v5)
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 uint8 = 0
                _ = len(v4)
                if (false || false) {
                    fmt.Println(fi8, fi7, fi7)
                    if ((true || true) || (32767 > 767)) {
                        v9 = 165
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v5 += (v5 ^ uint(v9))
                    }
                    fmt.Println(fi7, v9)
                }
                _ = fi8
                _ = v9
            }
        }
        _ = fi7
    }
    fmt.Println(v5)
    switch ("" + ("hello" + "")) {
    case "hello":
        v5 = v5
    case "foo bar":
        v10 := v4[32767]
        _ = v10
    case "a":
        var v11 bool = (v5 == (^(v5)))
        _ = v11
    }
    v12 := GenS1{F0: 32767}
    _ = len(v2)
    if (false && (488 >= 42)) {
        _ = len(v1)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v13 := GenS1{F0: 42}
    v1 = append(v1, -765)
    if (!true) {
        if ((float32(float64(-2.25)) / (1e10 / (((1.5 / 3.14159) + 1.5) / 3.14159))) <= 1.5) {
            _ = len(v3)
            var v14 int16 = 42
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := v4[fi15]
                for fi17 := 0; fi17 < 18; fi17++ {
                    v4[(^(fi17))] = (!(true && (255 == 389)))
                    v18 := GenS2{F0: (("" + "") + "a"), F1: 2}
                    switch "" {
                    case "hello":
                        v19 := GenS1{F0: 32767}
                        _ = v19
                    case "":
                        _ = len(v3)
                    case "a":
                        var v20 uint8 = 97
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        } else {
            var v21 uint = v5
            v22 := v4[141]
            if false {
                var v23 int32 = 1
                _ = v23
            } else {
                v21 -= v21
                if v22 {
                    v24 := GenS3{F0: 2}
                    _ = v24
                } else {
                    if ((v5 * v21) >= v21) {
                        v25 := v4[2147483647]
                        var v26 bool = v22
                        if v26 {
                            if v26 {
                                _ = len(v4)
                                v21 += 42
                                var v27 int8 = 0
                                _ = v27
                            } else {
                                v28 := v4[32767]
                                if len(v1) > 0 {
                                    v1[0] = 318
                                }
                                fmt.Println(v28, v25)
                                _ = v28
                            }
                            v29 := GenS4{F0: -1, F1: 2}
                            _ = v29
                        } else {
                            var v30 int32 = 0
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v30
                        }
                        _ = v25
                        _ = v26
                    }
                }
                v31 := GenS5{F0: v22, F1: 506}
                _ = v31
            }
            _ = v21
            _ = v22
        }
        v5 = v5
    } else {
        v32 := v4[2147483647]
        _ = v32
    }
    if (!true) {
        var v33 uint64 = 18446744073709551615
        _ = v33
    }
    switch 32767 {
    case 128:
        fmt.Println(v5)
    }
    fmt.Println(v5)
    v5 -= (v5 << (((^(v5))) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
