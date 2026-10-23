package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint8; F1 int }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]int)
    var v3 float64 = 0.0
    v4 := GenS1{F0: 685, F1: -718}
    v1 = v1
    v5 := GenS1{F0: 255, F1: 863}
    v3 += v3
    v6 := GenS2{F0: 255, F1: -9223372036854775808}
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint16 = 2
        _ = fi7
        _ = v8
    }
    v9 := []int{0, 221, 32768, 128}
    v2[820] = -268
    if true {
        v2[1] = 1
        v1 = v1
    } else {
        v1 = "\n\t"
        var v10 int = 586
        _ = v10
    }
    var v11 float64 = v3
    v12 := v2[2]
    var v13 uint64 = 1
    if false {
        v13 = 2
        v9 = append(v9, (v12 - (v12 & v12)))
    } else {
        var v14 int16 = int16(v13)
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
