package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 bool }
func genFn0() int8 {
    var v7 uint32 = 4294967295
    v7 += (v7 * v7)
    var v8 int = 32768
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int64, p1 int8, p2 int64) uint {
    p2 -= p0
    v9 := []int{42, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 561
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 0.0
    var v3 float32 = (0.0 / 3.14159)
    var v4 uint8 = 128
    v5 := []uint32{uint32(v4), 42, 596, 663}
    var v6 float32 = ((v3 - (v2 + ((v2 - (1e10 - (v2 * v2))) - 0.0))) * 3.14159)
    v6 -= 1.5
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS1{F0: 1, F1: 0, F2: (v3 != v3)}
        fmt.Println(v6, v2, v6)
        _ = fi10
        _ = v11
    }
    v2 -= v6
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 = v2
    var v12 float32 = ((v2 * -2.25) * v3)
    var v13 bool = false
    v4 -= v4
    _ = len(v5)
    switch "hello" {
    case "a":
        v12 -= (float32(1.5) - v12)
    case "":
        var v14 string = ""
        _ = v14
    }
    v6 += v6
    var v15 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
