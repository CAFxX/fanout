package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0() int32 {
    var v4 float32 = -2.25
    v5 := []int{-9223372036854775808, 722, 128, 2147483647}
    if len(v5) > 0 {
        v5[0] = 1
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(v4, v4)
        var v7 float32 = v4
        _ = fi6
        _ = v7
    }
    _ = v4
    _ = v5
    return 42
}
func genFn1() uint16 {
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 bool = (true || true)
        if false {
            v9 = ((!(!((0 == -295) || false))) && v9)
            v10 := []int{-749}
            fmt.Println(v9)
            _ = v10
        } else {
            fmt.Println(fi8)
            v11 := []int{(fi8 - (fi8 | fi8))}
            _ = v11
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            switch -795 {
            case 1:
                v9 = (!v9)
            case -9223372036854775808:
                v13 := GenS1{F0: v9, F1: fi12, F2: 141}
                _ = v13
            case 0:
                v9 = v9
            }
            _ = fi12
        }
        _ = fi8
        _ = v9
    }
    switch 127 {
    case -718:
        if false {
            if (4294967295 < 65535) {
                if false {
                    v14 := GenS1{F0: true, F1: -320, F2: 42}
                    var v15 int = -387
                    _ = v14
                    _ = v15
                }
                v16 := GenS1{F0: (false && ((float32(float64(-2.25)) + (((0.1 * -2.25) + float32(float64(-2.25))) / 1e10)) < ((1e10 - 1e10) - float32(1e10)))), F1: 1, F2: 65535}
                var v17 int64 = -9223372036854775808
                _ = v16
                _ = v17
            }
            var v18 uint = 18446744073709551615
            _ = v18
        } else {
            v19 := GenS1{F0: false, F1: -2147483648, F2: 42}
            if (("a" == "foo bar") && (18446744073709551615 != 514)) {
                var v20 int = -2147483648
                v20 -= 32767
                switch (v20 >> ((18446744073709551615) & 63)) {
                case -1:
                    v20 = (v20 ^ v20)
                default:
                    v21 := []int{128}
                    _ = v21
                }
                _ = v20
            } else {
                v22 := []int{224, -9223372036854775808}
                if ((583 != 18446744073709551615) || true) {
                    var v23 int16 = 1000
                    if (false || true) {
                        v24 := GenS1{F0: false, F1: 4294967295, F2: 255}
                        if (0.1 == (0.0 - 0.0)) {
                            switch "" {
                            case "foo bar":
                                var v25 float64 = -2.25
                                _ = v25
                            default:
                                var v26 uint32 = 4294967295
                                _ = v26
                            }
                            v23 -= v23
                            v23 = (v23 % ((v23) | 1))
                        } else {
                            v23 = v23
                            if (!false) {
                                v23 = 1000
                                v27 := GenS1{F0: (1e10 != 0.0), F1: 4294967295, F2: 1}
                                _ = v27
                            } else {
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                                for fi28 := 0; fi28 < 3; fi28++ {
                                    _ = len(v22)
                                    _ = len(v22)
                                    var v29 uint = 4294967295
                                    _ = fi28
                                    _ = v29
                                }
                                v23 = v23
                            }
                            v23 += v23
                        }
                        _ = v24
                    }
                    _ = v23
                }
                v22 = append(v22, 482)
                _ = v22
            }
            var v30 int64 = -1
            _ = v19
            _ = v30
        }
    }
    if ((!true) || false) {
        var v31 uint = 368
        fmt.Println(v31)
        _ = v31
    }
    return 92
}

func main() {
    var v1 int8 = 127
    v2 := []int{338}
    var v3 uint8 = 255
    fmt.Println(v1, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v32 uint64 = 19
    switch "\n\t" {
    case "":
        v33 := GenS2{F0: v3}
        _ = v33
    }
    if false {
        v3 = v3
        v34 := GenS2{F0: 227}
        _ = v34
    }
    v1 = v1
    var v35 uint16 = 255
    v1 -= 100
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
}
