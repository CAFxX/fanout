package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64 }

func main() {
    v1 := []string{(("foo bar" + ("foo bar" + (("foo bar" + "foo bar") + "a"))) + ("a" + ("a" + (("" + "foo bar") + ("\n\t" + "\n\t")))))}
    var v2 int8 = 68
    var v3 uint16 = 140
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "foo bar")
    }
    switch 4294967295 {
    case 4294967295:
        v4 := GenS1{F0: 1}
        _ = v4
    case 42:
        fmt.Println(v2, v3)
    }
    if (!(!false)) {
        switch ("\n\t" + (("foo bar" + "foo bar") + "foo bar")) {
        case "":
            var v5 uint64 = 522
            _ = v5
        case "\n\t":
            v3 -= 32767
        case "a":
            if false {
                _ = len(v1)
                if len(v1) > 0 {
                    v1[0] = ("hello" + "\n\t")
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        default:
            switch 9223372036854775807 {
            case 220:
                v2 = (-(v2))
            case 128:
                for fi6 := 0; fi6 < 1; fi6++ {
                    if (!true) {
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        for fi7 := 0; fi7 < 3; fi7++ {
                            if true {
                                var v8 uint16 = v3
                                v9 := GenS1{F0: -2147483648}
                                v2 = v2
                                _ = v8
                                _ = v9
                            } else {
                                fmt.Println(fi7, fi6, v2)
                                var v10 string = ""
                                var v11 uint32 = 66
                                var v12 int8 = (v2 >> ((2) & 63))
                                _ = v10
                                _ = v11
                                _ = v12
                            }
                            _ = fi7
                        }
                        if false {
                            switch "hello" {
                            case "foo bar":
                                var v13 uint16 = 65535
                                _ = v13
                            case "a":
                                for fi14 := 0; fi14 < 12; fi14++ {
                                    var v15 float64 = 3.14159
                                    v3 += (v3 | (v3 & v3))
                                    var v16 int32 = 2147483647
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                            }
                            v2 = (^(v2))
                        } else {
                            v3 = v3
                            v2 = v2
                        }
                    } else {
                        var v17 uint = 65535
                        _ = v17
                    }
                    v18 := GenS2{F0: "foo bar"}
                    _ = fi6
                    _ = v18
                }
            default:
                var v19 uint64 = 0
                _ = v19
            }
        }
        switch 127 {
        case -942:
            for fi20 := 0; fi20 < 17; fi20++ {
                if false {
                    v3 = v3
                    if ("\n\t" < "hello") {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v21 := GenS1{F0: 1}
                        _ = v21
                    }
                } else {
                    _ = len(v1)
                    v3 = v3
                }
                if true {
                    var v22 uint = 14
                    v2 = v2
                    _ = v22
                } else {
                    var v23 uint8 = 137
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        v24 := GenS1{F0: -492}
                        _ = v24
                    case "hello":
                        if (19 == 42) {
                            v25 := GenS3{F0: v23, F1: 1}
                            var v26 uint8 = 127
                            v27 := GenS2{F0: (("\n\t" + "\n\t") + ("foo bar" + "a"))}
                            _ = v25
                            _ = v26
                            _ = v27
                        } else {
                            v28 := GenS2{F0: ("hello" + ("\n\t" + (("a" + "foo bar") + "\n\t")))}
                            for fi29 := 0; fi29 < 9; fi29++ {
                                v30 := GenS1{F0: int32(v3)}
                                _ = fi29
                                _ = v30
                            }
                            var v31 uint32 = 42
                            _ = v28
                            _ = v31
                        }
                    default:
                        v32 := GenS4{F0: 4294967295, F1: 940}
                        _ = v32
                    }
                    _ = v23
                }
                _ = fi20
            }
        case -654:
            var v33 uint32 = uint32(v3)
            _ = v33
        case 128:
            switch "foo bar" {
            case "foo bar":
                switch "foo bar" {
                case "foo bar":
                    v2 -= v2
                default:
                    for fi34 := 0; fi34 < 1; fi34++ {
                        v2 = (v2 << ((233) & 63))
                        if ((0.1 < 3.14159) && false) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v35 float64 = -2.25
                            v2 += 16
                            _ = v35
                        } else {
                            if false {
                                v36 := GenS1{F0: 0}
                                var v37 int32 = -420
                                _ = v36
                                _ = v37
                            }
                        }
                        v2 += v2
                        _ = fi34
                    }
                }
            default:
                v38 := GenS5{F0: v2}
                _ = v38
            }
        default:
            if (true || (true && (!(4294967295 < 2)))) {
                var v39 int16 = -868
                var v40 uint8 = 0
                var v41 int16 = (v39 ^ (v39 ^ (v39 % ((1000) | 1))))
                _ = v39
                _ = v40
                _ = v41
            }
        }
        if true {
            var v42 float32 = 3.14159
            for fi43 := 0; fi43 < 20; fi43++ {
                v44 := GenS1{F0: 0}
                _ = len(v1)
                if false {
                    for fi45 := 0; fi45 < 7; fi45++ {
                        fmt.Println(v2, v42, fi43)
                        v2 -= v2
                        _ = fi45
                    }
                } else {
                    v3 = 2
                    v2 = (^(v2))
                }
                _ = fi43
                _ = v44
            }
            if false {
                for fi46 := 0; fi46 < 13; fi46++ {
                    fmt.Println(fi46, v2, v3)
                    fmt.Println(v42, v3)
                    _ = fi46
                }
                v42 += -2.25
            } else {
                v42 = (v42 + v42)
                fmt.Println(v2)
            }
            _ = v42
        } else {
            v2 = (v2 % (((v2 % (((v2 + (v2 * (v2 & v2)))) | 1))) | 1))
            for fi47 := 0; fi47 < 17; fi47++ {
                for fi48 := 0; fi48 < 5; fi48++ {
                    if (!(true || (-1 > 9223372036854775807))) {
                        v2 -= -97
                    }
                    v49 := GenS6{F0: 0}
                    v3 -= (v3 - v3)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch ((("\n\t" + "foo bar") + "") + "hello") {
                    case "a":
                        v50 := GenS2{F0: "\n\t"}
                        _ = v50
                    case "hello":
                        switch "hello" {
                        case "\n\t":
                            _ = len(v1)
                        case "a":
                            var v51 int16 = -32768
                            _ = v51
                        case "":
                            if (1 > 18446744073709551615) {
                                if (2 == 0) {
                                    v52 := GenS6{F0: -9223372036854775808}
                                    v2 -= v2
                                    _ = v52
                                } else {
                                    if (v2 < v2) {
                                        v2 -= (v2 + 1)
                                        if false {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            fmt.Println(fi48)
                                            var v53 uint8 = 255
                                            _ = v53
                                        }
                                    } else {
                                        if (fi48 > 2147483647) {
                                            if (!true) {
                                                fmt.Println(fi48)
                                                fmt.Println(fi48, v2, v3)
                                                fmt.Println(v2, v2)
                                                v2 = v2
                                                fmt.Println(fi47, v2, fi47)
                                            }
                                        } else {
                                            var v54 int16 = 42
                                            var v55 int8 = (v2 % ((v2) | 1))
                                            v55 -= -96
                                            _ = v54
                                            _ = v55
                                        }
                                        v56 := GenS5{F0: -49}
                                        fmt.Println(v3, v2)
                                        _ = v56
                                    }
                                    var v57 float32 = 0.1
                                    var v58 uint64 = 377
                                    _ = v57
                                    _ = v58
                                }
                                v2 += -55
                                v2 += v2
                            }
                        }
                    default:
                        v3 += 2
                    }
                    _ = fi48
                    _ = v49
                }
                _ = fi47
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (-2147483648 >= 1) {
        for fi59 := 0; fi59 < 5; fi59++ {
            v2 = v2
            v60 := GenS3{F0: 127, F1: 278}
            _ = fi59
            _ = v60
        }
    } else {
        fmt.Println(v3)
        switch "" {
        case "foo bar":
            v2 -= v2
        case "":
            switch 0 {
            case 0:
                var v61 int8 = v2
                _ = v61
            case -1:
                v1 = append(v1, "hello")
            case -68:
                v1 = append(v1, "a")
            default:
                v62 := GenS4{F0: 42, F1: 5}
                _ = v62
            }
        default:
            for fi63 := 0; fi63 < 17; fi63++ {
                var v64 int8 = v2
                v1 = append(v1, "foo bar")
                var v65 int16 = -32768
                _ = fi63
                _ = v64
                _ = v65
            }
        }
        var v66 int16 = -1
        _ = v66
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    fmt.Println(v2, v3)
    var v67 int = -9223372036854775808
    var v68 float32 = (1.5 / 0.1)
    v2 = (v2 - (v2 / ((127) | 1)))
    if (42 < 1) {
        v2 = v2
        fmt.Println(v3, v3)
        fmt.Println(v68, v68)
    }
    switch (v67 >> ((1) & 63)) {
    case 32767:
        v69 := GenS2{F0: ("\n\t" + "foo bar")}
        _ = v69
    }
    v67 += 339
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v67
    _ = v68
}
