package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 string; F2 int64 }

func main() {
    var v1 uint8 = 2
    var v2 uint = 1
    var v3 int64 = 528
    v4 := GenS1{F0: (uint16(v1) == 1), F1: -128}
    if true {
        v5 := []int{32768}
        var v6 int = 537
        _ = v5
        _ = v6
    } else {
        v7 := GenS2{F0: (v3 / (((-(v3))) | 1)), F1: 47, F2: (10 < 2)}
        _ = v7
    }
    var v8 uint = v2
    if (true || ((660 > 42) || (true || (1 == 42)))) {
        v8 = 0
    } else {
        v9 := GenS1{F0: true, F1: -1}
        var v10 int16 = -934
        _ = v10
        _ = v9
    }
    var v11 int32 = 2147483647
    v12 := []int{-163}
    fmt.Println(v3, v8)
    var v13 string = (("" + "hello") + "")
    var v14 int = 32768
    var v15 int8 = -50
    if len(v12) > 0 {
        v12[0] = (^(v14))
    }
    switch v14 {
    case 127:
        if false {
            if true {
                v16 := GenS3{F0: v14}
                v1 = v1
                if len(v12) > 0 {
                    v12[0] = int(v3)
                }
                _ = v16
            }
            v17 := GenS2{F0: -1, F1: (v15 / ((127) | 1)), F2: (319 < 756)}
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 uint16 = 500
                _ = fi18
                _ = v19
            }
            _ = v17
        }
    default:
        for fi20 := 0; fi20 < 7; fi20++ {
            fmt.Println(v15, v11)
            _ = fi20
        }
    }
    var v21 float32 = 0.1
    if len(v12) > 0 {
        v12[0] = v14
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        fmt.Println(v14, v15, v8)
        v13 = ("\n\t" + v13)
        _ = fi22
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    if (!(float32(-2.25) <= v21)) {
        switch -95 {
        case 9223372036854775807:
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        case 32768:
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 uint = 65535
                _ = fi23
                _ = v24
            }
        case 361:
            v1 -= 2
        }
    } else {
        for fi25 := 0; fi25 < 3; fi25++ {
            v12 = append(v12, v14)
            fmt.Println(v21, v15, v15)
            _ = fi25
        }
        if false {
            for fi26 := 0; fi26 < 16; fi26++ {
                v2 -= (v2 % (((v2 << ((v8) & 63))) | 1))
                _ = fi26
            }
            var v27 int = (-(v14))
            _ = v27
        }
        v28 := GenS4{F0: 2, F1: v13, F2: (v3 - 0)}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v8
}
