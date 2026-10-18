package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint16; F1 string; F2 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS8 struct { F0 uint; F1 uint16; F2 int64 }
type GenS9 struct { F0 uint32; F1 uint; F2 int32 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint; F1 string }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint; F1 int32 }
type GenS16 struct { F0 bool; F1 int8 }
type GenS17 struct { F0 int32; F1 int; F2 int32 }
func genFn0(p0 int16, p1 string, p2 int16) uint32 {
    v4 := GenS1{F0: 0}
    var v5 int = 42
    p1 = ((p1 + p1) + ("a" + "foo bar"))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 512
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 32767
    var v3 int64 = -9223372036854775808
    for fi6 := 0; fi6 < 15; fi6++ {
        v3 -= 9223372036854775807
        v7 := GenS1{F0: v3}
        if (537 == 65535) {
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 uint32 = 0
                if false {
                    switch 983 {
                    case -1:
                        var v10 string = "a"
                        _ = v10
                    case 2147483647:
                        if (1 < 18446744073709551615) {
                            if false {
                                fmt.Println(v3, v9, fi8)
                                v11 := GenS1{F0: 42}
                                v12 := []int{2147483647}
                                _ = v11
                                _ = v12
                            } else {
                                if false {
                                    switch (("\n\t" + "a") + "foo bar") {
                                    case "foo bar":
                                        var v13 int8 = 124
                                        _ = v13
                                    case "a":
                                        if true {
                                            fmt.Println(v9)
                                            if (4294967295 == 646) {
                                                v1[v9] = ((-1 == -51) || true)
                                                var v14 int = fi8
                                                v15 := []int{(v14 >> ((255) & 63)), (v14 | fi8), (fi6 | (fi8 / ((int(v3)) | 1))), fi8}
                                                v16 := GenS2{F0: (v2 ^ (v2 - int16(v9))), F1: (v9 * genFn0(v2, "hello", (v2 * (v2 / ((v2) | 1)))))}
                                                _ = v14
                                                _ = v15
                                                _ = v16
                                            }
                                            if false {
                                                var v17 int16 = 1
                                                _ = v17
                                            } else {
                                                v18 := GenS2{F0: v2, F1: genFn0(0, ("" + "hello"), (v2 >> ((2) & 63)))}
                                                _ = v18
                                            }
                                        }
                                    default:
                                        fmt.Println(v2, v2, fi6)
                                    }
                                    v3 += 1
                                }
                            }
                        } else {
                            v9 -= (v9 + v9)
                        }
                    case 32768:
                        switch ("hello" + "\n\t") {
                        case "\n\t":
                            if false {
                                switch (fi6 << ((898) & 63)) {
                                case 42:
                                    var v19 string = "\n\t"
                                    _ = v19
                                case -272:
                                    v20 := GenS3{F0: ("hello" + ("\n\t" + (("foo bar" + "a") + "")))}
                                    _ = v20
                                default:
                                    v2 = (v2 << ((255) & 63))
                                    v21 := []int{-1, (^(fi8)), 0, 866}
                                    _ = v21
                                }
                            }
                        case "a":
                            switch "" {
                            case "a":
                                var v22 float32 = -2.25
                                _ = v22
                            case "foo bar":
                                var v23 uint = 1
                                _ = v23
                            case "":
                                v1[(v9 ^ (v9 | 4294967295))] = false
                            }
                        case "foo bar":
                            v3 -= -701
                            var v24 int32 = 2147483647
                            v2 += v2
                            _ = v24
                        }
                    default:
                        var v25 int = (fi8 - -7)
                        var v26 uint32 = (^(v9))
                        v27 := GenS2{F0: v2, F1: (v9 * (v26 % (((v9 << ((42) & 63))) | 1)))}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS4{F0: 1, F1: 4294967295}
                    switch (fi8 << ((144) & 63)) {
                    case 202:
                        fmt.Println(v2, fi8)
                    case 128:
                        v9 = 2147483647
                    default:
                        if true {
                            var v29 bool = false
                            if false {
                                if (!true) {
                                    v3 -= (v3 / ((-9223372036854775808) | 1))
                                    var v30 uint16 = 32767
                                    if (false || (int32(fi6) != 2147483647)) {
                                        if true {
                                            v31 := GenS5{F0: (v30 - 1), F1: "a", F2: (v3 ^ (v3 >> ((4294967295) & 63)))}
                                            v3 += v3
                                            _ = v31
                                        } else {
                                            v32 := []int{}
                                            switch fi6 {
                                            case -500:
                                                var v33 int64 = (-(v3))
                                                _ = v33
                                            case -2147483648:
                                                var v34 uint8 = 114
                                                _ = v34
                                            default:
                                                if (float64(0.1) >= 0.0) {
                                                    v35 := GenS6{F0: 18446744073709551615}
                                                    v36 := GenS7{F0: 89, F1: (v3 / ((v3) | 1)), F2: 147}
                                                    if true {
                                                        v2 += 864
                                                        v37 := GenS8{F0: 277, F1: v30, F2: v3}
                                                        if (v29 || (2147483647 != 2147483647)) {
                                                            switch fi6 {
                                                            case 26:
                                                                _ = len(v32)
                                                            case 127:
                                                                switch "a" {
                                                                case "a":
                                                                    v38 := v1[v9]
                                                                    _ = v38
                                                                }
                                                            case 255:
                                                                v39 := GenS9{F0: 65535, F1: 74, F2: -1}
                                                                _ = v39
                                                            }
                                                            if (!false) {
                                                                if v29 {
                                                                    switch fi8 {
                                                                    case 127:
                                                                        v32 = append(v32, (fi6 % (((^(fi8))) | 1)))
                                                                    case 522:
                                                                        var v40 int64 = int64(v30)
                                                                        _ = v40
                                                                    }
                                                                    v41 := GenS10{F0: v9}
                                                                    if (38 != 18446744073709551615) {
                                                                        var v42 bool = v29
                                                                        v1[(v9 / (((v9 & uint32(v30))) | 1))] = ("" <= "")
                                                                        _ = v42
                                                                    } else {
                                                                        var v43 bool = (((fi8 >> ((1) & 63)) >= (fi8 | 127)) && (0 >= v2))
                                                                        if (genFn0(0, "foo bar", (v2 % ((v2) | 1))) != (v9 + v9)) {
                                                                            switch (fi8 ^ fi8) {
                                                                            case -1:
                                                                                var v44 int32 = -353
                                                                                _ = v44
                                                                            case 127:
                                                                                if len(v32) > 0 {
                                                                                    v32[0] = fi6
                                                                                }
                                                                            default:
                                                                                if (!v29) {
                                                                                    if (!v29) {
                                                                                        var v45 int = -323
                                                                                        fmt.Println(v3, fi6, v43)
                                                                                        var v46 uint64 = 1
                                                                                        delete(v1, v9)
                                                                                        if len(v32) >= 2 {
                                                                                            v32 = v32[1:len(v32)]
                                                                                        }
                                                                                        v47 := GenS1{F0: 0}
                                                                                        if false {
                                                                                            switch (fi8 / ((fi6) | 1)) {
                                                                                            case 1:
                                                                                                var v48 uint16 = v30
                                                                                                _ = v48
                                                                                            case 4294967295:
                                                                                                v2 += v2
                                                                                            }
                                                                                            var v49 int8 = 127
                                                                                            _ = v49
                                                                                        } else {
                                                                                            switch ("\n\t" + (("foo bar" + "hello") + "a")) {
                                                                                            case "":
                                                                                                if true {
                                                                                                    v32 = append(v32, (^(fi6)))
                                                                                                    var v50 uint32 = v9
                                                                                                    var v51 int = int(v2)
                                                                                                    if (84 <= 100) {
                                                                                                        v46 -= (v46 ^ v46)
                                                                                                        v52 := GenS7{F0: 2, F1: (^(v3)), F2: 72}
                                                                                                        var v53 int16 = 42
                                                                                                        _ = v52
                                                                                                        _ = v53
                                                                                                    } else {
                                                                                                        v46 = (v46 + (v46 >> ((21) & 63)))
                                                                                                        if true {
                                                                                                            var v54 int16 = (v2 + (v2 | -1))
                                                                                                            if (42 >= 114) {
                                                                                                                var v55 float32 = (3.14159 - 1.5)
                                                                                                                var v56 float32 = float32(0.1)
                                                                                                                _ = v55
                                                                                                                _ = v56
                                                                                                            }
                                                                                                            v51 = -695
                                                                                                            _ = v54
                                                                                                        } else {
                                                                                                            switch 32768 {
                                                                                                            case 4294967295:
                                                                                                                v57 := GenS11{F0: 4294967295, F1: v9, F2: 0}
                                                                                                                _ = v57
                                                                                                            case 9223372036854775807:
                                                                                                                if (501 > 1) {
                                                                                                                    var v58 int = v45
                                                                                                                    var v59 bool = false
                                                                                                                    var v60 string = "hello"
                                                                                                                    _ = v58
                                                                                                                    _ = v59
                                                                                                                    _ = v60
                                                                                                                }
                                                                                                            case 1:
                                                                                                                var v61 bool = v29
                                                                                                                _ = v61
                                                                                                            default:
                                                                                                                v43 = v43
                                                                                                            }
                                                                                                            fmt.Println(v2, v3)
                                                                                                        }
                                                                                                        var v62 int32 = 2147483647
                                                                                                        _ = v62
                                                                                                    }
                                                                                                    var v63 int16 = v2
                                                                                                    _ = v50
                                                                                                    _ = v51
                                                                                                    _ = v63
                                                                                                } else {
                                                                                                    v64 := GenS6{F0: v46}
                                                                                                    if ((((0.0 + 0.0) + 1.5) * 0.1) <= (3.14159 * 0.0)) {
                                                                                                        v65 := v1[v9]
                                                                                                        if false {
                                                                                                            switch (("" + "") + ("foo bar" + "\n\t")) {
                                                                                                            case "hello":
                                                                                                                if false {
                                                                                                                    v66 := GenS2{F0: (v2 + v2), F1: v9}
                                                                                                                    _ = v66
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                if v43 {
                                                                                                                    if len(v32) >= 2 {
                                                                                                                        v32 = v32[1:len(v32)]
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    if v65 {
                                                                                                                        v1[(v9 | (v9 | (v9 / ((2147483647) | 1))))] = (v3 == (v3 / ((42) | 1)))
                                                                                                                    } else {
                                                                                                                        switch "\n\t" {
                                                                                                                        case "a":
                                                                                                                            v29 = true
                                                                                                                        }
                                                                                                                        fmt.Println(v29, v45)
                                                                                                                    }
                                                                                                                    v43 = v29
                                                                                                                    if len(v32) >= 2 {
                                                                                                                        v32 = v32[1:len(v32)]
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                _ = len(v32)
                                                                                                            }
                                                                                                            switch "hello" {
                                                                                                            case "a":
                                                                                                                v29 = v29
                                                                                                            case "foo bar":
                                                                                                                switch int(v2) {
                                                                                                                case 127:
                                                                                                                    v9 -= 1
                                                                                                                case -61:
                                                                                                                    v65 = v65
                                                                                                                case 2147483647:
                                                                                                                    v67 := v1[930]
                                                                                                                    _ = v67
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                var v68 bool = true
                                                                                                                _ = v68
                                                                                                            }
                                                                                                            v69 := GenS3{F0: "\n\t"}
                                                                                                            v3 = -9223372036854775808
                                                                                                            _ = v69
                                                                                                        }
                                                                                                        _ = v65
                                                                                                    } else {
                                                                                                        _ = len(v32)
                                                                                                    }
                                                                                                    _ = v64
                                                                                                }
                                                                                            default:
                                                                                                switch (^(fi6)) {
                                                                                                case -9223372036854775808:
                                                                                                    var v70 int64 = v3
                                                                                                    if ((v30 * 188) >= v30) {
                                                                                                        if len(v32) >= 2 {
                                                                                                            v32 = v32[1:len(v32)]
                                                                                                        }
                                                                                                    } else {
                                                                                                        if (1 < 255) {
                                                                                                            v9 += (v9 - v9)
                                                                                                            v46 = v46
                                                                                                        } else {
                                                                                                            v46 += v46
                                                                                                        }
                                                                                                        fmt.Println(v9, v46, fi6)
                                                                                                    }
                                                                                                    switch (("a" + ("\n\t" + "hello")) + "hello") {
                                                                                                    case "\n\t":
                                                                                                        v70 -= v3
                                                                                                    case "foo bar":
                                                                                                        var v71 int16 = v2
                                                                                                        if len(v32) > 0 {
                                                                                                            v32[0] = 32767
                                                                                                        }
                                                                                                        var v72 int16 = v71
                                                                                                        _ = v71
                                                                                                        _ = v72
                                                                                                    case "hello":
                                                                                                        var v73 int64 = v3
                                                                                                        _ = v73
                                                                                                    }
                                                                                                    _ = v70
                                                                                                default:
                                                                                                    v29 = false
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v45
                                                                                        _ = v46
                                                                                        _ = v47
                                                                                    }
                                                                                } else {
                                                                                    v2 += (v2 << ((108) & 63))
                                                                                    v3 = int64(v2)
                                                                                }
                                                                            }
                                                                        }
                                                                        v32 = append(v32, 32768)
                                                                        _ = v43
                                                                    }
                                                                    _ = v41
                                                                } else {
                                                                    switch fi8 {
                                                                    case -9223372036854775808:
                                                                        if (-1 >= 42) {
                                                                            fmt.Println(v2, v29)
                                                                        }
                                                                    case 127:
                                                                        var v74 uint32 = (^(v9))
                                                                        _ = v74
                                                                    case -1:
                                                                        if v29 {
                                                                            v75 := GenS5{F0: 1, F1: "hello", F2: v3}
                                                                            _ = v75
                                                                        }
                                                                    default:
                                                                        var v76 uint32 = genFn0(0, ((("foo bar" + "\n\t") + "hello") + "a"), v2)
                                                                        _ = v76
                                                                    }
                                                                    var v77 int64 = (v3 % ((v3) | 1))
                                                                    _ = v77
                                                                }
                                                            } else {
                                                                v78 := GenS12{F0: 26}
                                                                switch fi6 {
                                                                case -687:
                                                                    v79 := GenS11{F0: 65535, F1: v9, F2: (v30 & v30)}
                                                                    _ = v79
                                                                case 4294967295:
                                                                    fmt.Println(v30, v9, v9)
                                                                case 868:
                                                                    _ = len(v32)
                                                                default:
                                                                    switch ("\n\t" + "\n\t") {
                                                                    case "hello":
                                                                        var v80 uint64 = 42
                                                                        _ = v80
                                                                    default:
                                                                        var v81 string = ""
                                                                        _ = v81
                                                                    }
                                                                    v82 := GenS12{F0: 100}
                                                                    _ = v82
                                                                }
                                                                _ = v78
                                                            }
                                                            v83 := GenS2{F0: v2, F1: (v9 % (((v9 * v9)) | 1))}
                                                            _ = v83
                                                        }
                                                        _ = v37
                                                    } else {
                                                        fmt.Println(v3)
                                                        if false {
                                                            if true {
                                                                v29 = true
                                                                v84 := GenS13{F0: 328, F1: "\n\t"}
                                                                _ = v84
                                                            }
                                                        } else {
                                                            v85 := GenS14{F0: 508}
                                                            delete(v1, v9)
                                                            if (v29 || (1 == -128)) {
                                                                var v86 uint8 = 247
                                                                if (v29 || (127 != v86)) {
                                                                    v3 -= (v3 | v3)
                                                                    if (18446744073709551615 <= 65535) {
                                                                        v87 := v1[(^(v9))]
                                                                        _ = v87
                                                                    } else {
                                                                        if v29 {
                                                                            fmt.Println(v29, v86, v30)
                                                                            v88 := GenS12{F0: -4}
                                                                            v86 = (v86 + 8)
                                                                            _ = v88
                                                                        }
                                                                        if len(v32) >= 2 {
                                                                            v32 = v32[1:len(v32)]
                                                                        }
                                                                    }
                                                                    v29 = true
                                                                } else {
                                                                    var v89 int64 = v3
                                                                    if (-37 <= -128) {
                                                                        v1[v9] = (!v29)
                                                                        var v90 bool = (694 != 762)
                                                                        if v29 {
                                                                            var v91 string = ("a" + "")
                                                                            _ = v91
                                                                        }
                                                                        _ = v90
                                                                    }
                                                                    fmt.Println(v29, v3)
                                                                    _ = v89
                                                                }
                                                                v29 = (v9 != v9)
                                                                _ = v86
                                                            }
                                                            if true {
                                                                fmt.Println(v9, v9, v29)
                                                                if true {
                                                                    var v92 int8 = 0
                                                                    var v93 uint8 = 127
                                                                    _ = v92
                                                                    _ = v93
                                                                } else {
                                                                    _ = len(v1)
                                                                    fmt.Println(v3, v29)
                                                                }
                                                                _ = len(v32)
                                                            } else {
                                                                var v94 uint32 = v9
                                                                _ = v94
                                                            }
                                                            _ = v85
                                                        }
                                                        if ((55 == 2147483647) && true) {
                                                            fmt.Println(fi6)
                                                        } else {
                                                            fmt.Println(v29, v9, v2)
                                                        }
                                                        v2 -= int16(v3)
                                                        var v95 uint16 = (^(v30))
                                                        _ = v95
                                                    }
                                                    _ = v35
                                                    _ = v36
                                                }
                                            }
                                            v3 = (v3 | (v3 << ((255) & 63)))
                                            _ = v32
                                        }
                                        v3 = 765
                                        var v96 float64 = 1e10
                                        _ = v96
                                    } else {
                                        if (false && ((^(v2)) >= v2)) {
                                            v29 = v29
                                        } else {
                                            if true {
                                                var v97 int = fi6
                                                var v98 uint32 = 4294967295
                                                if (255 >= 2) {
                                                    fmt.Println(v9, fi8, v2)
                                                }
                                                _ = len(v1)
                                                fmt.Println(v30)
                                                switch "\n\t" {
                                                case "a":
                                                    v99 := GenS4{F0: v3, F1: 3}
                                                    _ = v99
                                                }
                                                v30 = 255
                                                _ = v97
                                                _ = v98
                                            } else {
                                                v100 := GenS10{F0: v9}
                                                _ = v100
                                            }
                                        }
                                        fmt.Println(v3, v30, v2)
                                    }
                                    var v101 string = "hello"
                                    if (!((v3 - v3) != 636)) {
                                        v102 := []int{}
                                        v9 = (v9 % ((v9) | 1))
                                        var v103 uint16 = (v30 | v30)
                                        _ = v102
                                        _ = v103
                                    }
                                    fmt.Println(v2, v9, v9)
                                    v104 := GenS15{F0: 2, F1: -575}
                                    if (116 == 42) {
                                        v101 = v101
                                        v105 := GenS4{F0: (-(v3)), F1: 0}
                                        var v106 float64 = 1.5
                                        if (!v29) {
                                            v101 = v101
                                        }
                                        _ = v105
                                        _ = v106
                                    } else {
                                        if (fi8 <= (-(fi6))) {
                                            if v29 {
                                                v107 := GenS15{F0: 4294967295, F1: 2147483647}
                                                if v29 {
                                                    v29 = (v101 == ("hello" + "foo bar"))
                                                    if v29 {
                                                        var v108 bool = true
                                                        var v109 float32 = 0.1
                                                        _ = v108
                                                        _ = v109
                                                    } else {
                                                        v110 := []int{fi8, 9223372036854775807, (fi6 ^ fi8)}
                                                        if (!(!((^(v30)) > (^(v30))))) {
                                                            fmt.Println(v2, v3, v29)
                                                            var v111 float64 = -2.25
                                                            _ = v111
                                                        } else {
                                                            fmt.Println(v101)
                                                            switch ("" + "hello") {
                                                            case "":
                                                                _ = len(v1)
                                                            case "\n\t":
                                                                if (0.0 != (3.14159 / float32((-2.25 / 1.0)))) {
                                                                    var v112 uint8 = 239
                                                                    _ = v112
                                                                } else {
                                                                    if v29 {
                                                                        if v29 {
                                                                            if len(v110) >= 2 {
                                                                                v110 = v110[1:len(v110)]
                                                                            }
                                                                            v3 += (v3 * v3)
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            v113 := GenS5{F0: 0, F1: "hello", F2: (v3 - (^(v3)))}
                                                            _ = v113
                                                        }
                                                        _ = len(v1)
                                                        _ = v110
                                                    }
                                                    fmt.Println(v2, v30)
                                                    v3 += (v3 | (-(v3)))
                                                    v114 := []int{fi6, 576, fi6}
                                                    v3 += (v3 << ((301) & 63))
                                                    _ = v114
                                                }
                                                var v115 uint16 = v30
                                                fmt.Println(v9, v2, v3)
                                                _ = v107
                                                _ = v115
                                            }
                                            v9 -= (v9 << ((413) & 63))
                                            v30 += 1
                                        }
                                        v116 := v1[v9]
                                        fmt.Println(v3)
                                        _ = v116
                                    }
                                    _ = v101
                                    _ = v104
                                    _ = v30
                                } else {
                                    var v117 float32 = 0.1
                                    v118 := GenS3{F0: ("\n\t" + "hello")}
                                    if false {
                                        v29 = (!true)
                                        if v29 {
                                            switch "\n\t" {
                                            case "":
                                                if v29 {
                                                    var v119 float32 = (float32((1.5 / ((1e10 * (1e10 / (0.0 + (0.1 - 1.5)))) / 1.0))) * v117)
                                                    v120 := []int{fi8, 1}
                                                    if false {
                                                        v121 := GenS16{F0: v29, F1: -70}
                                                        var v122 uint16 = 32767
                                                        v117 = ((v119 / 3.14159) + (1e10 + v119))
                                                        _ = v121
                                                        _ = v122
                                                    }
                                                    v123 := v1[(^(v9))]
                                                    switch fi8 {
                                                    case 45:
                                                        v124 := GenS10{F0: 1}
                                                        _ = v124
                                                    }
                                                    v123 = ((0.0 * v117) >= v117)
                                                    _ = v119
                                                    _ = v120
                                                    _ = v123
                                                } else {
                                                    v117 = ((1.5 / 1.0) * 0.0)
                                                    switch (("" + "hello") + "foo bar") {
                                                    case "foo bar":
                                                        var v125 float32 = (((v117 / ((1.5 / (((0.1 / float32((3.14159 - ((0.1 * (3.14159 + 3.14159)) + 1.5)))) / v117) / 1e10)) + 0.1)) + 1.5) / 1.0)
                                                        _ = v125
                                                    }
                                                    v29 = (!v29)
                                                }
                                            case "a":
                                                var v126 int8 = -71
                                                _ = v126
                                            case "hello":
                                                v2 -= v2
                                            }
                                            var v127 int64 = v3
                                            var v128 uint = 1
                                            _ = v127
                                            _ = v128
                                        } else {
                                            v3 = (v3 << ((1) & 63))
                                        }
                                        switch "" {
                                        case "hello":
                                            v129 := []int{0, fi8}
                                            _ = v129
                                        case "":
                                            if (!v29) {
                                                fmt.Println(v2)
                                            }
                                        case "foo bar":
                                            var v130 int8 = -128
                                            _ = v130
                                        }
                                    } else {
                                        fmt.Println(fi6, v3, v29)
                                        v131 := GenS16{F0: false, F1: 1}
                                        switch "hello" {
                                        case "foo bar":
                                            var v132 int = (fi6 % (((fi6 ^ (fi6 * fi8))) | 1))
                                            if (42 != 4294967295) {
                                                fmt.Println(v3)
                                                var v133 int16 = -1
                                                _ = v133
                                            }
                                            _ = v132
                                        case "hello":
                                            var v134 float32 = (0.1 * v117)
                                            _ = v134
                                        case "":
                                            var v135 int32 = -639
                                            _ = v135
                                        }
                                        _ = v131
                                    }
                                    _ = v117
                                    _ = v118
                                }
                            }
                            _ = v29
                        } else {
                            var v136 int8 = -21
                            switch "hello" {
                            case "":
                                var v137 int32 = -619
                                _ = v137
                            case "foo bar":
                                var v138 uint8 = 127
                                _ = v138
                            case "\n\t":
                                delete(v1, 4294967295)
                            default:
                                v139 := GenS14{F0: 608}
                                _ = v139
                            }
                            _ = v136
                        }
                    }
                    _ = v28
                }
                switch "" {
                case "hello":
                    var v140 int16 = -1
                    _ = v140
                }
                _ = fi8
                _ = v9
            }
        } else {
            v141 := v1[26]
            v141 = v141
            v2 -= -32768
            _ = v141
        }
        _ = fi6
        _ = v7
    }
    var v142 int32 = -2147483648
    var v143 uint = 875
    fmt.Println(v3, v142)
    var v144 uint = v143
    v3 -= 0
    fmt.Println(v2)
    var v145 int16 = (v2 ^ (v2 ^ v2))
    fmt.Println(v143)
    if true {
        switch ("hello" + "foo bar") {
        case "\n\t":
            fmt.Println(v2, v142, v145)
        }
        switch -9223372036854775808 {
        case 127:
            var v146 uint64 = 42
            _ = v146
        }
    }
    v147 := GenS17{F0: v142, F1: 9223372036854775807, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v142:", v142)
    fmt.Println("v143:", v143)
    fmt.Println("v144:", v144)
    fmt.Println("v145:", v145)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v142
    _ = v143
    _ = v144
    _ = v145
    _ = v147
    _ = v2
    _ = v3
}
