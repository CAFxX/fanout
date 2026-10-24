package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint16 }
func genFn0(p0 uint32) string {
    var v7 uint = 0
    switch 428 {
    case -1:
        v8 := []int{}
        _ = v8
    case -2147483648:
        v9 := GenS1{F0: 2}
        _ = v9
    }
    _ = p0
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 262
    var v3 bool = false
    v4 := make(map[uint32]int)
    v5 := make(map[int]string)
    var v6 int32 = -1
    v10 := GenS2{F0: 9223372036854775807, F1: 1, F2: 255}
    var v11 int32 = v6
    switch -1 {
    case 9223372036854775807:
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 uint64 = 4294967295
            fmt.Println(v2)
            _ = fi12
            _ = v13
        }
    default:
        var v14 uint = 717
        _ = v14
    }
    v3 = v3
    var v15 uint32 = 2147483647
    if false {
        fmt.Println(v15, v11)
    } else {
        v16 := GenS1{F0: 42}
        v6 -= (-(v6))
        _ = v16
    }
    var v17 uint64 = 1
    var v18 uint16 = 1
    v2 += -870
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
