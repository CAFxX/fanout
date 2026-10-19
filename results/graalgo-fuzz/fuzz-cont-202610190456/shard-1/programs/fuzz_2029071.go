package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int64, p1 uint64, p2 int16) int {
    switch "foo bar" {
    case "\n\t":
        if true {
            p0 += -9223372036854775808
            v7 := GenS1{F0: "a"}
            _ = v7
        }
    default:
        if (2 < 18446744073709551615) {
            var v8 float32 = 3.14159
            v9 := []int{799}
            for fi10 := 0; fi10 < 14; fi10++ {
                for fi11 := 0; fi11 < 1; fi11++ {
                    switch "hello" {
                    case "foo bar":
                        p1 += (p1 ^ p1)
                    case "a":
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    default:
                        if (0 < 42) {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            var v12 uint = 2
                            _ = v12
                        } else {
                            p1 = p1
                            if ((1e10 + 0.1) >= (0.1 - v8)) {
                                v13 := GenS1{F0: "\n\t"}
                                p0 -= (p0 + p0)
                                _ = v13
                            }
                            if len(v9) > 0 {
                                v9[0] = fi10
                            }
                        }
                    }
                    var v14 bool = false
                    if len(v9) > 0 {
                        v9[0] = fi11
                    }
                    _ = fi11
                    _ = v14
                }
                for fi15 := 0; fi15 < 2; fi15++ {
                    v8 += ((v8 - v8) * v8)
                    _ = len(v9)
                    var v16 bool = false
                    _ = fi15
                    _ = v16
                }
                var v17 uint16 = 42
                _ = fi10
                _ = v17
            }
            _ = v8
            _ = v9
        } else {
            fmt.Println(p1)
            v18 := []int{127, -2147483648, 835}
            v18 = append(v18, 0)
            _ = v18
        }
    }
    p0 += 1
    _ = p0
    _ = p1
    _ = p2
    return 423
}
func genFn1(p0 uint) int8 {
    switch genFn0(195, 42, 1) {
    case 637:
        v19 := GenS1{F0: ("foo bar" + "hello")}
        _ = v19
    default:
        switch "\n\t" {
        case "":
            for fi20 := 0; fi20 < 4; fi20++ {
                v21 := []int{(^(fi20)), genFn0(-9223372036854775808, 4294967295, 1000), -960}
                _ = fi20
                _ = v21
            }
        case "foo bar":
            p0 += p0
        }
    }
    fmt.Println(p0)
    var v22 string = ("\n\t" + "a")
    var v23 uint16 = 255
    p0 -= p0
    _ = p0
    _ = v22
    _ = v23
    return 30
}

func main() {
    var v1 string = (("a" + "\n\t") + "foo bar")
    var v2 int16 = 1000
    var v3 int8 = 0
    var v4 uint8 = 69
    var v5 uint32 = 242
    var v6 int32 = 15
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS1{F0: (v1 + v1)}
        var v26 int8 = v3
        _ = fi24
        _ = v25
        _ = v26
    }
    v5 -= 2147483647
    var v27 float32 = 3.14159
    v28 := []int{-1, 127}
    for fi29 := 0; fi29 < 4; fi29++ {
        if len(v28) >= 2 {
            v28 = v28[1:len(v28)]
        }
        var v30 uint8 = (v4 - (v4 % ((127) | 1)))
        var v31 float64 = (0.0 * (-2.25 + 1e10))
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println(v6, v4)
    var v32 int8 = genFn1(43)
    fmt.Println(v1, v1, v6)
    var v33 int8 = v32
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
