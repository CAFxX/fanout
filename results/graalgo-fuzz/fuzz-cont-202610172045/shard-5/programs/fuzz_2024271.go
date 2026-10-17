package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := []int{0, 9223372036854775807, 32768}
    var v2 int8 = 42
    var v3 uint16 = 960
    var v4 uint64 = 18446744073709551615
    v5 := make(map[int]bool)
    for fi6 := 0; fi6 < 18; fi6++ {
        v4 = (v4 % ((18446744073709551615) | 1))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v7 := GenS1{F0: 696}
        _ = fi6
        _ = v7
    }
    var v8 float64 = 0.1
    if true {
        v3 = 255
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 float64 = (v8 - 1.5)
            _ = fi9
            _ = v10
        }
        v8 -= ((v8 + 0.1) * 3.14159)
    } else {
        v11 := GenS2{F0: 42}
        _ = v11
    }
    var v12 uint = 0
    v13 := GenS3{F0: 255, F1: v3}
    var v14 float32 = 1e10
    fmt.Println(v2, v12)
    switch "hello" {
    case "foo bar":
        if ((true && true) || (!(!false))) {
            if false {
                v8 = 0.0
            }
            v5[32767] = false
        }
    case "hello":
        if false {
            var v15 int = 9223372036854775807
            var v16 int = v15
            _ = v15
            _ = v16
        } else {
            v2 = (v2 << (((v12 / ((v12) | 1))) & 63))
        }
    case "a":
        fmt.Println(v2, v8, v3)
    }
    var v17 int8 = (-(v2))
    var v18 uint32 = 236
    if ((v4 | uint64(v18)) >= v4) {
        v4 += (v4 | v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
