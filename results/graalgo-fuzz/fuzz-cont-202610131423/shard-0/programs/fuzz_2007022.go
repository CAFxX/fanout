package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32; F1 uint; F2 int16 }
type GenS9 struct { F0 int32; F1 uint64 }
type GenS10 struct { F0 uint8; F1 int8 }
type GenS11 struct { F0 bool; F1 uint16; F2 bool }
func genFn0(p0 int16, p1 uint, p2 uint) uint8 {
    for fi7 := 0; fi7 < 8; fi7++ {
        if (!false) {
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 uint = p1
                _ = fi8
                _ = v9
            }
            v10 := []int{}
            _ = v10
        } else {
            v11 := []int{2147483647}
            _ = len(v11)
            _ = v11
        }
        _ = fi7
    }
    v12 := GenS1{F0: -128, F1: 2147483648, F2: "a"}
    var v13 int8 = 127
    v14 := GenS2{F0: 954}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 477
    v3 := make(map[uint32]string)
    var v4 bool = false
    var v5 int8 = 5
    var v6 float32 = (0.1 * (3.14159 / 3.14159))
    switch 0 {
    case 42:
        var v15 uint64 = 706
        _ = v15
    case 128:
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 int32 = 908
            var v18 int = (fi16 - 871)
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v19 := GenS2{F0: 255}
    v20 := []int{-2147483648, 0, 127, 4294967295}
    for fi21 := 0; fi21 < 4; fi21++ {
        if (!(true || (0 > -480))) {
            if true {
                var v22 bool = (1 < -1)
                if len(v20) > 0 {
                    v20[0] = 9223372036854775807
                }
                _ = v22
            }
            v2 += (v2 % (((^(v2))) | 1))
            v6 += (((-2.25 - v6) - (v6 / (((1.5 * ((1.5 + v6) + 1e10)) / v6) + v6))) / v6)
        } else {
            if v4 {
                v23 := GenS2{F0: 255}
                v24 := GenS2{F0: 42}
                var v25 int32 = -2147483648
                _ = v23
                _ = v24
                _ = v25
            } else {
                v26 := GenS3{F0: (!v4), F1: -2147483648}
                switch fi21 {
                case 42:
                    switch "foo bar" {
                    case "":
                        if ((fi21 >= 4294967295) && v4) {
                            if true {
                                fmt.Println(v4, v6, v5)
                            } else {
                                fmt.Println(v6, v4)
                                v4 = false
                            }
                            for fi27 := 0; fi27 < 9; fi27++ {
                                fmt.Println(v5, fi21)
                                _ = fi27
                            }
                            v4 = true
                        }
                    case "\n\t":
                        for fi28 := 0; fi28 < 17; fi28++ {
                            v29 := GenS1{F0: v5, F1: (v2 - (^(v2))), F2: "a"}
                            _ = len(v20)
                            v30 := GenS2{F0: 32767}
                            delete(v3, (v2 * 0))
                            fmt.Println(v5, v6)
                            v31 := GenS1{F0: -85, F1: (v2 + 2), F2: ((((("\n\t" + "\n\t") + "foo bar") + "\n\t") + "") + "foo bar")}
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    default:
                        v6 -= v6
                    }
                case 0:
                    if ((v2 + (v2 | (v2 ^ v2))) == 2147483647) {
                        v2 += v2
                        fmt.Println(v6, fi21)
                        v4 = false
                    } else {
                        for fi32 := 0; fi32 < 9; fi32++ {
                            v33 := GenS3{F0: v4, F1: -578}
                            v5 -= v5
                            _ = fi32
                            _ = v33
                        }
                        if true {
                            if v4 {
                                if (42 >= 0) {
                                    v3[(v2 * (^(v2)))] = ""
                                } else {
                                    v2 -= (^(v2))
                                    if v4 {
                                        var v34 uint64 = 40
                                        v35 := GenS2{F0: 42}
                                        v34 -= v34
                                        _ = v34
                                        _ = v35
                                    }
                                    for fi36 := 0; fi36 < 13; fi36++ {
                                        if v4 {
                                            v37 := GenS3{F0: v4, F1: 2147483647}
                                            switch (fi21 % ((fi21) | 1)) {
                                            case 32768:
                                                v4 = v4
                                            case 4294967295:
                                                var v38 string = "hello"
                                                v2 = v2
                                                if v4 {
                                                    fmt.Println(fi36, v38, v2)
                                                    v4 = ((!(fi36 != (fi21 & fi36))) && (false || true))
                                                    if (1 <= 46) {
                                                        v39 := GenS4{F0: 61, F1: (v2 / ((v2) | 1))}
                                                        var v40 bool = (!v4)
                                                        v41 := v3[(^(v2))]
                                                        v42 := GenS5{F0: 1}
                                                        var v43 int8 = v5
                                                        _ = v39
                                                        _ = v40
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    } else {
                                                        var v44 int8 = (v5 + -105)
                                                        v45 := GenS3{F0: v4, F1: 1}
                                                        v46 := GenS3{F0: v4, F1: 2147483647}
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                } else {
                                                    if true {
                                                        if v4 {
                                                            if ((v5 << ((0) & 63)) <= v5) {
                                                                if true {
                                                                    var v47 int = -381
                                                                    v48 := GenS3{F0: (!v4), F1: 0}
                                                                    v49 := GenS6{F0: fi36}
                                                                    switch v38 {
                                                                    case "":
                                                                        v47 = v47
                                                                    case "\n\t":
                                                                        if v4 {
                                                                            v50 := GenS2{F0: 0}
                                                                            _ = v50
                                                                        }
                                                                        fmt.Println(v6, v47, v47)
                                                                    case "hello":
                                                                        var v51 int32 = -1
                                                                        _ = v51
                                                                    default:
                                                                        v38 = (((v38 + v38) + (("" + ("foo bar" + v38)) + "\n\t")) + v38)
                                                                    }
                                                                    _ = v47
                                                                    _ = v48
                                                                    _ = v49
                                                                } else {
                                                                    var v52 int32 = -1
                                                                    v52 = (v52 + v52)
                                                                    _ = v52
                                                                }
                                                            }
                                                            v5 -= (v5 / ((v5) | 1))
                                                            var v53 string = "foo bar"
                                                            _ = v53
                                                        }
                                                        v5 = v5
                                                    } else {
                                                        var v54 int8 = 88
                                                        _ = len(v1)
                                                        var v55 int64 = -1
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                }
                                                _ = v38
                                            case -984:
                                                v2 -= 0
                                            }
                                            var v56 float64 = float64(((0.0 / 0.1) - (v6 / v6)))
                                            _ = v37
                                            _ = v56
                                        } else {
                                            var v57 uint16 = 1
                                            v5 = (v5 - -39)
                                            _ = v57
                                        }
                                        var v58 string = ""
                                        var v59 int = fi36
                                        _ = fi36
                                        _ = v58
                                        _ = v59
                                    }
                                }
                                for fi60 := 0; fi60 < 16; fi60++ {
                                    switch (("foo bar" + ((("hello" + "\n\t") + "a") + "\n\t")) + "\n\t") {
                                    case "a":
                                        v61 := v1[1]
                                        _ = v61
                                    }
                                    v2 -= v2
                                    var v62 float64 = (float64(v6) / 3.14159)
                                    _ = fi60
                                    _ = v62
                                }
                            }
                            var v63 float64 = 0.1
                            _ = v63
                        } else {
                            for fi64 := 0; fi64 < 16; fi64++ {
                                v6 = v6
                                _ = fi64
                            }
                            if (!false) {
                                v6 += (1.5 + 0.1)
                                v6 -= (0.0 * ((v6 + (v6 / v6)) + 0.1))
                                v5 += v5
                            } else {
                                if (!v4) {
                                    switch (fi21 - (fi21 & (fi21 << ((631) & 63)))) {
                                    case -529:
                                        for fi65 := 0; fi65 < 5; fi65++ {
                                            switch -756 {
                                            case 32768:
                                                fmt.Println(fi65, v5, v2)
                                            case 815:
                                                var v66 int32 = -545
                                                _ = v66
                                            }
                                            v67 := GenS7{F0: (fi65 / ((fi65) | 1))}
                                            var v68 bool = (0 < 0)
                                            _ = fi65
                                            _ = v67
                                            _ = v68
                                        }
                                    default:
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                    }
                                }
                            }
                            v5 = (v5 - 100)
                        }
                        var v69 int64 = 9223372036854775807
                        _ = v69
                    }
                case 127:
                    if len(v20) > 0 {
                        v20[0] = (fi21 << ((536) & 63))
                    }
                }
                var v70 float64 = float64(((float32(((1.5 + -2.25) * 0.1)) / float32(((3.14159 + 0.1) + -2.25))) - v6))
                _ = v26
                _ = v70
            }
            if (1 > 1000) {
                if false {
                    v4 = true
                    for fi71 := 0; fi71 < 3; fi71++ {
                        if (v4 || (42 < 397)) {
                            v72 := GenS8{F0: 0, F1: 86, F2: 1000}
                            var v73 float32 = v6
                            v5 += (v5 % (((-(v5))) | 1))
                            _ = v72
                            _ = v73
                        } else {
                            if false {
                                v74 := GenS9{F0: 2147483647, F1: 1}
                                if v4 {
                                    var v75 int16 = -436
                                    v76 := GenS6{F0: fi71}
                                    _ = v75
                                    _ = v76
                                }
                                v4 = v4
                                _ = v74
                            }
                            v5 -= v5
                        }
                        var v77 uint = 79
                        v78 := GenS6{F0: (fi71 ^ 917)}
                        _ = fi71
                        _ = v77
                        _ = v78
                    }
                    var v79 int = fi21
                    _ = v79
                } else {
                    if ("a" < "a") {
                        var v80 uint8 = 42
                        v2 += 616
                        v4 = true
                        _ = v80
                    } else {
                        v81 := v1[v2]
                        for fi82 := 0; fi82 < 14; fi82++ {
                            var v83 int = -2147483648
                            if v4 {
                                v6 += float32((-2.25 + ((0.1 - (1e10 / (-2.25 + 0.0))) + -2.25)))
                                fmt.Println(fi82)
                            }
                            var v84 int16 = -888
                            v85 := GenS3{F0: (0.0 >= (1.5 + 0.0)), F1: 284}
                            _ = fi82
                            _ = v83
                            _ = v84
                            _ = v85
                        }
                        _ = v81
                    }
                    var v86 uint8 = 2
                    _ = v86
                }
                if len(v20) > 0 {
                    v20[0] = 0
                }
            }
        }
        fmt.Println(v5)
        _ = fi21
    }
    var v87 uint64 = 1
    fmt.Println(v87, v2)
    v2 = 2147483648
    fmt.Println(v5, v5, v6)
    fmt.Println(v6, v2, v4)
    fmt.Println(v5, v87, v4)
    v88 := GenS10{F0: 1, F1: v5}
    v87 += 42
    if v4 {
        switch "a" {
        case "foo bar":
            v87 -= (v87 + 42)
        }
    }
    var v89 bool = (!false)
    if (!(132 > genFn0(42, 18446744073709551615, 546))) {
        for fi90 := 0; fi90 < 12; fi90++ {
            fmt.Println(v2, v2, v89)
            _ = fi90
        }
        fmt.Println(v87, v87, v4)
        if (!v4) {
            var v91 uint64 = 874
            v20 = append(v20, -9223372036854775808)
            switch ("" + "\n\t") {
            case "hello":
                if false {
                    v4 = (717 >= 32767)
                    fmt.Println(v87)
                    for fi92 := 0; fi92 < 15; fi92++ {
                        for fi93 := 0; fi93 < 10; fi93++ {
                            switch fi93 {
                            case -1:
                                fmt.Println(v89, fi93)
                            case -442:
                                v2 += (v2 * v2)
                            default:
                                var v94 float64 = 0.1
                                _ = v94
                            }
                            _ = fi93
                        }
                        _ = fi92
                    }
                } else {
                    var v95 uint16 = 65535
                    for fi96 := 0; fi96 < 2; fi96++ {
                        var v97 uint16 = (v95 << ((4294967295) & 63))
                        if (false || v4) {
                            for fi98 := 0; fi98 < 9; fi98++ {
                                v91 = 2
                                fmt.Println(v87)
                                _ = fi98
                            }
                        }
                        fmt.Println(v97)
                        _ = fi96
                        _ = v97
                    }
                    if (v95 != 255) {
                        v20 = append(v20, 0)
                    }
                    _ = v95
                }
            case "a":
                v5 += v5
            case "foo bar":
                switch -2147483648 {
                case 9223372036854775807:
                    if ((float32(float64(v6)) + 3.14159) < v6) {
                        var v99 int8 = 42
                        fmt.Println(v5, v89, v99)
                        v100 := GenS11{F0: v4, F1: 424, F2: v4}
                        _ = v100
                        _ = v99
                    } else {
                        fmt.Println(v91)
                    }
                default:
                    var v101 uint32 = 445
                    _ = v101
                }
            }
            _ = v91
        } else {
            v89 = (true && (!true))
            var v102 uint32 = 2
            if v89 {
                for fi103 := 0; fi103 < 13; fi103++ {
                    var v104 uint8 = genFn0(42, 42, uint(v87))
                    v105 := GenS1{F0: v5, F1: (v102 - 4294967295), F2: "\n\t"}
                    _ = fi103
                    _ = v104
                    _ = v105
                }
                v106 := GenS3{F0: true, F1: 0}
                _ = v106
            }
            _ = v102
        }
    }
    if len(v20) > 0 {
        v20[0] = 127
    }
    v87 -= v87
    var v107 int8 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v107
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v87
    _ = v88
    _ = v89
}
