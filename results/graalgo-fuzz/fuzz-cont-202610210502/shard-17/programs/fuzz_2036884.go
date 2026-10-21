package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 bool; F2 uint }
type GenS4 struct { F0 int8; F1 int16; F2 string }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 string; F1 uint16; F2 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 uint; F1 int; F2 uint16 }
type GenS10 struct { F0 int8; F1 uint64; F2 int8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16; F1 uint }
type GenS13 struct { F0 int16 }
func genFn0(p0 int8, p1 int, p2 int16) uint16 {
    if false {
        fmt.Println(p0, p0, p1)
    } else {
        p0 += -128
    }
    p0 = int8(p2)
    _ = p0
    _ = p1
    _ = p2
    return 832
}
func genFn1(p0 int8, p1 uint8, p2 int8) uint {
    v4 := GenS1{F0: 789, F1: "\n\t"}
    p1 -= 127
    fmt.Println(p1, p1, p2)
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return uint(p1)
}

func main() {
    var v1 uint32 = 613
    v2 := make(map[string]bool)
    var v3 uint16 = 550
    v2["foo bar"] = (0 <= 9223372036854775807)
    delete(v2, "hello")
    var v5 float64 = (1e10 + 0.1)
    v6 := v2["foo bar"]
    switch (("\n\t" + "foo bar") + "foo bar") {
    case "a":
        if v6 {
            if ("a" <= "\n\t") {
                v5 = 1e10
                var v7 float64 = v5
                _ = v7
            }
        }
    case "hello":
        v8 := []int{128}
        _ = v8
    case "":
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: 188, F1: ""}
            var v11 float64 = (v5 - float64(0.1))
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        v3 += genFn0(42, 127, 1000)
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS2{F0: 1}
        if true {
            v6 = ((1.5 * 1e10) != v5)
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 int64 = 0
                var v16 int64 = (v15 % (((v15 | v15)) | 1))
                v17 := GenS3{F0: (v3 * uint16(v1)), F1: (v6 && true), F2: 18446744073709551615}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            v18 := GenS3{F0: genFn0(-1, (^(fi12)), 1), F1: v6, F2: 987}
            _ = v18
        } else {
            v3 = 65535
            if ((-100 > 42) || (0 > -543)) {
                switch 2147483647 {
                case 635:
                    v1 = (v1 ^ (v1 << ((92) & 63)))
                case -156:
                    var v19 bool = (1 != 4294967295)
                    _ = v19
                case -685:
                    for fi20 := 0; fi20 < 11; fi20++ {
                        v3 += v3
                        switch "\n\t" {
                        case "foo bar":
                            v21 := GenS4{F0: -114, F1: -32768, F2: ("\n\t" + ("hello" + "a"))}
                            _ = v21
                        case "hello":
                            switch "foo bar" {
                            case "":
                                if true {
                                    if v6 {
                                        _ = len(v2)
                                        v22 := GenS1{F0: genFn1(-1, 2, -102), F1: (("\n\t" + "\n\t") + ("" + "foo bar"))}
                                        _ = v22
                                    }
                                }
                            default:
                                switch (-(fi12)) {
                                case -530:
                                    if v6 {
                                        v23 := GenS3{F0: 32767, F1: v6, F2: 4294967295}
                                        v24 := GenS4{F0: 37, F1: int16(v1), F2: (("a" + "a") + ("\n\t" + "foo bar"))}
                                        v25 := GenS4{F0: 42, F1: 1, F2: "foo bar"}
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                case 1:
                                    fmt.Println(v5, fi12)
                                case 448:
                                    var v26 uint32 = 2147483648
                                    _ = v26
                                }
                            }
                        case "a":
                            if false {
                                v27 := GenS1{F0: genFn1(-128, 35, 100), F1: ""}
                                _ = v27
                            }
                        default:
                            switch fi20 {
                            case -149:
                                v5 = v5
                            case -1:
                                v28 := GenS5{F0: v3, F1: 157}
                                _ = v28
                            case -68:
                                if (18446744073709551615 < 854) {
                                    var v29 int = (fi12 | fi20)
                                    v30 := GenS6{F0: "foo bar", F1: v3, F2: 222}
                                    if (fi20 > -215) {
                                        delete(v2, "foo bar")
                                    }
                                    _ = v29
                                    _ = v30
                                } else {
                                    fmt.Println(fi12, fi12)
                                }
                            }
                        }
                        v1 = (v1 ^ v1)
                        switch ("hello" + "foo bar") {
                        case "a":
                            v31 := v2["a"]
                            _ = v31
                        case "foo bar":
                            if v6 {
                                v32 := GenS7{F0: (v6 && v6)}
                                v33 := GenS8{F0: genFn1(6, 142, -56), F1: "hello"}
                                var v34 int = 2147483647
                                _ = v32
                                _ = v33
                                _ = v34
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    var v35 int16 = -524
                                    _ = v35
                                case "":
                                    v3 += 42
                                }
                            }
                        }
                        fmt.Println(v5, v3, v6)
                        _ = fi20
                    }
                default:
                    if (196 >= 2) {
                        switch ("foo bar" + "") {
                        case "\n\t":
                            v3 += genFn0(-111, fi12, 0)
                        case "hello":
                            var v36 int32 = -259
                            _ = v36
                        }
                        var v37 string = "a"
                        _ = v37
                    } else {
                        if true {
                            v6 = ((18446744073709551615 == 936) || v6)
                        } else {
                            for fi38 := 0; fi38 < 6; fi38++ {
                                v6 = v6
                                v5 -= v5
                                if (-760 > 56) {
                                    v39 := []int{}
                                    v1 = 4294967295
                                    _ = v39
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v40 := []int{fi12, fi38}
                                        _ = v40
                                    default:
                                        if v6 {
                                            v3 = v3
                                        } else {
                                            var v41 uint8 = 6
                                            var v42 int = 9223372036854775807
                                            fmt.Println(fi12)
                                            _ = v41
                                            _ = v42
                                        }
                                    }
                                    switch fi38 {
                                    case 1:
                                        v6 = true
                                    }
                                    var v43 string = ("hello" + ("foo bar" + (("\n\t" + ("foo bar" + "hello")) + ((("hello" + ("\n\t" + "")) + "") + "foo bar"))))
                                    v44 := GenS1{F0: 0, F1: v43}
                                    v45 := GenS8{F0: 694, F1: ((v43 + v43) + v43)}
                                    v3 -= 592
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                                _ = fi38
                            }
                        }
                        v3 = v3
                        for fi46 := 0; fi46 < 15; fi46++ {
                            if (!v6) {
                                v1 += 2
                            } else {
                                var v47 uint = genFn1(-128, 128, -13)
                                switch (fi46 % (((fi12 << ((189) & 63))) | 1)) {
                                case 9223372036854775807:
                                    var v48 int8 = -1
                                    _ = v48
                                case 0:
                                    v49 := GenS7{F0: (1000 < -1)}
                                    _ = v49
                                default:
                                    var v50 string = ""
                                    _ = v50
                                }
                                _ = v47
                            }
                            fmt.Println(fi46, v6)
                            _ = fi46
                        }
                    }
                }
                var v51 int64 = 1
                if v6 {
                    for fi52 := 0; fi52 < 8; fi52++ {
                        fmt.Println(v6, v5)
                        v53 := GenS7{F0: v6}
                        if true {
                            var v54 uint = 18446744073709551615
                            var v55 uint16 = 0
                            _ = v54
                            _ = v55
                        } else {
                            v56 := v2[("\n\t" + "")]
                            var v57 uint16 = v3
                            v58 := GenS3{F0: uint16(v1), F1: (!(-1 == 4)), F2: 255}
                            v5 -= (0.1 / 1.0)
                            var v59 uint = 746
                            if true {
                                switch "\n\t" {
                                case "a":
                                    if v6 {
                                        v60 := GenS7{F0: (!((!(!false)) && (!(v6 || (int32(v59) != 42)))))}
                                        switch fi12 {
                                        case -568:
                                            fmt.Println(fi52, v51)
                                        case 127:
                                            switch fi12 {
                                            case 430:
                                                var v61 uint64 = 2
                                                _ = v61
                                            case 0:
                                                v6 = ((("a" + ("\n\t" + "a")) + "foo bar") < "a")
                                            case -354:
                                                v5 = 1e10
                                            }
                                        case -2147483648:
                                            v62 := GenS7{F0: ((v5 >= v5) && false)}
                                            _ = v62
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v3)
                                        case "foo bar":
                                            var v63 int16 = 885
                                            _ = v63
                                        case "a":
                                            var v64 uint64 = 1
                                            v65 := []int{(fi12 >> ((v59) & 63)), (fi52 ^ (fi12 - (fi12 % ((fi52) | 1)))), (^(fi12))}
                                            _ = v64
                                            _ = v65
                                        }
                                        v66 := []int{}
                                        _ = v60
                                        _ = v66
                                    }
                                case "hello":
                                    var v67 int8 = 127
                                    _ = v67
                                case "\n\t":
                                    v68 := GenS7{F0: (v51 > -1)}
                                    _ = v68
                                default:
                                    var v69 float32 = -2.25
                                    _ = v69
                                }
                                v70 := []int{fi52, fi52, (fi52 + -9223372036854775808), (fi52 ^ fi12)}
                                _ = v70
                            }
                            v71 := GenS9{F0: (v59 | 65535), F1: fi52, F2: v3}
                            fmt.Println(fi12)
                            v57 -= (^(v57))
                            _ = v56
                            _ = v57
                            _ = v58
                            _ = v59
                            _ = v71
                        }
                        v1 = v1
                        switch fi12 {
                        case 128:
                            fmt.Println(v5, v3, v51)
                        case -9223372036854775808:
                            v72 := []int{fi12}
                            _ = v72
                        default:
                            if v6 {
                                var v73 int = fi12
                                _ = v73
                            }
                        }
                        var v74 int8 = 0
                        var v75 uint16 = (v3 % (((v3 % ((800) | 1))) | 1))
                        v76 := []int{(-(fi52)), 127}
                        _ = len(v76)
                        v77 := GenS10{F0: v74, F1: 854, F2: (-(v74))}
                        _ = fi52
                        _ = v53
                        _ = v74
                        _ = v75
                        _ = v76
                        _ = v77
                    }
                }
                _ = v51
            } else {
                v78 := GenS11{F0: -110}
                v79 := GenS12{F0: 1000, F1: 255}
                if true {
                    v80 := GenS7{F0: v6}
                    _ = v80
                }
                _ = v78
                _ = v79
            }
        }
        v81 := []int{}
        _ = fi12
        _ = v13
        _ = v81
    }
    v82 := GenS13{F0: 26}
    var v83 bool = (v6 || (-1 >= 42))
    var v84 int32 = int32(v3)
    for fi85 := 0; fi85 < 19; fi85++ {
        if (1 > v3) {
            switch (-(fi85)) {
            case 2147483647:
                v2[("" + "")] = false
            }
        }
        _ = fi85
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v82
    _ = v83
    _ = v84
}
