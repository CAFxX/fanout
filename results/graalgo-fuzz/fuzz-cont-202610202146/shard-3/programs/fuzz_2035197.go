package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int32; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32; F1 uint8; F2 uint }
type GenS10 struct { F0 uint8; F1 uint8 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint16; F1 uint16 }
type GenS13 struct { F0 int8; F1 uint8 }
type GenS14 struct { F0 bool; F1 uint32; F2 int }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 uint; F1 uint; F2 int32 }
type GenS18 struct { F0 int; F1 string; F2 int }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 int64 }
type GenS21 struct { F0 int32; F1 int32; F2 int64 }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 int32; F1 int32 }
type GenS24 struct { F0 uint64; F1 int32 }
type GenS25 struct { F0 int16; F1 uint32; F2 uint }
type GenS26 struct { F0 string; F1 uint8; F2 uint8 }
type GenS27 struct { F0 int8; F1 string }
type GenS28 struct { F0 string; F1 int32 }
type GenS29 struct { F0 uint8; F1 int }
type GenS30 struct { F0 uint16; F1 uint8 }
type GenS31 struct { F0 int32 }
type GenS32 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS33 struct { F0 uint8 }
type GenS34 struct { F0 string }
type GenS35 struct { F0 uint64; F1 bool }
func genFn0(p0 int64, p1 int8) int {
    v4 := GenS1{F0: (2 > 1)}
    var v5 float64 = 0.0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 9223372036854775807
}
func genFn1(p0 int32) int64 {
    p0 -= 392
    var v6 uint32 = 42
    v7 := GenS2{F0: (-(p0)), F1: 42, F2: 813}
    var v8 uint32 = v6
    for fi9 := 0; fi9 < 5; fi9++ {
        p0 += 0
        var v10 bool = false
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 842
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    var v11 int16 = 1
    if false {
        var v12 uint = 65535
        var v13 uint16 = 0
        var v14 int32 = -1
        _ = v12
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 15; fi16++ {
                v11 -= -10
                delete(v2, "foo bar")
                var v17 uint = 572
                if false {
                    fmt.Println(v11, fi15, v17)
                    v18 := GenS3{F0: 42, F1: 0}
                    switch "" {
                    case "\n\t":
                        var v19 uint16 = 65535
                        _ = v19
                    default:
                        v17 -= 18446744073709551615
                    }
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
            if true {
                v20 := []int{genFn0(0, 42), fi15}
                var v21 int = (^(v1))
                _ = v20
                _ = v21
            }
            _ = fi15
        }
        if (!((3.14159 * (3.14159 * (1e10 * 3.14159))) >= (1e10 - 3.14159))) {
            fmt.Println(v11, v11)
            var v22 int64 = genFn1(679)
            _ = v22
        }
    }
    v23 := GenS2{F0: -584, F1: -2147483648, F2: 114}
    v24 := GenS2{F0: 855, F1: -470, F2: 42}
    v25 := []int{(v1 & 32768), (^(v1)), genFn0(0, 0), (v1 | (v1 / (((v1 % (((v1 % ((822) | 1))) | 1))) | 1)))}
    for fi26 := 0; fi26 < 20; fi26++ {
        v11 = 0
        var v27 uint32 = 989
        _ = fi26
        _ = v27
    }
    switch "foo bar" {
    case "a":
        if true {
            v28 := GenS4{F0: v1}
            _ = v28
        }
    case "hello":
        fmt.Println(v1, v11)
    case "":
        if false {
            var v29 float32 = ((float32(0.1) / (1e10 / 1e10)) - 3.14159)
            _ = v29
        } else {
            v11 += -588
        }
    default:
        var v30 uint = 0
        _ = v30
    }
    var v31 int64 = 9223372036854775807
    v1 = 631
    fmt.Println(v11, v1)
    switch ("hello" + "") {
    case "":
        if false {
            for fi32 := 0; fi32 < 17; fi32++ {
                var v33 int32 = 2147483647
                _ = len(v3)
                _ = fi32
                _ = v33
            }
        }
    default:
        if (0 >= 42) {
            if (907 > 32767) {
                if ((^(v31)) == (v31 & -963)) {
                    for fi34 := 0; fi34 < 13; fi34++ {
                        var v35 uint16 = 2
                        if ((2 < 887) && ((!(0 < 18446744073709551615)) && false)) {
                            v1 += genFn0(-729, -93)
                            var v36 uint8 = 5
                            _ = v36
                        }
                        _ = fi34
                        _ = v35
                    }
                    if ((0 < 42) && false) {
                        for fi37 := 0; fi37 < 6; fi37++ {
                            if false {
                                if (((((1e10 + (0.0 / 1.0)) - (-2.25 / (3.14159 - 0.1))) + float32(1.5)) - 1e10) == 0.0) {
                                    var v38 uint = 0
                                    _ = v38
                                }
                            }
                            _ = fi37
                        }
                        switch 9223372036854775807 {
                        case -2147483648:
                            for fi39 := 0; fi39 < 18; fi39++ {
                                if (240 > 32767) {
                                    v40 := GenS5{F0: 0}
                                    v41 := GenS6{F0: "a"}
                                    _ = v40
                                    _ = v41
                                }
                                _ = fi39
                            }
                        case 819:
                            switch "hello" {
                            case "\n\t":
                                v42 := GenS6{F0: ""}
                                _ = v42
                            case "a":
                                for fi43 := 0; fi43 < 14; fi43++ {
                                    switch "" {
                                    case "foo bar":
                                        for fi44 := 0; fi44 < 18; fi44++ {
                                            var v45 float64 = (1.5 + 0.0)
                                            fmt.Println(fi44, fi43)
                                            _ = fi44
                                            _ = v45
                                        }
                                    }
                                    switch -9223372036854775808 {
                                    case 4294967295:
                                        fmt.Println(v11)
                                    case 9223372036854775807:
                                        for fi46 := 0; fi46 < 18; fi46++ {
                                            if (((100 <= 74) || (730 >= 711)) && true) {
                                                if (!(!false)) {
                                                    delete(v3, ("hello" + ("foo bar" + (("hello" + (("hello" + "foo bar") + "hello")) + "hello"))))
                                                    var v47 int16 = (v11 % (((v11 - -1)) | 1))
                                                    var v48 int16 = (v47 % ((1000) | 1))
                                                    _ = v47
                                                    _ = v48
                                                }
                                                v11 += v11
                                            } else {
                                                switch v1 {
                                                case 9223372036854775807:
                                                    v49 := GenS6{F0: "hello"}
                                                    _ = v49
                                                case 128:
                                                    v31 = v31
                                                    _ = len(v3)
                                                }
                                            }
                                            _ = fi46
                                        }
                                    case 127:
                                        for fi50 := 0; fi50 < 17; fi50++ {
                                            if true {
                                                if len(v25) >= 2 {
                                                    v25 = v25[1:len(v25)]
                                                }
                                                v2[""] = 2147483648
                                            } else {
                                                v11 -= v11
                                                var v51 int = (fi43 * (fi50 | fi50))
                                                fmt.Println(v11, v51)
                                                fmt.Println(fi50, v1)
                                                _ = v51
                                            }
                                            delete(v2, ((("" + "a") + ((("hello" + "a") + "\n\t") + ("" + (("\n\t" + "") + "")))) + ("hello" + ((("\n\t" + "hello") + "\n\t") + "foo bar"))))
                                            v31 -= (v31 - v31)
                                            if false {
                                                var v52 uint64 = 1
                                                fmt.Println(v31)
                                                var v53 float64 = 1.5
                                                _ = v52
                                                _ = v53
                                            } else {
                                                var v54 float32 = (0.0 / 1.5)
                                                var v55 int16 = v11
                                                _ = v54
                                                _ = v55
                                            }
                                            _ = fi50
                                        }
                                    default:
                                        v11 += v11
                                    }
                                    _ = fi43
                                }
                            }
                        case -609:
                            switch "" {
                            case "a":
                                if false {
                                    if len(v25) > 0 {
                                        v25[0] = v1
                                    }
                                } else {
                                    if (2 != 23) {
                                        if true {
                                            if (true && (42 <= 255)) {
                                                var v56 int8 = -1
                                                var v57 uint32 = 2
                                                _ = v56
                                                _ = v57
                                            }
                                            fmt.Println(v31, v1, v31)
                                            fmt.Println(v31, v1, v31)
                                        } else {
                                            fmt.Println(v1)
                                            var v58 float32 = 1.5
                                            _ = v58
                                        }
                                        switch (("\n\t" + ("" + "hello")) + "\n\t") {
                                        case "hello":
                                            for fi59 := 0; fi59 < 18; fi59++ {
                                                switch "hello" {
                                                case "":
                                                    var v60 int = (v1 >> ((2) & 63))
                                                    _ = v60
                                                case "foo bar":
                                                    v11 = v11
                                                case "hello":
                                                    for fi61 := 0; fi61 < 14; fi61++ {
                                                        if (-943 <= 2147483647) {
                                                            v62 := GenS1{F0: false}
                                                            var v63 uint16 = 2
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            v1 -= (-(fi61))
                                                            var v64 uint32 = 612
                                                            if ("hello" == "") {
                                                                v11 = 734
                                                            }
                                                            fmt.Println(fi61, fi59, v11)
                                                            _ = v64
                                                        }
                                                        if (false && (false || (genFn1(-1) == (v31 << ((2) & 63))))) {
                                                            switch genFn0(v31, 127) {
                                                            case -1:
                                                                var v65 uint64 = 2
                                                                fmt.Println(fi59)
                                                                _ = v65
                                                            case 2147483647:
                                                                if len(v25) >= 2 {
                                                                    v25 = v25[1:len(v25)]
                                                                }
                                                            default:
                                                                if false {
                                                                    var v66 int = (fi59 % ((fi61) | 1))
                                                                    var v67 uint8 = 0
                                                                    _ = v66
                                                                    _ = v67
                                                                } else {
                                                                    var v68 int = (fi61 - (fi59 / (((fi59 + fi61)) | 1)))
                                                                    v69 := GenS7{F0: (2147483647 == 4294967295), F1: (v11 | (v11 << ((752) & 63)))}
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                            }
                                                            var v70 uint16 = 42
                                                            _ = v70
                                                        }
                                                        v31 += (^(v31))
                                                        _ = fi61
                                                    }
                                                default:
                                                    fmt.Println(v11, fi59)
                                                }
                                                var v71 int32 = 42
                                                _ = fi59
                                                _ = v71
                                            }
                                        case "":
                                            var v72 uint64 = 59
                                            _ = v72
                                        case "a":
                                            fmt.Println(v1, v1, v1)
                                        }
                                    } else {
                                        if (1 != 0) {
                                            if false {
                                                switch genFn0(-1, 1) {
                                                case 0:
                                                    for fi73 := 0; fi73 < 15; fi73++ {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            fmt.Println(v1, v11)
                                                        case "\n\t":
                                                            switch "a" {
                                                            case "hello":
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        var v74 uint8 = 128
                                                                        _ = v74
                                                                    case "\n\t":
                                                                        var v75 int32 = 42
                                                                        _ = v75
                                                                    case "a":
                                                                        v11 -= 28
                                                                    default:
                                                                        if (true && (42 == 591)) {
                                                                            if (!true) {
                                                                                if false {
                                                                                    for fi76 := 0; fi76 < 17; fi76++ {
                                                                                        v1 = 9223372036854775807
                                                                                        _ = fi76
                                                                                    }
                                                                                    v77 := GenS4{F0: -671}
                                                                                    v1 -= fi73
                                                                                    _ = v77
                                                                                } else {
                                                                                    for fi78 := 0; fi78 < 6; fi78++ {
                                                                                        var v79 uint16 = 697
                                                                                        v79 += v79
                                                                                        v80 := GenS5{F0: 42}
                                                                                        _ = fi78
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                    }
                                                                                    if (4294967295 < 4294967295) {
                                                                                        if (("" + "") <= "foo bar") {
                                                                                            if len(v25) >= 2 {
                                                                                                v25 = v25[1:len(v25)]
                                                                                            }
                                                                                            _ = len(v25)
                                                                                        } else {
                                                                                            var v81 uint16 = 699
                                                                                            if (!false) {
                                                                                                var v82 int16 = 42
                                                                                                var v83 bool = ((v11 | (^(v11))) > 1)
                                                                                                switch "foo bar" {
                                                                                                case "a":
                                                                                                    var v84 float32 = 3.14159
                                                                                                    _ = v84
                                                                                                case "\n\t":
                                                                                                    if false {
                                                                                                        var v85 float64 = 3.14159
                                                                                                        _ = v85
                                                                                                    } else {
                                                                                                        for fi86 := 0; fi86 < 6; fi86++ {
                                                                                                            var v87 float64 = -2.25
                                                                                                            _ = fi86
                                                                                                            _ = v87
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = v82
                                                                                                _ = v83
                                                                                            } else {
                                                                                                v25 = append(v25, v1)
                                                                                            }
                                                                                            _ = v81
                                                                                        }
                                                                                        v11 = int16(v31)
                                                                                    } else {
                                                                                        for fi88 := 0; fi88 < 19; fi88++ {
                                                                                            var v89 bool = true
                                                                                            var v90 float32 = 1.5
                                                                                            switch fi88 {
                                                                                            case 32768:
                                                                                                v25 = append(v25, (fi73 ^ v1))
                                                                                            case -1:
                                                                                                v90 -= ((((0.1 - 1.5) / (v90 - float32((0.1 - 1.5)))) - (3.14159 * v90)) / 1e10)
                                                                                            }
                                                                                            _ = fi88
                                                                                            _ = v89
                                                                                            _ = v90
                                                                                        }
                                                                                        var v91 string = ("foo bar" + "hello")
                                                                                        v92 := GenS4{F0: 1}
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                    }
                                                                                }
                                                                                v3["foo bar"] = 0
                                                                            }
                                                                            if (((^(v11)) < v11) && false) {
                                                                                var v93 float32 = (0.1 * 1.5)
                                                                                if false {
                                                                                    var v94 int64 = (v31 ^ (^(v31)))
                                                                                    for fi95 := 0; fi95 < 18; fi95++ {
                                                                                        v1 -= 9223372036854775807
                                                                                        fmt.Println(fi73)
                                                                                        _ = fi95
                                                                                    }
                                                                                    _ = v94
                                                                                } else {
                                                                                    fmt.Println(v1, fi73, v93)
                                                                                    var v96 int = -2147483648
                                                                                    _ = v96
                                                                                }
                                                                                _ = v93
                                                                            }
                                                                        }
                                                                    }
                                                                case "":
                                                                    switch v1 {
                                                                    case 1:
                                                                        for fi97 := 0; fi97 < 6; fi97++ {
                                                                            var v98 float64 = (float64((0.1 * (0.0 + -2.25))) / 3.14159)
                                                                            _ = fi97
                                                                            _ = v98
                                                                        }
                                                                    case 246:
                                                                        v99 := GenS8{F0: 18446744073709551615}
                                                                        _ = v99
                                                                    case 963:
                                                                        if (759 < 42) {
                                                                            v100 := GenS8{F0: 1}
                                                                            if ((!true) && false) {
                                                                                var v101 uint64 = 2
                                                                                _ = v101
                                                                            } else {
                                                                                fmt.Println(v31)
                                                                            }
                                                                            _ = v100
                                                                        } else {
                                                                            if (("a" + ("hello" + "foo bar")) > "hello") {
                                                                                var v102 uint16 = 32767
                                                                                if true {
                                                                                    for fi103 := 0; fi103 < 4; fi103++ {
                                                                                        delete(v3, ("foo bar" + ""))
                                                                                        switch 974 {
                                                                                        case -2147483648:
                                                                                            if ((3.14159 / (-2.25 / float64(((3.14159 * -2.25) * (1e10 * -2.25))))) > 0.1) {
                                                                                                v104 := GenS4{F0: (fi103 - 746)}
                                                                                                switch (fi103 % (((-(fi103))) | 1)) {
                                                                                                case 128:
                                                                                                    if (v11 == (-(v11))) {
                                                                                                        var v105 string = ("a" + "")
                                                                                                        v105 = "foo bar"
                                                                                                        _ = v105
                                                                                                    } else {
                                                                                                        var v106 uint64 = 0
                                                                                                        _ = v106
                                                                                                    }
                                                                                                case -990:
                                                                                                    _ = len(v25)
                                                                                                }
                                                                                                switch ("foo bar" + ("\n\t" + "")) {
                                                                                                case "hello":
                                                                                                    switch "hello" {
                                                                                                    case "hello":
                                                                                                        if len(v25) > 0 {
                                                                                                            v25[0] = (fi103 - (fi103 % ((32768) | 1)))
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        v107 := GenS8{F0: 65535}
                                                                                                        _ = v107
                                                                                                    }
                                                                                                default:
                                                                                                    v102 += (v102 * (v102 >> ((1) & 63)))
                                                                                                }
                                                                                                fmt.Println(fi73, v1)
                                                                                                v108 := GenS9{F0: 1, F1: 160, F2: 4294967295}
                                                                                                _ = v104
                                                                                                _ = v108
                                                                                            }
                                                                                            var v109 uint32 = 2147483647
                                                                                            _ = v109
                                                                                        case -1:
                                                                                            if true {
                                                                                                v1 = int(v102)
                                                                                            }
                                                                                        default:
                                                                                            if ((3.14159 - float32(1.5)) == -2.25) {
                                                                                                v110 := GenS9{F0: 0, F1: 2, F2: 0}
                                                                                                v1 = (fi73 & 128)
                                                                                                if (((1.5 - (0.0 + 1.5)) / 0.1) <= 3.14159) {
                                                                                                    switch "a" {
                                                                                                    case "a":
                                                                                                        v1 = (fi103 % ((fi103) | 1))
                                                                                                    case "hello":
                                                                                                        var v111 float64 = (0.0 * (3.14159 + 1e10))
                                                                                                        fmt.Println(v31, fi103, v11)
                                                                                                        _ = v111
                                                                                                    default:
                                                                                                        var v112 uint = 764
                                                                                                        _ = v112
                                                                                                    }
                                                                                                    v11 -= v11
                                                                                                    if (2 != 0) {
                                                                                                        v113 := v3[(("\n\t" + ("a" + ("\n\t" + ""))) + "a")]
                                                                                                        fmt.Println(v113, v1, v31)
                                                                                                        _ = v113
                                                                                                    } else {
                                                                                                        fmt.Println(v1)
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(fi73, v1)
                                                                                                    var v114 uint64 = 1
                                                                                                    var v115 float64 = (((((0.0 + (1.5 * (float64((0.1 * 1.5)) / ((3.14159 + 3.14159) - 1e10)))) * ((((3.14159 + 1e10) / 1e10) * -2.25) / 1e10)) / -2.25) + -2.25) / 1.0)
                                                                                                    _ = v114
                                                                                                    _ = v115
                                                                                                }
                                                                                                _ = v110
                                                                                            }
                                                                                        }
                                                                                        _ = len(v2)
                                                                                        if true {
                                                                                            var v116 uint64 = 0
                                                                                            switch "hello" {
                                                                                            case "":
                                                                                                v116 += v116
                                                                                                v117 := v3[("foo bar" + "hello")]
                                                                                                v1 = genFn0(v31, -1)
                                                                                                _ = v117
                                                                                            case "hello":
                                                                                                var v118 uint64 = (v116 * v116)
                                                                                                _ = v118
                                                                                            case "foo bar":
                                                                                                var v119 int64 = v31
                                                                                                _ = v119
                                                                                            }
                                                                                            if ((fi73 ^ (fi103 | -942)) <= (^(v1))) {
                                                                                                v120 := GenS2{F0: -1, F1: 623, F2: (^(v116))}
                                                                                                switch 1 {
                                                                                                case 4294967295:
                                                                                                    v116 += 4294967295
                                                                                                default:
                                                                                                    v25 = append(v25, 32767)
                                                                                                }
                                                                                                var v121 int16 = 662
                                                                                                _ = v120
                                                                                                _ = v121
                                                                                            }
                                                                                            _ = v116
                                                                                        }
                                                                                        if (true || true) {
                                                                                            v102 = (v102 / (((v102 >> ((53) & 63))) | 1))
                                                                                            v122 := GenS10{F0: 255, F1: 42}
                                                                                            fmt.Println(fi103)
                                                                                            var v123 int16 = (v11 * 32767)
                                                                                            _ = v122
                                                                                            _ = v123
                                                                                        } else {
                                                                                            switch "a" {
                                                                                            case "a":
                                                                                                var v124 uint64 = 18446744073709551615
                                                                                                _ = v124
                                                                                            default:
                                                                                                v1 += fi103
                                                                                            }
                                                                                            fmt.Println(v102, v31, v11)
                                                                                            v125 := GenS11{F0: "a"}
                                                                                            _ = v125
                                                                                        }
                                                                                        fmt.Println(v31)
                                                                                        _ = fi103
                                                                                    }
                                                                                    v1 += v1
                                                                                    switch (("a" + ("hello" + "")) + "a") {
                                                                                    case "foo bar":
                                                                                        if true {
                                                                                            var v126 float32 = (1e10 + 3.14159)
                                                                                            _ = len(v3)
                                                                                            fmt.Println(v102, v102, v126)
                                                                                            _ = v126
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v127 := GenS8{F0: 65535}
                                                                                        _ = v127
                                                                                    case "hello":
                                                                                        for fi128 := 0; fi128 < 14; fi128++ {
                                                                                            if (v11 <= 42) {
                                                                                                if (v102 == v102) {
                                                                                                    fmt.Println(fi73, fi73, v11)
                                                                                                    var v129 bool = false
                                                                                                    v129 = v129
                                                                                                    _ = v129
                                                                                                }
                                                                                            }
                                                                                            v102 -= 32767
                                                                                            if (!false) {
                                                                                                if (true || true) {
                                                                                                    v102 += (v102 ^ v102)
                                                                                                    v130 := GenS12{F0: (v102 & v102), F1: 255}
                                                                                                    _ = v130
                                                                                                } else {
                                                                                                    v1 += fi128
                                                                                                    fmt.Println(v11, v102)
                                                                                                    delete(v2, (("a" + "hello") + "foo bar"))
                                                                                                }
                                                                                                var v131 uint8 = 2
                                                                                                fmt.Println(v11)
                                                                                                _ = v131
                                                                                            }
                                                                                            _ = fi128
                                                                                        }
                                                                                    default:
                                                                                        for fi132 := 0; fi132 < 17; fi132++ {
                                                                                            var v133 float32 = 0.0
                                                                                            var v134 int = v1
                                                                                            _ = fi132
                                                                                            _ = v133
                                                                                            _ = v134
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v102
                                                                            }
                                                                            for fi135 := 0; fi135 < 20; fi135++ {
                                                                                fmt.Println(v11, fi135)
                                                                                v11 -= (v11 >> ((0) & 63))
                                                                                switch "a" {
                                                                                case "":
                                                                                    v11 = (v11 % (((v11 + v11)) | 1))
                                                                                }
                                                                                _ = fi135
                                                                            }
                                                                        }
                                                                    default:
                                                                        v136 := GenS3{F0: 1, F1: 2}
                                                                        _ = v136
                                                                    }
                                                                default:
                                                                    switch (v1 / (((v1 % ((genFn0(genFn1(1), 1)) | 1))) | 1)) {
                                                                    case 32767:
                                                                        for fi137 := 0; fi137 < 20; fi137++ {
                                                                            v138 := GenS3{F0: -128, F1: 621}
                                                                            if false {
                                                                                _ = len(v25)
                                                                                v139 := GenS13{F0: 42, F1: 83}
                                                                                _ = v139
                                                                            }
                                                                            fmt.Println(v31, fi73, v1)
                                                                            _ = fi137
                                                                            _ = v138
                                                                        }
                                                                    case 1:
                                                                        switch (fi73 - genFn0((^(v31)), -84)) {
                                                                        case -537:
                                                                            switch v1 {
                                                                            case 559:
                                                                                switch (("\n\t" + "foo bar") + (("\n\t" + ("foo bar" + "a")) + "a")) {
                                                                                case "a":
                                                                                    var v140 uint32 = 2147483648
                                                                                    _ = v140
                                                                                }
                                                                            }
                                                                        case -60:
                                                                            v141 := GenS1{F0: ((1 >= 4294967295) && true)}
                                                                            _ = v141
                                                                        case 4294967295:
                                                                            v25 = append(v25, (fi73 & fi73))
                                                                        default:
                                                                            fmt.Println(v1)
                                                                        }
                                                                    case 128:
                                                                        v1 = genFn0(222, int8(v1))
                                                                    default:
                                                                        v142 := GenS11{F0: "hello"}
                                                                        _ = v142
                                                                    }
                                                                }
                                                            case "a":
                                                                v143 := GenS14{F0: (-2.25 >= -2.25), F1: 42, F2: (v1 >> ((1) & 63))}
                                                                _ = v143
                                                            }
                                                        case "hello":
                                                            if false {
                                                                for fi144 := 0; fi144 < 13; fi144++ {
                                                                    v25 = append(v25, genFn0((v31 % ((-760) | 1)), -78))
                                                                    var v145 string = ("a" + "\n\t")
                                                                    var v146 uint = 42
                                                                    _ = fi144
                                                                    _ = v145
                                                                    _ = v146
                                                                }
                                                                var v147 int64 = (v31 & (v31 % (((v31 | (v31 >> ((4294967295) & 63)))) | 1)))
                                                                _ = v147
                                                            }
                                                        }
                                                        for fi148 := 0; fi148 < 14; fi148++ {
                                                            if (false || ((!(v1 > (fi73 & 32767))) || false)) {
                                                                v149 := GenS8{F0: 42}
                                                                var v150 float64 = ((-2.25 / -2.25) + 0.1)
                                                                _ = v149
                                                                _ = v150
                                                            } else {
                                                                v151 := GenS1{F0: true}
                                                                _ = v151
                                                            }
                                                            _ = fi148
                                                        }
                                                        _ = fi73
                                                    }
                                                default:
                                                    switch (v1 | (v1 & (v1 & (v1 & v1)))) {
                                                    case 32768:
                                                        switch (v1 & 2147483647) {
                                                        case 32768:
                                                            for fi152 := 0; fi152 < 8; fi152++ {
                                                                var v153 uint = 65535
                                                                switch (fi152 >> ((4294967295) & 63)) {
                                                                case 1:
                                                                    var v154 string = ""
                                                                    _ = v154
                                                                case 42:
                                                                    v11 -= (v11 % (((v11 >> ((2) & 63))) | 1))
                                                                }
                                                                for fi155 := 0; fi155 < 15; fi155++ {
                                                                    v156 := v3[""]
                                                                    switch 128 {
                                                                    case 746:
                                                                        var v157 int8 = 42
                                                                        _ = v157
                                                                    case 142:
                                                                        fmt.Println(v156, v156)
                                                                    case -2147483648:
                                                                        if false {
                                                                            v1 += int(v153)
                                                                            var v158 int8 = 1
                                                                            v156 = (v156 >> (((v153 / ((v153) | 1))) & 63))
                                                                            switch (v1 & v1) {
                                                                            case 42:
                                                                                v158 += 0
                                                                            }
                                                                            if false {
                                                                                var v159 int = (fi152 & (fi155 - genFn0((-(v31)), v158)))
                                                                                fmt.Println(v1, v11, fi155)
                                                                                var v160 int = 1
                                                                                _ = v159
                                                                                _ = v160
                                                                            } else {
                                                                                _ = len(v25)
                                                                                var v161 int = 32767
                                                                                if (genFn1(42) >= (v31 + (v31 >> (((v153 + v153)) & 63)))) {
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        if (179 == 1) {
                                                                                            if false {
                                                                                                v162 := GenS15{F0: (^(v156))}
                                                                                                fmt.Println(v158)
                                                                                                v153 += 1
                                                                                                var v163 int16 = 0
                                                                                                _ = v162
                                                                                                _ = v163
                                                                                            } else {
                                                                                                v164 := GenS7{F0: true, F1: v11}
                                                                                                var v165 int64 = (v31 & v31)
                                                                                                _ = v164
                                                                                                _ = v165
                                                                                            }
                                                                                            var v166 int64 = v31
                                                                                            _ = v166
                                                                                        } else {
                                                                                            var v167 uint = v153
                                                                                            if len(v25) > 0 {
                                                                                                v25[0] = (v161 & -711)
                                                                                            }
                                                                                            _ = v167
                                                                                        }
                                                                                    case "hello":
                                                                                        if false {
                                                                                            _ = len(v25)
                                                                                            fmt.Println(v161)
                                                                                            var v168 float32 = 3.14159
                                                                                            v2["a"] = v156
                                                                                            _ = v168
                                                                                        } else {
                                                                                            v1 = (v161 / (((-(v1))) | 1))
                                                                                        }
                                                                                    default:
                                                                                        v31 = genFn1(555)
                                                                                    }
                                                                                } else {
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v158, v11, fi155)
                                                                                    case "\n\t":
                                                                                        v169 := GenS16{F0: ("foo bar" + (("foo bar" + "hello") + "foo bar"))}
                                                                                        v170 := GenS17{F0: (v153 % (((v153 ^ (v153 >> ((v153) & 63)))) | 1)), F1: v153, F2: 623}
                                                                                        fmt.Println(v161, v161)
                                                                                        _ = v169
                                                                                        _ = v170
                                                                                    default:
                                                                                        v153 += (v153 >> ((v153) & 63))
                                                                                    }
                                                                                }
                                                                                _ = v161
                                                                            }
                                                                            _ = v158
                                                                        } else {
                                                                            v171 := GenS7{F0: false, F1: 42}
                                                                            _ = v171
                                                                        }
                                                                    }
                                                                    if (v153 <= (v153 + (v153 >> ((255) & 63)))) {
                                                                        var v172 int = int(v153)
                                                                        _ = v172
                                                                    } else {
                                                                        if len(v25) >= 2 {
                                                                            v25 = v25[1:len(v25)]
                                                                        }
                                                                    }
                                                                    _ = fi155
                                                                    _ = v156
                                                                }
                                                                _ = fi152
                                                                _ = v153
                                                            }
                                                        case -435:
                                                            v173 := GenS5{F0: -1}
                                                            _ = v173
                                                        default:
                                                            v31 -= v31
                                                        }
                                                    case -1:
                                                        var v174 int8 = -128
                                                        _ = v174
                                                    case 952:
                                                        switch int(v11) {
                                                        case 0:
                                                            for fi175 := 0; fi175 < 8; fi175++ {
                                                                var v176 bool = false
                                                                _ = fi175
                                                                _ = v176
                                                            }
                                                        case -38:
                                                            v177 := v2["foo bar"]
                                                            _ = v177
                                                        case 973:
                                                            if (784 > 42) {
                                                                v31 = 0
                                                                var v178 uint8 = 0
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v31 = -9223372036854775808
                                                                }
                                                                _ = v178
                                                            }
                                                        default:
                                                            if true {
                                                                if false {
                                                                    delete(v2, "")
                                                                } else {
                                                                    var v179 uint = 65535
                                                                    _ = v179
                                                                }
                                                                v31 += v31
                                                                var v180 int64 = v31
                                                                _ = v180
                                                            }
                                                        }
                                                    default:
                                                        v181 := GenS16{F0: "hello"}
                                                        _ = v181
                                                    }
                                                }
                                                v182 := GenS18{F0: v1, F1: (("hello" + "") + ""), F2: 2147483647}
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v31 += (v31 - (v31 % ((0) | 1)))
                                                case "a":
                                                    if ((!true) && (1 <= 42)) {
                                                        if (((0.0 - 0.1) < 0.1) || (!(v31 > -759))) {
                                                            fmt.Println(v11, v11, v31)
                                                            v11 = v11
                                                            var v183 int = genFn0(genFn1(400), 41)
                                                            _ = v183
                                                        }
                                                        var v184 int8 = -1
                                                        if (401 != 42) {
                                                            switch v1 {
                                                            case -2147483648:
                                                                v184 -= 0
                                                            }
                                                            v2["a"] = 42
                                                        } else {
                                                            switch (v1 - genFn0(1, (v184 % ((127) | 1)))) {
                                                            case 32768:
                                                                for fi185 := 0; fi185 < 13; fi185++ {
                                                                    v3[""] = 475
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        if true {
                                                                            v186 := GenS19{F0: 2}
                                                                            v1 = fi185
                                                                            _ = v186
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            var v187 uint64 = 923
                                                                            v188 := GenS14{F0: (!true), F1: 0, F2: fi185}
                                                                            _ = v187
                                                                            _ = v188
                                                                        }
                                                                    }
                                                                    var v189 int16 = 0
                                                                    _ = fi185
                                                                    _ = v189
                                                                }
                                                            case -9223372036854775808:
                                                                for fi190 := 0; fi190 < 8; fi190++ {
                                                                    v191 := GenS20{F0: -9223372036854775808}
                                                                    _ = fi190
                                                                    _ = v191
                                                                }
                                                            }
                                                            if false {
                                                                v25 = append(v25, v1)
                                                            }
                                                        }
                                                        _ = v184
                                                    } else {
                                                        for fi192 := 0; fi192 < 1; fi192++ {
                                                            v1 -= (v1 ^ v1)
                                                            for fi193 := 0; fi193 < 6; fi193++ {
                                                                if (-2.25 == 1.5) {
                                                                    v31 = (v31 % ((genFn1(2147483647)) | 1))
                                                                    if false {
                                                                        _ = len(v3)
                                                                        v1 = (v1 << ((2) & 63))
                                                                        v31 -= v31
                                                                    } else {
                                                                        fmt.Println(fi193, fi193)
                                                                        if true {
                                                                            switch v1 {
                                                                            case 128:
                                                                                _ = len(v25)
                                                                            default:
                                                                                v11 = (v11 / ((v11) | 1))
                                                                            }
                                                                            fmt.Println(fi192, v31, v1)
                                                                            fmt.Println(v31, v31)
                                                                        } else {
                                                                            var v194 uint64 = 2
                                                                            v31 = -1
                                                                            var v195 string = "foo bar"
                                                                            _ = v194
                                                                            _ = v195
                                                                        }
                                                                        if (("hello" + "\n\t") < "a") {
                                                                            v1 -= 636
                                                                            if false {
                                                                                if (genFn0(7, 1) <= fi192) {
                                                                                    if (!false) {
                                                                                        v31 -= int64(fi193)
                                                                                        v196 := GenS13{F0: 42, F1: 255}
                                                                                        switch (-(fi193)) {
                                                                                        case -9223372036854775808:
                                                                                            v1 += (fi192 - 1)
                                                                                            v197 := v2["\n\t"]
                                                                                            _ = v197
                                                                                        }
                                                                                        _ = v196
                                                                                    }
                                                                                    var v198 uint8 = 161
                                                                                    v199 := GenS21{F0: 2147483647, F1: 2147483647, F2: (v31 % ((v31) | 1))}
                                                                                    v200 := GenS10{F0: v198, F1: 2}
                                                                                    _ = v198
                                                                                    _ = v199
                                                                                    _ = v200
                                                                                }
                                                                            } else {
                                                                                v201 := GenS20{F0: (v31 | genFn1(-2147483648))}
                                                                                fmt.Println(v31, fi193, v31)
                                                                                _ = v201
                                                                            }
                                                                            v11 = (v11 - -1)
                                                                            fmt.Println(v11, fi192)
                                                                        }
                                                                        if true {
                                                                            _ = len(v25)
                                                                            var v202 uint = 822
                                                                            v203 := GenS9{F0: 42, F1: 1, F2: (v202 / (((^(v202))) | 1))}
                                                                            _ = v202
                                                                            _ = v203
                                                                        }
                                                                        var v204 uint32 = 1
                                                                        _ = v204
                                                                    }
                                                                    v205 := GenS6{F0: ""}
                                                                    var v206 uint64 = 18446744073709551615
                                                                    _ = v205
                                                                    _ = v206
                                                                } else {
                                                                    v207 := GenS6{F0: "\n\t"}
                                                                    _ = v207
                                                                }
                                                                v208 := GenS9{F0: 1, F1: 2, F2: 840}
                                                                _ = fi193
                                                                _ = v208
                                                            }
                                                            fmt.Println(v11, v11)
                                                            _ = fi192
                                                        }
                                                    }
                                                case "hello":
                                                    v209 := GenS11{F0: ""}
                                                    _ = v209
                                                }
                                                _ = v182
                                            } else {
                                                var v210 int8 = -1
                                                _ = v210
                                            }
                                        } else {
                                            _ = len(v2)
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                            v211 := GenS22{F0: false}
                                            _ = v211
                                        }
                                        switch (("foo bar" + "a") + (("a" + ("a" + "")) + "hello")) {
                                        case "":
                                            switch (v1 << ((202) & 63)) {
                                            case 127:
                                                if ((!true) || ((v11 - (v11 % (((v11 + v11)) | 1))) <= 330)) {
                                                    if len(v25) > 0 {
                                                        v25[0] = (v1 / ((genFn0((v31 & genFn1(-484)), 0)) | 1))
                                                    }
                                                    var v212 string = (("\n\t" + "") + "")
                                                    switch (v1 % ((v1) | 1)) {
                                                    case -1:
                                                        _ = len(v25)
                                                    case 1:
                                                        if false {
                                                            if len(v25) >= 2 {
                                                                v25 = v25[1:len(v25)]
                                                            }
                                                            if ((v1 / (((v1 - (-(v1)))) | 1)) == genFn0(v31, 0)) {
                                                                v11 -= v11
                                                                v213 := GenS16{F0: ((v212 + (("" + v212) + v212)) + ((v212 + "\n\t") + v212))}
                                                                var v214 uint32 = 110
                                                                _ = v213
                                                                _ = v214
                                                            }
                                                        }
                                                    case -267:
                                                        v11 += (v11 << ((65535) & 63))
                                                    default:
                                                        v31 += genFn1(-507)
                                                    }
                                                    _ = v212
                                                }
                                            }
                                        case "hello":
                                            for fi215 := 0; fi215 < 3; fi215++ {
                                                var v216 uint16 = 1
                                                var v217 float64 = 1.5
                                                _ = fi215
                                                _ = v216
                                                _ = v217
                                            }
                                        case "foo bar":
                                            _ = len(v3)
                                        }
                                        var v218 int64 = -9223372036854775808
                                        _ = v218
                                    }
                                    for fi219 := 0; fi219 < 20; fi219++ {
                                        v11 = (v11 - 0)
                                        delete(v3, ("\n\t" + ("" + ("" + ("" + (("hello" + "hello") + "a"))))))
                                        v1 += (fi219 * (v1 * (^(v1))))
                                        _ = fi219
                                    }
                                }
                            }
                        default:
                            var v220 bool = (4294967295 > v1)
                            _ = v220
                        }
                        v1 -= -1
                    }
                    var v221 uint64 = 73
                    _ = v221
                }
                switch 722 {
                case 377:
                    var v222 float32 = (1.5 / (0.0 - 3.14159))
                    _ = v222
                case 127:
                    v11 = -686
                case 32768:
                    v31 += (^(v31))
                }
                if false {
                    if (-866 >= -770) {
                        var v223 uint64 = 4294967295
                        v223 += v223
                        delete(v3, "\n\t")
                        _ = v223
                    } else {
                        for fi224 := 0; fi224 < 2; fi224++ {
                            if (true || false) {
                                v1 = (v1 & genFn0(genFn1(42), 67))
                                var v225 uint64 = 4294967295
                                _ = v225
                            } else {
                                for fi226 := 0; fi226 < 19; fi226++ {
                                    switch "a" {
                                    case "hello":
                                        switch (("" + ("hello" + "hello")) + "a") {
                                        case "hello":
                                            v3["a"] = 4294967295
                                            if false {
                                                if (!true) {
                                                    if true {
                                                        fmt.Println(v11, fi224, v11)
                                                        var v227 string = (((("foo bar" + "") + ((("\n\t" + "\n\t") + "a") + "\n\t")) + "") + "\n\t")
                                                        _ = v227
                                                    }
                                                }
                                            }
                                            var v228 uint32 = 42
                                            _ = v228
                                        case "foo bar":
                                            if true {
                                                var v229 float64 = (0.1 * 0.1)
                                                if (!false) {
                                                    if false {
                                                        switch fi226 {
                                                        case -1:
                                                            _ = len(v2)
                                                        case 128:
                                                            v229 = ((1e10 - (v229 + (v229 + 1.5))) * 3.14159)
                                                        case -477:
                                                            switch "" {
                                                            case "hello":
                                                                var v230 uint16 = 32767
                                                                _ = v230
                                                            }
                                                            if ((false || (65535 > 18446744073709551615)) || (37 <= 2)) {
                                                                v1 -= genFn0(914, 90)
                                                                if ((0.1 * (v229 + v229)) <= v229) {
                                                                    v11 = (^(v11))
                                                                    if ((v11 << ((42) & 63)) == 42) {
                                                                        v231 := GenS6{F0: ""}
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            var v232 int = (v1 ^ 128)
                                                                            _ = v232
                                                                        case "\n\t":
                                                                            v11 += 42
                                                                        case "a":
                                                                            v233 := GenS11{F0: "a"}
                                                                            _ = v233
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            fmt.Println(v11, v31, v11)
                                                                        case "":
                                                                            v11 -= (v11 ^ (v11 | (^(v11))))
                                                                        case "a":
                                                                            if (true || ((!false) && true)) {
                                                                                v234 := GenS22{F0: false}
                                                                                v235 := GenS23{F0: 1, F1: 2147483647}
                                                                                _ = v234
                                                                                _ = v235
                                                                            }
                                                                            switch genFn0(genFn1(1), 100) {
                                                                            case 0:
                                                                                v229 -= (1e10 + float64((float32(1e10) / (-2.25 / 1.5))))
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    v229 = v229
                                                                                default:
                                                                                    v236 := GenS9{F0: 2, F1: 1, F2: 42}
                                                                                    _ = v236
                                                                                }
                                                                                if (true || ((!("hello" == ("" + "hello"))) || (!(!(true && ("foo bar" <= "foo bar")))))) {
                                                                                    if false {
                                                                                        v237 := GenS24{F0: 18446744073709551615, F1: 847}
                                                                                        var v238 int16 = 1000
                                                                                        _ = v237
                                                                                        _ = v238
                                                                                    } else {
                                                                                        var v239 uint16 = 191
                                                                                        _ = v239
                                                                                    }
                                                                                    switch genFn0(v31, 127) {
                                                                                    case 2147483647:
                                                                                        if (976 <= 0) {
                                                                                            v31 += (v31 & v31)
                                                                                            var v240 bool = (4294967295 >= 2147483647)
                                                                                            v229 += ((v229 * ((float64(0.1) + v229) + 1e10)) / 1e10)
                                                                                            _ = v240
                                                                                        }
                                                                                        v241 := GenS21{F0: 42, F1: 2147483647, F2: v31}
                                                                                        _ = v241
                                                                                    case 0:
                                                                                        var v242 uint = 42
                                                                                        _ = v242
                                                                                    default:
                                                                                        if ("hello" < "\n\t") {
                                                                                            switch (fi224 % (((-(fi226))) | 1)) {
                                                                                            case 2147483647:
                                                                                                v243 := GenS7{F0: false, F1: v11}
                                                                                                _ = v243
                                                                                            case -1:
                                                                                                v244 := v2[""]
                                                                                                v245 := GenS14{F0: false, F1: 39, F2: v1}
                                                                                                _ = v244
                                                                                                _ = v245
                                                                                            case 128:
                                                                                                switch "hello" {
                                                                                                case "foo bar":
                                                                                                    v229 = v229
                                                                                                case "\n\t":
                                                                                                    switch (("a" + ("hello" + "a")) + ("" + "foo bar")) {
                                                                                                    case "a":
                                                                                                        fmt.Println(v31)
                                                                                                    default:
                                                                                                        var v246 int32 = 0
                                                                                                        _ = v246
                                                                                                    }
                                                                                                case "a":
                                                                                                    fmt.Println(v1, v1)
                                                                                                default:
                                                                                                    v247 := GenS16{F0: "\n\t"}
                                                                                                    _ = v247
                                                                                                }
                                                                                            default:
                                                                                                v11 -= -550
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v229, v11)
                                                                                    v248 := GenS3{F0: 127, F1: 335}
                                                                                    v249 := v2["hello"]
                                                                                    _ = v248
                                                                                    _ = v249
                                                                                }
                                                                            default:
                                                                                v3[("a" + "\n\t")] = 413
                                                                            }
                                                                            v250 := v3[(("" + ("a" + ("" + "foo bar"))) + "")]
                                                                            if len(v25) > 0 {
                                                                                v25[0] = genFn0(382, -123)
                                                                            }
                                                                            _ = v250
                                                                        default:
                                                                            var v251 bool = (!false)
                                                                            _ = v251
                                                                        }
                                                                        _ = v231
                                                                    }
                                                                    v25 = append(v25, fi224)
                                                                    var v252 uint64 = 96
                                                                    _ = v252
                                                                }
                                                                fmt.Println(v1)
                                                                var v253 float64 = float64(3.14159)
                                                                _ = v253
                                                            } else {
                                                                var v254 int64 = int64(v11)
                                                                if true {
                                                                    if ((268 <= 1) && ((2 > 42) && true)) {
                                                                        var v255 float64 = v229
                                                                        _ = v255
                                                                    } else {
                                                                        var v256 uint32 = 367
                                                                        fmt.Println(fi224, fi226, v11)
                                                                        var v257 bool = (4294967295 < genFn0((v31 / ((-919) | 1)), -1))
                                                                        _ = v256
                                                                        _ = v257
                                                                    }
                                                                    switch (v1 + fi226) {
                                                                    case 4294967295:
                                                                        var v258 uint = 187
                                                                        _ = v258
                                                                    }
                                                                } else {
                                                                    delete(v3, "foo bar")
                                                                    var v259 float64 = (((3.14159 - v229) + v229) - v229)
                                                                    var v260 int16 = (v11 << ((1) & 63))
                                                                    _ = v259
                                                                    _ = v260
                                                                }
                                                                v25 = append(v25, int(v11))
                                                                _ = v254
                                                            }
                                                            _ = len(v2)
                                                        default:
                                                            v11 += (v11 - int16(v31))
                                                        }
                                                    } else {
                                                        v1 = fi226
                                                        if ((65535 <= 65535) && true) {
                                                            var v261 int = v1
                                                            switch ("foo bar" + "hello") {
                                                            case "":
                                                                if (!(((^(v11)) <= (-(v11))) || true)) {
                                                                    v262 := GenS25{F0: 32767, F1: 545, F2: 65535}
                                                                    _ = v262
                                                                }
                                                            case "hello":
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    var v263 uint = 0
                                                                    _ = v263
                                                                case "a":
                                                                    v229 = (v229 + (1e10 * float64(1e10)))
                                                                }
                                                            case "a":
                                                                v264 := GenS17{F0: 457, F1: 0, F2: 0}
                                                                _ = v264
                                                            }
                                                            _ = v261
                                                        } else {
                                                            var v265 float32 = ((1.5 + 1.5) * (1e10 + 0.0))
                                                            _ = v265
                                                        }
                                                    }
                                                    if ((((0.1 * (1.5 * ((0.1 + ((3.14159 / ((1e10 / float32(v229)) - float32(v229))) * 0.0)) / 1.5))) - -2.25) - 1e10) <= 0.0) {
                                                        var v266 int16 = 1000
                                                        v267 := GenS26{F0: "\n\t", F1: 42, F2: 42}
                                                        _ = v266
                                                        _ = v267
                                                    } else {
                                                        v268 := GenS27{F0: 0, F1: "hello"}
                                                        if true {
                                                            v3["hello"] = 0
                                                            fmt.Println(v229)
                                                            var v269 float32 = -2.25
                                                            _ = v269
                                                        }
                                                        switch "a" {
                                                        case "foo bar":
                                                            v270 := GenS28{F0: "hello", F1: 0}
                                                            _ = v270
                                                        case "":
                                                            if true {
                                                                v271 := v3["hello"]
                                                                v11 -= -71
                                                                switch "a" {
                                                                case "\n\t":
                                                                    switch "hello" {
                                                                    case "":
                                                                        v272 := v2["a"]
                                                                        _ = v272
                                                                    case "foo bar":
                                                                        var v273 int32 = 42
                                                                        _ = v273
                                                                    default:
                                                                        if len(v25) >= 2 {
                                                                            v25 = v25[1:len(v25)]
                                                                        }
                                                                    }
                                                                case "a":
                                                                    var v274 uint16 = 255
                                                                    _ = v274
                                                                case "":
                                                                    if true {
                                                                        delete(v3, "foo bar")
                                                                    }
                                                                }
                                                                var v275 float64 = 3.14159
                                                                _ = v271
                                                                _ = v275
                                                            }
                                                        case "\n\t":
                                                            switch (fi224 + (fi224 / ((fi224) | 1))) {
                                                            case 32767:
                                                                v25 = append(v25, (^(fi224)))
                                                            case -839:
                                                                v276 := GenS5{F0: -54}
                                                                _ = v276
                                                            default:
                                                                fmt.Println(v31)
                                                            }
                                                        }
                                                        v1 = 4294967295
                                                        _ = len(v2)
                                                        _ = v268
                                                    }
                                                    fmt.Println(fi224)
                                                }
                                                _ = v229
                                            }
                                        }
                                    case "\n\t":
                                        var v277 int32 = 42
                                        _ = v277
                                    case "a":
                                        var v278 uint64 = 4294967295
                                        _ = v278
                                    }
                                    var v279 uint = 4294967295
                                    _ = fi226
                                    _ = v279
                                }
                                v280 := GenS29{F0: 145, F1: 42}
                                fmt.Println(v11, v1)
                                _ = v280
                            }
                            var v281 float32 = ((1e10 + 3.14159) + float32((float64((0.0 * ((1e10 + (0.0 / -2.25)) / (float32(-2.25) * 1e10)))) + 3.14159)))
                            v281 -= v281
                            _ = fi224
                            _ = v281
                        }
                        v282 := GenS30{F0: 893, F1: 123}
                        _ = v282
                    }
                    fmt.Println(v11)
                    for fi283 := 0; fi283 < 9; fi283++ {
                        switch "foo bar" {
                        case "a":
                            switch "\n\t" {
                            case "\n\t":
                                for fi284 := 0; fi284 < 14; fi284++ {
                                    if false {
                                        v285 := GenS31{F0: 815}
                                        v286 := GenS8{F0: 366}
                                        if false {
                                            v11 += v11
                                            v1 += 42
                                            var v287 uint32 = 2
                                            _ = v287
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                delete(v2, "a")
                                                v11 = -1
                                                var v288 uint8 = 156
                                                _ = v288
                                            default:
                                                v1 += -545
                                            }
                                            if true {
                                                v25 = append(v25, genFn0((v31 >> ((899) & 63)), 42))
                                                if (0 == 65535) {
                                                    v289 := v3["hello"]
                                                    v31 -= -9223372036854775808
                                                    _ = len(v25)
                                                    _ = v289
                                                } else {
                                                    v290 := GenS32{F0: 910, F1: 1, F2: 0}
                                                    _ = v290
                                                }
                                            }
                                            v291 := GenS21{F0: 0, F1: int32(fi283), F2: v31}
                                            v292 := GenS33{F0: 128}
                                            _ = v291
                                            _ = v292
                                        }
                                        v31 = v31
                                        _ = v285
                                        _ = v286
                                    } else {
                                        if true {
                                            switch v1 {
                                            case -409:
                                                v31 += -9223372036854775808
                                            default:
                                                v293 := GenS34{F0: "\n\t"}
                                                _ = v293
                                            }
                                            var v294 int16 = v11
                                            _ = v294
                                        } else {
                                            switch "\n\t" {
                                            case "a":
                                                v295 := GenS20{F0: -37}
                                                _ = v295
                                            case "hello":
                                                v1 = v1
                                                var v296 uint64 = 42
                                                _ = v296
                                            case "foo bar":
                                                if (true && false) {
                                                    switch fi284 {
                                                    case -577:
                                                        v1 = fi284
                                                    case 9223372036854775807:
                                                        switch int(v31) {
                                                        case -425:
                                                            _ = len(v25)
                                                        case -2147483648:
                                                            switch (v1 << ((255) & 63)) {
                                                            case -323:
                                                                v25 = append(v25, (fi284 << ((65535) & 63)))
                                                                var v297 float64 = 3.14159
                                                                _ = v297
                                                            case 539:
                                                                if (false && (true || false)) {
                                                                    var v298 int16 = (-(v11))
                                                                    _ = len(v25)
                                                                    _ = v298
                                                                } else {
                                                                    v299 := GenS8{F0: 122}
                                                                    _ = v299
                                                                }
                                                            default:
                                                                switch fi284 {
                                                                case 32768:
                                                                    var v300 int32 = -2147483648
                                                                    v25 = append(v25, fi284)
                                                                    var v301 uint = 2
                                                                    v302 := GenS33{F0: 0}
                                                                    _ = v300
                                                                    _ = v301
                                                                    _ = v302
                                                                }
                                                            }
                                                        case 194:
                                                            var v303 float64 = -2.25
                                                            _ = v303
                                                        default:
                                                            fmt.Println(v11)
                                                        }
                                                    default:
                                                        v11 -= int16(v31)
                                                    }
                                                }
                                            default:
                                                var v304 int32 = 1
                                                _ = v304
                                            }
                                            fmt.Println(v1, fi284)
                                        }
                                    }
                                    v25 = append(v25, -283)
                                    _ = fi284
                                }
                            case "":
                                v305 := GenS32{F0: 1, F1: 179, F2: 171}
                                _ = v305
                            case "foo bar":
                                for fi306 := 0; fi306 < 9; fi306++ {
                                    v31 -= 982
                                    var v307 uint16 = 547
                                    _ = fi306
                                    _ = v307
                                }
                            }
                        case "hello":
                            for fi308 := 0; fi308 < 5; fi308++ {
                                if false {
                                    if (497 <= 1) {
                                        var v309 int = (-(v1))
                                        v310 := GenS35{F0: 4294967295, F1: false}
                                        var v311 string = ((("a" + "") + "\n\t") + "")
                                        _ = len(v25)
                                        v1 -= fi283
                                        _ = v309
                                        _ = v310
                                        _ = v311
                                    } else {
                                        var v312 int64 = v31
                                        switch "a" {
                                        case "foo bar":
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                        default:
                                            v313 := GenS23{F0: 729, F1: -2147483648}
                                            _ = v313
                                        }
                                        _ = v312
                                    }
                                    v11 = 42
                                    var v314 uint = 869
                                    var v315 float64 = ((1e10 * (((0.0 / 0.1) + (0.1 * 3.14159)) / 0.1)) - 0.0)
                                    v31 += 1
                                    if true {
                                        v1 -= 471
                                        if (true || true) {
                                            fmt.Println(fi308, fi308)
                                            v316 := GenS30{F0: 481, F1: 128}
                                            v317 := GenS10{F0: 127, F1: 26}
                                            _ = v316
                                            _ = v317
                                        } else {
                                            var v318 uint16 = 255
                                            v319 := GenS28{F0: ("foo bar" + "hello"), F1: -2147483648}
                                            var v320 string = "hello"
                                            _ = v318
                                            _ = v319
                                            _ = v320
                                        }
                                    } else {
                                        v315 -= 1.5
                                        var v321 int32 = 2147483647
                                        v321 = (v321 | 0)
                                        _ = v321
                                    }
                                    _ = v314
                                    _ = v315
                                }
                                var v322 uint = 0
                                _ = fi308
                                _ = v322
                            }
                        default:
                            var v323 uint16 = 540
                            _ = v323
                        }
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                        fmt.Println(v11, fi283, fi283)
                        _ = fi283
                    }
                }
            } else {
                delete(v3, "hello")
                v1 = 13
            }
        }
    }
    v31 += (v31 ^ (v31 + (v31 ^ v31)))
    var v324 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v324:", v324)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v324
}
