package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint8; F2 uint }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int8; F1 int32; F2 int }
type GenS11 struct { F0 uint32; F1 uint8; F2 int }
type GenS12 struct { F0 uint64 }
func genFn0(p0 uint8, p1 uint, p2 uint32) int {
    for fi5 := 0; fi5 < 2; fi5++ {
        switch "foo bar" {
        case "":
            v6 := GenS1{F0: (p0 >> (((p1 + 274)) & 63)), F1: -128}
            _ = v6
        case "a":
            switch fi5 {
            case -755:
                var v7 uint64 = 588
                _ = v7
            case -2147483648:
                var v8 int16 = -889
                _ = v8
            case 2147483647:
                var v9 uint = (p1 ^ (p1 - 2))
                _ = v9
            }
        case "hello":
            p2 += p2
        default:
            switch "a" {
            case "\n\t":
                v10 := GenS2{F0: (p0 + p0)}
                _ = v10
            case "foo bar":
                var v11 int = -9223372036854775808
                _ = v11
            }
        }
        _ = fi5
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(p1, p2, p0)
    case "hello":
        v12 := []int{-244, 9223372036854775807, 9223372036854775807}
        _ = v12
    case "a":
        v13 := GenS3{F0: 110, F1: -2147483648, F2: 64}
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS4{F0: 2}
        _ = fi14
        _ = v15
    }
    switch int(p2) {
    case 127:
        if (((float64((-2.25 + 1e10)) + -2.25) + -2.25) == 3.14159) {
            v16 := GenS5{F0: 497, F1: 181, F2: 0}
            p0 = 2
            switch ("\n\t" + "a") {
            case "\n\t":
                v17 := []int{-9223372036854775808, 127}
                _ = v17
            case "foo bar":
                p0 = p0
            default:
                fmt.Println(p0)
            }
            _ = v16
        }
    default:
        v18 := GenS3{F0: -10, F1: -820, F2: 0}
        _ = v18
    }
    var v19 float32 = (3.14159 + float32(1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return -92
}
func genFn1(p0 uint8, p1 bool, p2 int8) int {
    var v20 float32 = ((0.1 / 1.0) + ((float32(float64(3.14159)) / 1.0) / 1e10))
    if p1 {
        if (!p1) {
            fmt.Println(p0, p0, v20)
            var v21 float32 = v20
            _ = v21
        } else {
            for fi22 := 0; fi22 < 9; fi22++ {
                fmt.Println(p0)
                _ = fi22
            }
            var v23 int16 = -1
            _ = v23
        }
    }
    v24 := GenS6{F0: 2, F1: (-(p2))}
    switch ("foo bar" + "foo bar") {
    case "hello":
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v24
    return -308
}

func main() {
    var v1 int16 = 1
    var v2 bool = true
    v3 := make(map[string]bool)
    var v4 float64 = float64(((1.5 - 1.5) / 1.5))
    if (826 >= 0) {
        v25 := []int{32768, genFn1(255, (v2 || (183 > 699)), 127)}
        fmt.Println(v2, v4, v2)
        v2 = false
        _ = v25
    }
    switch "foo bar" {
    case "a":
        switch 276 {
        case 95:
            if v2 {
                for fi26 := 0; fi26 < 19; fi26++ {
                    v27 := GenS3{F0: 42, F1: -613, F2: 153}
                    v28 := v3["a"]
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
        case -844:
            fmt.Println(v4, v2, v1)
        }
    }
    v1 += v1
    for fi29 := 0; fi29 < 20; fi29++ {
        for fi30 := 0; fi30 < 6; fi30++ {
            delete(v3, (("foo bar" + "hello") + "foo bar"))
            _ = fi30
        }
        _ = fi29
    }
    if v2 {
        for fi31 := 0; fi31 < 11; fi31++ {
            if (0.0 <= ((3.14159 - float32((float64((0.0 * 0.1)) + v4))) - -2.25)) {
                var v32 uint32 = 4294967295
                for fi33 := 0; fi33 < 13; fi33++ {
                    if (!v2) {
                        v34 := GenS6{F0: 4294967295, F1: 0}
                        _ = v34
                    }
                    switch "hello" {
                    case "hello":
                        v35 := GenS7{F0: 100, F1: 31}
                        _ = v35
                    case "a":
                        v36 := GenS4{F0: 45}
                        _ = v36
                    }
                    _ = fi33
                }
                v1 += v1
                _ = v32
            } else {
                var v37 float32 = 0.0
                v37 = v37
                _ = v37
            }
            var v38 string = "foo bar"
            _ = fi31
            _ = v38
        }
    } else {
        var v39 uint8 = 141
        switch 32767 {
        case -9223372036854775808:
            v40 := []int{9223372036854775807, 1, 1, 2147483647}
            _ = v40
        case 42:
            var v41 int = -304
            _ = v41
        }
        _ = v39
    }
    if ((3.14159 + ((3.14159 - 0.0) - ((3.14159 / float32((float64(0.1) - v4))) / 3.14159))) != float32(v4)) {
        var v42 int8 = 42
        _ = v42
    }
    switch 32768 {
    case 0:
        v2 = v2
    case -793:
        var v43 float32 = ((float32(3.14159) - 1.5) / ((-2.25 * -2.25) / 3.14159))
        _ = v43
    }
    v44 := GenS8{F0: 42}
    var v45 uint = 174
    if v2 {
        var v46 int64 = 9223372036854775807
        for fi47 := 0; fi47 < 13; fi47++ {
            if ((v46 >> (((v45 / ((42) | 1))) & 63)) != 9223372036854775807) {
                v48 := []int{(fi47 << ((v45) & 63)), fi47, (fi47 | int(v46))}
                _ = len(v3)
                _ = v48
            } else {
                if (true && ((!(593 >= fi47)) && true)) {
                    v49 := []int{(fi47 ^ -1), (fi47 >> ((v45) & 63))}
                    _ = v49
                }
            }
            _ = fi47
        }
        switch ("foo bar" + "hello") {
        case "\n\t":
            if (v2 || v2) {
                var v50 uint8 = 39
                v51 := GenS1{F0: (v50 << ((v45) & 63)), F1: 1}
                _ = v50
                _ = v51
            } else {
                v52 := GenS7{F0: 0, F1: 42}
                _ = v52
            }
        default:
            v4 = v4
        }
        _ = v46
    }
    var v53 uint64 = 623
    for fi54 := 0; fi54 < 4; fi54++ {
        v55 := []int{(fi54 * fi54)}
        for fi56 := 0; fi56 < 12; fi56++ {
            v1 = v1
            _ = fi56
        }
        if ((-2147483648 >= 0) && true) {
            if ("a" == "foo bar") {
                v57 := GenS9{F0: v2}
                var v58 int16 = (v1 / (((v1 / ((v1) | 1))) | 1))
                for fi59 := 0; fi59 < 12; fi59++ {
                    var v60 int64 = 0
                    _ = fi59
                    _ = v60
                }
                _ = v57
                _ = v58
            }
        }
        _ = fi54
        _ = v55
    }
    v61 := GenS10{F0: 42, F1: -627, F2: -1}
    for fi62 := 0; fi62 < 8; fi62++ {
        _ = len(v3)
        if (0.0 >= 0.0) {
            v63 := GenS11{F0: 2, F1: 2, F2: fi62}
            var v64 int64 = -9223372036854775808
            v3[(("a" + "\n\t") + "foo bar")] = false
            _ = v63
            _ = v64
        }
        switch "" {
        case "a":
            if (!false) {
                v65 := v3[("hello" + "hello")]
                var v66 float32 = ((1.5 * 1.5) + 0.1)
                _ = v65
                _ = v66
            } else {
                if false {
                    v67 := v3[""]
                    v68 := GenS6{F0: v53, F1: -68}
                    for fi69 := 0; fi69 < 12; fi69++ {
                        v45 = v45
                        if (uint8(v45) > 42) {
                            switch genFn1(1, (genFn1(128, v67, -1) == (^(fi69))), -1) {
                            case 32768:
                                v70 := GenS12{F0: 2}
                                _ = v70
                            }
                            v1 += (^(v1))
                        }
                        fmt.Println(v4)
                        _ = len(v3)
                        _ = fi69
                    }
                    _ = v67
                    _ = v68
                }
                if v2 {
                    if (!v2) {
                        switch (fi62 - (fi62 >> ((18446744073709551615) & 63))) {
                        case -732:
                            v45 -= 42
                        }
                        if true {
                            v1 += v1
                            var v71 uint = (v45 ^ uint(v53))
                            _ = v71
                        }
                    }
                }
                var v72 int8 = -1
                _ = v72
            }
        case "foo bar":
            if v2 {
                var v73 int = fi62
                _ = v73
            } else {
                v74 := GenS10{F0: -128, F1: -2147483648, F2: (fi62 ^ (fi62 + (fi62 % ((fi62) | 1))))}
                _ = v74
            }
        default:
            for fi75 := 0; fi75 < 18; fi75++ {
                fmt.Println(fi62)
                _ = fi75
            }
        }
        _ = fi62
    }
    var v76 int64 = -1
    for fi77 := 0; fi77 < 1; fi77++ {
        var v78 uint32 = 241
        _ = fi77
        _ = v78
    }
    v79 := GenS6{F0: (v53 << ((42) & 63)), F1: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v53:", v53)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v53
    _ = v61
    _ = v76
    _ = v79
}
