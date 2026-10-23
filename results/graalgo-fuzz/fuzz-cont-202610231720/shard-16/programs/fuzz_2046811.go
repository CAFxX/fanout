package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 uint; F1 int; F2 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int; F1 uint; F2 uint }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int; F1 uint32 }
type GenS14 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint, p1 int) uint8 {
    var v5 float64 = -2.25
    fmt.Println(v5)
    var v6 bool = false
    for fi7 := 0; fi7 < 14; fi7++ {
        if true {
            v8 := GenS1{F0: -1, F1: fi7}
            switch ("hello" + "") {
            case "hello":
                v9 := GenS2{F0: -1}
                _ = v9
            }
            if (!(1.5 == -2.25)) {
                v6 = false
            }
            _ = v8
        } else {
            var v10 int64 = -273
            if false {
                var v11 uint32 = 2147483647
                if (123 < 2) {
                    switch "a" {
                    case "hello":
                        var v12 uint8 = 199
                        _ = v12
                    }
                    v13 := GenS3{F0: -1}
                    for fi14 := 0; fi14 < 17; fi14++ {
                        v10 = v10
                        v15 := GenS3{F0: 0}
                        _ = fi14
                        _ = v15
                    }
                    _ = v13
                }
                _ = v11
            }
            if (!(!((2 < 0) || v6))) {
                v6 = v6
                for fi16 := 0; fi16 < 19; fi16++ {
                    v17 := GenS4{F0: 2, F1: (v10 == v10)}
                    var v18 int32 = -1
                    var v19 int32 = -247
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
            } else {
                fmt.Println(v6, p1, p1)
                var v20 uint64 = 4294967295
                if v6 {
                    for fi21 := 0; fi21 < 2; fi21++ {
                        var v22 float64 = 1.5
                        var v23 int32 = 2147483647
                        var v24 uint = 980
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS5{F0: (("" + "foo bar") + "\n\t")}
                    switch "" {
                    case "foo bar":
                        p1 = -65
                    case "":
                        if ((!(fi7 > fi7)) && v6) {
                            for fi26 := 0; fi26 < 15; fi26++ {
                                var v27 int64 = v10
                                fmt.Println(v10, v20, p0)
                                var v28 int8 = 42
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            var v29 float64 = ((1e10 / v5) / float64(0.1))
                            for fi30 := 0; fi30 < 18; fi30++ {
                                if false {
                                    if (!(v6 && true)) {
                                        v10 -= v10
                                        var v31 int32 = 1
                                        _ = v31
                                    }
                                } else {
                                    fmt.Println(fi7, v6, fi7)
                                }
                                v10 -= v10
                                if v6 {
                                    if (-1 < 42) {
                                        if (790 != 952) {
                                            var v32 string = "foo bar"
                                            _ = v32
                                        }
                                        if v6 {
                                            v6 = false
                                        }
                                        v10 += (v10 & v10)
                                    }
                                }
                                _ = fi30
                            }
                            _ = v29
                        } else {
                            var v33 int64 = int64(v20)
                            _ = v33
                        }
                    }
                    _ = v25
                }
                _ = v20
            }
            _ = v10
        }
        fmt.Println(v6)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2
}
func genFn1() uint64 {
    v34 := []int{-9223372036854775808, 588, -9223372036854775808, 128}
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 bool = (!false)
        var v37 uint = 580
        for fi38 := 0; fi38 < 13; fi38++ {
            v39 := GenS6{F0: -1}
            _ = fi38
            _ = v39
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    if ((-2.25 * ((0.0 + float32(1.5)) + -2.25)) <= 1.5) {
        v40 := GenS7{F0: -201, F1: 510}
        v41 := GenS8{F0: 1, F1: 2147483647, F2: 6}
        _ = v40
        _ = v41
    }
    _ = v34
    return 491
}

func main() {
    var v1 int32 = -53
    var v2 float64 = (3.14159 + (3.14159 * 1e10))
    var v3 int = -198
    var v4 float32 = 3.14159
    var v42 uint8 = 189
    for fi43 := 0; fi43 < 10; fi43++ {
        switch v3 {
        case 2147483647:
            switch "foo bar" {
            case "":
                v3 -= (fi43 & (fi43 % (((v3 | v3)) | 1)))
            }
        case 9223372036854775807:
            var v44 int32 = (v1 & (v1 >> ((670) & 63)))
            _ = v44
        case 128:
            for fi45 := 0; fi45 < 20; fi45++ {
                v3 += fi43
                v46 := []int{128, v3, fi43}
                _ = fi45
                _ = v46
            }
        default:
            v47 := GenS1{F0: v1, F1: (-(v3))}
            _ = v47
        }
        v2 = 1.5
        fmt.Println(v3, v1)
        _ = fi43
    }
    v42 = v42
    for fi48 := 0; fi48 < 1; fi48++ {
        fmt.Println(v3, v3)
        for fi49 := 0; fi49 < 11; fi49++ {
            v3 += (v3 << ((18446744073709551615) & 63))
            _ = fi49
        }
        _ = fi48
    }
    switch 4294967295 {
    case -508:
        fmt.Println(v1, v3)
    case -1:
        for fi50 := 0; fi50 < 12; fi50++ {
            v51 := []int{}
            var v52 uint32 = 65535
            _ = fi50
            _ = v51
            _ = v52
        }
    default:
        for fi53 := 0; fi53 < 10; fi53++ {
            if ((true || true) || (!true)) {
                v1 += -465
                v54 := []int{}
                _ = v54
            }
            for fi55 := 0; fi55 < 3; fi55++ {
                v3 += (v3 ^ fi55)
                if (true || (1 >= 2147483648)) {
                    var v56 float32 = (((v4 + v4) * float32(v2)) + -2.25)
                    _ = v56
                }
                _ = fi55
            }
            v57 := []int{(fi53 >> ((65535) & 63)), 128, 17}
            _ = fi53
            _ = v57
        }
    }
    if true {
        v1 = (v1 >> ((1) & 63))
        v42 += 187
    }
    v4 = v4
    switch v3 {
    case 9223372036854775807:
        switch ("\n\t" + ((("\n\t" + ("hello" + ("foo bar" + ("foo bar" + ("hello" + ("a" + ("" + "foo bar"))))))) + ("foo bar" + "foo bar")) + "\n\t")) {
        case "a":
            for fi58 := 0; fi58 < 1; fi58++ {
                var v59 float64 = (-2.25 * (0.1 / ((((v2 + (1.5 / -2.25)) / 1.0) - (v2 / 1e10)) - v2)))
                var v60 uint8 = (v42 * (^(v42)))
                _ = fi58
                _ = v59
                _ = v60
            }
        }
    case 127:
        for fi61 := 0; fi61 < 11; fi61++ {
            for fi62 := 0; fi62 < 13; fi62++ {
                v63 := GenS9{F0: ("a" != "a")}
                if true {
                    var v64 uint64 = 130
                    fmt.Println(fi62)
                    var v65 int64 = -1
                    _ = v64
                    _ = v65
                }
                var v66 string = (("hello" + ("\n\t" + "")) + "foo bar")
                _ = fi62
                _ = v63
                _ = v66
            }
            v67 := GenS10{F0: int(v42), F1: 15, F2: 42}
            if (42 <= 0) {
                if false {
                    v3 += (fi61 << ((42) & 63))
                    switch "" {
                    case "\n\t":
                        var v68 float64 = v2
                        _ = v68
                    case "hello":
                        fmt.Println(fi61)
                    }
                } else {
                    if (!true) {
                        v1 += (v1 * v1)
                        var v69 float32 = v4
                        _ = v69
                    }
                    for fi70 := 0; fi70 < 13; fi70++ {
                        v42 += (v42 << ((uint(v42)) & 63))
                        _ = fi70
                    }
                    for fi71 := 0; fi71 < 9; fi71++ {
                        var v72 int8 = 17
                        _ = fi71
                        _ = v72
                    }
                }
            } else {
                var v73 float64 = (0.0 + ((0.0 + float64((0.1 - 3.14159))) / (-2.25 - v2)))
                _ = v73
            }
            _ = fi61
            _ = v67
        }
    case -562:
        switch v3 {
        case 2147483647:
            v74 := []int{v3, (^(v3)), (^(v3))}
            _ = v74
        case -397:
            switch v3 {
            case -27:
                v75 := []int{v3}
                _ = v75
            case -2147483648:
                var v76 int32 = v1
                _ = v76
            case 968:
                for fi77 := 0; fi77 < 11; fi77++ {
                    v78 := GenS5{F0: "a"}
                    switch (fi77 % ((v3) | 1)) {
                    case -759:
                        v79 := GenS11{F0: (fi77 / (((-(fi77))) | 1))}
                        _ = v79
                    }
                    _ = fi77
                    _ = v78
                }
            default:
                if true {
                    v80 := GenS12{F0: 4294967295}
                    _ = v80
                }
            }
        case 5:
            v2 = v2
        }
    default:
        v81 := GenS13{F0: (v3 + (v3 ^ v3)), F1: 0}
        _ = v81
    }
    if true {
        var v82 float64 = float64(v4)
        if false {
            v83 := []int{127, v3, (^(v3)), 32768}
            fmt.Println(v4, v82)
            v84 := GenS14{F0: (-2.25 < v82), F1: 17}
            _ = v83
            _ = v84
        } else {
            for fi85 := 0; fi85 < 15; fi85++ {
                var v86 float32 = (1e10 * v4)
                _ = fi85
                _ = v86
            }
        }
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
