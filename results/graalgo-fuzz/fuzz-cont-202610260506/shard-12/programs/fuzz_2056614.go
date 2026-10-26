package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 int32; F2 bool }
type GenS5 struct { F0 int16; F1 int64; F2 uint64 }
type GenS6 struct { F0 int64; F1 int8; F2 int64 }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int16; F1 int32; F2 int64 }
type GenS10 struct { F0 int; F1 int16 }
type GenS11 struct { F0 int8 }
func genFn0(p0 uint64) uint16 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{fi7, fi7}
        _ = fi7
        _ = v8
    }
    p0 += 1
    _ = p0
    return 255
}
func genFn1(p0 uint8) string {
    var v9 uint = 42
    if (!false) {
        fmt.Println(p0, v9)
        p0 = 2
        var v10 float32 = 0.1
        _ = v10
    }
    v11 := []int{42, 113, 32767}
    if true {
        switch -251 {
        case 32768:
            var v12 uint16 = genFn0(18446744073709551615)
            _ = v12
        case -80:
            if ((128 == 128) && (281 == -2147483648)) {
                var v13 bool = true
                for fi14 := 0; fi14 < 5; fi14++ {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    fmt.Println(fi14)
                    _ = fi14
                }
                _ = v13
            }
        case -418:
            var v15 float64 = (0.0 - -2.25)
            _ = v15
        default:
            if len(v11) > 0 {
                v11[0] = -6
            }
        }
        switch "\n\t" {
        case "foo bar":
            v16 := GenS1{F0: 1}
            _ = v16
        case "a":
            for fi17 := 0; fi17 < 15; fi17++ {
                if ((float32(3.14159) >= (-2.25 + (-2.25 - float32(0.0)))) || ((p0 - p0) > 0)) {
                    p0 += p0
                }
                var v18 int32 = -222
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi17
                _ = v18
            }
        case "":
            var v19 int16 = -414
            _ = v19
        }
    } else {
        var v20 int8 = -1
        v9 = 4294967295
        _ = v20
    }
    v9 = v9
    _ = p0
    _ = v11
    _ = v9
    return ("" + "")
}

func main() {
    var v1 uint8 = 0
    var v2 uint64 = 1
    v3 := []int{694, -1, 9223372036854775807}
    var v4 uint64 = v2
    var v5 uint8 = 42
    var v6 int8 = 4
    var v21 int = 9223372036854775807
    for fi22 := 0; fi22 < 10; fi22++ {
        switch v21 {
        case -587:
            var v23 int = (-(v21))
            _ = v23
        case 1:
            var v24 int = 148
            _ = v24
        default:
            for fi25 := 0; fi25 < 3; fi25++ {
                v1 -= (v1 & v5)
                var v26 uint64 = (v2 % (((^(v4))) | 1))
                _ = fi25
                _ = v26
            }
        }
        var v27 string = "\n\t"
        _ = fi22
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        fmt.Println(v2, v2)
        _ = fi28
    }
    v29 := GenS1{F0: -900}
    for fi30 := 0; fi30 < 17; fi30++ {
        for fi31 := 0; fi31 < 11; fi31++ {
            if (fi30 == v21) {
                var v32 uint16 = 97
                _ = v32
            } else {
                if len(v3) > 0 {
                    v3[0] = v21
                }
            }
            switch "foo bar" {
            case "\n\t":
                _ = len(v3)
            }
            if (-768 < 590) {
                var v33 string = "a"
                v34 := GenS2{F0: fi31, F1: (v33 + "hello")}
                if (!((-32768 < 42) || false)) {
                    v35 := GenS2{F0: 82, F1: ("\n\t" + v33)}
                    if len(v3) > 0 {
                        v3[0] = int(v1)
                    }
                    v36 := GenS2{F0: v21, F1: (v33 + "")}
                    _ = v35
                    _ = v36
                }
                _ = v33
                _ = v34
            }
            var v37 uint16 = 514
            _ = fi31
            _ = v37
        }
        v38 := GenS2{F0: (-(v21)), F1: ("" + "\n\t")}
        _ = fi30
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        v40 := GenS3{F0: 42}
        if (int8(v2) <= -128) {
            for fi41 := 0; fi41 < 20; fi41++ {
                _ = len(v3)
                fmt.Println(v21, v2, fi41)
                if (("foo bar" + ("" + "hello")) <= "hello") {
                    fmt.Println(v4, fi41, v21)
                    if len(v3) > 0 {
                        v3[0] = fi39
                    }
                } else {
                    var v42 uint32 = 123
                    if (true || false) {
                        v43 := GenS4{F0: "\n\t", F1: 2147483647, F2: false}
                        fmt.Println(v5)
                        var v44 int16 = 1
                        _ = v43
                        _ = v44
                    } else {
                        if false {
                            if (false && (!true)) {
                                switch 1 {
                                case 128:
                                    var v45 int32 = -1
                                    _ = v45
                                }
                                var v46 string = ("hello" + ("foo bar" + ("" + "hello")))
                                _ = v46
                            }
                            if true {
                                if (!((!false) || true)) {
                                    var v47 float32 = (float32(-2.25) * 1e10)
                                    if (uint16(v42) <= 65535) {
                                        v48 := GenS3{F0: v42}
                                        _ = v48
                                    }
                                    v1 += 127
                                    if true {
                                        var v49 uint16 = genFn0(42)
                                        _ = v49
                                    }
                                    var v50 float64 = -2.25
                                    if (fi39 >= (^(v21))) {
                                        switch fi39 {
                                        case -2147483648:
                                            var v51 uint32 = 2147483648
                                            _ = v51
                                        }
                                        if ((!(v2 >= v2)) && false) {
                                            switch "" {
                                            case "hello":
                                                v52 := GenS3{F0: (v42 ^ (v42 | v42))}
                                                _ = v52
                                            }
                                        }
                                        v5 = uint8(v42)
                                    }
                                    _ = v47
                                    _ = v50
                                }
                                var v53 int8 = v6
                                _ = v53
                            } else {
                                v5 -= (v1 - (v5 >> ((42) & 63)))
                            }
                            if true {
                                v5 = v1
                                var v54 int32 = -1
                                v4 += uint64(v5)
                                _ = v54
                            } else {
                                switch fi39 {
                                case 128:
                                    fmt.Println(fi41, fi39, v1)
                                    var v55 uint64 = v2
                                    _ = v55
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                if len(v3) > 0 {
                                    v3[0] = v21
                                }
                                fmt.Println(v1)
                                if (false || (false || (42 < 324))) {
                                    if false {
                                        switch (fi39 >> ((65535) & 63)) {
                                        case -571:
                                            v21 = fi39
                                        case 2147483647:
                                            if len(v3) > 0 {
                                                v3[0] = fi41
                                            }
                                        default:
                                            v3 = append(v3, (^(v21)))
                                        }
                                        v6 = v6
                                        v6 = 42
                                        if ((1 > -9223372036854775808) || (true || (false && false))) {
                                            v56 := GenS4{F0: "foo bar", F1: -635, F2: (!(!((v6 % (((v6 ^ v6)) | 1)) < -1)))}
                                            _ = v56
                                        }
                                    } else {
                                        if (((v1 + 255) == v1) || true) {
                                            var v57 int32 = -922
                                            v5 -= (v1 - (v5 << ((65535) & 63)))
                                            fmt.Println(v6, v6)
                                            _ = v57
                                        } else {
                                            if ("" != "a") {
                                                fmt.Println(fi41)
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                        }
                                        var v58 uint8 = (v1 - (v5 & (^(v1))))
                                        _ = v58
                                    }
                                } else {
                                    v59 := GenS1{F0: int64(v42)}
                                    switch -1 {
                                    case 2147483647:
                                        var v60 int64 = 35
                                        _ = v60
                                    default:
                                        if (((((^(v6)) <= 127) && (!true)) || (true && false)) || (0 != (^(v4)))) {
                                            v1 -= v1
                                            v2 = 573
                                            var v61 uint8 = v5
                                            _ = v61
                                        }
                                    }
                                    _ = v59
                                }
                                if true {
                                    switch "" {
                                    case "a":
                                        v42 -= (^(v42))
                                    case "foo bar":
                                        v62 := GenS3{F0: (^(v42))}
                                        var v63 int32 = -23
                                        v64 := GenS3{F0: v42}
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                    case "hello":
                                        v3 = append(v3, 145)
                                    default:
                                        var v65 int16 = -689
                                        _ = v65
                                    }
                                } else {
                                    v21 -= (fi41 * (v21 ^ fi41))
                                    if len(v3) > 0 {
                                        v3[0] = 884
                                    }
                                }
                                if true {
                                    v42 = 42
                                }
                                if len(v3) > 0 {
                                    v3[0] = fi39
                                }
                                var v66 int64 = 1
                                v3 = append(v3, (v21 % ((128) | 1)))
                                v1 += v5
                                v67 := GenS5{F0: 527, F1: (v66 * v66), F2: v4}
                                if ((!true) && ((v4 | (v4 << ((28) & 63))) <= (^(v4)))) {
                                    v68 := GenS1{F0: int64(v2)}
                                    var v69 int64 = v66
                                    if true {
                                        var v70 uint16 = 1
                                        _ = v70
                                    } else {
                                        fmt.Println(v21, v5, v66)
                                        fmt.Println(v4, v42, v5)
                                        var v71 float32 = 0.1
                                        _ = v71
                                    }
                                    _ = v68
                                    _ = v69
                                }
                                _ = v66
                                _ = v67
                            }
                        } else {
                            v6 += (^(v6))
                            var v72 uint64 = (^(v2))
                            _ = v72
                        }
                    }
                    v73 := GenS1{F0: -9223372036854775808}
                    _ = v42
                    _ = v73
                }
                _ = fi41
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (-2.25 != (0.0 * (float64(1e10) * 0.1))) {
                if (true && (452 != 776)) {
                    v74 := GenS6{F0: -9223372036854775808, F1: (^(v6)), F2: 699}
                    switch (("foo bar" + "a") + (("hello" + ("a" + "a")) + "a")) {
                    case "":
                        fmt.Println(v5, v2)
                    default:
                        var v75 uint64 = (v4 % (((v2 % (((v2 / ((v2) | 1))) | 1))) | 1))
                        _ = v75
                    }
                    _ = v74
                } else {
                    if (v4 == 916) {
                        v2 = v2
                        var v76 uint16 = 0
                        v77 := GenS7{F0: -52, F1: ""}
                        _ = v76
                        _ = v77
                    } else {
                        v78 := GenS7{F0: v6, F1: "hello"}
                        fmt.Println(v5)
                        _ = v78
                    }
                }
            } else {
                var v79 uint = 908
                var v80 float32 = (1.5 - 1e10)
                v81 := GenS6{F0: 42, F1: v6, F2: 9223372036854775807}
                _ = v79
                _ = v80
                _ = v81
            }
        } else {
            if false {
                var v82 uint64 = v2
                var v83 int8 = -128
                var v84 uint32 = 65535
                _ = v82
                _ = v83
                _ = v84
            } else {
                var v85 uint16 = genFn0((v2 * (v4 - 323)))
                switch "a" {
                case "":
                    for fi86 := 0; fi86 < 17; fi86++ {
                        if ((v85 + v85) > v85) {
                            var v87 int64 = 344
                            v5 = (v5 % (((v1 >> ((18446744073709551615) & 63))) | 1))
                            _ = v87
                        }
                        _ = fi86
                    }
                case "\n\t":
                    v88 := GenS7{F0: v6, F1: ""}
                    _ = v88
                case "foo bar":
                    for fi89 := 0; fi89 < 13; fi89++ {
                        var v90 uint64 = (v4 + (v4 / ((v2) | 1)))
                        var v91 float64 = (((0.1 * 3.14159) * (0.0 + (0.1 * (1e10 + 1.5)))) + -2.25)
                        _ = fi89
                        _ = v90
                        _ = v91
                    }
                }
                v3 = append(v3, fi39)
                _ = v85
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi39
        _ = v40
    }
    fmt.Println(v21)
    var v92 int32 = 1
    v21 = v21
    var v93 uint16 = genFn0(v2)
    for fi94 := 0; fi94 < 20; fi94++ {
        if (9223372036854775807 < 912) {
            _ = len(v3)
            v21 -= v21
        }
        switch "\n\t" {
        case "":
            if true {
                var v95 float32 = (1e10 + ((1e10 * (0.0 + (float32(3.14159) * ((1.5 - 3.14159) * 1.5)))) - -2.25))
                v96 := GenS5{F0: -532, F1: 9223372036854775807, F2: (v4 & v4)}
                _ = v95
                _ = v96
            } else {
                if len(v3) > 0 {
                    v3[0] = (v21 + v21)
                }
                v6 -= 42
            }
        default:
            for fi97 := 0; fi97 < 9; fi97++ {
                var v98 int16 = -1
                fmt.Println(fi94, v92)
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v99 := GenS8{F0: 255}
                    switch 724 {
                    case 0:
                        switch ("\n\t" + "hello") {
                        case "":
                            if (!false) {
                                var v100 float64 = 1e10
                                v1 = v5
                                _ = v100
                            }
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    default:
                        fmt.Println(v2, v93, v1)
                    }
                    fmt.Println(fi94, fi97, v6)
                    _ = v99
                }
                _ = fi97
                _ = v98
            }
        }
        for fi101 := 0; fi101 < 15; fi101++ {
            fmt.Println(v1, v21, v92)
            _ = fi101
        }
        _ = fi94
    }
    if (-1 <= 42) {
        v102 := GenS5{F0: 0, F1: 1, F2: (v4 << ((841) & 63))}
        switch 128 {
        case 32767:
            var v103 string = ("foo bar" + ("foo bar" + (("" + (("a" + (("\n\t" + ("" + ("foo bar" + "hello"))) + ("\n\t" + ("\n\t" + "foo bar")))) + (("a" + ("hello" + ("\n\t" + "hello"))) + "a"))) + ("" + "\n\t"))))
            _ = v103
        case 127:
            v3 = append(v3, v21)
        case 42:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            if len(v3) > 0 {
                v3[0] = v21
            }
        }
        v92 += (v92 & v92)
        _ = v102
    } else {
        v104 := GenS6{F0: int64(v1), F1: (v6 >> ((637) & 63)), F2: -1}
        v105 := GenS3{F0: 4294967295}
        v1 += (v5 / ((v5) | 1))
        _ = v104
        _ = v105
    }
    if true {
        v106 := GenS9{F0: -1, F1: (v92 / (((v92 >> ((uint(v93)) & 63))) | 1)), F2: 42}
        _ = v106
    } else {
        v93 -= (^(v93))
        if false {
            var v107 int64 = -1
            fmt.Println(v4, v2, v107)
            _ = v107
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v2, v1, v21)
            v108 := GenS10{F0: -612, F1: 207}
            _ = v108
        }
    }
    v109 := GenS11{F0: -128}
    v4 += v4
    switch "foo bar" {
    case "\n\t":
        if (true && true) {
            fmt.Println(v5)
            v110 := GenS7{F0: (v6 | (v6 * (v6 | (v6 / ((v6) | 1))))), F1: "hello"}
            _ = v110
        }
    case "hello":
        fmt.Println(v5, v2, v1)
    case "a":
        v93 += (^(v93))
    default:
        for fi111 := 0; fi111 < 19; fi111++ {
            fmt.Println(v2, v92, v5)
            _ = fi111
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v109
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v92
    _ = v93
}
