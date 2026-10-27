package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint64; F1 int8 }

func main() {
    var v1 uint8 = 106
    var v2 int64 = 143
    var v3 int = 4294967295
    v4 := make(map[uint32]uint32)
    if (!false) {
        v5 := GenS1{F0: 42, F1: true}
        _ = v5
    }
    var v6 uint = 42
    var v7 uint16 = uint16(v6)
    for fi8 := 0; fi8 < 7; fi8++ {
        v7 -= v7
        v7 += v7
        _ = fi8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        switch -2147483648 {
        case 42:
            v3 += (fi9 >> ((763) & 63))
        default:
            delete(v4, 138)
        }
        v10 := GenS1{F0: 45, F1: (!(485 < (v2 & -1)))}
        var v11 float32 = float32(3.14159)
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = 2
    var v13 bool = (!true)
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v12)
        _ = fi14
    }
    v15 := GenS2{F0: 18446744073709551615, F1: 127}
    v12 = 32767
    fmt.Println(v12, v6, v12)
    v16 := GenS2{F0: 741, F1: 100}
    v7 = (v7 / ((v12) | 1))
    var v17 bool = (!((-115 == -4) && ((0.0 * (-2.25 - 0.0)) == float64(1e10))))
    if ("hello" < "a") {
        v7 += (v7 << ((4294967295) & 63))
        v18 := GenS2{F0: 569, F1: 0}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
