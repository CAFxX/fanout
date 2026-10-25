package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 string, p1 int) uint {
    p0 = ("hello" + p0)
    p0 = "foo bar"
    var v7 int = (p1 | -2147483648)
    p0 = ("hello" + (("\n\t" + (((((("\n\t" + ("" + "a")) + p0) + p0) + p0) + (p0 + p0)) + ("\n\t" + (p0 + "\n\t")))) + "foo bar"))
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 float32 = (float32(3.14159) / 0.1)
        if false {
            var v10 int64 = 42
            var v11 uint16 = 1
            _ = v10
            _ = v11
        }
        v12 := []int{p1, (fi8 & fi8)}
        _ = fi8
        _ = v12
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 154
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    v3 := []int{-381, -9223372036854775808, -2147483648, 42}
    var v4 uint = 255
    v5 := make(map[uint32]string)
    v6 := make(map[int]string)
    fmt.Println(v4)
    for fi13 := 0; fi13 < 4; fi13++ {
        v4 = v4
        _ = fi13
    }
    v4 += (v4 | v4)
    v1[2] = (("" + "hello") + "hello")
    v14 := v1[65535]
    var v15 int = 42
    for fi16 := 0; fi16 < 2; fi16++ {
        switch (v14 + (v14 + (("foo bar" + (v14 + v14)) + "a"))) {
        case "hello":
            for fi17 := 0; fi17 < 18; fi17++ {
                v4 = (v4 & (v4 / (((v4 / ((628) | 1))) | 1)))
                var v18 float64 = ((((1.5 + 3.14159) / 3.14159) * float64((0.0 / -2.25))) / 0.1)
                v14 = v14
                v14 = ""
                _ = fi17
                _ = v18
            }
        case "a":
            var v19 bool = false
            _ = v19
        default:
            v14 = ("a" + (("" + v14) + v14))
        }
        _ = fi16
    }
    v15 = 1
    v15 -= (v15 >> (((v4 + genFn0("hello", v15))) & 63))
    if true {
        v14 = (("hello" + (v14 + v14)) + "foo bar")
    }
    if true {
        switch v15 {
        case 127:
            if len(v3) > 0 {
                v3[0] = v15
            }
        }
        v14 = (v14 + v14)
    } else {
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 bool = false
            _ = fi20
            _ = v21
        }
        switch v15 {
        case 0:
            if true {
                if ((!true) && true) {
                    delete(v5, 139)
                } else {
                    switch 1 {
                    case 9223372036854775807:
                        v2[(v15 ^ 9223372036854775807)] = v15
                    case -601:
                        for fi22 := 0; fi22 < 6; fi22++ {
                            v23 := GenS1{F0: 42, F1: 1}
                            v4 = (v4 % (((^(v4))) | 1))
                            v4 += v4
                            _ = fi22
                            _ = v23
                        }
                    default:
                        if (0.0 == ((0.1 / float64(3.14159)) * 0.1)) {
                            _ = len(v3)
                            v5[2] = ""
                        } else {
                            for fi24 := 0; fi24 < 20; fi24++ {
                                v25 := GenS1{F0: 0, F1: 163}
                                v14 = v14
                                if false {
                                    v4 -= genFn0("hello", v15)
                                    v14 = "foo bar"
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        if true {
                                            fmt.Println(v14)
                                            for fi26 := 0; fi26 < 2; fi26++ {
                                                v27 := GenS1{F0: 1, F1: 127}
                                                _ = fi26
                                                _ = v27
                                            }
                                        }
                                    default:
                                        _ = len(v3)
                                    }
                                    v28 := GenS1{F0: 2147483647, F1: 127}
                                    fmt.Println(fi24, v15, fi24)
                                    _ = v28
                                }
                                _ = fi24
                                _ = v25
                            }
                            if (false && false) {
                                for fi29 := 0; fi29 < 14; fi29++ {
                                    if true {
                                        fmt.Println(fi29, fi29, v15)
                                        v30 := GenS1{F0: 42, F1: 94}
                                        _ = v30
                                    }
                                    v3 = append(v3, v15)
                                    _ = len(v1)
                                    _ = fi29
                                }
                                v3 = append(v3, -2147483648)
                            }
                            var v31 uint8 = 42
                            _ = v31
                        }
                    }
                    v32 := GenS1{F0: -2147483648, F1: 1}
                    v33 := GenS2{F0: 42, F1: (v15 - (v15 / ((v15) | 1)))}
                    _ = v32
                    _ = v33
                }
                if (!(255 >= 2)) {
                    v15 = v15
                    v15 -= v15
                    v14 = ("foo bar" + v14)
                }
                v34 := GenS2{F0: 127, F1: -306}
                _ = v34
            } else {
                var v35 uint8 = 42
                v35 -= (v35 & (v35 / ((v35) | 1)))
                _ = v35
            }
        case 32768:
            v36 := GenS2{F0: 128, F1: v15}
            _ = v36
        default:
            for fi37 := 0; fi37 < 16; fi37++ {
                var v38 int16 = 32767
                _ = fi37
                _ = v38
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
