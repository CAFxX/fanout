package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }

func main() {
    v1 := []uint32{184, 2147483647, 818, 2147483648}
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    var v4 int8 = -108
    var v5 int8 = v4
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 bool = true
        _ = fi6
        _ = v7
    }
    var v8 uint16 = 955
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v4, v4)
        var v10 uint16 = v8
        _ = fi9
        _ = v10
    }
    var v11 int8 = v5
    var v12 int64 = -1
    _ = len(v2)
    for fi13 := 0; fi13 < 17; fi13++ {
        if false {
            if true {
                if false {
                    v2[fi13] = false
                } else {
                    v12 = (v12 * 679)
                    fmt.Println(v12)
                }
                fmt.Println(v11, v12)
                if ((-(v12)) >= -9223372036854775808) {
                    v14 := GenS1{F0: (v8 / ((v8) | 1)), F1: v8, F2: 1}
                    _ = v14
                }
            }
        }
        _ = fi13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
