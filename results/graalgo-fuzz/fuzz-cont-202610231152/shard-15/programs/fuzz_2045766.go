package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 uint8; F2 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 int; F2 int32 }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 string; F1 uint; F2 uint }
type GenS8 struct { F0 int16; F1 string; F2 uint16 }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 string }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 bool; F1 bool; F2 int64 }
type GenS13 struct { F0 uint32; F1 int16 }
type GenS14 struct { F0 uint64; F1 uint16 }
type GenS15 struct { F0 uint; F1 int16 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint; F1 int16 }
type GenS18 struct { F0 int8; F1 uint64 }
type GenS19 struct { F0 uint32; F1 uint32 }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 uint64; F1 int; F2 int16 }
type GenS22 struct { F0 int32; F1 uint }
type GenS23 struct { F0 int8; F1 bool }
type GenS24 struct { F0 uint32; F1 int }

func main() {
    v1 := []string{((((("foo bar" + "hello") + "a") + "\n\t") + "\n\t") + "a"), "\n\t", "hello"}
    var v2 string = "hello"
    v3 := make(map[string]uint32)
    switch -1 {
    case -2147483648:
        v2 = (((("\n\t" + "\n\t") + v2) + v2) + v2)
    }
    if (255 <= 403) {
        v4 := v3[v2]
        _ = v4
    }
    v5 := GenS1{F0: 32767}
    for fi6 := 0; fi6 < 20; fi6++ {
        v2 = (("" + v2) + v2)
        var v7 int8 = 42
        v8 := GenS2{F0: 2147483647}
        _ = fi6
        _ = v7
        _ = v8
    }
    switch (v2 + v2) {
    case "foo bar":
        if (!(!false)) {
            v9 := GenS3{F0: -1, F1: 255, F2: 4294967295}
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            switch (((v2 + "") + v2) + (v2 + "foo bar")) {
            case "a":
                v2 = v2
            case "\n\t":
                var v11 uint = 18446744073709551615
                _ = v11
            case "foo bar":
                v2 = "a"
            default:
                v12 := GenS4{F0: -309}
                _ = v12
            }
            v13 := GenS4{F0: 0}
            _ = fi10
            _ = v13
        }
    }
    if true {
        if (67 < 42) {
            if (0.0 >= (1e10 / 1.5)) {
                if (false && false) {
                    switch v2 {
                    case "\n\t":
                        for fi14 := 0; fi14 < 5; fi14++ {
                            for fi15 := 0; fi15 < 12; fi15++ {
                                var v16 bool = false
                                _ = fi15
                                _ = v16
                            }
                            _ = fi14
                        }
                    }
                    if true {
                        var v17 uint64 = 4294967295
                        v18 := GenS4{F0: 0}
                        v2 = v2
                        _ = v17
                        _ = v18
                    }
                } else {
                    if (!true) {
                        var v19 uint16 = 1
                        for fi20 := 0; fi20 < 8; fi20++ {
                            v19 = v19
                            if false {
                                fmt.Println(v19)
                                var v21 uint8 = 42
                                _ = v21
                            } else {
                                v19 = v19
                                for fi22 := 0; fi22 < 7; fi22++ {
                                    v19 = 65535
                                    _ = fi22
                                }
                                var v23 int64 = -9223372036854775808
                                _ = v23
                            }
                            _ = len(v1)
                            _ = fi20
                        }
                        _ = v19
                    }
                    switch v2 {
                    case "foo bar":
                        switch v2 {
                        case "hello":
                            if false {
                                for fi24 := 0; fi24 < 14; fi24++ {
                                    v25 := GenS5{F0: -32768, F1: (fi24 << ((65535) & 63)), F2: 505}
                                    _ = fi24
                                    _ = v25
                                }
                                switch v2 {
                                case "foo bar":
                                    var v26 int16 = 1
                                    _ = v26
                                case "":
                                    v2 = ""
                                case "hello":
                                    var v27 int64 = 1
                                    _ = v27
                                default:
                                    delete(v3, v2)
                                }
                                delete(v3, v2)
                            } else {
                                var v28 uint16 = 679
                                v29 := GenS6{F0: 32767, F1: 355}
                                _ = v28
                                _ = v29
                            }
                        case "\n\t":
                            var v30 uint16 = 65535
                            _ = v30
                        case "a":
                            v2 = v2
                        }
                    case "hello":
                        if ((!(true && (true || false))) || (false || (!false))) {
                            var v31 uint8 = 1
                            _ = v31
                        }
                    }
                    for fi32 := 0; fi32 < 17; fi32++ {
                        if (-9223372036854775808 == 43) {
                            var v33 uint32 = 0
                            for fi34 := 0; fi34 < 15; fi34++ {
                                if true {
                                    _ = len(v1)
                                }
                                fmt.Println(v2)
                                _ = fi34
                            }
                            if true {
                                v2 = v2
                                switch fi32 {
                                case -927:
                                    switch 371 {
                                    case 0:
                                        for fi35 := 0; fi35 < 13; fi35++ {
                                            v36 := GenS2{F0: -2147483648}
                                            if true {
                                                v37 := GenS7{F0: v2, F1: 1, F2: 18446744073709551615}
                                                v33 += (v33 ^ 2147483647)
                                                _ = v37
                                            }
                                            _ = fi35
                                            _ = v36
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v2, v33)
                                    default:
                                        for fi38 := 0; fi38 < 5; fi38++ {
                                            if false {
                                                delete(v3, "a")
                                                var v39 uint32 = v33
                                                switch v2 {
                                                case "foo bar":
                                                    if ((0.1 / -2.25) >= ((float32((0.0 - 0.0)) / 1.0) * 1e10)) {
                                                        v39 = 0
                                                        if false {
                                                            v33 += v33
                                                            var v40 int64 = -1
                                                            if false {
                                                                switch (fi38 << ((18446744073709551615) & 63)) {
                                                                case 1:
                                                                    v41 := GenS8{F0: -1, F1: ("hello" + v2), F2: 657}
                                                                    _ = v41
                                                                }
                                                                if true {
                                                                    switch (^(fi38)) {
                                                                    case 0:
                                                                        var v42 uint16 = 65535
                                                                        _ = v42
                                                                    case 650:
                                                                        var v43 uint64 = 0
                                                                        _ = v43
                                                                    }
                                                                    v1 = append(v1, "")
                                                                } else {
                                                                    fmt.Println(v40, v2, fi32)
                                                                }
                                                            } else {
                                                                var v44 string = "a"
                                                                var v45 string = "\n\t"
                                                                v40 += 9223372036854775807
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                            v46 := GenS3{F0: 32767, F1: 128, F2: fi32}
                                                            _ = v40
                                                            _ = v46
                                                        }
                                                    } else {
                                                        v39 += v33
                                                        v47 := GenS7{F0: ("\n\t" + ((("foo bar" + ((v2 + ("hello" + v2)) + v2)) + (v2 + "")) + "\n\t")), F1: 42, F2: 534}
                                                        if false {
                                                            fmt.Println(v33, v2)
                                                            v48 := GenS7{F0: v2, F1: 65535, F2: 42}
                                                            v39 += (v39 + v39)
                                                            _ = v48
                                                        } else {
                                                            if (true || (!false)) {
                                                                var v49 int32 = int32(v33)
                                                                _ = v49
                                                            }
                                                        }
                                                        _ = v47
                                                    }
                                                case "\n\t":
                                                    switch (fi38 * fi38) {
                                                    case 32768:
                                                        delete(v3, "a")
                                                    }
                                                case "a":
                                                    v3[v2] = v33
                                                }
                                                v33 += (v33 / ((v33) | 1))
                                                _ = v39
                                            } else {
                                                var v50 uint16 = 65535
                                                switch fi38 {
                                                case 102:
                                                    var v51 float32 = 3.14159
                                                    _ = v51
                                                default:
                                                    fmt.Println(fi38, v50, v2)
                                                    fmt.Println(v2, v50, fi32)
                                                    if false {
                                                        fmt.Println(fi38, v50)
                                                        v50 = v50
                                                        v33 -= v33
                                                    }
                                                    delete(v3, (v2 + v2))
                                                    if (-832 < -784) {
                                                        v52 := GenS9{F0: 1000, F1: true}
                                                        v2 = "foo bar"
                                                        if (14 == 210) {
                                                            v33 = (v33 ^ 42)
                                                            switch "a" {
                                                            case "foo bar":
                                                                if true {
                                                                    switch (fi38 - 32767) {
                                                                    case 1:
                                                                        var v53 uint32 = (v33 - (v33 << ((uint(v50)) & 63)))
                                                                        _ = v53
                                                                    default:
                                                                        v54 := GenS10{F0: v2}
                                                                        _ = v54
                                                                    }
                                                                    fmt.Println(v50, v50, v33)
                                                                    if true {
                                                                        if ((9223372036854775807 == 540) || ((!(2 != 159)) && false)) {
                                                                            if (93 < -32768) {
                                                                                var v55 int8 = -3
                                                                                _ = len(v1)
                                                                                switch (-(fi38)) {
                                                                                case 1:
                                                                                    v33 += v33
                                                                                default:
                                                                                    if (!(-218 == -402)) {
                                                                                        fmt.Println(v55)
                                                                                        var v56 int64 = 42
                                                                                        v2 = ""
                                                                                        _ = v56
                                                                                    } else {
                                                                                        var v57 int32 = 1
                                                                                        var v58 float64 = 1e10
                                                                                        v59 := GenS3{F0: 669, F1: 42, F2: fi32}
                                                                                        _ = v57
                                                                                        _ = v58
                                                                                        _ = v59
                                                                                    }
                                                                                }
                                                                                _ = v55
                                                                            }
                                                                            var v60 string = v2
                                                                            switch "" {
                                                                            case "":
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    fmt.Println(fi38)
                                                                                case "a":
                                                                                    var v61 float32 = 1e10
                                                                                    _ = v61
                                                                                case "foo bar":
                                                                                    v62 := GenS11{F0: (0 > 126)}
                                                                                    _ = v62
                                                                                }
                                                                            case "foo bar":
                                                                                fmt.Println(v2)
                                                                            case "\n\t":
                                                                                if (41 > (v33 % ((v33) | 1))) {
                                                                                    v60 = ((("foo bar" + v2) + (v2 + ("" + "\n\t"))) + "foo bar")
                                                                                    var v63 float64 = ((float64((3.14159 / 1.0)) + (1.5 * -2.25)) - float64((0.0 / (0.1 + ((1.5 - ((((-2.25 * 0.1) / (float32(0.1) * float32((((0.1 - ((0.0 + 1.5) + float64((float32(float64(0.0)) + (0.1 + 1e10))))) - 0.1) / 1.0)))) - (-2.25 - 0.0)) / 0.1)) / 0.1)))))
                                                                                    switch (fi32 ^ fi32) {
                                                                                    case -760:
                                                                                        if true {
                                                                                            fmt.Println(fi32, v33, v60)
                                                                                            switch (fi38 | fi32) {
                                                                                            case -712:
                                                                                                v3[v60] = v33
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v63
                                                                                } else {
                                                                                    switch v60 {
                                                                                    case "a":
                                                                                        switch 128 {
                                                                                        case 32767:
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                var v64 uint16 = 0
                                                                                                _ = v64
                                                                                            default:
                                                                                                var v65 uint32 = 688
                                                                                                _ = v65
                                                                                            }
                                                                                        case -9223372036854775808:
                                                                                            fmt.Println(fi38, v33)
                                                                                        default:
                                                                                            if true {
                                                                                                switch v2 {
                                                                                                case "foo bar":
                                                                                                    _ = len(v1)
                                                                                                default:
                                                                                                    v66 := GenS12{F0: true, F1: ((-7 > 1) && (-552 > 167)), F2: 763}
                                                                                                    _ = v66
                                                                                                }
                                                                                                if false {
                                                                                                    if (false && ((-2.25 > float64(0.0)) && (false && true))) {
                                                                                                        delete(v3, "\n\t")
                                                                                                        var v67 float64 = 1e10
                                                                                                        _ = v67
                                                                                                    } else {
                                                                                                        var v68 uint64 = 18446744073709551615
                                                                                                        if false {
                                                                                                            var v69 uint64 = (^(v68))
                                                                                                            _ = v69
                                                                                                        }
                                                                                                        v33 += 2147483647
                                                                                                        var v70 int = fi38
                                                                                                        _ = v68
                                                                                                        _ = v70
                                                                                                    }
                                                                                                }
                                                                                                var v71 int16 = -1
                                                                                                _ = v71
                                                                                            } else {
                                                                                                if (!false) {
                                                                                                    if false {
                                                                                                        v2 = ((v2 + (v2 + ((v2 + "a") + v2))) + v60)
                                                                                                        switch v2 {
                                                                                                        case "hello":
                                                                                                            v72 := v3[v2]
                                                                                                            _ = v72
                                                                                                        }
                                                                                                        v33 = (v33 | v33)
                                                                                                        if (1 < 53) {
                                                                                                            delete(v3, "foo bar")
                                                                                                        } else {
                                                                                                            fmt.Println(v33)
                                                                                                        }
                                                                                                        fmt.Println(fi32)
                                                                                                        var v73 bool = false
                                                                                                        _ = v73
                                                                                                    } else {
                                                                                                        fmt.Println(fi38, fi32)
                                                                                                        switch v2 {
                                                                                                        case "foo bar":
                                                                                                            v33 = (v33 & 635)
                                                                                                            v50 = (v50 + v50)
                                                                                                            v74 := GenS13{F0: (v33 * 1), F1: -1}
                                                                                                            _ = v74
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = len(v1)
                                                                                            }
                                                                                        }
                                                                                    case "":
                                                                                        v3[(((((v2 + "a") + (v2 + v2)) + ("" + (v60 + v60))) + v2) + (("\n\t" + (v60 + "foo bar")) + "hello"))] = v33
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v60
                                                                        }
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                var v75 float32 = -2.25
                                                                var v76 int64 = -187
                                                                _ = v75
                                                                _ = v76
                                                            case "hello":
                                                                v77 := GenS2{F0: 1}
                                                                _ = v77
                                                            }
                                                        }
                                                        var v78 int = (fi32 / ((fi38) | 1))
                                                        if ((1e10 != 1.5) && true) {
                                                            v3[(v2 + v2)] = (^(v33))
                                                            _ = len(v1)
                                                            fmt.Println(v33, fi38, fi38)
                                                            var v79 string = "foo bar"
                                                            fmt.Println(v2, v2, v50)
                                                            _ = v79
                                                        }
                                                        _ = v52
                                                        _ = v78
                                                    }
                                                    var v80 uint16 = 56
                                                    if (((fi38 == 1) || false) && (241 > 128)) {
                                                        if (!(!false)) {
                                                            v81 := GenS14{F0: 0, F1: (^(v50))}
                                                            _ = v81
                                                        } else {
                                                            if ((^(v80)) > (^(v80))) {
                                                                switch (fi32 / ((-335) | 1)) {
                                                                case -1:
                                                                    v2 = "hello"
                                                                    v50 += (v80 * v50)
                                                                    delete(v3, "a")
                                                                    v1 = append(v1, v2)
                                                                    delete(v3, "hello")
                                                                    v80 = 42
                                                                    if (("foo bar" + v2) > v2) {
                                                                        delete(v3, v2)
                                                                        var v82 int64 = 42
                                                                        var v83 bool = (false || true)
                                                                        if (!false) {
                                                                            v82 = -9223372036854775808
                                                                            if len(v1) > 0 {
                                                                                v1[0] = v2
                                                                            }
                                                                            var v84 string = ("a" + "foo bar")
                                                                            _ = v84
                                                                        } else {
                                                                            v3[""] = 595
                                                                            v85 := GenS15{F0: 0, F1: -608}
                                                                            _ = v85
                                                                        }
                                                                        delete(v3, v2)
                                                                        _ = v82
                                                                        _ = v83
                                                                    }
                                                                }
                                                                _ = len(v3)
                                                                _ = len(v3)
                                                            } else {
                                                                v86 := GenS13{F0: (v33 | (v33 ^ 4294967295)), F1: 1}
                                                                v2 = "hello"
                                                                v87 := GenS16{F0: true}
                                                                switch fi32 {
                                                                case 1:
                                                                    fmt.Println(v50, fi38, v50)
                                                                    v33 = (v33 + v33)
                                                                    if false {
                                                                        var v88 float32 = 0.1
                                                                        if true {
                                                                            switch (fi38 >> ((758) & 63)) {
                                                                            case -1:
                                                                                var v89 uint = 2
                                                                                _ = v89
                                                                            case 1:
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            case -2147483648:
                                                                                var v90 int8 = -1
                                                                                _ = v90
                                                                            }
                                                                            v88 -= 1.5
                                                                        }
                                                                        _ = v88
                                                                    } else {
                                                                        v91 := GenS17{F0: 0, F1: -22}
                                                                        var v92 float32 = (3.14159 / 1.0)
                                                                        v93 := GenS8{F0: -678, F1: (((("hello" + v2) + "\n\t") + v2) + (v2 + (v2 + ("hello" + v2)))), F2: (v80 & v50)}
                                                                        _ = v91
                                                                        _ = v92
                                                                        _ = v93
                                                                    }
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                                var v94 int32 = 2147483647
                                                                _ = v86
                                                                _ = v87
                                                                _ = v94
                                                            }
                                                        }
                                                    } else {
                                                        if true {
                                                            v2 = (v2 + v2)
                                                            fmt.Println(v50, fi38)
                                                            v95 := GenS18{F0: 107, F1: 18446744073709551615}
                                                            _ = v95
                                                        } else {
                                                            v96 := GenS16{F0: (-533 >= 9223372036854775807)}
                                                            v33 = (v33 << ((469) & 63))
                                                            v2 = v2
                                                            _ = v96
                                                        }
                                                        v2 = (v2 + (("hello" + v2) + v2))
                                                        switch ("hello" + v2) {
                                                        case "hello":
                                                            if true {
                                                                switch ((v2 + v2) + v2) {
                                                                case "":
                                                                    var v97 bool = (-1 < 1)
                                                                    _ = v97
                                                                case "\n\t":
                                                                    var v98 int = (fi38 << ((541) & 63))
                                                                    fmt.Println(v98, v50)
                                                                    var v99 bool = true
                                                                    _ = v98
                                                                    _ = v99
                                                                case "foo bar":
                                                                    fmt.Println(v80, fi38)
                                                                default:
                                                                    v50 -= (v80 % (((v80 + (v80 % ((v50) | 1)))) | 1))
                                                                }
                                                                v1 = append(v1, (("hello" + "\n\t") + "hello"))
                                                                v100 := GenS12{F0: (1 == 244), F1: false, F2: 1000}
                                                                _ = v100
                                                            } else {
                                                                if true {
                                                                    v101 := v3[((v2 + v2) + v2)]
                                                                    if len(v1) > 0 {
                                                                        v1[0] = v2
                                                                    }
                                                                    if ((42 < 597) && true) {
                                                                        fmt.Println(v2)
                                                                        v2 = "\n\t"
                                                                        v102 := GenS19{F0: v101, F1: v33}
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = v102
                                                                    } else {
                                                                        if (!(1 <= 0)) {
                                                                            v33 = (v101 + 2147483648)
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = v2
                                                                        }
                                                                    }
                                                                    _ = v101
                                                                } else {
                                                                    if true {
                                                                        if false {
                                                                            v33 = (v33 | 42)
                                                                            var v103 uint = 65535
                                                                            v33 = v33
                                                                            _ = v103
                                                                        } else {
                                                                            if ((!((3.14159 / 1.0) > ((0.1 * 1e10) + ((float64(1.5) - 3.14159) + ((3.14159 / 1.0) + (0.0 - -2.25)))))) || true) {
                                                                                switch (fi38 * (fi38 | (fi32 / (((-(fi38))) | 1)))) {
                                                                                case 9223372036854775807:
                                                                                    var v104 uint16 = 1
                                                                                    _ = v104
                                                                                case 32767:
                                                                                    v50 = v50
                                                                                }
                                                                                var v105 int64 = -9223372036854775808
                                                                                v106 := GenS7{F0: ((v2 + "a") + (v2 + ("a" + v2))), F1: 65535, F2: 0}
                                                                                _ = v105
                                                                                _ = v106
                                                                            }
                                                                            fmt.Println(fi38, v80, v2)
                                                                        }
                                                                    }
                                                                    _ = len(v1)
                                                                    var v107 uint8 = 113
                                                                    _ = v107
                                                                }
                                                                v108 := GenS2{F0: -558}
                                                                _ = v108
                                                            }
                                                        case "\n\t":
                                                            v109 := GenS8{F0: -40, F1: (v2 + v2), F2: (v50 / (((v80 & 2)) | 1))}
                                                            _ = v109
                                                        default:
                                                            v50 = 0
                                                        }
                                                        var v110 uint = 255
                                                        v50 = v80
                                                        switch v2 {
                                                        case "a":
                                                            v33 += v33
                                                        }
                                                        v50 = 255
                                                        _ = v110
                                                    }
                                                    _ = v80
                                                }
                                                _ = v50
                                            }
                                            var v111 string = (v2 + "\n\t")
                                            _ = fi38
                                            _ = v111
                                        }
                                    }
                                case 1:
                                    v33 -= v33
                                case 925:
                                    v112 := GenS17{F0: 160, F1: -513}
                                    _ = v112
                                default:
                                    var v113 uint = 961
                                    _ = v113
                                }
                                var v114 int8 = 42
                                _ = v114
                            } else {
                                if false {
                                    v2 = v2
                                }
                                v115 := GenS1{F0: 0}
                                for fi116 := 0; fi116 < 11; fi116++ {
                                    var v117 int = (fi116 + fi32)
                                    var v118 uint = 522
                                    v33 += v33
                                    _ = fi116
                                    _ = v117
                                    _ = v118
                                }
                                _ = v115
                            }
                            _ = v33
                        }
                        var v119 uint64 = 4294967295
                        _ = fi32
                        _ = v119
                    }
                }
                var v120 int8 = 127
                var v121 bool = ((698 <= 4294967295) || true)
                _ = v120
                _ = v121
            }
            v122 := GenS20{F0: 71}
            var v123 float32 = 0.0
            _ = v122
            _ = v123
        } else {
            v2 = "\n\t"
            for fi124 := 0; fi124 < 17; fi124++ {
                v125 := GenS21{F0: 0, F1: (fi124 >> ((1) & 63)), F2: 245}
                for fi126 := 0; fi126 < 11; fi126++ {
                    if (false && (0.0 < float32(0.1))) {
                        var v127 float64 = (0.0 * 1.5)
                        switch v2 {
                        case "a":
                            if (!false) {
                                if (!(((42 != 42) && true) && false)) {
                                    _ = len(v1)
                                }
                                v127 = v127
                            } else {
                                var v128 bool = (!false)
                                v127 = (0.0 - float64(((1.5 / -2.25) * (1e10 - 1e10))))
                                _ = v128
                            }
                        case "foo bar":
                            v129 := GenS16{F0: ((!(42 <= 32767)) && true)}
                            _ = v129
                        }
                        switch "hello" {
                        case "\n\t":
                            var v130 int = (fi124 >> ((255) & 63))
                            _ = v130
                        case "":
                            var v131 int8 = -128
                            _ = v131
                        case "foo bar":
                            var v132 int32 = 42
                            _ = v132
                        }
                        if false {
                            v127 = v127
                            fmt.Println(v2, fi124)
                            if true {
                                v2 = v2
                            } else {
                                v127 += 1.5
                                v133 := GenS12{F0: (!(false || (0 != 1))), F1: ((1e10 + 0.1) <= 0.0), F2: 540}
                                v127 -= 0.0
                                _ = v133
                            }
                        }
                        _ = v127
                    }
                    var v134 int8 = -79
                    _ = fi126
                    _ = v134
                }
                _ = fi124
                _ = v125
            }
            var v135 uint16 = 65535
            _ = v135
        }
        v2 = ("\n\t" + v2)
        if true {
            var v136 string = ("a" + v2)
            if (!(!false)) {
                var v137 int64 = -9223372036854775808
                fmt.Println(v136)
                _ = v137
            }
            _ = v136
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v138 uint = 1
            _ = v138
        }
    } else {
        if (!true) {
            v139 := GenS22{F0: -2147483648, F1: 18446744073709551615}
            v2 = v2
            v140 := GenS2{F0: -35}
            _ = v139
            _ = v140
        } else {
            v2 = v2
            if true {
                for fi141 := 0; fi141 < 1; fi141++ {
                    _ = len(v3)
                    _ = fi141
                }
            } else {
                var v142 int = 32767
                v2 = "a"
                v143 := GenS19{F0: 889, F1: 2}
                _ = v142
                _ = v143
            }
            if (!true) {
                v144 := GenS23{F0: 42, F1: (0 == 65535)}
                var v145 float32 = (1e10 - (1e10 + (1e10 - -2.25)))
                v146 := GenS24{F0: 2, F1: 127}
                _ = v144
                _ = v145
                _ = v146
            }
        }
        var v147 int32 = 0
        v147 = (v147 % (((-(v147))) | 1))
        _ = v147
    }
    v2 = ""
    _ = len(v1)
    v2 = v2
    if true {
        if (4294967295 <= -2147483648) {
            v2 = (v2 + v2)
            if false {
                for fi148 := 0; fi148 < 16; fi148++ {
                    for fi149 := 0; fi149 < 3; fi149++ {
                        var v150 int8 = -33
                        fmt.Println(fi149, fi148, v2)
                        _ = fi149
                        _ = v150
                    }
                    var v151 uint32 = 2147483648
                    _ = fi148
                    _ = v151
                }
                if (false || true) {
                    var v152 uint8 = 89
                    _ = v152
                }
                var v153 uint32 = 782
                _ = v153
            } else {
                v2 = ((v2 + v2) + "foo bar")
                var v154 int = 9223372036854775807
                for fi155 := 0; fi155 < 19; fi155++ {
                    for fi156 := 0; fi156 < 6; fi156++ {
                        switch (fi156 % (((fi155 / ((4294967295) | 1))) | 1)) {
                        case -1:
                            var v157 int16 = 32767
                            _ = v157
                        case 128:
                            var v158 int64 = 454
                            _ = v158
                        default:
                            fmt.Println(v2, fi156, fi156)
                        }
                        _ = fi156
                    }
                    _ = len(v1)
                    if true {
                        for fi159 := 0; fi159 < 5; fi159++ {
                            fmt.Println(fi159)
                            var v160 bool = (255 < 42)
                            v160 = v160
                            _ = fi159
                            _ = v160
                        }
                        var v161 int16 = 1
                        fmt.Println(v2, v154, v2)
                        _ = v161
                    }
                    _ = fi155
                }
                _ = v154
            }
            var v162 uint = 18446744073709551615
            _ = v162
        } else {
            var v163 float32 = ((0.0 - (-2.25 / 3.14159)) - (((1e10 + 1e10) / -2.25) - 0.0))
            v163 -= 0.0
            for fi164 := 0; fi164 < 14; fi164++ {
                v2 = v2
                _ = fi164
            }
            _ = v163
        }
    }
    if (79 <= 128) {
        v2 = ("a" + ("hello" + "hello"))
        switch "a" {
        case "a":
            var v165 int32 = 42
            _ = v165
        case "":
            _ = len(v1)
        }
        v2 = v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v5
}
