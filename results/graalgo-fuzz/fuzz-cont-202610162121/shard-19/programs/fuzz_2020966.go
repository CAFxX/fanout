package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint = 255
    var v2 uint8 = 1
    var v3 float32 = ((1.5 * 0.1) * 0.1)
    v4 := make(map[int]int)
    var v5 int = -9223372036854775808
    var v6 int16 = 0
    v7 := GenS1{F0: (v5 * (v5 + (^(v5)))), F1: 2}
    v2 = 92
    delete(v4, (v5 + (v5 << ((977) & 63))))
    var v8 uint16 = 65535
    fmt.Println(v5, v3)
    switch 42 {
    case 0:
        var v9 int16 = (^(v6))
        _ = v9
    case 2147483647:
        v10 := GenS1{F0: 4294967295, F1: 191}
        _ = v10
    case -2147483648:
        var v11 uint64 = 0
        _ = v11
    }
    v6 += (v6 | 330)
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 int16 = -285
        _ = fi12
        _ = v13
    }
    var v14 uint64 = 67
    if false {
        var v15 bool = true
        var v16 uint64 = (v14 % ((v14) | 1))
        _ = v15
        _ = v16
    }
    v1 = (^(v1))
    var v17 uint64 = uint64(v1)
    var v18 int16 = 1000
    v19 := []int{v5}
    if (511 <= (^(v1))) {
        var v20 float64 = (3.14159 + 0.0)
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 int16 = v6
            v5 = fi21
            _ = len(v4)
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    if false {
        v3 = v3
        var v23 uint64 = 42
        _ = v23
    } else {
        var v24 uint = 42
        if false {
            v3 -= 0.0
            var v25 uint32 = 636
            var v26 uint64 = 953
            _ = v25
            _ = v26
        }
        fmt.Println(v6, v8, v14)
        _ = v24
    }
    v27 := GenS2{F0: (^(v14))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
