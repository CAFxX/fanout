package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 int64; F1 string; F2 int }
type GenS6 struct { F0 int8; F1 int64; F2 int16 }
func genFn0(p0 uint32, p1 int16) int64 {
    var v6 float32 = (0.1 + (-2.25 * 3.14159))
    var v7 int8 = 100
    if (false && false) {
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := []int{}
            _ = fi8
            _ = v9
        }
    } else {
        var v10 uint64 = 1
        fmt.Println(v7, v6)
        v11 := GenS1{F0: 127, F1: ("hello" + "a")}
        _ = v10
        _ = v11
    }
    fmt.Println(p0, v7, v6)
    v7 = (^(v7))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1() bool {
    v12 := []int{502, -533, -9223372036854775808}
    var v13 int64 = -9223372036854775808
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 int8 = 100
        var v16 uint32 = 343
        if (-2.25 < float32(0.0)) {
            for fi17 := 0; fi17 < 8; fi17++ {
                fmt.Println(fi17, fi17, fi17)
                v15 = (v15 >> ((0) & 63))
                v13 += v13
                if len(v12) > 0 {
                    v12[0] = 1
                }
                switch "foo bar" {
                case "foo bar":
                    var v18 float64 = (float64(0.1) * 1e10)
                    _ = v18
                }
                _ = fi17
            }
        } else {
            switch 4294967295 {
            case -181:
                for fi19 := 0; fi19 < 7; fi19++ {
                    v13 -= v13
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    var v20 string = ""
                    _ = fi19
                    _ = v20
                }
            case 9223372036854775807:
                var v21 int16 = -32768
                _ = v21
            case 127:
                v13 = 42
            default:
                if len(v12) > 0 {
                    v12[0] = 0
                }
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    if (!false) {
        _ = len(v12)
        if (32767 <= 2147483647) {
            fmt.Println(v13)
            for fi22 := 0; fi22 < 3; fi22++ {
                if true {
                    var v23 int64 = (^(v13))
                    _ = v23
                }
                var v24 int32 = int32(fi22)
                _ = len(v12)
                _ = fi22
                _ = v24
            }
        } else {
            var v25 float32 = 3.14159
            _ = v25
        }
        for fi26 := 0; fi26 < 15; fi26++ {
            v13 = genFn0(215, -1)
            _ = fi26
        }
    }
    _ = v12
    _ = v13
    return true
}

func main() {
    v1 := []int{}
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 bool = (true || true)
    var v5 float64 = 0.1
    for fi27 := 0; fi27 < 14; fi27++ {
        if (0.0 <= (v5 / v5)) {
            v28 := v2[4294967295]
            for fi29 := 0; fi29 < 3; fi29++ {
                switch ("hello" + "") {
                case "":
                    if true {
                        fmt.Println(fi27, v4)
                        _ = len(v1)
                        fmt.Println(fi27, v4, v5)
                        var v30 float64 = 3.14159
                        _ = v30
                    } else {
                        switch ("" + "hello") {
                        case "hello":
                            var v31 uint32 = 4294967295
                            _ = v31
                        case "\n\t":
                            v4 = (127 < -18)
                        }
                        v32 := GenS1{F0: fi29, F1: "\n\t"}
                        _ = v32
                    }
                case "foo bar":
                    if (!true) {
                        delete(v3, "foo bar")
                        v28 = true
                        v28 = true
                        var v33 bool = (true && v4)
                        _ = v33
                    } else {
                        fmt.Println(fi27, fi29)
                        v34 := GenS1{F0: fi27, F1: ""}
                        v5 = v5
                        v4 = (73 < 11)
                        _ = v34
                    }
                case "hello":
                    var v35 int32 = 42
                    _ = v35
                default:
                    if v4 {
                        fmt.Println(v5)
                        v4 = true
                    }
                }
                _ = fi29
            }
            _ = v28
        } else {
            v5 -= 1.5
        }
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v36 uint8 = 90
            _ = v36
        }
        _ = fi27
    }
    if (42 < -1) {
        var v37 bool = (true || true)
        switch -9223372036854775808 {
        case 691:
            var v38 uint8 = 1
            _ = v38
        case 9223372036854775807:
            switch ("foo bar" + (("\n\t" + "") + ("\n\t" + ("hello" + "foo bar")))) {
            case "a":
                switch -1 {
                case 9223372036854775807:
                    if len(v1) > 0 {
                        v1[0] = -206
                    }
                case 128:
                    v39 := GenS1{F0: -2147483648, F1: ("" + ("" + ("a" + "a")))}
                    _ = v39
                case 4294967295:
                    if (145 < -1) {
                        v40 := v3["hello"]
                        v37 = true
                        _ = v40
                    } else {
                        for fi41 := 0; fi41 < 1; fi41++ {
                            v42 := v2[127]
                            var v43 float32 = (1e10 + 1.5)
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        fmt.Println(v4, v5, v37)
                        if len(v1) > 0 {
                            v1[0] = -651
                        }
                    }
                default:
                    _ = len(v1)
                }
            case "foo bar":
                v5 = 1.5
            case "hello":
                v44 := v2[4294967295]
                _ = v44
            default:
                switch ("hello" + "") {
                case "foo bar":
                    for fi45 := 0; fi45 < 19; fi45++ {
                        v37 = (0 > 787)
                        switch "\n\t" {
                        case "a":
                            switch fi45 {
                            case 127:
                                var v46 uint8 = 255
                                _ = v46
                            case 0:
                                v37 = true
                            default:
                                switch 193 {
                                case 128:
                                    v37 = (!true)
                                default:
                                    v47 := GenS1{F0: fi45, F1: (("hello" + "\n\t") + "")}
                                    _ = v47
                                }
                            }
                        case "\n\t":
                            if v4 {
                                if (999 > 47) {
                                    if len(v1) > 0 {
                                        v1[0] = fi45
                                    }
                                    var v48 uint64 = 581
                                    _ = v48
                                } else {
                                    fmt.Println(v4)
                                    v49 := GenS2{F0: 9223372036854775807}
                                    switch (fi45 * fi45) {
                                    case 9223372036854775807:
                                        for fi50 := 0; fi50 < 18; fi50++ {
                                            v37 = (("hello" + ("a" + "\n\t")) <= "\n\t")
                                            fmt.Println(fi45)
                                            var v51 int64 = 975
                                            _ = fi50
                                            _ = v51
                                        }
                                    }
                                    _ = v49
                                }
                            } else {
                                v37 = (0 <= -231)
                                switch "hello" {
                                case "\n\t":
                                    v37 = v4
                                }
                            }
                        case "":
                            v52 := GenS2{F0: 32768}
                            _ = v52
                        }
                        _ = fi45
                    }
                case "hello":
                    var v53 bool = v4
                    _ = v53
                case "\n\t":
                    v2[4294967295] = (!("hello" < ("" + "\n\t")))
                }
            }
        case 127:
            v4 = (("foo bar" + ("hello" + "hello")) < "a")
        default:
            if v37 {
                v54 := GenS3{F0: 0, F1: 42}
                var v55 float32 = -2.25
                if false {
                    switch -203 {
                    case -2147483648:
                        if v4 {
                            var v56 int = 1
                            _ = v56
                        } else {
                            for fi57 := 0; fi57 < 1; fi57++ {
                                for fi58 := 0; fi58 < 12; fi58++ {
                                    if (false || false) {
                                        var v59 int64 = -1
                                        _ = v59
                                    } else {
                                        if (1 <= 255) {
                                            switch (fi57 / ((fi58) | 1)) {
                                            case 2147483647:
                                                v60 := v2[(fi58 << ((0) & 63))]
                                                _ = v60
                                            default:
                                                fmt.Println(v55, v55, v37)
                                            }
                                            v37 = (217 >= 2)
                                            fmt.Println(v4, v55, v55)
                                        } else {
                                            v4 = true
                                            v4 = v37
                                        }
                                    }
                                    v37 = v4
                                    _ = fi58
                                }
                                v37 = (v37 || v37)
                                _ = fi57
                            }
                            var v61 int = 42
                            _ = v61
                        }
                    default:
                        switch -2147483648 {
                        case 2147483647:
                            var v62 int = 42
                            _ = v62
                        case 430:
                            if (-991 == 845) {
                                fmt.Println(v55, v4, v4)
                                if (2147483647 <= 452) {
                                    v63 := GenS4{F0: -32768, F1: ((false || true) || v37)}
                                    _ = v63
                                }
                                var v64 int8 = -13
                                _ = v64
                            } else {
                                v1 = append(v1, 9223372036854775807)
                                v1 = append(v1, 1)
                                v65 := GenS5{F0: genFn0(42, -32768), F1: ("" + (((("\n\t" + ("" + "foo bar")) + ("a" + "foo bar")) + "") + "foo bar")), F2: 127}
                                _ = v65
                            }
                        case -2147483648:
                            for fi66 := 0; fi66 < 20; fi66++ {
                                if true {
                                    switch 128 {
                                    case 4294967295:
                                        var v67 uint = 4294967295
                                        _ = v67
                                    }
                                    switch ("hello" + "foo bar") {
                                    case "\n\t":
                                        if v4 {
                                            v4 = true
                                            switch (fi66 - fi66) {
                                            case -9223372036854775808:
                                                v68 := GenS5{F0: 0, F1: "foo bar", F2: (fi66 << ((18446744073709551615) & 63))}
                                                _ = v68
                                            case 1:
                                                for fi69 := 0; fi69 < 13; fi69++ {
                                                    var v70 uint16 = 2
                                                    _ = fi69
                                                    _ = v70
                                                }
                                            case 128:
                                                if v4 {
                                                    var v71 float64 = ((v5 - 1e10) * v5)
                                                    var v72 uint = 2
                                                    _ = v71
                                                    _ = v72
                                                } else {
                                                    var v73 string = "hello"
                                                    _ = v73
                                                }
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    case "":
                                        var v74 float32 = v55
                                        _ = v74
                                    case "hello":
                                        v37 = true
                                    default:
                                        v75 := GenS5{F0: genFn0(2, 42), F1: "hello", F2: -9223372036854775808}
                                        _ = v75
                                    }
                                }
                                for fi76 := 0; fi76 < 14; fi76++ {
                                    v77 := GenS5{F0: -1, F1: "foo bar", F2: fi76}
                                    _ = fi76
                                    _ = v77
                                }
                                if v37 {
                                    v4 = v37
                                } else {
                                    var v78 int64 = 1
                                    if true {
                                        switch ("" + ("foo bar" + "\n\t")) {
                                        case "\n\t":
                                            var v79 float32 = v55
                                            _ = v79
                                        default:
                                            var v80 uint64 = 2
                                            _ = v80
                                        }
                                        var v81 int = (-(fi66))
                                        _ = v81
                                    }
                                    _ = v78
                                }
                                _ = fi66
                            }
                        }
                    }
                    v82 := GenS6{F0: 42, F1: -987, F2: 1}
                    v83 := GenS6{F0: -128, F1: 0, F2: 42}
                    _ = v82
                    _ = v83
                } else {
                    var v84 uint64 = 2
                    var v85 uint8 = 2
                    v5 += v5
                    _ = v84
                    _ = v85
                }
                _ = v54
                _ = v55
            }
        }
        v5 = v5
        _ = v37
    }
    v5 -= (1e10 - v5)
    var v86 uint32 = 2147483648
    v5 += v5
    v5 = 0.0
    v87 := GenS1{F0: 4294967295, F1: "foo bar"}
    if (2147483647 != 4294967295) {
        var v88 uint16 = 42
        for fi89 := 0; fi89 < 12; fi89++ {
            if len(v1) > 0 {
                v1[0] = (fi89 & fi89)
            }
            var v90 float32 = 3.14159
            var v91 float32 = (v90 - v90)
            _ = fi89
            _ = v90
            _ = v91
        }
        _ = v88
    } else {
        v5 += (v5 / 0.1)
        var v92 int32 = 717
        fmt.Println(v4, v92)
        _ = v92
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v86
    _ = v87
}
