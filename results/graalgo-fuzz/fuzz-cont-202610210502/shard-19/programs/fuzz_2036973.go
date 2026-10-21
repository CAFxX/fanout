package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int8; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int64; F1 int8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int32; F1 uint32 }
type GenS12 struct { F0 bool; F1 uint }
type GenS13 struct { F0 int32; F1 bool }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int; F1 int }
type GenS16 struct { F0 bool; F1 uint }
type GenS17 struct { F0 int64; F1 string; F2 uint32 }
type GenS18 struct { F0 uint32; F1 uint64 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 int8; F1 string }
type GenS22 struct { F0 uint16 }
type GenS23 struct { F0 int }
type GenS24 struct { F0 string; F1 uint64 }
type GenS25 struct { F0 int }
type GenS26 struct { F0 uint64; F1 uint8; F2 bool }
func genFn0(p0 int, p1 int64) int8 {
    if (("foo bar" < "foo bar") || true) {
        switch p0 {
        case 187:
            fmt.Println(p0, p0)
        case 42:
            switch ("foo bar" + (("" + "hello") + ("" + "a"))) {
            case "\n\t":
                v5 := GenS1{F0: p0, F1: 51}
                _ = v5
            case "":
                var v6 int16 = -719
                _ = v6
            case "a":
                v7 := []int{p0}
                _ = v7
            }
        }
        v8 := []int{p0}
        _ = v8
    }
    switch (^(p0)) {
    case 2147483647:
        var v9 uint32 = 1
        _ = v9
    case 9223372036854775807:
        for fi10 := 0; fi10 < 11; fi10++ {
            p0 = p0
            var v11 string = ("" + "\n\t")
            _ = fi10
            _ = v11
        }
    default:
        if (31 == 42) {
            var v12 float32 = (3.14159 - (float32(((float64((0.0 / 0.1)) - 3.14159) - 0.1)) / 1.5))
            fmt.Println(v12)
            var v13 int8 = 100
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    return 127
}
func genFn1(p0 int64, p1 int) uint8 {
    v14 := GenS1{F0: (p1 * (p1 + p1)), F1: 127}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v14
    return 255
}

func main() {
    var v1 int8 = 31
    v2 := []uint32{2, 701, 586}
    var v3 float64 = 0.0
    v4 := make(map[string]uint32)
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := GenS2{F0: 42}
        var v17 int = (fi15 ^ 128)
        fmt.Println(v17, fi15, v3)
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        v3 = v3
        v18 := GenS2{F0: 358}
        _ = v18
    } else {
        fmt.Println(v1)
    }
    switch ((("" + "") + "a") + "a") {
    case "hello":
        var v19 uint8 = 1
        _ = v19
    }
    if true {
        v1 += (^(v1))
        var v20 int = 0
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 uint = 468
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        var v23 uint8 = 127
        _ = v23
    }
    v1 += 42
    switch 32767 {
    case -9223372036854775808:
        var v24 bool = false
        _ = v24
    default:
        v3 += v3
    }
    v3 = v3
    var v25 int64 = -1
    v4["hello"] = 65535
    v3 += (v3 * 1e10)
    v26 := GenS2{F0: 18446744073709551615}
    if (true || false) {
        switch 2147483647 {
        case 42:
            fmt.Println(v1, v25, v1)
        case 128:
            fmt.Println(v1)
        }
        var v27 int64 = 0
        _ = v27
    }
    if true {
        var v28 bool = ((4294967295 <= 255) || (false || false))
        _ = v28
    } else {
        for fi29 := 0; fi29 < 12; fi29++ {
            fmt.Println(v3, v3)
            v3 -= 3.14159
            var v30 bool = (905 != 692)
            _ = fi29
            _ = v30
        }
        v31 := GenS1{F0: 127, F1: v1}
        var v32 int = 32768
        _ = v31
        _ = v32
    }
    v3 = v3
    for fi33 := 0; fi33 < 6; fi33++ {
        if false {
            for fi34 := 0; fi34 < 19; fi34++ {
                var v35 uint = 65535
                v25 = v25
                _ = fi34
                _ = v35
            }
            for fi36 := 0; fi36 < 12; fi36++ {
                var v37 string = "a"
                v3 += v3
                _ = fi36
                _ = v37
            }
            var v38 int16 = 1000
            _ = v38
        } else {
            for fi39 := 0; fi39 < 5; fi39++ {
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                _ = len(v4)
                var v40 int16 = -1
                _ = fi39
                _ = v40
            }
        }
        _ = fi33
    }
    if (41 != 5) {
        for fi41 := 0; fi41 < 7; fi41++ {
            for fi42 := 0; fi42 < 6; fi42++ {
                if false {
                    v43 := v4["foo bar"]
                    if false {
                        v25 = -9223372036854775808
                        v25 += v25
                    } else {
                        if len(v2) > 0 {
                            v2[0] = v43
                        }
                        var v44 int8 = (v1 >> ((990) & 63))
                        if (0 < 4294967295) {
                            fmt.Println(fi41, fi41)
                            if true {
                                v45 := GenS2{F0: 65535}
                                if false {
                                    if false {
                                        switch fi42 {
                                        case 42:
                                            if (int(v25) == (fi42 << ((18446744073709551615) & 63))) {
                                                fmt.Println(v43, v44, fi42)
                                            } else {
                                                v44 += (v1 ^ v1)
                                            }
                                        case -2147483648:
                                            if true {
                                                if (4294967295 == 0) {
                                                    v25 += (v25 / ((v25) | 1))
                                                } else {
                                                    if true {
                                                        v1 -= v44
                                                    }
                                                    if (false && false) {
                                                        var v46 uint64 = 448
                                                        _ = len(v4)
                                                        v25 -= (v25 >> ((4294967295) & 63))
                                                        _ = v46
                                                    }
                                                }
                                                v43 = v43
                                            }
                                        }
                                        fmt.Println(v25, v43, v44)
                                        switch "foo bar" {
                                        case "hello":
                                            v47 := GenS3{F0: 42, F1: (-(v44)), F2: (-(v44))}
                                            _ = v47
                                        case "foo bar":
                                            if (!false) {
                                                v43 += v43
                                            }
                                        default:
                                            v1 = v44
                                        }
                                    } else {
                                        if true {
                                            if ((fi42 >= (fi42 / (((fi41 | (fi42 & 0))) | 1))) && true) {
                                                if (true || true) {
                                                    v1 = 104
                                                    var v48 uint16 = 65535
                                                    var v49 uint64 = 18446744073709551615
                                                    switch "" {
                                                    case "\n\t":
                                                        v50 := v4["\n\t"]
                                                        _ = v50
                                                    }
                                                    v51 := GenS4{F0: v44}
                                                    _ = v48
                                                    _ = v49
                                                    _ = v51
                                                }
                                                fmt.Println(v3)
                                                if (true && false) {
                                                    v25 += 1
                                                } else {
                                                    if ((fi41 | (fi42 ^ (fi42 % (((fi41 ^ fi42)) | 1)))) != fi41) {
                                                        if (!false) {
                                                            switch (fi41 & 127) {
                                                            case 740:
                                                                if (false && (float64(0.1) > v3)) {
                                                                    switch "hello" {
                                                                    case "":
                                                                        var v52 uint8 = uint8(v43)
                                                                        _ = v52
                                                                    }
                                                                } else {
                                                                    var v53 int64 = (v25 % (((v25 << ((0) & 63))) | 1))
                                                                    fmt.Println(fi42)
                                                                    _ = v53
                                                                }
                                                            case 839:
                                                                if ("foo bar" < "a") {
                                                                    v54 := GenS4{F0: (v1 % ((v44) | 1))}
                                                                    var v55 uint = 4294967295
                                                                    v4["\n\t"] = 2147483648
                                                                    v43 = 4294967295
                                                                    if true {
                                                                        if (!false) {
                                                                            if (!false) {
                                                                                var v56 int = (fi42 % ((fi42) | 1))
                                                                                _ = v56
                                                                            }
                                                                            var v57 int16 = 42
                                                                            _ = v57
                                                                        }
                                                                        var v58 int8 = 0
                                                                        v2 = append(v2, v43)
                                                                        _ = v58
                                                                    }
                                                                    _ = v54
                                                                    _ = v55
                                                                } else {
                                                                    var v59 float32 = ((float32(v3) - ((1.5 + 1.5) / 1.5)) + -2.25)
                                                                    _ = v59
                                                                }
                                                            case 32767:
                                                                var v60 int8 = (v1 ^ (v44 + genFn0(fi41, -9223372036854775808)))
                                                                _ = v60
                                                            default:
                                                                if true {
                                                                    switch ("\n\t" + "hello") {
                                                                    case "":
                                                                        v3 += v3
                                                                    case "a":
                                                                        var v61 uint64 = 207
                                                                        _ = v61
                                                                    case "hello":
                                                                        var v62 uint = 1
                                                                        _ = v62
                                                                    default:
                                                                        v43 += (v43 ^ 2147483648)
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        var v63 float32 = 0.0
                                                                        _ = v63
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v64 uint32 = (v43 & (v43 ^ 512))
                                                    _ = v64
                                                }
                                            }
                                            if (42 <= 42) {
                                                var v65 int16 = -1
                                                if ((!((((65535 >= 643) || false) || ((-2147483648 == 1) || (-1 >= 2147483647))) && (127 != 2))) || ((v65 ^ 42) <= (v65 << ((215) & 63)))) {
                                                    v4["a"] = v43
                                                    v66 := GenS2{F0: 255}
                                                    _ = v66
                                                } else {
                                                    if false {
                                                        switch fi42 {
                                                        case 1:
                                                            if true {
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v67 := GenS1{F0: fi41, F1: (v1 >> ((42) & 63))}
                                                                    _ = v67
                                                                case "foo bar":
                                                                    fmt.Println(v1, fi41)
                                                                case "hello":
                                                                    v3 -= v3
                                                                default:
                                                                    v44 = int8(fi41)
                                                                }
                                                            } else {
                                                                v68 := GenS1{F0: (fi41 & (^(fi41))), F1: v1}
                                                                _ = v68
                                                            }
                                                        case 4294967295:
                                                            if true {
                                                                if ((!(false || false)) || (89 < 42)) {
                                                                    var v69 bool = false
                                                                    switch "hello" {
                                                                    case "a":
                                                                        if v69 {
                                                                            v70 := GenS5{F0: 388, F1: genFn1(v25, fi42)}
                                                                            _ = v70
                                                                        }
                                                                    case "":
                                                                        v2 = append(v2, 2147483647)
                                                                    case "foo bar":
                                                                        fmt.Println(v44)
                                                                    default:
                                                                        switch (fi42 ^ 34) {
                                                                        case 4294967295:
                                                                            var v71 uint = 8
                                                                            _ = v71
                                                                        }
                                                                    }
                                                                    _ = v69
                                                                } else {
                                                                    v1 -= -1
                                                                }
                                                                v25 = v25
                                                                v2 = append(v2, 326)
                                                            } else {
                                                                if (genFn1((v25 % ((v25) | 1)), (^(fi42))) > 128) {
                                                                    if (!(!false)) {
                                                                        var v72 uint16 = 255
                                                                        _ = v72
                                                                    } else {
                                                                        v1 -= v44
                                                                        v4["hello"] = (v43 >> ((42) & 63))
                                                                    }
                                                                } else {
                                                                    var v73 bool = true
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v2 = append(v2, (v43 + v43))
                                                                        v74 := GenS6{F0: fi41, F1: -1}
                                                                        _ = v74
                                                                    case "":
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    }
                                                                    var v75 float32 = 0.1
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        v76 := GenS2{F0: 2}
                                                                        _ = v76
                                                                    case "":
                                                                        v77 := GenS3{F0: (v65 - int16(v43)), F1: genFn0((fi42 & (fi42 | 966)), -9223372036854775808), F2: -112}
                                                                        _ = v77
                                                                    case "a":
                                                                        v78 := GenS7{F0: 349}
                                                                        _ = v78
                                                                    }
                                                                    _ = v73
                                                                    _ = v75
                                                                }
                                                                v79 := GenS6{F0: (^(fi42)), F1: 961}
                                                                _ = v79
                                                            }
                                                            v25 -= (-(v25))
                                                            v80 := GenS5{F0: 1, F1: 255}
                                                            var v81 int8 = genFn0((fi42 & (fi41 >> ((324) & 63))), v25)
                                                            v82 := GenS8{F0: v25, F1: (v81 & (-(v81)))}
                                                            _ = v80
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                        _ = len(v4)
                                                        switch "a" {
                                                        case "\n\t":
                                                            v83 := GenS9{F0: fi41}
                                                            _ = v83
                                                        default:
                                                            var v84 uint8 = 127
                                                            _ = v84
                                                        }
                                                    } else {
                                                        fmt.Println(v25)
                                                        switch "hello" {
                                                        case "\n\t":
                                                            if true {
                                                                v85 := GenS10{F0: genFn1(v25, fi42)}
                                                                var v86 int = fi41
                                                                _ = v85
                                                                _ = v86
                                                            } else {
                                                                fmt.Println(v43, v1)
                                                            }
                                                        case "foo bar":
                                                            var v87 int16 = 540
                                                            _ = v87
                                                        default:
                                                            v65 = (v65 % ((v65) | 1))
                                                            var v88 int64 = 1
                                                            switch (^(fi41)) {
                                                            case -2147483648:
                                                                _ = len(v4)
                                                            case 0:
                                                                v89 := v4["hello"]
                                                                _ = v89
                                                            default:
                                                                v44 -= 100
                                                            }
                                                            _ = v88
                                                        }
                                                    }
                                                    v90 := GenS8{F0: (v25 % ((v25) | 1)), F1: v1}
                                                    v4["foo bar"] = v43
                                                    _ = v90
                                                }
                                                if ((v43 >> ((42) & 63)) >= 910) {
                                                    fmt.Println(fi42, fi42)
                                                    if (fi41 >= fi41) {
                                                        switch (("foo bar" + ("" + "")) + "a") {
                                                        case "":
                                                            v91 := GenS11{F0: 0, F1: (v43 / ((1) | 1))}
                                                            _ = v91
                                                        case "hello":
                                                            fmt.Println(v65, fi41)
                                                        default:
                                                            switch -1 {
                                                            case 127:
                                                                v92 := GenS10{F0: 1}
                                                                _ = v92
                                                            case 32767:
                                                                var v93 bool = true
                                                                _ = v93
                                                            case 0:
                                                                var v94 int8 = genFn0(814, v25)
                                                                _ = v94
                                                            }
                                                            v95 := GenS7{F0: 0}
                                                            v96 := v4["foo bar"]
                                                            v97 := GenS12{F0: (280 > 255), F1: 400}
                                                            v4[((("a" + "\n\t") + "hello") + ((("foo bar" + "a") + ("hello" + "hello")) + ((("a" + "a") + (("\n\t" + "\n\t") + "foo bar")) + "foo bar")))] = v43
                                                            if false {
                                                                v98 := GenS13{F0: -2147483648, F1: true}
                                                                v99 := GenS14{F0: 18446744073709551615}
                                                                _ = v98
                                                                _ = v99
                                                            }
                                                            var v100 bool = (v1 <= (v44 | (v1 * (v44 / ((-128) | 1)))))
                                                            var v101 uint = 454
                                                            var v102 float32 = float32(float64(3.14159))
                                                            _ = len(v4)
                                                            _ = v100
                                                            _ = v101
                                                            _ = v102
                                                            _ = v95
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        v43 += (v43 * (v43 ^ (^(v43))))
                                                        var v103 int64 = (v25 % ((281) | 1))
                                                        _ = v103
                                                    }
                                                    var v104 uint16 = 65535
                                                    _ = v104
                                                } else {
                                                    v44 -= 29
                                                    v25 -= (v25 / ((v25) | 1))
                                                    _ = len(v2)
                                                    var v105 int16 = v65
                                                    v25 += v25
                                                    var v106 string = ""
                                                    _ = v105
                                                    _ = v106
                                                }
                                                _ = v65
                                            } else {
                                                if true {
                                                    var v107 uint64 = 146
                                                    switch "hello" {
                                                    case "":
                                                        v4[""] = 65535
                                                    case "\n\t":
                                                        var v108 uint8 = 66
                                                        _ = v108
                                                    case "hello":
                                                        v109 := GenS10{F0: 42}
                                                        _ = v109
                                                    }
                                                    fmt.Println(v44, v43, fi41)
                                                    _ = v107
                                                } else {
                                                    switch fi41 {
                                                    case 1:
                                                        v3 += float64(1e10)
                                                    case 32767:
                                                        v44 -= (v1 * v44)
                                                    case -901:
                                                        if true {
                                                            var v110 uint16 = 255
                                                            v111 := GenS5{F0: 255, F1: 127}
                                                            _ = v110
                                                            _ = v111
                                                        } else {
                                                            fmt.Println(v3)
                                                            var v112 bool = (-2147483648 >= 0)
                                                            _ = v112
                                                        }
                                                    default:
                                                        var v113 uint8 = 2
                                                        if (false || true) {
                                                            fmt.Println(v25, v44)
                                                            if (!false) {
                                                                delete(v4, "a")
                                                                v114 := v4[("a" + "\n\t")]
                                                                fmt.Println(v25)
                                                                fmt.Println(v43, v3, v3)
                                                                _ = v114
                                                            }
                                                            if (!true) {
                                                                if true {
                                                                    v3 -= v3
                                                                    if len(v2) > 0 {
                                                                        v2[0] = (v43 - 2)
                                                                    }
                                                                    _ = len(v2)
                                                                } else {
                                                                    var v115 uint32 = v43
                                                                    if ((v25 & v25) >= (v25 ^ (v25 ^ 1))) {
                                                                        if ((("foo bar" + "hello") != "foo bar") && true) {
                                                                            v116 := GenS15{F0: (-(fi42)), F1: (fi42 >> ((18446744073709551615) & 63))}
                                                                            var v117 int64 = 0
                                                                            v118 := GenS16{F0: true, F1: 420}
                                                                            _ = v116
                                                                            _ = v117
                                                                            _ = v118
                                                                        } else {
                                                                            fmt.Println(v43, v3, v113)
                                                                            switch ("hello" + "foo bar") {
                                                                            case "\n\t":
                                                                                v115 = v115
                                                                            default:
                                                                                var v119 bool = (false || true)
                                                                                _ = v119
                                                                            }
                                                                            v4["foo bar"] = (v43 ^ v115)
                                                                            v3 += (v3 + v3)
                                                                        }
                                                                    } else {
                                                                        if ((((v3 + 0.1) / ((v3 + v3) - v3)) - (1e10 / (1.5 + float64((0.1 / ((0.0 + 3.14159) * 3.14159)))))) != ((v3 * v3) - (1e10 - float64(0.1)))) {
                                                                            v1 += (-(v1))
                                                                        }
                                                                    }
                                                                    if false {
                                                                        v120 := GenS6{F0: -46, F1: -178}
                                                                        _ = v120
                                                                    }
                                                                    _ = v115
                                                                }
                                                            } else {
                                                                var v121 int64 = v25
                                                                var v122 int16 = -1
                                                                v123 := GenS1{F0: fi42, F1: v1}
                                                                _ = v121
                                                                _ = v122
                                                                _ = v123
                                                            }
                                                        }
                                                        var v124 int16 = 1000
                                                        _ = v113
                                                        _ = v124
                                                    }
                                                    if (int32(v1) != 42) {
                                                        v2 = append(v2, (v43 + 919))
                                                        v125 := GenS8{F0: v25, F1: v1}
                                                        v25 -= (^(v25))
                                                        v126 := GenS7{F0: 4294967295}
                                                        v44 += v44
                                                        _ = v125
                                                        _ = v126
                                                    }
                                                    var v127 float32 = float32(v3)
                                                    _ = v127
                                                }
                                            }
                                            fmt.Println(v1, fi42, fi41)
                                        } else {
                                            v43 = v43
                                            var v128 string = ""
                                            _ = v128
                                        }
                                        fmt.Println(v43)
                                    }
                                } else {
                                    var v129 string = ""
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v130 uint32 = 65535
                                    if ((0 <= 32767) && ((!true) || ((255 == 65535) || false))) {
                                        var v131 int8 = genFn0(fi42, (^(v25)))
                                        _ = v131
                                    } else {
                                        if false {
                                            var v132 uint16 = 22
                                            switch fi41 {
                                            case 9223372036854775807:
                                                v133 := GenS17{F0: (^(v25)), F1: (v129 + "\n\t"), F2: 778}
                                                _ = v133
                                            case 68:
                                                fmt.Println(v130, v132, v1)
                                            default:
                                                v134 := GenS18{F0: v130, F1: 42}
                                                switch v129 {
                                                case "hello":
                                                    v3 -= float64(-2.25)
                                                case "\n\t":
                                                    var v135 uint16 = 249
                                                    _ = v135
                                                case "a":
                                                    v136 := GenS18{F0: (v43 | v43), F1: 654}
                                                    _ = v136
                                                }
                                                if ((v132 ^ v132) < (v132 - (^(v132)))) {
                                                    if len(v2) > 0 {
                                                        v2[0] = v130
                                                    }
                                                    v130 += 4294967295
                                                    fmt.Println(fi41, fi41)
                                                } else {
                                                    if false {
                                                        var v137 uint32 = v43
                                                        if false {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v138 := GenS6{F0: (fi42 | -1), F1: 0}
                                                                if (false && (true && false)) {
                                                                    var v139 float64 = ((0.0 + (1e10 / (v3 + v3))) / (1e10 / v3))
                                                                    _ = v139
                                                                } else {
                                                                    switch (fi42 << ((18446744073709551615) & 63)) {
                                                                    case 0:
                                                                        v140 := GenS18{F0: v137, F1: 2}
                                                                        _ = v140
                                                                    case -301:
                                                                        var v141 float64 = v3
                                                                        _ = v141
                                                                    case -1:
                                                                        var v142 string = v129
                                                                        _ = v142
                                                                    }
                                                                }
                                                                if (false && false) {
                                                                    fmt.Println(v3)
                                                                    v4[("hello" + v129)] = 744
                                                                    if ((v132 | (v132 % ((v132) | 1))) < (^(v132))) {
                                                                        v137 -= v43
                                                                        v143 := GenS11{F0: 2147483647, F1: (v43 - v43)}
                                                                        var v144 bool = (18446744073709551615 < 1)
                                                                        switch ((v129 + v129) + "hello") {
                                                                        case "hello":
                                                                            v145 := GenS10{F0: genFn1(v25, 239)}
                                                                            _ = v145
                                                                        case "foo bar":
                                                                            fmt.Println(v44)
                                                                        }
                                                                        _ = v143
                                                                        _ = v144
                                                                    } else {
                                                                        var v146 uint8 = uint8(v132)
                                                                        var v147 uint = 1
                                                                        v129 = ((("foo bar" + v129) + v129) + (v129 + "hello"))
                                                                        var v148 int16 = -588
                                                                        var v149 uint = v147
                                                                        _ = v146
                                                                        _ = v147
                                                                        _ = v148
                                                                        _ = v149
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        var v150 uint32 = 4294967295
                                                                        if ((3.14159 * 3.14159) <= (0.1 - (-2.25 - 0.1))) {
                                                                            v129 = "hello"
                                                                        }
                                                                        switch (fi42 - (-(fi41))) {
                                                                        case 4294967295:
                                                                            var v151 bool = (v25 > v25)
                                                                            _ = v151
                                                                        }
                                                                        _ = v150
                                                                    } else {
                                                                        if (v137 != (^(v43))) {
                                                                            v2 = append(v2, 665)
                                                                            v152 := GenS11{F0: 0, F1: (^(v130))}
                                                                            var v153 bool = false
                                                                            _ = v152
                                                                            _ = v153
                                                                        }
                                                                    }
                                                                    var v154 float64 = v3
                                                                    var v155 uint16 = 32767
                                                                    v156 := GenS16{F0: ((^(v132)) < (^(v155))), F1: 348}
                                                                    v157 := GenS2{F0: 255}
                                                                    _ = v154
                                                                    _ = v155
                                                                    _ = v156
                                                                    _ = v157
                                                                }
                                                                if true {
                                                                    var v158 int8 = v1
                                                                    _ = v158
                                                                }
                                                                _ = v138
                                                            case "hello":
                                                                v159 := GenS19{F0: (-9223372036854775808 > 1)}
                                                                _ = v159
                                                            }
                                                        } else {
                                                            if len(v2) > 0 {
                                                                v2[0] = (v137 << ((42) & 63))
                                                            }
                                                            if true {
                                                                if (34 > genFn1(42, fi42)) {
                                                                    if false {
                                                                        v1 += v1
                                                                        v160 := GenS14{F0: 951}
                                                                        _ = v160
                                                                    }
                                                                }
                                                                if true {
                                                                    var v161 int16 = int16(fi41)
                                                                    var v162 uint64 = 42
                                                                    var v163 bool = (!false)
                                                                    _ = v161
                                                                    _ = v162
                                                                    _ = v163
                                                                } else {
                                                                    fmt.Println(v43, v130)
                                                                    fmt.Println(v43, v1, v1)
                                                                }
                                                            } else {
                                                                v25 = 1
                                                            }
                                                        }
                                                        v43 += (v137 ^ (^(v130)))
                                                        switch (fi41 | -921) {
                                                        case 304:
                                                            v164 := GenS20{F0: -1}
                                                            _ = v164
                                                        case 9223372036854775807:
                                                            if (false || true) {
                                                                fmt.Println(v3)
                                                                var v165 bool = true
                                                                var v166 int16 = -1
                                                                _ = v165
                                                                _ = v166
                                                            } else {
                                                                delete(v4, "hello")
                                                                if true {
                                                                    var v167 int8 = (v1 - (-(v1)))
                                                                    var v168 int32 = 42
                                                                    var v169 int = (fi42 | -1)
                                                                    var v170 float64 = v3
                                                                    var v171 int16 = int16(v168)
                                                                    _ = v167
                                                                    _ = v168
                                                                    _ = v169
                                                                    _ = v170
                                                                    _ = v171
                                                                }
                                                                v3 -= 3.14159
                                                                v3 -= v3
                                                                if false {
                                                                    switch v129 {
                                                                    case "":
                                                                        var v172 int32 = -294
                                                                        _ = v172
                                                                    case "foo bar":
                                                                        var v173 string = ""
                                                                        _ = v173
                                                                    case "\n\t":
                                                                        if true {
                                                                            var v174 uint = 255
                                                                            fmt.Println(v130, v132)
                                                                            _ = v174
                                                                        }
                                                                    }
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        var v175 int16 = -1
                                                                        _ = v175
                                                                    default:
                                                                        v129 = "foo bar"
                                                                    }
                                                                } else {
                                                                    if (18446744073709551615 <= 1) {
                                                                        var v176 float64 = 1.5
                                                                        var v177 float32 = -2.25
                                                                        v178 := GenS21{F0: v44, F1: v129}
                                                                        _ = v176
                                                                        _ = v177
                                                                        _ = v178
                                                                    } else {
                                                                        var v179 int8 = (v44 | 0)
                                                                        if (((v129 + ((v129 + "") + v129)) + ((v129 + "\n\t") + ((v129 + v129) + "hello"))) <= (((v129 + "\n\t") + v129) + "foo bar")) {
                                                                            var v180 uint16 = v132
                                                                            var v181 string = v129
                                                                            _ = v180
                                                                            _ = v181
                                                                        } else {
                                                                            v182 := GenS22{F0: (v132 & v132)}
                                                                            _ = v182
                                                                        }
                                                                        v183 := GenS23{F0: (fi41 << ((4294967295) & 63))}
                                                                        v3 += v3
                                                                        v44 = v179
                                                                        _ = v179
                                                                        _ = v183
                                                                    }
                                                                    v132 -= (v132 * v132)
                                                                    if true {
                                                                        v130 = (v43 + (^(v130)))
                                                                        if false {
                                                                            if (false || (((3.14159 + 3.14159) + 0.0) > 0.0)) {
                                                                                var v184 int64 = v25
                                                                                var v185 bool = (1e10 != -2.25)
                                                                                switch -495 {
                                                                                case 310:
                                                                                    v44 = (v44 & 23)
                                                                                case -307:
                                                                                    v130 -= 2147483647
                                                                                }
                                                                                _ = v184
                                                                                _ = v185
                                                                            }
                                                                            var v186 uint = 0
                                                                            _ = v186
                                                                        }
                                                                        fmt.Println(fi41, v25)
                                                                    } else {
                                                                        var v187 int32 = 122
                                                                        _ = v187
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            v188 := GenS24{F0: v129, F1: 0}
                                                            var v189 uint8 = uint8(v132)
                                                            switch fi41 {
                                                            case 4294967295:
                                                                v190 := GenS25{F0: (fi41 + (-(fi42)))}
                                                                _ = v190
                                                            case -2147483648:
                                                                fmt.Println(v130, v132, v132)
                                                            }
                                                            var v191 bool = ("foo bar" != ((v129 + v129) + (v129 + "hello")))
                                                            switch (fi41 & (fi41 - 2147483647)) {
                                                            case -1:
                                                                v192 := GenS1{F0: (fi42 | int(v130)), F1: (v1 - 49)}
                                                                _ = v192
                                                            default:
                                                                v129 = ""
                                                            }
                                                            _ = v188
                                                            _ = v189
                                                            _ = v191
                                                        }
                                                        fmt.Println(v43)
                                                        var v193 float32 = (1.5 / 1e10)
                                                        v132 = 42
                                                        _ = v137
                                                        _ = v193
                                                    }
                                                    if (674 <= v43) {
                                                        v194 := GenS8{F0: int64(v43), F1: v1}
                                                        _ = v194
                                                    } else {
                                                        v129 = v129
                                                    }
                                                }
                                                _ = v134
                                            }
                                            var v195 bool = ((^(v1)) != (v44 % ((v1) | 1)))
                                            _ = v132
                                            _ = v195
                                        }
                                        fmt.Println(fi41, v1)
                                    }
                                    v196 := GenS16{F0: (false || ((fi42 | (fi41 / ((32768) | 1))) >= (-(fi41)))), F1: 255}
                                    _ = v129
                                    _ = v130
                                    _ = v196
                                }
                                fmt.Println(v25, v25, fi41)
                                _ = v45
                            }
                        } else {
                            v1 -= (v44 << ((562) & 63))
                        }
                        _ = v44
                    }
                    _ = v43
                } else {
                    var v197 int32 = -787
                    v3 = -2.25
                    var v198 uint = 0
                    var v199 string = ""
                    if true {
                        var v200 string = (v199 + v199)
                        _ = v200
                    }
                    var v201 int32 = v197
                    _ = v197
                    _ = v198
                    _ = v199
                    _ = v201
                }
                var v202 string = "\n\t"
                _ = fi42
                _ = v202
            }
            for fi203 := 0; fi203 < 10; fi203++ {
                v1 += (-(v1))
                _ = fi203
            }
            v204 := GenS15{F0: fi41, F1: fi41}
            _ = fi41
            _ = v204
        }
        v205 := GenS17{F0: (-(v25)), F1: ("" + ("\n\t" + "hello")), F2: 660}
        if (-128 >= genFn0(2147483647, v25)) {
            if true {
                var v206 int64 = (^(v25))
                switch ("a" + "a") {
                case "a":
                    var v207 int64 = v25
                    _ = v207
                case "foo bar":
                    v25 = (v206 | (v206 << ((18446744073709551615) & 63)))
                }
                v208 := GenS18{F0: 570, F1: 1}
                _ = v206
                _ = v208
            }
            if len(v2) > 0 {
                v2[0] = 1
            }
            if (535 > 1) {
                for fi209 := 0; fi209 < 13; fi209++ {
                    v3 = 1e10
                    var v210 bool = true
                    _ = fi209
                    _ = v210
                }
                _ = len(v4)
            }
        }
        _ = v205
    }
    v211 := GenS26{F0: 0, F1: 2, F2: true}
    var v212 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v212:", v212)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v211
    _ = v212
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
