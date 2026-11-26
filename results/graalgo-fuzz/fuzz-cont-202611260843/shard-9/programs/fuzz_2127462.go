package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool; F1 int; F2 uint64 }
func genFn0(p0 int) uint32 {
    var v5 int32 = 0
    switch (("\n\t" + "") + ("a" + "\n\t")) {
    case "":
        var v6 string = "hello"
        _ = v6
    case "hello":
        v7 := []int{(p0 ^ (-(p0))), p0, (p0 & (p0 | p0)), 42}
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint64 = 42
        if true {
            if true {
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 uint16 = 42
                    v12 := GenS1{F0: (-75 < -9223372036854775808), F1: true}
                    switch (fi8 + fi8) {
                    case 32768:
                        var v13 string = ""
                        _ = v13
                    default:
                        fmt.Println(fi10, v11, fi8)
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v9 += (v9 * (v9 | v9))
            }
        } else {
            v9 = v9
        }
        fmt.Println(v5, p0)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    return 224
}
func genFn1(p0 uint32, p1 int32, p2 int64) uint {
    v14 := []int{845}
    var v15 uint = 408
    if (true || ("foo bar" < "a")) {
        for fi16 := 0; fi16 < 3; fi16++ {
            for fi17 := 0; fi17 < 12; fi17++ {
                v14 = append(v14, fi16)
                v18 := GenS2{F0: ("\n\t" + "hello"), F1: (!("foo bar" >= "\n\t"))}
                var v19 int64 = p2
                _ = fi17
                _ = v18
                _ = v19
            }
            if true {
                switch (^(fi16)) {
                case 127:
                    p1 = (p1 >> ((uint(p0)) & 63))
                case -175:
                    p2 = p2
                }
                var v20 uint64 = 0
                switch "a" {
                case "a":
                    if len(v14) > 0 {
                        v14[0] = fi16
                    }
                case "hello":
                    v21 := GenS2{F0: "\n\t", F1: (false && ((p2 >> ((2) & 63)) != 0))}
                    _ = v21
                case "\n\t":
                    var v22 int16 = 0
                    _ = v22
                default:
                    var v23 float32 = 1.5
                    _ = v23
                }
                _ = v20
            }
            var v24 float32 = 1e10
            _ = fi16
            _ = v24
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 4294967295
}

func main() {
    var v1 uint16 = 2
    var v2 int8 = 0
    v3 := make(map[int]bool)
    var v4 uint16 = v1
    v2 = v2
    if (968 == 42) {
        var v25 uint64 = 18446744073709551615
        switch "hello" {
        case "\n\t":
            fmt.Println(v4)
        case "hello":
            v25 = (v25 / (((v25 << ((1) & 63))) | 1))
        case "a":
            v26 := []int{-9223372036854775808}
            _ = v26
        default:
            if (0 >= genFn0(-805)) {
                var v27 float32 = float32(float64(1e10))
                v4 += (v4 + 255)
                _ = v27
            } else {
                v28 := []int{1, -9223372036854775808, 1}
                _ = v28
            }
        }
        _ = v25
    } else {
        v29 := GenS2{F0: ("foo bar" + ""), F1: true}
        switch -529 {
        case 830:
            var v30 float32 = -2.25
            _ = v30
        default:
            v1 = (v4 >> ((18446744073709551615) & 63))
        }
        _ = v29
    }
    var v31 int32 = 1
    fmt.Println(v2, v1)
    for fi32 := 0; fi32 < 11; fi32++ {
        if (1000 <= -32768) {
            if true {
                fmt.Println(v2, v4, v2)
                v33 := GenS1{F0: (false && (fi32 >= (fi32 & 42))), F1: (!true)}
                v3[(-(fi32))] = false
                _ = v33
            } else {
                var v34 int64 = -816
                v31 += -800
                var v35 string = "\n\t"
                _ = v34
                _ = v35
            }
            if (!false) {
                v3[(fi32 | fi32)] = false
                v2 -= -3
                v36 := []int{fi32, fi32, (^(fi32)), fi32}
                _ = v36
            }
        } else {
            var v37 float64 = (0.1 * (((0.0 / 1.0) - ((1e10 / 0.1) / 0.1)) - 3.14159))
            fmt.Println(v4, v31, v4)
            v1 = (v4 - 2)
            _ = v37
        }
        var v38 float64 = float64(0.1)
        if true {
            var v39 uint64 = 360
            _ = v39
        }
        _ = fi32
        _ = v38
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 uint8 = 127
        _ = fi40
        _ = v41
    }
    switch -1 {
    case 994:
        for fi42 := 0; fi42 < 2; fi42++ {
            v43 := []int{fi42, fi42, -1}
            if false {
                var v44 uint64 = 2
                if len(v43) > 0 {
                    v43[0] = fi42
                }
                v2 = -45
                _ = v44
            }
            var v45 string = ("foo bar" + "foo bar")
            _ = fi42
            _ = v43
            _ = v45
        }
    default:
        if false {
            fmt.Println(v1, v2)
            var v46 int64 = -34
            _ = v46
        }
    }
    v1 = v1
    var v47 int = -879
    var v48 uint = genFn1(2147483647, v31, -9223372036854775808)
    for fi49 := 0; fi49 < 8; fi49++ {
        var v50 float64 = (-2.25 - 3.14159)
        _ = fi49
        _ = v50
    }
    if (1 > 65535) {
        for fi51 := 0; fi51 < 6; fi51++ {
            var v52 int32 = (v31 * int32(v48))
            switch (v47 << ((v48) & 63)) {
            case -9223372036854775808:
                v31 -= (^(v31))
            case 2147483647:
                var v53 int = 127
                _ = v53
            }
            _ = fi51
            _ = v52
        }
        if true {
            if false {
                if false {
                    v47 = 202
                }
                switch v47 {
                case 1:
                    for fi54 := 0; fi54 < 6; fi54++ {
                        switch (fi54 >> ((v48) & 63)) {
                        case -233:
                            switch (fi54 * 1) {
                            case 492:
                                var v55 string = ("" + ("" + "hello"))
                                _ = v55
                            default:
                                v4 = 42
                            }
                        }
                        var v56 uint32 = 2147483648
                        v57 := GenS2{F0: ("a" + "a"), F1: (true || (!(v56 != (v56 + v56))))}
                        _ = fi54
                        _ = v56
                        _ = v57
                    }
                case 128:
                    switch ("\n\t" + ((("foo bar" + ("hello" + "\n\t")) + ("foo bar" + "hello")) + "")) {
                    case "":
                        var v58 uint = 902
                        _ = v58
                    case "hello":
                        v4 = (v1 - (^(v4)))
                    default:
                        fmt.Println(v47)
                    }
                default:
                    fmt.Println(v48, v1)
                }
                if false {
                    for fi59 := 0; fi59 < 10; fi59++ {
                        var v60 int64 = 42
                        v4 += 223
                        for fi61 := 0; fi61 < 6; fi61++ {
                            v62 := GenS1{F0: (!true), F1: (v31 < (v31 / ((v31) | 1)))}
                            _ = fi61
                            _ = v62
                        }
                        _ = fi59
                        _ = v60
                    }
                    if true {
                        if (1 >= 1) {
                            var v63 bool = (!((!(783 <= 4294967295)) && true))
                            _ = v63
                        } else {
                            if (-9223372036854775808 < 0) {
                                var v64 int = (v47 - 32768)
                                v65 := []int{127, 9223372036854775807}
                                switch (-(v47)) {
                                case 1:
                                    fmt.Println(v31)
                                case 2147483647:
                                    var v66 float64 = 1e10
                                    _ = v66
                                }
                                _ = v64
                                _ = v65
                            } else {
                                for fi67 := 0; fi67 < 15; fi67++ {
                                    for fi68 := 0; fi68 < 9; fi68++ {
                                        var v69 int16 = 1
                                        _ = fi68
                                        _ = v69
                                    }
                                    _ = fi67
                                }
                                var v70 int32 = (-(v31))
                                v31 = -612
                                _ = v70
                            }
                            switch (v47 ^ v47) {
                            case -1:
                                v71 := GenS1{F0: (false || false), F1: false}
                                _ = v71
                            }
                        }
                    }
                }
            }
        }
        fmt.Println(v4, v47)
    } else {
        v47 = 128
        for fi72 := 0; fi72 < 17; fi72++ {
            var v73 uint32 = 2147483647
            _ = fi72
            _ = v73
        }
        switch v47 {
        case 2147483647:
            if true {
                if false {
                    for fi74 := 0; fi74 < 8; fi74++ {
                        _ = len(v3)
                        var v75 int64 = 9223372036854775807
                        _ = fi74
                        _ = v75
                    }
                    v48 += genFn1(4294967295, (-(v31)), 1)
                    if true {
                        fmt.Println(v47)
                        for fi76 := 0; fi76 < 13; fi76++ {
                            for fi77 := 0; fi77 < 10; fi77++ {
                                v78 := GenS1{F0: false, F1: false}
                                _ = fi77
                                _ = v78
                            }
                            var v79 int32 = int32(v1)
                            _ = fi76
                            _ = v79
                        }
                    } else {
                        v31 += (-(v31))
                        switch v47 {
                        case 269:
                            v47 = 681
                        case 2147483647:
                            for fi80 := 0; fi80 < 19; fi80++ {
                                v81 := []int{}
                                for fi82 := 0; fi82 < 13; fi82++ {
                                    if false {
                                        if len(v81) >= 2 {
                                            v81 = v81[1:len(v81)]
                                        }
                                    } else {
                                        fmt.Println(v47)
                                        var v83 uint32 = 2
                                        _ = v83
                                    }
                                    _ = fi82
                                }
                                var v84 uint16 = v1
                                _ = fi80
                                _ = v81
                                _ = v84
                            }
                        }
                        v3[(v47 >> ((852) & 63))] = true
                    }
                }
                v85 := GenS2{F0: "foo bar", F1: ((!((v31 / ((v31) | 1)) <= -1)) && true)}
                _ = v85
            }
        case 1:
            var v86 uint8 = 1
            _ = v86
        default:
            v87 := GenS3{F0: false, F1: 0, F2: 2}
            _ = v87
        }
    }
    switch "" {
    case "\n\t":
        v88 := GenS3{F0: true, F1: 128, F2: 100}
        _ = v88
    case "foo bar":
        v89 := GenS1{F0: (3.14159 < (((-2.25 / float64((float32(-2.25) / 1e10))) + ((1.5 - 0.1) + ((0.0 - float64((3.14159 + 0.1))) / (float64((float32(float64(0.1)) * (1e10 * -2.25))) * 1e10)))) * 0.1)), F1: (!true)}
        _ = v89
    default:
        fmt.Println(v4, v1)
    }
    v90 := []int{-2147483648, (v47 / ((v47) | 1)), v47}
    for fi91 := 0; fi91 < 1; fi91++ {
        fmt.Println(v1)
        var v92 float32 = 0.1
        _ = fi91
        _ = v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v90:", len(v90), v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v47
    _ = v48
    _ = v90
}
