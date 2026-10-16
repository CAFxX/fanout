package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0(p0 int, p1 uint32) int16 {
    p1 = (p1 & (p1 / ((p1) | 1)))
    v5 := GenS1{F0: 42, F1: 18446744073709551615, F2: 1}
    p0 -= p0
    _ = p0
    _ = p1
    _ = v5
    return 656
}
func genFn1(p0 string, p1 string) uint16 {
    if true {
        p0 = p1
        p1 = p0
    }
    p0 = p1
    for fi6 := 0; fi6 < 3; fi6++ {
        if (!false) {
            fmt.Println(fi6, p1)
        } else {
            if ((float64(1e10) > 1e10) && (!false)) {
                if (1.5 != -2.25) {
                    for fi7 := 0; fi7 < 4; fi7++ {
                        if (830 >= 1) {
                            v8 := []int{fi6}
                            var v9 bool = (-2.25 > (((0.0 / 1.5) + 3.14159) / 1.0))
                            _ = v8
                            _ = v9
                        } else {
                            if (true || (65535 <= 953)) {
                                v10 := []int{}
                                v11 := GenS2{F0: p0, F1: 1000}
                                var v12 uint = 255
                                _ = v10
                                _ = v11
                                _ = v12
                            } else {
                                switch 32767 {
                                case 4294967295:
                                    v13 := GenS1{F0: 0, F1: 0, F2: 234}
                                    _ = v13
                                default:
                                    fmt.Println(fi7)
                                }
                                if (!true) {
                                    fmt.Println(p1, p1, fi6)
                                    switch p0 {
                                    case "hello":
                                        v14 := GenS1{F0: 4294967295, F1: 4294967295, F2: 349}
                                        _ = v14
                                    }
                                } else {
                                    fmt.Println(fi6)
                                }
                                v15 := GenS3{F0: (!true), F1: -32768}
                                _ = v15
                            }
                        }
                        fmt.Println(p0)
                        _ = fi7
                    }
                } else {
                    p0 = p1
                }
                v16 := []int{(fi6 | -9223372036854775808), 392, 42}
                _ = v16
            }
            var v17 int = (fi6 ^ -2147483648)
            _ = v17
        }
        _ = fi6
    }
    v18 := GenS4{F0: -1, F1: 4294967295}
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := []int{2147483647, (fi19 - (-(fi19))), (fi19 << ((993) & 63))}
        v21 := GenS4{F0: genFn0((-(fi19)), 228), F1: 4294967295}
        _ = fi19
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v18
    return 606
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]int)
    v3 += v3
    for fi22 := 0; fi22 < 16; fi22++ {
        switch ("" + (((("foo bar" + "a") + "") + "foo bar") + "a")) {
        case "\n\t":
            v23 := []int{fi22, -9223372036854775808, fi22, 188}
            _ = v23
        case "hello":
            v24 := GenS3{F0: true, F1: 1000}
            _ = v24
        }
        _ = fi22
    }
    _ = len(v2)
    v25 := GenS2{F0: "hello", F1: 1}
    v3 -= 2147483647
    switch (((("\n\t" + "foo bar") + ((("foo bar" + "foo bar") + "a") + (("" + "") + "foo bar"))) + "hello") + ("" + (("" + "a") + "\n\t"))) {
    case "\n\t":
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 int8 = -11
            var v28 int16 = 147
            if (9223372036854775807 == 0) {
                for fi29 := 0; fi29 < 20; fi29++ {
                    if true {
                        if true {
                            fmt.Println(fi29, v27, v28)
                            var v30 float32 = 0.0
                            _ = len(v2)
                            _ = v30
                        }
                    }
                    v31 := GenS3{F0: (-2.25 >= 0.0), F1: (v28 & -32768)}
                    var v32 int16 = v28
                    _ = fi29
                    _ = v31
                    _ = v32
                }
            }
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    var v33 float64 = 0.1
    v33 += float64(0.1)
    v33 -= ((((((v33 + ((v33 + float64(3.14159)) / (v33 * (v33 + v33)))) - v33) / v33) - float64(0.0)) - ((v33 * (v33 / (((v33 * -2.25) + 3.14159) * 3.14159))) + -2.25)) - v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v4
}
