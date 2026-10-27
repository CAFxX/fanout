package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint; F2 uint16 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string; F1 uint; F2 string }
type GenS9 struct { F0 int64 }

func main() {
    v1 := []uint32{4294967295, 530, 4294967295, 2147483648}
    var v2 int8 = -45
    var v3 int16 = int16(v2)
    v4 := []int{0, 0, 462}
    v5 := make(map[string]uint32)
    var v6 uint32 = 0
    var v7 float64 = 3.14159
    v6 = 2
    for fi8 := 0; fi8 < 2; fi8++ {
        delete(v5, ("\n\t" + ""))
        _ = fi8
    }
    if (488 > 42) {
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                v6 = v6
                v3 = (-(v3))
                v2 += 127
                var v11 string = "hello"
                var v12 uint32 = (v6 - v6)
                v3 = v3
                switch fi9 {
                case 32767:
                    v11 = v11
                case -1:
                    _ = len(v1)
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            if (true && true) {
                var v13 int = 794
                _ = v13
            }
            _ = fi9
        }
    } else {
        if (!((v7 - (v7 * (float64(1.5) + (v7 * -2.25)))) <= -2.25)) {
            var v14 int16 = (v3 ^ (^(v3)))
            _ = v14
        }
        v7 = -2.25
    }
    if false {
        switch "foo bar" {
        case "foo bar":
            v15 := v5["hello"]
            _ = v15
        case "":
            if false {
                for fi16 := 0; fi16 < 8; fi16++ {
                    for fi17 := 0; fi17 < 10; fi17++ {
                        if false {
                            var v18 uint32 = 2147483647
                            switch "" {
                            case "\n\t":
                                fmt.Println(v6)
                            case "hello":
                                if false {
                                    v3 = v3
                                    var v19 float32 = (1e10 * float32(3.14159))
                                    var v20 float64 = (v7 - v7)
                                    _ = v19
                                    _ = v20
                                }
                            case "":
                                if (!(-2.25 <= (0.0 * float32(v7)))) {
                                    v7 -= float64(0.0)
                                    fmt.Println(v18)
                                } else {
                                    switch ("foo bar" + "a") {
                                    case "foo bar":
                                        var v21 uint64 = 2
                                        _ = v21
                                    case "":
                                        var v22 float32 = 1e10
                                        _ = v22
                                    case "\n\t":
                                        v1 = append(v1, 0)
                                    default:
                                        fmt.Println(v3)
                                    }
                                    if true {
                                        var v23 float64 = float64((1.5 / 1.0))
                                        v18 -= (v18 % ((2147483648) | 1))
                                        var v24 uint64 = 1
                                        _ = v23
                                        _ = v24
                                    } else {
                                        var v25 int = (^(fi16))
                                        switch (-(fi17)) {
                                        case 9223372036854775807:
                                            _ = len(v5)
                                        default:
                                            _ = len(v5)
                                        }
                                        v3 = v3
                                        _ = len(v1)
                                        v25 += fi17
                                        fmt.Println(fi16)
                                        if len(v4) > 0 {
                                            v4[0] = int(v3)
                                        }
                                        v6 = (^(v18))
                                        _ = v25
                                    }
                                    v26 := GenS1{F0: false}
                                    _ = v26
                                }
                            default:
                                v6 = (v18 ^ v18)
                            }
                            v7 -= float64(0.1)
                            _ = v18
                        } else {
                            v2 -= (v2 ^ (v2 + 38))
                            var v27 string = ""
                            var v28 float32 = (((((-2.25 * 0.0) * float32(0.0)) + (1e10 - 0.0)) * (3.14159 * (1e10 + -2.25))) + 3.14159)
                            _ = v27
                            _ = v28
                        }
                        fmt.Println(v3, v7)
                        _ = fi17
                    }
                    for fi29 := 0; fi29 < 15; fi29++ {
                        var v30 float64 = v7
                        v3 += v3
                        var v31 uint8 = 10
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    if (((!true) && (!(!true))) || (1 <= 182)) {
                        for fi32 := 0; fi32 < 2; fi32++ {
                            v7 = v7
                            _ = fi32
                        }
                    }
                    _ = fi16
                }
                var v33 int8 = -128
                fmt.Println(v6, v7, v7)
                _ = v33
            }
        case "a":
            v34 := GenS2{F0: v6}
            _ = v34
        }
        if len(v1) > 0 {
            v1[0] = (v6 - (v6 * (v6 % ((v6) | 1))))
        }
        if (2147483647 > 2147483647) {
            var v35 uint16 = 65535
            var v36 uint16 = 2
            v5["hello"] = v6
            _ = v35
            _ = v36
        }
    } else {
        v7 += 1e10
        var v37 string = ""
        for fi38 := 0; fi38 < 19; fi38++ {
            var v39 float64 = ((1e10 * v7) / ((0.0 - v7) - (v7 + -2.25)))
            _ = fi38
            _ = v39
        }
        _ = v37
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v40 int64 = 9223372036854775807
    v3 += v3
    if (true && (((0.1 / float64((-2.25 * 3.14159))) * 1e10) >= ((((0.1 - (-2.25 + 1e10)) / v7) + v7) + v7))) {
        for fi41 := 0; fi41 < 20; fi41++ {
            if (42 >= 2) {
                var v42 float32 = -2.25
                switch fi41 {
                case 2147483647:
                    v43 := GenS3{F0: v40}
                    _ = v43
                }
                _ = v42
            } else {
                v40 -= v40
            }
            _ = fi41
        }
        v5[""] = (v6 % ((2147483648) | 1))
    } else {
        for fi44 := 0; fi44 < 17; fi44++ {
            if false {
                var v45 int64 = -44
                for fi46 := 0; fi46 < 4; fi46++ {
                    v2 += (-(v2))
                    v47 := GenS4{F0: (v45 >> ((255) & 63)), F1: 323}
                    _ = fi46
                    _ = v47
                }
                _ = v45
            } else {
                var v48 string = ""
                v49 := GenS5{F0: v3, F1: 42, F2: 768}
                v1 = append(v1, 294)
                _ = v48
                _ = v49
            }
            _ = len(v4)
            _ = fi44
        }
    }
    if (!true) {
        v50 := GenS3{F0: 9223372036854775807}
        var v51 float32 = 3.14159
        var v52 float64 = v7
        _ = v50
        _ = v51
        _ = v52
    } else {
        var v53 float32 = 3.14159
        _ = v53
    }
    switch 4294967295 {
    case 4294967295:
        if false {
            var v54 uint16 = 42
            _ = v54
        } else {
            switch ("a" + (((("foo bar" + ("" + "hello")) + "hello") + "") + ("a" + "hello"))) {
            case "foo bar":
                var v55 uint64 = 4294967295
                _ = v55
            case "\n\t":
                fmt.Println(v40, v7, v3)
            case "a":
                v56 := GenS6{F0: 340, F1: "foo bar"}
                _ = v56
            }
        }
    default:
        switch -955 {
        case 974:
            var v57 int64 = (v40 >> ((4294967295) & 63))
            _ = v57
        case 32768:
            if false {
                switch 2147483647 {
                case 765:
                    v58 := GenS7{F0: 1000}
                    _ = v58
                case -2147483648:
                    var v59 uint32 = (^(v6))
                    _ = v59
                case 718:
                    v60 := GenS2{F0: 523}
                    _ = v60
                default:
                    v40 += v40
                }
                delete(v5, "\n\t")
                switch 9223372036854775807 {
                case 32768:
                    var v61 float64 = (((3.14159 / ((0.1 / (v7 * v7)) + (v7 / -2.25))) / v7) * v7)
                    _ = v61
                case 2147483647:
                    if true {
                        switch "foo bar" {
                        case "foo bar":
                            v62 := GenS8{F0: "hello", F1: 42, F2: "foo bar"}
                            _ = v62
                        case "hello":
                            for fi63 := 0; fi63 < 16; fi63++ {
                                v7 = -2.25
                                v4 = append(v4, (fi63 & 9223372036854775807))
                                for fi64 := 0; fi64 < 13; fi64++ {
                                    switch "a" {
                                    case "\n\t":
                                        switch fi64 {
                                        case 960:
                                            var v65 float64 = float64(((-2.25 * 0.1) + 3.14159))
                                            _ = v65
                                        case 695:
                                            v66 := GenS6{F0: 651, F1: "hello"}
                                            v7 = v7
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v66
                                        case -1:
                                            v3 += (v3 | -190)
                                        default:
                                            _ = len(v4)
                                        }
                                    }
                                    _ = fi64
                                }
                                _ = fi63
                            }
                        default:
                            v5["foo bar"] = v6
                        }
                    } else {
                        var v67 uint16 = 0
                        if true {
                            fmt.Println(v2)
                            v68 := GenS9{F0: v40}
                            fmt.Println(v2, v3, v7)
                            _ = v68
                        }
                        _ = v67
                    }
                }
            } else {
                v69 := GenS1{F0: false}
                _ = v69
            }
        case 127:
            v40 = (v40 >> ((0) & 63))
        default:
            var v70 uint64 = 2
            _ = v70
        }
    }
    v71 := GenS5{F0: (v3 ^ (v3 ^ (v3 / (((v3 * 32767)) | 1)))), F1: 369, F2: 544}
    v40 -= 0
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v71
}
