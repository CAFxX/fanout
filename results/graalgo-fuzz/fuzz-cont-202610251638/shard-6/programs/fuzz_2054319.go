package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int64 }
func genFn0(p0 string) uint {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 bool = false
        _ = fi7
        _ = v8
    }
    v9 := []int{4294967295}
    _ = p0
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "\n\t"
    v3 := make(map[uint32]bool)
    v4 := []int{-526, 345, -1, 127}
    v5 := []int{}
    var v6 uint64 = 1
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS1{F0: 42, F1: ((v2 + ((v2 + v2) + v2)) + v2), F2: -1}
        _ = fi10
        _ = v11
    }
    delete(v3, 4294967295)
    v6 -= v6
    var v12 uint16 = 42
    fmt.Println(v12, v2)
    v6 -= v6
    v12 = v12
    switch v2 {
    case "foo bar":
        var v13 int8 = 42
        _ = v13
    case "\n\t":
        var v14 uint64 = v6
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
