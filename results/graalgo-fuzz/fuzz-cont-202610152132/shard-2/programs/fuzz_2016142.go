package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS2 struct { F0 bool; F1 bool; F2 int16 }
func genFn0(p0 string, p1 uint32, p2 uint) bool {
    if false {
        var v4 uint16 = 52
        p1 = 544
        _ = v4
    } else {
        v5 := []int{-179, -2147483648}
        p0 = p0
        fmt.Println(p0, p0, p0)
        _ = v5
    }
    p2 += (p2 | (p2 % ((p2) | 1)))
    fmt.Println(p2, p0)
    var v6 float32 = ((0.1 / 1.5) / 1.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return true
}
func genFn1(p0 uint8) int32 {
    v7 := GenS1{F0: p0, F1: -626, F2: 2}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 651
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]string)
    var v3 uint32 = 2147483648
    fmt.Println(v3)
    var v8 float32 = (((0.0 + ((0.0 - float32(-2.25)) + ((3.14159 / 1.5) + 3.14159))) / 1.0) * float32(3.14159))
    for fi9 := 0; fi9 < 7; fi9++ {
        v8 = (float32(((1.5 / 1.5) + 0.0)) * v8)
        v8 = (1e10 / v8)
        fmt.Println(fi9, v8)
        _ = fi9
    }
    v10 := GenS2{F0: true, F1: (1.5 != 0.0), F2: -1}
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 uint = 65535
        v13 := GenS1{F0: 42, F1: 1000, F2: 480}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int16 = 42
    switch -1 {
    case 1:
        v8 = v8
    }
    v14 += v14
    v1[((((("hello" + ("" + (("\n\t" + ("\n\t" + "foo bar")) + ("foo bar" + "foo bar")))) + ("\n\t" + "")) + "a") + "a") + "hello")] = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v8
}
