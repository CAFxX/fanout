package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 int; F1 int8; F2 bool }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 string; F1 int32; F2 uint }

func main() {
    v1 := make(map[string]string)
    var v2 string = ("a" + "foo bar")
    var v3 uint8 = 197
    var v4 float32 = (3.14159 - 1.5)
    switch -357 {
    case 0:
        var v5 int32 = -1
        _ = v5
    case 9223372036854775807:
        switch "\n\t" {
        case "hello":
            var v6 bool = (4294967295 > 42)
            _ = v6
        case "foo bar":
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 uint16 = 0
                var v9 float64 = 0.0
                _ = fi7
                _ = v8
                _ = v9
            }
        case "":
            delete(v1, "foo bar")
        default:
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 float32 = (float32(((1.5 - 1e10) + 1.5)) * ((3.14159 * (v4 / (float32(float64(v4)) / ((-2.25 * v4) * 0.1)))) + 3.14159))
                for fi12 := 0; fi12 < 2; fi12++ {
                    v13 := v1[v2]
                    v14 := GenS1{F0: -32768}
                    if false {
                        v15 := []int{fi10, 32767}
                        var v16 string = ("hello" + "a")
                        if true {
                            v17 := GenS2{F0: 509, F1: 2}
                            var v18 uint64 = 18446744073709551615
                            _ = v17
                            _ = v18
                        } else {
                            v3 = 1
                            if true {
                                v2 = v13
                                v16 = "hello"
                                if false {
                                    var v19 uint8 = (v3 & (v3 / (((v3 - v3)) | 1)))
                                    if (!true) {
                                        v3 += (v19 & v19)
                                        v2 = v13
                                        var v20 int32 = -2147483648
                                        _ = v20
                                    }
                                    switch -859 {
                                    case 2147483647:
                                        switch (fi12 % ((fi10) | 1)) {
                                        case 128:
                                            fmt.Println(v2)
                                        case -313:
                                            _ = len(v1)
                                        default:
                                            fmt.Println(v3)
                                        }
                                    case -1:
                                        if true {
                                            if ((v19 | (v19 * (v3 ^ 2))) == (v19 << ((2) & 63))) {
                                                if true {
                                                    var v21 uint8 = 128
                                                    var v22 int64 = 42
                                                    fmt.Println(v13, v13)
                                                    _ = v21
                                                    _ = v22
                                                } else {
                                                    fmt.Println(v16)
                                                    var v23 bool = ((!(-79 > 0)) || true)
                                                    _ = v23
                                                }
                                                var v24 string = v16
                                                _ = v24
                                            }
                                            v25 := v1[v2]
                                            var v26 int16 = -32768
                                            v25 = "hello"
                                            v27 := GenS1{F0: v26}
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        } else {
                                            v28 := GenS1{F0: 32767}
                                            _ = len(v15)
                                            _ = v28
                                        }
                                    }
                                    _ = len(v1)
                                    _ = v19
                                } else {
                                    v16 = (v13 + "")
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                }
                                v11 = (0.1 / v4)
                            }
                        }
                        _ = v15
                        _ = v16
                    } else {
                        var v29 uint = 549
                        v3 = (v3 >> ((v29) & 63))
                        fmt.Println(v13)
                        v4 = 0.1
                        v11 += (0.0 + v4)
                        _ = v29
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = fi10
                _ = v11
            }
        }
    case 1:
        fmt.Println(v2, v4, v4)
    default:
        v30 := GenS3{F0: v3, F1: "foo bar"}
        _ = v30
    }
    v3 -= v3
    v1[v2] = ((v2 + "hello") + (("hello" + ("\n\t" + v2)) + v2))
    for fi31 := 0; fi31 < 2; fi31++ {
        switch fi31 {
        case -9223372036854775808:
            v32 := []int{fi31, fi31}
            _ = v32
        case 128:
            v33 := GenS4{F0: 42, F1: true}
            _ = v33
        case -163:
            if (v2 == v2) {
                v1[v2] = v2
                for fi34 := 0; fi34 < 12; fi34++ {
                    var v35 int16 = 0
                    _ = fi34
                    _ = v35
                }
                if false {
                    var v36 int32 = 1
                    var v37 float32 = v4
                    _ = v36
                    _ = v37
                }
            }
        }
        var v38 uint32 = 1
        _ = fi31
        _ = v38
    }
    v2 = v2
    for fi39 := 0; fi39 < 4; fi39++ {
        v40 := []int{fi39, fi39, (fi39 | 9223372036854775807), fi39}
        _ = fi39
        _ = v40
    }
    if false {
        if (42 == 0) {
            for fi41 := 0; fi41 < 2; fi41++ {
                v42 := GenS1{F0: 32767}
                v43 := []int{(fi41 + fi41), (fi41 * (fi41 ^ fi41)), fi41}
                _ = fi41
                _ = v42
                _ = v43
            }
            v44 := GenS2{F0: uint64(v3), F1: 903}
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 7; fi45++ {
        v46 := []int{2147483647, (fi45 * (fi45 / ((0) | 1))), (fi45 & -1), fi45}
        switch v2 {
        case "":
            switch "" {
            case "":
                v47 := GenS2{F0: 1, F1: 592}
                _ = v47
            case "a":
                switch fi45 {
                case -95:
                    if true {
                        for fi48 := 0; fi48 < 14; fi48++ {
                            v49 := GenS3{F0: 128, F1: v2}
                            var v50 uint32 = 2
                            fmt.Println(v50, v50)
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        fmt.Println(v3, v4)
                        v51 := GenS4{F0: 646, F1: false}
                        _ = v51
                    } else {
                        for fi52 := 0; fi52 < 16; fi52++ {
                            switch (fi52 / (((fi45 + (fi52 + (fi45 ^ 475)))) | 1)) {
                            case 32767:
                                v3 -= (v3 ^ 79)
                            case 42:
                                var v53 uint64 = 473
                                _ = v53
                            default:
                                var v54 bool = false
                                _ = v54
                            }
                            var v55 int32 = 1
                            if true {
                                v56 := GenS4{F0: -552, F1: false}
                                var v57 uint32 = 1
                                var v58 int32 = 2147483647
                                _ = v56
                                _ = v57
                                _ = v58
                            }
                            _ = fi52
                            _ = v55
                        }
                        v59 := GenS5{F0: fi45, F1: 2}
                        _ = v59
                    }
                case 175:
                    for fi60 := 0; fi60 < 6; fi60++ {
                        if (32767 <= 1) {
                            var v61 int32 = 2147483647
                            _ = v61
                        }
                        v3 = v3
                        v2 = "a"
                        _ = fi60
                    }
                case 341:
                    var v62 float64 = (((-2.25 - (3.14159 + (float64(v4) + 0.0))) + 0.0) * float64(v4))
                    _ = v62
                }
            case "\n\t":
                v63 := GenS2{F0: 1, F1: 607}
                _ = v63
            default:
                for fi64 := 0; fi64 < 10; fi64++ {
                    if false {
                        v65 := GenS4{F0: 0, F1: true}
                        v66 := GenS1{F0: 42}
                        switch fi64 {
                        case 32767:
                            if (32767 != -216) {
                                v67 := GenS1{F0: -32768}
                                switch fi64 {
                                case 128:
                                    var v68 int = fi45
                                    _ = v68
                                default:
                                    v4 -= 1.5
                                }
                                fmt.Println(v3)
                                _ = v67
                            } else {
                                var v69 float64 = ((((1e10 + ((1e10 + ((3.14159 * (1e10 + (1e10 + -2.25))) - 3.14159)) * 0.0)) - 1.5) * 0.1) * 0.1)
                                var v70 float64 = ((v69 - 0.1) + (0.0 + (v69 + 3.14159)))
                                v3 -= v3
                                _ = v69
                                _ = v70
                            }
                        default:
                            var v71 uint = 635
                            _ = v71
                        }
                        _ = v65
                        _ = v66
                    }
                    _ = fi64
                }
            }
        case "\n\t":
            v4 = 1.5
        case "foo bar":
            switch -39 {
            case -318:
                var v72 bool = false
                _ = v72
            case 9223372036854775807:
                if ((v4 >= (-2.25 * float32((1e10 * 3.14159)))) || false) {
                    v2 = ""
                    v73 := GenS6{F0: uint64(v3), F1: 126}
                    _ = v73
                }
            case -1:
                var v74 uint = 42
                _ = v74
            }
        }
        _ = fi45
        _ = v46
    }
    v75 := GenS7{F0: 9223372036854775807, F1: 60, F2: (2147483647 == 2147483647)}
    if true {
        var v76 uint8 = v3
        var v77 uint64 = 42
        for fi78 := 0; fi78 < 12; fi78++ {
            if (false && (32767 == 1)) {
                var v79 int = (fi78 << ((65535) & 63))
                for fi80 := 0; fi80 < 7; fi80++ {
                    if true {
                        if true {
                            v81 := []int{(fi78 % ((fi78) | 1)), fi80, -8, fi80}
                            _ = v81
                        } else {
                            if ((1 == 2) || ((^(v3)) >= (v76 * (^(v3))))) {
                                var v82 uint = 4294967295
                                _ = v82
                            } else {
                                var v83 int32 = 42
                                var v84 int32 = 2147483647
                                _ = v83
                                _ = v84
                            }
                        }
                    }
                    if true {
                        if false {
                            v85 := GenS1{F0: 0}
                            v86 := []int{}
                            if false {
                                if true {
                                    if (2147483647 >= 1) {
                                        switch ("" + ("hello" + v2)) {
                                        case "\n\t":
                                            fmt.Println(v4, fi80, v4)
                                        case "hello":
                                            if false {
                                                fmt.Println(fi80, v77, v3)
                                            }
                                        case "":
                                            v77 += 1
                                        }
                                    }
                                }
                                var v87 int = fi80
                                switch ("foo bar" + v2) {
                                case "foo bar":
                                    v88 := GenS8{F0: (-(fi80)), F1: -1}
                                    _ = v88
                                case "":
                                    switch v79 {
                                    case 1:
                                        fmt.Println(v4, v76)
                                    case 32768:
                                        v3 -= 2
                                    case 255:
                                        if (127 >= -128) {
                                            if true {
                                                switch v2 {
                                                case "\n\t":
                                                    var v89 uint = 849
                                                    _ = v89
                                                }
                                                var v90 uint8 = 255
                                                v91 := GenS5{F0: fi78, F1: 4294967295}
                                                _ = v90
                                                _ = v91
                                            }
                                            v92 := GenS6{F0: 1, F1: 32767}
                                            v77 += (v77 & (^(v77)))
                                            _ = v92
                                        }
                                    }
                                default:
                                    v93 := GenS2{F0: (^(v77)), F1: (v77 ^ (v77 >> ((18446744073709551615) & 63)))}
                                    _ = v93
                                }
                                _ = v87
                            }
                            _ = len(v86)
                            fmt.Println(v4, fi78)
                            fmt.Println(v76, fi78)
                            _ = v85
                            _ = v86
                        } else {
                            var v94 string = (v2 + v2)
                            var v95 float64 = (0.1 + float64(v4))
                            _ = v94
                            _ = v95
                        }
                        if (3.14159 >= 0.0) {
                            var v96 bool = (1 > 1)
                            var v97 int64 = 42
                            _ = v96
                            _ = v97
                        }
                        v3 = v76
                    } else {
                        v76 = v3
                        v3 -= 42
                    }
                    v2 = v2
                    v98 := []int{-360}
                    if (1e10 >= 0.0) {
                        v99 := GenS9{F0: 499, F1: 255, F2: v76}
                        v3 = 72
                        v100 := GenS10{F0: (false || false)}
                        _ = v100
                        _ = v99
                    }
                    if len(v98) > 0 {
                        v98[0] = 9223372036854775807
                    }
                    _ = fi80
                    _ = v98
                }
                _ = v79
            } else {
                var v101 bool = (!false)
                for fi102 := 0; fi102 < 4; fi102++ {
                    var v103 uint64 = v77
                    var v104 int8 = 0
                    if (!(!v101)) {
                        var v105 uint8 = (v3 ^ (v76 / (((v3 & (v76 * (v3 >> ((223) & 63))))) | 1)))
                        delete(v1, "\n\t")
                        _ = v105
                    } else {
                        v2 = "hello"
                        if (-128 != -128) {
                            if (18446744073709551615 >= 0) {
                                if v101 {
                                    v103 = (v77 * (^(v77)))
                                } else {
                                    v76 = 128
                                }
                                v106 := []int{9223372036854775807, (fi78 | fi102)}
                                var v107 uint16 = 42
                                _ = v106
                                _ = v107
                            } else {
                                v4 = 3.14159
                            }
                            v108 := []int{(fi102 & 128)}
                            var v109 int = 0
                            _ = v108
                            _ = v109
                        } else {
                            v110 := GenS8{F0: fi102, F1: 42}
                            var v111 int16 = -32768
                            _ = v110
                            _ = v111
                        }
                        if ((v4 * v4) == (float32((1e10 - (0.0 * 0.1))) + v4)) {
                            var v112 int = (fi102 >> ((18446744073709551615) & 63))
                            v104 = v104
                            var v113 string = v2
                            _ = v112
                            _ = v113
                        } else {
                            v3 = (v3 + v76)
                            var v114 bool = (1 < 1)
                            v114 = true
                            _ = v114
                        }
                    }
                    _ = fi102
                    _ = v103
                    _ = v104
                }
                for fi115 := 0; fi115 < 18; fi115++ {
                    v116 := GenS10{F0: v101}
                    v117 := v1[v2]
                    v117 = (v2 + "foo bar")
                    _ = fi115
                    _ = v116
                    _ = v117
                }
                _ = v101
            }
            var v118 uint = 42
            _ = fi78
            _ = v118
        }
        _ = v76
        _ = v77
    }
    v4 = v4
    v119 := GenS11{F0: ((v2 + "hello") + (v2 + "")), F1: 42, F2: 42}
    v2 = ((v2 + "a") + v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v75
}
