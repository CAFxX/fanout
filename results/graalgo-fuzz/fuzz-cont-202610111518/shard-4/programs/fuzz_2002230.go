package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint; F2 int8 }
type GenS4 struct { F0 bool; F1 int16; F2 uint8 }
type GenS5 struct { F0 bool; F1 int64; F2 uint }
func genFn0(p0 uint16) uint32 {
    v7 := GenS1{F0: true, F1: 42}
    v8 := []int{339, 0}
    fmt.Println(p0)
    p0 += p0
    if false {
        for fi9 := 0; fi9 < 15; fi9++ {
            p0 -= p0
            _ = fi9
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = 1e10
    var v3 int32 = 1
    v4 := make(map[uint32]string)
    var v5 uint64 = 18446744073709551615
    v6 := []int{-1, 42, -2147483648}
    if (38 == 0) {
        for fi10 := 0; fi10 < 13; fi10++ {
            delete(v4, 65535)
            _ = len(v6)
            _ = fi10
        }
        var v11 uint16 = 986
        var v12 int16 = 32767
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 0}
    for fi14 := 0; fi14 < 18; fi14++ {
        if (false || (true || false)) {
            var v15 int = 1
            fmt.Println(v2, v2)
            _ = v15
        }
        if (!((false && true) || (2147483647 > v3))) {
            v2 = v2
            fmt.Println(v2, fi14, v1)
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := GenS3{F0: (!(!(!((genFn0(32767) > 0) && (1 >= 1))))), F1: 654, F2: 73}
                switch (fi16 / (((fi14 | (fi14 % (((-(fi16))) | 1)))) | 1)) {
                case 2147483647:
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        if (!true) {
                            v18 := GenS1{F0: true, F1: (^(fi14))}
                            var v19 int = (fi14 * fi16)
                            v20 := GenS2{F0: (v3 * 1)}
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    case "hello":
                        if (!false) {
                            v5 = 42
                        }
                    case "\n\t":
                        v1 -= v1
                    default:
                        if false {
                            var v21 int = 0
                            var v22 float32 = -2.25
                            v23 := GenS1{F0: ((!((v5 * (v5 * 2)) >= 1)) || true), F1: -1}
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            switch fi14 {
                            case -665:
                                v24 := GenS4{F0: false, F1: 42, F2: 127}
                                _ = v24
                            default:
                                var v25 int16 = -407
                                _ = v25
                            }
                            if (32767 == 0) {
                                if false {
                                    v1 -= v1
                                    if true {
                                        switch 4294967295 {
                                        case 32767:
                                            _ = len(v6)
                                        }
                                        v26 := v4[588]
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v26
                                    } else {
                                        v1 = v1
                                        v27 := GenS5{F0: (-1 > 109), F1: -9223372036854775808, F2: 1}
                                        _ = v27
                                    }
                                    switch (fi16 & (fi16 % ((4294967295) | 1))) {
                                    case -2147483648:
                                        v3 -= v3
                                    }
                                }
                                switch "a" {
                                case "hello":
                                    var v28 int64 = v1
                                    _ = v28
                                case "foo bar":
                                    switch ("a" + ("hello" + "hello")) {
                                    case "hello":
                                        var v29 uint8 = 1
                                        _ = v29
                                    case "foo bar":
                                        if ((1.5 - 1e10) == 1.5) {
                                            var v30 uint32 = 385
                                            var v31 int64 = (v1 & v1)
                                            delete(v4, v30)
                                            var v32 int64 = v1
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        }
                                    default:
                                        fmt.Println(v1, fi16)
                                    }
                                }
                            }
                        }
                    }
                case 831:
                    var v33 uint8 = 1
                    var v34 int8 = 100
                    _ = v33
                    _ = v34
                default:
                    var v35 int = (fi14 ^ 9223372036854775807)
                    _ = v35
                }
                var v36 uint16 = 32767
                _ = fi16
                _ = v17
                _ = v36
            }
        }
        _ = fi14
    }
    v1 = 9223372036854775807
    var v37 int = 127
    var v38 uint8 = 0
    v5 += (v5 & (^(v5)))
    for fi39 := 0; fi39 < 8; fi39++ {
        for fi40 := 0; fi40 < 7; fi40++ {
            if false {
                var v41 uint8 = v38
                v4[2147483647] = (("hello" + "\n\t") + (("hello" + "a") + ((("\n\t" + "foo bar") + "foo bar") + "\n\t")))
                _ = v41
            } else {
                if (2 <= 18) {
                    if (!true) {
                        var v42 int64 = (v1 % (((v1 & (-(v1)))) | 1))
                        var v43 float32 = float32((-2.25 / float64(float32(float64(1e10)))))
                        var v44 bool = ((v38 >> ((18446744073709551615) & 63)) < v38)
                        _ = v42
                        _ = v43
                        _ = v44
                    } else {
                        switch "hello" {
                        case "":
                            v6 = append(v6, (fi40 << ((799) & 63)))
                        }
                        v5 = v5
                        var v45 bool = (float64(float32(-2.25)) >= (((float64(3.14159) * (v2 - 0.1)) * float64(float32(float64(float32(v2))))) / (v2 + v2)))
                        _ = v45
                    }
                    _ = len(v4)
                } else {
                    fmt.Println(v2, v1)
                    v2 -= v2
                    v5 -= (^(v5))
                }
                v38 -= (^(v38))
                v3 += 2147483647
                v38 = (v38 - (v38 & 115))
            }
            _ = fi40
        }
        var v46 int = (v37 * fi39)
        _ = fi39
        _ = v46
    }
    v47 := v4[genFn0(42)]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
