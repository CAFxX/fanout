package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int64 }
func genFn0(p0 int32) uint16 {
    p0 = 2147483647
    var v5 int64 = 1
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    v6 := []int{-9223372036854775808, 105}
    fmt.Println(v1)
    if true {
        v1 = 42
        v1 -= v1
    } else {
        v1 = (^(v1))
        switch ("\n\t" + "\n\t") {
        case "a":
            var v7 int8 = 127
            _ = v7
        case "foo bar":
            var v8 uint32 = 0
            _ = v8
        default:
            var v9 float32 = ((3.14159 * float32(3.14159)) - float32(((float64(((-2.25 + -2.25) * 1e10)) - (0.1 * 1e10)) - (3.14159 * float64(float32(-2.25))))))
            _ = v9
        }
    }
    var v10 int16 = -951
    v10 = v10
    switch (("foo bar" + "hello") + "foo bar") {
    case "hello":
        v11 := GenS1{F0: 2147483648, F1: 32767, F2: 42}
        _ = v11
    default:
        var v12 uint16 = 255
        _ = v12
    }
    if ((18446744073709551615 >= 974) || (32767 != (v10 ^ v10))) {
        v10 = (v10 / ((42) | 1))
        v13 := v4[-507]
        var v14 int8 = -1
        _ = v13
        _ = v14
    } else {
        var v15 float32 = -2.25
        _ = v15
    }
    switch 128 {
    case 42:
        fmt.Println(v1)
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v16 := GenS1{F0: 2147483647, F1: 0, F2: -155}
    fmt.Println(v10)
    if (0 >= 42) {
        fmt.Println(v10, v1)
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v1 += v1
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
