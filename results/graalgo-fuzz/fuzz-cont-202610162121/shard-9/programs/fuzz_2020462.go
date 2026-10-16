package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint16; F1 int16; F2 bool }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS5 struct { F0 string; F1 uint64; F2 bool }
type GenS6 struct { F0 int64; F1 int; F2 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 int8; F1 int16; F2 int32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    var v3 uint8 = 2
    fmt.Println(v3)
    var v4 int = 1
    if true {
        v5 := []int{}
        var v6 float32 = 1.5
        _ = v5
        _ = v6
    }
    fmt.Println(v3, v4)
    if (1 <= 255) {
        v7 := GenS1{F0: 42, F1: 0}
        switch (v4 << ((322) & 63)) {
        case 889:
            v8 := GenS2{F0: 207, F1: 32767, F2: (v4 <= v4)}
            _ = v8
        case -256:
            v4 += v4
        case 9223372036854775807:
            for fi9 := 0; fi9 < 19; fi9++ {
                v1[14] = true
                for fi10 := 0; fi10 < 14; fi10++ {
                    if true {
                        var v11 int64 = 255
                        if false {
                            if ((v11 * v11) < (v11 - (v11 + (v11 ^ v11)))) {
                                v12 := GenS1{F0: 2147483647, F1: (v11 << ((269) & 63))}
                                var v13 float32 = ((0.0 * 1.5) / float32(1e10))
                                v14 := v2["a"]
                                if (int32(v11) != 0) {
                                    var v15 bool = (!((0.0 > -2.25) && false))
                                    v16 := []int{(fi9 / (((fi9 % (((v4 ^ (v4 & v4))) | 1))) | 1)), 4294967295}
                                    fmt.Println(fi10, fi10, v14)
                                    _ = v15
                                    _ = v16
                                }
                                _ = v12
                                _ = v13
                                _ = v14
                            } else {
                                if ((!(1.5 <= 0.0)) && (-1 < 1000)) {
                                    v3 = 94
                                    if ((432 <= 885) && false) {
                                        var v17 int8 = 42
                                        if ((42 >= uint(v3)) || (42 > 65535)) {
                                            delete(v2, ("hello" + "a"))
                                            if (("hello" > "\n\t") && (true || ((^(v17)) == v17))) {
                                                if ((true && (false || (!false))) || ((70 < 1) && true)) {
                                                    v18 := []int{fi9, (fi9 ^ 2147483647)}
                                                    var v19 int = 128
                                                    _ = v18
                                                    _ = v19
                                                }
                                                v20 := GenS3{F0: v3, F1: 817}
                                                _ = v20
                                            } else {
                                                var v21 string = ("a" + "")
                                                _ = v21
                                            }
                                        }
                                        _ = v17
                                    } else {
                                        if false {
                                            v11 = v11
                                            v22 := GenS2{F0: 42, F1: -841, F2: ("a" <= ("" + ""))}
                                            fmt.Println(v3, v3)
                                            _ = v22
                                        } else {
                                            v2["hello"] = 65535
                                            v23 := GenS3{F0: 42, F1: 4294967295}
                                            _ = v23
                                        }
                                        var v24 uint8 = v3
                                        _ = v24
                                    }
                                }
                            }
                            if (741 == 65535) {
                                if (1.5 != 0.1) {
                                    v25 := []int{}
                                    if true {
                                        var v26 int16 = 32767
                                        _ = v26
                                    }
                                    v27 := GenS4{F0: 65535, F1: (-(v11)), F2: 0}
                                    _ = v25
                                    _ = v27
                                }
                            }
                        } else {
                            switch ("hello" + "") {
                            case "foo bar":
                                var v28 uint64 = 42
                                _ = v28
                            }
                        }
                        _ = v11
                    }
                    v3 = v3
                    _ = fi10
                }
                _ = fi9
            }
        default:
            for fi29 := 0; fi29 < 1; fi29++ {
                v30 := GenS5{F0: "foo bar", F1: 2, F2: true}
                v3 = v3
                v3 += (^(v3))
                _ = fi29
                _ = v30
            }
        }
        _ = v7
    }
    fmt.Println(v3, v4)
    var v31 uint8 = 248
    if (v4 > v4) {
        v32 := GenS3{F0: (^(v3)), F1: 18446744073709551615}
        _ = v32
    } else {
        v4 = (^(v4))
    }
    v33 := GenS6{F0: -436, F1: v4, F2: v4}
    for fi34 := 0; fi34 < 15; fi34++ {
        var v35 float32 = 3.14159
        _ = fi34
        _ = v35
    }
    switch "foo bar" {
    case "a":
        var v36 int16 = 1000
        _ = v36
    case "hello":
        v31 = v3
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        var v38 int32 = 42
        var v39 bool = true
        _ = fi37
        _ = v38
        _ = v39
    }
    if (false && (v3 > 123)) {
        var v40 uint16 = 863
        _ = v40
    } else {
        switch 128 {
        case -376:
            fmt.Println(v31)
        case 2147483647:
            for fi41 := 0; fi41 < 6; fi41++ {
                if false {
                    if false {
                        if (42 >= 42) {
                            fmt.Println(fi41)
                            if (!(1 <= -100)) {
                                var v42 int16 = 684
                                v3 = 65
                                v3 += (v31 << ((1) & 63))
                                _ = v42
                            }
                            switch "foo bar" {
                            case "a":
                                if false {
                                    v4 = 127
                                } else {
                                    fmt.Println(v3)
                                }
                            default:
                                if (2147483648 <= 0) {
                                    delete(v1, 0)
                                    var v43 uint32 = 984
                                    switch "\n\t" {
                                    case "":
                                        for fi44 := 0; fi44 < 1; fi44++ {
                                            switch "hello" {
                                            case "\n\t":
                                                var v45 int32 = 42
                                                _ = v45
                                            case "a":
                                                var v46 uint64 = 42
                                                _ = v46
                                            case "foo bar":
                                                v47 := GenS7{F0: v4}
                                                _ = v47
                                            }
                                            fmt.Println(fi44, v43, fi44)
                                            switch fi44 {
                                            case -560:
                                                if (true || false) {
                                                    if true {
                                                        v1[2147483648] = (42 > -245)
                                                        v4 += 128
                                                        v43 = (v43 | 2147483648)
                                                    } else {
                                                        v31 += v31
                                                    }
                                                    v48 := GenS8{F0: 32767}
                                                    v49 := GenS9{F0: 9223372036854775807, F1: 254}
                                                    _ = v48
                                                    _ = v49
                                                }
                                                v4 = fi44
                                                v4 = (v4 | (fi41 + (fi41 + 0)))
                                            case 0:
                                                var v50 uint64 = 940
                                                _ = v50
                                            }
                                            v51 := []int{fi41}
                                            if (-1 >= -1) {
                                                var v52 int16 = 1
                                                var v53 int8 = -1
                                                fmt.Println(v43, v4, fi41)
                                                _ = v52
                                                _ = v53
                                            }
                                            v51 = append(v51, v4)
                                            _ = fi44
                                            _ = v51
                                        }
                                    default:
                                        v31 = (^(v3))
                                    }
                                    _ = v43
                                } else {
                                    for fi54 := 0; fi54 < 8; fi54++ {
                                        v3 -= (v31 + v31)
                                        _ = fi54
                                    }
                                    var v55 uint = 280
                                    fmt.Println(fi41)
                                    _ = v55
                                }
                            }
                        }
                    } else {
                        if false {
                            var v56 uint32 = 2147483647
                            switch "\n\t" {
                            case "a":
                                var v57 int8 = 127
                                _ = v57
                            case "foo bar":
                                for fi58 := 0; fi58 < 11; fi58++ {
                                    var v59 float64 = -2.25
                                    v60 := []int{(v4 + (-(fi41))), v4}
                                    switch -559 {
                                    case 127:
                                        if false {
                                            if (true || (true && (-32768 > 264))) {
                                                if len(v60) >= 2 {
                                                    v60 = v60[1:len(v60)]
                                                }
                                            }
                                        }
                                    case 0:
                                        v4 += 128
                                    case 32767:
                                        var v61 int64 = -1
                                        _ = v61
                                    }
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                            default:
                                fmt.Println(v56)
                            }
                            v62 := GenS5{F0: ("foo bar" + "a"), F1: 2, F2: false}
                            _ = v56
                            _ = v62
                        } else {
                            if (-2147483648 >= 379) {
                                v4 += v4
                                v4 = 128
                            } else {
                                for fi63 := 0; fi63 < 1; fi63++ {
                                    if false {
                                        if true {
                                            v3 += v31
                                            var v64 int = (v4 * -9223372036854775808)
                                            _ = v64
                                        }
                                        if (9223372036854775807 > 287) {
                                            v65 := GenS4{F0: 279, F1: -1, F2: 585}
                                            v31 = 0
                                            _ = v65
                                        } else {
                                            if true {
                                                if true {
                                                    v66 := []int{}
                                                    _ = v66
                                                } else {
                                                    _ = len(v1)
                                                    v67 := []int{(fi63 % ((59) | 1)), fi63, (fi63 - 127), -9223372036854775808}
                                                    _ = v67
                                                }
                                                var v68 uint = 2
                                                _ = v68
                                            }
                                            if (uint32(v3) < 2) {
                                                var v69 uint32 = 648
                                                if (!((!(1 <= -32768)) && (1.5 <= 0.1))) {
                                                    fmt.Println(v69, fi41, v3)
                                                    fmt.Println(v4, v69)
                                                    v70 := v2[("a" + ("\n\t" + ("" + "\n\t")))]
                                                    _ = v70
                                                }
                                                v1[(v69 << ((uint(v69)) & 63))] = false
                                                var v71 uint8 = v31
                                                var v72 uint16 = 306
                                                _ = v69
                                                _ = v71
                                                _ = v72
                                            } else {
                                                v73 := v2["foo bar"]
                                                fmt.Println(fi63, v4)
                                                if ((1e10 + ((1.5 / 0.1) / 1.0)) != float64(1.5)) {
                                                    v74 := []int{fi63, 0, v4}
                                                    v3 = (v3 % ((v31) | 1))
                                                    v3 -= v31
                                                    _ = v74
                                                } else {
                                                    var v75 uint32 = 1
                                                    _ = v75
                                                }
                                                v1[(^(v73))] = (42 != 9223372036854775807)
                                                delete(v2, "hello")
                                                _ = v73
                                            }
                                        }
                                        v4 = (fi63 / (((fi41 * -1)) | 1))
                                    } else {
                                        v76 := []int{(^(v4)), (^(fi63)), (^(fi63)), -2147483648}
                                        v4 = 32767
                                        _ = v76
                                    }
                                    switch ("hello" + (("foo bar" + "hello") + "hello")) {
                                    case "\n\t":
                                        var v77 string = ("hello" + (((("\n\t" + "a") + "") + ("hello" + "")) + ("a" + (("a" + "\n\t") + ("a" + ("hello" + "foo bar"))))))
                                        _ = v77
                                    default:
                                        v78 := GenS3{F0: (v3 ^ v3), F1: 4294967295}
                                        _ = v78
                                    }
                                    v4 = (fi63 + (v4 / ((0) | 1)))
                                    switch fi63 {
                                    case -121:
                                        fmt.Println(fi41, v31)
                                    case 42:
                                        switch (v4 >> ((381) & 63)) {
                                        case -674:
                                            v3 = (v3 << ((65535) & 63))
                                        case 32768:
                                            v79 := GenS4{F0: 702, F1: -944, F2: 42}
                                            _ = v79
                                        default:
                                            var v80 int = -2147483648
                                            _ = v80
                                        }
                                    case -230:
                                        v81 := GenS10{F0: 0, F1: 1000, F2: -1}
                                        _ = v81
                                    }
                                    _ = fi63
                                }
                                fmt.Println(v31)
                                for fi82 := 0; fi82 < 11; fi82++ {
                                    v3 = 1
                                    v1[2] = true
                                    _ = fi82
                                }
                            }
                        }
                    }
                    for fi83 := 0; fi83 < 2; fi83++ {
                        fmt.Println(fi83, v31, v3)
                        _ = fi83
                    }
                }
                if false {
                    var v84 string = "hello"
                    var v85 uint32 = 0
                    _ = v84
                    _ = v85
                }
                _ = fi41
            }
        case -331:
            v86 := GenS11{F0: 2147483648}
            _ = v86
        default:
            for fi87 := 0; fi87 < 14; fi87++ {
                v4 -= (v4 ^ 796)
                for fi88 := 0; fi88 < 13; fi88++ {
                    var v89 uint8 = (v3 / ((42) | 1))
                    _ = fi88
                    _ = v89
                }
                if false {
                    v90 := []int{32767}
                    var v91 uint32 = 882
                    switch (v4 << ((4294967295) & 63)) {
                    case -9223372036854775808:
                        v3 += (v31 & 2)
                    case 739:
                        for fi92 := 0; fi92 < 14; fi92++ {
                            v91 = (v91 | (^(v91)))
                            var v93 uint64 = 1
                            var v94 bool = false
                            var v95 int8 = 0
                            switch ("" + "hello") {
                            case "a":
                                v90 = append(v90, (fi87 - (v4 ^ int(v95))))
                            default:
                                if (32767 < 65535) {
                                    v4 = (fi92 * 521)
                                    var v96 int16 = 1000
                                    fmt.Println(v3)
                                    v97 := GenS1{F0: uint32(v93), F1: -9223372036854775808}
                                    _ = v96
                                    _ = v97
                                } else {
                                    var v98 int8 = 0
                                    fmt.Println(v94)
                                    _ = v98
                                }
                            }
                            if v94 {
                                var v99 float32 = 0.0
                                var v100 uint32 = 266
                                _ = v100
                                _ = v99
                            }
                            _ = fi92
                            _ = v93
                            _ = v94
                            _ = v95
                        }
                    }
                    _ = v90
                    _ = v91
                } else {
                    if false {
                        if true {
                            v4 += v4
                            if false {
                                v31 = 42
                                var v101 float32 = 0.1
                                if ((fi87 % ((fi87) | 1)) != fi87) {
                                    var v102 uint32 = 2147483648
                                    v31 = uint8(v102)
                                    v103 := GenS8{F0: 1000}
                                    _ = v102
                                    _ = v103
                                }
                                _ = v101
                            } else {
                                var v104 uint32 = 4294967295
                                v104 += (v104 * (v104 - v104))
                                var v105 int64 = 0
                                _ = v104
                                _ = v105
                            }
                            var v106 int = -9223372036854775808
                            _ = v106
                        } else {
                            v4 += v4
                        }
                        if true {
                            fmt.Println(fi87, v3)
                        } else {
                            if false {
                                for fi107 := 0; fi107 < 4; fi107++ {
                                    v108 := []int{fi107, -2147483648}
                                    v109 := GenS12{F0: int8(fi87), F1: 492}
                                    _ = fi107
                                    _ = v108
                                    _ = v109
                                }
                                var v110 bool = (false || true)
                                var v111 int16 = 1000
                                _ = v110
                                _ = v111
                            }
                            if false {
                                for fi112 := 0; fi112 < 18; fi112++ {
                                    var v113 bool = (false && false)
                                    v114 := []int{fi112}
                                    v115 := GenS11{F0: 283}
                                    var v116 string = ("hello" + "")
                                    var v117 uint64 = 18446744073709551615
                                    _ = fi112
                                    _ = v113
                                    _ = v114
                                    _ = v115
                                    _ = v116
                                    _ = v117
                                }
                            }
                        }
                        v3 -= v3
                    } else {
                        switch (^(fi87)) {
                        case 9223372036854775807:
                            v4 -= -392
                        case 879:
                            var v118 bool = true
                            _ = v118
                        case 32767:
                            fmt.Println(fi87, v31, v3)
                        }
                    }
                    if true {
                        v119 := []int{(v4 | (fi87 & -252)), 32768}
                        _ = v119
                    }
                }
                _ = fi87
            }
        }
        v120 := GenS9{F0: 9223372036854775807, F1: 65535}
        _ = v120
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v4
}
