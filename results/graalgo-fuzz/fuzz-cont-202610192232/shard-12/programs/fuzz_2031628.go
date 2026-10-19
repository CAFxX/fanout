package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 string; F2 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int16) int8 {
    p0 -= p0
    var v5 uint32 = 587
    _ = p0
    _ = v5
    return 79
}
func genFn1(p0 bool, p1 bool, p2 uint) uint64 {
    v6 := GenS1{F0: genFn0(-667)}
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 uint = 65535
    var v2 uint64 = 42
    var v3 int16 = -242
    v4 := make(map[int]string)
    delete(v4, -1)
    if (2147483647 > 2147483647) {
        v2 = genFn1((("hello" + "hello") >= ""), false, 1)
        v7 := GenS1{F0: 0}
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            if false {
                var v10 int = fi8
                var v11 int32 = 1
                v4[-9223372036854775808] = ("foo bar" + "\n\t")
                _ = v10
                _ = v11
            } else {
                fmt.Println(fi9, fi8)
                if false {
                    v12 := GenS2{F0: uint16(v2), F1: (("hello" + ("hello" + "hello")) + ""), F2: 0}
                    fmt.Println(v1, fi9, v2)
                    _ = v12
                } else {
                    v4[(fi9 + (fi9 | (fi8 + fi9)))] = "a"
                    var v13 float64 = 3.14159
                    var v14 uint = 255
                    _ = v13
                    _ = v14
                }
                v15 := GenS3{F0: v1}
                _ = v15
            }
            _ = fi9
        }
        _ = fi8
    }
    v2 = (v2 % ((v2) | 1))
    if true {
        var v16 uint = 1
        v1 += (v16 ^ (v16 - 42))
        _ = v16
    } else {
        for fi17 := 0; fi17 < 18; fi17++ {
            fmt.Println(v3, v1)
            switch fi17 {
            case 1:
                if ((fi17 | (fi17 * 2147483647)) > (fi17 >> ((v1) & 63))) {
                    if true {
                        fmt.Println(v3, v1)
                        v2 += v2
                    }
                    var v18 int64 = 42
                    v18 += (-(v18))
                    _ = v18
                }
            default:
                var v19 bool = false
                _ = v19
            }
            for fi20 := 0; fi20 < 12; fi20++ {
                v2 -= (^(v2))
                v1 = v1
                _ = fi20
            }
            _ = fi17
        }
        var v21 uint32 = 42
        _ = v21
    }
    fmt.Println(v2)
    v4[-2147483648] = ("a" + ("a" + ""))
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
