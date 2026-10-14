package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 float64 = float64((-2.25 - (0.1 + 0.1)))
    v2 := []uint32{111, 284}
    var v3 uint = 336
    var v4 int8 = -18
    var v5 int = 2147483647
    var v6 float64 = v1
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v4, v4)
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int64 = 178
            v9 = (v9 - v9)
            v1 = v1
            _ = fi8
            _ = v9
        }
        if true {
            v10 := GenS1{F0: 1, F1: 2147483648}
            _ = v10
        }
        _ = fi7
    }
    fmt.Println(v1, v3, v3)
    var v11 int = int(v3)
    v6 = -2.25
    if true {
        var v12 int32 = -1
        _ = v12
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 uint32 = 0
            _ = fi13
            _ = v14
        }
    }
    if (1e10 < 3.14159) {
        var v15 uint8 = 1
        if ((-(v4)) >= (^(v4))) {
            v4 = (^(v4))
        } else {
            v16 := GenS1{F0: 255, F1: uint32(v15)}
            var v17 float32 = 3.14159
            _ = v16
            _ = v17
        }
        _ = v15
    }
    _ = len(v2)
    if (true || (-9223372036854775808 == 42)) {
        if true {
            if (42 >= 249) {
                switch "\n\t" {
                case "hello":
                    var v18 int64 = 331
                    _ = v18
                }
            } else {
                v19 := GenS1{F0: 927, F1: 65535}
                _ = v19
            }
            var v20 uint16 = 0
            _ = v20
        }
        var v21 uint = (v3 + 208)
        _ = v21
    }
    v1 = v1
    var v22 uint64 = 670
    v2 = append(v2, 2)
    var v23 uint16 = 434
    v5 = -317
    for fi24 := 0; fi24 < 16; fi24++ {
        fmt.Println(v1, v23, v11)
        fmt.Println(v5, v5)
        v25 := GenS1{F0: (v23 * v23), F1: 894}
        _ = fi24
        _ = v25
    }
    if ((v3 >> (((v3 - v3)) & 63)) >= v3) {
        var v26 int8 = (^(v4))
        v23 -= (v23 ^ (v23 | 2))
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
