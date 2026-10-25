package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    var v1 int16 = 32767
    var v2 uint32 = 555
    v3 := make(map[int]int)
    if true {
        var v4 float64 = 1.5
        v5 := GenS1{F0: -1, F1: 100}
        _ = v4
        _ = v5
    } else {
        fmt.Println(v1, v2)
    }
    v6 := []int{1, -9223372036854775808, 672}
    for fi7 := 0; fi7 < 13; fi7++ {
        v2 = v2
        _ = fi7
    }
    var v8 uint32 = (v2 / (((^(v2))) | 1))
    v8 = 131
    var v9 string = ""
    fmt.Println(v8, v9)
    v2 = (v8 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v8
    _ = v9
}
