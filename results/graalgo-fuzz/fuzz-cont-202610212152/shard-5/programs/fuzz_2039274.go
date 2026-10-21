package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 3.14159
    var v3 float64 = -2.25
    var v4 uint8 = 127
    v5 := make(map[int]uint32)
    v6 := make(map[int]bool)
    var v7 uint16 = 42
    v7 = 1
    switch 128 {
    case 9223372036854775807:
        v2 += 3.14159
    case -149:
        for fi8 := 0; fi8 < 11; fi8++ {
            v2 += (float32(v3) * v2)
            switch (fi8 % (((fi8 >> ((609) & 63))) | 1)) {
            case 42:
                v9 := []int{}
                _ = v9
            case 32768:
                v10 := []int{fi8, fi8}
                _ = v10
            case -9223372036854775808:
                v11 := []int{32767}
                _ = v11
            }
            _ = fi8
        }
    default:
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 string = ("hello" + ("foo bar" + "a"))
            var v14 int64 = 460
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    v15 := v5[-9223372036854775808]
    v16 := []int{-9223372036854775808, 127, 680, 613}
    v4 = v4
    _ = len(v6)
    v17 := v5[42]
    v18 := GenS1{F0: 1, F1: "\n\t"}
    var v19 float32 = float32((-2.25 / v3))
    if len(v16) > 0 {
        v16[0] = 762
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
