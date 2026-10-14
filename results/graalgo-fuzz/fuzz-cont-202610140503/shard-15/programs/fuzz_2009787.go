package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 753
    v3 := make(map[uint32]uint32)
    var v4 int64 = -9223372036854775808
    v2 += (^(v2))
    v2 = v2
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v4, v4)
    }
    v2 = (v2 / (((v2 ^ (^(v2)))) | 1))
    v5 := []int{1}
    v6 := GenS1{F0: (v4 * (-(v4)))}
    switch "foo bar" {
    case "\n\t":
        v7 := GenS2{F0: 1}
        _ = v7
    case "hello":
        v8 := GenS1{F0: 1}
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint16 = 1
        var v11 int16 = 32767
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        v4 = v4
        v4 = v4
        v13 := GenS2{F0: 703}
        _ = fi12
        _ = v13
    }
    var v14 bool = true
    fmt.Println(v14, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
