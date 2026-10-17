package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 bool; F1 int64; F2 int64 }
type GenS12 struct { F0 int16 }
func genFn0(p0 string, p1 uint, p2 int32) uint16 {
    v4 := GenS1{F0: 251, F1: 0}
    fmt.Println(p1, p0)
    p0 = ("hello" + p0)
    for fi5 := 0; fi5 < 12; fi5++ {
        switch p0 {
        case "":
            for fi6 := 0; fi6 < 17; fi6++ {
                var v7 int8 = 14
                v8 := GenS1{F0: 2, F1: -400}
                switch "hello" {
                case "foo bar":
                    p0 = "\n\t"
                case "\n\t":
                    var v9 uint32 = 2
                    _ = v9
                default:
                    v7 = 0
                }
                _ = fi6
                _ = v7
                _ = v8
            }
        default:
            var v10 uint64 = 862
            _ = v10
        }
        _ = fi5
    }
    switch 127 {
    case 42:
        if ((-1 >= (p2 >> ((p1) & 63))) && false) {
            for fi11 := 0; fi11 < 12; fi11++ {
                fmt.Println(fi11, fi11)
                var v12 float64 = float64((1.5 * (0.1 + 1.5)))
                v13 := GenS1{F0: 1, F1: 0}
                _ = fi11
                _ = v12
                _ = v13
            }
        } else {
            if false {
                fmt.Println(p0, p2)
            }
            switch p0 {
            case "a":
                fmt.Println(p0, p1)
            case "":
                if false {
                    p2 += (p2 ^ -799)
                    var v14 int32 = (p2 - -2147483648)
                    _ = v14
                } else {
                    p2 = int32(p1)
                    var v15 int64 = -923
                    if true {
                        v16 := GenS2{F0: 0}
                        _ = v16
                    }
                    _ = v15
                }
            case "foo bar":
                p0 = p0
            }
            if false {
                v17 := GenS3{F0: 742, F1: (p0 + ("foo bar" + (("foo bar" + p0) + ("" + p0))))}
                for fi18 := 0; fi18 < 18; fi18++ {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        p0 = ("a" + "a")
                        p0 = p0
                        _ = fi19
                    }
                    switch fi18 {
                    case 2147483647:
                        var v20 int8 = 100
                        _ = v20
                    case 32767:
                        p0 = p0
                    case 907:
                        v21 := []int{114, fi18, (fi18 >> ((p1) & 63)), fi18}
                        _ = v21
                    default:
                        var v22 int64 = 446
                        _ = v22
                    }
                    _ = fi18
                }
                _ = v17
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}

func main() {
    var v1 uint8 = 2
    v2 := []string{("hello" + ("\n\t" + "a"))}
    var v3 uint = 255
    v23 := GenS4{F0: 42, F1: -2147483648, F2: 32767}
    fmt.Println(v1, v3)
    for fi24 := 0; fi24 < 12; fi24++ {
        if false {
            v3 = (v3 ^ 1)
            if false {
                if true {
                    v1 = (v1 >> ((4294967295) & 63))
                    for fi25 := 0; fi25 < 1; fi25++ {
                        var v26 float64 = ((-2.25 * 1.5) + (1e10 + 1.5))
                        fmt.Println(v3, fi25, v26)
                        v27 := GenS2{F0: 1000}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                var v28 uint = (v3 << (((v3 ^ 42)) & 63))
                _ = v28
            } else {
                fmt.Println(fi24, v3, fi24)
                for fi29 := 0; fi29 < 2; fi29++ {
                    v3 = (v3 / ((v3) | 1))
                    _ = fi29
                }
            }
            var v30 string = (("\n\t" + "hello") + "")
            _ = v30
        } else {
            v3 -= (v3 ^ 37)
        }
        switch "hello" {
        case "a":
            if true {
                v31 := GenS3{F0: uint64(v1), F1: "a"}
                var v32 uint16 = 2
                v3 += 0
                _ = v31
                _ = v32
            } else {
                v1 += (v1 * (^(v1)))
                for fi33 := 0; fi33 < 20; fi33++ {
                    v34 := GenS2{F0: 341}
                    var v35 uint8 = v1
                    v36 := GenS3{F0: 4294967295, F1: "\n\t"}
                    _ = fi33
                    _ = v34
                    _ = v35
                    _ = v36
                }
                v2 = append(v2, "")
            }
        case "\n\t":
            v1 += (^(v1))
        }
        _ = fi24
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 -= (v3 / (((v3 * v3)) | 1))
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi37 := 0; fi37 < 19; fi37++ {
            if false {
                if true {
                    var v38 uint16 = 643
                    var v39 float64 = float64(0.1)
                    _ = v38
                    _ = v39
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                switch ("\n\t" + ("a" + ((("foo bar" + ("" + "\n\t")) + ("foo bar" + ((((("a" + "\n\t") + "foo bar") + ("\n\t" + "hello")) + "foo bar") + "foo bar"))) + "hello"))) {
                case "foo bar":
                    switch 32768 {
                    case 0:
                        switch "foo bar" {
                        case "":
                            v1 = v1
                        case "foo bar":
                            var v40 uint = 141
                            _ = v40
                        }
                    case 42:
                        if true {
                            v3 = 0
                            fmt.Println(v1, v3, v3)
                            switch fi37 {
                            case 42:
                                v41 := GenS5{F0: fi37}
                                _ = v41
                            default:
                                if true {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(fi37, fi37)
                                    v3 = 1
                                }
                            }
                        } else {
                            v1 += 21
                            for fi42 := 0; fi42 < 3; fi42++ {
                                if (!(0 > -1)) {
                                    var v43 uint64 = 1
                                    var v44 int = fi37
                                    _ = v43
                                    _ = v44
                                }
                                _ = fi42
                            }
                        }
                    case 2147483647:
                        var v45 bool = true
                        _ = v45
                    default:
                        var v46 int64 = -246
                        _ = v46
                    }
                case "hello":
                    for fi47 := 0; fi47 < 17; fi47++ {
                        if ((true || (true || false)) || (-326 > -9223372036854775808)) {
                            var v48 float32 = (3.14159 - (float32(((0.1 - (((1.5 - 1e10) + 0.1) - 3.14159)) + 3.14159)) - (1.5 + (1.5 - float32(float64(((1.5 + (-2.25 * 3.14159)) / 1.0)))))))
                            v48 = (1.5 + 0.1)
                            _ = v48
                        } else {
                            var v49 bool = true
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v50 int64 = -755
                            if false {
                                switch fi37 {
                                case 4294967295:
                                    if (false || v49) {
                                        v51 := GenS4{F0: (^(v1)), F1: -1, F2: 65535}
                                        if v49 {
                                            v52 := GenS2{F0: 183}
                                            var v53 uint32 = 659
                                            _ = v52
                                            _ = v53
                                        }
                                        if (!v49) {
                                            switch -2147483648 {
                                            case -2147483648:
                                                if ((1.5 * 1e10) == 1.5) {
                                                    var v54 int16 = -32768
                                                    _ = v54
                                                }
                                            case 392:
                                                v50 = v50
                                            case 2147483647:
                                                v55 := GenS6{F0: 4294967295, F1: fi47}
                                                _ = v55
                                            default:
                                                v56 := GenS7{F0: (fi47 % ((fi37) | 1))}
                                                _ = v56
                                            }
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(v50)
                                            case "a":
                                                v1 -= 1
                                            }
                                        }
                                        _ = v51
                                    }
                                case 2147483647:
                                    _ = len(v2)
                                case -210:
                                    fmt.Println(v50, v3, v50)
                                }
                            }
                            _ = v49
                            _ = v50
                        }
                        fmt.Println(v1)
                        _ = fi47
                    }
                default:
                    fmt.Println(fi37, fi37)
                }
            } else {
                var v57 uint32 = 0
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v57
            }
            var v58 int = 1
            if (!true) {
                for fi59 := 0; fi59 < 7; fi59++ {
                    v60 := GenS8{F0: false}
                    _ = fi59
                    _ = v60
                }
            } else {
                v58 += fi37
                v61 := GenS9{F0: (^(v3))}
                for fi62 := 0; fi62 < 14; fi62++ {
                    v63 := GenS10{F0: 1}
                    switch "\n\t" {
                    case "a":
                        if (false && (v3 < v3)) {
                            var v64 int32 = int32(v3)
                            if ((true || false) && (1 == (fi37 | (fi37 << ((v3) & 63))))) {
                                var v65 uint = (v3 >> (((v3 ^ (v3 ^ v3))) & 63))
                                _ = v65
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v64
                        }
                    case "\n\t":
                        switch (("foo bar" + (("" + "") + ("\n\t" + ""))) + "") {
                        case "hello":
                            switch 4294967295 {
                            case -1:
                                v3 -= (v3 / ((v3) | 1))
                            case 2147483647:
                                var v66 int64 = 0
                                _ = v66
                            case 127:
                                v67 := GenS11{F0: true, F1: 804, F2: -9223372036854775808}
                                _ = v67
                            default:
                                var v68 float64 = float64((1.5 / 3.14159))
                                _ = v68
                            }
                        }
                    case "foo bar":
                        if true {
                            if false {
                                v58 += (fi62 ^ -2147483648)
                                fmt.Println(fi62)
                                v3 = v3
                            }
                        } else {
                            switch "" {
                            case "":
                                v3 += (v3 & v3)
                            case "hello":
                                var v69 uint32 = 2147483647
                                _ = v69
                            default:
                                _ = len(v2)
                            }
                            if (true || false) {
                                v58 = (v58 & 462)
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v70 string = "a"
                                _ = v70
                            }
                        }
                    default:
                        var v71 uint16 = 0
                        _ = v71
                    }
                    switch (v58 + v58) {
                    case 239:
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v58, v58)
                        } else {
                            if (true || (!("a" <= ("a" + "\n\t")))) {
                                var v72 bool = (1 <= -2147483648)
                                v2 = append(v2, "a")
                                var v73 int32 = 1
                                _ = v72
                                _ = v73
                            } else {
                                if true {
                                    _ = len(v2)
                                    v58 -= (-(fi62))
                                    if true {
                                        v58 = fi37
                                        var v74 string = "hello"
                                        _ = v74
                                    }
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                    v3 = (v3 >> ((49) & 63))
                                }
                            }
                            _ = len(v2)
                            fmt.Println(v3, fi37, v1)
                        }
                    default:
                        if (!(-2.25 >= ((1e10 + 1.5) / (-2.25 / (-2.25 / 1e10))))) {
                            v75 := GenS12{F0: 0}
                            switch "\n\t" {
                            case "foo bar":
                                switch (v58 * (-(v58))) {
                                case -747:
                                    fmt.Println(v3, v58)
                                case 801:
                                    if false {
                                        v1 = v1
                                        var v76 uint32 = uint32(v3)
                                        _ = v76
                                    }
                                }
                                var v77 int = v58
                                fmt.Println(v3, fi37, fi37)
                                _ = v77
                            case "\n\t":
                                var v78 uint16 = 719
                                _ = v78
                            }
                            var v79 uint64 = 750
                            _ = v75
                            _ = v79
                        }
                    }
                    _ = fi62
                    _ = v63
                }
                _ = v61
            }
            _ = fi37
            _ = v58
        }
    }
    if true {
        var v80 uint32 = 326
        _ = v80
    }
    var v81 int8 = -128
    if true {
        var v82 int64 = 985
        var v83 uint16 = 32767
        v3 += (v3 * v3)
        _ = v82
        _ = v83
    }
    switch "foo bar" {
    case "hello":
        for fi84 := 0; fi84 < 5; fi84++ {
            for fi85 := 0; fi85 < 8; fi85++ {
                if true {
                    switch fi84 {
                    case 0:
                        v81 = (v81 - (-(v81)))
                        var v86 int8 = (-(v81))
                        _ = v86
                    case 331:
                        var v87 int32 = 2147483647
                        _ = v87
                    default:
                        v88 := GenS9{F0: (v3 & v3)}
                        _ = v88
                    }
                }
                v89 := GenS3{F0: 2, F1: "a"}
                var v90 float32 = 1e10
                _ = fi85
                _ = v89
                _ = v90
            }
            _ = fi84
        }
    case "":
        v3 = (v3 % ((42) | 1))
    }
    _ = len(v2)
    switch ("a" + "\n\t") {
    case "hello":
        v1 -= (v1 ^ (v1 - 0))
    default:
        for fi91 := 0; fi91 < 12; fi91++ {
            _ = len(v2)
            fmt.Println(fi91, v1, v1)
            _ = fi91
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v81
}
