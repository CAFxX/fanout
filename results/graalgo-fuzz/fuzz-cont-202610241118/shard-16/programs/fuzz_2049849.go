package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int; F2 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 string; F2 uint16 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int64; F1 uint64; F2 string }
type GenS10 struct { F0 bool; F1 string }
type GenS11 struct { F0 string; F1 bool }
type GenS12 struct { F0 int32; F1 uint32 }
func genFn0() uint32 {
    v4 := []int{}
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int8 = -4
        if len(v4) > 0 {
            v4[0] = (^(fi5))
        }
        v6 = v6
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 4294967295}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int = (fi8 ^ fi8)
        _ = fi8
        _ = v9
    }
    _ = v4
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint8 = 127
    var v2 uint = 552
    var v3 uint = v2
    if (4294967295 >= v3) {
        v3 -= 239
        var v10 uint16 = 255
        v11 := GenS2{F0: 88}
        _ = v10
        _ = v11
    } else {
        var v12 bool = (-1 < 1)
        var v13 int8 = -1
        v3 += 42
        _ = v12
        _ = v13
    }
    v14 := []int{-9223372036854775808, -907, -186, 270}
    for fi15 := 0; fi15 < 6; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            if true {
                var v17 uint16 = 506
                var v18 uint64 = 91
                _ = v17
                _ = v18
            }
            _ = fi16
        }
        fmt.Println(v3)
        if true {
            v3 -= 972
            fmt.Println(v1)
            for fi19 := 0; fi19 < 12; fi19++ {
                fmt.Println(v1, v3, v1)
                var v20 float64 = float64(0.1)
                if (472 >= -2147483648) {
                    v2 = v3
                } else {
                    if (false || (!true)) {
                        v1 -= v1
                    } else {
                        var v21 uint64 = 245
                        v1 += v1
                        _ = v21
                    }
                }
                _ = fi19
                _ = v20
            }
        } else {
            v14 = append(v14, -994)
        }
        _ = fi15
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    switch -974 {
    case 32768:
        v14 = append(v14, -827)
    }
    var v22 int64 = 42
    v2 += 918
    if len(v14) > 0 {
        v14[0] = int(v1)
    }
    v3 = (v3 + (v2 % (((v3 | (^(v2)))) | 1)))
    for fi23 := 0; fi23 < 15; fi23++ {
        if ("" != "\n\t") {
            v24 := GenS3{F0: 2147483648, F1: (^(fi23)), F2: uint32(v1)}
            _ = v24
        } else {
            if (128 >= v1) {
                var v25 int16 = 1000
                if true {
                    v25 = -548
                    v3 -= (v3 << ((42) & 63))
                } else {
                    v2 = v3
                }
                _ = v25
            } else {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                var v26 uint32 = 363
                if false {
                    switch "\n\t" {
                    case "foo bar":
                        var v27 int = fi23
                        _ = v27
                    case "hello":
                        if (v1 < (v1 & (v1 / ((118) | 1)))) {
                            v28 := GenS3{F0: v26, F1: fi23, F2: 99}
                            _ = v28
                        } else {
                            switch fi23 {
                            case 1:
                                v14 = append(v14, 2147483647)
                            default:
                                v22 = 1
                            }
                            var v29 uint16 = 1
                            fmt.Println(v2)
                            _ = v29
                        }
                    default:
                        v30 := GenS4{F0: fi23}
                        _ = v30
                    }
                }
                _ = v26
            }
            v2 -= 65535
            if (true && (false && (true && (42 != 504)))) {
                if (false || (v22 < 1)) {
                    if false {
                        if len(v14) > 0 {
                            v14[0] = fi23
                        }
                        v31 := GenS3{F0: 2147483647, F1: (fi23 / ((fi23) | 1)), F2: 19}
                        switch "" {
                        case "":
                            if (0 == -1) {
                                v32 := GenS5{F0: v1, F1: "hello", F2: 0}
                                v2 += 1
                                if false {
                                    v2 = v2
                                } else {
                                    v33 := GenS6{F0: (("a" + "foo bar") + ""), F1: 4294967295}
                                    v34 := GenS5{F0: 158, F1: ("" + ("" + "\n\t")), F2: 255}
                                    switch fi23 {
                                    case -1:
                                        var v35 uint16 = 428
                                        _ = v35
                                    default:
                                        for fi36 := 0; fi36 < 6; fi36++ {
                                            v37 := GenS7{F0: v1}
                                            if len(v14) >= 2 {
                                                v14 = v14[1:len(v14)]
                                            }
                                            _ = fi36
                                            _ = v37
                                        }
                                    }
                                    _ = v33
                                    _ = v34
                                }
                                _ = v32
                            }
                        case "hello":
                            var v38 string = "a"
                            _ = v38
                        }
                        _ = v31
                    }
                }
                if false {
                    for fi39 := 0; fi39 < 2; fi39++ {
                        var v40 int = fi39
                        var v41 uint16 = 458
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    _ = len(v14)
                } else {
                    v3 = (v3 - 4294967295)
                }
            } else {
                switch fi23 {
                case 42:
                    fmt.Println(fi23, v22)
                }
                switch 1 {
                case 32767:
                    var v42 bool = true
                    _ = v42
                case -2147483648:
                    v43 := GenS7{F0: (v1 * v1)}
                    _ = v43
                default:
                    var v44 uint16 = 32767
                    _ = v44
                }
            }
        }
        for fi45 := 0; fi45 < 17; fi45++ {
            var v46 int = (fi23 * fi45)
            var v47 uint16 = 573
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = fi23
    }
    v14 = append(v14, 42)
    if ("foo bar" < "a") {
        if len(v14) > 0 {
            v14[0] = -713
        }
    }
    if false {
        v1 = uint8(v2)
    }
    if (!true) {
        for fi48 := 0; fi48 < 4; fi48++ {
            for fi49 := 0; fi49 < 14; fi49++ {
                fmt.Println(fi48, fi49)
                switch (("hello" + ("foo bar" + "\n\t")) + ((("\n\t" + (("\n\t" + ("hello" + "a")) + "")) + "foo bar") + "\n\t")) {
                case "\n\t":
                    var v50 int32 = 2147483647
                    _ = v50
                case "a":
                    if (-607 <= -1) {
                        fmt.Println(fi48)
                    }
                }
                switch "\n\t" {
                case "\n\t":
                    v51 := GenS8{F0: -9223372036854775808}
                    _ = v51
                case "a":
                    _ = len(v14)
                case "hello":
                    var v52 uint16 = 2
                    _ = v52
                default:
                    v53 := GenS9{F0: (-(v22)), F1: 1, F2: "hello"}
                    _ = v53
                }
                _ = fi49
            }
            switch (("" + (("hello" + ("foo bar" + (("hello" + "\n\t") + ("a" + ("" + ("foo bar" + "hello")))))) + "hello")) + "a") {
            case "\n\t":
                var v54 uint16 = 622
                _ = v54
            case "foo bar":
                if true {
                    if (v1 <= 2) {
                        switch "hello" {
                        case "\n\t":
                            switch "" {
                            case "hello":
                                for fi55 := 0; fi55 < 12; fi55++ {
                                    switch int(v3) {
                                    case -2147483648:
                                        v22 = (v22 & v22)
                                    default:
                                        if ((float32((3.14159 + 3.14159)) - 0.0) < 1.5) {
                                            if (!((false || (v22 >= (-(v22)))) || true)) {
                                                v56 := GenS5{F0: v1, F1: "hello", F2: 518}
                                                fmt.Println(v3, v1, v22)
                                                var v57 int64 = (v22 >> ((v3) & 63))
                                                v58 := GenS9{F0: v22, F1: 2, F2: "a"}
                                                var v59 uint = v3
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                                _ = v59
                                            }
                                            switch -50 {
                                            case 0:
                                                var v60 int8 = 91
                                                _ = v60
                                            }
                                        } else {
                                            switch (fi48 ^ fi48) {
                                            case 716:
                                                v2 = 18446744073709551615
                                            case 1:
                                                v2 -= v2
                                            case 2147483647:
                                                if ("foo bar" > "foo bar") {
                                                    var v61 float64 = 0.1
                                                    var v62 int16 = int16(fi55)
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    if false {
                                                        if len(v14) >= 2 {
                                                            v14 = v14[1:len(v14)]
                                                        }
                                                    } else {
                                                        var v63 int8 = 100
                                                        _ = v63
                                                    }
                                                    v64 := GenS10{F0: (42 <= 2), F1: ""}
                                                    _ = v64
                                                }
                                            default:
                                                v65 := GenS11{F0: "\n\t", F1: true}
                                                _ = v65
                                            }
                                            var v66 uint8 = 70
                                            _ = v66
                                        }
                                    }
                                    _ = fi55
                                }
                            case "foo bar":
                                v1 -= 45
                            default:
                                switch fi48 {
                                case 1:
                                    var v67 int8 = 127
                                    _ = v67
                                }
                            }
                        case "a":
                            v1 = (v1 + (v1 - v1))
                        case "hello":
                            fmt.Println(v1, v2, fi48)
                        default:
                            for fi68 := 0; fi68 < 18; fi68++ {
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                v69 := GenS2{F0: genFn0()}
                                v70 := GenS11{F0: ((("hello" + "hello") + "a") + (((("hello" + (("foo bar" + "") + "")) + "hello") + ((((("" + "foo bar") + (("\n\t" + "") + (("foo bar" + "\n\t") + "foo bar"))) + ("a" + "")) + ("a" + "a")) + ("a" + "foo bar"))) + ("\n\t" + "hello"))), F1: false}
                                _ = fi68
                                _ = v69
                                _ = v70
                            }
                        }
                        for fi71 := 0; fi71 < 6; fi71++ {
                            fmt.Println(v1, v22)
                            v22 = -21
                            var v72 uint8 = v1
                            _ = fi71
                            _ = v72
                        }
                        var v73 int8 = 1
                        _ = v73
                    } else {
                        var v74 uint32 = 770
                        v2 -= 0
                        v22 = (v22 + v22)
                        _ = v74
                    }
                } else {
                    v75 := GenS12{F0: int32(v22), F1: 500}
                    _ = v75
                }
            case "a":
                v1 += (v1 & 1)
            }
            v76 := GenS2{F0: 42}
            _ = fi48
            _ = v76
        }
        v77 := GenS6{F0: ("\n\t" + ""), F1: 42}
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
}
