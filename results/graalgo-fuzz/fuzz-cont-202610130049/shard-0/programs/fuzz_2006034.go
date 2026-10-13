package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint64 }
func genFn0(p0 uint, p1 string, p2 uint) uint8 {
    v6 := GenS1{F0: 530, F1: "foo bar"}
    v7 := []int{-1, 9223372036854775807, -616}
    fmt.Println(p0)
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []string{"hello"}
    var v2 string = "\n\t"
    v3 := []int{-859, -1, 803}
    var v4 uint = 4294967295
    v5 := []string{v2}
    v4 -= 42
    var v8 uint64 = 2
    if true {
        fmt.Println(v8, v8)
    } else {
        v1 = append(v1, "")
    }
    if true {
        v9 := GenS1{F0: 65535, F1: "hello"}
        switch -1 {
        case 32768:
            v8 = 508
        case 9223372036854775807:
            var v10 int32 = -523
            _ = v10
        case 611:
            if (true && (0.0 > (-2.25 - (1.5 + ((((((1e10 * 0.0) * ((-2.25 * ((3.14159 * 1e10) - 0.1)) / -2.25)) + (1.5 / 0.1)) - 0.1) - 0.0) / 1.5))))) {
                if true {
                    if (!false) {
                        switch -9223372036854775808 {
                        case -2147483648:
                            switch v2 {
                            case "\n\t":
                                v4 += (v4 / ((v4) | 1))
                            default:
                                v11 := GenS1{F0: 32767, F1: "a"}
                                _ = v11
                            }
                        case 32767:
                            var v12 int = -1
                            _ = v12
                        case 403:
                            v13 := GenS2{F0: -126, F1: -1}
                            _ = v13
                        default:
                            v2 = v2
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if false {
                            _ = len(v1)
                        } else {
                            v2 = (v2 + "hello")
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                }
                v4 = (v4 | (v4 / ((v4) | 1)))
            } else {
                var v14 bool = (0 <= 42)
                switch "foo bar" {
                case "a":
                    if (!(-9223372036854775808 > 1)) {
                        fmt.Println(v14)
                        v15 := GenS1{F0: 42, F1: "foo bar"}
                        _ = v15
                    } else {
                        v4 += v4
                        if false {
                            if len(v1) > 0 {
                                v1[0] = v2
                            }
                        }
                        var v16 uint64 = v8
                        _ = v16
                    }
                case "":
                    v2 = v2
                case "foo bar":
                    fmt.Println(v4)
                default:
                    for fi17 := 0; fi17 < 5; fi17++ {
                        if (1000 == -666) {
                            if len(v3) > 0 {
                                v3[0] = (fi17 + (^(fi17)))
                            }
                            var v18 int = fi17
                            v4 = 0
                            _ = v18
                        } else {
                            v14 = false
                        }
                        _ = fi17
                    }
                }
                v8 -= 18446744073709551615
                _ = v14
            }
        default:
            fmt.Println(v8)
        }
        _ = v9
    }
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    switch ("a" + v2) {
    case "foo bar":
        v8 = v8
    case "\n\t":
        var v19 uint32 = 149
        _ = v19
    default:
        var v20 uint = (v4 % ((0) | 1))
        _ = v20
    }
    switch (("\n\t" + v2) + v2) {
    case "\n\t":
        fmt.Println(v8)
    default:
        v21 := GenS1{F0: 255, F1: "hello"}
        _ = v21
    }
    switch 1 {
    case -2147483648:
        fmt.Println(v8, v2, v4)
    case 47:
        var v22 float64 = 0.1
        _ = v22
    case 560:
        switch 508 {
        case 0:
            v23 := GenS3{F0: v8, F1: 0, F2: 953}
            _ = v23
        case 32768:
            fmt.Println(v2)
        case 128:
            switch v2 {
            case "a":
                fmt.Println(v8, v2)
            }
        }
    default:
        v24 := GenS4{F0: 405, F1: genFn0(4294967295, v2, v4), F2: (^(v8))}
        _ = v24
    }
    v4 = v4
    switch 32767 {
    case 834:
        v8 = uint64(v4)
    }
    var v25 int = 112
    var v26 int32 = 1
    for fi27 := 0; fi27 < 6; fi27++ {
        v5 = append(v5, "\n\t")
        v26 = v26
        for fi28 := 0; fi28 < 8; fi28++ {
            fmt.Println(v26, v26)
            v4 = (v4 * 18446744073709551615)
            _ = fi28
        }
        _ = fi27
    }
    fmt.Println(v25, v26, v25)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
