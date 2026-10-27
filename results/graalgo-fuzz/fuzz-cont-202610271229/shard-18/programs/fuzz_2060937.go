package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint }
func genFn0(p0 uint16, p1 uint64, p2 uint) int {
    var v5 float64 = float64(0.1)
    v6 := GenS1{F0: (0.1 >= v5), F1: 0, F2: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    v1 := []int{-997, 2147483647}
    var v2 int = -2147483648
    var v3 float32 = 1e10
    v4 := []uint32{}
    var v7 int32 = 641
    v2 = (v2 * v2)
    fmt.Println(v2, v3, v7)
    if (((v7 % ((v7) | 1)) == (v7 / (((v7 << ((210) & 63))) | 1))) && (((true || false) && (!true)) && false)) {
        fmt.Println(v2, v3, v7)
    }
    if false {
        v7 = (v7 << ((0) & 63))
    }
    v8 := GenS2{F0: (0.0 != 3.14159)}
    if true {
        v1 = append(v1, -2147483648)
    }
    v1 = append(v1, v2)
    v9 := GenS3{F0: "a", F1: 4294967295}
    v10 := GenS1{F0: ((v7 ^ (v7 >> ((897) & 63))) != (v7 - v7)), F1: -72, F2: 1}
    var v11 int8 = 0
    v11 = (v11 - (v11 + v11))
    v11 = v11
    switch genFn0(244, 42, 65535) {
    case -746:
        if len(v4) > 0 {
            v4[0] = 235
        }
    case 460:
        fmt.Println(v7, v2, v7)
    case 833:
        switch (("a" + ("foo bar" + "")) + "a") {
        case "hello":
            v3 = v3
        case "foo bar":
            v12 := GenS1{F0: (!(!((true || (9223372036854775807 >= 662)) && false))), F1: v11, F2: 157}
            _ = v12
        case "":
            fmt.Println(v2)
        }
    default:
        v7 += -1
    }
    var v13 uint32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
