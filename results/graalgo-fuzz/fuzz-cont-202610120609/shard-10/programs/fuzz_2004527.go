package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []int{4294967295, 0}
    var v2 uint64 = 42
    var v3 int16 = 1
    var v4 float64 = (float64(3.14159) - 0.0)
    v5 := make(map[int]uint32)
    var v6 int64 = 1
    v7 := GenS1{F0: 42}
    var v8 uint = 0
    var v9 float32 = 0.1
    v4 += v4
    v10 := GenS1{F0: (v6 << (((v8 * (v8 | (v8 & (^(v8)))))) & 63))}
    fmt.Println(v2, v8, v3)
    var v11 int64 = (v6 & (v6 + v6))
    v6 = (v11 - v11)
    if (!(!(!(42 >= 100)))) {
        v2 = v2
        v8 = v8
    } else {
        switch 435 {
        case 843:
            var v12 int16 = v3
            _ = v12
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v13 string = "a"
        _ = v13
    }
    if ((true || true) || true) {
        v3 = -73
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int8 = 2
            _ = fi14
            _ = v15
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
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
    _ = v8
    _ = v9
}
