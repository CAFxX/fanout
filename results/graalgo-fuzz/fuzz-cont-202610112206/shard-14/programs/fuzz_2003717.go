package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int32; F1 uint64; F2 int32 }
type GenS5 struct { F0 int32; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 int16; F1 uint8; F2 uint }
type GenS8 struct { F0 int; F1 uint32; F2 int16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int16; F1 int }
type GenS13 struct { F0 uint; F1 int32 }
func genFn0() uint16 {
    v6 := []int{128, -9223372036854775808, 860, 42}
    var v7 uint16 = 2
    v7 -= v7
    var v8 float32 = float32(3.14159)
    _ = v6
    _ = v7
    _ = v8
    return v7
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "foo bar"
    var v3 int8 = 100
    var v4 int16 = -32768
    var v5 int32 = -2147483648
    v9 := GenS1{F0: ((v2 + v2) + "foo bar"), F1: (v3 + -128)}
    delete(v1, -1)
    if true {
        v3 = -17
        switch v2 {
        case "":
            if (128 <= 127) {
                var v10 int16 = 1
                for fi11 := 0; fi11 < 16; fi11++ {
                    v12 := GenS2{F0: v4}
                    if ((32767 > 847) && false) {
                        var v13 int16 = (v4 & (^(v10)))
                        if ((!(1 == -1)) || ((v2 + v2) <= "hello")) {
                            v10 = (v4 * (v10 ^ 42))
                        } else {
                            v14 := GenS2{F0: v4}
                            v10 = (-(v13))
                            _ = v14
                        }
                        if (true && (v3 <= int8(v5))) {
                            v1[fi11] = ((v2 + (v2 + "")) + "hello")
                            for fi15 := 0; fi15 < 2; fi15++ {
                                v4 += (v4 % ((0) | 1))
                                if true {
                                    v16 := GenS2{F0: 434}
                                    if true {
                                        fmt.Println(v10, v13, v2)
                                        v17 := GenS2{F0: (v13 ^ v10)}
                                        _ = v17
                                    }
                                    v18 := []int{}
                                    _ = v16
                                    _ = v18
                                }
                                _ = fi15
                            }
                            switch 1 {
                            case 32767:
                                v3 = v3
                            case 2147483647:
                                v19 := GenS1{F0: "hello", F1: (v3 & -114)}
                                _ = v19
                            case 871:
                                if false {
                                    if (489 == 639) {
                                        if false {
                                            v20 := GenS3{F0: -1, F1: fi11}
                                            _ = v20
                                        }
                                        var v21 int = fi11
                                        _ = v21
                                    } else {
                                        v13 += v10
                                        if false {
                                            if (607 <= 2147483647) {
                                                for fi22 := 0; fi22 < 4; fi22++ {
                                                    if true {
                                                        if false {
                                                            if (42 <= 8) {
                                                                v23 := GenS4{F0: (^(v5)), F1: 841, F2: v5}
                                                                v24 := v1[128]
                                                                _ = v23
                                                                _ = v24
                                                            }
                                                            if false {
                                                                if (false && (904 == 4294967295)) {
                                                                    fmt.Println(fi11, fi11)
                                                                }
                                                                fmt.Println(fi22)
                                                                var v25 uint = 255
                                                                _ = v25
                                                            }
                                                        }
                                                        var v26 uint32 = 0
                                                        switch 128 {
                                                        case 2147483647:
                                                            switch "a" {
                                                            case "hello":
                                                                v2 = (v2 + ((v2 + ("foo bar" + ("foo bar" + ((v2 + v2) + "\n\t")))) + v2))
                                                            case "":
                                                                v27 := v1[fi11]
                                                                _ = v27
                                                            case "a":
                                                                v13 -= v13
                                                            }
                                                        case 85:
                                                            v28 := GenS4{F0: v5, F1: 365, F2: v5}
                                                            _ = v28
                                                        case 9223372036854775807:
                                                            var v29 int16 = 42
                                                            _ = v29
                                                        }
                                                        v30 := GenS5{F0: 358, F1: 2, F2: 42}
                                                        _ = v26
                                                        _ = v30
                                                    } else {
                                                        switch -9223372036854775808 {
                                                        case 0:
                                                            var v31 uint = 0
                                                            _ = v31
                                                        default:
                                                            v32 := GenS6{F0: 128, F1: (v3 + (v3 | -1))}
                                                            _ = v32
                                                        }
                                                        var v33 int64 = 9223372036854775807
                                                        _ = v33
                                                    }
                                                    var v34 bool = false
                                                    _ = fi22
                                                    _ = v34
                                                }
                                                switch ((("a" + ("\n\t" + "foo bar")) + (("" + v2) + v2)) + v2) {
                                                case "\n\t":
                                                    var v35 bool = false
                                                    _ = v35
                                                }
                                                switch (fi11 ^ (fi11 & -48)) {
                                                case 128:
                                                    for fi36 := 0; fi36 < 10; fi36++ {
                                                        v10 += v10
                                                        var v37 float64 = (((3.14159 / 1.5) - 1e10) + 3.14159)
                                                        _ = fi36
                                                        _ = v37
                                                    }
                                                }
                                            } else {
                                                v13 = (v10 * v10)
                                            }
                                        } else {
                                            for fi38 := 0; fi38 < 14; fi38++ {
                                                if ((!false) || false) {
                                                    fmt.Println(v13, v10, fi11)
                                                }
                                                if (!(false || true)) {
                                                    v39 := GenS1{F0: v2, F1: v3}
                                                    if false {
                                                        var v40 uint8 = 255
                                                        switch "" {
                                                        case "foo bar":
                                                            fmt.Println(v13, v4)
                                                        case "a":
                                                            v2 = "foo bar"
                                                        case "hello":
                                                            v41 := []int{(fi11 * 32767), (fi38 & fi11)}
                                                            _ = v41
                                                        default:
                                                            switch "" {
                                                            case "\n\t":
                                                                fmt.Println(v40, v13)
                                                            default:
                                                                v42 := GenS4{F0: v5, F1: 29, F2: v5}
                                                                v43 := v1[fi11]
                                                                fmt.Println(v10)
                                                                _ = v42
                                                                _ = v43
                                                            }
                                                        }
                                                        _ = v40
                                                    }
                                                    if false {
                                                        v44 := GenS7{F0: v10, F1: 127, F2: 4294967295}
                                                        var v45 int16 = (v4 - 1000)
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    _ = v39
                                                } else {
                                                    var v46 string = "hello"
                                                    v47 := GenS8{F0: fi11, F1: 2, F2: v10}
                                                    _ = v46
                                                    _ = v47
                                                }
                                                _ = fi38
                                            }
                                            for fi48 := 0; fi48 < 14; fi48++ {
                                                v49 := GenS9{F0: v13}
                                                var v50 int8 = -97
                                                v10 += 32767
                                                _ = fi48
                                                _ = v49
                                                _ = v50
                                            }
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            if false {
                                                v13 = v10
                                                switch (fi11 | 32767) {
                                                case 9223372036854775807:
                                                    for fi51 := 0; fi51 < 4; fi51++ {
                                                        switch (fi11 + 615) {
                                                        case 9223372036854775807:
                                                            v5 += 0
                                                        case -49:
                                                            switch (v2 + (v2 + ("" + v2))) {
                                                            case "hello":
                                                                v5 += v5
                                                            case "a":
                                                                fmt.Println(v4, v13, v2)
                                                            case "\n\t":
                                                                var v52 string = ((v2 + ("\n\t" + v2)) + (v2 + (v2 + "\n\t")))
                                                                _ = v52
                                                            }
                                                        }
                                                        v3 -= v3
                                                        _ = fi51
                                                    }
                                                case -2147483648:
                                                    if true {
                                                        for fi53 := 0; fi53 < 18; fi53++ {
                                                            var v54 int8 = (v3 & 0)
                                                            _ = fi53
                                                            _ = v54
                                                        }
                                                        for fi55 := 0; fi55 < 15; fi55++ {
                                                            v3 = v3
                                                            _ = fi55
                                                        }
                                                        for fi56 := 0; fi56 < 2; fi56++ {
                                                            v57 := v1[-9223372036854775808]
                                                            var v58 int64 = -20
                                                            var v59 uint = 207
                                                            var v60 int32 = (v5 << ((v59) & 63))
                                                            _ = fi56
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    }
                                                default:
                                                    for fi61 := 0; fi61 < 3; fi61++ {
                                                        if true {
                                                            v2 = (v2 + v2)
                                                            var v62 uint64 = 636
                                                            _ = v62
                                                        }
                                                        _ = fi61
                                                    }
                                                }
                                            } else {
                                                v4 += -699
                                                if (true && false) {
                                                    if (!true) {
                                                        for fi63 := 0; fi63 < 3; fi63++ {
                                                            switch ("a" + ("hello" + v2)) {
                                                            case "hello":
                                                                v5 = int32(v10)
                                                            }
                                                            v64 := GenS10{F0: 1, F1: 2}
                                                            _ = fi63
                                                            _ = v64
                                                        }
                                                    }
                                                    v1[(fi11 ^ -389)] = (("foo bar" + v2) + v2)
                                                    for fi65 := 0; fi65 < 13; fi65++ {
                                                        var v66 float32 = 0.0
                                                        fmt.Println(v2, v2)
                                                        v67 := []int{fi11, int(v13), 1}
                                                        _ = fi65
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                } else {
                                                    v10 -= v4
                                                }
                                            }
                                        case "\n\t":
                                            fmt.Println(v4, v2)
                                        case "":
                                            for fi68 := 0; fi68 < 3; fi68++ {
                                                if false {
                                                    v3 -= -1
                                                    if true {
                                                        if (!true) {
                                                            v4 = v13
                                                            var v69 bool = false
                                                            v4 += (v4 % ((v4) | 1))
                                                            _ = v69
                                                        } else {
                                                            fmt.Println(fi11, v3)
                                                        }
                                                        switch fi68 {
                                                        case 0:
                                                            v1[-256] = (v2 + v2)
                                                        case 9223372036854775807:
                                                            var v70 uint = 214
                                                            _ = v70
                                                        default:
                                                            var v71 float32 = 0.1
                                                            _ = v71
                                                        }
                                                        v72 := []int{(fi11 | fi11), fi11, 270, 685}
                                                        _ = v72
                                                    }
                                                    _ = len(v1)
                                                }
                                                _ = fi68
                                            }
                                        }
                                    }
                                    v1[(fi11 & (fi11 | fi11))] = (v2 + ("\n\t" + (v2 + "")))
                                    v73 := GenS11{F0: (-(v3))}
                                    _ = v73
                                } else {
                                    var v74 float64 = 3.14159
                                    _ = v74
                                }
                            default:
                                var v75 bool = true
                                _ = v75
                            }
                        } else {
                            switch (fi11 >> ((4294967295) & 63)) {
                            case 32768:
                                fmt.Println(v10)
                            default:
                                v1[fi11] = (v2 + "hello")
                            }
                            switch 536 {
                            case -1:
                                for fi76 := 0; fi76 < 8; fi76++ {
                                    if (v4 < (v13 + 42)) {
                                        var v77 uint = 1
                                        fmt.Println(v4)
                                        v77 += 442
                                        _ = v77
                                    } else {
                                        if ((4294967295 != 2) && false) {
                                            v78 := []int{(fi76 + (fi76 * fi76)), (fi76 | 42), (^(fi76))}
                                            _ = v78
                                        } else {
                                            var v79 uint64 = 2
                                            v2 = v2
                                            _ = v79
                                        }
                                        v80 := []int{}
                                        _ = v80
                                    }
                                    v5 = v5
                                    v4 = v4
                                    if false {
                                        v1[fi76] = v2
                                    } else {
                                        var v81 uint8 = 255
                                        _ = v81
                                    }
                                    v3 = v3
                                    _ = fi76
                                }
                            }
                        }
                        _ = v13
                    }
                    if false {
                        v2 = v2
                        if false {
                            for fi82 := 0; fi82 < 7; fi82++ {
                                v83 := []int{(fi82 << ((172) & 63))}
                                fmt.Println(v4, v3, v4)
                                switch 2147483647 {
                                case 32768:
                                    v84 := GenS12{F0: 1000, F1: (fi82 ^ (fi11 / (((fi82 >> ((255) & 63))) | 1)))}
                                    _ = v84
                                case 128:
                                    var v85 int = (fi11 % ((fi11) | 1))
                                    _ = v85
                                case 32767:
                                    if len(v83) >= 2 {
                                        v83 = v83[1:len(v83)]
                                    }
                                default:
                                    var v86 uint64 = 755
                                    _ = v86
                                }
                                _ = fi82
                                _ = v83
                            }
                            for fi87 := 0; fi87 < 14; fi87++ {
                                v88 := []int{-1, (fi11 >> ((2) & 63)), -556}
                                v89 := GenS7{F0: v4, F1: 127, F2: 4294967295}
                                _ = fi87
                                _ = v88
                                _ = v89
                            }
                        }
                        v90 := GenS8{F0: (fi11 / ((0) | 1)), F1: 42, F2: v4}
                        _ = v90
                    }
                    _ = fi11
                    _ = v12
                }
                v91 := GenS13{F0: 2, F1: v5}
                _ = v10
                _ = v91
            } else {
                var v92 int = -646
                _ = len(v1)
                fmt.Println(v2, v5, v2)
                _ = v92
            }
        case "\n\t":
            for fi93 := 0; fi93 < 19; fi93++ {
                if false {
                    for fi94 := 0; fi94 < 11; fi94++ {
                        v3 = v3
                        fmt.Println(v3, v2, fi94)
                        _ = fi94
                    }
                    v5 += -508
                }
                _ = fi93
            }
        }
    }
    v5 += -2147483648
    v95 := GenS7{F0: v4, F1: 50, F2: 765}
    v96 := v1[128]
    v97 := v1[4294967295]
    var v98 string = v96
    v96 = ((("a" + "\n\t") + (v2 + ((v98 + ("a" + v97)) + ((v97 + v96) + v97)))) + v2)
    var v99 int16 = v4
    v100 := []int{4294967295, 2147483647, 42, 1}
    delete(v1, 127)
    v1[32768] = "\n\t"
    var v101 float64 = -2.25
    var v102 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v100:", len(v100), v100)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
    _ = v95
    _ = v96
    _ = v97
    _ = v98
    _ = v99
}
