package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int16; F2 uint16 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int; F1 int64; F2 int }
type GenS8 struct { F0 int16; F1 string; F2 string }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int8; F1 int16 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64; F1 int64 }
func genFn0() int {
    var v5 float64 = float64(0.1)
    v6 := []int{890, 32768}
    var v7 float64 = -2.25
    _ = v5
    _ = v6
    _ = v7
    return 32768
}
func genFn1(p0 uint, p1 string) uint64 {
    if (581 <= 0) {
        if (42 == (p0 % (((p0 * p0)) | 1))) {
            for fi8 := 0; fi8 < 3; fi8++ {
                for fi9 := 0; fi9 < 17; fi9++ {
                    if false {
                        var v10 int = (^(fi9))
                        _ = v10
                    } else {
                        var v11 int32 = -1
                        _ = v11
                    }
                    _ = fi9
                }
                v12 := []int{fi8, (fi8 << ((0) & 63)), 127, 0}
                for fi13 := 0; fi13 < 5; fi13++ {
                    v14 := GenS1{F0: 100, F1: -1, F2: 255}
                    if len(v12) > 0 {
                        v12[0] = genFn0()
                    }
                    _ = fi13
                    _ = v14
                }
                _ = fi8
                _ = v12
            }
            v15 := GenS2{F0: (p0 * 65535)}
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                if false {
                    p0 += p0
                    v18 := []int{}
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = v18
                }
                _ = fi17
            }
            switch (p1 + "a") {
            case "hello":
                v19 := []int{fi16, fi16, (fi16 | -9223372036854775808), fi16}
                _ = v19
            case "foo bar":
                v20 := GenS2{F0: p0}
                _ = v20
            default:
                var v21 int32 = 1
                _ = v21
            }
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS3{F0: -70, F1: 0, F2: 42}
                p0 -= p0
                _ = fi22
                _ = v23
            }
            _ = fi16
        }
        v24 := GenS1{F0: 100, F1: 1, F2: 128}
        _ = v24
    }
    fmt.Println(p0)
    switch p1 {
    case "foo bar":
        p1 = "\n\t"
    case "a":
        var v25 bool = true
        _ = v25
    default:
        fmt.Println(p1, p0)
    }
    if false {
        fmt.Println(p0, p1)
        var v26 float32 = ((-2.25 / -2.25) * 0.0)
        _ = v26
    }
    var v27 uint = p0
    _ = p0
    _ = p1
    _ = v27
    return 4294967295
}

func main() {
    var v1 string = ""
    var v2 uint = 2
    var v3 int = 2147483647
    v4 := make(map[string]uint32)
    fmt.Println(v3)
    var v28 float32 = (-2.25 - 1e10)
    if false {
        var v29 int32 = 0
        v30 := GenS3{F0: 0, F1: -984, F2: 1}
        v29 = (v29 % ((v29) | 1))
        _ = v29
        _ = v30
    } else {
        v31 := []int{v3}
        _ = v31
    }
    switch v1 {
    case "foo bar":
        var v32 int = (^(v3))
        _ = v32
    default:
        fmt.Println(v1, v3)
    }
    if true {
        if (true && (9223372036854775807 < 1)) {
            for fi33 := 0; fi33 < 9; fi33++ {
                v34 := GenS4{F0: 907, F1: (v1 + (((v1 + ("" + (((v1 + v1) + ("" + v1)) + v1))) + "\n\t") + (v1 + (v1 + v1))))}
                v4["\n\t"] = 65535
                _ = fi33
                _ = v34
            }
        }
        for fi35 := 0; fi35 < 14; fi35++ {
            var v36 uint16 = 42
            for fi37 := 0; fi37 < 11; fi37++ {
                switch v3 {
                case -9223372036854775808:
                    switch (fi37 * (fi35 / (((^(fi37))) | 1))) {
                    case 4294967295:
                        if true {
                            v38 := GenS2{F0: (v2 % ((399) | 1))}
                            v39 := GenS4{F0: 665, F1: v1}
                            var v40 uint = v2
                            _ = v38
                            _ = v39
                            _ = v40
                        } else {
                            fmt.Println(v36, v3, v36)
                            if false {
                                fmt.Println(v3)
                            }
                        }
                    case 2147483647:
                        v41 := []int{(^(fi37)), fi37}
                        switch ((v1 + "hello") + (v1 + (((v1 + v1) + "a") + (v1 + v1)))) {
                        case "":
                            if false {
                                v42 := GenS1{F0: 0, F1: 1, F2: 128}
                                var v43 int64 = -1
                                _ = v42
                                _ = v43
                            }
                        case "foo bar":
                            if (uint32(v36) == 207) {
                                if true {
                                    fmt.Println(v28, v3, v28)
                                    v41 = append(v41, genFn0())
                                } else {
                                    v44 := GenS5{F0: -9223372036854775808, F1: int16(v3)}
                                    if (!true) {
                                        var v45 int8 = int8(v36)
                                        _ = v45
                                    }
                                    switch (v1 + (v1 + v1)) {
                                    case "\n\t":
                                        v3 -= fi35
                                        switch ((v1 + ((v1 + v1) + ("" + v1))) + v1) {
                                        case "a":
                                            v3 += genFn0()
                                        case "foo bar":
                                            v3 -= fi37
                                        }
                                        fmt.Println(v2, v36, fi37)
                                    }
                                    switch ("foo bar" + (v1 + v1)) {
                                    case "":
                                        v28 -= v28
                                    case "\n\t":
                                        v46 := GenS4{F0: -375, F1: "hello"}
                                        _ = v46
                                    default:
                                        var v47 int8 = 127
                                        _ = v47
                                    }
                                    v48 := GenS6{F0: uint(v36)}
                                    switch (fi35 % (((fi37 ^ (v3 % ((fi37) | 1)))) | 1)) {
                                    case -1:
                                        v49 := GenS6{F0: (v2 - v2)}
                                        _ = v49
                                    case -453:
                                        var v50 int8 = 100
                                        _ = v50
                                    case 42:
                                        var v51 int = v3
                                        _ = v51
                                    }
                                    _ = v44
                                    _ = v48
                                }
                                v52 := GenS2{F0: 255}
                                _ = v52
                            } else {
                                switch "foo bar" {
                                case "a":
                                    if (-477 != 0) {
                                        v53 := GenS7{F0: (fi35 >> ((v2) & 63)), F1: -1, F2: (v3 - 0)}
                                        v1 = ((v1 + "hello") + v1)
                                        delete(v4, ((v1 + v1) + ((v1 + (v1 + v1)) + "")))
                                        _ = v53
                                    }
                                case "":
                                    var v54 int32 = 2147483647
                                    _ = v54
                                case "foo bar":
                                    v28 += 0.1
                                default:
                                    v3 = fi37
                                }
                                var v55 uint16 = (v36 >> (((^(v2))) & 63))
                                var v56 int16 = int16(v2)
                                _ = v55
                                _ = v56
                            }
                        case "hello":
                            var v57 int64 = 9223372036854775807
                            _ = v57
                        }
                        _ = v41
                    case 618:
                        v28 += (v28 / (1.5 + (v28 * ((0.1 / (v28 - v28)) / v28))))
                    default:
                        if (1e10 != (3.14159 / 0.1)) {
                            fmt.Println(fi37, v36, v3)
                            v58 := GenS8{F0: 0, F1: (("hello" + v1) + v1), F2: (v1 + v1)}
                            var v59 uint = 632
                            _ = v58
                            _ = v59
                        }
                    }
                case 2147483647:
                    if (v28 >= 0.1) {
                        v1 = "foo bar"
                    } else {
                        var v60 bool = true
                        v4[v1] = 2147483647
                        v36 = (v36 ^ (^(v36)))
                        switch 288 {
                        case 42:
                            switch 128 {
                            case -2147483648:
                                if (3.14159 == 1e10) {
                                    v61 := GenS9{F0: (v1 + "")}
                                    if v60 {
                                        v3 -= (fi35 / (((fi37 * 134)) | 1))
                                        v1 = ("" + v1)
                                    } else {
                                        var v62 uint = v2
                                        v3 = (-(v3))
                                        _ = v62
                                    }
                                    _ = v61
                                } else {
                                    fmt.Println(v1)
                                    v63 := GenS10{F0: 1, F1: 137}
                                    if (386 <= 54) {
                                        fmt.Println(fi37)
                                    } else {
                                        if v60 {
                                            v64 := []int{}
                                            v60 = true
                                            _ = v64
                                        }
                                        if ((v1 + ("a" + ("a" + v1))) > v1) {
                                            var v65 uint = (v2 >> (((v2 & v2)) & 63))
                                            switch -9223372036854775808 {
                                            case 2147483647:
                                                if true {
                                                    switch (v1 + v1) {
                                                    case "a":
                                                        switch (v3 * fi35) {
                                                        case -554:
                                                            v66 := GenS2{F0: v2}
                                                            _ = v66
                                                        default:
                                                            if (v36 > v36) {
                                                                v2 += v65
                                                            } else {
                                                                var v67 int64 = -1
                                                                fmt.Println(fi35, v28)
                                                                _ = v67
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v60 = true
                                                    switch v1 {
                                                    case "hello":
                                                        v68 := GenS11{F0: 957}
                                                        _ = v68
                                                    case "\n\t":
                                                        var v69 string = ""
                                                        _ = v69
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if v60 {
                                                            var v70 int64 = 783
                                                            _ = v70
                                                        } else {
                                                            v71 := GenS12{F0: v2}
                                                            var v72 int8 = 54
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                    case "foo bar":
                                                        if (!v60) {
                                                            if false {
                                                                var v73 float32 = (1.5 * (v28 - v28))
                                                                v74 := GenS7{F0: (^(fi35)), F1: -1, F2: v3}
                                                                switch (((v1 + "foo bar") + (v1 + (v1 + v1))) + v1) {
                                                                case "a":
                                                                    v75 := GenS10{F0: 1, F1: -32768}
                                                                    _ = v75
                                                                case "":
                                                                    v76 := GenS13{F0: 739, F1: 1}
                                                                    _ = v76
                                                                default:
                                                                    switch fi37 {
                                                                    case 4294967295:
                                                                        v3 += 0
                                                                        v77 := v4[""]
                                                                        _ = v77
                                                                    case 992:
                                                                        v1 = (v1 + v1)
                                                                    case 42:
                                                                        var v78 uint16 = v36
                                                                        _ = v78
                                                                    default:
                                                                        switch v1 {
                                                                        case "\n\t":
                                                                            if (v60 && true) {
                                                                                var v79 float64 = 0.1
                                                                                var v80 int64 = 9223372036854775807
                                                                                fmt.Println(v1, v60, v60)
                                                                                v81 := []int{fi37, (^(v3)), fi35}
                                                                                fmt.Println(fi37, v28, v36)
                                                                                var v82 int64 = (v80 + v80)
                                                                                _ = v79
                                                                                _ = v80
                                                                                _ = v81
                                                                                _ = v82
                                                                            }
                                                                        default:
                                                                            if (42 <= 42) {
                                                                                v73 -= v73
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = v73
                                                                _ = v74
                                                            } else {
                                                                var v83 uint = v2
                                                                fmt.Println(fi35, v83, v28)
                                                                if v60 {
                                                                    if (!v60) {
                                                                        var v84 float32 = v28
                                                                        _ = v84
                                                                    } else {
                                                                        switch 127 {
                                                                        case 32768:
                                                                            v65 -= v65
                                                                        }
                                                                        switch ("\n\t" + "hello") {
                                                                        case "hello":
                                                                            var v85 int8 = 127
                                                                            _ = v85
                                                                        case "foo bar":
                                                                            var v86 float64 = ((3.14159 - (1.5 * ((3.14159 / 1.0) - (-2.25 * (3.14159 * (((float64(1e10) / 1e10) + 0.1) - (-2.25 * 0.1))))))) / 1.0)
                                                                            _ = v86
                                                                        case "a":
                                                                            if true {
                                                                                var v87 uint64 = 835
                                                                                var v88 int = (-(fi35))
                                                                                _ = v87
                                                                                _ = v88
                                                                            }
                                                                        default:
                                                                            var v89 float32 = v28
                                                                            _ = v89
                                                                        }
                                                                    }
                                                                }
                                                                _ = v83
                                                            }
                                                            fmt.Println(v60, fi37)
                                                        } else {
                                                            fmt.Println(v65)
                                                        }
                                                    case "a":
                                                        v90 := GenS2{F0: 1}
                                                        _ = v90
                                                    default:
                                                        var v91 uint64 = 1
                                                        switch fi37 {
                                                        case -2147483648:
                                                            var v92 uint32 = 2147483648
                                                            _ = v92
                                                        case 223:
                                                            v93 := []int{}
                                                            _ = v93
                                                        }
                                                        if true {
                                                            var v94 int32 = 42
                                                            _ = v94
                                                        }
                                                        _ = v91
                                                    }
                                                }
                                            case -387:
                                                v95 := GenS1{F0: 127, F1: 42, F2: 42}
                                                _ = v95
                                            }
                                            v96 := GenS13{F0: 9223372036854775807, F1: -975}
                                            var v97 string = (("foo bar" + "a") + "hello")
                                            _ = v65
                                            _ = v96
                                            _ = v97
                                        }
                                        var v98 int32 = 0
                                        var v99 int = 32767
                                        _ = v98
                                        _ = v99
                                    }
                                    _ = v63
                                }
                            }
                        case 352:
                            v2 += (v2 ^ v2)
                        default:
                            delete(v4, "")
                        }
                        _ = v60
                    }
                default:
                    var v100 uint = (v2 - 651)
                    _ = v100
                }
                v3 -= (fi35 >> ((v2) & 63))
                v36 = 496
                _ = fi37
            }
            v3 = genFn0()
            _ = fi35
            _ = v36
        }
        var v101 bool = (false && true)
        _ = v101
    } else {
        var v102 uint32 = 65535
        var v103 int16 = 0
        if false {
            if (365 > 528) {
                switch "a" {
                case "a":
                    if false {
                        switch v3 {
                        case 32768:
                            v103 = (v103 & (v103 & v103))
                        default:
                            var v104 uint32 = 921
                            _ = v104
                        }
                    }
                }
            }
            v1 = v1
        } else {
            v103 -= (v103 & (v103 ^ v103))
        }
        _ = v102
        _ = v103
    }
    switch v1 {
    case "foo bar":
        fmt.Println(v2, v2)
    }
    v3 = v3
    var v105 float64 = 1.5
    var v106 uint64 = 0
    v107 := []int{-1, v3, v3}
    fmt.Println(v106, v106, v1)
    fmt.Println(v105)
    if false {
        if len(v107) > 0 {
            v107[0] = (^(v3))
        }
        for fi108 := 0; fi108 < 3; fi108++ {
            switch (v3 / (((v3 - v3)) | 1)) {
            case -737:
                for fi109 := 0; fi109 < 14; fi109++ {
                    var v110 uint8 = 136
                    _ = fi109
                    _ = v110
                }
            case -1:
                v1 = v1
            }
            _ = fi108
        }
    } else {
        v4[((((v1 + "\n\t") + ((v1 + ("foo bar" + v1)) + ("a" + (("a" + "") + "hello")))) + v1) + v1)] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v107:", len(v107), v107)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v105
    _ = v106
    _ = v107
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
