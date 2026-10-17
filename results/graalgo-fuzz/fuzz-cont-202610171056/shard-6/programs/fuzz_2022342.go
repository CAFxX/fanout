package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint; F2 int64 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 uint32; F1 int8; F2 string }
type GenS7 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 string; F1 int8 }
func genFn0(p0 int64) int8 {
    for fi7 := 0; fi7 < 2; fi7++ {
        p0 -= p0
        if true {
            p0 = (p0 * p0)
        }
        _ = fi7
    }
    p0 -= p0
    v8 := GenS1{F0: 2147483647, F1: -2147483648, F2: 0}
    _ = p0
    _ = v8
    return -19
}

func main() {
    var v1 uint32 = 725
    var v2 int = 130
    var v3 uint64 = 2
    var v4 uint32 = (v1 * v1)
    v5 := []uint32{v4, (v1 & 2147483647), v4}
    v6 := make(map[string]bool)
    v9 := GenS2{F0: 42, F1: (v3 >> ((18446744073709551615) & 63))}
    var v10 int16 = 42
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch ("\n\t" + "hello") {
    case "foo bar":
        for fi11 := 0; fi11 < 4; fi11++ {
            v6["a"] = (42 > 1)
            var v12 int = (fi11 | fi11)
            var v13 int32 = -1
            _ = fi11
            _ = v12
            _ = v13
        }
    case "hello":
        switch -9223372036854775808 {
        case -452:
            v2 -= (v2 ^ (-(v2)))
        case -952:
            switch v2 {
            case -895:
                v14 := GenS3{F0: 1}
                _ = v14
            case 32767:
                for fi15 := 0; fi15 < 13; fi15++ {
                    delete(v6, "foo bar")
                    v16 := GenS3{F0: 1}
                    _ = fi15
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 17; fi17++ {
                    if (float32(3.14159) <= -2.25) {
                        v1 = v4
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            v18 := GenS2{F0: 128, F1: v3}
                            _ = v18
                        case "foo bar":
                            for fi19 := 0; fi19 < 4; fi19++ {
                                switch ((((("\n\t" + "a") + "") + "a") + "a") + ("\n\t" + "hello")) {
                                case "\n\t":
                                    var v20 uint16 = 692
                                    _ = v20
                                }
                                delete(v6, "hello")
                                _ = fi19
                            }
                        case "hello":
                            if true {
                                if (v10 != 32767) {
                                    switch "a" {
                                    case "":
                                        for fi21 := 0; fi21 < 18; fi21++ {
                                            if true {
                                                if false {
                                                    var v22 bool = ((0.1 > float32((0.0 + (3.14159 + ((1e10 + (1e10 + (1.5 + 1.5))) * 1.5))))) || true)
                                                    v22 = false
                                                    v23 := GenS4{F0: "foo bar", F1: 18446744073709551615, F2: 1}
                                                    v10 -= v10
                                                    _ = v22
                                                    _ = v23
                                                } else {
                                                    v10 -= -771
                                                    v24 := GenS5{F0: v10, F1: 79}
                                                    var v25 int8 = -13
                                                    _ = v24
                                                    _ = v25
                                                }
                                            }
                                            var v26 uint16 = 65535
                                            _ = fi21
                                            _ = v26
                                        }
                                    default:
                                        v1 = 0
                                    }
                                    if (v2 > v2) {
                                        var v27 string = "hello"
                                        switch "\n\t" {
                                        case "\n\t":
                                            fmt.Println(v1, v1)
                                        default:
                                            if (42 <= genFn0(252)) {
                                                for fi28 := 0; fi28 < 6; fi28++ {
                                                    v10 -= (v10 - 32767)
                                                    _ = fi28
                                                }
                                                var v29 uint16 = 255
                                                var v30 uint8 = 42
                                                _ = v29
                                                _ = v30
                                            }
                                        }
                                        var v31 uint32 = 65535
                                        _ = v27
                                        _ = v31
                                    } else {
                                        v32 := GenS6{F0: v4, F1: 0, F2: ("a" + ("hello" + "a"))}
                                        v3 = (v3 & 517)
                                        v6[("" + "foo bar")] = true
                                        _ = v32
                                    }
                                    v4 = 4294967295
                                } else {
                                    if (!false) {
                                        delete(v6, ("" + ("\n\t" + "\n\t")))
                                    }
                                    if ((v10 % (((v10 | 976)) | 1)) > (v10 + -32768)) {
                                        switch 9223372036854775807 {
                                        case -164:
                                            var v33 int64 = -454
                                            _ = v33
                                        case 4294967295:
                                            switch v2 {
                                            case -916:
                                                v10 -= -1
                                            case 4294967295:
                                                _ = len(v5)
                                            default:
                                                v34 := GenS7{F0: (v3 / (((v3 >> ((1) & 63))) | 1)), F1: 1, F2: 127}
                                                _ = v34
                                            }
                                        default:
                                            v3 = v3
                                        }
                                        if false {
                                            v10 += (^(v10))
                                            if false {
                                                v5 = append(v5, (^(v1)))
                                                var v35 uint16 = 364
                                                delete(v6, "\n\t")
                                                _ = v35
                                            } else {
                                                v4 -= (v4 >> ((42) & 63))
                                                var v36 uint8 = 127
                                                _ = len(v5)
                                                _ = v36
                                            }
                                            v10 = (v10 * -1)
                                        }
                                        switch fi17 {
                                        case -828:
                                            switch v2 {
                                            case 32768:
                                                fmt.Println(v4, fi17)
                                            case 2147483647:
                                                for fi37 := 0; fi37 < 5; fi37++ {
                                                    fmt.Println(v2, v3)
                                                    if true {
                                                        v38 := GenS7{F0: 42, F1: 255, F2: genFn0(42)}
                                                        var v39 bool = (false || true)
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v10 += v10
                                                        default:
                                                            var v40 int8 = 31
                                                            _ = v40
                                                        }
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                    _ = fi37
                                                }
                                            default:
                                                v4 -= 523
                                            }
                                        }
                                    } else {
                                        if ((float64(3.14159) / float64((-2.25 - float32((1e10 + 1e10))))) < (0.1 / ((0.0 * (3.14159 / (1.5 - (-2.25 - (3.14159 + 0.1))))) + 1e10))) {
                                            if false {
                                                if false {
                                                    if (472 <= (v3 ^ 905)) {
                                                        for fi41 := 0; fi41 < 19; fi41++ {
                                                            switch "" {
                                                            case "\n\t":
                                                                fmt.Println(fi41)
                                                            case "a":
                                                                if (v10 == -1) {
                                                                    var v42 bool = false
                                                                    v2 = v2
                                                                    v43 := GenS4{F0: ("foo bar" + ""), F1: 1, F2: 9223372036854775807}
                                                                    _ = v42
                                                                    _ = v43
                                                                } else {
                                                                    v2 = 281
                                                                    if true {
                                                                        if false {
                                                                            if (1 == 2147483647) {
                                                                                var v44 int8 = 0
                                                                                v3 -= (v3 & 2)
                                                                                switch (-(fi17)) {
                                                                                case 839:
                                                                                    v1 -= (v4 % ((v1) | 1))
                                                                                case 795:
                                                                                    var v45 float64 = (-2.25 - float64(float32(-2.25)))
                                                                                    _ = v45
                                                                                case -9223372036854775808:
                                                                                    v44 = 96
                                                                                default:
                                                                                    var v46 float64 = (-2.25 / 1.0)
                                                                                    _ = v46
                                                                                }
                                                                                _ = v44
                                                                            }
                                                                            if (true && false) {
                                                                                switch (^(fi41)) {
                                                                                case 128:
                                                                                    if (!false) {
                                                                                        var v47 uint = 18446744073709551615
                                                                                        if (0.0 != (1e10 * 1e10)) {
                                                                                            v48 := GenS1{F0: 2147483647, F1: (fi41 / (((fi17 * (^(fi17)))) | 1)), F2: 4294967295}
                                                                                            var v49 float64 = float64(0.0)
                                                                                            var v50 int8 = 100
                                                                                            _ = v48
                                                                                            _ = v49
                                                                                            _ = v50
                                                                                        }
                                                                                        v51 := GenS8{F0: v47}
                                                                                        _ = v47
                                                                                        _ = v51
                                                                                    }
                                                                                default:
                                                                                    var v52 uint8 = 63
                                                                                    _ = v52
                                                                                }
                                                                            }
                                                                            if false {
                                                                                v53 := GenS4{F0: "\n\t", F1: 0, F2: int64(v1)}
                                                                                fmt.Println(v3)
                                                                                _ = v53
                                                                            } else {
                                                                                v54 := v6[""]
                                                                                var v55 float64 = 0.0
                                                                                _ = v54
                                                                                _ = v55
                                                                            }
                                                                            delete(v6, "a")
                                                                        }
                                                                        fmt.Println(v10)
                                                                        var v56 uint = 42
                                                                        _ = v56
                                                                    }
                                                                }
                                                            case "":
                                                                v1 = v4
                                                            }
                                                            if (0 >= 1) {
                                                                var v57 bool = false
                                                                switch fi41 {
                                                                case 32767:
                                                                    var v58 int16 = v10
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        if len(v5) > 0 {
                                                                            v5[0] = v4
                                                                        }
                                                                    default:
                                                                        if (255 <= 0) {
                                                                            if ((v3 & (^(v3))) <= v3) {
                                                                                v59 := GenS9{F0: (v58 + v58)}
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                                _ = v59
                                                                            }
                                                                            v60 := GenS2{F0: 2, F1: 18446744073709551615}
                                                                            var v61 float64 = 1e10
                                                                            v2 += (v2 << ((140) & 63))
                                                                            v10 = (v10 << ((65535) & 63))
                                                                            _ = v60
                                                                            _ = v61
                                                                        }
                                                                    }
                                                                    v3 = (v3 / ((v3) | 1))
                                                                    _ = v58
                                                                }
                                                                _ = v57
                                                            } else {
                                                                fmt.Println(v1)
                                                            }
                                                            v3 = v3
                                                            v3 = (v3 % (((v3 << ((639) & 63))) | 1))
                                                            _ = len(v6)
                                                            _ = fi41
                                                        }
                                                        for fi62 := 0; fi62 < 5; fi62++ {
                                                            switch ("hello" + "hello") {
                                                            case "hello":
                                                                if ((^(fi17)) < fi17) {
                                                                    var v63 int64 = 42
                                                                    v2 += (fi62 ^ (fi62 * (v2 + (v2 >> ((4294967295) & 63)))))
                                                                    v64 := GenS3{F0: 42}
                                                                    _ = v63
                                                                    _ = v64
                                                                } else {
                                                                    switch v2 {
                                                                    case 1:
                                                                        var v65 string = "foo bar"
                                                                        _ = v65
                                                                    case -1:
                                                                        var v66 int8 = 127
                                                                        _ = v66
                                                                    }
                                                                }
                                                            case "a":
                                                                fmt.Println(v3, v2)
                                                            default:
                                                                var v67 int32 = -3
                                                                _ = v67
                                                            }
                                                            var v68 int8 = 55
                                                            var v69 int = 32768
                                                            switch (("\n\t" + "hello") + "\n\t") {
                                                            case "hello":
                                                                var v70 uint8 = 255
                                                                _ = v70
                                                            default:
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            }
                                                            _ = fi62
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                    }
                                                    var v71 bool = (((1e10 * 0.1) > (1e10 - (float64((float32(float64(0.1)) + 3.14159)) / 3.14159))) && true)
                                                    v3 = 42
                                                    _ = v71
                                                }
                                                delete(v6, ("" + "foo bar"))
                                                fmt.Println(v4, v1)
                                            }
                                        }
                                    }
                                    delete(v6, "a")
                                }
                                for fi72 := 0; fi72 < 16; fi72++ {
                                    v5 = append(v5, 2147483647)
                                    _ = fi72
                                }
                                _ = len(v5)
                            } else {
                                for fi73 := 0; fi73 < 17; fi73++ {
                                    fmt.Println(v4, v3, fi73)
                                    _ = len(v5)
                                    v74 := GenS6{F0: (v4 >> ((0) & 63)), F1: 1, F2: ("foo bar" + ("hello" + "foo bar"))}
                                    switch ("foo bar" + "") {
                                    case "foo bar":
                                        var v75 string = "hello"
                                        _ = v75
                                    }
                                    _ = fi73
                                    _ = v74
                                }
                                var v76 float32 = 0.0
                                _ = v76
                            }
                        }
                        if false {
                            var v77 string = ""
                            _ = v77
                        } else {
                            if true {
                                switch "" {
                                case "\n\t":
                                    var v78 uint64 = 709
                                    _ = v78
                                case "a":
                                    v79 := GenS9{F0: (v10 ^ v10)}
                                    _ = v79
                                case "":
                                    v10 -= v10
                                default:
                                    var v80 uint16 = 318
                                    _ = v80
                                }
                            }
                            if len(v5) > 0 {
                                v5[0] = v1
                            }
                            var v81 uint8 = 127
                            _ = v81
                        }
                    } else {
                        var v82 int = (v2 >> ((2) & 63))
                        _ = v82
                    }
                    _ = fi17
                }
            }
        default:
            if (2 > 32767) {
                var v83 int64 = 1
                if true {
                    v84 := GenS6{F0: (v1 / ((v1) | 1)), F1: genFn0((v83 | (^(v83)))), F2: ("foo bar" + "a")}
                    _ = v84
                } else {
                    var v85 uint32 = 4294967295
                    _ = v85
                }
                _ = v83
            } else {
                var v86 uint = 42
                v3 = v3
                var v87 int64 = -9223372036854775808
                _ = v86
                _ = v87
            }
        }
    }
    if false {
        switch "a" {
        case "foo bar":
            for fi88 := 0; fi88 < 12; fi88++ {
                v6[(((("" + "foo bar") + (("foo bar" + "foo bar") + ("a" + "a"))) + ("a" + ("a" + (((("foo bar" + ((((("\n\t" + "a") + ("foo bar" + ((("foo bar" + "\n\t") + "hello") + ("hello" + "foo bar")))) + "foo bar") + "a") + "\n\t")) + "a") + "hello") + ("" + "foo bar"))))) + "\n\t")] = (true && ((^(v10)) == 1))
                v1 = 42
                _ = fi88
            }
        case "hello":
            fmt.Println(v4, v2)
        default:
            v89 := GenS3{F0: genFn0(-1)}
            _ = v89
        }
        v10 -= (-(v10))
        v1 = (v4 << ((313) & 63))
    }
    if len(v5) > 0 {
        v5[0] = (v4 & 2)
    }
    fmt.Println(v3, v1, v10)
    var v90 bool = false
    for fi91 := 0; fi91 < 5; fi91++ {
        switch v2 {
        case 346:
            v92 := GenS10{F0: ("hello" + "a"), F1: -128}
            _ = v92
        case 32768:
            fmt.Println(v4)
        }
        for fi93 := 0; fi93 < 16; fi93++ {
            v5 = append(v5, (v4 % ((v4) | 1)))
            _ = fi93
        }
        v90 = v90
        _ = fi91
    }
    v2 = (v2 / (((v2 << ((494) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v90:", v90)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
    _ = v90
}
