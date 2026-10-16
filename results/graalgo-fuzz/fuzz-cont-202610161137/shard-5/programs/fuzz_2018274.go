package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    var v4 int8 = -88
    v5 := GenS1{F0: -2147483648, F1: 70, F2: (v4 - v4)}
    v4 = (v4 / ((v4) | 1))
    switch "\n\t" {
    case "a":
        fmt.Println(v4)
    case "hello":
        if false {
            for fi6 := 0; fi6 < 11; fi6++ {
                v4 = (v4 >> ((386) & 63))
                _ = fi6
            }
            v2[(("hello" + "a") + "a")] = 128
        }
    default:
        fmt.Println(v4)
    }
    switch -1 {
    case 1:
        fmt.Println(v4)
    case 599:
        v7 := GenS1{F0: -1, F1: (v4 * -1), F2: v4}
        _ = v7
    }
    switch ("" + (((("" + "a") + "a") + "") + "a")) {
    case "":
        fmt.Println(v4)
    default:
        if true {
            var v8 uint16 = 0
            _ = v8
        }
    }
    var v9 uint = 2
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 int32 = 568
        fmt.Println(v9)
        v12 := GenS2{F0: 255, F1: 198}
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        v2["hello"] = -1
        _ = fi13
    }
    v14 := GenS1{F0: 2147483647, F1: (v4 & (v4 * 1)), F2: v4}
    v4 = v4
    v4 -= v4
    v9 -= v9
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS2{F0: 2, F1: 117}
        var v17 float32 = float32((0.0 + 0.1))
        if (!(437 < 2)) {
            v18 := GenS3{F0: v4, F1: 304}
            var v19 string = "\n\t"
            _ = v18
            _ = v19
        } else {
            var v20 float32 = (v17 + v17)
            _ = len(v3)
            fmt.Println(fi15, v17, fi15)
            _ = v20
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
