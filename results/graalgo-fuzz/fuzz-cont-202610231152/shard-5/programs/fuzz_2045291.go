package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint; F1 uint; F2 int64 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 int64; F2 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string }
func genFn0(p0 bool, p1 int32, p2 uint32) string {
    var v4 int64 = 0
    if p0 {
        var v5 int = 128
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 1; fi7++ {
                fmt.Println(p2)
                _ = fi7
            }
            var v8 int32 = (p1 >> ((709) & 63))
            v9 := GenS1{F0: v4, F1: 152, F2: 255}
            _ = fi6
            _ = v8
            _ = v9
        }
        _ = v5
    } else {
        switch 4294967295 {
        case -2147483648:
            fmt.Println(v4, v4)
        }
        v10 := []int{-9223372036854775808, -2147483648, 127}
        fmt.Println(p2, p1)
        _ = v10
    }
    fmt.Println(p1)
    if p0 {
        p0 = (-9223372036854775808 == 72)
    } else {
        switch 128 {
        case 32768:
            var v11 uint = 255
            _ = v11
        case -9223372036854775808:
            p0 = true
        case 0:
            if ((p2 <= (^(p2))) && p0) {
                var v12 int32 = (-(p1))
                if (!(((v4 & (v4 * v4)) > v4) && (-740 > (v4 + 786)))) {
                    var v13 uint = 18446744073709551615
                    _ = v13
                }
                _ = v12
            }
        default:
            fmt.Println(p2, p2)
        }
        p0 = (!(("foo bar" + "") <= "a"))
    }
    p1 = (p1 >> ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "\n\t"
}

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 1
    v3 := make(map[string]bool)
    v2 = (v2 % (((v2 + 1)) | 1))
    v14 := []int{int(v1)}
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS1{F0: -9223372036854775808, F1: v1, F2: (v1 ^ v1)}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        fmt.Println(v1, v1, v1)
        v18 := GenS1{F0: 9223372036854775807, F1: v1, F2: (^(v1))}
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 1, F1: v1, F2: 0}
    if ("" <= (("\n\t" + "\n\t") + "foo bar")) {
        var v20 int64 = 457
        _ = len(v14)
        _ = v20
    } else {
        delete(v3, ("\n\t" + "\n\t"))
    }
    if true {
        v2 = (v2 & 1)
        for fi21 := 0; fi21 < 19; fi21++ {
            for fi22 := 0; fi22 < 12; fi22++ {
                if false {
                    fmt.Println(fi22, v1)
                    switch "hello" {
                    case "":
                        var v23 uint64 = uint64(v1)
                        _ = v23
                    case "a":
                        switch "\n\t" {
                        case "a":
                            var v24 uint16 = 255
                            _ = v24
                        default:
                            fmt.Println(v1, v2)
                        }
                    }
                }
                if ((fi21 | fi21) < fi22) {
                    v25 := GenS1{F0: -946, F1: 194, F2: v1}
                    if (v1 >= 128) {
                        var v26 uint16 = 255
                        _ = v26
                    } else {
                        v27 := GenS2{F0: (v1 + (v1 - (v1 ^ v1))), F1: 12}
                        _ = v27
                    }
                    switch ("foo bar" + ("hello" + "\n\t")) {
                    case "\n\t":
                        if ((255 < 211) && true) {
                            if true {
                                fmt.Println(fi21, v2)
                                v14 = append(v14, fi22)
                            } else {
                                v1 += (v1 % (((v1 * (v1 - v1))) | 1))
                            }
                            v28 := GenS2{F0: (v1 + 0), F1: 100}
                            switch (fi21 * 359) {
                            case -9223372036854775808:
                                v29 := v3["hello"]
                                _ = v29
                            default:
                                v30 := GenS3{F0: 4294967295, F1: 999, F2: -9223372036854775808}
                                _ = v30
                            }
                            v1 -= 51
                            _ = v28
                        }
                    case "foo bar":
                        fmt.Println(fi22, fi22)
                    }
                    v31 := GenS2{F0: v1, F1: 1}
                    fmt.Println(v1, v2)
                    if (uint8(v2) >= v1) {
                        var v32 int = (fi22 >> ((906) & 63))
                        var v33 int = 152
                        _ = v32
                        _ = v33
                    }
                    _ = v25
                    _ = v31
                }
                var v34 int16 = 623
                v35 := v3["foo bar"]
                if true {
                    var v36 string = ("" + ("hello" + "\n\t"))
                    _ = v36
                }
                _ = fi22
                _ = v34
                _ = v35
            }
            if false {
                var v37 float32 = ((0.0 + (0.0 - ((float32((-2.25 / (-2.25 - 0.0))) + 3.14159) - float32(((0.0 + 1e10) - -2.25))))) - 0.0)
                _ = v37
            }
            _ = fi21
        }
        if false {
            for fi38 := 0; fi38 < 3; fi38++ {
                switch "" {
                case "hello":
                    for fi39 := 0; fi39 < 9; fi39++ {
                        if true {
                            v2 = (v2 / ((v2) | 1))
                            var v40 int16 = 42
                            _ = v40
                        } else {
                            if (-678 <= 32767) {
                                var v41 bool = false
                                _ = v41
                            }
                        }
                        var v42 int64 = 730
                        _ = fi39
                        _ = v42
                    }
                default:
                    if true {
                        v1 = (v1 << ((0) & 63))
                        v14 = append(v14, 4294967295)
                    } else {
                        if (0.1 == 1e10) {
                            fmt.Println(v2, fi38, fi38)
                        } else {
                            v1 = (^(v1))
                            _ = len(v3)
                        }
                    }
                }
                _ = fi38
            }
        }
    } else {
        fmt.Println(v2)
        var v43 uint = 739
        _ = v43
    }
    v1 += (v1 << ((902) & 63))
    v1 = (v1 + v1)
    for fi44 := 0; fi44 < 2; fi44++ {
        fmt.Println(v2, v2)
        var v45 float64 = 1e10
        _ = fi44
        _ = v45
    }
    v46 := GenS3{F0: 128, F1: 150, F2: 42}
    fmt.Println(v2, v1)
    switch 32767 {
    case 127:
        v1 = (v1 / ((v1) | 1))
    case 9223372036854775807:
        if (42 != 0) {
            v14 = append(v14, 128)
            v1 += (^(v1))
            switch -2147483648 {
            case 128:
                if (622 == -2147483648) {
                    v2 += 292
                    v2 -= 1
                    var v47 uint64 = (v2 - 741)
                    _ = v47
                } else {
                    switch 0 {
                    case 970:
                        var v48 bool = true
                        _ = v48
                    case -448:
                        v49 := GenS4{F0: ("hello" + "a"), F1: "\n\t"}
                        _ = v49
                    case 637:
                        v50 := GenS5{F0: ("hello" + "\n\t")}
                        _ = v50
                    default:
                        v3["hello"] = false
                    }
                }
            default:
                v1 = v1
            }
        } else {
            v3["a"] = (false || (false && (0 <= 0)))
            switch ("foo bar" + (("" + "") + ("hello" + ""))) {
            case "foo bar":
                var v51 uint16 = 624
                _ = v51
            case "hello":
                v52 := GenS6{F0: v2, F1: -1, F2: false}
                _ = v52
            }
            var v53 int8 = 1
            _ = v53
        }
    case 429:
        switch (((("" + ("foo bar" + "hello")) + ("foo bar" + ("foo bar" + "foo bar"))) + ("\n\t" + "a")) + "foo bar") {
        case "hello":
            var v54 uint32 = 2147483647
            _ = v54
        case "foo bar":
            var v55 string = "a"
            _ = v55
        case "\n\t":
            v2 += v2
        default:
            if ((false && (-687 < 42)) && (127 < 0)) {
                v1 = v1
                if (32767 >= 1000) {
                    v1 = 127
                } else {
                    if (1 <= 1) {
                        var v56 int8 = 81
                        if true {
                            v1 = v1
                            var v57 string = "hello"
                            var v58 int8 = v56
                            _ = v57
                            _ = v58
                        }
                        switch "\n\t" {
                        case "foo bar":
                            v2 = (v2 * 353)
                        case "":
                            delete(v3, "a")
                        default:
                            switch -2147483648 {
                            case -954:
                                fmt.Println(v56)
                            case 42:
                                v59 := GenS7{F0: "\n\t"}
                                _ = v59
                            case 370:
                                _ = len(v3)
                            }
                        }
                        _ = v56
                    }
                    v60 := GenS5{F0: "a"}
                    var v61 bool = (("foo bar" + ("foo bar" + (("foo bar" + (("foo bar" + "") + "")) + "hello"))) != "a")
                    _ = v60
                    _ = v61
                }
                switch (("\n\t" + "hello") + "a") {
                case "hello":
                    _ = len(v14)
                case "foo bar":
                    v62 := v3["a"]
                    _ = v62
                }
            }
        }
    }
    fmt.Println(v1)
    for fi63 := 0; fi63 < 20; fi63++ {
        if (1 <= 1) {
            v1 = 211
            delete(v3, ("" + "foo bar"))
            var v64 int64 = -212
            _ = v64
        }
        if true {
            for fi65 := 0; fi65 < 3; fi65++ {
                if (v2 == (v2 + 42)) {
                    var v66 int = (fi65 << ((65535) & 63))
                    if len(v14) > 0 {
                        v14[0] = fi63
                    }
                    _ = v66
                }
                _ = fi65
            }
            v67 := v3["\n\t"]
            _ = v67
        }
        var v68 uint32 = 42
        _ = fi63
        _ = v68
    }
    for fi69 := 0; fi69 < 19; fi69++ {
        switch "foo bar" {
        case "":
            v70 := GenS3{F0: 442, F1: 65535, F2: 1}
            _ = v70
        default:
            for fi71 := 0; fi71 < 16; fi71++ {
                fmt.Println(fi69, fi71)
                if false {
                    v72 := GenS4{F0: ("hello" + ""), F1: ("" + ("\n\t" + "hello"))}
                    _ = v72
                } else {
                    v1 = (v1 << ((1) & 63))
                    if ((true || true) || (-128 == -1)) {
                        switch (fi69 % (((fi71 | -2147483648)) | 1)) {
                        case 32767:
                            var v73 float32 = 1e10
                            _ = v73
                        case 32768:
                            if false {
                                if false {
                                    var v74 float32 = 0.1
                                    v74 = v74
                                    if (true && false) {
                                        if (100 >= -1) {
                                            fmt.Println(fi69, v74)
                                            v74 -= ((float32(0.1) - float32(1.5)) + (3.14159 / (1.5 / 1.5)))
                                        }
                                        v2 -= (^(v2))
                                    } else {
                                        switch (fi71 % (((fi69 * (fi71 >> ((2) & 63)))) | 1)) {
                                        case -9223372036854775808:
                                            v74 = ((v74 + 1.5) - -2.25)
                                        case -1:
                                            var v75 uint16 = 2
                                            v76 := GenS4{F0: (("\n\t" + "") + "\n\t"), F1: "foo bar"}
                                            if true {
                                                _ = len(v14)
                                                v77 := GenS8{F0: ("foo bar" + (("\n\t" + "hello") + "hello"))}
                                                if (!(1.5 != float64(v74))) {
                                                    v2 -= 4294967295
                                                } else {
                                                    var v78 float64 = ((float64(v74) * (-2.25 / 1.5)) + -2.25)
                                                    fmt.Println(v1, fi69)
                                                    _ = v78
                                                }
                                                _ = v77
                                            }
                                            _ = v75
                                            _ = v76
                                        default:
                                            var v79 int64 = 9223372036854775807
                                            _ = v79
                                        }
                                    }
                                    _ = v74
                                }
                                _ = len(v14)
                            } else {
                                if (945 < 870) {
                                    v1 -= v1
                                }
                                v1 -= v1
                            }
                        case 1:
                            var v80 string = "a"
                            _ = v80
                        default:
                            v81 := GenS5{F0: "hello"}
                            _ = v81
                        }
                        v2 = 4294967295
                    }
                }
                _ = fi71
            }
        }
        _ = fi69
    }
    var v82 float64 = ((float64(0.1) - 3.14159) + (3.14159 * (((0.0 - float64((((-2.25 - 1.5) + 1e10) / 1.5))) * 0.0) + -2.25)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v82:", v82)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v46
    _ = v82
}
