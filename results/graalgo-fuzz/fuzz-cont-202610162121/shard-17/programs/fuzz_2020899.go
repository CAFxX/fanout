package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 int16; F2 int64 }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64 }
func genFn0(p0 string) int64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        p0 = ((p0 + "a") + p0)
        switch -87 {
        case 2147483647:
            if false {
                var v8 uint = 1
                v9 := GenS1{F0: v8, F1: 32767}
                fmt.Println(fi7)
                _ = v8
                _ = v9
            }
        default:
            switch 9223372036854775807 {
            case 32768:
                var v10 uint16 = 32767
                _ = v10
            case 1:
                p0 = p0
            default:
                v11 := GenS2{F0: p0, F1: 0}
                _ = v11
            }
        }
        _ = fi7
    }
    var v12 float32 = ((1.5 - 1.5) + 3.14159)
    var v13 uint8 = 128
    _ = p0
    _ = v12
    _ = v13
    return -421
}
func genFn1() uint16 {
    if (true || false) {
        v14 := GenS1{F0: 2, F1: 1000}
        v15 := GenS2{F0: ("foo bar" + (("foo bar" + ("\n\t" + "\n\t")) + "\n\t")), F1: -32768}
        v16 := GenS1{F0: 1, F1: -952}
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 int64 = genFn0("a")
    var v18 uint8 = 42
    if true {
        fmt.Println(v18, v18)
        for fi19 := 0; fi19 < 16; fi19++ {
            if true {
                v20 := GenS1{F0: 812, F1: 42}
                _ = v20
            } else {
                if ("" >= (("hello" + "") + ("hello" + "hello"))) {
                    var v21 int8 = 42
                    switch "\n\t" {
                    case "hello":
                        v22 := GenS1{F0: 724, F1: -1}
                        _ = v22
                    case "a":
                        v23 := []int{fi19}
                        _ = v23
                    default:
                        if false {
                            v24 := GenS1{F0: 65535, F1: 42}
                            var v25 uint8 = v18
                            for fi26 := 0; fi26 < 17; fi26++ {
                                v17 += 42
                                v17 += v17
                                v27 := []int{fi19}
                                _ = fi26
                                _ = v27
                            }
                            _ = v24
                            _ = v25
                        }
                    }
                    v28 := []int{-9223372036854775808, (fi19 % (((fi19 | (fi19 ^ fi19))) | 1))}
                    _ = v21
                    _ = v28
                }
            }
            v18 = (v18 - (^(v18)))
            _ = fi19
        }
        var v29 int = -91
        _ = v29
    }
    _ = v17
    _ = v18
    return 65535
}

func main() {
    var v1 float64 = -2.25
    var v2 uint8 = 225
    var v3 int64 = 0
    v4 := []string{}
    var v5 uint32 = 2147483647
    var v6 int32 = 2147483647
    var v30 uint = 353
    for fi31 := 0; fi31 < 1; fi31++ {
        fmt.Println(v30, v30)
        v30 -= v30
        var v32 uint16 = 0
        _ = fi31
        _ = v32
    }
    if false {
        switch "" {
        case "\n\t":
            if (!(v2 < (^(v2)))) {
                if (!(false || (v5 >= (v5 * v5)))) {
                    switch (((("" + ("foo bar" + ("foo bar" + "hello"))) + "\n\t") + "") + (("hello" + ("hello" + "")) + ("" + "foo bar"))) {
                    case "\n\t":
                        for fi33 := 0; fi33 < 9; fi33++ {
                            v34 := GenS1{F0: v30, F1: 1}
                            if false {
                                for fi35 := 0; fi35 < 5; fi35++ {
                                    switch ("\n\t" + ((("hello" + ("a" + "hello")) + (("" + ("a" + ("a" + ""))) + "hello")) + "")) {
                                    case "a":
                                        v5 -= v5
                                    case "\n\t":
                                        var v36 int16 = -394
                                        _ = v36
                                    default:
                                        fmt.Println(v6)
                                    }
                                    var v37 int16 = -410
                                    var v38 string = ("hello" + ((("\n\t" + ("hello" + "hello")) + ("a" + "a")) + "a"))
                                    _ = fi35
                                    _ = v37
                                    _ = v38
                                }
                            } else {
                                fmt.Println(v30, fi33)
                            }
                            _ = fi33
                            _ = v34
                        }
                    case "a":
                        switch ("a" + ((("a" + ("hello" + "\n\t")) + "") + "foo bar")) {
                        case "foo bar":
                            v39 := GenS3{F0: 18446744073709551615}
                            _ = v39
                        case "":
                            if false {
                                var v40 string = ("\n\t" + "hello")
                                v5 -= v5
                                if (!(true && true)) {
                                    v41 := GenS4{F0: (v30 - v30)}
                                    fmt.Println(v6, v5, v30)
                                    v42 := GenS2{F0: ("\n\t" + v40), F1: -997}
                                    _ = v41
                                    _ = v42
                                }
                                _ = v40
                            }
                        case "hello":
                            var v43 string = ("hello" + ("foo bar" + "foo bar"))
                            _ = v43
                        default:
                            var v44 int16 = 1000
                            _ = v44
                        }
                    }
                    v45 := GenS2{F0: "", F1: 720}
                    _ = v45
                }
                if (65535 == 42) {
                    switch "a" {
                    case "":
                        for fi46 := 0; fi46 < 18; fi46++ {
                            var v47 int16 = 1
                            for fi48 := 0; fi48 < 15; fi48++ {
                                var v49 int8 = 54
                                var v50 string = ("hello" + "hello")
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 15; fi51++ {
                                var v52 uint32 = 2147483647
                                _ = len(v4)
                                _ = len(v4)
                                _ = fi51
                                _ = v52
                            }
                            _ = fi46
                            _ = v47
                        }
                    default:
                        v1 -= v1
                    }
                }
                v3 -= (v3 - int64(v2))
            } else {
                _ = len(v4)
            }
        case "":
            var v53 uint64 = 2
            _ = v53
        case "hello":
            var v54 uint16 = 2
            _ = v54
        default:
            if true {
                if true {
                    _ = len(v4)
                    v55 := GenS3{F0: 0}
                    _ = v55
                }
                fmt.Println(v1, v1)
                for fi56 := 0; fi56 < 8; fi56++ {
                    switch 982 {
                    case -1:
                        var v57 string = ("foo bar" + "")
                        _ = v57
                    }
                    var v58 bool = false
                    _ = fi56
                    _ = v58
                }
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        for fi59 := 0; fi59 < 7; fi59++ {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            _ = fi59
        }
        fmt.Println(v5, v3, v1)
    } else {
        for fi60 := 0; fi60 < 7; fi60++ {
            var v61 uint = v30
            v61 += v30
            if true {
                v30 += 255
            } else {
                switch ("a" + "hello") {
                case "":
                    if false {
                        fmt.Println(v2, v30, v5)
                    }
                case "a":
                    for fi62 := 0; fi62 < 19; fi62++ {
                        var v63 int64 = (v3 & 0)
                        fmt.Println(v1)
                        _ = fi62
                        _ = v63
                    }
                }
            }
            _ = fi60
            _ = v61
        }
    }
    v64 := GenS4{F0: 4294967295}
    for fi65 := 0; fi65 < 2; fi65++ {
        var v66 uint16 = uint16(v5)
        _ = fi65
        _ = v66
    }
    fmt.Println(v30)
    var v67 int32 = 2147483647
    if true {
        v5 += (v5 / (((^(v5))) | 1))
        if false {
            if (65535 != v5) {
                fmt.Println(v6)
                switch "a" {
                case "hello":
                    var v68 float32 = 0.1
                    _ = v68
                case "foo bar":
                    fmt.Println(v1)
                case "a":
                    v69 := GenS5{F0: -1}
                    _ = v69
                default:
                    for fi70 := 0; fi70 < 8; fi70++ {
                        for fi71 := 0; fi71 < 2; fi71++ {
                            _ = len(v4)
                            _ = fi71
                        }
                        _ = fi70
                    }
                }
            }
            for fi72 := 0; fi72 < 15; fi72++ {
                if (801 >= -1) {
                    v73 := GenS6{F0: ("\n\t" + ""), F1: int16(v5), F2: (v3 << ((842) & 63))}
                    var v74 int64 = (v3 | v3)
                    v75 := GenS7{F0: (^(fi72)), F1: (!(v5 <= v5))}
                    _ = v73
                    _ = v74
                    _ = v75
                } else {
                    var v76 int16 = 32767
                    for fi77 := 0; fi77 < 15; fi77++ {
                        if true {
                            v6 += (^(v6))
                            fmt.Println(v6, v76, fi77)
                            var v78 int32 = v6
                            _ = v78
                        } else {
                            v67 -= v67
                            var v79 int64 = int64(v2)
                            v6 += v67
                            _ = v79
                        }
                        if true {
                            var v80 uint64 = 2
                            if true {
                                var v81 int32 = v6
                                if true {
                                    var v82 uint16 = 42
                                    _ = v82
                                }
                                _ = v81
                            } else {
                                v3 = genFn0(("foo bar" + "a"))
                            }
                            v76 -= -803
                            _ = v80
                        }
                        _ = fi77
                    }
                    v4 = append(v4, "")
                    _ = v76
                }
                if true {
                    for fi83 := 0; fi83 < 18; fi83++ {
                        if len(v4) > 0 {
                            v4[0] = "\n\t"
                        }
                        _ = fi83
                    }
                }
                _ = fi72
            }
        } else {
            var v84 string = "a"
            if (genFn0(v84) >= (v3 - 0)) {
                switch 128 {
                case -2147483648:
                    for fi85 := 0; fi85 < 5; fi85++ {
                        var v86 int = 4294967295
                        var v87 float32 = 1.5
                        _ = fi85
                        _ = v86
                        _ = v87
                    }
                case -459:
                    var v88 uint32 = 388
                    _ = v88
                default:
                    var v89 uint32 = 4294967295
                    _ = v89
                }
            } else {
                for fi90 := 0; fi90 < 11; fi90++ {
                    v4 = append(v4, "foo bar")
                    switch -1 {
                    case 749:
                        for fi91 := 0; fi91 < 2; fi91++ {
                            var v92 uint8 = v2
                            v2 = (v92 >> ((75) & 63))
                            _ = fi91
                            _ = v92
                        }
                    }
                    v93 := GenS5{F0: 42}
                    _ = fi90
                    _ = v93
                }
                if ((v2 * v2) <= 233) {
                    v94 := GenS5{F0: -98}
                    var v95 int8 = 127
                    v96 := GenS8{F0: int64(v2)}
                    _ = v94
                    _ = v95
                    _ = v96
                }
            }
            if false {
                v5 += (v5 ^ v5)
                if (!true) {
                    for fi97 := 0; fi97 < 17; fi97++ {
                        fmt.Println(v2, fi97, v2)
                        v30 += v30
                        _ = fi97
                    }
                } else {
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                    v5 = (v5 << ((913) & 63))
                    _ = len(v4)
                }
                var v98 float32 = 0.1
                _ = v98
            }
            _ = v84
        }
    }
    v3 -= (-(v3))
    _ = len(v4)
    _ = len(v4)
    var v99 uint64 = 2
    if false {
        if false {
            var v100 int = 32767
            for fi101 := 0; fi101 < 5; fi101++ {
                v102 := GenS9{F0: 0}
                switch (-(v100)) {
                case 32767:
                    v30 -= v30
                default:
                    fmt.Println(v2, v30, v5)
                }
                _ = fi101
                _ = v102
            }
            _ = v100
        }
        for fi103 := 0; fi103 < 16; fi103++ {
            if true {
                var v104 int16 = int16(v67)
                _ = v104
            }
            _ = fi103
        }
    } else {
        _ = len(v4)
    }
    v1 = 3.14159
    switch "hello" {
    case "\n\t":
        var v105 uint8 = 235
        _ = v105
    case "foo bar":
        v30 = 284
    case "":
        for fi106 := 0; fi106 < 9; fi106++ {
            var v107 uint64 = (v99 + (v99 % ((568) | 1)))
            var v108 uint8 = 99
            _ = fi106
            _ = v107
            _ = v108
        }
    }
    v5 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v67
    _ = v99
}
