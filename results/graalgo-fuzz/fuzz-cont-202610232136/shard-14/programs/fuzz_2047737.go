package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 string; F2 int }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int; F1 bool }
func genFn0(p0 int64, p1 int64, p2 int32) uint {
    if false {
        v4 := GenS1{F0: 32767}
        _ = v4
    } else {
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 bool = ((0.1 > 3.14159) || true)
            v7 := []int{fi5, fi5, fi5}
            switch 1 {
            case 4294967295:
                fmt.Println(v6)
            default:
                var v8 uint32 = 0
                _ = v8
            }
            _ = fi5
            _ = v6
            _ = v7
        }
        var v9 int8 = -128
        _ = v9
    }
    fmt.Println(p1, p0)
    v10 := GenS1{F0: 0}
    if false {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int16 = -32768
            v13 := []int{int(p2), -464, (fi11 ^ (fi11 | (fi11 % (((fi11 - (fi11 - -647))) | 1))))}
            if true {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                var v14 int16 = (v12 - (v12 ^ (v12 << ((255) & 63))))
                _ = v14
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        var v15 float64 = (0.0 + float64(0.0))
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 65535
}
func genFn1(p0 uint16, p1 string) uint16 {
    fmt.Println(p0, p1)
    p0 = (p0 << ((genFn0(9223372036854775807, -9223372036854775808, 2147483647)) & 63))
    v16 := []int{9223372036854775807}
    v17 := GenS2{F0: true, F1: p1, F2: -2147483648}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 762
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "\n\t"
    v3 := []int{-2147483648, 32767, -824, -1}
    v18 := GenS3{F0: false, F1: v2}
    v19 := GenS4{F0: 32768, F1: (2147483647 >= 1)}
    if true {
        v20 := GenS1{F0: 705}
        var v21 float64 = (1e10 / (((float64(float32(-2.25)) / (float64(1.5) / 3.14159)) - ((0.0 * (float64((-2.25 + 1.5)) * 3.14159)) / 0.1)) / 1.5))
        if (true || (-2.25 > float32(((float64((-2.25 * (3.14159 + (0.1 + 1.5)))) + v21) - float64((1e10 * 1.5)))))) {
            v21 -= (0.1 / ((v21 + float64(1.5)) + 1.5))
            v22 := GenS3{F0: (genFn0(948, 1, 1) == 219), F1: v2}
            if (129 >= 1) {
                if ((-1 != 448) && true) {
                    if len(v3) > 0 {
                        v3[0] = 128
                    }
                }
            }
            _ = v22
        }
        _ = v20
        _ = v21
    }
    if (873 <= 255) {
        for fi23 := 0; fi23 < 13; fi23++ {
            v2 = v2
            v2 = v2
            _ = fi23
        }
    }
    switch "hello" {
    case "":
        switch v2 {
        case "\n\t":
            v24 := GenS1{F0: 0}
            _ = v24
        case "foo bar":
            v2 = v2
        case "hello":
            v25 := GenS1{F0: 42}
            _ = v25
        default:
            var v26 float32 = float32(3.14159)
            _ = v26
        }
    default:
        var v27 float32 = -2.25
        _ = v27
    }
    var v28 string = v2
    var v29 int64 = 9223372036854775807
    v28 = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
