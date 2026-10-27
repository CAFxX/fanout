package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32; F1 uint }
func genFn0(p0 int8, p1 int) uint64 {
    v4 := []int{(p1 % (((-(p1))) | 1)), (p1 * (p1 - (p1 % (((-(p1))) | 1))))}
    var v5 int16 = -1
    fmt.Println(p1)
    var v6 uint = 2
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1() int16 {
    switch "hello" {
    case "hello":
        v7 := []int{-108, 0}
        _ = v7
    }
    var v8 uint8 = 255
    if true {
        var v9 int16 = 273
        _ = v9
    }
    _ = v8
    return 1
}

func main() {
    var v1 int = -1
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    if true {
        v10 := GenS1{F0: true, F1: (v1 * 191)}
        v1 -= v1
        v11 := GenS1{F0: false, F1: (v1 & v1)}
        _ = v10
        _ = v11
    }
    var v12 int = -9223372036854775808
    v3[(v12 - -2147483648)] = (("a" + "a") + "foo bar")
    v13 := []int{(v1 | (v12 << ((2) & 63))), v1}
    if (207 <= 255) {
        _ = len(v2)
        v14 := GenS2{F0: ("hello" + "\n\t"), F1: 1}
        _ = v14
    } else {
        v1 = (v12 + -1)
    }
    fmt.Println(v1, v12)
    v1 = 1
    if true {
        var v15 int8 = 127
        _ = v15
    }
    var v16 bool = (true && true)
    fmt.Println(v16, v16)
    var v17 int16 = 0
    v18 := v2["\n\t"]
    fmt.Println(v16)
    if (-9 == -128) {
        switch "foo bar" {
        case "":
            switch "foo bar" {
            case "":
                var v19 float64 = float64((0.1 + 0.1))
                _ = v19
            case "\n\t":
                v18 = true
            case "a":
                if (v18 && (918 > 798)) {
                    switch v1 {
                    case -155:
                        for fi20 := 0; fi20 < 12; fi20++ {
                            switch v12 {
                            case -9223372036854775808:
                                v16 = v16
                            }
                            for fi21 := 0; fi21 < 15; fi21++ {
                                v12 = (-(fi20))
                                v17 += genFn1()
                                if (float32(((-2.25 / 1e10) / 1e10)) >= float32((1e10 / -2.25))) {
                                    v17 = v17
                                    var v22 uint8 = 113
                                    _ = v22
                                }
                                v1 -= (-(v12))
                                _ = fi21
                            }
                            _ = fi20
                        }
                    default:
                        fmt.Println(v16)
                    }
                    var v23 int = (-(v12))
                    _ = v23
                } else {
                    v17 -= v17
                    fmt.Println(v12, v16, v12)
                    if (1 != 1) {
                        if v16 {
                            v24 := GenS2{F0: "a", F1: 2}
                            for fi25 := 0; fi25 < 5; fi25++ {
                                var v26 uint = 65535
                                _ = fi25
                                _ = v26
                            }
                            var v27 int64 = 1
                            _ = v24
                            _ = v27
                        } else {
                            for fi28 := 0; fi28 < 7; fi28++ {
                                if (!(42 == 1)) {
                                    var v29 int8 = 1
                                    _ = v29
                                } else {
                                    for fi30 := 0; fi30 < 9; fi30++ {
                                        v1 = fi28
                                        fmt.Println(v16, v17, fi28)
                                        _ = fi30
                                    }
                                    fmt.Println(v16, v12, fi28)
                                    v31 := GenS2{F0: ("\n\t" + "\n\t"), F1: 686}
                                    _ = v31
                                }
                                v32 := GenS2{F0: "\n\t", F1: 559}
                                v2[("\n\t" + "foo bar")] = v18
                                _ = fi28
                                _ = v32
                            }
                        }
                        var v33 int8 = 42
                        var v34 float64 = float64(float32(((1.5 * 0.1) + float64((0.0 / 1.5)))))
                        _ = v33
                        _ = v34
                    } else {
                        v16 = (v18 || v16)
                        if (true && v16) {
                            v35 := GenS3{F0: 2}
                            v17 = (v17 << ((255) & 63))
                            if v16 {
                                v36 := GenS4{F0: -2147483648, F1: (v18 || true)}
                                _ = v36
                            } else {
                                v18 = (32 < 2)
                                var v37 float32 = (-2.25 / (0.0 + 3.14159))
                                _ = v37
                            }
                            _ = v35
                        }
                        v16 = false
                    }
                }
            default:
                for fi38 := 0; fi38 < 13; fi38++ {
                    for fi39 := 0; fi39 < 16; fi39++ {
                        fmt.Println(v1, v1, v12)
                        if v16 {
                            var v40 int16 = -784
                            v17 -= (v40 << ((2) & 63))
                            if (("a" + ("" + "foo bar")) <= "a") {
                                fmt.Println(v1, v16)
                                v18 = (!true)
                            } else {
                                var v41 uint64 = genFn0(0, 128)
                                var v42 int64 = -9223372036854775808
                                v12 = fi39
                                v43 := GenS5{F0: 2, F1: 42}
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            _ = v40
                        } else {
                            if false {
                                v12 = -973
                                v16 = v18
                                _ = len(v3)
                            } else {
                                if v16 {
                                    fmt.Println(v17)
                                    v12 = 0
                                    if (false || (-1 != 0)) {
                                        fmt.Println(v12, v1)
                                        v17 += (v17 % ((575) | 1))
                                        v3[v1] = "foo bar"
                                    }
                                    _ = len(v13)
                                } else {
                                    if v18 {
                                        fmt.Println(v12)
                                    } else {
                                        var v44 int32 = 0
                                        v45 := GenS6{F0: ((v1 == (v12 | v12)) || v16), F1: 339, F2: 42}
                                        fmt.Println(fi38)
                                        _ = v44
                                        _ = v45
                                    }
                                }
                                switch v1 {
                                case 127:
                                    switch "" {
                                    case "hello":
                                        var v46 float64 = 1.5
                                        fmt.Println(v46)
                                        _ = v46
                                    case "":
                                        v17 -= genFn1()
                                    }
                                }
                            }
                            v3[(fi39 ^ 43)] = ""
                            switch v12 {
                            case 32767:
                                v47 := GenS1{F0: (v18 || (v1 <= (-(fi38)))), F1: v1}
                                _ = v47
                            default:
                                var v48 uint32 = 825
                                _ = v48
                            }
                        }
                        v16 = true
                        _ = fi39
                    }
                    _ = fi38
                }
            }
        default:
            if true {
                for fi49 := 0; fi49 < 18; fi49++ {
                    for fi50 := 0; fi50 < 9; fi50++ {
                        v51 := GenS5{F0: 2147483647, F1: 675}
                        v16 = true
                        if ((v18 && v18) && false) {
                            var v52 bool = false
                            _ = v52
                        } else {
                            delete(v2, "hello")
                        }
                        _ = fi50
                        _ = v51
                    }
                    v1 = fi49
                    var v53 int64 = -1
                    _ = fi49
                    _ = v53
                }
            }
        }
        var v54 uint16 = 630
        v55 := GenS7{F0: -1}
        _ = v54
        _ = v55
    } else {
        for fi56 := 0; fi56 < 3; fi56++ {
            var v57 uint64 = 42
            v1 = 4294967295
            var v58 int8 = -101
            _ = fi56
            _ = v57
            _ = v58
        }
        for fi59 := 0; fi59 < 3; fi59++ {
            switch v1 {
            case 9223372036854775807:
                v18 = (255 > 125)
            case 105:
                _ = len(v2)
            default:
                if v16 {
                    fmt.Println(v18)
                } else {
                    var v60 float64 = 1.5
                    for fi61 := 0; fi61 < 16; fi61++ {
                        delete(v2, "\n\t")
                        var v62 string = "a"
                        v63 := GenS1{F0: v16, F1: (fi61 * (fi59 / (((v1 & -972)) | 1)))}
                        if (4294967295 >= 2) {
                            var v64 int32 = 2147483647
                            switch (fi61 >> ((18446744073709551615) & 63)) {
                            case 58:
                                v65 := GenS1{F0: false, F1: (v1 + 128)}
                                _ = v65
                            default:
                                fmt.Println(v64)
                            }
                            _ = v64
                        } else {
                            var v66 uint8 = 127
                            fmt.Println(v17, v66)
                            v67 := GenS6{F0: true, F1: 1, F2: uint16(v66)}
                            _ = v66
                            _ = v67
                        }
                        v17 = (v17 << ((4294967295) & 63))
                        _ = fi61
                        _ = v62
                        _ = v63
                    }
                    _ = v60
                }
            }
            v18 = v18
            if v18 {
                var v68 int64 = -488
                switch "" {
                case "\n\t":
                    v17 -= (v17 & v17)
                default:
                    switch (v12 + -2147483648) {
                    case -2147483648:
                        fmt.Println(v68, v18)
                    case 127:
                        v18 = v16
                    }
                }
                for fi69 := 0; fi69 < 20; fi69++ {
                    _ = len(v13)
                    var v70 uint64 = 42
                    if ((15 <= 2) && (!((1 >= -2147483648) && ((v18 && true) || false)))) {
                        var v71 uint16 = 0
                        v13 = append(v13, (^(v12)))
                        switch fi59 {
                        case -1:
                            var v72 float64 = (0.1 / -2.25)
                            _ = v72
                        case 128:
                            v73 := GenS3{F0: 32767}
                            _ = v73
                        case 42:
                            var v74 int8 = 42
                            _ = v74
                        default:
                            fmt.Println(v17)
                        }
                        _ = v71
                    } else {
                        if v18 {
                            var v75 uint8 = 80
                            _ = v75
                        } else {
                            v76 := GenS4{F0: 167, F1: (42 >= 2147483647)}
                            v1 -= fi59
                            _ = v76
                        }
                    }
                    _ = fi69
                    _ = v70
                }
                _ = v68
            } else {
                if v16 {
                    fmt.Println(v18, v18)
                    for fi77 := 0; fi77 < 2; fi77++ {
                        switch v12 {
                        case 128:
                            v78 := GenS6{F0: (!v18), F1: genFn0(100, (fi59 ^ (-(fi77)))), F2: 32767}
                            _ = v78
                        case 32767:
                            switch ("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))) {
                            case "\n\t":
                                v13 = append(v13, (fi59 ^ fi59))
                            default:
                                var v79 int16 = (^(v17))
                                _ = v79
                            }
                        }
                        v80 := GenS8{F0: 2147483647, F1: 990}
                        var v81 uint32 = 2147483647
                        _ = fi77
                        _ = v80
                        _ = v81
                    }
                    v82 := v3[v1]
                    _ = v82
                } else {
                    fmt.Println(v1, v12)
                    for fi83 := 0; fi83 < 12; fi83++ {
                        if (65535 <= 315) {
                            var v84 bool = true
                            _ = v84
                        }
                        var v85 int8 = 1
                        if true {
                            v16 = (true || v16)
                            v3[fi83] = ("" + "hello")
                        }
                        _ = fi83
                        _ = v85
                    }
                    v86 := GenS5{F0: 974, F1: 2147483648}
                    _ = v86
                }
                var v87 bool = v16
                var v88 float32 = float32(-2.25)
                _ = v87
                _ = v88
            }
            _ = fi59
        }
        fmt.Println(v18)
    }
    v17 = genFn1()
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
