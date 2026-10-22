package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
func genFn0(p0 int64) int8 {
    var v5 int16 = 0
    v5 = v5
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(v5, fi6)
        _ = fi6
    }
    p0 = (p0 >> ((980) & 63))
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    v4 := []string{}
    v7 := GenS1{F0: 512, F1: false}
    switch "a" {
    case "hello":
        v4 = append(v4, "")
    case "a":
        var v8 bool = true
        _ = v8
    default:
        v1 -= 1.5
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        switch fi9 {
        case -1:
            v1 -= 0.0
        }
        var v10 string = "hello"
        v1 -= ((float64((float32(v1) * 0.1)) / (v1 - 1.5)) - 1.5)
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 0
    var v12 int = 370
    v13 := GenS1{F0: 32767, F1: ((false && (42 > 753)) || (42 >= 100))}
    v12 = v12
    var v14 int8 = 127
    var v15 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
