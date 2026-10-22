package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int64; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint32; F1 int; F2 uint8 }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 int64; F1 uint8; F2 int64 }
type GenS9 struct { F0 int8; F1 uint64 }
type GenS10 struct { F0 int32; F1 uint16; F2 int16 }
type GenS11 struct { F0 int32; F1 uint; F2 int16 }
type GenS12 struct { F0 bool; F1 int }
type GenS13 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint32) uint16 {
    if ((-170 != -47) && false) {
        if (1e10 < (3.14159 - 0.1)) {
            fmt.Println(p0)
            for fi6 := 0; fi6 < 17; fi6++ {
                var v7 uint8 = 128
                _ = fi6
                _ = v7
            }
            v8 := []int{298}
            _ = v8
        } else {
            if false {
                var v9 bool = false
                p0 += p0
                var v10 float64 = -2.25
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 string = (("foo bar" + "") + "a")
                v13 := GenS1{F0: uint8(p0)}
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 float64 = 1e10
            _ = v14
        }
        var v15 int16 = -1
        _ = v15
    }
    v16 := []int{-670, 1}
    switch 9223372036854775807 {
    case 4294967295:
        var v17 float32 = 1e10
        _ = v17
    case 9223372036854775807:
        p0 = p0
    case 427:
        var v18 uint8 = 75
        _ = v18
    }
    _ = p0
    _ = v16
    return 0
}
func genFn1(p0 uint32, p1 string, p2 int) uint16 {
    if true {
        if (4294967295 <= 829) {
            p1 = p1
        }
    }
    switch (p2 - (-(p2))) {
    case 1:
        v19 := []int{}
        _ = v19
    case 127:
        for fi20 := 0; fi20 < 15; fi20++ {
            p0 = p0
            _ = fi20
        }
    default:
        p1 = p1
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 8; fi22++ {
            v23 := GenS2{F0: -117, F1: uint64(p0)}
            if true {
                v24 := GenS3{F0: 255, F1: 798}
                _ = v24
            } else {
                switch p1 {
                case "hello":
                    v25 := GenS4{F0: 42, F1: 42, F2: p0}
                    _ = v25
                default:
                    v26 := GenS4{F0: -285, F1: 42, F2: p0}
                    _ = v26
                }
            }
            if (!false) {
                var v27 int = (fi22 % ((fi22) | 1))
                _ = v27
            }
            _ = fi22
            _ = v23
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            if (1 != -1) {
                fmt.Println(fi28, p1, p2)
                var v29 uint64 = 840
                v30 := GenS2{F0: 1, F1: uint64(p0)}
                _ = v29
                _ = v30
            } else {
                v31 := GenS4{F0: 1, F1: 828, F2: 42}
                p2 = fi21
                v32 := GenS5{F0: 269, F1: ""}
                var v33 uint = 42
                switch (p2 ^ (fi28 << (((v33 - 42)) & 63))) {
                case 897:
                    v33 -= (v33 * v33)
                case 32767:
                    switch p1 {
                    case "foo bar":
                        if (-1 == 127) {
                            var v34 string = p1
                            v35 := GenS6{F0: p0, F1: -1, F2: 42}
                            _ = v34
                            _ = v35
                        } else {
                            fmt.Println(p2, fi28, p0)
                            v33 = v33
                        }
                    default:
                        v36 := []int{fi21}
                        _ = v36
                    }
                }
                _ = v31
                _ = v32
                _ = v33
            }
            p0 -= p0
            v37 := []int{(p2 & fi21), (fi28 ^ fi28), 32768, (fi28 % ((fi21) | 1))}
            if false {
                fmt.Println(p2)
                p1 = (((((p1 + "foo bar") + p1) + (p1 + "foo bar")) + (p1 + "\n\t")) + p1)
            }
            switch "foo bar" {
            case "hello":
                if len(v37) > 0 {
                    v37[0] = 741
                }
            case "foo bar":
                p1 = p1
            default:
                _ = len(v37)
            }
            _ = fi28
            _ = v37
        }
        for fi38 := 0; fi38 < 16; fi38++ {
            if (32768 <= fi21) {
                var v39 uint16 = 32767
                v39 += (v39 * genFn0(p0))
                _ = v39
            } else {
                v40 := []int{}
                v41 := GenS1{F0: 1}
                _ = v40
                _ = v41
            }
            _ = fi38
        }
        _ = fi21
    }
    v42 := GenS4{F0: -837, F1: 631, F2: (^(p0))}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v42
    return 0
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint16 = 32767
    var v3 int16 = -1
    var v4 int64 = -1
    var v5 float32 = float32(((0.0 / ((0.1 * 1e10) * 1e10)) - 0.1))
    v2 = 255
    switch (v1 | 128) {
    case -740:
        switch ("a" + "foo bar") {
        case "":
            switch v1 {
            case -936:
                v2 = genFn1(42, ("" + "hello"), v1)
            case 127:
                switch (v1 >> ((85) & 63)) {
                case 0:
                    if (((v5 + float32(-2.25)) >= v5) && true) {
                        v43 := []int{0, (^(v1)), (v1 << ((2) & 63))}
                        _ = v43
                    }
                default:
                    v44 := GenS6{F0: 722, F1: v1, F2: 1}
                    _ = v44
                }
            default:
                fmt.Println(v1, v5, v2)
            }
        }
    default:
        if (!((((("" + "a") + "foo bar") + "hello") < "a") || (((18446744073709551615 < 1) || true) && true))) {
            for fi45 := 0; fi45 < 14; fi45++ {
                v46 := GenS7{F0: -1, F1: v3}
                v47 := GenS8{F0: -1, F1: 128, F2: v4}
                if (42 > 127) {
                    var v48 int32 = -2147483648
                    switch 833 {
                    case 26:
                        if ((!true) || (2 < 2)) {
                            v49 := GenS9{F0: 1, F1: 0}
                            _ = v49
                        } else {
                            var v50 int32 = (v48 % ((v48) | 1))
                            var v51 float64 = 1e10
                            _ = v50
                            _ = v51
                        }
                    case 1:
                        v48 += v48
                    case 42:
                        fmt.Println(v48, v1, v48)
                    }
                    var v52 int16 = int16(fi45)
                    _ = v48
                    _ = v52
                } else {
                    v1 += fi45
                    fmt.Println(v2, v4)
                    fmt.Println(v2)
                }
                _ = fi45
                _ = v46
                _ = v47
            }
            for fi53 := 0; fi53 < 20; fi53++ {
                for fi54 := 0; fi54 < 20; fi54++ {
                    v5 += ((v5 + (v5 + (((1.5 + (v5 - float32((3.14159 + 3.14159)))) + v5) - (v5 / v5)))) * v5)
                    v55 := []int{fi53, (-(v1)), (v1 | -161), fi53}
                    v3 -= -32768
                    _ = fi54
                    _ = v55
                }
                v56 := GenS4{F0: (v4 % (((v4 | int64(v2))) | 1)), F1: 0, F2: 42}
                v1 += 1
                _ = fi53
                _ = v56
            }
            var v57 uint8 = 42
            _ = v57
        } else {
            if (((1e10 / 1e10) * 0.0) > ((3.14159 - ((1e10 * float64(float32(float64(v5)))) / (1e10 * 0.1))) - 3.14159)) {
                switch 128 {
                case -981:
                    switch "hello" {
                    case "\n\t":
                        if false {
                            v58 := []int{128, (v1 + (v1 - (v1 - v1))), -2147483648}
                            if len(v58) > 0 {
                                v58[0] = (v1 - v1)
                            }
                            _ = v58
                        } else {
                            v5 += ((1.5 + 1e10) - (3.14159 + (1.5 / 1.5)))
                            v5 = (v5 / v5)
                        }
                    case "":
                        switch (("foo bar" + "a") + "hello") {
                        case "\n\t":
                            var v59 float32 = (v5 * v5)
                            _ = v59
                        case "a":
                            if (!true) {
                                v2 = v2
                            }
                        case "":
                            fmt.Println(v3)
                        }
                    }
                case 32767:
                    v60 := GenS7{F0: 42, F1: (v3 >> ((18446744073709551615) & 63))}
                    _ = v60
                }
                v5 -= v5
                switch "" {
                case "hello":
                    var v61 bool = (v2 <= (v2 ^ genFn0(0)))
                    _ = v61
                case "\n\t":
                    v62 := GenS4{F0: (v4 >> ((0) & 63)), F1: 42, F2: 131}
                    _ = v62
                default:
                    v2 = 984
                }
            } else {
                if (-2.25 != (1e10 * 0.1)) {
                    fmt.Println(v1, v3)
                    var v63 uint64 = 0
                    v3 = 1
                    _ = v63
                } else {
                    v1 = v1
                }
                v1 = (v1 % ((v1) | 1))
                v4 += 42
            }
        }
    }
    v64 := GenS10{F0: int32(v2), F1: v2, F2: (-(v3))}
    var v65 int8 = 127
    var v66 float64 = float64(-2.25)
    v67 := GenS11{F0: 2147483647, F1: 18446744073709551615, F2: v3}
    if (!false) {
        v68 := GenS9{F0: v65, F1: 0}
        var v69 uint64 = 42
        _ = v68
        _ = v69
    } else {
        if true {
            fmt.Println(v2)
            var v70 string = "\n\t"
            _ = v70
        } else {
            var v71 int32 = 969
            v71 += v71
            for fi72 := 0; fi72 < 10; fi72++ {
                fmt.Println(v65, v4, fi72)
                var v73 float64 = v66
                v74 := GenS1{F0: 255}
                _ = fi72
                _ = v73
                _ = v74
            }
            _ = v71
        }
        var v75 int32 = 893
        _ = v75
    }
    fmt.Println(v66, v3)
    switch (v1 / ((v1) | 1)) {
    case 381:
        var v76 int64 = v4
        _ = v76
    case 0:
        v3 = v3
    case 2147483647:
        v77 := GenS12{F0: (42 > 307), F1: 42}
        _ = v77
    default:
        v66 = float64(v5)
    }
    if ((-1 >= -2147483648) && true) {
        for fi78 := 0; fi78 < 3; fi78++ {
            v5 = 1e10
            _ = fi78
        }
        var v79 bool = false
        _ = v79
    }
    if (genFn1(794, ((("hello" + "hello") + "hello") + ("a" + ("\n\t" + "\n\t"))), 376) < (v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))) {
        v80 := []int{v1, (v1 | (v1 ^ v1))}
        _ = v80
    }
    v66 = (float64((v5 * v5)) / v66)
    switch (v1 / ((-1) | 1)) {
    case 32767:
        if true {
            for fi81 := 0; fi81 < 5; fi81++ {
                for fi82 := 0; fi82 < 13; fi82++ {
                    if (true || false) {
                        switch (("hello" + ("hello" + "")) + ("" + (("a" + ("foo bar" + ("a" + "foo bar"))) + ""))) {
                        case "hello":
                            switch v1 {
                            case 773:
                                var v83 int8 = v65
                                _ = v83
                            default:
                                var v84 float32 = (1e10 * 0.1)
                                _ = v84
                            }
                        }
                    } else {
                        v85 := GenS13{F0: 433, F1: 575}
                        var v86 uint8 = 42
                        _ = v85
                        _ = v86
                    }
                    var v87 string = (("foo bar" + ("" + ("a" + ""))) + "foo bar")
                    v88 := []int{fi82, int(v3), (fi81 * v1)}
                    _ = fi82
                    _ = v87
                    _ = v88
                }
                v5 -= v5
                v89 := GenS13{F0: 42, F1: 725}
                _ = fi81
                _ = v89
            }
            v66 += 0.0
        } else {
            var v90 float64 = (float64(1.5) * v66)
            var v91 uint = 1
            v92 := []int{(v1 - (v1 + v1)), v1, (v1 << (((v91 / (((v91 % ((18446744073709551615) | 1))) | 1))) & 63))}
            _ = v90
            _ = v91
            _ = v92
        }
    case 127:
        if true {
            if false {
                var v93 int = (-(v1))
                fmt.Println(v5, v3, v3)
                var v94 uint64 = 1
                _ = v93
                _ = v94
            }
        } else {
            v95 := GenS8{F0: 442, F1: 0, F2: v4}
            var v96 uint64 = 1
            _ = v95
            _ = v96
        }
    case -2147483648:
        fmt.Println(v5, v5)
    default:
        if ((v5 + v5) >= ((0.1 + v5) - -2.25)) {
            var v97 uint8 = 127
            _ = v97
        }
    }
    fmt.Println(v5)
    v98 := []int{}
    v98 = append(v98, (v1 & (v1 % (((v1 >> ((71) & 63))) | 1))))
    v66 = (1.5 * float64(1e10))
    for fi99 := 0; fi99 < 12; fi99++ {
        switch 127 {
        case 976:
            for fi100 := 0; fi100 < 9; fi100++ {
                var v101 uint8 = 127
                v101 += (^(v101))
                v102 := GenS9{F0: (v65 >> ((1) & 63)), F1: 783}
                var v103 int32 = int32(v101)
                var v104 uint16 = (v2 / ((v2) | 1))
                _ = fi100
                _ = v101
                _ = v102
                _ = v103
                _ = v104
            }
        default:
            for fi105 := 0; fi105 < 20; fi105++ {
                var v106 float64 = float64(v5)
                if (1 != 67) {
                    if len(v98) > 0 {
                        v98[0] = fi99
                    }
                    fmt.Println(v3, v2)
                } else {
                    v1 = 1
                }
                _ = fi105
                _ = v106
            }
        }
        v107 := GenS1{F0: uint8(v2)}
        switch "a" {
        case "":
            fmt.Println(v65)
        case "\n\t":
            switch "foo bar" {
            case "\n\t":
                var v108 int32 = -1
                _ = v108
            case "foo bar":
                for fi109 := 0; fi109 < 2; fi109++ {
                    _ = len(v98)
                    _ = fi109
                }
            case "a":
                v4 += int64(fi99)
            }
        }
        _ = fi99
        _ = v107
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v98:", len(v98), v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v98
}
