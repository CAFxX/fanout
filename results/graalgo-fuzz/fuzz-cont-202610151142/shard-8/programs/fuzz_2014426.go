package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = float64(0.1)
    var v3 uint32 = 2147483648
    v4 := make(map[uint32]uint32)
    var v5 int64 = -1
    v6 := []uint32{(v3 ^ v3), v3}
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 42, F1: true, F2: (fi7 << ((255) & 63))}
            _ = fi8
            _ = v9
        }
        var v10 float64 = v2
        _ = fi7
        _ = v10
    }
    var v11 float64 = ((((1e10 + v2) * (0.1 * 1e10)) + 3.14159) - 3.14159)
    v12 := GenS1{F0: 0, F1: (((!(((true && true) && true) || ((false && false) || true))) && true) && false), F2: 2147483647}
    v2 = (v11 / -2.25)
    fmt.Println(v11, v5, v5)
    var v13 uint8 = 1
    var v14 bool = false
    if true {
        var v15 uint = uint(v13)
        var v16 uint8 = v13
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 9; fi17++ {
            if v14 {
                var v18 int64 = -351
                var v19 int64 = (v18 | v5)
                fmt.Println(v19, v2)
                _ = v18
                _ = v19
            }
            v2 = v2
            v2 = -2.25
            _ = fi17
        }
        v3 -= (v3 | 891)
    }
    v20 := GenS1{F0: -128, F1: v14, F2: 1}
    for fi21 := 0; fi21 < 17; fi21++ {
        v2 += ((v2 / (float64(float32(((v2 - v11) / v11))) - v2)) + 1e10)
        _ = fi21
    }
    var v22 uint16 = 32767
    for fi23 := 0; fi23 < 16; fi23++ {
        v3 = (v3 + (^(v3)))
        v6 = append(v6, (v3 ^ v3))
        _ = fi23
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 uint16 = 42
        v5 -= (v5 - 42)
        if (v14 && (int8(fi24) == int8(v3))) {
            switch (fi24 - fi24) {
            case -9223372036854775808:
                v5 = 696
            case 698:
                v26 := GenS1{F0: 1, F1: (("foo bar" + "") <= ""), F2: fi24}
                _ = v26
            case 2147483647:
                v2 -= v11
            }
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v14, v11)
    v2 = 0.1
    v27 := GenS1{F0: 100, F1: v14, F2: 32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
