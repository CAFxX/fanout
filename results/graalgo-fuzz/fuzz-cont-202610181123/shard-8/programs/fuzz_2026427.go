package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int8; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint64; F1 int32; F2 string }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64; F1 int8 }
type GenS10 struct { F0 uint; F1 uint; F2 int32 }
func genFn0(p0 bool, p1 string, p2 int) bool {
    fmt.Println(p2)
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: p1, F1: 42}
        for fi9 := 0; fi9 < 11; fi9++ {
            p0 = true
            if p0 {
                if (4294967295 != 425) {
                    var v10 uint64 = 2
                    v11 := GenS2{F0: 1, F1: v10, F2: 42}
                    fmt.Println(p0, fi9, p1)
                    _ = v10
                    _ = v11
                }
            } else {
                var v12 int = -896
                _ = v12
            }
            fmt.Println(p2, fi9, p0)
            _ = fi9
        }
        for fi13 := 0; fi13 < 16; fi13++ {
            if true {
                if false {
                    if ((-2.25 - 1.5) >= ((0.0 / 1.5) - -2.25)) {
                        var v14 float64 = 1e10
                        _ = v14
                    }
                    switch ((p1 + p1) + p1) {
                    case "hello":
                        if false {
                            p1 = "foo bar"
                        }
                        if (1.5 == -2.25) {
                            p0 = (1e10 < -2.25)
                        }
                    default:
                        v15 := []int{fi7, 127, (fi13 | (fi7 * fi7)), (-(p2))}
                        _ = v15
                    }
                    var v16 float32 = 0.1
                    _ = v16
                }
                switch p1 {
                case "":
                    v17 := GenS1{F0: "", F1: -10}
                    v18 := GenS3{F0: 1, F1: 1, F2: "hello"}
                    _ = v17
                    _ = v18
                case "hello":
                    switch p2 {
                    case 127:
                        var v19 bool = (-1 > -1)
                        p2 -= -191
                        _ = v19
                    case -977:
                        p1 = "\n\t"
                        p2 -= 128
                    }
                }
            } else {
                var v20 uint16 = 32767
                _ = v20
            }
            _ = fi13
        }
        _ = fi7
        _ = v8
    }
    var v21 float64 = (((3.14159 + 0.1) / (1.5 * (-2.25 - 1e10))) / float64((0.0 - (1e10 + 1.5))))
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return false
}

func main() {
    var v1 uint16 = 1
    var v2 bool = (-1 <= -2147483648)
    var v3 bool = true
    v4 := make(map[string]bool)
    var v5 uint64 = 2
    v6 := []uint32{2147483648, 2147483647, 65535}
    fmt.Println(v2)
    v22 := GenS2{F0: -53, F1: (v5 | v5), F2: 42}
    if (127 != 128) {
        fmt.Println(v1)
        for fi23 := 0; fi23 < 4; fi23++ {
            _ = len(v6)
            var v24 uint8 = 255
            _ = fi23
            _ = v24
        }
    } else {
        fmt.Println(v5, v3, v1)
        for fi25 := 0; fi25 < 16; fi25++ {
            switch 128 {
            case 1:
                v5 = (v5 / ((42) | 1))
            case 4294967295:
                v6 = append(v6, 42)
            default:
                if true {
                    if (!(255 != 65535)) {
                        v5 -= (^(v5))
                        var v26 string = ("hello" + "foo bar")
                        _ = v26
                    } else {
                        _ = len(v6)
                        fmt.Println(v3)
                        switch "\n\t" {
                        case "foo bar":
                            var v27 bool = ((v5 > (v5 * (v5 - (v5 * uint64(v1))))) || v3)
                            _ = v27
                        case "hello":
                            if v2 {
                                v28 := GenS1{F0: ("foo bar" + "hello"), F1: 210}
                                _ = v28
                            }
                        }
                    }
                    switch (((("foo bar" + "") + (("a" + "a") + ("foo bar" + "a"))) + "") + "\n\t") {
                    case "hello":
                        delete(v4, "hello")
                    case "":
                        v1 = (v1 >> ((18446744073709551615) & 63))
                    case "\n\t":
                        var v29 int16 = 42
                        _ = v29
                    default:
                        fmt.Println(v1, v5)
                    }
                }
            }
            v30 := GenS1{F0: "", F1: 2147483647}
            _ = fi25
            _ = v30
        }
        v31 := GenS3{F0: v5, F1: -1, F2: "foo bar"}
        _ = v31
    }
    if v3 {
        v32 := GenS3{F0: v5, F1: 653, F2: "foo bar"}
        _ = v32
    }
    if (!v3) {
        if (-2.25 <= (1.5 / 1.0)) {
            v1 = (v1 - 32767)
            for fi33 := 0; fi33 < 14; fi33++ {
                switch fi33 {
                case -480:
                    for fi34 := 0; fi34 < 10; fi34++ {
                        fmt.Println(v2, v1)
                        var v35 uint16 = 1
                        _ = fi34
                        _ = v35
                    }
                case -2147483648:
                    var v36 uint8 = 127
                    _ = v36
                }
                if v2 {
                    if v3 {
                        if v2 {
                            v1 = (v1 >> ((112) & 63))
                            for fi37 := 0; fi37 < 2; fi37++ {
                                v5 = v5
                                if (-1 > 42) {
                                    var v38 int32 = 1
                                    var v39 int32 = v38
                                    v1 = v1
                                    _ = v38
                                    _ = v39
                                } else {
                                    v3 = (v3 || (4294967295 <= (^(v5))))
                                    v40 := GenS4{F0: (v3 || (v5 >= (^(v5)))), F1: -1}
                                    v41 := GenS5{F0: (v1 | v1), F1: (v5 * (v5 ^ 42)), F2: 746}
                                    var v42 int64 = 9223372036854775807
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                                v43 := GenS6{F0: -1}
                                var v44 uint64 = 42
                                if len(v6) > 0 {
                                    v6[0] = 115
                                }
                                _ = fi37
                                _ = v43
                                _ = v44
                            }
                            v1 -= 42
                        }
                    }
                } else {
                    var v45 bool = v3
                    if ("a" < "foo bar") {
                        fmt.Println(v3, v3)
                        fmt.Println(v2, v2, v3)
                    }
                    v46 := v4["hello"]
                    _ = v45
                    _ = v46
                }
                v47 := GenS7{F0: uint32(v5), F1: uint32(v1)}
                _ = fi33
                _ = v47
            }
        } else {
            v48 := GenS8{F0: 42}
            var v49 uint32 = 131
            v49 += v49
            _ = v48
            _ = v49
        }
        if v3 {
            v50 := GenS5{F0: (v1 ^ (^(v1))), F1: (^(v5)), F2: 65535}
            v5 -= (v5 ^ (v5 % ((v5) | 1)))
            v51 := GenS5{F0: 605, F1: 555, F2: 65535}
            _ = v50
            _ = v51
        }
        if true {
            var v52 int32 = 2147483647
            fmt.Println(v52)
            fmt.Println(v5, v1, v3)
            _ = v52
        } else {
            if len(v6) > 0 {
                v6[0] = uint32(v5)
            }
        }
    }
    fmt.Println(v1, v3)
    if v3 {
        v5 += uint64(v1)
        var v53 uint16 = v1
        for fi54 := 0; fi54 < 8; fi54++ {
            if v3 {
                for fi55 := 0; fi55 < 5; fi55++ {
                    v56 := GenS7{F0: 2147483648, F1: 2147483647}
                    v57 := GenS9{F0: 42, F1: 27}
                    _ = fi55
                    _ = v56
                    _ = v57
                }
                if v3 {
                    if false {
                        for fi58 := 0; fi58 < 10; fi58++ {
                            switch "hello" {
                            case "foo bar":
                                var v59 uint = 1
                                _ = v59
                            case "\n\t":
                                var v60 int32 = -559
                                _ = v60
                            case "hello":
                                v61 := GenS10{F0: 671, F1: 0, F2: -80}
                                _ = v61
                            }
                            _ = fi58
                        }
                        fmt.Println(v5, v53, v53)
                    }
                    var v62 uint = 42
                    _ = v62
                }
                delete(v4, ("foo bar" + "a"))
            }
            v3 = (((0.0 / 0.1) + 3.14159) >= 1e10)
            _ = fi54
        }
        _ = v53
    } else {
        v3 = (!v3)
    }
    switch -12 {
    case 473:
        v5 = (^(v5))
    case -50:
        if (0.0 != 0.1) {
            v2 = v3
            for fi63 := 0; fi63 < 4; fi63++ {
                fmt.Println(v1, v2)
                var v64 int64 = 0
                _ = fi63
                _ = v64
            }
            var v65 uint8 = 42
            _ = v65
        } else {
            if (410 != 1000) {
                var v66 float64 = 0.1
                var v67 int16 = 0
                _ = len(v6)
                _ = v66
                _ = v67
            }
            v2 = ((((128 < 0) && (0 >= 4294967295)) || (255 < 42)) || v3)
            switch (("foo bar" + (("\n\t" + ("\n\t" + "hello")) + "hello")) + "\n\t") {
            case "foo bar":
                if ("hello" < ("hello" + ("a" + ""))) {
                    fmt.Println(v5, v5, v1)
                    var v68 uint32 = 930
                    if (587 != (^(v1))) {
                        var v69 uint8 = 221
                        _ = v69
                    } else {
                        var v70 float32 = 3.14159
                        fmt.Println(v70, v68, v70)
                        var v71 uint64 = v5
                        _ = v70
                        _ = v71
                    }
                    _ = v68
                } else {
                    switch 32768 {
                    case -642:
                        v5 += (v5 << ((42) & 63))
                    }
                }
            case "a":
                if ((((2 <= 396) && (2 != 42)) || v2) || ((-115 == 42) && (!v3))) {
                    if false {
                        for fi72 := 0; fi72 < 7; fi72++ {
                            for fi73 := 0; fi73 < 1; fi73++ {
                                if false {
                                    v1 += (v1 >> ((4294967295) & 63))
                                    var v74 int32 = -312
                                    v1 += v1
                                    _ = v74
                                }
                                if false {
                                    var v75 float64 = 0.0
                                    v4["hello"] = (!(v3 && true))
                                    _ = v75
                                } else {
                                    v76 := GenS3{F0: 0, F1: -314, F2: (("hello" + (("\n\t" + "") + "foo bar")) + "a")}
                                    _ = v76
                                }
                                _ = fi73
                            }
                            fmt.Println(v1, v5)
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = fi72
                        }
                    }
                }
            case "":
                v77 := GenS5{F0: v1, F1: v5, F2: 2147483647}
                _ = v77
            default:
                var v78 uint16 = v1
                _ = v78
            }
        }
    case -9223372036854775808:
        var v79 uint32 = 831
        _ = v79
    }
    fmt.Println(v5, v1, v3)
    if (!(2 <= 729)) {
        switch 927 {
        case 128:
            for fi80 := 0; fi80 < 1; fi80++ {
                if len(v6) > 0 {
                    v6[0] = 585
                }
                switch ("hello" + ("foo bar" + "foo bar")) {
                case "a":
                    var v81 float64 = (3.14159 - 0.0)
                    _ = v81
                }
                _ = fi80
            }
        default:
            v82 := GenS2{F0: -128, F1: (^(v5)), F2: -644}
            _ = v82
        }
        switch -2147483648 {
        case 127:
            var v83 uint32 = 958
            _ = v83
        default:
            if (!((0 < -23) || true)) {
                switch 2147483647 {
                case 128:
                    if true {
                        v84 := GenS3{F0: uint64(v1), F1: 1, F2: ("" + (("hello" + "") + ("foo bar" + "a")))}
                        var v85 int16 = -387
                        _ = v84
                        _ = v85
                    } else {
                        _ = len(v6)
                        var v86 float64 = (0.1 / -2.25)
                        _ = v86
                    }
                }
            }
        }
        v3 = (!(((42 < 896) || v2) || (("\n\t" + "a") != "")))
    } else {
        switch 32767 {
        case 42:
            var v87 int16 = -32768
            _ = v87
        case 1:
            switch "" {
            case "":
                var v88 uint8 = 0
                _ = v88
            }
        default:
            v4[(("foo bar" + "\n\t") + ("a" + (("\n\t" + "\n\t") + "\n\t")))] = v2
        }
    }
    var v89 bool = (!v3)
    var v90 uint32 = 42
    var v91 uint32 = 42
    var v92 uint = uint(v1)
    v1 = (v1 ^ 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v90
    _ = v91
    _ = v92
}
