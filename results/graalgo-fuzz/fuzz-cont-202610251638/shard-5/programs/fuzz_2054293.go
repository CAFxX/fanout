package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }
type GenS4 struct { F0 int16; F1 bool; F2 uint16 }
type GenS5 struct { F0 int; F1 uint32; F2 uint32 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 int; F1 uint; F2 string }
type GenS10 struct { F0 int16; F1 int32; F2 uint64 }
type GenS11 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{9223372036854775807, -796, -9223372036854775808}
    var v3 float64 = -2.25
    for fi4 := 0; fi4 < 15; fi4++ {
        switch (("\n\t" + "\n\t") + "a") {
        case "foo bar":
            var v5 bool = false
            _ = v5
        }
        var v6 int = -633
        _ = fi4
        _ = v6
    }
    if (-367 <= 32767) {
        v7 := GenS1{F0: ("hello" + ("" + "")), F1: (!false), F2: 1}
        var v8 uint16 = 0
        switch -9223372036854775808 {
        case -481:
            for fi9 := 0; fi9 < 9; fi9++ {
                v3 = v3
                fmt.Println(v3, v3, fi9)
                _ = fi9
            }
        case 47:
            delete(v1, ("" + "\n\t"))
        }
        _ = v7
        _ = v8
    } else {
        if len(v2) > 0 {
            v2[0] = 127
        }
        v3 -= 1.5
    }
    if (4294967295 < 2) {
        if (1 <= -9223372036854775808) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi10 := 0; fi10 < 15; fi10++ {
                delete(v1, "hello")
                for fi11 := 0; fi11 < 17; fi11++ {
                    fmt.Println(fi10)
                    _ = fi11
                }
                _ = fi10
            }
            var v12 uint8 = 2
            _ = v12
        } else {
            v13 := v1[""]
            for fi14 := 0; fi14 < 12; fi14++ {
                if (-125 >= 0) {
                    v3 += v3
                    v15 := GenS1{F0: (((("hello" + "foo bar") + ("foo bar" + "hello")) + "a") + ("hello" + "a")), F1: (-128 >= -121), F2: 42}
                    _ = v15
                } else {
                    if ((v13 + 131) <= v13) {
                        var v16 int16 = 0
                        _ = v16
                    } else {
                        v3 += ((float64(3.14159) / 3.14159) - v3)
                        var v17 int8 = 100
                        var v18 int64 = 1
                        _ = v17
                        _ = v18
                    }
                    delete(v1, "a")
                    for fi19 := 0; fi19 < 8; fi19++ {
                        if false {
                            if ("\n\t" < "a") {
                                if false {
                                    fmt.Println(v13, v3, fi19)
                                    if len(v2) > 0 {
                                        v2[0] = 0
                                    }
                                    var v20 string = "hello"
                                    var v21 int8 = -1
                                    switch (fi14 * fi14) {
                                    case -709:
                                        switch v20 {
                                        case "hello":
                                            v13 += (v13 ^ (v13 & v13))
                                        case "foo bar":
                                            v1[("hello" + ("foo bar" + v20))] = (^(v13))
                                        case "":
                                            v22 := v1[v20]
                                            _ = v22
                                        }
                                    case -1:
                                        fmt.Println(fi19)
                                    }
                                    _ = v20
                                    _ = v21
                                } else {
                                    switch fi19 {
                                    case 1:
                                        switch (("hello" + "a") + "") {
                                        case "hello":
                                            v13 -= (v13 * 4294967295)
                                        case "":
                                            if len(v2) > 0 {
                                                v2[0] = fi14
                                            }
                                        case "foo bar":
                                            v23 := GenS1{F0: "\n\t", F1: (898 < -9223372036854775808), F2: -440}
                                            _ = v23
                                        default:
                                            v13 += (v13 * (v13 % ((65535) | 1)))
                                        }
                                    }
                                }
                                var v24 int64 = 0
                                var v25 int16 = -32768
                                _ = v24
                                _ = v25
                            }
                        }
                        if false {
                            v26 := GenS2{F0: 0, F1: 65535}
                            var v27 int64 = 42
                            _ = v26
                            _ = v27
                        } else {
                            v13 = (^(v13))
                        }
                        if ((^(v13)) >= 1) {
                            v3 = float64(1.5)
                            var v28 bool = true
                            if v28 {
                                fmt.Println(v13)
                            } else {
                                v29 := GenS2{F0: 397, F1: 952}
                                var v30 int16 = 0
                                v3 = (-2.25 / ((((v3 * float64(3.14159)) + (v3 * v3)) + float64(-2.25)) / 1e10))
                                _ = v29
                                _ = v30
                            }
                            _ = v28
                        }
                        _ = fi19
                    }
                }
                _ = fi14
            }
            _ = v13
        }
        fmt.Println(v3)
        v3 -= v3
    } else {
        if (!false) {
            v31 := GenS3{F0: false, F1: ((("" + (("foo bar" + ("" + (("hello" + "\n\t") + ("hello" + "")))) + "a")) + ((("hello" + ("\n\t" + ("" + ((("a" + "a") + ((("" + "\n\t") + ("" + "\n\t")) + ("a" + "\n\t"))) + ("foo bar" + "hello"))))) + "foo bar") + "hello")) < ""), F2: 208}
            if false {
                v3 = (0.0 * -2.25)
                for fi32 := 0; fi32 < 2; fi32++ {
                    if (!false) {
                        if false {
                            if true {
                                var v33 int8 = -21
                                if true {
                                    for fi34 := 0; fi34 < 1; fi34++ {
                                        v3 = float64((3.14159 * (0.0 + -2.25)))
                                        if len(v2) > 0 {
                                            v2[0] = fi34
                                        }
                                        v35 := GenS4{F0: 1000, F1: (true && true), F2: 42}
                                        _ = fi34
                                        _ = v35
                                    }
                                    v36 := GenS2{F0: 0, F1: 65535}
                                    _ = v36
                                }
                                if (!(!false)) {
                                    switch 128 {
                                    case 2147483647:
                                        var v37 int64 = -34
                                        _ = v37
                                    case 4294967295:
                                        var v38 uint64 = 396
                                        _ = v38
                                    }
                                    switch fi32 {
                                    case -2147483648:
                                        if true {
                                            fmt.Println(v3)
                                            if false {
                                                v39 := GenS5{F0: 1, F1: 993, F2: 0}
                                                if false {
                                                    if (!false) {
                                                        delete(v1, "hello")
                                                        v33 = (-(v33))
                                                        for fi40 := 0; fi40 < 10; fi40++ {
                                                            if false {
                                                                var v41 uint = 42
                                                                if true {
                                                                    var v42 int64 = -1
                                                                    _ = v42
                                                                } else {
                                                                    v41 = (v41 + (v41 % ((v41) | 1)))
                                                                }
                                                                if false {
                                                                    _ = len(v1)
                                                                }
                                                                if true {
                                                                    var v43 int = (fi32 - (fi32 & fi32))
                                                                    _ = v43
                                                                } else {
                                                                    v2 = append(v2, 445)
                                                                    if false {
                                                                        v44 := GenS3{F0: (957 == -9223372036854775808), F1: (-782 >= 9223372036854775807), F2: 42}
                                                                        var v45 string = "hello"
                                                                        _ = v44
                                                                        _ = v45
                                                                    }
                                                                    v46 := GenS6{F0: (fi32 - (fi32 % (((fi32 + 2147483647)) | 1))), F1: (v33 * (v33 << ((2) & 63)))}
                                                                    _ = v46
                                                                }
                                                                _ = v41
                                                            }
                                                            switch ("a" + "foo bar") {
                                                            case "foo bar":
                                                                v1["\n\t"] = 4294967295
                                                            case "a":
                                                                if true {
                                                                    v47 := GenS4{F0: 407, F1: (!false), F2: 336}
                                                                    var v48 uint32 = 1
                                                                    fmt.Println(v33, v3)
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                            case "hello":
                                                                v33 = v33
                                                            }
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v3 += (v3 - float64((((1e10 * (0.0 + float32(v3))) / 0.1) + 0.0)))
                                                            case "\n\t":
                                                                if false {
                                                                    v3 -= ((v3 + (1.5 * v3)) * v3)
                                                                    if (!(false && false)) {
                                                                        v49 := GenS2{F0: 229, F1: 65535}
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        _ = v49
                                                                    }
                                                                }
                                                            case "a":
                                                                var v50 int32 = 42
                                                                _ = v50
                                                            }
                                                            if ((!(837 > 934)) && ((v33 & 0) <= (v33 - (v33 - v33)))) {
                                                                var v51 float32 = (1.5 * (3.14159 * (3.14159 - 1e10)))
                                                                _ = v51
                                                            } else {
                                                                var v52 int = fi40
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v53 := GenS1{F0: "hello", F1: (true && false), F2: -633}
                                                                _ = v52
                                                                _ = v53
                                                            }
                                                            _ = fi40
                                                        }
                                                    } else {
                                                        v54 := GenS6{F0: (^(fi32)), F1: v33}
                                                        _ = v54
                                                    }
                                                    v55 := GenS7{F0: 195, F1: 2147483647, F2: (v33 << ((18446744073709551615) & 63))}
                                                    _ = v55
                                                } else {
                                                    fmt.Println(v3, v33)
                                                }
                                                _ = v39
                                            }
                                            for fi56 := 0; fi56 < 16; fi56++ {
                                                v57 := GenS3{F0: (fi56 != (-(fi56))), F1: (true || false), F2: 128}
                                                var v58 uint16 = 2
                                                _ = fi56
                                                _ = v57
                                                _ = v58
                                            }
                                        }
                                    case -439:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case -9223372036854775808:
                                        var v59 uint64 = 0
                                        _ = v59
                                    }
                                }
                                _ = v33
                            }
                            v60 := GenS2{F0: 842, F1: 255}
                            var v61 uint8 = 42
                            _ = v60
                            _ = v61
                        } else {
                            if (int8(fi32) < 1) {
                                if true {
                                    for fi62 := 0; fi62 < 14; fi62++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v63 := GenS8{F0: "a", F1: "foo bar"}
                                        _ = fi62
                                        _ = v63
                                    }
                                    switch fi32 {
                                    case 2147483647:
                                        v3 -= 3.14159
                                    case 32767:
                                        if true {
                                            var v64 string = ""
                                            v65 := v1["\n\t"]
                                            var v66 uint16 = 42
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        }
                                    case -1:
                                        v3 = float64(-2.25)
                                    }
                                    fmt.Println(fi32, v3)
                                }
                            } else {
                                v67 := GenS9{F0: (fi32 | 32768), F1: 2, F2: ("hello" + "a")}
                                _ = len(v2)
                                fmt.Println(v3, v3)
                                _ = v67
                            }
                            v3 += (1.5 + 3.14159)
                        }
                    }
                    v3 = (0.0 * float64(1e10))
                    v68 := v1[(("a" + ((((((("a" + "foo bar") + "\n\t") + ("hello" + "a")) + "") + ("hello" + ("" + ("a" + "\n\t")))) + "") + ("a" + "foo bar"))) + "foo bar")]
                    _ = fi32
                    _ = v68
                }
            } else {
                for fi69 := 0; fi69 < 6; fi69++ {
                    v70 := GenS6{F0: fi69, F1: -56}
                    _ = fi69
                    _ = v70
                }
            }
            if (0 != 766) {
                fmt.Println(v3)
                var v71 uint16 = 2
                _ = v71
            } else {
                v3 = v3
                if false {
                    _ = len(v1)
                }
            }
            _ = v31
        } else {
            v3 = (float64(0.1) * 0.1)
            v3 = v3
        }
        for fi72 := 0; fi72 < 18; fi72++ {
            for fi73 := 0; fi73 < 3; fi73++ {
                v3 += v3
                _ = fi73
            }
            switch fi72 {
            case 127:
                v3 = 0.0
            case 471:
                if (false || true) {
                    for fi74 := 0; fi74 < 17; fi74++ {
                        v3 -= v3
                        v75 := GenS10{F0: 42, F1: 2147483647, F2: 231}
                        _ = fi74
                        _ = v75
                    }
                    v3 = v3
                    var v76 string = "a"
                    _ = v76
                }
            default:
                var v77 int8 = 81
                _ = v77
            }
            if (876 == 2) {
                v3 = (((0.1 - (v3 - v3)) / ((1.5 + 1e10) + 1e10)) + v3)
            }
            _ = fi72
        }
    }
    v3 = v3
    var v78 int16 = -476
    delete(v1, "\n\t")
    var v79 uint = 0
    var v80 int16 = (^(v78))
    if true {
        switch ("a" + ("" + "hello")) {
        case "a":
            for fi81 := 0; fi81 < 8; fi81++ {
                var v82 float32 = (float32(((v3 - (v3 * ((0.0 / v3) + v3))) * (v3 - 1e10))) * -2.25)
                v83 := GenS11{F0: 42}
                _ = fi81
                _ = v82
                _ = v83
            }
        default:
            var v84 int = 128
            _ = v84
        }
        v85 := v1[(("" + "") + "foo bar")]
        _ = v85
    }
    for fi86 := 0; fi86 < 17; fi86++ {
        if false {
            switch int(v79) {
            case 1:
                var v87 string = "\n\t"
                _ = v87
            case 4294967295:
                var v88 float32 = -2.25
                _ = v88
            }
        } else {
            delete(v1, "\n\t")
        }
        var v89 int16 = -652
        _ = fi86
        _ = v89
    }
    v90 := v1[("foo bar" + "hello")]
    v90 -= (v90 / ((v90) | 1))
    var v91 int16 = v80
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v78
    _ = v79
    _ = v80
    _ = v90
    _ = v91
}
