package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 int8; F2 int16 }
type GenS6 struct { F0 uint64; F1 int32; F2 bool }
type GenS7 struct { F0 uint; F1 uint32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint8; F1 string }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint8; F1 bool; F2 int32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int8; F1 int8; F2 uint }
type GenS15 struct { F0 int16; F1 string }
type GenS16 struct { F0 int64 }
func genFn0() int8 {
    if (((!true) && true) && (!(((1.5 - (1e10 + -2.25)) + 1e10) < 3.14159))) {
        var v5 float64 = 1.5
        _ = v5
    } else {
        v6 := []int{-2, -9223372036854775808}
        _ = v6
    }
    var v7 uint = 18446744073709551615
    switch -9223372036854775808 {
    case 32768:
        v8 := GenS1{F0: v7}
        _ = v8
    case 1:
        var v9 uint = 1
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 int8 = -1
        _ = fi10
        _ = v11
    }
    switch "foo bar" {
    case "a":
        v12 := GenS2{F0: -2147483648}
        _ = v12
    }
    _ = v7
    return -1
}
func genFn1(p0 int, p1 int64) int {
    p0 -= p0
    var v13 bool = (!(0 >= 65535))
    _ = p0
    _ = p1
    _ = v13
    return p0
}

func main() {
    var v1 int32 = -2147483648
    var v2 int8 = 1
    var v3 uint64 = 42
    v4 := []int{}
    v2 -= v2
    v1 = (v1 ^ (^(v1)))
    if ((4294967295 != 65535) || true) {
        fmt.Println(v2)
        v14 := GenS1{F0: 65535}
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 5; fi15++ {
            if (!((!(int8(fi15) != 1)) && true)) {
                switch (fi15 << ((2) & 63)) {
                case 32768:
                    v3 = (v3 / (((v3 * v3)) | 1))
                }
            }
            v3 = 971
            if true {
                switch "\n\t" {
                case "foo bar":
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v17 := GenS1{F0: 42}
                        _ = fi16
                        _ = v17
                    }
                case "\n\t":
                    if false {
                        if false {
                            v3 -= 1
                        }
                        if false {
                            fmt.Println(v1, v1)
                        }
                        fmt.Println(fi15)
                    } else {
                        var v18 int8 = v2
                        _ = v18
                    }
                case "hello":
                    v1 += v1
                }
                switch "a" {
                case "":
                    v1 = v1
                case "a":
                    v3 = v3
                case "hello":
                    var v19 uint32 = 200
                    _ = v19
                }
            } else {
                fmt.Println(v3, v1)
                v20 := GenS3{F0: false, F1: (("" + (("" + (("\n\t" + ("hello" + "a")) + "\n\t")) + "a")) + "a")}
                var v21 int16 = 1
                _ = v20
                _ = v21
            }
            _ = fi15
        }
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    }
    switch -1 {
    case 682:
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS1{F0: 4294967295}
            var v24 uint64 = 465
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi22
            _ = v23
            _ = v24
        }
    case 127:
        switch "foo bar" {
        case "hello":
            if len(v4) > 0 {
                v4[0] = -66
            }
        }
    default:
        switch 42 {
        case 127:
            v25 := GenS4{F0: ""}
            _ = v25
        case -588:
            if false {
                v1 = v1
                if true {
                    switch "hello" {
                    case "foo bar":
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                    case "hello":
                        fmt.Println(v2, v1, v1)
                    case "\n\t":
                        if len(v4) > 0 {
                            v4[0] = 931
                        }
                    }
                    fmt.Println(v1)
                    if (!(1 == 819)) {
                        v1 += v1
                        fmt.Println(v3, v2)
                    }
                } else {
                    switch 9223372036854775807 {
                    case 113:
                        switch -1 {
                        case -526:
                            for fi26 := 0; fi26 < 8; fi26++ {
                                switch "\n\t" {
                                case "foo bar":
                                    v27 := GenS5{F0: 4294967295, F1: (v2 - (v2 - v2)), F2: 1}
                                    _ = v27
                                case "hello":
                                    for fi28 := 0; fi28 < 10; fi28++ {
                                        v4 = append(v4, fi26)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = fi28
                                    }
                                case "a":
                                    v29 := GenS1{F0: 4294967295}
                                    _ = v29
                                }
                                if true {
                                    if false {
                                        switch "foo bar" {
                                        case "a":
                                            v4 = append(v4, genFn1((fi26 | fi26), -9223372036854775808))
                                        case "\n\t":
                                            v2 += genFn0()
                                        default:
                                            for fi30 := 0; fi30 < 4; fi30++ {
                                                if false {
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi26 << ((42) & 63))
                                                    }
                                                    var v31 int8 = genFn0()
                                                    var v32 bool = false
                                                    _ = v31
                                                    _ = v32
                                                }
                                                _ = fi30
                                            }
                                        }
                                        var v33 uint8 = 0
                                        var v34 float64 = 0.0
                                        _ = v33
                                        _ = v34
                                    } else {
                                        if true {
                                            fmt.Println(v3)
                                            for fi35 := 0; fi35 < 10; fi35++ {
                                                v36 := GenS2{F0: genFn1(128, 0)}
                                                _ = fi35
                                                _ = v36
                                            }
                                        }
                                        var v37 uint32 = 2
                                        if false {
                                            v1 = v1
                                            fmt.Println(fi26)
                                            if true {
                                                for fi38 := 0; fi38 < 18; fi38++ {
                                                    var v39 int64 = 1
                                                    _ = fi38
                                                    _ = v39
                                                }
                                                for fi40 := 0; fi40 < 8; fi40++ {
                                                    if true {
                                                        v2 -= v2
                                                    }
                                                    v41 := GenS6{F0: (^(v3)), F1: (v1 * 2147483647), F2: (int16(fi26) > -834)}
                                                    _ = fi40
                                                    _ = v41
                                                }
                                            } else {
                                                if (fi26 < (fi26 * (^(fi26)))) {
                                                    switch fi26 {
                                                    case 499:
                                                        v42 := GenS3{F0: (!(4294967295 > (^(v3)))), F1: ""}
                                                        _ = v42
                                                    default:
                                                        if false {
                                                            v43 := GenS3{F0: false, F1: ("foo bar" + ((("a" + "foo bar") + "hello") + "foo bar"))}
                                                            var v44 uint = 4294967295
                                                            v45 := GenS3{F0: (0.1 != 1e10), F1: ((("foo bar" + ("a" + "hello")) + "foo bar") + "a")}
                                                            _ = v43
                                                            _ = v44
                                                            _ = v45
                                                        } else {
                                                            fmt.Println(fi26, v37)
                                                            if (false && (!false)) {
                                                                v1 = v1
                                                                var v46 uint64 = 658
                                                                fmt.Println(v37, v2)
                                                                _ = v46
                                                            } else {
                                                                for fi47 := 0; fi47 < 10; fi47++ {
                                                                    v48 := GenS7{F0: 65535, F1: v37}
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    _ = fi47
                                                                    _ = v48
                                                                }
                                                                if false {
                                                                    v3 -= v3
                                                                    v49 := GenS8{F0: false}
                                                                    var v50 float32 = 0.1
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                                if (42 <= int64(v37)) {
                                                                    var v51 int32 = -604
                                                                    _ = v51
                                                                } else {
                                                                    var v52 int16 = 0
                                                                    switch (fi26 >> ((341) & 63)) {
                                                                    case -9223372036854775808:
                                                                        if true {
                                                                            if true {
                                                                                for fi53 := 0; fi53 < 17; fi53++ {
                                                                                    switch (^(fi53)) {
                                                                                    case 128:
                                                                                        v3 -= 4294967295
                                                                                    case -2147483648:
                                                                                        _ = len(v4)
                                                                                        if false {
                                                                                            if (705 == v3) {
                                                                                                switch ("\n\t" + "hello") {
                                                                                                case "":
                                                                                                    var v54 float32 = 0.1
                                                                                                    _ = v54
                                                                                                }
                                                                                                var v55 bool = ((v2 / (((^(v2))) | 1)) < (^(v2)))
                                                                                                if false {
                                                                                                    fmt.Println(fi53, v55)
                                                                                                    if (631 >= 127) {
                                                                                                        v52 = v52
                                                                                                        var v56 uint16 = 65535
                                                                                                        _ = v56
                                                                                                    }
                                                                                                } else {
                                                                                                    if v55 {
                                                                                                        v57 := GenS8{F0: (((v2 == int8(v3)) || false) && false)}
                                                                                                        fmt.Println(v3, fi53)
                                                                                                        var v58 float32 = float32(0.0)
                                                                                                        _ = v57
                                                                                                        _ = v58
                                                                                                    }
                                                                                                }
                                                                                                v3 = (v3 | (v3 - 4294967295))
                                                                                                v59 := GenS5{F0: v37, F1: genFn0(), F2: v52}
                                                                                                _ = v55
                                                                                                _ = v59
                                                                                            } else {
                                                                                                switch ("hello" + "") {
                                                                                                case "hello":
                                                                                                    switch fi53 {
                                                                                                    case -633:
                                                                                                        var v60 int16 = v52
                                                                                                        _ = v60
                                                                                                    case -1:
                                                                                                        fmt.Println(v37, v52)
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    var v61 int64 = -9223372036854775808
                                                                                                    _ = v61
                                                                                                case "foo bar":
                                                                                                    var v62 int = fi53
                                                                                                    _ = v62
                                                                                                default:
                                                                                                    var v63 bool = (v1 < v1)
                                                                                                    _ = v63
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v52, fi53, v37)
                                                                                            var v64 bool = true
                                                                                            _ = v64
                                                                                        }
                                                                                        v1 -= v1
                                                                                    case 1:
                                                                                        var v65 string = ("hello" + "foo bar")
                                                                                        _ = v65
                                                                                    default:
                                                                                        v1 = 0
                                                                                    }
                                                                                    var v66 uint16 = 255
                                                                                    fmt.Println(v37, v3)
                                                                                    _ = fi53
                                                                                    _ = v66
                                                                                }
                                                                                if (42 <= 457) {
                                                                                    var v67 float64 = 0.1
                                                                                    switch ("foo bar" + (("hello" + "") + "foo bar")) {
                                                                                    case "\n\t":
                                                                                        var v68 uint16 = 65535
                                                                                        _ = v68
                                                                                    case "hello":
                                                                                        switch fi26 {
                                                                                        case -647:
                                                                                            var v69 uint = 4294967295
                                                                                            _ = v69
                                                                                        case 1:
                                                                                            if (334 == 2) {
                                                                                                v70 := GenS9{F0: 40}
                                                                                                var v71 bool = true
                                                                                                var v72 float32 = 0.0
                                                                                                _ = v70
                                                                                                _ = v71
                                                                                                _ = v72
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if true {
                                                                                        if true {
                                                                                            var v73 uint64 = 18446744073709551615
                                                                                            if false {
                                                                                                fmt.Println(fi26, v37)
                                                                                            } else {
                                                                                                var v74 float64 = 0.0
                                                                                                var v75 string = (("" + "\n\t") + "")
                                                                                                _ = v74
                                                                                                _ = v75
                                                                                            }
                                                                                            _ = v73
                                                                                        }
                                                                                        if (!(18446744073709551615 <= 255)) {
                                                                                            switch (fi26 >> ((0) & 63)) {
                                                                                            case -1:
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                            case 2147483647:
                                                                                                switch (fi26 | (fi26 & int(v1))) {
                                                                                                case 32767:
                                                                                                    v4 = append(v4, (fi26 - fi26))
                                                                                                case 42:
                                                                                                    v1 = -1
                                                                                                }
                                                                                            case 42:
                                                                                                var v76 int64 = 1
                                                                                                _ = v76
                                                                                            default:
                                                                                                if true {
                                                                                                    for fi77 := 0; fi77 < 14; fi77++ {
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                        _ = fi77
                                                                                                    }
                                                                                                    v78 := GenS2{F0: 9223372036854775807}
                                                                                                    _ = v78
                                                                                                } else {
                                                                                                    var v79 bool = (false || false)
                                                                                                    if ((v79 && (!v79)) && (-1 <= -9223372036854775808)) {
                                                                                                        fmt.Println(v37, v2)
                                                                                                    }
                                                                                                    fmt.Println(v79)
                                                                                                    _ = v79
                                                                                                }
                                                                                            }
                                                                                            for fi80 := 0; fi80 < 20; fi80++ {
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = fi26
                                                                                                }
                                                                                                _ = fi80
                                                                                            }
                                                                                            var v81 uint64 = 18446744073709551615
                                                                                            _ = v81
                                                                                        } else {
                                                                                            v82 := GenS4{F0: ("\n\t" + "\n\t")}
                                                                                            _ = v82
                                                                                        }
                                                                                        fmt.Println(v1)
                                                                                    } else {
                                                                                        v83 := GenS1{F0: 0}
                                                                                        if false {
                                                                                            var v84 int16 = (v52 % (((v52 & (v52 << ((255) & 63)))) | 1))
                                                                                            _ = v84
                                                                                        }
                                                                                        _ = v83
                                                                                    }
                                                                                    _ = v67
                                                                                }
                                                                                _ = len(v4)
                                                                            }
                                                                            var v85 uint = 2
                                                                            _ = v85
                                                                        } else {
                                                                            fmt.Println(v2)
                                                                            v1 = v1
                                                                            v86 := GenS3{F0: ((v52 << ((18446744073709551615) & 63)) < v52), F1: "\n\t"}
                                                                            _ = v86
                                                                        }
                                                                    case 0:
                                                                        var v87 uint64 = (v3 & (^(v3)))
                                                                        _ = v87
                                                                    case 2147483647:
                                                                        var v88 string = (("foo bar" + "hello") + ("hello" + ("" + (("a" + "hello") + "a"))))
                                                                        _ = v88
                                                                    default:
                                                                        fmt.Println(v3, v2, fi26)
                                                                    }
                                                                    _ = v52
                                                                }
                                                            }
                                                            for fi89 := 0; fi89 < 13; fi89++ {
                                                                var v90 uint = 18446744073709551615
                                                                var v91 float32 = 1e10
                                                                _ = fi89
                                                                _ = v90
                                                                _ = v91
                                                            }
                                                        }
                                                    }
                                                }
                                                v3 += 838
                                            }
                                        }
                                        _ = v37
                                    }
                                }
                                _ = fi26
                            }
                        case 127:
                            if true {
                                if ((v3 * (v3 << ((42) & 63))) > v3) {
                                    switch ((("\n\t" + "foo bar") + "") + "hello") {
                                    case "a":
                                        for fi92 := 0; fi92 < 17; fi92++ {
                                            v2 = (v2 / (((v2 - v2)) | 1))
                                            for fi93 := 0; fi93 < 15; fi93++ {
                                                _ = len(v4)
                                                var v94 uint32 = 0
                                                var v95 int8 = int8(fi92)
                                                _ = fi93
                                                _ = v94
                                                _ = v95
                                            }
                                            fmt.Println(v2, v2)
                                            _ = fi92
                                        }
                                    case "":
                                        switch ("" + "") {
                                        case "\n\t":
                                            if true {
                                                var v96 string = "\n\t"
                                                if len(v4) > 0 {
                                                    v4[0] = 4294967295
                                                }
                                                _ = v96
                                            }
                                        }
                                    case "\n\t":
                                        var v97 uint = 0
                                        _ = v97
                                    }
                                    _ = len(v4)
                                    if true {
                                        for fi98 := 0; fi98 < 2; fi98++ {
                                            if false {
                                                v99 := GenS10{F0: 2, F1: ("foo bar" + "")}
                                                switch (fi98 >> ((18446744073709551615) & 63)) {
                                                case 863:
                                                    for fi100 := 0; fi100 < 4; fi100++ {
                                                        v101 := GenS8{F0: true}
                                                        v102 := GenS10{F0: 128, F1: (("foo bar" + ("\n\t" + "foo bar")) + "hello")}
                                                        v3 -= v3
                                                        _ = fi100
                                                        _ = v101
                                                        _ = v102
                                                    }
                                                }
                                                _ = v99
                                            } else {
                                                var v103 uint32 = 338
                                                var v104 int16 = -74
                                                _ = v103
                                                _ = v104
                                            }
                                            v1 = (v1 ^ (v1 >> ((255) & 63)))
                                            _ = fi98
                                        }
                                    }
                                } else {
                                    var v105 string = ("" + ("foo bar" + ("foo bar" + "hello")))
                                    v3 = 0
                                    _ = v105
                                }
                            }
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    case 4294967295:
                        fmt.Println(v1, v2, v3)
                    }
                }
            } else {
                if false {
                    for fi106 := 0; fi106 < 8; fi106++ {
                        v107 := GenS11{F0: 65535}
                        _ = fi106
                        _ = v107
                    }
                    fmt.Println(v2, v1)
                } else {
                    if false {
                        if (true || true) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = len(v4)
                        } else {
                            switch 1 {
                            case 42:
                                fmt.Println(v2, v1, v2)
                            case 127:
                                if (!true) {
                                    for fi108 := 0; fi108 < 2; fi108++ {
                                        var v109 float64 = 1.5
                                        switch "hello" {
                                        case "hello":
                                            v110 := GenS4{F0: "foo bar"}
                                            _ = v110
                                        case "":
                                            var v111 float64 = 3.14159
                                            _ = v111
                                        case "foo bar":
                                            if len(v4) > 0 {
                                                v4[0] = genFn1(fi108, -291)
                                            }
                                        default:
                                            var v112 int8 = (^(v2))
                                            _ = v112
                                        }
                                        _ = fi108
                                        _ = v109
                                    }
                                } else {
                                    fmt.Println(v1)
                                    v3 = (v3 << ((42) & 63))
                                }
                            }
                            for fi113 := 0; fi113 < 7; fi113++ {
                                for fi114 := 0; fi114 < 17; fi114++ {
                                    v115 := GenS7{F0: 65535, F1: 65535}
                                    _ = fi114
                                    _ = v115
                                }
                                _ = fi113
                            }
                        }
                        v3 = v3
                    }
                    if true {
                        v2 -= v2
                        v116 := GenS10{F0: 128, F1: ("" + "")}
                        _ = v116
                    } else {
                        v4 = append(v4, 32768)
                        switch 32767 {
                        case 32767:
                            v2 = (v2 / (((v2 | int8(v3))) | 1))
                        case 127:
                            if false {
                                v1 = (^(v1))
                            } else {
                                if ((65535 == 27) || ((929 < 0) || true)) {
                                    var v117 string = "\n\t"
                                    _ = v117
                                } else {
                                    v1 = 367
                                }
                                if ((v3 < 4294967295) || (2 <= 2147483647)) {
                                    var v118 string = "hello"
                                    _ = v118
                                } else {
                                    for fi119 := 0; fi119 < 3; fi119++ {
                                        var v120 int8 = (v2 & (v2 + v2))
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = fi119
                                        _ = v120
                                    }
                                    v1 = (v1 + v1)
                                    v121 := GenS6{F0: (v3 & (v3 + (^(v3)))), F1: int32(v2), F2: true}
                                    _ = v121
                                }
                                if (-1 == 128) {
                                    v1 -= v1
                                } else {
                                    v1 = v1
                                    v122 := GenS3{F0: (!(true && true)), F1: "a"}
                                    v123 := GenS5{F0: 65535, F1: 100, F2: -60}
                                    _ = v122
                                    _ = v123
                                }
                            }
                        default:
                            v124 := GenS1{F0: 741}
                            _ = v124
                        }
                    }
                    switch (("" + "") + "hello") {
                    case "a":
                        v125 := GenS7{F0: 636, F1: 299}
                        _ = v125
                    case "foo bar":
                        var v126 uint32 = 4294967295
                        _ = v126
                    default:
                        v2 = 0
                    }
                }
            }
        case 9223372036854775807:
            if (true && (1 <= -648)) {
                if len(v4) > 0 {
                    v4[0] = 710
                }
                for fi127 := 0; fi127 < 13; fi127++ {
                    for fi128 := 0; fi128 < 12; fi128++ {
                        var v129 bool = true
                        var v130 bool = true
                        if len(v4) > 0 {
                            v4[0] = -838
                        }
                        switch "foo bar" {
                        case "a":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "\n\t":
                            v131 := GenS9{F0: 1}
                            v132 := GenS8{F0: (518 < -1)}
                            _ = v131
                            _ = v132
                        }
                        if len(v4) > 0 {
                            v4[0] = (fi127 / (((fi127 << ((1) & 63))) | 1))
                        }
                        if (0 != 1) {
                            v4 = append(v4, genFn1(fi128, 9223372036854775807))
                            fmt.Println(fi127, v2)
                            var v133 int8 = (v2 << ((18446744073709551615) & 63))
                            _ = v133
                        } else {
                            v134 := GenS8{F0: v129}
                            var v135 float32 = ((0.0 - 1.5) * 1.5)
                            _ = v134
                            _ = v135
                        }
                        _ = fi128
                        _ = v129
                        _ = v130
                    }
                    _ = fi127
                }
            } else {
                var v136 float32 = 1e10
                _ = v136
            }
        default:
            if len(v4) > 0 {
                v4[0] = 385
            }
        }
    }
    var v137 uint16 = 32767
    switch ("hello" + (("hello" + "foo bar") + "\n\t")) {
    case "":
        v138 := GenS4{F0: "a"}
        _ = v138
    case "hello":
        v2 -= 42
    case "a":
        for fi139 := 0; fi139 < 2; fi139++ {
            var v140 uint32 = 582
            v140 -= v140
            _ = fi139
            _ = v140
        }
    default:
        for fi141 := 0; fi141 < 16; fi141++ {
            var v142 uint8 = 0
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi143 := 0; fi143 < 20; fi143++ {
                v3 += 558
                _ = fi143
            }
            _ = fi141
            _ = v142
        }
    }
    for fi144 := 0; fi144 < 13; fi144++ {
        v145 := GenS6{F0: v3, F1: v1, F2: true}
        v4 = append(v4, fi144)
        var v146 float32 = 3.14159
        _ = fi144
        _ = v145
        _ = v146
    }
    switch "" {
    case "a":
        fmt.Println(v1, v2)
    case "foo bar":
        v147 := GenS2{F0: 1}
        _ = v147
    case "":
        var v148 uint16 = v137
        _ = v148
    default:
        var v149 uint16 = v137
        _ = v149
    }
    v150 := GenS12{F0: 42, F1: (!false), F2: v1}
    v151 := GenS13{F0: -385}
    for fi152 := 0; fi152 < 15; fi152++ {
        v153 := GenS14{F0: v2, F1: v2, F2: 374}
        v3 += (v3 - v3)
        v154 := GenS15{F0: -125, F1: (((("hello" + "foo bar") + (("hello" + ("hello" + ("" + "a"))) + "")) + ("a" + "")) + "\n\t")}
        _ = fi152
        _ = v153
        _ = v154
    }
    var v155 float32 = (0.1 / 0.1)
    if true {
        v155 -= v155
    } else {
        var v156 float32 = float32(float64((v155 - v155)))
        v157 := GenS16{F0: 42}
        _ = v156
        _ = v157
    }
    if ((float32(-2.25) + v155) > float32(-2.25)) {
        v3 = v3
        fmt.Println(v137, v1, v3)
    }
    switch 33 {
    case 453:
        v158 := GenS2{F0: 1}
        _ = v158
    case 0:
        var v159 float64 = 0.0
        _ = v159
    }
    v1 = (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("v155:", v155)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v137
    _ = v150
    _ = v151
    _ = v155
    _ = v2
    _ = v3
    _ = v4
}
