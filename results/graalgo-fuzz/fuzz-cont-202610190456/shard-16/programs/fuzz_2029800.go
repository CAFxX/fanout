package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
func genFn0(p0 uint32, p1 uint, p2 uint8) uint {
    if (1 <= -9223372036854775808) {
        if false {
            var v7 bool = false
            var v8 int16 = 42
            _ = v7
            _ = v8
        }
        fmt.Println(p2, p0)
    } else {
        var v9 uint64 = 428
        fmt.Println(p0, p0, p1)
        fmt.Println(p0, p1)
        _ = v9
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p1
}

func main() {
    var v1 int32 = -1
    var v2 uint32 = 1
    var v3 bool = true
    v4 := []uint32{}
    var v5 uint32 = 2147483647
    var v6 int16 = 758
    fmt.Println(v2, v3, v6)
    var v11 string = "hello"
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(fi12, v1)
        _ = fi12
    }
    if (!((1.5 / 0.1) != 3.14159)) {
        v5 = v5
        v5 = v2
    } else {
        var v13 int32 = (^(v1))
        var v14 float64 = -2.25
        _ = v13
        _ = v14
    }
    var v15 int32 = v1
    v16 := GenS1{F0: 2, F1: v6}
    if (v11 <= v11) {
        var v17 uint8 = 2
        _ = v17
    } else {
        v1 = v15
        _ = len(v4)
        fmt.Println(v5)
    }
    var v18 uint16 = 857
    fmt.Println(v15, v3, v3)
    if v3 {
        switch -576 {
        case 4294967295:
            var v19 string = "\n\t"
            _ = v19
        default:
            v15 = 2147483647
        }
    }
    v20 := GenS1{F0: 255, F1: (v6 | (v6 / (((v6 >> ((230) & 63))) | 1)))}
    var v21 uint8 = 127
    fmt.Println(v15)
    v18 = (v18 << ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
