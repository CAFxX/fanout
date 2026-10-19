package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 int; F1 uint8; F2 int64 }
type GenS9 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 0
    var v3 uint8 = 1
    v4 := []string{"", "", (("a" + ("foo bar" + "")) + ((("" + "hello") + "a") + ("foo bar" + "a")))}
    v3 = 0
    for fi5 := 0; fi5 < 14; fi5++ {
        if true {
            v2 = v2
            v6 := GenS1{F0: int64(v3)}
            v7 := GenS1{F0: -9223372036854775808}
            _ = v6
            _ = v7
        }
        _ = fi5
    }
    v2 = v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        switch "" {
        case "\n\t":
            v3 = 101
        case "hello":
            for fi9 := 0; fi9 < 4; fi9++ {
                switch "\n\t" {
                case "\n\t":
                    if false {
                        if (-935 >= -871) {
                            switch "" {
                            case "a":
                                v3 -= (v3 + 1)
                            case "\n\t":
                                switch "hello" {
                                case "foo bar":
                                    if (0 <= 2147483648) {
                                        v2 = (^(v2))
                                    } else {
                                        fmt.Println(v3)
                                        var v10 int64 = -1
                                        v1[("" + "\n\t")] = ((("foo bar" + "foo bar") + "a") + ("" + "a"))
                                        _ = v10
                                    }
                                case "hello":
                                    v3 = (v3 >> ((490) & 63))
                                case "":
                                    if (!true) {
                                        v11 := GenS2{F0: -32768}
                                        v3 += v3
                                        _ = v11
                                    } else {
                                        var v12 uint8 = (v3 ^ (v3 + (v3 ^ v3)))
                                        _ = v12
                                    }
                                default:
                                    v3 = (^(v3))
                                    v13 := GenS1{F0: -9223372036854775808}
                                    _ = len(v1)
                                    _ = v13
                                }
                            default:
                                if false {
                                    var v14 int32 = 2147483647
                                    var v15 float32 = 3.14159
                                    v15 -= v15
                                    _ = v14
                                    _ = v15
                                }
                            }
                        } else {
                            v2 = v2
                        }
                        var v16 bool = (-103 >= 42)
                        _ = v16
                    } else {
                        v3 -= v3
                        switch (^(fi9)) {
                        case 4294967295:
                            var v17 int32 = 42
                            _ = v17
                        default:
                            fmt.Println(v2, fi8)
                        }
                        switch "foo bar" {
                        case "\n\t":
                            v3 -= v3
                        case "a":
                            switch 467 {
                            case -631:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 32767:
                                switch ("\n\t" + ("\n\t" + (("" + "a") + ("hello" + (("a" + "hello") + ""))))) {
                                case "foo bar":
                                    var v18 float32 = float32(-2.25)
                                    _ = v18
                                case "hello":
                                    var v19 int8 = -126
                                    _ = v19
                                default:
                                    if (!false) {
                                        v20 := GenS2{F0: 0}
                                        var v21 int = fi9
                                        var v22 uint8 = v3
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    } else {
                                        var v23 float32 = float32(1e10)
                                        if false {
                                            switch -947 {
                                            case 9223372036854775807:
                                                if false {
                                                    v2 += v2
                                                } else {
                                                    var v24 uint32 = 124
                                                    _ = v24
                                                }
                                            case -186:
                                                var v25 float32 = 0.1
                                                if false {
                                                    if (-956 < 1) {
                                                        var v26 uint8 = (^(v3))
                                                        _ = v26
                                                    } else {
                                                        v27 := GenS2{F0: -460}
                                                        fmt.Println(v2, v25)
                                                        v28 := GenS2{F0: -312}
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                }
                                                _ = v25
                                            case -1:
                                                if len(v4) > 0 {
                                                    v4[0] = "foo bar"
                                                }
                                                v29 := v1["foo bar"]
                                                fmt.Println(v3, fi9)
                                                _ = v29
                                            default:
                                                v30 := GenS2{F0: -25}
                                                _ = v30
                                            }
                                            var v31 bool = ((!false) || false)
                                            switch ((((("" + "foo bar") + "foo bar") + "") + "") + ("" + ("a" + ("\n\t" + (("a" + ((("hello" + "foo bar") + "hello") + (("foo bar" + "hello") + (((("" + ("foo bar" + "foo bar")) + ((("" + ("foo bar" + ((("\n\t" + "foo bar") + "\n\t") + "a"))) + ("a" + "hello")) + "\n\t")) + "") + "hello")))) + "foo bar"))))) {
                                            case "a":
                                                var v32 uint16 = 36
                                                _ = v32
                                            }
                                            _ = v31
                                        } else {
                                            var v33 int64 = -666
                                            v33 -= -274
                                            _ = v33
                                        }
                                        v23 = (((((v23 * v23) * v23) * 0.0) * 1e10) * (float32(3.14159) / 1.0))
                                        v3 -= v3
                                        _ = v23
                                    }
                                }
                            case 1:
                                if (!false) {
                                    if false {
                                        if (true || false) {
                                            var v34 uint8 = (v3 + 185)
                                            _ = v34
                                        }
                                        var v35 uint16 = (v2 / ((856) | 1))
                                        if (!true) {
                                            if (18446744073709551615 >= 368) {
                                                if true {
                                                    if ((v35 - v2) == v35) {
                                                        var v36 float64 = float64(0.0)
                                                        v37 := GenS1{F0: 0}
                                                        _ = v36
                                                        _ = v37
                                                    } else {
                                                        v35 = (v2 + uint16(v3))
                                                    }
                                                }
                                            } else {
                                                v35 -= v35
                                            }
                                        }
                                        if true {
                                            if true {
                                                v38 := GenS3{F0: 606}
                                                v35 = (^(v2))
                                                _ = v38
                                            } else {
                                                var v39 float64 = (((1.5 * (0.0 + (-2.25 + (-2.25 + 3.14159)))) / 1.5) / 1e10)
                                                _ = v39
                                            }
                                            fmt.Println(v35, v2, fi8)
                                        } else {
                                            var v40 int64 = -9223372036854775808
                                            _ = v40
                                        }
                                        _ = v35
                                    }
                                    var v41 float64 = 3.14159
                                    _ = len(v1)
                                    if (!(0 > 645)) {
                                        _ = len(v1)
                                        v42 := GenS3{F0: 2}
                                        _ = v42
                                    } else {
                                        v43 := GenS3{F0: 439}
                                        var v44 int64 = -9223372036854775808
                                        if (!true) {
                                            switch fi9 {
                                            case 32767:
                                                _ = len(v1)
                                            case 943:
                                                switch -949 {
                                                case 2147483647:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                default:
                                                    var v45 uint32 = 0
                                                    _ = v45
                                                }
                                            default:
                                                _ = len(v1)
                                            }
                                        }
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v41
                                }
                                fmt.Println(v2, fi9)
                                v3 -= 127
                                var v46 uint8 = 42
                                _ = v46
                            }
                        default:
                            switch "foo bar" {
                            case "\n\t":
                                if true {
                                    var v47 bool = ((v2 | (v2 ^ (v2 - 0))) != 405)
                                    _ = v47
                                }
                            }
                        }
                    }
                default:
                    v48 := GenS2{F0: -1}
                    _ = v48
                }
                var v49 float32 = (1e10 / (float32(-2.25) - (0.0 - (1e10 * ((float32((float64(1e10) + 0.0)) / 3.14159) + (((((0.0 * -2.25) * -2.25) - -2.25) / 1e10) * 0.0))))))
                _ = fi9
                _ = v49
            }
        }
        _ = fi8
    }
    fmt.Println(v2, v3)
    v2 += 716
    if false {
        for fi50 := 0; fi50 < 17; fi50++ {
            fmt.Println(v3, v2, v3)
            _ = fi50
        }
        v1[("a" + (("foo bar" + (("" + "a") + "\n\t")) + "hello"))] = "a"
        var v51 bool = (!(!(!(!(1 != 42)))))
        _ = v51
    } else {
        v3 = (v3 ^ (v3 >> ((65535) & 63)))
    }
    var v52 float64 = (0.0 * float64(1e10))
    var v53 int32 = -927
    for fi54 := 0; fi54 < 5; fi54++ {
        v55 := GenS2{F0: 1}
        var v56 bool = true
        _ = fi54
        _ = v55
        _ = v56
    }
    v57 := GenS4{F0: 802}
    delete(v1, "foo bar")
    if false {
        var v58 float32 = ((float32((v52 + 3.14159)) - 1e10) / ((1e10 + (-2.25 * 1.5)) * -2.25))
        switch 255 {
        case 9223372036854775807:
            v59 := GenS5{F0: uint(v3), F1: 4294967295}
            _ = v59
        case -229:
            v1["\n\t"] = "hello"
        default:
            for fi60 := 0; fi60 < 1; fi60++ {
                if (false && true) {
                    if (3.14159 >= ((0.0 * float32((v52 / (v52 / 0.1)))) * -2.25)) {
                        fmt.Println(v58, v58, v52)
                        if (-9223372036854775808 > -9223372036854775808) {
                            v2 += (v2 / (((v2 / (((v2 << ((4294967295) & 63))) | 1))) | 1))
                            var v61 uint16 = (^(v2))
                            _ = v61
                        } else {
                            switch 128 {
                            case 1:
                                v52 = float64(1e10)
                            case 9223372036854775807:
                                switch -679 {
                                case 674:
                                    v2 -= 0
                                }
                            default:
                                var v62 int16 = 42
                                _ = v62
                            }
                        }
                    } else {
                        for fi63 := 0; fi63 < 11; fi63++ {
                            v64 := GenS6{F0: "a"}
                            var v65 int64 = 1
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                        if false {
                            v53 += v53
                            for fi66 := 0; fi66 < 3; fi66++ {
                                v67 := v1["a"]
                                _ = len(v1)
                                if true {
                                    var v68 int32 = v53
                                    v68 -= (v53 << ((65535) & 63))
                                    v69 := GenS2{F0: -32768}
                                    _ = v68
                                    _ = v69
                                } else {
                                    var v70 int32 = v53
                                    _ = v70
                                }
                                _ = fi66
                                _ = v67
                            }
                            if false {
                                if true {
                                    for fi71 := 0; fi71 < 20; fi71++ {
                                        switch fi71 {
                                        case 9223372036854775807:
                                            fmt.Println(fi60)
                                        case -261:
                                            if (32767 >= -985) {
                                                var v72 int = 370
                                                switch (fi71 + (fi71 << ((729) & 63))) {
                                                case -9223372036854775808:
                                                    _ = len(v1)
                                                case 474:
                                                    v73 := GenS7{F0: "\n\t", F1: 128}
                                                    _ = v73
                                                case 32768:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v2 -= v2
                                                    default:
                                                        delete(v1, ("" + "a"))
                                                    }
                                                }
                                                fmt.Println(v72, v53, v72)
                                                var v74 bool = true
                                                _ = v72
                                                _ = v74
                                            }
                                        case 915:
                                            v2 -= v2
                                        }
                                        _ = fi71
                                    }
                                    v75 := GenS3{F0: 255}
                                    var v76 uint64 = 4294967295
                                    _ = v75
                                    _ = v76
                                } else {
                                    var v77 uint64 = 0
                                    v2 = v2
                                    _ = v77
                                }
                                if (!(v58 > v58)) {
                                    switch (fi60 + (^(fi60))) {
                                    case -1:
                                        var v78 uint = 838
                                        _ = v78
                                    case 1:
                                        switch 42 {
                                        case 4294967295:
                                            var v79 uint64 = 4294967295
                                            _ = v79
                                        case 2147483647:
                                            for fi80 := 0; fi80 < 13; fi80++ {
                                                v81 := GenS2{F0: 0}
                                                var v82 string = ((("" + (("" + ("hello" + ("hello" + "\n\t"))) + (("\n\t" + "a") + ("hello" + ("" + ""))))) + "foo bar") + ("hello" + ("\n\t" + "foo bar")))
                                                v52 += 1.5
                                                _ = fi80
                                                _ = v81
                                                _ = v82
                                            }
                                        default:
                                            var v83 int8 = 1
                                            _ = v83
                                        }
                                    }
                                    v84 := GenS4{F0: 414}
                                    v85 := GenS7{F0: "", F1: v3}
                                    _ = v84
                                    _ = v85
                                } else {
                                    v1["\n\t"] = "foo bar"
                                }
                            }
                        }
                    }
                    var v86 uint64 = 71
                    for fi87 := 0; fi87 < 10; fi87++ {
                        fmt.Println(v52)
                        v88 := GenS8{F0: 2147483647, F1: v3, F2: 0}
                        _ = len(v1)
                        _ = fi87
                        _ = v88
                    }
                    _ = v86
                }
                var v89 float32 = v58
                v90 := GenS9{F0: -45}
                _ = fi60
                _ = v89
                _ = v90
            }
        }
        _ = v58
    }
    fmt.Println(v53, v53, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v57
}
