package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 string }
func genFn0(p0 uint, p1 string, p2 uint64) int16 {
    p1 = "\n\t"
    v6 := []int{-9223372036854775808, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    var v3 uint8 = 128
    var v4 uint16 = 52
    var v5 uint8 = v3
    var v7 float64 = 1e10
    v8 := []int{127, 9223372036854775807, 2147483647, -2147483648}
    var v9 bool = (!((^(v3)) <= (v3 % ((v3) | 1))))
    var v10 float32 = ((1.5 + ((0.1 * 0.0) * (((3.14159 / ((1.5 / (1e10 * (3.14159 / 1.0))) - 0.0)) - (0.1 * float32(0.0))) - 0.0))) * 1.5)
    v11 := GenS1{F0: "hello", F1: 573, F2: (-118 <= 1)}
    v10 = (v10 / v10)
    if (100 >= -17) {
        if ((v7 * (0.0 + 1e10)) != v7) {
            if false {
                fmt.Println(v4, v9)
            }
            var v12 int = 9223372036854775807
            var v13 uint8 = v5
            _ = v12
            _ = v13
        } else {
            fmt.Println(v7, v10, v7)
            v14 := GenS1{F0: ("foo bar" + "\n\t"), F1: 1000, F2: true}
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 int8 = 73
                v17 := GenS2{F0: 179}
                var v18 int16 = 0
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = v14
        }
        if ((-87 == 1) || true) {
            switch -628 {
            case 42:
                var v19 int = 42
                _ = v19
            case -2147483648:
                var v20 int64 = 0
                _ = v20
            case -240:
                var v21 uint = 2
                _ = v21
            default:
                var v22 uint8 = (v3 + 0)
                _ = v22
            }
            fmt.Println(v3, v3, v9)
            v5 -= uint8(v4)
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            _ = len(v2)
            v4 = v4
            _ = fi23
        }
    } else {
        var v24 float64 = 1.5
        var v25 uint32 = 42
        v26 := GenS1{F0: ("a" + ((((("foo bar" + (("" + "\n\t") + "a")) + "") + "foo bar") + ("hello" + "")) + "")), F1: genFn0(335, "foo bar", 0), F2: v9}
        _ = v24
        _ = v25
        _ = v26
    }
    v27 := v1[-311]
    v28 := GenS3{F0: -650, F1: (("foo bar" + "a") + ((("hello" + ("" + "")) + "a") + "hello"))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
