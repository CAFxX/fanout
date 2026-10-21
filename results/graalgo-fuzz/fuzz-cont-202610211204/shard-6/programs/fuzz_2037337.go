package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int64; F2 bool }
type GenS4 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint16; F2 string }
type GenS7 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint16, p1 uint16) bool {
    v7 := []int{256, -151, -2147483648, 32767}
    var v8 int64 = 42
    p0 = (p0 + (p0 | p0))
    var v9 int64 = v8
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (2 == 34)
}
func genFn1() uint8 {
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 float64 = (((-2.25 - (3.14159 + 0.0)) / (0.1 + (1e10 - 0.0))) + -2.25)
        var v12 int32 = -366
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint64 = 18446744073709551615
    v13 = v13
    v14 := GenS1{F0: 30, F1: 0, F2: -960}
    _ = v13
    _ = v14
    return 255
}

func main() {
    var v1 string = ("" + ("hello" + (("\n\t" + ("a" + "a")) + "foo bar")))
    v2 := make(map[string]string)
    var v3 float32 = (1e10 + -2.25)
    var v4 bool = true
    v5 := []int{-1, -1}
    var v6 uint16 = 32767
    var v15 uint64 = 291
    var v16 int16 = 1000
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = (v4 || (97 >= uint32(v6)))
    v15 -= (^(v15))
    var v17 int = -576
    switch v1 {
    case "hello":
        v3 = (v3 / 1.0)
    case "":
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 uint8 = 0
            _ = fi18
            _ = v19
        }
    case "foo bar":
        for fi20 := 0; fi20 < 19; fi20++ {
            fmt.Println(v15)
            _ = fi20
        }
    default:
        v21 := GenS1{F0: 598, F1: 2147483648, F2: v17}
        _ = v21
    }
    fmt.Println(v15, v17)
    v1 = v1
    v22 := GenS1{F0: 42, F1: 65535, F2: v17}
    if (!false) {
        var v23 uint16 = (v6 / ((255) | 1))
        var v24 float64 = (float64(v3) * 1.5)
        if v4 {
            for fi25 := 0; fi25 < 11; fi25++ {
                v26 := v2[(v1 + v1)]
                if v4 {
                    delete(v2, v26)
                }
                switch v17 {
                case 127:
                    _ = len(v5)
                case 128:
                    if (255 != 0) {
                        for fi27 := 0; fi27 < 12; fi27++ {
                            var v28 uint8 = 42
                            v17 -= -165
                            if (((v3 + 3.14159) - ((0.1 * 1e10) / (3.14159 - 1.5))) <= ((v3 - float32(0.0)) * float32(((v24 + 1.5) + 0.0)))) {
                                v4 = (!v4)
                                var v29 uint8 = uint8(v15)
                                _ = v29
                            } else {
                                var v30 uint64 = v15
                                var v31 int = (fi27 & 127)
                                v26 = (v26 + ("a" + ""))
                                _ = v30
                                _ = v31
                            }
                            _ = fi27
                            _ = v28
                        }
                        var v32 int16 = (-(v16))
                        switch v17 {
                        case 919:
                            switch v1 {
                            case "foo bar":
                                for fi33 := 0; fi33 < 15; fi33++ {
                                    v16 += 1000
                                    v34 := GenS1{F0: 4294967295, F1: 1, F2: 0}
                                    _ = fi33
                                    _ = v34
                                }
                            case "hello":
                                if ((1e10 >= float32(v24)) || v4) {
                                    v5 = append(v5, (v17 >> ((4294967295) & 63)))
                                    for fi35 := 0; fi35 < 2; fi35++ {
                                        var v36 float32 = v3
                                        _ = fi35
                                        _ = v36
                                    }
                                }
                            case "":
                                if (-9223372036854775808 == int64(v16)) {
                                    v6 = 255
                                }
                            }
                        case -522:
                            v37 := GenS1{F0: 0, F1: 4294967295, F2: -9223372036854775808}
                            _ = v37
                        }
                        _ = v32
                    } else {
                        v17 = v17
                    }
                }
                _ = fi25
                _ = v26
            }
            fmt.Println(v1, v6, v3)
        } else {
            var v38 uint16 = v6
            _ = v38
        }
        _ = v23
        _ = v24
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v39 := GenS1{F0: 4294967295, F1: 653, F2: v17}
        switch v1 {
        case "a":
            for fi40 := 0; fi40 < 20; fi40++ {
                var v41 int64 = -9223372036854775808
                switch v17 {
                case 755:
                    _ = len(v5)
                case -9223372036854775808:
                    for fi42 := 0; fi42 < 3; fi42++ {
                        v1 = ("" + "a")
                        v41 = v41
                        if v4 {
                            var v43 string = ("hello" + "a")
                            v43 = ((((v1 + ((v43 + v43) + v43)) + v1) + "hello") + v43)
                            _ = v43
                        }
                        _ = fi42
                    }
                default:
                    switch "\n\t" {
                    case "a":
                        switch v1 {
                        case "\n\t":
                            v44 := GenS2{F0: (v1 + "foo bar")}
                            _ = v44
                        case "foo bar":
                            for fi45 := 0; fi45 < 2; fi45++ {
                                var v46 uint16 = v6
                                if (v3 != 3.14159) {
                                    var v47 int16 = v16
                                    v17 = (fi40 / ((4294967295) | 1))
                                    switch v1 {
                                    case "":
                                        switch -2147483648 {
                                        case 116:
                                            fmt.Println(v17, v47, v15)
                                        case 32767:
                                            var v48 uint16 = v46
                                            _ = v48
                                        }
                                    case "\n\t":
                                        v17 += (-(v17))
                                    case "hello":
                                        _ = len(v2)
                                        if false {
                                            switch v1 {
                                            case "hello":
                                                var v49 int16 = 1
                                                _ = v49
                                            case "":
                                                var v50 bool = true
                                                _ = v50
                                            }
                                            switch v17 {
                                            case -1:
                                                v1 = v1
                                            case -912:
                                                var v51 float64 = ((0.1 * 0.1) - 0.1)
                                                _ = v51
                                            default:
                                                if v4 {
                                                    if false {
                                                        var v52 int64 = (v41 << ((213) & 63))
                                                        _ = v52
                                                    }
                                                } else {
                                                    if (2147483648 >= 1) {
                                                        v1 = (v1 + "a")
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        var v53 float64 = 3.14159
                                                        _ = v53
                                                    }
                                                    var v54 float32 = 0.1
                                                    _ = v54
                                                }
                                            }
                                            var v55 string = v1
                                            _ = v55
                                        }
                                    }
                                    _ = len(v5)
                                    v56 := GenS3{F0: 42, F1: v41, F2: (!v4)}
                                    _ = v47
                                    _ = v56
                                }
                                _ = fi45
                                _ = v46
                            }
                        case "a":
                            if (v4 && (!(!v4))) {
                                switch (-(fi40)) {
                                case -1:
                                    for fi57 := 0; fi57 < 13; fi57++ {
                                        var v58 uint8 = 33
                                        var v59 uint64 = 1
                                        _ = fi57
                                        _ = v58
                                        _ = v59
                                    }
                                default:
                                    if false {
                                        v6 = (v6 & (v6 - (v6 >> ((65535) & 63))))
                                    }
                                }
                                v60 := v2[v1]
                                v61 := GenS4{F0: 9223372036854775807, F1: genFn1(), F2: 65535}
                                _ = v60
                                _ = v61
                            } else {
                                if false {
                                    var v62 float32 = v3
                                    _ = v62
                                }
                                v63 := GenS1{F0: 65535, F1: 0, F2: fi40}
                                _ = v63
                            }
                        default:
                            v2["\n\t"] = v1
                        }
                    default:
                        if v4 {
                            for fi64 := 0; fi64 < 15; fi64++ {
                                switch v1 {
                                case "hello":
                                    if true {
                                        v65 := GenS2{F0: (v1 + "")}
                                        var v66 bool = true
                                        _ = v65
                                        _ = v66
                                    }
                                case "\n\t":
                                    fmt.Println(v3, v6)
                                    v41 -= v41
                                    v67 := GenS5{F0: 0}
                                    _ = v67
                                }
                                v6 += v6
                                _ = fi64
                            }
                            v4 = v4
                        } else {
                            for fi68 := 0; fi68 < 5; fi68++ {
                                v41 = int64(v17)
                                switch fi40 {
                                case -969:
                                    v1 = v1
                                }
                                _ = fi68
                            }
                            if false {
                                v5 = append(v5, (fi40 * (fi40 + (^(v17)))))
                                if (139 == 137) {
                                    var v69 float32 = (v3 * (v3 - v3))
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v69
                                } else {
                                    for fi70 := 0; fi70 < 9; fi70++ {
                                        v15 -= (v15 + v15)
                                        if v4 {
                                            if true {
                                                switch (v17 ^ fi40) {
                                                case 9223372036854775807:
                                                    switch ((v1 + (v1 + (((((v1 + "") + ("\n\t" + v1)) + v1) + v1) + v1))) + v1) {
                                                    case "\n\t":
                                                        fmt.Println(v3, v3)
                                                    }
                                                case 0:
                                                    if true {
                                                        if (820 == v6) {
                                                            var v71 uint64 = v15
                                                            fmt.Println(v1, v1, v6)
                                                            if (v3 > 0.1) {
                                                                var v72 int = -620
                                                                _ = v72
                                                            } else {
                                                                v73 := GenS5{F0: -1}
                                                                _ = v73
                                                            }
                                                            switch ((v1 + v1) + v1) {
                                                            case "foo bar":
                                                                v74 := GenS4{F0: v41, F1: genFn1(), F2: 4294967295}
                                                                _ = v74
                                                            case "\n\t":
                                                                v6 -= v6
                                                            }
                                                            v41 -= -9223372036854775808
                                                            _ = v71
                                                        } else {
                                                            switch v1 {
                                                            case "a":
                                                                var v75 int16 = (-(v16))
                                                                _ = v75
                                                            case "\n\t":
                                                                var v76 uint32 = 2
                                                                _ = v76
                                                            }
                                                            if (v4 || v4) {
                                                                var v77 uint32 = 673
                                                                var v78 uint64 = v15
                                                                fmt.Println(v1, v4)
                                                                _ = v77
                                                                _ = v78
                                                            } else {
                                                                v79 := GenS3{F0: 0, F1: (v41 | 915), F2: (float64(((v3 / ((((v3 / float32((1.5 - ((0.0 - 0.1) - 0.1)))) - ((((1e10 * float32((1.5 * 1.5))) * v3) * v3) - 0.0)) - 3.14159) + v3)) - (-2.25 - v3))) != 0.0)}
                                                                v4 = (!true)
                                                                _ = v79
                                                            }
                                                            delete(v2, ("hello" + "foo bar"))
                                                        }
                                                    } else {
                                                        if ((1 != 100) || true) {
                                                            if (!(v4 && v4)) {
                                                                fmt.Println(v41, v16)
                                                                var v80 uint32 = 2
                                                                _ = v80
                                                            } else {
                                                                var v81 uint = 72
                                                                _ = v81
                                                            }
                                                            var v82 int64 = v41
                                                            var v83 string = v1
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                    }
                                                case 42:
                                                    fmt.Println(v4)
                                                }
                                                var v84 uint64 = (^(v15))
                                                _ = v84
                                            }
                                        }
                                        v15 = (v15 | v15)
                                        _ = fi70
                                    }
                                }
                            } else {
                                for fi85 := 0; fi85 < 13; fi85++ {
                                    var v86 int16 = -1
                                    _ = fi85
                                    _ = v86
                                }
                            }
                        }
                    }
                }
                if (v3 != ((((float32(((0.1 + float64(v3)) * 3.14159)) + v3) + float32(0.1)) / float32((1e10 + (-2.25 - float64(float32(1.5)))))) + (float32(0.1) + (v3 + (0.1 / 1.0))))) {
                    var v87 int16 = v16
                    _ = v87
                } else {
                    v41 += (v41 << ((18446744073709551615) & 63))
                }
                _ = fi40
                _ = v41
            }
        case "\n\t":
            var v88 uint32 = 4294967295
            _ = v88
        case "hello":
            var v89 float64 = ((-2.25 + (3.14159 + 1e10)) * 1.5)
            _ = v89
        default:
            v90 := GenS6{F0: 59, F1: v6, F2: v1}
            _ = v90
        }
        _ = v39
    }
    var v91 float64 = -2.25
    v2[(((v1 + v1) + "hello") + v1)] = (((((v1 + v1) + v1) + v1) + (("foo bar" + v1) + v1)) + ((v1 + v1) + v1))
    v92 := GenS7{F0: 1, F1: 2147483647}
    v93 := v2[("foo bar" + v1)]
    v91 = (float64(1e10) - (v91 - v91))
    v91 = (v91 / v91)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
    _ = v92
    _ = v93
}
