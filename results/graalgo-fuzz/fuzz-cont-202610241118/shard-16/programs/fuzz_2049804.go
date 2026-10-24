package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int; F2 int32 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint }
type GenS5 struct { F0 int; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8; F1 int64; F2 int8 }

func main() {
    var v1 uint32 = 966
    var v2 uint8 = 0
    var v3 float64 = 0.0
    var v4 int64 = 0
    var v5 uint16 = 750
    var v6 uint64 = 373
    if true {
        if true {
            var v7 string = "a"
            v2 = (^(v2))
            v6 = uint64(v2)
            _ = v7
        }
        fmt.Println(v4, v3)
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint16 = v5
            var v10 int8 = int8(v1)
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    switch 42 {
    case 128:
        v11 := []int{-2147483648, 9223372036854775807, 1, 127}
        _ = v11
    case 32768:
        fmt.Println(v6)
    case -2147483648:
        v12 := []int{32768, -1, 65}
        _ = v12
    }
    if true {
        var v13 int = -115
        v14 := []int{}
        _ = v13
        _ = v14
    } else {
        v15 := GenS1{F0: 127}
        v2 = 119
        _ = v15
    }
    switch -51 {
    case 141:
        var v16 uint32 = (v1 + (^(v1)))
        _ = v16
    case -9223372036854775808:
        for fi17 := 0; fi17 < 16; fi17++ {
            if (14 > 0) {
                v18 := GenS1{F0: 42}
                var v19 int64 = (v4 - int64(v5))
                _ = v18
                _ = v19
            } else {
                switch "a" {
                case "hello":
                    var v20 uint32 = v1
                    _ = v20
                case "a":
                    var v21 uint8 = v2
                    _ = v21
                default:
                    fmt.Println(v3)
                }
            }
            _ = fi17
        }
    }
    v22 := []int{-179, 42}
    if len(v22) > 0 {
        v22[0] = int(v1)
    }
    if true {
        if true {
            v23 := GenS2{F0: -659}
            for fi24 := 0; fi24 < 6; fi24++ {
                switch "foo bar" {
                case "foo bar":
                    if true {
                        fmt.Println(v1)
                        v2 = v2
                        fmt.Println(v1, v3, v4)
                    } else {
                        switch "hello" {
                        case "\n\t":
                            if (65535 != uint(v1)) {
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                                switch fi24 {
                                case 715:
                                    var v25 float32 = ((1e10 + 0.1) + -2.25)
                                    _ = v25
                                case 32768:
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                default:
                                    v26 := GenS1{F0: -1}
                                    _ = v26
                                }
                            } else {
                                var v27 int16 = 0
                                switch "a" {
                                case "":
                                    v4 = (v4 + v4)
                                case "a":
                                    v6 += (^(v6))
                                default:
                                    if (!true) {
                                        for fi28 := 0; fi28 < 4; fi28++ {
                                            _ = len(v22)
                                            v29 := GenS3{F0: 0, F1: (fi24 & fi24), F2: 2147483647}
                                            if false {
                                                if (false || false) {
                                                    if (uint8(v1) != (v2 ^ v2)) {
                                                        if true {
                                                            var v30 int8 = 1
                                                            var v31 int64 = 9223372036854775807
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        fmt.Println(v3)
                                                    }
                                                    v32 := GenS1{F0: 127}
                                                    _ = v32
                                                }
                                                fmt.Println(v27, v3)
                                                v4 = v4
                                            } else {
                                                v33 := GenS4{F0: v2, F1: (^(v27)), F2: 974}
                                                if true {
                                                    _ = len(v22)
                                                } else {
                                                    _ = len(v22)
                                                }
                                                var v34 int8 = 42
                                                _ = v33
                                                _ = v34
                                            }
                                            _ = fi28
                                            _ = v29
                                        }
                                        for fi35 := 0; fi35 < 17; fi35++ {
                                            var v36 string = "a"
                                            _ = fi35
                                            _ = v36
                                        }
                                    } else {
                                        fmt.Println(v5)
                                    }
                                }
                                var v37 int8 = -1
                                _ = v27
                                _ = v37
                            }
                        case "":
                            if (-71 == 23) {
                                v4 = -9223372036854775808
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            } else {
                                for fi38 := 0; fi38 < 1; fi38++ {
                                    if ((0.1 * 0.1) < 1e10) {
                                        if (!(true && (1000 <= -135))) {
                                            var v39 float32 = (1e10 * -2.25)
                                            fmt.Println(fi38)
                                            if len(v22) >= 2 {
                                                v22 = v22[1:len(v22)]
                                            }
                                            _ = v39
                                        } else {
                                            var v40 uint = 0
                                            _ = v40
                                        }
                                    }
                                    v41 := GenS3{F0: (v6 % ((v6) | 1)), F1: 32767, F2: -2147483648}
                                    if len(v22) > 0 {
                                        v22[0] = (fi38 % ((45) | 1))
                                    }
                                    _ = fi38
                                    _ = v41
                                }
                                for fi42 := 0; fi42 < 2; fi42++ {
                                    if true {
                                        var v43 uint16 = 1
                                        v5 -= (v43 | v43)
                                        if len(v22) > 0 {
                                            v22[0] = fi42
                                        }
                                        if (1000 >= 32767) {
                                            var v44 int16 = -32768
                                            v22 = append(v22, int(v6))
                                            if false {
                                                var v45 int32 = -721
                                                _ = v45
                                            } else {
                                                switch fi24 {
                                                case 127:
                                                    fmt.Println(v43, v44, v5)
                                                default:
                                                    if true {
                                                        if len(v22) >= 2 {
                                                            v22 = v22[1:len(v22)]
                                                        }
                                                        switch ("foo bar" + "") {
                                                        case "foo bar":
                                                            var v46 int16 = (v44 - v44)
                                                            _ = v46
                                                        case "hello":
                                                            var v47 int32 = 42
                                                            _ = v47
                                                        }
                                                    }
                                                    var v48 float32 = 1e10
                                                    _ = v48
                                                }
                                                v1 = (v1 * (v1 ^ (v1 << ((165) & 63))))
                                            }
                                            if (int32(v6) >= -2147483648) {
                                                v49 := GenS4{F0: 42, F1: v44, F2: 0}
                                                switch "foo bar" {
                                                case "a":
                                                    v4 += v4
                                                    v44 -= (v44 % ((v44) | 1))
                                                    v43 -= (v5 / ((0) | 1))
                                                case "foo bar":
                                                    var v50 int64 = (v4 | v4)
                                                    _ = v50
                                                case "hello":
                                                    v2 += v2
                                                default:
                                                    fmt.Println(v6, v6)
                                                    if true {
                                                        fmt.Println(v5, v5, v5)
                                                    }
                                                    v43 = v5
                                                    if false {
                                                        _ = len(v22)
                                                    }
                                                    v51 := GenS2{F0: (v44 ^ v44)}
                                                    _ = v51
                                                }
                                                _ = v49
                                            }
                                            _ = v44
                                        }
                                        if true {
                                            fmt.Println(fi42, v2, v3)
                                        }
                                        _ = v43
                                    } else {
                                        _ = len(v22)
                                        if false {
                                            fmt.Println(fi42)
                                        } else {
                                            var v52 int = int(v1)
                                            switch (((("hello" + "hello") + (("" + "foo bar") + ("" + ("\n\t" + "hello")))) + "\n\t") + (("\n\t" + ("hello" + "")) + ("foo bar" + "\n\t"))) {
                                            case "hello":
                                                v22 = append(v22, -2147483648)
                                            case "foo bar":
                                                var v53 uint32 = (v1 & 205)
                                                _ = v53
                                            case "a":
                                                var v54 uint = 255
                                                _ = v54
                                            }
                                            switch ("hello" + ("a" + "")) {
                                            case "foo bar":
                                                var v55 string = "a"
                                                _ = v55
                                            case "":
                                                v56 := GenS1{F0: 0}
                                                _ = v56
                                            default:
                                                v6 += (v6 & (v6 * (v6 / ((169) | 1))))
                                            }
                                            v57 := GenS3{F0: v6, F1: (fi24 * 32767), F2: 0}
                                            _ = v52
                                            _ = v57
                                        }
                                    }
                                    _ = fi42
                                }
                            }
                        default:
                            v5 -= 32767
                        }
                        for fi58 := 0; fi58 < 8; fi58++ {
                            switch "foo bar" {
                            case "a":
                                var v59 int8 = int8(fi24)
                                _ = v59
                            }
                            var v60 int8 = 100
                            v22 = append(v22, fi58)
                            _ = fi58
                            _ = v60
                        }
                    }
                case "":
                    for fi61 := 0; fi61 < 3; fi61++ {
                        switch int(v1) {
                        case -1:
                            if (-1 < 0) {
                                var v62 int64 = (v4 % ((v4) | 1))
                                switch fi24 {
                                case -296:
                                    switch int(v5) {
                                    case 4294967295:
                                        var v63 int32 = -1
                                        _ = v63
                                    case 559:
                                        if false {
                                            fmt.Println(fi61)
                                        } else {
                                            switch fi61 {
                                            case 32767:
                                                v5 = (v5 + v5)
                                            case 42:
                                                switch (fi24 % ((fi24) | 1)) {
                                                case 9223372036854775807:
                                                    if len(v22) >= 2 {
                                                        v22 = v22[1:len(v22)]
                                                    }
                                                case 2147483647:
                                                    var v64 float32 = 0.0
                                                    _ = v64
                                                case 128:
                                                    if len(v22) > 0 {
                                                        v22[0] = 4294967295
                                                    }
                                                }
                                            default:
                                                var v65 string = (("foo bar" + ("" + "")) + "")
                                                _ = v65
                                            }
                                        }
                                    case -9223372036854775808:
                                        v66 := GenS4{F0: (^(v2)), F1: 1000, F2: 4294967295}
                                        _ = v66
                                    default:
                                        v4 += 9223372036854775807
                                    }
                                }
                                _ = v62
                            }
                        case 32768:
                            fmt.Println(fi61)
                        }
                        _ = fi61
                    }
                case "\n\t":
                    fmt.Println(v6, v2, fi24)
                }
                switch "" {
                case "a":
                    var v67 string = "a"
                    _ = v67
                case "hello":
                    v5 -= v5
                case "":
                    if true {
                        fmt.Println(v5)
                        var v68 float64 = ((1.5 * v3) * v3)
                        var v69 uint = 677
                        _ = v68
                        _ = v69
                    }
                default:
                    v1 = 65535
                }
                fmt.Println(v3, v5)
                _ = fi24
            }
            if true {
                v6 -= 900
                v6 = (v6 >> ((42) & 63))
            } else {
                fmt.Println(v5, v5, v2)
                var v70 uint32 = 4294967295
                v71 := GenS5{F0: 291, F1: (^(v2)), F2: uint32(v6)}
                _ = v70
                _ = v71
            }
            _ = v23
        }
        if true {
            if (v1 <= 785) {
                var v72 string = "hello"
                _ = len(v22)
                var v73 int8 = 39
                _ = v72
                _ = v73
            } else {
                var v74 string = ((("" + ("a" + "\n\t")) + "a") + ((("hello" + ("\n\t" + "foo bar")) + "\n\t") + "foo bar"))
                if len(v22) > 0 {
                    v22[0] = 9223372036854775807
                }
                _ = v74
            }
            var v75 uint64 = (v6 >> ((978) & 63))
            switch -9223372036854775808 {
            case 32768:
                switch -735 {
                case -726:
                    v76 := GenS6{F0: uint(v1), F1: uint8(v6), F2: v5}
                    _ = v76
                default:
                    fmt.Println(v5, v6)
                }
            case 9223372036854775807:
                v3 += (-2.25 + 1.5)
            default:
                if len(v22) > 0 {
                    v22[0] = 42
                }
            }
            _ = v75
        } else {
            var v77 int64 = v4
            _ = v77
        }
    } else {
        if false {
            v2 += uint8(v1)
        }
        for fi78 := 0; fi78 < 11; fi78++ {
            for fi79 := 0; fi79 < 1; fi79++ {
                v6 -= v6
                if (!true) {
                    v80 := GenS7{F0: (^(v6))}
                    _ = v80
                } else {
                    v3 = (v3 - float64(3.14159))
                }
                fmt.Println(v5)
                v81 := GenS8{F0: uint8(v5), F1: (v4 / ((0) | 1)), F2: -128}
                v1 -= v1
                v3 = v3
                _ = fi79
                _ = v81
            }
            _ = fi78
        }
        var v82 bool = true
        _ = v82
    }
    var v83 int64 = (^(v4))
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v83
}
