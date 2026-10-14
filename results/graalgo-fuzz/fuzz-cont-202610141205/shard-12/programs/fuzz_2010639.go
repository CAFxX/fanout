package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int8; F2 int16 }
type GenS4 struct { F0 int; F1 int32; F2 string }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 int; F2 int64 }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint8; F1 uint32 }
type GenS13 struct { F0 int16; F1 int8 }
type GenS14 struct { F0 uint8; F1 string; F2 int64 }
func genFn0(p0 uint32, p1 int64, p2 int32) int64 {
    if (!true) {
        var v5 uint16 = 65535
        _ = v5
    }
    if (false || ((2 >= 0) || (!("\n\t" > "hello")))) {
        if false {
            v6 := []int{58, -153}
            p1 = p1
            _ = v6
        } else {
            v7 := GenS1{F0: 128, F1: p0}
            _ = v7
        }
        v8 := []int{128, -205}
        fmt.Println(p1)
        _ = v8
    } else {
        v9 := GenS1{F0: 0, F1: p0}
        fmt.Println(p1)
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int16 = -1
            for fi12 := 0; fi12 < 20; fi12++ {
                p0 -= 42
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    switch "hello" {
    case "":
        var v13 uint16 = 1
        _ = v13
    default:
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 17; fi15++ {
                v16 := []int{(fi15 << ((18446744073709551615) & 63)), (fi15 << ((uint(p0)) & 63)), fi14}
                fmt.Println(fi14, fi15, p1)
                p0 = (p0 << ((255) & 63))
                fmt.Println(p1, fi14)
                var v17 string = ("foo bar" + ("a" + ("" + ("" + ("hello" + "foo bar")))))
                var v18 int64 = p1
                var v19 uint64 = 18446744073709551615
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
                _ = v19
            }
            fmt.Println(p2, fi14)
            _ = fi14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 int16 = 1000
    var v2 uint8 = 10
    var v3 int32 = 0
    var v4 uint32 = 42
    v3 -= v3
    fmt.Println(v3)
    for fi20 := 0; fi20 < 13; fi20++ {
        for fi21 := 0; fi21 < 12; fi21++ {
            fmt.Println(v1, fi20, v1)
            v2 = v2
            v4 = 42
            if true {
                switch (fi20 + 32767) {
                case 1:
                    v22 := GenS2{F0: 913}
                    _ = v22
                case -9223372036854775808:
                    if (true || false) {
                        if false {
                            if (v1 == (v1 >> ((4294967295) & 63))) {
                                fmt.Println(v1, fi21, fi20)
                            }
                            if true {
                                v23 := GenS1{F0: fi20, F1: (v4 >> ((108) & 63))}
                                _ = v23
                            } else {
                                v2 -= (v2 * (v2 ^ (v2 >> ((255) & 63))))
                                fmt.Println(fi20, v3)
                                if false {
                                    v24 := []int{9223372036854775807}
                                    _ = v24
                                }
                                var v25 uint16 = 42
                                _ = v25
                            }
                        }
                    }
                    var v26 float64 = -2.25
                    _ = v26
                case 4294967295:
                    if false {
                        v27 := GenS1{F0: (fi20 ^ (fi20 - (-(fi20)))), F1: v4}
                        if true {
                            var v28 uint = 0
                            if ((-2.25 - (1.5 + 0.1)) == ((3.14159 * -2.25) - (1e10 / float64(1.5)))) {
                                var v29 bool = false
                                if v29 {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v29 = (!(!true))
                                        v30 := []int{fi20}
                                        if v29 {
                                            _ = len(v30)
                                            v30 = append(v30, (fi20 << ((v28) & 63)))
                                            v3 = 495
                                            v31 := GenS3{F0: 1, F1: 68, F2: (v1 & v1)}
                                            _ = v31
                                        } else {
                                            if v29 {
                                                v32 := GenS4{F0: 127, F1: 1, F2: "foo bar"}
                                                _ = v32
                                            } else {
                                                var v33 bool = true
                                                var v34 float64 = (1.5 * float64((1e10 / 3.14159)))
                                                if (v3 > -1) {
                                                    v34 = float64(float32((3.14159 / v34)))
                                                }
                                                var v35 uint64 = 4294967295
                                                v3 = v3
                                                if false {
                                                    v29 = (uint64(v4) < 2)
                                                    if v33 {
                                                        fmt.Println(v28, v4, v29)
                                                        v30 = append(v30, (^(fi21)))
                                                        _ = len(v30)
                                                    }
                                                }
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            }
                                        }
                                        _ = v30
                                    }
                                    v36 := GenS4{F0: 0, F1: v3, F2: "hello"}
                                    _ = v36
                                } else {
                                    switch fi21 {
                                    case 128:
                                        v37 := []int{fi20, 32768, (^(fi21)), fi21}
                                        _ = v37
                                    case 127:
                                        var v38 uint = v28
                                        _ = v38
                                    default:
                                        v39 := GenS4{F0: (fi21 ^ fi21), F1: v3, F2: "\n\t"}
                                        _ = v39
                                    }
                                }
                                var v40 string = ""
                                _ = v29
                                _ = v40
                            } else {
                                fmt.Println(v28)
                                v28 = (v28 & v28)
                            }
                            v4 = (v4 + (^(v4)))
                            _ = v28
                        } else {
                            switch "hello" {
                            case "":
                                v41 := GenS1{F0: (fi20 * fi21), F1: (v4 | (v4 & v4))}
                                _ = v41
                            default:
                                v42 := GenS5{F0: 205, F1: v1}
                                v3 = (v3 + -1)
                                v3 += (-(v3))
                                _ = v42
                            }
                            fmt.Println(v3, fi21)
                        }
                        var v43 int64 = int64(v2)
                        _ = v27
                        _ = v43
                    }
                default:
                    if true {
                        if false {
                            switch (fi21 * 32768) {
                            case 4294967295:
                                v1 -= (v1 % ((v1) | 1))
                            case -154:
                                v44 := []int{fi21, 2147483647}
                                _ = v44
                            default:
                                v45 := []int{fi20, (fi20 * (fi20 ^ 2147483647)), fi21}
                                _ = v45
                            }
                            v46 := GenS6{F0: v2}
                            _ = v46
                        } else {
                            v47 := GenS1{F0: (-(fi20)), F1: 4294967295}
                            v1 = v1
                            var v48 uint32 = 132
                            _ = v47
                            _ = v48
                        }
                        if false {
                            fmt.Println(fi20)
                            if (!(((1e10 + 0.0) * 1.5) >= (3.14159 - (3.14159 - 0.1)))) {
                                switch "hello" {
                                case "foo bar":
                                    var v49 int32 = (v3 - v3)
                                    _ = v49
                                case "\n\t":
                                    switch "foo bar" {
                                    case "foo bar":
                                        if true {
                                            if false {
                                                switch ("\n\t" + "hello") {
                                                case "hello":
                                                    v50 := GenS7{F0: (v2 + v2), F1: (fi20 / ((32768) | 1)), F2: -254}
                                                    _ = v50
                                                case "":
                                                    v51 := GenS4{F0: (fi21 ^ (fi21 & int(v3))), F1: (v3 * 0), F2: ("\n\t" + "\n\t")}
                                                    _ = v51
                                                }
                                                fmt.Println(v3, fi20, fi20)
                                                fmt.Println(v4, v4)
                                            }
                                            v52 := []int{32767, fi21}
                                            _ = v52
                                        }
                                    case "a":
                                        if false {
                                            if false {
                                                if (true && true) {
                                                    v3 += v3
                                                    if true {
                                                        v53 := GenS8{F0: (fi21 | 4294967295), F1: false}
                                                        _ = v53
                                                    } else {
                                                        v54 := GenS4{F0: 402, F1: v3, F2: ("hello" + "a")}
                                                        if (1.5 < 1e10) {
                                                            if false {
                                                                v3 = v3
                                                                if true {
                                                                    var v55 bool = (false && true)
                                                                    var v56 uint32 = v4
                                                                    _ = v55
                                                                    _ = v56
                                                                } else {
                                                                    if (351 <= 1) {
                                                                        switch fi21 {
                                                                        case 0:
                                                                            var v57 float32 = (((((-2.25 / 0.1) + 1.5) * (1e10 + (3.14159 * 0.0))) * float32((1.5 - -2.25))) + (0.1 / 1.0))
                                                                            _ = v57
                                                                        case 692:
                                                                            v4 -= v4
                                                                        default:
                                                                            fmt.Println(v4, v2)
                                                                        }
                                                                        switch (fi21 << ((0) & 63)) {
                                                                        case 127:
                                                                            v58 := GenS9{F0: (fi20 + int(v3)), F1: (v4 / (((^(v4))) | 1))}
                                                                            _ = v58
                                                                        case 128:
                                                                            if (!true) {
                                                                                var v59 uint32 = v4
                                                                                v60 := []int{(fi21 / (((fi20 % ((32768) | 1))) | 1)), fi21, (fi21 + (-(fi20)))}
                                                                                if true {
                                                                                    v1 -= (v1 ^ 1000)
                                                                                }
                                                                                _ = v59
                                                                                _ = v60
                                                                            }
                                                                        case 32768:
                                                                            fmt.Println(v3, v1)
                                                                        }
                                                                    }
                                                                }
                                                                if false {
                                                                    var v61 uint = 2
                                                                    _ = v61
                                                                } else {
                                                                    switch (("\n\t" + "a") + ((("\n\t" + "") + "") + "hello")) {
                                                                    case "foo bar":
                                                                        v4 -= v4
                                                                    case "":
                                                                        if (65535 <= 905) {
                                                                            v62 := GenS5{F0: v1, F1: v1}
                                                                            v63 := []int{(-(fi21))}
                                                                            var v64 uint8 = 138
                                                                            _ = v62
                                                                            _ = v63
                                                                            _ = v64
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            var v65 float64 = (((((1e10 / 1e10) - (1e10 + 0.0)) * -2.25) + -2.25) - 0.1)
                                                                            if (1e10 <= 0.1) {
                                                                                var v66 bool = (int(v1) != fi20)
                                                                                _ = v66
                                                                            }
                                                                            _ = v65
                                                                        }
                                                                    }
                                                                    switch (fi21 ^ 933) {
                                                                    case 42:
                                                                        v67 := []int{-2147483648}
                                                                        _ = v67
                                                                    case 531:
                                                                        var v68 int64 = -562
                                                                        _ = v68
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v69 := GenS2{F0: fi21}
                                                        _ = v54
                                                        _ = v69
                                                    }
                                                } else {
                                                    var v70 string = "hello"
                                                    var v71 uint64 = 42
                                                    _ = v70
                                                    _ = v71
                                                }
                                            } else {
                                                var v72 int64 = -1
                                                switch ("foo bar" + "hello") {
                                                case "foo bar":
                                                    v73 := GenS3{F0: (v3 + (v3 & v3)), F1: -51, F2: (v1 << ((18446744073709551615) & 63))}
                                                    _ = v73
                                                default:
                                                    if true {
                                                        v74 := GenS6{F0: v2}
                                                        v3 = v3
                                                        v75 := []int{fi20, fi21, (-(fi21))}
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                }
                                                v76 := []int{(-(fi20)), 4294967295}
                                                _ = v72
                                                _ = v76
                                            }
                                            var v77 uint64 = 42
                                            v78 := GenS10{F0: 42, F1: 277}
                                            _ = v77
                                            _ = v78
                                        } else {
                                            if true {
                                                fmt.Println(v1, fi20)
                                                switch 500 {
                                                case 4294967295:
                                                    fmt.Println(fi21, fi21)
                                                case -409:
                                                    var v79 string = ((("\n\t" + "a") + ("" + "")) + "foo bar")
                                                    _ = v79
                                                case -732:
                                                    var v80 int32 = int32(v2)
                                                    _ = v80
                                                }
                                                v2 -= 127
                                            } else {
                                                var v81 uint8 = v2
                                                v82 := GenS6{F0: v2}
                                                if true {
                                                    v83 := GenS9{F0: fi20, F1: 2147483647}
                                                    v84 := GenS11{F0: (v1 << ((251) & 63))}
                                                    v85 := GenS5{F0: v1, F1: 371}
                                                    _ = v83
                                                    _ = v84
                                                    _ = v85
                                                }
                                                var v86 uint8 = v2
                                                _ = v81
                                                _ = v82
                                                _ = v86
                                            }
                                        }
                                    }
                                case "hello":
                                    var v87 float32 = (3.14159 - -2.25)
                                    _ = v87
                                default:
                                    if false {
                                        v88 := []int{}
                                        v89 := GenS4{F0: fi20, F1: -2147483648, F2: "foo bar"}
                                        _ = v88
                                        _ = v89
                                    }
                                }
                                fmt.Println(v3, fi21)
                            }
                            var v90 uint32 = v4
                            if (65535 < 14) {
                                v4 += v4
                                var v91 int32 = (v3 ^ v3)
                                _ = v91
                            } else {
                                v92 := []int{}
                                if ((v1 ^ v1) != (v1 % ((v1) | 1))) {
                                    v90 = (v4 & v4)
                                    v3 = (^(v3))
                                }
                                var v93 int16 = (^(v1))
                                _ = v92
                                _ = v93
                            }
                            switch "a" {
                            case "hello":
                                if (1000 >= v1) {
                                    if (987 == uint(v2)) {
                                        v94 := GenS2{F0: (fi21 % ((fi20) | 1))}
                                        v95 := GenS1{F0: fi20, F1: v4}
                                        _ = v94
                                        _ = v95
                                    }
                                }
                            case "foo bar":
                                switch fi21 {
                                case 128:
                                    v96 := []int{}
                                    _ = v96
                                case -2147483648:
                                    v2 = v2
                                case 0:
                                    if (!true) {
                                        if true {
                                            if false {
                                                if true {
                                                    v2 = 128
                                                }
                                            } else {
                                                v97 := GenS8{F0: 9223372036854775807, F1: false}
                                                v98 := GenS12{F0: v2, F1: 714}
                                                if (0.0 == 1e10) {
                                                    fmt.Println(fi21)
                                                } else {
                                                    var v99 float64 = (-2.25 + 3.14159)
                                                    var v100 int64 = 516
                                                    var v101 uint8 = (v2 % ((42) | 1))
                                                    _ = v100
                                                    _ = v101
                                                    _ = v99
                                                }
                                                if (((1.5 / 1.5) - ((-2.25 / (3.14159 / float32(0.1))) + (3.14159 / 1e10))) >= (0.1 / 1e10)) {
                                                    var v102 uint32 = v4
                                                    v103 := []int{(fi20 / ((fi20) | 1)), fi21, fi21}
                                                    if len(v103) > 0 {
                                                        v103[0] = (^(fi20))
                                                    }
                                                    _ = v102
                                                    _ = v103
                                                }
                                                var v104 int64 = 42
                                                _ = v104
                                                _ = v97
                                                _ = v98
                                            }
                                            v105 := []int{32768, -268, fi21, int(v2)}
                                            fmt.Println(v3, v90, fi20)
                                            switch (("" + "foo bar") + "hello") {
                                            case "foo bar":
                                                fmt.Println(v2, fi20, fi20)
                                            }
                                            _ = v105
                                        }
                                        v106 := []int{}
                                        var v107 int = fi20
                                        v2 = (^(v2))
                                        v108 := GenS13{F0: 1000, F1: 24}
                                        _ = v106
                                        _ = v107
                                        _ = v108
                                    } else {
                                        v4 += (v4 | v4)
                                        v4 = (^(v90))
                                    }
                                default:
                                    var v109 uint16 = 255
                                    _ = v109
                                }
                            case "a":
                                var v110 uint32 = (v4 | (^(v90)))
                                _ = v110
                            }
                            _ = v90
                        }
                    }
                }
            }
            _ = fi21
        }
        var v111 uint16 = 83
        _ = fi20
        _ = v111
    }
    v112 := GenS10{F0: 370, F1: 255}
    v113 := []int{}
    v114 := GenS6{F0: v2}
    for fi115 := 0; fi115 < 17; fi115++ {
        v116 := GenS14{F0: (v2 & 2), F1: "hello", F2: 9223372036854775807}
        _ = fi115
        _ = v116
    }
    for fi117 := 0; fi117 < 19; fi117++ {
        v4 = v4
        v113 = append(v113, (^(fi117)))
        _ = fi117
    }
    v118 := GenS9{F0: 891, F1: 1}
    var v119 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v113:", len(v113), v113)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v112
    _ = v113
    _ = v114
    _ = v118
    _ = v119
    _ = v2
    _ = v3
    _ = v4
}
