package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 int }

func main() {
    v1 := []uint32{927, 2147483648, 0, 823}
    var v2 uint32 = 2
    var v3 uint = 0
    var v4 int = 32767
    fmt.Println(v4, v4)
    for fi5 := 0; fi5 < 20; fi5++ {
        v3 -= 65535
        var v6 string = ("" + "foo bar")
        v7 := GenS1{F0: 255, F1: 789}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 351, F1: (v2 << (((v3 / ((42) | 1))) & 63))}
    if true {
        if true {
            for fi9 := 0; fi9 < 9; fi9++ {
                if false {
                    _ = len(v1)
                    v2 = v2
                } else {
                    var v10 float64 = 1e10
                    fmt.Println(v4, fi9, fi9)
                    _ = v10
                }
                _ = fi9
            }
        }
        if (-216 > 9223372036854775807) {
            v2 = (v2 >> ((65535) & 63))
            v11 := GenS1{F0: 255, F1: (v2 << (((v3 << ((v3) & 63))) & 63))}
            _ = v11
        } else {
            var v12 int32 = 42
            _ = v12
        }
    } else {
        fmt.Println(v3)
        v13 := GenS1{F0: uint16(v3), F1: v2}
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int32 = 2147483647
        v16 := GenS2{F0: 42, F1: "a"}
        _ = fi14
        _ = v15
        _ = v16
    }
    if len(v1) > 0 {
        v1[0] = (^(v2))
    }
    var v17 float64 = -2.25
    var v18 int8 = 114
    var v19 uint = v3
    v20 := GenS3{F0: v4, F1: v2}
    switch int(v19) {
    case 128:
        switch (-(v4)) {
        case 2147483647:
            v21 := GenS4{F0: true, F1: 0}
            _ = v21
        case 1:
            var v22 bool = (435 >= 1)
            _ = v22
        case 0:
            v23 := GenS5{F0: -17}
            _ = v23
        default:
            for fi24 := 0; fi24 < 2; fi24++ {
                if false {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        var v26 float64 = v17
                        _ = fi25
                        _ = v26
                    }
                    var v27 int = v4
                    _ = v27
                }
                _ = fi24
            }
        }
    case 9223372036854775807:
        var v28 uint16 = 723
        _ = v28
    }
    v29 := GenS6{F0: v18, F1: v4}
    v3 -= (v19 & v3)
    switch v4 {
    case 42:
        v30 := GenS2{F0: -985, F1: ("hello" + "hello")}
        _ = v30
    case 32767:
        v1 = append(v1, 11)
    case 1:
        v1 = append(v1, 182)
    default:
        for fi31 := 0; fi31 < 10; fi31++ {
            v4 = 574
            _ = fi31
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v8
}
