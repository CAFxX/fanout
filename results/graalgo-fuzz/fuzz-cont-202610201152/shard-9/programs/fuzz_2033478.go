package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int; F1 int32; F2 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 bool; F2 string }
type GenS8 struct { F0 uint8; F1 int32; F2 int8 }
type GenS9 struct { F0 uint8; F1 uint16; F2 bool }
type GenS10 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    v4 := []int{9223372036854775807, 114}
    var v5 uint8 = 128
    switch "foo bar" {
    case "a":
        var v6 float64 = 0.0
        _ = v6
    }
    if (!false) {
        switch 32768 {
        case 1:
            if false {
                var v7 uint64 = 1
                _ = len(v1)
                v8 := GenS1{F0: v5, F1: -1}
                _ = v7
                _ = v8
            }
        case -568:
            var v9 float64 = (float64(v3) * float64(((v3 / (v3 + v3)) + (0.1 / (float32(float64(-2.25)) - v3)))))
            _ = v9
        case 32768:
            var v10 int64 = 0
            _ = v10
        default:
            if true {
                fmt.Println(v3, v5)
            }
        }
        fmt.Println(v3, v3)
    }
    v11 := GenS1{F0: (v5 - 127), F1: -24}
    for fi12 := 0; fi12 < 8; fi12++ {
        delete(v1, ("a" + ("\n\t" + "foo bar")))
        v5 += (^(v5))
        _ = fi12
    }
    if (0.0 > (1e10 * (0.0 - 3.14159))) {
        fmt.Println(v5)
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint = 18446744073709551615
            _ = fi13
            _ = v14
        }
    }
    var v15 string = "a"
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            v3 -= 3.14159
            fmt.Println(v5, v3)
            fmt.Println(v3, fi17, v5)
            fmt.Println(fi17, v3)
            _ = fi17
        }
        switch (fi16 & fi16) {
        case 46:
            if ((true || (2 == 1)) || (fi16 >= fi16)) {
                var v18 int64 = 42
                v3 = v3
                v15 = v15
                _ = v18
            }
        default:
            for fi19 := 0; fi19 < 17; fi19++ {
                switch 0 {
                case 128:
                    switch fi19 {
                    case 127:
                        fmt.Println(v15)
                    default:
                        var v20 int16 = 1
                        _ = v20
                    }
                case -275:
                    var v21 int = 798
                    v2[(fi16 - 2147483647)] = (v15 + ("\n\t" + ((v15 + ((v15 + (v15 + (v15 + "a"))) + ((v15 + v15) + v15))) + v15)))
                    v22 := GenS2{F0: 0, F1: 65535}
                    _ = v21
                    _ = v22
                default:
                    if (2 >= (v5 | v5)) {
                        if len(v4) > 0 {
                            v4[0] = 128
                        }
                        fmt.Println(fi16)
                        v3 -= (v3 * v3)
                    }
                }
                var v23 int8 = 66
                switch fi16 {
                case 32768:
                    switch -223 {
                    case 650:
                        if false {
                            v23 = 1
                            v3 = v3
                            v15 = v15
                            v24 := GenS2{F0: 19, F1: 821}
                            var v25 uint16 = 32767
                            _ = v24
                            _ = v25
                        }
                    default:
                        v3 += v3
                    }
                default:
                    switch fi19 {
                    case -192:
                        if false {
                            if false {
                                v23 = v23
                                v15 = ("a" + v15)
                                if (!true) {
                                    var v26 float64 = -2.25
                                    _ = v26
                                } else {
                                    var v27 float32 = float32(0.1)
                                    v28 := GenS2{F0: 1, F1: 0}
                                    var v29 int32 = 42
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                if false {
                                    switch 530 {
                                    case 32767:
                                        v23 += (v23 + -1)
                                    case 42:
                                        _ = len(v1)
                                    case 128:
                                        if false {
                                            if false {
                                                v30 := GenS1{F0: (v5 & 128), F1: 127}
                                                v31 := GenS1{F0: 1, F1: v23}
                                                fmt.Println(v5)
                                                _ = v30
                                                _ = v31
                                            }
                                        }
                                    default:
                                        var v32 uint = 4294967295
                                        _ = v32
                                    }
                                    var v33 bool = false
                                    _ = v33
                                } else {
                                    fmt.Println(v15)
                                    v5 = (v5 << ((42) & 63))
                                    var v34 int = (^(fi19))
                                    _ = v34
                                }
                                v5 = 0
                                v15 = ("\n\t" + "foo bar")
                                switch "hello" {
                                case "":
                                    switch "foo bar" {
                                    case "\n\t":
                                        v15 = v15
                                    }
                                case "a":
                                    _ = len(v2)
                                }
                                if (!true) {
                                    var v35 uint = 42
                                    _ = v35
                                } else {
                                    v36 := GenS3{F0: 128, F1: 1, F2: 560}
                                    var v37 uint8 = (v5 & v5)
                                    var v38 uint64 = 329
                                    if (!(!false)) {
                                        var v39 uint32 = 825
                                        var v40 uint16 = 42
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v39
                                        _ = v40
                                    } else {
                                        var v41 int16 = 1000
                                        if (1e10 == -2.25) {
                                            var v42 bool = (true && true)
                                            v38 = v38
                                            v43 := GenS2{F0: 0, F1: 255}
                                            _ = v42
                                            _ = v43
                                        } else {
                                            v3 += (v3 * v3)
                                            switch "foo bar" {
                                            case "\n\t":
                                                if ((v15 + (("foo bar" + v15) + "\n\t")) == "hello") {
                                                    var v44 int = fi16
                                                    v5 = (v5 + uint8(v38))
                                                    _ = v44
                                                } else {
                                                    var v45 int8 = (-(v23))
                                                    _ = v45
                                                }
                                            case "a":
                                                var v46 uint8 = 0
                                                _ = v46
                                            }
                                        }
                                        fmt.Println(v37)
                                        _ = v41
                                    }
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                                fmt.Println(v3)
                                switch v15 {
                                case "hello":
                                    v47 := GenS4{F0: 1}
                                    _ = v47
                                default:
                                    if false {
                                        var v48 uint = 18446744073709551615
                                        fmt.Println(v48, fi19, fi16)
                                        var v49 string = v15
                                        _ = v48
                                        _ = v49
                                    } else {
                                        if true {
                                            v3 -= (v3 / float32(0.1))
                                            fmt.Println(v15, fi16)
                                            delete(v2, (fi16 % (((fi19 & fi16)) | 1)))
                                        }
                                    }
                                }
                            }
                            v15 = "hello"
                            var v50 bool = ((v3 + v3) <= 0.0)
                            _ = v50
                        }
                    case 32768:
                        var v51 uint16 = 255
                        _ = v51
                    case 946:
                        v23 -= (v23 / ((v23) | 1))
                    default:
                        _ = len(v2)
                    }
                }
                _ = fi19
                _ = v23
            }
        }
        v5 = (v5 & 245)
        _ = fi16
    }
    var v52 int32 = 42
    if true {
        v15 = (("foo bar" + "\n\t") + v15)
        v53 := GenS5{F0: 727, F1: 868}
        _ = v53
    }
    v54 := GenS6{F0: 255}
    v55 := v1[v15]
    fmt.Println(v5, v55)
    for fi56 := 0; fi56 < 18; fi56++ {
        v3 -= v3
        _ = fi56
    }
    _ = len(v1)
    for fi57 := 0; fi57 < 5; fi57++ {
        var v58 int32 = (v52 | -2147483648)
        switch fi57 {
        case -859:
            if false {
                var v59 int8 = 113
                var v60 uint = 255
                if false {
                    v59 -= int8(v58)
                    delete(v1, ((("hello" + "\n\t") + "foo bar") + v15))
                    if (1 != 32767) {
                        v61 := GenS7{F0: 255, F1: (int8(fi57) != (v59 ^ (v59 >> ((0) & 63)))), F2: ("" + v15)}
                        for fi62 := 0; fi62 < 2; fi62++ {
                            var v63 int16 = 1000
                            if ((fi57 < 589) && (19 < 42)) {
                                var v64 uint32 = (^(v55))
                                _ = v64
                            } else {
                                v55 -= v55
                            }
                            if (("foo bar" + (v15 + v15)) == v15) {
                                fmt.Println(v60, v63)
                            } else {
                                v65 := GenS8{F0: (^(v5)), F1: int32(v55), F2: v59}
                                _ = len(v1)
                                v66 := GenS9{F0: 2, F1: uint16(v5), F2: (42 != 0)}
                                _ = v65
                                _ = v66
                            }
                            _ = len(v4)
                            switch (fi62 + (fi57 % ((fi62) | 1))) {
                            case -189:
                                var v67 uint16 = 255
                                _ = v67
                            case 1:
                                v55 -= 605
                            case -101:
                                var v68 int32 = 1
                                v69 := GenS10{F0: (v60 ^ (v60 & 18446744073709551615)), F1: -1}
                                v55 -= v55
                                _ = v68
                                _ = v69
                            }
                            _ = fi62
                            _ = v63
                        }
                        v55 -= v55
                        _ = v61
                    }
                } else {
                    if (0 < 797) {
                        v52 -= v58
                    } else {
                        for fi70 := 0; fi70 < 16; fi70++ {
                            if true {
                                v4 = append(v4, fi57)
                                var v71 int8 = v59
                                var v72 uint32 = v55
                                _ = v71
                                _ = v72
                            } else {
                                v59 += v59
                                var v73 bool = ((2 > (^(v5))) && false)
                                _ = v73
                            }
                            v60 = (v60 / ((v60) | 1))
                            var v74 uint64 = 358
                            _ = fi70
                            _ = v74
                        }
                    }
                    if (3.14159 != 0.0) {
                        if len(v4) > 0 {
                            v4[0] = (fi57 % (((-(fi57))) | 1))
                        }
                    } else {
                        for fi75 := 0; fi75 < 4; fi75++ {
                            v59 = v59
                            _ = fi75
                        }
                        var v76 uint16 = uint16(v60)
                        if false {
                            for fi77 := 0; fi77 < 20; fi77++ {
                                var v78 uint8 = v5
                                var v79 uint = v60
                                var v80 int64 = -1
                                _ = fi77
                                _ = v78
                                _ = v79
                                _ = v80
                            }
                            if len(v4) > 0 {
                                v4[0] = (fi57 ^ (^(fi57)))
                            }
                            fmt.Println(v76, v55, v3)
                        } else {
                            switch -444 {
                            case 34:
                                v52 += 2147483647
                            }
                        }
                        _ = v76
                    }
                    if false {
                        v81 := GenS8{F0: (v5 * (v5 & (v5 ^ v5))), F1: (-(v58)), F2: (v59 >> ((255) & 63))}
                        v3 = (1.5 / v3)
                        _ = v81
                    } else {
                        v82 := GenS3{F0: 390, F1: (v52 | int32(fi57)), F2: 1}
                        _ = v82
                    }
                }
                _ = v59
                _ = v60
            } else {
                for fi83 := 0; fi83 < 3; fi83++ {
                    v3 = v3
                    _ = fi83
                }
            }
        default:
            v84 := GenS7{F0: 2, F1: ((-2.25 < 1e10) && false), F2: (("\n\t" + v15) + (v15 + "hello"))}
            _ = v84
        }
        _ = fi57
        _ = v58
    }
    if true {
        v52 += v52
        delete(v2, 199)
        for fi85 := 0; fi85 < 4; fi85++ {
            if true {
                v52 += v52
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                var v86 int = fi85
                _ = v86
            }
            v87 := GenS8{F0: 128, F1: v52, F2: 10}
            v88 := GenS1{F0: (v5 * v5), F1: 100}
            _ = fi85
            _ = v87
            _ = v88
        }
    }
    var v89 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v54
    _ = v55
    _ = v89
}
