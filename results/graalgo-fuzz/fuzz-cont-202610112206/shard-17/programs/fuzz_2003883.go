package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint; F1 int64; F2 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS8 struct { F0 string; F1 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 255
    var v4 float64 = 3.14159
    v5 := []string{"hello", "", ("\n\t" + (("a" + ("\n\t" + ("a" + "a"))) + "\n\t"))}
    if true {
        for fi6 := 0; fi6 < 17; fi6++ {
            fmt.Println(v4)
            _ = fi6
        }
        fmt.Println(v3, v3)
    } else {
        switch "hello" {
        case "foo bar":
            if (42 > 32767) {
                var v7 int = 42
                for fi8 := 0; fi8 < 7; fi8++ {
                    v9 := GenS1{F0: 127}
                    var v10 float32 = 3.14159
                    v7 = (fi8 % ((v7) | 1))
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = v7
            } else {
                var v11 bool = (float64(3.14159) <= v4)
                v12 := GenS1{F0: (v3 & 88)}
                var v13 int16 = 0
                _ = v11
                _ = v12
                _ = v13
            }
        default:
            var v14 string = "\n\t"
            _ = v14
        }
    }
    var v15 float64 = float64(-2.25)
    fmt.Println(v3, v4, v4)
    v16 := GenS1{F0: v3}
    var v17 int16 = -32768
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            fmt.Println(fi18)
            _ = fi19
        }
        _ = len(v1)
        _ = fi18
    }
    v3 = v3
    var v20 string = "\n\t"
    switch (v20 + "foo bar") {
    case "a":
        v20 = v20
    default:
        _ = len(v1)
    }
    switch "\n\t" {
    case "foo bar":
        v21 := GenS2{F0: 32767, F1: 127}
        _ = v21
    }
    v17 -= v17
    for fi22 := 0; fi22 < 10; fi22++ {
        if true {
            for fi23 := 0; fi23 < 2; fi23++ {
                v20 = v20
                v4 = 0.1
                _ = fi23
            }
        } else {
            if len(v5) > 0 {
                v5[0] = ("hello" + v20)
            }
            switch (fi22 & 0) {
            case -2147483648:
                if false {
                    var v24 uint = 0
                    if (913 <= 728) {
                        switch (("\n\t" + v20) + v20) {
                        case "hello":
                            if false {
                                if true {
                                    v5 = append(v5, "hello")
                                } else {
                                    switch (fi22 - 284) {
                                    case 32768:
                                        var v25 uint8 = (^(v3))
                                        _ = v25
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    for fi26 := 0; fi26 < 8; fi26++ {
                                        if len(v5) > 0 {
                                            v5[0] = "foo bar"
                                        }
                                        delete(v2, 1)
                                        if (0 >= (v3 | (^(v3)))) {
                                            var v27 uint = uint(v3)
                                            _ = v27
                                        } else {
                                            switch (fi22 << ((v24) & 63)) {
                                            case -301:
                                                switch v20 {
                                                case "":
                                                    fmt.Println(fi26)
                                                case "foo bar":
                                                    v28 := GenS3{F0: 65535, F1: 64, F2: 1}
                                                    _ = v28
                                                default:
                                                    fmt.Println(v20)
                                                }
                                            default:
                                                v3 += (v3 / ((35) | 1))
                                            }
                                        }
                                        _ = fi26
                                    }
                                    for fi29 := 0; fi29 < 12; fi29++ {
                                        v30 := GenS1{F0: v3}
                                        _ = fi29
                                        _ = v30
                                    }
                                }
                                v31 := v1[-176]
                                var v32 uint32 = 2
                                _ = v31
                                _ = v32
                            } else {
                                for fi33 := 0; fi33 < 20; fi33++ {
                                    if true {
                                        v2[0] = (true || (false && true))
                                        v34 := GenS3{F0: 4294967295, F1: 925, F2: 42}
                                        v35 := GenS3{F0: 428, F1: 65535, F2: 0}
                                        fmt.Println(v3, v4, v4)
                                        _ = v34
                                        _ = v35
                                    } else {
                                        switch v20 {
                                        case "hello":
                                            if (v3 == (v3 * (v3 + (v3 >> (((v24 & v24)) & 63))))) {
                                                v17 = 652
                                                var v36 uint32 = uint32(v24)
                                                var v37 uint = (v24 * v24)
                                                _ = v36
                                                _ = v37
                                            } else {
                                                if (!((v17 * v17) == -70)) {
                                                    fmt.Println(v20, v24, v24)
                                                }
                                                fmt.Println(fi33)
                                            }
                                        case "":
                                            if (37 == 2147483647) {
                                                switch (fi33 & fi33) {
                                                case -200:
                                                    v38 := GenS2{F0: 781, F1: 42}
                                                    _ = v38
                                                case 2147483647:
                                                    v39 := GenS1{F0: v3}
                                                    _ = v39
                                                }
                                                var v40 int64 = -697
                                                v41 := v2[0]
                                                var v42 uint16 = 65535
                                                if len(v5) > 0 {
                                                    v5[0] = (("foo bar" + v20) + "a")
                                                }
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            }
                                            var v43 int = (-(fi22))
                                            var v44 uint32 = 2147483647
                                            _ = v43
                                            _ = v44
                                        default:
                                            switch -2147483648 {
                                            case 9223372036854775807:
                                                fmt.Println(v20)
                                            }
                                        }
                                    }
                                    if (v17 < v17) {
                                        v45 := GenS4{F0: v24, F1: 875, F2: "foo bar"}
                                        v46 := GenS3{F0: 318, F1: 32767, F2: 42}
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = fi33
                                }
                            }
                        default:
                            for fi47 := 0; fi47 < 3; fi47++ {
                                var v48 int8 = 86
                                switch ((("foo bar" + ("hello" + v20)) + v20) + "\n\t") {
                                case "\n\t":
                                    v4 += v15
                                case "hello":
                                    switch v20 {
                                    case "a":
                                        v49 := v1[fi22]
                                        _ = v49
                                    case "hello":
                                        fmt.Println(v4, fi47)
                                    }
                                default:
                                    var v50 int = (fi47 | (-(fi22)))
                                    v5 = append(v5, v20)
                                    var v51 int64 = -505
                                    _ = v50
                                    _ = v51
                                }
                                v52 := GenS5{F0: v24}
                                _ = fi47
                                _ = v48
                                _ = v52
                            }
                        }
                        var v53 string = (v20 + v20)
                        v54 := GenS2{F0: 0, F1: 127}
                        _ = v53
                        _ = v54
                    }
                    var v55 uint8 = (v3 >> ((v24) & 63))
                    _ = v24
                    _ = v55
                }
            case 4294967295:
                if (42 != 0) {
                    delete(v1, fi22)
                    v4 = (v15 + v4)
                    var v56 int8 = 0
                    _ = v56
                }
            }
        }
        for fi57 := 0; fi57 < 4; fi57++ {
            if (!true) {
                v58 := GenS2{F0: 32767, F1: -1}
                v17 += -32768
                var v59 int = fi57
                _ = v58
                _ = v59
            } else {
                if (v20 != v20) {
                    if len(v5) > 0 {
                        v5[0] = v20
                    }
                    v4 -= v4
                }
            }
            if false {
                var v60 int8 = 0
                var v61 uint16 = 0
                _ = v60
                _ = v61
            }
            _ = fi57
        }
        _ = fi22
    }
    var v62 uint = 65535
    for fi63 := 0; fi63 < 17; fi63++ {
        v3 = (^(v3))
        _ = fi63
    }
    if true {
        var v64 int8 = 100
        if false {
            switch 418 {
            case 13:
                switch (("foo bar" + (v20 + "hello")) + v20) {
                case "foo bar":
                    v64 = (^(v64))
                case "\n\t":
                    if (false && (100 == (v64 >> (((v62 ^ v62)) & 63)))) {
                        switch v20 {
                        case "a":
                            var v65 uint16 = 428
                            _ = v65
                        case "hello":
                            for fi66 := 0; fi66 < 2; fi66++ {
                                v62 -= v62
                                v62 -= 255
                                var v67 int64 = -1
                                _ = fi66
                                _ = v67
                            }
                        case "":
                            fmt.Println(v4)
                        }
                        v68 := GenS6{F0: (v20 + v20)}
                        for fi69 := 0; fi69 < 18; fi69++ {
                            var v70 int32 = 2147483647
                            _ = len(v5)
                            _ = fi69
                            _ = v70
                        }
                        _ = v68
                    } else {
                        v71 := GenS7{F0: (v62 ^ (v62 - (v62 ^ (v62 + v62)))), F1: 882, F2: 1}
                        _ = v71
                    }
                case "":
                    _ = len(v5)
                }
            default:
                v72 := GenS8{F0: ((v20 + "a") + ((((v20 + "foo bar") + v20) + (v20 + v20)) + "foo bar")), F1: 0}
                _ = v72
            }
            if false {
                v5 = append(v5, ("a" + ("a" + v20)))
                switch v20 {
                case "foo bar":
                    var v73 float64 = float64(((float32(0.0) - (3.14159 * -2.25)) * ((-2.25 * 0.0) * 0.1)))
                    _ = v73
                default:
                    var v74 string = ("" + "")
                    _ = v74
                }
                v5 = append(v5, v20)
            }
        } else {
            _ = len(v5)
            _ = len(v5)
            var v75 int = 42
            _ = v75
        }
        _ = v64
    }
    var v76 uint64 = 1
    switch -862 {
    case -2147483648:
        var v77 int8 = 42
        _ = v77
    case 100:
        v17 -= v17
    }
    var v78 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v62:", v62)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v76
    _ = v78
}
