package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 string = "a"
    var v4 uint = 1
    switch 127 {
    case 42:
        v4 = 42
    case -672:
        v3 = v3
    case 971:
        var v5 uint8 = 127
        _ = v5
    default:
        switch 43 {
        case 32768:
            v4 += 909
        }
    }
    var v6 bool = true
    var v7 float64 = 1.5
    v8 := GenS1{F0: v4, F1: 42}
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint64 = 659
        var v11 uint32 = 42
        v12 := []int{(fi9 & (fi9 % ((fi9) | 1)))}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 uint16 = 255
    v4 = v4
    for fi14 := 0; fi14 < 3; fi14++ {
        v6 = (v6 || true)
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
