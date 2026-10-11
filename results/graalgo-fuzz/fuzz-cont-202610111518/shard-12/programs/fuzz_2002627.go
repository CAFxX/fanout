package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int32; F2 int64 }
type GenS3 struct { F0 int8; F1 uint64; F2 bool }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int; F1 string; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 string; F2 int8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32 }
func genFn0(p0 uint16) int {
    switch -1 {
    case 0:
        p0 = 0
    case -1:
        v5 := []int{}
        _ = v5
    default:
        var v6 int = 9223372036854775807
        _ = v6
    }
    var v7 float32 = 1e10
    _ = p0
    _ = v7
    return 32768
}

func main() {
    var v1 int8 = 0
    v2 := []int{851}
    var v3 int64 = 0
    var v4 float32 = (0.1 - float32(3.14159))
    var v8 float64 = 1.5
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint32 = 2147483648
        _ = fi9
        _ = v10
    }
    if true {
        if len(v2) > 0 {
            v2[0] = -9223372036854775808
        }
    } else {
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 uint32 = 65535
                switch fi12 {
                case 128:
                    var v14 int = (fi12 + fi11)
                    _ = v14
                case 2147483647:
                    if true {
                        v1 = v1
                        v15 := GenS1{F0: 42}
                        switch ("hello" + ("\n\t" + "foo bar")) {
                        case "a":
                            fmt.Println(v8, v13)
                        case "":
                            v1 -= (-(v1))
                        }
                        v13 -= 637
                        _ = v15
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case 127:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    }
    if (32767 > 757) {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 float32 = float32(v8)
            _ = fi16
            _ = v17
        }
        v8 = (float64((v4 + (float32((float64(v4) + (v8 - (v8 * (3.14159 - v8))))) * v4))) * v8)
        v8 = v8
    } else {
        var v18 uint32 = 42
        for fi19 := 0; fi19 < 8; fi19++ {
            if true {
                var v20 string = "a"
                var v21 string = (v20 + v20)
                _ = v20
                _ = v21
            }
            if (!true) {
                switch ("foo bar" + "hello") {
                case "hello":
                    v8 -= (0.0 + (((v8 - v8) + v8) / ((0.0 * v8) + (v8 * float64(v4)))))
                }
                v3 = v3
            }
            if ((365 == 18446744073709551615) && (!(!false))) {
                if len(v2) > 0 {
                    v2[0] = (fi19 & -283)
                }
                v18 += (^(v18))
                if (!true) {
                    var v22 int8 = -128
                    var v23 float32 = (v4 - 0.0)
                    _ = v22
                    _ = v23
                } else {
                    var v24 float64 = v8
                    _ = v24
                }
            } else {
                for fi25 := 0; fi25 < 8; fi25++ {
                    v4 = 3.14159
                    _ = fi25
                }
                v4 = (v4 / float32(-2.25))
                v8 = 0.1
            }
            _ = fi19
        }
        _ = v18
    }
    v8 -= (v8 * v8)
    v4 = -2.25
    if len(v2) > 0 {
        v2[0] = 42
    }
    if len(v2) > 0 {
        v2[0] = 127
    }
    if false {
        var v26 int32 = -1
        var v27 int32 = (v26 & v26)
        v8 = v8
        _ = v26
        _ = v27
    } else {
        var v28 bool = false
        _ = v28
    }
    switch "" {
    case "foo bar":
        v4 -= float32(((v8 + 1.5) + (3.14159 - -2.25)))
    case "":
        fmt.Println(v3)
    case "\n\t":
        var v29 float64 = float64((0.1 - v4))
        _ = v29
    default:
        if true {
            for fi30 := 0; fi30 < 18; fi30++ {
                for fi31 := 0; fi31 < 6; fi31++ {
                    switch ("" + "") {
                    case "a":
                        v32 := GenS1{F0: 42}
                        _ = v32
                    case "\n\t":
                        var v33 int8 = v1
                        _ = v33
                    case "foo bar":
                        if true {
                            if false {
                                v34 := GenS2{F0: 255, F1: 0, F2: (v3 % ((v3) | 1))}
                                if len(v2) > 0 {
                                    v2[0] = (fi31 ^ (fi31 * fi31))
                                }
                                _ = v34
                            } else {
                                v3 = (v3 >> ((0) & 63))
                            }
                            switch 423 {
                            case 288:
                                switch (fi30 >> ((1) & 63)) {
                                case 9223372036854775807:
                                    var v35 string = ("" + "a")
                                    _ = v35
                                }
                                var v36 int16 = 1000
                                v37 := GenS2{F0: 0, F1: 474, F2: (v3 % ((42) | 1))}
                                _ = v36
                                _ = v37
                            case 127:
                                v1 = (v1 + (v1 % ((0) | 1)))
                            case 4294967295:
                                var v38 int64 = (v3 | (v3 << ((664) & 63)))
                                _ = v38
                            }
                        } else {
                            if len(v2) > 0 {
                                v2[0] = fi30
                            }
                            var v39 string = ("" + "a")
                            _ = v39
                        }
                        fmt.Println(v4, v1, v8)
                        v1 = 100
                    default:
                        v40 := GenS3{F0: v1, F1: 2, F2: false}
                        fmt.Println(fi30, fi30)
                        _ = v40
                    }
                    switch fi30 {
                    case 32768:
                        switch fi31 {
                        case 32768:
                            var v41 uint32 = 530
                            _ = v41
                        }
                    case 302:
                        v42 := GenS4{F0: (fi31 | 42), F1: -32768}
                        _ = v42
                    case 837:
                        if false {
                            switch genFn0(65535) {
                            case 612:
                                v43 := GenS5{F0: fi30, F1: (("" + ("foo bar" + "\n\t")) + ("a" + ((("" + (("a" + "hello") + ("" + ""))) + "hello") + ("a" + "hello")))), F2: ("" + "")}
                                _ = v43
                            case 254:
                                var v44 uint8 = 0
                                _ = v44
                            default:
                                if ("a" != ("\n\t" + "a")) {
                                    var v45 uint8 = 147
                                    if true {
                                        var v46 uint32 = 746
                                        _ = v46
                                    } else {
                                        var v47 uint64 = 0
                                        switch (("\n\t" + "foo bar") + "foo bar") {
                                        case "":
                                            v4 -= v4
                                        default:
                                            v48 := GenS4{F0: (fi30 ^ (fi30 * (fi30 * (fi31 * fi31)))), F1: 1}
                                            _ = v48
                                        }
                                        var v49 string = "foo bar"
                                        _ = v47
                                        _ = v49
                                    }
                                    v45 -= 197
                                    _ = v45
                                } else {
                                    if false {
                                        var v50 int8 = -28
                                        _ = v50
                                    }
                                }
                            }
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = (fi30 & (fi31 << ((2) & 63)))
                    }
                    _ = fi31
                }
                fmt.Println(v8, v8, v1)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi30
            }
            switch 42 {
            case 351:
                var v51 string = "\n\t"
                _ = v51
            case -2147483648:
                if true {
                    v52 := GenS6{F0: 354}
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 1; fi53++ {
                        for fi54 := 0; fi54 < 10; fi54++ {
                            switch ("\n\t" + "a") {
                            case "a":
                                var v55 int64 = (v3 / (((v3 & (v3 % ((-158) | 1)))) | 1))
                                _ = v55
                            case "":
                                var v56 int16 = 32767
                                _ = v56
                            case "hello":
                                v2 = append(v2, fi54)
                            }
                            _ = fi54
                        }
                        fmt.Println(v4, fi53, v1)
                        for fi57 := 0; fi57 < 7; fi57++ {
                            if (true || (!((v3 | v3) >= 1))) {
                                var v58 float32 = ((v4 * v4) / v4)
                                v59 := GenS7{F0: "a", F1: "a", F2: (v1 & (v1 >> ((42) & 63)))}
                                fmt.Println(fi57)
                                var v60 string = "a"
                                _ = v58
                                _ = v59
                                _ = v60
                            } else {
                                v4 = (v4 / (v4 + v4))
                                if true {
                                    v8 += (((float64(v4) * (((1.5 + (v8 + (v8 + float64(1e10)))) / (v8 + v8)) * (3.14159 - float64(v4)))) / (v8 * 1e10)) - (v8 + ((3.14159 - 1e10) / v8)))
                                } else {
                                    var v61 int = fi53
                                    v8 -= float64(v4)
                                    _ = v61
                                }
                                v62 := GenS1{F0: 65535}
                                _ = v62
                            }
                            v8 = v8
                            _ = fi57
                        }
                        _ = fi53
                    }
                    v1 -= 127
                    switch ("" + (("a" + "a") + "hello")) {
                    case "foo bar":
                        for fi63 := 0; fi63 < 1; fi63++ {
                            for fi64 := 0; fi64 < 18; fi64++ {
                                v3 += v3
                                v65 := GenS7{F0: "a", F1: ("\n\t" + ""), F2: v1}
                                _ = fi64
                                _ = v65
                            }
                            fmt.Println(v3, v3)
                            _ = fi63
                        }
                    case "hello":
                        if (200 > 2) {
                            v66 := GenS8{F0: 1}
                            _ = v66
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v3, v4, v1)
                        }
                    default:
                        v2 = append(v2, -142)
                    }
                }
            case -554:
                var v67 uint64 = 2
                _ = v67
            default:
                var v68 float32 = v4
                _ = v68
            }
            if (true || (false && false)) {
                var v69 uint32 = 2
                if true {
                    for fi70 := 0; fi70 < 7; fi70++ {
                        fmt.Println(v3)
                        switch fi70 {
                        case 32767:
                            v71 := GenS2{F0: 65535, F1: -1, F2: -319}
                            _ = v71
                        case -2147483648:
                            for fi72 := 0; fi72 < 12; fi72++ {
                                var v73 int8 = 1
                                var v74 uint = 54
                                _ = fi72
                                _ = v73
                                _ = v74
                            }
                        case 0:
                            var v75 int8 = (v1 >> ((uint(v69)) & 63))
                            _ = v75
                        }
                        v69 = v69
                        _ = fi70
                    }
                    switch "hello" {
                    case "foo bar":
                        var v76 uint = 2
                        _ = v76
                    case "a":
                        var v77 uint32 = (v69 << ((0) & 63))
                        _ = v77
                    case "\n\t":
                        for fi78 := 0; fi78 < 11; fi78++ {
                            switch ("hello" + "\n\t") {
                            case "":
                                switch fi78 {
                                case -114:
                                    var v79 uint8 = 42
                                    _ = v79
                                case -756:
                                    switch "hello" {
                                    case "a":
                                        if (v8 < ((-2.25 - ((0.1 * -2.25) + 1.5)) * v8)) {
                                            fmt.Println(v1, v1)
                                        }
                                    default:
                                        fmt.Println(v8)
                                    }
                                }
                            default:
                                v4 = (float32(v8) / 1.5)
                            }
                            _ = fi78
                        }
                    }
                    v80 := GenS4{F0: 42, F1: 1000}
                    _ = v80
                }
                switch 1 {
                case -2147483648:
                    if true {
                        if (false && (-412 <= -845)) {
                            v1 -= (^(v1))
                            for fi81 := 0; fi81 < 10; fi81++ {
                                v1 = (v1 | (v1 >> ((42) & 63)))
                                v69 = (v69 / (((v69 >> ((255) & 63))) | 1))
                                v82 := GenS9{F0: 2147483647}
                                _ = fi81
                                _ = v82
                            }
                        }
                    }
                default:
                    switch 32767 {
                    case 42:
                        v83 := GenS9{F0: (v69 >> ((0) & 63))}
                        _ = v83
                    default:
                        var v84 int64 = -1
                        _ = v84
                    }
                }
                _ = v69
            } else {
                v4 = v4
                var v85 int32 = 966
                _ = v85
            }
        } else {
            var v86 string = ("\n\t" + "a")
            fmt.Println(v4, v3, v4)
            v2 = append(v2, 901)
            _ = v86
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
