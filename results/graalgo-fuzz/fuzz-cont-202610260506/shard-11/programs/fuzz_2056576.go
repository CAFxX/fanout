package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint8, p1 int32) string {
    if true {
        p1 = p1
    } else {
        fmt.Println(p1, p1)
        switch "hello" {
        case "":
            v7 := []int{-1}
            _ = v7
        }
    }
    p1 -= (p1 << ((1) & 63))
    v8 := []int{-49, 32768, 128, -346}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v8
    return ("" + ("foo bar" + "hello"))
}
func genFn1() bool {
    var v9 int8 = -112
    var v10 float64 = 0.1
    _ = v10
    _ = v9
    return false
}

func main() {
    var v1 float64 = (((float64(1.5) - 0.0) - 0.1) + (float64(-2.25) * float64((1e10 * 1e10))))
    var v2 int = 32768
    v3 := make(map[string]uint32)
    var v4 int = (v2 + 2147483647)
    var v5 int = (-(v4))
    var v6 uint = 115
    v11 := GenS1{F0: -76}
    v4 -= 2147483647
    var v12 uint16 = 42
    var v13 int8 = 15
    for fi14 := 0; fi14 < 11; fi14++ {
        if true {
            v4 = 2147483647
        } else {
            v15 := GenS2{F0: 81}
            v6 -= (v6 >> (((v6 * (^(v6)))) & 63))
            v16 := GenS3{F0: 654, F1: v13}
            _ = v15
            _ = v16
        }
        v17 := GenS3{F0: 0, F1: v13}
        _ = fi14
        _ = v17
    }
    fmt.Println(v6)
    switch "" {
    case "foo bar":
        var v18 float64 = ((-2.25 / (1.5 + v1)) - 3.14159)
        _ = v18
    }
    var v19 int16 = 1000
    v19 -= (v19 ^ v19)
    var v20 uint8 = 255
    var v21 int = v5
    var v22 int = v5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
