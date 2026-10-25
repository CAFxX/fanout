package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int16; F1 int; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint32) string {
    var v7 bool = true
    if v7 {
        var v8 string = (("" + "foo bar") + "a")
        var v9 float64 = (-2.25 / ((1e10 / float64((0.0 + ((0.1 + (-2.25 / float32((0.0 - float64(1e10))))) - 1e10)))) - (-2.25 / (0.1 + 1e10))))
        _ = v8
        _ = v9
    } else {
        fmt.Println(v7, p0)
        switch "foo bar" {
        case "\n\t":
            v10 := GenS1{F0: 158}
            _ = v10
        case "":
            var v11 uint = 42
            _ = v11
        case "a":
            var v12 int64 = 892
            _ = v12
        }
    }
    _ = p0
    _ = v7
    return "\n\t"
}

func main() {
    v1 := []uint32{}
    var v2 int16 = 32767
    var v3 uint64 = 18446744073709551615
    var v4 uint64 = v3
    var v5 uint64 = 2
    var v6 bool = (false || false)
    v13 := GenS2{F0: 100}
    v3 -= (v4 % ((v4) | 1))
    var v14 uint8 = 255
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        if false {
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 int64 = 9223372036854775807
                for fi17 := 0; fi17 < 19; fi17++ {
                    var v18 string = ("" + "\n\t")
                    if ((fi17 | 1) != (-(fi15))) {
                        fmt.Println(v16)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        var v19 uint16 = 0
                        _ = v19
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
            var v20 int64 = 42
            _ = v20
        }
        v2 += (^(v2))
    } else {
        var v21 string = "hello"
        _ = v21
    }
    if v6 {
        for fi22 := 0; fi22 < 14; fi22++ {
            for fi23 := 0; fi23 < 10; fi23++ {
                fmt.Println(v14, v6)
                var v24 uint32 = 4294967295
                v5 -= v5
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        v25 := GenS2{F0: 100}
        var v26 uint32 = 195
        _ = v25
        _ = v26
    } else {
        if (42 < 994) {
            v1 = append(v1, 42)
            v5 = v3
            var v27 uint64 = 18446744073709551615
            _ = v27
        } else {
            for fi28 := 0; fi28 < 10; fi28++ {
                v6 = (!((1e10 != 1.5) && (v5 < (v3 + v3))))
                switch "foo bar" {
                case "a":
                    fmt.Println(v3, fi28)
                case "\n\t":
                    v5 = (v3 / (((v4 / (((v5 * v5)) | 1))) | 1))
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v4 = (v5 + (v4 >> ((255) & 63)))
                _ = fi28
            }
        }
        v14 -= 2
    }
    switch 2147483647 {
    case -1:
        switch 694 {
        case 32767:
            v29 := GenS2{F0: 47}
            _ = v29
        case 1:
            var v30 int = -583
            _ = v30
        default:
            v14 += v14
        }
    case -2147483648:
        if (((0.0 / (((1.5 + (0.1 * float64(-2.25))) * ((3.14159 - 3.14159) + float64(float32((1e10 + 1e10))))) + (1e10 * 0.0))) * 3.14159) != float64(-2.25)) {
            if true {
                var v31 int32 = -549
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v31
            } else {
                for fi32 := 0; fi32 < 17; fi32++ {
                    if v6 {
                        for fi33 := 0; fi33 < 16; fi33++ {
                            fmt.Println(v6, v6)
                            switch "\n\t" {
                            case "foo bar":
                                v6 = (!false)
                            case "":
                                var v34 int8 = -89
                                _ = v34
                            default:
                                v3 = (v4 >> ((865) & 63))
                            }
                            _ = fi33
                        }
                    } else {
                        if (((^(v5)) >= v5) || (("" + ("a" + "\n\t")) >= "a")) {
                            for fi35 := 0; fi35 < 17; fi35++ {
                                if (v14 == v14) {
                                    v3 = (v4 + 2)
                                    v1 = append(v1, 4294967295)
                                }
                                if ((v6 || true) || (!(!(false || true)))) {
                                    v3 += (^(v3))
                                    var v36 uint8 = v14
                                    var v37 float64 = (0.0 - (1.5 * ((1.5 + 1e10) - float64((-2.25 * 1.5)))))
                                    v38 := GenS3{F0: 1000, F1: 42}
                                    var v39 float64 = (v37 * v37)
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                var v40 int16 = (v2 - v2)
                                switch 32768 {
                                case 127:
                                    v41 := GenS3{F0: (v2 * v2), F1: -9223372036854775808}
                                    _ = v41
                                case 2147483647:
                                    if (v6 && false) {
                                        v4 += (v5 | 18446744073709551615)
                                    } else {
                                        v3 -= 42
                                    }
                                case 32768:
                                    switch 128 {
                                    case -477:
                                        var v42 int32 = 2147483647
                                        _ = v42
                                    }
                                }
                                switch "hello" {
                                case "":
                                    var v43 uint16 = 42
                                    _ = v43
                                }
                                fmt.Println(v5, v2, v2)
                                _ = fi35
                                _ = v40
                            }
                            if (65535 <= 192) {
                                switch "a" {
                                case "foo bar":
                                    _ = len(v1)
                                }
                                fmt.Println(v6)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                v3 += (v5 % ((v3) | 1))
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                            }
                            var v44 uint32 = 766
                            _ = v44
                        } else {
                            if (-1 > -92) {
                                v45 := GenS3{F0: (v2 / ((1000) | 1)), F1: 0}
                                _ = v45
                            } else {
                                if ((v4 % ((0) | 1)) > (^(v5))) {
                                    var v46 string = ""
                                    _ = v46
                                } else {
                                    v47 := GenS4{F0: (v2 & (v2 * -32768)), F1: fi32, F2: -1}
                                    _ = v47
                                }
                                var v48 int64 = -9223372036854775808
                                if (!v6) {
                                    if (100 < -57) {
                                        v49 := GenS2{F0: 100}
                                        _ = v49
                                    }
                                } else {
                                    v14 = 255
                                    var v50 uint8 = v14
                                    _ = v50
                                }
                                _ = v48
                            }
                        }
                        var v51 uint8 = 128
                        _ = v51
                    }
                    _ = fi32
                }
                for fi52 := 0; fi52 < 2; fi52++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi52
                }
                var v53 int8 = -88
                _ = v53
            }
            if v6 {
                v4 += (v4 + 2)
                fmt.Println(v2, v5)
                for fi54 := 0; fi54 < 20; fi54++ {
                    if (!v6) {
                        var v55 uint = 2
                        _ = v55
                    }
                    var v56 bool = true
                    _ = fi54
                    _ = v56
                }
            } else {
                v57 := GenS1{F0: (v14 ^ 127)}
                _ = v57
            }
        }
    case 32768:
        for fi58 := 0; fi58 < 17; fi58++ {
            v6 = v6
            _ = fi58
        }
    }
    v5 = (v4 / ((2) | 1))
    fmt.Println(v5, v14)
    for fi59 := 0; fi59 < 2; fi59++ {
        for fi60 := 0; fi60 < 1; fi60++ {
            var v61 float64 = (1.5 - 0.1)
            if v6 {
                v62 := GenS4{F0: (v2 - v2), F1: fi59, F2: 619}
                v63 := GenS2{F0: -1}
                _ = v62
                _ = v63
            } else {
                v1 = append(v1, 652)
                var v64 uint = 65535
                _ = v64
            }
            if (!(2 < 255)) {
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                var v65 uint32 = 88
                _ = len(v1)
                _ = v65
            }
            _ = fi60
            _ = v61
        }
        fmt.Println(v5, v2)
        _ = fi59
    }
    v4 = v3
    switch int(v5) {
    case -9223372036854775808:
        for fi66 := 0; fi66 < 3; fi66++ {
            v1 = append(v1, 2)
            fmt.Println(v14, v4, v4)
            _ = fi66
        }
    case -1:
        for fi67 := 0; fi67 < 11; fi67++ {
            v5 = (v5 - v4)
            for fi68 := 0; fi68 < 15; fi68++ {
                v3 -= 128
                if (v14 < (v14 / ((166) | 1))) {
                    if (false && v6) {
                        if (1 >= -65) {
                            var v69 uint32 = 0
                            v14 += v14
                            if v6 {
                                v2 = v2
                            } else {
                                v70 := GenS4{F0: 394, F1: fi68, F2: 950}
                                v71 := GenS5{F0: "a"}
                                var v72 uint8 = v14
                                var v73 int64 = -64
                                _ = v70
                                _ = v71
                                _ = v72
                                _ = v73
                            }
                            _ = v69
                        } else {
                            if ((v4 / (((^(v3))) | 1)) == 18446744073709551615) {
                                v3 += 948
                                var v74 string = ""
                                var v75 uint32 = 2147483648
                                _ = v74
                                _ = v75
                            }
                        }
                        v1 = append(v1, 2147483647)
                        v4 = 42
                    }
                    if (!false) {
                        var v76 float64 = (3.14159 * 0.0)
                        v1 = append(v1, 42)
                        _ = v76
                    } else {
                        v4 -= v3
                        v1 = append(v1, 2147483648)
                    }
                    v14 = (v14 / (((v14 * v14)) | 1))
                } else {
                    if v6 {
                        v77 := GenS6{F0: v5}
                        _ = v77
                    }
                    var v78 uint32 = 2147483648
                    _ = v78
                }
                _ = fi68
            }
            v2 = 0
            _ = fi67
        }
    case 32768:
        var v79 string = ("\n\t" + "")
        _ = v79
    }
    var v80 int = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v80
}
