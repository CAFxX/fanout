package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint16; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int32; F1 uint32; F2 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 uint8 }
func genFn0(p0 uint32) int8 {
    p0 -= (p0 << ((1) & 63))
    v5 := []int{-384, -477, 9223372036854775807, 32767}
    var v6 int = 32767
    v7 := GenS1{F0: ("hello" + "foo bar")}
    var v8 int32 = -1
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 100
}
func genFn1(p0 uint32) bool {
    v9 := GenS1{F0: "a"}
    p0 += (p0 + (^(p0)))
    if true {
        var v10 float32 = 0.0
        _ = v10
    }
    v11 := []int{}
    _ = p0
    _ = v11
    _ = v9
    return true
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]uint32)
    var v3 int16 = 627
    var v4 string = "foo bar"
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: v4}
        var v14 uint32 = v1
        if false {
            delete(v2, (fi12 & (fi12 / ((-2147483648) | 1))))
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint32 = 65535
    v16 := v2[-1]
    v17 := GenS1{F0: (v4 + (v4 + v4))}
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 13; fi19++ {
            if false {
                switch fi19 {
                case -2147483648:
                    v20 := GenS1{F0: v4}
                    _ = v20
                case -9223372036854775808:
                    var v21 int = (fi19 & fi18)
                    _ = v21
                case 2147483647:
                    switch fi19 {
                    case 4294967295:
                        if true {
                            v16 -= (v15 % (((v16 >> ((134) & 63))) | 1))
                            var v22 uint32 = v1
                            v23 := v2[-339]
                            _ = v22
                            _ = v23
                        }
                    }
                }
                fmt.Println(fi19, v4)
            } else {
                if (!(1 >= 255)) {
                    if (!(false || true)) {
                        v15 = (^(v16))
                        var v24 int8 = 52
                        switch fi19 {
                        case 127:
                            var v25 uint8 = 1
                            _ = v25
                        case -2147483648:
                            v26 := GenS1{F0: v4}
                            _ = v26
                        }
                        if false {
                            v16 = (^(v16))
                            if true {
                                v15 = v1
                            }
                        } else {
                            var v27 int32 = 2147483647
                            delete(v2, fi18)
                            _ = v27
                        }
                        _ = v24
                    }
                    v1 = 65535
                }
            }
            v28 := GenS1{F0: v4}
            _ = fi19
            _ = v28
        }
        for fi29 := 0; fi29 < 14; fi29++ {
            v4 = (v4 + "hello")
            _ = fi29
        }
        for fi30 := 0; fi30 < 18; fi30++ {
            if false {
                if false {
                    v31 := GenS1{F0: v4}
                    fmt.Println(fi30, v16)
                    var v32 int32 = 1
                    _ = v31
                    _ = v32
                } else {
                    v16 = v15
                    if (true || (v3 < (v3 << ((188) & 63)))) {
                        if false {
                            switch "" {
                            case "\n\t":
                                if ((false || true) && true) {
                                    if ((v4 + "a") > ("foo bar" + ("\n\t" + v4))) {
                                        fmt.Println(v16, v3, fi30)
                                        v33 := GenS1{F0: v4}
                                        var v34 int8 = -128
                                        v35 := []int{(fi18 << ((1) & 63))}
                                        switch fi18 {
                                        case -371:
                                            if (true || (false || false)) {
                                                v36 := GenS1{F0: (v4 + "\n\t")}
                                                v16 = v1
                                                if (((1e10 - (((3.14159 - -2.25) * 3.14159) + (1.5 / (0.1 + ((0.0 * (1e10 - 0.1)) / 1.5))))) * (0.1 + (float32(((3.14159 / 1.0) * 1e10)) - (1.5 + (float32(1.5) * 0.1))))) < 3.14159) {
                                                    v1 -= (v16 + (v15 * 346))
                                                } else {
                                                    if true {
                                                        v37 := GenS1{F0: (v4 + "\n\t")}
                                                        _ = v37
                                                    }
                                                }
                                                _ = v36
                                            } else {
                                                var v38 int8 = genFn0(v1)
                                                _ = len(v2)
                                                _ = v38
                                            }
                                        case 32768:
                                            fmt.Println(v1, fi18, v3)
                                        case 32767:
                                            v34 -= genFn0((v1 - 775))
                                        default:
                                            fmt.Println(v4)
                                        }
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                } else {
                                    var v39 float32 = 0.1
                                    v40 := []int{fi18}
                                    var v41 float32 = v39
                                    fmt.Println(fi30, v4, fi18)
                                    v42 := GenS1{F0: ("foo bar" + ("" + v4))}
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v43 uint64 = 951
                                        _ = v43
                                    default:
                                        var v44 int64 = 463
                                        _ = v44
                                    }
                                    v1 = (v15 ^ v1)
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                            case "hello":
                                var v45 int64 = 9223372036854775807
                                _ = v45
                            case "foo bar":
                                var v46 uint32 = (v1 | (v1 / ((v16) | 1)))
                                _ = v46
                            }
                            v47 := v2[fi30]
                            _ = v47
                        }
                        fmt.Println(fi30, v4)
                        v16 -= 2147483647
                    } else {
                        v3 = (v3 >> ((255) & 63))
                    }
                }
                v1 += v16
            }
            v15 += (v15 + v1)
            _ = fi30
        }
        _ = fi18
    }
    for fi48 := 0; fi48 < 5; fi48++ {
        for fi49 := 0; fi49 < 7; fi49++ {
            var v50 int = (^(fi49))
            if true {
                v16 = 335
                switch fi48 {
                case -446:
                    if (("\n\t" + v4) <= v4) {
                        var v51 uint16 = 65535
                        _ = v51
                    }
                }
                _ = len(v2)
                var v52 int64 = -9223372036854775808
                _ = v52
            }
            switch 42 {
            case 584:
                v50 += (fi48 << ((18446744073709551615) & 63))
                v53 := v2[(v50 % ((-9223372036854775808) | 1))]
                v1 += v53
                var v54 int64 = -9223372036854775808
                _ = v53
                _ = v54
            case 42:
                v55 := GenS2{F0: 257, F1: 42, F2: -246}
                _ = v55
            case -9223372036854775808:
                if true {
                    if (-2147483648 == 42) {
                        v2[(v50 / ((fi48) | 1))] = v16
                        v56 := []int{}
                        _ = v56
                    }
                    switch ((((v4 + "foo bar") + ("hello" + (("" + (v4 + v4)) + v4))) + "hello") + v4) {
                    case "":
                        v15 -= v1
                    case "\n\t":
                        v3 = (-(v3))
                    case "a":
                        v57 := v2[-9223372036854775808]
                        _ = v57
                    }
                    if (!((v50 & fi48) > v50)) {
                        switch fi48 {
                        case -207:
                            var v58 uint16 = 262
                            _ = v58
                        case 32768:
                            var v59 float32 = -2.25
                            _ = v59
                        case 127:
                            v60 := GenS3{F0: 0}
                            _ = v60
                        default:
                            v1 += (^(v1))
                        }
                        fmt.Println(fi48, v3)
                        v16 = (^(v15))
                        v61 := []int{(fi49 + v50), (v50 | (fi48 % ((fi49) | 1))), fi49}
                        _ = v61
                    }
                } else {
                    v1 = 0
                }
            }
            var v62 int32 = 781
            _ = fi49
            _ = v50
            _ = v62
        }
        switch ((("a" + v4) + "\n\t") + "hello") {
        case "a":
            v1 += 65535
        case "\n\t":
            v63 := []int{}
            _ = v63
        case "foo bar":
            for fi64 := 0; fi64 < 7; fi64++ {
                v65 := GenS4{F0: 897, F1: fi48}
                _ = fi64
                _ = v65
            }
        }
        if false {
            v66 := GenS5{F0: -674, F1: 2147483647, F2: 127}
            switch fi48 {
            case 128:
                _ = len(v2)
            case -175:
                v67 := []int{(fi48 & (fi48 ^ fi48))}
                _ = v67
            default:
                if (2 <= uint16(v1)) {
                    for fi68 := 0; fi68 < 7; fi68++ {
                        var v69 float64 = 0.1
                        switch (("foo bar" + (v4 + v4)) + v4) {
                        case "\n\t":
                            if (102 != genFn0((^(v1)))) {
                                fmt.Println(v16, v15)
                                if true {
                                    if ((941 <= -2147483648) || (100 <= 0)) {
                                        if (228 > 0) {
                                            var v70 uint16 = 733
                                            var v71 bool = true
                                            v72 := GenS6{F0: -2147483648}
                                            v73 := GenS7{F0: -1, F1: 127}
                                            v4 = v4
                                            _ = v70
                                            _ = v71
                                            _ = v72
                                            _ = v73
                                        } else {
                                            var v74 string = "foo bar"
                                            if true {
                                                if ((-(v3)) == -32768) {
                                                    if (float64((0.0 - float32((v69 * 0.1)))) == v69) {
                                                        var v75 int64 = 318
                                                        _ = v75
                                                    }
                                                } else {
                                                    var v76 int16 = v3
                                                    _ = v76
                                                }
                                                var v77 float32 = (0.1 + (float32((v69 * 1.5)) * (float32(1.5) * -2.25)))
                                                if false {
                                                    v3 = 156
                                                }
                                                _ = v77
                                            }
                                            var v78 uint = 255
                                            _ = v74
                                            _ = v78
                                        }
                                    }
                                } else {
                                    var v79 int32 = -2147483648
                                    _ = v79
                                }
                                if (0 >= 0) {
                                    v69 -= v69
                                    v80 := GenS6{F0: 433}
                                    _ = v80
                                }
                            }
                        case "a":
                            v81 := GenS2{F0: 0, F1: 526, F2: 42}
                            _ = v81
                        }
                        if (267 > 1) {
                            v2[(fi48 % ((int(v16)) | 1))] = v15
                            v82 := []int{42, fi48, fi48}
                            _ = v82
                        } else {
                            var v83 int64 = 1
                            v15 = v15
                            _ = len(v2)
                            v1 -= (v1 | (v15 << ((821) & 63)))
                            switch (fi68 * -762) {
                            case 32768:
                                v84 := GenS1{F0: (("foo bar" + v4) + ("hello" + (v4 + (v4 + "a"))))}
                                _ = v84
                            case 734:
                                var v85 int = fi68
                                _ = v85
                            default:
                                var v86 uint64 = 1
                                _ = v86
                            }
                            _ = v83
                        }
                        if true {
                            v87 := v2[(fi68 % ((0) | 1))]
                            if (uint64(v15) == 2) {
                                v88 := v2[(fi68 >> ((29) & 63))]
                                var v89 uint8 = 2
                                if (true || (!(3.14159 == (-2.25 - -2.25)))) {
                                    var v90 int32 = -1
                                    _ = v90
                                } else {
                                    var v91 string = (((v4 + (v4 + "hello")) + v4) + ((v4 + (v4 + v4)) + ("hello" + v4)))
                                    _ = v91
                                }
                                switch (fi68 + (fi68 / ((1) | 1))) {
                                case -585:
                                    v92 := []int{42, (fi48 - (fi68 & 2147483647)), 4294967295, (fi68 % ((fi68) | 1))}
                                    _ = v92
                                }
                                var v93 uint = 255
                                v94 := []int{(fi48 * 316), 32767}
                                if false {
                                    v95 := v2[127]
                                    if false {
                                        v69 += 0.0
                                        if false {
                                            var v96 bool = (false || true)
                                            _ = v96
                                        }
                                        v69 += (v69 / v69)
                                    }
                                    _ = v95
                                }
                                v3 -= -743
                                var v97 int32 = 2147483647
                                _ = v88
                                _ = v89
                                _ = v93
                                _ = v94
                                _ = v97
                            }
                            v69 = ((0.1 + -2.25) * -2.25)
                            _ = v87
                        }
                        _ = fi68
                        _ = v69
                    }
                    if false {
                        v15 = 296
                        for fi98 := 0; fi98 < 14; fi98++ {
                            v99 := GenS4{F0: 952, F1: 4294967295}
                            _ = fi98
                            _ = v99
                        }
                        var v100 uint8 = 1
                        _ = v100
                    } else {
                        var v101 uint8 = 2
                        var v102 bool = true
                        var v103 uint64 = 896
                        _ = v101
                        _ = v102
                        _ = v103
                    }
                    var v104 int64 = -224
                    _ = v104
                }
            }
            _ = v66
        } else {
            v105 := GenS1{F0: v4}
            _ = v105
        }
        _ = fi48
    }
    var v106 int64 = 663
    v3 -= 32767
    for fi107 := 0; fi107 < 14; fi107++ {
        delete(v2, fi107)
        _ = fi107
    }
    v15 += v15
    switch 460 {
    case 42:
        v108 := GenS1{F0: v4}
        _ = v108
    }
    v109 := GenS8{F0: (v15 << ((52) & 63)), F1: v3}
    v110 := GenS9{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v106
    _ = v109
    _ = v110
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
