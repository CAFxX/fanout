package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int16; F1 int8; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 uint32; F1 bool }
type GenS11 struct { F0 int8; F1 int32; F2 uint8 }

func main() {
    var v1 uint8 = 0
    var v2 float32 = 3.14159
    var v3 int16 = 1
    v4 := []string{("" + ""), "foo bar"}
    for fi5 := 0; fi5 < 13; fi5++ {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        v6 := GenS1{F0: 2, F1: 799}
        v7 := GenS1{F0: 42, F1: 0}
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v2)
    var v8 float32 = -2.25
    if (1 <= uint16(v1)) {
        switch "hello" {
        case "foo bar":
            if false {
                fmt.Println(v8, v3, v2)
                for fi9 := 0; fi9 < 13; fi9++ {
                    v8 -= (v2 - (v2 * (0.0 / 1.0)))
                    var v10 string = "hello"
                    _ = fi9
                    _ = v10
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        case "\n\t":
            if false {
                var v11 uint = 388
                for fi12 := 0; fi12 < 3; fi12++ {
                    v13 := GenS1{F0: 1, F1: 32767}
                    fmt.Println(v2, v11)
                    v14 := GenS2{F0: 0, F1: (^(v3))}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v15 := GenS2{F0: 516, F1: (v3 % (((v3 / ((v3) | 1))) | 1))}
                _ = v11
                _ = v15
            } else {
                if false {
                    v8 -= (float32(((3.14159 + (float64((v8 * v8)) / 1e10)) - ((0.0 * (1.5 / float64((v2 - ((v2 / v2) * (3.14159 * v2)))))) + (0.1 * -2.25)))) + v2)
                    switch "" {
                    case "hello":
                        v8 += v2
                    case "":
                        for fi16 := 0; fi16 < 2; fi16++ {
                            if true {
                                v1 += (v1 + v1)
                                v17 := GenS2{F0: -777, F1: -1}
                                var v18 float32 = v8
                                _ = v17
                                _ = v18
                            }
                            _ = len(v4)
                            _ = fi16
                        }
                    case "\n\t":
                        if false {
                            if false {
                                fmt.Println(v2, v3)
                                v19 := GenS3{F0: (v3 << ((42) & 63)), F1: 42, F2: "a"}
                                v2 = v2
                                _ = v19
                            }
                            v20 := GenS2{F0: -2147483648, F1: -32768}
                            _ = v20
                        }
                    default:
                        if false {
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                        } else {
                            fmt.Println(v3, v2)
                            if (0 < 42) {
                                fmt.Println(v3)
                            }
                            switch 32768 {
                            case 42:
                                fmt.Println(v1, v2, v1)
                            case -2147483648:
                                var v21 uint16 = 905
                                _ = v21
                            case -15:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                    }
                    v22 := GenS4{F0: v3}
                    _ = v22
                } else {
                    var v23 float32 = 3.14159
                    for fi24 := 0; fi24 < 2; fi24++ {
                        v25 := GenS1{F0: 42, F1: 32767}
                        _ = fi24
                        _ = v25
                    }
                    if true {
                        if true {
                            if true {
                                var v26 int32 = 0
                                _ = v26
                            } else {
                                v2 += 1.5
                            }
                        }
                        var v27 int = 4294967295
                        fmt.Println(v8, v23)
                        _ = v27
                    } else {
                        var v28 int8 = 1
                        var v29 int16 = v3
                        _ = v28
                        _ = v29
                    }
                    _ = v23
                }
            }
        case "a":
            fmt.Println(v8, v3, v2)
        }
        v4 = append(v4, (("" + "hello") + "a"))
    } else {
        var v30 float64 = -2.25
        _ = v30
    }
    if (2 == 205) {
        v31 := GenS5{F0: 65535}
        v1 = 128
        var v32 bool = ((((18446744073709551615 != 18446744073709551615) && (-297 >= -630)) && (!false)) || (((false || true) && (false && (3.14159 != ((v8 + v8) + (v2 * 0.0))))) || true))
        _ = v31
        _ = v32
    }
    fmt.Println(v1, v1, v8)
    for fi33 := 0; fi33 < 19; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            switch "" {
            case "foo bar":
                v3 -= -1
            case "":
                var v35 int32 = 0
                _ = v35
            case "\n\t":
                if false {
                    v8 = 0.1
                    v36 := GenS4{F0: 1000}
                    _ = v36
                } else {
                    var v37 uint16 = uint16(v1)
                    _ = v37
                }
            }
            _ = fi34
        }
        v4 = append(v4, ("\n\t" + ("hello" + ("hello" + ""))))
        _ = fi33
    }
    fmt.Println(v2, v8)
    var v38 bool = (!true)
    switch "foo bar" {
    case "":
        for fi39 := 0; fi39 < 3; fi39++ {
            v38 = (0 >= 613)
            _ = fi39
        }
    case "foo bar":
        v2 = ((v8 + 1.5) + v8)
    case "\n\t":
        var v40 float32 = v2
        _ = v40
    }
    var v41 float32 = (v8 * (1.5 / 1.0))
    if true {
        v41 += v2
        v2 -= 0.0
    } else {
        switch "" {
        case "":
            for fi42 := 0; fi42 < 3; fi42++ {
                switch fi42 {
                case 1:
                    var v43 uint = uint(v1)
                    _ = v43
                case -291:
                    v44 := GenS6{F0: -1, F1: v38}
                    _ = v44
                case 691:
                    var v45 uint64 = 4294967295
                    _ = v45
                default:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                for fi46 := 0; fi46 < 19; fi46++ {
                    if v38 {
                        v47 := GenS2{F0: -1, F1: -32768}
                        v3 = (v3 - (-(v3)))
                        switch (("\n\t" + ("" + ("hello" + "hello"))) + ("hello" + ("" + ("foo bar" + "\n\t")))) {
                        case "":
                            v38 = (!(42 < 42))
                        case "foo bar":
                            switch 128 {
                            case 128:
                                var v48 int32 = -683
                                _ = v48
                            case -804:
                                var v49 string = "foo bar"
                                _ = v49
                            default:
                                var v50 float64 = 0.1
                                _ = v50
                            }
                        }
                        _ = v47
                    } else {
                        var v51 uint8 = 170
                        _ = v51
                    }
                    var v52 int16 = (v3 % ((v3) | 1))
                    _ = fi46
                    _ = v52
                }
                v53 := GenS4{F0: (v3 << ((275) & 63))}
                _ = fi42
                _ = v53
            }
        case "\n\t":
            switch 127 {
            case -581:
                for fi54 := 0; fi54 < 3; fi54++ {
                    var v55 int64 = 42
                    _ = fi54
                    _ = v55
                }
            case 9223372036854775807:
                for fi56 := 0; fi56 < 13; fi56++ {
                    switch "hello" {
                    case "a":
                        v8 = ((float32(-2.25) - v8) - v8)
                    case "hello":
                        if (69 < -1) {
                            fmt.Println(v2)
                            for fi57 := 0; fi57 < 8; fi57++ {
                                _ = len(v4)
                                var v58 int16 = (v3 - v3)
                                v1 -= v1
                                _ = fi57
                                _ = v58
                            }
                        } else {
                            v4 = append(v4, "")
                        }
                    case "\n\t":
                        v59 := GenS7{F0: 58}
                        _ = v59
                    default:
                        for fi60 := 0; fi60 < 8; fi60++ {
                            v61 := GenS4{F0: v3}
                            v8 += (((float32(1.5) * v8) - (v2 / 3.14159)) * (1.5 - (float32(3.14159) / float32(float64(v8)))))
                            _ = fi60
                            _ = v61
                        }
                    }
                    switch (fi56 << ((65535) & 63)) {
                    case 9223372036854775807:
                        var v62 int32 = 2147483647
                        _ = v62
                    case 42:
                        var v63 int64 = -9223372036854775808
                        _ = v63
                    default:
                        var v64 uint = 0
                        _ = v64
                    }
                    switch ("\n\t" + "") {
                    case "hello":
                        switch "hello" {
                        case "a":
                            if len(v4) > 0 {
                                v4[0] = "a"
                            }
                        case "\n\t":
                            var v65 uint64 = 4294967295
                            _ = v65
                        }
                    case "\n\t":
                        if v38 {
                            for fi66 := 0; fi66 < 6; fi66++ {
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                                if v38 {
                                    var v67 uint64 = 0
                                    v1 = 202
                                    var v68 int64 = 752
                                    switch "\n\t" {
                                    case "":
                                        var v69 float32 = v8
                                        _ = v69
                                    case "hello":
                                        fmt.Println(fi56, fi56)
                                    case "a":
                                        v4 = append(v4, "a")
                                    }
                                    switch (("a" + ("\n\t" + ("\n\t" + ("" + "foo bar")))) + "foo bar") {
                                    case "foo bar":
                                        v70 := GenS8{F0: (v1 < (v1 >> ((0) & 63)))}
                                        _ = v70
                                    default:
                                        var v71 uint = 2
                                        _ = v71
                                    }
                                    _ = v67
                                    _ = v68
                                } else {
                                    v1 = v1
                                    var v72 int32 = -583
                                    _ = v72
                                }
                                v1 -= 71
                                v1 += v1
                                if ((32767 == v3) && v38) {
                                    if v38 {
                                        v73 := GenS9{F0: 353, F1: 1}
                                        var v74 uint = 141
                                        _ = v73
                                        _ = v74
                                    }
                                    if ((925 != 1) && (191 >= 65535)) {
                                        var v75 float64 = 3.14159
                                        _ = v75
                                    } else {
                                        switch fi66 {
                                        case -1:
                                            var v76 float64 = -2.25
                                            _ = v76
                                        case 564:
                                            v77 := GenS1{F0: 4294967295, F1: 65535}
                                            _ = v77
                                        case -2147483648:
                                            var v78 int16 = v3
                                            v79 := GenS4{F0: v3}
                                            fmt.Println(v41, v78)
                                            _ = v78
                                            _ = v79
                                        }
                                    }
                                    v80 := GenS10{F0: 1, F1: ((0.0 <= (v41 - ((((1e10 + v8) / 1e10) * -2.25) * (v41 / v41)))) || (1 == -1))}
                                    _ = v80
                                }
                                _ = fi66
                            }
                            for fi81 := 0; fi81 < 20; fi81++ {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                fmt.Println(v2)
                                _ = fi81
                            }
                            v82 := GenS11{F0: -128, F1: 1, F2: (v1 & (v1 * v1))}
                            _ = v82
                        } else {
                            v4 = append(v4, "hello")
                        }
                    default:
                        _ = len(v4)
                    }
                    _ = fi56
                }
            default:
                for fi83 := 0; fi83 < 15; fi83++ {
                    for fi84 := 0; fi84 < 17; fi84++ {
                        if len(v4) > 0 {
                            v4[0] = "foo bar"
                        }
                        var v85 string = (("hello" + "") + "a")
                        _ = fi84
                        _ = v85
                    }
                    _ = fi83
                }
            }
        }
        if v38 {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        var v86 int8 = 2
        _ = v86
    }
    switch 127 {
    case 934:
        var v87 string = ""
        _ = v87
    case -2147483648:
        var v88 bool = v38
        _ = v88
    default:
        switch 32768 {
        case 0:
            if false {
                v89 := GenS9{F0: 255, F1: 1}
                for fi90 := 0; fi90 < 14; fi90++ {
                    var v91 int8 = 1
                    _ = fi90
                    _ = v91
                }
                _ = v89
            }
        case 498:
            v4 = append(v4, "foo bar")
        default:
            v1 -= (v1 + (v1 & v1))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v8
}
