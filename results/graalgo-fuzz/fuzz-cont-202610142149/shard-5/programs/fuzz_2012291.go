package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint8; F2 int64 }
type GenS3 struct { F0 string; F1 uint; F2 uint32 }
func genFn0() int64 {
    var v6 uint16 = 134
    v6 -= v6
    _ = v6
    return -796
}
func genFn1(p0 uint8, p1 string, p2 uint64) int32 {
    if (42 > 15) {
        v7 := []int{165, 322, 2147483647}
        _ = v7
    } else {
        v8 := []int{-2147483648}
        _ = v8
    }
    var v9 int = 9223372036854775807
    v10 := []int{v9, (v9 * (v9 / ((int(p0)) | 1)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -1
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    var v5 int = 869
    var v11 uint64 = v1
    v12 := GenS1{F0: 204}
    var v13 uint8 = 127
    v14 := GenS1{F0: v13}
    if true {
        v15 := GenS2{F0: genFn1(uint8(v1), "a", v11), F1: v13, F2: -9223372036854775808}
        v16 := []int{0, -2147483648, (v5 & v5)}
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 14}
    switch "foo bar" {
    case "\n\t":
        var v18 uint16 = 2
        _ = v18
    case "hello":
        var v19 int32 = 42
        _ = v19
    default:
        v11 = (^(v11))
    }
    v20 := GenS2{F0: 1, F1: 146, F2: 491}
    switch 9223372036854775807 {
    case 128:
        v5 = 0
    default:
        v1 -= (v11 + 232)
    }
    v1 = 2
    if (!(2147483647 > v5)) {
        v13 = 255
        for fi21 := 0; fi21 < 19; fi21++ {
            if ((1 == 0) && (!false)) {
                if (true && true) {
                    var v22 int16 = -1
                    _ = v22
                } else {
                    v23 := v4[""]
                    v24 := []int{(v5 >> ((2) & 63))}
                    _ = v23
                    _ = v24
                }
            } else {
                if true {
                    v25 := GenS3{F0: "hello", F1: 4294967295, F2: 65535}
                    _ = v25
                }
            }
            fmt.Println(v13, fi21, v1)
            fmt.Println(v1, v13)
            _ = fi21
        }
        v3[1] = (-2.25 >= (0.1 / ((0.1 / 0.1) - 0.1)))
    }
    var v26 float64 = -2.25
    fmt.Println(v13)
    var v27 int = v5
    switch v27 {
    case -9223372036854775808:
        v28 := GenS2{F0: 2147483647, F1: v13, F2: int64(v13)}
        _ = v28
    case 127:
        for fi29 := 0; fi29 < 6; fi29++ {
            v2[fi29] = 314
            _ = fi29
        }
    default:
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
