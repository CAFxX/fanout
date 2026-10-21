package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 bool }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int = 42
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int8 = 40
        v10 -= -124
        v11 := []int{fi9, (fi9 & 2147483647), 875}
        _ = fi9
        _ = v10
        _ = v11
    }
    return 739
}
func genFn1() uint16 {
    var v12 uint32 = 0
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 int = (fi13 * 128)
        v14 = fi13
        _ = fi13
        _ = v14
    }
    v12 += 2147483648
    fmt.Println(v12)
    for fi15 := 0; fi15 < 3; fi15++ {
        v12 = 1
        if ((!(!(!true))) || true) {
            fmt.Println(fi15)
            v12 += 42
            var v16 uint = 616
            _ = v16
        } else {
            var v17 int8 = -128
            v17 -= v17
            v17 += 106
            _ = v17
        }
        _ = fi15
    }
    _ = v12
    return 32767
}

func main() {
    var v1 uint32 = 539
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 int32 = 42
    var v5 int16 = 42
    var v6 float64 = (1.5 * (float64(1e10) * (-2.25 - (0.1 - -2.25))))
    switch ("" + "") {
    case "\n\t":
        var v18 int32 = -2147483648
        _ = v18
    case "foo bar":
        fmt.Println(v5, v6, v4)
    default:
        v4 = -2147483648
    }
    if ((v4 <= int32(v1)) && false) {
        var v19 uint16 = 1
        if true {
            if (-128 >= 100) {
                v5 -= (v5 >> ((uint(v1)) & 63))
                v4 = 83
                for fi20 := 0; fi20 < 16; fi20++ {
                    v21 := GenS1{F0: genFn0(), F1: v5}
                    v22 := GenS1{F0: v19, F1: 32767}
                    v23 := []int{}
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            } else {
                v2[32767] = 2147483647
                if true {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        for fi25 := 0; fi25 < 6; fi25++ {
                            v4 += (-(v4))
                            var v26 uint = 2
                            if (2147483647 == (v1 << (((^(v26))) & 63))) {
                                var v27 uint16 = (v19 & (^(v19)))
                                if (true || (0.0 == (float32(v6) / 1.0))) {
                                    if (128 > 128) {
                                        var v28 int16 = v5
                                        _ = v28
                                    } else {
                                        v29 := v3[(v1 & v1)]
                                        var v30 uint = (v26 + 1)
                                        _ = v29
                                        _ = v30
                                    }
                                    _ = len(v2)
                                    v31 := GenS1{F0: (v19 << ((4294967295) & 63)), F1: v5}
                                    var v32 int64 = 1
                                    switch "a" {
                                    case "":
                                        if (672 == v32) {
                                            var v33 uint64 = 4294967295
                                            var v34 int = (fi25 + (fi25 + 1))
                                            switch 627 {
                                            case 9223372036854775807:
                                                v35 := v2[(fi24 ^ (fi24 * -387))]
                                                _ = v35
                                            case 2147483647:
                                                var v36 int64 = v32
                                                _ = v36
                                            case 42:
                                                v32 -= v32
                                            default:
                                                var v37 int = 4294967295
                                                switch 255 {
                                                case 42:
                                                    fmt.Println(v27, v6, v33)
                                                default:
                                                    switch v37 {
                                                    case -15:
                                                        v3[v1] = ("\n\t" + (("\n\t" + "foo bar") + "hello"))
                                                    default:
                                                        var v38 uint64 = (^(v33))
                                                        _ = v38
                                                    }
                                                }
                                                _ = v37
                                            }
                                            _ = v33
                                            _ = v34
                                        }
                                        v39 := GenS2{F0: (v27 / (((v19 & (v19 | 643))) | 1)), F1: (v32 * 1), F2: uint(v27)}
                                        _ = v39
                                    case "hello":
                                        switch 175 {
                                        case 1:
                                            if false {
                                                v40 := []int{}
                                                switch fi25 {
                                                case -382:
                                                    v1 = 463
                                                }
                                                _ = v40
                                            }
                                        case 0:
                                            var v41 int8 = int8(fi25)
                                            _ = v41
                                        case 32768:
                                            var v42 uint32 = (v1 % ((2147483648) | 1))
                                            v32 -= -1
                                            _ = v42
                                        }
                                    case "foo bar":
                                        v1 -= (v1 - (v1 & (v1 + 65535)))
                                    default:
                                        if true {
                                            fmt.Println(v1)
                                            v43 := GenS2{F0: v27, F1: -9223372036854775808, F2: v26}
                                            _ = v43
                                        } else {
                                            v27 = v27
                                        }
                                    }
                                    _ = v31
                                    _ = v32
                                } else {
                                    var v44 uint = v26
                                    fmt.Println(v6)
                                    _ = v44
                                }
                                _ = v27
                            }
                            v45 := GenS1{F0: uint16(v26), F1: v5}
                            _ = fi25
                            _ = v26
                            _ = v45
                        }
                        if false {
                            v3[0] = "\n\t"
                            for fi46 := 0; fi46 < 9; fi46++ {
                                v47 := GenS1{F0: (v19 ^ (v19 ^ 65535)), F1: (v5 & (v5 / ((0) | 1)))}
                                var v48 int16 = v5
                                switch ("\n\t" + ("foo bar" + "")) {
                                case "a":
                                    fmt.Println(v6)
                                case "foo bar":
                                    v1 = (v1 << ((65535) & 63))
                                }
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                        } else {
                            if ((true && false) || false) {
                                v19 += (v19 & (v19 & (^(v19))))
                            }
                        }
                        switch (fi24 ^ (fi24 * fi24)) {
                        case -2147483648:
                            v49 := []int{}
                            _ = v49
                        case 385:
                            var v50 uint8 = 128
                            _ = v50
                        }
                        _ = fi24
                    }
                    var v51 int = -2147483648
                    _ = v51
                } else {
                    var v52 float64 = 3.14159
                    _ = v52
                }
                v53 := GenS3{F0: (!true)}
                _ = v53
            }
            if true {
                v1 = 2147483647
                if true {
                    v54 := GenS3{F0: (!((!true) && (1 < uint8(v1))))}
                    _ = v54
                } else {
                    v6 = 3.14159
                    if false {
                        for fi55 := 0; fi55 < 16; fi55++ {
                            switch "\n\t" {
                            case "a":
                                var v56 bool = (259 != 1)
                                _ = v56
                            case "foo bar":
                                for fi57 := 0; fi57 < 17; fi57++ {
                                    v19 = (v19 + 32767)
                                    var v58 bool = false
                                    if (244 <= 255) {
                                        fmt.Println(v5)
                                    } else {
                                        var v59 int64 = 9223372036854775807
                                        _ = v59
                                    }
                                    fmt.Println(v58)
                                    _ = fi57
                                    _ = v58
                                }
                            case "":
                                v60 := []int{(fi55 << ((806) & 63))}
                                _ = v60
                            default:
                                var v61 bool = ((!false) && false)
                                _ = v61
                            }
                            fmt.Println(v4, fi55)
                            _ = fi55
                        }
                        if false {
                            switch "a" {
                            case "foo bar":
                                v62 := []int{}
                                _ = v62
                            case "a":
                                var v63 float64 = 3.14159
                                _ = v63
                            case "hello":
                                switch 32768 {
                                case -2147483648:
                                    v5 += (-(v5))
                                case 127:
                                    for fi64 := 0; fi64 < 18; fi64++ {
                                        var v65 int = 1
                                        var v66 float32 = (0.1 * 0.1)
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                }
                            }
                        }
                    } else {
                        var v67 uint = 42
                        _ = v67
                    }
                    for fi68 := 0; fi68 < 20; fi68++ {
                        switch (fi68 & (fi68 % ((fi68) | 1))) {
                        case 32768:
                            var v69 int16 = v5
                            _ = v69
                        }
                        var v70 int64 = -9223372036854775808
                        _ = fi68
                        _ = v70
                    }
                }
                if (!(!true)) {
                    v2[32767] = -2147483648
                }
            } else {
                v4 = -1
            }
            var v71 string = (("a" + "a") + "")
            _ = v71
        } else {
            if true {
                var v72 uint = 42
                for fi73 := 0; fi73 < 5; fi73++ {
                    v1 = (v1 >> ((v72) & 63))
                    var v74 int16 = v5
                    v75 := GenS3{F0: true}
                    _ = fi73
                    _ = v74
                    _ = v75
                }
                _ = v72
            }
            if (255 > 2) {
                fmt.Println(v5)
            }
            fmt.Println(v6, v4)
        }
        _ = v19
    } else {
        for fi76 := 0; fi76 < 16; fi76++ {
            var v77 uint = 255
            for fi78 := 0; fi78 < 11; fi78++ {
                if ((v4 / ((-918) | 1)) < v4) {
                    fmt.Println(v4, v1, fi76)
                }
                _ = fi78
            }
            _ = fi76
            _ = v77
        }
        fmt.Println(v6, v4, v1)
        _ = len(v2)
    }
    switch ("foo bar" + "") {
    case "":
        fmt.Println(v5, v5)
    case "a":
        for fi79 := 0; fi79 < 7; fi79++ {
            for fi80 := 0; fi80 < 9; fi80++ {
                v81 := GenS2{F0: 1, F1: -9223372036854775808, F2: 0}
                v5 = 1
                if true {
                    v82 := GenS4{F0: 128, F1: "foo bar"}
                    if (!(!true)) {
                        v2[128] = 1
                        var v83 float32 = 3.14159
                        if (65535 < 42) {
                            if true {
                                fmt.Println(v6)
                            }
                        } else {
                            fmt.Println(v83)
                            fmt.Println(v5, v6, v4)
                            var v84 uint16 = 2
                            _ = v84
                        }
                        _ = v83
                    }
                    _ = v82
                } else {
                    var v85 float32 = 1.5
                    if false {
                        if ((v5 | (^(v5))) <= 1) {
                            v86 := v2[2147483647]
                            switch v86 {
                            case 127:
                                fmt.Println(v1)
                            case 2147483647:
                                v87 := GenS4{F0: fi79, F1: ("hello" + "")}
                                _ = v87
                            case -1:
                                v86 -= (-(fi80))
                            }
                            delete(v3, 65535)
                            v4 = (v4 + v4)
                            _ = v86
                        } else {
                            fmt.Println(v85, v6, v4)
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v1, fi80, v85)
                                var v88 float32 = float32(((v6 + 3.14159) + 1.5))
                                var v89 int16 = v5
                                _ = v88
                                _ = v89
                            default:
                                v90 := GenS4{F0: (fi80 * fi80), F1: ""}
                                _ = v90
                            }
                        }
                        if ((^(fi80)) < 32768) {
                            var v91 float32 = (v85 * (float32(v6) / -2.25))
                            _ = v91
                        } else {
                            v6 -= (v6 * (float64(-2.25) * v6))
                            if (251 == 255) {
                                v1 = 65535
                                v92 := []int{}
                                switch -320 {
                                case 32767:
                                    if true {
                                        v93 := GenS3{F0: false}
                                        v94 := GenS5{F0: 1, F1: 223}
                                        v1 = 2147483648
                                        _ = v93
                                        _ = v94
                                    }
                                default:
                                    var v95 int = (fi79 + int(v1))
                                    _ = v95
                                }
                                _ = v92
                            }
                        }
                    }
                    delete(v3, (v1 << ((926) & 63)))
                    _ = v85
                }
                _ = fi80
                _ = v81
            }
            var v96 uint64 = 0
            _ = fi79
            _ = v96
        }
    default:
        fmt.Println(v5)
    }
    v6 += float64(-2.25)
    v4 = v4
    var v97 uint16 = 1
    var v98 float32 = (1e10 + (1e10 + (float32(0.0) + (1.5 / -2.25))))
    v98 = v98
    if ((-(v4)) >= (v4 - v4)) {
        if true {
            var v99 int64 = 1
            fmt.Println(v4, v6, v6)
            _ = v99
        }
        var v100 int16 = int16(v4)
        _ = v100
    } else {
        if true {
            if false {
                v4 += v4
                v101 := []int{246, int(v4)}
                var v102 int8 = -23
                _ = v101
                _ = v102
            }
            var v103 int = 127
            _ = v103
        }
        v97 = genFn0()
    }
    if true {
        v104 := []int{128, 42, 243, 128}
        v1 += v1
        v105 := GenS6{F0: (^(v1))}
        _ = v104
        _ = v105
    }
    v6 = v6
    v106 := GenS7{F0: 65535, F1: ((^(v1)) <= v1)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v97
    _ = v98
}
