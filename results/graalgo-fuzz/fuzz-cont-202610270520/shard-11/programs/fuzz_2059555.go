package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = (1e10 - 1.5)
    var v3 int = 127
    var v4 float32 = -2.25
    v5 := []uint32{4294967295, 2}
    v2 = v4
    for fi6 := 0; fi6 < 9; fi6++ {
        v2 = (1.5 + (v2 * (3.14159 + 3.14159)))
        _ = fi6
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 = v3
    var v7 uint16 = 42
    fmt.Println(v7, v2)
    var v8 uint64 = 4294967295
    if (34 <= 127) {
        v8 -= (v8 % (((v8 & v8)) | 1))
        switch (v3 / ((v3) | 1)) {
        case -93:
            v9 := GenS1{F0: 1, F1: 0}
            _ = v9
        case 0:
            var v10 float64 = ((3.14159 / 1e10) - float64(v2))
            _ = v10
        case 128:
            v2 -= float32(3.14159)
        }
        v11 := GenS1{F0: -2147483648, F1: 198}
        _ = v11
    }
    v7 = (v7 - v7)
    v5 = append(v5, 765)
    v12 := GenS2{F0: 24}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
