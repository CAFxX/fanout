package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int; F1 int8 }
func genFn0() bool {
    var v5 bool = true
    var v6 uint8 = 240
    _ = v5
    _ = v6
    return (!v5)
}

func main() {
    var v1 int8 = -17
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    var v4 int8 = (v1 % ((127) | 1))
    v4 += v4
    v1 -= (v4 - (v4 * 0))
    var v7 int32 = 2147483647
    v7 += 42
    if ((1.5 + -2.25) != -2.25) {
        v4 -= (v4 * v1)
        v4 -= v4
    }
    delete(v2, 65535)
    var v8 uint = 4294967295
    switch ("foo bar" + ("hello" + "hello")) {
    case "":
        v9 := []int{127, 712, -213}
        _ = v9
    case "a":
        var v10 float64 = -2.25
        _ = v10
    case "\n\t":
        v11 := v2[2147483647]
        _ = v11
    }
    v4 = v4
    v8 -= 1
    v12 := GenS1{F0: 758, F1: 65535}
    fmt.Println(v4)
    if false {
        v1 += v4
        var v13 int8 = (-(v4))
        _ = v13
    } else {
        v14 := GenS1{F0: 330, F1: 42}
        v1 -= (v1 & (v4 & (v4 * v1)))
        v15 := GenS2{F0: 2147483647, F1: v1}
        _ = v14
        _ = v15
    }
    var v16 int32 = v7
    v17 := GenS2{F0: -1, F1: v1}
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 int64 = 9223372036854775807
        _ = fi18
        _ = v19
    }
    v20 := v2[65535]
    fmt.Println(v7, v8, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
