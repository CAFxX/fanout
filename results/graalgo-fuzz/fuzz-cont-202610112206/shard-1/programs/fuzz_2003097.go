package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int8; F1 int16; F2 string }
type GenS3 struct { F0 int8; F1 int32 }
func genFn0(p0 uint64) int64 {
    var v7 uint64 = (p0 & 0)
    p0 = v7
    var v8 int16 = 0
    v9 := GenS1{F0: (v8 + 42), F1: (v7 >> ((2) & 63)), F2: 2}
    fmt.Println(v7)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -514
}

func main() {
    v1 := []string{""}
    var v2 int8 = 1
    var v3 uint = 4294967295
    v4 := make(map[int]string)
    var v5 float64 = 1.5
    v6 := make(map[string]uint32)
    fmt.Println(v3, v5)
    var v10 string = ""
    v3 -= 1
    if (((!(122 >= 255)) || true) && true) {
        v1 = append(v1, (("" + ("\n\t" + v10)) + ""))
        v1 = append(v1, "")
    } else {
        fmt.Println(v3, v10)
        var v11 uint64 = 42
        _ = v11
    }
    _ = len(v6)
    var v12 uint64 = 1
    if ((!false) || false) {
        v4[9223372036854775807] = v10
    }
    var v13 float32 = -2.25
    v10 = "\n\t"
    for fi14 := 0; fi14 < 4; fi14++ {
        v12 = 1
        switch fi14 {
        case 791:
            _ = len(v4)
        case -651:
            if false {
                var v15 int8 = v2
                if (v10 > v10) {
                    for fi16 := 0; fi16 < 2; fi16++ {
                        if true {
                            var v17 int64 = 9223372036854775807
                            v18 := GenS1{F0: 1000, F1: (v12 / (((v12 << (((^(v3))) & 63))) | 1)), F2: 2}
                            if false {
                                var v19 uint8 = uint8(v12)
                                _ = v19
                            }
                            v20 := GenS1{F0: 1000, F1: 1, F2: 952}
                            v21 := GenS2{F0: 0, F1: -1, F2: "foo bar"}
                            _ = v17
                            _ = v18
                            _ = v20
                            _ = v21
                        }
                        v13 = (v13 + (1.5 * float32(3.14159)))
                        fmt.Println(fi14)
                        _ = fi16
                    }
                }
                if (false && (1 > -1)) {
                    v3 = (v3 % ((18446744073709551615) | 1))
                    for fi22 := 0; fi22 < 12; fi22++ {
                        var v23 uint8 = 42
                        var v24 float64 = 0.1
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
                _ = v15
            } else {
                for fi25 := 0; fi25 < 10; fi25++ {
                    v12 -= (^(v12))
                    _ = fi25
                }
                v26 := v4[(fi14 >> (((^(v3))) & 63))]
                for fi27 := 0; fi27 < 8; fi27++ {
                    v28 := GenS3{F0: (v2 << (((v3 * (v3 * (v3 ^ 349)))) & 63)), F1: 474}
                    var v29 int = fi27
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v26
            }
        }
        _ = fi14
    }
    fmt.Println(v3, v3, v2)
    _ = len(v4)
    if false {
        _ = len(v4)
        if false {
            v10 = v10
        }
    } else {
        v12 += (^(v12))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
