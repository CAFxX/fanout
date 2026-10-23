package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint; F1 int32; F2 int64 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 int8; F2 bool }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint32; F1 int16 }
type GenS11 struct { F0 uint32; F1 uint8 }
type GenS12 struct { F0 int; F1 int8 }
type GenS13 struct { F0 uint32; F1 bool }
type GenS14 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 int8; F1 uint32; F2 int8 }
type GenS21 struct { F0 bool; F1 uint; F2 uint }
type GenS22 struct { F0 uint32; F1 uint; F2 uint }
type GenS23 struct { F0 int; F1 uint32; F2 int32 }
type GenS24 struct { F0 bool; F1 string; F2 uint32 }
type GenS25 struct { F0 uint64 }
func genFn0(p0 uint32, p1 bool) int32 {
    fmt.Println(p0, p0)
    var v6 int16 = 1000
    _ = p0
    _ = p1
    _ = v6
    return -1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    v3 := []string{"hello", (("hello" + "\n\t") + ("foo bar" + ("\n\t" + "\n\t"))), ""}
    var v4 int = -529
    v5 := []string{("a" + "a"), "\n\t", "a"}
    if false {
        if false {
            v7 := GenS1{F0: 954}
            v4 += (v4 / (((v4 % ((-769) | 1))) | 1))
            var v8 uint32 = 1
            _ = v7
            _ = v8
        }
    }
    var v9 bool = true
    fmt.Println(v4, v4)
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int8 = -1
        _ = fi10
        _ = v11
    }
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    delete(v1, v4)
    v12 := GenS1{F0: 686}
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint8 = 255
        v15 := GenS1{F0: 0}
        switch ("foo bar" + ("" + "hello")) {
        case "a":
            if (0 < -907) {
                v16 := GenS2{F0: v9}
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := GenS1{F0: 2}
                    v19 := v1[(-(v4))]
                    var v20 uint16 = 176
                    v21 := GenS2{F0: (1e10 < float64(1.5))}
                    var v22 bool = (743 == 115)
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                if ((("" + "a") + "") < "a") {
                    var v23 float32 = float32(3.14159)
                    switch "foo bar" {
                    case "\n\t":
                        var v24 int16 = 32767
                        _ = v24
                    case "foo bar":
                        v14 -= 255
                    default:
                        switch (fi13 ^ 32767) {
                        case -9223372036854775808:
                            v25 := GenS1{F0: 150}
                            _ = v25
                        case -2147483648:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        default:
                            for fi26 := 0; fi26 < 19; fi26++ {
                                v23 += v23
                                v9 = (true && ((v14 - (v14 ^ 128)) == (v14 % (((v14 % (((^(v14))) | 1))) | 1))))
                                if ((342 > 42) && (!(32767 > -1))) {
                                    v5 = append(v5, "a")
                                    v14 -= (v14 / ((42) | 1))
                                    v27 := GenS3{F0: 864, F1: 100}
                                    _ = v27
                                } else {
                                    var v28 uint = 1
                                    v9 = v9
                                    _ = len(v1)
                                    _ = v28
                                }
                                _ = fi26
                            }
                        }
                    }
                    _ = v23
                } else {
                    v4 = 174
                    if v9 {
                        v29 := GenS4{F0: 694, F1: genFn0(425, (32767 >= -691)), F2: -884}
                        _ = v29
                    }
                    var v30 float64 = 0.0
                    _ = v30
                }
                _ = v16
            } else {
                if (0 < v14) {
                    if (true || v9) {
                        v31 := GenS5{F0: 0, F1: 312}
                        fmt.Println(fi13, fi13)
                        _ = v31
                    }
                }
                if true {
                    v32 := GenS4{F0: 2, F1: -2147483648, F2: 1}
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                    for fi33 := 0; fi33 < 10; fi33++ {
                        fmt.Println(fi33, fi33)
                        var v34 uint16 = 204
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                } else {
                    if v9 {
                        var v35 uint8 = (v14 | v14)
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 16; fi36++ {
                        v37 := GenS6{F0: -2147483648, F1: -921}
                        fmt.Println(v9, v9)
                        _ = fi36
                        _ = v37
                    }
                    v38 := GenS5{F0: -128, F1: 1}
                    _ = v38
                }
            }
        case "hello":
            if v9 {
                v14 -= (v14 % ((2) | 1))
                if v9 {
                    v39 := GenS6{F0: int32(v4), F1: (^(fi13))}
                    if false {
                        v14 += 2
                    }
                    _ = v39
                }
            } else {
                var v40 int64 = 0
                var v41 uint = 63
                _ = v40
                _ = v41
            }
        default:
            var v42 int8 = 100
            _ = v42
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v4 -= v4
    v9 = (2 <= 532)
    if ("\n\t" < "hello") {
        v9 = false
        for fi43 := 0; fi43 < 7; fi43++ {
            for fi44 := 0; fi44 < 5; fi44++ {
                if (!(65535 != 42)) {
                    switch -1 {
                    case 2147483647:
                        if false {
                            var v45 int16 = 42
                            var v46 uint16 = 0
                            _ = v45
                            _ = v46
                        }
                    default:
                        if len(v5) > 0 {
                            v5[0] = "hello"
                        }
                    }
                    v47 := GenS6{F0: 603, F1: fi44}
                    if (1.5 > -2.25) {
                        fmt.Println(fi44)
                        v4 = v4
                    } else {
                        v4 = fi44
                        fmt.Println(fi44, v9)
                        if false {
                            if v9 {
                                var v48 uint8 = 0
                                v4 = v4
                                _ = v48
                            } else {
                                switch v4 {
                                case 128:
                                    if v9 {
                                        var v49 uint64 = 42
                                        _ = v49
                                    }
                                case 32768:
                                    if true {
                                        switch fi43 {
                                        case 1:
                                            v3 = append(v3, "hello")
                                        case -1:
                                            v4 -= v4
                                            var v50 int8 = 100
                                            _ = v50
                                        case 128:
                                            var v51 bool = v9
                                            _ = v51
                                        default:
                                            if true {
                                                v9 = (304 >= 1)
                                            } else {
                                                if true {
                                                    if (666 >= 42) {
                                                        if false {
                                                            switch v4 {
                                                            case -9223372036854775808:
                                                                v52 := GenS5{F0: -1, F1: 375}
                                                                _ = v52
                                                            case 127:
                                                                var v53 float64 = -2.25
                                                                _ = v53
                                                            case -1:
                                                                var v54 bool = v9
                                                                _ = v54
                                                            default:
                                                                v9 = (102 > -1)
                                                            }
                                                        }
                                                        v4 -= (fi43 / (((-(fi43))) | 1))
                                                        v9 = (-1 <= 79)
                                                    }
                                                } else {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        switch (^(v4)) {
                                                        case -955:
                                                            if v9 {
                                                                switch (("hello" + "hello") + "a") {
                                                                case "a":
                                                                    if v9 {
                                                                        fmt.Println(fi43)
                                                                        fmt.Println(v4)
                                                                    }
                                                                default:
                                                                    v55 := GenS7{F0: v9}
                                                                    _ = v55
                                                                }
                                                            } else {
                                                                if (-1 != -1) {
                                                                    if v9 {
                                                                        switch 32768 {
                                                                        case 32767:
                                                                            v56 := v1[4294967295]
                                                                            _ = v56
                                                                        case 128:
                                                                            v57 := GenS8{F0: 100, F1: -128, F2: v9}
                                                                            _ = v57
                                                                        default:
                                                                            fmt.Println(fi43, fi44, fi44)
                                                                        }
                                                                        v58 := v1[fi44]
                                                                        v4 -= 9223372036854775807
                                                                        _ = v58
                                                                    }
                                                                    v4 += 42
                                                                    v4 = (-(v4))
                                                                    if (!true) {
                                                                        v4 = -152
                                                                        switch fi44 {
                                                                        case 760:
                                                                            v59 := GenS2{F0: v9}
                                                                            _ = v59
                                                                        case 1:
                                                                            v9 = (!v9)
                                                                        default:
                                                                            v60 := GenS9{F0: 173}
                                                                            _ = v60
                                                                        }
                                                                        v61 := GenS10{F0: 2147483648, F1: -460}
                                                                        _ = v61
                                                                    } else {
                                                                        var v62 uint8 = 0
                                                                        fmt.Println(fi43, v62, v9)
                                                                        v62 = v62
                                                                        v63 := GenS7{F0: (v9 || ((2147483648 > 868) || (0 != 2147483647)))}
                                                                        v64 := GenS11{F0: 65535, F1: (v62 / (((v62 << ((85) & 63))) | 1))}
                                                                        _ = v62
                                                                        _ = v63
                                                                        _ = v64
                                                                    }
                                                                } else {
                                                                    if (!v9) {
                                                                        var v65 float64 = 0.0
                                                                        var v66 int64 = -1
                                                                        _ = v65
                                                                        _ = v66
                                                                    }
                                                                }
                                                                v4 += (^(v4))
                                                                var v67 uint64 = 0
                                                                _ = v67
                                                            }
                                                        case 32768:
                                                            fmt.Println(v4, v4)
                                                        case 2147483647:
                                                            if (1 >= genFn0(2147483648, true)) {
                                                                fmt.Println(fi44)
                                                                if v9 {
                                                                    fmt.Println(v9, v4)
                                                                    v9 = (!v9)
                                                                    v68 := GenS2{F0: (genFn0(690, false) > -929)}
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v69 := GenS4{F0: 515, F1: 0, F2: 1}
                                                                        _ = v69
                                                                    default:
                                                                        if (true && ((-2.25 - 0.1) > 1.5)) {
                                                                            var v70 string = ("a" + "a")
                                                                            fmt.Println(v70)
                                                                            if (4294967295 > 2) {
                                                                                v70 = v70
                                                                                if false {
                                                                                    var v71 int16 = 1000
                                                                                    fmt.Println(fi43)
                                                                                    var v72 float32 = 1e10
                                                                                    if (v9 && false) {
                                                                                        v72 = (v72 - (((3.14159 / (v72 - (-2.25 + 1e10))) / ((float32(1e10) * v72) - (v72 + -2.25))) - v72))
                                                                                        if (65535 <= 255) {
                                                                                            if v9 {
                                                                                                var v73 int8 = 1
                                                                                                v71 = 0
                                                                                                if (v9 && false) {
                                                                                                    fmt.Println(v4, v9)
                                                                                                    var v74 uint64 = 587
                                                                                                    _ = v74
                                                                                                }
                                                                                                _ = v73
                                                                                            }
                                                                                            fmt.Println(fi44, fi43, fi44)
                                                                                        } else {
                                                                                            var v75 bool = v9
                                                                                            if true {
                                                                                                v71 += (^(v71))
                                                                                                v71 += v71
                                                                                                fmt.Println(v72, v72)
                                                                                            }
                                                                                            var v76 float32 = v72
                                                                                            _ = v75
                                                                                            _ = v76
                                                                                        }
                                                                                    }
                                                                                    _ = v71
                                                                                    _ = v72
                                                                                } else {
                                                                                    var v77 int32 = 1
                                                                                    var v78 uint8 = 2
                                                                                    _ = v77
                                                                                    _ = v78
                                                                                }
                                                                            } else {
                                                                                fmt.Println(fi43, v70)
                                                                            }
                                                                            _ = v70
                                                                        } else {
                                                                            if v9 {
                                                                                v79 := GenS8{F0: 0, F1: 0, F2: v9}
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = (("foo bar" + ("" + "a")) + "a")
                                                                                }
                                                                                var v80 float32 = 3.14159
                                                                                if (1 != (^(v4))) {
                                                                                    if v9 {
                                                                                        v81 := GenS12{F0: fi43, F1: 42}
                                                                                        _ = v81
                                                                                    }
                                                                                    v3 = append(v3, ("hello" + "\n\t"))
                                                                                }
                                                                                var v82 float64 = -2.25
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    v82 += float64((v80 / ((1e10 / v80) + ((v80 + v80) * (((v80 + 0.1) - 3.14159) + float32(v82))))))
                                                                                }
                                                                                _ = v79
                                                                                _ = v80
                                                                                _ = v82
                                                                            } else {
                                                                                v9 = false
                                                                            }
                                                                        }
                                                                    }
                                                                    var v83 float32 = (0.0 + 0.0)
                                                                    _ = v68
                                                                    _ = v83
                                                                } else {
                                                                    var v84 uint64 = 753
                                                                    v4 = (fi44 | 4294967295)
                                                                    v85 := GenS10{F0: 784, F1: 42}
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                                var v86 uint32 = 639
                                                                _ = v86
                                                            } else {
                                                                switch 1 {
                                                                case -9223372036854775808:
                                                                    switch (fi44 >> ((1) & 63)) {
                                                                    case -1:
                                                                        var v87 uint = 4294967295
                                                                        _ = v87
                                                                    default:
                                                                        var v88 uint16 = 416
                                                                        _ = v88
                                                                    }
                                                                default:
                                                                    v4 = (fi44 | fi43)
                                                                }
                                                            }
                                                        default:
                                                            switch fi43 {
                                                            case 0:
                                                                var v89 uint32 = 118
                                                                _ = v89
                                                            case -325:
                                                                if false {
                                                                    v90 := GenS13{F0: 4294967295, F1: v9}
                                                                    var v91 uint8 = 76
                                                                    if ("foo bar" > "\n\t") {
                                                                        if (!(true || v9)) {
                                                                            v9 = (1 == 994)
                                                                            fmt.Println(fi44, v4, fi44)
                                                                            var v92 uint16 = 1
                                                                            switch "hello" {
                                                                            case "a":
                                                                                var v93 bool = v9
                                                                                _ = v93
                                                                            }
                                                                            var v94 int = 1
                                                                            _ = v92
                                                                            _ = v94
                                                                        }
                                                                    }
                                                                    if true {
                                                                        fmt.Println(v4, v91, fi44)
                                                                        var v95 int32 = genFn0(1, (0.1 <= 1.5))
                                                                        _ = v95
                                                                    }
                                                                    _ = v90
                                                                    _ = v91
                                                                } else {
                                                                    var v96 uint8 = 42
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        v97 := v2["hello"]
                                                                        fmt.Println(fi44, v9)
                                                                        _ = v97
                                                                    case "a":
                                                                        v4 = (fi43 % (((-(v4))) | 1))
                                                                    default:
                                                                        fmt.Println(v9, v4)
                                                                    }
                                                                    v98 := GenS10{F0: 65535, F1: 1}
                                                                    _ = v96
                                                                    _ = v98
                                                                }
                                                            }
                                                            if true {
                                                                v99 := v2[(("hello" + (("hello" + ("hello" + (("\n\t" + ("hello" + "hello")) + ("" + ("foo bar" + ""))))) + "a")) + "\n\t")]
                                                                v2[v99] = v99
                                                                _ = v99
                                                            }
                                                            var v100 int = fi44
                                                            var v101 uint8 = 86
                                                            var v102 float32 = 1e10
                                                            _ = v100
                                                            _ = v101
                                                            _ = v102
                                                        }
                                                    case "a":
                                                        v2["foo bar"] = "foo bar"
                                                        v4 -= v4
                                                        v9 = false
                                                    case "foo bar":
                                                        if (42 >= 100) {
                                                            var v103 int64 = int64(v4)
                                                            _ = v103
                                                        }
                                                    default:
                                                        _ = len(v5)
                                                    }
                                                    v104 := GenS3{F0: 42, F1: 42}
                                                    _ = v104
                                                }
                                                if v9 {
                                                    v105 := GenS14{F0: 0, F1: -84, F2: 42}
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        _ = len(v3)
                                                    case "a":
                                                        v106 := GenS8{F0: 127, F1: 42, F2: false}
                                                        var v107 int64 = 522
                                                        _ = v106
                                                        _ = v107
                                                    default:
                                                        v108 := GenS5{F0: 100, F1: 2}
                                                        _ = v108
                                                    }
                                                    _ = v105
                                                } else {
                                                    v109 := v1[(fi43 % ((fi44) | 1))]
                                                    if false {
                                                        switch ("\n\t" + "hello") {
                                                        case "hello":
                                                            fmt.Println(v4)
                                                        default:
                                                            fmt.Println(v4, v4, v4)
                                                        }
                                                        v9 = (!v109)
                                                        var v110 uint16 = 2
                                                        if v9 {
                                                            v111 := GenS2{F0: v9}
                                                            fmt.Println(fi43)
                                                            var v112 int = (fi43 * (v4 & fi44))
                                                            _ = v111
                                                            _ = v112
                                                        } else {
                                                            switch 32767 {
                                                            case -803:
                                                                var v113 float32 = 3.14159
                                                                _ = v113
                                                            default:
                                                                v114 := GenS15{F0: 1}
                                                                _ = v114
                                                            }
                                                            v109 = v9
                                                            var v115 float64 = 0.1
                                                            _ = v115
                                                        }
                                                        v4 = (-(fi43))
                                                        _ = v110
                                                    } else {
                                                        var v116 int8 = -76
                                                        var v117 uint64 = 4294967295
                                                        _ = v116
                                                        _ = v117
                                                    }
                                                    _ = v109
                                                }
                                            }
                                        }
                                        v118 := GenS16{F0: 794}
                                        _ = v118
                                    } else {
                                        var v119 bool = (0.1 < 1e10)
                                        _ = v119
                                    }
                                case 1:
                                    if false {
                                        if ((1.5 * (1.5 * 1.5)) >= float64(((-2.25 / 0.1) * -2.25))) {
                                            var v120 uint8 = 1
                                            switch (fi43 - fi44) {
                                            case 32768:
                                                var v121 int64 = -1
                                                _ = v121
                                            case 4294967295:
                                                switch "hello" {
                                                case "foo bar":
                                                    fmt.Println(fi44, fi43)
                                                case "\n\t":
                                                    if ((0.0 + 1e10) >= -2.25) {
                                                        v4 = -329
                                                        v120 = v120
                                                        var v122 uint64 = 18446744073709551615
                                                        fmt.Println(fi43, fi43, v9)
                                                        _ = v122
                                                    }
                                                    v9 = (65535 >= 4294967295)
                                                    if v9 {
                                                        if v9 {
                                                            var v123 uint16 = 65535
                                                            v124 := GenS3{F0: 654, F1: -62}
                                                            fmt.Println(v123, v4, v123)
                                                            _ = v123
                                                            _ = v124
                                                        } else {
                                                            v125 := GenS3{F0: 2, F1: 127}
                                                            delete(v2, "")
                                                            var v126 float64 = 1e10
                                                            switch fi43 {
                                                            case -1:
                                                                if false {
                                                                    var v127 float64 = (v126 - ((((3.14159 * v126) - float64((-2.25 * (((((float32((v126 / (-2.25 - (v126 / ((((-2.25 + v126) * v126) - v126) / v126))))) * ((0.1 + (1e10 / (0.1 + 0.1))) + 1e10)) / -2.25) - (1e10 - 3.14159)) + 0.0) + 1.5)))) / (v126 - -2.25)) / 1.5))
                                                                    if v9 {
                                                                        switch (fi44 - (v4 & (^(v4)))) {
                                                                        case 4294967295:
                                                                            var v128 float64 = (v127 - 1.5)
                                                                            _ = v128
                                                                        }
                                                                    } else {
                                                                        var v129 int64 = -1
                                                                        v130 := GenS4{F0: 1, F1: 1, F2: v129}
                                                                        delete(v2, "hello")
                                                                        _ = v129
                                                                        _ = v130
                                                                    }
                                                                    fmt.Println(fi43)
                                                                    _ = v127
                                                                } else {
                                                                    fmt.Println(v126)
                                                                }
                                                            case 973:
                                                                var v131 int16 = 1
                                                                _ = v131
                                                            case 32768:
                                                                fmt.Println(v126)
                                                            default:
                                                                fmt.Println(v120)
                                                            }
                                                            _ = v125
                                                            _ = v126
                                                        }
                                                        if v9 {
                                                            v9 = (2147483647 >= -2147483648)
                                                            v132 := GenS17{F0: 919}
                                                            v4 -= fi43
                                                            _ = v132
                                                        }
                                                        if v9 {
                                                            v133 := GenS18{F0: -9223372036854775808}
                                                            var v134 int64 = 9223372036854775807
                                                            switch (fi44 | 694) {
                                                            case 4294967295:
                                                                var v135 int32 = genFn0(2, (0 < -111))
                                                                _ = v135
                                                            }
                                                            _ = v133
                                                            _ = v134
                                                        } else {
                                                            var v136 uint = 236
                                                            if ((-686 > 2147483647) || ((!false) && (!("hello" > "hello")))) {
                                                                v137 := GenS12{F0: fi44, F1: 100}
                                                                _ = len(v3)
                                                                v9 = true
                                                                _ = v137
                                                            }
                                                            var v138 int8 = -128
                                                            delete(v1, 551)
                                                            _ = v136
                                                            _ = v138
                                                        }
                                                        v9 = (!(1 <= 42))
                                                        fmt.Println(v120, fi43, v4)
                                                        fmt.Println(v4)
                                                    } else {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v4 = v4
                                                        case "foo bar":
                                                            fmt.Println(v4, v9)
                                                        case "":
                                                            var v139 uint32 = 956
                                                            _ = v139
                                                        default:
                                                            _ = len(v5)
                                                        }
                                                        if ((0.1 - (1e10 * 0.1)) < (-2.25 * (0.0 + (1.5 * 0.0)))) {
                                                            switch (v4 ^ (fi44 - (fi44 * fi43))) {
                                                            case 1:
                                                                v120 = (v120 >> ((167) & 63))
                                                            case -9223372036854775808:
                                                                v140 := v2[(("\n\t" + "foo bar") + "hello")]
                                                                _ = v140
                                                            case 42:
                                                                var v141 int8 = 42
                                                                _ = v141
                                                            default:
                                                                switch ("hello" + "\n\t") {
                                                                case "":
                                                                    var v142 float64 = (float64(float32((((float64(0.1) * ((-2.25 / 1.0) - float64((-2.25 * 0.1)))) + (3.14159 / 1e10)) + ((float64((float32(((float64(1.5) - ((3.14159 + (float64(0.0) / ((3.14159 + -2.25) - 1.5))) + 0.1)) * ((float64((3.14159 * (0.0 / 3.14159))) / 1.0) - (((3.14159 - 1.5) + float64((3.14159 / float32(((1.5 - (0.1 * 0.0)) + float64(1e10)))))) / (float64(-2.25) - ((-2.25 + 0.1) - 3.14159)))))) * 3.14159)) + 0.1) - (((1.5 / 1.0) / 1.0) + 0.1))))) - -2.25)
                                                                    _ = v142
                                                                case "foo bar":
                                                                    var v143 bool = (986 > 18446744073709551615)
                                                                    _ = v143
                                                                case "a":
                                                                    v4 -= (^(v4))
                                                                default:
                                                                    v1[fi43] = v9
                                                                }
                                                            }
                                                            var v144 int = (^(fi44))
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            _ = v144
                                                        }
                                                    }
                                                default:
                                                    var v145 float64 = 0.0
                                                    _ = v145
                                                }
                                            case -1:
                                                v146 := GenS19{F0: v120}
                                                _ = v146
                                            default:
                                                if v9 {
                                                    if false {
                                                        var v147 float32 = float32(3.14159)
                                                        _ = v147
                                                    }
                                                    v120 = v120
                                                } else {
                                                    if (-128 == 0) {
                                                        v120 = (v120 >> ((uint(v120)) & 63))
                                                        v148 := GenS20{F0: 100, F1: 1, F2: -128}
                                                        _ = v148
                                                    }
                                                    fmt.Println(fi43, fi43, v9)
                                                    v120 = v120
                                                }
                                            }
                                            _ = v120
                                        } else {
                                            var v149 bool = (!v9)
                                            _ = v149
                                        }
                                        v150 := GenS19{F0: 127}
                                        v151 := GenS19{F0: 9}
                                        _ = v150
                                        _ = v151
                                    } else {
                                        fmt.Println(fi44, fi43, fi43)
                                    }
                                }
                                var v152 uint64 = 2
                                switch "a" {
                                case "\n\t":
                                    if (v9 && true) {
                                        var v153 int16 = 42
                                        var v154 uint64 = v152
                                        _ = v153
                                        _ = v154
                                    }
                                case "foo bar":
                                    if len(v3) > 0 {
                                        v3[0] = ((("hello" + ("\n\t" + "foo bar")) + "foo bar") + ("a" + ((("a" + ("hello" + "hello")) + "foo bar") + "foo bar")))
                                    }
                                default:
                                    v9 = false
                                }
                                _ = v152
                            }
                        }
                    }
                    if v9 {
                        if (!(v9 && ((541 < 2) && (2 > 2)))) {
                            if (255 >= 225) {
                                var v155 string = "foo bar"
                                _ = v155
                            }
                        }
                        if (!((false && (2 == 1)) || false)) {
                            v9 = (2147483647 <= 42)
                            v156 := GenS21{F0: (!v9), F1: 34, F2: 65535}
                            v157 := GenS22{F0: 2147483647, F1: 624, F2: 42}
                            _ = v156
                            _ = v157
                        }
                    } else {
                        v158 := GenS14{F0: 128, F1: 21, F2: 387}
                        _ = v158
                    }
                    if v9 {
                        var v159 float32 = 0.0
                        if (32767 > 255) {
                            switch -2147483648 {
                            case -1:
                                var v160 float32 = 1e10
                                _ = v160
                            case -2147483648:
                                v2[("foo bar" + (("foo bar" + "foo bar") + ""))] = "a"
                                if (0 == 4294967295) {
                                    fmt.Println(fi43, v4)
                                }
                                var v161 float64 = 1.5
                                _ = v161
                            case 2147483647:
                                var v162 uint64 = 956
                                _ = v162
                            default:
                                v163 := GenS23{F0: (v4 * (fi44 / ((fi43) | 1))), F1: 1, F2: -1}
                                _ = v163
                            }
                            if (("" + (("a" + "") + "foo bar")) == "hello") {
                                if (!false) {
                                    _ = len(v5)
                                } else {
                                    fmt.Println(v9)
                                }
                            }
                        } else {
                            v164 := GenS13{F0: 139, F1: (true || v9)}
                            switch (("" + "") + "hello") {
                            case "":
                                if (32767 != 759) {
                                    var v165 int16 = 367
                                    var v166 int64 = 42
                                    switch 1 {
                                    case 32768:
                                        fmt.Println(fi44, v166, v4)
                                    case 42:
                                        if (-2147483648 > 42) {
                                            v167 := GenS2{F0: (!((float32((0.1 - (1e10 * -2.25))) <= (v159 / v159)) || false))}
                                            _ = v167
                                        } else {
                                            v168 := GenS24{F0: (("\n\t" + "a") == ""), F1: ("foo bar" + ("\n\t" + "")), F2: 0}
                                            v9 = v9
                                            v3 = append(v3, "")
                                            _ = v168
                                        }
                                    case 128:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    _ = v165
                                    _ = v166
                                }
                            }
                            v159 = v159
                            _ = len(v3)
                            if v9 {
                                var v169 float32 = v159
                                v170 := GenS25{F0: 42}
                                v159 -= (v169 / float32(float64(((0.1 + (v159 * float32(1.5))) - 0.0))))
                                var v171 uint16 = 629
                                _ = v169
                                _ = v170
                                _ = v171
                            }
                            _ = v164
                        }
                        _ = v159
                    } else {
                        switch 9223372036854775807 {
                        case 756:
                            fmt.Println(v9, v9)
                        case 333:
                            var v172 uint8 = 1
                            _ = v172
                        case 1:
                            var v173 int16 = 32767
                            _ = v173
                        }
                        var v174 float32 = 1e10
                        _ = v174
                    }
                    _ = v47
                }
                _ = fi44
            }
            _ = fi43
        }
    } else {
        v4 = (v4 >> ((799) & 63))
        for fi175 := 0; fi175 < 16; fi175++ {
            v176 := v1[42]
            _ = fi175
            _ = v176
        }
        var v177 string = ("a" + "\n\t")
        _ = v177
    }
    delete(v1, (^(v4)))
    var v178 string = "\n\t"
    v3 = append(v3, (v178 + v178))
    v4 -= v4
    v9 = (42 < 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v178:", v178)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v178
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
