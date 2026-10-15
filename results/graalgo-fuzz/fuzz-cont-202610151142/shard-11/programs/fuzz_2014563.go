package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 bool }

func main() {
    var v1 bool = (!true)
    v2 := make(map[int]string)
    var v3 int16 = 1
    var v4 float64 = float64(3.14159)
    switch 9223372036854775807 {
    case 127:
        var v5 float64 = v4
        _ = v5
    }
    var v6 int16 = (v3 + 1)
    v1 = ("" != "foo bar")
    v1 = (2 <= 0)
    v7 := GenS1{F0: 334, F1: 0}
    v8 := GenS1{F0: 902, F1: 42}
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := []int{(fi9 << ((58) & 63)), 128, -1, (-(fi9))}
        var v11 int32 = -158
        _ = fi9
        _ = v10
        _ = v11
    }
    v1 = (4294967295 <= 2147483648)
    var v12 int8 = -84
    v13 := GenS1{F0: 576, F1: 2}
    v12 -= 0
    var v14 bool = false
    var v15 uint = 303
    if (0.1 <= 0.0) {
        v16 := GenS2{F0: 2, F1: 67, F2: (!v14)}
        _ = v16
    } else {
        v6 = v3
        fmt.Println(v4, v4, v15)
    }
    v3 += 802
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
