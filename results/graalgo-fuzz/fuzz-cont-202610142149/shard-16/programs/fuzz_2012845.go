package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
func genFn0() int64 {
    var v5 uint32 = 2147483648
    v5 -= v5
    var v6 uint32 = v5
    fmt.Println(v5)
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 uint16 = 913
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    var v4 uint = 827
    fmt.Println(v1, v1)
    var v7 uint32 = 0
    var v8 uint16 = 255
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 int64 = genFn0()
        v11 := GenS1{F0: 42, F1: 0}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = v1
    v4 += (v4 / ((v4) | 1))
    v13 := GenS1{F0: 9223372036854775807, F1: 85}
    fmt.Println(v12, v7, v4)
    v12 = v12
    v12 -= v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
