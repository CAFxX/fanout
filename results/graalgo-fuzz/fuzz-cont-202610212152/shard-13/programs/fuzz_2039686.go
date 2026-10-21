package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 bool; F1 int; F2 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 int64; F2 int }
type GenS7 struct { F0 string; F1 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 uint64; F1 string; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 32767
    var v3 bool = false
    var v4 uint8 = 2
    v5 := []string{("" + "hello"), ("" + "\n\t")}
    v6 := []string{(("foo bar" + "") + ""), "a", ("foo bar" + "a")}
    switch -872 {
    case 627:
        v7 := GenS1{F0: 0}
        _ = v7
    case 0:
        if v3 {
            var v8 uint16 = (^(v2))
            v3 = (!true)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v8
        }
    default:
        if false {
            v1[1] = ("foo bar" + "hello")
            var v9 int = 4294967295
            _ = v9
        }
    }
    var v10 int = -878
    var v11 bool = v3
    _ = len(v6)
    var v12 int64 = int64(v4)
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 string = ("foo bar" + "foo bar")
        delete(v1, 2)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        switch (fi15 & (fi15 >> ((1) & 63))) {
        case 4294967295:
            var v16 uint32 = 4294967295
            _ = v16
        case -376:
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 uint8 = v4
                var v19 uint8 = v4
                _ = fi17
                _ = v18
                _ = v19
            }
        default:
            if ((2 > 255) && (2 < 18446744073709551615)) {
                if v3 {
                    var v20 uint32 = 3
                    _ = v20
                }
                var v21 uint16 = (v2 ^ v2)
                _ = v21
            } else {
                if len(v6) > 0 {
                    v6[0] = "foo bar"
                }
                if v11 {
                    fmt.Println(v12, v4, v3)
                    if (true && v11) {
                        var v22 int32 = 2147483647
                        _ = v22
                    }
                }
                v23 := GenS2{F0: 0, F1: (^(v2)), F2: 18446744073709551615}
                _ = v23
            }
        }
        var v24 int8 = 127
        for fi25 := 0; fi25 < 3; fi25++ {
            v3 = v11
            v26 := GenS3{F0: (v4 ^ (^(v4))), F1: fi15}
            var v27 uint8 = (v4 * (^(v4)))
            if true {
                v10 = v10
                switch "\n\t" {
                case "foo bar":
                    if v3 {
                        switch ("hello" + "") {
                        case "a":
                            if (2147483648 > 2) {
                                v3 = (v3 && (3.14159 >= ((-2.25 * (float64((3.14159 + 3.14159)) - float64(((1.5 - 1e10) - 1e10)))) - 1.5)))
                                var v28 int = 693
                                _ = v28
                            }
                        case "foo bar":
                            v3 = true
                        case "hello":
                            delete(v1, 0)
                            fmt.Println(v10, fi25)
                            var v29 int32 = 1
                            _ = v29
                        default:
                            v12 = v12
                        }
                    }
                case "hello":
                    var v30 float64 = -2.25
                    if false {
                        var v31 uint64 = 2
                        _ = v31
                    }
                    switch "a" {
                    case "\n\t":
                        if v11 {
                            v2 -= (v2 >> ((0) & 63))
                            var v32 int64 = 9223372036854775807
                            if (!((^(v2)) == v2)) {
                                v11 = true
                            }
                            _ = v32
                        }
                    case "":
                        switch ("hello" + "hello") {
                        case "hello":
                            v24 += v24
                        case "":
                            var v33 int = fi25
                            _ = v33
                        case "\n\t":
                            if v3 {
                                _ = len(v1)
                                var v34 int8 = (v24 << ((uint(v4)) & 63))
                                _ = v34
                            }
                        default:
                            fmt.Println(v24)
                        }
                    }
                    _ = v30
                }
                if len(v6) > 0 {
                    v6[0] = "a"
                }
            }
            if (v24 < (-(v24))) {
                var v35 uint8 = (v27 | 127)
                fmt.Println(v11)
                if (!v11) {
                    if ((!((^(v27)) < 2)) && ("" < "hello")) {
                        v2 += v2
                    }
                }
                _ = v35
            } else {
                if ((((v2 % (((v2 + (v2 << ((346) & 63)))) | 1)) >= (v2 & v2)) || (true && true)) || v11) {
                    if true {
                        var v36 float32 = ((3.14159 + float32(0.1)) + 3.14159)
                        switch "foo bar" {
                        case "hello":
                            if ((v24 == (-(v24))) && v3) {
                                var v37 bool = (("\n\t" + "a") < "")
                                if (((v27 & v4) == v4) || (v11 || v11)) {
                                    _ = len(v5)
                                    v38 := GenS4{F0: (!(0 == (v12 + v12))), F1: fi15, F2: "\n\t"}
                                    _ = v38
                                }
                                if (v11 && (42 > 682)) {
                                    fmt.Println(v10, v12, v10)
                                } else {
                                    fmt.Println(v12, v3)
                                }
                                v4 = v4
                                var v39 uint = 18446744073709551615
                                switch "\n\t" {
                                case "foo bar":
                                    v40 := GenS5{F0: 65535}
                                    _ = v40
                                case "":
                                    if (512 > 140) {
                                        v41 := GenS6{F0: v2, F1: v12, F2: (^(fi15))}
                                        fmt.Println(v39, v2, v10)
                                        v27 -= (v4 | v27)
                                        v3 = v3
                                        switch "\n\t" {
                                        case "hello":
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        _ = v41
                                    }
                                }
                                v36 += v36
                                var v42 uint16 = v2
                                var v43 int8 = (v24 >> ((65535) & 63))
                                v37 = (("" + "a") >= "\n\t")
                                v44 := GenS7{F0: (("" + ("\n\t" + (("a" + "foo bar") + ""))) + ("foo bar" + "foo bar")), F1: (^(v43))}
                                var v45 float32 = v36
                                if (v39 >= v39) {
                                    v6 = append(v6, (("hello" + ("hello" + "\n\t")) + "a"))
                                    v4 -= v27
                                } else {
                                    var v46 uint16 = 255
                                    v1[2147483648] = ""
                                    _ = v46
                                }
                                _ = v37
                                _ = v39
                                _ = v42
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                        case "":
                            if v11 {
                                delete(v1, 42)
                            }
                        case "\n\t":
                            var v47 int = (fi25 >> ((255) & 63))
                            v6 = append(v6, (((("foo bar" + "a") + "foo bar") + ("hello" + "")) + "foo bar"))
                            var v48 uint16 = (v2 % ((v2) | 1))
                            _ = v47
                            _ = v48
                        }
                        var v49 string = "hello"
                        _ = v36
                        _ = v49
                    } else {
                        if len(v5) > 0 {
                            v5[0] = ""
                        }
                        var v50 int16 = 1
                        switch "" {
                        case "a":
                            v12 += -593
                            delete(v1, 4294967295)
                        case "foo bar":
                            v10 -= -9223372036854775808
                        case "":
                            var v51 float64 = (1.5 * float64((1.5 - 3.14159)))
                            _ = v51
                        }
                        v3 = v3
                        fmt.Println(v10, v24, v12)
                        _ = v50
                    }
                }
                v52 := v1[517]
                v53 := GenS8{F0: 23}
                _ = v52
                _ = v53
            }
            switch (v10 | (fi15 << ((uint(v27)) & 63))) {
            case -400:
                fmt.Println(v24, fi25, v11)
            default:
                if false {
                    if (!(!v3)) {
                        if (v11 && ((v12 - (-(v12))) <= v12)) {
                            var v54 int = fi15
                            v55 := GenS7{F0: ("a" + (("foo bar" + "") + "hello")), F1: (v24 | (v24 & 42))}
                            v54 = (v10 ^ -2147483648)
                            _ = v54
                            _ = v55
                        }
                    }
                    var v56 string = ("" + "a")
                    var v57 float64 = 3.14159
                    if v3 {
                        v58 := GenS9{F0: 255, F1: 0}
                        v27 += (v27 / (((v27 * (v4 - (v4 + (v4 % (((^(v27))) | 1)))))) | 1))
                        _ = v58
                    }
                    v59 := GenS9{F0: 42, F1: 2147483648}
                    _ = v56
                    _ = v57
                    _ = v59
                } else {
                    if (v11 || v3) {
                        v60 := GenS8{F0: (v4 % (((v27 | v27)) | 1))}
                        _ = v60
                    } else {
                        switch fi25 {
                        case 42:
                            delete(v1, 919)
                            v11 = ("\n\t" > "")
                        case 4294967295:
                            switch fi25 {
                            case 127:
                                var v61 float64 = ((float64(1.5) - 1.5) - 3.14159)
                                _ = v61
                            case 32768:
                                v62 := GenS3{F0: (v4 << ((0) & 63)), F1: (v10 % ((1) | 1))}
                                _ = v62
                            default:
                                v4 = (v4 << ((2) & 63))
                            }
                        case -2147483648:
                            v63 := GenS1{F0: 18446744073709551615}
                            _ = v63
                        }
                        var v64 float32 = 3.14159
                        if (((float64(v64) * (0.1 / (1.5 + float64(((v64 / float32((1.5 + 0.0))) - v64))))) - 1e10) < 0.1) {
                            _ = len(v1)
                            v2 -= (v2 % ((2) | 1))
                        } else {
                            var v65 string = "hello"
                            delete(v1, 2147483648)
                            _ = len(v5)
                            _ = len(v5)
                            fmt.Println(v12)
                            _ = v65
                        }
                        _ = len(v6)
                        _ = v64
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi15
        _ = v24
    }
    for fi66 := 0; fi66 < 16; fi66++ {
        v6 = append(v6, "hello")
        fmt.Println(v4, v11)
        _ = fi66
    }
    v3 = v11
    for fi67 := 0; fi67 < 5; fi67++ {
        for fi68 := 0; fi68 < 8; fi68++ {
            switch 4294967295 {
            case 32767:
                fmt.Println(v10, v12)
            }
            v69 := GenS2{F0: -1, F1: (v2 / (((v2 & (^(v2)))) | 1)), F2: 1}
            if v11 {
                var v70 int16 = 332
                v71 := GenS8{F0: uint8(v2)}
                _ = v70
                _ = v71
            } else {
                fmt.Println(v4, fi67)
                if (!(4294967295 <= 964)) {
                    var v72 uint16 = v2
                    _ = v72
                }
                v73 := GenS2{F0: 2147483647, F1: (v2 * v2), F2: 228}
                switch "hello" {
                case "":
                    var v74 int64 = 1
                    _ = v74
                default:
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                }
                _ = v73
            }
            v75 := GenS10{F0: 42, F1: "\n\t", F2: v2}
            fmt.Println(v3)
            var v76 float32 = 0.1
            _ = fi68
            _ = v69
            _ = v75
            _ = v76
        }
        v77 := v1[389]
        var v78 uint16 = (^(v2))
        _ = fi67
        _ = v77
        _ = v78
    }
    for fi79 := 0; fi79 < 3; fi79++ {
        var v80 float32 = 0.1
        _ = fi79
        _ = v80
    }
    fmt.Println(v11, v10, v3)
    switch v10 {
    case -900:
        if v11 {
            for fi81 := 0; fi81 < 2; fi81++ {
                if (54 == 127) {
                    for fi82 := 0; fi82 < 16; fi82++ {
                        v83 := GenS1{F0: 4294967295}
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi82
                        _ = v83
                    }
                }
                v84 := GenS2{F0: -1, F1: (v2 / (((v2 / ((v2) | 1))) | 1)), F2: 4294967295}
                if (true && v3) {
                    v85 := GenS1{F0: 2}
                    v12 += (v12 << ((2) & 63))
                    v86 := GenS9{F0: 1, F1: 662}
                    _ = v85
                    _ = v86
                }
                _ = fi81
                _ = v84
            }
            v87 := GenS2{F0: int32(v2), F1: v2, F2: 0}
            _ = v87
        } else {
            fmt.Println(v10, v10)
            for fi88 := 0; fi88 < 13; fi88++ {
                var v89 int16 = -457
                v90 := GenS4{F0: (((v2 & (v2 / ((v2) | 1))) == v2) || (480 > 42)), F1: (fi88 & v10), F2: "\n\t"}
                _ = fi88
                _ = v89
                _ = v90
            }
        }
    case -9223372036854775808:
        var v91 int16 = 1
        _ = v91
    case 42:
        v4 -= (^(v4))
    default:
        v4 = 61
    }
    switch (v10 / (((v10 % ((-1) | 1))) | 1)) {
    case 0:
        fmt.Println(v3, v3)
    }
    var v92 int32 = -254
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v92
}
