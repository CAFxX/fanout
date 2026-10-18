package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
func genFn0(p0 int16) uint {
    for fi6 := 0; fi6 < 2; fi6++ {
        switch -2147483648 {
        case -2147483648:
            v7 := []int{(fi6 - 32767), (fi6 - fi6)}
            _ = v7
        default:
            fmt.Println(fi6, fi6)
        }
        var v8 float32 = 3.14159
        fmt.Println(p0, p0)
        _ = fi6
        _ = v8
    }
    v9 := []int{-289, 161}
    switch 32768 {
    case 9223372036854775807:
        if (835 < 2) {
            var v10 int64 = 42
            var v11 float32 = ((0.1 / 1e10) - float32(0.1))
            p0 += 1000
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 int32 = 0
                v9 = append(v9, (fi12 << ((255) & 63)))
                var v14 float32 = -2.25
                _ = fi12
                _ = v13
                _ = v14
            }
            if len(v9) > 0 {
                v9[0] = -2147483648
            }
        }
    case 10:
        p0 = p0
    default:
        p0 = p0
    }
    _ = p0
    _ = v9
    return 18446744073709551615
}
func genFn1() int32 {
    var v15 int8 = 127
    var v16 uint32 = 810
    fmt.Println(v15, v15)
    v16 += (v16 * (v16 << ((genFn0(0)) & 63)))
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int64 = 263
        var v19 uint16 = 944
        v20 := []int{}
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    _ = v15
    _ = v16
    return -2147483648
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    var v3 int8 = -1
    v4 := []uint32{4294967295}
    v5 := []int{-921}
    if (42 >= -1) {
        var v21 bool = (4294967295 > 58)
        if ((false && v21) && v21) {
            v5 = append(v5, 32767)
        }
        _ = v21
    }
    var v22 float32 = float32(((3.14159 + 1e10) + 1.5))
    if (false || true) {
        fmt.Println(v22, v3)
        for fi23 := 0; fi23 < 3; fi23++ {
            if true {
                v24 := GenS1{F0: 144, F1: 24}
                _ = v24
            } else {
                var v25 int = (fi23 & (fi23 << ((494) & 63)))
                _ = v25
            }
            _ = fi23
        }
    }
    v22 -= ((-2.25 * (v22 / 1e10)) + 0.0)
    var v26 int32 = 0
    v26 = genFn1()
    var v27 float32 = v22
    if false {
        v3 += 127
        v28 := GenS1{F0: 0, F1: (-(v26))}
        for fi29 := 0; fi29 < 20; fi29++ {
            v30 := v1[("\n\t" + "\n\t")]
            v31 := GenS1{F0: 128, F1: (v26 / ((v26) | 1))}
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    }
    v4 = append(v4, 1)
    switch 42 {
    case 127:
        v22 = 1e10
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        fmt.Println(fi32, v27, v22)
        _ = fi32
    }
    _ = len(v1)
    fmt.Println(v3, v22)
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 uint64 = 18446744073709551615
        for fi35 := 0; fi35 < 8; fi35++ {
            switch ("hello" + "\n\t") {
            case "hello":
                v3 += v3
            case "\n\t":
                if true {
                    v26 = v26
                }
            case "a":
                var v36 uint32 = 0
                _ = v36
            default:
                v27 = (v27 * 0.0)
            }
            _ = fi35
        }
        v22 = (((1e10 - ((v27 / v22) * -2.25)) + v27) / 1.5)
        _ = fi33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
