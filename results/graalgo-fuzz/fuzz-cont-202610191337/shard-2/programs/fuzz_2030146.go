package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int64; F1 string; F2 bool }
type GenS4 struct { F0 int32; F1 uint16; F2 uint16 }
func genFn0() uint {
    v4 := GenS1{F0: 42}
    var v5 int64 = 434
    fmt.Println(v5)
    if (-692 >= -1) {
        v6 := GenS2{F0: 42, F1: 0}
        _ = v6
    }
    _ = v4
    _ = v5
    return 283
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = -2.25
    var v3 uint32 = 551
    if true {
        v2 = 1e10
    } else {
        v7 := []int{-913, -917, 128, 4294967295}
        v3 -= v3
        _ = v7
    }
    v3 = (v3 - v3)
    v8 := v1[42]
    v2 = float64(-2.25)
    v8 = v8
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            v3 = v3
            _ = fi10
        }
        _ = fi9
    }
    v11 := GenS3{F0: 9223372036854775807, F1: ("" + "\n\t"), F2: true}
    v12 := GenS4{F0: 2147483647, F1: 255, F2: 816}
    for fi13 := 0; fi13 < 1; fi13++ {
        switch (v8 + v8) {
        case "":
            v14 := v1[(-(fi13))]
            _ = v14
        case "a":
            v3 = (^(v3))
        default:
            v15 := GenS2{F0: 1, F1: 2}
            _ = v15
        }
        if (0 < 4) {
            fmt.Println(v8)
            _ = len(v1)
            fmt.Println(v3, v2)
        }
        _ = fi13
    }
    v8 = v8
    if ((v3 != v3) || false) {
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS1{F0: 630}
            _ = fi16
            _ = v17
        }
        v2 += v2
        fmt.Println(v8)
    }
    var v18 int = 4294967295
    fmt.Println(v2, v2, v2)
    var v19 uint = uint(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
}
