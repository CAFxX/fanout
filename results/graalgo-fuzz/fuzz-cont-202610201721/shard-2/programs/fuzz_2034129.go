package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int32 }
func genFn0(p0 int8, p1 int, p2 uint64) uint64 {
    switch p1 {
    case 32767:
        if ((false || true) || (!true)) {
            p2 = p2
        } else {
            v6 := GenS1{F0: "\n\t"}
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := []int{fi7, fi7}
                switch "\n\t" {
                case "":
                    p2 += (p2 ^ p2)
                default:
                    for fi9 := 0; fi9 < 2; fi9++ {
                        var v10 int16 = 32767
                        _ = fi9
                        _ = v10
                    }
                }
                fmt.Println(p0)
                _ = fi7
                _ = v8
            }
            for fi11 := 0; fi11 < 13; fi11++ {
                switch (("hello" + (("a" + "") + ("" + ("\n\t" + "a")))) + "hello") {
                case "":
                    for fi12 := 0; fi12 < 8; fi12++ {
                        if (!false) {
                            p0 = (p0 + p0)
                            var v13 uint16 = 0
                            p2 += p2
                            _ = v13
                        }
                        var v14 int = (p1 * 32767)
                        v15 := GenS2{F0: (true || true), F1: 562}
                        _ = fi12
                        _ = v14
                        _ = v15
                    }
                }
                _ = fi11
            }
            _ = v6
        }
    }
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 int = -1
    var v2 int16 = -32768
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]uint32)
    var v5 int32 = -1
    var v16 bool = (v5 != 2147483647)
    var v17 uint32 = 628
    v5 -= v5
    v18 := GenS2{F0: v16, F1: (-(v5))}
    if (-128 == -96) {
        var v19 int8 = -1
        for fi20 := 0; fi20 < 11; fi20++ {
            for fi21 := 0; fi21 < 8; fi21++ {
                switch (^(v1)) {
                case -9223372036854775808:
                    var v22 uint = 771
                    if (!(0.0 == 1.5)) {
                        v23 := []int{}
                        var v24 int8 = -128
                        var v25 int64 = 1
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                case -2147483648:
                    var v26 bool = (v16 || false)
                    _ = v26
                case 32768:
                    var v27 string = "foo bar"
                    _ = v27
                default:
                    v28 := []int{fi21, v1, 897, (fi20 / (((fi20 / ((v1) | 1))) | 1))}
                    _ = v28
                }
                _ = fi21
            }
            v3["a"] = v17
            v29 := v4[(v17 | v17)]
            _ = fi20
            _ = v29
        }
        var v30 bool = false
        _ = v19
        _ = v30
    }
    var v31 float64 = 1.5
    switch "a" {
    case "":
        v32 := GenS1{F0: "hello"}
        _ = v32
    default:
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 uint8 = 255
            v35 := []int{}
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    if (489 > 18446744073709551615) {
        for fi36 := 0; fi36 < 3; fi36++ {
            v37 := []int{(-(v1)), 56, 2147483647, 32768}
            _ = fi36
            _ = v37
        }
    }
    var v38 string = "\n\t"
    if v16 {
        var v39 uint16 = 32767
        v1 = (v1 + 128)
        _ = v39
    } else {
        for fi40 := 0; fi40 < 18; fi40++ {
            for fi41 := 0; fi41 < 14; fi41++ {
                v38 = (v38 + v38)
                v31 += v31
                v42 := GenS1{F0: ("foo bar" + v38)}
                _ = fi41
                _ = v42
            }
            _ = fi40
        }
    }
    fmt.Println(v16, v31, v2)
    v38 = v38
    if (true || (!(!v16))) {
        if v16 {
            v16 = ((float32(-2.25) + -2.25) == (3.14159 * (0.0 / -2.25)))
        }
    }
    if (false || v16) {
        var v43 int64 = -809
        var v44 int64 = 22
        v5 = (v5 + (v5 - v5))
        _ = v43
        _ = v44
    } else {
        var v45 int32 = v5
        _ = v45
    }
    v46 := GenS1{F0: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v4
    _ = v46
    _ = v5
}
