package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 int16; F2 int32 }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 uint; F1 string; F2 int32 }
type GenS9 struct { F0 int8; F1 string; F2 uint16 }
func genFn0(p0 uint) uint64 {
    for fi5 := 0; fi5 < 16; fi5++ {
        p0 -= p0
        _ = fi5
    }
    var v6 float64 = 3.14159
    _ = p0
    _ = v6
    return 4294967295
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint = 18446744073709551615
    v3 := make(map[int]string)
    v4 := []int{4294967295, -130, 1}
    _ = len(v3)
    switch 362 {
    case 4294967295:
        for fi7 := 0; fi7 < 12; fi7++ {
            fmt.Println(v2, v1, v1)
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 int64 = -604
                var v10 string = "\n\t"
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    case -854:
        if (true && true) {
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 uint32 = 1
                v12 += (^(v12))
                v2 += (^(v2))
                _ = fi11
                _ = v12
            }
            fmt.Println(v2)
        } else {
            v3[520] = ("foo bar" + "hello")
            v13 := GenS1{F0: 4294967295, F1: 255}
            _ = v13
        }
    case 32768:
        var v14 bool = (!false)
        _ = v14
    }
    if (150 > 600) {
        switch -9223372036854775808 {
        case 4294967295:
            v2 += (v2 + 65535)
        default:
            v2 += (v2 >> (((v2 << (((v2 / ((2) | 1))) & 63))) & 63))
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v15 uint = (v2 >> ((v2) & 63))
    for fi16 := 0; fi16 < 19; fi16++ {
        v1 -= v1
        if false {
            for fi17 := 0; fi17 < 15; fi17++ {
                if (107 > 127) {
                    if (true || false) {
                        if false {
                            v2 = (v15 - (v15 | v2))
                            var v18 int32 = (v1 >> ((v15) & 63))
                            fmt.Println(v18)
                            _ = v18
                        } else {
                            if (1 == -825) {
                                if true {
                                    var v19 string = (("a" + "hello") + ("foo bar" + "a"))
                                    if false {
                                        var v20 uint32 = 79
                                        v3[(fi16 << ((v2) & 63))] = (v19 + "hello")
                                        var v21 int64 = -148
                                        _ = v20
                                        _ = v21
                                    } else {
                                        v22 := GenS2{F0: -128, F1: 0}
                                        var v23 float32 = (0.1 / ((3.14159 / 1.5) + 1e10))
                                        _ = v22
                                        _ = v23
                                    }
                                    var v24 int64 = -1
                                    _ = v19
                                    _ = v24
                                }
                            } else {
                                if true {
                                    v15 -= v2
                                    if (uint32(v2) >= 0) {
                                        v4 = append(v4, 89)
                                        fmt.Println(v2)
                                        if true {
                                            var v25 uint32 = 142
                                            if true {
                                                fmt.Println(v1)
                                                v26 := GenS1{F0: v25, F1: 0}
                                                v1 -= v1
                                                _ = v26
                                            }
                                            delete(v3, fi16)
                                            _ = v25
                                        }
                                        switch fi17 {
                                        case -2147483648:
                                            var v27 int32 = (v1 % ((950) | 1))
                                            _ = v27
                                        case 200:
                                            if true {
                                                var v28 int = fi16
                                                fmt.Println(v15, v2, v2)
                                                _ = v28
                                            }
                                        case 1:
                                            if (-9223372036854775808 != -9223372036854775808) {
                                                if (((v1 % (((v1 / ((v1) | 1))) | 1)) <= (v1 % (((v1 - v1)) | 1))) && true) {
                                                    if (2 > 590) {
                                                        fmt.Println(v1, v2, fi16)
                                                        _ = len(v3)
                                                        delete(v3, (fi16 >> ((v15) & 63)))
                                                    }
                                                }
                                                var v29 bool = (65535 == 1)
                                                var v30 string = ("" + "")
                                                _ = v29
                                                _ = v30
                                            }
                                        }
                                        fmt.Println(fi16, fi17)
                                        v4 = append(v4, fi17)
                                        fmt.Println(fi16, v15)
                                        if true {
                                            var v31 uint16 = 2
                                            v32 := GenS3{F0: (v1 | v1), F1: 261}
                                            v33 := GenS2{F0: -1, F1: 1}
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        } else {
                                            v34 := GenS2{F0: -128, F1: -1}
                                            _ = v34
                                        }
                                        var v35 int = 127
                                        var v36 float64 = 1e10
                                        v15 -= (v2 >> ((v15) & 63))
                                        v37 := GenS1{F0: 2, F1: 0}
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    } else {
                                        fmt.Println(v15, fi17, v1)
                                        v38 := GenS4{F0: v1, F1: "hello"}
                                        _ = v38
                                    }
                                    v1 += v1
                                    switch "\n\t" {
                                    case "a":
                                        v39 := GenS5{F0: 241}
                                        _ = v39
                                    case "\n\t":
                                        if len(v4) > 0 {
                                            v4[0] = 127
                                        }
                                    case "hello":
                                        switch "\n\t" {
                                        case "hello":
                                            v15 += (v2 | v15)
                                        case "":
                                            if true {
                                                var v40 float64 = -2.25
                                                switch fi16 {
                                                case 9223372036854775807:
                                                    var v41 bool = (("foo bar" + (("" + "\n\t") + "\n\t")) < "hello")
                                                    _ = v41
                                                case 32767:
                                                    if false {
                                                        var v42 int32 = 42
                                                        _ = v42
                                                    }
                                                default:
                                                    v43 := v3[fi17]
                                                    _ = v43
                                                }
                                                if true {
                                                    v4 = append(v4, (fi17 << ((v15) & 63)))
                                                    delete(v3, fi16)
                                                    var v44 float64 = ((((-2.25 / 1.0) + 0.1) * (v40 + v40)) + ((((v40 * 1.5) - v40) - (v40 - v40)) + (3.14159 - (1e10 / 3.14159))))
                                                    _ = v44
                                                }
                                                _ = v40
                                            }
                                        default:
                                            if (!false) {
                                                _ = len(v4)
                                                var v45 string = ""
                                                v46 := GenS5{F0: 127}
                                                _ = v45
                                                _ = v46
                                            }
                                            var v47 int = (fi17 >> ((v15) & 63))
                                            _ = v47
                                        }
                                    default:
                                        var v48 int64 = 9223372036854775807
                                        _ = v48
                                    }
                                    fmt.Println(v2)
                                    var v49 float32 = -2.25
                                    var v50 int64 = -193
                                    _ = v49
                                    _ = v50
                                } else {
                                    if true {
                                        v15 = v15
                                        v51 := GenS4{F0: v1, F1: ("foo bar" + "\n\t")}
                                        _ = v51
                                    }
                                }
                            }
                            v15 -= (v2 >> ((4294967295) & 63))
                        }
                    } else {
                        var v52 bool = (-37 > (v1 / (((v1 ^ 2147483647)) | 1)))
                        v53 := GenS3{F0: 1, F1: 42}
                        _ = v52
                        _ = v53
                    }
                }
                switch 128 {
                case 127:
                    var v54 uint16 = 0
                    var v55 int32 = v1
                    if false {
                        var v56 bool = true
                        fmt.Println(v1)
                        fmt.Println(v56)
                        _ = v56
                    } else {
                        v54 += 65535
                    }
                    _ = len(v4)
                    _ = v54
                    _ = v55
                case 281:
                    var v57 int = fi17
                    switch fi16 {
                    case 9223372036854775807:
                        fmt.Println(fi17, v57)
                    }
                    v57 = (-(fi17))
                    v57 = (v57 / (((^(v57))) | 1))
                    v15 -= (v15 % (((^(v2))) | 1))
                    v2 += v2
                    _ = len(v3)
                    _ = v57
                default:
                    var v58 int = 128
                    _ = v58
                }
                if (4294967295 > 2) {
                    if true {
                        var v59 int64 = -1
                        v2 = 65535
                        v4 = append(v4, (fi17 - (fi17 << ((1) & 63))))
                        _ = v59
                    } else {
                        switch (fi16 + fi17) {
                        case 2147483647:
                            _ = len(v3)
                        case 9223372036854775807:
                            switch (-(fi16)) {
                            case 217:
                                var v60 uint8 = 42
                                _ = v60
                            case -898:
                                v15 -= (^(v15))
                            case 1:
                                var v61 string = "foo bar"
                                _ = v61
                            default:
                                v4 = append(v4, fi17)
                            }
                        case -1:
                            if ((-(fi17)) == fi16) {
                                v2 = v2
                                if (!(9223372036854775807 >= 0)) {
                                    var v62 uint16 = 632
                                    v1 += v1
                                    _ = v62
                                } else {
                                    fmt.Println(v15, v2, v15)
                                    v63 := GenS6{F0: (v1 - 789), F1: 1, F2: v1}
                                    _ = v63
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    var v64 uint8 = 0
                                    _ = v64
                                case "foo bar":
                                    fmt.Println(v1, fi17, v2)
                                case "a":
                                    v65 := GenS7{F0: (fi17 - 1), F1: 127}
                                    _ = v65
                                }
                                fmt.Println(v1)
                            }
                        }
                    }
                } else {
                    fmt.Println(v15, v2, fi16)
                    _ = len(v3)
                    if ((fi16 % (((fi17 | fi17)) | 1)) > fi17) {
                        v66 := GenS7{F0: fi16, F1: fi16}
                        v67 := GenS8{F0: v2, F1: "hello", F2: (v1 | v1)}
                        var v68 uint64 = genFn0((v2 ^ 255))
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                }
                _ = fi17
            }
            v69 := GenS7{F0: (fi16 >> ((v15) & 63)), F1: (fi16 / ((fi16) | 1))}
            switch ((("foo bar" + "foo bar") + ("" + "")) + "a") {
            case "hello":
                for fi70 := 0; fi70 < 14; fi70++ {
                    v1 -= v1
                    v71 := GenS8{F0: 4294967295, F1: "foo bar", F2: 42}
                    if false {
                        if true {
                            v72 := GenS8{F0: (v2 >> ((94) & 63)), F1: "\n\t", F2: v1}
                            _ = v72
                        }
                        if true {
                            var v73 string = ("" + "")
                            fmt.Println(v15, v15)
                            _ = v73
                        }
                    }
                    _ = fi70
                    _ = v71
                }
            case "a":
                v74 := GenS8{F0: v15, F1: "\n\t", F2: (v1 & 1)}
                _ = v74
            default:
                fmt.Println(v15, v2, v15)
            }
            _ = v69
        } else {
            switch fi16 {
            case 4294967295:
                var v75 int64 = -1
                _ = v75
            case -735:
                fmt.Println(v1)
            default:
                var v76 int = fi16
                _ = v76
            }
            var v77 int16 = 920
            _ = v77
        }
        for fi78 := 0; fi78 < 9; fi78++ {
            var v79 bool = (1 > -32768)
            _ = fi78
            _ = v79
        }
        _ = fi16
    }
    for fi80 := 0; fi80 < 13; fi80++ {
        delete(v3, fi80)
        var v81 uint8 = 1
        v82 := GenS9{F0: -1, F1: (("\n\t" + ("\n\t" + "")) + "\n\t"), F2: 32767}
        _ = fi80
        _ = v81
        _ = v82
    }
    switch 4294967295 {
    case 32768:
        v2 = v15
    case 127:
        fmt.Println(v2)
    }
    v4 = append(v4, 32767)
    fmt.Println(v2)
    var v83 int16 = 1
    var v84 uint64 = genFn0(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v84
}
