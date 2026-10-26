package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{}
    var v3 uint16 = 885
    v4 := []int{127, 128, 127, 0}
    var v5 string = ((("\n\t" + "hello") + "a") + "")
    v6 := make(map[string]uint32)
    var v7 float32 = (float32(((-2.25 - 0.0) + (0.0 - 1.5))) / 1.5)
    v7 = 1.5
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: 0, F1: 0, F2: 42}
        _ = fi8
        _ = v9
    }
    switch v5 {
    case "\n\t":
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 int8 = 100
            _ = fi10
            _ = v11
        }
    default:
        var v12 int32 = -965
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: 2, F1: 120, F2: -32768}
        _ = fi13
        _ = v14
    }
    var v15 int64 = 0
    var v16 uint8 = 128
    switch "" {
    case "":
        v17 := GenS2{F0: v5}
        _ = v17
    case "hello":
        var v18 int64 = (^(v15))
        _ = v18
    }
    var v19 uint8 = (v16 & v16)
    fmt.Println(v5, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
