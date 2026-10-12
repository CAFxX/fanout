package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 string; F1 string; F2 string }
type GenS3 struct { F0 int32 }
func genFn0(p0 int16, p1 uint32, p2 uint8) int {
    fmt.Println(p0)
    v7 := []int{-153}
    if false {
        if false {
            for fi8 := 0; fi8 < 16; fi8++ {
                if false {
                    switch (^(fi8)) {
                    case 42:
                        if (p2 == 2) {
                            var v9 bool = (!true)
                            if len(v7) > 0 {
                                v7[0] = (fi8 ^ fi8)
                            }
                            v7 = append(v7, fi8)
                            _ = v9
                        } else {
                            for fi10 := 0; fi10 < 8; fi10++ {
                                var v11 int16 = (p0 + p0)
                                fmt.Println(v11, fi10)
                                _ = fi10
                                _ = v11
                            }
                        }
                    default:
                        var v12 float64 = 0.0
                        _ = v12
                    }
                    v7 = append(v7, (fi8 * fi8))
                    var v13 uint = 65535
                    _ = v13
                } else {
                    var v14 int = fi8
                    fmt.Println(p0)
                    _ = v14
                }
                v15 := GenS1{F0: p0, F1: 127}
                _ = fi8
                _ = v15
            }
        }
        if (3.14159 != (float64(0.0) / -2.25)) {
            v16 := GenS1{F0: (p0 << ((1) & 63)), F1: -59}
            for fi17 := 0; fi17 < 17; fi17++ {
                v7 = append(v7, fi17)
                fmt.Println(fi17, p2)
                _ = fi17
            }
            _ = v16
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        p0 = (p0 * (-(p0)))
        v18 := GenS1{F0: p0, F1: 127}
        _ = v18
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483647
}
func genFn1() uint64 {
    if (257 <= 900) {
        var v19 uint32 = 406
        _ = v19
    }
    v20 := GenS2{F0: ("a" + ("\n\t" + "a")), F1: ((("\n\t" + "") + ("hello" + "hello")) + "foo bar"), F2: ""}
    _ = v20
    return 4294967295
}

func main() {
    var v1 uint8 = 127
    var v2 int = 32767
    v3 := []int{v2}
    v4 := make(map[string]string)
    var v5 bool = true
    v6 := []uint32{4294967295, 65535}
    if v5 {
        v2 = genFn0(1, 2, (v1 ^ (v1 % ((v1) | 1))))
    } else {
        delete(v4, "a")
        var v21 int32 = 0
        if v5 {
            fmt.Println(v2)
            v21 -= v21
            v22 := GenS3{F0: -2147483648}
            _ = v22
        } else {
            v1 = (v1 - (v1 | v1))
            if true {
                v2 = 339
                v1 = 108
                switch ("" + "a") {
                case "a":
                    fmt.Println(v1)
                default:
                    var v23 int64 = 1
                    _ = v23
                }
            } else {
                for fi24 := 0; fi24 < 18; fi24++ {
                    fmt.Println(v2, v2)
                    _ = fi24
                }
            }
            var v25 float32 = float32(0.0)
            _ = v25
        }
        _ = v21
    }
    var v26 uint32 = 42
    v6 = append(v6, 788)
    var v27 float32 = 3.14159
    var v28 int8 = -128
    fmt.Println(v2)
    fmt.Println(v27, v5, v28)
    if v5 {
        fmt.Println(v1, v28)
        var v29 int = (^(v2))
        _ = v29
    } else {
        var v30 int = genFn0(145, 484, (v1 - v1))
        _ = v30
    }
    var v31 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
