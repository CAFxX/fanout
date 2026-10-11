package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint8; F1 bool; F2 string }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 int16; F1 uint8; F2 string }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS11 struct { F0 uint16 }
func genFn0(p0 int, p1 int8, p2 int32) uint8 {
    p0 += p0
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 = p0
        p1 = -128
        _ = fi7
    }
    switch ((("" + "foo bar") + ("foo bar" + "\n\t")) + "a") {
    case "a":
        for fi8 := 0; fi8 < 4; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                v10 := []int{0, fi8, 895}
                switch "foo bar" {
                case "foo bar":
                    var v11 float32 = 1.5
                    _ = v11
                default:
                    if true {
                        if (p0 < (p0 - (fi9 >> ((2) & 63)))) {
                            p2 -= (-(p2))
                            p2 += p2
                            switch int(p2) {
                            case 1:
                                var v12 int8 = 125
                                _ = v12
                            case 32768:
                                v10 = append(v10, p0)
                            }
                        } else {
                            v13 := GenS1{F0: 4294967295, F1: 255, F2: 255}
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            v10 = append(v10, p0)
                            _ = v13
                        }
                    }
                }
                p1 = (p1 ^ (p1 << ((24) & 63)))
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    case "hello":
        if false {
            if true {
                var v14 uint16 = 820
                v14 = 582
                _ = v14
            } else {
                p2 += (p2 | p2)
                var v15 string = (((("a" + "\n\t") + "\n\t") + "") + "\n\t")
                _ = v15
            }
            p1 += (^(p1))
        }
    case "\n\t":
        var v16 float64 = ((-2.25 - -2.25) - float64(1.5))
        _ = v16
    default:
        var v17 float64 = ((1e10 - 0.1) * 0.0)
        _ = v17
    }
    v18 := []int{}
    v19 := GenS1{F0: (p0 - 1), F1: 255, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 169
}

func main() {
    var v1 int8 = -27
    var v2 bool = (1 != 2)
    var v3 int64 = -41
    var v4 int64 = -1
    var v5 string = (("a" + "\n\t") + "")
    v6 := make(map[string]string)
    if (127 >= 742) {
        v20 := GenS2{F0: 65535, F1: 4294967295, F2: (-2147483648 <= -601)}
        v5 = v5
        _ = v20
    } else {
        v2 = v2
    }
    var v21 string = "\n\t"
    if (722 < 32767) {
        if v2 {
            fmt.Println(v2, v5, v5)
        } else {
            var v22 uint16 = 771
            switch -521 {
            case 4294967295:
                v21 = "foo bar"
            case 32768:
                switch 0 {
                case 127:
                    v1 += (v1 - v1)
                default:
                    var v23 int8 = v1
                    _ = v23
                }
            case -445:
                for fi24 := 0; fi24 < 14; fi24++ {
                    if v2 {
                        for fi25 := 0; fi25 < 7; fi25++ {
                            if (false && v2) {
                                v26 := GenS1{F0: (fi24 / ((int(v1)) | 1)), F1: 467, F2: 1}
                                _ = v26
                            } else {
                                if v2 {
                                    if v2 {
                                        var v27 uint32 = 138
                                        _ = v27
                                    } else {
                                        v28 := GenS2{F0: 2147483647, F1: 255, F2: (v22 < (v22 >> ((42) & 63)))}
                                        _ = v28
                                    }
                                    switch fi25 {
                                    case 1:
                                        switch fi24 {
                                        case 0:
                                            v29 := v6[v5]
                                            _ = v29
                                        case 128:
                                            v21 = "a"
                                        case 32767:
                                            v22 = v22
                                        }
                                    default:
                                        v22 -= 32767
                                    }
                                    v30 := []int{fi24}
                                    _ = v30
                                }
                                var v31 int32 = 0
                                _ = v31
                            }
                            var v32 int8 = (v1 ^ v1)
                            _ = fi25
                            _ = v32
                        }
                        v6[v5] = v5
                    } else {
                        fmt.Println(v21, v5, v4)
                    }
                    if (989 <= 32767) {
                        var v33 uint32 = 406
                        if ((v21 + "foo bar") == (v21 + v5)) {
                            v2 = false
                            fmt.Println(v3, v4, v22)
                        }
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 16; fi34++ {
                        switch fi24 {
                        case -866:
                            if false {
                                var v35 int64 = (v4 | (-(v3)))
                                switch ("hello" + ((("\n\t" + (((("hello" + (v5 + v21)) + v21) + v21) + v5)) + (((("hello" + v5) + v5) + "\n\t") + v21)) + (("foo bar" + v5) + (((v21 + "a") + ("a" + v5)) + (v5 + ""))))) {
                                case "a":
                                    if false {
                                        var v36 uint32 = 2147483647
                                        if v2 {
                                            fmt.Println(v4, v5)
                                            v37 := GenS1{F0: (fi34 << ((18446744073709551615) & 63)), F1: v22, F2: 127}
                                            if v2 {
                                                v38 := GenS3{F0: v36}
                                                _ = v38
                                            } else {
                                                var v39 int64 = 9223372036854775807
                                                _ = v39
                                            }
                                            var v40 uint8 = 255
                                            v41 := GenS4{F0: (v36 / ((v36) | 1)), F1: v36, F2: uint16(v40)}
                                            if v2 {
                                                v5 = v5
                                            } else {
                                                v42 := v6[("a" + v5)]
                                                v2 = v2
                                                v43 := []int{fi34, (fi24 * (fi24 + fi34)), fi34}
                                                _ = v42
                                                _ = v43
                                            }
                                            _ = v37
                                            _ = v40
                                            _ = v41
                                        } else {
                                            if (v2 && (!((v2 || false) || v2))) {
                                                switch 405 {
                                                case 128:
                                                    if (((^(v4)) == 4) && v2) {
                                                        var v44 int32 = 942
                                                        v21 = "foo bar"
                                                        var v45 uint32 = 42
                                                        if (!false) {
                                                            v6[(v21 + "hello")] = v5
                                                            if (127 >= 2) {
                                                                v6["\n\t"] = v21
                                                                _ = len(v6)
                                                                v46 := GenS5{F0: 0, F1: v2, F2: v5}
                                                                _ = v46
                                                            }
                                                            v22 = v22
                                                            if (42 > 0) {
                                                                switch v5 {
                                                                case "hello":
                                                                    if v2 {
                                                                        v47 := []int{9223372036854775807, -2147483648, fi24, int(v44)}
                                                                        v4 -= v4
                                                                        _ = v47
                                                                    }
                                                                }
                                                                v48 := GenS6{F0: (v36 + (v36 - v45)), F1: v45}
                                                                v44 += v44
                                                                _ = v48
                                                            } else {
                                                                var v49 int32 = (-(v44))
                                                                v50 := GenS1{F0: (fi24 << ((0) & 63)), F1: (v22 | (v22 - (v22 >> ((18446744073709551615) & 63)))), F2: genFn0(fi24, v1, (v49 - int32(fi34)))}
                                                                if ((!v2) && (true || (!(!v2)))) {
                                                                    v51 := []int{(fi24 % ((fi34) | 1)), fi24, fi34}
                                                                    var v52 uint16 = (v22 - 728)
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                                if v2 {
                                                                    var v53 bool = (((-(v35)) >= -1) && v2)
                                                                    if (v2 && false) {
                                                                        v54 := GenS6{F0: 2, F1: v36}
                                                                        v55 := []int{(fi34 * (fi24 / (((fi24 & int(v22))) | 1))), fi24, (-(fi34))}
                                                                        v56 := GenS2{F0: v36, F1: 2, F2: v53}
                                                                        _ = v54
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                    _ = v53
                                                                } else {
                                                                    v57 := []int{int(v22), 9223372036854775807, 32767}
                                                                    var v58 string = v5
                                                                    v59 := GenS7{F0: 1, F1: 42, F2: v5}
                                                                    var v60 int16 = 360
                                                                    _ = v57
                                                                    _ = v58
                                                                    _ = v59
                                                                    _ = v60
                                                                }
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                        } else {
                                                            var v61 string = v5
                                                            v62 := GenS8{F0: fi24, F1: (-(fi24))}
                                                            v63 := []int{fi34}
                                                            _ = v61
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        fmt.Println(v1, v4)
                                                        v64 := GenS6{F0: (v36 - (^(v36))), F1: 0}
                                                        _ = v64
                                                    }
                                                case -2147483648:
                                                    fmt.Println(v4)
                                                default:
                                                    _ = len(v6)
                                                }
                                                var v65 int8 = 1
                                                var v66 int16 = -359
                                                _ = v65
                                                _ = v66
                                            }
                                            var v67 uint16 = (v22 * v22)
                                            _ = v67
                                        }
                                        var v68 bool = (3.14159 >= 0.1)
                                        _ = v36
                                        _ = v68
                                    }
                                case "hello":
                                    switch (fi34 / ((628) | 1)) {
                                    case 32767:
                                        v69 := GenS9{F0: 127}
                                        _ = v69
                                    case -2147483648:
                                        var v70 float64 = 0.1
                                        _ = v70
                                    case 4294967295:
                                        v4 -= -1
                                    }
                                }
                                _ = v35
                            }
                        case -2147483648:
                            var v71 float64 = ((1.5 + 0.0) - (1e10 / 1.0))
                            _ = v71
                        default:
                            var v72 uint = 4294967295
                            _ = v72
                        }
                        v73 := GenS9{F0: fi24}
                        _ = fi34
                        _ = v73
                    }
                    _ = fi24
                }
            }
            _ = v22
        }
    } else {
        var v74 string = ((("\n\t" + "hello") + "hello") + "")
        _ = v74
    }
    v6[((v5 + v5) + ("" + "hello"))] = "a"
    if false {
        v75 := GenS4{F0: 965, F1: 828, F2: 42}
        fmt.Println(v21)
        v5 = ("\n\t" + (v5 + ("\n\t" + v5)))
        _ = v75
    }
    v2 = false
    v21 = "foo bar"
    var v76 bool = v2
    for fi77 := 0; fi77 < 8; fi77++ {
        switch v5 {
        case "\n\t":
            v78 := []int{(fi77 % ((189) | 1)), fi77, fi77, (fi77 << ((728) & 63))}
            _ = v78
        default:
            for fi79 := 0; fi79 < 11; fi79++ {
                var v80 int64 = (^(v4))
                var v81 int = fi79
                switch 32767 {
                case -9223372036854775808:
                    var v82 float64 = 1e10
                    _ = v82
                case -1:
                    if (!(6 != 65535)) {
                        fmt.Println(fi77, v21)
                        if (1 != 255) {
                            v83 := GenS10{F0: 846, F1: 2147483648, F2: 1}
                            var v84 float32 = (0.1 + (-2.25 * 0.0))
                            v85 := GenS9{F0: (-(fi77))}
                            _ = v83
                            _ = v84
                            _ = v85
                        } else {
                            fmt.Println(v80, v81, v4)
                            if (v2 && true) {
                                delete(v6, v21)
                                fmt.Println(v5, v4)
                            } else {
                                delete(v6, v5)
                            }
                        }
                        v2 = (650 < 1)
                    } else {
                        var v86 string = v21
                        _ = v86
                    }
                }
                _ = fi79
                _ = v80
                _ = v81
            }
        }
        _ = fi77
    }
    v87 := GenS11{F0: 1}
    if true {
        switch 32768 {
        case 245:
            v88 := []int{127, -282, 785}
            _ = v88
        case -150:
            var v89 uint = 460
            _ = v89
        default:
            for fi90 := 0; fi90 < 13; fi90++ {
                fmt.Println(v2)
                switch v5 {
                case "hello":
                    switch (fi90 | 632) {
                    case 32768:
                        v76 = ((fi90 % ((4294967295) | 1)) < (fi90 >> ((1) & 63)))
                    case -2147483648:
                        for fi91 := 0; fi91 < 5; fi91++ {
                            v5 = ("\n\t" + v21)
                            var v92 int8 = (v1 * v1)
                            if v2 {
                                v5 = ""
                                v93 := v6["hello"]
                                _ = v93
                            }
                            _ = fi91
                            _ = v92
                        }
                    case 1:
                        delete(v6, "")
                    default:
                        for fi94 := 0; fi94 < 7; fi94++ {
                            var v95 int64 = v3
                            var v96 float64 = 0.0
                            var v97 int8 = 44
                            fmt.Println(v2)
                            _ = fi94
                            _ = v95
                            _ = v96
                            _ = v97
                        }
                    }
                default:
                    var v98 float64 = (-2.25 * 3.14159)
                    _ = v98
                }
                _ = fi90
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v87
}
