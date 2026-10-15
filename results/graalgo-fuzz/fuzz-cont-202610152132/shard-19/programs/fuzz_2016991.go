package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int64 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS6 struct { F0 int64; F1 bool; F2 uint16 }
func genFn0(p0 uint, p1 bool) int {
    fmt.Println(p1, p1)
    p1 = (p0 >= (p0 | p0))
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 uint32) int64 {
    var v5 uint = 255
    if true {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: 9223372036854775807, F1: p0, F2: -9223372036854775808}
            var v8 uint32 = (p0 << (((v5 - 255)) & 63))
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = v5
    return -9223372036854775808
}

func main() {
    v1 := []string{}
    var v2 uint16 = 638
    var v3 int = -124
    var v4 uint = 1
    v3 += (v3 ^ 127)
    for fi9 := 0; fi9 < 10; fi9++ {
        v4 = v4
        _ = fi9
    }
    v2 = 2
    var v10 uint64 = 18446744073709551615
    fmt.Println(v2, v3, v10)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v10 += 557
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    v2 = 0
    var v11 int32 = 64
    fmt.Println(v2, v2)
    switch ("\n\t" + ("foo bar" + "a")) {
    case "":
        var v12 uint = 4294967295
        _ = v12
    case "foo bar":
        fmt.Println(v3, v4)
    case "a":
        fmt.Println(v10)
    default:
        if (false || (!(((787 > 835) && false) && ((!false) && false)))) {
            if false {
                if (v2 < 65535) {
                    if (!(false || (((1e10 * (float32(-2.25) + 3.14159)) - 1.5) >= float32(3.14159)))) {
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                        v4 += 42
                        var v13 float64 = 0.0
                        _ = v13
                    }
                    v10 = v10
                    var v14 int16 = -32768
                    _ = v14
                } else {
                    _ = len(v1)
                }
                for fi15 := 0; fi15 < 13; fi15++ {
                    fmt.Println(v11, v3, v3)
                    _ = fi15
                }
                if (!false) {
                    v16 := GenS2{F0: 100, F1: 0}
                    v17 := GenS2{F0: 0, F1: 42}
                    _ = v16
                    _ = v17
                } else {
                    v2 = (v2 & 255)
                    v18 := GenS2{F0: 127, F1: 0}
                    if false {
                        v4 = uint(v2)
                        _ = len(v1)
                        var v19 uint8 = 137
                        _ = v19
                    }
                    _ = v18
                }
            }
            v20 := GenS1{F0: -9223372036854775808, F1: 434, F2: -245}
            for fi21 := 0; fi21 < 2; fi21++ {
                switch ("" + "foo bar") {
                case "":
                    v11 = -529
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi21
            }
            _ = v20
        } else {
            switch (v3 - 0) {
            case -280:
                v10 = 630
            case 934:
                for fi22 := 0; fi22 < 2; fi22++ {
                    if (!true) {
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        if (!(!false)) {
                            fmt.Println(v10, v10, fi22)
                            var v23 uint32 = 242
                            v2 = v2
                            _ = v23
                        } else {
                            for fi24 := 0; fi24 < 4; fi24++ {
                                var v25 int8 = 1
                                _ = fi24
                                _ = v25
                            }
                            v11 -= 2147483647
                        }
                        for fi26 := 0; fi26 < 6; fi26++ {
                            fmt.Println(v11, v4, v11)
                            v2 += v2
                            var v27 uint8 = 0
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        for fi28 := 0; fi28 < 18; fi28++ {
                            switch "foo bar" {
                            case "\n\t":
                                v10 -= (v10 - 4294967295)
                                if (true || false) {
                                    if false {
                                        var v29 int32 = 42
                                        v2 = v2
                                        var v30 uint8 = 1
                                        v31 := GenS3{F0: 1, F1: (v29 & v11)}
                                        var v32 bool = false
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    } else {
                                        v33 := GenS2{F0: -128, F1: 9223372036854775807}
                                        v34 := GenS2{F0: 1, F1: genFn1(65535)}
                                        _ = v33
                                        _ = v34
                                    }
                                    v4 += (^(v4))
                                }
                                if false {
                                    if len(v1) > 0 {
                                        v1[0] = "\n\t"
                                    }
                                    switch -658 {
                                    case 127:
                                        if (!(4294967295 < 1)) {
                                            var v35 float32 = 0.0
                                            if ((!false) && ((v2 - (v2 ^ (^(v2)))) > (v2 % ((v2) | 1)))) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v36 int16 = 42
                                                _ = v36
                                            }
                                            if true {
                                                if false {
                                                    _ = len(v1)
                                                    v3 -= (fi22 / (((fi22 + (fi22 << (((^(v4))) & 63)))) | 1))
                                                    var v37 uint32 = uint32(v4)
                                                    _ = v37
                                                }
                                                v10 = 4294967295
                                                if ((false && (2147483647 > 2)) || true) {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    if false {
                                                        v38 := GenS3{F0: 912, F1: -1}
                                                        var v39 uint32 = 936
                                                        v1 = append(v1, ("\n\t" + (((("" + ("\n\t" + "hello")) + "") + ("\n\t" + "")) + "a")))
                                                        _ = v38
                                                        _ = v39
                                                    } else {
                                                        v1 = append(v1, ("" + ("foo bar" + "\n\t")))
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            if true {
                                                                var v40 uint = v4
                                                                v41 := GenS1{F0: genFn1(968), F1: 1, F2: 0}
                                                                _ = v40
                                                                _ = v41
                                                            }
                                                        case "a":
                                                            v11 += (v11 >> ((v4) & 63))
                                                        case "":
                                                            var v42 int16 = -32768
                                                            _ = v42
                                                        }
                                                        fmt.Println(v11, fi22, v35)
                                                    }
                                                    switch 1 {
                                                    case 9223372036854775807:
                                                        if false {
                                                            if false {
                                                                v43 := GenS4{F0: "hello"}
                                                                _ = v43
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v35
                                        } else {
                                            fmt.Println(v2, fi22)
                                            var v44 bool = (!true)
                                            var v45 int8 = -128
                                            v1 = append(v1, "foo bar")
                                            _ = v44
                                            _ = v45
                                        }
                                    case 764:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case -2147483648:
                                        switch "foo bar" {
                                        case "":
                                            if true {
                                                switch ("a" + ("foo bar" + "")) {
                                                case "hello":
                                                    var v46 string = "a"
                                                    _ = v46
                                                }
                                                if true {
                                                    v3 = 955
                                                    _ = len(v1)
                                                    if (4294967295 <= (v10 % ((575) | 1))) {
                                                        v11 += 688
                                                    }
                                                } else {
                                                    v1 = append(v1, ("foo bar" + "a"))
                                                    var v47 uint = 2
                                                    v4 = uint(v10)
                                                    _ = v47
                                                }
                                                _ = len(v1)
                                                v2 = v2
                                            }
                                        case "hello":
                                            var v48 uint16 = 837
                                            _ = v48
                                        case "foo bar":
                                            fmt.Println(fi22, fi28, v2)
                                        default:
                                            switch fi28 {
                                            case 613:
                                                var v49 int32 = (v11 + 373)
                                                _ = v49
                                            case 1:
                                                v3 = -2147483648
                                            case 104:
                                                _ = len(v1)
                                            default:
                                                var v50 float32 = (3.14159 * -2.25)
                                                _ = v50
                                            }
                                            v10 = v10
                                        }
                                    }
                                    v10 -= (v10 ^ 712)
                                    if (!(255 != 2)) {
                                        if ((((1e10 - 0.0) + 0.1) > 3.14159) || ((1 == 178) && (!false))) {
                                            if ("foo bar" != "a") {
                                                var v51 uint16 = v2
                                                var v52 float32 = -2.25
                                                fmt.Println(v3, fi22, v51)
                                                _ = v51
                                                _ = v52
                                            } else {
                                                fmt.Println(v3, fi28, v3)
                                                v53 := GenS5{F0: 93, F1: uint64(v4), F2: v2}
                                                if (1 < (v11 / ((int32(v10)) | 1))) {
                                                    if ((v2 / ((v2) | 1)) >= (v2 << ((v4) & 63))) {
                                                        if len(v1) > 0 {
                                                            v1[0] = "\n\t"
                                                        }
                                                        var v54 int16 = 0
                                                        _ = v54
                                                    }
                                                    v1 = append(v1, "foo bar")
                                                    v4 = uint(v10)
                                                }
                                                _ = v53
                                            }
                                            if false {
                                                v55 := GenS4{F0: "a"}
                                                if false {
                                                    var v56 int = fi28
                                                    var v57 bool = ((!true) && false)
                                                    switch ("" + "a") {
                                                    case "foo bar":
                                                        _ = len(v1)
                                                    case "":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v58 int16 = -923
                                                            _ = v58
                                                        }
                                                    case "a":
                                                        v59 := GenS6{F0: 353, F1: true, F2: v2}
                                                        _ = v59
                                                    }
                                                    var v60 int8 = 0
                                                    _ = v56
                                                    _ = v57
                                                    _ = v60
                                                } else {
                                                    var v61 int = genFn0(v4, false)
                                                    var v62 int = v3
                                                    fmt.Println(fi28)
                                                    _ = v61
                                                    _ = v62
                                                }
                                                var v63 uint64 = (v10 % ((4294967295) | 1))
                                                _ = v55
                                                _ = v63
                                            } else {
                                                var v64 uint8 = 255
                                                _ = v64
                                            }
                                            if (!false) {
                                                v65 := GenS4{F0: "\n\t"}
                                                var v66 int8 = 100
                                                if len(v1) > 0 {
                                                    v1[0] = ("foo bar" + "\n\t")
                                                }
                                                _ = v65
                                                _ = v66
                                            } else {
                                                v67 := GenS5{F0: -1, F1: v10, F2: v2}
                                                switch "hello" {
                                                case "\n\t":
                                                    var v68 float32 = (0.0 - float32(float64(-2.25)))
                                                    _ = v68
                                                case "hello":
                                                    switch 4294967295 {
                                                    case 125:
                                                        if true {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                    case 42:
                                                        switch "" {
                                                        case "\n\t":
                                                            v11 -= v11
                                                        default:
                                                            switch fi22 {
                                                            case 216:
                                                                switch ("" + "a") {
                                                                case "foo bar":
                                                                    fmt.Println(fi22)
                                                                default:
                                                                    var v69 uint16 = 255
                                                                    _ = v69
                                                                }
                                                            case 32768:
                                                                var v70 uint = (v4 & (v4 % (((v4 * 541)) | 1)))
                                                                _ = v70
                                                            case 1:
                                                                v71 := GenS6{F0: genFn1(42), F1: true, F2: 32767}
                                                                _ = v71
                                                            }
                                                        }
                                                    }
                                                case "":
                                                    v4 -= uint(v2)
                                                default:
                                                    var v72 int8 = -62
                                                    _ = v72
                                                }
                                                _ = v67
                                            }
                                        } else {
                                            _ = len(v1)
                                        }
                                        v10 -= 2
                                        fmt.Println(fi28, v4)
                                    } else {
                                        var v73 uint16 = (^(v2))
                                        v1 = append(v1, "a")
                                        if len(v1) > 0 {
                                            v1[0] = "foo bar"
                                        }
                                        _ = v73
                                    }
                                }
                            case "a":
                                v74 := GenS4{F0: "hello"}
                                _ = v74
                            }
                            _ = fi28
                        }
                    }
                    _ = fi22
                }
            case 127:
                var v75 float32 = (3.14159 + (((0.1 * 1e10) * float32((((float64(3.14159) / -2.25) - (-2.25 * float64(1e10))) / 1e10))) - (3.14159 - 1e10)))
                _ = v75
            default:
                v4 = v4
            }
            if false {
                if len(v1) > 0 {
                    v1[0] = ""
                }
                for fi76 := 0; fi76 < 19; fi76++ {
                    fmt.Println(v11, v10)
                    _ = fi76
                }
                fmt.Println(v2, v4)
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
