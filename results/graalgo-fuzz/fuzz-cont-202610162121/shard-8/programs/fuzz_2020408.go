package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0() uint {
    var v4 int = -2147483648
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(fi5, fi5)
        _ = fi5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 65535, F1: 255}
        fmt.Println(fi6, v4)
        v4 -= (v4 >> ((65535) & 63))
        _ = fi6
        _ = v7
    }
    _ = v4
    return 18446744073709551615
}
func genFn1(p0 uint32) uint64 {
    if (-481 == -1) {
        if true {
            v8 := GenS1{F0: 18446744073709551615, F1: 1}
            _ = v8
        }
        v9 := GenS1{F0: 925, F1: 231}
        _ = v9
    }
    var v10 int = 42
    v11 := GenS1{F0: genFn0(), F1: 163}
    _ = p0
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int8 = -5
    var v2 float64 = 1e10
    var v3 uint32 = 65535
    v12 := GenS1{F0: 671, F1: 139}
    var v13 int64 = 42
    v14 := []int{381, 32768, -369, 1}
    v1 = (v1 ^ (v1 & (v1 + -128)))
    var v15 uint32 = (v3 % (((^(v3))) | 1))
    if (862 == 4294967295) {
        v13 = (v13 / (((v13 % (((v13 >> ((1) & 63))) | 1))) | 1))
    } else {
        _ = len(v14)
    }
    v2 = v2
    if (false && (true || (3.14159 < -2.25))) {
        var v16 int16 = 339
        v2 -= -2.25
        _ = v16
    }
    v15 -= v15
    if (uint8(v15) > 2) {
        v1 += 0
    } else {
        var v17 int8 = v1
        switch ("" + "a") {
        case "a":
            for fi18 := 0; fi18 < 1; fi18++ {
                for fi19 := 0; fi19 < 14; fi19++ {
                    var v20 string = "foo bar"
                    switch ("a" + v20) {
                    case "\n\t":
                        v21 := GenS2{F0: "\n\t", F1: v13}
                        _ = v21
                    case "hello":
                        var v22 uint64 = 168
                        _ = v22
                    case "a":
                        if (65535 != genFn0()) {
                            var v23 string = v20
                            _ = v23
                        }
                    }
                    _ = fi19
                    _ = v20
                }
                for fi24 := 0; fi24 < 2; fi24++ {
                    v13 -= (v13 + (v13 << ((18446744073709551615) & 63)))
                    _ = fi24
                }
                if (("hello" + ("a" + (("" + "hello") + ("a" + "")))) <= "") {
                    var v25 int = fi18
                    var v26 int = (v25 | v25)
                    switch "\n\t" {
                    case "hello":
                        v14 = append(v14, v25)
                    default:
                        var v27 float64 = (0.0 + 0.1)
                        _ = v27
                    }
                    _ = v25
                    _ = v26
                } else {
                    v13 += v13
                    fmt.Println(v1, v1)
                }
                _ = fi18
            }
        case "hello":
            if (824 <= 42) {
                var v28 bool = true
                v3 -= v15
                _ = v28
            } else {
                switch "" {
                case "a":
                    var v29 bool = false
                    _ = v29
                case "\n\t":
                    switch "" {
                    case "a":
                        v3 += (^(v3))
                    default:
                        fmt.Println(v13, v17)
                    }
                case "hello":
                    v30 := GenS1{F0: genFn0(), F1: 2}
                    _ = v30
                default:
                    var v31 uint8 = 2
                    _ = v31
                }
                if false {
                    v32 := GenS1{F0: 629, F1: 127}
                    _ = v32
                }
                if true {
                    v14 = append(v14, 2147483647)
                }
            }
        case "\n\t":
            v13 = 9223372036854775807
        default:
            fmt.Println(v17, v15)
        }
        fmt.Println(v15, v3)
        _ = v17
    }
    var v33 int = 4294967295
    fmt.Println(v33, v33)
    fmt.Println(v1, v15, v15)
    v33 = (v33 ^ (v33 << ((1) & 63)))
    v33 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v33
}
