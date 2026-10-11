package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 int8; F2 int32 }
type GenS5 struct { F0 uint32; F1 uint; F2 bool }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 uint8; F1 int32 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int; F1 uint16 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS16 struct { F0 int; F1 bool; F2 uint8 }
func genFn0() int {
    v4 := []int{-914, 32767, 9223372036854775807, 32767}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, 127)
    if len(v4) > 0 {
        v4[0] = -792
    }
    var v5 int8 = 101
    _ = v4
    _ = v5
    return 298
}

func main() {
    var v1 int32 = 0
    v2 := []int{2147483647}
    var v3 float64 = (3.14159 / 1.0)
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int32 = (-(v1))
        switch (fi6 << ((191) & 63)) {
        case 128:
            v8 := GenS1{F0: (v7 << ((133) & 63)), F1: 42}
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    var v9 int64 = -289
    fmt.Println(v1, v3)
    _ = len(v2)
    if true {
        v10 := GenS1{F0: v1, F1: (v9 + (v9 - (v9 << ((749) & 63))))}
        v11 := GenS1{F0: 1, F1: v9}
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        switch "\n\t" {
        case "foo bar":
            v13 := GenS2{F0: ""}
            _ = v13
        case "":
            var v14 bool = (4294967295 == 179)
            _ = v14
        }
        var v15 int = -146
        if false {
            v3 -= -2.25
        } else {
            v16 := GenS1{F0: -2147483648, F1: v9}
            switch "a" {
            case "a":
                v2 = append(v2, (v15 / (((fi12 << ((18446744073709551615) & 63))) | 1)))
            case "hello":
                var v17 uint = 657
                _ = v17
            case "":
                switch "hello" {
                case "foo bar":
                    fmt.Println(fi12)
                default:
                    if (87 != 255) {
                        for fi18 := 0; fi18 < 6; fi18++ {
                            var v19 uint32 = 10
                            var v20 bool = (!true)
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    } else {
                        if len(v2) > 0 {
                            v2[0] = v15
                        }
                        switch "foo bar" {
                        case "":
                            fmt.Println(v9)
                        case "foo bar":
                            if len(v2) > 0 {
                                v2[0] = (v15 << ((42) & 63))
                            }
                        case "a":
                            v21 := GenS3{F0: 32767}
                            _ = v21
                        default:
                            fmt.Println(fi12)
                        }
                    }
                }
            default:
                for fi22 := 0; fi22 < 18; fi22++ {
                    var v23 bool = false
                    if (127 != 0) {
                        v15 = 560
                        v24 := GenS2{F0: "foo bar"}
                        v25 := GenS3{F0: 0}
                        if (v23 || true) {
                            if true {
                                v9 = (v9 << ((593) & 63))
                                var v26 uint16 = 65535
                                v2 = append(v2, (-(fi22)))
                                _ = v26
                            } else {
                                v27 := GenS4{F0: 0, F1: 127, F2: v1}
                                if v23 {
                                    if (((("" + ("foo bar" + "\n\t")) + "foo bar") != "") || v23) {
                                        v1 += v1
                                        if (147 > 255) {
                                            switch ("" + ("a" + ("a" + ""))) {
                                            case "\n\t":
                                                _ = len(v2)
                                                var v28 int16 = 32767
                                                _ = v28
                                            case "a":
                                                v29 := GenS5{F0: 65535, F1: 42, F2: (fi12 <= (fi12 / ((-2147483648) | 1)))}
                                                _ = v29
                                            default:
                                                v15 = 32768
                                            }
                                        }
                                        var v30 string = "foo bar"
                                        var v31 float32 = 0.1
                                        _ = v30
                                        _ = v31
                                    }
                                    if v23 {
                                        if (0.1 > 1e10) {
                                            if (!(v23 || (!(!(944 > v9))))) {
                                                if true {
                                                    switch "a" {
                                                    case "hello":
                                                        v9 = v9
                                                    default:
                                                        var v32 uint64 = 4294967295
                                                        _ = v32
                                                    }
                                                    var v33 string = ("a" + "a")
                                                    _ = v33
                                                }
                                                if (fi12 < 128) {
                                                    var v34 float64 = ((float64(float32(v3)) / (v3 * float64((1e10 - (0.0 / 3.14159))))) + ((v3 / float64(3.14159)) / v3))
                                                    var v35 uint32 = 2147483648
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    fmt.Println(v3, v1, v1)
                                                    switch "hello" {
                                                    case "hello":
                                                        v1 -= 810
                                                    case "foo bar":
                                                        if v23 {
                                                            if true {
                                                                switch fi22 {
                                                                case 128:
                                                                    var v36 int = fi12
                                                                    _ = v36
                                                                case 0:
                                                                    if (483 <= 255) {
                                                                        if v23 {
                                                                            var v37 int8 = 0
                                                                            _ = v37
                                                                        } else {
                                                                            switch (^(fi22)) {
                                                                            case -9223372036854775808:
                                                                                switch (fi22 & fi12) {
                                                                                case -71:
                                                                                    fmt.Println(fi12)
                                                                                case 745:
                                                                                    if v23 {
                                                                                        v15 -= fi12
                                                                                        var v38 int64 = (v9 >> ((2) & 63))
                                                                                        fmt.Println(fi12)
                                                                                        _ = v38
                                                                                    } else {
                                                                                        if v23 {
                                                                                            _ = len(v2)
                                                                                            v15 -= (v15 / (((v15 * (fi22 / ((v15) | 1)))) | 1))
                                                                                        }
                                                                                        var v39 uint = 65535
                                                                                        var v40 int = (fi12 << (((v39 - 0)) & 63))
                                                                                        _ = v39
                                                                                        _ = v40
                                                                                    }
                                                                                }
                                                                            case 2147483647:
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    fmt.Println(fi22, fi12)
                                                                                default:
                                                                                    var v41 float64 = (v3 - float64(0.0))
                                                                                    _ = v41
                                                                                }
                                                                            case 128:
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = fi12
                                                                                }
                                                                            }
                                                                            var v42 int32 = 42
                                                                            var v43 int = 538
                                                                            _ = v42
                                                                            _ = v43
                                                                        }
                                                                        fmt.Println(v1, v1, fi22)
                                                                    }
                                                                case -329:
                                                                    switch (^(fi12)) {
                                                                    case 4294967295:
                                                                        fmt.Println(v1, fi22, v23)
                                                                        fmt.Println(v9)
                                                                        switch (v15 & (fi12 / ((v15) | 1))) {
                                                                        case 128:
                                                                            var v44 float64 = v3
                                                                            _ = v44
                                                                        default:
                                                                            fmt.Println(v23, v1)
                                                                        }
                                                                        if false {
                                                                            v3 = v3
                                                                            var v45 uint64 = 18446744073709551615
                                                                            _ = v45
                                                                        } else {
                                                                            fmt.Println(v15, v1)
                                                                            v2 = append(v2, (fi12 << ((303) & 63)))
                                                                            var v46 string = "a"
                                                                            _ = v46
                                                                        }
                                                                    default:
                                                                        var v47 uint64 = 0
                                                                        _ = v47
                                                                    }
                                                                }
                                                                var v48 uint = 554
                                                                if v23 {
                                                                    if false {
                                                                        v9 = (v9 - (-(v9)))
                                                                    } else {
                                                                        if v23 {
                                                                            v23 = v23
                                                                        } else {
                                                                            fmt.Println(v1, v48, v15)
                                                                        }
                                                                        var v49 int32 = (v1 - (^(v1)))
                                                                        _ = v49
                                                                    }
                                                                    if (v23 && ((v23 || false) || (true && (fi12 != (-(fi22)))))) {
                                                                        _ = len(v2)
                                                                    }
                                                                } else {
                                                                    v50 := GenS6{F0: (v1 % ((v1) | 1)), F1: 612}
                                                                    v15 += fi22
                                                                    _ = v50
                                                                }
                                                                _ = v48
                                                            }
                                                            _ = len(v2)
                                                            var v51 uint = 18446744073709551615
                                                            _ = v51
                                                        } else {
                                                            switch v15 {
                                                            case 9223372036854775807:
                                                                v52 := GenS7{F0: 0}
                                                                _ = v52
                                                            case 42:
                                                                var v53 uint = 2
                                                                _ = v53
                                                            case -2147483648:
                                                                var v54 int8 = int8(v1)
                                                                _ = v54
                                                            }
                                                        }
                                                    default:
                                                        v3 -= (v3 - (v3 - v3))
                                                    }
                                                    _ = len(v2)
                                                }
                                                v55 := GenS8{F0: 2}
                                                _ = v55
                                            }
                                            var v56 int16 = 0
                                            v57 := GenS9{F0: 1}
                                            _ = v56
                                            _ = v57
                                        } else {
                                            fmt.Println(fi12, fi12, v1)
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = (fi22 - fi12)
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 605
                                        }
                                    }
                                    v58 := GenS10{F0: 42, F1: -36}
                                    _ = v58
                                } else {
                                    v9 = (^(v9))
                                }
                                _ = v27
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v59 uint = 77
                            var v60 bool = false
                            var v61 int16 = int16(v9)
                            _ = v59
                            _ = v60
                            _ = v61
                        } else {
                            switch v15 {
                            case -17:
                                fmt.Println(v1)
                            default:
                                if (199 == v1) {
                                    var v62 int8 = 1
                                    if (((!((((((((-2.25 - 1.5) / (v3 + -2.25)) + 3.14159) - 0.1) / v3) / v3) == float64((1.5 + 1.5))) || v23)) || v23) && false) {
                                        _ = len(v2)
                                        var v63 int = 9223372036854775807
                                        _ = v63
                                    }
                                    _ = v62
                                } else {
                                    fmt.Println(v23, v3)
                                }
                            }
                            var v64 uint8 = 42
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if (v64 != v64) {
                                if len(v2) > 0 {
                                    v2[0] = v15
                                }
                            } else {
                                v3 = ((-2.25 - (v3 / 0.1)) - float64((((float32(0.1) - (-2.25 * 1.5)) - (0.1 / 1.5)) - 0.1)))
                                v15 += v15
                            }
                            _ = v64
                        }
                        if (!("a" != (("foo bar" + "a") + "a"))) {
                            v65 := GenS11{F0: 10, F1: (-(v1))}
                            if ((v3 + 3.14159) <= (v3 * float64(float32((1.5 - v3))))) {
                                if v23 {
                                    if ((0 == -1) || v23) {
                                        var v66 int64 = 9223372036854775807
                                        v67 := GenS10{F0: 42, F1: int8(v1)}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v66
                                        _ = v67
                                    } else {
                                        var v68 string = (("" + "a") + "\n\t")
                                        _ = v68
                                    }
                                    v23 = false
                                    v3 -= float64(1.5)
                                } else {
                                    var v69 int8 = -26
                                    _ = v69
                                }
                                if v23 {
                                    _ = len(v2)
                                }
                                if ((fi12 / ((32768) | 1)) > fi12) {
                                    var v70 bool = false
                                    if v70 {
                                        var v71 int8 = 100
                                        switch fi22 {
                                        case 4294967295:
                                            switch genFn0() {
                                            case -9223372036854775808:
                                                var v72 int16 = -919
                                                _ = v72
                                            case 32767:
                                                v9 = v9
                                            default:
                                                v73 := GenS9{F0: 500}
                                                _ = v73
                                            }
                                        case 128:
                                            fmt.Println(v15)
                                        }
                                        switch 9223372036854775807 {
                                        case 32768:
                                            var v74 uint64 = 879
                                            _ = v74
                                        case 32767:
                                            if ((234 > 65535) || v23) {
                                                if len(v2) > 0 {
                                                    v2[0] = (fi12 % ((genFn0()) | 1))
                                                }
                                                v75 := GenS8{F0: 65535}
                                                v76 := GenS12{F0: v71}
                                                if true {
                                                    v2 = append(v2, genFn0())
                                                } else {
                                                    v2 = append(v2, 32767)
                                                    v23 = (("hello" + "hello") != "")
                                                }
                                                v2 = append(v2, genFn0())
                                                _ = v75
                                                _ = v76
                                            } else {
                                                fmt.Println(v70, v9, v23)
                                            }
                                        default:
                                            v77 := GenS13{F0: (v15 | genFn0()), F1: 393}
                                            _ = v77
                                        }
                                        _ = v71
                                    } else {
                                        var v78 int8 = -1
                                        _ = v78
                                    }
                                    _ = v70
                                }
                            } else {
                                var v79 uint32 = 269
                                v1 += v1
                                var v80 string = "\n\t"
                                _ = v79
                                _ = v80
                            }
                            if true {
                                var v81 float64 = v3
                                if len(v2) > 0 {
                                    v2[0] = 2147483647
                                }
                                _ = v81
                            }
                            _ = len(v2)
                            var v82 string = "hello"
                            v83 := GenS6{F0: v1, F1: 2}
                            _ = v65
                            _ = v82
                            _ = v83
                        }
                        _ = v24
                        _ = v25
                    } else {
                        var v84 uint32 = 2147483648
                        var v85 float64 = v3
                        switch (("foo bar" + ("\n\t" + "foo bar")) + "") {
                        case "foo bar":
                            var v86 int8 = 127
                            _ = v86
                        case "\n\t":
                            if v23 {
                                if len(v2) > 0 {
                                    v2[0] = genFn0()
                                }
                                v15 = v15
                            } else {
                                var v87 int32 = -2147483648
                                _ = v87
                            }
                        }
                        v84 = (v84 | v84)
                        v88 := GenS5{F0: v84, F1: 18446744073709551615, F2: true}
                        _ = len(v2)
                        var v89 int64 = v9
                        _ = v84
                        _ = v85
                        _ = v88
                        _ = v89
                    }
                    _ = fi22
                    _ = v23
                }
            }
            _ = v16
        }
        _ = fi12
        _ = v15
    }
    switch 32768 {
    case 1:
        if len(v2) > 0 {
            v2[0] = 0
        }
    case 2147483647:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v90 uint16 = 991
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v91 uint32 = 451
        _ = v91
    } else {
        v92 := GenS14{F0: 0}
        _ = v92
    }
    for fi93 := 0; fi93 < 19; fi93++ {
        v1 -= 0
        v94 := GenS8{F0: v90}
        _ = fi93
        _ = v94
    }
    var v95 int8 = 0
    v96 := GenS2{F0: ("\n\t" + "foo bar")}
    for fi97 := 0; fi97 < 13; fi97++ {
        var v98 uint64 = 4294967295
        if false {
            v99 := GenS15{F0: 2, F1: 9223372036854775807, F2: 2}
            for fi100 := 0; fi100 < 14; fi100++ {
                var v101 int64 = 415
                v101 = v9
                v1 = (v1 % (((v1 >> ((42) & 63))) | 1))
                var v102 int16 = -32768
                _ = fi100
                _ = v101
                _ = v102
            }
            _ = v99
        } else {
            var v103 float64 = 1e10
            if false {
                var v104 int32 = (v1 / (((v1 >> ((2) & 63))) | 1))
                var v105 int64 = v9
                switch "foo bar" {
                case "hello":
                    for fi106 := 0; fi106 < 10; fi106++ {
                        v104 -= (^(v1))
                        var v107 uint8 = 0
                        _ = fi106
                        _ = v107
                    }
                }
                _ = v104
                _ = v105
            } else {
                for fi108 := 0; fi108 < 6; fi108++ {
                    switch "a" {
                    case "":
                        v1 -= v1
                    }
                    fmt.Println(fi108, v103, v9)
                    _ = fi108
                }
                v90 += v90
            }
            if false {
                var v109 float32 = 0.0
                fmt.Println(fi97)
                var v110 uint8 = 64
                _ = v109
                _ = v110
            } else {
                var v111 uint16 = 255
                for fi112 := 0; fi112 < 16; fi112++ {
                    var v113 int64 = v9
                    _ = fi112
                    _ = v113
                }
                for fi114 := 0; fi114 < 12; fi114++ {
                    switch (fi114 + (fi97 ^ (fi114 << ((661) & 63)))) {
                    case 128:
                        if len(v2) > 0 {
                            v2[0] = -1
                        }
                    case 1:
                        if len(v2) > 0 {
                            v2[0] = (fi114 << ((0) & 63))
                        }
                    case -2147483648:
                        if false {
                            v115 := GenS9{F0: (v111 % (((v90 << ((18446744073709551615) & 63))) | 1))}
                            _ = v115
                        } else {
                            v1 += 752
                        }
                    }
                    _ = fi114
                }
                _ = v111
            }
            _ = v103
        }
        _ = fi97
        _ = v98
    }
    v116 := GenS16{F0: 9223372036854775807, F1: false, F2: 42}
    v117 := GenS9{F0: 1}
    v90 -= v90
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    fmt.Println("v90:", v90)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v9
    _ = v90
    _ = v95
    _ = v96
}
