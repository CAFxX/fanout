package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 int32; F1 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 bool; F1 string }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint16 }
func genFn0(p0 uint32, p1 int64) int16 {
    p1 = (^(p1))
    v6 := []int{2147483647, -2147483648}
    if true {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint16 = 598
            _ = fi7
            _ = v8
        }
    } else {
        fmt.Println(p1)
        p0 -= 1
    }
    var v9 uint16 = 1
    _ = len(v6)
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return 1
}
func genFn1(p0 int64, p1 uint32) uint16 {
    var v10 int32 = -699
    var v11 int8 = 0
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        var v12 bool = (v10 > 0)
        _ = v12
    case "hello":
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 int32 = (v10 << ((255) & 63))
            var v15 int16 = 42
            v16 := GenS1{F0: 0}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
    }
    switch "hello" {
    case "\n\t":
        v17 := GenS1{F0: 1}
        _ = v17
    default:
        fmt.Println(p1, v11)
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 32767
}

func main() {
    var v1 bool = (true || (true || true))
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    var v4 float32 = 1.5
    v5 := []int{32768}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := GenS1{F0: 411}
        var v20 uint64 = 243
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS1{F0: 255}
        _ = fi21
        _ = v22
    }
    delete(v2, -801)
    delete(v2, 32768)
    if (1 >= 65535) {
        var v23 int32 = 42
        var v24 bool = (!v1)
        v5 = append(v5, 591)
        _ = v23
        _ = v24
    } else {
        fmt.Println(v1, v4)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println(v4, v1)
    fmt.Println(v4, v1)
    fmt.Println(v4, v4)
    v4 += float32(3.14159)
    switch ("a" + "") {
    case "foo bar":
        if true {
            v4 += (float32(3.14159) - ((v4 + (v4 + float32((-2.25 / -2.25)))) + v4))
        }
    }
    v25 := GenS2{F0: 0, F1: 0, F2: 2}
    if (v1 || (!v1)) {
        switch 2147483647 {
        case 1:
            v26 := GenS3{F0: 128, F1: ""}
            _ = v26
        case -2147483648:
            var v27 uint8 = 239
            _ = v27
        case 2147483647:
            for fi28 := 0; fi28 < 9; fi28++ {
                v4 = 0.1
                _ = fi28
            }
        default:
            for fi29 := 0; fi29 < 5; fi29++ {
                if v1 {
                    if (-1 >= fi29) {
                        switch "hello" {
                        case "\n\t":
                            var v30 int32 = 42
                            _ = v30
                        case "":
                            var v31 uint64 = 18446744073709551615
                            _ = v31
                        default:
                            v5 = append(v5, fi29)
                        }
                        if (fi29 > (fi29 >> ((2) & 63))) {
                            for fi32 := 0; fi32 < 1; fi32++ {
                                switch (fi29 << ((65535) & 63)) {
                                case -9223372036854775808:
                                    _ = len(v5)
                                default:
                                    var v33 uint8 = 128
                                    _ = v33
                                }
                                v5 = append(v5, fi29)
                                _ = fi32
                            }
                        } else {
                            switch "a" {
                            case "hello":
                                if (!v1) {
                                    for fi34 := 0; fi34 < 5; fi34++ {
                                        if ((4294967295 >= 255) || v1) {
                                            fmt.Println(fi29)
                                        } else {
                                            if v1 {
                                                if (255 <= 1) {
                                                    v5 = append(v5, (fi34 + 423))
                                                } else {
                                                    switch "hello" {
                                                    case "a":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v35 uint64 = 2
                                                            _ = v35
                                                        case "\n\t":
                                                            v4 = v4
                                                        default:
                                                            v36 := GenS4{F0: "\n\t"}
                                                            _ = v36
                                                        }
                                                    case "":
                                                        if v1 {
                                                            var v37 int32 = -2147483648
                                                            var v38 uint32 = 2147483648
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                    case "hello":
                                                        fmt.Println(v1, fi29)
                                                    }
                                                }
                                                v4 -= 3.14159
                                                v1 = (!(!(false || (!true))))
                                            } else {
                                                var v39 int64 = 0
                                                v4 = v4
                                                var v40 uint32 = 0
                                                switch fi29 {
                                                case 496:
                                                    var v41 uint64 = 0
                                                    _ = v41
                                                case -2147483648:
                                                    v42 := GenS2{F0: 0, F1: -1, F2: 706}
                                                    _ = v42
                                                case 4294967295:
                                                    var v43 int8 = 100
                                                    _ = v43
                                                }
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                        _ = fi34
                                    }
                                    if (65535 >= 180) {
                                        var v44 int8 = 63
                                        var v45 uint16 = 1
                                        if (!(295 == -2147483648)) {
                                            var v46 int16 = 32767
                                            _ = v46
                                        } else {
                                            for fi47 := 0; fi47 < 8; fi47++ {
                                                v48 := GenS2{F0: (-(v44)), F1: (v44 >> ((2) & 63)), F2: (v45 | (^(v45)))}
                                                _ = fi47
                                                _ = v48
                                            }
                                            if true {
                                                v45 = (v45 % ((2) | 1))
                                                v4 = v4
                                                v4 = (((1.5 / ((((1e10 / v4) - 0.0) * 1.5) / 1.5)) + v4) / (1.5 / (0.1 * ((-2.25 * 1.5) - float32(0.1)))))
                                            }
                                            v49 := GenS4{F0: (("" + "hello") + ((("\n\t" + ("\n\t" + (("foo bar" + "foo bar") + (("hello" + (("\n\t" + "hello") + ("" + "foo bar"))) + "foo bar")))) + "") + "hello"))}
                                            _ = v49
                                        }
                                        _ = v44
                                        _ = v45
                                    } else {
                                        var v50 int8 = 42
                                        _ = v50
                                    }
                                    fmt.Println(fi29)
                                } else {
                                    v51 := v2[fi29]
                                    var v52 int = (fi29 / ((32768) | 1))
                                    for fi53 := 0; fi53 < 13; fi53++ {
                                        if v1 {
                                            v52 = fi53
                                            if v1 {
                                                v4 -= 1e10
                                                v51 = "hello"
                                                var v54 float32 = (float32(1.5) / 1.5)
                                                _ = v54
                                            }
                                            v55 := GenS5{F0: 65535}
                                            _ = v55
                                        }
                                        _ = len(v3)
                                        switch fi29 {
                                        case 722:
                                            v56 := GenS3{F0: (fi29 ^ (v52 + 1)), F1: v51}
                                            _ = v56
                                        case 32768:
                                            switch fi53 {
                                            case 42:
                                                var v57 int = fi29
                                                _ = v57
                                            default:
                                                var v58 float32 = (v4 - ((float32((((1.5 / 1e10) / 1.5) + (0.0 * 0.0))) + float32(3.14159)) + v4))
                                                _ = v58
                                            }
                                        case 4294967295:
                                            var v59 uint8 = 221
                                            switch ("\n\t" + ("\n\t" + "foo bar")) {
                                            case "":
                                                var v60 string = "a"
                                                _ = v60
                                            case "foo bar":
                                                v61 := GenS5{F0: 2}
                                                _ = v61
                                            case "a":
                                                v62 := GenS2{F0: -1, F1: 0, F2: 1}
                                                _ = v62
                                            default:
                                                var v63 int16 = 32767
                                                _ = v63
                                            }
                                            if v1 {
                                                v51 = v51
                                                switch ("" + ("\n\t" + ("foo bar" + v51))) {
                                                case "":
                                                    var v64 int64 = 9223372036854775807
                                                    _ = v64
                                                case "hello":
                                                    if v1 {
                                                        v65 := GenS6{F0: -167, F1: 80, F2: 2}
                                                        v66 := GenS5{F0: 18446744073709551615}
                                                        if (((v4 - (0.1 / 0.1)) / ((v4 * 0.0) / v4)) > -2.25) {
                                                            switch 9223372036854775807 {
                                                            case -911:
                                                                fmt.Println(v52)
                                                            case -428:
                                                                var v67 int64 = -9223372036854775808
                                                                _ = v67
                                                            default:
                                                                if len(v5) > 0 {
                                                                    v5[0] = (v52 | -9223372036854775808)
                                                                }
                                                            }
                                                            var v68 int8 = 100
                                                            v69 := v2[fi29]
                                                            if v1 {
                                                                v68 -= int8(v52)
                                                                v51 = (v69 + (v51 + ((v69 + v51) + v51)))
                                                            } else {
                                                                fmt.Println(v1, v59, v69)
                                                            }
                                                            _ = v68
                                                            _ = v69
                                                        } else {
                                                            var v70 float32 = -2.25
                                                            _ = v70
                                                        }
                                                        _ = v65
                                                        _ = v66
                                                    } else {
                                                        var v71 uint16 = 1
                                                        _ = v71
                                                    }
                                                default:
                                                    v51 = v51
                                                }
                                            } else {
                                                var v72 bool = v1
                                                _ = v72
                                            }
                                            v2[v52] = v51
                                            v73 := GenS7{F0: 42, F1: 1}
                                            _ = v59
                                            _ = v73
                                        }
                                        _ = fi53
                                    }
                                    _ = v51
                                    _ = v52
                                }
                            default:
                                for fi74 := 0; fi74 < 19; fi74++ {
                                    switch "\n\t" {
                                    case "foo bar":
                                        v75 := GenS8{F0: 42, F1: ""}
                                        _ = v75
                                    case "\n\t":
                                        fmt.Println(fi74, v4, fi29)
                                    default:
                                        v4 = v4
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        v76 := GenS8{F0: 757, F1: "\n\t"}
                                        _ = v76
                                    case "":
                                        v1 = (v1 || (false || (false || false)))
                                    }
                                    v3[32] = fi74
                                    var v77 int32 = -2147483648
                                    var v78 int32 = 446
                                    _ = fi74
                                    _ = v77
                                    _ = v78
                                }
                            }
                            if (true && (!(210 >= 2))) {
                                v1 = v1
                            } else {
                                v1 = (v1 || ((float32(float64((((0.1 * (v4 + 1e10)) + (float32(float64(0.0)) / v4)) / (1.5 / v4)))) * 0.1) > v4))
                                v4 -= ((v4 / ((v4 * float32((1.5 / -2.25))) - v4)) / v4)
                                v79 := GenS9{F0: (fi29 >> ((4294967295) & 63))}
                                _ = v79
                            }
                        }
                    } else {
                        v4 = ((0.1 - -2.25) + ((v4 + (0.0 * (float32(((1e10 + 1e10) - -2.25)) + ((3.14159 / v4) / v4)))) * 0.1))
                        for fi80 := 0; fi80 < 9; fi80++ {
                            switch fi80 {
                            case -75:
                                switch 0 {
                                case 2147483647:
                                    switch "\n\t" {
                                    case "foo bar":
                                        if (-9223372036854775808 <= -9223372036854775808) {
                                            v81 := GenS10{F0: (42 <= -10), F1: ("" + ("a" + ("hello" + ("" + ("foo bar" + "hello")))))}
                                            var v82 float32 = v4
                                            if v1 {
                                                v82 = float32(float64(((float32(1.5) / float32(-2.25)) / v82)))
                                            }
                                            _ = v81
                                            _ = v82
                                        }
                                    case "a":
                                        var v83 int = 42
                                        _ = v83
                                    case "":
                                        v84 := GenS11{F0: 18446744073709551615}
                                        _ = v84
                                    }
                                case -9223372036854775808:
                                    var v85 int = (fi80 - fi80)
                                    _ = v85
                                default:
                                    fmt.Println(v4, v4)
                                }
                            }
                            _ = fi80
                        }
                        var v86 float32 = (((v4 / v4) + (v4 + float32((3.14159 * -2.25)))) * v4)
                        _ = v86
                    }
                    v1 = v1
                }
                _ = fi29
            }
        }
    } else {
        v1 = (v1 && true)
        switch 32767 {
        case 32767:
            v5 = append(v5, -9223372036854775808)
        default:
            if false {
                var v87 int = -2147483648
                v4 = (v4 * 0.0)
                _ = v87
            } else {
                v88 := v3[0]
                var v89 uint64 = 2
                _ = v88
                _ = v89
            }
        }
        for fi90 := 0; fi90 < 1; fi90++ {
            if len(v5) > 0 {
                v5[0] = fi90
            }
            switch "hello" {
            case "\n\t":
                delete(v3, 65535)
            case "a":
                v3[1] = -2147483648
            case "hello":
                for fi91 := 0; fi91 < 2; fi91++ {
                    v1 = (fi90 >= fi90)
                    if ((!v1) && v1) {
                        fmt.Println(fi91, v4)
                        v1 = v1
                    }
                    if (!true) {
                        var v92 float32 = (float32(1.5) * float32(1.5))
                        v93 := GenS2{F0: -1, F1: 42, F2: 1}
                        var v94 string = "\n\t"
                        _ = v92
                        _ = v93
                        _ = v94
                    }
                    _ = fi91
                }
            default:
                if true {
                    for fi95 := 0; fi95 < 1; fi95++ {
                        if (v1 || v1) {
                            if (!(((1e10 * (((0.0 * 1e10) * (0.1 * -2.25)) * (0.0 / ((3.14159 / 1e10) + 1.5)))) < (0.0 / -2.25)) && v1)) {
                                v96 := GenS12{F0: 2}
                                var v97 int32 = 0
                                v98 := GenS12{F0: 42}
                                v97 += 2147483647
                                var v99 uint64 = 1
                                var v100 float64 = 0.0
                                _ = v100
                                _ = v96
                                _ = v97
                                _ = v98
                                _ = v99
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v4 = v4
                            var v101 float64 = 1e10
                            v102 := GenS4{F0: "hello"}
                            _ = v101
                            _ = v102
                        } else {
                            v4 -= v4
                            _ = len(v5)
                            if false {
                                var v103 uint = 352
                                switch (fi90 | fi95) {
                                case -697:
                                    v4 = (3.14159 + 0.0)
                                }
                                var v104 float64 = (3.14159 * 3.14159)
                                _ = v103
                                _ = v104
                            }
                        }
                        v5 = append(v5, (fi95 - fi90))
                        _ = fi95
                    }
                } else {
                    var v105 string = ("" + "foo bar")
                    _ = v105
                }
            }
            _ = fi90
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
