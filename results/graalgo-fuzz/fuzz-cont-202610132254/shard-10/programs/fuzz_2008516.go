package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 string; F2 int32 }
type GenS6 struct { F0 int; F1 uint; F2 int16 }
type GenS7 struct { F0 bool; F1 int64; F2 uint }
type GenS8 struct { F0 uint16; F1 uint8; F2 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 int8; F2 int16 }
type GenS11 struct { F0 int64; F1 uint16; F2 int64 }
type GenS12 struct { F0 uint64; F1 uint }
type GenS13 struct { F0 int16; F1 string; F2 uint16 }
type GenS14 struct { F0 int8; F1 uint64 }
type GenS15 struct { F0 uint32; F1 int8; F2 bool }
type GenS16 struct { F0 uint16 }
func genFn0() int16 {
    var v7 uint32 = 622
    var v8 int16 = 42
    if (4294967295 <= v7) {
        switch "a" {
        case "hello":
            v7 = v7
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 string = "a"
            var v11 uint16 = 1
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 bool = ((!false) && true)
            v14 := GenS1{F0: 128, F1: (577 == 2), F2: 379}
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 int64 = -381
            _ = fi16
            _ = v17
        }
        v7 -= 663
        _ = fi15
    }
    v7 += v7
    _ = v7
    _ = v8
    return -32768
}
func genFn1(p0 int, p1 int64, p2 int64) uint32 {
    fmt.Println(p0, p0)
    v18 := GenS2{F0: "foo bar", F1: 100}
    switch (^(p0)) {
    case 1:
        if (-1 != 42) {
            switch (("hello" + "\n\t") + "foo bar") {
            case "foo bar":
                p2 = 561
            case "\n\t":
                p0 = (p0 << ((776) & 63))
            }
        } else {
            if (!(p1 <= (p2 ^ (p1 & (p2 + (^(p1))))))) {
                v19 := GenS3{F0: 530, F1: 577}
                fmt.Println(p0, p1, p1)
                _ = v19
            } else {
                var v20 uint16 = 2
                p1 -= p2
                fmt.Println(p0, v20)
                _ = v20
            }
            fmt.Println(p1, p0)
            p0 += p0
        }
    case 2147483647:
        fmt.Println(p0, p0, p2)
    case 979:
        if false {
            v21 := GenS2{F0: "\n\t", F1: int8(p0)}
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 uint32 = 4294967295
    v4 := make(map[int]bool)
    var v5 uint32 = 2147483647
    v6 := []int{127, 32767}
    _ = len(v6)
    fmt.Println(v5, v3)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if true {
        var v22 int16 = 32767
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS4{F0: 1}
            if len(v6) > 0 {
                v6[0] = fi23
            }
            _ = fi23
            _ = v24
        }
        fmt.Println(v22, v22, v5)
        _ = v22
    } else {
        switch "\n\t" {
        case "foo bar":
            if true {
                v3 += (v5 * (v5 % (((^(v5))) | 1)))
            } else {
                fmt.Println(v5)
                var v25 uint32 = genFn1(-585, -538, 0)
                _ = v25
            }
        case "":
            if true {
                var v26 float32 = 0.0
                var v27 uint64 = 42
                if false {
                    v5 = v3
                }
                _ = v26
                _ = v27
            }
        default:
            fmt.Println(v3, v3)
        }
        var v28 int32 = 42
        if false {
            var v29 uint8 = 42
            _ = v29
        }
        _ = v28
    }
    var v30 int16 = -1
    fmt.Println(v5)
    if true {
        v6 = append(v6, 9223372036854775807)
        for fi31 := 0; fi31 < 15; fi31++ {
            var v32 string = "hello"
            _ = fi31
            _ = v32
        }
        if false {
            for fi33 := 0; fi33 < 2; fi33++ {
                v30 = -951
                var v34 uint = 42
                var v35 int32 = int32(v30)
                _ = fi33
                _ = v34
                _ = v35
            }
            var v36 uint64 = 454
            _ = v36
        }
    } else {
        v5 = 2147483647
        for fi37 := 0; fi37 < 13; fi37++ {
            v30 = v30
            for fi38 := 0; fi38 < 18; fi38++ {
                var v39 float32 = 1.5
                var v40 float32 = (1.5 * (v39 + v39))
                _ = fi38
                _ = v39
                _ = v40
            }
            v41 := GenS5{F0: false, F1: "\n\t", F2: 2147483647}
            _ = fi37
            _ = v41
        }
    }
    v3 -= genFn1(-1, -121, 907)
    v30 = (v30 ^ 1)
    for fi42 := 0; fi42 < 4; fi42++ {
        v5 = (v5 % (((v3 >> ((381) & 63))) | 1))
        if (!false) {
            for fi43 := 0; fi43 < 6; fi43++ {
                var v44 int8 = -1
                _ = fi43
                _ = v44
            }
            switch ((("hello" + "foo bar") + "\n\t") + ((("a" + ("hello" + "foo bar")) + "\n\t") + ("" + "a"))) {
            case "\n\t":
                switch ("a" + "\n\t") {
                case "foo bar":
                    switch fi42 {
                    case 173:
                        fmt.Println(v30)
                    default:
                        if (!(9223372036854775807 != 9223372036854775807)) {
                            for fi45 := 0; fi45 < 2; fi45++ {
                                v3 += (v3 ^ 78)
                                var v46 uint8 = 127
                                _ = fi45
                                _ = v46
                            }
                            v30 += 42
                            switch fi42 {
                            case 9223372036854775807:
                                v30 = genFn0()
                            default:
                                for fi47 := 0; fi47 < 2; fi47++ {
                                    v30 = 5
                                    v5 = v5
                                    if false {
                                        var v48 float32 = (3.14159 + 0.1)
                                        _ = v48
                                    } else {
                                        v49 := v2[fi47]
                                        switch ("a" + (("a" + (("hello" + "") + "")) + ("" + "foo bar"))) {
                                        case "":
                                            if (!true) {
                                                fmt.Println(v3)
                                                var v50 uint32 = (^(v5))
                                                _ = v50
                                            } else {
                                                v6 = append(v6, 772)
                                                switch (-(fi47)) {
                                                case -1:
                                                    switch -995 {
                                                    case 1:
                                                        if true {
                                                            switch (fi47 & (fi47 / ((fi47) | 1))) {
                                                            case -9223372036854775808:
                                                                var v51 float64 = (3.14159 / (3.14159 / 1e10))
                                                                _ = v51
                                                            case -928:
                                                                v52 := v4[fi42]
                                                                _ = v52
                                                            case 127:
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                            default:
                                                                if (!v49) {
                                                                    if v49 {
                                                                        v53 := v2[fi47]
                                                                        switch fi47 {
                                                                        case 127:
                                                                            var v54 uint = 2
                                                                            _ = v54
                                                                        }
                                                                        fmt.Println(v49, v5)
                                                                        _ = v53
                                                                    }
                                                                    v55 := GenS6{F0: (fi42 * 843), F1: 18446744073709551615, F2: v30}
                                                                    _ = v55
                                                                }
                                                            }
                                                            var v56 int64 = 860
                                                            if false {
                                                                if len(v6) > 0 {
                                                                    v6[0] = int(v56)
                                                                }
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                            }
                                                            _ = v56
                                                        }
                                                    }
                                                default:
                                                    var v57 uint8 = 128
                                                    _ = v57
                                                }
                                            }
                                            if true {
                                                var v58 int32 = 2147483647
                                                _ = v58
                                            } else {
                                                v3 = v5
                                                if v49 {
                                                    v3 -= v3
                                                    v5 = 0
                                                    v59 := GenS2{F0: "hello", F1: int8(v30)}
                                                    _ = v59
                                                }
                                            }
                                            if false {
                                                switch fi42 {
                                                case 128:
                                                    if ((v30 & (v30 - genFn0())) <= genFn0()) {
                                                        v60 := GenS7{F0: v49, F1: -118, F2: 42}
                                                        switch fi42 {
                                                        case -656:
                                                            if false {
                                                                var v61 int32 = 299
                                                                _ = v61
                                                            } else {
                                                                v5 = v3
                                                                var v62 int64 = 0
                                                                _ = v62
                                                            }
                                                        case 127:
                                                            if true {
                                                                switch fi42 {
                                                                case 1:
                                                                    var v63 int32 = 0
                                                                    _ = v63
                                                                case -2147483648:
                                                                    v64 := GenS8{F0: 0, F1: 127, F2: (fi47 * (fi47 >> ((1) & 63)))}
                                                                    _ = v64
                                                                case -110:
                                                                    v3 = (v3 >> ((2) & 63))
                                                                }
                                                            }
                                                        }
                                                        _ = v60
                                                    }
                                                    v65 := v1[v5]
                                                    switch int(v3) {
                                                    case 32767:
                                                        v5 -= (v5 ^ (v5 / (((v3 | (v3 % (((v3 % (((v3 * v3)) | 1))) | 1)))) | 1)))
                                                        v4[(fi47 >> ((18446744073709551615) & 63))] = (42 != -1)
                                                    case -532:
                                                        v49 = (-1 <= 32768)
                                                        v5 = v5
                                                    default:
                                                        var v66 int = 0
                                                        _ = v66
                                                    }
                                                    _ = v65
                                                case 1:
                                                    var v67 bool = ((v30 % (((v30 | genFn0())) | 1)) > (v30 % ((int16(fi47)) | 1)))
                                                    _ = v67
                                                default:
                                                    if (true || (v49 || (-1 != 42))) {
                                                        var v68 uint16 = 1
                                                        _ = v68
                                                    } else {
                                                        v69 := GenS9{F0: 255}
                                                        v3 = genFn1(0, 773, -1)
                                                        v70 := GenS10{F0: v49, F1: 1, F2: 892}
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    v71 := GenS2{F0: "foo bar", F1: -125}
                                                    v72 := GenS10{F0: (!v49), F1: 0, F2: v30}
                                                    v73 := GenS11{F0: -641, F1: 42, F2: 9223372036854775807}
                                                    _ = v71
                                                    _ = v72
                                                    _ = v73
                                                }
                                                v6 = append(v6, (fi47 + (fi42 << ((297) & 63))))
                                            } else {
                                                var v74 uint = 255
                                                if v49 {
                                                    fmt.Println(fi47, fi47, v5)
                                                    var v75 uint = (v74 | (v74 & v74))
                                                    _ = v75
                                                } else {
                                                    var v76 int16 = -1
                                                    v77 := GenS12{F0: 18446744073709551615, F1: v74}
                                                    var v78 int32 = -263
                                                    var v79 int64 = 781
                                                    v3 -= (v5 / ((0) | 1))
                                                    var v80 bool = false
                                                    fmt.Println(fi42)
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                }
                                                v3 = (^(v3))
                                                _ = v74
                                            }
                                            if v49 {
                                                fmt.Println(v5)
                                                var v81 float32 = ((-2.25 - (0.0 + 1.5)) * 0.1)
                                                v81 -= 0.0
                                                _ = v81
                                            }
                                        case "foo bar":
                                            var v82 int = fi47
                                            if false {
                                                switch (fi42 << ((336) & 63)) {
                                                case 219:
                                                    var v83 uint32 = v3
                                                    _ = v83
                                                case 42:
                                                    v30 += (v30 ^ v30)
                                                    if ((!v49) || (-2.25 <= 0.1)) {
                                                        v84 := GenS9{F0: 42}
                                                        v5 -= (v5 + 1)
                                                        if v49 {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            if v49 {
                                                                v5 = v3
                                                            }
                                                            var v85 float32 = (0.0 / (1.5 - 0.0))
                                                            _ = v85
                                                        }
                                                        _ = v84
                                                    } else {
                                                        var v86 uint32 = (^(v3))
                                                        v82 = (^(fi47))
                                                        v86 += v86
                                                        _ = v86
                                                    }
                                                case 2147483647:
                                                    v30 += v30
                                                }
                                                if false {
                                                    var v87 int32 = 305
                                                    _ = v87
                                                } else {
                                                    v5 += (v5 + (v5 * v5))
                                                    switch ("" + "a") {
                                                    case "a":
                                                        var v88 int64 = 881
                                                        _ = v88
                                                    case "\n\t":
                                                        v3 += (^(v5))
                                                    }
                                                }
                                                v89 := GenS5{F0: ((v30 ^ (v30 >> ((0) & 63))) == -1), F1: "", F2: 0}
                                                var v90 int16 = -615
                                                if (v49 && true) {
                                                    switch fi47 {
                                                    case -92:
                                                        v5 = 948
                                                    case 32768:
                                                        var v91 int64 = -354
                                                        _ = v91
                                                    case -1:
                                                        var v92 int16 = genFn0()
                                                        _ = v92
                                                    }
                                                    if len(v6) > 0 {
                                                        v6[0] = int(v90)
                                                    }
                                                }
                                                v82 += v82
                                                _ = len(v6)
                                                v5 = (v3 + genFn1(999, 1, 597))
                                                _ = v89
                                                _ = v90
                                            } else {
                                                v4[fi47] = (v49 && (true || v49))
                                            }
                                            v30 = v30
                                            _ = v82
                                        case "a":
                                            delete(v2, fi42)
                                        default:
                                            if (!v49) {
                                                v5 += (v5 ^ v3)
                                                v3 -= v5
                                            }
                                        }
                                        _ = v49
                                    }
                                    _ = fi47
                                }
                            }
                        } else {
                            for fi93 := 0; fi93 < 6; fi93++ {
                                var v94 int16 = genFn0()
                                _ = fi93
                                _ = v94
                            }
                            _ = len(v4)
                            fmt.Println(v5, v30, v3)
                        }
                    }
                default:
                    for fi95 := 0; fi95 < 13; fi95++ {
                        switch fi95 {
                        case 1:
                            var v96 uint8 = 128
                            _ = v96
                        case 912:
                            var v97 uint16 = 65535
                            _ = v97
                        case 127:
                            var v98 uint8 = 255
                            _ = v98
                        default:
                            if (fi42 == (fi42 * fi95)) {
                                v99 := GenS13{F0: v30, F1: ("foo bar" + (("a" + ("a" + ((("a" + "\n\t") + ("foo bar" + ("\n\t" + (("foo bar" + ("" + "\n\t")) + "a")))) + (("hello" + "") + "")))) + "")), F2: 252}
                                v100 := GenS14{F0: -1, F1: 468}
                                if (!(!true)) {
                                    v5 = (v5 >> ((634) & 63))
                                    fmt.Println(fi42, fi95, v3)
                                    v101 := GenS15{F0: genFn1(fi95, 9223372036854775807, 42), F1: 100, F2: (!(!true))}
                                    _ = v101
                                }
                                _ = v100
                                _ = v99
                            } else {
                                v5 -= genFn1(fi42, 841, 1)
                                var v102 int = (fi42 | -9223372036854775808)
                                var v103 uint16 = 630
                                v103 = 2
                                v104 := v1[uint32(v103)]
                                v103 += v103
                                _ = v102
                                _ = v103
                                _ = v104
                            }
                        }
                        _ = fi95
                    }
                }
            default:
                v30 = v30
            }
        } else {
            var v105 int64 = 753
            fmt.Println(v105)
            _ = v105
        }
        for fi106 := 0; fi106 < 2; fi106++ {
            var v107 uint8 = 2
            _ = fi106
            _ = v107
        }
        _ = fi42
    }
    v108 := GenS16{F0: 42}
    var v109 bool = (!(3.14159 == (0.1 + 0.0)))
    fmt.Println(v3)
    v110 := GenS4{F0: 103}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v108
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
