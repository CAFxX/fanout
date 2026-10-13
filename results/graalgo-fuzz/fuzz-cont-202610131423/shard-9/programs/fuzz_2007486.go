package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int8; F1 int32; F2 uint32 }
type GenS3 struct { F0 int32 }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        if true {
            for fi8 := 0; fi8 < 4; fi8++ {
                fmt.Println(fi8, fi7)
                var v9 int = (^(fi8))
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 4; fi10++ {
                if false {
                    v11 := GenS1{F0: 127, F1: false}
                    var v12 int8 = 42
                    v12 = v12
                    _ = v11
                    _ = v12
                }
                v13 := []int{(fi10 % ((32768) | 1)), fi7, fi7}
                v14 := GenS1{F0: -128, F1: false}
                _ = fi10
                _ = v13
                _ = v14
            }
        }
        _ = fi7
    }
    var v15 int8 = 0
    v15 = 0
    _ = v15
    return 42
}
func genFn1(p0 int8, p1 bool) int16 {
    var v16 int32 = -2147483648
    v16 = (v16 >> ((959) & 63))
    if (18446744073709551615 != 178) {
        v16 = (v16 << ((0) & 63))
    } else {
        p0 = (p0 & (p0 & 42))
        v17 := GenS1{F0: -1, F1: p1}
        var v18 float32 = ((1e10 - 1.5) + -2.25)
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v16
    return 1
}

func main() {
    var v1 int64 = -9
    v2 := []uint32{0}
    var v3 int64 = (v1 ^ v1)
    v4 := make(map[int]uint32)
    var v5 uint = 42
    var v6 int = 127
    v3 = (^(v1))
    _ = len(v4)
    fmt.Println(v3, v1)
    v2 = append(v2, 1)
    _ = len(v4)
    v4[v6] = uint32(v5)
    if (!true) {
        switch (v6 << ((v5) & 63)) {
        case -2147483648:
            v19 := GenS2{F0: int8(v5), F1: -912, F2: 2147483647}
            _ = v19
        case -968:
            var v20 string = "foo bar"
            _ = v20
        default:
            var v21 uint64 = uint64(v5)
            _ = v21
        }
        if (!((!(!true)) && ((true && (1000 > -32768)) || ((0.1 + 0.0) != (((0.1 - 1.5) - 1.5) + 3.14159))))) {
            if false {
                var v22 uint16 = 2
                switch "a" {
                case "hello":
                    v6 = (v6 - v6)
                default:
                    for fi23 := 0; fi23 < 18; fi23++ {
                        for fi24 := 0; fi24 < 3; fi24++ {
                            var v25 int16 = genFn1(-81, (-128 < -92))
                            if false {
                                var v26 uint = (^(v5))
                                _ = v26
                            } else {
                                var v27 int32 = 42
                                var v28 int8 = 10
                                _ = v27
                                _ = v28
                            }
                            v25 -= v25
                            _ = fi24
                            _ = v25
                        }
                        var v29 uint64 = 42
                        _ = fi23
                        _ = v29
                    }
                }
                _ = v22
            } else {
                var v30 float64 = 3.14159
                _ = v30
            }
            v3 = (^(v3))
            v31 := GenS3{F0: int32(v6)}
            _ = v31
        }
    } else {
        delete(v4, 1)
    }
    var v32 int32 = -545
    fmt.Println(v1, v6)
    v3 -= v1
    v1 -= (^(v1))
    var v33 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
