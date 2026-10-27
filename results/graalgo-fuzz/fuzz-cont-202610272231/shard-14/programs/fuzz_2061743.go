package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint8, p1 uint16) int64 {
    v7 := []int{4294967295}
    var v8 int64 = 1
    v9 := GenS1{F0: 9223372036854775807}
    v8 = (v8 - (v8 | (v8 | 9223372036854775807)))
    v10 := GenS1{F0: 301}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 string = "\n\t"
    var v2 int = 0
    v3 := make(map[string]bool)
    v4 := []int{-1, (v2 * (v2 ^ (v2 >> ((42) & 63)))), 2147483647, v2}
    var v5 uint16 = 786
    v6 := make(map[string]int)
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 int8 = -128
        var v13 int8 = 100
        _ = fi11
        _ = v12
        _ = v13
    }
    v3[v1] = true
    v14 := GenS1{F0: (v2 | v2)}
    if (false || false) {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 int8 = 0
            _ = fi15
            _ = v16
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v6[(v1 + ("a" + v1))] = (v2 >> ((255) & 63))
        _ = len(v3)
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        v1 = v1
        _ = fi17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        if ((!true) || true) {
            v2 = (v2 + 546)
        } else {
            var v19 int16 = 0
            for fi20 := 0; fi20 < 19; fi20++ {
                v2 = 42
                var v21 uint8 = 66
                v22 := v6[v1]
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v19
        }
        delete(v6, "foo bar")
        for fi23 := 0; fi23 < 12; fi23++ {
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            v24 := v6[("hello" + "\n\t")]
            if true {
                v25 := GenS1{F0: (v2 | v24)}
                _ = v25
            }
            _ = fi23
            _ = v24
        }
        _ = fi18
    }
    switch 868 {
    case 0:
        for fi26 := 0; fi26 < 20; fi26++ {
            var v27 bool = (65535 != 676)
            if (v5 == v5) {
                var v28 bool = true
                _ = v28
            } else {
                for fi29 := 0; fi29 < 5; fi29++ {
                    v4 = append(v4, fi29)
                    _ = fi29
                }
                v30 := GenS2{F0: 170, F1: 0, F2: v5}
                _ = v30
            }
            if true {
                v31 := GenS2{F0: 2147483647, F1: 0, F2: v5}
                var v32 uint32 = 42
                _ = v31
                _ = v32
            }
            _ = fi26
            _ = v27
        }
    default:
        v2 += 2147483647
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        v5 = (^(v5))
        if (false && (true || true)) {
            if true {
                for fi34 := 0; fi34 < 17; fi34++ {
                    switch (fi34 / ((v2) | 1)) {
                    case 128:
                        switch v1 {
                        case "foo bar":
                            fmt.Println(fi34, v5, fi33)
                        case "a":
                            switch (fi34 << ((199) & 63)) {
                            case 128:
                                v5 += (v5 / (((v5 % ((255) | 1))) | 1))
                            case 2147483647:
                                switch v2 {
                                case -743:
                                    if true {
                                        fmt.Println(fi34, fi34)
                                        v5 += 52
                                        var v35 uint64 = 1
                                        _ = v35
                                    } else {
                                        v5 -= v5
                                        switch -650 {
                                        case 127:
                                            v1 = (v1 + v1)
                                        }
                                        v2 += (fi33 << ((255) & 63))
                                        var v36 uint8 = 69
                                        _ = v36
                                    }
                                default:
                                    switch "hello" {
                                    case "":
                                        v37 := v3[v1]
                                        _ = v37
                                    case "foo bar":
                                        _ = len(v4)
                                    }
                                }
                            }
                        case "":
                            var v38 int16 = -32768
                            _ = v38
                        default:
                            v4 = append(v4, 4294967295)
                        }
                    case 519:
                        switch (^(fi34)) {
                        case 2147483647:
                            if true {
                                var v39 float32 = ((((((float32((((0.0 - (0.0 - float64(0.0))) - -2.25) * 0.0)) / -2.25) + 0.1) / 1.0) * float32((3.14159 + 0.0))) + (((((1.5 + (0.0 + 3.14159)) + 3.14159) * (3.14159 / -2.25)) * 0.0) * 0.1)) + -2.25)
                                v40 := v3[((("hello" + (v1 + (v1 + v1))) + v1) + ("" + (v1 + ((v1 + (v1 + "foo bar")) + "hello"))))]
                                _ = v39
                                _ = v40
                            }
                        }
                    case 378:
                        v41 := v3[(v1 + v1)]
                        _ = v41
                    default:
                        v42 := GenS2{F0: 938, F1: int32(v5), F2: (v5 ^ v5)}
                        _ = v42
                    }
                    if false {
                        var v43 float64 = 0.0
                        switch (v2 - fi33) {
                        case 9223372036854775807:
                            v4 = append(v4, fi34)
                        case -598:
                            switch ((v1 + (v1 + ("" + "foo bar"))) + "") {
                            case "foo bar":
                                var v44 int16 = 42
                                _ = v44
                            case "":
                                var v45 uint16 = (v5 | v5)
                                _ = v45
                            default:
                                var v46 uint64 = 252
                                _ = v46
                            }
                            var v47 uint8 = 127
                            var v48 int32 = -753
                            v49 := GenS3{F0: v5, F1: 42}
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                        _ = v43
                    } else {
                        v50 := GenS4{F0: 2}
                        v5 = (v5 ^ (v5 / (((v5 >> ((255) & 63))) | 1)))
                        _ = v50
                    }
                    _ = fi34
                }
                var v51 bool = true
                _ = v51
            } else {
                if (v5 < v5) {
                    if false {
                        var v52 float64 = float64((1e10 / 3.14159))
                        _ = v52
                    } else {
                        var v53 int8 = 0
                        _ = v53
                    }
                }
                v54 := v3[(v1 + "\n\t")]
                _ = v54
            }
            var v55 int64 = 1
            _ = v55
        }
        fmt.Println(v2, v5)
        _ = fi33
    }
    if false {
        var v56 int64 = -501
        fmt.Println(v2, v56, v2)
        _ = v56
    } else {
        for fi57 := 0; fi57 < 14; fi57++ {
            v2 -= (v2 + (v2 ^ fi57))
            _ = fi57
        }
        v58 := GenS5{F0: false}
        _ = v58
    }
    v59 := GenS2{F0: 4294967295, F1: 715, F2: 405}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
