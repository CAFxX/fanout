package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int16; F2 string }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int; F1 int8; F2 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint16; F1 string; F2 int64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int; F1 int8 }
type GenS11 struct { F0 string; F1 uint32; F2 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = -50
    v3 := make(map[int]int)
    var v4 int16 = 1000
    var v5 uint32 = 847
    var v6 uint = 18446744073709551615
    v7 := GenS1{F0: 457, F1: 0, F2: 1}
    v8 := []int{0, -867, -542}
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(v4)
        switch fi9 {
        case 822:
            v10 := GenS2{F0: 71, F1: v4, F2: (("hello" + ("" + ("a" + ("" + "\n\t")))) + ("\n\t" + ("" + "a")))}
            _ = v10
        case 0:
            if (false && false) {
                switch -806 {
                case -1:
                    if ((!((true && true) || false)) || true) {
                        var v11 float32 = 3.14159
                        _ = v11
                    } else {
                        switch int(v4) {
                        case -770:
                            var v12 int32 = -2147483648
                            _ = v12
                        case -1:
                            for fi13 := 0; fi13 < 8; fi13++ {
                                var v14 int8 = v2
                                _ = fi13
                                _ = v14
                            }
                        case 343:
                            for fi15 := 0; fi15 < 13; fi15++ {
                                if true {
                                    v16 := GenS3{F0: (v2 + -128), F1: ("" < "foo bar")}
                                    var v17 uint = 714
                                    if true {
                                        if (4294967295 <= 2) {
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(fi9, v5)
                                            case "foo bar":
                                                v3[(fi9 - (fi9 * fi9))] = fi15
                                                if len(v8) >= 2 {
                                                    v8 = v8[1:len(v8)]
                                                }
                                                v18 := GenS4{F0: (-(fi9)), F1: int8(v4), F2: 42}
                                                _ = v18
                                            case "hello":
                                                var v19 float64 = (float64(((-2.25 / float32((float64(0.1) / 0.1))) / 1.0)) - (3.14159 - (float64((((((1.5 - float32(1.5)) - (0.1 / 1.0)) * 0.0) * -2.25) - (((((0.0 + -2.25) + 0.1) / 1.5) / -2.25) / float32(3.14159)))) * (1e10 / float64(1.5)))))
                                                _ = v19
                                            }
                                            fmt.Println(v5, v5, v4)
                                            var v20 int8 = 42
                                            _ = v20
                                        } else {
                                            if false {
                                                var v21 int16 = v4
                                                if true {
                                                    fmt.Println(v5)
                                                    switch (fi15 - (fi15 - -2147483648)) {
                                                    case 1:
                                                        if (fi9 < (fi15 ^ fi9)) {
                                                            if (1 == v21) {
                                                                v21 = v21
                                                            }
                                                        } else {
                                                            if false {
                                                                var v22 int16 = (v4 + (^(v21)))
                                                                if (!(!(!false))) {
                                                                    if true {
                                                                        v6 = v6
                                                                        v22 -= v21
                                                                    } else {
                                                                        v23 := GenS2{F0: 128, F1: v22, F2: "foo bar"}
                                                                        _ = v23
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        v24 := GenS2{F0: 1, F1: (v4 & v21), F2: ""}
                                                                        v5 = v5
                                                                        _ = v24
                                                                    }
                                                                }
                                                                _ = v22
                                                            }
                                                        }
                                                    case 18:
                                                        var v25 int8 = (-(v2))
                                                        _ = v25
                                                    default:
                                                        v26 := GenS3{F0: v2, F1: false}
                                                        _ = v26
                                                    }
                                                } else {
                                                    if true {
                                                        v21 = v4
                                                        var v27 uint64 = 843
                                                        v2 = v2
                                                        _ = v27
                                                    } else {
                                                        if (("\n\t" >= "") || true) {
                                                            v6 = (v17 << ((1) & 63))
                                                            v28 := GenS5{F0: 32767}
                                                            _ = v28
                                                        }
                                                        v29 := GenS5{F0: v4}
                                                        _ = v29
                                                    }
                                                    var v30 int32 = -53
                                                    var v31 int = (^(fi9))
                                                    if ((!(!false)) && true) {
                                                        v32 := GenS2{F0: 255, F1: v4, F2: "a"}
                                                        v33 := GenS1{F0: v5, F1: v30, F2: 42}
                                                        _ = v32
                                                        _ = v33
                                                    } else {
                                                        v2 = (v2 + v2)
                                                    }
                                                    switch ("a" + "\n\t") {
                                                    case "foo bar":
                                                        var v34 uint16 = 32767
                                                        _ = v34
                                                    case "a":
                                                        _ = len(v3)
                                                    }
                                                    var v35 int8 = int8(v21)
                                                    v2 += v2
                                                    fmt.Println(v21)
                                                    v8 = append(v8, (fi9 | (v31 ^ 427)))
                                                    v6 += 255
                                                    var v36 int32 = (v30 + (v30 ^ int32(v35)))
                                                    _ = v30
                                                    _ = v31
                                                    _ = v35
                                                    _ = v36
                                                }
                                                if (!(v2 < -128)) {
                                                    if (!(true || true)) {
                                                        var v37 float64 = 1.5
                                                        _ = v37
                                                    }
                                                }
                                                _ = v21
                                            } else {
                                                switch fi15 {
                                                case 9223372036854775807:
                                                    switch "" {
                                                    case "a":
                                                        var v38 uint = (v6 & (v6 | (v6 << ((792) & 63))))
                                                        _ = v38
                                                    case "\n\t":
                                                        v39 := GenS6{F0: 2}
                                                        var v40 int64 = -336
                                                        if true {
                                                            if ((v40 > 42) && (true || true)) {
                                                                var v41 int64 = (v40 ^ v40)
                                                                var v42 float32 = 0.1
                                                                var v43 uint16 = 32767
                                                                _ = v41
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                var v44 uint32 = v5
                                                                var v45 string = "\n\t"
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                            v6 = (v17 & (v6 >> (((v17 ^ v6)) & 63)))
                                                            v2 += (^(v2))
                                                        } else {
                                                            if false {
                                                                v46 := GenS3{F0: 0, F1: false}
                                                                v4 = 32767
                                                                if (-32768 >= 32767) {
                                                                    if false {
                                                                        var v47 bool = (((fi15 ^ (fi15 << (((v6 % ((874) | 1))) & 63))) > (fi9 + (fi9 << ((v6) & 63)))) || (((0.0 <= float32(((0.1 - -2.25) + -2.25))) && (34 <= 255)) || (!(v6 != v6))))
                                                                        v17 -= (v6 % (((v17 & 18446744073709551615)) | 1))
                                                                        v48 := GenS5{F0: (v4 & v4)}
                                                                        var v49 string = (((("foo bar" + "a") + "hello") + "foo bar") + ("" + ("\n\t" + "a")))
                                                                        _ = v47
                                                                        _ = v48
                                                                        _ = v49
                                                                    } else {
                                                                        var v50 int32 = -99
                                                                        v51 := GenS4{F0: (fi9 >> ((42) & 63)), F1: -1, F2: (-(v2))}
                                                                        if true {
                                                                            fmt.Println(v2)
                                                                            fmt.Println(fi15, v40, v2)
                                                                        }
                                                                        _ = v50
                                                                        _ = v51
                                                                    }
                                                                    v2 += 100
                                                                }
                                                                _ = v46
                                                            }
                                                            if (397 >= (^(fi15))) {
                                                                v6 = (v6 - (v17 - 18446744073709551615))
                                                            }
                                                        }
                                                        _ = v39
                                                        _ = v40
                                                    case "":
                                                        v3[(^(fi9))] = fi15
                                                    }
                                                case 662:
                                                    if true {
                                                        v52 := GenS3{F0: v2, F1: (true || (!true))}
                                                        if len(v8) > 0 {
                                                            v8[0] = fi9
                                                        }
                                                        _ = v52
                                                    }
                                                default:
                                                    var v53 uint64 = 18446744073709551615
                                                    _ = v53
                                                }
                                                v5 -= (v5 - (v5 << ((v6) & 63)))
                                                var v54 int8 = v2
                                                _ = v54
                                            }
                                        }
                                    }
                                    _ = v16
                                    _ = v17
                                }
                                _ = fi15
                            }
                        }
                        v55 := GenS2{F0: 128, F1: (v4 & v4), F2: "hello"}
                        v6 -= (v6 * v6)
                        _ = v55
                    }
                case 128:
                    var v56 float64 = (0.0 / (((1.5 * -2.25) / float64((3.14159 / 1e10))) - 3.14159))
                    _ = v56
                default:
                    for fi57 := 0; fi57 < 11; fi57++ {
                        var v58 uint64 = 2
                        var v59 uint16 = 1
                        var v60 int64 = -9223372036854775808
                        _ = fi57
                        _ = v58
                        _ = v59
                        _ = v60
                    }
                }
                v8 = append(v8, fi9)
            } else {
                fmt.Println(fi9, v5, v2)
                var v61 uint8 = 127
                switch fi9 {
                case -2147483648:
                    for fi62 := 0; fi62 < 6; fi62++ {
                        if false {
                            v4 = -353
                            v6 -= (v6 * uint(v5))
                            var v63 float64 = 0.1
                            _ = v63
                        } else {
                            var v64 bool = (!(!(!false)))
                            v65 := GenS3{F0: v2, F1: v64}
                            _ = v64
                            _ = v65
                        }
                        fmt.Println(v2)
                        _ = fi62
                    }
                case 1:
                    v4 -= int16(v2)
                case -985:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                default:
                    for fi66 := 0; fi66 < 16; fi66++ {
                        v67 := GenS5{F0: -1}
                        _ = fi66
                        _ = v67
                    }
                }
                _ = v61
            }
        case -435:
            fmt.Println(v5, v5)
        }
        _ = fi9
    }
    switch ("" + "a") {
    case "":
        for fi68 := 0; fi68 < 14; fi68++ {
            if false {
                for fi69 := 0; fi69 < 3; fi69++ {
                    var v70 float32 = 1e10
                    switch ("a" + "a") {
                    case "hello":
                        if ((v4 | v4) != (v4 | -2)) {
                            v71 := GenS4{F0: (fi69 & (fi69 + (-(fi68)))), F1: v2, F2: (-(v2))}
                            v72 := GenS7{F0: (v6 / ((v6) | 1))}
                            _ = v71
                            _ = v72
                        }
                    case "a":
                        fmt.Println(fi69, v4, fi68)
                    }
                    if len(v8) > 0 {
                        v8[0] = (fi69 >> ((v6) & 63))
                    }
                    _ = fi69
                    _ = v70
                }
            } else {
                v73 := GenS8{F0: 140, F1: "a", F2: 574}
                for fi74 := 0; fi74 < 8; fi74++ {
                    if ("" != "a") {
                        v5 -= v5
                        v4 += v4
                        fmt.Println(v6, fi74, fi74)
                    } else {
                        fmt.Println(v4, v2)
                    }
                    _ = fi74
                }
                var v75 int64 = -9223372036854775808
                _ = v73
                _ = v75
            }
            _ = fi68
        }
    case "foo bar":
        if (!(v4 <= 0)) {
            if (0 == 1) {
                if len(v8) > 0 {
                    v8[0] = 886
                }
                var v76 uint64 = 4294967295
                _ = v76
            } else {
                if ((-2.25 / ((float32(((3.14159 / 3.14159) + ((0.1 - (((0.1 - (float64(0.0) - float64(0.0))) + 0.1) / 0.1)) - (3.14159 + float64(float32((1.5 * (0.1 * 0.1)))))))) + (float32(3.14159) + 3.14159)) * float32(0.1))) != -2.25) {
                    switch 0 {
                    case 0:
                        var v77 uint64 = 962
                        _ = v77
                    }
                } else {
                    fmt.Println(v4)
                }
            }
            if (0.0 <= float64(-2.25)) {
                fmt.Println(v5)
                switch 4294967295 {
                case -2147483648:
                    var v78 int32 = -796
                    _ = v78
                case 952:
                    fmt.Println(v4, v2)
                case 203:
                    var v79 string = "a"
                    _ = v79
                }
            } else {
                var v80 uint = v6
                _ = v80
            }
        } else {
            for fi81 := 0; fi81 < 3; fi81++ {
                switch fi81 {
                case -9223372036854775808:
                    var v82 uint16 = 1
                    _ = v82
                case 732:
                    if (1e10 > float32(1.5)) {
                        v83 := v3[(fi81 << ((18446744073709551615) & 63))]
                        _ = v83
                    }
                }
                switch ("foo bar" + ("\n\t" + ("foo bar" + "foo bar"))) {
                case "a":
                    if (3.14159 == 3.14159) {
                        if false {
                            var v84 uint64 = 2
                            switch (fi81 - (fi81 + int(v2))) {
                            case 2147483647:
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            default:
                                for fi85 := 0; fi85 < 10; fi85++ {
                                    v84 -= (v84 + 397)
                                    switch int(v4) {
                                    case 127:
                                        v86 := GenS9{F0: 565}
                                        _ = v86
                                    case 32768:
                                        var v87 uint = v6
                                        _ = v87
                                    case -2147483648:
                                        v88 := GenS10{F0: fi85, F1: v2}
                                        _ = v88
                                    }
                                    _ = fi85
                                }
                            }
                            _ = v84
                        }
                        v5 = (v5 & (v5 << (((v6 / ((4294967295) | 1))) & 63)))
                        var v89 uint16 = 42
                        _ = v89
                    } else {
                        var v90 int16 = (^(v4))
                        var v91 int16 = (v90 | v4)
                        var v92 uint16 = 1
                        _ = v90
                        _ = v91
                        _ = v92
                    }
                case "\n\t":
                    var v93 int32 = -2147483648
                    _ = v93
                case "foo bar":
                    switch ("hello" + "a") {
                    case "\n\t":
                        v2 = (v2 ^ (^(v2)))
                    case "":
                        if (v6 != v6) {
                            if (!(2 >= 266)) {
                                v94 := GenS7{F0: v6}
                                _ = v94
                            }
                        } else {
                            delete(v3, fi81)
                            var v95 int = fi81
                            _ = v95
                        }
                    case "a":
                        v5 = v5
                    }
                }
                var v96 int8 = (v2 & int8(fi81))
                _ = fi81
                _ = v96
            }
            v2 = v2
            var v97 uint64 = 583
            _ = v97
        }
    case "hello":
        var v98 uint = 2
        _ = v98
    }
    var v99 int = 4294967295
    if false {
        for fi100 := 0; fi100 < 16; fi100++ {
            fmt.Println(v6, v5)
            _ = fi100
        }
        if len(v8) > 0 {
            v8[0] = v99
        }
        fmt.Println(v6)
    } else {
        if false {
            v5 -= v5
            var v101 int = 145
            _ = v101
        } else {
            for fi102 := 0; fi102 < 5; fi102++ {
                var v103 uint32 = v5
                fmt.Println(v99, v4)
                _ = fi102
                _ = v103
            }
        }
        var v104 int16 = v4
        var v105 int16 = v4
        _ = v104
        _ = v105
    }
    if (927 != 686) {
        if ((((("\n\t" + "") + ("\n\t" + ("" + ""))) + ("\n\t" + "foo bar")) + ("a" + ("foo bar" + ("" + "foo bar")))) >= "hello") {
            var v106 string = (("\n\t" + "hello") + "a")
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            fmt.Println(v6, v106, v6)
            _ = v106
        } else {
            var v107 uint32 = v5
            for fi108 := 0; fi108 < 5; fi108++ {
                v109 := GenS11{F0: ("" + "foo bar"), F1: v107, F2: v107}
                if false {
                    v110 := GenS5{F0: int16(v99)}
                    _ = v110
                }
                _ = fi108
                _ = v109
            }
            _ = v107
        }
        switch v99 {
        case -608:
            v111 := GenS5{F0: (v4 | v4)}
            _ = v111
        }
        v3[127] = v99
    } else {
        if (!((v99 % ((v99) | 1)) != (-(v99)))) {
            var v112 float32 = 0.0
            switch ("hello" + "a") {
            case "":
                var v113 int8 = v2
                _ = v113
            case "foo bar":
                if false {
                    v6 += v6
                }
            case "\n\t":
                v2 = (^(v2))
            }
            v114 := GenS1{F0: (v5 | v5), F1: 0, F2: 42}
            _ = v112
            _ = v114
        }
        var v115 uint16 = 255
        v115 = v115
        _ = v115
    }
    var v116 uint = v6
    var v117 uint64 = 454
    for fi118 := 0; fi118 < 15; fi118++ {
        var v119 int16 = (v4 << ((v116) & 63))
        v5 = 42
        switch ("" + "") {
        case "hello":
            var v120 uint = (v116 >> ((611) & 63))
            _ = v120
        case "foo bar":
            v116 = 4294967295
        case "":
            var v121 int32 = 0
            _ = v121
        }
        _ = fi118
        _ = v119
    }
    if len(v8) > 0 {
        v8[0] = v99
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v99
}
