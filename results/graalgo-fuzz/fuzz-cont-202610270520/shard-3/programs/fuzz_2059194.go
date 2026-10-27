package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 string; F1 int8; F2 int32 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 int64; F1 string; F2 uint }
type GenS8 struct { F0 int32; F1 int16; F2 uint64 }
func genFn0(p0 int8, p1 uint64, p2 int8) bool {
    v7 := []int{925, 9223372036854775807, -9223372036854775808, 128}
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v7 = append(v7, fi8)
            if (!false) {
                var v10 float64 = 3.14159
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                v11 := GenS1{F0: 188}
                _ = v10
                _ = v11
            }
            if (!(2 != 2147483648)) {
                fmt.Println(p0, p1)
            } else {
                if true {
                    switch fi9 {
                    case 2147483647:
                        switch "hello" {
                        case "hello":
                            v12 := GenS1{F0: 42}
                            _ = v12
                        default:
                            var v13 uint64 = p1
                            _ = v13
                        }
                    }
                }
                v14 := GenS1{F0: 32767}
                _ = v14
            }
            var v15 int8 = 100
            _ = fi9
            _ = v15
        }
        p2 = 0
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (false && true)
}
func genFn1(p0 uint, p1 int64, p2 string) uint64 {
    for fi16 := 0; fi16 < 10; fi16++ {
        p0 += 50
        v17 := GenS1{F0: 65535}
        _ = fi16
        _ = v17
    }
    fmt.Println(p0)
    switch p2 {
    case "":
        v18 := []int{}
        _ = v18
    case "a":
        v19 := []int{-181, 9223372036854775807}
        _ = v19
    case "\n\t":
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"a", "\n\t"}
    v3 := []uint32{534, 2, 0}
    v4 := make(map[int]int)
    var v5 int8 = 55
    var v6 int8 = (v5 | (v5 | v5))
    v20 := GenS1{F0: 255}
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 string = ("" + "\n\t")
        if false {
            v23 := GenS1{F0: 42}
            _ = v23
        } else {
            for fi24 := 0; fi24 < 3; fi24++ {
                switch 32767 {
                case 128:
                    v22 = "foo bar"
                case 127:
                    v25 := GenS1{F0: 14}
                    _ = v25
                }
                v6 = (v6 << ((42) & 63))
                _ = fi24
            }
            if true {
                v5 = (-(v5))
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 uint = 2
                    _ = fi26
                    _ = v27
                }
                switch -603 {
                case 4294967295:
                    fmt.Println(v6, fi21, fi21)
                case -1:
                    switch fi21 {
                    case -616:
                        for fi28 := 0; fi28 < 1; fi28++ {
                            v5 = 112
                            _ = fi28
                        }
                    case 32767:
                        if true {
                            var v29 bool = (110 == 255)
                            for fi30 := 0; fi30 < 16; fi30++ {
                                var v31 int8 = v5
                                _ = fi30
                                _ = v31
                            }
                            _ = v29
                        } else {
                            var v32 int32 = 42
                            _ = v32
                        }
                    default:
                        if ((!(!(!false))) || (462 < 42)) {
                            if ((496 < 18446744073709551615) && ((167 >= 42) || false)) {
                                v2 = append(v2, ((v22 + "\n\t") + ("" + (v22 + v22))))
                            } else {
                                if ((fi21 & (fi21 >> ((18446744073709551615) & 63))) >= fi21) {
                                    for fi33 := 0; fi33 < 8; fi33++ {
                                        fmt.Println(fi21, v6)
                                        _ = fi33
                                    }
                                }
                                if (!(42 > fi21)) {
                                    v34 := v4[(fi21 >> ((1) & 63))]
                                    _ = v34
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = 2147483648
                                    }
                                    for fi35 := 0; fi35 < 2; fi35++ {
                                        switch fi35 {
                                        case 128:
                                            switch fi35 {
                                            case 9223372036854775807:
                                                var v36 float32 = 1e10
                                                _ = v36
                                            case 32767:
                                                fmt.Println(v5, v22, v6)
                                            default:
                                                if false {
                                                    var v37 bool = (false || (!true))
                                                    v22 = "a"
                                                    if v37 {
                                                        v22 = (v22 + (("" + (((((v22 + "hello") + "\n\t") + v22) + v22) + "")) + v22))
                                                    }
                                                    _ = v37
                                                }
                                            }
                                        case 9223372036854775807:
                                            switch 2147483647 {
                                            case -9223372036854775808:
                                                fmt.Println(fi35, fi21, v6)
                                            case 52:
                                                switch v22 {
                                                case "hello":
                                                    if (float64(1.5) != (1e10 + 0.1)) {
                                                        if (42 > 0) {
                                                            switch (fi21 ^ (fi21 + (fi21 >> ((2) & 63)))) {
                                                            case 42:
                                                                v6 -= v5
                                                            case 32768:
                                                                if false {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    v38 := GenS2{F0: 42, F1: 223, F2: 42}
                                                                    _ = v38
                                                                } else {
                                                                    v5 += (v6 & v5)
                                                                    v4[fi35] = -2147483648
                                                                }
                                                            case 4294967295:
                                                                var v39 int32 = -239
                                                                _ = v39
                                                            }
                                                            v6 -= v6
                                                            var v40 float64 = 0.1
                                                            switch fi35 {
                                                            case 32767:
                                                                if (false && (v22 >= "hello")) {
                                                                    if true {
                                                                        v40 += float64(((0.0 + 0.0) * ((1e10 * 1.5) - -2.25)))
                                                                        if false {
                                                                            switch fi35 {
                                                                            case 1:
                                                                                if (!(-1 == 0)) {
                                                                                    v41 := GenS3{F0: fi35}
                                                                                    var v42 uint64 = 2
                                                                                    v5 += (v5 & -1)
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                } else {
                                                                                    var v43 int = (-(fi21))
                                                                                    _ = v43
                                                                                }
                                                                            default:
                                                                                var v44 int32 = 214
                                                                                _ = v44
                                                                            }
                                                                        }
                                                                        if true {
                                                                            var v45 uint = 45
                                                                            v40 = v40
                                                                            v46 := GenS1{F0: 389}
                                                                            _ = v45
                                                                            _ = v46
                                                                        } else {
                                                                            v22 = v22
                                                                            fmt.Println(v40, v5)
                                                                            v5 = v5
                                                                        }
                                                                    } else {
                                                                        var v47 uint64 = 214
                                                                        _ = v47
                                                                    }
                                                                    fmt.Println(fi21)
                                                                    switch v22 {
                                                                    case "foo bar":
                                                                        var v48 bool = (!(354 <= 2))
                                                                        _ = v48
                                                                    default:
                                                                        v49 := v4[(fi21 >> ((133) & 63))]
                                                                        _ = v49
                                                                    }
                                                                } else {
                                                                    var v50 float64 = (-2.25 - v40)
                                                                    _ = v50
                                                                }
                                                            }
                                                            _ = v40
                                                        }
                                                        if (float32(-2.25) <= 0.0) {
                                                            var v51 int8 = (v6 ^ v5)
                                                            switch (v22 + v22) {
                                                            case "\n\t":
                                                                v5 = (v6 | 44)
                                                            case "a":
                                                                var v52 string = (v22 + ("" + (("hello" + "a") + ("\n\t" + v22))))
                                                                _ = v52
                                                            case "foo bar":
                                                                v6 -= 127
                                                            default:
                                                                if false {
                                                                    v53 := GenS1{F0: 389}
                                                                    v6 = (v6 >> ((18446744073709551615) & 63))
                                                                    fmt.Println(v51, v6)
                                                                    _ = v53
                                                                } else {
                                                                    delete(v1, fi21)
                                                                }
                                                            }
                                                            v54 := GenS4{F0: fi21, F1: 174}
                                                            _ = v51
                                                            _ = v54
                                                        }
                                                    }
                                                case "":
                                                    var v55 uint16 = 172
                                                    _ = v55
                                                case "foo bar":
                                                    v56 := GenS5{F0: v22, F1: (v6 & int8(fi35)), F2: 42}
                                                    _ = v56
                                                }
                                            case 2147483647:
                                                switch "foo bar" {
                                                case "a":
                                                    var v57 uint64 = 18446744073709551615
                                                    _ = v57
                                                case "hello":
                                                    var v58 bool = (42 == 1)
                                                    _ = v58
                                                default:
                                                    if (true && false) {
                                                        v59 := v4[fi21]
                                                        v59 = fi35
                                                        v6 = (v6 * v6)
                                                        fmt.Println(v5)
                                                        _ = v59
                                                    }
                                                }
                                            default:
                                                if len(v2) > 0 {
                                                    v2[0] = v22
                                                }
                                                var v60 bool = true
                                                var v61 uint16 = 235
                                                _ = v60
                                                _ = v61
                                            }
                                        default:
                                            var v62 int32 = 0
                                            _ = v62
                                        }
                                        v6 = v6
                                        if false {
                                            v63 := GenS4{F0: (fi35 | fi35), F1: 18446744073709551615}
                                            _ = v63
                                        }
                                        v64 := GenS6{F0: 4294967295, F1: false}
                                        switch v22 {
                                        case "foo bar":
                                            switch v22 {
                                            case "foo bar":
                                                if (0 != 128) {
                                                    var v65 int64 = -9223372036854775808
                                                    var v66 int = (fi21 << ((336) & 63))
                                                    _ = v65
                                                    _ = v66
                                                }
                                            case "a":
                                                var v67 uint = 65535
                                                _ = v67
                                            case "hello":
                                                var v68 uint32 = 683
                                                _ = v68
                                            }
                                        case "":
                                            v69 := GenS1{F0: 255}
                                            _ = v69
                                        case "a":
                                            var v70 uint = 42
                                            _ = v70
                                        }
                                        _ = fi35
                                        _ = v64
                                    }
                                }
                            }
                        } else {
                            v6 = v5
                        }
                    }
                case 1:
                    if (true || (0 >= 842)) {
                        var v71 uint8 = 128
                        _ = v71
                    } else {
                        v72 := GenS4{F0: fi21, F1: 132}
                        _ = v72
                    }
                default:
                    var v73 uint = 255
                    _ = v73
                }
            } else {
                var v74 string = "hello"
                _ = v74
            }
            var v75 float32 = float32(1.5)
            _ = v75
        }
        _ = fi21
        _ = v22
    }
    for fi76 := 0; fi76 < 16; fi76++ {
        v77 := GenS7{F0: 14, F1: ("a" + "\n\t"), F2: 1}
        fmt.Println(fi76)
        var v78 uint16 = 921
        _ = fi76
        _ = v77
        _ = v78
    }
    if (42 != -954) {
        var v79 bool = (!((v5 >> ((79) & 63)) < (-(v5))))
        for fi80 := 0; fi80 < 19; fi80++ {
            var v81 string = "foo bar"
            v6 = v5
            _ = fi80
            _ = v81
        }
        _ = v79
    } else {
        var v82 int16 = -1
        if false {
            var v83 uint32 = 2
            _ = v83
        }
        fmt.Println(v82, v6, v6)
        _ = v82
    }
    if (2147483648 >= 0) {
        var v84 uint = 65535
        _ = v84
    }
    switch ((("\n\t" + ("foo bar" + (("hello" + "hello") + "\n\t"))) + ((("a" + "foo bar") + "") + "foo bar")) + ("foo bar" + "foo bar")) {
    case "foo bar":
        switch "foo bar" {
        case "":
            for fi85 := 0; fi85 < 17; fi85++ {
                var v86 float32 = float32(0.0)
                var v87 int16 = -32768
                _ = fi85
                _ = v86
                _ = v87
            }
        }
    case "hello":
        if (float32(1e10) >= (1e10 - ((1.5 / 1e10) - 1e10))) {
            if true {
                v88 := GenS4{F0: -300, F1: 452}
                _ = v88
            } else {
                _ = len(v2)
                for fi89 := 0; fi89 < 12; fi89++ {
                    if (249 == 563) {
                        fmt.Println(v6, v5)
                    }
                    for fi90 := 0; fi90 < 15; fi90++ {
                        v91 := GenS8{F0: -174, F1: 32767, F2: 4294967295}
                        fmt.Println(fi89)
                        var v92 int = (^(fi89))
                        _ = fi90
                        _ = v91
                        _ = v92
                    }
                    fmt.Println(fi89, v5)
                    _ = fi89
                }
            }
            var v93 float32 = 1e10
            _ = v93
        } else {
            v5 += (v6 | 1)
            delete(v4, 42)
        }
    case "a":
        for fi94 := 0; fi94 < 8; fi94++ {
            if (2 == 32767) {
                v5 = (^(v6))
                var v95 uint32 = 42
                _ = v95
            } else {
                if (true || true) {
                    v6 += (v5 * 42)
                    v3 = append(v3, 824)
                } else {
                    for fi96 := 0; fi96 < 6; fi96++ {
                        var v97 int32 = 0
                        v5 -= v6
                        _ = fi96
                        _ = v97
                    }
                    v5 += (v6 >> ((18446744073709551615) & 63))
                    v5 = -128
                }
                v98 := GenS8{F0: 119, F1: 344, F2: 2}
                _ = v98
            }
            v5 = (v6 * v5)
            _ = fi94
        }
    }
    var v99 bool = false
    for fi100 := 0; fi100 < 3; fi100++ {
        _ = len(v4)
        if (((0.0 * 0.1) > 1.5) || v99) {
            v101 := GenS6{F0: 4294967295, F1: v99}
            _ = v101
        } else {
            fmt.Println(v5, fi100, v6)
            var v102 float64 = float64(1e10)
            delete(v1, fi100)
            _ = v102
        }
        for fi103 := 0; fi103 < 10; fi103++ {
            v5 -= (v5 * (v5 + (v5 >> ((2) & 63))))
            delete(v4, fi103)
            _ = fi103
        }
        _ = fi100
    }
    var v104 uint = 18446744073709551615
    var v105 int = -80
    var v106 uint64 = 4294967295
    var v107 int8 = (v6 % ((-128) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v104
    _ = v105
    _ = v106
    _ = v107
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
