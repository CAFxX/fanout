package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 bool }

func main() {
    var v1 uint32 = 1
    var v2 uint8 = 47
    var v3 uint = 477
    v4 := []uint32{(v1 ^ 0), (v1 - (v1 << ((v3) & 63))), (v1 - (v1 ^ v1))}
    v5 := make(map[int]uint32)
    v3 = (v3 ^ v3)
    var v6 float64 = (-2.25 - -2.25)
    var v7 int = 9223372036854775807
    switch "foo bar" {
    case "hello":
        if (false || (0 <= 0)) {
            _ = len(v5)
            var v8 bool = (((v1 << (((v3 / (((v3 - v3)) | 1))) & 63)) < (v1 | (v1 | v1))) || (((v1 > (^(v1))) || (!false)) && (872 <= 422)))
            var v9 bool = (0 == 1)
            _ = v8
            _ = v9
        }
    case "a":
        if false {
            var v10 int64 = int64(v1)
            _ = v10
        }
    default:
        v11 := GenS1{F0: 650, F1: 2, F2: (((1e10 * float32(((0.0 * v6) - v6))) * -2.25) <= 3.14159)}
        _ = v11
    }
    v12 := GenS1{F0: 1000, F1: 18446744073709551615, F2: ((315 >= (^(v3))) && (-32768 >= 983))}
    v4 = append(v4, 65535)
    v3 -= 802
    var v13 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
