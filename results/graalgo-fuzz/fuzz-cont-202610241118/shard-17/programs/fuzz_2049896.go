package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 bool }
type GenS2 struct { F0 uint32; F1 bool; F2 int8 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS5 struct { F0 bool; F1 uint32; F2 bool }
type GenS6 struct { F0 uint; F1 uint; F2 uint8 }
type GenS7 struct { F0 int }

func main() {
    var v1 bool = false
    v2 := []string{"hello", "foo bar", (("a" + "foo bar") + ("a" + "a")), ("foo bar" + "a")}
    v3 := make(map[uint32]int)
    var v4 uint16 = 65535
    v1 = (2 >= 86)
    var v5 float32 = (float32(3.14159) - 0.1)
    v5 -= (((v5 - 3.14159) + 1.5) - v5)
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    if (!v1) {
        var v6 int = -2147483648
        fmt.Println(v4, v4, v6)
        for fi7 := 0; fi7 < 17; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(fi7)
                if v1 {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi8
            }
            _ = fi7
        }
        _ = v6
    } else {
        var v9 int64 = -9223372036854775808
        _ = v9
    }
    var v10 float64 = 1e10
    if v1 {
        if ((0.0 - (((v10 * v10) / v10) / v10)) > float64(0.0)) {
            var v11 int = 127
            v11 = v11
            _ = v11
        }
    }
    v1 = true
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 string = ("a" + "\n\t")
        _ = fi12
        _ = v13
    }
    delete(v3, 4294967295)
    v4 = v4
    switch "foo bar" {
    case "a":
        for fi14 := 0; fi14 < 11; fi14++ {
            v4 = 65535
            if v1 {
                var v15 uint = 65535
                switch "\n\t" {
                case "\n\t":
                    if true {
                        switch -1 {
                        case 258:
                            if v1 {
                                v16 := GenS1{F0: (v4 ^ (v4 - (v4 % (((^(v4))) | 1)))), F1: (v1 && v1), F2: ((v4 / (((v4 - v4)) | 1)) != v4)}
                                if (!(!v1)) {
                                    var v17 int8 = 54
                                    var v18 float64 = 3.14159
                                    _ = v17
                                    _ = v18
                                } else {
                                    if ((257 <= 125) && v1) {
                                        var v19 float64 = v10
                                        fmt.Println(v4)
                                        for fi20 := 0; fi20 < 14; fi20++ {
                                            switch ("hello" + ("foo bar" + ("" + "a"))) {
                                            case "a":
                                                var v21 bool = true
                                                _ = v21
                                            case "foo bar":
                                                switch fi20 {
                                                case 2147483647:
                                                    fmt.Println(v1, v1)
                                                default:
                                                    var v22 int8 = 0
                                                    _ = v22
                                                }
                                            case "hello":
                                                v10 -= (v19 + v10)
                                            default:
                                                switch "" {
                                                case "hello":
                                                    if v1 {
                                                        v5 = v5
                                                        fmt.Println(v1)
                                                    }
                                                case "":
                                                    v5 = v5
                                                }
                                            }
                                            if (int8(v15) >= 1) {
                                                switch -1 {
                                                case 854:
                                                    fmt.Println(v4, fi14)
                                                }
                                                v23 := v3[4294967295]
                                                _ = v23
                                            } else {
                                                v1 = (v15 > v15)
                                                var v24 string = ("foo bar" + "a")
                                                switch fi20 {
                                                case -873:
                                                    switch (-(fi14)) {
                                                    case 42:
                                                        if true {
                                                            v24 = "\n\t"
                                                        }
                                                    default:
                                                        switch v24 {
                                                        case "\n\t":
                                                            fmt.Println(v24, fi14, v19)
                                                        case "foo bar":
                                                            if v1 {
                                                                var v25 float64 = 0.0
                                                                fmt.Println(v4, v1, v24)
                                                                v5 += (((float32(v10) + v5) * v5) + (v5 / v5))
                                                                _ = v25
                                                            }
                                                        default:
                                                            if len(v2) > 0 {
                                                                v2[0] = v24
                                                            }
                                                        }
                                                    }
                                                case 267:
                                                    var v26 uint64 = 251
                                                    _ = v26
                                                case 127:
                                                    if (0.1 >= (-2.25 * 1.5)) {
                                                        fmt.Println(v15, v4)
                                                    }
                                                }
                                                _ = v24
                                            }
                                            var v27 int64 = -9223372036854775808
                                            _ = fi20
                                            _ = v27
                                        }
                                        _ = v19
                                    }
                                }
                                v28 := GenS2{F0: 315, F1: false, F2: -128}
                                _ = v16
                                _ = v28
                            } else {
                                var v29 int8 = -103
                                _ = v29
                            }
                        case 957:
                            for fi30 := 0; fi30 < 5; fi30++ {
                                if (!("foo bar" != "hello")) {
                                    var v31 int16 = 588
                                    v32 := GenS3{F0: (v31 + 524), F1: 535}
                                    _ = v31
                                    _ = v32
                                } else {
                                    v1 = false
                                    if v1 {
                                        v3[1] = fi30
                                        v2 = append(v2, "a")
                                    }
                                    var v33 float32 = (v5 - (v5 + ((float32(1.5) * v5) / (1.5 + -2.25))))
                                    _ = v33
                                }
                                fmt.Println(v15)
                                if false {
                                    fmt.Println(v15)
                                    v34 := GenS1{F0: (v4 * v4), F1: (544 == 1), F2: false}
                                    fmt.Println(v5)
                                    var v35 uint64 = 2
                                    _ = v34
                                    _ = v35
                                }
                                _ = fi30
                            }
                        case -995:
                            switch ("foo bar" + "hello") {
                            case "":
                                var v36 uint32 = 339
                                _ = v36
                            case "hello":
                                if v1 {
                                    v10 = v10
                                }
                            default:
                                delete(v3, 26)
                            }
                        }
                        if ((867 > (v4 & 42)) && true) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            for fi37 := 0; fi37 < 13; fi37++ {
                                var v38 uint16 = v4
                                _ = fi37
                                _ = v38
                            }
                            switch "" {
                            case "a":
                                switch (fi14 >> ((v15) & 63)) {
                                case 996:
                                    for fi39 := 0; fi39 < 13; fi39++ {
                                        var v40 bool = (323 > -757)
                                        var v41 int32 = 1
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                case 9223372036854775807:
                                    var v42 uint8 = 128
                                    _ = v42
                                }
                            case "foo bar":
                                if (v1 && v1) {
                                    delete(v3, 447)
                                    v43 := GenS2{F0: 2147483647, F1: v1, F2: 119}
                                    _ = v43
                                } else {
                                    v44 := GenS4{F0: v4, F1: 42, F2: 2}
                                    var v45 int16 = 1
                                    _ = v44
                                    _ = v45
                                }
                            }
                        } else {
                            switch (fi14 ^ 42) {
                            case -788:
                                switch fi14 {
                                case 463:
                                    v4 -= v4
                                case -9223372036854775808:
                                    v46 := GenS5{F0: v1, F1: 2, F2: true}
                                    _ = v46
                                case 784:
                                    fmt.Println(fi14)
                                }
                            }
                            for fi47 := 0; fi47 < 12; fi47++ {
                                v48 := GenS2{F0: 4294967295, F1: (-2147483648 >= 2147483647), F2: 1}
                                _ = fi47
                                _ = v48
                            }
                            v49 := GenS5{F0: v1, F1: 809, F2: true}
                            _ = v49
                        }
                        var v50 uint8 = 0
                        _ = v50
                    } else {
                        for fi51 := 0; fi51 < 3; fi51++ {
                            switch ((("" + "a") + "\n\t") + "a") {
                            case "a":
                                v52 := GenS4{F0: v4, F1: -2147483648, F2: 164}
                                _ = v52
                            }
                            v10 = float64(3.14159)
                            _ = fi51
                        }
                    }
                case "hello":
                    var v53 uint16 = v4
                    _ = v53
                case "":
                    switch "foo bar" {
                    case "foo bar":
                        for fi54 := 0; fi54 < 5; fi54++ {
                            var v55 float32 = v5
                            if false {
                                if v1 {
                                    var v56 bool = (42 != 702)
                                    _ = v56
                                } else {
                                    var v57 int32 = -2147483648
                                    fmt.Println(v57, v10)
                                    _ = v57
                                }
                                var v58 int64 = -685
                                _ = v58
                            } else {
                                _ = len(v2)
                            }
                            v59 := GenS6{F0: (v15 / ((v15) | 1)), F1: 978, F2: 127}
                            _ = fi54
                            _ = v55
                            _ = v59
                        }
                    default:
                        v60 := v3[4294967295]
                        _ = v60
                    }
                }
                fmt.Println(fi14, v5, v1)
                _ = v15
            } else {
                for fi61 := 0; fi61 < 4; fi61++ {
                    v10 += v10
                    v62 := GenS2{F0: 0, F1: ((v4 & v4) >= 65535), F2: -1}
                    _ = fi61
                    _ = v62
                }
                fmt.Println(v1)
            }
            _ = fi14
        }
    case "foo bar":
        switch "foo bar" {
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v63 bool = false
            _ = v63
        }
    case "":
        fmt.Println(v1, v4)
    }
    delete(v3, 407)
    v1 = v1
    if (0 > -9223372036854775808) {
        var v64 uint = 1
        if v1 {
            v2 = append(v2, "hello")
            v2 = append(v2, ("foo bar" + (("hello" + "") + "\n\t")))
            v65 := GenS7{F0: -864}
            _ = v65
        } else {
            fmt.Println(v5, v1)
            v4 = (v4 ^ v4)
            fmt.Println(v64, v64, v64)
        }
        _ = v64
    }
    v5 = v5
    v2 = append(v2, "")
    switch -113 {
    case 32767:
        v10 = 0.0
    case -9223372036854775808:
        for fi66 := 0; fi66 < 1; fi66++ {
            v4 -= (v4 + 65535)
            for fi67 := 0; fi67 < 17; fi67++ {
                if v1 {
                    fmt.Println(v4)
                    var v68 float64 = v10
                    fmt.Println(v10, v1, v68)
                    _ = v68
                }
                if v1 {
                    v10 -= (1e10 / v10)
                    if v1 {
                        var v69 int8 = 12
                        delete(v3, 42)
                        v4 -= 0
                        _ = v69
                    }
                }
                _ = fi67
            }
            _ = fi66
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
