package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 string; F2 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 int16; F2 int16 }
type GenS6 struct { F0 int16; F1 string; F2 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int16; F2 int }
type GenS9 struct { F0 uint; F1 int64 }
type GenS10 struct { F0 int8; F1 string; F2 int }
type GenS11 struct { F0 uint; F1 int }
type GenS12 struct { F0 uint32; F1 uint8; F2 string }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint64; F1 int16; F2 bool }
type GenS15 struct { F0 int32; F1 int64; F2 int16 }
func genFn0(p0 uint, p1 uint) uint16 {
    v4 := []int{-965, 32768, -144}
    for fi5 := 0; fi5 < 6; fi5++ {
        switch "hello" {
        case "":
            v6 := GenS1{F0: 42}
            _ = v6
        case "\n\t":
            v7 := GenS2{F0: (p0 | p1)}
            _ = v7
        default:
            var v8 int8 = 127
            _ = v8
        }
        var v9 int64 = 491
        switch (fi5 * 1) {
        case 4294967295:
            if false {
                for fi10 := 0; fi10 < 8; fi10++ {
                    _ = len(v4)
                    var v11 uint8 = 127
                    _ = fi10
                    _ = v11
                }
                v12 := GenS2{F0: p0}
                for fi13 := 0; fi13 < 19; fi13++ {
                    if ((483 >= 2) || true) {
                        var v14 int64 = v9
                        _ = v14
                    } else {
                        v9 = (^(v9))
                    }
                    _ = fi13
                }
                _ = v12
            }
        case -623:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 305:
            v4 = append(v4, 2147483647)
        default:
            fmt.Println(v9)
        }
        _ = fi5
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    return 829
}
func genFn1(p0 int64) uint8 {
    for fi15 := 0; fi15 < 6; fi15++ {
        if (!(!false)) {
            v16 := []int{-1, fi15, fi15, (fi15 ^ (fi15 / ((-2147483648) | 1)))}
            var v17 uint8 = 0
            switch (fi15 ^ -1) {
            case 0:
                var v18 float64 = float64(float32(0.0))
                _ = v18
            default:
                var v19 uint32 = 2147483647
                _ = v19
            }
            _ = v16
            _ = v17
        } else {
            v20 := GenS1{F0: 523}
            _ = v20
        }
        _ = fi15
    }
    var v21 uint8 = 59
    switch "a" {
    case "":
        switch "" {
        case "\n\t":
            switch 2147483647 {
            case 2147483647:
                v22 := []int{-106, 127, 4294967295}
                _ = v22
            case -868:
                var v23 uint = 4294967295
                _ = v23
            case -744:
                v24 := GenS3{F0: 0, F1: "", F2: true}
                _ = v24
            }
        }
    case "foo bar":
        v25 := GenS4{F0: 2147483647}
        _ = v25
    case "hello":
        v26 := GenS5{F0: -2, F1: 0, F2: 436}
        _ = v26
    }
    v27 := GenS2{F0: 4294967295}
    var v28 int64 = (p0 << ((255) & 63))
    _ = p0
    _ = v21
    _ = v27
    _ = v28
    return 42
}

func main() {
    var v1 int = 128
    var v2 uint16 = 491
    v3 := make(map[string]bool)
    v29 := GenS6{F0: 32767, F1: "foo bar", F2: (!false)}
    var v30 bool = false
    var v31 bool = (false && (v30 || false))
    for fi32 := 0; fi32 < 3; fi32++ {
        if v31 {
            v33 := []int{(fi32 & (v1 + int(v2))), (v1 >> ((707) & 63))}
            _ = v33
        } else {
            if (((0.0 / 1.0) == -2.25) || (42 == -69)) {
                switch "hello" {
                case "\n\t":
                    if (3.14159 == 1.5) {
                        var v34 int64 = 42
                        for fi35 := 0; fi35 < 15; fi35++ {
                            var v36 uint16 = (v2 | (v2 & genFn0(2, 4294967295)))
                            switch (("hello" + ((("\n\t" + "a") + "\n\t") + "foo bar")) + "hello") {
                            case "\n\t":
                                if (fi35 == fi35) {
                                    switch fi35 {
                                    case -581:
                                        if true {
                                            v2 += genFn0(0, 4294967295)
                                            v37 := GenS4{F0: 65535}
                                            _ = v37
                                        } else {
                                            var v38 uint = 42
                                            _ = v38
                                        }
                                    case 4294967295:
                                        if v30 {
                                            var v39 bool = ((false || ((false && (!true)) && (0.0 > 1e10))) || v30)
                                            var v40 uint32 = 2147483648
                                            _ = v39
                                            _ = v40
                                        }
                                    default:
                                        fmt.Println(v36)
                                    }
                                    v1 -= 2147483647
                                    v41 := GenS1{F0: 2}
                                    v42 := []int{(fi35 * v1), (fi35 << ((727) & 63)), (fi35 * (v1 >> ((42) & 63))), (-(fi35))}
                                    switch 0 {
                                    case -2147483648:
                                        var v43 string = "a"
                                        _ = v43
                                    case -85:
                                        switch "" {
                                        case "\n\t":
                                            var v44 bool = (("hello" + "hello") < "")
                                            _ = v44
                                        case "":
                                            v45 := GenS7{F0: v34}
                                            _ = v45
                                        case "a":
                                            v46 := GenS7{F0: -9223372036854775808}
                                            _ = v46
                                        }
                                    case 4294967295:
                                        var v47 int64 = -22
                                        _ = v47
                                    default:
                                        var v48 int16 = 1
                                        _ = v48
                                    }
                                    if ((!v30) || v31) {
                                        var v49 bool = (!v31)
                                        _ = v49
                                    } else {
                                        v2 -= v2
                                        var v50 uint8 = 128
                                        _ = v50
                                    }
                                    _ = len(v3)
                                    _ = v41
                                    _ = v42
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        if true {
                                            switch "\n\t" {
                                            case "foo bar":
                                                if ((true || v30) && (-93 >= 1)) {
                                                    v34 += v34
                                                    v2 = v36
                                                    fmt.Println(fi35)
                                                }
                                            case "hello":
                                                v36 = (v2 | v36)
                                            case "a":
                                                v51 := []int{}
                                                _ = v51
                                            default:
                                                if v30 {
                                                    var v52 uint32 = 2147483648
                                                    v34 -= (v34 - v34)
                                                    fmt.Println(fi32, v36)
                                                    _ = v52
                                                } else {
                                                    v31 = v31
                                                    var v53 int64 = (^(v34))
                                                    _ = v53
                                                }
                                            }
                                            fmt.Println(v2, fi32, v36)
                                        } else {
                                            switch "\n\t" {
                                            case "\n\t":
                                                if true {
                                                    v34 = (v34 | (v34 - 301))
                                                    delete(v3, "foo bar")
                                                } else {
                                                    var v54 uint32 = 2147483648
                                                    var v55 float64 = -2.25
                                                    if v31 {
                                                        var v56 string = ("a" + "a")
                                                        _ = v56
                                                    } else {
                                                        var v57 uint64 = 0
                                                        _ = v57
                                                    }
                                                    _ = v54
                                                    _ = v55
                                                }
                                            case "a":
                                                v31 = (true && v31)
                                            case "hello":
                                                var v58 uint64 = 342
                                                _ = v58
                                            }
                                            v59 := GenS5{F0: -38, F1: 42, F2: 32767}
                                            _ = v59
                                        }
                                    case "a":
                                        v30 = v30
                                    default:
                                        var v60 int8 = 103
                                        _ = v60
                                    }
                                    v30 = (!(473 > 701))
                                    var v61 int8 = 1
                                    _ = v61
                                }
                            case "":
                                switch -817 {
                                case -9223372036854775808:
                                    v62 := GenS8{F0: 10, F1: 32767, F2: fi35}
                                    _ = v62
                                case 322:
                                    var v63 uint = 584
                                    _ = v63
                                case 1:
                                    var v64 float64 = -2.25
                                    _ = v64
                                }
                            default:
                                delete(v3, ("" + "\n\t"))
                            }
                            _ = fi35
                            _ = v36
                        }
                        if (520 < (-(v34))) {
                            fmt.Println(v1)
                            v31 = v31
                            v65 := GenS9{F0: 65535, F1: (v34 & 87)}
                            _ = v65
                        } else {
                            v66 := GenS10{F0: -5, F1: "\n\t", F2: -109}
                            v31 = v31
                            _ = v66
                        }
                        _ = v34
                    } else {
                        switch "a" {
                        case "foo bar":
                            for fi67 := 0; fi67 < 7; fi67++ {
                                if false {
                                    v68 := []int{fi32, v1}
                                    if false {
                                        if true {
                                            switch "" {
                                            case "a":
                                                switch (v1 * (v1 ^ (fi32 << ((2) & 63)))) {
                                                case 909:
                                                    switch fi32 {
                                                    case 32768:
                                                        v31 = ((17 >= 0) || (!v30))
                                                    case -800:
                                                        if (690 == 1) {
                                                            v69 := GenS7{F0: -99}
                                                            _ = v69
                                                        } else {
                                                            v30 = true
                                                        }
                                                    default:
                                                        var v70 int64 = 865
                                                        _ = v70
                                                    }
                                                case 9223372036854775807:
                                                    v71 := GenS9{F0: 0, F1: 42}
                                                    _ = v71
                                                default:
                                                    v72 := GenS8{F0: 2, F1: -889, F2: (v1 << ((439) & 63))}
                                                    _ = v72
                                                }
                                            case "\n\t":
                                                var v73 string = ("hello" + ("a" + "hello"))
                                                _ = v73
                                            }
                                            if v30 {
                                                var v74 uint = 303
                                                _ = v74
                                            } else {
                                                v75 := GenS1{F0: 0}
                                                var v76 string = "a"
                                                fmt.Println(v1, fi67)
                                                _ = v75
                                                _ = v76
                                            }
                                            v31 = (v30 || (!(95 > -1)))
                                            var v77 int32 = 42
                                            var v78 uint8 = 42
                                            switch v1 {
                                            case 32767:
                                                if false {
                                                    var v79 bool = v30
                                                    v2 = 131
                                                    v80 := GenS5{F0: 101, F1: -201, F2: 32767}
                                                    _ = v79
                                                    _ = v80
                                                }
                                            case -9223372036854775808:
                                                switch fi67 {
                                                case 616:
                                                    var v81 uint64 = 0
                                                    _ = v81
                                                }
                                            case -207:
                                                _ = len(v3)
                                            }
                                            _ = v77
                                            _ = v78
                                        }
                                    }
                                    _ = v68
                                } else {
                                    switch (fi32 / ((fi67) | 1)) {
                                    case 69:
                                        var v82 float64 = -2.25
                                        _ = v82
                                    case 2147483647:
                                        v1 = (v1 + v1)
                                    case 0:
                                        v30 = (v31 || (v31 && (4294967295 == 42)))
                                    }
                                    var v83 int32 = -1
                                    if v30 {
                                        var v84 float32 = (3.14159 - 1e10)
                                        v2 = v2
                                        if v31 {
                                            var v85 int32 = (v83 - 2147483647)
                                            var v86 int8 = 0
                                            v87 := GenS11{F0: 175, F1: 4294967295}
                                            fmt.Println(fi67, v83)
                                            var v88 bool = true
                                            _ = v85
                                            _ = v86
                                            _ = v87
                                            _ = v88
                                        } else {
                                            v89 := GenS12{F0: 42, F1: 0, F2: ""}
                                            _ = v89
                                        }
                                        _ = v84
                                    } else {
                                        var v90 string = "foo bar"
                                        v91 := []int{fi32}
                                        if true {
                                            fmt.Println(fi67, v90)
                                        } else {
                                            var v92 float32 = (((float32(((-2.25 - (-2.25 * (0.1 + float64(float32((0.0 + (0.1 + 3.14159))))))) / 3.14159)) * float32(0.1)) / 1.5) - 0.0)
                                            if v31 {
                                                v3[v90] = true
                                            }
                                            _ = v92
                                        }
                                        v1 -= (fi67 | (v1 & v1))
                                        v30 = (1 == 0)
                                        _ = v90
                                        _ = v91
                                    }
                                    v93 := GenS11{F0: 255, F1: 4294967295}
                                    _ = v83
                                    _ = v93
                                }
                                v30 = false
                                _ = fi67
                            }
                        case "a":
                            if (924 > 42) {
                                v1 = (fi32 | (v1 | 868))
                            }
                        case "":
                            delete(v3, "a")
                        default:
                            for fi94 := 0; fi94 < 4; fi94++ {
                                v95 := GenS13{F0: ("" + "foo bar")}
                                var v96 int8 = 114
                                var v97 float64 = (0.0 - 1.5)
                                _ = fi94
                                _ = v95
                                _ = v96
                                _ = v97
                            }
                        }
                        var v98 int16 = 766
                        _ = v98
                    }
                case "a":
                    v1 = v1
                case "foo bar":
                    var v99 int8 = 1
                    _ = v99
                }
                var v100 uint64 = 42
                _ = v100
            } else {
                v30 = (255 <= 65535)
                for fi101 := 0; fi101 < 20; fi101++ {
                    v31 = ((!((0.1 + float32(float64(((3.14159 / ((float32(0.1) + -2.25) + 3.14159)) - (1e10 * (((3.14159 - 0.1) / float32(1e10)) / (float32(1.5) / 0.1))))))) != (0.1 / 1.0))) && (42 < -257))
                    var v102 uint8 = 123
                    var v103 uint16 = genFn0(729, 42)
                    _ = fi101
                    _ = v102
                    _ = v103
                }
                var v104 int32 = -2147483648
                _ = v104
            }
        }
        _ = fi32
    }
    v105 := []int{v1, v1, v1}
    switch v1 {
    case -621:
        if true {
            var v106 string = ("" + "a")
            v30 = (18446744073709551615 <= 42)
            _ = v106
        }
    case 32768:
        v2 = (v2 | 32767)
    case 32767:
        if v30 {
            v31 = v30
            v1 = 883
        }
    }
    var v107 int8 = 100
    switch ("hello" + "\n\t") {
    case "":
        v30 = v30
    case "hello":
        v2 = (v2 * 42)
    case "a":
        var v108 bool = (v30 && (2147483647 <= 184))
        _ = v108
    default:
        if len(v105) >= 2 {
            v105 = v105[1:len(v105)]
        }
    }
    if v31 {
        var v109 bool = (!false)
        if (216 >= 148) {
            if false {
                _ = len(v105)
                switch 9223372036854775807 {
                case 9223372036854775807:
                    if (((v1 & (v1 & v1)) < v1) || (!true)) {
                        switch 32767 {
                        case -1:
                            if v109 {
                                var v110 float64 = 1.5
                                var v111 float64 = v110
                                fmt.Println(v2, v31)
                                _ = v110
                                _ = v111
                            }
                        case 4294967295:
                            for fi112 := 0; fi112 < 18; fi112++ {
                                fmt.Println(v2)
                                _ = fi112
                            }
                        default:
                            v113 := GenS3{F0: 1, F1: "foo bar", F2: (v109 && v30)}
                            _ = v113
                        }
                        v105 = append(v105, v1)
                    } else {
                        v109 = false
                        var v114 int8 = -98
                        _ = v114
                    }
                case 4294967295:
                    v107 = v107
                case 1:
                    v115 := GenS3{F0: (v2 % ((genFn0(42, 65535)) | 1)), F1: "a", F2: (v109 || v30)}
                    _ = v115
                }
                v107 = (v107 * (v107 - 58))
            }
        }
        v105 = append(v105, v1)
        _ = v109
    } else {
        var v116 uint16 = (v2 & v2)
        v3[""] = v30
        _ = v116
    }
    switch 0 {
    case 128:
        fmt.Println(v31)
    case -55:
        v1 = v1
    case -1:
        var v117 uint32 = 2147483648
        _ = v117
    }
    v118 := GenS14{F0: 1, F1: 0, F2: (!v30)}
    v119 := GenS15{F0: -1, F1: 42, F2: 1}
    for fi120 := 0; fi120 < 3; fi120++ {
        if v30 {
            var v121 bool = (1 >= 0)
            v107 = 81
            v3["foo bar"] = v121
            _ = v121
        }
        fmt.Println(v1, v107, v107)
        _ = fi120
    }
    switch "a" {
    case "a":
        v1 -= v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v105:", len(v105), v105)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v105
    _ = v107
    _ = v118
    _ = v119
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
