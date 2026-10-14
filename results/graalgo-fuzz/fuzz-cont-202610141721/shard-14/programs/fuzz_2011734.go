package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 int; F1 uint16 }

func main() {
    var v1 string = ("\n\t" + ("a" + ("a" + "\n\t")))
    var v2 uint32 = 714
    var v3 int64 = 141
    var v4 float64 = 1.5
    var v5 int8 = -126
    var v6 int64 = (v3 & v3)
    var v7 int = -1
    var v8 uint16 = 910
    if (!((float64(1e10) - v4) != 1e10)) {
        switch ("\n\t" + ((v1 + "\n\t") + (v1 + v1))) {
        case "hello":
            fmt.Println(v1)
        case "":
            v3 = -9223372036854775808
        default:
            fmt.Println(v8, v4, v7)
        }
        v1 = "a"
    } else {
        switch "hello" {
        case "hello":
            var v9 uint = 18446744073709551615
            _ = v9
        case "\n\t":
            var v10 int32 = 2147483647
            _ = v10
        default:
            switch v1 {
            case "\n\t":
                if ((v1 + "foo bar") <= v1) {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        if true {
                            v4 += float64((0.0 / float32(((v4 / v4) + v4))))
                        } else {
                            if (-888 == int32(v5)) {
                                v12 := []int{-308, v7}
                                _ = v12
                            }
                            switch (fi11 >> ((2) & 63)) {
                            case 0:
                                var v13 int32 = -58
                                _ = v13
                            case -9223372036854775808:
                                v8 = (v8 * v8)
                            case 32768:
                                if (4294967295 < 0) {
                                    var v14 float64 = 0.0
                                    switch (("" + (v1 + v1)) + v1) {
                                    case "a":
                                        var v15 uint = 1
                                        _ = v15
                                    case "foo bar":
                                        var v16 uint = uint(v2)
                                        _ = v16
                                    case "hello":
                                        v2 += 484
                                    default:
                                        v17 := []int{v7, (fi11 >> ((42) & 63)), 541}
                                        _ = v17
                                    }
                                    if false {
                                        var v18 uint32 = 2
                                        for fi19 := 0; fi19 < 14; fi19++ {
                                            v20 := []int{(fi11 + -1), 1}
                                            v21 := GenS1{F0: 42, F1: (v5 & v5), F2: (true || (0 < 0))}
                                            _ = fi19
                                            _ = v20
                                            _ = v21
                                        }
                                        if (42 != 2147483647) {
                                            switch v1 {
                                            case "a":
                                                v8 -= (v8 ^ 255)
                                            case "foo bar":
                                                var v22 uint = 18446744073709551615
                                                _ = v22
                                            case "hello":
                                                v23 := []int{(fi11 & (fi11 + v7)), 127, fi11}
                                                _ = v23
                                            default:
                                                for fi24 := 0; fi24 < 12; fi24++ {
                                                    v5 = (^(v5))
                                                    _ = fi24
                                                }
                                            }
                                        } else {
                                            v25 := GenS1{F0: 1000, F1: v5, F2: (743 > 1)}
                                            var v26 int16 = 32767
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = v18
                                    } else {
                                        var v27 bool = false
                                        v28 := []int{0, v7, v7}
                                        _ = len(v28)
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = v14
                                }
                            }
                            fmt.Println(v2, v7)
                        }
                        v29 := []int{int(v3), (fi11 << ((9) & 63)), fi11, -815}
                        _ = fi11
                        _ = v29
                    }
                    v30 := GenS2{F0: v7, F1: v8}
                    _ = v30
                } else {
                    v5 = (v5 / ((v5) | 1))
                    v31 := []int{(-(v7)), (v7 >> ((4294967295) & 63)), (^(v7)), v7}
                    _ = v31
                }
            }
        }
    }
    var v32 uint64 = 42
    v6 = (^(v6))
    if false {
        if false {
            v33 := []int{(v7 & v7), (v7 ^ (v7 >> ((uint(v2)) & 63))), v7}
            v3 -= -1
            v3 = 209
            _ = v33
        } else {
            v32 = 1
        }
        v5 = v5
        v34 := []int{int(v6), (-(v7)), (v7 | (v7 + 128))}
        _ = v34
    }
    switch v1 {
    case "a":
        var v35 int16 = int16(v7)
        _ = v35
    default:
        v7 -= 9223372036854775807
    }
    fmt.Println(v7, v6, v8)
    fmt.Println(v4)
    v3 = 42
    var v36 float32 = (-2.25 / float32(-2.25))
    v37 := GenS1{F0: 0, F1: (v5 - int8(v6)), F2: (true && (true && false))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
