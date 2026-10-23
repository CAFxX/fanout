package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 string }

func main() {
    var v1 int64 = 448
    v2 := make(map[string]uint32)
    var v3 uint = 131
    var v4 uint = v3
    v5 := make(map[string]string)
    v6 := GenS1{F0: ""}
    var v7 float64 = (0.1 * -2.25)
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: "a"}
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: "foo bar"}
    v1 -= (v1 | v1)
    if false {
        v11 := GenS2{F0: (v3 - v4), F1: "foo bar"}
        _ = v11
    }
    v12 := []int{32767, -9223372036854775808, 505}
    for fi13 := 0; fi13 < 11; fi13++ {
        v4 += (v4 >> (((^(v4))) & 63))
        _ = fi13
    }
    if len(v12) > 0 {
        v12[0] = 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
