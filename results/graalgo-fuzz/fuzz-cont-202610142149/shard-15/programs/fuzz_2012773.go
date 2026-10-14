package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int; F2 uint64 }
type GenS5 struct { F0 int16; F1 int16; F2 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8; F1 int16; F2 bool }
func genFn0(p0 int32, p1 uint16) uint32 {
    v7 := []int{}
    if (("a" + ("" + "a")) != "") {
        var v8 int64 = -9223372036854775808
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: -1, F2: 0}
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint16 = 880
        _ = fi10
        _ = v11
    }
    if (!true) {
        var v12 float64 = ((-2.25 * 0.1) + 0.0)
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 0
}
func genFn1(p0 uint16, p1 int) uint64 {
    p0 = 42
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS2{F0: -2147483648, F1: 828}
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 uint8 = 255
            v17 := GenS1{F0: (p0 | p0), F1: (fi15 - (p1 >> ((540) & 63))), F2: 518}
            _ = fi15
            _ = v16
            _ = v17
        }
        fmt.Println(p1, p1)
        _ = fi13
        _ = v14
    }
    var v18 int = p1
    _ = p0
    _ = p1
    _ = v18
    return 553
}

func main() {
    var v1 float32 = 3.14159
    v2 := []int{-1, 127}
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    var v5 bool = (65535 != 2147483647)
    var v6 int = 723
    fmt.Println(v6, v5, v6)
    switch "a" {
    case "hello":
        if (719 != 18446744073709551615) {
            if true {
                for fi19 := 0; fi19 < 13; fi19++ {
                    var v20 int64 = 9223372036854775807
                    switch fi19 {
                    case 32768:
                        if true {
                            if (!("\n\t" < ("hello" + ("a" + "\n\t")))) {
                                var v21 uint = 2
                                delete(v4, (("a" + ("foo bar" + (("hello" + ((("" + "a") + "a") + "\n\t")) + "\n\t"))) + ""))
                                _ = v21
                            } else {
                                switch ("a" + ("" + (("hello" + ("foo bar" + "a")) + "foo bar"))) {
                                case "":
                                    v20 = v20
                                case "hello":
                                    for fi22 := 0; fi22 < 1; fi22++ {
                                        v4["hello"] = ""
                                        fmt.Println(v1, fi22, fi22)
                                        if ((((1.5 - (3.14159 / float32(float64(v1)))) - v1) + (v1 * 0.0)) == float32((1e10 / 1.5))) {
                                            switch ("hello" + "foo bar") {
                                            case "":
                                                v23 := GenS3{F0: 1}
                                                _ = v23
                                            case "hello":
                                                var v24 uint8 = 42
                                                _ = v24
                                            }
                                            var v25 int = 128
                                            if v5 {
                                                var v26 bool = (false || v5)
                                                v25 += (fi19 % (((fi19 * (^(v6)))) | 1))
                                                v26 = (0 > (^(fi22)))
                                                _ = v26
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                var v27 float32 = v1
                                                _ = v27
                                            case "hello":
                                                v28 := GenS4{F0: -128, F1: v25, F2: 1}
                                                _ = v28
                                            default:
                                                v29 := GenS5{F0: 1000, F1: 0, F2: ""}
                                                _ = v29
                                            }
                                            _ = v25
                                        }
                                        _ = fi22
                                    }
                                case "foo bar":
                                    var v30 int16 = -170
                                    _ = v30
                                }
                                for fi31 := 0; fi31 < 20; fi31++ {
                                    var v32 int16 = 0
                                    _ = fi31
                                    _ = v32
                                }
                                v33 := GenS1{F0: 32767, F1: (v6 / ((v6) | 1)), F2: 571}
                                _ = v33
                            }
                            v34 := GenS4{F0: -81, F1: (fi19 ^ fi19), F2: 4294967295}
                            _ = v34
                        }
                    default:
                        v35 := GenS4{F0: -128, F1: (v6 - (-(fi19))), F2: 1}
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 4; fi36++ {
                        fmt.Println(v1, v5, v6)
                        _ = fi36
                    }
                    _ = fi19
                    _ = v20
                }
                v6 = v6
            } else {
                v1 = (v1 * (v1 * 1e10))
                v5 = v5
                for fi37 := 0; fi37 < 16; fi37++ {
                    if (v5 || (!v5)) {
                        v38 := GenS3{F0: 888}
                        v6 += (fi37 - fi37)
                        if true {
                            for fi39 := 0; fi39 < 5; fi39++ {
                                var v40 bool = v5
                                _ = fi39
                                _ = v40
                            }
                            var v41 float32 = v1
                            _ = v41
                        } else {
                            switch fi37 {
                            case -643:
                                fmt.Println(v6, fi37)
                            }
                            var v42 int64 = 42
                            var v43 uint = 213
                            _ = v42
                            _ = v43
                        }
                        _ = v38
                    } else {
                        v44 := GenS6{F0: v5}
                        _ = v44
                    }
                    v45 := GenS4{F0: -1, F1: v6, F2: 18446744073709551615}
                    v1 -= (v1 * 0.0)
                    _ = fi37
                    _ = v45
                }
            }
        }
    case "\n\t":
        v1 = (v1 + (v1 + v1))
    case "foo bar":
        v46 := GenS7{F0: 918}
        _ = v46
    default:
        switch (v6 * v6) {
        case 1:
            if true {
                v1 -= 0.1
            }
        case 42:
            var v47 uint64 = 42
            _ = v47
        case -264:
            var v48 uint64 = 717
            _ = v48
        }
    }
    var v49 int64 = 640
    var v50 float64 = float64(v1)
    _ = len(v4)
    switch ("foo bar" + "\n\t") {
    case "":
        var v51 int32 = 42
        _ = v51
    case "hello":
        var v52 int64 = v49
        _ = v52
    }
    if ("hello" > ("" + "foo bar")) {
        _ = len(v3)
        v50 -= v50
    }
    var v53 uint8 = 0
    if ((42 != 4294967295) || (v53 < 2)) {
        v53 = 5
        v54 := v3[(v6 / ((v6) | 1))]
        var v55 float64 = (1.5 - ((v50 / (v50 + ((((v50 * -2.25) + 1e10) / v50) / v50))) - float64(((((0.0 - 3.14159) * v1) + ((1.5 - 1.5) + -2.25)) + -2.25))))
        _ = v54
        _ = v55
    } else {
        switch "a" {
        case "a":
            fmt.Println(v50, v6, v53)
        case "hello":
            v56 := GenS5{F0: 1000, F1: -691, F2: "foo bar"}
            _ = v56
        case "\n\t":
            var v57 int8 = -48
            _ = v57
        default:
            switch "\n\t" {
            case "\n\t":
                switch v6 {
                case -853:
                    for fi58 := 0; fi58 < 3; fi58++ {
                        fmt.Println(v50, fi58, fi58)
                        for fi59 := 0; fi59 < 15; fi59++ {
                            v53 += 2
                            _ = fi59
                        }
                        v60 := GenS8{F0: 42, F1: 1000, F2: true}
                        _ = fi58
                        _ = v60
                    }
                case 0:
                    v61 := v4[(("" + ("\n\t" + ((("\n\t" + ("\n\t" + "")) + "") + "a"))) + ("\n\t" + "\n\t"))]
                    _ = v61
                case -2147483648:
                    fmt.Println(v6)
                }
            default:
                if v5 {
                    v62 := GenS2{F0: 733, F1: 1}
                    v1 = 0.1
                    if ((v5 || (100 <= -124)) && true) {
                        var v63 uint64 = 86
                        v49 = v49
                        _ = v63
                    } else {
                        for fi64 := 0; fi64 < 18; fi64++ {
                            v5 = v5
                            _ = fi64
                        }
                        v65 := GenS5{F0: 1000, F1: 516, F2: ("a" + "a")}
                        _ = v65
                    }
                    _ = v62
                }
            }
        }
        var v66 uint = 588
        fmt.Println(v53, v66, v5)
        _ = v66
    }
    if ((genFn1(2, 4294967295) < genFn1(572, (-(v6)))) && v5) {
        if len(v2) > 0 {
            v2[0] = (v6 - (v6 << ((2) & 63)))
        }
    } else {
        v53 = (v53 + 255)
    }
    fmt.Println(v53, v6, v5)
    var v67 bool = (v5 && true)
    v1 -= v1
    if true {
        switch 0 {
        case 1:
            v68 := GenS4{F0: 42, F1: 491, F2: 2}
            _ = v68
        case 42:
            v6 -= 32767
        default:
            if ((1 >= 42) && true) {
                var v69 bool = true
                _ = v69
            }
        }
        if len(v2) > 0 {
            v2[0] = (v6 / ((v6) | 1))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v53
    _ = v6
    _ = v67
}
