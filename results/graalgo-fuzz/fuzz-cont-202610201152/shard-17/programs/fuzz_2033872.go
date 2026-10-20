package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16; F1 int32; F2 uint }
type GenS8 struct { F0 uint32; F1 uint; F2 string }
type GenS9 struct { F0 uint32; F1 uint }
func genFn0() uint {
    v5 := []int{128, -2147483648}
    if len(v5) > 0 {
        v5[0] = 127
    }
    v6 := GenS1{F0: 32768, F1: -1}
    var v7 int8 = 86
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint16 = 212
    var v2 float32 = -2.25
    var v3 string = ("" + "")
    var v4 uint64 = 4294967295
    v8 := []int{443, 42, -15}
    v3 = ("foo bar" + (("\n\t" + ((((((v3 + v3) + v3) + "foo bar") + "foo bar") + v3) + (("\n\t" + v3) + (v3 + ((v3 + (v3 + (v3 + v3))) + ("" + v3)))))) + v3))
    v3 = (((v3 + v3) + v3) + "")
    if (1 <= 1) {
        if true {
            if (1 < 1) {
                var v9 uint32 = 2147483647
                _ = v9
            } else {
                v8 = append(v8, -2147483648)
                if (true && false) {
                    if (245 <= 2147483648) {
                        var v10 uint8 = 255
                        fmt.Println(v3, v10)
                        switch 0 {
                        case 32767:
                            v2 = 0.0
                        case 4294967295:
                            if (!false) {
                                switch (("" + "\n\t") + v3) {
                                case "a":
                                    var v11 float64 = 0.0
                                    _ = v11
                                }
                                var v12 int = 127
                                _ = v12
                            }
                        case -634:
                            fmt.Println(v3)
                        }
                        _ = v10
                    }
                    v2 = ((((v2 - v2) * v2) + v2) + (v2 / (v2 * v2)))
                } else {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        v4 += (^(v4))
                        for fi14 := 0; fi14 < 13; fi14++ {
                            var v15 float32 = v2
                            _ = fi14
                            _ = v15
                        }
                        if ((fi13 ^ (fi13 ^ fi13)) < fi13) {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            v8 = append(v8, fi13)
                        } else {
                            v1 = v1
                            v2 = (-2.25 * 0.0)
                        }
                        _ = fi13
                    }
                    var v16 int8 = 0
                    switch 127 {
                    case 9223372036854775807:
                        for fi17 := 0; fi17 < 3; fi17++ {
                            if false {
                                switch v3 {
                                case "\n\t":
                                    if (!((true && (fi17 > -226)) && true)) {
                                        v18 := GenS1{F0: fi17, F1: -1}
                                        _ = len(v8)
                                        switch "foo bar" {
                                        case "hello":
                                            _ = len(v8)
                                        }
                                        _ = v18
                                    }
                                default:
                                    if (false || (!true)) {
                                        if true {
                                            for fi19 := 0; fi19 < 10; fi19++ {
                                                if (((!false) && false) && (v3 > "hello")) {
                                                    v8 = append(v8, fi17)
                                                    var v20 uint8 = 2
                                                    var v21 int = (fi17 * fi17)
                                                    v22 := GenS1{F0: (v21 / ((fi19) | 1)), F1: 1}
                                                    _ = v20
                                                    _ = v21
                                                    _ = v22
                                                } else {
                                                    fmt.Println(v1, v1, v1)
                                                    _ = len(v8)
                                                }
                                                switch v3 {
                                                case "foo bar":
                                                    v23 := GenS2{F0: 255, F1: 744, F2: v16}
                                                    _ = v23
                                                }
                                                fmt.Println(v16, v3)
                                                _ = fi19
                                            }
                                            if (v4 < 4294967295) {
                                                var v24 uint = 42
                                                if (v4 > v4) {
                                                    for fi25 := 0; fi25 < 7; fi25++ {
                                                        v26 := GenS3{F0: 0}
                                                        _ = fi25
                                                        _ = v26
                                                    }
                                                }
                                                _ = v24
                                            } else {
                                                var v27 int8 = v16
                                                _ = v27
                                            }
                                        } else {
                                            v16 -= v16
                                        }
                                    }
                                }
                            }
                            for fi28 := 0; fi28 < 10; fi28++ {
                                v3 = (("" + (("a" + "hello") + ("foo bar" + "a"))) + "hello")
                                if len(v8) > 0 {
                                    v8[0] = (-(fi17))
                                }
                                _ = fi28
                            }
                            _ = fi17
                        }
                    default:
                        var v29 uint64 = 855
                        _ = v29
                    }
                    _ = v16
                }
            }
            var v30 string = (v3 + ("hello" + ""))
            if true {
                fmt.Println(v4, v4)
            } else {
                v31 := GenS4{F0: 128, F1: 947}
                var v32 int = -174
                v8 = append(v8, v32)
                _ = v31
                _ = v32
            }
            _ = v30
        }
    } else {
        for fi33 := 0; fi33 < 12; fi33++ {
            switch v3 {
            case "hello":
                switch (fi33 % ((250) | 1)) {
                case -538:
                    var v34 uint64 = (v4 % (((v4 - (v4 + 18446744073709551615))) | 1))
                    _ = v34
                case 127:
                    var v35 uint16 = v1
                    _ = v35
                case 32768:
                    for fi36 := 0; fi36 < 6; fi36++ {
                        var v37 int8 = 127
                        v2 = 0.1
                        _ = fi36
                        _ = v37
                    }
                default:
                    v2 = float32(float64(((-2.25 / 1.5) - float32(1e10))))
                }
            case "foo bar":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            for fi38 := 0; fi38 < 12; fi38++ {
                if len(v8) > 0 {
                    v8[0] = fi33
                }
                v1 += 969
                _ = fi38
            }
            _ = fi33
        }
        var v39 int64 = 9223372036854775807
        _ = v39
    }
    fmt.Println(v1, v1, v4)
    for fi40 := 0; fi40 < 17; fi40++ {
        v2 -= float32((3.14159 - 1e10))
        _ = fi40
    }
    var v41 uint8 = 214
    if false {
        var v42 float32 = v2
        var v43 float32 = v42
        _ = v42
        _ = v43
    } else {
        if false {
            if len(v8) > 0 {
                v8[0] = -9223372036854775808
            }
            if (!(128 >= v41)) {
                v44 := GenS4{F0: (v41 >> ((18446744073709551615) & 63)), F1: 42}
                v4 -= (v4 | 18446744073709551615)
                var v45 uint8 = (v41 << ((586) & 63))
                _ = v44
                _ = v45
            }
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        v8 = append(v8, 202)
    }
    v46 := GenS3{F0: 1000}
    if (((float32(0.1) - v2) * v2) >= v2) {
        for fi47 := 0; fi47 < 17; fi47++ {
            v3 = (v3 + "\n\t")
            v2 = -2.25
            for fi48 := 0; fi48 < 16; fi48++ {
                if false {
                    switch v3 {
                    case "a":
                        var v49 uint32 = 1
                        _ = v49
                    default:
                        var v50 uint32 = 4294967295
                        _ = v50
                    }
                    switch (fi48 << ((255) & 63)) {
                    case 127:
                        if true {
                            fmt.Println(v41)
                            if (v4 > (v4 / ((v4) | 1))) {
                                if ((v4 | (v4 | v4)) != (v4 ^ 1)) {
                                    if true {
                                        switch v3 {
                                        case "hello":
                                            if len(v8) > 0 {
                                                v8[0] = (fi48 ^ -914)
                                            }
                                        default:
                                            switch ((v3 + v3) + "\n\t") {
                                            case "\n\t":
                                                v4 -= 2
                                            }
                                        }
                                        v51 := GenS4{F0: v41, F1: 278}
                                        _ = v51
                                    } else {
                                        if true {
                                            switch ("\n\t" + v3) {
                                            case "foo bar":
                                                switch (v3 + v3) {
                                                case "a":
                                                    v3 = "hello"
                                                case "":
                                                    switch "hello" {
                                                    case "hello":
                                                        var v52 bool = false
                                                        _ = v52
                                                    case "foo bar":
                                                        switch fi47 {
                                                        case 1:
                                                            v8 = append(v8, 1)
                                                        }
                                                    default:
                                                        var v53 float32 = 1.5
                                                        _ = v53
                                                    }
                                                default:
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                }
                                                v41 -= uint8(v1)
                                                fmt.Println(v41)
                                            case "\n\t":
                                                v3 = (v3 + "foo bar")
                                            }
                                            if len(v8) > 0 {
                                                v8[0] = (fi47 & 9223372036854775807)
                                            }
                                            var v54 int16 = -211
                                            _ = v54
                                        } else {
                                            if false {
                                                var v55 float64 = (0.1 / 1e10)
                                                v1 += (^(v1))
                                                _ = v55
                                            }
                                        }
                                        v56 := GenS5{F0: 32767, F1: 42}
                                        _ = v56
                                    }
                                    var v57 bool = ((^(fi48)) <= fi48)
                                    v41 += (^(v41))
                                    var v58 uint16 = v1
                                    v4 = v4
                                    var v59 uint16 = v58
                                    v60 := GenS6{F0: false}
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                } else {
                                    v61 := GenS2{F0: uint(v4), F1: 962, F2: 72}
                                    if false {
                                        if len(v8) > 0 {
                                            v8[0] = (fi47 ^ 9223372036854775807)
                                        }
                                        switch fi47 {
                                        case 32768:
                                            var v62 int16 = 1
                                            _ = v62
                                        }
                                    } else {
                                        var v63 int32 = 0
                                        v3 = "a"
                                        _ = v63
                                    }
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    v2 -= ((v2 + ((v2 * 3.14159) / float32(1e10))) * -2.25)
                                    _ = len(v8)
                                    _ = v61
                                }
                                v64 := GenS7{F0: (^(v1)), F1: -2147483648, F2: 306}
                                _ = v64
                            } else {
                                v2 += v2
                                v2 = (1e10 - v2)
                            }
                            fmt.Println(v2)
                        } else {
                            v65 := GenS8{F0: 42, F1: 65535, F2: (v3 + "foo bar")}
                            v66 := GenS6{F0: true}
                            _ = v65
                            _ = v66
                        }
                    case -1:
                        fmt.Println(v3)
                    default:
                        fmt.Println(fi48)
                    }
                    var v67 uint32 = 282
                    _ = v67
                }
                _ = fi48
            }
            _ = fi47
        }
        switch v3 {
        case "foo bar":
            switch 9223372036854775807 {
            case -2147483648:
                for fi68 := 0; fi68 < 10; fi68++ {
                    if false {
                        v41 = v41
                        v2 = 3.14159
                    }
                    var v69 uint = 18446744073709551615
                    _ = fi68
                    _ = v69
                }
            case 32767:
                var v70 bool = true
                _ = v70
            case -9223372036854775808:
                var v71 int = 9223372036854775807
                _ = v71
            }
        default:
            v72 := GenS9{F0: 2, F1: 4294967295}
            _ = v72
        }
        v41 = v41
    }
    v1 = 1
    for fi73 := 0; fi73 < 19; fi73++ {
        switch v3 {
        case "a":
            var v74 int = 127
            _ = v74
        case "\n\t":
            v3 = ((v3 + "a") + v3)
        case "foo bar":
            fmt.Println(v1)
        }
        _ = fi73
    }
    var v75 bool = false
    fmt.Println(v2, v2, v2)
    fmt.Println(v1, v2, v3)
    var v76 bool = v75
    for fi77 := 0; fi77 < 10; fi77++ {
        for fi78 := 0; fi78 < 1; fi78++ {
            switch 590 {
            case 127:
                fmt.Println(v2, v3, fi77)
            case -9223372036854775808:
                _ = len(v8)
            default:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                v76 = (genFn0() >= 42)
                if len(v8) > 0 {
                    v8[0] = fi77
                }
            }
            var v79 int64 = 1
            _ = fi78
            _ = v79
        }
        var v80 uint64 = (^(v4))
        _ = fi77
        _ = v80
    }
    for fi81 := 0; fi81 < 12; fi81++ {
        v1 = uint16(v4)
        _ = fi81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v46
    _ = v75
    _ = v76
    _ = v8
}
