package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint16; F1 int; F2 int64 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 bool; F2 uint }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 uint32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint32; F1 uint16; F2 uint }
type GenS17 struct { F0 int16; F1 uint64 }
type GenS18 struct { F0 string; F1 uint16 }
type GenS19 struct { F0 uint; F1 uint; F2 int16 }
type GenS20 struct { F0 uint32; F1 int32 }
type GenS21 struct { F0 uint8; F1 int64 }
type GenS22 struct { F0 string; F1 uint; F2 uint8 }
func genFn0(p0 uint16, p1 int16) uint8 {
    p1 += (p1 & -171)
    v5 := []int{-300, -1}
    if len(v5) > 0 {
        v5[0] = 691
    }
    switch 9223372036854775807 {
    case -2147483648:
        _ = len(v5)
    case 2147483647:
        p1 = p1
    case 32767:
        p1 = -224
    }
    v6 := GenS1{F0: 18446744073709551615, F1: (true && (2147483648 != 1)), F2: 0}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    var v3 int8 = -19
    v4 := make(map[uint32]int)
    v7 := GenS1{F0: 1, F1: true, F2: 163}
    var v8 uint8 = 83
    v8 += v8
    if true {
        var v9 int8 = -128
        if (false && (-1 == -1)) {
            fmt.Println(v3, v3)
            switch "hello" {
            case "\n\t":
                if (!(!false)) {
                    v10 := GenS1{F0: 65535, F1: (!(!(true && false))), F2: v8}
                    _ = v10
                } else {
                    if (255 >= 279) {
                        v11 := GenS1{F0: 18446744073709551615, F1: ((-9223372036854775808 == 9223372036854775807) && (-2.25 <= 0.0)), F2: (v8 % ((genFn0(42, 1000)) | 1))}
                        _ = v11
                    }
                }
            }
            var v12 float32 = 3.14159
            _ = v12
        } else {
            v8 = v8
            v13 := GenS2{F0: -1, F1: ""}
            switch 2147483647 {
            case -71:
                for fi14 := 0; fi14 < 18; fi14++ {
                    for fi15 := 0; fi15 < 18; fi15++ {
                        fmt.Println(v9, v3)
                        var v16 uint = 1
                        var v17 int = (fi14 - 32768)
                        var v18 int8 = -8
                        var v19 uint = 42
                        var v20 uint = 830
                        _ = len(v2)
                        if (0 != 0) {
                            switch v17 {
                            case 32767:
                                if (42 != -1) {
                                    v21 := GenS2{F0: int16(v9), F1: ("foo bar" + "\n\t")}
                                    fmt.Println(v9, v16, v18)
                                    _ = v21
                                } else {
                                    if false {
                                        var v22 uint8 = v8
                                        switch "foo bar" {
                                        case "\n\t":
                                            if true {
                                                switch ("foo bar" + "\n\t") {
                                                case "hello":
                                                    switch ("a" + "\n\t") {
                                                    case "hello":
                                                        v2[("foo bar" + "\n\t")] = "hello"
                                                    case "":
                                                        fmt.Println(v8, v18)
                                                    case "a":
                                                        if true {
                                                            _ = len(v1)
                                                            v23 := []int{fi14}
                                                            v18 += (v9 / (((v3 / ((v9) | 1))) | 1))
                                                            _ = v23
                                                        }
                                                    default:
                                                        var v24 uint64 = 18446744073709551615
                                                        _ = v24
                                                    }
                                                case "":
                                                    var v25 float32 = 1e10
                                                    _ = v25
                                                }
                                                switch "foo bar" {
                                                case "hello":
                                                    if ((v3 + v18) != 0) {
                                                        v17 -= (fi14 ^ (fi14 ^ (-(v17))))
                                                        if (true || false) {
                                                            switch ("" + ((((("\n\t" + "") + (("foo bar" + "\n\t") + "a")) + "a") + "\n\t") + "a")) {
                                                            case "":
                                                                switch ("foo bar" + "") {
                                                                case "":
                                                                    v26 := []int{}
                                                                    _ = v26
                                                                case "foo bar":
                                                                    fmt.Println(v17, v22, v8)
                                                                case "hello":
                                                                    v17 = 128
                                                                    if (1e10 != 1e10) {
                                                                        v27 := []int{(^(fi14)), (fi15 ^ fi15)}
                                                                        switch (fi15 % ((2147483647) | 1)) {
                                                                        case 1:
                                                                            if ((v8 + 255) != uint8(v20)) {
                                                                                fmt.Println(v8)
                                                                                delete(v1, 1)
                                                                            }
                                                                        default:
                                                                            _ = len(v27)
                                                                        }
                                                                        _ = v27
                                                                    } else {
                                                                        fmt.Println(v20, fi14, v9)
                                                                        v28 := GenS2{F0: 0, F1: "foo bar"}
                                                                        if (true || true) {
                                                                            switch v17 {
                                                                            case 32768:
                                                                                if true {
                                                                                    delete(v1, 42)
                                                                                }
                                                                            }
                                                                            var v29 uint32 = 4294967295
                                                                            v22 -= 42
                                                                            var v30 int16 = 0
                                                                            _ = v29
                                                                            _ = v30
                                                                        } else {
                                                                            if (false || (0 == -560)) {
                                                                                v31 := GenS3{F0: 1, F1: v17, F2: 9223372036854775807}
                                                                                _ = v31
                                                                            }
                                                                            if false {
                                                                                if true {
                                                                                    v32 := GenS4{F0: 0, F1: v9}
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        if false {
                                                                                            switch (fi14 << ((v19) & 63)) {
                                                                                            case 477:
                                                                                                v19 -= v16
                                                                                            }
                                                                                            var v33 uint16 = 397
                                                                                            v34 := GenS3{F0: 1, F1: (v17 & fi14), F2: -113}
                                                                                            _ = v33
                                                                                            _ = v34
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v4[2] = (v17 * 0)
                                                                                    case "hello":
                                                                                        v8 -= v8
                                                                                        _ = len(v1)
                                                                                    default:
                                                                                        var v35 uint8 = (^(v8))
                                                                                        _ = v35
                                                                                    }
                                                                                    var v36 int8 = v18
                                                                                    _ = v32
                                                                                    _ = v36
                                                                                }
                                                                            } else {
                                                                                v16 = (v20 | (^(v20)))
                                                                                v1[65535] = 42
                                                                            }
                                                                            var v37 uint32 = 1
                                                                            _ = v37
                                                                        }
                                                                        _ = v28
                                                                    }
                                                                    var v38 int64 = 9223372036854775807
                                                                    _ = v38
                                                                }
                                                            default:
                                                                v18 -= v3
                                                            }
                                                            var v39 float32 = (float32(3.14159) / 1.0)
                                                            _ = v39
                                                        } else {
                                                            v40 := []int{1}
                                                            v16 -= (v19 >> (((v16 << ((v19) & 63))) & 63))
                                                            _ = v40
                                                        }
                                                        v41 := GenS5{F0: 648}
                                                        v42 := GenS6{F0: (v16 | v20)}
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    if true {
                                                        v43 := GenS6{F0: (v16 + v16)}
                                                        _ = v43
                                                    }
                                                case "a":
                                                    delete(v2, "foo bar")
                                                case "foo bar":
                                                    if false {
                                                        fmt.Println(v18)
                                                    } else {
                                                        switch ((("foo bar" + "a") + ("hello" + "foo bar")) + "\n\t") {
                                                        case "":
                                                            if false {
                                                                v8 += (v22 - genFn0(255, 394))
                                                                v22 = (v8 | v8)
                                                            } else {
                                                                if ((594 != 4294967295) || true) {
                                                                    switch fi14 {
                                                                    case -663:
                                                                        v17 = (fi15 ^ fi15)
                                                                    case 708:
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            v44 := []int{127, fi15}
                                                                            _ = v44
                                                                        case "\n\t":
                                                                            var v45 bool = true
                                                                            _ = v45
                                                                        default:
                                                                            delete(v2, (("foo bar" + "hello") + ("hello" + "a")))
                                                                        }
                                                                    }
                                                                }
                                                                fmt.Println(fi15, v18, v19)
                                                                v46 := GenS5{F0: 42}
                                                                _ = v46
                                                            }
                                                        case "a":
                                                            switch (^(fi14)) {
                                                            case 278:
                                                                v20 = (^(v20))
                                                            case 127:
                                                                if true {
                                                                    v16 = 4294967295
                                                                    var v47 uint8 = v22
                                                                    switch ("\n\t" + ("a" + "")) {
                                                                    case "hello":
                                                                        var v48 uint64 = 1
                                                                        _ = v48
                                                                    default:
                                                                        fmt.Println(fi14, v9, v18)
                                                                    }
                                                                    v49 := GenS7{F0: true, F1: false, F2: (v16 % (((v16 >> ((v20) & 63))) | 1))}
                                                                    fmt.Println(v8, v3)
                                                                    _ = v47
                                                                    _ = v49
                                                                }
                                                            default:
                                                                fmt.Println(v19, v19)
                                                            }
                                                        default:
                                                            switch "" {
                                                            case "hello":
                                                                v3 += 32
                                                            case "\n\t":
                                                                if (484 < 42) {
                                                                    if false {
                                                                        v50 := v4[42]
                                                                        _ = v50
                                                                    }
                                                                    if (42 < 42) {
                                                                        v2["hello"] = ""
                                                                    }
                                                                    var v51 float64 = 0.0
                                                                    _ = v51
                                                                }
                                                            default:
                                                                switch 2147483647 {
                                                                case 474:
                                                                    v52 := []int{(fi14 | (fi15 | fi15)), 4294967295}
                                                                    _ = v52
                                                                case 365:
                                                                    v17 += (fi15 * (-(fi15)))
                                                                }
                                                            }
                                                        }
                                                        v53 := GenS8{F0: fi15}
                                                        _ = v53
                                                    }
                                                default:
                                                    v22 = genFn0(0, 1)
                                                }
                                                var v54 int32 = -233
                                                _ = v54
                                            } else {
                                                v55 := []int{fi15}
                                                var v56 uint = 1
                                                _ = v55
                                                _ = v56
                                            }
                                        default:
                                            if (v3 > v18) {
                                                v57 := GenS9{F0: 42, F1: 2147483648}
                                                if true {
                                                    switch (((("" + "foo bar") + "foo bar") + "\n\t") + "\n\t") {
                                                    case "a":
                                                        var v58 uint16 = 2
                                                        var v59 uint32 = uint32(v58)
                                                        var v60 int8 = v18
                                                        _ = v58
                                                        _ = v59
                                                        _ = v60
                                                    case "hello":
                                                        v61 := GenS10{F0: 4294967295}
                                                        _ = v61
                                                    default:
                                                        if false {
                                                            v62 := []int{179}
                                                            var v63 uint16 = 266
                                                            v64 := GenS11{F0: -1}
                                                            v65 := GenS1{F0: 18446744073709551615, F1: (!((65535 <= 2) || false)), F2: genFn0(32767, -55)}
                                                            _ = v62
                                                            _ = v63
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(fi15, v17, fi15)
                                                    var v66 int32 = -80
                                                    if (!((v16 ^ 148) > (v20 * (v19 - 255)))) {
                                                        switch "a" {
                                                        case "foo bar":
                                                            switch -2147483648 {
                                                            case 128:
                                                                if false {
                                                                    if (!(!(float64((0.1 / ((1.5 - 0.0) + float32(((0.0 + (((float64(1e10) + ((1.5 * 1e10) * 3.14159)) * 1e10) + ((float64((3.14159 / 0.1)) / ((3.14159 - 0.0) / -2.25)) / 3.14159))) / ((1e10 - -2.25) / (0.1 + 1e10))))))) == (1e10 * 1e10)))) {
                                                                        var v67 float32 = -2.25
                                                                        var v68 uint32 = 4294967295
                                                                        v16 += (^(v19))
                                                                        _ = v67
                                                                        _ = v68
                                                                    } else {
                                                                        if (!((1 <= 2) && false)) {
                                                                            v69 := []int{(fi14 + (fi15 | 4294967295)), fi15, 42}
                                                                            _ = v69
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                v70 := []int{}
                                                                _ = v70
                                                            }
                                                        }
                                                        v66 -= v66
                                                        var v71 int64 = 42
                                                        v72 := GenS10{F0: 2}
                                                        var v73 int16 = 0
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    _ = v66
                                                }
                                                _ = v57
                                            } else {
                                                switch "hello" {
                                                case "a":
                                                    if (!false) {
                                                        if ((406 == 521) && (-2147483648 != 42)) {
                                                            if ((true || (3.14159 <= 0.0)) && false) {
                                                                var v74 uint64 = 804
                                                                fmt.Println(v20, v17)
                                                                if ((false || (42 <= 0)) || false) {
                                                                    v75 := GenS12{F0: v74, F1: v74, F2: (v9 >> ((298) & 63))}
                                                                    fmt.Println(fi14)
                                                                    var v76 int64 = -1
                                                                    var v77 uint8 = genFn0(42, int16(fi14))
                                                                    if true {
                                                                        fmt.Println(v3, v16, v17)
                                                                        var v78 int16 = 1000
                                                                        _ = v78
                                                                    } else {
                                                                        if ((("a" + ("\n\t" + "")) >= "hello") && (13 > 89)) {
                                                                            var v79 uint8 = (v22 + genFn0(680, 32767))
                                                                            v20 -= v16
                                                                            _ = v79
                                                                        }
                                                                        v80 := []int{(v17 << ((v20) & 63)), fi15, 867}
                                                                        _ = len(v80)
                                                                        _ = v80
                                                                    }
                                                                    _ = v75
                                                                    _ = v76
                                                                    _ = v77
                                                                }
                                                                _ = v74
                                                            } else {
                                                                var v81 uint64 = 2
                                                                var v82 uint32 = 1
                                                                v83 := GenS9{F0: 42, F1: 2147483648}
                                                                _ = v81
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                            v9 = -28
                                                            v84 := GenS13{F0: -2147483648}
                                                            var v85 uint64 = 228
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                    }
                                                case "\n\t":
                                                    var v86 int8 = v18
                                                    _ = v86
                                                case "foo bar":
                                                    v3 = (v9 & (v18 * -1))
                                                default:
                                                    v87 := GenS14{F0: 1}
                                                    _ = v87
                                                }
                                                switch "hello" {
                                                case "hello":
                                                    if false {
                                                        switch (fi15 / ((fi14) | 1)) {
                                                        case 127:
                                                            v88 := GenS15{F0: v3}
                                                            _ = v88
                                                        case 4294967295:
                                                            var v89 int = (fi14 ^ (fi15 ^ fi15))
                                                            _ = v89
                                                        case 131:
                                                            if true {
                                                                if false {
                                                                    v90 := GenS13{F0: -2147483648}
                                                                    v91 := []int{-858, v17, -555, fi14}
                                                                    var v92 uint8 = 1
                                                                    _ = v90
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    v22 += genFn0(0, -32768)
                                                                    v93 := []int{fi15}
                                                                    _ = v93
                                                                }
                                                            }
                                                        }
                                                        if (!(-254 <= 42)) {
                                                            v20 += v19
                                                            var v94 uint64 = 0
                                                            var v95 uint8 = (v8 >> (((v19 - (v20 ^ (v19 - 0)))) & 63))
                                                            _ = v94
                                                            _ = v95
                                                        }
                                                        var v96 int16 = 0
                                                        v20 = v16
                                                        fmt.Println(v22)
                                                        _ = v96
                                                    } else {
                                                        fmt.Println(v3, v3, v8)
                                                        if (2 != 542) {
                                                            var v97 uint64 = 1
                                                            _ = v97
                                                        }
                                                        v98 := GenS10{F0: 18446744073709551615}
                                                        v16 -= v19
                                                        _ = v98
                                                    }
                                                }
                                            }
                                        }
                                        if (42 <= 65535) {
                                            if (792 >= 255) {
                                                v16 = (^(v19))
                                            } else {
                                                var v99 float64 = 3.14159
                                                v1[749] = 879
                                                var v100 uint32 = 2147483648
                                                _ = v100
                                                _ = v99
                                            }
                                        }
                                        fmt.Println(v3, fi15)
                                        switch ("" + "foo bar") {
                                        case "\n\t":
                                            var v101 uint64 = 0
                                            _ = v101
                                        case "a":
                                            var v102 uint = (v16 * v16)
                                            var v103 uint32 = 4294967295
                                            _ = v102
                                            _ = v103
                                        case "":
                                            v22 -= v22
                                        default:
                                            fmt.Println(v20)
                                        }
                                        if false {
                                            v17 -= (v17 * (v17 << ((743) & 63)))
                                            v18 = -128
                                            v104 := GenS9{F0: 32767, F1: 1}
                                            v105 := GenS8{F0: (v17 - 332)}
                                            _ = v104
                                            _ = v105
                                        }
                                        if (((0.1 * 1e10) * 1.5) >= 1.5) {
                                            v106 := GenS16{F0: 65535, F1: 2, F2: (^(v16))}
                                            v107 := GenS10{F0: 42}
                                            _ = v106
                                            _ = v107
                                        } else {
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v108 string = "foo bar"
                                                _ = v108
                                            case "hello":
                                                switch "" {
                                                case "hello":
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v20 = v20
                                                    case "":
                                                        v17 = (fi14 ^ fi15)
                                                    }
                                                case "foo bar":
                                                    var v109 uint64 = 658
                                                    _ = v109
                                                }
                                            case "\n\t":
                                                var v110 uint8 = 128
                                                _ = v110
                                            default:
                                                v9 = (v9 * v18)
                                            }
                                        }
                                        var v111 uint8 = 194
                                        _ = v111
                                        _ = v22
                                    }
                                }
                            case 2147483647:
                                switch "hello" {
                                case "hello":
                                    v112 := v2[("\n\t" + (("a" + "a") + "foo bar"))]
                                    _ = v112
                                case "a":
                                    if true {
                                        v113 := v2["foo bar"]
                                        var v114 int16 = 1000
                                        _ = v113
                                        _ = v114
                                    } else {
                                        v115 := []int{(fi14 + 609), (fi15 << ((v16) & 63)), -1, fi14}
                                        _ = len(v115)
                                        _ = v115
                                    }
                                    fmt.Println(v19)
                                }
                            default:
                                var v116 uint16 = 2
                                _ = v116
                            }
                            v8 = v8
                            fmt.Println(v18, fi15, v9)
                        } else {
                            var v117 uint64 = 1
                            v118 := GenS6{F0: (v16 + v19)}
                            _ = v117
                            _ = v118
                        }
                        if (true && true) {
                            v119 := GenS17{F0: 1000, F1: 18446744073709551615}
                            var v120 uint64 = 0
                            var v121 int8 = v9
                            fmt.Println(v9, v17, fi15)
                            _ = v119
                            _ = v120
                            _ = v121
                        } else {
                            if (true && (false || false)) {
                                var v122 uint64 = 988
                                _ = v122
                            }
                            if ((((1e10 - -2.25) < 0.0) || (1 >= 615)) && (-1 == 594)) {
                                if (!((!((1 == 496) || true)) || false)) {
                                    v19 += (^(v16))
                                    var v123 int8 = -128
                                    _ = v123
                                } else {
                                    delete(v4, 42)
                                    switch "hello" {
                                    case "foo bar":
                                        fmt.Println(v16, fi15, fi14)
                                    case "hello":
                                        v124 := GenS14{F0: 18446744073709551615}
                                        _ = v124
                                    }
                                    v1[2147483647] = 0
                                }
                            }
                            _ = len(v1)
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    switch "a" {
                    case "hello":
                        v125 := GenS2{F0: -271, F1: ("\n\t" + "")}
                        _ = v125
                    case "":
                        v126 := GenS18{F0: "\n\t", F1: 574}
                        _ = v126
                    case "\n\t":
                        v3 -= v3
                    default:
                        for fi127 := 0; fi127 < 13; fi127++ {
                            v9 += v9
                            if ((float64((1.5 / 1.0)) - -2.25) >= 3.14159) {
                                fmt.Println(fi14, fi14, v3)
                                var v128 float64 = -2.25
                                switch "foo bar" {
                                case "":
                                    if (0 > 2147483647) {
                                        if false {
                                            var v129 uint32 = 862
                                            switch (fi14 % ((fi127) | 1)) {
                                            case -2147483648:
                                                fmt.Println(v9)
                                            case 319:
                                                v130 := GenS19{F0: 2, F1: 423, F2: int16(fi14)}
                                                _ = v130
                                            }
                                            var v131 bool = false
                                            fmt.Println(v129, v8)
                                            v129 = 267
                                            _ = v129
                                            _ = v131
                                        }
                                        v132 := []int{-65, -153, fi127}
                                        _ = v132
                                    } else {
                                        var v133 int8 = (-(v9))
                                        var v134 bool = true
                                        _ = v133
                                        _ = v134
                                    }
                                default:
                                    v128 += v128
                                }
                                _ = v128
                            }
                            _ = fi127
                        }
                    }
                    _ = fi14
                }
            case -9223372036854775808:
                v135 := []int{32767}
                _ = v135
            }
            _ = v13
        }
        _ = v9
    }
    v3 = (v3 | v3)
    var v136 string = "\n\t"
    v137 := GenS9{F0: 433, F1: 767}
    v136 = v136
    v138 := []int{}
    v3 += (v3 - v3)
    for fi139 := 0; fi139 < 10; fi139++ {
        delete(v1, 4294967295)
        if true {
            switch v136 {
            case "hello":
                switch v136 {
                case "":
                    switch fi139 {
                    case -9223372036854775808:
                        for fi140 := 0; fi140 < 14; fi140++ {
                            switch v136 {
                            case "foo bar":
                                fmt.Println(v136, v136)
                            default:
                                var v141 uint32 = 970
                                _ = v141
                            }
                            if len(v138) > 0 {
                                v138[0] = (fi139 - 127)
                            }
                            _ = fi140
                        }
                    case -1:
                        if (!(true && true)) {
                            v142 := GenS20{F0: 2147483647, F1: -860}
                            v8 = 254
                            _ = v142
                        } else {
                            for fi143 := 0; fi143 < 12; fi143++ {
                                v8 = (v8 >> ((0) & 63))
                                if true {
                                    if ((2 != 18446744073709551615) || false) {
                                        v8 = v8
                                    }
                                } else {
                                    if false {
                                        _ = len(v138)
                                    } else {
                                        _ = len(v138)
                                    }
                                    fmt.Println(v3)
                                    v136 = v136
                                }
                                var v144 uint = 426
                                var v145 float64 = (1e10 - (1.5 + (3.14159 - 1.5)))
                                var v146 int = -301
                                var v147 float64 = -2.25
                                if false {
                                    v144 = (v144 << ((v144) & 63))
                                    if true {
                                        v146 = 128
                                    }
                                    v148 := GenS21{F0: (^(v8)), F1: -1}
                                    _ = v148
                                }
                                var v149 uint32 = 2
                                _ = fi143
                                _ = v144
                                _ = v145
                                _ = v146
                                _ = v147
                                _ = v149
                            }
                        }
                    default:
                        if true {
                            var v150 int32 = 418
                            _ = v150
                        }
                    }
                case "a":
                    for fi151 := 0; fi151 < 2; fi151++ {
                        v4[1] = fi139
                        var v152 float64 = float64((3.14159 * (1.5 + 3.14159)))
                        if (((true && false) || true) || false) {
                            fmt.Println(v152, fi151)
                            if true {
                                var v153 uint64 = 4294967295
                                var v154 float32 = (-2.25 * -2.25)
                                _ = v153
                                _ = v154
                            } else {
                                v136 = ("hello" + "hello")
                                var v155 bool = ((false && (-1 < 2147483647)) && (false || false))
                                var v156 float32 = (1e10 - ((1e10 / 1.0) / -2.25))
                                _ = v155
                                _ = v156
                            }
                        } else {
                            v157 := GenS22{F0: ((("" + v136) + v136) + (("hello" + ("a" + v136)) + (("hello" + (v136 + "foo bar")) + "a"))), F1: 0, F2: (v8 - v8)}
                            _ = v157
                        }
                        _ = fi151
                        _ = v152
                    }
                }
            case "foo bar":
                fmt.Println(v3, v3)
            case "\n\t":
                for fi158 := 0; fi158 < 6; fi158++ {
                    v136 = "a"
                    switch v136 {
                    case "foo bar":
                        v136 = "foo bar"
                    default:
                        var v159 float64 = 0.1
                        _ = v159
                    }
                    v3 -= v3
                    _ = fi158
                }
            default:
                var v160 int16 = -32768
                _ = v160
            }
            delete(v4, 2147483647)
        }
        switch (v136 + ((("\n\t" + v136) + "\n\t") + "")) {
        case "hello":
            for fi161 := 0; fi161 < 15; fi161++ {
                v162 := GenS8{F0: (fi139 + (fi161 | (fi139 << ((4294967295) & 63))))}
                var v163 int32 = -1
                if len(v138) > 0 {
                    v138[0] = fi161
                }
                _ = fi161
                _ = v162
                _ = v163
            }
        default:
            var v164 uint8 = (v8 & genFn0(2, -230))
            _ = v164
        }
        _ = fi139
    }
    var v165 string = v136
    v8 += (v8 - (v8 >> ((1) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v136:", v136)
    fmt.Println("slicelen v138:", len(v138), v138)
    fmt.Println("v165:", v165)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v136
    _ = v137
    _ = v138
    _ = v165
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
