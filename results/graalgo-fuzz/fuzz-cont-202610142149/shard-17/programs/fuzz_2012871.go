package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 uint8; F1 string; F2 int64 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint8, p2 int16) uint16 {
    fmt.Println(p2, p2)
    var v7 uint32 = 2147483647
    p1 = uint8(v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 1
    v4 := make(map[string]string)
    var v5 int16 = -32768
    v6 := []uint32{0, 429}
    var v8 int = 9223372036854775807
    var v9 uint8 = 42
    fmt.Println(v8, v9, v9)
    if ((48 < v9) || (true || false)) {
        v10 := GenS1{F0: "hello"}
        if ((42 == -2147483648) || false) {
            v11 := GenS2{F0: 2, F1: 4294967295, F2: 592}
            _ = v11
        } else {
            v12 := GenS3{F0: (^(v9)), F1: "hello", F2: -581}
            for fi13 := 0; fi13 < 17; fi13++ {
                fmt.Println(v3, v5, fi13)
                if false {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v9 = (v9 | v9)
                        v3 = 42
                        _ = fi14
                    }
                    fmt.Println(v8)
                }
                v15 := GenS4{F0: 42, F1: 127}
                _ = fi13
                _ = v15
            }
            _ = v12
        }
        var v16 uint64 = 42
        _ = v10
        _ = v16
    }
    if true {
        if ((("foo bar" + "a") + ((("a" + "\n\t") + "") + "foo bar")) > "hello") {
            if ((((float64((-2.25 / (((1e10 + 3.14159) + 0.0) * 3.14159))) * 3.14159) + 1.5) <= (float64(((((3.14159 / -2.25) / ((3.14159 * (((3.14159 + 1.5) / (3.14159 / -2.25)) * float32((0.1 / 1.0)))) - 0.1)) - 0.1) / -2.25)) + (0.1 + (((0.0 * (((1e10 + (1e10 + float64((-2.25 + ((-2.25 - 3.14159) / 0.1))))) - 0.0) * (0.1 / 3.14159))) * 1.5) / 1.0)))) || true) {
                v17 := GenS5{F0: 837, F1: genFn0(1, (v9 >> ((90) & 63)), v5)}
                for fi18 := 0; fi18 < 8; fi18++ {
                    v5 = v5
                    var v19 float32 = 0.0
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(v3, v9)
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        v5 += v5
        _ = fi20
    }
    v6 = append(v6, 859)
    if true {
        v5 -= int16(v9)
        switch v8 {
        case 42:
            var v21 uint32 = 0
            _ = v21
        case 18:
            var v22 uint64 = v3
            _ = v22
        case -175:
            v9 = v9
        default:
            v9 = 2
        }
    } else {
        fmt.Println(v5, v3, v5)
        var v23 uint16 = 0
        _ = v23
    }
    v3 -= 18446744073709551615
    var v24 uint16 = 32767
    if false {
        v25 := GenS6{F0: (v9 | (v9 / ((253) | 1)))}
        _ = v25
    }
    var v26 float64 = float64(1.5)
    fmt.Println(v9, v26)
    v3 = 723
    fmt.Println(v24, v8, v26)
    v1[42] = ("" + "a")
    var v27 int = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
