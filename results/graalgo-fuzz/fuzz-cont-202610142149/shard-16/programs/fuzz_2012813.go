package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 int16 = 32767
    v4 := []int{0, -2147483648}
    var v5 int16 = v3
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint32 = 0
        if ((v3 + v5) <= v5) {
            fmt.Println(v7, v5, v7)
        }
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 int32 = 0
            v10 := GenS1{F0: 127, F1: fi6, F2: "a"}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v11 uint64 = 925
    var v12 uint32 = 42
    switch 377 {
    case 4294967295:
        var v13 uint16 = 386
        _ = v13
    case 144:
        var v14 int8 = -6
        _ = v14
    case 1:
        for fi15 := 0; fi15 < 7; fi15++ {
            v12 -= v12
            v12 = (v12 - 2147483648)
            _ = fi15
        }
    default:
        var v16 int32 = 2147483647
        _ = v16
    }
    var v17 float32 = 1.5
    var v18 bool = false
    v5 -= (v3 << ((0) & 63))
    var v19 uint16 = 0
    v20 := v2[32768]
    var v21 string = "\n\t"
    switch v21 {
    case "a":
        switch -2147483648 {
        case -848:
            switch v21 {
            case "a":
                v4 = append(v4, -2147483648)
            default:
                for fi22 := 0; fi22 < 14; fi22++ {
                    v12 -= uint32(v19)
                    _ = fi22
                }
            }
        default:
            delete(v1, -2147483648)
        }
    default:
        var v23 int64 = 0
        _ = v23
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
