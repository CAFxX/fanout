package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint8; F2 int }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64; F1 int64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int8; F1 bool }
type GenS12 struct { F0 int32; F1 int16 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int16; F1 bool; F2 int8 }
type GenS15 struct { F0 string; F1 uint64; F2 int64 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 int32; F1 string }
type GenS18 struct { F0 uint16; F1 int32 }
type GenS19 struct { F0 uint16; F1 int8; F2 int }
type GenS20 struct { F0 int64; F1 uint8 }

func main() {
    var v1 uint = 417
    var v2 float32 = 0.1
    var v3 int = 341
    v4 := []uint32{0}
    var v5 int64 = -389
    var v6 uint = 42
    if true {
        switch ("a" + "a") {
        case "foo bar":
            var v7 int64 = (v5 << (((v1 + (v1 | (^(v1))))) & 63))
            _ = v7
        }
        fmt.Println(v6, v6, v6)
    } else {
        v5 -= v5
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        if false {
            fmt.Println(v2, v3, v3)
        } else {
            var v9 float32 = v2
            var v10 int64 = v5
            _ = v10
            _ = v9
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS1{F0: 1}
        if false {
            _ = len(v4)
        } else {
            var v13 uint8 = 255
            var v14 uint64 = 18446744073709551615
            if true {
                for fi15 := 0; fi15 < 15; fi15++ {
                    fmt.Println(v5, v13)
                    _ = len(v4)
                    _ = len(v4)
                    _ = fi15
                }
                switch "\n\t" {
                case "\n\t":
                    for fi16 := 0; fi16 < 19; fi16++ {
                        if (true || ((4 > 3) && true)) {
                            v5 = (^(v5))
                            if ((^(fi11)) != fi16) {
                                if (v14 > 944) {
                                    var v17 int = 127
                                    if (-125 <= -1) {
                                        var v18 uint64 = 42
                                        var v19 uint16 = 255
                                        var v20 uint = (^(v6))
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    } else {
                                        v2 = 1.5
                                        if len(v4) > 0 {
                                            v4[0] = 10
                                        }
                                    }
                                    v1 += (v6 % ((v1) | 1))
                                    if true {
                                        v4 = append(v4, 2147483647)
                                        v21 := GenS2{F0: (fi16 + (v3 / ((fi16) | 1))), F1: v13, F2: (-(v3))}
                                        fmt.Println(v3, v17)
                                        v2 -= (v2 / float32(-2.25))
                                        _ = v21
                                    }
                                    _ = v17
                                }
                                if true {
                                    switch "foo bar" {
                                    case "foo bar":
                                        _ = len(v4)
                                    case "a":
                                        v22 := GenS1{F0: 2147483647}
                                        _ = v22
                                    default:
                                        var v23 float32 = 0.1
                                        _ = v23
                                    }
                                    fmt.Println(v14, fi16, v2)
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        var v24 string = (("" + "foo bar") + ("foo bar" + ("hello" + "hello")))
                                        _ = v24
                                    default:
                                        fmt.Println(v6)
                                    }
                                    var v25 int32 = 1
                                    _ = v25
                                }
                                switch (((("\n\t" + ("\n\t" + "hello")) + "") + ("hello" + "a")) + ("hello" + "")) {
                                case "":
                                    var v26 uint8 = 49
                                    _ = v26
                                }
                            } else {
                                if (-703 >= fi16) {
                                    v5 -= -809
                                    switch -818 {
                                    case -81:
                                        var v27 uint16 = 0
                                        _ = v27
                                    }
                                }
                                switch "\n\t" {
                                case "a":
                                    var v28 int8 = -17
                                    _ = v28
                                case "foo bar":
                                    switch -808 {
                                    case 0:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case 562:
                                        switch "foo bar" {
                                        case "hello":
                                            var v29 uint8 = 0
                                            _ = v29
                                        }
                                    }
                                case "":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(v14)
                                case "":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    if (v6 <= (v1 & (v6 ^ (v6 / (((v6 + (^(v6)))) | 1))))) {
                                        v14 += 394
                                    }
                                }
                                switch "\n\t" {
                                case "a":
                                    switch (v3 + 923) {
                                    case 2147483647:
                                        v3 -= v3
                                    case -990:
                                        v30 := GenS3{F0: (^(v13)), F1: fi16}
                                        _ = v30
                                    case -1:
                                        if (!false) {
                                            if (((!true) || ((v13 <= v13) && false)) || (uint(v13) < 2)) {
                                                v31 := GenS1{F0: -570}
                                                if (42 != v6) {
                                                    if len(v4) > 0 {
                                                        v4[0] = 255
                                                    }
                                                    if ((float32(3.14159) > ((((v2 / (0.0 - ((0.1 * (((v2 / v2) - (v2 / 1e10)) * (3.14159 - (v2 * 1e10)))) - 0.1))) / 3.14159) - (v2 / (v2 / v2))) - v2)) && ("" <= "")) {
                                                        var v32 string = ("" + "\n\t")
                                                        _ = v32
                                                    } else {
                                                        fmt.Println(v3)
                                                        if true {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                if (!true) {
                                                                    switch 2147483647 {
                                                                    case 2147483647:
                                                                        v33 := GenS2{F0: v3, F1: (v13 % ((v13) | 1)), F2: (v3 & int(v14))}
                                                                        _ = v33
                                                                    default:
                                                                        v34 := GenS2{F0: (fi16 | (fi11 | fi16)), F1: v13, F2: 128}
                                                                        _ = v34
                                                                    }
                                                                    v14 = v14
                                                                    v35 := GenS4{F0: v13, F1: 255, F2: -372}
                                                                    _ = v35
                                                                } else {
                                                                    v1 = v1
                                                                    var v36 int32 = 2147483647
                                                                    if (v13 > (v13 * v13)) {
                                                                        v1 = (v1 << (((^(v1))) & 63))
                                                                        var v37 float32 = (0.0 - (v2 * -2.25))
                                                                        if false {
                                                                            var v38 float32 = 1e10
                                                                            _ = v38
                                                                        } else {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = 65535
                                                                            }
                                                                        }
                                                                        _ = v37
                                                                    } else {
                                                                        if false {
                                                                            fmt.Println(v6)
                                                                            if ((v36 % ((v36) | 1)) != (v36 << (((v1 << ((v6) & 63))) & 63))) {
                                                                                fmt.Println(v2, v14)
                                                                                if (v14 == 2) {
                                                                                    v5 -= v5
                                                                                    if (!false) {
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            if (float32(((1.5 + (float64(v2) + 3.14159)) - 1.5)) == (v2 + v2)) {
                                                                                                var v39 float32 = 0.1
                                                                                                _ = v39
                                                                                            }
                                                                                        case "a":
                                                                                            if true {
                                                                                                var v40 string = "hello"
                                                                                                _ = v40
                                                                                            } else {
                                                                                                v41 := GenS5{F0: 42, F1: "hello"}
                                                                                                v42 := GenS1{F0: v36}
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                                _ = v41
                                                                                                _ = v42
                                                                                            }
                                                                                        case "":
                                                                                            if true {
                                                                                                if (9223372036854775807 == v3) {
                                                                                                    v3 -= -2147483648
                                                                                                    var v43 float64 = 1.5
                                                                                                    v5 -= (-(v5))
                                                                                                    _ = v43
                                                                                                } else {
                                                                                                    fmt.Println(v14, fi11, v2)
                                                                                                }
                                                                                            } else {
                                                                                                if (!(!(("\n\t" + "a") <= ("\n\t" + "")))) {
                                                                                                    fmt.Println(v1)
                                                                                                    if ((^(v14)) != v14) {
                                                                                                        v6 = (v1 % (((v6 ^ v6)) | 1))
                                                                                                        if true {
                                                                                                            var v44 string = "hello"
                                                                                                            if (-503 >= fi16) {
                                                                                                                v44 = "hello"
                                                                                                                if len(v4) >= 2 {
                                                                                                                    v4 = v4[1:len(v4)]
                                                                                                                }
                                                                                                                if (float32(1e10) <= v2) {
                                                                                                                    if (false || (("a" == v44) || true)) {
                                                                                                                        v14 -= v14
                                                                                                                    }
                                                                                                                    fmt.Println(v5)
                                                                                                                    v44 = v44
                                                                                                                }
                                                                                                            } else {
                                                                                                                v45 := GenS1{F0: v36}
                                                                                                                var v46 uint64 = (^(v14))
                                                                                                                v4 = append(v4, 2147483648)
                                                                                                                _ = v45
                                                                                                                _ = v46
                                                                                                            }
                                                                                                            v5 = (v5 >> ((v1) & 63))
                                                                                                            _ = v44
                                                                                                        } else {
                                                                                                            fmt.Println(v2, v13, v36)
                                                                                                            var v47 uint8 = (^(v13))
                                                                                                            switch (v3 / ((32768) | 1)) {
                                                                                                            case 128:
                                                                                                                if len(v4) > 0 {
                                                                                                                    v4[0] = 699
                                                                                                                }
                                                                                                            default:
                                                                                                                var v48 uint64 = 42
                                                                                                                v36 += (v36 + (v36 >> (((v6 & 65535)) & 63)))
                                                                                                                _ = v48
                                                                                                            }
                                                                                                            _ = v47
                                                                                                        }
                                                                                                    }
                                                                                                    var v49 int = 0
                                                                                                    _ = v49
                                                                                                } else {
                                                                                                    if (!(1 != 0)) {
                                                                                                        v1 += v1
                                                                                                        v2 -= v2
                                                                                                        if (v14 > (^(v14))) {
                                                                                                            var v50 int16 = 32767
                                                                                                            if (v2 <= (v2 + (1.5 + (1.5 / ((v2 * float32(1.5)) / (((1e10 * (float32(1.5) / v2)) * float32(1e10)) * -2.25)))))) {
                                                                                                                if false {
                                                                                                                    if (((v13 * v13) >= v13) || (0.0 == 1.5)) {
                                                                                                                        v51 := GenS6{F0: (v13 << (((v6 << ((v1) & 63))) & 63)), F1: 822, F2: 633}
                                                                                                                        var v52 bool = (uint32(v1) >= 264)
                                                                                                                        _ = v51
                                                                                                                        _ = v52
                                                                                                                    } else {
                                                                                                                        var v53 bool = (294 <= 558)
                                                                                                                        v53 = v53
                                                                                                                        if (v53 || true) {
                                                                                                                            var v54 int8 = -1
                                                                                                                            if len(v4) >= 2 {
                                                                                                                                v4 = v4[1:len(v4)]
                                                                                                                            }
                                                                                                                            _ = v54
                                                                                                                        }
                                                                                                                        v6 = v1
                                                                                                                        _ = v53
                                                                                                                    }
                                                                                                                    var v55 float32 = (0.0 / v2)
                                                                                                                    _ = v55
                                                                                                                } else {
                                                                                                                    if false {
                                                                                                                        var v56 int16 = (v50 / (((v50 << ((v1) & 63))) | 1))
                                                                                                                        v1 -= (v6 + 275)
                                                                                                                        var v57 int = (v3 / ((fi16) | 1))
                                                                                                                        _ = v56
                                                                                                                        _ = v57
                                                                                                                    }
                                                                                                                    v58 := GenS6{F0: v13, F1: 910, F2: v6}
                                                                                                                    fmt.Println(v36)
                                                                                                                    if len(v4) >= 2 {
                                                                                                                        v4 = v4[1:len(v4)]
                                                                                                                    }
                                                                                                                    v50 = (v50 / ((v50) | 1))
                                                                                                                    _ = v58
                                                                                                                }
                                                                                                                v59 := GenS5{F0: 127, F1: ""}
                                                                                                                v4 = append(v4, 65535)
                                                                                                                _ = v59
                                                                                                            } else {
                                                                                                                v60 := GenS7{F0: -112}
                                                                                                                var v61 uint16 = 0
                                                                                                                _ = v60
                                                                                                                _ = v61
                                                                                                            }
                                                                                                            _ = v50
                                                                                                        } else {
                                                                                                            v62 := GenS2{F0: (fi16 % ((32768) | 1)), F1: v13, F2: (fi16 >> ((1) & 63))}
                                                                                                            _ = v62
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v63 int32 = 1
                                                                                                        v64 := GenS8{F0: (^(v14)), F1: v5}
                                                                                                        switch "" {
                                                                                                        case "\n\t":
                                                                                                            v4 = append(v4, 2147483647)
                                                                                                        case "foo bar":
                                                                                                            v1 = (v1 & 357)
                                                                                                        case "a":
                                                                                                            if false {
                                                                                                                if false {
                                                                                                                    switch fi16 {
                                                                                                                    case 2147483647:
                                                                                                                        var v65 uint64 = 0
                                                                                                                        _ = v65
                                                                                                                    case -2147483648:
                                                                                                                        v36 += (v36 % ((v63) | 1))
                                                                                                                    case 1:
                                                                                                                        v14 -= (v14 | v14)
                                                                                                                    }
                                                                                                                    v14 -= (^(v14))
                                                                                                                    var v66 int8 = 8
                                                                                                                    _ = v66
                                                                                                                } else {
                                                                                                                    if ((v6 | 18446744073709551615) <= (v1 + v1)) {
                                                                                                                        if true {
                                                                                                                            var v67 uint16 = 1
                                                                                                                            if false {
                                                                                                                                fmt.Println(v1, v13)
                                                                                                                                var v68 int = fi11
                                                                                                                                v68 += -192
                                                                                                                                _ = v68
                                                                                                                            } else {
                                                                                                                                fmt.Println(v6, v5)
                                                                                                                            }
                                                                                                                            _ = v67
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        if len(v4) > 0 {
                                                                                                                            v4[0] = 0
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if (94 > v36) {
                                                                                                                        fmt.Println(v36, v3)
                                                                                                                        if len(v4) >= 2 {
                                                                                                                            v4 = v4[1:len(v4)]
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        fmt.Println(v1, v3)
                                                                                                                        if true {
                                                                                                                            fmt.Println(v3)
                                                                                                                            v6 += 4294967295
                                                                                                                            v6 = v6
                                                                                                                        } else {
                                                                                                                            var v69 float64 = 0.0
                                                                                                                            _ = v69
                                                                                                                        }
                                                                                                                        v6 += v1
                                                                                                                    }
                                                                                                                    _ = len(v4)
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v70 float32 = float32(1.5)
                                                                                                                v71 := GenS9{F0: (v13 % ((127) | 1))}
                                                                                                                _ = v70
                                                                                                                _ = v71
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v63
                                                                                                        _ = v64
                                                                                                    }
                                                                                                    switch (-(fi11)) {
                                                                                                    case -2147483648:
                                                                                                        _ = len(v4)
                                                                                                    case 32768:
                                                                                                        if (false && false) {
                                                                                                            var v72 uint32 = 2147483647
                                                                                                            var v73 bool = false
                                                                                                            _ = v72
                                                                                                            _ = v73
                                                                                                        }
                                                                                                    default:
                                                                                                        _ = len(v4)
                                                                                                    }
                                                                                                    var v74 uint16 = 485
                                                                                                    if (true && false) {
                                                                                                        switch "a" {
                                                                                                        case "foo bar":
                                                                                                            v6 = 1
                                                                                                        }
                                                                                                        var v75 bool = (((-4 > 0) || ((v14 << ((v1) & 63)) <= 0)) && (127 != (^(v13))))
                                                                                                        var v76 int32 = 107
                                                                                                        _ = v75
                                                                                                        _ = v76
                                                                                                    } else {
                                                                                                        var v77 uint32 = 4294967295
                                                                                                        if (!true) {
                                                                                                            v13 = v13
                                                                                                        }
                                                                                                        v78 := GenS10{F0: fi11}
                                                                                                        _ = v77
                                                                                                        _ = v78
                                                                                                    }
                                                                                                    _ = v74
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            var v79 bool = ((int64(v36) > (v5 >> ((v6) & 63))) || (!(2 == uint16(v1))))
                                                                                            _ = v79
                                                                                        }
                                                                                        v14 += v14
                                                                                        var v80 uint32 = 382
                                                                                        _ = v80
                                                                                    }
                                                                                    v4 = append(v4, 406)
                                                                                    v2 = v2
                                                                                    v81 := GenS11{F0: 100, F1: (v5 <= v5)}
                                                                                    switch "foo bar" {
                                                                                    case "\n\t":
                                                                                        v82 := GenS12{F0: (v36 + v36), F1: 42}
                                                                                        _ = v82
                                                                                    case "hello":
                                                                                        if false {
                                                                                            var v83 string = ((("hello" + "foo bar") + "hello") + "")
                                                                                            if true {
                                                                                                if true {
                                                                                                    v4 = append(v4, 2147483647)
                                                                                                    if (755 == v6) {
                                                                                                        if (970 > int16(v1)) {
                                                                                                            v84 := GenS13{F0: v3}
                                                                                                            fmt.Println(v14, v6, v36)
                                                                                                            fmt.Println(fi11, v6)
                                                                                                            _ = v84
                                                                                                        }
                                                                                                    } else {
                                                                                                        v3 += 127
                                                                                                        v2 = -2.25
                                                                                                        var v85 uint64 = (v14 << ((65535) & 63))
                                                                                                        _ = v85
                                                                                                    }
                                                                                                }
                                                                                                if (true && (18446744073709551615 == v14)) {
                                                                                                    var v86 int16 = -502
                                                                                                    _ = v86
                                                                                                }
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = 258
                                                                                                }
                                                                                                switch fi11 {
                                                                                                case 127:
                                                                                                    v87 := GenS14{F0: 32767, F1: (2147483648 == 2), F2: 100}
                                                                                                    _ = v87
                                                                                                case 129:
                                                                                                    switch 1 {
                                                                                                    case -2147483648:
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                        var v88 uint16 = 964
                                                                                                        _ = v88
                                                                                                    case 32767:
                                                                                                        v89 := GenS7{F0: -55}
                                                                                                        _ = v89
                                                                                                    case -9223372036854775808:
                                                                                                        v90 := GenS15{F0: ("hello" + "a"), F1: 2, F2: v5}
                                                                                                        _ = v90
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v91 := GenS14{F0: int16(v13), F1: true, F2: 1}
                                                                                            var v92 uint64 = (v14 & 519)
                                                                                            _ = v83
                                                                                            _ = v91
                                                                                            _ = v92
                                                                                        }
                                                                                    case "a":
                                                                                        var v93 float32 = 0.0
                                                                                        _ = v93
                                                                                    default:
                                                                                        fmt.Println(v2, v13)
                                                                                    }
                                                                                    v36 -= 2147483647
                                                                                    _ = v81
                                                                                }
                                                                                if false {
                                                                                    v94 := GenS11{F0: 8, F1: ((!true) && (1.5 >= ((1e10 * 0.0) + 1.5)))}
                                                                                    if (18446744073709551615 < (v1 | 4294967295)) {
                                                                                        v36 = 42
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = 2147483648
                                                                                        }
                                                                                        _ = len(v4)
                                                                                    }
                                                                                    v1 += v6
                                                                                    _ = v94
                                                                                } else {
                                                                                    var v95 uint32 = 392
                                                                                    v36 = (^(v36))
                                                                                    switch "a" {
                                                                                    case "a":
                                                                                        fmt.Println(v14)
                                                                                    case "\n\t":
                                                                                        v4 = append(v4, 375)
                                                                                    default:
                                                                                        v3 -= v3
                                                                                    }
                                                                                    _ = v95
                                                                                }
                                                                            }
                                                                            var v96 uint16 = 65535
                                                                            _ = v96
                                                                        } else {
                                                                            var v97 uint64 = v14
                                                                            _ = v97
                                                                        }
                                                                    }
                                                                    _ = v36
                                                                }
                                                            case "a":
                                                                _ = len(v4)
                                                            case "foo bar":
                                                                var v98 int32 = 212
                                                                _ = v98
                                                            default:
                                                                v5 = (v5 / ((v5) | 1))
                                                            }
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                    switch (v3 >> ((v1) & 63)) {
                                                    case 671:
                                                        var v99 int16 = -1
                                                        _ = v99
                                                    case 162:
                                                        fmt.Println(fi11)
                                                    case 1:
                                                        if (((1.5 + 0.0) / (v2 + v2)) <= float32(((0.1 / 1e10) - (1e10 - 1.5)))) {
                                                            v100 := GenS12{F0: -2147483648, F1: 1}
                                                            if false {
                                                                fmt.Println(v3)
                                                                var v101 uint8 = (v13 % (((v13 & 1)) | 1))
                                                                var v102 string = "a"
                                                                _ = v101
                                                                _ = v102
                                                            } else {
                                                                var v103 uint32 = 0
                                                                var v104 bool = true
                                                                _ = v103
                                                                _ = v104
                                                            }
                                                            _ = v100
                                                        } else {
                                                            fmt.Println(v6, v13, v6)
                                                            var v105 int64 = (v5 * 189)
                                                            v1 += v6
                                                            _ = v105
                                                        }
                                                    }
                                                    v14 -= 18446744073709551615
                                                }
                                                switch "\n\t" {
                                                case "":
                                                    v6 = (v1 >> ((v1) & 63))
                                                case "hello":
                                                    var v106 uint8 = uint8(v6)
                                                    _ = v106
                                                }
                                                _ = v31
                                            }
                                            fmt.Println(fi16)
                                        }
                                    }
                                case "":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "hello":
                                    if true {
                                        v3 += 506
                                    } else {
                                        v2 = v2
                                        switch int(v1) {
                                        case 302:
                                            var v107 int64 = 42
                                            _ = v107
                                        case 4:
                                            fmt.Println(v13)
                                        case -1:
                                            v108 := GenS12{F0: -516, F1: -797}
                                            _ = len(v4)
                                            var v109 bool = (true || false)
                                            _ = v108
                                            _ = v109
                                        default:
                                            if false {
                                                if true {
                                                    switch v3 {
                                                    case 42:
                                                        var v110 bool = (v13 == (v13 | (^(v13))))
                                                        _ = v110
                                                    case 0:
                                                        v2 -= v2
                                                    case 4294967295:
                                                        if (!false) {
                                                            switch "foo bar" {
                                                            case "hello":
                                                                v1 = (v6 - v1)
                                                            case "foo bar":
                                                                v4 = append(v4, 2147483647)
                                                            case "a":
                                                                var v111 int = (fi16 & v3)
                                                                _ = v111
                                                            default:
                                                                switch "" {
                                                                case "a":
                                                                    fmt.Println(v1)
                                                                case "foo bar":
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                case "hello":
                                                                    fmt.Println(v14, v1)
                                                                default:
                                                                    fmt.Println(v6, v6, v13)
                                                                    if true {
                                                                        _ = len(v4)
                                                                        var v112 bool = false
                                                                        _ = v112
                                                                    } else {
                                                                        v14 = v14
                                                                        v113 := GenS14{F0: 113, F1: (!false), F2: 1}
                                                                        var v114 int16 = 0
                                                                        _ = v113
                                                                        _ = v114
                                                                    }
                                                                }
                                                            }
                                                            if false {
                                                                var v115 bool = (!("foo bar" == (("a" + "") + ("a" + "hello"))))
                                                                if true {
                                                                    v13 = (v13 ^ 42)
                                                                    var v116 uint64 = 42
                                                                    var v117 int = v3
                                                                    _ = v116
                                                                    _ = v117
                                                                }
                                                                _ = v115
                                                            } else {
                                                                v14 = v14
                                                            }
                                                        } else {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    default:
                                                        if ("hello" != "foo bar") {
                                                            var v118 bool = (-242 > -896)
                                                            _ = v118
                                                        }
                                                    }
                                                } else {
                                                    if (!(((((((3.14159 * (0.1 * (1e10 - (1e10 * 1.5)))) / (1.5 - 0.0)) / 1.5) + 0.1) + 0.1) - (float64(v2) + (1e10 / -2.25))) != 1e10)) {
                                                        switch "a" {
                                                        case "":
                                                            switch fi16 {
                                                            case 127:
                                                                switch "" {
                                                                case "":
                                                                    var v119 int8 = 100
                                                                    _ = v119
                                                                }
                                                            case 42:
                                                                switch (v3 / (((v3 >> ((v1) & 63))) | 1)) {
                                                                case 445:
                                                                    fmt.Println(v3, v14, v5)
                                                                }
                                                            case -1:
                                                                var v120 float64 = float64(v2)
                                                                _ = v120
                                                            }
                                                        case "\n\t":
                                                            if len(v4) > 0 {
                                                                v4[0] = uint32(v6)
                                                            }
                                                        default:
                                                            v13 = (v13 / ((v13) | 1))
                                                        }
                                                        v2 += (v2 / v2)
                                                        _ = len(v4)
                                                    }
                                                }
                                                v121 := GenS16{F0: 2147483648}
                                                _ = len(v4)
                                                var v122 float32 = (v2 + 0.1)
                                                v5 = v5
                                                var v123 uint32 = 944
                                                _ = v121
                                                _ = v122
                                                _ = v123
                                            } else {
                                                switch ("foo bar" + "\n\t") {
                                                case "foo bar":
                                                    if true {
                                                        fmt.Println(v1, v14)
                                                    }
                                                }
                                            }
                                        }
                                        v13 = 42
                                    }
                                default:
                                    v124 := GenS5{F0: 5, F1: ("\n\t" + "")}
                                    _ = v124
                                }
                            }
                            if false {
                                switch "a" {
                                case "":
                                    v125 := GenS5{F0: 3, F1: "a"}
                                    var v126 int8 = -81
                                    fmt.Println(v3, v1, v1)
                                    fmt.Println(v13, v2)
                                    _ = v125
                                    _ = v126
                                default:
                                    v127 := GenS6{F0: (v13 - (v13 << ((v1) & 63))), F1: 911, F2: 25}
                                    _ = v127
                                }
                                if ((v6 >> ((4294967295) & 63)) >= (^(v1))) {
                                    v2 -= 0.1
                                    v14 += (v14 >> ((255) & 63))
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                } else {
                                    fmt.Println(v2, v5, v1)
                                    switch -427 {
                                    case 122:
                                        if true {
                                            fmt.Println(fi11)
                                            var v128 string = (("" + "") + "\n\t")
                                            _ = v128
                                        } else {
                                            v13 = v13
                                        }
                                    case -108:
                                        if (("hello" + "hello") >= ("a" + "\n\t")) {
                                            if len(v4) > 0 {
                                                v4[0] = 4294967295
                                            }
                                            v6 = (v6 / (((v6 ^ 18446744073709551615)) | 1))
                                            var v129 uint16 = uint16(v14)
                                            _ = v129
                                        } else {
                                            var v130 int16 = -995
                                            fmt.Println(v2, v3, v13)
                                            _ = len(v4)
                                            _ = v130
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(fi16, v5)
                                v14 += v14
                                v131 := GenS14{F0: 513, F1: ((v6 % ((231) | 1)) != (^(v1))), F2: -37}
                                _ = v131
                            }
                        }
                        v132 := GenS14{F0: 32767, F1: (124 >= -591), F2: 127}
                        _ = fi16
                        _ = v132
                    }
                }
            } else {
                for fi133 := 0; fi133 < 16; fi133++ {
                    v134 := GenS17{F0: -2147483648, F1: ("hello" + "hello")}
                    _ = fi133
                    _ = v134
                }
                if true {
                    var v135 bool = ((("hello" + "\n\t") + "hello") != "")
                    v135 = false
                    _ = v135
                }
                fmt.Println(fi11, v13, v2)
            }
            _ = v13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    switch "a" {
    case "\n\t":
        v5 -= (v5 ^ -9223372036854775808)
    default:
        if false {
            var v136 int64 = 1
            _ = len(v4)
            _ = v136
        } else {
            v137 := GenS12{F0: -1, F1: -274}
            var v138 float32 = 3.14159
            _ = v137
            _ = v138
        }
    }
    if false {
        if false {
            v2 = v2
        } else {
            if true {
                v139 := GenS10{F0: (v3 | 32768)}
                _ = v139
            } else {
                if (false && true) {
                    if false {
                        for fi140 := 0; fi140 < 5; fi140++ {
                            if false {
                                v141 := GenS17{F0: -466, F1: ("hello" + "a")}
                                var v142 uint64 = 980
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v141
                                _ = v142
                            } else {
                                var v143 uint16 = 42
                                if len(v4) > 0 {
                                    v4[0] = 42
                                }
                                _ = v143
                            }
                            _ = fi140
                        }
                        v144 := GenS18{F0: 242, F1: 1}
                        _ = v144
                    }
                    v6 += (^(v6))
                    fmt.Println(v6, v6, v1)
                }
            }
        }
        v1 -= v1
        v1 = v1
    } else {
        v145 := GenS19{F0: 255, F1: 100, F2: 0}
        if len(v4) > 0 {
            v4[0] = 220
        }
        _ = v145
    }
    var v146 int64 = -537
    var v147 int16 = 42
    if true {
        v3 += 127
        for fi148 := 0; fi148 < 13; fi148++ {
            v149 := GenS20{F0: -9223372036854775808, F1: 128}
            var v150 uint64 = uint64(v6)
            _ = fi148
            _ = v149
            _ = v150
        }
    } else {
        if len(v4) > 0 {
            v4[0] = 0
        }
        fmt.Println(v6, v6, v6)
        v2 = 1.5
    }
    fmt.Println(v6, v1)
    _ = len(v4)
    var v151 int8 = 42
    for fi152 := 0; fi152 < 20; fi152++ {
        v1 = v1
        var v153 int16 = int16(v146)
        _ = fi152
        _ = v153
    }
    v151 = v151
    fmt.Println("v1:", v1)
    fmt.Println("v146:", v146)
    fmt.Println("v147:", v147)
    fmt.Println("v151:", v151)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v146
    _ = v147
    _ = v151
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
