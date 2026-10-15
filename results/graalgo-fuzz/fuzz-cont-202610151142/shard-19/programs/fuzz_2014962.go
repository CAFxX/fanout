package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 bool; F1 int; F2 uint8 }
func genFn0(p0 uint, p1 uint8, p2 int32) uint64 {
    var v4 uint = 166
    p1 += p1
    if (("foo bar" + "foo bar") == (("a" + ("\n\t" + "a")) + "\n\t")) {
        switch 867 {
        case 32768:
            v4 -= (v4 + 549)
        case 9223372036854775807:
            if true {
                v5 := GenS1{F0: ("foo bar" + (("\n\t" + "foo bar") + "hello"))}
                _ = v5
            }
        default:
            fmt.Println(p1, v4, p1)
        }
    } else {
        if false {
            fmt.Println(v4)
            var v6 float64 = (0.0 - (0.1 / 1.0))
            _ = v6
        }
        fmt.Println(v4, p2)
    }
    var v7 uint16 = 255
    fmt.Println(p2, p2, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 4294967295
}

func main() {
    var v1 int32 = 42
    var v2 uint = 237
    var v3 int32 = 1
    v8 := []int{9223372036854775807, -307}
    v9 := GenS2{F0: "\n\t", F1: -1}
    var v10 float32 = -2.25
    if (2 == 1) {
        v11 := GenS3{F0: 42}
        _ = v11
    } else {
        var v12 uint8 = 127
        if (!true) {
            var v13 uint8 = 2
            var v14 float32 = 0.0
            _ = v13
            _ = v14
        } else {
            v8 = append(v8, 1)
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            v15 := GenS4{F0: -1, F1: 1}
            _ = v15
        }
        _ = v12
    }
    _ = len(v8)
    if ((!true) || false) {
        for fi16 := 0; fi16 < 1; fi16++ {
            v2 = (v2 ^ v2)
            v17 := GenS1{F0: "a"}
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi16
            _ = v17
        }
        v1 += v3
        v3 = 380
    }
    if (genFn0((v2 & v2), 128, v3) >= 1) {
        switch "" {
        case "a":
            switch -975 {
            case -38:
                switch "" {
                case "":
                    v18 := GenS3{F0: 42}
                    _ = v18
                }
            case 2147483647:
                if ("\n\t" >= "foo bar") {
                    v3 = -2147483648
                } else {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        for fi20 := 0; fi20 < 18; fi20++ {
                            var v21 uint32 = 316
                            if true {
                                var v22 int16 = 539
                                _ = v22
                            }
                            _ = fi20
                            _ = v21
                        }
                        v23 := GenS2{F0: "hello", F1: (-(v1))}
                        v1 -= v3
                        _ = fi19
                        _ = v23
                    }
                    if (!false) {
                        fmt.Println(v1, v10)
                    }
                    if true {
                        if len(v8) > 0 {
                            v8[0] = 1
                        }
                    }
                }
            }
        default:
            var v24 int32 = v1
            _ = v24
        }
        fmt.Println(v3, v10, v10)
        v2 += (^(v2))
    }
    v8 = append(v8, 326)
    var v25 uint32 = 65535
    for fi26 := 0; fi26 < 14; fi26++ {
        switch "" {
        case "hello":
            var v27 uint = (v2 / (((^(v2))) | 1))
            _ = v27
        }
        _ = fi26
    }
    switch -20 {
    case 127:
        v8 = append(v8, -1)
    case 1:
        v28 := GenS5{F0: (!false), F1: 32767, F2: 100}
        _ = v28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        v30 := GenS3{F0: 2}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        var v31 bool = (true && (0 > -492))
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int8 = 58
    for fi33 := 0; fi33 < 13; fi33++ {
        v34 := GenS5{F0: false, F1: (^(fi33)), F2: 1}
        var v35 float32 = v10
        v32 += (^(v32))
        _ = fi33
        _ = v34
        _ = v35
    }
    if len(v8) > 0 {
        v8[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v8
    _ = v9
}
