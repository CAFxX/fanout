package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int8, p1 int64, p2 int16) int64 {
    v4 := []int{128, int(p1)}
    p2 -= (-(p2))
    fmt.Println(p2, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p1
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]string)
    var v3 float32 = ((3.14159 * (((float32(1.5) / 1.5) / (1.5 + 0.1)) / (((3.14159 + 1.5) / 1.5) + 3.14159))) + 1.5)
    fmt.Println(v3, v3)
    _ = len(v2)
    var v5 bool = true
    v6 := []int{2147483647, 2147483647}
    if v5 {
        v5 = v5
        v7 := GenS1{F0: 65535, F1: 42, F2: 9223372036854775807}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 uint16 = 660
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: 424, F1: 65535, F2: 9223372036854775807}
            switch "hello" {
            case "foo bar":
                switch "hello" {
                case "\n\t":
                    var v12 float32 = float32(1.5)
                    _ = v12
                case "foo bar":
                    v1 += -1
                case "hello":
                    for fi13 := 0; fi13 < 16; fi13++ {
                        switch ("\n\t" + (((("foo bar" + "\n\t") + "foo bar") + "a") + "")) {
                        case "a":
                            v14 := GenS1{F0: 1, F1: 562, F2: -9223372036854775808}
                            _ = v14
                        default:
                            fmt.Println(fi13)
                        }
                        _ = len(v2)
                        fmt.Println(fi8, v1)
                        _ = fi13
                    }
                default:
                    var v15 int8 = 0
                    _ = v15
                }
            }
            _ = fi8
            _ = v11
        }
    }
    v6 = append(v6, -86)
    var v16 bool = (v5 || (0 <= 664))
    var v17 int = -848
    if v5 {
        if (v16 && (v5 && v16)) {
            v18 := GenS2{F0: 42, F1: "foo bar"}
            _ = v18
        }
    }
    if (2 >= 539) {
        v5 = v16
        fmt.Println(v17, v3, v1)
        var v19 uint8 = 130
        _ = v19
    } else {
        v20 := v2[v17]
        _ = v20
    }
    var v21 float64 = 1e10
    for fi22 := 0; fi22 < 16; fi22++ {
        v23 := GenS3{F0: 1, F1: 42}
        _ = fi22
        _ = v23
    }
    if len(v6) > 0 {
        v6[0] = 42
    }
    switch v17 {
    case -114:
        v6 = append(v6, 32767)
    case 42:
        var v24 uint64 = 2
        _ = v24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        fmt.Println(v16, v5)
        _ = fi25
    }
    delete(v2, 1)
    switch (v17 / ((v17) | 1)) {
    case -9223372036854775808:
        v26 := GenS3{F0: 0, F1: 255}
        _ = v26
    default:
        v5 = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v5
    _ = v6
}
