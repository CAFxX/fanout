package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 uint = 18446744073709551615
    v3 = (v3 % (((v3 << (((v3 % ((255) | 1))) & 63))) | 1))
    v4 := GenS1{F0: -5, F1: 425, F2: 1000}
    v5 := GenS2{F0: -1, F1: ((413 > 8) || true), F2: 32767}
    v3 = v3
    var v6 int32 = 277
    switch "hello" {
    case "foo bar":
        v3 += (v3 | (v3 - v3))
    default:
        if false {
            v7 := v2[4294967295]
            if (!false) {
                var v8 uint32 = 65535
                var v9 int32 = v6
                _ = v8
                _ = v9
            } else {
                v10 := []int{2147483647}
                var v11 int = -992
                v3 += (v3 % ((4294967295) | 1))
                _ = v10
                _ = v11
            }
            if (1 <= 506) {
                var v12 uint8 = 185
                _ = v12
            }
            _ = v7
        } else {
            var v13 int16 = 469
            switch 32768 {
            case -970:
                v14 := GenS1{F0: 100, F1: -1, F2: (v13 * v13)}
                _ = v14
            }
            _ = v13
        }
    }
    if false {
        v15 := GenS2{F0: 1000, F1: (!(!(0 <= 42))), F2: 32767}
        _ = v15
    }
    var v16 uint = 540
    var v17 int8 = 0
    v17 -= -128
    fmt.Println(v17)
    v16 += 42
    fmt.Println(v17)
    fmt.Println(v16, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
