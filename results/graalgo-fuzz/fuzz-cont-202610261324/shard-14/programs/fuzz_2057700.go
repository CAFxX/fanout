package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint8 }

func main() {
    v1 := []int{127, 4294967295, 912, -2147483648}
    var v2 uint = 4294967295
    var v3 int32 = -2147483648
    var v4 bool = true
    var v5 string = "foo bar"
    v6 := GenS1{F0: v4}
    switch -1 {
    case 32768:
        v4 = (42 < 42)
    case -918:
        if len(v1) > 0 {
            v1[0] = 1
        }
    default:
        v1 = append(v1, 32767)
    }
    var v7 uint64 = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: v4}
        _ = fi8
        _ = v9
    }
    v3 += (v3 ^ -1)
    v10 := GenS2{F0: 821, F1: 216}
    fmt.Println(v2, v2)
    var v11 uint = 255
    fmt.Println(v7, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
