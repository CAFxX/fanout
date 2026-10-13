package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int16; F1 uint32; F2 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint; F1 uint; F2 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint8; F1 uint }
type GenS10 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int32, p1 uint, p2 int8) string {
    p0 = -1
    v6 := GenS1{F0: 32767, F1: 128, F2: 180}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ("foo bar" + "")
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 547
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    v5 := make(map[int]int)
    if ((true && ((1 != 255) || (317 == 330))) && true) {
        v7 := v1[88]
        v8 := GenS1{F0: 32767, F1: 0, F2: 150}
        for fi9 := 0; fi9 < 17; fi9++ {
            if ((!(-1 == 42)) && false) {
                if (true || false) {
                    if ((^(v2)) == 0) {
                        var v10 uint = 453
                        switch "a" {
                        case "":
                            var v11 bool = false
                            _ = v11
                        default:
                            v2 = uint64(v7)
                        }
                        delete(v4, v7)
                        _ = v10
                    }
                    v1[fi9] = v7
                }
                switch ("" + ("a" + ("hello" + "foo bar"))) {
                case "a":
                    delete(v1, fi9)
                case "":
                    v12 := GenS2{F0: false, F1: 1}
                    _ = v12
                case "hello":
                    v13 := []int{fi9, (-(fi9)), fi9, 32768}
                    _ = v13
                }
                var v14 float32 = 0.0
                _ = v14
            }
            _ = fi9
        }
        _ = v7
        _ = v8
    } else {
        v15 := GenS2{F0: false, F1: 2}
        if (1 <= -797) {
            var v16 int64 = -1
            v2 = (v2 | (v2 >> ((26) & 63)))
            var v17 int8 = 42
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            v2 = v2
            _ = fi18
        }
        _ = v15
    }
    fmt.Println(v2)
    if false {
        v19 := GenS2{F0: (1.5 != -2.25), F1: 4294967295}
        delete(v5, -1)
        _ = v19
    }
    if (42 > 0) {
        if true {
            for fi20 := 0; fi20 < 17; fi20++ {
                v2 += v2
                if true {
                    switch fi20 {
                    case -893:
                        v21 := GenS2{F0: (681 <= -1), F1: 161}
                        _ = v21
                    default:
                        fmt.Println(fi20, fi20)
                    }
                    v2 = 284
                }
                _ = fi20
            }
            v22 := GenS1{F0: 107, F1: 114, F2: -650}
            var v23 float32 = 0.0
            _ = v22
            _ = v23
        }
    } else {
        var v24 int32 = 404
        v25 := GenS2{F0: false, F1: 2}
        v2 = 4294967295
        _ = v24
        _ = v25
    }
    if (true || true) {
        v26 := GenS3{F0: 0, F1: 1, F2: true}
        switch -162 {
        case 621:
            v27 := []int{-718}
            _ = v27
        default:
            v28 := GenS3{F0: 32767, F1: 0, F2: (3.14159 < (1.5 * 1.5))}
            _ = v28
        }
        _ = v26
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        for fi30 := 0; fi30 < 14; fi30++ {
            if false {
                v31 := GenS1{F0: 0, F1: 2, F2: 127}
                v32 := GenS3{F0: 42, F1: 0, F2: false}
                v33 := []int{fi29}
                if true {
                    if false {
                        v34 := GenS1{F0: 42, F1: 128, F2: 9223372036854775807}
                        v35 := v1[fi30]
                        var v36 int = 4294967295
                        _ = v34
                        _ = v35
                        _ = v36
                    } else {
                        var v37 string = "foo bar"
                        v2 += 4294967295
                        _ = v37
                    }
                    fmt.Println(fi30, fi30)
                    fmt.Println(fi30, fi29, fi29)
                    v2 = v2
                } else {
                    v38 := v1[(fi30 % (((fi29 * fi30)) | 1))]
                    v2 = 4294967295
                    v2 -= 4294967295
                    _ = v38
                }
                _ = v31
                _ = v32
                _ = v33
            }
            switch "hello" {
            case "":
                if ((!(4294967295 > 2147483647)) && (!(0 == 42))) {
                    var v39 uint16 = 1
                    _ = v39
                } else {
                    if (!(!false)) {
                        switch 57 {
                        case 42:
                            switch 1 {
                            case 1:
                                v40 := GenS2{F0: (false || ((((true && (("" + "hello") >= "foo bar")) || true) && false) || true)), F1: 705}
                                _ = v40
                            default:
                                v41 := []int{}
                                _ = v41
                            }
                        case 2147483647:
                            var v42 bool = ((fi29 >> ((541) & 63)) == (fi30 ^ (^(fi30))))
                            _ = v42
                        case -556:
                            v2 = v2
                        }
                    } else {
                        v43 := GenS4{F0: 32767}
                        _ = v43
                    }
                    delete(v4, 358)
                }
            case "foo bar":
                var v44 uint8 = 30
                _ = v44
            case "\n\t":
                var v45 float64 = (0.1 / (1e10 - -2.25))
                _ = v45
            default:
                var v46 int32 = 42
                _ = v46
            }
            v47 := GenS3{F0: -409, F1: 2147483648, F2: (!(2 >= 188))}
            _ = fi30
            _ = v47
        }
        _ = fi29
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        v2 = (v2 / ((v2) | 1))
        var v49 int8 = 127
        _ = fi48
        _ = v49
    }
    var v50 int64 = 9223372036854775807
    v51 := GenS1{F0: 32767, F1: 127, F2: 1}
    for fi52 := 0; fi52 < 16; fi52++ {
        v50 = v50
        v2 -= 427
        _ = fi52
    }
    v2 += v2
    for fi53 := 0; fi53 < 6; fi53++ {
        for fi54 := 0; fi54 < 7; fi54++ {
            v2 = 18446744073709551615
            if true {
                switch "hello" {
                case "a":
                    if (113 < 65535) {
                        switch fi53 {
                        case 4294967295:
                            switch fi53 {
                            case -936:
                                if (!(v50 > (v50 + v50))) {
                                    var v55 uint64 = v2
                                    if (-2147483648 == -2147483648) {
                                        v56 := []int{}
                                        _ = len(v56)
                                        var v57 uint64 = (v55 - (v55 & 1))
                                        _ = v56
                                        _ = v57
                                    }
                                    v58 := GenS2{F0: ((v50 >> ((0) & 63)) <= (v50 % (((v50 + (v50 & 42))) | 1))), F1: 2}
                                    if false {
                                        v2 += v55
                                        var v59 int64 = (v50 | (^(v50)))
                                        if (float32(0.1) > (-2.25 + 0.0)) {
                                            fmt.Println(v59, fi53, fi54)
                                        }
                                        _ = v59
                                    }
                                    var v60 float32 = 0.1
                                    v55 = (v2 * v2)
                                    _ = v55
                                    _ = v58
                                    _ = v60
                                } else {
                                    v2 -= 264
                                    var v61 uint32 = 0
                                    _ = v61
                                }
                            case 980:
                                var v62 uint64 = v2
                                _ = v62
                            default:
                                var v63 uint32 = 2147483647
                                _ = v63
                            }
                        default:
                            v50 = (^(v50))
                        }
                        var v64 int8 = 57
                        var v65 float32 = (3.14159 * (1e10 - (0.1 * 1e10)))
                        if true {
                            switch 906 {
                            case 823:
                                var v66 int64 = (v50 << ((4294967295) & 63))
                                _ = v66
                            case -1:
                                v67 := GenS5{F0: 255, F1: v50, F2: 88}
                                _ = v67
                            }
                            v2 = (^(v2))
                            v68 := GenS3{F0: 0, F1: 2, F2: true}
                            v69 := v1[0]
                            var v70 string = "\n\t"
                            _ = v68
                            _ = v69
                            _ = v70
                        } else {
                            v3[(fi53 << ((255) & 63))] = 249
                            var v71 int8 = (v64 ^ v64)
                            _ = v71
                        }
                        _ = v64
                        _ = v65
                    } else {
                        var v72 float32 = 3.14159
                        v73 := []int{fi54, 32767}
                        var v74 float64 = float64(0.1)
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                default:
                    v4[42] = -676
                }
                if false {
                    var v75 int = (fi53 << ((uint(v2)) & 63))
                    switch "hello" {
                    case "":
                        var v76 uint32 = 1
                        _ = v76
                    case "a":
                        if (2 > 255) {
                            fmt.Println(v2)
                            if true {
                                switch (fi53 << ((42) & 63)) {
                                case -262:
                                    v50 = v50
                                case 127:
                                    var v77 int8 = 31
                                    _ = v77
                                case -9223372036854775808:
                                    v5[fi53] = (fi53 - (fi53 << ((857) & 63)))
                                default:
                                    delete(v4, 2147483647)
                                }
                            } else {
                                delete(v1, (^(fi54)))
                                v78 := GenS6{F0: 65535}
                                v75 = fi53
                                _ = v78
                            }
                        } else {
                            v50 = v50
                            v79 := []int{(^(v75)), fi53, 32767}
                            v80 := GenS2{F0: false, F1: uint(v2)}
                            _ = v79
                            _ = v80
                        }
                    case "hello":
                        fmt.Println(v75)
                    }
                    var v81 string = (("foo bar" + ("foo bar" + "hello")) + "\n\t")
                    fmt.Println(v2)
                    _ = v75
                    _ = v81
                } else {
                    v1[fi53] = 114
                    v82 := []int{(fi53 * fi53), fi54, fi54}
                    _ = v82
                }
                v83 := GenS7{F0: 4294967295, F1: 1, F2: fi53}
                switch (fi53 % ((-487) | 1)) {
                case 0:
                    v84 := []int{-1}
                    _ = v84
                case 1:
                    if (-32768 <= 32767) {
                        v2 -= v2
                        var v85 float32 = 0.0
                        _ = v85
                    }
                }
                _ = v83
            } else {
                if false {
                    v86 := GenS8{F0: (2 <= 1)}
                    _ = v86
                } else {
                    v2 += (v2 ^ (v2 * 414))
                    switch fi53 {
                    case -1:
                        if ((!true) && true) {
                            v50 = (v50 * 200)
                        } else {
                            fmt.Println(fi53, v50, v50)
                            if (71 == 78) {
                                delete(v4, 2)
                                v50 += (v50 ^ (v50 >> ((4294967295) & 63)))
                                v87 := GenS7{F0: 255, F1: 65535, F2: (fi54 | (fi53 ^ fi53))}
                                _ = v87
                            }
                        }
                    }
                    if true {
                        if true {
                            v88 := GenS8{F0: (!((v2 * v2) < 0))}
                            _ = v88
                        }
                        v50 = (v50 & (v50 | (v50 * v50)))
                    } else {
                        v89 := GenS6{F0: 0}
                        if false {
                            switch fi53 {
                            case 128:
                                fmt.Println(v2, v2)
                                fmt.Println(v50, fi53)
                                v50 = -1
                                var v90 uint = 65535
                                _ = v90
                            case 9223372036854775807:
                                v2 = (v2 ^ (v2 ^ (v2 | 0)))
                            case -2147483648:
                                if (-1 < -840) {
                                    var v91 float64 = (1e10 / ((1e10 + 0.1) / (1.5 * ((-2.25 - 0.0) * (0.1 + ((float64(float32(0.0)) * 1e10) + -2.25))))))
                                    fmt.Println(fi54, fi53, v50)
                                    _ = v91
                                }
                            }
                        } else {
                            v92 := []int{(fi54 + fi54)}
                            _ = v92
                        }
                        switch int(v50) {
                        case 9223372036854775807:
                            var v93 uint8 = 127
                            _ = v93
                        case -2147483648:
                            fmt.Println(fi53, fi54, fi54)
                        case 259:
                            var v94 float64 = 0.1
                            _ = v94
                        default:
                            if (v2 != v2) {
                                if false {
                                    v95 := GenS9{F0: 172, F1: 711}
                                    _ = v95
                                }
                            }
                        }
                        fmt.Println(v50, fi54, fi54)
                        _ = v89
                    }
                }
                fmt.Println(fi53, fi53, fi54)
                fmt.Println(fi53, v2, fi53)
            }
            _ = fi54
        }
        var v96 int16 = -712
        _ = fi53
        _ = v96
    }
    if (-32768 >= 1) {
        v97 := GenS10{F0: 128, F1: 0}
        v50 -= v50
        _ = v97
    }
    if false {
        v98 := GenS9{F0: 71, F1: uint(v2)}
        _ = v98
    } else {
        switch "" {
        case "hello":
            for fi99 := 0; fi99 < 15; fi99++ {
                v2 += v2
                _ = fi99
            }
        case "foo bar":
            var v100 int = 9223372036854775807
            _ = v100
        case "":
            v50 -= 1
        }
        var v101 int32 = 839
        var v102 uint64 = (v2 - (v2 ^ (v2 ^ (v2 / ((0) | 1)))))
        _ = v101
        _ = v102
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
}
