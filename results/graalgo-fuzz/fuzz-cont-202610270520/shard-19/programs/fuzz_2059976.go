package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 string; F1 uint64; F2 uint16 }
type GenS5 struct { F0 string; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int; F1 uint64; F2 int8 }
type GenS8 struct { F0 string; F1 uint32 }

func main() {
    var v1 uint32 = 65535
    var v2 float32 = -2.25
    v3 := []int{9223372036854775807, -464}
    var v4 int = 32767
    var v5 bool = (false || (false || (true || false)))
    v5 = (v4 < v4)
    if ((v2 - v2) <= (((3.14159 + v2) - float32(1e10)) - (v2 * v2))) {
        if (1 >= 42) {
            if (true && false) {
                var v6 int8 = 127
                v7 := GenS1{F0: v1}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v6
                _ = v7
            }
            v8 := GenS1{F0: (v1 / (((^(v1))) | 1))}
            switch (v4 + 128) {
            case 9223372036854775807:
                v5 = (65535 == 2)
            case 207:
                switch 1 {
                case 127:
                    fmt.Println(v5, v4)
                case 2147483647:
                    v9 := GenS2{F0: 127, F1: 32767}
                    _ = v9
                }
            case 2147483647:
                if ((^(v1)) != (v1 ^ v1)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch (^(v4)) {
                    case -9223372036854775808:
                        switch "a" {
                        case "\n\t":
                            switch v4 {
                            case -1:
                                var v10 uint32 = (^(v1))
                                _ = v10
                            case 32768:
                                var v11 float32 = v2
                                _ = v11
                            }
                        default:
                            if (1 >= 4294967295) {
                                if (v1 < (v1 >> ((255) & 63))) {
                                    var v12 int16 = 14
                                    for fi13 := 0; fi13 < 6; fi13++ {
                                        for fi14 := 0; fi14 < 17; fi14++ {
                                            v5 = ((v1 | (v1 % ((4294967295) | 1))) == (v1 ^ (v1 & v1)))
                                            var v15 uint8 = 128
                                            switch (fi14 >> ((42) & 63)) {
                                            case 4294967295:
                                                _ = len(v3)
                                            default:
                                                if (((fi14 << ((1) & 63)) == fi14) || v5) {
                                                    if len(v3) > 0 {
                                                        v3[0] = -871
                                                    }
                                                    if (957 < (fi14 ^ fi13)) {
                                                        v1 -= 2147483648
                                                    } else {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = v4
                                                        }
                                                        var v16 int = (v4 ^ (fi14 / ((fi14) | 1)))
                                                        v17 := GenS1{F0: v1}
                                                        _ = v16
                                                        _ = v17
                                                    }
                                                } else {
                                                    fmt.Println(v5, v1, v4)
                                                    if false {
                                                        fmt.Println(v2)
                                                        v18 := GenS2{F0: 0, F1: 42}
                                                        fmt.Println(v4, v4, v5)
                                                        _ = v18
                                                    }
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v4 = (v4 | 0)
                                            }
                                            _ = fi14
                                            _ = v15
                                        }
                                        v19 := GenS1{F0: v1}
                                        _ = fi13
                                        _ = v19
                                    }
                                    fmt.Println(v2, v4, v12)
                                    _ = v12
                                } else {
                                    var v20 int32 = 42
                                    for fi21 := 0; fi21 < 16; fi21++ {
                                        switch (v4 | (fi21 * fi21)) {
                                        case 42:
                                            var v22 string = ""
                                            _ = v22
                                        }
                                        var v23 int64 = 9223372036854775807
                                        _ = fi21
                                        _ = v23
                                    }
                                    v20 -= 42
                                    _ = v20
                                }
                                fmt.Println(v5, v5)
                            }
                        }
                    case 128:
                        if v5 {
                            var v24 uint16 = 255
                            if false {
                                var v25 int16 = 250
                                v4 += 67
                                if (-1 <= -66) {
                                    if len(v3) > 0 {
                                        v3[0] = (v4 | v4)
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch -2147483648 {
                                        case 0:
                                            for fi26 := 0; fi26 < 6; fi26++ {
                                                fmt.Println(fi26, fi26)
                                                _ = fi26
                                            }
                                        case 2147483647:
                                            fmt.Println(v5)
                                        case 32768:
                                            v27 := GenS3{F0: 220, F1: 4294967295}
                                            _ = v27
                                        default:
                                            fmt.Println(v4, v24, v25)
                                        }
                                    case "":
                                        v28 := GenS4{F0: ("hello" + "\n\t"), F1: 38, F2: 0}
                                        _ = v28
                                    case "hello":
                                        v1 += (v1 & v1)
                                    }
                                }
                                _ = v25
                            } else {
                                for fi29 := 0; fi29 < 7; fi29++ {
                                    var v30 int32 = -145
                                    v30 = 1
                                    _ = fi29
                                    _ = v30
                                }
                                for fi31 := 0; fi31 < 9; fi31++ {
                                    if (!v5) {
                                        var v32 bool = (354 != 2)
                                        var v33 float64 = 1e10
                                        _ = v32
                                        _ = v33
                                    } else {
                                        v2 = ((float32(1e10) / v2) * -2.25)
                                        if len(v3) > 0 {
                                            v3[0] = v4
                                        }
                                    }
                                    fmt.Println(v4, fi31)
                                    var v34 float32 = v2
                                    _ = fi31
                                    _ = v34
                                }
                            }
                            _ = v24
                        }
                    default:
                        if true {
                            fmt.Println(v5, v5)
                            var v35 int16 = -32768
                            switch "a" {
                            case "a":
                                v1 = 0
                            case "foo bar":
                                v35 += -645
                            default:
                                v36 := GenS5{F0: "foo bar", F1: 203, F2: (v35 & v35)}
                                _ = v36
                            }
                            _ = v35
                        }
                    }
                    if (-1 < -1) {
                        if (!((1e10 + (float64((((v2 * (((v2 * v2) - (v2 / 1.5)) + (((v2 - (v2 * (v2 - 3.14159))) - 1.5) * ((v2 * v2) / 1.0)))) - 1e10) - v2)) * -2.25)) < float64((v2 + v2)))) {
                            switch "foo bar" {
                            case "a":
                                v3 = append(v3, (^(v4)))
                            case "":
                                var v37 bool = (!true)
                                _ = v37
                            default:
                                v38 := GenS6{F0: 65535}
                                _ = v38
                            }
                            v2 += (v2 * v2)
                            for fi39 := 0; fi39 < 4; fi39++ {
                                for fi40 := 0; fi40 < 12; fi40++ {
                                    v5 = (v1 >= (^(v1)))
                                    fmt.Println(v2, fi39)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi40
                                }
                                v5 = (!v5)
                                _ = fi39
                            }
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        var v41 int32 = -65
                        var v42 uint8 = 2
                        _ = v41
                        _ = v42
                    }
                } else {
                    if ((v1 ^ (v1 * (v1 / ((v1) | 1)))) >= v1) {
                        switch (v4 % (((v4 >> ((42) & 63))) | 1)) {
                        case -625:
                            v4 = 130
                        default:
                            if len(v3) > 0 {
                                v3[0] = v4
                            }
                        }
                    } else {
                        v2 -= v2
                        if false {
                            var v43 int32 = 1
                            var v44 int = (v4 * -652)
                            v4 = 128
                            _ = v43
                            _ = v44
                        }
                    }
                    v3 = append(v3, v4)
                }
            }
            _ = v8
        } else {
            v45 := GenS7{F0: 193, F1: 18446744073709551615, F2: 0}
            var v46 uint32 = 1
            _ = v45
            _ = v46
        }
        v5 = false
    } else {
        var v47 uint8 = 127
        fmt.Println(v5, v2, v2)
        _ = v47
    }
    switch "foo bar" {
    case "":
        var v48 uint32 = v1
        _ = v48
    case "a":
        v2 += float32(1e10)
    case "foo bar":
        if v5 {
            v49 := GenS8{F0: "foo bar", F1: v1}
            _ = v49
        }
    default:
        var v50 int32 = -1
        _ = v50
    }
    v3 = append(v3, (v4 | (v4 >> ((0) & 63))))
    v5 = v5
    var v51 int32 = 2147483647
    v3 = append(v3, (^(v4)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
