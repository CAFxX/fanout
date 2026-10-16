package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint16 }

func main() {
    var v1 int = 4294967295
    v2 := make(map[uint32]string)
    v3 := []int{(v1 ^ 42), (^(v1))}
    v1 = v1
    var v4 int64 = 1
    if true {
        fmt.Println(v1, v1)
        if true {
            if len(v3) > 0 {
                v3[0] = -866
            }
        }
        fmt.Println(v4)
    }
    switch (^(v1)) {
    case 50:
        v5 := GenS1{F0: ("hello" >= "a"), F1: 246}
        _ = v5
    case 2147483647:
        v4 -= (v4 - (v4 / ((v4) | 1)))
    case -408:
        fmt.Println(v1, v1)
    }
    v4 = (v4 - v4)
    v1 = (v1 << ((320) & 63))
    v1 = (v1 / ((v1) | 1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v6 int16 = 1
    delete(v2, 2)
    v1 = (-(v1))
    var v7 float64 = ((1e10 + float64(-2.25)) - 1.5)
    var v8 int8 = 1
    v9 := GenS2{F0: -2147483648}
    v10 := GenS3{F0: 255, F1: 2}
    v4 -= (v4 & v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
