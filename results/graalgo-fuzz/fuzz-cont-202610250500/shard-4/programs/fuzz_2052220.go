package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 bool; F2 uint8 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int; F1 uint32 }
type GenS9 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint32) int8 {
    fmt.Println(p0)
    v7 := GenS1{F0: 2, F1: (p0 | 65535)}
    var v8 int = 1
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    v3 := []string{(((("foo bar" + "foo bar") + "") + "\n\t") + "")}
    v4 := make(map[int]string)
    v5 := make(map[int]uint32)
    var v6 string = ("a" + "hello")
    if (((!true) || (154 != 127)) || (462 <= 32767)) {
        switch 423 {
        case 32767:
            var v9 bool = true
            _ = v9
        default:
            if (1.5 > float64(0.1)) {
                if (32767 <= 238) {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        if (!(v6 == ((v6 + ((("" + v6) + v6) + v6)) + "a"))) {
                            var v11 bool = true
                            v12 := GenS1{F0: 406, F1: 1}
                            if len(v3) > 0 {
                                v3[0] = (v6 + v6)
                            }
                            _ = v11
                            _ = v12
                        }
                        v13 := GenS2{F0: v6, F1: 42}
                        var v14 int16 = 649
                        _ = fi10
                        _ = v13
                        _ = v14
                    }
                    if false {
                        v15 := GenS1{F0: 209, F1: 65535}
                        for fi16 := 0; fi16 < 9; fi16++ {
                            fmt.Println(v6, fi16)
                            _ = fi16
                        }
                        _ = v15
                    } else {
                        v6 = "a"
                    }
                    fmt.Println(v6)
                }
                for fi17 := 0; fi17 < 6; fi17++ {
                    fmt.Println(fi17)
                    if false {
                        v18 := GenS2{F0: v6, F1: 2}
                        for fi19 := 0; fi19 < 12; fi19++ {
                            _ = len(v2)
                            _ = fi19
                        }
                        switch "" {
                        case "":
                            var v20 uint32 = 19
                            _ = v20
                        default:
                            for fi21 := 0; fi21 < 10; fi21++ {
                                if (v6 > "foo bar") {
                                    delete(v4, fi21)
                                } else {
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(fi17, fi21)
                                    case "":
                                        v6 = (v6 + v6)
                                    case "a":
                                        v3 = append(v3, (v6 + (v6 + (v6 + ((v6 + ("\n\t" + "\n\t")) + v6)))))
                                    default:
                                        switch v6 {
                                        case "":
                                            v6 = "\n\t"
                                        default:
                                            v6 = v6
                                        }
                                    }
                                    v22 := v2[0]
                                    var v23 uint16 = 32767
                                    _ = v22
                                    _ = v23
                                }
                                _ = fi21
                            }
                        }
                        _ = v18
                    }
                    _ = fi17
                }
                if (-55 >= 1) {
                    v6 = v6
                    switch 42 {
                    case -1:
                        var v24 uint = 548
                        _ = v24
                    case -875:
                        var v25 bool = false
                        _ = v25
                    default:
                        if true {
                            delete(v1, 407)
                            v26 := GenS3{F0: -128}
                            _ = v26
                        }
                    }
                    var v27 bool = true
                    _ = v27
                }
            } else {
                if (("foo bar" <= "hello") && true) {
                    v6 = "hello"
                }
                v1[937] = true
                for fi28 := 0; fi28 < 4; fi28++ {
                    if (263 == 96) {
                        for fi29 := 0; fi29 < 17; fi29++ {
                            if ("a" < "a") {
                                v6 = (v6 + v6)
                                var v30 string = v6
                                var v31 float64 = ((3.14159 * (-2.25 * ((0.1 - 1e10) + 1.5))) / (-2.25 + 0.1))
                                if (!(127 <= 0)) {
                                    v6 = (v6 + v30)
                                    v32 := GenS4{F0: 18446744073709551615}
                                    _ = v32
                                }
                                fmt.Println(v6, fi29)
                                _ = v30
                                _ = v31
                            } else {
                                switch ((("hello" + (((v6 + v6) + "") + ((v6 + "hello") + (("" + (v6 + "a")) + v6)))) + v6) + "\n\t") {
                                case "":
                                    if (255 >= 2) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v33 := GenS5{F0: (fi28 + (fi28 >> ((205) & 63))), F1: (((2 <= 1) || (2147483648 > 937)) && (1 > 1)), F2: 0}
                                        v34 := GenS6{F0: 103, F1: 2147483648}
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v35 := GenS6{F0: 94, F1: 505}
                                        if false {
                                            switch v6 {
                                            case "":
                                                v36 := GenS7{F0: 1}
                                                _ = v36
                                            default:
                                                v37 := GenS8{F0: -2147483648, F1: 4294967295}
                                                _ = v37
                                            }
                                            if true {
                                                v6 = ""
                                            }
                                            v6 = v6
                                        } else {
                                            v38 := GenS6{F0: 1, F1: 2147483647}
                                            if (2 != 127) {
                                                var v39 int32 = -113
                                                if true {
                                                    _ = len(v4)
                                                    var v40 bool = false
                                                    _ = v40
                                                }
                                                _ = v39
                                            } else {
                                                if true {
                                                    var v41 float32 = float32(0.1)
                                                    var v42 float32 = 0.0
                                                    _ = v41
                                                    _ = v42
                                                }
                                            }
                                            var v43 uint64 = 1
                                            _ = v38
                                            _ = v43
                                        }
                                        v6 = ("foo bar" + "")
                                        _ = v35
                                    }
                                case "\n\t":
                                    v6 = v6
                                }
                            }
                            if ((-2.25 - (-2.25 - -2.25)) >= 1e10) {
                                var v44 float32 = 1.5
                                v6 = v6
                                _ = v44
                            }
                            _ = len(v3)
                            _ = fi29
                        }
                        if ((148 == 42) && false) {
                            v45 := GenS5{F0: fi28, F1: false, F2: 111}
                            var v46 string = (("hello" + (v6 + "foo bar")) + v6)
                            _ = v45
                            _ = v46
                        } else {
                            v2[65535] = fi28
                            v47 := GenS9{F0: 727, F1: 1, F2: 42}
                            _ = v47
                        }
                        var v48 string = (v6 + (v6 + v6))
                        _ = v48
                    }
                    _ = fi28
                }
            }
        }
    } else {
        var v49 float32 = 3.14159
        v6 = v6
        _ = v49
    }
    switch v6 {
    case "foo bar":
        var v50 float64 = ((3.14159 * 3.14159) * -2.25)
        _ = v50
    case "\n\t":
        for fi51 := 0; fi51 < 6; fi51++ {
            switch (v6 + v6) {
            case "a":
                if (!(2147483647 < 4294967295)) {
                    switch fi51 {
                    case 610:
                        if true {
                            if (!false) {
                                var v52 bool = true
                                v52 = (v52 && v52)
                                fmt.Println(v52, v6)
                                _ = v52
                            } else {
                                if true {
                                    var v53 float64 = 0.1
                                    _ = v53
                                }
                                v54 := v5[fi51]
                                var v55 uint64 = 4294967295
                                _ = v54
                                _ = v55
                            }
                            if false {
                                v6 = (v6 + "")
                            }
                            switch (^(fi51)) {
                            case 9223372036854775807:
                                v6 = v6
                            default:
                                var v56 int16 = -1
                                _ = v56
                            }
                        }
                    default:
                        v57 := GenS6{F0: 196, F1: 180}
                        _ = v57
                    }
                    var v58 uint8 = 2
                    _ = v58
                } else {
                    var v59 float64 = 0.1
                    _ = v59
                }
            case "":
                switch v6 {
                case "a":
                    v6 = "a"
                case "\n\t":
                    var v60 float64 = float64(3.14159)
                    _ = v60
                case "hello":
                    for fi61 := 0; fi61 < 15; fi61++ {
                        v62 := GenS3{F0: -120}
                        _ = fi61
                        _ = v62
                    }
                }
            case "foo bar":
                fmt.Println(fi51)
            }
            fmt.Println(v6, fi51)
            if true {
                v6 = "hello"
                v6 = ((v6 + "foo bar") + "hello")
                fmt.Println(v6, v6)
            }
            _ = fi51
        }
    }
    for fi63 := 0; fi63 < 4; fi63++ {
        _ = len(v5)
        v64 := GenS3{F0: 127}
        var v65 uint16 = 860
        _ = fi63
        _ = v64
        _ = v65
    }
    _ = len(v4)
    var v66 float32 = (3.14159 * 1.5)
    var v67 uint16 = 1
    var v68 uint8 = 235
    v6 = v6
    for fi69 := 0; fi69 < 15; fi69++ {
        v70 := GenS2{F0: (v6 + ""), F1: 0}
        if true {
            if true {
                v67 = v67
                for fi71 := 0; fi71 < 1; fi71++ {
                    v72 := GenS9{F0: 632, F1: 1, F2: 9223372036854775807}
                    var v73 int64 = 42
                    _ = fi71
                    _ = v72
                    _ = v73
                }
            }
            v66 = v66
            if false {
                for fi74 := 0; fi74 < 1; fi74++ {
                    _ = len(v2)
                    v66 = (v66 / 0.1)
                    _ = fi74
                }
                v75 := GenS5{F0: 0, F1: false, F2: (v68 & v68)}
                switch fi69 {
                case -9223372036854775808:
                    v6 = v6
                case -2147483648:
                    delete(v2, 2147483648)
                case -258:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    for fi76 := 0; fi76 < 2; fi76++ {
                        v77 := GenS5{F0: fi76, F1: true, F2: 160}
                        switch v6 {
                        case "\n\t":
                            switch -898 {
                            case -2147483648:
                                if true {
                                    v5[(^(fi69))] = 2147483647
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        if (-2.25 < float64(float32((1e10 * 0.1)))) {
                                            v78 := GenS4{F0: 65535}
                                            _ = v78
                                        }
                                    case "":
                                        if false {
                                            v79 := GenS8{F0: -108, F1: 107}
                                            var v80 int64 = 54
                                            _ = v79
                                            _ = v80
                                        }
                                        if (996 != 9223372036854775807) {
                                            switch 127 {
                                            case -97:
                                                var v81 int64 = -1
                                                _ = v81
                                            case 128:
                                                v68 = 42
                                            case 127:
                                                var v82 int16 = 16
                                                _ = v82
                                            }
                                            v68 -= (v68 << ((18446744073709551615) & 63))
                                        }
                                    case "\n\t":
                                        var v83 bool = ((v6 + "") == (v6 + v6))
                                        _ = v83
                                    default:
                                        if true {
                                            _ = len(v1)
                                        } else {
                                            fmt.Println(v68, fi76)
                                            v66 -= (v66 / v66)
                                            v84 := GenS10{F0: 4294967295}
                                            _ = v84
                                        }
                                    }
                                    v6 = v6
                                    fmt.Println(v67, fi76)
                                }
                                var v85 int16 = 457
                                if (!(((v68 << ((2) & 63)) > v68) || false)) {
                                    v68 -= 0
                                }
                                _ = v85
                            }
                        case "a":
                            v6 = (v6 + v6)
                        default:
                            if true {
                                v86 := GenS11{F0: (0 == 211), F1: 0}
                                _ = v86
                            }
                            _ = len(v3)
                        }
                        fmt.Println(v66)
                        _ = fi76
                        _ = v77
                    }
                }
                _ = v75
            } else {
                var v87 int8 = 58
                fmt.Println(v67, fi69)
                for fi88 := 0; fi88 < 7; fi88++ {
                    if (2147483648 == 2147483647) {
                        var v89 float32 = v66
                        _ = v89
                    }
                    v67 = (v67 % ((316) | 1))
                    delete(v1, uint32(v67))
                    _ = fi88
                }
                _ = v87
            }
        }
        _ = fi69
        _ = v70
    }
    v68 += v68
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v68
}
