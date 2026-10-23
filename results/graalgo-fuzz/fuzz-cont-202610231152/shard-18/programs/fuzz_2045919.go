package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0() int16 {
    var v5 uint8 = 190
    switch 128 {
    case 1:
        v6 := GenS1{F0: 100}
        _ = v6
    case 9223372036854775807:
        v7 := []int{-2147483648, -223}
        _ = v7
    case -945:
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int32 = 384
            _ = fi8
            _ = v9
        }
    default:
        v10 := []int{-9223372036854775808}
        _ = v10
    }
    _ = v5
    return -32768
}

func main() {
    var v1 int64 = -453
    v2 := make(map[uint32]bool)
    var v3 uint32 = 2147483648
    var v4 int16 = 42
    v1 += v1
    var v11 float64 = -2.25
    var v12 bool = ((((65535 <= (^(v3))) && (v1 >= 42)) && false) && false)
    var v13 uint8 = 2
    switch ("hello" + ("" + "")) {
    case "foo bar":
        var v14 uint = 18446744073709551615
        _ = v14
    default:
        v15 := GenS2{F0: 0, F1: -128}
        _ = v15
    }
    if ((-43 > -1) && v12) {
        switch "a" {
        case "hello":
            v3 = v3
        case "foo bar":
            fmt.Println(v11, v11)
        case "":
            var v16 int = 32768
            _ = v16
        }
    } else {
        v17 := v2[(v3 | (v3 ^ v3))]
        v11 += (float64(-2.25) - v11)
        if (v17 && v12) {
            var v18 uint8 = 128
            v17 = (false || (v13 >= (^(v18))))
            _ = v18
        } else {
            v19 := []int{}
            v20 := GenS2{F0: 0, F1: 42}
            _ = v19
            _ = v20
        }
        _ = v17
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        v22 := []int{128}
        for fi23 := 0; fi23 < 18; fi23++ {
            if v12 {
                v24 := v2[1]
                v13 -= v13
                _ = v24
            } else {
                var v25 uint32 = (v3 ^ v3)
                fmt.Println(v25, v13)
                _ = v25
            }
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    _ = len(v2)
    v4 -= genFn0()
    v2[(v3 + v3)] = v12
    v26 := GenS2{F0: -1, F1: 0}
    v12 = (v12 || ("foo bar" == ""))
    var v27 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
