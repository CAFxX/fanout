package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint64 }
type GenS2 struct { F0 int8; F1 int64; F2 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = -2.25
    var v3 int32 = 906
    var v4 int = -2147483648
    var v5 float32 = 0.0
    v1[800] = 0
    var v6 int64 = -923
    for fi7 := 0; fi7 < 16; fi7++ {
        v6 = -633
        switch (v4 ^ (v4 * fi7)) {
        case 32768:
            var v8 uint16 = 65535
            _ = v8
        case -2147483648:
            v6 = v6
        }
        v9 := GenS1{F0: 1000, F1: 1, F2: 1}
        _ = fi7
        _ = v9
    }
    v10 := GenS2{F0: -78, F1: (v6 << ((0) & 63)), F2: 32767}
    v11 := GenS2{F0: -1, F1: (v6 % (((v6 / ((870) | 1))) | 1)), F2: 454}
    v12 := GenS2{F0: 42, F1: (v6 ^ v6), F2: 65535}
    v3 += v3
    for fi13 := 0; fi13 < 8; fi13++ {
        v3 = 677
        _ = fi13
    }
    fmt.Println(v4, v5, v5)
    switch v4 {
    case 127:
        v14 := []int{(v4 - v4)}
        _ = v14
    }
    v2 = 3.14159
    var v15 uint16 = 32767
    var v16 float64 = ((1.5 + v2) - v2)
    v17 := GenS1{F0: -32768, F1: 100, F2: 925}
    v3 = (v3 - (-(v3)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
