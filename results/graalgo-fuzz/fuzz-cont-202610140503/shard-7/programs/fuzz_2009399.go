package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint8; F1 int8; F2 uint16 }

func main() {
    var v1 int64 = 0
    var v2 int8 = 53
    v3 := []string{"hello"}
    for fi4 := 0; fi4 < 17; fi4++ {
        v3 = append(v3, "hello")
        v1 += v1
        if false {
            v2 += -94
            for fi5 := 0; fi5 < 6; fi5++ {
                fmt.Println(v2, v2, v2)
                v2 -= v2
                _ = fi5
            }
        }
        _ = fi4
    }
    v6 := GenS1{F0: 633, F1: 466}
    var v7 bool = false
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint32 = 4294967295
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 2
    v11 := GenS2{F0: 67, F1: v2, F2: 973}
    v10 += (v10 >> ((2) & 63))
    if len(v3) > 0 {
        v3[0] = "a"
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v12 float64 = 0.0
    v13 := GenS2{F0: (^(v10)), F1: (v2 ^ (-(v2))), F2: 255}
    fmt.Println(v7, v10, v1)
    v3 = append(v3, "")
    for fi14 := 0; fi14 < 4; fi14++ {
        v7 = (811 < 231)
        _ = fi14
    }
    if v7 {
        var v15 int32 = -1
        _ = v15
    } else {
        switch 4294967295 {
        case 2147483647:
            if true {
                v16 := GenS2{F0: v10, F1: v2, F2: 0}
                switch "" {
                case "hello":
                    v2 += (v2 | v2)
                case "foo bar":
                    v10 += 11
                }
                _ = v16
            }
        }
        var v17 float32 = (-2.25 * (((0.1 + float32((v12 + 1e10))) * 1.5) / (1.5 * float32(0.1))))
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
