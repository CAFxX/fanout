package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 string }

func main() {
    var v1 int64 = 0
    var v2 bool = (914 >= 2147483648)
    v3 := make(map[string]bool)
    v4 := []int{32767, 32768, 2147483647}
    v5 := GenS1{F0: (true || (!v2))}
    var v6 int8 = -128
    if (-1 == 1) {
        for fi7 := 0; fi7 < 17; fi7++ {
            if v2 {
                v8 := GenS2{F0: 42, F1: 42}
                _ = v8
            }
            switch 243 {
            case 302:
                fmt.Println(v2)
            default:
                if v2 {
                    if ((18446744073709551615 < 1) || v2) {
                        var v9 float64 = 0.0
                        switch (("a" + "a") + "a") {
                        case "":
                            for fi10 := 0; fi10 < 6; fi10++ {
                                if ((1000 <= 42) || (127 == 104)) {
                                    var v11 uint = 2
                                    v9 = v9
                                    fmt.Println(v9, v9)
                                    _ = v11
                                }
                                v12 := GenS2{F0: -2147483648, F1: 42}
                                v6 = (v6 ^ int8(v1))
                                _ = fi10
                                _ = v12
                            }
                        default:
                            for fi13 := 0; fi13 < 7; fi13++ {
                                v6 -= v6
                                _ = fi13
                            }
                        }
                        switch 0 {
                        case 32767:
                            var v14 uint8 = 171
                            _ = v14
                        case 761:
                            fmt.Println(v9)
                        default:
                            v2 = v2
                        }
                        _ = v9
                    } else {
                        v1 = v1
                    }
                } else {
                    var v15 uint = 1
                    switch "\n\t" {
                    case "":
                        if v2 {
                            var v16 uint8 = 0
                            _ = v16
                        }
                    }
                    _ = v15
                }
            }
            _ = fi7
        }
        var v17 string = "hello"
        _ = v17
    } else {
        v1 = (-(v1))
        var v18 uint8 = 0
        v2 = v2
        _ = v18
    }
    v3["a"] = false
    var v19 string = "\n\t"
    var v20 uint32 = 42
    if false {
        for fi21 := 0; fi21 < 11; fi21++ {
            if v2 {
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 uint8 = 0
                    var v24 float64 = 1.5
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                if (2 > 481) {
                    v2 = (float64(0.1) == (-2.25 + (float64(1.5) * 0.1)))
                    if len(v4) > 0 {
                        v4[0] = (fi21 % (((fi21 % ((fi21) | 1))) | 1))
                    }
                    v20 += v20
                } else {
                    v19 = ("a" + ((v19 + (((("\n\t" + v19) + v19) + v19) + v19)) + v19))
                    var v25 int = fi21
                    v4 = append(v4, -2147483648)
                    _ = v25
                }
                v1 -= (v1 - 1)
            } else {
                switch fi21 {
                case -536:
                    var v26 uint8 = 2
                    _ = v26
                default:
                    v2 = false
                }
            }
            v27 := GenS3{F0: (v19 + ((v19 + v19) + ""))}
            _ = fi21
            _ = v27
        }
    } else {
        var v28 int = 2147483647
        _ = v28
    }
    var v29 uint = 0
    _ = len(v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v20 = (v20 / (((v20 / ((v20) | 1))) | 1))
    v30 := v3[v19]
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 int = (fi31 & (-(fi31)))
        v4 = append(v4, fi31)
        switch -733 {
        case 42:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 32768:
            var v33 int = -2147483648
            _ = v33
        default:
            fmt.Println(v32, v2, v32)
        }
        _ = fi31
        _ = v32
    }
    v34 := GenS2{F0: -800, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
