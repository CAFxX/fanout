package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
type GenS2 struct { F0 string; F1 int8; F2 string }
type GenS3 struct { F0 int64; F1 int32; F2 uint }
func genFn0() string {
    var v7 int64 = 1
    v7 += 9223372036854775807
    _ = v7
    return "a"
}

func main() {
    var v1 int8 = 127
    v2 := []string{}
    var v3 int16 = 0
    var v4 uint32 = 2
    var v5 string = ""
    v6 := make(map[int]uint32)
    var v8 bool = (!false)
    v9 := GenS1{F0: 0, F1: 255, F2: v8}
    v10 := GenS1{F0: v3, F1: 1, F2: true}
    switch -511 {
    case 226:
        v11 := GenS1{F0: (v3 / ((v3) | 1)), F1: 536, F2: v8}
        _ = v11
    case 937:
        v12 := GenS2{F0: ("hello" + v5), F1: (v1 ^ (^(v1))), F2: "a"}
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        delete(v6, int(v3))
        _ = fi13
    }
    v4 = (v4 / ((2) | 1))
    var v14 uint = 652
    for fi15 := 0; fi15 < 9; fi15++ {
        fmt.Println(fi15, v3, v14)
        v16 := GenS3{F0: 754, F1: -1, F2: 4294967295}
        _ = fi15
        _ = v16
    }
    var v17 int8 = v1
    v6[42] = v4
    v3 = (v3 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
