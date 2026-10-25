package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
func genFn0(p0 int32, p1 bool) bool {
    var v5 int8 = -128
    for fi6 := 0; fi6 < 18; fi6++ {
        if p1 {
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := []int{(fi6 - (fi6 << ((0) & 63))), (-(fi6))}
                var v9 int8 = 42
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 bool = true
            _ = v10
        }
        fmt.Println(p1, v5)
        p1 = (!(127 == v5))
        _ = fi6
    }
    if false {
        fmt.Println(v5, p1, p1)
    }
    _ = p0
    _ = p1
    _ = v5
    return p1
}
func genFn1(p0 uint8, p1 uint8, p2 string) int32 {
    fmt.Println(p1, p1)
    var v11 uint32 = 42
    v12 := GenS1{F0: p2, F1: 825}
    var v13 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return 32
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 1.5
    var v3 uint = 4294967295
    v4 := make(map[uint32]string)
    v3 = 278
    v3 -= 21
    var v14 float64 = (0.0 + ((0.0 + -2.25) / 3.14159))
    fmt.Println(v3)
    v14 += v14
    var v15 uint = v3
    for fi16 := 0; fi16 < 12; fi16++ {
        v3 = v3
        var v17 int16 = -728
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 2
    if (v14 > float64((v2 * v2))) {
        var v19 bool = true
        var v20 int16 = int16(v18)
        v19 = (!true)
        _ = v19
        _ = v20
    } else {
        var v21 uint = v3
        _ = v21
    }
    fmt.Println(v18, v15, v14)
    var v22 int64 = -1
    switch 310 {
    case -9223372036854775808:
        v23 := GenS1{F0: (("a" + "foo bar") + ("a" + "\n\t")), F1: -651}
        _ = v23
    case -328:
        switch "\n\t" {
        case "hello":
            fmt.Println(v2, v22, v3)
        }
    case -272:
        v24 := []int{-798, -2147483648, -1}
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
