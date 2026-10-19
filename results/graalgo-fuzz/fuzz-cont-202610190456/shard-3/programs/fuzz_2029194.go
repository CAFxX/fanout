package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() uint64 {
    switch "hello" {
    case "":
        var v4 uint32 = 0
        _ = v4
    default:
        switch 0 {
        case 32768:
            var v5 bool = true
            _ = v5
        default:
            var v6 float64 = ((1.5 * float64((-2.25 / 3.14159))) / (-2.25 * 3.14159))
            _ = v6
        }
    }
    var v7 uint32 = 58
    var v8 int32 = 2147483647
    _ = v7
    _ = v8
    return 32
}

func main() {
    var v1 int = 4294967295
    var v2 uint = 201
    var v3 uint64 = 4294967295
    v9 := []int{}
    v3 -= (v3 << ((v2) & 63))
    v3 = (v3 - (v3 % ((42) | 1)))
    v10 := GenS1{F0: 156}
    fmt.Println(v2, v3, v2)
    v11 := GenS1{F0: 1}
    var v12 bool = (!(!(255 > 255)))
    if v12 {
        var v13 uint8 = 128
        v12 = true
        v14 := GenS1{F0: 24}
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
