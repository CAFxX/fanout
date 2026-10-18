package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint8; F2 uint }
type GenS4 struct { F0 int16; F1 int16; F2 uint }
type GenS5 struct { F0 int8; F1 int16 }
func genFn0(p0 int16, p1 int16, p2 int8) uint64 {
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 string = "foo bar"
        v6 := GenS1{F0: p0, F1: (fi4 % (((fi4 >> ((18446744073709551615) & 63))) | 1))}
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{fi7, (fi7 + fi7), (-(fi7)), 2147483647}
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 uint8, p1 int64, p2 int32) string {
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int32 = (-(p2))
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: -32768, F1: -455}
    var v12 int8 = 127
    if (32767 == -1) {
        v13 := GenS1{F0: -804, F1: 127}
        var v14 uint16 = 216
        v12 += (v12 | 127)
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return ("hello" + "hello")
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = -2.25
    var v3 bool = (!false)
    var v15 int16 = 42
    v3 = v3
    v3 = v3
    v16 := GenS2{F0: 522}
    var v17 uint16 = 971
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 uint = 316
        fmt.Println(v3, v17)
        _ = fi18
        _ = v19
    }
    switch "foo bar" {
    case "foo bar":
        for fi20 := 0; fi20 < 6; fi20++ {
            if (255 < 988) {
                var v21 string = ("a" + "hello")
                v22 := GenS2{F0: 977}
                _ = v21
                _ = v22
            } else {
                var v23 uint8 = 42
                if v3 {
                    if (v15 != v15) {
                        v3 = ((v23 != (v23 & 128)) || v3)
                        var v24 bool = false
                        _ = v24
                    } else {
                        fmt.Println(v15)
                    }
                }
                switch -2147483648 {
                case 694:
                    for fi25 := 0; fi25 < 8; fi25++ {
                        v3 = v3
                        v15 = 901
                        var v26 float32 = 0.1
                        _ = fi25
                        _ = v26
                    }
                }
                _ = v23
            }
            _ = fi20
        }
    case "hello":
        v17 = 32767
    case "":
        fmt.Println(v3, v17, v2)
    }
    switch -470 {
    case 1:
        var v27 int8 = -65
        _ = v27
    case 42:
        if (v3 || (!false)) {
            switch "foo bar" {
            case "a":
                v28 := []int{-1}
                _ = v28
            case "hello":
                v29 := GenS2{F0: 652}
                _ = v29
            case "foo bar":
                v30 := []int{}
                _ = v30
            default:
                switch ("\n\t" + "") {
                case "foo bar":
                    v17 += v17
                }
            }
            for fi31 := 0; fi31 < 6; fi31++ {
                if true {
                    var v32 uint64 = 2
                    fmt.Println(v17)
                    if (v17 <= v17) {
                        var v33 int = (fi31 + 127)
                        var v34 uint64 = v32
                        _ = v33
                        _ = v34
                    } else {
                        var v35 int64 = int64(v32)
                        _ = v35
                    }
                    _ = v32
                } else {
                    for fi36 := 0; fi36 < 20; fi36++ {
                        v3 = (671 >= 0)
                        _ = fi36
                    }
                    v2 += v2
                }
                if (127 < 1) {
                    var v37 int16 = -1
                    _ = v37
                }
                if (v2 > (v2 / (((1e10 / v2) - v2) / float64(float32(1e10))))) {
                    for fi38 := 0; fi38 < 8; fi38++ {
                        if v3 {
                            v2 -= v2
                            v39 := GenS2{F0: 0}
                            var v40 float64 = v2
                            _ = v39
                            _ = v40
                        } else {
                            var v41 uint64 = 345
                            var v42 uint32 = 2
                            v43 := GenS2{F0: v42}
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        v44 := []int{(fi31 >> ((4294967295) & 63)), -1, fi38}
                        _ = fi38
                        _ = v44
                    }
                } else {
                    var v45 bool = v3
                    var v46 uint64 = 42
                    var v47 uint64 = v46
                    _ = v45
                    _ = v46
                    _ = v47
                }
                _ = fi31
            }
            v2 -= (v2 + 1e10)
        } else {
            v48 := []int{}
            if ((!v3) && v3) {
                v3 = v3
            } else {
                for fi49 := 0; fi49 < 8; fi49++ {
                    for fi50 := 0; fi50 < 4; fi50++ {
                        v3 = v3
                        v15 = (v15 * -32768)
                        _ = fi50
                    }
                    switch "" {
                    case "hello":
                        if true {
                            var v51 uint16 = v17
                            _ = v51
                        }
                    case "":
                        for fi52 := 0; fi52 < 8; fi52++ {
                            v53 := GenS3{F0: "hello", F1: 1, F2: 0}
                            v3 = (!v3)
                            switch "" {
                            case "foo bar":
                                v54 := v1[0]
                                _ = v54
                            case "":
                                var v55 uint8 = 121
                                _ = len(v1)
                                _ = v55
                            case "hello":
                                if (!((!v3) && (v2 < (1.5 * float64((0.0 / (((-2.25 * (-2.25 * 1.5)) + (3.14159 + 1.5)) - 0.1))))))) {
                                    v56 := GenS1{F0: v15, F1: (fi49 - fi49)}
                                    fmt.Println(v15)
                                    delete(v1, 4294967295)
                                    _ = v56
                                } else {
                                    v57 := GenS3{F0: "\n\t", F1: 2, F2: 4294967295}
                                    if v3 {
                                        v3 = true
                                        var v58 uint16 = 42
                                        v15 = v15
                                        _ = v58
                                    } else {
                                        var v59 uint16 = 255
                                        switch fi52 {
                                        case 967:
                                            var v60 int = (fi49 % ((fi49) | 1))
                                            _ = v60
                                        case 32768:
                                            if false {
                                                fmt.Println(v3, fi49)
                                                v2 -= (v2 + 0.0)
                                            } else {
                                                v61 := GenS1{F0: int16(v59), F1: (fi49 - 9223372036854775807)}
                                                v59 += v17
                                                var v62 uint64 = 725
                                                v62 += (v62 * genFn0(42, (v15 >> ((65535) & 63)), 102))
                                                v59 = v59
                                                if len(v48) > 0 {
                                                    v48[0] = fi52
                                                }
                                                _ = v61
                                                _ = v62
                                            }
                                        case -9223372036854775808:
                                            var v63 string = "a"
                                            _ = v63
                                        }
                                        _ = v59
                                    }
                                    switch "\n\t" {
                                    case "foo bar":
                                        v2 -= v2
                                    case "a":
                                        fmt.Println(v2)
                                    case "":
                                        switch (fi49 * (fi52 << ((4294967295) & 63))) {
                                        case 32768:
                                            v15 -= v15
                                        default:
                                            switch (^(fi52)) {
                                            case -195:
                                                v3 = v3
                                            case 442:
                                                v64 := GenS2{F0: 384}
                                                _ = v64
                                            default:
                                                var v65 float64 = v2
                                                _ = v65
                                            }
                                        }
                                    }
                                    var v66 float64 = float64(float32(-2.25))
                                    fmt.Println(v17)
                                    _ = v57
                                    _ = v66
                                }
                                if (9223372036854775807 != 0) {
                                    v67 := v1[2147483648]
                                    v2 -= (0.0 / (v2 / v2))
                                    var v68 uint32 = 379
                                    _ = v67
                                    _ = v68
                                }
                                v15 = -32768
                            }
                            fmt.Println(v17, v3, fi49)
                            _ = fi52
                            _ = v53
                        }
                    case "foo bar":
                        if len(v48) > 0 {
                            v48[0] = (^(fi49))
                        }
                    default:
                        if (int16(v17) != 886) {
                            v69 := GenS1{F0: 32767, F1: -2147483648}
                            _ = v69
                        } else {
                            for fi70 := 0; fi70 < 4; fi70++ {
                                var v71 int8 = 12
                                _ = fi70
                                _ = v71
                            }
                            if true {
                                fmt.Println(v15, v2, fi49)
                                v72 := v1[2147483647]
                                _ = v72
                            }
                        }
                    }
                    if len(v48) > 0 {
                        v48[0] = (fi49 / ((fi49) | 1))
                    }
                    _ = fi49
                }
                if ((4294967295 != 178) && false) {
                    v73 := GenS2{F0: 2147483648}
                    if (true && (!(-2.25 <= 1.5))) {
                        for fi74 := 0; fi74 < 6; fi74++ {
                            var v75 float64 = float64(3.14159)
                            if (!(true || false)) {
                                var v76 int32 = 510
                                for fi77 := 0; fi77 < 11; fi77++ {
                                    _ = len(v48)
                                    switch 42 {
                                    case -1:
                                        switch "foo bar" {
                                        case "hello":
                                            var v78 uint8 = 127
                                            _ = v78
                                        case "":
                                            var v79 bool = (!((v76 ^ v76) >= -2147483648))
                                            if v3 {
                                                var v80 uint64 = 4294967295
                                                switch (fi77 / (((^(fi77))) | 1)) {
                                                case 11:
                                                    var v81 string = ""
                                                    _ = v81
                                                case -9223372036854775808:
                                                    switch ("hello" + "foo bar") {
                                                    case "":
                                                        v79 = v79
                                                    case "a":
                                                        switch "" {
                                                        case "foo bar":
                                                            v82 := GenS1{F0: 492, F1: fi74}
                                                            _ = v82
                                                        case "\n\t":
                                                            v83 := v1[2147483648]
                                                            _ = v83
                                                        default:
                                                            var v84 uint16 = v17
                                                            _ = v84
                                                        }
                                                    }
                                                case 127:
                                                    v75 = float64((((-2.25 - 0.0) / (3.14159 - 1.5)) + 1.5))
                                                default:
                                                    v48 = append(v48, (-(fi74)))
                                                }
                                                var v85 uint = 351
                                                switch "" {
                                                case "hello":
                                                    if (127 >= 213) {
                                                        if true {
                                                            var v86 uint64 = uint64(v17)
                                                            v3 = v79
                                                            var v87 uint16 = 32767
                                                            if ("hello" == "foo bar") {
                                                                v80 += (^(v86))
                                                            } else {
                                                                if v79 {
                                                                    v88 := GenS3{F0: ("a" + ("" + "a")), F1: 2, F2: (v85 * v85)}
                                                                    _ = v88
                                                                } else {
                                                                    v87 = (^(v87))
                                                                }
                                                            }
                                                            _ = v86
                                                            _ = v87
                                                        } else {
                                                            fmt.Println(v3, v3)
                                                            v79 = ((v80 << ((v85) & 63)) < (v80 % ((v80) | 1)))
                                                            fmt.Println(v3, fi74, v85)
                                                        }
                                                        v48 = append(v48, 0)
                                                    }
                                                case "":
                                                    v2 = v75
                                                }
                                                _ = v80
                                                _ = v85
                                            }
                                            v89 := GenS4{F0: (v15 >> ((4294967295) & 63)), F1: 377, F2: 4294967295}
                                            _ = v79
                                            _ = v89
                                        case "\n\t":
                                            var v90 float64 = (-2.25 + v75)
                                            _ = v90
                                        default:
                                            v91 := GenS1{F0: int16(v76), F1: int(v17)}
                                            var v92 int8 = -66
                                            var v93 uint16 = v17
                                            _ = v91
                                            _ = v92
                                            _ = v93
                                        }
                                    }
                                    _ = fi77
                                }
                                _ = v76
                            }
                            _ = fi74
                            _ = v75
                        }
                        v17 = v17
                        v94 := GenS1{F0: v15, F1: 250}
                        _ = v94
                    } else {
                        if v3 {
                            v48 = append(v48, 891)
                            var v95 bool = ((v17 + v17) >= v17)
                            v96 := GenS5{F0: 42, F1: (v15 & 42)}
                            _ = v95
                            _ = v96
                        }
                        for fi97 := 0; fi97 < 15; fi97++ {
                            fmt.Println(v2, v17, v2)
                            v98 := GenS2{F0: 2147483647}
                            for fi99 := 0; fi99 < 4; fi99++ {
                                fmt.Println(v3, v17, v2)
                                fmt.Println(v3, fi99)
                                fmt.Println(fi99, fi97)
                                fmt.Println(v3, v17)
                                var v100 uint32 = 1
                                _ = fi99
                                _ = v100
                            }
                            _ = fi97
                            _ = v98
                        }
                        fmt.Println(v2, v3)
                    }
                    v3 = (!v3)
                    _ = v73
                } else {
                    if false {
                        fmt.Println(v3, v3, v17)
                    }
                    v17 += v17
                    if len(v48) > 0 {
                        v48[0] = -1
                    }
                }
                var v101 uint8 = 235
                _ = v101
            }
            _ = v48
        }
    case 32767:
        if true {
            switch "" {
            case "a":
                v15 = v15
            case "hello":
                v102 := []int{-1, 996}
                _ = v102
            case "\n\t":
                delete(v1, 435)
            }
        }
    default:
        v3 = v3
    }
    var v103 uint16 = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v103
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
