package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int; F1 int16 }

func main() {
    var v1 int = -560
    var v2 int64 = 1
    var v3 uint16 = 419
    if false {
        var v4 int16 = 0
        v1 = v1
        _ = v4
    } else {
        var v5 float64 = ((-2.25 - -2.25) - 0.1)
        _ = v5
    }
    fmt.Println(v1)
    v6 := GenS1{F0: 86, F1: 1}
    v7 := GenS1{F0: 1, F1: 1}
    var v8 int64 = (v2 << ((42) & 63))
    var v9 int64 = (v2 ^ (-(v2)))
    v8 = v9
    fmt.Println(v9, v9, v3)
    var v10 int64 = v2
    var v11 float32 = (0.0 + (float32((0.0 - float64((-2.25 * (3.14159 * 0.1))))) - ((float32(0.0) * 3.14159) * (-2.25 / -2.25))))
    v10 += (v2 >> ((0) & 63))
    var v12 int = v1
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            v15 := GenS2{F0: -1, F1: 0}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    switch (v1 & v12) {
    case 4294967295:
        v8 -= 9223372036854775807
    case 549:
        v16 := []int{(v12 << ((255) & 63)), 9223372036854775807}
        _ = v16
    default:
        var v17 uint8 = 128
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int64 = (v10 | v10)
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := []int{}
            _ = len(v21)
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
