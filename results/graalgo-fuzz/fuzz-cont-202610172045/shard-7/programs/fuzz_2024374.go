package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }
type GenS3 struct { F0 int64; F1 int8; F2 uint }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int8; F1 int8; F2 uint16 }
func genFn0(p0 uint32, p1 int64) uint16 {
    fmt.Println(p1, p0)
    var v6 uint8 = 255
    v7 := []int{1, 424}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint8, p1 int32, p2 uint16) uint64 {
    p1 = 0
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint32 = 65535
    var v2 float32 = (0.0 + 3.14159)
    var v3 float32 = 0.0
    var v4 uint8 = 0
    var v5 int16 = 0
    v4 = 255
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            switch "hello" {
            case "":
                var v10 int8 = 42
                _ = v10
            }
            _ = fi9
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if false {
            switch (-(fi11)) {
            case 42:
                var v12 float32 = 0.1
                _ = v12
            case 691:
                v13 := []int{(fi11 / ((128) | 1)), 4294967295, -2147483648}
                _ = v13
            case 4294967295:
                var v14 int8 = -128
                _ = v14
            }
        } else {
            v15 := []int{128, (fi11 ^ 2147483647), (fi11 ^ 42)}
            for fi16 := 0; fi16 < 6; fi16++ {
                v3 = (v3 - (v2 + v3))
                _ = fi16
            }
            switch fi11 {
            case 523:
                v17 := GenS1{F0: (v4 << ((0) & 63)), F1: v1, F2: (^(v4))}
                _ = v17
            case -2147483648:
                v1 = 1
            case 42:
                if true {
                    _ = len(v15)
                    var v18 uint64 = 1
                    _ = len(v15)
                    _ = v18
                } else {
                    if true {
                        v1 = (v1 / ((v1) | 1))
                        var v19 uint = 18446744073709551615
                        _ = v19
                    } else {
                        v5 = (^(v5))
                        if ((v5 & v5) < (v5 >> ((255) & 63))) {
                            v5 = (^(v5))
                            v2 = -2.25
                            for fi20 := 0; fi20 < 1; fi20++ {
                                if true {
                                    v2 = float32(3.14159)
                                    v15 = append(v15, fi20)
                                    v21 := GenS2{F0: 595, F1: false, F2: 236}
                                    v22 := GenS3{F0: 0, F1: 42, F2: 4294967295}
                                    v3 += v2
                                    v23 := GenS4{F0: (v1 - v1), F1: 1}
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                } else {
                                    fmt.Println(v5)
                                }
                                fmt.Println(v1)
                                _ = fi20
                            }
                        }
                    }
                    var v24 uint16 = 255
                    fmt.Println(v3)
                    _ = v24
                }
            }
            _ = v15
        }
        switch "\n\t" {
        case "\n\t":
            v25 := []int{(fi11 / (((fi11 + (fi11 | fi11))) | 1))}
            _ = v25
        case "foo bar":
            if (!((0 <= 1) && (((genFn0(v1, 535) >= 639) && false) || (-2.25 != 1e10)))) {
                var v26 int8 = 42
                v27 := GenS5{F0: (^(v26)), F1: (v26 & v26), F2: 745}
                _ = v26
                _ = v27
            } else {
                var v28 int = (-(fi11))
                v3 = (1e10 / (v2 - float32(0.1)))
                _ = v28
            }
        case "a":
            for fi29 := 0; fi29 < 9; fi29++ {
                if ((fi11 <= fi29) && (32767 == 42)) {
                    v5 = -406
                } else {
                    var v30 uint8 = 0
                    v31 := GenS2{F0: -168, F1: (16 > -1), F2: 255}
                    v32 := GenS3{F0: -814, F1: 100, F2: 42}
                    _ = v30
                    _ = v31
                    _ = v32
                }
                v4 = 185
                if ((fi29 ^ fi11) >= -683) {
                    v33 := GenS1{F0: v4, F1: (v1 + uint32(v4)), F2: v4}
                    _ = v33
                }
                _ = fi29
            }
        }
        v34 := GenS2{F0: -1, F1: true, F2: 1}
        _ = fi11
        _ = v34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        fmt.Println(fi35, v2)
        var v36 int16 = v5
        v5 = v36
        _ = fi35
        _ = v36
    }
    if true {
        v1 = v1
        fmt.Println(v5)
        fmt.Println(v3)
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        if (true && true) {
            v38 := []int{(fi37 ^ (fi37 & fi37)), fi37, fi37, 9223372036854775807}
            v39 := GenS2{F0: 9223372036854775807, F1: ((2 >= 18446744073709551615) || (true || ((true && (v5 != v5)) && ((!(-991 < 9223372036854775807)) && (!(v5 >= 690)))))), F2: 2}
            v5 += 989
            _ = v38
            _ = v39
        } else {
            var v40 float64 = (-2.25 + 3.14159)
            fmt.Println(v5, fi37)
            _ = v40
        }
        _ = fi37
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        v4 = v4
        _ = fi41
    }
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
