package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 uint16 }

func main() {
    var v1 float64 = (1e10 - (-2.25 * 0.0))
    var v2 float32 = (0.0 / 1.0)
    var v3 string = "a"
    var v4 int8 = -39
    v4 -= v4
    var v5 uint = 4294967295
    v4 = 127
    v2 -= v2
    var v6 float64 = float64(float32(float64(-2.25)))
    v4 -= 100
    var v7 float64 = v6
    if (true && (false && (!(!false)))) {
        v8 := GenS1{F0: 142, F1: -9223372036854775808, F2: 18446744073709551615}
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    v7 -= -2.25
    if ((("a" + "a") + "foo bar") > v3) {
        v11 := []int{827, 0, 128, -1}
        _ = v11
    }
    switch 2147483647 {
    case 32767:
        if (1 > 1) {
            var v12 uint32 = 2147483647
            v13 := []int{42}
            v7 += 1.5
            _ = v12
            _ = v13
        }
    case 964:
        v14 := []int{32767, 508}
        _ = v14
    }
    if ((v4 << ((v5) & 63)) > (v4 | (v4 - (v4 % ((127) | 1))))) {
        var v15 uint64 = 969
        v16 := GenS2{F0: v4, F1: 0, F2: 782}
        _ = v15
        _ = v16
    } else {
        fmt.Println(v4)
        var v17 float64 = (float64(v2) * float64(float32(v6)))
        v18 := []int{-650, -22, -55, 32768}
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
