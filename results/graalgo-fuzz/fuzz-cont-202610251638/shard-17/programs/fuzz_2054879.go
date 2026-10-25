package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint }
type GenS5 struct { F0 int32; F1 uint16; F2 int16 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 string; F1 bool; F2 int32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 string; F1 int8 }
type GenS12 struct { F0 uint32 }
func genFn0() int16 {
    if (1 >= 4294967295) {
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := []int{-707, -990}
            fmt.Println(fi6)
            _ = fi6
            _ = v7
        }
        var v8 bool = (272 < 1)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 19; fi9++ {
            if (-2.25 < (1.5 * 1e10)) {
                fmt.Println(fi9)
                var v10 uint = 42
                v11 := []int{(fi9 / ((fi9) | 1)), fi9, -9223372036854775808}
                _ = v10
                _ = v11
            } else {
                v12 := GenS1{F0: -1, F1: 2147483648, F2: 255}
                var v13 uint = 18446744073709551615
                _ = v12
                _ = v13
            }
            _ = fi9
        }
    }
    var v14 int64 = 42
    fmt.Println(v14)
    var v15 int64 = v14
    v15 = v15
    _ = v14
    _ = v15
    return -32768
}

func main() {
    v1 := []uint32{490, 346, 490}
    var v2 uint8 = 0
    var v3 bool = true
    v4 := []int{128, 32768, -9223372036854775808}
    v5 := make(map[uint32]string)
    v3 = v3
    switch -1 {
    case 127:
        v16 := GenS1{F0: 776, F1: 2, F2: 255}
        _ = v16
    case -492:
        for fi17 := 0; fi17 < 17; fi17++ {
            v2 += v2
            _ = fi17
        }
    case -9223372036854775808:
        v18 := v5[4294967295]
        _ = v18
    }
    _ = len(v5)
    if (128 != (^(v2))) {
        v19 := GenS2{F0: 0}
        var v20 int16 = 508
        _ = v19
        _ = v20
    }
    switch "hello" {
    case "hello":
        v2 -= v2
    case "":
        if len(v1) > 0 {
            v1[0] = 597
        }
    default:
        if len(v4) > 0 {
            v4[0] = 381
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = v2
    v21 := GenS3{F0: "a", F1: 65535}
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v22 float64 = float64((3.14159 / 3.14159))
        _ = v22
    case 127:
        var v23 float64 = 0.1
        _ = v23
    default:
        if v3 {
            v3 = v3
            fmt.Println(v2, v3)
        } else {
            fmt.Println(v3, v2)
            for fi24 := 0; fi24 < 10; fi24++ {
                if (v3 && v3) {
                    if (false || v3) {
                        v5[65535] = ("\n\t" + "a")
                    } else {
                        v25 := GenS3{F0: ("\n\t" + "foo bar"), F1: 32767}
                        _ = v25
                    }
                } else {
                    v26 := GenS3{F0: "a", F1: 1}
                    _ = v26
                }
                var v27 int8 = -100
                _ = fi24
                _ = v27
            }
        }
    }
    v3 = (!(-32768 > -32768))
    v28 := GenS2{F0: 1}
    v29 := GenS3{F0: "hello", F1: 65535}
    var v30 uint = 42
    if v3 {
        switch "hello" {
        case "\n\t":
            switch 4294967295 {
            case 872:
                fmt.Println(v30, v2, v2)
            case 127:
                if v3 {
                    if (1 != 4294967295) {
                        var v31 uint8 = 42
                        v31 += (v2 - (v31 + (v31 / ((v2) | 1))))
                        _ = v31
                    }
                    if (!v3) {
                        v32 := GenS4{F0: 3, F1: -1, F2: v30}
                        _ = v32
                    } else {
                        v4 = append(v4, -9223372036854775808)
                    }
                }
            case -9223372036854775808:
                var v33 float32 = (1.5 * 3.14159)
                _ = v33
            default:
                for fi34 := 0; fi34 < 16; fi34++ {
                    var v35 uint64 = 786
                    _ = fi34
                    _ = v35
                }
            }
        case "a":
            delete(v5, 4294967295)
        case "hello":
            v36 := GenS2{F0: -70}
            _ = v36
        }
    } else {
        for fi37 := 0; fi37 < 19; fi37++ {
            switch "hello" {
            case "foo bar":
                switch "" {
                case "foo bar":
                    fmt.Println(v30)
                default:
                    for fi38 := 0; fi38 < 18; fi38++ {
                        v1 = append(v1, 42)
                        v2 = (v2 - (v2 * (^(v2))))
                        _ = fi38
                    }
                }
            case "a":
                for fi39 := 0; fi39 < 6; fi39++ {
                    var v40 float32 = (3.14159 * 0.0)
                    v40 -= (v40 / 1e10)
                    v30 += (v30 + v30)
                    _ = fi39
                    _ = v40
                }
            case "\n\t":
                var v41 float32 = 1.5
                _ = v41
            }
            var v42 float64 = 1e10
            _ = fi37
            _ = v42
        }
        var v43 uint16 = 1
        fmt.Println(v3)
        _ = v43
    }
    if ((v3 || v3) && (2 < uint64(v2))) {
        if v3 {
            for fi44 := 0; fi44 < 14; fi44++ {
                v45 := GenS1{F0: -487, F1: 1, F2: v30}
                v2 -= 1
                _ = fi44
                _ = v45
            }
            for fi46 := 0; fi46 < 1; fi46++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi47 := 0; fi47 < 3; fi47++ {
                    if false {
                        var v48 int64 = -1
                        v3 = false
                        v5[67] = ("\n\t" + "\n\t")
                        _ = v48
                    }
                    var v49 uint16 = 42
                    if (-9223372036854775808 <= 1) {
                        switch fi47 {
                        case 32768:
                            v50 := GenS4{F0: v49, F1: 700, F2: (^(v30))}
                            _ = v50
                        }
                        if (9223372036854775807 == -9223372036854775808) {
                            if v3 {
                                var v51 bool = (!v3)
                                _ = v51
                            } else {
                                if (v3 && v3) {
                                    v5[4294967295] = ("a" + "")
                                    v30 -= (v30 >> ((v30) & 63))
                                    var v52 string = "\n\t"
                                    var v53 int16 = 1000
                                    _ = v52
                                    _ = v53
                                }
                                if (!(((v2 << ((v30) & 63)) <= (v2 * (v2 % ((114) | 1)))) && v3)) {
                                    var v54 bool = v3
                                    _ = v54
                                } else {
                                    v2 = v2
                                    var v55 float64 = ((3.14159 + ((3.14159 / 3.14159) / 1.0)) + ((1e10 / 1.0) + (3.14159 * 1e10)))
                                    if true {
                                        var v56 int16 = 42
                                        v55 += 0.1
                                        _ = v56
                                    }
                                    switch "hello" {
                                    case "hello":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "foo bar":
                                        if v3 {
                                            var v57 uint8 = v2
                                            var v58 int16 = 32767
                                            _ = v57
                                            _ = v58
                                        } else {
                                            v3 = v3
                                            v59 := GenS4{F0: (v49 >> ((761) & 63)), F1: -2147483648, F2: v30}
                                            _ = v59
                                        }
                                    case "\n\t":
                                        fmt.Println(fi46, fi47)
                                    }
                                    delete(v5, 1)
                                    var v60 int32 = -949
                                    if true {
                                        fmt.Println(v30)
                                    } else {
                                        v55 -= ((float64((-2.25 + 0.0)) + float64(((-2.25 * 1.5) / 1e10))) - (v55 * 1.5))
                                        switch (^(fi47)) {
                                        case -9223372036854775808:
                                            var v61 string = (("\n\t" + ("foo bar" + "hello")) + ("" + ""))
                                            _ = v61
                                        case 2147483647:
                                            _ = len(v5)
                                        }
                                        if (18446744073709551615 >= (v30 ^ (v30 >> ((v30) & 63)))) {
                                            var v62 int16 = genFn0()
                                            if (42 < 2) {
                                                v30 = (v30 ^ v30)
                                                v60 += (v60 & (^(v60)))
                                                v63 := GenS3{F0: "", F1: (v49 | (v49 | (v49 & v49)))}
                                                _ = v63
                                            }
                                            var v64 uint32 = 2
                                            _ = v62
                                            _ = v64
                                        } else {
                                            if true {
                                                v30 += (^(v30))
                                                switch "foo bar" {
                                                case "":
                                                    v65 := GenS5{F0: int32(fi47), F1: (v49 + 0), F2: 217}
                                                    _ = v65
                                                case "foo bar":
                                                    if (false && ("\n\t" >= ("hello" + "a"))) {
                                                        if ((false || (42 >= 601)) || (-32768 == 1)) {
                                                            v66 := GenS1{F0: -373, F1: 2147483647, F2: (v30 ^ (v30 - (v30 % ((v30) | 1))))}
                                                            delete(v5, 42)
                                                            if (4294967295 < 42) {
                                                                var v67 uint16 = v49
                                                                var v68 int8 = 100
                                                                fmt.Println(v49, v2, v67)
                                                                _ = v67
                                                                _ = v68
                                                            } else {
                                                                v55 -= 0.1
                                                                var v69 uint32 = 929
                                                                fmt.Println(v69, v55, v69)
                                                                fmt.Println(v3, v49)
                                                                _ = v69
                                                            }
                                                            _ = v66
                                                        } else {
                                                            if v3 {
                                                                var v70 uint = (v30 % ((v30) | 1))
                                                                var v71 float64 = ((v55 / 1e10) + v55)
                                                                _ = v70
                                                                _ = v71
                                                            } else {
                                                                v3 = ((v30 & v30) == v30)
                                                                v72 := GenS2{F0: -128}
                                                                if (uint16(v30) < 65535) {
                                                                    v73 := GenS6{F0: 1, F1: -754}
                                                                    _ = v73
                                                                } else {
                                                                    if v3 {
                                                                        v74 := GenS4{F0: (v49 * (^(v49))), F1: 2147483647, F2: 0}
                                                                        _ = v74
                                                                    } else {
                                                                        if v3 {
                                                                            v2 = (v2 & (v2 - v2))
                                                                            _ = len(v5)
                                                                            fmt.Println(v49, v49)
                                                                        } else {
                                                                            if false {
                                                                                fmt.Println(v55, v60)
                                                                                if (((v49 % ((346) | 1)) <= (v49 ^ 0)) || false) {
                                                                                    switch ("a" + "a") {
                                                                                    case "hello":
                                                                                        switch fi46 {
                                                                                        case 9223372036854775807:
                                                                                            var v75 uint64 = 2
                                                                                            _ = v75
                                                                                        case -591:
                                                                                            v3 = ((^(fi47)) > int(v30))
                                                                                        case -9223372036854775808:
                                                                                            v76 := GenS7{F0: 42, F1: (v2 >> (((v30 | v30)) & 63))}
                                                                                            _ = v76
                                                                                        default:
                                                                                            v55 = v55
                                                                                        }
                                                                                        v55 -= (v55 * v55)
                                                                                        var v77 uint8 = (v2 * v2)
                                                                                        _ = v77
                                                                                    default:
                                                                                        if (0.0 == 0.1) {
                                                                                            v60 += (v60 - v60)
                                                                                        }
                                                                                        v78 := GenS2{F0: -36}
                                                                                        var v79 int8 = 127
                                                                                        _ = v78
                                                                                        _ = v79
                                                                                    }
                                                                                } else {
                                                                                    switch fi46 {
                                                                                    case 9223372036854775807:
                                                                                        fmt.Println(v2, v49)
                                                                                    case 0:
                                                                                        v55 = (1e10 + float64(0.0))
                                                                                    }
                                                                                }
                                                                                var v80 uint16 = uint16(v2)
                                                                                _ = v80
                                                                            }
                                                                            var v81 float64 = 0.1
                                                                            _ = v81
                                                                        }
                                                                    }
                                                                    var v82 bool = (!(42 <= 675))
                                                                    v55 = (-2.25 - v55)
                                                                    var v83 uint16 = 2
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                                _ = v72
                                                            }
                                                        }
                                                    } else {
                                                        v30 -= v30
                                                        v84 := GenS8{F0: -9223372036854775808, F1: v2}
                                                        if (203 == -32768) {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v85 int16 = 1000
                                                            _ = v85
                                                        } else {
                                                            _ = len(v4)
                                                            fmt.Println(v49, v55)
                                                            var v86 int32 = v60
                                                            _ = v86
                                                        }
                                                        _ = v84
                                                    }
                                                default:
                                                    v2 += 203
                                                }
                                            } else {
                                                fmt.Println(fi46, v30, v30)
                                            }
                                            if (false || false) {
                                                _ = len(v1)
                                                if (4294967295 >= 2) {
                                                    v30 -= (v30 * (v30 & 2))
                                                    var v87 float32 = 3.14159
                                                    var v88 string = "hello"
                                                    _ = v87
                                                    _ = v88
                                                } else {
                                                    fmt.Println(v30, fi46)
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                        }
                                    }
                                    if (127 == 42) {
                                        v89 := GenS9{F0: "", F1: ((v2 - (v2 * (v2 + (v2 - (v2 * (^(v2))))))) == uint8(v30)), F2: v60}
                                        v90 := GenS2{F0: 127}
                                        _ = v89
                                        _ = v90
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 2147483648
                                    }
                                    v91 := GenS10{F0: (v60 % ((-2147483648) | 1))}
                                    var v92 int64 = 320
                                    _ = v55
                                    _ = v60
                                    _ = v91
                                    _ = v92
                                }
                            }
                            fmt.Println(v49, v2, v3)
                        } else {
                            var v93 uint8 = uint8(v49)
                            v94 := GenS11{F0: (("hello" + "") + ("" + ("" + ("foo bar" + ("a" + "\n\t"))))), F1: 100}
                            _ = v93
                            _ = v94
                        }
                        fmt.Println(fi46, v30)
                        var v95 float32 = ((float32((float64(0.0) / 1.0)) * 0.0) * -2.25)
                        var v96 int = (fi47 << (((^(v30))) & 63))
                        _ = v95
                        _ = v96
                    } else {
                        var v97 int32 = 0
                        switch 32768 {
                        case 9223372036854775807:
                            v98 := v5[uint32(v49)]
                            _ = v98
                        case 51:
                            fmt.Println(v97)
                        }
                        var v99 bool = (1 < 32767)
                        v100 := GenS12{F0: 121}
                        _ = v100
                        _ = v97
                        _ = v99
                    }
                    _ = fi47
                    _ = v49
                }
                _ = fi46
            }
        }
        if v3 {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v30 -= 1
            delete(v5, 4294967295)
        } else {
            v2 = (^(v2))
            var v101 uint32 = 2
            _ = v101
        }
    } else {
        delete(v5, 146)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
