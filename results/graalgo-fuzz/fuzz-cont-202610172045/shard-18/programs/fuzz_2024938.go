package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint }
func genFn0(p0 string, p1 int32, p2 uint64) int8 {
    p0 = p0
    v6 := []int{460, 4294967295, 127}
    if true {
        if len(v6) > 0 {
            v6[0] = 128
        }
        p1 -= 2147483647
        p0 = (p0 + (p0 + (p0 + p0)))
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        p0 = p0
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v8 := GenS1{F0: 42}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]uint32)
    v3 := []int{1, 0, 9223372036854775807}
    v4 := make(map[uint32]uint32)
    var v5 int64 = 9223372036854775807
    v1 = v1
    var v9 int64 = 42
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 float64 = 1e10
        var v12 uint16 = 2
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int8 = -40
    for fi14 := 0; fi14 < 19; fi14++ {
        if ((v13 - (-(v13))) > v13) {
            if (127 != 0) {
                v9 += (v9 / (((v9 - v9)) | 1))
            } else {
                if false {
                    _ = len(v2)
                }
                v2[fi14] = 0
            }
            v9 = 272
        } else {
            fmt.Println(fi14)
        }
        v5 += 42
        _ = fi14
    }
    fmt.Println(v5, v9, v13)
    var v15 uint = 18446744073709551615
    _ = len(v3)
    switch 9223372036854775807 {
    case -942:
        v16 := GenS2{F0: -1, F1: 875}
        _ = v16
    case 32768:
        v15 -= v15
    case 9223372036854775807:
        switch "" {
        case "":
            var v17 uint = 2
            _ = v17
        case "a":
            var v18 float64 = -2.25
            _ = v18
        case "\n\t":
            v19 := GenS1{F0: 0}
            _ = v19
        }
    default:
        var v20 uint8 = 127
        _ = v20
    }
    v21 := GenS1{F0: 42}
    v22 := GenS2{F0: 1000, F1: 1}
    fmt.Println(v5, v1, v13)
    v23 := GenS3{F0: v15, F1: (v15 / (((v15 / ((v15) | 1))) | 1))}
    fmt.Println(v15, v15)
    var v24 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
