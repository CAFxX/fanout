package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS9 struct { F0 int; F1 int64 }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 int16; F1 uint32; F2 int32 }
type GenS12 struct { F0 bool; F1 uint8; F2 uint }
type GenS13 struct { F0 uint64; F1 uint16 }
type GenS14 struct { F0 int16; F1 uint8; F2 uint }
type GenS15 struct { F0 uint8; F1 uint8 }
type GenS16 struct { F0 uint32 }
func genFn0() bool {
    var v5 uint16 = 255
    v5 = (v5 + (v5 % (((^(v5))) | 1)))
    _ = v5
    return true
}

func main() {
    v1 := []string{"a", ("\n\t" + ""), ""}
    var v2 uint64 = 1
    v3 := make(map[int]string)
    v4 := []int{218, 969}
    var v6 int16 = 1
    switch "" {
    case "a":
        if (false && true) {
            var v7 uint8 = 0
            if false {
                v7 += (v7 + (v7 & 139))
            } else {
                switch ("hello" + "hello") {
                case "\n\t":
                    v6 -= v6
                case "":
                    var v8 int8 = -98
                    _ = v8
                }
            }
            v9 := GenS1{F0: 18446744073709551615, F1: -2147483648, F2: true}
            _ = v7
            _ = v9
        } else {
            v6 = 765
            v10 := GenS2{F0: 0}
            if (9223372036854775807 != 9223372036854775807) {
                delete(v3, 1)
                v6 -= v6
                if true {
                    if (("hello" + ("a" + "")) <= "a") {
                        v2 = (v2 >> ((2) & 63))
                        v11 := GenS3{F0: 2147483647, F1: true}
                        _ = v11
                    }
                    switch ("a" + "hello") {
                    case "":
                        switch "a" {
                        case "":
                            switch ("hello" + "a") {
                            case "\n\t":
                                fmt.Println(v6, v6)
                            default:
                                switch (("\n\t" + ("foo bar" + ("foo bar" + "hello"))) + "\n\t") {
                                case "a":
                                    fmt.Println(v2)
                                case "\n\t":
                                    switch (("\n\t" + "a") + (("foo bar" + ("hello" + "\n\t")) + "\n\t")) {
                                    case "foo bar":
                                        var v12 int8 = 127
                                        _ = v12
                                    case "a":
                                        v2 = 4294967295
                                    case "":
                                        if true {
                                            var v13 int64 = 937
                                            for fi14 := 0; fi14 < 11; fi14++ {
                                                var v15 uint16 = 989
                                                for fi16 := 0; fi16 < 20; fi16++ {
                                                    switch fi14 {
                                                    case -829:
                                                        v2 = v2
                                                    case 127:
                                                        switch -205 {
                                                        case -717:
                                                            var v17 string = "a"
                                                            _ = v17
                                                        case -94:
                                                            _ = len(v1)
                                                        case 397:
                                                            v18 := GenS3{F0: 2147483647, F1: false}
                                                            _ = v18
                                                        }
                                                        var v19 int8 = 1
                                                        if (true && false) {
                                                            if false {
                                                                fmt.Println(fi16)
                                                                var v20 uint = 4294967295
                                                                _ = v20
                                                            }
                                                            v21 := GenS4{F0: (^(v19))}
                                                            fmt.Println(v6, v6, fi16)
                                                            _ = v21
                                                        } else {
                                                            var v22 int = 455
                                                            var v23 int = fi14
                                                            v13 = v13
                                                            _ = v22
                                                            _ = v23
                                                        }
                                                        _ = v19
                                                    }
                                                    if (!(!false)) {
                                                        v15 += 51
                                                        v24 := GenS3{F0: 829, F1: false}
                                                        v2 += v2
                                                        _ = v24
                                                    }
                                                    switch 9223372036854775807 {
                                                    case -2147483648:
                                                        _ = len(v3)
                                                    case -9223372036854775808:
                                                        v25 := GenS5{F0: v6, F1: ("foo bar" + "")}
                                                        v26 := GenS6{F0: (fi16 != fi16)}
                                                        _ = v25
                                                        _ = v26
                                                    case 32767:
                                                        if false {
                                                            var v27 int16 = 23
                                                            var v28 int64 = (v13 - 0)
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                    default:
                                                        if false {
                                                            if true {
                                                                v2 += (v2 << ((778) & 63))
                                                                v13 += v13
                                                            } else {
                                                                if (!true) {
                                                                    if (float32((float64((1.5 / 1.0)) / -2.25)) > 0.1) {
                                                                        v29 := GenS1{F0: 489, F1: (fi16 | (fi14 * 2147483647)), F2: ((329 > -1) || (4294967295 < 255))}
                                                                        _ = v29
                                                                    }
                                                                    var v30 uint8 = 255
                                                                    v31 := GenS7{F0: (v2 ^ 853)}
                                                                    v32 := GenS5{F0: (v6 % ((v6) | 1)), F1: "\n\t"}
                                                                    v33 := GenS8{F0: -1, F1: v30, F2: (v2 + 1)}
                                                                    _ = v30
                                                                    _ = v31
                                                                    _ = v32
                                                                    _ = v33
                                                                } else {
                                                                    if false {
                                                                        var v34 uint64 = v2
                                                                        _ = len(v1)
                                                                        switch "a" {
                                                                        case "":
                                                                            fmt.Println(v15, v13, fi14)
                                                                            if true {
                                                                                if (true || ((!(int(v2) >= fi14)) && (v6 >= (v6 >> ((946) & 63))))) {
                                                                                    var v35 uint32 = 2147483647
                                                                                    v36 := GenS4{F0: 100}
                                                                                    fmt.Println(v35, fi16)
                                                                                    _ = v35
                                                                                    _ = v36
                                                                                } else {
                                                                                    if (0.0 <= 1e10) {
                                                                                        var v37 float64 = 1e10
                                                                                        _ = len(v3)
                                                                                        _ = v37
                                                                                    } else {
                                                                                        switch (fi14 ^ 257) {
                                                                                        case 2147483647:
                                                                                            fmt.Println(v34)
                                                                                        case 32767:
                                                                                            v13 += v13
                                                                                        case -1:
                                                                                            v34 = v34
                                                                                        }
                                                                                        v34 = v34
                                                                                        v38 := GenS7{F0: (v34 - v2)}
                                                                                        _ = v38
                                                                                    }
                                                                                }
                                                                                switch ("foo bar" + "foo bar") {
                                                                                case "a":
                                                                                    var v39 float64 = 0.1
                                                                                    _ = v39
                                                                                case "foo bar":
                                                                                    if (1 < 42) {
                                                                                        var v40 int64 = v13
                                                                                        v41 := GenS9{F0: 32768, F1: v40}
                                                                                        _ = v40
                                                                                        _ = v41
                                                                                    } else {
                                                                                        v15 = 255
                                                                                    }
                                                                                default:
                                                                                    var v42 uint16 = 65535
                                                                                    _ = v42
                                                                                }
                                                                                var v43 int16 = (v6 % ((32767) | 1))
                                                                                v6 = v43
                                                                                var v44 int16 = (v43 >> ((18446744073709551615) & 63))
                                                                                _ = v43
                                                                                _ = v44
                                                                            } else {
                                                                                var v45 uint32 = 2147483648
                                                                                v46 := GenS8{F0: -1, F1: 212, F2: (v34 & (v34 + 2))}
                                                                                _ = v45
                                                                                _ = v46
                                                                            }
                                                                            _ = len(v4)
                                                                        default:
                                                                            var v47 int8 = 100
                                                                            _ = v47
                                                                        }
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        v15 += (v15 + uint16(v34))
                                                                        _ = v34
                                                                    }
                                                                    var v48 uint8 = 1
                                                                    _ = v48
                                                                }
                                                            }
                                                        } else {
                                                            if true {
                                                                var v49 bool = true
                                                                fmt.Println(v13, fi14)
                                                                v15 -= (v15 / ((v15) | 1))
                                                                if (!false) {
                                                                    switch (fi16 >> ((639) & 63)) {
                                                                    case 682:
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            var v50 int16 = 963
                                                                            _ = v50
                                                                        case "foo bar":
                                                                            v15 = v15
                                                                        case "hello":
                                                                            v2 = v2
                                                                        }
                                                                    case -871:
                                                                        v4 = append(v4, fi14)
                                                                    default:
                                                                        v49 = (v49 || false)
                                                                    }
                                                                    var v51 int32 = -1
                                                                    var v52 float64 = (1e10 * (0.1 * (3.14159 + ((1e10 * (3.14159 * 3.14159)) - (-2.25 * 0.1)))))
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                                _ = v49
                                                            } else {
                                                                var v53 uint8 = uint8(v15)
                                                                v15 += (v15 ^ uint16(v2))
                                                                _ = v53
                                                            }
                                                            v2 = uint64(v15)
                                                            v2 += 0
                                                        }
                                                    }
                                                    _ = fi16
                                                }
                                                _ = fi14
                                                _ = v15
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                switch 32768 {
                                                case 127:
                                                    switch -958 {
                                                    case -9223372036854775808:
                                                        if (!false) {
                                                            var v54 float64 = ((0.1 - float64(0.1)) / (1e10 + 1.5))
                                                            fmt.Println(v2, v6, v6)
                                                            _ = v54
                                                        }
                                                    case 32768:
                                                        fmt.Println(v2, v13, v13)
                                                    }
                                                default:
                                                    v13 = v13
                                                }
                                            case "":
                                                v55 := GenS6{F0: (!false)}
                                                _ = v55
                                            default:
                                                v2 += (v2 + v2)
                                            }
                                            _ = v13
                                        }
                                    default:
                                        v6 -= (v6 + 1000)
                                    }
                                case "foo bar":
                                    v6 = v6
                                default:
                                    var v56 uint = 1
                                    _ = v56
                                }
                            }
                        default:
                            if (1 < 195) {
                                for fi57 := 0; fi57 < 4; fi57++ {
                                    var v58 float32 = (0.0 - (-2.25 - 3.14159))
                                    fmt.Println(v2, v2, v2)
                                    _ = fi57
                                    _ = v58
                                }
                            }
                        }
                    case "hello":
                        switch ("a" + (("hello" + "foo bar") + "")) {
                        case "hello":
                            v6 -= -1
                        default:
                            var v59 int32 = 2147483647
                            _ = v59
                        }
                    default:
                        if false {
                            v3[128] = ("foo bar" + "hello")
                        } else {
                            v60 := GenS10{F0: (uint16(v2) > 42), F1: (^(v2))}
                            if true {
                                var v61 uint8 = 127
                                _ = v61
                            } else {
                                v62 := GenS8{F0: -2147483648, F1: 1, F2: 18446744073709551615}
                                v6 -= (v6 % (((-(v6))) | 1))
                                var v63 float64 = 1.5
                                _ = v62
                                _ = v63
                            }
                            _ = v60
                        }
                    }
                    v2 = v2
                } else {
                    for fi64 := 0; fi64 < 13; fi64++ {
                        switch "hello" {
                        case "a":
                            var v65 uint16 = 32767
                            _ = v65
                        case "foo bar":
                            var v66 int = fi64
                            _ = v66
                        default:
                            var v67 float64 = 1.5
                            _ = v67
                        }
                        _ = fi64
                    }
                    for fi68 := 0; fi68 < 20; fi68++ {
                        v69 := v3[(fi68 + -1)]
                        _ = fi68
                        _ = v69
                    }
                    if false {
                        v2 -= 0
                    } else {
                        v70 := GenS6{F0: true}
                        _ = v70
                    }
                }
            }
            _ = v10
        }
    default:
        v71 := GenS11{F0: v6, F1: 126, F2: 2147483647}
        _ = v71
    }
    switch 42 {
    case -774:
        var v72 float64 = 1.5
        _ = v72
    case 1:
        if true {
            fmt.Println(v2)
            var v73 int32 = -1
            _ = v73
        } else {
            switch 0 {
            case 4294967295:
                v2 += v2
            default:
                var v74 uint32 = 2147483647
                _ = v74
            }
        }
    case 127:
        v6 -= (v6 - (^(v6)))
    default:
        if (!(2147483647 <= 42)) {
            if false {
                v2 = (v2 << ((18446744073709551615) & 63))
                v75 := GenS8{F0: 42, F1: 0, F2: (v2 - v2)}
                if (!false) {
                    for fi76 := 0; fi76 < 7; fi76++ {
                        if (1 < -1) {
                            var v77 bool = false
                            _ = v77
                        } else {
                            delete(v3, fi76)
                        }
                        if (!true) {
                            if ((255 != 255) && (v6 < -1)) {
                                for fi78 := 0; fi78 < 10; fi78++ {
                                    v79 := GenS8{F0: 1, F1: 1, F2: (v2 | (v2 >> ((988) & 63)))}
                                    fmt.Println(fi76)
                                    _ = fi78
                                    _ = v79
                                }
                                switch 42 {
                                case 830:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            fmt.Println(v2, v2, v6)
                            var v80 uint8 = 136
                            _ = v80
                        } else {
                            v2 = v2
                            var v81 float64 = (1.5 * (3.14159 - 1.5))
                            v6 += 1
                            _ = v81
                        }
                        _ = fi76
                    }
                    var v82 bool = false
                    _ = v82
                }
                _ = v75
            } else {
                v83 := v3[0]
                _ = v83
            }
            v84 := v3[1]
            _ = v84
        }
    }
    if (127 >= 9223372036854775807) {
        for fi85 := 0; fi85 < 15; fi85++ {
            for fi86 := 0; fi86 < 4; fi86++ {
                v87 := GenS12{F0: (0 < 0), F1: 121, F2: 42}
                _ = fi86
                _ = v87
            }
            var v88 uint8 = 128
            if false {
                v6 += (v6 | (v6 ^ (v6 % ((v6) | 1))))
                switch (("" + "\n\t") + "foo bar") {
                case "":
                    for fi89 := 0; fi89 < 8; fi89++ {
                        if true {
                            var v90 uint8 = v88
                            if ((("foo bar" + "\n\t") < ("foo bar" + "a")) || (!(127 < 9223372036854775807))) {
                                var v91 bool = false
                                var v92 uint8 = v88
                                _ = v91
                                _ = v92
                            }
                            fmt.Println(v6, v90, v2)
                            _ = v90
                        }
                        v93 := GenS10{F0: ((v6 | 42) < (v6 % (((v6 + (v6 / ((-525) | 1)))) | 1))), F1: (v2 ^ uint64(v88))}
                        _ = fi89
                        _ = v93
                    }
                default:
                    v2 = v2
                }
            }
            _ = fi85
            _ = v88
        }
        if false {
            if false {
                v94 := GenS13{F0: v2, F1: 1}
                _ = v94
            } else {
                if true {
                    v2 += 0
                    for fi95 := 0; fi95 < 3; fi95++ {
                        var v96 uint32 = 65535
                        v97 := GenS14{F0: 42, F1: 91, F2: 275}
                        _ = fi95
                        _ = v96
                        _ = v97
                    }
                    v2 += (v2 - (v2 + (v2 / (((^(v2))) | 1))))
                }
                fmt.Println(v6, v6)
            }
            var v98 float64 = (float64(1.5) + 0.1)
            v99 := GenS15{F0: 125, F1: 127}
            _ = v98
            _ = v99
        } else {
            v6 = 1
        }
    }
    if false {
        if (0.0 != (((0.1 - 1e10) - (0.0 / (1.5 * 0.1))) - 1.5)) {
            switch 215 {
            case 4294967295:
                var v100 bool = true
                _ = v100
            case -511:
                v6 = v6
            default:
                var v101 float32 = ((((0.1 * float32(3.14159)) - (3.14159 - 0.1)) * 1e10) + (1e10 / -2.25))
                _ = v101
            }
            if (false && (true && ("a" < ((("hello" + "foo bar") + "") + "")))) {
                v1 = append(v1, "foo bar")
            }
        } else {
            switch 0 {
            case 795:
                for fi102 := 0; fi102 < 10; fi102++ {
                    v103 := GenS6{F0: (732 < 18)}
                    if false {
                        var v104 float32 = 1e10
                        _ = v104
                    }
                    _ = fi102
                    _ = v103
                }
            case -9223372036854775808:
                for fi105 := 0; fi105 < 20; fi105++ {
                    if true {
                        var v106 bool = ((v2 << ((42) & 63)) <= v2)
                        var v107 float64 = -2.25
                        fmt.Println(v106, v2)
                        _ = v106
                        _ = v107
                    }
                    _ = len(v1)
                    _ = fi105
                }
            }
        }
    } else {
        for fi108 := 0; fi108 < 19; fi108++ {
            v109 := GenS16{F0: 42}
            if (1 > -9223372036854775808) {
                for fi110 := 0; fi110 < 9; fi110++ {
                    if len(v4) > 0 {
                        v4[0] = -1
                    }
                    v2 = 0
                    _ = fi110
                }
                v6 += -55
            } else {
                v111 := GenS4{F0: 42}
                _ = v111
            }
            delete(v3, 4294967295)
            _ = fi108
            _ = v109
        }
        v2 = 567
    }
    fmt.Println(v6)
    var v112 int = 2147483647
    switch ("" + (("foo bar" + "hello") + "hello")) {
    case "hello":
        for fi113 := 0; fi113 < 20; fi113++ {
            var v114 int8 = -128
            v6 += -32768
            _ = fi113
            _ = v114
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v112:", v112)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v112
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
