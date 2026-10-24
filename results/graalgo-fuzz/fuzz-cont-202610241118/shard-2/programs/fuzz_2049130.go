package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int8; F1 int32; F2 uint16 }
func genFn0(p0 int, p1 uint8) uint16 {
    fmt.Println(p1)
    v4 := GenS1{F0: 42}
    p0 -= 9223372036854775807
    var v5 int8 = 100
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 651
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 2
    v3 := []int{4294967295, 296, 9223372036854775807, -915}
    v2 += 110
    v6 := GenS2{F0: 86, F1: 173}
    var v7 uint8 = 42
    if ((((!false) && (696 >= 2147483648)) && (3.14159 >= (-2.25 - float64(1.5)))) && true) {
        var v8 uint = 180
        _ = v8
    }
    v2 += (v2 >> ((370) & 63))
    switch 32768 {
    case -2147483648:
        v7 -= v7
    case 4294967295:
        if len(v3) > 0 {
            v3[0] = 127
        }
    case 42:
        v7 -= (v7 - v7)
    }
    v9 := GenS3{F0: 42, F1: 474, F2: (v2 ^ (v2 << ((uint(v7)) & 63)))}
    if (genFn0(-9223372036854775808, (v7 >> ((255) & 63))) > v2) {
        v2 = v2
        v7 = v7
        var v10 int32 = 0
        _ = v10
    }
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
