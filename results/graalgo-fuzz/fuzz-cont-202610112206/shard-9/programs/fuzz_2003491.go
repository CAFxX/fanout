package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32; F1 int8 }
type GenS11 struct { F0 int; F1 uint64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int8; F1 string }
func genFn0(p0 bool) uint16 {
    fmt.Println(p0)
    if (1 >= 127) {
        v4 := GenS1{F0: 0, F1: 127, F2: 112}
        _ = v4
    }
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(p0, p0)
        p0 = (p0 && ((fi5 & (fi5 - (fi5 & fi5))) > 884))
        _ = fi5
    }
    v6 := GenS2{F0: -9223372036854775808, F1: 920, F2: 42}
    _ = p0
    _ = v6
    return 527
}
func genFn1(p0 int32, p1 bool) int8 {
    p1 = ((255 != 4294967295) || ((p1 && ((("foo bar" + ("a" + "a")) <= "a") && (p0 > p0))) && p1))
    var v7 uint64 = 4294967295
    var v8 uint16 = 471
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 int = -284
    v2 := make(map[string]int)
    var v3 uint16 = 1
    var v9 uint = 18446744073709551615
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int8 = genFn1(-2147483648, (true || false))
        v2["foo bar"] = v1
        for fi12 := 0; fi12 < 10; fi12++ {
            v9 -= (v9 + 2)
            var v13 int = (v1 << ((v9) & 63))
            v13 = (v13 & fi10)
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v14 := []int{}
    v15 := GenS1{F0: 0, F1: -1, F2: 128}
    v9 = v9
    v16 := GenS2{F0: v1, F1: -473, F2: 42}
    if (-61 != -398) {
        v9 += (v9 >> (((v9 % ((823) | 1))) & 63))
        if (0.0 < ((1.5 / 0.1) + -2.25)) {
            v3 -= (v3 * v3)
        }
    }
    var v17 int = v1
    switch -1 {
    case -590:
        v9 = (v9 - v9)
    default:
        var v18 float32 = 0.1
        _ = v18
    }
    switch "" {
    case "foo bar":
        v2["\n\t"] = 32768
    case "a":
        v1 -= (v17 ^ (v1 << (((v9 / ((v9) | 1))) & 63)))
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v3 = v3
        _ = fi19
    }
    if true {
        switch -2147483648 {
        case -378:
            if (!((false || ("" < "foo bar")) && (4294967295 != int(v3)))) {
                switch (("" + "foo bar") + ((("\n\t" + ((("" + "a") + (("foo bar" + ("\n\t" + "")) + ("" + "\n\t"))) + "\n\t")) + "a") + "foo bar")) {
                case "foo bar":
                    if false {
                        v20 := GenS2{F0: v17, F1: int16(v17), F2: 112}
                        var v21 string = "hello"
                        switch "" {
                        case "hello":
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        case "\n\t":
                            v9 = v9
                        default:
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        }
                        _ = v20
                        _ = v21
                    }
                case "\n\t":
                    for fi22 := 0; fi22 < 7; fi22++ {
                        v23 := GenS3{F0: (^(v3)), F1: -1}
                        _ = fi22
                        _ = v23
                    }
                case "":
                    if (!false) {
                        v3 = 2
                        for fi24 := 0; fi24 < 17; fi24++ {
                            var v25 string = "hello"
                            var v26 float64 = 0.1
                            fmt.Println(v25, v26, v1)
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                default:
                    var v27 int16 = 1000
                    _ = v27
                }
            }
        case -135:
            switch ("hello" + "hello") {
            case "foo bar":
                var v28 uint32 = 0
                _ = v28
            case "\n\t":
                switch ("" + "") {
                case "":
                    if false {
                        v29 := GenS1{F0: 4294967295, F1: 100, F2: 128}
                        v30 := GenS4{F0: 1, F1: (2147483647 > 734)}
                        _ = v29
                        _ = v30
                    }
                case "hello":
                    if (100 == 127) {
                        var v31 uint32 = 42
                        v1 += -984
                        _ = v31
                    } else {
                        for fi32 := 0; fi32 < 14; fi32++ {
                            var v33 uint8 = 38
                            var v34 int8 = 42
                            v35 := GenS5{F0: (!true)}
                            _ = fi32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    }
                }
            case "a":
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            default:
                var v36 uint8 = 128
                _ = v36
            }
        default:
            v37 := GenS5{F0: true}
            _ = v37
        }
        var v38 float64 = -2.25
        var v39 uint = 2
        _ = v38
        _ = v39
    } else {
        var v40 float32 = 1.5
        _ = v40
    }
    if true {
        if (18446744073709551615 >= 2) {
            v1 = (v1 % ((v1) | 1))
        }
        var v41 float64 = float64(1e10)
        var v42 int16 = -55
        _ = v41
        _ = v42
    } else {
        if (128 < 42) {
            switch "\n\t" {
            case "hello":
                var v43 int = (-(v1))
                _ = v43
            }
            var v44 int64 = -1
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = v44
        }
    }
    if ((!false) && (!true)) {
        for fi45 := 0; fi45 < 2; fi45++ {
            fmt.Println(v3, v1, v3)
            if true {
                fmt.Println(v3)
                var v46 int64 = -322
                _ = v46
            } else {
                for fi47 := 0; fi47 < 20; fi47++ {
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    v48 := GenS5{F0: false}
                    _ = fi47
                    _ = v48
                }
                switch "foo bar" {
                case "":
                    if (122 <= 0) {
                        if false {
                            v9 = (v9 + (v9 ^ v9))
                            if true {
                                v17 = v17
                            } else {
                                v17 -= v17
                                v14 = append(v14, v1)
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            }
                            v9 += (v9 + v9)
                        }
                        switch (v17 / (((v1 | (fi45 - (-(v17))))) | 1)) {
                        case 127:
                            v49 := GenS6{F0: v9}
                            _ = v49
                        }
                        v50 := GenS6{F0: 18446744073709551615}
                        _ = v50
                    }
                }
                v2[""] = (-(v1))
            }
            _ = fi45
        }
        v3 += (^(v3))
        v51 := GenS7{F0: (v3 + 0), F1: v1}
        _ = v51
    } else {
        if ((0 != 1) && false) {
            v3 -= (v3 - (^(v3)))
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            if (-32768 == 32767) {
                for fi52 := 0; fi52 < 19; fi52++ {
                    for fi53 := 0; fi53 < 4; fi53++ {
                        fmt.Println(fi53, v17)
                        v3 -= v3
                        _ = fi53
                    }
                    if false {
                        var v54 uint8 = 127
                        if (uint32(v9) >= 0) {
                            if false {
                                var v55 uint8 = v54
                                v2["hello"] = 621
                                _ = v55
                            } else {
                                v56 := GenS5{F0: (9223372036854775807 == v1)}
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                _ = v56
                            }
                        }
                        _ = v54
                    }
                    _ = fi52
                }
                if (42 < 636) {
                    v17 += (^(v17))
                    if true {
                        switch (("a" + ("hello" + "hello")) + (("\n\t" + ((("hello" + (("a" + "foo bar") + "hello")) + "foo bar") + ("\n\t" + "foo bar"))) + "foo bar")) {
                        case "foo bar":
                            if false {
                                if (127 >= 131) {
                                    v1 -= (v17 - (v17 & (v17 + v17)))
                                }
                                v1 += (-(v17))
                            } else {
                                for fi57 := 0; fi57 < 12; fi57++ {
                                    v1 += 127
                                    var v58 uint16 = (v3 & v3)
                                    _ = fi57
                                    _ = v58
                                }
                            }
                        case "hello":
                            delete(v2, "foo bar")
                        }
                        if (!false) {
                            for fi59 := 0; fi59 < 16; fi59++ {
                                v1 = fi59
                                for fi60 := 0; fi60 < 10; fi60++ {
                                    v61 := v2["hello"]
                                    fmt.Println(v1, v1)
                                    _ = fi60
                                    _ = v61
                                }
                                v62 := GenS8{F0: 42, F1: genFn1(968, false)}
                                _ = fi59
                                _ = v62
                            }
                        } else {
                            v3 += (^(v3))
                            v2["foo bar"] = v1
                        }
                    } else {
                        var v63 uint8 = 2
                        v17 = (^(v1))
                        _ = v63
                    }
                    if (!true) {
                        var v64 uint = (^(v9))
                        _ = v64
                    } else {
                        _ = len(v14)
                        if true {
                            for fi65 := 0; fi65 < 2; fi65++ {
                                v66 := GenS8{F0: 1000, F1: -99}
                                _ = fi65
                                _ = v66
                            }
                        }
                    }
                }
            }
        }
        switch ("" + "hello") {
        case "a":
            v17 = -2147483648
        case "":
            if true {
                if (!(-9223372036854775808 != 9223372036854775807)) {
                    for fi67 := 0; fi67 < 7; fi67++ {
                        v68 := GenS6{F0: v9}
                        var v69 string = (("\n\t" + "") + "\n\t")
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                    switch v17 {
                    case 128:
                        var v70 int32 = -2147483648
                        _ = v70
                    case 4:
                        v71 := v2["foo bar"]
                        _ = v71
                    }
                }
                if true {
                    var v72 bool = (!true)
                    if (false && v72) {
                        v14 = append(v14, 42)
                    } else {
                        v73 := GenS9{F0: 1}
                        switch "a" {
                        case "a":
                            var v74 float32 = 0.0
                            _ = v74
                        case "foo bar":
                            if v72 {
                                for fi75 := 0; fi75 < 8; fi75++ {
                                    _ = len(v2)
                                    _ = fi75
                                }
                                v76 := GenS5{F0: (1.5 >= (float32((3.14159 + (float64(1.5) - (3.14159 - 0.1)))) + 3.14159))}
                                if v72 {
                                    var v77 string = ""
                                    var v78 string = v77
                                    _ = v77
                                    _ = v78
                                } else {
                                    var v79 string = ((("" + ("hello" + (("\n\t" + (("" + "\n\t") + "a")) + "hello"))) + "a") + ("a" + "foo bar"))
                                    var v80 int64 = 9223372036854775807
                                    _ = v79
                                    _ = v80
                                }
                                _ = v76
                            } else {
                                if (false && v72) {
                                    v81 := v2[("foo bar" + "foo bar")]
                                    for fi82 := 0; fi82 < 11; fi82++ {
                                        var v83 float32 = 1e10
                                        v84 := GenS2{F0: -802, F1: 430, F2: 1}
                                        v85 := GenS6{F0: v9}
                                        _ = fi82
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    }
                                    _ = v81
                                }
                            }
                        case "\n\t":
                            v3 = (v3 - (v3 % (((v3 | v3)) | 1)))
                        default:
                            fmt.Println(v1)
                        }
                        _ = v73
                    }
                    v72 = false
                    _ = v72
                }
                for fi86 := 0; fi86 < 8; fi86++ {
                    _ = len(v2)
                    if true {
                        var v87 uint32 = 2147483647
                        _ = v87
                    } else {
                        v88 := GenS10{F0: 2147483648, F1: genFn1(-1, true)}
                        _ = v88
                    }
                    if true {
                        fmt.Println(v9)
                        _ = len(v14)
                        var v89 uint8 = 33
                        _ = v89
                    } else {
                        var v90 int32 = 104
                        fmt.Println(v17)
                        var v91 int32 = (v90 & (v90 + (v90 ^ (v90 * v90))))
                        _ = v90
                        _ = v91
                    }
                    _ = fi86
                }
            } else {
                for fi92 := 0; fi92 < 2; fi92++ {
                    var v93 int = -488
                    var v94 int32 = 686
                    _ = fi92
                    _ = v93
                    _ = v94
                }
                if true {
                    for fi95 := 0; fi95 < 16; fi95++ {
                        var v96 int8 = 0
                        v97 := GenS11{F0: (v17 - 42), F1: 86}
                        v98 := GenS12{F0: (^(fi95))}
                        _ = fi95
                        _ = v96
                        _ = v97
                        _ = v98
                    }
                    delete(v2, ("\n\t" + "a"))
                    switch v17 {
                    case 2147483647:
                        switch (-(v17)) {
                        case 32767:
                            for fi99 := 0; fi99 < 8; fi99++ {
                                var v100 uint64 = 611
                                v14 = append(v14, -1)
                                v100 -= (v100 >> ((65535) & 63))
                                _ = fi99
                                _ = v100
                            }
                        default:
                            v101 := GenS13{F0: 127, F1: "foo bar"}
                            _ = v101
                        }
                    case 322:
                        if (!false) {
                            var v102 float32 = -2.25
                            _ = v102
                        }
                    }
                }
                var v103 string = "\n\t"
                _ = v103
            }
        case "hello":
            if true {
                v9 = v9
                switch "\n\t" {
                case "hello":
                    v104 := GenS1{F0: 636, F1: 127, F2: 157}
                    _ = v104
                }
                v1 -= (v17 / ((1) | 1))
            } else {
                var v105 float32 = (-2.25 + 1.5)
                v106 := GenS6{F0: (v9 ^ 255)}
                _ = v105
                _ = v106
            }
        default:
            var v107 string = "foo bar"
            _ = v107
        }
        v108 := GenS9{F0: -126}
        _ = v108
    }
    _ = len(v14)
    switch int(v3) {
    case 1:
        for fi109 := 0; fi109 < 5; fi109++ {
            v9 = v9
            var v110 int64 = -179
            _ = fi109
            _ = v110
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
