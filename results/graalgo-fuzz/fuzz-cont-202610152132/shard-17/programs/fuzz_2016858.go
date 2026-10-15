package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint }
func genFn0(p0 bool, p1 int, p2 uint64) uint32 {
    p1 -= p1
    p2 = p2
    p2 -= p2
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 int = (p1 % ((fi4) | 1))
        if p0 {
            v6 := []int{(fi4 / (((fi4 << ((65535) & 63))) | 1)), (fi4 >> ((65535) & 63)), p1, 2147483647}
            v7 := GenS1{F0: -2147483648, F1: (p2 * p2), F2: 768}
            if (float64(0.0) == 1e10) {
                var v8 int32 = -1
                for fi9 := 0; fi9 < 8; fi9++ {
                    var v10 uint16 = 0
                    v11 := GenS2{F0: 255}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = v8
            } else {
                v12 := GenS1{F0: -2147483648, F1: 42, F2: 923}
                var v13 int16 = 42
                for fi14 := 0; fi14 < 11; fi14++ {
                    v6 = append(v6, (-(fi14)))
                    if len(v6) > 0 {
                        v6[0] = fi4
                    }
                    var v15 string = "hello"
                    switch v15 {
                    case "hello":
                        var v16 int64 = int64(p2)
                        _ = v16
                    case "a":
                        fmt.Println(v13, v15, v15)
                    case "\n\t":
                        v6 = append(v6, int(p2))
                    }
                    _ = fi14
                    _ = v15
                }
                _ = v12
                _ = v13
            }
            _ = v6
            _ = v7
        }
        v17 := []int{(p1 + (p1 & 32768))}
        _ = fi4
        _ = v17
        _ = v5
    }
    p0 = (-2147483648 >= 2147483647)
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}
func genFn1(p0 uint64) uint16 {
    var v18 uint32 = 300
    switch "" {
    case "a":
        p0 -= p0
    }
    v19 := GenS2{F0: 42}
    _ = p0
    _ = v18
    _ = v19
    return 65535
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float64 = (3.14159 + 1e10)
    v3 := []uint32{v1}
    switch "foo bar" {
    case "foo bar":
        v2 += float64(3.14159)
    case "hello":
        v2 = (float64(3.14159) / v2)
    case "":
        fmt.Println(v1, v2)
    default:
        fmt.Println(v2, v2)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint8 = 1
            _ = fi21
            _ = v22
        }
        var v23 bool = (true || ((!((^(fi20)) == fi20)) || (v1 < v1)))
        v23 = v23
        _ = fi20
        _ = v23
    }
    var v24 int16 = -1
    v1 += (^(v1))
    v2 += (v2 * v2)
    v2 = float64(0.1)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
