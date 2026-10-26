package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int64; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int; F1 int32; F2 uint16 }
type GenS6 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint32, p1 int32) int16 {
    var v7 string = ("\n\t" + "a")
    for fi8 := 0; fi8 < 20; fi8++ {
        switch v7 {
        case "foo bar":
            v9 := GenS1{F0: v7, F1: 18446744073709551615}
            _ = v9
        case "\n\t":
            var v10 float64 = (1e10 / 1.0)
            _ = v10
        case "":
            if false {
                var v11 float32 = ((0.1 + 3.14159) + (3.14159 + (-2.25 + 0.0)))
                _ = v11
            }
        default:
            p1 -= 2147483647
        }
        v7 = v7
        v12 := []int{fi8, 42, 127, 32767}
        _ = fi8
        _ = v12
    }
    switch "a" {
    case "hello":
        v13 := []int{42, 42}
        _ = v13
    case "":
        if true {
            v7 = "a"
            var v14 uint16 = 3
            _ = v14
        }
    case "a":
        var v15 int8 = -1
        _ = v15
    default:
        p0 = p0
    }
    var v16 uint = 4294967295
    fmt.Println(p1, v7)
    _ = p0
    _ = p1
    _ = v16
    _ = v7
    return 708
}
func genFn1(p0 uint) uint8 {
    switch ("hello" + "a") {
    case "hello":
        var v17 uint16 = 32767
        _ = v17
    case "a":
        p0 = p0
    case "\n\t":
        for fi18 := 0; fi18 < 19; fi18++ {
            for fi19 := 0; fi19 < 16; fi19++ {
                fmt.Println(fi19)
                if (false && (1 == 747)) {
                    p0 -= 0
                    var v20 string = ""
                    fmt.Println(v20, v20, v20)
                    _ = v20
                }
                var v21 uint16 = 32767
                _ = fi19
                _ = v21
            }
            p0 = (p0 * p0)
            var v22 uint = (p0 - (p0 * (p0 * (p0 & p0))))
            _ = fi18
            _ = v22
        }
    }
    fmt.Println(p0)
    var v23 string = "a"
    var v24 bool = (!true)
    _ = p0
    _ = v23
    _ = v24
    return 216
}

func main() {
    var v1 float32 = (3.14159 * float32((1.5 / 3.14159)))
    var v2 int32 = -1
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    var v5 int16 = 1000
    v6 := []string{"a", "hello", ("\n\t" + ((("a" + ("\n\t" + "a")) + "foo bar") + (("" + "hello") + ("a" + ("" + "hello"))))), (("hello" + "\n\t") + "\n\t")}
    switch "" {
    case "a":
        v5 -= (v5 % (((v5 ^ (v5 >> ((2) & 63)))) | 1))
    case "foo bar":
        if ((("hello" + "foo bar") < "hello") && true) {
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 uint = 255
                _ = fi25
                _ = v26
            }
            v5 -= v5
            v5 -= (v5 * (v5 | v5))
        } else {
            v5 = v5
            v6 = append(v6, "\n\t")
            v27 := GenS1{F0: "foo bar", F1: 764}
            _ = v27
        }
    default:
        v1 -= (1.5 / v1)
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        delete(v4, 473)
    case "":
        switch "" {
        case "a":
            var v28 int8 = -51
            _ = v28
        case "\n\t":
            v2 = (v2 << ((2) & 63))
        }
    case "hello":
        var v29 int16 = (v5 | genFn0(221, (v2 >> ((18446744073709551615) & 63))))
        _ = v29
    }
    switch 541 {
    case -559:
        var v30 uint8 = genFn1(241)
        _ = v30
    case 285:
        if true {
            v31 := GenS1{F0: "a", F1: 773}
            switch 550 {
            case -376:
                if (42 == 1) {
                    v6 = append(v6, "foo bar")
                } else {
                    if (1e10 < float64((v1 * 0.0))) {
                        v6 = append(v6, ("a" + "a"))
                        var v32 int = 132
                        if (255 < 133) {
                            v32 -= -2147483648
                            for fi33 := 0; fi33 < 10; fi33++ {
                                v32 += -9223372036854775808
                                for fi34 := 0; fi34 < 9; fi34++ {
                                    switch "\n\t" {
                                    case "hello":
                                        var v35 uint8 = 2
                                        v36 := GenS2{F0: 2, F1: 42}
                                        var v37 uint64 = 2
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    case "":
                                        var v38 string = ""
                                        _ = v38
                                    case "\n\t":
                                        fmt.Println(fi34)
                                    }
                                    _ = fi34
                                }
                                v2 = 0
                                _ = fi33
                            }
                        }
                        _ = v32
                    } else {
                        for fi39 := 0; fi39 < 2; fi39++ {
                            v40 := GenS3{F0: 454, F1: v5, F2: genFn1(739)}
                            _ = fi39
                            _ = v40
                        }
                        v41 := GenS2{F0: 42, F1: 2}
                        _ = v41
                    }
                    if (9223372036854775807 <= 9223372036854775807) {
                        var v42 int32 = -1
                        v43 := GenS4{F0: 42, F1: (737 >= v2)}
                        if (v5 >= (v5 & (v5 >> ((42) & 63)))) {
                            var v44 uint8 = 1
                            v1 = v1
                            v4[9223372036854775807] = "hello"
                            _ = v44
                        }
                        _ = v42
                        _ = v43
                    }
                }
            case -1:
                for fi45 := 0; fi45 < 1; fi45++ {
                    for fi46 := 0; fi46 < 1; fi46++ {
                        v5 = v5
                        v47 := v3[(fi46 ^ (fi46 | fi46))]
                        _ = fi46
                        _ = v47
                    }
                    switch "hello" {
                    case "hello":
                        v48 := GenS5{F0: fi45, F1: v2, F2: 65535}
                        _ = v48
                    case "\n\t":
                        v5 += (v5 * v5)
                    }
                    v5 += v5
                    _ = fi45
                }
            case 31:
                for fi49 := 0; fi49 < 6; fi49++ {
                    if ((9223372036854775807 <= 42) || (!false)) {
                        v5 += v5
                        fmt.Println(fi49, v2)
                    } else {
                        fmt.Println(v2, v1)
                    }
                    for fi50 := 0; fi50 < 10; fi50++ {
                        v6 = append(v6, "\n\t")
                        var v51 uint = 255
                        _ = fi50
                        _ = v51
                    }
                    _ = fi49
                }
            default:
                var v52 uint8 = 89
                _ = v52
            }
            _ = v31
        } else {
            for fi53 := 0; fi53 < 14; fi53++ {
                fmt.Println(v1, v2)
                v2 = 73
                _ = fi53
            }
        }
    default:
        v1 = (3.14159 * v1)
    }
    v2 -= v2
    switch "hello" {
    case "hello":
        if true {
            switch (("" + "foo bar") + ("a" + ("" + "foo bar"))) {
            case "\n\t":
                var v54 uint32 = 2
                _ = v54
            case "hello":
                v5 += v5
            case "":
                fmt.Println(v5, v1)
            }
            v5 = (v5 ^ int16(v2))
        }
    default:
        v55 := GenS4{F0: 834, F1: (!true)}
        _ = v55
    }
    if false {
        v1 -= (v1 / 1.5)
        fmt.Println(v5)
    }
    v56 := v3[32768]
    var v57 uint = 255
    var v58 bool = true
    v59 := GenS5{F0: -406, F1: (^(v2)), F2: 219}
    for fi60 := 0; fi60 < 19; fi60++ {
        var v61 int = (fi60 << (((v57 & v57)) & 63))
        if ((v61 + 32768) <= (^(v61))) {
            if (v58 && v58) {
                var v62 int8 = 42
                v63 := GenS5{F0: v61, F1: v2, F2: 2}
                _ = v62
                _ = v63
            }
        } else {
            if (1 < 9223372036854775807) {
                for fi64 := 0; fi64 < 17; fi64++ {
                    var v65 int16 = (-(v5))
                    v66 := v3[(fi64 ^ fi60)]
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                _ = len(v6)
                v67 := GenS6{F0: (^(v2)), F1: 1}
                _ = v67
            } else {
                if (1 < 397) {
                    if (v56 && ((!v56) && (80 == 100))) {
                        v58 = (-1 != 0)
                    } else {
                        v58 = v56
                    }
                    var v68 bool = (820 != 4294967295)
                    v56 = (int16(v2) < genFn0(2147483648, v2))
                    _ = v68
                } else {
                    v1 = ((float32(3.14159) * float32(1.5)) + 1.5)
                }
            }
            fmt.Println(v61, v5, v1)
            var v69 int16 = v5
            _ = v69
        }
        _ = fi60
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
}
