package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int32; F1 bool; F2 int64 }
type GenS11 struct { F0 string; F1 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -1
    var v3 float64 = 1e10
    var v4 int16 = v2
    v5 := []uint32{}
    v5 = append(v5, 971)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (-1 == -320) {
        var v6 uint64 = 18446744073709551615
        switch 9223372036854775807 {
        case -457:
            var v7 uint = 141
            _ = v7
        case 127:
            fmt.Println(v2, v2, v2)
        case 9223372036854775807:
            if (false || (435 < 2147483647)) {
                switch (("foo bar" + "a") + "\n\t") {
                case "a":
                    for fi8 := 0; fi8 < 13; fi8++ {
                        var v9 int = 128
                        var v10 uint64 = (v6 | (v6 + (v6 / ((173) | 1))))
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
                var v11 float32 = float32(0.1)
                _ = v11
            }
        default:
            v12 := GenS1{F0: 0, F1: ((v6 >> ((18446744073709551615) & 63)) == (v6 & v6))}
            _ = v12
        }
        var v13 bool = true
        _ = v13
        _ = v6
    }
    if (((1.5 * ((float32((((v3 + 1.5) * 3.14159) + v3)) - 0.0) / (float32(((0.1 - v3) / v3)) / 0.1))) - ((0.0 + (0.0 - 1.5)) / -2.25)) < 1e10) {
        if true {
            for fi14 := 0; fi14 < 19; fi14++ {
                if false {
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v2, v4)
                    case "a":
                        delete(v1, "\n\t")
                    }
                }
                _ = fi14
            }
            switch "a" {
            case "\n\t":
                v4 = 0
            default:
                if true {
                    _ = len(v1)
                    switch "" {
                    case "foo bar":
                        for fi15 := 0; fi15 < 9; fi15++ {
                            var v16 uint8 = 1
                            var v17 int32 = -192
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    case "\n\t":
                        var v18 uint8 = 127
                        _ = v18
                    }
                    if (470 <= -842) {
                        var v19 uint64 = 4294967295
                        switch 438 {
                        case 42:
                            for fi20 := 0; fi20 < 8; fi20++ {
                                if (!(-60 >= 100)) {
                                    if false {
                                        var v21 string = "hello"
                                        fmt.Println(v19, v3)
                                        _ = v21
                                    }
                                } else {
                                    if (235 >= 0) {
                                        v22 := GenS2{F0: true, F1: 1}
                                        v23 := GenS1{F0: -9223372036854775808, F1: (67 == 42)}
                                        _ = v22
                                        _ = v23
                                    } else {
                                        var v24 bool = true
                                        _ = v24
                                    }
                                    v25 := v1["a"]
                                    _ = v25
                                }
                                for fi26 := 0; fi26 < 5; fi26++ {
                                    v19 -= v19
                                    v3 -= v3
                                    var v27 float32 = 0.1
                                    _ = fi26
                                    _ = v27
                                }
                                if len(v5) > 0 {
                                    v5[0] = 0
                                }
                                _ = fi20
                            }
                        case 32767:
                            var v28 uint8 = 127
                            _ = v28
                        case 391:
                            var v29 int64 = 42
                            _ = v29
                        default:
                            var v30 uint8 = 127
                            _ = v30
                        }
                        delete(v1, "hello")
                        _ = v19
                    } else {
                        v31 := GenS3{F0: (!(-59 >= -128)), F1: 834, F2: 156}
                        if true {
                            v5 = append(v5, 287)
                            v32 := GenS1{F0: -1, F1: (127 == -2147483648)}
                            _ = v32
                        } else {
                            v33 := GenS3{F0: (2 <= 4294967295), F1: 42, F2: 18446744073709551615}
                            _ = v33
                        }
                        _ = v31
                    }
                }
            }
        }
        switch "hello" {
        case "":
            fmt.Println(v2, v2)
        case "foo bar":
            fmt.Println(v2)
        case "\n\t":
            if (1 < 1) {
                v34 := GenS3{F0: false, F1: 336, F2: 242}
                v35 := GenS4{F0: "hello"}
                _ = v34
                _ = v35
            } else {
                fmt.Println(v2, v4)
            }
        default:
            for fi36 := 0; fi36 < 15; fi36++ {
                for fi37 := 0; fi37 < 6; fi37++ {
                    v38 := GenS1{F0: 37, F1: (!true)}
                    v39 := GenS5{F0: 4, F1: 746, F2: 2147483648}
                    v4 = int16(fi36)
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                var v40 int8 = -1
                _ = fi36
                _ = v40
            }
        }
        for fi41 := 0; fi41 < 4; fi41++ {
            switch "a" {
            case "":
                v42 := GenS6{F0: 12}
                _ = v42
            }
            for fi43 := 0; fi43 < 1; fi43++ {
                v44 := GenS7{F0: -1}
                switch ("\n\t" + ("\n\t" + "")) {
                case "":
                    v3 = 1e10
                case "a":
                    v3 += 1e10
                case "\n\t":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                default:
                    var v45 uint = 255
                    _ = v45
                }
                fmt.Println(fi41, v2)
                var v46 string = "\n\t"
                _ = fi43
                _ = v44
                _ = v46
            }
            _ = fi41
        }
    } else {
        v47 := GenS8{F0: (-(v4)), F1: 4294967295}
        var v48 uint8 = 255
        _ = v47
        _ = v48
    }
    switch -298 {
    case 42:
        switch 42 {
        case -1:
            var v49 int64 = 42
            _ = v49
        case 4294967295:
            for fi50 := 0; fi50 < 7; fi50++ {
                v51 := GenS4{F0: ""}
                var v52 uint8 = 128
                if true {
                    v53 := GenS9{F0: ("" + (((("a" + ("foo bar" + (("hello" + (("\n\t" + "") + "")) + ("a" + "foo bar")))) + "hello") + "\n\t") + "hello"))}
                    v1["a"] = ("a" + "hello")
                    _ = v53
                }
                _ = fi50
                _ = v51
                _ = v52
            }
        default:
            var v54 int64 = 1
            _ = v54
        }
    case 0:
        v55 := GenS6{F0: v2}
        _ = v55
    }
    for fi56 := 0; fi56 < 18; fi56++ {
        for fi57 := 0; fi57 < 8; fi57++ {
            v58 := GenS10{F0: -1, F1: (2147483647 >= 42), F2: 0}
            if ((!(560 == 1)) || ((v4 & v2) < (v4 % ((0) | 1)))) {
                var v59 uint32 = 401
                _ = v59
            } else {
                v60 := GenS2{F0: false, F1: fi57}
                var v61 bool = true
                v62 := GenS2{F0: true, F1: (fi56 >> ((880) & 63))}
                _ = v60
                _ = v61
                _ = v62
            }
            var v63 float32 = (1.5 * 0.1)
            _ = fi57
            _ = v58
            _ = v63
        }
        _ = fi56
    }
    v4 = (^(v4))
    v64 := GenS9{F0: ("a" + ("foo bar" + (("\n\t" + "hello") + "a")))}
    switch 0 {
    case 42:
        var v65 int16 = (^(v2))
        _ = v65
    case 32767:
        v3 -= (v3 + (((v3 / 1.0) + v3) / ((1.5 + v3) + ((v3 - (1e10 * v3)) * (v3 - (v3 * 0.0))))))
    case 32768:
        v3 = v3
    default:
        if true {
            if true {
                var v66 uint8 = 127
                v3 = (0.1 + ((1.5 * 1.5) - 1e10))
                _ = v66
            }
            v67 := v1[(("foo bar" + "") + "")]
            _ = v67
        }
    }
    v3 = v3
    _ = len(v1)
    if (true && true) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v68 := GenS6{F0: v2}
        _ = v68
    }
    if (!(18446744073709551615 < 2)) {
        switch (("a" + ("hello" + "a")) + "foo bar") {
        case "a":
            var v69 string = ("\n\t" + "foo bar")
            _ = v69
        case "hello":
            v70 := GenS11{F0: ("foo bar" + "foo bar"), F1: 2}
            _ = v70
        case "":
            fmt.Println(v2, v2)
        default:
            var v71 int8 = 0
            _ = v71
        }
        v72 := GenS4{F0: ("foo bar" + ("foo bar" + ("\n\t" + ("\n\t" + "hello"))))}
        _ = v72
    }
    if ("\n\t" == "a") {
        var v73 uint32 = 2147483648
        _ = v73
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
}
