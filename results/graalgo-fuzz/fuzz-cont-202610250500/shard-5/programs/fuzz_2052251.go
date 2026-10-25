package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int16; F2 uint }
type GenS5 struct { F0 uint64; F1 int64; F2 string }
type GenS6 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint16, p1 string) uint {
    switch 127 {
    case 32768:
        v6 := []int{32768, 501, 9223372036854775807}
        _ = v6
    default:
        switch ("hello" + (p1 + (p1 + p1))) {
        case "":
            var v7 uint = 42
            _ = v7
        case "\n\t":
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(p0, p0)
                var v9 uint = 65535
                var v10 int16 = 32767
                _ = fi8
                _ = v10
                _ = v9
            }
        }
    }
    var v11 int64 = 0
    fmt.Println(v11, v11)
    var v12 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 0
}
func genFn1(p0 int8, p1 uint64) int {
    var v13 uint32 = 2147483647
    p1 -= (p1 << ((2) & 63))
    var v14 int8 = p0
    var v15 int = 128
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return v15
}

func main() {
    var v1 bool = true
    var v2 uint = 65535
    var v3 uint32 = 2147483648
    v4 := make(map[uint32]string)
    v5 := make(map[string]bool)
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 uint32 = (^(v3))
        var v18 bool = v1
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint64 = 4294967295
    v20 := GenS1{F0: -72, F1: (v2 >> (((v2 ^ 614)) & 63))}
    switch 0 {
    case 42:
        if (-2.25 != 0.0) {
            if (54 == -1) {
                for fi21 := 0; fi21 < 7; fi21++ {
                    v1 = (!(0 != v19))
                    var v22 int16 = 42
                    _ = fi21
                    _ = v22
                }
                if (30 >= -973) {
                    for fi23 := 0; fi23 < 1; fi23++ {
                        for fi24 := 0; fi24 < 1; fi24++ {
                            _ = len(v5)
                            _ = fi24
                        }
                        for fi25 := 0; fi25 < 3; fi25++ {
                            var v26 uint = genFn0(255, ("foo bar" + ("" + "\n\t")))
                            if (!true) {
                                v27 := GenS1{F0: -1, F1: (^(v2))}
                                v19 = 135
                                _ = v27
                            } else {
                                fmt.Println(v19, fi23)
                                v28 := []int{fi23, (fi23 >> ((genFn0(65535, "\n\t")) & 63))}
                                if ("" > ((("foo bar" + "\n\t") + "") + "\n\t")) {
                                    v1 = false
                                } else {
                                    v2 = v2
                                    var v29 uint32 = v3
                                    _ = v29
                                }
                                _ = v28
                            }
                            _ = fi25
                            _ = v26
                        }
                        for fi30 := 0; fi30 < 1; fi30++ {
                            if false {
                                v31 := GenS1{F0: int8(v3), F1: v2}
                                _ = v31
                            } else {
                                v19 += (v19 + v19)
                                switch (fi30 >> (((^(v2))) & 63)) {
                                case -2147483648:
                                    var v32 uint = 65535
                                    v33 := GenS2{F0: 304, F1: v2}
                                    v34 := GenS3{F0: (("\n\t" + ("foo bar" + ("hello" + "hello"))) + ("hello" + "hello"))}
                                    v32 = 287
                                    if false {
                                        v2 -= 255
                                    }
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                                if false {
                                    v19 = v19
                                } else {
                                    var v35 uint16 = 32767
                                    _ = v35
                                }
                            }
                            if true {
                                var v36 int32 = 0
                                _ = v36
                            }
                            var v37 string = (("foo bar" + "foo bar") + "foo bar")
                            _ = fi30
                            _ = v37
                        }
                        _ = fi23
                    }
                } else {
                    for fi38 := 0; fi38 < 4; fi38++ {
                        var v39 float32 = (3.14159 / ((float32(3.14159) * 0.1) - (1.5 / 0.1)))
                        v40 := GenS4{F0: "hello", F1: 624, F2: v2}
                        var v41 bool = v1
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    var v42 string = "foo bar"
                    var v43 uint64 = v19
                    _ = v42
                    _ = v43
                }
                for fi44 := 0; fi44 < 6; fi44++ {
                    for fi45 := 0; fi45 < 20; fi45++ {
                        v46 := []int{int(v2), -515, (fi44 | -861)}
                        _ = len(v46)
                        var v47 int8 = 29
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    _ = fi44
                }
            } else {
                var v48 uint16 = 451
                var v49 int64 = -860
                _ = v48
                _ = v49
            }
            fmt.Println(v19, v3, v1)
            v50 := GenS2{F0: 0, F1: 262}
            _ = v50
        }
    case 128:
        v2 = genFn0(230, "\n\t")
    case -369:
        var v51 int = 32768
        _ = v51
    default:
        switch -9223372036854775808 {
        case 2147483647:
            var v52 uint32 = 1
            _ = v52
        case -116:
            if (v1 && v1) {
                v2 += v2
                v1 = v1
                fmt.Println(v19)
            } else {
                fmt.Println(v2, v2, v1)
            }
        case 32768:
            var v53 uint32 = (^(v3))
            _ = v53
        }
    }
    v3 = (v3 >> ((v2) & 63))
    fmt.Println(v19, v1, v3)
    v54 := []int{-903}
    v54 = append(v54, -308)
    switch "" {
    case "":
        var v55 uint = (^(v2))
        _ = v55
    default:
        var v56 int8 = -23
        _ = v56
    }
    v19 += v19
    v57 := GenS5{F0: (^(v19)), F1: -1, F2: "foo bar"}
    for fi58 := 0; fi58 < 7; fi58++ {
        fmt.Println(v3)
        switch ("a" + "") {
        case "\n\t":
            _ = len(v5)
        }
        if v1 {
            for fi59 := 0; fi59 < 15; fi59++ {
                if ((v3 + v3) != v3) {
                    if ((-2.25 / (1e10 / float64((0.1 - 1.5)))) != (((1e10 * float64(1e10)) * float64(((3.14159 / 3.14159) * ((1.5 * ((1e10 + float32((0.1 * 0.1))) + 1e10)) + (1e10 - (0.1 + 0.0)))))) / 1e10)) {
                        fmt.Println(fi58, v2, v19)
                        switch (fi59 % ((int(v3)) | 1)) {
                        case -408:
                            switch 42 {
                            case 10:
                                v60 := GenS5{F0: v19, F1: 0, F2: ""}
                                _ = v60
                            case 188:
                                var v61 uint16 = 255
                                _ = v61
                            case -330:
                                v62 := v5["a"]
                                _ = v62
                            default:
                                if true {
                                    var v63 uint = genFn0(538, "")
                                    v64 := GenS1{F0: 1, F1: v63}
                                    _ = v63
                                    _ = v64
                                } else {
                                    var v65 int32 = 42
                                    _ = v65
                                }
                            }
                        }
                    } else {
                        if ((v3 + 0) != 1) {
                            v3 = (^(v3))
                        } else {
                            if (!(0.1 == (3.14159 * 3.14159))) {
                                var v66 int32 = 119
                                _ = v66
                            } else {
                                switch ("a" + ((("a" + "a") + "foo bar") + "hello")) {
                                case "":
                                    _ = len(v54)
                                case "a":
                                    var v67 float32 = 3.14159
                                    _ = v67
                                case "\n\t":
                                    v54 = append(v54, fi59)
                                default:
                                    var v68 int = genFn1(-1, (v19 / ((362) | 1)))
                                    _ = v68
                                }
                                var v69 float32 = float32(((((3.14159 * float64((1.5 * 1.5))) - (0.0 + 1e10)) + (float64((0.1 + 1e10)) * ((0.1 + 1.5) * 1.5))) + 3.14159))
                                switch fi58 {
                                case 32768:
                                    v70 := GenS3{F0: ("hello" + ("a" + "foo bar"))}
                                    var v71 int32 = -232
                                    v2 = (v2 / ((0) | 1))
                                    v1 = (-9223372036854775808 > -1)
                                    var v72 uint8 = 42
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                }
                                _ = v69
                            }
                            v73 := GenS6{F0: 4294967295, F1: v19}
                            _ = v73
                        }
                        v19 = (^(v19))
                        switch "\n\t" {
                        case "":
                            v54 = append(v54, (fi59 + (fi59 ^ fi58)))
                        case "a":
                            var v74 int8 = 127
                            _ = v74
                        case "hello":
                            var v75 bool = (225 == v19)
                            _ = v75
                        default:
                            if true {
                                v1 = (v2 > uint(v3))
                                var v76 int32 = 1
                                _ = v76
                            } else {
                                var v77 int = 42
                                var v78 uint8 = 245
                                if true {
                                    v1 = v1
                                    if v1 {
                                        if ((394 <= v19) && v1) {
                                            if ((!true) && (fi59 != 1)) {
                                                v2 = genFn0(175, "")
                                                _ = len(v54)
                                                v1 = true
                                            }
                                        } else {
                                            v1 = true
                                        }
                                    }
                                    if false {
                                        v54 = append(v54, v77)
                                    }
                                }
                                _ = v77
                                _ = v78
                            }
                            var v79 int16 = 1
                            _ = v79
                        }
                    }
                    v80 := GenS4{F0: (("a" + ("a" + "\n\t")) + (("" + ((("a" + "a") + "") + "hello")) + "foo bar")), F1: -196, F2: (v2 | 1)}
                    fmt.Println(fi58, fi59)
                    v1 = v1
                    v19 -= v19
                    _ = v80
                } else {
                    var v81 uint = (v2 ^ (v2 + genFn0(255, ("hello" + "foo bar"))))
                    var v82 uint32 = (v3 << ((v2) & 63))
                    fmt.Println(v2, v1)
                    _ = v81
                    _ = v82
                }
                v83 := v5[("" + "\n\t")]
                if (v1 || ((fi59 | genFn1(0, 18446744073709551615)) <= (fi58 & fi58))) {
                    var v84 int32 = 1
                    _ = len(v5)
                    var v85 uint32 = (v3 & 65535)
                    if v83 {
                        v3 += (v85 >> ((genFn0(255, ("" + "hello"))) & 63))
                        fmt.Println(v85, v83)
                        if v1 {
                            v86 := GenS2{F0: -716, F1: v2}
                            switch (^(fi59)) {
                            case 32767:
                                fmt.Println(v3)
                            case 4294967295:
                                v19 -= (v19 | v19)
                            case 0:
                                var v87 int = (fi58 >> ((v2) & 63))
                                _ = v87
                            default:
                                v1 = (1 > 255)
                            }
                            if (("a" >= ("\n\t" + "\n\t")) && v83) {
                                v5[(("\n\t" + "hello") + "")] = v83
                                if (42 >= 926) {
                                    var v88 int16 = -1
                                    v19 = 4294967295
                                    v88 -= (v88 - (v88 ^ (-(v88))))
                                    _ = v88
                                }
                            }
                            _ = v86
                        }
                    }
                    _ = v84
                    _ = v85
                } else {
                    v83 = v83
                    var v89 float64 = (1.5 * 0.1)
                    v3 += v3
                    _ = v89
                }
                _ = fi59
                _ = v83
            }
            fmt.Println(v2, v1, fi58)
        } else {
            var v90 float32 = -2.25
            _ = v90
        }
        _ = fi58
    }
    if true {
        for fi91 := 0; fi91 < 9; fi91++ {
            fmt.Println(v3)
            for fi92 := 0; fi92 < 3; fi92++ {
                var v93 int8 = 0
                v19 = v19
                var v94 string = "\n\t"
                _ = fi92
                _ = v93
                _ = v94
            }
            _ = fi91
        }
        v95 := GenS4{F0: "", F1: 42, F2: (v2 ^ v2)}
        _ = v95
    }
    var v96 int64 = -9223372036854775808
    for fi97 := 0; fi97 < 4; fi97++ {
        v98 := GenS4{F0: "a", F1: 0, F2: v2}
        v2 += (v2 << ((v2) & 63))
        for fi99 := 0; fi99 < 12; fi99++ {
            v96 += v96
            _ = fi99
        }
        _ = fi97
        _ = v98
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v57
    _ = v96
}
