package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint }
func genFn0() uint {
    v5 := []int{-56, -2147483648, 980}
    var v6 float32 = (0.1 - 3.14159)
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int32) int32 {
    var v7 uint = 2
    fmt.Println(p0, v7)
    var v8 int32 = (-(p0))
    switch ((("" + "foo bar") + "hello") + "a") {
    case "hello":
        p0 += p0
    }
    _ = p0
    _ = v7
    _ = v8
    return (^(p0))
}

func main() {
    var v1 uint64 = 351
    var v2 int32 = 42
    var v3 uint8 = 128
    var v4 float32 = float32(float64((0.0 + ((1.5 / float32(float64(3.14159))) / 1.0))))
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(v4, v3)
        _ = fi9
    }
    v10 := GenS1{F0: -1, F1: v2}
    var v11 int64 = -63
    v12 := GenS1{F0: v2, F1: (v2 >> ((2) & 63))}
    var v13 uint16 = 2
    for fi14 := 0; fi14 < 6; fi14++ {
        if ((^(v11)) != v11) {
            var v15 float64 = 0.0
            v13 = uint16(v1)
            _ = v15
        } else {
            v3 -= v3
        }
        _ = fi14
    }
    v16 := GenS2{F0: 89, F1: 65535, F2: 1}
    v17 := []int{42, 128, 1}
    var v18 int64 = (v11 * (v11 >> ((uint(v3)) & 63)))
    var v19 int32 = int32(v18)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
