package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 string }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 int8; F1 uint16; F2 int32 }
type GenS7 struct { F0 int64; F1 uint16; F2 int32 }
type GenS8 struct { F0 uint64; F1 uint32; F2 string }
type GenS9 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS10 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint64; F1 int8; F2 int16 }
func genFn0(p0 int16) int32 {
    fmt.Println(p0)
    v7 := GenS1{F0: "hello", F1: -985, F2: ("\n\t" + "hello")}
    switch "hello" {
    case "hello":
        p0 -= (^(p0))
    case "":
        v8 := []int{-9223372036854775808, 1}
        _ = v8
    default:
        var v9 uint16 = 65535
        _ = v9
    }
    var v10 uint16 = 1
    var v11 uint64 = 4294967295
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    return 708
}
func genFn1(p0 uint64) uint8 {
    switch (("a" + ("\n\t" + "foo bar")) + "hello") {
    case "a":
        v12 := []int{}
        _ = v12
    case "foo bar":
        v13 := []int{2147483647, 325}
        _ = v13
    case "hello":
        if ((!(false && true)) && false) {
            var v14 int = -478
            var v15 int16 = -1
            _ = v14
            _ = v15
        } else {
            var v16 string = ("foo bar" + ((("foo bar" + ("a" + "hello")) + "hello") + ""))
            p0 -= (p0 | p0)
            v16 = (v16 + (v16 + v16))
            _ = v16
        }
    default:
        if (1 <= 100) {
            switch "hello" {
            case "hello":
                fmt.Println(p0)
            case "a":
                var v17 float32 = 3.14159
                _ = v17
            case "":
                var v18 uint = 2
                _ = v18
            }
            if false {
                var v19 int16 = -1
                switch "" {
                case "\n\t":
                    var v20 uint16 = 361
                    _ = v20
                case "":
                    v19 = (v19 % (((^(v19))) | 1))
                case "hello":
                    p0 -= 0
                }
                _ = v19
            }
            for fi21 := 0; fi21 < 4; fi21++ {
                fmt.Println(p0, p0)
                p0 = 432
                var v22 uint = 255
                _ = fi21
                _ = v22
            }
        } else {
            v23 := GenS1{F0: ((("foo bar" + "a") + "a") + ("foo bar" + "")), F1: 25, F2: "a"}
            _ = v23
        }
    }
    p0 += p0
    var v24 int = -1
    v24 = (v24 << ((4294967295) & 63))
    _ = p0
    _ = v24
    return 137
}

func main() {
    var v1 uint8 = 2
    v2 := []uint32{458, 101, 42, 37}
    var v3 int8 = 127
    v4 := []string{}
    var v5 uint = 1
    var v6 float64 = 1.5
    v25 := GenS1{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: 597, F2: "foo bar"}
    switch 206 {
    case 626:
        var v26 bool = (160 != 1)
        _ = v26
    case 128:
        for fi27 := 0; fi27 < 18; fi27++ {
            v6 -= v6
            switch fi27 {
            case 9223372036854775807:
                v3 = (v3 - v3)
            case 128:
                switch fi27 {
                case -2147483648:
                    for fi28 := 0; fi28 < 13; fi28++ {
                        v6 = 0.0
                        if true {
                            var v29 int8 = (v3 - v3)
                            var v30 uint32 = 2147483647
                            if false {
                                if false {
                                    v31 := GenS1{F0: ("\n\t" + ("" + "foo bar")), F1: (fi28 * 272), F2: "hello"}
                                    _ = v31
                                } else {
                                    var v32 float64 = (0.0 - float64((((1e10 / 0.1) / 1.5) * 1.5)))
                                    var v33 float32 = 1e10
                                    _ = v32
                                    _ = v33
                                }
                            } else {
                                v34 := GenS2{F0: 1, F1: 861}
                                _ = v34
                            }
                            _ = v29
                            _ = v30
                        } else {
                            switch "\n\t" {
                            case "hello":
                                v35 := GenS3{F0: genFn1(1), F1: (521 >= 43)}
                                _ = v35
                            case "\n\t":
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                            case "foo bar":
                                var v36 int8 = 0
                                _ = v36
                            default:
                                v37 := GenS4{F0: -9223372036854775808}
                                _ = v37
                            }
                            v3 += 42
                        }
                        v38 := GenS4{F0: -9223372036854775808}
                        _ = fi28
                        _ = v38
                    }
                case 9223372036854775807:
                    v3 -= (v3 % (((v3 / (((v3 * (v3 << ((v5) & 63)))) | 1))) | 1))
                case 4294967295:
                    var v39 int8 = 42
                    _ = v39
                }
            default:
                v40 := GenS4{F0: -9223372036854775808}
                _ = v40
            }
            _ = fi27
        }
    case 127:
        var v41 uint = 265
        _ = v41
    default:
        var v42 string = "\n\t"
        _ = v42
    }
    var v43 int8 = 0
    v44 := GenS3{F0: v1, F1: true}
    if (true || (4294967295 != 16)) {
        for fi45 := 0; fi45 < 13; fi45++ {
            v46 := GenS5{F0: 0, F1: true}
            _ = fi45
            _ = v46
        }
        var v47 bool = false
        _ = v47
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi48 := 0; fi48 < 9; fi48++ {
        for fi49 := 0; fi49 < 12; fi49++ {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            if (2147483647 >= 412) {
                if len(v4) > 0 {
                    v4[0] = ("foo bar" + "\n\t")
                }
                v43 -= (v3 | 0)
                v2 = append(v2, 2147483648)
            }
            v5 = (v5 ^ v5)
            _ = fi49
        }
        v6 = float64(3.14159)
        var v50 float32 = (1e10 * 0.1)
        _ = fi48
        _ = v50
    }
    v51 := GenS6{F0: v3, F1: 42, F2: -2147483648}
    v4 = append(v4, ("hello" + "\n\t"))
    v1 = (v1 << ((v5) & 63))
    v52 := GenS5{F0: (v43 & v43), F1: false}
    var v53 uint8 = (v1 % ((v1) | 1))
    v54 := GenS3{F0: 255, F1: false}
    if ((786 != 9223372036854775807) || true) {
        var v55 float32 = 0.0
        _ = v55
    } else {
        fmt.Println(v6, v3)
        var v56 float64 = v6
        _ = v56
    }
    if (true && (false && ((v1 & (v1 / (((v1 / (((v1 + genFn1(2))) | 1))) | 1))) != genFn1(4294967295)))) {
        if true {
            fmt.Println(v1)
        } else {
            v57 := GenS5{F0: (v3 | 1), F1: ("a" >= "\n\t")}
            if ((("foo bar" + "a") + "\n\t") > "") {
                switch 9223372036854775807 {
                case 2147483647:
                    v58 := GenS7{F0: 42, F1: 255, F2: -876}
                    _ = v58
                case 32768:
                    v53 += 229
                case 372:
                    if (false && (111 != 32767)) {
                        for fi59 := 0; fi59 < 4; fi59++ {
                            var v60 int = 42
                            _ = fi59
                            _ = v60
                        }
                    } else {
                        v61 := GenS1{F0: ("foo bar" + "\n\t"), F1: -9223372036854775808, F2: ""}
                        for fi62 := 0; fi62 < 19; fi62++ {
                            var v63 int16 = -1
                            _ = fi62
                            _ = v63
                        }
                        _ = v61
                    }
                }
                var v64 uint32 = 2
                _ = v64
            } else {
                if (65535 == (^(v5))) {
                    if false {
                        var v65 int16 = -162
                        for fi66 := 0; fi66 < 3; fi66++ {
                            var v67 uint32 = uint32(v53)
                            v53 = v1
                            if (v3 <= v3) {
                                var v68 uint16 = 32767
                                v68 = 32767
                                _ = v68
                            }
                            _ = fi66
                            _ = v67
                        }
                        if ((v6 * (v6 * v6)) >= 0.1) {
                            if (v1 == genFn1(42)) {
                                v69 := GenS8{F0: 428, F1: 42, F2: "\n\t"}
                                v70 := GenS9{F0: 1, F1: 950, F2: v1}
                                _ = v69
                                _ = v70
                            }
                            var v71 int64 = 1
                            v65 += 32767
                            _ = v71
                        } else {
                            for fi72 := 0; fi72 < 9; fi72++ {
                                fmt.Println(v43, v5)
                                _ = fi72
                            }
                        }
                        _ = v65
                    }
                    fmt.Println(v43, v6)
                }
                for fi73 := 0; fi73 < 18; fi73++ {
                    _ = len(v2)
                    _ = fi73
                }
                switch (("foo bar" + "") + "") {
                case "":
                    var v74 float64 = (3.14159 * v6)
                    _ = v74
                }
            }
            var v75 int64 = 337
            _ = v57
            _ = v75
        }
    } else {
        for fi76 := 0; fi76 < 16; fi76++ {
            if (42 > 4294967295) {
                for fi77 := 0; fi77 < 14; fi77++ {
                    var v78 uint8 = genFn1(42)
                    switch "a" {
                    case "hello":
                        v6 = (v6 - -2.25)
                    case "foo bar":
                        v79 := GenS2{F0: 775, F1: int16(v43)}
                        _ = v79
                    case "":
                        fmt.Println(v43, v78)
                    default:
                        v3 = (v43 - v3)
                    }
                    v78 = 128
                    _ = fi77
                    _ = v78
                }
                var v80 int32 = -1
                _ = v80
            }
            if false {
                v6 -= (((v6 / v6) + -2.25) - (v6 - ((-2.25 - v6) - 0.0)))
            }
            v81 := GenS2{F0: 42, F1: 0}
            _ = fi76
            _ = v81
        }
        for fi82 := 0; fi82 < 18; fi82++ {
            if (v53 < (v53 / (((v1 / ((genFn1(18446744073709551615)) | 1))) | 1))) {
                for fi83 := 0; fi83 < 12; fi83++ {
                    var v84 float32 = 0.1
                    v85 := GenS4{F0: -234}
                    _ = fi83
                    _ = v84
                    _ = v85
                }
                var v86 string = ""
                _ = v86
            } else {
                v87 := GenS2{F0: 0, F1: 758}
                var v88 int8 = 10
                _ = v87
                _ = v88
            }
            _ = fi82
        }
    }
    for fi89 := 0; fi89 < 10; fi89++ {
        if false {
            var v90 uint32 = 2147483648
            var v91 int8 = v43
            v92 := GenS4{F0: 9223372036854775807}
            _ = v90
            _ = v91
            _ = v92
        } else {
            v93 := GenS10{F0: 42, F1: uint32(v1), F2: 42}
            var v94 int32 = genFn0(-892)
            _ = v93
            _ = v94
        }
        if (false && true) {
            v95 := GenS11{F0: 0}
            _ = v95
        } else {
            v96 := GenS1{F0: "", F1: -81, F2: (("a" + "\n\t") + "")}
            for fi97 := 0; fi97 < 6; fi97++ {
                var v98 uint64 = 1
                _ = fi97
                _ = v98
            }
            v6 = 0.0
            _ = v96
        }
        _ = fi89
    }
    for fi99 := 0; fi99 < 19; fi99++ {
        var v100 float32 = 1.5
        _ = fi99
        _ = v100
    }
    if (0 < 1) {
        v1 -= (v53 >> (((v5 & 18446744073709551615)) & 63))
        if false {
            for fi101 := 0; fi101 < 20; fi101++ {
                v4 = append(v4, "a")
                _ = fi101
            }
            switch 556 {
            case -1:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                for fi102 := 0; fi102 < 10; fi102++ {
                    var v103 uint8 = v53
                    var v104 uint16 = 1
                    _ = fi102
                    _ = v103
                    _ = v104
                }
            }
        }
        var v105 int64 = 42
        _ = v105
    } else {
        var v106 float32 = ((1e10 + (float32(v6) * 0.1)) - -2.25)
        v107 := GenS12{F0: 449, F1: v43, F2: 565}
        _ = len(v4)
        _ = v106
        _ = v107
    }
    v53 -= uint8(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
