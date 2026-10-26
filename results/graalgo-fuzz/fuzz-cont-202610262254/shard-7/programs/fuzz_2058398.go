package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 float32 = (float32(float64(float32((0.0 - 0.0)))) / 3.14159)
    v2 := make(map[uint32]uint32)
    var v3 int16 = 1
    var v4 uint32 = 24
    switch 96 {
    case 32768:
        var v5 uint16 = 255
        _ = v5
    default:
        v6 := GenS1{F0: 969, F1: 18446744073709551615}
        _ = v6
    }
    v4 = v4
    v4 = 1
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint16 = 0
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v1 = float32(3.14159)
        _ = fi9
    }
    var v10 uint = 1
    fmt.Println(v3)
    v11 := GenS2{F0: (v3 / ((v3) | 1))}
    var v12 int8 = 1
    if (0.0 < -2.25) {
        _ = len(v2)
        fmt.Println(v1)
    }
    var v13 int8 = (v12 >> ((340) & 63))
    if (false && (127 != 127)) {
        var v14 int8 = v12
        var v15 int = -912
        var v16 uint = v10
        _ = v14
        _ = v15
        _ = v16
    }
    if false {
        var v17 uint16 = 2
        switch 127 {
        case 128:
            var v18 uint16 = v17
            _ = v18
        }
        v2[4294967295] = (v4 & 42)
        _ = v17
    }
    v2[(v4 / (((v4 + (v4 | v4))) | 1))] = (^(v4))
    if false {
        v1 = -2.25
        for fi19 := 0; fi19 < 2; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                v21 := GenS3{F0: 558}
                fmt.Println(v3, fi20, v12)
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        v22 := v2[4294967295]
        _ = v22
    }
    v23 := []int{4294967295, 127}
    fmt.Println(v4, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
