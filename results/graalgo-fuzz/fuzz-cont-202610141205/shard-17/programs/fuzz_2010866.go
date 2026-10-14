package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint, p1 bool, p2 string) uint8 {
    v6 := []int{2147483647, -44, 674}
    p0 = p0
    for fi7 := 0; fi7 < 3; fi7++ {
        v6 = append(v6, (fi7 + 127))
        var v8 bool = ((p2 + ("\n\t" + ((((((p2 + "") + "") + "a") + (("hello" + (p2 + p2)) + p2)) + p2) + "hello"))) > p2)
        var v9 uint = p0
        _ = fi7
        _ = v8
        _ = v9
    }
    if (!p1) {
        for fi10 := 0; fi10 < 20; fi10++ {
            if (p1 || true) {
                p2 = p2
                if p1 {
                    var v11 uint32 = 42
                    _ = v11
                } else {
                    var v12 uint16 = 335
                    _ = v12
                }
            }
            var v13 bool = (false && p1)
            switch "a" {
            case "foo bar":
                p0 -= (p0 + p0)
            case "":
                var v14 int64 = 1
                _ = v14
            case "hello":
                var v15 string = ("hello" + p2)
                _ = v15
            }
            _ = fi10
            _ = v13
        }
    } else {
        fmt.Println(p0)
        p1 = p1
        if (0.1 <= 1e10) {
            var v16 bool = (0 != 18446744073709551615)
            for fi17 := 0; fi17 < 2; fi17++ {
                switch ("foo bar" + ((p2 + "hello") + p2)) {
                case "foo bar":
                    if ((!(42 > -9223372036854775808)) || (false || p1)) {
                        var v18 int8 = 42
                        if ((^(p0)) != (p0 / ((p0) | 1))) {
                            if ((0.1 < 1e10) && p1) {
                                var v19 string = (p2 + (("\n\t" + p2) + (p2 + "\n\t")))
                                _ = v19
                            } else {
                                if (-1 == -2147483648) {
                                    p2 = (("" + "foo bar") + p2)
                                } else {
                                    p1 = (18446744073709551615 == 42)
                                }
                                var v20 int8 = v18
                                _ = v20
                            }
                            var v21 float32 = ((0.0 + 1e10) * 0.0)
                            if len(v6) > 0 {
                                v6[0] = -9223372036854775808
                            }
                            _ = v21
                        } else {
                            var v22 int = 128
                            _ = v22
                        }
                        v23 := GenS1{F0: 1, F1: 42}
                        _ = v18
                        _ = v23
                    } else {
                        var v24 int16 = -330
                        if len(v6) > 0 {
                            v6[0] = 32767
                        }
                        fmt.Println(p2, fi17, p1)
                        _ = v24
                    }
                }
                switch fi17 {
                case 569:
                    v25 := GenS2{F0: 656}
                    _ = v25
                default:
                    var v26 uint64 = 545
                    _ = v26
                }
                _ = fi17
            }
            _ = v16
        } else {
            switch 626 {
            case 4294967295:
                p1 = ((3.14159 / float32(1e10)) >= (-2.25 / (-2.25 / 0.1)))
            case -9223372036854775808:
                p0 = p0
            default:
                fmt.Println(p2, p1, p0)
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 uint32) string {
    for fi27 := 0; fi27 < 16; fi27++ {
        p0 -= 454
        for fi28 := 0; fi28 < 11; fi28++ {
            p0 = p0
            p0 -= p0
            v29 := []int{0}
            var v30 uint = 1
            if (true || false) {
                var v31 uint = (v30 >> ((v30) & 63))
                _ = len(v29)
                p0 += p0
                _ = v31
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        for fi32 := 0; fi32 < 13; fi32++ {
            v33 := []int{}
            _ = fi32
            _ = v33
        }
        _ = fi27
    }
    v34 := []int{983}
    _ = p0
    _ = v34
    return "\n\t"
}

func main() {
    var v1 float64 = (0.1 * float64(3.14159))
    var v2 int32 = 0
    v3 := []uint32{621, 42}
    var v4 uint16 = 42
    var v5 uint = 2
    v1 += (v1 + v1)
    var v35 float64 = (v1 / v1)
    if (!false) {
        v3 = append(v3, 42)
        var v36 uint16 = (v4 * v4)
        _ = v36
    } else {
        var v37 int16 = 0
        _ = v37
    }
    if len(v3) > 0 {
        v3[0] = 2
    }
    var v38 int8 = 42
    var v39 uint8 = 255
    for fi40 := 0; fi40 < 14; fi40++ {
        v41 := GenS2{F0: v2}
        _ = fi40
        _ = v41
    }
    var v42 int16 = 42
    if true {
        v42 -= v42
        v39 -= v39
    }
    v43 := GenS3{F0: v38, F1: (1 != 2147483647)}
    var v44 int8 = v38
    v4 -= (v4 % ((v4) | 1))
    v45 := GenS4{F0: v42}
    var v46 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
