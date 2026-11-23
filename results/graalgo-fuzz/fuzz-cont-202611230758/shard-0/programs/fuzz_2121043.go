package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
func genFn0(p0 uint16, p1 int) bool {
    p0 = 42
    fmt.Println(p0)
    v6 := []int{(^(p1)), (p1 * p1), 2147483647}
    p0 += (p0 - (^(p0)))
    _ = p0
    _ = p1
    _ = v6
    return true
}
func genFn1(p0 uint, p1 string) uint {
    p0 = (p0 + p0)
    v7 := []int{}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 uint16 = 65535
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (^(p0))
}

func main() {
    var v1 uint16 = 65535
    v2 := []int{42, 0, 1, -720}
    var v3 uint8 = 2
    var v4 int64 = -222
    v5 := make(map[string]uint32)
    var v9 int8 = 100
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v10 := GenS1{F0: -1, F1: (v4 | 1), F2: v9}
    var v11 int64 = -1
    var v12 uint16 = 1
    var v13 int32 = -2147483648
    fmt.Println(v1, v1, v4)
    v12 += 109
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 int16 = -32768
        var v16 int16 = (v15 + v15)
        _ = fi14
        _ = v15
        _ = v16
    }
    if len(v2) > 0 {
        v2[0] = -1
    }
    var v17 float32 = 0.1
    var v18 int = 268
    var v19 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
