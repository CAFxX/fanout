package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int64 }
func genFn0(p0 int64, p1 uint32, p2 uint8) uint {
    var v5 int = -9223372036854775808
    var v6 bool = (p2 <= p2)
    fmt.Println(v6, p1, p1)
    if v6 {
        for fi7 := 0; fi7 < 20; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                if v6 {
                    if (((((((float32((0.0 * float64(3.14159))) + 3.14159) / 1e10) - 0.1) - (((3.14159 / 1.5) + 1e10) / 1.0)) - 1e10) - 3.14159) != (3.14159 * 3.14159)) {
                        v9 := GenS1{F0: -1, F1: 42}
                        switch "\n\t" {
                        case "":
                            v10 := GenS1{F0: 100, F1: 575}
                            v11 := GenS2{F0: 330}
                            fmt.Println(p2, fi8, v5)
                            _ = v10
                            _ = v11
                        }
                        _ = v9
                    } else {
                        if (!((1e10 - (1e10 / (-2.25 + -2.25))) < (3.14159 - (1e10 + ((-2.25 - ((3.14159 / 0.1) + (0.1 / 0.1))) - 1e10))))) {
                            var v12 int32 = 977
                            _ = v12
                        }
                        var v13 uint8 = 5
                        _ = v13
                    }
                }
                v6 = (!v6)
                v6 = true
                _ = fi8
            }
            _ = fi7
        }
        v14 := GenS2{F0: 373}
        _ = v14
    } else {
        fmt.Println(p1, v5, p1)
        v15 := GenS1{F0: 1, F1: 42}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 uint16, p1 uint64, p2 int) uint {
    if (((1e10 - (1e10 + 1e10)) * 1.5) == (0.1 / 0.1)) {
        p1 = (p1 >> ((18446744073709551615) & 63))
        var v16 float64 = -2.25
        v17 := GenS3{F0: -9223372036854775808, F1: 0}
        _ = v16
        _ = v17
    }
    switch "\n\t" {
    case "hello":
        switch "hello" {
        case "\n\t":
            var v18 uint64 = (^(p1))
            _ = v18
        case "hello":
            if false {
                fmt.Println(p0, p2, p1)
                p1 += p1
            }
        }
    case "foo bar":
        p2 += (p2 << ((172) & 63))
    case "":
        v19 := []int{(^(p2))}
        _ = v19
    default:
        v20 := []int{p2, p2}
        _ = v20
    }
    var v21 int8 = 127
    v22 := []int{p2, (p2 % ((p2) | 1))}
    v23 := GenS1{F0: v21, F1: uint(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    _ = v23
    return 882
}

func main() {
    var v1 int16 = -1
    v2 := []int{9223372036854775807, 32767, 32767, -949}
    v3 := []string{"foo bar"}
    v4 := make(map[string]int)
    if len(v2) > 0 {
        v2[0] = 128
    }
    var v24 string = ("" + "")
    v24 = v24
    var v25 int = -795
    v1 -= v1
    if (2 >= 42) {
        for fi26 := 0; fi26 < 10; fi26++ {
            v25 += v25
            var v27 uint16 = 926
            _ = fi26
            _ = v27
        }
        var v28 uint = 811
        var v29 uint = v28
        _ = v28
        _ = v29
    } else {
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 bool = true
            _ = fi30
            _ = v31
        }
        if true {
            v1 = -32768
            v1 -= -179
            var v32 uint16 = 42
            _ = v32
        }
        v33 := GenS3{F0: 1, F1: 0}
        _ = v33
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        v4[((v24 + ((v24 + v24) + (v24 + "a"))) + v24)] = v25
        v2 = append(v2, (fi34 / ((v25) | 1)))
        _ = fi34
    }
    v25 -= (v25 + v25)
    if false {
        fmt.Println(v1, v1, v24)
        v24 = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
