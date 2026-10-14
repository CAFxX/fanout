package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint16 }

func main() {
    v1 := []uint32{373, 1, 4294967295}
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 float64 = -2.25
    var v5 string = ("" + "hello")
    v1 = append(v1, 65535)
    var v6 int16 = -32768
    v7 := GenS1{F0: -1, F1: 346, F2: 345}
    fmt.Println(v6)
    fmt.Println(v5, v4, v6)
    fmt.Println(v5)
    for fi8 := 0; fi8 < 15; fi8++ {
        v6 = (^(v6))
        _ = fi8
    }
    var v9 float64 = (0.1 - (3.14159 - v4))
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := v2[(v5 + (v5 + "foo bar"))]
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: v6, F1: 65535, F2: 31}
    v6 = (v6 % ((756) | 1))
    var v13 uint32 = 45
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
