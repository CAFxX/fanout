package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
type GenS2 struct { F0 string; F1 uint32; F2 uint16 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string; F1 int8; F2 int }
func genFn0(p0 string, p1 int16, p2 uint64) int {
    if (true || false) {
        p0 = (p0 + "foo bar")
        for fi4 := 0; fi4 < 12; fi4++ {
            v5 := []int{(fi4 << ((65535) & 63)), (-(fi4)), (fi4 | (fi4 << ((65535) & 63))), fi4}
            _ = fi4
            _ = v5
        }
    } else {
        p2 += (^(p2))
    }
    v6 := []int{-9223372036854775808, 127, 9223372036854775807}
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            if (1 == -72) {
                p0 = "hello"
                var v8 string = p0
                _ = v8
            } else {
                p0 = p0
                if (!true) {
                    p0 = p0
                    p0 = p0
                    _ = len(v6)
                }
                p1 -= 336
            }
        }
        var v9 string = (p0 + (p0 + p0))
        _ = len(v6)
        _ = fi7
        _ = v9
    }
    v6 = append(v6, 128)
    for fi10 := 0; fi10 < 7; fi10++ {
        switch (p0 + "hello") {
        case "\n\t":
            p1 += (p1 | (p1 * (p1 / (((p1 % ((1000) | 1))) | 1))))
        }
        var v11 int8 = 127
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 128
}
func genFn1(p0 int64, p1 bool) int64 {
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS1{F0: 2147483647, F1: -745, F2: 0}
        for fi14 := 0; fi14 < 4; fi14++ {
            fmt.Println(fi12, fi14, p0)
            var v15 int64 = (^(p0))
            fmt.Println(p1, p0, v15)
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            p1 = (!(true && ("\n\t" <= "hello")))
            var v17 int16 = 42
            _ = fi16
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    v18 := GenS1{F0: 42, F1: 42, F2: 1}
    if p1 {
        fmt.Println(p0, p1)
        var v19 int32 = -1
        _ = v19
    }
    v20 := []int{1}
    p0 = (p0 << ((42) & 63))
    _ = p0
    _ = p1
    _ = v18
    _ = v20
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 2147483647
    var v3 uint8 = 216
    v21 := []int{32767}
    for fi22 := 0; fi22 < 14; fi22++ {
        if (0.0 > -2.25) {
            _ = len(v21)
            switch "foo bar" {
            case "\n\t":
                for fi23 := 0; fi23 < 14; fi23++ {
                    var v24 uint = 42
                    _ = fi23
                    _ = v24
                }
            case "":
                for fi25 := 0; fi25 < 16; fi25++ {
                    v3 -= v3
                    if ((("a" + "a") + "foo bar") < "") {
                        var v26 uint16 = 934
                        var v27 uint = 42
                        _ = v26
                        _ = v27
                    } else {
                        v3 = v3
                        v2 -= (v2 << ((517) & 63))
                        if (0 <= 451) {
                            _ = len(v1)
                            v3 += v3
                            if (2 >= 255) {
                                fmt.Println(fi22, v3)
                                fmt.Println(fi25, fi22)
                                v28 := GenS2{F0: "foo bar", F1: 2147483647, F2: 2}
                                _ = v28
                            } else {
                                fmt.Println(v2, v2)
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                v2 = v2
                            }
                        }
                    }
                    v3 += (v3 | v3)
                    _ = fi25
                }
            case "foo bar":
                v29 := GenS2{F0: ("foo bar" + ""), F1: uint32(v3), F2: 0}
                _ = v29
            default:
                switch "a" {
                case "a":
                    for fi30 := 0; fi30 < 20; fi30++ {
                        v2 = (v2 >> ((18446744073709551615) & 63))
                        var v31 string = ("hello" + "")
                        _ = fi30
                        _ = v31
                    }
                }
            }
        }
        _ = fi22
    }
    if false {
        switch "\n\t" {
        case "\n\t":
            v2 = 0
        case "hello":
            v32 := GenS1{F0: v2, F1: 1, F2: 18446744073709551615}
            _ = v32
        case "":
            switch ("" + ("\n\t" + "")) {
            case "hello":
                v33 := GenS1{F0: (v2 + int32(v3)), F1: 0, F2: 0}
                _ = v33
            case "":
                var v34 uint = 1
                _ = v34
            default:
                v35 := GenS3{F0: ("a" + ""), F1: 9223372036854775807}
                _ = v35
            }
        }
        v3 = v3
        for fi36 := 0; fi36 < 3; fi36++ {
            delete(v1, 4294967295)
            _ = fi36
        }
    }
    if ((float32((1e10 / (((1e10 * -2.25) + 0.0) / 1e10))) + 0.0) < 1.5) {
        if true {
            v2 = -602
            v37 := GenS4{F0: 114, F1: 1, F2: 42}
            var v38 int32 = v2
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 2; fi39++ {
            for fi40 := 0; fi40 < 4; fi40++ {
                if true {
                    switch fi39 {
                    case 914:
                        var v41 int32 = -2147483648
                        _ = v41
                    case -9223372036854775808:
                        var v42 int32 = v2
                        _ = v42
                    case -2147483648:
                        switch 4294967295 {
                        case -795:
                            _ = len(v1)
                        case -9223372036854775808:
                            switch "" {
                            case "hello":
                                v43 := GenS5{F0: 2, F1: (v2 | (v2 / (((v2 ^ (v2 + 42))) | 1)))}
                                _ = v43
                            case "a":
                                switch fi39 {
                                case 42:
                                    v2 = 887
                                default:
                                    v21 = append(v21, (fi40 | (^(fi39))))
                                }
                            case "foo bar":
                                if false {
                                    switch ("a" + (("hello" + "hello") + "\n\t")) {
                                    case "a":
                                        v21 = append(v21, (-(fi39)))
                                    case "foo bar":
                                        var v44 int8 = -1
                                        _ = v44
                                    case "":
                                        v2 = 42
                                    default:
                                        v2 = -2147483648
                                    }
                                    v45 := GenS6{F0: 2147483647}
                                    var v46 float64 = (3.14159 / 1.5)
                                    _ = v45
                                    _ = v46
                                } else {
                                    var v47 uint8 = (v3 << ((18446744073709551615) & 63))
                                    switch "\n\t" {
                                    case "foo bar":
                                        switch genFn0(("a" + "foo bar"), 90, 944) {
                                        case 196:
                                            var v48 uint8 = (v47 / (((v3 + (v47 - (v3 + (v47 - 0))))) | 1))
                                            _ = v48
                                        case -2147483648:
                                            v49 := GenS1{F0: (v2 % (((^(v2))) | 1)), F1: -1, F2: 34}
                                            _ = v49
                                        case 950:
                                            v47 = (v3 ^ v47)
                                            var v50 int8 = -128
                                            v50 = 0
                                            fmt.Println(v2, fi40, v3)
                                            var v51 uint32 = 1
                                            switch fi39 {
                                            case 128:
                                                fmt.Println(v47, v3)
                                            default:
                                                switch "a" {
                                                case "":
                                                    if false {
                                                        if true {
                                                            v52 := v1[(v51 + v51)]
                                                            if true {
                                                                v3 = v3
                                                            } else {
                                                                v2 = v2
                                                            }
                                                            fmt.Println(v2, v2)
                                                            _ = v52
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v50
                                            _ = v51
                                        }
                                    case "a":
                                        var v53 int8 = int8(v47)
                                        _ = v53
                                    }
                                    _ = v47
                                }
                            }
                        case -1:
                            var v54 float32 = 0.1
                            _ = v54
                        default:
                            v3 += v3
                            var v55 int64 = 9223372036854775807
                            v56 := GenS1{F0: v2, F1: 0, F2: 1}
                            _ = v55
                            _ = v56
                        }
                    }
                }
                v2 = v2
                delete(v1, 1)
                if true {
                    v3 += v3
                }
                _ = fi40
            }
            _ = fi39
        }
        if true {
            var v57 int64 = 0
            fmt.Println(v2)
            v57 -= (v57 | genFn1((v57 - v57), (!false)))
            _ = v57
        }
    } else {
        for fi58 := 0; fi58 < 15; fi58++ {
            v2 += (v2 / (((v2 / ((42) | 1))) | 1))
            _ = fi58
        }
        v59 := GenS7{F0: "\n\t", F1: 70, F2: -924}
        for fi60 := 0; fi60 < 2; fi60++ {
            var v61 float64 = float64(float32(-2.25))
            var v62 uint64 = 2
            _ = fi60
            _ = v61
            _ = v62
        }
        _ = v59
    }
    var v63 string = "a"
    for fi64 := 0; fi64 < 1; fi64++ {
        if false {
            var v65 uint16 = 2
            for fi66 := 0; fi66 < 19; fi66++ {
                var v67 bool = (((-(v2)) >= -1) || (42 == (v3 - (v3 + v3))))
                _ = fi66
                _ = v67
            }
            _ = v65
        }
        var v68 uint16 = uint16(v3)
        _ = fi64
        _ = v68
    }
    var v69 int = -741
    v2 = v2
    v69 = v69
    switch v69 {
    case 32767:
        v63 = (("foo bar" + (((((v63 + "a") + "foo bar") + "\n\t") + "\n\t") + "")) + (v63 + (v63 + ("a" + (("\n\t" + (v63 + (v63 + "a"))) + "hello")))))
    }
    var v70 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v63
    _ = v69
    _ = v70
}
