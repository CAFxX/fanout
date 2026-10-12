package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 bool; F2 int16 }
func genFn0(p0 int32, p1 uint, p2 bool) string {
    fmt.Println(p1, p0, p1)
    var v6 int32 = (p0 << ((p1) & 63))
    v7 := GenS1{F0: 18446744073709551615, F1: 127, F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "hello"
}
func genFn1() uint8 {
    var v8 uint64 = 1
    v9 := []int{101}
    fmt.Println(v8)
    switch (("hello" + ("a" + "")) + "\n\t") {
    case "hello":
        var v10 bool = false
        _ = v10
    case "":
        v8 += 42
    case "foo bar":
        if (1 == 0) {
            switch 1 {
            case 1:
                var v11 uint32 = 375
                _ = v11
            default:
                if len(v9) > 0 {
                    v9[0] = 2147483647
                }
            }
        }
    default:
        var v12 uint8 = 117
        _ = v12
    }
    _ = v8
    _ = v9
    return 92
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (float32(0.0) * (((float32(3.14159) - -2.25) - 1.5) - (0.1 * 1e10)))
    var v3 float32 = (3.14159 * ((v2 - v2) * float32(0.0)))
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    v13 := []int{-2147483648, 678, 809, -2147483648}
    var v14 int8 = -8
    v2 = float32((((1e10 - 0.0) / 3.14159) - 3.14159))
    v15 := v5[""]
    if ((442 < 0) && false) {
        switch "hello" {
        case "foo bar":
            var v16 int16 = 42
            _ = v16
        default:
            v3 = -2.25
        }
    }
    if (((v2 + -2.25) / v3) > (((v3 + (v2 * float32(float64(v2)))) - (v2 - -2.25)) / v2)) {
        switch 765 {
        case -622:
            v17 := GenS2{F0: 958}
            _ = v17
        case 4294967295:
            switch v15 {
            case "\n\t":
                v3 = v2
            case "foo bar":
                v18 := GenS1{F0: 42, F1: v14, F2: 255}
                _ = v18
            default:
                if false {
                    var v19 int16 = 42
                    _ = len(v13)
                    _ = v19
                } else {
                    v20 := GenS1{F0: 995, F1: 125, F2: 1}
                    _ = v20
                }
            }
        }
        switch 1 {
        case 42:
            if (42 < 1) {
                var v21 int = 32768
                _ = v21
            } else {
                var v22 uint32 = 2147483647
                v1[v15] = true
                _ = v22
            }
        }
    } else {
        for fi23 := 0; fi23 < 20; fi23++ {
            v24 := GenS3{F0: -849}
            _ = fi23
            _ = v24
        }
        var v25 int16 = 42
        v26 := GenS4{F0: 32767, F1: true, F2: v25}
        _ = v25
        _ = v26
    }
    var v27 int = 503
    var v28 int = v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
