package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 int8 = 1
    var v2 int32 = -1
    v3 := make(map[string]string)
    _ = len(v3)
    v1 = 59
    var v4 int32 = v2
    v1 = -75
    fmt.Println(v1)
    v1 = (-(v1))
    var v5 float64 = 1.5
    var v6 uint64 = 18446744073709551615
    v7 := GenS1{F0: "foo bar", F1: v6}
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            v9 := GenS2{F0: 1}
            _ = v9
        } else {
            v6 -= (^(v6))
            var v10 uint16 = 2
            _ = v10
        }
        _ = fi8
    }
    v11 := []int{0, 0}
    if false {
        v12 := GenS3{F0: 42, F1: 127}
        _ = v12
    } else {
        v11 = append(v11, 513)
        v13 := GenS1{F0: "hello", F1: (v6 | v6)}
        _ = v13
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
