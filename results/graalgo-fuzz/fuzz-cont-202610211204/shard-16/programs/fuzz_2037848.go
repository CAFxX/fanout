package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int64 }
type GenS2 struct { F0 int64; F1 int8; F2 string }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 int64, p1 bool) int {
    p0 -= 1
    switch 0 {
    case 9223372036854775807:
        var v6 int = -58
        _ = v6
    case 127:
        if (205 >= 255) {
            p0 -= (p0 | 1)
        }
    case 128:
        var v7 int16 = 1
        _ = v7
    default:
        fmt.Println(p1, p1)
    }
    v8 := GenS1{F0: 18446744073709551615, F1: "a", F2: p0}
    var v9 int8 = 0
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -2147483648
}
func genFn1(p0 int32, p1 uint8) int {
    for fi10 := 0; fi10 < 14; fi10++ {
        p0 = p0
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := []int{}
            v13 := GenS1{F0: 18446744073709551615, F1: "a", F2: 0}
            var v14 string = "\n\t"
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        var v15 float64 = 0.0
        _ = fi10
        _ = v15
    }
    v16 := GenS1{F0: 18446744073709551615, F1: "\n\t", F2: 937}
    v17 := GenS1{F0: 1, F1: ("a" + ((("a" + "\n\t") + "hello") + "")), F2: 0}
    p1 = 1
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return genFn0(9223372036854775807, (110 > 100))
}

func main() {
    var v1 int8 = 42
    v2 := []string{((("a" + "a") + "foo bar") + ""), "a"}
    v3 := make(map[string]int)
    var v4 uint8 = 127
    var v5 uint = 4294967295
    v1 = v1
    _ = len(v2)
    fmt.Println(v4)
    if ((!(42 > 1000)) && false) {
        var v18 string = "\n\t"
        v18 = v18
        _ = v18
    } else {
        if ((v1 & (v1 << ((v5) & 63))) > int8(v5)) {
            var v19 int64 = -9223372036854775808
            var v20 int32 = 0
            _ = v19
            _ = v20
        } else {
            v2 = append(v2, "\n\t")
            v1 += v1
            fmt.Println(v1, v1)
        }
    }
    var v21 uint64 = 22
    v21 = (^(v21))
    v5 = (v5 ^ (v5 + (v5 + 1)))
    v22 := GenS2{F0: 42, F1: (v1 >> (((v5 >> ((v5) & 63))) & 63)), F2: ("a" + "")}
    switch "hello" {
    case "hello":
        v21 -= 2
    case "foo bar":
        v23 := GenS2{F0: -431, F1: 32, F2: ""}
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        if false {
            var v25 float32 = (-2.25 - (3.14159 * (float32(float64(1.5)) + 3.14159)))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi26 := 0; fi26 < 5; fi26++ {
                if true {
                    v1 -= 100
                    fmt.Println(v4, v25)
                } else {
                    v5 += v5
                }
                fmt.Println(v1, v21, v5)
                v2 = append(v2, ("\n\t" + ""))
                _ = fi26
            }
            _ = v25
        } else {
            if true {
                v4 -= (v4 ^ v4)
            } else {
                if true {
                    var v27 string = "\n\t"
                    _ = v27
                } else {
                    if false {
                        var v28 int16 = -830
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 2; fi29++ {
                            v21 = v21
                            v1 -= v1
                            _ = fi29
                        }
                        var v30 bool = (true || (!(2 >= 2)))
                        fmt.Println(v5, v1, v5)
                        _ = v30
                    }
                }
            }
        }
        _ = fi24
    }
    v21 += v21
    v5 = (v5 + 63)
    v31 := GenS3{F0: (80 >= 9223372036854775807), F1: 42}
    var v32 uint = 255
    v1 += 37
    switch 9223372036854775807 {
    case -1:
        v4 = (v4 & 1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
