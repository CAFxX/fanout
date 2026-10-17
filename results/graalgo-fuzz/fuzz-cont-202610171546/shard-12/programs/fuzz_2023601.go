package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int64; F2 int8 }
type GenS5 struct { F0 int16; F1 bool; F2 bool }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 string; F1 uint64; F2 int32 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 int; F1 uint32 }
type GenS11 struct { F0 uint32; F1 uint64; F2 uint8 }
func genFn0(p0 uint, p1 int32, p2 uint16) uint64 {
    var v6 int64 = -9223372036854775808
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int64 = v6
        var v9 int32 = p1
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1() uint16 {
    v10 := GenS1{F0: 1}
    if (587 == 1) {
        var v11 float32 = 3.14159
        for fi12 := 0; fi12 < 15; fi12++ {
            v11 = -2.25
            v13 := GenS1{F0: 1}
            var v14 string = ""
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    }
    switch 1 {
    case -9223372036854775808:
        v15 := []int{2147483647, 128, -318}
        _ = v15
    case 9223372036854775807:
        for fi16 := 0; fi16 < 10; fi16++ {
            if (-114 < (fi16 | fi16)) {
                for fi17 := 0; fi17 < 15; fi17++ {
                    fmt.Println(fi16)
                    var v18 uint32 = 42
                    _ = fi17
                    _ = v18
                }
            } else {
                v19 := GenS2{F0: 2, F1: 4294967295}
                for fi20 := 0; fi20 < 15; fi20++ {
                    if (0 >= -2147483648) {
                        var v21 int8 = 127
                        v22 := []int{}
                        if false {
                            fmt.Println(fi16, fi16, v21)
                            v22 = append(v22, 1)
                        } else {
                            if (true && false) {
                                if (true || (false && true)) {
                                    var v23 int8 = (v21 % ((-55) | 1))
                                    var v24 int64 = 0
                                    var v25 uint64 = genFn0(18446744073709551615, 2147483647, 667)
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 string = "\n\t"
                                    if (-34 == (fi16 + fi16)) {
                                        v22 = append(v22, (^(fi20)))
                                    }
                                    var v27 int32 = 0
                                    _ = v26
                                    _ = v27
                                }
                            }
                            if true {
                                _ = len(v22)
                            }
                        }
                        var v28 int = (fi16 / ((fi20) | 1))
                        if true {
                            v28 = -1
                        } else {
                            var v29 int16 = 239
                            if true {
                                if true {
                                    v29 -= -605
                                }
                                switch ("foo bar" + "foo bar") {
                                case "foo bar":
                                    switch v28 {
                                    case 1:
                                        var v30 int16 = v29
                                        _ = v30
                                    case -72:
                                        v31 := GenS1{F0: -2147483648}
                                        _ = v31
                                    case 32768:
                                        var v32 int16 = (v29 + (v29 % ((42) | 1)))
                                        _ = v32
                                    }
                                default:
                                    var v33 float64 = 1.5
                                    _ = v33
                                }
                                var v34 uint16 = 1
                                _ = v34
                            } else {
                                _ = len(v22)
                                switch (v28 & fi16) {
                                case -1:
                                    switch 127 {
                                    case 42:
                                        if len(v22) > 0 {
                                            v22[0] = (fi16 / ((v28) | 1))
                                        }
                                    case 1:
                                        v35 := GenS2{F0: 1, F1: 4294967295}
                                        _ = v35
                                    case 692:
                                        if true {
                                            if true {
                                                v36 := GenS3{F0: v29}
                                                var v37 uint16 = 1
                                                if (fi20 <= fi16) {
                                                    if (v21 > (^(v21))) {
                                                        var v38 float32 = -2.25
                                                        _ = v38
                                                    } else {
                                                        v21 = v21
                                                        switch 832 {
                                                        case 269:
                                                            v37 -= v37
                                                        case -903:
                                                            v28 += fi16
                                                        case 127:
                                                            v28 = v28
                                                        default:
                                                            if true {
                                                                var v39 float32 = 1e10
                                                                v40 := GenS4{F0: -2147483648, F1: -1, F2: -43}
                                                                if true {
                                                                    v21 = (^(v21))
                                                                    if (v37 == (v37 / ((32767) | 1))) {
                                                                        if len(v22) >= 2 {
                                                                            v22 = v22[1:len(v22)]
                                                                        }
                                                                        if ((42 == 2) || true) {
                                                                            v22 = append(v22, (v28 % (((v28 | 640)) | 1)))
                                                                            if false {
                                                                                var v41 int16 = v29
                                                                                _ = v41
                                                                            }
                                                                        } else {
                                                                            var v42 float64 = (-2.25 / float64(v39))
                                                                            switch ("\n\t" + "hello") {
                                                                            case "a":
                                                                                v22 = append(v22, 4294967295)
                                                                                var v43 bool = false
                                                                                _ = v43
                                                                            case "\n\t":
                                                                                v42 += v42
                                                                            default:
                                                                                if ("\n\t" <= "") {
                                                                                    v22 = append(v22, (v28 + (^(fi16))))
                                                                                    var v44 int32 = 42
                                                                                    _ = v44
                                                                                }
                                                                            }
                                                                            if len(v22) >= 2 {
                                                                                v22 = v22[1:len(v22)]
                                                                            }
                                                                            _ = v42
                                                                        }
                                                                    } else {
                                                                        switch fi16 {
                                                                        case 128:
                                                                            v45 := GenS5{F0: v29, F1: (!true), F2: false}
                                                                            _ = v45
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            var v46 uint64 = 0
                                                                            _ = v46
                                                                        case "":
                                                                            v37 = (v37 * v37)
                                                                        case "\n\t":
                                                                            v47 := GenS6{F0: ("foo bar" + "a"), F1: 136}
                                                                            _ = v47
                                                                        }
                                                                    }
                                                                }
                                                                _ = v39
                                                                _ = v40
                                                            } else {
                                                                if true {
                                                                    var v48 uint32 = 2
                                                                    _ = v48
                                                                } else {
                                                                    if len(v22) > 0 {
                                                                        v22[0] = (^(v28))
                                                                    }
                                                                    switch v28 {
                                                                    case 0:
                                                                        v49 := GenS7{F0: v37, F1: 2147483647}
                                                                        _ = v49
                                                                    case -408:
                                                                        var v50 string = "hello"
                                                                        _ = v50
                                                                    case 669:
                                                                        var v51 int8 = v21
                                                                        _ = v51
                                                                    default:
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            if true {
                                                                                v29 = 1000
                                                                            }
                                                                        }
                                                                        var v52 int8 = v21
                                                                        var v53 uint32 = 536
                                                                        v37 -= v37
                                                                        _ = v52
                                                                        _ = v53
                                                                    }
                                                                }
                                                                fmt.Println(fi16, fi16)
                                                                if false {
                                                                    fmt.Println(v28, v28)
                                                                    v29 = (v29 * v29)
                                                                    var v54 float64 = -2.25
                                                                    _ = v54
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v36
                                                _ = v37
                                            } else {
                                                fmt.Println(v28, fi16)
                                            }
                                        } else {
                                            v55 := GenS8{F0: "a", F1: 4294967295, F2: 2147483647}
                                            v56 := GenS1{F0: -2147483648}
                                            _ = v55
                                            _ = v56
                                        }
                                    default:
                                        var v57 uint = 223
                                        _ = v57
                                    }
                                default:
                                    switch fi16 {
                                    case -9223372036854775808:
                                        var v58 bool = false
                                        _ = v58
                                    case 1:
                                        var v59 uint64 = genFn0(679, 301, 65535)
                                        _ = v59
                                    default:
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                    }
                                }
                            }
                            _ = v29
                        }
                        v21 = 127
                        if true {
                            if (!(1 < genFn0(822, 2147483647, 758))) {
                                fmt.Println(v28)
                                _ = len(v22)
                            }
                            v60 := GenS2{F0: 1, F1: 721}
                            _ = v60
                        } else {
                            v61 := GenS9{F0: "", F1: 1}
                            switch (fi16 + fi16) {
                            case -111:
                                if false {
                                    var v62 uint64 = 4294967295
                                    if len(v22) > 0 {
                                        v22[0] = v28
                                    }
                                    if ((!(v28 <= (-(fi16)))) && true) {
                                        v63 := GenS2{F0: 42, F1: 2}
                                        _ = v63
                                    } else {
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                        switch fi16 {
                                        case -2147483648:
                                            if false {
                                                switch "" {
                                                case "a":
                                                    v62 -= (v62 | (v62 % ((v62) | 1)))
                                                case "\n\t":
                                                    var v64 bool = (1 >= 236)
                                                    _ = v64
                                                case "hello":
                                                    var v65 string = ("\n\t" + "a")
                                                    _ = v65
                                                default:
                                                    var v66 float64 = 0.1
                                                    _ = v66
                                                }
                                                if false {
                                                    if len(v22) >= 2 {
                                                        v22 = v22[1:len(v22)]
                                                    }
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v67 float32 = 0.1
                                                        _ = v67
                                                    }
                                                }
                                                var v68 int8 = v21
                                                _ = v68
                                            }
                                            if true {
                                                v62 += (v62 * (v62 / ((v62) | 1)))
                                                var v69 bool = (-1 <= 42)
                                                _ = v69
                                            } else {
                                                v62 = v62
                                                fmt.Println(fi20)
                                                v22 = append(v22, fi16)
                                                var v70 int16 = -1
                                                v70 += (v70 | (v70 >> ((99) & 63)))
                                                _ = v70
                                            }
                                        case 127:
                                            var v71 uint32 = 4294967295
                                            _ = v71
                                        }
                                        v28 = fi16
                                    }
                                    var v72 int32 = 674
                                    _ = v62
                                    _ = v72
                                }
                            }
                            _ = v61
                        }
                        v21 = v21
                        switch (-(fi16)) {
                        case -2147483648:
                            fmt.Println(fi16, v28, v21)
                        case 4294967295:
                            var v73 int8 = 127
                            _ = v73
                        case 127:
                            v21 += v21
                        default:
                            v21 = 0
                        }
                        _ = v21
                        _ = v22
                        _ = v28
                    } else {
                        var v74 string = ("hello" + "")
                        fmt.Println(fi20)
                        _ = v74
                    }
                    _ = fi20
                }
                var v75 uint8 = 1
                _ = v19
                _ = v75
            }
            _ = fi16
        }
    default:
        v76 := []int{32768}
        _ = v76
    }
    var v77 uint32 = 540
    var v78 uint8 = 128
    _ = v10
    _ = v77
    _ = v78
    return 2
}

func main() {
    var v1 int16 = -194
    var v2 int16 = -293
    var v3 uint64 = 985
    var v4 uint64 = v3
    v5 := []string{"foo bar", "", ""}
    for fi79 := 0; fi79 < 9; fi79++ {
        v80 := GenS2{F0: 42, F1: 1}
        for fi81 := 0; fi81 < 17; fi81++ {
            if true {
                var v82 int = (fi81 * 190)
                var v83 uint32 = 2147483647
                v3 = (v4 + (v3 + v3))
                _ = v82
                _ = v83
            }
            var v84 bool = false
            var v85 bool = true
            _ = fi81
            _ = v84
            _ = v85
        }
        _ = fi79
        _ = v80
    }
    _ = len(v5)
    var v86 uint64 = (v4 % ((v3) | 1))
    v1 -= -968
    v87 := GenS8{F0: "a", F1: genFn0(701, -322, 42), F2: 42}
    v88 := GenS10{F0: 2147483647, F1: 2}
    _ = len(v5)
    for fi89 := 0; fi89 < 5; fi89++ {
        if ((genFn0(18446744073709551615, 42, 881) > (v3 << ((18446744073709551615) & 63))) || false) {
            v86 += (^(v86))
        }
        var v90 int16 = -926
        switch (fi89 >> ((676) & 63)) {
        case -2147483648:
            if (v90 < (v1 / ((v2) | 1))) {
                var v91 uint = 64
                var v92 uint8 = 255
                _ = v91
                _ = v92
            } else {
                var v93 uint8 = 128
                _ = v93
            }
        default:
            v94 := GenS8{F0: "hello", F1: v4, F2: 42}
            _ = v94
        }
        _ = fi89
        _ = v90
    }
    v1 += v2
    var v95 int8 = -128
    v4 -= 866
    v2 = v1
    switch ("a" + "foo bar") {
    case "":
        _ = len(v5)
    }
    switch "" {
    case "":
        v96 := GenS11{F0: 1, F1: (v4 * v3), F2: 2}
        _ = v96
    case "a":
        for fi97 := 0; fi97 < 15; fi97++ {
            v3 -= v4
            v4 = (v86 << ((710) & 63))
            fmt.Println(v2)
            _ = fi97
        }
    }
    fmt.Println(v2, v1)
    if len(v5) > 0 {
        v5[0] = ""
    }
    v1 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v86:", v86)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v86
    _ = v87
    _ = v88
    _ = v95
}
