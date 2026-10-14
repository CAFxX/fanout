package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0() uint8 {
    var v6 int8 = 0
    var v7 int = -9223372036854775808
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int16 = 1000
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := []int{fi10}
        var v12 int = 363
        v13 := GenS1{F0: 128, F1: 42, F2: 65535}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    _ = v6
    _ = v7
    return 253
}
func genFn1(p0 uint32, p1 int, p2 string) uint32 {
    v14 := []int{}
    var v15 int32 = 1
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 uint32 = (p0 / (((p0 / ((2147483647) | 1))) | 1))
        fmt.Println(v15)
        _ = fi16
        _ = v17
    }
    var v18 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v18
    return 42
}

func main() {
    var v1 float64 = 1.5
    var v2 uint16 = 2
    var v3 uint16 = v2
    var v4 bool = false
    var v5 uint16 = v2
    switch "" {
    case "":
        var v19 int8 = 42
        _ = v19
    case "a":
        fmt.Println(v4, v3)
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        if false {
            if ("foo bar" <= "foo bar") {
                v2 += 1
            } else {
                v4 = v4
            }
        } else {
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(fi21)
                _ = fi21
            }
        }
        v3 = (v3 | v3)
        _ = fi20
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        fmt.Println(v3)
    case -1:
        var v22 int64 = 0
        _ = v22
    case 1:
        var v23 uint64 = 1
        _ = v23
    default:
        v24 := GenS1{F0: 4294967295, F1: 2147483647, F2: 2}
        _ = v24
    }
    v1 -= ((float64(1e10) * (v1 + 0.1)) * (v1 * v1))
    fmt.Println(v2, v5, v1)
    var v25 int = 9223372036854775807
    for fi26 := 0; fi26 < 10; fi26++ {
        v27 := GenS2{F0: 1, F1: (v25 ^ v25)}
        _ = fi26
        _ = v27
    }
    v25 = v25
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 uint8 = 42
        for fi30 := 0; fi30 < 9; fi30++ {
            fmt.Println(fi28, v29, fi30)
            fmt.Println(fi28)
            v5 += v2
            _ = fi30
        }
        for fi31 := 0; fi31 < 3; fi31++ {
            v32 := []int{-695, (fi31 + 0)}
            v33 := GenS1{F0: 127, F1: -2147483648, F2: 130}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi28
        _ = v29
    }
    v1 = (1.5 * v1)
    fmt.Println(v3, v2, v5)
    var v34 uint16 = 2
    fmt.Println(v2, v1)
    var v35 int16 = -45
    if false {
        v36 := GenS1{F0: 4294967295, F1: 42, F2: 255}
        fmt.Println(v1, v25, v3)
        v37 := GenS3{F0: 2, F1: 1}
        _ = v36
        _ = v37
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
