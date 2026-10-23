package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 int32; F1 int64; F2 bool }
type GenS7 struct { F0 uint32; F1 uint32; F2 uint8 }
func genFn0(p0 uint32, p1 uint8) int32 {
    if (18446744073709551615 == 18446744073709551615) {
        switch "" {
        case "foo bar":
            p1 = p1
        case "a":
            if (735 > 231) {
                p1 -= (p1 % ((2) | 1))
                v5 := GenS1{F0: (!true), F1: p0, F2: 42}
                v6 := GenS2{F0: 18446744073709551615}
                _ = v5
                _ = v6
            }
        }
    } else {
        v7 := GenS2{F0: uint(p1)}
        _ = v7
    }
    var v8 string = "a"
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]string)
    var v3 int = 32768
    var v4 uint8 = 0
    fmt.Println(v1)
    v9 := GenS3{F0: 32767, F1: 2}
    v10 := GenS4{F0: -258, F1: 464}
    v11 := []int{v3, v3}
    fmt.Println(v4, v3, v3)
    var v12 string = v1
    v1 = (v1 + "\n\t")
    v13 := GenS5{F0: 206, F1: v3}
    for fi14 := 0; fi14 < 18; fi14++ {
        v4 += v4
        v15 := GenS6{F0: -2147483648, F1: -9223372036854775808, F2: (65535 < 1)}
        _ = fi14
        _ = v15
    }
    v16 := GenS6{F0: 42, F1: 0, F2: (-2.25 > 3.14159)}
    v17 := GenS7{F0: 4294967295, F1: 2, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
