package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint64; F2 bool }
func genFn0(p0 uint8) int {
    var v7 int = -251
    var v8 int8 = 100
    switch "\n\t" {
    case "a":
        var v9 uint32 = 4294967295
        _ = v9
    case "":
        v10 := GenS1{F0: 42}
        _ = v10
    case "\n\t":
        if (true && false) {
            switch "foo bar" {
            case "":
                if true {
                    v8 = (v8 >> ((4294967295) & 63))
                    v7 = v7
                    var v11 uint8 = 1
                    _ = v11
                } else {
                    if true {
                        switch (v7 / (((^(v7))) | 1)) {
                        case -1:
                            v8 -= v8
                        case -2147483648:
                            v12 := []int{v7}
                            _ = v12
                        default:
                            switch v7 {
                            case 556:
                                v13 := GenS1{F0: 1}
                                _ = v13
                            case 32768:
                                var v14 int = 127
                                _ = v14
                            case 0:
                                var v15 int = v7
                                _ = v15
                            default:
                                if true {
                                    var v16 float64 = 1.5
                                    _ = v16
                                } else {
                                    if true {
                                        fmt.Println(p0)
                                    }
                                }
                            }
                        }
                    }
                    v17 := []int{9223372036854775807, 677}
                    switch v7 {
                    case -822:
                        var v18 string = "\n\t"
                        _ = v18
                    case 9223372036854775807:
                        var v19 uint = 441
                        _ = v19
                    case 533:
                        var v20 int = (v7 - v7)
                        _ = v20
                    }
                    _ = v17
                }
            case "a":
                v7 = 42
            case "\n\t":
                v7 = (v7 % ((v7) | 1))
            }
        } else {
            v21 := GenS2{F0: -2147483648}
            var v22 uint16 = 582
            _ = v21
            _ = v22
        }
    default:
        var v23 int = (^(v7))
        _ = v23
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        v7 = fi24
        _ = fi24
    }
    var v25 int16 = 1000
    _ = p0
    _ = v25
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 int, p1 int8, p2 uint32) int64 {
    p1 = (p1 + p1)
    v26 := GenS1{F0: 0}
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return 9223372036854775807
}

func main() {
    v1 := []int{}
    v2 := []int{}
    v3 := make(map[string]int)
    var v4 int32 = 979
    v5 := []string{"hello", "", "foo bar"}
    var v6 bool = true
    v27 := GenS2{F0: 0}
    if v6 {
        v3[(((("" + "hello") + ("" + "hello")) + (("\n\t" + "\n\t") + "a")) + "hello")] = -132
        for fi28 := 0; fi28 < 19; fi28++ {
            if v6 {
                if (547 <= 65535) {
                    if v6 {
                        v4 = (v4 * v4)
                    } else {
                        v2 = append(v2, (fi28 ^ (fi28 | fi28)))
                        v29 := GenS3{F0: -89}
                        if (!v6) {
                            v30 := GenS4{F0: 42, F1: 1}
                            v31 := GenS1{F0: (v4 ^ v4)}
                            _ = v30
                            _ = v31
                        } else {
                            var v32 int64 = genFn1(fi28, -16, 65535)
                            _ = v32
                        }
                        _ = v29
                    }
                    v3[""] = fi28
                    var v33 string = "foo bar"
                    _ = v33
                } else {
                    switch (^(fi28)) {
                    case 621:
                        v4 += v4
                    case -1:
                        var v34 float32 = 3.14159
                        _ = v34
                    default:
                        var v35 uint32 = 459
                        _ = v35
                    }
                    if true {
                        v6 = (!(!true))
                        if (!(!v6)) {
                            v3[""] = fi28
                        }
                        var v36 int32 = (-(v4))
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 17; fi37++ {
                        v38 := GenS1{F0: (v4 + v4)}
                        _ = fi37
                        _ = v38
                    }
                }
            } else {
                v4 = v4
                v39 := v3[("a" + "\n\t")]
                _ = v39
            }
            _ = fi28
        }
        if v6 {
            if (1 <= 0) {
                if (855 > 741) {
                    fmt.Println(v6)
                }
            } else {
                v6 = v6
            }
            var v40 float32 = 1.5
            _ = v40
        } else {
            var v41 int64 = 566
            v42 := GenS2{F0: v4}
            _ = v41
            _ = v42
        }
    }
    var v43 float32 = float32(3.14159)
    v44 := GenS5{F0: 915, F1: 0, F2: (42 < 4294967295)}
    fmt.Println(v6)
    var v45 float32 = 0.0
    var v46 int = 1
    v4 = 0
    fmt.Println(v45)
    fmt.Println(v43, v43)
    v47 := v3["hello"]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
