package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint32, p1 int64) string {
    var v7 string = "\n\t"
    var v8 uint32 = p0
    var v9 uint64 = 2
    fmt.Println(p0)
    v10 := []int{int(p0), 32768}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 uint16 = 738
    var v2 float32 = (3.14159 - -2.25)
    v3 := []int{-747, 758, 121}
    v4 := make(map[uint32]bool)
    var v5 int32 = 0
    var v6 float32 = 3.14159
    for fi11 := 0; fi11 < 11; fi11++ {
        switch "\n\t" {
        case "hello":
            if len(v3) > 0 {
                v3[0] = 42
            }
        }
        _ = fi11
    }
    if (127 != -128) {
        if (!(0.1 == 1.5)) {
            switch "" {
            case "hello":
                if (4294967295 == 42) {
                    v2 -= (v6 * ((v2 - ((0.0 - v6) - (float32((-2.25 * 0.1)) + float32(float64((3.14159 * (v2 / (float32(1.5) - (float32(0.0) * 1.5))))))))) * float32((-2.25 - 3.14159))))
                    for fi12 := 0; fi12 < 7; fi12++ {
                        if true {
                            var v13 int64 = 558
                            _ = v13
                        }
                        _ = fi12
                    }
                } else {
                    switch "a" {
                    case "foo bar":
                        for fi14 := 0; fi14 < 15; fi14++ {
                            if (false || (true || false)) {
                                for fi15 := 0; fi15 < 6; fi15++ {
                                    var v16 string = "\n\t"
                                    _ = fi15
                                    _ = v16
                                }
                                if false {
                                    if false {
                                        switch "" {
                                        case "\n\t":
                                            fmt.Println(v1, v5)
                                        case "a":
                                            fmt.Println(v6)
                                        }
                                    } else {
                                        if (v1 != (v1 ^ v1)) {
                                            var v17 uint16 = (v1 | 42)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v18 int8 = 58
                                            _ = v17
                                            _ = v18
                                        }
                                    }
                                    var v19 string = "\n\t"
                                    var v20 int16 = 32767
                                    _ = v19
                                    _ = v20
                                }
                                v21 := GenS1{F0: 127, F1: 2, F2: -1}
                                _ = v21
                            }
                            var v22 uint16 = (v1 - v1)
                            _ = fi14
                            _ = v22
                        }
                    case "\n\t":
                        fmt.Println(v1, v2, v6)
                    }
                    var v23 uint = 915
                    _ = v23
                }
            default:
                v3 = append(v3, 9223372036854775807)
            }
            v6 = 1.5
            for fi24 := 0; fi24 < 19; fi24++ {
                switch ("foo bar" + "") {
                case "a":
                    if (39 <= 0) {
                        v5 = v5
                        var v25 uint = 974
                        _ = v25
                    }
                default:
                    var v26 uint64 = 2
                    _ = v26
                }
                if false {
                    v27 := GenS2{F0: "\n\t", F1: 42}
                    var v28 uint64 = 2
                    for fi29 := 0; fi29 < 10; fi29++ {
                        var v30 uint64 = (^(v28))
                        _ = fi29
                        _ = v30
                    }
                    _ = v27
                    _ = v28
                } else {
                    v1 -= v1
                    for fi31 := 0; fi31 < 17; fi31++ {
                        if (("a" + "foo bar") == "hello") {
                            v2 -= float32(0.1)
                            if false {
                                delete(v4, 42)
                                var v32 int8 = -1
                                _ = v32
                            }
                            if true {
                                var v33 float64 = float64((-2.25 / 0.1))
                                var v34 uint64 = 943
                                var v35 int16 = 427
                                if len(v3) > 0 {
                                    v3[0] = fi31
                                }
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        }
                        switch -1 {
                        case 127:
                            var v36 string = ""
                            _ = v36
                        }
                        _ = fi31
                    }
                }
                var v37 int = fi24
                _ = fi24
                _ = v37
            }
        } else {
            v38 := GenS3{F0: false}
            _ = v38
        }
        var v39 uint16 = v1
        fmt.Println(v39, v2, v2)
        _ = v39
    }
    if (false && (true || ("" != "hello"))) {
        var v40 string = "a"
        v6 -= v6
        _ = v40
    }
    var v41 int16 = -1
    var v42 int16 = (v41 - (v41 / ((v41) | 1)))
    switch 4294967295 {
    case 964:
        switch int(v5) {
        case 9223372036854775807:
            if false {
                fmt.Println(v5, v6)
                var v43 int16 = 1000
                var v44 uint64 = 4294967295
                _ = v43
                _ = v44
            }
        }
    case -762:
        v2 += (v2 * -2.25)
    case 211:
        v45 := GenS3{F0: false}
        _ = v45
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        switch fi46 {
        case 127:
            var v47 int8 = 100
            _ = v47
        }
        var v48 uint64 = 18446744073709551615
        for fi49 := 0; fi49 < 2; fi49++ {
            v42 += (v42 ^ (v41 - v41))
            switch "hello" {
            case "":
                fmt.Println(v48, fi46)
            case "\n\t":
                v1 = v1
            }
            v50 := GenS3{F0: (int(v1) > int(v48))}
            var v51 int32 = 954
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = fi46
        _ = v48
    }
    if (v1 >= (^(v1))) {
        switch "" {
        case "":
            switch "hello" {
            case "foo bar":
                switch 127 {
                case -166:
                    if (1 != 32768) {
                        for fi52 := 0; fi52 < 3; fi52++ {
                            if true {
                                if (42 <= 42) {
                                    v53 := GenS1{F0: 219, F1: 2, F2: -963}
                                    var v54 uint8 = 17
                                    fmt.Println(v42, v6)
                                    _ = v53
                                    _ = v54
                                } else {
                                    fmt.Println(v42)
                                    for fi55 := 0; fi55 < 11; fi55++ {
                                        if (fi55 > fi52) {
                                            var v56 int64 = 9223372036854775807
                                            v57 := GenS1{F0: 42, F1: 65535, F2: v56}
                                            fmt.Println(v41)
                                            _ = v56
                                            _ = v57
                                        }
                                        v2 += (v2 - -2.25)
                                        switch "" {
                                        case "foo bar":
                                            fmt.Println(v2)
                                        case "":
                                            v42 += v41
                                        case "\n\t":
                                            var v58 uint16 = 2
                                            _ = v58
                                        default:
                                            var v59 int32 = 42
                                            _ = v59
                                        }
                                        _ = fi55
                                    }
                                    fmt.Println(v41, v2)
                                }
                                var v60 uint16 = v1
                                _ = v60
                            }
                            _ = fi52
                        }
                        for fi61 := 0; fi61 < 16; fi61++ {
                            if true {
                                fmt.Println(v6, v2)
                            }
                            var v62 int16 = (^(v42))
                            _ = fi61
                            _ = v62
                        }
                        fmt.Println(v41)
                    } else {
                        fmt.Println(v6, v5, v42)
                    }
                case 4294967295:
                    v63 := GenS2{F0: "hello", F1: 65535}
                    _ = v63
                case 794:
                    for fi64 := 0; fi64 < 17; fi64++ {
                        v65 := GenS4{F0: (v5 >> ((uint(v1)) & 63)), F1: 65535}
                        var v66 bool = (true && true)
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                default:
                    _ = len(v3)
                }
            }
        }
        v6 -= v2
    }
    switch 1 {
    case -2147483648:
        v2 -= v2
    default:
        if (0.1 == -2.25) {
            v67 := GenS4{F0: (v5 << ((0) & 63)), F1: 2}
            _ = v67
        } else {
            var v68 float64 = 0.1
            switch ("hello" + "\n\t") {
            case "":
                for fi69 := 0; fi69 < 6; fi69++ {
                    fmt.Println(v41)
                    if true {
                        fmt.Println(v2, v5, v5)
                        if false {
                            fmt.Println(v42, v2, v6)
                            var v70 uint32 = 614
                            for fi71 := 0; fi71 < 19; fi71++ {
                                var v72 float64 = 1.5
                                _ = fi71
                                _ = v72
                            }
                            _ = v70
                        } else {
                            switch ("" + ((("" + "a") + ("\n\t" + "hello")) + "hello")) {
                            case "foo bar":
                                v2 = 1e10
                            case "":
                                v6 += 0.1
                            default:
                                _ = len(v4)
                            }
                            var v73 float32 = v6
                            var v74 int = fi69
                            _ = v73
                            _ = v74
                        }
                        v3 = append(v3, -1)
                    }
                    _ = fi69
                }
            default:
                v1 = v1
            }
            v75 := GenS5{F0: (!true), F1: 127}
            _ = v68
            _ = v75
        }
    }
    var v76 float32 = (float32(0.0) * (((v6 + float32(1e10)) * v6) + float32(-2.25)))
    if ((^(v41)) == 32767) {
        v5 = (v5 * (v5 ^ (-(v5))))
        v77 := GenS6{F0: (v1 / ((v1) | 1))}
        if (float64((v2 / 3.14159)) <= (0.1 / 0.1)) {
            fmt.Println(v5, v2)
        }
        _ = v77
    }
    var v78 uint = 4294967295
    var v79 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v76
    _ = v78
    _ = v79
}
