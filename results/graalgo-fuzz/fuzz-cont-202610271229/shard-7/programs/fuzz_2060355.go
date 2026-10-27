package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 string; F1 uint64; F2 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 int8; F2 uint }
type GenS7 struct { F0 int8; F1 int32; F2 int }
type GenS8 struct { F0 bool; F1 int64; F2 int8 }
type GenS9 struct { F0 uint64; F1 uint8 }
type GenS10 struct { F0 int16; F1 uint64 }
type GenS11 struct { F0 int32; F1 int; F2 uint64 }
type GenS12 struct { F0 uint8; F1 int16; F2 int }

func main() {
    var v1 uint64 = 778
    v2 := make(map[uint32]int)
    var v3 uint64 = 166
    var v4 float64 = 0.1
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            v4 = v4
            _ = fi6
        }
        v7 := GenS1{F0: 338}
        _ = fi5
        _ = v7
    }
    if (!true) {
        var v8 int16 = 32767
        _ = v8
    } else {
        if false {
            var v9 uint64 = 95
            _ = v9
        } else {
            if false {
                if true {
                    fmt.Println(v3, v1, v4)
                    v10 := []int{9223372036854775807}
                    _ = v10
                } else {
                    var v11 int64 = 42
                    _ = v11
                }
            } else {
                if false {
                    if ((float64(1.5) != (v4 - v4)) || false) {
                        var v12 uint8 = 0
                        var v13 int32 = -1
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 2147483648}
                    v1 = v1
                    _ = v14
                } else {
                    if (-2147483648 > -578) {
                        var v15 int32 = 42
                        v4 += v4
                        _ = v15
                    }
                    fmt.Println(v3, v1)
                }
                switch "\n\t" {
                case "":
                    v4 = v4
                case "a":
                    v16 := GenS1{F0: 904}
                    _ = v16
                case "hello":
                    v17 := []int{-1, 32767, 42, 42}
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 1; fi18++ {
                        v1 = (^(v1))
                        _ = fi18
                    }
                }
                var v19 int8 = -1
                _ = v19
            }
            var v20 int64 = -9223372036854775808
            if false {
                if false {
                    v21 := GenS1{F0: 1}
                    _ = v21
                }
            } else {
                v22 := []int{42, 128, 42, 0}
                _ = v22
            }
            _ = v20
        }
        v23 := GenS2{F0: 866}
        _ = v23
    }
    v3 = (v1 >> ((1) & 63))
    if true {
        v24 := GenS3{F0: 52, F1: 328}
        delete(v2, 0)
        _ = v24
    }
    fmt.Println(v1, v4)
    if (42 == 306) {
        switch ("" + "\n\t") {
        case "hello":
            if (((float32((-2.25 - v4)) * (1.5 / (0.0 + float32(3.14159)))) > -2.25) && (255 <= 506)) {
                var v25 uint8 = 2
                v26 := []int{-219}
                _ = v25
                _ = v26
            } else {
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 int = (fi27 % (((fi27 & (-(fi27)))) | 1))
                    var v29 int32 = -2147483648
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                var v30 float32 = 3.14159
                var v31 uint = 0
                _ = v30
                _ = v31
            }
        case "foo bar":
            for fi32 := 0; fi32 < 11; fi32++ {
                fmt.Println(v3, v4)
                v33 := GenS3{F0: 1, F1: (fi32 / ((fi32) | 1))}
                v34 := []int{(fi32 - (fi32 >> ((65535) & 63)))}
                _ = fi32
                _ = v33
                _ = v34
            }
        case "\n\t":
            if false {
                if true {
                    v3 -= 42
                    v3 = v3
                    if (false || (v3 <= (v3 - (^(v1))))) {
                        v35 := GenS4{F0: "hello", F1: v3, F2: ((!(true || true)) && (42 == 2))}
                        v4 = (float64((3.14159 / 1e10)) - v4)
                        v3 -= (^(v1))
                        _ = v35
                    } else {
                        var v36 string = ""
                        _ = v36
                    }
                } else {
                    v37 := []int{9223372036854775807, 458, 32767, 864}
                    if false {
                        fmt.Println(v3)
                        if (!false) {
                            if (true || true) {
                                for fi38 := 0; fi38 < 19; fi38++ {
                                    v39 := v2[1]
                                    var v40 int32 = 0
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                                if len(v37) >= 2 {
                                    v37 = v37[1:len(v37)]
                                }
                                v41 := GenS5{F0: 1}
                                _ = v41
                            }
                            var v42 uint64 = v1
                            var v43 int64 = 215
                            _ = v42
                            _ = v43
                        } else {
                            switch 9223372036854775807 {
                            case -898:
                                var v44 int64 = -9223372036854775808
                                _ = v44
                            }
                        }
                    } else {
                        var v45 int64 = 1
                        for fi46 := 0; fi46 < 16; fi46++ {
                            switch "\n\t" {
                            case "hello":
                                v45 = (v45 * v45)
                            case "foo bar":
                                switch "a" {
                                case "a":
                                    v47 := GenS2{F0: -1}
                                    _ = v47
                                case "":
                                    for fi48 := 0; fi48 < 14; fi48++ {
                                        v49 := GenS6{F0: -697, F1: 88, F2: uint(v1)}
                                        var v50 bool = false
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                default:
                                    if len(v37) > 0 {
                                        v37[0] = 1
                                    }
                                }
                            }
                            if (576 == 2) {
                                for fi51 := 0; fi51 < 12; fi51++ {
                                    switch (^(fi46)) {
                                    case -9223372036854775808:
                                        v4 = v4
                                    }
                                    delete(v2, 42)
                                    var v52 int32 = 2147483647
                                    _ = fi51
                                    _ = v52
                                }
                                var v53 uint32 = 1
                                v4 = (v4 - -2.25)
                                _ = v53
                            } else {
                                if len(v37) > 0 {
                                    v37[0] = 230
                                }
                            }
                            _ = fi46
                        }
                        switch 128 {
                        case 32768:
                            v3 += (v3 | 20)
                        default:
                            var v54 int32 = -896
                            _ = v54
                        }
                        _ = v45
                    }
                    _ = v37
                }
                v55 := GenS4{F0: "\n\t", F1: (v3 * (v3 & v3)), F2: (127 == 529)}
                for fi56 := 0; fi56 < 12; fi56++ {
                    if (2 < 42) {
                        v3 = v3
                    } else {
                        for fi57 := 0; fi57 < 19; fi57++ {
                            if true {
                                switch 128 {
                                case -9223372036854775808:
                                    var v58 float32 = 0.1
                                    _ = v58
                                case -719:
                                    v4 += v4
                                default:
                                    switch "foo bar" {
                                    case "":
                                        v3 = 4294967295
                                    case "a":
                                        _ = len(v2)
                                    default:
                                        v1 -= (v3 ^ 2)
                                    }
                                }
                            }
                            v3 = v1
                            v3 = v3
                            _ = fi57
                        }
                        var v59 uint8 = 42
                        v60 := []int{32767}
                        _ = v59
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 6; fi61++ {
                        v62 := GenS7{F0: 0, F1: -158, F2: fi61}
                        v63 := GenS8{F0: (42 <= 8), F1: 1, F2: 42}
                        var v64 bool = false
                        _ = fi61
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                    delete(v2, 4294967295)
                    _ = fi56
                }
                _ = v55
            } else {
                v65 := GenS9{F0: (v1 | v3), F1: 1}
                v1 -= v1
                v66 := GenS9{F0: 42, F1: 2}
                _ = v65
                _ = v66
            }
        default:
            for fi67 := 0; fi67 < 3; fi67++ {
                var v68 float32 = 1e10
                switch (-(fi67)) {
                case 234:
                    v69 := GenS10{F0: -1, F1: 171}
                    _ = v69
                case -9223372036854775808:
                    fmt.Println(v4)
                }
                for fi70 := 0; fi70 < 7; fi70++ {
                    if (1000 >= 888) {
                        v71 := GenS2{F0: 42}
                        if (127 == -9) {
                            v72 := GenS11{F0: -2147483648, F1: 4294967295, F2: 552}
                            _ = v72
                        }
                        _ = v71
                    } else {
                        switch ("" + "foo bar") {
                        case "hello":
                            fmt.Println(v3, fi67, v68)
                        default:
                            var v73 float64 = float64(float32(v4))
                            _ = v73
                        }
                        v2[2147483647] = (fi67 << ((241) & 63))
                        if (773 >= 42) {
                            var v74 uint16 = 32767
                            v4 -= 3.14159
                            v1 = (v3 / ((v1) | 1))
                            v4 -= v4
                            _ = v74
                        }
                    }
                    var v75 uint8 = 1
                    if (42 == 42) {
                        fmt.Println(v3)
                        var v76 int64 = -9223372036854775808
                        _ = v76
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            v77 := []int{(^(fi67))}
                            if false {
                                var v78 float32 = 1e10
                                var v79 uint = 593
                                _ = v78
                                _ = v79
                            }
                            var v80 int64 = 1
                            _ = v77
                            _ = v80
                        }
                    }
                    _ = fi70
                    _ = v75
                }
                _ = fi67
                _ = v68
            }
        }
    }
    v81 := GenS12{F0: 127, F1: -82, F2: 32768}
    v1 = (v1 | (v1 % ((233) | 1)))
    if false {
        fmt.Println(v3, v4, v4)
    } else {
        for fi82 := 0; fi82 < 18; fi82++ {
            if (((!(410 != -1)) && true) || true) {
                if ("hello" != "") {
                    v1 += 42
                    v83 := []int{128, (^(fi82)), (fi82 | fi82), (^(fi82))}
                    fmt.Println(v3)
                    _ = v83
                } else {
                    v1 = v3
                    switch (("\n\t" + ("" + ((("hello" + ("a" + "foo bar")) + "a") + "foo bar"))) + "hello") {
                    case "foo bar":
                        var v84 uint64 = 1
                        _ = v84
                    case "hello":
                        fmt.Println(fi82)
                    default:
                        var v85 uint8 = 1
                        _ = v85
                    }
                    v4 -= v4
                }
                v4 = v4
            }
            _ = fi82
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v81
}
