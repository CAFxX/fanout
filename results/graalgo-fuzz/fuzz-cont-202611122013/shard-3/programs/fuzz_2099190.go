package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int64 = 0
    var v2 uint = 4294967295
    var v3 float32 = 0.1
    var v4 float64 = 0.0
    if false {
        switch "\n\t" {
        case "":
            v5 := GenS1{F0: (!false), F1: 2}
            _ = v5
        case "a":
            v6 := []int{-9223372036854775808, 162, 32768, 751}
            _ = v6
        case "foo bar":
            switch "\n\t" {
            case "":
                v7 := GenS2{F0: 1, F1: (v2 & v2), F2: -43}
                _ = v7
            case "foo bar":
                var v8 int8 = 42
                _ = v8
            case "hello":
                for fi9 := 0; fi9 < 18; fi9++ {
                    fmt.Println(v1, v3)
                    var v10 uint16 = 967
                    v10 = (^(v10))
                    _ = fi9
                    _ = v10
                }
            }
        }
        v11 := []int{}
        fmt.Println(v2)
        _ = v11
    }
    v12 := GenS3{F0: 150, F1: 18446744073709551615}
    v3 = (v3 * v3)
    v13 := []int{104, -430, 0, 4294967295}
    if (852 <= 32767) {
        var v14 int8 = -128
        switch "foo bar" {
        case "\n\t":
            if true {
                v15 := GenS4{F0: (-(v1)), F1: 4294967295, F2: (2147483647 >= 1)}
                v13 = append(v13, 1)
                switch "foo bar" {
                case "a":
                    v14 += (v14 << ((255) & 63))
                default:
                    if (false || true) {
                        v2 = v2
                        var v16 string = "\n\t"
                        v17 := GenS3{F0: 2, F1: 2}
                        _ = v16
                        _ = v17
                    } else {
                        var v18 int = 0
                        _ = v18
                    }
                }
                _ = v15
            }
        case "foo bar":
            _ = len(v13)
        default:
            if (!false) {
                if ((v2 >> (((v2 & v2)) & 63)) > v2) {
                    fmt.Println(v1)
                    var v19 float32 = (0.1 / (v3 + 3.14159))
                    switch "\n\t" {
                    case "":
                        if false {
                            var v20 uint32 = 910
                            v1 = (^(v1))
                            fmt.Println(v20, v14)
                            _ = v20
                        } else {
                            v21 := GenS3{F0: uint8(v2), F1: 883}
                            v13 = append(v13, 2147483647)
                            fmt.Println(v4)
                            _ = v21
                        }
                    }
                    _ = v19
                } else {
                    v3 -= v3
                    if false {
                        var v22 int8 = (v14 ^ (^(v14)))
                        v23 := GenS5{F0: 365, F1: 18446744073709551615}
                        for fi24 := 0; fi24 < 15; fi24++ {
                            v25 := GenS6{F0: 623}
                            for fi26 := 0; fi26 < 20; fi26++ {
                                var v27 uint64 = 2
                                var v28 bool = (true || true)
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            v2 = (v2 + 42)
                            _ = fi24
                            _ = v25
                        }
                        _ = v22
                        _ = v23
                    }
                }
            } else {
                var v29 int = 128
                v30 := GenS7{F0: 42, F1: v1}
                _ = v29
                _ = v30
            }
        }
        _ = v14
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        _ = len(v13)
        _ = fi31
    }
    var v32 uint16 = 134
    v32 += 367
    var v33 int16 = 1000
    if len(v13) > 0 {
        v13[0] = 4294967295
    }
    v34 := GenS1{F0: false, F1: 982}
    v32 -= (^(v32))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
