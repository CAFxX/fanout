package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint16 }

func main() {
    var v1 int16 = 366
    var v2 bool = false
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    var v5 uint16 = 0
    v6 := GenS1{F0: 128}
    var v7 float32 = -2.25
    v8 := GenS2{F0: 1}
    var v9 bool = (-2147483648 <= -1)
    v2 = ((^(v5)) == (^(v5)))
    v10 := GenS1{F0: 0}
    v11 := GenS3{F0: v2, F1: (v5 * v5)}
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int32 = 2147483647
        var v14 string = "foo bar"
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v7, v1)
    for fi15 := 0; fi15 < 7; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 uint64 = 42
            _ = fi16
            _ = v17
        }
        var v18 uint64 = 18446744073709551615
        _ = fi15
        _ = v18
    }
    if v9 {
        v2 = true
        _ = len(v4)
    } else {
        switch -933 {
        case -86:
            fmt.Println(v7, v1, v5)
        case -2147483648:
            v1 += v1
        default:
            v19 := []int{}
            _ = v19
        }
        v20 := []int{-383, -1, 495}
        _ = v20
    }
    var v21 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
