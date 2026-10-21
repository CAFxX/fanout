package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS7 struct { F0 bool; F1 uint; F2 string }
func genFn0(p0 int16, p1 int16, p2 uint16) int64 {
    fmt.Println(p1)
    v5 := GenS1{F0: 727, F1: 42}
    if (1 < 2) {
        if true {
            fmt.Println(p2, p1)
        } else {
            p0 = p0
            fmt.Println(p2, p0)
        }
        for fi6 := 0; fi6 < 13; fi6++ {
            p2 += 42
            v7 := []int{}
            _ = fi6
            _ = v7
        }
    } else {
        var v8 int64 = 42
        _ = v8
    }
    var v9 bool = false
    var v10 uint = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v9
    return -1
}
func genFn1() uint8 {
    switch ((("" + ("foo bar" + "hello")) + "hello") + "foo bar") {
    case "hello":
        if (738 >= 42) {
            for fi11 := 0; fi11 < 11; fi11++ {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := []int{}
                    v14 := GenS1{F0: 18446744073709551615, F1: 65535}
                    var v15 float64 = 3.14159
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                var v16 uint = 4294967295
                switch 760 {
                case -1:
                    v16 = (^(v16))
                case 128:
                    v16 = 352
                case 0:
                    v17 := GenS2{F0: (1.5 < float32(1e10))}
                    _ = v17
                }
                _ = fi11
                _ = v16
            }
        } else {
            var v18 bool = true
            var v19 float64 = 0.1
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 int = fi20
                _ = fi20
                _ = v21
            }
            _ = v18
            _ = v19
        }
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        for fi23 := 0; fi23 < 4; fi23++ {
            if true {
                var v24 uint = 4294967295
                _ = v24
            }
            _ = fi23
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        var v26 float32 = ((float32(((3.14159 * 0.0) + (1e10 - (3.14159 / float64(((1e10 / (-2.25 / float32(3.14159))) / 3.14159)))))) - 3.14159) - 0.0)
        _ = fi25
        _ = v26
    }
    var v27 bool = false
    switch "a" {
    case "":
        v28 := []int{}
        _ = v28
    case "a":
        v27 = v27
    default:
        v29 := GenS3{F0: 18}
        _ = v29
    }
    _ = v27
    return 1
}

func main() {
    var v1 int64 = 590
    var v2 string = ("\n\t" + (("hello" + ("\n\t" + "hello")) + ("foo bar" + "a")))
    var v3 bool = (233 != 255)
    v4 := []string{}
    if v3 {
        for fi30 := 0; fi30 < 4; fi30++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi30
        }
    } else {
        for fi31 := 0; fi31 < 12; fi31++ {
            v2 = ("" + (((v2 + (v2 + (v2 + "foo bar"))) + v2) + v2))
            _ = fi31
        }
        var v32 uint = 0
        _ = v32
    }
    v2 = "hello"
    switch (v2 + ((("foo bar" + ("hello" + (v2 + "a"))) + ("" + "hello")) + ("" + ""))) {
    case "hello":
        fmt.Println(v3, v2, v3)
    default:
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := GenS1{F0: 4294967295, F1: 813}
            _ = fi33
            _ = v34
        }
    }
    v35 := GenS3{F0: 0}
    if ("\n\t" == "foo bar") {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi36 := 0; fi36 < 4; fi36++ {
            for fi37 := 0; fi37 < 4; fi37++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (27 <= -128) {
                    v1 += v1
                    if v3 {
                        var v38 int64 = (v1 >> ((1) & 63))
                        switch ("" + v2) {
                        case "foo bar":
                            if v3 {
                                v3 = v3
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                fmt.Println(v38, v1, v38)
                                switch ((v2 + "\n\t") + "hello") {
                                case "a":
                                    if (42 < 1) {
                                        if len(v4) > 0 {
                                            v4[0] = "hello"
                                        }
                                        var v39 int8 = 121
                                        _ = v39
                                    } else {
                                        if (42 == 956) {
                                            switch v2 {
                                            case "a":
                                                if v3 {
                                                    var v40 uint = 18446744073709551615
                                                    _ = v40
                                                } else {
                                                    v3 = v3
                                                }
                                            case "":
                                                v41 := GenS4{F0: 1000, F1: 255}
                                                _ = v41
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = ""
                                            }
                                            fmt.Println(fi37)
                                        }
                                        v42 := GenS4{F0: 0, F1: 65535}
                                        switch 32768 {
                                        case 9223372036854775807:
                                            v2 = ((v2 + "hello") + v2)
                                        case 4294967295:
                                            var v43 int = fi36
                                            _ = v43
                                        default:
                                            fmt.Println(v2)
                                        }
                                        var v44 int = (fi37 ^ (-(fi37)))
                                        var v45 float64 = 1.5
                                        _ = v42
                                        _ = v44
                                        _ = v45
                                    }
                                case "":
                                    v46 := GenS3{F0: 1}
                                    _ = v46
                                case "foo bar":
                                    var v47 int64 = -1
                                    _ = v47
                                }
                                if (42 >= 2) {
                                    var v48 int32 = -525
                                    v1 += v1
                                    _ = v48
                                }
                                v3 = ((((((0.1 / float64(((1e10 + 0.1) / 0.1))) / 0.1) / 1.5) + ((-2.25 + ((-2.25 + 1.5) + 3.14159)) - ((-2.25 / (3.14159 - 1e10)) / (3.14159 + 1e10)))) * 0.0) != 0.0)
                                if len(v4) > 0 {
                                    v4[0] = "foo bar"
                                }
                                switch ((v2 + ("\n\t" + "a")) + "") {
                                case "\n\t":
                                    if (!(-1 == 0)) {
                                        var v49 float64 = float64((1e10 * 1.5))
                                        _ = v49
                                    } else {
                                        var v50 int8 = 109
                                        switch ("foo bar" + v2) {
                                        case "a":
                                            var v51 uint64 = 1
                                            _ = v51
                                        }
                                        fmt.Println(v2, fi37, v38)
                                        _ = v50
                                    }
                                default:
                                    fmt.Println(fi36, v3)
                                }
                            }
                        case "\n\t":
                            fmt.Println(v1, fi36, fi36)
                        case "hello":
                            if false {
                                if (fi36 > fi37) {
                                    v52 := GenS5{F0: -700}
                                    v53 := GenS3{F0: 127}
                                    v54 := GenS6{F0: 4294967295, F1: 42, F2: 127}
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                                var v55 uint32 = 788
                                v2 = ((("a" + ("" + v2)) + ((("a" + "\n\t") + "\n\t") + "hello")) + v2)
                                var v56 int8 = -1
                                _ = v55
                                _ = v56
                            }
                        default:
                            var v57 float32 = ((1e10 + 1e10) + 1e10)
                            _ = v57
                        }
                        _ = v38
                    } else {
                        var v58 uint = 42
                        _ = v58
                    }
                }
                v1 -= genFn0(439, 1, 0)
                switch (v2 + (v2 + v2)) {
                case "foo bar":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    _ = len(v4)
                }
                v2 = ("\n\t" + v2)
                fmt.Println(v2, v1)
                var v59 int16 = -984
                _ = fi37
                _ = v59
            }
            var v60 uint = 42
            for fi61 := 0; fi61 < 1; fi61++ {
                _ = len(v4)
                v60 += (v60 + v60)
                _ = fi61
            }
            _ = fi36
            _ = v60
        }
        if (312 == 324) {
            var v62 uint8 = 43
            v63 := GenS5{F0: -1}
            _ = v62
            _ = v63
        }
    } else {
        switch -2147483648 {
        case 42:
            switch 1 {
            case -9223372036854775808:
                _ = len(v4)
            case 32767:
                v3 = false
            default:
                v64 := GenS6{F0: 903, F1: 224, F2: 0}
                _ = v64
            }
        default:
            _ = len(v4)
        }
        if (v3 || v3) {
            var v65 int16 = 1000
            var v66 uint64 = 1
            _ = v65
            _ = v66
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v67 := GenS6{F0: 628, F1: 18446744073709551615, F2: 127}
            _ = v67
        }
        if (!true) {
            var v68 int16 = 1000
            _ = v68
        } else {
            var v69 uint8 = genFn1()
            v1 -= (v1 % ((v1) | 1))
            if v3 {
                var v70 uint64 = 897
                v4 = append(v4, v2)
                v69 -= 255
                _ = v70
            } else {
                if ((780 > 42) || v3) {
                    if v3 {
                        for fi71 := 0; fi71 < 7; fi71++ {
                            var v72 float64 = float64(float32(-2.25))
                            _ = fi71
                            _ = v72
                        }
                    } else {
                        switch v2 {
                        case "foo bar":
                            var v73 int = -578
                            _ = v73
                        case "":
                            var v74 string = "hello"
                            _ = v74
                        }
                    }
                }
                v1 -= genFn0(-1, 0, 65535)
            }
            _ = v69
        }
    }
    v3 = v3
    _ = len(v4)
    switch v2 {
    case "":
        for fi75 := 0; fi75 < 13; fi75++ {
            switch (((((v2 + "\n\t") + "foo bar") + "hello") + "foo bar") + v2) {
            case "a":
                v76 := GenS7{F0: v3, F1: 0, F2: (("" + v2) + v2)}
                _ = v76
            }
            for fi77 := 0; fi77 < 16; fi77++ {
                var v78 int = (fi77 | fi75)
                v2 = v2
                v78 -= (fi75 * fi75)
                _ = fi77
                _ = v78
            }
            _ = fi75
        }
    }
    fmt.Println(v3, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
