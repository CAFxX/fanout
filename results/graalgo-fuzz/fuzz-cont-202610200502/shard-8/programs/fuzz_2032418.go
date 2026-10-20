package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
func genFn0() uint8 {
    var v4 int16 = 0
    var v5 int16 = v4
    v6 := GenS1{F0: ("foo bar" + "foo bar"), F1: 9223372036854775807, F2: -9223372036854775808}
    _ = v4
    _ = v5
    _ = v6
    return 2
}

func main() {
    v1 := []uint32{2, 586, 2147483648, 2147483648}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 227
    if false {
        v3 = (v3 / (((v3 * (v3 ^ (v3 & (v3 + (v3 + v3)))))) | 1))
    } else {
        v7 := GenS2{F0: 42}
        v3 += v3
        _ = v7
    }
    v8 := GenS2{F0: -1}
    delete(v2, 2)
    v1 = append(v1, 708)
    v3 = v3
    if ((0.1 >= 0.1) || (!(18446744073709551615 == 0))) {
        v3 = v3
        _ = len(v1)
        switch 341 {
        case 0:
            var v9 float32 = 1e10
            _ = v9
        case 1:
            if false {
                var v10 uint8 = v3
                fmt.Println(v10)
                _ = v10
            }
        default:
            switch "a" {
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "\n\t":
                v3 -= 1
            case "a":
                v11 := GenS3{F0: 0}
                _ = v11
            }
        }
    }
    v1 = append(v1, 65535)
    v3 += (^(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
