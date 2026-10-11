package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int8; F2 uint }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int64 }
func genFn0(p0 string, p1 uint8) uint {
    p0 = p0
    fmt.Println(p0, p0)
    v6 := GenS1{F0: 290}
    _ = p0
    _ = p1
    _ = v6
    return 2
}
func genFn1(p0 bool) int32 {
    v7 := []int{32768, -9223372036854775808, 128, 42}
    var v8 float32 = (((0.0 * (((3.14159 / float32(3.14159)) / 0.1) / 1e10)) / 1.0) * float32(1e10))
    v9 := GenS1{F0: -1}
    if (-524 <= -32768) {
        _ = len(v7)
        if (-468 >= 1) {
            for fi10 := 0; fi10 < 4; fi10++ {
                if false {
                    if false {
                        if false {
                            v11 := GenS1{F0: -9223372036854775808}
                            _ = v11
                        }
                    } else {
                        switch (fi10 + fi10) {
                        case 127:
                            var v12 float64 = 0.1
                            _ = v12
                        case 128:
                            switch "" {
                            case "hello":
                                var v13 uint16 = 2
                                _ = v13
                            case "\n\t":
                                var v14 float32 = ((1.5 * 0.0) * v8)
                                _ = v14
                            case "foo bar":
                                switch ("a" + "a") {
                                case "a":
                                    v7 = append(v7, 32768)
                                case "foo bar":
                                    var v15 int64 = -318
                                    _ = v15
                                case "hello":
                                    v8 = v8
                                default:
                                    if (true && p0) {
                                        switch (fi10 - 9223372036854775807) {
                                        case -9223372036854775808:
                                            var v16 float32 = 0.0
                                            _ = v16
                                        case 9223372036854775807:
                                            switch "a" {
                                            case "a":
                                                var v17 uint16 = 0
                                                _ = v17
                                            case "":
                                                var v18 bool = p0
                                                _ = v18
                                            default:
                                                if len(v7) > 0 {
                                                    v7[0] = (^(fi10))
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        case -2147483648:
                            var v19 uint64 = 2
                            _ = v19
                        default:
                            if (57 < 127) {
                                var v20 float32 = -2.25
                                for fi21 := 0; fi21 < 11; fi21++ {
                                    var v22 uint16 = 2
                                    v23 := GenS1{F0: -439}
                                    v7 = append(v7, -159)
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                                fmt.Println(p0)
                                _ = v20
                            }
                        }
                        v8 = (v8 / (0.1 + 0.1))
                        var v24 int = (fi10 >> ((genFn0(("a" + "a"), 204)) & 63))
                        _ = v24
                    }
                } else {
                    var v25 float64 = -2.25
                    _ = v25
                }
                for fi26 := 0; fi26 < 2; fi26++ {
                    var v27 int8 = 76
                    v8 += v8
                    v7 = append(v7, (fi26 + fi10))
                    _ = fi26
                    _ = v27
                }
                for fi28 := 0; fi28 < 18; fi28++ {
                    p0 = (1 == 42)
                    v8 -= (v8 * v8)
                    p0 = true
                    v29 := GenS1{F0: -1}
                    p0 = p0
                    _ = fi28
                    _ = v29
                }
                _ = fi10
            }
            switch "hello" {
            case "":
                v30 := GenS2{F0: (("foo bar" + "") + ""), F1: true, F2: 127}
                _ = v30
            }
        } else {
            for fi31 := 0; fi31 < 11; fi31++ {
                v32 := GenS3{F0: 42}
                p0 = (2147483647 != 911)
                _ = fi31
                _ = v32
            }
        }
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []int{128}
    var v2 uint = 255
    v3 := []int{32768, 128, 32767}
    var v4 int16 = 42
    v5 := make(map[uint32]int)
    var v33 uint = (v2 >> ((v2) & 63))
    v4 = (v4 << ((65535) & 63))
    switch "hello" {
    case "a":
        if (true || (18446744073709551615 < 2)) {
            v34 := GenS1{F0: -885}
            fmt.Println(v2, v4)
            var v35 string = (("foo bar" + "foo bar") + "a")
            _ = v34
            _ = v35
        } else {
            fmt.Println(v33, v33)
            if ((true || false) && (true && true)) {
                delete(v5, 622)
            } else {
                var v36 int8 = 0
                switch "hello" {
                case "\n\t":
                    fmt.Println(v2, v36)
                }
                _ = v36
            }
            v37 := GenS4{F0: (v4 * 32767), F1: 127, F2: v33}
            _ = v37
        }
    case "foo bar":
        for fi38 := 0; fi38 < 1; fi38++ {
            for fi39 := 0; fi39 < 6; fi39++ {
                v2 -= (v33 % (((v2 - (^(v2)))) | 1))
                v33 = v33
                _ = fi39
            }
            _ = fi38
        }
    default:
        v40 := GenS2{F0: "\n\t", F1: (!true), F2: 1}
        _ = v40
    }
    switch 9223372036854775807 {
    case 32768:
        v41 := GenS2{F0: "a", F1: ((v2 | 42) < v2), F2: 52}
        _ = v41
    case 490:
        if true {
            if true {
                v33 = (v2 % (((v33 << (((v2 << ((v33) & 63))) & 63))) | 1))
                switch "a" {
                case "":
                    v33 = (v33 + (v2 ^ (v2 >> ((162) & 63))))
                }
            }
            var v42 string = ("" + "\n\t")
            for fi43 := 0; fi43 < 18; fi43++ {
                v4 -= v4
                v1 = append(v1, (fi43 * 1))
                _ = fi43
            }
            _ = v42
        }
    case -345:
        for fi44 := 0; fi44 < 16; fi44++ {
            v45 := GenS3{F0: 2}
            v46 := GenS3{F0: 0}
            _ = len(v5)
            _ = fi44
            _ = v45
            _ = v46
        }
    default:
        v33 = (v33 % ((genFn0(("\n\t" + ""), uint8(v33))) | 1))
    }
    v2 += v33
    fmt.Println(v2, v4, v4)
    var v47 bool = ((!(!true)) || ("" < ("" + "foo bar")))
    for fi48 := 0; fi48 < 6; fi48++ {
        v2 = v2
        _ = fi48
    }
    for fi49 := 0; fi49 < 11; fi49++ {
        v50 := GenS2{F0: "a", F1: v47, F2: 61}
        v51 := GenS1{F0: 0}
        _ = fi49
        _ = v50
        _ = v51
    }
    var v52 int8 = -1
    var v53 int64 = 9223372036854775807
    fmt.Println(v52)
    if (2 <= 515) {
        var v54 uint32 = 2
        if true {
            fmt.Println(v4, v53, v33)
        }
        var v55 uint32 = (v54 + v54)
        _ = v54
        _ = v55
    } else {
        v56 := GenS5{F0: 1, F1: ""}
        _ = v56
    }
    delete(v5, 42)
    for fi57 := 0; fi57 < 13; fi57++ {
        if true {
            if v47 {
                var v58 uint16 = 42
                v5[958] = int(v58)
                _ = v58
            }
            switch ((((((("a" + "\n\t") + ("hello" + (("" + ("" + "hello")) + ("hello" + ("" + "\n\t"))))) + (("" + "") + ("a" + ""))) + ("foo bar" + ("hello" + ""))) + "a") + "foo bar") + "\n\t") {
            case "":
                var v59 int32 = -144
                _ = v59
            case "hello":
                v52 += (v52 << (((v2 >> ((172) & 63))) & 63))
            case "a":
                v60 := GenS6{F0: v53}
                _ = v60
            default:
                fmt.Println(fi57, v52)
            }
        }
        _ = fi57
    }
    switch "" {
    case "a":
        if false {
            fmt.Println(v47, v2)
        }
    default:
        for fi61 := 0; fi61 < 4; fi61++ {
            var v62 int16 = 1000
            _ = fi61
            _ = v62
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v47
    _ = v5
    _ = v52
    _ = v53
}
