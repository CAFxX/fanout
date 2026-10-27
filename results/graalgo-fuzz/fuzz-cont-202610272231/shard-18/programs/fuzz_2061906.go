package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 uint32; F2 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32; F1 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int; F1 int32 }
type GenS10 struct { F0 uint; F1 int64; F2 uint }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint16, p1 uint32) int32 {
    fmt.Println(p1)
    switch ("\n\t" + ("foo bar" + "hello")) {
    case "foo bar":
        for fi4 := 0; fi4 < 18; fi4++ {
            v5 := []int{}
            switch fi4 {
            case 4294967295:
                if (42 >= 125) {
                    switch ((("hello" + "a") + ((("\n\t" + "") + ("\n\t" + "")) + ("\n\t" + "foo bar"))) + "a") {
                    case "hello":
                        for fi6 := 0; fi6 < 8; fi6++ {
                            var v7 uint8 = 2
                            var v8 int16 = 0
                            _ = fi6
                            _ = v7
                            _ = v8
                        }
                    }
                    _ = len(v5)
                    for fi9 := 0; fi9 < 2; fi9++ {
                        v10 := GenS1{F0: 42, F1: -487, F2: "foo bar"}
                        _ = fi9
                        _ = v10
                    }
                }
            case 379:
                var v11 int64 = 887
                _ = v11
            case -9223372036854775808:
                var v12 bool = true
                _ = v12
            default:
                v13 := GenS1{F0: 0, F1: 0, F2: (("" + "hello") + (("foo bar" + "\n\t") + "hello"))}
                _ = v13
            }
            v14 := GenS1{F0: 127, F1: -2147483648, F2: "a"}
            _ = fi4
            _ = v14
            _ = v5
        }
    case "":
        if false {
            for fi15 := 0; fi15 < 13; fi15++ {
                v16 := GenS2{F0: 204}
                var v17 int = (-(fi15))
                _ = fi15
                _ = v16
                _ = v17
            }
            switch "\n\t" {
            case "a":
                for fi18 := 0; fi18 < 6; fi18++ {
                    v19 := []int{910, (fi18 % (((fi18 / ((fi18) | 1))) | 1)), fi18}
                    fmt.Println(p1)
                    for fi20 := 0; fi20 < 9; fi20++ {
                        v21 := GenS3{F0: (p0 ^ 0)}
                        _ = fi20
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                }
            }
        }
    default:
        fmt.Println(p1, p1)
    }
    var v22 uint = 255
    var v23 int8 = 127
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = float64(1e10)
    v3 := []int{505, 32767, 4294967295}
    var v24 bool = true
    v24 = (0 != 1)
    if true {
        if (true && (0 >= 1)) {
            fmt.Println(v2)
        }
        var v25 int = 2147483647
        var v26 string = ("hello" + "hello")
        _ = v25
        _ = v26
    } else {
        var v27 uint64 = 676
        v2 -= (v2 / v2)
        if ((true && (!((v2 != (-2.25 - float64(0.0))) && v24))) || v24) {
            switch ("\n\t" + (("foo bar" + "hello") + "hello")) {
            case "foo bar":
                var v28 int16 = 0
                _ = v28
            default:
                var v29 uint32 = 127
                _ = v29
            }
            v2 = 1e10
        } else {
            v30 := GenS1{F0: 127, F1: genFn0(42, 783), F2: (("hello" + "foo bar") + "a")}
            for fi31 := 0; fi31 < 8; fi31++ {
                for fi32 := 0; fi32 < 17; fi32++ {
                    var v33 uint32 = 2
                    v34 := GenS4{F0: 255, F1: uint32(v27), F2: ("hello" + "foo bar")}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                v35 := GenS5{F0: 2147483648}
                v2 = v2
                _ = fi31
                _ = v35
            }
            if (!v24) {
                var v36 float32 = -2.25
                delete(v1, "\n\t")
                switch 9223372036854775807 {
                case 9223372036854775807:
                    var v37 float64 = 0.0
                    _ = v37
                case 127:
                    if (42 == 561) {
                        fmt.Println(v2, v27)
                        v38 := GenS6{F0: 153}
                        _ = v38
                    } else {
                        if (!false) {
                            v39 := GenS7{F0: -426, F1: 9223372036854775807}
                            fmt.Println(v27, v24)
                            _ = v39
                        }
                        v24 = v24
                    }
                case -908:
                    switch 42 {
                    case 128:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -2147483648:
                        var v40 int8 = -1
                        _ = v40
                    case 0:
                        v2 = ((-2.25 + v2) + (v2 / -2.25))
                    }
                default:
                    _ = len(v3)
                }
                _ = v36
            } else {
                v24 = (!(1.5 == (((v2 + (v2 / (v2 * v2))) + ((v2 * float64(0.1)) - v2)) / float64(3.14159))))
            }
            _ = v30
        }
        _ = v27
    }
    for fi41 := 0; fi41 < 7; fi41++ {
        for fi42 := 0; fi42 < 1; fi42++ {
            v24 = v24
            fmt.Println(v2, fi41, v2)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi42
        }
        delete(v1, "")
        switch ("foo bar" + (("\n\t" + ("a" + "hello")) + ("foo bar" + "\n\t"))) {
        case "":
            fmt.Println(v24, fi41)
        default:
            var v43 string = "foo bar"
            _ = v43
        }
        _ = fi41
    }
    if false {
        v44 := GenS8{F0: 0}
        v2 -= (-2.25 - (((float64(3.14159) - v2) * ((v2 + v2) + 0.0)) + v2))
        for fi45 := 0; fi45 < 4; fi45++ {
            if ((true || (!(v24 && ((float32(v2) / 1.0) < (((3.14159 - 0.0) / 1.0) / 3.14159))))) || false) {
                switch "\n\t" {
                case "a":
                    switch fi45 {
                    case 127:
                        var v46 int64 = 470
                        _ = v46
                    case -2147483648:
                        var v47 int8 = -1
                        _ = v47
                    case -367:
                        switch fi45 {
                        case 32768:
                            var v48 uint8 = 0
                            _ = v48
                        }
                    }
                default:
                    if false {
                        var v49 uint64 = 2
                        for fi50 := 0; fi50 < 8; fi50++ {
                            v51 := GenS9{F0: fi45, F1: -1}
                            _ = fi50
                            _ = v51
                        }
                        if (1.5 < 1.5) {
                            switch "" {
                            case "\n\t":
                                if len(v3) > 0 {
                                    v3[0] = 2147483647
                                }
                            }
                        } else {
                            for fi52 := 0; fi52 < 2; fi52++ {
                                var v53 int = 1
                                var v54 string = ""
                                _ = fi52
                                _ = v53
                                _ = v54
                            }
                            if v24 {
                                fmt.Println(v2, v2, v24)
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v24 = v24
                            }
                        }
                        _ = v49
                    } else {
                        var v55 uint64 = 966
                        _ = v55
                    }
                }
                switch "a" {
                case "\n\t":
                    var v56 bool = (454 <= 221)
                    _ = v56
                case "hello":
                    var v57 bool = false
                    _ = v57
                case "":
                    var v58 uint8 = 1
                    _ = v58
                default:
                    fmt.Println(fi45, fi45)
                }
            } else {
                var v59 bool = false
                _ = v59
            }
            v60 := GenS10{F0: 457, F1: -832, F2: 672}
            v2 += 1.5
            _ = fi45
            _ = v60
        }
        _ = v44
    }
    for fi61 := 0; fi61 < 17; fi61++ {
        fmt.Println(fi61, v2)
        _ = fi61
    }
    var v62 uint = 4294967295
    for fi63 := 0; fi63 < 8; fi63++ {
        if (42 < -2147483648) {
            var v64 int32 = 102
            v2 -= v2
            fmt.Println(fi63)
            _ = v64
        }
        _ = fi63
    }
    for fi65 := 0; fi65 < 11; fi65++ {
        v66 := GenS4{F0: 42, F1: 2147483648, F2: ("a" + "hello")}
        for fi67 := 0; fi67 < 19; fi67++ {
            fmt.Println(v2, fi65)
            _ = fi67
        }
        v68 := GenS11{F0: 1000}
        _ = fi65
        _ = v66
        _ = v68
    }
    var v69 uint = (^(v62))
    v69 = (v62 >> ((v62) & 63))
    if v24 {
        v69 = v69
        for fi70 := 0; fi70 < 13; fi70++ {
            for fi71 := 0; fi71 < 8; fi71++ {
                if ("" > "hello") {
                    v24 = v24
                    if v24 {
                        var v72 string = "\n\t"
                        _ = v72
                    } else {
                        delete(v1, "")
                        var v73 uint8 = 42
                        _ = v73
                    }
                    v74 := GenS6{F0: 1}
                    _ = v74
                }
                if len(v3) > 0 {
                    v3[0] = fi70
                }
                _ = fi71
            }
            fmt.Println(v2)
            switch (fi70 | 42) {
            case 21:
                v75 := GenS2{F0: 221}
                _ = v75
            case 128:
                if (v24 || ((50 == 42) && v24)) {
                    fmt.Println(v69, v24, v24)
                    var v76 uint32 = 1
                    v77 := GenS7{F0: 1, F1: 9223372036854775807}
                    _ = v76
                    _ = v77
                }
            case -9223372036854775808:
                v24 = v24
            }
            _ = fi70
        }
    }
    v78 := GenS12{F0: 1, F1: 42}
    if len(v3) > 0 {
        v3[0] = 731
    }
    v2 -= float64((-2.25 * 0.0))
    var v79 float64 = -2.25
    switch "\n\t" {
    case "hello":
        v80 := v1[("" + ("foo bar" + ("a" + "foo bar")))]
        _ = v80
    default:
        var v81 int16 = -1
        _ = v81
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v62:", v62)
    fmt.Println("v69:", v69)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v62
    _ = v69
    _ = v78
    _ = v79
}
