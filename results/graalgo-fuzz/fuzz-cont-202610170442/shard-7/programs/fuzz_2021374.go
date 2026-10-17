package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int }
type GenS3 struct { F0 int64; F1 int; F2 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 uint }

func main() {
    v1 := []string{("" + "\n\t"), ("\n\t" + (((("foo bar" + "foo bar") + "a") + ("a" + "foo bar")) + ("hello" + "a"))), "\n\t"}
    var v2 string = ("a" + "a")
    v3 := make(map[int]int)
    var v4 int8 = 100
    if (!false) {
        var v5 uint64 = 42
        _ = v5
    }
    if (true || (756 <= 0)) {
        var v6 int32 = -144
        _ = v6
    }
    v7 := GenS1{F0: 2147483647, F1: 149, F2: 776}
    _ = len(v1)
    var v8 int16 = -794
    v1 = append(v1, v2)
    switch v2 {
    case "hello":
        if ((32767 != 0) || (18446744073709551615 <= 4294967295)) {
            switch (v2 + "hello") {
            case "":
                if ((v4 ^ -17) <= -1) {
                    v2 = "hello"
                    var v9 float64 = float64(1e10)
                    _ = v9
                }
            case "\n\t":
                var v10 uint8 = 231
                _ = v10
            default:
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            }
            var v11 int8 = (v4 % (((v4 >> ((65535) & 63))) | 1))
            var v12 float32 = 1e10
            _ = v11
            _ = v12
        } else {
            var v13 uint16 = 802
            v13 -= v13
            _ = v13
        }
    case "a":
        for fi14 := 0; fi14 < 19; fi14++ {
            for fi15 := 0; fi15 < 8; fi15++ {
                fmt.Println(v4)
                v4 += (v4 - v4)
                _ = fi15
            }
            v4 = (v4 << ((65535) & 63))
            if false {
                fmt.Println(fi14)
            }
            _ = fi14
        }
    case "\n\t":
        for fi16 := 0; fi16 < 9; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                fmt.Println(v4, v2, fi16)
                if true {
                    switch (-(fi17)) {
                    case -901:
                        switch ("\n\t" + v2) {
                        case "a":
                            var v18 uint64 = 42
                            _ = v18
                        case "":
                            switch 407 {
                            case -76:
                                v19 := GenS2{F0: 2147483648, F1: 65535, F2: int(v8)}
                                _ = v19
                            case -9223372036854775808:
                                switch ((("" + "") + v2) + ("a" + v2)) {
                                case "":
                                    if false {
                                        if true {
                                            delete(v3, 42)
                                        } else {
                                            v20 := GenS1{F0: (fi17 * (^(fi17))), F1: 853, F2: 9223372036854775807}
                                            var v21 int16 = (v8 / (((-(v8))) | 1))
                                            fmt.Println(v8, v8)
                                            _ = v20
                                            _ = v21
                                        }
                                        if (838 <= 42) {
                                            if true {
                                                v4 -= (v4 % ((24) | 1))
                                            } else {
                                                v22 := GenS1{F0: (fi16 - fi17), F1: 18446744073709551615, F2: 0}
                                                switch ("a" + v2) {
                                                case "foo bar":
                                                    var v23 uint64 = 2
                                                    _ = v23
                                                }
                                                if (!false) {
                                                    var v24 int64 = -1
                                                    _ = v24
                                                }
                                                _ = v22
                                            }
                                            var v25 bool = (v4 <= v4)
                                            _ = v25
                                        }
                                        v3[99] = fi17
                                        var v26 float64 = 0.1
                                        _ = v26
                                    }
                                default:
                                    if false {
                                        if false {
                                            v27 := GenS3{F0: -9223372036854775808, F1: (fi17 % ((fi16) | 1)), F2: 2147483647}
                                            var v28 int = fi17
                                            v28 = (v28 | -727)
                                            var v29 int16 = 165
                                            v1 = append(v1, (((v2 + "") + v2) + (v2 + ("hello" + v2))))
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        }
                                        v8 -= 763
                                        v30 := GenS4{F0: (v8 | 32767)}
                                        _ = v30
                                    } else {
                                        if false {
                                            v8 -= v8
                                            v3[(-(fi17))] = fi17
                                        } else {
                                            var v31 int64 = 42
                                            _ = v31
                                        }
                                        if true {
                                            var v32 int64 = 1
                                            if len(v1) > 0 {
                                                v1[0] = ""
                                            }
                                            _ = v32
                                        }
                                        fmt.Println(v4, fi17, fi16)
                                    }
                                }
                            }
                        case "foo bar":
                            if (4294967295 >= 1) {
                                switch fi16 {
                                case -9223372036854775808:
                                    if true {
                                        v4 += v4
                                    }
                                case 1:
                                    var v33 uint16 = 2
                                    _ = v33
                                case 731:
                                    var v34 uint16 = 321
                                    _ = v34
                                default:
                                    if true {
                                        var v35 int16 = 1
                                        v36 := GenS5{F0: 18446744073709551615, F1: (!(!false)), F2: 2147483647}
                                        v37 := GenS5{F0: 559, F1: (false || false), F2: 623}
                                        v38 := GenS1{F0: fi17, F1: 65535, F2: 42}
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v2 = (v2 + "\n\t")
                                        if (210 != 1) {
                                            var v39 int16 = v8
                                            var v40 float64 = 0.0
                                            if (3.14159 >= (3.14159 - -2.25)) {
                                                v4 = 127
                                            }
                                            _ = v39
                                            _ = v40
                                        }
                                    }
                                }
                                if false {
                                    fmt.Println(fi16, fi17, v4)
                                    if false {
                                        fmt.Println(fi17, v2, v8)
                                        fmt.Println(v2)
                                    } else {
                                        if true {
                                            v41 := GenS3{F0: 9223372036854775807, F1: fi16, F2: 2147483647}
                                            var v42 int16 = (v8 << ((1) & 63))
                                            _ = v41
                                            _ = v42
                                        }
                                        if false {
                                            switch v2 {
                                            case "\n\t":
                                                if false {
                                                    v4 -= v4
                                                    var v43 bool = (42 <= 128)
                                                    v44 := GenS5{F0: 0, F1: v43, F2: 204}
                                                    var v45 bool = v43
                                                    _ = v43
                                                    _ = v44
                                                    _ = v45
                                                }
                                            case "hello":
                                                v4 -= v4
                                            case "foo bar":
                                                switch (((v2 + ((v2 + "\n\t") + v2)) + (v2 + ((v2 + v2) + ""))) + v2) {
                                                case "foo bar":
                                                    if true {
                                                        var v46 int32 = 2147483647
                                                        v47 := GenS4{F0: (v8 - v8)}
                                                        v46 = v46
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        if (v8 <= (v8 | v8)) {
                                                            var v48 int16 = (v8 | v8)
                                                            _ = v48
                                                        } else {
                                                            if (90 >= 42) {
                                                                v4 = (v4 - v4)
                                                                switch ("" + v2) {
                                                                case "hello":
                                                                    var v49 float32 = ((-2.25 - 1e10) * 1.5)
                                                                    switch fi16 {
                                                                    case 32767:
                                                                        if false {
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                v50 := GenS3{F0: -918, F1: (^(fi16)), F2: 50}
                                                                                _ = v50
                                                                            }
                                                                        } else {
                                                                            var v51 uint32 = 1
                                                                            _ = v51
                                                                        }
                                                                    case 4294967295:
                                                                        _ = len(v1)
                                                                    default:
                                                                        if len(v1) > 0 {
                                                                            v1[0] = v2
                                                                        }
                                                                    }
                                                                    v52 := GenS1{F0: -502, F1: 1, F2: 0}
                                                                    _ = v49
                                                                    _ = v52
                                                                }
                                                                v4 = v4
                                                                v8 -= v8
                                                                v2 = (v2 + (v2 + v2))
                                                                var v53 string = v2
                                                                _ = v53
                                                            } else {
                                                                v54 := GenS6{F0: 540}
                                                                v2 = v2
                                                                v2 = "a"
                                                                _ = v54
                                                            }
                                                            var v55 uint16 = 128
                                                            _ = v55
                                                        }
                                                        switch v2 {
                                                        case "\n\t":
                                                            if len(v1) > 0 {
                                                                v1[0] = ((v2 + v2) + v2)
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    var v56 string = v2
                                                    _ = v56
                                                }
                                            default:
                                                if ((fi16 ^ fi17) != fi16) {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    fmt.Println(v8, fi17, v8)
                                                    var v57 uint16 = 388
                                                    if false {
                                                        if (!false) {
                                                            v4 = 100
                                                            fmt.Println(v8, fi17)
                                                        } else {
                                                            if (false || (0 == 919)) {
                                                                var v58 float32 = 1e10
                                                                var v59 uint8 = 0
                                                                _ = v58
                                                                _ = v59
                                                            } else {
                                                                var v60 string = (v2 + "")
                                                                v61 := GenS1{F0: fi16, F1: 956, F2: 368}
                                                                delete(v3, int(v8))
                                                                if false {
                                                                    if true {
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            v4 -= (-(v4))
                                                                        case "a":
                                                                            fmt.Println(v2, v57, v2)
                                                                        }
                                                                        var v62 int64 = -9223372036854775808
                                                                        _ = v62
                                                                    } else {
                                                                        var v63 uint8 = 1
                                                                        var v64 float64 = 3.14159
                                                                        _ = v63
                                                                        _ = v64
                                                                    }
                                                                    v60 = v60
                                                                    fmt.Println(fi17)
                                                                    v8 += (^(v8))
                                                                    v8 -= v8
                                                                    var v65 int32 = -1
                                                                    _ = v65
                                                                }
                                                                v66 := GenS5{F0: 410, F1: (!false), F2: 65535}
                                                                _ = v60
                                                                _ = v61
                                                                _ = v66
                                                            }
                                                            fmt.Println(v8)
                                                            _ = len(v1)
                                                        }
                                                        if false {
                                                            var v67 uint32 = 418
                                                            var v68 bool = true
                                                            _ = v67
                                                            _ = v68
                                                        } else {
                                                            v69 := GenS1{F0: fi16, F1: 152, F2: 9223372036854775807}
                                                            v70 := GenS2{F0: 4294967295, F1: 534, F2: fi16}
                                                            var v71 uint32 = 2
                                                            _ = v69
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        v1 = append(v1, ((v2 + v2) + "hello"))
                                                        var v72 uint8 = 158
                                                        _ = v72
                                                    } else {
                                                        if (!(!(false && false))) {
                                                            var v73 int32 = 2147483647
                                                            _ = v73
                                                        }
                                                        fmt.Println(fi16, v57, fi16)
                                                        var v74 uint8 = 128
                                                        _ = v74
                                                    }
                                                    _ = v57
                                                } else {
                                                    switch "hello" {
                                                    case "hello":
                                                        switch (fi17 * (fi17 / ((-1) | 1))) {
                                                        case -762:
                                                            v2 = v2
                                                        case 622:
                                                            switch v2 {
                                                            case "a":
                                                                var v75 uint16 = 255
                                                                _ = v75
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            fmt.Println(v2, fi17)
                                            var v76 int64 = 9223372036854775807
                                            _ = v76
                                        } else {
                                            var v77 int32 = 472
                                            fmt.Println(v77, v4)
                                            _ = v77
                                        }
                                        var v78 float64 = 3.14159
                                        _ = v78
                                    }
                                    var v79 string = "foo bar"
                                    _ = v79
                                }
                            } else {
                                var v80 uint = 18446744073709551615
                                v4 += (^(v4))
                                delete(v3, fi16)
                                v8 = (-(v8))
                                v80 += (v80 >> ((v80) & 63))
                                v81 := GenS3{F0: 42, F1: int(v8), F2: 2147483647}
                                _ = v80
                                _ = v81
                            }
                        }
                    case 2147483647:
                        v8 += -1
                    case 42:
                        var v82 uint = 2
                        _ = v82
                    }
                    fmt.Println(v2, fi16, fi16)
                    if ((286 == 2147483648) || (false && false)) {
                        fmt.Println(fi17, fi17)
                        var v83 bool = (false || true)
                        _ = v83
                    } else {
                        var v84 int32 = 605
                        var v85 int = 772
                        _ = v84
                        _ = v85
                    }
                } else {
                    if false {
                        var v86 int16 = v8
                        _ = v86
                    }
                    v87 := GenS1{F0: fi17, F1: 395, F2: 628}
                    if true {
                        v4 += v4
                    }
                    _ = v87
                }
                v88 := GenS6{F0: 291}
                _ = fi17
                _ = v88
            }
            _ = fi16
        }
    }
    v4 += (v4 & 42)
    for fi89 := 0; fi89 < 13; fi89++ {
        v8 = 0
        v8 += v8
        var v90 bool = (42 != 629)
        _ = fi89
        _ = v90
    }
    var v91 uint32 = 2147483648
    var v92 uint16 = 1
    v8 -= 32767
    v91 += (v91 << ((255) & 63))
    v93 := GenS7{F0: (v8 % ((-32768) | 1)), F1: 481}
    switch "" {
    case "foo bar":
        for fi94 := 0; fi94 < 20; fi94++ {
            switch v2 {
            case "foo bar":
                if (!false) {
                    var v95 int = fi94
                    var v96 uint8 = 194
                    _ = v95
                    _ = v96
                }
            case "a":
                var v97 uint64 = 789
                _ = v97
            case "hello":
                var v98 float32 = 0.1
                _ = v98
            }
            _ = fi94
        }
    case "":
        v99 := GenS3{F0: 772, F1: 32767, F2: -739}
        _ = v99
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v91
    _ = v92
    _ = v93
}
