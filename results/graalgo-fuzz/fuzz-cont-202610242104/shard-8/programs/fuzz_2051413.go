package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int; F2 uint16 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 int16; F2 int32 }
func genFn0() uint16 {
    var v4 int = 127
    fmt.Println(v4)
    if true {
        v4 = v4
    }
    v5 := GenS1{F0: 127}
    _ = v4
    _ = v5
    return 65535
}
func genFn1(p0 string, p1 uint64) uint64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 uint = 2
        _ = fi6
        _ = v7
    }
    switch "" {
    case "hello":
        switch p0 {
        case "":
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 float32 = (0.0 / 1e10)
                var v10 uint16 = 42
                _ = fi8
                _ = v10
                _ = v9
            }
        default:
            v11 := GenS2{F0: "", F1: 0, F2: 1}
            _ = v11
        }
    }
    switch 4294967295 {
    case 127:
        var v12 bool = (true || true)
        _ = v12
    }
    var v13 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v13
    return (p1 + p1)
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 0
    var v3 int = 2147483647
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            v2 = v2
            v3 = v3
            v2 -= (v2 << ((0) & 63))
            _ = fi15
        }
        v16 := []int{}
        _ = fi14
        _ = v16
    }
    _ = len(v1)
    var v17 float32 = 0.1
    if ((-2.25 - v17) < v17) {
        v18 := []int{32767, 0}
        var v19 uint = 18446744073709551615
        _ = v18
        _ = v19
    }
    v20 := []int{(v3 ^ v3), v3, v3, 128}
    if true {
        var v21 int8 = -1
        var v22 float64 = 0.1
        if (v21 == (v21 % ((v21) | 1))) {
            switch (-(v3)) {
            case 302:
                v2 = v2
            case -777:
                switch ((((("\n\t" + "a") + "foo bar") + "\n\t") + "\n\t") + "a") {
                case "":
                    v23 := GenS3{F0: "hello", F1: "\n\t"}
                    _ = v23
                case "foo bar":
                    v2 += (v2 % ((255) | 1))
                default:
                    v3 = v3
                }
            case -2147483648:
                _ = len(v20)
            default:
                var v24 int = 2147483647
                _ = v24
            }
            if false {
                var v25 int64 = 42
                _ = v25
            } else {
                if (!true) {
                    var v26 uint64 = 0
                    if false {
                        v27 := GenS3{F0: "foo bar", F1: "\n\t"}
                        _ = v27
                    } else {
                        v28 := GenS4{F0: -300}
                        v26 += (v26 + (v26 * genFn1("\n\t", v26)))
                        for fi29 := 0; fi29 < 11; fi29++ {
                            if len(v20) > 0 {
                                v20[0] = v3
                            }
                            v22 += v22
                            v30 := GenS2{F0: "foo bar", F1: 32768, F2: 32767}
                            _ = fi29
                            _ = v30
                        }
                        _ = v28
                    }
                    fmt.Println(v17, v3, v17)
                    _ = v26
                }
                switch (("hello" + ((("a" + ("foo bar" + "hello")) + "") + ("hello" + ("" + ("" + ("a" + "foo bar")))))) + "a") {
                case "foo bar":
                    for fi31 := 0; fi31 < 18; fi31++ {
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                        v2 -= (v2 & (v2 + (v2 | v2)))
                        _ = fi31
                    }
                case "\n\t":
                    switch 127 {
                    case 128:
                        fmt.Println(v3, v22)
                    case 983:
                        v32 := GenS2{F0: "hello", F1: (v3 * v3), F2: 162}
                        _ = v32
                    case -42:
                        v33 := GenS1{F0: 255}
                        _ = v33
                    default:
                        v21 += 58
                    }
                case "hello":
                    if false {
                        v20 = append(v20, int(v2))
                        for fi34 := 0; fi34 < 16; fi34++ {
                            if false {
                                var v35 bool = false
                                v21 += v21
                                _ = v35
                            }
                            v36 := GenS5{F0: 1, F1: 42, F2: -814}
                            _ = fi34
                            _ = v36
                        }
                        v1[(v3 << ((4294967295) & 63))] = (-(v3))
                    } else {
                        switch v3 {
                        case 32767:
                            for fi37 := 0; fi37 < 7; fi37++ {
                                v38 := GenS5{F0: 255, F1: 1000, F2: 846}
                                fmt.Println(v3, v3, v2)
                                _ = fi37
                                _ = v38
                            }
                        case -9223372036854775808:
                            var v39 uint64 = 4294967295
                            _ = v39
                        }
                    }
                }
                for fi40 := 0; fi40 < 15; fi40++ {
                    if (4294967295 > 42) {
                        _ = len(v1)
                    }
                    _ = fi40
                }
            }
        }
        _ = v21
        _ = v22
    } else {
        var v41 uint32 = 2
        for fi42 := 0; fi42 < 14; fi42++ {
            v3 -= (v3 | (v3 * (-(v3))))
            _ = fi42
        }
        _ = v41
    }
    var v43 bool = ((false && false) || true)
    if len(v20) > 0 {
        v20[0] = -843
    }
    switch v3 {
    case -9223372036854775808:
        v17 += v17
    case 32767:
        if (v17 == ((1.5 - (1e10 * v17)) - v17)) {
            fmt.Println(v17, v2, v3)
            var v44 int16 = -32768
            _ = v44
        } else {
            v45 := v1[v3]
            _ = v45
        }
    case 1000:
        v2 = v2
    default:
        v46 := GenS5{F0: 136, F1: -1, F2: 2147483647}
        _ = v46
    }
    var v47 int8 = 13
    if (42 <= 4294967295) {
        for fi48 := 0; fi48 < 11; fi48++ {
            if v43 {
                v17 -= (v17 / 3.14159)
                v43 = (-2147483648 != 1)
                v3 += v3
            }
            _ = fi48
        }
        _ = len(v20)
    } else {
        var v49 int = v3
        for fi50 := 0; fi50 < 14; fi50++ {
            v51 := GenS4{F0: -616}
            for fi52 := 0; fi52 < 1; fi52++ {
                v49 -= fi52
                var v53 float32 = v17
                v49 -= 161
                _ = fi52
                _ = v53
            }
            _ = fi50
            _ = v51
        }
        var v54 uint8 = (v2 * 42)
        _ = v49
        _ = v54
    }
    fmt.Println(v3, v2, v3)
    if v43 {
        v55 := GenS3{F0: ("hello" + ("" + "a")), F1: "hello"}
        _ = v55
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    if (1 <= 0) {
        fmt.Println(v47, v43, v2)
        var v56 uint16 = uint16(v2)
        _ = len(v20)
        _ = v56
    } else {
        for fi57 := 0; fi57 < 8; fi57++ {
            switch "" {
            case "a":
                v43 = ((405 >= 552) && v43)
            }
            var v58 uint = 18446744073709551615
            var v59 uint32 = 0
            _ = fi57
            _ = v58
            _ = v59
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v43
    _ = v47
}
