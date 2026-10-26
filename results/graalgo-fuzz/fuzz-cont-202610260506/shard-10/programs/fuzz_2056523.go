package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint32; F2 int8 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint64; F1 int; F2 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16 }
func genFn0(p0 int) int32 {
    switch "hello" {
    case "\n\t":
        for fi5 := 0; fi5 < 3; fi5++ {
            v6 := []int{(fi5 * (fi5 / ((p0) | 1))), (fi5 + (p0 * 209)), 0, p0}
            for fi7 := 0; fi7 < 15; fi7++ {
                v6 = append(v6, 4294967295)
                p0 = (fi7 * fi7)
                _ = fi7
            }
            _ = fi5
            _ = v6
        }
    case "":
        var v8 bool = true
        _ = v8
    default:
        v9 := GenS1{F0: -979, F1: -32768, F2: (!("" >= ""))}
        _ = v9
    }
    var v10 uint8 = 0
    v11 := []int{(p0 | (p0 << ((4294967295) & 63))), (p0 / (((p0 >> ((2) & 63))) | 1)), 127, -1}
    switch (p0 / ((-2147483648) | 1)) {
    case 4294967295:
        var v12 string = "a"
        _ = v12
    case 42:
        var v13 string = ("a" + (("\n\t" + ("a" + "\n\t")) + "a"))
        _ = v13
    case 0:
        for fi14 := 0; fi14 < 12; fi14++ {
            for fi15 := 0; fi15 < 1; fi15++ {
                switch 647 {
                case 2147483647:
                    var v16 int16 = 42
                    _ = v16
                case 1:
                    if true {
                        v17 := GenS1{F0: (p0 >> ((255) & 63)), F1: 0, F2: false}
                        if true {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            var v18 float64 = (-2.25 / 1.5)
                            _ = v18
                        } else {
                            if true {
                                v19 := GenS2{F0: -1}
                                var v20 int16 = 0
                                v11 = append(v11, fi14)
                                _ = v19
                                _ = v20
                            }
                        }
                        p0 -= fi14
                        _ = v17
                    }
                case -2147483648:
                    v21 := GenS2{F0: 125}
                    _ = v21
                }
                p0 = (p0 ^ (fi15 / ((fi15) | 1)))
                _ = fi15
            }
            if true {
                if (-314 == 697) {
                    v22 := GenS1{F0: 2147483647, F1: 0, F2: ((v10 / (((^(v10))) | 1)) < v10)}
                    v10 -= (v10 / (((v10 & v10)) | 1))
                    _ = v22
                } else {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    v10 -= v10
                }
            } else {
                fmt.Println(fi14, fi14, fi14)
            }
            var v23 float32 = (-2.25 - 0.1)
            _ = fi14
            _ = v23
        }
    default:
        switch p0 {
        case 665:
            var v24 int = p0
            _ = v24
        default:
            p0 += (p0 >> ((2) & 63))
        }
    }
    _ = p0
    _ = v10
    _ = v11
    return 1
}
func genFn1(p0 int8, p1 int16) int32 {
    var v25 float32 = (1.5 - 0.0)
    if (!true) {
        if true {
            v25 += ((3.14159 / 1e10) / v25)
            for fi26 := 0; fi26 < 10; fi26++ {
                switch "\n\t" {
                case "a":
                    v27 := []int{(fi26 & 4294967295), 1, (fi26 / (((^(fi26))) | 1))}
                    _ = v27
                case "hello":
                    var v28 uint8 = 165
                    _ = v28
                case "foo bar":
                    for fi29 := 0; fi29 < 14; fi29++ {
                        v30 := GenS3{F0: (1 > 0), F1: 42, F2: -54}
                        _ = fi29
                        _ = v30
                    }
                }
                if (false && (true || (!(-1 < 1)))) {
                    p0 -= -128
                    switch ("foo bar" + "\n\t") {
                    case "\n\t":
                        var v31 float64 = (float64(float32(float64((((v25 * v25) - float32(float64(v25))) * v25)))) * 1.5)
                        _ = v31
                    }
                }
                v32 := GenS1{F0: fi26, F1: -1, F2: (164 <= fi26)}
                _ = fi26
                _ = v32
            }
        }
        v25 = (v25 - float32((1.5 / ((((((0.1 * 0.0) - 1.5) + (float64((v25 / float32(1.5))) - ((1.5 + (0.1 / (-2.25 + 0.1))) * (1e10 * (1.5 / 0.1))))) * 0.0) - 1e10) * (1.5 - 3.14159)))))
        var v33 string = ""
        _ = v33
    } else {
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v34 int16 = (p1 + (p1 & p1))
            _ = v34
        case -1:
            if false {
                p1 -= (p1 + p1)
                var v35 uint16 = 150
                v36 := []int{-9223372036854775808}
                _ = v35
                _ = v36
            }
        }
        var v37 int = 2147483647
        _ = v37
    }
    _ = p0
    _ = p1
    _ = v25
    return 42
}

func main() {
    var v1 int32 = -686
    v2 := make(map[string]uint32)
    var v3 bool = false
    var v4 uint64 = 238
    switch -946 {
    case -794:
        switch -9223372036854775808 {
        case -9223372036854775808:
            fmt.Println(v1, v3)
        case 2147483647:
            fmt.Println(v3)
        }
    case -1:
        if ((v1 + v1) >= v1) {
            delete(v2, "")
            v38 := GenS4{F0: -1, F1: "\n\t"}
            _ = v38
        }
    }
    v39 := GenS2{F0: 60}
    if (886 > 42) {
        var v40 bool = v3
        var v41 float64 = 3.14159
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 13; fi42++ {
        if (!(false || (!(fi42 != fi42)))) {
            v43 := []int{(fi42 % ((430) | 1)), (fi42 | (fi42 & fi42)), (fi42 * 42)}
            switch (fi42 / (((fi42 ^ (fi42 >> ((255) & 63)))) | 1)) {
            case 2147483647:
                for fi44 := 0; fi44 < 10; fi44++ {
                    var v45 int32 = v1
                    var v46 float32 = 0.0
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            case -9223372036854775808:
                v47 := GenS5{F0: (v4 / ((v4) | 1)), F1: (fi42 >> ((18446744073709551615) & 63)), F2: 0}
                _ = v47
            case 66:
                var v48 int = fi42
                _ = v48
            default:
                v4 = (v4 / ((v4) | 1))
            }
            _ = v43
        } else {
            switch (fi42 ^ 9223372036854775807) {
            case -2147483648:
                var v49 uint16 = 0
                _ = v49
            }
            v1 += (v1 ^ 0)
            v50 := GenS5{F0: (v4 / (((v4 / ((0) | 1))) | 1)), F1: -283, F2: 42}
            _ = v50
        }
        var v51 uint8 = 42
        _ = fi42
        _ = v51
    }
    v52 := GenS6{F0: -223}
    switch "" {
    case "hello":
        v1 = (-(v1))
    case "a":
        v1 = (v1 - v1)
    }
    v53 := GenS3{F0: (!(false || ("foo bar" == "a"))), F1: 2147483648, F2: 42}
    for fi54 := 0; fi54 < 20; fi54++ {
        v1 = v1
        var v55 uint8 = 129
        delete(v2, ((("hello" + ("\n\t" + ("a" + "foo bar"))) + "a") + ("a" + "hello")))
        _ = fi54
        _ = v55
    }
    v56 := []int{32767, -384}
    var v57 bool = (0.1 != (3.14159 * float64(1e10)))
    if (42 >= 32767) {
        var v58 int16 = 1000
        switch "foo bar" {
        case "":
            for fi59 := 0; fi59 < 19; fi59++ {
                if false {
                    for fi60 := 0; fi60 < 17; fi60++ {
                        v1 += int32(fi60)
                        _ = fi60
                    }
                }
                if v3 {
                    switch fi59 {
                    case -364:
                        for fi61 := 0; fi61 < 4; fi61++ {
                            var v62 bool = (-1 <= 9223372036854775807)
                            _ = fi61
                            _ = v62
                        }
                    case 127:
                        if len(v56) > 0 {
                            v56[0] = fi59
                        }
                    }
                    fmt.Println(v58, fi59, v1)
                    for fi63 := 0; fi63 < 17; fi63++ {
                        if len(v56) >= 2 {
                            v56 = v56[1:len(v56)]
                        }
                        _ = fi63
                    }
                } else {
                    v1 += 42
                }
                v64 := GenS7{F0: 0}
                _ = fi59
                _ = v64
            }
        case "a":
            v56 = append(v56, 363)
        case "hello":
            fmt.Println(v4, v4)
        default:
            for fi65 := 0; fi65 < 6; fi65++ {
                if v3 {
                    var v66 string = "foo bar"
                    if true {
                        v2["foo bar"] = 674
                        var v67 uint32 = 4294967295
                        v68 := GenS8{F0: v58}
                        _ = v67
                        _ = v68
                    } else {
                        var v69 int64 = 1
                        _ = v69
                    }
                    var v70 uint32 = 2147483647
                    _ = v66
                    _ = v70
                } else {
                    var v71 int64 = 42
                    var v72 uint8 = 127
                    _ = v71
                    _ = v72
                }
                _ = fi65
            }
        }
        _ = v58
    }
    switch "a" {
    case "\n\t":
        var v73 uint16 = 1
        _ = v73
    case "a":
        var v74 float32 = 1e10
        _ = v74
    default:
        if true {
            var v75 float32 = 1e10
            var v76 string = "foo bar"
            var v77 int16 = 1000
            _ = v75
            _ = v76
            _ = v77
        } else {
            var v78 int8 = 100
            _ = v78
        }
    }
    if len(v56) >= 2 {
        v56 = v56[1:len(v56)]
    }
    var v79 int32 = (v1 >> ((65535) & 63))
    v4 -= (v4 + (v4 % ((v4) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v57:", v57)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v52
    _ = v53
    _ = v56
    _ = v57
    _ = v79
}
