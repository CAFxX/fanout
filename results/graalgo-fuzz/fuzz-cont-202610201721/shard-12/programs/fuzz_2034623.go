package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }

func main() {
    v1 := []int{}
    v2 := []int{-46, 4294967295, 306}
    var v3 uint = 197
    var v4 uint8 = 1
    var v5 int32 = 42
    v6 := []int{42, 2147483647, 32768, 1}
    var v7 uint16 = 32767
    switch 222 {
    case 128:
        _ = len(v6)
    default:
        if ((0.1 != (1e10 + 0.0)) || (127 <= 42)) {
            if false {
                v1 = append(v1, int(v4))
            }
        }
    }
    v5 = (v5 / ((v5) | 1))
    v5 = (v5 - (-(v5)))
    var v8 int64 = int64(v5)
    v4 = v4
    if false {
        v5 += (v5 * (v5 << (((^(v3))) & 63)))
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            var v10 uint64 = 707
            _ = len(v2)
            v11 := GenS1{F0: v7, F1: v5, F2: uint(v4)}
            _ = v10
            _ = v11
        } else {
            switch (-(fi9)) {
            case 42:
                if (true || (0.0 >= float32((0.1 - (1e10 + 0.1))))) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v12 int32 = (v5 | 0)
                    if true {
                        v13 := GenS1{F0: v7, F1: v5, F2: v3}
                        _ = v13
                    } else {
                        var v14 uint32 = 2147483648
                        var v15 float32 = (float32(-2.25) - (((float32(-2.25) - 3.14159) - 3.14159) * (0.1 * (-2.25 + 0.0))))
                        _ = v14
                        _ = v15
                    }
                    _ = v12
                } else {
                    v3 += v3
                    fmt.Println(v7, fi9, v8)
                }
            }
            v4 = v4
            v4 += (v4 + (v4 >> (((v3 << (((v3 / ((v3) | 1))) & 63))) & 63)))
        }
        var v16 uint = 429
        _ = fi9
        _ = v16
    }
    if (1.5 <= 1e10) {
        for fi17 := 0; fi17 < 17; fi17++ {
            v4 += v4
            _ = fi17
        }
        _ = len(v1)
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 int32 = v5
        _ = fi18
        _ = v19
    }
    var v20 int16 = 307
    for fi21 := 0; fi21 < 5; fi21++ {
        fmt.Println(v7, v20, v8)
        v4 = 2
        v3 -= (^(v3))
        _ = fi21
    }
    var v22 int8 = -128
    if len(v6) > 0 {
        v6[0] = -1
    }
    var v23 uint = v3
    var v24 float64 = (0.0 * float64((float32(((0.1 / float64(0.1)) + 0.0)) - 1e10)))
    _ = len(v2)
    switch (("" + "a") + "\n\t") {
    case "\n\t":
        v1 = append(v1, 32767)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
