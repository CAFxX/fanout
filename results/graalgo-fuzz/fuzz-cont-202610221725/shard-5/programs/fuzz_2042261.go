package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 string; F2 uint32 }
type GenS6 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint64; F1 int; F2 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32; F1 bool; F2 int32 }
type GenS10 struct { F0 int64; F1 uint; F2 uint64 }
type GenS11 struct { F0 uint64; F1 bool; F2 uint }
type GenS12 struct { F0 bool; F1 int64; F2 uint64 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 string; F1 int8; F2 uint16 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint; F1 bool }
func genFn0(p0 string, p1 uint8) string {
    p1 -= (p1 * (p1 * p1))
    var v4 uint64 = 2
    _ = p0
    _ = p1
    _ = v4
    return (p0 + "foo bar")
}
func genFn1() uint16 {
    v5 := GenS1{F0: false, F1: 128, F2: 42}
    if false {
        if true {
            if (1 >= 0) {
                if false {
                    v6 := []int{469, -630, 4294967295}
                    var v7 bool = (18446744073709551615 >= 65535)
                    var v8 int = 579
                    _ = v6
                    _ = v7
                    _ = v8
                }
                if false {
                    v9 := GenS2{F0: 2147483647, F1: "foo bar"}
                    var v10 int64 = -9223372036854775808
                    _ = v10
                    _ = v9
                }
            }
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS3{F0: 0, F1: 1}
                v13 := GenS4{F0: 2}
                _ = fi11
                _ = v12
                _ = v13
            }
        } else {
            var v14 uint = 42
            fmt.Println(v14)
            v14 = 0
            _ = v14
        }
        v15 := []int{128, 9223372036854775807}
        var v16 int16 = -32768
        _ = v15
        _ = v16
    }
    var v17 uint = 4294967295
    v18 := GenS5{F0: 96, F1: "foo bar", F2: 132}
    var v19 int16 = -32768
    _ = v17
    _ = v18
    _ = v19
    _ = v5
    return 0
}

func main() {
    v1 := []int{-9223372036854775808, -1}
    var v2 float32 = 3.14159
    var v3 int16 = 1000
    v2 -= -2.25
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    fmt.Println(v2, v2)
    if false {
        switch -2147483648 {
        case 0:
            var v20 int16 = (v3 | v3)
            _ = v20
        default:
            if len(v1) > 0 {
                v1[0] = -580
            }
        }
    }
    v21 := GenS4{F0: 1}
    if (127 != 127) {
        var v22 bool = (-724 <= 0)
        v3 += v3
        _ = v22
    }
    var v23 float32 = v2
    if (9223372036854775807 != 9223372036854775807) {
        switch int(v3) {
        case 32768:
            fmt.Println(v2)
        }
        v3 = 0
    }
    var v24 bool = true
    v25 := GenS1{F0: true, F1: 172, F2: 332}
    for fi26 := 0; fi26 < 19; fi26++ {
        if (4294967295 != 65535) {
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 uint32 = 0
                if true {
                    fmt.Println(v3, v28)
                } else {
                    var v29 int32 = 2147483647
                    if (false && true) {
                        var v30 int32 = -644
                        _ = v30
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            _ = len(v1)
                        default:
                            v23 = v2
                        }
                        switch "" {
                        case "foo bar":
                            var v31 bool = (v24 || v24)
                            _ = v31
                        default:
                            v23 += (v2 * (v2 + v23))
                            switch 127 {
                            case 1:
                                if (1000 > (v3 - -32768)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if (v28 <= (v28 & v28)) {
                                        v32 := GenS3{F0: 1, F1: 570}
                                        v23 = float32((1.5 + 1.5))
                                        _ = v32
                                    }
                                }
                            }
                            switch (("" + ("hello" + "foo bar")) + "\n\t") {
                            case "a":
                                if (2 <= 958) {
                                    v23 -= v23
                                    var v33 uint = 0
                                    fmt.Println(v29)
                                    _ = v33
                                }
                            case "\n\t":
                                switch "\n\t" {
                                case "foo bar":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case "\n\t":
                                    fmt.Println(fi26)
                                case "hello":
                                    if v24 {
                                        if true {
                                            if v24 {
                                                switch fi26 {
                                                case 817:
                                                    switch -1 {
                                                    case 128:
                                                        switch int(v28) {
                                                        case 128:
                                                            if v24 {
                                                                fmt.Println(v24, v23)
                                                                v28 = 2147483647
                                                            }
                                                        }
                                                    case 127:
                                                        var v34 int = (fi27 | (fi27 * fi27))
                                                        _ = v34
                                                    case -2147483648:
                                                        if (42 > 0) {
                                                            var v35 uint64 = 42
                                                            _ = v35
                                                        } else {
                                                            fmt.Println(v29, fi26)
                                                        }
                                                    }
                                                default:
                                                    v29 -= (v29 & (v29 ^ 2147483647))
                                                }
                                                switch (fi26 & 9223372036854775807) {
                                                case 2147483647:
                                                    v36 := GenS6{F0: 191, F1: -1, F2: (v28 / ((v28) | 1))}
                                                    _ = v36
                                                case 32768:
                                                    switch fi26 {
                                                    case -2147483648:
                                                        var v37 uint = 255
                                                        _ = v37
                                                    case 32767:
                                                        switch ("a" + "hello") {
                                                        case "foo bar":
                                                            v29 += v29
                                                        case "":
                                                            fmt.Println(v3, v3)
                                                            switch (("\n\t" + "hello") + "foo bar") {
                                                            case "\n\t":
                                                                if (118 < 127) {
                                                                    if v24 {
                                                                        v23 = v2
                                                                        v29 = 1
                                                                        var v38 int32 = (v29 * v29)
                                                                        _ = v38
                                                                    } else {
                                                                        var v39 uint16 = 32767
                                                                        _ = v39
                                                                    }
                                                                    v1 = append(v1, (fi26 << ((uint(v28)) & 63)))
                                                                    fmt.Println(v23)
                                                                } else {
                                                                    var v40 uint64 = 42
                                                                    v3 += (v3 ^ (v3 << ((465) & 63)))
                                                                    v2 += (v2 * 1e10)
                                                                    _ = v40
                                                                }
                                                            case "":
                                                                v28 -= (v28 - (^(v28)))
                                                            }
                                                            v29 += (v29 * 0)
                                                        }
                                                    default:
                                                        if v24 {
                                                            v41 := GenS2{F0: v29, F1: ("hello" + "\n\t")}
                                                            v3 += -32768
                                                            v1 = append(v1, (fi26 + (fi26 % ((9223372036854775807) | 1))))
                                                            _ = v41
                                                        } else {
                                                            var v42 int16 = -934
                                                            if v24 {
                                                                if v24 {
                                                                    v28 += v28
                                                                    switch int(v29) {
                                                                    case -894:
                                                                        var v43 float64 = 1e10
                                                                        v24 = (4294967295 < 0)
                                                                        v29 = v29
                                                                        var v44 float64 = 3.14159
                                                                        switch "a" {
                                                                        case "a":
                                                                            var v45 uint8 = 127
                                                                            _ = v45
                                                                        }
                                                                        _ = v43
                                                                        _ = v44
                                                                    default:
                                                                        var v46 uint64 = 1
                                                                        _ = v46
                                                                    }
                                                                    var v47 float64 = (1e10 - ((3.14159 * 0.0) / 0.1))
                                                                    _ = v47
                                                                } else {
                                                                    fmt.Println(v29)
                                                                    var v48 int8 = 52
                                                                    if v24 {
                                                                        v28 = 625
                                                                        if v24 {
                                                                            v28 = (v28 / (((v28 % ((v28) | 1))) | 1))
                                                                            v49 := GenS7{F0: 573, F1: (fi27 % ((fi26) | 1)), F2: v48}
                                                                            v50 := GenS7{F0: 775, F1: fi27, F2: v48}
                                                                            _ = v49
                                                                            _ = v50
                                                                        }
                                                                        switch (fi26 | 0) {
                                                                        case 128:
                                                                            v3 += v3
                                                                        case 4294967295:
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        }
                                                                    }
                                                                    v3 = -32768
                                                                    v51 := GenS8{F0: -2147483648}
                                                                    _ = v48
                                                                    _ = v51
                                                                }
                                                                if len(v1) > 0 {
                                                                    v1[0] = 32767
                                                                }
                                                                var v52 string = "foo bar"
                                                                _ = v52
                                                            }
                                                            var v53 uint16 = 32767
                                                            v29 -= (v29 & (-(v29)))
                                                            _ = v42
                                                            _ = v53
                                                        }
                                                        fmt.Println(v24, fi26)
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if (!false) {
                                                            if false {
                                                                if len(v1) > 0 {
                                                                    v1[0] = fi26
                                                                }
                                                            }
                                                        }
                                                        var v54 int32 = v29
                                                        fmt.Println(v24, v24)
                                                        _ = v54
                                                    }
                                                case 42:
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v55 int16 = 0
                                                        _ = v55
                                                    case "a":
                                                        v29 = v29
                                                    default:
                                                        if (v24 && v24) {
                                                            var v56 uint = 42
                                                            _ = v56
                                                        }
                                                    }
                                                default:
                                                    v57 := GenS2{F0: (^(v29)), F1: "foo bar"}
                                                    if (!false) {
                                                        v58 := GenS3{F0: 4294967295, F1: 4294967295}
                                                        v2 = v23
                                                        fmt.Println(v29, fi26)
                                                        _ = v58
                                                    }
                                                    _ = v57
                                                }
                                                v59 := GenS8{F0: fi26}
                                                _ = v59
                                            } else {
                                                _ = len(v1)
                                                fmt.Println(v3, v23)
                                                var v60 uint = 516
                                                _ = v60
                                            }
                                        }
                                        if (42 != -70) {
                                            v1 = append(v1, fi27)
                                            v28 += v28
                                        } else {
                                            _ = len(v1)
                                            if (255 > 128) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                switch (fi27 << ((185) & 63)) {
                                                case 0:
                                                    v28 -= 2
                                                default:
                                                    if ("hello" != "\n\t") {
                                                        if ((3.14159 / 3.14159) >= 1.5) {
                                                            switch ("hello" + ("\n\t" + ("hello" + ("" + "hello")))) {
                                                            case "a":
                                                                var v61 int = (fi26 ^ 128)
                                                                _ = v61
                                                            }
                                                            var v62 string = "foo bar"
                                                            v63 := GenS8{F0: (-(fi27))}
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        fmt.Println(v2, v24)
                                                    }
                                                }
                                                if v24 {
                                                    switch fi27 {
                                                    case 32767:
                                                        v28 -= (v28 - v28)
                                                    case -294:
                                                        fmt.Println(v2)
                                                    case 127:
                                                        v29 -= 0
                                                    default:
                                                        var v64 int = 1
                                                        _ = v64
                                                    }
                                                    var v65 int = (-(fi27))
                                                    v66 := GenS1{F0: (true && (9223372036854775807 >= 532)), F1: 128, F2: 65535}
                                                    _ = v65
                                                    _ = v66
                                                } else {
                                                    if (42 == 669) {
                                                        if (4294967295 != 540) {
                                                            switch -2147483648 {
                                                            case -9223372036854775808:
                                                                v29 = (v29 >> ((4294967295) & 63))
                                                            }
                                                            v67 := GenS3{F0: 42, F1: 42}
                                                            _ = v67
                                                        }
                                                        if (!("hello" == "foo bar")) {
                                                            if false {
                                                                if (v24 && true) {
                                                                    var v68 int16 = -32768
                                                                    var v69 uint16 = 2
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                                v70 := GenS3{F0: 909, F1: 18446744073709551615}
                                                                _ = v70
                                                            }
                                                            var v71 int16 = v3
                                                            var v72 uint = 4294967295
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                    } else {
                                                        v1 = append(v1, 9223372036854775807)
                                                        var v73 string = "\n\t"
                                                        _ = v73
                                                    }
                                                    v23 += (3.14159 / ((v2 * (v23 - v23)) - ((-2.25 / -2.25) + (v23 + (((-2.25 - 1e10) + v23) / 0.1)))))
                                                }
                                            } else {
                                                switch ("foo bar" + "\n\t") {
                                                case "hello":
                                                    v28 = v28
                                                case "foo bar":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                default:
                                                    v2 = (v23 - v23)
                                                }
                                                if v24 {
                                                    if len(v1) > 0 {
                                                        v1[0] = 32768
                                                    }
                                                    if v24 {
                                                        v74 := GenS3{F0: 42, F1: 18446744073709551615}
                                                        v75 := GenS9{F0: (v29 << ((65535) & 63)), F1: (v24 || true), F2: (v29 - v29)}
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    if v24 {
                                                        if v24 {
                                                            if v24 {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    var v76 int32 = 0
                                                                    _ = v76
                                                                default:
                                                                    v77 := GenS6{F0: 51, F1: 1, F2: 2147483647}
                                                                    _ = v77
                                                                }
                                                                fmt.Println(v28, v24)
                                                                v78 := GenS3{F0: 18446744073709551615, F1: 1}
                                                                _ = v78
                                                            }
                                                            var v79 int32 = int32(v3)
                                                            var v80 uint64 = 4294967295
                                                            _ = v79
                                                            _ = v80
                                                        } else {
                                                            v81 := GenS10{F0: 0, F1: 271, F2: 1}
                                                            v1 = append(v1, 2147483647)
                                                            _ = v81
                                                        }
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            var v82 int32 = (v29 | v29)
                                                            _ = v82
                                                        case "":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        case "hello":
                                                            var v83 uint8 = 0
                                                            _ = v83
                                                        }
                                                        var v84 uint8 = 128
                                                        _ = v84
                                                    } else {
                                                        switch (("\n\t" + (("" + "hello") + ("\n\t" + (("a" + "hello") + "hello")))) + "foo bar") {
                                                        case "a":
                                                            v28 = (v28 / ((v28) | 1))
                                                        }
                                                    }
                                                }
                                                fmt.Println(v24)
                                            }
                                        }
                                    } else {
                                        switch (-(fi26)) {
                                        case 9223372036854775807:
                                            if false {
                                                if v24 {
                                                    v85 := GenS11{F0: 4294967295, F1: (v24 || true), F2: 255}
                                                    var v86 uint64 = 4294967295
                                                    var v87 float32 = 0.1
                                                    _ = v85
                                                    _ = v86
                                                    _ = v87
                                                } else {
                                                    v88 := GenS12{F0: true, F1: 358, F2: 4294967295}
                                                    var v89 bool = (18446744073709551615 != 835)
                                                    _ = v88
                                                    _ = v89
                                                }
                                                var v90 uint64 = 18446744073709551615
                                                if true {
                                                    var v91 uint64 = v90
                                                    v90 -= v90
                                                    fmt.Println(v24, v23, fi27)
                                                    _ = v91
                                                } else {
                                                    v29 = v29
                                                    v24 = (169 < 879)
                                                    v92 := GenS13{F0: 1}
                                                    _ = v92
                                                }
                                                if v24 {
                                                    var v93 uint = 1
                                                    switch (("foo bar" + ("" + "hello")) + "foo bar") {
                                                    case "hello":
                                                        if true {
                                                            var v94 int32 = v29
                                                            _ = v94
                                                        }
                                                    case "\n\t":
                                                        switch fi27 {
                                                        case 2147483647:
                                                            var v95 int64 = 0
                                                            _ = v95
                                                        }
                                                    case "":
                                                        var v96 bool = ((v28 >> (((v93 * (v93 | (v93 << ((0) & 63))))) & 63)) <= v28)
                                                        _ = v96
                                                    }
                                                    _ = v93
                                                }
                                                var v97 uint64 = (v90 << ((680) & 63))
                                                switch ("foo bar" + "") {
                                                case "\n\t":
                                                    fmt.Println(v3)
                                                case "":
                                                    if (321 == v29) {
                                                        v2 -= 0.0
                                                        v1 = append(v1, 32768)
                                                        if true {
                                                            if (v24 && ((^(v29)) != 42)) {
                                                                v98 := GenS14{F0: ("foo bar" + "a")}
                                                                var v99 uint64 = (v90 ^ v97)
                                                                _ = v98
                                                                _ = v99
                                                            } else {
                                                                if v24 {
                                                                    v97 = v97
                                                                } else {
                                                                    switch 32767 {
                                                                    case 0:
                                                                        v100 := GenS15{F0: "foo bar", F1: 42, F2: 65535}
                                                                        _ = v100
                                                                    }
                                                                    v90 = v90
                                                                }
                                                            }
                                                            if v24 {
                                                                fmt.Println(v24, v23, v28)
                                                                var v101 string = "\n\t"
                                                                v90 = v97
                                                                _ = v101
                                                            } else {
                                                                var v102 uint = 0
                                                                fmt.Println(v23, v29, v90)
                                                                fmt.Println(v90, v23)
                                                                _ = v102
                                                            }
                                                            fmt.Println(fi27, fi26, fi27)
                                                        } else {
                                                            v103 := GenS5{F0: 1, F1: "foo bar", F2: (^(v28))}
                                                            _ = v103
                                                        }
                                                        var v104 uint64 = (v90 ^ 1)
                                                        if len(v1) > 0 {
                                                            v1[0] = -798
                                                        }
                                                        _ = v104
                                                    }
                                                default:
                                                    fmt.Println(v90)
                                                    var v105 int64 = -1
                                                    if v24 {
                                                        v23 = (v23 * v2)
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v29 += (v29 & (v29 >> ((551) & 63)))
                                                        v28 -= (v28 - v28)
                                                    }
                                                    _ = v105
                                                }
                                                _ = v90
                                                _ = v97
                                            }
                                        }
                                        v106 := GenS16{F0: 0}
                                        switch "foo bar" {
                                        case "":
                                            var v107 int64 = 9223372036854775807
                                            _ = v107
                                        case "a":
                                            var v108 int8 = 65
                                            _ = v108
                                        }
                                        switch 9223372036854775807 {
                                        case 9223372036854775807:
                                            v109 := GenS14{F0: ((("\n\t" + "") + "") + (("\n\t" + "hello") + "foo bar"))}
                                            _ = v109
                                        case 32768:
                                            v23 = (v2 + (float32((0.0 * (float64(v23) + 1.5))) - (((v2 + v23) / 3.14159) * ((v23 + ((0.0 - v23) * v23)) / (((float32(0.1) - v23) + float32(1.5)) * ((((v23 / (((v2 / v23) + float32(1.5)) / -2.25)) / (float32(-2.25) - 1e10)) - (v2 - (3.14159 / 1.5))) - v2))))))
                                        }
                                        var v110 int64 = 1
                                        _ = v106
                                        _ = v110
                                    }
                                default:
                                    switch "foo bar" {
                                    case "":
                                        switch fi26 {
                                        case 0:
                                            var v111 int64 = 355
                                            _ = v111
                                        case -26:
                                            fmt.Println(v28, v2)
                                        case 32767:
                                            var v112 uint8 = 2
                                            _ = v112
                                        }
                                    default:
                                        if false {
                                            var v113 int64 = 0
                                            _ = v113
                                        } else {
                                            v29 += v29
                                            v114 := GenS9{F0: -1, F1: (v24 || false), F2: -435}
                                            if v24 {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v115 uint64 = 641
                                                _ = v115
                                            } else {
                                                if (v3 <= (v3 - (v3 | -1))) {
                                                    switch fi27 {
                                                    case 2147483647:
                                                        v28 -= 2147483648
                                                    case 4294967295:
                                                        fmt.Println(v29, v23, v24)
                                                    }
                                                }
                                            }
                                            fmt.Println(fi26)
                                            _ = v114
                                        }
                                    }
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "":
                                switch "hello" {
                                case "\n\t":
                                    _ = len(v1)
                                default:
                                    v3 -= -150
                                    v116 := GenS17{F0: (fi26 | fi26)}
                                    _ = v116
                                }
                            }
                            if ("hello" >= ("\n\t" + "")) {
                                _ = len(v1)
                                v28 = v28
                                var v117 int16 = (v3 / ((int16(fi27)) | 1))
                                _ = v117
                            }
                            v118 := GenS17{F0: 831}
                            var v119 bool = ((0 > 27) || (fi27 == fi27))
                            _ = v118
                            _ = v119
                        }
                        fmt.Println(v2, v2, v24)
                    }
                    _ = v29
                }
                v120 := GenS18{F0: 2, F1: ((v28 - v28) < 2)}
                _ = fi27
                _ = v120
                _ = v28
            }
            v1 = append(v1, (fi26 + fi26))
            if len(v1) > 0 {
                v1[0] = fi26
            }
        }
        _ = fi26
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
