package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 string; F2 int8 }
type GenS4 struct { F0 int32; F1 int64; F2 int8 }
type GenS5 struct { F0 int32; F1 int; F2 int }
type GenS6 struct { F0 uint8; F1 int8; F2 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS9 struct { F0 int8; F1 int64; F2 string }
type GenS10 struct { F0 bool; F1 uint16; F2 int32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int64; F1 uint64 }
type GenS13 struct { F0 int64; F1 int32; F2 bool }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 bool; F1 string }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint8 }
func genFn0() int {
    if (1.5 <= 0.1) {
        var v6 int8 = 57
        v7 := GenS1{F0: 235}
        if true {
            var v8 uint32 = 2147483648
            var v9 float32 = (3.14159 / 3.14159)
            _ = v8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    var v10 uint64 = 2
    v10 += (^(v10))
    v10 = (v10 - v10)
    v11 := GenS2{F0: 77}
    _ = v10
    _ = v11
    return 9223372036854775807
}
func genFn1() bool {
    if false {
        if (!false) {
            var v12 uint64 = 42
            if false {
                v13 := GenS1{F0: 2}
                v12 = 0
                _ = v13
            }
            switch "foo bar" {
            case "foo bar":
                if (true && (42 <= 42)) {
                    v14 := GenS1{F0: 412}
                    switch "hello" {
                    case "foo bar":
                        if (0 < 100) {
                            var v15 uint16 = 413
                            _ = v15
                        } else {
                            v16 := GenS2{F0: 0}
                            for fi17 := 0; fi17 < 19; fi17++ {
                                v18 := GenS3{F0: false, F1: "", F2: 0}
                                var v19 int16 = 1000
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            switch 1 {
                            case -9223372036854775808:
                                v12 = v12
                            default:
                                fmt.Println(v12)
                            }
                            _ = v16
                        }
                    case "a":
                        v12 = 2
                    }
                    _ = v14
                } else {
                    v20 := []int{-117, 92, 32767, 0}
                    _ = v20
                }
            case "":
                v21 := GenS1{F0: 4294967295}
                _ = v21
            case "a":
                v22 := GenS4{F0: -2147483648, F1: -9223372036854775808, F2: 0}
                _ = v22
            }
            _ = v12
        } else {
            v23 := GenS5{F0: 42, F1: 127, F2: 1}
            var v24 bool = (1.5 > 3.14159)
            for fi25 := 0; fi25 < 8; fi25++ {
                v26 := []int{2147483647, (fi25 % (((fi25 | fi25)) | 1)), fi25, genFn0()}
                _ = fi25
                _ = v26
            }
            _ = v23
            _ = v24
        }
    } else {
        if (1 != 0) {
            v27 := []int{-839, 1, 1}
            if false {
                for fi28 := 0; fi28 < 14; fi28++ {
                    if true {
                        if true {
                            switch fi28 {
                            case 42:
                                v29 := GenS5{F0: -2147483648, F1: 133, F2: fi28}
                                _ = v29
                            case 605:
                                if len(v27) > 0 {
                                    v27[0] = fi28
                                }
                            case 127:
                                if len(v27) > 0 {
                                    v27[0] = 127
                                }
                            default:
                                fmt.Println(fi28)
                            }
                        }
                    } else {
                        if false {
                            var v30 uint = 2
                            _ = v30
                        }
                        switch "a" {
                        case "foo bar":
                            v31 := GenS4{F0: 301, F1: 1, F2: 100}
                            _ = v31
                        case "a":
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                        case "":
                            if ((42 != 550) && true) {
                                switch "hello" {
                                case "foo bar":
                                    for fi32 := 0; fi32 < 2; fi32++ {
                                        if len(v27) >= 2 {
                                            v27 = v27[1:len(v27)]
                                        }
                                        var v33 uint = 1
                                        _ = fi32
                                        _ = v33
                                    }
                                case "":
                                    if true {
                                        if ((-32768 == 1000) || true) {
                                            fmt.Println(fi28)
                                            if len(v27) >= 2 {
                                                v27 = v27[1:len(v27)]
                                            }
                                        }
                                        var v34 bool = (!(true && false))
                                        v35 := GenS6{F0: 127, F1: -1, F2: 0}
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v36 := GenS7{F0: -1}
                                        if ((true || (!(false && true))) || (!true)) {
                                            var v37 string = ""
                                            for fi38 := 0; fi38 < 12; fi38++ {
                                                v39 := GenS4{F0: 0, F1: -9223372036854775808, F2: 42}
                                                v37 = "\n\t"
                                                if (44 != 42) {
                                                    v37 = v37
                                                } else {
                                                    fmt.Println(v37)
                                                }
                                                _ = fi38
                                                _ = v39
                                            }
                                            v40 := GenS8{F0: 42, F1: 2147483648, F2: -1}
                                            _ = v37
                                            _ = v40
                                        } else {
                                            var v41 bool = false
                                            _ = v41
                                        }
                                        _ = v36
                                    }
                                case "hello":
                                    v42 := GenS5{F0: 0, F1: fi28, F2: 128}
                                    _ = v42
                                }
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                                fmt.Println(fi28)
                            } else {
                                var v43 float64 = 0.1
                                _ = v43
                            }
                        }
                        fmt.Println(fi28)
                    }
                    switch fi28 {
                    case 1:
                        fmt.Println(fi28)
                    case 32768:
                        if false {
                            if (32767 > 1) {
                                fmt.Println(fi28)
                                v44 := GenS3{F0: false, F1: "a", F2: 84}
                                _ = v44
                            } else {
                                if (false && (fi28 != (fi28 & (fi28 + genFn0())))) {
                                    if (!true) {
                                        var v45 uint = 4294967295
                                        _ = v45
                                    } else {
                                        var v46 bool = (false && true)
                                        fmt.Println(v46)
                                        fmt.Println(fi28)
                                        _ = v46
                                    }
                                    var v47 uint16 = 42
                                    _ = v47
                                }
                            }
                            var v48 int8 = 31
                            _ = v48
                        } else {
                            for fi49 := 0; fi49 < 7; fi49++ {
                                v50 := GenS9{F0: 1, F1: 1, F2: "foo bar"}
                                _ = fi49
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 1; fi51++ {
                                var v52 uint32 = 2147483647
                                v52 = (v52 | (v52 ^ (v52 << ((459) & 63))))
                                if true {
                                    fmt.Println(fi28)
                                    if false {
                                        v52 = 471
                                        fmt.Println(v52, fi28)
                                        if ("\n\t" == "hello") {
                                            var v53 int32 = 1
                                            if (79 >= 25) {
                                                if len(v27) >= 2 {
                                                    v27 = v27[1:len(v27)]
                                                }
                                            } else {
                                                if false {
                                                    if true {
                                                        v52 = v52
                                                    } else {
                                                        v53 = (v53 + 369)
                                                        if (true || true) {
                                                            if false {
                                                                _ = len(v27)
                                                                var v54 uint8 = 0
                                                                _ = v54
                                                            }
                                                        }
                                                        v53 = (^(v53))
                                                    }
                                                }
                                                var v55 int = (fi28 >> ((18446744073709551615) & 63))
                                                switch "" {
                                                case "\n\t":
                                                    if false {
                                                        v53 += v53
                                                        fmt.Println(fi51)
                                                    } else {
                                                        v53 += v53
                                                        fmt.Println(v52, v55, fi28)
                                                    }
                                                default:
                                                    var v56 string = "foo bar"
                                                    _ = v56
                                                }
                                                _ = v55
                                            }
                                            if len(v27) >= 2 {
                                                v27 = v27[1:len(v27)]
                                            }
                                            _ = v53
                                        } else {
                                            if len(v27) >= 2 {
                                                v27 = v27[1:len(v27)]
                                            }
                                            var v57 uint64 = 38
                                            v58 := GenS10{F0: (100 > -112), F1: 42, F2: 594}
                                            _ = v57
                                            _ = v58
                                        }
                                        var v59 string = "a"
                                        _ = v59
                                    }
                                }
                                _ = fi51
                                _ = v52
                            }
                            if (true && false) {
                                if true {
                                    var v60 uint = 255
                                    if true {
                                        _ = len(v27)
                                    }
                                    v60 += (v60 & 42)
                                    _ = v60
                                }
                                var v61 uint64 = 382
                                for fi62 := 0; fi62 < 13; fi62++ {
                                    if (0 != 246) {
                                        fmt.Println(v61, v61)
                                    } else {
                                        v27 = append(v27, genFn0())
                                        v61 -= v61
                                        fmt.Println(fi62, v61)
                                    }
                                    if len(v27) > 0 {
                                        v27[0] = (fi28 - 32767)
                                    }
                                    _ = fi62
                                }
                                _ = v61
                            } else {
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                            }
                        }
                    case 2147483647:
                        for fi63 := 0; fi63 < 20; fi63++ {
                            v64 := GenS7{F0: 1}
                            if (-38 != 42) {
                                fmt.Println(fi28, fi63)
                            } else {
                                switch (fi28 * 42) {
                                case -750:
                                    var v65 float64 = 1.5
                                    _ = v65
                                default:
                                    fmt.Println(fi63, fi28)
                                }
                                v66 := GenS11{F0: 52}
                                if len(v27) > 0 {
                                    v27[0] = (fi28 - 833)
                                }
                                _ = v66
                            }
                            _ = fi63
                            _ = v64
                        }
                    }
                    _ = fi28
                }
            }
            _ = v27
        }
        v67 := GenS10{F0: true, F1: 1, F2: -1}
        for fi68 := 0; fi68 < 3; fi68++ {
            v69 := []int{fi68}
            switch "hello" {
            case "foo bar":
                var v70 int8 = 127
                _ = v70
            case "\n\t":
                v71 := GenS12{F0: 1, F1: 42}
                _ = v71
            case "hello":
                switch fi68 {
                case -461:
                    var v72 uint64 = 2
                    _ = v72
                }
            default:
                var v73 bool = true
                _ = v73
            }
            v74 := GenS12{F0: 42, F1: 4294967295}
            _ = fi68
            _ = v69
            _ = v74
        }
        _ = v67
    }
    if false {
        var v75 int16 = 637
        v75 -= (-(v75))
        _ = v75
    }
    var v76 uint16 = 255
    _ = v76
    return (!true)
}

func main() {
    var v1 int32 = 0
    var v2 float32 = 1.5
    var v3 uint16 = 42
    var v4 float64 = 0.0
    var v5 int32 = (v1 >> ((0) & 63))
    switch ("" + ("foo bar" + (("" + (("hello" + ("\n\t" + ("a" + "hello"))) + "")) + ("a" + "")))) {
    case "\n\t":
        if (!true) {
            v1 = (v1 | v1)
            for fi77 := 0; fi77 < 2; fi77++ {
                v5 += v5
                if true {
                    fmt.Println(v5)
                    v78 := []int{(fi77 + fi77), (fi77 << ((42) & 63)), 918, 431}
                    _ = v78
                } else {
                    v79 := []int{fi77, 9223372036854775807, fi77}
                    _ = v79
                }
                var v80 float32 = 0.1
                _ = fi77
                _ = v80
            }
        } else {
            v3 = 42
        }
    case "foo bar":
        v1 = v1
    case "a":
        switch (((("" + "hello") + "\n\t") + "hello") + "foo bar") {
        case "\n\t":
            var v81 int8 = 42
            _ = v81
        case "hello":
            if (genFn0() <= -578) {
                if false {
                    v2 = v2
                    fmt.Println(v4)
                } else {
                    v4 += v4
                    v82 := []int{int(v1), 534}
                    _ = v82
                }
                v83 := []int{128}
                _ = v83
            }
        case "foo bar":
            v84 := GenS13{F0: 137, F1: (^(v1)), F2: true}
            _ = v84
        default:
            for fi85 := 0; fi85 < 6; fi85++ {
                if (2 <= 664) {
                    var v86 uint8 = 128
                    var v87 int64 = 42
                    _ = v86
                    _ = v87
                }
                v88 := []int{(fi85 - fi85)}
                fmt.Println(v1)
                _ = fi85
                _ = v88
            }
        }
    default:
        var v89 float32 = v2
        _ = v89
    }
    v90 := []int{}
    switch "a" {
    case "\n\t":
        var v91 uint = 836
        _ = v91
    case "a":
        for fi92 := 0; fi92 < 9; fi92++ {
            var v93 float32 = 1e10
            v2 += ((float32(float64(v2)) + 3.14159) / v93)
            switch (^(fi92)) {
            case 127:
                var v94 int64 = -1
                _ = v94
            case 4294967295:
                if (!(!(v2 != (v2 + (v2 / 1e10))))) {
                    v95 := GenS14{F0: (v1 / (((v1 / ((v5) | 1))) | 1))}
                    _ = v95
                }
            }
            _ = fi92
            _ = v93
        }
    case "hello":
        v90 = append(v90, genFn0())
    }
    fmt.Println(v3, v4, v4)
    for fi96 := 0; fi96 < 8; fi96++ {
        if true {
            v97 := GenS15{F0: false, F1: "\n\t"}
            if true {
                v5 = (v5 / (((v1 % ((v1) | 1))) | 1))
                switch ("hello" + "\n\t") {
                case "\n\t":
                    v90 = append(v90, 128)
                case "hello":
                    if (((4294967295 >= 42) || true) || ((1 != uint64(v3)) || (-32768 >= 0))) {
                        v98 := GenS11{F0: -1}
                        if true {
                            v99 := GenS1{F0: 255}
                            var v100 int8 = 1
                            v101 := GenS11{F0: 116}
                            _ = v100
                            _ = v101
                            _ = v99
                        } else {
                            v90 = append(v90, 901)
                            var v102 uint64 = 600
                            _ = v102
                        }
                        _ = v98
                    }
                default:
                    switch "foo bar" {
                    case "hello":
                        var v103 int8 = 1
                        _ = v103
                    }
                }
            } else {
                v2 = v2
                if (1 >= 1) {
                    switch "hello" {
                    case "":
                        for fi104 := 0; fi104 < 15; fi104++ {
                            v105 := GenS16{F0: (-(fi104))}
                            _ = fi104
                            _ = v105
                        }
                    case "hello":
                        var v106 float32 = v2
                        _ = v106
                    default:
                        v5 = (v1 % (((v1 & 1)) | 1))
                    }
                    var v107 float64 = 0.1
                    _ = v107
                } else {
                    v1 -= (v1 * (v5 - v1))
                    v3 = (v3 & 65535)
                    if (v3 == (v3 >> ((255) & 63))) {
                        switch fi96 {
                        case -322:
                            for fi108 := 0; fi108 < 18; fi108++ {
                                if false {
                                    v109 := GenS6{F0: 15, F1: -1, F2: 196}
                                    _ = v109
                                }
                                switch "" {
                                case "foo bar":
                                    v4 += (v4 * v4)
                                case "a":
                                    v4 = (-2.25 - (v4 + 0.1))
                                }
                                v110 := GenS13{F0: 1, F1: v5, F2: (-9223372036854775808 == 424)}
                                _ = fi108
                                _ = v110
                            }
                        default:
                            v4 = (3.14159 / 1e10)
                        }
                    }
                }
            }
            v4 = v4
            _ = v97
        }
        v3 += v3
        if false {
            fmt.Println(v5, v3, fi96)
            var v111 uint16 = (v3 * v3)
            _ = v111
        }
        _ = fi96
    }
    if true {
        v3 -= v3
    }
    fmt.Println(v1)
    v112 := GenS8{F0: v3, F1: 1, F2: 108}
    if true {
        var v113 string = ("" + "")
        _ = v113
    } else {
        if len(v90) >= 2 {
            v90 = v90[1:len(v90)]
        }
        if (!(-416 != -246)) {
            for fi114 := 0; fi114 < 10; fi114++ {
                fmt.Println(v3)
                var v115 int8 = 7
                v116 := GenS13{F0: -153, F1: (v1 % (((^(v5))) | 1)), F2: (!true)}
                _ = fi114
                _ = v115
                _ = v116
            }
        } else {
            v117 := GenS17{F0: uint8(v3)}
            v4 = ((v4 - (((v4 - (((v4 - -2.25) + ((-2.25 + v4) - 3.14159)) / ((v4 - (v4 + v4)) / (v4 / v4)))) + v4) - float64(float32(v4)))) * v4)
            _ = v117
        }
    }
    fmt.Println(v4, v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v90:", len(v90), v90)
    _ = v1
    _ = v112
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v90
}
