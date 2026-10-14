package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint32 }

func main() {
    v1 := []uint32{2147483648}
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 uint = 42
    var v5 int8 = -1
    var v6 float32 = (0.1 / (1.5 / -2.25))
    fmt.Println(v4)
    if (42 > 1000) {
        switch 2147483647 {
        case 32768:
            var v7 uint64 = 18446744073709551615
            _ = v7
        case 127:
            v8 := GenS1{F0: 0, F1: 0, F2: 268}
            _ = v8
        }
        v5 += (v5 >> ((35) & 63))
    }
    fmt.Println(v5)
    v9 := GenS1{F0: 0, F1: -977, F2: 2147483648}
    v10 := GenS1{F0: 42, F1: 1000, F2: 65535}
    if false {
        if false {
            for fi11 := 0; fi11 < 12; fi11++ {
                v12 := GenS1{F0: 4294967295, F1: -577, F2: 134}
                if true {
                    v5 += (v5 * -1)
                    v6 = (v6 + (0.0 / v6))
                    v3[fi11] = 2147483648
                }
                _ = fi11
                _ = v12
            }
            v13 := v3[-688]
            _ = v13
        } else {
            var v14 int16 = 179
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 float32 = 1.5
        var v17 bool = false
        _ = fi15
        _ = v16
        _ = v17
    }
    if (false && (282 > 502)) {
        _ = len(v1)
        if (false && (42 <= 42)) {
            _ = len(v3)
            var v18 uint32 = 957
            _ = v18
        } else {
            v19 := v2[0]
            _ = v19
        }
        for fi20 := 0; fi20 < 14; fi20++ {
            v1 = append(v1, 0)
            v6 += (0.1 * 1.5)
            v1 = append(v1, 2147483648)
            _ = fi20
        }
    } else {
        v5 = v5
        var v21 int8 = v5
        if ((!false) || true) {
            v6 += v6
            v1 = append(v1, 631)
        } else {
            v1 = append(v1, 94)
            var v22 uint = v4
            _ = v22
        }
        _ = v21
    }
    v5 = (-(v5))
    var v23 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
