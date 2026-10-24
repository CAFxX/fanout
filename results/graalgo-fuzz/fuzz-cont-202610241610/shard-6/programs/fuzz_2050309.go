package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint8; F1 int64 }
func genFn0(p0 bool) int64 {
    p0 = true
    var v5 int16 = 1
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 uint16 = 312
    var v2 uint64 = 624
    var v3 int64 = 724
    var v4 string = "\n\t"
    fmt.Println(v4, v3, v4)
    v8 := GenS1{F0: 2, F1: true}
    v1 -= 246
    v4 = ""
    if (1 >= 1) {
        v1 = (^(v1))
    }
    v9 := GenS1{F0: 191, F1: (42 <= 0)}
    v2 = (v2 << ((255) & 63))
    if (v2 >= 42) {
        v10 := []int{32768}
        v11 := GenS2{F0: uint8(v2), F1: (v3 ^ 9223372036854775807)}
        _ = v10
        _ = v11
    }
    if false {
        v1 = 135
        fmt.Println(v1, v1)
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v4 = ("" + (v4 + v4))
        var v13 float32 = (0.1 / 0.1)
        var v14 uint64 = (^(v2))
        _ = fi12
        _ = v13
        _ = v14
    }
    if true {
        fmt.Println(v3, v1, v2)
        v15 := GenS2{F0: 1, F1: (v3 % (((v3 % (((v3 * (v3 << ((65535) & 63)))) | 1))) | 1))}
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
