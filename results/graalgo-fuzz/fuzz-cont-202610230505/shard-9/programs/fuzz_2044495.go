package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int64 }
func genFn0() int8 {
    switch "hello" {
    case "hello":
        var v6 int8 = 0
        _ = v6
    }
    var v7 bool = ((2147483647 >= 1) && true)
    if (1 > 1) {
        v8 := []int{2147483647, 1, -9223372036854775808, -2147483648}
        var v9 uint8 = 125
        switch (("\n\t" + "") + ("" + "\n\t")) {
        case "\n\t":
            _ = len(v8)
        case "hello":
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 int64 = -1
                _ = fi10
                _ = v11
            }
        default:
            fmt.Println(v9, v7)
        }
        _ = v8
        _ = v9
    } else {
        var v12 uint32 = 0
        if v7 {
            v13 := []int{9223372036854775807, 1, -580}
            fmt.Println(v12, v12)
            _ = v13
        } else {
            if v7 {
                switch "hello" {
                case "hello":
                    v12 = (v12 % ((v12) | 1))
                case "":
                    var v14 uint8 = 0
                    _ = v14
                }
                v7 = (383 == 1)
                var v15 uint32 = v12
                _ = v15
            } else {
                v12 = v12
            }
            v16 := []int{708, -831, -2147483648, -9223372036854775808}
            _ = v16
        }
        if (127 != 128) {
            var v17 int32 = 42
            v18 := []int{0, 1, 4294967295}
            switch "foo bar" {
            case "foo bar":
                v17 = v17
            case "hello":
                if len(v18) > 0 {
                    v18[0] = 42
                }
            default:
                v7 = (1 != 42)
            }
            _ = v17
            _ = v18
        }
        _ = v12
    }
    var v19 uint16 = 954
    _ = v19
    _ = v7
    return -1
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    v4 := make(map[int]int)
    var v5 float32 = 3.14159
    v20 := GenS1{F0: 0, F1: 0, F2: -21}
    for fi21 := 0; fi21 < 16; fi21++ {
        delete(v3, "a")
        _ = fi21
    }
    var v22 float32 = float32((1e10 * 3.14159))
    v22 = ((v22 - (float32((1.5 + 0.0)) + (((-2.25 / v22) + (1.5 + (0.1 - (v22 / (float32(0.1) * v22))))) + v5))) - -2.25)
    var v23 uint8 = 2
    var v24 uint = (^(v1))
    switch 827 {
    case 0:
        v25 := []int{2147483647}
        _ = v25
    case 22:
        v5 += (float32(-2.25) / ((v22 / v5) - 1.5))
    case 305:
        v26 := []int{128, 531, 86, -2147483648}
        _ = v26
    default:
        fmt.Println(v23, v5)
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        v23 = 158
        var v28 uint = v24
        _ = fi27
        _ = v28
    }
    fmt.Println(v5, v22)
    v24 = (v1 % (((^(v1))) | 1))
    if false {
        v1 += 818
    } else {
        switch "" {
        case "a":
            fmt.Println(v24, v23)
        case "":
            for fi29 := 0; fi29 < 19; fi29++ {
                if (("" + "") != ("a" + "a")) {
                    fmt.Println(v23, v22)
                }
                var v30 float64 = (((0.0 - (1.5 / 1.0)) - 0.0) / 1.5)
                fmt.Println(v5)
                _ = fi29
                _ = v30
            }
        }
        var v31 int64 = -659
        var v32 int8 = 93
        _ = v31
        _ = v32
    }
    v23 = (^(v23))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
