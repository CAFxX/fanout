package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int8; F1 int32 }
func genFn0() bool {
    var v5 uint = 18446744073709551615
    if (true && ((true || true) && false)) {
        switch "foo bar" {
        case "\n\t":
            switch -9223372036854775808 {
            case 1:
                v5 -= v5
            case 137:
                v6 := GenS1{F0: 4294967295, F1: -872}
                _ = v6
            default:
                v5 -= 486
            }
        default:
            v7 := GenS1{F0: 4294967295, F1: -9223372036854775808}
            _ = v7
        }
        switch "\n\t" {
        case "hello":
            v5 = 4294967295
        case "a":
            var v8 int32 = 1
            _ = v8
        case "":
            v5 -= 912
        }
    } else {
        fmt.Println(v5)
        v9 := []int{577, -1, 9223372036854775807, 0}
        _ = v9
    }
    var v10 float32 = 0.0
    _ = v10
    _ = v5
    return false
}

func main() {
    var v1 int16 = 32767
    var v2 uint = 375
    var v3 uint = 321
    var v4 int16 = v1
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := []int{}
        switch fi11 {
        case -956:
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        default:
            switch (-(fi11)) {
            case 128:
                _ = len(v12)
            case 0:
                v13 := GenS1{F0: -90, F1: fi11}
                _ = v13
            default:
                v14 := GenS2{F0: -1, F1: 42}
                _ = v14
            }
        }
        _ = fi11
        _ = v12
    }
    switch 42 {
    case 4294967295:
        if (1.5 != (0.0 * 1e10)) {
            fmt.Println(v1, v4, v1)
            fmt.Println(v1, v1)
        }
    case -507:
        v15 := GenS2{F0: -97, F1: 550}
        _ = v15
    }
    v1 = v1
    v2 = v3
    var v16 bool = true
    v17 := []int{-817, 128}
    if false {
        var v18 uint16 = 32767
        if len(v17) > 0 {
            v17[0] = 2147483647
        }
        _ = v18
    }
    switch "\n\t" {
    case "\n\t":
        v3 -= (^(v2))
    case "foo bar":
        fmt.Println(v2)
    }
    v17 = append(v17, -763)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
