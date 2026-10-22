package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v4 := []string{}
    var v5 float64 = 1e10
    v5 = 3.14159
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: fi6}
        v5 = (0.0 * v5)
        _ = fi6
        _ = v7
    }
    var v8 string = ""
    v8 = (v8 + ((("\n\t" + "\n\t") + v8) + "foo bar"))
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint = 320
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 88, F1: 2147483648}
    v5 += ((0.0 / v5) + v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
