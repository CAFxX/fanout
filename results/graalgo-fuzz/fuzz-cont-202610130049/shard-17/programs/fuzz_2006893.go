package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-9223372036854775808, 4294967295, 32768}
    var v3 int16 = 1
    v4 := make(map[int]uint32)
    var v5 int64 = -264
    v6 := make(map[uint32]int)
    fmt.Println(v5, v3)
    for fi7 := 0; fi7 < 6; fi7++ {
        v5 = 0
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 bool = true
            v10 := v6[2147483647]
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 float32 = 3.14159
        _ = fi7
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS1{F0: 25, F1: (^(v5))}
        _ = fi12
        _ = v13
    }
    fmt.Println(v5, v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3, v3)
    var v14 uint = 540
    v14 = (^(v14))
    fmt.Println(v3, v5)
    v5 += 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
