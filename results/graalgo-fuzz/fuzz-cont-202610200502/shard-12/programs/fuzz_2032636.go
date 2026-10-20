package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int; F2 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 uint; F2 int32 }
type GenS7 struct { F0 uint; F1 int16; F2 int32 }
type GenS8 struct { F0 int8; F1 uint; F2 string }
type GenS9 struct { F0 string; F1 uint32; F2 string }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 string; F1 uint32; F2 uint32 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int16; F1 bool; F2 int64 }
type GenS14 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int32 }
func genFn0(p0 int16) bool {
    v6 := []int{1}
    v7 := GenS1{F0: 1, F1: "hello"}
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 42, F1: ""}
        _ = fi8
        _ = v9
    }
    var v10 int = 32767
    v6 = append(v6, v10)
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    return ((true && false) && false)
}

func main() {
    var v1 uint8 = 221
    var v2 uint8 = v1
    v3 := make(map[string]int)
    var v4 int16 = -32768
    var v5 uint8 = v1
    var v11 uint8 = 42
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 int16 = (v4 / (((v4 >> ((279) & 63))) | 1))
        _ = fi12
        _ = v13
    }
    var v14 uint8 = 127
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS1{F0: 0, F1: ("a" + "hello")}
        _ = fi15
        _ = v16
    }
    v14 += 127
    v3["hello"] = -177
    v4 = (v4 << ((1) & 63))
    var v17 uint = 0
    for fi18 := 0; fi18 < 19; fi18++ {
        switch ((("a" + "\n\t") + "foo bar") + ((("\n\t" + "foo bar") + "hello") + ("foo bar" + "foo bar"))) {
        case "foo bar":
            if false {
                v17 += v17
                switch ("" + "hello") {
                case "hello":
                    v19 := GenS1{F0: v2, F1: "a"}
                    _ = v19
                }
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 int16 = v4
                    v22 := GenS1{F0: (v11 / (((^(v5))) | 1)), F1: "hello"}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                if ((v11 >> ((v17) & 63)) == uint8(v17)) {
                    v3["hello"] = fi18
                }
            }
        case "":
            v5 -= 53
        case "hello":
            v2 -= (v1 / ((v2) | 1))
        }
        var v23 uint8 = v14
        if true {
            var v24 string = ("hello" + ("hello" + ""))
            _ = v24
        } else {
            var v25 uint32 = 2147483648
            var v26 uint64 = 0
            _ = v25
            _ = v26
        }
        _ = fi18
        _ = v23
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        for fi28 := 0; fi28 < 17; fi28++ {
            switch "a" {
            case "a":
                switch fi27 {
                case 607:
                    var v29 uint = 294
                    _ = v29
                case 127:
                    fmt.Println(v1)
                case 32767:
                    if true {
                        v30 := GenS2{F0: 127}
                        var v31 bool = (42 < 221)
                        if ((v1 == 2) || false) {
                            fmt.Println(v5)
                            v32 := []int{}
                            _ = v32
                        } else {
                            v33 := GenS3{F0: 42, F1: fi28, F2: v4}
                            v34 := []int{fi27, (fi28 | fi27), 4294967295}
                            v35 := GenS4{F0: 52}
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        fmt.Println(v4, v17)
                        _ = v30
                        _ = v31
                    } else {
                        var v36 float32 = 0.0
                        switch "foo bar" {
                        case "hello":
                            v37 := GenS3{F0: 42, F1: (fi28 / ((fi28) | 1)), F2: (v4 + (^(v4)))}
                            _ = v37
                        case "foo bar":
                            if (v4 != v4) {
                                if false {
                                    v4 = (v4 << (((^(v17))) & 63))
                                    if true {
                                        v14 -= 42
                                        if true {
                                            fmt.Println(v36)
                                            v14 -= 21
                                            v38 := GenS5{F0: 447}
                                            _ = v38
                                        }
                                    } else {
                                        v17 = v17
                                        var v39 int = (fi28 << (((v17 * v17)) & 63))
                                        if false {
                                            v3[("\n\t" + "foo bar")] = -543
                                            var v40 uint8 = (v14 - (v5 - (v14 % ((v1) | 1))))
                                            _ = v40
                                        } else {
                                            var v41 int16 = -32768
                                            v4 = 131
                                            switch -347 {
                                            case 977:
                                                fmt.Println(v17, v11)
                                            default:
                                                v42 := GenS6{F0: 2, F1: v17, F2: -1}
                                                _ = v42
                                            }
                                            _ = v41
                                        }
                                        var v43 uint64 = 2
                                        _ = v39
                                        _ = v43
                                    }
                                    v36 = float32(0.0)
                                }
                            }
                        case "\n\t":
                            var v44 bool = true
                            _ = v44
                        }
                        if false {
                            var v45 uint16 = 161
                            var v46 float32 = (0.0 * v36)
                            _ = v45
                            _ = v46
                        } else {
                            fmt.Println(v4, fi27, v36)
                            var v47 string = ""
                            v48 := GenS7{F0: (v17 * v17), F1: (v4 * (v4 % ((-32768) | 1))), F2: 1}
                            switch ("foo bar" + ("foo bar" + v47)) {
                            case "":
                                v49 := v3[(v47 + "a")]
                                _ = v49
                            case "\n\t":
                                var v50 uint16 = 1
                                switch fi28 {
                                case 32767:
                                    switch fi27 {
                                    case 127:
                                        if true {
                                            var v51 float64 = (((1.5 / 1.5) + 1.5) + (-2.25 - 3.14159))
                                            switch (fi27 + 127) {
                                            case 42:
                                                v52 := []int{128, (fi28 % (((fi27 - (fi28 << (((v17 + 42)) & 63)))) | 1)), fi28, fi27}
                                                _ = v52
                                            case 856:
                                                switch fi27 {
                                                case -1:
                                                    switch fi28 {
                                                    case -9223372036854775808:
                                                        fmt.Println(v5, v17)
                                                    }
                                                    v53 := GenS8{F0: 127, F1: 255, F2: ((v47 + v47) + "a")}
                                                    _ = v53
                                                case 42:
                                                    if (v36 >= v36) {
                                                        v54 := []int{(fi28 << ((v17) & 63)), (fi28 & fi27), -2147483648}
                                                        v4 = -1
                                                        var v55 uint32 = 1
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                    fmt.Println(fi27)
                                                    if (((fi28 / (((fi28 - 4294967295)) | 1)) == -1) && false) {
                                                        if true {
                                                            v56 := GenS2{F0: 0}
                                                            if (true && (v4 < (v4 * v4))) {
                                                                if (!true) {
                                                                    switch 42 {
                                                                    case 0:
                                                                        if true {
                                                                            if (((!true) || (!false)) && false) {
                                                                                if (536 != (^(v4))) {
                                                                                    switch 861 {
                                                                                    case 80:
                                                                                        var v57 uint16 = v50
                                                                                        v51 = v51
                                                                                        if ((1 == -827) && (834 < 251)) {
                                                                                            _ = len(v3)
                                                                                            v14 += (v1 + 33)
                                                                                            var v58 int = (fi27 * fi27)
                                                                                            var v59 int32 = -669
                                                                                            _ = v58
                                                                                            _ = v59
                                                                                        } else {
                                                                                            v60 := []int{(fi27 | fi27), 1, (fi28 << (((v17 * 2)) & 63)), -2147483648}
                                                                                            v36 = ((v36 * v36) * -2.25)
                                                                                            var v61 uint = 1
                                                                                            _ = v60
                                                                                            _ = v61
                                                                                        }
                                                                                        _ = v57
                                                                                    case -576:
                                                                                        v62 := GenS8{F0: 100, F1: v17, F2: (v47 + "hello")}
                                                                                        _ = v62
                                                                                    }
                                                                                } else {
                                                                                    if false {
                                                                                        v11 += v11
                                                                                        if (1 == 61) {
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v5, v51, v4)
                                                                                            case "a":
                                                                                                v4 += (v4 ^ (v4 + (-(v4))))
                                                                                            }
                                                                                        }
                                                                                        v50 = v50
                                                                                    }
                                                                                    if (!false) {
                                                                                        fmt.Println(v5, v36)
                                                                                        switch 2147483647 {
                                                                                        case 128:
                                                                                            delete(v3, v47)
                                                                                        case 32768:
                                                                                            var v63 uint = 18446744073709551615
                                                                                            _ = v63
                                                                                        case 4294967295:
                                                                                            _ = len(v3)
                                                                                        default:
                                                                                            v64 := GenS4{F0: (^(v50))}
                                                                                            _ = v64
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v50, v1, v17)
                                                                                        fmt.Println(v2, v36)
                                                                                        v65 := GenS7{F0: 2, F1: v4, F2: 1}
                                                                                        if (false || (false || ((v51 - v51) >= (3.14159 - 1e10)))) {
                                                                                            var v66 float64 = -2.25
                                                                                            _ = v66
                                                                                        }
                                                                                        v67 := []int{9223372036854775807}
                                                                                        var v68 string = "hello"
                                                                                        _ = v65
                                                                                        _ = v67
                                                                                        _ = v68
                                                                                    }
                                                                                }
                                                                                var v69 uint32 = 1
                                                                                _ = v69
                                                                            }
                                                                            v70 := GenS9{F0: v47, F1: 2, F2: "foo bar"}
                                                                            v71 := []int{-735, fi27, 128}
                                                                            v11 = (^(v2))
                                                                            _ = v70
                                                                            _ = v71
                                                                        } else {
                                                                            var v72 uint = v17
                                                                            var v73 int64 = -1
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                    }
                                                                    switch fi28 {
                                                                    case -81:
                                                                        if false {
                                                                            var v74 int = (fi27 / ((42) | 1))
                                                                            _ = v74
                                                                        } else {
                                                                            switch (fi27 | fi28) {
                                                                            case -1:
                                                                                v75 := GenS10{F0: 74}
                                                                                _ = v75
                                                                            case 4294967295:
                                                                                if (-1 < 0) {
                                                                                    v1 = 65
                                                                                } else {
                                                                                    v51 -= 0.1
                                                                                    v76 := GenS1{F0: v14, F1: v47}
                                                                                    var v77 uint16 = v50
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                }
                                                                            default:
                                                                                var v78 uint64 = 553
                                                                                _ = v78
                                                                            }
                                                                            switch (fi27 / (((^(fi28))) | 1)) {
                                                                            case 42:
                                                                                v79 := GenS8{F0: 42, F1: 2, F2: v47}
                                                                                _ = v79
                                                                            case -1:
                                                                                v80 := GenS11{F0: v47, F1: 4294967295, F2: 65535}
                                                                                _ = v80
                                                                            default:
                                                                                v36 = ((v36 + (v36 * 1.5)) * v36)
                                                                            }
                                                                            var v81 int16 = v4
                                                                            var v82 int16 = v81
                                                                            var v83 bool = (!((45 > 0) && (!false)))
                                                                            fmt.Println(fi28, v11, v51)
                                                                            var v84 bool = (0.0 > (float64(v36) * ((v51 / (1e10 * v51)) - (v51 / v51))))
                                                                            _ = v81
                                                                            _ = v82
                                                                            _ = v83
                                                                            _ = v84
                                                                        }
                                                                    case 48:
                                                                        var v85 int32 = int32(v14)
                                                                        _ = v85
                                                                    case 32768:
                                                                        var v86 int8 = -75
                                                                        _ = v86
                                                                    default:
                                                                        v87 := []int{}
                                                                        _ = v87
                                                                    }
                                                                    fmt.Println(v5, v51)
                                                                }
                                                                var v88 uint = v17
                                                                v89 := GenS12{F0: v88}
                                                                v47 = "a"
                                                                var v90 string = "foo bar"
                                                                _ = v88
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                            _ = v56
                                                        } else {
                                                            var v91 int64 = 0
                                                            _ = v91
                                                        }
                                                        if (!((!(0 == 45)) && ((true && ((int(v17) < (fi28 >> ((2) & 63))) || (423 != int32(v4)))) || true))) {
                                                            var v92 uint64 = 42
                                                            v17 = 582
                                                            _ = v92
                                                        } else {
                                                            if (9223372036854775807 != -1) {
                                                                switch fi28 {
                                                                case 128:
                                                                    var v93 uint16 = 32767
                                                                    _ = v93
                                                                case 4294967295:
                                                                    v1 += (v5 % (((v14 ^ (v11 & (v2 - 127)))) | 1))
                                                                case -2147483648:
                                                                    if (true && (((fi27 + (fi28 % ((9223372036854775807) | 1))) != 127) || false)) {
                                                                        v47 = v47
                                                                    } else {
                                                                        v50 = 42
                                                                        var v94 int64 = 0
                                                                        fmt.Println(fi28)
                                                                        if (-65 != -109) {
                                                                            var v95 uint64 = 2
                                                                            _ = v95
                                                                        } else {
                                                                            var v96 uint16 = uint16(v14)
                                                                            var v97 float64 = v51
                                                                            _ = v96
                                                                            _ = v97
                                                                        }
                                                                        _ = v94
                                                                    }
                                                                default:
                                                                    v2 = v11
                                                                    v36 = float32((0.0 + 0.1))
                                                                    fmt.Println(v4, fi27)
                                                                }
                                                                var v98 int32 = 277
                                                                fmt.Println(v14)
                                                                _ = v98
                                                            }
                                                            delete(v3, (v47 + v47))
                                                        }
                                                    } else {
                                                        v1 = v2
                                                        v99 := GenS7{F0: (v17 % (((v17 << (((v17 & uint(v14))) & 63))) | 1)), F1: (v4 | 32767), F2: -768}
                                                        var v100 uint32 = 4294967295
                                                        _ = v100
                                                        _ = v99
                                                    }
                                                default:
                                                    if ((3.14159 - v36) != 0.0) {
                                                        if false {
                                                            v101 := GenS7{F0: (v17 % (((v17 / (((v17 | v17)) | 1))) | 1)), F1: 425, F2: 191}
                                                            if false {
                                                                var v102 int32 = 1
                                                                var v103 bool = (((v14 | v14) > v14) && false)
                                                                _ = v102
                                                                _ = v103
                                                            }
                                                            fmt.Println(v47, fi27)
                                                            _ = v101
                                                        }
                                                    }
                                                }
                                            case -1:
                                                v14 += 3
                                            }
                                            switch v47 {
                                            case "a":
                                                switch v47 {
                                                case "\n\t":
                                                    var v104 string = v47
                                                    _ = v104
                                                }
                                            case "foo bar":
                                                v105 := GenS13{F0: v4, F1: true, F2: int64(v17)}
                                                _ = v105
                                            default:
                                                var v106 uint = (v17 | (v17 + 425))
                                                _ = v106
                                            }
                                            v107 := []int{fi27, fi28, fi27}
                                            _ = v107
                                            _ = v51
                                        }
                                    case -237:
                                        v11 += 128
                                    case 508:
                                        v36 += 0.0
                                    default:
                                        v108 := []int{(fi28 << ((18446744073709551615) & 63)), int(v4)}
                                        _ = v108
                                    }
                                default:
                                    var v109 bool = false
                                    _ = v109
                                }
                                _ = v50
                            }
                            v11 -= (v5 * v2)
                            var v110 uint8 = (v14 + 2)
                            _ = v110
                            _ = v47
                            _ = v48
                        }
                        _ = v36
                    }
                }
            default:
                var v111 float32 = 0.1
                var v112 int = 32768
                _ = v111
                _ = v112
            }
            var v113 uint8 = (^(v5))
            _ = fi28
            _ = v113
        }
        for fi114 := 0; fi114 < 16; fi114++ {
            v115 := GenS1{F0: (v14 - v14), F1: ("a" + ("" + "hello"))}
            _ = fi114
            _ = v115
        }
        for fi116 := 0; fi116 < 10; fi116++ {
            fmt.Println(v17, v11, fi27)
            _ = fi116
        }
        _ = fi27
    }
    for fi117 := 0; fi117 < 2; fi117++ {
        for fi118 := 0; fi118 < 11; fi118++ {
            delete(v3, "hello")
            if false {
                fmt.Println(v14)
                if true {
                    v119 := GenS14{F0: 394, F1: v17, F2: 42}
                    v11 = v5
                    _ = v119
                }
            } else {
                v120 := GenS9{F0: "foo bar", F1: 65535, F2: "hello"}
                var v121 string = "hello"
                v14 += (v1 | (v11 & (v1 | (v2 + v5))))
                _ = v120
                _ = v121
            }
            v1 -= v11
            var v122 uint32 = 779
            v14 -= (v5 - (v14 | (v1 / ((v1) | 1))))
            _ = fi118
            _ = v122
        }
        if true {
            var v123 int16 = 0
            v124 := []int{}
            v125 := GenS15{F0: v4}
            _ = v123
            _ = v124
            _ = v125
        }
        _ = fi117
    }
    for fi126 := 0; fi126 < 20; fi126++ {
        v127 := GenS16{F0: 2147483647}
        v14 = v11
        v17 = uint(v1)
        _ = fi126
        _ = v127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
