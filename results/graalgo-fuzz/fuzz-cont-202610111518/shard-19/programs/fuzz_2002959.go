package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int16; F1 int32; F2 int8 }
type GenS5 struct { F0 uint16; F1 string; F2 string }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int8; F1 uint8 }
func genFn0() bool {
    switch "foo bar" {
    case "\n\t":
        var v4 uint8 = 72
        _ = v4
    case "a":
        var v5 float32 = 1.5
        _ = v5
    }
    v6 := []int{0, 32768, -513}
    _ = v6
    return false
}
func genFn1(p0 uint32, p1 string) bool {
    if (2147483648 <= p0) {
        if (!(1.5 != (float32(1.5) * 0.1))) {
            for fi7 := 0; fi7 < 17; fi7++ {
                p0 -= p0
                _ = fi7
            }
        } else {
            p1 = p1
            for fi8 := 0; fi8 < 5; fi8++ {
                if (-1 < -919) {
                    v9 := GenS1{F0: p1, F1: -1}
                    var v10 string = (("\n\t" + "") + (p1 + p1))
                    _ = v10
                    _ = v9
                }
                _ = fi8
            }
        }
        v11 := []int{-2147483648, -894, -1}
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 14; fi13++ {
                fmt.Println(fi13, p1)
                var v14 float64 = (-2.25 + 1.5)
                _ = fi13
                _ = v14
            }
            p1 = ""
            _ = fi12
        }
        _ = v11
    }
    v15 := GenS1{F0: "hello", F1: 9223372036854775807}
    var v16 uint = 1
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return true
}

func main() {
    var v1 float32 = 3.14159
    v2 := []int{32768}
    v3 := []int{}
    v17 := GenS2{F0: -1, F1: 2}
    var v18 uint64 = 42
    switch ("foo bar" + "\n\t") {
    case "a":
        v1 -= v1
    case "":
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 int16 = 42
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 uint64 = v18
                _ = fi21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
    default:
        switch "hello" {
        case "foo bar":
            v23 := GenS1{F0: "hello", F1: -1}
            _ = v23
        case "\n\t":
            fmt.Println(v18, v1)
        case "":
            for fi24 := 0; fi24 < 11; fi24++ {
                v25 := GenS3{F0: -728, F1: ((v18 % ((v18) | 1)) > (v18 | v18))}
                for fi26 := 0; fi26 < 6; fi26++ {
                    fmt.Println(fi26, fi26, v1)
                    v27 := GenS4{F0: 1, F1: 1, F2: 100}
                    _ = fi26
                    _ = v27
                }
                for fi28 := 0; fi28 < 6; fi28++ {
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v29 int16 = -545
                        v18 += (v18 >> ((4294967295) & 63))
                        _ = v29
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v30 uint8 = 34
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                        _ = v30
                    }
                    if (888 <= 1) {
                        v18 += (v18 % ((v18) | 1))
                        v31 := GenS2{F0: -527, F1: 127}
                        var v32 uint16 = 412
                        _ = v31
                        _ = v32
                    } else {
                        if false {
                            if true {
                                if (!false) {
                                    if (v18 >= (v18 + v18)) {
                                        fmt.Println(fi28)
                                        v33 := GenS2{F0: 432, F1: 0}
                                        _ = v33
                                    } else {
                                        var v34 bool = (2 != 405)
                                        if len(v3) > 0 {
                                            v3[0] = fi28
                                        }
                                        v35 := GenS3{F0: -717, F1: v34}
                                        _ = v34
                                        _ = v35
                                    }
                                    var v36 string = "a"
                                    _ = v36
                                }
                            } else {
                                switch ("foo bar" + "\n\t") {
                                case "foo bar":
                                    if true {
                                        v18 = 18446744073709551615
                                    }
                                case "a":
                                    var v37 float32 = v1
                                    _ = v37
                                }
                            }
                            if ((fi24 % (((-(fi28))) | 1)) > (-(fi24))) {
                                if ((((^(fi24)) != (fi24 >> ((1) & 63))) || (255 < 111)) || (25 < 18446744073709551615)) {
                                    var v38 float64 = 3.14159
                                    _ = v38
                                } else {
                                    fmt.Println(v18, fi28, fi28)
                                }
                                var v39 int32 = -651
                                var v40 float32 = v1
                                var v41 int = (fi24 ^ (fi28 >> ((560) & 63)))
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                            var v42 uint16 = 812
                            _ = v42
                        }
                        switch (-(fi28)) {
                        case 42:
                            v18 += (v18 >> ((0) & 63))
                        default:
                            v1 = v1
                        }
                    }
                    var v43 float32 = ((1e10 - ((v1 * float32(-2.25)) * ((v1 - (3.14159 + (-2.25 - 3.14159))) / v1))) + 3.14159)
                    v44 := GenS4{F0: 1, F1: -913, F2: -1}
                    v1 = 3.14159
                    _ = fi28
                    _ = v43
                    _ = v44
                }
                _ = fi24
                _ = v25
            }
        }
    }
    fmt.Println(v18, v18)
    for fi45 := 0; fi45 < 6; fi45++ {
        v18 = v18
        if len(v2) > 0 {
            v2[0] = (fi45 | -257)
        }
        switch fi45 {
        case -1:
            var v46 int8 = 42
            _ = v46
        default:
            v3 = append(v3, fi45)
        }
        _ = fi45
    }
    var v47 uint = 983
    var v48 float64 = (1e10 + 3.14159)
    var v49 float64 = v48
    for fi50 := 0; fi50 < 3; fi50++ {
        if len(v2) > 0 {
            v2[0] = fi50
        }
        fmt.Println(v49, v47, v47)
        var v51 int = (fi50 << ((v47) & 63))
        _ = fi50
        _ = v51
    }
    v49 = (-2.25 + float64(0.0))
    if false {
        _ = len(v2)
        switch -781 {
        case 127:
            v18 += 2
        }
        switch -86 {
        case 73:
            if false {
                v48 -= (0.1 - (float64((0.0 + v1)) - v49))
                for fi52 := 0; fi52 < 19; fi52++ {
                    v48 -= v49
                    _ = fi52
                }
                if ((-32768 <= -1) || ((((v18 * (v18 >> (((^(v47))) & 63))) > v18) && false) && (!(525 <= 65535)))) {
                    switch ((((((("hello" + "") + "\n\t") + (("" + "hello") + (("\n\t" + "a") + ("a" + "a")))) + "") + "foo bar") + "foo bar") + "") {
                    case "\n\t":
                        v53 := GenS4{F0: 0, F1: 528, F2: 47}
                        _ = v53
                    case "":
                        for fi54 := 0; fi54 < 8; fi54++ {
                            v55 := GenS5{F0: 907, F1: ("a" + "foo bar"), F2: "hello"}
                            var v56 int8 = -128
                            _ = fi54
                            _ = v55
                            _ = v56
                        }
                    case "foo bar":
                        _ = len(v3)
                    default:
                        for fi57 := 0; fi57 < 12; fi57++ {
                            v1 += (v1 - float32(v49))
                            _ = fi57
                        }
                    }
                    if (!(true || (2147483647 <= 457))) {
                        v58 := GenS6{F0: 81, F1: 1000}
                        var v59 float32 = v1
                        v48 = float64((0.1 * ((v59 / ((1e10 + float32(0.1)) * v59)) * v59)))
                        _ = v58
                        _ = v59
                    }
                } else {
                    v1 = v1
                }
            }
        case 42:
            _ = len(v3)
        default:
            v2 = append(v2, 128)
        }
    } else {
        v3 = append(v3, 127)
    }
    v60 := GenS5{F0: 0, F1: ("a" + "\n\t"), F2: ("" + "foo bar")}
    v61 := GenS7{F0: -68, F1: 161}
    v48 += v48
    var v62 int16 = 1
    fmt.Println(v18)
    var v63 uint64 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
