package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint64; F2 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 uint32; F1 uint64 }
func genFn0() int8 {
    var v4 int = 32767
    v5 := []int{(-(v4)), v4, 9223372036854775807}
    for fi6 := 0; fi6 < 17; fi6++ {
        v4 += (v4 ^ fi6)
        if (true && false) {
            if (false || true) {
                var v7 int64 = 263
                v7 += 9223372036854775807
                for fi8 := 0; fi8 < 20; fi8++ {
                    var v9 int = -9223372036854775808
                    v10 := GenS1{F0: 653}
                    _ = len(v5)
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = v7
            } else {
                var v11 bool = (false || (!true))
                var v12 uint32 = 710
                var v13 int64 = 83
                _ = v11
                _ = v12
                _ = v13
            }
            switch fi6 {
            case 32767:
                v4 -= 42
            }
        } else {
            fmt.Println(v4, fi6)
        }
        _ = fi6
    }
    var v14 uint = 4294967295
    _ = v14
    _ = v4
    _ = v5
    return int8(v4)
}
func genFn1(p0 uint, p1 uint64, p2 int32) uint {
    p0 = p0
    if true {
        p0 = (p0 | (^(p0)))
        v15 := []int{-2147483648, -9223372036854775808}
        _ = v15
    } else {
        fmt.Println(p2, p0)
        switch (("" + ((("hello" + ("hello" + ("a" + ("foo bar" + "\n\t")))) + ("" + "\n\t")) + "")) + "foo bar") {
        case "foo bar":
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS2{F0: fi16}
                v18 := GenS3{F0: 46, F1: uint64(p0), F2: (fi16 | (fi16 - -856))}
                _ = fi16
                _ = v17
                _ = v18
            }
        case "a":
            v19 := GenS2{F0: -647}
            _ = v19
        }
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    return (^(p0))
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 bool = false
    var v3 uint = 18446744073709551615
    if true {
        if (v1 < (v1 | (v1 + (^(v1))))) {
            switch "hello" {
            case "hello":
                if v2 {
                    v3 -= 667
                }
            }
        }
        var v20 float64 = (1e10 * 1e10)
        _ = v20
    } else {
        v1 = (v1 / ((v1) | 1))
        for fi21 := 0; fi21 < 2; fi21++ {
            switch "foo bar" {
            case "foo bar":
                v1 -= v1
            default:
                var v22 uint8 = 255
                _ = v22
            }
            _ = fi21
        }
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        v1 = v1
        var v24 string = (((("" + ("" + "\n\t")) + "foo bar") + "a") + "\n\t")
        _ = fi23
        _ = v24
    }
    var v25 uint32 = 4294967295
    v26 := GenS4{F0: 0}
    if false {
        for fi27 := 0; fi27 < 13; fi27++ {
            v25 += v25
            switch (fi27 + fi27) {
            case 9223372036854775807:
                v3 -= genFn1(1, uint64(v25), 1)
            case 374:
                for fi28 := 0; fi28 < 17; fi28++ {
                    var v29 uint = 480
                    var v30 uint16 = 32767
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
            _ = fi27
        }
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 int16 = -32768
            for fi33 := 0; fi33 < 11; fi33++ {
                switch ("" + "hello") {
                case "":
                    v3 = v3
                default:
                    if (!(639 != 2)) {
                        v1 += (-(v1))
                    } else {
                        v2 = v2
                        var v34 int = 2147483647
                        var v35 uint16 = 998
                        _ = v34
                        _ = v35
                    }
                }
                _ = fi33
            }
            _ = fi31
            _ = v32
        }
    } else {
        for fi36 := 0; fi36 < 19; fi36++ {
            fmt.Println(v2)
            _ = fi36
        }
    }
    var v37 int16 = 32767
    v38 := GenS5{F0: 42, F1: 100}
    v39 := []int{-2147483648, 0}
    switch "a" {
    case "hello":
        var v40 int8 = -114
        _ = v40
    case "foo bar":
        if (196 >= 255) {
            if true {
                v2 = v2
            } else {
                v41 := GenS2{F0: 1}
                var v42 uint32 = v25
                if ((2 == genFn1((v3 & (v3 * v3)), 18446744073709551615, 42)) && (100 < 1)) {
                    _ = len(v39)
                    switch ((("hello" + "\n\t") + "hello") + "foo bar") {
                    case "foo bar":
                        if v2 {
                            var v43 uint64 = 18446744073709551615
                            if (!true) {
                                v42 += 2
                                v44 := GenS1{F0: genFn1(65535, v43, 42)}
                                v45 := GenS5{F0: 1, F1: -47}
                                _ = v44
                                _ = v45
                            } else {
                                var v46 float64 = ((float64(float32((0.0 * -2.25))) - 0.0) * (3.14159 + 1e10))
                                _ = v46
                            }
                            if ((^(v43)) < v43) {
                                if len(v39) >= 2 {
                                    v39 = v39[1:len(v39)]
                                }
                            }
                            _ = v43
                        }
                    case "a":
                        for fi47 := 0; fi47 < 3; fi47++ {
                            for fi48 := 0; fi48 < 13; fi48++ {
                                var v49 int = fi47
                                var v50 uint = v3
                                v51 := GenS4{F0: 544}
                                _ = fi48
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                            switch "a" {
                            case "foo bar":
                                for fi52 := 0; fi52 < 11; fi52++ {
                                    v53 := GenS6{F0: false, F1: ""}
                                    _ = fi52
                                    _ = v53
                                }
                            default:
                                v37 += (v37 * v37)
                            }
                            _ = fi47
                        }
                    default:
                        fmt.Println(v3, v25)
                    }
                    if v2 {
                        fmt.Println(v42)
                        v25 -= (v25 << (((v3 % ((genFn1(v3, 4294967295, int32(v3))) | 1))) & 63))
                        for fi54 := 0; fi54 < 6; fi54++ {
                            if len(v39) > 0 {
                                v39[0] = (fi54 - 42)
                            }
                            _ = fi54
                        }
                    }
                } else {
                    if len(v39) >= 2 {
                        v39 = v39[1:len(v39)]
                    }
                    var v55 int16 = -50
                    for fi56 := 0; fi56 < 12; fi56++ {
                        var v57 float64 = 0.1
                        for fi58 := 0; fi58 < 13; fi58++ {
                            var v59 int64 = -9223372036854775808
                            v3 = 2
                            v37 = int16(fi58)
                            var v60 bool = v2
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                        _ = fi56
                        _ = v57
                    }
                    _ = v55
                }
                _ = v41
                _ = v42
            }
            if (42 <= 42) {
                if (((2147483647 > 1) && true) && (v2 || true)) {
                    var v61 uint = (^(v3))
                    if v2 {
                        v37 += v37
                        v2 = true
                        var v62 int32 = 0
                        _ = v62
                    } else {
                        v37 += v37
                        v63 := GenS6{F0: v2, F1: "hello"}
                        _ = v63
                    }
                    _ = v61
                }
            } else {
                var v64 uint8 = 92
                _ = len(v39)
                v37 = v37
                _ = v64
            }
            v3 -= 0
        } else {
            var v65 int8 = genFn0()
            if (v2 && true) {
                for fi66 := 0; fi66 < 16; fi66++ {
                    v25 = (v25 + (^(v25)))
                    fmt.Println(v25, v25)
                    _ = fi66
                }
                if true {
                    v67 := GenS3{F0: 216, F1: 42, F2: 32767}
                    _ = v67
                }
            } else {
                if false {
                    if len(v39) >= 2 {
                        v39 = v39[1:len(v39)]
                    }
                    v65 -= (v65 & -28)
                    var v68 int16 = (v37 | -878)
                    _ = v68
                }
            }
            var v69 float32 = 3.14159
            _ = v65
            _ = v69
        }
    }
    if (1 != 128) {
        for fi70 := 0; fi70 < 15; fi70++ {
            for fi71 := 0; fi71 < 1; fi71++ {
                v39 = append(v39, (-(fi70)))
                fmt.Println(v2, v2, fi71)
                v37 += v37
                _ = fi71
            }
            _ = fi70
        }
        fmt.Println(v1, v25, v1)
    } else {
        var v72 string = ("\n\t" + "a")
        _ = v72
    }
    for fi73 := 0; fi73 < 2; fi73++ {
        if len(v39) > 0 {
            v39[0] = (^(fi73))
        }
        var v74 int32 = -1
        v2 = v2
        _ = fi73
        _ = v74
    }
    v2 = v2
    v75 := GenS7{F0: v25, F1: 995}
    switch 641 {
    case -656:
        v2 = true
    case 725:
        for fi76 := 0; fi76 < 8; fi76++ {
            v1 -= (v1 << (((v3 - (v3 * v3))) & 63))
            for fi77 := 0; fi77 < 4; fi77++ {
                if len(v39) > 0 {
                    v39[0] = 127
                }
                fmt.Println(v2)
                _ = fi77
            }
            _ = fi76
        }
    }
    if len(v39) > 0 {
        v39[0] = 4294967295
    }
    fmt.Println(v1)
    v1 = v1
    v39 = append(v39, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v39:", len(v39), v39)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v75
}
