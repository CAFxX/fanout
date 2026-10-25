package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int8; F2 uint8 }
func genFn0(p0 uint16, p1 int32) uint64 {
    for fi5 := 0; fi5 < 11; fi5++ {
        if false {
            p0 -= (^(p0))
            switch "\n\t" {
            case "\n\t":
                v6 := []int{417}
                _ = v6
            case "a":
                if (true || (1e10 == (((-2.25 - 0.1) - 1e10) + 0.1))) {
                    var v7 float64 = ((0.1 - 0.0) + ((1.5 / 1.0) - (0.1 / float64(1e10))))
                    var v8 int8 = 1
                    _ = v7
                    _ = v8
                }
            }
        } else {
            p0 = p0
            for fi9 := 0; fi9 < 6; fi9++ {
                if false {
                    v10 := GenS1{F0: -1, F1: 651}
                    fmt.Println(p1, fi9)
                    _ = v10
                } else {
                    if (18446744073709551615 == 508) {
                        var v11 bool = false
                        var v12 uint64 = 2
                        fmt.Println(p1, fi5, v11)
                        var v13 int8 = 0
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    switch 65 {
                    case 127:
                        v14 := []int{(fi9 | fi5), -337, fi5}
                        _ = v14
                    }
                }
                _ = fi9
            }
        }
        _ = fi5
    }
    var v15 uint8 = 0
    var v16 int64 = -205
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 int8 = 28
            if false {
                var v20 uint8 = (v15 << ((18446744073709551615) & 63))
                v21 := GenS1{F0: int8(p1), F1: 42}
                _ = v20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        var v22 int = (fi17 + fi17)
        var v23 uint = 65535
        _ = fi17
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 174
}
func genFn1(p0 string, p1 uint64, p2 uint8) bool {
    v24 := []int{2147483647, 0}
    var v25 uint8 = (p2 ^ (p2 & (p2 & (p2 * 1))))
    switch "a" {
    case "foo bar":
        fmt.Println(p1, p0)
    case "":
        var v26 float32 = 0.0
        _ = v26
    case "hello":
        v25 += 207
    default:
        var v27 uint16 = uint16(p1)
        _ = v27
    }
    var v28 uint32 = 2
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    _ = v28
    return false
}

func main() {
    var v1 int64 = -905
    var v2 uint64 = 293
    var v3 uint64 = v2
    var v4 int = 405
    v4 = -854
    var v29 int16 = 422
    v30 := []int{32768, (-(v4)), v4, v4}
    var v31 int8 = -112
    switch v4 {
    case 2147483647:
        for fi32 := 0; fi32 < 18; fi32++ {
            v33 := GenS2{F0: (v4 / (((-(v4))) | 1))}
            v34 := GenS1{F0: 42, F1: 0}
            _ = fi32
            _ = v33
            _ = v34
        }
    case 127:
        var v35 float64 = (0.1 - 0.0)
        _ = v35
    }
    var v36 string = (("" + ("foo bar" + "\n\t")) + "\n\t")
    v37 := GenS3{F0: 255, F1: (v31 + v31), F2: 255}
    v36 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v4
}
