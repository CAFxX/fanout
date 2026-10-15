package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 bool; F1 bool; F2 int16 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint }
type GenS5 struct { F0 uint; F1 uint8; F2 int }
type GenS6 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16; F1 int64; F2 uint }
type GenS11 struct { F0 bool; F1 string; F2 int64 }
type GenS12 struct { F0 uint8; F1 int32 }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            switch (-(fi4)) {
            case -2147483648:
                var v6 int16 = -247
                _ = v6
            default:
                var v7 uint = 776
                _ = v7
            }
            _ = fi5
        }
        v8 := GenS1{F0: -9223372036854775808}
        _ = fi4
        _ = v8
    }
    if (((!true) && true) && (-1 != 0)) {
        switch "\n\t" {
        case "a":
            v9 := GenS2{F0: 42, F1: 175}
            _ = v9
        }
        var v10 int16 = 1
        var v11 int16 = v10
        _ = v10
        _ = v11
    } else {
        if (42 < -128) {
            var v12 int8 = 1
            var v13 int32 = -165
            v14 := GenS1{F0: 717}
            _ = v12
            _ = v13
            _ = v14
        }
        v15 := []int{42, -183, 2147483647, 489}
        if len(v15) > 0 {
            v15[0] = -488
        }
        _ = v15
    }
    v16 := GenS3{F0: (("a" + ("a" + "")) == ""), F1: (2 != 2), F2: 42}
    v17 := []int{361, -1, 886}
    var v18 float64 = 1.5
    _ = v16
    _ = v17
    _ = v18
    return 2
}

func main() {
    var v1 uint64 = 290
    v2 := make(map[uint32]string)
    var v3 bool = (!false)
    v3 = (!(true && ((-893 == 2147483647) && (!v3))))
    if (true && (2147483647 != 2147483647)) {
        if v3 {
            _ = len(v2)
        }
        v1 = v1
        switch 520 {
        case 9223372036854775807:
            if false {
                switch "hello" {
                case "\n\t":
                    v1 -= (v1 >> ((700) & 63))
                case "a":
                    switch 4294967295 {
                    case -836:
                        v19 := []int{42, 2147483647, -9223372036854775808, 9223372036854775807}
                        _ = v19
                    }
                default:
                    v1 += (v1 % ((v1) | 1))
                }
                switch "hello" {
                case "hello":
                    switch "" {
                    case "":
                        fmt.Println(v3)
                    case "hello":
                        fmt.Println(v3)
                    default:
                        v20 := []int{42}
                        _ = v20
                    }
                case "":
                    var v21 int16 = -32768
                    _ = v21
                case "foo bar":
                    var v22 uint32 = 2147483648
                    _ = v22
                }
            }
        case -140:
            v23 := GenS4{F0: 915, F1: 1, F2: 525}
            _ = v23
        default:
            for fi24 := 0; fi24 < 18; fi24++ {
                if (!v3) {
                    v3 = (220 == 26)
                }
                _ = fi24
            }
        }
    } else {
        v3 = ((1.5 - 0.0) == 1e10)
        for fi25 := 0; fi25 < 14; fi25++ {
            fmt.Println(fi25, v3)
            for fi26 := 0; fi26 < 17; fi26++ {
                if (2 == 762) {
                    v2[0] = "a"
                    var v27 uint64 = v1
                    switch (^(fi26)) {
                    case 726:
                        var v28 uint8 = 0
                        _ = v28
                    case 127:
                        switch fi26 {
                        case 4294967295:
                            if v3 {
                                v29 := []int{-1, fi26, fi26, fi25}
                                if v3 {
                                    delete(v2, 65535)
                                    switch 4294967295 {
                                    case 4294967295:
                                        switch fi25 {
                                        case -9223372036854775808:
                                            v30 := GenS3{F0: (("a" + "a") > (("" + "hello") + ("" + ""))), F1: (-128 >= 86), F2: -32768}
                                            var v31 int64 = 1
                                            switch fi25 {
                                            case 9223372036854775807:
                                                if (2 != 255) {
                                                    var v32 string = "a"
                                                    v31 -= v31
                                                    v2[65535] = ""
                                                    _ = v32
                                                } else {
                                                    var v33 uint32 = 65535
                                                    var v34 bool = true
                                                    _ = v33
                                                    _ = v34
                                                }
                                            case 42:
                                                var v35 bool = false
                                                _ = v35
                                            default:
                                                var v36 int8 = 1
                                                _ = v36
                                            }
                                            if true {
                                                v37 := GenS5{F0: 18446744073709551615, F1: 49, F2: fi25}
                                                v1 = (v27 << ((697) & 63))
                                                _ = v37
                                            } else {
                                                v38 := GenS6{F0: 0, F1: 20, F2: 42}
                                                v39 := GenS4{F0: 0, F1: v31, F2: 2}
                                                v40 := GenS6{F0: 42, F1: (-(v31)), F2: 65535}
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                            }
                                            if false {
                                                switch (-(fi25)) {
                                                case -368:
                                                    var v41 string = "hello"
                                                    _ = v41
                                                case 32768:
                                                    v42 := GenS4{F0: 32767, F1: 9223372036854775807, F2: 554}
                                                    _ = v42
                                                case 4294967295:
                                                    v43 := GenS7{F0: 0, F1: 42}
                                                    var v44 uint8 = 42
                                                    v27 = 519
                                                    _ = v43
                                                    _ = v44
                                                default:
                                                    var v45 int64 = (v31 / ((v31) | 1))
                                                    _ = v45
                                                }
                                            }
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                            _ = v30
                                            _ = v31
                                        case 42:
                                            if v3 {
                                                if v3 {
                                                    if (v3 || (42 == 447)) {
                                                        v27 = 735
                                                    } else {
                                                        var v46 int64 = 0
                                                        fmt.Println(fi26, v46, v46)
                                                        _ = v46
                                                    }
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if false {
                                                            var v47 uint8 = 127
                                                            switch (^(fi25)) {
                                                            case -79:
                                                                switch "" {
                                                                case "":
                                                                    v47 -= v47
                                                                }
                                                            case 32767:
                                                                if (-32768 != 42) {
                                                                    v47 = v47
                                                                }
                                                            case 329:
                                                                if v3 {
                                                                    var v48 float32 = 0.1
                                                                    v49 := GenS1{F0: 32767}
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                            default:
                                                                v3 = v3
                                                            }
                                                            var v50 float32 = 3.14159
                                                            _ = v47
                                                            _ = v50
                                                        }
                                                        v27 += (v27 ^ (v1 & (v1 | v1)))
                                                    default:
                                                        var v51 uint16 = 255
                                                        _ = v51
                                                    }
                                                }
                                            } else {
                                                v2[42] = "a"
                                            }
                                        case -256:
                                            var v52 int64 = 9223372036854775807
                                            _ = v52
                                        default:
                                            var v53 float32 = 0.0
                                            _ = v53
                                        }
                                    default:
                                        v3 = (fi26 > 32768)
                                    }
                                    if v3 {
                                        var v54 int8 = -18
                                        v27 += (v1 << ((18446744073709551615) & 63))
                                        var v55 uint8 = 0
                                        _ = v54
                                        _ = v55
                                    } else {
                                        fmt.Println(fi25, v1)
                                        if (0.0 < 1.5) {
                                            v56 := GenS8{F0: 1, F1: 0}
                                            _ = v56
                                        } else {
                                            if v3 {
                                                var v57 string = "hello"
                                                v57 = (("hello" + "hello") + ("foo bar" + v57))
                                                v58 := v2[773]
                                                v59 := GenS9{F0: 42}
                                                var v60 uint64 = (v1 | v1)
                                                _ = v57
                                                _ = v58
                                                _ = v59
                                                _ = v60
                                            }
                                            var v61 uint16 = 32767
                                            fmt.Println(fi25, fi25, v1)
                                            _ = v61
                                        }
                                    }
                                    v62 := GenS10{F0: 255, F1: 0, F2: 255}
                                    _ = v62
                                } else {
                                    _ = len(v29)
                                }
                                _ = v29
                            }
                        case 32768:
                            v63 := GenS4{F0: 255, F1: 150, F2: 255}
                            _ = v63
                        default:
                            v3 = v3
                        }
                    case 32767:
                        var v64 int16 = 392
                        _ = v64
                    }
                    _ = v27
                }
                v1 += 2
                v65 := []int{fi26}
                _ = len(v65)
                _ = fi26
                _ = v65
            }
            _ = fi25
        }
        v3 = false
    }
    v1 = (v1 << ((216) & 63))
    v66 := []int{-2147483648, -1, 32767}
    switch 260 {
    case -2147483648:
        for fi67 := 0; fi67 < 4; fi67++ {
            for fi68 := 0; fi68 < 16; fi68++ {
                var v69 float64 = 3.14159
                fmt.Println(v3, fi68)
                _ = fi68
                _ = v69
            }
            _ = fi67
        }
    case 42:
        v3 = v3
    case 584:
        for fi70 := 0; fi70 < 13; fi70++ {
            fmt.Println(v3, fi70, v3)
            var v71 int16 = 0
            fmt.Println(v71)
            _ = fi70
            _ = v71
        }
    default:
        v3 = (4294967295 <= -2147483648)
    }
    for fi72 := 0; fi72 < 3; fi72++ {
        for fi73 := 0; fi73 < 17; fi73++ {
            var v74 uint8 = 61
            _ = fi73
            _ = v74
        }
        v75 := GenS3{F0: (100 >= 100), F1: ((180 != 1) && ((1e10 / 1.0) == -2.25)), F2: -469}
        v76 := GenS7{F0: 78, F1: 1}
        _ = fi72
        _ = v75
        _ = v76
    }
    var v77 uint64 = 55
    v77 = v77
    if (((!(!v3)) || v3) && v3) {
        if v3 {
            switch "a" {
            case "":
                if v3 {
                    v78 := GenS3{F0: v3, F1: (v3 && (0 != int16(v1))), F2: -369}
                    v77 = 842
                    var v79 int32 = 2147483647
                    _ = v78
                    _ = v79
                } else {
                    var v80 uint8 = 1
                    _ = v80
                }
            case "\n\t":
                v81 := GenS5{F0: 255, F1: 2, F2: 642}
                _ = v81
            case "hello":
                for fi82 := 0; fi82 < 16; fi82++ {
                    fmt.Println(fi82)
                    _ = fi82
                }
            }
        }
        v83 := GenS8{F0: 42, F1: 0}
        _ = v83
    }
    fmt.Println(v77, v77, v1)
    for fi84 := 0; fi84 < 5; fi84++ {
        for fi85 := 0; fi85 < 15; fi85++ {
            v86 := GenS11{F0: v3, F1: ("a" + ""), F2: -1}
            v66 = append(v66, -236)
            fmt.Println(v3)
            _ = fi85
            _ = v86
        }
        for fi87 := 0; fi87 < 4; fi87++ {
            switch "a" {
            case "foo bar":
                if v3 {
                    _ = len(v66)
                    v77 = 18446744073709551615
                    var v88 int = 2147483647
                    _ = v88
                } else {
                    v89 := GenS12{F0: 0, F1: 0}
                    v77 = v77
                    var v90 uint8 = 128
                    _ = v89
                    _ = v90
                }
            case "":
                var v91 uint32 = 688
                _ = v91
            case "hello":
                v1 = v1
            default:
                v1 -= (v77 & (v1 % ((genFn0()) | 1)))
            }
            v3 = v3
            _ = fi87
        }
        var v92 float64 = 0.1
        _ = fi84
        _ = v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v66
    _ = v77
}
