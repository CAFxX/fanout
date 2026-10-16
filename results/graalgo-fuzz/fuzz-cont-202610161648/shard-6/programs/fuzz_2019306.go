package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 int, p1 uint8) uint16 {
    v5 := []int{(p0 & (p0 | 128)), (p0 >> ((4294967295) & 63)), p0}
    var v6 uint32 = 519
    if (1.5 > (1.5 / 1e10)) {
        fmt.Println(p1, p0)
        switch "" {
        case "foo bar":
            var v7 bool = false
            _ = v7
        default:
            var v8 uint16 = 1
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 487
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{4294967295}
    var v3 uint8 = 128
    var v4 bool = (2147483647 < -2147483648)
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int32 = 2147483647
        _ = fi9
        _ = v10
    }
    if ((660 > 255) || (-937 > 218)) {
        fmt.Println(v3, v3)
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 float32 = (1e10 + float32((float64(((0.0 - (-2.25 - (float32(1.5) * float32(0.1)))) + (1e10 - 0.0))) + 1e10)))
            fmt.Println(fi11)
            _ = fi11
            _ = v12
        }
        var v13 uint = 255
        _ = v13
    } else {
        v3 += (^(v3))
        var v14 float32 = ((0.1 * 0.1) - 3.14159)
        _ = v14
    }
    var v15 int16 = 104
    v4 = ((1 >= 42) || ((v3 | (v3 % ((v3) | 1))) < v3))
    v1[301] = (false && (v4 || (100 == 1)))
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(v3)
        var v17 int32 = 1
        fmt.Println(v3, v17)
        _ = fi16
        _ = v17
    }
    v4 = ((v3 | v3) >= (^(v3)))
    v1[117] = (v4 || false)
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS1{F0: 2147483647, F1: 42}
        switch "hello" {
        case "hello":
            v4 = (!true)
        case "a":
            var v20 int = (^(fi18))
            _ = v20
        default:
            if v4 {
                switch (fi18 + fi18) {
                case 951:
                    v4 = false
                case 1:
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                case -176:
                    v15 += 0
                }
            } else {
                switch (fi18 & (fi18 ^ fi18)) {
                case 1:
                    var v21 uint = 62
                    _ = v21
                case 32767:
                    if (!false) {
                        for fi22 := 0; fi22 < 18; fi22++ {
                            switch fi22 {
                            case 9223372036854775807:
                                v23 := GenS1{F0: 268, F1: 42}
                                _ = v23
                            case 312:
                                if (v3 == v3) {
                                    var v24 int64 = 521
                                    v15 = 1
                                    _ = v24
                                } else {
                                    v15 = v15
                                    v25 := GenS1{F0: 2, F1: 100}
                                    var v26 uint16 = 0
                                    _ = v25
                                    _ = v26
                                }
                            case 128:
                                v4 = (!(((v15 ^ v15) <= v15) || (1.5 > 0.0)))
                            default:
                                var v27 uint = 18446744073709551615
                                _ = v27
                            }
                            _ = fi22
                        }
                        for fi28 := 0; fi28 < 2; fi28++ {
                            if true {
                                var v29 int = fi18
                                fmt.Println(fi18, fi18)
                                _ = v29
                            } else {
                                _ = len(v2)
                            }
                            _ = fi28
                        }
                    }
                case -706:
                    v3 = v3
                }
            }
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v4)
    for fi30 := 0; fi30 < 15; fi30++ {
        v3 += v3
        for fi31 := 0; fi31 < 9; fi31++ {
            v4 = false
            if (("foo bar" <= (("hello" + "a") + "a")) || v4) {
                v4 = true
                var v32 uint = 229
                v3 -= v3
                _ = v32
            } else {
                fmt.Println(fi31, fi30)
            }
            _ = fi31
        }
        var v33 float32 = 1e10
        _ = fi30
        _ = v33
    }
    v15 -= v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
