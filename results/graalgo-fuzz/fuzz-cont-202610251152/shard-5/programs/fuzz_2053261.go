package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint8; F2 string }
type GenS5 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 bool; F1 int; F2 bool }
type GenS10 struct { F0 int; F1 uint16 }

func main() {
    v1 := []string{"a"}
    v2 := []int{}
    v3 := []int{-765, 127, 128}
    var v4 bool = (false || (("" <= "\n\t") || false))
    v5 := []string{("a" + "\n\t"), ""}
    v3 = append(v3, -9223372036854775808)
    v6 := GenS1{F0: 0, F1: 360, F2: -1}
    _ = len(v1)
    v4 = v4
    if false {
        fmt.Println(v4)
        v3 = append(v3, -9223372036854775808)
    } else {
        var v7 int8 = -30
        v8 := GenS1{F0: -1, F1: 694, F2: -574}
        if false {
            v7 = -128
            if (2 > 255) {
                fmt.Println(v7, v4)
                var v9 int64 = 42
                var v10 uint32 = 0
                _ = v10
                _ = v9
            } else {
                v11 := GenS2{F0: -136, F1: 2, F2: ("" + "foo bar")}
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 uint8 = 109
                    var v14 int16 = 0
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                switch 0 {
                case 1:
                    var v15 int32 = -1
                    _ = v15
                }
                _ = v11
            }
        } else {
            v16 := GenS2{F0: 0, F1: 2147483648, F2: (("a" + "hello") + "foo bar")}
            _ = v16
        }
        _ = v7
        _ = v8
    }
    v4 = true
    if false {
        var v17 int = 658
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(v17)
            _ = fi18
        }
        _ = v17
    } else {
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 uint64 = 549
            if ((0.0 - -2.25) > (0.0 / (float32((float64(1.5) / (-2.25 / 1.0))) * 0.1))) {
                var v21 int8 = 106
                _ = v21
            }
            _ = fi19
            _ = v20
        }
    }
    var v22 uint64 = 248
    v4 = (!v4)
    v22 -= v22
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 uint16 = 1
        v5 = append(v5, "foo bar")
        _ = fi23
        _ = v24
    }
    if true {
        v25 := GenS1{F0: -776, F1: 4294967295, F2: 1000}
        _ = v25
    } else {
        if false {
            v4 = v4
            if true {
                switch "hello" {
                case "foo bar":
                    fmt.Println(v4, v22)
                case "\n\t":
                    var v26 int8 = 1
                    _ = v26
                default:
                    var v27 int32 = -1
                    _ = v27
                }
                v28 := GenS3{F0: v22}
                _ = v28
            } else {
                var v29 float64 = (1e10 / 1e10)
                var v30 int32 = -2147483648
                v4 = (42 > 42)
                _ = v29
                _ = v30
            }
            if ((((1.5 - ((1.5 - 1e10) * 1.5)) - (3.14159 - 1e10)) > ((0.1 + float32(0.1)) / 1e10)) && (!true)) {
                v22 = 1
            }
        } else {
            if (v4 || v4) {
                if false {
                    v4 = true
                    if (4294967295 > 2147483648) {
                        if (0 == 4294967295) {
                            v5 = append(v5, "\n\t")
                            fmt.Println(v22, v4)
                            var v31 uint8 = 127
                            _ = v31
                        } else {
                            var v32 int8 = 86
                            v1 = append(v1, "")
                            _ = v32
                        }
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            var v33 uint64 = 4294967295
                            _ = v33
                        case "a":
                            if len(v3) > 0 {
                                v3[0] = -792
                            }
                        case "":
                            if false {
                                v4 = v4
                                v34 := GenS1{F0: 128, F1: 421, F2: 0}
                                v35 := GenS4{F0: v22, F1: 128, F2: "foo bar"}
                                _ = v34
                                _ = v35
                            }
                        default:
                            v22 = (v22 & v22)
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch ("foo bar" + "a") {
                        case "foo bar":
                            if v4 {
                                v4 = v4
                                for fi36 := 0; fi36 < 15; fi36++ {
                                    for fi37 := 0; fi37 < 18; fi37++ {
                                        switch ("\n\t" + ("hello" + "foo bar")) {
                                        case "a":
                                            if v4 {
                                                v4 = (1 > 491)
                                            } else {
                                                if v4 {
                                                    var v38 float64 = 0.0
                                                    v39 := GenS5{F0: 42, F1: 808, F2: 1}
                                                    _ = v38
                                                    _ = v39
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    v3 = append(v3, fi36)
                                                case "":
                                                    v40 := GenS3{F0: 816}
                                                    _ = v40
                                                case "hello":
                                                    switch ("a" + ("hello" + "foo bar")) {
                                                    case "\n\t":
                                                        var v41 float32 = ((((-2.25 * 1e10) - -2.25) + 1.5) * float32((3.14159 - 3.14159)))
                                                        _ = v41
                                                    case "hello":
                                                        v42 := GenS4{F0: 1, F1: 127, F2: ""}
                                                        _ = v42
                                                    }
                                                }
                                            }
                                        }
                                        switch "hello" {
                                        case "":
                                            v43 := GenS6{F0: 65535, F1: 713}
                                            _ = v43
                                        default:
                                            if v4 {
                                                v44 := GenS5{F0: 2, F1: 1, F2: 2}
                                                fmt.Println(v4)
                                                fmt.Println(fi37)
                                                _ = v44
                                            } else {
                                                fmt.Println(v22)
                                            }
                                        }
                                        var v45 int64 = 208
                                        _ = fi37
                                        _ = v45
                                    }
                                    _ = fi36
                                }
                            } else {
                                switch "" {
                                case "a":
                                    for fi46 := 0; fi46 < 3; fi46++ {
                                        v22 = 4294967295
                                        if v4 {
                                            var v47 uint16 = 495
                                            _ = v47
                                        } else {
                                            for fi48 := 0; fi48 < 20; fi48++ {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v1 = append(v1, "")
                                                _ = len(v2)
                                                var v49 uint64 = (v22 - v22)
                                                var v50 int16 = 1
                                                var v51 float64 = float64(float32(3.14159))
                                                _ = fi48
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                            }
                                            v4 = (!v4)
                                            var v52 uint = 18446744073709551615
                                            _ = v52
                                        }
                                        var v53 int16 = 42
                                        _ = fi46
                                        _ = v53
                                    }
                                case "":
                                    var v54 int16 = 32767
                                    _ = v54
                                default:
                                    switch 0 {
                                    case -1:
                                        var v55 uint32 = 0
                                        _ = v55
                                    case -2147483648:
                                        if v4 {
                                            var v56 uint64 = (v22 << ((42) & 63))
                                            v57 := GenS2{F0: 2147483647, F1: 2, F2: ("" + "foo bar")}
                                            _ = v56
                                            _ = v57
                                        } else {
                                            var v58 uint64 = 2
                                            v22 = (^(v22))
                                            _ = v58
                                        }
                                    }
                                }
                            }
                        case "":
                            for fi59 := 0; fi59 < 9; fi59++ {
                                for fi60 := 0; fi60 < 7; fi60++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi60
                                }
                                if (65535 != 951) {
                                    fmt.Println(v4)
                                } else {
                                    if v4 {
                                        var v61 int = (fi59 % ((fi59) | 1))
                                        var v62 int32 = 2147483647
                                        _ = v61
                                        _ = v62
                                    } else {
                                        var v63 uint16 = 26
                                        switch (fi59 - fi59) {
                                        case 128:
                                            v63 = (v63 - v63)
                                        }
                                        v63 += v63
                                        _ = v63
                                    }
                                    if (v4 || true) {
                                        fmt.Println(v22)
                                        if v4 {
                                            for fi64 := 0; fi64 < 14; fi64++ {
                                                var v65 float64 = ((1.5 / 1e10) * 3.14159)
                                                var v66 uint32 = 2147483648
                                                if (-1 < (^(fi59))) {
                                                    v3 = append(v3, -94)
                                                    v66 += (v66 % ((v66) | 1))
                                                    fmt.Println(v4)
                                                    switch "" {
                                                    case "hello":
                                                        if len(v5) > 0 {
                                                            v5[0] = (("" + "hello") + "a")
                                                        }
                                                        v67 := GenS6{F0: 899, F1: 42}
                                                        v68 := GenS5{F0: 1, F1: 4294967295, F2: 310}
                                                        _ = v67
                                                        _ = v68
                                                    case "foo bar":
                                                        var v69 int8 = -1
                                                        _ = v69
                                                    }
                                                    if false {
                                                        v70 := GenS7{F0: 0}
                                                        _ = v70
                                                    }
                                                    var v71 bool = (false && (587 < -1))
                                                    if v71 {
                                                        var v72 bool = (!(!(!v71)))
                                                        if true {
                                                            var v73 int = 127
                                                            _ = v73
                                                        } else {
                                                            if len(v1) > 0 {
                                                                v1[0] = "hello"
                                                            }
                                                        }
                                                        var v74 float64 = (3.14159 / (3.14159 / 1.0))
                                                        v65 = v65
                                                        _ = v72
                                                        _ = v74
                                                    }
                                                    _ = v71
                                                } else {
                                                    var v75 uint16 = 466
                                                    if (v22 >= v22) {
                                                        var v76 int16 = 42
                                                        var v77 int = (fi64 % ((fi59) | 1))
                                                        v78 := GenS8{F0: -589, F1: (-(v76))}
                                                        v79 := GenS9{F0: ((^(v75)) != v75), F1: (fi64 / ((538) | 1)), F2: v4}
                                                        _ = v76
                                                        _ = v77
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                    _ = v75
                                                }
                                                _ = fi64
                                                _ = v65
                                                _ = v66
                                            }
                                            switch (^(fi59)) {
                                            case 42:
                                                v80 := GenS10{F0: 804, F1: 2}
                                                _ = v80
                                            case 127:
                                                for fi81 := 0; fi81 < 7; fi81++ {
                                                    fmt.Println(fi59, fi81)
                                                    _ = fi81
                                                }
                                            default:
                                                for fi82 := 0; fi82 < 4; fi82++ {
                                                    v4 = v4
                                                    if (v4 && (40 >= 34)) {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if v4 {
                                                            fmt.Println(fi82, v22, v22)
                                                        } else {
                                                            var v83 int16 = 1
                                                            if (767 < 42) {
                                                                var v84 uint32 = 4294967295
                                                                v85 := GenS5{F0: 255, F1: 18446744073709551615, F2: 65535}
                                                                fmt.Println(v4)
                                                                _ = v84
                                                                _ = v85
                                                            }
                                                            var v86 float64 = (3.14159 / 1.5)
                                                            _ = v83
                                                            _ = v86
                                                        }
                                                    }
                                                    v87 := GenS4{F0: v22, F1: 0, F2: "foo bar"}
                                                    _ = fi82
                                                    _ = v87
                                                }
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = -430
                                            }
                                        }
                                        var v88 string = ("a" + (("a" + ("\n\t" + ("foo bar" + "a"))) + "foo bar"))
                                        _ = v88
                                    }
                                    fmt.Println(v4, v22)
                                }
                                fmt.Println(fi59, v4)
                                _ = fi59
                            }
                        case "a":
                            var v89 uint = 452
                            _ = v89
                        }
                    }
                } else {
                    v90 := GenS5{F0: 803, F1: 255, F2: 2}
                    _ = v90
                }
            }
        }
        for fi91 := 0; fi91 < 11; fi91++ {
            v4 = v4
            switch 32767 {
            case -674:
                if false {
                    var v92 int32 = 624
                    _ = v92
                }
            case 0:
                var v93 int = fi91
                _ = v93
            default:
                if (v4 && (!((v4 && v4) && (!false)))) {
                    var v94 int16 = -32768
                    v94 = 32767
                    fmt.Println(v4)
                    _ = v94
                }
            }
            fmt.Println(v4, v4)
            _ = fi91
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
