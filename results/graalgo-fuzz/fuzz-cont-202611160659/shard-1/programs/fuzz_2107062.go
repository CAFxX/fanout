package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int8; F2 bool }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 float32 = -2.25
    var v2 uint = 0
    var v3 uint = (v2 & v2)
    var v4 float64 = 1e10
    v5 := []uint32{}
    for fi6 := 0; fi6 < 20; fi6++ {
        v2 = v3
        v1 -= v1
        _ = fi6
    }
    v3 -= v2
    fmt.Println(v4, v2, v4)
    var v7 float64 = v4
    _ = len(v5)
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            if (2 == 4294967295) {
                v1 -= 3.14159
                if (!false) {
                    _ = len(v5)
                    switch fi9 {
                    case 9223372036854775807:
                        fmt.Println(v2, fi9, v3)
                        _ = len(v5)
                        var v10 int8 = -107
                        _ = v10
                    default:
                        if false {
                            var v11 uint32 = 2147483647
                            _ = v11
                        } else {
                            var v12 uint16 = 32767
                            var v13 uint16 = 477
                            _ = v12
                            _ = v13
                        }
                    }
                } else {
                    v7 = v7
                    switch 9223372036854775807 {
                    case 1:
                        var v14 float64 = float64(v1)
                        _ = v14
                    case -1:
                        var v15 int = fi8
                        _ = v15
                    }
                    v16 := GenS1{F0: true, F1: 2}
                    _ = v16
                }
                var v17 float64 = 1.5
                _ = v17
            } else {
                var v18 string = ((("foo bar" + ("\n\t" + "\n\t")) + "\n\t") + "\n\t")
                _ = len(v5)
                if (542 < 448) {
                    if len(v5) > 0 {
                        v5[0] = 2
                    }
                    switch fi8 {
                    case 0:
                        v19 := GenS2{F0: 4294967295, F1: 65535}
                        var v20 uint8 = 42
                        var v21 uint16 = 0
                        _ = v19
                        _ = v20
                        _ = v21
                    case 42:
                        switch fi9 {
                        case -55:
                            var v22 uint32 = 65535
                            _ = v22
                        case 868:
                            v5 = append(v5, 0)
                        case -167:
                            v7 = v7
                            switch v18 {
                            case "\n\t":
                                if len(v5) > 0 {
                                    v5[0] = 0
                                }
                            default:
                                var v23 uint32 = 1
                                _ = v23
                            }
                            var v24 string = (v18 + v18)
                            _ = v24
                        }
                    default:
                        v1 = v1
                    }
                    var v25 int32 = 1
                    _ = v25
                }
                v26 := GenS1{F0: ((961 < 0) && (false && false)), F1: (^(v2))}
                _ = v18
                _ = v26
            }
            var v27 uint8 = 1
            var v28 float32 = v1
            _ = fi9
            _ = v27
            _ = v28
        }
        _ = fi8
    }
    v7 = 0.0
    if true {
        v3 -= (v3 << ((0) & 63))
        v5 = append(v5, 65535)
    } else {
        v29 := GenS3{F0: 0, F1: -81, F2: 908}
        if ("" >= "hello") {
            var v30 bool = (!(2 <= 2))
            v5 = append(v5, 991)
            v31 := GenS4{F0: 18446744073709551615, F1: -1, F2: true}
            _ = v30
            _ = v31
        }
        _ = v29
    }
    v32 := GenS5{F0: uint16(v3), F1: 1}
    v4 -= v7
    fmt.Println(v4)
    v33 := GenS4{F0: 356, F1: 1, F2: (false || ((-1 > 2147483647) || (!true)))}
    for fi34 := 0; fi34 < 16; fi34++ {
        for fi35 := 0; fi35 < 19; fi35++ {
            if true {
                switch fi34 {
                case 32767:
                    fmt.Println(v3)
                case 9223372036854775807:
                    v2 = (v3 - (v2 | (v3 * 1)))
                case -2147483648:
                    v1 = 1.5
                default:
                    if ((-786 == fi34) || ((2 == 395) || (-327 >= -474))) {
                        v5 = append(v5, 1)
                        fmt.Println(v2, fi34, v7)
                    }
                }
                fmt.Println(v3, v3)
            } else {
                var v36 uint = (v3 ^ (^(v2)))
                _ = v36
            }
            v4 += (v7 * ((v4 + (v4 / (v7 * 0.0))) * v7))
            _ = len(v5)
            _ = fi35
        }
        _ = fi34
    }
    if true {
        v7 -= (1.5 / v4)
        v3 -= v2
    }
    v37 := GenS6{F0: 255, F1: 2147483647}
    v7 -= v4
    fmt.Println(v3)
    v38 := GenS5{F0: 0, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v7
}
