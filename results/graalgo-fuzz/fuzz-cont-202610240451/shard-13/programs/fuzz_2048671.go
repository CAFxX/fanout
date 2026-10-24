package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 string; F1 uint64; F2 uint32 }
type GenS8 struct { F0 int32; F1 uint64; F2 int }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint32; F1 int16 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint8; F1 uint8 }
type GenS14 struct { F0 int8; F1 uint64; F2 uint }

func main() {
    var v1 bool = false
    v2 := []int{-646, 0, 32768, -663}
    var v3 float64 = 1.5
    var v4 int8 = 60
    var v5 uint64 = 4294967295
    var v6 float32 = 3.14159
    v7 := GenS1{F0: (v5 + (v5 - 325)), F1: 138}
    v6 = v6
    var v8 bool = (-455 < int(v4))
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int8 = (v4 ^ (v4 ^ (v4 % ((v4) | 1))))
        var v11 int16 = 1000
        var v12 float64 = (((3.14159 / -2.25) - v3) + v3)
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 uint = 255
    if (false || (0 < 301)) {
        v14 := GenS1{F0: 657, F1: 199}
        if len(v2) > 0 {
            v2[0] = 9223372036854775807
        }
        v13 += v13
        _ = v14
    }
    v6 = v6
    fmt.Println(v6, v3, v3)
    var v15 string = ("a" + (("a" + (("a" + "a") + (("a" + "foo bar") + "hello"))) + (((("\n\t" + "") + "\n\t") + "hello") + "hello")))
    var v16 int64 = 9223372036854775807
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 int = -1
        var v19 uint8 = 2
        if (70 < 42) {
            v20 := GenS1{F0: v5, F1: (v19 | 2)}
            if v8 {
                var v21 uint16 = 1
                v22 := GenS1{F0: 2, F1: (v19 * (^(v19)))}
                _ = v21
                _ = v22
            } else {
                if (0 >= 1) {
                    v18 = (fi17 ^ 1)
                    if (((v3 - float64(0.1)) / v3) >= (float64(v6) / ((0.0 * 3.14159) - -2.25))) {
                        for fi23 := 0; fi23 < 10; fi23++ {
                            fmt.Println(v16, v19)
                            _ = fi23
                        }
                        v8 = v8
                        switch -2147483648 {
                        case -3:
                            switch ("" + v15) {
                            case "hello":
                                if ((v4 / ((v4) | 1)) <= 127) {
                                    var v24 uint8 = 128
                                    for fi25 := 0; fi25 < 6; fi25++ {
                                        v3 -= v3
                                        _ = fi25
                                    }
                                    _ = v24
                                }
                            case "a":
                                v26 := GenS1{F0: 834, F1: (v19 & 128)}
                                _ = v26
                            default:
                                switch "hello" {
                                case "":
                                    if v8 {
                                        if (true || (1 > 32767)) {
                                            v27 := GenS2{F0: int8(v18)}
                                            _ = v27
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = -1
                                            }
                                            v4 = 74
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        fmt.Println(v13)
                                        if false {
                                            var v28 bool = v1
                                            switch int(v5) {
                                            case 127:
                                                var v29 int = -506
                                                _ = v29
                                            case 32767:
                                                switch "" {
                                                case "\n\t":
                                                    switch (v15 + "a") {
                                                    case "":
                                                        switch v15 {
                                                        case "":
                                                            if (!(v28 || false)) {
                                                                fmt.Println(v8, v6, v1)
                                                                v15 = v15
                                                            } else {
                                                                v30 := GenS1{F0: (v5 % ((18446744073709551615) | 1)), F1: v19}
                                                                _ = v30
                                                            }
                                                        case "hello":
                                                            var v31 uint64 = 0
                                                            _ = v31
                                                        default:
                                                            for fi32 := 0; fi32 < 6; fi32++ {
                                                                var v33 float32 = (float32((v3 / 3.14159)) / v6)
                                                                var v34 int64 = -1
                                                                v28 = false
                                                                _ = fi32
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v19 += (^(v19))
                                                    default:
                                                        for fi35 := 0; fi35 < 4; fi35++ {
                                                            v1 = false
                                                            _ = fi35
                                                        }
                                                    }
                                                case "hello":
                                                    if len(v2) > 0 {
                                                        v2[0] = fi17
                                                    }
                                                case "foo bar":
                                                    v15 = ("a" + v15)
                                                default:
                                                    for fi36 := 0; fi36 < 6; fi36++ {
                                                        v28 = ("foo bar" >= v15)
                                                        var v37 int16 = 42
                                                        var v38 int = fi17
                                                        _ = fi36
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                }
                                            default:
                                                if (!v28) {
                                                    v15 = v15
                                                    v3 = 0.0
                                                }
                                            }
                                            _ = v28
                                        } else {
                                            if (0 >= 32767) {
                                                for fi39 := 0; fi39 < 3; fi39++ {
                                                    if v8 {
                                                        if false {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            if len(v2) > 0 {
                                                                v2[0] = (fi39 - fi39)
                                                            }
                                                        }
                                                    } else {
                                                        v4 = 127
                                                        if len(v2) > 0 {
                                                            v2[0] = -9223372036854775808
                                                        }
                                                        fmt.Println(v5, v13, v1)
                                                    }
                                                    _ = fi39
                                                }
                                                v8 = (!v1)
                                                if false {
                                                    var v40 int = -942
                                                    var v41 float32 = v6
                                                    fmt.Println(v13, v8, v18)
                                                    _ = v40
                                                    _ = v41
                                                } else {
                                                    var v42 uint8 = 128
                                                    v2 = append(v2, v18)
                                                    for fi43 := 0; fi43 < 12; fi43++ {
                                                        _ = len(v2)
                                                        fmt.Println(fi43)
                                                        _ = fi43
                                                    }
                                                    _ = v42
                                                }
                                            }
                                            _ = len(v2)
                                            if (!true) {
                                                fmt.Println(v15, v4, v8)
                                                switch (v18 + -9) {
                                                case 127:
                                                    _ = len(v2)
                                                case 32767:
                                                    if (uint32(v19) <= 340) {
                                                        for fi44 := 0; fi44 < 2; fi44++ {
                                                            _ = len(v2)
                                                            _ = fi44
                                                        }
                                                    } else {
                                                        var v45 bool = false
                                                        v46 := GenS3{F0: (^(v16))}
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                case -867:
                                                    fmt.Println(v3, v1)
                                                }
                                            } else {
                                                fmt.Println(fi17, v8, v19)
                                                if (!((v18 & v18) >= (v18 << ((v13) & 63)))) {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    fmt.Println(v18, v13)
                                                } else {
                                                    fmt.Println(v13, v4)
                                                    if (0.0 >= float64(float32(v3))) {
                                                        for fi47 := 0; fi47 < 20; fi47++ {
                                                            var v48 uint64 = 2
                                                            v15 = (v15 + (v15 + ("a" + "a")))
                                                            var v49 int64 = (-(v16))
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = fi47
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                    } else {
                                                        for fi50 := 0; fi50 < 3; fi50++ {
                                                            var v51 uint = (v13 & v13)
                                                            v1 = (0.0 != 0.1)
                                                            if (!v1) {
                                                                if true {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v5 += 710
                                                                    v18 += int(v19)
                                                                }
                                                                var v52 bool = v8
                                                                var v53 uint64 = v5
                                                                switch (v18 % ((42) | 1)) {
                                                                case -9223372036854775808:
                                                                    _ = len(v2)
                                                                case -529:
                                                                    var v54 uint32 = 4294967295
                                                                    _ = v54
                                                                case 9223372036854775807:
                                                                    var v55 uint = v13
                                                                    _ = v55
                                                                default:
                                                                    var v56 float64 = v3
                                                                    _ = v56
                                                                }
                                                                var v57 int16 = -594
                                                                _ = v52
                                                                _ = v53
                                                                _ = v57
                                                            }
                                                            _ = fi50
                                                            _ = v51
                                                        }
                                                    }
                                                    for fi58 := 0; fi58 < 20; fi58++ {
                                                        v18 = fi58
                                                        v19 += v19
                                                        if v1 {
                                                            fmt.Println(v3, v5)
                                                            var v59 float64 = v3
                                                            _ = v59
                                                        } else {
                                                            var v60 float64 = (v3 + float64((v6 * (v6 - float32((-2.25 - (v3 * float64((-2.25 * -2.25)))))))))
                                                            var v61 uint64 = v5
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                        _ = fi58
                                                    }
                                                }
                                            }
                                        }
                                        var v62 uint16 = 65535
                                        _ = v62
                                    }
                                }
                            }
                        }
                    }
                    var v63 float64 = float64(((v6 * (v6 - v6)) + 3.14159))
                    _ = v63
                }
            }
            switch -1 {
            case 0:
                var v64 float64 = v3
                _ = v64
            case -9223372036854775808:
                for fi65 := 0; fi65 < 16; fi65++ {
                    v16 += v16
                    _ = fi65
                }
            }
            _ = v20
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    if (v1 && false) {
        switch v15 {
        case "a":
            v16 += 0
        case "foo bar":
            v3 = 0.1
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v66 int16 = 1000
            _ = v66
        }
        var v67 uint16 = 65535
        _ = v67
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v4, v4, v4)
        fmt.Println(v3)
    }
    fmt.Println(v6, v13, v5)
    for fi68 := 0; fi68 < 5; fi68++ {
        v69 := GenS4{F0: 255, F1: 128}
        fmt.Println(v6, v16)
        _ = fi68
        _ = v69
    }
    switch v15 {
    case "hello":
        if v8 {
            for fi70 := 0; fi70 < 20; fi70++ {
                fmt.Println(v1, v1)
                if (1000 <= 1000) {
                    if len(v2) > 0 {
                        v2[0] = (fi70 * 127)
                    }
                    var v71 string = (v15 + v15)
                    _ = v71
                } else {
                    for fi72 := 0; fi72 < 16; fi72++ {
                        if (!(1 != v5)) {
                            switch (fi72 / ((2147483647) | 1)) {
                            case 0:
                                v6 += 0.1
                            case 32768:
                                v8 = ((0 != (v5 - 138)) && false)
                            case 968:
                                switch 32767 {
                                case -2147483648:
                                    v3 = 1e10
                                case 101:
                                    switch "" {
                                    case "foo bar":
                                        v73 := GenS5{F0: int16(v16)}
                                        _ = v73
                                    case "a":
                                        v74 := GenS1{F0: (v5 & v5), F1: 255}
                                        _ = v74
                                    case "hello":
                                        var v75 int32 = int32(v4)
                                        _ = v75
                                    }
                                }
                            }
                            if ((v4 >> ((1) & 63)) != (v4 >> ((v13) & 63))) {
                                var v76 bool = (0 != 0)
                                var v77 uint64 = v5
                                _ = v76
                                _ = v77
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v78 uint32 = 42
                                v2 = append(v2, (fi70 | fi72))
                                _ = v78
                            }
                            var v79 uint8 = 128
                            _ = v79
                        } else {
                            if len(v2) > 0 {
                                v2[0] = (fi70 & (fi70 & int(v16)))
                            }
                        }
                        v80 := GenS2{F0: v4}
                        _ = fi72
                        _ = v80
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v4 += v4
                }
                _ = fi70
            }
        } else {
            v1 = v1
        }
    case "foo bar":
        v81 := GenS6{F0: v4, F1: (v4 / ((100) | 1))}
        _ = v81
    case "a":
        if ((42 >= 42) && v8) {
            if false {
                switch "" {
                case "":
                    v15 = (v15 + v15)
                default:
                    v82 := GenS7{F0: "foo bar", F1: v5, F2: 74}
                    _ = v82
                }
                v1 = ((v4 / ((v4) | 1)) < v4)
                var v83 uint16 = 1
                _ = v83
            }
            if v1 {
                v84 := GenS8{F0: 1, F1: (v5 - v5), F2: -9223372036854775808}
                _ = v84
            } else {
                fmt.Println(v6, v13, v1)
            }
        } else {
            fmt.Println(v8, v3, v8)
            if (!v8) {
                if (-1 > 588) {
                    if v8 {
                        for fi85 := 0; fi85 < 7; fi85++ {
                            for fi86 := 0; fi86 < 10; fi86++ {
                                v13 = v13
                                _ = fi86
                            }
                            _ = fi85
                        }
                        var v87 float32 = v6
                        _ = v87
                    } else {
                        switch -394 {
                        case 32768:
                            v15 = "a"
                        default:
                            v88 := GenS9{F0: 1, F1: -32768}
                            _ = v88
                        }
                    }
                    v89 := GenS10{F0: 1}
                    if ((v13 | (v13 - (v13 >> ((1) & 63)))) < 255) {
                        if (v1 || (32768 < 637)) {
                            v90 := GenS6{F0: v4, F1: v4}
                            var v91 int16 = -32768
                            var v92 float64 = (3.14159 + 1e10)
                            _ = v90
                            _ = v91
                            _ = v92
                        } else {
                            var v93 uint32 = 65535
                            var v94 int = 926
                            _ = v93
                            _ = v94
                        }
                        switch v15 {
                        case "\n\t":
                            for fi95 := 0; fi95 < 18; fi95++ {
                                if (v4 >= (v4 >> ((42) & 63))) {
                                    if (!((255 >= (v13 & (v13 << ((65535) & 63)))) || (v16 <= (v16 - 9223372036854775807)))) {
                                        fmt.Println(v1, v15, v8)
                                        v2 = append(v2, -228)
                                        v96 := GenS10{F0: 1}
                                        _ = v96
                                    } else {
                                        _ = len(v2)
                                        v1 = (!v1)
                                        fmt.Println(fi95, v6)
                                    }
                                }
                                v4 -= (v4 | 30)
                                var v97 uint8 = 146
                                _ = fi95
                                _ = v97
                            }
                        case "":
                            v5 += (^(v5))
                        case "hello":
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                        default:
                            var v98 string = "hello"
                            _ = v98
                        }
                        switch (v15 + ((("foo bar" + v15) + (v15 + v15)) + v15)) {
                        case "hello":
                            if v8 {
                                switch "a" {
                                case "a":
                                    fmt.Println(v6, v3)
                                default:
                                    var v99 int32 = 830
                                    _ = v99
                                }
                            } else {
                                switch -975 {
                                case -316:
                                    _ = len(v2)
                                case -6:
                                    _ = len(v2)
                                default:
                                    v2 = append(v2, 4294967295)
                                }
                                for fi100 := 0; fi100 < 3; fi100++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v8, v4, v4)
                                    _ = fi100
                                }
                                fmt.Println(v4, v5)
                            }
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    } else {
                        if true {
                            v13 = 1
                            for fi101 := 0; fi101 < 8; fi101++ {
                                fmt.Println(v3)
                                v13 = v13
                                v102 := GenS10{F0: 42}
                                _ = fi101
                                _ = v102
                            }
                            for fi103 := 0; fi103 < 17; fi103++ {
                                for fi104 := 0; fi104 < 17; fi104++ {
                                    v4 = (v4 << ((v13) & 63))
                                    if len(v2) > 0 {
                                        v2[0] = (fi103 << (((v13 | (v13 << (((v13 / ((v13) | 1))) & 63)))) & 63))
                                    }
                                    fmt.Println(v16, v16)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if v8 {
                                        v16 += (v16 & (v16 * -9223372036854775808))
                                        fmt.Println(v15)
                                        if ((0.1 - -2.25) >= 0.0) {
                                            var v105 uint16 = 255
                                            _ = v105
                                        }
                                    } else {
                                        if v8 {
                                            v1 = v1
                                            v15 = (v15 + "")
                                            _ = len(v2)
                                        } else {
                                            fmt.Println(v6, fi103)
                                            v6 = v6
                                            var v106 int16 = -1
                                            _ = v106
                                        }
                                        v107 := GenS7{F0: v15, F1: v5, F2: 4294967295}
                                        _ = v107
                                    }
                                    _ = fi104
                                }
                                v6 = v6
                                v13 -= (v13 >> ((v13) & 63))
                                _ = fi103
                            }
                        }
                        if v8 {
                            var v108 uint32 = 65535
                            _ = v108
                        } else {
                            v16 = v16
                        }
                        var v109 bool = v8
                        _ = v109
                    }
                    _ = v89
                } else {
                    if v1 {
                        v110 := GenS11{F0: 0, F1: -690}
                        v6 = v6
                        _ = v110
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            v111 := GenS12{F0: (!(v5 < 0))}
                            _ = v111
                        case "":
                            v112 := GenS13{F0: 42, F1: 1}
                            _ = v112
                        case "\n\t":
                            v13 = (v13 & (v13 | (v13 + (v13 | (v13 / ((1) | 1))))))
                        default:
                            var v113 float32 = 1e10
                            _ = v113
                        }
                        switch v15 {
                        case "a":
                            var v114 bool = (true || ((v15 + "hello") >= "\n\t"))
                            _ = v114
                        case "":
                            if true {
                                switch 128 {
                                case -1:
                                    if v8 {
                                        v2 = append(v2, 32768)
                                        v5 = v5
                                        if (173 == 0) {
                                            v4 -= (-(v4))
                                            for fi115 := 0; fi115 < 11; fi115++ {
                                                fmt.Println(v13, v15, v16)
                                                var v116 int32 = 42
                                                _ = fi115
                                                _ = v116
                                            }
                                        }
                                    } else {
                                        v117 := GenS7{F0: v15, F1: uint64(v13), F2: 4294967295}
                                        _ = v117
                                    }
                                }
                                v4 -= (-(v4))
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        default:
                            v118 := GenS8{F0: 1, F1: (v5 - v5), F2: -9223372036854775808}
                            _ = v118
                        }
                        v2 = append(v2, 32768)
                    }
                    switch ("hello" + "a") {
                    case "a":
                        var v119 uint32 = 145
                        _ = v119
                    case "foo bar":
                        if v8 {
                            for fi120 := 0; fi120 < 17; fi120++ {
                                var v121 int16 = 283
                                var v122 int8 = (v4 * v4)
                                _ = fi120
                                _ = v121
                                _ = v122
                            }
                            var v123 uint64 = (^(v5))
                            if false {
                                switch v15 {
                                case "":
                                    v124 := GenS14{F0: v4, F1: v123, F2: v13}
                                    _ = v124
                                case "foo bar":
                                    var v125 string = "hello"
                                    _ = v125
                                default:
                                    v16 = (^(v16))
                                }
                                switch -2147483648 {
                                case 127:
                                    v4 += -85
                                case 4294967295:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case -460:
                                    var v126 uint = v13
                                    _ = v126
                                }
                            }
                            _ = v123
                        }
                    }
                    v4 -= v4
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
