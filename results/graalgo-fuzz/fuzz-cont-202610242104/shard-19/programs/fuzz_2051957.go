package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint; F1 uint64; F2 bool }

func main() {
    var v1 string = "a"
    v2 := []int{}
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    if (1000 >= 32767) {
        var v5 uint8 = 255
        v5 += v5
        var v6 uint = 4294967295
        _ = v5
        _ = v6
    }
    fmt.Println(v1)
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            switch v1 {
            case "a":
                v9 := GenS1{F0: 0, F1: 1}
                _ = v9
            case "\n\t":
                var v10 float64 = (1.5 / 1.5)
                _ = v10
            case "hello":
                var v11 int16 = 0
                _ = v11
            default:
                var v12 int = 127
                _ = v12
            }
            _ = fi8
        }
        var v13 string = (v1 + v1)
        _ = fi7
        _ = v13
    }
    v14 := GenS2{F0: 32, F1: v1}
    switch 343 {
    case 1:
        for fi15 := 0; fi15 < 20; fi15++ {
            if (!(5 != 2147483647)) {
                var v16 int8 = 100
                var v17 int16 = -1
                _ = v16
                _ = v17
            } else {
                v1 = v1
                v1 = "\n\t"
            }
            _ = fi15
        }
    case 9223372036854775807:
        v18 := GenS1{F0: 255, F1: 42}
        _ = v18
    case -9223372036854775808:
        v19 := GenS1{F0: 16, F1: 18446744073709551615}
        _ = v19
    default:
        v20 := GenS1{F0: 127, F1: 0}
        _ = v20
    }
    var v21 uint = 42
    v21 = v21
    _ = len(v2)
    if (true && (!false)) {
        if (v21 < (v21 + v21)) {
            v22 := GenS3{F0: 89}
            if len(v2) > 0 {
                v2[0] = 9223372036854775807
            }
            _ = v22
        }
        if (706 != 9223372036854775807) {
            fmt.Println(v21, v1)
            if ((!((v21 & 710) == v21)) && (!false)) {
                var v23 int32 = -1
                switch 2147483647 {
                case 716:
                    switch 42 {
                    case 4294967295:
                        v24 := GenS3{F0: 1}
                        _ = v24
                    case -1:
                        switch -91 {
                        case 42:
                            fmt.Println(v21, v1, v23)
                        }
                    default:
                        if (127 != 42) {
                            var v25 uint8 = 0
                            var v26 uint8 = 179
                            var v27 uint32 = 42
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    }
                case 828:
                    _ = len(v2)
                case -9223372036854775808:
                    v4[2147483648] = 1
                default:
                    v28 := v3[-1]
                    _ = v28
                }
                _ = v23
            } else {
                delete(v3, 128)
                var v29 int16 = 1
                _ = v29
            }
        } else {
            for fi30 := 0; fi30 < 8; fi30++ {
                for fi31 := 0; fi31 < 10; fi31++ {
                    var v32 int32 = -2147483648
                    _ = fi31
                    _ = v32
                }
                v33 := GenS2{F0: -128, F1: v1}
                _ = fi30
                _ = v33
            }
            v34 := GenS1{F0: 127, F1: 0}
            _ = v34
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 = (v1 + (v1 + v1))
    }
    var v35 string = "hello"
    v1 = "\n\t"
    switch v1 {
    case "":
        for fi36 := 0; fi36 < 13; fi36++ {
            var v37 int = (fi36 + 286)
            switch v37 {
            case 32767:
                var v38 uint8 = 219
                _ = v38
            default:
                v2 = append(v2, v37)
            }
            _ = fi36
            _ = v37
        }
    case "a":
        var v39 int32 = 1
        _ = v39
    case "\n\t":
        if (42 > 0) {
            fmt.Println(v35)
        }
    default:
        if false {
            v35 = v1
            switch v1 {
            case "":
                if false {
                    for fi40 := 0; fi40 < 2; fi40++ {
                        if true {
                            if false {
                                var v41 uint = v21
                                for fi42 := 0; fi42 < 17; fi42++ {
                                    fmt.Println(v1)
                                    switch ("hello" + "foo bar") {
                                    case "hello":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v43 uint8 = 227
                                        v1 = v35
                                        fmt.Println(v21, v43, v41)
                                        _ = v43
                                    }
                                    _ = fi42
                                }
                                _ = v41
                            } else {
                                var v44 uint8 = 60
                                var v45 int8 = 42
                                v46 := GenS3{F0: (fi40 * (^(fi40)))}
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                            for fi47 := 0; fi47 < 8; fi47++ {
                                var v48 string = v1
                                _ = fi47
                                _ = v48
                            }
                            var v49 string = "a"
                            _ = v49
                        } else {
                            v50 := GenS4{F0: "hello", F1: 558, F2: -395}
                            v51 := GenS5{F0: 42, F1: 42, F2: 1}
                            v21 += (^(v21))
                            _ = v50
                            _ = v51
                        }
                        _ = fi40
                    }
                    if false {
                        v4[0] = 4294967295
                    } else {
                        v2 = append(v2, -845)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v35 = (("" + "\n\t") + ("a" + (("" + (v35 + "a")) + v1)))
                } else {
                    fmt.Println(v35, v1, v1)
                    v35 = ("foo bar" + (v1 + (v1 + "")))
                }
            case "foo bar":
                if (0 >= 78) {
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        switch v35 {
                        case "":
                            for fi52 := 0; fi52 < 4; fi52++ {
                                var v53 float64 = (3.14159 / (3.14159 + float64(3.14159)))
                                v2 = append(v2, (fi52 << (((^(v21))) & 63)))
                                if len(v2) > 0 {
                                    v2[0] = (fi52 << ((v21) & 63))
                                }
                                _ = fi52
                                _ = v53
                            }
                        case "hello":
                            fmt.Println(v1)
                        default:
                            switch v35 {
                            case "":
                                v54 := GenS5{F0: 292, F1: 1, F2: 1}
                                _ = v54
                            case "a":
                                switch v1 {
                                case "a":
                                    if (100 <= -9) {
                                        v1 = ("foo bar" + (((v35 + "hello") + (v1 + "\n\t")) + v1))
                                        var v55 uint64 = 1
                                        fmt.Println(v55, v35)
                                        _ = v55
                                    } else {
                                        fmt.Println(v21, v35, v21)
                                    }
                                default:
                                    for fi56 := 0; fi56 < 12; fi56++ {
                                        if (!true) {
                                            fmt.Println(v21, v1, v1)
                                            if ((v21 << ((388) & 63)) >= v21) {
                                                if (!(2 <= 110)) {
                                                    v57 := GenS3{F0: (fi56 ^ fi56)}
                                                    _ = v57
                                                } else {
                                                    v58 := GenS1{F0: 87, F1: 158}
                                                    switch v35 {
                                                    case "foo bar":
                                                        for fi59 := 0; fi59 < 3; fi59++ {
                                                            if (!(false && false)) {
                                                                var v60 int32 = 639
                                                                _ = v60
                                                            }
                                                            v61 := GenS6{F0: 18446744073709551615, F1: false}
                                                            _ = fi59
                                                            _ = v61
                                                        }
                                                    case "":
                                                        if (!true) {
                                                            var v62 uint64 = 244
                                                            _ = len(v3)
                                                            v35 = "\n\t"
                                                            _ = v62
                                                        } else {
                                                            var v63 bool = (1 > 231)
                                                            if v63 {
                                                                for fi64 := 0; fi64 < 15; fi64++ {
                                                                    v63 = (!false)
                                                                    v35 = (("foo bar" + v1) + v35)
                                                                    var v65 int = (fi64 % (((-(fi56))) | 1))
                                                                    if (32767 > 32767) {
                                                                        fmt.Println(v35, v65)
                                                                    } else {
                                                                        var v66 string = (("" + "") + (v35 + ""))
                                                                        if v63 {
                                                                            _ = len(v2)
                                                                            if (379 <= 1) {
                                                                                v65 = (v65 ^ (fi64 % (((-(v65))) | 1)))
                                                                                v67 := GenS1{F0: 127, F1: 18446744073709551615}
                                                                                if (!false) {
                                                                                    if (!(-92 != 100)) {
                                                                                        v1 = "hello"
                                                                                        var v68 int8 = 0
                                                                                        _ = v68
                                                                                    } else {
                                                                                        v69 := v4[1]
                                                                                        switch (v65 >> (((v21 | (^(v21)))) & 63)) {
                                                                                        case 4294967295:
                                                                                            if v63 {
                                                                                                var v70 int32 = 2147483647
                                                                                                switch v66 {
                                                                                                case "":
                                                                                                    var v71 float32 = 0.0
                                                                                                    var v72 int64 = 9223372036854775807
                                                                                                    if (("a" + v1) != ("hello" + (v35 + v1))) {
                                                                                                        v65 = 0
                                                                                                        v4[452] = v69
                                                                                                    } else {
                                                                                                        v1 = ((v1 + "hello") + (("foo bar" + "a") + ("foo bar" + v1)))
                                                                                                        if (42 == 0) {
                                                                                                            switch v35 {
                                                                                                            case "hello":
                                                                                                                var v73 uint = v21
                                                                                                                _ = v73
                                                                                                            }
                                                                                                            v63 = (!(!(42 < 127)))
                                                                                                            v74 := GenS7{F0: 1}
                                                                                                            _ = v74
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                if (v72 > (v72 + (v72 ^ (-(v72))))) {
                                                                                                                    switch v35 {
                                                                                                                    case "hello":
                                                                                                                        v70 -= -1
                                                                                                                    case "a":
                                                                                                                        var v75 uint64 = 927
                                                                                                                        _ = v75
                                                                                                                    }
                                                                                                                    v3[(^(fi56))] = (fi64 & 42)
                                                                                                                    _ = len(v4)
                                                                                                                }
                                                                                                            }
                                                                                                            var v76 float64 = 0.0
                                                                                                            v72 = v72
                                                                                                            _ = v76
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v71
                                                                                                    _ = v72
                                                                                                case "foo bar":
                                                                                                    v35 = "\n\t"
                                                                                                }
                                                                                                switch v1 {
                                                                                                case "a":
                                                                                                    v3[-1] = v65
                                                                                                }
                                                                                                _ = v70
                                                                                            }
                                                                                        case 9223372036854775807:
                                                                                            switch v35 {
                                                                                            case "\n\t":
                                                                                                switch int(v21) {
                                                                                                case -2147483648:
                                                                                                    var v77 bool = ((!(true || (!(-128 == 1)))) || (-2.25 == ((-2.25 * 0.1) * (3.14159 / 3.14159))))
                                                                                                    _ = v77
                                                                                                case 32768:
                                                                                                    v21 -= 255
                                                                                                default:
                                                                                                    v65 -= fi64
                                                                                                }
                                                                                            default:
                                                                                                v78 := GenS3{F0: v65}
                                                                                                _ = v78
                                                                                            }
                                                                                        }
                                                                                        _ = v69
                                                                                    }
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = fi64
                                                                                    }
                                                                                } else {
                                                                                    var v79 float64 = -2.25
                                                                                    _ = v79
                                                                                }
                                                                                _ = v67
                                                                            }
                                                                        } else {
                                                                            var v80 uint32 = 2
                                                                            if (float64((((-2.25 + 3.14159) / -2.25) / 1.5)) > 0.1) {
                                                                                if v63 {
                                                                                    if (fi64 >= fi56) {
                                                                                        if (810 != 42) {
                                                                                            switch -9223372036854775808 {
                                                                                            case -330:
                                                                                                v81 := v3[(fi56 >> ((v21) & 63))]
                                                                                                switch -2147483648 {
                                                                                                case 160:
                                                                                                    v80 = (v80 ^ v80)
                                                                                                case -1:
                                                                                                    v1 = "a"
                                                                                                case -9223372036854775808:
                                                                                                    if (1000 == 42) {
                                                                                                        var v82 string = ("a" + "foo bar")
                                                                                                        v80 -= (v80 * v80)
                                                                                                        fmt.Println(v81, v80, fi64)
                                                                                                        _ = v82
                                                                                                    }
                                                                                                default:
                                                                                                    v83 := GenS8{F0: v21, F1: 18446744073709551615, F2: (!("hello" != ((((v35 + "") + "a") + v1) + (v35 + "\n\t"))))}
                                                                                                    _ = v83
                                                                                                }
                                                                                                var v84 uint32 = (^(v80))
                                                                                                _ = v81
                                                                                                _ = v84
                                                                                            default:
                                                                                                v66 = (("foo bar" + ((((("foo bar" + "") + v1) + (v35 + ("" + ((v66 + "a") + "hello")))) + v66) + (v35 + v1))) + "foo bar")
                                                                                                v35 = v66
                                                                                                var v85 uint = (v21 | (v21 + 2))
                                                                                                var v86 uint32 = uint32(v21)
                                                                                                _ = v85
                                                                                                _ = v86
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v87 := GenS3{F0: 128}
                                                                                    var v88 uint8 = 2
                                                                                    _ = v87
                                                                                    _ = v88
                                                                                }
                                                                                if (true && (false || (v66 < "a"))) {
                                                                                    var v89 int8 = -93
                                                                                    switch v1 {
                                                                                    case "\n\t":
                                                                                        v80 -= v80
                                                                                    case "a":
                                                                                        fmt.Println(fi64, v21)
                                                                                        var v90 float32 = float32(0.1)
                                                                                        v91 := v3[128]
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    case "foo bar":
                                                                                        v35 = (v35 + v35)
                                                                                    }
                                                                                    fmt.Println(v35)
                                                                                    _ = v89
                                                                                }
                                                                                var v92 int64 = int64(v80)
                                                                                _ = v92
                                                                            }
                                                                            _ = v80
                                                                        }
                                                                        _ = v66
                                                                    }
                                                                    v63 = true
                                                                    var v93 uint16 = 255
                                                                    _ = fi64
                                                                    _ = v65
                                                                    _ = v93
                                                                }
                                                                v1 = v35
                                                                for fi94 := 0; fi94 < 18; fi94++ {
                                                                    v21 += 833
                                                                    _ = fi94
                                                                }
                                                            }
                                                            _ = v63
                                                        }
                                                    default:
                                                        if (704 <= 96) {
                                                            fmt.Println(v35)
                                                        }
                                                    }
                                                    _ = v58
                                                }
                                            }
                                            v21 = 1
                                        } else {
                                            v35 = (v35 + v35)
                                            var v95 int8 = 42
                                            _ = v95
                                        }
                                        v35 = v35
                                        _ = fi56
                                    }
                                }
                            default:
                                fmt.Println(v21, v1, v1)
                            }
                        }
                    case 4294967295:
                        var v96 int8 = -16
                        _ = v96
                    case 128:
                        var v97 int = 779
                        _ = v97
                    default:
                        for fi98 := 0; fi98 < 2; fi98++ {
                            switch (fi98 / (((fi98 >> ((255) & 63))) | 1)) {
                            case -87:
                                v35 = "\n\t"
                            case -467:
                                if true {
                                    for fi99 := 0; fi99 < 4; fi99++ {
                                        if true {
                                            v1 = v35
                                            switch (v1 + ((v1 + v1) + v35)) {
                                            case "\n\t":
                                                v35 = ""
                                            default:
                                                if true {
                                                    v1 = (v1 + "foo bar")
                                                    fmt.Println(fi98, fi99)
                                                    _ = len(v4)
                                                } else {
                                                    v21 = (^(v21))
                                                    var v100 uint = 1
                                                    _ = v100
                                                }
                                            }
                                        }
                                        if ((true || true) && (-20 <= 127)) {
                                            v35 = v1
                                            fmt.Println(v1, fi99, v1)
                                        }
                                        _ = fi99
                                    }
                                }
                            default:
                                v4[65535] = 985
                            }
                            fmt.Println(v35, fi98, v1)
                            _ = fi98
                        }
                    }
                }
            case "a":
                if (!false) {
                    var v101 uint8 = 127
                    _ = v101
                } else {
                    for fi102 := 0; fi102 < 11; fi102++ {
                        v103 := GenS2{F0: -128, F1: ((v1 + "\n\t") + "hello")}
                        var v104 uint8 = 2
                        v105 := GenS8{F0: v21, F1: 2, F2: false}
                        _ = fi102
                        _ = v103
                        _ = v104
                        _ = v105
                    }
                }
            }
        } else {
            v1 = (v35 + (v35 + "hello"))
            var v106 float32 = ((-2.25 * (0.0 + (1.5 + 1.5))) - 1.5)
            v1 = "a"
            _ = v106
        }
    }
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v4
}
