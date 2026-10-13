package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint64; F1 int64; F2 int }

func main() {
    var v1 uint64 = 28
    var v2 int16 = -488
    v3 := make(map[int]uint32)
    v4 := []uint32{2147483648, 0}
    var v5 float32 = (((-2.25 * 3.14159) / -2.25) * 0.1)
    v6 := GenS1{F0: -2147483648, F1: uint(v1)}
    fmt.Println(v1, v1, v1)
    v5 += v5
    fmt.Println(v1, v1)
    var v7 float64 = ((((1e10 + -2.25) - 0.1) - (float64(v5) + 3.14159)) * 1.5)
    var v8 string = ""
    var v9 uint64 = v1
    fmt.Println(v8, v2)
    switch v8 {
    case "\n\t":
        v10 := GenS1{F0: -405, F1: 255}
        _ = v10
    case "foo bar":
        var v11 int64 = -996
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 int64 = 0
            _ = fi12
            _ = v13
        }
    }
    switch 127 {
    case 475:
        v2 = v2
    default:
        switch (v8 + (((((v8 + "hello") + v8) + v8) + (v8 + v8)) + (("a" + v8) + (v8 + v8)))) {
        case "":
            if (uint32(v1) != 308) {
                if true {
                    var v14 uint32 = 348
                    _ = v14
                }
            } else {
                var v15 bool = (false && false)
                var v16 int64 = 42
                v5 = 1e10
                _ = v15
                _ = v16
            }
        case "hello":
            if true {
                switch -510 {
                case 236:
                    v2 -= (-(v2))
                case -799:
                    v8 = (((v8 + v8) + v8) + "\n\t")
                case -696:
                    fmt.Println(v2)
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = 155
                }
                var v17 int8 = 44
                for fi18 := 0; fi18 < 4; fi18++ {
                    fmt.Println(v7, fi18, fi18)
                    for fi19 := 0; fi19 < 3; fi19++ {
                        switch "" {
                        case "":
                            var v20 int64 = 0
                            _ = v20
                        case "a":
                            v17 = (v17 >> ((42) & 63))
                        case "foo bar":
                            v21 := GenS1{F0: 0, F1: 501}
                            _ = v21
                        default:
                            if false {
                                var v22 bool = false
                                _ = v22
                            } else {
                                v5 += (v5 + float32(1e10))
                                switch v8 {
                                case "foo bar":
                                    v8 = v8
                                default:
                                    if true {
                                        if true {
                                            v23 := GenS1{F0: 54, F1: 2}
                                            _ = v23
                                        } else {
                                            v24 := GenS1{F0: -730, F1: 4294967295}
                                            _ = len(v3)
                                            _ = v24
                                        }
                                        var v25 float64 = ((v7 + -2.25) * float64(((v5 + ((((((((v5 * ((1.5 * 0.1) * v5)) + -2.25) + ((v5 / (v5 / 1e10)) - float32(0.1))) / ((v5 / v5) * 3.14159)) + 3.14159) / 3.14159) / ((v5 + (v5 - float32(float64(v5)))) - v5)) / (1e10 + v5))) - v5)))
                                        _ = len(v4)
                                        _ = v25
                                    } else {
                                        var v26 uint8 = 128
                                        _ = v26
                                    }
                                }
                                v27 := GenS2{F0: -2147483648, F1: 694}
                                var v28 bool = false
                                _ = v27
                                _ = v28
                            }
                        }
                        v4 = append(v4, 2147483647)
                        var v29 string = "foo bar"
                        _ = fi19
                        _ = v29
                    }
                    v30 := GenS3{F0: 1, F1: 42}
                    _ = fi18
                    _ = v30
                }
                _ = v17
            }
        case "foo bar":
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
        default:
            v31 := GenS4{F0: (v9 ^ v9), F1: 42, F2: 0}
            _ = v31
        }
    }
    v8 = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
