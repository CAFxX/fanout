package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 int16; F2 int }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int64; F1 uint32; F2 uint }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int; F1 uint16 }
type GenS14 struct { F0 int64; F1 uint16; F2 bool }
type GenS15 struct { F0 uint }
func genFn0(p0 bool, p1 int16, p2 uint) string {
    v4 := GenS1{F0: 825}
    fmt.Println(p0, p0)
    v5 := GenS2{F0: 227, F1: 128}
    v6 := GenS3{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return ""
}
func genFn1() int32 {
    if (!true) {
        if (!(197 == 42)) {
            var v7 bool = true
            v8 := []int{}
            _ = v7
            _ = v8
        }
    }
    if (4294967295 <= 345) {
        v9 := []int{-712, 4294967295, 1, 378}
        v10 := GenS4{F0: 1}
        _ = v10
        _ = v9
    }
    var v11 uint64 = 768
    var v12 int64 = 9223372036854775807
    _ = v11
    _ = v12
    return -2147483648
}

func main() {
    var v1 int16 = -845
    var v2 int8 = 0
    var v3 float64 = 0.0
    if true {
        fmt.Println(v1)
        if false {
            fmt.Println(v3, v2, v3)
            v13 := []int{-628, 128, 128, 42}
            _ = v13
        } else {
            v14 := GenS4{F0: 2147483647}
            fmt.Println(v3, v3)
            _ = v14
        }
    } else {
        var v15 int16 = v1
        _ = v15
    }
    if (!(1e10 != float32((0.0 * 0.0)))) {
        if true {
            fmt.Println(v2, v2)
            if true {
                v16 := GenS5{F0: 1}
                _ = v16
            }
            if true {
                if false {
                    var v17 int64 = -1
                    _ = v17
                } else {
                    var v18 int64 = 324
                    switch 4294967295 {
                    case 2147483647:
                        v3 = (v3 * ((3.14159 * v3) - (1.5 - float64(1.5))))
                    case 9223372036854775807:
                        if false {
                            for fi19 := 0; fi19 < 5; fi19++ {
                                v20 := GenS6{F0: 839}
                                for fi21 := 0; fi21 < 15; fi21++ {
                                    v18 -= v18
                                    if true {
                                        var v22 uint16 = 510
                                        v23 := GenS7{F0: v22, F1: -32768, F2: 128}
                                        _ = v22
                                        _ = v23
                                    }
                                    if true {
                                        v24 := GenS6{F0: 639}
                                        v18 = (v18 - (v18 << ((0) & 63)))
                                        if (398 < 42) {
                                            var v25 int32 = 1
                                            var v26 uint64 = 42
                                            var v27 uint16 = 0
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        } else {
                                            fmt.Println(v1)
                                        }
                                        _ = v24
                                    } else {
                                        v3 = (v3 - 0.1)
                                    }
                                    _ = fi21
                                }
                                _ = fi19
                                _ = v20
                            }
                        } else {
                            var v28 float64 = 0.1
                            v18 = 408
                            v29 := GenS3{F0: -1}
                            _ = v28
                            _ = v29
                        }
                    }
                    var v30 int64 = v18
                    _ = v18
                    _ = v30
                }
                v1 -= 548
                v31 := GenS3{F0: 1}
                _ = v31
            } else {
                if true {
                    for fi32 := 0; fi32 < 5; fi32++ {
                        var v33 string = "\n\t"
                        if (582 < 18446744073709551615) {
                            if (true && true) {
                                fmt.Println(v33)
                                for fi34 := 0; fi34 < 1; fi34++ {
                                    v2 += -1
                                    if false {
                                        switch v33 {
                                        case "hello":
                                            v2 = (v2 / ((v2) | 1))
                                        case "a":
                                            v35 := []int{fi32, (fi32 / ((fi32) | 1)), 4294967295, fi32}
                                            switch v33 {
                                            case "hello":
                                                if (!(127 == fi32)) {
                                                    switch ((v33 + ((v33 + "a") + v33)) + v33) {
                                                    case "\n\t":
                                                        var v36 bool = false
                                                        v2 = (v2 | (v2 >> ((42) & 63)))
                                                        _ = v36
                                                    case "foo bar":
                                                        fmt.Println(fi32, v1)
                                                    case "a":
                                                        v37 := GenS7{F0: 136, F1: 0, F2: 42}
                                                        _ = v37
                                                    default:
                                                        v38 := GenS3{F0: 42}
                                                        _ = v38
                                                    }
                                                    if (true && true) {
                                                        if ((2 == 1) || (false || false)) {
                                                            fmt.Println(fi34, v2)
                                                            v39 := GenS8{F0: 32767, F1: 1}
                                                            v40 := GenS9{F0: -182}
                                                            _ = v39
                                                            _ = v40
                                                        } else {
                                                            v2 = v2
                                                        }
                                                        var v41 uint = 13
                                                        _ = v41
                                                    } else {
                                                        switch -475 {
                                                        case -870:
                                                            fmt.Println(v1)
                                                        case -2147483648:
                                                            if ((!((0.1 + 1e10) >= float32(((v3 * float64(3.14159)) - (float64((-2.25 * 3.14159)) - v3))))) && true) {
                                                                fmt.Println(v33, v33, fi34)
                                                                var v42 string = (v33 + "hello")
                                                                _ = v42
                                                            }
                                                        default:
                                                            _ = len(v35)
                                                        }
                                                    }
                                                    v3 -= (float64(((float32(v3) / float32(((((0.0 + ((v3 / -2.25) + (v3 * v3))) - float64(3.14159)) + v3) + (v3 - (v3 / v3))))) * 0.0)) - v3)
                                                }
                                            case "":
                                                if len(v35) > 0 {
                                                    v35[0] = 127
                                                }
                                            default:
                                                v35 = append(v35, (fi34 / ((fi32) | 1)))
                                            }
                                            _ = v35
                                        case "foo bar":
                                            var v43 uint8 = 128
                                            _ = v43
                                        }
                                    }
                                    _ = fi34
                                }
                            } else {
                                fmt.Println(v2, v3, v33)
                                if false {
                                    for fi44 := 0; fi44 < 13; fi44++ {
                                        v45 := GenS10{F0: 33}
                                        v46 := []int{(fi44 << ((18446744073709551615) & 63)), fi32, fi44, fi32}
                                        v1 = -659
                                        v47 := GenS9{F0: -2147483648}
                                        _ = len(v46)
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    }
                                    if true {
                                        if (!(((40 < 200) && (false || false)) && false)) {
                                            v33 = v33
                                            var v48 int16 = (v1 | 42)
                                            _ = v48
                                        }
                                    } else {
                                        var v49 int64 = 0
                                        v33 = "\n\t"
                                        var v50 float32 = -2.25
                                        _ = v49
                                        _ = v50
                                    }
                                    var v51 int = (-(fi32))
                                    _ = v51
                                } else {
                                    var v52 int32 = 2147483647
                                    v33 = ("\n\t" + ("a" + (((v33 + v33) + (v33 + v33)) + "foo bar")))
                                    _ = v52
                                }
                                switch "hello" {
                                case "":
                                    var v53 float64 = float64((1e10 - 0.1))
                                    _ = v53
                                default:
                                    for fi54 := 0; fi54 < 18; fi54++ {
                                        v55 := GenS11{F0: -1, F1: 304, F2: 2}
                                        var v56 int32 = 419
                                        _ = fi54
                                        _ = v55
                                        _ = v56
                                    }
                                }
                            }
                            v3 = v3
                            if false {
                                v2 = (-(v2))
                                v57 := GenS12{F0: v33}
                                _ = v57
                            } else {
                                if false {
                                    switch v33 {
                                    case "\n\t":
                                        var v58 bool = true
                                        _ = v58
                                    default:
                                        for fi59 := 0; fi59 < 13; fi59++ {
                                            var v60 uint32 = 42
                                            v1 -= 395
                                            _ = fi59
                                            _ = v60
                                        }
                                    }
                                } else {
                                    v1 -= (v1 << ((0) & 63))
                                    for fi61 := 0; fi61 < 5; fi61++ {
                                        var v62 float32 = 1e10
                                        v33 = (v33 + v33)
                                        v63 := []int{fi61, fi61, (-(fi61)), (fi32 ^ (fi61 - fi32))}
                                        _ = fi61
                                        _ = v62
                                        _ = v63
                                    }
                                    fmt.Println(v3)
                                }
                            }
                        } else {
                            v64 := GenS4{F0: (fi32 & 128)}
                            _ = v64
                        }
                        _ = fi32
                        _ = v33
                    }
                }
                if false {
                    switch "foo bar" {
                    case "hello":
                        if (1 > 0) {
                            var v65 uint16 = 865
                            _ = v65
                        }
                    case "a":
                        if (v3 != 0.0) {
                            fmt.Println(v3, v1, v1)
                            if (false && false) {
                                v66 := GenS9{F0: 2147483647}
                                _ = v66
                            }
                        } else {
                            if ((18446744073709551615 == 65535) || false) {
                                for fi67 := 0; fi67 < 16; fi67++ {
                                    v68 := GenS13{F0: fi67, F1: 525}
                                    var v69 bool = false
                                    _ = fi67
                                    _ = v68
                                    _ = v69
                                }
                                fmt.Println(v3, v1)
                                v1 = (^(v1))
                            }
                            if ("\n\t" < "\n\t") {
                                var v70 int = -271
                                v1 = v1
                                _ = v70
                            } else {
                                var v71 uint32 = 0
                                var v72 int16 = v1
                                fmt.Println(v3, v2, v72)
                                _ = v71
                                _ = v72
                            }
                        }
                    case "":
                        var v73 uint = 18446744073709551615
                        _ = v73
                    }
                    var v74 float64 = 0.1
                    _ = v74
                }
                var v75 int64 = -1
                _ = v75
            }
        } else {
            var v76 int8 = -1
            _ = v76
        }
        var v77 bool = (!(false && (true || (!false))))
        if false {
            for fi78 := 0; fi78 < 9; fi78++ {
                v79 := GenS9{F0: 0}
                _ = fi78
                _ = v79
            }
            v1 = v1
        }
        _ = v77
    } else {
        for fi80 := 0; fi80 < 3; fi80++ {
            v81 := GenS1{F0: 227}
            v82 := []int{fi80, 295, fi80, fi80}
            if true {
                switch (fi80 >> ((42) & 63)) {
                case 32768:
                    var v83 string = "\n\t"
                    _ = v83
                case 0:
                    var v84 uint16 = 0
                    _ = v84
                case -2147483648:
                    v85 := GenS14{F0: -1, F1: 450, F2: true}
                    _ = v85
                }
                v86 := GenS8{F0: 1, F1: 1}
                _ = v86
            } else {
                var v87 int32 = 2147483647
                if len(v82) >= 2 {
                    v82 = v82[1:len(v82)]
                }
                v88 := GenS13{F0: -613, F1: 255}
                _ = v87
                _ = v88
            }
            _ = fi80
            _ = v81
            _ = v82
        }
    }
    if true {
        v89 := []int{9223372036854775807}
        _ = v89
    }
    v3 = v3
    fmt.Println(v2, v3, v2)
    var v90 int = 458
    v91 := []int{(v90 | (v90 * v90)), 9223372036854775807, (v90 % ((v90) | 1))}
    for fi92 := 0; fi92 < 6; fi92++ {
        switch (("hello" + "\n\t") + "") {
        case "foo bar":
            if len(v91) > 0 {
                v91[0] = (^(fi92))
            }
        case "":
            v2 += v2
        default:
            v93 := GenS2{F0: 42, F1: 2}
            _ = v93
        }
        for fi94 := 0; fi94 < 19; fi94++ {
            fmt.Println(v90, fi94, fi92)
            _ = fi94
        }
        _ = fi92
    }
    var v95 int8 = (v2 % (((v2 ^ (v2 / ((v2) | 1)))) | 1))
    for fi96 := 0; fi96 < 13; fi96++ {
        for fi97 := 0; fi97 < 14; fi97++ {
            if true {
                v90 = (v90 / (((-(fi97))) | 1))
                v98 := GenS7{F0: 124, F1: -141, F2: fi96}
                fmt.Println(v2, v90, v95)
                var v99 float32 = (3.14159 * ((((1e10 / 1e10) / (0.0 - 3.14159)) + 0.0) + 3.14159))
                var v100 int64 = 321
                _ = v100
                _ = v98
                _ = v99
            } else {
                var v101 bool = (float64(1e10) >= (1.5 + 1e10))
                var v102 int16 = v1
                v95 += (v95 + (v95 | v2))
                _ = v101
                _ = v102
            }
            _ = fi97
        }
        var v103 string = "hello"
        if len(v91) >= 2 {
            v91 = v91[1:len(v91)]
        }
        _ = fi96
        _ = v103
    }
    v104 := GenS15{F0: 0}
    _ = len(v91)
    if true {
        var v105 int16 = (v1 + v1)
        _ = v105
    } else {
        v106 := GenS7{F0: 474, F1: v1, F2: (v90 | v90)}
        _ = v106
    }
    fmt.Println(v3, v3, v1)
    var v107 uint = 318
    v108 := GenS4{F0: v90}
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v90:", v90)
    fmt.Println("slicelen v91:", len(v91), v91)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v104
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v90
    _ = v91
    _ = v95
}
