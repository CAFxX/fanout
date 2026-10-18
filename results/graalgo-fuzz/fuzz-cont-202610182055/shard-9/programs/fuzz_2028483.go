package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []string{("a" + "a"), ((("\n\t" + "foo bar") + "foo bar") + "foo bar")}
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    v4 := []uint32{579, 2, 65535, 1}
    _ = len(v2)
    var v5 uint32 = 42
    for fi6 := 0; fi6 < 5; fi6++ {
        if true {
            v5 += v5
            v5 = v5
        } else {
            switch 9223372036854775807 {
            case -9223372036854775808:
                fmt.Println(fi6, v5)
            case 9223372036854775807:
                for fi7 := 0; fi7 < 3; fi7++ {
                    v2[(v5 - 42)] = (^(fi7))
                    v2[v5] = (fi6 & (fi6 ^ fi7))
                    v5 = (v5 - 0)
                    _ = fi7
                }
            case -91:
                for fi8 := 0; fi8 < 16; fi8++ {
                    switch "" {
                    case "":
                        switch fi8 {
                        case -1:
                            var v9 bool = true
                            v9 = false
                            _ = v9
                        case -9223372036854775808:
                            var v10 string = ("hello" + (("a" + "") + "foo bar"))
                            _ = v10
                        case 1:
                            var v11 int8 = 0
                            _ = v11
                        default:
                            v12 := GenS1{F0: 605, F1: 128}
                            _ = v12
                        }
                    case "a":
                        delete(v3, (v5 - v5))
                    case "\n\t":
                        var v13 float32 = 0.0
                        if ((((1.5 + 0.0) / 3.14159) != 0.1) && true) {
                            fmt.Println(v13, v13, fi8)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v14 float64 = 3.14159
                            _ = v14
                        }
                        var v15 uint = 54
                        _ = v13
                        _ = v15
                    default:
                        if false {
                            v16 := v3[2147483647]
                            _ = v16
                        } else {
                            var v17 bool = (65535 > 2)
                            if (1e10 != (1e10 - (((1.5 / ((((1.5 + 1e10) + -2.25) - 0.1) / 1e10)) / 1e10) - 0.1))) {
                                if (v5 <= v5) {
                                    fmt.Println(fi6, fi8, v5)
                                } else {
                                    v18 := GenS2{F0: 159}
                                    v19 := GenS2{F0: 42}
                                    _ = v18
                                    _ = v19
                                }
                            }
                            v17 = ("\n\t" <= ("hello" + "a"))
                            _ = v17
                        }
                    }
                    _ = fi8
                }
            default:
                for fi20 := 0; fi20 < 20; fi20++ {
                    v5 -= v5
                    fmt.Println(fi6)
                    var v21 int64 = 42
                    _ = fi20
                    _ = v21
                }
            }
        }
        v5 = (v5 % (((v5 << ((833) & 63))) | 1))
        _ = fi6
    }
    v5 -= v5
    v5 = (v5 & v5)
    var v22 bool = false
    if (((!true) || false) || v22) {
        v23 := GenS1{F0: 18446744073709551615, F1: 129}
        _ = v23
    } else {
        v5 += (v5 | v5)
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 int8 = -128
            _ = fi24
            _ = v25
        }
    }
    fmt.Println(v22, v22)
    v26 := GenS2{F0: 127}
    var v27 float32 = (1e10 + (((float32(((float64(0.0) + 1.5) + 1e10)) / float32(3.14159)) - 1.5) - 1e10))
    var v28 uint32 = v5
    v27 = 1.5
    v29 := GenS1{F0: 274, F1: 4}
    var v30 int = 9223372036854775807
    v31 := GenS1{F0: 1, F1: 200}
    if len(v4) > 0 {
        v4[0] = v5
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
