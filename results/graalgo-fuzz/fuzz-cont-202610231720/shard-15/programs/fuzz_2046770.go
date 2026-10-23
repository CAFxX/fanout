package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int8; F1 bool; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 string; F1 int16; F2 int32 }
type GenS7 struct { F0 int; F1 int; F2 uint16 }
type GenS8 struct { F0 int8; F1 int; F2 int16 }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int32 }
func genFn0() int8 {
    if false {
        var v7 uint = 65535
        v8 := []int{1, -1, 661}
        _ = v7
        _ = v8
    } else {
        var v9 bool = true
        var v10 float32 = (1.5 * ((0.1 - 3.14159) + 1e10))
        _ = v10
        _ = v9
    }
    var v11 int32 = 461
    _ = v11
    return -22
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    v3 := []int{0, 42, 121}
    var v4 int16 = -32768
    var v5 uint = 18446744073709551615
    var v6 float32 = (0.1 / ((-2.25 / 1.0) * -2.25))
    v12 := v1[4294967295]
    v13 := GenS1{F0: 9223372036854775807}
    v14 := GenS1{F0: 635}
    v15 := v2[128]
    if (18446744073709551615 <= 18446744073709551615) {
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := GenS2{F0: 694, F1: 999}
            var v18 int64 = 9223372036854775807
            v6 -= v6
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 bool = (42 != 9223372036854775807)
        if v15 {
            v6 = v6
            switch 127 {
            case -654:
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 uint64 = 0
                    v1[4294967295] = ((v12 + "foo bar") + ((((((v12 + v12) + (("a" + "foo bar") + ((v12 + "foo bar") + ("\n\t" + "\n\t")))) + ((v12 + "foo bar") + "foo bar")) + ((("\n\t" + v12) + v12) + ("hello" + ""))) + v12) + ""))
                    if (!(0 >= 255)) {
                        if v15 {
                            v22 := v2[fi20]
                            v5 -= (v5 << ((479) & 63))
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 3; fi23++ {
                                v12 = v12
                                _ = fi23
                            }
                        }
                    }
                    _ = fi20
                    _ = v21
                }
            case 42:
                if true {
                    v24 := GenS3{F0: genFn0(), F1: (false && (42 < 2)), F2: v5}
                    _ = v24
                } else {
                    switch 32767 {
                    case -2147483648:
                        v6 = float32(1.5)
                    case 1:
                        v25 := GenS4{F0: v5}
                        _ = v25
                    }
                }
            default:
                fmt.Println(v6)
            }
            if (((1e10 / float64((v6 + (v6 * (v6 / 3.14159))))) + float64(v6)) != -2.25) {
                for fi26 := 0; fi26 < 7; fi26++ {
                    _ = len(v3)
                    _ = fi26
                }
            } else {
                v15 = false
                if ((32767 == 2) || true) {
                    v15 = false
                } else {
                    var v27 float32 = float32(1.5)
                    v6 += (v27 / v6)
                    for fi28 := 0; fi28 < 11; fi28++ {
                        v29 := GenS1{F0: 968}
                        delete(v1, 65535)
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                }
                v4 = v4
            }
        } else {
            if true {
                switch 840 {
                case 4294967295:
                    var v30 uint64 = 0
                    _ = v30
                case 0:
                    if (v15 || true) {
                        fmt.Println(v15, v4)
                        v4 = v4
                    } else {
                        v31 := GenS5{F0: (v5 ^ v5), F1: 617}
                        _ = v31
                    }
                default:
                    v19 = (v19 && (v6 == v6))
                }
            } else {
                if (!true) {
                    if len(v3) > 0 {
                        v3[0] = 440
                    }
                    for fi32 := 0; fi32 < 19; fi32++ {
                        if (80 < int8(v5)) {
                            var v33 int64 = 1
                            _ = v33
                        } else {
                            var v34 int = (^(fi32))
                            v4 = v4
                            switch (v12 + v12) {
                            case "hello":
                                if true {
                                    if (v19 && ((0 == 65535) || (!v15))) {
                                        var v35 float64 = 0.1
                                        _ = v35
                                    } else {
                                        v4 = (v4 * v4)
                                    }
                                    var v36 int = -966
                                    fmt.Println(v5, v5, v19)
                                    _ = v36
                                } else {
                                    for fi37 := 0; fi37 < 6; fi37++ {
                                        v12 = v12
                                        _ = len(v1)
                                        _ = fi37
                                    }
                                }
                            }
                            _ = v34
                        }
                        _ = fi32
                    }
                    var v38 int = 32768
                    _ = v38
                }
            }
            v12 = v12
        }
        _ = v19
    } else {
        if len(v3) > 0 {
            v3[0] = 128
        }
        var v39 int64 = 9223372036854775807
        fmt.Println(v4, v15)
        _ = v39
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        var v41 int64 = -99
        _ = fi40
        _ = v41
    }
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    var v42 string = (v12 + v12)
    if false {
        var v43 float64 = 1e10
        switch 32768 {
        case -2147483648:
            v44 := GenS6{F0: v42, F1: (v4 - v4), F2: 53}
            _ = v44
        case 4294967295:
            for fi45 := 0; fi45 < 6; fi45++ {
                v46 := GenS7{F0: (-(fi45)), F1: 42, F2: 65535}
                var v47 int64 = -9223372036854775808
                _ = fi45
                _ = v46
                _ = v47
            }
        default:
            switch -1 {
            case 127:
                if v15 {
                    if v15 {
                        delete(v1, 1)
                    }
                    if (18446744073709551615 > 690) {
                        for fi48 := 0; fi48 < 10; fi48++ {
                            fmt.Println(v5, v5, v15)
                            switch 745 {
                            case 42:
                                switch (fi48 & fi48) {
                                case -9223372036854775808:
                                    v5 = (v5 >> ((v5) & 63))
                                case 128:
                                    fmt.Println(fi48, v42, v15)
                                default:
                                    for fi49 := 0; fi49 < 15; fi49++ {
                                        switch ("" + v42) {
                                        case "\n\t":
                                            v50 := GenS8{F0: 0, F1: (fi49 % ((32768) | 1)), F2: -32768}
                                            _ = v50
                                        case "":
                                            var v51 string = ""
                                            _ = v51
                                        case "hello":
                                            delete(v2, 1)
                                        default:
                                            var v52 bool = (v15 && (!v15))
                                            _ = v52
                                        }
                                        switch (v42 + v42) {
                                        case "a":
                                            v2[(fi49 / (((fi48 + fi49)) | 1))] = (float32(v43) > 0.0)
                                            if v15 {
                                                v53 := GenS4{F0: v5}
                                                _ = v53
                                            } else {
                                                if (804 < 1) {
                                                    if (127 >= 0) {
                                                        switch (v42 + (v12 + ((("\n\t" + "foo bar") + "") + v42))) {
                                                        case "a":
                                                            _ = len(v2)
                                                            var v54 int16 = v4
                                                            v55 := GenS9{F0: 127, F1: (v15 || (842 != 42))}
                                                            _ = v54
                                                            _ = v55
                                                        case "foo bar":
                                                            fmt.Println(fi49)
                                                        case "hello":
                                                            var v56 uint64 = 886
                                                            _ = v56
                                                        default:
                                                            v43 -= float64(0.0)
                                                        }
                                                    }
                                                    if (!(!v15)) {
                                                        v43 += v43
                                                    } else {
                                                        fmt.Println(v5, v15, v15)
                                                        switch (fi48 / ((-149) | 1)) {
                                                        case 892:
                                                            switch (fi49 << (((v5 - (v5 & (v5 + (v5 / (((v5 | v5)) | 1)))))) & 63)) {
                                                            case 127:
                                                                if (1.5 >= (v43 + ((1.5 / float64(((v6 * v6) / v6))) * v43))) {
                                                                    v57 := GenS10{F0: 0}
                                                                    _ = v57
                                                                } else {
                                                                    var v58 int = fi48
                                                                    _ = v58
                                                                }
                                                            case 32767:
                                                                var v59 string = v42
                                                                _ = v59
                                                            default:
                                                                v5 = v5
                                                            }
                                                        }
                                                        if v15 {
                                                            var v60 int = 460
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v60
                                                        } else {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            switch 78 {
                                                            case 42:
                                                                v61 := GenS4{F0: v5}
                                                                _ = v61
                                                            case -580:
                                                                v62 := GenS5{F0: (^(v5)), F1: fi49}
                                                                _ = v62
                                                            default:
                                                                var v63 uint = v5
                                                                _ = v63
                                                            }
                                                            var v64 float64 = v43
                                                            _ = v64
                                                        }
                                                    }
                                                }
                                                if ((v42 + (((v42 + (v12 + v12)) + (((v12 + "\n\t") + v42) + "\n\t")) + "\n\t")) <= v42) {
                                                    v6 += v6
                                                    v65 := GenS9{F0: -1, F1: false}
                                                    _ = v65
                                                }
                                                v12 = v42
                                                if len(v3) > 0 {
                                                    v3[0] = int(v5)
                                                }
                                                _ = len(v1)
                                                fmt.Println(fi49, v42)
                                            }
                                            var v66 int8 = 100
                                            _ = v66
                                        default:
                                            v67 := GenS10{F0: 16}
                                            _ = v67
                                        }
                                        if (960 == 9223372036854775807) {
                                            if v15 {
                                                var v68 uint32 = 1
                                                _ = v68
                                            }
                                        }
                                        _ = fi49
                                    }
                                }
                            default:
                                var v69 int16 = (v4 % (((v4 | (v4 | (v4 ^ (v4 | int16(fi48)))))) | 1))
                                _ = v69
                            }
                            var v70 int8 = 56
                            _ = fi48
                            _ = v70
                        }
                        if true {
                            switch ("hello" + "foo bar") {
                            case "\n\t":
                                var v71 bool = v15
                                _ = v71
                            }
                        }
                    }
                } else {
                    v12 = "a"
                    v15 = v15
                    if (v15 && v15) {
                        v42 = v42
                    }
                }
            default:
                v43 = (3.14159 / (float64(v6) - 0.0))
            }
        }
        _ = v43
    } else {
        switch -9223372036854775808 {
        case 4294967295:
            v42 = ((v42 + v42) + "hello")
        case 45:
            fmt.Println(v12, v4, v15)
        default:
            v42 = v12
        }
        if (!v15) {
            var v72 int16 = v4
            _ = len(v3)
            _ = v72
        } else {
            if len(v3) > 0 {
                v3[0] = -9223372036854775808
            }
            switch v12 {
            case "":
                v12 = v42
            case "\n\t":
                fmt.Println(v42, v6, v42)
            }
            v5 -= 939
        }
    }
    for fi73 := 0; fi73 < 17; fi73++ {
        fmt.Println(v12, v6)
        _ = fi73
    }
    v12 = (v42 + (v42 + v12))
    fmt.Println(v6, v42)
    fmt.Println(v15, v42)
    v42 = ("hello" + v12)
    for fi74 := 0; fi74 < 5; fi74++ {
        if v15 {
            _ = len(v1)
        }
        v5 += v5
        var v75 uint32 = 65535
        _ = fi74
        _ = v75
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    if v15 {
        if v15 {
            if (v5 >= 18446744073709551615) {
                v6 -= (v6 * v6)
            } else {
                var v76 uint32 = 768
                if v15 {
                    v77 := GenS11{F0: 42}
                    _ = len(v3)
                    _ = v77
                } else {
                    v6 += v6
                    if v15 {
                        var v78 int64 = 0
                        v4 += (v4 / ((v4) | 1))
                        _ = v78
                    }
                    var v79 bool = (v15 && ((v4 >> ((v5) & 63)) <= (v4 >> ((v5) & 63))))
                    _ = v79
                }
                _ = v76
            }
        }
        if len(v3) > 0 {
            v3[0] = 904
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
