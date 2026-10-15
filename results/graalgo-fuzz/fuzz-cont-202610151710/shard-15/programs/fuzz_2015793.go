package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []uint32{373, 0, 0, 286}
    var v2 int8 = 73
    var v3 string = "hello"
    v4 := make(map[int]bool)
    var v5 string = "\n\t"
    v6 := []string{v5, v3, v3}
    fmt.Println(v5, v5, v3)
    var v7 int32 = -932
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 uint8 = 0
    v8 = (v8 | v8)
    for fi9 := 0; fi9 < 18; fi9++ {
        _ = len(v4)
        _ = fi9
    }
    v10 := GenS1{F0: 1}
    for fi11 := 0; fi11 < 13; fi11++ {
        v8 += 112
        fmt.Println(v5, v5)
        _ = fi11
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    var v12 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
