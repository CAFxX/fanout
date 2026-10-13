package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint; F1 int32; F2 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint; F2 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8; F1 uint64; F2 uint }
type GenS9 struct { F0 uint32; F1 int64; F2 int16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int; F1 int32 }
type GenS12 struct { F0 int32; F1 int; F2 int }
type GenS13 struct { F0 uint64; F1 bool }
type GenS14 struct { F0 int16; F1 int }
type GenS15 struct { F0 int16; F1 int64; F2 uint }
type GenS16 struct { F0 uint32; F1 uint32 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 int8; F1 uint16 }
type GenS21 struct { F0 uint8; F1 uint16; F2 int }
type GenS22 struct { F0 uint64; F1 uint }

func main() {
    var v1 string = ("foo bar" + "a")
    v2 := []string{v1}
    v3 := []uint32{42, 4294967295, 315}
    var v4 string = (v1 + "\n\t")
    v5 := GenS1{F0: v4, F1: 32767}
    switch (v1 + (v4 + "\n\t")) {
    case "hello":
        v3 = append(v3, 2)
    case "\n\t":
        switch "foo bar" {
        case "a":
            v4 = (v1 + (v1 + v1))
        default:
            switch -9223372036854775808 {
            case -250:
                switch 32768 {
                case -1:
                    var v6 bool = false
                    _ = v6
                case 1:
                    for fi7 := 0; fi7 < 9; fi7++ {
                        var v8 int8 = -128
                        _ = fi7
                        _ = v8
                    }
                }
            default:
                if false {
                    v1 = (v4 + v1)
                }
            }
        }
    }
    if (!false) {
        switch "hello" {
        case "hello":
            v3 = append(v3, 42)
        default:
            switch (("foo bar" + "foo bar") + v1) {
            case "":
                v4 = (v1 + ((v4 + v4) + ("a" + "")))
            case "hello":
                var v9 uint8 = 0
                _ = v9
            default:
                if true {
                    v4 = ("foo bar" + v4)
                    if true {
                        for fi10 := 0; fi10 < 8; fi10++ {
                            for fi11 := 0; fi11 < 2; fi11++ {
                                if false {
                                    var v12 int8 = 0
                                    var v13 uint64 = 1
                                    _ = v12
                                    _ = v13
                                }
                                v14 := GenS1{F0: v4, F1: (fi11 - 32768)}
                                v1 = "hello"
                                _ = fi11
                                _ = v14
                            }
                            for fi15 := 0; fi15 < 15; fi15++ {
                                if (-6 >= -1) {
                                    var v16 uint32 = 908
                                    _ = v16
                                } else {
                                    var v17 uint16 = 418
                                    var v18 uint = 73
                                    v17 -= v17
                                    _ = v17
                                    _ = v18
                                }
                                if true {
                                    if (false || (v4 != v1)) {
                                        if ((fi10 ^ fi15) != fi10) {
                                            if (!(0 > 224)) {
                                                fmt.Println(v1, fi15, v1)
                                                v4 = (v1 + "a")
                                                v3 = append(v3, 1)
                                                fmt.Println(v4, fi10)
                                                switch (fi15 ^ fi15) {
                                                case 128:
                                                    var v19 bool = true
                                                    _ = v19
                                                default:
                                                    var v20 bool = false
                                                    _ = v20
                                                }
                                            } else {
                                                var v21 int32 = -825
                                                v4 = v1
                                                _ = v21
                                            }
                                            v1 = ("foo bar" + v1)
                                            v1 = v4
                                            v4 = ("a" + v1)
                                        } else {
                                            v1 = (("hello" + v1) + v1)
                                            v2 = append(v2, (("a" + ("\n\t" + ("" + v1))) + ("a" + ((v4 + "\n\t") + v1))))
                                        }
                                    }
                                    switch 32768 {
                                    case 128:
                                        switch fi10 {
                                        case 2147483647:
                                            v1 = v1
                                            v22 := GenS1{F0: ("hello" + v4), F1: (^(fi15))}
                                            _ = v22
                                        case 4294967295:
                                            var v23 int8 = 42
                                            _ = v23
                                        default:
                                            if (!false) {
                                                if false {
                                                    v24 := GenS1{F0: v4, F1: fi10}
                                                    _ = v24
                                                }
                                                switch (fi10 << ((18446744073709551615) & 63)) {
                                                case -364:
                                                    v1 = v4
                                                }
                                                var v25 int64 = -561
                                                _ = v25
                                            }
                                        }
                                        v26 := GenS1{F0: v1, F1: fi10}
                                        _ = v26
                                    }
                                    v1 = v1
                                }
                                switch "a" {
                                case "":
                                    v4 = "foo bar"
                                case "hello":
                                    var v27 int64 = 1
                                    _ = v27
                                case "\n\t":
                                    v4 = "hello"
                                default:
                                    var v28 uint8 = 89
                                    _ = v28
                                }
                                _ = fi15
                            }
                            if (2 > 2147483648) {
                                switch -6 {
                                case 697:
                                    v2 = append(v2, "hello")
                                case -662:
                                    v29 := GenS1{F0: v1, F1: -2147483648}
                                    _ = v29
                                case 4294967295:
                                    if false {
                                        v1 = (v1 + "foo bar")
                                    } else {
                                        for fi30 := 0; fi30 < 11; fi30++ {
                                            v1 = "foo bar"
                                            _ = fi30
                                        }
                                    }
                                }
                                v1 = v4
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = ((v4 + v1) + (v1 + v4))
                                    }
                                } else {
                                    for fi31 := 0; fi31 < 18; fi31++ {
                                        if (1 > 2) {
                                            v4 = "foo bar"
                                            switch v4 {
                                            case "a":
                                                v4 = ("" + v4)
                                            case "\n\t":
                                                switch 42 {
                                                case 2147483647:
                                                    var v32 int32 = 42
                                                    _ = v32
                                                }
                                            default:
                                                switch (((("\n\t" + v4) + v4) + v4) + v4) {
                                                case "":
                                                    _ = len(v2)
                                                case "\n\t":
                                                    v2 = append(v2, v4)
                                                case "a":
                                                    if (487 <= 42) {
                                                        var v33 uint16 = 255
                                                        v34 := GenS1{F0: "\n\t", F1: 4294967295}
                                                        v3 = append(v3, 0)
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                default:
                                                    var v35 uint32 = 765
                                                    _ = v35
                                                }
                                            }
                                            v1 = v1
                                        } else {
                                            var v36 uint32 = 1
                                            v1 = v1
                                            v37 := GenS2{F0: 255, F1: 1, F2: 42}
                                            _ = v36
                                            _ = v37
                                        }
                                        switch (fi31 << ((101) & 63)) {
                                        case 42:
                                            _ = len(v3)
                                        default:
                                            var v38 uint32 = 408
                                            v39 := GenS3{F0: v38}
                                            fmt.Println(fi10, fi31, v1)
                                            _ = v38
                                            _ = v39
                                        }
                                        _ = fi31
                                    }
                                    v40 := GenS4{F0: false, F1: 18446744073709551615, F2: 2}
                                    if true {
                                        v41 := GenS3{F0: 63}
                                        _ = len(v2)
                                        if (-280 != 0) {
                                            if (1e10 == 1e10) {
                                                if false {
                                                    for fi42 := 0; fi42 < 10; fi42++ {
                                                        if (42 == 778) {
                                                            var v43 int64 = 42
                                                            _ = v43
                                                        }
                                                        _ = fi42
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v44 string = (v1 + "a")
                                                    _ = v44
                                                } else {
                                                    fmt.Println(v1)
                                                    v45 := GenS4{F0: true, F1: 322, F2: 255}
                                                    _ = v45
                                                }
                                                var v46 string = v4
                                                v47 := GenS5{F0: -2147483648}
                                                _ = v46
                                                _ = v47
                                            } else {
                                                v48 := GenS6{F0: true, F1: v1}
                                                _ = v48
                                            }
                                            fmt.Println(v4)
                                            var v49 uint8 = 127
                                            _ = v49
                                        }
                                        _ = v41
                                    }
                                    _ = v40
                                }
                            } else {
                                fmt.Println(v1, v4, v1)
                                v1 = "\n\t"
                                var v50 uint32 = 433
                                _ = v50
                            }
                            _ = fi10
                        }
                        v4 = "foo bar"
                    } else {
                        for fi51 := 0; fi51 < 11; fi51++ {
                            var v52 bool = (!false)
                            _ = fi51
                            _ = v52
                        }
                    }
                } else {
                    var v53 int8 = -128
                    var v54 float32 = (float32(0.1) * 3.14159)
                    _ = v53
                    _ = v54
                }
            }
        }
        for fi55 := 0; fi55 < 9; fi55++ {
            if (true || false) {
                v4 = ("foo bar" + v1)
                switch (fi55 & fi55) {
                case 577:
                    var v56 int16 = 662
                    _ = v56
                case 896:
                    var v57 bool = (-2147483648 == 1)
                    _ = v57
                default:
                    for fi58 := 0; fi58 < 11; fi58++ {
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch v1 {
                            case "foo bar":
                                v59 := GenS7{F0: 42}
                                _ = v59
                            case "\n\t":
                                fmt.Println(v1, v1, fi55)
                            case "":
                                if (v1 >= (v1 + "hello")) {
                                    v4 = (v1 + "\n\t")
                                } else {
                                    var v60 uint = 750
                                    v60 = 18446744073709551615
                                    switch (v1 + v4) {
                                    case "hello":
                                        v4 = "\n\t"
                                    default:
                                        var v61 int8 = 1
                                        _ = v61
                                    }
                                    _ = v60
                                }
                            }
                        } else {
                            var v62 int32 = 2147483647
                            var v63 uint8 = 2
                            _ = v62
                            _ = v63
                        }
                        _ = fi58
                    }
                }
            }
            fmt.Println(v4, v4, fi55)
            var v64 string = "hello"
            _ = fi55
            _ = v64
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            if true {
                var v65 int16 = 42
                v66 := GenS8{F0: 128, F1: 0, F2: 0}
                if false {
                    for fi67 := 0; fi67 < 10; fi67++ {
                        if (2 == 18446744073709551615) {
                            if false {
                                var v68 int8 = -57
                                if false {
                                    _ = len(v2)
                                } else {
                                    var v69 string = v1
                                    if (975 > 0) {
                                        v3 = append(v3, 65535)
                                        var v70 int64 = -9223372036854775808
                                        for fi71 := 0; fi71 < 12; fi71++ {
                                            v70 = (v70 - (v70 | -1))
                                            _ = fi71
                                        }
                                        _ = v70
                                    } else {
                                        var v72 int = fi67
                                        _ = v72
                                    }
                                    _ = v69
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v68
                            }
                            v1 = (v1 + "")
                        }
                        _ = fi67
                    }
                    var v73 int16 = 1
                    v1 = ("foo bar" + v1)
                    _ = v73
                } else {
                    v74 := GenS9{F0: 642, F1: 9223372036854775807, F2: -1}
                    v4 = v1
                    for fi75 := 0; fi75 < 8; fi75++ {
                        var v76 bool = true
                        v77 := GenS10{F0: 102}
                        fmt.Println(fi75, v76, v65)
                        _ = fi75
                        _ = v76
                        _ = v77
                    }
                    _ = v74
                }
                _ = v65
                _ = v66
            } else {
                if (((false || false) || true) || true) {
                    if ("\n\t" != "") {
                        v78 := GenS11{F0: 0, F1: 0}
                        v1 = ("\n\t" + "a")
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                        _ = v78
                    } else {
                        for fi79 := 0; fi79 < 2; fi79++ {
                            v4 = v4
                            fmt.Println(fi79, v1)
                            _ = fi79
                        }
                        v1 = "a"
                    }
                    if true {
                        switch -841 {
                        case 0:
                            v80 := GenS12{F0: 0, F1: 32767, F2: -885}
                            _ = v80
                        default:
                            if (42 > 65535) {
                                var v81 int = 599
                                _ = v81
                            }
                        }
                        var v82 float64 = -2.25
                        fmt.Println(v4, v82)
                        _ = v82
                    }
                } else {
                    if (!false) {
                        var v83 uint64 = 473
                        _ = v83
                    } else {
                        v4 = ("foo bar" + v4)
                    }
                }
                for fi84 := 0; fi84 < 1; fi84++ {
                    for fi85 := 0; fi85 < 18; fi85++ {
                        _ = len(v3)
                        v86 := GenS7{F0: 1}
                        if false {
                            _ = len(v2)
                        } else {
                            v1 = v1
                            var v87 string = v4
                            if false {
                                v87 = "\n\t"
                                v88 := GenS7{F0: 42}
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                                _ = v88
                            }
                            _ = len(v3)
                            v4 = ""
                            _ = v87
                        }
                        _ = fi85
                        _ = v86
                    }
                    v89 := GenS13{F0: 0, F1: ("" != (v1 + v4))}
                    v90 := GenS14{F0: 1, F1: (fi84 - (-(fi84)))}
                    _ = fi84
                    _ = v89
                    _ = v90
                }
                if false {
                    v1 = "a"
                    if (0 > 0) {
                        if true {
                            if (-528 >= 2147483647) {
                                switch ("" + "a") {
                                case "foo bar":
                                    v4 = (v4 + v1)
                                case "a":
                                    _ = len(v3)
                                case "hello":
                                    for fi91 := 0; fi91 < 14; fi91++ {
                                        v1 = "\n\t"
                                        v92 := GenS15{F0: 507, F1: -9223372036854775808, F2: 65535}
                                        v93 := GenS8{F0: 127, F1: 18446744073709551615, F2: 708}
                                        _ = fi91
                                        _ = v92
                                        _ = v93
                                    }
                                }
                                v94 := GenS2{F0: 255, F1: 1, F2: 42}
                                _ = v94
                            }
                        } else {
                            fmt.Println(v1)
                            switch "hello" {
                            case "a":
                                if false {
                                    for fi95 := 0; fi95 < 12; fi95++ {
                                        fmt.Println(fi95, v4, v1)
                                        var v96 uint64 = 2
                                        v96 += 42
                                        _ = fi95
                                        _ = v96
                                    }
                                    _ = len(v3)
                                    v4 = v1
                                } else {
                                    if (42 < 0) {
                                        var v97 float32 = (0.0 - float32(1.5))
                                        var v98 int8 = 1
                                        fmt.Println(v98)
                                        _ = v97
                                        _ = v98
                                    } else {
                                        var v99 string = (v4 + "\n\t")
                                        switch v1 {
                                        case "hello":
                                            switch v1 {
                                            case "":
                                                var v100 uint64 = 42
                                                _ = v100
                                            default:
                                                if (false && (!true)) {
                                                    for fi101 := 0; fi101 < 2; fi101++ {
                                                        for fi102 := 0; fi102 < 10; fi102++ {
                                                            v103 := GenS16{F0: 42, F1: 1}
                                                            if (-539 >= 996) {
                                                                fmt.Println(v1, fi101)
                                                                var v104 int64 = -1
                                                                _ = v104
                                                            }
                                                            _ = fi102
                                                            _ = v103
                                                        }
                                                        for fi105 := 0; fi105 < 3; fi105++ {
                                                            var v106 float32 = (1.5 + 3.14159)
                                                            var v107 float64 = ((((0.0 / 1.5) + (float64((v106 - (1e10 * 1.5))) - (((((1.5 + -2.25) - (1.5 * 0.1)) - 1.5) - 0.0) * 0.0))) + (-2.25 * 0.1)) / 1e10)
                                                            _ = fi105
                                                            _ = v106
                                                            _ = v107
                                                        }
                                                        _ = fi101
                                                    }
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = 2
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            if len(v3) > 0 {
                                                v3[0] = 42
                                            }
                                        case "foo bar":
                                            var v108 int16 = 1000
                                            _ = v108
                                        }
                                        if ((129 <= -1) || false) {
                                            v109 := GenS17{F0: 1}
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v109
                                        } else {
                                            v99 = (("foo bar" + ((v1 + "\n\t") + "a")) + (("" + ((v99 + v4) + v4)) + ((v4 + v1) + (v4 + "hello"))))
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v4 = v99
                                        }
                                        _ = v99
                                    }
                                    v110 := GenS6{F0: true, F1: v4}
                                    _ = v110
                                }
                            case "foo bar":
                                fmt.Println(v4, v4)
                            }
                        }
                        if (!((-17 < 0) || false)) {
                            if (119 >= 255) {
                                if (-2147483648 != -2147483648) {
                                    switch 32768 {
                                    case 2147483647:
                                        if false {
                                            v4 = "hello"
                                            fmt.Println(v1)
                                        } else {
                                            v4 = v4
                                        }
                                    case 47:
                                        for fi111 := 0; fi111 < 15; fi111++ {
                                            v112 := GenS7{F0: 0}
                                            v113 := GenS18{F0: ""}
                                            _ = fi111
                                            _ = v112
                                            _ = v113
                                        }
                                    case 1:
                                        var v114 uint = 4294967295
                                        _ = v114
                                    }
                                    v4 = ("hello" + v4)
                                    var v115 int = 9223372036854775807
                                    _ = v115
                                }
                                for fi116 := 0; fi116 < 10; fi116++ {
                                    if (false && (1 < 9223372036854775807)) {
                                        for fi117 := 0; fi117 < 2; fi117++ {
                                            switch "hello" {
                                            case "a":
                                                _ = len(v3)
                                            case "":
                                                v4 = v4
                                            case "hello":
                                                if (false || false) {
                                                    var v118 float64 = (((1.5 + (0.1 / 1e10)) - -2.25) * 3.14159)
                                                    v118 -= -2.25
                                                    v118 = (0.1 - 0.0)
                                                    _ = v118
                                                } else {
                                                    if (("\n\t" + "") == (v4 + "foo bar")) {
                                                        switch "a" {
                                                        case "a":
                                                            var v119 uint16 = 0
                                                            _ = v119
                                                        case "":
                                                            v1 = ("a" + "\n\t")
                                                            var v120 int = (fi116 >> ((196) & 63))
                                                            _ = v120
                                                        }
                                                        fmt.Println(v4)
                                                        v121 := GenS9{F0: 4294967295, F1: -9223372036854775808, F2: 0}
                                                        _ = v121
                                                    } else {
                                                        switch v4 {
                                                        case "foo bar":
                                                            v122 := GenS13{F0: 4294967295, F1: true}
                                                            _ = v122
                                                        case "hello":
                                                            fmt.Println(v4)
                                                        case "\n\t":
                                                            v123 := GenS6{F0: (-1 < 1), F1: v4}
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            if len(v2) > 0 {
                                                                v2[0] = "\n\t"
                                                            }
                                                            _ = v123
                                                        default:
                                                            var v124 int64 = 0
                                                            var v125 bool = false
                                                            v1 = v4
                                                            _ = v124
                                                            _ = v125
                                                        }
                                                    }
                                                    var v126 float32 = 0.1
                                                    v4 = "hello"
                                                    _ = v126
                                                }
                                            }
                                            var v127 float64 = 1e10
                                            _ = fi117
                                            _ = v127
                                        }
                                        _ = len(v3)
                                    } else {
                                        _ = len(v2)
                                        if (!(!true)) {
                                            if (255 < 1) {
                                                fmt.Println(fi116, fi116, v1)
                                                var v128 uint16 = 65535
                                                _ = v128
                                            } else {
                                                v4 = "a"
                                                for fi129 := 0; fi129 < 19; fi129++ {
                                                    fmt.Println(v4)
                                                    var v130 bool = false
                                                    v4 = "\n\t"
                                                    _ = fi129
                                                    _ = v130
                                                }
                                            }
                                        } else {
                                            for fi131 := 0; fi131 < 14; fi131++ {
                                                fmt.Println(fi116)
                                                fmt.Println(fi116, fi131, v1)
                                                switch (v4 + ((v4 + "hello") + v1)) {
                                                case "\n\t":
                                                    switch 361 {
                                                    case 2147483647:
                                                        var v132 bool = (0 >= 38)
                                                        _ = v132
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                default:
                                                    v1 = v4
                                                }
                                                var v133 int16 = 1
                                                _ = fi131
                                                _ = v133
                                            }
                                            v134 := GenS15{F0: 0, F1: 94, F2: 324}
                                            _ = v134
                                        }
                                    }
                                    _ = fi116
                                }
                            } else {
                                for fi135 := 0; fi135 < 3; fi135++ {
                                    var v136 string = (v1 + v4)
                                    var v137 int8 = 42
                                    _ = fi135
                                    _ = v136
                                    _ = v137
                                }
                                var v138 uint8 = 251
                                _ = v138
                            }
                            v139 := GenS19{F0: -1}
                            _ = v139
                        } else {
                            var v140 int16 = -32768
                            var v141 int16 = v140
                            v140 += -32768
                            _ = v140
                            _ = v141
                        }
                    } else {
                        v4 = v1
                    }
                    switch ((((v1 + "\n\t") + v4) + (v4 + "foo bar")) + (v1 + ("\n\t" + ""))) {
                    case "\n\t":
                        v4 = v1
                    case "hello":
                        switch v4 {
                        case "\n\t":
                            if false {
                                switch 42 {
                                case -353:
                                    for fi142 := 0; fi142 < 20; fi142++ {
                                        var v143 int16 = 1
                                        _ = fi142
                                        _ = v143
                                    }
                                case 4294967295:
                                    if (2 == 368) {
                                        var v144 float64 = 1e10
                                        _ = v144
                                    } else {
                                        v4 = (v4 + "a")
                                    }
                                }
                                var v145 uint8 = 255
                                _ = v145
                            } else {
                                v4 = v1
                            }
                        }
                    }
                } else {
                    var v146 float32 = 0.0
                    _ = v146
                }
            }
        } else {
            v147 := GenS20{F0: 39, F1: 0}
            v4 = ((v4 + v4) + "foo bar")
            _ = v147
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if true {
        switch v4 {
        case "a":
            v4 = ""
        case "\n\t":
            fmt.Println(v1, v4)
        case "":
            if (true || (!false)) {
                if false {
                    if false {
                        fmt.Println(v4)
                        if len(v2) > 0 {
                            v2[0] = (v4 + v1)
                        }
                        switch 32767 {
                        case 42:
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                        case 32768:
                            switch 934 {
                            case 9223372036854775807:
                                if true {
                                    var v148 int32 = -2147483648
                                    if false {
                                        if len(v2) > 0 {
                                            v2[0] = (v1 + "a")
                                        }
                                    }
                                    _ = v148
                                }
                            case 0:
                                if len(v2) > 0 {
                                    v2[0] = ((v4 + (v4 + ("" + v1))) + v4)
                                }
                            case -1:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                v4 = ""
                            }
                        case 1:
                            for fi149 := 0; fi149 < 19; fi149++ {
                                v1 = v1
                                _ = fi149
                            }
                        default:
                            v1 = v4
                        }
                    } else {
                        for fi150 := 0; fi150 < 6; fi150++ {
                            v3 = append(v3, 110)
                            _ = fi150
                        }
                        if (255 == 67) {
                            if (591 != 844) {
                                v151 := GenS21{F0: 255, F1: 1, F2: 9223372036854775807}
                                _ = v151
                            }
                            v4 = "foo bar"
                            for fi152 := 0; fi152 < 10; fi152++ {
                                var v153 float64 = (0.0 - (-2.25 / 1.5))
                                var v154 bool = true
                                for fi155 := 0; fi155 < 16; fi155++ {
                                    if (2 >= 904) {
                                        v153 -= v153
                                        v4 = v4
                                        switch -2147483648 {
                                        case 4294967295:
                                            fmt.Println(v154, v154, v4)
                                        case 1:
                                            v154 = false
                                        case 567:
                                            v156 := GenS14{F0: -477, F1: (-(fi155))}
                                            _ = v156
                                        }
                                    } else {
                                        if true {
                                            v154 = v154
                                            v1 = "foo bar"
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    }
                                    v157 := GenS12{F0: -1, F1: -419, F2: fi152}
                                    _ = fi155
                                    _ = v157
                                }
                                _ = fi152
                                _ = v153
                                _ = v154
                            }
                        }
                        v4 = v1
                    }
                    v158 := GenS3{F0: 181}
                    _ = v158
                }
                var v159 int64 = 1
                switch "foo bar" {
                case "\n\t":
                    v160 := GenS20{F0: 127, F1: 32767}
                    _ = v160
                }
                _ = v159
            }
        default:
            for fi161 := 0; fi161 < 6; fi161++ {
                v162 := GenS14{F0: 242, F1: (fi161 >> ((425) & 63))}
                _ = fi161
                _ = v162
            }
        }
    }
    var v163 string = v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v164 int8 = -5
    for fi165 := 0; fi165 < 20; fi165++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v166 uint16 = 42
        switch (fi165 | fi165) {
        case 42:
            switch (("hello" + v4) + ("hello" + (v1 + "a"))) {
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "":
                for fi167 := 0; fi167 < 13; fi167++ {
                    v164 = (v164 - (v164 >> ((42) & 63)))
                    _ = fi167
                }
            default:
                v168 := GenS9{F0: 357, F1: -9223372036854775808, F2: 32767}
                _ = v168
            }
        case -349:
            v3 = append(v3, 2147483648)
        }
        _ = fi165
        _ = v166
    }
    for fi169 := 0; fi169 < 4; fi169++ {
        for fi170 := 0; fi170 < 11; fi170++ {
            v1 = v163
            _ = len(v3)
            var v171 uint64 = 42
            if ((!(!((65535 <= 18446744073709551615) || false))) && (48 > 255)) {
                var v172 string = "foo bar"
                _ = v172
            } else {
                var v173 uint32 = 210
                if ((fi170 << ((462) & 63)) < (fi170 & fi169)) {
                    var v174 float32 = (0.1 / (-2.25 - 0.0))
                    var v175 string = ((v1 + v163) + v163)
                    _ = v174
                    _ = v175
                } else {
                    v173 += (^(v173))
                    switch v1 {
                    case "":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "\n\t":
                        var v176 bool = (!(1e10 == ((float32((1.5 * (0.1 * 3.14159))) * float32(1.5)) - 1e10)))
                        _ = v176
                    case "hello":
                        switch (v1 + "\n\t") {
                        case "\n\t":
                            var v177 int64 = 9223372036854775807
                            _ = v177
                        case "a":
                            var v178 float32 = 3.14159
                            _ = v178
                        }
                    default:
                        switch fi170 {
                        case -398:
                            switch v1 {
                            case "foo bar":
                                var v179 int = 32767
                                _ = v179
                            }
                        case -1:
                            var v180 bool = false
                            _ = v180
                        case 32767:
                            var v181 int32 = 2147483647
                            _ = v181
                        }
                    }
                    fmt.Println(v164, v163)
                }
                v163 = v1
                if (132 != 127) {
                    v182 := GenS22{F0: v171, F1: 252}
                    if (0 >= 127) {
                        fmt.Println(fi170, v164, v1)
                        _ = len(v2)
                        switch 32767 {
                        case 32768:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case -131:
                            v183 := GenS20{F0: v164, F1: 129}
                            _ = v183
                        case -1:
                            fmt.Println(fi170, v4)
                        }
                    }
                    v184 := GenS4{F0: true, F1: 65535, F2: 255}
                    _ = v182
                    _ = v184
                }
                _ = v173
            }
            var v185 int16 = -721
            _ = fi170
            _ = v171
            _ = v185
        }
        var v186 bool = (1000 == 1000)
        _ = fi169
        _ = v186
    }
    for fi187 := 0; fi187 < 15; fi187++ {
        var v188 float32 = 1.5
        v188 += 0.0
        fmt.Println(v188)
        _ = fi187
        _ = v188
    }
    fmt.Println("v1:", v1)
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v163
    _ = v164
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
