package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int16; F2 bool }
func genFn0(p0 bool) int64 {
    var v4 int64 = 9223372036854775807
    p0 = (p0 && (("a" + "\n\t") < "hello"))
    _ = p0
    _ = v4
    return v4
}

func main() {
    var v1 float64 = float64((1e10 - (1.5 * 1e10)))
    v2 := make(map[string]uint32)
    var v3 uint8 = 255
    v1 = float64(0.1)
    var v5 int16 = 37
    if (65535 < 1) {
        var v6 uint8 = (v3 + (v3 >> ((412) & 63)))
        _ = v6
    }
    v3 = v3
    v5 = (v5 / (((v5 / ((42) | 1))) | 1))
    var v7 uint = 450
    v8 := []int{4294967295, -741}
    switch 9223372036854775807 {
    case 9223372036854775807:
        _ = len(v8)
    case 127:
        fmt.Println(v5, v5)
    default:
        var v9 uint16 = 32767
        _ = v9
    }
    v2[("hello" + "")] = 0
    v10 := GenS1{F0: 145, F1: (^(v3))}
    fmt.Println(v1, v1, v7)
    _ = len(v8)
    var v11 uint64 = 2
    var v12 string = "a"
    v12 = v12
    v13 := GenS1{F0: (v3 >> ((v7) & 63)), F1: (^(v3))}
    if (uint32(v11) == 978) {
        v14 := GenS2{F0: (v11 % (((v11 * v11)) | 1)), F1: (v5 / (((v5 % (((v5 << ((v7) & 63))) | 1))) | 1)), F2: false}
        for fi15 := 0; fi15 < 14; fi15++ {
            switch -1 {
            case 451:
                var v16 uint8 = v3
                _ = v16
            case -1:
                v12 = v12
            case -155:
                var v17 float64 = v1
                _ = v17
            }
            v18 := GenS2{F0: v11, F1: v5, F2: true}
            v5 -= v5
            _ = fi15
            _ = v18
        }
        _ = v14
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v7
    _ = v8
}
