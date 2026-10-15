package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 bool; F1 int8; F2 int32 }
func genFn0() uint64 {
    var v5 bool = true
    var v6 int64 = -9223372036854775808
    fmt.Println(v5)
    switch -2147483648 {
    case -193:
        var v7 int8 = -1
        _ = v7
    case -9223372036854775808:
        var v8 uint16 = 1
        _ = v8
    case 127:
        switch "a" {
        case "hello":
            var v9 float32 = (0.0 * 0.0)
            _ = v9
        case "\n\t":
            if v5 {
                for fi10 := 0; fi10 < 16; fi10++ {
                    v11 := GenS1{F0: 255, F1: 2}
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: 177, F1: 42}
                _ = v12
            }
        case "foo bar":
            if v5 {
                v6 = (v6 & -1)
                v13 := GenS1{F0: 128, F1: 0}
                _ = v13
            } else {
                for fi14 := 0; fi14 < 16; fi14++ {
                    fmt.Println(fi14, fi14)
                    _ = fi14
                }
                for fi15 := 0; fi15 < 7; fi15++ {
                    switch 127 {
                    case 2147483647:
                        fmt.Println(v6)
                    case 510:
                        if true {
                            for fi16 := 0; fi16 < 6; fi16++ {
                                if v5 {
                                    var v17 uint = 815
                                    var v18 int = fi16
                                    _ = v17
                                    _ = v18
                                } else {
                                    fmt.Println(fi15, v5, fi16)
                                    v19 := GenS1{F0: 74, F1: 18446744073709551615}
                                    _ = v19
                                }
                                switch 128 {
                                case 9223372036854775807:
                                    switch -2147483648 {
                                    case 834:
                                        switch (("a" + ("hello" + (("\n\t" + "foo bar") + "hello"))) + ("foo bar" + ("a" + "\n\t"))) {
                                        case "\n\t":
                                            var v20 int32 = 1
                                            _ = v20
                                        case "":
                                            fmt.Println(v6)
                                        case "hello":
                                            fmt.Println(v6)
                                        default:
                                            if true {
                                                v21 := []int{127}
                                                _ = v21
                                            } else {
                                                fmt.Println(v6)
                                            }
                                        }
                                    default:
                                        if false {
                                            v6 -= -1
                                            var v22 int32 = 0
                                            switch "\n\t" {
                                            case "a":
                                                var v23 uint64 = 1
                                                v24 := []int{}
                                                _ = v23
                                                _ = v24
                                            }
                                            _ = v22
                                        }
                                    }
                                case 42:
                                    var v25 float64 = ((1.5 * 1.5) * float64(3.14159))
                                    _ = v25
                                case 1:
                                    fmt.Println(fi16, fi16, fi15)
                                default:
                                    if (36 >= 85) {
                                        v26 := []int{fi16, 1, fi16}
                                        fmt.Println(fi15, fi16, v5)
                                        fmt.Println(fi15, v5, v6)
                                        _ = v26
                                    }
                                }
                                if (!(((2 >= 255) || (v5 && (19 <= 18446744073709551615))) || v5)) {
                                    v27 := []int{(fi16 << ((0) & 63)), fi16}
                                    _ = v27
                                }
                                _ = fi16
                            }
                        }
                    case -508:
                        switch -88 {
                        case 0:
                            v28 := GenS1{F0: 5, F1: 307}
                            _ = v28
                        default:
                            fmt.Println(v6)
                        }
                    default:
                        v6 = (v6 ^ v6)
                    }
                    v6 += (v6 % ((v6) | 1))
                    _ = fi15
                }
            }
        default:
            v5 = (("a" + "hello") == ("foo bar" + "a"))
        }
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 uint8 = 0
        _ = fi29
        _ = v30
    }
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint8) uint8 {
    v31 := []int{-16, 4294967295, -1}
    v32 := GenS1{F0: p0, F1: 476}
    for fi33 := 0; fi33 < 2; fi33++ {
        p0 -= (p0 >> ((1) & 63))
        _ = fi33
    }
    fmt.Println(p0)
    _ = p0
    _ = v31
    _ = v32
    return p0
}

func main() {
    var v1 uint32 = 197
    var v2 int8 = 100
    v3 := []uint32{(v1 / (((v1 + 0)) | 1))}
    v4 := make(map[int]bool)
    for fi34 := 0; fi34 < 7; fi34++ {
        delete(v4, (fi34 | 128))
        var v35 uint8 = 128
        _ = fi34
        _ = v35
    }
    v1 += (v1 / ((2147483647) | 1))
    if (false || ((v2 * -128) == (v2 * v2))) {
        if (!(!(!false))) {
            var v36 uint8 = 0
            _ = v36
        }
    } else {
        v37 := GenS1{F0: 2, F1: 471}
        _ = v37
    }
    v1 = v1
    v2 = v2
    var v38 float32 = (-2.25 * 3.14159)
    v1 += (^(v1))
    var v39 int64 = -9223372036854775808
    v2 -= (v2 ^ (^(v2)))
    fmt.Println(v39)
    var v40 int = -2147483648
    v1 = 2147483647
    v4[-773] = (!(v39 == -1))
    v41 := GenS2{F0: v40, F1: ("hello" + "foo bar")}
    v42 := GenS3{F0: true, F1: (v2 % ((-16) | 1)), F2: -360}
    v43 := GenS3{F0: (false || false), F1: (v2 << ((0) & 63)), F2: -1}
    var v44 uint32 = 653
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
