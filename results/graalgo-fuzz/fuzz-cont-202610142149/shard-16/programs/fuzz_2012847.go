package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 float32 = 0.0
    v4 := make(map[string]int)
    fmt.Println(v3)
    v3 -= (float32((3.14159 + (3.14159 / -2.25))) / (float32(1.5) / v3))
    if false {
        var v5 uint32 = 1
        switch -822 {
        case 0:
            v6 := v1[2147483648]
            _ = v6
        default:
            for fi7 := 0; fi7 < 12; fi7++ {
                v3 = (0.1 / (v3 * v3))
                v8 := GenS1{F0: true, F1: true, F2: 65535}
                _ = fi7
                _ = v8
            }
        }
        _ = v5
    }
    var v9 uint8 = 255
    var v10 int32 = -1
    var v11 uint32 = 0
    switch 2147483647 {
    case 128:
        for fi12 := 0; fi12 < 2; fi12++ {
            v9 -= (v9 << ((1) & 63))
            if (18446744073709551615 > 18446744073709551615) {
                v13 := GenS2{F0: 1, F1: 42}
                v10 = (v10 & (v10 - (v10 >> ((0) & 63))))
                v10 = v10
                _ = v13
            } else {
                var v14 int = (fi12 % (((fi12 ^ fi12)) | 1))
                _ = v14
            }
            _ = fi12
        }
    case -2147483648:
        switch -820 {
        case -87:
            v9 = (^(v9))
        default:
            v15 := GenS2{F0: 42, F1: v9}
            _ = v15
        }
    case 1:
        v11 += (v11 % ((v11) | 1))
    }
    fmt.Println(v3, v3, v9)
    var v16 int64 = 1
    var v17 uint = 65535
    v16 = v16
    for fi18 := 0; fi18 < 7; fi18++ {
        v16 = -1
        _ = fi18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
