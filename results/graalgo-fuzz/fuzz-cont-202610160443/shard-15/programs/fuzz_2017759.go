package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -732
    var v3 int8 = 127
    v4 := []int{127, -9223372036854775808, 4294967295}
    fmt.Println(v3, v2)
    var v5 float64 = (-2.25 / 1e10)
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: -1, F1: ("hello" + "\n\t")}
        _ = fi6
        _ = v7
    }
    v8 := v1[4294967295]
    v3 = v3
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int8 = (v3 >> ((255) & 63))
        _ = fi9
        _ = v10
    }
    var v11 int64 = int64(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
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
