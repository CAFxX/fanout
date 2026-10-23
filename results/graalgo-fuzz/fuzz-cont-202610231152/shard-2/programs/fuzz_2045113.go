package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 bool }
func genFn0(p0 int8) uint32 {
    var v6 int = 127
    switch (v6 & v6) {
    case 180:
        fmt.Println(v6, v6)
    }
    if false {
        v7 := GenS1{F0: 100, F1: 529, F2: (2 <= 255)}
        _ = v7
    }
    p0 = 83
    _ = p0
    _ = v6
    return 2147483648
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]uint32)
    var v3 bool = true
    v4 := make(map[uint32]bool)
    var v5 bool = (!v3)
    v8 := []int{}
    v9 := v2[777]
    v5 = true
    switch "a" {
    case "hello":
        v1 += 127
    }
    v3 = (0.1 >= float32(-2.25))
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(v9, v5)
        var v11 float64 = float64((((3.14159 / 1.5) / (1.5 - (((1.5 - 0.1) - (-2.25 * (1.5 + 1.5))) - (((0.1 + ((1e10 / ((3.14159 / 1e10) / ((float32((float64((0.1 * 1.5)) - 3.14159)) / (((0.1 + ((3.14159 / 1e10) / -2.25)) * float32(3.14159)) / (float32(-2.25) - (0.1 - 0.1)))) + 3.14159))) + 0.0)) - 1.5) - 1e10)))) * ((0.1 - ((0.0 + ((((1e10 * (((1.5 + (1.5 - 3.14159)) - 0.1) * (3.14159 / 1.0))) + 3.14159) - 1e10) + 0.0)) / (0.1 + float32(1.5)))) + 0.1)))
        var v12 uint64 = 42
        _ = fi10
        _ = v11
        _ = v12
    }
    delete(v4, genFn0(1))
    for fi13 := 0; fi13 < 5; fi13++ {
        v9 += 1
        _ = fi13
    }
    v9 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
