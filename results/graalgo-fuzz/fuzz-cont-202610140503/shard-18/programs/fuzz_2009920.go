package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint8; F2 int }
func genFn0(p0 uint64) uint64 {
    var v7 bool = false
    fmt.Println(p0)
    v7 = v7
    _ = p0
    _ = v7
    return p0
}
func genFn1(p0 int) uint64 {
    if (((3.14159 / ((0.0 / (0.0 + 1e10)) + 0.1)) + (1.5 + 1.5)) > (0.1 / (0.1 * (0.0 - -2.25)))) {
        fmt.Println(p0)
        v8 := GenS1{F0: 858}
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        p0 = -9223372036854775808
        _ = fi9
    }
    _ = p0
    return 18446744073709551615
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]bool)
    var v3 int64 = 0
    var v4 int8 = -99
    v5 := make(map[int]int)
    v6 := make(map[uint32]uint32)
    v1 = (v1 / (((v1 * 0)) | 1))
    v10 := []int{42, 1, 32768, 9223372036854775807}
    _ = len(v10)
    switch ("a" + "\n\t") {
    case "hello":
        var v11 int32 = v1
        _ = v11
    case "\n\t":
        switch 128 {
        case 676:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        case 32767:
            if len(v10) > 0 {
                v10[0] = -302
            }
        default:
            if false {
                for fi12 := 0; fi12 < 15; fi12++ {
                    if (true && ((v1 < v1) || (0 <= 4294967295))) {
                        switch (fi12 | (-(fi12))) {
                        case 623:
                            if (3.14159 >= 3.14159) {
                                v3 = (v3 << ((380) & 63))
                                for fi13 := 0; fi13 < 15; fi13++ {
                                    var v14 string = "foo bar"
                                    if (42 == 1) {
                                        var v15 int = (fi13 ^ (fi13 ^ (fi13 << ((18446744073709551615) & 63))))
                                        _ = len(v10)
                                        _ = v15
                                    }
                                    _ = fi13
                                    _ = v14
                                }
                                var v16 uint64 = genFn1((fi12 | (fi12 ^ (^(fi12)))))
                                _ = v16
                            } else {
                                if ((493 == 391) || false) {
                                    for fi17 := 0; fi17 < 7; fi17++ {
                                        v6[39] = 2147483647
                                        if (true && (false && false)) {
                                            fmt.Println(v1, v3, fi12)
                                        }
                                        _ = fi17
                                    }
                                    var v18 uint64 = 2
                                    _ = v18
                                }
                            }
                        case 128:
                            v19 := GenS2{F0: 136}
                            _ = v19
                        default:
                            v1 = v1
                        }
                        for fi20 := 0; fi20 < 14; fi20++ {
                            fmt.Println(fi20)
                            v10 = append(v10, fi20)
                            var v21 float64 = 3.14159
                            _ = fi20
                            _ = v21
                        }
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    }
                    var v22 int32 = -1
                    _ = fi12
                    _ = v22
                }
                v3 = (v3 / (((^(v3))) | 1))
                fmt.Println(v4, v1)
            }
        }
    case "foo bar":
        for fi23 := 0; fi23 < 3; fi23++ {
            for fi24 := 0; fi24 < 3; fi24++ {
                v10 = append(v10, 128)
                v3 = 0
                switch "hello" {
                case "foo bar":
                    if len(v10) > 0 {
                        v10[0] = int(v4)
                    }
                case "a":
                    fmt.Println(v1)
                case "":
                    v4 += v4
                default:
                    fmt.Println(fi24, fi24)
                }
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi24
            }
            switch 32768 {
            case 4294967295:
                v25 := GenS2{F0: 2}
                _ = v25
            }
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 int64 = v3
                _ = fi26
                _ = v27
            }
            _ = fi23
        }
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        fmt.Println(v3, v1)
        _ = len(v5)
        for fi29 := 0; fi29 < 19; fi29++ {
            fmt.Println(v3, v3, v3)
            _ = fi29
        }
        _ = fi28
    }
    v30 := GenS1{F0: 65535}
    if len(v10) > 0 {
        v10[0] = -1
    }
    v4 += v4
    v31 := GenS3{F0: (v1 / (((v1 << ((4294967295) & 63))) | 1)), F1: 0, F2: -490}
    var v32 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
