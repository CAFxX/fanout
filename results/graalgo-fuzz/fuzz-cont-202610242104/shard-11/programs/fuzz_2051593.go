package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64; F1 int64; F2 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool }

func main() {
    var v1 int16 = 1
    var v2 int32 = 639
    var v3 uint8 = 0
    if false {
        if (((false && false) || (false && false)) && false) {
            v4 := []int{920, 127, 0}
            for fi5 := 0; fi5 < 3; fi5++ {
                for fi6 := 0; fi6 < 2; fi6++ {
                    v1 = v1
                    if false {
                        fmt.Println(v3)
                        v1 = -1
                        if true {
                            v4 = append(v4, fi5)
                            if (!true) {
                                v1 = 82
                                v1 += v1
                            } else {
                                var v7 uint64 = 42
                                v2 -= v2
                                _ = v7
                            }
                        } else {
                            var v8 float64 = (3.14159 - 1.5)
                            _ = v8
                        }
                    } else {
                        if (!(1 > 1)) {
                            fmt.Println(v2, fi6)
                        } else {
                            v3 -= v3
                            switch -768 {
                            case -9223372036854775808:
                                var v9 int32 = (^(v2))
                                _ = v9
                            case 0:
                                var v10 uint = 18446744073709551615
                                _ = v10
                            case 32768:
                                v11 := GenS1{F0: "hello", F1: 42, F2: v2}
                                _ = v11
                            default:
                                switch "" {
                                case "hello":
                                    var v12 uint32 = 2
                                    _ = v12
                                case "a":
                                    fmt.Println(fi6, v2)
                                default:
                                    v3 += v3
                                }
                            }
                            if true {
                                fmt.Println(v3, fi6)
                            }
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v2 = (v2 + (v2 ^ v2))
                    }
                    switch "a" {
                    case "\n\t":
                        switch (("foo bar" + "hello") + "\n\t") {
                        case "a":
                            v13 := GenS2{F0: int(v1), F1: 235}
                            _ = v13
                        case "\n\t":
                            switch "hello" {
                            case "a":
                                if ("hello" == "\n\t") {
                                    var v14 int = fi5
                                    v3 = (v3 / ((v3) | 1))
                                    var v15 string = "foo bar"
                                    _ = v14
                                    _ = v15
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    switch fi6 {
                                    case 235:
                                        switch "\n\t" {
                                        case "":
                                            v1 = 970
                                        case "a":
                                            var v16 string = "hello"
                                            _ = v16
                                        case "\n\t":
                                            _ = len(v4)
                                        }
                                    case -1:
                                        if (v1 <= v1) {
                                            var v17 uint32 = 26
                                            v4 = append(v4, fi5)
                                            var v18 float32 = float32(((1.5 + 0.1) * 1.5))
                                            _ = v17
                                            _ = v18
                                        }
                                    default:
                                        v19 := GenS2{F0: fi5, F1: 564}
                                        _ = v19
                                    }
                                }
                            case "hello":
                                var v20 bool = (1 <= -1)
                                _ = v20
                            }
                        case "foo bar":
                            v2 += (v2 + v2)
                        }
                    case "hello":
                        fmt.Println(fi6, v1)
                    default:
                        v21 := GenS3{F0: 567, F1: 0}
                        _ = v21
                    }
                    _ = fi6
                }
                _ = fi5
            }
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
            _ = v4
        }
    }
    var v22 uint16 = 1
    switch 4294967295 {
    case -372:
        if true {
            if true {
                v23 := GenS2{F0: 32768, F1: 2}
                v24 := GenS1{F0: (("\n\t" + "foo bar") + "hello"), F1: 65535, F2: 0}
                _ = v23
                _ = v24
            }
            switch "\n\t" {
            case "foo bar":
                var v25 uint = 2
                _ = v25
            case "":
                fmt.Println(v2, v2, v22)
            default:
                v3 = 42
            }
        }
    case 42:
        var v26 uint8 = (v3 % ((v3) | 1))
        _ = v26
    case 9223372036854775807:
        v27 := GenS4{F0: "hello", F1: 922}
        _ = v27
    }
    var v28 bool = false
    v2 += (^(v2))
    var v29 uint8 = 2
    v30 := GenS1{F0: "hello", F1: 2147483647, F2: v2}
    v31 := GenS5{F0: 2, F1: -9223372036854775808}
    switch 797 {
    case 128:
        if false {
            var v32 int32 = v2
            v33 := []int{32767, -9223372036854775808, 127}
            fmt.Println(v1, v32)
            _ = v32
            _ = v33
        }
    case -572:
        v28 = (620 > 65535)
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        if ((-2.25 >= (0.1 - ((((-2.25 / (3.14159 + (1.5 / (3.14159 - 0.1)))) + (1.5 / -2.25)) / (0.0 + float64(float32((0.1 * 0.1))))) / 3.14159))) && (-23 <= -19)) {
            var v35 uint32 = 42
            _ = v35
        }
        for fi36 := 0; fi36 < 15; fi36++ {
            var v37 uint64 = 1
            if (!(100 >= 42)) {
                if ("" >= "") {
                    if (42 < int64(fi36)) {
                        v38 := []int{(fi36 ^ fi36), -570}
                        fmt.Println(fi36, fi34)
                        _ = v38
                    } else {
                        v28 = v28
                        v39 := GenS5{F0: 841, F1: (fi36 / ((int(v2)) | 1))}
                        _ = v39
                    }
                }
                v40 := []int{(fi34 >> ((189) & 63))}
                v41 := GenS6{F0: v2, F1: (-(v2))}
                _ = v40
                _ = v41
            } else {
                if (uint64(v22) <= (v37 << ((18446744073709551615) & 63))) {
                    if v28 {
                        var v42 bool = (v28 || true)
                        _ = v42
                    } else {
                        switch "" {
                        case "hello":
                            fmt.Println(v37, v2, v22)
                        case "":
                            var v43 uint64 = 1
                            _ = v43
                        default:
                            fmt.Println(v29, fi36)
                        }
                        v44 := GenS7{F0: v28, F1: (v37 - (v37 | (v37 & v37)))}
                        _ = v44
                    }
                    fmt.Println(v1, fi36, v22)
                } else {
                    v22 -= 255
                }
                v29 -= (v3 << ((252) & 63))
                var v45 uint8 = (^(v3))
                _ = v45
            }
            v46 := []int{372, (fi36 * (fi36 & int(v3))), (-(fi34))}
            _ = fi36
            _ = v37
            _ = v46
        }
        if false {
            var v47 string = "\n\t"
            v47 = ("\n\t" + "a")
            if v28 {
                switch v47 {
                case "a":
                    var v48 uint16 = 32767
                    _ = v48
                }
                var v49 uint = 42
                v50 := GenS7{F0: v28, F1: 511}
                _ = v49
                _ = v50
            } else {
                var v51 int = fi34
                v52 := GenS8{F0: (v28 && false)}
                _ = v51
                _ = v52
            }
            _ = v47
        } else {
            for fi53 := 0; fi53 < 12; fi53++ {
                var v54 uint32 = 2147483647
                v55 := []int{}
                switch "a" {
                case "\n\t":
                    if v28 {
                        if len(v55) > 0 {
                            v55[0] = (fi53 - fi53)
                        }
                        var v56 uint16 = (^(v22))
                        switch (-(fi53)) {
                        case -1:
                            v55 = append(v55, -2147483648)
                        case 4294967295:
                            if v28 {
                                v55 = append(v55, (fi53 % (((fi34 >> ((1) & 63))) | 1)))
                                switch 127 {
                                case 42:
                                    var v57 uint = 42
                                    _ = v57
                                }
                                var v58 bool = v28
                                _ = v58
                            }
                        }
                        _ = v56
                    } else {
                        if (uint64(v22) == 4294967295) {
                            v28 = (((v29 ^ (v29 | (v3 >> ((0) & 63)))) >= 9) && (!(v22 > (v22 * v22))))
                            v1 -= v1
                        } else {
                            if (!(fi34 <= 32768)) {
                                var v59 uint32 = v54
                                v60 := GenS9{F0: uint64(v59), F1: 9223372036854775807, F2: -9223372036854775808}
                                _ = v59
                                _ = v60
                            }
                        }
                        if false {
                            v54 = (^(v54))
                            switch ((((("a" + ("\n\t" + "foo bar")) + "\n\t") + "a") + "a") + "") {
                            case "a":
                                var v61 int32 = (v2 / ((v2) | 1))
                                _ = v61
                            case "":
                                v22 = 32767
                            }
                            var v62 uint64 = 355
                            _ = v62
                        } else {
                            var v63 uint8 = (v29 * 82)
                            if false {
                                switch "" {
                                case "\n\t":
                                    switch (fi34 >> ((592) & 63)) {
                                    case 858:
                                        if (0.1 == 0.1) {
                                            v64 := GenS10{F0: (fi53 | 32767)}
                                            _ = v64
                                        }
                                    }
                                case "hello":
                                    fmt.Println(v2, v28, v22)
                                case "":
                                    v65 := GenS9{F0: 2, F1: -9223372036854775808, F2: 1}
                                    _ = v65
                                default:
                                    var v66 float32 = (1.5 + 3.14159)
                                    _ = v66
                                }
                                v67 := GenS8{F0: false}
                                v68 := GenS11{F0: true}
                                _ = v67
                                _ = v68
                            }
                            _ = v63
                        }
                    }
                }
                _ = fi53
                _ = v54
                _ = v55
            }
        }
        _ = fi34
    }
    fmt.Println(v2, v22, v1)
    v3 = 2
    v3 -= (v3 % ((v3) | 1))
    v69 := []int{-503}
    switch -690 {
    case 1:
        switch 42 {
        case 2147483647:
            var v70 uint = 316
            _ = v70
        case 128:
            switch (("hello" + "hello") + "hello") {
            case "":
                v3 -= 99
            case "a":
                v28 = (145 != v3)
            case "foo bar":
                v3 = v29
            default:
                var v71 int16 = (v1 / ((v1) | 1))
                _ = v71
            }
        case 9223372036854775807:
            fmt.Println(v28, v29, v3)
        default:
            for fi72 := 0; fi72 < 9; fi72++ {
                switch "a" {
                case "":
                    var v73 int = (^(fi72))
                    _ = v73
                case "a":
                    fmt.Println(v22, v28, v29)
                case "hello":
                    fmt.Println(v3, v29, fi72)
                }
                for fi74 := 0; fi74 < 15; fi74++ {
                    if (100 <= -116) {
                        v1 += 42
                        var v75 int32 = v2
                        _ = v75
                    }
                    fmt.Println(v28)
                    _ = fi74
                }
                if v28 {
                    var v76 string = "\n\t"
                    v29 -= (v29 / (((v29 >> ((4294967295) & 63))) | 1))
                    var v77 uint32 = uint32(v22)
                    _ = v76
                    _ = v77
                } else {
                    fmt.Println(v29, fi72)
                }
                _ = fi72
            }
        }
    default:
        var v78 int8 = 0
        _ = v78
    }
    _ = len(v69)
    switch "\n\t" {
    case "hello":
        if len(v69) >= 2 {
            v69 = v69[1:len(v69)]
        }
    default:
        var v79 int32 = int32(v3)
        _ = v79
    }
    if v28 {
        fmt.Println(v2, v1, v1)
        v80 := GenS7{F0: true, F1: 328}
        v2 -= v2
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v69:", len(v69), v69)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v69
}
