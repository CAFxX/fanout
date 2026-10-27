package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
func genFn0(p0 uint16, p1 int64, p2 int64) int8 {
    p2 -= p2
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: 0, F1: 2147483647}
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 float32 = 1.5
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return -128
}
func genFn1(p0 uint64) uint16 {
    p0 = (p0 / ((p0) | 1))
    p0 -= (p0 & (p0 ^ p0))
    v9 := []int{}
    for fi10 := 0; fi10 < 5; fi10++ {
        if len(v9) > 0 {
            v9[0] = fi10
        }
        _ = fi10
    }
    _ = p0
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    var v4 bool = false
    if (65535 < 65535) {
        v11 := GenS1{F0: (v1 % (((v1 % ((42) | 1))) | 1)), F1: 32767}
        _ = v11
    } else {
        fmt.Println(v1)
        switch 4294967295 {
        case 42:
            delete(v2, "")
        case -1:
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 int32 = 42
                _ = fi12
                _ = v13
            }
        case 0:
            v1 = v1
        default:
            v14 := GenS1{F0: v1, F1: -772}
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := v2["hello"]
        _ = fi15
        _ = v16
    }
    var v17 uint16 = 255
    v18 := GenS1{F0: v1, F1: 32768}
    var v19 uint = 255
    fmt.Println(v4)
    v4 = (((v17 & 611) > 329) && (!true))
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            v22 := []int{2147483647, (fi21 << ((4294967295) & 63))}
            if ((0.1 / (3.14159 - (0.1 + 3.14159))) >= -2.25) {
                if v4 {
                    v19 -= v19
                    var v23 string = "hello"
                    fmt.Println(fi20)
                    _ = v23
                } else {
                    v3[2147483648] = 813
                    v3[468] = 2147483647
                }
                v4 = v4
            } else {
                var v24 uint = (v19 & (v19 % (((v19 | v19)) | 1)))
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        v25 := GenS1{F0: 255, F1: fi20}
        _ = fi20
        _ = v25
    }
    switch 851 {
    case 9223372036854775807:
        var v26 string = "\n\t"
        _ = v26
    case -2147483648:
        v1 = 42
    case 0:
        var v27 int = -1
        _ = v27
    }
    v28 := []int{42, 143, 2147483647}
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 int32 = -426
            v32 := GenS1{F0: v1, F1: fi29}
            _ = fi30
            _ = v31
            _ = v32
        }
        for fi33 := 0; fi33 < 3; fi33++ {
            v17 -= 930
            _ = fi33
        }
        if v4 {
            v4 = (v4 && (1 <= 0))
            if false {
                v19 += (^(v19))
                v19 = (v19 / ((v19) | 1))
            } else {
                var v34 uint8 = v1
                _ = v34
            }
            if len(v28) > 0 {
                v28[0] = fi29
            }
        } else {
            fmt.Println(v4, v19)
        }
        _ = fi29
    }
    v1 += (v1 >> (((v19 >> ((v19) & 63))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
