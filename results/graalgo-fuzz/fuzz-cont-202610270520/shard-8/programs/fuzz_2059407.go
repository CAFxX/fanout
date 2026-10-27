package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int8 = 6
    var v2 uint16 = 0
    var v3 int64 = -1
    var v4 float64 = 0.0
    v5 := []int{4294967295}
    v5 = append(v5, -1)
    var v6 int64 = 901
    _ = len(v5)
    if false {
        switch 136 {
        case -2147483648:
            if false {
                fmt.Println(v6)
                v4 = 1.5
                var v7 uint32 = 896
                _ = v7
            } else {
                fmt.Println(v4, v3)
                var v8 float32 = (float32(v4) * 1.5)
                _ = v8
            }
        default:
            v9 := GenS1{F0: 231}
            _ = v9
        }
        v10 := GenS1{F0: 0}
        _ = v10
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v11 uint16 = (v2 >> ((0) & 63))
    v5 = append(v5, int(v2))
    v1 += (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
