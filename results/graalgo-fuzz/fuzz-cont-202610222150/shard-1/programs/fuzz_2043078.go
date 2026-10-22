package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint32; F1 int64; F2 int32 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint64 }
func genFn0(p0 int64, p1 uint32, p2 int) uint16 {
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = 42
        p2 = (p2 & p2)
        _ = fi7
    }
    v8 := GenS1{F0: 4294967295}
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS2{F0: p1, F1: false}
            var v12 uint16 = 42
            var v13 int8 = 127
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 2
    var v3 uint64 = 194
    var v4 float32 = -2.25
    var v5 float32 = v4
    v6 := []uint32{2147483648, 42}
    v6 = append(v6, 2147483648)
    var v14 int = 2147483647
    if false {
        for fi15 := 0; fi15 < 10; fi15++ {
            delete(v1, 2147483647)
            _ = fi15
        }
        var v16 uint64 = (^(v3))
        _ = v16
    } else {
        if (true || false) {
            switch v14 {
            case 0:
                var v17 uint16 = v2
                _ = v17
            default:
                v1[2147483648] = "hello"
            }
            if (1 < -9223372036854775808) {
                for fi18 := 0; fi18 < 2; fi18++ {
                    v5 = (float32((-2.25 + float64(float32(((0.0 / 1e10) + (1.5 + -2.25)))))) / 1e10)
                    v19 := GenS3{F0: 1, F1: (1000 > 0)}
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 14; fi20++ {
                    v3 += uint64(v2)
                    _ = fi20
                }
            } else {
                if true {
                    v21 := GenS4{F0: 879, F1: -1, F2: 69}
                    v5 += 0.0
                    _ = v21
                }
                for fi22 := 0; fi22 < 1; fi22++ {
                    v2 = 591
                    if (-48 != -1) {
                        if false {
                            var v23 uint = 42
                            for fi24 := 0; fi24 < 15; fi24++ {
                                _ = len(v6)
                                v25 := GenS5{F0: 2, F1: 127, F2: (v3 / (((v3 >> ((v23) & 63))) | 1))}
                                _ = fi24
                                _ = v25
                            }
                            _ = len(v6)
                            _ = v23
                        } else {
                            v2 = 1
                        }
                        v5 = (v5 + 1e10)
                        _ = len(v1)
                    } else {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            var v27 bool = true
                            var v28 uint8 = 199
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        fmt.Println(v2, fi22, v2)
                    }
                    _ = fi22
                }
            }
            var v29 uint = 18446744073709551615
            _ = v29
        } else {
            var v30 bool = ((!(-546 == 9223372036854775807)) || ((v14 ^ 127) >= v14))
            _ = v30
        }
    }
    fmt.Println(v5, v5)
    for fi31 := 0; fi31 < 8; fi31++ {
        v3 = v3
        v6 = append(v6, 2147483648)
        _ = fi31
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v2 = v2
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 uint16 = (^(v2))
        _ = fi32
        _ = v33
    }
    var v34 float32 = 0.1
    var v35 bool = (((("a" + "") + "hello") >= (("a" + "hello") + (("hello" + "a") + "hello"))) && false)
    fmt.Println(v34, v2, v4)
    var v36 uint64 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
