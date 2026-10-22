package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int; F1 int8; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 string; F2 uint16 }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int; F1 uint64 }
type GenS13 struct { F0 int8; F1 int32 }
type GenS14 struct { F0 uint8; F1 int }
type GenS15 struct { F0 int16; F1 int32; F2 uint8 }
type GenS16 struct { F0 uint64; F1 int; F2 int64 }
type GenS17 struct { F0 bool; F1 string }
func genFn0(p0 int16, p1 string, p2 string) uint8 {
    if true {
        fmt.Println(p0)
    }
    var v5 int8 = 115
    switch ((("a" + p2) + (p1 + "hello")) + "") {
    case "hello":
        v5 += (v5 / ((100) | 1))
    case "a":
        p0 = (-(p0))
    default:
        switch -951 {
        case 4294967295:
            if (((!true) || (true || true)) && (2147483647 < 234)) {
                for fi6 := 0; fi6 < 8; fi6++ {
                    v5 += (v5 - -55)
                    if (-9223372036854775808 > -1) {
                        var v7 uint16 = 0
                        var v8 bool = (false || false)
                        _ = v7
                        _ = v8
                    }
                    _ = fi6
                }
                v5 -= (v5 | v5)
                for fi9 := 0; fi9 < 7; fi9++ {
                    v10 := GenS1{F0: 2147483647, F1: false, F2: (fi9 * 9223372036854775807)}
                    switch (p1 + p1) {
                    case "foo bar":
                        if (true || (true && (((p0 + p0) == p0) && ((p1 + (p2 + p2)) == p2)))) {
                            if (78 < 234) {
                                if false {
                                    v11 := []int{(fi9 & 9223372036854775807), 2147483647, (fi9 >> ((0) & 63)), 300}
                                    if (fi9 == 32767) {
                                        var v12 float64 = 3.14159
                                        var v13 uint8 = 42
                                        _ = v12
                                        _ = v13
                                    } else {
                                        v14 := GenS2{F0: 9223372036854775807}
                                        var v15 uint8 = 42
                                        _ = v14
                                        _ = v15
                                    }
                                    _ = v11
                                } else {
                                    var v16 string = p2
                                    var v17 float64 = 3.14159
                                    _ = v16
                                    _ = v17
                                }
                            }
                            var v18 float64 = (((((1e10 * (0.0 / (1e10 * (((0.1 + 3.14159) - 0.1) + -2.25)))) / 1e10) + 0.0) / 0.1) + 0.1)
                            var v19 uint32 = 65535
                            _ = v18
                            _ = v19
                        }
                    case "a":
                        var v20 uint = 65535
                        _ = v20
                    default:
                        var v21 int16 = (^(p0))
                        _ = v21
                    }
                    p2 = ""
                    _ = fi9
                    _ = v10
                }
            }
        case 42:
            if (!true) {
                fmt.Println(v5, p1)
                var v22 float64 = 0.0
                _ = v22
            }
        default:
            var v23 int16 = p0
            _ = v23
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 10
}
func genFn1(p0 int8, p1 string) uint {
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 int16 = 1000
        p0 = (-(p0))
        p1 = "hello"
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        p0 = 0
        var v27 int32 = -2147483648
        if (false || (-202 < 0)) {
            switch (fi26 + fi26) {
            case 32768:
                v28 := []int{fi26, fi26, (^(fi26)), fi26}
                _ = v28
            default:
                var v29 float64 = (0.1 / 1.0)
                _ = v29
            }
            if true {
                var v30 uint64 = 61
                switch fi26 {
                case 42:
                    fmt.Println(p0, v27, p0)
                case 32767:
                    switch 9223372036854775807 {
                    case 42:
                        v27 = v27
                    case 0:
                        v30 = v30
                    }
                case 78:
                    var v31 int = fi26
                    _ = v31
                default:
                    v32 := GenS3{F0: 65535, F1: 338}
                    _ = v32
                }
                _ = v30
            }
        } else {
            var v33 bool = true
            var v34 int8 = p0
            _ = v33
            _ = v34
        }
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    return 369
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint64 = 42
    var v3 int8 = 127
    var v4 int = 150
    v4 -= int(v3)
    var v35 uint8 = 189
    switch "a" {
    case "":
        v1 -= 3.14159
    case "foo bar":
        var v36 string = "\n\t"
        _ = v36
    case "hello":
        if false {
            if true {
                var v37 uint64 = (^(v2))
                _ = v37
            }
            if true {
                v3 += v3
                fmt.Println(v3, v2, v1)
            }
        }
    }
    fmt.Println(v4, v4, v4)
    var v38 int16 = -1
    v4 -= -421
    if false {
        switch -2147483648 {
        case 0:
            for fi39 := 0; fi39 < 17; fi39++ {
                var v40 uint = 487
                for fi41 := 0; fi41 < 2; fi41++ {
                    v38 = 1000
                    v42 := []int{}
                    _ = fi41
                    _ = v42
                }
                for fi43 := 0; fi43 < 12; fi43++ {
                    var v44 string = ""
                    var v45 uint64 = (v2 % ((v2) | 1))
                    v46 := []int{fi43, (v4 >> ((1) & 63)), (fi39 ^ fi43)}
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v46
                }
                _ = fi39
                _ = v40
            }
        default:
            v3 += v3
        }
    } else {
        switch v4 {
        case 128:
            v1 += 0.1
        case 127:
            switch int(v38) {
            case 32768:
                switch v4 {
                case 237:
                    v47 := []int{9223372036854775807, (-(v4)), 128, v4}
                    _ = v47
                case 2147483647:
                    var v48 string = "a"
                    _ = v48
                }
            case 32767:
                for fi49 := 0; fi49 < 6; fi49++ {
                    v50 := []int{v4}
                    fmt.Println(fi49)
                    for fi51 := 0; fi51 < 4; fi51++ {
                        v52 := GenS2{F0: -1}
                        if false {
                            v2 += v2
                            if false {
                                if false {
                                    if ((!false) && true) {
                                        v53 := GenS4{F0: fi49, F1: -1, F2: 2147483648}
                                        fmt.Println(v4, fi49)
                                        _ = v53
                                    }
                                    switch ("\n\t" + (("\n\t" + "a") + (("\n\t" + "foo bar") + ("hello" + ("\n\t" + ""))))) {
                                    case "hello":
                                        switch v4 {
                                        case 32767:
                                            if len(v50) > 0 {
                                                v50[0] = (-(fi49))
                                            }
                                        case -1:
                                            fmt.Println(v2, v4)
                                        default:
                                            if len(v50) >= 2 {
                                                v50 = v50[1:len(v50)]
                                            }
                                        }
                                    }
                                    switch ("\n\t" + "hello") {
                                    case "hello":
                                        var v54 uint = 4294967295
                                        _ = v54
                                    case "a":
                                        var v55 int16 = v38
                                        _ = v55
                                    case "foo bar":
                                        v56 := GenS5{F0: v2}
                                        _ = v56
                                    default:
                                        switch (fi51 / ((fi49) | 1)) {
                                        case 326:
                                            v57 := GenS6{F0: (!(v38 < (v38 * (^(v38)))))}
                                            _ = v57
                                        case 42:
                                            if true {
                                                v35 = 128
                                                v58 := GenS7{F0: (false && (!true)), F1: "foo bar", F2: 369}
                                                _ = v58
                                            }
                                        default:
                                            if (v3 == 100) {
                                                if len(v50) > 0 {
                                                    v50[0] = 32767
                                                }
                                            } else {
                                                var v59 int16 = (v38 / ((42) | 1))
                                                if (int64(v35) < 0) {
                                                    fmt.Println(v3, v4)
                                                }
                                                _ = v59
                                            }
                                        }
                                    }
                                }
                                if true {
                                    v4 -= (fi49 << ((382) & 63))
                                } else {
                                    v2 = (v2 & v2)
                                    v38 -= (v38 | (v38 % (((-(v38))) | 1)))
                                    v50 = append(v50, fi49)
                                }
                                var v60 uint = 42
                                _ = v60
                            } else {
                                var v61 bool = true
                                v4 = -1
                                _ = v61
                            }
                        } else {
                            v50 = append(v50, fi51)
                            v62 := GenS8{F0: genFn1(-128, "hello"), F1: 1}
                            switch "" {
                            case "":
                                switch ("\n\t" + "hello") {
                                case "a":
                                    v63 := GenS8{F0: 952, F1: 1}
                                    _ = v63
                                case "":
                                    if (2147483647 < 0) {
                                        v4 -= -2147483648
                                        v35 -= (v35 << ((0) & 63))
                                        v64 := GenS3{F0: 0, F1: 2}
                                        v50 = append(v50, fi49)
                                        _ = v64
                                    } else {
                                        v65 := GenS6{F0: false}
                                        _ = v65
                                    }
                                case "hello":
                                    var v66 int64 = 9223372036854775807
                                    _ = v66
                                }
                            case "\n\t":
                                if len(v50) >= 2 {
                                    v50 = v50[1:len(v50)]
                                }
                            }
                            _ = v62
                        }
                        _ = fi51
                        _ = v52
                    }
                    _ = fi49
                    _ = v50
                }
            case 0:
                switch (v4 % ((v4) | 1)) {
                case -557:
                    var v67 bool = true
                    _ = v67
                }
            }
        case -2147483648:
            if (!(42 > 42)) {
                for fi68 := 0; fi68 < 6; fi68++ {
                    v69 := []int{(fi68 >> ((517) & 63)), v4, 1}
                    fmt.Println(v2, v2)
                    _ = fi68
                    _ = v69
                }
                v70 := []int{}
                _ = v70
            } else {
                switch "" {
                case "foo bar":
                    for fi71 := 0; fi71 < 8; fi71++ {
                        v1 = (v1 - ((1.5 + v1) / 1.0))
                        for fi72 := 0; fi72 < 14; fi72++ {
                            var v73 int32 = 619
                            if (!true) {
                                if true {
                                    var v74 int = (fi72 - (fi71 | fi72))
                                    if true {
                                        var v75 uint = 511
                                        _ = v75
                                    } else {
                                        if (!true) {
                                            v2 = v2
                                            var v76 string = ("\n\t" + "\n\t")
                                            var v77 bool = false
                                            _ = v76
                                            _ = v77
                                        }
                                        if false {
                                            switch ("a" + ("hello" + "")) {
                                            case "hello":
                                                if false {
                                                    v74 = (v4 << ((18446744073709551615) & 63))
                                                    v78 := GenS9{F0: 2}
                                                    if true {
                                                        var v79 float32 = 0.1
                                                        _ = v79
                                                    }
                                                    v80 := []int{}
                                                    var v81 int32 = 348
                                                    _ = v78
                                                    _ = v80
                                                    _ = v81
                                                }
                                            case "":
                                                var v82 int32 = (-(v73))
                                                _ = v82
                                            }
                                            v35 = v35
                                        }
                                        v3 = v3
                                    }
                                    v83 := GenS10{F0: 1, F1: (v35 & (v35 << ((65535) & 63))), F2: v73}
                                    _ = v74
                                    _ = v83
                                } else {
                                    v84 := []int{-9223372036854775808}
                                    var v85 int32 = v73
                                    var v86 int64 = -373
                                    var v87 int64 = (v86 - v86)
                                    _ = v84
                                    _ = v85
                                    _ = v86
                                    _ = v87
                                }
                                var v88 int8 = (-(v3))
                                v2 -= (v2 & (v2 << ((909) & 63)))
                                _ = v88
                            }
                            fmt.Println(fi72)
                            _ = fi72
                            _ = v73
                        }
                        if (((v35 | 222) != (v35 * v35)) && true) {
                            var v89 bool = true
                            _ = v89
                        }
                        _ = fi71
                    }
                case "\n\t":
                    if true {
                        if (v35 != genFn0(v38, "foo bar", "")) {
                            if true {
                                v90 := GenS11{F0: -452}
                                var v91 uint32 = 2147483647
                                var v92 int32 = -57
                                _ = v90
                                _ = v91
                                _ = v92
                            }
                            if false {
                                switch int(v3) {
                                case 644:
                                    if ((42 >= 928) || true) {
                                        v93 := GenS12{F0: v4, F1: v2}
                                        switch (v4 ^ v4) {
                                        case -455:
                                            v2 = v2
                                        default:
                                            v94 := GenS13{F0: v3, F1: 2147483647}
                                            _ = v94
                                        }
                                        _ = v93
                                    } else {
                                        v95 := GenS8{F0: 65535, F1: 255}
                                        v2 = 530
                                        _ = v95
                                    }
                                case 479:
                                    var v96 uint = genFn1(v3, (("hello" + ("hello" + ("\n\t" + ("hello" + "")))) + "foo bar"))
                                    _ = v96
                                case 128:
                                    if (("\n\t" + ((("\n\t" + "a") + ("\n\t" + ("foo bar" + ""))) + "foo bar")) >= "") {
                                        v97 := []int{(-(v4)), v4, (v4 & 32768), v4}
                                        var v98 uint64 = v2
                                        var v99 uint32 = 65535
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    } else {
                                        for fi100 := 0; fi100 < 15; fi100++ {
                                            var v101 uint16 = uint16(v2)
                                            _ = fi100
                                            _ = v101
                                        }
                                        var v102 uint16 = 212
                                        _ = v102
                                    }
                                }
                                fmt.Println(v1, v3, v35)
                                v103 := GenS8{F0: 877, F1: 32767}
                                _ = v103
                            } else {
                                v104 := []int{979}
                                if false {
                                    if true {
                                        if len(v104) >= 2 {
                                            v104 = v104[1:len(v104)]
                                        }
                                    }
                                } else {
                                    _ = len(v104)
                                }
                                _ = v104
                            }
                            if (1 == 65535) {
                                switch (v4 * (v4 * v4)) {
                                case 128:
                                    var v105 uint = 2
                                    _ = v105
                                case -2147483648:
                                    var v106 float64 = (v1 + v1)
                                    _ = v106
                                default:
                                    v107 := GenS2{F0: 9223372036854775807}
                                    _ = v107
                                }
                            }
                        }
                    }
                }
                fmt.Println(v4, v4)
            }
        }
    }
    var v108 int = 2147483647
    if false {
        var v109 uint32 = 42
        _ = v109
    }
    v108 -= v4
    switch v108 {
    case 32768:
        v35 = uint8(v2)
    default:
        v110 := GenS14{F0: (v35 >> ((4294967295) & 63)), F1: v108}
        _ = v110
    }
    var v111 uint16 = 255
    for fi112 := 0; fi112 < 16; fi112++ {
        v113 := GenS10{F0: 723, F1: v35, F2: -251}
        _ = fi112
        _ = v113
    }
    switch (v4 * 836) {
    case 127:
        for fi114 := 0; fi114 < 10; fi114++ {
            v115 := GenS15{F0: v38, F1: 42, F2: v35}
            var v116 uint8 = genFn0(v38, "a", (("" + "a") + "\n\t"))
            _ = fi114
            _ = v115
            _ = v116
        }
    case -324:
        v3 = (v3 + 42)
    case 35:
        v1 -= (((-2.25 + v1) + v1) - v1)
    default:
        var v117 string = "hello"
        _ = v117
    }
    if ((v111 + (v111 ^ uint16(v2))) != (v111 & v111)) {
        v118 := GenS7{F0: false, F1: ("a" + "\n\t"), F2: 65535}
        if (1 != -730) {
            v119 := GenS16{F0: (^(v2)), F1: int(v3), F2: 9223372036854775807}
            _ = v119
        } else {
            switch -119 {
            case 127:
                v120 := []int{-178, v4}
                _ = v120
            case -9223372036854775808:
                v121 := []int{127, (v108 & int(v2)), (v108 >> ((0) & 63))}
                _ = v121
            case 2147483647:
                v122 := GenS8{F0: 4294967295, F1: v111}
                _ = v122
            }
            v123 := GenS17{F0: (1 < -2147483648), F1: ("" + ("a" + ""))}
            _ = v123
        }
        _ = v118
    }
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v108
    _ = v111
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v4
}
