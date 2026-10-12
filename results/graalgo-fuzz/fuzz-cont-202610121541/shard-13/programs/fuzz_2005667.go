package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int }

func main() {
    var v1 bool = true
    var v2 string = ("\n\t" + "")
    v3 := make(map[int]string)
    var v4 uint64 = 347
    v5 := make(map[string]uint32)
    var v6 int = -318
    for fi7 := 0; fi7 < 4; fi7++ {
        if ((fi7 | v6) != fi7) {
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(v2)
                _ = fi8
            }
            var v9 uint32 = 2
            v10 := GenS1{F0: v4, F1: 2}
            _ = v10
            _ = v9
        } else {
            if (2 < 4294967295) {
                v6 -= -970
                v6 += -9223372036854775808
                var v11 bool = false
                _ = v11
            }
        }
        _ = fi7
    }
    v12 := GenS2{F0: (v6 / (((v6 / ((v6) | 1))) | 1))}
    if false {
        v2 = (v2 + ("" + v2))
        var v13 uint32 = 952
        v13 -= v13
        _ = v13
    } else {
        v1 = v1
        var v14 uint16 = 42
        _ = v14
    }
    v6 = v6
    var v15 int32 = -2147483648
    var v16 uint16 = 0
    var v17 float64 = 3.14159
    v18 := GenS1{F0: (v4 - v4), F1: 334}
    v19 := GenS1{F0: 355, F1: 11}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
