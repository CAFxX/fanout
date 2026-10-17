package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int32; F1 int8; F2 uint16 }
type GenS5 struct { F0 bool; F1 uint; F2 string }

func main() {
    var v1 int8 = 100
    var v2 uint64 = 0
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    _ = len(v4)
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{32767, (fi5 & -309)}
        _ = fi5
        _ = v6
    }
    v2 -= v2
    if true {
        if false {
            v1 = v1
            for fi7 := 0; fi7 < 10; fi7++ {
                v8 := GenS1{F0: ("a" + "\n\t"), F1: 2, F2: 42}
                for fi9 := 0; fi9 < 6; fi9++ {
                    v10 := GenS1{F0: "", F1: 42, F2: (v1 - (v1 % ((42) | 1)))}
                    fmt.Println(v2, v2)
                    v1 -= (v1 | (v1 & (v1 | v1)))
                    _ = fi9
                    _ = v10
                }
                v1 = v1
                _ = fi7
                _ = v8
            }
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := []int{(fi11 ^ (fi11 / ((fi11) | 1)))}
                var v13 float32 = 0.0
                if true {
                    v14 := GenS2{F0: 42}
                    _ = v14
                } else {
                    var v15 float64 = (3.14159 + -2.25)
                    var v16 float64 = (v15 / v15)
                    fmt.Println(fi11, v2)
                    _ = v15
                    _ = v16
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        v17 := []int{0, 52}
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = v17
    }
    fmt.Println(v2, v2)
    if false {
        var v18 int32 = 2147483647
        v3[4294967295] = true
        _ = v18
    }
    switch "hello" {
    case "foo bar":
        var v19 uint8 = 239
        _ = v19
    case "a":
        var v20 float64 = 0.1
        _ = v20
    default:
        _ = len(v3)
    }
    fmt.Println(v2)
    v21 := []int{2147483647}
    switch "hello" {
    case "\n\t":
        for fi22 := 0; fi22 < 16; fi22++ {
            fmt.Println(fi22)
            fmt.Println(fi22)
            switch (fi22 << ((368) & 63)) {
            case -9223372036854775808:
                if (false || (255 >= 4294967295)) {
                    var v23 bool = ((0 < 42) && (!true))
                    switch "a" {
                    case "":
                        var v24 uint16 = 1
                        _ = v24
                    default:
                        var v25 uint8 = 128
                        _ = v25
                    }
                    _ = v23
                } else {
                    var v26 int64 = 42
                    v27 := GenS3{F0: 42, F1: 32767}
                    _ = v26
                    _ = v27
                }
            case 127:
                switch "\n\t" {
                case "a":
                    var v28 uint = 916
                    _ = v28
                case "hello":
                    if true {
                        if (65535 != 255) {
                            var v29 int32 = -1
                            _ = v29
                        }
                    } else {
                        fmt.Println(fi22, v1)
                        v1 -= 52
                        if true {
                            v1 -= (v1 + 127)
                            var v30 string = ""
                            v2 = (v2 << ((22) & 63))
                            _ = v30
                        } else {
                            for fi31 := 0; fi31 < 11; fi31++ {
                                v32 := v4["\n\t"]
                                v33 := GenS3{F0: 32767, F1: 0}
                                var v34 int8 = (v1 / ((v1) | 1))
                                _ = len(v4)
                                if (!true) {
                                    if (!(42 <= 32767)) {
                                        v32 = (fi31 * fi31)
                                    } else {
                                        var v35 int = 32767
                                        v2 += 18446744073709551615
                                        _ = v35
                                    }
                                    switch "a" {
                                    case "hello":
                                        var v36 int8 = -25
                                        _ = v36
                                    case "\n\t":
                                        v37 := v3[fi31]
                                        _ = v37
                                    case "":
                                        fmt.Println(v32)
                                    default:
                                        v34 -= v34
                                    }
                                    v1 -= (v34 % ((v34) | 1))
                                }
                                _ = fi31
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                        }
                    }
                case "":
                    var v38 int = (fi22 * fi22)
                    _ = v38
                default:
                    v39 := GenS4{F0: 782, F1: v1, F2: 615}
                    _ = v39
                }
            case 1:
                v1 += v1
            }
            _ = fi22
        }
    case "foo bar":
        switch "\n\t" {
        case "hello":
            if (1e10 < 1e10) {
                for fi40 := 0; fi40 < 2; fi40++ {
                    for fi41 := 0; fi41 < 12; fi41++ {
                        if false {
                            v21 = append(v21, fi40)
                            var v42 uint8 = 2
                            _ = v42
                        }
                        _ = fi41
                    }
                    _ = fi40
                }
            }
        default:
            if true {
                if len(v21) > 0 {
                    v21[0] = 32767
                }
            }
        }
    case "hello":
        switch 32767 {
        case 32767:
            v2 = (^(v2))
        case 32768:
            var v43 int = 42
            _ = v43
        case 9223372036854775807:
            if true {
                var v44 int = -287
                _ = v44
            } else {
                v2 = v2
            }
        }
    }
    if false {
        v2 = (^(v2))
    } else {
        if (!false) {
            v2 -= (^(v2))
        }
        v21 = append(v21, 821)
        v45 := GenS5{F0: ((v2 - (^(v2))) < 18446744073709551615), F1: 857, F2: ""}
        _ = v45
    }
    v46 := GenS5{F0: ("hello" < (("foo bar" + ("" + (("" + ((("a" + (("foo bar" + "hello") + ("foo bar" + "foo bar"))) + "a") + "\n\t")) + ("a" + ("foo bar" + ""))))) + (("hello" + "a") + "a"))), F1: 4294967295, F2: "\n\t"}
    v1 -= v1
    v1 = v1
    var v47 int8 = v1
    v1 += (v1 * v47)
    _ = len(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v46
    _ = v47
}
