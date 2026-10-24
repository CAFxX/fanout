package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint64; F1 bool }
func genFn0() int32 {
    var v5 int64 = -709
    fmt.Println(v5)
    _ = v5
    return 2147483647
}
func genFn1() string {
    switch -407 {
    case 0:
        if true {
            if false {
                v6 := GenS1{F0: 65535, F1: 18446744073709551615}
                _ = v6
            } else {
                var v7 uint8 = 42
                _ = v7
            }
        } else {
            var v8 uint16 = 255
            v9 := GenS1{F0: 2, F1: 42}
            v10 := []int{-1}
            _ = v10
            _ = v8
            _ = v9
        }
    case 32768:
        switch "hello" {
        case "hello":
            v11 := []int{}
            _ = v11
        case "":
            var v12 int32 = 2147483647
            _ = v12
        case "foo bar":
            var v13 int8 = 40
            _ = v13
        default:
            var v14 uint16 = 0
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint8 = 68
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 18446744073709551615, F1: ((4294967295 > 1) || (-462 >= 805))}
    _ = v17
    return ""
}

func main() {
    v1 := []uint32{447}
    v2 := make(map[string]bool)
    var v3 uint8 = 128
    var v4 float64 = 1.5
    var v18 float32 = 0.1
    v4 = (v4 * (float64((((v18 * (float32((v4 - 1.5)) - (v18 * v18))) + v18) - 1e10)) + 1e10))
    v19 := GenS2{F0: 18446744073709551615, F1: (((-22 <= 1) || (-84 >= 42)) || true)}
    var v20 uint32 = 1
    v18 = 0.1
    v3 = 1
    _ = len(v2)
    v21 := GenS2{F0: 0, F1: (false && false)}
    var v22 float32 = ((v18 + ((v18 * (v18 * ((1.5 - (1.5 / v18)) - v18))) * 1e10)) + (float32(1.5) - 0.0))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
