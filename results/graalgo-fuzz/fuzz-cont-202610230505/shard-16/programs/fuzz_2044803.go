package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    var v1 bool = (!false)
    var v2 bool = ((400 != -1) && (457 >= 4294967295))
    var v3 float32 = 1e10
    v3 += (v3 / ((v3 / 1e10) - (v3 * v3)))
    v1 = true
    v4 := []int{}
    _ = len(v4)
    _ = len(v4)
    v5 := GenS1{F0: -852, F1: 32767}
    v6 := GenS1{F0: -609, F1: 32768}
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            v1 = (32767 < 990)
            _ = fi8
        }
        v1 = v1
        _ = fi7
    }
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
