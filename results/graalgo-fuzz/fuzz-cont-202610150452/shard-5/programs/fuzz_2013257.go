package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint64 }
func genFn0(p0 uint8, p1 uint8) int {
    var v6 uint16 = 0
    switch ("hello" + "a") {
    case "":
        if true {
            switch 570 {
            case 908:
                p1 -= (p1 / (((^(p0))) | 1))
            case 20:
                for fi7 := 0; fi7 < 15; fi7++ {
                    p1 += (p1 << ((4294967295) & 63))
                    _ = fi7
                }
            }
        } else {
            switch "" {
            case "":
                v8 := []int{32768, 2147483647, 127}
                _ = v8
            }
        }
    case "\n\t":
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: (p1 * p0), F1: 862, F2: 189}
            _ = fi9
            _ = v10
        }
    case "foo bar":
        var v11 uint8 = 127
        _ = v11
    }
    var v12 int64 = 0
    v13 := GenS1{F0: (^(p0)), F1: 492, F2: 4294967295}
    for fi14 := 0; fi14 < 12; fi14++ {
        fmt.Println(p1, p1)
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v6
    return -1
}
func genFn1(p0 int8, p1 uint16, p2 int8) int {
    var v15 float64 = (0.0 + ((3.14159 / 1.0) / 1.5))
    var v16 int16 = -32768
    v17 := GenS2{F0: 435, F1: 1, F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v17
    return genFn0(74, 0)
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 79
    var v3 int64 = 109
    v4 := make(map[int]uint32)
    var v5 uint32 = 710
    v18 := []int{}
    var v19 uint64 = uint64(v5)
    v19 = v19
    v20 := GenS2{F0: (v5 & 1), F1: 1, F2: (^(v19))}
    var v21 uint = 42
    v22 := GenS1{F0: 178, F1: (v21 % (((v21 & v21)) | 1)), F2: 1}
    _ = len(v18)
    var v23 uint32 = v5
    v23 = (v23 + v23)
    var v24 int = 32768
    var v25 uint8 = 17
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 uint32 = 42
        var v28 int64 = (v3 % ((632) | 1))
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := v4[(v24 ^ v24)]
    fmt.Println(v19, v19, v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
