package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 string; F2 int8 }
type GenS4 struct { F0 int32; F1 bool; F2 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string; F1 string }
type GenS10 struct { F0 uint64; F1 int64; F2 uint32 }
func genFn0(p0 string, p1 uint32, p2 uint32) uint {
    for fi5 := 0; fi5 < 2; fi5++ {
        if true {
            p2 = p2
            var v6 int = fi5
            _ = v6
        }
        var v7 uint8 = 100
        p0 = "\n\t"
        _ = fi5
        _ = v7
    }
    var v8 int32 = -1
    v9 := []int{-305, 1, 128}
    var v10 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 255
}
func genFn1(p0 int) int {
    v11 := GenS1{F0: -9223372036854775808, F1: (!(0.1 != -2.25))}
    v12 := GenS2{F0: 226}
    var v13 float32 = 1.5
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 128
}

func main() {
    var v1 int8 = 127
    var v2 int16 = 42
    var v3 int = 32767
    var v4 uint32 = 2147483647
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(v1, v2)
        _ = fi14
    }
    switch ("hello" + "") {
    case "a":
        v2 = (^(v2))
    case "foo bar":
        v15 := []int{(-(v3))}
        _ = v15
    default:
        fmt.Println(v3, v4, v2)
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 int32 = 2147483647
        var v18 uint8 = 127
        _ = fi16
        _ = v17
        _ = v18
    }
    if true {
        if false {
            var v19 int16 = v2
            _ = v19
        }
    } else {
        v20 := []int{(v3 % ((-2147483648) | 1)), (v3 | 4294967295), (v3 + (v3 - genFn1(int(v1))))}
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        fmt.Println(v1, v3)
        v2 += (v2 | (^(v2)))
        _ = fi21
    }
    fmt.Println(v3, v4)
    for fi22 := 0; fi22 < 6; fi22++ {
        v23 := []int{genFn1((v3 % ((fi22) | 1)))}
        v1 += (v1 ^ v1)
        if true {
            switch v3 {
            case 1:
                v23 = append(v23, fi22)
            case 915:
                v24 := GenS3{F0: 255, F1: ("foo bar" + ("" + "\n\t")), F2: v1}
                _ = v24
            case 42:
                fmt.Println(v3, fi22)
            default:
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
            }
            var v25 uint64 = 152
            _ = v25
        } else {
            switch 847 {
            case -26:
                switch (v3 << ((1) & 63)) {
                case 42:
                    if true {
                        switch "a" {
                        case "\n\t":
                            var v26 string = ("hello" + "a")
                            _ = v26
                        case "foo bar":
                            if (!true) {
                                v27 := GenS4{F0: -747, F1: true, F2: genFn0(("" + ""), v4, v4)}
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                                _ = v27
                            } else {
                                for fi28 := 0; fi28 < 8; fi28++ {
                                    v29 := GenS3{F0: 144, F1: "foo bar", F2: v1}
                                    _ = fi28
                                    _ = v29
                                }
                                if (!false) {
                                    v30 := GenS5{F0: false}
                                    _ = v30
                                } else {
                                    var v31 float64 = 1e10
                                    if len(v23) > 0 {
                                        v23[0] = (v3 & (v3 & fi22))
                                    }
                                    _ = v31
                                }
                            }
                        case "":
                            var v32 int8 = (v1 << ((255) & 63))
                            _ = v32
                        }
                        v33 := GenS6{F0: 0, F1: 32767}
                        _ = v33
                    }
                case -63:
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                case 107:
                    if false {
                        if (-2.25 == 1e10) {
                            var v34 bool = ((((^(v4)) >= v4) || (true && (!(1000 <= (v2 << ((genFn0(((("a" + "hello") + "foo bar") + "\n\t"), 42, v4)) & 63)))))) || false)
                            var v35 uint16 = 540
                            _ = v34
                            _ = v35
                        } else {
                            fmt.Println(v3)
                            var v36 uint8 = 2
                            _ = v36
                        }
                        var v37 int32 = -25
                        fmt.Println(v1, fi22)
                        _ = v37
                    }
                default:
                    v38 := GenS7{F0: 913, F1: 648}
                    _ = v38
                }
            default:
                var v39 uint16 = 319
                _ = v39
            }
            fmt.Println(v1, v1)
            switch "" {
            case "a":
                switch ("hello" + ("hello" + "hello")) {
                case "hello":
                    if true {
                        for fi40 := 0; fi40 < 15; fi40++ {
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            if false {
                                var v41 int = genFn1(v3)
                                _ = v41
                            } else {
                                var v42 int64 = 9223372036854775807
                                if ("foo bar" == "") {
                                    v1 += (v1 / (((-(v1))) | 1))
                                }
                                switch genFn1((v3 & 32768)) {
                                case 32767:
                                    var v43 uint64 = 0
                                    _ = v43
                                case 385:
                                    v3 += 127
                                    v44 := GenS8{F0: 65535}
                                    _ = v44
                                }
                                fmt.Println(v4, v42, v42)
                                if false {
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    var v45 uint = 330
                                    var v46 string = ""
                                    _ = v45
                                    _ = v46
                                }
                                _ = v42
                            }
                            if (((-2.25 - -2.25) / (3.14159 - -2.25)) >= ((3.14159 * ((0.0 + -2.25) / 1e10)) * 0.1)) {
                                var v47 uint8 = 127
                                _ = v47
                            } else {
                                v48 := GenS8{F0: 0}
                                var v49 uint32 = (v4 % ((v4) | 1))
                                _ = v48
                                _ = v49
                            }
                            _ = fi40
                        }
                        v3 += -715
                        switch -38 {
                        case 127:
                            v50 := GenS9{F0: ("" + "foo bar"), F1: "foo bar"}
                            _ = v50
                        case -1:
                            for fi51 := 0; fi51 < 17; fi51++ {
                                if len(v23) > 0 {
                                    v23[0] = (v3 * -173)
                                }
                                v52 := GenS6{F0: 0, F1: 57}
                                var v53 int = (v3 - genFn1((^(fi22))))
                                fmt.Println(fi51, v4)
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                        case 32767:
                            v3 = -2147483648
                        }
                    } else {
                        fmt.Println(v4, v2, fi22)
                        if (((false || false) || (4294967295 <= 431)) && (true || true)) {
                            v4 += (v4 / ((v4) | 1))
                            if (1.5 == 1e10) {
                                var v54 uint64 = 4294967295
                                switch v3 {
                                case -2147483648:
                                    v3 = (v3 % ((32767) | 1))
                                default:
                                    fmt.Println(v4, fi22)
                                }
                                _ = v54
                            }
                            var v55 uint16 = 32767
                            _ = v55
                        }
                        v4 = (v4 << ((65535) & 63))
                    }
                }
            case "hello":
                var v56 uint64 = 406
                _ = v56
            case "foo bar":
                var v57 int64 = -1
                _ = v57
            }
        }
        _ = fi22
        _ = v23
    }
    for fi58 := 0; fi58 < 7; fi58++ {
        for fi59 := 0; fi59 < 12; fi59++ {
            v60 := GenS6{F0: 60, F1: 2}
            if (((-(v2)) != v2) && true) {
                var v61 bool = (-275 <= v3)
                _ = v61
            }
            _ = fi59
            _ = v60
        }
        _ = fi58
    }
    if false {
        var v62 float64 = 1e10
        v63 := GenS8{F0: 464}
        _ = v62
        _ = v63
    } else {
        v64 := []int{}
        switch (^(v3)) {
        case -903:
            if (!(("\n\t" < "") && true)) {
                if true {
                    fmt.Println(v4, v3, v4)
                    fmt.Println(v4)
                } else {
                    v65 := GenS6{F0: 522, F1: 42}
                    _ = v65
                }
                if true {
                    if (229 < 42) {
                        if (((!false) && true) && true) {
                            if false {
                                var v66 int8 = -1
                                _ = v66
                            }
                            v2 = (v2 + (v2 % ((v2) | 1)))
                            switch ("\n\t" + ("\n\t" + "")) {
                            case "\n\t":
                                if false {
                                    fmt.Println(v1, v3, v2)
                                    v1 = (v1 * (-(v1)))
                                } else {
                                    v2 = -284
                                    v1 = (v1 - (v1 | v1))
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(v1, v3, v1)
                                    case "foo bar":
                                        var v67 uint = 290
                                        _ = v67
                                    default:
                                        v68 := GenS1{F0: 455, F1: false}
                                        _ = v68
                                    }
                                }
                            case "foo bar":
                                v64 = append(v64, (^(v3)))
                            }
                        }
                        if true {
                            var v69 bool = false
                            if v69 {
                                fmt.Println(v1, v69, v2)
                                v1 = (v1 << ((390) & 63))
                                _ = len(v64)
                            }
                            _ = v69
                        } else {
                            for fi70 := 0; fi70 < 7; fi70++ {
                                v71 := GenS10{F0: 655, F1: -1, F2: (v4 >> ((2) & 63))}
                                _ = fi70
                                _ = v71
                            }
                        }
                    }
                }
                fmt.Println(v3, v1, v2)
            }
        case -1:
            v1 = v1
        case 32768:
            if false {
                v64 = append(v64, (-(v3)))
                v3 = -2147483648
                fmt.Println(v1, v2, v4)
            }
        default:
            var v72 float64 = -2.25
            _ = v72
        }
        var v73 int = 128
        _ = v64
        _ = v73
    }
    v1 += v1
    switch ("foo bar" + "hello") {
    case "foo bar":
        for fi74 := 0; fi74 < 12; fi74++ {
            v75 := []int{32768, (v3 >> ((255) & 63)), (v3 | (fi74 + (fi74 >> ((genFn0("foo bar", (v4 | v4), v4)) & 63))))}
            _ = fi74
            _ = v75
        }
    case "":
        v3 = v3
    case "a":
        v76 := []int{(^(v3)), genFn1((v3 * 127)), (v3 >> ((999) & 63)), 127}
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
