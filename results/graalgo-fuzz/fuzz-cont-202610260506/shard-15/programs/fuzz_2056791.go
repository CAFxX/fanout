package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16; F1 int8; F2 uint }
type GenS10 struct { F0 uint32; F1 uint64; F2 uint64 }
func genFn0(p0 uint64) int64 {
    var v6 uint16 = 1
    var v7 uint = 255
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 string, p1 int32) uint16 {
    fmt.Println(p0, p1)
    v8 := GenS1{F0: (((p0 + ("hello" + "\n\t")) + p0) + p0)}
    p0 = "a"
    var v9 int64 = -410
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]bool)
    var v4 string = "a"
    v5 := make(map[string]string)
    if true {
        var v10 uint16 = 353
        switch v4 {
        case "foo bar":
            var v11 uint8 = 2
            _ = v11
        default:
            v12 := []int{32768, -9223372036854775808, -693}
            _ = v12
        }
        if false {
            var v13 bool = (!false)
            _ = v13
        } else {
            v4 = v4
            var v14 uint = 4294967295
            v15 := GenS1{F0: v4}
            _ = v14
            _ = v15
        }
        _ = v10
    } else {
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 string = v4
            switch fi16 {
            case -1:
                var v18 int = (fi16 - fi16)
                _ = v18
            }
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 int64 = (^(v2))
                v4 = v4
                _ = fi19
                _ = v20
            }
            _ = fi16
            _ = v17
        }
        fmt.Println(v4, v4)
        for fi21 := 0; fi21 < 14; fi21++ {
            v22 := []int{fi21, (fi21 / ((fi21) | 1)), fi21}
            if ((v4 + "a") == v4) {
                var v23 uint32 = 2147483647
                if (true && (0 == 100)) {
                    v23 = v23
                    v1["hello"] = v23
                }
                _ = v23
            } else {
                for fi24 := 0; fi24 < 19; fi24++ {
                    var v25 float32 = (-2.25 - 0.1)
                    if (2147483647 >= 0) {
                        v26 := GenS2{F0: (v2 - 9223372036854775807)}
                        var v27 uint16 = 2
                        v27 = 255
                        _ = v26
                        _ = v27
                    }
                    if true {
                        v2 = genFn0(42)
                        _ = len(v5)
                    }
                    _ = fi24
                    _ = v25
                }
            }
            _ = len(v22)
            _ = fi21
            _ = v22
        }
    }
    v28 := GenS1{F0: "\n\t"}
    if (915 < 65535) {
        var v29 bool = false
        v29 = (65535 > 42)
        _ = v29
    }
    if (32767 > 32768) {
        v4 = "\n\t"
        if (0 < -950) {
            fmt.Println(v4, v4)
        } else {
            var v30 float64 = 1e10
            _ = v30
        }
        if ((!true) || (42 > 100)) {
            v4 = (("foo bar" + v4) + v4)
            v31 := GenS2{F0: (v2 * genFn0(2))}
            _ = v31
        } else {
            for fi32 := 0; fi32 < 17; fi32++ {
                v1[v4] = 65535
                v4 = v4
                _ = fi32
            }
            if (true || true) {
                fmt.Println(v2, v2)
                for fi33 := 0; fi33 < 19; fi33++ {
                    v4 = v4
                    _ = fi33
                }
            }
            for fi34 := 0; fi34 < 18; fi34++ {
                fmt.Println(v2)
                if (false && (18446744073709551615 > 2)) {
                    v35 := v1[v4]
                    delete(v3, "")
                    _ = v35
                } else {
                    if (false || (v2 < v2)) {
                        for fi36 := 0; fi36 < 10; fi36++ {
                            var v37 float64 = 0.1
                            delete(v1, v4)
                            fmt.Println(v4, v37)
                            v38 := v1["\n\t"]
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    } else {
                        for fi39 := 0; fi39 < 12; fi39++ {
                            v40 := GenS3{F0: fi39, F1: 635}
                            _ = fi39
                            _ = v40
                        }
                        v2 += (v2 - int64(fi34))
                    }
                }
                if ((-405 == 1) && false) {
                    if true {
                        if (65535 == 65535) {
                            switch fi34 {
                            case -231:
                                switch fi34 {
                                case 1:
                                    _ = len(v3)
                                }
                            default:
                                for fi41 := 0; fi41 < 12; fi41++ {
                                    fmt.Println(fi41, v4, fi41)
                                    var v42 uint32 = 1
                                    _ = fi41
                                    _ = v42
                                }
                            }
                            v2 = 9223372036854775807
                        } else {
                            v43 := []int{(-(fi34)), (fi34 << ((4294967295) & 63))}
                            var v44 bool = (true || (42 != 42))
                            var v45 int32 = 295
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                        fmt.Println(v4, v4)
                        v46 := []int{(^(fi34)), fi34}
                        _ = v46
                    }
                    if true {
                        v47 := GenS1{F0: v4}
                        switch "" {
                        case "hello":
                            v48 := []int{(fi34 - 0)}
                            _ = v48
                        case "\n\t":
                            v2 = genFn0(1)
                        }
                        v2 = genFn0(42)
                        _ = v47
                    }
                } else {
                    if false {
                        var v49 uint64 = 1
                        var v50 int32 = int32(v2)
                        v50 -= v50
                        _ = v49
                        _ = v50
                    } else {
                        for fi51 := 0; fi51 < 1; fi51++ {
                            var v52 uint32 = 2147483648
                            if true {
                                if false {
                                    v53 := GenS4{F0: true}
                                    _ = v53
                                } else {
                                    if (fi34 <= fi34) {
                                        switch 4294967295 {
                                        case 910:
                                            v52 = v52
                                        case -1:
                                            v54 := GenS4{F0: (fi34 > fi34)}
                                            if true {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if (127 == 42) {
                                                        var v55 int16 = 32767
                                                        var v56 uint64 = 42
                                                        switch v4 {
                                                        case "hello":
                                                            if false {
                                                                v4 = "foo bar"
                                                                switch (v4 + v4) {
                                                                case "foo bar":
                                                                    _ = len(v1)
                                                                case "a":
                                                                    v57 := GenS2{F0: genFn0(v56)}
                                                                    _ = v57
                                                                }
                                                            } else {
                                                                var v58 int16 = int16(v52)
                                                                _ = v58
                                                            }
                                                        default:
                                                            fmt.Println(v56, v2)
                                                        }
                                                        var v59 int32 = 2147483647
                                                        _ = v55
                                                        _ = v56
                                                        _ = v59
                                                    }
                                                case "":
                                                    if true {
                                                        v60 := GenS5{F0: -482, F1: v52}
                                                        fmt.Println(fi51)
                                                        if true {
                                                            v61 := []int{-2147483648}
                                                            var v62 string = v4
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                        _ = v60
                                                    }
                                                case "\n\t":
                                                    if false {
                                                        _ = len(v3)
                                                    } else {
                                                        if false {
                                                            fmt.Println(v52)
                                                            v4 = "hello"
                                                        } else {
                                                            var v63 int32 = 0
                                                            _ = v63
                                                        }
                                                        var v64 float64 = 0.1
                                                        _ = v64
                                                    }
                                                    var v65 int16 = -1
                                                    fmt.Println(fi34)
                                                    _ = v65
                                                default:
                                                    v3[v4] = true
                                                    v66 := GenS6{F0: true, F1: 16, F2: 0}
                                                    v52 = v52
                                                    _ = v66
                                                }
                                                if (false && (0.0 < (3.14159 / 1e10))) {
                                                    if (v4 < v4) {
                                                        if (-944 <= 1) {
                                                            var v67 float64 = 1.5
                                                            switch 127 {
                                                            case -9223372036854775808:
                                                                if (18446744073709551615 < 809) {
                                                                    switch fi51 {
                                                                    case -94:
                                                                        v68 := GenS7{F0: -9223372036854775808, F1: 993}
                                                                        _ = v68
                                                                    }
                                                                }
                                                            case 127:
                                                                fmt.Println(v4)
                                                            case 621:
                                                                if false {
                                                                    v52 = v52
                                                                    var v69 uint16 = 32767
                                                                    _ = v69
                                                                } else {
                                                                    v70 := GenS8{F0: 897}
                                                                    v4 = ((v4 + ("\n\t" + ((v4 + v4) + v4))) + v4)
                                                                    _ = v70
                                                                }
                                                                var v71 float64 = v67
                                                                var v72 bool = true
                                                                _ = v71
                                                                _ = v72
                                                            default:
                                                                v73 := GenS2{F0: v2}
                                                                _ = v73
                                                            }
                                                            _ = v67
                                                        } else {
                                                            fmt.Println(fi34)
                                                        }
                                                        fmt.Println(fi51, v52, v2)
                                                        var v74 int32 = 1
                                                        _ = v74
                                                    }
                                                }
                                            }
                                            fmt.Println(v2)
                                            v75 := GenS7{F0: (fi34 / (((fi34 << ((4294967295) & 63))) | 1)), F1: 212}
                                            _ = v54
                                            _ = v75
                                        default:
                                            var v76 int = (^(fi51))
                                            var v77 int64 = 0
                                            if (!(false || (v4 != (("\n\t" + "a") + ("a" + v4))))) {
                                                v78 := []int{(^(fi51)), (fi34 | (fi34 & fi34)), fi51, fi34}
                                                v52 = 42
                                                v79 := v3[v4]
                                                _ = v78
                                                _ = v79
                                            }
                                            var v80 uint = 576
                                            delete(v3, v4)
                                            if true {
                                                var v81 int = v76
                                                v82 := GenS5{F0: 82, F1: v52}
                                                _ = v81
                                                _ = v82
                                            } else {
                                                if (54 != -1) {
                                                    v83 := GenS6{F0: ("a" >= "foo bar"), F1: 127, F2: 18446744073709551615}
                                                    _ = v83
                                                } else {
                                                    fmt.Println(v77)
                                                }
                                                if (283 < 1) {
                                                    fmt.Println(v80, v4)
                                                    if (!(-1 <= -91)) {
                                                        var v84 uint16 = 705
                                                        v52 -= 2
                                                        _ = v84
                                                    }
                                                    fmt.Println(fi51, v52)
                                                    v85 := []int{(fi51 << (((v80 / (((v80 & v80)) | 1))) & 63)), (fi51 * (fi51 & fi51))}
                                                    _ = v85
                                                } else {
                                                    if true {
                                                        v86 := []int{}
                                                        _ = v86
                                                    } else {
                                                        _ = len(v3)
                                                        var v87 int16 = -1
                                                        var v88 int = 32768
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                    fmt.Println(v52, v80)
                                                    var v89 int32 = -2147483648
                                                    _ = v89
                                                }
                                                v90 := GenS3{F0: (fi34 / ((fi34) | 1)), F1: 2}
                                                _ = v90
                                            }
                                            _ = v76
                                            _ = v77
                                            _ = v80
                                        }
                                        if (!false) {
                                            switch (fi34 - fi51) {
                                            case -1:
                                                if true {
                                                    var v91 int16 = -178
                                                    _ = v91
                                                }
                                            case 1:
                                                v92 := GenS5{F0: 0, F1: (v52 - (v52 - 34))}
                                                _ = v92
                                            case -202:
                                                var v93 int32 = 42
                                                _ = v93
                                            }
                                        }
                                        v4 = "foo bar"
                                    }
                                    var v94 float32 = 1.5
                                    _ = v94
                                }
                                _ = len(v1)
                                v95 := GenS9{F0: 1000, F1: -128, F2: 4294967295}
                                v52 = (v52 & 4294967295)
                                v96 := []int{fi51, fi34, (fi34 >> ((0) & 63)), (fi51 * fi34)}
                                _ = v95
                                _ = v96
                            }
                            fmt.Println(fi34, fi51, v4)
                            _ = fi51
                            _ = v52
                        }
                    }
                }
                _ = fi34
            }
        }
    } else {
        for fi97 := 0; fi97 < 8; fi97++ {
            fmt.Println(v4, v4)
            v2 = v2
            _ = fi97
        }
        switch v4 {
        case "foo bar":
            fmt.Println(v4, v2)
        case "":
            var v98 uint16 = 2
            _ = v98
        case "hello":
            for fi99 := 0; fi99 < 12; fi99++ {
                for fi100 := 0; fi100 < 9; fi100++ {
                    v101 := GenS7{F0: (-(fi99)), F1: 323}
                    _ = fi100
                    _ = v101
                }
                _ = fi99
            }
        default:
            v4 = v4
        }
    }
    v102 := GenS3{F0: 32767, F1: 377}
    var v103 bool = ((!true) || (0 == 1))
    v2 = (v2 * (v2 / ((genFn0(4294967295)) | 1)))
    v104 := []int{}
    for fi105 := 0; fi105 < 9; fi105++ {
        v5[v4] = v4
        var v106 int32 = -2147483648
        _ = fi105
        _ = v106
    }
    v104 = append(v104, -1)
    var v107 uint64 = 264
    var v108 uint64 = (v107 % ((v107) | 1))
    var v109 uint64 = v108
    if v103 {
        if v103 {
            v110 := GenS10{F0: 841, F1: 679, F2: (v107 / ((986) | 1))}
            _ = v110
        }
    } else {
        for fi111 := 0; fi111 < 1; fi111++ {
            if (0 == -9) {
                v4 = "a"
            }
            var v112 uint32 = 2147483647
            _ = fi111
            _ = v112
        }
    }
    fmt.Println(v2, v107)
    switch 9223372036854775807 {
    case 2147483647:
        var v113 uint = 4294967295
        _ = v113
    case 0:
        if len(v104) > 0 {
            v104[0] = 744
        }
    default:
        fmt.Println(v2, v4, v109)
    }
    var v114 uint16 = 661
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("slicelen v104:", len(v104), v104)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v107
    _ = v108
    _ = v109
    _ = v114
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
