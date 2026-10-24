package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint32, p1 int8) int8 {
    v6 := []int{}
    v7 := GenS1{F0: 799}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -92
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 int64 = 42
    var v4 int16 = 946
    var v5 float64 = (0.0 / 1.0)
    delete(v2, 2147483648)
    for fi8 := 0; fi8 < 6; fi8++ {
        switch (fi8 / (((fi8 + (fi8 / ((fi8) | 1)))) | 1)) {
        case 127:
            fmt.Println(v5, fi8, v4)
        case 486:
            v9 := []int{}
            _ = v9
        }
        _ = len(v1)
        _ = fi8
    }
    var v10 int64 = (v3 ^ (v3 / ((0) | 1)))
    fmt.Println(v3)
    var v11 int64 = (^(v3))
    var v12 float64 = 0.0
    var v13 int16 = (^(v4))
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint16 = 255
        fmt.Println(v11, v5, v4)
        v2[42] = 2147483648
        _ = fi14
        _ = v15
    }
    v11 = (v3 << ((2) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
