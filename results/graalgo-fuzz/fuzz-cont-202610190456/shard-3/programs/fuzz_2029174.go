package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8; F1 uint16; F2 uint }

func main() {
    v1 := []string{"", (("hello" + ("foo bar" + ("a" + "foo bar"))) + "")}
    var v2 int8 = 55
    v3 := make(map[uint32]int)
    var v4 int16 = 42
    v5 := make(map[string]int)
    var v6 float64 = 1e10
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v7 int8 = -128
    var v8 float32 = ((-2.25 - (((0.1 / 1.0) * 0.0) - 3.14159)) - float32(0.1))
    fmt.Println(v6, v8, v2)
    v2 -= 0
    v9 := GenS1{F0: (v4 / ((v4) | 1))}
    v10 := GenS2{F0: "\n\t"}
    var v11 int8 = v2
    var v12 bool = false
    var v13 float32 = float32(v6)
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int8 = (v7 % ((v2) | 1))
        for fi16 := 0; fi16 < 8; fi16++ {
            switch ("foo bar" + "foo bar") {
            case "a":
                var v17 uint32 = 2147483648
                _ = v17
            case "\n\t":
                v4 += -515
            case "hello":
                var v18 uint8 = 99
                _ = v18
            default:
                var v19 int16 = 388
                _ = v19
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    var v20 int64 = 1
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 uint16 = 255
            _ = fi22
            _ = v23
        }
        if ("" == (((("hello" + (("hello" + "\n\t") + "\n\t")) + ("a" + "\n\t")) + "a") + ("a" + ""))) {
            var v24 string = ""
            _ = v24
        } else {
            v25 := GenS3{F0: true}
            for fi26 := 0; fi26 < 11; fi26++ {
                if v12 {
                    var v27 int32 = 0
                    _ = v27
                } else {
                    v20 = v20
                    v28 := GenS4{F0: (!(-639 == 0)), F1: 32767}
                    if (!(!true)) {
                        var v29 int16 = v4
                        _ = v29
                    } else {
                        v30 := GenS4{F0: (2 > 42), F1: 489}
                        var v31 float64 = (3.14159 - v6)
                        var v32 uint64 = 4294967295
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    _ = v28
                }
                v20 = int64(fi26)
                _ = fi26
            }
            _ = v25
        }
        if (v11 <= int8(v20)) {
            for fi33 := 0; fi33 < 19; fi33++ {
                _ = len(v1)
                switch ("\n\t" + (("a" + "") + "")) {
                case "\n\t":
                    if v12 {
                        if false {
                            switch ((("" + ("a" + "hello")) + ("a" + ("" + "hello"))) + "hello") {
                            case "a":
                                var v34 int16 = (v4 % (((-(v4))) | 1))
                                _ = v34
                            case "":
                                var v35 int64 = v20
                                _ = v35
                            case "\n\t":
                                v8 -= (float32((float64((float32(3.14159) / v13)) + 1e10)) / v8)
                            default:
                                var v36 uint = 255
                                _ = v36
                            }
                            var v37 float64 = float64(v8)
                            _ = v37
                        } else {
                            v38 := GenS5{F0: (fi33 + fi21)}
                            fmt.Println(v13, v20)
                            if (v6 >= v6) {
                                switch fi33 {
                                case 127:
                                    fmt.Println(v6)
                                }
                                if v12 {
                                    v1 = append(v1, ("\n\t" + ""))
                                } else {
                                    v1 = append(v1, (("\n\t" + (("foo bar" + "hello") + "a")) + "\n\t"))
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v6 = (((v6 * v6) - v6) * v6)
                                }
                            }
                            var v39 int32 = -616
                            var v40 uint32 = 571
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        if v12 {
                            var v41 int16 = (v4 / ((v4) | 1))
                            var v42 uint32 = 42
                            _ = v41
                            _ = v42
                        } else {
                            if (-2.25 > 0.1) {
                                fmt.Println(v4, fi21)
                                if v12 {
                                    v4 = (v4 * (-(v4)))
                                    switch (fi33 & fi21) {
                                    case 4294967295:
                                        if ((-2147483648 > 0) && (!((0 == 65535) && true))) {
                                            if (18446744073709551615 <= 18446744073709551615) {
                                                v6 = v6
                                                fmt.Println(v11, v2)
                                                var v43 int = fi33
                                                _ = v43
                                            } else {
                                                var v44 int64 = (-(v20))
                                                switch -2147483648 {
                                                case -2147483648:
                                                    fmt.Println(v20, v20, v8)
                                                    if ((^(v44)) <= 740) {
                                                        var v45 bool = ((v4 / ((-32768) | 1)) == (v4 * v4))
                                                        var v46 bool = (false && v12)
                                                        if len(v1) > 0 {
                                                            v1[0] = "foo bar"
                                                        }
                                                        v45 = v46
                                                        var v47 string = ("a" + "")
                                                        _ = v45
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                case 32768:
                                                    v48 := v3[2]
                                                    _ = v48
                                                case -40:
                                                    var v49 int16 = (-(v4))
                                                    _ = v49
                                                default:
                                                    var v50 uint32 = 2
                                                    _ = v50
                                                }
                                                _ = v44
                                            }
                                            if v12 {
                                                if (!v12) {
                                                    if true {
                                                        switch "\n\t" {
                                                        case "":
                                                            switch (("foo bar" + ((("a" + "hello") + ("" + "foo bar")) + ("foo bar" + (("foo bar" + "hello") + "foo bar")))) + "foo bar") {
                                                            case "a":
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            default:
                                                                if (false || true) {
                                                                    var v51 int32 = 246
                                                                    _ = v51
                                                                } else {
                                                                    v6 -= float64((((((((0.1 - v13) * float32(v6)) + float32(v6)) / 1e10) * v13) - 1.5) * (v13 / (v8 + v8))))
                                                                }
                                                            }
                                                        case "\n\t":
                                                            v7 = (v7 % ((v2) | 1))
                                                        }
                                                        var v52 uint8 = 2
                                                        if ("hello" > "foo bar") {
                                                            v53 := GenS6{F0: 919}
                                                            if v12 {
                                                                v13 = v13
                                                                v6 = (v6 - (-2.25 * 3.14159))
                                                                switch (fi21 % (((fi21 - (fi21 / (((fi33 / ((189) | 1))) | 1)))) | 1)) {
                                                                case 4294967295:
                                                                    v54 := GenS7{F0: 255}
                                                                    _ = v54
                                                                case 32767:
                                                                    fmt.Println(v2, fi21, v12)
                                                                }
                                                            } else {
                                                                switch "foo bar" {
                                                                case "":
                                                                    v12 = v12
                                                                case "\n\t":
                                                                    if v12 {
                                                                        v13 += float32((float64(1.5) * 1.5))
                                                                        v55 := GenS5{F0: fi33}
                                                                        fmt.Println(v6)
                                                                        _ = v55
                                                                    } else {
                                                                        if (!v12) {
                                                                            var v56 uint = 784
                                                                            var v57 int8 = 100
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                        var v58 int16 = v4
                                                                        if (324 <= 0) {
                                                                            var v59 int8 = (v7 * v7)
                                                                            _ = v59
                                                                        } else {
                                                                            var v60 int32 = 2147483647
                                                                            var v61 uint = 168
                                                                            var v62 uint8 = 4
                                                                            _ = v60
                                                                            _ = v61
                                                                            _ = v62
                                                                        }
                                                                        _ = v58
                                                                    }
                                                                case "hello":
                                                                    if true {
                                                                        fmt.Println(v11, v52, v11)
                                                                        if (!v12) {
                                                                            v13 = (1e10 * v13)
                                                                        }
                                                                        var v63 float32 = v13
                                                                        _ = v63
                                                                    }
                                                                }
                                                            }
                                                            _ = v53
                                                        } else {
                                                            v64 := GenS7{F0: 322}
                                                            var v65 bool = v12
                                                            switch (fi33 | (fi21 >> ((18446744073709551615) & 63))) {
                                                            case 4294967295:
                                                                fmt.Println(v12, v7)
                                                            case 127:
                                                                if ("a" <= ("a" + "")) {
                                                                    var v66 float64 = 0.1
                                                                    _ = v66
                                                                } else {
                                                                    var v67 uint8 = (v52 ^ (v52 * v52))
                                                                    v4 = v4
                                                                    _ = v67
                                                                }
                                                            case -1:
                                                                var v68 int32 = 791
                                                                _ = v68
                                                            default:
                                                                fmt.Println(fi33)
                                                            }
                                                            v69 := GenS6{F0: v20}
                                                            v70 := GenS8{F0: v7}
                                                            var v71 uint = 333
                                                            _ = v64
                                                            _ = v65
                                                            _ = v69
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        _ = v52
                                                    }
                                                    var v72 int32 = 42
                                                    v73 := v3[65535]
                                                    _ = v72
                                                    _ = v73
                                                }
                                                v74 := GenS9{F0: 2}
                                                var v75 int32 = int32(fi21)
                                                if (int64(fi33) != (^(v20))) {
                                                    var v76 uint16 = 485
                                                    var v77 int16 = (v4 ^ v4)
                                                    var v78 uint64 = 4294967295
                                                    v79 := GenS3{F0: ((fi21 < (fi33 ^ fi33)) || true)}
                                                    v4 = v4
                                                    var v80 uint = 42
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                }
                                                fmt.Println(v13, v12, v12)
                                                _ = v74
                                                _ = v75
                                            }
                                            if (!(-1 <= (v2 | v2))) {
                                                v4 = -599
                                            }
                                        }
                                    case 9223372036854775807:
                                        if (0 <= 249) {
                                            if len(v1) > 0 {
                                                v1[0] = "hello"
                                            }
                                        }
                                    case 2147483647:
                                        switch fi21 {
                                        case 4294967295:
                                            var v81 int64 = (v20 + v20)
                                            _ = v81
                                        case 214:
                                            var v82 string = "hello"
                                            _ = v82
                                        default:
                                            v83 := GenS5{F0: (fi33 - fi33)}
                                            _ = v83
                                        }
                                    }
                                }
                            } else {
                                var v84 uint16 = 65535
                                v20 = (v20 / ((v20) | 1))
                                var v85 uint16 = 32767
                                _ = v84
                                _ = v85
                            }
                        }
                    } else {
                        var v86 float64 = 0.0
                        _ = v86
                    }
                case "a":
                    if v12 {
                        var v87 int8 = 0
                        var v88 float64 = 0.1
                        _ = v87
                        _ = v88
                    } else {
                        switch (("a" + "hello") + (("\n\t" + ("foo bar" + "a")) + "\n\t")) {
                        case "":
                            _ = len(v3)
                        }
                        _ = len(v3)
                        v12 = ((v12 && (4294967295 <= 2)) || v12)
                    }
                case "hello":
                    switch ("\n\t" + "\n\t") {
                    case "foo bar":
                        var v89 float64 = 1e10
                        _ = v89
                    case "hello":
                        v90 := GenS10{F0: (^(v7)), F1: 65535, F2: 639}
                        _ = v90
                    case "\n\t":
                        fmt.Println(v20)
                    }
                }
                v91 := GenS8{F0: (v11 | v7)}
                _ = fi33
                _ = v91
            }
            var v92 int16 = v4
            _ = v92
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
