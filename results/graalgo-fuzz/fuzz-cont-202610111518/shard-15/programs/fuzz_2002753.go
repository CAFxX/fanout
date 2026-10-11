package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 string; F1 int16; F2 int }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 18446744073709551615
    var v3 uint32 = 1
    v4 := []int{0, 590, -1, 128}
    v5 := make(map[string]int)
    v6 := []int{9223372036854775807, 42, -286, -842}
    v7 := GenS1{F0: -1, F1: -161}
    v8 := GenS2{F0: -1}
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v2, v2)
    default:
        v2 = 871
    }
    fmt.Println(v2, v2)
    v2 += (v2 & (v2 * (v2 % ((303) | 1))))
    if true {
        switch (("a" + "a") + ("foo bar" + ("hello" + "\n\t"))) {
        case "foo bar":
            var v9 uint16 = 0
            _ = v9
        case "a":
            v10 := v1["foo bar"]
            _ = v10
        case "\n\t":
            var v11 int32 = -1
            _ = v11
        default:
            for fi12 := 0; fi12 < 1; fi12++ {
                v2 = (v2 << ((v2) & 63))
                _ = fi12
            }
        }
        v13 := GenS1{F0: 83, F1: -327}
        var v14 int32 = 2147483647
        _ = v13
        _ = v14
    }
    switch ("a" + (("\n\t" + "") + "foo bar")) {
    case "foo bar":
        switch ("foo bar" + "foo bar") {
        case "":
            if true {
                var v15 int64 = -177
                var v16 uint32 = (^(v3))
                if false {
                    if ("\n\t" < "a") {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            v18 := GenS3{F0: 789}
                            if false {
                                var v19 uint64 = 18446744073709551615
                                _ = v19
                            }
                            if ("\n\t" != ("\n\t" + "a")) {
                                fmt.Println(v3, fi17)
                            } else {
                                for fi20 := 0; fi20 < 12; fi20++ {
                                    switch (("a" + "foo bar") + ("hello" + "\n\t")) {
                                    case "a":
                                        v16 -= v3
                                    }
                                    _ = fi20
                                }
                            }
                            _ = fi17
                            _ = v18
                        }
                        v2 = 4294967295
                        for fi21 := 0; fi21 < 7; fi21++ {
                            v22 := GenS3{F0: 0}
                            _ = fi21
                            _ = v22
                        }
                    } else {
                        var v23 int64 = v15
                        for fi24 := 0; fi24 < 1; fi24++ {
                            var v25 int8 = -66
                            for fi26 := 0; fi26 < 13; fi26++ {
                                v15 += (v23 / ((v15) | 1))
                                var v27 uint = 4294967295
                                delete(v1, "a")
                                if len(v6) > 0 {
                                    v6[0] = (fi24 / ((int(v16)) | 1))
                                }
                                _ = fi26
                                _ = v27
                            }
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    }
                    v16 = v16
                    for fi28 := 0; fi28 < 13; fi28++ {
                        v4 = append(v4, 2147483647)
                        fmt.Println(v16)
                        _ = fi28
                    }
                } else {
                    if true {
                        switch -921 {
                        case -534:
                            switch ("hello" + "") {
                            case "hello":
                                var v29 uint32 = (v3 & 2)
                                _ = v29
                            case "":
                                v30 := GenS2{F0: (^(v15))}
                                _ = v30
                            default:
                                for fi31 := 0; fi31 < 13; fi31++ {
                                    for fi32 := 0; fi32 < 20; fi32++ {
                                        if ((0 != 127) || (!true)) {
                                            fmt.Println(fi32, v15)
                                            v16 = (v3 >> (((v2 + (v2 - 255))) & 63))
                                        } else {
                                            var v33 uint = v2
                                            _ = v33
                                        }
                                        _ = fi32
                                    }
                                    var v34 float32 = (1.5 + ((0.1 - (-2.25 + 3.14159)) - -2.25))
                                    _ = fi31
                                    _ = v34
                                }
                            }
                        case -37:
                            v2 += (^(v2))
                        case 32768:
                            v35 := GenS4{F0: 1, F1: (v3 / ((0) | 1)), F2: v16}
                            _ = v35
                        }
                        for fi36 := 0; fi36 < 15; fi36++ {
                            for fi37 := 0; fi37 < 18; fi37++ {
                                fmt.Println(v2)
                                fmt.Println(v2, v15, fi37)
                                v3 = (v16 | 4294967295)
                                _ = fi37
                            }
                            _ = fi36
                        }
                        for fi38 := 0; fi38 < 4; fi38++ {
                            v2 = (v2 << ((uint(v3)) & 63))
                            fmt.Println(v16)
                            v39 := GenS2{F0: 0}
                            _ = fi38
                            _ = v39
                        }
                    } else {
                        var v40 int = 514
                        fmt.Println(v2)
                        _ = v40
                    }
                }
                _ = v15
                _ = v16
            } else {
                if (255 == 255) {
                    if true {
                        if (49 == 1) {
                            v3 = v3
                            if (!(!true)) {
                                v41 := GenS3{F0: 1}
                                v3 = 0
                                delete(v1, (("hello" + (("\n\t" + "\n\t") + (("foo bar" + "hello") + ("hello" + "a")))) + "foo bar"))
                                _ = v41
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    switch "foo bar" {
                                    case "a":
                                        var v42 uint32 = (v3 | v3)
                                        _ = v42
                                    }
                                case "":
                                    if false {
                                        switch 349 {
                                        case 0:
                                            delete(v1, ("a" + ("\n\t" + ("" + ("\n\t" + ((("hello" + "hello") + ("a" + "\n\t")) + "foo bar"))))))
                                        }
                                    }
                                case "hello":
                                    v43 := v1["foo bar"]
                                    _ = v43
                                }
                                if ((((-270 != -636) || false) && (-9223372036854775808 <= -239)) && (((!true) && (("\n\t" + "foo bar") > ("" + "hello"))) && (1 < 800))) {
                                    delete(v5, "\n\t")
                                    switch "\n\t" {
                                    case "hello":
                                        var v44 float64 = (1e10 + (1e10 * 0.1))
                                        _ = v44
                                    case "\n\t":
                                        var v45 uint32 = (v3 ^ (v3 >> ((872) & 63)))
                                        _ = v45
                                    default:
                                        if false {
                                            var v46 float32 = 3.14159
                                            fmt.Println(v3)
                                            if false {
                                                v3 -= (^(v3))
                                                fmt.Println(v2)
                                            }
                                            _ = v46
                                        } else {
                                            delete(v1, ("hello" + "foo bar"))
                                        }
                                    }
                                }
                            }
                            switch 0 {
                            case -356:
                                var v47 uint8 = 128
                                _ = v47
                            case 1:
                                var v48 int = 1
                                _ = v48
                            case 192:
                                var v49 string = "\n\t"
                                _ = v49
                            default:
                                v50 := GenS4{F0: 0, F1: v3, F2: 2147483648}
                                _ = v50
                            }
                        }
                    }
                }
                if false {
                    v2 = (^(v2))
                } else {
                    var v51 bool = false
                    for fi52 := 0; fi52 < 12; fi52++ {
                        v51 = ((0.0 < 0.0) && (((((0.0 + ((float32(3.14159) / (0.1 / 3.14159)) + (float32((1.5 / 0.1)) - (-2.25 * (3.14159 - 0.0))))) - float32(-2.25)) + 1e10) < 1e10) || (584 < 2)))
                        v51 = false
                        v53 := v1[("hello" + "\n\t")]
                        _ = fi52
                        _ = v53
                    }
                    _ = v51
                }
                var v54 uint16 = 32767
                _ = v54
            }
        }
    case "a":
        fmt.Println(v2, v3)
    }
    for fi55 := 0; fi55 < 20; fi55++ {
        for fi56 := 0; fi56 < 9; fi56++ {
            v2 -= 18446744073709551615
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v3 += (v3 ^ v3)
            _ = fi56
        }
        for fi57 := 0; fi57 < 7; fi57++ {
            var v58 int64 = int64(fi55)
            var v59 float32 = 1.5
            if (124 <= 0) {
                v5["foo bar"] = (fi57 + int(v2))
                if true {
                    if (1 == 32767) {
                        v58 = (v58 / ((0) | 1))
                    } else {
                        v3 = 2
                        var v60 int64 = (v58 * (v58 * (v58 ^ (v58 ^ (^(v58))))))
                        _ = v60
                    }
                    v58 += v58
                    v58 -= v58
                }
            } else {
                fmt.Println(v59)
                fmt.Println(fi57)
                v61 := GenS5{F0: 509, F1: 255}
                _ = v61
            }
            _ = fi57
            _ = v58
            _ = v59
        }
        _ = fi55
    }
    if true {
        fmt.Println(v3, v3)
        if ((!(-1 < -1)) && true) {
            v5["foo bar"] = 0
            v3 = (v3 / ((v3) | 1))
            fmt.Println(v3, v3)
        }
        if ((false && (188 < 127)) || false) {
            switch -9223372036854775808 {
            case 42:
                v62 := GenS6{F0: 32767}
                _ = v62
            default:
                v3 += 4294967295
            }
            v63 := GenS4{F0: 0, F1: (v3 * v3), F2: v3}
            _ = v63
        }
    }
    fmt.Println(v3, v3)
    if ((uint8(v2) != 183) || false) {
        for fi64 := 0; fi64 < 18; fi64++ {
            v3 += (v3 | 446)
            _ = fi64
        }
        delete(v1, ("hello" + "foo bar"))
    } else {
        _ = len(v4)
        v65 := GenS7{F0: (^(v2)), F1: ("foo bar" + "a")}
        _ = v65
    }
    for fi66 := 0; fi66 < 17; fi66++ {
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v67 int = 9223372036854775807
            var v68 float32 = -2.25
            _ = v67
            _ = v68
        } else {
            v2 = v2
            v3 = 4294967295
        }
        _ = fi66
    }
    var v69 uint64 = 1
    if (!false) {
        fmt.Println(v69)
        var v70 uint64 = (v69 << ((65535) & 63))
        v2 += (v2 | (v2 - (v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))))
        _ = v70
    } else {
        fmt.Println(v69, v3, v69)
        for fi71 := 0; fi71 < 5; fi71++ {
            for fi72 := 0; fi72 < 6; fi72++ {
                v73 := GenS8{F0: "foo bar", F1: 42, F2: (fi71 + fi72)}
                var v74 uint64 = v69
                var v75 float64 = -2.25
                _ = fi72
                _ = v73
                _ = v74
                _ = v75
            }
            v3 = 247
            if true {
                var v76 int = fi71
                v77 := GenS9{F0: "\n\t"}
                _ = v76
                _ = v77
            } else {
                delete(v1, (("foo bar" + ("hello" + "foo bar")) + ("" + "foo bar")))
            }
            _ = fi71
        }
        v78 := GenS10{F0: -32768, F1: -128}
        _ = v78
    }
    var v79 bool = false
    _ = len(v4)
    fmt.Println(v69)
    v80 := v1[(("" + (("" + "a") + "foo bar")) + "")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v69:", v69)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v7
    _ = v79
    _ = v8
    _ = v80
}
