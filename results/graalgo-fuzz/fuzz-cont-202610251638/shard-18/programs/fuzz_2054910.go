package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint32 }
func genFn0() uint64 {
    var v7 int64 = 42
    for fi8 := 0; fi8 < 14; fi8++ {
        v7 += (^(v7))
        _ = fi8
    }
    _ = v7
    return 412
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 bool = true
    var v5 uint32 = 923
    var v6 bool = ((((true || v4) && (uint64(v5) > 1)) && v4) && false)
    var v9 uint32 = v5
    v10 := GenS1{F0: int32(v9), F1: 42, F2: 1}
    var v11 uint16 = 42
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 uint8 = (v1 / (((v1 / (((v1 >> ((1) & 63))) | 1))) | 1))
        _ = fi12
        _ = v13
    }
    var v14 uint64 = 619
    var v15 uint64 = v14
    var v16 int = 257
    var v17 float32 = (float32(1e10) / 1e10)
    var v18 float64 = 3.14159
    if (((false && false) || ("hello" <= "foo bar")) && (v17 != ((v17 * v17) + (((v17 / v17) / 1.5) - v17)))) {
        v19 := []int{-315, (v16 << ((255) & 63))}
        switch "hello" {
        case "hello":
            fmt.Println(v15, v9, v11)
        default:
            v20 := GenS1{F0: -2147483648, F1: v5, F2: v9}
            _ = v20
        }
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
