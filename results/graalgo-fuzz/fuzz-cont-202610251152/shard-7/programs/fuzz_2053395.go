package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 uint8; F2 int }
type GenS7 struct { F0 string; F1 uint32 }

func main() {
    var v1 string = "a"
    var v2 uint8 = 54
    v3 := make(map[int]uint32)
    var v4 int32 = 1
    v5 := []uint32{2147483648, 1}
    v6 := make(map[uint32]bool)
    var v7 int16 = 32767
    var v8 int32 = v4
    switch 4294967295 {
    case 9223372036854775807:
        v9 := GenS1{F0: 128, F1: uint16(v2)}
        _ = v9
    case 32767:
        var v10 string = v1
        _ = v10
    case -196:
        var v11 int16 = v7
        _ = v11
    }
    v12 := GenS2{F0: 0, F1: v4, F2: v7}
    v6[2] = (!false)
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint64 = 4294967295
        v2 += v2
        _ = fi13
        _ = v14
    }
    switch v1 {
    case "foo bar":
        for fi15 := 0; fi15 < 7; fi15++ {
            _ = len(v5)
            var v16 float64 = 1.5
            _ = fi15
            _ = v16
        }
    case "\n\t":
        switch "\n\t" {
        case "hello":
            v8 += (v4 & v4)
        case "foo bar":
            var v17 float64 = (0.0 / 3.14159)
            _ = v17
        }
    case "":
        for fi18 := 0; fi18 < 1; fi18++ {
            if (976 > 174) {
                if ((!(185 > 1)) && (!false)) {
                    v19 := GenS3{F0: (v7 & v7), F1: 4294967295}
                    v20 := GenS4{F0: v1, F1: v8}
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS3{F0: v7, F1: 4294967295}
                    _ = v21
                }
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 float64 = (((0.0 * ((1e10 / 3.14159) / 1.5)) + 0.0) - (1e10 / 3.14159))
                    v24 := GenS1{F0: 0, F1: 0}
                    var v25 int64 = 0
                    var v26 string = "hello"
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
            v5 = append(v5, 640)
            if (!true) {
                v4 = (v8 - 42)
            } else {
                for fi27 := 0; fi27 < 8; fi27++ {
                    var v28 uint = 4294967295
                    var v29 int32 = v4
                    delete(v6, 0)
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                fmt.Println(v4)
            }
            _ = fi18
        }
    }
    var v30 int = 128
    var v31 int64 = -1
    v7 += int16(v31)
    switch "" {
    case "":
        var v32 uint32 = 185
        _ = v32
    case "a":
        var v33 int64 = v31
        _ = v33
    }
    v1 = v1
    if ((v4 / (((v4 / ((v4) | 1))) | 1)) > v8) {
        var v34 bool = (v1 >= v1)
        for fi35 := 0; fi35 < 6; fi35++ {
            for fi36 := 0; fi36 < 14; fi36++ {
                v37 := GenS4{F0: ((v1 + (v1 + (v1 + (v1 + v1)))) + ((v1 + "hello") + ((((((v1 + v1) + ("" + "")) + v1) + ((v1 + "\n\t") + (v1 + v1))) + v1) + "hello"))), F1: (v4 | (v4 - 2147483647))}
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        v38 := GenS5{F0: 95}
        _ = v34
        _ = v38
    } else {
        v39 := GenS6{F0: (v31 + 809), F1: (v2 - 0), F2: 9223372036854775807}
        _ = v39
    }
    if ((v31 != (^(v31))) && (!(!((v31 * -9223372036854775808) != v31)))) {
        var v40 int16 = 42
        var v41 int64 = (v31 & (^(v31)))
        _ = v40
        _ = v41
    } else {
        v31 -= (v31 % ((v31) | 1))
        for fi42 := 0; fi42 < 12; fi42++ {
            delete(v6, 0)
            _ = fi42
        }
    }
    v43 := GenS7{F0: (v1 + v1), F1: 369}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
