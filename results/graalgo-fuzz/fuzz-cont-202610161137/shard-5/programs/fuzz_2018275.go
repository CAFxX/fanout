package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    var v4 uint64 = (v1 << ((65535) & 63))
    v5 := GenS1{F0: ("\n\t" + "foo bar"), F1: -57}
    if (-214 == 4294967295) {
        v6 := GenS1{F0: "hello", F1: 508}
        _ = v6
    }
    if true {
        var v7 float32 = ((1e10 / 0.1) + 0.0)
        _ = v7
    }
    v1 = 537
    var v8 uint64 = (v1 / ((v4) | 1))
    v9 := GenS1{F0: "a", F1: -1}
    var v10 float64 = 1e10
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 uint64 = (v1 ^ v1)
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(v8, v8)
            _ = fi13
        }
        var v14 int16 = 234
        _ = fi11
        _ = v12
        _ = v14
    }
    v15 := []int{547, 163, -1}
    v4 = v1
    var v16 uint8 = 122
    if true {
        var v17 uint32 = 4294967295
        v18 := GenS2{F0: 42, F1: 76}
        v19 := GenS1{F0: "a", F1: 9223372036854775807}
        _ = v17
        _ = v18
        _ = v19
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
