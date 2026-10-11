package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 string }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int; F1 int32; F2 uint64 }
type GenS4 struct { F0 int64; F1 int64; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS7 struct { F0 string; F1 int; F2 uint32 }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 int16 }
func genFn0(p0 uint32, p1 string) int32 {
    var v6 uint64 = 18446744073709551615
    for fi7 := 0; fi7 < 17; fi7++ {
        if false {
            switch "" {
            case "":
                fmt.Println(p1, fi7)
            case "foo bar":
                var v8 string = ((p1 + p1) + "hello")
                _ = v8
            case "hello":
                fmt.Println(v6)
            }
            var v9 int = 199
            v10 := GenS1{F0: 1, F1: "a", F2: "a"}
            _ = v10
            _ = v9
        }
        var v11 int16 = 42
        v6 -= (v6 << ((0) & 63))
        _ = fi7
        _ = v11
    }
    switch 9223372036854775807 {
    case 1:
        v12 := GenS1{F0: -1, F1: p1, F2: ""}
        _ = v12
    case 32767:
        var v13 bool = true
        _ = v13
    case -9223372036854775808:
        if false {
            v6 += (v6 + v6)
        } else {
            v14 := []int{128, -9223372036854775808, 128}
            p0 = p0
            if (v6 <= (v6 >> ((18446744073709551615) & 63))) {
                if true {
                    var v15 int64 = -9223372036854775808
                    var v16 bool = ((0.0 > (3.14159 * (((((1.5 + (((1e10 * 3.14159) - -2.25) / float32(1e10))) - (1e10 + (0.1 - 0.1))) / 1.0) / (0.1 + 0.1)) - 3.14159))) && false)
                    _ = v15
                    _ = v16
                }
            }
            _ = v14
        }
    default:
        var v17 int32 = 42
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v6
    return 932
}

func main() {
    var v1 float32 = (0.0 * -2.25)
    var v2 string = "a"
    v3 := make(map[int]uint32)
    var v4 int = 958
    var v5 int32 = -1
    switch v4 {
    case -2147483648:
        _ = len(v3)
    case 42:
        _ = len(v3)
    case 730:
        fmt.Println(v5, v2)
    default:
        var v18 float64 = 1e10
        _ = v18
    }
    v1 = 1.5
    v19 := GenS2{F0: 0, F1: 75}
    if true {
        _ = len(v3)
        if true {
            fmt.Println(v1, v5)
            var v20 int32 = -2147483648
            _ = v20
        } else {
            var v21 float32 = v1
            _ = v21
        }
        switch v2 {
        case "hello":
            var v22 uint32 = 4294967295
            _ = v22
        default:
            for fi23 := 0; fi23 < 8; fi23++ {
                fmt.Println(v5, v5, v5)
                var v24 uint32 = 1
                for fi25 := 0; fi25 < 3; fi25++ {
                    switch v2 {
                    case "a":
                        if ((^(v24)) < v24) {
                            if true {
                                v4 = -1
                                var v26 bool = true
                                _ = v26
                            } else {
                                switch v2 {
                                case "hello":
                                    var v27 string = v2
                                    _ = v27
                                case "\n\t":
                                    v28 := GenS2{F0: v5, F1: 2}
                                    v4 -= (^(fi23))
                                    _ = v28
                                case "foo bar":
                                    v5 -= 42
                                }
                            }
                            v29 := GenS2{F0: v5, F1: 4294967295}
                            _ = v29
                        } else {
                            if false {
                                if true {
                                    v5 -= (v5 | (v5 - (v5 >> ((4294967295) & 63))))
                                } else {
                                    var v30 float32 = (0.0 + ((v1 / v1) - float32(1.5)))
                                    if true {
                                        v5 += v5
                                        v31 := GenS1{F0: -9223372036854775808, F1: (v2 + ((v2 + v2) + "hello")), F2: v2}
                                        v32 := GenS3{F0: (fi23 % ((4294967295) | 1)), F1: v5, F2: 605}
                                        _ = v31
                                        _ = v32
                                    } else {
                                        if false {
                                            var v33 bool = false
                                            var v34 int32 = (v5 % (((v5 - genFn0(2, v2))) | 1))
                                            _ = v33
                                            _ = v34
                                        }
                                        v35 := GenS4{F0: 898, F1: -9223372036854775808, F2: v5}
                                        _ = v35
                                    }
                                    v36 := []int{v4, fi23, (v4 & fi23)}
                                    _ = v30
                                    _ = v36
                                }
                            } else {
                                var v37 int64 = -9223372036854775808
                                var v38 uint64 = 4294967295
                                _ = v37
                                _ = v38
                            }
                        }
                    case "foo bar":
                        v39 := GenS5{F0: v2}
                        _ = v39
                    case "\n\t":
                        v40 := []int{9223372036854775807, v4, (fi23 ^ (v4 % ((v4) | 1))), (v4 >> ((997) & 63))}
                        _ = v40
                    }
                    var v41 bool = (!true)
                    var v42 int64 = -1
                    _ = fi25
                    _ = v41
                    _ = v42
                }
                _ = fi23
                _ = v24
            }
        }
    } else {
        for fi43 := 0; fi43 < 3; fi43++ {
            if false {
                v4 -= v4
            }
            v2 = v2
            v5 = v5
            _ = fi43
        }
        v4 = 2147483647
    }
    fmt.Println(v4, v5)
    switch (v2 + v2) {
    case "a":
        for fi44 := 0; fi44 < 2; fi44++ {
            if (!true) {
                fmt.Println(fi44, fi44, fi44)
                for fi45 := 0; fi45 < 8; fi45++ {
                    if false {
                        var v46 uint16 = 0
                        if true {
                            v47 := []int{(v4 & 1), (v4 - (v4 >> ((429) & 63)))}
                            v48 := GenS1{F0: 1, F1: (("hello" + (v2 + v2)) + ("foo bar" + v2)), F2: v2}
                            _ = v47
                            _ = v48
                        }
                        if (!(3.14159 != 0.1)) {
                            v49 := v3[(fi44 >> ((2) & 63))]
                            delete(v3, 32767)
                            _ = v49
                        }
                        _ = v46
                    }
                    fmt.Println(fi45, v1, v4)
                    _ = fi45
                }
                v5 -= genFn0(2147483648, ("hello" + v2))
            } else {
                v50 := []int{(-(fi44)), (fi44 | (v4 - 634)), (^(v4))}
                _ = v50
            }
            for fi51 := 0; fi51 < 18; fi51++ {
                var v52 uint = 255
                var v53 int16 = -32768
                _ = fi51
                _ = v52
                _ = v53
            }
            var v54 uint64 = 482
            _ = fi44
            _ = v54
        }
    }
    if true {
        fmt.Println(v2)
        v3[127] = 65535
        v55 := GenS1{F0: 0, F1: v2, F2: "hello"}
        _ = v55
    } else {
        if true {
            v2 = v2
        } else {
            v4 += -990
        }
    }
    fmt.Println(v2, v2, v4)
    var v56 uint32 = 542
    if true {
        v56 -= 950
        var v57 uint8 = 50
        for fi58 := 0; fi58 < 13; fi58++ {
            if (!true) {
                if (42 == 9223372036854775807) {
                    v57 = v57
                    if false {
                        v56 += v56
                        for fi59 := 0; fi59 < 8; fi59++ {
                            v4 -= fi58
                            switch v2 {
                            case "a":
                                v60 := []int{v4, fi58}
                                _ = v60
                            }
                            _ = fi59
                        }
                        var v61 uint16 = 676
                        _ = v61
                    }
                }
                v4 = (fi58 << ((65535) & 63))
            } else {
                v5 -= -1
                if true {
                    if (!true) {
                        v62 := GenS3{F0: v4, F1: (v5 * 971), F2: 2}
                        delete(v3, 128)
                        switch (-(fi58)) {
                        case 4294967295:
                            v4 += 144
                        case 440:
                            var v63 uint = 744
                            _ = v63
                        case 32767:
                            for fi64 := 0; fi64 < 9; fi64++ {
                                var v65 uint8 = (v57 | (v57 - uint8(v56)))
                                if false {
                                    if (true && (v2 >= v2)) {
                                        var v66 string = "a"
                                        v67 := GenS5{F0: ((v66 + v2) + "")}
                                        switch (fi64 ^ 2147483647) {
                                        case -9223372036854775808:
                                            v68 := GenS6{F0: -1, F1: 255, F2: 2}
                                            _ = v68
                                        default:
                                            switch "" {
                                            case "foo bar":
                                                v2 = (v66 + v66)
                                            }
                                        }
                                        _ = v66
                                        _ = v67
                                    } else {
                                        var v69 string = v2
                                        _ = v69
                                    }
                                }
                                _ = fi64
                                _ = v65
                            }
                        default:
                            var v70 uint16 = 1
                            _ = v70
                        }
                        _ = v62
                    } else {
                        v71 := GenS3{F0: (^(v4)), F1: v5, F2: 978}
                        _ = v71
                    }
                }
                for fi72 := 0; fi72 < 3; fi72++ {
                    if (127 > int8(v5)) {
                        var v73 bool = (!((((v2 + ((v2 + "foo bar") + "hello")) + "hello") + v2) == ((("foo bar" + v2) + v2) + v2)))
                        _ = v73
                    } else {
                        switch v2 {
                        case "foo bar":
                            v4 -= (v4 | (fi58 + -2147483648))
                        }
                    }
                    v74 := GenS7{F0: v2, F1: fi72, F2: (^(v56))}
                    if true {
                        var v75 uint8 = v57
                        var v76 int16 = 806
                        _ = v75
                        _ = v76
                    }
                    switch fi72 {
                    case 42:
                        v4 = 1
                        v1 = v1
                    default:
                        var v77 string = ("" + (v2 + v2))
                        var v78 uint64 = 18446744073709551615
                        var v79 int8 = 0
                        _ = v77
                        _ = v78
                        _ = v79
                    }
                    _ = fi72
                    _ = v74
                }
            }
            if false {
                for fi80 := 0; fi80 < 20; fi80++ {
                    fmt.Println(fi58)
                    fmt.Println(v5)
                    _ = fi80
                }
            }
            _ = fi58
        }
        _ = v57
    }
    if true {
        var v81 int64 = -9223372036854775808
        _ = v81
    } else {
        v56 += v56
    }
    if (-1 != 42) {
        var v82 float64 = (((-2.25 - 0.1) * 3.14159) + 0.0)
        v83 := GenS8{F0: (v56 - 2147483648), F1: 0}
        _ = v82
        _ = v83
    } else {
        switch -149 {
        case 839:
            _ = len(v3)
        case -2147483648:
            if true {
                v84 := GenS2{F0: genFn0((v56 - (v56 << ((1) & 63))), ""), F1: 4294967295}
                for fi85 := 0; fi85 < 8; fi85++ {
                    fmt.Println(v4, v4)
                    _ = fi85
                }
                v56 += v56
                _ = v84
            }
        }
    }
    v86 := GenS9{F0: -32768}
    if false {
        var v87 int8 = 103
        _ = v87
    } else {
        if false {
            v5 = (v5 | genFn0(v56, "hello"))
            var v88 uint = 773
            _ = v88
        }
        if false {
            v89 := []int{v4}
            v2 = v2
            v5 = 0
            _ = v89
        }
    }
    v56 = (v56 & (v56 - (^(v56))))
    var v90 int8 = -128
    if ((v56 * (v56 % ((v56) | 1))) != 2147483647) {
        v91 := v3[v4]
        var v92 uint = 255
        _ = v91
        _ = v92
    }
    if (float64((((1.5 + 0.0) / (1e10 / float32((((3.14159 / (3.14159 + float64(((1.5 * v1) / float32(float64(v1)))))) - -2.25) / 1.5)))) / v1)) > 1e10) {
        v93 := GenS5{F0: "a"}
        v94 := GenS3{F0: int(v90), F1: v5, F2: 0}
        for fi95 := 0; fi95 < 17; fi95++ {
            var v96 float32 = v1
            _ = fi95
            _ = v96
        }
        _ = v93
        _ = v94
    } else {
        for fi97 := 0; fi97 < 7; fi97++ {
            var v98 string = v2
            v5 -= genFn0((v56 / ((2147483647) | 1)), v98)
            if false {
                fmt.Println(fi97)
                if false {
                    var v99 int16 = -1
                    var v100 uint64 = 42
                    _ = v100
                    _ = v99
                } else {
                    var v101 int32 = (v5 / (((v5 << ((65535) & 63))) | 1))
                    for fi102 := 0; fi102 < 17; fi102++ {
                        switch (v4 - (^(v4))) {
                        case 128:
                            v103 := []int{}
                            _ = v103
                        case -1:
                            v101 = 42
                        }
                        _ = fi102
                    }
                    v98 = ("hello" + v2)
                    _ = v101
                }
            } else {
                v104 := GenS3{F0: v4, F1: -887, F2: 1}
                if true {
                    fmt.Println(v4, fi97)
                    switch int(v90) {
                    case 2147483647:
                        v105 := GenS7{F0: v2, F1: (v4 / (((^(fi97))) | 1)), F2: v56}
                        _ = v105
                    case -174:
                        switch "a" {
                        case "foo bar":
                            if true {
                                v2 = v98
                                var v106 int = 32768
                                _ = v106
                            } else {
                                v107 := []int{fi97, v4}
                                _ = v107
                            }
                        case "\n\t":
                            v3[v4] = v56
                        }
                    case 128:
                        fmt.Println(v4)
                    default:
                        fmt.Println(fi97, v90)
                    }
                }
                _ = v104
            }
            _ = fi97
            _ = v98
        }
        v2 = "\n\t"
        for fi108 := 0; fi108 < 16; fi108++ {
            switch fi108 {
            case 42:
                fmt.Println(fi108, v4)
            }
            var v109 uint8 = 127
            _ = fi108
            _ = v109
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v86
    _ = v90
}
