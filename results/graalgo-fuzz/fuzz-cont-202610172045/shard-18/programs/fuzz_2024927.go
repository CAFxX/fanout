package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint32 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint32, p1 int) int8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        switch (p1 << ((919) & 63)) {
        case -1:
            var v7 int16 = 1000
            _ = v7
        case 32768:
            switch (p1 - (^(p1))) {
            case 4294967295:
                if false {
                    switch "hello" {
                    case "hello":
                        v8 := GenS1{F0: fi6, F1: "foo bar", F2: (^(p0))}
                        _ = v8
                    default:
                        p1 -= (fi6 & p1)
                    }
                    v9 := []int{p1, (fi6 | fi6)}
                    _ = v9
                } else {
                    switch (p1 % ((fi6) | 1)) {
                    case -1:
                        if (!false) {
                            p0 = p0
                        }
                    case 9223372036854775807:
                        p1 -= -9223372036854775808
                    }
                    v10 := GenS1{F0: (-(fi6)), F1: "", F2: p0}
                    p0 += p0
                    _ = v10
                }
            default:
                p1 = fi6
            }
        case 957:
            var v11 int16 = 42
            _ = v11
        default:
            switch "a" {
            case "a":
                if false {
                    switch (p1 ^ fi6) {
                    case 271:
                        var v12 int8 = 42
                        _ = v12
                    case -9223372036854775808:
                        var v13 int8 = 127
                        _ = v13
                    default:
                        if true {
                            v14 := GenS1{F0: 828, F1: (("foo bar" + "") + "hello"), F2: 65535}
                            v15 := GenS1{F0: fi6, F1: "\n\t", F2: 231}
                            switch p1 {
                            case 1:
                                switch (fi6 << ((0) & 63)) {
                                case -9223372036854775808:
                                    p0 -= (p0 & (p0 * p0))
                                case 128:
                                    if (131 == 1) {
                                        switch "a" {
                                        case "":
                                            if ((float32((((float64(0.0) - 0.1) / 1.5) + 0.1)) / 1.5) > ((1.5 - (((1.5 * float32(1.5)) - 3.14159) * 3.14159)) + ((1e10 + 1.5) * 0.1))) {
                                                if false {
                                                    v16 := GenS2{F0: 190, F1: 804}
                                                    var v17 uint8 = 255
                                                    _ = v16
                                                    _ = v17
                                                }
                                                v18 := []int{(p1 & int(p0)), p1}
                                                v18 = append(v18, (fi6 * 1))
                                                _ = v18
                                            } else {
                                                fmt.Println(p0)
                                                if true {
                                                    v19 := []int{0, -715}
                                                    _ = v19
                                                } else {
                                                    p0 += p0
                                                }
                                            }
                                        case "hello":
                                            v20 := GenS1{F0: (^(fi6)), F1: "hello", F2: (p0 ^ (p0 * p0))}
                                            _ = v20
                                        case "foo bar":
                                            p0 -= 1
                                        default:
                                            p1 = fi6
                                        }
                                        p1 = (p1 >> ((664) & 63))
                                        p1 -= 32767
                                    } else {
                                        p0 = (^(p0))
                                    }
                                default:
                                    v21 := []int{}
                                    _ = v21
                                }
                            case 0:
                                v22 := []int{(-(fi6))}
                                _ = v22
                            }
                            _ = v14
                            _ = v15
                        }
                    }
                    p1 -= fi6
                }
            case "hello":
                switch 42 {
                case 1:
                    v23 := []int{(-(fi6)), p1, fi6}
                    _ = v23
                case 0:
                    v24 := GenS1{F0: p1, F1: "a", F2: (^(p0))}
                    _ = v24
                }
            case "foo bar":
                p0 = 820
            }
        }
        v25 := []int{}
        v25 = append(v25, 641)
        _ = fi6
        _ = v25
    }
    v26 := GenS2{F0: 518, F1: -32768}
    _ = p0
    _ = p1
    _ = v26
    return 0
}
func genFn1(p0 uint8, p1 string, p2 int) uint8 {
    v27 := GenS3{F0: false, F1: 0, F2: 865}
    var v28 float64 = float64((float32(0.0) - 1e10))
    var v29 float32 = 1.5
    var v30 int16 = -728
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    _ = v29
    _ = v30
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -1
    v3 := []uint32{}
    v4 := make(map[uint32]string)
    var v5 int32 = 42
    v3 = append(v3, 407)
    v2 = v2
    var v31 float64 = 3.14159
    fmt.Println(v5, v31, v2)
    v32 := GenS2{F0: 257, F1: 778}
    var v33 uint8 = 42
    fmt.Println(v5, v33, v31)
    v34 := GenS4{F0: 0}
    var v35 uint16 = 65535
    v5 = (-(v5))
    if false {
        fmt.Println(v35)
        var v36 uint = 65535
        var v37 uint8 = genFn1(v33, (("a" + "foo bar") + (("foo bar" + ("hello" + "a")) + "foo bar")), -9223372036854775808)
        _ = v36
        _ = v37
    } else {
        v35 += (v35 ^ (v35 + v35))
    }
    fmt.Println(v5, v33, v33)
    var v38 int = 4294967295
    var v39 int = v38
    var v40 bool = ((v35 + (^(v35))) != 2)
    for fi41 := 0; fi41 < 16; fi41++ {
        for fi42 := 0; fi42 < 6; fi42++ {
            v38 = v39
            var v43 uint32 = 2147483648
            v31 += v31
            _ = fi42
            _ = v43
        }
        _ = fi41
    }
    v44 := GenS5{F0: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v5
}
