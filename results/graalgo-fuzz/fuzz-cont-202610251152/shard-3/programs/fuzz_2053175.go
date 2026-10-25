package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS4 struct { F0 string; F1 uint32 }
func genFn0() string {
    if false {
        var v6 int64 = 1
        var v7 uint32 = 1
        v6 = v6
        _ = v6
        _ = v7
    } else {
        var v8 string = "hello"
        _ = v8
    }
    var v9 string = ""
    v9 = ((v9 + "hello") + ("foo bar" + ("\n\t" + v9)))
    _ = v9
    return v9
}
func genFn1(p0 bool) uint64 {
    v10 := GenS1{F0: p0}
    switch 32768 {
    case 42:
        fmt.Println(p0)
    }
    var v11 int32 = 2147483647
    _ = p0
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    var v1 int16 = 32767
    v2 := []int{32768, -9223372036854775808, 127, -1}
    v3 := []uint32{2147483648}
    v4 := make(map[string]string)
    v5 := make(map[string]bool)
    v1 = 1
    var v12 string = "hello"
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 float32 = 0.1
        _ = len(v3)
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 0
    fmt.Println(v12)
    v1 = (-(v1))
    if (!(-1 > -128)) {
        switch 9223372036854775807 {
        case -327:
            if (((1.5 * (float32((((float64(1.5) / 3.14159) + 1e10) + 1e10)) / 1.0)) != (float32(0.1) + 0.0)) || false) {
                var v16 int = -790
                if true {
                    v5[v12] = (209 < 42)
                } else {
                    if len(v2) > 0 {
                        v2[0] = (v16 & (v16 / ((9223372036854775807) | 1)))
                    }
                }
                v15 -= (v15 + (v15 << ((2) & 63)))
                _ = v16
            } else {
                var v17 int64 = -792
                v1 -= v1
                _ = v17
            }
        case 4294967295:
            if (v15 != v15) {
                v12 = ((("" + v12) + "") + ("\n\t" + (v12 + v12)))
            }
        case 42:
            v2 = append(v2, 127)
        default:
            for fi18 := 0; fi18 < 12; fi18++ {
                delete(v5, "a")
                _ = fi18
            }
        }
    } else {
        for fi19 := 0; fi19 < 5; fi19++ {
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 float32 = (3.14159 + 0.1)
                var v22 uint64 = 0
                v2 = append(v2, (fi19 + 0))
                _ = fi20
                _ = v21
                _ = v22
            }
            fmt.Println(fi19, v1, v15)
            for fi23 := 0; fi23 < 15; fi23++ {
                if true {
                    if true {
                        v12 = "hello"
                        switch fi19 {
                        case 42:
                            v24 := GenS1{F0: (-2.25 != float64(3.14159))}
                            _ = v24
                        case 403:
                            v25 := GenS2{F0: 42, F1: (-(fi19))}
                            _ = v25
                        case 0:
                            v12 = "a"
                        }
                    }
                }
                var v26 uint = 601
                _ = fi23
                _ = v26
            }
            _ = fi19
        }
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := GenS3{F0: 0, F1: true, F2: 255}
        var v29 uint64 = 1
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := GenS2{F0: -607, F1: 4294967295}
        fmt.Println(fi30, fi30, v15)
        if (!false) {
            if false {
                var v32 int8 = 105
                _ = v32
            }
            for fi33 := 0; fi33 < 14; fi33++ {
                v15 -= v15
                v1 = v1
                var v34 uint64 = 18446744073709551615
                _ = fi33
                _ = v34
            }
            v35 := GenS4{F0: (((v12 + v12) + "a") + v12), F1: 810}
            _ = v35
        } else {
            if (2147483648 <= 946) {
                var v36 uint8 = 0
                v15 = v15
                _ = v36
            }
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
