package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 uint; F1 uint8 }
type GenS8 struct { F0 uint8; F1 int16; F2 int16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint; F1 int8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint8; F1 uint32 }
type GenS16 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 bool; F1 int8; F2 int8 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 uint }
type GenS21 struct { F0 int32; F1 uint; F2 int64 }
type GenS22 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 uint32 = 334
    var v4 float64 = 0.1
    var v5 uint16 = 255
    if (42 < 9223372036854775807) {
        var v6 float64 = (3.14159 * v4)
        if (-754 != 0) {
            var v7 string = (("a" + ((((((("foo bar" + "foo bar") + "") + "\n\t") + ("hello" + "")) + (((("foo bar" + "\n\t") + "\n\t") + (("a" + ("hello" + "hello")) + ((("a" + "foo bar") + "hello") + ("hello" + "")))) + "")) + "a") + "")) + "\n\t")
            _ = v7
        } else {
            var v8 uint32 = v3
            var v9 int32 = 0
            _ = v8
            _ = v9
        }
        var v10 uint = 554
        _ = v10
        _ = v6
    } else {
        v1["hello"] = (50 <= 104)
    }
    if false {
        if (-1 == -2147483648) {
            v5 = (v5 ^ (v5 | (^(v5))))
            v11 := GenS1{F0: 4294967295}
            if (false && (true && true)) {
                v5 = 65535
            }
            _ = v11
        }
        if (!(false || (!false))) {
            if (v5 == (v5 * (^(v5)))) {
                v12 := GenS2{F0: 203, F1: 4294967295}
                var v13 float64 = v4
                _ = v12
                _ = v13
            } else {
                v14 := GenS3{F0: (v3 * 2147483647)}
                v15 := GenS1{F0: 42}
                _ = v14
                _ = v15
            }
        } else {
            var v16 uint16 = (v5 - uint16(v3))
            var v17 int32 = 42
            v4 = 0.0
            _ = v16
            _ = v17
        }
    } else {
        v4 -= ((v4 / v4) + v4)
    }
    v4 = v4
    var v18 float32 = float32(float64(-2.25))
    v19 := GenS3{F0: (v3 / (((v3 * (v3 & 2147483647))) | 1))}
    switch "a" {
    case "\n\t":
        if false {
            var v20 float64 = 3.14159
            if false {
                var v21 uint = 18446744073709551615
                v4 -= 0.0
                _ = v21
            }
            _ = v20
        } else {
            var v22 uint32 = v3
            fmt.Println(v3)
            if (42 >= 112) {
                var v23 int = -1
                _ = v23
            }
            _ = v22
        }
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS1{F0: 18446744073709551615}
            v27 := []int{fi25, fi24}
            switch (fi25 ^ (fi25 ^ 265)) {
            case 135:
                if true {
                    var v28 uint8 = 2
                    var v29 uint8 = 255
                    _ = v28
                    _ = v29
                }
            case 1:
                v3 = v3
                if (v5 == (v5 ^ 65535)) {
                    if (!(19 < 100)) {
                        if false {
                            var v30 int = -2147483648
                            var v31 bool = (!((-643 >= 67) || false))
                            var v32 float64 = ((((1e10 * v4) + 0.0) / v4) + v4)
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    }
                    v33 := GenS3{F0: (v3 | (v3 + (v3 - 2)))}
                    v18 += v18
                    _ = v33
                }
            case -885:
                v2[("foo bar" + (("hello" + "hello") + ""))] = true
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        v1[""] = (true && (0 < 370))
        _ = fi24
    }
    fmt.Println(v4, v5, v4)
    v5 = (^(v5))
    if (v18 > ((0.0 + v18) + 1e10)) {
        v34 := GenS4{F0: 1, F1: -91}
        fmt.Println(v3)
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 uint16 = (v5 & v5)
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        var v38 uint16 = v5
        v5 -= v38
        if false {
            v38 += 42
            if false {
                v18 = v18
                var v39 float32 = (v18 / 1.5)
                for fi40 := 0; fi40 < 11; fi40++ {
                    switch "hello" {
                    case "foo bar":
                        var v41 int16 = 70
                        var v42 int64 = -684
                        v3 -= (^(v3))
                        v43 := []int{(fi37 * 127), (fi37 * -12), -9223372036854775808}
                        if ((v41 | 32767) == v41) {
                            var v44 uint = 18446744073709551615
                            v44 = (v44 + v44)
                            _ = v44
                        } else {
                            switch (-(fi37)) {
                            case -403:
                                v5 += v38
                            case -9223372036854775808:
                                v45 := GenS5{F0: 4294967295}
                                _ = v45
                            }
                            v39 = (v39 - 3.14159)
                        }
                        switch 1 {
                        case 32768:
                            var v46 int64 = v42
                            _ = v46
                        case 42:
                            fmt.Println(v5)
                            v42 += 9223372036854775807
                            switch ("foo bar" + ("" + "\n\t")) {
                            case "\n\t":
                                var v47 uint16 = v38
                                if len(v43) > 0 {
                                    v43[0] = (-(fi40))
                                }
                                if false {
                                    var v48 string = ""
                                    var v49 int32 = -1
                                    _ = v48
                                    _ = v49
                                } else {
                                    var v50 int8 = 0
                                    _ = len(v43)
                                    v41 += (^(v41))
                                    v51 := GenS5{F0: 4294967295}
                                    _ = v50
                                    _ = v51
                                }
                                _ = v47
                            }
                            if (32768 > (fi40 % (((^(fi37))) | 1))) {
                                if false {
                                    _ = len(v43)
                                    v18 = ((v18 - 1.5) - float32(-2.25))
                                    switch "hello" {
                                    case "":
                                        fmt.Println(v42)
                                    case "foo bar":
                                        if true {
                                            if ((("\n\t" + "") + ("foo bar" + ("" + "a"))) == ((("a" + ("a" + "")) + (("\n\t" + "a") + "\n\t")) + "foo bar")) {
                                                fmt.Println(v18, v41)
                                            } else {
                                                if (1 <= -1) {
                                                    v42 += 154
                                                    var v52 int = (fi40 + fi37)
                                                    v43 = append(v43, v52)
                                                    var v53 uint64 = 146
                                                    _ = v52
                                                    _ = v53
                                                }
                                                if false {
                                                    _ = len(v43)
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        fmt.Println(v42)
                                                    case "hello":
                                                        if true {
                                                            if ("hello" == ("foo bar" + "\n\t")) {
                                                                v41 = v41
                                                            }
                                                            v54 := v2[((("a" + "foo bar") + "") + "a")]
                                                            _ = v54
                                                        }
                                                        fmt.Println(v38, v5, v3)
                                                        if (!false) {
                                                            var v55 float64 = 1.5
                                                            if true {
                                                                v56 := GenS6{F0: 1, F1: (fi37 & (^(fi37)))}
                                                                _ = v56
                                                            }
                                                            var v57 int = 2147483647
                                                            _ = v55
                                                            _ = v57
                                                        }
                                                    case "a":
                                                        v3 = v3
                                                    }
                                                }
                                                v58 := GenS7{F0: 373, F1: 2}
                                                switch (-(fi37)) {
                                                case 4294967295:
                                                    switch "a" {
                                                    case "foo bar":
                                                        switch (fi37 + (-(fi40))) {
                                                        case 4294967295:
                                                            v43 = append(v43, fi37)
                                                        case 1:
                                                            switch "a" {
                                                            case "":
                                                                v18 = 0.0
                                                            case "a":
                                                                v5 = (v5 | v5)
                                                            }
                                                        case 2147483647:
                                                            v1["hello"] = ("" == "\n\t")
                                                        }
                                                    case "hello":
                                                        var v59 uint64 = 18446744073709551615
                                                        _ = v59
                                                    case "a":
                                                        v18 -= float32(v4)
                                                    default:
                                                        v39 -= 1e10
                                                    }
                                                case 128:
                                                    var v60 int = fi37
                                                    _ = v60
                                                default:
                                                    if ((-(v42)) > v42) {
                                                        var v61 string = "hello"
                                                        var v62 string = v61
                                                        v43 = append(v43, (fi37 + int(v41)))
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                }
                                                var v63 int32 = 421
                                                switch (^(fi40)) {
                                                case 1:
                                                    var v64 float32 = v18
                                                    _ = v64
                                                }
                                                _ = v58
                                                _ = v63
                                            }
                                            fmt.Println(v41, v5)
                                        } else {
                                            fmt.Println(v38)
                                            var v65 int = 32767
                                            v66 := GenS1{F0: 2}
                                            if (v39 <= v18) {
                                                switch (v65 / (((fi40 | v65)) | 1)) {
                                                case -367:
                                                    switch v65 {
                                                    case 0:
                                                        var v67 uint8 = 2
                                                        _ = v67
                                                    case 4294967295:
                                                        fmt.Println(v4, v65)
                                                    }
                                                case 630:
                                                    v4 = v4
                                                case 0:
                                                    v42 = -1
                                                }
                                                v68 := GenS8{F0: 127, F1: -1, F2: (v41 ^ v41)}
                                                if len(v43) >= 2 {
                                                    v43 = v43[1:len(v43)]
                                                }
                                                _ = v68
                                            }
                                            var v69 float32 = v39
                                            _ = v65
                                            _ = v66
                                            _ = v69
                                        }
                                    case "hello":
                                        switch -691 {
                                        case 2147483647:
                                            if (0.1 >= (v39 * v39)) {
                                                var v70 float64 = v4
                                                fmt.Println(v3)
                                                if (false && (true || false)) {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v41 = 712
                                                    case "":
                                                        var v71 int8 = -1
                                                        _ = v71
                                                    default:
                                                        v43 = append(v43, fi40)
                                                    }
                                                }
                                                fmt.Println(v18, v41)
                                                v18 = float32((v70 * v4))
                                                v5 += 777
                                                _ = v70
                                            } else {
                                                var v72 int = fi40
                                                var v73 bool = (130 >= 4294967295)
                                                v74 := GenS9{F0: (v41 % (((v41 & -1)) | 1))}
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                            }
                                        case 127:
                                            fmt.Println(fi40, v4, fi40)
                                        default:
                                            v41 = 995
                                        }
                                        v5 += 1
                                    default:
                                        if len(v43) > 0 {
                                            v43[0] = fi37
                                        }
                                        v42 = (v42 % ((v42) | 1))
                                        if (false && false) {
                                            v4 = 1.5
                                            var v75 int16 = v41
                                            if true {
                                                var v76 int8 = 23
                                                _ = v76
                                            }
                                            _ = v75
                                        } else {
                                            fmt.Println(v42, fi40)
                                        }
                                    }
                                }
                                var v77 bool = false
                                _ = v77
                            } else {
                                v38 = v38
                                var v78 float32 = v39
                                _ = v78
                            }
                            if (false || false) {
                                v2[("hello" + "hello")] = true
                                var v79 int16 = (-(v41))
                                fmt.Println(fi40, v42)
                                _ = v79
                            } else {
                                v80 := GenS10{F0: 553, F1: 42}
                                v41 = -702
                                _ = v80
                            }
                        default:
                            v81 := GenS7{F0: 42, F1: 127}
                            _ = v81
                        }
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    switch "foo bar" {
                    case "foo bar":
                        v38 = 255
                    case "hello":
                        v82 := GenS11{F0: fi40}
                        _ = v82
                    case "":
                        v83 := GenS9{F0: 0}
                        _ = v83
                    default:
                        v84 := []int{(fi37 - 377), fi37, -96}
                        _ = v84
                    }
                    var v85 float64 = (v4 * v4)
                    _ = fi40
                    _ = v85
                }
                _ = v39
            } else {
                v38 = (v38 / ((v38) | 1))
                var v86 uint64 = 18446744073709551615
                _ = v86
            }
            v87 := GenS12{F0: v3}
            _ = v87
        } else {
            var v88 float64 = float64((float32((v4 * 0.0)) + ((v18 / v18) - 0.0)))
            _ = v88
        }
        _ = fi37
        _ = v38
    }
    switch (("" + ("" + ((("a" + (("\n\t" + "\n\t") + "\n\t")) + ("foo bar" + "hello")) + "foo bar"))) + "foo bar") {
    case "a":
        switch 2147483647 {
        case 4294967295:
            if true {
                var v89 bool = true
                v18 = float32((((1e10 - float64(-2.25)) * (float64(0.0) + v4)) * v4))
                _ = v89
            } else {
                v4 -= (1e10 / float64(v18))
                var v90 bool = true
                _ = v90
            }
        case 127:
            var v91 int32 = -1
            _ = v91
        }
    case "foo bar":
        switch 128 {
        case 4294967295:
            for fi92 := 0; fi92 < 9; fi92++ {
                for fi93 := 0; fi93 < 10; fi93++ {
                    if false {
                        v94 := GenS13{F0: 127}
                        fmt.Println(v3, fi92)
                        _ = v94
                    } else {
                        v95 := GenS14{F0: 27}
                        v96 := []int{fi93}
                        if false {
                            v5 = (v5 ^ 1)
                            var v97 uint8 = 118
                            _ = v97
                        } else {
                            v4 = 1e10
                            var v98 bool = (!false)
                            _ = v98
                        }
                        _ = v95
                        _ = v96
                    }
                    _ = fi93
                }
                fmt.Println(v4, v3)
                _ = fi92
            }
        }
    }
    var v99 int8 = 0
    switch "" {
    case "foo bar":
        if true {
            var v100 float32 = 1e10
            _ = v100
        } else {
            if true {
                for fi101 := 0; fi101 < 1; fi101++ {
                    v18 += (((v18 - 3.14159) - v18) / v18)
                    _ = fi101
                }
            }
            if ((false || (v99 <= (v99 - v99))) && true) {
                if false {
                    fmt.Println(v3, v18, v99)
                    if (!(42 < 1)) {
                        if true {
                            var v102 uint8 = 128
                            _ = v102
                        } else {
                            for fi103 := 0; fi103 < 20; fi103++ {
                                v18 += (0.1 * ((((0.0 - (v18 * (v18 / 0.1))) + v18) / 1.5) * ((v18 + (-2.25 / 1.0)) * float32(3.14159))))
                                fmt.Println(v5)
                                _ = fi103
                            }
                            var v104 int64 = 9223372036854775807
                            _ = v104
                        }
                        if false {
                            fmt.Println(v4, v5)
                            v105 := []int{4294967295, 128}
                            for fi106 := 0; fi106 < 19; fi106++ {
                                var v107 int8 = (v99 - 100)
                                v108 := GenS2{F0: 0, F1: 65535}
                                v99 = (v99 * (v107 | (v107 % ((127) | 1))))
                                _ = fi106
                                _ = v107
                                _ = v108
                            }
                            _ = v105
                        } else {
                            for fi109 := 0; fi109 < 11; fi109++ {
                                var v110 int64 = 812
                                _ = fi109
                                _ = v110
                            }
                            if false {
                                if (0 > 1) {
                                    var v111 float32 = v18
                                    var v112 int64 = -425
                                    v99 += (v99 | v99)
                                    _ = v111
                                    _ = v112
                                } else {
                                    v4 = 1.5
                                }
                                switch ("hello" + "") {
                                case "foo bar":
                                    var v113 float64 = -2.25
                                    _ = v113
                                case "hello":
                                    v99 += v99
                                case "":
                                    v5 = 0
                                }
                                switch 1 {
                                case 32768:
                                    v114 := GenS7{F0: 42, F1: 42}
                                    _ = v114
                                }
                            } else {
                                switch 42 {
                                case -2147483648:
                                    v115 := GenS15{F0: 255, F1: (^(v3))}
                                    _ = v115
                                case 212:
                                    for fi116 := 0; fi116 < 20; fi116++ {
                                        v117 := []int{(fi116 | (fi116 + -443)), fi116, 32767, (-(fi116))}
                                        for fi118 := 0; fi118 < 19; fi118++ {
                                            v4 = ((v4 / float64((v18 / (float32((((v4 / float64(v18)) * (3.14159 / (v4 / v4))) / (float64(0.1) - (((0.1 - float64(((v18 + float32(1e10)) / v18))) * -2.25) - v4)))) / (v18 - 1e10))))) + (v4 + 1e10))
                                            if false {
                                                var v119 float32 = 0.0
                                                v120 := GenS10{F0: 985, F1: v99}
                                                v121 := GenS11{F0: (fi116 % ((4294967295) | 1))}
                                                v122 := GenS16{F0: (v99 ^ -128), F1: v5, F2: 0}
                                                _ = v119
                                                _ = v120
                                                _ = v121
                                                _ = v122
                                            }
                                            v123 := GenS17{F0: 664}
                                            if ((1.5 < v4) && true) {
                                                v5 += (v5 << ((255) & 63))
                                                fmt.Println(v99, fi116, v99)
                                                v4 -= float64(v18)
                                            } else {
                                                v117 = append(v117, fi118)
                                            }
                                            _ = fi118
                                            _ = v123
                                        }
                                        v18 -= (-2.25 * -2.25)
                                        _ = fi116
                                        _ = v117
                                    }
                                case 4294967295:
                                    v99 = (v99 >> ((18446744073709551615) & 63))
                                default:
                                    for fi124 := 0; fi124 < 2; fi124++ {
                                        v125 := GenS18{F0: true, F1: 1, F2: v99}
                                        v126 := GenS19{F0: false}
                                        v99 -= v99
                                        _ = fi124
                                        _ = v125
                                        _ = v126
                                    }
                                }
                            }
                        }
                    }
                }
            } else {
                v127 := GenS14{F0: -1}
                switch 128 {
                case -875:
                    v128 := []int{}
                    _ = v128
                case 716:
                    for fi129 := 0; fi129 < 18; fi129++ {
                        var v130 uint = 34
                        _ = fi129
                        _ = v130
                    }
                case 128:
                    switch -484 {
                    case 127:
                        if (1 <= -776) {
                            v131 := []int{9223372036854775807}
                            var v132 uint = 1
                            _ = v131
                            _ = v132
                        } else {
                            v18 += v18
                        }
                    case 0:
                        v3 += (v3 / ((4294967295) | 1))
                    default:
                        var v133 int16 = -1
                        _ = v133
                    }
                default:
                    v5 = v5
                }
                _ = v127
            }
            for fi134 := 0; fi134 < 19; fi134++ {
                v135 := GenS19{F0: (1e10 <= v18)}
                _ = fi134
                _ = v135
            }
        }
    case "hello":
        if (false || true) {
            for fi136 := 0; fi136 < 18; fi136++ {
                delete(v1, "a")
                _ = fi136
            }
            var v137 string = "foo bar"
            var v138 float32 = v18
            _ = v137
            _ = v138
        } else {
            var v139 int32 = -2147483648
            _ = v139
        }
    case "\n\t":
        delete(v2, ("hello" + "foo bar"))
    default:
        if true {
            v99 = v99
            v140 := GenS20{F0: 1}
            if (!false) {
                v3 = (v3 + (^(v3)))
                v3 = (v3 ^ 65535)
            }
            _ = v140
        } else {
            var v141 string = ("hello" + "")
            v142 := GenS21{F0: 0, F1: 1, F2: 1}
            _ = v141
            _ = v142
        }
    }
    v99 += (v99 << ((1) & 63))
    v143 := GenS22{F0: (^(v5)), F1: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v143
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v99
}
