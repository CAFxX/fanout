package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int = 294
    var v2 uint = 255
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    v5 := []int{(v1 * v1), (^(v1)), 32767, (v1 >> ((2) & 63))}
    fmt.Println(v2)
    v6 := GenS1{F0: 100, F1: (^(v2))}
    v2 = (v2 / ((v2) | 1))
    v7 := GenS2{F0: -1, F1: 1, F2: 0}
    switch "hello" {
    case "foo bar":
        v8 := GenS3{F0: 1}
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 10; fi9++ {
            if true {
                var v10 float32 = 1e10
                _ = v10
            }
            _ = fi9
        }
    case "\n\t":
        v1 = (v1 * 127)
    }
    _ = len(v5)
    fmt.Println(v2, v1)
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
