package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 814
    var v3 int = 128
    var v4 float64 = 1e10
    v5 := []int{4294967295, -1, v3, v3}
    v6 := GenS1{F0: 1, F1: (v2 + v2), F2: 75}
    var v7 uint32 = 56
    if false {
        var v8 uint32 = 638
        var v9 int16 = -393
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v3 = 503
        _ = len(v1)
        _ = fi10
    }
    var v11 uint32 = (^(v7))
    var v12 bool = (!false)
    v4 += float64(3.14159)
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
