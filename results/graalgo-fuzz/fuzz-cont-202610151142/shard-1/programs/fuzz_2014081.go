package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 string }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    var v1 bool = false
    var v2 int = 128
    v3 := make(map[uint32]int)
    var v4 string = "hello"
    v5 := []int{v2, (v2 << ((255) & 63)), v2, (v2 * (v2 * -237))}
    _ = len(v5)
    v2 = (-(v2))
    for fi6 := 0; fi6 < 5; fi6++ {
        delete(v3, 864)
        v1 = (42 > 32767)
        _ = fi6
    }
    var v7 uint16 = 923
    switch v2 {
    case 4294967295:
        for fi8 := 0; fi8 < 2; fi8++ {
            if v1 {
                if (true && v1) {
                    var v9 uint64 = 0
                    v10 := GenS1{F0: 171, F1: v7, F2: v4}
                    _ = v10
                    _ = v9
                }
            } else {
                v1 = v1
                if v1 {
                    v11 := GenS2{F0: 42, F1: 113}
                    v7 += v7
                    v12 := GenS2{F0: 666, F1: 0}
                    _ = v11
                    _ = v12
                }
            }
            v1 = v1
            _ = fi8
        }
    case 1:
        v1 = (113 > 2)
    }
    var v13 float64 = (1e10 - ((3.14159 / 1.5) * -2.25))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
