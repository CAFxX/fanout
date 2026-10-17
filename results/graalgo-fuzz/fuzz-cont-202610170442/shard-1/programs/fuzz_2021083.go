package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int }
type GenS2 struct { F0 uint64; F1 int8; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 string; F1 uint16; F2 uint32 }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 uint16; F1 int32; F2 int }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 int8; F1 bool }
func genFn0(p0 uint64, p1 uint) bool {
    v6 := GenS1{F0: 9223372036854775807, F1: 1, F2: 124}
    var v7 string = "hello"
    if true {
        var v8 uint = p1
        v9 := []int{32768, 32767, 672}
        _ = v8
        _ = v9
    } else {
        var v10 bool = true
        _ = v10
    }
    if true {
        v11 := GenS1{F0: -9223372036854775808, F1: 0, F2: 0}
        v12 := []int{87, -1, int(p1), 91}
        _ = v11
        _ = v12
    } else {
        v13 := GenS1{F0: 1, F1: 1, F2: -9223372036854775808}
        switch "foo bar" {
        case "\n\t":
            var v14 int8 = 100
            _ = v14
        case "":
            if ((2147483648 < 1) && (749 == 1)) {
                p0 += 4294967295
                var v15 uint = 4294967295
                var v16 string = "\n\t"
                _ = v15
                _ = v16
            }
        case "hello":
            v17 := []int{}
            _ = v17
        }
        p1 -= p1
        _ = v13
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 int) uint {
    var v18 int = p0
    v19 := GenS2{F0: 1, F1: 1, F2: 636}
    _ = p0
    _ = v18
    _ = v19
    return 897
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 int64 = -9223372036854775808
    var v4 bool = (0 >= -1)
    var v5 uint = 65535
    var v20 bool = v4
    fmt.Println(v3, v3)
    fmt.Println(v20, v5)
    fmt.Println(v5)
    fmt.Println(v5, v3)
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := []int{(^(fi21)), (fi21 & (fi21 | 9223372036854775807)), fi21}
        v2[(fi21 + (fi21 ^ -2147483648))] = (!(v20 || v4))
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        if v20 {
            v3 += 1
            switch (fi23 % (((fi23 * (fi23 - (fi23 / (((-(fi23))) | 1))))) | 1)) {
            case 42:
                fmt.Println(v4, v20, fi23)
            case 1:
                fmt.Println(v4, v3, fi23)
            }
            fmt.Println(v5, v20)
        }
        if true {
            switch fi23 {
            case -638:
                v20 = true
            case 32768:
                switch ("a" + "foo bar") {
                case "foo bar":
                    v4 = false
                case "hello":
                    v24 := GenS1{F0: (v3 & v3), F1: 64, F2: fi23}
                    _ = v24
                }
            }
        }
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := []int{fi23}
            if (true && ((!(2 >= 255)) || v4)) {
                fmt.Println(v20, v4)
                var v27 float32 = 1.5
                v20 = (397 >= 42)
                _ = v27
            } else {
                v5 -= 18446744073709551615
                var v28 int32 = 42
                _ = v28
            }
            _ = fi25
            _ = v26
        }
        _ = fi23
    }
    delete(v2, -524)
    switch 127 {
    case 42:
        v20 = true
    case -9223372036854775808:
        var v29 uint32 = 2
        _ = v29
    default:
        if v20 {
            var v30 int = -193
            _ = v30
        }
    }
    var v31 uint16 = 2
    var v32 string = ((("hello" + "foo bar") + ("\n\t" + "hello")) + "")
    for fi33 := 0; fi33 < 5; fi33++ {
        for fi34 := 0; fi34 < 1; fi34++ {
            fmt.Println(v5)
            v31 += uint16(v5)
            v31 += v31
            _ = fi34
        }
        for fi35 := 0; fi35 < 13; fi35++ {
            fmt.Println(fi35)
            _ = fi35
        }
        v36 := GenS3{F0: 127}
        _ = fi33
        _ = v36
    }
    if (109 >= 74) {
        if (v4 && v20) {
            v37 := v2[127]
            if (127 != 42) {
                for fi38 := 0; fi38 < 11; fi38++ {
                    switch v32 {
                    case "\n\t":
                        for fi39 := 0; fi39 < 6; fi39++ {
                            switch "hello" {
                            case "\n\t":
                                v3 = (v3 << ((0) & 63))
                            case "foo bar":
                                v40 := []int{-804, fi38, 128, (fi38 & (fi38 % (((fi38 + fi39)) | 1)))}
                                if (-128 <= 42) {
                                    var v41 int16 = 1000
                                    var v42 float64 = ((0.0 * ((float64(3.14159) - -2.25) * 1.5)) - (((0.1 - 0.0) + 1.5) / (0.0 + (3.14159 / (1e10 + 0.0)))))
                                    _ = v41
                                    _ = v42
                                } else {
                                    v5 += (v5 % ((genFn1((fi39 << (((v5 * genFn1(fi39))) & 63)))) | 1))
                                    fmt.Println(v32)
                                }
                                v32 = v32
                                _ = v40
                            case "a":
                                fmt.Println(v3, v3)
                            default:
                                if v20 {
                                    fmt.Println(fi38, fi38)
                                }
                            }
                            switch (fi39 | (-(fi39))) {
                            case 32767:
                                fmt.Println(fi39, fi39)
                            case 2:
                                v43 := GenS2{F0: 741, F1: -128, F2: (v5 - v5)}
                                _ = v43
                            case -9223372036854775808:
                                var v44 int32 = -876
                                _ = v44
                            default:
                                v37 = v4
                            }
                            switch "a" {
                            case "hello":
                                v45 := GenS4{F0: v31, F1: 77}
                                _ = v45
                            case "\n\t":
                                var v46 int8 = 1
                                _ = v46
                            }
                            _ = fi39
                        }
                    case "":
                        if (v3 == v3) {
                            if (v3 > v3) {
                                if v4 {
                                    fmt.Println(v5, fi38, fi38)
                                    if v4 {
                                        v20 = (!true)
                                    } else {
                                        v47 := GenS5{F0: -1, F1: 500}
                                        var v48 int8 = -37
                                        _ = v47
                                        _ = v48
                                    }
                                    var v49 uint32 = 42
                                    _ = v49
                                } else {
                                    if v4 {
                                        v50 := []int{fi38}
                                        switch 32768 {
                                        case 900:
                                            if v4 {
                                                var v51 uint64 = 1
                                                _ = v51
                                            } else {
                                                v52 := GenS5{F0: -1, F1: 4294967295}
                                                _ = v52
                                            }
                                        default:
                                            var v53 float64 = 1e10
                                            _ = v53
                                        }
                                        var v54 float64 = (float64(float32((1.5 * (1.5 + 1e10)))) - -2.25)
                                        _ = v50
                                        _ = v54
                                    } else {
                                        for fi55 := 0; fi55 < 1; fi55++ {
                                            v32 = v32
                                            delete(v2, fi55)
                                            if true {
                                                v56 := v2[(-(fi38))]
                                                var v57 uint = (v5 + v5)
                                                _ = v56
                                                _ = v57
                                            }
                                            var v58 float32 = ((1.5 * (1.5 / (3.14159 / (-2.25 / float32(1.5))))) * (0.1 * -2.25))
                                            v58 += (v58 + 1.5)
                                            v5 = (v5 % (((v5 | (v5 + (v5 | uint(v31))))) | 1))
                                            v32 = v32
                                            _ = fi55
                                            _ = v58
                                        }
                                        for fi59 := 0; fi59 < 9; fi59++ {
                                            v60 := v2[fi38]
                                            v20 = ((!false) && (!((v31 << ((168) & 63)) != 255)))
                                            if v37 {
                                                fmt.Println(v32, v5, v3)
                                                if true {
                                                    var v61 float32 = 1.5
                                                    if true {
                                                        v3 -= 42
                                                        var v62 uint8 = 42
                                                        var v63 int16 = 42
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        v64 := GenS2{F0: 795, F1: 0, F2: v5}
                                                        var v65 float64 = 1.5
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    _ = v61
                                                } else {
                                                    v66 := []int{fi59}
                                                    fmt.Println(v3)
                                                    v32 = "\n\t"
                                                    v2[fi38] = (v32 >= "")
                                                    var v67 int8 = 120
                                                    _ = v66
                                                    _ = v67
                                                }
                                            }
                                            _ = fi59
                                            _ = v60
                                        }
                                    }
                                    var v68 uint32 = 404
                                    var v69 float64 = (0.0 + (0.1 / (3.14159 / 1.5)))
                                    _ = v68
                                    _ = v69
                                }
                            } else {
                                for fi70 := 0; fi70 < 5; fi70++ {
                                    v3 = v3
                                    var v71 string = ((v32 + (("" + ((v32 + ("a" + v32)) + v32)) + v32)) + "\n\t")
                                    if true {
                                        var v72 uint8 = 32
                                        if false {
                                            v31 -= 255
                                        }
                                        v20 = (123 <= 100)
                                        _ = v72
                                    } else {
                                        var v73 float32 = 1e10
                                        v4 = v37
                                        _ = v73
                                    }
                                    _ = fi70
                                    _ = v71
                                }
                                var v74 uint8 = 0
                                v75 := GenS4{F0: v31, F1: 123}
                                _ = v74
                                _ = v75
                            }
                        } else {
                            v37 = true
                        }
                    }
                    var v76 int8 = 127
                    _ = fi38
                    _ = v76
                }
                var v77 uint = v5
                _ = v77
            }
            _ = v37
        }
        _ = len(v2)
    } else {
        var v78 uint32 = 1
        var v79 bool = v20
        _ = v78
        _ = v79
    }
    v80 := GenS6{F0: "foo bar", F1: (^(v31)), F2: 0}
    var v81 uint32 = 2147483648
    var v82 bool = false
    for fi83 := 0; fi83 < 12; fi83++ {
        for fi84 := 0; fi84 < 8; fi84++ {
            v4 = true
            _ = fi84
        }
        for fi85 := 0; fi85 < 14; fi85++ {
            var v86 uint = (^(v5))
            if (!false) {
                fmt.Println(v32, v4, v86)
            } else {
                var v87 int64 = v3
                if v82 {
                    if false {
                        v87 = v87
                        switch ("hello" + (v32 + (((("hello" + ("\n\t" + "\n\t")) + "a") + v32) + ("a" + (v32 + "foo bar"))))) {
                        case "foo bar":
                            fmt.Println(v87, fi83)
                        case "":
                            switch (("hello" + "\n\t") + ("a" + "hello")) {
                            case "\n\t":
                                if false {
                                    v81 = (^(v81))
                                } else {
                                    v88 := GenS7{F0: 188, F1: v31}
                                    v89 := GenS6{F0: v32, F1: (^(v31)), F2: v81}
                                    var v90 int32 = 0
                                    _ = v88
                                    _ = v89
                                    _ = v90
                                }
                            case "hello":
                                v86 += (v5 / ((v5) | 1))
                            case "a":
                                delete(v1, (fi83 * 294))
                            default:
                                var v91 bool = false
                                _ = v91
                            }
                        case "a":
                            v92 := GenS5{F0: 42, F1: 2147483648}
                            _ = v92
                        }
                        if (0 >= -1) {
                            v93 := GenS8{F0: v31, F1: int32(v81), F2: 128}
                            fmt.Println(fi83)
                            var v94 int64 = v3
                            _ = v93
                            _ = v94
                        }
                    } else {
                        v95 := v1[int(v86)]
                        if (4294967295 != uint64(v81)) {
                            if (18446744073709551615 >= 42) {
                                v86 += 255
                                if v4 {
                                    switch fi83 {
                                    case -1:
                                        v96 := GenS9{F0: 1, F1: int(v3)}
                                        v97 := GenS7{F0: 1, F1: 32767}
                                        v86 -= uint(v31)
                                        _ = v96
                                        _ = v97
                                    default:
                                        fmt.Println(v5, v20)
                                    }
                                    v98 := GenS4{F0: v31, F1: 0}
                                    _ = len(v1)
                                    _ = v98
                                }
                            }
                            v31 += 65535
                        }
                        _ = v95
                    }
                } else {
                    fmt.Println(v4)
                    fmt.Println(v5, v20, v3)
                    var v99 float64 = (1.5 * 0.1)
                    _ = v99
                }
                _ = v87
            }
            v100 := GenS6{F0: "hello", F1: v31, F2: v81}
            switch "a" {
            case "":
                v101 := []int{fi85}
                _ = v101
            case "\n\t":
                v3 += 9223372036854775807
            default:
                switch (v32 + v32) {
                case "a":
                    var v102 int = (fi85 << ((255) & 63))
                    _ = v102
                default:
                    v4 = v20
                }
                if (1 == -62) {
                    v3 -= (v3 - v3)
                    v103 := GenS10{F0: int8(v81), F1: true}
                    v104 := []int{}
                    var v105 int32 = 42
                    _ = v103
                    _ = v104
                    _ = v105
                } else {
                    var v106 int8 = 42
                    _ = v106
                }
                v31 = (v31 * (^(v31)))
            }
            _ = fi85
            _ = v100
            _ = v86
        }
        v5 = uint(v81)
        _ = fi83
    }
    var v107 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v107:", v107)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v107
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v80
    _ = v81
    _ = v82
}
