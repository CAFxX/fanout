package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint }
func genFn0(p0 int, p1 int64) int {
    if true {
        var v7 float64 = 1e10
        var v8 string = "a"
        var v9 uint8 = 0
        _ = v7
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: -1}
        _ = v10
    }
    p1 = 0
    _ = p0
    _ = p1
    return (-(p0))
}
func genFn1(p0 int32) uint32 {
    var v11 int = 1
    if false {
        v11 += genFn0(genFn0(v11, -2), -833)
        var v12 uint64 = 172
        switch "" {
        case "":
            if true {
                v12 = (^(v12))
            } else {
                v13 := GenS1{F0: 0}
                p0 += p0
                _ = v13
            }
        case "a":
            fmt.Println(p0, p0, v11)
        case "foo bar":
            switch (v11 | (v11 >> ((157) & 63))) {
            case 308:
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 int32 = -1
                    _ = fi14
                    _ = v15
                }
            case 9223372036854775807:
                var v16 uint32 = 211
                _ = v16
            default:
                if ((p0 + p0) == (p0 ^ p0)) {
                    p0 += p0
                    v12 = (v12 % ((v12) | 1))
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        var v17 uint16 = 414
                        _ = v17
                    default:
                        v18 := GenS1{F0: 127}
                        _ = v18
                    }
                }
            }
        default:
            var v19 bool = false
            _ = v19
        }
        _ = v12
    }
    if (42 > 0) {
        if true {
            v20 := []int{genFn0(v11, -9223372036854775808), (-(v11)), v11, genFn0((v11 >> ((0) & 63)), 174)}
            _ = v20
        } else {
            var v21 int64 = 0
            for fi22 := 0; fi22 < 6; fi22++ {
                fmt.Println(fi22, v11, p0)
                v23 := []int{(v11 + -1)}
                _ = fi22
                _ = v23
            }
            _ = v21
        }
        if (!(false || ((float32(0.1) / ((0.1 - 1e10) + 0.0)) != (1.5 + 0.0)))) {
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 float64 = 0.1
                _ = fi24
                _ = v25
            }
        } else {
            p0 = -2147483648
        }
        v26 := []int{9223372036854775807}
        _ = v26
    } else {
        var v27 string = "a"
        v28 := GenS2{F0: -128, F1: 1}
        v27 = ("foo bar" + "\n\t")
        _ = v27
        _ = v28
    }
    p0 -= (p0 % ((p0) | 1))
    _ = p0
    _ = v11
    return 1
}

func main() {
    var v1 uint = 42
    var v2 string = "a"
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 uint64 = 738
    v6 := make(map[int]bool)
    fmt.Println(v1)
    fmt.Println(v1, v1)
    v29 := GenS3{F0: (v1 + 4294967295)}
    v1 -= 0
    var v30 uint8 = 214
    var v31 string = "foo bar"
    var v32 float64 = -2.25
    v32 -= -2.25
    fmt.Println(v30)
    var v33 float32 = 3.14159
    v2 = (v31 + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
