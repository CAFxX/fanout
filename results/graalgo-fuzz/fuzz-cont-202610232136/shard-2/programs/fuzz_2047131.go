package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int16; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int16; F1 uint64; F2 int32 }
type GenS7 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS8 struct { F0 int16; F1 bool }
type GenS9 struct { F0 string; F1 string }
type GenS10 struct { F0 uint8; F1 string; F2 uint8 }

func main() {
    v1 := []int{32768, -2147483648}
    v2 := make(map[int]int)
    var v3 bool = (65535 < 0)
    v4 := make(map[uint32]string)
    var v5 string = (("a" + "foo bar") + "a")
    var v6 int = 1
    var v7 int16 = 313
    var v8 float32 = 1.5
    var v9 int32 = 42
    var v10 string = v5
    v8 += v8
    var v11 int16 = (v7 - (v7 + (v7 ^ v7)))
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: (v7 * (-(v11)))}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi14
    }
    var v15 uint64 = 4294967295
    if (42 < 0) {
        var v16 int16 = v11
        _ = v16
    } else {
        for fi17 := 0; fi17 < 5; fi17++ {
            v9 = (v9 % ((v9) | 1))
            switch (fi17 | (^(fi17))) {
            case 95:
                if true {
                    v18 := GenS2{F0: ("hello" + (v5 + v10)), F1: (v11 - -506), F2: (v6 / (((fi17 ^ (v6 + (v6 >> ((493) & 63))))) | 1))}
                    switch ((v10 + (v5 + "\n\t")) + "a") {
                    case "foo bar":
                        for fi19 := 0; fi19 < 11; fi19++ {
                            var v20 bool = (74 == -1)
                            v21 := GenS2{F0: "\n\t", F1: (v11 >> ((19) & 63)), F2: fi19}
                            _ = len(v1)
                            v7 -= (v11 - v11)
                            v15 -= (v15 ^ 18446744073709551615)
                            v22 := GenS1{F0: v11}
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    case "\n\t":
                        for fi23 := 0; fi23 < 13; fi23++ {
                            if (((v11 | int16(v6)) == (v7 & (v7 << ((500) & 63)))) && (0.0 != 1e10)) {
                                switch (fi23 ^ (-(fi17))) {
                                case -9223372036854775808:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                var v24 float32 = 1e10
                                v25 := GenS1{F0: (v7 + v7)}
                                switch v5 {
                                case "\n\t":
                                    v26 := GenS1{F0: v11}
                                    _ = v26
                                }
                                _ = v24
                                _ = v25
                            } else {
                                fmt.Println(v10)
                                if v3 {
                                    switch (^(fi23)) {
                                    case 21:
                                        v11 = (v11 - (v11 - -697))
                                    case 696:
                                        var v27 string = v5
                                        _ = v27
                                    case 15:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        var v28 uint32 = 2
                                        _ = v28
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = -743
                                    }
                                    switch (v6 + (fi23 * (fi17 ^ (v6 | -33)))) {
                                    case -768:
                                        var v29 int16 = v7
                                        _ = v29
                                    default:
                                        var v30 uint64 = v15
                                        _ = v30
                                    }
                                }
                            }
                            if v3 {
                                v6 = 127
                                if (!(v10 < v5)) {
                                    fmt.Println(v15)
                                    switch "\n\t" {
                                    case "\n\t":
                                        if false {
                                            if (4294967295 == 379) {
                                                v31 := GenS3{F0: 2147483647}
                                                _ = v31
                                            }
                                            v32 := GenS1{F0: 32767}
                                            var v33 int64 = 508
                                            var v34 uint64 = v15
                                            v33 = 580
                                            v35 := GenS1{F0: v7}
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        } else {
                                            fmt.Println(v11, fi23)
                                            v36 := GenS4{F0: 1}
                                            var v37 int8 = 100
                                            _ = v36
                                            _ = v37
                                        }
                                    default:
                                        switch (v10 + "foo bar") {
                                        case "":
                                            var v38 float32 = ((-2.25 * v8) / (float32((-2.25 + (1e10 * (-2.25 + 3.14159)))) / v8))
                                            _ = v38
                                        default:
                                            if ((false || v3) || (346 <= 2147483647)) {
                                                var v39 float64 = 0.1
                                                var v40 string = v10
                                                switch fi17 {
                                                case 2147483647:
                                                    var v41 int = (v6 * -116)
                                                    _ = v41
                                                case -1:
                                                    v42 := GenS4{F0: 65535}
                                                    _ = v42
                                                }
                                                _ = v39
                                                _ = v40
                                            } else {
                                                v3 = ((v6 % (((fi23 + 2147483647)) | 1)) < (v6 - (v6 << ((798) & 63))))
                                                v43 := GenS5{F0: v15, F1: 18446744073709551615}
                                                _ = v43
                                            }
                                        }
                                    }
                                }
                            }
                            fmt.Println(v6, v10)
                            _ = fi23
                        }
                    case "hello":
                        switch v10 {
                        case "\n\t":
                            if (v8 > ((v8 / 1e10) * (v8 / v8))) {
                                var v44 uint8 = 36
                                var v45 float32 = v8
                                var v46 string = v5
                                _ = v44
                                _ = v45
                                _ = v46
                            } else {
                                var v47 string = (v5 + v10)
                                _ = v47
                            }
                        case "hello":
                            v48 := GenS3{F0: (-(v9))}
                            _ = v48
                        default:
                            var v49 float32 = v8
                            _ = v49
                        }
                    default:
                        var v50 string = ""
                        _ = v50
                    }
                    _ = v18
                }
            case 42:
                var v51 bool = v3
                _ = v51
            }
            fmt.Println(fi17, v3)
            _ = fi17
        }
        delete(v4, 2147483647)
    }
    for fi52 := 0; fi52 < 1; fi52++ {
        if (22 <= 221) {
            if false {
                v11 = 0
            }
        } else {
            switch ("hello" + v5) {
            case "hello":
                switch ("\n\t" + v10) {
                case "foo bar":
                    var v53 bool = (859 > 42)
                    _ = v53
                }
            }
            switch v10 {
            case "foo bar":
                if true {
                    v8 = v8
                } else {
                    var v54 float64 = 0.0
                    if v3 {
                        var v55 int16 = 0
                        _ = v55
                    } else {
                        v56 := v2[32768]
                        if v3 {
                            fmt.Println(v6, v10)
                        } else {
                            for fi57 := 0; fi57 < 19; fi57++ {
                                if (!(v10 == v10)) {
                                    fmt.Println(fi52)
                                    v54 = v54
                                }
                                v11 -= (v7 & v7)
                                if v3 {
                                    switch v5 {
                                    case "a":
                                        fmt.Println(v11, v7)
                                    }
                                    v58 := GenS6{F0: 1, F1: v15, F2: (v9 | (v9 ^ 948))}
                                    _ = v58
                                }
                                _ = fi57
                            }
                        }
                        _ = v56
                    }
                    fmt.Println(v8)
                    _ = v54
                }
            case "hello":
                for fi59 := 0; fi59 < 2; fi59++ {
                    switch v10 {
                    case "\n\t":
                        var v60 int32 = v9
                        _ = v60
                    case "hello":
                        if (!v3) {
                            delete(v2, 1)
                            var v61 string = v10
                            v4[0] = v5
                            v62 := GenS2{F0: ("" + v5), F1: 0, F2: fi52}
                            _ = v61
                            _ = v62
                        } else {
                            v10 = v5
                        }
                    case "a":
                        v63 := GenS7{F0: 32767, F1: 0, F2: 2147483647}
                        _ = v63
                    default:
                        switch v10 {
                        case "\n\t":
                            switch ("foo bar" + v10) {
                            case "":
                                fmt.Println(v5, v6, v7)
                            case "foo bar":
                                var v64 float64 = (0.0 + -2.25)
                                _ = v64
                            default:
                                if (float64(0.0) > (3.14159 * (((1.5 - 0.0) + (0.1 / (0.1 / 1.5))) / 1.0))) {
                                    switch 4294967295 {
                                    case 953:
                                        var v65 int8 = -29
                                        _ = v65
                                    }
                                    v7 = v11
                                    v9 += (v9 - (-(v9)))
                                } else {
                                    switch (v10 + v5) {
                                    case "":
                                        if true {
                                            fmt.Println(fi52)
                                        } else {
                                            if false {
                                                v9 = v9
                                            } else {
                                                var v66 float32 = 0.0
                                                _ = v66
                                            }
                                            switch fi59 {
                                            case 128:
                                                v7 += v7
                                            }
                                            v67 := v4[2147483648]
                                            _ = v67
                                        }
                                    case "foo bar":
                                        var v68 int = fi59
                                        _ = v68
                                    case "\n\t":
                                        v6 = fi52
                                    }
                                    var v69 bool = (786 >= 1)
                                    _ = v69
                                }
                            }
                        case "":
                            var v70 uint8 = 147
                            _ = v70
                        case "a":
                            var v71 float64 = float64(0.0)
                            _ = v71
                        }
                    }
                    var v72 int16 = v7
                    _ = fi59
                    _ = v72
                }
            }
            if v3 {
                for fi73 := 0; fi73 < 9; fi73++ {
                    v6 += v6
                    var v74 bool = (2 >= 828)
                    _ = fi73
                    _ = v74
                }
                for fi75 := 0; fi75 < 8; fi75++ {
                    var v76 uint64 = (v15 - v15)
                    v77 := GenS8{F0: -1, F1: (v3 && v3)}
                    _ = fi75
                    _ = v76
                    _ = v77
                }
            }
        }
        if false {
            var v78 int16 = v11
            v79 := GenS9{F0: (("hello" + "hello") + (v5 + v10)), F1: v5}
            _ = v78
            _ = v79
        } else {
            switch (fi52 ^ 637) {
            case 0:
                switch (("foo bar" + v5) + ((v10 + v5) + ((v10 + "\n\t") + v10))) {
                case "\n\t":
                    for fi80 := 0; fi80 < 4; fi80++ {
                        var v81 uint32 = 65535
                        _ = fi80
                        _ = v81
                    }
                }
            case 32768:
                v1 = append(v1, fi52)
            case -1:
                v5 = ((v10 + (v10 + v10)) + v10)
            default:
                v82 := GenS10{F0: 127, F1: (v5 + "hello"), F2: 0}
                _ = v82
            }
            for fi83 := 0; fi83 < 7; fi83++ {
                var v84 uint64 = v15
                var v85 uint16 = 617
                var v86 int = (fi52 >> ((0) & 63))
                _ = fi83
                _ = v84
                _ = v85
                _ = v86
            }
            delete(v2, (v6 | 1))
        }
        for fi87 := 0; fi87 < 11; fi87++ {
            v7 = 32767
            var v88 int64 = -9223372036854775808
            v5 = ("hello" + ("hello" + v10))
            _ = fi87
            _ = v88
        }
        _ = fi52
    }
    v5 = (("foo bar" + v10) + ("hello" + (("\n\t" + ("a" + (v10 + (v5 + ("" + ""))))) + v5)))
    var v89 uint = 857
    v11 = (^(v7))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v89:", v89)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v89
    _ = v9
}
