package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int16; F1 string; F2 uint }
func genFn0() int32 {
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 float32 = 3.14159
        _ = fi4
        _ = v5
    }
    var v6 int16 = 349
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int64 = 9223372036854775807
        switch fi7 {
        case 673:
            if (0.1 <= (3.14159 - 3.14159)) {
                if ("hello" <= "hello") {
                    v6 -= (v6 >> ((396) & 63))
                    v9 := GenS1{F0: false}
                    for fi10 := 0; fi10 < 5; fi10++ {
                        v8 -= (v8 * v8)
                        _ = fi10
                    }
                    _ = v9
                } else {
                    var v11 int16 = (^(v6))
                    v8 = (v8 * v8)
                    var v12 int = (fi7 ^ fi7)
                    _ = v11
                    _ = v12
                }
                var v13 int8 = 31
                if (!true) {
                    if true {
                        v14 := []int{fi7, (fi7 | -95)}
                        var v15 int16 = v6
                        var v16 string = "hello"
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        v17 := []int{(fi7 >> ((0) & 63)), 2147483647, (fi7 | int(v8)), fi7}
                        fmt.Println(v8, v6, v8)
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        _ = v17
                    }
                    var v18 int32 = -2147483648
                    _ = v18
                } else {
                    v19 := []int{(fi7 / (((^(fi7))) | 1)), 127}
                    _ = v19
                }
                _ = v13
            } else {
                fmt.Println(fi7, v6, v6)
            }
        }
        _ = fi7
        _ = v8
    }
    _ = v6
    return 42
}

func main() {
    v1 := []int{-2147483648, 0, -259}
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 int = fi20
        v21 += fi20
        _ = fi20
        _ = v21
    }
    var v22 uint64 = 2
    v23 := GenS2{F0: -126}
    v22 -= 4294967295
    var v24 int16 = -32768
    var v25 int = 221
    if false {
        for fi26 := 0; fi26 < 18; fi26++ {
            if true {
                v27 := GenS1{F0: (true || (2147483647 > 688))}
                _ = v27
            } else {
                if ((v24 / ((268) | 1)) < 42) {
                    if (false || true) {
                        v24 += v24
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                } else {
                    switch "\n\t" {
                    case "":
                        if false {
                            v28 := GenS2{F0: -95}
                            var v29 float64 = (float64(1.5) * 0.1)
                            v30 := GenS2{F0: 127}
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            fmt.Println(v25, v24)
                            switch v25 {
                            case 762:
                                var v31 int64 = -9223372036854775808
                                _ = v31
                            default:
                                v22 += (v22 / ((v22) | 1))
                            }
                        }
                    case "\n\t":
                        var v32 bool = false
                        _ = v32
                    case "hello":
                        var v33 int16 = v24
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 14; fi34++ {
                        v25 = (fi34 % ((v25) | 1))
                        _ = fi34
                    }
                    fmt.Println(v25, fi26, fi26)
                }
                v1 = append(v1, v25)
            }
            fmt.Println(v22, v24, v22)
            var v35 uint16 = 32767
            _ = fi26
            _ = v35
        }
        v36 := GenS3{F0: v22, F1: 18446744073709551615}
        v37 := GenS2{F0: 0}
        _ = v36
        _ = v37
    }
    var v38 int32 = -2147483648
    v24 = v24
    v39 := v3[v25]
    switch 0 {
    case 128:
        fmt.Println(v38, v38, v22)
    case -9223372036854775808:
        v40 := GenS4{F0: (v24 >> ((744) & 63)), F1: (((("a" + ("\n\t" + v39)) + v39) + (("\n\t" + "") + ("\n\t" + ""))) + v39), F2: 291}
        _ = v40
    case -294:
        if true {
            v41 := GenS1{F0: true}
            _ = v41
        }
    }
    var v42 uint16 = 65535
    v22 -= v22
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v39
    _ = v42
}
