package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int64, p1 int16) int8 {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float32 = 1e10
        _ = fi7
        _ = v8
    }
    v9 := []int{-2147483648, 4294967295, 42, 0}
    _ = len(v9)
    fmt.Println(p1)
    switch 9223372036854775807 {
    case 359:
        p1 -= 1000
    case 42:
        var v10 int64 = -754
        _ = v10
    case 128:
        fmt.Println(p0, p1)
    default:
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = v9
    return 127
}

func main() {
    var v1 uint8 = 2
    v2 := []int{1}
    v3 := make(map[uint32]bool)
    var v4 int16 = -32768
    var v5 int8 = 100
    v6 := make(map[uint32]uint32)
    var v11 uint = 2
    v6[0] = 0
    v5 = (-(v5))
    var v12 int64 = 42
    var v13 bool = true
    v14 := GenS1{F0: v1}
    fmt.Println(v12, v5, v4)
    v4 = v4
    if v13 {
        v5 = (^(v5))
    }
    if len(v2) > 0 {
        v2[0] = -1
    }
    v2 = append(v2, 1)
    fmt.Println(v5, v5)
    var v15 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
