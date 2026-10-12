package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int64; F2 int }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 uint16; F1 int; F2 int64 }
type GenS7 struct { F0 int16; F1 uint64; F2 uint }
type GenS8 struct { F0 bool; F1 uint; F2 bool }
type GenS9 struct { F0 uint16; F1 int32; F2 string }
func genFn0(p0 int, p1 int64, p2 string) uint16 {
    var v7 uint16 = 0
    var v8 int16 = -1
    v9 := GenS1{F0: p0, F1: 1, F2: 864}
    if ((^(p1)) <= -16) {
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := []int{p0, (^(p0)), p0, (^(fi10))}
            _ = fi10
            _ = v11
        }
    } else {
        if false {
            var v12 int32 = -2147483648
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (v7 ^ 1)
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -32768
    var v3 int64 = 942
    var v4 bool = ((false && (1 != -384)) || false)
    v5 := make(map[int]uint32)
    var v6 int32 = 2147483647
    v13 := GenS2{F0: 764}
    var v14 int = -550
    var v15 float32 = 1e10
    var v16 uint64 = 42
    if (971 <= genFn0((v14 / (((v14 | int(v3))) | 1)), (v3 % ((9223372036854775807) | 1)), "\n\t")) {
        fmt.Println(v15, v14)
    } else {
        switch "\n\t" {
        case "foo bar":
            switch v14 {
            case 32768:
                if (true || false) {
                    v4 = v4
                    var v17 int32 = v6
                    v16 += 42
                    _ = v17
                } else {
                    if (((float32((0.1 - (3.14159 - (0.1 / 1.0)))) - v15) - v15) > (float32(0.0) * -2.25)) {
                        v18 := GenS3{F0: v6, F1: v3, F2: (v14 << ((65535) & 63))}
                        var v19 int64 = (v3 << ((0) & 63))
                        _ = v18
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            switch ("foo bar" + ("hello" + "\n\t")) {
                            case "a":
                                var v21 uint = 585
                                _ = v21
                            }
                            _ = fi20
                        }
                        var v22 float32 = v15
                        for fi23 := 0; fi23 < 13; fi23++ {
                            v24 := GenS4{F0: true, F1: v3}
                            _ = len(v1)
                            v25 := GenS4{F0: true, F1: (v3 & v3)}
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                    }
                    var v26 uint8 = 255
                    if true {
                        v4 = ((v3 + v3) <= 42)
                        var v27 uint64 = v16
                        for fi28 := 0; fi28 < 9; fi28++ {
                            if ((4294967295 == 2147483648) || (false || (0 < 127))) {
                                if v4 {
                                    var v29 uint16 = 32767
                                    for fi30 := 0; fi30 < 2; fi30++ {
                                        var v31 uint8 = v26
                                        switch "a" {
                                        case "\n\t":
                                            v32 := GenS4{F0: false, F1: v3}
                                            _ = v32
                                        case "a":
                                            if v4 {
                                                v33 := v5[int(v3)]
                                                _ = v33
                                            } else {
                                                var v34 uint64 = (^(v27))
                                                _ = v34
                                            }
                                        }
                                        if (v2 == v2) {
                                            v35 := GenS1{F0: fi28, F1: 18446744073709551615, F2: 479}
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v36 bool = false
                                                _ = v36
                                            }
                                            v37 := []int{2147483647}
                                            _ = v35
                                            _ = v37
                                        } else {
                                            fmt.Println(fi30, v27)
                                        }
                                        _ = fi30
                                        _ = v31
                                    }
                                    if v4 {
                                        switch fi28 {
                                        case -1:
                                            v5[128] = 4294967295
                                        case 697:
                                            v4 = false
                                        default:
                                            v3 += v3
                                        }
                                        v2 = (v2 >> ((42) & 63))
                                        v16 += (^(v16))
                                    }
                                    _ = v29
                                } else {
                                    var v38 uint32 = 0
                                    if v4 {
                                        var v39 uint32 = 0
                                        var v40 int64 = (v3 >> ((18446744073709551615) & 63))
                                        v15 = v15
                                        _ = v39
                                        _ = v40
                                    }
                                    _ = v38
                                }
                                v41 := GenS1{F0: fi28, F1: 1, F2: 0}
                                v42 := []int{}
                                _ = v41
                                _ = v42
                            } else {
                                for fi43 := 0; fi43 < 7; fi43++ {
                                    if (-1 != (v3 % ((-38) | 1))) {
                                        v15 = -2.25
                                        var v44 float32 = v15
                                        _ = v44
                                    }
                                    var v45 bool = true
                                    v15 -= v15
                                    _ = fi43
                                    _ = v45
                                }
                                v46 := GenS3{F0: (v6 - v6), F1: (v3 % ((v3) | 1)), F2: -2147483648}
                                _ = v46
                            }
                            var v47 string = ""
                            _ = fi28
                            _ = v47
                        }
                        _ = v27
                    }
                    _ = v26
                }
            default:
                if (42 >= 42) {
                    v3 += 0
                    if (0 != 1) {
                        var v48 int32 = int32(v16)
                        v49 := []int{v14, (v14 | v14), v14}
                        _ = v48
                        _ = v49
                    }
                    v50 := []int{(v14 | (v14 >> ((65535) & 63))), 128, (^(v14))}
                    _ = v50
                }
            }
        case "hello":
            var v51 int16 = v2
            _ = v51
        case "\n\t":
            if (v15 <= (((v15 * v15) + (1e10 + (v15 * float32(3.14159)))) * ((v15 + ((v15 + v15) - 1.5)) / (v15 + v15)))) {
                v52 := GenS5{F0: (-(v2)), F1: (v14 << ((4294967295) & 63))}
                v53 := GenS6{F0: 397, F1: -2147483648, F2: (v3 / (((v3 | (v3 >> ((1) & 63)))) | 1))}
                _ = v52
                _ = v53
            }
        default:
            switch -9223372036854775808 {
            case 850:
                v15 = (float32(0.1) - 1e10)
            case 0:
                v54 := GenS3{F0: (v6 * v6), F1: v3, F2: 878}
                _ = v54
            case -213:
                switch v14 {
                case -1:
                    fmt.Println(v14)
                case 32767:
                    var v55 uint = 2
                    _ = v55
                case -2147483648:
                    v56 := GenS1{F0: (v14 * (v14 / ((2147483647) | 1))), F1: 1, F2: 255}
                    _ = v56
                default:
                    var v57 uint64 = (v16 << ((423) & 63))
                    _ = v57
                }
            default:
                var v58 int32 = v6
                _ = v58
            }
        }
    }
    for fi59 := 0; fi59 < 14; fi59++ {
        switch int(v6) {
        case -9223372036854775808:
            if (!(2147483647 != (v6 + (v6 / ((v6) | 1))))) {
                var v60 int = fi59
                var v61 int8 = 42
                _ = v60
                _ = v61
            }
        case 4294967295:
            switch fi59 {
            case 32768:
                if (int(v3) <= (-(fi59))) {
                    delete(v5, 32768)
                    v62 := []int{fi59, (-(fi59))}
                    _ = v62
                } else {
                    v63 := GenS7{F0: (v2 - (v2 * (v2 * v2))), F1: (v16 ^ v16), F2: 389}
                    fmt.Println(v14)
                    fmt.Println(v4, fi59, v3)
                    _ = v63
                }
            case 127:
                if v4 {
                    v14 = (fi59 - (v14 ^ (fi59 >> ((1) & 63))))
                } else {
                    var v64 float64 = -2.25
                    _ = v64
                }
            case -9223372036854775808:
                var v65 int = (^(v14))
                _ = v65
            default:
                delete(v1, (fi59 | fi59))
            }
        default:
            var v66 float32 = v15
            _ = v66
        }
        var v67 int8 = 80
        _ = fi59
        _ = v67
    }
    var v68 float32 = (float32(3.14159) - 0.0)
    v69 := []int{(v14 & v14), (v14 % ((v14) | 1))}
    if v4 {
        var v70 uint8 = 255
        fmt.Println(v70, v6)
        v68 = v15
        _ = v70
    } else {
        var v71 int64 = v3
        var v72 uint16 = 893
        for fi73 := 0; fi73 < 5; fi73++ {
            fmt.Println(v68, v6)
            var v74 float32 = (v68 * (((v15 / v68) - (v15 - v68)) / (v15 / (float32(((float64(v15) * ((float64(v15) / float64(float32(1e10))) + 0.1)) / (((0.0 - 1.5) * (((1e10 + 0.1) + 1e10) + -2.25)) - (1.5 + (1e10 / (0.1 + float64((-2.25 * v15)))))))) - v68))))
            if v4 {
                for fi75 := 0; fi75 < 14; fi75++ {
                    if v4 {
                        if len(v69) >= 2 {
                            v69 = v69[1:len(v69)]
                        }
                        delete(v5, (fi75 - (v14 * (fi75 % ((629) | 1)))))
                    }
                    if false {
                        delete(v1, int(v71))
                        if false {
                            fmt.Println(v4, v15, v3)
                            if v4 {
                                v5[127] = 1
                                fmt.Println(v3)
                                fmt.Println(v4)
                            } else {
                                fmt.Println(v71)
                                var v76 int64 = (v3 + 1)
                                _ = v76
                            }
                            switch "\n\t" {
                            case "a":
                                var v77 bool = (v6 < (-(v6)))
                                _ = v77
                            default:
                                v68 -= v74
                            }
                        } else {
                            switch "\n\t" {
                            case "":
                                if (v4 || (((v72 >> ((65535) & 63)) <= (^(v72))) || v4)) {
                                    v14 = 163
                                    var v78 bool = false
                                    _ = v78
                                }
                            case "\n\t":
                                var v79 int64 = v3
                                _ = v79
                            case "hello":
                                if ("" > ("\n\t" + ("hello" + ""))) {
                                    if (v4 && v4) {
                                        var v80 string = ""
                                        _ = v80
                                    } else {
                                        v71 += v71
                                    }
                                    var v81 string = "foo bar"
                                    switch v81 {
                                    case "":
                                        delete(v5, (v14 ^ (-(v14))))
                                    }
                                    _ = v81
                                } else {
                                    var v82 int64 = (v71 * (v3 + (v3 >> ((2) & 63))))
                                    switch "\n\t" {
                                    case "\n\t":
                                        if v4 {
                                            v74 -= (1e10 * v15)
                                            var v83 float32 = v15
                                            _ = v83
                                        }
                                    }
                                    fmt.Println(v2, v68, v71)
                                    if v4 {
                                        _ = len(v69)
                                        _ = len(v69)
                                        v14 = fi73
                                    }
                                    fmt.Println(v3, v72, v2)
                                    switch "\n\t" {
                                    case "":
                                        var v84 int = -979
                                        _ = v84
                                    default:
                                        fmt.Println(v6, v2)
                                    }
                                    _ = v82
                                }
                                v74 = (float32(0.0) * v15)
                            }
                            fmt.Println(v3)
                            v68 += (v68 * ((v68 / (1e10 * v74)) / 1.5))
                        }
                        var v85 int16 = (v2 + (v2 >> ((65535) & 63)))
                        _ = v85
                    }
                    _ = fi75
                }
            } else {
                fmt.Println(v16)
                v15 += 1.5
                fmt.Println(fi73, v72)
            }
            _ = fi73
            _ = v74
        }
        _ = v71
        _ = v72
    }
    for fi86 := 0; fi86 < 3; fi86++ {
        var v87 int64 = v3
        if len(v69) > 0 {
            v69[0] = fi86
        }
        switch 9223372036854775807 {
        case -277:
            if ((v6 + -949) <= (v6 ^ v6)) {
                for fi88 := 0; fi88 < 3; fi88++ {
                    v2 = (v2 >> ((1) & 63))
                    _ = fi88
                }
                v69 = append(v69, -2147483648)
            } else {
                if true {
                    fmt.Println(v3, v14, v16)
                } else {
                    _ = len(v5)
                    fmt.Println(v15)
                    v6 -= (-(v6))
                }
                fmt.Println(v14, v4)
                for fi89 := 0; fi89 < 13; fi89++ {
                    var v90 int = (v14 >> ((4294967295) & 63))
                    var v91 int8 = -1
                    var v92 uint32 = 673
                    _ = fi89
                    _ = v90
                    _ = v91
                    _ = v92
                }
            }
        case -493:
            v93 := GenS8{F0: (!(!(!(!v4)))), F1: 0, F2: v4}
            _ = v93
        case 1:
            var v94 int64 = v3
            _ = v94
        default:
            v14 = (fi86 * 682)
        }
        _ = fi86
        _ = v87
    }
    v6 = (v6 - v6)
    v68 = float32(float64(v68))
    v4 = v4
    if v4 {
        v95 := GenS3{F0: (v6 >> ((0) & 63)), F1: v3, F2: 42}
        _ = v95
    }
    v4 = false
    switch (("foo bar" + "a") + ("\n\t" + "hello")) {
    case "hello":
        if v4 {
            v96 := GenS1{F0: (v14 / ((-9223372036854775808) | 1)), F1: 4294967295, F2: 23}
            v97 := GenS9{F0: 760, F1: v6, F2: "\n\t"}
            for fi98 := 0; fi98 < 13; fi98++ {
                if len(v69) >= 2 {
                    v69 = v69[1:len(v69)]
                }
                if ((v16 * (v16 & 4294967295)) >= v16) {
                    v6 -= (v6 - v6)
                    for fi99 := 0; fi99 < 8; fi99++ {
                        v6 = (v6 + 572)
                        if (v4 && (127 <= 100)) {
                            fmt.Println(v16, fi98)
                            fmt.Println(v15, v68, v16)
                        }
                        if (v4 && v4) {
                            v3 += (v3 / ((v3) | 1))
                            if v4 {
                                if (0 < 1) {
                                    fmt.Println(v68, fi98)
                                    var v100 uint64 = v16
                                    var v101 int32 = v6
                                    _ = v100
                                    _ = v101
                                } else {
                                    v15 -= ((v15 / 1.5) * (1e10 * 0.1))
                                    var v102 uint16 = 1
                                    _ = v102
                                }
                                v68 -= (v15 * v68)
                                v15 += (v68 / 1.5)
                                v2 = (^(v2))
                            } else {
                                var v103 string = "\n\t"
                                v104 := v1[-2147483648]
                                _ = v103
                                _ = v104
                            }
                            var v105 int16 = (v2 * v2)
                            _ = v105
                        }
                        _ = fi99
                    }
                    fmt.Println(v68, v3, v14)
                } else {
                    var v106 uint32 = 976
                    _ = v106
                }
                fmt.Println(v14, v14)
                _ = fi98
            }
            _ = v96
            _ = v97
        } else {
            for fi107 := 0; fi107 < 14; fi107++ {
                for fi108 := 0; fi108 < 7; fi108++ {
                    switch fi108 {
                    case 42:
                        v69 = append(v69, -1)
                    default:
                        v1[v14] = (-1 <= 100)
                    }
                    v109 := GenS1{F0: (fi108 * 127), F1: 18446744073709551615, F2: 2}
                    _ = fi108
                    _ = v109
                }
                _ = fi107
            }
            delete(v5, v14)
        }
    case "\n\t":
        v15 -= 1.5
    default:
        v16 += v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v69:", len(v69), v69)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
}
