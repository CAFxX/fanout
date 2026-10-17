package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 string; F1 uint64; F2 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64; F1 uint16; F2 bool }
type GenS7 struct { F0 int64 }
func genFn0(p0 int32, p1 int8, p2 string) uint16 {
    p1 -= p1
    v4 := []int{32768, -1, -2147483648}
    v5 := GenS1{F0: p2}
    switch "" {
    case "":
        if false {
            for fi6 := 0; fi6 < 15; fi6++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi7 := 0; fi7 < 12; fi7++ {
                    p1 -= 42
                    _ = fi7
                }
                v8 := GenS2{F0: 374}
                _ = fi6
                _ = v8
            }
        } else {
            var v9 uint32 = 1
            fmt.Println(p2, p1, p1)
            p2 = p2
            _ = v9
        }
    case "hello":
        v10 := GenS3{F0: 818, F1: 42}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 282
}
func genFn1(p0 string, p1 int, p2 string) int8 {
    v11 := GenS1{F0: p0}
    var v12 float32 = float32((((0.1 * 0.1) * 0.1) / -2.25))
    fmt.Println(p2)
    var v13 float64 = (3.14159 * 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return -128
}

func main() {
    v1 := []int{-491}
    var v2 float64 = 0.1
    var v3 float64 = 0.0
    fmt.Println(v3, v2)
    v3 = (v2 * (0.1 * 1.5))
    var v14 int64 = -714
    v15 := GenS4{F0: "a", F1: 654, F2: 0}
    fmt.Println(v14)
    if false {
        switch (("foo bar" + "foo bar") + ("foo bar" + "hello")) {
        case "\n\t":
            for fi16 := 0; fi16 < 14; fi16++ {
                if len(v1) > 0 {
                    v1[0] = fi16
                }
                if false {
                    fmt.Println(v14)
                    var v17 int = (-(fi16))
                    _ = v17
                } else {
                    if true {
                        v14 -= 1
                        _ = len(v1)
                    }
                    for fi18 := 0; fi18 < 7; fi18++ {
                        v3 -= v3
                        _ = fi18
                    }
                }
                _ = fi16
            }
        case "":
            if ((-2147483648 <= 1) || false) {
                switch ("" + ("a" + "")) {
                case "foo bar":
                    v19 := GenS1{F0: ""}
                    _ = v19
                }
            }
        }
        var v20 int16 = 1
        _ = v20
    }
    v3 = 0.0
    v14 += (v14 + v14)
    for fi21 := 0; fi21 < 17; fi21++ {
        for fi22 := 0; fi22 < 18; fi22++ {
            switch (fi21 + -839) {
            case 32767:
                v3 = v2
            }
            v23 := GenS4{F0: ("hello" + "\n\t"), F1: 18446744073709551615, F2: -9223372036854775808}
            switch "a" {
            case "\n\t":
                fmt.Println(v2, fi22, v2)
            case "a":
                v14 += v14
            case "foo bar":
                v2 = float64(float32(3.14159))
                if true {
                    var v24 int32 = -1
                    v25 := GenS4{F0: "", F1: 2, F2: (fi22 / (((fi22 % ((-198) | 1))) | 1))}
                    fmt.Println(v24, fi22, v14)
                    if len(v1) > 0 {
                        v1[0] = (fi22 - fi22)
                    }
                    _ = v24
                    _ = v25
                } else {
                    if len(v1) > 0 {
                        v1[0] = (fi22 >> ((411) & 63))
                    }
                    fmt.Println(fi21, v3)
                    if (2147483648 != 547) {
                        var v26 int64 = (v14 ^ v14)
                        if (float64(1e10) >= v3) {
                            var v27 int = 271
                            _ = v27
                        }
                        var v28 float64 = 1.5
                        _ = v26
                        _ = v28
                    }
                }
                switch "\n\t" {
                case "a":
                    switch fi21 {
                    case 1:
                        var v29 int16 = -32768
                        v30 := GenS5{F0: 65535}
                        _ = v29
                        _ = v30
                    case 128:
                        fmt.Println(v14)
                    case 32768:
                        var v31 uint8 = 127
                        _ = v31
                    default:
                        if len(v1) > 0 {
                            v1[0] = (fi21 << ((4294967295) & 63))
                        }
                    }
                case "":
                    v14 += (v14 << ((255) & 63))
                case "foo bar":
                    v3 = (v3 - v2)
                default:
                    if false {
                        switch fi22 {
                        case 127:
                            v3 = 0.0
                        case 2147483647:
                            if (42 <= 0) {
                                if (((v14 % (((v14 % ((-915) | 1))) | 1)) < (v14 + (v14 % (((^(v14))) | 1)))) || false) {
                                    var v32 bool = false
                                    v3 += ((((v2 / 1e10) + (((1e10 * -2.25) + 0.1) - float64(float32(-2.25)))) - float64(1e10)) - -2.25)
                                    v33 := GenS6{F0: v14, F1: 498, F2: ((!(v32 || v32)) && (42 != 1))}
                                    _ = v32
                                    _ = v33
                                }
                                if (genFn0(3, -128, "\n\t") == 32767) {
                                    v34 := GenS6{F0: v14, F1: 65535, F2: (42 > 1)}
                                    _ = v34
                                }
                                v14 += (v14 - v14)
                                v3 = 0.0
                                v2 = v2
                            }
                        }
                        switch (fi22 & (fi21 / (((fi22 + int(v14))) | 1))) {
                        case 4294967295:
                            v3 += 0.0
                        default:
                            var v35 int64 = v14
                            _ = v35
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v36 := GenS1{F0: "foo bar"}
                        _ = v36
                    }
                }
                var v37 uint64 = 0
                fmt.Println(v37)
                v2 -= (v2 / v2)
                var v38 int8 = 100
                v37 = (v37 - v37)
                if false {
                    var v39 int16 = 42
                    _ = v39
                } else {
                    switch 9223372036854775807 {
                    case 4294967295:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
                _ = v37
                _ = v38
            }
            _ = fi22
            _ = v23
        }
        v3 = ((v2 - v3) / 1.0)
        _ = fi21
    }
    v14 = -749
    v14 = v14
    var v40 int32 = 1
    switch (((("hello" + "") + ((("" + "\n\t") + "\n\t") + (("hello" + ((("" + "a") + ("foo bar" + "")) + "")) + (((("\n\t" + "foo bar") + "hello") + (("foo bar" + "hello") + ("a" + "\n\t"))) + "")))) + "foo bar") + "hello") {
    case "hello":
        v3 -= (v2 / float64(float32(v2)))
    default:
        var v41 float32 = -2.25
        _ = v41
    }
    v14 = v14
    v42 := GenS7{F0: int64(v40)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v40
    _ = v42
}
