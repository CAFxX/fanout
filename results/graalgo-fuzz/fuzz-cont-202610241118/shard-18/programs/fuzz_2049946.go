package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 string; F1 int8; F2 uint32 }
type GenS5 struct { F0 int64; F1 uint64; F2 int }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint32) int32 {
    if false {
        var v4 bool = (!(true || false))
        _ = v4
    } else {
        var v5 string = (((((("foo bar" + "a") + "hello") + "\n\t") + ("a" + (("hello" + (("hello" + "\n\t") + "foo bar")) + "foo bar"))) + "\n\t") + (("\n\t" + "\n\t") + (("\n\t" + "\n\t") + "")))
        v5 = ((v5 + v5) + v5)
        _ = v5
    }
    v6 := GenS1{F0: "hello"}
    v7 := GenS1{F0: "\n\t"}
    if ((9223372036854775807 != -1) || (true || false)) {
        v8 := GenS1{F0: ""}
        var v9 int8 = -32
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := []int{(fi10 * fi10), fi10, 32768}
            _ = len(v11)
            p0 = 0
            _ = fi10
            _ = v11
        }
    }
    v12 := GenS2{F0: 18, F1: -1}
    _ = p0
    _ = v12
    _ = v6
    _ = v7
    return -431
}
func genFn1() uint32 {
    v13 := GenS2{F0: 127, F1: -266}
    var v14 string = "\n\t"
    v14 = "a"
    for fi15 := 0; fi15 < 8; fi15++ {
        if false {
            var v16 string = "hello"
            _ = v16
        }
        if true {
            if true {
                for fi17 := 0; fi17 < 19; fi17++ {
                    v14 = (v14 + "hello")
                    _ = fi17
                }
                if true {
                    fmt.Println(fi15, fi15)
                    var v18 int16 = 354
                    if true {
                        for fi19 := 0; fi19 < 3; fi19++ {
                            switch v14 {
                            case "":
                                var v20 int16 = -636
                                _ = v20
                            case "a":
                                if false {
                                    v18 += v18
                                } else {
                                    if true {
                                        v21 := GenS3{F0: 0, F1: 65535}
                                        _ = v21
                                    }
                                }
                            case "hello":
                                v22 := []int{(fi15 ^ fi19)}
                                _ = v22
                            }
                            _ = fi19
                        }
                        for fi23 := 0; fi23 < 11; fi23++ {
                            switch (v14 + ((v14 + "foo bar") + ((v14 + v14) + "foo bar"))) {
                            case "foo bar":
                                v14 = (v14 + v14)
                            case "":
                                v24 := []int{}
                                _ = v24
                            default:
                                if false {
                                    fmt.Println(v14, fi23)
                                    if true {
                                        v25 := []int{(fi23 + (fi15 | (fi15 | int(v18)))), fi15}
                                        if (!true) {
                                            _ = len(v25)
                                            v14 = "a"
                                        } else {
                                            v25 = append(v25, fi15)
                                        }
                                        v18 = (v18 + v18)
                                        v26 := GenS4{F0: (v14 + v14), F1: 42, F2: 42}
                                        var v27 int8 = 100
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                }
                            }
                            fmt.Println(fi23, fi15)
                            _ = fi23
                        }
                        fmt.Println(v14)
                    }
                    _ = v18
                } else {
                    v14 = ("foo bar" + ((v14 + (((((((((v14 + v14) + ("\n\t" + v14)) + (("" + (((((v14 + ("foo bar" + v14)) + "\n\t") + (v14 + ("a" + (v14 + (v14 + v14))))) + ("" + ((v14 + ("foo bar" + v14)) + (("" + "\n\t") + "\n\t")))) + v14)) + "a")) + "a") + v14) + "\n\t") + (((v14 + "foo bar") + ("foo bar" + v14)) + "\n\t")) + (v14 + ("hello" + v14))) + "hello")) + v14))
                    v28 := GenS5{F0: 659, F1: 585, F2: fi15}
                    var v29 int = (fi15 ^ fi15)
                    _ = v28
                    _ = v29
                }
            } else {
                switch 1 {
                case -703:
                    fmt.Println(fi15, v14)
                default:
                    var v30 int = 744
                    _ = v30
                }
                for fi31 := 0; fi31 < 13; fi31++ {
                    v32 := []int{}
                    fmt.Println(fi31)
                    var v33 int32 = -2147483648
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                switch fi15 {
                case 1:
                    v34 := GenS5{F0: -325, F1: 4294967295, F2: (fi15 ^ fi15)}
                    _ = v34
                case 127:
                    v14 = v14
                case 370:
                    var v35 int16 = -155
                    _ = v35
                default:
                    fmt.Println(v14, v14)
                }
            }
            var v36 int8 = 0
            v36 -= (-(v36))
            _ = v36
        }
        v37 := []int{(fi15 + (fi15 & -982))}
        _ = fi15
        _ = v37
    }
    _ = v13
    _ = v14
    return 822
}

func main() {
    var v1 float64 = 1.5
    var v2 uint64 = 1
    var v3 uint32 = 65535
    var v38 uint = 18446744073709551615
    v3 += 249
    for fi39 := 0; fi39 < 20; fi39++ {
        var v40 uint64 = (v2 / ((v2) | 1))
        _ = fi39
        _ = v40
    }
    v41 := []int{4294967295, 0}
    if ((true && (!true)) && true) {
        v42 := GenS1{F0: ("foo bar" + (((("" + "") + ((("foo bar" + "hello") + "a") + (("\n\t" + "foo bar") + "a"))) + "\n\t") + "\n\t"))}
        if (uint8(v38) <= uint8(v3)) {
            switch 127 {
            case 9223372036854775807:
                switch "a" {
                case "\n\t":
                    if (35 != 128) {
                        fmt.Println(v2, v38)
                        for fi43 := 0; fi43 < 16; fi43++ {
                            v41 = append(v41, fi43)
                            _ = fi43
                        }
                        v44 := GenS5{F0: -980, F1: (v2 | v2), F2: -2147483648}
                        _ = v44
                    } else {
                        for fi45 := 0; fi45 < 14; fi45++ {
                            v2 = 593
                            fmt.Println(v3, fi45)
                            _ = fi45
                        }
                    }
                case "a":
                    var v46 int8 = 42
                    _ = v46
                default:
                    var v47 int = 0
                    _ = v47
                }
            case 152:
                var v48 int = 565
                _ = v48
            case -9223372036854775808:
                if (127 != -106) {
                    for fi49 := 0; fi49 < 12; fi49++ {
                        if len(v41) > 0 {
                            v41[0] = (fi49 - int(v2))
                        }
                        _ = fi49
                    }
                }
            default:
                if false {
                    v50 := GenS2{F0: -1, F1: -181}
                    _ = v50
                }
            }
        }
        v1 -= v1
        _ = v42
    }
    v2 = 611
    v1 = v1
    if len(v41) > 0 {
        v41[0] = 128
    }
    if false {
        v51 := GenS1{F0: "hello"}
        var v52 uint = v38
        var v53 int = int(v3)
        _ = v51
        _ = v52
        _ = v53
    }
    var v54 uint = (v38 - v38)
    if true {
        if true {
            _ = len(v41)
            var v55 uint = v38
            if (!true) {
                v56 := GenS5{F0: 1, F1: (v2 >> ((0) & 63)), F2: -572}
                if ((9223372036854775807 == 670) && (!false)) {
                    var v57 int8 = 31
                    v58 := GenS3{F0: 1, F1: 65535}
                    if true {
                        if true {
                            var v59 int32 = 42
                            v59 -= -1
                            _ = v59
                        }
                    } else {
                        switch "hello" {
                        case "foo bar":
                            v60 := GenS6{F0: genFn1()}
                            _ = v60
                        case "\n\t":
                            for fi61 := 0; fi61 < 2; fi61++ {
                                var v62 float32 = 1e10
                                _ = fi61
                                _ = v62
                            }
                        }
                    }
                    _ = v57
                    _ = v58
                } else {
                    v2 += (v2 >> ((255) & 63))
                }
                _ = v56
            }
            _ = v55
        }
        for fi63 := 0; fi63 < 13; fi63++ {
            v41 = append(v41, 9223372036854775807)
            _ = fi63
        }
        if (!((0.1 == (-2.25 / ((3.14159 * 3.14159) + 1.5))) || ((!(128 > 32768)) || (!(true || (42 != v2)))))) {
            if true {
                var v64 int8 = 42
                v3 += 1
                _ = v64
            }
            if false {
                fmt.Println(v1, v54)
                var v65 bool = (1 >= -706)
                for fi66 := 0; fi66 < 1; fi66++ {
                    var v67 int8 = -30
                    _ = fi66
                    _ = v67
                }
                _ = v65
            }
        } else {
            if len(v41) >= 2 {
                v41 = v41[1:len(v41)]
            }
            if true {
                var v68 uint32 = 4294967295
                for fi69 := 0; fi69 < 15; fi69++ {
                    v1 -= (v1 - (v1 * v1))
                    _ = fi69
                }
                fmt.Println(v3, v1)
                _ = v68
            }
        }
    }
    for fi70 := 0; fi70 < 9; fi70++ {
        v38 = (v54 - (v38 + v38))
        _ = fi70
    }
    v71 := GenS6{F0: v3}
    v54 = v54
    v41 = append(v41, -306)
    v1 = (0.0 - 3.14159)
    var v72 int64 = 330
    switch 1 {
    case 9223372036854775807:
        var v73 string = ("hello" + (("hello" + (("hello" + ("hello" + "hello")) + "")) + ("a" + ("foo bar" + ("a" + "\n\t")))))
        _ = v73
    default:
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v54:", v54)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v41
    _ = v54
    _ = v71
    _ = v72
}
