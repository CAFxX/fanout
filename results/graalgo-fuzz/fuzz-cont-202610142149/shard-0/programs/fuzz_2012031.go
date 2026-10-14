package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint }
type GenS2 struct { F0 bool }
func genFn0() uint32 {
    switch 42 {
    case 1:
        var v4 string = "\n\t"
        _ = v4
    case 42:
        v5 := GenS1{F0: (42 != 0), F1: 1, F2: 566}
        _ = v5
    case -67:
        v6 := []int{951, 127, 765, 968}
        _ = v6
    }
    if false {
        var v7 float64 = 1e10
        var v8 uint32 = 2147483647
        var v9 uint8 = 255
        _ = v7
        _ = v8
        _ = v9
    } else {
        if (4294967295 != 42) {
            v10 := []int{9223372036854775807, 777, -235, -965}
            if (602 != 2) {
                var v11 int32 = 433
                _ = len(v10)
                fmt.Println(v11)
                _ = v11
            }
            var v12 int16 = -196
            _ = v10
            _ = v12
        }
    }
    v13 := []int{}
    var v14 uint32 = 1
    v15 := GenS1{F0: (1000 <= 0), F1: 0, F2: 494}
    _ = v13
    _ = v14
    _ = v15
    return (^(v14))
}
func genFn1() string {
    var v16 uint64 = 42
    var v17 int64 = 1
    v18 := GenS1{F0: false, F1: 1, F2: 1}
    _ = v16
    _ = v17
    _ = v18
    return "\n\t"
}

func main() {
    var v1 int64 = 370
    var v2 int64 = (v1 - v1)
    var v3 uint32 = 2147483648
    if true {
        switch "hello" {
        case "foo bar":
            var v19 float64 = ((3.14159 - ((0.1 * (-2.25 + 1.5)) / (float64(float32((1.5 + -2.25))) * 0.1))) * float64(float32(-2.25)))
            _ = v19
        case "":
            fmt.Println(v1)
        default:
            var v20 int8 = -128
            _ = v20
        }
        v1 += (v2 / (((v2 & (v2 << ((18446744073709551615) & 63)))) | 1))
        v21 := GenS1{F0: true, F1: 2147483647, F2: 255}
        _ = v21
    } else {
        v22 := GenS2{F0: true}
        _ = v22
    }
    if ((!true) && ((2147483647 <= 0) && (607 == -1))) {
        var v23 uint8 = 250
        _ = v23
    }
    v3 = v3
    v24 := []int{-294}
    v24 = append(v24, 42)
    var v25 uint = 532
    v26 := GenS1{F0: true, F1: 42, F2: 42}
    var v27 uint8 = 133
    if (genFn0() != (v3 * (v3 - genFn0()))) {
        switch 443 {
        case 702:
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 int16 = -32768
                _ = len(v24)
                _ = fi28
                _ = v29
            }
        case 1:
            if false {
                var v30 uint16 = 255
                var v31 float64 = 0.1
                _ = v30
                _ = v31
            }
        case -2147483648:
            v3 += (v3 % ((4294967295) | 1))
        }
    } else {
        _ = len(v24)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
