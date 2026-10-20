package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint8 }
func genFn0(p0 int32, p1 int16) bool {
    if false {
        var v7 int32 = p0
        _ = v7
    } else {
        if false {
            var v8 uint16 = 42
            _ = v8
        } else {
            fmt.Println(p1, p0)
            var v9 uint32 = 2147483647
            p1 = int16(v9)
            _ = v9
        }
    }
    switch -902 {
    case 9223372036854775807:
        v10 := []int{32767, 1, -429, -864}
        _ = v10
    case 466:
        v11 := []int{32768, 0, -9223372036854775808}
        _ = v11
    default:
        p0 = -1
    }
    p1 = 1000
    fmt.Println(p1, p0)
    v12 := []int{836, -1}
    _ = p0
    _ = p1
    _ = v12
    return false
}

func main() {
    var v1 int16 = -485
    var v2 float32 = 1e10
    v3 := []uint32{}
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 128
    v6 := make(map[uint32]int)
    v5 -= (v5 % ((v5) | 1))
    v5 = (v5 / ((v5) | 1))
    v13 := GenS1{F0: 159, F1: 551, F2: (v5 / ((v5) | 1))}
    delete(v4, 1)
    var v14 float32 = v2
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 string = "\n\t"
        var v17 int8 = 0
        _ = fi15
        _ = v16
        _ = v17
    }
    v5 += v5
    fmt.Println(v1, v5)
    var v18 uint = 178
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
