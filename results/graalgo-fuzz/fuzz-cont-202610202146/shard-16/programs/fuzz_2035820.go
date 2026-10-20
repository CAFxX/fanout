package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 2
    var v3 float32 = 1e10
    var v4 string = ""
    v1[42] = uint32(v2)
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint8 = 42
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 42
    fmt.Println(v4, v3)
    v8 := GenS1{F0: (v2 / (((v2 * (v2 << ((84) & 63)))) | 1)), F1: 582, F2: (v7 + (v7 ^ v7))}
    v9 := []int{}
    var v10 int16 = -32768
    for fi11 := 0; fi11 < 12; fi11++ {
        v4 = v4
        _ = fi11
    }
    var v12 uint8 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
