package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 uint; F1 uint }
func genFn0() string {
    switch 73 {
    case 1:
        switch 42 {
        case 128:
            v7 := GenS1{F0: 18446744073709551615}
            _ = v7
        default:
            v8 := []int{-91, -115, 1}
            _ = v8
        }
    default:
        if true {
            var v9 bool = false
            v9 = true
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := []int{}
                v9 = v9
                _ = fi10
                _ = v11
            }
            _ = v9
        }
    }
    var v12 float64 = 1.5
    v12 += 0.0
    v13 := []int{-2147483648, 32768}
    v12 += v12
    _ = v12
    _ = v13
    return "hello"
}
func genFn1() uint {
    var v14 int8 = 12
    var v15 int8 = v14
    switch 127 {
    case -9223372036854775808:
        v16 := []int{-649}
        _ = v16
    default:
        switch "foo bar" {
        case "foo bar":
            if (3.14159 <= float64(0.0)) {
                var v17 string = "\n\t"
                v18 := []int{-288, 382, -423, 32767}
                var v19 int = -374
                _ = v17
                _ = v18
                _ = v19
            }
        case "\n\t":
            v20 := []int{339, 127, 2147483647, 4294967295}
            _ = v20
        case "":
            var v21 uint = 964
            _ = v21
        default:
            if (1e10 >= 1.5) {
                if (42 != 1000) {
                    switch "\n\t" {
                    case "a":
                        v14 = v15
                    case "\n\t":
                        fmt.Println(v15)
                    case "foo bar":
                        if ((615 >= 42) && true) {
                            for fi22 := 0; fi22 < 17; fi22++ {
                                var v23 bool = true
                                if (0 >= 255) {
                                    for fi24 := 0; fi24 < 9; fi24++ {
                                        v15 = v15
                                        _ = fi24
                                    }
                                    v25 := []int{0, 149}
                                    if v23 {
                                        for fi26 := 0; fi26 < 11; fi26++ {
                                            v14 -= (v14 & v15)
                                            _ = fi26
                                        }
                                        fmt.Println(v15, fi22)
                                    }
                                    _ = v25
                                } else {
                                    fmt.Println(v14)
                                }
                                v27 := GenS1{F0: 1}
                                _ = fi22
                                _ = v23
                                _ = v27
                            }
                            var v28 bool = false
                            var v29 float64 = 1e10
                            _ = v28
                            _ = v29
                        } else {
                            fmt.Println(v14, v14)
                            fmt.Println(v15, v14)
                        }
                    default:
                        if true {
                            var v30 uint64 = 2
                            v15 += v14
                            v15 -= (v14 % ((v15) | 1))
                            _ = v30
                        }
                    }
                    if false {
                        v14 = 0
                    } else {
                        for fi31 := 0; fi31 < 18; fi31++ {
                            v15 = (-(v14))
                            _ = fi31
                        }
                        v32 := GenS1{F0: 872}
                        _ = v32
                    }
                } else {
                    for fi33 := 0; fi33 < 1; fi33++ {
                        for fi34 := 0; fi34 < 19; fi34++ {
                            v15 = (v14 | v15)
                            if false {
                                var v35 bool = false
                                var v36 uint16 = 2
                                v37 := GenS2{F0: 632}
                                var v38 int32 = -1
                                var v39 uint = 883
                                _ = v35
                                _ = v36
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            v40 := GenS3{F0: 885}
                            _ = fi34
                            _ = v40
                        }
                        var v41 uint = 527
                        for fi42 := 0; fi42 < 15; fi42++ {
                            switch fi42 {
                            case 9223372036854775807:
                                v43 := []int{(fi42 % (((fi42 ^ (fi42 | (-(fi33))))) | 1)), 32768, 4294967295}
                                _ = v43
                            }
                            _ = fi42
                        }
                        _ = fi33
                        _ = v41
                    }
                    v44 := GenS4{F0: 2147483647, F1: 987}
                    _ = v44
                }
            } else {
                switch "\n\t" {
                case "foo bar":
                    var v45 uint32 = 1
                    _ = v45
                default:
                    var v46 int64 = -723
                    _ = v46
                }
                v47 := []int{-2147483648, 127, 128}
                _ = v47
            }
        }
    }
    if (18446744073709551615 <= 18446744073709551615) {
        var v48 uint = 1
        _ = v48
    }
    _ = v14
    _ = v15
    return 18446744073709551615
}

func main() {
    var v1 int16 = -32768
    var v2 int8 = -1
    var v3 uint16 = 65535
    var v4 string = "foo bar"
    v5 := make(map[int]bool)
    var v6 int8 = v2
    delete(v5, 2147483647)
    fmt.Println(v3, v3, v2)
    var v49 uint8 = 0
    for fi50 := 0; fi50 < 10; fi50++ {
        var v51 int64 = 605
        _ = fi50
        _ = v51
    }
    fmt.Println(v49)
    fmt.Println(v49, v6)
    for fi52 := 0; fi52 < 7; fi52++ {
        v2 = (v2 - (v6 >> ((genFn1()) & 63)))
        for fi53 := 0; fi53 < 19; fi53++ {
            if true {
                v54 := GenS4{F0: fi52, F1: (v3 - (v3 / (((^(v3))) | 1)))}
                v4 = ("hello" + (v4 + v4))
                var v55 uint = 558
                _ = v54
                _ = v55
            }
            if false {
                v5[fi52] = false
                fmt.Println(v1, v49)
                v56 := GenS1{F0: 42}
                if false {
                    var v57 int64 = -1
                    _ = v57
                } else {
                    if (!true) {
                        v2 = (v6 + -107)
                    } else {
                        v1 = 32767
                        fmt.Println(v49, v4)
                        fmt.Println(fi52, v6, v3)
                    }
                    v58 := []int{}
                    fmt.Println(fi52)
                    _ = len(v5)
                    _ = v58
                }
                _ = v56
            }
            var v59 uint32 = 1
            _ = fi53
            _ = v59
        }
        v60 := GenS1{F0: 18446744073709551615}
        _ = fi52
        _ = v60
    }
    if (0 < 900) {
        if true {
            if true {
                if false {
                    switch (v4 + v4) {
                    case "a":
                        fmt.Println(v3, v1, v6)
                    }
                    v61 := []int{int(v49)}
                    switch -1 {
                    case 128:
                        var v62 uint64 = 777
                        _ = v62
                    case 1:
                        if len(v61) >= 2 {
                            v61 = v61[1:len(v61)]
                        }
                    }
                    _ = v61
                } else {
                    if (65535 > 0) {
                        v63 := GenS1{F0: 42}
                        if false {
                            for fi64 := 0; fi64 < 6; fi64++ {
                                fmt.Println(fi64)
                                _ = fi64
                            }
                            for fi65 := 0; fi65 < 2; fi65++ {
                                v6 = v2
                                v66 := GenS5{F0: 1, F1: fi65}
                                _ = fi65
                                _ = v66
                            }
                            var v67 uint32 = 2147483648
                            _ = v67
                        }
                        _ = v63
                    }
                    var v68 int32 = 42
                    if ((v68 | (v68 ^ v68)) < -2147483648) {
                        v69 := GenS6{F0: (v1 > 828), F1: false}
                        var v70 string = v4
                        var v71 uint32 = 2147483647
                        _ = v69
                        _ = v70
                        _ = v71
                    }
                    _ = v68
                }
                fmt.Println(v3, v6)
                for fi72 := 0; fi72 < 3; fi72++ {
                    v73 := GenS5{F0: 996, F1: fi72}
                    _ = fi72
                    _ = v73
                }
            }
            switch -1 {
            case -924:
                var v74 int16 = v1
                _ = v74
            case 446:
                if (2147483648 <= 1) {
                    v75 := GenS7{F0: 1, F1: 2}
                    _ = v75
                }
            case 0:
                var v76 int32 = 2147483647
                _ = v76
            }
            for fi77 := 0; fi77 < 12; fi77++ {
                switch fi77 {
                case -1:
                    switch fi77 {
                    case 2147483647:
                        switch "" {
                        case "a":
                            v78 := GenS7{F0: 0, F1: (v3 << ((0) & 63))}
                            _ = v78
                        case "hello":
                            for fi79 := 0; fi79 < 14; fi79++ {
                                fmt.Println(fi77)
                                v1 -= v1
                                var v80 int = (fi79 & fi77)
                                _ = fi79
                                _ = v80
                            }
                        case "\n\t":
                            v81 := []int{fi77}
                            _ = v81
                        }
                    case 128:
                        v82 := []int{fi77, 357, 0, 127}
                        _ = v82
                    }
                default:
                    var v83 uint = genFn1()
                    _ = v83
                }
                fmt.Println(v3, fi77)
                v84 := GenS8{F0: 42, F1: 310}
                _ = fi77
                _ = v84
            }
        } else {
            v3 -= (^(v3))
            var v85 bool = true
            switch ("foo bar" + ((v4 + "") + ("\n\t" + ((("a" + v4) + (("\n\t" + v4) + ((("foo bar" + v4) + (v4 + v4)) + v4))) + (v4 + ((((("\n\t" + "foo bar") + "foo bar") + (v4 + "hello")) + v4) + (v4 + ("" + v4)))))))) {
            case "":
                for fi86 := 0; fi86 < 7; fi86++ {
                    v87 := v5[(fi86 << ((genFn1()) & 63))]
                    _ = len(v5)
                    var v88 uint64 = 853
                    _ = fi86
                    _ = v87
                    _ = v88
                }
            case "hello":
                switch v4 {
                case "":
                    if ((0.0 * 0.1) != float64(0.1)) {
                        if v85 {
                            fmt.Println(v3, v85, v49)
                            switch "\n\t" {
                            case "a":
                                var v89 int = 128
                                _ = v89
                            }
                            var v90 int32 = 1
                            _ = v90
                        } else {
                            var v91 int16 = (v1 - (v1 % (((v1 << ((4294967295) & 63))) | 1)))
                            _ = v91
                        }
                        fmt.Println(v1, v85, v49)
                    } else {
                        var v92 string = v4
                        _ = v92
                    }
                case "hello":
                    v4 = v4
                }
            }
            _ = v85
        }
    }
    switch "a" {
    case "":
        if (false || true) {
            switch -309 {
            case 538:
                v49 = (v49 | v49)
            case 895:
                var v93 int64 = 1
                _ = v93
            }
            if true {
                if true {
                    v94 := v5[-2147483648]
                    _ = v94
                }
                if false {
                    v95 := GenS8{F0: 42, F1: 754}
                    if ((-2.25 < (float32((-2.25 - 1.5)) * 3.14159)) || (228 > 65535)) {
                        if (205 > 876) {
                            var v96 int16 = -447
                            var v97 uint64 = 4294967295
                            _ = v96
                            _ = v97
                        }
                        v4 = "foo bar"
                    }
                    for fi98 := 0; fi98 < 18; fi98++ {
                        switch (("" + "\n\t") + v4) {
                        case "\n\t":
                            fmt.Println(fi98)
                        case "a":
                            fmt.Println(fi98, v49, v2)
                        default:
                            switch (fi98 + int(v3)) {
                            case 907:
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(v6, v49, v6)
                                }
                            case 128:
                                fmt.Println(v3)
                            }
                        }
                        var v99 int16 = v1
                        _ = fi98
                        _ = v99
                    }
                    _ = v95
                } else {
                    switch "hello" {
                    case "hello":
                        fmt.Println(v4)
                    case "a":
                        if ((42 != int(v3)) && true) {
                            switch (v4 + "") {
                            case "":
                                v3 += v3
                            default:
                                for fi100 := 0; fi100 < 3; fi100++ {
                                    v4 = (("a" + v4) + "a")
                                    fmt.Println(v2, v1)
                                    _ = fi100
                                }
                            }
                            delete(v5, 32768)
                            for fi101 := 0; fi101 < 19; fi101++ {
                                fmt.Println(v2, v3)
                                _ = fi101
                            }
                        }
                    case "foo bar":
                        for fi102 := 0; fi102 < 5; fi102++ {
                            var v103 int = (^(fi102))
                            v104 := GenS2{F0: 0}
                            _ = fi102
                            _ = v103
                            _ = v104
                        }
                    }
                    fmt.Println(v4)
                    for fi105 := 0; fi105 < 3; fi105++ {
                        switch v4 {
                        case "":
                            for fi106 := 0; fi106 < 13; fi106++ {
                                if false {
                                    fmt.Println(v6, v6)
                                } else {
                                    v107 := []int{}
                                    _ = v107
                                }
                                v108 := GenS4{F0: 0, F1: (^(v3))}
                                _ = fi106
                                _ = v108
                            }
                        case "foo bar":
                            v6 += (v6 & 1)
                        }
                        var v109 int32 = 1
                        _ = fi105
                        _ = v109
                    }
                }
            } else {
                for fi110 := 0; fi110 < 15; fi110++ {
                    var v111 float64 = 1e10
                    var v112 uint32 = 2147483648
                    v2 = 1
                    _ = fi110
                    _ = v111
                    _ = v112
                }
                var v113 bool = false
                switch 32767 {
                case 305:
                    v49 = v49
                case 128:
                    v1 += (v1 << ((18446744073709551615) & 63))
                }
                _ = v113
            }
            if true {
                var v114 int64 = -1
                if true {
                    var v115 int64 = v114
                    v116 := GenS8{F0: 65535, F1: 18446744073709551615}
                    v117 := []int{-336, -1}
                    _ = v115
                    _ = v116
                    _ = v117
                } else {
                    var v118 uint64 = 1
                    v3 -= (v3 & (v3 * v3))
                    v119 := GenS6{F0: false, F1: true}
                    _ = v118
                    _ = v119
                }
                _ = v114
            } else {
                switch v4 {
                case "hello":
                    v120 := GenS4{F0: -795, F1: 2}
                    _ = v120
                }
                for fi121 := 0; fi121 < 10; fi121++ {
                    var v122 uint8 = 172
                    _ = fi121
                    _ = v122
                }
            }
        } else {
            var v123 uint8 = (v49 % ((v49) | 1))
            fmt.Println(v49, v6, v49)
            _ = v123
        }
    default:
        switch "a" {
        case "a":
            var v124 int32 = 1
            _ = v124
        default:
            var v125 uint8 = (v49 + (v49 ^ 255))
            _ = v125
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
