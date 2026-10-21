package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{-666, 4294967295}
    var v3 float32 = (1e10 / 1e10)
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]string)
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(v6, v3, v3)
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 bool = (42 == 124)
        for fi10 := 0; fi10 < 17; fi10++ {
            v6 += (v6 ^ v6)
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    v11 := GenS1{F0: (v6 >> ((18446744073709551615) & 63))}
    var v12 uint32 = 1
    var v13 uint = 0
    v3 -= (v3 * (float32((0.1 - (1.5 / 1.5))) - ((v3 / ((v3 * v3) - (3.14159 + v3))) / v3)))
    var v14 bool = ((v12 / (((v12 ^ (v12 % ((v12) | 1)))) | 1)) > 2)
    v15 := GenS1{F0: v6}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
