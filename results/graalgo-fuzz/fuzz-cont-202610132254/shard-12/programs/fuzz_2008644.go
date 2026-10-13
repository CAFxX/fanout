package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int; F1 int8; F2 bool }

func main() {
    var v1 uint8 = 72
    var v2 uint = 607
    var v3 string = "foo bar"
    var v4 int = 127
    v5 := []int{v4, v4, v4, -33}
    fmt.Println(v1)
    for fi6 := 0; fi6 < 6; fi6++ {
        v3 = v3
        _ = fi6
    }
    v5 = append(v5, v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 float64 = float64((0.1 / float32(1.5)))
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int32 = -2147483648
        v4 += v4
        _ = fi8
        _ = v9
    }
    fmt.Println(v2)
    if len(v5) > 0 {
        v5[0] = (v4 / ((v4) | 1))
    }
    var v10 string = (v3 + "\n\t")
    v10 = v3
    v1 = (v1 << (((v2 / ((v2) | 1))) & 63))
    v3 = v10
    switch v4 {
    case 32768:
        var v11 bool = true
        _ = v11
    case 2147483647:
        v12 := GenS1{F0: 42, F1: -40}
        _ = v12
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println(v4)
    if true {
        v13 := GenS2{F0: (v4 ^ v4), F1: 83, F2: false}
        _ = len(v5)
        _ = v13
    }
    var v14 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
