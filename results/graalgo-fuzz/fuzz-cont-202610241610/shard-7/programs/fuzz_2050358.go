package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int16 }
func genFn0() uint8 {
    var v4 uint16 = 271
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    var v6 uint = 65535
    v4 += 0
    v7 := GenS1{F0: 0, F1: 2, F2: 265}
    _ = v4
    _ = v6
    _ = v7
    return 124
}

func main() {
    var v1 float32 = float32(-2.25)
    var v2 float64 = (1.5 / 1.5)
    var v3 uint32 = 52
    if (!false) {
        v3 -= v3
    } else {
        var v8 uint32 = v3
        v9 := []int{-2147483648}
        switch 1 {
        case 32767:
            var v10 int16 = -555
            _ = v10
        case 9223372036854775807:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
        _ = v8
        _ = v9
    }
    var v11 uint16 = 466
    var v12 int8 = -1
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS1{F0: -1, F1: v3, F2: 42}
        switch (fi13 << ((65535) & 63)) {
        case 2147483647:
            v1 = v1
        default:
            for fi15 := 0; fi15 < 17; fi15++ {
                switch (^(fi13)) {
                case -1:
                    v16 := []int{fi13}
                    _ = v16
                case 32767:
                    v11 -= uint16(v3)
                case -690:
                    v17 := []int{42, (fi15 / ((fi15) | 1))}
                    _ = v17
                }
                switch fi15 {
                case 288:
                    if (-1 < fi13) {
                        v12 = (v12 * v12)
                        fmt.Println(v3, v3)
                        v1 = (v1 * (float32(v2) - -2.25))
                    }
                case 128:
                    var v18 int64 = -9223372036854775808
                    _ = v18
                default:
                    v19 := GenS1{F0: 2147483647, F1: v3, F2: 692}
                    _ = v19
                }
                var v20 uint8 = 255
                _ = fi15
                _ = v20
            }
        }
        _ = fi13
        _ = v14
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        for fi21 := 0; fi21 < 3; fi21++ {
            fmt.Println(fi21)
            var v22 uint = 12
            _ = fi21
            _ = v22
        }
    case 2147483647:
        v2 = (1e10 - v2)
    }
    v23 := GenS1{F0: 42, F1: 2, F2: 1}
    var v24 uint32 = v3
    fmt.Println(v2)
    v25 := GenS1{F0: -961, F1: v3, F2: -32768}
    for fi26 := 0; fi26 < 10; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            fmt.Println(v24)
            fmt.Println(v12)
            if (true || (((-2147483648 >= (fi27 + (fi27 % (((fi26 * fi26)) | 1)))) && true) || false)) {
                v3 -= v3
            } else {
                var v28 int64 = -9223372036854775808
                v12 = v12
                _ = v28
            }
            _ = fi27
        }
        _ = fi26
    }
    var v29 uint32 = v24
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
}
