package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint32; F1 int32; F2 int8 }
type GenS4 struct { F0 int32; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int64; F1 int64 }
func genFn0(p0 bool) int16 {
    switch "hello" {
    case "foo bar":
        p0 = true
    case "a":
        fmt.Println(p0)
    case "hello":
        for fi4 := 0; fi4 < 16; fi4++ {
            if (false || p0) {
                p0 = false
            }
            v5 := []int{fi4}
            _ = fi4
            _ = v5
        }
    }
    var v6 uint = 224
    if (307 != 548) {
        switch "foo bar" {
        case "":
            v7 := []int{-576, 32768, -1, 32768}
            _ = v7
        case "a":
            v8 := GenS1{F0: -1, F1: 494}
            _ = v8
        default:
            if (42 == 0) {
                if (!(false && (!((((!(((2 != 1) || p0) && p0)) && false) || p0) && (550 > -32768))))) {
                    v9 := []int{}
                    var v10 bool = false
                    for fi11 := 0; fi11 < 18; fi11++ {
                        var v12 int32 = 0
                        v13 := GenS2{F0: ("hello" + ("foo bar" + "\n\t")), F1: fi11}
                        if true {
                            for fi14 := 0; fi14 < 1; fi14++ {
                                switch (("\n\t" + ("foo bar" + "\n\t")) + "hello") {
                                case "a":
                                    fmt.Println(v12, v6)
                                case "\n\t":
                                    var v15 uint8 = 0
                                    _ = v15
                                }
                                v16 := GenS2{F0: ((("foo bar" + "") + ("" + ("foo bar" + ""))) + "a"), F1: -1}
                                _ = fi14
                                _ = v16
                            }
                        } else {
                            v17 := GenS3{F0: 2147483648, F1: (v12 * (v12 / ((v12) | 1))), F2: 110}
                            v12 += -688
                            v9 = append(v9, 905)
                            _ = v17
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = v10
                    _ = v9
                }
                var v18 uint = v6
                _ = v18
            }
        }
    }
    _ = p0
    _ = v6
    return -32768
}

func main() {
    v1 := []uint32{2147483647}
    var v2 int8 = -128
    v3 := make(map[string]string)
    switch "a" {
    case "":
        var v19 int = 9223372036854775807
        _ = v19
    case "foo bar":
        v2 += 100
    case "\n\t":
        v2 -= v2
    }
    v2 = v2
    if false {
        if len(v1) > 0 {
            v1[0] = 2
        }
        fmt.Println(v2)
        v2 -= v2
    }
    fmt.Println(v2)
    if false {
        var v20 int = 4294967295
        v21 := GenS1{F0: 992, F1: -32768}
        for fi22 := 0; fi22 < 9; fi22++ {
            if (65535 == 2) {
                _ = len(v3)
            } else {
                var v23 float64 = (0.1 / ((3.14159 - float64((-2.25 * -2.25))) * 1e10))
                v20 += fi22
                if ((false && ((true || (-9223372036854775808 < -923)) && false)) || (!true)) {
                    switch "a" {
                    case "hello":
                        for fi24 := 0; fi24 < 16; fi24++ {
                            v2 = (v2 & 0)
                            _ = fi24
                        }
                    }
                    v23 += ((3.14159 + 0.0) * (v23 / v23))
                }
                _ = v23
            }
            _ = fi22
        }
        _ = v20
        _ = v21
    }
    v25 := GenS4{F0: 0, F1: 4294967295, F2: 42}
    var v26 uint = 18446744073709551615
    for fi27 := 0; fi27 < 13; fi27++ {
        v2 -= (v2 * v2)
        v28 := GenS5{F0: 2, F1: "foo bar"}
        if (true && false) {
            if (false || true) {
                v26 = (v26 % ((782) | 1))
                if ((fi27 * (fi27 ^ fi27)) > 32767) {
                    v29 := GenS1{F0: -2147483648, F1: -1}
                    var v30 uint64 = 42
                    v30 = (v30 * v30)
                    _ = v29
                    _ = v30
                } else {
                    if (1 <= 32767) {
                        if len(v1) > 0 {
                            v1[0] = 944
                        }
                    }
                    v31 := GenS5{F0: 0, F1: "foo bar"}
                    for fi32 := 0; fi32 < 10; fi32++ {
                        switch fi27 {
                        case 0:
                            fmt.Println(v26, fi32, fi32)
                        case 128:
                            v33 := GenS2{F0: "a", F1: 2147483647}
                            _ = v33
                        }
                        var v34 uint64 = 2
                        v34 += v34
                        _ = fi32
                        _ = v34
                    }
                    _ = v31
                }
                v1 = append(v1, 65535)
            } else {
                var v35 uint = (v26 & (^(v26)))
                _ = v35
            }
            var v36 float64 = 1e10
            for fi37 := 0; fi37 < 17; fi37++ {
                v38 := GenS6{F0: -726, F1: -872}
                _ = fi37
                _ = v38
            }
            _ = v36
        }
        _ = fi27
        _ = v28
    }
    v2 += 88
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
