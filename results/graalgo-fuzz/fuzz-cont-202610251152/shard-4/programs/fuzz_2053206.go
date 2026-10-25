package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint16 = 689
    v2 := make(map[int]uint32)
    v3 := []uint32{42}
    if true {
        v4 := GenS1{F0: 65535, F1: true}
        _ = v4
    } else {
        fmt.Println(v1)
    }
    var v5 int16 = -1
    var v6 float64 = (0.0 / 1.5)
    for fi7 := 0; fi7 < 15; fi7++ {
        v1 += (v1 << ((1) & 63))
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 int = fi7
            v10 := GenS1{F0: 1, F1: (1 == 100)}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    var v11 float32 = (-2.25 / float32((v6 * (v6 + 0.0))))
    if false {
        v12 := v2[9223372036854775807]
        v1 += v1
        _ = v12
    } else {
        v1 = v1
        _ = len(v3)
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            switch int(v1) {
            case 128:
                var v15 int64 = 42
                _ = v15
            case 2147483647:
                var v16 uint16 = 32767
                _ = v16
            case -330:
                v1 = (v1 / (((v1 >> ((377) & 63))) | 1))
            default:
                var v17 uint32 = 468
                _ = v17
            }
            if true {
                v2[(fi13 | (fi13 + fi13))] = 65535
                var v18 int32 = 1
                v19 := GenS1{F0: 289, F1: false}
                _ = v18
                _ = v19
            } else {
                v20 := GenS1{F0: 65535, F1: (v5 < (v5 ^ v5))}
                fmt.Println(v1)
                fmt.Println(fi13)
                _ = v20
            }
            v2[fi14] = 477
            _ = fi14
        }
        _ = fi13
    }
    var v21 uint32 = 368
    var v22 bool = (4294967295 >= 18446744073709551615)
    delete(v2, 718)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v5
    _ = v6
}
