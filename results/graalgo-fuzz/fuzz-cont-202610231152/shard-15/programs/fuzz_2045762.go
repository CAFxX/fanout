package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8; F1 uint8 }
type GenS11 struct { F0 string; F1 int }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint64; F1 string }
type GenS14 struct { F0 uint8; F1 int64; F2 uint32 }
func genFn0(p0 uint8) int32 {
    v7 := []int{375, 0}
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: 18446744073709551615, F1: 42}
        _ = fi8
        _ = v9
    }
    p0 = p0
    switch "" {
    case "a":
        p0 -= p0
    case "":
        v10 := GenS1{F0: 4294967295, F1: 1}
        _ = v10
    case "\n\t":
        if len(v7) > 0 {
            v7[0] = 127
        }
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1() int64 {
    var v11 uint = 668
    var v12 uint = (^(v11))
    if false {
        v13 := GenS2{F0: 0, F1: 71, F2: false}
        var v14 int = 32768
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS3{F0: (^(v11))}
            v11 = v12
            _ = fi15
            _ = v16
        }
    }
    _ = v11
    _ = v12
    return -9223372036854775808
}

func main() {
    var v1 int64 = 716
    var v2 string = ("foo bar" + "foo bar")
    v3 := []int{-1, 32768, 668}
    var v4 uint16 = 759
    var v5 int64 = (^(v1))
    var v6 string = ""
    v3 = append(v3, 1)
    for fi17 := 0; fi17 < 13; fi17++ {
        v5 = 42
        _ = fi17
    }
    var v18 uint = uint(v4)
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 uint8 = 2
        _ = fi19
        _ = v20
    }
    v4 -= 2
    v2 = (v6 + v6)
    var v21 bool = false
    if (65535 == 2147483648) {
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 uint8 = 41
            v18 -= 18446744073709551615
            for fi24 := 0; fi24 < 20; fi24++ {
                v5 = genFn1()
                var v25 uint16 = v4
                v6 = "hello"
                switch "" {
                case "\n\t":
                    v21 = false
                case "":
                    var v26 int16 = 0
                    _ = v26
                default:
                    var v27 string = "foo bar"
                    _ = v27
                }
                var v28 float32 = float32((1.5 * (float64(float32(1e10)) + 0.1)))
                _ = fi24
                _ = v25
                _ = v28
            }
            _ = fi22
            _ = v23
        }
        v2 = ((v6 + "a") + ("foo bar" + v2))
        switch 131 {
        case 42:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -701:
            var v29 int8 = -11
            _ = v29
        case 2147483647:
            for fi30 := 0; fi30 < 20; fi30++ {
                v5 = (v5 % (((-(v1))) | 1))
                if (0 < 837) {
                    var v31 uint8 = 127
                    if v21 {
                        v32 := GenS3{F0: v18}
                        v3 = append(v3, (fi30 | fi30))
                        _ = v32
                    }
                    _ = v31
                } else {
                    if false {
                        switch 32768 {
                        case 2147483647:
                            v33 := GenS4{F0: 42, F1: 2, F2: 104}
                            _ = v33
                        case 9223372036854775807:
                            for fi34 := 0; fi34 < 16; fi34++ {
                                v35 := GenS5{F0: "hello", F1: 0}
                                _ = fi34
                                _ = v35
                            }
                        default:
                            v21 = (!(v18 != (v18 & v18)))
                        }
                    }
                    switch "a" {
                    case "\n\t":
                        if len(v3) > 0 {
                            v3[0] = (fi30 + 1)
                        }
                    default:
                        switch ((((v6 + v2) + v2) + ((v2 + v6) + v2)) + "a") {
                        case "hello":
                            var v36 uint32 = 65535
                            _ = v36
                        case "":
                            if v21 {
                                var v37 uint = 42
                                var v38 uint32 = 65535
                                _ = len(v3)
                                _ = v37
                                _ = v38
                            }
                        default:
                            var v39 uint32 = 65535
                            _ = v39
                        }
                    }
                    var v40 uint = 868
                    _ = v40
                }
                v5 -= (v5 | 304)
                _ = fi30
            }
        default:
            switch "hello" {
            case "foo bar":
                v3 = append(v3, 32767)
            case "hello":
                var v41 uint = 489
                _ = v41
            case "":
                var v42 int16 = 1000
                _ = v42
            }
        }
    } else {
        for fi43 := 0; fi43 < 8; fi43++ {
            if v21 {
                switch (v6 + v2) {
                case "foo bar":
                    var v44 bool = v21
                    _ = v44
                case "":
                    for fi45 := 0; fi45 < 14; fi45++ {
                        var v46 uint64 = 345
                        var v47 float32 = ((float32(1e10) - 3.14159) + float32((1.5 * 0.0)))
                        if true {
                            fmt.Println(v21, v46)
                        }
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                default:
                    if (-1 != -35) {
                        fmt.Println(v4, v4)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        v21 = ((1e10 + 1e10) != (1e10 + -2.25))
                    }
                }
            } else {
                v6 = v2
                for fi48 := 0; fi48 < 16; fi48++ {
                    var v49 uint8 = 47
                    var v50 uint8 = 184
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                switch v2 {
                case "":
                    if v21 {
                        for fi51 := 0; fi51 < 8; fi51++ {
                            if (!(v21 && true)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                _ = len(v3)
                            }
                            var v52 uint64 = 1
                            var v53 uint16 = v4
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                        var v54 float32 = 1.5
                        _ = v54
                    }
                case "\n\t":
                    for fi55 := 0; fi55 < 9; fi55++ {
                        v18 = (v18 | v18)
                        v56 := GenS5{F0: ("\n\t" + ((v6 + "a") + ("a" + (v6 + ("hello" + v6))))), F1: fi43}
                        _ = fi55
                        _ = v56
                    }
                }
            }
            for fi57 := 0; fi57 < 12; fi57++ {
                switch fi57 {
                case -2147483648:
                    if (!(v21 && v21)) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        v21 = v21
                    }
                case -860:
                    v58 := GenS5{F0: ("foo bar" + ""), F1: fi57}
                    _ = v58
                }
                _ = fi57
            }
            v6 = ""
            _ = fi43
        }
    }
    var v59 int8 = int8(v1)
    v60 := GenS3{F0: 281}
    for fi61 := 0; fi61 < 12; fi61++ {
        var v62 float32 = (-2.25 * -2.25)
        var v63 uint32 = 366
        _ = fi61
        _ = v62
        _ = v63
    }
    if (("a" + v6) <= ("a" + v6)) {
        if v21 {
            for fi64 := 0; fi64 < 9; fi64++ {
                v65 := GenS3{F0: v18}
                v66 := GenS6{F0: 0, F1: 0}
                if true {
                    if v21 {
                        var v67 int16 = 478
                        v4 -= (^(v4))
                        _ = v67
                    } else {
                        if (65535 < v18) {
                            v68 := GenS6{F0: 65535, F1: 0}
                            _ = v68
                        }
                        var v69 bool = false
                        for fi70 := 0; fi70 < 18; fi70++ {
                            if (1 != 138) {
                                v1 -= v1
                                switch (fi70 ^ -1) {
                                case 127:
                                    var v71 string = v6
                                    fmt.Println(v71)
                                    if false {
                                        v72 := GenS7{F0: (0 >= 2), F1: (v2 + v71)}
                                        if false {
                                            if v69 {
                                                v2 = (v71 + v71)
                                                fmt.Println(fi70, v18)
                                                v71 = "hello"
                                                v6 = "a"
                                            } else {
                                                if (uint32(v4) > 2147483648) {
                                                    var v73 uint64 = 4294967295
                                                    if (v6 >= "a") {
                                                        v21 = true
                                                        v74 := GenS3{F0: (^(v18))}
                                                        var v75 uint16 = v4
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    _ = v73
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = -9223372036854775808
                                                    }
                                                    var v76 int64 = int64(v59)
                                                    _ = v76
                                                }
                                            }
                                            switch (fi70 >> ((v18) & 63)) {
                                            case 9223372036854775807:
                                                v18 = (v18 ^ (v18 << ((1) & 63)))
                                            case 128:
                                                if false {
                                                    v21 = (v4 >= 167)
                                                    if (644 > 9) {
                                                        v69 = (1 != -406)
                                                        fmt.Println(v69)
                                                        fmt.Println(v69, v6, fi70)
                                                    } else {
                                                        v77 := GenS8{F0: true, F1: (fi64 >> ((v18) & 63))}
                                                        fmt.Println(v4)
                                                        _ = v77
                                                    }
                                                } else {
                                                    if v69 {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    v78 := GenS9{F0: 880}
                                                    _ = len(v3)
                                                    _ = v78
                                                }
                                                v79 := GenS8{F0: (v69 || (1 >= 18446744073709551615)), F1: (fi70 & fi70)}
                                                fmt.Println(v21, fi64)
                                                _ = v79
                                            case -1:
                                                fmt.Println(v18, fi70)
                                            }
                                        } else {
                                            fmt.Println(v69)
                                            var v80 float64 = 1e10
                                            switch -129 {
                                            case 4294967295:
                                                v71 = v71
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v81 uint16 = v4
                                            v82 := GenS10{F0: (v59 ^ (v59 << ((v18) & 63))), F1: 128}
                                            _ = v80
                                            _ = v81
                                            _ = v82
                                        }
                                        v83 := GenS11{F0: v6, F1: 0}
                                        _ = v72
                                        _ = v83
                                    }
                                    _ = v71
                                case -913:
                                    v84 := GenS6{F0: 0, F1: -361}
                                    _ = v84
                                }
                                var v85 int16 = 1000
                                _ = v85
                            }
                            v86 := GenS12{F0: v4}
                            _ = fi70
                            _ = v86
                        }
                        _ = v69
                    }
                }
                _ = fi64
                _ = v65
                _ = v66
            }
            var v87 uint32 = 135
            var v88 uint32 = (v87 - (^(v87)))
            _ = v87
            _ = v88
        }
        for fi89 := 0; fi89 < 20; fi89++ {
            var v90 bool = (v21 || v21)
            _ = fi89
            _ = v90
        }
    }
    if v21 {
        if (v21 && false) {
            var v91 uint16 = v4
            v92 := GenS5{F0: "hello", F1: 42}
            fmt.Println(v18, v21, v2)
            _ = v91
            _ = v92
        }
    }
    if (765 >= 796) {
        var v93 int64 = 0
        _ = v93
    } else {
        var v94 string = ("foo bar" + (v6 + "a"))
        _ = v94
    }
    if ((1e10 > -2.25) && true) {
        v18 -= (v18 >> ((v18) & 63))
    } else {
        fmt.Println(v21, v59)
        v2 = "\n\t"
        switch -287 {
        case 128:
            v6 = ""
        case -664:
            _ = len(v3)
        case 0:
            v18 = (^(v18))
        }
    }
    if v21 {
        v95 := GenS12{F0: (v4 << ((v18) & 63))}
        v1 = (v5 * genFn1())
        var v96 string = (("a" + "\n\t") + ((v2 + v6) + v2))
        _ = v95
        _ = v96
    } else {
        v97 := GenS13{F0: 42, F1: "\n\t"}
        fmt.Println(v59)
        v98 := GenS14{F0: 183, F1: v5, F2: 236}
        _ = v97
        _ = v98
    }
    for fi99 := 0; fi99 < 8; fi99++ {
        if len(v3) > 0 {
            v3[0] = -567
        }
        v1 = int64(v18)
        for fi100 := 0; fi100 < 14; fi100++ {
            v1 += 9223372036854775807
            _ = fi100
        }
        _ = fi99
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
}
