package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 uint; F1 uint8 }
type GenS9 struct { F0 string; F1 uint32; F2 int }

func main() {
    var v1 bool = false
    var v2 float32 = 1e10
    var v3 int64 = 42
    v4 := []int{}
    v5 := []int{}
    var v6 int = 32767
    fmt.Println(v1)
    var v7 uint32 = 65535
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int8 = 25
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 int = int(v7)
            v11 += fi10
            _ = fi10
            _ = v11
        }
        v5 = append(v5, (fi8 + 42))
        _ = fi8
        _ = v9
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 uint32 = v7
        v14 := GenS1{F0: -2147483648}
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        if (-107 == 0) {
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 int8 = 42
                v7 += 2147483648
                v7 = v7
                _ = fi16
                _ = v17
            }
            var v18 int8 = 23
            for fi19 := 0; fi19 < 11; fi19++ {
                v1 = v1
                if ((v7 & (v7 | (^(v7)))) > (v7 * v7)) {
                    v20 := GenS1{F0: -2147483648}
                    _ = v20
                } else {
                    v5 = append(v5, 32768)
                    if true {
                        if len(v5) > 0 {
                            v5[0] = (fi15 % ((fi15) | 1))
                        }
                    }
                    if (622 <= 18446744073709551615) {
                        var v21 bool = (!(true || false))
                        v22 := GenS1{F0: 42}
                        _ = v21
                        _ = v22
                    }
                }
                switch ("foo bar" + "\n\t") {
                case "hello":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "a":
                    if (!v1) {
                        if (!(1e10 > ((1.5 - 1.5) / (0.0 - (3.14159 + ((0.1 * 0.1) - 0.0)))))) {
                            var v23 float32 = (v2 * v2)
                            fmt.Println(fi15, v7, v1)
                            _ = v23
                        }
                        var v24 int16 = 1000
                        fmt.Println(v24, fi15, v18)
                        _ = v24
                    } else {
                        if v1 {
                            v4 = append(v4, v6)
                        } else {
                            v2 += v2
                            if (v3 != v3) {
                                v2 = v2
                            }
                            v7 -= 2147483648
                            if (v2 >= (v2 - v2)) {
                                if v1 {
                                    v25 := GenS1{F0: 42}
                                    _ = v25
                                }
                            } else {
                                v1 = v1
                                var v26 uint32 = 457
                                _ = v26
                            }
                        }
                        var v27 uint16 = 2
                        _ = v27
                    }
                case "":
                    switch 32767 {
                    case 127:
                        var v28 bool = (105 <= 1)
                        _ = v28
                    }
                    var v29 uint16 = 966
                    _ = len(v4)
                    _ = v29
                }
                var v30 string = "a"
                fmt.Println(v7)
                _ = fi19
                _ = v30
            }
            _ = v18
        }
        _ = fi15
    }
    switch "foo bar" {
    case "\n\t":
        var v31 float64 = 0.0
        _ = v31
    case "a":
        for fi32 := 0; fi32 < 12; fi32++ {
            for fi33 := 0; fi33 < 4; fi33++ {
                fmt.Println(v6, v2, v2)
                _ = fi33
            }
            switch -154 {
            case 32768:
                v4 = append(v4, (fi32 & (^(fi32))))
            case 127:
                if (v1 || (0.0 != (1.5 + 1e10))) {
                    for fi34 := 0; fi34 < 2; fi34++ {
                        v1 = (!(v1 || (32767 >= 990)))
                        _ = fi34
                    }
                }
            default:
                v2 -= -2.25
            }
            _ = fi32
        }
    default:
        v1 = v1
    }
    if v1 {
        _ = len(v4)
    }
    switch -89 {
    case 9223372036854775807:
        fmt.Println(v3, v7, v2)
    case 32768:
        var v35 int8 = 127
        _ = v35
    }
    v36 := GenS1{F0: -2147483648}
    switch "a" {
    case "hello":
        switch "foo bar" {
        case "":
            if (0 >= 0) {
                if v1 {
                    if (v1 && (v7 != 4294967295)) {
                        for fi37 := 0; fi37 < 17; fi37++ {
                            var v38 float32 = 1e10
                            var v39 int32 = -2147483648
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        v40 := GenS2{F0: 311, F1: (("" == "hello") && (!true))}
                        v2 = float32(((1.5 + ((((3.14159 - 0.1) - 0.0) / (1e10 + 1e10)) * float64(((float32((((0.1 * ((0.1 * 0.1) / 1.5)) - -2.25) / 1.0)) + float32(3.14159)) - 0.1)))) - 1.5))
                        _ = v40
                    }
                    v3 = 576
                    v6 = (v6 % ((128) | 1))
                } else {
                    if true {
                        fmt.Println(v7, v1)
                        switch "a" {
                        case "a":
                            if false {
                                v3 = (v3 & v3)
                                v2 = 3.14159
                                if (!(-1 >= 0)) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            }
                        default:
                            var v41 int = v6
                            _ = v41
                        }
                        fmt.Println(v7, v1, v7)
                    } else {
                        if false {
                            var v42 uint64 = 722
                            for fi43 := 0; fi43 < 1; fi43++ {
                                var v44 float64 = (1.5 / 1.5)
                                _ = fi43
                                _ = v44
                            }
                            _ = v42
                        } else {
                            v7 -= (^(v7))
                        }
                        var v45 int8 = -1
                        _ = v45
                    }
                }
            }
        case "foo bar":
            var v46 int64 = -9223372036854775808
            _ = v46
        }
    }
    if (!true) {
        var v47 uint64 = 708
        v4 = append(v4, v6)
        _ = v47
    } else {
        switch (("hello" + "") + "foo bar") {
        case "":
            var v48 float32 = v2
            _ = v48
        case "foo bar":
            if false {
                if false {
                    if v1 {
                        for fi49 := 0; fi49 < 5; fi49++ {
                            v2 = 0.1
                            var v50 int = (^(v6))
                            _ = fi49
                            _ = v50
                        }
                        if (2 != 255) {
                            if false {
                                var v51 uint = 1
                                if (false || ((v6 == (v6 & (v6 + v6))) || true)) {
                                    var v52 uint16 = 65535
                                    _ = v52
                                } else {
                                    for fi53 := 0; fi53 < 18; fi53++ {
                                        _ = len(v5)
                                        _ = fi53
                                    }
                                    v7 = v7
                                    if v1 {
                                        for fi54 := 0; fi54 < 17; fi54++ {
                                            if len(v5) > 0 {
                                                v5[0] = fi54
                                            }
                                            _ = fi54
                                        }
                                    }
                                }
                                _ = v51
                            } else {
                                v55 := GenS3{F0: 1, F1: -1}
                                switch "" {
                                case "foo bar":
                                    v7 += v7
                                case "":
                                    if (255 < 128) {
                                        if false {
                                            for fi56 := 0; fi56 < 15; fi56++ {
                                                var v57 int8 = -1
                                                _ = fi56
                                                _ = v57
                                            }
                                            v58 := GenS3{F0: 24, F1: 100}
                                            v7 = (v7 - 2147483648)
                                            _ = v58
                                        }
                                    } else {
                                        v3 -= (v3 + -537)
                                        var v59 uint16 = 0
                                        _ = v59
                                    }
                                default:
                                    var v60 uint64 = 18446744073709551615
                                    _ = v60
                                }
                                v61 := GenS4{F0: 78, F1: 2}
                                _ = v55
                                _ = v61
                            }
                            var v62 bool = v1
                            _ = v62
                        } else {
                            if v1 {
                                fmt.Println(v1, v3)
                                v4 = append(v4, (v6 / ((v6) | 1)))
                            } else {
                                var v63 int32 = -50
                                _ = v63
                            }
                            fmt.Println(v3, v6)
                        }
                        for fi64 := 0; fi64 < 16; fi64++ {
                            v7 = 378
                            if (42 != 939) {
                                var v65 uint32 = v7
                                if (!false) {
                                    fmt.Println(v65, v65)
                                    if false {
                                        var v66 int8 = 80
                                        v67 := GenS5{F0: 1}
                                        switch "\n\t" {
                                        case "hello":
                                            var v68 uint32 = 6
                                            _ = v68
                                        }
                                        _ = v66
                                        _ = v67
                                    }
                                }
                                switch (fi64 % (((v6 ^ 32768)) | 1)) {
                                case 9223372036854775807:
                                    v65 -= v65
                                case 0:
                                    var v69 string = "a"
                                    _ = v69
                                }
                                _ = v65
                            }
                            _ = fi64
                        }
                    } else {
                        var v70 float32 = v2
                        v71 := GenS2{F0: 18446744073709551615, F1: (v1 && (1000 > -1))}
                        var v72 int = v6
                        _ = v70
                        _ = v71
                        _ = v72
                    }
                    var v73 uint64 = 42
                    _ = v73
                } else {
                    for fi74 := 0; fi74 < 2; fi74++ {
                        var v75 float64 = (3.14159 * 3.14159)
                        _ = fi74
                        _ = v75
                    }
                    var v76 int16 = 60
                    if (!(float32((3.14159 + 1.5)) > v2)) {
                        for fi77 := 0; fi77 < 14; fi77++ {
                            if ((v7 <= 1) && ((3.14159 + 0.0) >= 0.1)) {
                                for fi78 := 0; fi78 < 17; fi78++ {
                                    switch ("hello" + "foo bar") {
                                    case "\n\t":
                                        v1 = false
                                    case "":
                                        var v79 int = -982
                                        _ = v79
                                    }
                                    switch "hello" {
                                    case "":
                                        _ = len(v4)
                                    case "foo bar":
                                        var v80 bool = (!false)
                                        _ = v80
                                    default:
                                        v81 := GenS3{F0: 0, F1: -128}
                                        _ = v81
                                    }
                                    _ = fi78
                                }
                                v82 := GenS6{F0: -1}
                                var v83 int = fi77
                                _ = v82
                                _ = v83
                            }
                            var v84 float32 = (v2 - -2.25)
                            _ = fi77
                            _ = v84
                        }
                    } else {
                        var v85 bool = true
                        var v86 int8 = -8
                        _ = len(v4)
                        _ = v85
                        _ = v86
                    }
                    _ = v76
                }
            } else {
                var v87 string = ((("\n\t" + "hello") + ("hello" + ("" + ("hello" + ("" + ("foo bar" + ("a" + ("a" + "a")))))))) + ("" + "a"))
                _ = v87
            }
        default:
            if v1 {
                switch v6 {
                case -2147483648:
                    var v88 string = "a"
                    _ = v88
                case 2147483647:
                    if (!true) {
                        var v89 bool = v1
                        switch "a" {
                        case "":
                            var v90 int8 = 1
                            _ = v90
                        default:
                            v91 := GenS7{F0: 441, F1: (9223372036854775807 >= v6)}
                            _ = v91
                        }
                        var v92 int8 = 100
                        _ = v89
                        _ = v92
                    } else {
                        switch v6 {
                        case 32767:
                            if (!v1) {
                                var v93 int32 = 1
                                for fi94 := 0; fi94 < 3; fi94++ {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    for fi95 := 0; fi95 < 19; fi95++ {
                                        var v96 int16 = 42
                                        if true {
                                            v97 := GenS5{F0: 864}
                                            var v98 float32 = v2
                                            if (v1 || v1) {
                                                v7 += 627
                                                v93 = (v93 & 0)
                                                v96 = v96
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            _ = v97
                                            _ = v98
                                        } else {
                                            fmt.Println(v6, v7)
                                            fmt.Println(v1, fi94, v1)
                                            v6 = fi95
                                            fmt.Println(v3)
                                        }
                                        v1 = ((v6 < (fi95 / ((32768) | 1))) && (18446744073709551615 <= 18446744073709551615))
                                        if v1 {
                                            v99 := GenS7{F0: (fi94 / ((42) | 1)), F1: (1e10 == 1e10)}
                                            var v100 uint16 = 2
                                            _ = v100
                                            _ = v99
                                        } else {
                                            var v101 int32 = (v93 * -1)
                                            _ = len(v4)
                                            _ = v101
                                        }
                                        _ = fi95
                                        _ = v96
                                    }
                                    var v102 uint16 = 5
                                    _ = fi94
                                    _ = v102
                                }
                                v93 = v93
                                _ = v93
                            }
                        case -9223372036854775808:
                            for fi103 := 0; fi103 < 14; fi103++ {
                                v7 += 2147483648
                                _ = fi103
                            }
                        case -763:
                            v104 := GenS8{F0: 2, F1: 1}
                            _ = v104
                        default:
                            if (65535 > 729) {
                                v6 += (^(v6))
                                for fi105 := 0; fi105 < 20; fi105++ {
                                    var v106 uint32 = v7
                                    _ = fi105
                                    _ = v106
                                }
                                switch v6 {
                                case -695:
                                    for fi107 := 0; fi107 < 11; fi107++ {
                                        var v108 uint16 = 426
                                        v1 = v1
                                        v109 := GenS2{F0: 1, F1: v1}
                                        _ = fi107
                                        _ = v108
                                        _ = v109
                                    }
                                }
                            } else {
                                v6 = v6
                                for fi110 := 0; fi110 < 6; fi110++ {
                                    v7 += v7
                                    fmt.Println(v7, v6)
                                    _ = fi110
                                }
                            }
                        }
                        if false {
                            var v111 float32 = v2
                            if (v1 && (1 == -2147483648)) {
                                v112 := GenS5{F0: 1000}
                                if v1 {
                                    fmt.Println(v7, v7)
                                }
                                _ = v112
                            } else {
                                v2 += v2
                                var v113 float32 = (0.0 / 1e10)
                                for fi114 := 0; fi114 < 14; fi114++ {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v1 = (!v1)
                                    case "":
                                        for fi115 := 0; fi115 < 14; fi115++ {
                                            var v116 uint32 = v7
                                            var v117 uint16 = 1
                                            var v118 uint8 = 198
                                            _ = fi115
                                            _ = v116
                                            _ = v117
                                            _ = v118
                                        }
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = (v6 - fi114)
                                        }
                                    }
                                    var v119 uint = 255
                                    if (42 == 647) {
                                        switch v6 {
                                        case -767:
                                            var v120 uint32 = (^(v7))
                                            _ = v120
                                        case -9223372036854775808:
                                            var v121 uint16 = 2
                                            _ = v121
                                        case -251:
                                            var v122 float64 = float64(1.5)
                                            _ = v122
                                        default:
                                            if true {
                                                for fi123 := 0; fi123 < 5; fi123++ {
                                                    v111 -= float32(1.5)
                                                    var v124 int8 = 127
                                                    _ = fi123
                                                    _ = v124
                                                }
                                                var v125 string = ""
                                                _ = v125
                                            } else {
                                                var v126 uint = 65535
                                                _ = v126
                                            }
                                        }
                                        fmt.Println(fi114, v111, v113)
                                        _ = len(v4)
                                    } else {
                                        var v127 uint8 = 0
                                        if (1000 == 32767) {
                                            v128 := GenS6{F0: -32768}
                                            _ = v128
                                        }
                                        if false {
                                            var v129 string = "hello"
                                            _ = v129
                                        } else {
                                            if true {
                                                switch v6 {
                                                case -668:
                                                    if len(v4) > 0 {
                                                        v4[0] = 127
                                                    }
                                                case 9223372036854775807:
                                                    var v130 uint8 = (^(v127))
                                                    _ = v130
                                                case 127:
                                                    var v131 uint = v119
                                                    _ = v131
                                                default:
                                                    fmt.Println(fi114, v113, v7)
                                                }
                                                _ = len(v4)
                                                switch (^(fi114)) {
                                                case -438:
                                                    v111 -= 0.0
                                                }
                                            } else {
                                                if (-332 <= 2147483647) {
                                                    v132 := GenS9{F0: "", F1: 65535, F2: 0}
                                                    _ = v132
                                                } else {
                                                    if len(v5) > 0 {
                                                        v5[0] = fi114
                                                    }
                                                }
                                            }
                                            var v133 uint16 = 32767
                                            _ = v133
                                        }
                                        _ = v127
                                    }
                                    _ = fi114
                                    _ = v119
                                }
                                _ = v113
                            }
                            _ = v111
                        } else {
                            v134 := GenS8{F0: 4294967295, F1: 231}
                            _ = v134
                        }
                    }
                }
                _ = len(v4)
                if len(v5) > 0 {
                    v5[0] = (v6 & 42)
                }
            }
        }
    }
    v3 -= 1
    _ = len(v5)
    v135 := GenS4{F0: 32767, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v135
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
