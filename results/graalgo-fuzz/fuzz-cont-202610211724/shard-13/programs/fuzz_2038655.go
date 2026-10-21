package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int; F1 int16; F2 uint64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int8 }
func genFn0(p0 string) string {
    var v7 int = 2147483647
    var v8 int = v7
    switch (v8 | (v7 | v7)) {
    case 32768:
        var v9 int32 = 0
        _ = v9
    case 9223372036854775807:
        var v10 uint = 936
        _ = v10
    }
    switch v7 {
    case 127:
        if (true && true) {
            var v11 uint8 = 196
            _ = v11
        }
    default:
        fmt.Println(p0, v7, v7)
    }
    _ = p0
    _ = v7
    _ = v8
    return "a"
}

func main() {
    var v1 float64 = (1.5 + -2.25)
    var v2 int64 = -771
    var v3 float32 = 0.0
    var v4 float64 = 0.0
    var v5 uint16 = 42
    var v6 uint = 107
    v12 := []int{42, -2147483648}
    v2 -= (v2 & v2)
    if (42 <= -1) {
        var v13 int64 = v2
        _ = v13
    }
    v1 = float64(float32((1e10 * v4)))
    var v14 bool = (!false)
    var v15 float32 = -2.25
    switch ("foo bar" + ("a" + "a")) {
    case "a":
        switch ("hello" + "\n\t") {
        case "":
            switch "a" {
            case "":
                var v16 int = -1
                _ = v16
            default:
                switch "a" {
                case "":
                    var v17 uint32 = 1
                    _ = v17
                default:
                    var v18 int32 = -1
                    _ = v18
                }
            }
        case "\n\t":
            fmt.Println(v2, v3, v1)
        case "hello":
            v15 -= v3
        }
    case "":
        v12 = append(v12, 2147483647)
    case "\n\t":
        var v19 uint64 = 2
        _ = v19
    }
    v12 = append(v12, -1)
    switch 128 {
    case -1:
        for fi20 := 0; fi20 < 4; fi20++ {
            if v14 {
                fmt.Println(v3, v14)
                var v21 uint64 = 4294967295
                for fi22 := 0; fi22 < 9; fi22++ {
                    v23 := GenS1{F0: fi22, F1: ("" + ("\n\t" + "\n\t"))}
                    v3 -= v15
                    if v14 {
                        switch "" {
                        case "a":
                            fmt.Println(v15, v6, v15)
                        case "\n\t":
                            var v24 uint = v6
                            _ = v24
                        case "hello":
                            switch fi20 {
                            case 2147483647:
                                v25 := GenS2{F0: v14, F1: fi22}
                                _ = v25
                            }
                        default:
                            fmt.Println(v21, v1)
                        }
                    }
                    _ = fi22
                    _ = v23
                }
                _ = v21
            }
            v12 = append(v12, fi20)
            v15 += v15
            _ = fi20
        }
    case 128:
        var v26 bool = v14
        _ = v26
    case -149:
        v3 = (float32(0.1) - 3.14159)
    }
    switch "foo bar" {
    case "hello":
        v2 -= (v2 - v2)
    case "\n\t":
        if (v14 || v14) {
            var v27 float32 = ((float32(1.5) - 0.1) / v15)
            _ = v27
        } else {
            v5 = v5
        }
    case "":
        v4 -= 1e10
    default:
        if len(v12) > 0 {
            v12[0] = 146
        }
    }
    v12 = append(v12, 9223372036854775807)
    var v28 int8 = 1
    for fi29 := 0; fi29 < 7; fi29++ {
        v15 = v3
        if ((!(v2 == 42)) && (v14 && v14)) {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
        _ = fi29
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        v2 += -87
        _ = fi30
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 uint32 = 118
            if v14 {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                v34 := GenS2{F0: v14, F1: fi32}
                _ = v34
            } else {
                if (!v14) {
                    var v35 int16 = 0
                    if v14 {
                        v36 := GenS3{F0: (fi32 ^ fi31), F1: (v5 & uint16(v33)), F2: 2}
                        v3 -= v15
                        _ = v36
                    } else {
                        if (v14 || (v14 && v14)) {
                            var v37 uint16 = (^(v5))
                            v33 = (v33 << ((876) & 63))
                            if (v4 == 1.5) {
                                if ((v3 - 0.1) >= ((v3 - 3.14159) - v15)) {
                                    fmt.Println(v5)
                                    if (int16(v6) <= (^(v35))) {
                                        v33 -= 0
                                        fmt.Println(v4, v28)
                                        var v38 int64 = v2
                                        switch fi31 {
                                        case 9223372036854775807:
                                            v12 = append(v12, 4294967295)
                                        default:
                                            if false {
                                                switch fi31 {
                                                case 0:
                                                    var v39 uint = 0
                                                    _ = v39
                                                default:
                                                    var v40 int8 = v28
                                                    _ = v40
                                                }
                                            } else {
                                                var v41 float64 = ((v1 + ((v1 * (float64(float32(1.5)) - v1)) + 1e10)) / (float64(((v15 / (v15 * (v15 + (-2.25 - v3)))) - (float32(float64(-2.25)) + 1.5))) * v1))
                                                switch "foo bar" {
                                                case "":
                                                    var v42 string = "foo bar"
                                                    _ = v42
                                                case "hello":
                                                    v43 := GenS3{F0: (fi31 | 492), F1: v5, F2: v33}
                                                    var v44 uint8 = 127
                                                    if (fi31 <= int(v28)) {
                                                        v41 = v4
                                                    } else {
                                                        if ((!v14) || v14) {
                                                            var v45 int16 = -32768
                                                            var v46 int8 = (v28 >> (((v6 % ((v6) | 1))) & 63))
                                                            v41 = (v1 * v1)
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        if (!false) {
                                                            v47 := GenS1{F0: fi32, F1: ""}
                                                            _ = v47
                                                        }
                                                        if len(v12) > 0 {
                                                            v12[0] = -120
                                                        }
                                                    }
                                                    v28 = -3
                                                    _ = v43
                                                    _ = v44
                                                case "foo bar":
                                                    if (v28 > (v28 % (((v28 ^ 0)) | 1))) {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if ((v1 + (v1 - (v41 - v4))) == (v4 + 1.5)) {
                                                                fmt.Println(v35, v41)
                                                                v33 -= 65535
                                                                if true {
                                                                    fmt.Println(fi31, v3, v6)
                                                                    v48 := GenS2{F0: (!v14), F1: fi32}
                                                                    v49 := GenS3{F0: fi32, F1: (v37 + (v37 + 748)), F2: uint32(v37)}
                                                                    v38 = 0
                                                                    v50 := GenS1{F0: (fi31 ^ 241), F1: "a"}
                                                                    v51 := GenS2{F0: ((v35 ^ (v35 << (((v6 - (v6 << ((18446744073709551615) & 63)))) & 63))) <= v35), F1: (fi31 << ((v6) & 63))}
                                                                    _ = v48
                                                                    _ = v49
                                                                    _ = v50
                                                                    _ = v51
                                                                }
                                                            } else {
                                                                switch (fi32 | (fi31 / ((fi32) | 1))) {
                                                                case 127:
                                                                    v12 = append(v12, fi31)
                                                                }
                                                                v12 = append(v12, (fi31 % ((-30) | 1)))
                                                                v41 -= (v41 - (((1e10 / v4) - v4) - (v1 / v4)))
                                                                if (v5 == (v37 + v5)) {
                                                                    if len(v12) >= 2 {
                                                                        v12 = v12[1:len(v12)]
                                                                    }
                                                                } else {
                                                                    if (v14 && (!(128 == 85))) {
                                                                        v52 := GenS3{F0: (fi32 * 251), F1: v5, F2: (v33 ^ (v33 % ((v33) | 1)))}
                                                                        v3 -= v3
                                                                        var v53 uint8 = 0
                                                                        var v54 uint32 = v33
                                                                        v15 = 0.0
                                                                        v55 := GenS3{F0: fi31, F1: 591, F2: v33}
                                                                        _ = v52
                                                                        _ = v53
                                                                        _ = v54
                                                                        _ = v55
                                                                    }
                                                                }
                                                                var v56 int8 = (v28 >> ((v6) & 63))
                                                                _ = v56
                                                            }
                                                        }
                                                        if (v14 && (false && (!((!(141 > 128)) || v14)))) {
                                                            fmt.Println(v41, v14)
                                                        } else {
                                                            v57 := GenS1{F0: (fi32 | (fi32 % (((fi32 >> ((v6) & 63))) | 1))), F1: ("foo bar" + "")}
                                                            _ = v57
                                                        }
                                                        switch "hello" {
                                                        case "foo bar":
                                                            if len(v12) > 0 {
                                                                v12[0] = (fi32 << ((uint(v33)) & 63))
                                                            }
                                                        case "hello":
                                                            v28 = (v28 << (((v6 + v6)) & 63))
                                                        case "\n\t":
                                                            v58 := GenS4{F0: v2, F1: 128, F2: 42}
                                                            _ = v58
                                                        default:
                                                            switch "hello" {
                                                            case "hello":
                                                                v5 = v5
                                                            case "a":
                                                                v15 -= float32((v41 * ((((v1 / -2.25) - v1) - v41) - ((v41 / v1) * -2.25))))
                                                                if len(v12) >= 2 {
                                                                    v12 = v12[1:len(v12)]
                                                                }
                                                            case "\n\t":
                                                                if len(v12) >= 2 {
                                                                    v12 = v12[1:len(v12)]
                                                                }
                                                            default:
                                                                var v59 string = "hello"
                                                                _ = v59
                                                            }
                                                        }
                                                        if v14 {
                                                            v3 += (float32((v4 / v4)) + ((1e10 / -2.25) / v3))
                                                            if v14 {
                                                                v12 = append(v12, 32767)
                                                                if (824 >= v37) {
                                                                    v60 := GenS3{F0: fi31, F1: (v5 ^ 32767), F2: (v33 << ((v6) & 63))}
                                                                    _ = v60
                                                                } else {
                                                                    v28 = -1
                                                                    var v61 int = (fi31 << ((v6) & 63))
                                                                    _ = v61
                                                                }
                                                            } else {
                                                                if v14 {
                                                                    if false {
                                                                        var v62 float32 = v3
                                                                        var v63 float32 = -2.25
                                                                        if len(v12) >= 2 {
                                                                            v12 = v12[1:len(v12)]
                                                                        }
                                                                        _ = v62
                                                                        _ = v63
                                                                    }
                                                                    switch 42 {
                                                                    case -737:
                                                                        v64 := GenS5{F0: (^(v37))}
                                                                        _ = v64
                                                                    case 4294967295:
                                                                        if ((!(v2 > (v38 & 0))) || v14) {
                                                                            if false {
                                                                                var v65 float64 = (1e10 * v4)
                                                                                if v14 {
                                                                                    v66 := GenS6{F0: (("a" + ("hello" + "hello")) + "\n\t"), F1: -1}
                                                                                    var v67 bool = ("foo bar" > ("a" + "foo bar"))
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                } else {
                                                                                    switch (fi32 & fi31) {
                                                                                    case 9223372036854775807:
                                                                                        var v68 float32 = (((v3 / v3) + v3) + (v3 + v3))
                                                                                        _ = v68
                                                                                    case 1:
                                                                                        v41 = v41
                                                                                    default:
                                                                                        var v69 int8 = (v28 & v28)
                                                                                        _ = v69
                                                                                    }
                                                                                }
                                                                                _ = v65
                                                                            } else {
                                                                                v15 -= 1e10
                                                                                fmt.Println(v2, v14, v15)
                                                                                var v70 int = -2147483648
                                                                                _ = v70
                                                                            }
                                                                            v35 = v35
                                                                        } else {
                                                                            switch ("a" + (("\n\t" + "hello") + "foo bar")) {
                                                                            case "\n\t":
                                                                                _ = len(v12)
                                                                            case "a":
                                                                                v28 = v28
                                                                            case "foo bar":
                                                                                if ((v15 * v15) != 1e10) {
                                                                                    v71 := GenS6{F0: ("foo bar" + "hello"), F1: 0}
                                                                                    var v72 int8 = 0
                                                                                    _ = v71
                                                                                    _ = v72
                                                                                } else {
                                                                                    if ((v38 * 744) < v38) {
                                                                                        if ((v33 >> ((409) & 63)) < v33) {
                                                                                            fmt.Println(v1)
                                                                                        } else {
                                                                                            if (((!v14) && v14) && v14) {
                                                                                                var v73 int16 = v35
                                                                                                switch fi31 {
                                                                                                case 4294967295:
                                                                                                    var v74 string = "a"
                                                                                                    _ = v74
                                                                                                case -2147483648:
                                                                                                    switch fi31 {
                                                                                                    case 610:
                                                                                                        v2 -= 9223372036854775807
                                                                                                    }
                                                                                                case 127:
                                                                                                    if (!((v14 || (-671 != 2147483647)) && v14)) {
                                                                                                        v3 = v3
                                                                                                    }
                                                                                                default:
                                                                                                    _ = len(v12)
                                                                                                }
                                                                                                _ = v73
                                                                                            } else {
                                                                                                var v75 float64 = (v1 + 3.14159)
                                                                                                _ = v75
                                                                                            }
                                                                                            v41 = v1
                                                                                            var v76 uint64 = 1
                                                                                            _ = v76
                                                                                        }
                                                                                        if (!(!v14)) {
                                                                                            v2 -= v38
                                                                                            var v77 float32 = 0.0
                                                                                            if v14 {
                                                                                                var v78 uint64 = 18446744073709551615
                                                                                                v1 = 1.5
                                                                                                v79 := GenS4{F0: 597, F1: 1, F2: -2147483648}
                                                                                                _ = v78
                                                                                                _ = v79
                                                                                            } else {
                                                                                                var v80 int16 = -32768
                                                                                                switch "hello" {
                                                                                                case "a":
                                                                                                    _ = len(v12)
                                                                                                case "\n\t":
                                                                                                    v81 := GenS7{F0: v6}
                                                                                                    _ = v81
                                                                                                case "foo bar":
                                                                                                    if len(v12) >= 2 {
                                                                                                        v12 = v12[1:len(v12)]
                                                                                                    }
                                                                                                }
                                                                                                _ = v80
                                                                                            }
                                                                                            _ = v77
                                                                                        } else {
                                                                                            var v82 uint32 = (^(v33))
                                                                                            v83 := GenS8{F0: fi32, F1: v35, F2: 42}
                                                                                            v84 := GenS6{F0: "a", F1: 1}
                                                                                            if len(v12) > 0 {
                                                                                                v12[0] = (fi31 ^ fi31)
                                                                                            }
                                                                                            _ = v82
                                                                                            _ = v83
                                                                                            _ = v84
                                                                                        }
                                                                                        v85 := GenS5{F0: 65535}
                                                                                        var v86 uint64 = 42
                                                                                        v2 = 569
                                                                                        fmt.Println(fi32)
                                                                                        _ = v85
                                                                                        _ = v86
                                                                                    }
                                                                                    if (!(2 <= 32767)) {
                                                                                        if len(v12) >= 2 {
                                                                                            v12 = v12[1:len(v12)]
                                                                                        }
                                                                                    } else {
                                                                                        if len(v12) >= 2 {
                                                                                            v12 = v12[1:len(v12)]
                                                                                        }
                                                                                    }
                                                                                    var v87 int = int(v35)
                                                                                    v28 = -65
                                                                                    v88 := GenS8{F0: v87, F1: 32767, F2: 18446744073709551615}
                                                                                    _ = v87
                                                                                    _ = v88
                                                                                }
                                                                            default:
                                                                                v89 := GenS9{F0: (!v14)}
                                                                                _ = v89
                                                                            }
                                                                            v33 -= 1
                                                                            v90 := GenS9{F0: v14}
                                                                            _ = v90
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v91 uint = (v6 | v6)
                                                                    fmt.Println(v41, v33)
                                                                    if len(v12) > 0 {
                                                                        v12[0] = fi31
                                                                    }
                                                                    v38 = v2
                                                                    _ = v91
                                                                }
                                                            }
                                                            if ("\n\t" > "foo bar") {
                                                                var v92 uint32 = (v33 ^ v33)
                                                                _ = v92
                                                            } else {
                                                                _ = len(v12)
                                                            }
                                                        }
                                                        if (fi32 <= fi32) {
                                                            v93 := GenS10{F0: (v28 % (((v28 * (v28 - v28))) | 1))}
                                                            var v94 int32 = -747
                                                            v2 -= 0
                                                            _ = v93
                                                            _ = v94
                                                        } else {
                                                            v95 := GenS8{F0: fi32, F1: (^(v35)), F2: 925}
                                                            _ = v95
                                                        }
                                                    } else {
                                                        var v96 int8 = 70
                                                        v2 = int64(fi31)
                                                        v96 -= (v28 + v28)
                                                        _ = v96
                                                    }
                                                }
                                                var v97 uint = (v6 << (((v6 | uint(v37))) & 63))
                                                _ = v41
                                                _ = v97
                                            }
                                        }
                                        _ = v38
                                    }
                                    v5 = v37
                                }
                            } else {
                                if v14 {
                                    _ = len(v12)
                                    v4 -= 1e10
                                    switch fi31 {
                                    case 1:
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                    default:
                                        fmt.Println(v6, v6, v28)
                                    }
                                } else {
                                    var v98 uint32 = 0
                                    var v99 int32 = -1
                                    switch (fi32 / ((fi32) | 1)) {
                                    case 128:
                                        v100 := GenS3{F0: 127, F1: (v5 / ((32767) | 1)), F2: uint32(v6)}
                                        _ = v100
                                    }
                                    _ = v98
                                    _ = v99
                                }
                            }
                            _ = v37
                        }
                        var v101 uint = (v6 / (((v6 | (v6 * 42))) | 1))
                        var v102 float64 = v4
                        fmt.Println(v6, v101, v102)
                        _ = v101
                        _ = v102
                    }
                    _ = v35
                } else {
                    var v103 string = "a"
                    _ = v103
                }
                v104 := GenS10{F0: 1}
                _ = v104
            }
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    if v14 {
        v5 += v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
