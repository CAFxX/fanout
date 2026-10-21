package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 int32 }

func main() {
    var v1 int16 = -439
    var v2 float32 = (1.5 + (float32(0.1) - (float32(-2.25) + 3.14159)))
    v3 := []string{"hello", "hello", "\n\t", "foo bar"}
    v4 := make(map[int]uint32)
    var v5 uint32 = 2
    var v6 int32 = -1
    v7 := GenS1{F0: 23}
    if (true && (255 > 189)) {
        v8 := GenS1{F0: 0}
        fmt.Println(v2, v5, v6)
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 string = ("hello" + "hello")
            if false {
                v1 -= v1
                v4[int(v1)] = (v5 % (((v5 * (v5 & v5))) | 1))
                for fi11 := 0; fi11 < 10; fi11++ {
                    v10 = "\n\t"
                    _ = fi11
                }
            }
            var v12 int16 = (v1 | (v1 % ((v1) | 1)))
            _ = fi9
            _ = v10
            _ = v12
        }
        _ = v8
    } else {
        fmt.Println(v2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v5, v6)
    v13 := GenS1{F0: 4294967295}
    switch "foo bar" {
    case "hello":
        v2 = (v2 / v2)
    case "a":
        fmt.Println(v5, v1, v6)
    case "foo bar":
        switch "foo bar" {
        case "a":
            if false {
                if false {
                    v1 += -520
                }
                if (232 != 817) {
                    var v14 uint = 42
                    var v15 int16 = v1
                    if false {
                        fmt.Println(v15)
                        for fi16 := 0; fi16 < 2; fi16++ {
                            v5 = 42
                            var v17 uint = (v14 & (^(v14)))
                            for fi18 := 0; fi18 < 11; fi18++ {
                                if true {
                                    v6 = (v6 ^ (v6 ^ (v6 << (((v17 * v14)) & 63))))
                                }
                                var v19 float32 = v2
                                var v20 int = (-(fi16))
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            _ = fi16
                            _ = v17
                        }
                        var v21 int8 = 42
                        _ = v21
                    }
                    _ = v14
                    _ = v15
                }
                v22 := GenS2{F0: -128}
                _ = v22
            } else {
                var v23 int = -638
                v6 -= -1
                v24 := GenS1{F0: 195}
                _ = v23
                _ = v24
            }
        case "hello":
            if (!false) {
                var v25 int32 = (-(v6))
                v6 -= v6
                v26 := GenS3{F0: (v1 >> ((4294967295) & 63)), F1: 0}
                _ = v25
                _ = v26
            } else {
                if false {
                    if ((!(("foo bar" + "a") <= "foo bar")) || false) {
                        v27 := GenS4{F0: 32768, F1: 2}
                        var v28 uint = 154
                        var v29 bool = true
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                }
                fmt.Println(v1)
                v30 := GenS5{F0: (v6 - v6)}
                _ = v30
            }
        }
    default:
        var v31 int = -9223372036854775808
        _ = v31
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        v33 := GenS6{F0: true, F1: v6}
        if (1 == fi32) {
            var v34 string = ("foo bar" + "\n\t")
            var v35 int8 = -128
            _ = v34
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    switch "hello" {
    case "hello":
        v6 = 0
    case "a":
        switch 128 {
        case -2147483648:
            v4[624] = 4294967295
        default:
            v5 = (v5 * v5)
        }
    }
    var v36 int32 = v6
    if (!(1000 < v1)) {
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 uint16 = 1
            v2 -= (v2 * ((v2 + 1e10) / float32(1.5)))
            if false {
                v1 = (v1 << ((uint(v38)) & 63))
                if true {
                    v36 = (v6 + v6)
                    _ = len(v3)
                }
                var v39 uint = 494
                _ = v39
            }
            _ = fi37
            _ = v38
        }
        v5 += (v5 * 42)
        var v40 int8 = 42
        _ = v40
    } else {
        v41 := GenS1{F0: 251}
        for fi42 := 0; fi42 < 7; fi42++ {
            var v43 int = fi42
            _ = fi42
            _ = v43
        }
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
