package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
func genFn0(p0 int64) int32 {
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{(^(fi7)), (fi7 % ((fi7) | 1)), fi7, (fi7 >> ((255) & 63))}
        p0 = p0
        v9 := GenS1{F0: true, F1: 0}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        p0 = -9223372036854775808
        _ = fi10
    }
    var v11 float64 = 0.1
    _ = p0
    _ = v11
    return -1
}

func main() {
    var v1 uint64 = 28
    var v2 uint32 = 4294967295
    var v3 int16 = 294
    v4 := []uint32{v2}
    v5 := []string{"\n\t", "", ((("" + (((("a" + "a") + "foo bar") + "\n\t") + "a")) + "\n\t") + ("foo bar" + "foo bar")), "a"}
    var v6 int64 = 930
    if false {
        v2 -= v2
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        switch fi12 {
        case -2147483648:
            if true {
                v13 := GenS1{F0: (230 < v2), F1: v1}
                var v14 uint32 = (v2 | v2)
                _ = v13
                _ = v14
            }
        case 1:
            switch (fi12 / ((fi12) | 1)) {
            case 2147483647:
                for fi15 := 0; fi15 < 17; fi15++ {
                    v2 += 2
                    switch (fi12 | (fi12 + 480)) {
                    case 128:
                        switch 128 {
                        case -1:
                            var v16 int8 = 127
                            _ = v16
                        case 9223372036854775807:
                            v6 -= -1
                            _ = len(v5)
                        case -810:
                            v17 := GenS2{F0: "\n\t"}
                            _ = v17
                        }
                    default:
                        if len(v4) > 0 {
                            v4[0] = v2
                        }
                    }
                    var v18 bool = ((!(91 >= 18446744073709551615)) || (725 >= 328))
                    _ = fi15
                    _ = v18
                }
            case 128:
                var v19 bool = false
                _ = v19
            case 4294967295:
                v2 += (v2 ^ (v2 + 2147483647))
            }
        case 9223372036854775807:
            v3 = (v3 * 42)
        default:
            _ = len(v4)
        }
        _ = fi12
    }
    var v20 uint32 = (v2 >> ((42) & 63))
    v6 -= (v6 & v6)
    var v21 uint = 2
    v4 = append(v4, v2)
    if true {
        if true {
            var v22 int64 = v6
            var v23 uint8 = 183
            v24 := GenS1{F0: false, F1: (v1 | v1)}
            _ = v22
            _ = v23
            _ = v24
        }
        fmt.Println(v1, v20, v21)
    }
    v21 = (v21 * v21)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 += uint64(v2)
    fmt.Println(v1, v20)
    if (false || false) {
        fmt.Println(v6, v20, v3)
        v20 = v20
        var v25 float32 = (1e10 / (1.5 + 1.5))
        _ = v25
    }
    var v26 int64 = (v6 - 9223372036854775807)
    var v27 int16 = (v3 * -32768)
    var v28 int = 2147483647
    for fi29 := 0; fi29 < 16; fi29++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi29
    }
    if true {
        for fi30 := 0; fi30 < 12; fi30++ {
            v31 := GenS3{F0: uint(v20)}
            var v32 float32 = 1.5
            _ = fi30
            _ = v31
            _ = v32
        }
        if (!false) {
            fmt.Println(v27, v2, v1)
        }
    } else {
        v27 = 1
        v20 -= (v20 >> ((55) & 63))
    }
    v28 = (v28 % (((v28 + (v28 / ((v28) | 1)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
