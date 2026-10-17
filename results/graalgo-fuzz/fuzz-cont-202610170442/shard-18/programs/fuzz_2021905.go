package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 int; F1 int32 }
type GenS10 struct { F0 string; F1 int16; F2 uint32 }
type GenS11 struct { F0 uint32; F1 uint8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int64; F1 bool }
func genFn0() uint16 {
    if (88 >= 109) {
        v7 := GenS1{F0: 0, F1: 100}
        switch 9223372036854775807 {
        case -751:
            var v8 int16 = 32767
            _ = v8
        case 4294967295:
            switch 2147483647 {
            case -559:
                for fi9 := 0; fi9 < 10; fi9++ {
                    fmt.Println(fi9)
                    var v10 int8 = 42
                    var v11 int32 = 912
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case 32768:
                var v12 int = -1
                _ = v12
            case -379:
                v13 := GenS2{F0: 0}
                _ = v13
            }
        case -127:
            var v14 uint = 441
            _ = v14
        }
        var v15 bool = true
        _ = v15
        _ = v7
    } else {
        var v16 int = -1
        if (!false) {
            v17 := []int{}
            var v18 uint = 309
            _ = v17
            _ = v18
        } else {
            if (false || false) {
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 int = (v16 * v16)
                    v20 = (v16 ^ (v16 << ((120) & 63)))
                    _ = fi19
                    _ = v20
                }
            } else {
                v21 := GenS1{F0: 1, F1: -1}
                var v22 int64 = 627
                _ = v21
                _ = v22
            }
            v23 := []int{(v16 % ((v16) | 1)), v16}
            v24 := GenS3{F0: 1, F1: 655, F2: -39}
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 uint64 = 855
            if true {
                for fi27 := 0; fi27 < 3; fi27++ {
                    if true {
                        var v28 int8 = -128
                        v29 := []int{32767}
                        v26 -= (^(v26))
                        _ = v28
                        _ = v29
                    } else {
                        var v30 uint16 = 255
                        if true {
                            v16 = 127
                        } else {
                            v31 := GenS2{F0: v26}
                            _ = v31
                        }
                        _ = v30
                    }
                    v32 := []int{v16, fi27}
                    _ = fi27
                    _ = v32
                }
            }
            var v33 uint32 = uint32(v26)
            _ = fi25
            _ = v26
            _ = v33
        }
        _ = v16
    }
    var v34 int = 838
    if (!false) {
        v34 -= v34
    } else {
        fmt.Println(v34)
        switch (v34 >> ((4294967295) & 63)) {
        case -826:
            switch v34 {
            case 42:
                v35 := []int{v34, v34, v34, (-(v34))}
                _ = v35
            case 32767:
                var v36 int64 = -9223372036854775808
                _ = v36
            case 2147483647:
                for fi37 := 0; fi37 < 17; fi37++ {
                    v38 := []int{(fi37 & (v34 + (v34 % ((fi37) | 1)))), fi37, fi37, (fi37 | v34)}
                    var v39 uint32 = 42
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            default:
                switch v34 {
                case 157:
                    if (2 == 0) {
                        for fi40 := 0; fi40 < 12; fi40++ {
                            v34 -= (fi40 / (((fi40 | (v34 + fi40))) | 1))
                            _ = fi40
                        }
                        fmt.Println(v34)
                        v34 -= 209
                    }
                }
            }
        }
    }
    for fi41 := 0; fi41 < 2; fi41++ {
        v42 := []int{146}
        v43 := GenS4{F0: 128, F1: 348}
        _ = fi41
        _ = v42
        _ = v43
    }
    v44 := GenS4{F0: 127, F1: 674}
    _ = v34
    _ = v44
    return 42
}

func main() {
    v1 := []uint32{1000, 65535, 312}
    var v2 int16 = 1
    v3 := make(map[uint32]string)
    var v4 string = "foo bar"
    var v5 bool = ((-2.25 * -2.25) <= 1.5)
    v6 := make(map[int]string)
    for fi45 := 0; fi45 < 12; fi45++ {
        if (true || false) {
            if (1 >= 127) {
                var v46 int64 = -1
                var v47 int8 = 57
                for fi48 := 0; fi48 < 7; fi48++ {
                    fmt.Println(v47)
                    if (((1e10 / 3.14159) <= 0.0) && ((v5 && v5) || ((v2 == (-(v2))) && v5))) {
                        var v49 int = (fi45 * -586)
                        var v50 float32 = 1e10
                        v2 = v2
                        _ = v49
                        _ = v50
                    }
                    _ = fi48
                }
                _ = v46
                _ = v47
            } else {
                fmt.Println(v2)
                var v51 int32 = -845
                v2 = v2
                _ = v51
            }
            var v52 int16 = (v2 | (v2 << ((255) & 63)))
            v2 -= (v2 | v52)
            _ = v52
        } else {
            fmt.Println(fi45)
            v5 = v5
        }
        switch (-(fi45)) {
        case 1:
            v53 := GenS5{F0: 1, F1: 1, F2: (v2 * v2)}
            _ = v53
        case 32767:
            v4 = v4
        default:
            var v54 bool = v5
            _ = v54
        }
        _ = fi45
    }
    fmt.Println(v4, v5, v4)
    var v55 string = v4
    _ = len(v3)
    for fi56 := 0; fi56 < 15; fi56++ {
        var v57 int8 = 1
        switch ("foo bar" + (v4 + v4)) {
        case "foo bar":
            var v58 float32 = (1.5 - ((1e10 - ((1e10 - (1.5 - (float32(1.5) + 1.5))) + ((1e10 / -2.25) / 1.5))) - 0.1))
            _ = v58
        case "":
            if v5 {
                switch v4 {
                case "a":
                    for fi59 := 0; fi59 < 18; fi59++ {
                        v4 = v4
                        if false {
                            v60 := GenS6{F0: -9223372036854775808}
                            var v61 float32 = 0.1
                            _ = v60
                            _ = v61
                        } else {
                            v57 += (v57 - v57)
                            v62 := GenS7{F0: 138, F1: v55}
                            _ = v62
                        }
                        switch (fi56 * fi59) {
                        case -874:
                            _ = len(v3)
                        case -2147483648:
                            v57 += v57
                        case -459:
                            var v63 int32 = -818
                            _ = v63
                        default:
                            v64 := v3[2147483647]
                            _ = v64
                        }
                        _ = fi59
                    }
                case "":
                    switch v4 {
                    case "a":
                        for fi65 := 0; fi65 < 1; fi65++ {
                            switch "\n\t" {
                            case "\n\t":
                                v2 = (v2 * v2)
                            case "hello":
                                if (0 == -1) {
                                    v2 = (-(v2))
                                    var v66 float64 = (1e10 + float64(1e10))
                                    _ = v66
                                } else {
                                    v2 += 1000
                                    var v67 uint64 = 0
                                    _ = v67
                                }
                            case "":
                                v5 = ((v57 << ((1) & 63)) >= v57)
                            }
                            _ = fi65
                        }
                    }
                case "hello":
                    v68 := GenS8{F0: 72, F1: v57}
                    _ = v68
                }
                v4 = ("hello" + v4)
                v55 = "a"
            } else {
                _ = len(v6)
                var v69 float32 = 1.5
                _ = v69
            }
        case "\n\t":
            var v70 int64 = -857
            _ = v70
        default:
            v71 := GenS3{F0: 42, F1: 0, F2: int32(v2)}
            _ = v71
        }
        _ = fi56
        _ = v57
    }
    switch v4 {
    case "foo bar":
        var v72 bool = v5
        _ = v72
    case "\n\t":
        v1 = append(v1, 65535)
    case "hello":
        v73 := GenS9{F0: 2147483647, F1: 0}
        _ = v73
    }
    if (1000 >= v2) {
        if (!(!((!v5) && ((((-2.25 / -2.25) - -2.25) - 1e10) > float64(3.14159))))) {
            switch v4 {
            case "foo bar":
                for fi74 := 0; fi74 < 4; fi74++ {
                    v5 = v5
                    fmt.Println(fi74, v55, fi74)
                    _ = fi74
                }
            case "":
                var v75 uint8 = 68
                _ = v75
            }
            switch (("\n\t" + (((v4 + v4) + v4) + v55)) + "foo bar") {
            case "":
                v76 := GenS10{F0: "", F1: -1, F2: 306}
                _ = v76
            case "hello":
                v77 := GenS11{F0: 65535, F1: 128}
                _ = v77
            case "\n\t":
                v78 := GenS12{F0: 230}
                _ = v78
            default:
                var v79 int = 9223372036854775807
                _ = v79
            }
        } else {
            var v80 uint16 = 32767
            var v81 uint64 = 1
            fmt.Println(v2, v2, v2)
            _ = v80
            _ = v81
        }
        if (v5 && ((("a" + "hello") <= v4) || (!v5))) {
            switch v4 {
            case "\n\t":
                for fi82 := 0; fi82 < 9; fi82++ {
                    if v5 {
                        _ = len(v1)
                        for fi83 := 0; fi83 < 4; fi83++ {
                            v84 := v3[910]
                            _ = fi83
                            _ = v84
                        }
                        if v5 {
                            switch (((v4 + ((("a" + v4) + "\n\t") + v4)) + ("hello" + v4)) + (v55 + "a")) {
                            case "":
                                if false {
                                    var v85 int64 = 9223372036854775807
                                    for fi86 := 0; fi86 < 17; fi86++ {
                                        _ = len(v1)
                                        _ = fi86
                                    }
                                    _ = v85
                                }
                            case "hello":
                                v2 = 262
                            default:
                                v55 = ""
                            }
                        }
                    }
                    var v87 uint16 = 65535
                    _ = fi82
                    _ = v87
                }
            case "hello":
                fmt.Println(v5)
            case "a":
                for fi88 := 0; fi88 < 18; fi88++ {
                    if v5 {
                        if (-2.25 == -2.25) {
                            switch fi88 {
                            case 128:
                                var v89 int = (fi88 + (^(fi88)))
                                _ = v89
                            default:
                                v90 := GenS10{F0: ("a" + ""), F1: 32767, F2: 444}
                                _ = v90
                            }
                            v91 := v6[fi88]
                            v55 = v4
                            _ = v91
                        }
                        var v92 uint = 393
                        for fi93 := 0; fi93 < 20; fi93++ {
                            if (0 < -9223372036854775808) {
                                var v94 int16 = (v2 | (v2 * (v2 - 1000)))
                                if (180 == 62) {
                                    if (495 >= 42) {
                                        fmt.Println(v55, v94, fi88)
                                        v95 := GenS13{F0: 1000}
                                        _ = v95
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v94
                            } else {
                                var v96 bool = v5
                                v6[fi93] = "foo bar"
                                _ = v96
                            }
                            _ = fi93
                        }
                        _ = v92
                    } else {
                        var v97 uint64 = 1
                        _ = v97
                    }
                    _ = fi88
                }
            }
        } else {
            v55 = "\n\t"
            _ = len(v3)
            var v98 float64 = 3.14159
            _ = v98
        }
    } else {
        _ = len(v3)
    }
    fmt.Println(v4, v2)
    v5 = (!(2 < 2))
    v5 = (!(-536 >= -1))
    if v5 {
        v4 = ("foo bar" + (v4 + ((v4 + "hello") + (v55 + v4))))
    }
    switch 9223372036854775807 {
    case 127:
        v99 := GenS14{F0: 0, F1: (v5 || v5)}
        _ = v99
    case 42:
        v100 := GenS3{F0: 2147483647, F1: 698, F2: 591}
        _ = v100
    case -484:
        _ = len(v6)
    }
    switch 713 {
    case -9223372036854775808:
        v55 = v55
    default:
        var v101 int32 = 0
        _ = v101
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
