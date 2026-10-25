package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int }
func genFn0(p0 bool, p1 int32) int64 {
    var v5 uint32 = 2147483647
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: (-(fi6))}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return -9223372036854775808
}

func main() {
    v1 := []uint32{}
    var v2 uint = 905
    v3 := make(map[int]bool)
    var v4 int8 = 0
    delete(v3, int(v4))
    if true {
        var v8 int = 1
        v9 := GenS2{F0: 42, F1: 127}
        _ = v8
        _ = v9
    }
    var v10 int16 = 42
    v3[32767] = false
    var v11 string = "a"
    v12 := GenS2{F0: 2147483648, F1: 2147483647}
    if len(v1) > 0 {
        v1[0] = 218
    }
    var v13 int8 = 42
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
