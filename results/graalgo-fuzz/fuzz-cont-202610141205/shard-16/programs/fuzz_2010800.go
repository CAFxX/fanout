package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint64; F1 bool; F2 string }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 bool; F1 int; F2 string }
type GenS6 struct { F0 int64; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int16, p1 int32) int64 {
    p0 += (p0 >> ((918) & 63))
    v6 := GenS1{F0: (p0 & p0), F1: -2147483648, F2: 0}
    v7 := GenS1{F0: -32768, F1: 42, F2: 2}
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            if false {
                switch -506 {
                case -9223372036854775808:
                    v10 := []int{int(p1)}
                    _ = v10
                default:
                    p0 = p0
                }
                v11 := GenS1{F0: 1, F1: fi8, F2: 4}
                _ = v11
            }
            _ = fi9
        }
        _ = fi8
    }
    var v12 float64 = (((1.5 * 3.14159) + ((float64(1.5) - 1.5) / 3.14159)) - (1.5 / 1.5))
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    _ = v7
    return 42
}
func genFn1() uint8 {
    var v13 string = ("a" + "a")
    fmt.Println(v13)
    fmt.Println(v13)
    v14 := []int{-629}
    _ = v13
    _ = v14
    return 255
}

func main() {
    v1 := []string{"foo bar", "foo bar", "foo bar", "a"}
    v2 := []uint32{337, 2}
    var v3 uint16 = 485
    var v4 bool = ("foo bar" == "hello")
    v5 := []int{952}
    fmt.Println(v3)
    if (912 != -9223372036854775808) {
        var v15 uint64 = 0
        var v16 uint64 = (v15 & v15)
        _ = v15
        _ = v16
    } else {
        fmt.Println(v4, v4)
        var v17 uint16 = 2
        if true {
            var v18 string = "foo bar"
            for fi19 := 0; fi19 < 18; fi19++ {
                v17 = (v3 ^ (v17 + (^(v17))))
                v20 := GenS2{F0: 790, F1: 65535}
                v21 := GenS2{F0: uint32(v17), F1: v3}
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = v18
        }
        _ = v17
    }
    if v4 {
        _ = len(v5)
        var v22 int8 = 127
        _ = v22
    }
    if (!v4) {
        _ = len(v2)
        switch -664 {
        case 128:
            v4 = false
        default:
            v4 = false
        }
        if (!v4) {
            v23 := GenS3{F0: 0, F1: (323 > -614), F2: "a"}
            _ = v23
        } else {
            if v4 {
                if v4 {
                    switch "" {
                    case "hello":
                        for fi24 := 0; fi24 < 8; fi24++ {
                            for fi25 := 0; fi25 < 12; fi25++ {
                                v26 := GenS4{F0: 1, F1: 255}
                                if (!false) {
                                    switch int(v3) {
                                    case 9223372036854775807:
                                        if (true || v4) {
                                            v3 = (v3 ^ (v3 / (((v3 ^ (v3 >> ((65535) & 63)))) | 1)))
                                            switch (("foo bar" + ((("\n\t" + ("foo bar" + ("a" + "a"))) + "hello") + "hello")) + "") {
                                            case "hello":
                                                var v27 int64 = 27
                                                v28 := GenS3{F0: 0, F1: (!((3.14159 < -2.25) && false)), F2: "a"}
                                                v29 := GenS5{F0: ((-2.25 == 3.14159) || true), F1: (-(fi25)), F2: "a"}
                                                v4 = (!v4)
                                                _ = len(v2)
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            case "\n\t":
                                                v3 = v3
                                                switch "a" {
                                                case "a":
                                                    v30 := GenS1{F0: -637, F1: 32768, F2: 334}
                                                    _ = v30
                                                case "\n\t":
                                                    switch fi25 {
                                                    case 127:
                                                        var v31 uint = 847
                                                        _ = v31
                                                    case -745:
                                                        fmt.Println(fi24)
                                                        var v32 int = -2147483648
                                                        fmt.Println(v32, fi24)
                                                        _ = v32
                                                    }
                                                case "hello":
                                                    if true {
                                                        fmt.Println(v4, v3)
                                                        v33 := GenS5{F0: false, F1: fi24, F2: ("\n\t" + "\n\t")}
                                                        _ = v33
                                                    } else {
                                                        var v34 bool = (698 > 18446744073709551615)
                                                        switch "\n\t" {
                                                        case "hello":
                                                            _ = len(v5)
                                                            var v35 int16 = -32768
                                                            _ = v35
                                                        case "foo bar":
                                                            switch "" {
                                                            case "foo bar":
                                                                v36 := GenS4{F0: 42, F1: 0}
                                                                _ = v36
                                                            case "":
                                                                switch fi24 {
                                                                case 42:
                                                                    if v34 {
                                                                        v4 = ((!(v34 || (65535 > (v3 ^ v3)))) || v4)
                                                                        var v37 float32 = (0.1 + 0.1)
                                                                        _ = v37
                                                                    }
                                                                default:
                                                                    v3 += v3
                                                                }
                                                            }
                                                        default:
                                                            v3 -= v3
                                                        }
                                                        var v38 int32 = -1
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            fmt.Println(v38)
                                                        default:
                                                            fmt.Println(v34, v3)
                                                        }
                                                        _ = v34
                                                        _ = v38
                                                    }
                                                default:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if v4 {
                                                        var v39 uint8 = 14
                                                        if v4 {
                                                            var v40 uint64 = 507
                                                            var v41 int = -9223372036854775808
                                                            if false {
                                                                var v42 uint16 = (v3 >> ((255) & 63))
                                                                if (265 > -1) {
                                                                    v43 := GenS6{F0: -1, F1: (^(v42)), F2: 42}
                                                                    var v44 int32 = 2147483647
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                                v45 := GenS3{F0: v40, F1: (!true), F2: "a"}
                                                                _ = v42
                                                                _ = v45
                                                            } else {
                                                                v40 += v40
                                                                var v46 uint16 = (v3 % ((v3) | 1))
                                                                fmt.Println(fi25, v41, v3)
                                                                _ = v46
                                                            }
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        if v4 {
                                                            var v47 uint64 = 4294967295
                                                            if ((v47 | 0) != (v47 ^ (v47 | v47))) {
                                                                if ((v39 | v39) == (v39 - genFn1())) {
                                                                    v48 := GenS7{F0: 128, F1: "\n\t"}
                                                                    switch (fi24 ^ (^(fi25))) {
                                                                    case -823:
                                                                        var v49 uint32 = 722
                                                                        _ = v49
                                                                    case 0:
                                                                        v3 -= (v3 << ((4294967295) & 63))
                                                                    default:
                                                                        if false {
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                switch (fi25 * (-(fi25))) {
                                                                                case -1:
                                                                                    if ((!v4) && true) {
                                                                                        fmt.Println(v47, v47)
                                                                                    } else {
                                                                                        var v50 int8 = 31
                                                                                        var v51 int16 = 1000
                                                                                        if v4 {
                                                                                            if (v4 || (v4 && (v47 < 18446744073709551615))) {
                                                                                                fmt.Println(v47, v51)
                                                                                            } else {
                                                                                                switch "" {
                                                                                                case "hello":
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                case "":
                                                                                                    fmt.Println(v4, v3)
                                                                                                }
                                                                                                var v52 float32 = float32(-2.25)
                                                                                                var v53 int = fi24
                                                                                                if false {
                                                                                                    switch "\n\t" {
                                                                                                    case "":
                                                                                                        if v4 {
                                                                                                            v39 -= genFn1()
                                                                                                            v2 = append(v2, 1)
                                                                                                        } else {
                                                                                                            v52 -= 0.1
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        v3 -= (v3 ^ 2)
                                                                                                    case "foo bar":
                                                                                                        var v54 int16 = v51
                                                                                                        _ = v54
                                                                                                    default:
                                                                                                        fmt.Println(v3)
                                                                                                    }
                                                                                                    var v55 uint = 18446744073709551615
                                                                                                    _ = v55
                                                                                                }
                                                                                                _ = v52
                                                                                                _ = v53
                                                                                            }
                                                                                            if ((v47 & 0) < v47) {
                                                                                                _ = len(v1)
                                                                                                v56 := GenS8{F0: 1}
                                                                                                _ = v56
                                                                                            } else {
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = (^(fi24))
                                                                                                }
                                                                                                v39 += v39
                                                                                                _ = len(v1)
                                                                                                if (585 >= -367) {
                                                                                                    if v4 {
                                                                                                        v47 = v47
                                                                                                        fmt.Println(v4, fi25)
                                                                                                        v57 := GenS4{F0: 538, F1: (v39 ^ v39)}
                                                                                                        switch "" {
                                                                                                        case "":
                                                                                                            v58 := GenS4{F0: 2, F1: v39}
                                                                                                            _ = v58
                                                                                                        }
                                                                                                        _ = v57
                                                                                                    } else {
                                                                                                        if v4 {
                                                                                                            switch fi24 {
                                                                                                            case 216:
                                                                                                                var v59 uint8 = v39
                                                                                                                _ = v59
                                                                                                            case 42:
                                                                                                                v60 := GenS8{F0: (v3 + v3)}
                                                                                                                _ = v60
                                                                                                            default:
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = 65535
                                                                                                                }
                                                                                                            }
                                                                                                            var v61 uint64 = (v47 << ((2) & 63))
                                                                                                            _ = v61
                                                                                                        }
                                                                                                        var v62 int16 = -1
                                                                                                        _ = v62
                                                                                                    }
                                                                                                    if ((!true) || v4) {
                                                                                                        var v63 uint64 = v47
                                                                                                        _ = v63
                                                                                                    }
                                                                                                    var v64 string = "a"
                                                                                                    _ = v64
                                                                                                } else {
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = 1
                                                                                                    }
                                                                                                    var v65 float32 = 3.14159
                                                                                                    v65 = 1.5
                                                                                                    var v66 uint = 407
                                                                                                    v65 = v65
                                                                                                    if false {
                                                                                                        var v67 uint64 = v47
                                                                                                        _ = v67
                                                                                                    } else {
                                                                                                        if (!((v39 / (((^(v39))) | 1)) >= genFn1())) {
                                                                                                            fmt.Println(v65)
                                                                                                            fmt.Println(fi25)
                                                                                                        }
                                                                                                        fmt.Println(v51, v65, v3)
                                                                                                        var v68 int8 = v50
                                                                                                        _ = v68
                                                                                                    }
                                                                                                    _ = v65
                                                                                                    _ = v66
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v50, v4)
                                                                                        }
                                                                                        _ = v50
                                                                                        _ = v51
                                                                                    }
                                                                                case -214:
                                                                                    var v69 int64 = 1
                                                                                    _ = v69
                                                                                }
                                                                            case "hello":
                                                                                v39 -= (v39 * (v39 * (v39 % (((v39 << ((4294967295) & 63))) | 1))))
                                                                            }
                                                                        } else {
                                                                            var v70 int16 = 1000
                                                                            _ = v70
                                                                        }
                                                                    }
                                                                    _ = v48
                                                                }
                                                                fmt.Println(v4, v4)
                                                                switch fi25 {
                                                                case 42:
                                                                    fmt.Println(v39, fi25, fi25)
                                                                default:
                                                                    v47 += 977
                                                                }
                                                            } else {
                                                                v1 = append(v1, ("\n\t" + ""))
                                                                if (v4 && true) {
                                                                    if len(v5) > 0 {
                                                                        v5[0] = fi25
                                                                    }
                                                                    fmt.Println(fi25)
                                                                }
                                                                switch (fi24 / (((fi24 + -701)) | 1)) {
                                                                case 0:
                                                                    v4 = v4
                                                                case -986:
                                                                    if (false && ((((3.14159 * 3.14159) / 1.0) / 1.5) < ((float64(0.0) * 3.14159) * (0.0 / 3.14159)))) {
                                                                        var v71 int16 = 505
                                                                        var v72 float64 = ((1.5 * 0.1) * 1.5)
                                                                        var v73 uint8 = genFn1()
                                                                        _ = v71
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                case 32768:
                                                                    var v74 string = ""
                                                                    _ = v74
                                                                }
                                                                var v75 int = (-(fi25))
                                                                _ = len(v5)
                                                                _ = v75
                                                            }
                                                            _ = v47
                                                        }
                                                        _ = v39
                                                    }
                                                    switch "hello" {
                                                    case "":
                                                        var v76 uint = 2
                                                        _ = v76
                                                    case "hello":
                                                        _ = len(v1)
                                                    default:
                                                        v4 = false
                                                    }
                                                }
                                                switch "foo bar" {
                                                case "a":
                                                    v3 -= 0
                                                case "":
                                                    switch 654 {
                                                    case 9223372036854775807:
                                                        if (3.14159 >= 3.14159) {
                                                            if true {
                                                                v4 = v4
                                                                v4 = (529 != 1)
                                                            } else {
                                                                if (-1 > 9223372036854775807) {
                                                                    v77 := GenS7{F0: 128, F1: "hello"}
                                                                    v78 := GenS4{F0: 65535, F1: 2}
                                                                    _ = v77
                                                                    _ = v78
                                                                } else {
                                                                    fmt.Println(fi24, v3)
                                                                    fmt.Println(fi25, v4)
                                                                    var v79 int16 = -958
                                                                    _ = v79
                                                                }
                                                                if (1.5 >= (1e10 - (1.5 / 1.0))) {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 0
                                                                    }
                                                                    switch "" {
                                                                    case "a":
                                                                        var v80 uint32 = 2147483647
                                                                        _ = v80
                                                                    default:
                                                                        var v81 int64 = 0
                                                                        _ = v81
                                                                    }
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 821
                                                                    }
                                                                }
                                                                var v82 float64 = -2.25
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    var v83 int = (fi24 % ((fi25) | 1))
                                                                    _ = v83
                                                                case "":
                                                                    var v84 int16 = 42
                                                                    _ = v84
                                                                case "hello":
                                                                    var v85 int16 = 0
                                                                    _ = v85
                                                                }
                                                                _ = v82
                                                            }
                                                            fmt.Println(v3, fi24, fi25)
                                                            if ((v4 || (2 > 1)) || true) {
                                                                if v4 {
                                                                    var v86 uint16 = v3
                                                                    _ = v86
                                                                }
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                            fmt.Println(fi25, v4, v3)
                                                            if v4 {
                                                                fmt.Println(fi25, fi24)
                                                                if (!true) {
                                                                    var v87 uint = 303
                                                                    v4 = (!(127 > 128))
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (("" + "hello") + "\n\t")
                                                                    }
                                                                    _ = v87
                                                                } else {
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v3 = (v3 | v3)
                                                                    case "hello":
                                                                        v88 := GenS5{F0: false, F1: 0, F2: (("hello" + "hello") + "a")}
                                                                        _ = v88
                                                                    case "\n\t":
                                                                        var v89 uint = 42
                                                                        _ = v89
                                                                    }
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        var v90 bool = false
                                                                        _ = v90
                                                                    case "":
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    }
                                                                    v91 := GenS9{F0: v3}
                                                                    _ = v91
                                                                }
                                                            }
                                                        }
                                                    case -1:
                                                        var v92 int64 = 628
                                                        _ = v92
                                                    }
                                                default:
                                                    v93 := GenS9{F0: (v3 << ((384) & 63))}
                                                    _ = v93
                                                }
                                                var v94 int32 = 42
                                                fmt.Println(fi25)
                                                _ = v94
                                            case "":
                                                var v95 int32 = int32(fi25)
                                                _ = v95
                                            }
                                        } else {
                                            if v4 {
                                                v3 = (v3 + (v3 - v3))
                                                if v4 {
                                                    v4 = (!(v4 && ((fi25 & fi24) <= fi25)))
                                                }
                                            }
                                            v4 = (129 != -1)
                                        }
                                        v4 = v4
                                    case 1:
                                        if (true || v4) {
                                            fmt.Println(v3, v4, fi25)
                                            var v96 uint64 = 42
                                            if ("hello" != "") {
                                                switch ("foo bar" + ((("hello" + "") + "foo bar") + "")) {
                                                case "":
                                                    if v4 {
                                                        var v97 float64 = (-2.25 / -2.25)
                                                        _ = v97
                                                    } else {
                                                        var v98 uint64 = 18446744073709551615
                                                        fmt.Println(v3)
                                                        switch "" {
                                                        case "hello":
                                                            if len(v2) > 0 {
                                                                v2[0] = 2147483648
                                                            }
                                                        case "foo bar":
                                                            var v99 int64 = -1
                                                            _ = v99
                                                        default:
                                                            v100 := GenS9{F0: (^(v3))}
                                                            _ = v100
                                                        }
                                                        _ = v98
                                                    }
                                                    var v101 int16 = 42
                                                    _ = v101
                                                case "foo bar":
                                                    if v4 {
                                                        if (1 <= 1) {
                                                            switch 2147483647 {
                                                            case 42:
                                                                v102 := GenS2{F0: 25, F1: 255}
                                                                _ = v102
                                                            case 32768:
                                                                fmt.Println(fi25, fi25)
                                                            case -760:
                                                                if false {
                                                                    var v103 uint = 4294967295
                                                                    v4 = (v4 && (-1 == -111))
                                                                    _ = v103
                                                                } else {
                                                                    var v104 uint = 65535
                                                                    _ = v104
                                                                }
                                                            }
                                                        } else {
                                                            v3 = (v3 << ((860) & 63))
                                                            fmt.Println(v3)
                                                            v2 = append(v2, 642)
                                                        }
                                                        v105 := GenS10{F0: 128}
                                                        _ = v105
                                                    }
                                                case "\n\t":
                                                    var v106 uint8 = uint8(v3)
                                                    _ = v106
                                                }
                                            }
                                            fmt.Println(fi24)
                                            v4 = v4
                                            _ = v96
                                        }
                                    default:
                                        if true {
                                            v3 += 2
                                            var v107 int = 2147483647
                                            if v4 {
                                                fmt.Println(v3)
                                            }
                                            _ = len(v5)
                                            _ = v107
                                        } else {
                                            switch (fi25 & 128) {
                                            case 280:
                                                fmt.Println(fi25)
                                            case 32767:
                                                v108 := GenS10{F0: 0}
                                                _ = len(v5)
                                                switch 0 {
                                                case 2147483647:
                                                    var v109 int = (fi24 % ((fi24) | 1))
                                                    _ = v109
                                                case 42:
                                                    v3 += 0
                                                default:
                                                    switch -9223372036854775808 {
                                                    case 32768:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    case 1:
                                                        fmt.Println(fi24)
                                                    }
                                                }
                                                _ = v108
                                            default:
                                                var v110 uint8 = 230
                                                _ = v110
                                            }
                                            if v4 {
                                                if ((!(0 != -1)) && v4) {
                                                    if len(v1) > 0 {
                                                        v1[0] = ("hello" + "a")
                                                    }
                                                    var v111 int32 = -1
                                                    _ = v111
                                                } else {
                                                    v4 = (2 >= 1)
                                                    var v112 int8 = -32
                                                    _ = v112
                                                }
                                            } else {
                                                var v113 uint8 = genFn1()
                                                _ = v113
                                            }
                                            fmt.Println(v4, fi24)
                                        }
                                    }
                                } else {
                                    var v114 uint16 = (v3 >> ((4294967295) & 63))
                                    _ = v114
                                }
                                if len(v5) > 0 {
                                    v5[0] = fi24
                                }
                                _ = fi25
                                _ = v26
                            }
                            var v115 float32 = (-2.25 / 1.0)
                            _ = fi24
                            _ = v115
                        }
                    case "\n\t":
                        fmt.Println(v3, v4)
                    }
                    _ = len(v5)
                    if (false || true) {
                        v116 := GenS6{F0: genFn0(-725, -1), F1: (v3 * (^(v3))), F2: 0}
                        v3 = 65535
                        v4 = v4
                        _ = v116
                    } else {
                        v3 = (v3 << ((65535) & 63))
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                var v117 int32 = 1
                var v118 uint64 = 0
                _ = v117
                _ = v118
            }
        }
    }
    var v119 uint8 = 255
    if (1000 == -470) {
        if v4 {
            v3 += v3
            v120 := GenS6{F0: 0, F1: 255, F2: genFn0(0, 0)}
            _ = v120
        }
        v119 += (^(v119))
        v3 += v3
    }
    var v121 uint32 = 1
    switch "a" {
    case "foo bar":
        v4 = (9223372036854775807 >= -710)
    default:
        v122 := GenS10{F0: genFn1()}
        _ = v122
    }
    var v123 uint16 = v3
    var v124 float64 = ((-2.25 + ((1.5 / 0.1) / (1.5 * 0.1))) * 1.5)
    for fi125 := 0; fi125 < 3; fi125++ {
        if true {
            fmt.Println(fi125)
        }
        fmt.Println(v119)
        v126 := GenS8{F0: (^(v123))}
        _ = fi125
        _ = v126
    }
    _ = len(v1)
    var v127 uint8 = 1
    switch 0 {
    case 883:
        switch ("\n\t" + "a") {
        case "\n\t":
            var v128 int8 = -1
            _ = v128
        default:
            if (v121 < v121) {
                switch ("" + "") {
                case "a":
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                case "":
                    v129 := GenS11{F0: v127, F1: v3}
                    _ = v129
                case "\n\t":
                    v121 -= v121
                }
                v130 := GenS10{F0: (v127 << ((4294967295) & 63))}
                for fi131 := 0; fi131 < 12; fi131++ {
                    var v132 int64 = 375
                    v133 := GenS11{F0: (v119 << ((4294967295) & 63)), F1: (^(v123))}
                    v134 := GenS7{F0: 42, F1: ""}
                    _ = fi131
                    _ = v132
                    _ = v133
                    _ = v134
                }
                _ = v130
            }
        }
    case 127:
        v2 = append(v2, v121)
    default:
        fmt.Println(v119)
    }
    v135 := GenS2{F0: (v121 % ((1) | 1)), F1: 192}
    var v136 uint16 = 42
    v127 -= (v119 * v127)
    v124 += v124
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v119:", v119)
    fmt.Println("v121:", v121)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v127:", v127)
    fmt.Println("v136:", v136)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v119
    _ = v121
    _ = v123
    _ = v124
    _ = v127
    _ = v135
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
