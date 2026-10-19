package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string; F2 int16 }

func main() {
    var v1 int8 = 104
    v2 := make(map[uint32]bool)
    var v3 int64 = 1
    var v4 int16 = 1000
    var v5 int32 = -2147483648
    var v6 uint = 1
    fmt.Println(v5, v5)
    var v7 int64 = 268
    v8 := v2[65535]
    switch -9223372036854775808 {
    case -617:
        v9 := []int{42, 4294967295, -9223372036854775808, -746}
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS1{F0: v1}
        var v12 bool = v8
        var v13 uint8 = 128
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    if v8 {
        var v14 int8 = v1
        _ = v14
    }
    var v15 bool = (v1 > v1)
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            v18 := []int{(fi17 / ((628) | 1)), -876, 42, (fi17 - (fi16 * 32768))}
            _ = fi17
            _ = v18
        }
        v5 = (v5 - 973)
        v19 := []int{fi16, fi16, fi16}
        _ = fi16
        _ = v19
    }
    v4 -= -290
    for fi20 := 0; fi20 < 13; fi20++ {
        for fi21 := 0; fi21 < 8; fi21++ {
            delete(v2, 2147483647)
            if (!(1 > 128)) {
                if v8 {
                    v22 := v2[0]
                    v23 := GenS1{F0: (v1 * (v1 % (((v1 / (((v1 << ((v6) & 63))) | 1))) | 1)))}
                    _ = v22
                    _ = v23
                } else {
                    delete(v2, 0)
                    v7 = (v7 - 379)
                }
            } else {
                v24 := GenS2{F0: v6, F1: (("" + "foo bar") + ""), F2: v4}
                fmt.Println(fi21, fi21)
                _ = v24
            }
            _ = fi21
        }
        _ = fi20
    }
    v25 := []int{765, 42, -386}
    for fi26 := 0; fi26 < 9; fi26++ {
        v3 -= -1
        v3 = v7
        switch "" {
        case "a":
            v27 := GenS1{F0: (v1 / ((-71) | 1))}
            _ = v27
        case "\n\t":
            var v28 uint16 = 2
            _ = v28
        case "foo bar":
            var v29 int16 = 1000
            _ = v29
        default:
            var v30 int32 = v5
            _ = v30
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
