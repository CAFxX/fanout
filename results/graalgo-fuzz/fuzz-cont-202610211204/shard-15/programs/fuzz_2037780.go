package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int8; F1 string; F2 int8 }

func main() {
    var v1 uint8 = 0
    v2 := []int{4294967295, 42}
    v3 := make(map[int]string)
    var v4 int8 = -59
    var v5 int32 = 42
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int64 = 9223372036854775807
        var v8 float32 = 3.14159
        v3[int(v4)] = "foo bar"
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int32 = (-(v5))
    var v10 uint16 = 53
    v11 := v3[2147483647]
    v4 += (v4 * (v4 / ((-38) | 1)))
    switch "foo bar" {
    case "\n\t":
        v12 := GenS1{F0: 4294967295}
        _ = v12
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: 65535}
        if (!false) {
            var v15 int8 = (-(v4))
            if false {
                var v16 int64 = int64(fi13)
                _ = v16
            } else {
                var v17 string = ""
                _ = v17
            }
            _ = v15
        } else {
            _ = len(v3)
        }
        _ = fi13
        _ = v14
    }
    v18 := GenS2{F0: (v5 / ((v5) | 1))}
    if true {
        if true {
            v9 -= 0
            if ("\n\t" > v11) {
                v9 -= v9
                var v19 int8 = v4
                _ = v19
            }
        } else {
            var v20 uint64 = 42
            v5 += int32(v20)
            _ = v20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            v22 := v3[(fi21 << ((65535) & 63))]
            if true {
                v23 := GenS1{F0: 4294967295}
                _ = v23
            } else {
                if false {
                    if len(v2) > 0 {
                        v2[0] = fi21
                    }
                } else {
                    v24 := GenS3{F0: ("\n\t" + v11)}
                    _ = v24
                }
                if (-30 == -199) {
                    for fi25 := 0; fi25 < 6; fi25++ {
                        var v26 float32 = (0.0 - (-2.25 * 1.5))
                        switch "hello" {
                        case "a":
                            v3[int(v5)] = v11
                        case "foo bar":
                            switch fi21 {
                            case -397:
                                if false {
                                    v5 -= 42
                                    if false {
                                        var v27 uint = 984
                                        _ = v27
                                    } else {
                                        if (false && true) {
                                            var v28 uint64 = 4294967295
                                            fmt.Println(v5, fi21)
                                            fmt.Println(fi21, v22)
                                            if (1 == 663) {
                                                if ((-9223372036854775808 >= -9223372036854775808) && false) {
                                                    var v29 int32 = v9
                                                    v11 = v22
                                                    _ = v29
                                                } else {
                                                    if false {
                                                        switch fi21 {
                                                        case 599:
                                                            var v30 uint8 = v1
                                                            _ = v30
                                                        case 545:
                                                            fmt.Println(v5)
                                                        case 4294967295:
                                                            if (!true) {
                                                                v1 += 0
                                                                switch fi25 {
                                                                case -764:
                                                                    switch v22 {
                                                                    case "\n\t":
                                                                        _ = len(v3)
                                                                    case "a":
                                                                        if false {
                                                                            v4 -= (v4 - (v4 >> ((65535) & 63)))
                                                                            var v31 uint8 = v1
                                                                            _ = v31
                                                                        }
                                                                    case "foo bar":
                                                                        var v32 uint64 = uint64(v1)
                                                                        _ = v32
                                                                    default:
                                                                        var v33 bool = (769 >= 65535)
                                                                        _ = v33
                                                                    }
                                                                }
                                                            } else {
                                                                v26 = 3.14159
                                                            }
                                                            switch "\n\t" {
                                                            case "a":
                                                                _ = len(v2)
                                                            default:
                                                                switch fi25 {
                                                                case -9223372036854775808:
                                                                    v34 := GenS3{F0: v11}
                                                                    _ = v34
                                                                case -376:
                                                                    var v35 uint8 = (^(v1))
                                                                    _ = v35
                                                                }
                                                            }
                                                            var v36 int64 = 1
                                                            _ = v36
                                                        default:
                                                            v26 -= ((((v26 * float32((3.14159 + ((((((1.5 + (1e10 * -2.25)) / 0.1) - 1.5) + ((0.0 + 0.1) * 0.0)) - 1e10) - 1.5)))) - v26) * (float32(0.0) - v26)) + 0.0)
                                                        }
                                                    }
                                                    v26 = v26
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            var v37 uint8 = v1
                                            _ = v28
                                            _ = v37
                                        } else {
                                            v38 := GenS3{F0: v22}
                                            v10 += 42
                                            if (0 <= 2) {
                                                fmt.Println(v10, v9, v9)
                                                v39 := GenS3{F0: ((((((("hello" + v11) + (((v22 + (v22 + v11)) + "foo bar") + "foo bar")) + v22) + "\n\t") + v11) + "hello") + ((v11 + "\n\t") + "a"))}
                                                _ = v39
                                            } else {
                                                v40 := GenS4{F0: fi25}
                                                v41 := GenS5{F0: 42, F1: (v4 | (v4 >> ((255) & 63)))}
                                                var v42 float64 = 3.14159
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            }
                                            _ = v38
                                        }
                                    }
                                    var v43 int = fi25
                                    _ = v43
                                }
                            case 32767:
                                v44 := GenS3{F0: v22}
                                _ = v44
                            case 0:
                                v9 -= (v9 ^ (v9 / ((-12) | 1)))
                            }
                        case "hello":
                            switch (v22 + ((v22 + "") + "\n\t")) {
                            case "\n\t":
                                switch "" {
                                case "hello":
                                    switch -9223372036854775808 {
                                    case 25:
                                        switch fi25 {
                                        case 42:
                                            v5 = v9
                                        case 32767:
                                            if false {
                                                switch "hello" {
                                                case "foo bar":
                                                    if false {
                                                        delete(v3, (fi21 << ((255) & 63)))
                                                    } else {
                                                        if ((620 == 4294967295) && ((v4 + -1) > (v4 >> ((255) & 63)))) {
                                                            if true {
                                                                v4 += int8(v9)
                                                                delete(v3, (fi21 & (fi21 / (((fi21 >> ((2) & 63))) | 1))))
                                                                var v45 int16 = 42
                                                                _ = v45
                                                            } else {
                                                                var v46 float64 = 3.14159
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    v26 -= ((float32(3.14159) + (((-2.25 * (v26 / 1e10)) * (v26 / v26)) - -2.25)) + v26)
                                                                case "":
                                                                    v47 := GenS6{F0: 4294967295, F1: v11}
                                                                    _ = v47
                                                                case "foo bar":
                                                                    var v48 uint = 4294967295
                                                                    _ = v48
                                                                default:
                                                                    if (v11 > v11) {
                                                                        if false {
                                                                            v46 -= v46
                                                                            if false {
                                                                                var v49 int64 = 493
                                                                                _ = v49
                                                                            } else {
                                                                                v10 -= (v10 << ((uint(v1)) & 63))
                                                                                if (v26 >= v26) {
                                                                                    if (97 != 816) {
                                                                                        v11 = v22
                                                                                        var v50 int16 = 0
                                                                                        _ = v50
                                                                                    } else {
                                                                                        if ((false || (!true)) && true) {
                                                                                            var v51 string = "a"
                                                                                            _ = v51
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v9, v11, v10)
                                                                                    var v52 uint32 = 42
                                                                                    _ = v52
                                                                                }
                                                                                fmt.Println(v26, v46, v46)
                                                                            }
                                                                            fmt.Println(v9, v46, v22)
                                                                        } else {
                                                                            fmt.Println(v4, fi21, fi25)
                                                                            fmt.Println(v4, v26, fi21)
                                                                            delete(v3, (^(fi25)))
                                                                        }
                                                                    }
                                                                }
                                                                v10 -= 2
                                                                _ = v46
                                                            }
                                                            if true {
                                                                _ = len(v3)
                                                                switch fi21 {
                                                                case 42:
                                                                    fmt.Println(v10)
                                                                case 53:
                                                                    v53 := GenS7{F0: v10, F1: "foo bar"}
                                                                    _ = v53
                                                                case 2147483647:
                                                                    v22 = v11
                                                                }
                                                                if ((((0.0 - (((0.0 / 0.1) * 0.1) + (-2.25 * 1.5))) / (((1e10 + float64(3.14159)) + 1e10) - -2.25)) <= (3.14159 / 1e10)) && false) {
                                                                    fmt.Println(v5, v9, v4)
                                                                    v11 = v22
                                                                } else {
                                                                    if true {
                                                                        var v54 float32 = (float32(0.1) - ((-2.25 * (v26 + v26)) + v26))
                                                                        _ = v54
                                                                    } else {
                                                                        var v55 bool = (393 >= 71)
                                                                        _ = v55
                                                                    }
                                                                }
                                                            } else {
                                                                fmt.Println(v22, v10, fi25)
                                                            }
                                                            var v56 int32 = v9
                                                            _ = v56
                                                        }
                                                        if false {
                                                            if (!(!((-(fi25)) >= 585))) {
                                                                if false {
                                                                    _ = len(v3)
                                                                } else {
                                                                    v11 = "foo bar"
                                                                    var v57 float32 = (v26 / (float32(-2.25) / v26))
                                                                    _ = v57
                                                                }
                                                                switch v22 {
                                                                case "foo bar":
                                                                    v10 -= (v10 / (((^(v10))) | 1))
                                                                case "\n\t":
                                                                    fmt.Println(fi21)
                                                                case "a":
                                                                    switch v22 {
                                                                    case "foo bar":
                                                                        var v58 int16 = 1000
                                                                        _ = v58
                                                                    case "":
                                                                        fmt.Println(v5, v11)
                                                                    case "\n\t":
                                                                        var v59 float32 = -2.25
                                                                        _ = v59
                                                                    }
                                                                }
                                                                v60 := GenS8{F0: 842, F1: (v1 * v1), F2: 1}
                                                                _ = v60
                                                            } else {
                                                                fmt.Println(v22)
                                                                v61 := GenS6{F0: 2147483647, F1: (v22 + "a")}
                                                                if len(v2) > 0 {
                                                                    v2[0] = 0
                                                                }
                                                                _ = v61
                                                            }
                                                            v10 = (v10 / ((255) | 1))
                                                            var v62 float64 = -2.25
                                                            v63 := GenS9{F0: v11, F1: -2147483648}
                                                            fmt.Println(v62, v5, v11)
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            if false {
                                                                v64 := GenS10{F0: v10}
                                                                _ = v64
                                                            }
                                                            switch (fi25 >> ((0) & 63)) {
                                                            case 992:
                                                                switch (fi25 + fi21) {
                                                                case 4294967295:
                                                                    v65 := GenS3{F0: "\n\t"}
                                                                    _ = v65
                                                                case -968:
                                                                    if true {
                                                                        fmt.Println(v9, v11)
                                                                    }
                                                                }
                                                            case 1:
                                                                v5 = (v9 / ((-165) | 1))
                                                                fmt.Println(v1, v4)
                                                            }
                                                            var v66 uint64 = 0
                                                            _ = v66
                                                        }
                                                    }
                                                case "":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                default:
                                                    v11 = ((v11 + (("a" + "a") + (v22 + (v11 + ((v11 + v22) + v22))))) + ("\n\t" + (v22 + "hello")))
                                                }
                                                v26 = float32((1.5 * 3.14159))
                                            }
                                        default:
                                            switch 1 {
                                            case 42:
                                                v3[(fi25 + 942)] = v11
                                            }
                                        }
                                    }
                                case "a":
                                    var v67 uint64 = 631
                                    _ = v67
                                default:
                                    v9 = v5
                                }
                            default:
                                var v68 uint = 475
                                _ = v68
                            }
                        }
                        _ = fi25
                        _ = v26
                    }
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v69 uint32 = 2147483647
                    v11 = (v22 + "\n\t")
                    _ = v69
                }
            }
            for fi70 := 0; fi70 < 11; fi70++ {
                v9 = 0
                if (!((fi21 - (-(fi70))) >= fi70)) {
                    var v71 int16 = -965
                    var v72 int8 = v4
                    v72 = -46
                    _ = v71
                    _ = v72
                }
                _ = fi70
            }
            _ = fi21
            _ = v22
        }
        if (false && true) {
            switch (v11 + v11) {
            case "hello":
                v4 = v4
            default:
                var v73 uint64 = 18446744073709551615
                _ = v73
            }
            v74 := GenS9{F0: "foo bar", F1: int32(v10)}
            var v75 uint64 = 4294967295
            _ = v74
            _ = v75
        } else {
            var v76 int32 = v9
            _ = v76
        }
    } else {
        switch int(v1) {
        case -1:
            v1 = 232
        case 32767:
            delete(v3, -2147483648)
        case 54:
            for fi77 := 0; fi77 < 20; fi77++ {
                v78 := GenS4{F0: (fi77 ^ fi77)}
                _ = fi77
                _ = v78
            }
        }
        fmt.Println(v10)
        if ((!(float32((0.1 - 1.5)) < 0.1)) && true) {
            if false {
                fmt.Println(v9, v1)
            } else {
                switch -577 {
                case -1:
                    v5 = v9
                case -2147483648:
                    for fi79 := 0; fi79 < 12; fi79++ {
                        if false {
                            v9 = (v9 | v9)
                        } else {
                            if (int32(fi79) == (-(v5))) {
                                var v80 uint = 568
                                v81 := GenS11{F0: -128, F1: "", F2: v4}
                                var v82 uint64 = 2
                                _ = v80
                                _ = v81
                                _ = v82
                            }
                        }
                        _ = fi79
                    }
                case 4294967295:
                    switch 494 {
                    case -266:
                        v10 = (^(v10))
                    case 32768:
                        var v83 bool = false
                        _ = v83
                    case 2147483647:
                        fmt.Println(v1, v9)
                    default:
                        var v84 int16 = 1
                        _ = v84
                    }
                default:
                    v2 = append(v2, -336)
                }
            }
            v10 = uint16(v1)
        } else {
            v1 = v1
        }
    }
    switch -2147483648 {
    case 2147483647:
        var v85 int32 = (v5 ^ (v9 / ((42) | 1)))
        _ = v85
    case 1:
        var v86 int8 = v4
        _ = v86
    case -9223372036854775808:
        var v87 uint16 = (^(v10))
        _ = v87
    default:
        switch 543 {
        case 32768:
            v5 -= (-(v9))
        case 0:
            v4 = v4
        }
    }
    var v88 bool = true
    var v89 bool = (!v88)
    if true {
        var v90 int64 = -513
        if v88 {
            var v91 uint64 = 42
            fmt.Println(v90, v91)
            var v92 float64 = (1.5 / -2.25)
            _ = v91
            _ = v92
        } else {
            v4 -= v4
            v9 += v5
        }
        _ = v90
    } else {
        fmt.Println(v4, v4, v1)
        v3[-9223372036854775808] = v11
        v88 = ((v9 ^ (-(v5))) >= v9)
    }
    if (1 != 1) {
        var v93 string = v11
        if (!(v89 || v89)) {
            v93 = ""
            v5 = v5
            var v94 int8 = 42
            _ = v94
        }
        _ = v93
    } else {
        var v95 int16 = -1
        v96 := v3[0]
        v4 = (v4 - 106)
        _ = v95
        _ = v96
    }
    v4 -= (v4 / ((42) | 1))
    v11 = (v11 + v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v88
    _ = v89
    _ = v9
}
