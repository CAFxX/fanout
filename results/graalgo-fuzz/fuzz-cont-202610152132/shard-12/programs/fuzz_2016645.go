package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int32; F1 int64; F2 int8 }
type GenS3 struct { F0 uint; F1 int64; F2 uint }
type GenS4 struct { F0 uint32; F1 int; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint32; F2 string }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int32; F1 int64; F2 uint }
type GenS8 struct { F0 uint; F1 int32 }
func genFn0() bool {
    v7 := GenS1{F0: false, F1: 611, F2: 32767}
    var v8 int8 = -1
    v9 := GenS2{F0: 2147483647, F1: 6, F2: (-(v8))}
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]uint32)
    var v3 int8 = 0
    v4 := []string{}
    var v5 int64 = v1
    var v6 uint32 = 2147483647
    if (true || false) {
        _ = len(v4)
        fmt.Println(v1)
        for fi10 := 0; fi10 < 18; fi10++ {
            if ((v5 % ((-313) | 1)) != (v1 - v1)) {
                fmt.Println(v1, fi10)
                switch "a" {
                case "\n\t":
                    switch (fi10 ^ -799) {
                    case -9223372036854775808:
                        v11 := GenS3{F0: 0, F1: v1, F2: 573}
                        _ = v11
                    case 128:
                        v12 := GenS1{F0: true, F1: 4294967295, F2: 529}
                        _ = v12
                    }
                case "":
                    if true {
                        v13 := GenS2{F0: -482, F1: v1, F2: (-(v3))}
                        var v14 string = ""
                        _ = v13
                        _ = v14
                    }
                default:
                    v15 := GenS1{F0: ((true && false) || false), F1: 65535, F2: 747}
                    _ = v15
                }
            }
            if false {
                v4 = append(v4, "hello")
            } else {
                if (18446744073709551615 != 63) {
                    var v16 float32 = 1.5
                    var v17 uint16 = 32767
                    _ = v16
                    _ = v17
                }
            }
            for fi18 := 0; fi18 < 17; fi18++ {
                v19 := GenS2{F0: 1, F1: v5, F2: 1}
                var v20 bool = false
                v2[(-(fi10))] = 65535
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi10
        }
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(fi21, v6, v6)
        _ = fi21
    }
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        for fi22 := 0; fi22 < 12; fi22++ {
            switch 1 {
            case -1:
                var v23 string = ("hello" + "")
                _ = v23
            case 127:
                var v24 uint64 = 828
                _ = v24
            default:
                switch (-(fi22)) {
                case 4294967295:
                    for fi25 := 0; fi25 < 1; fi25++ {
                        switch fi25 {
                        case -213:
                            var v26 uint = 2
                            _ = v26
                        case 2147483647:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            v27 := GenS2{F0: -1, F1: v1, F2: v3}
                            _ = v27
                        }
                        delete(v2, 112)
                        v4 = append(v4, "hello")
                        _ = fi25
                    }
                }
            }
            var v28 string = ""
            v29 := GenS1{F0: (((-2.25 * (float64((1e10 * 3.14159)) / float64(-2.25))) * (float64((1e10 / -2.25)) / (1e10 + 1.5))) >= 3.14159), F1: 878, F2: 724}
            _ = fi22
            _ = v28
            _ = v29
        }
    }
    if (!true) {
        v3 = (v3 >> ((1) & 63))
        var v30 uint = 42
        _ = v30
    }
    var v31 float32 = float32(float64(float32(0.1)))
    switch "" {
    case "a":
        if false {
            for fi32 := 0; fi32 < 17; fi32++ {
                for fi33 := 0; fi33 < 18; fi33++ {
                    switch "a" {
                    case "a":
                        v31 += v31
                    case "\n\t":
                        v3 += 127
                    default:
                        var v34 uint8 = 0
                        _ = v34
                    }
                    if (1 >= 1) {
                        if false {
                            v35 := GenS3{F0: 336, F1: v5, F2: 340}
                            _ = v35
                        }
                        switch fi32 {
                        case -495:
                            var v36 float64 = 0.0
                            _ = v36
                        case 32768:
                            v4 = append(v4, ("hello" + ((("hello" + ("foo bar" + "\n\t")) + "hello") + "\n\t")))
                        }
                        fmt.Println(fi33)
                        var v37 uint64 = 18446744073709551615
                        fmt.Println(v5, v3, v1)
                        _ = len(v2)
                        v38 := GenS1{F0: false, F1: (v6 | (v6 & (^(v6)))), F2: 2}
                        v39 := GenS2{F0: -2147483648, F1: v5, F2: -128}
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    if (42 < 190) {
                        fmt.Println(v5, v3, v31)
                        v40 := GenS4{F0: v6, F1: fi32, F2: 65535}
                        v41 := GenS1{F0: (544 != 4294967295), F1: v6, F2: 842}
                        v42 := GenS4{F0: v6, F1: fi33, F2: 644}
                        v43 := GenS5{F0: 1, F1: (v6 + (v6 ^ 65535)), F2: ("\n\t" + ((("foo bar" + ("foo bar" + (("\n\t" + ((("hello" + "foo bar") + "a") + ("a" + ("hello" + "foo bar")))) + ("\n\t" + ("a" + "a"))))) + "\n\t") + "hello"))}
                        switch "a" {
                        case "hello":
                            if false {
                                var v44 string = "a"
                                _ = v44
                            }
                        case "a":
                            if true {
                                v45 := GenS1{F0: (((((-2.25 - (float64(float32((1.5 - float64(float32((((0.1 + (0.1 / 1.0)) - 1.5) - 1e10)))))) - 1.5)) * -2.25) + 0.0) <= (0.1 + 1.5)) || ((!false) && (v6 >= 812))), F1: 1, F2: 255}
                                var v46 int = (fi32 | (fi32 << ((714) & 63)))
                                _ = v45
                                _ = v46
                            }
                        case "":
                            v3 = v3
                        }
                        _ = v40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    _ = fi33
                }
                var v47 uint64 = 0
                if ((true || (!((v47 + v47) == uint64(v6)))) && false) {
                    var v48 uint8 = 104
                    fmt.Println(v31, v31)
                    _ = v48
                } else {
                    v6 = v6
                    fmt.Println(v1)
                }
                _ = fi32
                _ = v47
            }
            if false {
                if (230 > 1) {
                    var v49 int16 = 32767
                    if (2 > 32767) {
                        var v50 int8 = -99
                        v6 -= (v6 + (v6 >> ((42) & 63)))
                        _ = v50
                    } else {
                        fmt.Println(v3)
                        var v51 int = -1
                        var v52 uint16 = 322
                        _ = v51
                        _ = v52
                    }
                    _ = v49
                } else {
                    if ("hello" < "foo bar") {
                        v3 += v3
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            } else {
                for fi53 := 0; fi53 < 13; fi53++ {
                    switch "" {
                    case "hello":
                        if (553 >= 255) {
                            v54 := GenS6{F0: 65535, F1: "foo bar"}
                            switch ("\n\t" + ("\n\t" + "\n\t")) {
                            case "":
                                var v55 uint = 18446744073709551615
                                _ = v55
                            default:
                                fmt.Println(fi53)
                            }
                            fmt.Println(v6, v3, v1)
                            _ = v54
                        } else {
                            if (v6 <= (v6 & (v6 * (v6 - (v6 << ((255) & 63)))))) {
                                v3 = (v3 + (v3 >> ((2) & 63)))
                            }
                        }
                    case "foo bar":
                        for fi56 := 0; fi56 < 7; fi56++ {
                            v57 := GenS7{F0: -138, F1: v1, F2: 0}
                            var v58 int64 = v5
                            v3 = v3
                            _ = fi56
                            _ = v57
                            _ = v58
                        }
                    case "\n\t":
                        v59 := GenS1{F0: ((v3 % ((0) | 1)) == v3), F1: (v6 / (((^(v6))) | 1)), F2: 1}
                        _ = v59
                    }
                    var v60 float32 = (1e10 / v31)
                    v31 -= float32(1e10)
                    _ = fi53
                    _ = v60
                }
            }
        } else {
            v61 := GenS8{F0: 255, F1: 2147483647}
            _ = v61
        }
    case "\n\t":
        var v62 int8 = (v3 % ((0) | 1))
        _ = v62
    }
    v5 -= v5
    var v63 uint16 = 32767
    v64 := GenS7{F0: -153, F1: (v5 % ((9223372036854775807) | 1)), F2: 0}
    switch ("foo bar" + "a") {
    case "foo bar":
        if (2147483647 < 1) {
            if false {
                var v65 uint16 = 927
                v66 := GenS6{F0: (v65 & v65), F1: "hello"}
                _ = v65
                _ = v66
            }
            var v67 bool = false
            _ = v67
        } else {
            if (127 >= 127) {
                if (!true) {
                    v68 := GenS3{F0: 18446744073709551615, F1: -502, F2: uint(v63)}
                    v2[0] = v6
                    _ = v68
                }
            }
            v31 -= v31
            for fi69 := 0; fi69 < 14; fi69++ {
                if (v3 <= v3) {
                    v1 -= (v5 ^ v1)
                    if true {
                        v6 = (^(v6))
                        v31 -= 1e10
                    } else {
                        switch (-(fi69)) {
                        case 879:
                            var v70 uint = 42
                            _ = v70
                        case 1:
                            switch ("a" + "a") {
                            case "":
                                fmt.Println(fi69, v5)
                            case "\n\t":
                                var v71 uint = 1
                                _ = v71
                            case "a":
                                var v72 bool = (537 <= uint64(v6))
                                _ = v72
                            }
                        case -9223372036854775808:
                            v3 -= (v3 ^ (^(v3)))
                        }
                        for fi73 := 0; fi73 < 16; fi73++ {
                            var v74 uint = 2
                            _ = fi73
                            _ = v74
                        }
                        var v75 int8 = -128
                        _ = v75
                    }
                    v5 = v5
                } else {
                    v76 := v2[(fi69 / ((fi69) | 1))]
                    _ = v76
                }
                _ = fi69
            }
        }
    case "hello":
        var v77 uint16 = (^(v63))
        _ = v77
    case "\n\t":
        var v78 int64 = (v5 % (((^(v1))) | 1))
        _ = v78
    default:
        for fi79 := 0; fi79 < 12; fi79++ {
            if true {
                if false {
                    var v80 uint8 = 1
                    v31 = v31
                    for fi81 := 0; fi81 < 12; fi81++ {
                        var v82 uint32 = v6
                        var v83 int16 = 1
                        var v84 string = (("\n\t" + "\n\t") + "\n\t")
                        var v85 int = (fi79 ^ (-(fi79)))
                        _ = fi81
                        _ = v82
                        _ = v83
                        _ = v84
                        _ = v85
                    }
                    _ = v80
                }
                var v86 float32 = (0.0 * v31)
                _ = v86
            }
            for fi87 := 0; fi87 < 18; fi87++ {
                v3 = (v3 + v3)
                _ = fi87
            }
            _ = fi79
        }
    }
    fmt.Println(v6, v5)
    v88 := GenS4{F0: v6, F1: 32768, F2: v63}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v88
}
