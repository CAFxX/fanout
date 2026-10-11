package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }

func main() {
    v1 := []uint32{4294967295}
    var v2 uint32 = 4294967295
    v3 := []string{(("foo bar" + "a") + "\n\t"), ("" + "\n\t"), "a"}
    var v4 uint8 = 128
    v5 := make(map[int]uint32)
    var v6 string = ""
    v7 := GenS1{F0: 0, F1: 92}
    switch 128 {
    case 32767:
        _ = len(v3)
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        _ = len(v1)
        fmt.Println(v2, v4, fi8)
        _ = fi8
    }
    var v9 int16 = 1
    delete(v5, -9223372036854775808)
    v6 = v6
    fmt.Println(v6)
    switch 4294967295 {
    case 127:
        var v10 int16 = 1
        _ = v10
    case -845:
        if (!false) {
            if (65535 != 886) {
                var v11 int = 0
                if (342 < v2) {
                    var v12 int64 = 0
                    if (((^(v4)) == v4) && (32767 != 72)) {
                        v11 = v11
                        var v13 float64 = 1.5
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 15; fi14++ {
                        var v15 int64 = -658
                        _ = fi14
                        _ = v15
                    }
                    _ = v12
                }
                _ = v11
            }
        } else {
            v16 := v5[42]
            v2 -= (v16 * v2)
            _ = v16
        }
    case 0:
        v17 := GenS1{F0: 2, F1: 1}
        _ = v17
    default:
        v3 = append(v3, "hello")
    }
    v9 += (-(v9))
    fmt.Println(v2, v4, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
