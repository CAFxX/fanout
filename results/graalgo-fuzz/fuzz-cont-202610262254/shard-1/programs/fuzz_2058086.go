package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }

func main() {
    var v1 int = 2147483647
    var v2 int8 = -1
    v3 := make(map[int]int)
    var v4 int32 = -1
    v5 := make(map[uint32]string)
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int64 = 0
        fmt.Println(v8)
        _ = fi7
        _ = v8
    }
    fmt.Println(v4, v2)
    var v9 float64 = (0.0 / 0.1)
    v4 = v4
    var v10 uint64 = 4294967295
    var v11 float64 = v9
    switch v1 {
    case 50:
        v9 = v11
    case -222:
        v11 = float64((0.1 * (3.14159 / -2.25)))
    case 32768:
        v12 := []int{}
        _ = v12
    }
    v4 += v4
    v13 := []int{}
    if (false && true) {
        var v14 uint32 = 2
        if false {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            v9 = v11
        }
        var v15 uint64 = (^(v10))
        _ = v14
        _ = v15
    } else {
        fmt.Println(v4)
        v16 := GenS1{F0: 2, F1: 0, F2: 1}
        _ = v16
    }
    var v17 uint = 269
    v9 += ((1.5 - (v11 * float64(1.5))) + v9)
    var v18 int16 = int16(v2)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v2 += 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
