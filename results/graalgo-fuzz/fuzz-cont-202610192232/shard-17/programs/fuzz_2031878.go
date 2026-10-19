package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint; F2 bool }
type GenS5 struct { F0 bool; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 uint }
type GenS10 struct { F0 int32; F1 string }
type GenS11 struct { F0 string; F1 uint; F2 int16 }
type GenS12 struct { F0 uint64; F1 uint64; F2 bool }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 string; F1 uint64 }
type GenS16 struct { F0 bool; F1 uint8; F2 int }
type GenS17 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS18 struct { F0 string; F1 uint64 }
type GenS19 struct { F0 int; F1 uint16 }
type GenS20 struct { F0 uint; F1 string; F2 int16 }

func main() {
    var v1 int8 = -1
    var v2 int64 = -891
    var v3 bool = true
    var v4 int64 = (v2 & (v2 | -9223372036854775808))
    switch 42 {
    case -61:
        v3 = (1.5 == (float32(1e10) * (0.0 / -2.25)))
    default:
        v5 := []int{}
        _ = v5
    }
    fmt.Println(v1)
    var v6 int64 = (v2 << ((65535) & 63))
    for fi7 := 0; fi7 < 2; fi7++ {
        switch fi7 {
        case 4294967295:
            v8 := []int{}
            _ = v8
        case 32767:
            v9 := GenS1{F0: 18446744073709551615, F1: -1}
            _ = v9
        }
        switch "" {
        case "foo bar":
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 int = (fi10 / ((829) | 1))
                _ = fi10
                _ = v11
            }
        case "\n\t":
            for fi12 := 0; fi12 < 10; fi12++ {
                if false {
                    var v13 int = 0
                    v13 = v13
                    _ = v13
                }
                if false {
                    v14 := []int{fi12, (fi12 * (^(fi7))), (fi7 & fi7), fi7}
                    v1 -= (v1 | v1)
                    v6 = 1
                    _ = v14
                } else {
                    if v3 {
                        if v3 {
                            var v15 float64 = (float64(0.1) / 1.0)
                            var v16 uint32 = 924
                            switch (fi12 % (((fi12 ^ fi12)) | 1)) {
                            case -901:
                                v16 += 2
                            case 755:
                                var v17 uint64 = 42
                                _ = v17
                            case 4294967295:
                                var v18 string = ("foo bar" + "foo bar")
                                _ = v18
                            default:
                                if true {
                                    var v19 uint = 18446744073709551615
                                    v20 := []int{(fi7 - fi7), -9223372036854775808, (fi12 + (fi7 ^ (fi12 << (((v19 ^ v19)) & 63))))}
                                    var v21 uint32 = 2
                                    var v22 uint = 2
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                } else {
                                    var v23 int16 = 0
                                    if (v3 && v3) {
                                        fmt.Println(v15, fi7, v1)
                                        var v24 float64 = float64(0.0)
                                        var v25 float32 = 3.14159
                                        _ = v24
                                        _ = v25
                                    } else {
                                        v26 := []int{}
                                        _ = v26
                                    }
                                    switch ((("" + "") + "a") + "") {
                                    case "a":
                                        v23 += v23
                                    case "hello":
                                        var v27 float64 = 1e10
                                        _ = v27
                                    default:
                                        if (false || (v3 || false)) {
                                            v28 := GenS2{F0: uint16(v16)}
                                            _ = v28
                                        } else {
                                            v3 = (42 == 120)
                                            v1 -= int8(v16)
                                            if v3 {
                                                var v29 uint = 0
                                                v29 = 65535
                                                _ = v29
                                            }
                                            var v30 int64 = v2
                                            v31 := []int{(fi7 | 32768), (fi12 << ((42) & 63)), fi12}
                                            var v32 int16 = (v23 >> ((255) & 63))
                                            v33 := GenS2{F0: 301}
                                            var v34 bool = ((v32 ^ (v32 % ((1) | 1))) < 719)
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                        }
                                    }
                                    switch -9223372036854775808 {
                                    case 4294967295:
                                        v1 = v1
                                    case 32768:
                                        if false {
                                            var v35 int16 = (v23 + (-(v23)))
                                            _ = v35
                                        } else {
                                            v36 := GenS3{F0: (true || (v3 || true)), F1: (v16 & v16)}
                                            _ = v36
                                        }
                                    case 977:
                                        fmt.Println(v23, v15, fi7)
                                    default:
                                        if v3 {
                                            if ((v23 * -32768) > v23) {
                                                if v3 {
                                                    var v37 int32 = -1
                                                    _ = v37
                                                }
                                            } else {
                                                switch ("foo bar" + "hello") {
                                                case "hello":
                                                    var v38 int64 = (v4 ^ v2)
                                                    _ = v38
                                                case "a":
                                                    v4 += (v6 % ((v6) | 1))
                                                }
                                            }
                                            fmt.Println(v6, fi12)
                                            v4 = (v6 % ((v4) | 1))
                                        }
                                    }
                                    if (("" != ("hello" + ("" + ("foo bar" + "\n\t")))) || true) {
                                        v16 = (v16 + (v16 | v16))
                                        v39 := []int{(fi7 << ((65535) & 63)), (fi12 * (fi7 & fi7)), 32767, (fi7 * 127)}
                                        _ = v39
                                    } else {
                                        var v40 uint64 = 18446744073709551615
                                        if false {
                                            var v41 bool = v3
                                            v42 := GenS1{F0: 678, F1: v23}
                                            v4 -= v6
                                            switch "hello" {
                                            case "a":
                                                v43 := GenS4{F0: 428, F1: 153, F2: v41}
                                                _ = v43
                                            case "foo bar":
                                                v6 = int64(fi7)
                                            }
                                            v16 = (v16 / (((v16 & v16)) | 1))
                                            v44 := GenS3{F0: v41, F1: v16}
                                            v2 -= (v2 - int64(v40))
                                            _ = v41
                                            _ = v42
                                            _ = v44
                                        }
                                        _ = v40
                                    }
                                    _ = v23
                                }
                            }
                            _ = v15
                            _ = v16
                        }
                        if v3 {
                            v1 += v1
                            v2 -= (^(v6))
                        } else {
                            fmt.Println(v1, fi12)
                            v45 := []int{(fi12 ^ fi12)}
                            _ = v45
                        }
                        v46 := []int{}
                        _ = v46
                    }
                    if (1e10 == (3.14159 / (float32(1e10) / ((0.1 + -2.25) - 0.0)))) {
                        v47 := GenS1{F0: 862, F1: -32768}
                        v48 := []int{fi12, 128, (fi7 % ((1) | 1))}
                        if true {
                            var v49 int32 = 1
                            _ = v49
                        } else {
                            v3 = true
                            v50 := GenS5{F0: v3, F1: -1000, F2: 255}
                            _ = v50
                        }
                        switch 32768 {
                        case 9223372036854775807:
                            if false {
                                v6 = v2
                                v1 -= (-(v1))
                            }
                            var v51 uint = 2
                            v3 = true
                            _ = v51
                        }
                        if len(v48) > 0 {
                            v48[0] = -442
                        }
                        _ = v47
                        _ = v48
                    }
                    var v52 int64 = v6
                    fmt.Println(v4, v4)
                    if v3 {
                        v6 += v2
                        if true {
                            var v53 int16 = 42
                            if (!(127 != 0)) {
                                v54 := GenS5{F0: (42 == (v1 << ((0) & 63))), F1: 1, F2: 2}
                                var v55 uint8 = 255
                                if false {
                                    if (true && v3) {
                                        switch "hello" {
                                        case "a":
                                            v52 += (v4 >> ((709) & 63))
                                        default:
                                            fmt.Println(v1)
                                        }
                                        v56 := GenS1{F0: 18446744073709551615, F1: 1000}
                                        if false {
                                            var v57 string = ""
                                            var v58 uint8 = 2
                                            _ = v57
                                            _ = v58
                                        }
                                        v3 = (false || v3)
                                        _ = v56
                                    } else {
                                        v1 = v1
                                        var v59 string = "hello"
                                        v59 = (v59 + v59)
                                        _ = v59
                                    }
                                    var v60 int16 = (v53 % (((v53 / ((v53) | 1))) | 1))
                                    var v61 uint64 = 42
                                    if (true || (!true)) {
                                        var v62 int64 = (^(v52))
                                        var v63 int32 = 1
                                        v64 := GenS6{F0: (v61 | 0), F1: 591, F2: v55}
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                    }
                                    if (((v53 & (^(v53))) <= v60) || false) {
                                        v60 = 0
                                    }
                                    _ = v60
                                    _ = v61
                                } else {
                                    var v65 uint16 = 0
                                    if true {
                                        if (float64((float32((float64((3.14159 / 1.5)) + (((float64(((0.1 / 1.0) - 1.5)) + 1.5) + (((0.0 / 1e10) / (((float64(((((0.0 * -2.25) / 1.0) / -2.25) - -2.25)) / -2.25) / 0.1) + 0.1)) * 1.5)) * 3.14159))) - 1e10)) <= (-2.25 + 1e10)) {
                                            var v66 uint8 = (v55 | v55)
                                            fmt.Println(v53)
                                            _ = v66
                                        } else {
                                            if (!false) {
                                                switch ("" + "hello") {
                                                case "":
                                                    if (!(!v3)) {
                                                        var v67 string = "foo bar"
                                                        v68 := GenS1{F0: 4294967295, F1: (^(v53))}
                                                        v67 = (v67 + v67)
                                                        _ = v67
                                                        _ = v68
                                                    } else {
                                                        fmt.Println(v52)
                                                        switch (fi12 - 611) {
                                                        case 0:
                                                            if v3 {
                                                                v69 := GenS7{F0: 0, F1: -188}
                                                                if v3 {
                                                                    if ((3.14159 + 1e10) > (1e10 - (-2.25 / -2.25))) {
                                                                        if v3 {
                                                                            if (2 < 1) {
                                                                                v2 = (v6 - (^(v2)))
                                                                                switch ("hello" + ("" + "foo bar")) {
                                                                                case "a":
                                                                                    v70 := []int{(fi12 % (((fi7 << ((42) & 63))) | 1)), (fi12 & fi12), fi7}
                                                                                    _ = v70
                                                                                case "":
                                                                                    var v71 int8 = v1
                                                                                    _ = v71
                                                                                case "hello":
                                                                                    v72 := []int{(-(fi7))}
                                                                                    _ = v72
                                                                                }
                                                                                var v73 uint8 = 2
                                                                                _ = v73
                                                                            } else {
                                                                                v2 = (-(v52))
                                                                                v74 := GenS1{F0: 4294967295, F1: (-(v53))}
                                                                                v75 := GenS8{F0: -1}
                                                                                _ = v74
                                                                                _ = v75
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v53)
                                                                            v76 := GenS9{F0: 121, F1: 619}
                                                                            v2 += v2
                                                                            fmt.Println(v4)
                                                                            _ = v76
                                                                        }
                                                                        if ("" <= ("foo bar" + (("" + "a") + (("a" + ("a" + (((((("\n\t" + "a") + "\n\t") + ((((("\n\t" + "\n\t") + "hello") + "a") + "a") + "a")) + (("hello" + "a") + "hello")) + "") + ("foo bar" + "hello")))) + ("hello" + "foo bar"))))) {
                                                                            fmt.Println(v52, v53, v2)
                                                                            if (((1 == -79) && (!(1e10 >= float32(0.0)))) || (!((^(v2)) >= (v2 ^ (^(v6)))))) {
                                                                                if ((!(!(4294967295 < 0))) && ((v53 >> ((50) & 63)) < -615)) {
                                                                                    v3 = true
                                                                                    var v77 string = ""
                                                                                    v3 = (-2147483648 > -2147483648)
                                                                                    _ = v77
                                                                                } else {
                                                                                    v2 += v6
                                                                                    if ((int16(v55) < v53) || (!(false || v3))) {
                                                                                        v78 := GenS2{F0: (v65 & v65)}
                                                                                        if v3 {
                                                                                            v1 = 127
                                                                                            var v79 int16 = (v53 ^ 32767)
                                                                                            v53 -= v53
                                                                                            _ = v79
                                                                                        }
                                                                                        _ = v78
                                                                                    }
                                                                                    switch "" {
                                                                                    case "":
                                                                                        v3 = false
                                                                                    }
                                                                                    var v80 int16 = v53
                                                                                    if (!true) {
                                                                                        var v81 uint = 18446744073709551615
                                                                                        var v82 float64 = 0.1
                                                                                        _ = v81
                                                                                        _ = v82
                                                                                    } else {
                                                                                        switch (fi7 / ((fi12) | 1)) {
                                                                                        case 127:
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                var v83 int32 = -732
                                                                                                _ = v83
                                                                                            case "a":
                                                                                                v80 -= (v80 | (v53 | 42))
                                                                                            case "hello":
                                                                                                var v84 int = (fi12 + (fi12 + fi7))
                                                                                                _ = v84
                                                                                            }
                                                                                            v2 -= 0
                                                                                        }
                                                                                        if v3 {
                                                                                            var v85 int32 = int32(v65)
                                                                                            _ = v85
                                                                                        } else {
                                                                                            v86 := GenS10{F0: -1, F1: "\n\t"}
                                                                                            _ = v86
                                                                                        }
                                                                                        if v3 {
                                                                                            fmt.Println(v65, fi7, v53)
                                                                                            v87 := []int{(-(fi7)), (fi12 - 212), -572}
                                                                                            _ = v87
                                                                                        }
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            v88 := []int{fi7, fi7, fi7, (fi7 % ((fi7) | 1))}
                                                                                            _ = v88
                                                                                        case "a":
                                                                                            v89 := GenS11{F0: "", F1: 65535, F2: v53}
                                                                                            _ = v89
                                                                                        case "":
                                                                                            var v90 float64 = 3.14159
                                                                                            _ = v90
                                                                                        }
                                                                                    }
                                                                                    _ = v80
                                                                                }
                                                                                if v3 {
                                                                                    fmt.Println(v53, v2)
                                                                                }
                                                                                fmt.Println(v3)
                                                                                v91 := []int{(fi7 << ((14) & 63)), (fi12 % (((fi7 - fi7)) | 1)), fi7}
                                                                                var v92 uint16 = v65
                                                                                _ = v91
                                                                                _ = v92
                                                                            } else {
                                                                                var v93 int32 = 0
                                                                                _ = v93
                                                                            }
                                                                            v94 := []int{}
                                                                            _ = v94
                                                                        }
                                                                    } else {
                                                                        if v3 {
                                                                            var v95 float64 = ((1.5 + (1e10 - 3.14159)) - (1e10 * (3.14159 / 1.5)))
                                                                            fmt.Println(v4, v55)
                                                                            _ = v95
                                                                        }
                                                                        v96 := []int{(fi7 ^ fi7), -2147483648}
                                                                        _ = v96
                                                                    }
                                                                    v55 = v55
                                                                    v53 = -431
                                                                    var v97 int8 = v1
                                                                    v98 := GenS9{F0: v1, F1: 710}
                                                                    v99 := GenS9{F0: v1, F1: 0}
                                                                    _ = v97
                                                                    _ = v98
                                                                    _ = v99
                                                                }
                                                                _ = v69
                                                            } else {
                                                                if v3 {
                                                                    v100 := GenS11{F0: "hello", F1: 0, F2: (v53 >> ((42) & 63))}
                                                                    switch (fi12 % (((fi12 * (fi12 % ((fi7) | 1)))) | 1)) {
                                                                    case -532:
                                                                        if true {
                                                                            var v101 int = fi7
                                                                            var v102 float32 = float32(0.1)
                                                                            v1 += (^(v1))
                                                                            _ = v101
                                                                            _ = v102
                                                                        }
                                                                    case 127:
                                                                        if true {
                                                                            fmt.Println(fi7, v3, v55)
                                                                        }
                                                                    case 2147483647:
                                                                        v52 = v4
                                                                    default:
                                                                        v103 := GenS12{F0: 4294967295, F1: 428, F2: (-388 == -1)}
                                                                        _ = v103
                                                                    }
                                                                    if ((539 != 0) || v3) {
                                                                        v104 := []int{fi12, fi12}
                                                                        _ = v104
                                                                    } else {
                                                                        v105 := GenS13{F0: 42}
                                                                        v106 := GenS14{F0: v1}
                                                                        var v107 float32 = (1.5 / -2.25)
                                                                        _ = v105
                                                                        _ = v106
                                                                        _ = v107
                                                                    }
                                                                    _ = v100
                                                                } else {
                                                                    v108 := GenS3{F0: false, F1: 903}
                                                                    v2 += (^(v4))
                                                                    _ = v108
                                                                }
                                                            }
                                                        case 658:
                                                            switch 127 {
                                                            case 2147483647:
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    var v109 float32 = 0.0
                                                                    _ = v109
                                                                case "foo bar":
                                                                    fmt.Println(v52)
                                                                }
                                                            }
                                                            var v110 int16 = (v53 * v53)
                                                            _ = v110
                                                        case 32768:
                                                            var v111 uint8 = v55
                                                            _ = v111
                                                        default:
                                                            v112 := GenS15{F0: "\n\t", F1: 2}
                                                            _ = v112
                                                        }
                                                    }
                                                case "foo bar":
                                                    v55 -= 42
                                                case "\n\t":
                                                    v6 = int64(v1)
                                                }
                                                fmt.Println(v65, v1)
                                            }
                                            fmt.Println(v52, v65)
                                        }
                                        switch "hello" {
                                        case "hello":
                                            switch "foo bar" {
                                            case "hello":
                                                var v113 uint64 = 0
                                                _ = v113
                                            }
                                        case "foo bar":
                                            var v114 uint16 = (^(v65))
                                            _ = v114
                                        }
                                        switch -9223372036854775808 {
                                        case 759:
                                            if v3 {
                                                var v115 uint32 = 42
                                                _ = v115
                                            } else {
                                                var v116 string = ("a" + ("\n\t" + "a"))
                                                _ = v116
                                            }
                                        case 1:
                                            if v3 {
                                                if v3 {
                                                    var v117 uint32 = 766
                                                    fmt.Println(v53, v2, v117)
                                                    if v3 {
                                                        v52 = (v2 | v4)
                                                        v6 = 42
                                                    }
                                                    _ = v117
                                                } else {
                                                    v53 = 385
                                                    var v118 uint64 = 1
                                                    v119 := []int{}
                                                    _ = v118
                                                    _ = v119
                                                }
                                                v53 += (v53 % (((^(v53))) | 1))
                                                v2 -= v6
                                                v120 := []int{(fi7 / ((fi7) | 1)), 32767}
                                                _ = v120
                                            } else {
                                                var v121 uint = 331
                                                if (v121 <= 255) {
                                                    v65 = (v65 >> (((v121 / ((v121) | 1))) & 63))
                                                    v122 := []int{(fi7 >> (((v121 ^ 2)) & 63))}
                                                    fmt.Println(v2)
                                                    v121 = v121
                                                    _ = v122
                                                }
                                                _ = v121
                                            }
                                        case 2147483647:
                                            v123 := GenS16{F0: (v65 > (v65 / ((uint16(v55)) | 1))), F1: 255, F2: 32768}
                                            _ = v123
                                        }
                                    }
                                    switch ("foo bar" + "a") {
                                    case "":
                                        fmt.Println(v6, fi12, v3)
                                        switch "" {
                                        case "foo bar":
                                            switch (fi12 / (((^(fi7))) | 1)) {
                                            case 127:
                                                if (!true) {
                                                    v124 := GenS17{F0: 65535, F1: v65, F2: (v65 << ((665) & 63))}
                                                    if (358 != (v52 & (v52 ^ v4))) {
                                                        var v125 bool = false
                                                        if ((!v125) || (-2147483648 <= 2147483647)) {
                                                            v126 := []int{fi12, fi12, fi12}
                                                            if v3 {
                                                                v53 += v53
                                                                var v127 int = fi7
                                                                v6 = 0
                                                                _ = v127
                                                            }
                                                            v6 -= (v4 >> ((0) & 63))
                                                            _ = v126
                                                        }
                                                        _ = v125
                                                    }
                                                    fmt.Println(v53, v1)
                                                    v4 += (^(v6))
                                                    switch "" {
                                                    case "hello":
                                                        if ((v65 == (v65 + 0)) || ((v3 || false) || ((fi7 ^ fi7) != fi12))) {
                                                            if v3 {
                                                                v128 := []int{fi7}
                                                                var v129 int = 2147483647
                                                                if (!(int32(v55) < -1)) {
                                                                    switch ("foo bar" + ("\n\t" + ("" + "a"))) {
                                                                    case "":
                                                                        fmt.Println(v55)
                                                                    }
                                                                    if true {
                                                                        if false {
                                                                            v1 = 1
                                                                            v129 = -9223372036854775808
                                                                        } else {
                                                                            if (true || (32767 < (v129 - (fi12 >> ((899) & 63))))) {
                                                                                var v130 string = "\n\t"
                                                                                _ = v130
                                                                            }
                                                                        }
                                                                        v53 = (v53 * v53)
                                                                    }
                                                                    v52 -= v4
                                                                } else {
                                                                    var v131 int8 = v1
                                                                    var v132 uint8 = (v55 - v55)
                                                                    _ = v131
                                                                    _ = v132
                                                                }
                                                                switch "a" {
                                                                case "":
                                                                    switch fi7 {
                                                                    case 286:
                                                                        v2 += (v6 / ((v52) | 1))
                                                                    case 127:
                                                                        if v3 {
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                var v133 int8 = v1
                                                                                _ = v133
                                                                            case "a":
                                                                                fmt.Println(v6, v65)
                                                                            default:
                                                                                fmt.Println(v3, v2)
                                                                            }
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    switch v129 {
                                                                    case 32768:
                                                                        fmt.Println(v52)
                                                                    }
                                                                default:
                                                                    if (!(-2147483648 <= 983)) {
                                                                        v2 = (v52 ^ v52)
                                                                        var v134 uint64 = 42
                                                                        _ = v134
                                                                    }
                                                                }
                                                                _ = v128
                                                                _ = v129
                                                            } else {
                                                                v6 += (v52 - v2)
                                                                v52 += v52
                                                                if false {
                                                                    v135 := []int{}
                                                                    _ = v135
                                                                }
                                                            }
                                                            v136 := GenS18{F0: "\n\t", F1: 472}
                                                            fmt.Println(v52, v53, fi12)
                                                            _ = v136
                                                        }
                                                    default:
                                                        v137 := GenS9{F0: (v1 << ((4294967295) & 63)), F1: 365}
                                                        _ = v137
                                                    }
                                                    v138 := GenS19{F0: 0, F1: (^(v65))}
                                                    fmt.Println(v3, fi7)
                                                    _ = v124
                                                    _ = v138
                                                } else {
                                                    if true {
                                                        v139 := GenS19{F0: 128, F1: v65}
                                                        fmt.Println(v3, fi12, fi7)
                                                        _ = v139
                                                    }
                                                }
                                            }
                                        case "a":
                                            var v140 bool = v3
                                            _ = v140
                                        }
                                        if (0 >= 4294967295) {
                                            v141 := GenS19{F0: fi12, F1: v65}
                                            fmt.Println(v53)
                                            _ = v141
                                        }
                                    case "\n\t":
                                        if (false && (false && (!v3))) {
                                            v142 := GenS7{F0: 65535, F1: v53}
                                            _ = v142
                                        } else {
                                            v3 = (int64(v55) < v52)
                                            v143 := GenS10{F0: int32(v1), F1: (("hello" + "hello") + "")}
                                            var v144 int = fi12
                                            _ = v143
                                            _ = v144
                                        }
                                    case "a":
                                        v145 := GenS13{F0: 113}
                                        _ = v145
                                    default:
                                        var v146 uint64 = 0
                                        _ = v146
                                    }
                                    _ = v65
                                }
                                var v147 uint16 = 792
                                _ = v147
                                _ = v54
                                _ = v55
                            } else {
                                v148 := []int{(fi7 | 4294967295), fi12}
                                var v149 string = ""
                                var v150 int64 = (v6 - v6)
                                _ = v148
                                _ = v149
                                _ = v150
                            }
                            var v151 int64 = (v52 & (v2 & (v52 - (v52 / ((512) | 1)))))
                            _ = v151
                            _ = v53
                        }
                        v52 -= (-(v4))
                    } else {
                        var v152 string = ("" + "a")
                        _ = v152
                    }
                    _ = v52
                }
                _ = fi12
            }
        default:
            fmt.Println(v4, v4, v1)
        }
        var v153 uint32 = 42
        _ = fi7
        _ = v153
    }
    var v154 uint = 1
    v6 = v6
    fmt.Println(v1)
    if ((float64(1.5) <= float64((float32(-2.25) * 0.1))) && ((!(1 == 4294967295)) || true)) {
        var v155 uint16 = 876
        var v156 uint = (v154 ^ (^(v154)))
        _ = v155
        _ = v156
    }
    switch (("foo bar" + ("hello" + "foo bar")) + "a") {
    case "foo bar":
        for fi157 := 0; fi157 < 3; fi157++ {
            v158 := GenS20{F0: (v154 % ((v154) | 1)), F1: "\n\t", F2: 32767}
            if (v3 && (v3 && (42 != 564))) {
                var v159 uint = 33
                var v160 int16 = 1
                var v161 uint64 = 0
                _ = v159
                _ = v160
                _ = v161
            } else {
                switch fi157 {
                case 9223372036854775807:
                    var v162 float64 = (1e10 + 3.14159)
                    _ = v162
                case 127:
                    v163 := []int{}
                    _ = v163
                case -9223372036854775808:
                    v3 = true
                }
            }
            _ = fi157
            _ = v158
        }
    case "":
        var v164 uint32 = 306
        _ = v164
    default:
        v3 = (((v154 << (((v154 - (^(v154)))) & 63)) > (v154 + (^(v154)))) && (2147483647 == -1))
    }
    v1 -= (v1 & 42)
    var v165 float32 = -2.25
    v4 = 961
    fmt.Println(v154)
    fmt.Println("v1:", v1)
    fmt.Println("v154:", v154)
    fmt.Println("v165:", v165)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v154
    _ = v165
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
