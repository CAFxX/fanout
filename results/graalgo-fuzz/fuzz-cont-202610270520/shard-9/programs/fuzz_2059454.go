package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int8; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 string; F2 uint32 }
type GenS8 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS9 struct { F0 bool }
func genFn0(p0 int32) bool {
    p0 -= p0
    fmt.Println(p0)
    v6 := GenS1{F0: 55, F1: -69, F2: 26}
    if (!(-1 < 14)) {
        fmt.Println(p0)
        p0 -= (p0 - 2147483647)
    } else {
        v7 := GenS1{F0: 1, F1: 1, F2: 128}
        v8 := []int{2147483647}
        switch 4294967295 {
        case -9223372036854775808:
            var v9 uint16 = 661
            _ = v9
        case 0:
            var v10 int8 = 0
            _ = v10
        case 776:
            switch 128 {
            case 0:
                p0 -= (p0 * p0)
            case 108:
                fmt.Println(p0)
            }
        }
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v6
    return (2 != 1)
}

func main() {
    var v1 uint64 = 42
    var v2 uint16 = 114
    var v3 string = "a"
    var v4 float32 = 1.5
    var v5 uint8 = 0
    v11 := GenS1{F0: 100, F1: -128, F2: v5}
    v12 := GenS1{F0: 1, F1: 19, F2: 42}
    if (0.1 >= (float64(v4) / 1.0)) {
        v13 := GenS1{F0: 1, F1: -1, F2: (v5 >> ((65535) & 63))}
        switch (("" + "a") + v3) {
        case "\n\t":
            fmt.Println(v2, v4, v3)
        case "":
            v3 = "foo bar"
        }
        _ = v13
    }
    if (42 > -2147483648) {
        v14 := []int{400, 4294967295}
        _ = v14
    } else {
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 uint64 = 1
            var v17 int = fi15
            _ = fi15
            _ = v16
            _ = v17
        }
        if ((v2 << ((4294967295) & 63)) > v2) {
            var v18 float64 = 3.14159
            fmt.Println(v2)
            switch "\n\t" {
            case "a":
                var v19 int8 = -1
                _ = v19
            default:
                var v20 uint64 = 2
                _ = v20
            }
            _ = v18
        } else {
            for fi21 := 0; fi21 < 9; fi21++ {
                v22 := []int{4294967295, -878, fi21}
                v2 = v2
                v1 = (^(v1))
                _ = fi21
                _ = v22
            }
            v23 := GenS1{F0: -128, F1: 0, F2: v5}
            _ = v23
        }
    }
    var v24 uint32 = uint32(v5)
    v5 = v5
    if false {
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 uint64 = (^(v1))
            for fi27 := 0; fi27 < 20; fi27++ {
                if false {
                    var v28 uint16 = v2
                    var v29 int8 = 100
                    if (v2 == 42) {
                        v30 := GenS1{F0: v29, F1: (^(v29)), F2: 124}
                        if (!(!((v28 * v2) > v2))) {
                            if true {
                                fmt.Println(v1)
                                v31 := GenS2{F0: 1000, F1: (v1 & (^(v26))), F2: (v2 * 1)}
                                _ = v31
                            } else {
                                switch (v3 + "foo bar") {
                                case "":
                                    if (!(!true)) {
                                        switch (v3 + v3) {
                                        case "":
                                            fmt.Println(v28, v3, v28)
                                        case "foo bar":
                                            v3 = ""
                                        case "a":
                                            var v32 int8 = (v29 & 1)
                                            _ = v32
                                        }
                                        fmt.Println(v1)
                                    } else {
                                        var v33 int = (fi27 * fi27)
                                        var v34 bool = (true && true)
                                        var v35 int8 = (-(v29))
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                default:
                                    if ((^(v29)) < v29) {
                                        v36 := []int{(fi27 * (^(fi25))), fi27, fi25}
                                        var v37 uint32 = v24
                                        _ = v36
                                        _ = v37
                                    } else {
                                        v1 += v1
                                    }
                                }
                            }
                            switch v3 {
                            case "a":
                                fmt.Println(fi27, v28)
                            case "hello":
                                switch "foo bar" {
                                case "a":
                                    v5 = v5
                                    if ((fi25 % (((fi27 * (fi25 | int(v29)))) | 1)) != (fi25 | fi25)) {
                                        v24 -= v24
                                        if ((v26 + (^(v1))) >= v26) {
                                            var v38 uint = 482
                                            v28 += (^(v2))
                                            fmt.Println(v5, fi27, fi27)
                                            _ = v38
                                        } else {
                                            if (-2147483648 == -2147483648) {
                                                var v39 string = "foo bar"
                                                v40 := GenS2{F0: -1, F1: v1, F2: (v28 % (((v28 | 2)) | 1))}
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                        v41 := []int{}
                                        _ = v41
                                    } else {
                                        v42 := GenS1{F0: (v29 | v29), F1: v29, F2: 221}
                                        var v43 int64 = 9223372036854775807
                                        _ = v42
                                        _ = v43
                                    }
                                    v44 := []int{}
                                    var v45 int8 = v29
                                    _ = v44
                                    _ = v45
                                case "hello":
                                    v46 := []int{(fi27 * 356), (fi27 >> ((65535) & 63))}
                                    _ = v46
                                }
                            case "foo bar":
                                var v47 int8 = -100
                                _ = v47
                            }
                            v26 = v26
                            var v48 int32 = 2147483647
                            v49 := GenS3{F0: v1, F1: v29, F2: 0}
                            _ = v48
                            _ = v49
                        } else {
                            v28 -= (v2 * 615)
                            v28 = (v28 >> ((1) & 63))
                        }
                        _ = v30
                    } else {
                        var v50 uint8 = v5
                        _ = v50
                    }
                    v51 := []int{(fi27 << ((42) & 63)), 2147483647}
                    _ = v28
                    _ = v29
                    _ = v51
                } else {
                    var v52 int32 = -237
                    v1 -= uint64(v5)
                    var v53 uint32 = 42
                    v54 := []int{}
                    v55 := GenS1{F0: 100, F1: 127, F2: (v5 >> ((719) & 63))}
                    _ = len(v54)
                    if (!false) {
                        if (42 == (v2 >> ((4294967295) & 63))) {
                            var v56 int16 = 42
                            _ = v56
                        }
                        v2 = 42
                        v4 += 3.14159
                    } else {
                        var v57 string = (v3 + v3)
                        fmt.Println(v26)
                        _ = v57
                    }
                    _ = v52
                    _ = v53
                    _ = v54
                    _ = v55
                }
                var v58 uint16 = v2
                v59 := []int{}
                _ = fi27
                _ = v58
                _ = v59
            }
            var v60 bool = true
            _ = fi25
            _ = v26
            _ = v60
        }
        v61 := []int{}
        _ = v61
    } else {
        if (int8(v1) != -1) {
            v5 -= v5
        }
        for fi62 := 0; fi62 < 11; fi62++ {
            v63 := GenS4{F0: (v5 & v5)}
            switch ("foo bar" + v3) {
            case "":
                var v64 float32 = v4
                _ = v64
            }
            _ = fi62
            _ = v63
        }
    }
    if (18446744073709551615 == 2) {
        var v65 int8 = 0
        var v66 string = v3
        v4 = float32(0.1)
        _ = v65
        _ = v66
    } else {
        var v67 uint64 = 1
        _ = v67
    }
    for fi68 := 0; fi68 < 16; fi68++ {
        v1 -= 196
        switch (((v3 + (("a" + ("hello" + ((("\n\t" + v3) + v3) + v3))) + "\n\t")) + v3) + v3) {
        case "\n\t":
            for fi69 := 0; fi69 < 8; fi69++ {
                v70 := GenS1{F0: 1, F1: int8(v2), F2: (^(v5))}
                if (!false) {
                    v3 = v3
                    v71 := GenS1{F0: 42, F1: 100, F2: (v5 % ((v5) | 1))}
                    var v72 float64 = 1e10
                    _ = v71
                    _ = v72
                } else {
                    switch fi69 {
                    case -9223372036854775808:
                        var v73 int = (fi69 >> ((2) & 63))
                        _ = v73
                    case -231:
                        switch ((v3 + v3) + "") {
                        case "hello":
                            switch (("\n\t" + v3) + (((("hello" + "\n\t") + v3) + ("foo bar" + "foo bar")) + v3)) {
                            case "":
                                switch fi68 {
                                case 155:
                                    v74 := []int{fi69, (fi69 * (fi69 + fi68)), fi68, fi69}
                                    _ = v74
                                }
                            case "foo bar":
                                switch (fi69 % (((fi68 ^ (fi68 | fi68))) | 1)) {
                                case 815:
                                    if false {
                                        v1 = 688
                                        v1 -= v1
                                        if (((42 > 32767) || false) || true) {
                                            var v75 bool = true
                                            var v76 int = (fi69 - (fi69 ^ (fi69 - int(v1))))
                                            v5 -= 2
                                            _ = v75
                                            _ = v76
                                        } else {
                                            v77 := []int{(^(fi68)), (fi68 & (fi68 >> ((278) & 63))), (^(fi69))}
                                            v78 := GenS5{F0: v3, F1: v3}
                                            switch 279 {
                                            case 1:
                                                if (v5 >= 254) {
                                                    v1 = v1
                                                } else {
                                                    v1 = (v1 >> ((18446744073709551615) & 63))
                                                    _ = len(v77)
                                                }
                                            default:
                                                v79 := GenS6{F0: -1}
                                                _ = v79
                                            }
                                            var v80 float64 = 1e10
                                            _ = v77
                                            _ = v78
                                            _ = v80
                                        }
                                    } else {
                                        fmt.Println(fi69, v1)
                                        v81 := GenS4{F0: v5}
                                        v3 = v3
                                        if (!((0.0 != (float32(1.5) - (float32(0.1) * v4))) && false)) {
                                            if ((32767 > -32768) || false) {
                                                if true {
                                                    v82 := GenS7{F0: 42, F1: v3, F2: v24}
                                                    _ = v82
                                                } else {
                                                    v83 := GenS8{F0: (v24 * (v24 + v24)), F1: 255, F2: (v1 * (v1 | (v1 | v1)))}
                                                    fmt.Println(fi68, v24)
                                                    switch fi69 {
                                                    case -213:
                                                        v84 := []int{4294967295, (-(fi69))}
                                                        _ = v84
                                                    case -1:
                                                        v24 -= 2147483648
                                                        switch (("a" + ((v3 + v3) + ("" + (((v3 + v3) + ((("a" + v3) + ((v3 + v3) + ("\n\t" + v3))) + "foo bar")) + "\n\t")))) + v3) {
                                                        case "foo bar":
                                                            v2 -= v2
                                                        case "hello":
                                                            v85 := []int{926}
                                                            _ = v85
                                                        case "":
                                                            var v86 uint32 = (v24 >> ((1) & 63))
                                                            _ = v86
                                                        }
                                                        if false {
                                                            v87 := GenS5{F0: v3, F1: "foo bar"}
                                                            var v88 float32 = v4
                                                            v89 := []int{}
                                                            _ = v87
                                                            _ = v88
                                                            _ = v89
                                                        } else {
                                                            v24 = 653
                                                            if true {
                                                                var v90 uint64 = v1
                                                                v4 = float32(3.14159)
                                                                switch (v3 + "foo bar") {
                                                                case "foo bar":
                                                                    var v91 int32 = -2147483648
                                                                    _ = v91
                                                                case "\n\t":
                                                                    v92 := GenS8{F0: v24, F1: (v5 + (v5 | 128)), F2: (^(v1))}
                                                                    _ = v92
                                                                case "hello":
                                                                    if true {
                                                                        v93 := GenS9{F0: (4294967295 > fi69)}
                                                                        var v94 float32 = 3.14159
                                                                        v95 := []int{}
                                                                        _ = v93
                                                                        _ = v94
                                                                        _ = v95
                                                                    }
                                                                }
                                                                _ = v90
                                                            }
                                                            var v96 int = fi68
                                                            _ = v96
                                                        }
                                                    case 32768:
                                                        switch v3 {
                                                        case "":
                                                            v4 = float32(3.14159)
                                                        }
                                                        if true {
                                                            fmt.Println(v5)
                                                            var v97 string = "a"
                                                            _ = v97
                                                        } else {
                                                            fmt.Println(v4, fi69, v3)
                                                            var v98 uint16 = uint16(v1)
                                                            fmt.Println(fi68)
                                                            _ = v98
                                                        }
                                                    default:
                                                        fmt.Println(fi68, v3)
                                                    }
                                                    _ = v83
                                                }
                                                var v99 uint8 = (v5 ^ 219)
                                                _ = v99
                                            } else {
                                                var v100 uint64 = (v1 * v1)
                                                _ = v100
                                            }
                                        } else {
                                            var v101 uint32 = 65535
                                            var v102 int32 = 1
                                            v24 = (v101 & uint32(v2))
                                            _ = v101
                                            _ = v102
                                        }
                                        _ = v81
                                    }
                                default:
                                    switch fi68 {
                                    case 2147483647:
                                        var v103 uint = 255
                                        _ = v103
                                    default:
                                        v104 := []int{}
                                        _ = v104
                                    }
                                }
                            }
                        }
                    case 4294967295:
                        if true {
                            if ((^(v24)) < v24) {
                                var v105 int16 = -709
                                var v106 int16 = v105
                                v1 += v1
                                _ = v105
                                _ = v106
                            }
                        } else {
                            var v107 float64 = 1e10
                            var v108 float32 = v4
                            var v109 float64 = (v107 / (v107 - (float64(v108) - v107)))
                            _ = v107
                            _ = v108
                            _ = v109
                        }
                    default:
                        var v110 uint32 = 4294967295
                        _ = v110
                    }
                    if (false || false) {
                        if ((v3 + v3) == (v3 + v3)) {
                            var v111 string = "a"
                            var v112 string = (v111 + v111)
                            v112 = (v112 + v112)
                            _ = v111
                            _ = v112
                        }
                        v5 += (v5 | (v5 + 0))
                        var v113 uint16 = 1
                        _ = v113
                    }
                }
                switch (fi69 ^ fi68) {
                case 128:
                    if (-1 == 1) {
                        v114 := GenS4{F0: (v5 + 64)}
                        _ = v114
                    } else {
                        if ((v3 + v3) != v3) {
                            var v115 uint16 = v2
                            _ = v115
                        }
                        fmt.Println(v24, v24, v1)
                        var v116 int64 = 42
                        _ = v116
                    }
                }
                _ = fi69
                _ = v70
            }
        case "foo bar":
            v5 -= v5
        }
        v3 = v3
        _ = fi68
    }
    v117 := []int{}
    v1 = v1
    var v118 string = (v3 + "a")
    fmt.Println(v2, v4, v3)
    v1 += uint64(v2)
    var v119 string = v118
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v117:", len(v117), v117)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v117
    _ = v118
    _ = v119
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
