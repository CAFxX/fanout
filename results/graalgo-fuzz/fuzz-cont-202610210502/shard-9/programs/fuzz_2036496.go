package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int16; F2 int16 }
func genFn0(p0 uint, p1 int16) uint8 {
    var v7 float64 = (float64(3.14159) * float64(1e10))
    var v8 int32 = -1
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            p0 += (p0 / ((4294967295) | 1))
            var v11 int8 = -127
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := []int{-2147483648, 128}
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS1{F0: 752}
        if (!((v7 - 3.14159) > 1.5)) {
            v8 = (v8 ^ (v8 - (^(v8))))
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 int32 = v8
                fmt.Println(fi13)
                v16 -= v16
                _ = fi15
                _ = v16
            }
        } else {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 uint64 = 42
                var v19 int32 = (v8 * v8)
                p0 += 86
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 uint32 = 65535
            if len(v12) > 0 {
                v12[0] = fi13
            }
            _ = fi20
            _ = v21
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    _ = v8
    return 237
}

func main() {
    var v1 int16 = 32767
    var v2 int8 = -128
    v3 := make(map[string]int)
    var v4 float32 = 1e10
    var v5 bool = (65535 > 601)
    v6 := []int{}
    var v22 uint8 = genFn0(16, (v1 - (v1 % ((965) | 1))))
    fmt.Println(v5, v5)
    var v23 string = "\n\t"
    switch v23 {
    case "hello":
        for fi24 := 0; fi24 < 1; fi24++ {
            if (2147483647 <= 42) {
                v1 = (-(v1))
                if (1.5 != 0.1) {
                    v23 = v23
                    for fi25 := 0; fi25 < 4; fi25++ {
                        if v5 {
                            if len(v6) > 0 {
                                v6[0] = (fi24 % ((fi25) | 1))
                            }
                            if v5 {
                                if (v5 || false) {
                                    fmt.Println(fi24, v2)
                                }
                                if v5 {
                                    fmt.Println(v23, v1)
                                    switch 32767 {
                                    case 1:
                                        v26 := GenS2{F0: v23, F1: int16(v22), F2: -1}
                                        _ = v26
                                    }
                                    var v27 uint16 = 531
                                    _ = v27
                                } else {
                                    var v28 float32 = 1.5
                                    fmt.Println(v4, v2, fi25)
                                    _ = v28
                                }
                            } else {
                                if v5 {
                                    var v29 string = "\n\t"
                                    _ = v29
                                }
                                v30 := v3[(v23 + v23)]
                                fmt.Println(v22, v5)
                                switch (fi24 - fi24) {
                                case 4294967295:
                                    v3[v23] = fi24
                                case 0:
                                    fmt.Println(v5, fi25)
                                case -1:
                                    v23 = ("hello" + v23)
                                default:
                                    if true {
                                        v23 = v23
                                    } else {
                                        v22 -= (v22 + 1)
                                        fmt.Println(fi24, v23)
                                        v23 = (v23 + v23)
                                    }
                                }
                                _ = v30
                            }
                            v31 := v3[v23]
                            _ = v31
                        }
                        fmt.Println(v1, v23, v5)
                        _ = fi25
                    }
                } else {
                    v2 += v2
                }
            } else {
                if (672 == 2147483647) {
                    fmt.Println(v4, v23)
                    v3[("foo bar" + (v23 + ("a" + v23)))] = fi24
                } else {
                    var v32 int8 = (v2 / ((v2) | 1))
                    _ = v32
                }
                fmt.Println(v5)
            }
            _ = fi24
        }
    }
    var v33 uint16 = 1
    v34 := GenS1{F0: 0}
    fmt.Println(v4, v23)
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    switch v23 {
    case "a":
        v35 := GenS1{F0: -9223372036854775808}
        _ = v35
    case "foo bar":
        v23 = (v23 + "hello")
    default:
        _ = len(v6)
    }
    var v36 int16 = 42
    v37 := GenS1{F0: 602}
    var v38 string = ((v23 + ("foo bar" + v23)) + "hello")
    v33 -= v33
    if v5 {
        v39 := GenS2{F0: v23, F1: v36, F2: v36}
        v22 -= uint8(v33)
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
