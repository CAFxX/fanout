package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint16 }
func genFn0(p0 uint8) int32 {
    var v5 uint32 = 552
    fmt.Println(p0, p0)
    if ((9223372036854775807 > 42) || (18446744073709551615 == 733)) {
        if (((1e10 * 0.0) + 1e10) != (-2.25 * 1e10)) {
            v6 := []int{42, 4294967295, 42, 320}
            fmt.Println(v5)
            _ = v6
        } else {
            v7 := []int{}
            _ = v7
        }
        p0 -= (p0 - (^(p0)))
    }
    _ = p0
    _ = v5
    return -952
}
func genFn1() bool {
    v8 := []int{419, 32768, -1}
    v9 := GenS1{F0: ((((true || (-9 > -41)) && true) || (493 <= 1000)) && true), F1: 1, F2: 32767}
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 int64 = 0
    var v2 uint = 65535
    var v3 int8 = 42
    v4 := []int{}
    v4 = append(v4, -66)
    var v10 float64 = (3.14159 * -2.25)
    v1 = -9223372036854775808
    fmt.Println(v3, v10, v2)
    v11 := GenS1{F0: false, F1: v2, F2: 65535}
    var v12 int64 = (v1 & (v1 ^ (v1 >> (((v2 >> ((v2) & 63))) & 63))))
    fmt.Println(v10)
    v3 -= 0
    v10 = (((v10 / v10) + v10) + float64(0.0))
    v4 = append(v4, 55)
    v2 -= v2
    if (true || (848 > (v2 * v2))) {
        _ = len(v4)
    }
    v4 = append(v4, -14)
    switch 4294967295 {
    case -1:
        v10 = v10
    case -314:
        v4 = append(v4, 128)
    default:
        v2 += 498
    }
    v12 = v12
    if ((1 != int(v3)) && true) {
        v3 = (v3 >> ((v2) & 63))
    }
    v1 = int64(v3)
    v13 := GenS1{F0: (false && true), F1: (v2 % ((v2) | 1)), F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
