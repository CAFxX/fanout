package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint32; F1 int; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 int64; F1 string; F2 uint32 }
type GenS9 struct { F0 string; F1 uint32; F2 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 bool; F1 int16; F2 uint }
func genFn0(p0 uint16, p1 bool, p2 int16) uint32 {
    v7 := []int{}
    v7 = append(v7, 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 bool, p1 int8, p2 bool) uint64 {
    v8 := []int{-245, 4294967295, -199}
    fmt.Println(p1, p0, p2)
    for fi9 := 0; fi9 < 8; fi9++ {
        _ = len(v8)
        var v10 float64 = 0.0
        p1 = (p1 | 100)
        _ = fi9
        _ = v10
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 42
    var v3 uint16 = 21
    v4 := []int{42, -2147483648, -576, 128}
    v5 := make(map[string]int)
    v6 := []string{}
    delete(v5, "foo bar")
    v2 = 42
    if (545 == 663) {
        var v11 float32 = (1.5 - 1.5)
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := v1[2147483647]
                var v15 float32 = v11
                if true {
                    var v16 float64 = float64(-2.25)
                    _ = len(v6)
                    _ = v16
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
        _ = v11
    } else {
        if (42 < 2) {
            v3 += 913
        } else {
            var v17 bool = (genFn1((!((-2.25 >= (-2.25 - (float32((1e10 + float64(float32(0.1)))) * 1e10))) && (-128 == -102))), 100, true) == 18446744073709551615)
            _ = v17
        }
        var v18 float64 = (3.14159 + 1e10)
        v5["foo bar"] = 318
        _ = v18
    }
    v19 := GenS1{F0: (v3 >> ((18446744073709551615) & 63))}
    delete(v1, 456)
    fmt.Println(v2, v3)
    var v20 float32 = -2.25
    for fi21 := 0; fi21 < 17; fi21++ {
        _ = len(v1)
        _ = fi21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        _ = len(v4)
        _ = fi22
    }
    if true {
        var v23 uint64 = 0
        _ = v23
    } else {
        if false {
            if true {
                v24 := GenS1{F0: v3}
                _ = v24
            } else {
                if ((2147483647 >= 2147483647) || true) {
                    v3 += 89
                    delete(v1, 4294967295)
                    v25 := GenS1{F0: (v3 - (v3 % ((v3) | 1)))}
                    _ = v25
                } else {
                    switch "" {
                    case "\n\t":
                        for fi26 := 0; fi26 < 9; fi26++ {
                            v20 = v20
                            if false {
                                v3 += v3
                            }
                            var v27 bool = (0 >= -34)
                            _ = fi26
                            _ = v27
                        }
                    default:
                        if true {
                            for fi28 := 0; fi28 < 11; fi28++ {
                                var v29 bool = (false && false)
                                for fi30 := 0; fi30 < 7; fi30++ {
                                    if v29 {
                                        v31 := GenS2{F0: 128, F1: (-(v2))}
                                        if (-128 == 1) {
                                            v29 = (!true)
                                            v32 := GenS2{F0: 196, F1: -481}
                                            v3 = 65535
                                            _ = v32
                                        } else {
                                            switch (-(fi28)) {
                                            case 32768:
                                                if false {
                                                    fmt.Println(v3)
                                                    fmt.Println(v3, v2, v29)
                                                    v2 = v2
                                                    switch "a" {
                                                    case "":
                                                        var v33 uint16 = v3
                                                        _ = v33
                                                    default:
                                                        var v34 int64 = 1
                                                        _ = v34
                                                    }
                                                }
                                            case 2147483647:
                                                if len(v4) > 0 {
                                                    v4[0] = 4294967295
                                                }
                                            default:
                                                var v35 string = "\n\t"
                                                _ = v35
                                            }
                                        }
                                        _ = v31
                                    } else {
                                        var v36 float32 = 1.5
                                        _ = v36
                                    }
                                    v20 -= v20
                                    v2 += (v2 * (^(v2)))
                                    fmt.Println(v20)
                                    _ = fi30
                                }
                                _ = fi28
                                _ = v29
                            }
                            v3 = v3
                        }
                    }
                }
                var v37 float64 = 3.14159
                if ((65535 == (v3 + (v3 & v3))) || (42 >= 800)) {
                    if true {
                        v38 := GenS3{F0: 980, F1: 32767, F2: 198}
                        v39 := GenS4{F0: v3}
                        _ = v38
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 13; fi40++ {
                        for fi41 := 0; fi41 < 10; fi41++ {
                            var v42 uint = 65535
                            var v43 float32 = (v20 - (v20 + v20))
                            var v44 bool = false
                            _ = fi41
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        v2 += v2
                        for fi45 := 0; fi45 < 12; fi45++ {
                            if true {
                                if false {
                                    if (1 >= 416) {
                                        var v46 int = (fi45 << ((255) & 63))
                                        _ = v46
                                    } else {
                                        if (1 < 255) {
                                            var v47 bool = false
                                            _ = v47
                                        }
                                        fmt.Println(v3, v3, v20)
                                    }
                                    v48 := GenS1{F0: 1}
                                    _ = v48
                                }
                            }
                            v6 = append(v6, ((("hello" + ("\n\t" + "")) + ("a" + ("foo bar" + "\n\t"))) + ""))
                            v20 = v20
                            _ = fi45
                        }
                        _ = fi40
                    }
                }
                _ = v37
            }
            v20 = v20
            var v49 uint32 = 832
            _ = v49
        }
        v50 := GenS5{F0: -100}
        var v51 uint = 4294967295
        _ = v50
        _ = v51
    }
    v6 = append(v6, "\n\t")
    if (v3 < v3) {
        if false {
            switch "foo bar" {
            case "foo bar":
                for fi52 := 0; fi52 < 14; fi52++ {
                    _ = len(v1)
                    switch "\n\t" {
                    case "a":
                        switch "foo bar" {
                        case "foo bar":
                            delete(v5, "foo bar")
                        case "a":
                            if (255 == 2) {
                                v3 = (^(v3))
                            } else {
                                v53 := GenS6{F0: 42}
                                for fi54 := 0; fi54 < 2; fi54++ {
                                    switch 4294967295 {
                                    case -796:
                                        var v55 float32 = ((0.0 / (v20 / ((3.14159 - ((float32(-2.25) * v20) * v20)) - 0.0))) - float32(-2.25))
                                        _ = v55
                                    default:
                                        switch ("a" + ("hello" + "")) {
                                        case "foo bar":
                                            var v56 bool = (!true)
                                            _ = v56
                                        case "hello":
                                            var v57 int16 = v2
                                            _ = v57
                                        case "\n\t":
                                            v2 = (^(v2))
                                        }
                                    }
                                    if false {
                                        if true {
                                            var v58 uint = 0
                                            _ = v58
                                        } else {
                                            v59 := GenS4{F0: 255}
                                            v60 := GenS7{F0: 0, F1: v2}
                                            v20 = v20
                                            _ = v59
                                            _ = v60
                                        }
                                        var v61 int32 = 650
                                        _ = v61
                                    } else {
                                        v62 := GenS2{F0: 128, F1: (v2 % (((^(v2))) | 1))}
                                        _ = v62
                                    }
                                    v3 += (^(v3))
                                    _ = fi54
                                }
                                if len(v4) > 0 {
                                    v4[0] = fi52
                                }
                                _ = v53
                            }
                        case "":
                            var v63 int8 = -2
                            _ = v63
                        default:
                            for fi64 := 0; fi64 < 14; fi64++ {
                                switch fi64 {
                                case 42:
                                    if true {
                                        switch "hello" {
                                        case "hello":
                                            v65 := GenS8{F0: int64(v3), F1: "a", F2: 65535}
                                            _ = v65
                                        case "\n\t":
                                            v66 := GenS6{F0: 852}
                                            _ = v66
                                        case "":
                                            if (true && (-32768 < (-(v2)))) {
                                                var v67 int32 = 42
                                                _ = v67
                                            }
                                        }
                                    } else {
                                        var v68 int8 = 127
                                        var v69 uint8 = 2
                                        _ = v68
                                        _ = v69
                                    }
                                    if false {
                                        var v70 bool = false
                                        var v71 int = fi64
                                        _ = v70
                                        _ = v71
                                    }
                                    v20 = v20
                                case 1:
                                    var v72 uint16 = v3
                                    _ = v72
                                case 541:
                                    switch "hello" {
                                    case "":
                                        v6 = append(v6, ("a" + ("\n\t" + (("foo bar" + ("a" + "a")) + ""))))
                                    case "hello":
                                        var v73 int16 = (v2 * v2)
                                        _ = v73
                                    }
                                default:
                                    var v74 uint64 = 155
                                    _ = v74
                                }
                                fmt.Println(fi64, fi52, v20)
                                switch "a" {
                                case "foo bar":
                                    if true {
                                        var v75 string = ("a" + ("hello" + ""))
                                        v3 = v3
                                        _ = v75
                                    } else {
                                        var v76 int64 = -1
                                        var v77 int64 = -1
                                        v78 := GenS9{F0: "\n\t", F1: 408, F2: 42}
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    }
                                default:
                                    switch fi52 {
                                    case 4294967295:
                                        fmt.Println(fi52)
                                    case 625:
                                        v79 := GenS1{F0: v3}
                                        _ = v79
                                    case 1:
                                        v2 = int16(fi52)
                                        var v80 int8 = -128
                                        _ = v80
                                    }
                                }
                                _ = fi64
                            }
                        }
                    }
                    var v81 uint32 = 477
                    _ = fi52
                    _ = v81
                }
            case "hello":
                v20 -= 1e10
            case "\n\t":
                if false {
                    if true {
                        v82 := GenS9{F0: "hello", F1: 0, F2: 735}
                        _ = v82
                    }
                    var v83 float32 = v20
                    _ = v83
                }
            default:
                if true {
                    v84 := GenS3{F0: 4294967295, F1: 128, F2: 42}
                    v1[940] = "hello"
                    fmt.Println(v3)
                    _ = v84
                } else {
                    v1[2147483647] = "\n\t"
                }
            }
        } else {
            var v85 string = ("a" + (("a" + "\n\t") + "foo bar"))
            _ = v85
        }
    } else {
        for fi86 := 0; fi86 < 17; fi86++ {
            v87 := GenS10{F0: (!true)}
            _ = fi86
            _ = v87
        }
        v3 -= v3
    }
    v20 = float32(1.5)
    for fi88 := 0; fi88 < 18; fi88++ {
        var v89 int8 = 127
        v90 := GenS2{F0: 102, F1: (v2 << ((4294967295) & 63))}
        _ = fi88
        _ = v89
        _ = v90
    }
    if (0.0 == (v20 + (v20 * -2.25))) {
        switch -532 {
        case 9223372036854775807:
            if (!(false || true)) {
                v91 := GenS10{F0: (float32(1e10) < v20)}
                v20 += v20
                _ = v91
            } else {
                v20 = 3.14159
            }
        case -9223372036854775808:
            v4 = append(v4, int(v2))
        case 32767:
            fmt.Println(v3)
        default:
            v92 := GenS11{F0: false, F1: (v2 % ((v2) | 1)), F2: 868}
            _ = v92
        }
        var v93 uint64 = 1
        v93 -= (v93 ^ v93)
        _ = v93
    }
    _ = len(v6)
    var v94 bool = (-861 != -1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v94
}
