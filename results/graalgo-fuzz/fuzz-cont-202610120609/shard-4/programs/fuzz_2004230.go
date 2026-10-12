package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint }
type GenS2 struct { F0 int }

func main() {
    v1 := []int{0, 42, 128, -9223372036854775808}
    v2 := []uint32{1, 509, 101, 323}
    var v3 int64 = 1
    var v4 uint16 = 0
    if (35 >= 127) {
        if len(v2) > 0 {
            v2[0] = 42
        }
    } else {
        var v5 int32 = -2147483648
        v3 += -1
        _ = v5
    }
    v4 -= 32767
    switch 42 {
    case 868:
        if len(v1) > 0 {
            v1[0] = -1
        }
    }
    if ((v3 >= v3) || false) {
        v1 = append(v1, 128)
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 int32 = 0
            var v8 uint = 2
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 127, F1: (false || false), F2: 255}
        _ = v9
    } else {
        v3 = (v3 ^ v3)
        var v10 float32 = 1e10
        if true {
            v11 := GenS1{F0: -1, F1: true, F2: 930}
            _ = v11
        } else {
            var v12 uint16 = 542
            v13 := GenS1{F0: -145, F1: false, F2: 562}
            _ = v12
            _ = v13
        }
        _ = v10
    }
    var v14 float32 = 3.14159
    var v15 int = 32768
    v14 += v14
    v15 -= -9223372036854775808
    var v16 int32 = 2147483647
    v16 += 42
    for fi17 := 0; fi17 < 16; fi17++ {
        v14 -= v14
        _ = fi17
    }
    fmt.Println(v14, v4)
    v14 = (v14 * (v14 + (v14 + 0.1)))
    v18 := GenS1{F0: (v15 % ((-2147483648) | 1)), F1: false, F2: 65535}
    var v19 uint8 = 1
    switch (v15 % ((835) | 1)) {
    case -2147483648:
        if false {
            fmt.Println(v14)
            var v20 int = v15
            v21 := GenS1{F0: (v15 << ((uint(v4)) & 63)), F1: false, F2: 280}
            _ = v20
            _ = v21
        } else {
            var v22 uint = 0
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := GenS2{F0: (-(v15))}
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
