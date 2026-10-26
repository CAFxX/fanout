package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint; F2 int16 }
type GenS4 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint16 = 127
    var v2 float32 = ((0.1 - 1e10) - 0.0)
    v3 := make(map[uint32]bool)
    var v4 int8 = 127
    var v5 string = "hello"
    v2 = (v2 - (0.1 + float32(0.1)))
    v6 := []int{-490, 9223372036854775807, 1, -2147483648}
    var v7 string = v5
    var v8 uint16 = (v1 ^ 222)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint8 = 0
        fmt.Println(v5, v5)
        v4 = 42
        _ = fi9
        _ = v10
    }
    switch 575 {
    case -1:
        v4 += v4
    case 919:
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := v3[729]
            _ = fi11
            _ = v12
        }
    }
    if (1 == 1) {
        v2 += 0.0
        if ((v8 << ((18446744073709551615) & 63)) == (v1 << ((65535) & 63))) {
            fmt.Println(v7)
            fmt.Println(v7, v5)
        }
    } else {
        v8 = 276
        fmt.Println(v4, v7)
    }
    var v13 uint = 0
    v14 := GenS1{F0: -31, F1: 957, F2: -9223372036854775808}
    v7 = v5
    v15 := GenS2{F0: 128}
    for fi16 := 0; fi16 < 18; fi16++ {
        if false {
            v17 := GenS2{F0: fi16}
            v5 = ((v5 + v5) + v5)
            _ = v17
        } else {
            if false {
                var v18 int8 = int8(v8)
                fmt.Println(v1)
                v19 := GenS3{F0: v5, F1: 4294967295, F2: -32768}
                _ = v18
                _ = v19
            } else {
                v8 += 255
            }
        }
        if ((false && true) || false) {
            for fi20 := 0; fi20 < 7; fi20++ {
                v21 := GenS4{F0: (v4 < -1), F1: 1}
                var v22 int16 = 1000
                var v23 float64 = 0.1
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        } else {
            fmt.Println(v1, v13)
            if ((87 >= v4) && (!(!false))) {
                var v24 float32 = (v2 - (v2 * v2))
                _ = v24
            }
        }
        var v25 float32 = ((v2 / ((v2 + (0.0 - (1e10 + ((v2 + 3.14159) - float32(0.0))))) - (1.5 + 0.0))) / v2)
        _ = fi16
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
