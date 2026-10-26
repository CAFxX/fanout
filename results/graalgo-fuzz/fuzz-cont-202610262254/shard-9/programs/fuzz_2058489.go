package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 string }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 int64; F2 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32; F1 int64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int32; F1 uint }
type GenS12 struct { F0 int64; F1 uint64 }
type GenS13 struct { F0 uint; F1 uint16 }
type GenS14 struct { F0 uint32; F1 int8; F2 int8 }
type GenS15 struct { F0 int16; F1 uint16 }
type GenS16 struct { F0 int16; F1 int16; F2 uint64 }
func genFn0(p0 uint64, p1 uint8) uint64 {
    var v7 string = ("a" + (("\n\t" + "") + "\n\t"))
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(v7, v7, fi8)
        v9 := []int{(fi8 + (fi8 ^ fi8)), fi8, -2147483648, (fi8 | -1)}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return (^(p0))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 32767
    var v3 int8 = -128
    var v4 uint8 = 42
    var v5 uint64 = 0
    var v6 int16 = (v2 >> ((556) & 63))
    v10 := GenS1{F0: (v2 * v2), F1: (true && ("a" < "hello")), F2: ("foo bar" + "\n\t")}
    if false {
        v11 := GenS2{F0: (v4 * (v4 * (v4 ^ 127))), F1: 42}
        var v12 float32 = float32((0.0 / 1e10))
        _ = v11
        _ = v12
    } else {
        fmt.Println(v4, v5)
    }
    if false {
        if (!(255 > 18446744073709551615)) {
            for fi13 := 0; fi13 < 2; fi13++ {
                for fi14 := 0; fi14 < 14; fi14++ {
                    v15 := []int{fi13, 32768, (fi14 / (((-(fi13))) | 1)), (fi14 + -9223372036854775808)}
                    v2 = (^(v6))
                    if (uint64(v4) <= v5) {
                        v1[380] = "foo bar"
                        switch fi13 {
                        case 552:
                            switch (fi13 << ((401) & 63)) {
                            case -53:
                                v16 := GenS3{F0: "hello"}
                                _ = v16
                            case -636:
                                switch "\n\t" {
                                case "foo bar":
                                    v17 := GenS1{F0: v6, F1: true, F2: ""}
                                    _ = v17
                                case "\n\t":
                                    fmt.Println(v3, fi14)
                                }
                            }
                        default:
                            var v18 float64 = (float64((3.14159 - 1.5)) + (0.1 + ((float64(1e10) + 0.1) + 0.1)))
                            _ = v18
                        }
                        v19 := GenS1{F0: 1, F1: (439 <= 18446744073709551615), F2: "foo bar"}
                        v20 := GenS4{F0: 0, F1: -110, F2: ""}
                        fmt.Println(v3, v6)
                        _ = v19
                        _ = v20
                    }
                    _ = fi14
                    _ = v15
                }
                for fi21 := 0; fi21 < 11; fi21++ {
                    if ((fi13 == (fi13 % ((fi13) | 1))) || (!("" <= "foo bar"))) {
                        v3 = 100
                        if (1 != 2) {
                            switch fi21 {
                            case -2147483648:
                                var v22 uint8 = v4
                                _ = v22
                            case 2147483647:
                                v4 -= (v4 ^ 1)
                            case 32767:
                                switch "foo bar" {
                                case "\n\t":
                                    v3 -= (v3 & v3)
                                case "hello":
                                    v5 -= (v5 & (v5 - v5))
                                }
                            }
                            switch (^(fi21)) {
                            case -1:
                                v2 = v2
                            case 2147483647:
                                v23 := GenS5{F0: 557}
                                _ = v23
                            }
                            v2 -= v2
                        } else {
                            if false {
                                switch "\n\t" {
                                case "":
                                    switch 447 {
                                    case -876:
                                        v3 = (^(v3))
                                        v24 := GenS2{F0: v4, F1: 2147483648}
                                        v6 += v6
                                        _ = v24
                                    case 2147483647:
                                        v25 := []int{}
                                        _ = v25
                                    default:
                                        if true {
                                            v1[uint32(v5)] = "hello"
                                            v6 = v2
                                        } else {
                                            v26 := v1[838]
                                            v3 = 76
                                            _ = v26
                                        }
                                    }
                                case "a":
                                    if (1e10 >= (((0.1 + (float64(-2.25) - 3.14159)) + ((0.1 - 1.5) * 1.5)) / -2.25)) {
                                        if (!(false || (!false))) {
                                            v4 += 0
                                            switch "hello" {
                                            case "":
                                                fmt.Println(v6, v6)
                                            case "a":
                                                switch "foo bar" {
                                                case "hello":
                                                    v27 := GenS6{F0: ("\n\t" + "\n\t"), F1: false}
                                                    _ = v27
                                                case "":
                                                    v1[42] = "foo bar"
                                                default:
                                                    switch -480 {
                                                    case -518:
                                                        v28 := GenS6{F0: ("a" + ("hello" + "foo bar")), F1: (2147483647 > 0)}
                                                        _ = v28
                                                    default:
                                                        var v29 uint32 = 917
                                                        _ = v29
                                                    }
                                                }
                                            case "foo bar":
                                                fmt.Println(v4, v2)
                                            default:
                                                if (0 != -1) {
                                                    var v30 int8 = (v3 ^ 1)
                                                    fmt.Println(v4)
                                                    _ = v30
                                                } else {
                                                    v31 := GenS3{F0: "a"}
                                                    _ = v31
                                                }
                                            }
                                        } else {
                                            v6 = v2
                                        }
                                        var v32 float64 = (1.5 * float64(float32(1.5)))
                                        var v33 uint64 = v5
                                        var v34 int64 = -1
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                case "foo bar":
                                    v35 := GenS7{F0: (!(true && (float64((0.0 + 1.5)) < 0.0)))}
                                    _ = v35
                                default:
                                    v36 := GenS8{F0: (fi13 - (fi13 / ((fi21) | 1)))}
                                    _ = v36
                                }
                            }
                            v37 := GenS9{F0: 42, F1: 42}
                            var v38 uint16 = 255
                            _ = v37
                            _ = v38
                        }
                    }
                    delete(v1, 233)
                    _ = fi21
                }
                for fi39 := 0; fi39 < 16; fi39++ {
                    v3 += (v3 - 1)
                    _ = fi39
                }
                _ = fi13
            }
        }
        v40 := GenS8{F0: -427}
        if (true || (float64((0.1 + float32(1.5))) > 1.5)) {
            v41 := GenS4{F0: 42, F1: 1, F2: "a"}
            var v42 uint8 = v4
            _ = v41
            _ = v42
        }
        _ = v40
    } else {
        v43 := GenS10{F0: "\n\t"}
        for fi44 := 0; fi44 < 10; fi44++ {
            if false {
                if false {
                    for fi45 := 0; fi45 < 10; fi45++ {
                        if false {
                            var v46 float32 = 1e10
                            var v47 int = fi44
                            delete(v1, uint32(v5))
                            _ = v46
                            _ = v47
                        } else {
                            v5 = genFn0(v5, (v4 ^ (v4 ^ (v4 % (((v4 & (v4 & 128))) | 1)))))
                            var v48 int = 250
                            v49 := GenS8{F0: (^(fi44))}
                            _ = v48
                            _ = v49
                        }
                        switch (fi45 << ((863) & 63)) {
                        case 32767:
                            v5 = v5
                        case -810:
                            if false {
                                if ((v4 << ((65535) & 63)) >= v4) {
                                    v2 += 579
                                } else {
                                    v2 = v2
                                }
                            } else {
                                if (437 >= 32767) {
                                    var v50 int64 = 9223372036854775807
                                    if (-839 == -2147483648) {
                                        fmt.Println(fi45, v3, v3)
                                    } else {
                                        fmt.Println(v6)
                                        if (0.1 != 0.0) {
                                            var v51 uint32 = 65535
                                            var v52 string = "foo bar"
                                            var v53 bool = ((477 < 2) && false)
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        }
                                    }
                                    _ = v50
                                } else {
                                    v54 := GenS11{F0: 42, F1: 962}
                                    v55 := []int{}
                                    fmt.Println(v3)
                                    v4 += (v4 - 215)
                                    v56 := GenS8{F0: fi44}
                                    var v57 int32 = 369
                                    fmt.Println(v5)
                                    v55 = append(v55, 9223372036854775807)
                                    var v58 string = ("\n\t" + "foo bar")
                                    var v59 uint16 = uint16(v4)
                                    var v60 uint32 = 46
                                    if true {
                                        var v61 float64 = ((1e10 * (1e10 / -2.25)) - -2.25)
                                        v6 += (v2 / (((v2 & int16(v5))) | 1))
                                        v5 = v5
                                        var v62 int = fi44
                                        switch v58 {
                                        case "foo bar":
                                            fmt.Println(v6, v59)
                                        case "hello":
                                            var v63 uint16 = (v59 + 0)
                                            _ = v63
                                        default:
                                            var v64 int64 = 0
                                            _ = v64
                                        }
                                        fmt.Println(v57)
                                        _ = v61
                                        _ = v62
                                    } else {
                                        var v65 uint16 = 2
                                        v65 = v59
                                        _ = v65
                                    }
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                fmt.Println(v2, v5)
                            }
                        case 98:
                            if false {
                                v2 -= (v6 ^ (v2 | v2))
                                if false {
                                    var v66 int64 = 0
                                    v67 := GenS12{F0: 833, F1: 215}
                                    _ = v66
                                    _ = v67
                                } else {
                                    if (255 > 2) {
                                        if (!true) {
                                            v6 -= 1
                                        } else {
                                            switch (("foo bar" + "") + "hello") {
                                            case "foo bar":
                                                fmt.Println(v6, fi44)
                                            case "a":
                                                v68 := GenS5{F0: 9223372036854775807}
                                                _ = v68
                                            default:
                                                fmt.Println(v5, v3, v5)
                                            }
                                            if (659 != 14) {
                                                var v69 int32 = 580
                                                var v70 string = "\n\t"
                                                v71 := GenS5{F0: 0}
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                            } else {
                                                if false {
                                                    if false {
                                                        if true {
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v2 = (^(v2))
                                                            }
                                                        }
                                                    } else {
                                                        if false {
                                                            var v72 uint64 = v5
                                                            _ = v72
                                                        } else {
                                                            var v73 int64 = -9223372036854775808
                                                            v2 += (v6 * int16(fi44))
                                                            switch fi45 {
                                                            case 127:
                                                                v5 = genFn0((v5 << ((747) & 63)), v4)
                                                            case -661:
                                                                v4 = v4
                                                            case -52:
                                                                switch int(v73) {
                                                                case -329:
                                                                    var v74 int = fi45
                                                                    _ = v74
                                                                }
                                                            }
                                                            var v75 float64 = 0.1
                                                            _ = v73
                                                            _ = v75
                                                        }
                                                        v6 = (-(v6))
                                                        v76 := []int{9223372036854775807, (^(fi44)), fi45, 2147483647}
                                                        _ = v76
                                                    }
                                                    v77 := GenS13{F0: 228, F1: 530}
                                                    _ = v77
                                                } else {
                                                    v78 := GenS3{F0: (("hello" + ((("a" + "hello") + "\n\t") + "\n\t")) + "\n\t")}
                                                    _ = v78
                                                }
                                                fmt.Println(v4, v6, v5)
                                            }
                                        }
                                        v79 := []int{-9223372036854775808, fi45, fi45}
                                        _ = v79
                                    } else {
                                        v3 -= v3
                                        switch (fi45 + (fi45 >> ((1) & 63))) {
                                        case 296:
                                            if false {
                                                v4 += v4
                                                if false {
                                                    v3 = (v3 & v3)
                                                    var v80 int64 = -288
                                                    v2 = (v6 % ((-32768) | 1))
                                                    v3 += v3
                                                    _ = v80
                                                }
                                                v2 += v2
                                            } else {
                                                v2 = v6
                                                v1[4294967295] = ("" + "foo bar")
                                                if ((385 > 65535) && true) {
                                                    v6 -= int16(fi44)
                                                    var v81 uint8 = 0
                                                    _ = v81
                                                } else {
                                                    var v82 uint = 1
                                                    v6 += v2
                                                    v83 := []int{-579}
                                                    _ = v82
                                                    _ = v83
                                                }
                                            }
                                        }
                                        v84 := GenS4{F0: 0, F1: 0, F2: "a"}
                                        _ = v84
                                    }
                                }
                                v5 = v5
                            }
                        default:
                            switch fi44 {
                            case 2147483647:
                                if (!(false && (0.1 != (3.14159 / 1.0)))) {
                                    v85 := GenS14{F0: 1, F1: -86, F2: 15}
                                    switch fi44 {
                                    case 128:
                                        switch "a" {
                                        case "a":
                                            fmt.Println(v4, v4)
                                        case "foo bar":
                                            if (!true) {
                                                var v86 uint16 = 2
                                                _ = v86
                                            } else {
                                                var v87 uint16 = 281
                                                v88 := GenS10{F0: ("\n\t" + "\n\t")}
                                                var v89 int64 = -794
                                                _ = v87
                                                _ = v88
                                                _ = v89
                                            }
                                        default:
                                            fmt.Println(v4, fi45, fi44)
                                        }
                                    case 4294967295:
                                        v3 -= (^(v3))
                                        v2 += (-(v2))
                                        v6 += (-(v2))
                                    case 9223372036854775807:
                                        if (708 <= -672) {
                                            var v90 int64 = 9223372036854775807
                                            _ = v90
                                        }
                                    }
                                    _ = v85
                                }
                            default:
                                v5 += v5
                            }
                        }
                        _ = fi45
                    }
                    var v91 uint32 = 0
                    _ = v91
                } else {
                    v92 := []int{(fi44 - fi44), (^(fi44))}
                    _ = v92
                }
            } else {
                switch ("foo bar" + (("foo bar" + "") + "")) {
                case "":
                    var v93 float32 = -2.25
                    _ = v93
                }
            }
            _ = fi44
        }
        _ = v43
    }
    for fi94 := 0; fi94 < 10; fi94++ {
        switch fi94 {
        case -858:
            var v95 uint64 = (v5 + (v5 & 2))
            _ = v95
        case 9223372036854775807:
            if (float64(-2.25) <= ((-2.25 - float64((0.0 / float32(1e10)))) - 1e10)) {
                for fi96 := 0; fi96 < 20; fi96++ {
                    v2 -= v6
                    var v97 uint = 797
                    var v98 uint64 = (v5 | v5)
                    _ = fi96
                    _ = v97
                    _ = v98
                }
                var v99 float32 = 0.1
                var v100 string = ((((("foo bar" + "") + "hello") + "") + (("a" + "\n\t") + ((("a" + "hello") + "") + "foo bar"))) + "\n\t")
                _ = v100
                _ = v99
            }
        }
        if false {
            var v101 uint64 = genFn0(uint64(v4), v4)
            v102 := []int{fi94}
            _ = v101
            _ = v102
        }
        _ = fi94
    }
    for fi103 := 0; fi103 < 20; fi103++ {
        v4 = (v4 & v4)
        _ = fi103
    }
    switch -977 {
    case 356:
        fmt.Println(v3, v2)
    default:
        if true {
            v5 = v5
        }
    }
    if true {
        v104 := []int{int(v5), -496}
        var v105 uint8 = (v4 | v4)
        for fi106 := 0; fi106 < 14; fi106++ {
            var v107 float64 = 0.0
            v108 := GenS15{F0: (v2 >> ((42) & 63)), F1: 65535}
            _ = fi106
            _ = v107
            _ = v108
        }
        _ = v104
        _ = v105
    }
    fmt.Println(v5, v3)
    if (0 > 1) {
        v109 := []int{32767, -2147483648}
        _ = v109
    }
    for fi110 := 0; fi110 < 3; fi110++ {
        for fi111 := 0; fi111 < 9; fi111++ {
            var v112 int64 = 1
            _ = fi111
            _ = v112
        }
        if ("foo bar" <= "hello") {
            switch 9223372036854775807 {
            case 481:
                v6 = v2
            case -1000:
                for fi113 := 0; fi113 < 19; fi113++ {
                    switch -9223372036854775808 {
                    case -202:
                        v2 -= v2
                    case -918:
                        v114 := v1[uint32(v5)]
                        _ = v114
                    case -788:
                        var v115 uint64 = 852
                        _ = v115
                    }
                    v6 -= -750
                    v116 := []int{}
                    if (!((!false) && true)) {
                        v3 = (v3 | v3)
                        v117 := GenS16{F0: int16(v4), F1: -204, F2: uint64(v4)}
                        if (!true) {
                            fmt.Println(v3, v3, v5)
                            var v118 int64 = -837
                            _ = v118
                        }
                        _ = v117
                    }
                    v4 += v4
                    _ = fi113
                    _ = v116
                }
            case 733:
                if false {
                    v5 = genFn0(0, (^(v4)))
                    v4 = (v4 / ((v4) | 1))
                }
            }
            v3 -= v3
            v5 -= 677
        } else {
            delete(v1, 608)
        }
        _ = fi110
    }
    v119 := GenS7{F0: true}
    v4 -= 2
    if (true || (false || false)) {
        if ("hello" <= "hello") {
            if true {
                var v120 int8 = v3
                v121 := GenS8{F0: 2147483647}
                switch ("hello" + ("a" + ("" + ""))) {
                case "foo bar":
                    v4 -= v4
                }
                _ = v120
                _ = v121
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
