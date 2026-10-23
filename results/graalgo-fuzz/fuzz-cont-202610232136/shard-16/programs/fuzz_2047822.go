package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int }
type GenS4 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint64, p2 int) uint {
    v5 := GenS1{F0: 186, F1: 65535}
    var v6 int32 = 227
    switch "\n\t" {
    case "a":
        v7 := []int{p2}
        _ = v7
    case "hello":
        if p0 {
            var v8 float32 = 0.0
            _ = v8
        }
    case "foo bar":
        var v9 uint16 = 961
        _ = v9
    default:
        v10 := GenS2{F0: 1, F1: 2147483648, F2: 4294967295}
        _ = v10
    }
    var v11 bool = ((-(p2)) >= 32768)
    fmt.Println(p0, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []uint32{2147483648, 2, 4294967295, 4294967295}
    var v2 int8 = -86
    v3 := make(map[int]uint32)
    var v4 int8 = (v2 | (v2 ^ (-(v2))))
    var v12 float64 = (1e10 / (1e10 + 1e10))
    if (float64(1e10) == float64(-2.25)) {
        if (42 > 4294967295) {
            v3[737] = 2
            for fi13 := 0; fi13 < 11; fi13++ {
                _ = len(v3)
                _ = fi13
            }
        }
        v2 = 1
        for fi14 := 0; fi14 < 1; fi14++ {
            switch 896 {
            case 1:
                fmt.Println(v4, fi14, v4)
            case 32767:
                var v15 uint = 204
                _ = v15
            }
            _ = fi14
        }
    } else {
        var v16 int = 32767
        switch v16 {
        case 9223372036854775807:
            v16 = 127
        case 0:
            var v17 int8 = v4
            _ = v17
        case -1:
            if (1 != 128) {
                if (388 > 331) {
                    switch (v16 ^ (v16 ^ v16)) {
                    case -988:
                        if (471 != -904) {
                            v18 := GenS3{F0: 633, F1: 2, F2: -487}
                            _ = v18
                        } else {
                            var v19 uint64 = 18446744073709551615
                            _ = v19
                        }
                    case -9223372036854775808:
                        var v20 int8 = v4
                        _ = v20
                    }
                    var v21 uint = 18446744073709551615
                    _ = v21
                } else {
                    var v22 int = (v16 >> ((genFn0(false, 0, v16)) & 63))
                    _ = len(v1)
                    v23 := GenS3{F0: 831, F1: 523, F2: (^(v16))}
                    _ = v22
                    _ = v23
                }
                for fi24 := 0; fi24 < 3; fi24++ {
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    if false {
                        switch -9223372036854775808 {
                        case -1:
                            v4 -= (-(v4))
                        case -9223372036854775808:
                            fmt.Println(v12, v16, v12)
                        case 822:
                            if false {
                                var v25 float64 = 3.14159
                                v26 := v3[(^(v16))]
                                _ = v25
                                _ = v26
                            }
                        default:
                            v2 = (v2 & 0)
                        }
                        v4 += 1
                    }
                    _ = fi24
                }
            }
        default:
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        }
        _ = v16
    }
    if false {
        var v27 string = "hello"
        _ = v27
    }
    v28 := GenS4{F0: 222}
    for fi29 := 0; fi29 < 3; fi29++ {
        v12 = -2.25
        var v30 int32 = 0
        _ = fi29
        _ = v30
    }
    var v31 int32 = -154
    v1 = append(v1, 1)
    var v32 int64 = 360
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
