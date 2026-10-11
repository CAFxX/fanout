package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 uint64 = 163
    var v5 float64 = 0.0
    v2["a"] = 921
    var v6 int8 = -1
    v4 -= v4
    v7 := GenS1{F0: v6, F1: 9223372036854775807}
    delete(v2, ("hello" + ("" + "\n\t")))
    for fi8 := 0; fi8 < 8; fi8++ {
        v6 = (v6 << ((4294967295) & 63))
        fmt.Println(v6)
        _ = fi8
    }
    v9 := GenS2{F0: (v4 % ((v4) | 1))}
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            v5 += (float64(-2.25) * 1.5)
            _ = fi11
        }
        var v12 uint = 42
        _ = fi10
        _ = v12
    }
    var v13 float32 = 1e10
    v14 := []int{716, 32768}
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS1{F0: v6, F1: -1}
        v17 := GenS2{F0: v4}
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(v5)
    var v18 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
