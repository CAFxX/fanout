package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8; F1 int; F2 int }
func genFn0(p0 int8, p1 string, p2 int64) uint16 {
    fmt.Println(p2, p1, p2)
    var v5 int8 = (p0 / ((p0) | 1))
    var v6 int32 = -2147483648
    p0 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 32767
}
func genFn1() uint64 {
    var v7 int64 = -871
    v8 := []int{-406, -238, 0, 0}
    var v9 uint = 376
    var v10 int16 = 1000
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 string = ((("" + (("foo bar" + ("a" + "hello")) + "foo bar")) + (("" + "hello") + "a")) + "hello")
    var v2 uint16 = 32767
    var v3 float32 = 1e10
    v4 := make(map[int]uint32)
    v3 = v3
    var v11 uint16 = 0
    fmt.Println(v1)
    var v12 bool = (false || (634 != 2147483648))
    fmt.Println(v2, v3, v1)
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 string = (v1 + ((v1 + "") + ("hello" + "\n\t")))
        switch 805 {
        case -67:
            v12 = (42 > 65535)
        case 1:
            fmt.Println(v1, v12)
        case 128:
            v3 -= v3
        }
        var v15 uint16 = v2
        _ = fi13
        _ = v14
        _ = v15
    }
    v12 = (-124 > -1)
    if (0 != 2) {
        switch 239 {
        case -2147483648:
            v1 = ((v1 + "") + (((("foo bar" + ("hello" + v1)) + "a") + ("\n\t" + (("a" + "") + (v1 + v1)))) + (v1 + v1)))
        case 127:
            if ((!(!(3.14159 != 0.1))) && false) {
                if ((v12 || (!((((v2 + v11) <= genFn0(-1, v1, 153)) && v12) || false))) && false) {
                    if true {
                        if ((1 == genFn1()) && (v12 && (!v12))) {
                            switch "" {
                            case "foo bar":
                                switch 2147483647 {
                                case 1:
                                    v4[1] = 2147483648
                                }
                            default:
                                switch v1 {
                                case "a":
                                    v3 = v3
                                case "foo bar":
                                    var v16 uint = 903
                                    _ = v16
                                case "":
                                    var v17 uint16 = genFn0(-128, "hello", 1)
                                    _ = v17
                                }
                            }
                            if v12 {
                                v18 := []int{32768, 32768, 32768, 9223372036854775807}
                                _ = len(v18)
                                _ = v18
                            } else {
                                var v19 int64 = 9223372036854775807
                                if false {
                                    switch ((v1 + (v1 + v1)) + "hello") {
                                    case "":
                                        var v20 bool = v12
                                        _ = v20
                                    case "hello":
                                        fmt.Println(v1, v11, v11)
                                    default:
                                        v2 += (v11 >> ((927) & 63))
                                    }
                                    v11 += v11
                                    var v21 int32 = -2147483648
                                    _ = v21
                                } else {
                                    if ((0.1 * ((0.1 + (((float64(((float32((0.1 - 1.5)) / v3) + (v3 + 1e10))) * (3.14159 / 1.5)) + (3.14159 / 1.0)) / float64((v3 + -2.25)))) * 0.0)) == 0.1) {
                                        switch -2147483648 {
                                        case -982:
                                            fmt.Println(v12)
                                        default:
                                            if true {
                                                for fi22 := 0; fi22 < 12; fi22++ {
                                                    v4[fi22] = 1
                                                    _ = fi22
                                                }
                                            } else {
                                                if true {
                                                    v3 -= ((v3 / 0.1) + (3.14159 / (1.5 / -2.25)))
                                                    for fi23 := 0; fi23 < 6; fi23++ {
                                                        switch v1 {
                                                        case "\n\t":
                                                            v1 = v1
                                                        case "":
                                                            for fi24 := 0; fi24 < 20; fi24++ {
                                                                if (((v2 * (v2 + v11)) <= (^(v2))) && true) {
                                                                    v25 := v4[fi23]
                                                                    v3 += (float32(-2.25) + v3)
                                                                    v1 = v1
                                                                    var v26 string = v1
                                                                    _ = v25
                                                                    _ = v26
                                                                }
                                                                v3 = float32(0.1)
                                                                _ = fi24
                                                            }
                                                        }
                                                        _ = fi23
                                                    }
                                                }
                                                fmt.Println(v1, v19)
                                            }
                                        }
                                        v27 := []int{2147483647, 1, 32768, -9223372036854775808}
                                        _ = v27
                                    }
                                    var v28 int16 = 0
                                    v29 := GenS1{F0: 9223372036854775807, F1: 72, F2: v11}
                                    _ = v28
                                    _ = v29
                                }
                                _ = v19
                            }
                            switch 127 {
                            case 294:
                                delete(v4, 568)
                            default:
                                for fi30 := 0; fi30 < 9; fi30++ {
                                    v31 := GenS2{F0: int64(v11), F1: -628}
                                    _ = fi30
                                    _ = v31
                                }
                            }
                        } else {
                            v2 = genFn0(int8(v11), "hello", -9223372036854775808)
                            v3 = 0.0
                        }
                        var v32 uint = 65535
                        for fi33 := 0; fi33 < 17; fi33++ {
                            v11 -= (v11 << ((v32) & 63))
                            switch fi33 {
                            case 0:
                                switch fi33 {
                                case -2147483648:
                                    v34 := []int{fi33, (^(fi33))}
                                    _ = v34
                                case -758:
                                    v3 -= (float32(float64(0.0)) + ((v3 * 1e10) / float32(float64(v3))))
                                case 739:
                                    v35 := GenS3{F0: -894, F1: ""}
                                    _ = v35
                                }
                            default:
                                if (173 <= 4294967295) {
                                    var v36 uint8 = 1
                                    if (v12 && ((-2.25 * (1.5 * 1.5)) <= 3.14159)) {
                                        for fi37 := 0; fi37 < 18; fi37++ {
                                            v36 = v36
                                            v38 := GenS3{F0: 32767, F1: v1}
                                            var v39 uint32 = 42
                                            _ = fi37
                                            _ = v38
                                            _ = v39
                                        }
                                        if v12 {
                                            v40 := v4[fi33]
                                            _ = v40
                                        }
                                    }
                                    v36 = v36
                                    _ = v36
                                } else {
                                    v41 := GenS3{F0: fi33, F1: ((v1 + ("\n\t" + v1)) + v1)}
                                    v12 = (233 <= 2147483648)
                                    _ = v41
                                }
                            }
                            _ = fi33
                        }
                        _ = v32
                    }
                    v3 = float32(1.5)
                    v12 = false
                }
            } else {
                v42 := []int{-1, -9223372036854775808, 42}
                _ = v42
            }
        }
    } else {
        var v43 int8 = 100
        _ = v43
    }
    var v44 int32 = -534
    delete(v4, 207)
    if (v12 && false) {
        for fi45 := 0; fi45 < 15; fi45++ {
            v46 := GenS4{F0: 1, F1: v44}
            for fi47 := 0; fi47 < 7; fi47++ {
                var v48 uint64 = 18446744073709551615
                var v49 int64 = 1
                if (42 == 0) {
                    if false {
                        v50 := GenS5{F0: 2, F1: 1}
                        var v51 uint = 935
                        _ = v50
                        _ = v51
                    } else {
                        if true {
                            v48 = v48
                            switch (^(fi45)) {
                            case 420:
                                if ((v1 <= ("foo bar" + v1)) || false) {
                                    var v52 uint32 = 42
                                    _ = v52
                                }
                            case 128:
                                v53 := GenS2{F0: -800, F1: 1000}
                                _ = v53
                            }
                            switch (fi47 & -201) {
                            case 32767:
                                switch "hello" {
                                case "\n\t":
                                    var v54 int8 = -88
                                    _ = v54
                                }
                            case -1:
                                v55 := GenS6{F0: v1, F1: 396}
                                _ = v55
                            case 0:
                                v1 = (v1 + v1)
                            default:
                                fmt.Println(v1, v11, v49)
                            }
                            var v56 uint32 = 848
                            var v57 uint8 = 42
                            _ = v56
                            _ = v57
                        } else {
                            var v58 uint64 = 1
                            fmt.Println(v12, v58)
                            fmt.Println(v1, v2, v1)
                            _ = v58
                        }
                        var v59 float64 = 0.1
                        _ = v59
                    }
                    v60 := []int{220, fi45, (fi45 * (fi45 << ((18446744073709551615) & 63))), fi45}
                    delete(v4, 32767)
                    _ = v60
                }
                if (348 > v2) {
                    fmt.Println(v49, v48)
                    v61 := GenS1{F0: v49, F1: 0, F2: (v2 | v11)}
                    _ = v61
                }
                if (!v12) {
                    if v12 {
                        v48 = (^(v48))
                        switch (^(fi47)) {
                        case 0:
                            v11 -= (v11 + genFn0(0, "\n\t", 9223372036854775807))
                        case 1:
                            var v62 uint32 = 2147483648
                            _ = v62
                        }
                        var v63 int32 = v44
                        v63 -= v44
                        v3 = float32((0.1 - 3.14159))
                        v63 = 0
                        _ = v63
                    }
                }
                delete(v4, fi45)
                fmt.Println(v44, fi45)
                v1 = (v1 + v1)
                _ = fi47
                _ = v48
                _ = v49
            }
            _ = fi45
            _ = v46
        }
    } else {
        var v64 uint = 65535
        _ = v64
    }
    if (!(v12 && (!v12))) {
        if false {
            v65 := []int{32767, 128, 9223372036854775807}
            if len(v65) >= 2 {
                v65 = v65[1:len(v65)]
            }
            _ = v65
        } else {
            var v66 int32 = v44
            for fi67 := 0; fi67 < 12; fi67++ {
                v11 += 42
                v68 := GenS2{F0: 0, F1: -360}
                _ = fi67
                _ = v68
            }
            v11 = (v2 - v2)
            _ = v66
        }
        var v69 int64 = 735
        if (0 != 1) {
            switch 32768 {
            case 127:
                for fi70 := 0; fi70 < 4; fi70++ {
                    delete(v4, (fi70 << ((255) & 63)))
                    _ = fi70
                }
            case -9223372036854775808:
                fmt.Println(v69)
            default:
                if v12 {
                    switch (v1 + v1) {
                    case "\n\t":
                        switch (v1 + "\n\t") {
                        case "":
                            v71 := GenS4{F0: 18446744073709551615, F1: v44}
                            _ = v71
                        case "\n\t":
                            var v72 uint32 = 2
                            _ = v72
                        case "hello":
                            if false {
                                for fi73 := 0; fi73 < 19; fi73++ {
                                    v11 += 42
                                    _ = fi73
                                }
                                v74 := GenS3{F0: -728, F1: ((("a" + "hello") + v1) + v1)}
                                v75 := []int{}
                                _ = v74
                                _ = v75
                            } else {
                                v11 = (v11 - 255)
                                var v76 string = ""
                                _ = v76
                            }
                        default:
                            v77 := []int{30, -536}
                            _ = v77
                        }
                    }
                    switch v1 {
                    case "\n\t":
                        v78 := GenS5{F0: 680, F1: 80}
                        _ = v78
                    }
                    if v12 {
                        if (!false) {
                            for fi79 := 0; fi79 < 18; fi79++ {
                                for fi80 := 0; fi80 < 18; fi80++ {
                                    fmt.Println(fi79)
                                    _ = len(v4)
                                    _ = fi80
                                }
                                fmt.Println(v11, v12)
                                _ = fi79
                            }
                        }
                    } else {
                        var v81 uint64 = 18446744073709551615
                        _ = v81
                    }
                }
            }
            if (!(!v12)) {
                v11 += (v2 + (v11 % ((v2) | 1)))
                v82 := GenS7{F0: 32767, F1: 1, F2: -1}
                for fi83 := 0; fi83 < 18; fi83++ {
                    if v12 {
                        v11 = (v2 / ((v2) | 1))
                    } else {
                        var v84 int64 = (v69 ^ v69)
                        _ = v84
                    }
                    switch v1 {
                    case "":
                        v11 += genFn0(0, ("hello" + v1), (v69 % (((v69 & 1)) | 1)))
                    case "foo bar":
                        switch v1 {
                        case "":
                            var v85 float32 = v3
                            _ = v85
                        }
                    }
                    _ = fi83
                }
                _ = v82
            } else {
                v69 = v69
            }
        }
        _ = v69
    }
    fmt.Println(v1)
    if (123 < 0) {
        switch 32768 {
        case -647:
            if v12 {
                v4[-1] = 0
                var v86 string = "foo bar"
                _ = v86
            } else {
                if (-232 != -1) {
                    v44 = (^(v44))
                    v87 := GenS6{F0: ("" + v1), F1: 698}
                    for fi88 := 0; fi88 < 15; fi88++ {
                        v3 = (v3 + ((v3 - float32(3.14159)) + (1.5 + 0.0)))
                        v11 = v2
                        v89 := []int{1, (-(fi88))}
                        _ = fi88
                        _ = v89
                    }
                    _ = v87
                } else {
                    v44 += v44
                }
                delete(v4, -949)
                switch 2147483647 {
                case 1:
                    v44 += v44
                case -154:
                    v2 -= v2
                case 32767:
                    for fi90 := 0; fi90 < 3; fi90++ {
                        v12 = v12
                        delete(v4, -2147483648)
                        var v91 int8 = -72
                        _ = fi90
                        _ = v91
                    }
                default:
                    v92 := GenS8{F0: 1}
                    _ = v92
                }
            }
        case 2147483647:
            v93 := GenS5{F0: 1, F1: 255}
            _ = v93
        default:
            v94 := GenS6{F0: v1, F1: 2}
            _ = v94
        }
    } else {
        for fi95 := 0; fi95 < 18; fi95++ {
            v96 := GenS9{F0: 255, F1: (fi95 % (((fi95 % (((fi95 & fi95)) | 1))) | 1)), F2: fi95}
            _ = fi95
            _ = v96
        }
    }
    v97 := []int{}
    v3 = (1e10 * (float32(1.5) - 0.0))
    v98 := GenS5{F0: 65535, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v97:", len(v97), v97)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v97
    _ = v98
}
