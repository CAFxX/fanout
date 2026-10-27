package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint32 }
func genFn0(p0 string, p1 int64, p2 bool) string {
    fmt.Println(p0, p0, p2)
    v4 := GenS1{F0: 65535, F1: -525, F2: 214}
    p2 = (-2147483648 == -602)
    var v5 int32 = -703
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return "a"
}
func genFn1(p0 int8, p1 uint32) int8 {
    v6 := GenS2{F0: 2, F1: true}
    p0 = (p0 | -1)
    p1 -= 595
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{(fi8 | fi8), (fi8 * (fi7 / (((fi7 % ((fi8) | 1))) | 1))), 128, fi8}
            _ = fi8
            _ = v9
        }
        if ((2 != 336) && false) {
            p0 += (p0 / ((0) | 1))
            p0 = (^(p0))
        }
        var v10 float32 = ((0.0 / 1e10) / 1.0)
        _ = fi7
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 uint8 = 64
    var v2 string = ""
    v3 := []string{v2, "\n\t", ("a" + v2), "a"}
    var v11 uint64 = 42
    v1 = v1
    switch ("foo bar" + v2) {
    case "":
        v12 := GenS3{F0: 309}
        _ = v12
    case "a":
        v3 = append(v3, ("foo bar" + v2))
    }
    fmt.Println(v11, v11, v1)
    var v13 uint32 = 4294967295
    switch v2 {
    case "":
        var v14 string = (v2 + "foo bar")
        _ = v14
    case "\n\t":
        fmt.Println(v2, v11, v13)
    }
    var v15 int64 = -999
    if false {
        v15 -= 42
        if len(v3) > 0 {
            v3[0] = "a"
        }
    } else {
        v11 = v11
        v11 -= 353
    }
    if (2 <= 4294967295) {
        switch ("hello" + ("foo bar" + (v2 + (v2 + v2)))) {
        case "a":
            fmt.Println(v15, v11)
        default:
            var v16 float64 = (0.0 - 1.5)
            _ = v16
        }
        v3 = append(v3, "\n\t")
    } else {
        v2 = "foo bar"
        switch 413 {
        case 9223372036854775807:
            if len(v3) > 0 {
                v3[0] = v2
            }
        }
        v15 -= v15
    }
    v13 -= 2
    for fi17 := 0; fi17 < 5; fi17++ {
        if true {
            for fi18 := 0; fi18 < 17; fi18++ {
                switch fi17 {
                case 32768:
                    var v19 uint64 = (v11 + 2)
                    _ = v19
                case -1:
                    var v20 uint = 653
                    var v21 uint16 = 911
                    _ = v20
                    _ = v21
                default:
                    v1 += v1
                }
                _ = fi18
            }
            if (false && true) {
                var v22 uint = 42
                v23 := GenS4{F0: 454, F1: 255, F2: 4294967295}
                _ = v22
                _ = v23
            } else {
                var v24 uint = 42
                var v25 string = ""
                _ = v24
                _ = v25
            }
            fmt.Println(v1)
        }
        _ = fi17
    }
    _ = len(v3)
    v3 = append(v3, (((v2 + v2) + (v2 + "\n\t")) + ("" + "\n\t")))
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 uint32 = v13
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v13
    _ = v15
    _ = v2
    _ = v3
}
