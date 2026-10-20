package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS5 struct { F0 string; F1 uint; F2 int16 }
type GenS6 struct { F0 uint32; F1 string; F2 uint32 }
type GenS7 struct { F0 int; F1 int; F2 bool }
type GenS8 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 bool; F1 uint64 }
func genFn0(p0 int, p1 bool) int8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(p0, p1)
        var v7 uint = 255
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := GenS1{F0: (v7 % ((464) | 1)), F1: 127}
            v10 := []int{}
            var v11 uint8 = 121
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    if p1 {
        v12 := GenS1{F0: 255, F1: 240}
        v13 := []int{p0}
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1() string {
    switch ("hello" + ("foo bar" + "hello")) {
    case "":
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: 255, F1: 255}
            v16 := GenS1{F0: 1, F1: 127}
            _ = fi14
            _ = v15
            _ = v16
        }
    case "a":
        v17 := GenS1{F0: 186, F1: 0}
        _ = v17
    case "hello":
        v18 := GenS2{F0: "a", F1: 127, F2: 127}
        _ = v18
    default:
        var v19 float64 = 0.0
        _ = v19
    }
    if (20 != 208) {
        for fi20 := 0; fi20 < 16; fi20++ {
            v21 := GenS1{F0: 0, F1: 42}
            for fi22 := 0; fi22 < 1; fi22++ {
                if true {
                    v23 := []int{}
                    fmt.Println(fi22)
                    if true {
                        v24 := GenS3{F0: 998}
                        switch fi20 {
                        case 9223372036854775807:
                            var v25 int32 = 238
                            _ = v25
                        case -224:
                            if true {
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                            } else {
                                var v26 int32 = -517
                                v26 -= (-(v26))
                                fmt.Println(fi22, fi20)
                                _ = v26
                            }
                        case 127:
                            var v27 int64 = 9223372036854775807
                            _ = v27
                        default:
                            switch (fi22 ^ fi20) {
                            case -9223372036854775808:
                                v28 := GenS1{F0: 1, F1: 232}
                                _ = v28
                            default:
                                switch fi20 {
                                case 0:
                                    var v29 int32 = -847
                                    _ = v29
                                case 32767:
                                    v30 := GenS4{F0: 2, F1: false, F2: 0}
                                    _ = v30
                                }
                            }
                            if (float64(0.1) > float64(float32((float64(-2.25) - 1e10)))) {
                                v31 := GenS4{F0: 229, F1: (4294967295 >= 307), F2: 2147483648}
                                var v32 uint8 = 1
                                _ = v31
                                _ = v32
                            }
                            v23 = append(v23, (fi20 ^ fi20))
                        }
                        var v33 int8 = 100
                        _ = v24
                        _ = v33
                    } else {
                        var v34 string = "\n\t"
                        fmt.Println(v34, fi22)
                        if (0.0 != 0.1) {
                            fmt.Println(fi20, fi22)
                            if (!true) {
                                if false {
                                    switch (^(fi20)) {
                                    case 909:
                                        if len(v23) >= 2 {
                                            v23 = v23[1:len(v23)]
                                        }
                                    case 127:
                                        if len(v23) > 0 {
                                            v23[0] = fi20
                                        }
                                    case 696:
                                        if (((fi22 ^ 32768) != 127) && ((!(!true)) && false)) {
                                            var v35 int16 = 73
                                            fmt.Println(fi20, fi22)
                                            _ = v35
                                        } else {
                                            if false {
                                                if len(v23) >= 2 {
                                                    v23 = v23[1:len(v23)]
                                                }
                                                v36 := GenS4{F0: 1, F1: (-669 >= 713), F2: 2147483647}
                                                _ = v36
                                            }
                                            v34 = "hello"
                                            fmt.Println(v34, fi20, fi22)
                                            var v37 uint64 = 0
                                            v38 := GenS5{F0: v34, F1: 0, F2: -1}
                                            _ = v37
                                            _ = v38
                                        }
                                    default:
                                        var v39 uint = 65535
                                        _ = v39
                                    }
                                    v40 := GenS6{F0: 1, F1: ("\n\t" + "hello"), F2: 4294967295}
                                    fmt.Println(fi20, v34)
                                    _ = v40
                                } else {
                                    var v41 int16 = -1
                                    switch "hello" {
                                    case "":
                                        v42 := GenS6{F0: 0, F1: v34, F2: 307}
                                        v43 := GenS7{F0: (^(fi20)), F1: (fi20 | 128), F2: true}
                                        _ = v42
                                        _ = v43
                                    default:
                                        if ((643 == 682) || (!(!true))) {
                                            v44 := GenS8{F0: 255, F1: 18446744073709551615, F2: v41}
                                            v45 := GenS6{F0: 1, F1: "a", F2: 788}
                                            _ = v44
                                            _ = v45
                                        } else {
                                            if true {
                                                v41 -= 46
                                                fmt.Println(fi20, v34, v41)
                                                fmt.Println(fi20, fi22, fi20)
                                                switch (fi22 << ((4294967295) & 63)) {
                                                case 42:
                                                    v34 = v34
                                                }
                                            } else {
                                                if (false && true) {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if true {
                                                            if len(v23) >= 2 {
                                                                v23 = v23[1:len(v23)]
                                                            }
                                                            switch "\n\t" {
                                                            case "a":
                                                                v41 += (v41 ^ (v41 << ((18446744073709551615) & 63)))
                                                            case "foo bar":
                                                                fmt.Println(v34, v41, fi22)
                                                            default:
                                                                if (v34 >= "a") {
                                                                    fmt.Println(fi22, fi20, v34)
                                                                    v34 = (v34 + ("foo bar" + (v34 + v34)))
                                                                }
                                                            }
                                                        }
                                                    case "":
                                                        switch v34 {
                                                        case "foo bar":
                                                            v41 = 1
                                                        }
                                                    case "a":
                                                        if false {
                                                            v46 := GenS1{F0: 734, F1: 0}
                                                            v34 = ("foo bar" + ((v34 + ("" + v34)) + v34))
                                                            if ((18446744073709551615 != 57) && true) {
                                                                if false {
                                                                    if (v34 == "hello") {
                                                                        var v47 int8 = genFn0((-(fi20)), true)
                                                                        fmt.Println(v47)
                                                                        switch 9223372036854775807 {
                                                                        case 577:
                                                                            fmt.Println(fi20)
                                                                        case -177:
                                                                            v41 = v41
                                                                        }
                                                                        _ = v47
                                                                    }
                                                                    var v48 float64 = -2.25
                                                                    v49 := GenS9{F0: 255}
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                                v50 := GenS1{F0: 0, F1: 0}
                                                                switch fi22 {
                                                                case 42:
                                                                    if (0 <= 1) {
                                                                        v51 := GenS9{F0: 65535}
                                                                        _ = v51
                                                                    }
                                                                case 276:
                                                                    v34 = "foo bar"
                                                                }
                                                                var v52 uint16 = 42
                                                                _ = v50
                                                                _ = v52
                                                            }
                                                            _ = v46
                                                        }
                                                    default:
                                                        var v53 int64 = 541
                                                        _ = v53
                                                    }
                                                    var v54 string = v34
                                                    v55 := GenS8{F0: 2, F1: 496, F2: -730}
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    var v56 uint8 = 1
                                                    fmt.Println(v34)
                                                    _ = v56
                                                }
                                                if false {
                                                    var v57 bool = ((v34 + v34) < "a")
                                                    var v58 bool = (v34 == v34)
                                                    _ = v57
                                                    _ = v58
                                                } else {
                                                    if len(v23) > 0 {
                                                        v23[0] = (-(fi22))
                                                    }
                                                    fmt.Println(v34, v41)
                                                    var v59 float64 = 1e10
                                                    _ = v59
                                                }
                                                v34 = "foo bar"
                                            }
                                            if false {
                                                var v60 string = v34
                                                v61 := GenS8{F0: 134, F1: 42, F2: v41}
                                                _ = v60
                                                _ = v61
                                            }
                                        }
                                    }
                                    var v62 uint16 = 255
                                    _ = v41
                                    _ = v62
                                }
                                var v63 int8 = 42
                                _ = v63
                            }
                            var v64 int8 = 42
                            v23 = append(v23, fi20)
                            fmt.Println(fi20, v64)
                            _ = v64
                        }
                        _ = v34
                    }
                    _ = v23
                }
                var v65 int16 = 562
                _ = fi22
                _ = v65
            }
            _ = fi20
            _ = v21
        }
        if (!((!true) && false)) {
            var v66 uint8 = 42
            v67 := []int{339}
            if (!(false || (!false))) {
                v67 = append(v67, 0)
                v66 -= v66
                v68 := GenS8{F0: (v66 / (((v66 ^ 128)) | 1)), F1: 897, F2: 723}
                _ = v68
            }
            _ = v66
            _ = v67
        } else {
            v69 := GenS5{F0: (("" + ("hello" + "a")) + ("" + "\n\t")), F1: 0, F2: 925}
            var v70 uint8 = 0
            var v71 int64 = 1
            _ = v69
            _ = v70
            _ = v71
        }
    }
    var v72 string = "foo bar"
    v73 := []int{-779}
    _ = v72
    _ = v73
    return v72
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 uint8 = 1
    v4 := make(map[uint32]bool)
    var v5 uint8 = v3
    var v74 float32 = 3.14159
    v3 -= (v5 / ((v5) | 1))
    if true {
        fmt.Println(v5)
        fmt.Println(v3, v3, v3)
    }
    if ((!("" >= "hello")) && true) {
        v75 := GenS9{F0: 255}
        _ = v75
    } else {
        v3 = (v5 * (v3 - 0))
    }
    fmt.Println(v74, v5)
    var v76 int = 0
    fmt.Println(v74, v5, v74)
    switch (("foo bar" + ("a" + "a")) + ("hello" + "a")) {
    case "a":
        v77 := []int{v76, v76, v76}
        _ = v77
    }
    for fi78 := 0; fi78 < 13; fi78++ {
        _ = len(v2)
        switch "" {
        case "":
            v79 := v1["a"]
            _ = v79
        case "foo bar":
            v80 := []int{fi78, v76, (^(fi78)), (fi78 * (v76 ^ 0))}
            _ = v80
        case "a":
            for fi81 := 0; fi81 < 13; fi81++ {
                switch ((("hello" + "a") + ("foo bar" + "")) + (("hello" + "\n\t") + "\n\t")) {
                case "a":
                    fmt.Println(fi78, fi78)
                case "":
                    if true {
                        switch fi81 {
                        case 32768:
                            v76 = fi81
                            v82 := GenS10{F0: -2147483648}
                            v74 = v74
                            _ = v82
                        default:
                            v83 := GenS8{F0: v5, F1: 525, F2: 1000}
                            _ = v83
                        }
                        v74 = -2.25
                        if (32767 <= 2) {
                            fmt.Println(v76)
                        } else {
                            if (!(2 > 0)) {
                                delete(v2, ("hello" + "\n\t"))
                                fmt.Println(fi78, v76, v76)
                                v5 = v3
                            } else {
                                var v84 string = "a"
                                _ = v84
                            }
                            var v85 uint32 = 790
                            _ = v85
                        }
                    } else {
                        v5 -= 0
                        var v86 uint8 = v5
                        _ = v86
                    }
                case "\n\t":
                    v74 -= (v74 - (1e10 - v74))
                default:
                    v76 = 127
                    v87 := []int{(fi81 & -1)}
                    v88 := GenS2{F0: (("foo bar" + "") + "hello"), F1: 127, F2: 24}
                    _ = v87
                    _ = v88
                }
                v89 := GenS11{F0: (!((4294967295 <= 674) && true)), F1: 2}
                _ = fi81
                _ = v89
            }
        default:
            v90 := []int{48, fi78, 243, 128}
            _ = v90
        }
        _ = fi78
    }
    v3 = v3
    var v91 int32 = -1
    var v92 int8 = -122
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    fmt.Println("v76:", v76)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
    _ = v76
    _ = v91
    _ = v92
}
