package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 uint8; F1 int; F2 bool }
type GenS7 struct { F0 int64; F1 int8; F2 uint32 }
type GenS8 struct { F0 int8; F1 int32; F2 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64; F1 uint; F2 uint }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint; F1 uint }
type GenS13 struct { F0 string; F1 bool; F2 int32 }
type GenS14 struct { F0 int16; F1 uint8; F2 string }
type GenS15 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint32 = 710
    v2 := make(map[string]int)
    var v3 int8 = 127
    if (v1 < 4294967295) {
        v4 := v2[("\n\t" + "foo bar")]
        _ = v4
    } else {
        v1 += (v1 + (v1 | 65535))
        var v5 uint64 = 78
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int64 = 0
        v8 := GenS1{F0: (^(fi6)), F1: v3, F2: v3}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 120, F1: (-(v3))}
    if false {
        v1 = v1
        fmt.Println(v3, v1)
        v2["a"] = 32767
    } else {
        v1 = v1
    }
    switch -2147483648 {
    case -9223372036854775808:
        if true {
            v10 := GenS1{F0: 9223372036854775807, F1: v3, F2: v3}
            _ = v10
        } else {
            var v11 uint64 = 689
            for fi12 := 0; fi12 < 10; fi12++ {
                switch 804 {
                case 4294967295:
                    if false {
                        v2[(("hello" + "foo bar") + "")] = -18
                        v3 = v3
                        if ((1 <= 2) && ((v3 << ((42) & 63)) == (v3 * 47))) {
                            if true {
                                if (v11 <= (^(v11))) {
                                    var v13 int = 2147483647
                                    var v14 uint64 = v11
                                    var v15 uint64 = v14
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                } else {
                                    var v16 int = fi12
                                    _ = v16
                                }
                            }
                            for fi17 := 0; fi17 < 5; fi17++ {
                                v18 := GenS3{F0: v11, F1: 1}
                                switch "" {
                                case "hello":
                                    switch -8 {
                                    case -636:
                                        v11 += v11
                                    case 4294967295:
                                        v1 -= v1
                                    default:
                                        switch -755 {
                                        case -401:
                                            var v19 int64 = 42
                                            _ = v19
                                        default:
                                            switch fi17 {
                                            case -2147483648:
                                                var v20 int8 = v3
                                                var v21 float32 = 0.1
                                                if ((!false) && ((v20 - (v3 / ((0) | 1))) > (v3 % ((v20) | 1)))) {
                                                    v21 = float32(1e10)
                                                } else {
                                                    switch fi17 {
                                                    case -9223372036854775808:
                                                        var v22 bool = ((-(v20)) == v3)
                                                        _ = v22
                                                    case 32767:
                                                        v23 := GenS4{F0: "", F1: -1}
                                                        v1 -= (^(v1))
                                                        if true {
                                                            switch fi17 {
                                                            case 21:
                                                                var v24 int16 = 42
                                                                _ = v24
                                                            }
                                                            var v25 int32 = 819
                                                            _ = v25
                                                        }
                                                        _ = v23
                                                    }
                                                }
                                                var v26 float64 = (-2.25 * ((((float64(float32(float64(0.0))) / 1.0) / 1.5) + 1e10) / 1.0))
                                                var v27 int32 = 42
                                                _ = v20
                                                _ = v21
                                                _ = v26
                                                _ = v27
                                            case 673:
                                                var v28 uint8 = 36
                                                _ = v28
                                            default:
                                                var v29 float32 = 1e10
                                                _ = v29
                                            }
                                        }
                                    }
                                default:
                                    v3 -= v3
                                }
                                v1 = uint32(v11)
                                _ = fi17
                                _ = v18
                            }
                            if ((-2.25 + ((3.14159 + 0.0) + 0.0)) <= (0.1 / float32(1.5))) {
                                v2["foo bar"] = fi12
                                for fi30 := 0; fi30 < 17; fi30++ {
                                    v11 -= v11
                                    if true {
                                        delete(v2, "hello")
                                        v31 := GenS5{F0: (1 <= 1000), F1: "\n\t"}
                                        _ = v31
                                    } else {
                                        v32 := GenS6{F0: 0, F1: fi30, F2: (!false)}
                                        v3 += (v3 / (((v3 + (v3 % ((-102) | 1)))) | 1))
                                        v33 := GenS7{F0: 1, F1: (v3 << ((uint(v1)) & 63)), F2: v1}
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = fi30
                                }
                            } else {
                                var v34 float64 = ((float64(0.1) * 3.14159) + 0.0)
                                for fi35 := 0; fi35 < 1; fi35++ {
                                    var v36 uint32 = v1
                                    _ = fi35
                                    _ = v36
                                }
                                switch "" {
                                case "a":
                                    switch ("" + "foo bar") {
                                    case "foo bar":
                                        if (false && (!true)) {
                                            if (true || true) {
                                                switch ("a" + "a") {
                                                case "":
                                                    v1 = uint32(v11)
                                                case "a":
                                                    v11 = 0
                                                }
                                                if false {
                                                    v11 = v11
                                                    switch ("foo bar" + ((("\n\t" + "a") + "") + "hello")) {
                                                    case "hello":
                                                        v37 := GenS7{F0: 515, F1: v3, F2: (v1 | (v1 - v1))}
                                                        _ = v37
                                                    case "foo bar":
                                                        var v38 float32 = (1e10 - (0.0 * -2.25))
                                                        _ = v38
                                                    default:
                                                        var v39 uint32 = v1
                                                        _ = v39
                                                    }
                                                    v40 := []int{(fi12 * fi12), fi12}
                                                    _ = v40
                                                } else {
                                                    if (42 >= 1) {
                                                        if ((((v34 + v34) / (3.14159 / v34)) - float64(1.5)) >= 3.14159) {
                                                            v41 := []int{287, fi12}
                                                            if len(v41) >= 2 {
                                                                v41 = v41[1:len(v41)]
                                                            }
                                                            _ = v41
                                                        }
                                                    } else {
                                                        fmt.Println(v11, fi12)
                                                        v42 := GenS7{F0: -1, F1: -1, F2: v1}
                                                        _ = v42
                                                    }
                                                    v11 = (v11 << ((2) & 63))
                                                }
                                            } else {
                                                v43 := GenS8{F0: (v3 >> ((4294967295) & 63)), F1: 832, F2: "\n\t"}
                                                v2["\n\t"] = -343
                                                _ = v43
                                            }
                                            for fi44 := 0; fi44 < 15; fi44++ {
                                                var v45 uint8 = 0
                                                if false {
                                                    v11 = v11
                                                }
                                                v2[("a" + ("foo bar" + ("" + "hello")))] = (fi12 & 1)
                                                v46 := []int{974}
                                                fmt.Println(fi12, fi12)
                                                var v47 string = ("a" + "hello")
                                                _ = fi44
                                                _ = v45
                                                _ = v46
                                                _ = v47
                                            }
                                        }
                                    case "a":
                                        v1 += v1
                                    default:
                                        for fi48 := 0; fi48 < 16; fi48++ {
                                            var v49 string = "a"
                                            if false {
                                                v50 := GenS9{F0: (-(fi48))}
                                                switch v49 {
                                                case "foo bar":
                                                    if true {
                                                        fmt.Println(v11, v1, v11)
                                                    } else {
                                                        v51 := v2["hello"]
                                                        var v52 bool = (!false)
                                                        if (!(!v52)) {
                                                            if v52 {
                                                                switch (v51 - v51) {
                                                                case 300:
                                                                    var v53 int32 = 2147483647
                                                                    _ = v53
                                                                case 32768:
                                                                    v54 := []int{}
                                                                    v2[("a" + ((("a" + ("hello" + v49)) + v49) + v49))] = fi48
                                                                    _ = v54
                                                                }
                                                                var v55 int64 = 42
                                                                _ = v55
                                                            } else {
                                                                if ((-2.25 / 1.5) > 3.14159) {
                                                                    if (0.0 != (v34 - ((float64(1.5) - v34) + float64(0.1)))) {
                                                                        fmt.Println(v49, fi48)
                                                                    } else {
                                                                        v56 := GenS10{F0: (^(v11)), F1: 255, F2: 18446744073709551615}
                                                                        fmt.Println(v51)
                                                                        _ = v56
                                                                    }
                                                                }
                                                            }
                                                            if ((!v52) || v52) {
                                                                v2[v49] = (v51 << ((245) & 63))
                                                            } else {
                                                                var v57 int64 = 1
                                                                _ = v57
                                                            }
                                                            if ("" > v49) {
                                                                v52 = (float64(3.14159) > float64(float32(v34)))
                                                                var v58 uint8 = 0
                                                                v59 := GenS11{F0: -53}
                                                                _ = v58
                                                                _ = v59
                                                            } else {
                                                                if v52 {
                                                                    switch ((v49 + v49) + v49) {
                                                                    case "\n\t":
                                                                        v60 := []int{(v51 << ((0) & 63))}
                                                                        _ = v60
                                                                    case "":
                                                                        v52 = v52
                                                                    }
                                                                    switch (v49 + v49) {
                                                                    case "a":
                                                                        var v61 uint16 = 1
                                                                        if true {
                                                                            v52 = v52
                                                                            v52 = (((42 == -734) || ((^(v61)) <= (v61 / ((65535) | 1)))) || v52)
                                                                            if v52 {
                                                                                var v62 uint64 = v11
                                                                                var v63 int64 = 42
                                                                                _ = v62
                                                                                _ = v63
                                                                            }
                                                                            v64 := GenS10{F0: uint64(v1), F1: 1, F2: 0}
                                                                            var v65 int8 = v3
                                                                            v66 := GenS10{F0: v11, F1: 0, F2: 1}
                                                                            v67 := GenS12{F0: 18446744073709551615, F1: uint(v61)}
                                                                            _ = v64
                                                                            _ = v65
                                                                            _ = v66
                                                                            _ = v67
                                                                        }
                                                                        _ = v61
                                                                    }
                                                                    fmt.Println(v34, v3)
                                                                    var v68 uint64 = 730
                                                                    _ = v68
                                                                }
                                                            }
                                                        }
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                case "\n\t":
                                                    _ = len(v2)
                                                    var v69 int64 = -9223372036854775808
                                                    if false {
                                                        v70 := []int{fi12, (^(fi48)), (fi48 >> ((18446744073709551615) & 63)), fi48}
                                                        var v71 float64 = 3.14159
                                                        v72 := GenS5{F0: false, F1: v49}
                                                        _ = v70
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    _ = v69
                                                default:
                                                    fmt.Println(v3, v3, fi48)
                                                    if true {
                                                        v49 = "\n\t"
                                                        v73 := GenS12{F0: 42, F1: 255}
                                                        _ = v73
                                                    }
                                                    delete(v2, "hello")
                                                }
                                                if true {
                                                    var v74 int32 = -2147483648
                                                    _ = v74
                                                } else {
                                                    var v75 uint32 = v1
                                                    var v76 float64 = v34
                                                    fmt.Println(v75, v49, fi48)
                                                    _ = v75
                                                    _ = v76
                                                }
                                                _ = v50
                                            } else {
                                                _ = len(v2)
                                                v77 := v2[v49]
                                                _ = v77
                                            }
                                            v1 = (v1 * v1)
                                            fmt.Println(fi12, v34)
                                            _ = fi48
                                            _ = v49
                                        }
                                    }
                                case "":
                                    switch 908 {
                                    case 0:
                                        delete(v2, "")
                                    default:
                                        var v78 uint16 = 607
                                        _ = v78
                                    }
                                case "foo bar":
                                    for fi79 := 0; fi79 < 17; fi79++ {
                                        var v80 uint = 2
                                        v34 += v34
                                        _ = fi79
                                        _ = v80
                                    }
                                default:
                                    if true {
                                        for fi81 := 0; fi81 < 17; fi81++ {
                                            if false {
                                                switch "a" {
                                                case "foo bar":
                                                    fmt.Println(fi81, v11, v11)
                                                    switch ("" + "a") {
                                                    case "\n\t":
                                                        var v82 bool = false
                                                        _ = v82
                                                    }
                                                    v83 := []int{(-(fi81)), (fi81 >> ((4294967295) & 63)), (fi12 + (fi12 ^ (fi12 - (fi81 | fi12)))), -239}
                                                    _ = v83
                                                case "\n\t":
                                                    if true {
                                                        var v84 string = ""
                                                        var v85 uint32 = (^(v1))
                                                        switch (-(fi12)) {
                                                        case 541:
                                                            if (!true) {
                                                                v86 := []int{fi12, fi81}
                                                                _ = v86
                                                            } else {
                                                                var v87 int32 = 0
                                                                fmt.Println(fi12, v11, v87)
                                                                _ = v87
                                                            }
                                                        case -1:
                                                            fmt.Println(v85, v85, v1)
                                                        case -930:
                                                            v88 := GenS7{F0: 1, F1: (v3 % ((v3) | 1)), F2: v1}
                                                            _ = v88
                                                        default:
                                                            switch fi12 {
                                                            case 920:
                                                                v84 = ((((v84 + "\n\t") + (v84 + "\n\t")) + v84) + v84)
                                                                v1 = v1
                                                            case 261:
                                                                fmt.Println(v1)
                                                            }
                                                        }
                                                        _ = v84
                                                        _ = v85
                                                    } else {
                                                        v34 += 1e10
                                                        fmt.Println(fi81, v11)
                                                    }
                                                case "":
                                                    if (0 == 39) {
                                                        fmt.Println(fi81, fi81, fi12)
                                                        fmt.Println(v1, v11)
                                                    }
                                                default:
                                                    v89 := []int{fi81}
                                                    _ = v89
                                                }
                                                switch "" {
                                                case "foo bar":
                                                    v34 -= float64(float32((((1e10 / v34) / v34) / ((v34 + 1.5) - float64((-2.25 + 0.1))))))
                                                default:
                                                    v1 = v1
                                                }
                                            }
                                            fmt.Println(v3)
                                            _ = fi81
                                        }
                                        v11 += v11
                                        var v90 bool = (!((fi12 & (fi12 ^ (fi12 ^ fi12))) > 32768))
                                        _ = v90
                                    } else {
                                        if (v34 == v34) {
                                            v11 += (v11 & (v11 - (v11 - (v11 + (^(v11))))))
                                            var v91 uint = 255
                                            _ = v91
                                        } else {
                                            var v92 bool = (!(!(((1.5 - v34) + ((1.5 * (float64(3.14159) / (v34 + float64(3.14159)))) + -2.25)) >= v34)))
                                            _ = v92
                                        }
                                        if false {
                                            v93 := GenS9{F0: (^(fi12))}
                                            for fi94 := 0; fi94 < 20; fi94++ {
                                                var v95 string = "hello"
                                                var v96 int = 32768
                                                _ = fi94
                                                _ = v95
                                                _ = v96
                                            }
                                            var v97 int32 = -116
                                            _ = v93
                                            _ = v97
                                        } else {
                                            v1 = (v1 << ((42) & 63))
                                        }
                                    }
                                }
                                _ = v34
                            }
                        } else {
                            var v98 uint = 313
                            if (!(fi12 > 305)) {
                                v99 := []int{(fi12 - (^(fi12))), 2147483647, (-(fi12)), fi12}
                                _ = v99
                            } else {
                                switch (fi12 >> ((42) & 63)) {
                                case -1:
                                    var v100 int32 = -559
                                    _ = v100
                                case 128:
                                    if true {
                                        if false {
                                            v101 := GenS13{F0: "", F1: false, F2: -1}
                                            _ = v101
                                        } else {
                                            if (true && false) {
                                                var v102 int32 = 0
                                                var v103 uint64 = v11
                                                _ = v102
                                                _ = v103
                                            } else {
                                                if true {
                                                    for fi104 := 0; fi104 < 3; fi104++ {
                                                        var v105 uint16 = 255
                                                        v106 := GenS14{F0: 304, F1: 2, F2: "hello"}
                                                        _ = len(v2)
                                                        _ = fi104
                                                        _ = v105
                                                        _ = v106
                                                    }
                                                }
                                            }
                                        }
                                        if false {
                                            var v107 uint16 = 2
                                            var v108 int = (fi12 * fi12)
                                            _ = v107
                                            _ = v108
                                        }
                                    } else {
                                        v3 += (v3 - 127)
                                        for fi109 := 0; fi109 < 5; fi109++ {
                                            switch (fi109 * (-(fi12))) {
                                            case 32768:
                                                v98 += v98
                                            case 925:
                                                if true {
                                                    fmt.Println(v11)
                                                    var v110 int8 = (v3 | (v3 % ((v3) | 1)))
                                                    var v111 int = -81
                                                    _ = v110
                                                    _ = v111
                                                }
                                            case 42:
                                                var v112 uint64 = 499
                                                _ = v112
                                            default:
                                                var v113 uint16 = 979
                                                _ = v113
                                            }
                                            _ = fi109
                                        }
                                    }
                                case 127:
                                    v1 = 2147483647
                                }
                                delete(v2, ("foo bar" + ""))
                            }
                            _ = v98
                        }
                    } else {
                        fmt.Println(v11, v11)
                    }
                }
                fmt.Println(v11, v11)
                switch "foo bar" {
                case "hello":
                    var v114 int32 = -1
                    _ = v114
                case "\n\t":
                    fmt.Println(v3, v11)
                default:
                    var v115 int8 = v3
                    _ = v115
                }
                _ = fi12
            }
            for fi116 := 0; fi116 < 20; fi116++ {
                v2[""] = fi116
                var v117 int = (fi116 + 32768)
                v11 = uint64(v1)
                _ = fi116
                _ = v117
            }
            _ = v11
        }
    case 32768:
        v1 = (^(v1))
    default:
        fmt.Println(v1, v3)
    }
    v3 = (v3 | (v3 - -128))
    for fi118 := 0; fi118 < 8; fi118++ {
        if true {
            if false {
                for fi119 := 0; fi119 < 3; fi119++ {
                    switch (fi118 % ((4294967295) | 1)) {
                    case 487:
                        var v120 int32 = 42
                        _ = v120
                    }
                    fmt.Println(fi119, fi119)
                    _ = fi119
                }
            } else {
                if true {
                    for fi121 := 0; fi121 < 2; fi121++ {
                        v122 := GenS12{F0: 1, F1: 2}
                        fmt.Println(fi118)
                        _ = fi121
                        _ = v122
                    }
                    switch (fi118 | fi118) {
                    case -1:
                        v123 := []int{1, 1, 128}
                        _ = v123
                    case 42:
                        if (1.5 > -2.25) {
                            v124 := GenS3{F0: 764, F1: 128}
                            var v125 int16 = 1000
                            _ = v124
                            _ = v125
                        }
                    default:
                        var v126 uint32 = (^(v1))
                        _ = v126
                    }
                } else {
                    var v127 int32 = -1
                    _ = v127
                }
                var v128 uint = 18446744073709551615
                _ = v128
            }
            v129 := v2["a"]
            _ = v129
        } else {
            v130 := GenS15{F0: (false || (v3 == v3)), F1: 209}
            if true {
                var v131 uint64 = 42
                _ = v131
            }
            fmt.Println(v1, v1, fi118)
            _ = v130
        }
        _ = fi118
    }
    if (true || (v1 > (v1 / (((v1 >> ((42) & 63))) | 1)))) {
        for fi132 := 0; fi132 < 10; fi132++ {
            v133 := GenS13{F0: "", F1: (false || false), F2: 2147483647}
            _ = fi132
            _ = v133
        }
    } else {
        v3 = -128
        fmt.Println(v3, v1)
    }
    v134 := []int{-9223372036854775808, -320}
    v1 = 4294967295
    fmt.Println(v3, v1)
    if false {
        v1 = (v1 - v1)
        switch 2147483647 {
        case -196:
            var v135 int8 = (v3 / (((v3 * v3)) | 1))
            _ = v135
        case 42:
            for fi136 := 0; fi136 < 6; fi136++ {
                var v137 uint32 = 4294967295
                v3 -= -128
                v138 := GenS14{F0: -99, F1: 255, F2: ("a" + (("a" + ("" + "hello")) + "a"))}
                _ = fi136
                _ = v137
                _ = v138
            }
        case 1:
            v1 -= v1
        }
        fmt.Println(v1, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v134:", len(v134), v134)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v134
    _ = v2
    _ = v3
    _ = v9
}
