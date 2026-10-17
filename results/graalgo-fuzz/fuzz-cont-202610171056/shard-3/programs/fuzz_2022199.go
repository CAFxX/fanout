package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint; F1 int64; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint; F1 bool; F2 uint32 }
type GenS6 struct { F0 int16; F1 uint64; F2 int64 }
type GenS7 struct { F0 int32; F1 int8; F2 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16; F1 bool; F2 string }

func main() {
    var v1 uint32 = 823
    var v2 bool = false
    v3 := make(map[string]string)
    v4 := GenS1{F0: (("a" + "a") + "a"), F1: 1}
    var v5 int = 2147483647
    v6 := GenS2{F0: 4294967295, F1: 509, F2: 42}
    if v2 {
        var v7 int = (v5 | v5)
        var v8 int32 = 42
        _ = v7
        _ = v8
    }
    switch "hello" {
    case "foo bar":
        v1 = v1
    case "":
        if (!(!((^(v1)) < v1))) {
            var v9 int16 = 1
            v9 = v9
            _ = v9
        } else {
            v1 = 4294967295
            var v10 bool = true
            fmt.Println(v5, v5)
            _ = v10
        }
    case "hello":
        v11 := GenS1{F0: "hello", F1: 0}
        _ = v11
    }
    fmt.Println(v2, v1, v5)
    for fi12 := 0; fi12 < 17; fi12++ {
        v5 -= (fi12 * (v5 * fi12))
        for fi13 := 0; fi13 < 7; fi13++ {
            if ("hello" > (("" + "") + "hello")) {
                var v14 uint16 = 32767
                _ = len(v3)
                v14 = v14
                switch ("foo bar" + ("" + "")) {
                case "hello":
                    fmt.Println(fi13, v14)
                    v15 := GenS1{F0: "a", F1: 100}
                    var v16 uint64 = 300
                    _ = v15
                    _ = v16
                case "":
                    v17 := GenS3{F0: v2, F1: 127}
                    _ = v17
                case "\n\t":
                    v1 = 65535
                    v5 += fi12
                    var v18 int = 935
                    _ = v18
                }
                _ = v14
            } else {
                var v19 bool = true
                if true {
                    v2 = (104 <= 0)
                    v5 += fi12
                } else {
                    var v20 uint64 = 4294967295
                    v20 = 0
                    v20 = v20
                    fmt.Println(v2, fi12)
                    v19 = (100 != -128)
                    _ = v20
                }
                v21 := GenS2{F0: 4294967295, F1: 42, F2: 1}
                _ = v19
                _ = v21
            }
            v1 += 65535
            v5 = 928
            _ = fi13
        }
        _ = fi12
    }
    v1 = 0
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int32 = -2147483648
        delete(v3, ((("hello" + (("hello" + ("foo bar" + ("hello" + "a"))) + "a")) + "hello") + ("a" + ("hello" + ("hello" + "a")))))
        _ = fi22
        _ = v23
    }
    switch 127 {
    case 9223372036854775807:
        for fi24 := 0; fi24 < 19; fi24++ {
            var v25 int16 = 1
            for fi26 := 0; fi26 < 1; fi26++ {
                if ((("foo bar" + "a") + ((("a" + ("" + "a")) + "") + (((("hello" + "hello") + "\n\t") + "foo bar") + "\n\t"))) >= "foo bar") {
                    if ((fi24 % ((v5) | 1)) < fi26) {
                        var v27 string = "\n\t"
                        v28 := v3[("hello" + "a")]
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(fi26, v2)
                    var v29 uint8 = 1
                    var v30 int8 = 50
                    _ = v29
                    _ = v30
                }
                _ = fi26
            }
            delete(v3, ("" + "\n\t"))
            _ = fi24
            _ = v25
        }
    case -661:
        v31 := GenS1{F0: "hello", F1: -128}
        _ = v31
    case 1:
        var v32 uint16 = 2
        _ = v32
    }
    switch "" {
    case "":
        v33 := GenS4{F0: 4294967295, F1: 135, F2: 1}
        _ = v33
    default:
        for fi34 := 0; fi34 < 4; fi34++ {
            if (9223372036854775807 < 0) {
                if (((0 < 944) || (!(v2 && (v5 < (v5 * (fi34 & 32767)))))) && (!v2)) {
                    for fi35 := 0; fi35 < 6; fi35++ {
                        switch "a" {
                        case "hello":
                            fmt.Println(v2, fi34)
                        case "a":
                            fmt.Println(v2, v1)
                        case "foo bar":
                            v36 := GenS2{F0: 0, F1: 560, F2: 4294967295}
                            _ = v36
                        default:
                            var v37 int16 = 42
                            v38 := []int{fi34, fi34}
                            var v39 uint64 = 0
                            v39 -= 4294967295
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        v1 = v1
                        v2 = v2
                        _ = fi35
                    }
                    v2 = (9223372036854775807 > -1)
                    if v2 {
                        if (false || v2) {
                            if v2 {
                                var v40 uint16 = 42
                                _ = v40
                            }
                        } else {
                            var v41 float64 = 1.5
                            var v42 int16 = -1
                            switch ("foo bar" + "foo bar") {
                            case "":
                                if v2 {
                                    if false {
                                        switch fi34 {
                                        case 0:
                                            v43 := GenS1{F0: ("hello" + "\n\t"), F1: 99}
                                            _ = v43
                                        default:
                                            v42 += v42
                                        }
                                        switch (("a" + "a") + "hello") {
                                        case "":
                                            fmt.Println(v2, v41, v5)
                                        default:
                                            var v44 uint32 = 2147483648
                                            _ = v44
                                        }
                                        fmt.Println(v41)
                                    }
                                    v45 := GenS5{F0: 255, F1: false, F2: 330}
                                    v46 := []int{(fi34 + (fi34 * (v5 ^ 723))), fi34}
                                    _ = v45
                                    _ = v46
                                }
                            }
                            _ = v41
                            _ = v42
                        }
                    }
                } else {
                    var v47 uint32 = v1
                    _ = v47
                }
            } else {
                fmt.Println(v1, v2)
                fmt.Println(v1, v5)
                for fi48 := 0; fi48 < 20; fi48++ {
                    v2 = v2
                    v49 := GenS2{F0: 678, F1: -1, F2: 473}
                    var v50 int8 = 82
                    v50 = v50
                    if false {
                        if (32767 < 1) {
                            var v51 uint = 4294967295
                            _ = v51
                        }
                        fmt.Println(fi48, v1)
                        if (-1 <= 0) {
                            v50 = 127
                        }
                    } else {
                        if v2 {
                            var v52 int8 = (^(v50))
                            _ = v52
                        } else {
                            v2 = v2
                            v2 = true
                            switch "\n\t" {
                            case "":
                                var v53 uint16 = 314
                                if ((v53 & 72) > v53) {
                                    fmt.Println(v50, v2)
                                    if v2 {
                                        switch ("" + "\n\t") {
                                        case "foo bar":
                                            if true {
                                                if v2 {
                                                    var v54 int8 = (v50 * (v50 % (((v50 << ((4294967295) & 63))) | 1)))
                                                    _ = v54
                                                }
                                                switch (fi48 & fi34) {
                                                case 935:
                                                    v50 -= v50
                                                case 32768:
                                                    var v55 float32 = 0.1
                                                    _ = v55
                                                case 2147483647:
                                                    if (((((-994 <= -1) || v2) || false) && true) && v2) {
                                                        if ((127 < 127) || (779 > v1)) {
                                                            v56 := []int{fi34, (v5 - (fi34 * 9223372036854775807)), (v5 >> ((998) & 63)), fi48}
                                                            if v2 {
                                                                if true {
                                                                    var v57 int = fi48
                                                                    var v58 int64 = 848
                                                                    v2 = v2
                                                                    delete(v3, "")
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    if v2 {
                                                                        var v59 string = ("foo bar" + ("a" + "foo bar"))
                                                                        v59 = ""
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            v60 := v3[""]
                                                                            _ = v60
                                                                        case "foo bar":
                                                                            v61 := GenS2{F0: 0, F1: -9223372036854775808, F2: 0}
                                                                            _ = v61
                                                                        case "":
                                                                            var v62 uint8 = 127
                                                                            _ = v62
                                                                        default:
                                                                            var v63 int = -92
                                                                            _ = v63
                                                                        }
                                                                        _ = v59
                                                                    }
                                                                }
                                                                fmt.Println(v50, fi48, v1)
                                                            } else {
                                                                if len(v56) >= 2 {
                                                                    v56 = v56[1:len(v56)]
                                                                }
                                                            }
                                                            switch (("foo bar" + ("a" + ("" + ("" + "")))) + "") {
                                                            case "hello":
                                                                var v64 uint32 = v1
                                                                _ = v64
                                                            case "foo bar":
                                                                var v65 int32 = -66
                                                                var v66 int32 = (v65 >> ((0) & 63))
                                                                _ = v65
                                                                _ = v66
                                                            case "\n\t":
                                                                fmt.Println(fi48)
                                                            }
                                                            _ = v56
                                                        } else {
                                                            var v67 int16 = -654
                                                            _ = v67
                                                        }
                                                        if (v2 && v2) {
                                                            fmt.Println(v1, fi34, v1)
                                                        } else {
                                                            switch ("hello" + "\n\t") {
                                                            case "":
                                                                var v68 int8 = 21
                                                                _ = v68
                                                            case "foo bar":
                                                                var v69 float32 = 0.0
                                                                _ = v69
                                                            }
                                                            var v70 uint64 = 42
                                                            if (1 == -1) {
                                                                v1 = v1
                                                            } else {
                                                                v71 := GenS6{F0: 0, F1: (v70 * 934), F2: int64(v70)}
                                                                var v72 int8 = (v50 << ((65535) & 63))
                                                                _ = v71
                                                                _ = v72
                                                            }
                                                            _ = v70
                                                        }
                                                        switch -2147483648 {
                                                        case -1:
                                                            if (1 != 2) {
                                                                v73 := GenS1{F0: "", F1: (v50 * v50)}
                                                                v74 := GenS7{F0: -600, F1: (v50 / ((42) | 1)), F2: (v50 ^ (^(v50)))}
                                                                if v2 {
                                                                    var v75 uint = 716
                                                                    _ = v75
                                                                } else {
                                                                    if (((("foo bar" + (("a" + "foo bar") + ("foo bar" + "a"))) + "foo bar") + "a") > "a") {
                                                                        if true {
                                                                            if v2 {
                                                                                fmt.Println(v53, v50)
                                                                                v3[("" + ("a" + ("" + (("foo bar" + ("foo bar" + ("foo bar" + "foo bar"))) + (("" + ("\n\t" + ("foo bar" + ("\n\t" + ((("hello" + "") + "") + ""))))) + "\n\t")))))] = "foo bar"
                                                                            } else {
                                                                                delete(v3, ("" + ""))
                                                                                v5 += fi48
                                                                                if (32767 > 42) {
                                                                                    v1 -= 703
                                                                                    fmt.Println(v2)
                                                                                } else {
                                                                                    var v76 int = fi48
                                                                                    switch (("" + "") + "foo bar") {
                                                                                    case "a":
                                                                                        var v77 uint8 = 42
                                                                                        _ = v77
                                                                                    default:
                                                                                        if false {
                                                                                            var v78 uint32 = v1
                                                                                            v76 = -1
                                                                                            switch (fi48 * fi34) {
                                                                                            case -698:
                                                                                                v79 := []int{fi34, -1}
                                                                                                _ = v79
                                                                                            case 2147483647:
                                                                                                var v80 float32 = float32((1e10 * 0.1))
                                                                                                _ = v80
                                                                                            case -9223372036854775808:
                                                                                                if (!v2) {
                                                                                                    v76 -= (-(v5))
                                                                                                    v81 := GenS1{F0: "hello", F1: (v50 + v50)}
                                                                                                    v76 -= int(v50)
                                                                                                    if v2 {
                                                                                                        v82 := GenS7{F0: 2147483647, F1: v50, F2: (v50 | v50)}
                                                                                                        switch 32768 {
                                                                                                        case 32768:
                                                                                                            var v83 float32 = float32(1e10)
                                                                                                            _ = v83
                                                                                                        case 4294967295:
                                                                                                            var v84 string = "a"
                                                                                                            if true {
                                                                                                                var v85 bool = v2
                                                                                                                _ = v85
                                                                                                            } else {
                                                                                                                var v86 bool = (!((v78 % (((v1 % ((42) | 1))) | 1)) > v1))
                                                                                                                if (-2.25 < ((1e10 * 1.5) * (((float32(0.0) * 0.0) + 0.0) + (-2.25 - 1e10)))) {
                                                                                                                    v3[v84] = v84
                                                                                                                    v50 -= v50
                                                                                                                }
                                                                                                                _ = v86
                                                                                                            }
                                                                                                            v87 := []int{v5, fi34}
                                                                                                            _ = v84
                                                                                                            _ = v87
                                                                                                        case -125:
                                                                                                            v76 = v76
                                                                                                        }
                                                                                                        v88 := GenS7{F0: 0, F1: (v50 << ((0) & 63)), F2: v50}
                                                                                                        switch "\n\t" {
                                                                                                        case "\n\t":
                                                                                                            delete(v3, "foo bar")
                                                                                                        case "hello":
                                                                                                            var v89 int = fi34
                                                                                                            _ = v89
                                                                                                        case "a":
                                                                                                            v90 := GenS5{F0: 191, F1: v2, F2: 65535}
                                                                                                            _ = v90
                                                                                                        default:
                                                                                                            v91 := []int{(fi34 % ((2147483647) | 1)), -609, fi34}
                                                                                                            _ = v91
                                                                                                        }
                                                                                                        v92 := v3[((("a" + "foo bar") + "hello") + "hello")]
                                                                                                        _ = v82
                                                                                                        _ = v88
                                                                                                        _ = v92
                                                                                                    }
                                                                                                    _ = v81
                                                                                                } else {
                                                                                                    if (uint8(v78) <= 0) {
                                                                                                        var v93 int = (^(fi34))
                                                                                                        _ = v93
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v94 := GenS8{F0: 0}
                                                                                            var v95 int8 = v50
                                                                                            var v96 uint = 580
                                                                                            _ = v78
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                            _ = v96
                                                                                        }
                                                                                    }
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                        }
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            var v97 uint64 = 42
                                                                            v97 -= v97
                                                                            _ = v97
                                                                        case "hello":
                                                                            v98 := GenS9{F0: 1, F1: (54 > 228), F2: (("a" + "\n\t") + "hello")}
                                                                            _ = v98
                                                                        case "":
                                                                            var v99 uint16 = (v53 & v53)
                                                                            _ = v99
                                                                        default:
                                                                            v53 += v53
                                                                        }
                                                                        fmt.Println(v53, v2)
                                                                    }
                                                                    fmt.Println(v53)
                                                                    v100 := []int{}
                                                                    _ = v100
                                                                }
                                                                fmt.Println(fi34, v2, v2)
                                                                v101 := GenS7{F0: 1, F1: -1, F2: v50}
                                                                if (9223372036854775807 != -9223372036854775808) {
                                                                    v102 := []int{}
                                                                    _ = v102
                                                                }
                                                                v103 := GenS1{F0: "\n\t", F1: v50}
                                                                _ = v101
                                                                _ = v103
                                                                _ = v73
                                                                _ = v74
                                                            } else {
                                                                v104 := GenS2{F0: 1, F1: 967, F2: 413}
                                                                _ = v104
                                                            }
                                                        case 1:
                                                            v5 = v5
                                                        }
                                                    } else {
                                                        var v105 uint8 = 255
                                                        _ = v105
                                                    }
                                                }
                                                var v106 bool = v2
                                                _ = v106
                                            }
                                        case "":
                                            var v107 int = -9223372036854775808
                                            _ = v107
                                        case "hello":
                                            v2 = (!v2)
                                            delete(v3, "a")
                                        }
                                        v108 := GenS9{F0: 1, F1: (uint32(v53) == (v1 % ((14) | 1))), F2: "hello"}
                                        var v109 bool = ((1.5 - -2.25) == ((0.1 - ((1.5 * 0.1) * 3.14159)) / 1.0))
                                        _ = v108
                                        _ = v109
                                    }
                                }
                                delete(v3, (("" + "hello") + "hello"))
                                _ = v53
                            case "\n\t":
                                switch fi34 {
                                case 1:
                                    v3[("\n\t" + "a")] = ((("foo bar" + "hello") + (("\n\t" + "\n\t") + ("\n\t" + "foo bar"))) + ("\n\t" + (("foo bar" + "hello") + "\n\t")))
                                case 128:
                                    v1 += 4294967295
                                case 9223372036854775807:
                                    var v110 uint16 = 65535
                                    _ = v110
                                default:
                                    v2 = (((-2.25 / 1e10) / ((3.14159 - 0.1) + 1.5)) > 0.1)
                                }
                            }
                        }
                    }
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            }
            _ = fi34
        }
    }
    v111 := v3[""]
    var v112 uint64 = 496
    v113 := []int{v5, 128}
    if false {
        var v114 string = (((v111 + (v111 + v111)) + "a") + v111)
        v1 += v1
        _ = v114
    } else {
        v5 = 128
        for fi115 := 0; fi115 < 13; fi115++ {
            var v116 uint32 = v1
            var v117 uint64 = (v112 / (((v112 & v112)) | 1))
            var v118 bool = v2
            _ = fi115
            _ = v116
            _ = v117
            _ = v118
        }
        if false {
            switch v111 {
            case "a":
                delete(v3, v111)
            default:
                var v119 uint = 42
                _ = v119
            }
        } else {
            if len(v113) >= 2 {
                v113 = v113[1:len(v113)]
            }
        }
    }
    v5 = v5
    fmt.Println(v1, v111)
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("slicelen v113:", len(v113), v113)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v111
    _ = v112
    _ = v113
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
