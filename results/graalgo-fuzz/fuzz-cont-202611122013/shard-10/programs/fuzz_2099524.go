package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int32; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool }
func genFn0(p0 int32, p1 uint8, p2 uint8) int32 {
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            switch (fi5 + fi6) {
            case 2147483647:
                var v7 float64 = 1.5
                _ = v7
            case 4294967295:
                switch fi5 {
                case -635:
                    var v8 uint16 = 255
                    _ = v8
                default:
                    if (true || ((65535 > 2) && false)) {
                        if (-849 >= p0) {
                            var v9 int8 = -32
                            _ = v9
                        }
                        p0 -= (^(p0))
                        var v10 int64 = -293
                        _ = v10
                    }
                }
            default:
                v11 := GenS1{F0: (("\n\t" + "") + "foo bar")}
                p0 += -382
                v12 := GenS1{F0: ((("a" + "\n\t") + ("a" + "")) + "hello")}
                _ = v11
                _ = v12
            }
            _ = fi6
        }
        v13 := []int{(^(fi5)), 32768, fi5}
        if true {
            switch -555 {
            case 32767:
                p1 -= p1
            case -9223372036854775808:
                p1 -= p1
            case -2147483648:
                for fi14 := 0; fi14 < 17; fi14++ {
                    if (1e10 > 0.1) {
                        p2 += p1
                        var v15 float32 = float32(3.14159)
                        if true {
                            p0 -= (p0 - (p0 | 294))
                            if false {
                                var v16 string = "a"
                                fmt.Println(p2, fi5)
                                var v17 uint16 = uint16(p2)
                                _ = v16
                                _ = v17
                            } else {
                                v18 := GenS1{F0: ("" + "a")}
                                if (fi14 >= (fi14 >> ((255) & 63))) {
                                    v15 += ((float32(((-2.25 * (float64(1.5) * (0.1 * (1e10 + 0.1)))) - float64((v15 - (((0.1 - 3.14159) - v15) + (float32((1e10 + -2.25)) + (1e10 / ((v15 + -2.25) / v15)))))))) * 0.1) / (((float32((1.5 * 3.14159)) * (1.5 / -2.25)) / 1e10) + (((v15 - 0.0) * v15) / ((float32(0.0) / v15) - -2.25))))
                                }
                                switch (-(fi14)) {
                                case 32767:
                                    v19 := GenS1{F0: ("hello" + "\n\t")}
                                    _ = v19
                                case 32768:
                                    fmt.Println(p1, v15, fi14)
                                    v20 := GenS1{F0: "a"}
                                    _ = v20
                                case 1:
                                    var v21 uint = 171
                                    _ = v21
                                }
                                _ = v18
                            }
                            if (!true) {
                                v13 = append(v13, (fi5 ^ (fi14 | (-(fi5)))))
                                v22 := GenS1{F0: "a"}
                                var v23 int8 = 1
                                _ = v22
                                _ = v23
                            } else {
                                if true {
                                    v24 := GenS1{F0: "hello"}
                                    _ = v24
                                } else {
                                    switch fi5 {
                                    case 2147483647:
                                        var v25 int64 = 42
                                        _ = v25
                                    case 951:
                                        if ((p0 * (p0 & (-(p0)))) > p0) {
                                            var v26 string = "hello"
                                            _ = v26
                                        }
                                    }
                                    var v27 float32 = float32(float64(v15))
                                    _ = v27
                                }
                            }
                            if ((!(!(65535 >= 926))) || (((("foo bar" + "a") + (("" + ("a" + ((("foo bar" + "foo bar") + "foo bar") + "a"))) + "foo bar")) + "\n\t") <= "hello")) {
                                if (1 == 1) {
                                    var v28 uint = 18446744073709551615
                                    _ = len(v13)
                                    _ = v28
                                }
                                fmt.Println(v15, p0)
                            }
                        }
                        _ = v15
                    } else {
                        if true {
                            switch 727 {
                            case 32768:
                                p1 = p2
                            case 9223372036854775807:
                                var v29 uint64 = 2
                                _ = v29
                            case 2147483647:
                                var v30 int = fi14
                                _ = v30
                            }
                        }
                        var v31 float64 = (3.14159 - (1e10 * float64((((float32((1e10 - 1e10)) - float32(float64(((0.0 * 1e10) + (0.1 + ((-2.25 + ((0.1 / (-2.25 + (3.14159 + 3.14159))) - ((3.14159 - 0.0) * (1e10 * ((1e10 * -2.25) + 1.5))))) * (1e10 / float32(1.5)))))))) * ((((0.0 - (float32((0.1 - 3.14159)) / 1e10)) + float32(0.0)) + -2.25) - -2.25)) - 1e10))))
                        _ = v31
                    }
                    _ = fi14
                }
            }
        }
        _ = fi5
        _ = v13
    }
    var v32 uint8 = (p2 >> ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return 1
}

func main() {
    var v1 float32 = -2.25
    var v2 int32 = -551
    v3 := make(map[string]int)
    v4 := []string{("" + "a"), "foo bar", ""}
    fmt.Println(v2)
    if false {
        fmt.Println(v1, v1)
        var v33 int32 = v2
        v2 = v2
        _ = v33
    }
    v1 = ((v1 / -2.25) - 1e10)
    var v34 int8 = 100
    if (116 != 246) {
        var v35 int16 = 696
        _ = v35
    }
    v2 += (v2 ^ v2)
    fmt.Println(v34)
    fmt.Println(v1, v2)
    for fi36 := 0; fi36 < 15; fi36++ {
        v34 -= (v34 - (v34 - v34))
        _ = fi36
    }
    var v37 int32 = (v2 / ((42) | 1))
    v1 = v1
    for fi38 := 0; fi38 < 15; fi38++ {
        if true {
            fmt.Println(v34, v1, fi38)
            if (!(true || (false || (genFn0(-1, 42, 1) != genFn0(-625, 252, 199))))) {
                var v39 int64 = -1
                v40 := GenS2{F0: true}
                _ = len(v4)
                _ = v39
                _ = v40
            }
        } else {
            var v41 int64 = 42
            _ = v41
        }
        v42 := GenS1{F0: "hello"}
        _ = fi38
        _ = v42
    }
    if (32767 > 42) {
        for fi43 := 0; fi43 < 6; fi43++ {
            v44 := v3[(("foo bar" + "") + "hello")]
            var v45 uint = 0
            var v46 bool = (false || (false && true))
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
        v47 := GenS3{F0: "a", F1: genFn0(v2, 127, 255), F2: 2147483647}
        if ("hello" <= ("foo bar" + "hello")) {
            switch "hello" {
            case "\n\t":
                v34 = v34
            case "":
                v48 := GenS1{F0: "a"}
                _ = v48
            case "foo bar":
                v34 -= (v34 ^ v34)
            }
        } else {
            switch "\n\t" {
            case "a":
                var v49 int = 2147483647
                _ = v49
            case "hello":
                var v50 int64 = 0
                _ = v50
            }
            if (!false) {
                v1 += v1
            }
        }
        _ = v47
    } else {
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = len(v3)
        } else {
            var v51 int64 = 9223372036854775807
            for fi52 := 0; fi52 < 7; fi52++ {
                var v53 uint64 = 18446744073709551615
                var v54 uint64 = 4294967295
                _ = fi52
                _ = v53
                _ = v54
            }
            _ = v51
        }
        for fi55 := 0; fi55 < 7; fi55++ {
            var v56 int8 = (v34 % (((v34 / (((v34 >> ((1) & 63))) | 1))) | 1))
            if true {
                v57 := GenS2{F0: true}
                _ = v57
            }
            v58 := GenS4{F0: 42}
            _ = fi55
            _ = v56
            _ = v58
        }
    }
    switch ("a" + "foo bar") {
    case "foo bar":
        if true {
            for fi59 := 0; fi59 < 3; fi59++ {
                v1 += 1e10
                fmt.Println(v34)
                var v60 int64 = 9223372036854775807
                _ = fi59
                _ = v60
            }
        } else {
            for fi61 := 0; fi61 < 3; fi61++ {
                switch "" {
                case "":
                    var v62 uint16 = 0
                    _ = v62
                }
                v63 := GenS3{F0: "a", F1: (v37 + (v2 + v37)), F2: (-(fi61))}
                _ = fi61
                _ = v63
            }
        }
    }
    if (((v2 / ((v37) | 1)) <= v2) || (v2 < -1)) {
        v4 = append(v4, "\n\t")
        var v64 int8 = (v34 & v34)
        _ = v64
    } else {
        v2 += 2147483647
        v65 := GenS5{F0: 255, F1: (v34 >> ((0) & 63))}
        _ = v65
    }
    for fi66 := 0; fi66 < 1; fi66++ {
        switch fi66 {
        case 127:
            fmt.Println(v1)
        case -9223372036854775808:
            var v67 bool = false
            _ = v67
        case 9223372036854775807:
            switch (fi66 ^ (fi66 | (fi66 << ((0) & 63)))) {
            case 127:
                for fi68 := 0; fi68 < 3; fi68++ {
                    var v69 uint8 = 250
                    v2 = (v37 / ((genFn0(42, (^(v69)), 64)) | 1))
                    var v70 bool = (!(((!(("foo bar" + "a") == "")) || (!true)) || (-128 == 127)))
                    _ = fi68
                    _ = v69
                    _ = v70
                }
            case -534:
                v71 := GenS6{F0: -32768}
                _ = v71
            case -9223372036854775808:
                if (v34 != v34) {
                    v34 = (v34 | (v34 & 0))
                    for fi72 := 0; fi72 < 7; fi72++ {
                        v1 -= v1
                        if true {
                            var v73 bool = (123 < 255)
                            _ = v73
                        }
                        _ = fi72
                    }
                }
            default:
                delete(v3, (("\n\t" + "") + "a"))
            }
        default:
            v74 := v3["foo bar"]
            _ = v74
        }
        v2 = v2
        _ = fi66
    }
    for fi75 := 0; fi75 < 15; fi75++ {
        if true {
            var v76 uint = 198
            _ = v76
        } else {
            switch fi75 {
            case 127:
                switch ("foo bar" + ((("foo bar" + "hello") + "hello") + "hello")) {
                case "hello":
                    v77 := GenS1{F0: "hello"}
                    _ = v77
                case "foo bar":
                    fmt.Println(v34)
                default:
                    v37 = v37
                }
            case 128:
                v34 = (v34 % ((v34) | 1))
            case 558:
                var v78 float32 = (v1 * 0.1)
                _ = v78
            }
            for fi79 := 0; fi79 < 11; fi79++ {
                var v80 int32 = v2
                v81 := GenS2{F0: (!false)}
                if false {
                    v82 := GenS7{F0: true}
                    v83 := GenS6{F0: 487}
                    switch -1 {
                    case -1:
                        v84 := GenS4{F0: 255}
                        _ = v84
                    }
                    _ = v82
                    _ = v83
                } else {
                    var v85 uint64 = 2
                    v2 += int32(fi79)
                    _ = v85
                }
                _ = fi79
                _ = v80
                _ = v81
            }
            _ = len(v4)
        }
        var v86 string = ""
        for fi87 := 0; fi87 < 8; fi87++ {
            var v88 string = (v86 + (v86 + "foo bar"))
            _ = fi87
            _ = v88
        }
        _ = fi75
        _ = v86
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v4
}
