package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint8; F2 int8 }
type GenS4 struct { F0 string; F1 bool; F2 uint8 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS8 struct { F0 uint16; F1 string; F2 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint8; F1 int8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint; F1 uint8 }
type GenS13 struct { F0 uint; F1 int8 }
type GenS14 struct { F0 int16; F1 string; F2 int }
type GenS15 struct { F0 int32; F1 string; F2 int32 }
type GenS16 struct { F0 uint32; F1 int32 }
type GenS17 struct { F0 int8; F1 bool }
type GenS18 struct { F0 uint16; F1 uint64 }
type GenS19 struct { F0 string }
type GenS20 struct { F0 string; F1 uint64 }
type GenS21 struct { F0 uint16 }
type GenS22 struct { F0 int16 }
func genFn0(p0 string, p1 int16) uint64 {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{fi7, fi7, (fi7 | fi7)}
        var v9 bool = false
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: 0, F2: p1}
    v11 := []int{402, -632, -704}
    var v12 int = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 0
}
func genFn1(p0 int16, p1 int) int64 {
    switch "a" {
    case "foo bar":
        if false {
            p1 = p1
            switch "foo bar" {
            case "hello":
                v13 := GenS1{F0: 0, F1: 906, F2: 171}
                _ = v13
            case "":
                if (int16(p1) != 1) {
                    if false {
                        v14 := []int{25, (p1 - p1)}
                        var v15 int = (p1 & (p1 / ((-483) | 1)))
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = v14
                        _ = v15
                    }
                    var v16 int64 = 1
                    var v17 float32 = 0.1
                    _ = v16
                    _ = v17
                } else {
                    p1 += p1
                    p0 = (p0 ^ 1)
                    switch 9223372036854775807 {
                    case 4294967295:
                        for fi18 := 0; fi18 < 4; fi18++ {
                            switch (("\n\t" + "") + "foo bar") {
                            case "hello":
                                v19 := GenS1{F0: 0, F1: 255, F2: (p0 * p0)}
                                _ = v19
                            case "a":
                                for fi20 := 0; fi20 < 2; fi20++ {
                                    p1 += 127
                                    var v21 uint8 = 1
                                    v22 := GenS2{F0: false}
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            case "":
                                v23 := []int{fi18}
                                _ = v23
                            }
                            if ((p0 << ((42) & 63)) == 994) {
                                p1 -= (p1 % ((p1) | 1))
                                for fi24 := 0; fi24 < 10; fi24++ {
                                    v25 := GenS3{F0: -499, F1: 42, F2: -128}
                                    switch "" {
                                    case "":
                                        v26 := []int{(^(fi18)), (p1 ^ fi24), fi18, 32768}
                                        _ = v26
                                    default:
                                        if (!(4 <= 2)) {
                                            p1 = (^(fi18))
                                        } else {
                                            v27 := []int{(p1 & (fi24 | (p1 - (p1 | -2147483648))))}
                                            p1 += (^(fi18))
                                            if (genFn0("\n\t", (-(p0))) >= 955) {
                                                if false {
                                                    var v28 int = 1
                                                    if true {
                                                        var v29 int8 = 42
                                                        _ = v29
                                                    }
                                                    _ = v28
                                                } else {
                                                    v30 := GenS4{F0: "a", F1: false, F2: 50}
                                                    v31 := GenS2{F0: true}
                                                    if (true || ((-722 >= 42) || false)) {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v32 := GenS5{F0: fi24, F1: 808}
                                                            _ = v32
                                                        case "":
                                                            switch fi24 {
                                                            case -343:
                                                                var v33 uint16 = 255
                                                                if len(v27) >= 2 {
                                                                    v27 = v27[1:len(v27)]
                                                                }
                                                                _ = v33
                                                            case 127:
                                                                fmt.Println(p1)
                                                            default:
                                                                v34 := GenS3{F0: 0, F1: 196, F2: 127}
                                                                _ = v34
                                                            }
                                                        default:
                                                            var v35 string = "hello"
                                                            p1 = (p1 - fi24)
                                                            _ = v35
                                                        }
                                                        v27 = append(v27, fi18)
                                                        p1 -= fi24
                                                    } else {
                                                        if (1000 == 32767) {
                                                            fmt.Println(fi18)
                                                        }
                                                        if len(v27) > 0 {
                                                            v27[0] = 1
                                                        }
                                                        var v36 uint8 = 65
                                                        _ = v36
                                                    }
                                                    var v37 uint32 = 138
                                                    _ = v30
                                                    _ = v31
                                                    _ = v37
                                                }
                                                v27 = append(v27, (^(fi18)))
                                            } else {
                                                v38 := GenS2{F0: false}
                                                var v39 float32 = float32((0.0 + -2.25))
                                                _ = v38
                                                _ = v39
                                            }
                                            _ = v27
                                        }
                                    }
                                    var v40 uint = 862
                                    _ = fi24
                                    _ = v25
                                    _ = v40
                                }
                                if false {
                                    v41 := GenS5{F0: p1, F1: 65535}
                                    p1 -= 4294967295
                                    for fi42 := 0; fi42 < 5; fi42++ {
                                        var v43 bool = false
                                        fmt.Println(p1)
                                        var v44 int8 = 42
                                        _ = fi42
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v41
                                }
                            }
                            _ = fi18
                        }
                    case 32767:
                        p1 = (p1 + 32767)
                    case 0:
                        if (true || false) {
                            p0 -= (p0 * p0)
                            if true {
                                if true {
                                    v45 := []int{}
                                    v46 := GenS5{F0: p1, F1: 2147483648}
                                    _ = v45
                                    _ = v46
                                } else {
                                    for fi47 := 0; fi47 < 10; fi47++ {
                                        fmt.Println(p1)
                                        fmt.Println(fi47, fi47)
                                        _ = fi47
                                    }
                                    if (1 <= 336) {
                                        v48 := []int{-1}
                                        var v49 bool = false
                                        for fi50 := 0; fi50 < 20; fi50++ {
                                            for fi51 := 0; fi51 < 5; fi51++ {
                                                switch "a" {
                                                case "hello":
                                                    if (!(fi51 >= (p1 | (fi51 * 2147483647)))) {
                                                        var v52 string = ("\n\t" + "\n\t")
                                                        _ = v52
                                                    } else {
                                                        v49 = v49
                                                        if v49 {
                                                            if len(v48) >= 2 {
                                                                v48 = v48[1:len(v48)]
                                                            }
                                                            v53 := GenS3{F0: 746, F1: 2, F2: 127}
                                                            if v49 {
                                                                var v54 uint16 = 2
                                                                switch fi51 {
                                                                case 310:
                                                                    if v49 {
                                                                        switch fi51 {
                                                                        case 715:
                                                                            var v55 int8 = 0
                                                                            _ = v55
                                                                        default:
                                                                            v49 = (false || v49)
                                                                        }
                                                                        if v49 {
                                                                            v56 := GenS6{F0: 0}
                                                                            if true {
                                                                                v57 := GenS3{F0: 1, F1: 1, F2: -53}
                                                                                _ = v57
                                                                            }
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                var v58 uint32 = 86
                                                                                _ = v58
                                                                            case "":
                                                                                if len(v48) >= 2 {
                                                                                    v48 = v48[1:len(v48)]
                                                                                }
                                                                            case "a":
                                                                                if v49 {
                                                                                    v49 = false
                                                                                    switch (fi51 >> ((18446744073709551615) & 63)) {
                                                                                    case -352:
                                                                                        switch ("hello" + "a") {
                                                                                        case "foo bar":
                                                                                            var v59 uint8 = 255
                                                                                            _ = v59
                                                                                        default:
                                                                                            switch p1 {
                                                                                            case 295:
                                                                                                v49 = ((-(p0)) > (p0 >> ((2) & 63)))
                                                                                            case 4294967295:
                                                                                                v60 := GenS1{F0: 0, F1: 4294967295, F2: p0}
                                                                                                _ = v60
                                                                                            case 9223372036854775807:
                                                                                                fmt.Println(v54, v54, v54)
                                                                                            default:
                                                                                                fmt.Println(fi51, p0)
                                                                                            }
                                                                                            fmt.Println(fi51, fi51)
                                                                                            var v61 int = p1
                                                                                            _ = v61
                                                                                        }
                                                                                    default:
                                                                                        if v49 {
                                                                                            if (true && (!true)) {
                                                                                                v62 := GenS7{F0: genFn0("hello", p0), F1: 18446744073709551615, F2: 4294967295}
                                                                                                _ = v62
                                                                                            }
                                                                                            p1 -= fi51
                                                                                            _ = len(v48)
                                                                                        }
                                                                                    }
                                                                                    _ = len(v48)
                                                                                }
                                                                            }
                                                                            _ = v56
                                                                        } else {
                                                                            v49 = v49
                                                                            v63 := GenS8{F0: (v54 << ((564) & 63)), F1: "hello", F2: 32767}
                                                                            fmt.Println(v54, v49, v54)
                                                                            _ = v63
                                                                        }
                                                                        fmt.Println(p0, p0)
                                                                    }
                                                                case -2147483648:
                                                                    switch ((("" + ((("\n\t" + ("foo bar" + ("" + "a"))) + "a") + ("hello" + "foo bar"))) + "hello") + (("\n\t" + "hello") + "")) {
                                                                    case "\n\t":
                                                                        if (v49 || ((3.14159 - 1e10) >= 1e10)) {
                                                                            v64 := GenS9{F0: (p0 * (-(p0)))}
                                                                            v54 -= v54
                                                                            v65 := GenS10{F0: 1, F1: 42}
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                    case "a":
                                                                        var v66 int8 = 42
                                                                        _ = v66
                                                                    case "hello":
                                                                        var v67 float64 = -2.25
                                                                        _ = v67
                                                                    default:
                                                                        v48 = append(v48, -156)
                                                                    }
                                                                default:
                                                                    v54 = v54
                                                                }
                                                                _ = v54
                                                            }
                                                            _ = v53
                                                        }
                                                        v68 := GenS11{F0: -236}
                                                        if v49 {
                                                            if (0 == 0) {
                                                                switch fi51 {
                                                                case 1:
                                                                    if (1 >= 255) {
                                                                        if v49 {
                                                                            if true {
                                                                                var v69 bool = (!v49)
                                                                                var v70 int64 = 0
                                                                                var v71 int16 = (p0 & p0)
                                                                                _ = v69
                                                                                _ = v70
                                                                                _ = v71
                                                                            } else {
                                                                                if v49 {
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        var v72 uint8 = 2
                                                                                        _ = v72
                                                                                    }
                                                                                    if v49 {
                                                                                        p0 += p0
                                                                                        if v49 {
                                                                                            v73 := GenS8{F0: 1, F1: "", F2: 1}
                                                                                            _ = v73
                                                                                        }
                                                                                        v49 = (-32768 >= p0)
                                                                                    } else {
                                                                                        switch ("a" + "hello") {
                                                                                        case "a":
                                                                                            p1 = (-(fi50))
                                                                                        case "\n\t":
                                                                                            var v74 float64 = (((1.5 / 1.0) + (0.0 / -2.25)) / 1.0)
                                                                                            _ = v74
                                                                                        case "foo bar":
                                                                                            switch ("a" + "a") {
                                                                                            case "hello":
                                                                                                v49 = (42 >= 0)
                                                                                                var v75 uint8 = 86
                                                                                                _ = v75
                                                                                            case "\n\t":
                                                                                                v76 := GenS9{F0: p0}
                                                                                                _ = v76
                                                                                            }
                                                                                        }
                                                                                        var v77 float32 = ((-2.25 / -2.25) / 1.0)
                                                                                        v48 = append(v48, fi51)
                                                                                        _ = v77
                                                                                    }
                                                                                } else {
                                                                                    var v78 float32 = float32(0.0)
                                                                                    _ = v78
                                                                                }
                                                                            }
                                                                            fmt.Println(p1)
                                                                            if len(v48) > 0 {
                                                                                v48[0] = fi50
                                                                            }
                                                                        }
                                                                        var v79 uint = 255
                                                                        v79 = v79
                                                                        _ = v79
                                                                    } else {
                                                                        p0 = p0
                                                                    }
                                                                case 32768:
                                                                    if len(v48) >= 2 {
                                                                        v48 = v48[1:len(v48)]
                                                                    }
                                                                case 2147483647:
                                                                    var v80 int16 = 1000
                                                                    _ = v80
                                                                }
                                                                if (!v49) {
                                                                    p0 = 1
                                                                }
                                                                v49 = (!((-(p0)) != (p0 % ((p0) | 1))))
                                                            } else {
                                                                if v49 {
                                                                    var v81 float64 = ((0.0 * (1e10 * float64(0.1))) + 1e10)
                                                                    var v82 float64 = ((((v81 + (1e10 * (float64(0.1) + (-2.25 - v81)))) * (v81 / (v81 - -2.25))) - v81) * (float64(3.14159) / float64(float32(3.14159))))
                                                                    p1 = fi51
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                                var v83 float32 = (0.1 / 0.1)
                                                                _ = v83
                                                            }
                                                            fmt.Println(p1)
                                                            fmt.Println(fi51, p1, v49)
                                                            v49 = ((1.5 * 0.0) > float32(3.14159))
                                                        } else {
                                                            v84 := GenS12{F0: 617, F1: 0}
                                                            v49 = (true && v49)
                                                            v85 := GenS10{F0: 128, F1: -100}
                                                            if v49 {
                                                                var v86 bool = ((1 < 634) && v49)
                                                                _ = v86
                                                            }
                                                            fmt.Println(fi51)
                                                            switch ("foo bar" + "") {
                                                            case "a":
                                                                var v87 float32 = (1.5 * 3.14159)
                                                                _ = v87
                                                            case "":
                                                                if v49 {
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        var v88 int64 = 9223372036854775807
                                                                        _ = v88
                                                                    case "foo bar":
                                                                        switch 0 {
                                                                        case 32768:
                                                                            fmt.Println(p1, p1)
                                                                        default:
                                                                            if v49 {
                                                                                if v49 {
                                                                                    var v89 int32 = -2147483648
                                                                                    var v90 int8 = 1
                                                                                    _ = v89
                                                                                    _ = v90
                                                                                } else {
                                                                                    v49 = true
                                                                                    if len(v48) > 0 {
                                                                                        v48[0] = p1
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v91 := GenS13{F0: 2, F1: 117}
                                                                                fmt.Println(fi50, p1, v49)
                                                                                _ = v91
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        fmt.Println(fi50, p1, p0)
                                                                    default:
                                                                        switch ("foo bar" + ((("hello" + "\n\t") + ("hello" + "foo bar")) + "\n\t")) {
                                                                        case "foo bar":
                                                                            fmt.Println(fi51)
                                                                        case "\n\t":
                                                                            if (-592 <= 32) {
                                                                                v92 := GenS9{F0: p0}
                                                                                _ = v92
                                                                            }
                                                                        case "":
                                                                            fmt.Println(p0, fi50, v49)
                                                                        }
                                                                    }
                                                                    fmt.Println(p0, v49, p1)
                                                                    _ = len(v48)
                                                                } else {
                                                                    var v93 float64 = 0.0
                                                                    if false {
                                                                        v48 = append(v48, (fi51 >> ((450) & 63)))
                                                                        fmt.Println(p1, v49)
                                                                        if (!(!false)) {
                                                                            if true {
                                                                                fmt.Println(p1)
                                                                            }
                                                                        }
                                                                        p0 += p0
                                                                        var v94 int64 = -9223372036854775808
                                                                        v48 = append(v48, (fi51 * fi51))
                                                                        var v95 uint = 558
                                                                        _ = v94
                                                                        _ = v95
                                                                    } else {
                                                                        v96 := GenS6{F0: 20}
                                                                        switch fi50 {
                                                                        case 32767:
                                                                            v97 := GenS1{F0: 18446744073709551615, F1: 65535, F2: (p0 + p0)}
                                                                            _ = v97
                                                                        case -320:
                                                                            var v98 int16 = p0
                                                                            _ = v98
                                                                        default:
                                                                            var v99 uint16 = 137
                                                                            _ = v99
                                                                        }
                                                                        var v100 uint = 4294967295
                                                                        _ = v100
                                                                        _ = v96
                                                                    }
                                                                    _ = v93
                                                                }
                                                                if (32767 < -702) {
                                                                    v49 = (9223372036854775807 < -1)
                                                                    v48 = append(v48, fi51)
                                                                    if true {
                                                                        v101 := GenS10{F0: 220, F1: -1}
                                                                        if len(v48) > 0 {
                                                                            v48[0] = (fi50 + -2147483648)
                                                                        }
                                                                        v102 := GenS6{F0: 128}
                                                                        _ = v101
                                                                        _ = v102
                                                                    }
                                                                }
                                                            default:
                                                                if (!(v49 && (432 != (p0 ^ (p0 * (p0 * 42)))))) {
                                                                    p0 = (p0 ^ -1)
                                                                    var v103 uint16 = 65535
                                                                    _ = v103
                                                                }
                                                            }
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                        p1 -= fi51
                                                        p0 -= (p0 | p0)
                                                        if len(v48) >= 2 {
                                                            v48 = v48[1:len(v48)]
                                                        }
                                                        _ = v68
                                                    }
                                                case "a":
                                                    _ = len(v48)
                                                case "":
                                                    var v104 float64 = ((0.1 - ((((((0.0 / 1.0) / (1e10 * ((1.5 / -2.25) + ((0.1 + float64(3.14159)) * 0.0)))) + -2.25) / 1e10) - 0.1) / 1.0)) / 1.0)
                                                    _ = v104
                                                }
                                                v105 := GenS14{F0: (p0 + (p0 << ((2) & 63))), F1: "", F2: (-(p1))}
                                                _ = fi51
                                                _ = v105
                                            }
                                            if (-1 != 42) {
                                                fmt.Println(fi50)
                                                switch (p1 * (p1 % ((fi50) | 1))) {
                                                case 128:
                                                    fmt.Println(v49, p0)
                                                case 9223372036854775807:
                                                    if (!v49) {
                                                        var v106 int8 = -128
                                                        v106 = v106
                                                        _ = v106
                                                    }
                                                }
                                                var v107 int = fi50
                                                _ = v107
                                            }
                                            switch "\n\t" {
                                            case "\n\t":
                                                if true {
                                                    if len(v48) >= 2 {
                                                        v48 = v48[1:len(v48)]
                                                    }
                                                    p0 += p0
                                                } else {
                                                    var v108 int64 = 42
                                                    if ((192 < 128) || (!false)) {
                                                        v49 = false
                                                        fmt.Println(v108, v108)
                                                        if true {
                                                            v109 := GenS3{F0: -1, F1: 127, F2: 127}
                                                            _ = v109
                                                        } else {
                                                            v110 := GenS13{F0: 1, F1: 64}
                                                            if len(v48) >= 2 {
                                                                v48 = v48[1:len(v48)]
                                                            }
                                                            fmt.Println(p0)
                                                            _ = v110
                                                        }
                                                    }
                                                    var v111 int64 = 42
                                                    _ = v108
                                                    _ = v111
                                                }
                                            }
                                            _ = fi50
                                        }
                                        _ = v48
                                        _ = v49
                                    }
                                    if (!((!true) && (2147483647 <= 0))) {
                                        p1 = p1
                                        v112 := []int{}
                                        var v113 float64 = -2.25
                                        _ = v112
                                        _ = v113
                                    } else {
                                        for fi114 := 0; fi114 < 3; fi114++ {
                                            if true {
                                                var v115 bool = true
                                                _ = v115
                                            } else {
                                                switch (fi114 & (^(p1))) {
                                                case -650:
                                                    v116 := GenS15{F0: 984, F1: ("hello" + ("hello" + "foo bar")), F2: -1}
                                                    _ = v116
                                                case 32768:
                                                    v117 := GenS10{F0: 255, F1: -1}
                                                    _ = v117
                                                case -663:
                                                    p0 = p0
                                                default:
                                                    fmt.Println(p1, p1, fi114)
                                                }
                                                var v118 bool = false
                                                _ = v118
                                            }
                                            if (1.5 > 1.5) {
                                                var v119 uint64 = 42
                                                _ = v119
                                            }
                                            v120 := []int{-9223372036854775808, p1, (p1 | p1)}
                                            _ = fi114
                                            _ = v120
                                        }
                                        fmt.Println(p0)
                                        for fi121 := 0; fi121 < 1; fi121++ {
                                            fmt.Println(p0, fi121, p1)
                                            switch -2147483648 {
                                            case 42:
                                                if (42 < 4294967295) {
                                                    p0 = (p0 >> ((4294967295) & 63))
                                                    var v122 int32 = 0
                                                    v123 := GenS16{F0: 2147483648, F1: v122}
                                                    _ = v122
                                                    _ = v123
                                                } else {
                                                    v124 := GenS17{F0: 9, F1: (1e10 != (-2.25 + 0.0))}
                                                    _ = v124
                                                }
                                            }
                                            _ = fi121
                                        }
                                    }
                                }
                                v125 := []int{p1, (p1 * p1), int(p0), p1}
                                for fi126 := 0; fi126 < 12; fi126++ {
                                    p0 -= p0
                                    _ = fi126
                                }
                                _ = v125
                            }
                            switch "\n\t" {
                            case "a":
                                v127 := []int{(p1 & p1), p1, (p1 ^ p1)}
                                _ = v127
                            case "\n\t":
                                v128 := []int{(p1 % ((p1) | 1)), (p1 | (p1 - (p1 - p1)))}
                                _ = v128
                            case "":
                                fmt.Println(p0, p0)
                            }
                        } else {
                            switch (p1 << ((42) & 63)) {
                            case 0:
                                var v129 int16 = 42
                                _ = v129
                            case 32767:
                                fmt.Println(p1, p0)
                            default:
                                if false {
                                    p1 = (p1 - (p1 << ((802) & 63)))
                                }
                            }
                        }
                    }
                }
            case "a":
                var v130 int = (p1 * p1)
                _ = v130
            default:
                v131 := GenS4{F0: ("\n\t" + "\n\t"), F1: (!(1 >= -1)), F2: 255}
                _ = v131
            }
            for fi132 := 0; fi132 < 5; fi132++ {
                if false {
                    var v133 int = fi132
                    _ = v133
                } else {
                    v134 := []int{42, p1}
                    p1 += (fi132 & fi132)
                    if (true && true) {
                        switch (("a" + "foo bar") + "a") {
                        case "\n\t":
                            fmt.Println(fi132, p0, fi132)
                        case "foo bar":
                            p0 += p0
                        case "hello":
                            v135 := GenS5{F0: fi132, F1: 480}
                            _ = v135
                        default:
                            p1 -= (p1 | fi132)
                        }
                        v136 := GenS18{F0: 2, F1: 4294967295}
                        for fi137 := 0; fi137 < 9; fi137++ {
                            if len(v134) >= 2 {
                                v134 = v134[1:len(v134)]
                            }
                            _ = fi137
                        }
                        _ = v136
                    }
                    _ = v134
                }
                for fi138 := 0; fi138 < 10; fi138++ {
                    if (9223372036854775807 < -43) {
                        var v139 int = (fi132 - (fi138 * 128))
                        var v140 uint32 = 396
                        _ = v139
                        _ = v140
                    } else {
                        v141 := []int{fi138, (fi138 / ((-194) | 1)), (p1 | (-(p1)))}
                        _ = v141
                    }
                    _ = fi138
                }
                var v142 float64 = 0.1
                _ = fi132
                _ = v142
            }
        } else {
            var v143 int32 = -128
            switch (p1 + p1) {
            case -616:
                v144 := []int{p1, (p1 - p1), 1, 128}
                _ = v144
            case -540:
                var v145 uint8 = 2
                _ = v145
            case -210:
                for fi146 := 0; fi146 < 5; fi146++ {
                    if false {
                        v147 := GenS7{F0: 18446744073709551615, F1: 4294967295, F2: 0}
                        v148 := GenS19{F0: "a"}
                        _ = v147
                        _ = v148
                    } else {
                        for fi149 := 0; fi149 < 2; fi149++ {
                            v150 := []int{fi149, fi149, fi146, p1}
                            _ = fi149
                            _ = v150
                        }
                    }
                    _ = fi146
                }
            default:
                switch 2147483647 {
                case 32767:
                    var v151 int = (p1 / (((^(p1))) | 1))
                    _ = v151
                default:
                    if false {
                        v143 = (v143 * v143)
                        var v152 bool = ("foo bar" >= ("" + ""))
                        for fi153 := 0; fi153 < 20; fi153++ {
                            fmt.Println(fi153, fi153)
                            _ = fi153
                        }
                        _ = v152
                    } else {
                        var v154 bool = (("foo bar" + "a") == "foo bar")
                        switch (p1 / ((-977) | 1)) {
                        case 2147483647:
                            var v155 int32 = (v143 - (v143 % ((42) | 1)))
                            _ = v155
                        default:
                            v143 += (v143 * -2147483648)
                        }
                        v156 := GenS15{F0: (v143 - v143), F1: "hello", F2: v143}
                        _ = v154
                        _ = v156
                    }
                }
            }
            switch "foo bar" {
            case "":
                fmt.Println(v143, v143, p1)
            }
            _ = v143
        }
    case "a":
        fmt.Println(p0, p1)
    default:
        var v157 string = "a"
        _ = v157
    }
    v158 := []int{(p1 - (p1 & p1)), (p1 + p1), int(p0)}
    var v159 int16 = (^(p0))
    _ = p0
    _ = p1
    _ = v158
    _ = v159
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    v3 := []int{}
    var v4 float64 = 0.0
    v5 := make(map[int]string)
    var v6 uint16 = 0
    switch "\n\t" {
    case "":
        if ((!((42 == 938) && (((v6 >> ((18446744073709551615) & 63)) != 2) || (0 >= genFn0("", 42))))) || ("hello" >= "\n\t")) {
            if ((0 < -1) && false) {
                v160 := GenS12{F0: 0, F1: 179}
                for fi161 := 0; fi161 < 2; fi161++ {
                    var v162 uint = 255
                    v162 -= (v162 & 184)
                    _ = fi161
                    _ = v162
                }
                _ = v160
            }
            v4 = v4
            var v163 bool = false
            _ = v163
        } else {
            if (float64(0.0) <= v4) {
                v5[2147483647] = "foo bar"
                v164 := GenS20{F0: "hello", F1: 1}
                _ = v164
            }
        }
    case "hello":
        v165 := v5[-580]
        _ = v165
    default:
        fmt.Println(v4, v4)
    }
    switch 430 {
    case -9223372036854775808:
        v4 += ((1e10 * 0.1) / v4)
    }
    if true {
        fmt.Println(v6, v4)
        v166 := GenS14{F0: -378, F1: "\n\t", F2: 128}
        for fi167 := 0; fi167 < 12; fi167++ {
            var v168 bool = true
            v4 = v4
            _ = fi167
            _ = v168
        }
        _ = v166
    }
    for fi169 := 0; fi169 < 5; fi169++ {
        v170 := GenS21{F0: 2}
        v6 = 399
        _ = fi169
        _ = v170
    }
    if ((v4 + v4) != v4) {
        var v171 bool = true
        _ = v171
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = v4
    if false {
        v4 = 1e10
        if false {
            v6 = 255
            v6 = (v6 >> ((2) & 63))
            v4 += v4
        } else {
            v4 = v4
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v172 uint64 = 367
            _ = v172
        }
        if false {
            switch 1 {
            case -1:
                if false {
                    fmt.Println(v6)
                    if (0 <= 2) {
                        var v173 uint32 = 2147483648
                        v173 -= (^(v173))
                        var v174 bool = false
                        _ = v173
                        _ = v174
                    }
                    var v175 uint = 42
                    _ = v175
                }
            case -382:
                v6 = (^(v6))
            default:
                var v176 uint8 = 1
                _ = v176
            }
            if false {
                v177 := GenS9{F0: 1000}
                var v178 float32 = (1e10 * 0.0)
                _ = v177
                _ = v178
            } else {
                var v179 float32 = 1e10
                fmt.Println(v179)
                _ = v179
            }
        }
    } else {
        v180 := GenS22{F0: -1}
        _ = v180
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v6)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
