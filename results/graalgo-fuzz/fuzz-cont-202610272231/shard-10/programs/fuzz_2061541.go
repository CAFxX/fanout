package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    var v1 bool = (255 == 1)
    var v2 uint16 = 1
    v3 := make(map[int]bool)
    var v4 int32 = 42
    v3[10] = true
    fmt.Println(v2, v1, v1)
    if (4294967295 <= 133) {
        var v5 uint8 = 128
        v6 := GenS1{F0: 0}
        fmt.Println(v2, v4)
        _ = v5
        _ = v6
    }
    v4 += v4
    delete(v3, -556)
    v4 = (v4 - (v4 & (v4 | (v4 << ((2) & 63)))))
    var v7 int32 = v4
    v2 -= 65535
    v8 := GenS1{F0: 819}
    switch 0 {
    case 9223372036854775807:
        var v9 int32 = 728
        _ = v9
    case 128:
        v10 := GenS1{F0: 396}
        _ = v10
    case 685:
        if (!v1) {
            fmt.Println(v2)
            v11 := GenS1{F0: 4294967295}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 14; fi12++ {
                _ = len(v3)
                _ = fi12
            }
        }
    }
    v7 += -2147483648
    v3[-2147483648] = v1
    v4 = v7
    v13 := GenS2{F0: 1, F1: 9223372036854775807}
    delete(v3, 128)
    var v14 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
