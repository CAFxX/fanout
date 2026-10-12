package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 uint }
type GenS3 struct { F0 bool; F1 bool; F2 int8 }
type GenS4 struct { F0 int8; F1 uint8 }
func genFn0() int8 {
    var v4 bool = (-1 != 0)
    var v5 float32 = 0.1
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(v5, fi6)
            v4 = (v4 || false)
            v8 := []int{(-(fi7))}
            _ = fi7
            _ = v8
        }
        switch -281 {
        case 9223372036854775807:
            var v9 bool = (18446744073709551615 > 783)
            _ = v9
        }
        _ = fi6
    }
    var v10 int16 = 853
    for fi11 := 0; fi11 < 1; fi11++ {
        if (false && false) {
            if false {
                var v12 uint32 = 989
                _ = v12
            }
            var v13 int64 = 196
            _ = v13
        } else {
            var v14 string = "a"
            v4 = v4
            switch ("hello" + (((((v14 + v14) + v14) + (v14 + ((v14 + (v14 + v14)) + "\n\t"))) + (v14 + "\n\t")) + v14)) {
            case "foo bar":
                if v4 {
                    var v15 uint64 = 893
                    if (v4 || false) {
                        var v16 int32 = int32(v10)
                        for fi17 := 0; fi17 < 9; fi17++ {
                            v16 -= (v16 * 1)
                            _ = fi17
                        }
                        v18 := []int{-9223372036854775808, (-(fi11)), 986}
                        _ = v16
                        _ = v18
                    }
                    if v4 {
                        if false {
                            if (-2.25 <= 3.14159) {
                                if (2 < 2) {
                                    v4 = v4
                                    v4 = false
                                    v10 = (v10 + -646)
                                }
                            } else {
                                v10 = (v10 & (v10 + (v10 & v10)))
                            }
                            if true {
                                v15 += 4294967295
                                var v19 float32 = -2.25
                                v20 := []int{fi11, 127}
                                _ = v19
                                _ = v20
                            }
                        }
                        fmt.Println(v10, v4, v5)
                        if (-35 == 0) {
                            var v21 uint = 2
                            if (-2147483648 <= -811) {
                                v5 -= ((v5 / (float32(1e10) - v5)) * (v5 / float32(0.1)))
                            }
                            _ = v21
                        } else {
                            fmt.Println(v10, v4, v14)
                        }
                    } else {
                        var v22 uint = 18446744073709551615
                        _ = v22
                    }
                    _ = v15
                } else {
                    var v23 int64 = -642
                    v24 := []int{fi11}
                    var v25 bool = (v10 < 0)
                    _ = v23
                    _ = v24
                    _ = v25
                }
            default:
                fmt.Println(v5, v10)
            }
            _ = v14
        }
        switch fi11 {
        case 128:
            v26 := GenS1{F0: fi11, F1: 42, F2: 2}
            _ = v26
        case 42:
            var v27 uint8 = 253
            _ = v27
        default:
            if (65535 > 42) {
                v5 += (v5 - 1.5)
            }
        }
        v28 := []int{}
        _ = fi11
        _ = v28
    }
    _ = v10
    _ = v4
    _ = v5
    return -128
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 100
    v3 := []string{(("a" + ("" + "a")) + "a"), "foo bar"}
    for fi29 := 0; fi29 < 6; fi29++ {
        fmt.Println(fi29, fi29)
        v2 += v2
        _ = fi29
    }
    v30 := GenS1{F0: 42, F1: 42, F2: 65535}
    if ((v2 - (v2 * 42)) <= (v2 << ((255) & 63))) {
        for fi31 := 0; fi31 < 6; fi31++ {
            var v32 int = fi31
            _ = fi31
            _ = v32
        }
        _ = len(v1)
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        v2 -= (v2 & v2)
        _ = fi33
    }
    v34 := GenS2{F0: -314, F1: -32768, F2: 393}
    for fi35 := 0; fi35 < 15; fi35++ {
        if true {
            v36 := GenS3{F0: (1e10 != 3.14159), F1: (4294967295 <= 890), F2: v2}
            var v37 int64 = 42
            _ = v36
            _ = v37
        }
        fmt.Println(v2)
        _ = fi35
    }
    v2 -= (v2 * v2)
    var v38 bool = (true && false)
    v1[0] = 42
    switch 32768 {
    case -1:
        v38 = (-1 != 1)
    }
    v39 := GenS4{F0: (v2 * v2), F1: 208}
    for fi40 := 0; fi40 < 12; fi40++ {
        for fi41 := 0; fi41 < 5; fi41++ {
            v38 = (0.0 >= 0.0)
            _ = fi41
        }
        _ = fi40
    }
    var v42 uint64 = 820
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v38
    _ = v39
    _ = v42
}
