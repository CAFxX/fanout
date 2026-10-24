package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 int8; F2 int32 }
type GenS8 struct { F0 bool; F1 int16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint8 }
func genFn0(p0 bool, p1 int64) int {
    p0 = false
    if p0 {
        v5 := GenS1{F0: 2, F1: 127}
        if p0 {
            if (18446744073709551615 >= 4294967295) {
                var v6 string = "\n\t"
                _ = v6
            } else {
                var v7 int16 = 32767
                for fi8 := 0; fi8 < 1; fi8++ {
                    for fi9 := 0; fi9 < 4; fi9++ {
                        if p0 {
                            v10 := GenS2{F0: 65535, F1: int16(fi8), F2: 452}
                            p1 = p1
                            _ = v10
                        } else {
                            v11 := GenS3{F0: fi9}
                            v12 := GenS2{F0: 565, F1: v7, F2: 42}
                            _ = v11
                            _ = v12
                        }
                        switch fi8 {
                        case 32768:
                            if p0 {
                                var v13 uint16 = 1
                                var v14 int8 = 95
                                _ = v13
                                _ = v14
                            } else {
                                switch -98 {
                                case 127:
                                    v15 := []int{}
                                    if ((0.1 + float64(3.14159)) > (-2.25 + ((1e10 + 1e10) / 3.14159))) {
                                        var v16 int64 = (p1 ^ -314)
                                        v17 := GenS3{F0: (fi8 / ((32768) | 1))}
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                        _ = v16
                                        _ = v17
                                    } else {
                                        var v18 float32 = 1e10
                                        var v19 float32 = (v18 - v18)
                                        v19 = v18
                                        _ = v18
                                        _ = v19
                                    }
                                    _ = v15
                                case 0:
                                    var v20 uint8 = 124
                                    _ = v20
                                default:
                                    v7 = 42
                                }
                                v21 := []int{fi8, (fi9 % ((fi8) | 1))}
                                _ = v21
                            }
                        case 2147483647:
                            v22 := GenS4{F0: (p1 << ((255) & 63)), F1: 982}
                            _ = v22
                        default:
                            switch "hello" {
                            case "a":
                                var v23 bool = p0
                                _ = v23
                            case "":
                                fmt.Println(p1, v7)
                            }
                        }
                        p0 = (((1e10 * ((0.0 * 1.5) * 3.14159)) * 0.1) == 3.14159)
                        if true {
                            if (2147483647 == 4294967295) {
                                var v24 bool = p0
                                _ = v24
                            } else {
                                if false {
                                    p0 = (2 > 0)
                                    var v25 uint32 = 1
                                    switch 9223372036854775807 {
                                    case -2147483648:
                                        if false {
                                            v26 := GenS5{F0: 1}
                                            fmt.Println(v25, v25, p0)
                                            p0 = p0
                                            _ = v26
                                        }
                                    case 54:
                                        v27 := []int{(fi9 & (fi9 | 4294967295)), (fi8 << ((2) & 63)), 4294967295}
                                        _ = v27
                                    default:
                                        v7 = (v7 / ((v7) | 1))
                                    }
                                    _ = v25
                                } else {
                                    v7 += v7
                                    p1 -= int64(v7)
                                    if (95 >= -108) {
                                        if (p0 && (!(!((4294967295 < 2) || p0)))) {
                                            var v28 uint8 = 127
                                            _ = v28
                                        } else {
                                            var v29 int16 = (v7 - (v7 / (((v7 % (((v7 & (v7 << ((1) & 63)))) | 1))) | 1)))
                                            v30 := GenS5{F0: -54}
                                            v31 := GenS6{F0: 304}
                                            var v32 float64 = (-2.25 - 0.1)
                                            v33 := []int{fi9, (fi9 % (((fi8 | fi9)) | 1)), fi9, (fi8 % ((-676) | 1))}
                                            p1 = -1
                                            _ = v29
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        }
                                        var v34 int16 = (^(v7))
                                        if p0 {
                                            var v35 uint64 = 0
                                            v36 := []int{-1, fi9, fi9}
                                            _ = v35
                                            _ = v36
                                        }
                                        _ = v34
                                    }
                                }
                                v37 := GenS3{F0: (fi9 * (fi9 * -2147483648))}
                                _ = v37
                            }
                        } else {
                            var v38 int8 = 1
                            _ = v38
                        }
                        _ = fi9
                    }
                    _ = fi8
                }
                _ = v7
            }
            for fi39 := 0; fi39 < 16; fi39++ {
                if false {
                    var v40 int32 = -2147483648
                    if (1 < 2) {
                        v40 -= 1
                        v41 := GenS7{F0: v40, F1: -107, F2: v40}
                        v40 += v40
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 2; fi42++ {
                        v43 := GenS3{F0: (fi42 - fi39)}
                        v40 += (v40 | (v40 << ((42) & 63)))
                        _ = fi42
                        _ = v43
                    }
                    _ = v40
                }
                if ((65535 < 4294967295) && p0) {
                    fmt.Println(fi39, p0, p0)
                    for fi44 := 0; fi44 < 4; fi44++ {
                        var v45 float32 = ((float32(1e10) + (-2.25 / 1e10)) - 0.0)
                        _ = fi44
                        _ = v45
                    }
                } else {
                    for fi46 := 0; fi46 < 7; fi46++ {
                        var v47 int8 = 48
                        var v48 int32 = -2147483648
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    fmt.Println(p0, p1, fi39)
                    v49 := GenS2{F0: 738, F1: 622, F2: 523}
                    _ = v49
                }
                if (((!p0) || false) && (!(45 == 545))) {
                    var v50 uint8 = 0
                    p1 = (p1 & (^(p1)))
                    _ = v50
                } else {
                    v51 := []int{(fi39 * fi39)}
                    v52 := GenS1{F0: 874, F1: 110}
                    fmt.Println(fi39)
                    _ = v51
                    _ = v52
                }
                _ = fi39
            }
            fmt.Println(p1, p1)
        } else {
            p1 += (p1 << ((18446744073709551615) & 63))
            for fi53 := 0; fi53 < 13; fi53++ {
                v54 := []int{}
                _ = fi53
                _ = v54
            }
            v55 := GenS8{F0: ((3.14159 + float32((3.14159 + (0.0 / 3.14159)))) == 0.0), F1: -32768}
            _ = v55
        }
        v56 := []int{-2147483648, -623, 2147483647, 30}
        _ = v5
        _ = v56
    }
    switch "foo bar" {
    case "":
        if (65535 < 351) {
            for fi57 := 0; fi57 < 9; fi57++ {
                switch (fi57 ^ 4294967295) {
                case 42:
                    v58 := GenS9{F0: 65535}
                    _ = v58
                case 1:
                    p1 = 1
                default:
                    var v59 int = (fi57 - -2147483648)
                    _ = v59
                }
                _ = fi57
            }
            fmt.Println(p1, p0)
            for fi60 := 0; fi60 < 7; fi60++ {
                var v61 uint8 = 255
                var v62 uint32 = 2147483647
                _ = fi60
                _ = v61
                _ = v62
            }
        } else {
            v63 := GenS2{F0: 1, F1: 0, F2: 2147483647}
            _ = v63
        }
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    v1 := []string{}
    var v2 bool = false
    var v3 int32 = -835
    var v4 uint8 = 1
    v4 -= (v4 >> ((1) & 63))
    if v2 {
        var v64 int32 = v3
        var v65 uint8 = v4
        var v66 bool = (!(((genFn0(v2, 1) < 128) || (-9223372036854775808 != -828)) || v2))
        _ = v64
        _ = v65
        _ = v66
    } else {
        v67 := GenS9{F0: 2}
        fmt.Println(v2, v3, v3)
        _ = v67
    }
    if v2 {
        var v68 string = "hello"
        var v69 int32 = 42
        if (-32768 == 716) {
            switch 128 {
            case -9223372036854775808:
                var v70 uint64 = 18446744073709551615
                _ = v70
            case 533:
                v2 = ((42 < 23) && (!v2))
            }
            var v71 int32 = 2147483647
            _ = v71
        }
        _ = v68
        _ = v69
    } else {
        if (32767 >= 32767) {
            var v72 string = "foo bar"
            _ = v72
        } else {
            var v73 uint = 650
            v4 = 42
            if (1 <= 2) {
                if (9223372036854775807 > 32768) {
                    var v74 uint = (v73 << (((v73 << ((2) & 63))) & 63))
                    _ = v74
                }
                var v75 float32 = 1e10
                _ = v75
            } else {
                var v76 uint32 = 4294967295
                _ = len(v1)
                _ = v76
            }
            _ = v73
        }
        fmt.Println(v4, v2, v4)
        var v77 uint = 2
        _ = v77
    }
    if v2 {
        var v78 uint16 = 65535
        _ = v78
    } else {
        for fi79 := 0; fi79 < 2; fi79++ {
            fmt.Println(v2, v3)
            v2 = true
            _ = fi79
        }
    }
    v4 = 42
    for fi80 := 0; fi80 < 20; fi80++ {
        if ((!(false && (!v2))) && (false && true)) {
            var v81 string = ("a" + ("a" + "hello"))
            _ = v81
        } else {
            v82 := GenS7{F0: v3, F1: 113, F2: (v3 | v3)}
            v3 = (v3 ^ int32(fi80))
            var v83 uint = 2
            _ = v82
            _ = v83
        }
        switch (fi80 << ((359) & 63)) {
        case 892:
            v84 := GenS4{F0: 1, F1: (v3 & -1)}
            _ = v84
        case -756:
            for fi85 := 0; fi85 < 7; fi85++ {
                v86 := GenS10{F0: (!v2)}
                v2 = (!v2)
                var v87 bool = v2
                var v88 uint32 = 2147483647
                _ = fi85
                _ = v86
                _ = v87
                _ = v88
            }
        case 543:
            v2 = v2
        }
        fmt.Println(v3, fi80, v4)
        _ = fi80
    }
    switch "" {
    case "hello":
        v89 := GenS5{F0: 114}
        _ = v89
    case "a":
        for fi90 := 0; fi90 < 12; fi90++ {
            switch 888 {
            case 1:
                fmt.Println(fi90, v4)
            case -614:
                for fi91 := 0; fi91 < 9; fi91++ {
                    switch (fi90 >> ((950) & 63)) {
                    case -42:
                        var v92 int8 = 100
                        _ = v92
                    case -2147483648:
                        switch ("\n\t" + (("foo bar" + "foo bar") + (("foo bar" + "") + ("" + ("\n\t" + "\n\t"))))) {
                        case "a":
                            if (!((true || false) && (0.0 == -2.25))) {
                                switch ("" + "\n\t") {
                                case "a":
                                    v4 -= (v4 | (^(v4)))
                                case "\n\t":
                                    var v93 uint16 = 42
                                    _ = v93
                                case "hello":
                                    var v94 uint16 = 0
                                    _ = v94
                                }
                            }
                        }
                    case -1:
                        switch (-(fi90)) {
                        case 0:
                            var v95 uint16 = 255
                            _ = v95
                        case 1:
                            var v96 uint8 = v4
                            _ = v96
                        case 190:
                            v97 := GenS7{F0: (v3 - v3), F1: -1, F2: (-(v3))}
                            _ = v97
                        }
                        v98 := GenS9{F0: 2}
                        var v99 bool = (0 > 65535)
                        _ = v98
                        _ = v99
                    default:
                        v100 := GenS11{F0: (v4 | v4)}
                        v3 = (v3 << ((42) & 63))
                        _ = v100
                    }
                    _ = fi91
                }
            default:
                v1 = append(v1, "hello")
            }
            fmt.Println(v2, v2, v3)
            _ = fi90
        }
    }
    var v101 uint = 315
    v4 += (^(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
}
