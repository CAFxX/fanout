package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool; F1 uint8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64; F1 int64 }
type GenS12 struct { F0 int8; F1 int32 }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 uint; F1 uint8 }
type GenS15 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int8; F1 int64 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 string }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 int64; F1 int8 }
type GenS22 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS23 struct { F0 string; F1 string }
type GenS24 struct { F0 uint8; F1 uint32 }
type GenS25 struct { F0 uint32; F1 uint; F2 string }
type GenS26 struct { F0 int32; F1 uint32 }
type GenS27 struct { F0 uint; F1 int8; F2 uint8 }
type GenS28 struct { F0 uint; F1 int; F2 uint }
type GenS29 struct { F0 uint16; F1 int }

func main() {
    var v1 int16 = -705
    var v2 int64 = 42
    var v3 int8 = 0
    v3 = v3
    if (false && false) {
        if false {
            var v4 uint64 = 179
            for fi5 := 0; fi5 < 2; fi5++ {
                var v6 uint32 = 2
                fmt.Println(v2, v4)
                if ((v3 >> ((1) & 63)) <= (v3 * -1)) {
                    v6 -= 42
                } else {
                    var v7 uint = 628
                    if true {
                        if false {
                            for fi8 := 0; fi8 < 10; fi8++ {
                                v6 = 65535
                                switch (("foo bar" + ("a" + "a")) + "foo bar") {
                                case "hello":
                                    var v9 float32 = -2.25
                                    _ = v9
                                case "":
                                    var v10 bool = (111 == 247)
                                    _ = v10
                                case "a":
                                    if true {
                                        v11 := GenS1{F0: (-(fi5)), F1: 42, F2: -72}
                                        fmt.Println(v1, v3, v7)
                                        if false {
                                            v12 := GenS2{F0: 255, F1: v3}
                                            v13 := GenS3{F0: (v4 * v4), F1: 1}
                                            if ((("" + ("" + "hello")) == ("\n\t" + "hello")) || true) {
                                                v1 += v1
                                                switch fi8 {
                                                case 9223372036854775807:
                                                    switch "hello" {
                                                    case "foo bar":
                                                        switch "a" {
                                                        case "a":
                                                            fmt.Println(v2, fi8, v4)
                                                        case "foo bar":
                                                            v14 := GenS4{F0: (fi5 << ((v7) & 63)), F1: (!(v1 == v1))}
                                                            _ = v14
                                                        default:
                                                            var v15 uint32 = (v6 % (((v6 << ((255) & 63))) | 1))
                                                            _ = v15
                                                        }
                                                    case "":
                                                        var v16 uint = v7
                                                        _ = v16
                                                    case "\n\t":
                                                        if (float64((0.1 + -2.25)) <= 0.1) {
                                                            if false {
                                                                switch "" {
                                                                case "a":
                                                                    v17 := GenS5{F0: v2}
                                                                    _ = v17
                                                                case "foo bar":
                                                                    fmt.Println(v3, fi8, v4)
                                                                default:
                                                                    v18 := GenS4{F0: 128, F1: false}
                                                                    _ = v18
                                                                }
                                                            } else {
                                                                switch ("\n\t" + ("" + "a")) {
                                                                case "hello":
                                                                    if true {
                                                                        v19 := []int{(fi8 | (fi5 * (fi8 >> (((v7 ^ (v7 ^ 65535))) & 63)))), fi8}
                                                                        _ = v19
                                                                    }
                                                                case "a":
                                                                    v20 := GenS3{F0: (v4 / (((v4 << ((v7) & 63))) | 1)), F1: (v4 | v4)}
                                                                    _ = v20
                                                                }
                                                                if false {
                                                                    var v21 uint8 = 14
                                                                    v2 = (v2 * (v2 - v2))
                                                                    fmt.Println(fi8, v2)
                                                                    _ = v21
                                                                } else {
                                                                    if (!("foo bar" <= "a")) {
                                                                        v3 = (v3 >> ((1) & 63))
                                                                        fmt.Println(v6, v4, v3)
                                                                    } else {
                                                                        if ("foo bar" == "\n\t") {
                                                                            fmt.Println(v7, v2, v7)
                                                                            if false {
                                                                                v7 -= (^(v7))
                                                                                var v22 uint16 = 169
                                                                                var v23 int8 = 17
                                                                                _ = v22
                                                                                _ = v23
                                                                            } else {
                                                                                fmt.Println(fi8, v1, v2)
                                                                            }
                                                                            fmt.Println(v3, v6)
                                                                            fmt.Println(v7, v6, v4)
                                                                            v24 := []int{}
                                                                            v24 = append(v24, (^(fi5)))
                                                                            _ = v24
                                                                        } else {
                                                                            var v25 uint32 = (^(v6))
                                                                            var v26 int64 = (v2 & (v2 * (-(v2))))
                                                                            _ = v25
                                                                            _ = v26
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            var v27 float32 = 1e10
                                                            _ = v27
                                                        }
                                                    default:
                                                        switch fi5 {
                                                        case 42:
                                                            switch ((("a" + "foo bar") + "foo bar") + "") {
                                                            case "\n\t":
                                                                v28 := GenS6{F0: ("a" >= ("hello" + "foo bar")), F1: (v6 - 65535)}
                                                                _ = v28
                                                            case "":
                                                                switch ("\n\t" + "foo bar") {
                                                                case "\n\t":
                                                                    fmt.Println(v6, fi8)
                                                                case "foo bar":
                                                                    if true {
                                                                        var v29 float32 = float32(((((0.1 + 1e10) / 1e10) / 3.14159) - 1.5))
                                                                        var v30 string = ""
                                                                        var v31 uint8 = 255
                                                                        _ = v29
                                                                        _ = v30
                                                                        _ = v31
                                                                    } else {
                                                                        v4 = v4
                                                                        if (true && false) {
                                                                            v7 = uint(v6)
                                                                            v32 := []int{int(v7), fi8, (-(fi5))}
                                                                            switch (fi8 * 4294967295) {
                                                                            case 32767:
                                                                                var v33 uint8 = 215
                                                                                var v34 int16 = v1
                                                                                _ = v33
                                                                                _ = v34
                                                                            case -1:
                                                                                fmt.Println(v3)
                                                                            case 127:
                                                                                fmt.Println(v2, v1, v6)
                                                                            }
                                                                            _ = v32
                                                                        } else {
                                                                            var v35 string = "\n\t"
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                if (((!true) || ((-2.25 * (0.1 - (3.14159 * (-2.25 + (0.1 - (((1e10 / 1.0) - -2.25) / 1.5)))))) <= 3.14159)) && true) {
                                                                                    var v36 int = (fi5 | fi5)
                                                                                    v37 := GenS2{F0: 35, F1: (v3 << ((0) & 63))}
                                                                                    if (false && ((v4 % ((v4) | 1)) < 18446744073709551615)) {
                                                                                        var v38 float32 = -2.25
                                                                                        switch (v35 + (v35 + v35)) {
                                                                                        case "":
                                                                                            if (((v2 > (^(v2))) || true) && (!(false && ((v6 >> (((^(v7))) & 63)) > v6)))) {
                                                                                                if false {
                                                                                                    var v39 int16 = v1
                                                                                                    var v40 int = (v36 << ((65535) & 63))
                                                                                                    _ = v39
                                                                                                    _ = v40
                                                                                                } else {
                                                                                                    v38 = v38
                                                                                                }
                                                                                                if ((!(v6 >= uint32(v7))) || false) {
                                                                                                    var v41 int = v36
                                                                                                    fmt.Println(v1)
                                                                                                    var v42 string = v35
                                                                                                    _ = v41
                                                                                                    _ = v42
                                                                                                } else {
                                                                                                    v43 := GenS7{F0: 255, F1: 0, F2: (v6 & v6)}
                                                                                                    fmt.Println(v6)
                                                                                                    _ = v43
                                                                                                }
                                                                                                v38 += v38
                                                                                            }
                                                                                        }
                                                                                        v3 += (v3 >> (((v7 + (v7 << (((v7 + (v7 * v7))) & 63)))) & 63))
                                                                                        _ = v38
                                                                                    }
                                                                                    v2 = (v2 * (v2 | int64(v7)))
                                                                                    v44 := GenS1{F0: fi8, F1: 2, F2: v3}
                                                                                    v45 := []int{32767, 4294967295, fi8}
                                                                                    if false {
                                                                                        var v46 int8 = (v3 << (((^(v7))) & 63))
                                                                                        fmt.Println(v4, v2, v7)
                                                                                        var v47 uint32 = (v6 * 955)
                                                                                        _ = v46
                                                                                        _ = v47
                                                                                    } else {
                                                                                        switch 192 {
                                                                                        case 9223372036854775807:
                                                                                            switch (fi5 - (-(v36))) {
                                                                                            case 42:
                                                                                                switch v36 {
                                                                                                case 378:
                                                                                                    v48 := GenS8{F0: v1}
                                                                                                    _ = v48
                                                                                                default:
                                                                                                    fmt.Println(v6, v4, v1)
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                var v49 string = v35
                                                                                                _ = v49
                                                                                            }
                                                                                        default:
                                                                                            if (0.1 > 0.1) {
                                                                                                v50 := GenS5{F0: v2}
                                                                                                _ = v50
                                                                                            }
                                                                                        }
                                                                                        switch (v35 + v35) {
                                                                                        case "hello":
                                                                                            fmt.Println(v2, v2)
                                                                                        }
                                                                                    }
                                                                                    _ = v36
                                                                                    _ = v37
                                                                                    _ = v44
                                                                                    _ = v45
                                                                                }
                                                                            }
                                                                            v7 = v7
                                                                            v4 += (^(v4))
                                                                            _ = v35
                                                                        }
                                                                        var v51 int32 = 1
                                                                        v6 += v6
                                                                        v4 = (v4 * v4)
                                                                        v4 = (^(v4))
                                                                        var v52 uint16 = 1
                                                                        _ = v51
                                                                        _ = v52
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            if false {
                                                                if true {
                                                                    var v53 uint32 = 65535
                                                                    var v54 uint8 = 1
                                                                    _ = v53
                                                                    _ = v54
                                                                } else {
                                                                    var v55 int32 = 926
                                                                    v55 = (^(v55))
                                                                    _ = v55
                                                                }
                                                                v56 := GenS9{F0: false, F1: 179}
                                                                _ = v56
                                                            }
                                                        }
                                                    }
                                                case 128:
                                                    v57 := GenS8{F0: int16(v2)}
                                                    _ = v57
                                                default:
                                                    v58 := []int{4294967295, fi5}
                                                    _ = v58
                                                }
                                                var v59 int = 2147483647
                                                _ = v59
                                            } else {
                                                v1 = (^(v1))
                                                switch (fi5 % ((fi5) | 1)) {
                                                case 32767:
                                                    v60 := GenS10{F0: 655}
                                                    _ = v60
                                                case -567:
                                                    if (!false) {
                                                        v2 -= v2
                                                        fmt.Println(v1)
                                                        v3 = (v3 % ((v3) | 1))
                                                    }
                                                default:
                                                    v61 := GenS6{F0: ((!(!true)) && true), F1: (v6 - 4294967295)}
                                                    _ = v61
                                                }
                                                v62 := GenS4{F0: fi8, F1: (false && false)}
                                                _ = v62
                                            }
                                            _ = v12
                                            _ = v13
                                        }
                                        v4 = v4
                                        var v63 int8 = v3
                                        _ = v11
                                        _ = v63
                                    } else {
                                        v64 := GenS11{F0: v4, F1: (v2 | v2)}
                                        if ((true && (true && (1 < 128))) && false) {
                                            if false {
                                                if true {
                                                    var v65 string = "a"
                                                    v1 -= (v1 + int16(v3))
                                                    _ = v65
                                                } else {
                                                    switch (("a" + "") + ("foo bar" + ("hello" + "a"))) {
                                                    case "hello":
                                                        var v66 int32 = -105
                                                        _ = v66
                                                    case "":
                                                        v2 += (v2 | 9223372036854775807)
                                                    default:
                                                        switch fi8 {
                                                        case 733:
                                                            var v67 uint64 = 265
                                                            _ = v67
                                                        case -9223372036854775808:
                                                            var v68 float32 = 3.14159
                                                            v1 = v1
                                                            _ = v68
                                                        }
                                                    }
                                                    v6 = uint32(v7)
                                                }
                                                switch "" {
                                                case "a":
                                                    var v69 uint32 = (^(v6))
                                                    _ = v69
                                                }
                                                fmt.Println(v4, v1, v2)
                                            } else {
                                                switch "\n\t" {
                                                case "":
                                                    var v70 int32 = 0
                                                    _ = v70
                                                case "hello":
                                                    var v71 int8 = (v3 * (v3 << (((v7 >> (((v7 / ((v7) | 1))) & 63))) & 63)))
                                                    _ = v71
                                                case "\n\t":
                                                    v4 -= (v4 * (^(v4)))
                                                }
                                                if ((v2 >> ((v7) & 63)) != v2) {
                                                    if false {
                                                        var v72 uint8 = 128
                                                        _ = v72
                                                    }
                                                    v73 := []int{-240, fi8}
                                                    _ = v73
                                                }
                                                switch (fi5 - (fi5 - fi8)) {
                                                case 830:
                                                    if false {
                                                        if ((^(fi8)) == -247) {
                                                            v74 := []int{fi8, (fi8 | fi8)}
                                                            switch 2147483647 {
                                                            case 32767:
                                                                v4 -= 775
                                                            }
                                                            var v75 int8 = (^(v3))
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                    }
                                                case 0:
                                                    if (v6 == 2147483647) {
                                                        var v76 int8 = (^(v3))
                                                        _ = v76
                                                    } else {
                                                        v2 -= -1
                                                        if true {
                                                            var v77 int = int(v1)
                                                            _ = v77
                                                        } else {
                                                            v4 = uint64(v6)
                                                            var v78 uint32 = v6
                                                            var v79 int = (fi8 % ((127) | 1))
                                                            if true {
                                                                var v80 int16 = -1
                                                                var v81 uint64 = 2
                                                                if ((!true) && ((127 > 97) && false)) {
                                                                    if (v78 <= v78) {
                                                                        fmt.Println(fi8, v78, v1)
                                                                        var v82 uint32 = (v78 & v6)
                                                                        v83 := []int{(v79 ^ 127), fi8, -9223372036854775808}
                                                                        _ = v82
                                                                        _ = v83
                                                                    } else {
                                                                        var v84 int8 = int8(v80)
                                                                        _ = v84
                                                                    }
                                                                    v4 = v4
                                                                }
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                            switch (fi8 * fi5) {
                                                            case -59:
                                                                var v85 float32 = 3.14159
                                                                _ = v85
                                                            case 1:
                                                                if false {
                                                                    if (84 > (v4 * (^(v4)))) {
                                                                        if (1 < uint16(v7)) {
                                                                            var v86 uint16 = 32767
                                                                            _ = v86
                                                                        } else {
                                                                            if (((v4 % (((v4 / ((v4) | 1))) | 1)) < v4) || (v7 == (^(v7)))) {
                                                                                v87 := GenS9{F0: (!false), F1: 35}
                                                                                switch int(v2) {
                                                                                case -74:
                                                                                    fmt.Println(fi5, fi5, v4)
                                                                                }
                                                                                v88 := GenS12{F0: 127, F1: 2147483647}
                                                                                _ = v87
                                                                                _ = v88
                                                                            } else {
                                                                                v3 = (^(v3))
                                                                                var v89 uint64 = (v4 - (^(v4)))
                                                                                v90 := []int{(-(fi5))}
                                                                                _ = v89
                                                                                _ = v90
                                                                            }
                                                                        }
                                                                        var v91 int32 = -1
                                                                        if ((((((float32((0.1 + 1e10)) - 1e10) + ((((-2.25 / -2.25) + (((((0.0 + 1.5) / (0.1 * 3.14159)) * 3.14159) - 1e10) / 1.0)) * (3.14159 * (float32(3.14159) / (float32((3.14159 * (1e10 * ((0.0 - 0.0) - ((((float64(1e10) - 1e10) - ((3.14159 + 3.14159) / 3.14159)) * 1.5) / 1e10))))) / 0.1)))) - (0.0 - (1e10 / 3.14159)))) + -2.25) + -2.25) - (((1e10 + 0.0) - (0.1 * ((0.1 + 0.1) / 0.1))) + 0.1)) <= 1.5) {
                                                                            if ((v91 ^ v91) == v91) {
                                                                                switch (fi5 & 35) {
                                                                                case -843:
                                                                                    v1 = (v1 | (v1 + v1))
                                                                                default:
                                                                                    v4 -= (v4 >> ((822) & 63))
                                                                                }
                                                                                v92 := GenS13{F0: true, F1: (v2 | int64(v78))}
                                                                                v1 = (^(v1))
                                                                                _ = v92
                                                                            } else {
                                                                                if false {
                                                                                    v93 := GenS14{F0: v7, F1: 42}
                                                                                    var v94 string = "foo bar"
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                }
                                                                                var v95 int16 = 1
                                                                                _ = v95
                                                                            }
                                                                            fmt.Println(fi8)
                                                                            var v96 int32 = (v91 ^ (v91 / ((-804) | 1)))
                                                                            _ = v96
                                                                        } else {
                                                                            v97 := []int{v79}
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                var v98 int = v79
                                                                                _ = v98
                                                                            case "":
                                                                                switch "" {
                                                                                case "hello":
                                                                                    v79 += fi8
                                                                                default:
                                                                                    if false {
                                                                                        v79 = (fi5 ^ (fi8 & fi5))
                                                                                        v99 := GenS3{F0: v4, F1: v4}
                                                                                        if len(v97) >= 2 {
                                                                                            v97 = v97[1:len(v97)]
                                                                                        }
                                                                                        _ = v99
                                                                                    } else {
                                                                                        var v100 int32 = v91
                                                                                        v101 := GenS15{F0: 65535, F1: v1, F2: 65535}
                                                                                        var v102 bool = (v4 > v4)
                                                                                        _ = v100
                                                                                        _ = v101
                                                                                        _ = v102
                                                                                    }
                                                                                }
                                                                            case "\n\t":
                                                                                fmt.Println(v6)
                                                                            default:
                                                                                if true {
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        var v103 uint = 0
                                                                                        _ = v103
                                                                                    case "hello":
                                                                                        fmt.Println(fi8, v3, v79)
                                                                                    default:
                                                                                        if false {
                                                                                            if true {
                                                                                                v7 = 433
                                                                                                var v104 int = v79
                                                                                                v97 = append(v97, -372)
                                                                                                if false {
                                                                                                    v4 -= v4
                                                                                                    var v105 float32 = 3.14159
                                                                                                    _ = v105
                                                                                                }
                                                                                                if len(v97) >= 2 {
                                                                                                    v97 = v97[1:len(v97)]
                                                                                                }
                                                                                                if true {
                                                                                                    switch (fi8 + (v104 % ((128) | 1))) {
                                                                                                    case -9223372036854775808:
                                                                                                        v1 = (v1 / (((v1 + v1)) | 1))
                                                                                                    }
                                                                                                    v106 := GenS16{F0: "\n\t"}
                                                                                                    v107 := GenS17{F0: -128, F1: 761}
                                                                                                    _ = v106
                                                                                                    _ = v107
                                                                                                } else {
                                                                                                    var v108 int = 127
                                                                                                    if (1.5 <= 0.1) {
                                                                                                        fmt.Println(v108, v3)
                                                                                                    }
                                                                                                    _ = v108
                                                                                                }
                                                                                                fmt.Println(v6)
                                                                                                if len(v97) > 0 {
                                                                                                    v97[0] = fi5
                                                                                                }
                                                                                                _ = v104
                                                                                            } else {
                                                                                                fmt.Println(v78)
                                                                                            }
                                                                                            v97 = append(v97, -964)
                                                                                        } else {
                                                                                            v109 := GenS11{F0: v4, F1: (v2 << (((v7 / ((v7) | 1))) & 63))}
                                                                                            v6 += (v6 / ((42) | 1))
                                                                                            _ = len(v97)
                                                                                            _ = v109
                                                                                        }
                                                                                        if false {
                                                                                            v6 -= 2147483648
                                                                                            v110 := GenS18{F0: (v2 % ((v2) | 1))}
                                                                                            if len(v97) >= 2 {
                                                                                                v97 = v97[1:len(v97)]
                                                                                            }
                                                                                            _ = v110
                                                                                        }
                                                                                    }
                                                                                    v3 = v3
                                                                                    if false {
                                                                                        var v111 string = ("foo bar" + "\n\t")
                                                                                        var v112 uint32 = (v78 * (^(v78)))
                                                                                        v112 += v78
                                                                                        _ = v111
                                                                                        _ = v112
                                                                                    } else {
                                                                                        var v113 float64 = -2.25
                                                                                        var v114 int32 = (^(v91))
                                                                                        if (false || false) {
                                                                                            if false {
                                                                                                if (1e10 != 0.1) {
                                                                                                    var v115 uint64 = (v4 ^ 930)
                                                                                                    _ = v115
                                                                                                } else {
                                                                                                    v116 := GenS10{F0: 42}
                                                                                                    fmt.Println(v79)
                                                                                                    _ = len(v97)
                                                                                                    _ = v116
                                                                                                }
                                                                                                if len(v97) > 0 {
                                                                                                    v97[0] = v79
                                                                                                }
                                                                                                v97 = append(v97, -2147483648)
                                                                                            } else {
                                                                                                v91 += 2147483647
                                                                                                var v117 int64 = (v2 % (((^(v2))) | 1))
                                                                                                v6 -= 4294967295
                                                                                                var v118 int8 = (v3 / ((-124) | 1))
                                                                                                _ = v117
                                                                                                _ = v118
                                                                                            }
                                                                                            v119 := GenS12{F0: v3, F1: v114}
                                                                                            v97 = append(v97, fi5)
                                                                                            _ = v119
                                                                                        } else {
                                                                                            if false {
                                                                                                var v120 float32 = 0.0
                                                                                                var v121 bool = true
                                                                                                var v122 int64 = v2
                                                                                                var v123 int = (v79 >> ((18446744073709551615) & 63))
                                                                                                v124 := GenS2{F0: 255, F1: 72}
                                                                                                _ = v120
                                                                                                _ = v121
                                                                                                _ = v122
                                                                                                _ = v123
                                                                                                _ = v124
                                                                                            } else {
                                                                                                if false {
                                                                                                    if (uint8(v4) != 128) {
                                                                                                        if (v3 <= v3) {
                                                                                                            var v125 float64 = v113
                                                                                                            _ = v125
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v126 uint8 = 1
                                                                                                        if false {
                                                                                                            if true {
                                                                                                                var v127 float32 = 3.14159
                                                                                                                v6 = (v78 ^ (v6 * (v6 % ((2147483648) | 1))))
                                                                                                                var v128 int16 = (^(v1))
                                                                                                                _ = v127
                                                                                                                _ = v128
                                                                                                            } else {
                                                                                                                var v129 uint32 = (v6 + (^(v6)))
                                                                                                                _ = v129
                                                                                                            }
                                                                                                        } else {
                                                                                                            v130 := GenS19{F0: "a"}
                                                                                                            if len(v97) > 0 {
                                                                                                                v97[0] = (fi5 & fi8)
                                                                                                            }
                                                                                                            if true {
                                                                                                                if len(v97) >= 2 {
                                                                                                                    v97 = v97[1:len(v97)]
                                                                                                                }
                                                                                                                if (uint8(v7) != uint8(v4)) {
                                                                                                                    if ((!false) || ((v7 | v7) == (v7 + v7))) {
                                                                                                                        v131 := GenS19{F0: "foo bar"}
                                                                                                                        var v132 uint = (^(v7))
                                                                                                                        _ = v131
                                                                                                                        _ = v132
                                                                                                                    } else {
                                                                                                                        v133 := GenS20{F0: v4}
                                                                                                                        v2 -= (v2 - (v2 * v2))
                                                                                                                        var v134 uint16 = 42
                                                                                                                        _ = v133
                                                                                                                        _ = v134
                                                                                                                    }
                                                                                                                    if (0 != 65535) {
                                                                                                                        v97 = append(v97, int(v78))
                                                                                                                        v114 = (v91 + v91)
                                                                                                                        switch "foo bar" {
                                                                                                                        case "a":
                                                                                                                            fmt.Println(v113)
                                                                                                                            fmt.Println(v79, fi5, v6)
                                                                                                                        case "":
                                                                                                                            v4 = 869
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if ((v113 / v113) <= (3.14159 - 0.0)) {
                                                                                                                        var v135 uint = 255
                                                                                                                        if (!((v1 + 1) <= v1)) {
                                                                                                                            v136 := GenS9{F0: false, F1: (v126 ^ (v126 * v126))}
                                                                                                                            fmt.Println(v126, v2, v1)
                                                                                                                            var v137 uint16 = 237
                                                                                                                            _ = v136
                                                                                                                            _ = v137
                                                                                                                        }
                                                                                                                        _ = v135
                                                                                                                    }
                                                                                                                    v2 -= v2
                                                                                                                    v78 += 42
                                                                                                                    if true {
                                                                                                                        v7 += 4294967295
                                                                                                                        v138 := GenS6{F0: false, F1: (^(v6))}
                                                                                                                        var v139 float64 = 3.14159
                                                                                                                        _ = v138
                                                                                                                        _ = v139
                                                                                                                    } else {
                                                                                                                        var v140 bool = true
                                                                                                                        _ = v140
                                                                                                                    }
                                                                                                                    _ = len(v97)
                                                                                                                    if true {
                                                                                                                        switch 1 {
                                                                                                                        case 2147483647:
                                                                                                                            var v141 int16 = int16(v91)
                                                                                                                            _ = v141
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v142 uint32 = (v78 - uint32(v4))
                                                                                                                    _ = v142
                                                                                                                }
                                                                                                                v97 = append(v97, (v79 - (fi5 << ((v7) & 63))))
                                                                                                            } else {
                                                                                                                switch (-(fi5)) {
                                                                                                                case 128:
                                                                                                                    if len(v97) > 0 {
                                                                                                                        v97[0] = (v79 / ((fi8) | 1))
                                                                                                                    }
                                                                                                                case -543:
                                                                                                                    v78 = (v6 * 954)
                                                                                                                default:
                                                                                                                    v91 = v91
                                                                                                                    v143 := GenS12{F0: v3, F1: (v91 / ((v91) | 1))}
                                                                                                                    if len(v97) >= 2 {
                                                                                                                        v97 = v97[1:len(v97)]
                                                                                                                    }
                                                                                                                    _ = len(v97)
                                                                                                                    v91 = 2147483647
                                                                                                                    _ = v143
                                                                                                                }
                                                                                                                var v144 string = "foo bar"
                                                                                                                switch v144 {
                                                                                                                case "a":
                                                                                                                    v114 += (^(v114))
                                                                                                                case "":
                                                                                                                    var v145 uint8 = uint8(v4)
                                                                                                                    _ = v145
                                                                                                                case "\n\t":
                                                                                                                    v146 := GenS19{F0: ((v144 + ("a" + v144)) + v144)}
                                                                                                                    _ = v146
                                                                                                                default:
                                                                                                                    v147 := GenS15{F0: 42, F1: int16(v2), F2: 641}
                                                                                                                    _ = v147
                                                                                                                }
                                                                                                                _ = v144
                                                                                                            }
                                                                                                            _ = v130
                                                                                                        }
                                                                                                        _ = v126
                                                                                                    }
                                                                                                }
                                                                                                var v148 int32 = v114
                                                                                                _ = v148
                                                                                            }
                                                                                        }
                                                                                        _ = v113
                                                                                        _ = v114
                                                                                    }
                                                                                } else {
                                                                                    var v149 int8 = (v3 % ((v3) | 1))
                                                                                    _ = v149
                                                                                }
                                                                            }
                                                                            if ((v4 / ((354) | 1)) < v4) {
                                                                                var v150 bool = true
                                                                                _ = v150
                                                                            }
                                                                            if (128 < 38) {
                                                                                v151 := GenS21{F0: (v2 ^ 0), F1: 127}
                                                                                _ = len(v97)
                                                                                var v152 uint8 = 127
                                                                                _ = v151
                                                                                _ = v152
                                                                            } else {
                                                                                var v153 float32 = 0.1
                                                                                if true {
                                                                                    switch 126 {
                                                                                    case 128:
                                                                                        if true {
                                                                                            v154 := GenS5{F0: (v2 % ((v2) | 1))}
                                                                                            v155 := GenS19{F0: "\n\t"}
                                                                                            _ = len(v97)
                                                                                            if (v4 == v4) {
                                                                                                if (false || false) {
                                                                                                    if len(v97) >= 2 {
                                                                                                        v97 = v97[1:len(v97)]
                                                                                                    }
                                                                                                    v78 = v78
                                                                                                    var v156 int = (fi8 - fi8)
                                                                                                    _ = v156
                                                                                                } else {
                                                                                                    var v157 int16 = 1
                                                                                                    switch 2147483647 {
                                                                                                    case -840:
                                                                                                        switch fi5 {
                                                                                                        case 0:
                                                                                                            if len(v97) >= 2 {
                                                                                                                v97 = v97[1:len(v97)]
                                                                                                            }
                                                                                                        case -487:
                                                                                                            var v158 uint = (v7 & (^(v7)))
                                                                                                            _ = v158
                                                                                                        }
                                                                                                    case 127:
                                                                                                        v7 = v7
                                                                                                    case 32767:
                                                                                                        var v159 uint64 = v4
                                                                                                        _ = v159
                                                                                                    }
                                                                                                    v97 = append(v97, int(v7))
                                                                                                    var v160 int = -602
                                                                                                    fmt.Println(v153, v160)
                                                                                                    if (620 > 32767) {
                                                                                                        var v161 int = 1
                                                                                                        v162 := GenS22{F0: v78, F1: 42, F2: (-(v157))}
                                                                                                        fmt.Println(fi5)
                                                                                                        if (false || true) {
                                                                                                            v7 = 18446744073709551615
                                                                                                            switch "\n\t" {
                                                                                                            case "hello":
                                                                                                                if true {
                                                                                                                    var v163 uint16 = 32767
                                                                                                                    switch "a" {
                                                                                                                    case "":
                                                                                                                        v164 := GenS23{F0: "foo bar", F1: ""}
                                                                                                                        _ = v164
                                                                                                                    case "hello":
                                                                                                                        if ("a" != ((((("\n\t" + "\n\t") + ("hello" + ("a" + "foo bar"))) + "foo bar") + "hello") + "foo bar")) {
                                                                                                                            v79 = v161
                                                                                                                            v161 = 4294967295
                                                                                                                            var v165 uint32 = (^(v6))
                                                                                                                            _ = v165
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v163
                                                                                                                } else {
                                                                                                                    v2 = (v2 & int64(v4))
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                if false {
                                                                                                                    if len(v97) > 0 {
                                                                                                                        v97[0] = (fi8 - (fi5 ^ (fi5 << ((v7) & 63))))
                                                                                                                    }
                                                                                                                    v97 = append(v97, v160)
                                                                                                                    v166 := GenS7{F0: 255, F1: v91, F2: (v6 % (((v78 - v78)) | 1))}
                                                                                                                    _ = v166
                                                                                                                } else {
                                                                                                                    if len(v97) > 0 {
                                                                                                                        v97[0] = (^(fi5))
                                                                                                                    }
                                                                                                                    if true {
                                                                                                                        v167 := GenS20{F0: uint64(v7)}
                                                                                                                        if false {
                                                                                                                            if ((!true) || (!((v157 >> ((323) & 63)) != -942))) {
                                                                                                                                v79 += 154
                                                                                                                            } else {
                                                                                                                                if false {
                                                                                                                                    v160 = 127
                                                                                                                                    var v168 uint = 1
                                                                                                                                    _ = v168
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v3 += v3
                                                                                                                        }
                                                                                                                        fmt.Println(v153, v7)
                                                                                                                        _ = v167
                                                                                                                    } else {
                                                                                                                        var v169 bool = false
                                                                                                                        _ = v169
                                                                                                                    }
                                                                                                                    v170 := GenS15{F0: 0, F1: (v157 - v157), F2: 32767}
                                                                                                                    _ = v170
                                                                                                                }
                                                                                                            default:
                                                                                                                v171 := GenS1{F0: (-(fi8)), F1: 105, F2: (v3 * (v3 | v3))}
                                                                                                                _ = v171
                                                                                                            }
                                                                                                            var v172 float32 = (v153 - 3.14159)
                                                                                                            _ = v172
                                                                                                        }
                                                                                                        _ = len(v97)
                                                                                                        _ = v161
                                                                                                        _ = v162
                                                                                                    }
                                                                                                    var v173 int8 = v3
                                                                                                    _ = v157
                                                                                                    _ = v160
                                                                                                    _ = v173
                                                                                                }
                                                                                                fmt.Println(fi8, v91)
                                                                                            } else {
                                                                                                v2 += (^(v2))
                                                                                                if len(v97) >= 2 {
                                                                                                    v97 = v97[1:len(v97)]
                                                                                                }
                                                                                            }
                                                                                            _ = v154
                                                                                            _ = v155
                                                                                        }
                                                                                    case 1:
                                                                                        if false {
                                                                                            v174 := GenS24{F0: 1, F1: v78}
                                                                                            _ = len(v97)
                                                                                            if len(v97) >= 2 {
                                                                                                v97 = v97[1:len(v97)]
                                                                                            }
                                                                                            _ = v174
                                                                                        } else {
                                                                                            v153 = (0.1 - v153)
                                                                                            var v175 uint = (v7 / ((1) | 1))
                                                                                            _ = v175
                                                                                        }
                                                                                    default:
                                                                                        var v176 uint = (v7 * v7)
                                                                                        _ = v176
                                                                                    }
                                                                                } else {
                                                                                    v177 := GenS18{F0: v2}
                                                                                    _ = v177
                                                                                }
                                                                                fmt.Println(v4, v79, fi8)
                                                                                _ = v153
                                                                            }
                                                                            if true {
                                                                                _ = len(v97)
                                                                                v6 -= v6
                                                                            } else {
                                                                                v79 = (-(fi5))
                                                                                var v178 int64 = (v2 | v2)
                                                                                var v179 uint8 = 255
                                                                                _ = v178
                                                                                _ = v179
                                                                            }
                                                                            _ = v97
                                                                        }
                                                                        _ = v91
                                                                    } else {
                                                                        v79 = -502
                                                                    }
                                                                    var v180 int16 = v1
                                                                    _ = v180
                                                                } else {
                                                                    var v181 int64 = 260
                                                                    _ = v181
                                                                }
                                                            case -519:
                                                                v182 := []int{fi8, v79, (v79 / ((-9223372036854775808) | 1)), 128}
                                                                _ = v182
                                                            }
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                    }
                                                default:
                                                    if true {
                                                        var v183 int32 = 42
                                                        v184 := []int{fi8, (-(fi5)), fi8, (fi5 | int(v7))}
                                                        var v185 int8 = 42
                                                        _ = v183
                                                        _ = v184
                                                        _ = v185
                                                    } else {
                                                        v3 -= (v3 << ((v7) & 63))
                                                        v2 -= v2
                                                        v186 := []int{(fi8 - 32768), fi8}
                                                        _ = v186
                                                    }
                                                }
                                            }
                                            fmt.Println(v4, v3, fi5)
                                        }
                                        v187 := []int{}
                                        _ = v187
                                        _ = v64
                                    }
                                }
                                v188 := GenS25{F0: v6, F1: 0, F2: "\n\t"}
                                _ = fi8
                                _ = v188
                            }
                        }
                        var v189 int32 = 406
                        _ = v189
                    } else {
                        for fi190 := 0; fi190 < 16; fi190++ {
                            if true {
                                v3 += v3
                                v191 := []int{(fi5 * (fi5 / ((-1) | 1))), 1, (-(fi190)), fi190}
                                if ((!(1 >= -1)) && true) {
                                    var v192 uint32 = 4294967295
                                    _ = v192
                                } else {
                                    if (-1 > 513) {
                                        var v193 int8 = -5
                                        var v194 float32 = (0.0 / 1.0)
                                        if len(v191) > 0 {
                                            v191[0] = (fi190 >> (((v7 * 1)) & 63))
                                        }
                                        _ = v193
                                        _ = v194
                                    }
                                }
                                _ = v191
                            } else {
                                v2 = (v2 + (v2 % (((v2 | (v2 / ((v2) | 1)))) | 1)))
                                switch "foo bar" {
                                case "":
                                    v7 = v7
                                case "\n\t":
                                    v195 := []int{}
                                    _ = v195
                                case "foo bar":
                                    var v196 int64 = (v2 + (v2 ^ (v2 - 0)))
                                    _ = v196
                                default:
                                    var v197 string = ((("a" + ("" + "")) + "foo bar") + "\n\t")
                                    _ = v197
                                }
                                v198 := GenS26{F0: 0, F1: v6}
                                _ = v198
                            }
                            _ = fi190
                        }
                        var v199 int16 = v1
                        switch (fi5 % (((fi5 - (fi5 / ((128) | 1)))) | 1)) {
                        case 32768:
                            fmt.Println(fi5, v3)
                        case -1:
                            var v200 string = "hello"
                            _ = v200
                        }
                        _ = v199
                    }
                    _ = v7
                }
                _ = fi5
                _ = v6
            }
            var v201 bool = ((v3 % ((v3) | 1)) != v3)
            _ = v201
            _ = v4
        } else {
            if false {
                var v202 string = (("a" + "\n\t") + "a")
                if (!false) {
                    v203 := GenS23{F0: v202, F1: v202}
                    switch (((v202 + v202) + v202) + ("\n\t" + v202)) {
                    case "a":
                        for fi204 := 0; fi204 < 1; fi204++ {
                            fmt.Println(v3, v2)
                            switch v202 {
                            case "\n\t":
                                v202 = (v202 + v202)
                            case "":
                                if false {
                                    var v205 uint64 = 18446744073709551615
                                    for fi206 := 0; fi206 < 4; fi206++ {
                                        v207 := []int{}
                                        v208 := GenS3{F0: 282, F1: (v205 & (v205 + v205))}
                                        _ = fi206
                                        _ = v207
                                        _ = v208
                                    }
                                    _ = v205
                                } else {
                                    v1 -= (v1 + (v1 & v1))
                                    switch ("\n\t" + v202) {
                                    case "foo bar":
                                        if true {
                                            var v209 string = (("" + "foo bar") + (v202 + ""))
                                            for fi210 := 0; fi210 < 19; fi210++ {
                                                if (!(int64(v3) < (v2 / ((-1) | 1)))) {
                                                    var v211 int16 = v1
                                                    _ = v211
                                                } else {
                                                    v209 = ("foo bar" + v209)
                                                    fmt.Println(v1)
                                                }
                                                var v212 int32 = 42
                                                v213 := []int{(fi210 & 637)}
                                                if (1 >= 128) {
                                                    var v214 int32 = (v212 / ((42) | 1))
                                                    if len(v213) >= 2 {
                                                        v213 = v213[1:len(v213)]
                                                    }
                                                    v1 += v1
                                                    fmt.Println(v212, fi204)
                                                    v209 = v202
                                                    _ = v214
                                                } else {
                                                    v202 = v202
                                                    v215 := GenS27{F0: 42, F1: v3, F2: 127}
                                                    v2 += (-(v2))
                                                    _ = v215
                                                }
                                                _ = fi210
                                                _ = v212
                                                _ = v213
                                            }
                                            var v216 uint = 0
                                            _ = v209
                                            _ = v216
                                        } else {
                                            v217 := []int{(-(fi204)), -1, -117, fi204}
                                            _ = v217
                                        }
                                    }
                                }
                            case "foo bar":
                                switch -2 {
                                case 4294967295:
                                    for fi218 := 0; fi218 < 2; fi218++ {
                                        v219 := GenS3{F0: 345, F1: 984}
                                        v220 := GenS24{F0: 157, F1: 42}
                                        v221 := []int{int(v3), fi218}
                                        _ = fi218
                                        _ = v219
                                        _ = v220
                                        _ = v221
                                    }
                                case 32768:
                                    fmt.Println(fi204)
                                }
                            }
                            _ = fi204
                        }
                    default:
                        for fi222 := 0; fi222 < 10; fi222++ {
                            v223 := GenS3{F0: 2, F1: 1}
                            for fi224 := 0; fi224 < 14; fi224++ {
                                v202 = "a"
                                _ = fi224
                            }
                            _ = fi222
                            _ = v223
                        }
                    }
                    _ = v203
                }
                v225 := []int{-981, 766}
                _ = v202
                _ = v225
            }
        }
    }
    for fi226 := 0; fi226 < 10; fi226++ {
        v3 -= (v3 >> ((185) & 63))
        _ = fi226
    }
    var v227 int = -1
    if (v227 != v227) {
        for fi228 := 0; fi228 < 16; fi228++ {
            if false {
                v229 := GenS28{F0: 180, F1: (fi228 / ((0) | 1)), F2: 262}
                v230 := []int{}
                _ = v229
                _ = v230
            } else {
                v227 += (fi228 + v227)
                var v231 bool = (!true)
                _ = v231
            }
            v232 := GenS29{F0: 487, F1: int(v2)}
            _ = fi228
            _ = v232
        }
        var v233 int16 = v1
        fmt.Println(v2)
        _ = v233
    } else {
        v1 = v1
        var v234 float64 = 0.1
        _ = v234
    }
    for fi235 := 0; fi235 < 20; fi235++ {
        v236 := []int{v227, (fi235 % (((v227 * (v227 / ((fi235) | 1)))) | 1)), v227, (v227 >> ((995) & 63))}
        _ = fi235
        _ = v236
    }
    v237 := GenS13{F0: true, F1: (^(v2))}
    fmt.Println(v3, v2, v2)
    switch "\n\t" {
    case "":
        var v238 float64 = -2.25
        _ = v238
    case "\n\t":
        switch ("" + "\n\t") {
        case "a":
            var v239 uint = 65535
            _ = v239
        case "\n\t":
            if false {
                switch (^(v227)) {
                case -2147483648:
                    if (v3 == 14) {
                        fmt.Println(v1)
                        v3 = (v3 % ((int8(v1)) | 1))
                    } else {
                        for fi240 := 0; fi240 < 4; fi240++ {
                            v3 -= (v3 ^ v3)
                            for fi241 := 0; fi241 < 17; fi241++ {
                                v227 += 343
                                v242 := []int{(fi240 ^ (fi241 * v227)), fi240, fi240}
                                _ = fi241
                                _ = v242
                            }
                            _ = fi240
                        }
                        v243 := []int{}
                        _ = v243
                    }
                }
                var v244 uint = 58
                v245 := []int{(v227 >> ((255) & 63))}
                _ = v244
                _ = v245
            }
        }
    }
    v246 := GenS25{F0: 65535, F1: 814, F2: "foo bar"}
    var v247 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v227:", v227)
    fmt.Println("v247:", v247)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v227
    _ = v237
    _ = v246
    _ = v247
    _ = v3
}
