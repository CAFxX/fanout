package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }

func main() {
    var v1 uint = 0
    var v2 uint64 = 473
    var v3 int8 = 1
    var v4 uint32 = 4294967295
    var v5 uint8 = 0
    var v6 float64 = 1.5
    var v7 float32 = (0.1 / 1e10)
    var v8 uint64 = (v2 >> (((v1 ^ v1)) & 63))
    switch 4294967295 {
    case 127:
        if (int64(v1) >= -1) {
            if false {
                if false {
                    var v9 int16 = 1
                    _ = v9
                } else {
                    v3 += (v3 << (((^(v1))) & 63))
                }
                v3 = int8(v2)
                switch -840 {
                case -9223372036854775808:
                    var v10 uint = v1
                    _ = v10
                case 2147483647:
                    fmt.Println(v7)
                default:
                    v11 := []int{128, -491, -799, -2147483648}
                    _ = v11
                }
            }
        }
    case 42:
        v12 := GenS1{F0: v3}
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v5 = v5
        var v14 int64 = 693
        switch "a" {
        case "hello":
            v1 -= 42
        case "":
            var v15 int16 = -1
            _ = v15
        case "a":
            if (0 <= (^(v14))) {
                v4 = 99
                v16 := GenS1{F0: (v3 | -1)}
                _ = v16
            }
        default:
            switch "foo bar" {
            case "foo bar":
                v17 := GenS2{F0: (v5 >> (((v1 / ((v1) | 1))) & 63)), F1: 534, F2: (v4 * 1)}
                _ = v17
            default:
                v18 := []int{}
                _ = v18
            }
        }
        _ = fi13
        _ = v14
    }
    var v19 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
