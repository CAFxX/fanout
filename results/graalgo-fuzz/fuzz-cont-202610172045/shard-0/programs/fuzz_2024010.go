package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint8; F2 bool }
type GenS4 struct { F0 uint64; F1 int16; F2 int64 }
func genFn0() int8 {
    var v4 int32 = -2147483648
    switch "foo bar" {
    case "":
        v5 := []int{}
        _ = v5
    case "foo bar":
        if (18446744073709551615 < 18446744073709551615) {
            fmt.Println(v4)
        } else {
            v4 = (v4 >> ((255) & 63))
        }
    }
    var v6 int32 = (v4 % (((v4 % ((-2147483648) | 1))) | 1))
    var v7 int64 = 0
    _ = v4
    _ = v6
    _ = v7
    return 127
}
func genFn1() int64 {
    var v8 bool = true
    switch 32768 {
    case -79:
        var v9 uint = 65535
        _ = v9
    case 2:
        fmt.Println(v8)
    default:
        for fi10 := 0; fi10 < 17; fi10++ {
            v8 = ((65535 < 2) && v8)
            v11 := []int{9223372036854775807}
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v8)
    v8 = ((!(!(42 < -1))) && false)
    _ = v8
    return 42
}

func main() {
    var v1 int = 127
    var v2 uint32 = 227
    v3 := make(map[string]bool)
    v1 -= (v1 % (((v1 ^ (-(v1)))) | 1))
    var v12 int64 = genFn1()
    if true {
        v13 := GenS1{F0: 1}
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := GenS2{F0: 973}
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        v16 := []int{}
        switch v1 {
        case 128:
            if true {
                fmt.Println(v12, v12)
                fmt.Println(v12, v12)
            } else {
                v17 := GenS3{F0: 32767, F1: 185, F2: true}
                if false {
                    var v18 uint = 236
                    _ = v18
                } else {
                    fmt.Println(v12)
                    var v19 float64 = 3.14159
                    for fi20 := 0; fi20 < 18; fi20++ {
                        v1 = v1
                        _ = fi20
                    }
                    _ = v19
                }
                _ = len(v3)
                _ = v17
            }
        case 845:
            _ = len(v16)
        }
        _ = v16
    }
    _ = len(v3)
    for fi21 := 0; fi21 < 13; fi21++ {
        switch 0 {
        case -235:
            v12 -= v12
        case 2147483647:
            v2 = (v2 - v2)
        case 42:
            var v22 uint16 = 351
            _ = v22
        default:
            for fi23 := 0; fi23 < 12; fi23++ {
                v12 = v12
                v2 = v2
                _ = fi23
            }
        }
        fmt.Println(fi21, fi21, v12)
        var v24 uint32 = 1
        _ = fi21
        _ = v24
    }
    switch (v1 ^ -2147483648) {
    case 9223372036854775807:
        if true {
            v25 := []int{v1, (v1 / ((v1) | 1))}
            v26 := GenS2{F0: 2}
            for fi27 := 0; fi27 < 3; fi27++ {
                v12 -= (-(v12))
                _ = fi27
            }
            _ = v25
            _ = v26
        }
    case -1:
        v3[(("foo bar" + ("foo bar" + "")) + ("" + "foo bar"))] = true
    }
    v28 := GenS3{F0: 577, F1: 127, F2: false}
    v3["foo bar"] = true
    var v29 int8 = genFn0()
    v30 := GenS4{F0: 42, F1: 0, F2: (^(v12))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
