package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int8; F2 int8 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint8; F1 int8; F2 int16 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool; F1 int8; F2 uint }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 bool; F1 uint64 }
type GenS13 struct { F0 string; F1 uint64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int16; F1 int32 }
func genFn0(p0 uint16, p1 uint32, p2 int16) uint16 {
    var v5 int = 42
    p2 -= -1
    v5 -= (v5 ^ v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 uint16) uint8 {
    fmt.Println(p0)
    v6 := GenS1{F0: true, F1: 65535, F2: 9223372036854775807}
    var v7 int16 = -1
    v8 := []int{-307}
    v9 := GenS1{F0: (!false), F1: 65535, F2: -773}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 157
}

func main() {
    var v1 uint = 486
    var v2 float64 = 0.1
    var v3 float64 = v2
    v4 := make(map[string]bool)
    var v10 uint64 = 18446744073709551615
    fmt.Println(v2, v10)
    fmt.Println(v2)
    if true {
        if (2 <= 255) {
            if (!(128 <= 1)) {
                delete(v4, "")
                v11 := GenS1{F0: (-128 <= 42), F1: (v1 | 65535), F2: -611}
                v12 := GenS2{F0: int16(v1)}
                _ = v11
                _ = v12
            }
            v10 = (v10 << (((v1 * 0)) & 63))
            if (-249 > 2147483647) {
                if true {
                    fmt.Println(v2, v2, v3)
                    for fi13 := 0; fi13 < 20; fi13++ {
                        for fi14 := 0; fi14 < 10; fi14++ {
                            switch "" {
                            case "\n\t":
                                v15 := GenS3{F0: 0, F1: 127, F2: 127}
                                if (630 > fi13) {
                                    var v16 float32 = 3.14159
                                    _ = v16
                                } else {
                                    if (v1 > v1) {
                                        v3 = (1.5 + v2)
                                        var v17 uint8 = 125
                                        v18 := GenS3{F0: -1, F1: 40, F2: 42}
                                        v1 = 1
                                        _ = v17
                                        _ = v18
                                    } else {
                                        var v19 int16 = 1000
                                        switch fi14 {
                                        case 42:
                                            fmt.Println(fi13, v19, v3)
                                            var v20 int64 = 9223372036854775807
                                            _ = v20
                                        case -925:
                                            switch 0 {
                                            case -2147483648:
                                                v21 := []int{fi13}
                                                _ = v21
                                            case 42:
                                                v22 := []int{42}
                                                _ = v22
                                            case 145:
                                                v3 += 1.5
                                            }
                                        default:
                                            v1 -= (^(v1))
                                        }
                                        if true {
                                            v3 = v2
                                            if (2 == 1) {
                                                var v23 uint64 = 42
                                                fmt.Println(v23, v1)
                                                if true {
                                                    v19 = v19
                                                    v2 = float64(3.14159)
                                                    var v24 uint32 = 404
                                                    v10 = (v23 << (((v1 << ((v1) & 63))) & 63))
                                                    _ = v24
                                                }
                                                _ = v23
                                            }
                                        } else {
                                            if false {
                                                var v25 uint16 = 32767
                                                _ = v25
                                            } else {
                                                var v26 string = "hello"
                                                v27 := []int{1, (^(fi14)), (fi13 / (((fi13 - (fi14 | (fi14 % ((fi13) | 1))))) | 1))}
                                                if false {
                                                    switch (fi14 ^ 193) {
                                                    case 128:
                                                        var v28 uint32 = 2147483647
                                                        _ = v28
                                                    case -1:
                                                        v1 = (^(v1))
                                                        var v29 float64 = 1.5
                                                        fmt.Println(v10)
                                                        _ = v29
                                                    case 354:
                                                        v30 := GenS3{F0: v19, F1: -128, F2: 127}
                                                        _ = v30
                                                    default:
                                                        v3 = (float64((-2.25 + (0.0 / 1.0))) + v3)
                                                        var v31 float32 = ((1.5 / 1e10) + -2.25)
                                                        if false {
                                                            var v32 int = fi14
                                                            _ = v32
                                                        }
                                                        _ = v31
                                                    }
                                                    fmt.Println(fi13, fi14, v3)
                                                } else {
                                                    if false {
                                                        v10 = v10
                                                    } else {
                                                        if true {
                                                            switch (fi13 / ((int(v10)) | 1)) {
                                                            case -564:
                                                                fmt.Println(v3)
                                                            default:
                                                                if (42 == 1) {
                                                                    fmt.Println(v10, fi14, v19)
                                                                    _ = len(v4)
                                                                    fmt.Println(fi13, v26, fi13)
                                                                }
                                                            }
                                                            var v33 int32 = -2147483648
                                                            fmt.Println(v19)
                                                            _ = v33
                                                        } else {
                                                            v10 -= (v10 + (v10 << ((18446744073709551615) & 63)))
                                                            v1 = v1
                                                            switch (fi13 ^ fi13) {
                                                            case 32767:
                                                                if len(v27) > 0 {
                                                                    v27[0] = fi13
                                                                }
                                                            case 311:
                                                                v27 = append(v27, (^(fi13)))
                                                                v1 = (v1 + v1)
                                                            default:
                                                                v34 := GenS4{F0: (v19 | v19), F1: 42}
                                                                _ = v34
                                                            }
                                                            v10 = (v10 % (((v10 & (v10 - v10))) | 1))
                                                        }
                                                    }
                                                    v35 := GenS5{F0: 127, F1: 100, F2: 0}
                                                    v36 := GenS3{F0: 32767, F1: -128, F2: -41}
                                                    v37 := v4[v26]
                                                    if true {
                                                        v4[(("hello" + v26) + "foo bar")] = v37
                                                        fmt.Println(fi14, v19)
                                                        fmt.Println(v3)
                                                    }
                                                    fmt.Println(v10, v3, v37)
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                }
                                                var v38 uint16 = 591
                                                fmt.Println(v38)
                                                _ = v26
                                                _ = v27
                                                _ = v38
                                            }
                                            v3 -= float64((((((1.5 + -2.25) + 1.5) / (-2.25 - (3.14159 / 1.0))) + (1.5 - (float32(v2) - (float32(float64((0.0 + 1e10))) / -2.25)))) / 1e10))
                                        }
                                        _ = v19
                                    }
                                    v39 := GenS3{F0: 246, F1: 21, F2: -128}
                                    _ = v39
                                }
                                if false {
                                    var v40 uint64 = (v10 - v10)
                                    _ = v40
                                }
                                _ = v15
                            case "hello":
                                fmt.Println(v10, v2, v2)
                            case "":
                                switch (fi14 << ((2) & 63)) {
                                case 128:
                                    switch "hello" {
                                    case "hello":
                                        var v41 uint64 = (v10 * v10)
                                        _ = v41
                                    default:
                                        switch 1 {
                                        case 127:
                                            var v42 int = fi14
                                            _ = v42
                                        case 112:
                                            if (32767 < 590) {
                                                v4["hello"] = ((!(!false)) && ((v1 & (v1 << (((v1 - (v1 + 2))) & 63))) != 42))
                                            } else {
                                                v43 := GenS4{F0: -1, F1: 893}
                                                var v44 float32 = 0.0
                                                if (!true) {
                                                    if false {
                                                        if true {
                                                            if false {
                                                                var v45 int = fi13
                                                                v46 := v4[("foo bar" + "foo bar")]
                                                                switch "" {
                                                                case "hello":
                                                                    if (false && ((3.14159 + v44) == 3.14159)) {
                                                                        if (-76 > 100) {
                                                                            if v46 {
                                                                                if v46 {
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        switch 9223372036854775807 {
                                                                                        case 128:
                                                                                            if false {
                                                                                                if true {
                                                                                                    if (((v1 ^ v1) >= 255) && v46) {
                                                                                                        v47 := GenS2{F0: 0}
                                                                                                        v1 = uint(v10)
                                                                                                        _ = v47
                                                                                                    } else {
                                                                                                        var v48 uint = 18446744073709551615
                                                                                                        v1 -= (v48 / ((uint(v10)) | 1))
                                                                                                        if v46 {
                                                                                                            var v49 uint8 = genFn1(0)
                                                                                                            v49 -= 2
                                                                                                            v50 := GenS6{F0: (true || (335 <= -1)), F1: 127}
                                                                                                            v2 = v2
                                                                                                            _ = v49
                                                                                                            _ = v50
                                                                                                        } else {
                                                                                                            fmt.Println(v1, fi14)
                                                                                                            switch ("\n\t" + "") {
                                                                                                            case "":
                                                                                                                v1 += 18446744073709551615
                                                                                                            }
                                                                                                        }
                                                                                                        v51 := []int{(v45 % ((4294967295) | 1)), (fi14 >> (((v48 | 255)) & 63)), -1, v45}
                                                                                                        v46 = false
                                                                                                        fmt.Println(fi14, fi13, v1)
                                                                                                        if len(v51) >= 2 {
                                                                                                            v51 = v51[1:len(v51)]
                                                                                                        }
                                                                                                        _ = v48
                                                                                                        _ = v51
                                                                                                    }
                                                                                                    v52 := GenS6{F0: v46, F1: -128}
                                                                                                    var v53 int = (fi13 ^ fi14)
                                                                                                    v45 = fi14
                                                                                                    _ = v52
                                                                                                    _ = v53
                                                                                                } else {
                                                                                                    var v54 uint16 = 42
                                                                                                    var v55 float64 = v2
                                                                                                    _ = v54
                                                                                                    _ = v55
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(fi13, fi13)
                                                                                            }
                                                                                        default:
                                                                                            v45 += fi14
                                                                                        }
                                                                                    case "":
                                                                                        v4[("hello" + "foo bar")] = v46
                                                                                    case "a":
                                                                                        if ((true || ((!(v46 && false)) || v46)) || (v46 || (2 <= 127))) {
                                                                                            v3 -= ((0.0 * (0.0 - v3)) + 3.14159)
                                                                                            switch ("hello" + "\n\t") {
                                                                                            case "":
                                                                                                var v56 uint = v1
                                                                                                _ = v56
                                                                                            case "a":
                                                                                                switch -2147483648 {
                                                                                                case -784:
                                                                                                    var v57 uint32 = 274
                                                                                                    _ = v57
                                                                                                case 42:
                                                                                                    fmt.Println(v45, v46)
                                                                                                case -25:
                                                                                                    var v58 uint32 = 2147483647
                                                                                                    _ = v58
                                                                                                }
                                                                                            default:
                                                                                                if (v10 < (v10 / (((v10 - 769)) | 1))) {
                                                                                                    switch (v45 % ((fi14) | 1)) {
                                                                                                    case -1:
                                                                                                        switch "a" {
                                                                                                        case "\n\t":
                                                                                                            v3 += -2.25
                                                                                                        case "":
                                                                                                            v45 = v45
                                                                                                        default:
                                                                                                            if v46 {
                                                                                                                var v59 string = ""
                                                                                                                var v60 float64 = 0.1
                                                                                                                _ = v59
                                                                                                                _ = v60
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        v4[("foo bar" + ("" + ("a" + "a")))] = ((false || (-125 >= 100)) && v46)
                                                                                                    }
                                                                                                    v61 := GenS6{F0: ((((("hello" + "a") + "") + "a") + (("" + (("\n\t" + "") + "a")) + "foo bar")) > "\n\t"), F1: -100}
                                                                                                    switch (((("\n\t" + ("" + ("" + "\n\t"))) + ("a" + "")) + "a") + "") {
                                                                                                    case "a":
                                                                                                        v44 -= v44
                                                                                                    case "":
                                                                                                        if (!(genFn1(565) != 2)) {
                                                                                                            v62 := GenS3{F0: 1000, F1: 100, F2: -128}
                                                                                                            _ = v62
                                                                                                        } else {
                                                                                                            v63 := GenS3{F0: -299, F1: 1, F2: 0}
                                                                                                            var v64 uint = (v1 & v1)
                                                                                                            v65 := GenS3{F0: -510, F1: 2, F2: 100}
                                                                                                            _ = v63
                                                                                                            _ = v64
                                                                                                            _ = v65
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        v1 -= 442
                                                                                                    }
                                                                                                    fmt.Println(v45)
                                                                                                    switch v45 {
                                                                                                    case 128:
                                                                                                        var v66 int8 = -1
                                                                                                        _ = v66
                                                                                                    case -512:
                                                                                                        v67 := []int{fi13, fi13, 2147483647, (fi13 + -1)}
                                                                                                        _ = v67
                                                                                                    case 576:
                                                                                                        v45 = fi13
                                                                                                        _ = len(v4)
                                                                                                    }
                                                                                                    v1 = 18446744073709551615
                                                                                                    v68 := GenS2{F0: 0}
                                                                                                    _ = v61
                                                                                                    _ = v68
                                                                                                }
                                                                                            }
                                                                                            v46 = (!((v44 - ((v44 * float32((1.5 * v3))) * (((v44 + v44) + (1e10 / 0.1)) * ((float32((v2 + 1.5)) / (float32(v3) + (v44 + float32((v3 + (v2 / (v2 * float64(v44)))))))) * (v44 + 0.0))))) != v44))
                                                                                        } else {
                                                                                            v4[("\n\t" + "")] = true
                                                                                            v69 := GenS1{F0: ((!((!v46) && ((^(v10)) == 18446744073709551615))) || (true && (1 == 127))), F1: 0, F2: 0}
                                                                                            switch fi13 {
                                                                                            case -349:
                                                                                                switch -2147483648 {
                                                                                                case -2147483648:
                                                                                                    switch "\n\t" {
                                                                                                    case "":
                                                                                                        v46 = true
                                                                                                    case "hello":
                                                                                                        delete(v4, (("hello" + "") + (("hello" + ("a" + "hello")) + "foo bar")))
                                                                                                    case "a":
                                                                                                        fmt.Println(v46)
                                                                                                    }
                                                                                                case 2147483647:
                                                                                                    var v70 float64 = (-2.25 / 1.5)
                                                                                                    _ = v70
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                fmt.Println(fi14, v46, v10)
                                                                                                if (v46 || v46) {
                                                                                                    v71 := v4[(("foo bar" + "hello") + "")]
                                                                                                    _ = v71
                                                                                                } else {
                                                                                                    v3 = 3.14159
                                                                                                    v72 := []int{fi13}
                                                                                                    var v73 uint16 = 0
                                                                                                    _ = v72
                                                                                                    _ = v73
                                                                                                }
                                                                                            case 42:
                                                                                                fmt.Println(v45)
                                                                                                var v74 int32 = 0
                                                                                                delete(v4, "")
                                                                                                _ = v74
                                                                                            }
                                                                                            _ = v69
                                                                                        }
                                                                                    }
                                                                                    var v75 int = v45
                                                                                    _ = v75
                                                                                } else {
                                                                                    v46 = v46
                                                                                }
                                                                                var v76 uint64 = (v10 % ((v10) | 1))
                                                                                v77 := []int{(fi14 << (((v1 << ((v1) & 63))) & 63))}
                                                                                _ = v76
                                                                                _ = v77
                                                                            } else {
                                                                                v78 := GenS6{F0: (true || v46), F1: -128}
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    var v79 int = v45
                                                                                    _ = v79
                                                                                }
                                                                                var v80 uint32 = 2147483648
                                                                                _ = v78
                                                                                _ = v80
                                                                            }
                                                                            v81 := GenS7{F0: 630, F1: false}
                                                                            _ = v81
                                                                        }
                                                                    } else {
                                                                        v45 += (fi13 << (((v1 | (v1 | 255))) & 63))
                                                                        v82 := GenS8{F0: 65535}
                                                                        _ = v82
                                                                    }
                                                                case "foo bar":
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        if false {
                                                                            if false {
                                                                                switch (^(fi14)) {
                                                                                case -2147483648:
                                                                                    v83 := []int{-699, (-(fi13)), (fi14 ^ v45), 695}
                                                                                    _ = v83
                                                                                case -69:
                                                                                    var v84 uint8 = uint8(v10)
                                                                                    _ = v84
                                                                                default:
                                                                                    v46 = ((v1 == (^(v1))) || true)
                                                                                }
                                                                            } else {
                                                                                v85 := GenS9{F0: false, F1: 0, F2: (v1 | (v1 & (v1 / ((v1) | 1))))}
                                                                                v86 := GenS10{F0: v1}
                                                                                _ = v85
                                                                                _ = v86
                                                                            }
                                                                            if v46 {
                                                                                var v87 int16 = 1
                                                                                _ = v87
                                                                            } else {
                                                                                var v88 int64 = -1
                                                                                delete(v4, "")
                                                                                v1 = (v1 * (v1 << ((972) & 63)))
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    v1 = v1
                                                                                case "a":
                                                                                    v89 := GenS10{F0: (v1 & 255)}
                                                                                    _ = v89
                                                                                case "hello":
                                                                                    var v90 float32 = float32(v3)
                                                                                    _ = v90
                                                                                default:
                                                                                    v1 = v1
                                                                                }
                                                                                v91 := []int{(^(fi13))}
                                                                                _ = v88
                                                                                _ = v91
                                                                            }
                                                                        }
                                                                    case "foo bar":
                                                                        var v92 int32 = 1
                                                                        _ = v92
                                                                    default:
                                                                        var v93 float64 = float64(0.1)
                                                                        v94 := v4["a"]
                                                                        _ = v93
                                                                        _ = v94
                                                                    }
                                                                case "":
                                                                    v95 := GenS11{F0: (128 == 2), F1: 319}
                                                                    _ = v95
                                                                default:
                                                                    var v96 int16 = 42
                                                                    _ = v96
                                                                }
                                                                _ = v45
                                                                _ = v46
                                                            } else {
                                                                var v97 int16 = 1000
                                                                var v98 float32 = v44
                                                                _ = v97
                                                                _ = v98
                                                            }
                                                        }
                                                    }
                                                    var v99 int16 = -539
                                                    fmt.Println(v2, fi13, v10)
                                                    _ = v99
                                                } else {
                                                    if true {
                                                        var v100 uint32 = 2147483647
                                                        v101 := GenS3{F0: -999, F1: 42, F2: 1}
                                                        _ = v100
                                                        _ = v101
                                                    }
                                                    v10 += (v10 % ((v10) | 1))
                                                    v102 := []int{}
                                                    _ = v102
                                                }
                                                _ = v43
                                                _ = v44
                                            }
                                        case 2147483647:
                                            v103 := GenS7{F0: 2147483648, F1: true}
                                            _ = v103
                                        }
                                    }
                                case -956:
                                    v104 := []int{(fi13 % ((fi13) | 1))}
                                    _ = v104
                                }
                            default:
                                v2 = v2
                            }
                            v4[""] = (v2 == 1.5)
                            _ = fi14
                        }
                        _ = fi13
                    }
                } else {
                    for fi105 := 0; fi105 < 8; fi105++ {
                        var v106 string = "hello"
                        _ = fi105
                        _ = v106
                    }
                }
                for fi107 := 0; fi107 < 6; fi107++ {
                    var v108 string = "foo bar"
                    if (!true) {
                        for fi109 := 0; fi109 < 18; fi109++ {
                            v1 = v1
                            v1 = (v1 + v1)
                            _ = fi109
                        }
                        var v110 string = (v108 + "foo bar")
                        fmt.Println(v1, v2, v1)
                        _ = v110
                    } else {
                        for fi111 := 0; fi111 < 16; fi111++ {
                            var v112 uint64 = 4294967295
                            v113 := []int{(fi107 / (((fi107 ^ -995)) | 1)), (fi107 % ((fi111) | 1)), (^(fi111))}
                            v114 := GenS3{F0: 0, F1: -43, F2: -102}
                            _ = fi111
                            _ = v112
                            _ = v113
                            _ = v114
                        }
                        switch "foo bar" {
                        case "hello":
                            for fi115 := 0; fi115 < 14; fi115++ {
                                v108 = ((v108 + (((v108 + ("hello" + ("" + v108))) + ((v108 + ("foo bar" + "")) + v108)) + v108)) + v108)
                                v116 := GenS2{F0: 32767}
                                _ = fi115
                                _ = v116
                            }
                        case "a":
                            for fi117 := 0; fi117 < 18; fi117++ {
                                if (!(false && false)) {
                                    fmt.Println(v1)
                                    var v118 string = v108
                                    var v119 int8 = 42
                                    _ = v118
                                    _ = v119
                                }
                                var v120 int16 = -880
                                _ = fi117
                                _ = v120
                            }
                        case "foo bar":
                            v10 = (v10 - 18446744073709551615)
                        default:
                            v2 += (3.14159 / v3)
                        }
                        switch (fi107 >> ((v1) & 63)) {
                        case 31:
                            v121 := []int{(fi107 >> ((v1) & 63)), fi107, fi107}
                            _ = v121
                        case 1:
                            var v122 uint = 4294967295
                            _ = v122
                        case 9223372036854775807:
                            v123 := GenS12{F0: false, F1: 2}
                            _ = v123
                        default:
                            v124 := GenS13{F0: "\n\t", F1: 42}
                            _ = v124
                        }
                    }
                    _ = fi107
                    _ = v108
                }
                v10 -= (^(v10))
            } else {
                if (500 <= 1) {
                    for fi125 := 0; fi125 < 1; fi125++ {
                        var v126 float64 = ((0.0 * v3) + (-2.25 / 1e10))
                        _ = fi125
                        _ = v126
                    }
                    v127 := GenS8{F0: genFn0(65535, 42, 1)}
                    var v128 int16 = -32768
                    _ = v127
                    _ = v128
                } else {
                    switch ("\n\t" + "foo bar") {
                    case "":
                        var v129 int8 = 0
                        _ = v129
                    case "a":
                        v10 += v10
                    case "hello":
                        var v130 int16 = 1
                        _ = v130
                    }
                    for fi131 := 0; fi131 < 6; fi131++ {
                        v4["\n\t"] = (v10 >= (v10 >> ((v1) & 63)))
                        _ = fi131
                    }
                    v2 += 0.1
                }
            }
        }
        v132 := GenS2{F0: -704}
        v133 := GenS6{F0: false, F1: -1}
        _ = v132
        _ = v133
    } else {
        v134 := GenS1{F0: false, F1: 4294967295, F2: -9223372036854775808}
        v3 = ((((((float64((float32(3.14159) / 1.5)) - float64(((0.1 + 1e10) / (0.0 - (3.14159 / 1.0))))) + v2) * (float64(3.14159) * v3)) * ((0.0 * float64(float32(((0.0 / float64((1e10 + -2.25))) + 1.5)))) + v3)) - 0.0) - ((3.14159 * 3.14159) + v2))
        fmt.Println(v10, v10)
        _ = v134
    }
    fmt.Println(v1, v1)
    for fi135 := 0; fi135 < 14; fi135++ {
        v136 := []int{1}
        if (-97 > 76) {
            v1 += 1
            if false {
                var v137 int8 = -128
                v10 = 2
                var v138 string = ((("hello" + "a") + "hello") + "")
                _ = v137
                _ = v138
            }
        } else {
            for fi139 := 0; fi139 < 11; fi139++ {
                if false {
                    if true {
                        fmt.Println(v3, v3, v2)
                        v3 += v3
                        switch "foo bar" {
                        case "":
                            v140 := GenS14{F0: 0}
                            _ = v140
                        case "hello":
                            v10 -= 267
                        }
                    } else {
                        v141 := GenS15{F0: 1000, F1: 2147483647}
                        var v142 uint16 = 42
                        _ = v141
                        _ = v142
                    }
                    var v143 uint = (^(v1))
                    _ = v143
                } else {
                    if len(v136) >= 2 {
                        v136 = v136[1:len(v136)]
                    }
                }
                _ = fi139
            }
        }
        var v144 uint8 = 235
        _ = fi135
        _ = v136
        _ = v144
    }
    switch 4294967295 {
    case -642:
        var v145 bool = true
        _ = v145
    }
    v146 := []int{}
    switch "hello" {
    case "foo bar":
        if (158 <= 101) {
            v10 = v10
            var v147 int = -1
            fmt.Println(v147, v3, v10)
            _ = v147
        }
    default:
        for fi148 := 0; fi148 < 8; fi148++ {
            var v149 uint64 = v10
            _ = fi148
            _ = v149
        }
    }
    for fi150 := 0; fi150 < 13; fi150++ {
        switch ("" + "foo bar") {
        case "a":
            v10 -= (^(v10))
        }
        _ = len(v4)
        _ = fi150
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v146:", len(v146), v146)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v146
    _ = v2
    _ = v3
    _ = v4
}
