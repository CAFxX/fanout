package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint64 }
type GenS2 struct { F0 bool; F1 int }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint8 = 68
    var v3 uint = 42
    var v4 int64 = (v1 - -23)
    v5 := []int{-9223372036854775808}
    var v6 int64 = (v1 | 565)
    switch -2147483648 {
    case -959:
        v4 = (v4 / (((v6 + (v1 | v1))) | 1))
    case -145:
        v5 = append(v5, -337)
    case 128:
        v3 -= (v3 << ((255) & 63))
    default:
        v4 -= (v4 ^ v4)
    }
    v7 := GenS1{F0: 0, F1: 42, F2: 1}
    v8 := GenS1{F0: 317, F1: 110, F2: 0}
    var v9 float32 = ((0.0 * (((((((float32(float64(-2.25)) * 1.5) + ((3.14159 * -2.25) + 1e10)) - 3.14159) * (0.1 / float32(1e10))) - 0.1) + ((0.0 * 1e10) / (float32(1.5) * 1.5))) + 1.5)) / (-2.25 + (-2.25 - (3.14159 / 3.14159))))
    _ = len(v5)
    v4 = (^(v4))
    v4 = v4
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v10 := GenS2{F0: ((v2 & (v2 + v2)) != (v2 >> ((65535) & 63))), F1: 32768}
    if ((false || ((!true) || (!(!false)))) || false) {
        _ = len(v5)
        _ = len(v5)
    }
    v9 = v9
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
