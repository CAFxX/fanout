package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 string; F1 uint; F2 int32 }
func genFn0() int32 {
    var v4 int = 127
    v4 -= 0
    _ = v4
    return 2147483647
}
func genFn1(p0 uint64, p1 uint16) int8 {
    p0 += p0
    p1 += 33
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 0.1
    var v3 string = "hello"
    v5 := []int{42, 9223372036854775807, 913, -2147483648}
    v3 = v3
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int64 = 876
        v3 = "a"
        v7 += v7
        _ = fi6
        _ = v7
    }
    if (!false) {
        v8 := GenS1{F0: 657, F1: 4294967295, F2: 0}
        var v9 int8 = 127
        _ = v8
        _ = v9
    } else {
        v2 = (-2.25 / 0.1)
    }
    v10 := GenS2{F0: v3, F1: 678, F2: 0}
    switch v3 {
    case "hello":
        v2 = ((v2 / v2) / float64(3.14159))
    default:
        var v11 int8 = 0
        _ = v11
    }
    v3 = ((("\n\t" + "") + v3) + "foo bar")
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(v3)
        if true {
            var v13 int64 = 0
            _ = v13
        } else {
            if (65535 != 42) {
                fmt.Println(v2, v2, v3)
                var v14 uint = 46
                v3 = v3
                _ = v14
            }
            v2 -= ((((((-2.25 + 3.14159) / v2) * (v2 - 1e10)) * (v2 * float64((-2.25 + float32(0.1))))) / (-2.25 + -2.25)) + (v2 - ((0.0 * v2) * 1.5)))
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 float64 = v2
                _ = fi15
                _ = v16
            }
        }
        _ = fi12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
}
