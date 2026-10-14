package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 bool; F1 int; F2 uint16 }
type GenS6 struct { F0 int64; F1 string; F2 uint16 }
type GenS7 struct { F0 int; F1 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int; F1 bool; F2 uint }
type GenS10 struct { F0 int64; F1 int; F2 int64 }
type GenS11 struct { F0 int; F1 int32 }
type GenS12 struct { F0 uint; F1 uint16; F2 string }
type GenS13 struct { F0 int; F1 uint32; F2 bool }
type GenS14 struct { F0 string; F1 bool; F2 uint64 }
type GenS15 struct { F0 int8; F1 int8 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 bool; F1 uint64 }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 int8; F1 int64; F2 uint32 }
type GenS20 struct { F0 uint64; F1 int }
type GenS21 struct { F0 int32 }
func genFn0(p0 uint16) uint64 {
    v7 := []int{-815}
    for fi8 := 0; fi8 < 10; fi8++ {
        if false {
            v9 := GenS1{F0: (!false), F1: 40}
            var v10 int32 = 1
            var v11 int8 = -99
            _ = v10
            _ = v11
            _ = v9
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        p0 += p0
        _ = fi12
    }
    _ = p0
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int32) uint {
    v13 := GenS1{F0: true, F1: 100}
    fmt.Println(p0)
    if (392 > 494) {
        var v14 int8 = 100
        fmt.Println(v14, v14)
        var v15 string = ("a" + "foo bar")
        _ = v14
        _ = v15
    } else {
        v16 := GenS1{F0: (2 <= 444), F1: 42}
        p0 = p0
        _ = v16
    }
    _ = p0
    _ = v13
    return 255
}

func main() {
    v1 := []uint32{2147483647}
    var v2 bool = true
    var v3 float64 = 1e10
    var v4 int32 = -2147483648
    var v5 int64 = 42
    v6 := make(map[string]int)
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := GenS2{F0: 42}
        _ = fi17
        _ = v18
    }
    fmt.Println(v2, v5)
    for fi19 := 0; fi19 < 13; fi19++ {
        fmt.Println(fi19, v3)
        _ = len(v1)
        v2 = v2
        _ = fi19
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 uint = 4294967295
        _ = fi20
        _ = v21
    }
    if (0 >= -1) {
        var v22 int32 = 42
        switch 32767 {
        case -573:
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := GenS3{F0: 199, F1: 32767}
                v2 = true
                _ = fi23
                _ = v24
            }
        case -1:
            v25 := GenS4{F0: genFn1(v4), F1: 128}
            _ = v25
        case 893:
            var v26 float32 = float32((-2.25 - (0.1 - v3)))
            _ = v26
        }
        _ = v22
    }
    switch "\n\t" {
    case "":
        if ((!(v2 || ((v5 | v5) < (v5 / ((-1) | 1))))) || v2) {
            var v27 float64 = v3
            var v28 uint64 = 2
            for fi29 := 0; fi29 < 5; fi29++ {
                if v2 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch fi29 {
                    case 2147483647:
                        for fi30 := 0; fi30 < 18; fi30++ {
                            v31 := GenS1{F0: true, F1: 0}
                            fmt.Println(v4, v2, v28)
                            fmt.Println(v2, v27, v28)
                            _ = fi30
                            _ = v31
                        }
                    case -2147483648:
                        fmt.Println(v5)
                    case -9223372036854775808:
                        if (v2 && v2) {
                            for fi32 := 0; fi32 < 20; fi32++ {
                                switch fi29 {
                                case -312:
                                    v33 := GenS5{F0: false, F1: 42, F2: 30}
                                    _ = v33
                                case -1:
                                    v34 := GenS3{F0: 0, F1: -611}
                                    fmt.Println(v4, v2, v3)
                                    _ = v34
                                default:
                                    v27 = ((float64(1.5) + v3) + ((1.5 + (v27 + float64(3.14159))) * (0.1 / v3)))
                                }
                                v28 += genFn0(255)
                                fmt.Println(fi32, v5, v27)
                                var v35 string = ""
                                v36 := GenS6{F0: (v5 >> ((genFn1(-1)) & 63)), F1: (("" + "foo bar") + (((((v35 + v35) + ((v35 + v35) + ((v35 + v35) + ""))) + v35) + "a") + v35)), F2: 255}
                                _ = fi32
                                _ = v35
                                _ = v36
                            }
                            if v2 {
                                for fi37 := 0; fi37 < 11; fi37++ {
                                    var v38 int32 = v4
                                    _ = fi37
                                    _ = v38
                                }
                            } else {
                                v39 := GenS6{F0: (v5 - v5), F1: "\n\t", F2: 856}
                                switch "\n\t" {
                                case "\n\t":
                                    v40 := GenS7{F0: 4294967295, F1: int8(v4)}
                                    _ = v40
                                case "hello":
                                    var v41 bool = v2
                                    _ = v41
                                default:
                                    for fi42 := 0; fi42 < 3; fi42++ {
                                        v3 += v27
                                        switch ("foo bar" + "") {
                                        case "\n\t":
                                            var v43 uint = 255
                                            _ = v43
                                        case "a":
                                            switch (-(fi42)) {
                                            case -9223372036854775808:
                                                v44 := GenS8{F0: 2}
                                                _ = v44
                                            case 2147483647:
                                                fmt.Println(v28, v28, v3)
                                            case -1:
                                                switch (-(fi29)) {
                                                case 924:
                                                    if (!(48 != 2)) {
                                                        switch fi42 {
                                                        case -9223372036854775808:
                                                            switch 677 {
                                                            case -9223372036854775808:
                                                                v45 := GenS9{F0: fi42, F1: false, F2: 0}
                                                                _ = v45
                                                            case -1:
                                                                v4 += (v4 ^ (-(v4)))
                                                                var v46 uint = 4294967295
                                                                v47 := GenS10{F0: (v5 ^ (v5 * (v5 - (^(v5))))), F1: fi42, F2: v5}
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                        }
                                                    } else {
                                                        if (480 < 0) {
                                                            switch fi29 {
                                                            case 4294967295:
                                                                var v48 float64 = 0.1
                                                                _ = v48
                                                            case 9223372036854775807:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            case 127:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        } else {
                                                            v2 = v2
                                                            if len(v1) > 0 {
                                                                v1[0] = 1
                                                            }
                                                            var v49 uint8 = 0
                                                            _ = v49
                                                        }
                                                    }
                                                case -682:
                                                    delete(v6, "hello")
                                                default:
                                                    v50 := GenS2{F0: (v28 % ((v28) | 1))}
                                                    _ = v50
                                                }
                                                if (v28 != v28) {
                                                    if true {
                                                        v3 += v27
                                                        v4 += (v4 & (v4 | int32(v5)))
                                                        v51 := GenS6{F0: -1, F1: (("" + "foo bar") + "a"), F2: 128}
                                                        _ = v51
                                                    }
                                                }
                                                if (0 >= 252) {
                                                    var v52 int32 = 2147483647
                                                    fmt.Println(v4, fi29, v3)
                                                    _ = v52
                                                }
                                            default:
                                                v53 := GenS9{F0: fi42, F1: v2, F2: 65535}
                                                _ = v53
                                            }
                                        case "hello":
                                            if ((3.14159 / -2.25) >= 0.0) {
                                                switch int(v5) {
                                                case 9223372036854775807:
                                                    v4 += -1
                                                case -237:
                                                    var v54 float64 = 0.0
                                                    _ = v54
                                                case 1:
                                                    if (1 <= 127) {
                                                        v2 = (-128 > -128)
                                                        switch ("hello" + "hello") {
                                                        case "a":
                                                            v4 += v4
                                                        case "hello":
                                                            if (65535 == 2) {
                                                                fmt.Println(v28)
                                                                v27 += float64(1e10)
                                                                if v2 {
                                                                    fmt.Println(v3)
                                                                    v27 = 1.5
                                                                }
                                                                if true {
                                                                    fmt.Println(v5, v27, v4)
                                                                    fmt.Println(fi42, v27, v2)
                                                                    var v55 uint32 = 2
                                                                    _ = v55
                                                                } else {
                                                                    var v56 float64 = 0.1
                                                                    v3 = (v27 - 0.1)
                                                                    _ = v56
                                                                }
                                                                fmt.Println(v5)
                                                            } else {
                                                                v57 := GenS7{F0: int(v28), F1: -124}
                                                                _ = v57
                                                            }
                                                        }
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                default:
                                                    var v58 string = ("" + "")
                                                    _ = v58
                                                }
                                            } else {
                                                switch "hello" {
                                                case "":
                                                    fmt.Println(v27)
                                                }
                                                fmt.Println(v5, v5, v5)
                                                var v59 float64 = v3
                                                switch fi29 {
                                                case -437:
                                                    delete(v6, "\n\t")
                                                case 2147483647:
                                                    v2 = v2
                                                case 58:
                                                    v60 := GenS4{F0: 18446744073709551615, F1: (fi42 << ((511) & 63))}
                                                    _ = v60
                                                default:
                                                    var v61 bool = v2
                                                    _ = v61
                                                }
                                                v62 := GenS11{F0: -27, F1: v4}
                                                _ = v59
                                                _ = v62
                                            }
                                            if v2 {
                                                switch (fi42 ^ 32768) {
                                                case 32767:
                                                    var v63 string = "hello"
                                                    _ = v63
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = 480
                                                    }
                                                }
                                                if ((^(fi42)) >= 128) {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    fmt.Println(v28, v27, v4)
                                                    switch "hello" {
                                                    case "a":
                                                        var v64 int32 = v4
                                                        _ = v64
                                                    case "":
                                                        v1 = append(v1, 586)
                                                    }
                                                    if false {
                                                        v27 = float64(3.14159)
                                                    } else {
                                                        var v65 uint32 = 277
                                                        v66 := GenS12{F0: 255, F1: 1, F2: "foo bar"}
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                }
                                            }
                                            var v67 uint = 18446744073709551615
                                            _ = v67
                                        }
                                        _ = fi42
                                    }
                                }
                                v3 = (float64(1.5) * ((v27 - 3.14159) / v3))
                                _ = v39
                            }
                            fmt.Println(v27, v27)
                        }
                    }
                    var v68 int = (fi29 >> ((974) & 63))
                    _ = v68
                } else {
                    var v69 int64 = (v5 ^ v5)
                    v3 = (v3 / 3.14159)
                    v3 = v3
                    _ = v69
                }
                if v2 {
                    fmt.Println(fi29)
                    if ((v5 >> ((893) & 63)) <= (v5 | -777)) {
                        v5 -= v5
                    } else {
                        var v70 int = (fi29 >> ((141) & 63))
                        fmt.Println(v70, v27, v70)
                        _ = len(v6)
                        _ = v70
                    }
                    fmt.Println(v3)
                }
                v1 = append(v1, 2)
                _ = fi29
            }
            _ = v27
            _ = v28
        }
    case "\n\t":
        var v71 int16 = 42
        _ = v71
    case "a":
        switch 1 {
        case 42:
            for fi72 := 0; fi72 < 18; fi72++ {
                _ = len(v1)
                _ = fi72
            }
        }
    }
    var v73 int8 = -128
    v74 := GenS2{F0: 18446744073709551615}
    v2 = v2
    if true {
        for fi75 := 0; fi75 < 14; fi75++ {
            v3 = (0.1 * ((-2.25 * (float64(float32(v3)) * 0.1)) - (v3 - v3)))
            _ = fi75
        }
    } else {
        var v76 uint64 = 0
        v77 := GenS13{F0: int(v73), F1: 0, F2: v2}
        _ = v76
        _ = v77
    }
    for fi78 := 0; fi78 < 8; fi78++ {
        var v79 bool = (v3 <= v3)
        var v80 uint8 = 255
        _ = fi78
        _ = v79
        _ = v80
    }
    if (193 != 42) {
        for fi81 := 0; fi81 < 18; fi81++ {
            var v82 int16 = 32767
            for fi83 := 0; fi83 < 2; fi83++ {
                var v84 uint8 = 40
                v85 := GenS14{F0: "\n\t", F1: (true && (1 > genFn0(42))), F2: 2}
                var v86 uint8 = (^(v84))
                _ = fi83
                _ = v84
                _ = v85
                _ = v86
            }
            _ = fi81
            _ = v82
        }
        if true {
            v87 := GenS15{F0: -126, F1: (v73 * v73)}
            if ((v2 || v2) && (981 > int16(v73))) {
                if v2 {
                    for fi88 := 0; fi88 < 18; fi88++ {
                        v3 += ((1.5 - v3) - v3)
                        if v2 {
                            delete(v6, "foo bar")
                            switch "" {
                            case "a":
                                var v89 bool = true
                                _ = v89
                            case "":
                                fmt.Println(v3, v5)
                            case "hello":
                                for fi90 := 0; fi90 < 20; fi90++ {
                                    v4 += v4
                                    v73 = v73
                                    delete(v6, "")
                                    _ = fi90
                                }
                            }
                            if ("a" <= "") {
                                delete(v6, ("" + "\n\t"))
                                v91 := v6["a"]
                                _ = v91
                            } else {
                                v73 = v73
                                v2 = v2
                            }
                        } else {
                            delete(v6, "foo bar")
                            v2 = (!false)
                            for fi92 := 0; fi92 < 20; fi92++ {
                                var v93 string = ((("hello" + "a") + ("" + (((("hello" + "hello") + "a") + "foo bar") + "hello"))) + "")
                                if ((-2.25 - (float64(3.14159) - (v3 * 1.5))) != float64((0.0 + 1.5))) {
                                    var v94 int64 = v5
                                    var v95 int = (^(fi92))
                                    v2 = (1 >= -22)
                                    v96 := GenS2{F0: 42}
                                    _ = v94
                                    _ = v95
                                    _ = v96
                                }
                                _ = fi92
                                _ = v93
                            }
                        }
                        _ = fi88
                    }
                }
                if (170 == 1) {
                    if v2 {
                        var v97 float32 = -2.25
                        switch 128 {
                        case 1:
                            var v98 int32 = (v4 ^ v4)
                            _ = v98
                        }
                        v2 = (4294967295 < 65535)
                        _ = v97
                    } else {
                        fmt.Println(v2)
                        var v99 int16 = -970
                        _ = v99
                    }
                    for fi100 := 0; fi100 < 4; fi100++ {
                        if ((-2.25 + 3.14159) > 3.14159) {
                            switch "\n\t" {
                            case "":
                                v3 -= v3
                            case "a":
                                v3 = 1.5
                            default:
                                fmt.Println(fi100, v2, v5)
                            }
                            if (568 > 203) {
                                var v101 uint16 = 1
                                var v102 int16 = 1000
                                _ = v101
                                _ = v102
                            } else {
                                v73 = v73
                            }
                        } else {
                            var v103 int8 = v73
                            switch fi100 {
                            case -201:
                                v73 = v73
                            case -2147483648:
                                switch fi100 {
                                case 9223372036854775807:
                                    if (!(false && true)) {
                                        for fi104 := 0; fi104 < 3; fi104++ {
                                            fmt.Println(fi104)
                                            _ = fi104
                                        }
                                        for fi105 := 0; fi105 < 16; fi105++ {
                                            v2 = (v2 && false)
                                            if (true || (1 > (-(v5)))) {
                                                var v106 uint8 = 0
                                                v5 = (^(v5))
                                                _ = v106
                                            } else {
                                                switch 2147483647 {
                                                case 127:
                                                    delete(v6, ("hello" + ("hello" + "hello")))
                                                case -150:
                                                    v103 = v73
                                                case -950:
                                                    v73 += int8(fi105)
                                                default:
                                                    switch fi105 {
                                                    case 32768:
                                                        if (!((v4 ^ v4) < (v4 | (-(v4))))) {
                                                            var v107 float64 = 3.14159
                                                            var v108 float32 = (1.5 + (((((((0.0 * ((0.0 + (0.0 * 1e10)) * -2.25)) - 0.0) * float32(3.14159)) / ((3.14159 - -2.25) / 0.1)) / ((0.1 / 1.0) + 0.1)) / 0.1) + 1.5))
                                                            switch ("a" + ("foo bar" + "")) {
                                                            case "foo bar":
                                                                if v2 {
                                                                    v5 = (v5 % ((42) | 1))
                                                                }
                                                            case "hello":
                                                                switch "hello" {
                                                                case "a":
                                                                    v2 = v2
                                                                default:
                                                                    v109 := GenS16{F0: 128}
                                                                    _ = v109
                                                                }
                                                            }
                                                            _ = v107
                                                            _ = v108
                                                        } else {
                                                            var v110 uint64 = 307
                                                            _ = v110
                                                        }
                                                    default:
                                                        fmt.Println(v2)
                                                    }
                                                }
                                                v4 = int32(v5)
                                            }
                                            if (!v2) {
                                                v111 := GenS8{F0: 41}
                                                v73 = v73
                                                _ = len(v1)
                                                fmt.Println(fi105)
                                                var v112 int64 = v5
                                                v4 = 0
                                                v5 = (-(v5))
                                                v113 := GenS17{F0: (1.5 < v3), F1: 1}
                                                var v114 uint = 4294967295
                                                _ = v111
                                                _ = v112
                                                _ = v113
                                                _ = v114
                                            } else {
                                                v115 := GenS18{F0: 1}
                                                v103 += v103
                                                _ = v115
                                            }
                                            _ = fi105
                                        }
                                        v73 = v103
                                    } else {
                                        if ((v5 + v5) > v5) {
                                            if (fi100 >= -874) {
                                                for fi116 := 0; fi116 < 9; fi116++ {
                                                    var v117 float64 = ((v3 + 1e10) / v3)
                                                    var v118 int8 = (-(v73))
                                                    _ = fi116
                                                    _ = v117
                                                    _ = v118
                                                }
                                                v4 = (v4 << ((0) & 63))
                                            } else {
                                                v2 = true
                                                v2 = v2
                                                v119 := GenS17{F0: (42 < 600), F1: 0}
                                                _ = v119
                                            }
                                            var v120 int = (fi100 / ((-379) | 1))
                                            _ = v120
                                        } else {
                                            v121 := GenS5{F0: v2, F1: (-(fi100)), F2: 32767}
                                            for fi122 := 0; fi122 < 12; fi122++ {
                                                var v123 uint32 = 42
                                                v124 := GenS8{F0: 174}
                                                v2 = (v4 > (v4 << ((1) & 63)))
                                                var v125 int16 = 1
                                                _ = fi122
                                                _ = v123
                                                _ = v124
                                                _ = v125
                                            }
                                            if v2 {
                                                if (v4 > (v4 % ((2147483647) | 1))) {
                                                    v103 = (v73 >> ((4294967295) & 63))
                                                    var v126 int32 = v4
                                                    _ = v126
                                                } else {
                                                    v127 := GenS15{F0: (v73 + 1), F1: int8(v5)}
                                                    var v128 uint64 = 2
                                                    _ = v127
                                                    _ = v128
                                                }
                                                v73 -= (v103 / ((-1) | 1))
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v129 := GenS6{F0: (v5 - (^(v5))), F1: "a", F2: 1}
                                                    _ = v129
                                                case "a":
                                                    _ = len(v1)
                                                }
                                            }
                                            _ = v121
                                        }
                                    }
                                case -1:
                                    v103 = (v103 - v103)
                                case 32768:
                                    v130 := GenS17{F0: v2, F1: 18446744073709551615}
                                    _ = v130
                                default:
                                    v73 -= (v73 & (v73 & 1))
                                }
                            case 1:
                                var v131 uint64 = 251
                                _ = v131
                            }
                            if false {
                                var v132 string = "\n\t"
                                v5 = (v5 ^ (v5 & 42))
                                v132 = "a"
                                _ = v132
                            } else {
                                v133 := v6["hello"]
                                v2 = v2
                                _ = v133
                            }
                            _ = v103
                        }
                        _ = fi100
                    }
                    var v134 int8 = (v73 % ((1) | 1))
                    _ = v134
                } else {
                    var v135 string = "\n\t"
                    _ = len(v1)
                    _ = v135
                }
                switch -9223372036854775808 {
                case -631:
                    var v136 bool = true
                    _ = v136
                case 0:
                    v137 := v6[("hello" + "foo bar")]
                    _ = v137
                default:
                    v4 -= (v4 / (((v4 << ((255) & 63))) | 1))
                }
            }
            _ = v87
        } else {
            var v138 int8 = v73
            _ = v138
        }
    } else {
        if v2 {
            switch ((((("\n\t" + "a") + "foo bar") + ("" + ((("foo bar" + "hello") + ("\n\t" + "a")) + ""))) + "foo bar") + ("hello" + "a")) {
            case "\n\t":
                fmt.Println(v3)
            case "a":
                fmt.Println(v3, v73)
            default:
                var v139 int8 = v73
                _ = v139
            }
            v73 = -1
            var v140 uint16 = 1
            _ = v140
        }
        v5 = (v5 & 1)
    }
    switch ("foo bar" + ("\n\t" + "foo bar")) {
    case "hello":
        if false {
            if v2 {
                switch -421 {
                case 4294967295:
                    var v141 uint16 = 1
                    _ = v141
                case 9223372036854775807:
                    if v2 {
                        fmt.Println(v2, v2, v2)
                        v73 -= (v73 | (v73 % ((42) | 1)))
                    }
                case 2147483647:
                    v142 := GenS19{F0: v73, F1: 0, F2: 42}
                    _ = v142
                default:
                    for fi143 := 0; fi143 < 12; fi143++ {
                        switch fi143 {
                        case 135:
                            if true {
                                var v144 float64 = (float64(0.1) * (float64(1.5) - v3))
                                v145 := GenS4{F0: 4294967295, F1: (fi143 | (fi143 << ((18446744073709551615) & 63)))}
                                v6[("foo bar" + "a")] = fi143
                                _ = v144
                                _ = v145
                            }
                        default:
                            for fi146 := 0; fi146 < 3; fi146++ {
                                fmt.Println(fi146, fi146, v5)
                                var v147 int32 = v4
                                fmt.Println(v5)
                                _ = fi146
                                _ = v147
                            }
                        }
                        _ = fi143
                    }
                }
            }
            switch 609 {
            case 32768:
                var v148 float32 = (3.14159 / -2.25)
                _ = v148
            case 578:
                v4 = v4
            default:
                v2 = (!(!v2))
            }
        }
    case "a":
        switch (("" + "") + "\n\t") {
        case "foo bar":
            for fi149 := 0; fi149 < 17; fi149++ {
                var v150 uint8 = 68
                _ = fi149
                _ = v150
            }
        }
    default:
        v3 -= ((float64(0.1) - 3.14159) + float64((3.14159 / 3.14159)))
    }
    for fi151 := 0; fi151 < 20; fi151++ {
        v73 -= int8(v5)
        var v152 uint32 = 224
        _ = fi151
        _ = v152
    }
    var v153 float64 = 0.0
    switch ("a" + "") {
    case "foo bar":
        v4 -= (v4 / (((-(v4))) | 1))
    case "hello":
        for fi154 := 0; fi154 < 8; fi154++ {
            var v155 bool = ((1e10 < ((1.5 / 0.1) + 1.5)) || v2)
            v73 = (v73 >> ((65535) & 63))
            var v156 int64 = 320
            _ = fi154
            _ = v155
            _ = v156
        }
    case "\n\t":
        v157 := GenS20{F0: 18446744073709551615, F1: 388}
        _ = v157
    default:
        v158 := GenS5{F0: v2, F1: 574, F2: 1}
        _ = v158
    }
    v159 := GenS21{F0: 360}
    var v160 uint64 = 758
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v153:", v153)
    fmt.Println("v160:", v160)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v73:", v73)
    _ = v1
    _ = v153
    _ = v159
    _ = v160
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
}
