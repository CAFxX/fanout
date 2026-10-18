package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int32 }

func main() {
    v1 := []uint32{2147483648, 2147483647, 42}
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 string = "a"
    v5 := make(map[int]uint32)
    var v6 string = "foo bar"
    v7 := GenS1{F0: 255, F1: (v6 + v6), F2: -686}
    v3[4294967295] = true
    for fi8 := 0; fi8 < 5; fi8++ {
        v6 = v6
        _ = fi8
    }
    var v9 float64 = -2.25
    var v10 uint = 2
    if (0 < -1) {
        switch 179 {
        case -896:
            v10 -= v10
        case -2147483648:
            v4 = v6
        }
    }
    if true {
        v2[1] = -1
    } else {
        fmt.Println(v9, v10, v10)
        v9 -= float64((((1e10 - 1.5) * float32((1e10 - (v9 + float64((((0.0 - 0.0) + (float32((0.0 - v9)) + ((((1.5 / 1.0) + 0.1) + 0.0) * -2.25))) * 0.0)))))) + (-2.25 + 0.1)))
    }
    var v11 uint64 = 4294967295
    v11 += 946
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
