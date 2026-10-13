package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 int8; F1 int }
func genFn0(p0 uint32, p1 int8, p2 uint64) int32 {
    p2 += p2
    fmt.Println(p1, p2)
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := []uint32{65535, 307, 2147483648}
    v2 := make(map[int]uint32)
    var v3 bool = (!("a" == "a"))
    var v4 string = ""
    v3 = (!(-9223372036854775808 > 127))
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: 1, F1: 1000, F2: (v4 + v4)}
        v7 := GenS2{F0: 1, F1: 32767}
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v4)
    v3 = v3
    if (true && (-396 >= -1)) {
        v2[42] = 231
    } else {
        var v8 uint64 = 18446744073709551615
        if (100 <= 100) {
            for fi9 := 0; fi9 < 14; fi9++ {
                v1 = append(v1, 4294967295)
                v3 = v3
                switch fi9 {
                case -1:
                    if (0 <= 2147483647) {
                        for fi10 := 0; fi10 < 7; fi10++ {
                            v4 = v4
                            v11 := GenS3{F0: 268, F1: genFn0(65535, 42, (^(v8)))}
                            v8 = (^(v8))
                            v8 = v8
                            switch fi10 {
                            case 1:
                                v3 = (42 > v8)
                            case 32768:
                                var v12 int8 = 1
                                _ = v12
                            case -1:
                                if (v3 && (42 == 65535)) {
                                    switch ("hello" + "a") {
                                    case "":
                                        if v3 {
                                            v3 = v3
                                        }
                                    case "a":
                                        v13 := GenS2{F0: 0, F1: 0}
                                        _ = v13
                                    default:
                                        if (-556 == fi9) {
                                            switch ((v4 + "hello") + "foo bar") {
                                            case "hello":
                                                v4 = "hello"
                                            case "foo bar":
                                                delete(v2, (fi10 & 127))
                                            case "\n\t":
                                                if (1 > -1) {
                                                    fmt.Println(v3)
                                                    var v14 uint32 = 2147483648
                                                    v1 = append(v1, (v14 - (v14 * (^(v14)))))
                                                    _ = v14
                                                } else {
                                                    if v3 {
                                                        var v15 uint8 = 154
                                                        _ = v15
                                                    }
                                                    if (-1 == 9223372036854775807) {
                                                        v4 = v4
                                                        v1 = append(v1, 735)
                                                    }
                                                    switch fi10 {
                                                    case -565:
                                                        _ = len(v2)
                                                    default:
                                                        v4 = v4
                                                    }
                                                }
                                            }
                                            if v3 {
                                                if (!((1.5 + (float64((0.1 * 1e10)) / ((0.0 + 1e10) * -2.25))) > ((float64((float32(float64((0.0 / (0.0 - 1.5)))) * 1.5)) - (0.0 * float64(1.5))) * (0.1 - 1.5)))) {
                                                    v16 := GenS4{F0: 42, F1: 42}
                                                    _ = v16
                                                } else {
                                                    v17 := GenS1{F0: (v8 << ((959) & 63)), F1: (v8 << ((4294967295) & 63)), F2: v4}
                                                    fmt.Println(v8, v4, v4)
                                                    _ = v17
                                                }
                                                switch fi10 {
                                                case 0:
                                                    v4 = ((v4 + (("hello" + v4) + v4)) + "foo bar")
                                                case -2147483648:
                                                    v8 += 759
                                                    if (((!v3) && ((2147483648 < 1) || (0 >= -460))) && v3) {
                                                        var v18 int8 = -128
                                                        _ = v18
                                                    } else {
                                                        v3 = (-80 < 127)
                                                    }
                                                }
                                                var v19 float32 = (float32(((((float64((0.1 - 0.0)) / 1.0) * 1.5) - 0.1) + 1e10)) * (1.5 * float32(((0.1 - -2.25) - (1.5 + (((float64(1e10) - 0.1) / 3.14159) / (-2.25 - float64((1e10 + float32(1.5))))))))))
                                                v3 = (127 == 42)
                                                _ = v19
                                            }
                                            if true {
                                                fmt.Println(v3, v3)
                                                if (!true) {
                                                    var v20 uint32 = 14
                                                    v20 = v20
                                                    v21 := GenS5{F0: 1000, F1: (v20 | v20)}
                                                    fmt.Println(v8, fi10, v3)
                                                    var v22 float64 = 3.14159
                                                    _ = v20
                                                    _ = v21
                                                    _ = v22
                                                }
                                            }
                                        } else {
                                            fmt.Println(v4, v3, fi10)
                                            var v23 uint8 = 2
                                            v8 += 2
                                            _ = v23
                                        }
                                    }
                                    v24 := GenS6{F0: 42, F1: 994}
                                    _ = v24
                                } else {
                                    if false {
                                        v25 := v2[(-(fi9))]
                                        _ = v25
                                    } else {
                                        var v26 uint64 = (v8 & (v8 % ((2) | 1)))
                                        v3 = ((!(553 != 42)) && (-128 < 0))
                                        fmt.Println(fi9, v4, fi10)
                                        _ = v26
                                    }
                                }
                            }
                            _ = fi10
                            _ = v11
                        }
                    } else {
                        switch v4 {
                        case "a":
                            fmt.Println(v8)
                        case "":
                            v3 = (false && v3)
                        case "hello":
                            if (!(422 != 9223372036854775807)) {
                                _ = len(v2)
                                if v3 {
                                    var v27 float32 = (3.14159 / (0.1 - 3.14159))
                                    var v28 string = "hello"
                                    var v29 uint64 = (v8 / (((v8 | (v8 + v8))) | 1))
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                if v3 {
                                    v1 = append(v1, 42)
                                    for fi30 := 0; fi30 < 5; fi30++ {
                                        if ((v4 + v4) <= "") {
                                            v31 := GenS1{F0: 837, F1: 551, F2: (v4 + (v4 + ("hello" + (((((v4 + (v4 + ((v4 + ((((v4 + "foo bar") + v4) + v4) + v4)) + ((v4 + "foo bar") + "foo bar")))) + (v4 + ("" + "\n\t"))) + ("a" + v4)) + "") + v4))))}
                                            var v32 string = v4
                                            _ = v31
                                            _ = v32
                                        }
                                        fmt.Println(fi30, fi9, v4)
                                        _ = fi30
                                    }
                                }
                            }
                        default:
                            v8 = (v8 >> ((65535) & 63))
                        }
                        v4 = "a"
                        switch "hello" {
                        case "\n\t":
                            for fi33 := 0; fi33 < 17; fi33++ {
                                if ((0.0 + 3.14159) == 1.5) {
                                    switch 0 {
                                    case 95:
                                        v34 := GenS3{F0: 858, F1: int32(fi33)}
                                        _ = v34
                                    }
                                    v35 := GenS7{F0: 8, F1: 2147483647}
                                    _ = v35
                                }
                                fmt.Println(v4)
                                var v36 int32 = int32(v8)
                                _ = fi33
                                _ = v36
                            }
                        default:
                            var v37 uint8 = 1
                            _ = v37
                        }
                    }
                default:
                    switch v4 {
                    case "foo bar":
                        switch ((((v4 + ("foo bar" + (v4 + v4))) + "a") + v4) + v4) {
                        case "a":
                            v4 = v4
                        }
                    case "a":
                        for fi38 := 0; fi38 < 13; fi38++ {
                            v8 = 468
                            fmt.Println(v4, v8, v4)
                            _ = fi38
                        }
                    case "\n\t":
                        v39 := GenS5{F0: -363, F1: 1}
                        _ = v39
                    }
                }
                _ = fi9
            }
            _ = len(v1)
        }
        _ = v8
    }
    if (916 == 1) {
        var v40 int = -2147483648
        _ = v40
    } else {
        v41 := GenS8{F0: v4, F1: -1}
        _ = v41
    }
    v4 = v4
    if true {
        switch ("\n\t" + "") {
        case "foo bar":
            var v42 uint8 = 92
            _ = v42
        case "\n\t":
            if (4294967295 == 4294967295) {
                var v43 uint8 = 190
                _ = v43
            }
        default:
            v44 := GenS5{F0: 32767, F1: 390}
            _ = v44
        }
        for fi45 := 0; fi45 < 15; fi45++ {
            v1 = append(v1, 65535)
            v46 := GenS3{F0: 2, F1: 2147483647}
            _ = fi45
            _ = v46
        }
        for fi47 := 0; fi47 < 8; fi47++ {
            var v48 bool = true
            for fi49 := 0; fi49 < 12; fi49++ {
                var v50 int8 = -128
                _ = fi49
                _ = v50
            }
            _ = fi47
            _ = v48
        }
    }
    var v51 uint8 = 220
    switch -934 {
    case -1:
        v52 := GenS7{F0: (v51 / ((146) | 1)), F1: 65535}
        _ = v52
    default:
        var v53 uint = 257
        _ = v53
    }
    if false {
        if (0.0 > (3.14159 - float64(1e10))) {
            if true {
                _ = len(v2)
                v54 := GenS9{F0: 42, F1: -1}
                v4 = v4
                _ = v54
            } else {
                fmt.Println(v4)
                var v55 int16 = 1
                switch -50 {
                case -31:
                    if true {
                        var v56 int32 = genFn0(65535, -1, 303)
                        v57 := GenS2{F0: 751, F1: 1}
                        _ = v56
                        _ = v57
                    }
                case -1:
                    v55 = (v55 / ((v55) | 1))
                case 348:
                    for fi58 := 0; fi58 < 19; fi58++ {
                        var v59 float32 = 0.0
                        var v60 float32 = float32((-2.25 - (-2.25 - 1.5)))
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                }
                _ = v55
            }
            for fi61 := 0; fi61 < 18; fi61++ {
                v62 := GenS8{F0: v4, F1: -1}
                fmt.Println(fi61)
                _ = fi61
                _ = v62
            }
            switch (v4 + (v4 + "")) {
            case "":
                if (true && v3) {
                    var v63 uint = 42
                    _ = v63
                }
            case "\n\t":
                for fi64 := 0; fi64 < 18; fi64++ {
                    v51 = 47
                    _ = fi64
                }
            default:
                v3 = (1 > (v51 ^ v51))
            }
        } else {
            for fi65 := 0; fi65 < 12; fi65++ {
                v4 = v4
                v4 = ("a" + "\n\t")
                v66 := GenS6{F0: (v51 >> ((42) & 63)), F1: 42}
                _ = fi65
                _ = v66
            }
        }
        v67 := GenS9{F0: 94, F1: 0}
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch ((("hello" + ("a" + v4)) + v4) + v4) {
            case "hello":
                var v68 uint = 212
                _ = v68
            case "\n\t":
                var v69 uint = 255
                _ = v69
            case "":
                v2[9223372036854775807] = 2147483647
            }
        } else {
            v4 = v4
            v51 -= (v51 & 32)
            if v3 {
                var v70 float64 = ((((1.5 - (1.5 + float64((-2.25 * 3.14159)))) + float64(3.14159)) / -2.25) * 0.1)
                _ = v70
            } else {
                if v3 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if ("a" >= v4) {
                        if ("foo bar" < v4) {
                            v4 = ((("a" + (v4 + v4)) + ((v4 + v4) + ((v4 + v4) + v4))) + v4)
                        }
                    } else {
                        var v71 int32 = genFn0(42, -1, uint64(v51))
                        _ = v71
                    }
                }
            }
        }
        _ = v67
    } else {
        v72 := GenS3{F0: 0, F1: 1}
        _ = v72
    }
    v73 := GenS4{F0: 66, F1: 32767}
    if ("foo bar" == ("" + v4)) {
        v3 = (0 > 4294967295)
    }
    if ((^(v51)) >= v51) {
        var v74 int16 = 1000
        v3 = ((v74 <= -556) || (1 >= 32768))
        _ = v74
    } else {
        v75 := GenS2{F0: 1, F1: 32767}
        _ = v75
    }
    var v76 uint8 = (v51 * (v51 >> ((0) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v73
    _ = v76
}
