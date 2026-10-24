package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int }
type GenS3 struct { F0 bool; F1 uint16; F2 bool }

func main() {
    var v1 int8 = 83
    var v2 int32 = 2147483647
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 19; fi5++ {
        if (0 >= 255) {
            v2 -= -1
        } else {
            _ = len(v4)
            v6 := []int{(fi5 % ((fi5) | 1)), fi5}
            v2 = int32(fi5)
            _ = v6
        }
        v2 += v2
        v7 := GenS1{F0: 753, F1: -9223372036854775808, F2: 713}
        _ = fi5
        _ = v7
    }
    v8 := []int{-9223372036854775808, 32767, -181, 32768}
    var v9 int8 = v1
    v2 += v2
    v10 := GenS2{F0: 18446744073709551615, F1: 1, F2: 32768}
    for fi11 := 0; fi11 < 6; fi11++ {
        if (true && false) {
            v2 = 42
            var v12 int16 = 169
            _ = v12
        } else {
            v13 := GenS3{F0: (!((!(9223372036854775807 <= -452)) && true)), F1: 1, F2: false}
            v14 := GenS2{F0: 42, F1: 0, F2: (-(fi11))}
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    fmt.Println(v1, v9)
    var v15 uint16 = 255
    v15 = v15
    var v16 uint32 = 2147483647
    v17 := GenS3{F0: false, F1: (v15 * 387), F2: false}
    var v18 float32 = 1.5
    _ = len(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
