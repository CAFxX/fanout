package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }
type GenS2 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS3 struct { F0 bool }
func genFn0(p0 int8) uint8 {
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(p0, p0)
        _ = fi7
    }
    var v8 int32 = 532
    var v9 int64 = int64(p0)
    if false {
        v8 += v8
    } else {
        var v10 bool = ("a" >= "hello")
        _ = v10
    }
    _ = p0
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 65535
    var v3 uint64 = 18446744073709551615
    var v4 int16 = -630
    var v5 bool = false
    var v6 uint32 = 65535
    switch 32767 {
    case 127:
        fmt.Println(v3)
    case 515:
        var v11 float64 = (0.0 - -2.25)
        _ = v11
    case 38:
        fmt.Println(v5, v5)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v12 := GenS1{F0: 69, F1: 255, F2: "\n\t"}
    v4 -= v4
    v2 = (v2 / (((v2 * v2)) | 1))
    for fi13 := 0; fi13 < 16; fi13++ {
        v2 -= 744
        v3 += 42
        _ = fi13
    }
    _ = len(v1)
    v14 := GenS2{F0: v2, F1: 222, F2: v6}
    v15 := GenS3{F0: v5}
    switch 71 {
    case 127:
        switch 4294967295 {
        case -1:
            switch ("a" + ("" + "")) {
            case "\n\t":
                v6 -= 480
            default:
                v16 := GenS2{F0: 790, F1: -2147483648, F2: (v6 | 465)}
                _ = v16
            }
        case 32768:
            if len(v1) > 0 {
                v1[0] = 65535
            }
        case 2147483647:
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 uint32 = v6
                _ = fi17
                _ = v18
            }
        default:
            fmt.Println(v3, v5)
        }
    }
    v4 += 32767
    v2 += (v2 % ((871) | 1))
    v6 -= (v6 * v6)
    var v19 uint = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
