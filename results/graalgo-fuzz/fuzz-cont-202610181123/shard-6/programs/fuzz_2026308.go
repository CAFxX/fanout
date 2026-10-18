package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int16; F2 int64 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int16; F1 bool; F2 string }
type GenS11 struct { F0 uint8; F1 uint64 }
type GenS12 struct { F0 string; F1 int16; F2 int16 }
type GenS13 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int8) uint64 {
    var v4 int32 = 0
    var v5 int8 = (p0 + 1)
    _ = p0
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 bool) uint32 {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{fi6, fi6}
        v8 := GenS1{F0: 397, F1: -32768, F2: (-1 > 3)}
        for fi9 := 0; fi9 < 9; fi9++ {
            p0 = (!(172 < 0))
            switch fi9 {
            case 32767:
                if p0 {
                    v7 = append(v7, fi6)
                    v7 = append(v7, -865)
                    if (-9223372036854775808 > -19) {
                        var v10 int16 = 1000
                        var v11 uint = 1
                        var v12 int16 = (^(v10))
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        var v13 int64 = 707
                        fmt.Println(p0)
                        v14 := GenS2{F0: fi6}
                        _ = v13
                        _ = v14
                    }
                } else {
                    if p0 {
                        var v15 int8 = -128
                        var v16 int16 = 1
                        _ = v15
                        _ = v16
                    } else {
                        var v17 int8 = -57
                        if len(v7) > 0 {
                            v7[0] = (fi9 / ((fi6) | 1))
                        }
                        switch 972 {
                        case 581:
                            fmt.Println(fi6, fi6)
                        case 4294967295:
                            var v18 uint64 = 949
                            _ = v18
                        default:
                            fmt.Println(fi6, fi9, fi9)
                        }
                        _ = v17
                    }
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        fmt.Println(fi6)
                    case 9223372036854775807:
                        var v19 uint64 = 208
                        _ = v19
                    case 128:
                        var v20 int16 = 42
                        _ = v20
                    default:
                        switch "foo bar" {
                        case "a":
                            fmt.Println(p0, fi9, p0)
                        case "hello":
                            fmt.Println(p0, fi9, p0)
                        case "":
                            fmt.Println(p0, p0, fi6)
                        default:
                            var v21 uint64 = 42
                            _ = v21
                        }
                    }
                    p0 = p0
                }
                fmt.Println(fi6, fi6, fi6)
            case 1:
                p0 = ((!(118 <= 2)) || p0)
            }
            _ = fi9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 = (840 < -314)
    var v22 uint = 4294967295
    if (!false) {
        var v23 uint = v22
        var v24 uint32 = 0
        fmt.Println(v22)
        _ = v23
        _ = v24
    } else {
        var v25 bool = p0
        v26 := GenS1{F0: 411, F1: 282, F2: (65535 <= 0)}
        fmt.Println(p0, v25, v22)
        _ = v25
        _ = v26
    }
    _ = p0
    _ = v22
    return 229
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint64 = 42
    var v3 float64 = float64((1.5 / (3.14159 + 3.14159)))
    var v27 string = (("foo bar" + "a") + "\n\t")
    if true {
        var v28 uint16 = 0
        v29 := []int{-307, 42, 146, 704}
        _ = v28
        _ = v29
    } else {
        fmt.Println(v27, v3, v3)
    }
    v3 -= (v3 + (1.5 / float64((float32((float64(-2.25) * (v3 / (0.1 / v3)))) + -2.25))))
    v1 = (v1 * 674)
    v2 = v2
    fmt.Println(v3, v3)
    switch v27 {
    case "hello":
        for fi30 := 0; fi30 < 18; fi30++ {
            v31 := GenS3{F0: v27, F1: -32768, F2: -573}
            var v32 string = ("hello" + v27)
            var v33 uint64 = 42
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
    }
    if ((!true) || (!(42 != 0))) {
        if true {
            if true {
                if true {
                    v34 := []int{-2147483648, 32767, -1}
                    switch v27 {
                    case "\n\t":
                        switch v27 {
                        case "a":
                            var v35 uint8 = 0
                            _ = v35
                        default:
                            v3 = v3
                        }
                    case "":
                        var v36 int64 = 802
                        _ = v36
                    }
                    var v37 string = v27
                    _ = v34
                    _ = v37
                }
            }
            for fi38 := 0; fi38 < 4; fi38++ {
                v3 -= -2.25
                for fi39 := 0; fi39 < 15; fi39++ {
                    v27 = ""
                    v3 -= (((v3 * (0.0 * v3)) / 1.5) - v3)
                    v2 = (^(v2))
                    switch (fi39 - -235) {
                    case 127:
                        var v40 uint32 = (v1 & 2147483647)
                        _ = v40
                    case 42:
                        v41 := GenS3{F0: v27, F1: -1, F2: 42}
                        _ = v41
                    }
                    _ = fi39
                }
                v27 = v27
                _ = fi38
            }
            var v42 uint32 = genFn1((true || false))
            _ = v42
        }
    } else {
        if (321 < 255) {
            for fi43 := 0; fi43 < 9; fi43++ {
                for fi44 := 0; fi44 < 1; fi44++ {
                    v3 -= v3
                    var v45 float64 = 0.1
                    switch ((("foo bar" + v27) + "hello") + "a") {
                    case "hello":
                        v46 := []int{}
                        _ = v46
                    default:
                        fmt.Println(v2, fi43)
                    }
                    v47 := []int{}
                    _ = fi44
                    _ = v45
                    _ = v47
                }
                if (!true) {
                    v48 := GenS1{F0: v1, F1: 798, F2: (962 < (fi43 * int(v2)))}
                    if (!false) {
                        if true {
                            v49 := []int{(fi43 << ((255) & 63)), (^(fi43))}
                            var v50 int = fi43
                            _ = v49
                            _ = v50
                        } else {
                            v51 := GenS1{F0: (v1 % ((genFn1((1 < 1))) | 1)), F1: 0, F2: (true || false)}
                            _ = v51
                        }
                        if (1 >= 26) {
                            var v52 uint8 = uint8(v1)
                            if (806 <= 2) {
                                var v53 int64 = 900
                                switch ((v27 + ("foo bar" + v27)) + v27) {
                                case "hello":
                                    var v54 float32 = 3.14159
                                    _ = v54
                                case "":
                                    var v55 float64 = v3
                                    _ = v55
                                }
                                _ = v53
                            } else {
                                for fi56 := 0; fi56 < 18; fi56++ {
                                    v52 = 0
                                    var v57 int = fi56
                                    _ = fi56
                                    _ = v57
                                }
                            }
                            _ = v52
                        }
                        v1 = genFn1(false)
                    }
                    _ = v48
                } else {
                    if ((!(true && true)) || true) {
                        if true {
                            v58 := GenS3{F0: "hello", F1: 42, F2: 1}
                            var v59 bool = false
                            var v60 uint16 = 42
                            _ = v58
                            _ = v59
                            _ = v60
                        } else {
                            if (2 <= 145) {
                                v61 := GenS4{F0: (v1 * (^(v1))), F1: (!true)}
                                switch "" {
                                case "foo bar":
                                    var v62 int8 = 100
                                    _ = v62
                                case "\n\t":
                                    v63 := GenS3{F0: (v27 + ""), F1: 1, F2: -130}
                                    _ = v63
                                case "a":
                                    v64 := GenS1{F0: 395, F1: 42, F2: false}
                                    _ = v64
                                default:
                                    if (fi43 < (fi43 << ((18446744073709551615) & 63))) {
                                        v65 := GenS5{F0: v1}
                                        v3 = 1e10
                                        switch 42 {
                                        case 32767:
                                            var v66 int64 = -1
                                            _ = v66
                                        case -271:
                                            if (((v2 ^ v2) <= (^(v2))) && true) {
                                                fmt.Println(fi43, v27)
                                                if ((-(fi43)) == fi43) {
                                                    if (198 != 1) {
                                                        switch (((v27 + (("foo bar" + v27) + v27)) + "\n\t") + "") {
                                                        case "":
                                                            for fi67 := 0; fi67 < 19; fi67++ {
                                                                v68 := GenS6{F0: v27}
                                                                v69 := GenS3{F0: v27, F1: 341, F2: -921}
                                                                switch v27 {
                                                                case "a":
                                                                    fmt.Println(fi43, v1)
                                                                default:
                                                                    v70 := []int{(fi43 * fi43), 0, -755}
                                                                    _ = v70
                                                                }
                                                                _ = fi67
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                        }
                                                        v71 := []int{}
                                                        for fi72 := 0; fi72 < 14; fi72++ {
                                                            if len(v71) > 0 {
                                                                v71[0] = 4294967295
                                                            }
                                                            if (!(-336 < 650)) {
                                                                _ = len(v71)
                                                                if true {
                                                                    v73 := GenS1{F0: (v1 << ((18446744073709551615) & 63)), F1: 1000, F2: ((3.14159 * v3) != (0.1 - 0.0))}
                                                                    v74 := GenS7{F0: (v2 / (((^(v2))) | 1))}
                                                                    _ = v73
                                                                    _ = v74
                                                                }
                                                                if false {
                                                                    var v75 uint64 = (v2 / (((v2 / (((v2 + genFn0(100))) | 1))) | 1))
                                                                    v71 = append(v71, (fi43 / (((fi72 | fi43)) | 1)))
                                                                    _ = v75
                                                                } else {
                                                                    _ = len(v71)
                                                                }
                                                            }
                                                            v76 := GenS7{F0: genFn0(0)}
                                                            v77 := GenS2{F0: fi43}
                                                            _ = fi72
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        _ = v71
                                                    }
                                                    if (!false) {
                                                        v78 := GenS8{F0: -128}
                                                        v79 := []int{fi43, fi43, (fi43 ^ (fi43 + fi43))}
                                                        fmt.Println(v3, v1)
                                                        _ = v78
                                                        _ = v79
                                                    } else {
                                                        var v80 int32 = 1000
                                                        v81 := []int{fi43, 128, 127}
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                } else {
                                                    v82 := GenS7{F0: (v2 >> ((2) & 63))}
                                                    for fi83 := 0; fi83 < 11; fi83++ {
                                                        fmt.Println(v27)
                                                        _ = fi83
                                                    }
                                                    var v84 uint64 = (v2 - (^(v2)))
                                                    _ = v82
                                                    _ = v84
                                                }
                                                if ((2 != 94) || false) {
                                                    var v85 uint8 = 49
                                                    v85 = (v85 & v85)
                                                    _ = v85
                                                }
                                            }
                                        case 9223372036854775807:
                                            for fi86 := 0; fi86 < 11; fi86++ {
                                                if ((0.0 * ((1.5 / 1.5) * -2.25)) <= ((1e10 / 0.1) - 0.0)) {
                                                    if (int64(v2) != 622) {
                                                        v87 := []int{-1, fi43, fi43}
                                                        v2 = uint64(v1)
                                                        v1 = genFn1(true)
                                                        v88 := GenS9{F0: 1}
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        v89 := GenS8{F0: 116}
                                                        _ = v89
                                                    }
                                                    v90 := GenS10{F0: 32767, F1: true, F2: "hello"}
                                                    v91 := GenS3{F0: "a", F1: 42, F2: 9223372036854775807}
                                                    if true {
                                                        v2 += genFn0(30)
                                                        var v92 uint = 65535
                                                        _ = v92
                                                    } else {
                                                        v3 = v3
                                                        var v93 string = v27
                                                        _ = v93
                                                    }
                                                    _ = v90
                                                    _ = v91
                                                } else {
                                                    v94 := GenS3{F0: v27, F1: 764, F2: 42}
                                                    switch v27 {
                                                    case "foo bar":
                                                        var v95 int16 = 32767
                                                        _ = v95
                                                    case "hello":
                                                        v1 = v1
                                                    default:
                                                        v96 := []int{}
                                                        v2 += 319
                                                        v2 = 252
                                                        _ = v96
                                                    }
                                                    v2 = (v2 >> ((442) & 63))
                                                    _ = v94
                                                }
                                                v1 += (v1 % ((4294967295) | 1))
                                                v3 = v3
                                                _ = fi86
                                            }
                                        }
                                        _ = v65
                                    } else {
                                        if (!true) {
                                            for fi97 := 0; fi97 < 5; fi97++ {
                                                v98 := GenS11{F0: 255, F1: genFn0(100)}
                                                _ = fi97
                                                _ = v98
                                            }
                                            if (18446744073709551615 > 255) {
                                                v27 = v27
                                                var v99 float32 = (1.5 + 1.5)
                                                _ = v99
                                            } else {
                                                v100 := []int{fi43, (fi43 & fi43), (-(fi43)), -2147483648}
                                                for fi101 := 0; fi101 < 20; fi101++ {
                                                    var v102 int = fi101
                                                    fmt.Println(v27, v1)
                                                    v102 -= 9223372036854775807
                                                    _ = fi101
                                                    _ = v102
                                                }
                                                v103 := GenS12{F0: v27, F1: -32768, F2: 1000}
                                                _ = v100
                                                _ = v103
                                            }
                                            fmt.Println(v3)
                                        } else {
                                            var v104 float32 = -2.25
                                            _ = v104
                                        }
                                        for fi105 := 0; fi105 < 16; fi105++ {
                                            v106 := []int{(^(fi105))}
                                            if ((v1 ^ (v1 / (((v1 | (v1 | (v1 & v1)))) | 1))) != v1) {
                                                v106 = append(v106, fi105)
                                                var v107 uint8 = 128
                                                v2 += (^(v2))
                                                _ = v107
                                            } else {
                                                fmt.Println(fi105)
                                                var v108 string = v27
                                                _ = v108
                                            }
                                            _ = fi105
                                            _ = v106
                                        }
                                    }
                                }
                                _ = v61
                            }
                        }
                        v109 := []int{fi43}
                        v27 = v27
                        _ = v109
                    } else {
                        v110 := GenS13{F0: 2, F1: 42}
                        _ = v110
                    }
                }
                _ = fi43
            }
            if true {
                v2 = genFn0(-82)
                fmt.Println(v3, v1, v27)
                v27 = v27
            }
            var v111 uint32 = 102
            _ = v111
        }
        fmt.Println(v3)
        v112 := []int{32767, -9223372036854775808, 42}
        _ = v112
    }
    for fi113 := 0; fi113 < 1; fi113++ {
        var v114 float64 = v3
        _ = fi113
        _ = v114
    }
    v27 = "\n\t"
    v27 = v27
    v27 = v27
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
}
