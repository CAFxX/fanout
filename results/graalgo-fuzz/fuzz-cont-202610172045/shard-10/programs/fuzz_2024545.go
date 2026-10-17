package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    var v1 int16 = -1
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    v4 := []string{(("foo bar" + "hello") + (("a" + (("" + "\n\t") + "a")) + "a"))}
    v5 := make(map[uint32]int)
    v6 := make(map[string]uint32)
    var v7 bool = (0 >= 1)
    var v8 float32 = (-2.25 / 1.0)
    v9 := v5[0]
    fmt.Println(v8, v7)
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v7)
        var v11 int = (fi10 * v9)
        _ = fi10
        _ = v11
    }
    var v12 int64 = 9223372036854775807
    v1 = v1
    v13 := GenS1{F0: (v1 & v1), F1: 733}
    var v14 float32 = 1e10
    var v15 uint16 = 2
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 uint32 = 2
        var v18 float64 = float64(float32(-2.25))
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: (v1 >> ((234) & 63)), F1: 0}
    v20 := v5[1]
    var v21 uint64 = 2
    var v22 int8 = 0
    v23 := GenS1{F0: v1, F1: v21}
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
