package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 int32; F2 bool }
type GenS9 struct { F0 bool; F1 int8 }
type GenS10 struct { F0 bool; F1 uint }
type GenS11 struct { F0 uint; F1 int64; F2 uint64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint; F1 uint }
type GenS14 struct { F0 uint64; F1 int16; F2 bool }
type GenS15 struct { F0 bool; F1 int32; F2 int16 }
type GenS16 struct { F0 int16; F1 uint64; F2 int8 }
type GenS17 struct { F0 int8; F1 int8 }
func genFn0(p0 int, p1 bool) uint {
    if p1 {
        v4 := GenS1{F0: 1, F1: 2147483647, F2: 892}
        _ = v4
    } else {
        p0 += 128
        fmt.Println(p0, p1)
    }
    if true {
        var v5 int64 = 1
        _ = v5
    } else {
        p0 = 1
        var v6 bool = (((100 != -128) && ((9223372036854775807 != 9223372036854775807) && (!p1))) || (false && true))
        var v7 float32 = (1.5 / float32(1e10))
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    return 596
}

func main() {
    var v1 int32 = 42
    v2 := []string{}
    var v3 uint32 = 69
    var v8 string = ""
    switch -2147483648 {
    case 26:
        v1 += 659
    case 102:
        v1 -= v1
    case -1:
        if (v1 < v1) {
            if (!true) {
                for fi9 := 0; fi9 < 10; fi9++ {
                    fmt.Println(v1, v3)
                    v10 := GenS1{F0: 42, F1: v1, F2: v3}
                    for fi11 := 0; fi11 < 9; fi11++ {
                        if true {
                            var v12 float64 = -2.25
                            v12 = 1.5
                            var v13 uint = 49
                            _ = v12
                            _ = v13
                        } else {
                            if ((100 == -1) || false) {
                                if (fi9 >= (fi9 - fi11)) {
                                    var v14 uint8 = uint8(v3)
                                    _ = v14
                                } else {
                                    var v15 int8 = -6
                                    switch 128 {
                                    case 267:
                                        v16 := GenS1{F0: genFn0(fi9, false), F1: (-(v1)), F2: v3}
                                        _ = v16
                                    case 2147483647:
                                        if true {
                                            switch 1 {
                                            case 9223372036854775807:
                                                var v17 float32 = 0.1
                                                _ = v17
                                            }
                                            fmt.Println(v15)
                                            if true {
                                                _ = len(v2)
                                                var v18 int32 = (v1 & (v1 >> ((65535) & 63)))
                                                switch v8 {
                                                case "foo bar":
                                                    v2 = append(v2, v8)
                                                case "hello":
                                                    if len(v2) > 0 {
                                                        v2[0] = (v8 + v8)
                                                    }
                                                default:
                                                    v19 := GenS2{F0: -32768, F1: 128}
                                                    _ = v19
                                                }
                                                if true {
                                                    v2 = append(v2, "\n\t")
                                                    if len(v2) > 0 {
                                                        v2[0] = ("\n\t" + "foo bar")
                                                    }
                                                    var v20 float64 = (1.5 - 0.1)
                                                    _ = v20
                                                } else {
                                                    if true {
                                                        switch v8 {
                                                        case "":
                                                            switch -1 {
                                                            case 286:
                                                                if false {
                                                                    v8 = ((((((("" + v8) + "hello") + "foo bar") + (v8 + "hello")) + v8) + "\n\t") + v8)
                                                                    if (!(128 != 128)) {
                                                                        var v21 int64 = 0
                                                                        v22 := GenS3{F0: v3, F1: -696}
                                                                        v23 := GenS4{F0: 1}
                                                                        _ = v21
                                                                        _ = v22
                                                                        _ = v23
                                                                    }
                                                                    v8 = v8
                                                                }
                                                            case -2147483648:
                                                                var v24 uint = 4294967295
                                                                _ = v24
                                                            default:
                                                                v25 := GenS3{F0: v3, F1: (fi9 * fi9)}
                                                                _ = v25
                                                            }
                                                        case "\n\t":
                                                            v3 = (v3 % ((42) | 1))
                                                        }
                                                    }
                                                    v15 += (-(v15))
                                                    if (false || true) {
                                                        if (128 <= 215) {
                                                            if (false && false) {
                                                                if true {
                                                                    if (!(2 >= 2)) {
                                                                        v1 = -2147483648
                                                                        var v26 int32 = (v1 << ((697) & 63))
                                                                        _ = v26
                                                                    }
                                                                    if len(v2) > 0 {
                                                                        v2[0] = ("hello" + (v8 + v8))
                                                                    }
                                                                } else {
                                                                    fmt.Println(v15)
                                                                    var v27 float32 = ((0.1 * ((0.1 / 1.5) / 3.14159)) / (0.0 - 1e10))
                                                                    var v28 float64 = 1.5
                                                                    _ = v27
                                                                    _ = v28
                                                                }
                                                                switch v8 {
                                                                case "foo bar":
                                                                    if false {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = "foo bar"
                                                                        }
                                                                        if true {
                                                                            var v29 uint = 82
                                                                            _ = v29
                                                                        }
                                                                    } else {
                                                                        switch v8 {
                                                                        case "":
                                                                            v8 = v8
                                                                        case "\n\t":
                                                                            v30 := GenS4{F0: 237}
                                                                            _ = v30
                                                                        default:
                                                                            var v31 float64 = -2.25
                                                                            _ = v31
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    v32 := GenS1{F0: 1, F1: 2147483647, F2: 1}
                                                                    _ = v32
                                                                default:
                                                                    var v33 uint32 = (v3 | (v3 % (((v3 % ((2) | 1))) | 1)))
                                                                    _ = v33
                                                                }
                                                            } else {
                                                                if (2 != 2) {
                                                                    if (!((1 == 0) && false)) {
                                                                        _ = len(v2)
                                                                    }
                                                                    switch (v8 + v8) {
                                                                    case "\n\t":
                                                                        var v34 bool = true
                                                                        _ = v34
                                                                    case "":
                                                                        v15 = v15
                                                                    }
                                                                    v1 -= (v1 | v1)
                                                                    if false {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        var v35 uint64 = 18446744073709551615
                                                                        _ = v35
                                                                    }
                                                                }
                                                            }
                                                            v18 -= v18
                                                            if true {
                                                                switch 0 {
                                                                case -1:
                                                                    var v36 int16 = int16(fi11)
                                                                    _ = v36
                                                                }
                                                                var v37 int16 = 945
                                                                _ = len(v2)
                                                                _ = v37
                                                            } else {
                                                                var v38 string = "foo bar"
                                                                if true {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        switch ("\n\t" + ((("a" + v8) + "\n\t") + (v38 + v38))) {
                                                                        case "hello":
                                                                            v18 = v18
                                                                        case "a":
                                                                            var v39 int8 = (-(v15))
                                                                            _ = v39
                                                                        default:
                                                                            var v40 bool = ((fi11 ^ (fi11 + fi11)) >= fi9)
                                                                            _ = v40
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi11, v1)
                                                                    }
                                                                } else {
                                                                    var v41 int8 = v15
                                                                    if false {
                                                                        var v42 float32 = (0.0 + 0.1)
                                                                        _ = v42
                                                                    } else {
                                                                        _ = len(v2)
                                                                        if (4294967295 < 1) {
                                                                            v38 = "a"
                                                                        }
                                                                        v8 = (((v8 + "hello") + ((v38 + ("hello" + "a")) + v38)) + "hello")
                                                                        var v43 float32 = (1e10 * 3.14159)
                                                                        _ = v43
                                                                    }
                                                                    _ = v41
                                                                }
                                                                _ = v38
                                                            }
                                                            if true {
                                                                var v44 bool = ((0 != 65535) || true)
                                                                v45 := GenS5{F0: 2, F1: 127, F2: (-(v1))}
                                                                _ = v44
                                                                _ = v45
                                                            } else {
                                                                v8 = ("a" + (v8 + ("hello" + "\n\t")))
                                                                switch (v8 + v8) {
                                                                case "":
                                                                    v15 = (v15 * (v15 >> ((828) & 63)))
                                                                case "\n\t":
                                                                    v1 -= 2147483647
                                                                    var v46 float64 = 0.1
                                                                    _ = v46
                                                                default:
                                                                    if true {
                                                                        var v47 float64 = 3.14159
                                                                        v48 := GenS6{F0: 65535, F1: 185}
                                                                        var v49 float32 = (float32(v47) * -2.25)
                                                                        switch fi11 {
                                                                        case 0:
                                                                            v15 = (^(v15))
                                                                        case -991:
                                                                            v50 := GenS1{F0: 65535, F1: v1, F2: v3}
                                                                            _ = v50
                                                                        case -2147483648:
                                                                            if (("hello" + v8) != v8) {
                                                                                if (654 != 2) {
                                                                                    v51 := GenS7{F0: 812}
                                                                                    fmt.Println(v47, v15)
                                                                                    switch v8 {
                                                                                    case "foo bar":
                                                                                        switch fi11 {
                                                                                        case 42:
                                                                                            var v52 int8 = (v15 + v15)
                                                                                            switch v8 {
                                                                                            case "\n\t":
                                                                                                v3 = (^(v3))
                                                                                            case "a":
                                                                                                v3 = 2147483647
                                                                                            default:
                                                                                                fmt.Println(v47)
                                                                                            }
                                                                                            v53 := GenS2{F0: 32767, F1: 2}
                                                                                            v3 = 4294967295
                                                                                            var v54 int8 = v52
                                                                                            _ = v52
                                                                                            _ = v53
                                                                                            _ = v54
                                                                                        default:
                                                                                            v18 -= (v18 << ((genFn0(32767, (false || (true && (!(18446744073709551615 >= 0)))))) & 63))
                                                                                        }
                                                                                    case "":
                                                                                        if false {
                                                                                            fmt.Println(fi11, v15, v18)
                                                                                        } else {
                                                                                            var v55 bool = (80 >= 71)
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            _ = v55
                                                                                        }
                                                                                    default:
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = (v8 + (v8 + "hello"))
                                                                                        }
                                                                                    }
                                                                                    v15 = (v15 & (v15 * v15))
                                                                                    _ = v51
                                                                                }
                                                                            } else {
                                                                                v2 = append(v2, v8)
                                                                            }
                                                                        default:
                                                                            v2 = append(v2, (((v8 + v8) + (v8 + ("\n\t" + ((v8 + "hello") + ((v8 + ("hello" + v8)) + v8))))) + v8))
                                                                        }
                                                                        var v56 uint64 = 4294967295
                                                                        _ = v47
                                                                        _ = v48
                                                                        _ = v49
                                                                        _ = v56
                                                                    }
                                                                }
                                                                v15 += (-(v15))
                                                            }
                                                            v8 = ((v8 + v8) + ("foo bar" + v8))
                                                            v57 := GenS8{F0: "\n\t", F1: v1, F2: false}
                                                            _ = v57
                                                        } else {
                                                            if len(v2) > 0 {
                                                                v2[0] = (v8 + "")
                                                            }
                                                            if (0.0 > 3.14159) {
                                                                var v58 uint64 = 1
                                                                if len(v2) > 0 {
                                                                    v2[0] = v8
                                                                }
                                                                var v59 uint32 = (v3 + v3)
                                                                v60 := GenS3{F0: v3, F1: fi9}
                                                                _ = v58
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                        }
                                                    } else {
                                                        var v61 int8 = -128
                                                        if len(v2) > 0 {
                                                            v2[0] = v8
                                                        }
                                                        v62 := GenS9{F0: ((-(v1)) <= v1), F1: (v61 - v61)}
                                                        var v63 float64 = -2.25
                                                        _ = v61
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                }
                                                fmt.Println(v18, fi11)
                                                fmt.Println(v18)
                                                _ = v18
                                            } else {
                                                v8 = v8
                                                var v64 float64 = 0.0
                                                if (true && true) {
                                                    if true {
                                                        var v65 uint64 = 1
                                                        _ = v65
                                                    } else {
                                                        v64 += (0.0 - (v64 / v64))
                                                    }
                                                    fmt.Println(v15, fi11, v15)
                                                }
                                                _ = v64
                                            }
                                        } else {
                                            if (true || (1 != 0)) {
                                                v66 := GenS10{F0: true, F1: 1}
                                                _ = v66
                                            }
                                        }
                                    case 432:
                                        v3 = 4294967295
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    _ = v15
                                }
                                var v67 uint64 = uint64(v3)
                                _ = v67
                            } else {
                                v68 := GenS11{F0: 255, F1: 0, F2: 42}
                                fmt.Println(v3, fi9)
                                _ = v68
                            }
                            var v69 uint8 = 255
                            _ = v69
                        }
                        v8 = v8
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v70 int32 = (v1 & v1)
                            _ = v70
                        } else {
                            v3 += 2
                            v8 = (v8 + v8)
                            v3 = 0
                        }
                        _ = fi11
                    }
                    _ = fi9
                    _ = v10
                }
                for fi71 := 0; fi71 < 15; fi71++ {
                    v72 := GenS12{F0: genFn0((^(fi71)), (((true || (-533 <= fi71)) || (false || (true || true))) || (!("" == "foo bar"))))}
                    v3 += v3
                    _ = fi71
                    _ = v72
                }
                var v73 uint64 = 42
                _ = v73
            } else {
                switch v8 {
                case "a":
                    var v74 int = 4294967295
                    _ = v74
                default:
                    for fi75 := 0; fi75 < 9; fi75++ {
                        var v76 int32 = v1
                        v77 := GenS13{F0: 352, F1: 42}
                        _ = fi75
                        _ = v76
                        _ = v77
                    }
                }
                v78 := GenS14{F0: 0, F1: 1000, F2: (float32((1.5 + 0.0)) == (0.0 - (float32(3.14159) * (1.5 / 0.1))))}
                for fi79 := 0; fi79 < 6; fi79++ {
                    v1 = (v1 + (v1 - v1))
                    v80 := GenS15{F0: false, F1: (v1 % ((v1) | 1)), F2: 0}
                    _ = fi79
                    _ = v80
                }
                _ = v78
            }
            var v81 uint = 727
            _ = v81
        }
    }
    switch "hello" {
    case "hello":
        var v82 float32 = -2.25
        _ = v82
    case "a":
        if true {
            var v83 string = "foo bar"
            v1 -= 792
            _ = v83
        }
    }
    var v84 uint = 255
    v85 := GenS16{F0: -898, F1: 4294967295, F2: 100}
    v84 = (v84 * genFn0(310, false))
    fmt.Println(v84, v84, v1)
    var v86 string = "a"
    v86 = v8
    var v87 uint = 252
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v88 int64 = -1
    var v89 float32 = (0.0 / 1.0)
    v90 := GenS15{F0: false, F1: v1, F2: int16(v1)}
    var v91 bool = true
    switch 597 {
    case 128:
        if true {
            v92 := GenS17{F0: 112, F1: -1}
            var v93 float64 = -2.25
            if false {
                var v94 float64 = ((3.14159 + (v93 + (0.0 * v93))) / v93)
                _ = v94
            }
            _ = v92
            _ = v93
        }
    case 127:
        v84 = v87
    default:
        for fi95 := 0; fi95 < 1; fi95++ {
            v96 := GenS15{F0: v91, F1: (^(v1)), F2: -32768}
            for fi97 := 0; fi97 < 19; fi97++ {
                v86 = ("a" + ((v86 + ("" + ("hello" + v86))) + (("foo bar" + v8) + "hello")))
                var v98 uint32 = 0
                fmt.Println(v8, v91, fi97)
                _ = fi97
                _ = v98
            }
            var v99 float32 = v89
            _ = fi95
            _ = v96
            _ = v99
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v84:", v84)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v84
    _ = v85
    _ = v86
    _ = v87
    _ = v88
    _ = v89
    _ = v90
    _ = v91
}
