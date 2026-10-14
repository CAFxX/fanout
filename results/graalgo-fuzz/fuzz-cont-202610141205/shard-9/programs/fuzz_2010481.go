package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int8; F2 int64 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint, p1 bool) uint64 {
    switch "hello" {
    case "hello":
        var v7 int = 9223372036854775807
        _ = v7
    case "a":
        v8 := GenS1{F0: 0, F1: 65535, F2: -1}
        _ = v8
    case "foo bar":
        fmt.Println(p0)
    default:
        var v9 bool = (616 > 0)
        _ = v9
    }
    p0 = p0
    var v10 uint = 4294967295
    if false {
        for fi11 := 0; fi11 < 16; fi11++ {
            for fi12 := 0; fi12 < 20; fi12++ {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(fi11, v10, fi11)
                }
                v10 -= v10
                if false {
                    fmt.Println(p1, p1, p1)
                }
                _ = fi12
            }
            _ = fi11
        }
        v10 += (p0 % ((817) | 1))
    }
    _ = p0
    _ = p1
    _ = v10
    return 737
}

func main() {
    var v1 uint = 46
    v2 := make(map[int]string)
    v3 := []int{0, -1, 9223372036854775807, -650}
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    var v6 int8 = -19
    fmt.Println(v6, v1)
    v1 = 391
    for fi13 := 0; fi13 < 1; fi13++ {
        if true {
            v14 := GenS2{F0: (fi13 / ((fi13) | 1))}
            if (!false) {
                for fi15 := 0; fi15 < 15; fi15++ {
                    if false {
                        var v16 bool = (!(860 >= 32767))
                        _ = v16
                    }
                    delete(v5, 742)
                    _ = fi15
                }
            } else {
                if false {
                    var v17 uint32 = 4294967295
                    var v18 string = "hello"
                    delete(v5, (^(v17)))
                    _ = v17
                    _ = v18
                } else {
                    switch fi13 {
                    case 127:
                        var v19 string = ("a" + "foo bar")
                        _ = v19
                    default:
                        v6 = 0
                    }
                    for fi20 := 0; fi20 < 12; fi20++ {
                        fmt.Println(v6)
                        v6 -= (v6 << ((18446744073709551615) & 63))
                        _ = fi20
                    }
                    switch (fi13 % (((fi13 << ((788) & 63))) | 1)) {
                    case 42:
                        if (0 > 0) {
                            if ((0 < 0) || false) {
                                var v21 bool = true
                                _ = v21
                            }
                        }
                    default:
                        var v22 uint64 = 18446744073709551615
                        _ = v22
                    }
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            for fi23 := 0; fi23 < 11; fi23++ {
                if false {
                    v6 += (v6 << (((v1 % ((v1) | 1))) & 63))
                    v4["a"] = 775
                    v24 := GenS3{F0: 127, F1: v6, F2: 1}
                    _ = v24
                } else {
                    var v25 float64 = float64((1.5 - ((3.14159 / (3.14159 * (-2.25 - 0.0))) * 1e10)))
                    var v26 int16 = 0
                    switch (fi13 * (fi23 - 536)) {
                    case 0:
                        if (1e10 > 1e10) {
                            if (false && (275 != 721)) {
                                switch fi23 {
                                case 9223372036854775807:
                                    v4[("a" + "a")] = 0
                                case 94:
                                    v26 -= (v26 ^ v26)
                                case -2147483648:
                                    if len(v3) > 0 {
                                        v3[0] = 274
                                    }
                                }
                            }
                            var v27 int64 = -9223372036854775808
                            _ = v27
                        }
                    case 578:
                        v26 = v26
                    case 32767:
                        var v28 float32 = (((1.5 + 1e10) + 3.14159) / 1e10)
                        _ = v28
                    }
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
            _ = v14
        } else {
            fmt.Println(fi13, v1, v6)
        }
        _ = fi13
    }
    v29 := GenS4{F0: 42, F1: 18446744073709551615, F2: 246}
    var v30 uint64 = 18446744073709551615
    v31 := GenS5{F0: 339, F1: 65535}
    _ = len(v2)
    fmt.Println(v30)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
