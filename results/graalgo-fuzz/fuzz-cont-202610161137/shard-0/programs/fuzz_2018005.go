package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0(p0 int32, p1 int64) uint32 {
    for fi5 := 0; fi5 < 14; fi5++ {
        p1 = p1
        _ = fi5
    }
    var v6 float32 = (0.0 + (1e10 - (0.0 / -2.25)))
    p1 = (p1 ^ p1)
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 int8, p1 int, p2 int) uint16 {
    p1 = (-(p2))
    p2 = p2
    switch "hello" {
    case "\n\t":
        var v7 int = p2
        _ = v7
    case "foo bar":
        if true {
            fmt.Println(p2)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 int32 = 1
    var v2 int16 = 32767
    var v3 uint64 = 18446744073709551615
    v4 := []int{32768, -240, 2147483647}
    fmt.Println(v3, v3, v1)
    var v8 uint64 = v3
    fmt.Println(v1)
    var v9 uint16 = 32767
    var v10 uint32 = 1
    _ = len(v4)
    v11 := GenS1{F0: 42}
    v12 := GenS2{F0: 0, F1: 0}
    var v13 uint16 = v9
    v14 := GenS1{F0: 127}
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v15 := GenS3{F0: v9, F1: 42}
    _ = len(v4)
    fmt.Println(v9, v10, v2)
    fmt.Println(v13, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
