package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 float32 = 0.1
    var v4 uint64 = 1
    var v5 uint32 = 42
    delete(v2, -9223372036854775808)
    var v6 uint8 = 196
    if (v6 >= v6) {
        fmt.Println(v5)
        switch 1 {
        case 961:
            v5 -= 179
        case 42:
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := []int{}
                _ = fi7
                _ = v8
            }
        }
        _ = len(v2)
    } else {
        var v9 int8 = 63
        _ = v9
    }
    var v10 int16 = -432
    var v11 uint64 = v4
    v12 := []int{}
    if false {
        for fi13 := 0; fi13 < 4; fi13++ {
            if (65535 != 247) {
                v5 += 1
            }
            _ = fi13
        }
        _ = len(v1)
        v14 := GenS1{F0: 525, F1: v10}
        _ = v14
    }
    switch 9223372036854775807 {
    case 127:
        var v15 uint = 2
        _ = v15
    }
    if (0 > (v4 >> ((42) & 63))) {
        var v16 float32 = 3.14159
        _ = v16
    }
    v17 := GenS1{F0: 171, F1: 13}
    v3 -= -2.25
    if false {
        var v18 int8 = 100
        var v19 int64 = int64(v6)
        _ = v18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
