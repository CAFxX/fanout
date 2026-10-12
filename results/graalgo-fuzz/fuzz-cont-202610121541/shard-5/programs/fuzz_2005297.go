package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 string; F1 uint16; F2 bool }
type GenS4 struct { F0 string; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint32; F1 bool }
func genFn0(p0 int8, p1 uint, p2 uint8) uint16 {
    if (p1 == (p1 + p1)) {
        var v5 uint64 = 2
        for fi6 := 0; fi6 < 1; fi6++ {
            p1 += p1
            var v7 bool = true
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    v8 := GenS1{F0: -99, F1: 109, F2: 1}
    switch "" {
    case "hello":
        for fi9 := 0; fi9 < 14; fi9++ {
            if ((3.14159 > 0.0) || ((0.1 / (1e10 - -2.25)) > float64(0.1))) {
                if true {
                    for fi10 := 0; fi10 < 12; fi10++ {
                        var v11 uint32 = 1
                        v12 := []int{fi9}
                        if false {
                            switch fi10 {
                            case 431:
                                p2 = p2
                            case -9223372036854775808:
                                v13 := GenS1{F0: p0, F1: 652, F2: 2}
                                _ = v13
                            case -272:
                                switch ("hello" + "") {
                                case "":
                                    var v14 float64 = -2.25
                                    _ = v14
                                }
                            default:
                                v15 := GenS2{F0: p2, F1: "foo bar"}
                                _ = v15
                            }
                            var v16 uint16 = 8
                            p2 = p2
                            _ = v16
                        }
                        var v17 uint64 = 2
                        var v18 int64 = 9223372036854775807
                        if len(v12) > 0 {
                            v12[0] = 555
                        }
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v17
                        _ = v18
                    }
                    if ("foo bar" > "\n\t") {
                        v19 := []int{(fi9 | fi9), fi9, (fi9 << (((p1 >> ((p1) & 63))) & 63))}
                        _ = v19
                    }
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(p1, p2)
                    case "a":
                        v20 := GenS1{F0: p0, F1: 1000, F2: 2}
                        _ = v20
                    }
                } else {
                    var v21 float32 = -2.25
                    var v22 uint8 = (p2 - (^(p2)))
                    _ = v21
                    _ = v22
                }
                v23 := []int{fi9, (-(fi9)), (fi9 % (((fi9 | (fi9 & (fi9 ^ fi9)))) | 1))}
                var v24 uint8 = (p2 - 2)
                _ = v23
                _ = v24
            } else {
                v25 := GenS1{F0: 127, F1: 32767, F2: 2}
                _ = v25
            }
            if (-2147483648 > 42) {
                v26 := GenS1{F0: -1, F1: 1000, F2: 2}
                _ = v26
            }
            v27 := []int{(^(fi9))}
            _ = fi9
            _ = v27
        }
    case "foo bar":
        if (((2147483647 < -2147483648) && (1 <= 2147483647)) || true) {
            var v28 int8 = p0
            p1 = 4294967295
            for fi29 := 0; fi29 < 13; fi29++ {
                var v30 uint8 = (p2 / (((^(p2))) | 1))
                _ = fi29
                _ = v30
            }
            _ = v28
        }
    }
    fmt.Println(p2, p0, p1)
    v31 := GenS2{F0: 0, F1: ""}
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    _ = v8
    return 1
}

func main() {
    var v1 float32 = 1e10
    v2 := []int{}
    var v3 int16 = 59
    v4 := make(map[uint32]bool)
    v3 = 1000
    if (((100 <= 0) && (42 >= -1)) && (42 == 255)) {
        fmt.Println(v3, v3)
    }
    var v32 int16 = (^(v3))
    if true {
        for fi33 := 0; fi33 < 4; fi33++ {
            if true {
                if (!false) {
                    var v34 float64 = 3.14159
                    v35 := GenS3{F0: "a", F1: genFn0(-1, 18446744073709551615, 39), F2: (2 < 208)}
                    switch (("foo bar" + (("\n\t" + "") + (("foo bar" + "foo bar") + "a"))) + "a") {
                    case "\n\t":
                        for fi36 := 0; fi36 < 18; fi36++ {
                            v37 := GenS3{F0: "hello", F1: 42, F2: (1 <= 4294967295)}
                            v32 -= v3
                            v38 := GenS1{F0: -1, F1: (v3 * (v3 & -374)), F2: 861}
                            switch "foo bar" {
                            case "a":
                                _ = len(v2)
                            }
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    case "foo bar":
                        if false {
                            switch ("foo bar" + ("hello" + "a")) {
                            case "\n\t":
                                if false {
                                    v39 := GenS3{F0: "", F1: 2, F2: true}
                                    for fi40 := 0; fi40 < 7; fi40++ {
                                        v32 += -32768
                                        fmt.Println(v1, fi40, v32)
                                        v32 = v32
                                        _ = fi40
                                    }
                                    _ = v39
                                }
                            case "foo bar":
                                v41 := GenS3{F0: "a", F1: 65535, F2: false}
                                _ = v41
                            default:
                                _ = len(v4)
                            }
                            _ = len(v4)
                        }
                    case "":
                        v42 := GenS2{F0: 2, F1: ("" + "")}
                        _ = v42
                    }
                    _ = v34
                    _ = v35
                }
                var v43 int64 = 0
                v3 = v32
                _ = v43
            }
            _ = fi33
        }
        var v44 uint8 = 2
        v45 := GenS4{F0: ("" + "a"), F1: 2, F2: -810}
        _ = v44
        _ = v45
    }
    v32 = v32
    v1 -= (v1 * -2.25)
    v46 := GenS4{F0: "\n\t", F1: 0, F2: 2147483647}
    var v47 uint32 = 0
    fmt.Println(v1)
    v32 -= 42
    v48 := GenS5{F0: v47, F1: false}
    delete(v4, v47)
    switch ("hello" + ("foo bar" + ("" + "hello"))) {
    case "":
        if false {
            var v49 float32 = (0.1 / v1)
            _ = v49
        } else {
            if ((v1 * (float32(0.0) / v1)) > v1) {
                _ = len(v4)
            }
        }
    default:
        var v50 int = -2147483648
        _ = v50
    }
    fmt.Println(v3)
    v51 := GenS4{F0: "a", F1: (^(v47)), F2: -372}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v51
}
