package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int32; F1 int; F2 int64 }
type GenS5 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint; F1 bool; F2 string }
func genFn0(p0 int, p1 string) int {
    v7 := []int{p0, p0}
    if true {
        var v8 int32 = -2147483648
        fmt.Println(v8, v8)
        p1 = p1
        _ = v8
    } else {
        p0 = p0
        if true {
            p1 = (p1 + p1)
            var v9 int = (p0 * (p0 + 9223372036854775807))
            _ = v9
        } else {
            if false {
                for fi10 := 0; fi10 < 20; fi10++ {
                    var v11 int32 = -1
                    if true {
                        p0 += fi10
                    }
                    _ = fi10
                    _ = v11
                }
                if false {
                    p0 = -9223372036854775808
                    v12 := GenS1{F0: 42}
                    _ = v12
                } else {
                    p1 = p1
                    switch ((((p1 + "\n\t") + ((("foo bar" + p1) + p1) + (("a" + ("" + p1)) + ("a" + (p1 + "foo bar"))))) + ((p1 + "foo bar") + p1)) + (p1 + ((p1 + p1) + ((p1 + (p1 + p1)) + (("foo bar" + "hello") + (("hello" + p1) + p1)))))) {
                    case "foo bar":
                        var v13 int32 = 1
                        _ = v13
                    }
                    switch p1 {
                    case "hello":
                        var v14 uint16 = 0
                        _ = v14
                    case "\n\t":
                        for fi15 := 0; fi15 < 15; fi15++ {
                            var v16 int32 = 2147483647
                            _ = fi15
                            _ = v16
                        }
                    case "foo bar":
                        p0 -= (p0 | 128)
                    default:
                        if len(v7) > 0 {
                            v7[0] = p0
                        }
                    }
                }
            } else {
                var v17 float32 = 0.0
                _ = v17
            }
            v18 := GenS2{F0: p1, F1: (42 > 574)}
            _ = v18
        }
        switch p0 {
        case -2147483648:
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 int8 = 1
                if false {
                    fmt.Println(p1)
                }
                v21 := GenS3{F0: 517, F1: 42}
                _ = fi19
                _ = v20
                _ = v21
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 933
}
func genFn1(p0 int, p1 bool, p2 int64) int32 {
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 14; fi23++ {
            p2 = p2
            v24 := []int{(p0 / ((p0) | 1)), fi22}
            if p1 {
                var v25 uint8 = 0
                _ = v25
            } else {
                p0 -= fi22
            }
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    if false {
        for fi26 := 0; fi26 < 4; fi26++ {
            if (0.0 < -2.25) {
                if p1 {
                    v27 := GenS3{F0: 0, F1: 2}
                    _ = v27
                }
                p2 = p2
            } else {
                p0 += (fi26 & fi26)
                if (int32(p2) != 766) {
                    p0 += p0
                    p0 = (p0 / (((p0 * (^(fi26)))) | 1))
                    if p1 {
                        switch fi26 {
                        case 42:
                            for fi28 := 0; fi28 < 15; fi28++ {
                                var v29 uint32 = 837
                                if p1 {
                                    v30 := GenS2{F0: "", F1: false}
                                    p0 = (fi26 * genFn0((fi26 * fi26), ("a" + ("a" + "hello"))))
                                    var v31 uint64 = 18446744073709551615
                                    _ = v30
                                    _ = v31
                                } else {
                                    v32 := GenS2{F0: "\n\t", F1: true}
                                    var v33 uint8 = 2
                                    _ = v32
                                    _ = v33
                                }
                                _ = fi28
                                _ = v29
                            }
                        default:
                            var v34 int64 = (p2 * p2)
                            _ = v34
                        }
                        var v35 uint32 = 2147483647
                        _ = v35
                    } else {
                        p2 -= -358
                        p1 = (!(255 < 42))
                    }
                } else {
                    var v36 uint32 = 0
                    v37 := []int{p0}
                    _ = v36
                    _ = v37
                }
            }
            if (!(0.1 <= -2.25)) {
                v38 := []int{}
                _ = v38
            }
            _ = fi26
        }
    } else {
        v39 := GenS4{F0: 2147483647, F1: p0, F2: (p2 % ((42) | 1))}
        var v40 uint8 = 1
        v41 := GenS1{F0: uint(v40)}
        _ = v39
        _ = v40
        _ = v41
    }
    _ = p0
    _ = p1
    _ = p2
    return 665
}

func main() {
    var v1 uint64 = 1
    var v2 uint = 18446744073709551615
    var v3 string = "foo bar"
    v4 := []string{v3, "foo bar"}
    var v5 int32 = 2147483647
    v6 := make(map[uint32]string)
    v42 := GenS3{F0: 0, F1: (^(v1))}
    v1 -= v1
    v43 := GenS4{F0: v5, F1: genFn0(42, "\n\t"), F2: 1}
    switch 1 {
    case 9223372036854775807:
        v44 := GenS1{F0: v2}
        _ = v44
    default:
        var v45 int8 = 42
        _ = v45
    }
    var v46 float64 = ((1.5 - 0.1) / (1.5 - (0.0 * 3.14159)))
    v47 := GenS3{F0: 540, F1: (v1 & 18446744073709551615)}
    if false {
        for fi48 := 0; fi48 < 10; fi48++ {
            v49 := GenS5{F0: 65535, F1: (!(uint64(v2) <= (^(v1)))), F2: 564}
            v50 := GenS6{F0: v2, F1: (genFn1(fi48, true, 93) >= v5), F2: v3}
            _ = len(v6)
            _ = fi48
            _ = v49
            _ = v50
        }
    } else {
        v6[65535] = v3
    }
    v6[42] = (v3 + v3)
    v3 = v3
    v4 = append(v4, (v3 + (v3 + v3)))
    if (-2.25 < 3.14159) {
        v51 := GenS3{F0: 2147483648, F1: v1}
        _ = len(v6)
        _ = v51
    }
    v6[65535] = (("" + v3) + "hello")
    if (!(v3 > ((v3 + "") + ((v3 + "foo bar") + "\n\t")))) {
        switch 128 {
        case 2147483647:
            fmt.Println(v5)
        case 32767:
            v5 = v5
        case 4294967295:
            v46 = v46
        default:
            _ = len(v4)
        }
        var v52 int8 = 100
        _ = v52
    }
    var v53 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v46
    _ = v47
    _ = v5
    _ = v53
    _ = v6
}
