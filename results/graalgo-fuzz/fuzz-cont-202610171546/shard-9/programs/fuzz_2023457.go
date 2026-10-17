package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint; F1 uint64; F2 int16 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 uint32; F1 int16 }
type GenS10 struct { F0 int; F1 string; F2 string }
type GenS11 struct { F0 bool; F1 int32; F2 uint16 }
func genFn0() uint8 {
    var v4 uint = 691
    if (-795 != 1000) {
        var v5 float32 = -2.25
        _ = v5
    } else {
        fmt.Println(v4)
        fmt.Println(v4)
        v4 = v4
    }
    v6 := GenS1{F0: 0}
    if true {
        fmt.Println(v4)
    } else {
        var v7 uint64 = 2
        fmt.Println(v4, v4)
        _ = v7
    }
    _ = v4
    _ = v6
    return 127
}
func genFn1(p0 uint64, p1 int16, p2 uint8) int32 {
    for fi8 := 0; fi8 < 3; fi8++ {
        if ((fi8 >= fi8) || ((3.14159 * -2.25) <= 1.5)) {
            v9 := GenS2{F0: genFn0(), F1: ("hello" + "")}
            _ = v9
        }
        _ = fi8
    }
    var v10 uint8 = 1
    var v11 uint16 = 0
    var v12 float32 = 1e10
    v12 += (v12 * v12)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return -269
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = false
    var v3 int64 = 9223372036854775807
    var v13 bool = ((v3 * (v3 - v3)) != (^(v3)))
    if v2 {
        fmt.Println(v2)
        if true {
            var v14 uint16 = 255
            if true {
                if v2 {
                    v13 = v2
                    for fi15 := 0; fi15 < 2; fi15++ {
                        for fi16 := 0; fi16 < 9; fi16++ {
                            v3 += (-(v3))
                            v2 = (506 < 855)
                            _ = fi16
                        }
                        _ = fi15
                    }
                    v13 = (false && (100 == 100))
                } else {
                    v1["a"] = 262
                }
            }
            _ = v14
        } else {
            v17 := GenS2{F0: 151, F1: ""}
            switch -81 {
            case 32768:
                v3 -= -1
            default:
                if (!(1.5 < 3.14159)) {
                    v18 := GenS3{F0: false}
                    _ = v18
                }
            }
            v1[("foo bar" + "foo bar")] = 4294967295
            _ = v17
        }
    }
    fmt.Println(v13, v3, v2)
    v1[("foo bar" + "a")] = 128
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 uint = 1
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        _ = len(v1)
        _ = fi21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        fmt.Println(v3, fi22)
        v2 = (v13 && (v13 && v2))
        _ = fi22
    }
    if false {
        var v23 float64 = 0.1
        for fi24 := 0; fi24 < 9; fi24++ {
            for fi25 := 0; fi25 < 2; fi25++ {
                v2 = v2
                v2 = v2
                v26 := GenS4{F0: 0, F1: 881}
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        if v13 {
            var v27 bool = (!v2)
            if (v27 && (((v23 + (((v23 + ((v23 * 0.1) * float64(((float32((((((v23 + v23) * v23) * v23) - (0.0 - (v23 / 0.1))) / ((v23 / 1e10) * float64(1.5)))) * (1e10 / 1.5)) * (1.5 * 0.0))))) - 0.1) - float64(1e10))) == 1.5) || v2)) {
                v28 := v1["hello"]
                _ = v28
            }
            _ = v27
        } else {
            for fi29 := 0; fi29 < 1; fi29++ {
                v3 -= int64(fi29)
                switch int(v3) {
                case -9223372036854775808:
                    if v2 {
                        switch ("a" + "hello") {
                        case "foo bar":
                            if v13 {
                                var v30 uint64 = 42
                                var v31 uint8 = 255
                                _ = v30
                                _ = v31
                            }
                        case "a":
                            v32 := GenS3{F0: false}
                            _ = v32
                        default:
                            if v2 {
                                v23 = float64(-2.25)
                                var v33 uint16 = 255
                                for fi34 := 0; fi34 < 16; fi34++ {
                                    var v35 uint = 813
                                    fmt.Println(v13, v3, v13)
                                    _ = fi34
                                    _ = v35
                                }
                                _ = v33
                            }
                        }
                        v13 = v2
                    }
                case -2147483648:
                    for fi36 := 0; fi36 < 17; fi36++ {
                        if true {
                            var v37 int8 = -63
                            v37 = v37
                            _ = v37
                        }
                        v2 = ((((v23 + 3.14159) - v23) != -2.25) && (199 >= 2))
                        _ = fi36
                    }
                default:
                    if v13 {
                        if v2 {
                            var v38 float32 = 1.5
                            _ = v38
                        }
                        v39 := []int{-2147483648, fi29, (fi29 + fi29)}
                        _ = v39
                    } else {
                        for fi40 := 0; fi40 < 13; fi40++ {
                            v23 = v23
                            switch -270 {
                            case -9223372036854775808:
                                if v13 {
                                    var v41 int64 = v3
                                    var v42 float32 = 1e10
                                    _ = v41
                                    _ = v42
                                }
                            }
                            fmt.Println(v13, fi40, fi40)
                            v2 = (0 != 291)
                            if (79 < -128) {
                                fmt.Println(v2, v23, v23)
                                fmt.Println(fi29, v23)
                                fmt.Println(v3, v23)
                            }
                            _ = fi40
                        }
                        if v2 {
                            if false {
                                if ((fi29 ^ 128) == fi29) {
                                    fmt.Println(v3, v23, v2)
                                    _ = len(v1)
                                } else {
                                    v23 = v23
                                }
                                v23 = float64(((0.0 - 1.5) - 1.5))
                                v43 := GenS5{F0: 824, F1: 4294967295}
                                _ = v43
                            }
                            var v44 int8 = -128
                            _ = v44
                        } else {
                            switch fi29 {
                            case 4294967295:
                                if v2 {
                                    v1["foo bar"] = fi29
                                    if (!false) {
                                        switch "" {
                                        case "foo bar":
                                            var v45 int = fi29
                                            _ = v45
                                        default:
                                            delete(v1, "\n\t")
                                        }
                                        delete(v1, "a")
                                    } else {
                                        fmt.Println(v23, v23, fi29)
                                    }
                                    v46 := GenS6{F0: 18446744073709551615}
                                    _ = v46
                                } else {
                                    v47 := GenS4{F0: -899, F1: 32767}
                                    _ = len(v1)
                                    if (false && v13) {
                                        fmt.Println(v23, v23)
                                    } else {
                                        if (v3 > v3) {
                                            for fi48 := 0; fi48 < 9; fi48++ {
                                                v2 = v13
                                                v49 := []int{(fi48 ^ (fi48 >> ((424) & 63))), (fi29 & fi48), fi48, (fi48 % (((fi29 / (((-(fi48))) | 1))) | 1))}
                                                _ = fi48
                                                _ = v49
                                            }
                                        } else {
                                            v13 = false
                                        }
                                        if v2 {
                                            v2 = ((-2.25 * -2.25) >= (v23 * float64(3.14159)))
                                            v13 = false
                                        }
                                        if true {
                                            v1["\n\t"] = (fi29 + fi29)
                                        } else {
                                            var v50 string = "a"
                                            _ = v50
                                        }
                                    }
                                    _ = v47
                                }
                            }
                            switch 0 {
                            case 230:
                                if v2 {
                                    fmt.Println(fi29)
                                }
                            case 9223372036854775807:
                                var v51 int16 = 42
                                _ = v51
                            case 127:
                                var v52 uint8 = 128
                                _ = v52
                            }
                            fmt.Println(fi29)
                        }
                        for fi53 := 0; fi53 < 16; fi53++ {
                            fmt.Println(v23, v2)
                            switch (fi29 | (-(fi29))) {
                            case 757:
                                fmt.Println(v2)
                            }
                            switch (fi29 << ((255) & 63)) {
                            case 0:
                                var v54 uint16 = 2
                                _ = v54
                            case 1:
                                var v55 uint = 65535
                                _ = v55
                            default:
                                v13 = true
                            }
                            _ = fi53
                        }
                    }
                }
                _ = fi29
            }
        }
        _ = v23
    } else {
        v56 := GenS2{F0: 128, F1: ("" + "hello")}
        if (!v13) {
            v57 := GenS7{F0: 340, F1: 4294967295, F2: -32768}
            _ = v57
        } else {
            fmt.Println(v2, v13, v2)
            switch -339 {
            case 42:
                for fi58 := 0; fi58 < 12; fi58++ {
                    v3 += v3
                    for fi59 := 0; fi59 < 9; fi59++ {
                        v60 := []int{(fi58 >> ((0) & 63)), (fi59 & 2147483647), 32768}
                        _ = fi59
                        _ = v60
                    }
                    v61 := GenS7{F0: 255, F1: 212, F2: 1}
                    _ = fi58
                    _ = v61
                }
            case 4294967295:
                for fi62 := 0; fi62 < 2; fi62++ {
                    if v13 {
                        v63 := GenS8{F0: "\n\t", F1: 987}
                        for fi64 := 0; fi64 < 4; fi64++ {
                            var v65 uint = 850
                            switch (fi64 << (((^(v65))) & 63)) {
                            case 965:
                                var v66 string = ""
                                _ = v66
                            default:
                                v67 := GenS9{F0: 2147483648, F1: int16(v3)}
                                _ = v67
                            }
                            _ = fi64
                            _ = v65
                        }
                        _ = v63
                    } else {
                        v68 := GenS10{F0: (fi62 % (((fi62 >> ((232) & 63))) | 1)), F1: "", F2: ("a" + "a")}
                        v69 := []int{}
                        _ = v68
                        _ = v69
                    }
                    for fi70 := 0; fi70 < 5; fi70++ {
                        v71 := GenS10{F0: (-(fi62)), F1: ("foo bar" + "foo bar"), F2: "foo bar"}
                        _ = fi70
                        _ = v71
                    }
                    var v72 int32 = 42
                    _ = fi62
                    _ = v72
                }
            default:
                var v73 bool = v13
                _ = v73
            }
        }
        _ = v56
    }
    v74 := []int{}
    if ((!v13) && ((42 == 609) || false)) {
        switch -1 {
        case 283:
            v75 := GenS11{F0: v13, F1: 42, F2: 1}
            _ = v75
        default:
            for fi76 := 0; fi76 < 6; fi76++ {
                fmt.Println(fi76)
                _ = fi76
            }
        }
        switch "\n\t" {
        case "":
            v13 = true
        case "\n\t":
            _ = len(v74)
        default:
            for fi77 := 0; fi77 < 8; fi77++ {
                delete(v1, "foo bar")
                fmt.Println(v2)
                _ = fi77
            }
        }
    } else {
        for fi78 := 0; fi78 < 9; fi78++ {
            v79 := v1[""]
            var v80 int64 = v3
            _ = fi78
            _ = v79
            _ = v80
        }
        delete(v1, "foo bar")
        for fi81 := 0; fi81 < 10; fi81++ {
            _ = len(v74)
            _ = fi81
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v74:", len(v74), v74)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v74
}
