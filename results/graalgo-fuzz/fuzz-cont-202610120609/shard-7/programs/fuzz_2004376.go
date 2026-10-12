package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint) uint {
    var v5 int = -539
    if false {
        var v6 uint = 42
        _ = v6
    } else {
        for fi7 := 0; fi7 < 13; fi7++ {
            v5 = (fi7 | (-(v5)))
            var v8 int64 = 0
            var v9 float32 = 3.14159
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v10 := []int{(v5 / (((v5 ^ v5)) | 1)), 42, (v5 - v5), (v5 >> ((p0) & 63))}
    switch (v5 << (((p0 & 2)) & 63)) {
    case -553:
        var v11 uint64 = 933
        _ = v11
    case -464:
        var v12 uint64 = 18446744073709551615
        _ = v12
    case 2147483647:
        switch "hello" {
        case "":
            var v13 uint32 = 1
            _ = v13
        case "hello":
            var v14 float32 = ((0.1 + (float32((float64((0.1 * (-2.25 * 1.5))) / (1.5 + float64(((1.5 * 1.5) * -2.25))))) * 0.0)) + 0.1)
            _ = v14
        case "\n\t":
            if true {
                v10 = append(v10, (v5 << ((p0) & 63)))
            }
        }
    }
    _ = p0
    _ = v10
    _ = v5
    return 42
}

func main() {
    v1 := []int{-1, -153, 9223372036854775807}
    v2 := make(map[int]string)
    var v3 float64 = 0.1
    var v4 uint8 = 1
    var v15 int64 = 42
    if (-663 < v15) {
        v16 := GenS1{F0: 0}
        v3 = v3
        var v17 uint16 = 568
        _ = v16
        _ = v17
    } else {
        _ = len(v2)
        v1 = append(v1, 1)
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        fmt.Println(fi18)
        v19 := GenS1{F0: 1}
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: int16(v4), F1: -595}
    switch -1 {
    case -915:
        _ = len(v2)
    case 32767:
        var v21 bool = (v4 == (v4 >> ((4294967295) & 63)))
        _ = v21
    default:
        v15 -= 9223372036854775807
    }
    v3 -= (-2.25 + (-2.25 * v3))
    if (((false && true) || true) || (0 < -2147483648)) {
        v22 := GenS3{F0: -9223372036854775808, F1: 2}
        _ = len(v2)
        switch 32768 {
        case 534:
            v3 = 0.1
        case 0:
            for fi23 := 0; fi23 < 18; fi23++ {
                delete(v2, fi23)
                _ = fi23
            }
        default:
            if (100 < 127) {
                delete(v2, 2147483647)
                v3 = -2.25
            }
        }
        _ = v22
    }
    var v24 uint = genFn0(18446744073709551615)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
