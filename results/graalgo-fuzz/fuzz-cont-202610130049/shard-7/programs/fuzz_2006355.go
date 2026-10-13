package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int32; F2 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int; F1 uint; F2 uint8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    var v3 string = ("\n\t" + ("foo bar" + "a"))
    v4 := make(map[uint32]string)
    v5 := []string{"foo bar"}
    v6 := []uint32{596, 4294967295, 2147483648, 42}
    v3 = (("a" + ("hello" + v3)) + "hello")
    var v7 uint64 = 379
    fmt.Println(v3, v7)
    if (((v3 + "\n\t") + (v3 + "hello")) == v3) {
        var v8 int8 = 113
        var v9 int16 = -588
        v9 -= (v9 * (v9 / ((-661) | 1)))
        _ = v8
        _ = v9
    } else {
        v7 += (v7 | v7)
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(fi10)
            _ = fi10
        }
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS1{F0: 9223372036854775807}
        _ = fi11
        _ = v12
    }
    fmt.Println(v3, v3)
    v3 = v3
    if (!(127 < 15)) {
        v13 := GenS2{F0: (v7 << ((812) & 63)), F1: (((!(true && false)) || (864 == -2147483648)) || true)}
        v3 = v3
        _ = v13
    }
    switch 42 {
    case 128:
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := GenS2{F0: (v7 + 0), F1: true}
            var v16 int8 = 0
            var v17 float32 = (1e10 + -2.25)
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    }
    v3 = (((v3 + ("foo bar" + v3)) + v3) + ((v3 + "\n\t") + v3))
    var v18 uint8 = 0
    if len(v5) > 0 {
        v5[0] = v3
    }
    delete(v1, 2)
    if true {
        for fi19 := 0; fi19 < 14; fi19++ {
            if false {
                v18 = 128
            }
            if ((!false) || false) {
                switch fi19 {
                case 0:
                    for fi20 := 0; fi20 < 7; fi20++ {
                        var v21 int64 = 42
                        var v22 uint32 = 2147483647
                        var v23 uint16 = 147
                        _ = len(v2)
                        fmt.Println(v3, v7)
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                }
                for fi24 := 0; fi24 < 5; fi24++ {
                    if (-2.25 > 0.0) {
                        if (-2.25 != 3.14159) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v2[fi19] = true
                            v3 = v3
                        }
                        v7 += 889
                    } else {
                        v18 += (v18 & 128)
                        v3 = ("a" + (("\n\t" + v3) + v3))
                    }
                    _ = fi24
                }
            } else {
                v7 += v7
                switch fi19 {
                case 665:
                    switch 42 {
                    case 1:
                        v25 := GenS3{F0: 2}
                        _ = v25
                    default:
                        var v26 float64 = float64(0.1)
                        _ = v26
                    }
                case -134:
                    var v27 int16 = 32767
                    _ = v27
                default:
                    var v28 int64 = 0
                    _ = v28
                }
                var v29 uint32 = 0
                _ = v29
            }
            if true {
                for fi30 := 0; fi30 < 10; fi30++ {
                    if ((!(408 >= 65535)) && true) {
                        v31 := GenS4{F0: v18, F1: 319, F2: "hello"}
                        fmt.Println(fi19, v7)
                        if len(v6) > 0 {
                            v6[0] = 823
                        }
                        switch "a" {
                        case "foo bar":
                            if true {
                                if true {
                                    switch 780 {
                                    case 246:
                                        fmt.Println(fi30)
                                    case 32768:
                                        if false {
                                            v3 = v3
                                            switch (^(fi19)) {
                                            case 2147483647:
                                                v3 = "foo bar"
                                            case -744:
                                                v3 = v3
                                            default:
                                                v7 += (v7 - 42)
                                            }
                                            var v32 int16 = 1000
                                            _ = v32
                                        } else {
                                            switch v3 {
                                            case "hello":
                                                switch "a" {
                                                case "\n\t":
                                                    var v33 uint64 = v7
                                                    _ = v33
                                                }
                                            case "":
                                                var v34 int16 = 42
                                                _ = v34
                                            case "a":
                                                var v35 bool = true
                                                _ = v35
                                            default:
                                                _ = len(v4)
                                            }
                                            if (2 != 65535) {
                                                v18 = v18
                                                var v36 int8 = 12
                                                _ = v36
                                            }
                                            switch fi30 {
                                            case 42:
                                                v7 += (^(v7))
                                            default:
                                                v3 = v3
                                            }
                                            var v37 int64 = int64(fi30)
                                            v38 := GenS5{F0: 6}
                                            _ = v37
                                            _ = v38
                                        }
                                    case 589:
                                        if false {
                                            v4[2147483648] = "a"
                                            _ = len(v1)
                                            v7 = v7
                                        }
                                    default:
                                        var v39 float64 = 3.14159
                                        _ = v39
                                    }
                                    var v40 float64 = 1.5
                                    var v41 bool = (v7 <= 42)
                                    v42 := GenS6{F0: 2147483648, F1: -1}
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                                v18 = 2
                            } else {
                                var v43 bool = true
                                fmt.Println(fi19, fi19)
                                if (3.14159 != 3.14159) {
                                    delete(v4, 1)
                                    v44 := GenS7{F0: 42, F1: (v3 + ((("\n\t" + v3) + (v3 + v3)) + v3))}
                                    var v45 uint16 = 630
                                    v7 = 940
                                    v46 := v4[0]
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                } else {
                                    v47 := v1[2147483647]
                                    v48 := GenS8{F0: 255}
                                    var v49 uint = 948
                                    _ = v47
                                    _ = v48
                                    _ = v49
                                }
                                _ = v43
                            }
                        case "a":
                            var v50 float32 = -2.25
                            _ = v50
                        }
                        v51 := GenS3{F0: 2}
                        _ = v31
                        _ = v51
                    }
                    _ = fi30
                }
                v7 = 740
                fmt.Println(v18)
            }
            _ = fi19
        }
        for fi52 := 0; fi52 < 17; fi52++ {
            v18 = v18
            _ = fi52
        }
    } else {
        var v53 int16 = 42
        switch "a" {
        case "\n\t":
            for fi54 := 0; fi54 < 18; fi54++ {
                for fi55 := 0; fi55 < 13; fi55++ {
                    if len(v6) > 0 {
                        v6[0] = 0
                    }
                    if true {
                        v53 = (v53 * -32768)
                        v53 -= (v53 + v53)
                    } else {
                        var v56 int64 = -9223372036854775808
                        v7 -= (^(v7))
                        if len(v5) > 0 {
                            v5[0] = ("" + (((v3 + v3) + v3) + "a"))
                        }
                        _ = v56
                    }
                    v53 += v53
                    _ = fi55
                }
                fmt.Println(v53)
                _ = fi54
            }
        case "a":
            if ((-2147483648 < 1) || (((0.1 - 0.1) > float32((0.1 - ((0.1 / 3.14159) * (1.5 / (0.1 * 0.1)))))) && false)) {
                if (false && (((255 < 18446744073709551615) && (!false)) && (-1 <= 24))) {
                    if true {
                        for fi57 := 0; fi57 < 12; fi57++ {
                            if true {
                                v18 = v18
                            }
                            v5 = append(v5, "a")
                            _ = fi57
                        }
                    } else {
                        fmt.Println(v53, v3, v3)
                        if (!(false || (42 <= 255))) {
                            v58 := GenS4{F0: 176, F1: -1, F2: ((("\n\t" + (v3 + (v3 + ""))) + (v3 + v3)) + v3)}
                            for fi59 := 0; fi59 < 1; fi59++ {
                                _ = len(v6)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if (true || false) {
                                    _ = len(v4)
                                    var v60 uint = 0
                                    for fi61 := 0; fi61 < 9; fi61++ {
                                        var v62 string = "hello"
                                        _ = fi61
                                        _ = v62
                                    }
                                    _ = v60
                                }
                                _ = fi59
                            }
                            _ = v58
                        }
                        switch -9223372036854775808 {
                        case 1:
                            var v63 int32 = 42
                            _ = v63
                        case 9223372036854775807:
                            v3 = v3
                        case -1:
                            fmt.Println(v18, v3, v7)
                        default:
                            fmt.Println(v53, v18, v7)
                        }
                    }
                    v18 = v18
                    var v64 float32 = 1.5
                    _ = v64
                } else {
                    fmt.Println(v18, v53, v3)
                    switch ("a" + "foo bar") {
                    case "foo bar":
                        v65 := GenS3{F0: 42}
                        _ = v65
                    default:
                        switch 1 {
                        case -979:
                            if len(v5) > 0 {
                                v5[0] = v3
                            }
                        case 32767:
                            var v66 bool = true
                            _ = v66
                        }
                    }
                    var v67 uint32 = 2
                    _ = v67
                }
                _ = len(v4)
            } else {
                var v68 int32 = 42
                var v69 string = "\n\t"
                _ = v68
                _ = v69
            }
        default:
            if true {
                fmt.Println(v7, v18, v7)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            } else {
                var v70 int32 = -1
                if (!true) {
                    for fi71 := 0; fi71 < 1; fi71++ {
                        _ = len(v1)
                        _ = fi71
                    }
                    v72 := GenS8{F0: 4294967295}
                    for fi73 := 0; fi73 < 15; fi73++ {
                        var v74 uint = 0
                        for fi75 := 0; fi75 < 8; fi75++ {
                            var v76 int = fi73
                            delete(v1, 0)
                            _ = fi75
                            _ = v76
                        }
                        for fi77 := 0; fi77 < 8; fi77++ {
                            v78 := GenS9{F0: -1, F1: v74, F2: v18}
                            v79 := GenS1{F0: -1}
                            _ = fi77
                            _ = v78
                            _ = v79
                        }
                        _ = fi73
                        _ = v74
                    }
                    _ = v72
                }
                v70 -= v70
                _ = v70
            }
        }
        v7 = 1
        _ = v53
    }
    fmt.Println(v18, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
