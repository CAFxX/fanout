package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint; F2 bool }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"a", "hello", "hello"}
    v3 := []string{"foo bar"}
    v4 := make(map[int]int)
    switch "\n\t" {
    case "\n\t":
        var v5 int64 = -114
        _ = v5
    case "":
        switch ("" + "") {
        case "foo bar":
            switch -2147483648 {
            case 32768:
                var v6 int32 = 128
                _ = v6
            case 128:
                var v7 int = -881
                _ = v7
            case 4294967295:
                if false {
                    var v8 int = -9223372036854775808
                    var v9 int8 = -128
                    v8 += v8
                    _ = v8
                    _ = v9
                }
            }
        case "a":
            var v10 uint64 = 0
            _ = v10
        case "":
            var v11 int8 = -128
            _ = v11
        }
    }
    if (127 != 2147483647) {
        v12 := GenS1{F0: 0, F1: 0, F2: 42}
        _ = v12
    } else {
        var v13 uint8 = 0
        var v14 int8 = 127
        _ = v13
        _ = v14
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    var v15 int16 = -32768
    _ = len(v2)
    switch "hello" {
    case "a":
        var v16 float64 = float64(0.1)
        _ = v16
    case "":
        if true {
            switch int(v15) {
            case 909:
                if false {
                    if ((!(true && false)) && false) {
                        v15 = (v15 << ((65535) & 63))
                        for fi17 := 0; fi17 < 11; fi17++ {
                            var v18 int8 = -1
                            _ = fi17
                            _ = v18
                        }
                    } else {
                        if true {
                            for fi19 := 0; fi19 < 2; fi19++ {
                                v20 := GenS1{F0: 42, F1: 18446744073709551615, F2: 255}
                                if (!false) {
                                    for fi21 := 0; fi21 < 14; fi21++ {
                                        var v22 uint8 = 1
                                        if false {
                                            v22 = v22
                                            fmt.Println(fi21, v15)
                                            fmt.Println(fi19, v22)
                                            fmt.Println(fi19, v15, fi21)
                                            fmt.Println(v22)
                                        } else {
                                            var v23 int = fi19
                                            var v24 int32 = int32(v22)
                                            _ = v23
                                            _ = v24
                                        }
                                        var v25 uint8 = 11
                                        _ = fi21
                                        _ = v22
                                        _ = v25
                                    }
                                }
                                _ = len(v3)
                                _ = fi19
                                _ = v20
                            }
                            v26 := GenS1{F0: 468, F1: 849, F2: 1}
                            if true {
                                for fi27 := 0; fi27 < 7; fi27++ {
                                    for fi28 := 0; fi28 < 13; fi28++ {
                                        switch (fi27 * (fi27 ^ -9223372036854775808)) {
                                        case -9223372036854775808:
                                            switch "" {
                                            case "":
                                                var v29 uint = 673
                                                _ = v29
                                            case "a":
                                                if false {
                                                    var v30 uint16 = 542
                                                    v31 := GenS1{F0: 9223372036854775807, F1: 73, F2: (^(v30))}
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    var v32 int = (fi27 >> ((110) & 63))
                                                    _ = v32
                                                }
                                            default:
                                                var v33 float64 = 1.5
                                                _ = v33
                                            }
                                        }
                                        switch (^(fi28)) {
                                        case 2147483647:
                                            v15 = v15
                                        case -860:
                                            fmt.Println(v15, fi28, fi27)
                                        case 32768:
                                            var v34 int = fi28
                                            _ = v34
                                        }
                                        _ = fi28
                                    }
                                    _ = fi27
                                }
                                delete(v1, 2)
                            }
                            _ = v26
                        }
                        if true {
                            v15 = (v15 << ((576) & 63))
                            v15 = 42
                        }
                        v15 = -32768
                    }
                } else {
                    switch ("a" + ("hello" + ("foo bar" + (("foo bar" + ((("hello" + ((("hello" + (((((("\n\t" + ("foo bar" + ("" + ""))) + "a") + "\n\t") + "foo bar") + "") + ("a" + "\n\t"))) + "hello") + ("a" + "a"))) + "foo bar") + ("hello" + ((("foo bar" + "foo bar") + "hello") + "\n\t")))) + (("" + "foo bar") + "a"))))) {
                    case "hello":
                        for fi35 := 0; fi35 < 5; fi35++ {
                            for fi36 := 0; fi36 < 3; fi36++ {
                                _ = len(v3)
                                v15 += v15
                                if (true || true) {
                                    var v37 float64 = 1.5
                                    v15 = v15
                                    _ = v37
                                }
                                _ = fi36
                            }
                            if (1 == 42) {
                                if true {
                                    for fi38 := 0; fi38 < 3; fi38++ {
                                        var v39 int = (fi35 % ((fi38) | 1))
                                        v39 += v39
                                        _ = fi38
                                        _ = v39
                                    }
                                } else {
                                    for fi40 := 0; fi40 < 1; fi40++ {
                                        var v41 uint64 = 18446744073709551615
                                        v41 = (v41 ^ 711)
                                        _ = fi40
                                        _ = v41
                                    }
                                    var v42 uint16 = 42
                                    _ = v42
                                }
                                var v43 int8 = -128
                                v15 = (v15 - -1)
                                _ = v43
                            }
                            fmt.Println(v15, fi35)
                            _ = fi35
                        }
                    default:
                        var v44 int32 = 2147483647
                        _ = v44
                    }
                }
            case -477:
                switch "a" {
                case "":
                    v45 := GenS1{F0: -471, F1: 42, F2: 77}
                    _ = v45
                case "\n\t":
                    if true {
                        fmt.Println(v15)
                    }
                case "a":
                    var v46 int32 = 2147483647
                    _ = v46
                }
            case 32768:
                var v47 uint = 255
                _ = v47
            default:
                var v48 uint64 = 4294967295
                _ = v48
            }
        }
    default:
        fmt.Println(v15)
    }
    if (!true) {
        if (((((!((v15 & (v15 - v15)) <= v15)) || ((((float32(0.1) - 1.5) * 3.14159) / 1.0) != 0.0)) || (-2.25 >= 1e10)) && false) || (2 >= 405)) {
            for fi49 := 0; fi49 < 11; fi49++ {
                _ = len(v3)
                if false {
                    v50 := GenS1{F0: -285, F1: 2, F2: 483}
                    _ = v50
                } else {
                    for fi51 := 0; fi51 < 1; fi51++ {
                        var v52 int = (fi49 / (((fi49 - fi51)) | 1))
                        fmt.Println(fi49, fi51, fi49)
                        _ = fi51
                        _ = v52
                    }
                    delete(v1, 1)
                }
                _ = fi49
            }
            var v53 float64 = 1.5
            _ = v53
        } else {
            v15 = (v15 * (v15 / (((v15 >> ((1) & 63))) | 1)))
        }
        if (899 != 533) {
            switch ("foo bar" + ("hello" + "\n\t")) {
            case "\n\t":
                v15 -= (v15 / ((-638) | 1))
            default:
                fmt.Println(v15)
            }
            for fi54 := 0; fi54 < 18; fi54++ {
                fmt.Println(fi54)
                v55 := GenS1{F0: -296, F1: 816, F2: 65535}
                _ = fi54
                _ = v55
            }
        }
        for fi56 := 0; fi56 < 18; fi56++ {
            v15 = 32767
            v15 += v15
            _ = fi56
        }
    }
    v2 = append(v2, "")
    v57 := GenS1{F0: 1, F1: 65535, F2: 684}
    var v58 uint8 = 42
    if (false || true) {
        v58 += 110
        if ((!((0.1 + -2.25) == 1e10)) && true) {
            v59 := GenS1{F0: 693, F1: 0, F2: 258}
            v15 = v15
            _ = v59
        } else {
            if len(v3) > 0 {
                v3[0] = "a"
            }
        }
    } else {
        if false {
            var v60 uint16 = 292
            _ = v60
        } else {
            for fi61 := 0; fi61 < 14; fi61++ {
                var v62 int64 = -1
                var v63 uint8 = (v58 / (((^(v58))) | 1))
                _ = fi61
                _ = v62
                _ = v63
            }
            for fi64 := 0; fi64 < 9; fi64++ {
                switch fi64 {
                case -9223372036854775808:
                    fmt.Println(v15, v58, v15)
                case 42:
                    _ = len(v3)
                default:
                    for fi65 := 0; fi65 < 16; fi65++ {
                        if false {
                            v66 := GenS1{F0: 1, F1: 65535, F2: 42}
                            v15 = (v15 * v15)
                            var v67 bool = true
                            _ = v66
                            _ = v67
                        } else {
                            if true {
                                if (false || (!true)) {
                                    v3 = append(v3, "\n\t")
                                    var v68 int32 = -534
                                    var v69 uint64 = 787
                                    _ = len(v1)
                                    _ = v68
                                    _ = v69
                                }
                                _ = len(v1)
                                v58 = v58
                            } else {
                                fmt.Println(fi64)
                            }
                            v3 = append(v3, "a")
                            v58 -= (v58 | v58)
                        }
                        fmt.Println(fi64, fi65)
                        _ = fi65
                    }
                }
                var v70 float64 = -2.25
                _ = fi64
                _ = v70
            }
            v71 := GenS1{F0: 446, F1: 297, F2: 65535}
            _ = v71
        }
        v72 := GenS2{F0: 42, F1: 0}
        _ = v72
    }
    var v73 uint64 = 151
    switch ("a" + "foo bar") {
    case "":
        switch "foo bar" {
        case "hello":
            var v74 float64 = 1.5
            _ = v74
        }
    default:
        if false {
            if true {
                var v75 string = ""
                _ = v75
            }
            for fi76 := 0; fi76 < 19; fi76++ {
                if len(v2) > 0 {
                    v2[0] = "a"
                }
                v15 += v15
                _ = fi76
            }
        }
    }
    v15 += (v15 << ((0) & 63))
    if true {
        switch "foo bar" {
        case "":
            delete(v1, 42)
        case "a":
            if (true || (4294967295 <= 42)) {
                v77 := GenS1{F0: -1, F1: 255, F2: 0}
                _ = v77
            }
        case "hello":
            _ = len(v2)
        default:
            v15 -= v15
        }
        v1[1] = 127
    }
    if (119 > (^(v58))) {
        v58 -= v58
        v73 = (v73 % ((v73) | 1))
        v78 := GenS3{F0: 42}
        _ = v78
    } else {
        v79 := GenS4{F0: ("hello" + (("a" + ("foo bar" + "a")) + "a")), F1: 679, F2: false}
        var v80 int = 364
        _ = v79
        _ = v80
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v58:", v58)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
    _ = v73
}
