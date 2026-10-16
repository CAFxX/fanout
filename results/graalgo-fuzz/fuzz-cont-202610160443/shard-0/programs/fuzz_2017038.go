package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32 }
func genFn0(p0 int32, p1 int64, p2 string) uint {
    if false {
        v7 := GenS1{F0: 684, F1: 599}
        fmt.Println(p0, p2)
        _ = v7
    }
    p1 = -9223372036854775808
    if (!true) {
        var v8 uint32 = 2147483648
        for fi9 := 0; fi9 < 11; fi9++ {
            p1 = p1
            fmt.Println(v8, p1, p0)
            switch (fi9 | 9223372036854775807) {
            case 532:
                fmt.Println(p1)
            case -1:
                fmt.Println(p0, v8)
            case 32767:
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 bool = ((0.1 + (((3.14159 * ((float32(float64(0.0)) * 1e10) / 1e10)) - 1e10) / float32(3.14159))) != 3.14159)
                    _ = fi10
                    _ = v11
                }
            default:
                fmt.Println(fi9, p2)
            }
            _ = fi9
        }
        _ = v8
    } else {
        v12 := []int{32767, 1, 32767}
        v13 := GenS2{F0: 42}
        v14 := GenS3{F0: -1, F1: 0}
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return 914
}
func genFn1(p0 int, p1 string) uint8 {
    if true {
        var v15 int = (p0 & (p0 >> ((genFn0(-2147483648, -33, "a")) & 63)))
        switch p0 {
        case 126:
            v16 := GenS4{F0: 2, F1: 1}
            _ = v16
        case 0:
            p1 = "hello"
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            p1 = p1
            for fi18 := 0; fi18 < 15; fi18++ {
                v15 = (fi18 + v15)
                switch p1 {
                case "foo bar":
                    v19 := GenS4{F0: 212, F1: -491}
                    _ = v19
                case "\n\t":
                    p1 = ((p1 + (("hello" + p1) + p1)) + p1)
                default:
                    var v20 bool = (-128 >= 1)
                    switch p1 {
                    case "a":
                        if v20 {
                            switch p0 {
                            case 435:
                                var v21 int32 = 167
                                _ = v21
                            case 2147483647:
                                var v22 int16 = -94
                                _ = v22
                            }
                            var v23 uint = 799
                            _ = v23
                        }
                    default:
                        switch p1 {
                        case "":
                            switch "foo bar" {
                            case "a":
                                v15 -= fi18
                            case "foo bar":
                                var v24 float32 = 3.14159
                                v25 := GenS1{F0: 2, F1: 32}
                                _ = v24
                                _ = v25
                            default:
                                v15 += (-(v15))
                            }
                        case "a":
                            if (!true) {
                                v26 := []int{}
                                var v27 int32 = 42
                                var v28 uint32 = 2147483648
                                v26 = append(v26, (fi17 << ((853) & 63)))
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                        default:
                            p1 = p1
                        }
                    }
                    v29 := []int{(-(fi18)), (v15 & fi17)}
                    _ = v20
                    _ = v29
                }
                var v30 int8 = -1
                _ = fi18
                _ = v30
            }
            _ = fi17
        }
        _ = v15
    }
    if false {
        var v31 uint64 = 18446744073709551615
        switch p1 {
        case "hello":
            if (42 < 732) {
                fmt.Println(p1, p0)
                var v32 float32 = 3.14159
                switch (p0 & p0) {
                case 9223372036854775807:
                    fmt.Println(v32, v32)
                case -394:
                    v31 = (v31 ^ (v31 - (v31 >> ((419) & 63))))
                case 42:
                    for fi33 := 0; fi33 < 1; fi33++ {
                        switch 356 {
                        case 128:
                            for fi34 := 0; fi34 < 3; fi34++ {
                                v35 := []int{4294967295, fi34}
                                fmt.Println(fi33, v32, v32)
                                p0 = (fi33 % (((fi34 * fi33)) | 1))
                                _ = fi34
                                _ = v35
                            }
                        }
                        if (!((-738 < 42) || true)) {
                            v36 := []int{128}
                            v37 := GenS4{F0: 65535, F1: 0}
                            _ = v36
                            _ = v37
                        }
                        _ = fi33
                    }
                default:
                    fmt.Println(p0, p1, p0)
                }
                _ = v32
            }
        }
        _ = v31
    } else {
        for fi38 := 0; fi38 < 7; fi38++ {
            fmt.Println(fi38, p1, p0)
            _ = fi38
        }
        v39 := GenS3{F0: -229, F1: 459}
        v40 := GenS2{F0: 0}
        _ = v39
        _ = v40
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := []string{("foo bar" + "foo bar"), "a", "foo bar", "a"}
    var v2 uint64 = 0
    v3 := []uint32{633, 4294967295, 245}
    v4 := make(map[int]int)
    var v5 int64 = 42
    v6 := []string{"\n\t"}
    switch "\n\t" {
    case "\n\t":
        for fi41 := 0; fi41 < 8; fi41++ {
            if false {
                v2 -= 0
            } else {
                fmt.Println(v2, fi41)
                if true {
                    switch "a" {
                    case "hello":
                        v3 = append(v3, 4294967295)
                    case "":
                        switch (fi41 - 2147483647) {
                        case 42:
                            if (false || (1 < 2)) {
                                switch (("\n\t" + "foo bar") + ("\n\t" + "hello")) {
                                case "hello":
                                    fmt.Println(v2)
                                }
                                var v42 float64 = (1.5 + 3.14159)
                                for fi43 := 0; fi43 < 11; fi43++ {
                                    fmt.Println(fi43, v2)
                                    v44 := GenS5{F0: genFn1(fi41, "hello")}
                                    if false {
                                        v5 = v5
                                    }
                                    _ = fi43
                                    _ = v44
                                }
                                _ = v42
                            }
                        case 127:
                            _ = len(v3)
                        case -2147483648:
                            if false {
                                if true {
                                    switch "" {
                                    case "hello":
                                        v5 = 775
                                    case "\n\t":
                                        v2 += v2
                                    default:
                                        if (v5 < (v5 << ((255) & 63))) {
                                            v5 -= -887
                                            var v45 uint32 = 2147483648
                                            _ = v45
                                        }
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            }
                        }
                    default:
                        v5 += (v5 ^ (v5 / (((v5 | (v5 - v5))) | 1)))
                    }
                    for fi46 := 0; fi46 < 18; fi46++ {
                        switch ("" + (("\n\t" + ("foo bar" + "a")) + "\n\t")) {
                        case "\n\t":
                            var v47 float64 = 1e10
                            _ = v47
                        case "hello":
                            v48 := GenS5{F0: 42}
                            _ = v48
                        case "":
                            v3 = append(v3, 964)
                        default:
                            v2 -= (v2 / ((v2) | 1))
                        }
                        _ = len(v4)
                        _ = fi46
                    }
                } else {
                    v49 := GenS6{F0: "a", F1: 638, F2: 2}
                    _ = v49
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(fi41, v2)
            _ = fi41
        }
    case "hello":
        if (((128 != 32767) || (!false)) && (32767 >= 42)) {
            fmt.Println(v2, v2)
            _ = len(v3)
        } else {
            for fi50 := 0; fi50 < 6; fi50++ {
                v2 += v2
                v2 += 151
                if false {
                    fmt.Println(fi50, fi50)
                }
                _ = fi50
            }
            if (65535 <= 2) {
                var v51 float64 = (1.5 / 1.0)
                v52 := GenS6{F0: (((("" + "foo bar") + "foo bar") + ("hello" + "\n\t")) + ""), F1: (v2 % (((v2 >> ((genFn0(42, v5, ("foo bar" + "hello"))) & 63))) | 1)), F2: (v2 | 2)}
                fmt.Println(v51, v5, v2)
                _ = v51
                _ = v52
            } else {
                if true {
                    v2 += v2
                } else {
                    v2 += (v2 << ((65535) & 63))
                }
                var v53 float32 = 1.5
                _ = v53
            }
            switch 297 {
            case 128:
                if false {
                    v5 -= v5
                    for fi54 := 0; fi54 < 19; fi54++ {
                        var v55 uint = 42
                        v55 -= v55
                        _ = fi54
                        _ = v55
                    }
                }
            case 747:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case -9223372036854775808:
                switch 9223372036854775807 {
                case 127:
                    if false {
                        v56 := GenS7{F0: 0, F1: 748}
                        _ = v56
                    }
                case -181:
                    fmt.Println(v5, v5)
                case 508:
                    fmt.Println(v2)
                }
            }
        }
    case "foo bar":
        if true {
            if true {
                for fi57 := 0; fi57 < 14; fi57++ {
                    for fi58 := 0; fi58 < 16; fi58++ {
                        var v59 uint32 = 2147483648
                        _ = fi58
                        _ = v59
                    }
                    v2 = (^(v2))
                    v6 = append(v6, ("a" + "foo bar"))
                    _ = fi57
                }
                _ = len(v4)
                v60 := GenS8{F0: 32767}
                _ = v60
            }
            for fi61 := 0; fi61 < 8; fi61++ {
                for fi62 := 0; fi62 < 6; fi62++ {
                    v6 = append(v6, "")
                    if ((("hello" + "a") + ("" + ("hello" + (("\n\t" + "") + "foo bar")))) != ("a" + "hello")) {
                        v2 += v2
                        v4[(-(fi61))] = int(v5)
                        v2 -= v2
                    }
                    _ = fi62
                }
                v2 += (v2 - v2)
                if false {
                    v63 := GenS5{F0: 42}
                    delete(v4, 32768)
                    _ = v63
                } else {
                    if false {
                        for fi64 := 0; fi64 < 6; fi64++ {
                            v65 := GenS4{F0: 255, F1: -2147483648}
                            _ = fi64
                            _ = v65
                        }
                    }
                }
                _ = fi61
            }
            var v66 uint64 = v2
            _ = v66
        } else {
            delete(v4, 128)
        }
    }
    var v67 float64 = 3.14159
    v67 = v67
    if (1.5 > float32(v67)) {
        v5 = v5
        for fi68 := 0; fi68 < 3; fi68++ {
            var v69 uint64 = v2
            fmt.Println(v67, v2)
            var v70 int8 = int8(fi68)
            _ = fi68
            _ = v69
            _ = v70
        }
        v71 := GenS9{F0: 621}
        _ = v71
    }
    switch "\n\t" {
    case "foo bar":
        var v72 int8 = -128
        _ = v72
    default:
        v73 := GenS2{F0: 42}
        _ = v73
    }
    v67 = (1e10 + 1.5)
    v67 = v67
    var v74 int64 = v5
    v5 = (v5 ^ v74)
    var v75 uint64 = (v2 * (v2 + v2))
    switch (("a" + "") + "a") {
    case "\n\t":
        fmt.Println(v67, v74)
    default:
        fmt.Println(v74, v5)
    }
    var v76 uint8 = genFn1(128, "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v67:", v67)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v74
    _ = v75
    _ = v76
}
