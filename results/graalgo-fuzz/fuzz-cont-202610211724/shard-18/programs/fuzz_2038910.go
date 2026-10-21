package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 string }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint; F2 int32 }
type GenS5 struct { F0 bool; F1 uint8; F2 uint }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 uint64 }
type GenS9 struct { F0 int8; F1 uint32 }
type GenS10 struct { F0 int; F1 uint }
func genFn0() int {
    v5 := []int{}
    var v6 int64 = 83
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v7 := GenS1{F0: 0, F1: -711}
    switch "" {
    case "a":
        if true {
            v8 := GenS1{F0: -1, F1: 1}
            var v9 uint8 = 1
            _ = v8
            _ = v9
        }
    case "":
        var v10 uint16 = 32767
        _ = v10
    case "foo bar":
        switch 32767 {
        case -9223372036854775808:
            v6 = v6
        case 32768:
            var v11 float64 = float64(((1e10 + -2.25) + (0.1 / ((1e10 * 1.5) / 3.14159))))
            _ = v11
        case 4294967295:
            switch ("hello" + "") {
            case "foo bar":
                if (true || ("" == "hello")) {
                    var v12 float32 = (-2.25 - (0.0 + 0.0))
                    var v13 int = 388
                    _ = v12
                    _ = v13
                } else {
                    var v14 uint64 = 18446744073709551615
                    for fi15 := 0; fi15 < 20; fi15++ {
                        if len(v5) > 0 {
                            v5[0] = (fi15 & (fi15 % ((fi15) | 1)))
                        }
                        if (!true) {
                            fmt.Println(fi15, v14)
                        }
                        if false {
                            for fi16 := 0; fi16 < 18; fi16++ {
                                var v17 string = "a"
                                _ = fi16
                                _ = v17
                            }
                        }
                        _ = fi15
                    }
                    _ = v14
                }
            case "":
                for fi18 := 0; fi18 < 7; fi18++ {
                    if true {
                        v19 := GenS1{F0: -57, F1: 1}
                        if true {
                            if ((4294967295 >= 0) || true) {
                                v5 = append(v5, fi18)
                            }
                            v6 = (v6 * v6)
                        } else {
                            var v20 float64 = 3.14159
                            for fi21 := 0; fi21 < 12; fi21++ {
                                v6 -= (v6 | (v6 >> ((2) & 63)))
                                var v22 bool = false
                                v20 = 1.5
                                _ = fi21
                                _ = v22
                            }
                            if false {
                                var v23 float64 = (0.1 - float64(float32(float64((0.0 / -2.25)))))
                                var v24 int8 = -33
                                _ = v23
                                _ = v24
                            }
                            _ = v20
                        }
                        fmt.Println(fi18)
                        _ = v19
                    } else {
                        v6 = (v6 + (v6 / (((v6 ^ -9223372036854775808)) | 1)))
                        v6 = v6
                    }
                    _ = fi18
                }
            default:
                v25 := GenS2{F0: (-(v6)), F1: 50, F2: "\n\t"}
                _ = v25
            }
        default:
            switch 30 {
            case 2147483647:
                _ = len(v5)
            case -213:
                for fi26 := 0; fi26 < 8; fi26++ {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        if ((-128 < -128) && true) {
                            if (128 >= 128) {
                                fmt.Println(v6)
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = fi26
                                }
                                var v28 uint = 2
                                v6 = (v6 ^ (v6 + 42))
                                _ = v28
                            }
                            _ = len(v5)
                        } else {
                            var v29 int64 = (v6 % ((v6) | 1))
                            v30 := GenS1{F0: 1, F1: 42}
                            var v31 int64 = v29
                            if true {
                                var v32 float64 = (float64(1.5) * -2.25)
                                if len(v5) > 0 {
                                    v5[0] = (fi27 << ((1) & 63))
                                }
                                var v33 float64 = v32
                                _ = v32
                                _ = v33
                            } else {
                                if false {
                                    v34 := GenS1{F0: -932, F1: 0}
                                    v6 = 0
                                    if ((!((-252 != -1) || (702 > 4294967295))) && false) {
                                        v5 = append(v5, (^(fi27)))
                                        var v35 uint64 = 4294967295
                                        v5 = append(v5, (fi27 & fi26))
                                        v36 := GenS3{F0: 1, F1: 2}
                                        _ = v35
                                        _ = v36
                                    } else {
                                        var v37 bool = ((fi27 / ((fi26) | 1)) >= (fi27 ^ (fi26 * (fi26 + fi26))))
                                        _ = v37
                                    }
                                    _ = v34
                                } else {
                                    var v38 bool = false
                                    var v39 uint32 = 67
                                    switch fi26 {
                                    case -1:
                                        v6 -= (v31 ^ -9223372036854775808)
                                    case 4294967295:
                                        if false {
                                            var v40 uint = 18446744073709551615
                                            _ = v40
                                        }
                                    case 32767:
                                        v5 = append(v5, -1)
                                    }
                                    _ = v38
                                    _ = v39
                                }
                                var v41 int = (fi26 * fi26)
                                var v42 uint16 = 1
                                _ = v41
                                _ = v42
                            }
                            if ((1 < 1) && true) {
                                var v43 uint16 = 2
                                _ = v43
                            }
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = fi27
                    }
                    var v44 string = "\n\t"
                    _ = fi26
                    _ = v44
                }
            case 931:
                _ = len(v5)
            }
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 string = (("foo bar" + ("\n\t" + "a")) + "foo bar")
    var v2 uint64 = 18446744073709551615
    var v3 int64 = 1
    var v4 int8 = 1
    v4 += v4
    v45 := GenS3{F0: 0, F1: 65535}
    switch v1 {
    case "a":
        if (!(v3 > -1)) {
            var v46 float64 = 1e10
            v47 := []int{}
            _ = v46
            _ = v47
        } else {
            for fi48 := 0; fi48 < 7; fi48++ {
                switch v1 {
                case "a":
                    switch fi48 {
                    case -2147483648:
                        var v49 int = genFn0()
                        _ = v49
                    case 603:
                        if false {
                            if false {
                                var v50 int8 = v4
                                _ = v50
                            } else {
                                var v51 uint32 = 4294967295
                                _ = v51
                            }
                        } else {
                            if (94 == 65535) {
                                v3 = (v3 | (v3 << ((255) & 63)))
                            } else {
                                fmt.Println(fi48, v2, v2)
                            }
                            if false {
                                v52 := []int{fi48, fi48}
                                v4 -= 127
                                _ = v52
                            } else {
                                switch fi48 {
                                case 1:
                                    for fi53 := 0; fi53 < 7; fi53++ {
                                        v54 := GenS2{F0: v3, F1: 780, F2: "hello"}
                                        _ = fi53
                                        _ = v54
                                    }
                                case 89:
                                    var v55 float32 = float32(((0.0 / (0.0 + (1e10 + 1.5))) / 1.5))
                                    _ = v55
                                }
                                var v56 uint16 = 323
                                switch v1 {
                                case "foo bar":
                                    if (((v4 * (^(v4))) != 1) || (v4 < v4)) {
                                        if (0 == -1) {
                                            v4 = (v4 ^ v4)
                                            var v57 uint8 = 2
                                            switch fi48 {
                                            case -9223372036854775808:
                                                var v58 uint16 = v56
                                                _ = v58
                                            case 4294967295:
                                                switch v1 {
                                                case "hello":
                                                    var v59 int8 = v4
                                                    _ = v59
                                                default:
                                                    var v60 float32 = 1e10
                                                    _ = v60
                                                }
                                            default:
                                                switch (fi48 - fi48) {
                                                case -2147483648:
                                                    if true {
                                                        if (!false) {
                                                            v61 := GenS2{F0: 144, F1: uint32(v56), F2: (v1 + v1)}
                                                            _ = v61
                                                        } else {
                                                            switch (fi48 + (fi48 % ((genFn0()) | 1))) {
                                                            case -493:
                                                                var v62 float32 = ((3.14159 * 1e10) + 1e10)
                                                                _ = v62
                                                            }
                                                            switch genFn0() {
                                                            case 792:
                                                                if ((false || false) && false) {
                                                                    var v63 int8 = v4
                                                                    for fi64 := 0; fi64 < 18; fi64++ {
                                                                        var v65 int32 = 1
                                                                        v2 = 42
                                                                        v1 = "a"
                                                                        if false {
                                                                            v56 = uint16(v2)
                                                                            v66 := []int{fi48}
                                                                            _ = v66
                                                                        }
                                                                        if false {
                                                                            v56 = (v56 + 42)
                                                                            fmt.Println(v63, fi64)
                                                                        } else {
                                                                            v67 := []int{(fi48 / (((fi64 << ((4294967295) & 63))) | 1))}
                                                                            var v68 int64 = 0
                                                                            v67 = append(v67, 128)
                                                                            _ = v67
                                                                            _ = v68
                                                                        }
                                                                        _ = fi64
                                                                        _ = v65
                                                                    }
                                                                    _ = v63
                                                                } else {
                                                                    for fi69 := 0; fi69 < 8; fi69++ {
                                                                        v4 += v4
                                                                        _ = fi69
                                                                    }
                                                                    v70 := []int{(fi48 << ((281) & 63))}
                                                                    _ = v70
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        for fi71 := 0; fi71 < 7; fi71++ {
                                                            switch v1 {
                                                            case "a":
                                                                v1 = "\n\t"
                                                            case "foo bar":
                                                                fmt.Println(v56, fi71)
                                                            default:
                                                                v56 -= (v56 | 2)
                                                            }
                                                            _ = fi71
                                                        }
                                                        v72 := GenS2{F0: v3, F1: 65535, F2: ""}
                                                        for fi73 := 0; fi73 < 10; fi73++ {
                                                            v74 := GenS4{F0: 2147483647, F1: 2, F2: int32(v3)}
                                                            var v75 int64 = (v3 >> ((42) & 63))
                                                            var v76 uint = 42
                                                            fmt.Println(v57)
                                                            _ = fi73
                                                            _ = v74
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        _ = v72
                                                    }
                                                case 1:
                                                    var v77 float64 = 1e10
                                                    _ = v77
                                                case 0:
                                                    fmt.Println(v57, v3)
                                                default:
                                                    var v78 int8 = -60
                                                    _ = v78
                                                }
                                            }
                                            _ = v57
                                        } else {
                                            fmt.Println(fi48, v4, v4)
                                            switch ("hello" + "foo bar") {
                                            case "foo bar":
                                                v79 := GenS5{F0: (false && (4294967295 >= 42)), F1: 245, F2: 0}
                                                _ = v79
                                            case "":
                                                if (!((^(v2)) <= (v2 ^ (v2 << ((255) & 63))))) {
                                                    var v80 float64 = 0.1
                                                    _ = v80
                                                } else {
                                                    if false {
                                                        if true {
                                                            v81 := []int{-779, fi48, 2147483647}
                                                            _ = v81
                                                        } else {
                                                            if false {
                                                                if false {
                                                                    v4 = (v4 / (((v4 * v4)) | 1))
                                                                    v82 := GenS6{F0: 1, F1: -589}
                                                                    _ = v82
                                                                }
                                                            }
                                                            for fi83 := 0; fi83 < 9; fi83++ {
                                                                if (-122 >= (v4 - 42)) {
                                                                    if false {
                                                                        fmt.Println(fi48)
                                                                        var v84 float64 = 0.1
                                                                        _ = v84
                                                                    } else {
                                                                        switch v1 {
                                                                        case "\n\t":
                                                                            v85 := GenS7{F0: (v4 - (v4 % (((v4 - v4)) | 1)))}
                                                                            _ = v85
                                                                        case "hello":
                                                                            fmt.Println(v1, v3, fi83)
                                                                        }
                                                                    }
                                                                    v86 := GenS3{F0: uint(v56), F1: 637}
                                                                    _ = v86
                                                                }
                                                                var v87 int8 = v4
                                                                _ = fi83
                                                                _ = v87
                                                            }
                                                            if false {
                                                                fmt.Println(v4)
                                                                var v88 uint64 = (v2 + 0)
                                                                _ = v88
                                                            } else {
                                                                v89 := GenS3{F0: 65535, F1: 0}
                                                                var v90 uint = 140
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                        }
                                                        if (65535 != 4294967295) {
                                                            v91 := []int{-1, 4294967295, fi48, genFn0()}
                                                            if len(v91) >= 2 {
                                                                v91 = v91[1:len(v91)]
                                                            }
                                                            _ = v91
                                                        } else {
                                                            v1 = (v1 + v1)
                                                            v92 := GenS3{F0: 42, F1: 2147483648}
                                                            _ = v92
                                                        }
                                                        var v93 int16 = 184
                                                        _ = v93
                                                    }
                                                    v4 = -128
                                                }
                                            case "hello":
                                                if (v3 >= 1) {
                                                    v94 := []int{int(v2), -459, (fi48 & (-(fi48)))}
                                                    v2 -= v2
                                                    v95 := GenS8{F0: 0, F1: (v2 / (((v2 * 2)) | 1))}
                                                    _ = v94
                                                    _ = v95
                                                }
                                            default:
                                                v96 := []int{(-(fi48)), (fi48 >> ((18446744073709551615) & 63))}
                                                _ = v96
                                            }
                                            for fi97 := 0; fi97 < 16; fi97++ {
                                                v4 = v4
                                                var v98 uint8 = 19
                                                v99 := GenS5{F0: false, F1: v98, F2: 255}
                                                var v100 int = genFn0()
                                                _ = fi97
                                                _ = v100
                                                _ = v98
                                                _ = v99
                                            }
                                        }
                                        var v101 bool = false
                                        _ = v101
                                    }
                                case "":
                                    for fi102 := 0; fi102 < 14; fi102++ {
                                        var v103 float32 = -2.25
                                        v104 := []int{32768, fi102}
                                        _ = fi102
                                        _ = v103
                                        _ = v104
                                    }
                                default:
                                    v105 := []int{}
                                    _ = v105
                                }
                                _ = v56
                            }
                        }
                    }
                }
                v1 = v1
                fmt.Println(v3)
                _ = fi48
            }
        }
    case "":
        for fi106 := 0; fi106 < 16; fi106++ {
            var v107 int = -134
            v107 -= v107
            _ = fi106
            _ = v107
        }
    default:
        for fi108 := 0; fi108 < 2; fi108++ {
            v109 := GenS3{F0: 4294967295, F1: 1}
            if (v3 > (v3 ^ v3)) {
                v3 -= -1
                v110 := GenS9{F0: int8(v2), F1: 36}
                v111 := []int{(fi108 / ((4294967295) | 1)), (fi108 ^ (fi108 - genFn0())), fi108, genFn0()}
                _ = v110
                _ = v111
            }
            _ = fi108
            _ = v109
        }
    }
    v112 := GenS2{F0: 1, F1: 29, F2: "\n\t"}
    v3 += -1
    switch 127 {
    case 604:
        var v113 int16 = 0
        _ = v113
    default:
        v114 := GenS10{F0: 924, F1: uint(v2)}
        _ = v114
    }
    switch 32767 {
    case 1:
        v115 := []int{42, 2147483647, 32767}
        _ = v115
    case 127:
        fmt.Println(v1)
    case -542:
        fmt.Println(v3)
    default:
        v3 = (v3 % ((v3) | 1))
    }
    var v116 uint32 = 1
    v4 -= v4
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v112
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v45
}
