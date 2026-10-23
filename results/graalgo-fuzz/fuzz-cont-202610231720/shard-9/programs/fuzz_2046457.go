package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
type GenS3 struct { F0 uint64; F1 uint16; F2 int }
type GenS4 struct { F0 int8; F1 uint32; F2 bool }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 string; F1 int8; F2 uint32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int; F1 int8; F2 int }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 string; F1 uint8 }
type GenS14 struct { F0 int64; F1 int }
type GenS15 struct { F0 uint16; F1 int8 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 bool; F1 int32 }
type GenS18 struct { F0 uint64; F1 uint; F2 int64 }
type GenS19 struct { F0 string; F1 uint64 }
type GenS20 struct { F0 int }
type GenS21 struct { F0 int; F1 int16 }
type GenS22 struct { F0 int8; F1 int; F2 uint }
type GenS23 struct { F0 bool; F1 int32; F2 string }
type GenS24 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS25 struct { F0 int64; F1 uint32; F2 int32 }
type GenS26 struct { F0 int; F1 int64; F2 bool }
type GenS27 struct { F0 int16; F1 uint16; F2 uint8 }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]uint32)
    var v3 float32 = 3.14159
    var v4 int16 = 32767
    v5 := []int{-535}
    var v6 int32 = 342
    for fi7 := 0; fi7 < 13; fi7++ {
        v4 += (^(v4))
        if true {
            if false {
                if false {
                    for fi8 := 0; fi8 < 14; fi8++ {
                        switch "" {
                        case "":
                            fmt.Println(v6)
                        case "a":
                            var v9 string = "foo bar"
                            _ = v9
                        case "foo bar":
                            v3 = v1
                        }
                        _ = fi8
                    }
                }
                if len(v5) > 0 {
                    v5[0] = (^(fi7))
                }
            } else {
                for fi10 := 0; fi10 < 15; fi10++ {
                    if (0 <= 128) {
                        v11 := GenS1{F0: (358 < 0), F1: 379, F2: (^(v4))}
                        v3 -= v1
                        if (1e10 > v1) {
                            v12 := GenS2{F0: 2, F1: 635, F2: "\n\t"}
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v12
                        } else {
                            fmt.Println(v6, v3, v3)
                            switch "foo bar" {
                            case "hello":
                                var v13 int8 = 127
                                _ = v13
                            }
                            switch fi10 {
                            case 198:
                                var v14 uint16 = 42
                                _ = v14
                            case 9223372036854775807:
                                if (!((42 <= 25) && false)) {
                                    var v15 uint16 = 727
                                    if true {
                                        v4 += (v4 / ((1) | 1))
                                        v16 := GenS2{F0: 255, F1: v15, F2: ""}
                                        if true {
                                            var v17 int8 = -10
                                            _ = v17
                                        }
                                        v6 -= v6
                                        v3 = float32(-2.25)
                                        _ = v16
                                    } else {
                                        if true {
                                            if ((3.14159 * (float32((1e10 * 1e10)) - ((v3 + float32(0.1)) - float32(3.14159)))) == (v1 * v1)) {
                                                if ((^(v4)) == -32768) {
                                                    var v18 uint8 = 1
                                                    _ = v18
                                                }
                                                var v19 int = int(v4)
                                                _ = v19
                                            }
                                            _ = len(v5)
                                        }
                                        v20 := v2[32768]
                                        switch int(v4) {
                                        case -661:
                                            if true {
                                                v21 := GenS3{F0: 0, F1: v15, F2: (fi7 ^ fi10)}
                                                _ = v21
                                            }
                                        case 4294967295:
                                            fmt.Println(v6, v15, v15)
                                        }
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v22 float64 = 1.5
                                            _ = v22
                                        case "a":
                                            var v23 bool = false
                                            _ = v23
                                        default:
                                            var v24 uint32 = v20
                                            _ = v24
                                        }
                                        _ = v20
                                    }
                                    if false {
                                        v25 := GenS4{F0: 127, F1: 2147483647, F2: (409 >= 2)}
                                        v3 = 0.1
                                        _ = v25
                                    } else {
                                        if (true && (!("hello" <= "foo bar"))) {
                                            v15 += (v15 / ((v15) | 1))
                                            v4 -= (v4 - (v4 % ((v4) | 1)))
                                            v26 := GenS2{F0: 105, F1: 1, F2: ""}
                                            _ = v26
                                        }
                                        if true {
                                            v27 := GenS2{F0: 42, F1: v15, F2: "\n\t"}
                                            var v28 float32 = (-2.25 / 1.5)
                                            fmt.Println(v28)
                                            v15 = 0
                                            _ = v27
                                            _ = v28
                                        } else {
                                            delete(v2, (fi7 & fi10))
                                            var v29 uint16 = (v15 >> ((18446744073709551615) & 63))
                                            _ = v29
                                        }
                                        delete(v2, -2147483648)
                                    }
                                    _ = v15
                                } else {
                                    var v30 bool = ((!(false || false)) && ((!true) && false))
                                    switch ("a" + "a") {
                                    case "a":
                                        v30 = ((fi7 + 4294967295) >= (fi10 + fi7))
                                    case "":
                                        if true {
                                            if v30 {
                                                if v30 {
                                                    var v31 float32 = v3
                                                    _ = v31
                                                } else {
                                                    v30 = (3.14159 <= 0.0)
                                                    v32 := GenS5{F0: -102, F1: false}
                                                    v6 -= (v6 | v6)
                                                    _ = v32
                                                }
                                                v33 := GenS6{F0: (v6 % ((2147483647) | 1)), F1: 128}
                                                delete(v2, fi10)
                                                if (2147483647 < 65535) {
                                                    if len(v5) > 0 {
                                                        v5[0] = fi10
                                                    }
                                                }
                                                _ = v33
                                            }
                                            switch -9223372036854775808 {
                                            case 127:
                                                if v30 {
                                                    var v34 uint64 = 0
                                                    fmt.Println(v6, v34, fi7)
                                                    var v35 int8 = 127
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    var v36 int = -2147483648
                                                    v37 := GenS4{F0: 0, F1: 2, F2: false}
                                                    _ = v36
                                                    _ = v37
                                                }
                                            case 32767:
                                                _ = len(v2)
                                            case -9223372036854775808:
                                                v38 := GenS7{F0: 698, F1: (v4 * (-(v4)))}
                                                _ = v38
                                            default:
                                                v39 := GenS8{F0: 0, F1: 65535}
                                                _ = v39
                                            }
                                            fmt.Println(v30, v1)
                                        }
                                    case "\n\t":
                                        var v40 uint = 254
                                        _ = v40
                                    }
                                    v6 = (-(v6))
                                    _ = v30
                                }
                            }
                        }
                        v4 += 1000
                        if true {
                            var v41 bool = true
                            switch ("foo bar" + "\n\t") {
                            case "hello":
                                var v42 bool = (147 != 4294967295)
                                _ = v42
                            }
                            _ = v41
                        } else {
                            if false {
                                v43 := GenS9{F0: "\n\t", F1: 1, F2: 1}
                                var v44 uint = 255
                                _ = v43
                                _ = v44
                            } else {
                                switch fi7 {
                                case -237:
                                    if false {
                                        if (((1e10 * ((0.1 * ((1.5 * (0.1 + 1e10)) + 3.14159)) + 3.14159)) / 1.5) == 1e10) {
                                            if (0 < 2147483648) {
                                                if (((255 >= 15) || (!(false || true))) || (65535 > 4294967295)) {
                                                    v3 += 0.1
                                                    v6 = v6
                                                }
                                                v1 -= 1e10
                                                if false {
                                                    var v45 int64 = 736
                                                    var v46 float32 = 1e10
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    if (-289 < 0) {
                                                        fmt.Println(v3, v3, v4)
                                                        if true {
                                                            v4 = v4
                                                        }
                                                    }
                                                    if (!true) {
                                                        if true {
                                                            v1 -= (3.14159 / v1)
                                                            if true {
                                                                var v47 int8 = -1
                                                                if false {
                                                                    if len(v5) > 0 {
                                                                        v5[0] = 42
                                                                    }
                                                                }
                                                                _ = v47
                                                            }
                                                            var v48 string = ("hello" + "foo bar")
                                                            _ = v48
                                                        }
                                                        v4 = 525
                                                        switch (fi7 / ((127) | 1)) {
                                                        case 2147483647:
                                                            v3 -= v3
                                                        case -837:
                                                            var v49 float64 = 1e10
                                                            _ = v49
                                                        default:
                                                            v50 := GenS5{F0: v6, F1: false}
                                                            _ = v50
                                                        }
                                                    }
                                                }
                                            }
                                            if false {
                                                fmt.Println(v6, v3, v3)
                                                switch (fi10 << ((0) & 63)) {
                                                case 63:
                                                    v6 = -447
                                                case 4294967295:
                                                    v6 = 711
                                                case 42:
                                                    if false {
                                                        var v51 int64 = -9223372036854775808
                                                        switch ("a" + "a") {
                                                        case "hello":
                                                            if true {
                                                                if (false || false) {
                                                                    if false {
                                                                        var v52 int64 = v51
                                                                        var v53 int64 = v52
                                                                        v54 := GenS10{F0: 65535}
                                                                        _ = v52
                                                                        _ = v53
                                                                        _ = v54
                                                                    } else {
                                                                        delete(v2, fi10)
                                                                        switch (fi7 ^ (-(fi7))) {
                                                                        case -1:
                                                                            if false {
                                                                                fmt.Println(v4, v4)
                                                                                delete(v2, (fi10 & (fi7 | int(v51))))
                                                                            } else {
                                                                                fmt.Println(fi7, v4, fi10)
                                                                                if (42 == 2) {
                                                                                    _ = len(v5)
                                                                                }
                                                                            }
                                                                        case -145:
                                                                            var v55 int32 = (v6 + -672)
                                                                            _ = v55
                                                                        case 42:
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        }
                                                                        var v56 uint = 628
                                                                        _ = v56
                                                                    }
                                                                    var v57 uint16 = 1
                                                                    _ = v57
                                                                }
                                                            }
                                                        default:
                                                            v3 += (0.0 + float32((float64(float32(0.0)) - 3.14159)))
                                                            fmt.Println(fi7, fi10, v1)
                                                            switch fi10 {
                                                            case 85:
                                                                fmt.Println(fi7, fi7, v1)
                                                            case 42:
                                                                if ((255 > 621) || true) {
                                                                    if false {
                                                                        v58 := GenS1{F0: true, F1: 65535, F2: (v4 + (v4 | 1000))}
                                                                        if false {
                                                                            switch ((("a" + ("" + ((("hello" + ("a" + "\n\t")) + "") + "\n\t"))) + (("hello" + "a") + "hello")) + "foo bar") {
                                                                            case "\n\t":
                                                                                fmt.Println(fi7, v4, v3)
                                                                            case "foo bar":
                                                                                if false {
                                                                                    var v59 uint8 = 2
                                                                                    fmt.Println(v4, fi7)
                                                                                    _ = v59
                                                                                }
                                                                            case "hello":
                                                                                var v60 uint8 = 95
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    if ((fi7 % (((fi10 ^ (fi10 + (^(fi7))))) | 1)) == -2147483648) {
                                                                                        if false {
                                                                                            v6 = v6
                                                                                        } else {
                                                                                            v60 += 2
                                                                                            if true {
                                                                                                v6 += 153
                                                                                                var v61 uint32 = 501
                                                                                                if (1.5 >= float32((0.0 / ((float64(v3) + ((1.5 / 0.1) - float64(1e10))) / float64(((float32(-2.25) / 0.1) * -2.25)))))) {
                                                                                                    fmt.Println(v61, v4, v61)
                                                                                                    v62 := GenS9{F0: "hello", F1: 127, F2: v61}
                                                                                                    var v63 bool = true
                                                                                                    _ = v62
                                                                                                    _ = v63
                                                                                                }
                                                                                                _ = v61
                                                                                            } else {
                                                                                                if false {
                                                                                                    var v64 string = "foo bar"
                                                                                                    v3 = 0.0
                                                                                                    _ = v64
                                                                                                }
                                                                                                v65 := GenS7{F0: 1, F1: v4}
                                                                                                _ = v65
                                                                                            }
                                                                                        }
                                                                                        switch (fi10 + 127) {
                                                                                        case 32767:
                                                                                            v66 := GenS11{F0: (fi10 ^ 9223372036854775807), F1: 48, F2: fi7}
                                                                                            _ = v66
                                                                                        }
                                                                                        if (v60 != (v60 << ((899) & 63))) {
                                                                                            if false {
                                                                                                if true {
                                                                                                    var v67 bool = false
                                                                                                    _ = v67
                                                                                                }
                                                                                                switch (^(fi10)) {
                                                                                                case 42:
                                                                                                    switch "\n\t" {
                                                                                                    case "a":
                                                                                                        v6 = (v6 + -322)
                                                                                                    }
                                                                                                case -9223372036854775808:
                                                                                                    if (1e10 > float64(v1)) {
                                                                                                        v68 := GenS12{F0: 2147483647}
                                                                                                        _ = v68
                                                                                                    }
                                                                                                case -2147483648:
                                                                                                    if ((v51 << ((1) & 63)) > int64(fi10)) {
                                                                                                        v69 := v2[fi10]
                                                                                                        _ = v69
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            var v70 uint16 = 765
                                                                                            fmt.Println(v4, v6, v51)
                                                                                            var v71 float32 = v3
                                                                                            switch (-(fi7)) {
                                                                                            case 4294967295:
                                                                                                v71 = v3
                                                                                            }
                                                                                            _ = v70
                                                                                            _ = v71
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    switch fi7 {
                                                                                    case -9223372036854775808:
                                                                                        switch ("hello" + ("hello" + "\n\t")) {
                                                                                        case "a":
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                v72 := GenS7{F0: 32767, F1: (v4 * (v4 >> ((uint(v60)) & 63)))}
                                                                                                _ = v72
                                                                                            case "hello":
                                                                                                if false {
                                                                                                    v73 := v2[fi10]
                                                                                                    _ = v73
                                                                                                } else {
                                                                                                    var v74 uint16 = 65535
                                                                                                    _ = v74
                                                                                                }
                                                                                            default:
                                                                                                var v75 float64 = 1e10
                                                                                                _ = v75
                                                                                            }
                                                                                        case "hello":
                                                                                            v5 = append(v5, fi7)
                                                                                        default:
                                                                                            var v76 int64 = v51
                                                                                            _ = v76
                                                                                        }
                                                                                    case 128:
                                                                                        _ = len(v2)
                                                                                    case 0:
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = (fi7 & 32767)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v60 -= (^(v60))
                                                                                _ = v60
                                                                            default:
                                                                                var v77 uint16 = 598
                                                                                _ = v77
                                                                            }
                                                                            if ((!true) && false) {
                                                                                switch "hello" {
                                                                                case "hello":
                                                                                    v51 = v51
                                                                                case "foo bar":
                                                                                    _ = len(v2)
                                                                                default:
                                                                                    v78 := GenS8{F0: 2, F1: 127}
                                                                                    _ = v78
                                                                                }
                                                                                _ = len(v2)
                                                                                v51 += 796
                                                                            }
                                                                        } else {
                                                                            var v79 uint16 = 255
                                                                            var v80 int8 = 127
                                                                            fmt.Println(v6, v3)
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                        _ = v58
                                                                    }
                                                                    v6 = (v6 * (v6 << ((255) & 63)))
                                                                    if (v6 >= v6) {
                                                                        if (3.14159 <= (3.14159 + 0.1)) {
                                                                            if true {
                                                                                v51 += v51
                                                                            } else {
                                                                                var v81 int64 = (^(v51))
                                                                                _ = v81
                                                                            }
                                                                            if (!false) {
                                                                                var v82 uint16 = 42
                                                                                switch fi10 {
                                                                                case 32767:
                                                                                    fmt.Println(v6, v6)
                                                                                    if ((v1 < ((0.1 - v1) + (float32(1e10) - 1.5))) && false) {
                                                                                        v83 := GenS13{F0: "foo bar", F1: 1}
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = (fi7 / ((fi7) | 1))
                                                                                        }
                                                                                        _ = v83
                                                                                    }
                                                                                    v84 := GenS14{F0: (v51 + int64(v4)), F1: 160}
                                                                                    _ = v84
                                                                                default:
                                                                                    fmt.Println(v4)
                                                                                }
                                                                                _ = v82
                                                                            } else {
                                                                                var v85 uint = 4294967295
                                                                                v86 := GenS15{F0: 65535, F1: 55}
                                                                                _ = len(v2)
                                                                                var v87 int8 = -128
                                                                                _ = v85
                                                                                _ = v86
                                                                                _ = v87
                                                                            }
                                                                            var v88 uint64 = 42
                                                                            _ = v88
                                                                        }
                                                                        if true {
                                                                            var v89 uint = 500
                                                                            v90 := GenS15{F0: 32767, F1: 1}
                                                                            if true {
                                                                                var v91 string = "foo bar"
                                                                                if true {
                                                                                    v1 = (float32(3.14159) + (1e10 - v1))
                                                                                    v92 := GenS16{F0: (v89 % ((0) | 1))}
                                                                                    _ = v92
                                                                                } else {
                                                                                    if (!(42 == 2)) {
                                                                                        if (!(false && (0 < 100))) {
                                                                                            v3 += v3
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v4)
                                                                                        switch ("" + "") {
                                                                                        case "\n\t":
                                                                                            v93 := GenS12{F0: (v6 << (((v89 >> (((v89 ^ 853)) & 63))) & 63))}
                                                                                            _ = v93
                                                                                        }
                                                                                        var v94 uint64 = 661
                                                                                        _ = v94
                                                                                    }
                                                                                    var v95 float32 = (float32((float64(v3) - 3.14159)) + float32(3.14159))
                                                                                    _ = v95
                                                                                }
                                                                                if false {
                                                                                    fmt.Println(v4, v91, v6)
                                                                                }
                                                                                _ = v91
                                                                            }
                                                                            _ = v89
                                                                            _ = v90
                                                                        } else {
                                                                            if (!(622 != 892)) {
                                                                                if false {
                                                                                    var v96 bool = (1 != 2147483647)
                                                                                    v3 -= v1
                                                                                    if v96 {
                                                                                        var v97 uint = 10
                                                                                        if v96 {
                                                                                            var v98 int64 = (v51 / ((v51) | 1))
                                                                                            if v96 {
                                                                                                if (v96 && v96) {
                                                                                                    var v99 int64 = v51
                                                                                                    _ = v99
                                                                                                }
                                                                                            } else {
                                                                                                v100 := GenS10{F0: 255}
                                                                                                var v101 int32 = -209
                                                                                                var v102 int16 = (v4 ^ v4)
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                                _ = v102
                                                                                            }
                                                                                            _ = v98
                                                                                        }
                                                                                        var v103 uint = v97
                                                                                        _ = v103
                                                                                        _ = v97
                                                                                    } else {
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                        fmt.Println(v1, v3)
                                                                                    }
                                                                                    _ = v96
                                                                                }
                                                                                v3 -= v1
                                                                                var v104 uint16 = 255
                                                                                _ = v104
                                                                            }
                                                                            if ((-2.25 / -2.25) <= -2.25) {
                                                                                v105 := GenS9{F0: "hello", F1: 100, F2: 2147483647}
                                                                                v3 -= 0.1
                                                                                if (!true) {
                                                                                    v106 := v2[32768]
                                                                                    _ = v106
                                                                                } else {
                                                                                    v1 += -2.25
                                                                                    fmt.Println(fi7, fi10)
                                                                                    v107 := GenS17{F0: false, F1: int32(v51)}
                                                                                    _ = v107
                                                                                }
                                                                                _ = v105
                                                                            }
                                                                            v5 = append(v5, (fi10 * 0))
                                                                        }
                                                                        v108 := GenS11{F0: (^(fi7)), F1: -21, F2: (fi7 + 179)}
                                                                        _ = v108
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v51
                                                    } else {
                                                        if true {
                                                            if false {
                                                                v1 += 1.5
                                                            } else {
                                                                fmt.Println(fi7)
                                                                if (3.14159 < 3.14159) {
                                                                    v109 := GenS16{F0: 0}
                                                                    v2[fi7] = 2147483648
                                                                    v110 := GenS4{F0: -1, F1: 776, F2: false}
                                                                    v3 = v3
                                                                    _ = v109
                                                                    _ = v110
                                                                }
                                                                v4 = (v4 % ((v4) | 1))
                                                                if false {
                                                                    var v111 uint32 = 4294967295
                                                                    if ((^(fi7)) <= 127) {
                                                                        var v112 int32 = -583
                                                                        _ = v112
                                                                    } else {
                                                                        fmt.Println(fi10)
                                                                        var v113 int64 = 57
                                                                        v113 = v113
                                                                        _ = v113
                                                                    }
                                                                    _ = v111
                                                                } else {
                                                                    v5 = append(v5, (fi10 * (fi10 ^ -9223372036854775808)))
                                                                    if (v6 < (v6 | v6)) {
                                                                        var v114 uint32 = 498
                                                                        if false {
                                                                            v115 := GenS5{F0: v6, F1: (1 < 142)}
                                                                            _ = v115
                                                                        }
                                                                        v3 = (v3 * (v3 * v3))
                                                                        _ = v114
                                                                    }
                                                                    v3 = 1.5
                                                                }
                                                            }
                                                            v6 += (v6 * v6)
                                                            fmt.Println(v4, fi10, fi10)
                                                            if len(v5) > 0 {
                                                                v5[0] = (fi10 << ((2) & 63))
                                                            }
                                                            v116 := GenS7{F0: 32767, F1: 236}
                                                            _ = v116
                                                        }
                                                        var v117 uint8 = 128
                                                        var v118 int64 = -181
                                                        v118 -= 378
                                                        _ = v117
                                                        _ = v118
                                                    }
                                                }
                                            }
                                            var v119 uint64 = 506
                                            _ = v119
                                        }
                                    } else {
                                        if ((4294967295 == 1) && false) {
                                            var v120 int32 = (^(v6))
                                            if (true || false) {
                                                v4 = (v4 | v4)
                                                v121 := GenS18{F0: 18446744073709551615, F1: 18446744073709551615, F2: 0}
                                                switch "a" {
                                                case "\n\t":
                                                    v3 = ((((0.1 * (v3 - v3)) / ((1.5 + (v3 + (v3 - (1e10 + v3)))) - v1)) / (3.14159 / (((v1 * float32(-2.25)) * float32((-2.25 * 1.5))) + v3))) - v1)
                                                case "":
                                                    v120 = -602
                                                }
                                                _ = v121
                                            } else {
                                                var v122 int32 = (v120 - v6)
                                                if (255 == 36) {
                                                    v123 := GenS19{F0: ("foo bar" + "\n\t"), F1: 489}
                                                    v1 -= (float32(-2.25) - -2.25)
                                                    var v124 uint = 1
                                                    var v125 float64 = ((0.0 / 1.5) / float64(float32(1.5)))
                                                    _ = v123
                                                    _ = v124
                                                    _ = v125
                                                }
                                                _ = v122
                                            }
                                            _ = v120
                                        }
                                    }
                                case -308:
                                    switch -9223372036854775808 {
                                    case 725:
                                        fmt.Println(v3, fi7)
                                    case 0:
                                        var v126 uint = 460
                                        _ = v126
                                    }
                                case 128:
                                    v1 = v3
                                }
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v127 := GenS13{F0: "", F1: 1}
                            _ = v127
                        }
                        _ = v11
                    } else {
                        v6 = (v6 / (((v6 >> ((508) & 63))) | 1))
                        fmt.Println(v6)
                    }
                    var v128 int32 = (v6 | (^(v6)))
                    var v129 int16 = v4
                    fmt.Println(v128)
                    v1 += v3
                    _ = fi10
                    _ = v128
                    _ = v129
                }
            }
            if (!true) {
                var v130 uint16 = 32767
                switch ("a" + "a") {
                case "foo bar":
                    if ("foo bar" > "foo bar") {
                        for fi131 := 0; fi131 < 11; fi131++ {
                            var v132 bool = ((0.1 != 3.14159) || (!true))
                            v6 += v6
                            v130 = (v130 % (((v130 | (v130 + v130))) | 1))
                            _ = fi131
                            _ = v132
                        }
                        if (80 == 255) {
                            v3 -= v3
                            v133 := GenS3{F0: 4294967295, F1: (v130 >> ((42) & 63)), F2: 32768}
                            switch fi7 {
                            case 0:
                                v130 = (v130 - 2)
                            case 4294967295:
                                v134 := GenS20{F0: fi7}
                                _ = v134
                            default:
                                for fi135 := 0; fi135 < 20; fi135++ {
                                    if false {
                                        if false {
                                            v136 := GenS16{F0: 42}
                                            _ = v136
                                        }
                                        switch fi7 {
                                        case 352:
                                            v6 = v6
                                        }
                                        switch ("a" + "hello") {
                                        case "hello":
                                            v130 = v130
                                        }
                                        var v137 uint16 = (v130 & v130)
                                        fmt.Println(v130, v3, fi7)
                                        if (3.14159 == (float32(-2.25) * ((float32(float64((0.0 / v1))) - v1) - (-2.25 / (((1e10 - v3) + 3.14159) + 3.14159))))) {
                                            var v138 uint64 = 952
                                            if ("foo bar" > ("a" + "hello")) {
                                                switch -428 {
                                                case 128:
                                                    v137 = (v130 | 1)
                                                case -1:
                                                    switch "hello" {
                                                    case "a":
                                                        var v139 uint64 = v138
                                                        _ = v139
                                                    case "\n\t":
                                                        _ = len(v2)
                                                    case "":
                                                        v140 := GenS21{F0: (fi7 % ((fi135) | 1)), F1: (v4 % ((-32768) | 1))}
                                                        _ = v140
                                                    }
                                                }
                                                v1 -= (v3 * (v1 - v3))
                                            } else {
                                                v138 += (v138 * 42)
                                                fmt.Println(v3, fi135)
                                                switch "hello" {
                                                case "\n\t":
                                                    switch (-(fi7)) {
                                                    case 9223372036854775807:
                                                        if (1 == 9223372036854775807) {
                                                            v141 := GenS17{F0: true, F1: (^(v6))}
                                                            v6 += v6
                                                            _ = v141
                                                        }
                                                        var v142 int64 = -1
                                                        _ = v142
                                                    }
                                                case "a":
                                                    v130 -= 923
                                                case "foo bar":
                                                    if (!false) {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    } else {
                                                        v130 += 294
                                                    }
                                                }
                                            }
                                            var v143 uint8 = 136
                                            v138 += v138
                                            _ = v138
                                            _ = v143
                                        }
                                        _ = v137
                                    }
                                    if (!(((!(128 < 42)) || (185 != 0)) && (!(true && (!false))))) {
                                        v144 := GenS22{F0: -1, F1: 885, F2: 382}
                                        v145 := GenS23{F0: (0 < 18446744073709551615), F1: (v6 ^ v6), F2: "\n\t"}
                                        _ = v144
                                        _ = v145
                                    } else {
                                        v146 := v2[(fi7 % ((fi7) | 1))]
                                        var v147 int16 = (v4 % ((int16(v146)) | 1))
                                        fmt.Println(v6)
                                        _ = v146
                                        _ = v147
                                    }
                                    _ = fi135
                                }
                            }
                            _ = v133
                        }
                        var v148 int8 = 26
                        _ = v148
                    }
                case "hello":
                    var v149 int = (fi7 * fi7)
                    _ = v149
                }
                _ = v130
            }
        } else {
            delete(v2, fi7)
            v5 = append(v5, (^(fi7)))
        }
        _ = fi7
    }
    for fi150 := 0; fi150 < 16; fi150++ {
        var v151 uint32 = 0
        _ = fi150
        _ = v151
    }
    switch "\n\t" {
    case "":
        for fi152 := 0; fi152 < 8; fi152++ {
            var v153 int32 = v6
            var v154 int8 = 33
            _ = fi152
            _ = v153
            _ = v154
        }
    case "\n\t":
        v155 := GenS22{F0: -128, F1: 4294967295, F2: 1}
        _ = v155
    }
    if true {
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            if (v3 < float32(float64(v1))) {
                fmt.Println(v1, v6, v1)
            }
        case "":
            var v156 string = (("hello" + "\n\t") + "\n\t")
            _ = v156
        case "hello":
            var v157 int = 128
            _ = v157
        default:
            if (v1 != v3) {
                for fi158 := 0; fi158 < 3; fi158++ {
                    v3 = v1
                    if (true && true) {
                        switch "\n\t" {
                        case "a":
                            if (925 > 42) {
                                if true {
                                    if (true || false) {
                                        switch fi158 {
                                        case 9223372036854775807:
                                            for fi159 := 0; fi159 < 4; fi159++ {
                                                _ = len(v2)
                                                v2[-39] = 2
                                                _ = len(v5)
                                                _ = fi159
                                            }
                                        default:
                                            var v160 float32 = v3
                                            _ = v160
                                        }
                                        _ = len(v2)
                                    } else {
                                        var v161 uint64 = 4294967295
                                        _ = v161
                                    }
                                } else {
                                    if (-6 == 42) {
                                        if (249 == 46) {
                                            if false {
                                                if false {
                                                    for fi162 := 0; fi162 < 3; fi162++ {
                                                        switch -404 {
                                                        case 32768:
                                                            var v163 string = ("hello" + ("foo bar" + "foo bar"))
                                                            _ = v163
                                                        case 9223372036854775807:
                                                            v164 := GenS19{F0: "", F1: 18446744073709551615}
                                                            _ = v164
                                                        case -9223372036854775808:
                                                            var v165 int32 = (v6 >> ((42) & 63))
                                                            _ = v165
                                                        default:
                                                            var v166 uint8 = 1
                                                            if (62 > 100) {
                                                                v4 = v4
                                                            }
                                                            _ = v166
                                                        }
                                                        _ = fi162
                                                    }
                                                }
                                            }
                                            fmt.Println(v1)
                                        } else {
                                            fmt.Println(fi158, v1)
                                        }
                                        var v167 uint32 = 2147483648
                                        v3 -= (v1 + ((v1 / (v3 + v1)) * v1))
                                        _ = v167
                                    }
                                }
                                v168 := GenS24{F0: v4, F1: 450, F2: 0}
                                _ = v168
                            }
                        case "\n\t":
                            for fi169 := 0; fi169 < 10; fi169++ {
                                v3 = 1e10
                                _ = fi169
                            }
                        case "hello":
                            switch "a" {
                            case "\n\t":
                                v170 := GenS25{F0: 42, F1: 674, F2: v6}
                                _ = v170
                            case "":
                                fmt.Println(fi158, v1)
                            }
                        default:
                            v1 += (v3 * float32(float64(1e10)))
                        }
                        for fi171 := 0; fi171 < 1; fi171++ {
                            fmt.Println(fi158, v4)
                            _ = len(v2)
                            _ = fi171
                        }
                        if true {
                            fmt.Println(v4)
                        } else {
                            v3 = v3
                        }
                    } else {
                        if (!(true || false)) {
                            fmt.Println(v3, v3)
                        }
                        var v172 float64 = 1e10
                        _ = v172
                    }
                    _ = fi158
                }
                fmt.Println(v3)
                var v173 uint64 = 1
                _ = v173
            }
        }
        if false {
            for fi174 := 0; fi174 < 9; fi174++ {
                switch "a" {
                case "\n\t":
                    if false {
                        for fi175 := 0; fi175 < 18; fi175++ {
                            fmt.Println(v4, fi174, fi174)
                            v176 := v2[980]
                            _ = fi175
                            _ = v176
                        }
                    } else {
                        v177 := GenS3{F0: 42, F1: 255, F2: fi174}
                        if false {
                            var v178 int8 = 127
                            _ = v178
                        }
                        if (v6 >= 42) {
                            _ = len(v5)
                        } else {
                            var v179 uint16 = 2
                            switch "" {
                            case "a":
                                var v180 int = (fi174 - (fi174 << ((0) & 63)))
                                _ = v180
                            case "foo bar":
                                fmt.Println(v6, v3)
                            }
                            _ = v179
                        }
                        _ = v177
                    }
                case "foo bar":
                    v6 -= (v6 & -2147483648)
                case "a":
                    if (!true) {
                        v181 := GenS19{F0: ("foo bar" + ""), F1: 1}
                        for fi182 := 0; fi182 < 6; fi182++ {
                            fmt.Println(v1)
                            var v183 int16 = v4
                            v184 := GenS19{F0: "hello", F1: 4294967295}
                            _ = fi182
                            _ = v183
                            _ = v184
                        }
                        var v185 bool = false
                        _ = v181
                        _ = v185
                    } else {
                        v186 := GenS17{F0: false, F1: 939}
                        for fi187 := 0; fi187 < 15; fi187++ {
                            var v188 float64 = (1e10 / 1e10)
                            v1 = (v1 + 1e10)
                            _ = fi187
                            _ = v188
                        }
                        if false {
                            switch (fi174 % (((fi174 - fi174)) | 1)) {
                            case 4294967295:
                                if (v3 == (0.0 * 1e10)) {
                                    var v189 uint32 = 407
                                    v1 += 0.1
                                    var v190 uint8 = 128
                                    _ = v189
                                    _ = v190
                                }
                            case -9223372036854775808:
                                switch "a" {
                                case "a":
                                    v1 = (v3 - v1)
                                case "\n\t":
                                    for fi191 := 0; fi191 < 8; fi191++ {
                                        _ = len(v5)
                                        _ = fi191
                                    }
                                case "foo bar":
                                    for fi192 := 0; fi192 < 1; fi192++ {
                                        v4 += v4
                                        v2[(^(fi174))] = 4294967295
                                        _ = fi192
                                    }
                                }
                            case 128:
                                _ = len(v5)
                            default:
                                for fi193 := 0; fi193 < 8; fi193++ {
                                    var v194 float64 = -2.25
                                    v3 += 1e10
                                    _ = fi193
                                    _ = v194
                                }
                            }
                            var v195 float32 = v1
                            _ = v195
                        }
                        _ = v186
                    }
                }
                if (119 >= 111) {
                    v196 := GenS7{F0: 2, F1: v4}
                    _ = len(v2)
                    _ = v196
                }
                var v197 int16 = (v4 % ((32767) | 1))
                _ = fi174
                _ = v197
            }
            v198 := GenS26{F0: -547, F1: -483, F2: (!((18446744073709551615 >= 42) || ((((((v3 - -2.25) - -2.25) - 1.5) / v1) >= (v1 / (v3 + float32(1.5)))) || false)))}
            _ = len(v2)
            _ = v198
        } else {
            var v199 uint32 = 0
            var v200 bool = false
            var v201 uint32 = v199
            _ = v199
            _ = v200
            _ = v201
        }
        var v202 uint64 = 2
        _ = v202
    } else {
        v3 -= v3
        v6 = 42
    }
    v203 := GenS9{F0: (("a" + "a") + "hello"), F1: 0, F2: 356}
    var v204 int16 = (-(v4))
    v3 += 1.5
    v205 := GenS27{F0: v4, F1: 696, F2: 95}
    v206 := GenS23{F0: false, F1: v6, F2: "foo bar"}
    var v207 int = 32768
    var v208 uint32 = 65535
    v4 = (v4 >> ((255) & 63))
    for fi209 := 0; fi209 < 7; fi209++ {
        var v210 int64 = -9223372036854775808
        _ = fi209
        _ = v210
    }
    var v211 float64 = ((1.5 + ((0.0 * -2.25) * 3.14159)) / (((1e10 - 1e10) + (-2.25 / 1.0)) * -2.25))
    for fi212 := 0; fi212 < 14; fi212++ {
        var v213 uint16 = 42
        var v214 string = "foo bar"
        v5 = append(v5, v207)
        _ = fi212
        _ = v213
        _ = v214
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v204:", v204)
    fmt.Println("v207:", v207)
    fmt.Println("v208:", v208)
    fmt.Println("v211:", v211)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v203
    _ = v204
    _ = v205
    _ = v206
    _ = v207
    _ = v208
    _ = v211
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
