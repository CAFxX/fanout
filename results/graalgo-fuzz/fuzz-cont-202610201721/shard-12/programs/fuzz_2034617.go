package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint; F1 string; F2 uint }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int16; F1 int16; F2 int }
type GenS6 struct { F0 int16 }
func genFn0() uint8 {
    var v6 uint64 = 42
    if true {
        if false {
            v6 -= (^(v6))
            var v7 uint64 = 4294967295
            var v8 uint32 = 4294967295
            _ = v7
            _ = v8
        } else {
            var v9 uint8 = 1
            var v10 uint = 349
            v11 := []int{626}
            _ = v10
            _ = v11
            _ = v9
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            switch -9223372036854775808 {
            case 842:
                v13 := GenS1{F0: "hello", F1: 42, F2: 255}
                _ = v13
            }
            if (1 > -1) {
                v14 := GenS1{F0: ("" + ("a" + ("" + ("" + "hello")))), F1: 42, F2: 42}
                v6 += v6
                var v15 int32 = 14
                _ = v14
                _ = v15
            } else {
                v16 := GenS1{F0: "foo bar", F1: 1, F2: 2}
                var v17 uint8 = 227
                var v18 int64 = 42
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi12
        }
        for fi19 := 0; fi19 < 17; fi19++ {
            v6 += (v6 << ((512) & 63))
            _ = fi19
        }
    } else {
        v20 := GenS1{F0: "\n\t", F1: -1, F2: 42}
        _ = v20
    }
    _ = v6
    return 128
}
func genFn1(p0 uint8, p1 int64) uint32 {
    p1 = p1
    var v21 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = v21
    return 34
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 124
    var v3 uint8 = 128
    v4 := make(map[int]uint32)
    var v5 uint32 = 220
    v22 := []int{127}
    var v23 uint64 = 0
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 bool = ((!false) || false)
        if (-128 >= -128) {
            switch "hello" {
            case "":
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 uint64 = v23
                    fmt.Println(fi24, v5, v2)
                    _ = fi26
                    _ = v27
                }
            }
        } else {
            v3 -= v3
            v22 = append(v22, int(v2))
        }
        if (v25 || (745 > 885)) {
            for fi28 := 0; fi28 < 7; fi28++ {
                if v25 {
                    if v25 {
                        v1[(fi24 >> ((1) & 63))] = (^(fi28))
                    }
                    switch (fi24 ^ (fi28 / ((42) | 1))) {
                    case 128:
                        v3 = 1
                    case 2147483647:
                        v29 := v1[int(v3)]
                        _ = v29
                    default:
                        if v25 {
                            v30 := GenS1{F0: "a", F1: -2147483648, F2: v3}
                            _ = v30
                        }
                    }
                    fmt.Println(v3)
                } else {
                    fmt.Println(v3, v2, v2)
                    if false {
                        var v31 bool = (((0.0 - (1e10 + 0.0)) - 1.5) != float64(0.1))
                        switch 2147483647 {
                        case 0:
                            v25 = (((42 <= 757) || false) || (true && ((1e10 == 1.5) && ((975 >= 2) || false))))
                        }
                        v31 = v31
                        _ = v31
                    } else {
                        _ = len(v22)
                        fmt.Println(v3, v2)
                        if (127 == 100) {
                            if (!true) {
                                if (false || false) {
                                    v32 := v4[-9223372036854775808]
                                    var v33 uint8 = 192
                                    v34 := GenS2{F0: 42, F1: v5}
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                } else {
                                    switch "a" {
                                    case "\n\t":
                                        var v35 uint32 = (^(v5))
                                        _ = v35
                                    case "foo bar":
                                        fmt.Println(v2, fi24)
                                    }
                                }
                                v5 = v5
                            }
                            v25 = ((true && true) || ((42 <= 42) && v25))
                        } else {
                            switch -9223372036854775808 {
                            case 571:
                                var v36 bool = (0 >= -20)
                                _ = v36
                            default:
                                v5 = 0
                            }
                            v23 -= v23
                        }
                        v4[fi28] = (v5 + v5)
                    }
                    var v37 int64 = -824
                    _ = v37
                }
                var v38 float32 = (0.1 / (((1e10 / -2.25) / 1.0) / float32(1.5)))
                if len(v22) > 0 {
                    v22[0] = (fi24 << ((v2) & 63))
                }
                var v39 int16 = 32767
                v40 := GenS1{F0: "foo bar", F1: 592, F2: (v3 ^ 0)}
                if v25 {
                    fmt.Println(v23, v2, v5)
                    v41 := GenS1{F0: (("" + "") + "\n\t"), F1: -1, F2: v3}
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    _ = v41
                } else {
                    if v25 {
                        v42 := GenS2{F0: -128, F1: genFn1(v3, 15)}
                        v39 += v39
                        var v43 int64 = -9223372036854775808
                        _ = v42
                        _ = v43
                    }
                    if v25 {
                        v44 := GenS1{F0: ("a" + "\n\t"), F1: -738, F2: (v3 + (v3 >> ((635) & 63)))}
                        _ = len(v1)
                        v45 := GenS3{F0: (v2 >> ((v2) & 63)), F1: (("" + ("foo bar" + ("foo bar" + ("\n\t" + "")))) + "\n\t"), F2: 255}
                        _ = v44
                        _ = v45
                    } else {
                        var v46 bool = false
                        switch -874 {
                        case 32768:
                            if true {
                                var v47 uint8 = v3
                                v5 = v5
                                _ = v47
                            } else {
                                var v48 int8 = 42
                                _ = len(v4)
                                v48 = -1
                                if (uint32(v2) == v5) {
                                    var v49 uint16 = 634
                                    v50 := v1[fi28]
                                    switch -2147483648 {
                                    case -9223372036854775808:
                                        switch ((("foo bar" + "") + "foo bar") + ("hello" + ("hello" + "\n\t"))) {
                                        case "\n\t":
                                            switch (^(fi28)) {
                                            case 861:
                                                if len(v22) >= 2 {
                                                    v22 = v22[1:len(v22)]
                                                }
                                            default:
                                                var v51 bool = true
                                                _ = v51
                                            }
                                        case "foo bar":
                                            v52 := GenS2{F0: -128, F1: (v5 | v5)}
                                            _ = v52
                                        }
                                    case 9223372036854775807:
                                        v53 := GenS3{F0: (v2 << ((v2) & 63)), F1: "", F2: v2}
                                        _ = v53
                                    case 127:
                                        if v46 {
                                            if (int64(v39) == 1) {
                                                v22 = append(v22, v50)
                                                switch ("foo bar" + ("a" + "")) {
                                                case "\n\t":
                                                    if false {
                                                        _ = len(v4)
                                                        var v54 int = (v50 - int(v3))
                                                        _ = v54
                                                    } else {
                                                        var v55 int32 = 42
                                                        _ = v55
                                                    }
                                                }
                                            }
                                            var v56 uint16 = v49
                                            var v57 float32 = v38
                                            fmt.Println(v23, v25)
                                            _ = v56
                                            _ = v57
                                        } else {
                                            fmt.Println(v49, v49)
                                        }
                                    default:
                                        if v46 {
                                            if len(v22) > 0 {
                                                v22[0] = (fi24 ^ -9223372036854775808)
                                            }
                                        } else {
                                            fmt.Println(v46)
                                            v58 := GenS2{F0: 0, F1: genFn1(genFn0(), -9223372036854775808)}
                                            v59 := GenS3{F0: (v2 & (^(v2))), F1: "hello", F2: v2}
                                            _ = v58
                                            _ = v59
                                        }
                                    }
                                    var v60 int64 = 1
                                    var v61 float32 = ((v38 - 1e10) - v38)
                                    _ = v49
                                    _ = v50
                                    _ = v60
                                    _ = v61
                                } else {
                                    if v25 {
                                        v62 := GenS4{F0: (^(v2)), F1: 319}
                                        _ = v62
                                    } else {
                                        v63 := GenS1{F0: ("a" + ""), F1: 1, F2: v3}
                                        _ = v63
                                    }
                                }
                                _ = v48
                            }
                        case 9223372036854775807:
                            delete(v1, (^(fi24)))
                        case 555:
                            if (v2 >= uint(v23)) {
                                v64 := GenS3{F0: (v2 >> (((v2 | uint(v23))) & 63)), F1: "\n\t", F2: v2}
                                _ = v64
                            } else {
                                v65 := GenS2{F0: -15, F1: genFn1(v3, int64(v3))}
                                v23 -= (v23 / (((v23 * v23)) | 1))
                                v66 := GenS4{F0: (v2 * uint(v5)), F1: v2}
                                _ = v65
                                _ = v66
                            }
                        default:
                            var v67 int16 = (v39 % (((v39 & (v39 ^ (v39 * (v39 & 1))))) | 1))
                            _ = v67
                        }
                        if (!v25) {
                            var v68 string = ""
                            switch (v68 + v68) {
                            case "":
                                switch v68 {
                                case "":
                                    var v69 int64 = 0
                                    _ = v69
                                }
                            default:
                                v38 = v38
                            }
                            v70 := GenS5{F0: v39, F1: v39, F2: (fi28 - (fi24 * fi28))}
                            _ = v68
                            _ = v70
                        } else {
                            _ = len(v22)
                            if v46 {
                                v38 = v38
                                fmt.Println(v25)
                            }
                        }
                        _ = len(v22)
                        v2 += uint(v3)
                        _ = v46
                    }
                    fmt.Println(v3, fi24)
                }
                _ = len(v22)
                if v25 {
                    var v71 uint = (v2 % (((v2 + (^(v2)))) | 1))
                    fmt.Println(v39)
                    v39 = (v39 + (^(v39)))
                    _ = v71
                } else {
                    var v72 int64 = -1
                    if false {
                        var v73 string = "a"
                        switch fi28 {
                        case 1:
                            var v74 float32 = v38
                            _ = v74
                        case 2147483647:
                            if v25 {
                                fmt.Println(v38, v5)
                                var v75 uint64 = (v23 % ((v23) | 1))
                                _ = v75
                            }
                        case 32768:
                            if true {
                                v76 := GenS3{F0: v2, F1: "", F2: v2}
                                v3 = (v3 + v3)
                                _ = v76
                            } else {
                                switch (v73 + ("\n\t" + v73)) {
                                case "\n\t":
                                    _ = len(v22)
                                default:
                                    var v77 int64 = 293
                                    _ = v77
                                }
                                switch fi28 {
                                case -128:
                                    v39 = int16(v23)
                                case 2147483647:
                                    v72 = 1
                                }
                                fmt.Println(v38, fi24, v73)
                                switch fi28 {
                                case -634:
                                    v5 -= v5
                                case 1:
                                    var v78 float64 = ((-2.25 * (0.0 + (0.1 / 1.0))) / (1.5 / 1.5))
                                    var v79 uint16 = 691
                                    _ = v78
                                    _ = v79
                                case -9223372036854775808:
                                    var v80 uint = (v2 | 0)
                                    _ = v80
                                default:
                                    switch (fi28 % (((fi24 >> ((v2) & 63))) | 1)) {
                                    case 0:
                                        fmt.Println(fi24, fi28)
                                    }
                                }
                            }
                        default:
                            var v81 int16 = v39
                            _ = v81
                        }
                        v82 := GenS5{F0: (v39 >> ((v2) & 63)), F1: 0, F2: (fi28 | 1)}
                        _ = v73
                        _ = v82
                    } else {
                        fmt.Println(v72, fi24)
                    }
                    v23 -= 922
                    _ = v72
                }
                _ = fi28
                _ = v38
                _ = v39
                _ = v40
            }
            for fi83 := 0; fi83 < 5; fi83++ {
                switch 127 {
                case 4294967295:
                    var v84 uint32 = v5
                    _ = v84
                case 317:
                    switch "a" {
                    case "":
                        switch fi24 {
                        case 956:
                            if len(v22) > 0 {
                                v22[0] = (fi24 & (fi83 + fi24))
                            }
                        case 127:
                            fmt.Println(fi24)
                        default:
                            v23 += 42
                        }
                    case "\n\t":
                        v5 -= v5
                    default:
                        var v85 bool = true
                        _ = v85
                    }
                }
                _ = fi83
            }
            v86 := GenS6{F0: -977}
            _ = v86
        } else {
            if ("\n\t" < "hello") {
                switch fi24 {
                case 234:
                    fmt.Println(v5, v3)
                case -1:
                    for fi87 := 0; fi87 < 19; fi87++ {
                        v23 = 2
                        v88 := GenS6{F0: 1}
                        v2 = (v2 >> ((uint(v23)) & 63))
                        _ = fi87
                        _ = v88
                    }
                case 304:
                    var v89 string = ""
                    _ = v89
                }
                var v90 int = (-(fi24))
                _ = v90
            } else {
                for fi91 := 0; fi91 < 20; fi91++ {
                    v92 := v1[(fi24 >> ((65535) & 63))]
                    switch fi91 {
                    case 0:
                        var v93 uint16 = 32767
                        _ = v93
                    }
                    _ = fi91
                    _ = v92
                }
            }
            v94 := v4[(fi24 / (((-(fi24))) | 1))]
            _ = v94
        }
        _ = fi24
        _ = v25
    }
    var v95 int64 = 104
    for fi96 := 0; fi96 < 15; fi96++ {
        var v97 float32 = float32(float64((0.1 / ((1.5 + ((1e10 - (1.5 + 3.14159)) - 3.14159)) + 3.14159))))
        var v98 uint16 = 65535
        _ = fi96
        _ = v97
        _ = v98
    }
    v99 := GenS2{F0: 0, F1: v5}
    v2 = (v2 / ((v2) | 1))
    for fi100 := 0; fi100 < 2; fi100++ {
        if true {
            var v101 int = (fi100 ^ (fi100 & (fi100 ^ (fi100 ^ -1))))
            _ = v101
        }
        v102 := v4[(fi100 >> (((v2 + v2)) & 63))]
        _ = fi100
        _ = v102
    }
    for fi103 := 0; fi103 < 8; fi103++ {
        v3 -= v3
        switch fi103 {
        case -1:
            v5 = (v5 ^ (v5 ^ (v5 & (v5 / (((^(v5))) | 1)))))
        }
        _ = fi103
    }
    v23 -= (^(v23))
    fmt.Println(v23)
    v3 = v3
    var v104 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v104
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v95
    _ = v99
}
