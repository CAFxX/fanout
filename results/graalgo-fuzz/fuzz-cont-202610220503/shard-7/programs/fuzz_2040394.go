package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint8; F1 int64; F2 string }
type GenS5 struct { F0 int16; F1 uint }
func genFn0(p0 string, p1 uint, p2 int32) int {
    v6 := GenS1{F0: p0}
    fmt.Println(p2, p1, p2)
    var v7 int8 = -30
    v8 := GenS2{F0: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 9223372036854775807
}
func genFn1(p0 int, p1 int16, p2 uint) uint64 {
    if false {
        v9 := GenS3{F0: 886, F1: 465}
        switch "hello" {
        case "foo bar":
            v10 := GenS1{F0: "a"}
            _ = v10
        case "":
            var v11 int64 = 1
            _ = v11
        }
        _ = v9
    } else {
        var v12 uint16 = 42
        for fi13 := 0; fi13 < 4; fi13++ {
            if false {
                p1 -= 1
                var v14 uint64 = 18446744073709551615
                _ = v14
            }
            p0 -= fi13
            _ = fi13
        }
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 string = ("a" + ("hello" + "foo bar"))
            var v17 uint = (^(p2))
            switch (fi15 | p0) {
            case 32768:
                v18 := []int{(fi15 % (((p0 << (((p2 >> (((^(p2))) & 63))) & 63))) | 1)), fi15, fi15}
                _ = v18
            case -9223372036854775808:
                p0 -= int(p1)
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v12
    }
    if true {
        switch ("hello" + "\n\t") {
        case "a":
            if (1 == 0) {
                var v19 uint = (^(p2))
                v20 := GenS3{F0: 889, F1: 284}
                switch (p0 / (((^(p0))) | 1)) {
                case 42:
                    for fi21 := 0; fi21 < 1; fi21++ {
                        for fi22 := 0; fi22 < 17; fi22++ {
                            v23 := []int{fi21, (^(fi21)), (p0 | (^(fi21)))}
                            v24 := GenS1{F0: ("a" + "\n\t")}
                            var v25 uint16 = 0
                            _ = fi22
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 12; fi26++ {
                            var v27 int8 = int8(fi21)
                            p1 -= (p1 | 570)
                            _ = fi26
                            _ = v27
                        }
                        _ = fi21
                    }
                default:
                    if (!((903 >= 424) && true)) {
                        v28 := []int{(p0 & (p0 << ((p2) & 63))), (p0 % (((p0 & (p0 / ((genFn0(("\n\t" + "hello"), 65535, -92)) | 1)))) | 1)), genFn0("hello", v19, -1), 42}
                        _ = len(v28)
                        fmt.Println(p1, p2, p0)
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 13; fi29++ {
                            if ((4294967295 < 2147483647) || true) {
                                fmt.Println(v19)
                            }
                            fmt.Println(p0, fi29)
                            _ = fi29
                        }
                        v30 := GenS4{F0: 240, F1: -731, F2: "foo bar"}
                        switch 2147483647 {
                        case 42:
                            p2 = 1
                        case 433:
                            var v31 int32 = 577
                            _ = v31
                        case 799:
                            for fi32 := 0; fi32 < 7; fi32++ {
                                v33 := GenS4{F0: 214, F1: 9223372036854775807, F2: ("" + "hello")}
                                var v34 float64 = float64(1.5)
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                        }
                        _ = v30
                    }
                }
                _ = v19
                _ = v20
            }
        default:
            p0 = p0
        }
        p2 = (^(p2))
    } else {
        var v35 int32 = -2147483648
        p2 = 2
        p1 -= 53
        _ = v35
    }
    v36 := []int{p0, (p0 ^ genFn0("hello", (p2 + 2), -2147483648)), p0, 42}
    v37 := GenS2{F0: 621}
    switch ("foo bar" + (("foo bar" + "a") + "\n\t")) {
    case "\n\t":
        if len(v36) > 0 {
            v36[0] = p0
        }
    case "a":
        var v38 uint16 = 256
        _ = v38
    default:
        if len(v36) > 0 {
            v36[0] = p0
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    _ = v37
    return 706
}

func main() {
    v1 := []string{("foo bar" + ("\n\t" + "\n\t")), ("a" + ((("\n\t" + ((("hello" + "\n\t") + ("" + "\n\t")) + ("a" + "foo bar"))) + "hello") + "")), ""}
    v2 := []string{"\n\t", ""}
    v3 := []uint32{510, 4294967295, 2}
    var v4 int32 = 1
    v5 := []int{32768}
    for fi39 := 0; fi39 < 18; fi39++ {
        for fi40 := 0; fi40 < 13; fi40++ {
            var v41 uint64 = 4294967295
            _ = fi40
            _ = v41
        }
        var v42 uint64 = 0
        _ = fi39
        _ = v42
    }
    if (-32768 != -32768) {
        v4 = -2147483648
        v43 := GenS4{F0: 0, F1: 9223372036854775807, F2: (("foo bar" + ("a" + "\n\t")) + ((((("hello" + ("" + "")) + ("" + "a")) + "") + "a") + ""))}
        _ = len(v3)
        _ = v43
    } else {
        var v44 uint8 = 69
        switch "hello" {
        case "":
            var v45 int8 = 1
            _ = v45
        case "hello":
            var v46 int16 = -688
            _ = v46
        default:
            for fi47 := 0; fi47 < 13; fi47++ {
                var v48 uint32 = 0
                var v49 int64 = 1
                _ = fi47
                _ = v48
                _ = v49
            }
        }
        _ = v44
    }
    if (32767 == -215) {
        switch 1 {
        case 2147483647:
            if false {
                var v50 uint32 = 767
                _ = v50
            } else {
                var v51 uint = 648
                v52 := GenS1{F0: ""}
                switch genFn0(("\n\t" + ""), v51, 42) {
                case 42:
                    fmt.Println(v51)
                case 365:
                    switch "hello" {
                    case "\n\t":
                        if false {
                            if (((42 > -128) || false) || false) {
                                if true {
                                    var v53 float64 = 1.5
                                    v53 = v53
                                    _ = v53
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v51 = 2
                                }
                                for fi54 := 0; fi54 < 16; fi54++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi54
                                }
                                var v55 int16 = -32768
                                _ = v55
                            } else {
                                for fi56 := 0; fi56 < 11; fi56++ {
                                    v51 -= v51
                                    for fi57 := 0; fi57 < 14; fi57++ {
                                        switch "" {
                                        case "hello":
                                            var v58 uint64 = 1
                                            _ = v58
                                        }
                                        _ = fi57
                                    }
                                    v4 -= 0
                                    _ = fi56
                                }
                                var v59 float32 = float32((1.5 + 0.1))
                                _ = v59
                            }
                            v51 += (^(v51))
                            var v60 bool = ((1000 > -662) || (-819 < -953))
                            _ = v60
                        }
                    case "a":
                        for fi61 := 0; fi61 < 7; fi61++ {
                            var v62 int8 = 100
                            _ = fi61
                            _ = v62
                        }
                    }
                case -2147483648:
                    var v63 bool = ((genFn0("hello", v51, -275) <= 0) || (!false))
                    _ = v63
                default:
                    var v64 int = int(v4)
                    _ = v64
                }
                _ = v51
                _ = v52
            }
        }
        var v65 uint = 255
        _ = v65
    }
    for fi66 := 0; fi66 < 2; fi66++ {
        var v67 bool = (4294967295 >= (fi66 + fi66))
        var v68 bool = false
        switch "" {
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi66
        _ = v67
        _ = v68
    }
    v69 := GenS3{F0: 1, F1: (-(v4))}
    var v70 int = 4294967295
    if true {
        v4 += (^(v4))
        var v71 int8 = 42
        v70 -= genFn0((("foo bar" + (("\n\t" + ("\n\t" + "foo bar")) + "a")) + ((("hello" + ("foo bar" + ("\n\t" + (("" + "\n\t") + "\n\t")))) + ((("hello" + "") + "hello") + "hello")) + "foo bar")), 1, -563)
        _ = v71
    } else {
        v72 := GenS5{F0: -1, F1: 1}
        var v73 int64 = 0
        _ = v72
        _ = v73
    }
    for fi74 := 0; fi74 < 16; fi74++ {
        v70 = 315
        _ = fi74
    }
    v4 = v4
    v75 := GenS1{F0: "foo bar"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
    _ = v75
}
