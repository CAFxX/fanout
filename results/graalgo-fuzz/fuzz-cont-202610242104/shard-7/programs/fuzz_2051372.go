package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }

func main() {
    var v1 int = 32767
    v2 := make(map[string]int)
    v3 := []uint32{100, 0}
    var v4 float32 = 3.14159
    var v5 uint8 = 128
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 0, F1: (v1 & (^(v1)))}
        _ = fi6
        _ = v7
    }
    v3 = append(v3, 1)
    if ((!false) && (false || true)) {
        delete(v2, "\n\t")
        var v8 bool = (2147483647 == 0)
        var v9 uint32 = 806
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: -1, F1: (v1 << ((117) & 63))}
    switch (v1 / (((v1 & v1)) | 1)) {
    case -575:
        if (false || false) {
            v2["a"] = -693
            v1 -= (v1 | v1)
            var v11 int8 = 127
            _ = v11
        }
    }
    if (1e10 == -2.25) {
        v4 -= ((0.1 / v4) + ((v4 - v4) * v4))
        v1 = (v1 << ((65535) & 63))
        fmt.Println(v1)
    }
    fmt.Println(v4, v1)
    _ = len(v3)
    v12 := GenS1{F0: -128, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
