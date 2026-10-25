package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 bool }

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]string)
    var v3 uint16 = 862
    var v4 float64 = -2.25
    var v5 uint8 = 168
    fmt.Println(v4)
    if false {
        v4 += v4
        v6 := GenS1{F0: (true || true)}
        _ = v6
    }
    if ((0.0 + float32((v4 - v4))) == -2.25) {
        v7 := GenS1{F0: (!true)}
        v8 := []int{-9223372036854775808, 128}
        _ = v7
        _ = v8
    } else {
        v1 = 1
        var v9 float32 = -2.25
        var v10 uint = 2
        _ = v10
        _ = v9
    }
    v2[(v1 / ((4294967295) | 1))] = "hello"
    var v11 bool = false
    if (true && false) {
        v1 = (v1 + (^(v1)))
        _ = len(v2)
    } else {
        var v12 int64 = 42
        _ = v12
    }
    v13 := []int{638, 1}
    switch (("\n\t" + "a") + "a") {
    case "a":
        if ((4294967295 > 18446744073709551615) && false) {
            v11 = (42 <= 2)
            fmt.Println(v11, v5)
        }
    case "foo bar":
        for fi14 := 0; fi14 < 8; fi14++ {
            v2[v1] = "foo bar"
            fmt.Println(v11, v3)
            _ = fi14
        }
    default:
        switch ("a" + "") {
        case "\n\t":
            var v15 int = 0
            _ = v15
        case "a":
            for fi16 := 0; fi16 < 4; fi16++ {
                fmt.Println(v4, fi16, v3)
                var v17 int32 = 206
                if ((v17 > v17) && v11) {
                    switch fi16 {
                    case -2147483648:
                        switch "\n\t" {
                        case "foo bar":
                            if ((v11 || (!false)) && (v11 || (1 >= (-(fi16))))) {
                                if v11 {
                                    for fi18 := 0; fi18 < 9; fi18++ {
                                        var v19 uint64 = 929
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        _ = fi18
                                        _ = v19
                                    }
                                    var v20 uint8 = 166
                                    _ = v20
                                }
                            } else {
                                var v21 bool = v11
                                for fi22 := 0; fi22 < 14; fi22++ {
                                    if v11 {
                                        fmt.Println(fi22, v17, v5)
                                        var v23 uint8 = (v5 >> ((969) & 63))
                                        _ = v23
                                    }
                                    _ = fi22
                                }
                                _ = v21
                            }
                        case "":
                            v4 = ((v4 + (0.1 / (((v4 / (((v4 / 1.5) - 3.14159) - v4)) * -2.25) * float64((3.14159 / ((1.5 * float32(v4)) * (1e10 + 0.0))))))) + (1.5 * 3.14159))
                        }
                    }
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v1 -= v1
                        if (v3 != v3) {
                            v25 := GenS2{F0: 42, F1: v11}
                            _ = v25
                        }
                        _ = fi24
                    }
                } else {
                    v2[(v1 * (v1 / ((464) | 1)))] = ("\n\t" + "foo bar")
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    var v26 uint64 = 0
                    _ = v26
                }
                _ = fi16
                _ = v17
            }
        case "":
            var v27 uint64 = 4294967295
            _ = v27
        default:
            v1 += v1
        }
    }
    var v28 int8 = 1
    v29 := v2[4294967295]
    v11 = (20 != (v28 / ((v28) | 1)))
    v29 = v29
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
