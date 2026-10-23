package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint32; F1 uint16; F2 string }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 int16 }
func genFn0() int8 {
    if true {
        var v6 float64 = float64(-2.25)
        fmt.Println(v6)
        _ = v6
    } else {
        v7 := []int{-810, -1}
        _ = v7
    }
    var v8 int16 = -32768
    v9 := []int{0, -9}
    switch (("foo bar" + ("a" + ("foo bar" + "a"))) + "\n\t") {
    case "\n\t":
        var v10 uint32 = 2147483648
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 8; fi11++ {
            switch ((("a" + "foo bar") + "") + "foo bar") {
            case "hello":
                v12 := GenS1{F0: 135, F1: 556}
                _ = v12
            default:
                v8 = 1000
            }
            v13 := GenS1{F0: 0, F1: 42}
            fmt.Println(fi11, v8)
            _ = fi11
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        fmt.Println(fi14)
        if (true && (!(42 != 1))) {
            for fi15 := 0; fi15 < 18; fi15++ {
                v16 := GenS1{F0: 550, F1: 4294967295}
                _ = fi15
                _ = v16
            }
            v8 = (^(v8))
            var v17 float64 = ((((-2.25 + (0.1 - 1e10)) + float64((0.0 - 0.0))) * (0.0 / 0.1)) - (-2.25 * ((((3.14159 / 0.1) + (float64((0.0 + float32(1.5))) - ((0.0 - (1e10 - 1.5)) + 3.14159))) + 0.1) - (3.14159 + (float64(0.0) + 0.0)))))
            _ = v17
        } else {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            fmt.Println(fi14)
            if (4294967295 != 999) {
                switch "" {
                case "a":
                    v18 := GenS1{F0: 2147483647, F1: 42}
                    _ = v18
                case "":
                    if ((42 < (fi14 | fi14)) && false) {
                        if false {
                            switch fi14 {
                            case -1:
                                v8 = (^(v8))
                            case 2147483647:
                                if (0.0 >= 1.5) {
                                    v8 += (v8 >> ((18446744073709551615) & 63))
                                    switch "foo bar" {
                                    case "\n\t":
                                        if len(v9) > 0 {
                                            v9[0] = 98
                                        }
                                    case "foo bar":
                                        var v19 int64 = -247
                                        _ = v19
                                    case "hello":
                                        v8 += (v8 % ((v8) | 1))
                                    }
                                } else {
                                    var v20 int64 = -1
                                    _ = v20
                                }
                            }
                            v8 = v8
                        } else {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            v8 = v8
                        }
                        if ((fi14 + -421) > 42) {
                            var v21 int = fi14
                            _ = v21
                        }
                    } else {
                        v9 = append(v9, (fi14 % ((fi14) | 1)))
                        v8 = v8
                        var v22 uint16 = 2
                        _ = v22
                    }
                case "foo bar":
                    v9 = append(v9, (fi14 * (fi14 % ((fi14) | 1))))
                }
            } else {
                v9 = append(v9, (-(fi14)))
                var v23 string = "\n\t"
                if len(v9) > 0 {
                    v9[0] = fi14
                }
                _ = v23
            }
        }
        _ = fi14
    }
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := []int{238, -382, -2147483648}
    var v2 int = 128
    var v3 int8 = 127
    var v4 string = ""
    var v5 int = v2
    switch v5 {
    case 4294967295:
        v4 = v4
    case 338:
        var v24 string = v4
        _ = v24
    }
    var v25 uint8 = 1
    _ = len(v1)
    for fi26 := 0; fi26 < 13; fi26++ {
        v25 = v25
        var v27 bool = (128 != v5)
        v1 = append(v1, -513)
        _ = fi26
        _ = v27
    }
    var v28 uint = 892
    var v29 string = ((v4 + (v4 + v4)) + v4)
    fmt.Println(v29, v29, v29)
    var v30 int8 = -1
    fmt.Println(v30)
    var v31 uint = v28
    v32 := GenS1{F0: -2147483648, F1: 483}
    fmt.Println(v31, v29)
    if false {
        for fi33 := 0; fi33 < 1; fi33++ {
            for fi34 := 0; fi34 < 13; fi34++ {
                if false {
                    var v35 int16 = 1000
                    var v36 bool = true
                    _ = v35
                    _ = v36
                } else {
                    v3 = v30
                    var v37 float64 = 0.1
                    var v38 uint = v28
                    _ = v37
                    _ = v38
                }
                switch "\n\t" {
                case "foo bar":
                    v28 += v28
                    v28 -= v28
                case "":
                    v4 = "\n\t"
                }
                var v39 uint = (v31 * v31)
                _ = fi34
                _ = v39
            }
            _ = fi33
        }
        if (((-(v30)) > 1) && ((false && false) && (!true))) {
            v1 = append(v1, v5)
            v40 := GenS1{F0: 355, F1: 999}
            if true {
                for fi41 := 0; fi41 < 3; fi41++ {
                    if (false && false) {
                        v42 := GenS2{F0: (-(v3))}
                        var v43 uint64 = 0
                        fmt.Println(v28, v25)
                        _ = v42
                        _ = v43
                    } else {
                        var v44 uint8 = v25
                        _ = v44
                    }
                    _ = fi41
                }
                v1 = append(v1, v5)
            } else {
                v25 -= uint8(v28)
            }
            _ = v40
        } else {
            if true {
                fmt.Println(v2)
                var v45 int = 32768
                _ = v45
            } else {
                fmt.Println(v31, v31)
                if false {
                    fmt.Println(v2, v25, v5)
                }
                for fi46 := 0; fi46 < 17; fi46++ {
                    var v47 uint = (^(v31))
                    v48 := GenS2{F0: 0}
                    _ = fi46
                    _ = v47
                    _ = v48
                }
            }
            if true {
                switch v29 {
                case "\n\t":
                    var v49 uint = (^(v31))
                    _ = v49
                case "a":
                    v1 = append(v1, (v2 & (v5 & v2)))
                case "":
                    for fi50 := 0; fi50 < 18; fi50++ {
                        if true {
                            if false {
                                for fi51 := 0; fi51 < 12; fi51++ {
                                    v52 := GenS1{F0: 781, F1: 186}
                                    _ = fi51
                                    _ = v52
                                }
                                switch int(v25) {
                                case -1:
                                    switch (v2 + (v5 << (((v31 & 42)) & 63))) {
                                    case 769:
                                        switch 589 {
                                        case 9223372036854775807:
                                            var v53 float64 = (((3.14159 * 0.1) + -2.25) / float64(0.1))
                                            _ = v53
                                        default:
                                            for fi54 := 0; fi54 < 4; fi54++ {
                                                if true {
                                                    if true {
                                                        var v55 int32 = -1
                                                        switch fi54 {
                                                        case 0:
                                                            v25 -= v25
                                                            v30 = genFn0()
                                                            v5 = (fi54 >> (((v28 ^ uint(v25))) & 63))
                                                            var v56 int8 = (v3 % (((v30 ^ (v3 >> ((v31) & 63)))) | 1))
                                                            _ = v56
                                                        case -2147483648:
                                                            switch 256 {
                                                            case 2147483647:
                                                                fmt.Println(v30)
                                                            default:
                                                                switch v4 {
                                                                case "\n\t":
                                                                    v3 -= v3
                                                                case "a":
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        switch (v5 / (((v5 >> (((v28 / ((v28) | 1))) & 63))) | 1)) {
                                                                        case 127:
                                                                            if (!(!(!false))) {
                                                                                v5 = v2
                                                                                var v57 int8 = (^(v30))
                                                                                _ = v57
                                                                            }
                                                                        case -456:
                                                                            _ = len(v1)
                                                                        case 0:
                                                                            fmt.Println(v28, v31)
                                                                            v58 := GenS1{F0: -923, F1: 266}
                                                                            _ = v58
                                                                        }
                                                                    default:
                                                                        var v59 int32 = -2147483648
                                                                        _ = v59
                                                                    }
                                                                    switch fi50 {
                                                                    case 2147483647:
                                                                        v30 += int8(fi50)
                                                                    case 465:
                                                                        v60 := GenS2{F0: int8(v55)}
                                                                        switch v29 {
                                                                        case "":
                                                                            if ((v55 * int32(v25)) >= (v55 ^ 1)) {
                                                                                v4 = (v29 + ("foo bar" + v29))
                                                                                v3 += genFn0()
                                                                                fmt.Println(v29, v30, v5)
                                                                            }
                                                                        case "hello":
                                                                            _ = len(v1)
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = (^(fi50))
                                                                        }
                                                                        var v61 uint32 = 450
                                                                        if (false && false) {
                                                                            var v62 int16 = 0
                                                                            v63 := GenS3{F0: (32767 != v62)}
                                                                            v1 = append(v1, (fi54 << (((^(v28))) & 63)))
                                                                            _ = v62
                                                                            _ = v63
                                                                        }
                                                                        _ = v60
                                                                        _ = v61
                                                                    default:
                                                                        var v64 uint64 = 2
                                                                        _ = v64
                                                                    }
                                                                default:
                                                                    v65 := GenS1{F0: 0, F1: 1}
                                                                    _ = v65
                                                                }
                                                            }
                                                        default:
                                                            var v66 uint = (v28 / (((v31 << ((700) & 63))) | 1))
                                                            _ = v66
                                                        }
                                                        _ = v55
                                                    }
                                                    fmt.Println(v28, v4, v4)
                                                } else {
                                                    if ((!false) && (!true)) {
                                                        v4 = v29
                                                        v4 = v29
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v67 := GenS4{F0: 601}
                                                        _ = v67
                                                    }
                                                    var v68 int = fi54
                                                    _ = v68
                                                }
                                                switch (^(v2)) {
                                                case -2147483648:
                                                    if (3.14159 <= 1e10) {
                                                        v69 := GenS5{F0: -1, F1: 0}
                                                        v70 := GenS2{F0: (v30 * v3)}
                                                        _ = v69
                                                        _ = v70
                                                    } else {
                                                        fmt.Println(v30)
                                                        v5 = v5
                                                        v71 := GenS2{F0: (^(v30))}
                                                        _ = v71
                                                    }
                                                case 4294967295:
                                                    if (0 != 42) {
                                                        fmt.Println(v5)
                                                        if ((v25 & v25) >= v25) {
                                                            v1 = append(v1, (v2 * (-(v2))))
                                                            switch ("" + "foo bar") {
                                                            case "":
                                                                switch ("foo bar" + "") {
                                                                case "\n\t":
                                                                    var v72 int16 = int16(v3)
                                                                    _ = v72
                                                                }
                                                            case "hello":
                                                                var v73 int16 = 42
                                                                _ = v73
                                                            case "foo bar":
                                                                v74 := GenS3{F0: true}
                                                                _ = v74
                                                            }
                                                            if false {
                                                                var v75 uint64 = 1
                                                                if true {
                                                                    if ((v28 & (^(v31))) != v28) {
                                                                        if false {
                                                                            v4 = v29
                                                                            switch (v2 & (fi50 % ((-2147483648) | 1))) {
                                                                            case 32767:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 269
                                                                                }
                                                                            case -207:
                                                                                if true {
                                                                                    v5 = 0
                                                                                    var v76 uint = 42
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                            var v77 uint64 = (v75 + 118)
                                                                            _ = v77
                                                                        } else {
                                                                            if true {
                                                                                v2 = fi50
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = (fi54 - (v2 | (fi50 - -195)))
                                                                                }
                                                                                if (true || (true && (v75 < v75))) {
                                                                                    v1 = append(v1, (v5 - 9223372036854775807))
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = v5
                                                                                    }
                                                                                    var v78 float64 = float64(-2.25)
                                                                                    _ = v78
                                                                                } else {
                                                                                    var v79 int = fi50
                                                                                    _ = v79
                                                                                }
                                                                            } else {
                                                                                if true {
                                                                                    v80 := GenS1{F0: -342, F1: 435}
                                                                                    var v81 uint8 = uint8(v75)
                                                                                    var v82 uint = (^(v31))
                                                                                    _ = v80
                                                                                    _ = v81
                                                                                    _ = v82
                                                                                }
                                                                                v83 := GenS1{F0: 0, F1: 2}
                                                                                _ = v83
                                                                            }
                                                                            v84 := GenS5{F0: 1, F1: 776}
                                                                            v31 = (v31 >> (((v31 >> (((v31 ^ v28)) & 63))) & 63))
                                                                            var v85 int = (fi50 + v2)
                                                                            _ = v84
                                                                            _ = v85
                                                                        }
                                                                    } else {
                                                                        var v86 uint = (v31 | (v28 >> ((v28) & 63)))
                                                                        var v87 bool = true
                                                                        v88 := GenS1{F0: -563, F1: uint64(v28)}
                                                                        _ = v86
                                                                        _ = v87
                                                                        _ = v88
                                                                    }
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (^(fi54))
                                                                    }
                                                                    v3 = 0
                                                                } else {
                                                                    switch (v5 | (^(v2))) {
                                                                    case 128:
                                                                        v5 += 32768
                                                                    case -2147483648:
                                                                        v89 := GenS3{F0: ((v28 * (^(v31))) <= 18446744073709551615)}
                                                                        _ = v89
                                                                    case -793:
                                                                        var v90 int32 = 2147483647
                                                                        _ = v90
                                                                    }
                                                                    var v91 uint16 = uint16(v28)
                                                                    var v92 int32 = -851
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                _ = v75
                                                            } else {
                                                                v25 = v25
                                                            }
                                                        } else {
                                                            v29 = "foo bar"
                                                            v31 -= (^(v31))
                                                            if (false || false) {
                                                                v2 -= (fi54 & (^(fi54)))
                                                                v5 = v2
                                                            } else {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        }
                                                    }
                                                default:
                                                    if false {
                                                        if len(v1) > 0 {
                                                            v1[0] = fi54
                                                        }
                                                        v31 += v28
                                                        if (!(2 > 788)) {
                                                            _ = len(v1)
                                                            var v93 uint8 = 128
                                                            v93 = v93
                                                            _ = v93
                                                        } else {
                                                            if (((1.5 / (1.5 * (1e10 + 1e10))) / 1.5) < -2.25) {
                                                                var v94 uint32 = 2
                                                                if true {
                                                                    var v95 float32 = 3.14159
                                                                    v1 = append(v1, (fi50 | (v2 ^ 32767)))
                                                                    _ = v95
                                                                } else {
                                                                    v2 -= (v2 | -463)
                                                                    var v96 string = (v4 + v29)
                                                                    _ = v96
                                                                }
                                                                _ = v94
                                                            }
                                                            fmt.Println(v3, fi54, v28)
                                                        }
                                                    }
                                                }
                                                v97 := GenS6{F0: 4294967295, F1: 1, F2: "\n\t"}
                                                _ = fi54
                                                _ = v97
                                            }
                                        }
                                    }
                                default:
                                    v98 := GenS7{F0: ("foo bar" + v4), F1: v31}
                                    _ = v98
                                }
                                fmt.Println(v3, fi50, v31)
                            }
                        } else {
                            v30 = v30
                            v99 := GenS8{F0: 0}
                            _ = v99
                        }
                        _ = fi50
                    }
                default:
                    for fi100 := 0; fi100 < 7; fi100++ {
                        switch ("" + (v4 + v29)) {
                        case "":
                            var v101 float64 = ((((0.1 / 1e10) + 0.1) - (3.14159 * 3.14159)) / 1.0)
                            _ = v101
                        }
                        _ = fi100
                    }
                }
            } else {
                v31 = (v31 & (v28 % ((v28) | 1)))
            }
            for fi102 := 0; fi102 < 13; fi102++ {
                v28 = (v31 * (v31 + v28))
                var v103 float64 = 1.5
                _ = fi102
                _ = v103
            }
        }
    } else {
        v29 = v29
        if len(v1) > 0 {
            v1[0] = v5
        }
    }
    switch v2 {
    case 127:
        switch v29 {
        case "":
            v104 := GenS1{F0: 2147483647, F1: 0}
            _ = v104
        default:
            var v105 int = (v5 + v5)
            _ = v105
        }
    case -1:
        v1 = append(v1, v5)
    default:
        v29 = v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
