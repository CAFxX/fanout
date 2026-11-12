package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint32; F2 int16 }

func main() {
    var v1 uint64 = 42
    var v2 uint32 = 2147483648
    var v3 int32 = 2147483647
    v4 := make(map[uint32]bool)
    var v5 uint64 = (^(v1))
    var v6 int32 = (v3 | -573)
    if ((0.0 >= 1.5) || true) {
        fmt.Println(v5, v6, v6)
        fmt.Println(v5)
    } else {
        v7 := GenS1{F0: v2, F1: 0}
        _ = v7
    }
    v2 = v2
    if true {
        _ = len(v4)
        var v8 int = 36
        _ = v8
    }
    var v9 int64 = -1
    v3 += (v6 / ((int32(v1)) | 1))
    switch "a" {
    case "a":
        if true {
            fmt.Println(v2, v9)
            if (32767 == 42) {
                v1 = uint64(v2)
                v1 = (v1 & (v1 ^ 0))
            } else {
                switch "" {
                case "foo bar":
                    for fi10 := 0; fi10 < 9; fi10++ {
                        switch "\n\t" {
                        case "hello":
                            switch "\n\t" {
                            case "foo bar":
                                var v11 uint = 18446744073709551615
                                _ = v11
                            case "hello":
                                v2 -= (v2 / ((v2) | 1))
                            }
                        case "a":
                            v12 := GenS1{F0: v2, F1: (fi10 ^ fi10)}
                            _ = v12
                        case "":
                            v13 := []int{(fi10 / (((fi10 % ((fi10) | 1))) | 1))}
                            _ = v13
                        default:
                            v14 := GenS2{F0: -1, F1: -218}
                            _ = v14
                        }
                        switch (fi10 >> ((2) & 63)) {
                        case 865:
                            var v15 int16 = 704
                            _ = v15
                        case 32768:
                            var v16 string = ("\n\t" + "foo bar")
                            _ = v16
                        }
                        _ = fi10
                    }
                case "a":
                    for fi17 := 0; fi17 < 3; fi17++ {
                        if true {
                            v18 := v4[(v2 << ((1) & 63))]
                            if (false || (false && (155 == v2))) {
                                fmt.Println(v5)
                            }
                            _ = v18
                        } else {
                            var v19 float64 = 1.5
                            if (!(!("a" == "foo bar"))) {
                                v20 := GenS3{F0: v9, F1: 0}
                                switch 601 {
                                case -795:
                                    v21 := GenS3{F0: v9, F1: 32767}
                                    _ = v21
                                case 128:
                                    v22 := GenS4{F0: v5, F1: 2}
                                    _ = v22
                                default:
                                    v1 = v1
                                }
                                fmt.Println(v1, fi17, v9)
                                _ = v20
                            }
                            _ = v19
                        }
                        v1 = 2
                        fmt.Println(fi17, v5, v2)
                        _ = fi17
                    }
                case "\n\t":
                    v23 := []int{}
                    _ = v23
                default:
                    var v24 uint16 = 351
                    _ = v24
                }
                var v25 int32 = v6
                _ = v25
            }
            v9 = -567
        }
    case "foo bar":
        if false {
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 uint8 = 247
                _ = fi26
                _ = v27
            }
            v2 = (v2 * (v2 / (((v2 >> ((4294967295) & 63))) | 1)))
            var v28 float32 = 0.0
            _ = v28
        } else {
            for fi29 := 0; fi29 < 8; fi29++ {
                switch "foo bar" {
                case "hello":
                    if (32767 <= 0) {
                        if ((true || ((v3 + v6) != -1)) && true) {
                            v1 = (v1 >> ((42) & 63))
                        }
                    } else {
                        v2 += 2147483648
                        v5 += 18446744073709551615
                    }
                case "foo bar":
                    for fi30 := 0; fi30 < 19; fi30++ {
                        v2 = v2
                        v31 := GenS5{F0: (9223372036854775807 > (fi29 / (((fi29 | (fi29 + (fi29 >> ((620) & 63))))) | 1)))}
                        _ = fi30
                        _ = v31
                    }
                }
                fmt.Println(v6, v3, v3)
                var v32 int64 = (v9 << ((0) & 63))
                _ = fi29
                _ = v32
            }
            var v33 uint32 = (v2 | 42)
            _ = v33
        }
    case "":
        for fi34 := 0; fi34 < 3; fi34++ {
            for fi35 := 0; fi35 < 19; fi35++ {
                var v36 uint64 = 0
                v6 = (^(v3))
                _ = fi35
                _ = v36
            }
            var v37 uint64 = (v5 & (v5 ^ 1))
            _ = fi34
            _ = v37
        }
    }
    if true {
        var v38 float32 = ((float32(((-2.25 / (1e10 - 3.14159)) / 0.1)) + 1.5) * 3.14159)
        _ = v38
    }
    v39 := GenS5{F0: false}
    v40 := GenS6{F0: ((("\n\t" + "") + "a") + "hello"), F1: v2, F2: -32768}
    delete(v4, 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
