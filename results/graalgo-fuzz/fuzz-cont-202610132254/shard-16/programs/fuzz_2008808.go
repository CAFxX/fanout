package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 uint32; F2 uint16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint32; F1 int64 }
type GenS12 struct { F0 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    var v3 int64 = 42
    v4 := make(map[string]uint32)
    v5 := make(map[int]string)
    v6 := []string{"foo bar", "hello"}
    var v7 int64 = v3
    v8 := GenS1{F0: 2}
    if (65535 >= 1) {
        for fi9 := 0; fi9 < 19; fi9++ {
            switch ("foo bar" + ((("" + "\n\t") + "hello") + "foo bar")) {
            case "\n\t":
                var v10 uint64 = 61
                _ = v10
            case "hello":
                v11 := GenS2{F0: 768, F1: 2, F2: 1}
                _ = v11
            case "a":
                if (128 >= 255) {
                    var v12 bool = (-2.25 == (float32(((1.5 + (1.5 + (1.5 * 0.1))) * (-2.25 * -2.25))) * 3.14159))
                    if true {
                        fmt.Println(v7, v12, v3)
                        v13 := GenS3{F0: "hello", F1: -66}
                        switch "hello" {
                        case "\n\t":
                            v14 := GenS4{F0: -722, F1: (fi9 * 9223372036854775807)}
                            _ = v14
                        case "a":
                            var v15 int8 = 100
                            _ = v15
                        case "foo bar":
                            if false {
                                if (((0.1 - ((1.5 / 3.14159) * 0.0)) * (((3.14159 + 0.0) - (float64(0.0) * 3.14159)) / 1.0)) >= 0.0) {
                                    var v16 uint = 65535
                                    var v17 float64 = float64(((float32(float64(((float32(((1e10 + 3.14159) + (0.1 / 0.1))) * (3.14159 + -2.25)) - (1e10 + -2.25)))) + (1e10 * float32(3.14159))) - 3.14159))
                                    switch fi9 {
                                    case 0:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    case 32767:
                                        for fi18 := 0; fi18 < 1; fi18++ {
                                            var v19 float64 = v17
                                            fmt.Println(v16, v19)
                                            var v20 int8 = 127
                                            _ = fi18
                                            _ = v19
                                            _ = v20
                                        }
                                    default:
                                        switch (^(fi9)) {
                                        case 0:
                                            switch fi9 {
                                            case 864:
                                                var v21 int64 = (v3 | -1)
                                                _ = v21
                                            }
                                        case -1:
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v22 float64 = (((float64(0.1) - -2.25) - (float64(((-2.25 + -2.25) - float32(v17))) * 3.14159)) - (v17 / v17))
                                                _ = v22
                                            case "a":
                                                switch 2147483647 {
                                                case 32768:
                                                    fmt.Println(v12, fi9, v3)
                                                default:
                                                    v7 = v3
                                                }
                                            }
                                        case 128:
                                            v23 := GenS5{F0: 704}
                                            _ = v23
                                        default:
                                            var v24 uint64 = 685
                                            _ = v24
                                        }
                                    }
                                    _ = v16
                                    _ = v17
                                }
                                var v25 int = fi9
                                v3 = 1
                                _ = v25
                            } else {
                                if false {
                                    v12 = (v12 || ((!v12) || (!true)))
                                    if true {
                                        var v26 int64 = 1
                                        _ = v26
                                    } else {
                                        var v27 bool = true
                                        if (v27 || v12) {
                                            switch ("a" + "\n\t") {
                                            case "hello":
                                                v6 = append(v6, "a")
                                            case "\n\t":
                                                v7 = (v3 ^ 0)
                                            case "a":
                                                v6 = append(v6, "\n\t")
                                            default:
                                                if (945 >= 1) {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v28 := GenS6{F0: 1, F1: 127}
                                                        _ = v28
                                                    case "hello":
                                                        var v29 int8 = -92
                                                        _ = v29
                                                    case "":
                                                        var v30 uint8 = 2
                                                        _ = v30
                                                    }
                                                    var v31 uint32 = 1
                                                    var v32 string = ("" + (("\n\t" + "\n\t") + ("hello" + "a")))
                                                    _ = v31
                                                    _ = v32
                                                } else {
                                                    v12 = (222 == 191)
                                                }
                                            }
                                            for fi33 := 0; fi33 < 11; fi33++ {
                                                var v34 int32 = 318
                                                _ = len(v6)
                                                var v35 uint = 78
                                                _ = fi33
                                                _ = v34
                                                _ = v35
                                            }
                                        } else {
                                            switch ("foo bar" + (("" + ("a" + "foo bar")) + (("a" + "\n\t") + "foo bar"))) {
                                            case "foo bar":
                                                v12 = v27
                                            case "\n\t":
                                                _ = len(v6)
                                            case "":
                                                switch -955 {
                                                case -9223372036854775808:
                                                    v36 := GenS4{F0: (v7 & v3), F1: fi9}
                                                    _ = v36
                                                case 127:
                                                    delete(v2, "a")
                                                case 479:
                                                    for fi37 := 0; fi37 < 10; fi37++ {
                                                        v38 := GenS7{F0: 2}
                                                        if ((v7 << ((18446744073709551615) & 63)) > v3) {
                                                            var v39 bool = (255 < 343)
                                                            fmt.Println(fi9, v7, v39)
                                                            v40 := GenS2{F0: 4294967295, F1: 32767, F2: 702}
                                                            _ = v39
                                                            _ = v40
                                                        } else {
                                                            var v41 bool = ((fi9 | (^(fi37))) != 32767)
                                                            v27 = false
                                                            _ = len(v5)
                                                            _ = v41
                                                        }
                                                        _ = fi37
                                                        _ = v38
                                                    }
                                                }
                                            }
                                        }
                                        _ = v27
                                    }
                                }
                                v12 = (v12 && v12)
                                v7 = 172
                            }
                        default:
                            switch "hello" {
                            case "hello":
                                var v42 uint8 = 244
                                _ = v42
                            case "a":
                                var v43 int32 = 1
                                _ = v43
                            case "":
                                delete(v2, "")
                            }
                        }
                        _ = v13
                    } else {
                        fmt.Println(fi9, v7)
                        v7 = (v7 - v3)
                    }
                    if (99 != 2147483648) {
                        var v44 bool = true
                        v12 = (v12 || (((((1.5 - float32(1.5)) * float32((0.0 - (1e10 * 0.1)))) > -2.25) && v12) && (false && (-32768 <= -32768))))
                        _ = v44
                    }
                    _ = v12
                } else {
                    var v45 int = fi9
                    for fi46 := 0; fi46 < 17; fi46++ {
                        var v47 float64 = ((float64(float32((1e10 - 1e10))) / 0.1) / 0.1)
                        switch 0 {
                        case 42:
                            fmt.Println(v7, fi9)
                        case 9223372036854775807:
                            v48 := GenS6{F0: 255, F1: 196}
                            _ = v48
                        case 127:
                            var v49 uint = 2
                            _ = v49
                        default:
                            v2["hello"] = (0 < 0)
                        }
                        delete(v4, (("\n\t" + "") + ("\n\t" + "hello")))
                        _ = fi46
                        _ = v47
                    }
                    var v50 float64 = (3.14159 / (0.1 / -2.25))
                    _ = v45
                    _ = v50
                }
            default:
                switch 32767 {
                case 1:
                    for fi51 := 0; fi51 < 4; fi51++ {
                        var v52 float64 = 0.1
                        _ = fi51
                        _ = v52
                    }
                case -1:
                    v3 = v3
                default:
                    v53 := GenS8{F0: 1}
                    _ = v53
                }
            }
            _ = fi9
        }
        if (32768 >= 32767) {
            _ = len(v6)
            var v54 uint8 = 2
            if (-480 == v3) {
                for fi55 := 0; fi55 < 8; fi55++ {
                    var v56 uint16 = 65535
                    v56 -= (v56 - v56)
                    _ = fi55
                    _ = v56
                }
                v57 := GenS6{F0: 255, F1: (v54 - 72)}
                _ = v57
            } else {
                switch 2147483647 {
                case 1:
                    v58 := GenS3{F0: "\n\t", F1: 127}
                    _ = v58
                default:
                    _ = len(v6)
                }
            }
            _ = v54
        }
        fmt.Println(v3, v3)
    }
    v59 := GenS1{F0: 128}
    var v60 uint = 735
    var v61 uint64 = 283
    var v62 uint16 = 65535
    for fi63 := 0; fi63 < 13; fi63++ {
        var v64 uint16 = 255
        _ = fi63
        _ = v64
    }
    if (128 == 0) {
        v62 -= (v62 * v62)
        v65 := GenS9{F0: ((("hello" + ("a" + ("hello" + "\n\t"))) + "foo bar") + "\n\t"), F1: 42, F2: v62}
        _ = v65
    }
    v66 := GenS10{F0: v62}
    v67 := GenS11{F0: 42, F1: (v3 * (-(v3)))}
    fmt.Println(v61, v60)
    switch "foo bar" {
    case "foo bar":
        var v68 int32 = -2147483648
        _ = v68
    case "\n\t":
        v69 := GenS12{F0: (v3 / ((v7) | 1))}
        _ = v69
    case "":
        fmt.Println(v61, v3, v3)
    }
    v3 -= v7
    var v70 float32 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v66
    _ = v67
    _ = v7
    _ = v70
    _ = v8
}
