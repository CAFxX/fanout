package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int64; F1 string; F2 bool }
func genFn0(p0 uint32, p1 uint64, p2 uint8) int8 {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 1}
        v7 := []int{(fi5 | fi5)}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            switch int(p1) {
            case 2147483647:
                v10 := []int{fi8}
                p0 = p0
                _ = v10
            default:
                switch fi8 {
                case 291:
                    v11 := GenS1{F0: 9223372036854775807}
                    var v12 uint32 = p0
                    _ = v11
                    _ = v12
                }
            }
            if false {
                switch "" {
                case "":
                    p2 += 78
                case "\n\t":
                    var v13 uint16 = 164
                    _ = v13
                }
                fmt.Println(p0)
                v14 := GenS1{F0: 1}
                _ = v14
            } else {
                fmt.Println(fi9, fi8, fi8)
            }
            _ = fi9
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(fi8, fi15)
            _ = fi15
        }
        v16 := []int{(fi8 % ((312) | 1)), (^(fi8)), fi8, fi8}
        _ = fi8
        _ = v16
    }
    v17 := []int{}
    fmt.Println(p2, p1, p1)
    v17 = append(v17, -453)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 70
}
func genFn1(p0 int32, p1 uint32) string {
    v18 := GenS1{F0: 1}
    v19 := []int{42}
    v20 := GenS2{F0: 42, F1: (p1 ^ p1)}
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := GenS2{F0: 0, F1: p1}
        var v23 uint16 = 65535
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        if len(v19) > 0 {
            v19[0] = (fi24 << ((65535) & 63))
        }
        var v25 int16 = 1
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v20
    return ("a" + ("\n\t" + "foo bar"))
}

func main() {
    var v1 bool = false
    v2 := make(map[string]string)
    var v3 float64 = (-2.25 * -2.25)
    var v4 bool = (2 > 269)
    v26 := GenS1{F0: -9223372036854775808}
    var v27 int8 = 42
    for fi28 := 0; fi28 < 4; fi28++ {
        if (v4 || v1) {
            v29 := GenS2{F0: 18446744073709551615, F1: 808}
            _ = v29
        } else {
            v30 := GenS3{F0: -1}
            fmt.Println(v1)
            var v31 uint = 18446744073709551615
            _ = v30
            _ = v31
        }
        v3 = 0.0
        _ = fi28
    }
    var v32 int = 32768
    switch v32 {
    case 32767:
        v2["\n\t"] = "a"
    case -87:
        fmt.Println(v3, v4)
    case 2147483647:
        v3 = v3
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        _ = len(v2)
        _ = fi33
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        if false {
            v3 += v3
            for fi35 := 0; fi35 < 12; fi35++ {
                var v36 int8 = genFn0(162, 0, 72)
                fmt.Println(v27, fi35)
                var v37 int32 = -943
                _ = fi35
                _ = v36
                _ = v37
            }
            var v38 int = fi34
            _ = v38
        } else {
            v27 += (v27 - (v27 + 127))
            v4 = v4
        }
        _ = fi34
    }
    var v39 float64 = v3
    var v40 uint8 = 71
    if (v32 != (v32 ^ v32)) {
        var v41 bool = (v32 <= v32)
        for fi42 := 0; fi42 < 9; fi42++ {
            var v43 float32 = float32((1.5 + ((float64(3.14159) / v3) / (float64(float32(v39)) * (-2.25 + v39)))))
            for fi44 := 0; fi44 < 11; fi44++ {
                var v45 uint8 = v40
                v46 := GenS2{F0: 4294967295, F1: 1}
                _ = fi44
                _ = v45
                _ = v46
            }
            if v1 {
                v27 = v27
                if v41 {
                    var v47 string = (("" + "hello") + "a")
                    _ = v47
                } else {
                    v48 := []int{v32, v32, fi42}
                    switch 127 {
                    case -145:
                        for fi49 := 0; fi49 < 11; fi49++ {
                            if len(v48) >= 2 {
                                v48 = v48[1:len(v48)]
                            }
                            fmt.Println(fi42)
                            var v50 bool = v4
                            if v50 {
                                if v50 {
                                    v43 -= v43
                                    var v51 uint = 65535
                                    _ = v51
                                } else {
                                    var v52 uint32 = 9
                                    v53 := GenS3{F0: 113}
                                    v54 := v2["foo bar"]
                                    var v55 int = v32
                                    v55 = -9223372036854775808
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                                var v56 int16 = 902
                                _ = v56
                            } else {
                                v48 = append(v48, (fi42 & (fi42 + v32)))
                                var v57 int32 = -1
                                _ = len(v48)
                                _ = v57
                            }
                            var v58 uint16 = 663
                            _ = fi49
                            _ = v50
                            _ = v58
                        }
                    case -1:
                        v59 := GenS4{F0: 1, F1: fi42}
                        _ = v59
                    case -138:
                        var v60 int16 = -1
                        _ = v60
                    }
                    _ = v48
                }
                if (!true) {
                    switch "\n\t" {
                    case "":
                        v27 = v27
                    case "foo bar":
                        if (!(-810 < 9223372036854775807)) {
                            v61 := GenS2{F0: 517, F1: 0}
                            _ = v61
                        } else {
                            v62 := GenS5{F0: 2147483648, F1: 831, F2: 2}
                            var v63 uint8 = (v40 >> ((18446744073709551615) & 63))
                            _ = v62
                            _ = v63
                        }
                    case "\n\t":
                        v39 = 1.5
                    default:
                        v27 = (v27 / (((v27 % ((100) | 1))) | 1))
                    }
                    if ((2147483647 == 274) || v4) {
                        var v64 uint32 = 2147483647
                        _ = v64
                    }
                    if true {
                        var v65 uint64 = 660
                        for fi66 := 0; fi66 < 18; fi66++ {
                            v67 := []int{fi66, (fi42 | (v32 & fi66))}
                            v67 = append(v67, -410)
                            _ = fi66
                            _ = v67
                        }
                        _ = v65
                    }
                }
            } else {
                _ = len(v2)
                switch 2147483647 {
                case 128:
                    v41 = (1 < -3)
                case -737:
                    v68 := []int{}
                    _ = v68
                }
                for fi69 := 0; fi69 < 18; fi69++ {
                    v70 := GenS3{F0: -1}
                    v43 += 0.0
                    _ = fi69
                    _ = v70
                }
            }
            _ = fi42
            _ = v43
        }
        switch "a" {
        case "":
            fmt.Println(v4, v40)
        case "a":
            for fi71 := 0; fi71 < 3; fi71++ {
                switch ("\n\t" + "") {
                case "foo bar":
                    for fi72 := 0; fi72 < 6; fi72++ {
                        v32 += fi71
                        v32 -= (fi71 | (-(fi71)))
                        var v73 int64 = 831
                        _ = fi72
                        _ = v73
                    }
                case "\n\t":
                    v74 := GenS3{F0: 32767}
                    _ = v74
                case "a":
                    if v1 {
                        for fi75 := 0; fi75 < 13; fi75++ {
                            if v4 {
                                var v76 int16 = int16(v27)
                                if true {
                                    var v77 int8 = v27
                                    v2["\n\t"] = ("hello" + ("\n\t" + ("a" + "foo bar")))
                                    _ = v77
                                } else {
                                    v78 := GenS6{F0: -508, F1: "", F2: v4}
                                    v79 := []int{(v32 >> ((2) & 63)), fi75}
                                    var v80 uint64 = 2
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                }
                                _ = v76
                            } else {
                                if (65535 >= 2) {
                                    if v1 {
                                        var v81 int64 = 1
                                        v32 += v32
                                        _ = v81
                                    }
                                    if v4 {
                                        v32 -= v32
                                        var v82 float32 = (1.5 - float32(float64((float32(((v3 + (3.14159 - v39)) / v39)) * 0.1))))
                                        v83 := GenS5{F0: 2147483647, F1: 18446744073709551615, F2: 800}
                                        v39 = (v3 / (((float64(3.14159) - (v39 + (0.0 / 1.5))) / -2.25) - (0.0 * v39)))
                                        _ = v82
                                        _ = v83
                                    } else {
                                        fmt.Println(v39, v32, v39)
                                        v1 = ((3.14159 != v39) && (1e10 >= v39))
                                    }
                                } else {
                                    var v84 uint32 = 294
                                    _ = v84
                                }
                                fmt.Println(v4, v39)
                                switch fi75 {
                                case 2147483647:
                                    v85 := []int{v32, fi71}
                                    _ = v85
                                case 9223372036854775807:
                                    var v86 uint32 = 431
                                    _ = v86
                                case 871:
                                    var v87 float32 = 0.0
                                    _ = v87
                                }
                            }
                            _ = fi75
                        }
                    }
                default:
                    _ = len(v2)
                }
                if (0 != 0) {
                    v88 := []int{}
                    v89 := GenS3{F0: 1000}
                    var v90 int8 = -128
                    _ = v88
                    _ = v89
                    _ = v90
                }
                for fi91 := 0; fi91 < 12; fi91++ {
                    fmt.Println(v39, v3, v4)
                    _ = fi91
                }
                _ = fi71
            }
        default:
            var v92 uint16 = 0
            _ = v92
        }
        _ = v41
    } else {
        if (false || v1) {
            if false {
                var v93 uint32 = 4294967295
                v39 -= 1e10
                fmt.Println(v1, v1, v27)
                _ = v93
            }
            switch (v32 / ((v32) | 1)) {
            case -792:
                v32 = 2147483647
            default:
                var v94 float64 = v3
                _ = v94
            }
        }
        v40 = v40
    }
    v27 += (v27 - -123)
    v95 := []int{181}
    v96 := GenS4{F0: 237, F1: (v32 >> ((562) & 63))}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v95:", len(v95), v95)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v40
    _ = v95
    _ = v96
}
