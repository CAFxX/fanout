package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 string; F2 int8 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int16; F1 string; F2 uint16 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := GenS1{F0: 1}
        var v7 int16 = 42
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint16 = 42
    v8 += 295
    var v9 uint = 42
    fmt.Println(v8, v9)
    _ = v8
    _ = v9
    return 1000
}
func genFn1() int {
    var v10 int16 = 581
    var v11 int8 = 100
    var v12 uint8 = 42
    if true {
        v12 += v12
    }
    if false {
        var v13 int8 = -1
        v14 := []int{-782}
        var v15 uint16 = 32767
        _ = v13
        _ = v14
        _ = v15
    }
    _ = v10
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 float64 = 1e10
    v2 := []string{"foo bar", (("foo bar" + ("" + "")) + ""), "", ("a" + "foo bar")}
    var v3 uint64 = 4294967295
    v4 := []int{-936, 128, 32768}
    v3 = 18446744073709551615
    v2 = append(v2, (("\n\t" + "") + "foo bar"))
    if len(v2) > 0 {
        v2[0] = ((("\n\t" + "\n\t") + "a") + "a")
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 string = ("a" + "foo bar")
        switch (fi16 - -739) {
        case -2147483648:
            var v18 uint16 = 65535
            _ = v18
        case 127:
            var v19 string = ""
            _ = v19
        case -138:
            var v20 int = genFn1()
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    var v21 uint32 = 2147483648
    if (true || (-1 < 1)) {
        for fi22 := 0; fi22 < 8; fi22++ {
            if len(v2) > 0 {
                v2[0] = ("" + "")
            }
            var v23 uint64 = 251
            _ = fi22
            _ = v23
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi24 := 0; fi24 < 7; fi24++ {
            v2 = append(v2, "\n\t")
            v1 += v1
            _ = fi24
        }
    }
    if (genFn1() != -1) {
        switch "foo bar" {
        case "hello":
            if (v3 != v3) {
                for fi25 := 0; fi25 < 17; fi25++ {
                    var v26 int8 = 103
                    if (float32(v1) <= 1.5) {
                        v27 := GenS2{F0: v26, F1: "foo bar", F2: (v26 * (v26 * v26))}
                        v26 -= v26
                        _ = v27
                    } else {
                        switch "" {
                        case "hello":
                            v28 := GenS3{F0: (fi25 | genFn1()), F1: 499}
                            _ = v28
                        default:
                            if (65535 > uint(v3)) {
                                v29 := GenS2{F0: (v26 / ((42) | 1)), F1: "\n\t", F2: (v26 | v26)}
                                if true {
                                    v26 -= -1
                                    var v30 uint64 = (v3 * v3)
                                    for fi31 := 0; fi31 < 5; fi31++ {
                                        if len(v2) > 0 {
                                            v2[0] = ("" + ("foo bar" + "\n\t"))
                                        }
                                        _ = fi31
                                    }
                                    _ = v30
                                }
                                _ = v29
                            }
                        }
                        var v32 float64 = float64((0.1 - -2.25))
                        _ = v32
                    }
                    _ = fi25
                    _ = v26
                }
                for fi33 := 0; fi33 < 7; fi33++ {
                    fmt.Println(v1, v3, v3)
                    for fi34 := 0; fi34 < 9; fi34++ {
                        v35 := GenS2{F0: -128, F1: "foo bar", F2: -113}
                        switch "hello" {
                        case "a":
                            var v36 int32 = 183
                            _ = v36
                        }
                        var v37 uint = 18446744073709551615
                        _ = fi34
                        _ = v35
                        _ = v37
                    }
                    _ = fi33
                }
                if (42 > 127) {
                    var v38 uint16 = 0
                    _ = v38
                }
            }
        case "a":
            for fi39 := 0; fi39 < 13; fi39++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v3 -= (v3 + 858)
                _ = fi39
            }
        case "foo bar":
            _ = len(v2)
        default:
            v1 = 0.1
        }
    }
    for fi40 := 0; fi40 < 6; fi40++ {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        _ = fi40
    }
    v3 = v3
    v1 = (v1 / (0.1 / 1.0))
    var v41 uint8 = 2
    var v42 int = 128
    if false {
        var v43 uint8 = (^(v41))
        if (!true) {
            v3 = v3
        }
        _ = v43
    } else {
        v44 := GenS3{F0: int(v41), F1: 1}
        v1 -= (v1 * -2.25)
        _ = v44
    }
    if ((9223372036854775807 != 42) && (!true)) {
        var v45 uint8 = (^(v41))
        if (float32(((v1 / (v1 + (0.1 * float64(1.5)))) * (v1 + 1.5))) == -2.25) {
            if false {
                switch "hello" {
                case "a":
                    var v46 uint64 = v3
                    _ = v46
                case "hello":
                    v45 -= (^(v41))
                }
                v47 := GenS2{F0: 42, F1: "\n\t", F2: -128}
                switch "\n\t" {
                case "":
                    fmt.Println(v45)
                case "foo bar":
                    v1 -= ((((v1 - ((((((v1 - v1) - v1) * 3.14159) / 1e10) * float64(3.14159)) / ((((1e10 - v1) / 1.0) / v1) - (float64((-2.25 * 1e10)) / v1)))) + (((v1 + v1) - (v1 / v1)) - (v1 * (3.14159 * v1)))) * v1) + v1)
                case "\n\t":
                    if (!true) {
                        for fi48 := 0; fi48 < 8; fi48++ {
                            var v49 int32 = 2147483647
                            if false {
                                switch 128 {
                                case 42:
                                    for fi50 := 0; fi50 < 20; fi50++ {
                                        var v51 float32 = float32(v1)
                                        _ = fi50
                                        _ = v51
                                    }
                                }
                            }
                            _ = fi48
                            _ = v49
                        }
                        if true {
                            v21 = v21
                        }
                    }
                }
                _ = v47
            }
            v52 := GenS1{F0: (v3 << ((4294967295) & 63))}
            _ = v52
        } else {
            v3 = v3
        }
        _ = v45
    }
    for fi53 := 0; fi53 < 8; fi53++ {
        v54 := GenS3{F0: int(v3), F1: 42}
        var v55 uint = 1
        var v56 string = "hello"
        _ = fi53
        _ = v54
        _ = v55
        _ = v56
    }
    fmt.Println(v42, v41)
    if (true && (255 < 65535)) {
        for fi57 := 0; fi57 < 15; fi57++ {
            var v58 uint8 = v41
            for fi59 := 0; fi59 < 15; fi59++ {
                v60 := GenS4{F0: int16(v58), F1: ("hello" + ("foo bar" + ("a" + ("foo bar" + ("a" + "foo bar"))))), F2: 42}
                _ = fi59
                _ = v60
            }
            fmt.Println(fi57, v42)
            _ = fi57
            _ = v58
        }
        fmt.Println(v41, v3, v1)
        _ = len(v2)
    }
    var v61 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v61
}
