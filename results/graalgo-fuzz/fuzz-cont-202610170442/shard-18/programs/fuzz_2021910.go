package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int; F1 uint64; F2 string }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int16; F1 int64; F2 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 bool; F2 int }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 int32; F1 int }
type GenS10 struct { F0 uint32; F1 int }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint8; F1 int64 }
func genFn0(p0 int8, p1 int) int {
    v7 := GenS1{F0: (p0 * (p0 % (((p0 << ((2) & 63))) | 1)))}
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 bool, p1 string, p2 uint8) int64 {
    fmt.Println(p0, p1, p0)
    v8 := GenS2{F0: 0, F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 480
    var v3 int = 128
    var v4 uint = uint(v2)
    v5 := []string{"foo bar", "foo bar", "a"}
    var v6 string = ("" + "")
    v6 = "\n\t"
    v9 := v1["a"]
    if ("\n\t" < ((("foo bar" + (v6 + (v6 + v6))) + "") + "")) {
        if (3.14159 > float64(float32((float64(float32(0.0)) + 1.5)))) {
            if (uint(v2) <= v4) {
                if ((571 != 506) && v9) {
                    fmt.Println(v2, v4)
                    if ("\n\t" == v6) {
                        var v10 uint = (v4 % ((v4) | 1))
                        fmt.Println(v6)
                        _ = v10
                    }
                    var v11 uint8 = 127
                    _ = v11
                } else {
                    if ((v2 ^ (v2 + (v2 ^ (v2 ^ 65535)))) >= (v2 / ((v2) | 1))) {
                        if (42 != -128) {
                            for fi12 := 0; fi12 < 3; fi12++ {
                                var v13 int64 = 354
                                if (-94 == 0) {
                                    v14 := GenS1{F0: -116}
                                    var v15 uint32 = 42
                                    _ = v14
                                    _ = v15
                                }
                                _ = fi12
                                _ = v13
                            }
                        }
                    } else {
                        for fi16 := 0; fi16 < 5; fi16++ {
                            v17 := GenS3{F0: (v3 % ((-1) | 1)), F1: 4294967295, F2: ("\n\t" + (v6 + (("\n\t" + v6) + ((v6 + v6) + "foo bar"))))}
                            switch (^(v3)) {
                            case 150:
                                if v9 {
                                    v4 = 65535
                                    var v18 float32 = 3.14159
                                    v5 = append(v5, ((v6 + v6) + (v6 + ("\n\t" + ((("" + v6) + ("foo bar" + (v6 + "\n\t"))) + (v6 + "a"))))))
                                    _ = v18
                                }
                            }
                            _ = fi16
                            _ = v17
                        }
                        for fi19 := 0; fi19 < 1; fi19++ {
                            v20 := GenS4{F0: -1, F1: 9223372036854775807}
                            _ = fi19
                            _ = v20
                        }
                        var v21 int16 = 0
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 14; fi22++ {
                        for fi23 := 0; fi23 < 17; fi23++ {
                            fmt.Println(fi23, v6, v4)
                            if (0.0 <= (1.5 - float32(1.5))) {
                                v24 := GenS5{F0: -403, F1: 1, F2: int(v2)}
                                _ = v24
                            }
                            fmt.Println(v6)
                            var v25 bool = (!(0.0 != 1.5))
                            _ = fi23
                            _ = v25
                        }
                        fmt.Println(v4)
                        _ = fi22
                    }
                    v1[("a" + (((v6 + "a") + v6) + v6))] = ((!(v4 == (^(v4)))) || ((((!((128 != 88) && (0 >= 9223372036854775807))) || ((v4 | (v4 ^ v4)) < 0)) || true) && (v9 && ((true && false) && ((0 <= genFn1(v9, (v6 + v6), 183)) && (0 >= 18446744073709551615))))))
                }
                v26 := GenS4{F0: 1, F1: 453}
                _ = v26
            }
        }
    } else {
        for fi27 := 0; fi27 < 14; fi27++ {
            v2 -= uint16(v4)
            _ = fi27
        }
    }
    if (v6 < "\n\t") {
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 uint16 = v2
            _ = fi28
            _ = v29
        }
        fmt.Println(v9, v6)
        switch v6 {
        case "hello":
            switch -787 {
            case -1:
                var v30 int8 = 42
                _ = v30
            case 42:
                switch "hello" {
                case "a":
                    if (!true) {
                        v31 := v1[v6]
                        var v32 int64 = 42
                        _ = v31
                        _ = v32
                    }
                case "hello":
                    if (((v9 && true) && false) || (v9 || (false || false))) {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v9, v6, v3)
                        default:
                            if (float64((0.0 - float32(float64(-2.25)))) != ((0.0 / (0.1 * 3.14159)) * float64((-2.25 - 0.0)))) {
                                switch (v3 >> ((42) & 63)) {
                                case 32768:
                                    v9 = v9
                                default:
                                    fmt.Println(v9)
                                }
                                v33 := GenS2{F0: v3, F1: 0}
                                for fi34 := 0; fi34 < 10; fi34++ {
                                    v6 = v6
                                    _ = fi34
                                }
                                _ = v33
                            }
                        }
                        fmt.Println(v6, v6, v4)
                        if v9 {
                            v9 = v9
                        } else {
                            v4 += v4
                            var v35 int32 = -1
                            _ = v35
                        }
                    } else {
                        v4 = 614
                    }
                }
            case -9223372036854775808:
                for fi36 := 0; fi36 < 15; fi36++ {
                    fmt.Println(fi36, v3, v9)
                    fmt.Println(v2, fi36, v2)
                    fmt.Println(v6)
                    _ = fi36
                }
            }
        case "a":
            if v9 {
                v1["foo bar"] = (255 <= 0)
                if v9 {
                    for fi37 := 0; fi37 < 16; fi37++ {
                        switch v6 {
                        case "foo bar":
                            v2 -= v2
                        case "\n\t":
                            switch -9223372036854775808 {
                            case -9223372036854775808:
                                v5 = append(v5, v6)
                            case 42:
                                v38 := v1[""]
                                _ = v38
                            default:
                                for fi39 := 0; fi39 < 1; fi39++ {
                                    if true {
                                        if (!true) {
                                            var v40 int64 = genFn1(v9, (v6 + ((("hello" + "a") + ((v6 + ((("a" + ((v6 + v6) + ("hello" + (v6 + v6)))) + "foo bar") + v6)) + "")) + ((v6 + (("\n\t" + v6) + "foo bar")) + ("hello" + v6)))), 42)
                                            _ = v40
                                        } else {
                                            v41 := GenS5{F0: 0, F1: 0, F2: (fi39 / (((v3 * (v3 - (v3 / ((-2147483648) | 1))))) | 1))}
                                            fmt.Println(v3, v6, v3)
                                            if ((-2.25 + 1e10) > float64(float32(1.5))) {
                                                v9 = (18446744073709551615 <= 682)
                                                fmt.Println(v9)
                                                _ = len(v1)
                                                v3 += (fi37 - (fi37 >> ((4294967295) & 63)))
                                                v3 -= (fi39 + (fi37 ^ -9223372036854775808))
                                            }
                                            _ = v41
                                        }
                                    }
                                    _ = fi39
                                }
                            }
                        }
                        for fi42 := 0; fi42 < 12; fi42++ {
                            _ = len(v5)
                            var v43 int8 = 1
                            if false {
                                switch v6 {
                                case "\n\t":
                                    v9 = (!v9)
                                case "a":
                                    switch ("" + (v6 + ("" + v6))) {
                                    case "a":
                                        var v44 uint32 = 150
                                        _ = v44
                                    case "\n\t":
                                        switch fi37 {
                                        case 42:
                                            if (!(!v9)) {
                                                switch ((v6 + "a") + ("a" + v6)) {
                                                case "":
                                                    _ = len(v1)
                                                case "a":
                                                    if (false && true) {
                                                        fmt.Println(fi42)
                                                        fmt.Println(fi42)
                                                        switch (fi42 - fi37) {
                                                        case 0:
                                                            var v45 bool = true
                                                            _ = v45
                                                        default:
                                                            v43 -= 1
                                                        }
                                                    } else {
                                                        var v46 string = v6
                                                        if (18446744073709551615 < 18446744073709551615) {
                                                            v47 := GenS1{F0: 42}
                                                            var v48 int64 = 1
                                                            switch ((("hello" + ("a" + v46)) + v6) + "a") {
                                                            case "":
                                                                v5 = append(v5, v46)
                                                            case "foo bar":
                                                                v3 -= 0
                                                            default:
                                                                v2 += uint16(v4)
                                                            }
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        var v49 float32 = 0.0
                                                        _ = v46
                                                        _ = v49
                                                    }
                                                    if false {
                                                        if (v2 != v2) {
                                                            switch v6 {
                                                            case "":
                                                                v3 = (fi42 * int(v2))
                                                            case "\n\t":
                                                                v50 := GenS4{F0: v43, F1: fi42}
                                                                _ = v50
                                                            case "foo bar":
                                                                var v51 int8 = v43
                                                                _ = v51
                                                            default:
                                                                var v52 int = (fi42 & genFn0((v43 >> (((v4 - v4)) & 63)), fi42))
                                                                _ = v52
                                                            }
                                                            var v53 int = fi37
                                                            var v54 int = (fi42 ^ v53)
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                        v3 = fi42
                                                        if len(v5) > 0 {
                                                            v5[0] = v6
                                                        }
                                                    } else {
                                                        v3 += (fi37 / (((^(v3))) | 1))
                                                        var v55 int8 = (v43 * 42)
                                                        _ = v55
                                                    }
                                                default:
                                                    var v56 string = ("foo bar" + v6)
                                                    _ = v56
                                                }
                                                switch fi37 {
                                                case 42:
                                                    fmt.Println(v4, v3)
                                                case -490:
                                                    v57 := GenS6{F0: 918}
                                                    _ = v57
                                                case 0:
                                                    var v58 uint32 = 0
                                                    _ = v58
                                                default:
                                                    v43 -= 100
                                                }
                                                var v59 int16 = -827
                                                _ = v59
                                            } else {
                                                if v9 {
                                                    var v60 float32 = -2.25
                                                    _ = v60
                                                } else {
                                                    if (!(!v9)) {
                                                        if ((-(v3)) < genFn0(v43, int(v2))) {
                                                            if len(v5) > 0 {
                                                                v5[0] = v6
                                                            }
                                                            switch v3 {
                                                            case 805:
                                                                var v61 uint8 = 42
                                                                _ = v61
                                                            case 319:
                                                                var v62 uint8 = 1
                                                                _ = v62
                                                            }
                                                        } else {
                                                            v2 = (v2 ^ 2)
                                                        }
                                                    }
                                                    var v63 uint8 = 184
                                                    _ = v63
                                                }
                                            }
                                        default:
                                            v4 = (^(v4))
                                        }
                                    }
                                case "foo bar":
                                    v4 += 42
                                }
                                switch int(v4) {
                                case 32767:
                                    if v9 {
                                        fmt.Println(v4, v9, v4)
                                        v6 = (v6 + (v6 + (((v6 + (v6 + (((v6 + v6) + "a") + v6))) + v6) + (v6 + "foo bar"))))
                                    } else {
                                        v3 += (fi42 * fi42)
                                        if (!((!v9) && v9)) {
                                            v2 = v2
                                            v64 := GenS1{F0: 54}
                                            _ = v64
                                        } else {
                                            var v65 int16 = 32767
                                            _ = len(v1)
                                            _ = v65
                                        }
                                        v2 = (v2 >> ((v4) & 63))
                                        v2 -= (v2 - (^(v2)))
                                    }
                                case -391:
                                    v66 := GenS6{F0: -1}
                                    _ = v66
                                case -465:
                                    if false {
                                        if v9 {
                                            switch v6 {
                                            case "":
                                                v67 := GenS7{F0: 4294967295, F1: (894 >= v4), F2: -9223372036854775808}
                                                _ = v67
                                            case "foo bar":
                                                v2 += 913
                                                if (v9 && (v2 == 443)) {
                                                    v68 := v1[""]
                                                    v3 = (fi42 << ((18446744073709551615) & 63))
                                                    _ = v68
                                                } else {
                                                    v9 = v9
                                                }
                                            }
                                        }
                                    } else {
                                        _ = len(v5)
                                        if (42 < 0) {
                                            v1["foo bar"] = v9
                                            fmt.Println(v43, v9, v3)
                                        } else {
                                            var v69 int8 = v43
                                            v70 := GenS7{F0: 1, F1: true, F2: 127}
                                            v43 = v43
                                            v4 += v4
                                            _ = v69
                                            _ = v70
                                        }
                                        v43 = -1
                                    }
                                default:
                                    switch v3 {
                                    case 128:
                                        switch fi42 {
                                        case 32767:
                                            fmt.Println(v4)
                                        case 623:
                                            v71 := GenS8{F0: v4, F1: (fi37 << ((876) & 63))}
                                            _ = v71
                                        }
                                    case 0:
                                        switch "foo bar" {
                                        case "\n\t":
                                            v72 := v1["foo bar"]
                                            _ = v72
                                        case "a":
                                            v1[v6] = (!(!(!true)))
                                        }
                                    default:
                                        switch 9223372036854775807 {
                                        case 1:
                                            v2 = (v2 | v2)
                                        case -9223372036854775808:
                                            if (-1 <= 876) {
                                                switch v6 {
                                                case "\n\t":
                                                    fmt.Println(v6, v2)
                                                }
                                                v73 := GenS9{F0: 0, F1: fi42}
                                                if (-1 != 0) {
                                                    v74 := v1[("foo bar" + (v6 + v6))]
                                                    v75 := GenS7{F0: 18446744073709551615, F1: (0.1 > (0.0 * -2.25)), F2: (fi37 * fi42)}
                                                    fmt.Println(fi42, v4, v4)
                                                    _ = v74
                                                    _ = v75
                                                } else {
                                                    var v76 int8 = (v43 ^ (v43 << ((v4) & 63)))
                                                    v9 = v9
                                                    var v77 int8 = (v43 << ((v4) & 63))
                                                    _ = v76
                                                    _ = v77
                                                }
                                                var v78 uint32 = 2147483647
                                                v6 = (v6 + v6)
                                                v79 := GenS2{F0: (^(fi37)), F1: (^(v78))}
                                                v78 = v78
                                                _ = v73
                                                _ = v78
                                                _ = v79
                                            } else {
                                                v80 := GenS9{F0: -2147483648, F1: (fi37 * (-(fi37)))}
                                                _ = v80
                                            }
                                        default:
                                            v2 = 65535
                                        }
                                    }
                                }
                            }
                            _ = fi42
                            _ = v43
                        }
                        var v81 bool = false
                        _ = fi37
                        _ = v81
                    }
                    if v9 {
                        fmt.Println(v2, v4, v3)
                    }
                } else {
                    for fi82 := 0; fi82 < 19; fi82++ {
                        var v83 uint = v4
                        var v84 int16 = -323
                        for fi85 := 0; fi85 < 12; fi85++ {
                            if v9 {
                                var v86 int = -674
                                v87 := GenS10{F0: 333, F1: 127}
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v86
                                _ = v87
                            } else {
                                if (v9 && (v3 < genFn0(127, (fi85 | fi82)))) {
                                    v88 := GenS11{F0: (!v9)}
                                    switch (v3 / ((fi82) | 1)) {
                                    case -866:
                                        v4 -= (v4 * uint(v2))
                                    case 1:
                                        v4 = (^(v83))
                                    default:
                                        switch (-(v3)) {
                                        case 0:
                                            v9 = v9
                                        case 32767:
                                            v2 = v2
                                        default:
                                            switch 0 {
                                            case 860:
                                                var v89 int8 = 100
                                                _ = v89
                                            default:
                                                var v90 int64 = 9223372036854775807
                                                _ = v90
                                            }
                                        }
                                    }
                                    var v91 string = v6
                                    v3 += -957
                                    _ = v88
                                    _ = v91
                                } else {
                                    var v92 uint = (v83 ^ v4)
                                    if (-1 >= 1) {
                                        _ = len(v5)
                                        switch v6 {
                                        case "hello":
                                            v93 := GenS6{F0: (v84 % ((-1) | 1))}
                                            delete(v1, v6)
                                            var v94 int = (fi82 + fi82)
                                            _ = v93
                                            _ = v94
                                        case "foo bar":
                                            var v95 uint = (v4 - 349)
                                            _ = v95
                                        }
                                        if v9 {
                                            var v96 int8 = -33
                                            var v97 int = (^(v3))
                                            v98 := GenS7{F0: 18446744073709551615, F1: v9, F2: (fi85 | (^(fi82)))}
                                            _ = v96
                                            _ = v97
                                            _ = v98
                                        }
                                        v83 -= 65535
                                    }
                                    _ = v92
                                }
                                switch (fi85 - (^(fi82))) {
                                case 0:
                                    _ = len(v5)
                                case 32768:
                                    var v99 uint32 = uint32(v4)
                                    _ = v99
                                default:
                                    var v100 int64 = -80
                                    _ = v100
                                }
                                v2 -= (v2 - 178)
                            }
                            _ = fi85
                        }
                        _ = fi82
                        _ = v83
                        _ = v84
                    }
                    if true {
                        v101 := GenS4{F0: int8(v3), F1: (v3 | (v3 ^ (v3 ^ (v3 - genFn0(-112, v3)))))}
                        fmt.Println(v4)
                        for fi102 := 0; fi102 < 16; fi102++ {
                            v103 := GenS2{F0: (v3 % ((23) | 1)), F1: 50}
                            if false {
                                fmt.Println(v9, fi102, v9)
                                var v104 uint8 = 42
                                _ = v104
                            }
                            v2 += 32767
                            _ = fi102
                            _ = v103
                        }
                        _ = v101
                    } else {
                        v2 = (v2 | v2)
                    }
                }
                v9 = v9
            }
        case "\n\t":
            switch v6 {
            case "a":
                if (!(0 == -643)) {
                    var v105 uint8 = 97
                    switch -82 {
                    case 9223372036854775807:
                        var v106 int8 = -19
                        _ = v106
                    case 128:
                        if ((0.1 * 0.1) != 1e10) {
                            switch (v3 & v3) {
                            case 9223372036854775807:
                                var v107 string = v6
                                _ = v107
                            default:
                                var v108 uint16 = 42
                                _ = v108
                            }
                        }
                    case 32768:
                        var v109 int8 = 127
                        _ = v109
                    default:
                        if (v9 || (103 > 100)) {
                            if (922 < 0) {
                                if (-435 <= (-(v3))) {
                                    v105 = (^(v105))
                                    switch v3 {
                                    case -2147483648:
                                        var v110 int = int(v4)
                                        _ = v110
                                    case 4294967295:
                                        var v111 uint32 = 0
                                        _ = v111
                                    default:
                                        var v112 float32 = -2.25
                                        _ = v112
                                    }
                                } else {
                                    fmt.Println(v3, v105, v3)
                                    v113 := GenS8{F0: 730, F1: v3}
                                    _ = v113
                                }
                                fmt.Println(v2)
                            } else {
                                if v9 {
                                    var v114 int64 = -9223372036854775808
                                    delete(v1, "a")
                                    for fi115 := 0; fi115 < 20; fi115++ {
                                        delete(v1, v6)
                                        _ = fi115
                                    }
                                    _ = v114
                                } else {
                                    switch v6 {
                                    case "hello":
                                        v5 = append(v5, ((v6 + v6) + v6))
                                    default:
                                        var v116 uint = (^(v4))
                                        _ = v116
                                    }
                                    for fi117 := 0; fi117 < 8; fi117++ {
                                        delete(v1, "foo bar")
                                        fmt.Println(v105, v3)
                                        _ = fi117
                                    }
                                    v9 = (!v9)
                                }
                                for fi118 := 0; fi118 < 10; fi118++ {
                                    v3 += genFn0(1, (fi118 << (((^(v4))) & 63)))
                                    _ = fi118
                                }
                            }
                            var v119 uint = (v4 + (^(v4)))
                            _ = v119
                        } else {
                            v120 := GenS10{F0: 0, F1: (^(v3))}
                            var v121 uint64 = 18446744073709551615
                            _ = v120
                            _ = v121
                        }
                    }
                    _ = v105
                } else {
                    var v122 string = ("\n\t" + v6)
                    _ = v122
                }
            case "foo bar":
                for fi123 := 0; fi123 < 7; fi123++ {
                    if ((v9 && v9) && v9) {
                        for fi124 := 0; fi124 < 11; fi124++ {
                            if (v6 > (v6 + v6)) {
                                if (v9 && false) {
                                    var v125 uint = 255
                                    var v126 uint32 = 4294967295
                                    _ = v125
                                    _ = v126
                                } else {
                                    if v9 {
                                        delete(v1, "a")
                                        var v127 float32 = -2.25
                                        var v128 bool = true
                                        _ = v127
                                        _ = v128
                                    }
                                    if ((0.0 * ((float64(((1e10 / 3.14159) + 0.0)) - 1.5) + (1e10 - ((3.14159 - 1.5) - float64(1.5))))) >= float64(float32(-2.25))) {
                                        v129 := GenS5{F0: 32767, F1: 9223372036854775807, F2: (fi124 + 32768)}
                                        _ = v129
                                    } else {
                                        if v9 {
                                            v3 = -279
                                            var v130 uint64 = 42
                                            v3 = genFn0(42, genFn0(1, 9223372036854775807))
                                            fmt.Println(v3)
                                            fmt.Println(fi123, v3)
                                            if v9 {
                                                switch fi123 {
                                                case 42:
                                                    v131 := GenS12{F0: 85, F1: 42}
                                                    _ = v131
                                                }
                                            }
                                            _ = v130
                                        }
                                    }
                                    fmt.Println(fi124)
                                }
                                fmt.Println(v6, v3)
                                var v132 uint32 = 1
                                v133 := GenS3{F0: (fi124 - v3), F1: 784, F2: "\n\t"}
                                v132 -= v132
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v134 float32 = (0.1 + 3.14159)
                                _ = v132
                                _ = v133
                                _ = v134
                            } else {
                                v3 = 1
                            }
                            _ = fi124
                        }
                        var v135 uint16 = v2
                        var v136 int64 = 9223372036854775807
                        _ = v135
                        _ = v136
                    }
                    _ = fi123
                }
            case "":
                var v137 uint8 = 255
                _ = v137
            }
        }
    }
    delete(v1, "\n\t")
    v2 += (v2 / ((v2) | 1))
    for fi138 := 0; fi138 < 3; fi138++ {
        v139 := v1[""]
        _ = fi138
        _ = v139
    }
    _ = len(v1)
    var v140 float32 = (((float32(0.1) * 0.1) - (1e10 * -2.25)) * float32((1e10 / 1.5)))
    for fi141 := 0; fi141 < 20; fi141++ {
        var v142 int16 = 894
        _ = fi141
        _ = v142
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v140
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
