package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint64; F1 int8; F2 uint8 }
func genFn0(p0 int32) int8 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int64 = 1
        var v8 int64 = (v7 * (v7 - 1))
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 2147483647 {
    case 42:
        fmt.Println(p0)
    }
    p0 += p0
    _ = p0
    return -128
}
func genFn1(p0 uint) bool {
    v9 := []int{-856, 2147483647}
    if len(v9) > 0 {
        v9[0] = 2147483647
    }
    if false {
        var v10 float32 = 1.5
        var v11 int8 = -1
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v10
        _ = v11
    } else {
        _ = len(v9)
        var v12 uint8 = 92
        _ = v12
    }
    _ = p0
    _ = v9
    return (!(!false))
}

func main() {
    var v1 int32 = 0
    v2 := []string{("\n\t" + "\n\t"), "hello", "foo bar", "hello"}
    var v3 int16 = -1
    v4 := make(map[int]int)
    var v5 uint64 = 331
    if true {
        v3 = v3
        v4[32767] = int(v1)
        var v13 float64 = 1.5
        _ = v13
    }
    fmt.Println(v1, v1, v3)
    if ((^(v3)) <= v3) {
        v14 := GenS1{F0: 0, F1: "", F2: 2}
        if ((v5 + 433) >= v5) {
            for fi15 := 0; fi15 < 15; fi15++ {
                v1 += v1
                v5 += (v5 * 0)
                _ = fi15
            }
            switch 2147483647 {
            case 2147483647:
                if false {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        if (v3 <= v3) {
                            var v17 uint32 = 42
                            v1 = (v1 + 280)
                            switch (fi16 / ((2147483647) | 1)) {
                            case -9223372036854775808:
                                var v18 uint16 = 42
                                _ = v18
                            case 32767:
                                var v19 int32 = (v1 % ((v1) | 1))
                                _ = v19
                            case -46:
                                v20 := GenS1{F0: (fi16 ^ (fi16 & fi16)), F1: "hello", F2: 1}
                                _ = v20
                            }
                            _ = v17
                        } else {
                            switch "" {
                            case "hello":
                                v1 = (v1 / ((-557) | 1))
                            default:
                                if (!true) {
                                    if (("" + ((("" + "") + "hello") + "hello")) != "") {
                                        v21 := GenS2{F0: 2147483648, F1: true}
                                        v3 = v3
                                        _ = v21
                                    }
                                    var v22 int16 = (v3 & (v3 << ((uint(v5)) & 63)))
                                    _ = v22
                                }
                            }
                        }
                        var v23 uint64 = (v5 + 912)
                        var v24 int = (-(fi16))
                        _ = fi16
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(v1, v5)
                }
            case 4294967295:
                v3 -= (v3 + (-(v3)))
            case 841:
                fmt.Println(v5)
            default:
                v5 += (v5 * 42)
            }
        }
        fmt.Println(v5)
        _ = v14
    }
    v5 -= v5
    v1 += v1
    v5 += v5
    var v25 int32 = 0
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 float32 = (1e10 + (0.1 - ((3.14159 + 1e10) * (3.14159 + 3.14159))))
        v4[fi26] = (fi26 % ((fi26) | 1))
        _ = len(v2)
        _ = fi26
        _ = v27
    }
    v25 -= v1
    v1 = v25
    switch 4294967295 {
    case -286:
        if true {
            v28 := GenS1{F0: -955, F1: ("a" + "foo bar"), F2: 754}
            _ = v28
        } else {
            var v29 int64 = 0
            _ = v29
        }
    case 2147483647:
        var v30 int = 128
        _ = v30
    case 953:
        var v31 float32 = ((0.1 + float32(((3.14159 - 0.1) / (1.5 / 1.0)))) - (1e10 * (-2.25 / 1e10)))
        _ = v31
    default:
        v32 := GenS3{F0: 361, F1: 100, F2: 255}
        _ = v32
    }
    v33 := GenS1{F0: -9223372036854775808, F1: ("hello" + (((("\n\t" + "hello") + ("foo bar" + "foo bar")) + ("\n\t" + ("foo bar" + "a"))) + "")), F2: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
