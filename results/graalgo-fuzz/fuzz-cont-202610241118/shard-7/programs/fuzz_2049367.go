package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS7 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 bool; F1 uint8; F2 int8 }
func genFn0() int {
    var v7 float32 = 0.1
    if (true || (-32768 != 1000)) {
        var v8 int16 = -1
        v8 = (v8 >> ((65535) & 63))
        _ = v8
    }
    var v9 int = -9223372036854775808
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint8 = 128
        v9 -= (^(v9))
        var v12 int16 = -1
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v7
    _ = v9
    return v9
}
func genFn1() int64 {
    var v13 uint8 = 181
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int64 = -978
        switch fi14 {
        case 9223372036854775807:
            if true {
                for fi16 := 0; fi16 < 20; fi16++ {
                    if ("hello" <= "\n\t") {
                        v17 := []int{fi14, fi16, 9223372036854775807, fi16}
                        var v18 int = genFn0()
                        _ = v17
                        _ = v18
                    }
                    v15 = v15
                    v13 += (v13 ^ v13)
                    v19 := GenS1{F0: (^(v13)), F1: 1, F2: 1000}
                    _ = fi16
                    _ = v19
                }
                for fi20 := 0; fi20 < 3; fi20++ {
                    fmt.Println(v13)
                    v21 := GenS1{F0: 127, F1: 4294967295, F2: -780}
                    _ = fi20
                    _ = v21
                }
            } else {
                var v22 float64 = 0.0
                v23 := GenS1{F0: (v13 % ((127) | 1)), F1: 4294967295, F2: -407}
                var v24 int16 = -1
                _ = v22
                _ = v23
                _ = v24
            }
        default:
            fmt.Println(v13)
        }
        _ = fi14
        _ = v15
    }
    v25 := GenS1{F0: (^(v13)), F1: 500, F2: 326}
    _ = v13
    _ = v25
    return -1
}

func main() {
    var v1 float32 = float32((0.0 / (0.0 - (3.14159 + ((3.14159 + 1.5) + 1e10)))))
    var v2 string = "hello"
    var v3 uint = 0
    var v4 uint = v3
    var v5 uint = (v3 >> (((v4 / ((255) | 1))) & 63))
    v6 := []string{}
    _ = len(v6)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if ((!false) || (1 != 1)) {
        var v26 int16 = 0
        _ = v26
    } else {
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 int16 = -499
            _ = fi27
            _ = v28
        }
    }
    var v29 string = v2
    v2 = v29
    for fi30 := 0; fi30 < 15; fi30++ {
        if true {
            if (!false) {
                if len(v6) > 0 {
                    v6[0] = v2
                }
                switch ("\n\t" + (v2 + "a")) {
                case "hello":
                    fmt.Println(v1, v2)
                default:
                    var v31 uint32 = 2147483647
                    _ = v31
                }
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v2 = v29
                v5 -= 358
            }
        } else {
            var v32 string = (v2 + v29)
            v4 = v5
            v6 = append(v6, (v32 + (v2 + v32)))
            _ = v32
        }
        _ = fi30
    }
    if len(v6) > 0 {
        v6[0] = v29
    }
    v4 -= v5
    v33 := GenS1{F0: 160, F1: 18446744073709551615, F2: -1}
    if false {
        fmt.Println(v4, v2)
        var v34 int16 = -1
        fmt.Println(v29, v5)
        _ = v34
    } else {
        switch (v29 + "a") {
        case "foo bar":
            switch ("" + v2) {
            case "foo bar":
                v4 -= v3
            case "a":
                v35 := GenS2{F0: (v5 >> (((v3 >> ((4294967295) & 63))) & 63)), F1: 42}
                _ = v35
            case "hello":
                var v36 int8 = 106
                _ = v36
            default:
                if len(v6) > 0 {
                    v6[0] = v29
                }
            }
        case "":
            switch -57 {
            case 2147483647:
                v5 -= (v5 * v5)
            case 32768:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case 575:
                if len(v6) > 0 {
                    v6[0] = (v29 + (("" + (("foo bar" + ((v29 + (v2 + v29)) + ("\n\t" + ("a" + v29)))) + v2)) + v2))
                }
            default:
                v37 := GenS3{F0: 4294967295, F1: 664}
                _ = v37
            }
        case "hello":
            if true {
                for fi38 := 0; fi38 < 18; fi38++ {
                    if (v4 > v5) {
                        v39 := GenS2{F0: (^(v4)), F1: 2147483647}
                        _ = v39
                    } else {
                        switch 1 {
                        case -2147483648:
                            fmt.Println(v29)
                        case 2147483647:
                            switch (-(fi38)) {
                            case 9223372036854775807:
                                switch fi38 {
                                case -409:
                                    switch (v29 + "\n\t") {
                                    case "":
                                        switch "" {
                                        case "":
                                            switch ("a" + v29) {
                                            case "\n\t":
                                                v5 += v4
                                            case "hello":
                                                v40 := GenS4{F0: 100, F1: -117}
                                                _ = v40
                                            default:
                                                v2 = "a"
                                            }
                                        case "a":
                                            if false {
                                                switch "\n\t" {
                                                case "hello":
                                                    v41 := GenS2{F0: (^(v5)), F1: 2}
                                                    _ = v41
                                                default:
                                                    _ = len(v6)
                                                }
                                                v42 := GenS3{F0: 1, F1: 2147483647}
                                                _ = v42
                                            } else {
                                                v43 := GenS5{F0: 255, F1: (v3 / ((v3) | 1))}
                                                var v44 string = "a"
                                                _ = v43
                                                _ = v44
                                            }
                                        default:
                                            if ((!true) || true) {
                                                var v45 string = (v2 + ((v2 + ("" + ("hello" + (("a" + ("a" + v2)) + "a")))) + ((v2 + ((v2 + "hello") + (("foo bar" + ("" + v2)) + v29))) + v2)))
                                                if true {
                                                    if (true && (3.14159 != 1e10)) {
                                                        switch fi38 {
                                                        case 32767:
                                                            for fi46 := 0; fi46 < 9; fi46++ {
                                                                var v47 int32 = -232
                                                                if true {
                                                                    var v48 int = 42
                                                                    var v49 uint16 = 2
                                                                    v2 = "foo bar"
                                                                    v48 = v48
                                                                    _ = v48
                                                                    _ = v49
                                                                } else {
                                                                    var v50 uint64 = 4294967295
                                                                    _ = v50
                                                                }
                                                                _ = fi46
                                                                _ = v47
                                                            }
                                                        default:
                                                            v1 = v1
                                                        }
                                                    }
                                                    for fi51 := 0; fi51 < 9; fi51++ {
                                                        if true {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            v52 := GenS6{F0: 218, F1: (false || false), F2: 127}
                                                            v53 := GenS4{F0: 100, F1: -736}
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        var v54 uint16 = 238
                                                        _ = fi51
                                                        _ = v54
                                                    }
                                                    var v55 uint16 = 652
                                                    _ = v55
                                                } else {
                                                    var v56 uint8 = 2
                                                    var v57 int8 = 1
                                                    _ = v56
                                                    _ = v57
                                                }
                                                if false {
                                                    var v58 int8 = -1
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    _ = v58
                                                } else {
                                                    if false {
                                                        v1 = v1
                                                        v5 += (v4 + 76)
                                                    } else {
                                                        for fi59 := 0; fi59 < 13; fi59++ {
                                                            v3 -= v4
                                                            v6 = append(v6, v45)
                                                            if true {
                                                                v60 := GenS2{F0: (v4 | (v3 - v4)), F1: 4294967295}
                                                                v2 = ""
                                                                _ = v60
                                                            } else {
                                                                v4 -= (v5 & 18446744073709551615)
                                                            }
                                                            _ = fi59
                                                        }
                                                    }
                                                    v61 := GenS7{F0: 42, F1: 35, F2: 325}
                                                    switch v45 {
                                                    case "":
                                                        if (false || ((v1 <= v1) || false)) {
                                                            if false {
                                                                v62 := GenS2{F0: (v3 ^ v3), F1: 42}
                                                                _ = v62
                                                            } else {
                                                                v63 := GenS8{F0: fi38, F1: (v5 - v3)}
                                                                _ = v63
                                                            }
                                                            v64 := GenS9{F0: 671}
                                                            v29 = v29
                                                            _ = v64
                                                        } else {
                                                            var v65 bool = (((127 > 255) || (!false)) && false)
                                                            _ = v65
                                                        }
                                                    }
                                                    _ = v61
                                                }
                                                _ = v45
                                            }
                                        }
                                    case "a":
                                        switch (-(fi38)) {
                                        case -2147483648:
                                            v3 += v5
                                        }
                                    case "\n\t":
                                        v66 := GenS5{F0: 2, F1: v5}
                                        _ = v66
                                    default:
                                        if (686 >= 1) {
                                            var v67 string = "hello"
                                            fmt.Println(v67, v29, v4)
                                            if false {
                                                v6 = append(v6, v67)
                                            }
                                            _ = v67
                                        }
                                    }
                                case 127:
                                    if ((0.0 < 3.14159) || false) {
                                        v29 = v29
                                        if true {
                                            var v68 int = fi38
                                            v2 = (("a" + "\n\t") + "hello")
                                            _ = v68
                                        }
                                        switch fi38 {
                                        case 2147483647:
                                            var v69 int = (^(fi38))
                                            _ = v69
                                        case -1:
                                            for fi70 := 0; fi70 < 16; fi70++ {
                                                switch genFn0() {
                                                case -133:
                                                    v1 = (v1 - float32(0.0))
                                                case 986:
                                                    var v71 int = (fi70 | fi70)
                                                    v5 = v4
                                                    var v72 bool = (true || false)
                                                    _ = v71
                                                    _ = v72
                                                case 0:
                                                    fmt.Println(v5, fi70)
                                                }
                                                _ = fi70
                                            }
                                        case 192:
                                            if (4294967295 < 18446744073709551615) {
                                                v1 = (v1 - v1)
                                            }
                                        }
                                    }
                                default:
                                    var v73 int = (fi38 >> (((^(v3))) & 63))
                                    _ = v73
                                }
                            case -9223372036854775808:
                                v74 := GenS2{F0: (v4 / (((v3 * 18446744073709551615)) | 1)), F1: 144}
                                _ = v74
                            }
                        default:
                            v3 = (^(v5))
                        }
                    }
                    v29 = "a"
                    _ = fi38
                }
                if len(v6) > 0 {
                    v6[0] = ((((("" + (v29 + "a")) + (((v29 + "hello") + v29) + "hello")) + "") + "hello") + v29)
                }
                var v75 float64 = (3.14159 / float64(1e10))
                _ = v75
            } else {
                var v76 float64 = (1.5 - 0.0)
                for fi77 := 0; fi77 < 2; fi77++ {
                    _ = len(v6)
                    var v78 int8 = -128
                    fmt.Println(v29, v2, v29)
                    _ = fi77
                    _ = v78
                }
                switch genFn0() {
                case 32768:
                    if false {
                        var v79 uint = v3
                        _ = v79
                    }
                case -9223372036854775808:
                    fmt.Println(v3)
                }
                _ = v76
            }
        }
        fmt.Println(v3, v4, v5)
    }
    var v80 int64 = 0
    fmt.Println(v3, v80, v80)
    for fi81 := 0; fi81 < 15; fi81++ {
        for fi82 := 0; fi82 < 2; fi82++ {
            if (false || ((fi81 - fi82) != fi82)) {
                fmt.Println(v5)
                v29 = v29
                var v83 int64 = (v80 >> (((v5 | v4)) & 63))
                _ = v83
            }
            v2 = "a"
            _ = fi82
        }
        switch v2 {
        case "a":
            for fi84 := 0; fi84 < 5; fi84++ {
                v29 = (v2 + (v2 + "\n\t"))
                _ = fi84
            }
        default:
            v4 = 0
        }
        _ = fi81
    }
    v85 := GenS10{F0: (false && false), F1: 128, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v80
    _ = v85
}
