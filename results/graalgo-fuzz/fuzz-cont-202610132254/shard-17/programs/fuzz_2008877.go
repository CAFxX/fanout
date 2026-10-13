package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }

func main() {
    v1 := []string{}
    v2 := []string{"", "foo bar", "foo bar", "hello"}
    var v3 int8 = -92
    v4 := make(map[int]string)
    var v5 int64 = 0
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(v5)
        fmt.Println(v3, fi6)
        fmt.Println(v5, v5, fi6)
        _ = fi6
    }
    if false {
        var v7 uint64 = 228
        var v8 uint = 4294967295
        _ = v7
        _ = v8
    } else {
        var v9 int64 = (v5 % (((-(v5))) | 1))
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: 18446744073709551615}
    v11 := GenS1{F0: 14, F1: 1}
    v5 = -415
    var v12 float64 = float64((-2.25 + 3.14159))
    switch -2147483648 {
    case -918:
        var v13 uint32 = 223
        _ = v13
    default:
        var v14 uint32 = 2147483648
        _ = v14
    }
    if (((float32(v12) * 3.14159) < 0.0) || true) {
        var v15 float64 = float64(float32(float64(1e10)))
        v16 := GenS2{F0: v3, F1: v3, F2: 441}
        var v17 uint8 = 127
        _ = v15
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 2; fi18++ {
            if (false || true) {
                if false {
                    fmt.Println(v3)
                }
                var v19 uint = 18446744073709551615
                _ = v19
            }
            v20 := GenS2{F0: v3, F1: (v3 + 1), F2: 2}
            _ = fi18
            _ = v20
        }
        var v21 uint16 = 65535
        _ = v21
    }
    var v22 int16 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
