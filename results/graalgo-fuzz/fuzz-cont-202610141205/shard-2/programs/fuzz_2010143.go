package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint; F2 int64 }
type GenS3 struct { F0 bool; F1 uint32; F2 int32 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 int64; F1 string; F2 int32 }
type GenS8 struct { F0 string; F1 uint32; F2 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS12 struct { F0 bool; F1 int32; F2 uint32 }
type GenS13 struct { F0 int; F1 string }
func genFn0(p0 int32, p1 uint64, p2 uint8) string {
    var v4 int16 = 0
    fmt.Println(p2)
    fmt.Println(p2)
    if false {
        for fi5 := 0; fi5 < 17; fi5++ {
            for fi6 := 0; fi6 < 16; fi6++ {
                var v7 float32 = -2.25
                if false {
                    v8 := []int{}
                    _ = v8
                } else {
                    var v9 int32 = (p0 * (p0 & 1))
                    v10 := GenS1{F0: (v4 << ((65535) & 63))}
                    _ = v10
                    _ = v9
                }
                if false {
                    v11 := GenS2{F0: 18446744073709551615, F1: 0, F2: -9223372036854775808}
                    p2 += (p2 & (^(p2)))
                    var v12 string = "a"
                    if true {
                        var v13 uint = 255
                        _ = v13
                    } else {
                        var v14 uint32 = 844
                        if false {
                            var v15 uint = 827
                            if (v14 == (v14 % ((uint32(v15)) | 1))) {
                                var v16 float64 = -2.25
                                _ = v16
                            }
                            switch v12 {
                            case "hello":
                                if (34 < 45) {
                                    if (false || false) {
                                        var v17 uint = 561
                                        var v18 string = (v12 + "")
                                        _ = v17
                                        _ = v18
                                    }
                                    switch 4294967295 {
                                    case 128:
                                        p1 = 18446744073709551615
                                        var v19 uint8 = (p2 + p2)
                                        fmt.Println(p0)
                                        _ = v19
                                    default:
                                        v12 = "foo bar"
                                    }
                                }
                            case "foo bar":
                                v20 := []int{fi6, (fi6 ^ 9223372036854775807), (-(fi6)), fi6}
                                _ = v20
                            case "":
                                v21 := []int{(-(fi6)), (fi5 - (fi5 + (fi6 - (^(fi6)))))}
                                _ = v21
                            default:
                                var v22 uint = uint(p1)
                                _ = v22
                            }
                            var v23 uint16 = 2
                            var v24 int64 = 1
                            _ = v15
                            _ = v23
                            _ = v24
                        }
                        _ = v14
                    }
                    v25 := GenS3{F0: true, F1: 0, F2: p0}
                    _ = v11
                    _ = v12
                    _ = v25
                }
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
    } else {
        v26 := GenS4{F0: (("a" + "\n\t") + ("" + "foo bar")), F1: "hello"}
        p0 = (p0 ^ (p0 % ((p0) | 1)))
        _ = v26
    }
    var v27 uint8 = (p2 - p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v4
    return "a"
}

func main() {
    var v1 float64 = 0.0
    var v2 int16 = -359
    v3 := make(map[string]uint32)
    switch 184 {
    case 32768:
        var v28 uint32 = 65535
        _ = v28
    default:
        var v29 int8 = 80
        _ = v29
    }
    delete(v3, "\n\t")
    if true {
        fmt.Println(v2)
        fmt.Println(v2, v1)
        if (0 > 763) {
            v30 := []int{-1, 2147483647, 554, 9223372036854775807}
            var v31 float64 = float64(0.1)
            if len(v30) > 0 {
                v30[0] = 4294967295
            }
            _ = v30
            _ = v31
        }
    } else {
        var v32 int16 = (v2 & (-(v2)))
        fmt.Println(v2, v2)
        fmt.Println(v32, v2, v32)
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := GenS3{F0: ((("hello" + "foo bar") >= "foo bar") && (18446744073709551615 != 42)), F1: 4294967295, F2: -842}
        if true {
            for fi35 := 0; fi35 < 1; fi35++ {
                v1 -= ((3.14159 / v1) - float64((0.1 * 0.1)))
                fmt.Println(fi33, v2)
                v36 := []int{(fi33 % (((fi35 + fi33)) | 1)), fi35, 252, fi35}
                v2 -= (-(v2))
                v37 := GenS5{F0: v2, F1: 42}
                if ((1e10 + v1) <= (float64(3.14159) / ((0.0 - v1) - v1))) {
                    var v38 uint16 = 1
                    fmt.Println(v2, v1)
                    _ = v38
                }
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        var v39 bool = true
        _ = fi33
        _ = v34
        _ = v39
    }
    if false {
        v40 := v3[("a" + ((("\n\t" + "") + ((("foo bar" + "hello") + "foo bar") + "foo bar")) + ("" + (("a" + "hello") + "hello"))))]
        _ = v40
    } else {
        v3["\n\t"] = 12
    }
    var v41 string = "foo bar"
    if false {
        for fi42 := 0; fi42 < 7; fi42++ {
            switch ("a" + v41) {
            case "\n\t":
                if (0.0 == v1) {
                    v43 := GenS4{F0: "", F1: v41}
                    var v44 int64 = 673
                    var v45 float64 = 1e10
                    _ = v43
                    _ = v44
                    _ = v45
                } else {
                    v2 = (^(v2))
                }
            case "a":
                for fi46 := 0; fi46 < 18; fi46++ {
                    if false {
                        var v47 string = (v41 + v41)
                        v48 := GenS6{F0: 2, F1: ("a" + ("a" + v41))}
                        _ = v47
                        _ = v48
                    }
                    fmt.Println(v1)
                    var v49 float32 = 3.14159
                    _ = fi46
                    _ = v49
                }
            }
            fmt.Println(v41)
            _ = fi42
        }
        v50 := []int{915, -1, 128, -533}
        var v51 float64 = v1
        _ = v50
        _ = v51
    } else {
        v1 -= v1
    }
    switch v41 {
    case "":
        var v52 uint64 = 519
        _ = v52
    case "a":
        for fi53 := 0; fi53 < 7; fi53++ {
            v54 := []int{fi53, -295}
            v55 := GenS7{F0: 9223372036854775807, F1: "", F2: 2147483647}
            if true {
                var v56 int8 = -1
                var v57 float64 = ((v1 + v1) + float64((0.1 + 3.14159)))
                v1 = v57
                _ = v56
                _ = v57
            } else {
                if (0 != 42) {
                    for fi58 := 0; fi58 < 11; fi58++ {
                        fmt.Println(fi53)
                        switch fi53 {
                        case -901:
                            var v59 int16 = (v2 - v2)
                            _ = v59
                        case 32767:
                            v60 := GenS8{F0: v41, F1: 42, F2: v2}
                            _ = v60
                        case 9223372036854775807:
                            var v61 uint = 1
                            _ = v61
                        }
                        fmt.Println(v41)
                        var v62 uint16 = 32767
                        _ = fi58
                        _ = v62
                    }
                } else {
                    for fi63 := 0; fi63 < 16; fi63++ {
                        switch fi53 {
                        case 1:
                            v54 = append(v54, 42)
                        case -1:
                            v64 := GenS7{F0: -283, F1: v41, F2: 828}
                            v65 := GenS6{F0: 42, F1: "hello"}
                            var v66 int64 = 42
                            _ = v64
                            _ = v65
                            _ = v66
                        case -2147483648:
                            v54 = append(v54, (fi63 * (fi63 + (fi53 % (((fi53 * (fi53 ^ 128))) | 1)))))
                        default:
                            if false {
                                var v67 uint8 = 1
                                _ = v67
                            }
                        }
                        _ = fi63
                    }
                    if (128 < 127) {
                        v68 := GenS9{F0: 4294967295}
                        var v69 int = (fi53 | (fi53 % ((-1) | 1)))
                        if (float64((1e10 * 3.14159)) == (((v1 - float64(-2.25)) / float64(3.14159)) / float64((-2.25 * (3.14159 - (0.0 / (0.1 / 1.5))))))) {
                            fmt.Println(fi53)
                            v70 := GenS2{F0: 65535, F1: 2, F2: 42}
                            var v71 uint8 = 127
                            _ = v70
                            _ = v71
                        } else {
                            switch int(v2) {
                            case 0:
                                v54 = append(v54, -214)
                            case 42:
                                if (2 <= 2) {
                                    var v72 bool = ((true && false) && (-380 >= 1))
                                    _ = v72
                                }
                            }
                        }
                        _ = v68
                        _ = v69
                    }
                    if true {
                        var v73 bool = false
                        v73 = (2 <= 255)
                        if (true && v73) {
                            for fi74 := 0; fi74 < 13; fi74++ {
                                if v73 {
                                    fmt.Println(fi74, v73)
                                    if (1 >= -320) {
                                        var v75 uint = 269
                                        v76 := GenS9{F0: 2}
                                        _ = v75
                                        _ = v76
                                    } else {
                                        v41 = (v41 + (("\n\t" + v41) + v41))
                                        if (false && (2 >= 42)) {
                                            var v77 uint16 = 1
                                            if (!false) {
                                                v78 := GenS10{F0: (v77 + v77)}
                                                if len(v54) >= 2 {
                                                    v54 = v54[1:len(v54)]
                                                }
                                                _ = v78
                                            } else {
                                                v3[(v41 + "\n\t")] = 65535
                                                var v79 int32 = -566
                                                var v80 int16 = -825
                                                _ = v79
                                                _ = v80
                                            }
                                            _ = v77
                                        }
                                        fmt.Println(v1, v1, fi53)
                                    }
                                    var v81 uint8 = 128
                                    _ = v81
                                } else {
                                    fmt.Println(fi53, v41, v73)
                                    v3[(v41 + v41)] = 65535
                                }
                                v82 := GenS11{F0: 2147483647, F1: 890, F2: 1}
                                v83 := GenS12{F0: (3.14159 <= float32(float64(0.1))), F1: -1, F2: 65535}
                                v73 = v73
                                v54 = append(v54, (-(fi53)))
                                _ = fi74
                                _ = v82
                                _ = v83
                            }
                            for fi84 := 0; fi84 < 8; fi84++ {
                                switch "\n\t" {
                                case "hello":
                                    v85 := GenS10{F0: 42}
                                    _ = v85
                                }
                                v2 += v2
                                if ((-(fi53)) <= (fi53 | fi53)) {
                                    v86 := GenS13{F0: (fi84 % ((-1) | 1)), F1: v41}
                                    _ = v86
                                }
                                _ = fi84
                            }
                        }
                        _ = v73
                    } else {
                        for fi87 := 0; fi87 < 15; fi87++ {
                            v1 += v1
                            _ = fi87
                        }
                        v1 = float64(0.0)
                        switch v41 {
                        case "foo bar":
                            v88 := GenS12{F0: (92 < 19), F1: 2147483647, F2: 4294967295}
                            _ = v88
                        case "":
                            if (((-9223372036854775808 <= int64(fi53)) && (((3.14159 + v1) * ((v1 + float64((0.1 + (1.5 * 0.1)))) / v1)) >= v1)) && true) {
                                var v89 int16 = v2
                                _ = v89
                            }
                        case "\n\t":
                            for fi90 := 0; fi90 < 3; fi90++ {
                                var v91 bool = (-128 == -112)
                                switch ("hello" + "hello") {
                                case "hello":
                                    v92 := GenS7{F0: 0, F1: (("\n\t" + (v41 + "\n\t")) + "hello"), F2: 925}
                                    _ = v92
                                case "":
                                    var v93 int32 = 0
                                    _ = v93
                                }
                                var v94 int32 = 1
                                _ = fi90
                                _ = v91
                                _ = v94
                            }
                        default:
                            fmt.Println(v41)
                        }
                    }
                }
                var v95 bool = (false && false)
                _ = v95
            }
            _ = fi53
            _ = v54
            _ = v55
        }
    case "\n\t":
        for fi96 := 0; fi96 < 10; fi96++ {
            v97 := []int{fi96, fi96, (fi96 + (fi96 + (fi96 ^ fi96))), 32768}
            v98 := GenS8{F0: (v41 + v41), F1: 2147483648, F2: v2}
            if (573 > 1) {
                switch "" {
                case "\n\t":
                    var v99 bool = (false || true)
                    _ = v99
                }
                v97 = append(v97, (fi96 / (((fi96 << ((4294967295) & 63))) | 1)))
                for fi100 := 0; fi100 < 11; fi100++ {
                    var v101 int32 = -1
                    _ = fi100
                    _ = v101
                }
            } else {
                if len(v97) >= 2 {
                    v97 = v97[1:len(v97)]
                }
                fmt.Println(v2, v1, v1)
                var v102 bool = false
                _ = v102
            }
            _ = fi96
            _ = v97
            _ = v98
        }
    default:
        v41 = v41
    }
    var v103 float32 = ((0.0 - 3.14159) - 0.0)
    if (0.1 == float64(v103)) {
        var v104 bool = true
        _ = v104
    }
    switch -2147483648 {
    case 4294967295:
        for fi105 := 0; fi105 < 17; fi105++ {
            v106 := GenS13{F0: fi105, F1: (("\n\t" + ((v41 + "hello") + ("\n\t" + "foo bar"))) + (v41 + v41))}
            fmt.Println(v2, fi105)
            v41 = (v41 + (("\n\t" + v41) + ((v41 + (("hello" + v41) + v41)) + v41)))
            _ = fi105
            _ = v106
        }
    default:
        var v107 int8 = -70
        _ = v107
    }
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v41
}
