package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 string; F2 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = float64(float32(1.5))
    var v3 int16 = -213
    v4 := []uint32{0, 1}
    var v5 uint8 = 128
    var v6 uint = 255
    v7 := GenS1{F0: 2147483647}
    v6 -= v6
    var v8 float32 = 1e10
    v8 -= v8
    if (522 > 65535) {
        fmt.Println(v2, v8)
        v4 = append(v4, 4294967295)
    } else {
        v5 = 15
        var v9 uint64 = uint64(v5)
        _ = v9
    }
    v5 -= v5
    if true {
        var v10 int = 2147483647
        if true {
            var v11 float64 = v2
            _ = v11
        }
        v12 := GenS1{F0: 0}
        _ = v10
        _ = v12
    } else {
        var v13 uint16 = 32767
        _ = v13
    }
    if (v5 < 255) {
        v14 := GenS2{F0: (^(v6)), F1: "", F2: 486}
        _ = v14
    }
    if ((^(v5)) != (^(v5))) {
        var v15 int32 = 933
        v15 -= 196
        _ = v15
    } else {
        v16 := GenS2{F0: (v6 >> ((v6) & 63)), F1: "a", F2: 0}
        v3 -= (v3 | 32767)
        var v17 int16 = (v3 << ((v6) & 63))
        _ = v16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
