package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 uint32; F2 uint8 }

func main() {
    var v1 int16 = 1
    var v2 int32 = 42
    var v3 int16 = 42
    fmt.Println(v3)
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 int16 = (v3 ^ (v3 >> ((4294967295) & 63)))
        var v6 int = 895
        if true {
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := []int{}
                v9 := GenS1{F0: 4294967295, F1: 2, F2: "hello"}
                _ = fi7
                _ = v8
                _ = v9
            }
            v3 -= v1
            if false {
                v10 := GenS2{F0: false}
                v11 := GenS1{F0: 2, F1: 65535, F2: ("a" + "\n\t")}
                v12 := GenS2{F0: (!(false && true))}
                _ = v10
                _ = v11
                _ = v12
            }
        } else {
            switch "foo bar" {
            case "hello":
                var v13 int64 = 992
                _ = v13
            default:
                v14 := []int{(fi4 & (-(fi4))), fi4, fi4}
                _ = v14
            }
            if true {
                v15 := []int{}
                var v16 uint32 = 460
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 int64 = 1
                    fmt.Println(v16, v2, v1)
                    if true {
                        var v19 int64 = (v18 & (v18 << ((18446744073709551615) & 63)))
                        _ = v19
                    } else {
                        fmt.Println(v2)
                    }
                    _ = fi17
                    _ = v18
                }
                _ = v15
                _ = v16
            }
            var v20 uint32 = 1
            _ = v20
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(v2, v3, v3)
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            switch (fi21 * fi21) {
            case 0:
                v2 = (v2 | (-(v2)))
            case -971:
                fmt.Println(v1, fi21, v3)
            case 42:
                var v23 int8 = 45
                if ((1 > 42) || (-2.25 <= float32((3.14159 - 0.0)))) {
                    fmt.Println(v1, v2)
                } else {
                    v3 -= (v3 * 32767)
                }
                _ = v23
            }
            v24 := GenS2{F0: (-9223372036854775808 == int64(v2))}
            if (65535 >= 65535) {
                if (!false) {
                    var v25 uint8 = 128
                    _ = v25
                }
                v26 := GenS2{F0: ((-874 > -1) || (230 < 1))}
                v27 := GenS2{F0: ((v2 > 0) && true)}
                _ = v26
                _ = v27
            } else {
                if false {
                    var v28 uint16 = 32767
                    if false {
                        fmt.Println(v1)
                        v2 = (v2 ^ 42)
                        switch "\n\t" {
                        case "hello":
                            v29 := GenS1{F0: 0, F1: v28, F2: "hello"}
                            _ = v29
                        case "":
                            var v30 int = (fi21 % (((fi22 ^ 101)) | 1))
                            _ = v30
                        case "\n\t":
                            fmt.Println(fi21, v2, fi22)
                        default:
                            v31 := []int{(fi22 * fi21), fi21, (fi21 + fi22)}
                            _ = v31
                        }
                    }
                    if (0 <= 42) {
                        var v32 uint = 0
                        _ = v32
                    } else {
                        v3 = 52
                    }
                    _ = v28
                } else {
                    if (!true) {
                        fmt.Println(v2, v2)
                        fmt.Println(v1, v1, v2)
                        var v33 int32 = v2
                        _ = v33
                    } else {
                        v3 += v3
                    }
                    v34 := []int{9223372036854775807, fi22, fi22}
                    var v35 uint = 42
                    _ = v34
                    _ = v35
                }
                fmt.Println(v2, fi21, fi21)
                v36 := []int{-598, fi22, 2147483647}
                _ = v36
            }
            _ = fi22
            _ = v24
        }
        v37 := GenS2{F0: (!true)}
        _ = fi21
        _ = v37
    }
    v3 = (v3 & (v1 << ((2) & 63)))
    v3 -= 42
    var v38 uint16 = 61
    if (("a" + "\n\t") > "\n\t") {
        fmt.Println(v2, v1, v3)
        v38 -= (^(v38))
        switch "a" {
        case "hello":
            v39 := GenS1{F0: 2, F1: (v38 * v38), F2: "foo bar"}
            _ = v39
        case "":
            v3 -= 42
        case "a":
            var v40 uint32 = 2147483647
            _ = v40
        default:
            v41 := GenS3{F0: v2}
            _ = v41
        }
    } else {
        if true {
            for fi42 := 0; fi42 < 16; fi42++ {
                var v43 int = fi42
                _ = fi42
                _ = v43
            }
            for fi44 := 0; fi44 < 2; fi44++ {
                if true {
                    v45 := []int{(^(fi44)), fi44, (^(fi44)), (fi44 << ((65535) & 63))}
                    if false {
                        var v46 bool = (("hello" + "a") == "a")
                        for fi47 := 0; fi47 < 4; fi47++ {
                            switch ("foo bar" + "\n\t") {
                            case "foo bar":
                                v48 := GenS4{F0: uint64(v38)}
                                _ = v48
                            case "a":
                                v49 := GenS5{F0: v46, F1: -1, F2: (^(v38))}
                                _ = v49
                            case "hello":
                                switch (fi44 - 893) {
                                case 2147483647:
                                    v2 += v2
                                case 741:
                                    v3 = 32767
                                }
                            }
                            fmt.Println(v2, v46, v2)
                            v38 += 42
                            _ = fi47
                        }
                        if (65535 >= 1) {
                            var v50 string = ("" + "foo bar")
                            switch (^(fi44)) {
                            case -484:
                                if (!(!false)) {
                                    v3 -= (v1 + (v3 >> ((42) & 63)))
                                    for fi51 := 0; fi51 < 3; fi51++ {
                                        var v52 float32 = -2.25
                                        v53 := GenS6{F0: 375}
                                        _ = fi51
                                        _ = v52
                                        _ = v53
                                    }
                                }
                            case 1:
                                fmt.Println(v38)
                            }
                            _ = v50
                        }
                        _ = v46
                    } else {
                        var v54 float64 = float64(float32(1e10))
                        switch "a" {
                        case "":
                            var v55 uint = 255
                            _ = v55
                        case "foo bar":
                            v3 = v3
                        default:
                            v38 = v38
                        }
                        fmt.Println(v38, v38, fi44)
                        _ = v54
                    }
                    v3 += (v1 & (v3 % ((v1) | 1)))
                    _ = v45
                }
                var v56 uint64 = 1
                v38 -= 32767
                _ = fi44
                _ = v56
            }
        }
        var v57 float32 = (0.1 / 0.1)
        if (v1 != (v1 + (v3 ^ (v3 << ((255) & 63))))) {
            v38 = (^(v38))
            switch (("foo bar" + "\n\t") + ("hello" + "\n\t")) {
            case "":
                switch -1 {
                case 2147483647:
                    var v58 bool = true
                    _ = v58
                case 32767:
                    var v59 uint8 = 242
                    _ = v59
                case -442:
                    var v60 int8 = 1
                    _ = v60
                }
            case "a":
                v61 := GenS6{F0: 1}
                _ = v61
            case "hello":
                if ((v38 << ((319) & 63)) != (v38 & v38)) {
                    v62 := GenS7{F0: "a", F1: 2147483648, F2: 2}
                    for fi63 := 0; fi63 < 9; fi63++ {
                        if false {
                            v3 -= 0
                            v1 = int16(v2)
                            v64 := GenS4{F0: 42}
                            _ = v64
                        } else {
                            var v65 int8 = 127
                            for fi66 := 0; fi66 < 10; fi66++ {
                                v38 = (v38 ^ (v38 & (v38 - v38)))
                                var v67 int8 = (-(v65))
                                fmt.Println(v38, v38, v2)
                                _ = fi66
                                _ = v67
                            }
                            _ = v65
                        }
                        v68 := []int{(fi63 - (fi63 | (fi63 ^ 32767))), fi63}
                        if (("a" < "hello") && ((v3 | (v1 * (v1 >> ((849) & 63)))) > v1)) {
                            if false {
                                fmt.Println(v3)
                            } else {
                                for fi69 := 0; fi69 < 5; fi69++ {
                                    _ = len(v68)
                                    _ = fi69
                                }
                            }
                        }
                        _ = fi63
                        _ = v68
                    }
                    if false {
                        if false {
                            if (-2.25 > ((((v57 / (v57 * v57)) - v57) + v57) - (3.14159 / v57))) {
                                var v70 int32 = v2
                                _ = v70
                            } else {
                                switch ("foo bar" + ("foo bar" + ("\n\t" + ("a" + "foo bar")))) {
                                case "\n\t":
                                    v57 -= float32((-2.25 / 1.0))
                                case "foo bar":
                                    if ("foo bar" <= "foo bar") {
                                        if false {
                                            fmt.Println(v1)
                                            switch "a" {
                                            case "":
                                                for fi71 := 0; fi71 < 18; fi71++ {
                                                    fmt.Println(v3, v38, v1)
                                                    v72 := []int{}
                                                    _ = fi71
                                                    _ = v72
                                                }
                                            default:
                                                switch "hello" {
                                                case "a":
                                                    v73 := GenS1{F0: 0, F1: v38, F2: "a"}
                                                    _ = v73
                                                case "":
                                                    v74 := []int{42}
                                                    _ = v74
                                                case "\n\t":
                                                    v38 += 0
                                                }
                                            }
                                        }
                                        switch "a" {
                                        case "":
                                            fmt.Println(v57)
                                        case "\n\t":
                                            if true {
                                                v75 := []int{-460, 9223372036854775807}
                                                _ = v75
                                            } else {
                                                var v76 uint64 = 889
                                                _ = v76
                                            }
                                        default:
                                            for fi77 := 0; fi77 < 7; fi77++ {
                                                var v78 int = fi77
                                                _ = fi77
                                                _ = v78
                                            }
                                        }
                                        switch 674 {
                                        case 32768:
                                            v79 := GenS6{F0: 844}
                                            _ = v79
                                        default:
                                            if false {
                                                v1 = v3
                                                v80 := []int{}
                                                _ = v80
                                            }
                                        }
                                    }
                                }
                                var v81 float32 = (v57 * (v57 * v57))
                                v82 := []int{}
                                _ = v81
                                _ = v82
                            }
                        }
                    }
                    _ = v62
                }
            default:
                v83 := []int{-1}
                _ = v83
            }
        } else {
            v1 += -1
        }
        _ = v57
    }
    var v84 uint64 = 676
    var v85 uint32 = 498
    v3 -= (v3 - v1)
    v86 := []int{337, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("slicelen v86:", len(v86), v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v84
    _ = v85
    _ = v86
}
