package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := []uint32{}
    var v2 int = -132
    var v3 int16 = 0
    var v4 float32 = 0.1
    var v5 uint = 2
    v6 := []int{448, (v2 << (((v5 >> ((v5) & 63))) & 63)), (v2 | v2)}
    var v7 float64 = -2.25
    if true {
        for fi8 := 0; fi8 < 8; fi8++ {
            if true {
                v5 -= (v5 * v5)
                var v9 uint16 = 32767
                switch "hello" {
                case "hello":
                    v10 := GenS1{F0: -32768}
                    _ = v10
                case "":
                    for fi11 := 0; fi11 < 8; fi11++ {
                        if (!false) {
                            v4 += v4
                            v7 = (v7 / v7)
                            var v12 float32 = v4
                            _ = v12
                        } else {
                            fmt.Println(v2, v7, v2)
                            if (v9 >= 1) {
                                v9 = v9
                                switch fi11 {
                                case 1:
                                    v13 := GenS1{F0: (v3 >> (((^(v5))) & 63))}
                                    _ = v13
                                case 4294967295:
                                    if len(v6) > 0 {
                                        v6[0] = -949
                                    }
                                default:
                                    if (0 >= 1) {
                                        if true {
                                            var v14 uint16 = v9
                                            _ = v14
                                        }
                                        var v15 int = 4294967295
                                        fmt.Println(fi11, v7, v9)
                                        _ = v15
                                    } else {
                                        v16 := GenS1{F0: (v3 - (v3 >> (((v5 & (v5 >> (((v5 / ((v5) | 1))) & 63)))) & 63)))}
                                        _ = v16
                                    }
                                }
                            } else {
                                var v17 float64 = float64((v4 * (v4 * v4)))
                                _ = v17
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v18 float32 = 1.5
                        fmt.Println(v3, v9)
                        _ = fi11
                        _ = v18
                    }
                case "\n\t":
                    for fi19 := 0; fi19 < 16; fi19++ {
                        fmt.Println(v4, v9)
                        v20 := GenS1{F0: -1}
                        _ = fi19
                        _ = v20
                    }
                }
                _ = v9
            }
            v4 = (v4 / (float32(float64(v4)) - 1.5))
            v21 := GenS2{F0: 65535, F1: v3, F2: 4294967295}
            _ = fi8
            _ = v21
        }
    } else {
        if true {
            if false {
                v22 := GenS3{F0: 0}
                v4 = 0.0
                _ = v22
            } else {
                var v23 float64 = -2.25
                _ = v23
            }
            v3 -= v3
        } else {
            v7 -= float64(0.1)
        }
    }
    _ = len(v6)
    for fi24 := 0; fi24 < 10; fi24++ {
        switch (("hello" + "\n\t") + ("hello" + ("" + ""))) {
        case "a":
            v4 += ((v4 - v4) / 0.1)
        case "hello":
            switch v2 {
            case 127:
                fmt.Println(fi24, v5)
            case 4294967295:
                v25 := GenS2{F0: 1, F1: v3, F2: 18446744073709551615}
                _ = v25
            }
        }
        _ = fi24
    }
    if true {
        v26 := GenS4{F0: 118, F1: 9223372036854775807}
        switch (v2 / ((v2) | 1)) {
        case 272:
            var v27 int32 = 2147483647
            _ = v27
        case 806:
            var v28 uint64 = 1
            _ = v28
        case 4294967295:
            var v29 int32 = 532
            _ = v29
        default:
            var v30 int = (v2 - v2)
            _ = v30
        }
        _ = v26
    } else {
        v5 = 4294967295
        for fi31 := 0; fi31 < 3; fi31++ {
            switch "foo bar" {
            case "\n\t":
                v2 = 765
            case "hello":
                v2 = v2
            case "foo bar":
                var v32 int16 = v3
                _ = v32
            }
            v33 := GenS4{F0: 1, F1: 887}
            _ = fi31
            _ = v33
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v34 := GenS5{F0: -66}
    var v35 uint32 = 930
    var v36 int8 = 127
    v37 := GenS5{F0: (v3 - v3)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
