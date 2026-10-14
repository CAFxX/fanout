package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 string; F1 int32; F2 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint8) int16 {
    switch ("foo bar" + "") {
    case "a":
        v4 := []int{-2147483648, 208}
        _ = v4
    case "\n\t":
        v5 := GenS1{F0: 835, F1: 1, F2: 275}
        _ = v5
    case "hello":
        v6 := GenS1{F0: 268, F1: 527, F2: 2147483648}
        _ = v6
    }
    var v7 float32 = float32((0.1 + -2.25))
    var v8 float32 = (v7 / 3.14159)
    if false {
        v9 := []int{-1, 1, -973, 9223372036854775807}
        v8 += float32((3.14159 - (float64(v8) / ((0.0 - 1.5) + (1e10 * float64((v8 + (float32(1e10) / v8))))))))
        v10 := GenS1{F0: 1, F1: 2, F2: 2147483648}
        _ = v10
        _ = v9
    }
    var v11 string = ("" + (("\n\t" + "a") + "\n\t"))
    _ = p0
    _ = v11
    _ = v7
    _ = v8
    return 1000
}
func genFn1(p0 uint32) string {
    switch 849 {
    case 32768:
        v12 := []int{-9223372036854775808, 128, -783, -898}
        _ = v12
    case -1:
        v13 := []int{645, 2147483647, 42, -531}
        _ = v13
    }
    v14 := []int{9223372036854775807, 127, 32767, 464}
    switch 2147483647 {
    case 127:
        fmt.Println(p0)
    case 777:
        v15 := GenS2{F0: ("" >= "a"), F1: 255}
        _ = v15
    case 128:
        var v16 uint64 = 18446744073709551615
        _ = v16
    default:
        v17 := GenS3{F0: "\n\t", F1: 312, F2: 128}
        _ = v17
    }
    var v18 float32 = 1.5
    var v19 uint8 = 42
    _ = p0
    _ = v14
    _ = v18
    _ = v19
    return "\n\t"
}

func main() {
    var v1 float64 = ((1.5 + 1e10) * 0.1)
    var v2 uint16 = 2
    var v3 int16 = -803
    v3 = (v3 - (v3 - genFn0(1)))
    v1 = v1
    switch 1 {
    case 42:
        v3 = v3
    case 0:
        if true {
            if (true && (0 < 0)) {
                v2 = v2
                var v20 float64 = ((v1 * 1e10) * (1e10 * v1))
                _ = v20
            }
            v3 = genFn0(127)
            v21 := []int{}
            _ = v21
        } else {
            v22 := []int{2147483647, 9223372036854775807, 806}
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            _ = v22
        }
    }
    v2 += (v2 / (((v2 << ((0) & 63))) | 1))
    v23 := GenS2{F0: true, F1: (v2 % ((v2) | 1))}
    v2 -= 42
    fmt.Println(v2, v2)
    v24 := GenS4{F0: 40}
    var v25 float64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
