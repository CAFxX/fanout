package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint16; F2 bool }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int8; F2 int8 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 int64; F1 bool; F2 uint16 }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 uint64 }

func main() {
    v1 := []string{(("foo bar" + "foo bar") + (("" + ("" + "")) + "foo bar")), ("\n\t" + "hello")}
    v2 := make(map[uint32]uint32)
    v3 := []uint32{65535, 0}
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    if (-9223372036854775808 <= 42) {
        var v6 uint32 = 2
        var v7 float32 = 0.0
        _ = v6
        _ = v7
    } else {
        if true {
            var v8 int16 = 42
            for fi9 := 0; fi9 < 10; fi9++ {
                if true {
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    v10 := GenS1{F0: false}
                    v8 = v8
                    _ = v10
                }
                _ = fi9
            }
            v5["a"] = "\n\t"
            _ = v8
        } else {
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 uint = 18446744073709551615
                v12 -= v12
                v13 := v5[("foo bar" + "a")]
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 int = 0
            for fi15 := 0; fi15 < 14; fi15++ {
                for fi16 := 0; fi16 < 14; fi16++ {
                    switch "a" {
                    case "\n\t":
                        v17 := GenS1{F0: false}
                        _ = v17
                    }
                    var v18 float32 = 3.14159
                    _ = fi16
                    _ = v18
                }
                _ = fi15
            }
            _ = v14
        }
        var v19 int64 = 9223372036854775807
        _ = v19
    }
    v3 = append(v3, 2)
    var v20 int = 1
    v20 = v20
    var v21 int8 = 100
    if true {
        var v22 int = (-(v20))
        _ = v22
    }
    var v23 int8 = v21
    if true {
        v24 := GenS1{F0: true}
        for fi25 := 0; fi25 < 10; fi25++ {
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := v2[236]
                if true {
                    if ((3.14159 >= -2.25) || (1 < 2)) {
                        _ = len(v3)
                        v23 = -128
                        v28 := GenS1{F0: true}
                        _ = v28
                    }
                    var v29 float64 = ((1e10 - 1e10) - 1e10)
                    v20 += -100
                    _ = v29
                } else {
                    if false {
                        if (19 < 128) {
                            switch fi25 {
                            case -2147483648:
                                var v30 uint64 = 2
                                _ = v30
                            case 4294967295:
                                delete(v2, v27)
                            case -105:
                                fmt.Println(v20, fi25, v23)
                            }
                            switch -902 {
                            case 810:
                                v23 = (^(v23))
                            case -2147483648:
                                v27 += 65535
                            case -829:
                                v31 := GenS2{F0: 255}
                                var v32 string = "foo bar"
                                _ = v31
                                _ = v32
                            default:
                                if true {
                                    v27 = (v27 / (((^(v27))) | 1))
                                }
                                var v33 int8 = (^(v23))
                                v27 = v27
                                _ = v33
                            }
                        } else {
                            var v34 uint32 = v27
                            var v35 string = ""
                            _ = v34
                            _ = v35
                        }
                        if false {
                            var v36 uint16 = 32767
                            switch fi25 {
                            case 112:
                                fmt.Println(fi25)
                            case 127:
                                if (!(true && true)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v37 := GenS1{F0: (false || true)}
                                    v27 -= v27
                                    _ = v37
                                }
                            default:
                                switch "\n\t" {
                                case "foo bar":
                                    fmt.Println(v23, v21)
                                case "hello":
                                    if true {
                                        fmt.Println(v27, v36)
                                        v38 := v2[(v27 >> ((42) & 63))]
                                        _ = v38
                                    } else {
                                        var v39 uint32 = (v27 ^ (v27 - v27))
                                        _ = v39
                                    }
                                }
                            }
                            switch "a" {
                            case "hello":
                                v21 -= 91
                            case "foo bar":
                                v40 := GenS2{F0: 65535}
                                _ = v40
                            case "\n\t":
                                if (true && false) {
                                    v21 += (v23 ^ (v21 + v23))
                                    var v41 string = "\n\t"
                                    if (true && false) {
                                        if false {
                                            var v42 float32 = ((1e10 - 0.0) / 0.1)
                                            _ = v42
                                        } else {
                                            v36 += (v36 * v36)
                                        }
                                        fmt.Println(v21, fi25, fi26)
                                        v5[(("hello" + (v41 + v41)) + "\n\t")] = ("hello" + v41)
                                        var v43 float64 = 1e10
                                        var v44 int64 = -819
                                        _ = v43
                                        _ = v44
                                    } else {
                                        v41 = (v41 + ((v41 + v41) + v41))
                                    }
                                    _ = v41
                                }
                            }
                            _ = v36
                        } else {
                            fmt.Println(v27)
                        }
                        v27 += (v27 >> ((0) & 63))
                    }
                    switch v20 {
                    case -685:
                        v45 := GenS1{F0: (1 >= -1)}
                        _ = v45
                    case 9223372036854775807:
                        var v46 int = v20
                        _ = v46
                    case -2147483648:
                        if (255 > 0) {
                            v21 -= 42
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if false {
                                var v47 uint8 = 0
                                var v48 float64 = -2.25
                                _ = len(v5)
                                _ = v47
                                _ = v48
                            }
                            if true {
                                if (!(true || (!true))) {
                                    delete(v5, "\n\t")
                                    fmt.Println(v23, v20, fi26)
                                    if (681 != 18446744073709551615) {
                                        if (!("foo bar" < ("foo bar" + "\n\t"))) {
                                            v49 := GenS3{F0: (123 != 2), F1: 42, F2: ((v20 % ((fi26) | 1)) > 32768)}
                                            var v50 int = (fi26 ^ fi25)
                                            if true {
                                                if false {
                                                    v50 -= 518
                                                    var v51 string = "foo bar"
                                                    switch ((v51 + "hello") + (v51 + "hello")) {
                                                    case "":
                                                        var v52 uint8 = 123
                                                        _ = v52
                                                    default:
                                                        var v53 uint64 = 825
                                                        switch "a" {
                                                        case "hello":
                                                            if true {
                                                                switch v51 {
                                                                case "foo bar":
                                                                    fmt.Println(v50, v51, v23)
                                                                case "\n\t":
                                                                    v5["foo bar"] = v51
                                                                case "a":
                                                                    v54 := GenS4{F0: (v23 % (((v21 * (v23 >> ((792) & 63)))) | 1)), F1: 2}
                                                                    _ = v54
                                                                }
                                                                switch v51 {
                                                                case "hello":
                                                                    var v55 string = (v51 + v51)
                                                                    _ = v55
                                                                default:
                                                                    v20 -= v50
                                                                }
                                                                var v56 int = v50
                                                                switch (fi25 % ((v20) | 1)) {
                                                                case -9223372036854775808:
                                                                    v20 -= v56
                                                                case -2147483648:
                                                                    v57 := GenS3{F0: (true && ((!(!false)) || ((v53 - v53) != 937))), F1: 129, F2: ((v27 << ((0) & 63)) <= v27)}
                                                                    _ = v57
                                                                }
                                                                _ = v56
                                                            }
                                                        case "a":
                                                            if (!(!((false && (!(4294967295 > v27))) && (v53 <= v53)))) {
                                                                var v58 string = v51
                                                                fmt.Println(v21, v20)
                                                                if ((((((v58 + (v58 + ("foo bar" + (v51 + v58)))) + "hello") + v58) + v51) + v51) < v58) {
                                                                    v59 := GenS5{F0: 2147483648}
                                                                    var v60 int8 = (v23 << ((255) & 63))
                                                                    v27 = (v27 >> ((673) & 63))
                                                                    var v61 int8 = 15
                                                                    _ = v59
                                                                    _ = v60
                                                                    _ = v61
                                                                } else {
                                                                    var v62 string = (("\n\t" + v58) + v51)
                                                                    var v63 float32 = 1.5
                                                                    switch -1 {
                                                                    case 42:
                                                                        v53 = v53
                                                                    case -272:
                                                                        v23 -= v21
                                                                    case 815:
                                                                        v21 = v21
                                                                    }
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                                v21 -= v21
                                                                if (9223372036854775807 >= 42) {
                                                                    v3 = append(v3, 2147483647)
                                                                    var v64 float64 = 0.1
                                                                    _ = v64
                                                                } else {
                                                                    var v65 float32 = 1.5
                                                                    switch int(v21) {
                                                                    case 9223372036854775807:
                                                                        delete(v2, (v27 - v27))
                                                                    case 32767:
                                                                        v66 := GenS3{F0: true, F1: 65535, F2: ((v58 + ((v51 + (v51 + v51)) + ((((v51 + v51) + "foo bar") + v51) + v58))) == (("" + v51) + (v58 + v58)))}
                                                                        _ = v66
                                                                    default:
                                                                        v20 = (fi26 ^ 42)
                                                                    }
                                                                    var v67 uint8 = 42
                                                                    _ = v65
                                                                    _ = v67
                                                                }
                                                                fmt.Println(v21)
                                                                _ = v58
                                                            }
                                                            var v68 int32 = -2147483648
                                                            _ = v68
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        _ = v53
                                                    }
                                                    _ = v51
                                                } else {
                                                    v69 := GenS4{F0: v23, F1: 2}
                                                    if len(v1) > 0 {
                                                        v1[0] = (("\n\t" + "foo bar") + "a")
                                                    }
                                                    v4[fi26] = v50
                                                    _ = v69
                                                }
                                                switch fi25 {
                                                case 2147483647:
                                                    v3 = append(v3, (v27 | v27))
                                                case 128:
                                                    var v70 int = (^(fi25))
                                                    _ = v70
                                                default:
                                                    v71 := GenS6{F0: 392, F1: -7, F2: 42}
                                                    _ = v71
                                                }
                                                v23 += -1
                                            } else {
                                                if (v21 >= -128) {
                                                    if (fi26 <= v50) {
                                                        v23 += v21
                                                        if len(v1) > 0 {
                                                            v1[0] = ("hello" + "hello")
                                                        }
                                                    } else {
                                                        if len(v1) > 0 {
                                                            v1[0] = (((("foo bar" + "") + ("\n\t" + "")) + "foo bar") + "hello")
                                                        }
                                                    }
                                                    v72 := GenS6{F0: 1, F1: v21, F2: (^(v21))}
                                                    _ = v72
                                                } else {
                                                    switch fi25 {
                                                    case 2147483647:
                                                        v73 := GenS2{F0: 65535}
                                                        _ = v73
                                                    case 767:
                                                        var v74 float64 = 0.1
                                                        _ = v74
                                                    case 9223372036854775807:
                                                        switch (-(v20)) {
                                                        case 127:
                                                            var v75 uint16 = 65535
                                                            _ = v75
                                                        default:
                                                            v50 += 812
                                                        }
                                                    default:
                                                        var v76 uint32 = 2
                                                        _ = v76
                                                    }
                                                }
                                                if (2 > 255) {
                                                    var v77 uint = 2
                                                    v78 := GenS2{F0: v77}
                                                    var v79 int16 = 32767
                                                    _ = v77
                                                    _ = v78
                                                    _ = v79
                                                } else {
                                                    fmt.Println(v21)
                                                    if ((1 < 32767) || (!true)) {
                                                        var v80 uint8 = 74
                                                        _ = v80
                                                    } else {
                                                        fmt.Println(v50)
                                                        var v81 int = -878
                                                        var v82 float64 = (0.0 * float64(float32(0.1)))
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                }
                                            }
                                            _ = v49
                                            _ = v50
                                        } else {
                                            if (!false) {
                                                var v83 int32 = -884
                                                _ = v83
                                            }
                                            v27 -= v27
                                        }
                                    } else {
                                        v20 += fi25
                                        var v84 string = ("\n\t" + ("foo bar" + ("" + "")))
                                        var v85 uint64 = 880
                                        _ = v84
                                        _ = v85
                                    }
                                } else {
                                    if false {
                                        var v86 int8 = (v21 + (^(v21)))
                                        var v87 int32 = 2147483647
                                        var v88 uint16 = 0
                                        _ = v86
                                        _ = v87
                                        _ = v88
                                    }
                                    v89 := v5[""]
                                    _ = v89
                                }
                                v90 := GenS1{F0: true}
                                _ = v90
                            }
                            v20 -= 0
                            var v91 int64 = -9223372036854775808
                            if true {
                                switch "a" {
                                case "foo bar":
                                    var v92 uint16 = 2
                                    _ = v92
                                case "hello":
                                    v27 -= (^(v27))
                                default:
                                    fmt.Println(v91, v20)
                                    switch fi25 {
                                    case 42:
                                        switch (fi25 >> ((1) & 63)) {
                                        case -9223372036854775808:
                                            if true {
                                                var v93 uint = 65535
                                                fmt.Println(v27, fi26, fi25)
                                                switch -955 {
                                                case 9223372036854775807:
                                                    v94 := GenS6{F0: 18446744073709551615, F1: (v21 * v23), F2: v21}
                                                    _ = v94
                                                case 4294967295:
                                                    var v95 uint32 = (v27 % ((v27) | 1))
                                                    _ = v95
                                                default:
                                                    switch (fi26 << ((2) & 63)) {
                                                    case 398:
                                                        var v96 float32 = 3.14159
                                                        _ = v96
                                                    }
                                                }
                                                _ = v93
                                            }
                                        case 127:
                                            if true {
                                                if ((v91 << ((590) & 63)) != -1) {
                                                    if ((v91 % ((v91) | 1)) > v91) {
                                                        v20 = fi26
                                                        v20 -= 127
                                                        v20 += v20
                                                    } else {
                                                        v97 := GenS7{F0: 0, F1: "foo bar"}
                                                        var v98 string = "foo bar"
                                                        v27 = (v27 + (^(v27)))
                                                        v99 := GenS8{F0: (v91 * v91), F1: ((^(v20)) <= (fi26 << ((4294967295) & 63))), F2: 2}
                                                        _ = v97
                                                        _ = v98
                                                        _ = v99
                                                    }
                                                }
                                            } else {
                                                v20 = (fi25 + fi25)
                                            }
                                        case 1:
                                            var v100 int8 = v23
                                            _ = v100
                                        }
                                    case 4294967295:
                                        _ = len(v2)
                                    case 9223372036854775807:
                                        v101 := GenS5{F0: v27}
                                        _ = v101
                                    }
                                    var v102 float32 = -2.25
                                    _ = v102
                                }
                                if false {
                                    var v103 int32 = -106
                                    v104 := GenS6{F0: 18446744073709551615, F1: v23, F2: 127}
                                    _ = v103
                                    _ = v104
                                } else {
                                    if true {
                                        var v105 int = (-(fi26))
                                        fmt.Println(v91, v20, v21)
                                        fmt.Println(fi26)
                                        _ = v105
                                    }
                                }
                            }
                            _ = v91
                        } else {
                            v106 := GenS9{F0: (v20 ^ (^(fi25))), F1: (^(v27))}
                            v107 := GenS10{F0: 4294967295}
                            _ = v106
                            _ = v107
                        }
                        var v108 uint8 = 2
                        v109 := GenS7{F0: 2147483647, F1: "\n\t"}
                        _ = v108
                        _ = v109
                    }
                    fmt.Println(fi26, fi25)
                }
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        if ("hello" < "a") {
            v110 := v5[("\n\t" + ("hello" + ("\n\t" + "a")))]
            _ = len(v1)
            fmt.Println(v21)
            _ = v110
        }
        _ = v24
    }
    var v111 uint = 2
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        if (166 <= 2) {
            for fi112 := 0; fi112 < 4; fi112++ {
                for fi113 := 0; fi113 < 13; fi113++ {
                    var v114 int = (fi112 - (fi113 + v20))
                    _ = fi113
                    _ = v114
                }
                for fi115 := 0; fi115 < 14; fi115++ {
                    v116 := v2[2147483647]
                    _ = fi115
                    _ = v116
                }
                _ = fi112
            }
            var v117 uint64 = 443
            var v118 string = "foo bar"
            _ = v117
            _ = v118
        } else {
            var v119 int32 = -1
            var v120 uint64 = 4294967295
            _ = v119
            _ = v120
        }
    }
    if (2 >= 42) {
        if (v20 < v20) {
            var v121 int32 = 1
            for fi122 := 0; fi122 < 14; fi122++ {
                v123 := GenS7{F0: (v121 / (((v121 * -2147483648)) | 1)), F1: ("foo bar" + "\n\t")}
                _ = fi122
                _ = v123
            }
            _ = v121
        }
    }
    if (255 <= 1) {
        var v124 uint = v111
        _ = v124
    }
    v125 := GenS8{F0: -9223372036854775808, F1: false, F2: 281}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v111:", v111)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v111
    _ = v125
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
