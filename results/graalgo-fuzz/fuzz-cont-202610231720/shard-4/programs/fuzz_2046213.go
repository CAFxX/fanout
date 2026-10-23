package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int64 = 0
    var v2 float32 = -2.25
    var v3 int = 128
    v1 -= v1
    var v4 uint8 = 118
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int32 = 0
        v7 := []int{-1, 4294967295}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 float32 = 0.0
        var v10 string = "hello"
        v10 = (((v10 + "foo bar") + "foo bar") + "")
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v4)
    v1 = -84
    fmt.Println(v1, v4, v4)
    v4 = (v4 - (v4 / (((v4 & 237)) | 1)))
    if (136 >= 42) {
        v11 := []int{(v3 & v3), v3, 678, 28}
        v4 = v4
        if true {
            v12 := GenS1{F0: 0}
            _ = v12
        }
        _ = v11
    } else {
        var v13 int8 = 42
        v14 := []int{(-(v3)), (v3 - (v3 >> ((255) & 63))), (v3 >> ((1) & 63))}
        v15 := GenS2{F0: 1}
        _ = v13
        _ = v14
        _ = v15
    }
    switch 32768 {
    case -1:
        v4 -= (v4 % (((v4 * v4)) | 1))
    default:
        v3 -= (v3 - (v3 * v3))
    }
    v16 := []int{}
    var v17 float32 = ((v2 + (3.14159 * 1e10)) / (-2.25 + v2))
    var v18 int32 = -462
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
