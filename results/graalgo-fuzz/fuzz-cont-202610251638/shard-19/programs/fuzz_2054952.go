package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 int; F2 int16 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    var v3 string = "\n\t"
    v4 := make(map[string]bool)
    v5 := make(map[uint32]bool)
    var v6 float32 = 0.1
    v5[797] = false
    v7 := []int{-77, 2147483647, -353, -780}
    v3 = "\n\t"
    v8 := GenS1{F0: "hello", F1: 1}
    if (true || false) {
        v9 := GenS2{F0: (4294967295 != -2147483648), F1: 0, F2: 1}
        _ = v9
    }
    fmt.Println(v6)
    if ((("hello" + "foo bar") + v3) == v3) {
        _ = len(v7)
        v3 = ("\n\t" + "hello")
    }
    if true {
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(fi10)
            _ = fi10
        }
        if false {
            if (1 >= 65535) {
                fmt.Println(v3)
                var v11 uint = 645
                _ = v11
            }
            v3 = "a"
            var v12 float32 = float32((1.5 / 1e10))
            _ = v12
        } else {
            if (!false) {
                v3 = "foo bar"
                for fi13 := 0; fi13 < 16; fi13++ {
                    switch 0 {
                    case 998:
                        for fi14 := 0; fi14 < 20; fi14++ {
                            if false {
                                var v15 uint64 = 872
                                switch (fi13 % (((fi13 / ((fi14) | 1))) | 1)) {
                                case -353:
                                    if (("hello" >= "") || (100 == -105)) {
                                        v6 += ((((1.5 / v6) / (((-2.25 * (-2.25 + 1e10)) - (v6 / v6)) + v6)) / v6) / 1e10)
                                    } else {
                                        if false {
                                            v15 += (v15 | (^(v15)))
                                            v16 := GenS1{F0: (v3 + v3), F1: 65535}
                                            switch fi13 {
                                            case -9223372036854775808:
                                                if ("hello" < v3) {
                                                    v17 := GenS1{F0: (v3 + ((("\n\t" + (("" + "foo bar") + ("" + (v3 + "\n\t")))) + v3) + v3)), F1: 2}
                                                    var v18 uint8 = 2
                                                    _ = v17
                                                    _ = v18
                                                }
                                            case 652:
                                                switch "hello" {
                                                case "a":
                                                    var v19 bool = (!(true && true))
                                                    _ = v19
                                                }
                                            case -2147483648:
                                                v15 = 42
                                            default:
                                                if ((1 <= 0) && false) {
                                                    v20 := v2["foo bar"]
                                                    if ((false || true) || false) {
                                                        if (42 <= 1) {
                                                            switch -713 {
                                                            case -1:
                                                                switch fi13 {
                                                                case -9223372036854775808:
                                                                    v15 -= v15
                                                                case -964:
                                                                    var v21 float64 = 0.1
                                                                    _ = v21
                                                                case 42:
                                                                    v3 = "a"
                                                                    var v22 float32 = (v6 - v6)
                                                                    v1[("foo bar" + v3)] = ((v3 + "") + ((v3 + v3) + "\n\t"))
                                                                    switch 2147483647 {
                                                                    case 1:
                                                                        if false {
                                                                            var v23 int8 = 1
                                                                            _ = v23
                                                                        } else {
                                                                            var v24 int32 = 53
                                                                            v25 := GenS3{F0: 1}
                                                                            v3 = "hello"
                                                                            fmt.Println(v20, fi14, v24)
                                                                            v24 += (v24 ^ v24)
                                                                            switch v3 {
                                                                            case "foo bar":
                                                                                v26 := GenS1{F0: v3, F1: 1}
                                                                                _ = v26
                                                                            case "\n\t":
                                                                                var v27 uint = 179
                                                                                _ = v27
                                                                            case "":
                                                                                switch (fi14 << ((65535) & 63)) {
                                                                                case 4294967295:
                                                                                    v5[4294967295] = true
                                                                                case -2147483648:
                                                                                    v7 = append(v7, v20)
                                                                                case -1:
                                                                                    switch (fi13 - -9223372036854775808) {
                                                                                    case 1:
                                                                                        if (!true) {
                                                                                            switch ("hello" + v3) {
                                                                                            case "":
                                                                                                v24 = -2147483648
                                                                                            case "hello":
                                                                                                v15 -= v15
                                                                                            default:
                                                                                                switch (fi13 - 128) {
                                                                                                case 32768:
                                                                                                    v28 := GenS3{F0: 510}
                                                                                                    _ = v28
                                                                                                case -1:
                                                                                                    var v29 uint32 = 65535
                                                                                                    _ = v29
                                                                                                case 4294967295:
                                                                                                    v30 := v2[v3]
                                                                                                    _ = v30
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case 128:
                                                                                        var v31 float64 = 1.5
                                                                                        _ = v31
                                                                                    case 0:
                                                                                        v6 -= ((3.14159 / (v6 + v6)) * v22)
                                                                                    }
                                                                                }
                                                                            }
                                                                            v32 := GenS1{F0: (v3 + ((("hello" + "") + "a") + "hello")), F1: 32767}
                                                                            v3 = "hello"
                                                                            v33 := GenS3{F0: 2147483647}
                                                                            var v34 string = v3
                                                                            var v35 int64 = 217
                                                                            _ = v24
                                                                            _ = v25
                                                                            _ = v32
                                                                            _ = v33
                                                                            _ = v34
                                                                            _ = v35
                                                                        }
                                                                    default:
                                                                        v36 := GenS2{F0: ((1e10 * (1e10 - 0.0)) > 0.1), F1: 78, F2: (v15 * 1)}
                                                                        _ = v36
                                                                    }
                                                                    var v37 int32 = 0
                                                                    _ = v22
                                                                    _ = v37
                                                                default:
                                                                    if len(v7) > 0 {
                                                                        v7[0] = fi14
                                                                    }
                                                                }
                                                            case -9223372036854775808:
                                                                fmt.Println(fi13, v6, v3)
                                                                _ = len(v4)
                                                                var v38 uint8 = 42
                                                                _ = v38
                                                            case 32767:
                                                                var v39 uint16 = 32767
                                                                _ = v39
                                                            }
                                                        } else {
                                                            var v40 uint = 42
                                                            _ = v40
                                                        }
                                                        if true {
                                                            v20 = 32767
                                                        }
                                                    }
                                                    _ = v20
                                                } else {
                                                    fmt.Println(v6, v3, fi13)
                                                    var v41 string = v3
                                                    var v42 uint16 = 65535
                                                    _ = v41
                                                    _ = v42
                                                }
                                            }
                                            _ = v16
                                        } else {
                                            if (!true) {
                                                fmt.Println(fi14, v6, v15)
                                            }
                                            delete(v5, 1)
                                        }
                                        switch 42 {
                                        case 2147483647:
                                            var v43 float64 = ((3.14159 - 3.14159) - 1.5)
                                            _ = v43
                                        case 127:
                                            fmt.Println(v3, fi14)
                                        default:
                                            fmt.Println(v15, fi13)
                                        }
                                        v15 = v15
                                    }
                                }
                                fmt.Println(fi14, v15)
                                _ = v15
                            }
                            v44 := GenS1{F0: ("a" + v3), F1: 667}
                            _ = fi14
                            _ = v44
                        }
                    case 1:
                        v45 := GenS3{F0: 42}
                        _ = v45
                    default:
                        var v46 uint16 = 66
                        _ = v46
                    }
                    _ = fi13
                }
            } else {
                if (!false) {
                    var v47 int = 42
                    _ = v47
                } else {
                    v48 := GenS1{F0: "", F1: 255}
                    var v49 int16 = 1000
                    _ = v48
                    _ = v49
                }
                for fi50 := 0; fi50 < 10; fi50++ {
                    switch fi50 {
                    case 2147483647:
                        var v51 uint8 = 228
                        _ = v51
                    default:
                        for fi52 := 0; fi52 < 2; fi52++ {
                            v2[(v3 + "\n\t")] = fi50
                            var v53 int8 = 1
                            fmt.Println(fi52)
                            _ = fi52
                            _ = v53
                        }
                    }
                    _ = fi50
                }
            }
            if len(v7) > 0 {
                v7[0] = 2147483647
            }
            v6 = float32(float64((1e10 / 0.1)))
        }
        switch v3 {
        case "foo bar":
            if true {
                fmt.Println(v6)
            } else {
                v3 = ((((v3 + "\n\t") + v3) + "") + v3)
                var v54 uint = 42
                _ = v54
            }
        default:
            v55 := v5[2147483647]
            _ = v55
        }
    } else {
        if (!(42 == 0)) {
            for fi56 := 0; fi56 < 10; fi56++ {
                v3 = "hello"
                switch (fi56 + (fi56 % ((fi56) | 1))) {
                case 871:
                    v6 = v6
                default:
                    var v57 uint64 = 1
                    _ = v57
                }
                v6 += v6
                _ = fi56
            }
            v3 = ("foo bar" + (("a" + "foo bar") + v3))
            fmt.Println(v3, v3)
        }
        v58 := v1["hello"]
        var v59 uint16 = 65535
        _ = v58
        _ = v59
    }
    var v60 int8 = -100
    if false {
        if false {
            fmt.Println(v6, v60, v3)
            v6 -= v6
        } else {
            var v61 float32 = 1.5
            _ = v61
        }
    } else {
        var v62 int = 721
        if (-464 >= 1000) {
            v62 = -9223372036854775808
            v63 := GenS3{F0: 2147483648}
            _ = v63
        }
        var v64 int32 = 0
        _ = v62
        _ = v64
    }
    _ = len(v5)
    var v65 float32 = (0.0 + 1e10)
    switch v3 {
    case "foo bar":
        if (!(-210 < 712)) {
            for fi66 := 0; fi66 < 15; fi66++ {
                v67 := GenS3{F0: 65535}
                _ = fi66
                _ = v67
            }
            if true {
                var v68 int8 = 1
                v69 := GenS4{F0: 18446744073709551615, F1: 856, F2: 32767}
                fmt.Println(v60, v65)
                _ = v68
                _ = v69
            }
        } else {
            switch v3 {
            case "hello":
                v6 = (((((float32(1e10) + (v65 + v6)) / float32(1.5)) / ((v6 / v6) + ((-2.25 * v6) - (v65 - v65)))) / (-2.25 / (v6 / v65))) + (v65 / float32(((0.1 * (3.14159 + float64(float32((((-2.25 / 1.0) / (0.0 + (0.1 / 1e10))) / 3.14159))))) * (-2.25 / ((0.1 - 3.14159) * float64((-2.25 * float32(1e10)))))))))
            case "foo bar":
                switch ("\n\t" + v3) {
                case "\n\t":
                    v7 = append(v7, 32767)
                }
            case "a":
                for fi70 := 0; fi70 < 9; fi70++ {
                    if false {
                        if (!(!true)) {
                            for fi71 := 0; fi71 < 2; fi71++ {
                                var v72 float64 = 0.0
                                v73 := GenS4{F0: 550, F1: (fi71 - fi70), F2: -32768}
                                v4[v3] = false
                                _ = fi71
                                _ = v72
                                _ = v73
                            }
                            if (false || ((v3 + v3) > v3)) {
                                var v74 float32 = (((v65 / 1.5) + float32(-2.25)) * float32(0.0))
                                var v75 int8 = -22
                                v76 := GenS2{F0: true, F1: 129, F2: 42}
                                _ = v74
                                _ = v75
                                _ = v76
                            } else {
                                var v77 int32 = 1
                                _ = v77
                            }
                            v78 := GenS2{F0: (v60 >= 0), F1: 1, F2: 4294967295}
                            _ = v78
                        } else {
                            if true {
                                v65 -= (((-2.25 * v6) * (((((v65 / (v6 + 1.5)) + float32((-2.25 * float64((v65 - v6))))) * (0.1 / 3.14159)) + v6) + v6)) - 0.1)
                            }
                        }
                    } else {
                        for fi79 := 0; fi79 < 12; fi79++ {
                            if false {
                                var v80 int8 = v60
                                v6 -= (v65 / ((v6 + ((v6 / float32(-2.25)) + ((0.1 * v65) + v6))) / 3.14159))
                                fmt.Println(v6, v65, v80)
                                _ = v80
                            } else {
                                if true {
                                    var v81 int64 = 712
                                    var v82 int = (fi70 + (fi70 ^ 4294967295))
                                    _ = v81
                                    _ = v82
                                } else {
                                    switch "hello" {
                                    case "\n\t":
                                        if false {
                                            var v83 uint16 = 255
                                            _ = v83
                                        }
                                    }
                                    switch fi70 {
                                    case -905:
                                        fmt.Println(v3, v6, v65)
                                    case 374:
                                        var v84 string = "\n\t"
                                        _ = v84
                                    case -769:
                                        if false {
                                            fmt.Println(v65, v3, fi70)
                                            if (4294967295 >= 522) {
                                                v60 = 3
                                                if ((fi70 ^ (^(fi70))) == (fi70 ^ (fi70 & (fi70 / ((fi79) | 1))))) {
                                                    var v85 int16 = -1
                                                    v86 := GenS5{F0: fi79, F1: 310}
                                                    var v87 uint64 = 0
                                                    if false {
                                                        var v88 int64 = -1
                                                        _ = v88
                                                    }
                                                    _ = v85
                                                    _ = v86
                                                    _ = v87
                                                } else {
                                                    v3 = v3
                                                }
                                            }
                                        } else {
                                            if ((102 < 127) || false) {
                                                v89 := GenS2{F0: (((!false) && (42 != 2)) || (229 <= 0)), F1: 127, F2: 697}
                                                v60 = (-(v60))
                                                switch v3 {
                                                case "":
                                                    _ = len(v2)
                                                case "hello":
                                                    var v90 uint16 = 0
                                                    _ = v90
                                                default:
                                                    v60 = v60
                                                    switch 42 {
                                                    case -692:
                                                        v4[v3] = false
                                                    case 0:
                                                        v65 = v6
                                                    case 2147483647:
                                                        _ = len(v1)
                                                    }
                                                    var v91 float32 = v65
                                                    var v92 uint = 1
                                                    v93 := GenS6{F0: -319}
                                                    v3 = "\n\t"
                                                    _ = v91
                                                    _ = v92
                                                    _ = v93
                                                }
                                                _ = v89
                                            } else {
                                                v6 = -2.25
                                            }
                                            v3 = (v3 + "")
                                        }
                                    default:
                                        switch (v3 + v3) {
                                        case "a":
                                            var v94 uint64 = 2
                                            _ = v94
                                        case "hello":
                                            if true {
                                                var v95 int8 = 127
                                                _ = v95
                                            }
                                        case "":
                                            v60 = 100
                                        }
                                        var v96 uint = 42
                                        v65 = v65
                                        _ = v96
                                    }
                                    v7 = append(v7, 1)
                                    v60 = 56
                                    v3 = "hello"
                                }
                                if (1.5 >= v65) {
                                    v7 = append(v7, (fi70 + (-(fi70))))
                                    v60 = 13
                                } else {
                                    fmt.Println(fi70)
                                }
                                v3 = v3
                            }
                            fmt.Println(fi79, v60, fi79)
                            _ = fi79
                        }
                    }
                    if len(v7) > 0 {
                        v7[0] = 4294967295
                    }
                    _ = fi70
                }
            default:
                v97 := GenS7{F0: -1}
                _ = v97
            }
        }
    }
    v60 += v60
    var v98 bool = true
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v65
    _ = v7
    _ = v8
    _ = v98
}
