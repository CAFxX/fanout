package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 bool; F2 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 uint16; F2 int8 }
type GenS8 struct { F0 int8; F1 uint8; F2 uint }
type GenS9 struct { F0 int; F1 uint16; F2 uint8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint16; F1 uint8 }
type GenS12 struct { F0 string; F1 int16; F2 uint32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 bool; F1 uint }

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]string)
    var v3 uint32 = 2
    var v4 uint32 = v3
    var v5 bool = ((477 < 457) || true)
    fmt.Println(v3)
    if (("foo bar" >= (("\n\t" + ("hello" + ("hello" + ("" + "hello")))) + "\n\t")) || (!false)) {
        fmt.Println(v1, v1)
        var v6 float32 = 1.5
        for fi7 := 0; fi7 < 10; fi7++ {
            if false {
                for fi8 := 0; fi8 < 1; fi8++ {
                    v1 += (v1 | (v1 ^ 0))
                    delete(v2, (v3 | v3))
                    var v9 int16 = 572
                    _ = len(v2)
                    delete(v2, v3)
                    var v10 int16 = 1000
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                delete(v2, v3)
            } else {
                var v11 uint16 = v1
                if (251 != uint8(v4)) {
                    v11 += 65535
                    v12 := GenS1{F0: v3, F1: v5}
                    _ = v12
                }
                v13 := GenS2{F0: (v1 / ((2) | 1)), F1: ""}
                _ = v11
                _ = v13
            }
            v4 += (v4 & v3)
            v14 := GenS2{F0: (^(v1)), F1: "foo bar"}
            _ = fi7
            _ = v14
        }
        _ = v6
    } else {
        switch -241 {
        case -1:
            v5 = true
        case 32768:
            if v5 {
                if (false || (23 == 9223372036854775807)) {
                    var v15 float32 = 1.5
                    v16 := GenS3{F0: -939}
                    _ = v15
                    _ = v16
                } else {
                    fmt.Println(v5, v1, v5)
                    v3 += (v3 % ((v3) | 1))
                }
                v17 := []int{-2147483648}
                for fi18 := 0; fi18 < 7; fi18++ {
                    if (-1 <= 756) {
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            if ((1000 <= -1) || true) {
                                v5 = false
                            }
                        case "a":
                            v19 := v2[(v3 << ((uint(v4)) & 63))]
                            _ = v19
                        case "hello":
                            v5 = (!v5)
                        default:
                            if (2147483647 >= 42) {
                                if v5 {
                                    delete(v2, (^(v4)))
                                    switch 42 {
                                    case -8:
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                    case 127:
                                        for fi20 := 0; fi20 < 9; fi20++ {
                                            fmt.Println(v1, v4)
                                            var v21 int8 = -49
                                            v4 += 2147483648
                                            _ = fi20
                                            _ = v21
                                        }
                                    default:
                                        v4 = (v4 ^ (^(v4)))
                                    }
                                    if true {
                                        fmt.Println(v3, v3, fi18)
                                        v4 = v3
                                    } else {
                                        v1 = (v1 ^ v1)
                                        v4 = 0
                                    }
                                } else {
                                    v22 := GenS3{F0: 74}
                                    _ = v22
                                }
                                fmt.Println(v4)
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(v4, v3)
                                case "foo bar":
                                    if ((!((^(v1)) != v1)) && true) {
                                        v23 := GenS4{F0: 2}
                                        _ = v23
                                    }
                                case "":
                                    if ((v4 | v4) >= v4) {
                                        delete(v2, 403)
                                    }
                                }
                            }
                        }
                    }
                    _ = fi18
                }
                _ = v17
            } else {
                v3 = (v4 ^ v3)
                switch 4294967295 {
                case 206:
                    v24 := []int{-86, 623, -946}
                    _ = v24
                default:
                    var v25 bool = (v1 >= (v1 - (v1 % (((v1 & (v1 + (^(v1))))) | 1))))
                    _ = v25
                }
                v1 = 227
            }
        default:
            v26 := GenS4{F0: 0}
            _ = v26
        }
    }
    fmt.Println(v5, v3)
    v27 := GenS1{F0: (^(v3)), F1: v5}
    v28 := []int{}
    v29 := GenS1{F0: v4, F1: ((!(v5 || (!true))) && (!(!("\n\t" == "foo bar"))))}
    v3 = (v3 | (v3 - (v4 * 2)))
    if v5 {
        v1 -= 2
        v30 := v2[v3]
        _ = v30
    } else {
        v3 = (v4 * v3)
    }
    v31 := v2[42]
    if len(v28) > 0 {
        v28[0] = 128
    }
    if (9223372036854775807 == -2147483648) {
        fmt.Println(v1)
        v5 = (-128 > 100)
        if v5 {
            if (!(!true)) {
                v5 = v5
                v32 := GenS3{F0: 42}
                fmt.Println(v31)
                _ = v32
            } else {
                if (false || false) {
                    if v5 {
                        if false {
                            if ((-29 < -1) && true) {
                                v33 := GenS5{F0: 914, F1: false, F2: 0}
                                for fi34 := 0; fi34 < 19; fi34++ {
                                    switch 32767 {
                                    case -293:
                                        if ((v1 * (v1 << ((584) & 63))) < 0) {
                                            var v35 uint32 = (v3 + v3)
                                            v5 = v5
                                            var v36 int64 = 9223372036854775807
                                            _ = v35
                                            _ = v36
                                        }
                                    case -9223372036854775808:
                                        var v37 uint64 = 4294967295
                                        _ = v37
                                    case 42:
                                        if v5 {
                                            v5 = true
                                            for fi38 := 0; fi38 < 20; fi38++ {
                                                v39 := GenS6{F0: -1}
                                                v28 = append(v28, fi38)
                                                v5 = v5
                                                _ = fi38
                                                _ = v39
                                            }
                                            var v40 uint8 = uint8(v4)
                                            _ = v40
                                        }
                                    default:
                                        v1 = v1
                                    }
                                    v31 = (((v31 + ("a" + "foo bar")) + v31) + (((v31 + "a") + "a") + v31))
                                    for fi41 := 0; fi41 < 8; fi41++ {
                                        v1 = (v1 % ((v1) | 1))
                                        fmt.Println(v5, v4)
                                        if (((v1 - v1) > 371) && (9223372036854775807 >= 151)) {
                                            fmt.Println(v4, v31)
                                            fmt.Println(v4, v31, v3)
                                        }
                                        v1 += v1
                                        v28 = append(v28, (fi34 & fi41))
                                        _ = fi41
                                    }
                                    _ = fi34
                                }
                                if v5 {
                                    v42 := GenS2{F0: v1, F1: (v31 + ("" + ("a" + "")))}
                                    fmt.Println(v31, v3)
                                    v5 = (!v5)
                                    _ = v42
                                }
                                _ = v33
                            }
                        } else {
                            v1 = 42
                        }
                        switch v31 {
                        case "foo bar":
                            if ((!v5) && (-261 != -9223372036854775808)) {
                                if (v5 && v5) {
                                    v31 = v31
                                    var v43 int32 = -946
                                    _ = v43
                                } else {
                                    if true {
                                        v3 += 0
                                    }
                                    v3 += v3
                                }
                            } else {
                                v44 := GenS6{F0: 0}
                                v5 = (1.5 >= (0.1 / 1.5))
                                v1 = v1
                                _ = v44
                            }
                        case "hello":
                            v3 -= v4
                        }
                        for fi45 := 0; fi45 < 20; fi45++ {
                            v2[(v4 + v3)] = v31
                            fmt.Println(v31)
                            if (!false) {
                                switch (v31 + "hello") {
                                case "a":
                                    for fi46 := 0; fi46 < 20; fi46++ {
                                        v31 = (v31 + "a")
                                        v3 = v4
                                        _ = len(v28)
                                        _ = fi46
                                    }
                                case "":
                                    var v47 float64 = 1e10
                                    _ = v47
                                }
                                var v48 int = fi45
                                _ = v48
                            } else {
                                if v5 {
                                    if (!v5) {
                                        switch (^(fi45)) {
                                        case 32767:
                                            v49 := GenS7{F0: "hello", F1: v1, F2: 37}
                                            _ = v49
                                        default:
                                            v50 := GenS4{F0: 2}
                                            _ = v50
                                        }
                                        for fi51 := 0; fi51 < 1; fi51++ {
                                            var v52 int8 = 127
                                            _ = fi51
                                            _ = v52
                                        }
                                        if (!v5) {
                                            if false {
                                                v5 = (-32768 <= 42)
                                            } else {
                                                v31 = "\n\t"
                                                v1 = (v1 | (v1 % (((v1 | v1)) | 1)))
                                            }
                                        }
                                    } else {
                                        v28 = append(v28, (fi45 | fi45))
                                    }
                                } else {
                                    fmt.Println(v1)
                                    switch v31 {
                                    case "a":
                                        switch (^(fi45)) {
                                        case -9223372036854775808:
                                            if len(v28) >= 2 {
                                                v28 = v28[1:len(v28)]
                                            }
                                        case 419:
                                            if v5 {
                                                if true {
                                                    for fi53 := 0; fi53 < 19; fi53++ {
                                                        if (v5 || ("" > (v31 + ("hello" + v31)))) {
                                                            v54 := GenS1{F0: v3, F1: v5}
                                                            v55 := GenS8{F0: -1, F1: 1, F2: 666}
                                                            _ = v54
                                                            _ = v55
                                                        } else {
                                                            var v56 uint16 = 989
                                                            var v57 bool = false
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        _ = fi53
                                                    }
                                                    for fi58 := 0; fi58 < 11; fi58++ {
                                                        var v59 uint = 4294967295
                                                        if v5 {
                                                            v59 -= (v59 | (v59 * uint(v1)))
                                                            v59 = (v59 << ((v59) & 63))
                                                        } else {
                                                            if v5 {
                                                                fmt.Println(v1, fi45, v31)
                                                            } else {
                                                                v60 := GenS9{F0: (fi58 % ((fi45) | 1)), F1: v1, F2: 0}
                                                                v61 := GenS10{F0: v31}
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    v62 := v2[v4]
                                                                    _ = v62
                                                                default:
                                                                    if len(v28) > 0 {
                                                                        v28[0] = (^(fi58))
                                                                    }
                                                                }
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                            v5 = ((!((4294967295 < fi58) && (!true))) || true)
                                                            _ = len(v2)
                                                        }
                                                        var v63 uint = v59
                                                        delete(v2, v3)
                                                        _ = fi58
                                                        _ = v59
                                                        _ = v63
                                                    }
                                                } else {
                                                    for fi64 := 0; fi64 < 20; fi64++ {
                                                        var v65 string = "a"
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v66 int8 = 126
                                                            _ = v66
                                                        }
                                                        _ = fi64
                                                        _ = v65
                                                    }
                                                }
                                                v3 += (v4 & 2147483648)
                                            }
                                        }
                                    case "":
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                    case "hello":
                                        switch 9223372036854775807 {
                                        case 748:
                                            var v67 int64 = 1
                                            _ = v67
                                        case 0:
                                            fmt.Println(v3, v5)
                                        default:
                                            fmt.Println(v31, v1)
                                        }
                                    }
                                    v3 = v4
                                }
                                switch (fi45 - -9223372036854775808) {
                                case 1:
                                    if (int64(v4) >= -1) {
                                        switch ("\n\t" + ((("hello" + ("a" + (v31 + v31))) + v31) + ("\n\t" + (((v31 + (v31 + v31)) + v31) + "hello")))) {
                                        case "foo bar":
                                            var v68 int32 = -2147483648
                                            _ = v68
                                        case "a":
                                            var v69 string = v31
                                            _ = v69
                                        case "":
                                            var v70 int32 = 1
                                            _ = v70
                                        default:
                                            var v71 uint8 = 2
                                            _ = v71
                                        }
                                    }
                                case 127:
                                    var v72 uint64 = 697
                                    _ = v72
                                case 0:
                                    var v73 uint16 = v1
                                    _ = v73
                                }
                            }
                            _ = fi45
                        }
                    }
                    v74 := GenS6{F0: 0}
                    fmt.Println(v1, v4, v3)
                    _ = v74
                } else {
                    fmt.Println(v3)
                    var v75 bool = v5
                    _ = v75
                }
            }
            v3 -= v4
        } else {
            v76 := GenS7{F0: v31, F1: 0, F2: -80}
            v3 += (v3 - v3)
            _ = len(v2)
            _ = v76
        }
    }
    fmt.Println(v31, v3)
    if v5 {
        for fi77 := 0; fi77 < 4; fi77++ {
            fmt.Println(fi77, v3, v5)
            v78 := GenS7{F0: v31, F1: 223, F2: 108}
            _ = fi77
            _ = v78
        }
        delete(v2, v3)
    } else {
        _ = len(v2)
    }
    if v5 {
        v3 += v3
        if false {
            fmt.Println(v31)
            for fi79 := 0; fi79 < 11; fi79++ {
                for fi80 := 0; fi80 < 12; fi80++ {
                    switch (v31 + v31) {
                    case "a":
                        if v5 {
                            v31 = (v31 + ("\n\t" + (((v31 + "") + ((v31 + ("\n\t" + "a")) + "a")) + "a")))
                            v4 = 538
                        } else {
                            v81 := GenS11{F0: (v1 & v1), F1: 128}
                            v5 = (42 < 42)
                            _ = v81
                        }
                    case "foo bar":
                        var v82 string = (v31 + "\n\t")
                        _ = v82
                    default:
                        v83 := GenS1{F0: v3, F1: (((float32(1.5) - 1.5) / 1e10) < (((3.14159 / 0.1) / 1.0) / (1e10 * 0.1)))}
                        _ = v83
                    }
                    switch "\n\t" {
                    case "a":
                        switch "hello" {
                        case "":
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                        case "foo bar":
                            if (-2.25 >= (-2.25 / 1.0)) {
                                var v84 string = ((((v31 + v31) + (v31 + ("\n\t" + "a"))) + "foo bar") + v31)
                                v85 := GenS10{F0: ((("foo bar" + v31) + "") + v31)}
                                v4 = (v4 % ((v3) | 1))
                                _ = v84
                                _ = v85
                            } else {
                                var v86 float64 = (1e10 + 3.14159)
                                _ = v86
                            }
                            var v87 float64 = 0.1
                            _ = v87
                        default:
                            v88 := GenS6{F0: 0}
                            _ = v88
                        }
                    case "foo bar":
                        if ((fi80 / ((fi80) | 1)) >= fi80) {
                            if v5 {
                                v89 := GenS10{F0: ("" + (v31 + ((((v31 + "") + (("" + v31) + v31)) + v31) + (v31 + v31))))}
                                var v90 int16 = -673
                                v2[760] = v31
                                _ = v89
                                _ = v90
                            } else {
                                if true {
                                    v91 := GenS12{F0: v31, F1: 1, F2: (v3 & (v3 % ((v4) | 1)))}
                                    var v92 bool = (float32(((0.1 * (3.14159 - 3.14159)) + 1e10)) >= float32((1e10 + (1.5 / 1e10))))
                                    var v93 uint32 = v3
                                    _ = v91
                                    _ = v92
                                    _ = v93
                                }
                                v94 := v2[(^(v3))]
                                v95 := GenS4{F0: 1}
                                _ = v94
                                _ = v95
                            }
                        } else {
                            switch (fi79 * fi79) {
                            case 4294967295:
                                if len(v28) > 0 {
                                    v28[0] = -236
                                }
                            default:
                                v5 = v5
                            }
                        }
                    default:
                        switch fi79 {
                        case -981:
                            switch v31 {
                            case "hello":
                                delete(v2, v4)
                            case "a":
                                if false {
                                    v96 := GenS13{F0: -362}
                                    _ = v96
                                } else {
                                    var v97 int8 = -128
                                    fmt.Println(v1, v1, v5)
                                    switch "a" {
                                    case "a":
                                        switch (v31 + "") {
                                        case "hello":
                                            v98 := GenS5{F0: 1, F1: (!v5), F2: 181}
                                            _ = v98
                                        }
                                    case "hello":
                                        switch "" {
                                        case "a":
                                            if false {
                                                v99 := GenS14{F0: v97}
                                                var v100 float64 = (0.0 + 0.0)
                                                _ = v100
                                                _ = v99
                                            } else {
                                                if true {
                                                    var v101 uint32 = (v4 >> ((0) & 63))
                                                    v3 += uint32(v1)
                                                    if (v1 <= 669) {
                                                        if false {
                                                            switch fi79 {
                                                            case -1:
                                                                var v102 uint32 = v4
                                                                _ = v102
                                                            }
                                                        } else {
                                                            switch v31 {
                                                            case "foo bar":
                                                                var v103 uint = 909
                                                                switch v31 {
                                                                case "hello":
                                                                    if (v5 && false) {
                                                                        v2[(v3 << (((^(v103))) & 63))] = "\n\t"
                                                                    }
                                                                case "a":
                                                                    if (65535 > v4) {
                                                                        var v104 int = fi80
                                                                        v104 = v104
                                                                        _ = v104
                                                                    }
                                                                }
                                                                _ = v103
                                                            case "hello":
                                                                if (0 <= 255) {
                                                                    v31 = ""
                                                                }
                                                            }
                                                            var v105 uint8 = 45
                                                            v106 := GenS13{F0: 756}
                                                            v107 := v2[(v4 * (v3 >> ((18446744073709551615) & 63)))]
                                                            _ = v105
                                                            _ = v106
                                                            _ = v107
                                                        }
                                                        var v108 uint = 18446744073709551615
                                                        v31 = ((("\n\t" + v31) + (v31 + "")) + v31)
                                                        _ = v108
                                                    } else {
                                                        _ = len(v2)
                                                    }
                                                    _ = v101
                                                }
                                            }
                                        case "foo bar":
                                            var v109 int32 = -278
                                            _ = v109
                                        case "\n\t":
                                            var v110 float64 = 0.0
                                            _ = v110
                                        }
                                    default:
                                        var v111 uint8 = 1
                                        _ = v111
                                    }
                                    _ = v97
                                }
                            case "\n\t":
                                v5 = true
                            default:
                                var v112 int64 = 0
                                _ = v112
                            }
                        case 0:
                            var v113 int32 = -1
                            _ = v113
                        case 584:
                            v114 := GenS15{F0: true, F1: 765}
                            v1 += 65535
                            if v5 {
                                fmt.Println(v4, fi79)
                            }
                            _ = v114
                        }
                    }
                    _ = fi80
                }
                var v115 int32 = 1
                _ = fi79
                _ = v115
            }
            if (3.14159 >= 3.14159) {
                v31 = v31
            } else {
                v3 += v3
                for fi116 := 0; fi116 < 11; fi116++ {
                    var v117 uint32 = 243
                    v3 += 0
                    _ = fi116
                    _ = v117
                }
            }
        } else {
            switch ("hello" + "hello") {
            case "\n\t":
                delete(v2, (v4 * v3))
            default:
                var v118 float64 = (1.5 / -2.25)
                _ = v118
            }
        }
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
