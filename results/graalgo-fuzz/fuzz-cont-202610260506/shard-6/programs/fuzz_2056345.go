package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 bool; F2 int64 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 string; F1 int16; F2 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16; F1 uint }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int; F1 int }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int; F1 int64; F2 int }
type GenS15 struct { F0 int16; F1 uint8 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int16; F1 string; F2 int64 }
func genFn0(p0 int) uint64 {
    if true {
        for fi5 := 0; fi5 < 1; fi5++ {
            fmt.Println(p0, p0)
            v6 := GenS1{F0: 2}
            _ = fi5
            _ = v6
        }
        if (false || (false || true)) {
            v7 := []int{-1, p0, 9223372036854775807, 32768}
            _ = v7
        } else {
            if true {
                var v8 int16 = 1000
                v9 := GenS2{F0: 7, F1: (18446744073709551615 >= 486), F2: 42}
                v10 := []int{(p0 | (p0 & (^(p0)))), (^(p0))}
                _ = v10
                _ = v8
                _ = v9
            } else {
                if true {
                    p0 = (p0 / ((p0) | 1))
                    for fi11 := 0; fi11 < 13; fi11++ {
                        v12 := []int{fi11}
                        for fi13 := 0; fi13 < 12; fi13++ {
                            if len(v12) > 0 {
                                v12[0] = 982
                            }
                            p0 = fi11
                            p0 = p0
                            v14 := GenS1{F0: 387}
                            _ = fi13
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 18; fi15++ {
                            if len(v12) > 0 {
                                v12[0] = (fi11 + 441)
                            }
                            switch (fi15 >> ((255) & 63)) {
                            case 32767:
                                var v16 uint64 = 0
                                _ = v16
                            case 9223372036854775807:
                                p0 = -9223372036854775808
                            case -259:
                                v17 := GenS3{F0: 85, F1: 1000, F2: 518}
                                _ = v17
                            }
                            _ = len(v12)
                            _ = fi15
                        }
                        _ = fi11
                        _ = v12
                    }
                    var v18 int8 = 81
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                p0 += p0
                var v20 float32 = (0.1 * (0.1 - (-2.25 / 1.5)))
                for fi21 := 0; fi21 < 1; fi21++ {
                    var v22 int8 = 42
                    v23 := GenS4{F0: 42, F1: 1}
                    v24 := GenS4{F0: 2, F1: -32768}
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi19
                _ = v20
            }
            v25 := GenS5{F0: 128}
            _ = v25
        }
    }
    v26 := GenS6{F0: 2147483647, F1: -104}
    _ = p0
    _ = v26
    return 13
}
func genFn1() uint8 {
    v27 := GenS7{F0: 1, F1: "\n\t"}
    var v28 int64 = -1
    fmt.Println(v28)
    if false {
        v29 := []int{32767, -862, -2147483648}
        v30 := GenS8{F0: ("hello" + ("\n\t" + "hello")), F1: 1, F2: -62}
        v28 -= v28
        _ = v29
        _ = v30
    } else {
        var v31 int = 127
        var v32 uint64 = 42
        _ = v31
        _ = v32
    }
    v33 := []int{32768, 128}
    _ = v27
    _ = v28
    _ = v33
    return 0
}

func main() {
    var v1 int = -115
    v2 := []string{("" + "\n\t"), ("" + "hello"), ("foo bar" + "a"), "foo bar"}
    v3 := make(map[uint32]string)
    var v4 uint8 = 255
    if true {
        for fi34 := 0; fi34 < 15; fi34++ {
            delete(v3, 65535)
            _ = fi34
        }
    }
    v4 -= genFn1()
    v1 += (v1 ^ v1)
    fmt.Println(v4, v4)
    var v35 int32 = 0
    v36 := v3[959]
    v3[881] = (v36 + v36)
    for fi37 := 0; fi37 < 3; fi37++ {
        if false {
            var v38 int64 = -9223372036854775808
            for fi39 := 0; fi39 < 6; fi39++ {
                switch v36 {
                case "":
                    if true {
                        if true {
                            if len(v2) > 0 {
                                v2[0] = ((v36 + (v36 + "foo bar")) + (v36 + v36))
                            }
                            var v40 float64 = 0.1
                            v41 := GenS2{F0: (v4 - v4), F1: false, F2: (^(v38))}
                            _ = v40
                            _ = v41
                        }
                    }
                case "hello":
                    if false {
                        v42 := GenS7{F0: -293, F1: v36}
                        _ = v42
                    } else {
                        v43 := GenS9{F0: 462}
                        if (!false) {
                            var v44 float64 = 1.5
                            v3[857] = "hello"
                            _ = v44
                        } else {
                            var v45 uint8 = v4
                            v46 := v3[4294967295]
                            var v47 float32 = float32(0.0)
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        v3[574] = "\n\t"
                        v38 -= v38
                        _ = v43
                    }
                case "a":
                    v1 = 4294967295
                }
                v38 += -1
                _ = fi39
            }
            _ = v38
        } else {
            var v48 int = 668
            var v49 int = -103
            _ = v48
            _ = v49
        }
        var v50 uint = 42
        v51 := GenS7{F0: 9223372036854775807, F1: (v36 + ("hello" + "hello"))}
        _ = fi37
        _ = v50
        _ = v51
    }
    switch ("a" + (v36 + v36)) {
    case "":
        if true {
            var v52 int8 = -128
            _ = v52
        } else {
            switch v36 {
            case "a":
                fmt.Println(v1)
            case "hello":
                v1 = (v1 / (((v1 % ((766) | 1))) | 1))
            default:
                switch v1 {
                case 32768:
                    v53 := GenS10{F0: -32768, F1: 352}
                    _ = v53
                case 32767:
                    if true {
                        for fi54 := 0; fi54 < 1; fi54++ {
                            switch v1 {
                            case 4294967295:
                                var v55 int = (v1 << ((65535) & 63))
                                _ = v55
                            }
                            v36 = "hello"
                            var v56 uint = 4294967295
                            _ = fi54
                            _ = v56
                        }
                        if (float32(((0.1 * 1.5) - (float64(1e10) + float64((1e10 * 0.1))))) >= 1.5) {
                            fmt.Println(v1, v35)
                        } else {
                            var v57 int32 = v35
                            v58 := GenS11{F0: true}
                            fmt.Println(v57, v4, v35)
                            _ = v57
                            _ = v58
                        }
                    }
                }
            }
        }
    }
    for fi59 := 0; fi59 < 16; fi59++ {
        v2 = append(v2, v36)
        var v60 bool = false
        for fi61 := 0; fi61 < 3; fi61++ {
            fmt.Println(v1, v36)
            var v62 int64 = 9223372036854775807
            switch "a" {
            case "foo bar":
                fmt.Println(v60, v1)
            case "\n\t":
                _ = len(v2)
            default:
                v63 := GenS7{F0: 9223372036854775807, F1: ""}
                v64 := GenS12{F0: (fi61 / ((32767) | 1)), F1: (v1 * (fi61 & 113))}
                switch v1 {
                case -2147483648:
                    v65 := GenS8{F0: v36, F1: 31, F2: -128}
                    _ = v65
                default:
                    var v66 int = 2147483647
                    var v67 int = int(v35)
                    v66 += v66
                    _ = v66
                    _ = v67
                }
                _ = v63
                _ = v64
            }
            v4 -= 2
            _ = fi61
            _ = v62
        }
        _ = fi59
        _ = v60
    }
    v4 = (v4 << ((255) & 63))
    fmt.Println(v36, v4, v4)
    if false {
        switch v1 {
        case 42:
            v4 = genFn1()
        }
        for fi68 := 0; fi68 < 6; fi68++ {
            if (557 < 42) {
                if ((((2 <= 42) || true) || false) && ((false && true) && true)) {
                    fmt.Println(v4)
                    var v69 uint64 = 1
                    _ = v69
                }
                for fi70 := 0; fi70 < 5; fi70++ {
                    if (true || false) {
                        var v71 float64 = 3.14159
                        var v72 bool = true
                        if false {
                            fmt.Println(v1, v1, v72)
                            v3[65535] = "foo bar"
                            if v72 {
                                var v73 uint64 = 513
                                var v74 int = fi68
                                var v75 uint32 = 65535
                                _ = v73
                                _ = v74
                                _ = v75
                            }
                        } else {
                            fmt.Println(v36, fi68)
                            v76 := GenS5{F0: (v4 & v4)}
                            var v77 int32 = v35
                            switch (v1 ^ fi68) {
                            case 312:
                                var v78 float32 = (-2.25 / 1.5)
                                _ = v78
                            case 1:
                                v79 := GenS10{F0: 0, F1: 0}
                                _ = v79
                            default:
                                if false {
                                    v80 := GenS4{F0: 18446744073709551615, F1: 0}
                                    if v72 {
                                        v81 := GenS7{F0: -419, F1: "a"}
                                        var v82 int32 = (v35 | (-(v35)))
                                        _ = v81
                                        _ = v82
                                    } else {
                                        if false {
                                            var v83 uint = 985
                                            var v84 float32 = float32((float64(1e10) + 3.14159))
                                            _ = v83
                                            _ = v84
                                        }
                                        if true {
                                            if v72 {
                                                delete(v3, 4294967295)
                                                var v85 float64 = (v71 * ((float64(3.14159) / v71) * (v71 + v71)))
                                                _ = v85
                                            } else {
                                                if v72 {
                                                    fmt.Println(fi68, v71)
                                                    v35 -= (v35 >> ((997) & 63))
                                                }
                                            }
                                            v2 = append(v2, v36)
                                        }
                                    }
                                    _ = v80
                                } else {
                                    if (625 < 1) {
                                        v86 := v3[674]
                                        _ = v86
                                    } else {
                                        switch (fi70 % ((fi68) | 1)) {
                                        case 9223372036854775807:
                                            if v72 {
                                                switch (v36 + v36) {
                                                case "\n\t":
                                                    var v87 int = v1
                                                    _ = v87
                                                }
                                                fmt.Println(v1)
                                                if ((^(v1)) > (fi70 & fi68)) {
                                                    var v88 float64 = (float64(1.5) * v71)
                                                    if (true || (!v72)) {
                                                        switch ((("\n\t" + v36) + v36) + v36) {
                                                        case "hello":
                                                            var v89 uint8 = v4
                                                            _ = v89
                                                        }
                                                        var v90 int32 = 2147483647
                                                        _ = len(v3)
                                                        if ((v35 < v77) || (v88 != (v88 / v88))) {
                                                            fmt.Println(v1)
                                                            if len(v2) > 0 {
                                                                v2[0] = v36
                                                            }
                                                            var v91 int32 = v77
                                                            _ = v91
                                                        } else {
                                                            v36 = "a"
                                                            if (v72 && v72) {
                                                                v3[0] = (((v36 + "foo bar") + v36) + ((v36 + (v36 + v36)) + (v36 + "")))
                                                                var v92 int64 = 364
                                                                _ = v92
                                                            } else {
                                                                v36 = v36
                                                            }
                                                        }
                                                        _ = v90
                                                    } else {
                                                        v93 := GenS13{F0: true}
                                                        if len(v2) > 0 {
                                                            v2[0] = v36
                                                        }
                                                        _ = v93
                                                    }
                                                    switch ("\n\t" + "") {
                                                    case "a":
                                                        v77 = v77
                                                    }
                                                    _ = v88
                                                } else {
                                                    var v94 uint16 = 65535
                                                    if v72 {
                                                        v94 += (v94 / (((v94 ^ v94)) | 1))
                                                        v95 := GenS14{F0: 4294967295, F1: -1, F2: fi68}
                                                        _ = v95
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = (v36 + v36)
                                                        }
                                                    }
                                                    v72 = (v72 && (!false))
                                                    _ = v94
                                                }
                                            }
                                        }
                                    }
                                    v35 += v35
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                            _ = v76
                            _ = v77
                        }
                        _ = v71
                        _ = v72
                    } else {
                        var v96 uint = 18446744073709551615
                        v97 := GenS4{F0: 13, F1: -32768}
                        var v98 float32 = (0.1 / 1.0)
                        switch v36 {
                        case "hello":
                            v99 := GenS1{F0: 2}
                            _ = v99
                        case "":
                            v100 := GenS15{F0: 1, F1: (v4 & v4)}
                            _ = v100
                        case "foo bar":
                            v96 = v96
                        default:
                            v35 = v35
                            var v101 float64 = 1.5
                            _ = v101
                        }
                        delete(v3, 2147483648)
                        if (false && true) {
                            if true {
                                var v102 int64 = 9223372036854775807
                                var v103 float64 = (1.5 * 3.14159)
                                v36 = (v36 + v36)
                                if true {
                                    v96 = (v96 / ((v96) | 1))
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v104 := GenS16{F0: -483}
                                    _ = v104
                                }
                                v105 := GenS17{F0: 0, F1: v36, F2: int64(v96)}
                                var v106 uint = (v96 * v96)
                                _ = v102
                                _ = v103
                                _ = v105
                                _ = v106
                            }
                            v96 = v96
                            v2 = append(v2, v36)
                        } else {
                            if false {
                                fmt.Println(v98, v35)
                                v98 += v98
                                if len(v2) > 0 {
                                    v2[0] = (v36 + (("\n\t" + "\n\t") + (("hello" + (v36 + "a")) + (v36 + v36))))
                                }
                            }
                            v35 = v35
                        }
                        var v107 float32 = v98
                        _ = v107
                        _ = v96
                        _ = v97
                        _ = v98
                    }
                    v4 += (v4 & (v4 / ((v4) | 1)))
                    fmt.Println(v35, v35)
                    _ = fi70
                }
                var v108 int64 = 42
                _ = v108
            } else {
                v4 = genFn1()
                v4 -= 202
            }
            fmt.Println(v36, fi68)
            for fi109 := 0; fi109 < 2; fi109++ {
                fmt.Println(fi109, v1, fi68)
                _ = fi109
            }
            _ = fi68
        }
    } else {
        var v110 uint32 = 2147483647
        v1 -= v1
        _ = v110
    }
    if (42 <= 42) {
        fmt.Println(v36)
    } else {
        v35 += (v35 % (((v35 + 630)) | 1))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
