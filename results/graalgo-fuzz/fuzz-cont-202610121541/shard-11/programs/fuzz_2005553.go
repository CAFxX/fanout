package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int8 }
type GenS2 struct { F0 int32; F1 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    v4 := []string{("\n\t" + "foo bar"), "a"}
    v5 := make(map[string]uint32)
    var v6 int = 32768
    var v7 int = (v6 | (v6 - v6))
    v6 += v6
    var v8 int64 = 953
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: -2147483648, F1: 255, F2: 46}
        var v11 float32 = float32((1e10 * 1e10))
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 42, F1: v6}
    fmt.Println(v7, v6)
    v13 := GenS1{F0: 2147483647, F1: 2, F2: 1}
    v14 := GenS2{F0: -2147483648, F1: v6}
    _ = len(v3)
    var v15 int64 = v8
    var v16 string = ("a" + "\n\t")
    var v17 uint8 = 2
    v8 = (v8 | v8)
    for fi18 := 0; fi18 < 6; fi18++ {
        _ = len(v1)
        fmt.Println(v17, v16)
        _ = fi18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 float32 = 1.5
        v15 += (v8 + v15)
        var v21 float32 = 0.0
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
