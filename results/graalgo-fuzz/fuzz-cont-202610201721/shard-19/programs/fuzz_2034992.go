package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 bool; F2 uint32 }
type GenS3 struct { F0 int; F1 int; F2 uint }
type GenS4 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint64; F1 uint16 }
type GenS12 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS13 struct { F0 uint32; F1 int32 }
type GenS14 struct { F0 uint32; F1 int32; F2 int64 }
type GenS15 struct { F0 int32; F1 string; F2 uint32 }
type GenS16 struct { F0 int; F1 int16 }
type GenS17 struct { F0 uint32; F1 int; F2 string }
type GenS18 struct { F0 int; F1 uint }
type GenS19 struct { F0 int32; F1 uint16 }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 uint32 }
type GenS22 struct { F0 bool; F1 int64; F2 uint64 }
type GenS23 struct { F0 string; F1 bool; F2 int32 }
func genFn0() int {
    if false {
        switch "foo bar" {
        case "":
            for fi6 := 0; fi6 < 18; fi6++ {
                var v7 uint64 = 18446744073709551615
                var v8 int64 = 9223372036854775807
                if true {
                    var v9 bool = ((1.5 > float64((1.5 + -2.25))) && true)
                    _ = v9
                }
                _ = fi6
                _ = v7
                _ = v8
            }
        case "hello":
            var v10 uint16 = 533
            _ = v10
        case "foo bar":
            if false {
                v11 := GenS1{F0: -1}
                var v12 int8 = -39
                _ = v11
                _ = v12
            }
        default:
            if (true && (false && ((-567 != 503) || ((!false) || (!false))))) {
                if true {
                    var v13 int64 = -9223372036854775808
                    var v14 uint64 = 0
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS1{F0: 100}
                    _ = v15
                }
                v16 := GenS1{F0: 0}
                var v17 uint64 = 804
                _ = v16
                _ = v17
            } else {
                v18 := []int{-1, 668}
                _ = v18
            }
        }
        var v19 int16 = -32768
        v20 := []int{375, int(v19)}
        _ = v19
        _ = v20
    }
    v21 := []int{-9223372036854775808, -411, 128}
    _ = v21
    return 9223372036854775807
}
func genFn1(p0 uint) uint16 {
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := []int{}
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        _ = fi22
        _ = v23
    }
    var v24 int64 = -9223372036854775808
    v25 := GenS2{F0: v24, F1: true, F2: 2147483648}
    v26 := GenS3{F0: 4294967295, F1: genFn0(), F2: (p0 - p0)}
    var v27 float32 = (0.0 - 3.14159)
    _ = p0
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    return 65535
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    v3 := []string{"\n\t", "\n\t"}
    var v4 bool = (true && false)
    var v5 string = "a"
    v4 = true
    fmt.Println(v4, v4)
    for fi28 := 0; fi28 < 14; fi28++ {
        for fi29 := 0; fi29 < 14; fi29++ {
            v5 = "hello"
            v30 := GenS1{F0: 108}
            _ = fi29
            _ = v30
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi28
    }
    switch 128 {
    case 4294967295:
        for fi31 := 0; fi31 < 5; fi31++ {
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 int8 = 127
                if v4 {
                    v33 = v33
                    v4 = v4
                    if v4 {
                        switch (^(fi32)) {
                        case 1:
                            var v34 uint = 471
                            _ = v34
                        case 127:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            v35 := GenS4{F0: 42, F1: 0, F2: 0}
                            _ = v35
                        }
                        switch genFn0() {
                        case 42:
                            var v36 int32 = 2147483647
                            _ = v36
                        default:
                            if v4 {
                                v37 := v2[2147483647]
                                var v38 int = 1
                                switch (v38 + fi32) {
                                case 127:
                                    var v39 int32 = 496
                                    _ = v39
                                case 128:
                                    switch v5 {
                                    case "foo bar":
                                        switch (v37 + "") {
                                        case "\n\t":
                                            switch v37 {
                                            case "a":
                                                v3 = append(v3, "foo bar")
                                            case "\n\t":
                                                if len(v3) > 0 {
                                                    v3[0] = v37
                                                }
                                            case "foo bar":
                                                v37 = ""
                                            default:
                                                fmt.Println(v37)
                                            }
                                        case "hello":
                                            if (32767 == 4294967295) {
                                                v40 := GenS5{F0: 2147483647}
                                                _ = v40
                                            } else {
                                                v2[0] = "foo bar"
                                                switch (fi31 + fi32) {
                                                case 0:
                                                    var v41 int32 = 42
                                                    _ = v41
                                                case 42:
                                                    v42 := GenS1{F0: -1}
                                                    _ = v42
                                                case 128:
                                                    v37 = (((v37 + ((v37 + (v5 + (("a" + (v37 + v5)) + v37))) + v37)) + (v5 + (v37 + v5))) + v5)
                                                default:
                                                    if v4 {
                                                        if v4 {
                                                            if true {
                                                                v4 = ((((0.0 - -2.25) - 3.14159) + 1.5) < ((3.14159 - 1e10) * 1e10))
                                                            }
                                                            if false {
                                                                var v43 string = ""
                                                                v4 = v4
                                                                _ = v43
                                                            }
                                                            if true {
                                                                v44 := GenS3{F0: (v38 | genFn0()), F1: -980, F2: 42}
                                                                v33 -= 1
                                                                _ = v44
                                                            } else {
                                                                _ = len(v3)
                                                                switch v5 {
                                                                case "":
                                                                    v45 := GenS5{F0: (fi31 + fi32)}
                                                                    _ = v45
                                                                case "foo bar":
                                                                    var v46 int8 = (-(v33))
                                                                    _ = v46
                                                                }
                                                                if true {
                                                                    fmt.Println(v37, v4, v38)
                                                                } else {
                                                                    if v4 {
                                                                        v38 -= v38
                                                                        _ = len(v2)
                                                                    } else {
                                                                        v47 := GenS3{F0: v38, F1: -9223372036854775808, F2: 18446744073709551615}
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (v5 + "\n\t")
                                                                        }
                                                                        v4 = ((652 <= 255) && true)
                                                                        _ = v47
                                                                    }
                                                                }
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = v5
                                                            }
                                                            var v48 int64 = -9223372036854775808
                                                            _ = v48
                                                        }
                                                        v49 := v1[v5]
                                                        if true {
                                                            v50 := GenS5{F0: 2147483647}
                                                            _ = v50
                                                        } else {
                                                            if (623 == 42) {
                                                                if (!v4) {
                                                                    if v4 {
                                                                        v51 := GenS5{F0: (fi31 ^ genFn0())}
                                                                        _ = v51
                                                                    }
                                                                } else {
                                                                    switch (v37 + ("\n\t" + "foo bar")) {
                                                                    case "\n\t":
                                                                        if v4 {
                                                                            var v52 uint8 = 128
                                                                            if false {
                                                                                var v53 bool = false
                                                                                _ = v53
                                                                            } else {
                                                                                v54 := GenS6{F0: -1, F1: 4294967295}
                                                                                var v55 uint64 = 1
                                                                                _ = v54
                                                                                _ = v55
                                                                            }
                                                                            v3 = append(v3, "a")
                                                                            _ = v52
                                                                        }
                                                                    }
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                switch v5 {
                                                                case "":
                                                                    if false {
                                                                        switch v5 {
                                                                        case "foo bar":
                                                                            switch genFn0() {
                                                                            case -9223372036854775808:
                                                                                fmt.Println(v5)
                                                                            case 0:
                                                                                if v4 {
                                                                                    var v56 bool = true
                                                                                    _ = v56
                                                                                }
                                                                            }
                                                                        case "hello":
                                                                            if v4 {
                                                                                v57 := GenS7{F0: -32768, F1: 2}
                                                                                switch "a" {
                                                                                case "":
                                                                                    if (v5 >= v5) {
                                                                                        v58 := GenS3{F0: 32767, F1: (fi31 - (fi31 | 2147483647)), F2: 610}
                                                                                        _ = v58
                                                                                    }
                                                                                case "a":
                                                                                    var v59 uint16 = 42
                                                                                    _ = v59
                                                                                default:
                                                                                    if (v49 <= v49) {
                                                                                        v49 = v49
                                                                                    } else {
                                                                                        var v60 string = "a"
                                                                                        var v61 uint16 = 0
                                                                                        var v62 string = (("foo bar" + (v5 + v5)) + ((v60 + "a") + v60))
                                                                                        _ = v60
                                                                                        _ = v61
                                                                                        _ = v62
                                                                                    }
                                                                                }
                                                                                v63 := GenS6{F0: 42, F1: (v49 | 42)}
                                                                                _ = v57
                                                                                _ = v63
                                                                            }
                                                                        case "":
                                                                            var v64 int64 = -9223372036854775808
                                                                            fmt.Println(v4, v49, v33)
                                                                            fmt.Println(v4)
                                                                            _ = v64
                                                                        }
                                                                        var v65 uint = 950
                                                                        v66 := GenS8{F0: (^(v49))}
                                                                        _ = v65
                                                                        _ = v66
                                                                    } else {
                                                                        if (255 != 65535) {
                                                                            v67 := GenS3{F0: (fi32 + genFn0()), F1: genFn0(), F2: 515}
                                                                            _ = v67
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    if v4 {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = v5
                                                                        }
                                                                        v5 = ((v5 + v37) + v37)
                                                                        v68 := v2[v49]
                                                                        _ = v68
                                                                    }
                                                                case "foo bar":
                                                                    if (v4 && false) {
                                                                        if (!v4) {
                                                                            fmt.Println(fi32, v33)
                                                                            v33 += v33
                                                                            fmt.Println(fi31, fi32, v37)
                                                                        } else {
                                                                            var v69 uint8 = 255
                                                                            if true {
                                                                                v5 = v37
                                                                                var v70 uint16 = 255
                                                                                _ = v70
                                                                            } else {
                                                                                v71 := GenS6{F0: 42, F1: 1}
                                                                                v72 := GenS3{F0: (v38 & (v38 % ((fi31) | 1))), F1: (v38 >> ((4294967295) & 63)), F2: 575}
                                                                                if (!(true && v4)) {
                                                                                    var v73 uint32 = v49
                                                                                    v5 = ("hello" + v5)
                                                                                    _ = v73
                                                                                }
                                                                                var v74 uint = 388
                                                                                _ = v71
                                                                                _ = v72
                                                                                _ = v74
                                                                            }
                                                                            v75 := GenS3{F0: fi32, F1: genFn0(), F2: 18446744073709551615}
                                                                            _ = v69
                                                                            _ = v75
                                                                        }
                                                                        switch v5 {
                                                                        case "":
                                                                            switch v5 {
                                                                            case "":
                                                                                v37 = ("foo bar" + "\n\t")
                                                                            case "\n\t":
                                                                                if (-1 == 1) {
                                                                                    delete(v2, v49)
                                                                                    var v76 int = -1
                                                                                    _ = v76
                                                                                }
                                                                                if (0.0 > 1e10) {
                                                                                    if (!v4) {
                                                                                        fmt.Println(v37)
                                                                                    } else {
                                                                                        v77 := GenS9{F0: 42, F1: (-873 > -988)}
                                                                                        var v78 string = v5
                                                                                        var v79 int = v38
                                                                                        _ = v77
                                                                                        _ = v78
                                                                                        _ = v79
                                                                                    }
                                                                                    var v80 int32 = 896
                                                                                    v81 := GenS10{F0: 851}
                                                                                    v82 := GenS11{F0: 4294967295, F1: genFn1(4294967295)}
                                                                                    _ = v80
                                                                                    _ = v81
                                                                                    _ = v82
                                                                                } else {
                                                                                    var v83 int16 = int16(v33)
                                                                                    v33 -= (v33 % (((v33 % (((^(v33))) | 1))) | 1))
                                                                                    _ = v83
                                                                                }
                                                                                v2[v49] = v5
                                                                                if v4 {
                                                                                    v84 := GenS12{F0: (v49 % (((^(v49))) | 1)), F1: 11, F2: -1}
                                                                                    switch (fi31 % (((fi31 - (fi31 << ((0) & 63)))) | 1)) {
                                                                                    case 1:
                                                                                        var v85 uint32 = (v49 << ((18446744073709551615) & 63))
                                                                                        _ = v85
                                                                                    case 42:
                                                                                        if (((true || v4) && v4) || (!(-2147483648 <= -2147483648))) {
                                                                                            if false {
                                                                                                switch ((v37 + (v37 + "\n\t")) + v5) {
                                                                                                case "":
                                                                                                    if (18446744073709551615 > 1) {
                                                                                                        var v86 uint = 541
                                                                                                        switch genFn0() {
                                                                                                        case 128:
                                                                                                            if (!true) {
                                                                                                                delete(v2, v49)
                                                                                                                var v87 uint32 = v49
                                                                                                                _ = v87
                                                                                                            }
                                                                                                        case -1:
                                                                                                            var v88 bool = (!(v86 >= 643))
                                                                                                            _ = v88
                                                                                                        case 1:
                                                                                                            var v89 uint8 = 186
                                                                                                            _ = v89
                                                                                                        }
                                                                                                        _ = v86
                                                                                                    }
                                                                                                case "hello":
                                                                                                    var v90 float32 = 1.5
                                                                                                    _ = v90
                                                                                                case "\n\t":
                                                                                                    v91 := GenS5{F0: v38}
                                                                                                    _ = len(v3)
                                                                                                    if (!v4) {
                                                                                                        switch ("a" + v37) {
                                                                                                        case "a":
                                                                                                            var v92 uint32 = (v49 >> ((161) & 63))
                                                                                                            _ = v92
                                                                                                        case "":
                                                                                                            v93 := GenS13{F0: 841, F1: 2147483647}
                                                                                                            _ = v93
                                                                                                        }
                                                                                                        v94 := GenS14{F0: (v49 / ((v49) | 1)), F1: 0, F2: 42}
                                                                                                        _ = v94
                                                                                                    } else {
                                                                                                        if v4 {
                                                                                                            fmt.Println(fi31, fi32)
                                                                                                            if v4 {
                                                                                                                var v95 float64 = 0.1
                                                                                                                fmt.Println(v37, v37, v37)
                                                                                                                _ = v95
                                                                                                            } else {
                                                                                                                if v4 {
                                                                                                                    v49 = v49
                                                                                                                    v49 -= v49
                                                                                                                    v37 = (v5 + (v5 + "foo bar"))
                                                                                                                    var v96 string = v5
                                                                                                                    v5 = "foo bar"
                                                                                                                    switch ("hello" + "") {
                                                                                                                    case "a":
                                                                                                                        var v97 int8 = v33
                                                                                                                        _ = v97
                                                                                                                    }
                                                                                                                    v33 = (v33 ^ 127)
                                                                                                                    _ = v96
                                                                                                                } else {
                                                                                                                    var v98 int16 = -487
                                                                                                                    _ = v98
                                                                                                                }
                                                                                                            }
                                                                                                            var v99 string = "foo bar"
                                                                                                            _ = len(v1)
                                                                                                            if (false || v4) {
                                                                                                                var v100 int = fi31
                                                                                                                switch "foo bar" {
                                                                                                                case "a":
                                                                                                                    if v4 {
                                                                                                                        fmt.Println(v100)
                                                                                                                        v2[2147483648] = (v37 + ((v5 + ((v99 + (v5 + "")) + v99)) + (("hello" + v99) + "foo bar")))
                                                                                                                        v99 = v37
                                                                                                                    } else {
                                                                                                                        v4 = v4
                                                                                                                        v49 += v49
                                                                                                                        v99 = v5
                                                                                                                    }
                                                                                                                default:
                                                                                                                    var v101 float64 = float64(0.0)
                                                                                                                    _ = v101
                                                                                                                }
                                                                                                                if (!(42 >= -2147483648)) {
                                                                                                                    if len(v3) > 0 {
                                                                                                                        v3[0] = ""
                                                                                                                    }
                                                                                                                    switch genFn0() {
                                                                                                                    case 9223372036854775807:
                                                                                                                        if v4 {
                                                                                                                            var v102 uint64 = 815
                                                                                                                            _ = v102
                                                                                                                        }
                                                                                                                    case 42:
                                                                                                                        if true {
                                                                                                                            var v103 uint = 18446744073709551615
                                                                                                                            _ = v103
                                                                                                                        }
                                                                                                                    case 32768:
                                                                                                                        if v4 {
                                                                                                                            fmt.Println(fi31, v99, fi31)
                                                                                                                            fmt.Println(fi32)
                                                                                                                            fmt.Println(fi32, v33)
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    switch v5 {
                                                                                                                    case "hello":
                                                                                                                        var v104 uint32 = v49
                                                                                                                        _ = v104
                                                                                                                    case "\n\t":
                                                                                                                        v100 = 452
                                                                                                                    }
                                                                                                                    fmt.Println(v33, v100)
                                                                                                                }
                                                                                                                if false {
                                                                                                                    var v105 uint32 = v49
                                                                                                                    var v106 int64 = -9223372036854775808
                                                                                                                    _ = v105
                                                                                                                    _ = v106
                                                                                                                } else {
                                                                                                                    if (v49 == (^(v49))) {
                                                                                                                        _ = len(v2)
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v100
                                                                                                            } else {
                                                                                                                if v4 {
                                                                                                                    var v107 int = (fi32 % (((fi32 % (((-(fi32))) | 1))) | 1))
                                                                                                                    if v4 {
                                                                                                                        if len(v3) >= 2 {
                                                                                                                            v3 = v3[1:len(v3)]
                                                                                                                        }
                                                                                                                        v3 = append(v3, "hello")
                                                                                                                        v33 = v33
                                                                                                                    } else {
                                                                                                                        v107 += -9223372036854775808
                                                                                                                    }
                                                                                                                    if len(v3) > 0 {
                                                                                                                        v3[0] = "hello"
                                                                                                                    }
                                                                                                                    _ = v107
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v99
                                                                                                        } else {
                                                                                                            v108 := GenS15{F0: 42, F1: "", F2: v49}
                                                                                                            _ = v108
                                                                                                        }
                                                                                                        var v109 float32 = (0.0 - (1e10 * 3.14159))
                                                                                                        v49 -= v49
                                                                                                        _ = v109
                                                                                                    }
                                                                                                    _ = v91
                                                                                                default:
                                                                                                    _ = len(v3)
                                                                                                }
                                                                                                fmt.Println(fi32, fi32)
                                                                                                if (42 != 0) {
                                                                                                    v110 := GenS15{F0: 42, F1: "foo bar", F2: v49}
                                                                                                    switch v37 {
                                                                                                    case "hello":
                                                                                                        var v111 uint8 = uint8(v49)
                                                                                                        _ = v111
                                                                                                    case "":
                                                                                                        if (-2.25 == ((0.1 - 0.1) * 1.5)) {
                                                                                                            var v112 float32 = 1.5
                                                                                                            _ = v112
                                                                                                        } else {
                                                                                                            fmt.Println(v5, v49)
                                                                                                            v5 = v37
                                                                                                            v113 := GenS16{F0: v38, F1: 0}
                                                                                                            _ = v113
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        delete(v2, v49)
                                                                                                    }
                                                                                                    var v114 uint16 = 2
                                                                                                    var v115 string = (((("foo bar" + ((v37 + v5) + ((v37 + v5) + "a"))) + v5) + v37) + "")
                                                                                                    v37 = v115
                                                                                                    var v116 int = genFn0()
                                                                                                    _ = v110
                                                                                                    _ = v114
                                                                                                    _ = v115
                                                                                                    _ = v116
                                                                                                }
                                                                                                v2[(v49 / ((2) | 1))] = "foo bar"
                                                                                            }
                                                                                        } else {
                                                                                            v117 := GenS17{F0: (v49 & 65535), F1: 1, F2: "\n\t"}
                                                                                            _ = v117
                                                                                        }
                                                                                    }
                                                                                    _ = v84
                                                                                } else {
                                                                                    switch (v37 + v5) {
                                                                                    case "\n\t":
                                                                                        v3 = append(v3, v37)
                                                                                    default:
                                                                                        var v118 uint64 = 867
                                                                                        _ = v118
                                                                                    }
                                                                                    fmt.Println(fi32)
                                                                                    v5 = (v37 + "")
                                                                                }
                                                                                var v119 int32 = 0
                                                                                if true {
                                                                                    if true {
                                                                                        v119 -= -1
                                                                                        var v120 int32 = (v119 ^ -2147483648)
                                                                                        var v121 uint16 = 1
                                                                                        _ = v120
                                                                                        _ = v121
                                                                                    } else {
                                                                                        v33 = (^(v33))
                                                                                    }
                                                                                }
                                                                                _ = v119
                                                                            case "hello":
                                                                                v49 = (^(v49))
                                                                            }
                                                                        case "\n\t":
                                                                            switch (("a" + (v5 + v37)) + ((v37 + "hello") + v5)) {
                                                                            case "\n\t":
                                                                                if v4 {
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = v5
                                                                                    }
                                                                                    v38 = genFn0()
                                                                                    v122 := GenS18{F0: 32767, F1: 42}
                                                                                    var v123 int8 = (v33 * v33)
                                                                                    _ = v122
                                                                                    _ = v123
                                                                                }
                                                                            case "":
                                                                                switch 2147483647 {
                                                                                case -412:
                                                                                    var v124 uint64 = 866
                                                                                    _ = v124
                                                                                case -398:
                                                                                    if (!v4) {
                                                                                        v37 = (("foo bar" + ("foo bar" + "a")) + v37)
                                                                                        fmt.Println(v4, v49)
                                                                                    }
                                                                                }
                                                                            default:
                                                                                fmt.Println(fi31, v5)
                                                                            }
                                                                        case "foo bar":
                                                                            var v125 int64 = -1
                                                                            _ = v125
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v4, v37, fi32)
                                                                        fmt.Println(v49)
                                                                    }
                                                                default:
                                                                    v37 = v5
                                                                }
                                                            }
                                                        }
                                                        v49 += 253
                                                        v126 := GenS16{F0: fi32, F1: -1}
                                                        _ = v126
                                                        _ = v49
                                                    }
                                                }
                                                fmt.Println(v5, fi31, v33)
                                            }
                                        case "":
                                            var v127 int = (fi32 * (fi31 % ((fi31) | 1)))
                                            _ = v127
                                        }
                                    }
                                case 32767:
                                    v128 := GenS13{F0: 2, F1: 0}
                                    _ = v128
                                default:
                                    fmt.Println(v37, fi32)
                                }
                                _ = v37
                                _ = v38
                            } else {
                                var v129 string = ("hello" + v5)
                                _ = v129
                            }
                        }
                    }
                } else {
                    fmt.Println(v4)
                }
                _ = fi32
                _ = v33
            }
            _ = len(v3)
            var v130 int32 = -1
            _ = fi31
            _ = v130
        }
    case 42:
        if (0.1 >= -2.25) {
            v5 = (((v5 + v5) + ((((v5 + (v5 + (((("a" + v5) + (((((((("\n\t" + v5) + "\n\t") + (v5 + ((v5 + (v5 + (("hello" + ((((v5 + "hello") + v5) + v5) + v5)) + (v5 + v5)))) + (v5 + v5)))) + v5) + v5) + "a") + v5) + "hello")) + "a") + v5))) + v5) + v5) + (v5 + v5))) + "")
        } else {
            if (0 < -76) {
                for fi131 := 0; fi131 < 1; fi131++ {
                    v132 := GenS19{F0: 447, F1: genFn1(320)}
                    _ = fi131
                    _ = v132
                }
            } else {
                var v133 uint = 4294967295
                var v134 float32 = -2.25
                _ = v133
                _ = v134
            }
            fmt.Println(v5, v4)
            if v4 {
                fmt.Println(v4, v4)
            }
        }
    }
    for fi135 := 0; fi135 < 20; fi135++ {
        for fi136 := 0; fi136 < 5; fi136++ {
            delete(v1, "\n\t")
            _ = fi136
        }
        v137 := GenS15{F0: 20, F1: ("\n\t" + v5), F2: 238}
        _ = fi135
        _ = v137
    }
    if v4 {
        var v138 float64 = 1.5
        _ = v138
    } else {
        var v139 uint8 = 127
        switch v5 {
        case "":
            v140 := GenS15{F0: 0, F1: ((v5 + v5) + (((v5 + "hello") + ("hello" + v5)) + v5)), F2: 1}
            _ = v140
        case "hello":
            if (4294967295 > 371) {
                var v141 int64 = -1
                for fi142 := 0; fi142 < 11; fi142++ {
                    fmt.Println(v4, v139)
                    for fi143 := 0; fi143 < 7; fi143++ {
                        if (v5 > ((v5 + v5) + (("hello" + "\n\t") + v5))) {
                            fmt.Println(v5, v4)
                            fmt.Println(v5, v141)
                            v1[v5] = 2147483648
                        } else {
                            var v144 int8 = -128
                            _ = v144
                        }
                        delete(v2, 322)
                        _ = fi143
                    }
                    v139 = 73
                    _ = fi142
                }
                if v4 {
                    v141 = v141
                    v145 := GenS20{F0: 0}
                    _ = v145
                } else {
                    _ = len(v2)
                }
                _ = v141
            }
        case "a":
            for fi146 := 0; fi146 < 11; fi146++ {
                v4 = (!(72 <= 419))
                v139 -= (^(v139))
                _ = fi146
            }
        }
        var v147 uint64 = 2
        _ = v139
        _ = v147
    }
    var v148 float64 = 1e10
    v148 -= float64(0.1)
    switch -2147483648 {
    case -533:
        v149 := GenS21{F0: 324}
        _ = v149
    default:
        for fi150 := 0; fi150 < 9; fi150++ {
            switch v5 {
            case "foo bar":
                if (4294967295 < 1) {
                    fmt.Println(v5, v4, v5)
                    var v151 uint32 = 2
                    v2[v151] = ("hello" + v5)
                    _ = v151
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v152 int = (fi150 + fi150)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v152
                }
            case "hello":
                v153 := GenS21{F0: 0}
                _ = v153
            case "a":
                if v4 {
                    _ = len(v2)
                } else {
                    v4 = v4
                }
            }
            if (v4 && false) {
                fmt.Println(fi150, fi150, fi150)
                fmt.Println(v5)
                v154 := GenS18{F0: (fi150 << ((18446744073709551615) & 63)), F1: 42}
                _ = v154
            }
            _ = fi150
        }
    }
    var v155 int32 = 0
    var v156 int16 = 529
    var v157 uint16 = 19
    switch -2147483648 {
    case -9223372036854775808:
        if v4 {
            if (0 >= 495) {
                switch (v5 + v5) {
                case "hello":
                    for fi158 := 0; fi158 < 11; fi158++ {
                        for fi159 := 0; fi159 < 3; fi159++ {
                            fmt.Println(v148, fi158, fi158)
                            v160 := v2[42]
                            _ = fi159
                            _ = v160
                        }
                        var v161 string = "\n\t"
                        v2[42] = v5
                        _ = fi158
                        _ = v161
                    }
                case "":
                    for fi162 := 0; fi162 < 15; fi162++ {
                        for fi163 := 0; fi163 < 11; fi163++ {
                            switch v5 {
                            case "foo bar":
                                fmt.Println(v148, v4, v5)
                                switch v5 {
                                case "foo bar":
                                    if v4 {
                                        delete(v2, 910)
                                        fmt.Println(v148, v148)
                                        var v164 uint16 = v157
                                        v148 = (v148 / (float64(-2.25) / ((float64((float32(v148) - ((float32(1.5) / (1e10 / (-2.25 / 1.0))) / 1.0))) / float64(((1e10 / 3.14159) + (0.1 + ((-2.25 * ((float32(-2.25) / 1e10) / 1e10)) / 1.0))))) * v148)))
                                        if true {
                                            v155 -= (v155 % (((v155 - -547)) | 1))
                                            if len(v3) > 0 {
                                                v3[0] = v5
                                            }
                                        } else {
                                            switch ("a" + v5) {
                                            case "a":
                                                var v165 string = v5
                                                _ = v165
                                            case "foo bar":
                                                v157 -= (v164 & (v157 >> ((42) & 63)))
                                            default:
                                                var v166 uint64 = 2
                                                if v4 {
                                                    var v167 bool = v4
                                                    switch genFn0() {
                                                    case -275:
                                                        var v168 string = v5
                                                        _ = v168
                                                    case 0:
                                                        if (4294967295 <= 18446744073709551615) {
                                                            if true {
                                                                var v169 int8 = int8(v156)
                                                                _ = v169
                                                            } else {
                                                                var v170 int16 = v156
                                                                fmt.Println(v156, v164)
                                                                v171 := GenS13{F0: 43, F1: 232}
                                                                var v172 int64 = 537
                                                                switch v5 {
                                                                case "\n\t":
                                                                    var v173 float32 = 3.14159
                                                                    _ = v173
                                                                }
                                                                fmt.Println(fi163, v148)
                                                                v167 = v167
                                                                var v174 float64 = (0.1 - v148)
                                                                var v175 int8 = 127
                                                                _ = v170
                                                                _ = v171
                                                                _ = v172
                                                                _ = v174
                                                                _ = v175
                                                            }
                                                            var v176 float32 = 3.14159
                                                            switch v5 {
                                                            case "a":
                                                                v157 -= 2
                                                            case "\n\t":
                                                                var v177 uint8 = 127
                                                                _ = v177
                                                            case "hello":
                                                                v178 := v1[v5]
                                                                _ = v178
                                                            }
                                                            _ = len(v3)
                                                            v3 = append(v3, "hello")
                                                            var v179 uint64 = v166
                                                            _ = v176
                                                            _ = v179
                                                        } else {
                                                            var v180 bool = (-9223372036854775808 >= 42)
                                                            v166 += v166
                                                            v180 = v180
                                                            var v181 uint8 = 0
                                                            if v4 {
                                                                v182 := v1[v5]
                                                                v157 = (v157 % ((2) | 1))
                                                                v183 := GenS22{F0: true, F1: 691, F2: (v166 * v166)}
                                                                _ = v182
                                                                _ = v183
                                                            }
                                                            _ = v180
                                                            _ = v181
                                                        }
                                                        v5 = v5
                                                    case 4294967295:
                                                        fmt.Println(v157, v166, v4)
                                                    }
                                                    v148 = v148
                                                    _ = v167
                                                }
                                                var v184 float64 = -2.25
                                                _ = v166
                                                _ = v184
                                            }
                                        }
                                        _ = v164
                                    }
                                default:
                                    v4 = (((!(((v155 > v155) || (!false)) || (("hello" + ("hello" + v5)) >= "foo bar"))) || (1 > int8(v155))) && (v4 || (1e10 != (-2.25 - 3.14159))))
                                    var v185 string = "foo bar"
                                    v157 -= (v157 % (((v157 << ((4294967295) & 63))) | 1))
                                    _ = v185
                                }
                            case "hello":
                                if (!v4) {
                                    v155 -= (v155 ^ (v155 * 42))
                                }
                            }
                            _ = fi163
                        }
                        for fi186 := 0; fi186 < 11; fi186++ {
                            switch fi162 {
                            case 127:
                                var v187 bool = (!((v157 | 134) != v157))
                                _ = v187
                            case 9223372036854775807:
                                fmt.Println(fi186, fi162)
                            }
                            if ((44 != 1) || (!(((true && (fi186 >= (fi186 - 1))) || (563 > 0)) && v4))) {
                                var v188 uint32 = 625
                                var v189 int64 = 42
                                _ = v188
                                _ = v189
                            }
                            var v190 int8 = 1
                            _ = fi186
                            _ = v190
                        }
                        _ = fi162
                    }
                default:
                    v191 := GenS23{F0: (((v5 + "a") + v5) + ""), F1: (341 != 357), F2: (v155 % ((v155) | 1))}
                    _ = v191
                }
            }
            if v4 {
                v3 = append(v3, ((("" + ((v5 + v5) + v5)) + "a") + (v5 + v5)))
            } else {
                switch (v5 + v5) {
                case "foo bar":
                    v192 := GenS23{F0: v5, F1: (0.1 != (3.14159 * v148)), F2: v155}
                    _ = v192
                case "\n\t":
                    v148 += -2.25
                case "":
                    for fi193 := 0; fi193 < 17; fi193++ {
                        _ = len(v2)
                        fmt.Println(v5, v155, v157)
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        _ = fi193
                    }
                default:
                    switch ("hello" + v5) {
                    case "\n\t":
                        fmt.Println(v4)
                    case "":
                        if true {
                            var v194 bool = v4
                            v148 += v148
                            _ = v194
                        }
                    case "foo bar":
                        v156 = -32768
                    default:
                        fmt.Println(v157, v157)
                    }
                }
                switch ("" + v5) {
                case "foo bar":
                    var v195 float32 = (1.5 * -2.25)
                    _ = v195
                }
            }
        }
    case -2147483648:
        var v196 uint64 = 4294967295
        _ = v196
    case 4294967295:
        if len(v3) > 0 {
            v3[0] = v5
        }
    default:
        v197 := GenS11{F0: 263, F1: 32767}
        _ = v197
    }
    v157 -= v157
    v155 -= 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v148:", v148)
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("v157:", v157)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v148
    _ = v155
    _ = v156
    _ = v157
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
