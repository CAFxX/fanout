package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 2
    var v3 float64 = -2.25
    v2 -= 255
    var v4 uint8 = v2
    if true {
        v5 := GenS1{F0: 4294967295, F1: 0}
        v2 -= 127
        var v6 float32 = 0.1
        _ = v5
        _ = v6
    }
    v7 := []int{4294967295, 32768, 625}
    var v8 string = "foo bar"
    v2 -= (v4 ^ (v4 - v4))
    v9 := GenS1{F0: 65535, F1: 0}
    fmt.Println(v2)
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS1{F0: 18446744073709551615, F1: 2}
        _ = fi10
        _ = v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
