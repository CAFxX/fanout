package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int32; F2 uint }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"a", "a", "foo bar", "a"}
    v3 := make(map[string]uint32)
    v4 := []string{"hello", "", ("\n\t" + ((("\n\t" + "a") + "") + "foo bar"))}
    switch "foo bar" {
    case "":
        var v5 int = -1
        _ = v5
    case "foo bar":
        if true {
            if true {
                switch 4294967295 {
                case 9223372036854775807:
                    var v6 bool = ((false || true) || ((893 == 205) && (0.0 >= 0.0)))
                    _ = v6
                case 799:
                    if true {
                        for fi7 := 0; fi7 < 8; fi7++ {
                            var v8 int8 = -1
                            _ = fi7
                            _ = v8
                        }
                        var v9 uint16 = 65535
                        _ = v9
                    } else {
                        switch "foo bar" {
                        case "hello":
                            v10 := GenS1{F0: "a", F1: 65535}
                            _ = v10
                        default:
                            v11 := GenS1{F0: ((("hello" + "a") + "foo bar") + ""), F1: 967}
                            _ = v11
                        }
                        v12 := GenS2{F0: 433, F1: -470}
                        if true {
                            var v13 uint16 = 559
                            fmt.Println(v13)
                            _ = v13
                        }
                        _ = v12
                    }
                case -9223372036854775808:
                    var v14 uint32 = 4294967295
                    _ = v14
                default:
                    var v15 float32 = (1e10 + -2.25)
                    _ = v15
                }
                _ = len(v1)
            } else {
                if true {
                    if true {
                        var v16 float32 = 1.5
                        if true {
                            v17 := GenS1{F0: "a", F1: 0}
                            switch "foo bar" {
                            case "a":
                                for fi18 := 0; fi18 < 6; fi18++ {
                                    var v19 uint64 = 844
                                    _ = fi18
                                    _ = v19
                                }
                            default:
                                var v20 uint = 255
                                _ = v20
                            }
                            _ = v17
                        } else {
                            if (!(true && false)) {
                                v16 = float32(((0.1 - 0.0) + (0.1 + 1.5)))
                                fmt.Println(v16)
                                if (((true && false) && true) || (4294967295 < 721)) {
                                    v16 -= 3.14159
                                } else {
                                    v21 := GenS2{F0: 1, F1: 313}
                                    var v22 uint = 4294967295
                                    _ = v21
                                    _ = v22
                                }
                            }
                            if (!("a" > "\n\t")) {
                                var v23 uint16 = 754
                                switch "hello" {
                                case "foo bar":
                                    for fi24 := 0; fi24 < 13; fi24++ {
                                        for fi25 := 0; fi25 < 17; fi25++ {
                                            v23 = v23
                                            var v26 string = "\n\t"
                                            fmt.Println(v26)
                                            _ = fi25
                                            _ = v26
                                        }
                                        if true {
                                            v23 -= v23
                                            switch "foo bar" {
                                            case "a":
                                                var v27 uint32 = 65535
                                                _ = v27
                                            default:
                                                for fi28 := 0; fi28 < 11; fi28++ {
                                                    v29 := GenS2{F0: 2, F1: fi28}
                                                    v23 = 255
                                                    if (!(-1 != -1)) {
                                                        v30 := GenS1{F0: "a", F1: 4294967295}
                                                        if false {
                                                            v16 += (0.0 - 0.0)
                                                            var v31 float64 = float64(v16)
                                                            var v32 uint = 18446744073709551615
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                        if len(v4) > 0 {
                                                            v4[0] = "hello"
                                                        }
                                                        _ = v30
                                                    } else {
                                                        v23 -= (v23 + 0)
                                                        var v33 uint = 2
                                                        _ = v33
                                                    }
                                                    _ = fi28
                                                    _ = v29
                                                }
                                            }
                                            switch "a" {
                                            case "":
                                                for fi34 := 0; fi34 < 12; fi34++ {
                                                    var v35 bool = false
                                                    v35 = ((!(true || (2147483647 != int32(fi34)))) && (v35 || v35))
                                                    var v36 int8 = 0
                                                    _ = fi34
                                                    _ = v35
                                                    _ = v36
                                                }
                                            default:
                                                var v37 int32 = 851
                                                _ = v37
                                            }
                                        }
                                        v23 -= 255
                                        _ = fi24
                                    }
                                case "\n\t":
                                    var v38 uint16 = 760
                                    _ = v38
                                default:
                                    switch "foo bar" {
                                    case "hello":
                                        delete(v1, 76)
                                    }
                                }
                                if true {
                                    if (!(false || (!(v16 != (0.0 * 1e10))))) {
                                        var v39 string = "hello"
                                        _ = v39
                                    }
                                    v16 -= 1e10
                                    v16 += v16
                                } else {
                                    var v40 uint = 255
                                    _ = v40
                                }
                                _ = v23
                            } else {
                                if true {
                                    if (true && false) {
                                        v16 -= (0.0 / 1.5)
                                        var v41 int8 = 42
                                        _ = v41
                                    } else {
                                        if (0 > -2147483648) {
                                            v16 += (((float32(((0.0 * 0.0) - 3.14159)) + 1.5) / ((1e10 * (0.0 * (3.14159 + v16))) + v16)) / v16)
                                        } else {
                                            v16 = 0.0
                                        }
                                        fmt.Println(v16)
                                        for fi42 := 0; fi42 < 13; fi42++ {
                                            for fi43 := 0; fi43 < 4; fi43++ {
                                                var v44 float32 = float32((-2.25 / 1.0))
                                                _ = fi43
                                                _ = v44
                                            }
                                            fmt.Println(fi42)
                                            switch fi42 {
                                            case 3:
                                                v16 += ((((v16 / v16) - v16) / (((v16 + 3.14159) / (v16 - v16)) - float32(3.14159))) / (v16 * 1.5))
                                            case -1:
                                                v16 += 1.5
                                            case 0:
                                                if false {
                                                    switch 477 {
                                                    case 0:
                                                        v45 := GenS3{F0: 18446744073709551615, F1: 127}
                                                        _ = v45
                                                    case 42:
                                                        if (18446744073709551615 < 881) {
                                                            if false {
                                                                var v46 uint32 = 0
                                                                for fi47 := 0; fi47 < 6; fi47++ {
                                                                    if (false || (435 > 1)) {
                                                                        v48 := GenS4{F0: 42, F1: 1, F2: 78}
                                                                        _ = v48
                                                                    } else {
                                                                        var v49 uint32 = 142
                                                                        fmt.Println(fi47, v46)
                                                                        _ = v49
                                                                    }
                                                                    if true {
                                                                        v50 := GenS1{F0: "foo bar", F1: 255}
                                                                        fmt.Println(v46)
                                                                        fmt.Println(fi47, v46, fi47)
                                                                        _ = v50
                                                                    }
                                                                    _ = fi47
                                                                }
                                                                var v51 uint = 51
                                                                _ = v46
                                                                _ = v51
                                                            }
                                                        }
                                                    case -1:
                                                        if (v16 > v16) {
                                                            v16 -= 0.0
                                                            v16 = 1.5
                                                        } else {
                                                            fmt.Println(fi42, v16)
                                                            v16 = v16
                                                            v52 := GenS3{F0: 42, F1: 127}
                                                            _ = v52
                                                        }
                                                    }
                                                    v16 = (float32(3.14159) - v16)
                                                    fmt.Println(fi42)
                                                }
                                            default:
                                                var v53 string = "hello"
                                                _ = v53
                                            }
                                            _ = fi42
                                        }
                                    }
                                }
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                                if (!false) {
                                    var v54 int32 = -1
                                    _ = v54
                                } else {
                                    v16 = v16
                                    for fi55 := 0; fi55 < 20; fi55++ {
                                        v16 -= (0.1 + -2.25)
                                        v16 -= v16
                                        _ = fi55
                                    }
                                }
                            }
                        }
                        v16 = (float32(-2.25) / (v16 / 1.0))
                        _ = v16
                    } else {
                        for fi56 := 0; fi56 < 13; fi56++ {
                            v57 := GenS5{F0: 2, F1: -32768}
                            _ = fi56
                            _ = v57
                        }
                        var v58 uint8 = 127
                        var v59 int = 437
                        _ = v58
                        _ = v59
                    }
                    var v60 uint16 = 365
                    var v61 uint = 1
                    _ = v60
                    _ = v61
                }
                if (623 == 524) {
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    for fi62 := 0; fi62 < 1; fi62++ {
                        if (9223372036854775807 >= 0) {
                            switch (fi62 + (fi62 << ((18446744073709551615) & 63))) {
                            case 876:
                                fmt.Println(fi62)
                            case -566:
                                var v63 float32 = ((-2.25 * (((((0.0 + (0.1 / (3.14159 + 0.0))) / 1e10) + ((float32(0.1) - ((1.5 + float32(-2.25)) - 3.14159)) - 1e10)) - -2.25) * 1e10)) / (3.14159 + 0.0))
                                _ = v63
                            default:
                                fmt.Println(fi62)
                            }
                        }
                        var v64 int8 = -1
                        _ = fi62
                        _ = v64
                    }
                }
                var v65 int64 = -1
                _ = v65
            }
        } else {
            var v66 int = 4294967295
            for fi67 := 0; fi67 < 3; fi67++ {
                var v68 float32 = 0.1
                for fi69 := 0; fi69 < 9; fi69++ {
                    switch v66 {
                    case 757:
                        v68 = v68
                    }
                    switch (v66 - (^(fi67))) {
                    case 445:
                        v68 += (((v68 * (v68 * float32(1e10))) + float32((0.1 * float64(-2.25)))) - v68)
                    case -9223372036854775808:
                        fmt.Println(fi69, v66, v68)
                        v70 := GenS6{F0: 42}
                        var v71 float64 = 3.14159
                        _ = v70
                        _ = v71
                    default:
                        v66 = -500
                    }
                    fmt.Println(fi69, v68)
                    _ = fi69
                }
                _ = fi67
                _ = v68
            }
            var v72 int64 = 0
            _ = v66
            _ = v72
        }
    }
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    for fi73 := 0; fi73 < 2; fi73++ {
        var v74 string = "foo bar"
        v74 = ((v74 + v74) + v74)
        _ = fi73
        _ = v74
    }
    if true {
        var v75 int16 = -602
        _ = v75
    }
    var v76 int16 = -978
    var v77 uint = 4294967295
    switch (("hello" + "a") + "foo bar") {
    case "\n\t":
        var v78 bool = (0 != (v77 + 844))
        _ = v78
    case "foo bar":
        var v79 int16 = v76
        _ = v79
    case "":
        var v80 float64 = -2.25
        _ = v80
    }
    for fi81 := 0; fi81 < 1; fi81++ {
        v76 = (v76 << (((v77 ^ 0)) & 63))
        var v82 bool = true
        _ = fi81
        _ = v82
    }
    if ((1e10 < -2.25) && false) {
        if ("\n\t" > ("" + ("\n\t" + "a"))) {
            for fi83 := 0; fi83 < 13; fi83++ {
                v77 = (v77 % ((42) | 1))
                _ = fi83
            }
            switch 225 {
            case -262:
                v3[""] = 363
            }
        }
        v84 := v3[("hello" + ("foo bar" + "foo bar"))]
        _ = v84
    } else {
        _ = len(v3)
    }
    if (42 <= 127) {
        delete(v3, "foo bar")
        var v85 int = 2147483647
        for fi86 := 0; fi86 < 13; fi86++ {
            var v87 float64 = 1e10
            if false {
                if true {
                    if (1 < 1) {
                        for fi88 := 0; fi88 < 16; fi88++ {
                            v77 -= (v77 + 0)
                            v87 += (float64(0.0) / float64((((((1.5 * 1.5) * ((3.14159 / -2.25) * (-2.25 * 0.1))) - 3.14159) * ((3.14159 - (((0.1 - (((-2.25 * float32(0.0)) + (-2.25 / (0.1 - ((-2.25 - 0.1) * 3.14159)))) * 3.14159)) * (0.1 + 1.5)) * float32(0.0))) - (0.1 + (1.5 * -2.25)))) + float32((v87 - float64(0.0))))))
                            _ = fi88
                        }
                    } else {
                        v89 := GenS5{F0: 255, F1: -248}
                        if false {
                            if false {
                                v87 -= v87
                                for fi90 := 0; fi90 < 1; fi90++ {
                                    if true {
                                        v91 := GenS4{F0: 649, F1: 0, F2: 1}
                                        _ = v91
                                    }
                                    _ = fi90
                                }
                            }
                        }
                        for fi92 := 0; fi92 < 2; fi92++ {
                            v87 = ((float64(-2.25) + v87) / ((((v87 / v87) * v87) * v87) - v87))
                            if true {
                                v93 := GenS7{F0: -1}
                                v94 := GenS7{F0: 0}
                                _ = v93
                                _ = v94
                            }
                            _ = fi92
                        }
                        _ = v89
                    }
                    v4 = append(v4, "a")
                }
                _ = len(v1)
                for fi95 := 0; fi95 < 9; fi95++ {
                    var v96 uint = 18446744073709551615
                    _ = fi95
                    _ = v96
                }
            } else {
                var v97 float64 = (-2.25 - (v87 * ((((1.5 - (float64(0.0) + v87)) + v87) + (v87 / float64(((1e10 / -2.25) - -2.25)))) - 0.1)))
                _ = v97
            }
            var v98 uint16 = 255
            _ = fi86
            _ = v87
            _ = v98
        }
        _ = v85
    } else {
        for fi99 := 0; fi99 < 16; fi99++ {
            switch (fi99 << ((198) & 63)) {
            case -2147483648:
                var v100 bool = true
                _ = v100
            case -1:
                v77 -= 255
            case 164:
                fmt.Println(v77, v77, fi99)
            default:
                if (true || (-128 < 100)) {
                    var v101 int8 = 100
                    if false {
                        if len(v4) > 0 {
                            v4[0] = ("hello" + "a")
                        }
                    }
                    _ = v101
                } else {
                    v76 += 32767
                }
            }
            v102 := GenS8{F0: "a", F1: fi99}
            var v103 string = "hello"
            _ = fi99
            _ = v102
            _ = v103
        }
    }
    if (127 != 1) {
        delete(v1, 42)
        var v104 int32 = -684
        _ = v104
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v76
    _ = v77
}
