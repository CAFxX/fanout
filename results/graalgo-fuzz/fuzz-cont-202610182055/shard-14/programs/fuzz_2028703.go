package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 bool }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint32 }
func genFn0() int32 {
    switch 4294967295 {
    case 128:
        v5 := []int{0, -946, 769}
        _ = v5
    case 42:
        var v6 bool = (true || true)
        _ = v6
    }
    v7 := []int{4294967295, -708, 334}
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 float64 = (((1.5 / 1.0) - 0.0) / -2.25)
        _ = fi8
        _ = v9
    }
    var v10 bool = (128 == 0)
    _ = v10
    _ = v7
    return -2147483648
}
func genFn1(p0 int32, p1 string, p2 string) uint {
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS1{F0: 906, F1: (-(p0)), F2: false}
        v13 := GenS2{F0: -1, F1: 18446744073709551615}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 string = (p2 + (("foo bar" + p1) + p2))
    var v15 int8 = -22
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 162
}

func main() {
    var v1 string = "foo bar"
    var v2 float64 = 1.5
    v3 := []string{v1, (((("" + "hello") + "foo bar") + (("a" + "a") + (v1 + v1))) + (v1 + v1))}
    var v4 int8 = 127
    v1 = v1
    fmt.Println(v2)
    v3 = append(v3, v1)
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 float64 = v2
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v17 -= 3.14159
            v1 = v1
        }
        _ = fi16
        _ = v17
    }
    var v18 uint32 = 0
    v3 = append(v3, v1)
    fmt.Println(v4, v18, v2)
    if len(v3) > 0 {
        v3[0] = ("" + (v1 + v1))
    }
    var v19 uint8 = 42
    v18 += (v18 + (^(v18)))
    v20 := GenS1{F0: v18, F1: 123, F2: (!(!true))}
    var v21 uint = 18446744073709551615
    for fi22 := 0; fi22 < 6; fi22++ {
        v1 = v1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        v24 := GenS3{F0: v18}
        _ = fi23
        _ = v24
    }
    v1 = v1
    if true {
        fmt.Println(v19, v18, v21)
        _ = len(v3)
    } else {
        var v25 int64 = int64(v4)
        if (!false) {
            v18 += (v18 % (((v18 << (((^(v21))) & 63))) | 1))
            for fi26 := 0; fi26 < 15; fi26++ {
                if false {
                    v27 := GenS3{F0: 948}
                    for fi28 := 0; fi28 < 18; fi28++ {
                        switch 32768 {
                        case 2147483647:
                            var v29 float32 = 1.5
                            _ = v29
                        default:
                            v25 -= 315
                        }
                        _ = fi28
                    }
                    fmt.Println(v19, v4)
                    _ = v27
                }
                _ = fi26
            }
            var v30 int = -394
            _ = v30
        }
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
