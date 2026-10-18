package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int16 }
func genFn0() int8 {
    var v5 int16 = 1
    v5 = 1000
    switch 351 {
    case 630:
        v5 = v5
    }
    fmt.Println(v5)
    _ = v5
    return 97
}
func genFn1(p0 bool) bool {
    if p0 {
        var v6 string = ""
        _ = v6
    } else {
        for fi7 := 0; fi7 < 15; fi7++ {
            p0 = p0
            for fi8 := 0; fi8 < 2; fi8++ {
                p0 = (((255 != 89) && (((true || false) && true) && (p0 && p0))) && p0)
                switch "" {
                case "\n\t":
                    fmt.Println(fi7, p0)
                case "foo bar":
                    v9 := GenS1{F0: 145}
                    _ = v9
                default:
                    v10 := GenS2{F0: 0, F1: 682, F2: 461}
                    _ = v10
                }
                var v11 int = (fi7 << ((18446744073709551615) & 63))
                var v12 uint8 = 127
                v12 -= (v12 << ((683) & 63))
                _ = fi8
                _ = v11
                _ = v12
            }
            _ = fi7
        }
    }
    v13 := []int{128}
    var v14 uint16 = 42
    if p0 {
        if p0 {
            if p0 {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                if p0 {
                    p0 = (1e10 <= (1e10 / (float64((3.14159 / -2.25)) + 0.1)))
                } else {
                    fmt.Println(p0)
                    v15 := GenS3{F0: 374}
                    v14 = v14
                    _ = v15
                }
            } else {
                for fi16 := 0; fi16 < 12; fi16++ {
                    v14 += v14
                    p0 = p0
                    _ = fi16
                }
            }
            if p0 {
                for fi17 := 0; fi17 < 20; fi17++ {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        var v19 int32 = 2147483647
                        p0 = (p0 || p0)
                        _ = fi18
                        _ = v19
                    }
                    _ = fi17
                }
                fmt.Println(v14)
                switch "a" {
                case "\n\t":
                    v20 := GenS3{F0: -1}
                    _ = v20
                case "":
                    p0 = (!false)
                }
            } else {
                var v21 int8 = 1
                _ = v21
            }
        }
        v14 += (v14 / ((0) | 1))
    }
    _ = p0
    _ = v13
    _ = v14
    return p0
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 bool = false
    var v3 uint16 = 32767
    var v4 int64 = v1
    v22 := []int{0, 32768}
    var v23 uint64 = 466
    v4 -= int64(v3)
    switch "hello" {
    case "\n\t":
        var v24 uint32 = 245
        _ = v24
    }
    if false {
        fmt.Println(v2, v3, v4)
        var v25 int64 = (v1 ^ -9223372036854775808)
        _ = v25
    } else {
        var v26 float32 = 0.1
        v23 += v23
        var v27 int64 = -1
        _ = v26
        _ = v27
    }
    v28 := GenS3{F0: 1000}
    var v29 uint8 = 0
    for fi30 := 0; fi30 < 14; fi30++ {
        fmt.Println(v2)
        var v31 string = ("" + "hello")
        _ = len(v22)
        _ = fi30
        _ = v31
    }
    fmt.Println(v3, v23, v2)
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    v2 = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
