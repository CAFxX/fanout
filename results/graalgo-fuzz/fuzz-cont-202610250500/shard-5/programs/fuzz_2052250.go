package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float32 = float32(1.5)
    var v2 uint16 = 699
    v3 := []int{-9223372036854775808, 249, 2147483647}
    v4 := []uint32{}
    var v5 bool = ((((true || true) || (!false)) && false) && (false && (-2147483648 >= -966)))
    var v6 uint64 = 2
    v7 := GenS1{F0: 603}
    v1 = v1
    v1 = 3.14159
    var v8 int16 = 1000
    if false {
        var v9 float32 = float32(1e10)
        switch "" {
        case "hello":
            if len(v4) > 0 {
                v4[0] = 1
            }
        }
        _ = v9
    } else {
        var v10 float32 = v1
        var v11 uint16 = (v2 / (((^(v2))) | 1))
        var v12 uint16 = 220
        _ = v10
        _ = v11
        _ = v12
    }
    v1 -= v1
    if v5 {
        if len(v3) > 0 {
            v3[0] = 104
        }
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 uint8 = 127
            _ = fi13
            _ = v14
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
