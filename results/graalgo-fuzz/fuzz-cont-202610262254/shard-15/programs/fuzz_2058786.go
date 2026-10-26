package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int16; F1 bool; F2 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 uint; F2 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint; F1 uint64; F2 uint64 }
func genFn0(p0 uint, p1 int16, p2 int64) string {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 float32 = 3.14159
        switch ("hello" + ("\n\t" + "a")) {
        case "":
            fmt.Println(p0, fi7, p1)
        case "hello":
            if (((float64(v8) + -2.25) / (1.5 + (float64(v8) * ((0.0 + float64(v8)) * -2.25)))) > 1e10) {
                if true {
                    fmt.Println(p2, p0, p0)
                } else {
                    v9 := []int{fi7}
                    p2 -= p2
                    _ = len(v9)
                    _ = v9
                }
                p2 += -9223372036854775808
            } else {
                var v10 int16 = 42
                _ = v10
            }
        case "\n\t":
            v8 += -2.25
        default:
            v11 := []int{(fi7 * fi7)}
            _ = v11
        }
        var v12 uint16 = 872
        _ = fi7
        _ = v12
        _ = v8
    }
    v13 := GenS1{F0: 1, F1: 1}
    if (!(p0 < p0)) {
        if (!(672 < 4294967295)) {
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 uint32 = 1
                var v16 float64 = (0.1 - -2.25)
                _ = fi14
                _ = v15
                _ = v16
            }
            if (!(42 == 18)) {
                switch "a" {
                case "foo bar":
                    switch "hello" {
                    case "foo bar":
                        v17 := GenS1{F0: 127, F1: 1}
                        _ = v17
                    }
                case "\n\t":
                    fmt.Println(p2)
                }
                if ("foo bar" != "a") {
                    if false {
                        p1 = p1
                    } else {
                        var v18 float32 = ((float32(float64(3.14159)) * 3.14159) * (1e10 - -2.25))
                        var v19 uint64 = 2
                        _ = v18
                        _ = v19
                    }
                    var v20 uint64 = 18446744073709551615
                    v21 := GenS2{F0: true, F1: v20}
                    _ = v20
                    _ = v21
                }
            } else {
                var v22 int8 = -71
                _ = v22
            }
        }
        v23 := GenS2{F0: true, F1: 530}
        _ = v23
    } else {
        v24 := GenS3{F0: (p0 + p0), F1: 0}
        if ((((1.5 * 3.14159) / -2.25) * float64(float32(0.0))) >= 0.0) {
            p0 = p0
            if (120 > 127) {
                var v25 bool = (!false)
                fmt.Println(p0, p0, p1)
                fmt.Println(v25)
                _ = v25
            } else {
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 string = ""
                    var v28 uint32 = 42
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                var v29 uint16 = 2
                fmt.Println(p1, p0, v29)
                _ = v29
            }
            v30 := GenS3{F0: p0, F1: 2}
            _ = v30
        } else {
            switch "" {
            case "a":
                for fi31 := 0; fi31 < 18; fi31++ {
                    if false {
                        if (!(fi31 >= fi31)) {
                            v32 := GenS4{F0: "hello", F1: 1}
                            v33 := GenS5{F0: p1, F1: (false && (100 >= int8(p0))), F2: 17}
                            _ = v32
                            _ = v33
                        }
                        if (!(2 == 184)) {
                            p0 = 4294967295
                            v34 := GenS4{F0: "hello", F1: 90}
                            fmt.Println(p0, p2)
                            _ = v34
                        } else {
                            if true {
                                v35 := GenS3{F0: p0, F1: 172}
                                _ = v35
                            }
                        }
                        v36 := GenS6{F0: -128}
                        _ = v36
                    } else {
                        fmt.Println(p1)
                    }
                    var v37 int16 = 42
                    _ = fi31
                    _ = v37
                }
            default:
                v38 := GenS7{F0: 1}
                _ = v38
            }
        }
        _ = v24
    }
    fmt.Println(p0, p1)
    var v39 int = 124
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v39
    return "foo bar"
}

func main() {
    var v1 float64 = -2.25
    v2 := []int{32768}
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    v5 := []uint32{573, 2147483648, 561}
    var v6 uint16 = 65535
    fmt.Println(v6, v1)
    if false {
        delete(v4, 128)
        if true {
            v40 := GenS7{F0: -832}
            v6 = (v6 - v6)
            _ = v40
        } else {
            fmt.Println(v6, v6)
        }
        v1 = v1
    } else {
        v41 := GenS6{F0: 56}
        v6 = v6
        _ = v41
    }
    v42 := GenS6{F0: 42}
    var v43 uint16 = (v6 >> ((1) & 63))
    v1 = 0.1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 -= (1.5 / (v1 + float64((3.14159 + -2.25))))
    switch -470 {
    case -9223372036854775808:
        if true {
            fmt.Println(v6, v6)
            v6 = 75
        }
    case -2147483648:
        for fi44 := 0; fi44 < 13; fi44++ {
            v4[1] = ("\n\t" <= "hello")
            v1 = float64(float32(v1))
            if false {
                if (!true) {
                    if (9223372036854775807 <= 42) {
                        v45 := GenS3{F0: 479, F1: 127}
                        var v46 uint64 = 2
                        _ = v45
                        _ = v46
                    }
                    v47 := GenS6{F0: 127}
                    _ = v47
                } else {
                    v43 = (v6 << ((0) & 63))
                    v1 = v1
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v48 := GenS4{F0: ("" + "\n\t"), F1: 165}
                v43 = 32767
                _ = v48
            }
            _ = fi44
        }
    case 2147483647:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        if (true || (87 >= 0)) {
            for fi49 := 0; fi49 < 15; fi49++ {
                v50 := GenS8{F0: "a", F1: 65535, F2: 0}
                fmt.Println(fi49, fi49)
                _ = fi49
                _ = v50
            }
            for fi51 := 0; fi51 < 16; fi51++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi51
            }
        } else {
            var v52 int8 = -128
            v6 += (v43 << ((590) & 63))
            if false {
                var v53 int = 0
                v6 -= (v6 | (v6 | 42))
                _ = v53
            } else {
                if (42 > 42) {
                    v43 += v43
                    v6 += v43
                } else {
                    for fi54 := 0; fi54 < 5; fi54++ {
                        v6 -= 65535
                        v5 = append(v5, 2147483647)
                        _ = fi54
                    }
                    for fi55 := 0; fi55 < 20; fi55++ {
                        v56 := GenS9{F0: ("a" + "hello")}
                        for fi57 := 0; fi57 < 11; fi57++ {
                            var v58 uint = 42
                            v1 += 1.5
                            var v59 int = 42
                            var v60 uint64 = 1
                            if true {
                                fmt.Println(v58, v52, v59)
                            }
                            _ = fi57
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        for fi61 := 0; fi61 < 10; fi61++ {
                            if (65535 == 65535) {
                                v3[("foo bar" + "\n\t")] = "\n\t"
                                var v62 int8 = 106
                                fmt.Println(fi61, v52, v6)
                                _ = v62
                            }
                            if false {
                                switch "" {
                                case "a":
                                    v63 := GenS9{F0: "hello"}
                                    _ = v63
                                case "":
                                    if len(v5) > 0 {
                                        v5[0] = 4294967295
                                    }
                                default:
                                    v6 = v6
                                }
                                var v64 int64 = 1
                                if (32767 == -32768) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v1 -= (v1 * v1)
                                    switch 127 {
                                    case 617:
                                        var v65 float64 = (3.14159 - -2.25)
                                        _ = v65
                                    default:
                                        fmt.Println(v6, v1)
                                    }
                                }
                                _ = v64
                            } else {
                                var v66 uint64 = 535
                                if false {
                                    fmt.Println(v6)
                                }
                                fmt.Println(fi55)
                                _ = v66
                            }
                            if (false || (!true)) {
                                var v67 uint8 = 128
                                fmt.Println(fi55, fi61, fi55)
                                v68 := GenS6{F0: (v52 - (v52 * v52))}
                                switch "foo bar" {
                                case "foo bar":
                                    switch int(v6) {
                                    case 2147483647:
                                        if ("" != "a") {
                                            var v69 uint32 = 65535
                                            _ = v69
                                        }
                                    case -331:
                                        v52 = v52
                                    default:
                                        fmt.Println(v43)
                                    }
                                case "":
                                    v52 += v52
                                default:
                                    if (65535 != 2147483647) {
                                        var v70 float64 = v1
                                        v71 := GenS10{F0: 1, F1: uint64(v43), F2: 315}
                                        _ = v70
                                        _ = v71
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = 32767
                                        }
                                        var v72 uint64 = 185
                                        var v73 int = (-(fi61))
                                        _ = v72
                                        _ = v73
                                    }
                                }
                                _ = v67
                                _ = v68
                            }
                            _ = fi61
                        }
                        _ = fi55
                        _ = v56
                    }
                    for fi74 := 0; fi74 < 10; fi74++ {
                        fmt.Println(v1, fi74, v52)
                        if (!true) {
                            if len(v2) > 0 {
                                v2[0] = (fi74 + (fi74 & 128))
                            }
                            switch fi74 {
                            case 1:
                                if (int32(v52) > -224) {
                                    v6 = 0
                                    var v75 int16 = 32767
                                    var v76 uint64 = 1
                                    _ = v75
                                    _ = v76
                                } else {
                                    fmt.Println(v6)
                                    switch "\n\t" {
                                    case "\n\t":
                                        for fi77 := 0; fi77 < 10; fi77++ {
                                            v1 = 0.1
                                            _ = fi77
                                        }
                                    case "":
                                        v43 = v6
                                    case "a":
                                        var v78 int16 = 42
                                        _ = v78
                                    }
                                    v1 += float64((float32((v1 + v1)) + (1.5 - (((3.14159 - 3.14159) + (1e10 / 1.5)) - ((1.5 * (1.5 + 3.14159)) * 0.0)))))
                                }
                            case 127:
                                v52 = (v52 * v52)
                            case 894:
                                v6 = 113
                            }
                            _ = len(v3)
                        }
                        _ = fi74
                    }
                }
            }
            _ = v52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
