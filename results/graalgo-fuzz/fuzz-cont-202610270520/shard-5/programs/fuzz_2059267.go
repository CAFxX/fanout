package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int32; F1 int8; F2 int8 }
func genFn0(p0 int, p1 int16, p2 bool) string {
    p2 = ((0 == 763) && p2)
    for fi6 := 0; fi6 < 3; fi6++ {
        p1 = (p1 / ((105) | 1))
        _ = fi6
    }
    v7 := GenS1{F0: 150, F1: 1}
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "hello"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 806
    var v3 int32 = 0
    var v4 int64 = -502
    var v5 int16 = -618
    var v8 uint32 = 500
    v9 := []int{-772, -2147483648, 4294967295}
    var v10 uint32 = 42
    if (-128 == 1) {
        for fi11 := 0; fi11 < 4; fi11++ {
            v2 += (v2 / ((601) | 1))
            _ = fi11
        }
        _ = len(v9)
        v12 := GenS1{F0: 35, F1: 42}
        _ = v12
    } else {
        var v13 int8 = 42
        _ = v13
    }
    var v14 uint = 4294967295
    fmt.Println(v2, v5)
    fmt.Println(v14)
    switch "" {
    case "a":
        if (!false) {
            if true {
                if (true || true) {
                    if false {
                        v3 = 2147483647
                    }
                } else {
                    if (true && (0 != 2147483647)) {
                        if (456 == (v4 >> ((18446744073709551615) & 63))) {
                            v5 = -32768
                            if (!true) {
                                var v15 uint32 = 4294967295
                                _ = v15
                            } else {
                                var v16 int8 = -1
                                var v17 int = -628
                                _ = v16
                                _ = v17
                            }
                            fmt.Println(v8, v3)
                        } else {
                            v18 := GenS2{F0: int32(v5), F1: 1, F2: 127}
                            fmt.Println(v8)
                            _ = v18
                        }
                    } else {
                        var v19 bool = true
                        v20 := v1[v10]
                        _ = v19
                        _ = v20
                    }
                    switch "" {
                    case "foo bar":
                        switch "\n\t" {
                        case "\n\t":
                            for fi21 := 0; fi21 < 20; fi21++ {
                                v8 = v10
                                v22 := GenS1{F0: 1, F1: 92}
                                _ = fi21
                                _ = v22
                            }
                        case "foo bar":
                            v10 -= (v8 + 886)
                        default:
                            delete(v1, 2)
                        }
                    }
                    var v23 int32 = v3
                    _ = v23
                }
            } else {
                if (!(!false)) {
                    if true {
                        v24 := v1[(v8 << (((^(v2))) & 63))]
                        _ = v24
                    } else {
                        v25 := GenS2{F0: (v3 * (v3 >> (((v2 << ((v14) & 63))) & 63))), F1: 42, F2: 0}
                        _ = len(v9)
                        _ = len(v1)
                        _ = v25
                    }
                }
            }
            fmt.Println(v14)
        }
    case "hello":
        var v26 float64 = 0.1
        _ = v26
    default:
        if false {
            _ = len(v1)
            for fi27 := 0; fi27 < 8; fi27++ {
                if (!(0.0 < 3.14159)) {
                    var v28 int8 = 100
                    _ = len(v1)
                    v29 := GenS2{F0: v3, F1: -1, F2: (v28 * v28)}
                    _ = v28
                    _ = v29
                } else {
                    v30 := GenS1{F0: 255, F1: -128}
                    for fi31 := 0; fi31 < 18; fi31++ {
                        v4 += v4
                        var v32 uint64 = 0
                        _ = fi31
                        _ = v32
                    }
                    fmt.Println(v14)
                    _ = v30
                }
                v10 = (v8 * v8)
                v8 = v8
                _ = fi27
            }
            v3 = (v3 * (v3 - -1))
        }
    }
    if true {
        for fi33 := 0; fi33 < 15; fi33++ {
            _ = len(v1)
            v34 := GenS1{F0: 255, F1: 100}
            v8 = v10
            _ = fi33
            _ = v34
        }
        if true {
            var v35 int = -9223372036854775808
            switch "hello" {
            case "a":
                fmt.Println(v14, v35, v14)
            case "hello":
                for fi36 := 0; fi36 < 9; fi36++ {
                    var v37 int8 = 62
                    v38 := v1[1]
                    fmt.Println(v37, v14)
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            case "foo bar":
                if (false || (886 > 4294967295)) {
                    v4 = (v4 ^ v4)
                }
            }
            _ = v35
        } else {
            var v39 uint32 = (v8 | (^(v8)))
            v3 += 415
            v40 := GenS1{F0: 127, F1: 42}
            _ = v39
            _ = v40
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
