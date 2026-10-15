package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint; F2 string }
type GenS5 struct { F0 int8; F1 int8; F2 uint16 }
type GenS6 struct { F0 bool; F1 int; F2 uint16 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 uint32; F1 bool; F2 int }
type GenS9 struct { F0 bool; F1 int32; F2 uint16 }
type GenS10 struct { F0 int8; F1 int16; F2 uint64 }
type GenS11 struct { F0 uint8; F1 int64; F2 int }
func genFn0(p0 string, p1 uint16, p2 uint32) uint16 {
    fmt.Println(p2, p1)
    var v4 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p1 ^ p1)
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (-2.25 - 0.1)
    var v3 string = "foo bar"
    var v5 uint8 = 128
    var v6 int = 0
    v6 = 32767
    delete(v1, v3)
    v7 := GenS1{F0: "a"}
    v8 := GenS1{F0: (("a" + "foo bar") + "foo bar")}
    if ((-1 < 9223372036854775807) || false) {
        if false {
            v6 = v6
            v3 = "foo bar"
        }
    } else {
        if true {
            if true {
                if (true || (!true)) {
                    switch v6 {
                    case 42:
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v6, v2, v3)
                        default:
                            if (!(1 <= 1000)) {
                                v3 = (((v3 + "foo bar") + v3) + (v3 + v3))
                                for fi9 := 0; fi9 < 4; fi9++ {
                                    if true {
                                        fmt.Println(v2)
                                    } else {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v3 = ""
                                        case "foo bar":
                                            v6 = fi9
                                        case "a":
                                            for fi10 := 0; fi10 < 5; fi10++ {
                                                v11 := GenS1{F0: ""}
                                                if (465 <= fi10) {
                                                    var v12 int8 = -60
                                                    _ = v12
                                                } else {
                                                    delete(v1, "\n\t")
                                                    var v13 int64 = 0
                                                    v3 = "\n\t"
                                                    if false {
                                                        v14 := GenS2{F0: 42, F1: 65535, F2: 4294967295}
                                                        if false {
                                                            v15 := GenS3{F0: 1, F1: 419}
                                                            v13 = v13
                                                            _ = v15
                                                        } else {
                                                            var v16 int = (fi10 | (v6 / ((fi9) | 1)))
                                                            if (v5 > (^(v5))) {
                                                                v17 := GenS1{F0: ((("a" + "hello") + v3) + "hello")}
                                                                if (false || (1.5 > 3.14159)) {
                                                                    var v18 string = "foo bar"
                                                                    _ = v18
                                                                } else {
                                                                    switch ("foo bar" + "a") {
                                                                    case "a":
                                                                        v13 -= v13
                                                                    case "foo bar":
                                                                        var v19 int16 = 933
                                                                        _ = v19
                                                                    }
                                                                    var v20 uint64 = uint64(v5)
                                                                    _ = v20
                                                                }
                                                                var v21 uint32 = 853
                                                                var v22 uint8 = (v5 * v5)
                                                                _ = v17
                                                                _ = v21
                                                                _ = v22
                                                            }
                                                            if true {
                                                                if true {
                                                                    var v23 uint16 = 42
                                                                    v2 = (0.0 * ((v2 + 0.0) + (float64(0.1) + v2)))
                                                                    var v24 int16 = 346
                                                                    _ = v23
                                                                    _ = v24
                                                                } else {
                                                                    if (!(!((v3 > "foo bar") || true))) {
                                                                        v25 := GenS1{F0: ((v3 + "") + v3)}
                                                                        v26 := GenS2{F0: 2147483647, F1: 865, F2: 981}
                                                                        _ = v25
                                                                        _ = v26
                                                                    } else {
                                                                        v16 = v16
                                                                    }
                                                                    if true {
                                                                        v13 = v13
                                                                        v2 -= v2
                                                                        if false {
                                                                            v1[((("foo bar" + v3) + v3) + (v3 + v3))] = v6
                                                                        }
                                                                    }
                                                                    if (!true) {
                                                                        v1[((v3 + v3) + "hello")] = (fi10 - (fi9 % ((-531) | 1)))
                                                                    } else {
                                                                        var v27 bool = (!false)
                                                                        _ = v27
                                                                    }
                                                                }
                                                            }
                                                            v3 = (v3 + (v3 + ((v3 + v3) + v3)))
                                                            _ = v16
                                                        }
                                                        _ = v14
                                                    } else {
                                                        v28 := GenS2{F0: 0, F1: 42, F2: 42}
                                                        _ = v28
                                                    }
                                                    v5 = 0
                                                    _ = v13
                                                }
                                                _ = fi10
                                                _ = v11
                                            }
                                        }
                                        v29 := GenS4{F0: 1, F1: 255, F2: v3}
                                        _ = v29
                                    }
                                    _ = fi9
                                }
                            }
                        }
                    case 486:
                        v30 := GenS4{F0: 226, F1: 255, F2: v3}
                        _ = v30
                    default:
                        var v31 uint64 = 408
                        _ = v31
                    }
                    if false {
                        if false {
                            v3 = "hello"
                            v2 += 1e10
                        } else {
                            fmt.Println(v2)
                            v5 += (v5 ^ (v5 - (v5 * 0)))
                            switch "a" {
                            case "":
                                v2 -= (0.0 * float64((1e10 + (0.1 - 0.1))))
                            case "foo bar":
                                v3 = v3
                            case "\n\t":
                                if false {
                                    for fi32 := 0; fi32 < 14; fi32++ {
                                        for fi33 := 0; fi33 < 2; fi33++ {
                                            _ = len(v1)
                                            _ = fi33
                                        }
                                        var v34 int16 = 1
                                        _ = fi32
                                        _ = v34
                                    }
                                    v35 := []int{}
                                    if false {
                                        _ = len(v1)
                                        if (-1 >= -1) {
                                            delete(v1, v3)
                                        }
                                        if (122 < 2) {
                                            var v36 int = v6
                                            if (2147483647 == uint32(v5)) {
                                                var v37 uint64 = 42
                                                if ((v37 | (v37 & (^(v37)))) < v37) {
                                                    switch (v3 + v3) {
                                                    case "a":
                                                        v38 := GenS2{F0: 0, F1: 901, F2: 266}
                                                        _ = v38
                                                    case "\n\t":
                                                        if false {
                                                            for fi39 := 0; fi39 < 14; fi39++ {
                                                                delete(v1, v3)
                                                                _ = fi39
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v40 int16 = -959
                                                        _ = v40
                                                    }
                                                }
                                                _ = v37
                                            }
                                            _ = v36
                                        } else {
                                            _ = len(v1)
                                            v41 := GenS3{F0: 51, F1: 1}
                                            for fi42 := 0; fi42 < 4; fi42++ {
                                                fmt.Println(fi42)
                                                v6 += fi42
                                                fmt.Println(v5, v5)
                                                _ = fi42
                                            }
                                            _ = v41
                                        }
                                    }
                                    _ = v35
                                } else {
                                    v43 := GenS2{F0: -1, F1: 42, F2: 57}
                                    v44 := GenS2{F0: -347, F1: 1, F2: 2147483647}
                                    _ = v43
                                    _ = v44
                                }
                            default:
                                if (false && false) {
                                    v45 := GenS1{F0: ("\n\t" + ((v3 + "hello") + v3))}
                                    var v46 float32 = -2.25
                                    for fi47 := 0; fi47 < 18; fi47++ {
                                        var v48 int32 = -2147483648
                                        v49 := GenS1{F0: v3}
                                        if (799 != 336) {
                                            var v50 uint32 = 65535
                                            _ = v50
                                        }
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                    }
                                    _ = v45
                                    _ = v46
                                } else {
                                    fmt.Println(v3)
                                }
                            }
                        }
                        var v51 string = v3
                        v52 := []int{}
                        _ = v51
                        _ = v52
                    } else {
                        v3 = v3
                    }
                } else {
                    v6 = v6
                    v53 := GenS5{F0: -74, F1: 42, F2: genFn0(v3, 1, 0)}
                    v5 = v5
                    _ = v53
                }
                v5 = (v5 - (^(v5)))
                switch (v6 % ((v6) | 1)) {
                case -1:
                    var v54 int64 = 0
                    _ = v54
                case -774:
                    for fi55 := 0; fi55 < 7; fi55++ {
                        v56 := []int{}
                        if true {
                            var v57 int16 = 1000
                            var v58 int = 32767
                            _ = v57
                            _ = v58
                        }
                        if false {
                            fmt.Println(v2)
                            v59 := GenS3{F0: 1, F1: 510}
                            if (1 >= 0) {
                                var v60 int64 = 311
                                fmt.Println(fi55, fi55, v60)
                                v61 := GenS6{F0: ((v60 << ((80) & 63)) != v60), F1: v6, F2: 886}
                                _ = v60
                                _ = v61
                            }
                            _ = v59
                        }
                        _ = fi55
                        _ = v56
                    }
                case 2147483647:
                    _ = len(v1)
                }
            }
            if false {
                var v62 float32 = 1e10
                for fi63 := 0; fi63 < 12; fi63++ {
                    for fi64 := 0; fi64 < 14; fi64++ {
                        var v65 float64 = 1e10
                        _ = fi64
                        _ = v65
                    }
                    if true {
                        for fi66 := 0; fi66 < 14; fi66++ {
                            var v67 int8 = 0
                            fmt.Println(v3, v62)
                            v68 := GenS7{F0: v5, F1: (false || true)}
                            _ = fi66
                            _ = v67
                            _ = v68
                        }
                        var v69 int = fi63
                        for fi70 := 0; fi70 < 10; fi70++ {
                            v2 -= v2
                            _ = fi70
                        }
                        _ = v69
                    }
                    fmt.Println(v6)
                    _ = fi63
                }
                _ = v62
            }
        } else {
            v2 = (v2 - ((1e10 * v2) - 1e10))
            v5 -= (v5 - (v5 | v5))
            for fi71 := 0; fi71 < 6; fi71++ {
                var v72 int16 = -10
                _ = fi71
                _ = v72
            }
        }
        for fi73 := 0; fi73 < 19; fi73++ {
            v2 -= v2
            for fi74 := 0; fi74 < 5; fi74++ {
                var v75 uint8 = v5
                _ = fi74
                _ = v75
            }
            _ = fi73
        }
        switch (((v3 + v3) + ("" + v3)) + v3) {
        case "a":
            var v76 uint = 960
            _ = v76
        }
    }
    v6 -= (^(v6))
    if (0 == 42) {
        var v77 uint8 = 1
        for fi78 := 0; fi78 < 9; fi78++ {
            v79 := GenS8{F0: 0, F1: false, F2: 32767}
            if (v5 != (v5 - 127)) {
                fmt.Println(v6)
            } else {
                switch v3 {
                case "":
                    for fi80 := 0; fi80 < 12; fi80++ {
                        var v81 uint64 = 680
                        var v82 int8 = -114
                        var v83 uint32 = 42
                        _ = fi80
                        _ = v81
                        _ = v82
                        _ = v83
                    }
                case "a":
                    v3 = (v3 + v3)
                default:
                    v3 = "hello"
                }
                v84 := GenS9{F0: (true && true), F1: 926, F2: genFn0(v3, 2, 1)}
                _ = v84
            }
            for fi85 := 0; fi85 < 20; fi85++ {
                v6 = (fi78 ^ (fi78 / ((32767) | 1)))
                _ = fi85
            }
            _ = fi78
            _ = v79
        }
        _ = v77
    } else {
        if false {
            for fi86 := 0; fi86 < 20; fi86++ {
                if false {
                    var v87 float32 = -2.25
                    _ = v87
                } else {
                    var v88 float64 = v2
                    v89 := GenS10{F0: -128, F1: 32767, F2: 722}
                    _ = v88
                    _ = v89
                }
                _ = fi86
            }
        } else {
            if true {
                var v90 int16 = -90
                v3 = ""
                _ = v90
            } else {
                for fi91 := 0; fi91 < 18; fi91++ {
                    var v92 float64 = v2
                    _ = fi91
                    _ = v92
                }
                v93 := GenS1{F0: "hello"}
                v6 = 9223372036854775807
                _ = v93
            }
        }
        if (true || (695 <= 668)) {
            v94 := []int{}
            for fi95 := 0; fi95 < 19; fi95++ {
                var v96 uint16 = 32767
                var v97 uint16 = v96
                v96 = genFn0(v3, (v96 >> ((1) & 63)), 42)
                _ = fi95
                _ = v96
                _ = v97
            }
            switch v6 {
            case 9223372036854775807:
                var v98 float64 = ((((-2.25 * v2) * v2) / v2) - float64(0.0))
                _ = v98
            case -2147483648:
                v5 = v5
            }
            _ = v94
        }
    }
    fmt.Println(v6, v6)
    for fi99 := 0; fi99 < 6; fi99++ {
        if false {
            v100 := []int{}
            for fi101 := 0; fi101 < 14; fi101++ {
                var v102 uint8 = (v5 / (((^(v5))) | 1))
                v102 = (v102 % ((v5) | 1))
                _ = fi101
                _ = v102
            }
            if false {
                for fi103 := 0; fi103 < 14; fi103++ {
                    fmt.Println(v2, v5, fi99)
                    v2 -= -2.25
                    _ = fi103
                }
            } else {
                var v104 float32 = float32(3.14159)
                var v105 float32 = ((v104 / 3.14159) - 1e10)
                _ = v104
                _ = v105
            }
            _ = v100
        } else {
            for fi106 := 0; fi106 < 14; fi106++ {
                v2 += (v2 - v2)
                v107 := GenS9{F0: true, F1: 238, F2: 1}
                _ = fi106
                _ = v107
            }
            v6 -= v6
            v108 := GenS11{F0: (v5 ^ (v5 & v5)), F1: 835, F2: v6}
            _ = v108
        }
        if (1 >= 42) {
            if true {
                if ((-821 < 0) && true) {
                    v6 -= 1
                    var v109 uint64 = 2
                    _ = v109
                }
                v3 = (v3 + v3)
            }
        }
        var v110 bool = (!(!(-1 <= 2)))
        _ = fi99
        _ = v110
    }
    fmt.Println(v3)
    v3 = v3
    for fi111 := 0; fi111 < 12; fi111++ {
        var v112 int16 = -933
        _ = fi111
        _ = v112
    }
    delete(v1, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
