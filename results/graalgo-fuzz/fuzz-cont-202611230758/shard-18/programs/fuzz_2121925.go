package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint64; F2 string }
type GenS6 struct { F0 int16; F1 bool; F2 int }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 int32; F1 bool }
type GenS9 struct { F0 int64; F1 int8; F2 string }
type GenS10 struct { F0 string; F1 int32; F2 uint8 }
type GenS11 struct { F0 int8 }
func genFn0(p0 int64, p1 uint64, p2 uint) uint64 {
    var v5 uint64 = (p1 / (((p1 >> ((p2) & 63))) | 1))
    fmt.Println(v5, p1)
    var v6 int64 = p0
    if (float64(1e10) != 3.14159) {
        v7 := GenS1{F0: 2147483648, F1: -128}
        if false {
            fmt.Println(p1, v5, v6)
        }
        if (!(!true)) {
            switch "hello" {
            case "\n\t":
                var v8 uint16 = 65535
                _ = v8
            case "a":
                var v9 uint64 = (^(p1))
                _ = v9
            }
            var v10 float32 = float32(((0.0 / float64(float32(((-2.25 * 0.1) * (1.5 + ((float64(0.0) * 1.5) / float64(0.1))))))) + ((1.5 * (0.0 + (0.0 * 1.5))) - 1e10)))
            _ = v10
        }
        _ = v7
    } else {
        fmt.Println(p0, v5, v5)
        if true {
            v11 := []int{}
            p1 += (p1 & 42)
            switch "a" {
            case "a":
                switch 2147483647 {
                case 2147483647:
                    p0 -= p0
                }
            default:
                p0 = int64(p2)
            }
            _ = v11
        }
        var v12 bool = (("" + ("hello" + "foo bar")) > ((("foo bar" + "hello") + "a") + "a"))
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return v5
}
func genFn1() int32 {
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 float32 = -2.25
        var v15 int32 = -206
        switch "a" {
        case "a":
            if false {
                fmt.Println(fi13, v15)
            } else {
                v15 = (v15 & 1)
            }
        case "hello":
            if false {
                fmt.Println(fi13)
                fmt.Println(v14)
                v15 += int32(fi13)
            } else {
                v15 -= -2147483648
                var v16 float32 = float32(1e10)
                _ = v16
            }
        case "":
            var v17 int64 = 9223372036854775807
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    var v18 uint32 = 2
    v19 := GenS2{F0: 42, F1: (-2.25 < (1.5 + float64(((((0.0 + 1e10) * -2.25) * 0.1) / 1.0)))), F2: 42}
    fmt.Println(v18)
    var v20 uint8 = 127
    _ = v18
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 string = "\n\t"
    v2 := []int{537, 4294967295}
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    v1 = "\n\t"
    v1 = ((v1 + ("\n\t" + v1)) + (("foo bar" + (v1 + ("\n\t" + "foo bar"))) + "a"))
    v21 := GenS3{F0: 0}
    var v22 int32 = -2147483648
    var v23 int = 9223372036854775807
    if ((1.5 == -2.25) && false) {
        var v24 int8 = 0
        _ = v24
    } else {
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 uint = 65535
            if false {
                if (!(!false)) {
                    v4[1] = false
                    var v27 uint32 = 4294967295
                    var v28 uint16 = uint16(v26)
                    _ = v27
                    _ = v28
                }
                switch (("" + (v1 + "")) + ("\n\t" + (v1 + (v1 + "\n\t")))) {
                case "\n\t":
                    var v29 float64 = 0.0
                    _ = v29
                case "a":
                    switch "" {
                    case "\n\t":
                        for fi30 := 0; fi30 < 6; fi30++ {
                            if true {
                                v22 = (v22 | v22)
                                if false {
                                    var v31 int32 = (-(v22))
                                    if ((((true && false) || (127 <= 127)) && (false || (v23 <= v23))) && false) {
                                        v32 := v3[v1]
                                        v31 = v31
                                        _ = v32
                                    } else {
                                        v4[(^(fi30))] = true
                                        fmt.Println(v26, v1, fi30)
                                        v33 := v4[32767]
                                        _ = v33
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = v23
                                    }
                                    _ = v31
                                } else {
                                    v23 += fi25
                                    var v34 uint8 = 131
                                    v23 += (v23 - (-(fi30)))
                                    v22 = v22
                                    _ = v34
                                }
                                v23 -= 2147483647
                            }
                            _ = fi30
                        }
                    }
                default:
                    var v35 uint = v26
                    _ = v35
                }
            }
            if false {
                var v36 int8 = -1
                var v37 uint16 = 32767
                _ = v36
                _ = v37
            }
            _ = fi25
            _ = v26
        }
        var v38 int8 = 1
        for fi39 := 0; fi39 < 14; fi39++ {
            switch 42 {
            case 42:
                if ((v1 + v1) >= ("foo bar" + "")) {
                    var v40 int8 = v38
                    _ = v40
                }
            }
            _ = fi39
        }
        _ = v38
    }
    var v41 uint16 = 2
    v42 := GenS1{F0: 42, F1: 6}
    v43 := GenS1{F0: 2147483647, F1: int8(v41)}
    for fi44 := 0; fi44 < 16; fi44++ {
        fmt.Println(v1, v41)
        if false {
            v23 = fi44
            switch 1 {
            case 4294967295:
                switch (v1 + v1) {
                case "":
                    var v45 float64 = 0.0
                    _ = v45
                case "a":
                    if true {
                        v23 = v23
                        if false {
                            for fi46 := 0; fi46 < 18; fi46++ {
                                if true {
                                    if (true && (fi46 == 9223372036854775807)) {
                                        v47 := GenS4{F0: -11, F1: (v41 | v41)}
                                        _ = v47
                                    } else {
                                        if (true && true) {
                                            switch (v1 + ((("foo bar" + v1) + (v1 + v1)) + v1)) {
                                            case "foo bar":
                                                switch fi46 {
                                                case -676:
                                                    fmt.Println(fi44, fi46, v22)
                                                case -228:
                                                    if true {
                                                        var v48 bool = ((!(!true)) || true)
                                                        _ = v48
                                                    }
                                                default:
                                                    v1 = v1
                                                }
                                                if (42 > 42) {
                                                    var v49 bool = false
                                                    _ = v49
                                                }
                                                if true {
                                                    v50 := GenS3{F0: v41}
                                                    v51 := GenS1{F0: 2147483647, F1: 73}
                                                    _ = v50
                                                    _ = v51
                                                }
                                            case "a":
                                                var v52 int32 = v22
                                                _ = v52
                                            }
                                            v23 = (fi44 / ((0) | 1))
                                        } else {
                                            fmt.Println(v22, v23)
                                            switch v23 {
                                            case -1:
                                                v53 := GenS5{F0: 2, F1: 91, F2: "foo bar"}
                                                _ = v53
                                            case 33:
                                                switch ("foo bar" + v1) {
                                                case "":
                                                    v54 := GenS2{F0: 42, F1: (!(84 < 128)), F2: -32768}
                                                    _ = v54
                                                }
                                                v55 := GenS6{F0: 1, F1: (!true), F2: (fi46 % (((fi46 / ((9223372036854775807) | 1))) | 1))}
                                                _ = v55
                                            case 0:
                                                var v56 float64 = 0.1
                                                _ = v56
                                            default:
                                                v41 -= v41
                                            }
                                            v1 = ((v1 + "foo bar") + "\n\t")
                                        }
                                        v23 = fi46
                                        var v57 uint = 424
                                        _ = v57
                                    }
                                } else {
                                    v22 = v22
                                    if (true && (1 <= 65535)) {
                                        switch v1 {
                                        case "":
                                            switch v23 {
                                            case 0:
                                                v58 := GenS5{F0: 1, F1: 93, F2: v1}
                                                _ = v58
                                            case 857:
                                                v59 := GenS1{F0: 147, F1: 0}
                                                _ = v59
                                            case 318:
                                                var v60 int = fi46
                                                _ = v60
                                            }
                                            v23 = (v23 % ((4294967295) | 1))
                                            switch (v23 ^ (fi46 * (fi46 * (fi46 & 32768)))) {
                                            case 2147483647:
                                                var v61 int16 = 185
                                                _ = v61
                                            default:
                                                v41 = (v41 % ((v41) | 1))
                                            }
                                            v62 := GenS7{F0: (v22 >> ((65535) & 63)), F1: 699}
                                            _ = v62
                                        case "\n\t":
                                            v22 += (v22 & (^(v22)))
                                        case "hello":
                                            var v63 float32 = 1.5
                                            _ = v63
                                        default:
                                            v1 = (v1 + "a")
                                        }
                                        v64 := GenS5{F0: 555, F1: 527, F2: "\n\t"}
                                        _ = v64
                                    }
                                }
                                if true {
                                    v65 := v3[(("" + v1) + ("hello" + (v1 + "\n\t")))]
                                    v66 := GenS6{F0: 1000, F1: true, F2: v23}
                                    _ = v65
                                    _ = v66
                                }
                                v67 := GenS3{F0: (v41 ^ (v41 << ((2) & 63)))}
                                _ = fi46
                                _ = v67
                            }
                            if (9223372036854775807 < 0) {
                                var v68 int = 162
                                for fi69 := 0; fi69 < 16; fi69++ {
                                    var v70 int32 = 0
                                    if true {
                                        if (2 != 2147483647) {
                                            v71 := GenS3{F0: (v41 >> ((0) & 63))}
                                            if (-2.25 < 1.5) {
                                                var v72 string = (v1 + v1)
                                                var v73 bool = (2147483647 <= 412)
                                                switch -1 {
                                                case 2147483647:
                                                    _ = len(v2)
                                                case -401:
                                                    switch fi69 {
                                                    case 9223372036854775807:
                                                        if v73 {
                                                            switch v72 {
                                                            case "a":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                        } else {
                                                            if ((!((v72 + (v72 + (v1 + v1))) == "hello")) || v73) {
                                                                if (!false) {
                                                                    fmt.Println(fi69)
                                                                }
                                                            }
                                                            _ = len(v2)
                                                        }
                                                    case 0:
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v74 := GenS3{F0: (v41 - v41)}
                                                            _ = v74
                                                        case "\n\t":
                                                            var v75 uint = 0
                                                            _ = v75
                                                        default:
                                                            v3[""] = ""
                                                        }
                                                    case 127:
                                                        var v76 int = fi44
                                                        _ = v76
                                                    }
                                                    v23 = -295
                                                case 127:
                                                    var v77 string = v72
                                                    _ = v77
                                                }
                                                _ = v72
                                                _ = v73
                                            } else {
                                                v70 += (v70 / (((v70 - -2147483648)) | 1))
                                                if len(v2) > 0 {
                                                    v2[0] = v23
                                                }
                                            }
                                            _ = v71
                                        }
                                        v41 += v41
                                        switch v1 {
                                        case "":
                                            var v78 int32 = v22
                                            _ = v78
                                        default:
                                            switch v1 {
                                            case "foo bar":
                                                v79 := GenS8{F0: (v22 % ((genFn1()) | 1)), F1: false}
                                                _ = v79
                                            case "a":
                                                v80 := GenS6{F0: -1, F1: false, F2: 32768}
                                                _ = v80
                                            }
                                        }
                                    }
                                    _ = fi69
                                    _ = v70
                                }
                                for fi81 := 0; fi81 < 9; fi81++ {
                                    if false {
                                        fmt.Println(v41, v41, fi81)
                                        v23 -= (^(fi44))
                                        v82 := GenS5{F0: 4294967295, F1: 69, F2: "a"}
                                        var v83 int32 = v22
                                        _ = v82
                                        _ = v83
                                    }
                                    _ = fi81
                                }
                                _ = v68
                            } else {
                                v84 := GenS8{F0: (v22 / (((v22 << ((1) & 63))) | 1)), F1: (!(1 <= 42))}
                                var v85 bool = true
                                for fi86 := 0; fi86 < 8; fi86++ {
                                    var v87 bool = v85
                                    var v88 float64 = ((3.14159 - ((-2.25 + 1.5) / (((3.14159 - 3.14159) + (((float64((float32(0.0) + 1.5)) / 1.5) * 1.5) * (0.0 / 1.5))) + 3.14159))) / 1.5)
                                    var v89 bool = (!(uint32(v41) <= 4294967295))
                                    _ = fi86
                                    _ = v87
                                    _ = v88
                                    _ = v89
                                }
                                _ = v84
                                _ = v85
                            }
                            var v90 int = (fi44 >> ((65535) & 63))
                            _ = v90
                        } else {
                            _ = len(v4)
                        }
                        for fi91 := 0; fi91 < 11; fi91++ {
                            v41 -= (v41 >> ((42) & 63))
                            v92 := GenS6{F0: -1, F1: (-1 >= int8(v22)), F2: fi91}
                            v41 = (v41 | v41)
                            _ = fi91
                            _ = v92
                        }
                    } else {
                        fmt.Println(v22, v23, v1)
                        v93 := GenS1{F0: 42, F1: 74}
                        _ = v93
                    }
                case "hello":
                    v41 = 2
                default:
                    switch (-(fi44)) {
                    case 9223372036854775807:
                        fmt.Println(v23, v1)
                    case -263:
                        var v94 int8 = 10
                        _ = v94
                    case -402:
                        for fi95 := 0; fi95 < 9; fi95++ {
                            v1 = (v1 + v1)
                            _ = fi95
                        }
                    }
                }
            case 521:
                v41 = v41
            case -9223372036854775808:
                v96 := GenS1{F0: 1, F1: -128}
                _ = v96
            }
            fmt.Println(fi44)
        } else {
            var v97 int16 = -1
            v98 := GenS9{F0: 9223372036854775807, F1: -128, F2: v1}
            if (!true) {
                v1 = (((v1 + "") + "hello") + v1)
                for fi99 := 0; fi99 < 8; fi99++ {
                    v100 := GenS10{F0: (v1 + "a"), F1: v22, F2: 21}
                    v101 := GenS7{F0: genFn1(), F1: v97}
                    v41 = (v41 * 255)
                    _ = fi99
                    _ = v100
                    _ = v101
                }
                for fi102 := 0; fi102 < 11; fi102++ {
                    v2 = append(v2, 2147483647)
                    v23 = (v23 * fi44)
                    _ = fi102
                }
            } else {
                if (!true) {
                    fmt.Println(v97, v1, v97)
                } else {
                    var v103 int16 = v97
                    var v104 int = fi44
                    if (v22 <= v22) {
                        fmt.Println(v104, v23)
                        fmt.Println(v104, v104)
                    }
                    _ = v103
                    _ = v104
                }
                v1 = ""
            }
            _ = v97
            _ = v98
        }
        v105 := GenS11{F0: 100}
        _ = fi44
        _ = v105
    }
    for fi106 := 0; fi106 < 10; fi106++ {
        v1 = v1
        var v107 uint32 = 4294967295
        switch (fi106 ^ fi106) {
        case 481:
            for fi108 := 0; fi108 < 7; fi108++ {
                if (848 != 0) {
                    var v109 int32 = v22
                    _ = v109
                }
                _ = fi108
            }
        default:
            var v110 uint = 0
            _ = v110
        }
        _ = fi106
        _ = v107
    }
    var v111 uint64 = 923
    var v112 int8 = 0
    v22 = genFn1()
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v111
    _ = v112
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
