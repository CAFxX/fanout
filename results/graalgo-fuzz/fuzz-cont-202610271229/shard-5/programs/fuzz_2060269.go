package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0() uint32 {
    switch -74 {
    case -341:
        v6 := GenS1{F0: 320}
        _ = v6
    case 127:
        var v7 int8 = -45
        _ = v7
    }
    v8 := []int{-431, 32768, 9223372036854775807}
    _ = v8
    return 42
}
func genFn1(p0 uint64, p1 bool, p2 uint16) uint16 {
    for fi9 := 0; fi9 < 4; fi9++ {
        switch "foo bar" {
        case "\n\t":
            for fi10 := 0; fi10 < 13; fi10++ {
                if ((p0 >> ((18446744073709551615) & 63)) != (p0 >> ((1) & 63))) {
                    switch (^(fi10)) {
                    case -9223372036854775808:
                        p2 = p2
                    case 1:
                        var v11 uint16 = (p2 | p2)
                        _ = v11
                    }
                    switch "\n\t" {
                    case "foo bar":
                        if true {
                            v12 := []int{fi10}
                            fmt.Println(fi10, p2, p2)
                            p1 = p1
                            if p1 {
                                var v13 float64 = ((1e10 + 3.14159) / float64(1.5))
                                v14 := GenS2{F0: 127, F1: 1}
                                fmt.Println(p0, fi10)
                                _ = v13
                                _ = v14
                            }
                            switch "" {
                            case "hello":
                                var v15 uint32 = 65535
                                _ = v15
                            default:
                                var v16 bool = p1
                                _ = v16
                            }
                            _ = len(v12)
                            _ = v12
                        } else {
                            p1 = true
                        }
                    }
                    switch "\n\t" {
                    case "hello":
                        p1 = (fi10 <= 9223372036854775807)
                    case "foo bar":
                        var v17 string = (("\n\t" + "\n\t") + ("hello" + "a"))
                        _ = v17
                    default:
                        p0 += p0
                    }
                }
                if (uint(p2) == 2) {
                    if ((-2.25 / 1.5) >= (0.0 * (float32((((0.1 + ((-2.25 + float64(0.1)) + (((0.0 / 1.0) / 1.0) / 1.0))) - (3.14159 + 3.14159)) + (-2.25 - 1e10))) / float32(3.14159)))) {
                        var v18 float32 = 0.0
                        var v19 int = fi10
                        fmt.Println(fi9, v18, p2)
                        _ = v18
                        _ = v19
                    }
                    p0 = p0
                }
                _ = fi10
            }
        }
        v20 := []int{}
        var v21 int = (fi9 >> ((255) & 63))
        _ = fi9
        _ = v20
        _ = v21
    }
    if false {
        if true {
            var v22 bool = (-2.25 < -2.25)
            _ = v22
        } else {
            var v23 uint64 = (^(p0))
            var v24 bool = (float64((1e10 - 3.14159)) >= 1.5)
            v25 := GenS2{F0: 0, F1: 2}
            _ = v23
            _ = v24
            _ = v25
        }
    } else {
        p2 += (p2 & (p2 / (((^(p2))) | 1)))
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    v1 := []int{-2147483648, -2147483648}
    var v2 int8 = -1
    var v3 uint32 = 4294967295
    var v4 uint64 = 18446744073709551615
    var v5 int32 = 1
    v1 = append(v1, 184)
    v26 := GenS2{F0: v2, F1: uint32(v4)}
    switch 127 {
    case -810:
        v5 = (v5 ^ (v5 * v5))
    case 127:
        fmt.Println(v4)
    default:
        switch (("hello" + "") + "foo bar") {
        case "a":
            var v27 bool = true
            _ = v27
        }
    }
    var v28 uint16 = 0
    switch 9223372036854775807 {
    case -166:
        var v29 uint16 = (v28 + v28)
        _ = v29
    case 128:
        for fi30 := 0; fi30 < 2; fi30++ {
            v3 = v3
            v28 += 6
            _ = fi30
        }
    }
    v4 += (v4 % ((822) | 1))
    v2 = (v2 ^ (-(v2)))
    var v31 int8 = v2
    fmt.Println(v5, v4)
    var v32 uint64 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
