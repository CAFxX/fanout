package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 bool; F2 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint64; F2 int16 }

func main() {
    var v1 float32 = (3.14159 - 0.0)
    v2 := make(map[string]int)
    var v3 string = ("" + "\n\t")
    for fi4 := 0; fi4 < 6; fi4++ {
        v3 = ("hello" + v3)
        v1 = (v1 * (v1 / 1.0))
        switch "\n\t" {
        case "\n\t":
            var v5 float32 = (0.1 + ((v1 + v1) - v1))
            _ = v5
        case "":
            if true {
                var v6 uint = 318
                var v7 uint16 = 32767
                _ = v6
                _ = v7
            } else {
                v8 := GenS1{F0: 0, F1: 35, F2: 2147483647}
                _ = v8
            }
        default:
            var v9 bool = false
            _ = v9
        }
        _ = fi4
    }
    v1 = (v1 + (((v1 / (float32((1e10 - 0.1)) / ((3.14159 * v1) + float32(((0.1 * (1e10 + (0.1 / 1.0))) + 0.1))))) + v1) / ((v1 * v1) / 1.0)))
    v10 := GenS1{F0: -9223372036854775808, F1: 98, F2: -2147483648}
    delete(v2, v3)
    v11 := GenS1{F0: -9223372036854775808, F1: 131, F2: 42}
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint8 = 128
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 uint = 18446744073709551615
        _ = fi14
        _ = v15
    }
    if false {
        var v16 int64 = -1
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 string = "a"
            v3 = v3
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        var v19 int = 4294967295
        for fi20 := 0; fi20 < 18; fi20++ {
            switch (^(fi20)) {
            case 32767:
                var v21 int64 = 347
                _ = v21
            case -9223372036854775808:
                switch (v19 % ((v19) | 1)) {
                case -674:
                    for fi22 := 0; fi22 < 17; fi22++ {
                        var v23 uint64 = 18446744073709551615
                        var v24 uint = 42
                        v25 := GenS1{F0: 1, F1: 230, F2: 0}
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                case 538:
                    v26 := GenS1{F0: 42, F1: 127, F2: 0}
                    _ = v26
                default:
                    v27 := []int{(v19 + fi20), v19}
                    _ = v27
                }
            case 0:
                var v28 string = v3
                _ = v28
            }
            for fi29 := 0; fi29 < 20; fi29++ {
                switch (v3 + "hello") {
                case "":
                    if true {
                        if (!true) {
                            v1 += ((((((0.0 - v1) - (float32((0.1 + 0.1)) / 3.14159)) / (1.5 - float32((0.0 - 3.14159)))) - 3.14159) + (((1e10 + (v1 / v1)) / (float32(1.5) + (v1 / v1))) + (v1 * v1))) - v1)
                        } else {
                            if false {
                                var v30 uint64 = 4294967295
                                var v31 int16 = -539
                                _ = v30
                                _ = v31
                            }
                        }
                        if (!(((968 <= 65535) || (!true)) || true)) {
                            v3 = (v3 + ("hello" + ""))
                        }
                        v32 := []int{(fi20 & 764), v19}
                        _ = v32
                    }
                case "hello":
                    var v33 float64 = 0.0
                    v19 = (fi29 / (((fi20 * 42)) | 1))
                    _ = v33
                default:
                    if (-9223372036854775808 != -794) {
                        if false {
                            v1 += 0.0
                            v3 = v3
                        } else {
                            if true {
                                if false {
                                    v1 -= -2.25
                                } else {
                                    if ((426 >= 333) || (100 != 0)) {
                                        v19 -= v19
                                        v34 := GenS1{F0: 42, F1: 1, F2: 2147483647}
                                        _ = v34
                                    } else {
                                        var v35 uint32 = 791
                                        _ = v35
                                    }
                                    var v36 uint8 = 128
                                    if true {
                                        if ((float64(v1) + 1.5) > -2.25) {
                                            fmt.Println(fi29, v1, v19)
                                            v37 := GenS2{F0: false}
                                            _ = v37
                                        }
                                        v19 = int(v36)
                                    }
                                    _ = v36
                                }
                                v3 = v3
                            } else {
                                v38 := GenS2{F0: (1 >= 0)}
                                _ = v38
                            }
                        }
                        switch fi29 {
                        case 611:
                            v19 = -1
                        case 42:
                            v39 := []int{v19}
                            _ = v39
                        case 2147483647:
                            var v40 string = v3
                            _ = v40
                        }
                        v19 = (fi20 + (fi29 << ((42) & 63)))
                        fmt.Println(v19)
                    } else {
                        var v41 int8 = 0
                        switch (v19 & 32767) {
                        case -9223372036854775808:
                            v42 := []int{}
                            _ = v42
                        }
                        v43 := []int{v19, fi20, (v19 ^ -777), (v19 >> ((0) & 63))}
                        _ = v41
                        _ = v43
                    }
                }
                v44 := []int{2147483647, (v19 >> ((2) & 63))}
                v45 := GenS2{F0: ("" == v3)}
                _ = fi29
                _ = v44
                _ = v45
            }
            _ = fi20
        }
        fmt.Println(v1, v1)
        _ = v19
    }
    for fi46 := 0; fi46 < 17; fi46++ {
        for fi47 := 0; fi47 < 7; fi47++ {
            switch ("" + "") {
            case "hello":
                v48 := []int{(fi47 >> ((524) & 63)), 269, fi47, fi47}
                _ = v48
            case "\n\t":
                if false {
                    var v49 uint64 = 667
                    _ = v49
                } else {
                    v3 = ("" + v3)
                    v1 = (((v1 / (v1 + v1)) / (v1 - (1e10 - ((-2.25 - -2.25) - v1)))) / v1)
                }
            case "foo bar":
                v2[("" + v3)] = (fi47 / ((4294967295) | 1))
            }
            v50 := []int{fi47, fi46, fi46}
            v3 = v3
            _ = fi47
            _ = v50
        }
        _ = fi46
    }
    if true {
        v3 = ((v3 + v3) + v3)
        v2[v3] = -641
        v1 = v1
    }
    v51 := GenS3{F0: 4294967295, F1: false, F2: 827}
    if true {
        switch v3 {
        case "a":
            var v52 uint16 = 32767
            _ = v52
        case "foo bar":
            switch "foo bar" {
            case "a":
                var v53 float64 = 0.0
                _ = v53
            case "foo bar":
                if ((v3 + v3) < v3) {
                    v3 = "\n\t"
                    if true {
                        v54 := GenS1{F0: -371, F1: 0, F2: 2147483647}
                        if true {
                            for fi55 := 0; fi55 < 8; fi55++ {
                                v2[v3] = fi55
                                _ = fi55
                            }
                        } else {
                            v3 = v3
                        }
                        _ = v54
                    }
                } else {
                    if (-584 < 2147483647) {
                        if true {
                            v1 += (v1 + v1)
                        } else {
                            for fi56 := 0; fi56 < 17; fi56++ {
                                v3 = v3
                                if true {
                                    v57 := []int{-651, (-(fi56))}
                                    _ = v57
                                }
                                _ = fi56
                            }
                            if (v1 != 0.1) {
                                if (737 >= 0) {
                                    var v58 float32 = (((v1 * 0.1) + 0.1) * v1)
                                    v59 := GenS4{F0: 0}
                                    _ = v58
                                    _ = v59
                                } else {
                                    switch 858 {
                                    case 822:
                                        if ((!false) || ((1 > 42) && false)) {
                                            var v60 uint8 = 1
                                            v1 += 0.1
                                            v1 -= v1
                                            _ = v60
                                        }
                                    case 32767:
                                        v3 = ""
                                    case 0:
                                        v3 = ((((("foo bar" + "\n\t") + v3) + v3) + ("foo bar" + "")) + v3)
                                    }
                                    v61 := []int{9223372036854775807, 127, 32768, 461}
                                    switch "hello" {
                                    case "hello":
                                        v3 = (("a" + "foo bar") + ((v3 + "") + "hello"))
                                    }
                                    _ = v61
                                }
                                fmt.Println(v3)
                                fmt.Println(v3, v3)
                            } else {
                                v62 := []int{-552, 127}
                                var v63 string = v3
                                _ = v62
                                _ = v63
                            }
                        }
                    }
                }
            }
        case "":
            if true {
                fmt.Println(v1, v3)
            } else {
                var v64 uint32 = 0
                v65 := GenS3{F0: 0, F1: (42 >= 1), F2: -234}
                _ = v64
                _ = v65
            }
        }
        var v66 int32 = -477
        if true {
            v67 := GenS3{F0: 497, F1: (149 >= 18446744073709551615), F2: 1}
            v68 := []int{1, 9223372036854775807, 789}
            for fi69 := 0; fi69 < 19; fi69++ {
                fmt.Println(v1)
                var v70 int = (fi69 << ((2) & 63))
                _ = fi69
                _ = v70
            }
            _ = v67
            _ = v68
        } else {
            fmt.Println(v3, v1)
        }
        _ = v66
    } else {
        var v71 uint16 = 0
        var v72 uint64 = 884
        _ = v71
        _ = v72
    }
    if false {
        var v73 int = 4294967295
        if (42 >= 1) {
            v74 := GenS2{F0: (!true)}
            switch (v73 | (v73 * (v73 << ((475) & 63)))) {
            case 0:
                var v75 int32 = -38
                _ = v75
            case -358:
                var v76 int16 = -417
                _ = v76
            }
            _ = v74
        } else {
            for fi77 := 0; fi77 < 14; fi77++ {
                fmt.Println(fi77, v1, fi77)
                fmt.Println(v73, v1, v3)
                _ = fi77
            }
        }
        _ = v73
    } else {
        var v78 bool = false
        v1 = 1.5
        _ = v78
    }
    var v79 uint64 = 42
    if (2147483647 == 2147483647) {
        v1 = 1.5
    } else {
        switch -239 {
        case -509:
            for fi80 := 0; fi80 < 8; fi80++ {
                v81 := []int{}
                var v82 int = fi80
                switch ((v3 + v3) + v3) {
                case "\n\t":
                    fmt.Println(v1)
                case "foo bar":
                    v83 := GenS1{F0: 0, F1: 2, F2: 0}
                    _ = v83
                default:
                    for fi84 := 0; fi84 < 2; fi84++ {
                        if false {
                            if (949 == 255) {
                                v79 -= v79
                            } else {
                                var v85 bool = ((v1 * (0.1 + float32(0.0))) >= v1)
                                v1 -= 0.1
                                if (((fi80 + (v82 >> ((0) & 63))) > -86) || v85) {
                                    v1 = float32((0.1 + 0.1))
                                    if len(v81) > 0 {
                                        v81[0] = v82
                                    }
                                }
                                _ = v85
                            }
                            switch v3 {
                            case "":
                                v86 := GenS5{F0: v3, F1: (v79 / ((1) | 1)), F2: 336}
                                _ = v86
                            case "\n\t":
                                if true {
                                    v79 += 2
                                    v3 = v3
                                    var v87 int8 = 127
                                    var v88 uint = 2
                                    _ = v87
                                    _ = v88
                                }
                            default:
                                v1 -= v1
                            }
                        } else {
                            v89 := GenS4{F0: 127}
                            _ = v89
                        }
                        var v90 uint32 = 42
                        _ = fi84
                        _ = v90
                    }
                }
                _ = fi80
                _ = v81
                _ = v82
            }
        default:
            if (2 <= 4294967295) {
                if false {
                    v91 := GenS4{F0: 56}
                    _ = v91
                } else {
                    v79 = 324
                    var v92 bool = false
                    var v93 float32 = ((v1 / v1) + v1)
                    _ = v92
                    _ = v93
                }
                var v94 float64 = float64(v1)
                fmt.Println(v94, v79)
                _ = v94
            } else {
                fmt.Println(v1, v79)
                for fi95 := 0; fi95 < 17; fi95++ {
                    if ((3.14159 >= 1.5) && false) {
                        v3 = ((v3 + (v3 + ("\n\t" + v3))) + v3)
                        v1 -= 3.14159
                    } else {
                        v96 := v2[""]
                        _ = v96
                    }
                    v3 = ((v3 + v3) + v3)
                    if (41 <= 0) {
                        var v97 uint64 = v79
                        var v98 int32 = 1
                        if (2 >= 0) {
                            switch ((v3 + "a") + ((("" + (("a" + "hello") + ("hello" + (v3 + v3)))) + v3) + v3)) {
                            case "":
                                if (65535 != 2) {
                                    v1 = v1
                                    if (!((797 >= 1) || false)) {
                                        v99 := GenS2{F0: (42 <= (^(v79)))}
                                        _ = v99
                                    } else {
                                        for fi100 := 0; fi100 < 8; fi100++ {
                                            v97 = (v97 & 1)
                                            delete(v2, (v3 + v3))
                                            var v101 float32 = v1
                                            var v102 int8 = -128
                                            v101 = v101
                                            fmt.Println(v1, v97, v98)
                                            var v103 float32 = v101
                                            _ = fi100
                                            _ = v101
                                            _ = v102
                                            _ = v103
                                        }
                                        var v104 uint16 = 487
                                        if false {
                                            v104 = (v104 & v104)
                                            v105 := GenS5{F0: "foo bar", F1: (v79 + (^(v97))), F2: 514}
                                            _ = v105
                                        }
                                        _ = v104
                                    }
                                }
                            }
                            if false {
                                if (-2.25 < float64(-2.25)) {
                                    v1 -= 1.5
                                    var v106 uint16 = 626
                                    _ = v106
                                } else {
                                    switch v3 {
                                    case "hello":
                                        if false {
                                            switch 0 {
                                            case 2147483647:
                                                v107 := []int{fi95}
                                                _ = v107
                                            case 0:
                                                var v108 int64 = 202
                                                _ = v108
                                            }
                                            v1 += 0.1
                                            v109 := v2[(v3 + "")]
                                            _ = v109
                                        }
                                    case "foo bar":
                                        var v110 float32 = -2.25
                                        _ = v110
                                    default:
                                        v111 := v2[(("a" + ((v3 + "") + "a")) + "foo bar")]
                                        _ = v111
                                    }
                                }
                                v112 := GenS4{F0: -43}
                                v3 = ((("a" + "") + "hello") + (("hello" + v3) + v3))
                                _ = v112
                            }
                        }
                        _ = v97
                        _ = v98
                    }
                    _ = fi95
                }
                v1 += (1.5 / v1)
            }
        }
        var v113 bool = true
        delete(v2, v3)
        _ = v113
    }
    var v114 uint64 = (^(v79))
    var v115 uint64 = (v79 << ((18446744073709551615) & 63))
    var v116 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v116:", v116)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v10
    _ = v11
    _ = v114
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v51
    _ = v79
}
