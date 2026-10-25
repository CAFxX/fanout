package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint8; F2 int32 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 uint32; F2 int64 }
func genFn0() bool {
    var v7 float32 = 1e10
    var v8 bool = (42 <= 18446744073709551615)
    switch "foo bar" {
    case "":
        switch 42 {
        case -1:
            v9 := []int{}
            _ = v9
        case 0:
            var v10 uint64 = 42
            _ = v10
        default:
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := []int{(^(fi11)), 4294967295, 32768, (fi11 << ((183) & 63))}
                var v13 uint64 = 502
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    case "\n\t":
        var v14 float32 = v7
        _ = v14
    case "hello":
        v15 := []int{}
        _ = v15
    }
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 int = -9223372036854775808
    var v2 uint32 = 4294967295
    var v3 uint8 = 0
    var v4 string = "a"
    v5 := make(map[int]bool)
    v6 := make(map[string]bool)
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint16 = 65535
        fmt.Println(v17, v1)
        _ = fi16
        _ = v17
    }
    if false {
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 int8 = 55
            delete(v5, int(v3))
            _ = fi18
            _ = v19
        }
    } else {
        fmt.Println(v4, v4)
        var v20 string = ""
        _ = v20
    }
    v21 := GenS1{F0: -444}
    v22 := []int{(v1 * (v1 >> ((65535) & 63))), (v1 / ((v1) | 1))}
    switch (("hello" + "") + ("\n\t" + "a")) {
    case "":
        if (-2.25 <= 1.5) {
            v22 = append(v22, 1)
            v23 := GenS1{F0: 0}
            var v24 int64 = 907
            _ = v23
            _ = v24
        } else {
            switch v1 {
            case 128:
                v25 := GenS1{F0: -9223372036854775808}
                _ = v25
            case -947:
                for fi26 := 0; fi26 < 20; fi26++ {
                    var v27 int64 = 894
                    _ = fi26
                    _ = v27
                }
            case 830:
                switch v4 {
                case "a":
                    var v28 uint8 = (v3 * (v3 * v3))
                    _ = v28
                default:
                    v29 := GenS1{F0: 1}
                    _ = v29
                }
            }
            v2 += 320
            if false {
                for fi30 := 0; fi30 < 10; fi30++ {
                    switch v4 {
                    case "\n\t":
                        if (!false) {
                            delete(v5, v1)
                            if len(v22) > 0 {
                                v22[0] = (v1 << ((uint(v3)) & 63))
                            }
                            for fi31 := 0; fi31 < 13; fi31++ {
                                switch fi31 {
                                case -419:
                                    var v32 uint64 = 50
                                    _ = v32
                                case 127:
                                    if (v3 < v3) {
                                        switch fi30 {
                                        case 128:
                                            v3 += v3
                                        case 127:
                                            v33 := GenS1{F0: 302}
                                            _ = v33
                                        case 2147483647:
                                            if (42 == 2) {
                                                if true {
                                                    if false {
                                                        v1 = fi30
                                                        switch fi31 {
                                                        case 32767:
                                                            var v34 uint16 = 983
                                                            _ = v34
                                                        }
                                                        v22 = append(v22, 217)
                                                        switch -2147483648 {
                                                        case 127:
                                                            if (false || ((!((v2 % (((^(v2))) | 1)) >= 491)) || false)) {
                                                                switch (v4 + ("hello" + v4)) {
                                                                case "\n\t":
                                                                    switch v4 {
                                                                    case "":
                                                                        var v35 float64 = 3.14159
                                                                        _ = v35
                                                                    case "a":
                                                                        fmt.Println(v3)
                                                                    case "foo bar":
                                                                        v22 = append(v22, fi31)
                                                                    default:
                                                                        var v36 string = v4
                                                                        _ = v36
                                                                    }
                                                                case "a":
                                                                    var v37 int32 = -1
                                                                    _ = v37
                                                                }
                                                            } else {
                                                                var v38 int8 = int8(v2)
                                                                if false {
                                                                    v39 := v5[2147483647]
                                                                    v40 := GenS1{F0: -739}
                                                                    v41 := GenS1{F0: 9223372036854775807}
                                                                    if (("\n\t" + v4) >= v4) {
                                                                        var v42 uint32 = v2
                                                                        v38 = (^(v38))
                                                                        if (2 >= 4294967295) {
                                                                            v3 += v3
                                                                        }
                                                                        _ = v42
                                                                    } else {
                                                                        if (!false) {
                                                                            if v39 {
                                                                                v43 := GenS2{F0: (v2 * v2), F1: 18446744073709551615, F2: 42}
                                                                                v44 := GenS3{F0: 42, F1: (v3 + v3), F2: 2147483647}
                                                                                if false {
                                                                                    var v45 bool = (1 >= 1)
                                                                                    var v46 int = (fi31 + v1)
                                                                                    var v47 uint = 4294967295
                                                                                    v45 = (int8(v3) != -102)
                                                                                    _ = v45
                                                                                    _ = v46
                                                                                    _ = v47
                                                                                }
                                                                                var v48 int = 244
                                                                                _ = v43
                                                                                _ = v44
                                                                                _ = v48
                                                                            }
                                                                            var v49 uint = 493
                                                                            _ = v49
                                                                        } else {
                                                                            v50 := GenS4{F0: -689, F1: 0}
                                                                            var v51 string = v4
                                                                            _ = v50
                                                                            _ = v51
                                                                        }
                                                                        switch v4 {
                                                                        case "":
                                                                            delete(v5, (fi30 & 304))
                                                                        }
                                                                        v4 = "foo bar"
                                                                        if ((860 != 981) || (((1.5 * 1.5) - 1.5) < (float32(0.1) / 1.0))) {
                                                                            v3 -= 180
                                                                            v38 += v38
                                                                        }
                                                                    }
                                                                    _ = v39
                                                                    _ = v40
                                                                    _ = v41
                                                                } else {
                                                                    var v52 int32 = 147
                                                                    _ = v52
                                                                }
                                                                _ = v38
                                                            }
                                                        case 4294967295:
                                                            v2 -= (v2 / ((2147483648) | 1))
                                                        }
                                                        var v53 int64 = 9223372036854775807
                                                        _ = v53
                                                    } else {
                                                        var v54 int64 = 288
                                                        if len(v22) >= 2 {
                                                            v22 = v22[1:len(v22)]
                                                        }
                                                        _ = v54
                                                    }
                                                    if (true || ((3.14159 == (((-2.25 / 1.5) + (0.1 * ((0.1 * (0.1 * (((float64((0.1 + (0.1 - 1e10))) + 0.0) / float64((-2.25 + float32(float64(0.1))))) / 0.1))) + ((-2.25 + ((3.14159 * 1.5) * -2.25)) * ((1.5 - 0.0) * ((((1.5 * 0.1) - 1.5) + 0.0) * (1.5 * (0.0 - 0.1)))))))) - 3.14159)) && true)) {
                                                        var v55 uint32 = 407
                                                        v4 = ((v4 + ("a" + v4)) + (v4 + v4))
                                                        _ = v55
                                                    }
                                                    v4 = (v4 + (("" + v4) + ((v4 + v4) + v4)))
                                                } else {
                                                    if (!false) {
                                                        var v56 bool = true
                                                        _ = v56
                                                    } else {
                                                        v57 := v6[""]
                                                        v4 = v4
                                                        v58 := GenS3{F0: -1, F1: 128, F2: 1}
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    var v59 int = -2147483648
                                                    _ = v59
                                                }
                                                v60 := GenS5{F0: v4}
                                                _ = v60
                                            }
                                        }
                                    } else {
                                        if false {
                                            v61 := GenS4{F0: 1, F1: 467}
                                            fmt.Println(fi31, v2)
                                            v6["\n\t"] = false
                                            var v62 bool = true
                                            _ = v61
                                            _ = v62
                                        }
                                    }
                                default:
                                    v4 = ("foo bar" + v4)
                                }
                                fmt.Println(v1, v3)
                                _ = fi31
                            }
                        }
                    }
                    _ = fi30
                }
            }
        }
    case "foo bar":
        v63 := GenS3{F0: 42, F1: (v3 | 255), F2: 42}
        _ = v63
    case "a":
        if ((float64((0.1 + (((0.0 * 1e10) + ((0.0 / float32(-2.25)) / 1.0)) / 1e10))) + 0.0) != 1.5) {
            fmt.Println(v4, v2, v2)
            var v64 float32 = float32((float64(3.14159) / (1.5 * float64(-2.25))))
            v22 = append(v22, (-(v1)))
            _ = v64
        }
    default:
        var v65 int = (v1 / (((^(v1))) | 1))
        _ = v65
    }
    if (!false) {
        _ = len(v22)
    } else {
        v66 := GenS6{F0: 2}
        var v67 bool = (false && (-32768 >= -317))
        v3 = (v3 - (v3 ^ (v3 | 2)))
        _ = v66
        _ = v67
    }
    var v68 bool = true
    v3 += (v3 >> ((65535) & 63))
    if v68 {
        if v68 {
            v69 := GenS2{F0: v2, F1: 42, F2: 4294967295}
            _ = v69
        }
        switch (-(v1)) {
        case -9223372036854775808:
            v3 -= (v3 - v3)
        }
        switch v4 {
        case "foo bar":
            if false {
                var v70 uint16 = 946
                fmt.Println(v2)
                _ = v70
            }
        case "\n\t":
            if v68 {
                if (-50 > -128) {
                    v71 := GenS6{F0: 255}
                    _ = v71
                }
            } else {
                for fi72 := 0; fi72 < 20; fi72++ {
                    var v73 int8 = int8(v3)
                    _ = fi72
                    _ = v73
                }
            }
        }
    } else {
        switch 601 {
        case -736:
            if false {
                v3 = v3
                var v74 uint = 295
                switch ("a" + "\n\t") {
                case "hello":
                    for fi75 := 0; fi75 < 6; fi75++ {
                        var v76 int16 = 42
                        _ = fi75
                        _ = v76
                    }
                }
                _ = v74
            } else {
                var v77 int64 = -502
                _ = v77
            }
        case 9223372036854775807:
            var v78 string = ("" + (v4 + (v4 + v4)))
            _ = v78
        case 128:
            _ = len(v5)
        default:
            v79 := GenS1{F0: 0}
            _ = v79
        }
        v80 := GenS7{F0: 0, F1: 0, F2: 9223372036854775807}
        _ = v80
    }
    v68 = false
    v3 = (v3 | v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
}
