package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int; F2 string }
type GenS5 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool; F1 int32 }
type GenS10 struct { F0 int32; F1 int }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 string; F1 int16; F2 uint16 }
type GenS13 struct { F0 int32; F1 string; F2 bool }
type GenS14 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint) uint32 {
    v6 := GenS1{F0: 1000, F1: 1, F2: 42}
    v7 := GenS1{F0: 1000, F1: -745, F2: 0}
    v8 := []int{2147483647}
    _ = len(v8)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int32) uint16 {
    v9 := GenS1{F0: 1, F1: 42, F2: p0}
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := []int{(fi10 >> ((841) & 63))}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v9
    return 1
}

func main() {
    var v1 float64 = 3.14159
    v2 := []int{319, 32767}
    var v3 int16 = -32768
    var v4 float64 = v1
    var v5 int32 = 42
    fmt.Println(v1)
    v12 := GenS2{F0: (v3 / ((1) | 1))}
    var v13 uint32 = 2
    if len(v2) > 0 {
        v2[0] = -272
    }
    fmt.Println(v4, v1, v5)
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    if len(v2) > 0 {
        v2[0] = 35
    }
    if true {
        if false {
            var v14 uint64 = 42
            _ = v14
        } else {
            if len(v2) > 0 {
                v2[0] = int(v13)
            }
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 int = (-(fi15))
                if false {
                    var v17 uint64 = 18446744073709551615
                    for fi18 := 0; fi18 < 5; fi18++ {
                        v19 := GenS3{F0: (v5 - (^(v5))), F1: 255}
                        v3 += (v3 | (v3 | (^(v3))))
                        v1 += 3.14159
                        _ = fi18
                        _ = v19
                    }
                    switch (v16 & (v16 * v16)) {
                    case 9223372036854775807:
                        for fi20 := 0; fi20 < 9; fi20++ {
                            switch "" {
                            case "\n\t":
                                v16 = 32768
                            default:
                                var v21 bool = (v3 < (v3 + 1000))
                                _ = v21
                            }
                            var v22 int32 = 2147483647
                            _ = fi20
                            _ = v22
                        }
                    case -1:
                        _ = len(v2)
                    }
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
        }
        v23 := GenS4{F0: (v13 & (v13 / ((59) | 1))), F1: -603, F2: "foo bar"}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 int16 = (v3 ^ v3)
            if (419 == 2147483647) {
                var v26 int32 = v5
                _ = v26
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v27 := GenS4{F0: v13, F1: (fi24 | fi24), F2: ""}
                var v28 int8 = 127
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 14; fi29++ {
                if false {
                    fmt.Println(fi29)
                    if false {
                        v3 = (v3 | v3)
                        var v30 int16 = (^(v3))
                        _ = v30
                    } else {
                        if false {
                            v5 -= -2147483648
                            v31 := GenS5{F0: uint64(v13), F1: 127, F2: 112}
                            _ = v31
                        } else {
                            var v32 float64 = (float64(-2.25) + -2.25)
                            if false {
                                var v33 uint = 893
                                switch fi29 {
                                case 1:
                                    fmt.Println(v1, fi29, v1)
                                }
                                _ = v33
                            } else {
                                switch "a" {
                                case "hello":
                                    var v34 uint = 0
                                    _ = v34
                                }
                                var v35 int = 42
                                if false {
                                    v36 := GenS6{F0: v13}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v36
                                }
                                _ = v35
                            }
                            _ = v32
                        }
                        if true {
                            v13 = genFn0(uint(v13))
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if false {
                                v37 := GenS7{F0: 42, F1: v13, F2: 128}
                                v38 := GenS8{F0: 183}
                                var v39 uint32 = (v13 * 2147483648)
                                _ = v37
                                _ = v38
                                _ = v39
                            } else {
                                switch "\n\t" {
                                case "":
                                    switch ("foo bar" + ("hello" + "hello")) {
                                    case "foo bar":
                                        if ((v5 + (v5 % ((v5) | 1))) == -1) {
                                            _ = len(v2)
                                            if (255 > 1) {
                                                _ = len(v2)
                                                var v40 uint32 = (v13 & 4294967295)
                                                fmt.Println(v25)
                                                _ = v40
                                            }
                                            fmt.Println(fi24, v1)
                                            if (!true) {
                                                _ = len(v2)
                                                switch (fi24 / (((fi24 + fi24)) | 1)) {
                                                case 128:
                                                    if len(v2) > 0 {
                                                        v2[0] = fi29
                                                    }
                                                case 127:
                                                    switch fi29 {
                                                    case 237:
                                                        v1 = v4
                                                    case -2147483648:
                                                        var v41 float32 = (0.0 * (1e10 + (((((1e10 - ((1e10 + ((0.1 - ((((1e10 - (1.5 - -2.25)) / (float32(v1) * (0.0 * 1.5))) + 0.1) * 1.5)) * 3.14159)) - (float32(v4) - 0.0))) * 1e10) * -2.25) - -2.25) / 1e10)))
                                                        if false {
                                                            var v42 int64 = 598
                                                            switch (("hello" + ("\n\t" + ("" + "hello"))) + ("foo bar" + "hello")) {
                                                            case "a":
                                                                fmt.Println(v4)
                                                            default:
                                                                v1 = 0.0
                                                            }
                                                            _ = v42
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = fi24
                                                        }
                                                        _ = v41
                                                    }
                                                default:
                                                    if (255 > genFn1(2147483647)) {
                                                        v1 = v4
                                                    }
                                                    v1 = v1
                                                    switch (fi29 | fi29) {
                                                    case 128:
                                                        v43 := GenS9{F0: true, F1: (-(v5))}
                                                        _ = v43
                                                    case 127:
                                                        v25 -= (v3 - v3)
                                                    case -49:
                                                        var v44 bool = (true || (true && true))
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            if (149 != 255) {
                                                                if (0 >= uint8(v13)) {
                                                                    var v45 int64 = 50
                                                                    _ = v45
                                                                }
                                                                var v46 string = (("\n\t" + "\n\t") + "\n\t")
                                                                if v44 {
                                                                    var v47 int = fi29
                                                                    var v48 int64 = -9223372036854775808
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 888
                                                                    }
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                                _ = v46
                                                            }
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = fi24
                                                        }
                                                        _ = v44
                                                    }
                                                }
                                                var v49 bool = (!false)
                                                _ = v49
                                            }
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = fi29
                                            }
                                            v5 = v5
                                            if (!(true && ((v3 % (((v3 >> ((uint(v13)) & 63))) | 1)) < v25))) {
                                                var v50 string = ("a" + "foo bar")
                                                _ = v50
                                            } else {
                                                v3 = v3
                                                fmt.Println(v3)
                                            }
                                        }
                                    }
                                default:
                                    var v51 uint8 = 223
                                    _ = v51
                                }
                                v52 := GenS10{F0: (-(v5)), F1: (fi29 / (((fi24 >> ((1) & 63))) | 1))}
                                if (false || true) {
                                    _ = len(v2)
                                    _ = len(v2)
                                    v53 := GenS1{F0: v3, F1: v5, F2: (v5 % ((v5) | 1))}
                                    v54 := GenS10{F0: (v5 - -1), F1: (^(fi24))}
                                    v2 = append(v2, fi24)
                                    v55 := GenS11{F0: 1000}
                                    v2 = append(v2, int(v25))
                                    var v56 int8 = 0
                                    var v57 bool = ((255 == 1) || false)
                                    v58 := GenS8{F0: 199}
                                    fmt.Println(fi24, fi29, v57)
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                }
                                _ = v52
                            }
                            if true {
                                fmt.Println(v1)
                                v2 = append(v2, fi24)
                                fmt.Println(v3)
                            } else {
                                var v59 uint = 2
                                v60 := GenS4{F0: (^(v13)), F1: -9223372036854775808, F2: ("foo bar" + "")}
                                _ = v59
                                _ = v60
                            }
                            switch (fi29 * fi24) {
                            case 2147483647:
                                if ((false || (0 < (v5 % (((v5 * (v5 | 2147483647))) | 1)))) || true) {
                                    if true {
                                        var v61 uint32 = genFn0(255)
                                        if (float32(v4) >= float32((1.5 / (v1 + v1)))) {
                                            v62 := GenS1{F0: (v3 & (v25 * v25)), F1: (v5 & (v5 / ((v5) | 1))), F2: v5}
                                            _ = v62
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = (fi29 << ((2) & 63))
                                            }
                                            v25 += v3
                                            if (1 == 42) {
                                                v63 := GenS2{F0: (v3 % (((v25 + (v25 + v3))) | 1))}
                                                if (!false) {
                                                    fmt.Println(v1, v1)
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v63
                                            }
                                        }
                                        var v64 uint8 = 42
                                        _ = v61
                                        _ = v64
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        var v65 uint16 = 735
                                        _ = v65
                                    }
                                    var v66 uint16 = 42
                                    _ = v66
                                }
                            case 4294967295:
                                var v67 uint8 = 2
                                _ = v67
                            }
                        }
                        var v68 int16 = v25
                        _ = v68
                    }
                    switch fi29 {
                    case 32767:
                        fmt.Println(v3, v1)
                    case 9223372036854775807:
                        if (!true) {
                            v69 := GenS12{F0: ("" + ("foo bar" + ("" + "foo bar"))), F1: int16(fi24), F2: 2}
                            _ = v69
                        } else {
                            if false {
                                if (1 == 458) {
                                    switch fi24 {
                                    case 127:
                                        v5 = (v5 / ((v5) | 1))
                                    default:
                                        var v70 uint64 = 0
                                        _ = v70
                                    }
                                } else {
                                    switch 40 {
                                    case 964:
                                        var v71 bool = (-1 > 1)
                                        _ = v71
                                    default:
                                        v13 += (v13 << ((4294967295) & 63))
                                    }
                                    if (1 >= 400) {
                                        var v72 int64 = -182
                                        _ = v72
                                    } else {
                                        if (true || (int16(fi29) <= 1)) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            switch "foo bar" {
                                            case "foo bar":
                                                v13 = (v13 / (((v13 << ((0) & 63))) | 1))
                                            case "\n\t":
                                                if (true || (((v25 >= 0) && false) && (((128 >= fi29) || (genFn1(v5) == 363)) || (42 > fi24)))) {
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi24 & (fi24 >> ((65535) & 63)))
                                                    }
                                                } else {
                                                    fmt.Println(v25)
                                                    v73 := GenS13{F0: (v5 - (^(v5))), F1: ("" + "a"), F2: false}
                                                    _ = v73
                                                }
                                            default:
                                                if (40 >= 255) {
                                                    v5 += v5
                                                    fmt.Println(v25, v5)
                                                }
                                            }
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v5 = 0
                                            fmt.Println(v3, fi29)
                                            v74 := GenS11{F0: v3}
                                            v1 = 1.5
                                            v75 := GenS14{F0: 1, F1: v25}
                                            _ = v74
                                            _ = v75
                                        }
                                        fmt.Println(v1, fi24)
                                        v2 = append(v2, (fi24 >> ((65535) & 63)))
                                    }
                                    if ((!((fi24 - (fi24 ^ (fi29 >> ((uint(v13)) & 63)))) <= (fi24 & 4294967295))) || true) {
                                        v76 := GenS4{F0: v13, F1: fi29, F2: ("a" + "\n\t")}
                                        _ = v76
                                    } else {
                                        v2 = append(v2, fi24)
                                        if (1 >= 65535) {
                                            var v77 string = "foo bar"
                                            v78 := GenS8{F0: 2}
                                            switch "hello" {
                                            case "\n\t":
                                                var v79 int = (fi29 & 127)
                                                _ = v79
                                            case "a":
                                                v3 = (^(v25))
                                            }
                                            _ = v77
                                            _ = v78
                                        }
                                    }
                                    v3 = 1
                                }
                                v80 := GenS9{F0: false, F1: 0}
                                v25 += (v25 ^ 0)
                                _ = v80
                            } else {
                                var v81 uint64 = 925
                                _ = v81
                            }
                            fmt.Println(v4)
                            v5 += v5
                        }
                    case 4294967295:
                        switch ("" + "foo bar") {
                        case "hello":
                            v25 = -678
                            v25 = 332
                            v82 := GenS2{F0: v3}
                            _ = v82
                        }
                    }
                    v13 = genFn0(0)
                    v3 = (v25 + (v3 / (((-(v3))) | 1)))
                    if false {
                        v83 := GenS4{F0: genFn0(146), F1: (fi24 >> ((1) & 63)), F2: "foo bar"}
                        _ = v83
                    } else {
                        var v84 int8 = 42
                        v13 += v13
                        _ = v84
                    }
                    var v85 uint64 = 1
                    var v86 uint32 = (v13 * 546)
                    _ = v85
                    _ = v86
                } else {
                    if true {
                        v87 := GenS9{F0: true, F1: (v5 & (^(v5)))}
                        v5 += 1
                        v13 += v13
                        _ = v87
                    } else {
                        v5 -= (v5 | (-(v5)))
                        fmt.Println(v5, fi24)
                    }
                    v2 = append(v2, fi29)
                }
                var v88 float64 = ((v4 * (v1 * float64(((0.1 + 0.0) - (((3.14159 - float32(3.14159)) - float32(v4)) - 3.14159))))) - v1)
                var v89 uint8 = 1
                _ = fi29
                _ = v88
                _ = v89
            }
            _ = fi24
            _ = v25
        }
        if (1 >= 0) {
            var v90 int32 = v5
            _ = v90
        }
        if ((!(-180 != int64(v5))) || (!(-207 == -9223372036854775808))) {
            fmt.Println(v1, v1)
        } else {
            var v91 uint16 = 954
            _ = v91
        }
    }
    switch int(v3) {
    case -1:
        if len(v2) > 0 {
            v2[0] = 1
        }
    case 0:
        switch 1 {
        case -9223372036854775808:
            if len(v2) > 0 {
                v2[0] = -963
            }
        }
    default:
        for fi92 := 0; fi92 < 18; fi92++ {
            var v93 int = (fi92 / ((2147483647) | 1))
            v4 -= 1e10
            switch "a" {
            case "\n\t":
                v4 -= (v1 / v1)
            case "a":
                var v94 string = "hello"
                _ = v94
            }
            _ = fi92
            _ = v93
        }
    }
    var v95 int = 324
    v95 = (v95 + v95)
    v2 = append(v2, v95)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v95
}
