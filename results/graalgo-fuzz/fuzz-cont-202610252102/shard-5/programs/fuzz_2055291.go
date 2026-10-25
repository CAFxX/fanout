package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int64; F1 int64; F2 int }
type GenS7 struct { F0 int64; F1 int32; F2 uint8 }
type GenS8 struct { F0 uint }
func genFn0() int {
    v6 := []int{2147483647, 0}
    var v7 uint32 = 107
    var v8 int16 = -227
    fmt.Println(v8, v7)
    var v9 uint16 = 714
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint8, p1 int64, p2 bool) uint64 {
    p1 = p1
    if ((-2.25 * 1.5) == (-2.25 / 1e10)) {
        fmt.Println(p1, p0)
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 uint16 = 0
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(fi10, p1, fi12)
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(p0)
    v13 := GenS1{F0: -1, F1: 1, F2: int(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 4294967295
}

func main() {
    v1 := []string{(("foo bar" + "\n\t") + "a"), "\n\t"}
    var v2 int = 127
    var v3 float32 = 1e10
    var v4 int = v2
    var v5 int8 = 127
    for fi14 := 0; fi14 < 7; fi14++ {
        if true {
            if (156 > 42) {
                if len(v1) > 0 {
                    v1[0] = (("\n\t" + ("\n\t" + ("" + "a"))) + "hello")
                }
                for fi15 := 0; fi15 < 18; fi15++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi15
                }
            }
            if (true && false) {
                v1 = append(v1, ("" + "foo bar"))
            } else {
                v16 := GenS1{F0: (-(v5)), F1: 42, F2: fi14}
                _ = v16
            }
            v2 -= genFn0()
        }
        _ = fi14
    }
    v1 = append(v1, "hello")
    v2 += (v4 ^ v2)
    if ((!(!true)) || (float64(v3) >= (1e10 + 1e10))) {
        v1 = append(v1, ("" + "hello"))
        var v17 int64 = -173
        v3 -= (v3 - ((3.14159 / -2.25) - (-2.25 / ((float32((float64((1.5 / float32(0.1))) - 1.5)) / (0.0 - 1e10)) + v3))))
        _ = v17
    }
    switch "foo bar" {
    case "\n\t":
        if (false && (((v4 & v2) > 32767) || (true || (0 < 42)))) {
            for fi18 := 0; fi18 < 10; fi18++ {
                _ = len(v1)
                var v19 float64 = (1e10 + 3.14159)
                _ = fi18
                _ = v19
            }
            var v20 int16 = 42
            if true {
                if true {
                    fmt.Println(v4, v4)
                    var v21 uint8 = 212
                    _ = v21
                } else {
                    if (42 <= 323) {
                        var v22 int8 = (v5 >> ((255) & 63))
                        _ = v22
                    }
                    fmt.Println(v4, v20)
                    fmt.Println(v4, v3)
                }
                if ((v20 * v20) == v20) {
                    for fi23 := 0; fi23 < 5; fi23++ {
                        for fi24 := 0; fi24 < 15; fi24++ {
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            _ = fi24
                        }
                        if false {
                            var v25 int = (v4 ^ (v4 >> ((648) & 63)))
                            for fi26 := 0; fi26 < 14; fi26++ {
                                v27 := GenS2{F0: 875, F1: (("\n\t" + "") + "\n\t")}
                                fmt.Println(v2)
                                var v28 uint = 1
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            _ = v25
                        } else {
                            for fi29 := 0; fi29 < 20; fi29++ {
                                _ = len(v1)
                                fmt.Println(fi23, fi23)
                                switch 42 {
                                case 4294967295:
                                    v3 = ((v3 - (v3 / v3)) - float32(1.5))
                                case 128:
                                    v2 -= (^(fi29))
                                case 2147483647:
                                    var v30 uint = 121
                                    _ = v30
                                default:
                                    if false {
                                        var v31 uint8 = 164
                                        v32 := GenS2{F0: 4294967295, F1: "hello"}
                                        var v33 bool = (((!(!((1 <= 0) || true))) && (v5 >= v5)) || false)
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                }
                                _ = fi29
                            }
                            v34 := GenS2{F0: genFn1(0, 0, false), F1: (("a" + "\n\t") + "")}
                            for fi35 := 0; fi35 < 10; fi35++ {
                                if false {
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v4, v4, v3)
                                        if false {
                                            v2 += v4
                                            v36 := GenS3{F0: 32767, F1: 0}
                                            _ = v36
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = "hello"
                                            }
                                        }
                                    }
                                    if (!(!(false && (!true)))) {
                                        v2 -= (fi35 >> ((334) & 63))
                                        if (318 <= 0) {
                                            if (65535 >= 827) {
                                                v5 -= (^(v5))
                                            }
                                            var v37 float32 = float32(0.1)
                                            _ = v37
                                        }
                                        v38 := GenS3{F0: 42, F1: 42}
                                        var v39 int8 = 42
                                        _ = v38
                                        _ = v39
                                    }
                                    v3 = v3
                                }
                                _ = fi35
                            }
                            _ = v34
                        }
                        _ = fi23
                    }
                    for fi40 := 0; fi40 < 2; fi40++ {
                        fmt.Println(v20, fi40)
                        v5 -= 1
                        _ = fi40
                    }
                }
                for fi41 := 0; fi41 < 18; fi41++ {
                    if false {
                        fmt.Println(v5, v4)
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        fmt.Println(fi41, v20)
                    }
                    _ = fi41
                }
            } else {
                v42 := GenS1{F0: int8(v2), F1: -287, F2: v2}
                v43 := GenS3{F0: 0, F1: -2147483648}
                _ = len(v1)
                _ = v42
                _ = v43
            }
            _ = v20
        } else {
            fmt.Println(v2)
            var v44 uint8 = 0
            _ = v44
        }
    default:
        var v45 uint64 = 824
        _ = v45
    }
    v2 = (-(v2))
    var v46 uint = 30
    var v47 int16 = 42
    for fi48 := 0; fi48 < 20; fi48++ {
        v2 = (-(fi48))
        _ = fi48
    }
    v49 := GenS2{F0: 0, F1: ("foo bar" + ("foo bar" + "\n\t"))}
    var v50 int32 = 2147483647
    if true {
        var v51 string = "foo bar"
        _ = v51
    }
    for fi52 := 0; fi52 < 9; fi52++ {
        var v53 float64 = 0.1
        _ = fi52
        _ = v53
    }
    for fi54 := 0; fi54 < 5; fi54++ {
        if true {
            var v55 uint8 = 119
            for fi56 := 0; fi56 < 10; fi56++ {
                if (!(683 != 353)) {
                    if (!((v55 / ((141) | 1)) <= (v55 - (v55 ^ (^(v55)))))) {
                        v57 := GenS4{F0: 42}
                        _ = v57
                    }
                    var v58 int = (fi56 % ((v2) | 1))
                    _ = v58
                }
                _ = fi56
            }
            _ = v55
        } else {
            if len(v1) > 0 {
                v1[0] = ""
            }
            if true {
                v59 := GenS4{F0: -333}
                _ = v59
            } else {
                for fi60 := 0; fi60 < 2; fi60++ {
                    var v61 bool = (((((816 <= 0) || (!false)) && false) && false) && false)
                    var v62 uint = v46
                    _ = fi60
                    _ = v61
                    _ = v62
                }
                var v63 float32 = 3.14159
                v46 = v46
                _ = v63
            }
        }
        v64 := GenS4{F0: 42}
        switch (("a" + (("" + "a") + "a")) + ("" + "\n\t")) {
        case "\n\t":
            if (831 != 65535) {
                v2 = 42
                var v65 int32 = 0
                _ = v65
            } else {
                switch "foo bar" {
                case "":
                    var v66 int32 = v50
                    _ = v66
                case "\n\t":
                    if (1 != -128) {
                        for fi67 := 0; fi67 < 2; fi67++ {
                            v5 = (^(v5))
                            v68 := GenS4{F0: 42}
                            if false {
                                v1 = append(v1, "hello")
                                fmt.Println(v2, fi54, v2)
                            } else {
                                fmt.Println(v2, v5, v2)
                                v69 := GenS5{F0: 0, F1: "hello"}
                                if ((42 >= 2) && (!false)) {
                                    _ = len(v1)
                                } else {
                                    v5 = 0
                                    fmt.Println(v4, v5)
                                    fmt.Println(v46)
                                    v1 = append(v1, "a")
                                }
                                if true {
                                    switch (^(v2)) {
                                    case 434:
                                        switch "a" {
                                        case "a":
                                            var v70 int32 = 2147483647
                                            _ = v70
                                        case "\n\t":
                                            v71 := GenS6{F0: 42, F1: -1, F2: v2}
                                            _ = v71
                                        }
                                    case 32768:
                                        v50 = (v50 % ((v50) | 1))
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch 317 {
                                    case -3:
                                        var v72 int64 = -1
                                        _ = v72
                                    case -228:
                                        switch v2 {
                                        case 9223372036854775807:
                                            var v73 float64 = 0.0
                                            v74 := GenS5{F0: 32767, F1: ""}
                                            var v75 int64 = -9223372036854775808
                                            if (!true) {
                                                if false {
                                                    fmt.Println(v47, v2)
                                                } else {
                                                    var v76 float64 = 0.0
                                                    var v77 string = "a"
                                                    _ = v76
                                                    _ = v77
                                                }
                                            }
                                            _ = v73
                                            _ = v74
                                            _ = v75
                                        default:
                                            var v78 float64 = 1e10
                                            _ = v78
                                        }
                                        var v79 uint8 = 208
                                        var v80 int = v2
                                        _ = v79
                                        _ = v80
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v69
                            }
                            _ = fi67
                            _ = v68
                        }
                    } else {
                        v81 := GenS7{F0: -1, F1: v50, F2: 2}
                        if false {
                            v4 -= (-(fi54))
                        } else {
                            switch fi54 {
                            case 2147483647:
                                var v82 uint16 = 1
                                _ = v82
                            }
                        }
                        _ = v81
                    }
                }
            }
        case "hello":
            for fi83 := 0; fi83 < 20; fi83++ {
                v4 = int(v47)
                if true {
                    if false {
                        v84 := GenS8{F0: 4294967295}
                        _ = v84
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v2 = v2
                }
                _ = fi83
            }
        default:
            var v85 uint32 = 0
            _ = v85
        }
        _ = fi54
        _ = v64
    }
    v46 = (v46 << ((v46) & 63))
    if true {
        fmt.Println(v4, v3)
        switch "" {
        case "a":
            var v86 float32 = 1e10
            _ = v86
        case "\n\t":
            for fi87 := 0; fi87 < 15; fi87++ {
                _ = len(v1)
                v88 := GenS1{F0: (^(v5)), F1: -2147483648, F2: v2}
                _ = fi87
                _ = v88
            }
        case "hello":
            v4 = v2
        }
        v4 += (v2 & (v4 ^ v2))
    }
    for fi89 := 0; fi89 < 2; fi89++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi89
    }
    v1 = append(v1, ("hello" + ("" + "a")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v49
    _ = v5
    _ = v50
}
