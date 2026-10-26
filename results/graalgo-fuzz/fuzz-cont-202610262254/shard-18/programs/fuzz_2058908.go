package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS3 struct { F0 int8; F1 uint32; F2 int64 }
func genFn0(p0 int16, p1 bool, p2 int64) int {
    fmt.Println(p2)
    v6 := GenS1{F0: 2147483648, F1: -638, F2: 178}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 128
}
func genFn1() uint64 {
    v7 := []int{}
    var v8 uint16 = 2
    _ = len(v7)
    for fi9 := 0; fi9 < 6; fi9++ {
        switch 0 {
        case 192:
            v10 := GenS1{F0: 2147483647, F1: 0, F2: 237}
            _ = v10
        case 9223372036854775807:
            if false {
                v8 = v8
            }
        case -672:
            _ = len(v7)
        }
        v11 := GenS2{F0: 858, F1: 2, F2: -471}
        var v12 uint16 = v8
        _ = fi9
        _ = v11
        _ = v12
    }
    fmt.Println(v8)
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []string{("a" + ((("" + "a") + "") + ("foo bar" + "hello"))), "\n\t", "", ("foo bar" + "hello")}
    var v2 int64 = 1
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 2
    if true {
        _ = len(v1)
    } else {
        v4[100] = 2147483648
        var v13 float64 = 3.14159
        var v14 float64 = (float64(0.1) + (v13 * ((-2.25 / v13) * 1e10)))
        _ = v13
        _ = v14
    }
    v2 -= (v2 / (((v2 * v2)) | 1))
    if (127 == 42) {
        for fi15 := 0; fi15 < 19; fi15++ {
            v5 += v5
            v16 := GenS1{F0: 42, F1: -4, F2: v5}
            v4[0] = 1
            _ = fi15
            _ = v16
        }
        var v17 uint8 = v5
        v18 := GenS2{F0: 0, F1: 32767, F2: -1}
        _ = v17
        _ = v18
    } else {
        switch "" {
        case "hello":
            var v19 int = 4294967295
            _ = v19
        case "foo bar":
            v2 -= v2
        default:
            fmt.Println(v5)
        }
        fmt.Println(v5)
        var v20 int64 = (v2 >> ((65535) & 63))
        _ = v20
    }
    v21 := GenS2{F0: 58, F1: 230, F2: 0}
    var v22 int32 = -2147483648
    if false {
        var v23 uint8 = (v5 | (^(v5)))
        v24 := GenS2{F0: 18446744073709551615, F1: 1, F2: v22}
        _ = v23
        _ = v24
    }
    var v25 int = 2147483647
    if false {
        v5 += 255
        var v26 string = "hello"
        _ = v26
    }
    if false {
        var v27 string = (("\n\t" + "\n\t") + "hello")
        _ = v27
    }
    v28 := GenS1{F0: 4294967295, F1: -32768, F2: (v5 + (v5 + v5))}
    for fi29 := 0; fi29 < 17; fi29++ {
        v5 = v5
        _ = fi29
    }
    switch genFn0(int16(v2), ((v22 * v22) >= (v22 + (v22 >> ((18446744073709551615) & 63)))), (v2 - 135)) {
    case 0:
        if (!true) {
            _ = len(v1)
        } else {
            if true {
                fmt.Println(v5, v5, v2)
                v5 += (v5 + (v5 & (v5 & v5)))
            }
            var v30 uint32 = 65535
            _ = v30
        }
    default:
        var v31 float32 = 3.14159
        _ = v31
    }
    v32 := GenS1{F0: 2147483647, F1: 32767, F2: v5}
    fmt.Println(v22, v2, v2)
    for fi33 := 0; fi33 < 12; fi33++ {
        v2 -= 42
        _ = fi33
    }
    fmt.Println(v22)
    v34 := GenS3{F0: 84, F1: 1, F2: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v28
    _ = v3
    _ = v32
    _ = v34
    _ = v4
    _ = v5
}
