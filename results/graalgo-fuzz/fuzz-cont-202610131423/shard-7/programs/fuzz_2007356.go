package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int64; F2 int64 }
type GenS3 struct { F0 int32; F1 string; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS5 struct { F0 uint; F1 int; F2 uint32 }
type GenS6 struct { F0 uint32; F1 int32; F2 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 int; F2 uint32 }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 bool; F1 bool }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint; F1 uint64 }
type GenS16 struct { F0 bool }

func main() {
    var v1 uint8 = 128
    var v2 int64 = 453
    var v3 int16 = -1
    v4 := make(map[uint32]uint32)
    var v5 int64 = v2
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint16 = 762
        _ = fi6
        _ = v7
    }
    switch 42 {
    case 558:
        var v8 int16 = (v3 | v3)
        _ = v8
    case 32768:
        fmt.Println(v2, v3)
    }
    v9 := GenS1{F0: 0}
    if (true || (!(42 != 2147483647))) {
        var v10 int64 = v5
        v11 := GenS1{F0: 255}
        v12 := GenS1{F0: 0}
        _ = v10
        _ = v11
        _ = v12
    } else {
        v1 = v1
    }
    var v13 uint32 = 1
    v13 += v13
    delete(v4, (^(v13)))
    for fi14 := 0; fi14 < 12; fi14++ {
        if (!true) {
            if (621 <= 0) {
                v15 := GenS2{F0: (v3 * int16(fi14)), F1: 38, F2: (-(v5))}
                v16 := []int{(-(fi14)), -1, fi14, int(v3)}
                var v17 int32 = -113
                _ = v15
                _ = v16
                _ = v17
            } else {
                if (1 == v5) {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        fmt.Println(fi18, v5, v13)
                        v19 := v4[v13]
                        _ = fi18
                        _ = v19
                    }
                    v20 := GenS3{F0: 2147483647, F1: "foo bar", F2: 1}
                    _ = v20
                } else {
                    v21 := GenS4{F0: 65535, F1: 42, F2: v13}
                    if (305 <= 805) {
                        v22 := GenS1{F0: 2}
                        v23 := []int{fi14, (fi14 + (fi14 & (fi14 ^ (fi14 / ((int(v1)) | 1))))), fi14, fi14}
                        fmt.Println(v2, v2)
                        _ = v22
                        _ = v23
                    } else {
                        fmt.Println(v5)
                        var v24 int8 = -40
                        _ = v24
                    }
                    _ = v21
                }
                for fi25 := 0; fi25 < 19; fi25++ {
                    var v26 uint8 = v1
                    if true {
                        switch fi14 {
                        case 128:
                            var v27 int32 = -2147483648
                            v5 += (v5 / (((^(v5))) | 1))
                            v28 := []int{-18}
                            _ = v27
                            _ = v28
                        }
                        var v29 int32 = 1
                        v30 := GenS4{F0: 226, F1: 0, F2: (v13 >> ((uint(v1)) & 63))}
                        switch "\n\t" {
                        case "":
                            v31 := GenS1{F0: 6}
                            _ = v31
                        default:
                            v32 := GenS2{F0: (v3 + v3), F1: v5, F2: (v2 / (((v2 & v2)) | 1))}
                            _ = v32
                        }
                        if true {
                            v33 := GenS1{F0: 724}
                            var v34 bool = (((1 <= v3) && ((42 != -1) && false)) || (((1e10 / ((1e10 / 3.14159) / float64(((((-2.25 - 1.5) - float32(((0.0 * 1.5) - 1e10))) * (1.5 + 1e10)) * 1.5)))) * 3.14159) > 0.1))
                            switch "hello" {
                            case "foo bar":
                                if false {
                                    if (629 <= 1) {
                                        if (true && v34) {
                                            if ((!v34) && (1 > 4294967295)) {
                                                if (!v34) {
                                                    if v34 {
                                                        v29 = (-(v29))
                                                    }
                                                    var v35 uint8 = v26
                                                    v2 = (-(v5))
                                                    _ = v35
                                                }
                                                var v36 int16 = v3
                                                var v37 float64 = 0.1
                                                if false {
                                                    switch ((("foo bar" + "\n\t") + ("a" + "\n\t")) + "hello") {
                                                    case "":
                                                        if (!(!false)) {
                                                            v37 -= (v37 - v37)
                                                            if ((v34 || (239 >= (-(v36)))) || ((int(v3) >= fi25) && v34)) {
                                                                var v38 float64 = (float64((1.5 / 1e10)) - ((v37 / 0.1) / (v37 * (v37 + v37))))
                                                                switch (^(fi14)) {
                                                                case 32768:
                                                                    if (18446744073709551615 <= 0) {
                                                                        var v39 float64 = v38
                                                                        var v40 float64 = (v38 / (v38 / (v39 * v37)))
                                                                        v1 += v1
                                                                        _ = v39
                                                                        _ = v40
                                                                    }
                                                                case 838:
                                                                    var v41 uint16 = 928
                                                                    _ = v41
                                                                case 127:
                                                                    fmt.Println(fi14)
                                                                default:
                                                                    var v42 uint32 = (v13 % ((1) | 1))
                                                                    _ = v42
                                                                }
                                                                var v43 string = ("foo bar" + "hello")
                                                                var v44 uint64 = 18446744073709551615
                                                                v45 := GenS5{F0: 555, F1: (fi14 / (((-(fi25))) | 1)), F2: v13}
                                                                switch (fi14 << ((1) & 63)) {
                                                                case 32767:
                                                                    v34 = (!((1e10 / float32((((v38 / v38) / 1e10) + v38))) == (1.5 - 3.14159)))
                                                                default:
                                                                    if v34 {
                                                                        fmt.Println(v34)
                                                                    }
                                                                }
                                                                _ = v38
                                                                _ = v43
                                                                _ = v44
                                                                _ = v45
                                                            } else {
                                                                v26 = uint8(v13)
                                                            }
                                                        } else {
                                                            _ = len(v4)
                                                        }
                                                    case "\n\t":
                                                        v46 := GenS6{F0: (v13 / (((v13 * v13)) | 1)), F1: (v29 >> ((261) & 63)), F2: (v5 <= -880)}
                                                        _ = v46
                                                    }
                                                } else {
                                                    v47 := GenS6{F0: 787, F1: v29, F2: (("a" + "\n\t") > "a")}
                                                    _ = v47
                                                }
                                                fmt.Println(v34)
                                                _ = v36
                                                _ = v37
                                            }
                                        } else {
                                            fmt.Println(fi14, v29, v34)
                                            switch ("hello" + "hello") {
                                            case "":
                                                v5 = v5
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                v48 := []int{}
                                                _ = v48
                                            }
                                        }
                                        v3 -= (v3 % ((1000) | 1))
                                        v49 := v4[441]
                                        _ = v49
                                    } else {
                                        if v34 {
                                            v50 := v4[uint32(v1)]
                                            switch "\n\t" {
                                            case "foo bar":
                                                v13 -= 1
                                            case "a":
                                                if (true || v34) {
                                                    v51 := []int{788}
                                                    v51 = append(v51, (fi14 >> ((65535) & 63)))
                                                    _ = v51
                                                } else {
                                                    v52 := GenS6{F0: (^(v50)), F1: 103, F2: v34}
                                                    v5 = (v2 - (v5 - (v2 * (-(v2)))))
                                                    _ = v52
                                                }
                                            }
                                            _ = v50
                                        }
                                        switch 4294967295 {
                                        case 42:
                                            _ = len(v4)
                                        case 9223372036854775807:
                                            switch ("a" + ("\n\t" + "foo bar")) {
                                            case "a":
                                                v53 := GenS5{F0: 65535, F1: (^(fi25)), F2: 838}
                                                _ = v53
                                            case "":
                                                v26 -= v1
                                            }
                                        case 32767:
                                            fmt.Println(fi14, fi14, v3)
                                        default:
                                            if (!(!(42 > 655))) {
                                                v29 = (v29 % ((v29) | 1))
                                            }
                                        }
                                    }
                                    switch fi14 {
                                    case 832:
                                        if (!true) {
                                            v13 = v13
                                            var v54 uint64 = 1
                                            _ = v54
                                        } else {
                                            v55 := GenS7{F0: v29}
                                            v56 := v4[v13]
                                            v56 += v56
                                            v1 -= (v26 << ((4294967295) & 63))
                                            _ = v55
                                            _ = v56
                                        }
                                    case -497:
                                        v4[v13] = (v13 ^ v13)
                                    }
                                }
                            }
                            _ = v33
                            _ = v34
                        } else {
                            v57 := GenS3{F0: v29, F1: (("\n\t" + "hello") + (("hello" + "") + "hello")), F2: uint64(v1)}
                            fmt.Println(fi14, v2, fi14)
                            var v58 uint8 = (v1 >> ((1) & 63))
                            _ = v57
                            _ = v58
                        }
                        _ = v29
                        _ = v30
                    }
                    v59 := GenS8{F0: 189, F1: fi14, F2: (v13 - v13)}
                    _ = fi25
                    _ = v26
                    _ = v59
                }
                if false {
                    var v60 string = ("" + ("" + "a"))
                    v2 = v2
                    v5 = int64(v13)
                    _ = v60
                } else {
                    if (!true) {
                        if false {
                            switch (fi14 << ((4294967295) & 63)) {
                            case 42:
                                if true {
                                    switch ("foo bar" + "a") {
                                    case "hello":
                                        v61 := GenS7{F0: 2147483647}
                                        _ = v61
                                    case "":
                                        if ((-1 >= 839) || false) {
                                            for fi62 := 0; fi62 < 15; fi62++ {
                                                v1 += 65
                                                switch ("" + (("a" + "hello") + "")) {
                                                case "a":
                                                    var v63 uint32 = v13
                                                    _ = v63
                                                default:
                                                    v64 := GenS8{F0: 36, F1: (fi14 % (((fi14 ^ 740)) | 1)), F2: (v13 - 1)}
                                                    var v65 bool = (float64(-2.25) <= 1.5)
                                                    _ = v64
                                                    _ = v65
                                                }
                                                v3 = 32767
                                                _ = fi62
                                            }
                                            v66 := []int{0}
                                            _ = v66
                                        }
                                    }
                                }
                            case 1:
                                v3 -= (v3 | (v3 / ((1) | 1)))
                            default:
                                var v67 uint = 2
                                _ = v67
                            }
                            var v68 int32 = 659
                            v69 := GenS2{F0: (^(v3)), F1: (v2 * v5), F2: int64(v3)}
                            _ = v68
                            _ = v69
                        } else {
                            if true {
                                v4[v13] = (v13 / ((1) | 1))
                                v13 = 2147483647
                            } else {
                                v70 := GenS4{F0: 1, F1: 57, F2: (v13 | (^(v13)))}
                                _ = v70
                            }
                        }
                    } else {
                        switch (fi14 % ((-795) | 1)) {
                        case 2147483647:
                            v71 := v4[(v13 + (v13 + v13))]
                            _ = v71
                        default:
                            if (v2 <= int64(v1)) {
                                var v72 bool = true
                                _ = v72
                            } else {
                                if false {
                                    if (v1 > v1) {
                                        var v73 int16 = (v3 | (v3 / (((^(v3))) | 1)))
                                        for fi74 := 0; fi74 < 6; fi74++ {
                                            if false {
                                                if ((v2 - -969) <= -940) {
                                                    if false {
                                                        v75 := GenS1{F0: 42}
                                                        _ = v75
                                                    } else {
                                                        var v76 uint = 1
                                                        var v77 bool = true
                                                        v78 := GenS8{F0: 18446744073709551615, F1: (fi14 & (fi14 & fi74)), F2: v13}
                                                        _ = v76
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    v79 := []int{(fi74 - (fi14 ^ fi14))}
                                                    var v80 int8 = 0
                                                    _ = v79
                                                    _ = v80
                                                }
                                                v81 := GenS9{F0: 1, F1: (((1e10 * float32((1.5 * (3.14159 / 3.14159)))) * ((float32(float64(float32(((float64(1.5) * 1.5) * -2.25)))) - ((1e10 + ((((1.5 - (((float32((3.14159 / 3.14159)) + 3.14159) / 1.0) / 1e10)) - ((0.1 + 0.1) - (3.14159 - ((3.14159 * 0.0) / (3.14159 - 1.5))))) / float32(0.1)) / (3.14159 - (1.5 - 1.5)))) * 0.1)) + 1e10)) == 3.14159)}
                                                _ = v81
                                            }
                                            var v82 uint8 = (v1 & v1)
                                            v82 += (v82 % ((154) | 1))
                                            _ = fi74
                                            _ = v82
                                        }
                                        for fi83 := 0; fi83 < 12; fi83++ {
                                            var v84 uint32 = v13
                                            switch (fi83 - fi83) {
                                            case 1:
                                                if true {
                                                    switch fi14 {
                                                    case 128:
                                                        v85 := GenS10{F0: v5, F1: 57, F2: (^(v1))}
                                                        _ = v85
                                                    case 32767:
                                                        var v86 uint32 = v84
                                                        _ = v86
                                                    }
                                                }
                                            case -9223372036854775808:
                                                v87 := GenS11{F0: true}
                                                _ = v87
                                            }
                                            if (42 == uint64(v1)) {
                                                v73 -= 32767
                                                v88 := GenS12{F0: (!false), F1: false}
                                                if (!true) {
                                                    if ((true || (0.1 != (((1e10 * 1.5) * 0.1) - 1e10))) && false) {
                                                        fmt.Println(fi14)
                                                    } else {
                                                        var v89 int32 = 42
                                                        _ = v89
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v90 := GenS3{F0: -2147483648, F1: (("\n\t" + "a") + ("foo bar" + "hello")), F2: 2}
                                                        _ = v90
                                                    case "\n\t":
                                                        v91 := []int{4294967295, (-(fi83)), (^(fi14)), (fi83 | 32768)}
                                                        _ = v91
                                                    case "hello":
                                                        v5 = 42
                                                    default:
                                                        if false {
                                                            v2 = 42
                                                        } else {
                                                            v92 := []int{fi83, fi83, 1}
                                                            switch ("\n\t" + "\n\t") {
                                                            case "hello":
                                                                if true {
                                                                    if (!true) {
                                                                        v93 := GenS11{F0: (v2 != (v5 & v5))}
                                                                        v94 := GenS13{F0: 0}
                                                                        v95 := GenS14{F0: 431}
                                                                        _ = v93
                                                                        _ = v94
                                                                        _ = v95
                                                                    } else {
                                                                        var v96 uint8 = (v1 | v1)
                                                                        if false {
                                                                            v3 = (v3 / (((v73 ^ 0)) | 1))
                                                                        }
                                                                        _ = v96
                                                                    }
                                                                    v4[v13] = (v84 & v84)
                                                                    if len(v92) >= 2 {
                                                                        v92 = v92[1:len(v92)]
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                v84 -= 65535
                                                            case "a":
                                                                var v97 int = 9223372036854775807
                                                                _ = v97
                                                            }
                                                            _ = v92
                                                        }
                                                    }
                                                    var v98 int8 = 23
                                                    _ = v98
                                                } else {
                                                    fmt.Println(fi83)
                                                    v73 = v73
                                                }
                                                _ = v88
                                            }
                                            var v99 float64 = (0.1 / 1e10)
                                            v5 = (v5 + (v2 << ((18446744073709551615) & 63)))
                                            _ = fi83
                                            _ = v84
                                            _ = v99
                                        }
                                        _ = v73
                                    } else {
                                        var v100 uint64 = 2
                                        fmt.Println(v13, v1, fi14)
                                        v13 += 660
                                        _ = v100
                                    }
                                } else {
                                    for fi101 := 0; fi101 < 18; fi101++ {
                                        if (false || (!true)) {
                                            var v102 int = (-(fi14))
                                            var v103 int = v102
                                            var v104 int32 = 1
                                            _ = v102
                                            _ = v103
                                            _ = v104
                                        }
                                        v5 += (-(v2))
                                        _ = fi101
                                    }
                                    switch 247 {
                                    case -250:
                                        v105 := []int{326, fi14, (fi14 - 2147483647)}
                                        _ = v105
                                    default:
                                        var v106 int8 = -44
                                        _ = v106
                                    }
                                    v107 := GenS15{F0: 255, F1: uint64(v13)}
                                    _ = v107
                                }
                                var v108 bool = (true && true)
                                var v109 int8 = 1
                                _ = v108
                                _ = v109
                            }
                        }
                        v110 := GenS15{F0: 4294967295, F1: 2}
                        _ = v110
                    }
                    v111 := GenS7{F0: 0}
                    var v112 int = (fi14 + (fi14 - (fi14 * -115)))
                    _ = v111
                    _ = v112
                }
            }
        } else {
            v113 := GenS16{F0: true}
            _ = v113
        }
        v5 -= (v5 & (v2 & v2))
        var v114 uint16 = 625
        _ = fi14
        _ = v114
    }
    var v115 uint8 = (^(v1))
    for fi116 := 0; fi116 < 20; fi116++ {
        v117 := GenS7{F0: 0}
        for fi118 := 0; fi118 < 11; fi118++ {
            var v119 uint8 = (v115 | (v115 << ((764) & 63)))
            if true {
                v120 := []int{(-(fi116)), (fi116 % ((fi118) | 1)), fi118}
                _ = len(v4)
                _ = v120
            } else {
                v121 := GenS4{F0: 65535, F1: -1, F2: (v13 & 65535)}
                fmt.Println(v1)
                switch fi116 {
                case -885:
                    var v122 int64 = (^(v2))
                    _ = v122
                case 128:
                    switch "a" {
                    case "a":
                        switch ((("foo bar" + "") + "") + "") {
                        case "":
                            var v123 uint8 = v1
                            _ = v123
                        case "\n\t":
                            v1 -= 255
                        }
                    }
                default:
                    fmt.Println(v3)
                }
                _ = v121
            }
            _ = fi118
            _ = v119
        }
        _ = fi116
        _ = v117
    }
    var v124 uint32 = v13
    for fi125 := 0; fi125 < 13; fi125++ {
        fmt.Println(v1, v2, v1)
        v3 = v3
        v3 = v3
        _ = fi125
    }
    if ((v124 << ((18446744073709551615) & 63)) == v124) {
        var v126 int = 0
        _ = v126
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        if (!(true && (42 > -114))) {
            switch 347 {
            case 32768:
                v13 -= v124
            default:
                switch 1 {
                case 727:
                    v127 := []int{4294967295, 2147483647, int(v13), -343}
                    _ = v127
                default:
                    v128 := GenS3{F0: 2147483647, F1: "a", F2: 42}
                    _ = v128
                }
            }
        } else {
            var v129 int32 = -224
            var v130 uint32 = 2
            _ = v129
            _ = v130
        }
    }
    for fi131 := 0; fi131 < 14; fi131++ {
        for fi132 := 0; fi132 < 13; fi132++ {
            fmt.Println(v3, v124, v1)
            _ = fi132
        }
        _ = fi131
    }
    var v133 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v124:", v124)
    fmt.Println("v13:", v13)
    fmt.Println("v133:", v133)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v115
    _ = v124
    _ = v13
    _ = v133
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
