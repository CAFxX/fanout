package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int64; F2 int32 }
type GenS3 struct { F0 string; F1 uint64; F2 int32 }
func genFn0(p0 uint32, p1 int32, p2 uint) int8 {
    v7 := GenS1{F0: 887}
    if (127 >= uint8(p2)) {
        var v8 uint32 = (p0 ^ p0)
        _ = v8
    }
    var v9 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return -25
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 int = -184
    v4 := make(map[string]uint32)
    var v5 string = "a"
    v6 := make(map[uint32]bool)
    v10 := []int{v3, v3}
    var v11 uint64 = 1
    for fi12 := 0; fi12 < 12; fi12++ {
        fmt.Println(v5, v5, v11)
        v13 := GenS1{F0: 18446744073709551615}
        _ = fi12
        _ = v13
    }
    var v14 int16 = 501
    v5 = (v5 + "hello")
    switch v3 {
    case 531:
        if ((^(v11)) < v11) {
            switch v5 {
            case "":
                switch v5 {
                case "foo bar":
                    for fi15 := 0; fi15 < 7; fi15++ {
                        for fi16 := 0; fi16 < 19; fi16++ {
                            var v17 string = v5
                            var v18 int8 = 100
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        var v19 int32 = -194
                        v20 := GenS1{F0: (v11 << ((377) & 63))}
                        _ = fi15
                        _ = v19
                        _ = v20
                    }
                case "hello":
                    if ((!((^(v11)) != v11)) || false) {
                        var v21 int64 = -9223372036854775808
                        if true {
                            v2[-300] = v3
                            fmt.Println(v11, v14, v21)
                        }
                        _ = len(v4)
                        _ = v21
                    }
                default:
                    var v22 uint32 = 65535
                    _ = v22
                }
            default:
                var v23 int16 = (v14 + v14)
                _ = v23
            }
            if (!(-2.25 < -2.25)) {
                var v24 uint64 = (v11 % ((4294967295) | 1))
                if (-498 != -706) {
                    switch (v3 - v3) {
                    case 4294967295:
                        v25 := GenS2{F0: v5, F1: 302, F2: 1}
                        _ = v25
                    case 42:
                        var v26 int8 = 127
                        _ = v26
                    }
                } else {
                    fmt.Println(v5)
                    fmt.Println(v14, v3, v11)
                }
                _ = len(v2)
                _ = v24
            } else {
                var v27 float64 = (0.1 * ((float64(float32((((3.14159 / (3.14159 + 1.5)) + 1e10) * ((1e10 / 1.5) - 0.1)))) / -2.25) - -2.25))
                v6[2147483647] = (v5 < v5)
                fmt.Println(v3)
                _ = v27
            }
            v14 = v14
        }
    case -1:
        v10 = append(v10, int(v11))
    case -2147483648:
        v28 := GenS3{F0: "", F1: v11, F2: 2147483647}
        _ = v28
    }
    var v29 int16 = v14
    if (18446744073709551615 != 697) {
        var v30 string = "foo bar"
        _ = v30
    } else {
        var v31 uint32 = 539
        _ = v31
    }
    var v32 uint8 = 128
    v5 = ("" + "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
