package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int64; F1 uint64; F2 bool }
type GenS4 struct { F0 int64; F1 int16; F2 uint }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint; F1 uint64; F2 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64; F1 uint; F2 bool }
type GenS9 struct { F0 string; F1 int16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64 }
func genFn0() uint64 {
    var v7 int = 42
    for fi8 := 0; fi8 < 10; fi8++ {
        if true {
            var v9 int32 = -1
            switch "\n\t" {
            case "":
                if false {
                    v9 = v9
                } else {
                    var v10 uint = 507
                    _ = v10
                }
            default:
                v7 += (v7 + (fi8 | (fi8 | -2147483648)))
            }
            var v11 string = "hello"
            _ = v11
            _ = v9
        } else {
            v7 -= (-(v7))
        }
        _ = fi8
    }
    _ = v7
    return 42
}
func genFn1(p0 uint64) uint {
    var v12 bool = (-2.25 >= 3.14159)
    var v13 string = "foo bar"
    _ = p0
    _ = v12
    _ = v13
    return 837
}

func main() {
    var v1 uint64 = 98
    var v2 uint8 = 42
    v3 := make(map[string]string)
    var v4 string = ((("\n\t" + "hello") + "foo bar") + "")
    v5 := make(map[string]uint32)
    var v6 uint16 = 1
    var v14 bool = (!(int64(v2) >= 1))
    v1 += uint64(v2)
    switch -2147483648 {
    case 128:
        var v15 uint32 = 42
        _ = v15
    case 4294967295:
        v16 := []int{}
        _ = v16
    }
    v6 = 65535
    var v17 uint16 = v6
    var v18 string = v4
    v6 -= v17
    if false {
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS1{F0: -49}
            _ = fi19
            _ = v20
        }
    } else {
        fmt.Println(v4, v4, v1)
        v18 = (v4 + "\n\t")
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        switch v18 {
        case "\n\t":
            switch fi21 {
            case -896:
                for fi22 := 0; fi22 < 7; fi22++ {
                    if v14 {
                        v14 = (v14 || v14)
                        v23 := GenS1{F0: 1}
                        var v24 int16 = 0
                        v25 := v3[v18]
                        v14 = true
                        _ = v23
                        _ = v24
                        _ = v25
                    } else {
                        fmt.Println(v4, fi21, v18)
                        fmt.Println(v2, v1, v17)
                        if v14 {
                            if (-1 <= -209) {
                                var v26 uint16 = (v17 + v17)
                                if v14 {
                                    fmt.Println(v1)
                                    var v27 float64 = ((3.14159 * 3.14159) + ((3.14159 - (0.0 + (3.14159 - 3.14159))) / 1.0))
                                    _ = len(v5)
                                    _ = v27
                                } else {
                                    fmt.Println(fi21, v1)
                                    if false {
                                        var v28 uint32 = 73
                                        var v29 uint8 = 0
                                        v30 := v3["hello"]
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    }
                                }
                                var v31 int32 = -1
                                _ = v26
                                _ = v31
                            } else {
                                if ((false || (v14 || false)) && (!((v17 >> ((0) & 63)) == (v6 | (v6 + (v6 - v6)))))) {
                                    var v32 float32 = 3.14159
                                    fmt.Println(v18)
                                    _ = v32
                                } else {
                                    var v33 uint = 641
                                    v34 := GenS2{F0: 1, F1: v14}
                                    var v35 int8 = 5
                                    if (826 > v6) {
                                        switch (fi21 << (((v33 | 255)) & 63)) {
                                        case 32768:
                                            if v14 {
                                                v36 := GenS2{F0: (v35 % ((42) | 1)), F1: v14}
                                                _ = v36
                                            }
                                        }
                                        switch ((v18 + v4) + (v4 + "foo bar")) {
                                        case "foo bar":
                                            var v37 string = v4
                                            _ = v37
                                        case "hello":
                                            v38 := GenS1{F0: (v35 | (-(v35)))}
                                            _ = v38
                                        case "a":
                                            fmt.Println(v1)
                                        }
                                    } else {
                                        switch ((v18 + v4) + "hello") {
                                        case "foo bar":
                                            v4 = (v18 + v4)
                                        case "a":
                                            v33 -= (v33 - (v33 >> (((v33 * 4294967295)) & 63)))
                                        }
                                    }
                                    v4 = ((v4 + v4) + (v4 + v4))
                                    var v39 float32 = ((1e10 * 0.1) + 0.0)
                                    v39 = (float32(0.0) - (v39 * 1e10))
                                    var v40 int = fi22
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                    _ = v39
                                    _ = v40
                                }
                                fmt.Println(v4, v17)
                            }
                            v4 = v4
                            fmt.Println(v1, fi22)
                        } else {
                            if (-1 == 1000) {
                                v41 := GenS1{F0: -1}
                                if (!v14) {
                                    var v42 uint8 = v2
                                    v17 = (v6 % (((v6 & v6)) | 1))
                                    var v43 bool = ((^(fi21)) > fi21)
                                    _ = v42
                                    _ = v43
                                } else {
                                    v18 = ("" + ((("a" + "hello") + v18) + v4))
                                }
                                _ = v41
                            } else {
                                switch ("hello" + ((v18 + v18) + (v4 + (v18 + v18)))) {
                                case "foo bar":
                                    var v44 int8 = 42
                                    _ = v44
                                case "a":
                                    if v14 {
                                        var v45 uint16 = (^(v6))
                                        v6 = v45
                                        _ = v45
                                    } else {
                                        var v46 int32 = 42
                                        v6 += v17
                                        switch (v4 + (v18 + v4)) {
                                        case "\n\t":
                                            fmt.Println(fi22, v46, v14)
                                        default:
                                            var v47 float32 = ((1e10 / 1.5) / -2.25)
                                            _ = v47
                                        }
                                        _ = v46
                                    }
                                case "":
                                    switch (v4 + "\n\t") {
                                    case "a":
                                        if false {
                                            var v48 bool = (fi22 > fi22)
                                            _ = v48
                                        } else {
                                            if true {
                                                fmt.Println(v6)
                                            } else {
                                                v49 := []int{}
                                                var v50 bool = v14
                                                _ = v49
                                                _ = v50
                                            }
                                            fmt.Println(v6, v14, v6)
                                        }
                                    case "foo bar":
                                        v5[v4] = 2147483647
                                    default:
                                        if (!v14) {
                                            if (42 == 421) {
                                                v1 = (^(v1))
                                            } else {
                                                v51 := v5[v18]
                                                var v52 int32 = 42
                                                v14 = (((1.5 - (1e10 - 1.5)) >= 0.1) || true)
                                                _ = v51
                                                _ = v52
                                            }
                                        }
                                        v18 = v4
                                    }
                                }
                            }
                            v53 := []int{-187, fi22, fi22, (fi22 * fi22)}
                            _ = v53
                        }
                    }
                    if true {
                        v54 := GenS1{F0: 100}
                        var v55 bool = v14
                        v56 := GenS3{F0: -311, F1: v1, F2: (1.5 <= (((((float32(float64((3.14159 - 1.5))) - 0.0) - float32(((((1.5 + ((3.14159 - (-2.25 + (float64(3.14159) - 1e10))) + ((-2.25 / 1e10) - 0.0))) / ((0.1 / 1.5) + -2.25)) * (0.1 - 1e10)) / 3.14159))) / (1.5 / 1.5)) * 1e10) * 3.14159))}
                        switch "hello" {
                        case "a":
                            v57 := GenS1{F0: -128}
                            _ = v57
                        case "hello":
                            var v58 bool = (v2 == uint8(v17))
                            _ = v58
                        case "foo bar":
                            var v59 float64 = 1.5
                            _ = v59
                        default:
                            fmt.Println(fi21, v18)
                        }
                        v60 := GenS4{F0: -1, F1: 592, F2: 324}
                        _ = v54
                        _ = v55
                        _ = v56
                        _ = v60
                    } else {
                        fmt.Println(v1, v14, fi22)
                    }
                    switch (v18 + (v18 + "hello")) {
                    case "\n\t":
                        v4 = "hello"
                    case "hello":
                        var v61 uint = genFn1(genFn0())
                        _ = v61
                    case "a":
                        var v62 uint = 824
                        _ = v62
                    default:
                        switch 847 {
                        case 32767:
                            switch 32767 {
                            case 127:
                                if (v4 >= v18) {
                                    var v63 bool = v14
                                    switch 4294967295 {
                                    case 0:
                                        fmt.Println(v4, fi21, v2)
                                    case 128:
                                        var v64 uint = 4294967295
                                        _ = v64
                                    case -574:
                                        fmt.Println(v14, v63, v18)
                                    default:
                                        var v65 uint = 4
                                        v66 := GenS5{F0: 363, F1: -1}
                                        fmt.Println(v18, v6, v17)
                                        v14 = (uint16(v65) > v17)
                                        v67 := []int{(fi21 >> (((v65 >> ((genFn1(v1)) & 63))) & 63)), (fi22 % (((-(fi21))) | 1))}
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    }
                                    v2 = (v2 % ((1) | 1))
                                    _ = v63
                                } else {
                                    switch 128 {
                                    case 32768:
                                        if true {
                                            fmt.Println(v2, v2)
                                            var v68 uint32 = 2
                                            _ = v68
                                        }
                                    case 127:
                                        var v69 int64 = 42
                                        _ = v69
                                    }
                                    if (!(!true)) {
                                        var v70 uint16 = 32767
                                        v5[(((((((v18 + "hello") + "\n\t") + v18) + (("\n\t" + v4) + v18)) + "hello") + v18) + ((v4 + v18) + "a"))] = 2147483647
                                        _ = v70
                                    } else {
                                        v17 = (v6 * (^(v17)))
                                    }
                                    if v14 {
                                        var v71 int32 = 1
                                        fmt.Println(v4)
                                        _ = v71
                                    }
                                    v5["foo bar"] = 4294967295
                                }
                            case 115:
                                var v72 uint64 = 2
                                _ = v72
                            case 456:
                                v3[v18] = (("a" + v4) + ((v18 + "foo bar") + v18))
                            default:
                                switch v18 {
                                case "a":
                                    if false {
                                        v1 -= (^(v1))
                                        if (!true) {
                                            if true {
                                                v73 := GenS6{F0: 255, F1: genFn0(), F2: -198}
                                                v74 := GenS6{F0: genFn1((v1 << ((65535) & 63))), F1: v1, F2: 42}
                                                _ = v73
                                                _ = v74
                                            } else {
                                                v4 = ((((v4 + v4) + (v4 + "a")) + v4) + v4)
                                                v75 := GenS1{F0: 42}
                                                fmt.Println(fi21, fi22)
                                                switch 4294967295 {
                                                case -661:
                                                    v4 = "\n\t"
                                                case -535:
                                                    var v76 float64 = float64((((1e10 * (0.0 + (float32(3.14159) * 1.5))) / 1e10) / (1.5 - 1e10)))
                                                    _ = v76
                                                default:
                                                    v77 := GenS6{F0: 2, F1: (^(v1)), F2: 1000}
                                                    _ = v77
                                                }
                                                _ = v75
                                            }
                                            v18 = "a"
                                            switch -798 {
                                            case 2147483647:
                                                if ((fi21 - 0) < fi21) {
                                                    v78 := GenS1{F0: -82}
                                                    _ = v78
                                                }
                                            case 127:
                                                switch v4 {
                                                case "foo bar":
                                                    switch v4 {
                                                    case "hello":
                                                        v79 := GenS7{F0: (v14 || (!(!v14)))}
                                                        _ = v79
                                                    default:
                                                        if v14 {
                                                            var v80 int8 = -30
                                                            v17 = (v17 - v6)
                                                            v81 := GenS6{F0: 2, F1: (^(v1)), F2: -1}
                                                            _ = v80
                                                            _ = v81
                                                        } else {
                                                            var v82 string = v18
                                                            _ = v82
                                                        }
                                                    }
                                                default:
                                                    switch v4 {
                                                    case "a":
                                                        v83 := GenS8{F0: -348, F1: 376, F2: v14}
                                                        _ = v83
                                                    }
                                                }
                                            default:
                                                var v84 int32 = 931
                                                _ = v84
                                            }
                                        } else {
                                            v14 = false
                                            var v85 bool = (42 != -128)
                                            _ = v85
                                        }
                                        v1 += (^(v1))
                                    }
                                    v86 := GenS7{F0: v14}
                                    _ = v86
                                default:
                                    v87 := GenS6{F0: 18446744073709551615, F1: 0, F2: -1}
                                    _ = v87
                                }
                            }
                        case 0:
                            var v88 uint32 = uint32(v1)
                            _ = v88
                        case 42:
                            if v14 {
                                v4 = (v18 + v4)
                            }
                            fmt.Println(v4, v4, v6)
                            fmt.Println(fi22, fi22)
                        }
                    }
                    _ = fi22
                }
            default:
                fmt.Println(fi21)
            }
        case "":
            switch fi21 {
            case -2147483648:
                v2 = v2
            case -9223372036854775808:
                v14 = v14
            }
        case "a":
            var v89 int32 = 755
            _ = v89
        default:
            _ = len(v5)
        }
        v2 += v2
        _ = fi21
    }
    switch -2147483648 {
    case 2147483647:
        var v90 int = 32768
        _ = v90
    }
    fmt.Println(v14)
    v91 := []int{-499, 32767, -169}
    for fi92 := 0; fi92 < 7; fi92++ {
        v17 = (v6 >> ((571) & 63))
        switch fi92 {
        case 32767:
            switch "\n\t" {
            case "hello":
                for fi93 := 0; fi93 < 19; fi93++ {
                    v94 := GenS1{F0: 127}
                    if (((0.1 / ((-2.25 * (0.0 + 1.5)) * 1e10)) / (0.0 + 1e10)) <= (((0.1 / 1.5) * 0.1) / float64(1.5))) {
                        if (!v14) {
                            v95 := GenS8{F0: -1, F1: 18446744073709551615, F2: (v18 <= (v4 + v18))}
                            if ((!(-2147483648 >= -1)) || ((v1 << ((65535) & 63)) > (v1 + (v1 * (v1 % (((v1 << ((65535) & 63))) | 1)))))) {
                                var v96 int64 = 9223372036854775807
                                _ = v96
                            } else {
                                v4 = v4
                            }
                            v97 := v5[v18]
                            _ = v95
                            _ = v97
                        } else {
                            var v98 bool = (((3.14159 + 0.0) + 0.0) <= ((((1e10 + float64(0.0)) + ((1.5 / 3.14159) - 3.14159)) - (float64(3.14159) * 1e10)) * 1.5))
                            v99 := GenS9{F0: (("" + "\n\t") + (v18 + v4)), F1: int16(v2)}
                            delete(v3, "")
                            _ = v98
                            _ = v99
                        }
                        switch ((v4 + v4) + "foo bar") {
                        case "":
                            if true {
                                if (!((v1 - v1) == (v1 + (v1 >> ((652) & 63))))) {
                                    var v100 uint32 = 2147483647
                                    v2 -= (v2 & v2)
                                    v101 := GenS10{F0: 42}
                                    _ = v100
                                    _ = v101
                                } else {
                                    if v14 {
                                        fmt.Println(v14, v1)
                                        var v102 float32 = (0.0 / 1.0)
                                        var v103 float32 = v102
                                        _ = v102
                                        _ = v103
                                    }
                                }
                                v91 = append(v91, (fi93 >> ((2) & 63)))
                                v6 = (v6 * (v6 ^ (^(v6))))
                                _ = len(v3)
                            } else {
                                v4 = (v18 + ("hello" + ((v18 + (v18 + v18)) + "")))
                                v5[v4] = 2147483647
                            }
                        }
                        v1 += genFn0()
                    } else {
                        var v104 bool = false
                        var v105 uint8 = (v2 | v2)
                        _ = v104
                        _ = v105
                    }
                    _ = fi93
                    _ = v94
                }
            case "foo bar":
                var v106 int16 = -792
                _ = v106
            case "\n\t":
                for fi107 := 0; fi107 < 6; fi107++ {
                    fmt.Println(v1)
                    v108 := v3["\n\t"]
                    v17 -= (^(v17))
                    switch "\n\t" {
                    case "hello":
                        v109 := GenS11{F0: (v1 & 0)}
                        _ = v109
                    case "":
                        v17 = v6
                    }
                    v2 += (v2 & (v2 >> ((4294967295) & 63)))
                    _ = fi107
                    _ = v108
                }
            }
        }
        var v110 int = (fi92 / (((^(fi92))) | 1))
        _ = fi92
        _ = v110
    }
    v111 := GenS5{F0: 66, F1: 383}
    switch v4 {
    case "foo bar":
        for fi112 := 0; fi112 < 1; fi112++ {
            if len(v91) > 0 {
                v91[0] = fi112
            }
            v17 = v6
            var v113 uint = 4294967295
            _ = fi112
            _ = v113
        }
    case "hello":
        v2 = v2
    default:
        v17 = (v17 ^ (v6 & v6))
    }
    switch -9223372036854775808 {
    case 128:
        v2 += v2
    case 4294967295:
        delete(v5, v18)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v91:", len(v91), v91)
    _ = v1
    _ = v111
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
}
