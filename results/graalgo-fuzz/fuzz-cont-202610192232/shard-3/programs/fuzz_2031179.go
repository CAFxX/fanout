package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16; F1 int32; F2 uint64 }
type GenS8 struct { F0 bool; F1 string; F2 string }
type GenS9 struct { F0 uint; F1 int; F2 int16 }
type GenS10 struct { F0 int32; F1 uint; F2 int }
type GenS11 struct { F0 uint8 }
func genFn0(p0 int64) uint64 {
    v7 := []int{}
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: "", F1: 162}
        fmt.Println(p0, fi8)
        for fi10 := 0; fi10 < 17; fi10++ {
            if false {
                p0 -= 1
                var v11 uint8 = 1
                _ = v11
            }
            var v12 uint16 = 32767
            _ = fi10
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if (!true) {
        var v13 int32 = 2147483647
        if true {
            fmt.Println(p0, v13)
            switch -9223372036854775808 {
            case -9223372036854775808:
                v14 := GenS1{F0: ("" + ("" + ("a" + ("foo bar" + "")))), F1: 2}
                _ = v14
            case 128:
                switch ("hello" + "a") {
                case "a":
                    for fi15 := 0; fi15 < 20; fi15++ {
                        p0 = (p0 >> ((42) & 63))
                        for fi16 := 0; fi16 < 13; fi16++ {
                            v7 = append(v7, (fi16 * (fi15 >> ((18446744073709551615) & 63))))
                            switch (-(fi15)) {
                            case -116:
                                if (0 >= 18446744073709551615) {
                                    var v17 bool = (!true)
                                    _ = v17
                                } else {
                                    _ = len(v7)
                                }
                            case 539:
                                v18 := GenS1{F0: "", F1: 127}
                                _ = v18
                            }
                            switch "\n\t" {
                            case "\n\t":
                                if true {
                                    v13 += v13
                                    v19 := GenS1{F0: "a", F1: 127}
                                    _ = v19
                                }
                            }
                            _ = fi16
                        }
                        _ = fi15
                    }
                case "\n\t":
                    var v20 bool = (!false)
                    _ = v20
                default:
                    p0 = 9223372036854775807
                }
            default:
                switch "" {
                case "":
                    v13 += (^(v13))
                case "foo bar":
                    fmt.Println(v13)
                case "a":
                    p0 = (p0 ^ p0)
                default:
                    fmt.Println(p0, v13)
                }
            }
            if (!false) {
                if (760 == 65535) {
                    for fi21 := 0; fi21 < 8; fi21++ {
                        var v22 uint8 = 128
                        for fi23 := 0; fi23 < 2; fi23++ {
                            if ((("foo bar" > (("\n\t" + "hello") + (("" + "hello") + (((("" + "a") + "") + ("foo bar" + "foo bar")) + "a")))) || false) || (!false)) {
                                v22 -= (v22 | (v22 ^ (v22 >> ((2) & 63))))
                                var v24 uint64 = 667
                                fmt.Println(v22, fi23, v24)
                                _ = v24
                            } else {
                                var v25 bool = false
                                _ = v25
                            }
                            var v26 uint = 536
                            var v27 int32 = -2147483648
                            v13 += v27
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            v27 += v27
                            p0 += int64(v13)
                            _ = fi23
                            _ = v26
                            _ = v27
                        }
                        if true {
                            var v28 int = 0
                            v29 := GenS2{F0: true, F1: (p0 / ((-1) | 1))}
                            _ = v28
                            _ = v29
                        }
                        _ = fi21
                        _ = v22
                    }
                } else {
                    for fi30 := 0; fi30 < 9; fi30++ {
                        var v31 uint = 0
                        v32 := GenS1{F0: "hello", F1: 127}
                        var v33 int = (fi30 << ((295) & 63))
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    v7 = append(v7, 849)
                    p0 = 0
                }
                if (1 <= 445) {
                    var v34 int32 = (-(v13))
                    if false {
                        var v35 uint32 = 65535
                        v36 := GenS2{F0: (-1 < 42), F1: -359}
                        fmt.Println(v13, v35, v35)
                        _ = v35
                        _ = v36
                    } else {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    _ = v34
                } else {
                    fmt.Println(v13, p0)
                    var v37 int16 = -32768
                    v38 := GenS2{F0: true, F1: 9223372036854775807}
                    _ = v37
                    _ = v38
                }
                p0 += p0
            } else {
                var v39 int32 = v13
                v13 += (v13 % (((v39 - 1)) | 1))
                v40 := GenS3{F0: 26}
                _ = v39
                _ = v40
            }
        } else {
            var v41 float32 = 0.1
            _ = v41
        }
        _ = v13
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1() uint8 {
    v42 := []int{128, 4294967295, -2147483648}
    var v43 int16 = -957
    switch -1 {
    case 32767:
        var v44 float32 = (0.0 + float32(0.1))
        _ = v44
    case -276:
        var v45 int = 30
        _ = v45
    case -992:
        v43 = (v43 | v43)
    }
    var v46 string = "a"
    for fi47 := 0; fi47 < 5; fi47++ {
        v43 = (v43 | (v43 * (v43 + v43)))
        _ = fi47
    }
    _ = v42
    _ = v43
    _ = v46
    return 255
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    var v3 uint64 = 18446744073709551615
    var v4 float32 = ((((3.14159 / (0.0 - 0.1)) + 3.14159) / (float32(3.14159) / 1.0)) * (((0.1 * -2.25) - float32((0.1 * (1e10 / (float64(3.14159) + (1e10 - (1e10 + 0.0))))))) / 3.14159))
    var v5 bool = ((42 != 65535) && false)
    var v6 int8 = 116
    var v48 string = "foo bar"
    v5 = v5
    var v49 uint = 255
    if false {
        var v50 int32 = 0
        var v51 bool = false
        if false {
            var v52 int8 = 127
            if v51 {
                var v53 string = "a"
                for fi54 := 0; fi54 < 10; fi54++ {
                    v55 := GenS2{F0: (137 > 1), F1: 1}
                    var v56 int = (fi54 * (fi54 * fi54))
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                v1[("hello" + "")] = v53
                _ = v53
            }
            for fi57 := 0; fi57 < 7; fi57++ {
                fmt.Println(v4, v5, v52)
                if v51 {
                    v58 := GenS1{F0: v48, F1: 255}
                    v59 := GenS2{F0: true, F1: -9223372036854775808}
                    _ = v58
                    _ = v59
                } else {
                    switch (v48 + (("foo bar" + v48) + (v48 + v48))) {
                    case "\n\t":
                        if v5 {
                            if (!v51) {
                                switch "\n\t" {
                                case "foo bar":
                                    switch v48 {
                                    case "\n\t":
                                        var v60 uint16 = 313
                                        _ = v60
                                    case "foo bar":
                                        v50 += (-(v50))
                                    case "a":
                                        if v5 {
                                            v61 := GenS1{F0: v48, F1: genFn1()}
                                            var v62 uint32 = 2
                                            _ = v61
                                            _ = v62
                                        }
                                    }
                                case "\n\t":
                                    var v63 int64 = 42
                                    _ = v63
                                }
                                var v64 uint8 = 127
                                _ = v64
                            } else {
                                v51 = v5
                            }
                            switch -460 {
                            case 2147483647:
                                fmt.Println(fi57)
                            case 1:
                                if v5 {
                                    v65 := GenS4{F0: v5}
                                    _ = v65
                                } else {
                                    for fi66 := 0; fi66 < 10; fi66++ {
                                        v50 -= 889
                                        v67 := GenS2{F0: ((311 > 455) || false), F1: 42}
                                        _ = fi66
                                        _ = v67
                                    }
                                    v49 += (v49 ^ (v49 + 4294967295))
                                }
                            case 42:
                                fmt.Println(v52)
                            }
                        } else {
                            v68 := v1[v48]
                            _ = v68
                        }
                    case "a":
                        var v69 uint64 = (^(v3))
                        _ = v69
                    default:
                        v51 = (fi57 != (fi57 / ((9223372036854775807) | 1)))
                    }
                    v3 += genFn0(9223372036854775807)
                }
                _ = fi57
            }
            _ = v52
        } else {
            v49 = (v49 % (((^(v49))) | 1))
            for fi70 := 0; fi70 < 13; fi70++ {
                var v71 int8 = (-(v6))
                _ = fi70
                _ = v71
            }
        }
        _ = v50
        _ = v51
    }
    switch 9223372036854775807 {
    case 128:
        var v72 int32 = 0
        _ = v72
    case 4294967295:
        switch v48 {
        case "":
            v73 := v2[(v48 + "\n\t")]
            _ = v73
        case "a":
            if false {
                fmt.Println(v5, v5, v3)
                var v74 uint64 = 2
                _ = v74
            } else {
                delete(v1, v48)
            }
        case "hello":
            if (42 > 0) {
                var v75 bool = v5
                _ = v75
            } else {
                for fi76 := 0; fi76 < 7; fi76++ {
                    v77 := []int{int(v49), fi76}
                    for fi78 := 0; fi78 < 17; fi78++ {
                        var v79 int32 = 2147483647
                        switch (v48 + "a") {
                        case "":
                            v3 -= 602
                        default:
                            if (!(v6 != 127)) {
                                var v80 int8 = -128
                                v81 := GenS3{F0: (fi78 >> ((1) & 63))}
                                _ = v80
                                _ = v81
                            } else {
                                v82 := GenS5{F0: -67}
                                _ = v82
                            }
                        }
                        v77 = append(v77, fi76)
                        _ = fi78
                        _ = v79
                    }
                    _ = fi76
                    _ = v77
                }
                v83 := GenS6{F0: 32767}
                for fi84 := 0; fi84 < 18; fi84++ {
                    for fi85 := 0; fi85 < 11; fi85++ {
                        var v86 string = ((v48 + "a") + v48)
                        v87 := v2[v48]
                        _ = fi85
                        _ = v86
                        _ = v87
                    }
                    if v5 {
                        v49 = v49
                        var v88 int8 = -1
                        _ = v88
                    } else {
                        if (v5 && (!true)) {
                            var v89 int64 = -9223372036854775808
                            _ = len(v2)
                            _ = v89
                        } else {
                            v90 := GenS3{F0: (^(fi84))}
                            v91 := GenS7{F0: -32768, F1: 0, F2: (^(v3))}
                            _ = v90
                            _ = v91
                        }
                        switch v48 {
                        case "\n\t":
                            switch (fi84 / ((900) | 1)) {
                            case 128:
                                v6 = v6
                            case 2147483647:
                                _ = len(v1)
                            case 9223372036854775807:
                                v92 := []int{-2147483648, fi84, fi84}
                                _ = v92
                            default:
                                var v93 uint64 = (v3 / ((v3) | 1))
                                _ = v93
                            }
                        default:
                            var v94 float64 = 3.14159
                            _ = v94
                        }
                        fmt.Println(v3, v6)
                    }
                    var v95 int64 = -9223372036854775808
                    _ = fi84
                    _ = v95
                }
                _ = v83
            }
        default:
            switch 42 {
            case 2147483647:
                var v96 uint8 = 178
                _ = v96
            case 32768:
                if (!v5) {
                    v5 = (0 >= 0)
                }
            }
        }
    default:
        for fi97 := 0; fi97 < 18; fi97++ {
            v98 := GenS8{F0: (1 == 1000), F1: v48, F2: v48}
            if (v5 && v5) {
                v99 := []int{(fi97 - fi97), (fi97 | (-(fi97))), fi97, fi97}
                if len(v99) > 0 {
                    v99[0] = 1
                }
                _ = v99
            }
            if true {
                var v100 int32 = -1
                var v101 int8 = v6
                _ = v100
                _ = v101
            }
            _ = fi97
            _ = v98
        }
    }
    switch v48 {
    case "\n\t":
        switch -1 {
        case 0:
            v4 -= 0.0
        default:
            if (!v5) {
                v5 = (42 < 85)
                v102 := GenS9{F0: v49, F1: -771, F2: 1}
                v103 := GenS8{F0: (862 <= 807), F1: "", F2: (v48 + "a")}
                _ = v102
                _ = v103
            } else {
                v6 = (v6 & v6)
                v2[("hello" + v48)] = -694
                if v5 {
                    v104 := GenS10{F0: 42, F1: v49, F2: -916}
                    if (4294967295 == 32767) {
                        fmt.Println(v48, v49)
                        var v105 uint64 = v3
                        _ = v105
                    }
                    v106 := GenS5{F0: (v6 | (v6 % (((v6 & v6)) | 1)))}
                    _ = v104
                    _ = v106
                }
            }
        }
    case "":
        if true {
            v5 = ((2 >= (v49 << ((65535) & 63))) && false)
        } else {
            v3 = 813
        }
    }
    if v5 {
        var v107 string = ("foo bar" + (v48 + (v48 + "hello")))
        for fi108 := 0; fi108 < 19; fi108++ {
            v3 = 260
            v109 := GenS5{F0: (^(v6))}
            v110 := GenS11{F0: 255}
            _ = fi108
            _ = v109
            _ = v110
        }
        _ = v107
    } else {
        var v111 int32 = -2147483648
        _ = v111
    }
    v112 := GenS5{F0: (-(v6))}
    if false {
        v6 = v6
    }
    var v113 uint16 = 2
    v114 := GenS10{F0: -1, F1: (v49 % (((v49 << (((v49 >> ((65535) & 63))) & 63))) | 1)), F2: 733}
    v115 := []int{1}
    v116 := GenS8{F0: true, F1: (v48 + "\n\t"), F2: v48}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v113:", v113)
    fmt.Println("slicelen v115:", len(v115), v115)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v112
    _ = v113
    _ = v114
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
