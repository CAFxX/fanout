package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint; F1 int32; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint16; F2 int }

func main() {
    var v1 bool = false
    var v2 int8 = 100
    v3 := make(map[uint32]uint32)
    var v4 bool = v1
    var v5 int32 = 2147483647
    v6 := []int{0, 127, -1}
    if ((v1 || v4) || v4) {
        var v7 int32 = 2147483647
        var v8 uint64 = 0
        _ = v7
        _ = v8
    } else {
        var v9 uint64 = 0
        if (!(!false)) {
            v2 = 100
            var v10 bool = ((-750 > int16(v5)) && (1 <= 774))
            v11 := GenS1{F0: ""}
            _ = v10
            _ = v11
        }
        _ = v9
    }
    fmt.Println(v1)
    var v12 float32 = 3.14159
    fmt.Println(v4, v12, v12)
    v13 := GenS1{F0: "hello"}
    var v14 bool = v1
    switch 32768 {
    case -612:
        if ((0.0 + 1e10) <= (1.5 - (((1e10 / (0.1 + -2.25)) * 0.1) / 0.1))) {
            v14 = (v14 || v14)
        }
    case 127:
        var v15 int8 = v2
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v3[173] = 849
        if (!(v14 && (3.14159 >= ((-2.25 - (((3.14159 * (((3.14159 * 3.14159) - (1e10 - float64(v12))) - ((0.1 * float64(v12)) + 0.0))) / 1.0) / (((0.0 * (((1e10 - 1.5) / (((((0.1 + 3.14159) / (((3.14159 * 1.5) - 0.0) + (3.14159 + 1.5))) * float64(((1e10 - v12) * 1e10))) - 1e10) - (3.14159 / 3.14159))) + (1e10 - (1.5 * (0.0 + 0.1))))) - 1.5) + (0.1 / 1.0)))) + 0.0)))) {
            var v17 float64 = 1e10
            if true {
                switch "foo bar" {
                case "\n\t":
                    for fi18 := 0; fi18 < 11; fi18++ {
                        v4 = v1
                        var v19 uint64 = 42
                        _ = fi18
                        _ = v19
                    }
                }
                fmt.Println(fi16, v14, fi16)
                for fi20 := 0; fi20 < 2; fi20++ {
                    fmt.Println(v2, v12, v17)
                    switch -489 {
                    case 32768:
                        v21 := GenS2{F0: fi20, F1: "hello"}
                        _ = v21
                    case 0:
                        fmt.Println(v5)
                    case -1:
                        if (289 > 2147483647) {
                            _ = len(v6)
                            if true {
                                var v22 uint32 = 884
                                _ = v22
                            } else {
                                var v23 int8 = -117
                                fmt.Println(v14, v4, v4)
                                var v24 int = (fi16 << ((2) & 63))
                                v4 = false
                                _ = v23
                                _ = v24
                            }
                            v14 = v1
                            if len(v6) > 0 {
                                v6[0] = (-(fi20))
                            }
                        } else {
                            v2 = v2
                            var v25 string = ("a" + ("\n\t" + "a"))
                            v26 := GenS3{F0: 42, F1: v5, F2: ((!v4) || true)}
                            _ = v25
                            _ = v26
                        }
                    default:
                        v27 := GenS3{F0: 18446744073709551615, F1: v5, F2: (!(!(v1 || (!(!true)))))}
                        _ = v27
                    }
                    _ = fi20
                }
            }
            if (921 < 32767) {
                delete(v3, 1)
            }
            _ = v17
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v1 = false
            if (341 == 1) {
                if v1 {
                    v28 := GenS1{F0: "hello"}
                    v29 := GenS4{F0: v2}
                    _ = v28
                    _ = v29
                }
                v1 = v1
                switch "" {
                case "foo bar":
                    switch fi16 {
                    case 0:
                        for fi30 := 0; fi30 < 13; fi30++ {
                            if v4 {
                                v31 := GenS2{F0: fi30, F1: "hello"}
                                delete(v3, 2)
                                _ = v31
                            }
                            if (v2 != v2) {
                                var v32 int = (fi30 / ((fi16) | 1))
                                var v33 bool = v4
                                fmt.Println(v32, fi16)
                                v34 := GenS1{F0: ""}
                                var v35 bool = v14
                                v33 = v35
                                delete(v3, 2147483648)
                                _ = v32
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            _ = fi30
                        }
                    case -2147483648:
                        v4 = ((fi16 % ((fi16) | 1)) == -1)
                    }
                default:
                    if (-2.25 != 3.14159) {
                        v1 = (65535 >= 32767)
                        v36 := GenS3{F0: 23, F1: (v5 / (((v5 | (v5 | v5))) | 1)), F2: (!(false || ("\n\t" < ((((("hello" + "foo bar") + "") + "a") + ("hello" + "")) + "foo bar"))))}
                        var v37 float32 = v12
                        _ = v36
                        _ = v37
                    } else {
                        fmt.Println(v2, v1)
                        fmt.Println(v1, v1, v14)
                    }
                }
            } else {
                fmt.Println(v1, v12)
            }
        }
        v1 = ("a" >= "\n\t")
        _ = fi16
    }
    var v38 int16 = 42
    switch "\n\t" {
    case "hello":
        var v39 bool = v4
        _ = v39
    case "a":
        delete(v3, 2147483647)
    case "\n\t":
        v5 -= (v5 >> ((4294967295) & 63))
    }
    if (!true) {
        switch "" {
        case "":
            var v40 float32 = (v12 * (0.0 + v12))
            _ = v40
        case "hello":
            v1 = (v2 != (v2 & v2))
        case "foo bar":
            if v14 {
                if v14 {
                    for fi41 := 0; fi41 < 13; fi41++ {
                        for fi42 := 0; fi42 < 12; fi42++ {
                            switch "" {
                            case "hello":
                                fmt.Println(v38, v14, v5)
                            case "":
                                v1 = v4
                            }
                            _ = fi42
                        }
                        _ = fi41
                    }
                }
            } else {
                if true {
                    var v43 int16 = v38
                    _ = v43
                } else {
                    if v1 {
                        for fi44 := 0; fi44 < 18; fi44++ {
                            if (684 <= 1) {
                                fmt.Println(v4)
                                for fi45 := 0; fi45 < 20; fi45++ {
                                    v46 := GenS3{F0: 4294967295, F1: v5, F2: (9223372036854775807 == 1)}
                                    if len(v6) > 0 {
                                        v6[0] = (fi44 + (fi45 & 4294967295))
                                    }
                                    _ = fi45
                                    _ = v46
                                }
                                var v47 int64 = 587
                                _ = v47
                            }
                            v48 := GenS3{F0: 1, F1: 2147483647, F2: (int8(v5) >= v2)}
                            _ = fi44
                            _ = v48
                        }
                        for fi49 := 0; fi49 < 12; fi49++ {
                            fmt.Println(v12)
                            v50 := GenS1{F0: "a"}
                            v51 := GenS1{F0: "a"}
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                    } else {
                        v52 := GenS3{F0: 18446744073709551615, F1: v5, F2: true}
                        _ = v52
                    }
                }
                for fi53 := 0; fi53 < 12; fi53++ {
                    v5 -= (v5 % ((v5) | 1))
                    var v54 string = "hello"
                    _ = fi53
                    _ = v54
                }
                if v1 {
                    for fi55 := 0; fi55 < 12; fi55++ {
                        if (0.1 > 1.5) {
                            if true {
                                fmt.Println(v4)
                            }
                            var v56 bool = ((v2 - (v2 - 127)) > v2)
                            v5 = v5
                            _ = v56
                        } else {
                            for fi57 := 0; fi57 < 4; fi57++ {
                                fmt.Println(v4)
                                _ = fi57
                            }
                            for fi58 := 0; fi58 < 9; fi58++ {
                                if (!(464 <= -203)) {
                                    v38 = v38
                                    v5 = v5
                                    v38 = 330
                                    if (!(688 > 2147483648)) {
                                        if true {
                                            v2 -= int8(v5)
                                        }
                                        var v59 string = (("hello" + "hello") + "")
                                        _ = v59
                                    }
                                }
                                if ("" > "foo bar") {
                                    var v60 int = (fi58 % ((0) | 1))
                                    var v61 int64 = -781
                                    _ = v60
                                    _ = v61
                                }
                                v38 = (v38 - (v38 ^ v38))
                                _ = fi58
                            }
                            var v62 string = ("foo bar" + ("hello" + ((("foo bar" + "\n\t") + (((("hello" + "hello") + (("hello" + "") + "hello")) + "") + "foo bar")) + "hello")))
                            _ = v62
                        }
                        v4 = (!v4)
                        _ = fi55
                    }
                } else {
                    var v63 string = "\n\t"
                    var v64 uint64 = 1
                    for fi65 := 0; fi65 < 8; fi65++ {
                        if true {
                            if ((-(v2)) >= v2) {
                                _ = len(v3)
                            }
                            v66 := GenS5{F0: v63, F1: 32767, F2: (fi65 + -435)}
                            _ = v66
                        } else {
                            switch 128 {
                            case 32768:
                                v2 = v2
                            case 127:
                                fmt.Println(v64, v12, v14)
                            default:
                                if (true && v1) {
                                    var v67 float64 = 3.14159
                                    switch v63 {
                                    case "a":
                                        v67 = (v67 + v67)
                                    case "":
                                        if len(v6) > 0 {
                                            v6[0] = 1
                                        }
                                    case "\n\t":
                                        v68 := GenS2{F0: fi65, F1: v63}
                                        _ = v68
                                    }
                                    v69 := GenS5{F0: v63, F1: 1, F2: fi65}
                                    _ = v67
                                    _ = v69
                                }
                            }
                        }
                        if len(v6) > 0 {
                            v6[0] = fi65
                        }
                        _ = fi65
                    }
                    _ = v63
                    _ = v64
                }
            }
        }
        v14 = v1
        fmt.Println(v5)
    } else {
        fmt.Println(v5)
        if false {
            v4 = (181 > 255)
            if v14 {
                v4 = (!v4)
                v4 = (517 >= 704)
            } else {
                var v70 float32 = v12
                _ = v70
            }
            for fi71 := 0; fi71 < 2; fi71++ {
                switch fi71 {
                case 4294967295:
                    if len(v6) > 0 {
                        v6[0] = fi71
                    }
                case 531:
                    v2 += v2
                case 197:
                    v12 = v12
                }
                v72 := GenS3{F0: 42, F1: (^(v5)), F2: (("foo bar" >= ("foo bar" + "foo bar")) && v14)}
                _ = fi71
                _ = v72
            }
        } else {
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
            var v73 bool = false
            _ = v73
        }
    }
    if v1 {
        v3[1] = 0
    }
    switch 963 {
    case 2147483647:
        for fi74 := 0; fi74 < 19; fi74++ {
            for fi75 := 0; fi75 < 8; fi75++ {
                fmt.Println(fi75, v12)
                var v76 string = "hello"
                _ = fi75
                _ = v76
            }
            _ = fi74
        }
    default:
        for fi77 := 0; fi77 < 4; fi77++ {
            v38 = v38
            _ = fi77
        }
    }
    var v78 bool = v14
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v78
}
