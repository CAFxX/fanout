package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 bool; F1 int64; F2 int8 }
type GenS5 struct { F0 int; F1 string; F2 uint }
func genFn0(p0 uint16) bool {
    v6 := []int{}
    var v7 int = 128
    _ = p0
    _ = v6
    _ = v7
    return (!(65535 < 274))
}
func genFn1(p0 uint32) int64 {
    v8 := []int{61, -644}
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(fi9, p0)
        fmt.Println(p0, fi9)
        _ = fi9
    }
    fmt.Println(p0)
    _ = len(v8)
    _ = p0
    _ = v8
    return 0
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := []int{128}
    var v5 uint = 18446744073709551615
    v10 := GenS1{F0: ("" + ""), F1: 360}
    var v11 uint = (v5 + (v5 % (((v5 - (^(v5)))) | 1)))
    fmt.Println(v5)
    v11 = 18446744073709551615
    v5 -= v11
    v12 := v2[2]
    v13 := GenS2{F0: -128, F1: -682}
    v12 = (42 > 100)
    var v14 uint16 = 872
    if v12 {
        switch 2147483647 {
        case 128:
            switch ("" + ("foo bar" + "hello")) {
            case "a":
                v5 += v11
            case "foo bar":
                for fi15 := 0; fi15 < 17; fi15++ {
                    fmt.Println(v5, v14, v12)
                    fmt.Println(fi15)
                    if (!v12) {
                        if (true && true) {
                            if (18446744073709551615 <= v5) {
                                switch fi15 {
                                case 42:
                                    if len(v4) > 0 {
                                        v4[0] = (fi15 * -1)
                                    }
                                case 2147483647:
                                    var v16 int64 = -1
                                    _ = v16
                                case 798:
                                    v14 += uint16(v11)
                                }
                                v14 = 0
                                for fi17 := 0; fi17 < 18; fi17++ {
                                    v18 := GenS1{F0: "", F1: 1}
                                    v19 := GenS2{F0: 42, F1: 2147483647}
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                            }
                        }
                    } else {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v12 = v12
                            fmt.Println(v5, fi15)
                            _ = fi20
                        }
                    }
                    _ = fi15
                }
            case "hello":
                if (!(!true)) {
                    fmt.Println(v11, v5)
                }
            }
        }
        v11 += v5
    }
    if v12 {
        if len(v4) > 0 {
            v4[0] = 1
        }
        var v21 uint8 = 2
        _ = v21
    }
    v22 := GenS3{F0: v5, F1: "foo bar"}
    v11 -= (v11 >> ((0) & 63))
    if (((1 != -2147483648) || (889 <= 866)) || (4294967295 >= 18446744073709551615)) {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 bool = (!true)
            v25 := GenS4{F0: (0 >= 152), F1: 42, F2: 87}
            _ = fi23
            _ = v24
            _ = v25
        }
        v11 += 255
    } else {
        for fi26 := 0; fi26 < 10; fi26++ {
            for fi27 := 0; fi27 < 5; fi27++ {
                var v28 uint = (v5 ^ (v11 % (((^(v11))) | 1)))
                _ = fi27
                _ = v28
            }
            if (!(127 >= 128)) {
                for fi29 := 0; fi29 < 12; fi29++ {
                    v30 := GenS2{F0: 127, F1: -1}
                    v4 = append(v4, fi29)
                    _ = fi29
                    _ = v30
                }
                v12 = true
            }
            _ = fi26
        }
        var v31 bool = (1e10 <= 1.5)
        _ = v31
    }
    if (-60 == 100) {
        v5 -= (v11 << (((v5 | (^(v5)))) & 63))
        for fi32 := 0; fi32 < 17; fi32++ {
            switch 2147483647 {
            case 4294967295:
                v33 := GenS1{F0: ("a" + ("" + "\n\t")), F1: -1}
                _ = v33
            case -9223372036854775808:
                v12 = (4294967295 != 2)
            case 1:
                for fi34 := 0; fi34 < 20; fi34++ {
                    var v35 uint = v5
                    if (2147483647 <= -902) {
                        v36 := GenS5{F0: fi34, F1: "a", F2: (v5 * 255)}
                        _ = v36
                    }
                    var v37 uint = v5
                    _ = fi34
                    _ = v35
                    _ = v37
                }
            default:
                var v38 uint8 = 42
                _ = v38
            }
            v5 -= 18446744073709551615
            _ = fi32
        }
        for fi39 := 0; fi39 < 11; fi39++ {
            if ((fi39 | (fi39 + fi39)) >= (fi39 % (((^(fi39))) | 1))) {
                v1 = append(v1, fi39)
            }
            v1 = append(v1, fi39)
            _ = fi39
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
