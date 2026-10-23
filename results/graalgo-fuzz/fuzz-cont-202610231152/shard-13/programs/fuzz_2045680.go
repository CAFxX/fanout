package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 string; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 int64 }
func genFn0(p0 int16) uint8 {
    v5 := []int{int(p0), 222, 42}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        switch -680 {
        case 32767:
            var v6 uint32 = 885
            _ = v6
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        fmt.Println(p0)
    }
    v5 = append(v5, 128)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 float32 = 1.5
        for fi9 := 0; fi9 < 15; fi9++ {
            if (!true) {
                var v10 uint8 = 2
                fmt.Println(fi9, p0, fi7)
                var v11 int8 = int8(p0)
                var v12 int16 = p0
                _ = v10
                _ = v11
                _ = v12
            } else {
                p0 = (p0 - (p0 + p0))
                if (-740 != 1) {
                    var v13 string = "\n\t"
                    _ = v13
                } else {
                    if (true || (42 <= 42)) {
                        v8 = float32(3.14159)
                        if true {
                            var v14 int = (fi9 - 4294967295)
                            v14 = fi9
                            v14 = (v14 | fi7)
                            _ = v14
                        } else {
                            var v15 int8 = 42
                            if true {
                                fmt.Println(v8)
                                p0 = (^(p0))
                                switch fi9 {
                                case 737:
                                    fmt.Println(fi9, v8, v15)
                                case 127:
                                    if ((-(p0)) != (p0 / (((-(p0))) | 1))) {
                                        v16 := GenS1{F0: (558 != 1)}
                                        _ = v16
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = (fi9 << ((309) & 63))
                                        }
                                        v17 := GenS1{F0: (0.0 < 1.5)}
                                        _ = v17
                                    }
                                case 32767:
                                    var v18 int32 = 42
                                    _ = v18
                                default:
                                    fmt.Println(fi9)
                                }
                            }
                            var v19 int8 = v15
                            _ = v15
                            _ = v19
                        }
                        switch (fi9 << ((1) & 63)) {
                        case 501:
                            if false {
                                var v20 int64 = 42
                                _ = v20
                            } else {
                                fmt.Println(p0, v8, p0)
                                var v21 float64 = 1.5
                                var v22 int8 = 100
                                _ = v21
                                _ = v22
                            }
                        }
                        if true {
                            v8 = v8
                            v23 := GenS1{F0: true}
                            _ = v23
                        } else {
                            if false {
                                p0 = p0
                                if ((v8 >= 0.0) && false) {
                                    var v24 string = ""
                                    _ = v24
                                } else {
                                    v8 = (v8 - -2.25)
                                    v25 := GenS1{F0: (false && true)}
                                    _ = v25
                                }
                            }
                            v26 := GenS1{F0: true}
                            _ = v26
                        }
                        var v27 bool = ("hello" != "\n\t")
                        _ = v27
                    } else {
                        v8 += v8
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if (!false) {
                        var v28 int = (-(fi9))
                        v28 = (-(fi9))
                        _ = v28
                    }
                }
            }
            var v29 bool = (!(1 >= 42))
            _ = fi9
            _ = v29
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v5
    return 255
}
func genFn1() int8 {
    var v30 uint8 = genFn0(32767)
    for fi31 := 0; fi31 < 17; fi31++ {
        v30 = (v30 | (v30 + v30))
        var v32 int16 = -646
        v32 = -1
        _ = fi31
        _ = v32
    }
    _ = v30
    return 71
}

func main() {
    var v1 int8 = 100
    v2 := []int{32768, 128, -1, -9223372036854775808}
    var v3 bool = (false || (!false))
    var v4 int16 = 1
    _ = len(v2)
    for fi33 := 0; fi33 < 10; fi33++ {
        v34 := GenS1{F0: (false && false)}
        _ = fi33
        _ = v34
    }
    v35 := GenS2{F0: v4, F1: v1}
    if true {
        v36 := GenS3{F0: 364, F1: 9223372036854775807}
        _ = v36
    } else {
        for fi37 := 0; fi37 < 12; fi37++ {
            for fi38 := 0; fi38 < 3; fi38++ {
                if ((float32(((1e10 - 1.5) / (-2.25 - float64(1e10)))) * 3.14159) >= 0.1) {
                    var v39 int64 = 1
                    _ = v39
                }
                if false {
                    if (-128 >= fi37) {
                        v2 = append(v2, -9223372036854775808)
                    }
                    v40 := GenS4{F0: 0, F1: 42}
                    if (fi38 > fi37) {
                        if v3 {
                            switch "\n\t" {
                            case "foo bar":
                                v4 = 42
                            default:
                                fmt.Println(fi38)
                            }
                            if (v4 == 222) {
                                if (((3.14159 * (1e10 / 3.14159)) * 3.14159) > 3.14159) {
                                    if true {
                                        var v41 float64 = -2.25
                                        fmt.Println(fi37, fi37)
                                        _ = v41
                                    }
                                    var v42 uint8 = 85
                                    switch ("hello" + "a") {
                                    case "\n\t":
                                        var v43 bool = ("" == "")
                                        _ = v43
                                    case "hello":
                                        fmt.Println(fi37, v3)
                                    case "a":
                                        var v44 int16 = -32768
                                        _ = v44
                                    default:
                                        switch "foo bar" {
                                        case "hello":
                                            v45 := GenS5{F0: "", F1: 65535, F2: 2}
                                            _ = v45
                                        case "":
                                            _ = len(v2)
                                        case "a":
                                            var v46 int = int(v4)
                                            _ = v46
                                        }
                                    }
                                    var v47 uint16 = 212
                                    var v48 int32 = 42
                                    var v49 uint = 186
                                    _ = v42
                                    _ = v47
                                    _ = v48
                                    _ = v49
                                } else {
                                    v50 := GenS6{F0: v1}
                                    _ = v50
                                }
                            }
                            var v51 float64 = ((0.1 - 3.14159) * 0.0)
                            v52 := GenS3{F0: 301, F1: -885}
                            fmt.Println(v3, v1, v3)
                            var v53 uint = 65535
                            v54 := GenS7{F0: 2147483647, F1: -102}
                            _ = v51
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                        var v55 int64 = 42
                        _ = v55
                    }
                    _ = v40
                } else {
                    var v56 float64 = 3.14159
                    _ = v56
                }
                _ = fi38
            }
            for fi57 := 0; fi57 < 2; fi57++ {
                fmt.Println(v1, fi37)
                switch fi57 {
                case -2147483648:
                    if len(v2) > 0 {
                        v2[0] = (-(fi57))
                    }
                }
                _ = fi57
            }
            _ = fi37
        }
        var v58 uint8 = 1
        _ = v58
    }
    switch "a" {
    case "":
        for fi59 := 0; fi59 < 16; fi59++ {
            v3 = (42 <= 4294967295)
            for fi60 := 0; fi60 < 14; fi60++ {
                v4 = v4
                if true {
                    var v61 string = ""
                    _ = len(v2)
                    _ = v61
                } else {
                    v4 -= (^(v4))
                    v2 = append(v2, fi59)
                    switch ("a" + ("a" + ("a" + ""))) {
                    case "hello":
                        var v62 uint32 = 42
                        _ = v62
                    case "":
                        v3 = v3
                    case "a":
                        if (("foo bar" + (("hello" + "foo bar") + ("hello" + "foo bar"))) > "a") {
                            var v63 int = (-(fi60))
                            v64 := GenS3{F0: v4, F1: 145}
                            _ = v63
                            _ = v64
                        }
                    default:
                        var v65 float64 = 1e10
                        _ = v65
                    }
                    if ((42 <= 9223372036854775807) && v3) {
                        if (476 > 1) {
                            var v66 int64 = 1
                            v1 += v1
                            _ = v66
                        } else {
                            if (v3 && (733 != 65535)) {
                                var v67 float32 = 1e10
                                if len(v2) > 0 {
                                    v2[0] = fi60
                                }
                                _ = v67
                            } else {
                                var v68 uint32 = 2147483647
                                switch fi59 {
                                case 127:
                                    var v69 uint = 807
                                    _ = v69
                                case 128:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case 42:
                                    v68 = (v68 - v68)
                                }
                                v3 = (0.0 == -2.25)
                                _ = v68
                            }
                        }
                        fmt.Println(v1)
                    }
                }
                _ = fi60
            }
            _ = fi59
        }
    case "\n\t":
        for fi70 := 0; fi70 < 5; fi70++ {
            for fi71 := 0; fi71 < 7; fi71++ {
                var v72 int = (fi71 >> ((255) & 63))
                if len(v2) > 0 {
                    v2[0] = v72
                }
                _ = fi71
                _ = v72
            }
            _ = fi70
        }
    case "a":
        v73 := GenS2{F0: (v4 >> ((2) & 63)), F1: (v1 ^ (v1 >> ((65535) & 63)))}
        _ = v73
    default:
        var v74 int64 = 0
        _ = v74
    }
    var v75 bool = (((((v1 >> ((65535) & 63)) < (v1 % ((v1) | 1))) && (0.0 < 0.0)) || v3) || v3)
    var v76 int16 = (v4 & (v4 ^ v4))
    v75 = (v3 || false)
    v1 += genFn1()
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v75
    _ = v76
}
