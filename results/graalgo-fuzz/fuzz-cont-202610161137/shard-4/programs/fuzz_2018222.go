package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 42
    v3 := make(map[string]uint32)
    if (true && false) {
        fmt.Println(v2)
        v2 += v2
    }
    var v4 int = v2
    var v5 uint32 = 2147483647
    if false {
        fmt.Println(v5)
    }
    v3[""] = 42
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := v3["a"]
        _ = fi6
        _ = v7
    }
    v4 = (v4 / ((32767) | 1))
    if true {
        switch 32767 {
        case 294:
            v4 += (-(v2))
        }
        v4 = 32767
    } else {
        v8 := GenS1{F0: 1, F1: 1, F2: 416}
        _ = v8
    }
    v4 += v2
    v9 := v3["a"]
    var v10 uint = 456
    v10 += v10
    if true {
        v9 = (^(v5))
        var v11 float32 = -2.25
        _ = v11
    }
    var v12 float32 = 0.1
    v12 = float32(float64((v12 + (v12 * v12))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
