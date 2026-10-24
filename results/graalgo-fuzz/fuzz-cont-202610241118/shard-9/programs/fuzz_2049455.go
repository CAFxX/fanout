package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 string; F2 bool }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 bool; F1 uint; F2 int32 }
type GenS10 struct { F0 string }

func main() {
    v1 := []string{"\n\t"}
    v2 := []uint32{65535, 0, 2147483647}
    var v3 bool = (!(float64(0.0) == 3.14159))
    v4 := []uint32{107, 1, 0}
    var v5 uint64 = 227
    v6 := make(map[string]string)
    var v7 float64 = ((-2.25 - (1e10 * (1.5 - (1.5 - 0.1)))) * float64((((1.5 / ((3.14159 * float32(0.0)) + ((1.5 * ((float32(1.5) / -2.25) * 3.14159)) / 1e10))) - 0.0) / (3.14159 * 0.1))))
    if (!((v7 * float64(3.14159)) > -2.25)) {
        v8 := GenS1{F0: "foo bar"}
        fmt.Println(v3, v7, v5)
        for fi9 := 0; fi9 < 11; fi9++ {
            if v3 {
                v10 := GenS1{F0: "a"}
                _ = len(v4)
                _ = v10
            }
            var v11 int8 = 20
            fmt.Println(fi9, v3)
            _ = fi9
            _ = v11
        }
        _ = v8
    }
    fmt.Println(v3)
    v12 := GenS1{F0: "hello"}
    v5 += v5
    _ = len(v1)
    var v13 uint32 = 4294967295
    if len(v2) > 0 {
        v2[0] = v13
    }
    v13 -= v13
    for fi14 := 0; fi14 < 10; fi14++ {
        if v3 {
            v15 := v6[""]
            if ((!(!(!true))) || false) {
                v7 = float64(float32(1e10))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 string = v15
                    _ = fi16
                    _ = v17
                }
            } else {
                var v18 uint64 = (v5 % ((1) | 1))
                var v19 uint64 = (v18 - (^(v5)))
                _ = v18
                _ = v19
            }
            _ = v15
        }
        v7 = v7
        switch (fi14 ^ fi14) {
        case 9223372036854775807:
            for fi20 := 0; fi20 < 16; fi20++ {
                fmt.Println(v5, fi14)
                switch ((("hello" + "a") + "hello") + "foo bar") {
                case "foo bar":
                    v3 = v3
                }
                if (!(127 < -59)) {
                    switch (^(fi14)) {
                    case -1:
                        v5 = 328
                    case -10:
                        if true {
                            if (1.5 >= 1e10) {
                                if (36 != -105) {
                                    v13 = (v13 & v13)
                                    v5 += (v5 % ((v5) | 1))
                                    v21 := v6["\n\t"]
                                    _ = v21
                                }
                                if false {
                                    var v22 int8 = 42
                                    var v23 int16 = 0
                                    _ = v22
                                    _ = v23
                                } else {
                                    var v24 uint = 42
                                    if v3 {
                                        if true {
                                            var v25 int64 = 9223372036854775807
                                            _ = v25
                                        }
                                        if false {
                                            var v26 int64 = 903
                                            _ = v26
                                        }
                                        var v27 uint64 = 692
                                        if v3 {
                                            var v28 uint = (^(v24))
                                            fmt.Println(fi20, fi20)
                                            v29 := v6["hello"]
                                            _ = v28
                                            _ = v29
                                        } else {
                                            v3 = v3
                                        }
                                        _ = v27
                                    }
                                    if false {
                                        var v30 string = "\n\t"
                                        v13 = v13
                                        _ = v30
                                    } else {
                                        if (-1 == 109) {
                                            _ = len(v6)
                                            if (42 >= 127) {
                                                if ((!((v5 % (((^(v5))) | 1)) == v5)) && (v5 != (v5 >> ((v24) & 63)))) {
                                                    v31 := GenS2{F0: (^(v13)), F1: 4294967295, F2: 0}
                                                    if (!(!(true || (true || (1 != uint8(v13)))))) {
                                                        switch "a" {
                                                        case "hello":
                                                            v6["a"] = (("foo bar" + ("a" + "foo bar")) + (("foo bar" + "hello") + "foo bar"))
                                                        default:
                                                            var v32 int16 = 558
                                                            _ = v32
                                                        }
                                                    } else {
                                                        fmt.Println(v24, v3, v7)
                                                        fmt.Println(v3, fi14, v3)
                                                        var v33 bool = false
                                                        _ = v33
                                                    }
                                                    _ = v31
                                                } else {
                                                    switch "a" {
                                                    case "":
                                                        var v34 uint32 = (v13 % (((v13 / (((v13 - v13)) | 1))) | 1))
                                                        _ = v34
                                                    case "\n\t":
                                                        if (v7 <= 1.5) {
                                                            fmt.Println(fi14, v24)
                                                            v5 = (v5 << ((v24) & 63))
                                                            var v35 int32 = 42
                                                            _ = v35
                                                        } else {
                                                            v36 := GenS1{F0: ("\n\t" + ("foo bar" + ("" + ("a" + ""))))}
                                                            var v37 int32 = 1
                                                            var v38 uint = 2
                                                            if ((-1 != 42) || true) {
                                                                v24 = v24
                                                                delete(v6, "hello")
                                                                var v39 int64 = -391
                                                                _ = v39
                                                            } else {
                                                                v40 := GenS1{F0: (("a" + "hello") + "\n\t")}
                                                                _ = v40
                                                            }
                                                            var v41 uint8 = 234
                                                            fmt.Println(v38)
                                                            _ = v36
                                                            _ = v37
                                                            _ = v38
                                                            _ = v41
                                                        }
                                                    default:
                                                        v5 = v5
                                                    }
                                                    fmt.Println(v13, v13)
                                                }
                                            } else {
                                                _ = len(v1)
                                                var v42 uint = v24
                                                _ = v42
                                            }
                                            v24 -= 255
                                            v7 -= v7
                                        }
                                        if v3 {
                                            fmt.Println(v13, v24, v24)
                                            v24 = v24
                                        } else {
                                            if v3 {
                                                fmt.Println(v7, v5)
                                                switch fi20 {
                                                case 42:
                                                    switch ("a" + "a") {
                                                    case "":
                                                        fmt.Println(v13)
                                                    case "hello":
                                                        v43 := GenS3{F0: 255}
                                                        _ = v43
                                                    }
                                                case 4294967295:
                                                    if v3 {
                                                        if true {
                                                            var v44 int16 = -932
                                                            _ = v44
                                                        } else {
                                                            v45 := GenS4{F0: v3}
                                                            if v3 {
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v46 := GenS5{F0: (v5 / (((v5 + 1)) | 1)), F1: "a", F2: v3}
                                                                    _ = len(v4)
                                                                    v47 := v6[""]
                                                                    _ = v46
                                                                    _ = v47
                                                                case "\n\t":
                                                                    v48 := GenS3{F0: v24}
                                                                    if true {
                                                                        var v49 bool = ((v7 / v7) < (0.0 + (-2.25 / 1e10)))
                                                                        _ = len(v6)
                                                                        v7 = 1e10
                                                                        _ = v49
                                                                    } else {
                                                                        v50 := GenS4{F0: false}
                                                                        var v51 uint = 242
                                                                        _ = v50
                                                                        _ = v51
                                                                    }
                                                                    var v52 int32 = 0
                                                                    _ = v48
                                                                    _ = v52
                                                                case "":
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        switch fi20 {
                                                                        case 1:
                                                                            fmt.Println(v24, v3)
                                                                        default:
                                                                            v3 = v3
                                                                        }
                                                                    case "foo bar":
                                                                        switch fi14 {
                                                                        case 1:
                                                                            fmt.Println(v7)
                                                                        case 32768:
                                                                            var v53 uint = (v24 / ((508) | 1))
                                                                            _ = v53
                                                                        case 9223372036854775807:
                                                                            fmt.Println(fi14, fi14, v13)
                                                                        default:
                                                                            if (!v3) {
                                                                                v13 = v13
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    switch fi14 {
                                                                    case -436:
                                                                        var v54 uint = 4294967295
                                                                        _ = v54
                                                                    case -2147483648:
                                                                        _ = len(v4)
                                                                    }
                                                                    var v55 int8 = 1
                                                                    v56 := GenS6{F0: v5, F1: v13, F2: uint16(v24)}
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                            } else {
                                                                if v3 {
                                                                    v7 = v7
                                                                    v3 = v3
                                                                    v5 += v5
                                                                    v3 = (!false)
                                                                }
                                                            }
                                                            if (1 < 127) {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    switch int(v24) {
                                                                    case -511:
                                                                        if (!(32767 == 2)) {
                                                                            v57 := GenS7{F0: (fi14 % ((fi20) | 1)), F1: (v13 * 42)}
                                                                            var v58 uint32 = v13
                                                                            fmt.Println(v24, v5, v24)
                                                                            _ = v57
                                                                            _ = v58
                                                                        } else {
                                                                            fmt.Println(v3)
                                                                        }
                                                                    case 2147483647:
                                                                        v7 += v7
                                                                    default:
                                                                        fmt.Println(v13, v3, v7)
                                                                    }
                                                                case "a":
                                                                    switch ("" + "") {
                                                                    case "a":
                                                                        v3 = v3
                                                                    case "\n\t":
                                                                        fmt.Println(fi20, v5, v5)
                                                                    case "":
                                                                        var v59 int16 = -577
                                                                        _ = v59
                                                                    default:
                                                                        v60 := GenS5{F0: uint64(v13), F1: (("a" + "a") + ""), F2: true}
                                                                        _ = v60
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(fi20, fi14, fi14)
                                                                default:
                                                                    if (v3 || false) {
                                                                        if ("foo bar" == ("a" + "")) {
                                                                            if (v3 || (746 < 42)) {
                                                                                v7 -= ((float64(float32(1e10)) + float64(1.5)) * v7)
                                                                                v61 := GenS3{F0: uint(v5)}
                                                                                _ = v61
                                                                            }
                                                                            fmt.Println(v7, v13, v13)
                                                                            if ((v5 % ((v5) | 1)) > v5) {
                                                                                var v62 int32 = 1
                                                                                fmt.Println(v24, v3, v13)
                                                                                var v63 int = -4
                                                                                _ = v62
                                                                                _ = v63
                                                                            } else {
                                                                                fmt.Println(v5, fi20)
                                                                                var v64 uint = (v24 ^ (v24 + v24))
                                                                                _ = v64
                                                                            }
                                                                        } else {
                                                                            var v65 string = ("foo bar" + "\n\t")
                                                                            var v66 int16 = 1000
                                                                            _ = v65
                                                                            _ = v66
                                                                        }
                                                                    }
                                                                }
                                                                if (!(959 != 1)) {
                                                                    v7 = v7
                                                                    var v67 uint32 = uint32(v24)
                                                                    _ = v67
                                                                }
                                                                _ = len(v6)
                                                            }
                                                            var v68 float64 = (float64(3.14159) - ((1e10 * 0.0) - (v7 + v7)))
                                                            _ = v45
                                                            _ = v68
                                                        }
                                                        _ = len(v4)
                                                    }
                                                default:
                                                    var v69 int16 = 355
                                                    _ = v69
                                                }
                                                v5 = (v5 | v5)
                                            }
                                            if (int8(v5) > 100) {
                                                var v70 string = ""
                                                _ = v70
                                            }
                                            var v71 int16 = 1000
                                            _ = v71
                                        }
                                    }
                                    _ = v24
                                }
                                var v72 uint16 = 149
                                _ = v72
                            } else {
                                var v73 int8 = 0
                                fmt.Println(fi20, v73, v5)
                                _ = v73
                            }
                        } else {
                            v74 := GenS8{F0: (!v3)}
                            _ = v74
                        }
                    default:
                        var v75 int32 = 825
                        if (v3 || (1 > 4294967295)) {
                            var v76 string = ""
                            _ = v76
                        } else {
                            v77 := GenS9{F0: v3, F1: 1, F2: (-(v75))}
                            if (v3 || v3) {
                                if len(v4) > 0 {
                                    v4[0] = v13
                                }
                            } else {
                                v3 = ((false && v3) && true)
                            }
                            var v78 bool = true
                            v7 = float64(-2.25)
                            _ = v77
                            _ = v78
                        }
                        v7 = -2.25
                        _ = v75
                    }
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                    v79 := v6[(("" + ("foo bar" + "foo bar")) + "")]
                    _ = v79
                }
                _ = len(v6)
                if v3 {
                    v5 += 1
                    switch fi14 {
                    case 4294967295:
                        v5 -= v5
                    default:
                        _ = len(v6)
                    }
                }
                if len(v2) > 0 {
                    v2[0] = v13
                }
                v7 -= v7
                if v3 {
                    v80 := GenS8{F0: (0 < 42)}
                    if (((fi14 * fi14) < (fi14 << ((0) & 63))) || v3) {
                        var v81 string = "a"
                        _ = v81
                    } else {
                        fmt.Println(v5, fi14)
                        var v82 uint8 = 255
                        var v83 bool = v3
                        if (721 > v5) {
                            var v84 int16 = 1000
                            v7 = float64((0.0 - 0.1))
                            if (v7 <= 1.5) {
                                v85 := GenS10{F0: ""}
                                _ = v85
                            }
                            _ = v84
                        }
                        _ = v82
                        _ = v83
                    }
                    _ = len(v1)
                    v86 := GenS9{F0: (255 != 18446744073709551615), F1: 42, F2: -2147483648}
                    _ = v80
                    _ = v86
                } else {
                    v87 := GenS8{F0: v3}
                    v5 = v5
                    _ = v87
                }
                _ = fi20
            }
        default:
            fmt.Println(v13, v13, fi14)
        }
        _ = fi14
    }
    v7 = 0.0
    var v88 int = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v88
}
