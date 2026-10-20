package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 float32 = -2.25
    var v2 uint64 = 1
    var v3 int16 = 42
    var v4 int32 = 1
    v5 := make(map[string]int)
    var v6 float32 = v1
    var v7 uint8 = 128
    fmt.Println(v7)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 string = ""
        if true {
            for fi10 := 0; fi10 < 3; fi10++ {
                fmt.Println(v9)
                if ((fi8 | 127) <= 36) {
                    var v11 uint32 = 83
                    _ = v11
                }
                v12 := GenS1{F0: true, F1: -1, F2: 2}
                _ = fi10
                _ = v12
            }
            v6 = (v1 * (v1 - v6))
            var v13 int8 = 100
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    v14 := GenS2{F0: uint32(v7), F1: 708}
    v1 -= v6
    v15 := GenS1{F0: false, F1: -834, F2: v7}
    if false {
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 int8 = -51
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: ((32768 >= 128) && false), F1: 42, F2: (v7 - uint8(v2))}
        _ = v18
    } else {
        var v19 string = ""
        for fi20 := 0; fi20 < 18; fi20++ {
            v1 = ((1e10 * (-2.25 + v1)) + 1e10)
            v21 := []int{-1, (fi20 << ((1) & 63)), 604}
            _ = fi20
            _ = v21
        }
        switch -1 {
        case -261:
            v22 := GenS3{F0: (v7 ^ v7), F1: 143}
            _ = v22
        case 32767:
            var v23 uint64 = v2
            _ = v23
        case 0:
            var v24 uint = 4294967295
            _ = v24
        default:
            v25 := GenS1{F0: true, F1: 9223372036854775807, F2: (^(v7))}
            _ = v25
        }
        _ = v19
    }
    fmt.Println(v4, v6)
    fmt.Println(v3)
    v4 += 2147483647
    var v26 bool = false
    if (!("hello" == "")) {
        switch 128 {
        case 4294967295:
            v27 := []int{32767, 0, 128}
            _ = v27
        case 32767:
            var v28 uint8 = v7
            _ = v28
        default:
            var v29 int8 = 114
            _ = v29
        }
        fmt.Println(v7, v2)
        v30 := GenS1{F0: (716 >= 255), F1: 0, F2: (v7 & (v7 >> ((65535) & 63)))}
        _ = v30
    } else {
        if ((32767 >= -175) || (!(v6 <= v1))) {
            fmt.Println(v1, v1)
        } else {
            for fi31 := 0; fi31 < 7; fi31++ {
                fmt.Println(v4, v26)
                var v32 uint = 255
                _ = fi31
                _ = v32
            }
            v33 := []int{32767}
            _ = v33
        }
        var v34 string = "\n\t"
        _ = v34
    }
    if (-63 < 18) {
        v1 = (v1 * 1e10)
        v7 = (^(v7))
    } else {
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 uint16 = 240
            _ = fi35
            _ = v36
        }
        v3 += (^(v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
