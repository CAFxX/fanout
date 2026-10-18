package main

import (
    "fmt"
)

func genFn0(p0 uint16) uint8 {
    p0 += 1
    var v6 uint = uint(p0)
    _ = p0
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    var v4 int = -1
    v5 := make(map[string]int)
    v2 = -107
    v7 := []int{9223372036854775807, 127}
    if true {
        var v8 bool = true
        v4 = (v4 * (-(v4)))
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint64 = 18446744073709551615
        var v11 int16 = -129
        fmt.Println(v2)
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v4, v4)
    fmt.Println(v4)
    _ = len(v1)
    var v12 int16 = 640
    var v13 int = v4
    v13 += v13
    v12 = v12
    v1[(^(v4))] = 736
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
