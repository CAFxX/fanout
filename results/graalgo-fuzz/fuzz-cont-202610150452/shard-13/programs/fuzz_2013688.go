package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 int8; F2 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 string }
type GenS9 struct { F0 uint; F1 uint }
func genFn0() bool {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = -1
    }
    v8 := GenS1{F0: 9223372036854775807, F1: 2, F2: 42}
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 uint16 = 42
    var v2 bool = false
    var v3 float64 = 3.14159
    var v4 uint64 = 42
    v5 := make(map[uint32]string)
    v6 := []string{""}
    for fi9 := 0; fi9 < 19; fi9++ {
        _ = len(v5)
        if true {
            v10 := GenS2{F0: int8(v4)}
            var v11 uint16 = (v1 << ((255) & 63))
            var v12 uint32 = 2147483648
            _ = v10
            _ = v11
            _ = v12
        }
        if v2 {
            var v13 uint = 42
            v14 := GenS3{F0: 0}
            _ = v13
            _ = v14
        }
        _ = fi9
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            if v2 {
                if len(v6) > 0 {
                    v6[0] = (("hello" + "\n\t") + "\n\t")
                }
                if len(v6) > 0 {
                    v6[0] = ("hello" + "hello")
                }
            }
            _ = fi16
        }
        var v17 int64 = 42
        _ = fi15
        _ = v17
    }
    if (((false || v2) && (!v2)) && (2147483647 > 4294967295)) {
        switch ("\n\t" + "") {
        case "hello":
            if (40 < 42) {
                var v18 bool = (true && (false || v2))
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    v3 = 1.5
                }
                if (1000 > 325) {
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v3 = ((v3 / (float64(1e10) / float64(3.14159))) / ((0.1 + v3) / (v3 - float64(((1e10 + float32(v3)) * (3.14159 / 3.14159))))))
                        v3 += -2.25
                        _ = fi19
                    }
                } else {
                    var v20 int8 = 42
                    switch int(v20) {
                    case 2147483647:
                        v21 := GenS3{F0: -1}
                        _ = v21
                    case -2147483648:
                        v22 := GenS2{F0: v20}
                        _ = v22
                    default:
                        v23 := GenS2{F0: (v20 * (^(v20)))}
                        _ = v23
                    }
                    _ = v20
                }
                _ = v18
            }
        }
    }
    v4 = 777
    for fi24 := 0; fi24 < 16; fi24++ {
        v25 := GenS2{F0: 1}
        _ = fi24
        _ = v25
    }
    _ = len(v6)
    v2 = (("\n\t" + ("" + "hello")) == "foo bar")
    if v2 {
        var v26 float64 = v3
        switch "a" {
        case "\n\t":
            for fi27 := 0; fi27 < 14; fi27++ {
                if (!("" != "a")) {
                    var v28 int = (^(fi27))
                    _ = v28
                }
                var v29 uint32 = 2
                var v30 uint8 = 255
                _ = fi27
                _ = v29
                _ = v30
            }
        case "foo bar":
            if false {
                v31 := GenS4{F0: (v4 - 42), F1: (^(v4))}
                _ = v31
            }
        case "hello":
            v26 = v3
        }
        var v32 int32 = 378
        _ = v26
        _ = v32
    } else {
        var v33 uint64 = v4
        for fi34 := 0; fi34 < 15; fi34++ {
            for fi35 := 0; fi35 < 9; fi35++ {
                if (-631 > -1) {
                    v1 += 836
                    fmt.Println(v3)
                    delete(v5, 2147483648)
                } else {
                    var v36 bool = true
                    v1 -= (v1 % ((32767) | 1))
                    _ = v36
                }
                v37 := GenS3{F0: 2147483647}
                if false {
                    if (v4 > v33) {
                        var v38 uint32 = 1
                        var v39 uint16 = (v1 ^ (v1 | v1))
                        if true {
                            if (128 <= 128) {
                                var v40 uint = 1
                                v41 := v5[v38]
                                _ = v40
                                _ = v41
                            } else {
                                v42 := GenS1{F0: 9223372036854775807, F1: 2, F2: 42}
                                v33 = v33
                                _ = v42
                            }
                            v4 = (^(v4))
                            switch "hello" {
                            case "hello":
                                var v43 uint = 255
                                _ = v43
                            case "a":
                                switch "\n\t" {
                                case "a":
                                    if v2 {
                                        var v44 int = fi35
                                        fmt.Println(v3)
                                        var v45 float32 = 1e10
                                        v1 = v1
                                        var v46 uint16 = (v39 << ((18446744073709551615) & 63))
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    } else {
                                        v47 := GenS5{F0: v1}
                                        if false {
                                            if v2 {
                                                switch "a" {
                                                case "\n\t":
                                                    v48 := GenS2{F0: 42}
                                                    _ = v48
                                                case "hello":
                                                    var v49 uint16 = v39
                                                    _ = v49
                                                case "":
                                                    var v50 uint32 = (v38 * v38)
                                                    _ = v50
                                                }
                                                if false {
                                                    if v2 {
                                                        v1 -= 2
                                                        if (1.5 == -2.25) {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            if (v2 || v2) {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                            } else {
                                                                if (128 != 127) {
                                                                    var v51 int64 = -222
                                                                    switch "a" {
                                                                    case "":
                                                                        v52 := GenS5{F0: (^(v39))}
                                                                        _ = v52
                                                                    case "foo bar":
                                                                        switch ("" + "a") {
                                                                        case "a":
                                                                            if v2 {
                                                                                fmt.Println(v1, v39, v39)
                                                                                var v53 int32 = 203
                                                                                switch (fi35 + int(v53)) {
                                                                                case 32767:
                                                                                    v54 := GenS6{F0: 1, F1: 100, F2: v53}
                                                                                    _ = v54
                                                                                case 959:
                                                                                    var v55 float32 = (0.1 + 0.1)
                                                                                    _ = v55
                                                                                case -103:
                                                                                    if false {
                                                                                        if ((((!v2) && v2) || v2) && (v51 != v51)) {
                                                                                            var v56 float64 = 0.1
                                                                                            _ = v56
                                                                                        }
                                                                                        v6 = append(v6, "")
                                                                                        var v57 uint16 = (^(v1))
                                                                                        _ = v57
                                                                                    } else {
                                                                                        v1 += v39
                                                                                        fmt.Println(fi35, v51, fi35)
                                                                                    }
                                                                                default:
                                                                                    v1 = (v39 >> ((1) & 63))
                                                                                }
                                                                                _ = len(v6)
                                                                                _ = v53
                                                                            } else {
                                                                                delete(v5, v38)
                                                                                if ("a" > "") {
                                                                                    v58 := GenS6{F0: -1, F1: 42, F2: -697}
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        if (v39 != v1) {
                                                                                            v39 = v1
                                                                                            fmt.Println(v4)
                                                                                        } else {
                                                                                            if (int8(v1) <= -1) {
                                                                                                v59 := GenS5{F0: v39}
                                                                                                v60 := GenS3{F0: -2147483648}
                                                                                                switch "a" {
                                                                                                case "":
                                                                                                    fmt.Println(v4, v51)
                                                                                                }
                                                                                                var v61 int8 = -128
                                                                                                v39 -= 2
                                                                                                var v62 int64 = v51
                                                                                                if false {
                                                                                                    var v63 uint32 = 2147483648
                                                                                                    _ = v63
                                                                                                }
                                                                                                _ = v59
                                                                                                _ = v60
                                                                                                _ = v61
                                                                                                _ = v62
                                                                                            } else {
                                                                                                switch ("" + (("hello" + "\n\t") + ("foo bar" + "hello"))) {
                                                                                                case "foo bar":
                                                                                                    if len(v6) >= 2 {
                                                                                                        v6 = v6[1:len(v6)]
                                                                                                    }
                                                                                                case "":
                                                                                                    var v64 int32 = 0
                                                                                                    _ = v64
                                                                                                case "hello":
                                                                                                    if ((false || (!v2)) && v2) {
                                                                                                        v5[(^(v38))] = "hello"
                                                                                                        v4 = v4
                                                                                                        var v65 bool = (v2 && ((4294967295 > fi35) && false))
                                                                                                        _ = v65
                                                                                                    }
                                                                                                }
                                                                                                v4 = (v33 / ((v4) | 1))
                                                                                                fmt.Println(v51, v33, v51)
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        if ((^(v39)) == (v39 | 944)) {
                                                                                            if (v2 && (!true)) {
                                                                                                var v66 uint64 = v4
                                                                                                if true {
                                                                                                    switch "a" {
                                                                                                    case "hello":
                                                                                                        var v67 int = fi34
                                                                                                        _ = v67
                                                                                                    case "\n\t":
                                                                                                        var v68 string = "\n\t"
                                                                                                        _ = v68
                                                                                                    }
                                                                                                }
                                                                                                switch (((("a" + (("a" + "a") + (("\n\t" + ((("\n\t" + ("" + ("foo bar" + ("hello" + "a")))) + "a") + ("foo bar" + "hello"))) + "hello"))) + ((("\n\t" + "foo bar") + ("foo bar" + ("a" + ("" + "foo bar")))) + "foo bar")) + ("foo bar" + "hello")) + (("" + ("hello" + "\n\t")) + "a")) {
                                                                                                case "hello":
                                                                                                    fmt.Println(fi34, fi35, v51)
                                                                                                }
                                                                                                _ = v66
                                                                                            } else {
                                                                                                var v69 float64 = v3
                                                                                                v70 := v5[v38]
                                                                                                v71 := GenS7{F0: (((v70 + ("\n\t" + ((v70 + "a") + "hello"))) + v70) + (v70 + v70))}
                                                                                                _ = v69
                                                                                                _ = v70
                                                                                                _ = v71
                                                                                            }
                                                                                            if (fi34 >= 468) {
                                                                                                v39 = (v1 % ((v39) | 1))
                                                                                                v39 = v1
                                                                                                switch (-(fi35)) {
                                                                                                case 128:
                                                                                                    var v72 uint64 = (v33 & v4)
                                                                                                    _ = v72
                                                                                                default:
                                                                                                    var v73 int = (fi34 % (((fi35 / (((fi34 * fi34)) | 1))) | 1))
                                                                                                    _ = v73
                                                                                                }
                                                                                            } else {
                                                                                                v51 = (-(v51))
                                                                                                var v74 uint32 = (^(v38))
                                                                                                _ = v74
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v51 -= (-(v51))
                                                                                    _ = v58
                                                                                } else {
                                                                                    v5[v38] = ""
                                                                                    var v75 bool = v2
                                                                                    _ = v75
                                                                                }
                                                                            }
                                                                        case "":
                                                                            v1 = v1
                                                                        case "hello":
                                                                            fmt.Println(fi35, v38)
                                                                        default:
                                                                            _ = len(v6)
                                                                        }
                                                                    case "hello":
                                                                        v1 += 651
                                                                        v38 += 4294967295
                                                                    }
                                                                    _ = v51
                                                                } else {
                                                                    v1 += (v1 * (v39 * v1))
                                                                    fmt.Println(v38)
                                                                }
                                                                var v76 uint8 = 255
                                                                _ = v76
                                                            }
                                                        }
                                                        var v77 uint64 = 18446744073709551615
                                                        _ = v77
                                                    }
                                                } else {
                                                    if (v2 && ((0.1 <= (float64(1.5) + (v3 / (v3 + v3)))) && v2)) {
                                                        v1 += (^(v39))
                                                        var v78 uint16 = (v1 | 379)
                                                        _ = v78
                                                    }
                                                }
                                            }
                                            _ = len(v5)
                                        }
                                        _ = v47
                                    }
                                default:
                                    v79 := GenS8{F0: (v1 - v1), F1: "a"}
                                    _ = v79
                                }
                            case "\n\t":
                                v80 := GenS5{F0: v1}
                                _ = v80
                            }
                        } else {
                            v4 -= 308
                            if v2 {
                                if ((((126 >= 95) && (v2 || (true || (v2 && v2)))) || v2) || false) {
                                    fmt.Println(v2, fi35, v1)
                                    v2 = (9223372036854775807 < -1)
                                }
                                if false {
                                    if len(v6) > 0 {
                                        v6[0] = "hello"
                                    }
                                }
                            } else {
                                switch "\n\t" {
                                case "a":
                                    v81 := GenS7{F0: "\n\t"}
                                    _ = v81
                                case "hello":
                                    var v82 uint = 1
                                    _ = v82
                                case "foo bar":
                                    var v83 float32 = 3.14159
                                    _ = v83
                                default:
                                    var v84 int8 = int8(fi35)
                                    _ = v84
                                }
                            }
                        }
                        _ = v38
                        _ = v39
                    } else {
                        fmt.Println(v3)
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v85 int8 = -128
                    if true {
                        if (2147483648 <= 65535) {
                            v6 = append(v6, "foo bar")
                        } else {
                            var v86 float64 = ((float64((0.0 - (1.5 * (0.1 / 1e10)))) / 3.14159) - -2.25)
                            v3 = v3
                            _ = v86
                        }
                    } else {
                        var v87 uint16 = (v1 ^ 2)
                        v85 += 100
                        var v88 bool = false
                        _ = v87
                        _ = v88
                    }
                    _ = v85
                } else {
                    if v2 {
                        v3 = float64(0.0)
                        v89 := GenS7{F0: ("foo bar" + "hello")}
                        var v90 string = ("" + "\n\t")
                        _ = v89
                        _ = v90
                    }
                }
                _ = fi35
                _ = v37
            }
            _ = fi34
        }
        _ = v33
    }
    var v91 float64 = 1e10
    v92 := GenS9{F0: 255, F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
    _ = v92
}
