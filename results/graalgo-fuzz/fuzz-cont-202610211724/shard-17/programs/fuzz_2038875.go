package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }
type GenS2 struct { F0 string; F1 int8; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16; F1 int64; F2 int64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 int16; F1 uint32 }
type GenS11 struct { F0 uint16; F1 uint16 }
type GenS12 struct { F0 int16; F1 uint8 }
type GenS13 struct { F0 bool; F1 uint32; F2 uint32 }
func genFn0(p0 int8, p1 int32, p2 string) bool {
    if true {
        p1 -= 2147483647
        var v4 int16 = 32767
        _ = v4
    }
    p1 -= p1
    switch 42 {
    case 4294967295:
        var v5 float64 = 3.14159
        _ = v5
    }
    if true {
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}
func genFn1(p0 int64) uint8 {
    var v6 uint32 = 65535
    fmt.Println(p0, p0)
    _ = p0
    _ = v6
    return 127
}

func main() {
    var v1 uint64 = 580
    var v2 uint8 = 90
    v3 := make(map[int]bool)
    fmt.Println(v2, v1)
    var v7 string = "foo bar"
    fmt.Println(v1)
    var v8 int = 1
    if true {
        switch v7 {
        case "foo bar":
            if true {
                if true {
                    var v9 int8 = 1
                    switch v8 {
                    case 258:
                        if false {
                            v2 += (v2 % (((^(v2))) | 1))
                            var v10 uint16 = 65535
                            _ = v10
                        } else {
                            v1 = 4294967295
                        }
                    }
                    _ = v9
                }
            } else {
                v11 := GenS1{F0: -1, F1: v7, F2: v8}
                var v12 int8 = -71
                _ = v11
                _ = v12
            }
        case "":
            fmt.Println(v1, v2)
        default:
            v1 = (v1 / ((0) | 1))
        }
        var v13 string = v7
        if true {
            v14 := GenS1{F0: -1, F1: "a", F2: (v8 - (v8 / ((v8) | 1)))}
            v7 = ("a" + "hello")
            _ = v14
        }
        _ = v13
    } else {
        if (127 > 100) {
            switch v7 {
            case "a":
                for fi15 := 0; fi15 < 5; fi15++ {
                    if (!(1 == 0)) {
                        for fi16 := 0; fi16 < 1; fi16++ {
                            fmt.Println(v1)
                            switch (v7 + ((((v7 + v7) + v7) + "a") + v7)) {
                            case "\n\t":
                                var v17 uint = 18446744073709551615
                                _ = v17
                            }
                            _ = fi16
                        }
                        fmt.Println(v8, v2, v8)
                        if false {
                            v18 := GenS2{F0: v7, F1: -21, F2: (v7 + v7)}
                            _ = len(v3)
                            v19 := []int{32768, (^(fi15)), fi15}
                            _ = v18
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 10; fi20++ {
                                v21 := v3[(v8 | fi20)]
                                var v22 int64 = 1
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        }
                    } else {
                        switch v7 {
                        case "a":
                            v7 = ("\n\t" + v7)
                        case "foo bar":
                            v1 -= 823
                        default:
                            fmt.Println(v8)
                        }
                        v23 := []int{(^(fi15)), fi15, (v8 - v8)}
                        for fi24 := 0; fi24 < 11; fi24++ {
                            var v25 int8 = -1
                            v25 = v25
                            var v26 int16 = 0
                            v2 = (v2 + (v2 & v2))
                            v1 = v1
                            if false {
                                var v27 float32 = 1.5
                                _ = v27
                            }
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        _ = v23
                    }
                    var v28 float32 = 0.0
                    _ = fi15
                    _ = v28
                }
            case "":
                var v29 int32 = 2147483647
                _ = v29
            case "foo bar":
                for fi30 := 0; fi30 < 4; fi30++ {
                    var v31 int16 = 1
                    var v32 uint16 = 42
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            default:
                fmt.Println(v8, v2)
            }
        }
        for fi33 := 0; fi33 < 14; fi33++ {
            for fi34 := 0; fi34 < 6; fi34++ {
                if true {
                    v35 := GenS3{F0: 255}
                    v8 = (fi34 ^ (v8 & (fi33 * (fi33 & (fi34 + 2147483647)))))
                    v36 := GenS1{F0: -709, F1: v7, F2: fi33}
                    var v37 string = "a"
                    _ = v35
                    _ = v36
                    _ = v37
                } else {
                    if false {
                        var v38 uint = uint(v1)
                        v39 := GenS4{F0: (v8 & -2147483648), F1: 0}
                        v3[-991] = false
                        switch ((("\n\t" + v7) + v7) + "a") {
                        case "foo bar":
                            var v40 int = -254
                            _ = v40
                        case "hello":
                            switch fi33 {
                            case 32768:
                                fmt.Println(v7, v38, v7)
                            case -217:
                                _ = len(v3)
                            case 214:
                                v2 -= (v2 >> ((0) & 63))
                            }
                        default:
                            v41 := []int{}
                            _ = v41
                        }
                        v42 := []int{v8, (fi34 / (((fi34 + fi34)) | 1)), -720}
                        v42 = append(v42, (v8 >> (((v38 & (v38 % (((v38 / ((18446744073709551615) | 1))) | 1)))) & 63)))
                        fmt.Println(v8)
                        if true {
                            var v43 int = fi34
                            _ = v43
                        } else {
                            v44 := GenS1{F0: 0, F1: v7, F2: v8}
                            switch int(v1) {
                            case -826:
                                if true {
                                    if (-490 != -9223372036854775808) {
                                        var v45 int64 = 42
                                        var v46 int8 = 0
                                        v7 = v7
                                        _ = v45
                                        _ = v46
                                    } else {
                                        v1 = v1
                                        if true {
                                            v2 -= v2
                                            if true {
                                                if (0.1 == 0.0) {
                                                    v47 := GenS5{F0: v38}
                                                    _ = v47
                                                } else {
                                                    switch (fi34 / ((fi33) | 1)) {
                                                    case -606:
                                                        var v48 float32 = (float32(float64(1e10)) + 3.14159)
                                                        _ = v48
                                                    case 0:
                                                        switch v7 {
                                                        case "":
                                                            fmt.Println(v1, v1, v8)
                                                        default:
                                                            fmt.Println(v38)
                                                        }
                                                    }
                                                    if false {
                                                        v49 := GenS1{F0: 60, F1: v7, F2: (v8 * fi34)}
                                                        _ = v49
                                                    }
                                                }
                                                delete(v3, (v8 - fi34))
                                                fmt.Println(v38, fi33)
                                            }
                                            if (!false) {
                                                var v50 uint64 = (v1 | 380)
                                                v50 -= 0
                                                var v51 int8 = -1
                                                _ = v50
                                                _ = v51
                                            } else {
                                                fmt.Println(v7)
                                                if (-2.25 < (1e10 * (3.14159 / 1.0))) {
                                                    v52 := GenS6{F0: 1}
                                                    _ = v52
                                                } else {
                                                    v1 = 2
                                                    v53 := GenS2{F0: v7, F1: -79, F2: "\n\t"}
                                                    if (1.5 > 3.14159) {
                                                        var v54 float32 = (0.0 / 1.0)
                                                        _ = v54
                                                    } else {
                                                        if len(v42) >= 2 {
                                                            v42 = v42[1:len(v42)]
                                                        }
                                                        v38 = (v38 / (((v38 << ((uint(v1)) & 63))) | 1))
                                                    }
                                                    _ = v53
                                                }
                                                v2 -= v2
                                                v42 = append(v42, fi33)
                                                if true {
                                                    v55 := GenS7{F0: 0, F1: 42, F2: 623}
                                                    _ = v55
                                                }
                                            }
                                        } else {
                                            var v56 uint8 = v2
                                            if (535 >= 704) {
                                                if false {
                                                    if true {
                                                        if true {
                                                            v57 := GenS5{F0: (v38 % (((^(v38))) | 1))}
                                                            switch (v8 | (-(fi33))) {
                                                            case -9223372036854775808:
                                                                fmt.Println(fi34, v56, v56)
                                                            case -2147483648:
                                                                if true {
                                                                    switch v7 {
                                                                    case "foo bar":
                                                                        var v58 int8 = 1
                                                                        _ = v58
                                                                    case "\n\t":
                                                                        v56 = genFn1(484)
                                                                    default:
                                                                        if false {
                                                                            if false {
                                                                                switch ((v7 + ("hello" + v7)) + v7) {
                                                                                case "\n\t":
                                                                                    var v59 float32 = 1e10
                                                                                    _ = v59
                                                                                case "a":
                                                                                    var v60 int16 = 783
                                                                                    _ = v60
                                                                                case "hello":
                                                                                    switch ((v7 + "\n\t") + ((((v7 + v7) + v7) + "\n\t") + v7)) {
                                                                                    case "hello":
                                                                                        var v61 int = (v8 - (fi33 ^ fi34))
                                                                                        _ = v61
                                                                                    case "a":
                                                                                        if len(v42) > 0 {
                                                                                            v42[0] = v8
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v62 := GenS3{F0: 220}
                                                                                    fmt.Println(fi34, v2)
                                                                                    _ = v62
                                                                                }
                                                                            } else {
                                                                                v8 -= fi33
                                                                                var v63 string = "hello"
                                                                                var v64 int64 = 0
                                                                                if (true || false) {
                                                                                    if (!true) {
                                                                                        v3[(v8 * (v8 | -933))] = true
                                                                                        var v65 float64 = float64(((((1.5 + 0.1) / 1.0) * -2.25) / ((3.14159 - (1.5 + 1.5)) - (1.5 / 1.0))))
                                                                                        var v66 float64 = v65
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    } else {
                                                                                        var v67 int = fi33
                                                                                        _ = v67
                                                                                    }
                                                                                    v68 := GenS6{F0: 209}
                                                                                    _ = v68
                                                                                }
                                                                                var v69 float32 = (((1.5 + (-2.25 + 0.0)) + float32(float64(0.1))) * (0.1 - 0.0))
                                                                                _ = v63
                                                                                _ = v64
                                                                                _ = v69
                                                                            }
                                                                            var v70 uint16 = uint16(v38)
                                                                            v71 := GenS2{F0: "a", F1: 113, F2: (("hello" + v7) + (("a" + v7) + v7))}
                                                                            fmt.Println(fi34, v8)
                                                                            _ = v70
                                                                            _ = v71
                                                                        } else {
                                                                            v72 := GenS5{F0: (v38 >> ((2) & 63))}
                                                                            var v73 uint64 = 1
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                    }
                                                                    v7 = (v7 + v7)
                                                                } else {
                                                                    v74 := GenS3{F0: v38}
                                                                    fmt.Println(fi34, fi33, v7)
                                                                    _ = v74
                                                                }
                                                            case 2147483647:
                                                                v42 = append(v42, (v8 % ((int(v38)) | 1)))
                                                                fmt.Println(v56, fi33, v1)
                                                            default:
                                                                if (true && (-582 >= -1)) {
                                                                    v75 := GenS8{F0: v38}
                                                                    v42 = append(v42, v8)
                                                                    _ = v75
                                                                } else {
                                                                    v76 := v3[fi34]
                                                                    var v77 uint8 = (v2 ^ 1)
                                                                    _ = v76
                                                                    _ = v77
                                                                }
                                                            }
                                                            v1 = (^(v1))
                                                            _ = v57
                                                        } else {
                                                            if true {
                                                                var v78 string = v7
                                                                _ = v78
                                                            }
                                                            if ((v1 | (v1 + uint64(v38))) >= uint64(v38)) {
                                                                v3[fi33] = true
                                                                if (-679 >= -32768) {
                                                                    if true {
                                                                        var v79 uint64 = (v1 & (v1 & v1))
                                                                        _ = v79
                                                                    }
                                                                    if len(v42) >= 2 {
                                                                        v42 = v42[1:len(v42)]
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        var v80 bool = (-35 >= 9223372036854775807)
                                                        var v81 uint8 = (^(v2))
                                                        var v82 uint16 = 0
                                                        if (true || v80) {
                                                            fmt.Println(v82, v38, fi34)
                                                        }
                                                        _ = v80
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                } else {
                                                    var v83 int16 = 0
                                                    var v84 uint32 = 338
                                                    _ = v83
                                                    _ = v84
                                                }
                                                if (v38 > v38) {
                                                    if true {
                                                        if (v2 < v2) {
                                                            v8 = (fi33 * 32768)
                                                            var v85 uint = v38
                                                            fmt.Println(v56, v8, v2)
                                                            _ = v85
                                                        } else {
                                                            var v86 uint = 861
                                                            _ = v86
                                                        }
                                                        fmt.Println(v38, v38, v2)
                                                        v1 = (v1 * v1)
                                                    } else {
                                                        v56 = 255
                                                    }
                                                }
                                                switch (("foo bar" + v7) + "a") {
                                                case "\n\t":
                                                    fmt.Println(v56)
                                                case "hello":
                                                    _ = len(v42)
                                                case "a":
                                                    var v87 uint = 42
                                                    _ = v87
                                                default:
                                                    if false {
                                                        if (-1 != 32767) {
                                                            v88 := GenS7{F0: 42, F1: 9223372036854775807, F2: 0}
                                                            fmt.Println(fi34, fi33)
                                                            _ = v88
                                                        }
                                                    } else {
                                                        var v89 uint64 = v1
                                                        v90 := GenS9{F0: 32767, F1: 848}
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                }
                                            } else {
                                                switch fi33 {
                                                case 9223372036854775807:
                                                    if false {
                                                        fmt.Println(fi34)
                                                        switch "" {
                                                        case "foo bar":
                                                            v91 := GenS10{F0: 509, F1: 65535}
                                                            _ = v91
                                                        }
                                                        v38 = (v38 / ((968) | 1))
                                                    }
                                                case 786:
                                                    v92 := GenS8{F0: v38}
                                                    _ = v92
                                                default:
                                                    if true {
                                                        var v93 uint64 = v1
                                                        if len(v42) >= 2 {
                                                            v42 = v42[1:len(v42)]
                                                        }
                                                        fmt.Println(v38, fi34, v2)
                                                        _ = v93
                                                    }
                                                }
                                                _ = len(v42)
                                                switch (v7 + "a") {
                                                case "":
                                                    fmt.Println(v56, fi33)
                                                case "\n\t":
                                                    var v94 int = fi34
                                                    _ = v94
                                                case "hello":
                                                    var v95 int16 = 1
                                                    _ = v95
                                                }
                                            }
                                            v1 = v1
                                            _ = v56
                                        }
                                    }
                                }
                            default:
                                fmt.Println(fi34)
                            }
                            var v96 uint16 = 883
                            _ = v44
                            _ = v96
                        }
                        _ = v38
                        _ = v39
                        _ = v42
                    }
                    if true {
                        var v97 string = (v7 + v7)
                        _ = v97
                    } else {
                        v7 = v7
                        v98 := GenS11{F0: 32767, F1: 255}
                        _ = v98
                    }
                }
                fmt.Println(v2)
                v1 = v1
                _ = fi34
            }
            v2 += v2
            _ = fi33
        }
    }
    fmt.Println(v7, v1)
    fmt.Println(v8)
    if ((-(v8)) >= v8) {
        v99 := []int{(v8 << ((20) & 63)), -2147483648, -9223372036854775808}
        fmt.Println(v8, v8)
        var v100 bool = true
        _ = v100
        _ = v99
    }
    v101 := GenS12{F0: -792, F1: v2}
    v102 := []int{}
    v8 += (v8 / (((v8 * v8)) | 1))
    v8 = v8
    if false {
        if true {
            v103 := GenS13{F0: false, F1: 65535, F2: 2}
            v104 := GenS10{F0: 1000, F1: 1}
            v105 := v3[(^(v8))]
            _ = v103
            _ = v104
            _ = v105
        }
        var v106 int8 = -76
        _ = v106
    }
    for fi107 := 0; fi107 < 16; fi107++ {
        _ = len(v3)
        _ = fi107
    }
    var v108 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v102:", len(v102), v102)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v101
    _ = v102
    _ = v108
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
