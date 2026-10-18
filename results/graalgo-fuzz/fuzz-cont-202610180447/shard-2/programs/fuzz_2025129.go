package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint64 }
type GenS2 struct { F0 int64; F1 int16 }

func main() {
    v1 := []uint32{2147483647, 336, 212}
    v2 := []uint32{51, 393, 717, 42}
    v3 := make(map[uint32]bool)
    var v4 uint8 = 128
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := v3[12]
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        v3[155] = false
        _ = fi7
    }
    switch 42 {
    case 32767:
        fmt.Println(v4)
    }
    var v8 uint16 = 258
    v9 := GenS1{F0: (!(false && (0.1 > (((1.5 * 3.14159) * 3.14159) + 1e10)))), F1: 2147483647, F2: 18446744073709551615}
    fmt.Println(v4, v8)
    var v10 int8 = 85
    v11 := GenS2{F0: 1, F1: 0}
    fmt.Println(v10)
    v4 -= (v4 & v4)
    v3[407] = true
    v10 = (v10 | v10)
    v12 := GenS2{F0: -1, F1: 1}
    var v13 uint16 = v8
    v10 = (-(v10))
    var v14 uint = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
