package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = (("" + "a") + "hello")
    var v3 float64 = 1.5
    v4 := make(map[int]bool)
    v3 += v3
    v5 := GenS1{F0: -1, F1: 1, F2: 1}
    fmt.Println(v2, v2)
    var v6 uint32 = 2147483647
    if (4294967295 <= 192) {
        v2 = v2
    } else {
        fmt.Println(v3)
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            v2 = v2
            v9 := GenS2{F0: 2}
            var v10 int16 = 560
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v2 = v2
    v6 -= v6
    var v11 uint8 = 2
    v11 = v11
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := v1[0]
        _ = fi12
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
