package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int8; F1 bool; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 int16; F1 uint64; F2 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string; F1 bool; F2 int8 }
type GenS12 struct { F0 int; F1 uint8; F2 int }
type GenS13 struct { F0 int64; F1 uint16 }
type GenS14 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS15 struct { F0 int }
func genFn0(p0 string, p1 int) string {
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 float64 = 3.14159
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(p1, fi8)
            _ = fi8
        }
        fmt.Println(p0, v7, fi6)
        _ = fi6
        _ = v7
    }
    v9 := []int{p1}
    if (18446744073709551615 >= 2) {
        var v10 uint32 = 1
        _ = v10
    } else {
        p0 = p0
        p1 = -291
    }
    var v11 uint16 = 2
    _ = p0
    _ = p1
    _ = v11
    _ = v9
    return ("hello" + ("\n\t" + (p0 + "")))
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 2
    var v3 uint32 = 481
    var v4 uint = 1
    v5 := []string{"hello", "", "a"}
    switch 9223372036854775807 {
    case 2147483647:
        v2 -= v3
    case 127:
        var v12 uint8 = uint8(v4)
        _ = v12
    case -2147483648:
        var v13 uint8 = uint8(v2)
        _ = v13
    }
    v3 += 376
    if ((202 > int64(v4)) || (32768 == -614)) {
        if false {
            if (!true) {
                v2 = (v2 | (v3 >> (((^(v4))) & 63)))
            } else {
                for fi14 := 0; fi14 < 1; fi14++ {
                    if false {
                        v15 := GenS1{F0: (fi14 | 44), F1: (fi14 + -9223372036854775808)}
                        _ = v15
                    }
                    v4 = (^(v4))
                    v16 := GenS1{F0: fi14, F1: fi14}
                    _ = fi14
                    _ = v16
                }
                for fi17 := 0; fi17 < 1; fi17++ {
                    var v18 uint64 = 18446744073709551615
                    if (0.0 < 0.0) {
                        var v19 int8 = -115
                        for fi20 := 0; fi20 < 8; fi20++ {
                            switch fi17 {
                            case -397:
                                var v21 int = (fi20 ^ (fi17 * (-(fi17))))
                                var v22 int8 = (v19 & v19)
                                v2 += 918
                                switch "foo bar" {
                                case "\n\t":
                                    v23 := GenS1{F0: 32767, F1: fi17}
                                    _ = v23
                                }
                                v21 = fi17
                                v18 = v18
                                _ = v21
                                _ = v22
                            case 4294967295:
                                if true {
                                    var v24 string = "a"
                                    _ = v24
                                } else {
                                    var v25 int16 = 1000
                                    fmt.Println(v18, v19, v2)
                                    var v26 uint = v4
                                    var v27 bool = (!((-9223372036854775808 != 1) || true))
                                    var v28 uint16 = 2
                                    var v29 bool = true
                                    delete(v1, fi17)
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            case -9223372036854775808:
                                if (-2.25 > 0.0) {
                                    if false {
                                        v3 = v3
                                        v30 := GenS1{F0: (fi20 % (((fi20 * (fi20 + 1))) | 1)), F1: 127}
                                        var v31 bool = (770 == 24)
                                        _ = v30
                                        _ = v31
                                    }
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = "foo bar"
                                    }
                                }
                            default:
                                v32 := GenS1{F0: 950, F1: -450}
                                _ = v32
                            }
                            if ((false || (32768 != fi20)) || true) {
                                v4 += 0
                                v4 = v4
                                v18 = (v18 * v18)
                                if ((v4 + (v4 >> ((v4) & 63))) > 445) {
                                    v33 := GenS1{F0: fi20, F1: (fi17 % ((9223372036854775807) | 1))}
                                    v34 := GenS1{F0: (fi20 - -87), F1: (fi20 >> (((^(v4))) & 63))}
                                    _ = v33
                                    _ = v34
                                } else {
                                    if (!false) {
                                        fmt.Println(fi20)
                                    }
                                }
                                var v35 int = (fi20 ^ 32767)
                                if len(v5) > 0 {
                                    v5[0] = "\n\t"
                                }
                                var v36 int = (fi17 | (fi20 << ((755) & 63)))
                                v37 := GenS1{F0: 0, F1: 820}
                                var v38 int64 = 9223372036854775807
                                _ = v35
                                _ = v36
                                _ = v37
                                _ = v38
                            } else {
                                if false {
                                    switch (fi17 << ((2) & 63)) {
                                    case 42:
                                        fmt.Println(v18, fi17)
                                    case 32767:
                                        v39 := GenS2{F0: 0, F1: false, F2: v4}
                                        _ = v39
                                    }
                                    fmt.Println(fi17)
                                }
                            }
                            v2 = (v3 - 2147483648)
                            _ = fi20
                        }
                        if false {
                            switch ("foo bar" + "") {
                            case "foo bar":
                                for fi40 := 0; fi40 < 2; fi40++ {
                                    if (v18 >= 2) {
                                        v18 = v18
                                    } else {
                                        var v41 float32 = ((((3.14159 + 3.14159) * (3.14159 * 0.1)) + (0.0 + 1.5)) * 3.14159)
                                        _ = len(v5)
                                        if (false && false) {
                                            var v42 uint = (^(v4))
                                            var v43 int16 = 32767
                                            fmt.Println(v3, v3)
                                            var v44 uint16 = 340
                                            var v45 uint32 = (v3 ^ 214)
                                            var v46 float64 = ((1.5 / (float64(v41) * 3.14159)) * 3.14159)
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        }
                                        v4 += v4
                                        _ = v41
                                    }
                                    v2 = 0
                                    fmt.Println(fi40, v2, v4)
                                    _ = fi40
                                }
                            case "hello":
                                v47 := v1[fi17]
                                _ = v47
                            }
                            var v48 float32 = ((3.14159 - float32((0.0 / -2.25))) + 1e10)
                            _ = v48
                        } else {
                            if false {
                                var v49 int16 = 1000
                                if false {
                                    switch fi17 {
                                    case 0:
                                        var v50 int8 = 1
                                        _ = v50
                                    case -9223372036854775808:
                                        if len(v5) > 0 {
                                            v5[0] = (("\n\t" + "foo bar") + "")
                                        }
                                    default:
                                        if (!true) {
                                            v2 -= (v3 >> ((v4) & 63))
                                            v4 = v4
                                        }
                                    }
                                } else {
                                    switch (fi17 >> (((v4 << (((^(v4))) & 63))) & 63)) {
                                    case 4294967295:
                                        var v51 int8 = (v19 ^ (v19 + v19))
                                        _ = v51
                                    case 9223372036854775807:
                                        v52 := GenS3{F0: 42}
                                        _ = v52
                                    case -9223372036854775808:
                                        var v53 int = fi17
                                        _ = v53
                                    default:
                                        if false {
                                            v54 := GenS1{F0: fi17, F1: fi17}
                                            var v55 bool = true
                                            _ = v54
                                            _ = v55
                                        }
                                    }
                                    delete(v1, (fi17 >> ((65535) & 63)))
                                }
                                _ = v49
                            }
                        }
                        _ = v19
                    } else {
                        if true {
                            var v56 int = (fi17 % ((-9223372036854775808) | 1))
                            if (203 >= 128) {
                                v2 -= (v3 * (^(v3)))
                                v3 += v3
                            }
                            if (!("\n\t" >= ("a" + (("" + "hello") + ("foo bar" + (("\n\t" + (((("foo bar" + "hello") + "a") + "foo bar") + "hello")) + (("" + "\n\t") + "\n\t"))))))) {
                                v57 := v1[(fi17 << (((v4 ^ v4)) & 63))]
                                _ = v57
                            }
                            _ = v56
                        } else {
                            v3 -= (v2 << ((v4) & 63))
                            var v58 bool = true
                            _ = v58
                        }
                        if true {
                            if (true || (v18 <= (v18 * (^(v18))))) {
                                v4 = (v4 * (v4 + v4))
                                if (!(false && true)) {
                                    _ = len(v1)
                                    v59 := GenS1{F0: 127, F1: -1}
                                    _ = v59
                                }
                                for fi60 := 0; fi60 < 18; fi60++ {
                                    var v61 int16 = 297
                                    _ = len(v5)
                                    _ = fi60
                                    _ = v61
                                }
                            } else {
                                v4 -= v4
                                v62 := GenS2{F0: 127, F1: true, F2: v4}
                                for fi63 := 0; fi63 < 2; fi63++ {
                                    var v64 int = 1
                                    _ = fi63
                                    _ = v64
                                }
                                _ = v62
                            }
                        }
                        for fi65 := 0; fi65 < 13; fi65++ {
                            if (((v18 * 997) != (^(v18))) && (((fi65 ^ fi17) < (fi65 ^ 0)) && (0.1 <= (1.5 + 1.5)))) {
                                if (!(-119 == -1)) {
                                    var v66 int = 32767
                                    var v67 uint8 = 42
                                    var v68 int16 = 1
                                    var v69 int = fi65
                                    fmt.Println(v4)
                                    _ = len(v5)
                                    if ((v67 - v67) < v67) {
                                        if false {
                                            var v70 string = ""
                                            v71 := GenS3{F0: v67}
                                            _ = v70
                                            _ = v71
                                        }
                                        if (v18 < v18) {
                                            fmt.Println(v68)
                                            v18 = v18
                                        } else {
                                            v66 = int(v18)
                                            v72 := GenS4{F0: 1, F1: 42}
                                            v2 = v3
                                            _ = v72
                                        }
                                        v73 := GenS5{F0: v2, F1: (!(!(false || (!false)))), F2: 1}
                                        _ = v73
                                    } else {
                                        v18 -= v18
                                        v69 = -9223372036854775808
                                    }
                                    v67 += v67
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                    _ = v69
                                } else {
                                    var v74 string = ("a" + "\n\t")
                                    _ = v74
                                }
                                v1[42] = false
                            } else {
                                v75 := GenS6{F0: v4, F1: (1.5 < ((0.1 + -2.25) + 1e10))}
                                v76 := GenS7{F0: 1, F1: v18, F2: (^(v4))}
                                _ = v75
                                _ = v76
                            }
                            _ = fi65
                        }
                    }
                    v2 = (v3 >> (((^(v4))) & 63))
                    _ = fi17
                    _ = v18
                }
            }
        }
    }
    for fi77 := 0; fi77 < 4; fi77++ {
        v3 += (^(v3))
        v3 -= 2147483648
        _ = fi77
    }
    for fi78 := 0; fi78 < 9; fi78++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi78
    }
    var v79 string = ""
    var v80 uint32 = v2
    for fi81 := 0; fi81 < 9; fi81++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v82 int = (fi81 ^ fi81)
        _ = fi81
        _ = v82
    }
    v80 += (v80 * (v3 ^ 65535))
    if len(v5) > 0 {
        v5[0] = v79
    }
    if true {
        fmt.Println(v80)
        v79 = (v79 + (("" + ("hello" + "foo bar")) + v79))
        var v83 int16 = 1000
        _ = v83
    } else {
        if false {
            for fi84 := 0; fi84 < 13; fi84++ {
                switch (fi84 % ((fi84) | 1)) {
                case -887:
                    v4 += (v4 ^ v4)
                case -805:
                    v85 := GenS8{F0: 127}
                    _ = v85
                case -1:
                    v4 = 2
                default:
                    v2 = v80
                }
                for fi86 := 0; fi86 < 17; fi86++ {
                    switch "foo bar" {
                    case "\n\t":
                        var v87 int32 = 1
                        _ = v87
                    case "foo bar":
                        fmt.Println(v4, v2)
                    case "hello":
                        v88 := GenS9{F0: v3}
                        _ = v88
                    default:
                        v89 := GenS10{F0: -1}
                        _ = v89
                    }
                    fmt.Println(v3)
                    _ = fi86
                }
                for fi90 := 0; fi90 < 17; fi90++ {
                    v91 := GenS4{F0: 42, F1: 42}
                    var v92 bool = ((v79 + v79) > "hello")
                    v93 := v1[1]
                    if true {
                        var v94 uint16 = 514
                        fmt.Println(v3)
                        v2 = v2
                        _ = v94
                    } else {
                        if v92 {
                            _ = len(v5)
                            v95 := GenS1{F0: fi84, F1: 0}
                            v80 = 862
                            _ = v95
                        } else {
                            fmt.Println(fi84, v92, v93)
                            if true {
                                if false {
                                    var v96 uint32 = 103
                                    _ = v96
                                } else {
                                    if true {
                                        switch (fi84 << ((v4) & 63)) {
                                        case -9223372036854775808:
                                            switch (fi84 >> ((v4) & 63)) {
                                            case -9223372036854775808:
                                                v97 := GenS7{F0: 0, F1: uint64(v4), F2: (v4 % ((65535) | 1))}
                                                _ = v97
                                            default:
                                                v98 := GenS11{F0: ("\n\t" + ("" + v79)), F1: (true && (212 == 0)), F2: 127}
                                                _ = v98
                                            }
                                        case -2147483648:
                                            v93 = v92
                                        case 0:
                                            var v99 uint32 = 4294967295
                                            if (!true) {
                                                switch (fi90 >> (((^(v4))) & 63)) {
                                                case 32768:
                                                    v93 = v93
                                                default:
                                                    fmt.Println(fi84)
                                                }
                                                _ = len(v5)
                                            }
                                            _ = v99
                                        default:
                                            v5 = append(v5, "a")
                                        }
                                        switch "\n\t" {
                                        case "\n\t":
                                            switch "" {
                                            case "a":
                                                v100 := GenS6{F0: 731, F1: true}
                                                _ = v100
                                            case "hello":
                                                if true {
                                                    fmt.Println(fi84)
                                                    _ = len(v5)
                                                }
                                            }
                                        case "hello":
                                            if (!v93) {
                                                v93 = v92
                                                _ = len(v5)
                                            }
                                        }
                                        v101 := GenS12{F0: (-(fi84)), F1: 127, F2: (fi90 & fi84)}
                                        _ = v101
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = "a"
                                    }
                                    var v102 int = (fi90 / (((fi84 >> ((v4) & 63))) | 1))
                                    _ = v102
                                }
                            }
                            var v103 int64 = -1
                            _ = v103
                        }
                        switch v79 {
                        case "":
                            var v104 uint = v4
                            _ = v104
                        default:
                            var v105 int64 = 0
                            _ = v105
                        }
                    }
                    _ = fi90
                    _ = v91
                    _ = v92
                    _ = v93
                }
                _ = fi84
            }
            _ = len(v1)
            if (!true) {
                if (92 != 86) {
                    v79 = v79
                    if true {
                        if false {
                            switch 1 {
                            case 9223372036854775807:
                                var v106 int64 = 370
                                _ = v106
                            case -2147483648:
                                _ = len(v1)
                            }
                            var v107 int8 = -1
                            _ = v107
                        } else {
                            fmt.Println(v2)
                        }
                        fmt.Println(v3, v80)
                        delete(v1, 32767)
                    } else {
                        var v108 bool = (99 >= 1)
                        for fi109 := 0; fi109 < 10; fi109++ {
                            var v110 uint32 = v3
                            v111 := GenS8{F0: 127}
                            _ = fi109
                            _ = v110
                            _ = v111
                        }
                        fmt.Println(v79, v3, v108)
                        _ = v108
                    }
                    v3 += 65535
                }
                if len(v5) > 0 {
                    v5[0] = (v79 + v79)
                }
                var v112 uint16 = 32767
                _ = v112
            }
        }
        for fi113 := 0; fi113 < 12; fi113++ {
            v114 := GenS13{F0: int64(fi113), F1: 664}
            for fi115 := 0; fi115 < 18; fi115++ {
                v116 := GenS14{F0: (v2 * v3), F1: v2, F2: -978}
                var v117 int = fi115
                v118 := GenS15{F0: fi115}
                _ = fi115
                _ = v116
                _ = v117
                _ = v118
            }
            fmt.Println(v79, v80, v2)
            _ = fi113
            _ = v114
        }
    }
    if (v4 != v4) {
        v119 := GenS1{F0: -822, F1: 0}
        _ = v119
    } else {
        v2 = 4294967295
        for fi120 := 0; fi120 < 20; fi120++ {
            fmt.Println(fi120)
            _ = fi120
        }
    }
    var v121 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v79
    _ = v80
}
