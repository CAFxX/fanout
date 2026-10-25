package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint; F2 int32 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 bool; F1 int16; F2 uint16 }
type GenS7 struct { F0 int64; F1 bool; F2 int16 }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 int32 }
func genFn0(p0 uint32) uint64 {
    var v7 float64 = 3.14159
    var v8 int8 = 82
    v9 := []int{-2147483648, 0, 615, 32768}
    var v10 uint8 = 42
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 float64 = -2.25
    var v2 string = (((("hello" + ("hello" + "")) + (("" + "\n\t") + "")) + "hello") + "foo bar")
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    var v5 int32 = -1
    var v6 float32 = 1.5
    v11 := []int{128}
    v1 = float64(v6)
    if true {
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := v3[940]
            if v13 {
                if v13 {
                    v14 := v3[556]
                    if v14 {
                        v5 -= v5
                        switch (fi12 << ((455) & 63)) {
                        case 0:
                            v14 = (1000 == -32768)
                        default:
                            fmt.Println(v13, v6)
                        }
                    } else {
                        _ = len(v4)
                    }
                    v14 = v13
                    _ = v14
                } else {
                    if ((-2.25 != 0.0) || v13) {
                        if v13 {
                            for fi15 := 0; fi15 < 19; fi15++ {
                                var v16 uint16 = 1
                                v17 := GenS1{F0: 18446744073709551615, F1: 0}
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 4; fi18++ {
                                if v13 {
                                    v13 = (v13 && ((!(v1 == v1)) || (18446744073709551615 != 1)))
                                    var v19 uint16 = 419
                                    _ = v19
                                }
                                var v20 uint = 4294967295
                                _ = fi18
                                _ = v20
                            }
                            switch (v2 + v2) {
                            case "hello":
                                var v21 uint8 = 0
                                _ = v21
                            case "a":
                                if true {
                                    var v22 int16 = 1000
                                    fmt.Println(v5)
                                    v1 = (float64(((((float32((v1 / (3.14159 / -2.25))) - (v6 - 0.1)) - v6) * v6) / v6)) + v1)
                                    _ = v22
                                } else {
                                    if (-1 >= 42) {
                                        if false {
                                            for fi23 := 0; fi23 < 16; fi23++ {
                                                v6 += (v6 * (v6 - 0.0))
                                                if len(v11) >= 2 {
                                                    v11 = v11[1:len(v11)]
                                                }
                                                v6 -= v6
                                                v24 := GenS1{F0: 0, F1: 0}
                                                _ = fi23
                                                _ = v24
                                            }
                                            if (0 <= 32767) {
                                                if false {
                                                    delete(v4, (-(fi12)))
                                                    switch v2 {
                                                    case "\n\t":
                                                        v6 = (3.14159 + (v6 / v6))
                                                    case "a":
                                                        v25 := GenS1{F0: 1, F1: 2}
                                                        _ = v25
                                                    case "":
                                                        v2 = v2
                                                    default:
                                                        var v26 float64 = 1e10
                                                        _ = v26
                                                    }
                                                    v1 += ((v1 - v1) + (1e10 - ((0.0 + (float64(v6) - (1e10 * v1))) + (v1 / v1))))
                                                }
                                            } else {
                                                v27 := GenS1{F0: 1, F1: 2}
                                                _ = v27
                                            }
                                            if v13 {
                                                v13 = (v13 || true)
                                                v4[fi12] = "hello"
                                                v28 := GenS2{F0: v5, F1: 42}
                                                _ = v28
                                            } else {
                                                if v13 {
                                                    switch (fi12 % (((fi12 << ((255) & 63))) | 1)) {
                                                    case 4294967295:
                                                        v29 := GenS3{F0: (v2 + (v2 + v2))}
                                                        _ = v29
                                                    case 862:
                                                        for fi30 := 0; fi30 < 18; fi30++ {
                                                            v2 = ""
                                                            v6 -= float32(v1)
                                                            var v31 uint8 = 255
                                                            _ = fi30
                                                            _ = v31
                                                        }
                                                    case 32768:
                                                        v32 := GenS4{F0: (100 > 26), F1: 42, F2: v5}
                                                        _ = v32
                                                    }
                                                    v13 = v13
                                                }
                                            }
                                        } else {
                                            var v33 int32 = (v5 - v5)
                                            _ = v33
                                        }
                                        fmt.Println(v2)
                                    } else {
                                        var v34 uint8 = 255
                                        v13 = (0 > 255)
                                        v1 = v1
                                        _ = v34
                                    }
                                    var v35 uint64 = 1
                                    fmt.Println(v2, v6, v6)
                                    _ = v35
                                }
                            default:
                                for fi36 := 0; fi36 < 20; fi36++ {
                                    v2 = (v2 + ((((v2 + v2) + v2) + v2) + "foo bar"))
                                    _ = fi36
                                }
                            }
                        } else {
                            v11 = append(v11, (fi12 / ((fi12) | 1)))
                            var v37 int = 128
                            _ = v37
                        }
                        v38 := GenS3{F0: v2}
                        switch "foo bar" {
                        case "a":
                            var v39 bool = true
                            _ = v39
                        case "\n\t":
                            var v40 int64 = 42
                            _ = v40
                        case "":
                            if len(v11) > 0 {
                                v11[0] = fi12
                            }
                        default:
                            v41 := v3[665]
                            _ = v41
                        }
                        _ = v38
                    }
                    var v42 int64 = -846
                    _ = v42
                }
                var v43 float64 = (float64(1.5) - 0.0)
                v2 = "a"
                _ = v43
            } else {
                var v44 int16 = -32768
                fmt.Println(v44, v2)
                v1 -= (float64(0.1) - float64(1e10))
                _ = v44
            }
            _ = fi12
            _ = v13
        }
    } else {
        v45 := GenS5{F0: 2, F1: "a"}
        _ = v45
    }
    v1 = (v1 / (v1 + v1))
    if false {
        var v46 bool = true
        if v46 {
            switch v2 {
            case "\n\t":
                switch ((((v2 + "a") + "\n\t") + v2) + "hello") {
                case "\n\t":
                    for fi47 := 0; fi47 < 13; fi47++ {
                        var v48 bool = false
                        _ = fi47
                        _ = v48
                    }
                case "a":
                    if v46 {
                        v11 = append(v11, 42)
                        fmt.Println(v5)
                    }
                case "foo bar":
                    var v49 int = 159
                    _ = v49
                default:
                    if (4294967295 <= 533) {
                        var v50 float32 = 1e10
                        var v51 uint32 = 2147483647
                        if true {
                            if ((v51 + (v51 ^ (v51 / ((65535) | 1)))) != 65535) {
                                if (-2147483648 >= 0) {
                                    v6 = (((v6 / float32((v1 - v1))) + float32(((0.1 / float64(v50)) + float64(1e10)))) / 0.1)
                                    v52 := GenS6{F0: (v6 != v50), F1: 298, F2: 198}
                                    var v53 string = v2
                                    _ = v52
                                    _ = v53
                                } else {
                                    for fi54 := 0; fi54 < 3; fi54++ {
                                        for fi55 := 0; fi55 < 12; fi55++ {
                                            var v56 uint = 255
                                            _ = fi55
                                            _ = v56
                                        }
                                        v57 := GenS6{F0: false, F1: 42, F2: 2}
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                        _ = fi54
                                        _ = v57
                                    }
                                    if (0 < -334) {
                                        var v58 float64 = 1e10
                                        for fi59 := 0; fi59 < 19; fi59++ {
                                            v5 = -876
                                            var v60 int32 = -123
                                            _ = fi59
                                            _ = v60
                                        }
                                        _ = v58
                                    }
                                }
                                fmt.Println(v46, v51, v51)
                            }
                            v46 = v46
                        } else {
                            if v46 {
                                if false {
                                    v6 = ((((v6 / float32(v1)) - v6) / (0.0 - v6)) * (float32((float64(0.0) / (1e10 * v1))) + -2.25))
                                    if true {
                                        var v61 uint64 = 680
                                        for fi62 := 0; fi62 < 4; fi62++ {
                                            if (14 < 49) {
                                                for fi63 := 0; fi63 < 19; fi63++ {
                                                    var v64 uint32 = (^(v51))
                                                    v50 = 0.1
                                                    var v65 float64 = v1
                                                    _ = fi63
                                                    _ = v64
                                                    _ = v65
                                                }
                                            } else {
                                                var v66 int64 = 9223372036854775807
                                                fmt.Println(fi62)
                                                var v67 int64 = (v66 * -217)
                                                _ = v66
                                                _ = v67
                                            }
                                            var v68 int16 = 841
                                            var v69 uint32 = (v51 >> ((18446744073709551615) & 63))
                                            _ = fi62
                                            _ = v68
                                            _ = v69
                                        }
                                        _ = v61
                                    }
                                }
                                delete(v3, (v51 - v51))
                            } else {
                                if len(v11) > 0 {
                                    v11[0] = 4294967295
                                }
                            }
                            v46 = v46
                        }
                        _ = v50
                        _ = v51
                    }
                }
            }
            for fi70 := 0; fi70 < 18; fi70++ {
                switch "a" {
                case "":
                    v71 := GenS5{F0: 197, F1: (((("" + v2) + v2) + v2) + "\n\t")}
                    _ = v71
                }
                switch v2 {
                case "hello":
                    v5 -= 0
                case "":
                    delete(v3, 878)
                default:
                    fmt.Println(v5, v5)
                }
                if (114 != 2) {
                    v11 = append(v11, (-(fi70)))
                    v6 = 0.1
                } else {
                    fmt.Println(fi70, v6)
                    if v46 {
                        var v72 int32 = v5
                        _ = v72
                    } else {
                        fmt.Println(v5)
                    }
                    if (255 > 958) {
                        if (566 > 0) {
                            var v73 uint16 = 180
                            if v46 {
                                v6 += (v6 + (v6 + v6))
                            } else {
                                var v74 int8 = 42
                                _ = v74
                            }
                            if v46 {
                                if v46 {
                                    var v75 uint64 = 0
                                    _ = v75
                                }
                                var v76 string = (v2 + v2)
                                _ = v76
                            }
                            _ = v73
                        }
                        switch ((v2 + (v2 + v2)) + ("" + (v2 + v2))) {
                        case "foo bar":
                            for fi77 := 0; fi77 < 10; fi77++ {
                                v78 := GenS2{F0: (v5 & (v5 * v5)), F1: 38}
                                fmt.Println(fi77)
                                v79 := GenS7{F0: int64(fi70), F1: v46, F2: 32767}
                                fmt.Println(v5, v46)
                                _ = fi77
                                _ = v78
                                _ = v79
                            }
                        case "a":
                            if ((1000 != int16(v5)) && (false || true)) {
                                var v80 int64 = 42
                                v1 = float64(float32((float64(float32(3.14159)) / v1)))
                                for fi81 := 0; fi81 < 14; fi81++ {
                                    switch v2 {
                                    case "":
                                        v5 -= 876
                                        fmt.Println(v5)
                                    case "hello":
                                        _ = len(v4)
                                    case "foo bar":
                                        v82 := GenS8{F0: 0, F1: v80}
                                        _ = v82
                                    }
                                    v1 = v1
                                    fmt.Println(v46)
                                    _ = fi81
                                }
                                _ = v80
                            }
                        }
                    } else {
                        for fi83 := 0; fi83 < 11; fi83++ {
                            var v84 int = (fi70 ^ (-(fi83)))
                            _ = fi83
                            _ = v84
                        }
                        fmt.Println(v2, v5)
                        v1 = v1
                    }
                }
                _ = fi70
            }
        }
        _ = v46
    } else {
        var v85 float32 = float32((float64(v6) - (v1 - 0.0)))
        _ = v85
    }
    if len(v11) > 0 {
        v11[0] = 4294967295
    }
    for fi86 := 0; fi86 < 1; fi86++ {
        if false {
            v5 += v5
            v5 += v5
            v3[4294967295] = false
        } else {
            for fi87 := 0; fi87 < 11; fi87++ {
                var v88 uint8 = 26
                v89 := GenS9{F0: v5}
                _ = fi87
                _ = v88
                _ = v89
            }
            _ = len(v3)
            var v90 int = fi86
            _ = v90
        }
        v11 = append(v11, 32767)
        _ = fi86
    }
    if (-423 <= -9223372036854775808) {
        v91 := GenS6{F0: (((966 > 255) || true) || true), F1: 404, F2: 956}
        _ = v91
    } else {
        var v92 int64 = 9223372036854775807
        v92 += (v92 ^ 9223372036854775807)
        _ = v92
    }
    var v93 int8 = -29
    v93 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v93
}
