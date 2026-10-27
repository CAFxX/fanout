package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 string; F2 uint32 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 int; F1 uint; F2 int16 }
type GenS8 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS9 struct { F0 int16; F1 uint64 }
type GenS10 struct { F0 int; F1 string; F2 uint8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint64; F1 int64 }
type GenS13 struct { F0 int16; F1 uint32; F2 string }
func genFn0(p0 int) int32 {
    for fi6 := 0; fi6 < 12; fi6++ {
        p0 += fi6
        _ = fi6
    }
    if false {
        v7 := []int{}
        _ = v7
    }
    var v8 uint64 = 4294967295
    for fi9 := 0; fi9 < 8; fi9++ {
        if false {
            switch "a" {
            case "foo bar":
                v8 += 1
            case "":
                v10 := []int{-9223372036854775808, 32768, p0, (-(p0))}
                _ = v10
            }
        } else {
            if (1 == 255) {
                p0 = 9223372036854775807
            }
            if ((0 == 2) || (false || false)) {
                switch "foo bar" {
                case "a":
                    var v11 uint32 = 295
                    _ = v11
                case "\n\t":
                    var v12 uint8 = 0
                    _ = v12
                case "hello":
                    fmt.Println(p0, fi9, fi9)
                }
            } else {
                p0 = p0
            }
            if ((1 > 2147483647) || (!true)) {
                if (0 == -2147483648) {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        fmt.Println(v8)
                        _ = fi13
                    }
                    var v14 string = ((("foo bar" + "a") + "\n\t") + "foo bar")
                    if (1000 > 521) {
                        switch v14 {
                        case "":
                            v15 := []int{fi9, 1}
                            _ = v15
                        case "a":
                            for fi16 := 0; fi16 < 4; fi16++ {
                                var v17 float64 = 3.14159
                                var v18 uint16 = 501
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        }
                    } else {
                        v8 = 796
                    }
                    _ = v14
                } else {
                    switch -999 {
                    case 32768:
                        v19 := GenS1{F0: ((("" + "hello") + ("\n\t" + ("\n\t" + "hello"))) + "foo bar"), F1: (v8 | (v8 | v8)), F2: v8}
                        _ = v19
                    case 9223372036854775807:
                        fmt.Println(fi9, fi9)
                    case -1:
                        for fi20 := 0; fi20 < 19; fi20++ {
                            p0 += fi20
                            _ = fi20
                        }
                    }
                    v21 := GenS1{F0: "a", F1: v8, F2: (^(v8))}
                    _ = v21
                }
                var v22 uint64 = 42
                _ = v22
            } else {
                if (p0 == 1) {
                    var v23 int16 = 42
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v25 := GenS2{F0: ("a" < ((("foo bar" + "") + ("foo bar" + (("hello" + "hello") + "hello"))) + "a")), F1: 255, F2: 0}
                        var v26 int64 = 0
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS3{F0: uint32(v8), F1: v8}
                    _ = v23
                    _ = v27
                } else {
                    var v28 uint64 = (^(v8))
                    v29 := []int{fi9, -727, (p0 | fi9)}
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                    _ = v28
                    _ = v29
                }
                v30 := GenS1{F0: "\n\t", F1: 197, F2: (v8 & v8)}
                _ = v30
            }
        }
        var v31 int32 = 1
        for fi32 := 0; fi32 < 10; fi32++ {
            v31 = v31
            _ = fi32
        }
        _ = fi9
        _ = v31
    }
    _ = p0
    _ = v8
    return 1
}

func main() {
    var v1 bool = (42 == 4294967295)
    var v2 string = "foo bar"
    v3 := make(map[string]bool)
    var v4 uint64 = 2
    var v5 bool = v1
    var v33 float64 = 0.0
    for fi34 := 0; fi34 < 9; fi34++ {
        v1 = (2 == 260)
        if v5 {
            v35 := []int{}
            v2 = v2
            v5 = ((v33 * 1e10) >= v33)
            _ = v35
        }
        v5 = v1
        _ = fi34
    }
    v36 := v3[v2]
    var v37 uint32 = 65535
    for fi38 := 0; fi38 < 4; fi38++ {
        if v5 {
            v37 = 4294967295
        }
        var v39 uint8 = 127
        _ = fi38
        _ = v39
    }
    v5 = true
    v40 := GenS4{F0: "foo bar"}
    var v41 string = "a"
    v36 = false
    for fi42 := 0; fi42 < 13; fi42++ {
        var v43 string = (v2 + v2)
        v44 := []int{2147483647, (-(fi42)), (fi42 ^ fi42)}
        switch 454 {
        case -2147483648:
            var v45 uint64 = (v4 | 366)
            _ = v45
        case 127:
            switch 34 {
            case -9223372036854775808:
                var v46 int8 = -128
                _ = v46
            case -926:
                _ = len(v44)
            }
        case -914:
            switch (v2 + "\n\t") {
            case "\n\t":
                v43 = "a"
            case "":
                for fi47 := 0; fi47 < 3; fi47++ {
                    var v48 uint = 2
                    var v49 uint = v48
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            case "hello":
                if true {
                    v50 := v3[v2]
                    var v51 int16 = 42
                    _ = v50
                    _ = v51
                }
            default:
                v52 := GenS5{F0: false, F1: v43, F2: v37}
                _ = v52
            }
        default:
            var v53 int8 = 42
            _ = v53
        }
        _ = fi42
        _ = v43
        _ = v44
    }
    if ((0 > 42) || (1 >= -511)) {
        v54 := []int{-610}
        _ = v54
    } else {
        var v55 int8 = 0
        _ = v55
    }
    switch 9223372036854775807 {
    case 127:
        switch v41 {
        case "":
            if v36 {
                var v56 float64 = (0.1 * v33)
                _ = v56
            }
        case "hello":
            if ((((1e10 + v33) + (v33 + (v33 * v33))) / (v33 * v33)) <= v33) {
                v57 := []int{128, 26, -9223372036854775808, 4294967295}
                fmt.Println(v5, v2, v41)
                v5 = (true || ((65535 < 42) && (v1 && true)))
                _ = v57
            } else {
                for fi58 := 0; fi58 < 13; fi58++ {
                    switch v2 {
                    case "":
                        switch (fi58 << ((4294967295) & 63)) {
                        case 128:
                            v3[v2] = (v37 != 2147483648)
                        default:
                            if v36 {
                                for fi59 := 0; fi59 < 17; fi59++ {
                                    v60 := GenS3{F0: 42, F1: v4}
                                    v61 := GenS5{F0: false, F1: "", F2: v37}
                                    var v62 uint64 = (v4 >> ((42) & 63))
                                    _ = fi59
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                }
                                var v63 bool = ((!v5) && ((v4 >> ((255) & 63)) <= (^(v4))))
                                _ = v63
                            }
                        }
                    case "\n\t":
                        v64 := GenS6{F0: (fi58 & 0), F1: -28}
                        _ = v64
                    case "foo bar":
                        v33 += (((0.0 - (v33 / ((v33 - v33) + v33))) + float64((0.1 - (0.0 / ((-2.25 / 1e10) - -2.25))))) - 1e10)
                    }
                    var v65 uint64 = (^(v4))
                    for fi66 := 0; fi66 < 15; fi66++ {
                        v36 = (!(!(!((v4 - 1) == 652))))
                        _ = len(v3)
                        _ = fi66
                    }
                    _ = fi58
                    _ = v65
                }
                switch 2147483647 {
                case 719:
                    fmt.Println(v1, v5)
                case -2147483648:
                    switch -251 {
                    case 2147483647:
                        for fi67 := 0; fi67 < 15; fi67++ {
                            if (-32768 == 42) {
                                delete(v3, "a")
                            } else {
                                v4 = 2
                                if true {
                                    if (966 >= genFn0(128)) {
                                        var v68 int8 = -128
                                        switch ((((("hello" + v2) + (v41 + v41)) + v2) + (v2 + (((v41 + (("\n\t" + v41) + ("\n\t" + v2))) + v2) + (v2 + "a")))) + ("\n\t" + "foo bar")) {
                                        case "":
                                            for fi69 := 0; fi69 < 4; fi69++ {
                                                v5 = (-2.25 <= float32(v33))
                                                var v70 uint64 = (v4 | v4)
                                                _ = fi69
                                                _ = v70
                                            }
                                        case "foo bar":
                                            var v71 float64 = v33
                                            _ = v71
                                        case "a":
                                            switch -978 {
                                            case 811:
                                                switch "a" {
                                                case "foo bar":
                                                    for fi72 := 0; fi72 < 7; fi72++ {
                                                        v73 := GenS7{F0: 56, F1: 946, F2: -627}
                                                        v74 := []int{}
                                                        v75 := GenS8{F0: 2, F1: 0, F2: -2147483648}
                                                        _ = fi72
                                                        _ = v73
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                case "a":
                                                    v76 := GenS9{F0: 0, F1: v4}
                                                    _ = v76
                                                }
                                            default:
                                                if v1 {
                                                    var v77 uint = 1
                                                    _ = v77
                                                } else {
                                                    v78 := GenS3{F0: v37, F1: (^(v4))}
                                                    _ = v78
                                                }
                                            }
                                        }
                                        _ = v68
                                    }
                                }
                            }
                            v79 := GenS10{F0: (fi67 ^ (fi67 ^ fi67)), F1: (("a" + v2) + "a"), F2: 51}
                            _ = fi67
                            _ = v79
                        }
                    case -1:
                        var v80 float32 = (-2.25 * ((0.0 * 0.1) * float32(0.0)))
                        _ = v80
                    default:
                        var v81 int16 = -919
                        _ = v81
                    }
                }
                var v82 uint16 = 1
                _ = v82
            }
        case "a":
            var v83 uint = uint(v37)
            _ = v83
        }
    }
    v84 := GenS11{F0: 18446744073709551615}
    if false {
        var v85 int32 = 1
        for fi86 := 0; fi86 < 16; fi86++ {
            fmt.Println(v2, fi86, v37)
            for fi87 := 0; fi87 < 15; fi87++ {
                var v88 uint64 = uint64(v37)
                v85 = -1
                var v89 int8 = -128
                var v90 int32 = genFn0(9223372036854775807)
                _ = fi87
                _ = v88
                _ = v89
                _ = v90
            }
            v33 = 0.1
            _ = fi86
        }
        var v91 string = "foo bar"
        _ = v85
        _ = v91
    } else {
        v92 := []int{}
        v93 := GenS12{F0: 2, F1: 1}
        if v1 {
            switch 128 {
            case 2147483647:
                v33 = v33
            case 0:
                var v94 int16 = 1000
                _ = v94
            case -3:
                v95 := GenS10{F0: 2147483647, F1: v2, F2: 251}
                _ = v95
            default:
                v96 := GenS13{F0: -312, F1: (v37 - v37), F2: v41}
                _ = v96
            }
            var v97 float64 = v33
            v41 = v2
            _ = v97
        }
        _ = v92
        _ = v93
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v84
}
