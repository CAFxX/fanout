package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int }
func genFn0() string {
    var v4 float32 = 0.1
    fmt.Println(v4)
    v4 -= float32((-2.25 / (float64(v4) * -2.25)))
    fmt.Println(v4)
    _ = v4
    return "a"
}
func genFn1(p0 uint32, p1 int8, p2 int8) uint8 {
    p1 = int8(p0)
    v5 := GenS1{F0: 0, F1: -223}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 255
}

func main() {
    var v1 bool = true
    var v2 uint = 18446744073709551615
    var v3 int8 = -128
    v6 := GenS2{F0: -9223372036854775808, F1: 42}
    v1 = (((32767 < 1) || v1) && v1)
    v7 := GenS1{F0: (^(v2)), F1: 127}
    var v8 int = 42
    v8 -= (^(v8))
    var v9 int16 = 32767
    if false {
        v10 := []int{v8}
        var v11 int8 = 42
        _ = v10
        _ = v11
    }
    if (v8 >= v8) {
        v2 = (v2 + 2)
    }
    v12 := GenS3{F0: ("hello" + ""), F1: 255}
    fmt.Println(v9)
    switch 327 {
    case 9223372036854775807:
        v13 := GenS4{F0: (v8 % ((v8) | 1))}
        _ = v13
    case 475:
        switch "" {
        case "foo bar":
            v14 := GenS2{F0: -1, F1: 2}
            _ = v14
        }
    case 32768:
        v15 := []int{v8, (v8 << (((v2 ^ (v2 ^ 42))) & 63)), (v8 / ((v8) | 1)), v8}
        _ = v15
    default:
        var v16 int8 = v3
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
