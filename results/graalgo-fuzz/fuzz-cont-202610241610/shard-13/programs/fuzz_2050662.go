package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int64; F2 uint8 }
func genFn0(p0 int64, p1 string, p2 int) bool {
    var v7 uint = 2
    v8 := []int{p2, 2147483647, (-(p2))}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (1 < 0)
}
func genFn1() bool {
    v9 := GenS1{F0: 2, F1: -117}
    var v10 int8 = -1
    _ = v10
    _ = v9
    return false
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 int64 = -1
    v4 := make(map[string]bool)
    v5 := []uint32{4294967295, 503, 510, 4294967295}
    var v6 uint = 65535
    var v11 float32 = 1.5
    var v12 bool = true
    switch 4294967295 {
    case -699:
        for fi13 := 0; fi13 < 6; fi13++ {
            if v12 {
                switch ("\n\t" + "hello") {
                case "":
                    v3 += (v3 * v3)
                case "a":
                    for fi14 := 0; fi14 < 6; fi14++ {
                        v11 = (v11 / (v11 + v11))
                        v3 = (v3 * v3)
                        switch ("" + "a") {
                        case "a":
                            if (v3 == (v3 >> ((18446744073709551615) & 63))) {
                                fmt.Println(v11, fi13, v12)
                            }
                        }
                        v11 += v11
                        _ = fi14
                    }
                default:
                    var v15 float32 = v11
                    _ = v15
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v16 := GenS1{F0: 1, F1: 35}
                _ = v16
            } else {
                var v17 int = fi13
                for fi18 := 0; fi18 < 10; fi18++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = len(v5)
                    _ = fi18
                }
                v19 := GenS1{F0: 65535, F1: -6}
                _ = v17
                _ = v19
            }
            fmt.Println(v11, v6, v12)
            _ = fi13
        }
    }
    v5 = append(v5, 65535)
    switch 32768 {
    case 32767:
        var v20 int32 = 0
        _ = v20
    case -9223372036854775808:
        v21 := GenS2{F0: 223}
        _ = v21
    case 399:
        switch 579 {
        case 42:
            v22 := GenS2{F0: 382}
            _ = v22
        default:
            for fi23 := 0; fi23 < 9; fi23++ {
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := GenS1{F0: 65535, F1: 0}
                    switch (fi23 ^ -630) {
                    case 2147483647:
                        var v26 bool = ((1.5 <= (0.0 + (v11 + v11))) && v12)
                        _ = v26
                    case 32767:
                        if (2147483648 < 1) {
                            v6 = (v6 + v6)
                            var v27 uint32 = 1
                            v28 := GenS3{F0: 1, F1: v3, F2: 255}
                            _ = v27
                            _ = v28
                        }
                    case 510:
                        fmt.Println(v12, fi24)
                    }
                    _ = fi24
                    _ = v25
                }
                v11 = ((0.0 - (1e10 / 1e10)) - float32(((0.0 - float64(((1.5 + (v11 / 1.5)) - ((v11 + (v11 / 0.1)) / (v11 - -2.25))))) + float64(0.1))))
                var v29 int32 = 42
                _ = fi23
                _ = v29
            }
        }
    }
    v6 = v6
    fmt.Println(v6, v12)
    v6 -= v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
