package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 string; F2 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint; F1 int16; F2 int8 }
type GenS8 struct { F0 int8 }
func genFn0(p0 uint, p1 uint, p2 uint8) int {
    v7 := GenS1{F0: -1, F1: p2, F2: -537}
    fmt.Println(p0, p2)
    p2 -= (p2 | 127)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 bool, p1 uint64, p2 uint) uint8 {
    for fi8 := 0; fi8 < 2; fi8++ {
        switch "" {
        case "hello":
            v9 := GenS2{F0: 1}
            _ = v9
        case "foo bar":
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := GenS2{F0: 0}
                _ = fi10
                _ = v11
            }
        }
        p0 = p0
        p1 -= (p1 * p1)
        _ = fi8
    }
    p0 = (1 < 0)
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            fmt.Println(p1, p2)
            p1 -= (p1 - 4294967295)
            switch fi13 {
            case 128:
                var v14 int = genFn0((p2 >> ((4294967295) & 63)), 18446744073709551615, 127)
                _ = v14
            case 127:
                if ((("\n\t" + ("\n\t" + "a")) + "hello") <= "a") {
                    v15 := []int{}
                    fmt.Println(p2)
                    _ = v15
                } else {
                    if (-767 <= 1000) {
                        if (!((1 > 108) || (32767 > 127))) {
                            v16 := GenS3{F0: p1, F1: 124, F2: 1}
                            p2 += p2
                            _ = v16
                        }
                        p2 = (p2 << ((0) & 63))
                    } else {
                        v17 := GenS1{F0: 1, F1: 224, F2: 42}
                        var v18 int = (fi12 ^ (-(fi13)))
                        switch "hello" {
                        case "hello":
                            v19 := []int{(fi12 >> (((p2 - (p2 - p2))) & 63)), (fi13 >> (((p2 & p2)) & 63))}
                            _ = v19
                        case "\n\t":
                            v20 := GenS4{F0: 2}
                            _ = v20
                        default:
                            v21 := []int{genFn0(p2, p2, 1), genFn0(p2, 0, 42), fi12, fi12}
                            _ = v21
                        }
                        if (157 > 1) {
                            var v22 uint = (p2 & 18446744073709551615)
                            _ = v22
                        }
                        var v23 uint16 = 2
                        _ = v17
                        _ = v18
                        _ = v23
                    }
                    p1 += (p1 - p1)
                }
            case -1:
                p2 = (^(p2))
                v24 := []int{}
                if p0 {
                    p2 = p2
                    fmt.Println(p2, p1, p0)
                }
                _ = v24
            default:
                switch (-(fi12)) {
                case 729:
                    var v25 uint8 = uint8(p1)
                    _ = v25
                case -9223372036854775808:
                    v26 := []int{}
                    _ = v26
                case 32767:
                    p2 -= (p2 << ((p2) & 63))
                default:
                    v27 := []int{(fi13 ^ (-(fi13))), (fi12 ^ (fi13 << ((p2) & 63)))}
                    _ = v27
                }
            }
            _ = fi13
        }
        if true {
            v28 := []int{genFn0((p2 / ((p2) | 1)), (^(p2)), 2), genFn0((p2 - (p2 * 249)), p2, 2), fi12, fi12}
            switch "\n\t" {
            case "\n\t":
                if (259 == p1) {
                    var v29 uint8 = uint8(p1)
                    switch "foo bar" {
                    case "hello":
                        if p0 {
                            v29 = v29
                        } else {
                            var v30 int16 = 0
                            fmt.Println(p0, p0, v30)
                            v31 := GenS5{F0: -1, F1: "hello", F2: p0}
                            _ = v30
                            _ = v31
                        }
                    }
                    fmt.Println(v29, p1)
                    _ = v29
                }
            case "foo bar":
                p2 -= (^(p2))
            case "a":
                fmt.Println(p0)
            }
            switch ((("" + "\n\t") + ("\n\t" + "\n\t")) + "") {
            case "":
                if p0 {
                    var v32 int64 = 0
                    for fi33 := 0; fi33 < 10; fi33++ {
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                        p1 -= (p1 * 18446744073709551615)
                        fmt.Println(p0)
                        _ = fi33
                    }
                    p0 = (42 != -789)
                    _ = v32
                }
            case "a":
                var v34 string = ("hello" + (("foo bar" + "hello") + "foo bar"))
                _ = v34
            case "\n\t":
                if false {
                    for fi35 := 0; fi35 < 17; fi35++ {
                        p1 = (^(p1))
                        v28 = append(v28, 4294967295)
                        var v36 int = (-(fi12))
                        _ = fi35
                        _ = v36
                    }
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                    fmt.Println(p1)
                }
            }
            _ = v28
        }
        _ = fi12
    }
    var v37 uint64 = (p1 & (p1 & (p1 >> ((124) & 63))))
    v38 := GenS2{F0: int64(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    _ = v38
    return 127
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []uint32{42, 905, 4294967295}
    v3 := []string{"a", (("\n\t" + "") + "")}
    var v4 int64 = 42
    v5 := []string{"\n\t", "foo bar", "foo bar"}
    v6 := make(map[uint32]bool)
    var v39 float32 = 1.5
    v40 := GenS4{F0: 1}
    v4 = v4
    if (11 > 4294967295) {
        if (32768 >= -9223372036854775808) {
            v41 := GenS6{F0: true}
            switch ("" + "a") {
            case "\n\t":
                _ = len(v6)
            }
            fmt.Println(v39, v39, v1)
            _ = v41
        }
        switch 42 {
        case 679:
            for fi42 := 0; fi42 < 15; fi42++ {
                _ = len(v2)
                if (((true && true) && (((!false) && false) && (!(334 != 4294967295)))) || false) {
                    var v43 uint = 2
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                    switch "a" {
                    case "hello":
                        v4 += 518
                    case "\n\t":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "foo bar":
                        var v44 float32 = (0.1 * 0.1)
                        _ = v44
                    }
                    _ = v43
                }
                var v45 uint64 = 42
                _ = fi42
                _ = v45
            }
        case -9223372036854775808:
            var v46 int64 = v1
            _ = v46
        }
        if false {
            var v47 int64 = (v4 ^ 42)
            for fi48 := 0; fi48 < 9; fi48++ {
                _ = len(v6)
                v1 = (v47 & v47)
                v49 := v6[910]
                _ = fi48
                _ = v49
            }
            _ = v47
        } else {
            if true {
                fmt.Println(v1, v39)
                if true {
                    switch 128 {
                    case -606:
                        v4 += (v1 & -313)
                    }
                } else {
                    var v50 uint32 = 2
                    delete(v6, (^(v50)))
                    _ = v50
                }
            }
            delete(v6, 892)
        }
    }
    for fi51 := 0; fi51 < 5; fi51++ {
        switch "foo bar" {
        case "\n\t":
            v4 = -511
        case "a":
            var v52 uint16 = 674
            _ = v52
        case "":
            var v53 float64 = 0.1
            _ = v53
        default:
            switch (fi51 << ((669) & 63)) {
            case 2147483647:
                for fi54 := 0; fi54 < 13; fi54++ {
                    delete(v6, 65535)
                    v55 := GenS7{F0: 2, F1: -38, F2: -1}
                    var v56 uint64 = 709
                    v57 := GenS5{F0: 42, F1: ("foo bar" + "a"), F2: (false && true)}
                    var v58 uint16 = 643
                    _ = fi54
                    _ = v55
                    _ = v56
                    _ = v57
                    _ = v58
                }
            case 4294967295:
                var v59 uint8 = 232
                _ = v59
            case 9223372036854775807:
                for fi60 := 0; fi60 < 11; fi60++ {
                    v1 = v1
                    v4 += (^(v1))
                    _ = fi60
                }
            }
        }
        if (false && (0.0 == float32(3.14159))) {
            v1 = v1
            v1 += (v1 & (v1 << ((255) & 63)))
        }
        v39 = float32(-2.25)
        _ = fi51
    }
    v4 -= v4
    var v61 float32 = ((v39 / -2.25) / ((v39 / float32(0.1)) * (v39 * float32(0.1))))
    v1 = v1
    v4 = (v4 | 308)
    switch 503 {
    case -2147483648:
        switch ("foo bar" + "a") {
        case "\n\t":
            var v62 int = -1
            _ = v62
        case "hello":
            if ((!(v4 < (v1 | (v4 ^ 0)))) && (!(-313 != -44))) {
                var v63 int16 = -32768
                fmt.Println(v61, v61)
                _ = v63
            } else {
                if (1 <= 255) {
                    v64 := v6[2]
                    _ = v64
                }
            }
        }
    case 496:
        v6[2147483648] = false
    }
    if len(v5) > 0 {
        v5[0] = "a"
    }
    for fi65 := 0; fi65 < 8; fi65++ {
        var v66 string = "foo bar"
        _ = fi65
        _ = v66
    }
    switch ("hello" + "") {
    case "a":
        v39 = -2.25
    case "foo bar":
        _ = len(v6)
    case "":
        var v67 uint8 = 2
        _ = v67
    default:
        var v68 int = 2147483647
        _ = v68
    }
    v39 = v39
    v69 := GenS7{F0: 252, F1: -780, F2: 0}
    v70 := GenS8{F0: 127}
    for fi71 := 0; fi71 < 11; fi71++ {
        switch genFn0(42, 275, 1) {
        case 9223372036854775807:
            v6[210] = (!(!false))
        case 501:
            if true {
                fmt.Println(fi71)
                var v72 uint64 = 606
                _ = v72
            }
        }
        var v73 float64 = (-2.25 + (float64(v61) + 0.0))
        _ = fi71
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v61
    _ = v69
    _ = v70
}
