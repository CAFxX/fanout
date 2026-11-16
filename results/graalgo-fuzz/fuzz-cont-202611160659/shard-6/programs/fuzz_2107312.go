package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int16 }
type GenS2 struct { F0 int; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 int8; F2 int64 }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 string; F1 uint16; F2 uint32 }
type GenS10 struct { F0 int; F1 int64; F2 uint }
type GenS11 struct { F0 uint16; F1 int8; F2 int16 }
type GenS12 struct { F0 int16; F1 string; F2 bool }
type GenS13 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 int8 = -1
    var v2 uint32 = 4294967295
    v3 := []uint32{v2, (v2 | v2), (^(v2)), (v2 * (v2 | v2))}
    v4 := []string{"", "foo bar"}
    if (((1e10 - 0.1) + 0.0) < (0.0 * 0.0)) {
        v1 = (v1 | 0)
        fmt.Println(v1, v2)
    } else {
        fmt.Println(v1)
    }
    v1 -= 42
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 int32 = -20
        _ = fi5
        _ = v6
    }
    fmt.Println(v1, v2)
    var v7 float64 = ((float64(0.0) * float64(0.0)) / 1e10)
    var v8 int32 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v2)
    for fi9 := 0; fi9 < 10; fi9++ {
        if true {
            v10 := GenS1{F0: v1, F1: 39, F2: 0}
            v7 += ((float64((1.5 + -2.25)) + v7) - 1.5)
            var v11 bool = (int(v8) > fi9)
            _ = v10
            _ = v11
        } else {
            var v12 float64 = float64((((((0.0 * float32(1.5)) - (float32(float64(0.1)) * (float32(v7) * ((((float32(((0.0 + -2.25) - v7)) - 3.14159) + (1.5 * ((1.5 / (float32((v7 / 3.14159)) / 0.1)) + (1.5 - 1e10)))) + 0.1) + 3.14159)))) - (((0.1 * (-2.25 + 0.0)) / 1e10) / (-2.25 / 0.1))) + (0.1 * 3.14159)) * -2.25))
            var v13 uint16 = 65535
            _ = v12
            _ = v13
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v14 := GenS1{F0: (v1 + 49), F1: (-(v1)), F2: 32767}
        _ = fi9
        _ = v14
    }
    v8 += 42
    if true {
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 uint16 = 65535
            var v17 uint8 = 2
            if (0 < 18446744073709551615) {
                for fi18 := 0; fi18 < 5; fi18++ {
                    fmt.Println(v17, fi18, v7)
                    v1 -= int8(v17)
                    v17 = 96
                    _ = fi18
                }
                if (!(347 == 2)) {
                    v2 = (v2 / (((v2 * v2)) | 1))
                    if ((int8(v17) <= v1) && true) {
                        v19 := GenS1{F0: (v1 ^ v1), F1: 100, F2: 0}
                        v3 = append(v3, 2147483647)
                        v20 := GenS1{F0: (v1 & v1), F1: v1, F2: int16(v8)}
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 3; fi21++ {
                        v22 := GenS1{F0: (v1 ^ (v1 ^ v1)), F1: (-(v1)), F2: -94}
                        _ = fi21
                        _ = v22
                    }
                }
            } else {
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 uint16 = (v16 * v16)
                    if (false && true) {
                        var v25 uint = 0
                        v26 := GenS2{F0: fi15, F1: (v1 << ((v25) & 63)), F2: 4294967295}
                        _ = v25
                        _ = v26
                    } else {
                        v27 := GenS3{F0: 2}
                        _ = v27
                    }
                    v28 := GenS4{F0: int(v1), F1: v1}
                    if ((^(v16)) == v24) {
                        v16 -= (v16 << ((65535) & 63))
                        fmt.Println(v7, fi23, v24)
                        var v29 float64 = v7
                        _ = v29
                    }
                    _ = fi23
                    _ = v24
                    _ = v28
                }
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            v8 = (^(v8))
            v2 -= v2
            var v31 uint64 = 18446744073709551615
            _ = fi30
            _ = v31
        }
        var v32 int8 = v1
        _ = v32
    } else {
        var v33 float32 = 3.14159
        _ = v33
    }
    if false {
        v2 -= (v2 >> ((4294967295) & 63))
    } else {
        v7 -= v7
    }
    v2 -= 0
    for fi34 := 0; fi34 < 1; fi34++ {
        for fi35 := 0; fi35 < 16; fi35++ {
            var v36 bool = true
            if v36 {
                v3 = append(v3, 1)
                var v37 int = fi34
                _ = v37
            } else {
                if v36 {
                    if (0 == v8) {
                        var v38 uint32 = 2147483648
                        switch fi35 {
                        case -992:
                            var v39 bool = v36
                            _ = v39
                        case 895:
                            v40 := GenS2{F0: -77, F1: v1, F2: v2}
                            _ = v40
                        }
                        v41 := GenS5{F0: v1}
                        _ = v38
                        _ = v41
                    }
                } else {
                    var v42 int32 = v8
                    _ = v42
                }
                v43 := GenS4{F0: fi34, F1: (v1 & 100)}
                _ = v43
            }
            _ = fi35
            _ = v36
        }
        if (3.14159 == (1.5 * -2.25)) {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            if (!(true && (!(true && false)))) {
                if (float32((v7 + v7)) <= ((float32((float64(((float32(v7) * 0.0) - ((3.14159 * 1.5) - 0.1))) - v7)) - (3.14159 + 1.5)) * (0.0 * 1e10))) {
                    v2 = v2
                    fmt.Println(v8, fi34)
                }
                if true {
                    v44 := GenS6{F0: 42, F1: (^(v1)), F2: 1}
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v7, fi34, fi34)
                    }
                    _ = v44
                }
                for fi45 := 0; fi45 < 8; fi45++ {
                    v2 = 4294967295
                    v46 := GenS6{F0: v2, F1: v1, F2: 1}
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                    _ = fi45
                    _ = v46
                }
            } else {
                v1 = (v1 >> ((18446744073709551615) & 63))
            }
            if true {
                v2 = (v2 * (v2 ^ v2))
                var v47 int = 127
                for fi48 := 0; fi48 < 5; fi48++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v2 = (v2 | 767)
                    switch v47 {
                    case 127:
                        v1 += (v1 << ((869) & 63))
                    default:
                        if (888 > 2) {
                            fmt.Println(fi48, v47, fi34)
                            v49 := GenS3{F0: (v2 * (v2 << ((65535) & 63)))}
                            v50 := GenS6{F0: v2, F1: v1, F2: 42}
                            _ = v49
                            _ = v50
                        } else {
                            switch "hello" {
                            case "hello":
                                switch (fi48 | (v47 << ((0) & 63))) {
                                case 0:
                                    var v51 int64 = 1
                                    v52 := GenS4{F0: (v47 | (^(fi48))), F1: (v1 - v1)}
                                    _ = v51
                                    _ = v52
                                }
                            case "\n\t":
                                v2 -= v2
                            case "":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                    }
                    v53 := GenS1{F0: v1, F1: (v1 % ((v1) | 1)), F2: 42}
                    _ = fi48
                    _ = v53
                }
                _ = v47
            } else {
                for fi54 := 0; fi54 < 3; fi54++ {
                    v8 += (v8 | v8)
                    v55 := GenS2{F0: (-(fi54)), F1: v1, F2: 2147483647}
                    _ = fi54
                    _ = v55
                }
                v8 += v8
            }
        } else {
            fmt.Println(v1)
            v56 := GenS5{F0: v1}
            v57 := GenS7{F0: 2, F1: 216}
            _ = v56
            _ = v57
        }
        fmt.Println(v8, v7)
        _ = fi34
    }
    for fi58 := 0; fi58 < 17; fi58++ {
        v3 = append(v3, v2)
        var v59 int = fi58
        _ = fi58
        _ = v59
    }
    if ((^(v2)) < (^(v2))) {
        v7 = 1.5
    } else {
        v8 -= v8
        var v60 int32 = v8
        switch "" {
        case "a":
            v8 = 1
        case "\n\t":
            for fi61 := 0; fi61 < 8; fi61++ {
                v7 = (3.14159 + v7)
                v8 = v8
                if (65535 >= 18446744073709551615) {
                    for fi62 := 0; fi62 < 9; fi62++ {
                        v63 := GenS5{F0: (v1 & 127)}
                        v64 := GenS5{F0: (v1 | (v1 >> ((1) & 63)))}
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                    fmt.Println(v7, v1)
                } else {
                    for fi65 := 0; fi65 < 5; fi65++ {
                        switch "" {
                        case "\n\t":
                            switch (-(fi65)) {
                            case -201:
                                v66 := GenS1{F0: (v1 % ((v1) | 1)), F1: 0, F2: -816}
                                _ = v66
                            case 2147483647:
                                v8 = v8
                            }
                        default:
                            if false {
                                switch (fi65 % ((fi61) | 1)) {
                                case -9223372036854775808:
                                    var v67 bool = (0 == 2147483647)
                                    _ = v67
                                default:
                                    switch -126 {
                                    case 0:
                                        switch 654 {
                                        case 372:
                                            var v68 uint32 = (v2 + 0)
                                            _ = v68
                                        case 127:
                                            if false {
                                                var v69 uint16 = 65535
                                                if (int8(fi65) <= int8(v2)) {
                                                    fmt.Println(v60, fi61, v8)
                                                    v1 = (v1 * v1)
                                                    var v70 float64 = (v7 + (-2.25 - ((-2.25 / (float64(1e10) / 1e10)) * v7)))
                                                    _ = v70
                                                }
                                                switch ("foo bar" + ("" + "")) {
                                                case "hello":
                                                    if (!false) {
                                                        var v71 uint64 = 42
                                                        if false {
                                                            v72 := GenS1{F0: v1, F1: (-(v1)), F2: 274}
                                                            _ = v72
                                                        } else {
                                                            if len(v4) > 0 {
                                                                v4[0] = (("a" + "") + ("hello" + ("\n\t" + "")))
                                                            }
                                                            var v73 bool = (1 == 255)
                                                            _ = v73
                                                        }
                                                        v60 = v8
                                                        switch ("a" + "hello") {
                                                        case "\n\t":
                                                            v74 := GenS8{F0: 4294967295, F1: v1}
                                                            _ = v74
                                                        case "hello":
                                                            if true {
                                                                v75 := GenS3{F0: 65535}
                                                                var v76 int64 = 9223372036854775807
                                                                if false {
                                                                    v76 -= v76
                                                                }
                                                                _ = v75
                                                                _ = v76
                                                            } else {
                                                                v60 += (^(v60))
                                                                _ = len(v4)
                                                            }
                                                        }
                                                        _ = v71
                                                    } else {
                                                        v2 = (v2 * (v2 ^ v2))
                                                        v77 := GenS9{F0: "foo bar", F1: 1, F2: (^(v2))}
                                                        var v78 int16 = 97
                                                        switch -44 {
                                                        case -831:
                                                            _ = len(v4)
                                                        case 32767:
                                                            if len(v4) > 0 {
                                                                v4[0] = "foo bar"
                                                            }
                                                        case 128:
                                                            fmt.Println(v2)
                                                        default:
                                                            v69 = (v69 % ((v69) | 1))
                                                        }
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                case "":
                                                    fmt.Println(v8, v69, v2)
                                                case "\n\t":
                                                    switch 9223372036854775807 {
                                                    case -151:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    case -393:
                                                        v79 := GenS2{F0: fi65, F1: v1, F2: 2147483648}
                                                        _ = v79
                                                    default:
                                                        v69 += v69
                                                    }
                                                }
                                                _ = v69
                                            } else {
                                                v80 := GenS10{F0: fi61, F1: 42, F2: 573}
                                                if (false || false) {
                                                    v81 := GenS8{F0: 2, F1: v1}
                                                    if false {
                                                        if len(v4) > 0 {
                                                            v4[0] = "\n\t"
                                                        }
                                                        if false {
                                                            v2 -= v2
                                                            v60 = v60
                                                            var v82 uint = 18446744073709551615
                                                            _ = v82
                                                        } else {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            switch (fi65 % (((fi65 - fi65)) | 1)) {
                                                            case -2147483648:
                                                                v7 = ((-2.25 + v7) / -2.25)
                                                            case -267:
                                                                var v83 uint32 = (^(v2))
                                                                _ = v83
                                                            default:
                                                                v1 += v1
                                                            }
                                                            v84 := GenS11{F0: 42, F1: -128, F2: -32768}
                                                            _ = v84
                                                        }
                                                        v85 := GenS9{F0: "a", F1: 0, F2: v2}
                                                        var v86 uint16 = 2
                                                        _ = v85
                                                        _ = v86
                                                    } else {
                                                        if false {
                                                            var v87 int16 = 32767
                                                            v88 := GenS12{F0: (v87 >> ((975) & 63)), F1: "\n\t", F2: true}
                                                            var v89 int32 = (v8 + int32(v1))
                                                            _ = v87
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                        v3 = append(v3, 553)
                                                    }
                                                    _ = len(v4)
                                                    if (true || (42 <= 18446744073709551615)) {
                                                        _ = len(v3)
                                                    }
                                                    _ = v81
                                                }
                                                _ = v80
                                            }
                                            var v90 uint16 = 42
                                            _ = v90
                                        default:
                                            if (!false) {
                                                var v91 int8 = v1
                                                if true {
                                                    _ = len(v3)
                                                    v92 := GenS13{F0: uint8(v2), F1: 42}
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v93 int8 = 100
                                                        _ = v93
                                                    case "\n\t":
                                                        v2 += v2
                                                    case "":
                                                        fmt.Println(v60)
                                                    }
                                                    _ = v92
                                                } else {
                                                    v7 -= -2.25
                                                    v91 += v91
                                                }
                                                _ = v91
                                            }
                                        }
                                    }
                                }
                                var v94 uint8 = 2
                                v8 -= 42
                                _ = v94
                            }
                        }
                        _ = fi65
                    }
                }
                _ = fi61
            }
        }
        _ = v60
    }
    var v95 uint8 = 1
    v96 := GenS5{F0: (v1 - v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v95
    _ = v96
}
