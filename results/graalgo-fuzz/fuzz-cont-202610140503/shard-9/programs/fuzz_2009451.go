package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint; F2 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 uint64; F2 int }
func genFn0() uint {
    var v7 int64 = 523
    if false {
        v7 = 533
        var v8 uint64 = 789
        if ((2 > 65535) || true) {
            v8 -= v8
            v9 := GenS1{F0: 245}
            _ = v9
        } else {
            v10 := []int{128, 9223372036854775807, -2147483648}
            v8 = v8
            if (v8 >= 18446744073709551615) {
                for fi11 := 0; fi11 < 11; fi11++ {
                    if (!false) {
                        var v12 int64 = -346
                        v10 = append(v10, fi11)
                        v12 = (v7 * -513)
                        _ = v12
                    }
                    if true {
                        if (true && (!true)) {
                            v7 = (v7 & v7)
                            v8 -= (v8 * (v8 / (((v8 + (v8 & (v8 << ((641) & 63))))) | 1)))
                        } else {
                            var v13 uint16 = 0
                            if (int(v13) >= -2147483648) {
                                v14 := GenS1{F0: 42}
                                fmt.Println(fi11, v8, v8)
                                _ = v14
                            } else {
                                for fi15 := 0; fi15 < 16; fi15++ {
                                    var v16 int32 = -1
                                    if false {
                                        v17 := GenS1{F0: 0}
                                        switch ("foo bar" + ("" + "a")) {
                                        case "a":
                                            fmt.Println(v16, v8, v16)
                                        case "":
                                            var v18 int = 83
                                            _ = v18
                                        case "foo bar":
                                            v13 += (^(v13))
                                            if false {
                                                _ = len(v10)
                                            } else {
                                                v7 += 660
                                                var v19 float32 = -2.25
                                                fmt.Println(fi11, v13)
                                                _ = v19
                                            }
                                            v13 = (^(v13))
                                        }
                                        var v20 uint8 = 0
                                        _ = v17
                                        _ = v20
                                    } else {
                                        _ = len(v10)
                                    }
                                    var v21 int8 = 100
                                    var v22 int = fi15
                                    fmt.Println(v16, v21, v22)
                                    _ = fi15
                                    _ = v16
                                    _ = v21
                                    _ = v22
                                }
                                v8 = (v8 % ((42) | 1))
                            }
                            var v23 uint64 = (v8 & (^(v8)))
                            _ = v13
                            _ = v23
                        }
                        fmt.Println(v8, v8, v7)
                    }
                    switch "foo bar" {
                    case "hello":
                        v24 := GenS1{F0: 90}
                        _ = v24
                    }
                    _ = fi11
                }
            } else {
                var v25 int = -84
                var v26 float32 = 1.5
                _ = v25
                _ = v26
            }
            _ = v10
        }
        _ = v8
    }
    fmt.Println(v7)
    _ = v7
    return 632
}
func genFn1() uint64 {
    v27 := GenS1{F0: 1}
    var v28 uint16 = 2
    fmt.Println(v28)
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := []int{fi29, fi29}
        v31 := GenS1{F0: -739}
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 uint16 = (v28 % ((v28) | 1))
    _ = v27
    _ = v28
    _ = v32
    return 0
}

func main() {
    var v1 uint64 = 0
    v2 := []string{"foo bar", ((("\n\t" + "a") + "a") + "")}
    v3 := []string{("\n\t" + "foo bar"), ""}
    var v4 bool = (-41 >= 127)
    var v5 int = -2147483648
    var v6 uint8 = 128
    fmt.Println(v6, v4)
    v2 = append(v2, ("hello" + ("a" + "\n\t")))
    if v4 {
        v33 := GenS2{F0: -140}
        v6 = (v6 / ((181) | 1))
        _ = v33
    }
    if v4 {
        var v34 int8 = 1
        v35 := GenS2{F0: 799}
        var v36 uint = 4294967295
        _ = v34
        _ = v35
        _ = v36
    }
    v37 := GenS3{F0: ("a" + "hello"), F1: 543, F2: (v1 >> ((4294967295) & 63))}
    fmt.Println(v1, v6)
    v5 -= 127
    if v4 {
        v1 += v1
        var v38 float64 = 0.0
        _ = v38
    } else {
        var v39 int32 = 2147483647
        var v40 uint64 = v1
        v41 := GenS1{F0: 1}
        _ = v39
        _ = v40
        _ = v41
    }
    v1 = (^(v1))
    for fi42 := 0; fi42 < 1; fi42++ {
        for fi43 := 0; fi43 < 17; fi43++ {
            if len(v2) > 0 {
                v2[0] = ("hello" + ("foo bar" + "\n\t"))
            }
            _ = fi43
        }
        var v44 string = "foo bar"
        v45 := GenS4{F0: 1000}
        _ = fi42
        _ = v44
        _ = v45
    }
    var v46 int8 = -1
    var v47 int8 = (v46 & 42)
    if false {
        var v48 uint = 255
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v49 := GenS2{F0: 9223372036854775807}
        _ = v48
        _ = v49
    } else {
        v2 = append(v2, "foo bar")
        var v50 uint64 = genFn1()
        _ = v50
    }
    v6 = (v6 << ((255) & 63))
    switch "" {
    case "":
        var v51 uint32 = 0
        _ = v51
    case "foo bar":
        for fi52 := 0; fi52 < 14; fi52++ {
            _ = len(v3)
            switch ((("a" + ("foo bar" + ("a" + ""))) + ("hello" + "foo bar")) + "foo bar") {
            case "foo bar":
                if (true || (!true)) {
                    v47 = v47
                    switch fi52 {
                    case -39:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 42:
                        if ("" <= "\n\t") {
                            v53 := GenS4{F0: 42}
                            _ = v53
                        } else {
                            if v4 {
                                for fi54 := 0; fi54 < 17; fi54++ {
                                    v55 := GenS2{F0: -1}
                                    _ = fi54
                                    _ = v55
                                }
                            } else {
                                v4 = (!v4)
                                if v4 {
                                    if (genFn0() >= 65535) {
                                        if ((-2.25 >= ((((0.1 + 3.14159) / 1.0) * 1e10) / 1e10)) || v4) {
                                            var v56 bool = v4
                                            _ = v56
                                        }
                                        fmt.Println(v47)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    if true {
                                        _ = len(v3)
                                        switch "hello" {
                                        case "foo bar":
                                            v57 := GenS5{F0: 533, F1: 0, F2: (^(v1))}
                                            _ = v57
                                        case "":
                                            if len(v2) > 0 {
                                                v2[0] = ("foo bar" + "\n\t")
                                            }
                                        case "\n\t":
                                            v1 += genFn1()
                                        default:
                                            v4 = false
                                        }
                                        v5 += (v5 * 42)
                                    }
                                    var v58 bool = (!false)
                                    _ = v58
                                } else {
                                    for fi59 := 0; fi59 < 9; fi59++ {
                                        v60 := GenS6{F0: 255}
                                        v5 = (fi59 / ((-2147483648) | 1))
                                        switch -1 {
                                        case 128:
                                            switch (^(v5)) {
                                            case 32767:
                                                var v61 uint = 255
                                                _ = v61
                                            case 128:
                                                if ((((!true) && v4) && (32767 > 1)) || (1 == -932)) {
                                                    var v62 int = v5
                                                    v1 = v1
                                                    var v63 uint8 = v6
                                                    _ = v62
                                                    _ = v63
                                                } else {
                                                    v64 := GenS4{F0: 32767}
                                                    var v65 int = (^(fi52))
                                                    _ = v64
                                                    _ = v65
                                                }
                                            case 4294967295:
                                                switch "\n\t" {
                                                case "":
                                                    if v4 {
                                                        if ((65535 >= 0) || v4) {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            var v66 bool = false
                                                            _ = v66
                                                        }
                                                    }
                                                case "\n\t":
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v67 string = (("hello" + "hello") + "")
                                                        _ = v67
                                                    case "hello":
                                                        fmt.Println(v4)
                                                        if (v4 || (42 >= 65535)) {
                                                            _ = len(v3)
                                                            var v68 int8 = 42
                                                            _ = v68
                                                        }
                                                        v69 := GenS6{F0: 4294967295}
                                                        _ = v69
                                                    case "foo bar":
                                                        switch ("" + "hello") {
                                                        case "foo bar":
                                                            if (v4 || (!false)) {
                                                                if true {
                                                                    if (!v4) {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = ("foo bar" + "foo bar")
                                                                        }
                                                                    }
                                                                } else {
                                                                    fmt.Println(fi52, v4, v1)
                                                                    fmt.Println(v5, v5)
                                                                }
                                                                if ((1e10 > 1.5) || v4) {
                                                                    v47 = (v46 & v46)
                                                                }
                                                            }
                                                        case "\n\t":
                                                            if (!(v4 && v4)) {
                                                                var v70 int32 = 0
                                                                if (v4 || true) {
                                                                    var v71 uint32 = 4294967295
                                                                    fmt.Println(fi52, fi52, fi59)
                                                                    _ = v71
                                                                }
                                                                _ = v70
                                                            } else {
                                                                var v72 string = "foo bar"
                                                                _ = v72
                                                            }
                                                        case "hello":
                                                            if v4 {
                                                                v4 = (42 == 4294967295)
                                                            }
                                                        default:
                                                            if v4 {
                                                                fmt.Println(v4)
                                                                var v73 bool = v4
                                                                _ = v73
                                                            } else {
                                                                var v74 uint64 = (v1 << ((4294967295) & 63))
                                                                _ = len(v2)
                                                                _ = v74
                                                            }
                                                            v75 := GenS7{F0: -1, F1: 42}
                                                            var v76 bool = v4
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                    }
                                                }
                                            default:
                                                fmt.Println(v6, v1, v6)
                                            }
                                        case 42:
                                            fmt.Println(v47, v1, v46)
                                        case 2147483647:
                                            v77 := GenS8{F0: "\n\t"}
                                            _ = v77
                                        }
                                        _ = fi59
                                        _ = v60
                                    }
                                }
                                v6 -= uint8(v1)
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if true {
                                v78 := GenS5{F0: 2147483648, F1: 2147483647, F2: 4294967295}
                                if (!(v6 < (^(v6)))) {
                                    v6 = (v6 & v6)
                                }
                                var v79 uint64 = (v1 ^ (^(v1)))
                                _ = v78
                                _ = v79
                            }
                        }
                    }
                }
            case "\n\t":
                v5 += (fi52 % ((fi52) | 1))
            case "hello":
                v46 = 127
            default:
                switch v5 {
                case 544:
                    v2 = append(v2, (("foo bar" + "\n\t") + (("" + ((("hello" + ("\n\t" + ("hello" + "a"))) + ("hello" + (("" + "a") + ("hello" + ((("" + "a") + "\n\t") + "foo bar"))))) + ((("\n\t" + "\n\t") + (((("" + (("hello" + "hello") + ((("" + ("\n\t" + "")) + (("a" + "hello") + "foo bar")) + "foo bar"))) + (("\n\t" + ("hello" + "\n\t")) + ((((("" + "") + "foo bar") + ("" + "foo bar")) + "a") + "foo bar"))) + (((("a" + "") + ("a" + ("hello" + ((("hello" + (((("hello" + "a") + "hello") + ("foo bar" + ("foo bar" + "a"))) + ("a" + "a"))) + "") + "\n\t")))) + "hello") + "foo bar")) + "hello")) + "hello"))) + "hello")))
                case 2147483647:
                    v80 := GenS9{F0: 18446744073709551615}
                    _ = v80
                case 32767:
                    if (v4 && (true || v4)) {
                        v81 := GenS3{F0: "foo bar", F1: 18446744073709551615, F2: (v1 + 214)}
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                        _ = v81
                    }
                default:
                    switch v5 {
                    case 921:
                        for fi82 := 0; fi82 < 9; fi82++ {
                            var v83 uint64 = v1
                            switch v5 {
                            case 32767:
                                if v4 {
                                    if ((1000 >= -32768) && (!false)) {
                                        var v84 int = fi52
                                        fmt.Println(v47, fi52, v47)
                                        v3 = append(v3, ("foo bar" + "\n\t"))
                                        _ = v84
                                    } else {
                                        v5 = fi52
                                        v6 += (^(v6))
                                        v5 = 128
                                    }
                                    v6 -= (v6 >> ((1) & 63))
                                    var v85 int16 = -490
                                    _ = v85
                                }
                            default:
                                fmt.Println(fi52, v83)
                            }
                            if false {
                                v86 := GenS5{F0: 1, F1: 42, F2: 2}
                                fmt.Println(v1, v83, fi82)
                                _ = v86
                            }
                            _ = fi82
                            _ = v83
                        }
                    case 9223372036854775807:
                        for fi87 := 0; fi87 < 12; fi87++ {
                            v47 = v46
                            _ = fi87
                        }
                    }
                }
            }
            switch fi52 {
            case -9223372036854775808:
                v88 := GenS10{F0: (true && ((!true) || (-1 >= -355))), F1: v1, F2: (v5 % ((fi52) | 1))}
                _ = v88
            default:
                for fi89 := 0; fi89 < 14; fi89++ {
                    var v90 string = ("" + (("foo bar" + "hello") + ""))
                    _ = fi89
                    _ = v90
                }
            }
            _ = fi52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
