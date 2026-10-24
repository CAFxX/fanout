package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int32; F2 int32 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int8; F1 bool; F2 uint }
type GenS7 struct { F0 int8; F1 int8; F2 int }
type GenS8 struct { F0 int64; F1 int }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string; F1 int64; F2 uint }
type GenS11 struct { F0 uint16; F1 string; F2 int32 }
type GenS12 struct { F0 uint32; F1 uint; F2 bool }
type GenS13 struct { F0 int64; F1 int32; F2 int16 }
type GenS14 struct { F0 uint8; F1 uint64 }
type GenS15 struct { F0 uint32; F1 uint64; F2 int }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint; F1 bool; F2 int64 }
type GenS18 struct { F0 uint32; F1 string; F2 int16 }
func genFn0(p0 int64) int {
    p0 = (p0 | 0)
    p0 = 1
    v4 := GenS1{F0: 140, F1: -244, F2: p0}
    v5 := []int{}
    var v6 uint8 = 0
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 128
}
func genFn1() int16 {
    v7 := GenS1{F0: 2, F1: -2147483648, F2: -9223372036854775808}
    switch 127 {
    case 42:
        v8 := GenS1{F0: 0, F1: -2147483648, F2: -123}
        _ = v8
    case 127:
        v9 := []int{}
        _ = v9
    case -384:
        var v10 int8 = -42
        _ = v10
    default:
        var v11 float32 = (1e10 + -2.25)
        _ = v11
    }
    _ = v7
    return 32767
}

func main() {
    var v1 uint64 = 0
    v2 := []string{"foo bar"}
    v3 := make(map[int]bool)
    v1 = v1
    var v12 uint8 = 1
    switch 127 {
    case 2147483647:
        if (0.0 < 1e10) {
            v1 = (v1 & 82)
            var v13 int = 127
            _ = v13
        } else {
            var v14 string = (("hello" + "\n\t") + "foo bar")
            _ = v14
        }
    default:
        _ = len(v3)
    }
    v15 := GenS1{F0: 1, F1: 42, F2: 492}
    if (2 != 2) {
        _ = len(v3)
        var v16 int = -55
        switch (-(v16)) {
        case -1:
            var v17 int8 = 127
            _ = v17
        case 42:
            if true {
                v18 := GenS2{F0: (v12 - (v12 * v12)), F1: -32768}
                for fi19 := 0; fi19 < 1; fi19++ {
                    v12 -= v12
                    var v20 int64 = 787
                    v16 -= (^(v16))
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
        }
        _ = v16
    } else {
        v21 := v3[127]
        var v22 uint = 18446744073709551615
        _ = len(v3)
        _ = v21
        _ = v22
    }
    var v23 bool = ((0 < 255) && ((float32(0.0) > ((1.5 - (((float32(1.5) - (-2.25 / 0.1)) * (0.0 / 3.14159)) + (1e10 * 1e10))) * 3.14159)) || false))
    var v24 uint = 65535
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := v3[fi25]
        if v23 {
            if (v12 == uint8(v24)) {
                fmt.Println(fi25)
                v3[(fi25 * (-(fi25)))] = (true && (((3.14159 - 0.1) - 0.1) != 3.14159))
            }
            v27 := GenS2{F0: 128, F1: -323}
            _ = v27
        } else {
            if false {
                if (true && false) {
                    v28 := GenS3{F0: (v1 * v1)}
                    _ = v28
                } else {
                    switch (fi25 / ((genFn0(1)) | 1)) {
                    case -730:
                        var v29 int32 = -2147483648
                        _ = v29
                    case 32767:
                        for fi30 := 0; fi30 < 9; fi30++ {
                            switch (fi25 / ((0) | 1)) {
                            case 42:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 791:
                                v31 := GenS4{F0: (("foo bar" + (("a" + ("foo bar" + (("" + "foo bar") + ""))) + ("" + "hello"))) + ("foo bar" + "hello")), F1: -1, F2: -972}
                                _ = v31
                            }
                            _ = fi30
                        }
                    default:
                        _ = len(v2)
                    }
                    for fi32 := 0; fi32 < 6; fi32++ {
                        v33 := GenS2{F0: 128, F1: -929}
                        v23 = (0 <= int16(v12))
                        _ = fi32
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 10; fi34++ {
                        if ((100 < -128) && v26) {
                            if v26 {
                                v1 += 4294967295
                            } else {
                                var v35 int64 = -809
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v36 uint64 = (v1 * 18446744073709551615)
                                _ = v35
                                _ = v36
                            }
                            delete(v3, (fi25 / ((genFn0(0)) | 1)))
                            var v37 float64 = 3.14159
                            _ = v37
                        } else {
                            switch "\n\t" {
                            case "hello":
                                v23 = v23
                            case "\n\t":
                                var v38 uint16 = 0
                                _ = v38
                            default:
                                var v39 uint32 = 2
                                v24 = 1
                                var v40 int8 = -128
                                _ = v39
                                _ = v40
                            }
                            switch fi25 {
                            case -2147483648:
                                v26 = false
                            case 113:
                                if false {
                                    switch 573 {
                                    case 88:
                                        var v41 bool = (!("a" >= ("foo bar" + "foo bar")))
                                        _ = v41
                                    default:
                                        if (v26 && (genFn0(9223372036854775807) >= fi25)) {
                                            if (1.5 < 1.5) {
                                                v26 = false
                                            }
                                            v1 = uint64(v12)
                                            if (!v26) {
                                                v26 = (!true)
                                            } else {
                                                if (!("hello" != (("foo bar" + ("hello" + "a")) + "\n\t"))) {
                                                    v12 = (v12 & v12)
                                                    v26 = (232 >= -1)
                                                    v42 := GenS2{F0: 152, F1: -106}
                                                    var v43 int = fi25
                                                    fmt.Println(v23)
                                                    switch "" {
                                                    case "a":
                                                        fmt.Println(v12, v12, fi34)
                                                    default:
                                                        v44 := GenS3{F0: uint64(v12)}
                                                        _ = v44
                                                    }
                                                    v45 := GenS5{F0: "hello", F1: 42}
                                                    _ = v42
                                                    _ = v43
                                                    _ = v45
                                                }
                                                v46 := GenS3{F0: v1}
                                                v47 := GenS6{F0: -1, F1: (!v26), F2: v24}
                                                if false {
                                                    v48 := GenS7{F0: 63, F1: 127, F2: fi34}
                                                    fmt.Println(v24)
                                                    fmt.Println(fi25, v12, v26)
                                                    _ = v48
                                                } else {
                                                    var v49 string = "\n\t"
                                                    var v50 float64 = 3.14159
                                                    v50 = v50
                                                    _ = v49
                                                    _ = v50
                                                }
                                                _ = v46
                                                _ = v47
                                            }
                                        } else {
                                            if v23 {
                                                v51 := GenS7{F0: 1, F1: 85, F2: genFn0(-9223372036854775808)}
                                                if v23 {
                                                    v2 = append(v2, "foo bar")
                                                    switch (fi34 % ((fi34) | 1)) {
                                                    case 577:
                                                        if false {
                                                            fmt.Println(fi34, v26, v24)
                                                        } else {
                                                            if (((((3.14159 / 1.0) * (-2.25 + 1.5)) + (1e10 * -2.25)) / 1.0) <= 1e10) {
                                                                var v52 uint8 = 255
                                                                _ = len(v3)
                                                                fmt.Println(v26, v26, fi34)
                                                                _ = v52
                                                            } else {
                                                                _ = len(v3)
                                                                v53 := GenS8{F0: -9223372036854775808, F1: (fi34 << (((v24 * v24)) & 63))}
                                                                _ = v53
                                                            }
                                                        }
                                                    case 9223372036854775807:
                                                        v54 := GenS9{F0: 100}
                                                        _ = v54
                                                    case 4294967295:
                                                        fmt.Println(v23, fi25, v24)
                                                    default:
                                                        var v55 float32 = (0.0 * 1.5)
                                                        _ = v55
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = ("hello" + "")
                                                    }
                                                }
                                                v56 := GenS4{F0: "hello", F1: 0, F2: -2147483648}
                                                v24 = v24
                                                if true {
                                                    switch fi25 {
                                                    case 32768:
                                                        switch genFn0(1) {
                                                        case -269:
                                                            fmt.Println(v26)
                                                        case 9223372036854775807:
                                                            if ("a" <= "") {
                                                                v57 := GenS9{F0: -128}
                                                                v23 = v26
                                                                var v58 int32 = -1
                                                                v59 := GenS10{F0: "", F1: 0, F2: (^(v24))}
                                                                fmt.Println(fi34, v23)
                                                                if v23 {
                                                                    var v60 uint = (v24 % ((2) | 1))
                                                                    switch (("" + "hello") + ("a" + "a")) {
                                                                    case "hello":
                                                                        var v61 uint16 = 255
                                                                        _ = v61
                                                                    case "foo bar":
                                                                        var v62 uint = v60
                                                                        _ = v62
                                                                    default:
                                                                        v23 = v26
                                                                    }
                                                                    _ = v60
                                                                } else {
                                                                    switch (fi34 & (fi34 ^ 630)) {
                                                                    case -992:
                                                                        if false {
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        }
                                                                    case 42:
                                                                        v1 += v1
                                                                    default:
                                                                        var v63 uint8 = v12
                                                                        if (!false) {
                                                                            if false {
                                                                                var v64 uint16 = 0
                                                                                v1 += (^(v1))
                                                                                _ = v64
                                                                            }
                                                                            v65 := GenS3{F0: v1}
                                                                            v66 := GenS11{F0: 255, F1: "foo bar", F2: 289}
                                                                            fmt.Println(v23, v24)
                                                                            _ = v65
                                                                            _ = v66
                                                                        } else {
                                                                            switch "" {
                                                                            case "hello":
                                                                                _ = len(v3)
                                                                            }
                                                                            v2 = append(v2, "hello")
                                                                            v26 = (v26 && ((2 <= 65535) || v26))
                                                                        }
                                                                        _ = v63
                                                                    }
                                                                }
                                                                _ = v57
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                        case 32767:
                                                            if (30 > 0) {
                                                                if v23 {
                                                                    v23 = v26
                                                                    if (uint32(v1) < 1) {
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        case "hello":
                                                                            delete(v3, genFn0(-38))
                                                                        case "foo bar":
                                                                            v67 := GenS12{F0: 65535, F1: v24, F2: v23}
                                                                            _ = v67
                                                                        default:
                                                                            var v68 int = (-(fi25))
                                                                            _ = v68
                                                                        }
                                                                    }
                                                                    fmt.Println(v23, v1)
                                                                } else {
                                                                    v1 = (v1 ^ v1)
                                                                }
                                                                v69 := GenS13{F0: 1, F1: 540, F2: genFn1()}
                                                                _ = v69
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    if (!(true && (42 < 131))) {
                                                        if false {
                                                            var v70 int16 = 42
                                                            v12 = (^(v12))
                                                            var v71 float64 = 0.0
                                                            var v72 uint64 = (v1 << (((v24 % ((uint(v1)) | 1))) & 63))
                                                            v73 := GenS9{F0: 42}
                                                            _ = v70
                                                            _ = v71
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    } else {
                                                        v74 := GenS10{F0: ("a" + "hello"), F1: -9223372036854775808, F2: 2}
                                                        fmt.Println(fi34)
                                                        if (((((float32((0.1 - 0.0)) - float32(3.14159)) - 0.0) - -2.25) >= (float32(1e10) + (0.1 - -2.25))) && v26) {
                                                            var v75 int8 = -128
                                                            fmt.Println(fi34)
                                                            _ = v75
                                                        } else {
                                                            v1 = (^(v1))
                                                            if len(v2) > 0 {
                                                                v2[0] = "a"
                                                            }
                                                            fmt.Println(v24)
                                                        }
                                                        _ = v74
                                                    }
                                                    switch "" {
                                                    case "\n\t":
                                                        v76 := GenS14{F0: 222, F1: (v1 >> ((v24) & 63))}
                                                        _ = v76
                                                    }
                                                }
                                                _ = v51
                                                _ = v56
                                            }
                                        }
                                        switch genFn0(-9223372036854775808) {
                                        case -2147483648:
                                            var v77 float64 = 1.5
                                            _ = v77
                                        default:
                                            if (v26 && v23) {
                                                fmt.Println(v12, fi25, v1)
                                            }
                                        }
                                    }
                                    v78 := GenS15{F0: 2, F1: (^(v1)), F2: fi34}
                                    v2 = append(v2, "\n\t")
                                    _ = v78
                                } else {
                                    if v26 {
                                        v23 = (571 == 727)
                                    }
                                    v24 = (v24 + v24)
                                    var v79 int16 = 1
                                    v3[(fi25 ^ (fi25 | (fi25 - (fi25 & fi34))))] = v26
                                    _ = v79
                                }
                            default:
                                v1 = v1
                            }
                            v80 := GenS16{F0: 1}
                            v81 := GenS17{F0: v24, F1: v26, F2: -780}
                            _ = v80
                            _ = v81
                        }
                        switch "a" {
                        case "foo bar":
                            v3[int(v24)] = v26
                        case "a":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                        default:
                            v26 = v23
                        }
                        _ = fi34
                    }
                }
                v23 = (!(248 != v12))
                v12 = v12
            }
        }
        switch "hello" {
        case "foo bar":
            v23 = ((v12 >> (((v24 + 1)) & 63)) > uint8(v1))
        }
        _ = fi25
        _ = v26
    }
    v1 += 41
    v82 := v3[4294967295]
    switch 32768 {
    case 1:
        v82 = (v82 && v82)
    case 189:
        v23 = v82
    case 2147483647:
        v12 = 1
    default:
        v83 := GenS18{F0: 42, F1: "hello", F2: -32768}
        _ = v83
    }
    v84 := GenS14{F0: 120, F1: v1}
    for fi85 := 0; fi85 < 20; fi85++ {
        if len(v2) > 0 {
            v2[0] = ("a" + "\n\t")
        }
        v82 = (v24 < v24)
        switch "" {
        case "\n\t":
            if (((1e10 * (0.0 - ((float64(0.0) * 1.5) / 1.0))) + (1.5 + 1.5)) < 0.0) {
                _ = len(v2)
                v86 := GenS15{F0: 583, F1: 789, F2: genFn0(-222)}
                var v87 float64 = float64(1.5)
                _ = v86
                _ = v87
            } else {
                v23 = v23
                var v88 int8 = 100
                _ = v88
            }
        case "foo bar":
            switch ("a" + "a") {
            case "\n\t":
                var v89 int32 = 0
                _ = v89
            case "hello":
                v1 = (^(v1))
            case "foo bar":
                for fi90 := 0; fi90 < 1; fi90++ {
                    if (!v82) {
                        v91 := GenS1{F0: 0, F1: -1, F2: -285}
                        fmt.Println(v24, v24, v24)
                        v12 = 0
                        _ = v91
                    } else {
                        v12 = (v12 ^ v12)
                    }
                    if ("hello" < "\n\t") {
                        var v92 bool = v82
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        v93 := GenS8{F0: int64(fi85), F1: 32767}
                        _ = v92
                        _ = v93
                    } else {
                        var v94 int = fi85
                        var v95 uint64 = uint64(v12)
                        var v96 int64 = -1
                        _ = v94
                        _ = v95
                        _ = v96
                    }
                    _ = fi90
                }
            default:
                for fi97 := 0; fi97 < 4; fi97++ {
                    v98 := GenS4{F0: "\n\t", F1: -2147483648, F2: 1}
                    v99 := GenS9{F0: -128}
                    v12 = (v12 % (((v12 | (v12 + (v12 % (((v12 ^ v12)) | 1))))) | 1))
                    v23 = true
                    fmt.Println(fi85)
                    fmt.Println(v82)
                    _ = fi97
                    _ = v98
                    _ = v99
                }
            }
        default:
            v100 := GenS5{F0: "", F1: -32768}
            _ = v100
        }
        _ = fi85
    }
    v1 += v1
    if v82 {
        fmt.Println(v12)
        v82 = v23
    }
    v12 = v12
    var v101 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v82:", v82)
    _ = v1
    _ = v101
    _ = v12
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v82
    _ = v84
}
