package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint32; F1 string; F2 bool }

func main() {
    var v1 int16 = -406
    v2 := make(map[string]uint32)
    var v3 float32 = 1.5
    var v4 int16 = v1
    v5 := make(map[string]bool)
    fmt.Println(v1, v1)
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 float32 = (0.0 + float32(1e10))
        _ = fi6
        _ = v7
    }
    if false {
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := []int{fi8}
            if true {
                fmt.Println(v1, v3, v4)
                if true {
                    v10 := GenS1{F0: 1, F1: ((42 < 1) || false), F2: true}
                    v11 := GenS2{F0: 128}
                    _ = v10
                    _ = v11
                }
                v1 = v1
            }
            _ = fi8
            _ = v9
        }
        var v12 float64 = 1e10
        _ = len(v5)
        _ = v12
    } else {
        var v13 uint = 42
        v1 = -533
        _ = v13
    }
    v4 -= 1000
    var v14 uint = 2
    v15 := GenS3{F0: 0, F1: 42}
    var v16 uint = v14
    v17 := GenS3{F0: 128, F1: -126}
    v1 = (v4 | -1)
    var v18 int32 = 0
    switch "a" {
    case "foo bar":
        var v19 float32 = float32(0.0)
        _ = v19
    default:
        v20 := GenS4{F0: 0, F1: "", F2: (33 <= 42)}
        _ = v20
    }
    var v21 float32 = (-2.25 + v3)
    v22 := GenS3{F0: 32768, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
