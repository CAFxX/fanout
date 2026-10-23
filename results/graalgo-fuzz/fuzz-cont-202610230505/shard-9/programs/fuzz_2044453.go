package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int32; F2 bool }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 int16; F1 uint; F2 uint16 }
type GenS10 struct { F0 int32; F1 bool }
type GenS11 struct { F0 uint64; F1 uint16; F2 int }
func genFn0() uint32 {
    v6 := GenS1{F0: 42, F1: 1, F2: 0}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 bool = true
        if true {
            for fi9 := 0; fi9 < 14; fi9++ {
                fmt.Println(fi9)
                var v10 int64 = 9223372036854775807
                v11 := []int{32767, fi7, (^(fi9))}
                v11 = append(v11, -455)
                _ = fi9
                _ = v10
                _ = v11
            }
            v8 = (355 < 0)
            v12 := []int{(fi7 << ((18446744073709551615) & 63)), (fi7 + (^(fi7)))}
            _ = v12
        } else {
            v13 := []int{fi7}
            v8 = (false && v8)
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    v14 := GenS2{F0: 255, F1: -109}
    _ = v14
    _ = v6
    return 1
}
func genFn1() int64 {
    if (42 != 42) {
        for fi15 := 0; fi15 < 13; fi15++ {
            v16 := GenS1{F0: 1, F1: 0, F2: 116}
            v17 := []int{}
            var v18 int = fi15
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    } else {
        if ((1 < -653) && false) {
            var v19 float32 = (((0.1 + 3.14159) + 1.5) / (-2.25 + (-2.25 - (float32(float64(0.0)) + -2.25))))
            v20 := GenS3{F0: 32767, F1: 1}
            var v21 uint32 = 2
            _ = v19
            _ = v20
            _ = v21
        } else {
            var v22 uint8 = 255
            _ = v22
        }
        var v23 int16 = 1000
        v24 := []int{284, -356}
        _ = v23
        _ = v24
    }
    v25 := []int{4294967295}
    if ((float64((0.1 - (3.14159 + 1.5))) + 1e10) > (0.0 * 1e10)) {
        if (-107 != -64) {
            var v26 uint16 = 32767
            if (408 > 1) {
                for fi27 := 0; fi27 < 5; fi27++ {
                    switch ("a" + (("hello" + ("hello" + "")) + (("foo bar" + "hello") + (("a" + "") + ("\n\t" + "hello"))))) {
                    case "":
                        var v28 int8 = -1
                        _ = v28
                    default:
                        if true {
                            fmt.Println(v26)
                            if true {
                                v26 = (v26 % ((v26) | 1))
                                switch fi27 {
                                case -829:
                                    var v29 int32 = 2147483647
                                    _ = v29
                                case 1:
                                    if (("hello" < (("a" + "foo bar") + "")) && (fi27 != (fi27 / ((fi27) | 1)))) {
                                        for fi30 := 0; fi30 < 12; fi30++ {
                                            v26 = 255
                                            fmt.Println(fi30, fi30, fi27)
                                            var v31 float64 = (1e10 / 1.0)
                                            if (!(18446744073709551615 > 0)) {
                                                var v32 int16 = 32767
                                                v26 -= v26
                                                if ((9223372036854775807 < 9223372036854775807) && (0 == -2147483648)) {
                                                    var v33 uint32 = 319
                                                    var v34 float32 = 3.14159
                                                    _ = v33
                                                    _ = v34
                                                } else {
                                                    var v35 int8 = 0
                                                    if false {
                                                        if true {
                                                            _ = len(v25)
                                                            v36 := GenS2{F0: 1, F1: (^(v35))}
                                                            var v37 int = 128
                                                            var v38 int16 = (v32 << ((1) & 63))
                                                            v35 += v35
                                                            _ = v36
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                        var v39 uint8 = 159
                                                        if (22 < 923) {
                                                            v26 -= (v26 + 1)
                                                            v35 -= (v35 ^ 1)
                                                            if false {
                                                                var v40 uint8 = (v39 >> ((721) & 63))
                                                                _ = v40
                                                            } else {
                                                                v39 = (v39 ^ 255)
                                                                v41 := GenS2{F0: (v39 >> ((7) & 63)), F1: v35}
                                                                var v42 int64 = 9223372036854775807
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v43 := GenS2{F0: v39, F1: 42}
                                                                    _ = v43
                                                                default:
                                                                    v26 += v26
                                                                }
                                                                switch "" {
                                                                case "foo bar":
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        v31 += v31
                                                                    case "":
                                                                        v39 += v39
                                                                    case "foo bar":
                                                                        if true {
                                                                            var v44 int16 = (v32 | (^(v32)))
                                                                            _ = v44
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    v45 := GenS3{F0: 0, F1: -2147483648}
                                                                    v26 = 570
                                                                    var v46 float32 = 0.0
                                                                    _ = v45
                                                                    _ = v46
                                                                case "hello":
                                                                    var v47 float32 = 0.0
                                                                    v48 := GenS1{F0: 0, F1: 2147483647, F2: 0}
                                                                    _ = v47
                                                                    _ = v48
                                                                default:
                                                                    var v49 float32 = -2.25
                                                                    _ = v49
                                                                }
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                            if true {
                                                                switch (fi27 * 2147483647) {
                                                                case 403:
                                                                    fmt.Println(v26)
                                                                case 32768:
                                                                    var v50 int32 = 1
                                                                    _ = v50
                                                                }
                                                            }
                                                        }
                                                        var v51 int16 = -851
                                                        _ = v39
                                                        _ = v51
                                                    } else {
                                                        var v52 bool = (true || (!(1 < 2)))
                                                        switch (fi27 & -47) {
                                                        case -9223372036854775808:
                                                            v53 := GenS4{F0: 801, F1: 42}
                                                            _ = v53
                                                        case 616:
                                                            var v54 float32 = float32(-2.25)
                                                            _ = v54
                                                        }
                                                        _ = v52
                                                    }
                                                    _ = v35
                                                }
                                                _ = v32
                                            }
                                            if true {
                                                var v55 float64 = v31
                                                if (32767 != 1) {
                                                    switch (fi27 | fi27) {
                                                    case 128:
                                                        var v56 float64 = ((((v31 * v31) / -2.25) * ((((0.0 * float64(((((1.5 / float32(1.5)) / 1e10) * (3.14159 + ((((0.1 + -2.25) * -2.25) / (((float32(1e10) * 0.1) * 1e10) / 3.14159)) - float32((v55 * (v55 + 0.0)))))) - 1e10))) * v55) * ((v55 / 1e10) / (v55 / (0.1 + ((3.14159 * 0.1) + float64(1e10)))))) + float64(0.0))) / (v31 - 1e10))
                                                        _ = v56
                                                    }
                                                    v55 += float64(float32(((((float64(float32(0.0)) * v31) * float64(-2.25)) / ((v31 + v55) / v55)) + 0.1)))
                                                } else {
                                                    var v57 int32 = 0
                                                    v58 := GenS3{F0: 255, F1: (v57 | (-(v57)))}
                                                    v59 := GenS5{F0: v26, F1: (-(v57)), F2: false}
                                                    v31 = float64((1.5 + (0.0 + -2.25)))
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                }
                                                _ = v55
                                            }
                                            _ = fi30
                                            _ = v31
                                        }
                                        if false {
                                            v60 := GenS5{F0: (^(v26)), F1: 1, F2: false}
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                            if ((float32(float64(0.0)) < 3.14159) && (130 <= 65535)) {
                                                v26 = v26
                                            }
                                            _ = v60
                                        } else {
                                            v26 -= 42
                                        }
                                        v61 := GenS6{F0: 0, F1: 4294967295}
                                        _ = v61
                                    }
                                }
                            }
                        }
                    }
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                    switch (fi27 / (((^(fi27))) | 1)) {
                    case 790:
                        _ = len(v25)
                    case 32767:
                        var v62 float32 = 3.14159
                        _ = v62
                    }
                    _ = fi27
                }
                switch "foo bar" {
                case "\n\t":
                    v63 := GenS2{F0: 167, F1: 100}
                    _ = v63
                case "a":
                    switch "\n\t" {
                    case "hello":
                        v26 = v26
                    case "\n\t":
                        v64 := GenS7{F0: (v26 % ((v26) | 1)), F1: 9223372036854775807}
                        _ = v64
                    }
                }
            } else {
                var v65 uint8 = 0
                if true {
                    if (4294967295 != 2) {
                        if false {
                            v66 := GenS7{F0: v26, F1: 0}
                            switch "a" {
                            case "\n\t":
                                var v67 float32 = 3.14159
                                _ = v67
                            case "foo bar":
                                var v68 string = "hello"
                                _ = v68
                            case "hello":
                                for fi69 := 0; fi69 < 9; fi69++ {
                                    fmt.Println(v26)
                                    var v70 string = "a"
                                    v71 := GenS8{F0: "foo bar", F1: 42}
                                    _ = fi69
                                    _ = v70
                                    _ = v71
                                }
                            default:
                                fmt.Println(v65, v26)
                            }
                            _ = v66
                        } else {
                            for fi72 := 0; fi72 < 16; fi72++ {
                                v26 -= 255
                                if len(v25) > 0 {
                                    v25[0] = (fi72 % ((fi72) | 1))
                                }
                                fmt.Println(v26)
                                _ = fi72
                            }
                        }
                        if (!true) {
                            for fi73 := 0; fi73 < 18; fi73++ {
                                v74 := GenS6{F0: uint(v65), F1: 0}
                                var v75 uint64 = 1
                                _ = fi73
                                _ = v74
                                _ = v75
                            }
                            var v76 uint32 = 2
                            if true {
                                if false {
                                    fmt.Println(v76)
                                    v65 = (v65 | 253)
                                    v77 := GenS9{F0: -32768, F1: 65535, F2: v26}
                                    _ = v77
                                }
                                if (0 < -1) {
                                    var v78 int8 = -24
                                    _ = len(v25)
                                    var v79 string = "foo bar"
                                    _ = v78
                                    _ = v79
                                } else {
                                    switch ("hello" + "foo bar") {
                                    case "foo bar":
                                        var v80 uint = 0
                                        _ = v80
                                    case "":
                                        v81 := GenS2{F0: 98, F1: 127}
                                        _ = v81
                                    case "hello":
                                        fmt.Println(v65)
                                    default:
                                        var v82 bool = false
                                        _ = v82
                                    }
                                    for fi83 := 0; fi83 < 6; fi83++ {
                                        if true {
                                            fmt.Println(v26, v26, fi83)
                                            var v84 float64 = (-2.25 / ((float64(3.14159) * -2.25) * 0.1))
                                            _ = v84
                                        }
                                        fmt.Println(v76)
                                        _ = fi83
                                    }
                                    var v85 uint32 = uint32(v26)
                                    _ = v85
                                }
                            } else {
                                switch "" {
                                case "\n\t":
                                    for fi86 := 0; fi86 < 7; fi86++ {
                                        if len(v25) >= 2 {
                                            v25 = v25[1:len(v25)]
                                        }
                                        v76 -= genFn0()
                                        v65 = 128
                                        _ = fi86
                                    }
                                case "":
                                    var v87 uint8 = (v65 << ((518) & 63))
                                    _ = v87
                                default:
                                    v26 += uint16(v65)
                                }
                            }
                            _ = v76
                        }
                    }
                } else {
                    v88 := GenS10{F0: 2147483647, F1: true}
                    _ = v88
                }
                _ = v65
            }
            var v89 float64 = ((-2.25 - 0.0) + -2.25)
            _ = v26
            _ = v89
        } else {
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
        }
        var v90 uint8 = 49
        _ = v90
    }
    _ = v25
    return -9223372036854775808
}

func main() {
    var v1 bool = true
    var v2 int = 2147483647
    v3 := make(map[string]uint32)
    var v4 uint64 = 0
    var v5 uint = 4294967295
    v1 = false
    for fi91 := 0; fi91 < 19; fi91++ {
        var v92 float32 = 1.5
        v93 := GenS7{F0: 876, F1: -576}
        switch (("\n\t" + "a") + (((((("foo bar" + ("foo bar" + ("foo bar" + ("a" + ("hello" + ("\n\t" + "")))))) + "a") + "foo bar") + "\n\t") + ("" + (((("hello" + ((("foo bar" + "\n\t") + (("hello" + ("foo bar" + (("" + "") + "\n\t"))) + "\n\t")) + "")) + ("" + ("a" + ("" + "hello")))) + "a") + ("foo bar" + "\n\t")))) + "a")) {
        case "hello":
            v2 = (fi91 - v2)
        case "a":
            for fi94 := 0; fi94 < 1; fi94++ {
                v5 -= (v5 * (^(v5)))
                v95 := GenS8{F0: "\n\t", F1: -852}
                v96 := GenS8{F0: ((("foo bar" + "hello") + (("foo bar" + "a") + "foo bar")) + (("\n\t" + ((("foo bar" + "hello") + "") + "foo bar")) + "")), F1: -9223372036854775808}
                _ = fi94
                _ = v95
                _ = v96
            }
        }
        _ = fi91
        _ = v92
        _ = v93
    }
    var v97 string = ((("a" + "\n\t") + ("\n\t" + "\n\t")) + ("foo bar" + "a"))
    fmt.Println(v1, v1, v2)
    v98 := []int{(v2 % ((v2) | 1)), v2}
    for fi99 := 0; fi99 < 20; fi99++ {
        fmt.Println(v2)
        if (!v1) {
            fmt.Println(v1, fi99)
        }
        _ = fi99
    }
    var v100 int32 = 94
    delete(v3, v97)
    v101 := v3[(v97 + v97)]
    v101 = uint32(v5)
    v102 := v3[v97]
    switch "foo bar" {
    case "a":
        v98 = append(v98, v2)
    default:
        if ((^(v4)) < v4) {
            v2 -= v2
            v4 += (v4 >> ((42) & 63))
        }
    }
    if (3.14159 > 1.5) {
        v5 += 4294967295
    } else {
        for fi103 := 0; fi103 < 3; fi103++ {
            v104 := GenS7{F0: 0, F1: 9223372036854775807}
            v105 := GenS10{F0: v100, F1: (!(1 == 9223372036854775807))}
            _ = fi103
            _ = v104
            _ = v105
        }
        var v106 string = v97
        fmt.Println(v2, v5, v100)
        _ = v106
    }
    if (9223372036854775807 == -907) {
        if (genFn0() < 430) {
            var v107 int32 = v100
            for fi108 := 0; fi108 < 1; fi108++ {
                var v109 int8 = 99
                _ = fi108
                _ = v109
            }
            _ = v107
        } else {
            if v1 {
                fmt.Println(v5, v102, v101)
                v110 := GenS7{F0: 32767, F1: -206}
                v111 := GenS7{F0: 255, F1: -1}
                _ = v110
                _ = v111
            } else {
                var v112 float32 = 1e10
                _ = v112
            }
            for fi113 := 0; fi113 < 6; fi113++ {
                if (((^(v101)) > (v102 + (v101 * v101))) || (-2.25 <= 0.1)) {
                    v101 = 632
                    if len(v98) >= 2 {
                        v98 = v98[1:len(v98)]
                    }
                } else {
                    fmt.Println(v2, v101)
                    var v114 uint32 = (^(v102))
                    v115 := GenS11{F0: (v4 | (^(v4))), F1: 32767, F2: v2}
                    _ = v114
                    _ = v115
                }
                var v116 uint8 = 127
                if v1 {
                    var v117 uint16 = 32767
                    _ = v117
                }
                _ = fi113
                _ = v116
            }
        }
    }
    var v118 float64 = (float64(float32(-2.25)) + (-2.25 * 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v97:", v97)
    fmt.Println("slicelen v98:", len(v98), v98)
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v97
    _ = v98
}
