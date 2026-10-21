package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int32; F1 int64; F2 string }
func genFn0(p0 uint8, p1 uint32, p2 int8) int {
    p1 += 4294967295
    if true {
        p2 += p2
        v4 := []int{32768, 32767}
        _ = v4
    } else {
        v5 := []int{42, -2147483648}
        if (!(-480 >= 925)) {
            var v6 float32 = 0.1
            _ = v6
        } else {
            switch "" {
            case "a":
                v7 := GenS1{F0: -2147483648, F1: (p0 % (((p0 * p0)) | 1))}
                _ = v7
            case "foo bar":
                p1 = p1
            }
        }
        p1 -= (p1 | p1)
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 391
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 59
    var v3 int = 1
    var v8 string = ((("" + "foo bar") + ("foo bar" + "")) + "a")
    v9 := []int{v3, v3, (v3 / ((-774) | 1))}
    v3 += (v3 - 9223372036854775807)
    var v10 float32 = (1e10 + (1.5 - (1e10 / 0.1)))
    v3 = v3
    var v11 uint64 = 2
    v2 -= (v2 % (((v2 | (^(v2)))) | 1))
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 float64 = (1.5 / 3.14159)
        if false {
            v14 := GenS1{F0: 9223372036854775807, F1: v2}
            for fi15 := 0; fi15 < 11; fi15++ {
                switch v8 {
                case "hello":
                    v2 = v2
                }
                v16 := GenS1{F0: 169, F1: (v2 - (v2 ^ 0))}
                v13 = float64((0.0 + 0.1))
                v17 := GenS1{F0: fi12, F1: 128}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
        var v18 uint32 = 2147483648
        _ = fi12
        _ = v13
        _ = v18
    }
    switch (v8 + (v8 + v8)) {
    case "foo bar":
        v8 = (v8 + "hello")
    default:
        var v19 int32 = -2147483648
        _ = v19
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        if (-2147483648 == 1) {
            v21 := GenS2{F0: -2147483648, F1: 0, F2: v8}
            v1[(v3 | (-(v3)))] = 2147483648
            _ = v21
        } else {
            var v22 uint8 = (v2 / ((2) | 1))
            v2 = (v2 | v2)
            var v23 string = v8
            _ = v22
            _ = v23
        }
        if len(v9) > 0 {
            v9[0] = (fi20 + 127)
        }
        _ = fi20
    }
    switch v8 {
    case "hello":
        if true {
            var v24 string = (v8 + ((v8 + (v8 + "hello")) + "foo bar"))
            _ = v24
        }
    case "":
        if false {
            switch (v8 + v8) {
            case "":
                v11 = (v11 >> ((71) & 63))
            }
            var v25 int64 = 298
            delete(v1, int(v11))
            _ = v25
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
