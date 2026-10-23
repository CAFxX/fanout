package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int; F1 int; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 int64; F1 bool; F2 uint }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 uint16 = 1
    var v4 float64 = (float64(0.0) + 3.14159)
    var v5 bool = ("" > "\n\t")
    v6 := []int{-660, -2147483648}
    v4 -= 0.0
    fmt.Println(v5, v5)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint64 = 2
        switch "" {
        case "a":
            v3 += v3
        default:
            v8 += 4294967295
        }
        _ = fi7
        _ = v8
    }
    if len(v6) > 0 {
        v6[0] = 32768
    }
    fmt.Println(v5, v3, v5)
    if ((v4 > 0.0) || false) {
        if (!(42 == 1)) {
            if (v5 && false) {
                if ((false && v5) && true) {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        v10 := GenS1{F0: 127, F1: 1}
                        for fi11 := 0; fi11 < 19; fi11++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v3 = (v3 + 932)
                            v3 = v3
                            _ = fi11
                        }
                        for fi12 := 0; fi12 < 6; fi12++ {
                            v5 = v5
                            if len(v6) > 0 {
                                v6[0] = (fi9 + fi12)
                            }
                            _ = fi12
                        }
                        _ = fi9
                        _ = v10
                    }
                    var v13 int64 = -9223372036854775808
                    var v14 string = (((("" + ("hello" + "a")) + "hello") + "foo bar") + "hello")
                    _ = v13
                    _ = v14
                }
                v15 := GenS2{F0: 1, F1: 4294967295, F2: "a"}
                for fi16 := 0; fi16 < 8; fi16++ {
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v5 = ("\n\t" > "\n\t")
                        _ = fi17
                    }
                    _ = fi16
                }
                _ = v15
            } else {
                switch "a" {
                case "foo bar":
                    var v18 int64 = 9223372036854775807
                    _ = v18
                case "hello":
                    v19 := GenS2{F0: 128, F1: 2147483647, F2: ("foo bar" + "\n\t")}
                    _ = v19
                case "\n\t":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    var v20 uint32 = 432
                    _ = v20
                }
                v3 = v3
            }
        }
        switch 136 {
        case 555:
            delete(v1, 32767)
        case -2147483648:
            switch 42 {
            case 9223372036854775807:
                var v21 float64 = (v4 + float64(0.0))
                _ = v21
            case -432:
                for fi22 := 0; fi22 < 15; fi22++ {
                    var v23 uint8 = 2
                    var v24 bool = (-1 <= 0)
                    fmt.Println(v23)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
        default:
            var v25 int8 = 127
            _ = v25
        }
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 uint = 530
            if false {
                v28 := GenS3{F0: ""}
                for fi29 := 0; fi29 < 15; fi29++ {
                    v5 = (v3 < (^(v3)))
                    var v30 int32 = 0
                    fmt.Println(v4)
                    v2[fi26] = true
                    _ = fi29
                    _ = v30
                }
                _ = v28
            } else {
                if (((^(v27)) != (v27 + 604)) && v5) {
                    switch "" {
                    case "\n\t":
                        if (!(v5 || (0 == 219))) {
                            fmt.Println(v5, fi26)
                            for fi31 := 0; fi31 < 6; fi31++ {
                                v5 = v5
                                _ = fi31
                            }
                            v27 -= 1
                        } else {
                            v3 = v3
                        }
                    case "hello":
                        v5 = v5
                    default:
                        v3 -= (^(v3))
                    }
                    var v32 uint32 = 0
                    v33 := GenS4{F0: -492, F1: 734}
                    _ = v32
                    _ = v33
                } else {
                    fmt.Println(fi26)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    switch -9223372036854775808 {
                    case 4294967295:
                        v34 := GenS1{F0: 127, F1: 42}
                        _ = v34
                    case 128:
                        if v5 {
                            v3 -= 32767
                            for fi35 := 0; fi35 < 10; fi35++ {
                                switch ((("foo bar" + ("\n\t" + ("a" + "foo bar"))) + "a") + "hello") {
                                case "\n\t":
                                    switch "a" {
                                    case "a":
                                        _ = len(v6)
                                    }
                                case "a":
                                    if true {
                                        var v36 int8 = 11
                                        _ = v36
                                    } else {
                                        switch (fi35 / (((fi35 / (((fi35 / (((fi26 + 1)) | 1))) | 1))) | 1)) {
                                        case -2147483648:
                                            var v37 int32 = 1
                                            _ = v37
                                        default:
                                            v3 = (v3 - 42)
                                        }
                                        var v38 float32 = float32(v4)
                                        _ = v38
                                    }
                                case "foo bar":
                                    v39 := GenS5{F0: -718, F1: (2147483647 > 1), F2: (v27 << (((v27 - v27)) & 63))}
                                    fmt.Println(v4, fi26, v5)
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v39
                                }
                                _ = fi35
                            }
                        }
                    case 273:
                        if v5 {
                            for fi40 := 0; fi40 < 16; fi40++ {
                                var v41 float32 = (0.0 + 1.5)
                                var v42 uint64 = 834
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = fi40
                                _ = v41
                                _ = v42
                            }
                        }
                    }
                }
                v27 = (^(v27))
            }
            _ = fi26
            _ = v27
        }
    } else {
        for fi43 := 0; fi43 < 4; fi43++ {
            _ = len(v2)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            for fi44 := 0; fi44 < 2; fi44++ {
                if len(v6) > 0 {
                    v6[0] = fi43
                }
                if (1e10 < -2.25) {
                    v5 = ((!v5) && false)
                } else {
                    v3 = (v3 >> ((724) & 63))
                    if v5 {
                        v45 := GenS1{F0: 42, F1: 0}
                        _ = v45
                    } else {
                        v5 = (1 <= -1)
                        switch "a" {
                        case "a":
                            if true {
                                var v46 uint16 = 0
                                v2[fi44] = v5
                                v46 = v46
                                if (!v5) {
                                    var v47 int16 = -1
                                    _ = v47
                                }
                                _ = v46
                            }
                        case "hello":
                            v48 := v1[fi44]
                            _ = v48
                        case "\n\t":
                            var v49 float32 = ((float32((1e10 * ((v4 * 0.1) - 1.5))) * ((-2.25 * 1.5) / 1e10)) + float32(v4))
                            _ = v49
                        default:
                            v3 = (v3 - 2)
                        }
                    }
                }
                v50 := GenS5{F0: 1, F1: v5, F2: 65535}
                _ = fi44
                _ = v50
            }
            _ = fi43
        }
        for fi51 := 0; fi51 < 6; fi51++ {
            v5 = true
            v3 = v3
            v52 := GenS6{F0: -571, F1: 2147483647}
            _ = fi51
            _ = v52
        }
    }
    for fi53 := 0; fi53 < 15; fi53++ {
        for fi54 := 0; fi54 < 2; fi54++ {
            v6 = append(v6, (-(fi53)))
            if v5 {
                _ = len(v6)
                v55 := GenS7{F0: 605}
                v3 = 32767
                _ = v55
            } else {
                fmt.Println(fi54)
            }
            _ = fi54
        }
        var v56 float32 = (1.5 * -2.25)
        _ = fi53
        _ = v56
    }
    if (32767 > -1) {
        switch "hello" {
        case "":
            delete(v1, 127)
        case "hello":
            if ((v4 + v4) != ((v4 - (v4 - 3.14159)) / (float64(((((((3.14159 / 1e10) - 0.1) - 1.5) * 0.1) / 1.0) + 1.5)) / v4))) {
                if len(v6) > 0 {
                    v6[0] = -158
                }
            } else {
                var v57 uint16 = 278
                v58 := GenS4{F0: -451, F1: 42}
                v6 = append(v6, 42)
                _ = v57
                _ = v58
            }
        }
    }
    for fi59 := 0; fi59 < 9; fi59++ {
        v5 = (42 != -128)
        v60 := GenS7{F0: 4294967295}
        _ = fi59
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
