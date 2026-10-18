package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int = 37
    v2 := []string{"a", ("a" + (("a" + "hello") + ((("\n\t" + "\n\t") + "foo bar") + ("hello" + "a"))))}
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    var v6 int32 = 2147483647
    var v7 float64 = 1e10
    for fi8 := 0; fi8 < 3; fi8++ {
        fmt.Println(v1, v7, v1)
        for fi9 := 0; fi9 < 20; fi9++ {
            v7 += float64(0.1)
            if (1 >= -36) {
                var v10 int16 = 1
                _ = v10
            } else {
                var v11 uint32 = 4294967295
                var v12 uint32 = 65535
                if false {
                    switch (fi9 / (((fi9 | (-(fi9)))) | 1)) {
                    case 747:
                        v11 = v12
                    case -1:
                        _ = len(v3)
                    case -727:
                        v13 := GenS1{F0: 42, F1: 1000, F2: "hello"}
                        if (!true) {
                            if false {
                                v11 -= 4294967295
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                                v14 := GenS1{F0: 1, F1: -900, F2: "foo bar"}
                                _ = v14
                            }
                            fmt.Println(v1, v1)
                        } else {
                            var v15 uint8 = 128
                            v1 = fi9
                            _ = v15
                        }
                        _ = v13
                    }
                } else {
                    v7 = 0.1
                    if (!(2 != 994)) {
                        v7 += (((v7 + (0.0 * (((v7 * v7) - 1e10) + v7))) * v7) + v7)
                        var v16 int = fi9
                        _ = len(v2)
                        _ = v16
                    }
                    if true {
                        fmt.Println(fi9)
                    } else {
                        v6 = v6
                        var v17 uint16 = 911
                        v18 := v4[fi8]
                        _ = v17
                        _ = v18
                    }
                }
                var v19 uint32 = 566
                fmt.Println(v6, fi8, v6)
                _ = v11
                _ = v12
                _ = v19
            }
            fmt.Println(fi9, v7)
            _ = fi9
        }
        if false {
            var v20 float32 = 1.5
            switch 672 {
            case 0:
                fmt.Println(v6, v7)
            case 4294967295:
                switch fi8 {
                case 0:
                    for fi21 := 0; fi21 < 2; fi21++ {
                        var v22 string = "\n\t"
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi21
                        _ = v22
                    }
                case 127:
                    v23 := GenS1{F0: -822, F1: 1000, F2: "\n\t"}
                    _ = v23
                case 32768:
                    fmt.Println(v20, v7)
                }
            default:
                v24 := GenS1{F0: 515, F1: 1, F2: ("hello" + ((("a" + ("" + ((("hello" + ("hello" + "hello")) + "hello") + (((("\n\t" + ((("" + ("a" + ("a" + "a"))) + "\n\t") + "hello")) + ("a" + ("\n\t" + "\n\t"))) + "a") + "foo bar")))) + "hello") + ("hello" + "\n\t")))}
                _ = v24
            }
            var v25 bool = (3.14159 < (-2.25 * v7))
            _ = v20
            _ = v25
        }
        _ = fi8
    }
    if false {
        _ = len(v2)
        v26 := GenS2{F0: 100}
        fmt.Println(v7, v7, v7)
        _ = v26
    }
    v27 := GenS1{F0: 105, F1: 1000, F2: "foo bar"}
    v7 += v7
    v1 += v1
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 int = (fi28 + fi28)
        for fi30 := 0; fi30 < 16; fi30++ {
            switch ("foo bar" + (("hello" + ("foo bar" + (("foo bar" + "hello") + "\n\t"))) + ("hello" + "hello"))) {
            case "hello":
                var v31 float64 = v7
                v2 = append(v2, "a")
                v32 := v5[((("hello" + ("foo bar" + "")) + ("\n\t" + "\n\t")) + "a")]
                _ = v31
                _ = v32
            default:
                var v33 int64 = -1
                _ = v33
            }
            var v34 int8 = 42
            v35 := v4[fi28]
            _ = fi30
            _ = v34
            _ = v35
        }
        if true {
            v36 := GenS2{F0: 100}
            for fi37 := 0; fi37 < 14; fi37++ {
                v3[128] = 2
                _ = fi37
            }
            var v38 string = ("foo bar" + ("a" + "hello"))
            _ = v36
            _ = v38
        } else {
            var v39 int = 128
            _ = v39
        }
        _ = fi28
        _ = v29
    }
    v6 += (v6 >> ((42) & 63))
    if (v6 <= -2147483648) {
        var v40 int8 = -109
        var v41 int32 = 83
        _ = v40
        _ = v41
    } else {
        v42 := GenS2{F0: 1}
        if true {
            fmt.Println(v6)
            var v43 int16 = -1
            _ = v43
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi44 := 0; fi44 < 13; fi44++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v1, v7, v1)
                var v45 uint32 = 677
                _ = fi44
                _ = v45
            }
            if false {
                delete(v4, (v1 & (v1 % (((v1 / (((v1 >> ((976) & 63))) | 1))) | 1))))
                for fi46 := 0; fi46 < 5; fi46++ {
                    switch v1 {
                    case 32767:
                        v47 := v5["hello"]
                        _ = v47
                    case -1:
                        var v48 uint64 = 18446744073709551615
                        _ = v48
                    case 128:
                        switch fi46 {
                        case -2147483648:
                            var v49 uint = 703
                            _ = v49
                        case 2147483647:
                            if (!(0 != 32767)) {
                                v7 = v7
                            }
                        case -22:
                            switch fi46 {
                            case -1:
                                if (0 == 4294967295) {
                                    v7 += v7
                                } else {
                                    fmt.Println(v1, fi46, v6)
                                    fmt.Println(fi46, v1, v6)
                                    switch ("" + "\n\t") {
                                    case "":
                                        var v50 uint = 42
                                        _ = v50
                                    case "\n\t":
                                        switch "" {
                                        case "":
                                            v7 = ((1e10 / (v7 + 0.0)) + ((float64((0.1 - (1.5 + (-2.25 / 1.0)))) * (float64(0.1) / v7)) - ((v7 - 0.1) * v7)))
                                        case "foo bar":
                                            v7 -= (float64(0.1) + v7)
                                        case "\n\t":
                                            for fi51 := 0; fi51 < 5; fi51++ {
                                                if true {
                                                    if (!(255 != 255)) {
                                                        var v52 int16 = -897
                                                        switch "" {
                                                        case "\n\t":
                                                            v7 = 1e10
                                                        case "hello":
                                                            if (919 >= 2147483648) {
                                                                v3[fi46] = 352
                                                                v2 = append(v2, "\n\t")
                                                                delete(v5, "a")
                                                            }
                                                        default:
                                                            fmt.Println(v1, fi51)
                                                        }
                                                        switch "foo bar" {
                                                        case "a":
                                                            fmt.Println(v7)
                                                        case "":
                                                            v1 += (fi46 / ((fi46) | 1))
                                                        case "hello":
                                                            v6 = v6
                                                        }
                                                        _ = v52
                                                    } else {
                                                        v1 -= v1
                                                        var v53 float64 = (v7 * -2.25)
                                                        v54 := GenS2{F0: 0}
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    if (!false) {
                                                        v7 += 3.14159
                                                    }
                                                }
                                                switch fi46 {
                                                case 127:
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if (float64(1.5) > -2.25) {
                                                            v6 += v6
                                                            v55 := GenS2{F0: -128}
                                                            _ = v55
                                                        } else {
                                                            switch "a" {
                                                            case "foo bar":
                                                                var v56 string = "hello"
                                                                _ = v56
                                                            case "hello":
                                                                v7 = (v7 + 0.0)
                                                            default:
                                                                v6 = 780
                                                                var v57 int64 = -1
                                                                _ = v57
                                                            }
                                                        }
                                                    case "":
                                                        var v58 uint = 4294967295
                                                        _ = v58
                                                    default:
                                                        v3[4294967295] = 4294967295
                                                    }
                                                case -295:
                                                    v2 = append(v2, "")
                                                case -9223372036854775808:
                                                    if (1000 > -1) {
                                                        delete(v3, (fi46 << ((255) & 63)))
                                                        var v59 bool = true
                                                        if len(v2) > 0 {
                                                            v2[0] = ((("" + (("hello" + ("\n\t" + (((("foo bar" + "foo bar") + "hello") + ("foo bar" + "a")) + "\n\t"))) + "a")) + ("foo bar" + "hello")) + (("foo bar" + "hello") + ("hello" + ((("foo bar" + "foo bar") + ("\n\t" + ("a" + ("foo bar" + ("hello" + ("" + "hello")))))) + "foo bar"))))
                                                        }
                                                        _ = v59
                                                    }
                                                default:
                                                    if (2 >= 128) {
                                                        fmt.Println(v1, v7, v7)
                                                    } else {
                                                        if (!(v1 > 127)) {
                                                            v60 := GenS3{F0: v6}
                                                            v1 -= fi51
                                                            _ = v60
                                                        } else {
                                                            if ((v6 % ((v6) | 1)) == -1) {
                                                                fmt.Println(v7)
                                                                switch (fi51 - -879) {
                                                                case -2147483648:
                                                                    var v61 int8 = -1
                                                                    _ = v61
                                                                case 127:
                                                                    var v62 uint16 = 42
                                                                    _ = v62
                                                                default:
                                                                    if (215 >= 65535) {
                                                                        v1 -= 128
                                                                        if (true && (v6 < (^(v6)))) {
                                                                            v7 = v7
                                                                        }
                                                                        if true {
                                                                            if len(v2) > 0 {
                                                                                v2[0] = "a"
                                                                            }
                                                                        } else {
                                                                            v63 := GenS4{F0: 609}
                                                                            _ = v63
                                                                        }
                                                                    } else {
                                                                        v1 = -995
                                                                    }
                                                                }
                                                            } else {
                                                                v64 := v4[-2147483648]
                                                                _ = len(v2)
                                                                _ = v64
                                                            }
                                                            v6 += (v6 & (v6 >> ((4294967295) & 63)))
                                                        }
                                                        if true {
                                                            v1 -= (fi51 - fi46)
                                                            v1 = v1
                                                            var v65 int16 = 0
                                                            v65 += (v65 << ((18446744073709551615) & 63))
                                                            if false {
                                                                _ = len(v5)
                                                                var v66 float64 = (float64(-2.25) / v7)
                                                                switch fi51 {
                                                                case -2147483648:
                                                                    v6 -= v6
                                                                case 42:
                                                                    if (2 != 361) {
                                                                        v7 -= 0.1
                                                                        v66 += v66
                                                                    } else {
                                                                        var v67 float64 = ((1.5 + v66) * float64(0.1))
                                                                        v6 = (v6 | (v6 ^ (v6 & v6)))
                                                                        v66 -= 0.1
                                                                        v68 := v5["\n\t"]
                                                                        v66 -= 3.14159
                                                                        _ = v67
                                                                        _ = v68
                                                                    }
                                                                case 32767:
                                                                    v69 := GenS5{F0: 161, F1: false, F2: 32767}
                                                                    _ = v69
                                                                default:
                                                                    if false {
                                                                        var v70 string = ""
                                                                        v7 = 1.5
                                                                        _ = v70
                                                                    } else {
                                                                        v7 = (1.5 - v66)
                                                                        v66 += (3.14159 * float64(1e10))
                                                                        fmt.Println(v1)
                                                                    }
                                                                    if ((("a" + "a") != "a") || false) {
                                                                        v1 = (fi51 % ((fi46) | 1))
                                                                        var v71 float64 = ((0.1 + float64((((3.14159 / 3.14159) + 1.5) / (3.14159 - 1e10)))) / 1.0)
                                                                        if ((32767 == 32767) && false) {
                                                                            delete(v3, -2147483648)
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (("\n\t" + "foo bar") + ("" + "a"))
                                                                            }
                                                                        }
                                                                        _ = v71
                                                                    } else {
                                                                        v2 = append(v2, "a")
                                                                    }
                                                                    fmt.Println(fi51)
                                                                }
                                                                _ = v66
                                                            }
                                                            var v72 int32 = (v6 - (v6 & v6))
                                                            var v73 int = (^(v1))
                                                            _ = v65
                                                            _ = v72
                                                            _ = v73
                                                        } else {
                                                            var v74 float32 = 1.5
                                                            _ = v74
                                                        }
                                                        if false {
                                                            fmt.Println(v7)
                                                        } else {
                                                            if true {
                                                                var v75 float32 = (0.0 / 0.1)
                                                                _ = v75
                                                            }
                                                            var v76 string = "foo bar"
                                                            fmt.Println(v6)
                                                            var v77 int16 = 0
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        if false {
                                                            v7 += v7
                                                        }
                                                        var v78 bool = (false || true)
                                                        var v79 uint = 42
                                                        v6 = -1
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                }
                                                if (("hello" + "foo bar") <= "a") {
                                                    var v80 int8 = 42
                                                    v1 += (v1 | fi46)
                                                    v7 = v7
                                                    _ = v80
                                                } else {
                                                    v7 = (((v7 * (0.1 - v7)) * v7) - float64(float32(-2.25)))
                                                    v7 = v7
                                                    v7 = (float64(3.14159) * v7)
                                                }
                                                _ = fi51
                                            }
                                        }
                                    default:
                                        for fi81 := 0; fi81 < 1; fi81++ {
                                            switch -2147483648 {
                                            case 1:
                                                v82 := GenS3{F0: -904}
                                                _ = v82
                                            case 32768:
                                                var v83 string = ("a" + (((("\n\t" + "\n\t") + ("foo bar" + "\n\t")) + "hello") + ("hello" + "hello")))
                                                _ = v83
                                            case 9223372036854775807:
                                                var v84 float64 = float64(0.1)
                                                _ = v84
                                            default:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            var v85 uint16 = 0
                                            _ = fi81
                                            _ = v85
                                        }
                                    }
                                }
                            }
                        }
                    }
                    var v86 int16 = 32767
                    _ = fi46
                    _ = v86
                }
            }
        }
        _ = v42
    }
    v5[(("" + "") + ("foo bar" + "hello"))] = (!false)
    v5[("" + "")] = ((((float32(v7) - 1.5) / (((-2.25 - (0.0 + 1e10)) - ((1e10 / 1.5) + float32(((0.1 + (3.14159 * float64(3.14159))) / 1.0)))) / 1e10)) * 1.5) >= 3.14159)
    switch v1 {
    case 9223372036854775807:
        var v87 float32 = float32(1.5)
        _ = v87
    }
    v88 := GenS5{F0: 42, F1: false, F2: 480}
    v89 := GenS6{F0: 282, F1: 1}
    v90 := GenS2{F0: -128}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v88
    _ = v89
    _ = v90
}
