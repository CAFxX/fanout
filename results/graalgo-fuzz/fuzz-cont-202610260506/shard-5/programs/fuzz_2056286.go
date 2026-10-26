package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16; F1 int; F2 int }
type GenS7 struct { F0 int64; F1 int8; F2 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64; F1 string; F2 string }

func main() {
    v1 := []int{-207}
    v2 := []uint32{}
    var v3 bool = true
    v4 := make(map[string]string)
    fmt.Println(v3)
    v5 := v4["\n\t"]
    v3 = true
    v1 = append(v1, -322)
    v3 = false
    var v6 int16 = -32768
    if v3 {
        v7 := GenS1{F0: 50, F1: 2147483648, F2: 0}
        _ = v7
    } else {
        if (-11 > -21) {
            var v8 float64 = (1e10 - 1.5)
            v9 := GenS2{F0: 977}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v8
            _ = v9
        } else {
            var v10 float32 = 1e10
            var v11 int = 2147483647
            _ = v10
            _ = v11
        }
        v12 := v4[v5]
        _ = v12
    }
    if (false && (!("\n\t" < ("foo bar" + v5)))) {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 bool = v3
            var v15 bool = v3
            _ = fi13
            _ = v14
            _ = v15
        }
        switch v5 {
        case "\n\t":
            for fi16 := 0; fi16 < 1; fi16++ {
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 int64 = 861
                    if false {
                        if v3 {
                            fmt.Println(fi16)
                            var v19 int16 = 499
                            var v20 uint64 = 2
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(v6, v3)
                        }
                    } else {
                        if (93 < -313) {
                            if false {
                                v3 = true
                                fmt.Println(fi16, fi17)
                                v21 := GenS3{F0: v3, F1: (v6 + (v6 + (v6 - (v6 - v6))))}
                                _ = v21
                            } else {
                                var v22 int64 = (v18 / (((v18 % ((v18) | 1))) | 1))
                                _ = v22
                            }
                            v23 := GenS1{F0: 2147483647, F1: 1, F2: 1}
                            var v24 uint8 = 246
                            _ = v23
                            _ = v24
                        } else {
                            v4[((v5 + (("hello" + "a") + "hello")) + "a")] = v5
                            if true {
                                v4[""] = v5
                            } else {
                                var v25 int16 = -787
                                _ = v25
                            }
                        }
                        fmt.Println(v18)
                        switch ("foo bar" + v5) {
                        case "a":
                            v5 = v5
                        default:
                            v18 = 588
                        }
                    }
                    v26 := GenS4{F0: (v6 * -1), F1: 1}
                    _ = fi17
                    _ = v18
                    _ = v26
                }
                v27 := GenS5{F0: 261}
                fmt.Println(v3, v3)
                _ = fi16
                _ = v27
            }
        case "a":
            v28 := GenS6{F0: 0, F1: -321, F2: -1}
            _ = v28
        case "hello":
            if (0.1 == 3.14159) {
                if (v5 >= v5) {
                    var v29 uint32 = 394
                    _ = v29
                }
                if (255 >= 850) {
                    var v30 uint16 = 0
                    if v3 {
                        v31 := GenS6{F0: v6, F1: -684, F2: 1}
                        var v32 int16 = (v6 * (v6 << ((42) & 63)))
                        v3 = false
                        _ = v31
                        _ = v32
                    }
                    if v3 {
                        for fi33 := 0; fi33 < 4; fi33++ {
                            if v3 {
                                v5 = "a"
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    var v35 int8 = -37
                                    if (v3 && (0 >= 18446744073709551615)) {
                                        v36 := GenS1{F0: 0, F1: 2147483648, F2: v30}
                                        if (-2.25 < -2.25) {
                                            v5 = v5
                                        } else {
                                            fmt.Println(fi34, v5)
                                        }
                                        if v3 {
                                            var v37 uint = 1
                                            fmt.Println(v6)
                                            switch -2147483648 {
                                            case 0:
                                                v38 := GenS1{F0: 2147483647, F1: 827, F2: v30}
                                                _ = v38
                                            case -9223372036854775808:
                                                if (fi33 == fi34) {
                                                    v2 = append(v2, 2147483647)
                                                } else {
                                                    v37 = 1
                                                    v37 = (v37 * (v37 >> (((^(v37))) & 63)))
                                                }
                                                v35 = (v35 << (((v37 & v37)) & 63))
                                            case 42:
                                                v30 = 42
                                            }
                                            _ = v37
                                        } else {
                                            var v39 string = v5
                                            _ = v39
                                        }
                                        delete(v4, "foo bar")
                                        _ = v36
                                    } else {
                                        if (v3 || v3) {
                                            if (42 == 1) {
                                                switch 583 {
                                                case 1:
                                                    fmt.Println(v30, v6)
                                                case -2147483648:
                                                    switch (fi33 / (((-(fi34))) | 1)) {
                                                    case 4294967295:
                                                        v40 := GenS6{F0: v6, F1: (fi34 ^ (fi33 / ((fi33) | 1))), F2: 32768}
                                                        _ = v40
                                                    }
                                                default:
                                                    var v41 int64 = -9223372036854775808
                                                    _ = v41
                                                }
                                                switch v5 {
                                                case "\n\t":
                                                    v6 = 32767
                                                }
                                                var v42 uint = 876
                                                _ = v42
                                            }
                                            v3 = false
                                            switch ("" + v5) {
                                            case "\n\t":
                                                v30 -= 1
                                            case "":
                                                var v43 uint32 = 1
                                                _ = v43
                                            case "a":
                                                fmt.Println(v35, v6, v6)
                                            default:
                                                v44 := GenS7{F0: 1, F1: v35, F2: (v5 + "foo bar")}
                                                _ = v44
                                            }
                                        } else {
                                            v45 := GenS5{F0: -343}
                                            _ = v45
                                        }
                                        switch "a" {
                                        case "a":
                                            v3 = true
                                        case "":
                                            var v46 uint16 = v30
                                            _ = v46
                                        }
                                        if v3 {
                                            var v47 int8 = (v35 & (^(v35)))
                                            _ = v47
                                        }
                                        var v48 int32 = 0
                                        v5 = "foo bar"
                                        v30 += v30
                                        _ = v48
                                    }
                                    _ = fi34
                                    _ = v35
                                }
                            }
                            fmt.Println(v30)
                            switch (fi33 >> ((255) & 63)) {
                            case 127:
                                if (v3 || (!(v30 <= v30))) {
                                    if (51 != (fi33 ^ 42)) {
                                        fmt.Println(v3, v6)
                                        fmt.Println(v30, v30)
                                        for fi49 := 0; fi49 < 17; fi49++ {
                                            v3 = ((v6 | (v6 & -32768)) > v6)
                                            fmt.Println(v30, v5)
                                            _ = fi49
                                        }
                                    }
                                    v5 = (v5 + v5)
                                }
                            case 657:
                                var v50 uint16 = (v30 % ((v30) | 1))
                                _ = v50
                            case 32768:
                                v51 := GenS8{F0: -128}
                                _ = v51
                            default:
                                var v52 int16 = (v6 + (v6 >> ((4294967295) & 63)))
                                _ = v52
                            }
                            _ = fi33
                        }
                        var v53 string = v5
                        var v54 int = -886
                        _ = v53
                        _ = v54
                    } else {
                        if true {
                            v6 += (v6 - (v6 ^ (v6 << ((42) & 63))))
                            _ = len(v4)
                            v6 += v6
                        }
                        if true {
                            v1 = append(v1, 4294967295)
                            for fi55 := 0; fi55 < 14; fi55++ {
                                fmt.Println(v6, v3, v6)
                                v5 = (((("\n\t" + v5) + v5) + ((((v5 + ("hello" + "hello")) + "foo bar") + v5) + (v5 + "\n\t"))) + v5)
                                _ = fi55
                            }
                            v5 = v5
                        }
                        if false {
                            if (!true) {
                                var v56 int8 = 60
                                _ = v56
                            }
                            v30 = v30
                        } else {
                            for fi57 := 0; fi57 < 16; fi57++ {
                                for fi58 := 0; fi58 < 15; fi58++ {
                                    switch (v5 + "hello") {
                                    case "a":
                                        v30 += 740
                                    case "hello":
                                        if len(v2) > 0 {
                                            v2[0] = 156
                                        }
                                    }
                                    _ = fi58
                                }
                                v59 := GenS2{F0: 2}
                                v30 -= (v30 % ((137) | 1))
                                _ = fi57
                                _ = v59
                            }
                            if ((4294967295 != 0) || false) {
                                v5 = "foo bar"
                                v30 = (v30 * v30)
                            }
                            v60 := GenS9{F0: 18446744073709551615, F1: (v5 + "hello"), F2: "foo bar"}
                            _ = v60
                        }
                    }
                    _ = v30
                }
            }
        }
        switch (("hello" + (((v5 + "") + ("\n\t" + ((((v5 + "\n\t") + (v5 + v5)) + v5) + v5))) + v5)) + ((("\n\t" + (("" + "hello") + v5)) + "a") + v5)) {
        case "hello":
            var v61 int32 = -1
            _ = v61
        case "foo bar":
            for fi62 := 0; fi62 < 6; fi62++ {
                switch fi62 {
                case 4294967295:
                    for fi63 := 0; fi63 < 16; fi63++ {
                        var v64 uint64 = 830
                        _ = fi63
                        _ = v64
                    }
                case 703:
                    var v65 int = fi62
                    _ = v65
                default:
                    var v66 string = (((v5 + "") + v5) + "")
                    _ = v66
                }
                if (-596 >= -684) {
                    v67 := GenS3{F0: (!("" > "")), F1: (-(v6))}
                    _ = v67
                } else {
                    var v68 float64 = 0.0
                    switch 127 {
                    case 4294967295:
                        v3 = (255 <= 142)
                    case 128:
                        var v69 int8 = 127
                        _ = v69
                    case 42:
                        fmt.Println(v5, v6)
                    default:
                        var v70 bool = ((v5 + v5) >= "a")
                        _ = v70
                    }
                    _ = v68
                }
                switch fi62 {
                case 2147483647:
                    v3 = (3.14159 < ((float64(3.14159) + ((0.1 / ((0.0 / (3.14159 * 0.1)) + ((0.1 / 1.0) + (3.14159 - 1.5)))) * (1e10 - -2.25))) * 3.14159))
                case 128:
                    v5 = "hello"
                }
                _ = fi62
            }
        case "a":
            var v71 int = -2147483648
            _ = v71
        default:
            fmt.Println(v6, v5)
        }
    } else {
        var v72 uint8 = 152
        v5 = "\n\t"
        _ = v72
    }
    v6 = 32767
    switch 32767 {
    case -1:
        fmt.Println(v6, v5, v3)
    }
    _ = len(v1)
    v5 = (("" + "foo bar") + "\n\t")
    var v73 int16 = v6
    var v74 int8 = -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
}
