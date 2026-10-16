package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int64; F2 int64 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS4 struct { F0 int; F1 uint; F2 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 int8; F1 int; F2 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint; F1 int16; F2 uint8 }
type GenS10 struct { F0 int64; F1 string }
type GenS11 struct { F0 int32; F1 int32; F2 uint }
type GenS12 struct { F0 int8; F1 int16 }
type GenS13 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS14 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS15 struct { F0 int8 }
func genFn0(p0 bool) string {
    var v7 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int = (fi8 / (((^(fi8))) | 1))
        var v10 uint64 = 555
        fmt.Println(fi8, v10)
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 128 {
    case 0:
        fmt.Println(v7, v7)
    }
    _ = p0
    _ = v7
    return (("\n\t" + "a") + ("" + "foo bar"))
}

func main() {
    var v1 int16 = 687
    v2 := []string{"\n\t", "foo bar", "hello"}
    var v3 uint32 = 65535
    v4 := make(map[uint32]string)
    var v5 int64 = 42
    v6 := []string{"\n\t", (("\n\t" + "hello") + ""), ("hello" + "hello"), "foo bar"}
    var v11 int64 = v5
    var v12 int32 = 1
    if (true && true) {
        var v13 int16 = -799
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 uint = 4294967295
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        for fi16 := 0; fi16 < 16; fi16++ {
            v11 = v5
            if (2 > 4294967295) {
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 uint8 = 128
                    var v19 int32 = (^(v12))
                    v3 = (^(v3))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v20 := GenS1{F0: 65535}
                _ = v20
            } else {
                fmt.Println(v11)
            }
            fmt.Println(v3, v5, v3)
            _ = fi16
        }
        v1 = v1
    }
    if (!false) {
        _ = len(v2)
    }
    var v21 bool = (!(true || (v12 < (v12 % ((v12) | 1)))))
    v22 := GenS2{F0: int16(v5), F1: v11, F2: 9223372036854775807}
    if (((!(!v21)) || ((int8(v1) <= 0) || ((v3 << ((42) & 63)) <= v3))) && (true && (v21 || (v21 || v21)))) {
        if v21 {
            var v23 uint16 = 32767
            switch -2147483648 {
            case 9223372036854775807:
                v21 = (851 > 868)
            case 231:
                for fi24 := 0; fi24 < 3; fi24++ {
                    if true {
                        if v21 {
                            var v25 uint64 = 512
                            delete(v4, (^(v3)))
                            _ = v25
                        } else {
                            v26 := GenS1{F0: v23}
                            _ = v26
                        }
                        if true {
                            for fi27 := 0; fi27 < 8; fi27++ {
                                switch "hello" {
                                case "":
                                    if false {
                                        var v28 int64 = 9223372036854775807
                                        v5 = v11
                                        _ = v28
                                    } else {
                                        v1 = v1
                                    }
                                case "a":
                                    if v21 {
                                        v3 = 4294967295
                                        switch "a" {
                                        case "a":
                                            delete(v4, 42)
                                            v11 -= v11
                                            if len(v2) > 0 {
                                                v2[0] = (("\n\t" + ("hello" + ("a" + ("\n\t" + "foo bar")))) + "")
                                            }
                                        case "foo bar":
                                            var v29 int8 = 100
                                            _ = v29
                                        case "hello":
                                            if len(v6) > 0 {
                                                v6[0] = ""
                                            }
                                        default:
                                            v12 -= (v12 / ((v12) | 1))
                                        }
                                        var v30 uint64 = 2
                                        var v31 float64 = (0.0 * 3.14159)
                                        _ = v30
                                        _ = v31
                                    } else {
                                        if (0.1 > (0.0 * 3.14159)) {
                                            var v32 uint = 998
                                            v4[v3] = ""
                                            _ = v32
                                        } else {
                                            switch (fi27 & fi27) {
                                            case -2147483648:
                                                fmt.Println(fi27, fi27, v3)
                                            case 128:
                                                fmt.Println(v5, v23, v1)
                                            case 0:
                                                if (42 >= -2147483648) {
                                                    v3 += v3
                                                    var v33 uint32 = (v3 / ((2147483648) | 1))
                                                    v21 = (!true)
                                                    _ = v33
                                                }
                                            default:
                                                if v21 {
                                                    var v34 int8 = -43
                                                    var v35 int = fi27
                                                    _ = len(v4)
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    var v36 uint64 = 4
                                                    v37 := GenS3{F0: (v36 & v36), F1: -104, F2: v23}
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v38 uint64 = (v36 << ((0) & 63))
                                                        _ = v38
                                                    case "a":
                                                        v12 = (^(v12))
                                                    default:
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v39 uint32 = v3
                                                            _ = v39
                                                        default:
                                                            v5 = (v11 & 0)
                                                        }
                                                    }
                                                    _ = v36
                                                    _ = v37
                                                }
                                            }
                                            switch fi27 {
                                            case -45:
                                                fmt.Println(fi24, v12, v21)
                                            default:
                                                var v40 int32 = (v12 | 2147483647)
                                                _ = v40
                                            }
                                        }
                                    }
                                case "foo bar":
                                    var v41 uint = 314
                                    _ = v41
                                }
                                fmt.Println(v23)
                                fmt.Println(v1, fi27, v5)
                                if (!v21) {
                                    v21 = (v3 < v3)
                                } else {
                                    var v42 int32 = 2147483647
                                    if false {
                                        switch 774 {
                                        case 32767:
                                            _ = len(v4)
                                        case -2147483648:
                                            v4[1] = ((("a" + "\n\t") + "\n\t") + "\n\t")
                                            var v43 int8 = -1
                                            switch "foo bar" {
                                            case "hello":
                                                v3 -= (v3 >> ((1) & 63))
                                            }
                                            _ = v43
                                        default:
                                            if ((("a" + "\n\t") + "\n\t") < "\n\t") {
                                                var v44 int8 = -1
                                                _ = v44
                                            } else {
                                                switch "a" {
                                                case "":
                                                    v23 -= (v23 << ((589) & 63))
                                                case "\n\t":
                                                    var v45 uint = 255
                                                    _ = v45
                                                default:
                                                    fmt.Println(v1, v11)
                                                }
                                                if v21 {
                                                    v42 = (v42 >> ((4294967295) & 63))
                                                    v46 := GenS4{F0: (fi27 - (fi24 + 162)), F1: 2, F2: -647}
                                                    v47 := GenS5{F0: "hello"}
                                                    v6 = append(v6, ("a" + "\n\t"))
                                                    switch (-(fi27)) {
                                                    case -2147483648:
                                                        v21 = v21
                                                    case 42:
                                                        fmt.Println(v1, v42, fi27)
                                                    case -496:
                                                        var v48 string = "a"
                                                        var v49 int16 = (v1 * (v1 << ((65535) & 63)))
                                                        delete(v4, (v3 % ((v3) | 1)))
                                                        _ = v48
                                                        _ = v49
                                                    default:
                                                        v42 = v42
                                                        v50 := GenS6{F0: 787, F1: v23}
                                                        _ = v50
                                                    }
                                                    fmt.Println(v5, v11, v23)
                                                    v1 -= (v1 - (v1 + int16(v23)))
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    fmt.Println(v23)
                                                    v1 -= 898
                                                    v12 = -2147483648
                                                    v51 := GenS7{F0: 100, F1: (fi24 * fi24), F2: (v1 << ((2) & 63))}
                                                    _ = v51
                                                }
                                                if v21 {
                                                    fmt.Println(fi27)
                                                } else {
                                                    v52 := GenS1{F0: (v23 >> ((65535) & 63))}
                                                    if (!((1 == 187) && ((v11 ^ (v5 % (((v11 | v11)) | 1))) != v11))) {
                                                        switch 2147483647 {
                                                        case 919:
                                                            if (!v21) {
                                                                fmt.Println(v23, v21, v42)
                                                            }
                                                        case 127:
                                                            v23 = 65535
                                                        case 42:
                                                            switch fi27 {
                                                            case 32767:
                                                                v23 = (v23 << ((4294967295) & 63))
                                                            }
                                                        default:
                                                            if len(v2) > 0 {
                                                                v2[0] = "a"
                                                            }
                                                        }
                                                    } else {
                                                        var v53 int32 = (v12 ^ v42)
                                                        _ = v53
                                                    }
                                                    v11 = v11
                                                    _ = v52
                                                }
                                            }
                                        }
                                        var v54 uint64 = uint64(v3)
                                        _ = v54
                                    } else {
                                        if ((v5 >> ((378) & 63)) >= (v5 % ((v5) | 1))) {
                                            switch "foo bar" {
                                            case "":
                                                v6 = append(v6, "\n\t")
                                            case "\n\t":
                                                if v21 {
                                                    fmt.Println(v1)
                                                }
                                            }
                                            var v55 float32 = (-2.25 + 0.1)
                                            var v56 int = (fi27 << ((287) & 63))
                                            _ = v55
                                            _ = v56
                                        } else {
                                            fmt.Println(fi24, v11)
                                            v57 := GenS5{F0: "foo bar"}
                                            v58 := GenS7{F0: 0, F1: -402, F2: v1}
                                            _ = v57
                                            _ = v58
                                        }
                                    }
                                    v3 = (v3 << ((0) & 63))
                                    _ = v42
                                }
                                _ = fi27
                            }
                            var v59 float64 = (float64(1.5) * 1e10)
                            if (32 != -84) {
                                v60 := GenS8{F0: 42}
                                var v61 int = (fi24 << ((42) & 63))
                                if false {
                                    var v62 uint = 255
                                    v3 += v3
                                    var v63 float64 = v59
                                    _ = v62
                                    _ = v63
                                }
                                _ = v60
                                _ = v61
                            } else {
                                if (v1 < (v1 * v1)) {
                                    if (v59 != 1.5) {
                                        v59 = (v59 + (v59 / 0.1))
                                        switch (fi24 & (fi24 >> ((65535) & 63))) {
                                        case 128:
                                            for fi64 := 0; fi64 < 18; fi64++ {
                                                v1 += (v1 | v1)
                                                _ = fi64
                                            }
                                        case 127:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case 42:
                                            fmt.Println(v23, v21)
                                        default:
                                            for fi65 := 0; fi65 < 8; fi65++ {
                                                v66 := GenS9{F0: 18446744073709551615, F1: (v1 % (((v1 - -1)) | 1)), F2: 135}
                                                v67 := GenS1{F0: v23}
                                                _ = fi65
                                                _ = v66
                                                _ = v67
                                            }
                                        }
                                    } else {
                                        switch (^(fi24)) {
                                        case 1:
                                            v68 := GenS10{F0: 9223372036854775807, F1: "a"}
                                            _ = v68
                                        case 0:
                                            var v69 float64 = v59
                                            _ = v69
                                        }
                                    }
                                    v21 = v21
                                } else {
                                    v70 := GenS1{F0: (v23 >> ((0) & 63))}
                                    fmt.Println(v11, v5)
                                    _ = v70
                                }
                            }
                            _ = v59
                        }
                        v11 -= (v11 / ((v11) | 1))
                    }
                    var v71 uint8 = 128
                    if true {
                        for fi72 := 0; fi72 < 8; fi72++ {
                            fmt.Println(v11, fi24)
                            v11 -= (v11 << ((2) & 63))
                            _ = fi72
                        }
                    }
                    _ = fi24
                    _ = v71
                }
            case 1:
                var v73 int = 12
                _ = v73
            default:
                if false {
                    v4[v3] = "a"
                }
            }
            switch (("hello" + "hello") + (("\n\t" + (("foo bar" + (((("hello" + "hello") + "\n\t") + "\n\t") + "\n\t")) + "\n\t")) + ("hello" + ("foo bar" + ("a" + "foo bar"))))) {
            case "a":
                if v21 {
                    switch -738 {
                    case -953:
                        if ((v5 + (v11 - v5)) >= 888) {
                            switch "foo bar" {
                            case "hello":
                                for fi74 := 0; fi74 < 17; fi74++ {
                                    switch "a" {
                                    case "\n\t":
                                        switch "" {
                                        case "\n\t":
                                            v23 -= v23
                                        case "a":
                                            for fi75 := 0; fi75 < 18; fi75++ {
                                                var v76 int32 = (v12 % ((int32(v3)) | 1))
                                                v12 = v76
                                                v76 += int32(v1)
                                                _ = fi75
                                                _ = v76
                                            }
                                        case "foo bar":
                                            v5 = 42
                                        }
                                    case "foo bar":
                                        var v77 uint16 = v23
                                        _ = v77
                                    case "a":
                                        v21 = (1 >= 255)
                                    default:
                                        v78 := GenS11{F0: 0, F1: v12, F2: 744}
                                        _ = v78
                                    }
                                    var v79 int8 = -44
                                    _ = fi74
                                    _ = v79
                                }
                            case "":
                                if true {
                                    for fi80 := 0; fi80 < 12; fi80++ {
                                        switch (fi80 ^ -9223372036854775808) {
                                        case 2147483647:
                                            var v81 int8 = 100
                                            _ = v81
                                        case 9223372036854775807:
                                            for fi82 := 0; fi82 < 20; fi82++ {
                                                v83 := GenS2{F0: v1, F1: (v11 >> ((440) & 63)), F2: (v11 | (v5 - v11))}
                                                if (151 != uint(v23)) {
                                                    fmt.Println(v12)
                                                } else {
                                                    var v84 uint16 = (v23 + (v23 | 820))
                                                    v5 -= 9223372036854775807
                                                    _ = len(v4)
                                                    _ = v84
                                                }
                                                var v85 uint64 = 4294967295
                                                _ = fi82
                                                _ = v83
                                                _ = v85
                                            }
                                        }
                                        if v21 {
                                            fmt.Println(v5)
                                            if false {
                                                var v86 int8 = 127
                                                _ = v86
                                            }
                                        } else {
                                            v87 := GenS9{F0: 2, F1: (v1 / ((-513) | 1)), F2: 127}
                                            v88 := GenS12{F0: 127, F1: 42}
                                            switch (fi80 | fi80) {
                                            case 32768:
                                                delete(v4, v3)
                                            case -471:
                                                for fi89 := 0; fi89 < 2; fi89++ {
                                                    _ = len(v2)
                                                    v90 := GenS13{F0: 88, F1: (-(v1)), F2: (v3 - 65535)}
                                                    _ = fi89
                                                    _ = v90
                                                }
                                            case -542:
                                                v91 := GenS5{F0: ("\n\t" + "\n\t")}
                                                _ = v91
                                            }
                                            _ = v87
                                            _ = v88
                                        }
                                        _ = fi80
                                    }
                                    var v92 bool = false
                                    _ = v92
                                }
                            default:
                                for fi93 := 0; fi93 < 4; fi93++ {
                                    v94 := v4[uint32(v23)]
                                    if (-2147483648 == fi93) {
                                        v95 := GenS13{F0: 255, F1: (v1 + (v1 & v1)), F2: 65535}
                                        v23 = 255
                                        _ = v95
                                    }
                                    fmt.Println(v21, v3, v21)
                                    _ = fi93
                                    _ = v94
                                }
                            }
                            delete(v4, v3)
                            var v96 int8 = 100
                            _ = v96
                        } else {
                            fmt.Println(v1)
                        }
                    default:
                        fmt.Println(v23, v12, v21)
                    }
                    _ = len(v4)
                    v12 += v12
                }
            }
            _ = v23
        }
        for fi97 := 0; fi97 < 2; fi97++ {
            v12 = (v12 | v12)
            var v98 float32 = float32(-2.25)
            _ = fi97
            _ = v98
        }
    } else {
        if v21 {
            var v99 string = ""
            _ = v99
        } else {
            var v100 int16 = -1
            v101 := GenS1{F0: 96}
            var v102 uint64 = 42
            _ = v100
            _ = v101
            _ = v102
        }
        switch ("\n\t" + "foo bar") {
        case "hello":
            v103 := GenS14{F0: (v12 * (v12 * v12)), F1: 2, F2: 42}
            _ = v103
        case "a":
            var v104 uint = 1
            _ = v104
        }
    }
    switch "" {
    case "hello":
        v12 = (^(v12))
    default:
        if v21 {
            var v105 int = 2147483647
            var v106 bool = (18446744073709551615 >= 60)
            if true {
                var v107 uint8 = 2
                v106 = (2 < 255)
                _ = v107
            } else {
                switch -1 {
                case 4294967295:
                    delete(v4, (v3 * (v3 / ((831) | 1))))
                case 32768:
                    if v106 {
                        var v108 uint = 42
                        v21 = false
                        _ = v108
                    } else {
                        var v109 uint64 = 766
                        _ = v109
                    }
                default:
                    var v110 uint = 586
                    _ = v110
                }
            }
            _ = v105
            _ = v106
        } else {
            fmt.Println(v11, v5, v5)
        }
    }
    var v111 float32 = -2.25
    fmt.Println(v111, v21, v21)
    fmt.Println(v12)
    fmt.Println(v21)
    for fi112 := 0; fi112 < 14; fi112++ {
        switch "\n\t" {
        case "\n\t":
            var v113 int16 = (v1 >> ((4294967295) & 63))
            _ = v113
        case "":
            for fi114 := 0; fi114 < 13; fi114++ {
                fmt.Println(v12, v21)
                _ = fi114
            }
        default:
            v5 += (^(v5))
        }
        switch (^(fi112)) {
        case 9223372036854775807:
            var v115 uint = 0
            _ = v115
        case 817:
            v111 += (v111 - -2.25)
        case -889:
            if ((987 > 317) && (v21 || ((284 <= (^(v3))) && v21))) {
                v3 += (v3 << ((3) & 63))
                delete(v4, 817)
            } else {
                fmt.Println(fi112, v21, v11)
                v116 := GenS15{F0: 60}
                _ = v116
            }
        }
        _ = fi112
    }
    switch -478 {
    case 32767:
        delete(v4, v3)
    case 9223372036854775807:
        var v117 int8 = -1
        _ = v117
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v111:", v111)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v111
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
