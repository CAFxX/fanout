package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int16; F1 uint64; F2 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 int16 }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 bool; F1 uint; F2 string }
type GenS12 struct { F0 int32; F1 int }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint8; F1 bool }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint32; F1 string; F2 uint32 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 uint16; F1 int64 }
type GenS20 struct { F0 uint16; F1 bool; F2 int32 }
func genFn0(p0 uint, p1 int32, p2 int16) int16 {
    var v5 uint64 = 2
    v6 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p2
}

func main() {
    var v1 uint8 = 42
    var v2 float64 = float64(-2.25)
    v3 := make(map[string]int)
    var v4 uint32 = 0
    var v7 float64 = v2
    var v8 uint64 = 315
    if (!false) {
        var v9 float64 = float64(-2.25)
        _ = v9
    } else {
        if (49 != int8(v8)) {
            if ((("\n\t" + ("foo bar" + (((("\n\t" + "foo bar") + ("\n\t" + ("hello" + "a"))) + "a") + "\n\t"))) + ("a" + ("\n\t" + "foo bar"))) < "hello") {
                if true {
                    delete(v3, "hello")
                } else {
                    var v10 int32 = int32(v4)
                    _ = v10
                }
                for fi11 := 0; fi11 < 17; fi11++ {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        switch "\n\t" {
                        case "":
                            v4 -= v4
                        case "a":
                            fmt.Println(v7, fi11, v4)
                        case "foo bar":
                            fmt.Println(fi12)
                        default:
                            _ = len(v3)
                        }
                        fmt.Println(fi11)
                        if (82 <= v4) {
                            if true {
                                var v13 int32 = -108
                                v4 -= v4
                                switch "\n\t" {
                                case "hello":
                                    switch "" {
                                    case "":
                                        fmt.Println(v7, v7, v2)
                                    default:
                                        v14 := GenS2{F0: (^(v4))}
                                        var v15 int16 = 42
                                        _ = v14
                                        _ = v15
                                    }
                                case "":
                                    fmt.Println(v7, v13, v8)
                                }
                                if (true || true) {
                                    var v16 int16 = 1000
                                    _ = v16
                                } else {
                                    var v17 int16 = 1
                                    switch (fi11 + (^(fi12))) {
                                    case 4294967295:
                                        if (18446744073709551615 > (^(v8))) {
                                            var v18 uint64 = (^(v8))
                                            _ = v18
                                        } else {
                                            v19 := GenS3{F0: uint64(v1)}
                                            if true {
                                                var v20 uint8 = (v1 | v1)
                                                var v21 uint64 = (^(v8))
                                                _ = v20
                                                _ = v21
                                            } else {
                                                v22 := GenS2{F0: (v4 | (v4 & (v4 % ((793) | 1))))}
                                                var v23 int64 = -9223372036854775808
                                                if false {
                                                    v24 := GenS4{F0: 293, F1: v1, F2: 255}
                                                    v25 := GenS3{F0: (v8 % (((^(v8))) | 1))}
                                                    var v26 int = (fi12 / ((4294967295) | 1))
                                                    _ = v24
                                                    _ = v25
                                                    _ = v26
                                                } else {
                                                    v13 -= 38
                                                    v2 -= v2
                                                }
                                                _ = v22
                                                _ = v23
                                            }
                                            _ = v19
                                        }
                                    case -722:
                                        if true {
                                            var v27 float32 = float32(float64(((3.14159 / 1.5) * 0.1)))
                                            if false {
                                                v8 = v8
                                            } else {
                                                if (true && (4294967295 <= fi12)) {
                                                    if (1 <= 1) {
                                                        v17 -= -298
                                                        v28 := GenS5{F0: (v17 / ((v17) | 1)), F1: v8, F2: v13}
                                                        v7 = 3.14159
                                                        v29 := GenS1{F0: 1}
                                                        var v30 float64 = v2
                                                        switch fi12 {
                                                        case -2147483648:
                                                            v31 := GenS6{F0: 127}
                                                            _ = v31
                                                        case 38:
                                                            fmt.Println(v27, v4, fi11)
                                                        default:
                                                            var v32 bool = false
                                                            _ = v32
                                                        }
                                                        if false {
                                                            v8 += (v8 * (v8 + (v8 * v8)))
                                                            v27 = -2.25
                                                            switch ("hello" + "foo bar") {
                                                            case "a":
                                                                v7 = v30
                                                            }
                                                        }
                                                        _ = v28
                                                        _ = v29
                                                        _ = v30
                                                    } else {
                                                        v4 += 4294967295
                                                    }
                                                    if ((0.1 / (float64(v27) * v2)) == float64(v27)) {
                                                        if ((v4 - 583) >= v4) {
                                                            if (!true) {
                                                                fmt.Println(v1, v4, v7)
                                                                if (!false) {
                                                                    v33 := GenS6{F0: -1}
                                                                    _ = v33
                                                                }
                                                                if true {
                                                                    v13 -= (-(v13))
                                                                    var v34 uint16 = uint16(v1)
                                                                    _ = v34
                                                                }
                                                                if (v4 != v4) {
                                                                    v35 := []int{fi11, fi12, fi11}
                                                                    if len(v35) >= 2 {
                                                                        v35 = v35[1:len(v35)]
                                                                    }
                                                                    v17 = (v17 >> ((6) & 63))
                                                                    if (!true) {
                                                                        v36 := GenS1{F0: (v17 | v17)}
                                                                        switch (fi11 - (fi12 | 985)) {
                                                                        case -485:
                                                                            v37 := GenS2{F0: (v4 & (^(v4)))}
                                                                            _ = v37
                                                                        case 127:
                                                                            _ = len(v3)
                                                                        default:
                                                                            var v38 string = ((("foo bar" + ("hello" + "a")) + "\n\t") + "a")
                                                                            _ = v38
                                                                        }
                                                                        var v39 bool = true
                                                                        _ = v36
                                                                        _ = v39
                                                                    } else {
                                                                        fmt.Println(v4)
                                                                    }
                                                                    _ = v35
                                                                }
                                                                v40 := GenS7{F0: (v1 ^ 42), F1: 638}
                                                                _ = v40
                                                            }
                                                            var v41 uint8 = 2
                                                            var v42 float32 = (float32((v7 + v2)) / (v27 - 1.5))
                                                            _ = v41
                                                            _ = v42
                                                        } else {
                                                            var v43 bool = false
                                                            switch (fi11 >> ((536) & 63)) {
                                                            case -1:
                                                                if (genFn0(65535, v13, v17) <= v17) {
                                                                    v44 := GenS7{F0: v1, F1: genFn0(1, (v13 >> ((42) & 63)), (v17 - v17))}
                                                                    v1 = 255
                                                                    v45 := GenS7{F0: uint8(v4), F1: v17}
                                                                    v46 := []int{fi12, (-(fi12)), -380, (fi11 % ((fi11) | 1))}
                                                                    v3["foo bar"] = (fi11 - (fi12 << ((uint(v8)) & 63)))
                                                                    v47 := GenS8{F0: (239 > v1), F1: (v8 / (((v8 << ((494) & 63))) | 1))}
                                                                    v2 = ((v7 + v7) / v7)
                                                                    _ = v44
                                                                    _ = v45
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                            case 513:
                                                                var v48 bool = (85 != 65535)
                                                                _ = v48
                                                            case -2147483648:
                                                                var v49 uint16 = uint16(v4)
                                                                _ = v49
                                                            default:
                                                                v1 += (v1 & v1)
                                                            }
                                                            _ = v43
                                                        }
                                                        v27 += ((v27 * v27) * 0.1)
                                                        v50 := v3["\n\t"]
                                                        v51 := GenS5{F0: v17, F1: v8, F2: (^(v13))}
                                                        var v52 float64 = float64(v27)
                                                        _ = v50
                                                        _ = v51
                                                        _ = v52
                                                    } else {
                                                        switch ("hello" + "foo bar") {
                                                        case "foo bar":
                                                            if false {
                                                                switch "hello" {
                                                                case "":
                                                                    if false {
                                                                        if (1 != 922) {
                                                                            v53 := GenS9{F0: -103}
                                                                            v7 += (v7 / v2)
                                                                            v17 += v17
                                                                            _ = v53
                                                                        }
                                                                    } else {
                                                                        if (genFn0(4294967295, v13, (v17 >> ((18446744073709551615) & 63))) != v17) {
                                                                            var v54 int32 = (v13 * 2147483647)
                                                                            _ = v54
                                                                        }
                                                                        v55 := GenS10{F0: v4}
                                                                        _ = v55
                                                                    }
                                                                case "hello":
                                                                    fmt.Println(v27)
                                                                    var v56 int64 = -497
                                                                    switch ("hello" + (("foo bar" + "\n\t") + (("hello" + "hello") + "foo bar"))) {
                                                                    case "\n\t":
                                                                        v57 := []int{(fi12 >> ((65535) & 63))}
                                                                        _ = v57
                                                                    case "":
                                                                        fmt.Println(v17, fi12)
                                                                    case "foo bar":
                                                                        var v58 uint8 = (v1 - v1)
                                                                        fmt.Println(v7, v4)
                                                                        _ = v58
                                                                    default:
                                                                        fmt.Println(v8, v4)
                                                                    }
                                                                    _ = v56
                                                                }
                                                                v59 := GenS11{F0: ((true && (false || true)) || false), F1: 725, F2: ("hello" + ("\n\t" + "hello"))}
                                                                if (v1 >= v1) {
                                                                    if (1 != 100) {
                                                                        fmt.Println(v13, v17, fi11)
                                                                        switch (fi12 ^ fi11) {
                                                                        case 909:
                                                                            v17 = (v17 % (((^(v17))) | 1))
                                                                        case 1:
                                                                            var v60 int16 = (-(v17))
                                                                            _ = v60
                                                                        }
                                                                        var v61 int = -933
                                                                        _ = v61
                                                                    } else {
                                                                        fmt.Println(v7, v8)
                                                                        v62 := GenS12{F0: 0, F1: -2147483648}
                                                                        v63 := []int{-1, (^(fi11)), int(v13)}
                                                                        _ = v62
                                                                        _ = v63
                                                                    }
                                                                }
                                                                switch "" {
                                                                case "hello":
                                                                    v64 := GenS4{F0: 1, F1: (v1 >> ((65535) & 63)), F2: 42}
                                                                    _ = v64
                                                                }
                                                                if (((true || false) || true) && true) {
                                                                    fmt.Println(v4, v4)
                                                                } else {
                                                                    switch "" {
                                                                    case "a":
                                                                        v65 := v3["\n\t"]
                                                                        _ = v65
                                                                    case "hello":
                                                                        var v66 int16 = -32768
                                                                        _ = v66
                                                                    }
                                                                    v67 := []int{127}
                                                                    delete(v3, "a")
                                                                    _ = v67
                                                                }
                                                                _ = v59
                                                            }
                                                        case "":
                                                            v3[("foo bar" + "foo bar")] = (fi11 / ((fi12) | 1))
                                                            if false {
                                                                var v68 int16 = v17
                                                                v69 := GenS13{F0: 758}
                                                                v13 = (-(v13))
                                                                _ = v68
                                                                _ = v69
                                                            } else {
                                                                var v70 float64 = (v7 + (3.14159 / 0.1))
                                                                fmt.Println(v17)
                                                                _ = v70
                                                            }
                                                        default:
                                                            var v71 uint64 = 42
                                                            _ = v71
                                                        }
                                                        v27 = (3.14159 / v27)
                                                        v4 += (^(v4))
                                                    }
                                                    switch 42 {
                                                    case 0:
                                                        v4 += v4
                                                    case -554:
                                                        v4 = v4
                                                        v72 := []int{fi12, 42}
                                                        if true {
                                                            v13 = -2147483648
                                                            if true {
                                                                var v73 uint64 = 0
                                                                switch 0 {
                                                                case 128:
                                                                    v74 := v3[""]
                                                                    _ = v74
                                                                case -9223372036854775808:
                                                                    var v75 uint8 = (^(v1))
                                                                    _ = v75
                                                                default:
                                                                    var v76 uint64 = v73
                                                                    fmt.Println(v13)
                                                                    _ = v76
                                                                }
                                                                v77 := GenS14{F0: v1, F1: ((v73 % ((v73) | 1)) >= 403)}
                                                                _ = v73
                                                                _ = v77
                                                            } else {
                                                                delete(v3, ("foo bar" + ("a" + (("a" + ("a" + "hello")) + ("" + "")))))
                                                            }
                                                            fmt.Println(v2, v13, v1)
                                                            fmt.Println(v27)
                                                        } else {
                                                            v78 := GenS5{F0: genFn0(18446744073709551615, -892, (v17 ^ (v17 | v17))), F1: 18446744073709551615, F2: (v13 & 0)}
                                                            v79 := GenS15{F0: (v8 - v8)}
                                                            fmt.Println(fi12, v17, v13)
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                        _ = v72
                                                    case 128:
                                                        v80 := []int{(fi11 | (^(fi11))), fi12, (fi11 - 9223372036854775807), (fi12 / (((-(fi12))) | 1))}
                                                        _ = v80
                                                    }
                                                } else {
                                                    var v81 bool = ((("\n\t" + "foo bar") + ("" + ("" + "\n\t"))) <= ("\n\t" + ("" + "a")))
                                                    fmt.Println(v2, v81, v1)
                                                    var v82 uint64 = (v8 + v8)
                                                    _ = v81
                                                    _ = v82
                                                }
                                                v2 = v2
                                                var v83 int32 = (v13 + 86)
                                                if (!false) {
                                                    var v84 int = (fi11 | fi11)
                                                    v8 -= (^(v8))
                                                    _ = v84
                                                }
                                                var v85 uint64 = 18446744073709551615
                                                _ = v83
                                                _ = v85
                                            }
                                            _ = v27
                                        } else {
                                            var v86 uint16 = 81
                                            _ = v86
                                        }
                                        v87 := v3[("a" + ("\n\t" + ("\n\t" + "\n\t")))]
                                        switch (v87 ^ fi11) {
                                        case 32768:
                                            var v88 string = ("foo bar" + ("\n\t" + "a"))
                                            _ = v88
                                        case 0:
                                            v13 = v13
                                        default:
                                            v13 = (v13 % (((v13 & v13)) | 1))
                                        }
                                        v1 -= (v1 << ((2) & 63))
                                        var v89 uint32 = (v4 % ((1) | 1))
                                        v90 := GenS5{F0: genFn0(0, (v13 >> ((942) & 63)), (v17 << ((2) & 63))), F1: (v8 | v8), F2: int32(v17)}
                                        var v91 uint64 = 522
                                        v8 = v91
                                        _ = v87
                                        _ = v89
                                        _ = v90
                                        _ = v91
                                    }
                                    v92 := GenS16{F0: 1}
                                    v8 += v8
                                    _ = v17
                                    _ = v92
                                }
                                _ = v13
                            }
                            switch "a" {
                            case "hello":
                                switch "hello" {
                                case "":
                                    v93 := []int{32768, fi11, fi12, (fi12 / (((fi11 * (^(fi11)))) | 1))}
                                    _ = v93
                                case "\n\t":
                                    fmt.Println(v8, v1)
                                default:
                                    var v94 string = ("\n\t" + ("hello" + ""))
                                    _ = v94
                                }
                            case "":
                                _ = len(v3)
                            default:
                                delete(v3, "hello")
                            }
                            if true {
                                fmt.Println(v4, fi12, v4)
                            }
                        } else {
                            v8 += 4294967295
                            var v95 uint32 = v4
                            var v96 uint16 = 1
                            v97 := []int{}
                            v98 := GenS17{F0: v4, F1: "hello", F2: (v95 | v95)}
                            _ = v95
                            _ = v96
                            _ = v97
                            _ = v98
                        }
                        _ = fi12
                    }
                    _ = fi11
                }
                fmt.Println(v2)
            } else {
                if true {
                    var v99 uint64 = (v8 / ((694) | 1))
                    var v100 int64 = 880
                    _ = v100
                    _ = v99
                } else {
                    v101 := GenS5{F0: 32767, F1: (v8 - v8), F2: 2147483647}
                    var v102 uint16 = 483
                    _ = v101
                    _ = v102
                }
            }
        } else {
            v1 = (^(v1))
            for fi103 := 0; fi103 < 20; fi103++ {
                fmt.Println(v2, v8, v7)
                v104 := GenS18{F0: 65535}
                switch "" {
                case "hello":
                    v2 = v7
                case "a":
                    v105 := GenS16{F0: 0}
                    _ = v105
                default:
                    v4 = 132
                }
                _ = fi103
                _ = v104
            }
        }
        var v106 int64 = 42
        _ = v106
    }
    for fi107 := 0; fi107 < 11; fi107++ {
        v4 = (v4 - (v4 + 2))
        v7 -= float64(3.14159)
        var v108 uint8 = (^(v1))
        _ = fi107
        _ = v108
    }
    var v109 int = 127
    switch "hello" {
    case "a":
        fmt.Println(v1, v7, v8)
    case "hello":
        var v110 string = ("\n\t" + ("hello" + "\n\t"))
        _ = v110
    case "\n\t":
        for fi111 := 0; fi111 < 8; fi111++ {
            v112 := GenS19{F0: 542, F1: -1}
            _ = fi111
            _ = v112
        }
    default:
        fmt.Println(v109, v8)
    }
    v2 -= 3.14159
    fmt.Println(v4, v4)
    v113 := GenS20{F0: 65535, F1: false, F2: 42}
    v2 = v7
    v114 := v3[("a" + ("\n\t" + "hello"))]
    for fi115 := 0; fi115 < 20; fi115++ {
        for fi116 := 0; fi116 < 9; fi116++ {
            delete(v3, "\n\t")
            _ = fi116
        }
        v117 := GenS9{F0: -23}
        _ = fi115
        _ = v117
    }
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v109
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
