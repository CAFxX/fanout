package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }
func genFn0(p0 uint8, p1 int8) string {
    if (42 > 0) {
        var v7 uint64 = 1
        if false {
            v8 := []int{-1, -559}
            var v9 int64 = 42
            _ = v8
            _ = v9
        } else {
            p0 -= p0
            var v10 int = -9223372036854775808
            _ = v10
        }
        var v11 int8 = 5
        _ = v11
        _ = v7
    } else {
        p0 = p0
    }
    var v12 uint64 = 4294967295
    if true {
        var v13 int32 = 0
        v14 := []int{}
        for fi15 := 0; fi15 < 9; fi15++ {
            if true {
                var v16 bool = false
                var v17 uint64 = v12
                _ = v16
                _ = v17
            } else {
                for fi18 := 0; fi18 < 4; fi18++ {
                    switch "\n\t" {
                    case "a":
                        v14 = append(v14, fi15)
                    case "foo bar":
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                    default:
                        p0 += (p0 * p0)
                    }
                    _ = fi18
                }
            }
            _ = fi15
        }
        _ = v13
        _ = v14
    } else {
        fmt.Println(p0, v12)
        v12 += (v12 << ((708) & 63))
    }
    fmt.Println(v12, v12, p0)
    v19 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v12
    _ = v19
    return "\n\t"
}

func main() {
    var v1 int32 = -356
    var v2 uint16 = 289
    v3 := make(map[int]string)
    var v4 float64 = ((3.14159 * -2.25) + (0.0 * float64(1.5)))
    var v5 int = 9223372036854775807
    var v6 int = (v5 << ((377) & 63))
    fmt.Println(v2, v5, v1)
    switch v6 {
    case 127:
        switch (v5 / (((v5 + v5)) | 1)) {
        case 32768:
            if true {
                v3[(^(v5))] = "a"
                v2 -= (v2 % ((v2) | 1))
            } else {
                v20 := []int{(v5 << ((255) & 63)), v6, (v5 % ((v6) | 1)), v6}
                _ = v20
            }
        case 9223372036854775807:
            for fi21 := 0; fi21 < 12; fi21++ {
                v22 := GenS2{F0: 1, F1: 40}
                v2 = v2
                _ = fi21
                _ = v22
            }
        case 127:
            delete(v3, (v6 & (v5 << ((24) & 63))))
        default:
            v1 += (v1 + (v1 - (v1 ^ -2147483648)))
        }
    }
    delete(v3, v6)
    var v23 uint16 = (v2 ^ (v2 << ((927) & 63)))
    fmt.Println(v6)
    delete(v3, (^(v5)))
    v4 = 0.0
    for fi24 := 0; fi24 < 6; fi24++ {
        fmt.Println(v1, v2)
        if false {
            v25 := []int{-1, (v6 & 32768), (v5 | v6)}
            _ = v25
        } else {
            v26 := GenS3{F0: 282, F1: (((("\n\t" + "foo bar") + "foo bar") + "\n\t") + "\n\t"), F2: v23}
            v6 = (v5 * fi24)
            var v27 uint8 = 0
            _ = v26
            _ = v27
        }
        _ = fi24
    }
    v28 := []int{(v6 - v5), (v5 * (v5 << ((670) & 63)))}
    var v29 int8 = 75
    var v30 int32 = (v1 / ((440) | 1))
    fmt.Println(v29, v2, v30)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
