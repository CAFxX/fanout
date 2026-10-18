package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 int = 0
    v3 += 128
    if (-128 >= -1) {
        v3 += (v3 + (v3 / ((v3) | 1)))
        var v4 int64 = 0
        fmt.Println(v4)
        _ = v4
    }
    v3 += (v3 & v3)
    v2["a"] = (v3 ^ v3)
    if (1e10 <= 1e10) {
        var v5 int32 = -986
        _ = v5
    }
    var v6 int32 = -2147483648
    v3 = v3
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int32 = 42
        v9 := GenS1{F0: (^(v3)), F1: v3, F2: 2}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: (-(v3)), F1: v3, F2: 42}
    var v11 float64 = (float64(0.0) + 1e10)
    v12 := []int{(v3 / ((0) | 1)), (v3 + 42), v3}
    var v13 int = (-(v3))
    v2[(("hello" + ("\n\t" + (("hello" + "hello") + ("" + "\n\t")))) + ("" + ("foo bar" + ("a" + ""))))] = v13
    var v14 uint8 = 128
    if (!false) {
        fmt.Println(v11, v6, v14)
        v3 += v3
    } else {
        switch (("\n\t" + "") + "") {
        case "\n\t":
            v15 := GenS1{F0: v3, F1: (-(v13)), F2: 2}
            _ = v15
        case "hello":
            v13 = (v3 % (((^(v3))) | 1))
        case "foo bar":
            v16 := GenS1{F0: v3, F1: (v13 / ((128) | 1)), F2: 862}
            _ = v16
        default:
            v13 = v13
        }
        if (((("foo bar" + "a") + "foo bar") != ((("\n\t" + "hello") + "a") + "foo bar")) && (((true && false) || (true || false)) || (true || (!false)))) {
            var v17 uint8 = v14
            _ = v17
        } else {
            v6 = v6
        }
    }
    var v18 uint8 = 195
    for fi19 := 0; fi19 < 3; fi19++ {
        switch fi19 {
        case 32768:
            var v20 bool = ((v11 * v11) >= float64(0.0))
            _ = v20
        case 1:
            v21 := GenS1{F0: fi19, F1: (v3 + v3), F2: 4294967295}
            _ = v21
        case 756:
            v1[0] = 435
        default:
            v22 := GenS1{F0: v3, F1: (v13 - (-(v13))), F2: uint32(v14)}
            _ = v22
        }
        if len(v12) > 0 {
            v12[0] = fi19
        }
        if (148 <= 2147483648) {
            for fi23 := 0; fi23 < 2; fi23++ {
                if true {
                    fmt.Println(fi23, fi23)
                }
                var v24 int8 = 42
                _ = fi23
                _ = v24
            }
        } else {
            var v25 uint16 = 255
            var v26 int64 = -222
            if (v26 < int64(v6)) {
                v27 := GenS1{F0: v13, F1: (v3 * v3), F2: 2}
                _ = v27
            }
            _ = v25
            _ = v26
        }
        _ = fi19
    }
    switch v3 {
    case 127:
        v6 -= v6
    case 128:
        var v28 bool = true
        _ = v28
    case 15:
        v13 = -9223372036854775808
    default:
        var v29 int8 = 42
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v6
}
