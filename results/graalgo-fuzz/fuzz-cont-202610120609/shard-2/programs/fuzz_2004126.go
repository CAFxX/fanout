package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 uint; F1 uint; F2 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint32; F1 string }
type GenS12 struct { F0 int16; F1 uint64 }
type GenS13 struct { F0 int; F1 int64; F2 uint32 }
type GenS14 struct { F0 string; F1 uint; F2 uint64 }

func main() {
    var v1 uint32 = 42
    v2 := []uint32{v1}
    var v3 float32 = 1.5
    var v4 uint64 = 57
    var v5 int8 = -1
    var v6 float32 = v3
    v7 := GenS1{F0: int32(v5), F1: 32767, F2: v5}
    if (("a" + (("a" + "foo bar") + "a")) == (("" + "\n\t") + "foo bar")) {
        var v8 string = ""
        if ((v5 & -18) < -49) {
            var v9 uint16 = 470
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v9
        }
        _ = v8
    }
    if (-1 != -9223372036854775808) {
        if (42 == v1) {
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 bool = true
                if ((!v11) || (fi10 == fi10)) {
                    var v12 int16 = -1
                    if (v1 <= 750) {
                        if (v11 || (true || v11)) {
                            if false {
                                for fi13 := 0; fi13 < 3; fi13++ {
                                    v14 := GenS1{F0: 170, F1: 438, F2: (v5 / ((v5) | 1))}
                                    var v15 uint64 = v4
                                    v16 := GenS2{F0: 661, F1: v1, F2: 1}
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                            } else {
                                for fi17 := 0; fi17 < 17; fi17++ {
                                    if (int32(fi17) > int32(v1)) {
                                        if (!v11) {
                                            fmt.Println(v11, v5, v6)
                                            fmt.Println(v3, v4, v1)
                                        } else {
                                            var v18 uint32 = (v1 * (v1 - 195))
                                            v2 = append(v2, 1)
                                            var v19 int16 = (v12 >> ((18446744073709551615) & 63))
                                            var v20 string = ("hello" + "foo bar")
                                            if v11 {
                                                fmt.Println(v20, v1)
                                            }
                                            _ = v18
                                            _ = v19
                                            _ = v20
                                        }
                                        v21 := GenS3{F0: 0, F1: v4, F2: (v5 & (v5 / (((v5 >> ((4294967295) & 63))) | 1)))}
                                        _ = v21
                                    }
                                    _ = fi17
                                }
                                v2 = append(v2, (v1 | 2147483647))
                                switch (fi10 + -534) {
                                case -1:
                                    var v22 int64 = 42
                                    _ = v22
                                default:
                                    v6 += 0.1
                                }
                            }
                            v23 := GenS4{F0: 2}
                            _ = v23
                        }
                    }
                    fmt.Println(v12, v1, v4)
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
        } else {
            if true {
                var v24 uint16 = 42
                switch 32768 {
                case 626:
                    v24 = 1
                case -2147483648:
                    if true {
                        var v25 float32 = (v3 - v3)
                        for fi26 := 0; fi26 < 19; fi26++ {
                            if true {
                                v27 := GenS5{F0: (v4 | (^(v4))), F1: false}
                                var v28 float32 = float32(-2.25)
                                v24 = v24
                                _ = v27
                                _ = v28
                            }
                            var v29 float32 = ((0.0 - (v6 / -2.25)) * v25)
                            _ = fi26
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 15; fi30++ {
                            var v31 uint16 = v24
                            _ = fi30
                            _ = v31
                        }
                        _ = v25
                    }
                case 259:
                    v32 := GenS4{F0: 0}
                    _ = v32
                default:
                    v3 = 3.14159
                }
                _ = v24
            } else {
                var v33 string = "\n\t"
                if true {
                    if (("hello" > v33) && false) {
                        for fi34 := 0; fi34 < 13; fi34++ {
                            for fi35 := 0; fi35 < 8; fi35++ {
                                if (((!((107 >= 402) && false)) || (!(1 < v5))) || (!true)) {
                                    if ((!(!false)) || true) {
                                        v33 = (v33 + v33)
                                    }
                                    _ = len(v2)
                                    var v36 uint32 = 1
                                    _ = v36
                                } else {
                                    fmt.Println(v33)
                                    v4 = (v4 / ((v4) | 1))
                                }
                                var v37 uint16 = 2
                                fmt.Println(fi35, fi35, v37)
                                v38 := GenS5{F0: uint64(v37), F1: false}
                                fmt.Println(v5, v33, v3)
                                var v39 uint16 = (v37 >> ((1) & 63))
                                v2 = append(v2, v1)
                                _ = fi35
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            v1 += (v1 >> ((704) & 63))
                            var v40 string = v33
                            _ = fi34
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 19; fi41++ {
                            var v42 int32 = 0
                            v43 := GenS4{F0: 1}
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        v44 := GenS6{F0: 18446744073709551615, F1: 18446744073709551615, F2: 359}
                        _ = v44
                    }
                    var v45 int32 = -2147483648
                    _ = v45
                } else {
                    var v46 uint64 = (v4 >> ((65535) & 63))
                    v47 := GenS7{F0: 2}
                    _ = v46
                    _ = v47
                }
                v4 = v4
                _ = v33
            }
            var v48 int32 = -1
            _ = v48
        }
        switch -468 {
        case 128:
            var v49 int32 = 1
            _ = v49
        default:
            if len(v2) > 0 {
                v2[0] = (v1 ^ v1)
            }
        }
        var v50 bool = false
        _ = v50
    } else {
        switch ("" + (("hello" + "\n\t") + ("foo bar" + ("hello" + "\n\t")))) {
        case "a":
            var v51 uint64 = 1
            _ = v51
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v52 := GenS8{F0: 1}
            _ = v52
        }
    }
    v6 -= v3
    var v53 uint8 = 128
    if false {
        if true {
            if ((42 != 0) && (true || ((v5 >> ((366) & 63)) > v5))) {
                for fi54 := 0; fi54 < 19; fi54++ {
                    if len(v2) > 0 {
                        v2[0] = 10
                    }
                    _ = fi54
                }
            } else {
                for fi55 := 0; fi55 < 12; fi55++ {
                    var v56 float32 = ((v3 - v6) - (-2.25 * (v6 * 0.0)))
                    for fi57 := 0; fi57 < 17; fi57++ {
                        v58 := GenS3{F0: 255, F1: (^(v4)), F2: v5}
                        v59 := GenS6{F0: 904, F1: 255, F2: (fi57 * (fi57 >> ((747) & 63)))}
                        if len(v2) > 0 {
                            v2[0] = (^(v1))
                        }
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                    _ = fi55
                    _ = v56
                }
                v60 := GenS9{F0: 2, F1: 65535}
                _ = v60
            }
            var v61 string = "\n\t"
            _ = v61
        } else {
            for fi62 := 0; fi62 < 20; fi62++ {
                v5 = v5
                _ = fi62
            }
            if true {
                if false {
                    fmt.Println(v4, v6)
                    var v63 int32 = -2147483648
                    v63 += int32(v5)
                    _ = v63
                }
            }
        }
        var v64 float32 = v3
        _ = v64
    } else {
        var v65 uint = 65535
        _ = v65
    }
    var v66 uint16 = 32767
    for fi67 := 0; fi67 < 9; fi67++ {
        for fi68 := 0; fi68 < 17; fi68++ {
            v6 = v6
            _ = len(v2)
            var v69 bool = (!(true || false))
            v4 = 42
            if (2147483647 > 1) {
                var v70 uint32 = v1
                _ = v70
            } else {
                v71 := GenS10{F0: "foo bar"}
                _ = v71
            }
            _ = fi68
            _ = v69
        }
        switch "\n\t" {
        case "foo bar":
            v72 := GenS11{F0: v1, F1: ((((("" + "foo bar") + "") + "a") + "\n\t") + "")}
            _ = v72
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v1 -= v1
        }
        _ = len(v2)
        _ = fi67
    }
    for fi73 := 0; fi73 < 1; fi73++ {
        if len(v2) > 0 {
            v2[0] = v1
        }
        if ((0 != -1) || true) {
            v74 := GenS1{F0: 0, F1: 42, F2: (v5 / (((v5 << ((447) & 63))) | 1))}
            var v75 bool = true
            _ = v74
            _ = v75
        } else {
            if ((v53 ^ 1) >= v53) {
                v6 -= v6
                for fi76 := 0; fi76 < 13; fi76++ {
                    if (!(!true)) {
                        switch "a" {
                        case "":
                            v66 = 108
                            var v77 float32 = ((v6 + float32((0.0 * 0.0))) + float32(((((float64((((float32((-2.25 - (0.1 * 0.1))) - float32(((0.0 + 1.5) + float64(((v6 * ((v6 - v3) / 0.1)) + 0.1))))) - v6) - float32(((float64(0.1) + -2.25) / float64((float32(-2.25) + ((0.1 * (v3 * 0.0)) / v6))))))) / (float64(0.0) + (float64((v6 / v3)) * 1e10))) / 1.0) / 0.1) * -2.25)))
                            _ = v77
                        case "\n\t":
                            fmt.Println(v6)
                        default:
                            fmt.Println(v1)
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = (^(v1))
                    }
                    _ = fi76
                }
                var v78 uint = 2
                _ = v78
            }
            v79 := GenS11{F0: (v1 | (v1 / ((2147483648) | 1))), F1: "a"}
            _ = v79
        }
        _ = fi73
    }
    fmt.Println(v4, v3)
    v2 = append(v2, 2)
    v80 := GenS12{F0: 0, F1: 825}
    var v81 int = -1
    v53 = v53
    v82 := GenS13{F0: v81, F1: -9223372036854775808, F2: v1}
    switch 2147483647 {
    case 4294967295:
        var v83 uint16 = 2
        _ = v83
    case -812:
        var v84 int64 = 42
        _ = v84
    default:
        var v85 int32 = 2147483647
        _ = v85
    }
    v86 := GenS14{F0: "", F1: 0, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v66
    _ = v7
    _ = v80
    _ = v81
    _ = v82
    _ = v86
}
