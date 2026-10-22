package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int16 }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    var v1 float64 = (1e10 * float64(1.5))
    v2 := make(map[int]uint32)
    var v3 uint8 = 194
    v4 := GenS1{F0: 779, F1: (!(v1 < float64(3.14159))), F2: 1}
    if true {
        var v5 bool = (((0 >= -2147483648) && true) && false)
        v6 := GenS2{F0: 0, F1: -1}
        _ = v5
        _ = v6
    }
    var v7 uint = 1
    v3 += 42
    fmt.Println(v1, v3, v3)
    var v8 int = 637
    v9 := GenS2{F0: 100, F1: int8(v8)}
    var v10 int64 = -9223372036854775808
    var v11 uint16 = 32767
    if ((v7 >= 255) && true) {
        if (!(798 != 4294967295)) {
            v12 := GenS1{F0: (v10 / ((v10) | 1)), F1: false, F2: 42}
            _ = v12
        }
        switch "\n\t" {
        case "hello":
            var v13 int32 = -438
            _ = v13
        case "":
            switch v8 {
            case 2147483647:
                v3 += (v3 << ((v7) & 63))
            }
        default:
            v2[v8] = 42
        }
        v7 += v7
    } else {
        var v14 bool = (!false)
        var v15 uint8 = uint8(v11)
        v10 += v10
        _ = v14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
