package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 string }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 string, p1 int32) uint8 {
    var v7 int = -395
    var v8 int16 = 1
    for fi9 := 0; fi9 < 8; fi9++ {
        if true {
            var v10 string = "\n\t"
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := GenS1{F0: -1, F1: 2}
                if false {
                    if false {
                        v13 := []int{}
                        fmt.Println(p0, fi9, v7)
                        _ = v13
                    }
                    var v14 string = "a"
                    switch fi11 {
                    case 337:
                        v15 := GenS2{F0: 203, F1: 485, F2: v14}
                        _ = v15
                    case -9223372036854775808:
                        if false {
                            fmt.Println(v7)
                            v8 = (v8 * v8)
                        } else {
                            if (!(1 == -192)) {
                                v16 := []int{fi11, fi9, (v7 >> ((4294967295) & 63))}
                                _ = v16
                            } else {
                                if false {
                                    v17 := GenS3{F0: p1, F1: 1000}
                                    if true {
                                        var v18 int64 = 0
                                        _ = v18
                                    }
                                    var v19 uint64 = 293
                                    _ = v17
                                    _ = v19
                                }
                            }
                        }
                    }
                    _ = v14
                }
                v20 := GenS4{F0: 0}
                v7 -= fi11
                p1 -= (p1 & p1)
                p0 = (((("" + "hello") + p0) + (v10 + p0)) + p0)
                _ = fi11
                _ = v12
                _ = v20
            }
            switch -822 {
            case 9223372036854775807:
                v8 = 1000
            case 127:
                v8 = v8
            case 128:
                v8 = 487
            }
            _ = v10
        }
        v21 := []int{v7, fi9, v7}
        _ = fi9
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int8, p1 uint, p2 bool) int16 {
    var v22 string = "foo bar"
    var v23 string = "foo bar"
    v24 := GenS3{F0: -2147483648, F1: -1}
    if true {
        v22 = "hello"
        v25 := GenS2{F0: 1, F1: 42, F2: ("\n\t" + v23)}
        _ = v25
    } else {
        if p2 {
            var v26 uint8 = 42
            _ = v26
        } else {
            var v27 float64 = (0.1 * 0.0)
            p1 = (p1 * p1)
            fmt.Println(v23, p2)
            _ = v27
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    _ = v24
    return 729
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 789
    var v3 int32 = -2147483648
    var v4 int16 = (v2 + v2)
    var v5 float32 = (1e10 / 1e10)
    var v6 uint8 = 2
    var v28 uint64 = 1
    v29 := GenS5{F0: "\n\t", F1: 2}
    v30 := GenS6{F0: genFn0("foo bar", v3), F1: "foo bar"}
    v31 := []int{}
    if (!(false && (112 < 127))) {
        v4 = -669
        v32 := GenS7{F0: -128}
        _ = v32
    }
    v3 = v3
    var v33 bool = (false || ((true && true) && (-1 == 9223372036854775807)))
    fmt.Println(v5, v4, v5)
    v5 = v5
    var v34 float32 = v5
    v6 = v6
    if ((-(v4)) > 1) {
        switch ("a" + "a") {
        case "\n\t":
            for fi35 := 0; fi35 < 12; fi35++ {
                v36 := GenS7{F0: -20}
                _ = fi35
                _ = v36
            }
        }
        switch ("a" + "hello") {
        case "a":
            switch -979 {
            case 42:
                fmt.Println(v5)
            default:
                for fi37 := 0; fi37 < 7; fi37++ {
                    if (((v28 + (^(v28))) == v28) || false) {
                        if ((!true) && v33) {
                            if len(v31) >= 2 {
                                v31 = v31[1:len(v31)]
                            }
                            var v38 int64 = 42
                            fmt.Println(v34)
                            _ = v38
                        } else {
                            v39 := GenS8{F0: 899}
                            var v40 uint = 271
                            for fi41 := 0; fi41 < 14; fi41++ {
                                var v42 string = "foo bar"
                                v5 += -2.25
                                var v43 uint16 = 255
                                v28 = 764
                                fmt.Println(v3, v28, v43)
                                var v44 int = fi37
                                _ = fi41
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            _ = v39
                            _ = v40
                        }
                    }
                    _ = fi37
                }
            }
        case "\n\t":
            var v45 int8 = 94
            _ = v45
        default:
            if ((false || v33) && false) {
                v3 = (v3 | v3)
            } else {
                v6 = v6
                v2 -= (v2 / ((v4) | 1))
            }
        }
    }
    if len(v31) > 0 {
        v31[0] = 127
    }
    var v46 uint8 = v6
    var v47 int64 = 9223372036854775807
    if (!(v33 || false)) {
        v4 = (v2 % ((v4) | 1))
    }
    v34 = v5
    if v33 {
        delete(v1, 2147483647)
        v34 = v34
        switch "foo bar" {
        case "hello":
            v31 = append(v31, 2147483647)
        case "":
            for fi48 := 0; fi48 < 18; fi48++ {
                fmt.Println(v4, v34)
                var v49 uint8 = v6
                _ = fi48
                _ = v49
            }
        case "\n\t":
            fmt.Println(v2, v47)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
