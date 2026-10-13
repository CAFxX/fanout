package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int32; F2 bool }
type GenS6 struct { F0 uint; F1 int8; F2 int64 }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 int32; F1 int32; F2 string }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := []int{}
        for fi7 := 0; fi7 < 20; fi7++ {
            if (1 >= 2) {
                var v8 float64 = 3.14159
                _ = v8
            }
            var v9 int = (fi5 + 32768)
            if (0 < 61) {
                var v10 int16 = 32767
                v9 += (fi5 / (((fi5 % ((v9) | 1))) | 1))
                switch ((("" + "hello") + ("hello" + "hello")) + ((("\n\t" + "") + "") + "\n\t")) {
                case "hello":
                    if len(v6) > 0 {
                        v6[0] = fi5
                    }
                default:
                    _ = len(v6)
                }
                _ = v10
            } else {
                v11 := GenS1{F0: (-128 >= -83)}
                v12 := GenS2{F0: 0, F1: 1}
                fmt.Println(fi7)
                if false {
                    var v13 float64 = 1e10
                    _ = v13
                } else {
                    fmt.Println(v9)
                    v9 -= fi7
                    v9 -= 0
                }
                v9 += (v9 ^ 0)
                var v14 uint64 = 4294967295
                var v15 int8 = -1
                switch (-(v9)) {
                case 643:
                    if (!(true && (2 <= 42))) {
                        v16 := GenS3{F0: -584, F1: true}
                        switch "hello" {
                        case "\n\t":
                            var v17 int = (^(fi7))
                            _ = v17
                        case "hello":
                            fmt.Println(fi5, v14)
                        case "foo bar":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            fmt.Println(v15)
                            var v18 bool = true
                            _ = v18
                        }
                        _ = v16
                    } else {
                        var v19 int64 = 0
                        fmt.Println(fi5, fi7)
                        var v20 float64 = (0.0 * 1e10)
                        _ = v19
                        _ = v20
                    }
                case -372:
                    v21 := GenS1{F0: false}
                    var v22 uint16 = 32767
                    v23 := GenS1{F0: (!((!((v14 ^ v14) < 83)) || true))}
                    _ = v21
                    _ = v22
                    _ = v23
                default:
                    _ = len(v6)
                }
                var v24 uint = 255
                _ = v11
                _ = v12
                _ = v14
                _ = v15
                _ = v24
            }
            switch ("" + ("" + "")) {
            case "":
                var v25 int16 = int16(v9)
                _ = v25
            case "hello":
                if true {
                    if true {
                        var v26 int8 = 14
                        switch v9 {
                        case -499:
                            fmt.Println(v26, v9, fi7)
                        case -729:
                            var v27 uint32 = 2147483647
                            _ = v27
                        default:
                            if len(v6) > 0 {
                                v6[0] = v9
                            }
                        }
                        _ = v26
                    } else {
                        if (!false) {
                            _ = len(v6)
                            v28 := GenS3{F0: -2147483648, F1: false}
                            v9 += -998
                            v9 += 9223372036854775807
                            if (9223372036854775807 >= 42) {
                                v29 := GenS1{F0: true}
                                v9 += 0
                                v9 -= v9
                                _ = v29
                            } else {
                                v30 := GenS4{F0: 1}
                                switch 9223372036854775807 {
                                case 897:
                                    v31 := GenS3{F0: 0, F1: true}
                                    _ = v31
                                case -2147483648:
                                    v9 = (^(fi7))
                                default:
                                    if (1 >= 15) {
                                        _ = len(v6)
                                        v9 -= v9
                                        var v32 int64 = int64(fi7)
                                        _ = v32
                                    } else {
                                        fmt.Println(fi5, v9)
                                    }
                                }
                                if (1 >= 186) {
                                    v33 := GenS3{F0: 218, F1: (32767 != 610)}
                                    if true {
                                        if len(v6) > 0 {
                                            v6[0] = (fi7 - (v9 << ((255) & 63)))
                                        }
                                    } else {
                                        var v34 uint = 613
                                        _ = v34
                                    }
                                    fmt.Println(v9, v9, v9)
                                    v9 -= (fi7 >> ((2) & 63))
                                    _ = v33
                                }
                                _ = v30
                            }
                            _ = v28
                        }
                        var v35 int8 = 127
                        if len(v6) > 0 {
                            v6[0] = fi5
                        }
                        _ = v35
                    }
                    v9 -= (fi7 + (-(fi7)))
                    v9 = (v9 ^ fi7)
                    v6 = append(v6, fi5)
                }
            case "\n\t":
                v9 += fi5
            default:
                var v36 uint = 958
                _ = v36
            }
            v6 = append(v6, 621)
            _ = fi7
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    var v37 float32 = 1e10
    fmt.Println(v37)
    var v38 uint64 = 18446744073709551615
    _ = v37
    _ = v38
    return 0
}
func genFn1(p0 bool, p1 int, p2 string) int {
    for fi39 := 0; fi39 < 7; fi39++ {
        v40 := GenS5{F0: p1, F1: 42, F2: (false || (4294967295 == 2))}
        switch (fi39 >> ((255) & 63)) {
        case 42:
            p1 -= p1
        default:
            v41 := []int{fi39, -1}
            _ = v41
        }
        _ = fi39
        _ = v40
    }
    switch p2 {
    case "":
        if p0 {
            fmt.Println(p1)
            p0 = p0
            p1 = p1
        } else {
            var v42 float64 = ((0.1 * (0.0 - -2.25)) * -2.25)
            v43 := []int{(p1 * (^(p1))), (-(p1))}
            if len(v43) >= 2 {
                v43 = v43[1:len(v43)]
            }
            _ = v42
            _ = v43
        }
    case "foo bar":
        v44 := GenS6{F0: 775, F1: 42, F2: 1}
        _ = v44
    case "a":
        v45 := GenS4{F0: 128}
        _ = v45
    }
    var v46 uint = 485
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v46
    return 32768
}

func main() {
    var v1 float64 = 0.0
    var v2 int8 = 42
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    delete(v3, 74)
    v2 = -17
    if (!(0 < 1)) {
        if false {
            switch 0 {
            case -1:
                if false {
                    switch "a" {
                    case "a":
                        if (!false) {
                            v1 = v1
                        } else {
                            fmt.Println(v2)
                        }
                    case "hello":
                        _ = len(v4)
                    }
                    var v47 float32 = float32(float64(3.14159))
                    _ = v47
                }
            }
            if (((true && (-238 > 2147483647)) && true) || (988 < 18446744073709551615)) {
                var v48 int16 = -192
                fmt.Println(v1, v2)
                if (true || false) {
                    if false {
                        fmt.Println(v2)
                    } else {
                        switch -2147483648 {
                        case 42:
                            v2 -= v2
                        case -528:
                            v1 = v1
                        }
                        if true {
                            v4[""] = 2147483648
                            v2 -= (^(v2))
                            var v49 uint16 = 591
                            _ = v49
                        } else {
                            var v50 uint8 = 42
                            _ = v50
                        }
                    }
                    for fi51 := 0; fi51 < 9; fi51++ {
                        for fi52 := 0; fi52 < 20; fi52++ {
                            v48 -= (v48 * (v48 * (-(v48))))
                            v53 := GenS1{F0: (("" != "foo bar") || true)}
                            fmt.Println(v2, v2, v1)
                            _ = fi52
                            _ = v53
                        }
                        if true {
                            v2 -= v2
                            for fi54 := 0; fi54 < 19; fi54++ {
                                v2 = (v2 % (((v2 % ((v2) | 1))) | 1))
                                _ = fi54
                            }
                        }
                        v55 := []int{(fi51 % (((-(fi51))) | 1)), fi51}
                        _ = fi51
                        _ = v55
                    }
                }
                _ = v48
            }
            v1 = v1
        } else {
            var v56 int16 = 1000
            _ = v56
        }
        if true {
            v57 := GenS7{F0: 541, F1: 556}
            v4[""] = 2
            var v58 float64 = (float64(0.0) / 3.14159)
            _ = v57
            _ = v58
        } else {
            v59 := GenS4{F0: 0}
            _ = v59
        }
    } else {
        v2 = -128
        v1 -= 1e10
    }
    if false {
        v1 = (v1 * 3.14159)
        switch -798 {
        case 32767:
            v60 := GenS1{F0: (false || (("" <= (("" + ("a" + "foo bar")) + ("foo bar" + "\n\t"))) || false))}
            _ = v60
        case 9223372036854775807:
            for fi61 := 0; fi61 < 13; fi61++ {
                v62 := GenS2{F0: 0, F1: 2147483647}
                v63 := GenS8{F0: 2147483647, F1: -2147483648, F2: (("a" + (((((("hello" + "foo bar") + "\n\t") + "foo bar") + "") + ("a" + (("" + "hello") + "a"))) + "foo bar")) + ("\n\t" + "hello"))}
                _ = fi61
                _ = v62
                _ = v63
            }
        }
    }
    v64 := GenS6{F0: 42, F1: v2, F2: -1}
    delete(v4, "\n\t")
    if ((!true) && ((true || true) || false)) {
        fmt.Println(v1, v2)
        var v65 uint16 = 0
        _ = v65
    }
    var v66 uint16 = 255
    v67 := GenS1{F0: (!(false || true))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v66
    _ = v67
}
