package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint64) uint {
    p0 = p0
    if (0 != 42) {
        fmt.Println(p0)
        p0 += p0
        v5 := []int{}
        _ = v5
    } else {
        if true {
            v6 := GenS1{F0: 128, F1: 355, F2: -1}
            for fi7 := 0; fi7 < 11; fi7++ {
                p0 = p0
                v8 := GenS2{F0: 191}
                _ = fi7
                _ = v8
            }
            _ = v6
        }
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int = 0
        _ = fi9
        _ = v10
    }
    _ = p0
    return 764
}
func genFn1() uint64 {
    var v11 uint8 = 153
    v12 := []int{127, 2147483647, 32767, 0}
    _ = len(v12)
    _ = v11
    _ = v12
    return 18446744073709551615
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]int)
    var v3 int = 2147483647
    var v4 bool = false
    var v13 int16 = (-(v1))
    if true {
        v14 := GenS3{F0: 765, F1: 127}
        v15 := GenS4{F0: 810, F1: v3}
        v3 -= (^(v3))
        _ = v14
        _ = v15
    }
    v4 = false
    var v16 int = (v3 / (((v3 << ((0) & 63))) | 1))
    v4 = ((((1e10 / 1.0) / 1.5) - (0.0 * (1.5 - -2.25))) >= ((1e10 * (1.5 * 1.5)) / ((1.5 + (0.0 / -2.25)) + float64(((0.0 * float32(float64(3.14159))) / (-2.25 / 0.1))))))
    for fi17 := 0; fi17 < 15; fi17++ {
        v1 = (v1 / ((1000) | 1))
        _ = fi17
    }
    var v18 int8 = 42
    fmt.Println(v18, v16)
    var v19 uint8 = 160
    v3 += (v16 + (v3 * v16))
    var v20 int32 = 0
    switch "" {
    case "a":
        if ((v13 == (^(v1))) || false) {
            if ((v13 + 1000) != v13) {
                fmt.Println(v3, v19, v1)
                v21 := GenS5{F0: ("a" + "")}
                _ = v21
            }
            if (-1 >= -9223372036854775808) {
                v13 = v13
                v22 := v2[v3]
                var v23 float32 = float32(3.14159)
                _ = v22
                _ = v23
            }
            if true {
                switch "a" {
                case "a":
                    var v24 int = v16
                    _ = v24
                case "hello":
                    v18 = (v18 << ((445) & 63))
                }
                v18 -= v18
                v16 = v16
            }
        }
    }
    v25 := GenS6{F0: genFn1()}
    switch "" {
    case "":
        for fi26 := 0; fi26 < 5; fi26++ {
            _ = len(v2)
            _ = fi26
        }
    case "a":
        v13 = 718
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
}
