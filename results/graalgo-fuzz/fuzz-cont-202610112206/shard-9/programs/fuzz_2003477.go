package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 int8 }

func main() {
    var v1 int8 = 100
    v2 := []string{"a"}
    v3 := []uint32{}
    var v4 bool = (true && true)
    v5 := make(map[int]uint32)
    v4 = v4
    v2 = append(v2, "foo bar")
    v3 = append(v3, 0)
    v6 := GenS1{F0: 128, F1: 1}
    if v4 {
        var v7 uint = 2
        _ = v7
    } else {
        var v8 string = "foo bar"
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(v1, fi9)
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS2{F0: 900, F1: 0, F2: (v1 - 97)}
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        v2 = append(v2, "hello")
        _ = v8
    }
    v4 = (68 < 42)
    var v12 float32 = 0.0
    v12 = (v12 * 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
