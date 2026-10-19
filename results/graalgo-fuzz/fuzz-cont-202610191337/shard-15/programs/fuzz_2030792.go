package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint64 }

func main() {
    var v1 float64 = 1e10
    var v2 int64 = 379
    v3 := make(map[int]uint32)
    fmt.Println(v1, v2)
    v1 -= v1
    v4 := []int{-1}
    fmt.Println(v1)
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 string = ("a" + "foo bar")
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi5
        _ = v6
    }
    var v7 float32 = (float32(v1) + ((1.5 - ((-2.25 - 3.14159) - 0.1)) + 0.0))
    switch "foo bar" {
    case "foo bar":
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: 42}
            var v10 int32 = 0
            _ = fi8
            _ = v10
            _ = v9
        }
    case "":
        v11 := GenS1{F0: 32767}
        _ = v11
    }
    v12 := GenS2{F0: 610, F1: 4294967295, F2: 0}
    v1 -= ((v1 / 0.1) + 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
