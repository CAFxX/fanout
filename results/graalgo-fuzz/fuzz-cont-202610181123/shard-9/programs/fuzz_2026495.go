package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 string; F1 string; F2 uint8 }
type GenS10 struct { F0 uint16; F1 int32 }
type GenS11 struct { F0 uint16; F1 string }
type GenS12 struct { F0 uint8; F1 uint32 }
type GenS13 struct { F0 uint64; F1 int64; F2 int32 }
func genFn0(p0 int16, p1 int32, p2 uint64) uint8 {
    v4 := []int{127, -431}
    switch ("hello" + "\n\t") {
    case "":
        if true {
            for fi5 := 0; fi5 < 16; fi5++ {
                var v6 int32 = p1
                _ = fi5
                _ = v6
            }
        }
    case "hello":
        if (2 == 51) {
            p1 = (-(p1))
            p2 = (^(p2))
            if (p2 < (^(p2))) {
                var v7 string = ("foo bar" + "")
                _ = v7
            }
        }
    }
    var v8 int64 = 346
    v9 := GenS1{F0: ((!true) && (true || (127 > 219))), F1: (^(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[int]int)
    var v3 uint = 1
    if true {
        for fi10 := 0; fi10 < 17; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                switch "hello" {
                case "":
                    switch "a" {
                    case "hello":
                        v12 := GenS1{F0: ((v3 / ((v3) | 1)) > (v3 % ((2) | 1))), F1: 4294967295}
                        _ = v12
                    case "a":
                        v3 += 65535
                    case "":
                        fmt.Println(v3, fi10)
                    default:
                        var v13 int32 = -2147483648
                        _ = v13
                    }
                    if ((^(v1)) == (^(v1))) {
                        v14 := GenS2{F0: 1}
                        var v15 int = 0
                        if ((v1 ^ (v1 / (((v1 << ((v3) & 63))) | 1))) > v1) {
                            switch fi10 {
                            case -415:
                                v1 = (v1 | (v1 & (^(v1))))
                            case 128:
                                _ = len(v2)
                            }
                            var v16 int8 = -1
                            _ = v16
                        } else {
                            v15 = (^(fi11))
                            var v17 uint = (v3 << ((v3) & 63))
                            var v18 int32 = -29
                            _ = v17
                            _ = v18
                        }
                        _ = v14
                        _ = v15
                    } else {
                        switch ("a" + ((("\n\t" + "") + "") + "foo bar")) {
                        case "":
                            if false {
                                if true {
                                    v1 = v1
                                    v19 := GenS3{F0: 164, F1: 0}
                                    _ = v19
                                }
                                var v20 uint64 = 0
                                v21 := GenS2{F0: 2147483647}
                                _ = v20
                                _ = v21
                            }
                        case "a":
                            var v22 uint64 = 601
                            _ = v22
                        }
                        v3 += 917
                        var v23 int8 = 92
                        _ = v23
                    }
                    v1 = v1
                    fmt.Println(v1, v3)
                case "a":
                    var v24 uint32 = (v1 % (((v1 / ((1) | 1))) | 1))
                    _ = v24
                case "foo bar":
                    switch "a" {
                    case "":
                        var v25 uint32 = 65535
                        _ = v25
                    default:
                        v26 := GenS1{F0: ((42 > -1) && (255 == 65535)), F1: 4294967295}
                        _ = v26
                    }
                default:
                    if false {
                        var v27 uint64 = 42
                        switch fi11 {
                        case -2147483648:
                            var v28 int = fi11
                            v28 += -1
                            _ = v28
                        case 32768:
                            if ((3.14159 * ((3.14159 / ((0.0 / -2.25) + -2.25)) / 1.0)) >= (float64(float32(0.1)) - (float64(float32(float64(1.5))) / 1e10))) {
                                var v29 string = "\n\t"
                                v27 = v27
                                fmt.Println(v27, v3)
                                _ = v29
                            } else {
                                if false {
                                    if true {
                                        switch fi10 {
                                        case 0:
                                            v27 += (v27 - (v27 / ((2) | 1)))
                                        default:
                                            v2[fi11] = 32767
                                        }
                                        switch ("a" + "") {
                                        case "foo bar":
                                            var v30 uint8 = genFn0(32767, 1, (v27 % (((v27 * uint64(v3))) | 1)))
                                            _ = v30
                                        case "":
                                            var v31 uint8 = genFn0(1000, -1, v27)
                                            _ = v31
                                        case "hello":
                                            if (32767 < 255) {
                                                v32 := []int{127, (fi11 << ((v3) & 63)), fi10, (fi11 / (((fi10 & fi10)) | 1))}
                                                v33 := GenS4{F0: (^(v27))}
                                                _ = v32
                                                _ = v33
                                            } else {
                                                fmt.Println(fi10, v27, fi10)
                                                if false {
                                                    if true {
                                                        v2[(fi10 + 0)] = fi10
                                                        if false {
                                                            fmt.Println(v1)
                                                        }
                                                    } else {
                                                        v34 := GenS4{F0: v27}
                                                        _ = v34
                                                    }
                                                    if true {
                                                        delete(v2, 32768)
                                                        v35 := GenS2{F0: 1}
                                                        _ = v35
                                                    }
                                                    if (!false) {
                                                        if false {
                                                            fmt.Println(fi10, v27, v1)
                                                            if true {
                                                                if false {
                                                                    v27 += v27
                                                                    v27 = (v27 << ((v3) & 63))
                                                                } else {
                                                                    fmt.Println(v1, v1)
                                                                }
                                                                var v36 int64 = 648
                                                                switch (fi10 & (-(fi11))) {
                                                                case -1:
                                                                    v37 := GenS3{F0: v36, F1: -721}
                                                                    var v38 bool = false
                                                                    _ = v37
                                                                    _ = v38
                                                                case -9223372036854775808:
                                                                    if (v27 > (v27 % ((816) | 1))) {
                                                                        var v39 uint = (v3 << ((947) & 63))
                                                                        _ = v39
                                                                    } else {
                                                                        v40 := GenS5{F0: 142, F1: 1000}
                                                                        v36 = (v36 >> (((v3 >> (((v3 / ((955) | 1))) & 63))) & 63))
                                                                        _ = v40
                                                                    }
                                                                }
                                                                _ = v36
                                                            }
                                                            fmt.Println(fi11)
                                                        } else {
                                                            v1 = (v1 % ((v1) | 1))
                                                            var v41 int64 = -315
                                                            v1 += uint32(v3)
                                                            _ = v41
                                                        }
                                                        var v42 uint64 = (v27 - (v27 % ((382) | 1)))
                                                        v3 += (v3 & v3)
                                                        _ = v42
                                                    }
                                                    if true {
                                                        var v43 int16 = 1
                                                        _ = v43
                                                    }
                                                }
                                            }
                                        default:
                                            delete(v2, fi11)
                                        }
                                    } else {
                                        v1 += 4294967295
                                        if false {
                                            var v44 bool = ((^(v3)) == v3)
                                            _ = v44
                                        } else {
                                            var v45 float64 = 0.0
                                            v46 := GenS6{F0: 255, F1: 65535, F2: int64(v3)}
                                            var v47 int16 = 1
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        }
                                        v48 := GenS7{F0: v27}
                                        _ = v48
                                    }
                                } else {
                                    var v49 uint16 = 1
                                    _ = v49
                                }
                                fmt.Println(v3, v27)
                                v50 := []int{fi11, (fi11 >> ((uint(v27)) & 63)), fi10}
                                switch int(v27) {
                                case 0:
                                    var v51 string = (("foo bar" + "foo bar") + "foo bar")
                                    _ = v51
                                default:
                                    v52 := GenS6{F0: 42, F1: 665, F2: 0}
                                    _ = v52
                                }
                                _ = v50
                            }
                        default:
                            switch (("" + (("foo bar" + (("\n\t" + ("foo bar" + "hello")) + (("foo bar" + ("a" + "hello")) + (("a" + ("a" + ("a" + "a"))) + "foo bar")))) + "\n\t")) + "\n\t") {
                            case "hello":
                                v1 = 2147483648
                            case "foo bar":
                                v3 = v3
                            case "a":
                                if false {
                                    v3 += (^(v3))
                                }
                            default:
                                fmt.Println(v27, fi11)
                            }
                        }
                        v3 = v3
                        _ = v27
                    } else {
                        if ((1 >= -1) && true) {
                            switch "" {
                            case "foo bar":
                                fmt.Println(fi11, fi11)
                            case "a":
                                var v53 uint64 = 0
                                _ = v53
                            default:
                                var v54 uint32 = v1
                                _ = v54
                            }
                        } else {
                            var v55 int = fi10
                            if true {
                                var v56 int32 = -1
                                if true {
                                    v57 := GenS8{F0: -1}
                                    v58 := []int{-2147483648, (-(fi11))}
                                    if false {
                                        var v59 uint64 = 42
                                        if ((v3 % ((v3) | 1)) >= 255) {
                                            var v60 int64 = -150
                                            _ = v60
                                        }
                                        v61 := GenS4{F0: 484}
                                        if (v3 >= 2) {
                                            _ = len(v2)
                                            fmt.Println(v3, fi10)
                                            v62 := GenS3{F0: 212, F1: 445}
                                            if (-32768 != 32767) {
                                                var v63 uint64 = (v59 + v59)
                                                _ = v63
                                            }
                                            var v64 float32 = 1.5
                                            _ = v62
                                            _ = v64
                                        } else {
                                            v65 := v2[(^(fi11))]
                                            var v66 uint16 = 65535
                                            v59 -= (^(v59))
                                            _ = v65
                                            _ = v66
                                        }
                                        _ = v59
                                        _ = v61
                                    }
                                    _ = v57
                                    _ = v58
                                }
                                _ = v56
                            }
                            v1 -= (v1 | v1)
                            _ = v55
                        }
                    }
                }
                v2[fi11] = (fi10 % (((fi11 << ((2) & 63))) | 1))
                v1 += (v1 ^ 370)
                _ = fi11
            }
            _ = fi10
        }
    } else {
        for fi67 := 0; fi67 < 6; fi67++ {
            _ = len(v2)
            fmt.Println(v1, fi67)
            _ = fi67
        }
        v68 := []int{-9223372036854775808}
        fmt.Println(v3, v1)
        _ = v68
    }
    var v69 int8 = -128
    v3 -= v3
    v70 := GenS9{F0: (("\n\t" + "\n\t") + "a"), F1: "\n\t", F2: 135}
    var v71 uint8 = 1
    var v72 string = "hello"
    switch ((v72 + "hello") + v72) {
    case "foo bar":
        if (((!(42 == 18446744073709551615)) && false) || (271 < 1)) {
            fmt.Println(v1, v71, v1)
            var v73 uint = (v3 ^ v3)
            switch ("a" + (v72 + "\n\t")) {
            case "foo bar":
                v71 += (v71 << (((^(v3))) & 63))
            default:
                fmt.Println(v1)
            }
            _ = v73
        }
    case "":
        switch "a" {
        case "hello":
            for fi74 := 0; fi74 < 8; fi74++ {
                v1 += v1
                var v75 int = fi74
                if false {
                    var v76 float32 = 1.5
                    v77 := []int{fi74, v75}
                    _ = v76
                    _ = v77
                } else {
                    if (false || false) {
                        v78 := GenS10{F0: 730, F1: 1}
                        _ = v78
                    } else {
                        for fi79 := 0; fi79 < 20; fi79++ {
                            var v80 int16 = 0
                            v81 := []int{(fi79 + (v75 ^ int(v3))), (fi74 * fi74)}
                            if (587 != 0) {
                                fmt.Println(v80, v69, v75)
                                var v82 uint8 = v71
                                _ = v82
                            }
                            _ = fi79
                            _ = v80
                            _ = v81
                        }
                    }
                    _ = len(v2)
                }
                _ = fi74
                _ = v75
            }
        default:
            if false {
                var v83 uint8 = (v71 | genFn0(-32768, 2147483647, 0))
                var v84 int8 = (v69 >> (((v3 * 2)) & 63))
                _ = v83
                _ = v84
            }
        }
    case "hello":
        var v85 float64 = 0.1
        _ = v85
    default:
        switch 710 {
        case 4294967295:
            v2[2147483647] = 32768
        case 2147483647:
            switch -2147483648 {
            case -311:
                v2[69] = -1
            }
        case 623:
            var v86 uint8 = genFn0(-119, 1, 539)
            _ = v86
        default:
            for fi87 := 0; fi87 < 20; fi87++ {
                var v88 float32 = 1.5
                _ = fi87
                _ = v88
            }
        }
    }
    var v89 int64 = 9223372036854775807
    switch (v72 + "a") {
    case "hello":
        var v90 string = "foo bar"
        _ = v90
    case "a":
        var v91 string = ("" + ("" + v72))
        _ = v91
    case "\n\t":
        var v92 int = -9223372036854775808
        _ = v92
    default:
        if false {
            for fi93 := 0; fi93 < 16; fi93++ {
                if false {
                    if true {
                        v89 += 129
                        for fi94 := 0; fi94 < 13; fi94++ {
                            fmt.Println(fi94, v89, fi94)
                            v95 := GenS11{F0: 814, F1: ("" + "a")}
                            _ = fi94
                            _ = v95
                        }
                        if true {
                            var v96 uint64 = 42
                            fmt.Println(v3)
                            _ = len(v2)
                            _ = v96
                        }
                    }
                } else {
                    for fi97 := 0; fi97 < 9; fi97++ {
                        v98 := GenS4{F0: 859}
                        switch (fi93 >> ((v3) & 63)) {
                        case 4294967295:
                            var v99 int8 = -1
                            _ = v99
                        default:
                            v1 = (v1 >> ((3) & 63))
                        }
                        _ = fi97
                        _ = v98
                    }
                }
                _ = fi93
            }
            switch (v72 + v72) {
            case "hello":
                v89 += (v89 ^ v89)
            case "a":
                fmt.Println(v1, v71)
            }
        } else {
            _ = len(v2)
            if (true || false) {
                fmt.Println(v72, v72, v3)
                _ = len(v2)
                var v100 bool = (-2.25 == 0.1)
                _ = v100
            } else {
                v71 += uint8(v3)
            }
        }
    }
    v101 := []int{int(v1), 755, -9223372036854775808}
    v102 := GenS12{F0: genFn0(-1, 42, 18446744073709551615), F1: (v1 & (^(v1)))}
    v3 = 955
    v69 -= (v69 % (((v69 | v69)) | 1))
    for fi103 := 0; fi103 < 8; fi103++ {
        v104 := GenS13{F0: 1, F1: -9223372036854775808, F2: 0}
        _ = fi103
        _ = v104
    }
    v105 := GenS13{F0: uint64(v3), F1: -293, F2: 777}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v101:", len(v101), v101)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v69:", v69)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v101
    _ = v102
    _ = v105
    _ = v2
    _ = v3
    _ = v69
    _ = v70
    _ = v71
    _ = v72
    _ = v89
}
