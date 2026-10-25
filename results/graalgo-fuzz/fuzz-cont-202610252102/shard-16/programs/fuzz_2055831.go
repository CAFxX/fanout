package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
func genFn0(p0 int8, p1 int32) uint32 {
    p0 = p0
    v7 := GenS1{F0: -767, F1: 275}
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 uint32) uint8 {
    var v8 uint8 = 20
    fmt.Println(p0, v8)
    fmt.Println(p0, p0)
    p0 = p0
    _ = p0
    _ = v8
    return v8
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -1
    var v3 int8 = v2
    var v4 string = ""
    v5 := []string{(v4 + "a"), (("a" + "foo bar") + (v4 + "foo bar"))}
    var v6 uint8 = 13
    var v9 int64 = 1
    fmt.Println(v3)
    var v10 float64 = float64(0.1)
    v11 := GenS1{F0: 496, F1: 2}
    fmt.Println(v10, v3, v3)
    var v12 uint64 = 18446744073709551615
    var v13 float64 = float64(0.0)
    v4 = v4
    delete(v1, 32768)
    for fi14 := 0; fi14 < 7; fi14++ {
        fmt.Println(v13, v13)
        v4 = v4
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
