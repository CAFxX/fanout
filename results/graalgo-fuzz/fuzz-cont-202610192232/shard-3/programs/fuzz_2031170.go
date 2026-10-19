package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0(p0 int8) int {
    switch -2147483648 {
    case -9223372036854775808:
        var v6 int = 485
        _ = v6
    case -2147483648:
        if (true && true) {
            v7 := []int{}
            fmt.Println(p0)
            v8 := GenS1{F0: 499, F1: 1}
            _ = v7
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 10; fi9++ {
            p0 -= -128
            fmt.Println(fi9)
            var v10 bool = (!true)
            _ = fi9
            _ = v10
        }
    }
    p0 = (p0 & 127)
    p0 += (p0 + (p0 / (((^(p0))) | 1)))
    _ = p0
    return 338
}
func genFn1(p0 int16, p1 string) int32 {
    var v11 int64 = -1
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int16 = p0
        fmt.Println(v13, v11, p1)
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v11 = v11
        v15 := GenS2{F0: 255, F1: (42 != 1)}
        var v16 uint64 = 4294967295
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: v11, F1: genFn0(93)}
    _ = p0
    _ = p1
    _ = v11
    _ = v17
    return 2147483647
}

func main() {
    var v1 float32 = float32(((-2.25 + (1e10 * 0.1)) + float64(((3.14159 * (1e10 * float32((0.0 + 3.14159)))) - (-2.25 / 1.0)))))
    var v2 uint16 = 2
    var v3 int8 = 42
    v4 := make(map[int]bool)
    var v5 int = 32768
    v3 += 42
    v3 = v3
    var v18 string = "\n\t"
    v3 = (v3 % (((v3 / ((0) | 1))) | 1))
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 4; fi20++ {
            delete(v4, (fi20 | (-(fi20))))
            _ = fi20
        }
        fmt.Println(fi19, v3)
        _ = fi19
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        v18 = ""
        _ = fi21
    }
    v22 := GenS2{F0: (^(v2)), F1: (false && true)}
    var v23 bool = (("a" + v18) >= v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
