package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint; F1 int32; F2 int64 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int8; F1 uint32 }
type GenS12 struct { F0 string }
func genFn0(p0 uint, p1 uint, p2 uint) uint32 {
    var v5 int32 = 42
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{fi7}
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi7
            _ = v8
        }
        fmt.Println(p2)
        var v9 int8 = -73
        _ = fi6
        _ = v9
    }
    p1 = (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    var v1 uint32 = 91
    var v2 uint64 = 718
    var v3 uint8 = 62
    var v4 bool = false
    fmt.Println(v2, v1)
    v3 = (^(v3))
    for fi10 := 0; fi10 < 6; fi10++ {
        if v4 {
            for fi11 := 0; fi11 < 5; fi11++ {
                switch "a" {
                case "hello":
                    switch fi10 {
                    case 128:
                        var v12 uint = 2
                        if true {
                            v12 = v12
                        } else {
                            v13 := []int{}
                            _ = v13
                        }
                        fmt.Println(v1)
                        _ = v12
                    }
                case "\n\t":
                    if true {
                        fmt.Println(fi11)
                        if v4 {
                            var v14 int16 = -724
                            v15 := GenS1{F0: 42, F1: v1, F2: int8(v3)}
                            _ = v14
                            _ = v15
                        }
                    } else {
                        if false {
                            var v16 uint16 = 262
                            _ = v16
                        } else {
                            var v17 uint32 = 2
                            switch fi11 {
                            case 1:
                                if (v1 < (^(v1))) {
                                    if v4 {
                                        fmt.Println(fi10)
                                        var v18 int16 = 42
                                        var v19 int8 = 100
                                        v2 += v2
                                        _ = v18
                                        _ = v19
                                    } else {
                                        v20 := GenS1{F0: (v2 % ((v2) | 1)), F1: (^(v17)), F2: 0}
                                        _ = v20
                                    }
                                    var v21 float32 = 0.0
                                    v4 = false
                                    fmt.Println(fi10, v2)
                                    v22 := GenS1{F0: v2, F1: uint32(v2), F2: -128}
                                    _ = v21
                                    _ = v22
                                }
                            }
                            if (!(v3 < v3)) {
                                var v23 uint = uint(v3)
                                v23 -= (v23 << (((v23 + (v23 | (v23 ^ v23)))) & 63))
                                _ = v23
                            } else {
                                switch fi11 {
                                case -9223372036854775808:
                                    var v24 float32 = 0.1
                                    _ = v24
                                case 2147483647:
                                    v1 = v17
                                    var v25 int16 = 632
                                    switch fi11 {
                                    case -1:
                                        if (true && v4) {
                                            var v26 uint16 = 32767
                                            _ = v26
                                        }
                                    }
                                    v2 = (^(v2))
                                    _ = v25
                                default:
                                    v4 = (1 < 42)
                                }
                            }
                            _ = v17
                        }
                        if v4 {
                            if (420 > 0) {
                                switch ("" + ("" + "a")) {
                                case "a":
                                    var v27 uint64 = 42
                                    _ = v27
                                case "\n\t":
                                    v28 := []int{fi11, fi10, 354, fi10}
                                    _ = v28
                                }
                            }
                            v29 := []int{fi11}
                            v1 = (v1 & (v1 ^ 108))
                            if (v4 && v4) {
                                var v30 uint = 2
                                fmt.Println(v30, fi10)
                                v31 := GenS2{F0: -32768, F1: (fi11 / ((fi10) | 1)), F2: -1}
                                _ = v30
                                _ = v31
                            }
                            fmt.Println(v2, v2, v1)
                            _ = v29
                        }
                        var v32 uint32 = v1
                        _ = v32
                    }
                case "foo bar":
                    if ((1e10 >= ((float64(0.1) - (3.14159 - 3.14159)) / -2.25)) || (v4 && false)) {
                        v33 := GenS1{F0: (v2 >> ((18446744073709551615) & 63)), F1: (v1 | (v1 | 2147483648)), F2: 82}
                        v3 -= 2
                        if (((3.14159 + float32(1e10)) * -2.25) > 1e10) {
                            fmt.Println(v2, v1)
                            v34 := GenS3{F0: 1}
                            v35 := GenS1{F0: (v2 % ((v2) | 1)), F1: (v1 - 2), F2: -128}
                            _ = v34
                            _ = v35
                        } else {
                            v36 := GenS2{F0: 1, F1: fi11, F2: -1}
                            if v4 {
                                v2 = v2
                                if v4 {
                                    v37 := GenS4{F0: 599, F1: 823}
                                    v2 -= 2
                                    if ((v2 + (v2 ^ v2)) < 0) {
                                        if (0 >= 42) {
                                            v38 := GenS5{F0: 83, F1: -265, F2: -1}
                                            fmt.Println(v2, fi10)
                                            v39 := GenS6{F0: v2, F1: (39 > 2)}
                                            v40 := GenS7{F0: int64(v3)}
                                            v3 = 0
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        } else {
                                            v1 -= genFn0(18446744073709551615, 2, 65535)
                                            var v41 uint = 1
                                            _ = v41
                                        }
                                    }
                                    _ = v37
                                }
                            } else {
                                var v42 float64 = (1.5 - 3.14159)
                                _ = v42
                            }
                            _ = v36
                        }
                        v43 := GenS3{F0: 65535}
                        if v4 {
                            switch fi10 {
                            case 93:
                                fmt.Println(v2, v2)
                            }
                            if v4 {
                                if (-698 > 292) {
                                    v4 = true
                                    var v44 int64 = 0
                                    v44 -= 1
                                    _ = v44
                                }
                                v45 := []int{(fi10 >> ((2) & 63)), fi10}
                                _ = v45
                            } else {
                                v3 -= 20
                                fmt.Println(fi11, fi11)
                            }
                        } else {
                            fmt.Println(v4)
                            v3 -= (v3 % (((v3 - (v3 + 69))) | 1))
                        }
                        _ = v33
                        _ = v43
                    } else {
                        switch 42 {
                        case 927:
                            v46 := GenS6{F0: (v2 & (^(v2))), F1: (32767 >= 1)}
                            _ = v46
                        case -907:
                            v4 = v4
                        case -437:
                            v4 = ((!(-896 == -16)) || (42 <= 227))
                        default:
                            if ("\n\t" < "foo bar") {
                                if (-2147483648 > 2147483647) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v47 bool = ((v3 >> ((221) & 63)) <= 196)
                                        v48 := []int{fi10, (fi11 | fi11), fi10}
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                v49 := GenS8{F0: 42, F1: 1}
                                switch ((("a" + ((("\n\t" + "") + "foo bar") + "foo bar")) + "") + (("foo bar" + "") + "a")) {
                                case "\n\t":
                                    v50 := []int{32767, fi11}
                                    _ = v50
                                case "hello":
                                    switch "\n\t" {
                                    case "\n\t":
                                        if true {
                                            v51 := []int{fi10, fi11, fi10}
                                            v4 = v4
                                            v52 := GenS9{F0: 331, F1: (v1 - 519)}
                                            _ = v51
                                            _ = v52
                                        } else {
                                            v53 := []int{fi10, (fi10 * 128), fi11, (fi11 * fi10)}
                                            _ = v53
                                        }
                                    default:
                                        switch ("foo bar" + ("" + "\n\t")) {
                                        case "\n\t":
                                            if false {
                                                v2 = 18446744073709551615
                                            }
                                        case "":
                                            var v54 uint = uint(v3)
                                            if (-1 >= 0) {
                                                v55 := GenS2{F0: 42, F1: (fi11 / (((fi10 + (fi11 % (((fi11 + fi11)) | 1)))) | 1)), F2: 0}
                                                v56 := GenS10{F0: (v3 << ((v54) & 63))}
                                                switch fi11 {
                                                case 32768:
                                                    v3 -= v3
                                                }
                                                _ = v55
                                                _ = v56
                                            }
                                            _ = v54
                                        case "a":
                                            var v57 int8 = 95
                                            _ = v57
                                        }
                                    }
                                }
                                fmt.Println(fi10, v4, fi11)
                                _ = v49
                            }
                        }
                        v1 = 0
                    }
                default:
                    if v4 {
                        v3 -= v3
                    }
                }
                _ = fi11
            }
            if v4 {
                var v58 int64 = 411
                switch "" {
                case "\n\t":
                    v59 := []int{fi10, (fi10 | (fi10 >> ((639) & 63))), -847, (fi10 & fi10)}
                    _ = v59
                case "":
                    var v60 bool = v4
                    _ = v60
                case "hello":
                    for fi61 := 0; fi61 < 10; fi61++ {
                        if ((v3 | (v3 >> ((42) & 63))) == 222) {
                            fmt.Println(v3)
                            var v62 uint16 = 0
                            var v63 int32 = 42
                            fmt.Println(v3)
                            _ = v62
                            _ = v63
                        } else {
                            v2 = 2
                            v64 := []int{(fi61 >> ((168) & 63)), (fi10 << ((409) & 63)), int(v2), (-(fi10))}
                            _ = v64
                        }
                        _ = fi61
                    }
                }
                v65 := []int{}
                _ = v58
                _ = v65
            }
            v66 := []int{}
            _ = v66
        }
        if (v4 && true) {
            v2 += v2
        }
        var v67 uint64 = v2
        _ = fi10
        _ = v67
    }
    v68 := []int{}
    var v69 uint32 = genFn0(65535, 130, 1)
    if v4 {
        if len(v68) > 0 {
            v68[0] = -9223372036854775808
        }
        var v70 int64 = 950
        _ = v70
    }
    if (361 >= 2) {
        switch "hello" {
        case "a":
            if len(v68) > 0 {
                v68[0] = int(v69)
            }
        default:
            for fi71 := 0; fi71 < 8; fi71++ {
                var v72 int8 = -128
                if true {
                    fmt.Println(v72, fi71, v72)
                    _ = len(v68)
                    if len(v68) >= 2 {
                        v68 = v68[1:len(v68)]
                    }
                }
                fmt.Println(v1, v3)
                _ = fi71
                _ = v72
            }
        }
    }
    v1 += (v1 + v1)
    for fi73 := 0; fi73 < 1; fi73++ {
        v74 := GenS11{F0: -128, F1: v69}
        _ = fi73
        _ = v74
    }
    v69 += v1
    for fi75 := 0; fi75 < 5; fi75++ {
        var v76 uint8 = v3
        v68 = append(v68, fi75)
        var v77 uint32 = (v1 & v69)
        _ = fi75
        _ = v76
        _ = v77
    }
    v78 := GenS8{F0: 520, F1: 127}
    var v79 int16 = -978
    if len(v68) >= 2 {
        v68 = v68[1:len(v68)]
    }
    if (0.0 != (float32(1.5) - -2.25)) {
        var v80 int32 = -1
        var v81 int8 = 127
        switch -657 {
        case -559:
            switch ("" + ("hello" + "foo bar")) {
            case "hello":
                fmt.Println(v1, v79, v81)
            case "":
                for fi82 := 0; fi82 < 3; fi82++ {
                    fmt.Println(v80)
                    var v83 int16 = v79
                    fmt.Println(v3)
                    _ = fi82
                    _ = v83
                }
            default:
                var v84 int32 = (v80 + v80)
                _ = v84
            }
        case 32768:
            var v85 bool = (128 != 128)
            _ = v85
        default:
            if true {
                var v86 int = -2147483648
                _ = v86
            }
        }
        _ = v80
        _ = v81
    } else {
        var v87 int32 = 1
        _ = v87
    }
    for fi88 := 0; fi88 < 2; fi88++ {
        var v89 uint64 = (v2 | (v2 % (((^(v2))) | 1)))
        v90 := GenS12{F0: ("" + ("\n\t" + "a"))}
        for fi91 := 0; fi91 < 20; fi91++ {
            v79 = v79
            _ = fi91
        }
        _ = fi88
        _ = v89
        _ = v90
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v68:", len(v68), v68)
    fmt.Println("v69:", v69)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v68
    _ = v69
    _ = v78
    _ = v79
}
