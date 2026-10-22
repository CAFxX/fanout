package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint }
func genFn0(p0 uint, p1 int8) uint32 {
    for fi5 := 0; fi5 < 12; fi5++ {
        p1 = -1
        _ = fi5
    }
    var v6 float32 = 3.14159
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{(fi7 & (-(fi7)))}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 2
}

func main() {
    v1 := []uint32{1, 4294967295, 2147483648}
    v2 := make(map[int]int)
    v3 := []int{-9223372036854775808, 688, 9223372036854775807, -332}
    v4 := make(map[int]int)
    var v9 int32 = 0
    if (!false) {
        v9 = v9
        if (42 >= 9223372036854775807) {
            v10 := GenS1{F0: 127}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := GenS2{F0: 0, F1: 42}
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 3; fi13++ {
                v9 += v9
                switch "foo bar" {
                case "foo bar":
                    if (-590 <= 1) {
                        var v14 int8 = -128
                        _ = v14
                    }
                default:
                    v9 -= v9
                }
                _ = fi13
            }
        }
        v4[32767] = 4294967295
    } else {
        v15 := GenS1{F0: 27}
        _ = v15
    }
    _ = len(v4)
    v9 -= -166
    if (2147483648 != 261) {
        var v16 int = 32768
        _ = v16
    } else {
        switch 0 {
        case 32768:
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            switch "" {
            case "":
                v18 := v4[32767]
                _ = v18
            case "a":
                var v19 uint64 = 2
                _ = v19
            case "\n\t":
                var v20 int16 = 42
                _ = v20
            }
            v9 += 0
            _ = fi17
        }
    }
    switch 32767 {
    case 42:
        var v21 int = 48
        _ = v21
    default:
        var v22 int16 = 1
        _ = v22
    }
    v23 := GenS3{F0: 1, F1: genFn0(1, 0), F2: 385}
    v24 := GenS2{F0: 42, F1: 127}
    fmt.Println(v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
