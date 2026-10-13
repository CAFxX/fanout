package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 bool; F1 uint32; F2 bool }
type GenS7 struct { F0 int8; F1 string; F2 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string; F1 string }
type GenS10 struct { F0 uint16; F1 int }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int32; F1 int }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int32; F1 int16 }
type GenS16 struct { F0 string; F1 uint32; F2 uint8 }
type GenS17 struct { F0 uint64; F1 int; F2 uint32 }
type GenS18 struct { F0 uint; F1 int32; F2 uint16 }
type GenS19 struct { F0 int64; F1 string; F2 string }
type GenS20 struct { F0 uint }
type GenS21 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 int16 }
type GenS24 struct { F0 int; F1 int8; F2 uint32 }
type GenS25 struct { F0 uint }
type GenS26 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS27 struct { F0 uint8 }
type GenS28 struct { F0 int }
type GenS29 struct { F0 uint; F1 int16; F2 string }
type GenS30 struct { F0 string; F1 bool }
type GenS31 struct { F0 uint16; F1 string; F2 int }
type GenS32 struct { F0 string }
func genFn0() string {
    var v5 float32 = -2.25
    v5 -= float32(1e10)
    v5 = 0.0
    var v6 string = "hello"
    fmt.Println(v6, v5)
    _ = v5
    _ = v6
    return ("foo bar" + "hello")
}

func main() {
    var v1 uint8 = 128
    v2 := []int{820, 1, -9223372036854775808}
    v3 := []string{("hello" + "foo bar"), "", ""}
    v4 := []string{("\n\t" + "foo bar"), "a", ("a" + ((("" + ("a" + "\n\t")) + "a") + (("a" + ("" + "hello")) + ("a" + (("hello" + "") + ("\n\t" + "foo bar"))))))}
    if ((0.1 != float64((((0.1 - (1.5 + 0.0)) * 0.0) - 3.14159))) && (true || (!(4294967295 != 485)))) {
        v3 = append(v3, ("\n\t" + ("a" + "")))
        var v7 uint16 = uint16(v1)
        _ = v7
    } else {
        var v8 uint16 = 65535
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 42}
            _ = len(v3)
            v1 = 128
            _ = fi9
            _ = v10
        }
        switch ("hello" + "hello") {
        case "":
            var v11 uint64 = 603
            _ = v11
        case "hello":
            _ = len(v4)
        }
        _ = v8
    }
    v1 = (^(v1))
    switch 1 {
    case -646:
        v1 += 2
    case 1:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 97:
        v1 = 127
    default:
        var v12 uint32 = 337
        _ = v12
    }
    _ = len(v4)
    if false {
        var v13 string = (("" + "foo bar") + "foo bar")
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 float32 = 3.14159
            v15 = v15
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    if true {
        for fi16 := 0; fi16 < 3; fi16++ {
            if (!(false && false)) {
                v17 := GenS1{F0: 42}
                _ = v17
            }
            fmt.Println(v1)
            _ = fi16
        }
    } else {
        switch 32768 {
        case 848:
            for fi18 := 0; fi18 < 9; fi18++ {
                fmt.Println(fi18)
                v1 = (v1 << ((895) & 63))
                _ = fi18
            }
        case 51:
            if (true && ((float64(0.0) - ((1e10 / 0.1) - 3.14159)) != (3.14159 + (1.5 * float64(1e10))))) {
                if true {
                    switch -9223372036854775808 {
                    case -418:
                        for fi19 := 0; fi19 < 3; fi19++ {
                            var v20 float64 = 0.1
                            fmt.Println(fi19, v1)
                            var v21 uint16 = 32767
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    case 32768:
                        var v22 uint64 = 822
                        _ = v22
                    case -1:
                        v23 := GenS2{F0: 918, F1: 127}
                        _ = v23
                    default:
                        v1 = (v1 & (v1 >> ((4294967295) & 63)))
                    }
                    if (4294967295 >= 611) {
                        fmt.Println(v1)
                    } else {
                        v1 -= (v1 * (^(v1)))
                        fmt.Println(v1)
                    }
                    fmt.Println(v1)
                } else {
                    v1 += v1
                    if true {
                        switch "hello" {
                        case "foo bar":
                            if true {
                                if true {
                                    fmt.Println(v1)
                                    var v24 int16 = 1
                                    var v25 bool = (127 <= 0)
                                    _ = v24
                                    _ = v25
                                }
                            }
                        }
                    }
                }
                v26 := GenS1{F0: -32768}
                _ = v26
            }
        default:
            var v27 int32 = 0
            _ = v27
        }
        v28 := GenS1{F0: -1}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v28
    }
    var v29 string = ""
    if true {
        var v30 int64 = 42
        var v31 uint16 = 940
        for fi32 := 0; fi32 < 9; fi32++ {
            var v33 int8 = 127
            if len(v2) > 0 {
                v2[0] = fi32
            }
            if (false || true) {
                fmt.Println(v33)
            }
            _ = fi32
            _ = v33
        }
        _ = v30
        _ = v31
    } else {
        v2 = append(v2, 42)
    }
    v1 -= (^(v1))
    if ((!(2 <= 1)) || ((0.0 != 1.5) && (-2147483648 > -965))) {
        for fi34 := 0; fi34 < 10; fi34++ {
            for fi35 := 0; fi35 < 11; fi35++ {
                switch (^(fi34)) {
                case 32768:
                    v29 = v29
                case -233:
                    var v36 float64 = (1.5 * 0.0)
                    _ = v36
                case -606:
                    switch v29 {
                    case "":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
                _ = fi35
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                v1 = v1
            }
            _ = fi34
        }
        for fi37 := 0; fi37 < 1; fi37++ {
            v1 += (v1 / ((v1) | 1))
            if ((785 >= (fi37 ^ fi37)) || (false && (!((v1 * v1) <= v1)))) {
                if ((146 > 979) && ((538 != -32768) && true)) {
                    v38 := GenS2{F0: 2147483647, F1: 127}
                    fmt.Println(v1, fi37)
                    if (-128 != 11) {
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                        v1 = 42
                        if (((-2.25 + (((3.14159 * 0.0) + 0.1) + 1e10)) / 0.1) == (float64(0.0) * (float64(0.0) / 3.14159))) {
                            v39 := GenS2{F0: 2147483647, F1: -128}
                            for fi40 := 0; fi40 < 16; fi40++ {
                                switch -334 {
                                case 1:
                                    var v41 int16 = 1000
                                    v3 = append(v3, "")
                                    _ = v41
                                case -227:
                                    v42 := GenS3{F0: 42, F1: 255}
                                    _ = v42
                                case -9223372036854775808:
                                    switch -876 {
                                    case 1:
                                        v43 := GenS4{F0: 9223372036854775807}
                                        _ = v43
                                    }
                                }
                                var v44 uint = 103
                                var v45 uint8 = uint8(v44)
                                _ = fi40
                                _ = v44
                                _ = v45
                            }
                            v1 -= 42
                            _ = v39
                        }
                    }
                    _ = v38
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v1 = v1
                }
            }
            if (376 == 18446744073709551615) {
                var v46 float64 = 0.1
                _ = v46
            }
            _ = fi37
        }
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        v1 = 127
        _ = fi47
    }
    for fi48 := 0; fi48 < 9; fi48++ {
        fmt.Println(fi48)
        v1 += (v1 ^ 2)
        _ = fi48
    }
    v1 -= (v1 >> ((18446744073709551615) & 63))
    switch (v29 + v29) {
    case "hello":
        if (!true) {
            if false {
                for fi49 := 0; fi49 < 17; fi49++ {
                    v1 = 42
                    v1 += 255
                    if false {
                        var v50 string = (v29 + (v29 + v29))
                        if (!((!true) || ((!(!true)) || (!true)))) {
                            var v51 int8 = -128
                            _ = v51
                        }
                        _ = v50
                    } else {
                        fmt.Println(fi49, v29)
                        v52 := GenS3{F0: 32767, F1: 215}
                        v53 := GenS3{F0: 0, F1: v1}
                        _ = v52
                        _ = v53
                    }
                    _ = fi49
                }
                v54 := GenS5{F0: (("\n\t" + "\n\t") <= (v29 + (v29 + "a"))), F1: uint(v1)}
                var v55 string = v29
                _ = v54
                _ = v55
            } else {
                for fi56 := 0; fi56 < 11; fi56++ {
                    var v57 uint = 18446744073709551615
                    _ = fi56
                    _ = v57
                }
            }
            if ((!(!(112 == 1))) || false) {
                switch (v29 + v29) {
                case "hello":
                    v1 -= (v1 << ((42) & 63))
                default:
                    fmt.Println(v29, v1)
                }
                var v58 uint8 = 42
                _ = v58
            }
            switch -1 {
            case -2147483648:
                for fi59 := 0; fi59 < 11; fi59++ {
                    fmt.Println(v29, fi59, fi59)
                    var v60 float32 = 3.14159
                    var v61 uint32 = 42
                    _ = fi59
                    _ = v60
                    _ = v61
                }
            default:
                v29 = (v29 + "")
            }
        }
    case "":
        var v62 int16 = 42
        _ = v62
    case "a":
        if false {
            var v63 uint8 = (v1 + v1)
            v1 = (v1 ^ v1)
            _ = v63
        } else {
            if true {
                for fi64 := 0; fi64 < 20; fi64++ {
                    if true {
                        switch fi64 {
                        case 9223372036854775807:
                            v29 = "a"
                        case 618:
                            if (!(true && true)) {
                                v29 = (((("foo bar" + v29) + v29) + v29) + (v29 + "hello"))
                            }
                        case 967:
                            v3 = append(v3, "")
                        default:
                            for fi65 := 0; fi65 < 10; fi65++ {
                                v1 += (^(v1))
                                if (!(v1 >= (^(v1)))) {
                                    if (!false) {
                                        var v66 uint = uint(v1)
                                        fmt.Println(v29, fi65)
                                        _ = v66
                                    }
                                    v67 := GenS4{F0: 42}
                                    var v68 string = ("a" + v29)
                                    switch fi64 {
                                    case -307:
                                        var v69 int8 = 125
                                        _ = v69
                                    case -2147483648:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    v68 = (("foo bar" + ("a" + ("\n\t" + v68))) + ((v68 + v68) + ("foo bar" + "")))
                                    _ = v67
                                    _ = v68
                                } else {
                                    var v70 float64 = 3.14159
                                    switch "a" {
                                    case "\n\t":
                                        v29 = "a"
                                    default:
                                        if false {
                                            var v71 float32 = (float32((v70 / v70)) * (((float32(0.1) - (1e10 - 1e10)) * 0.1) + ((1e10 - 1.5) / (((1e10 / 1e10) + -2.25) + 0.0))))
                                            if (!(!false)) {
                                                v72 := GenS6{F0: (2147483647 == 970), F1: 4294967295, F2: ((18446744073709551615 > 65535) || (((v1 / (((v1 % ((v1) | 1))) | 1)) == (v1 / (((v1 % ((v1) | 1))) | 1))) || (487 > 255)))}
                                                var v73 uint32 = 65535
                                                v74 := GenS6{F0: false, F1: 525, F2: true}
                                                if (!false) {
                                                    var v75 float64 = (float64((-2.25 / float32(v70))) / 0.1)
                                                    switch 9223372036854775807 {
                                                    case 128:
                                                        v76 := GenS7{F0: -1, F1: "hello", F2: v29}
                                                        _ = v76
                                                    case -503:
                                                        switch "foo bar" {
                                                        case "":
                                                            var v77 uint = 65535
                                                            _ = v77
                                                        case "\n\t":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        default:
                                                            if len(v3) > 0 {
                                                                v3[0] = (v29 + "")
                                                            }
                                                        }
                                                    default:
                                                        var v78 uint = 2
                                                        _ = v78
                                                    }
                                                    _ = v75
                                                } else {
                                                    var v79 uint64 = 4294967295
                                                    _ = v79
                                                }
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                            } else {
                                                if (-1 == 79) {
                                                    if true {
                                                        v3 = append(v3, (v29 + v29))
                                                        v80 := GenS8{F0: (664 <= 42)}
                                                        if true {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            var v81 int = fi65
                                                            _ = v81
                                                        } else {
                                                            fmt.Println(v29)
                                                            var v82 float32 = v71
                                                            _ = v82
                                                        }
                                                        _ = v80
                                                    }
                                                    if true {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        v83 := GenS8{F0: (true || true)}
                                                        _ = v83
                                                    } else {
                                                        var v84 float64 = v70
                                                        _ = v84
                                                    }
                                                    switch v29 {
                                                    case "\n\t":
                                                        v70 = (v70 - (v70 * ((v70 / v70) - v70)))
                                                    case "":
                                                        var v85 uint32 = 180
                                                        _ = v85
                                                    case "foo bar":
                                                        if len(v2) > 0 {
                                                            v2[0] = (fi65 - (fi65 | fi64))
                                                        }
                                                    }
                                                    v29 = v29
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v86 int64 = -103
                                                    v1 -= 121
                                                    _ = v86
                                                }
                                            }
                                            _ = len(v4)
                                            _ = v71
                                        } else {
                                            v29 = v29
                                        }
                                    }
                                    _ = v70
                                }
                                _ = fi65
                            }
                        }
                    } else {
                        v1 = (v1 & (v1 - v1))
                    }
                    if (-32768 < 1) {
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v87 := GenS8{F0: ((!(294 != 1)) || true)}
                            if false {
                                switch ("\n\t" + v29) {
                                case "hello":
                                    v1 += v1
                                case "a":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "\n\t":
                                    switch 2147483647 {
                                    case -1:
                                        if true {
                                            for fi88 := 0; fi88 < 17; fi88++ {
                                                if (2 == 56) {
                                                    var v89 bool = true
                                                    _ = v89
                                                }
                                                _ = fi88
                                            }
                                        } else {
                                            v1 -= v1
                                            v1 = (v1 + (v1 ^ v1))
                                        }
                                    case 128:
                                        v29 = v29
                                    default:
                                        fmt.Println(v1)
                                    }
                                }
                                v1 = (v1 / ((37) | 1))
                            }
                            _ = v87
                        } else {
                            v3 = append(v3, v29)
                            for fi90 := 0; fi90 < 3; fi90++ {
                                fmt.Println(v29)
                                if len(v2) > 0 {
                                    v2[0] = fi64
                                }
                                var v91 uint16 = 0
                                _ = fi90
                                _ = v91
                            }
                        }
                        var v92 string = v29
                        if (65535 != 65535) {
                            v4 = append(v4, "hello")
                        }
                        _ = v92
                    }
                    _ = fi64
                }
            } else {
                for fi93 := 0; fi93 < 10; fi93++ {
                    var v94 bool = (0.0 < 3.14159)
                    v94 = v94
                    if false {
                        fmt.Println(v94, v1)
                        v1 = 221
                        var v95 uint32 = 2147483648
                        _ = v95
                    }
                    _ = fi93
                    _ = v94
                }
                for fi96 := 0; fi96 < 16; fi96++ {
                    for fi97 := 0; fi97 < 5; fi97++ {
                        if false {
                            var v98 int = fi96
                            _ = v98
                        } else {
                            v29 = "a"
                            fmt.Println(fi96, v29, v29)
                        }
                        var v99 int64 = -1
                        if (!true) {
                            switch (fi96 & fi96) {
                            case 32767:
                                v1 += (v1 / ((v1) | 1))
                            case 0:
                                v99 = (v99 & v99)
                            case -545:
                                if true {
                                    v3 = append(v3, "")
                                } else {
                                    switch -70 {
                                    case -322:
                                        var v100 uint = 1
                                        _ = v100
                                    case -9223372036854775808:
                                        switch "a" {
                                        case "hello":
                                            v29 = ""
                                        case "":
                                            fmt.Println(fi97, v1, v1)
                                        case "a":
                                            var v101 int64 = (v99 ^ v99)
                                            _ = v101
                                        default:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    case 0:
                                        if false {
                                            var v102 float32 = float32(0.1)
                                            _ = v102
                                        } else {
                                            _ = len(v4)
                                            var v103 uint8 = v1
                                            _ = v103
                                        }
                                    }
                                }
                            }
                            switch -2147483648 {
                            case 32767:
                                var v104 uint16 = 65535
                                _ = v104
                            case 32768:
                                switch fi96 {
                                case 2147483647:
                                    if true {
                                        var v105 int = fi97
                                        v1 += 42
                                        _ = v105
                                    } else {
                                        fmt.Println(v99, v1, v1)
                                        switch (fi97 + fi97) {
                                        case 4294967295:
                                            var v106 bool = true
                                            _ = v106
                                        case 2147483647:
                                            v99 = -1
                                        case 0:
                                            var v107 int32 = -2147483648
                                            _ = v107
                                        default:
                                            v1 += (v1 | (^(v1)))
                                        }
                                    }
                                case 426:
                                    var v108 int = fi96
                                    var v109 int8 = -1
                                    fmt.Println(v99)
                                    _ = v108
                                    _ = v109
                                case -9223372036854775808:
                                    if len(v2) > 0 {
                                        v2[0] = fi96
                                    }
                                }
                            }
                        } else {
                            var v110 int16 = -32768
                            var v111 int32 = -1
                            _ = v110
                            _ = v111
                        }
                        v3 = append(v3, ((((v29 + "foo bar") + (v29 + "hello")) + "\n\t") + (("a" + "a") + "foo bar")))
                        v112 := GenS9{F0: "a", F1: (v29 + v29)}
                        if len(v2) > 0 {
                            v2[0] = (fi97 / (((fi97 << ((692) & 63))) | 1))
                        }
                        v113 := GenS3{F0: 2, F1: (v1 % ((v1) | 1))}
                        _ = fi97
                        _ = v112
                        _ = v113
                        _ = v99
                    }
                    _ = fi96
                }
                if (15 != -2147483648) {
                    for fi114 := 0; fi114 < 14; fi114++ {
                        v2 = append(v2, fi114)
                        v29 = v29
                        _ = fi114
                    }
                }
            }
            switch -496 {
            case 1:
                for fi115 := 0; fi115 < 13; fi115++ {
                    v116 := GenS10{F0: 2, F1: 2147483647}
                    _ = fi115
                    _ = v116
                }
            case -10:
                v29 = v29
            }
            var v117 uint64 = 4294967295
            _ = v117
        }
    default:
        if (true || (!false)) {
            v118 := GenS6{F0: (!(306 <= 2147483648)), F1: 2147483647, F2: ((1000 != 473) || (-621 > int64(v1)))}
            for fi119 := 0; fi119 < 12; fi119++ {
                if false {
                    for fi120 := 0; fi120 < 15; fi120++ {
                        v3 = append(v3, v29)
                        _ = fi120
                    }
                    v121 := GenS11{F0: uint32(v1)}
                    for fi122 := 0; fi122 < 20; fi122++ {
                        fmt.Println(v29, v29, v1)
                        _ = fi122
                    }
                    _ = v121
                } else {
                    for fi123 := 0; fi123 < 11; fi123++ {
                        if (174 > 65535) {
                            v124 := GenS4{F0: -1}
                            if (fi123 > 32767) {
                                if true {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v125 := GenS12{F0: -313, F1: fi119}
                                        _ = v125
                                    default:
                                        v1 += v1
                                    }
                                    v1 = (v1 % (((v1 * (v1 | v1))) | 1))
                                } else {
                                    if (32 > -841) {
                                        switch fi123 {
                                        case 127:
                                            v1 -= (v1 ^ 200)
                                        }
                                        if true {
                                            if (((307 <= 813) && true) || (!(((^(fi119)) < (fi123 >> ((42) & 63))) && (1000 >= 1)))) {
                                                if false {
                                                    switch (fi119 % ((fi123) | 1)) {
                                                    case 458:
                                                        v1 += v1
                                                    case -623:
                                                        v29 = ("hello" + (v29 + v29))
                                                    }
                                                }
                                                fmt.Println(fi123)
                                                if false {
                                                    if true {
                                                        switch (fi119 + fi119) {
                                                        case -973:
                                                            _ = len(v4)
                                                        case 652:
                                                            v29 = "foo bar"
                                                        default:
                                                            v126 := GenS3{F0: 0, F1: (^(v1))}
                                                            _ = v126
                                                        }
                                                        if true {
                                                            v29 = v29
                                                        } else {
                                                            v1 += v1
                                                            v127 := GenS4{F0: int64(v1)}
                                                            v1 = (^(v1))
                                                            _ = len(v4)
                                                            _ = v127
                                                        }
                                                    } else {
                                                        var v128 uint = 18446744073709551615
                                                        _ = v128
                                                    }
                                                    v129 := GenS13{F0: 113}
                                                    _ = v129
                                                } else {
                                                    if len(v4) > 0 {
                                                        v4[0] = (("foo bar" + (v29 + (v29 + ""))) + v29)
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v29)
                                            }
                                            fmt.Println(fi123, fi123)
                                        } else {
                                            v29 = "hello"
                                            if (-9223372036854775808 >= 9223372036854775807) {
                                                v29 = "a"
                                                switch fi119 {
                                                case -9223372036854775808:
                                                    switch v29 {
                                                    case "":
                                                        v130 := GenS10{F0: 0, F1: fi119}
                                                        _ = v130
                                                    case "foo bar":
                                                        v1 = (v1 - 2)
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                case 9223372036854775807:
                                                    fmt.Println(fi123, v1)
                                                case 4294967295:
                                                    if (!(!(79 >= 79))) {
                                                        var v131 uint32 = 42
                                                        _ = v131
                                                    } else {
                                                        var v132 bool = false
                                                        var v133 float64 = (float64((0.0 * -2.25)) / 1.0)
                                                        v134 := GenS8{F0: (v1 >= (v1 & 2))}
                                                        _ = v132
                                                        _ = v133
                                                        _ = v134
                                                    }
                                                }
                                            } else {
                                                if (("\n\t" + (("a" + "a") + ("foo bar" + v29))) < v29) {
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi123 << ((1) & 63))
                                                    }
                                                }
                                                var v135 int8 = 0
                                                _ = v135
                                            }
                                            switch fi119 {
                                            case 128:
                                                v29 = "foo bar"
                                            }
                                        }
                                    } else {
                                        v29 = ("hello" + ("hello" + "\n\t"))
                                        v29 = (("" + v29) + v29)
                                        var v136 bool = (0 >= 477)
                                        _ = v136
                                    }
                                }
                            } else {
                                v29 = v29
                                switch (fi119 & fi119) {
                                case 32767:
                                    var v137 uint = 18446744073709551615
                                    _ = v137
                                case -1:
                                    if false {
                                        fmt.Println(fi119)
                                    }
                                    switch fi123 {
                                    case -917:
                                        switch -2147483648 {
                                        case 128:
                                            var v138 int = fi119
                                            _ = v138
                                        case -1:
                                            var v139 uint = 697
                                            _ = v139
                                        case 127:
                                            var v140 uint16 = uint16(v1)
                                            _ = v140
                                        }
                                    case -249:
                                        switch ("a" + "hello") {
                                        case "":
                                            if false {
                                                switch fi119 {
                                                case 0:
                                                    if (2 > 18446744073709551615) {
                                                        fmt.Println(v29)
                                                        switch fi123 {
                                                        case 32767:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        case -282:
                                                            v29 = v29
                                                        case -2147483648:
                                                            var v141 uint16 = 578
                                                            _ = v141
                                                        }
                                                        v4 = append(v4, v29)
                                                    }
                                                }
                                                v142 := GenS14{F0: 4294967295}
                                                v29 = (v29 + v29)
                                                fmt.Println(v1, v29)
                                                if true {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if ((v29 + v29) != v29) {
                                                        switch fi119 {
                                                        case -894:
                                                            fmt.Println(v1, fi123)
                                                        }
                                                    } else {
                                                        v1 += v1
                                                        if false {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            switch "a" {
                                                            case "a":
                                                                v3 = append(v3, v29)
                                                            default:
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            }
                                                            v1 += 2
                                                        } else {
                                                            if true {
                                                                switch v29 {
                                                                case "":
                                                                    var v143 uint32 = 2147483647
                                                                    _ = v143
                                                                case "hello":
                                                                    var v144 int32 = 2147483647
                                                                    _ = v144
                                                                case "\n\t":
                                                                    if false {
                                                                        var v145 int16 = 11
                                                                        v145 -= v145
                                                                        _ = v145
                                                                    } else {
                                                                        v146 := GenS10{F0: 65535, F1: fi119}
                                                                        var v147 int8 = 0
                                                                        switch fi119 {
                                                                        case 9223372036854775807:
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        }
                                                                        _ = v146
                                                                        _ = v147
                                                                    }
                                                                }
                                                            } else {
                                                                switch v29 {
                                                                case "":
                                                                    switch (fi123 ^ (fi123 >> ((255) & 63))) {
                                                                    case -706:
                                                                        if (1 == 245) {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            v148 := GenS15{F0: 1, F1: 1000}
                                                                            _ = v148
                                                                        }
                                                                    case -793:
                                                                        switch v29 {
                                                                        case "":
                                                                            v1 -= 0
                                                                        case "a":
                                                                            var v149 float32 = (-2.25 * 1.5)
                                                                            _ = v149
                                                                        case "hello":
                                                                            switch v29 {
                                                                            case "foo bar":
                                                                                v2 = append(v2, 4294967295)
                                                                            default:
                                                                                if false {
                                                                                    if (-1 <= 136) {
                                                                                        v150 := GenS9{F0: v29, F1: (v29 + v29)}
                                                                                        _ = v150
                                                                                    }
                                                                                    v1 -= 128
                                                                                    if false {
                                                                                        var v151 int32 = -1
                                                                                        _ = v151
                                                                                    } else {
                                                                                        fmt.Println(v1, v29, v1)
                                                                                        if (-1 <= -2147483648) {
                                                                                            v152 := GenS16{F0: "", F1: 2147483647, F2: v1}
                                                                                            _ = v152
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    case 4294967295:
                                                                        v153 := GenS17{F0: 1, F1: fi123, F2: 2147483648}
                                                                        _ = v153
                                                                    }
                                                                case "\n\t":
                                                                    v29 = v29
                                                                    v2 = append(v2, (fi123 - fi119))
                                                                    var v154 int32 = 1
                                                                    _ = v154
                                                                case "a":
                                                                    var v155 string = ""
                                                                    _ = v155
                                                                default:
                                                                    if false {
                                                                        if ((1 > 2147483648) && false) {
                                                                            fmt.Println(v29, fi123, v29)
                                                                            v156 := GenS16{F0: (v29 + "\n\t"), F1: 0, F2: v1}
                                                                            var v157 uint = 2
                                                                            fmt.Println(fi119, v29, v29)
                                                                            switch fi123 {
                                                                            case -1:
                                                                                switch "\n\t" {
                                                                                case "\n\t":
                                                                                    var v158 int8 = 100
                                                                                    _ = v158
                                                                                }
                                                                                var v159 uint16 = 1
                                                                                _ = v159
                                                                            case 4294967295:
                                                                                if true {
                                                                                    v160 := GenS13{F0: -128}
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = (fi123 / (((fi123 | (fi119 / ((fi119) | 1)))) | 1))
                                                                                    }
                                                                                    var v161 int32 = 0
                                                                                    _ = v160
                                                                                    _ = v161
                                                                                } else {
                                                                                    var v162 int32 = 1
                                                                                    var v163 int64 = -1
                                                                                    v162 = (v162 - -339)
                                                                                    _ = v162
                                                                                    _ = v163
                                                                                }
                                                                            default:
                                                                                fmt.Println(v157)
                                                                            }
                                                                            _ = v156
                                                                            _ = v157
                                                                        }
                                                                        v29 = "hello"
                                                                        fmt.Println(v1)
                                                                    } else {
                                                                        if true {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            v29 = v29
                                                                            var v164 uint32 = 597
                                                                            _ = v164
                                                                        }
                                                                        if (int64(fi119) <= -1) {
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                if (2147483647 == -2147483648) {
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                } else {
                                                                                    v1 = v1
                                                                                    v1 = (v1 | (v1 - v1))
                                                                                }
                                                                            }
                                                                        }
                                                                        if true {
                                                                            v29 = v29
                                                                            var v165 uint32 = 380
                                                                            _ = v165
                                                                        } else {
                                                                            var v166 int8 = 0
                                                                            if false {
                                                                                var v167 float64 = 0.1
                                                                                _ = v167
                                                                            } else {
                                                                                fmt.Println(fi119, fi123, v29)
                                                                                v166 += (-(v166))
                                                                                var v168 int16 = -32768
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = ("foo bar" + (v29 + (v29 + (v29 + v29))))
                                                                                }
                                                                                v4 = append(v4, v29)
                                                                                _ = v168
                                                                            }
                                                                            fmt.Println(v166, fi123, v1)
                                                                            _ = v166
                                                                        }
                                                                        v169 := GenS18{F0: 4294967295, F1: -2147483648, F2: 42}
                                                                        var v170 float64 = 0.1
                                                                        _ = v169
                                                                        _ = v170
                                                                    }
                                                                }
                                                                if ((-1 <= 0) || true) {
                                                                    var v171 uint = 42
                                                                    _ = v171
                                                                } else {
                                                                    if true {
                                                                        if false {
                                                                            v172 := GenS8{F0: (0 <= 562)}
                                                                            var v173 uint8 = 127
                                                                            switch fi123 {
                                                                            case 461:
                                                                                var v174 bool = (!false)
                                                                                _ = v174
                                                                            }
                                                                            _ = v172
                                                                            _ = v173
                                                                        } else {
                                                                            var v175 float32 = 0.1
                                                                            if (!true) {
                                                                                if (true || true) {
                                                                                    fmt.Println(fi123, fi123, v175)
                                                                                    v176 := GenS1{F0: 1}
                                                                                    v3 = append(v3, v29)
                                                                                    _ = v176
                                                                                } else {
                                                                                    var v177 int32 = 2147483647
                                                                                    if ((v1 + 255) > (v1 ^ (v1 + v1))) {
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            v2 = append(v2, 32767)
                                                                                        case "a":
                                                                                            switch 0 {
                                                                                            case -2147483648:
                                                                                                switch 32767 {
                                                                                                case 860:
                                                                                                    v178 := GenS19{F0: -483, F1: v29, F2: (v29 + ("a" + "hello"))}
                                                                                                    _ = v178
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if false {
                                                                                            switch (fi119 * fi119) {
                                                                                            case 9223372036854775807:
                                                                                                v29 = v29
                                                                                            case 667:
                                                                                                v179 := GenS11{F0: uint32(v1)}
                                                                                                _ = v179
                                                                                            }
                                                                                        } else {
                                                                                            var v180 int16 = int16(v1)
                                                                                            v180 = v180
                                                                                            _ = v180
                                                                                        }
                                                                                    } else {
                                                                                        if (-1 >= fi123) {
                                                                                            var v181 float32 = v175
                                                                                            v1 = v1
                                                                                            _ = v181
                                                                                        }
                                                                                        if (true && true) {
                                                                                            if false {
                                                                                                v4 = append(v4, (v29 + v29))
                                                                                                switch v29 {
                                                                                                case "a":
                                                                                                    var v182 uint32 = 4294967295
                                                                                                    _ = v182
                                                                                                case "hello":
                                                                                                    fmt.Println(v177, v177)
                                                                                                default:
                                                                                                    if false {
                                                                                                        v183 := GenS20{F0: 65535}
                                                                                                        if ((v177 + v177) == v177) {
                                                                                                            v184 := GenS5{F0: true, F1: 0}
                                                                                                            var v185 string = v29
                                                                                                            _ = v184
                                                                                                            _ = v185
                                                                                                        } else {
                                                                                                            var v186 int = fi119
                                                                                                            _ = v186
                                                                                                        }
                                                                                                        _ = v183
                                                                                                    } else {
                                                                                                        v177 -= v177
                                                                                                        v177 += -57
                                                                                                    }
                                                                                                }
                                                                                                var v187 string = ("" + v29)
                                                                                                _ = v187
                                                                                            } else {
                                                                                                _ = len(v4)
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v188 int32 = 1
                                                                                        v29 = v29
                                                                                        var v189 float64 = 1e10
                                                                                        _ = v188
                                                                                        _ = v189
                                                                                    }
                                                                                    if (true && (!true)) {
                                                                                        if (95 >= 32767) {
                                                                                            var v190 int = (-(fi119))
                                                                                            if (268 < 42) {
                                                                                                if (true && (uint(v1) != 1)) {
                                                                                                    _ = len(v2)
                                                                                                    v191 := GenS9{F0: v29, F1: v29}
                                                                                                    var v192 int16 = 32767
                                                                                                    _ = v191
                                                                                                    _ = v192
                                                                                                }
                                                                                                v193 := GenS8{F0: true}
                                                                                                switch "" {
                                                                                                case "a":
                                                                                                    v194 := GenS21{F0: -1, F1: 2, F2: 32767}
                                                                                                    _ = v194
                                                                                                case "hello":
                                                                                                    v195 := GenS5{F0: (95 <= 877), F1: 1}
                                                                                                    _ = v195
                                                                                                case "\n\t":
                                                                                                    if false {
                                                                                                        v175 -= float32(0.1)
                                                                                                    }
                                                                                                }
                                                                                                _ = v193
                                                                                            } else {
                                                                                                v29 = v29
                                                                                            }
                                                                                            var v196 int = (v190 + -1)
                                                                                            _ = v190
                                                                                            _ = v196
                                                                                        }
                                                                                        var v197 int16 = 42
                                                                                        _ = v197
                                                                                    } else {
                                                                                        v198 := GenS5{F0: false, F1: 4294967295}
                                                                                        _ = v198
                                                                                    }
                                                                                    v29 = v29
                                                                                    v199 := GenS12{F0: v177, F1: (fi123 ^ 9223372036854775807)}
                                                                                    _ = v177
                                                                                    _ = v199
                                                                                }
                                                                                v29 = v29
                                                                                v200 := GenS22{F0: true}
                                                                                _ = v200
                                                                            } else {
                                                                                v29 = v29
                                                                                fmt.Println(v29, fi119, fi119)
                                                                            }
                                                                            _ = v175
                                                                        }
                                                                    } else {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        v29 = v29
                                                                        _ = len(v3)
                                                                    }
                                                                }
                                                            }
                                                            if (124 == -128) {
                                                                if (("foo bar" + ("foo bar" + v29)) != v29) {
                                                                    fmt.Println(v29)
                                                                    v1 -= 128
                                                                } else {
                                                                    var v201 int32 = -274
                                                                    _ = v201
                                                                }
                                                                v4 = append(v4, (v29 + v29))
                                                                _ = len(v2)
                                                            }
                                                        }
                                                        var v202 int16 = 1
                                                        _ = v202
                                                    }
                                                }
                                                _ = v142
                                            }
                                        }
                                    default:
                                        v1 -= v1
                                    }
                                    v29 = v29
                                }
                                if (true || (1e10 >= (1.5 - 0.1))) {
                                    v29 = v29
                                }
                            }
                            v203 := GenS23{F0: 1000}
                            _ = v124
                            _ = v203
                        }
                        if false {
                            var v204 bool = (42 < 127)
                            var v205 uint32 = 42
                            var v206 int64 = -417
                            if false {
                                fmt.Println(v205, fi119, v205)
                                if (1 != 2147483647) {
                                    fmt.Println(v204, v206)
                                    switch v29 {
                                    case "foo bar":
                                        if (v204 && (18446744073709551615 < 4294967295)) {
                                            if false {
                                                var v207 int64 = v206
                                                v1 -= (v1 % ((v1) | 1))
                                                switch fi123 {
                                                case -9223372036854775808:
                                                    v208 := GenS21{F0: -32768, F1: 172, F2: 42}
                                                    _ = v208
                                                case 165:
                                                    if false {
                                                        var v209 uint32 = v205
                                                        var v210 int64 = (v207 - v207)
                                                        var v211 float64 = 3.14159
                                                        _ = v209
                                                        _ = v210
                                                        _ = v211
                                                    }
                                                case 128:
                                                    v29 = v29
                                                default:
                                                    var v212 int8 = 110
                                                    _ = v212
                                                }
                                                _ = v207
                                            }
                                        }
                                    case "\n\t":
                                        var v213 string = v29
                                        _ = v213
                                    }
                                    v205 = (v205 - 42)
                                } else {
                                    var v214 uint64 = 42
                                    var v215 uint64 = v214
                                    _ = v214
                                    _ = v215
                                }
                            } else {
                                fmt.Println(v204, v29, v206)
                                v29 = "hello"
                                switch fi119 {
                                case 128:
                                    if (42 > 42) {
                                        if (!false) {
                                            v205 = 4294967295
                                            v204 = (-32768 > 396)
                                        } else {
                                            v206 = (v206 * v206)
                                            switch "foo bar" {
                                            case "":
                                                var v216 uint8 = v1
                                                _ = v216
                                            case "a":
                                                if true {
                                                    var v217 int8 = -1
                                                    var v218 string = ("foo bar" + ((v29 + "a") + v29))
                                                    var v219 uint32 = (v205 % (((v205 ^ v205)) | 1))
                                                    _ = v217
                                                    _ = v218
                                                    _ = v219
                                                }
                                            case "foo bar":
                                                v3 = append(v3, v29)
                                            }
                                            if (!false) {
                                                if ((((1e10 - 0.1) * 0.0) <= 3.14159) && v204) {
                                                    v220 := GenS12{F0: 0, F1: fi123}
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if (!(!((v206 << ((65535) & 63)) > v206))) {
                                                        switch (fi119 & fi119) {
                                                        case 49:
                                                            v29 = ((v29 + "\n\t") + (((v29 + ("\n\t" + v29)) + (v29 + (v29 + ("\n\t" + (v29 + v29))))) + v29))
                                                        case 128:
                                                            v221 := GenS24{F0: (-(fi119)), F1: -28, F2: 324}
                                                            _ = v221
                                                        case 1:
                                                            if (32767 <= uint16(v1)) {
                                                                switch (^(fi119)) {
                                                                case 2147483647:
                                                                    v205 -= v205
                                                                case 1:
                                                                    v204 = v204
                                                                default:
                                                                    switch (fi123 - -2147483648) {
                                                                    case 32767:
                                                                        _ = len(v3)
                                                                    default:
                                                                        fmt.Println(v29)
                                                                    }
                                                                    if len(v4) > 0 {
                                                                        v4[0] = (("" + v29) + "")
                                                                    }
                                                                    v206 -= v206
                                                                    v222 := GenS16{F0: v29, F1: (v205 + v205), F2: (v1 & v1)}
                                                                    v204 = true
                                                                    v205 = 4294967295
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    v1 -= 234
                                                                    _ = v222
                                                                }
                                                            } else {
                                                                var v223 uint = 4294967295
                                                                _ = v223
                                                            }
                                                        }
                                                        v206 = v206
                                                        var v224 int16 = -121
                                                        _ = v224
                                                    } else {
                                                        switch ((v29 + "a") + "hello") {
                                                        case "\n\t":
                                                            if v204 {
                                                                var v225 uint32 = (v205 - v205)
                                                                switch "" {
                                                                case "":
                                                                    var v226 int = fi123
                                                                    _ = v226
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = ("" + "")
                                                                }
                                                                _ = v225
                                                            } else {
                                                                fmt.Println(fi123, v206, fi123)
                                                                switch fi123 {
                                                                case -9223372036854775808:
                                                                    var v227 int32 = 42
                                                                    _ = v227
                                                                default:
                                                                    v204 = true
                                                                }
                                                                switch (fi123 - fi123) {
                                                                case 1:
                                                                    var v228 uint8 = v1
                                                                    _ = v228
                                                                case -89:
                                                                    var v229 int = fi119
                                                                    _ = v229
                                                                case 0:
                                                                    fmt.Println(v205, v206)
                                                                }
                                                            }
                                                        case "hello":
                                                            v1 = (v1 ^ 42)
                                                        case "a":
                                                            var v230 bool = ((false && (-1 > 93)) && (!(!(-2.25 < (-2.25 + 3.14159)))))
                                                            _ = v230
                                                        default:
                                                            if (42 != 893) {
                                                                v29 = v29
                                                                var v231 uint64 = 76
                                                                fmt.Println(v205, v1, fi119)
                                                                _ = v231
                                                            } else {
                                                                var v232 int = fi119
                                                                var v233 int16 = -1
                                                                if false {
                                                                    var v234 uint32 = (v205 >> ((2) & 63))
                                                                    var v235 uint32 = v205
                                                                    _ = v234
                                                                    _ = v235
                                                                } else {
                                                                    v29 = v29
                                                                    var v236 int64 = (v206 & v206)
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v232
                                                                    }
                                                                    _ = v236
                                                                }
                                                                _ = v232
                                                                _ = v233
                                                            }
                                                        }
                                                        fmt.Println(v206, v205)
                                                    }
                                                    v4 = append(v4, "a")
                                                    v237 := GenS25{F0: 1}
                                                    var v238 uint32 = 171
                                                    v206 = v206
                                                    v206 -= (v206 & v206)
                                                    _ = v220
                                                    _ = v237
                                                    _ = v238
                                                } else {
                                                    if true {
                                                        var v239 uint16 = 907
                                                        var v240 int32 = -2147483648
                                                        _ = v239
                                                        _ = v240
                                                    }
                                                }
                                                v204 = v204
                                                v241 := GenS16{F0: v29, F1: (v205 & (v205 % ((v205) | 1))), F2: v1}
                                                _ = v241
                                            }
                                            if v204 {
                                                if ((v1 * v1) < v1) {
                                                    var v242 uint = 42
                                                    v206 = (v206 / ((v206) | 1))
                                                    if (!false) {
                                                        v206 = v206
                                                        switch int(v206) {
                                                        case 9223372036854775807:
                                                            v243 := GenS26{F0: 2, F1: (v1 * v1), F2: 710}
                                                            _ = v243
                                                        case 175:
                                                            if len(v2) > 0 {
                                                                v2[0] = fi119
                                                            }
                                                        default:
                                                            var v244 uint16 = 32767
                                                            _ = v244
                                                        }
                                                    }
                                                    _ = v242
                                                } else {
                                                    var v245 uint16 = 0
                                                    if len(v2) > 0 {
                                                        v2[0] = fi119
                                                    }
                                                    _ = len(v3)
                                                    _ = v245
                                                }
                                            }
                                            v204 = (((v206 & v206) > (v206 - (-(v206)))) && false)
                                        }
                                    }
                                case 32768:
                                    v205 = 570
                                    switch "a" {
                                    case "hello":
                                        if (int32(v1) != 0) {
                                            switch "hello" {
                                            case "\n\t":
                                                var v246 int16 = -32768
                                                _ = v246
                                            case "foo bar":
                                                var v247 uint8 = 0
                                                _ = v247
                                            }
                                            var v248 int32 = 2147483647
                                            fmt.Println(v248, v29)
                                            _ = v248
                                        }
                                    }
                                    var v249 uint32 = v205
                                    if v204 {
                                        if (v29 < "foo bar") {
                                            v249 = 77
                                            _ = len(v3)
                                            if v204 {
                                                v250 := GenS27{F0: (v1 ^ 42)}
                                                _ = v250
                                            }
                                            switch v29 {
                                            case "\n\t":
                                                v204 = (1 >= -551)
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = "hello"
                                            }
                                        } else {
                                            var v251 string = "hello"
                                            _ = v251
                                        }
                                        v1 -= v1
                                    }
                                    _ = v249
                                }
                                var v252 bool = ((-(v206)) <= (v206 | -755))
                                _ = v252
                            }
                            switch ("foo bar" + ("foo bar" + v29)) {
                            case "foo bar":
                                _ = len(v2)
                            case "\n\t":
                                v204 = (!(414 >= 2))
                            }
                            _ = v204
                            _ = v205
                            _ = v206
                        }
                        if (false || (!true)) {
                            v1 = v1
                        }
                        switch (v29 + "\n\t") {
                        case "\n\t":
                            v29 = (v29 + v29)
                        case "hello":
                            var v253 int8 = 127
                            _ = v253
                        default:
                            if false {
                                var v254 bool = (493 < 42)
                                var v255 int16 = 42
                                v255 = 93
                                _ = v254
                                _ = v255
                            } else {
                                var v256 int8 = 100
                                var v257 uint32 = 44
                                if ((0 != -9223372036854775808) || ((!true) || (!(471 != 400)))) {
                                    v4 = append(v4, ("hello" + (v29 + v29)))
                                    if false {
                                        v258 := GenS1{F0: -108}
                                        _ = v258
                                    }
                                } else {
                                    v29 = ("\n\t" + v29)
                                    var v259 int32 = 2147483647
                                    _ = v259
                                }
                                _ = v256
                                _ = v257
                            }
                        }
                        fmt.Println(v1, fi123)
                        var v260 int8 = -50
                        var v261 string = "a"
                        _ = fi123
                        _ = v260
                        _ = v261
                    }
                    v1 = (v1 / ((127) | 1))
                }
                if false {
                    v29 = v29
                    switch (fi119 | (fi119 % (((fi119 & (-(fi119)))) | 1))) {
                    case -108:
                        switch 127 {
                        case 127:
                            v262 := GenS13{F0: -90}
                            _ = v262
                        case -574:
                            var v263 float64 = 1.5
                            _ = v263
                        case 146:
                            if false {
                                v1 = v1
                                for fi264 := 0; fi264 < 16; fi264++ {
                                    v265 := GenS23{F0: 30}
                                    _ = fi264
                                    _ = v265
                                }
                            }
                        default:
                            if ((^(v1)) <= 2) {
                                if (false && true) {
                                    var v266 string = "foo bar"
                                    fmt.Println(v266, v29, v1)
                                    _ = v266
                                }
                            }
                        }
                    }
                    if (!((2 <= 272) && true)) {
                        if true {
                            if false {
                                switch (fi119 % (((fi119 ^ (fi119 | fi119))) | 1)) {
                                case 128:
                                    for fi267 := 0; fi267 < 8; fi267++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v1 = 127
                                        _ = fi267
                                    }
                                case 9223372036854775807:
                                    switch (^(fi119)) {
                                    case -1:
                                        for fi268 := 0; fi268 < 20; fi268++ {
                                            var v269 uint = 1
                                            var v270 uint64 = 1
                                            var v271 int8 = int8(v269)
                                            _ = fi268
                                            _ = v269
                                            _ = v270
                                            _ = v271
                                        }
                                    case 128:
                                        if true {
                                            for fi272 := 0; fi272 < 15; fi272++ {
                                                fmt.Println(fi119, v1)
                                                fmt.Println(v29, fi119)
                                                _ = fi272
                                            }
                                        } else {
                                            switch (fi119 / (((fi119 - (fi119 | fi119))) | 1)) {
                                            case 1:
                                                fmt.Println(fi119, fi119)
                                            }
                                            var v273 uint16 = 2
                                            if (!(2 <= (v273 | (v273 * (v273 + (^(v273))))))) {
                                                v274 := GenS28{F0: (-(fi119))}
                                                if (!false) {
                                                    if (112 < 65535) {
                                                        var v275 uint = 255
                                                        _ = v275
                                                    }
                                                }
                                                _ = v274
                                            }
                                            _ = v273
                                        }
                                    case -2147483648:
                                        if len(v3) > 0 {
                                            v3[0] = (v29 + "")
                                        }
                                    }
                                }
                                fmt.Println(fi119)
                                for fi276 := 0; fi276 < 9; fi276++ {
                                    v29 = v29
                                    fmt.Println(fi119, fi119, fi119)
                                    _ = fi276
                                }
                            }
                            v277 := GenS29{F0: 65535, F1: 402, F2: (v29 + ("hello" + "foo bar"))}
                            for fi278 := 0; fi278 < 11; fi278++ {
                                var v279 int32 = 42
                                if (2147483647 >= 2147483648) {
                                    if (353 == -254) {
                                        if (-9223372036854775808 <= 51) {
                                            switch "foo bar" {
                                            case "hello":
                                                v1 -= 211
                                            }
                                            v29 = ("foo bar" + v29)
                                        } else {
                                            v1 += 2
                                            if (18446744073709551615 == 42) {
                                                _ = len(v4)
                                                v280 := GenS3{F0: 65535, F1: 80}
                                                _ = v280
                                            }
                                            v1 = 127
                                        }
                                    } else {
                                        v1 = (^(v1))
                                        if (18446744073709551615 < uint(v1)) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v281 := GenS30{F0: "hello", F1: (!true)}
                                            _ = v281
                                        }
                                        fmt.Println(v279)
                                    }
                                    if true {
                                        fmt.Println(v29, v279)
                                    }
                                } else {
                                    v279 = v279
                                }
                                _ = fi278
                                _ = v279
                            }
                            _ = v277
                        } else {
                            var v282 float64 = -2.25
                            var v283 string = ((("\n\t" + v29) + ("hello" + (v29 + ("a" + "")))) + "")
                            _ = v282
                            _ = v283
                        }
                        var v284 uint32 = 468
                        _ = v284
                    }
                }
                _ = fi119
            }
            if false {
                if true {
                    if len(v2) > 0 {
                        v2[0] = 32767
                    }
                    for fi285 := 0; fi285 < 19; fi285++ {
                        switch 9223372036854775807 {
                        case 4294967295:
                            for fi286 := 0; fi286 < 18; fi286++ {
                                v287 := GenS31{F0: 160, F1: v29, F2: fi285}
                                _ = fi286
                                _ = v287
                            }
                        default:
                            _ = len(v3)
                        }
                        if (0 != 65535) {
                            v29 = ""
                        } else {
                            v29 = "hello"
                            v29 = ("foo bar" + "\n\t")
                            var v288 uint32 = 484
                            _ = v288
                        }
                        _ = fi285
                    }
                    v29 = v29
                }
            }
            _ = v118
        } else {
            v1 -= v1
        }
    }
    v29 = v29
    for fi289 := 0; fi289 < 7; fi289++ {
        _ = len(v2)
        v29 = (v29 + "")
        v1 = 251
        _ = fi289
    }
    v290 := GenS14{F0: 85}
    for fi291 := 0; fi291 < 14; fi291++ {
        switch "" {
        case "":
            v292 := GenS32{F0: v29}
            _ = v292
        case "\n\t":
            for fi293 := 0; fi293 < 18; fi293++ {
                fmt.Println(fi291, fi293)
                _ = fi293
            }
        default:
            var v294 uint16 = 905
            _ = v294
        }
        var v295 int32 = 42
        _ = fi291
        _ = v295
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v290
    _ = v3
    _ = v4
}
