package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[int]int)
    v2 := []int{}
    v3 := make(map[int]int)
    var v4 int64 = 1
    v4 += v4
    v4 += v4
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint16 = 0
        v7 := v3[fi5]
        if true {
            v8 := GenS1{F0: 342}
            _ = v8
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v9 int32 = int32(v4)
    var v10 float64 = float64((1e10 * -2.25))
    var v11 int16 = 32767
    var v12 int64 = v4
    v11 += v11
    var v13 uint = 1
    if ((!(false && true)) || false) {
        v14 := GenS2{F0: v9, F1: true}
        _ = v14
    } else {
        var v15 int16 = -397
        _ = v15
    }
    v2 = append(v2, 13)
    fmt.Println(v11, v12)
    v12 -= (v12 - (v4 >> ((248) & 63)))
    v16 := GenS2{F0: v9, F1: (false && ((v9 | (v9 & v9)) != v9))}
    for fi17 := 0; fi17 < 17; fi17++ {
        v12 -= (v4 >> ((18446744073709551615) & 63))
        v10 = v10
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
