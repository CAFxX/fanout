package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint32; F2 int16 }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]bool)
    var v3 int8 = -1
    var v4 uint = 18446744073709551615
    var v5 int32 = 42
    var v6 float64 = ((3.14159 + (1e10 * 0.1)) - 1e10)
    v3 = (v3 | (v3 + v3))
    var v7 int16 = v1
    fmt.Println(v3, v7, v4)
    for fi8 := 0; fi8 < 4; fi8++ {
        if false {
            v1 -= (v7 % ((v1) | 1))
            for fi9 := 0; fi9 < 7; fi9++ {
                v6 += ((v6 * v6) + 1.5)
                switch ("foo bar" + "") {
                case "\n\t":
                    var v10 int16 = (v7 ^ (^(v7)))
                    _ = v10
                default:
                    v6 = v6
                }
                v6 = float64(((1e10 - (0.1 * 0.1)) + 0.0))
                _ = fi9
            }
        } else {
            if true {
                var v11 uint64 = 1
                _ = v11
            }
            var v12 uint16 = 32767
            _ = v12
        }
        v3 -= (v3 & v3)
        _ = fi8
    }
    if (false && false) {
        if ((false || true) && ((183 == 255) || false)) {
            for fi13 := 0; fi13 < 18; fi13++ {
                for fi14 := 0; fi14 < 17; fi14++ {
                    if (1 == 2) {
                        v4 = 0
                        var v15 int64 = -284
                        v16 := GenS1{F0: fi13}
                        _ = v15
                        _ = v16
                    }
                    v1 = (v1 << ((255) & 63))
                    fmt.Println(v5, v5, v4)
                    _ = fi14
                }
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 int64 = -22
                    if true {
                        _ = len(v2)
                    }
                    v1 = v7
                    switch ("" + "a") {
                    case "a":
                        var v19 float32 = 3.14159
                        _ = v19
                    case "hello":
                        delete(v2, 42)
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi13
            }
        }
        switch "a" {
        case "foo bar":
            var v20 int = 2147483647
            _ = v20
        default:
            for fi21 := 0; fi21 < 15; fi21++ {
                v6 = v6
                var v22 int = -769
                _ = fi21
                _ = v22
            }
        }
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 uint8 = 166
            if (!false) {
                var v25 uint8 = (v24 << (((v4 >> ((2) & 63))) & 63))
                var v26 int64 = -1
                v6 = (v6 * v6)
                _ = v25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
    } else {
        v7 -= v7
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 string = ("\n\t" + ("" + "\n\t"))
        _ = fi27
        _ = v28
    }
    switch 32768 {
    case -919:
        if (v4 > v4) {
            v5 += (v5 + v5)
            v2[621] = false
            var v29 string = ((("a" + "\n\t") + "foo bar") + ("\n\t" + "foo bar"))
            _ = v29
        }
    case 935:
        v7 += v7
    case 32767:
        v30 := GenS2{F0: int16(v5)}
        _ = v30
    default:
        delete(v2, 699)
    }
    v31 := []int{42, 128}
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 uint = (v4 / ((42) | 1))
        if false {
            var v34 uint32 = 4294967295
            if true {
                if (128 == 1) {
                    for fi35 := 0; fi35 < 7; fi35++ {
                        _ = len(v2)
                        _ = fi35
                    }
                    var v36 uint64 = 4294967295
                    _ = v36
                } else {
                    fmt.Println(fi32)
                    if (v6 > (((((0.1 / (v6 / v6)) * 1.5) * 3.14159) + v6) * v6)) {
                        v37 := GenS2{F0: (v1 & (v7 | (^(v1))))}
                        _ = v37
                    }
                }
            } else {
                v38 := GenS3{F0: (v5 + 0), F1: (v34 | (^(v34))), F2: (v1 ^ v1)}
                _ = v38
            }
            _ = v34
        }
        if false {
            if len(v31) > 0 {
                v31[0] = -266
            }
        } else {
            fmt.Println(v6, v7, v3)
            v39 := GenS2{F0: -65}
            _ = v39
        }
        _ = fi32
        _ = v33
    }
    var v40 int16 = (v1 & (v1 - v1))
    var v41 float32 = (1e10 - (0.0 / float32((v6 / v6))))
    var v42 uint64 = 42
    if (-2.25 == v6) {
        if ((255 == 1) || (v41 == (1e10 + v41))) {
            fmt.Println(v42, v5)
        } else {
            _ = len(v31)
        }
        var v43 float32 = -2.25
        _ = v43
    } else {
        var v44 uint8 = 213
        switch "hello" {
        case "hello":
            v1 += v1
        case "":
            _ = len(v31)
        default:
            for fi45 := 0; fi45 < 9; fi45++ {
                _ = len(v2)
                var v46 bool = ((v5 ^ v5) == -590)
                fmt.Println(v3)
                _ = fi45
                _ = v46
            }
        }
        var v47 uint8 = (v44 - (v44 % (((^(v44))) | 1)))
        _ = v44
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
