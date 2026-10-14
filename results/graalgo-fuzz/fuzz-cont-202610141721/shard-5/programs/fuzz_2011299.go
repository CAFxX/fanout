package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 uint64; F1 int16; F2 uint16 }

func main() {
    var v1 uint = 255
    var v2 uint16 = 32767
    var v3 float64 = ((((1.5 / 1.5) - float64(1e10)) * 1.5) / 1.0)
    var v4 uint16 = (^(v2))
    v5 := make(map[uint32]int)
    v4 = 32767
    var v6 uint16 = (v4 | v2)
    switch 32767 {
    case -1:
        var v7 uint64 = 18446744073709551615
        _ = v7
    case 128:
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: ("a" + ("hello" + "\n\t")), F1: v6}
            _ = fi8
            _ = v9
        }
    }
    v5[865] = 693
    if false {
        v4 = (v6 - 2)
        var v10 uint32 = 4294967295
        v11 := []int{}
        _ = v10
        _ = v11
    }
    var v12 uint16 = (v6 * (v4 << ((v1) & 63)))
    switch "\n\t" {
    case "foo bar":
        switch 0 {
        case -2147483648:
            if (true || (-2147483648 != 2147483647)) {
                v1 = v1
                v13 := GenS2{F0: 4294967295}
                if (!(true || true)) {
                    v14 := []int{1}
                    v6 += (v4 | (v12 & 2))
                    _ = v14
                }
                _ = v13
            } else {
                var v15 int64 = -9223372036854775808
                _ = v15
            }
        case 1:
            v16 := GenS3{F0: 100, F1: 4294967295}
            _ = v16
        default:
            for fi17 := 0; fi17 < 15; fi17++ {
                v6 = 2
                v18 := v5[0]
                v6 = v2
                _ = fi17
                _ = v18
            }
        }
    case "hello":
        fmt.Println(v12)
    case "":
        for fi19 := 0; fi19 < 8; fi19++ {
            v4 = v4
            switch fi19 {
            case 4294967295:
                fmt.Println(v1, v1)
            }
            _ = fi19
        }
    default:
        var v20 uint64 = 1
        _ = v20
    }
    v6 += v12
    if true {
        for fi21 := 0; fi21 < 20; fi21++ {
            if ((false || ((!false) && ("\n\t" < ("" + (("hello" + (("a" + "") + (((("hello" + "hello") + ((("hello" + ("a" + ("a" + "a"))) + ("foo bar" + "\n\t")) + "a")) + (("\n\t" + ("" + "hello")) + "\n\t")) + "hello"))) + "foo bar"))))) || (v3 < 3.14159)) {
                switch fi21 {
                case -409:
                    v22 := []int{}
                    _ = v22
                }
            }
            if (!true) {
                var v23 int = (fi21 * (-(fi21)))
                _ = v23
            } else {
                var v24 int8 = -128
                for fi25 := 0; fi25 < 18; fi25++ {
                    if (v24 == (^(v24))) {
                        switch "" {
                        case "a":
                            var v26 string = ""
                            _ = v26
                        case "hello":
                            v27 := GenS4{F0: -993}
                            _ = v27
                        case "":
                            v2 = (v2 + 236)
                        }
                        v28 := GenS3{F0: v24, F1: 468}
                        _ = v28
                    }
                    switch (fi25 + (fi21 + fi25)) {
                    case 888:
                        v2 = 1
                    case 127:
                        if (false || false) {
                            var v29 int16 = 32767
                            v24 = v24
                            _ = v29
                        } else {
                            v30 := GenS2{F0: (v1 << (((v1 << ((930) & 63))) & 63))}
                            var v31 uint8 = 68
                            v12 -= v4
                            _ = v30
                            _ = v31
                        }
                        v32 := []int{2147483647}
                        var v33 uint16 = 2
                        _ = v32
                        _ = v33
                    case -9223372036854775808:
                        if (9223372036854775807 >= 42) {
                            v34 := GenS5{F0: true}
                            v6 = v6
                            v24 -= (v24 % ((v24) | 1))
                            var v35 int16 = -32768
                            _ = v34
                            _ = v35
                        }
                    }
                    v36 := GenS6{F0: 1, F1: 904}
                    _ = fi25
                    _ = v36
                }
                if false {
                    switch (^(fi21)) {
                    case -1:
                        switch "" {
                        case "foo bar":
                            var v37 string = ((("" + "foo bar") + ("\n\t" + "\n\t")) + "\n\t")
                            _ = v37
                        case "a":
                            var v38 uint8 = 11
                            _ = v38
                        case "":
                            if (32767 >= 42) {
                                v39 := []int{(-(fi21)), (fi21 / (((fi21 * int(v1))) | 1))}
                                _ = v39
                            } else {
                                v1 = (v1 % (((^(v1))) | 1))
                                fmt.Println(v3, v3, fi21)
                            }
                        default:
                            v40 := []int{(fi21 + fi21), (fi21 % (((fi21 ^ (fi21 * 0))) | 1)), fi21, fi21}
                            _ = v40
                        }
                    default:
                        v1 = v1
                    }
                    for fi41 := 0; fi41 < 9; fi41++ {
                        v12 -= (^(v6))
                        if true {
                            var v42 int8 = v24
                            var v43 int8 = (-(v42))
                            _ = v42
                            _ = v43
                        }
                        _ = fi41
                    }
                    var v44 bool = false
                    _ = v44
                } else {
                    v45 := GenS7{F0: uint8(v1), F1: (v4 % (((v12 % (((v12 - 255)) | 1))) | 1))}
                    v6 = v12
                    v6 = v4
                    _ = v45
                }
                _ = v24
            }
            if (float32(v3) >= (1e10 + 3.14159)) {
                for fi46 := 0; fi46 < 13; fi46++ {
                    var v47 int32 = 986
                    if ("" < "") {
                        fmt.Println(v3, v12)
                        fmt.Println(fi21, v3)
                        if (2147483648 != 0) {
                            v3 += float64(3.14159)
                            _ = len(v5)
                            v48 := []int{fi46}
                            _ = v48
                        } else {
                            var v49 bool = (true || (-1 > int64(v1)))
                            fmt.Println(fi46, v49, v2)
                            v2 = 0
                            _ = v49
                        }
                    } else {
                        fmt.Println(v1, v4, v47)
                        v50 := GenS2{F0: v1}
                        _ = v50
                    }
                    _ = fi46
                    _ = v47
                }
            }
            _ = fi21
        }
        switch -1 {
        case 42:
            fmt.Println(v6, v1)
        case -9223372036854775808:
            var v51 float64 = 3.14159
            _ = v51
        case 4294967295:
            var v52 uint = (v1 % ((v1) | 1))
            _ = v52
        default:
            v53 := GenS6{F0: (v1 << (((v1 ^ (v1 % ((v1) | 1)))) & 63)), F1: 723}
            _ = v53
        }
        v12 += 803
    }
    for fi54 := 0; fi54 < 20; fi54++ {
        if false {
            for fi55 := 0; fi55 < 18; fi55++ {
                v2 += 0
                _ = fi55
            }
        }
        v2 -= v6
        _ = fi54
    }
    for fi56 := 0; fi56 < 6; fi56++ {
        var v57 uint32 = 842
        _ = fi56
        _ = v57
    }
    fmt.Println(v3, v4, v1)
    if true {
        var v58 uint16 = v2
        _ = v58
    }
    var v59 uint8 = 190
    v4 += (v12 + v6)
    switch 4294967295 {
    case 128:
        for fi60 := 0; fi60 < 7; fi60++ {
            v12 = (v6 & 255)
            v6 -= v4
            var v61 bool = (fi60 >= fi60)
            _ = fi60
            _ = v61
        }
    default:
        switch (((("foo bar" + "a") + "a") + ("\n\t" + ((((("hello" + ("hello" + "hello")) + (("foo bar" + "foo bar") + ("\n\t" + "\n\t"))) + (("hello" + ("foo bar" + (("\n\t" + ("hello" + "a")) + "a"))) + "foo bar")) + "") + (("\n\t" + "hello") + "")))) + ((("" + "hello") + "a") + "")) {
        case "":
            switch 2147483647 {
            case 839:
                switch "foo bar" {
                case "hello":
                    v4 += v6
                case "foo bar":
                    var v62 int16 = 0
                    _ = v62
                case "a":
                    if ((3.14159 / 1.0) < 0.1) {
                        for fi63 := 0; fi63 < 7; fi63++ {
                            v1 -= (^(v1))
                            _ = fi63
                        }
                        for fi64 := 0; fi64 < 5; fi64++ {
                            v12 += (v6 << (((v1 - (v1 * (v1 % (((v1 ^ (v1 + (v1 % ((42) | 1))))) | 1))))) & 63))
                            v65 := []int{}
                            _ = fi64
                            _ = v65
                        }
                    } else {
                        fmt.Println(v4)
                        v12 -= (v2 - v12)
                        switch 438 {
                        case 42:
                            v2 += v6
                        case 503:
                            for fi66 := 0; fi66 < 7; fi66++ {
                                var v67 float32 = (float32(1.5) - 1e10)
                                _ = fi66
                                _ = v67
                            }
                        case 32767:
                            for fi68 := 0; fi68 < 13; fi68++ {
                                fmt.Println(v3)
                                _ = fi68
                            }
                        }
                    }
                default:
                    var v69 uint64 = 18446744073709551615
                    _ = v69
                }
            case 32768:
                var v70 int16 = 0
                _ = v70
            case -9223372036854775808:
                v12 = (v2 + 2)
            }
        case "\n\t":
            if ("" >= ("\n\t" + "a")) {
                var v71 uint32 = 2147483648
                var v72 uint64 = 828
                v3 = ((float64((0.0 * 0.0)) - ((((0.0 + v3) + ((((1e10 * v3) - ((float64(float32((float64(1e10) * (v3 + (-2.25 * -2.25))))) + (v3 * ((((v3 / (3.14159 + 0.1)) - v3) * v3) - (v3 / v3)))) + v3)) - v3) / -2.25)) / v3) * 1.5)) + float64(3.14159))
                _ = v71
                _ = v72
            } else {
                var v73 uint8 = v59
                if (!true) {
                    v74 := []int{0, -801, 0}
                    switch -1 {
                    case 127:
                        var v75 int16 = 1000
                        _ = v75
                    case -1:
                        v12 += (v4 & 32767)
                    default:
                        v2 = v12
                    }
                    for fi76 := 0; fi76 < 14; fi76++ {
                        var v77 int32 = 655
                        v78 := GenS8{F0: 42, F1: 42, F2: (v4 + v6)}
                        _ = fi76
                        _ = v77
                        _ = v78
                    }
                    _ = v74
                } else {
                    var v79 uint8 = uint8(v1)
                    fmt.Println(v2)
                    _ = v79
                }
                _ = v73
            }
        default:
            var v80 int16 = -83
            _ = v80
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
