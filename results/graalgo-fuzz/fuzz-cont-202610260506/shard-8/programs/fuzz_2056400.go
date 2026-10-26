package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint }
type GenS7 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS8 struct { F0 string; F1 uint16; F2 int16 }
func genFn0() int32 {
    v4 := GenS1{F0: 1, F1: 2147483647}
    switch ("\n\t" + "") {
    case "\n\t":
        var v5 int16 = 1000
        _ = v5
    default:
        var v6 int16 = -32768
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint8 = 2
        _ = fi7
        _ = v8
    }
    if (!false) {
        v9 := []int{127, 0, 218, 128}
        var v10 int8 = 0
        _ = v10
        _ = v9
    }
    _ = v4
    return 2147483647
}

func main() {
    var v1 int32 = -1
    var v2 int64 = -1
    var v3 uint64 = 911
    v2 = (v2 | (^(v2)))
    fmt.Println(v1)
    for fi11 := 0; fi11 < 15; fi11++ {
        v2 = int64(v1)
        if (true || true) {
            var v12 int8 = 1
            if (3.14159 >= 1.5) {
                v13 := GenS1{F0: 42, F1: (^(fi11))}
                v2 = v2
                for fi14 := 0; fi14 < 12; fi14++ {
                    v3 -= v3
                    fmt.Println(v1)
                    _ = fi14
                }
                _ = v13
            } else {
                if (583 >= 65535) {
                    var v15 uint32 = 0
                    v16 := []int{(fi11 * (^(fi11)))}
                    _ = v15
                    _ = v16
                } else {
                    var v17 int8 = (v12 ^ -33)
                    v2 -= v2
                    fmt.Println(v1, v12)
                    _ = v17
                }
            }
            _ = v12
        } else {
            if (!(1.5 >= 1e10)) {
                v18 := GenS1{F0: 42, F1: -837}
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 uint16 = 255
                    v21 := []int{32767, fi11, 2147483647, fi11}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 3; fi22++ {
                    v23 := GenS1{F0: 42, F1: (fi22 / (((fi11 >> ((2) & 63))) | 1))}
                    if (true && ((("hello" + ("foo bar" + ("" + ("\n\t" + "\n\t")))) > "foo bar") || ((^(v2)) < v2))) {
                        switch ((("foo bar" + ("" + "foo bar")) + ("\n\t" + "hello")) + (("a" + "foo bar") + "\n\t")) {
                        case "\n\t":
                            v1 += 990
                        case "hello":
                            fmt.Println(v1, fi11)
                        default:
                            v24 := GenS2{F0: v2}
                            _ = v24
                        }
                        var v25 uint16 = 182
                        v26 := GenS3{F0: (v1 - genFn0()), F1: 2}
                        var v27 int64 = (v2 << ((361) & 63))
                        v28 := []int{fi11, fi22}
                        v29 := GenS2{F0: (v2 / ((v27) | 1))}
                        if false {
                            v25 = v25
                        } else {
                            if false {
                                var v30 uint16 = v25
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                                var v31 bool = (true && (849 > uint32(v3)))
                                var v32 uint = 0
                                if (842 < 967) {
                                    var v33 uint64 = v3
                                    v34 := GenS3{F0: v1, F1: 128}
                                    if len(v28) > 0 {
                                        v28[0] = 42
                                    }
                                    _ = v33
                                    _ = v34
                                }
                                if ((127 > -25) || false) {
                                    var v35 uint = 910
                                    if v31 {
                                        v30 -= (v25 - uint16(v32))
                                        switch ("a" + "hello") {
                                        case "a":
                                            v35 = v32
                                        case "foo bar":
                                            v27 -= (v27 - (v27 & (v27 + -871)))
                                        }
                                        if (204 <= 65535) {
                                            v28 = append(v28, (fi11 ^ fi22))
                                            v32 -= v32
                                        }
                                    } else {
                                        v36 := GenS1{F0: (v3 & (v3 % ((1) | 1))), F1: fi11}
                                        v27 = (v27 / ((v27) | 1))
                                        _ = v36
                                    }
                                    fmt.Println(v31)
                                    _ = v35
                                } else {
                                    var v37 string = ""
                                    v3 = (v3 ^ v3)
                                    if (true || (((32767 == v25) && v31) && (32767 <= 0))) {
                                        v25 = 0
                                        v28 = append(v28, (fi22 & (fi22 & (^(fi22)))))
                                    }
                                    if len(v28) > 0 {
                                        v28[0] = (fi22 + (fi11 ^ (fi22 - (fi11 % ((fi22) | 1)))))
                                    }
                                    _ = v37
                                }
                                switch "hello" {
                                case "foo bar":
                                    if len(v28) > 0 {
                                        v28[0] = (fi22 % (((fi11 << ((18446744073709551615) & 63))) | 1))
                                    }
                                case "\n\t":
                                    if len(v28) >= 2 {
                                        v28 = v28[1:len(v28)]
                                    }
                                case "hello":
                                    if len(v28) >= 2 {
                                        v28 = v28[1:len(v28)]
                                    }
                                default:
                                    if (!v31) {
                                        fmt.Println(v3, v1, v27)
                                        if (889 >= v30) {
                                            var v38 uint32 = 2147483648
                                            _ = v38
                                        }
                                    }
                                }
                                _ = v30
                                _ = v31
                                _ = v32
                            } else {
                                var v39 float32 = ((((-2.25 + 0.1) + float32(-2.25)) / 3.14159) + (float32(float64(float32(float64(0.1)))) * -2.25))
                                if (((-2.25 < (1.5 + 1e10)) || (("" + "foo bar") <= "")) && (!false)) {
                                    v28 = append(v28, -655)
                                    var v40 uint = 948
                                    var v41 int32 = (v1 << ((0) & 63))
                                    _ = v40
                                    _ = v41
                                }
                                fmt.Println(fi11, fi11, fi11)
                                _ = v39
                            }
                            if false {
                                if true {
                                    _ = len(v28)
                                    var v42 uint8 = 127
                                    if false {
                                        var v43 int64 = v27
                                        if (209 != 2) {
                                            v44 := GenS2{F0: v2}
                                            v1 -= 2147483647
                                            v1 = v1
                                            _ = v44
                                        }
                                        _ = v43
                                    } else {
                                        v45 := GenS1{F0: 1, F1: fi11}
                                        _ = v45
                                    }
                                    _ = v42
                                } else {
                                    fmt.Println(v25, v27, v25)
                                    v2 = int64(v3)
                                }
                                if len(v28) > 0 {
                                    v28[0] = 0
                                }
                                var v46 int32 = v1
                                var v47 uint64 = v3
                                v48 := GenS1{F0: (v3 * 4294967295), F1: (fi22 & (fi11 ^ fi11))}
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                            var v49 uint8 = 128
                            _ = v49
                        }
                        _ = v25
                        _ = v26
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        v50 := GenS1{F0: v3, F1: 128}
                        _ = v50
                    }
                    switch 977 {
                    case 104:
                        var v51 int16 = 0
                        _ = v51
                    case -9223372036854775808:
                        v52 := GenS4{F0: v1}
                        fmt.Println(fi11)
                        var v53 int64 = 9223372036854775807
                        v54 := []int{fi22, 32767, fi11}
                        if len(v54) > 0 {
                            v54[0] = fi11
                        }
                        _ = v52
                        _ = v53
                        _ = v54
                    default:
                        switch -56 {
                        case -189:
                            v55 := GenS2{F0: (^(v2))}
                            _ = v55
                        case 2147483647:
                            var v56 uint = 400
                            _ = v56
                        }
                    }
                    _ = fi22
                    _ = v23
                }
                _ = v18
            } else {
                v57 := []int{fi11, (^(fi11)), (fi11 << ((uint(v3)) & 63)), (^(fi11))}
                _ = v57
            }
            v58 := []int{fi11, (fi11 >> ((18446744073709551615) & 63)), fi11}
            if len(v58) >= 2 {
                v58 = v58[1:len(v58)]
            }
            _ = v58
        }
        _ = fi11
    }
    v2 += -9223372036854775808
    v59 := GenS2{F0: (v2 * 1)}
    if ((1 == 0) || (267 == 375)) {
        var v60 uint32 = 4294967295
        fmt.Println(v60, v60)
        _ = v60
    } else {
        v61 := GenS5{F0: 255, F1: 510}
        _ = v61
    }
    switch "" {
    case "a":
        if (42 != 71) {
            v1 += genFn0()
            for fi62 := 0; fi62 < 12; fi62++ {
                var v63 int64 = (-(v2))
                v64 := GenS6{F0: 1, F1: v3, F2: 0}
                var v65 int16 = int16(v1)
                _ = fi62
                _ = v63
                _ = v64
                _ = v65
            }
            fmt.Println(v1, v1)
        }
    default:
        v66 := []int{1, 618}
        _ = v66
    }
    var v67 string = "a"
    for fi68 := 0; fi68 < 20; fi68++ {
        if false {
            v69 := GenS5{F0: 203, F1: (-(v2))}
            _ = v69
        }
        var v70 int16 = 32767
        _ = fi68
        _ = v70
    }
    for fi71 := 0; fi71 < 8; fi71++ {
        if (100 < 57) {
            v1 += (v1 & (^(v1)))
        } else {
            if false {
                v72 := GenS1{F0: (v3 << ((1) & 63)), F1: fi71}
                if true {
                    var v73 uint64 = (^(v3))
                    if (!(1 >= 127)) {
                        var v74 uint16 = 33
                        v75 := []int{fi71, 580}
                        for fi76 := 0; fi76 < 5; fi76++ {
                            fmt.Println(v67)
                            if (!false) {
                                v74 += 42
                                v77 := GenS2{F0: (^(v2))}
                                var v78 uint8 = 0
                                _ = v77
                                _ = v78
                            }
                            _ = fi76
                        }
                        _ = v74
                        _ = v75
                    }
                    var v79 uint64 = v3
                    _ = v73
                    _ = v79
                }
                _ = v72
            } else {
                fmt.Println(v67)
                switch v67 {
                case "":
                    v1 += 1
                case "hello":
                    v80 := []int{}
                    _ = v80
                case "foo bar":
                    if (((("" + "a") + v67) + v67) < "a") {
                        for fi81 := 0; fi81 < 14; fi81++ {
                            switch "\n\t" {
                            case "\n\t":
                                var v82 int8 = 42
                                _ = v82
                            case "foo bar":
                                fmt.Println(fi81)
                            default:
                                v83 := []int{fi81, fi81, fi71}
                                _ = v83
                            }
                            if true {
                                var v84 uint32 = 42
                                if false {
                                    fmt.Println(v84, v3, v1)
                                } else {
                                    v2 += (v2 >> ((206) & 63))
                                }
                                var v85 uint32 = (v84 << ((18446744073709551615) & 63))
                                _ = v84
                                _ = v85
                            }
                            _ = fi81
                        }
                        fmt.Println(fi71)
                    }
                }
                var v86 uint16 = 1
                _ = v86
            }
            v87 := GenS2{F0: v2}
            v88 := GenS7{F0: 65535, F1: (v3 >> ((255) & 63)), F2: 1}
            _ = v87
            _ = v88
        }
        v89 := []int{(fi71 & (fi71 | fi71))}
        if (!(4294967295 <= 2147483647)) {
            for fi90 := 0; fi90 < 14; fi90++ {
                var v91 int16 = 1000
                _ = fi90
                _ = v91
            }
            var v92 uint16 = 255
            _ = v92
        }
        _ = fi71
        _ = v89
    }
    for fi93 := 0; fi93 < 8; fi93++ {
        var v94 uint64 = (v3 >> ((667) & 63))
        v95 := GenS5{F0: 99, F1: (v2 - -9223372036854775808)}
        _ = fi93
        _ = v94
        _ = v95
    }
    var v96 int64 = v2
    v97 := GenS8{F0: ("a" + v67), F1: 1, F2: 1}
    var v98 int32 = genFn0()
    v1 = v1
    var v99 int8 = 42
    fmt.Println(v1, v98, v96)
    var v100 int32 = (v1 + (v98 & (v1 | (v1 % ((0) | 1)))))
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    fmt.Println("v96:", v96)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v59
    _ = v67
    _ = v96
    _ = v97
    _ = v98
    _ = v99
}
