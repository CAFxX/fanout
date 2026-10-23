package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 uint32; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int16; F2 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS10 struct { F0 uint; F1 uint8 }
type GenS11 struct { F0 int8; F1 int8 }
type GenS12 struct { F0 uint64; F1 int32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 bool; F1 int }
type GenS16 struct { F0 string; F1 uint32 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint8; F1 int64 }
type GenS19 struct { F0 int; F1 uint32 }
type GenS20 struct { F0 string }
type GenS21 struct { F0 int8; F1 int8 }
type GenS22 struct { F0 int16; F1 int16 }
type GenS23 struct { F0 int32; F1 int8 }
type GenS24 struct { F0 uint32; F1 int8; F2 int64 }
type GenS25 struct { F0 uint32; F1 string }
type GenS26 struct { F0 uint8; F1 string; F2 string }
type GenS27 struct { F0 uint8; F1 int16 }
type GenS28 struct { F0 uint64 }
func genFn0(p0 int, p1 int16) string {
    switch 128 {
    case 128:
        var v5 float64 = 0.1
        _ = v5
    case -9223372036854775808:
        if (!(true || true)) {
            switch -575 {
            case -932:
                p0 = p0
            case -373:
                p1 = (^(p1))
            }
            switch ("" + "foo bar") {
            case "foo bar":
                for fi6 := 0; fi6 < 10; fi6++ {
                    p1 = -776
                    var v7 bool = (-2147483648 == 0)
                    _ = fi6
                    _ = v7
                }
            case "":
                p0 = p0
            default:
                v8 := GenS1{F0: 72, F1: 42, F2: 80}
                _ = v8
            }
        } else {
            fmt.Println(p0)
        }
    default:
        var v9 uint32 = 2147483648
        _ = v9
    }
    var v10 uint64 = 0
    for fi11 := 0; fi11 < 5; fi11++ {
        p0 = p0
        p1 += p1
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return "hello"
}
func genFn1(p0 int8, p1 int64) int32 {
    var v12 int16 = 995
    fmt.Println(v12)
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS1{F0: 2147483648, F1: -372, F2: 0}
        switch "foo bar" {
        case "hello":
            v15 := GenS2{F0: ("a" + ("" + "a")), F1: 65535, F2: "foo bar"}
            _ = v15
        case "a":
            if (583 <= 18446744073709551615) {
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 uint16 = 42
                    if true {
                        switch "" {
                        case "\n\t":
                            if false {
                                if ((fi13 | (fi13 % (((fi13 / (((fi13 % ((-1) | 1))) | 1))) | 1))) < fi16) {
                                    v18 := GenS3{F0: 32767}
                                    v19 := []int{int(v12), -179, fi13, (fi13 << ((4294967295) & 63))}
                                    _ = v18
                                    _ = v19
                                } else {
                                    var v20 string = "hello"
                                    var v21 int64 = p1
                                    _ = v20
                                    _ = v21
                                }
                                v22 := []int{}
                                _ = v22
                            } else {
                                var v23 bool = true
                                if (true && false) {
                                    v12 = (-(v12))
                                    v17 += v17
                                }
                                _ = v23
                            }
                        default:
                            var v24 float32 = -2.25
                            fmt.Println(fi13, v24, v17)
                            _ = v24
                        }
                        p0 = (p0 * (^(p0)))
                        var v25 int32 = -1
                        _ = v25
                    }
                    var v26 string = "hello"
                    _ = fi16
                    _ = v17
                    _ = v26
                }
                switch 32767 {
                case -606:
                    switch ("" + "") {
                    case "a":
                        v27 := []int{(fi13 | (fi13 << ((303) & 63)))}
                        _ = v27
                    }
                default:
                    v28 := []int{900, (^(fi13)), fi13, (fi13 << ((4294967295) & 63))}
                    _ = v28
                }
            }
        }
        _ = fi13
        _ = v14
    }
    v29 := GenS4{F0: 128}
    fmt.Println(p0, p1, v12)
    _ = p0
    _ = p1
    _ = v12
    _ = v29
    return -2147483648
}

func main() {
    var v1 string = ("hello" + ("hello" + "foo bar"))
    var v2 string = v1
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    v2 = "\n\t"
    v3[v1] = v1
    if true {
        v30 := GenS2{F0: (v1 + (v1 + (v2 + "hello"))), F1: 2147483647, F2: "a"}
        for fi31 := 0; fi31 < 11; fi31++ {
            var v32 uint32 = 994
            v1 = v1
            v33 := GenS5{F0: 127}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = v30
    } else {
        v34 := GenS6{F0: 1, F1: 32767, F2: -434}
        delete(v3, (((v1 + v1) + v1) + ""))
        _ = v34
    }
    var v35 uint = 0
    v1 = (v1 + ("a" + (((((v2 + "\n\t") + v2) + v1) + "hello") + ("hello" + "hello"))))
    for fi36 := 0; fi36 < 3; fi36++ {
        for fi37 := 0; fi37 < 7; fi37++ {
            fmt.Println(fi37)
            if ((582 == 618) || false) {
                var v38 int64 = -9223372036854775808
                switch "" {
                case "hello":
                    if false {
                        if false {
                            v38 = (v38 + (v38 - (v38 % ((v38) | 1))))
                        } else {
                            v39 := []int{(fi37 * (fi37 >> ((v35) & 63)))}
                            v4["\n\t"] = 65535
                            _ = v39
                        }
                        if false {
                            fmt.Println(v1)
                            v1 = v1
                            v40 := []int{}
                            _ = v40
                        }
                    }
                case "foo bar":
                    switch ((v2 + v1) + ("hello" + "foo bar")) {
                    case "a":
                        v2 = "\n\t"
                    case "":
                        if true {
                            switch fi36 {
                            case 804:
                                v41 := GenS2{F0: (v2 + v2), F1: 807, F2: (v1 + ((v2 + "\n\t") + v1))}
                                _ = v41
                            }
                        }
                    default:
                        v42 := []int{fi36, fi36, 4294967295}
                        _ = v42
                    }
                default:
                    if true {
                        var v43 int = fi36
                        switch ("\n\t" + v2) {
                        case "a":
                            if false {
                                var v44 bool = (65535 <= 42)
                                var v45 uint8 = 127
                                _ = v44
                                _ = v45
                            } else {
                                var v46 string = v2
                                var v47 int64 = 9223372036854775807
                                if (0 <= 2147483647) {
                                    if true {
                                        var v48 int8 = -111
                                        switch (fi36 * 2147483647) {
                                        case 9223372036854775807:
                                            v49 := GenS7{F0: uint64(v35)}
                                            _ = v49
                                        case -270:
                                            v47 = v47
                                        case -9223372036854775808:
                                            var v50 bool = (!((^(v47)) > v38))
                                            _ = v50
                                        }
                                        _ = v48
                                    } else {
                                        if false {
                                            if true {
                                                _ = len(v4)
                                            } else {
                                                var v51 float64 = ((0.0 + ((3.14159 / 1.0) / 1e10)) * 0.1)
                                                if false {
                                                    fmt.Println(v2)
                                                    switch -1 {
                                                    case 1:
                                                        fmt.Println(v35, v46, fi36)
                                                    case 834:
                                                        v1 = v1
                                                    }
                                                    if true {
                                                        v38 += -537
                                                    } else {
                                                        v46 = v1
                                                    }
                                                } else {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v52 := GenS8{F0: 565}
                                                        _ = v52
                                                    case "foo bar":
                                                        var v53 int = v43
                                                        _ = v53
                                                    case "hello":
                                                        v46 = ("a" + v2)
                                                    default:
                                                        v47 = v38
                                                    }
                                                    var v54 uint = 42
                                                    _ = v54
                                                }
                                                var v55 uint32 = 2147483648
                                                v3[v46] = (v2 + v1)
                                                _ = v51
                                                _ = v55
                                            }
                                        }
                                        if (v38 <= (v47 | v47)) {
                                            var v56 uint32 = 857
                                            v57 := GenS8{F0: 42}
                                            _ = v56
                                            _ = v57
                                        } else {
                                            var v58 int8 = 25
                                            var v59 string = (((v1 + (("hello" + "hello") + v2)) + v1) + "")
                                            _ = v58
                                            _ = v59
                                        }
                                    }
                                    fmt.Println(v35, v35, v46)
                                    if (4294967295 != 503) {
                                        var v60 int8 = -110
                                        v43 -= (fi37 % ((42) | 1))
                                        _ = v60
                                    } else {
                                        v61 := GenS2{F0: (v2 + (v46 + (((((((v2 + v46) + "a") + (("hello" + ("hello" + ((((v2 + v1) + v1) + v1) + "a"))) + v46)) + (v2 + v2)) + (v2 + (("" + ((((v46 + v1) + ((v1 + (("hello" + v46) + "\n\t")) + "")) + ((v1 + (v1 + "a")) + v46)) + v2)) + v46))) + ("\n\t" + ("\n\t" + v46))) + ((v2 + ("foo bar" + v46)) + v1)))), F1: 1, F2: v46}
                                        _ = v61
                                    }
                                }
                                _ = v46
                                _ = v47
                            }
                        case "":
                            var v62 uint = (v35 << (((v35 >> (((v35 / ((v35) | 1))) & 63))) & 63))
                            _ = v62
                        }
                        v63 := GenS9{F0: 255, F1: 732, F2: 1}
                        fmt.Println(v1, v35, v35)
                        _ = v43
                        _ = v63
                    } else {
                        v35 -= v35
                        fmt.Println(v38, v1)
                        if (v38 > (-(v38))) {
                            if true {
                                var v64 uint16 = 32767
                                switch v2 {
                                case "foo bar":
                                    v38 = (v38 / ((v38) | 1))
                                case "a":
                                    if true {
                                        v64 = (v64 | (^(v64)))
                                        fmt.Println(v1, fi37, v64)
                                        if ((!(1e10 < 3.14159)) || (v35 != v35)) {
                                            switch ("" + (v2 + "hello")) {
                                            case "\n\t":
                                                if ((^(v38)) <= 165) {
                                                    v65 := GenS2{F0: v2, F1: 2147483647, F2: "foo bar"}
                                                    v66 := []int{128, 950}
                                                    var v67 bool = (240 > 1)
                                                    if ((^(v64)) != (v64 ^ (^(v64)))) {
                                                        fmt.Println(v35)
                                                    } else {
                                                        fmt.Println(fi37, fi36, v67)
                                                    }
                                                    _ = v65
                                                    _ = v66
                                                    _ = v67
                                                }
                                            case "a":
                                                if (v2 >= "hello") {
                                                    v68 := GenS4{F0: 100}
                                                    if true {
                                                        if (-2.25 != (0.1 + (3.14159 + 0.0))) {
                                                            fmt.Println(v2)
                                                            switch (v2 + "\n\t") {
                                                            case "a":
                                                                if true {
                                                                    v1 = ("a" + "\n\t")
                                                                    fmt.Println(v64, v38, v35)
                                                                    v69 := []int{(-(fi37))}
                                                                    _ = v69
                                                                } else {
                                                                    switch (v2 + v2) {
                                                                    case "foo bar":
                                                                        if (!true) {
                                                                            switch ("foo bar" + (v2 + ("hello" + ("\n\t" + v1)))) {
                                                                            case "\n\t":
                                                                                fmt.Println(fi37, v38, v38)
                                                                            default:
                                                                                v35 += 1
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        if (127 > 246) {
                                                                            var v70 uint16 = 2
                                                                            _ = v70
                                                                        } else {
                                                                            switch (fi36 >> (((v35 >> ((v35) & 63))) & 63)) {
                                                                            case -9223372036854775808:
                                                                                switch 32767 {
                                                                                case 32767:
                                                                                    if (!(false || (3.14159 <= (3.14159 * 3.14159)))) {
                                                                                        var v71 uint16 = 65535
                                                                                        fmt.Println(v38)
                                                                                        _ = v71
                                                                                    }
                                                                                case -374:
                                                                                    var v72 string = "a"
                                                                                    _ = v72
                                                                                case 42:
                                                                                    fmt.Println(v38, fi36)
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    switch v2 {
                                                                    case "a":
                                                                        var v73 int16 = -32768
                                                                        _ = v73
                                                                    case "hello":
                                                                        var v74 int64 = v38
                                                                        _ = v74
                                                                    case "\n\t":
                                                                        v3[(v1 + v1)] = (("hello" + "hello") + ("hello" + "a"))
                                                                    }
                                                                    switch (fi36 + (fi36 >> (((v35 / ((v35) | 1))) & 63))) {
                                                                    case 9223372036854775807:
                                                                        var v75 uint64 = 359
                                                                        _ = v75
                                                                    default:
                                                                        delete(v3, v2)
                                                                    }
                                                                }
                                                            default:
                                                                v76 := GenS3{F0: (fi37 ^ 781)}
                                                                _ = v76
                                                            }
                                                        }
                                                        var v77 int64 = v38
                                                        v78 := []int{(-(fi36)), -2147483648, (fi37 - 441), (fi37 >> (((v35 ^ 770)) & 63))}
                                                        _ = v77
                                                        _ = v78
                                                    } else {
                                                        v2 = v1
                                                    }
                                                    _ = v68
                                                }
                                            case "hello":
                                                v4[""] = 2147483648
                                            }
                                            if (-1 == 168) {
                                                v38 = v38
                                                v2 = "a"
                                                v64 = (v64 / (((v64 ^ v64)) | 1))
                                            } else {
                                                if false {
                                                    var v79 int32 = -2147483648
                                                    v80 := GenS10{F0: 65535, F1: 128}
                                                    v81 := []int{-1, 0}
                                                    _ = v79
                                                    _ = v80
                                                    _ = v81
                                                }
                                            }
                                            fmt.Println(v38, v64)
                                        } else {
                                            switch v1 {
                                            case "foo bar":
                                                var v82 int8 = 1
                                                v83 := []int{fi37, fi36, fi36}
                                                fmt.Println(fi37)
                                                _ = v82
                                                _ = v83
                                            case "hello":
                                                v2 = v2
                                            }
                                            var v84 string = ("\n\t" + (v1 + v2))
                                            var v85 uint = v35
                                            if false {
                                                var v86 int64 = v38
                                                var v87 bool = ((true || (!true)) || false)
                                                _ = v86
                                                _ = v87
                                            } else {
                                                if true {
                                                    var v88 float64 = (0.1 / float64(0.1))
                                                    switch (-(fi37)) {
                                                    case 127:
                                                        var v89 uint16 = v64
                                                        _ = v89
                                                    }
                                                    _ = v88
                                                }
                                                switch fi37 {
                                                case -2147483648:
                                                    var v90 uint8 = 127
                                                    v2 = "\n\t"
                                                    _ = v90
                                                case 4294967295:
                                                    var v91 int32 = genFn1(-64, v38)
                                                    _ = v91
                                                }
                                                if false {
                                                    v92 := GenS11{F0: 42, F1: -128}
                                                    _ = v92
                                                }
                                                if (-911 >= -44) {
                                                    var v93 float64 = (3.14159 / 3.14159)
                                                    v94 := []int{(fi37 % ((fi36) | 1))}
                                                    _ = v93
                                                    _ = v94
                                                }
                                            }
                                            var v95 uint16 = (v64 | (v64 / (((v64 + 32767)) | 1)))
                                            v1 = v84
                                            v96 := GenS12{F0: 18446744073709551615, F1: -1}
                                            var v97 int16 = 711
                                            v98 := GenS13{F0: v38}
                                            v64 -= v64
                                            switch fi37 {
                                            case 32767:
                                                v1 = v2
                                            case 32768:
                                                switch ("" + v2) {
                                                case "hello":
                                                    var v99 int8 = 42
                                                    _ = v99
                                                case "\n\t":
                                                    var v100 uint64 = 80
                                                    if (1.5 == (0.0 - float64(3.14159))) {
                                                        if false {
                                                            v2 = v84
                                                            if false {
                                                                v101 := GenS14{F0: uint8(v100)}
                                                                switch v1 {
                                                                case "hello":
                                                                    v102 := v4[v84]
                                                                    v103 := []int{(fi37 | fi37), fi37, int(v38), fi36}
                                                                    v104 := GenS11{F0: 37, F1: 1}
                                                                    _ = v102
                                                                    _ = v103
                                                                    _ = v104
                                                                }
                                                                if ((v84 <= "\n\t") && true) {
                                                                    v64 = (^(v64))
                                                                    var v105 int8 = -48
                                                                    v106 := []int{0, 32767, 4294967295, 9223372036854775807}
                                                                    var v107 bool = (!(1.5 > float32(1e10)))
                                                                    switch fi36 {
                                                                    case -9223372036854775808:
                                                                        var v108 float32 = float32(1e10)
                                                                        _ = v108
                                                                    case 22:
                                                                        v85 = (^(v85))
                                                                    case 128:
                                                                        v109 := GenS3{F0: 164}
                                                                        _ = v109
                                                                    }
                                                                    _ = v105
                                                                    _ = v106
                                                                    _ = v107
                                                                }
                                                                _ = v101
                                                            } else {
                                                                v85 = v35
                                                            }
                                                            var v110 uint64 = v100
                                                            _ = v110
                                                        } else {
                                                            v64 -= 42
                                                            var v111 uint8 = 1
                                                            v112 := GenS14{F0: (^(v111))}
                                                            _ = v111
                                                            _ = v112
                                                        }
                                                        var v113 uint16 = 982
                                                        v114 := GenS11{F0: 1, F1: -86}
                                                        var v115 int16 = v97
                                                        v116 := GenS15{F0: (!(!(!true))), F1: fi37}
                                                        _ = v113
                                                        _ = v114
                                                        _ = v115
                                                        _ = v116
                                                    } else {
                                                        v117 := GenS16{F0: (v2 + v2), F1: 2147483648}
                                                        switch 4294967295 {
                                                        case 32768:
                                                            v1 = (v2 + v2)
                                                        case -9223372036854775808:
                                                            delete(v4, (("hello" + "") + "hello"))
                                                        }
                                                        _ = v117
                                                    }
                                                    v118 := GenS17{F0: fi36}
                                                    var v119 uint64 = v100
                                                    _ = v100
                                                    _ = v118
                                                    _ = v119
                                                case "foo bar":
                                                    var v120 string = "a"
                                                    _ = v120
                                                default:
                                                    fmt.Println(v2)
                                                }
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v121 := []int{}
                                                    _ = v121
                                                case "a":
                                                    v122 := []int{fi37, int(v97), 409, fi37}
                                                    _ = v122
                                                }
                                            case -9223372036854775808:
                                                v123 := GenS4{F0: 70}
                                                _ = v123
                                            }
                                            _ = v84
                                            _ = v85
                                            _ = v95
                                            _ = v96
                                            _ = v97
                                            _ = v98
                                        }
                                    }
                                }
                                _ = v64
                            } else {
                                var v124 int32 = 891
                                var v125 bool = (!(42 <= 1))
                                var v126 uint32 = 2
                                v127 := GenS18{F0: 2, F1: (v38 >> ((0) & 63))}
                                var v128 bool = (-975 != -32768)
                                if (255 >= 202) {
                                    v126 = (v126 << ((v35) & 63))
                                    v129 := GenS5{F0: 100}
                                    if v128 {
                                        var v130 float64 = float64(float32(1e10))
                                        if ((v38 % (((v38 % (((v38 & (v38 >> ((v35) & 63)))) | 1))) | 1)) != (^(v38))) {
                                            v131 := []int{(fi37 & fi37)}
                                            var v132 int = (fi37 >> ((v35) & 63))
                                            if v128 {
                                                v130 -= v130
                                                var v133 int16 = -32768
                                                if len(v131) > 0 {
                                                    v131[0] = fi37
                                                }
                                                _ = v133
                                            }
                                            _ = v131
                                            _ = v132
                                        }
                                        if (255 < 675) {
                                            if (318 == 707) {
                                                delete(v3, "")
                                                fmt.Println(v128)
                                                v3["a"] = "\n\t"
                                            } else {
                                                var v134 int64 = 42
                                                var v135 int64 = v38
                                                _ = v134
                                                _ = v135
                                            }
                                            if v125 {
                                                switch fi37 {
                                                case 974:
                                                    if v125 {
                                                        switch "a" {
                                                        case "foo bar":
                                                            v38 += (v38 & (v38 & v38))
                                                        case "":
                                                            v136 := []int{fi36, int(v126), int(v35), int(v38)}
                                                            _ = v136
                                                        }
                                                        v137 := GenS19{F0: fi37, F1: (^(v126))}
                                                        v138 := GenS20{F0: v1}
                                                        var v139 uint8 = 207
                                                        var v140 uint64 = uint64(v35)
                                                        _ = v137
                                                        _ = v138
                                                        _ = v139
                                                        _ = v140
                                                    }
                                                case 937:
                                                    var v141 int16 = -32768
                                                    _ = v141
                                                default:
                                                    fmt.Println(v130)
                                                }
                                                v35 += 283
                                                fmt.Println(v126, v128, v128)
                                            }
                                            var v142 uint16 = 42
                                            _ = v142
                                        } else {
                                            v143 := []int{4294967295}
                                            _ = v143
                                        }
                                        _ = v130
                                    }
                                    _ = v129
                                } else {
                                    v35 -= (^(v35))
                                }
                                _ = v124
                                _ = v125
                                _ = v126
                                _ = v127
                                _ = v128
                            }
                            fmt.Println(v38)
                        }
                        if (int8(fi37) != 127) {
                            v35 = v35
                            if (0.0 < 3.14159) {
                                v38 = (-(v38))
                                v144 := GenS21{F0: 42, F1: 127}
                                _ = v144
                            } else {
                                var v145 float64 = 0.0
                                if false {
                                    var v146 uint64 = 18446744073709551615
                                    fmt.Println(v2, v1)
                                    var v147 int = (fi36 / (((fi37 ^ (^(fi36)))) | 1))
                                    _ = v146
                                    _ = v147
                                }
                                _ = v145
                            }
                        } else {
                            switch (fi37 << ((2) & 63)) {
                            case 0:
                                v148 := []int{}
                                _ = v148
                            case -1:
                                var v149 float32 = (((((0.1 / (0.1 - 1.5)) * 0.0) / 1.0) * float32((3.14159 * 0.1))) * -2.25)
                                _ = v149
                            }
                            if (v35 <= 1) {
                                v150 := []int{fi36, fi37}
                                v3[v2] = v2
                                v38 -= v38
                                v1 = ((("foo bar" + v1) + ((v1 + v1) + ("\n\t" + v1))) + "\n\t")
                                if (((0 <= 1) || true) || true) {
                                    v150 = append(v150, (fi37 / ((-2147483648) | 1)))
                                }
                                _ = v150
                            } else {
                                fmt.Println(fi37)
                                v151 := GenS7{F0: 1}
                                var v152 uint = (^(v35))
                                _ = v151
                                _ = v152
                            }
                        }
                        v153 := []int{(fi36 / ((352) | 1)), (fi37 + 111), fi37, (fi37 ^ fi36)}
                        _ = v153
                    }
                }
                _ = v38
            } else {
                var v154 int = (^(fi36))
                v155 := GenS22{F0: -32768, F1: 830}
                if (!false) {
                    v156 := []int{fi37, int(v35), fi36}
                    switch 0 {
                    case -977:
                        var v157 string = v1
                        _ = v157
                    case -9223372036854775808:
                        var v158 float64 = (float64(((float32(1.5) + (0.0 - 1e10)) + (3.14159 - 3.14159))) / 1e10)
                        _ = v158
                    default:
                        var v159 uint = v35
                        _ = v159
                    }
                    var v160 int64 = 42
                    v160 = v160
                    if (!true) {
                        var v161 uint32 = 1
                        _ = v161
                    } else {
                        if (2147483647 <= 2147483647) {
                            v1 = (v1 + ("foo bar" + v1))
                            var v162 string = v1
                            v163 := GenS22{F0: 42, F1: 32767}
                            _ = v162
                            _ = v163
                        }
                        var v164 float64 = 3.14159
                        var v165 string = v1
                        v166 := GenS23{F0: 2147483647, F1: -48}
                        _ = v164
                        _ = v165
                        _ = v166
                    }
                    _ = v156
                    _ = v160
                } else {
                    switch "" {
                    case "a":
                        var v167 uint64 = 1
                        var v168 bool = false
                        if (!(42 > -1)) {
                            if v168 {
                                var v169 int = v154
                                _ = v169
                            }
                            v170 := []int{(fi36 * 4294967295), (fi37 >> ((255) & 63))}
                            var v171 int8 = 127
                            _ = v170
                            _ = v171
                        }
                        _ = v167
                        _ = v168
                    case "foo bar":
                        _ = len(v4)
                        v172 := v3[v1]
                        switch v1 {
                        case "a":
                            v172 = "foo bar"
                        case "\n\t":
                            switch 658 {
                            case 237:
                                v173 := GenS15{F0: (true && (true && false)), F1: v154}
                                _ = v173
                            case 32767:
                                fmt.Println(v35)
                            case 128:
                                switch ("foo bar" + (v1 + ((v172 + v172) + v1))) {
                                case "foo bar":
                                    v1 = "a"
                                case "":
                                    v2 = v172
                                    _ = len(v3)
                                    v174 := []int{852}
                                    _ = v174
                                }
                                var v175 uint = v35
                                v176 := GenS1{F0: 2147483647, F1: genFn1(42, 1), F2: 18446744073709551615}
                                var v177 string = "hello"
                                delete(v4, (v2 + v177))
                                var v178 int32 = -362
                                _ = v175
                                _ = v176
                                _ = v177
                                _ = v178
                            }
                        }
                        _ = v172
                    case "\n\t":
                        var v179 uint8 = 255
                        _ = v179
                    default:
                        if ("a" == ("a" + (v2 + ("foo bar" + ("foo bar" + ("a" + "foo bar")))))) {
                            v1 = "a"
                            v180 := v3[v2]
                            if (635 != 694) {
                                fmt.Println(v154)
                            } else {
                                v154 = 2147483647
                                switch (v154 & fi37) {
                                case 381:
                                    delete(v4, "foo bar")
                                case -850:
                                    v181 := GenS24{F0: 4294967295, F1: 127, F2: 42}
                                    _ = v181
                                case 882:
                                    v182 := GenS25{F0: 4294967295, F1: (((((v1 + v180) + "a") + (v1 + (v180 + ""))) + "") + v1)}
                                    _ = v182
                                default:
                                    var v183 string = (v2 + v2)
                                    _ = v183
                                }
                                v184 := GenS26{F0: 0, F1: "hello", F2: v180}
                                _ = v184
                            }
                            v185 := GenS27{F0: 58, F1: 806}
                            _ = v180
                            _ = v185
                        }
                    }
                    v186 := GenS28{F0: 1}
                    v2 = ((v2 + v2) + "a")
                    _ = v186
                }
                _ = v154
                _ = v155
            }
            var v187 uint64 = 2
            if (!(false || ((v187 * v187) > 387))) {
                v1 = ("foo bar" + v2)
            }
            v188 := GenS17{F0: fi36}
            var v189 string = (v1 + v2)
            _ = fi37
            _ = v187
            _ = v188
            _ = v189
        }
        v35 -= 2
        for fi190 := 0; fi190 < 1; fi190++ {
            v191 := []int{128, -9223372036854775808, fi190}
            v192 := v3[(("foo bar" + ("\n\t" + ("\n\t" + ("" + (("" + ("a" + ("a" + ("\n\t" + v1)))) + "\n\t"))))) + v1)]
            _ = fi190
            _ = v191
            _ = v192
        }
        _ = fi36
    }
    var v193 uint = 65535
    var v194 uint64 = 2
    var v195 int8 = -22
    fmt.Println("v1:", v1)
    fmt.Println("v193:", v193)
    fmt.Println("v194:", v194)
    fmt.Println("v195:", v195)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v193
    _ = v194
    _ = v195
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
