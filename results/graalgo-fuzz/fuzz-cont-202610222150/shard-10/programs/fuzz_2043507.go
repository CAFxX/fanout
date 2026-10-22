package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int32 }

func main() {
    v1 := []string{("hello" + "hello"), "a"}
    v2 := make(map[string]int)
    var v3 uint = 2
    var v4 string = ("hello" + "")
    v5 := []int{-9223372036854775808, 24, 128}
    v5 = append(v5, 32768)
    switch ("hello" + v4) {
    case "\n\t":
        var v6 float64 = 0.1
        _ = v6
    }
    if (!(v3 == v3)) {
        if (0 != 0) {
            if false {
                v7 := GenS1{F0: 42}
                if false {
                    var v8 uint = v3
                    fmt.Println(v4, v3, v3)
                    _ = v8
                }
                _ = v7
            }
            var v9 uint = v3
            v4 = (((v4 + ("foo bar" + ("hello" + v4))) + v4) + "a")
            _ = v9
        } else {
            if (34 < 1) {
                var v10 uint64 = 1
                if true {
                    fmt.Println(v3)
                    if true {
                        if (!true) {
                            var v11 int32 = -1
                            fmt.Println(v4, v4, v4)
                            fmt.Println(v4, v4)
                            _ = v11
                        }
                    } else {
                        fmt.Println(v4, v3, v10)
                        v10 -= 4294967295
                    }
                    v12 := GenS2{F0: 1, F1: -645}
                    _ = v12
                } else {
                    v2["foo bar"] = 32767
                    var v13 int32 = 0
                    _ = v13
                }
                if true {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        var v15 uint64 = v10
                        _ = fi14
                        _ = v15
                    }
                    if (((float32(-2.25) / 1.5) == (((1e10 / ((3.14159 - 0.0) * (1e10 - 0.0))) * ((1e10 / ((1e10 + 0.0) + 0.0)) + 3.14159)) * (1.5 / ((3.14159 - (((0.1 / 1.5) / 1e10) / (((float32(0.0) + 0.0) + -2.25) - (3.14159 * 1e10)))) / 1e10)))) && false) {
                        var v16 uint8 = 254
                        v10 -= v10
                        _ = v16
                    } else {
                        var v17 int = 1
                        if false {
                            v4 = v4
                            var v18 int32 = 1
                            fmt.Println(v4, v4)
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 13; fi19++ {
                            if false {
                                fmt.Println(v10, v17, v3)
                            } else {
                                var v20 uint64 = (v10 - 4294967295)
                                _ = v20
                            }
                            v10 -= v10
                            fmt.Println(v3)
                            _ = fi19
                        }
                        _ = v17
                    }
                } else {
                    v3 = v3
                    v10 -= (v10 / (((v10 * (v10 >> ((65535) & 63)))) | 1))
                    if (v10 != v10) {
                        v3 -= v3
                        for fi21 := 0; fi21 < 7; fi21++ {
                            switch fi21 {
                            case 32768:
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    var v23 uint64 = v10
                                    var v24 float32 = ((1.5 / (0.1 - 1e10)) + (((-2.25 + float32(0.1)) - 1.5) * (0.1 + 0.0)))
                                    var v25 string = v4
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                            case 42:
                                fmt.Println(fi21)
                            case -726:
                                for fi26 := 0; fi26 < 8; fi26++ {
                                    v10 = 42
                                    var v27 int = (fi26 ^ -127)
                                    _ = fi26
                                    _ = v27
                                }
                            default:
                                var v28 int8 = -118
                                _ = v28
                            }
                            switch (v4 + "") {
                            case "foo bar":
                                if false {
                                    var v29 uint32 = 65535
                                    _ = v29
                                } else {
                                    v30 := GenS3{F0: 817}
                                    v31 := GenS2{F0: 42, F1: -870}
                                    _ = v30
                                    _ = v31
                                }
                            case "a":
                                if (3.14159 == float64(0.1)) {
                                    fmt.Println(fi21, v10, v4)
                                    var v32 int8 = 0
                                    fmt.Println(fi21, fi21, v3)
                                    _ = v32
                                } else {
                                    delete(v2, v4)
                                    fmt.Println(fi21, v3)
                                }
                            }
                            _ = fi21
                        }
                        v4 = v4
                    } else {
                        fmt.Println(v4, v10, v10)
                    }
                }
                _ = v10
            }
        }
        if (65535 <= 1) {
            fmt.Println(v3, v3)
            v4 = "foo bar"
            switch 0 {
            case 32768:
                var v33 uint64 = 1
                _ = v33
            case 9223372036854775807:
                var v34 int = 128
                _ = v34
            case 51:
                for fi35 := 0; fi35 < 3; fi35++ {
                    var v36 int8 = 1
                    v3 -= v3
                    _ = fi35
                    _ = v36
                }
            }
        }
        for fi37 := 0; fi37 < 13; fi37++ {
            switch (fi37 * -2147483648) {
            case -797:
                v4 = "a"
            case 2147483647:
                if true {
                    var v38 int8 = 0
                    _ = v38
                }
            default:
                v4 = (v4 + v4)
            }
            v1 = append(v1, "")
            var v39 int8 = 0
            _ = fi37
            _ = v39
        }
    } else {
        fmt.Println(v3, v4)
        _ = len(v2)
    }
    v3 = 255
    fmt.Println(v3, v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (234 <= 9223372036854775807) {
        var v40 int64 = 42
        _ = v40
    } else {
        for fi41 := 0; fi41 < 20; fi41++ {
            if true {
                var v42 uint64 = 42
                _ = v42
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if (!(1 > 717)) {
                    var v43 string = v4
                    for fi44 := 0; fi44 < 7; fi44++ {
                        fmt.Println(fi41)
                        var v45 uint32 = 2
                        v46 := GenS4{F0: 0, F1: 9223372036854775807}
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                    _ = v43
                }
            }
            _ = fi41
        }
        var v47 uint16 = 961
        for fi48 := 0; fi48 < 8; fi48++ {
            for fi49 := 0; fi49 < 5; fi49++ {
                v4 = "\n\t"
                _ = fi49
            }
            _ = fi48
        }
        _ = v47
    }
    v2["foo bar"] = 67
    v3 = (v3 % (((v3 % ((681) | 1))) | 1))
    if len(v5) > 0 {
        v5[0] = -301
    }
    var v50 int = -1
    if (-64 >= -1) {
        v51 := GenS2{F0: 24, F1: 1}
        _ = v51
    } else {
        if (-2147483648 < 2147483647) {
            var v52 int16 = 1000
            v53 := GenS1{F0: 1}
            _ = v52
            _ = v53
        } else {
            if false {
                v54 := GenS2{F0: 2, F1: 1}
                var v55 int = v50
                _ = v54
                _ = v55
            } else {
                if (0 > 9223372036854775807) {
                    _ = len(v1)
                    for fi56 := 0; fi56 < 18; fi56++ {
                        switch (fi56 + fi56) {
                        case 719:
                            if (1e10 <= (1.5 + -2.25)) {
                                var v57 uint32 = 2147483648
                                _ = v57
                            }
                        case -1:
                            fmt.Println(fi56)
                        }
                        v58 := GenS4{F0: -1, F1: -1}
                        switch (v50 >> ((v3) & 63)) {
                        case 32767:
                            switch 0 {
                            case -402:
                                if ((true || false) || true) {
                                    var v59 int = 4
                                    v3 -= v3
                                    switch v4 {
                                    case "a":
                                        switch v59 {
                                        case -2147483648:
                                            if (-58 > 127) {
                                                var v60 int32 = 902
                                                fmt.Println(v3, v4)
                                                _ = v60
                                            } else {
                                                switch v4 {
                                                case "":
                                                    if (!true) {
                                                        var v61 uint16 = 255
                                                        if true {
                                                            var v62 float64 = 0.1
                                                            v61 -= v61
                                                            _ = v62
                                                        } else {
                                                            fmt.Println(fi56, v59)
                                                            var v63 float32 = 0.1
                                                            var v64 uint32 = 758
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                        v1 = append(v1, v4)
                                                        _ = v61
                                                    } else {
                                                        for fi65 := 0; fi65 < 9; fi65++ {
                                                            _ = len(v1)
                                                            v3 -= v3
                                                            _ = fi65
                                                        }
                                                        v59 -= fi56
                                                        for fi66 := 0; fi66 < 9; fi66++ {
                                                            switch fi66 {
                                                            case 9223372036854775807:
                                                                var v67 string = ""
                                                                _ = v67
                                                            case 127:
                                                                var v68 bool = true
                                                                _ = v68
                                                            default:
                                                                var v69 uint32 = 0
                                                                _ = v69
                                                            }
                                                            var v70 uint64 = 4294967295
                                                            switch v59 {
                                                            case 9223372036854775807:
                                                                v3 = 555
                                                            }
                                                            _ = fi66
                                                            _ = v70
                                                        }
                                                    }
                                                case "a":
                                                    _ = len(v1)
                                                case "hello":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            }
                                        case 9223372036854775807:
                                            v59 = (v50 | (v50 >> ((v3) & 63)))
                                        case 4294967295:
                                            for fi71 := 0; fi71 < 9; fi71++ {
                                                if false {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v72 float32 = -2.25
                                                    _ = len(v2)
                                                    _ = v72
                                                }
                                                v73 := GenS5{F0: 548}
                                                v74 := GenS2{F0: 255, F1: 1000}
                                                var v75 int = fi71
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v76 uint = 0
                                                    _ = v76
                                                case "hello":
                                                    fmt.Println(v3, v50, v3)
                                                }
                                                delete(v2, (v4 + v4))
                                                _ = fi71
                                                _ = v73
                                                _ = v74
                                                _ = v75
                                            }
                                        }
                                    case "\n\t":
                                        for fi77 := 0; fi77 < 8; fi77++ {
                                            var v78 int16 = 42
                                            if (716 == 26) {
                                                var v79 bool = (65 <= 164)
                                                v80 := GenS4{F0: -9223372036854775808, F1: 0}
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                _ = v79
                                                _ = v80
                                            }
                                            _ = fi77
                                            _ = v78
                                        }
                                    case "foo bar":
                                        v59 = fi56
                                    }
                                    _ = v59
                                } else {
                                    if true {
                                        v3 += (v3 ^ (v3 ^ 4294967295))
                                    }
                                    var v81 bool = true
                                    _ = v81
                                }
                            case 2147483647:
                                var v82 float64 = (float64(-2.25) + 3.14159)
                                _ = v82
                            case 127:
                                var v83 bool = (2147483647 != 42)
                                _ = v83
                            default:
                                if (!false) {
                                    var v84 int64 = -10
                                    _ = v84
                                }
                            }
                        }
                        _ = fi56
                        _ = v58
                    }
                    var v85 bool = true
                    _ = v85
                }
                v3 += 255
                v3 = (v3 << ((v3) & 63))
            }
            v86 := GenS4{F0: 1, F1: 9223372036854775807}
            _ = v86
        }
        var v87 bool = true
        _ = v87
    }
    v88 := v2[("" + "foo bar")]
    var v89 bool = ((2 <= 4294967295) && false)
    var v90 int = v50
    switch 1 {
    case 127:
        if v89 {
            v5 = append(v5, (v50 / ((v90) | 1)))
        } else {
            v50 = v50
            var v91 uint16 = 42
            _ = v91
        }
    }
    switch (-(v50)) {
    case 9223372036854775807:
        v4 = v4
    case 127:
        v92 := GenS1{F0: 4294967295}
        _ = v92
    }
    switch "a" {
    case "foo bar":
        if false {
            delete(v2, v4)
        }
    case "hello":
        v89 = false
    case "\n\t":
        var v93 bool = (-366 != -32768)
        _ = v93
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v88
    _ = v89
    _ = v90
}
