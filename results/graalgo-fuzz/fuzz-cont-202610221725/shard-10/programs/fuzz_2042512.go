package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float32 = (3.14159 / 1e10)
    var v2 int = 32768
    v3 := []uint32{}
    var v4 uint8 = 1
    v5 := GenS1{F0: 65535, F1: 1}
    if (false || (true && true)) {
        var v6 int64 = -507
        v2 = v2
        _ = v6
    } else {
        v1 = v1
        var v7 float64 = 1e10
        _ = v7
    }
    var v8 float64 = -2.25
    if len(v3) > 0 {
        v3[0] = 2
    }
    fmt.Println(v2, v4)
    switch v2 {
    case 4294967295:
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: 2147483647, F1: -32768}
            for fi11 := 0; fi11 < 14; fi11++ {
                if true {
                    v3 = append(v3, 928)
                    var v12 int8 = 1
                    if false {
                        _ = len(v3)
                    }
                    _ = v12
                } else {
                    var v13 uint = 103
                    var v14 int = fi9
                    _ = v13
                    _ = v14
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v15 := GenS2{F0: int16(v4), F1: 42, F2: 1}
                _ = fi11
                _ = v15
            }
            for fi16 := 0; fi16 < 20; fi16++ {
                if false {
                    v2 -= 9223372036854775807
                    var v17 int64 = 483
                    _ = v17
                }
                if false {
                    fmt.Println(v2, fi16, v2)
                    switch v2 {
                    case -9223372036854775808:
                        v18 := GenS3{F0: 710}
                        _ = v18
                    }
                    _ = len(v3)
                }
                _ = fi16
            }
            _ = fi9
            _ = v10
        }
    case -9223372036854775808:
        var v19 uint32 = 65535
        _ = v19
    case 1:
        fmt.Println(v1, v1, v8)
    default:
        if (false && true) {
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 uint32 = 4294967295
                _ = fi20
                _ = v21
            }
            var v22 uint16 = 255
            _ = v22
        } else {
            var v23 int64 = 784
            _ = v23
        }
    }
    _ = len(v3)
    v1 += float32(((v8 / (((v8 / 1.0) / v8) / (3.14159 / v8))) - (3.14159 + (((-2.25 / 0.1) - (v8 - -2.25)) + (((((v8 / (v8 - -2.25)) + (v8 / 1.5)) - 1.5) * v8) * v8)))))
    v2 = (v2 % (((v2 & (^(v2)))) | 1))
    v24 := GenS1{F0: 1, F1: -448}
    fmt.Println(v2, v1)
    v1 -= (-2.25 * (float32((0.1 / v8)) + v1))
    switch "\n\t" {
    case "foo bar":
        var v25 int8 = 127
        _ = v25
    case "":
        v26 := GenS3{F0: 0}
        _ = v26
    case "hello":
        for fi27 := 0; fi27 < 5; fi27++ {
            fmt.Println(v4, v8, v2)
            for fi28 := 0; fi28 < 18; fi28++ {
                switch 433 {
                case 115:
                    fmt.Println(v2, fi27)
                case 0:
                    if len(v3) > 0 {
                        v3[0] = 483
                    }
                case -701:
                    v2 = fi27
                }
                _ = fi28
            }
            _ = fi27
        }
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        v8 += v8
        v30 := GenS3{F0: 0}
        _ = fi29
        _ = v30
    }
    var v31 uint16 = 851
    fmt.Println(v31)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v8
}
