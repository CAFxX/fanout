package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint8 = 176
    var v2 string = "\n\t"
    var v3 float32 = 0.1
    var v4 int = 32768
    v5 := []uint32{1, 42}
    var v6 int8 = 127
    var v7 int = v4
    var v8 string = v2
    if len(v5) > 0 {
        v5[0] = 1
    }
    v7 += 128
    for fi9 := 0; fi9 < 10; fi9++ {
        v8 = "\n\t"
        _ = fi9
    }
    v10 := GenS1{F0: 42, F1: 9223372036854775807, F2: 0}
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(v6)
        v3 = -2.25
        if (1 == 475) {
            var v12 uint16 = 65535
            var v13 float32 = (v3 / -2.25)
            switch "\n\t" {
            case "hello":
                v13 += v3
            default:
                if false {
                    if (("hello" + v8) < (v8 + "hello")) {
                        if ((fi11 / (((v4 - 4294967295)) | 1)) >= (v4 - v4)) {
                            if ((v7 - (fi11 * 32768)) == (fi11 % (((v7 << ((532) & 63))) | 1))) {
                                var v14 uint8 = v1
                                _ = v14
                            } else {
                                fmt.Println(v6)
                            }
                            switch v2 {
                            case "a":
                                v15 := GenS1{F0: 42, F1: (v7 % ((v4) | 1)), F2: 2}
                                _ = v15
                            case "foo bar":
                                for fi16 := 0; fi16 < 18; fi16++ {
                                    var v17 float32 = ((3.14159 / ((v13 * v3) / v3)) * v13)
                                    if false {
                                        var v18 uint = 0
                                        v13 = v17
                                        var v19 int64 = -9223372036854775808
                                        _ = v18
                                        _ = v19
                                    }
                                    if (false && true) {
                                        var v20 int8 = v6
                                        var v21 string = v8
                                        fmt.Println(v13, v13, v17)
                                        _ = v20
                                        _ = v21
                                    } else {
                                        var v22 uint = 4294967295
                                        _ = len(v5)
                                        v12 = v12
                                        _ = v22
                                    }
                                    _ = fi16
                                    _ = v17
                                }
                            }
                            for fi23 := 0; fi23 < 17; fi23++ {
                                switch (^(fi11)) {
                                case 127:
                                    v24 := GenS2{F0: (v6 + v6), F1: 767}
                                    v2 = (v8 + (v8 + v8))
                                    _ = len(v5)
                                    _ = v24
                                case 32768:
                                    switch -2147483648 {
                                    case 660:
                                        switch fi11 {
                                        case 42:
                                            v4 = fi23
                                        case 32768:
                                            if true {
                                                var v25 uint16 = 42
                                                _ = v25
                                            }
                                        default:
                                            switch (v7 << ((0) & 63)) {
                                            case -1:
                                                var v26 float64 = 3.14159
                                                _ = v26
                                            case 869:
                                                if ((((v2 < ((((("hello" + v2) + (("" + v8) + (v2 + "foo bar"))) + (v8 + v8)) + v2) + "hello")) || (((v1 + (v1 >> ((2) & 63))) <= v1) && true)) && false) && ((1.5 != 1.5) || false)) {
                                                    if (!(!false)) {
                                                        if (v2 >= ((v2 + v8) + "")) {
                                                            v7 -= -621
                                                            switch (v4 | (v7 << ((uint(v1)) & 63))) {
                                                            case 32767:
                                                                fmt.Println(v8, v7)
                                                            }
                                                            if false {
                                                                v27 := GenS1{F0: 0, F1: 803, F2: 823}
                                                                var v28 string = ((v2 + v8) + v8)
                                                                _ = v27
                                                                _ = v28
                                                            } else {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                v6 -= (v6 & 127)
                                                                v2 = "foo bar"
                                                            }
                                                        }
                                                        v29 := GenS1{F0: 778, F1: 1, F2: 2}
                                                        fmt.Println(v3, v6)
                                                        v1 -= v1
                                                        _ = len(v5)
                                                        var v30 float64 = (1e10 * 0.1)
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                    fmt.Println(v12, v8)
                                                } else {
                                                    v7 = v7
                                                    switch ((v2 + "\n\t") + v2) {
                                                    case "hello":
                                                        fmt.Println(v6)
                                                    }
                                                    if (65535 <= 2) {
                                                        var v31 int64 = -1
                                                        _ = v31
                                                    }
                                                }
                                            case 434:
                                                v2 = "foo bar"
                                            }
                                        }
                                    case 1:
                                        var v32 float32 = v13
                                        _ = v32
                                    case 2147483647:
                                        switch v2 {
                                        case "foo bar":
                                            if true {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            } else {
                                                switch "" {
                                                case "a":
                                                    v8 = v2
                                                case "foo bar":
                                                    v2 = ""
                                                default:
                                                    v5 = append(v5, 0)
                                                }
                                            }
                                        case "hello":
                                            v33 := GenS1{F0: 399, F1: (v7 >> ((18446744073709551615) & 63)), F2: 455}
                                            _ = v33
                                        default:
                                            v5 = append(v5, 100)
                                            v12 = 692
                                        }
                                    }
                                    v34 := GenS2{F0: 67, F1: 255}
                                    _ = v34
                                }
                                if (42 != 1) {
                                    var v35 int = (^(fi11))
                                    if (42 <= 1) {
                                        if true {
                                            if true {
                                                fmt.Println(v13, v3, v4)
                                                fmt.Println(v7, v2, v1)
                                            }
                                            var v36 int32 = -500
                                            v13 += v13
                                            v35 = (fi23 - v4)
                                            var v37 string = ("\n\t" + v2)
                                            var v38 int = v35
                                            var v39 uint8 = (v1 % ((v1) | 1))
                                            if (true || (!false)) {
                                                v5 = append(v5, 342)
                                            }
                                            var v40 uint16 = 1
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        } else {
                                            if (42 >= 42) {
                                                v35 -= v35
                                                var v41 int32 = 1
                                                _ = v41
                                            }
                                            var v42 uint8 = v1
                                            var v43 float32 = 1.5
                                            _ = v42
                                            _ = v43
                                        }
                                        v7 = 815
                                        var v44 int16 = 12
                                        _ = v44
                                    }
                                    var v45 int16 = 0
                                    _ = v35
                                    _ = v45
                                } else {
                                    v1 = 127
                                    if (false || true) {
                                        if true {
                                            v46 := GenS1{F0: 42, F1: -610, F2: 576}
                                            _ = v46
                                        }
                                    }
                                    v47 := GenS3{F0: 259}
                                    switch v8 {
                                    case "\n\t":
                                        var v48 float64 = float64(0.0)
                                        _ = v48
                                    }
                                    v2 = v2
                                    _ = v47
                                }
                                var v49 bool = true
                                _ = fi23
                                _ = v49
                            }
                        } else {
                            var v50 bool = (1 == (v12 ^ v12))
                            v3 -= (float32(1e10) / 1e10)
                            _ = v50
                        }
                    }
                    v2 = (v8 + "hello")
                }
            }
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    var v51 bool = true
    for fi52 := 0; fi52 < 3; fi52++ {
        var v53 int = v4
        v54 := GenS1{F0: 1, F1: v53, F2: 235}
        _ = fi52
        _ = v53
        _ = v54
    }
    if (v7 <= (v7 << ((2) & 63))) {
        v2 = v2
    } else {
        switch ("a" + ("foo bar" + "hello")) {
        case "hello":
            v51 = (3.14159 != (0.1 + 1e10))
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
        case "":
            v55 := GenS4{F0: -251, F1: 0}
            _ = v55
        default:
            v8 = ((v8 + (v2 + v2)) + "a")
        }
    }
    _ = len(v5)
    if (!v51) {
        v6 = (v6 / ((v6) | 1))
        v4 = v7
        if v51 {
            if ("hello" <= "") {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v56 int8 = (v6 >> ((42) & 63))
                _ = v56
            }
        } else {
            v5 = append(v5, 42)
            v1 = (^(v1))
        }
    }
    var v57 uint64 = 4294967295
    _ = len(v5)
    if v51 {
        v58 := GenS2{F0: (v6 - v6), F1: 1}
        fmt.Println(v3)
        _ = v58
    } else {
        var v59 bool = v51
        for fi60 := 0; fi60 < 18; fi60++ {
            for fi61 := 0; fi61 < 20; fi61++ {
                v5 = append(v5, 687)
                fmt.Println(v1, v4)
                v62 := GenS3{F0: -1}
                _ = fi61
                _ = v62
            }
            v3 += 1.5
            for fi63 := 0; fi63 < 4; fi63++ {
                if len(v5) > 0 {
                    v5[0] = 1
                }
                fmt.Println(v59, v59)
                v51 = ((v59 && (0.1 < 0.1)) || v51)
                _ = fi63
            }
            _ = fi60
        }
        var v64 uint32 = 881
        _ = v59
        _ = v64
    }
    var v65 bool = v51
    v66 := GenS3{F0: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v57
    _ = v6
    _ = v65
    _ = v66
    _ = v7
    _ = v8
}
