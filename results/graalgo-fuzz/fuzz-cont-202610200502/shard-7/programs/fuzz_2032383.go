package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint8; F1 string; F2 int32 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS12 struct { F0 int64; F1 uint64 }
type GenS13 struct { F0 int64; F1 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{}
    var v3 float32 = (3.14159 / float32(float64(((0.1 / 3.14159) + 1.5))))
    var v4 uint64 = 903
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 string = ("foo bar" + "hello")
        for fi7 := 0; fi7 < 10; fi7++ {
            switch v6 {
            case "foo bar":
                v8 := GenS1{F0: -197, F1: 804, F2: false}
                _ = v8
            case "\n\t":
                switch v6 {
                case "a":
                    v2 = append(v2, 65535)
                default:
                    fmt.Println(v4, fi5)
                }
            case "a":
                var v9 float32 = v3
                _ = v9
            }
            fmt.Println(fi7, v4, v4)
            if len(v2) > 0 {
                v2[0] = 789
            }
            if true {
                fmt.Println(fi7, v6, fi7)
                switch fi7 {
                case -1:
                    if false {
                        switch "foo bar" {
                        case "foo bar":
                            var v10 string = v6
                            _ = v10
                        case "a":
                            if (false || false) {
                                var v11 uint8 = 255
                                fmt.Println(v6, v4, fi7)
                                _ = v11
                            }
                        case "\n\t":
                            fmt.Println(fi5)
                        }
                        var v12 uint = 4294967295
                        switch fi5 {
                        case -838:
                            var v13 uint8 = 127
                            _ = v13
                        case 32767:
                            var v14 bool = (false || false)
                            _ = v14
                        case 528:
                            var v15 uint = v12
                            _ = v15
                        default:
                            v2 = append(v2, 809)
                        }
                        var v16 float32 = ((((v3 + v3) - -2.25) - v3) * 0.1)
                        _ = v12
                        _ = v16
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v17 uint = 721
                        _ = v17
                    }
                case -853:
                    if (!(true || (!(false || true)))) {
                        v18 := GenS1{F0: -1, F1: 65535, F2: false}
                        if true {
                            v19 := v1[v6]
                            _ = v19
                        } else {
                            v20 := v1["a"]
                            v21 := GenS2{F0: 42}
                            _ = v20
                            _ = v21
                        }
                        _ = v18
                    } else {
                        v6 = "foo bar"
                    }
                default:
                    v22 := GenS1{F0: -9223372036854775808, F1: 292, F2: true}
                    _ = v22
                }
            } else {
                if (-1 > 97) {
                    v23 := GenS2{F0: -54}
                    _ = v23
                }
            }
            _ = fi7
        }
        v4 += v4
        _ = fi5
        _ = v6
    }
    if (uint16(v4) <= 255) {
        v1[((("a" + "\n\t") + ("foo bar" + "a")) + "foo bar")] = false
        for fi24 := 0; fi24 < 8; fi24++ {
            if false {
                var v25 uint8 = 42
                if (!true) {
                    v3 -= (v3 / 1.0)
                    if false {
                        v3 += v3
                        v1[("a" + "foo bar")] = (0 == -608)
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "foo bar":
                            if (!false) {
                                var v26 float64 = -2.25
                                _ = len(v1)
                                _ = v26
                            } else {
                                v3 = (((v3 + v3) / 3.14159) + v3)
                                fmt.Println(v25, fi24)
                                if (true || (!(!false))) {
                                    for fi27 := 0; fi27 < 17; fi27++ {
                                        switch ("a" + "hello") {
                                        case "a":
                                            var v28 int16 = -1
                                            fmt.Println(v3, v3)
                                            fmt.Println(v28)
                                            _ = v28
                                        case "\n\t":
                                            var v29 uint8 = (v25 ^ 68)
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(fi27, v25, v4)
                                            default:
                                                switch (fi27 >> ((0) & 63)) {
                                                case 4294967295:
                                                    switch 0 {
                                                    case -2147483648:
                                                        v1["foo bar"] = ((2 != uint16(v25)) && (0 > 18446744073709551615))
                                                    }
                                                case 1:
                                                    if true {
                                                        v29 = uint8(v4)
                                                    } else {
                                                        if (true && false) {
                                                            v3 -= (float32(0.1) / (v3 * (((v3 / float32(3.14159)) / (float32(-2.25) / v3)) - -2.25)))
                                                            if (0 < -2147483648) {
                                                                var v30 float64 = -2.25
                                                                _ = v30
                                                            }
                                                            if false {
                                                                switch -354 {
                                                                case 9223372036854775807:
                                                                    var v31 string = "foo bar"
                                                                    _ = v31
                                                                case 4294967295:
                                                                    v3 = 1.5
                                                                default:
                                                                    var v32 int32 = 2147483647
                                                                    _ = v32
                                                                }
                                                                v33 := GenS2{F0: -655}
                                                                var v34 uint16 = 255
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                            switch 128 {
                                                            case -693:
                                                                v35 := GenS1{F0: -9223372036854775808, F1: 0, F2: (42 < 925)}
                                                                _ = v35
                                                            case 127:
                                                                v36 := GenS3{F0: -797, F1: 42}
                                                                _ = v36
                                                            default:
                                                                v25 = (v29 % (((v25 << ((uint(v4)) & 63))) | 1))
                                                            }
                                                        }
                                                        var v37 int = 9223372036854775807
                                                        _ = v37
                                                    }
                                                }
                                            }
                                            _ = v29
                                        case "":
                                            switch "hello" {
                                            case "\n\t":
                                                if false {
                                                    v38 := GenS4{F0: 885, F1: -1}
                                                    if false {
                                                        if true {
                                                            var v39 int32 = 42
                                                            v40 := GenS5{F0: 91, F1: "\n\t", F2: (v39 / (((-(v39))) | 1))}
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                    } else {
                                                        var v41 int = fi27
                                                        _ = v41
                                                    }
                                                    _ = v38
                                                }
                                            case "":
                                                var v42 string = "\n\t"
                                                _ = v42
                                            case "hello":
                                                if ((v4 <= 1) || true) {
                                                    if ((("" <= "foo bar") && false) || false) {
                                                        if ((9223372036854775807 >= fi27) || false) {
                                                            var v43 uint64 = v4
                                                            _ = v43
                                                        }
                                                        v44 := GenS6{F0: -2147483648, F1: (fi24 << ((255) & 63))}
                                                        if (14 == 255) {
                                                            switch (fi27 * fi27) {
                                                            case 32767:
                                                                v1[("hello" + "")] = (!false)
                                                            case 32768:
                                                                delete(v1, ("foo bar" + (("foo bar" + "foo bar") + "foo bar")))
                                                            }
                                                            v45 := GenS6{F0: 2147483647, F1: (fi27 % ((0) | 1))}
                                                            _ = v45
                                                        } else {
                                                            fmt.Println(fi24, fi24, v3)
                                                        }
                                                        fmt.Println(fi24, v25, fi27)
                                                        var v46 string = ("hello" + "foo bar")
                                                        _ = v44
                                                        _ = v46
                                                    }
                                                    var v47 int64 = 0
                                                    _ = v47
                                                } else {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v48 int32 = -216
                                                        _ = v48
                                                    case "":
                                                        if true {
                                                            var v49 float64 = 0.1
                                                            var v50 string = ("hello" + "a")
                                                            _ = len(v2)
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            _ = len(v2)
                                                            if true {
                                                                if (false && false) {
                                                                    v51 := GenS7{F0: 18446744073709551615, F1: (v4 % ((v4) | 1))}
                                                                    v4 = v4
                                                                    var v52 int = fi27
                                                                    v1["\n\t"] = false
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                                v4 = (v4 << ((0) & 63))
                                                            }
                                                            fmt.Println(fi24)
                                                        }
                                                    case "\n\t":
                                                        switch (fi27 | (-(fi27))) {
                                                        case -758:
                                                            if (float32(0.0) < ((-2.25 / 1.5) / (((v3 * 1e10) - 0.0) + (0.1 + (0.1 / float32(-2.25)))))) {
                                                                var v53 uint = 728
                                                                switch fi24 {
                                                                case -61:
                                                                    v54 := GenS7{F0: (v53 ^ (^(v53))), F1: 0}
                                                                    _ = v54
                                                                case 42:
                                                                    switch 533 {
                                                                    case 32767:
                                                                        v25 = v25
                                                                    case 127:
                                                                        if (58 > -26) {
                                                                            v3 = v3
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                v3 = 0.1
                                                                            case "":
                                                                                v55 := GenS8{F0: "\n\t"}
                                                                                _ = v55
                                                                            case "hello":
                                                                                var v56 string = ("\n\t" + "foo bar")
                                                                                _ = v56
                                                                            default:
                                                                                var v57 float32 = v3
                                                                                _ = v57
                                                                            }
                                                                            var v58 uint16 = 2
                                                                            _ = v58
                                                                        } else {
                                                                            var v59 int64 = 1
                                                                            if ((!((v3 * v3) == v3)) || false) {
                                                                                var v60 int8 = 0
                                                                                var v61 int8 = v60
                                                                                if (32767 > 42) {
                                                                                    v62 := GenS8{F0: "hello"}
                                                                                    if ((1e10 + v3) != (3.14159 - float32(1.5))) {
                                                                                        if (!false) {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        }
                                                                                        v61 += -1
                                                                                    }
                                                                                    _ = v62
                                                                                }
                                                                                fmt.Println(v59, v25, fi27)
                                                                                v61 = 0
                                                                                v63 := GenS3{F0: -1, F1: -794}
                                                                                _ = v60
                                                                                _ = v61
                                                                                _ = v63
                                                                            }
                                                                            _ = len(v2)
                                                                            var v64 bool = true
                                                                            _ = v59
                                                                            _ = v64
                                                                        }
                                                                    case 42:
                                                                        var v65 float64 = (1e10 / 1.0)
                                                                        _ = v65
                                                                    }
                                                                case 127:
                                                                    v66 := GenS9{F0: (v25 | v25)}
                                                                    _ = v66
                                                                default:
                                                                    v67 := GenS5{F0: 147, F1: "foo bar", F2: 329}
                                                                    _ = v67
                                                                }
                                                                v25 -= v25
                                                                delete(v1, ("" + "a"))
                                                                _ = v53
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        if false {
                                            if (false && true) {
                                                v3 -= 0.0
                                                var v68 int16 = 367
                                                v69 := GenS9{F0: v25}
                                                v68 += (v68 / ((32767) | 1))
                                                _ = v68
                                                _ = v69
                                            } else {
                                                var v70 uint64 = (v4 / ((v4) | 1))
                                                v2 = append(v2, 0)
                                                _ = v70
                                            }
                                            v4 -= v4
                                            v71 := GenS10{F0: 32767}
                                            _ = v71
                                        } else {
                                            v3 = 1.5
                                        }
                                        _ = fi27
                                    }
                                    var v72 uint32 = 1
                                    var v73 uint = 18446744073709551615
                                    _ = v72
                                    _ = v73
                                } else {
                                    fmt.Println(v4, fi24)
                                }
                            }
                        case "a":
                            v74 := GenS11{F0: v4, F1: 1, F2: (v25 * (v25 + v25))}
                            _ = v74
                        }
                        var v75 uint16 = 2
                        v75 = (v75 * (v75 & 32767))
                        _ = v75
                    }
                }
                _ = v25
            }
            v4 += v4
            _ = fi24
        }
    }
    fmt.Println(v4)
    switch 837 {
    case 79:
        if false {
            if (0 <= 551) {
                v76 := GenS12{F0: 993, F1: (^(v4))}
                v4 = v4
                _ = v76
            } else {
                var v77 uint16 = 255
                var v78 int8 = -128
                _ = v77
                _ = v78
            }
            fmt.Println(v4, v4)
            var v79 float32 = ((float32((1e10 * -2.25)) / ((((v3 * float32((0.0 + float64((0.0 + (v3 - float32(float64((v3 - v3))))))))) * (v3 + v3)) - v3) / v3)) + (1e10 + (v3 + v3)))
            _ = v79
        }
    case 42:
        if (!true) {
            v3 = 1e10
        }
    case -1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v80 bool = false
    switch "foo bar" {
    case "\n\t":
        switch -9223372036854775808 {
        case 32767:
            fmt.Println(v80, v3)
        default:
            v4 = v4
        }
    case "hello":
        if true {
            if v80 {
                switch -1 {
                case 127:
                    var v81 float64 = -2.25
                    _ = v81
                case 42:
                    if ((v3 / 1e10) <= ((float32((((1e10 + 1e10) - 1.5) * 1e10)) * float32((1e10 - 0.0))) * (1e10 / v3))) {
                        var v82 int16 = 649
                        v83 := GenS5{F0: 2, F1: "\n\t", F2: -2147483648}
                        _ = v82
                        _ = v83
                    }
                case 782:
                    if true {
                        var v84 int64 = int64(v4)
                        _ = v84
                    }
                }
                v80 = v80
                v2 = append(v2, 1)
            } else {
                fmt.Println(v3, v3, v80)
                for fi85 := 0; fi85 < 19; fi85++ {
                    fmt.Println(v4, v3, v3)
                    _ = fi85
                }
                switch -1 {
                case 42:
                    for fi86 := 0; fi86 < 1; fi86++ {
                        var v87 uint16 = 0
                        _ = fi86
                        _ = v87
                    }
                case 4294967295:
                    v80 = ((true && (-1 <= 32767)) || v80)
                case -9223372036854775808:
                    if (!v80) {
                        switch 413 {
                        case 0:
                            v80 = (v80 || (!v80))
                        case 32768:
                            v3 = float32(0.0)
                        default:
                            v88 := GenS7{F0: 4294967295, F1: v4}
                            _ = v88
                        }
                        if (!(265 >= 321)) {
                            for fi89 := 0; fi89 < 5; fi89++ {
                                _ = len(v2)
                                _ = fi89
                            }
                        }
                        for fi90 := 0; fi90 < 18; fi90++ {
                            fmt.Println(v4)
                            v80 = ((fi90 - 42) < fi90)
                            _ = fi90
                        }
                    } else {
                        var v91 float64 = 0.1
                        _ = v91
                    }
                }
            }
            fmt.Println(v4, v3, v4)
        }
    default:
        for fi92 := 0; fi92 < 2; fi92++ {
            v3 -= 0.1
            var v93 uint16 = 814
            _ = fi92
            _ = v93
        }
    }
    v94 := GenS13{F0: -577, F1: 0}
    var v95 float32 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v80:", v80)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v80
    _ = v94
    _ = v95
}
