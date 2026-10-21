package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS4 struct { F0 int64; F1 int32; F2 int8 }
type GenS5 struct { F0 string; F1 int16; F2 bool }
type GenS6 struct { F0 int32; F1 int32; F2 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint; F1 int32 }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 int16; F1 uint }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int8; F1 uint8 }
type GenS13 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS14 struct { F0 uint8; F1 uint8 }
type GenS15 struct { F0 uint; F1 string }
type GenS16 struct { F0 int8 }
func genFn0(p0 int) int16 {
    p0 -= 32767
    fmt.Println(p0)
    if true {
        fmt.Println(p0)
        for fi4 := 0; fi4 < 5; fi4++ {
            p0 += 42
            for fi5 := 0; fi5 < 9; fi5++ {
                v6 := GenS1{F0: 226}
                p0 -= (fi4 | fi5)
                _ = fi5
                _ = v6
            }
            _ = fi4
        }
        v7 := GenS2{F0: 116}
        _ = v7
    } else {
        p0 -= 9223372036854775807
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS3{F0: 178, F1: -709, F2: 4294967295}
            _ = fi8
            _ = v9
        }
        p0 = p0
    }
    p0 = p0
    if (42 == 386) {
        var v10 uint16 = 0
        v10 = (v10 + 65535)
        _ = v10
    } else {
        fmt.Println(p0)
        for fi11 := 0; fi11 < 19; fi11++ {
            p0 += p0
            _ = fi11
        }
    }
    _ = p0
    return 216
}
func genFn1(p0 int64, p1 uint8) int64 {
    v12 := GenS4{F0: 9223372036854775807, F1: 42, F2: 1}
    var v13 uint8 = (p1 | (p1 - 42))
    if true {
        var v14 int8 = 69
        _ = v14
    } else {
        v15 := GenS2{F0: 42}
        if (false || ((p0 % ((-926) | 1)) == (p0 % ((p0) | 1)))) {
            v16 := GenS3{F0: 42, F1: -689, F2: 4294967295}
            fmt.Println(v13)
            _ = v16
        } else {
            if (2 <= (^(v13))) {
                p0 -= (-(p0))
                v13 = (p1 + 2)
            }
            var v17 int64 = (p0 << ((65535) & 63))
            _ = v17
        }
        _ = v15
    }
    v13 -= p1
    v18 := GenS5{F0: "hello", F1: genFn0(4294967295), F2: (-1 <= -69)}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v18
    return 9223372036854775807
}

func main() {
    var v1 float64 = (1e10 / 1e10)
    var v2 uint8 = 114
    var v3 uint64 = 364
    var v19 bool = (-2147483648 >= 9223372036854775807)
    fmt.Println(v1, v2, v1)
    if ((v2 % (((v2 / ((101) | 1))) | 1)) == v2) {
        if v19 {
            if v19 {
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 uint8 = v2
                    for fi22 := 0; fi22 < 8; fi22++ {
                        v19 = v19
                        fmt.Println(v21, v19, v1)
                        var v23 string = "hello"
                        _ = fi22
                        _ = v23
                    }
                    fmt.Println(v1)
                    _ = fi20
                    _ = v21
                }
                v24 := GenS1{F0: v2}
                for fi25 := 0; fi25 < 19; fi25++ {
                    var v26 bool = v19
                    v27 := []int{(fi25 * fi25), fi25, fi25, 553}
                    var v28 int32 = int32(v2)
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = v24
            }
            v2 = v2
            switch "\n\t" {
            case "foo bar":
                v29 := GenS6{F0: -862, F1: 1, F2: 128}
                _ = v29
            case "a":
                v30 := GenS7{F0: 1}
                _ = v30
            default:
                fmt.Println(v1, v3)
            }
        } else {
            if (9223372036854775807 <= genFn1(0, (^(v2)))) {
                if v19 {
                    v31 := GenS7{F0: 0}
                    _ = v31
                }
            } else {
                v32 := GenS8{F0: 882, F1: 0}
                v3 = 1
                switch "foo bar" {
                case "":
                    var v33 int16 = -146
                    _ = v33
                }
                _ = v32
            }
            for fi34 := 0; fi34 < 19; fi34++ {
                v1 = ((-2.25 + v1) - -2.25)
                v3 += (v3 + 2)
                if (290 != -2147483648) {
                    var v35 int8 = 100
                    var v36 bool = (247 < 9223372036854775807)
                    _ = v35
                    _ = v36
                } else {
                    if v19 {
                        v37 := []int{fi34, fi34, (-(fi34)), fi34}
                        _ = len(v37)
                        _ = v37
                    } else {
                        v3 = (v3 | (v3 | 240))
                    }
                    var v38 int32 = -2147483648
                    _ = v38
                }
                _ = fi34
            }
            v39 := GenS9{F0: 32767, F1: -599}
            _ = v39
        }
        if true {
            v2 = v2
        } else {
            for fi40 := 0; fi40 < 14; fi40++ {
                v41 := []int{fi40, (fi40 - int(v2))}
                v3 = (v3 << ((234) & 63))
                v3 += v3
                _ = fi40
                _ = v41
            }
            var v42 int = 9223372036854775807
            _ = v42
        }
        var v43 uint64 = 955
        _ = v43
    } else {
        for fi44 := 0; fi44 < 18; fi44++ {
            if false {
                v45 := GenS10{F0: genFn0((fi44 >> ((65535) & 63))), F1: 666}
                if (1.5 != 0.0) {
                    v3 += 42
                }
                for fi46 := 0; fi46 < 10; fi46++ {
                    v47 := []int{}
                    var v48 uint8 = v2
                    v49 := GenS7{F0: -62}
                    _ = fi46
                    _ = v47
                    _ = v48
                    _ = v49
                }
                _ = v45
            }
            _ = fi44
        }
        var v50 string = ("" + "a")
        _ = v50
    }
    v2 += v2
    v51 := []int{127, -2147483648}
    var v52 uint8 = uint8(v3)
    switch 9223372036854775807 {
    case -862:
        switch 9223372036854775807 {
        case 9223372036854775807:
            if (1 < 1) {
                v53 := GenS9{F0: 677, F1: -1}
                v54 := GenS4{F0: genFn1(-1, (^(v52))), F1: 1, F2: 10}
                _ = v53
                _ = v54
            } else {
                for fi55 := 0; fi55 < 6; fi55++ {
                    if len(v51) >= 2 {
                        v51 = v51[1:len(v51)]
                    }
                    var v56 uint8 = (v2 ^ 0)
                    v57 := GenS11{F0: 42}
                    _ = fi55
                    _ = v56
                    _ = v57
                }
                if (!v19) {
                    if len(v51) >= 2 {
                        v51 = v51[1:len(v51)]
                    }
                    if ((v1 / -2.25) == float64(((-2.25 * 3.14159) - 0.0))) {
                        if ((!true) || v19) {
                            for fi58 := 0; fi58 < 17; fi58++ {
                                fmt.Println(v52, v1)
                                v2 = v52
                                _ = fi58
                            }
                            switch "a" {
                            case "":
                                v59 := GenS12{F0: 110, F1: 98}
                                _ = v59
                            default:
                                var v60 int64 = 752
                                _ = v60
                            }
                            if v19 {
                                v1 += v1
                                if len(v51) >= 2 {
                                    v51 = v51[1:len(v51)]
                                }
                                var v61 uint32 = 619
                                _ = v61
                            } else {
                                v62 := GenS6{F0: -1, F1: 1, F2: 32767}
                                v63 := GenS11{F0: -1}
                                _ = v62
                                _ = v63
                            }
                        } else {
                            var v64 int = 128
                            if true {
                                if (453 != 255) {
                                    fmt.Println(v3, v64, v64)
                                    fmt.Println(v19, v64, v52)
                                } else {
                                    v64 = -400
                                    switch ("hello" + "a") {
                                    case "hello":
                                        for fi65 := 0; fi65 < 3; fi65++ {
                                            if (128 > (v64 ^ (fi65 ^ 2147483647))) {
                                                if false {
                                                    for fi66 := 0; fi66 < 8; fi66++ {
                                                        v64 = (fi65 - fi65)
                                                        var v67 float32 = ((3.14159 * float32(3.14159)) + float32((v1 + v1)))
                                                        v68 := GenS13{F0: 65535, F1: -128, F2: (v3 & uint64(v2))}
                                                        var v69 uint16 = 32767
                                                        if (v19 || v19) {
                                                            var v70 int64 = -9223372036854775808
                                                            _ = v70
                                                        }
                                                        _ = fi66
                                                        _ = v67
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    v52 = (v2 - (^(v52)))
                                                } else {
                                                    v51 = append(v51, (fi65 % (((-(v64))) | 1)))
                                                    if len(v51) >= 2 {
                                                        v51 = v51[1:len(v51)]
                                                    }
                                                    v64 -= fi65
                                                }
                                                v71 := GenS3{F0: v3, F1: -2147483648, F2: (v3 ^ (v3 | 1))}
                                                var v72 float32 = (-2.25 * (0.0 - 0.0))
                                                _ = v71
                                                _ = v72
                                            }
                                            for fi73 := 0; fi73 < 12; fi73++ {
                                                v1 = v1
                                                var v74 string = ((((("a" + "hello") + ("" + ("" + ("hello" + "foo bar")))) + "foo bar") + ("a" + "foo bar")) + "hello")
                                                _ = fi73
                                                _ = v74
                                            }
                                            if false {
                                                v75 := GenS7{F0: 100}
                                                v2 += v52
                                                switch "hello" {
                                                case "\n\t":
                                                    if len(v51) >= 2 {
                                                        v51 = v51[1:len(v51)]
                                                    }
                                                case "hello":
                                                    var v76 float32 = 0.0
                                                    _ = v76
                                                }
                                                _ = v75
                                            } else {
                                                if (255 > 419) {
                                                    switch "a" {
                                                    case "foo bar":
                                                        switch 1 {
                                                        case 9223372036854775807:
                                                            v1 = -2.25
                                                        case -529:
                                                            fmt.Println(v2, v64)
                                                        default:
                                                            fmt.Println(v3)
                                                        }
                                                    default:
                                                        for fi77 := 0; fi77 < 10; fi77++ {
                                                            v64 -= int(v52)
                                                            var v78 int8 = 100
                                                            if (((fi65 << ((2) & 63)) != (fi77 / ((127) | 1))) && v19) {
                                                                v51 = append(v51, 9223372036854775807)
                                                            } else {
                                                                v3 += 0
                                                                fmt.Println(v52, fi65)
                                                                if ((fi77 > fi77) || false) {
                                                                    var v79 int16 = 1
                                                                    v80 := GenS10{F0: 1, F1: 2}
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                            }
                                                            switch (fi77 - (v64 * fi77)) {
                                                            case -302:
                                                                switch v64 {
                                                                case 4294967295:
                                                                    v81 := GenS3{F0: (v3 % (((v3 >> ((2) & 63))) | 1)), F1: -2147483648, F2: (v3 * (v3 ^ (v3 * v3)))}
                                                                    _ = v81
                                                                case -9223372036854775808:
                                                                    v78 -= v78
                                                                default:
                                                                    v82 := GenS6{F0: 330, F1: 1, F2: 1}
                                                                    _ = v82
                                                                }
                                                            }
                                                            fmt.Println(fi77, v2)
                                                            var v83 int16 = 187
                                                            fmt.Println(fi77, v52, v3)
                                                            var v84 int8 = v78
                                                            fmt.Println(fi65)
                                                            v2 += 255
                                                            _ = fi77
                                                            _ = v78
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                    }
                                                    var v85 float32 = float32((v1 + (((float64(-2.25) * v1) - (v1 * v1)) + v1)))
                                                    _ = v85
                                                } else {
                                                    switch ("hello" + ("a" + "\n\t")) {
                                                    case "":
                                                        v86 := GenS8{F0: 4294967295, F1: 878}
                                                        _ = v86
                                                    }
                                                }
                                                v1 = (float64((0.1 + 1e10)) - v1)
                                                var v87 bool = v19
                                                _ = v87
                                            }
                                            _ = fi65
                                        }
                                    }
                                }
                                v19 = v19
                            }
                            v51 = append(v51, (v64 * (v64 ^ (v64 >> ((0) & 63)))))
                            _ = v64
                        }
                        v19 = (int16(v3) <= -865)
                    }
                    for fi88 := 0; fi88 < 4; fi88++ {
                        v51 = append(v51, (fi88 + fi88))
                        if len(v51) >= 2 {
                            v51 = v51[1:len(v51)]
                        }
                        var v89 uint16 = 255
                        _ = fi88
                        _ = v89
                    }
                } else {
                    v3 = v3
                }
                if len(v51) >= 2 {
                    v51 = v51[1:len(v51)]
                }
            }
        }
    case 128:
        var v90 int32 = 0
        _ = v90
    default:
        for fi91 := 0; fi91 < 14; fi91++ {
            for fi92 := 0; fi92 < 15; fi92++ {
                var v93 int8 = -1
                v93 -= (v93 << ((741) & 63))
                if (!true) {
                    if len(v51) > 0 {
                        v51[0] = (-(fi91))
                    }
                } else {
                    switch "a" {
                    case "":
                        if (!v19) {
                            v94 := GenS14{F0: v2, F1: v52}
                            fmt.Println(fi91, v52, v93)
                            _ = v94
                        }
                    default:
                        var v95 int64 = 351
                        _ = v95
                    }
                }
                _ = fi92
                _ = v93
            }
            var v96 uint64 = (v3 ^ (v3 >> ((18446744073709551615) & 63)))
            var v97 float64 = v1
            _ = fi91
            _ = v96
            _ = v97
        }
    }
    v98 := GenS3{F0: (v3 - v3), F1: -1, F2: 42}
    v99 := GenS3{F0: (v3 ^ 919), F1: 377, F2: v3}
    var v100 string = "foo bar"
    fmt.Println(v100, v52)
    if ((4294967295 <= (^(v3))) || (true || (1 > 2147483647))) {
        v101 := GenS15{F0: 255, F1: (v100 + ("a" + "hello"))}
        fmt.Println(v1, v1)
        _ = v101
    } else {
        v102 := GenS16{F0: 82}
        fmt.Println(v19, v100)
        _ = v102
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v100
    _ = v19
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v98
    _ = v99
}
